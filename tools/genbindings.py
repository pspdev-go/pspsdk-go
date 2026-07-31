#!/usr/bin/env python3
"""Generate direct TinyGo bindings for user-mode PSPSDK headers."""

import json
import os
import re
import subprocess
import tempfile
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
PSPDEV = Path(os.environ["PSPDEV"])
INCLUDE = PSPDEV / "psp/sdk/include"
OUT = ROOT / "psp"
SKIP_WORDS = ("_kernel", "forkernel", "driver")
HANDWRITTEN = {"pspctrl.h", "pspdebug.h", "pspdisplay.h", "pspkernel.h"}
KERNEL_LIBRARIES = {
    "pspaudio_kernel.h": "pspaudio_driver",
    "pspctrl_kernel.h": "pspctrl_driver",
    "pspdisplay_kernel.h": "pspdisplay_driver",
    "pspimpose_driver.h": "pspkernel",
    "pspintrman_kernel.h": "pspkernel",
    "pspiofilemgr_kernel.h": "pspkernel",
    "psploadexec_kernel.h": "pspkernel",
    "pspmodulemgr_kernel.h": "pspkernel",
    "pspnand_driver.h": "pspnand_driver",
    "pspstdio_kernel.h": "pspkernel",
    "pspsysmem_kernel.h": "pspkernel",
    "pspthreadman_kernel.h": "pspkernel",
    "psputilsforkernel.h": "pspkernel",
}


def package_name(header: str) -> str:
    name = header[:-2]
    if name.startswith("psp"):
        name = name[3:]
    name = re.sub(r"[^a-zA-Z0-9_]", "_", name).lower()
    if not name or name[0].isdigit():
        name = "h_" + name
    return name


def target_headers():
    return [
        p for p in sorted(INCLUDE.glob("psp*.h"))
    ]


def source_function_names(text: str):
    text = re.sub(r"/\*.*?\*/|//[^\n]*", " ", text, flags=re.S)
    # False positives (macros and calls) are harmless: this set is intersected
    # with Clang FunctionDecl nodes. Keeping this deliberately broad also
    # handles declarations following function-like macros without semicolons.
    pattern = re.compile(r"\b([A-Za-z_]\w*)\s*\(")
    blocked = {"if", "while", "switch", "return", "sizeof"}
    return {m.group(1) for m in pattern.finditer(text) if m.group(1) not in blocked}


def ast_for(header: Path):
    with tempfile.TemporaryDirectory(prefix="pspsdk-go-") as tmp:
        tmp = Path(tmp)
        source = tmp / "header.c"
        preprocessed = tmp / "header.i"
        source.write_text(f"#include <{header.name}>\n")
        with preprocessed.open("wb") as out:
            subprocess.run(
                [str(PSPDEV / "bin/psp-gcc"), "-E", "-P",
                 f"-I{INCLUDE}", str(source)],
                stdout=out, check=True,
            )
        result = subprocess.run(
            ["clang", "-x", "c", "-target", "mipsel-none-eabi",
             "-fsyntax-only", "-Xclang", "-ast-dump=json", str(preprocessed)],
            capture_output=True, check=False, text=True,
        )
        if not result.stdout.strip():
            raise RuntimeError(result.stderr.strip())
        return json.loads(result.stdout)


def walk(node):
    yield node
    for child in node.get("inner", []):
        yield from walk(child)


SCALARS = {
    "void": "",
    "_Bool": "bool",
    "char": "int8", "signed char": "int8", "unsigned char": "uint8",
    "short": "int16", "short int": "int16", "signed short": "int16",
    "unsigned short": "uint16", "unsigned short int": "uint16",
    "int": "int32", "signed int": "int32", "unsigned int": "uint32",
    "long": "int32", "long int": "int32", "unsigned long": "uint32",
    "long long": "int64", "long long int": "int64",
    "unsigned long long": "uint64", "unsigned long long int": "uint64",
    "float": "float32", "double": "float64",
}


def go_type(type_info, is_return=False):
    ctype = type_info.get("desugaredQualType", type_info.get("qualType", ""))
    ctype = re.sub(r"\b(const|volatile|restrict)\b", "", ctype)
    ctype = " ".join(ctype.split())
    if "(*)" in ctype or ctype.endswith("(*)"):
        return "unsafe.Pointer"
    if "*" in ctype or "[" in ctype:
        return "unsafe.Pointer"
    if ctype in SCALARS:
        return SCALARS[ctype]
    # PSPSDK scalar typedefs are 32-bit unless their desugared type says otherwise.
    if re.search(r"(u64|uint64|SceInt64)", ctype):
        return "uint64"
    if re.search(r"(s64|int64|SceOff)", ctype):
        return "int64"
    if is_return and ctype == "void":
        return ""
    return "int32"


def enum_value(node):
    if "value" in node:
        return node["value"]
    for child in node.get("inner", []):
        value = enum_value(child)
        if value is not None:
            return value
    return None


def exported(name: str):
    return name[:1].upper() + name[1:]


def generate(header: Path):
    ast = ast_for(header)
    wanted = source_function_names(header.read_text(errors="ignore"))
    functions = []
    enums = []
    seen_enum = set()
    seen_function = set()
    skipped = []

    for node in walk(ast):
        if node.get("kind") == "EnumConstantDecl":
            name = node.get("name")
            value = enum_value(node)
            if name and value is not None and name not in seen_enum:
                enums.append((name, value))
                seen_enum.add(name)
        if node.get("kind") != "FunctionDecl" or node.get("name") not in wanted:
            continue
        name = node["name"]
        if name in seen_function:
            continue
        seen_function.add(name)
        qual = node.get("type", {}).get("qualType", "")
        if "..." in qual:
            skipped.append(f"{name}: variadic")
            continue
        params = []
        for index, parm in enumerate(
            child for child in node.get("inner", [])
            if child.get("kind") == "ParmVarDecl"
        ):
            pname = parm.get("name") or f"arg{index}"
            pname = re.sub(r"\W", "_", pname)
            if pname in {"type", "func", "map", "range", "var", "chan", "go", "defer", "select", "interface"}:
                pname += "Value"
            params.append((pname, go_type(parm.get("type", {}))))
        result = go_type({"qualType": qual.split(" (", 1)[0]}, True)
        functions.append((name, params, result))

    pkg = package_name(header.name)
    lines = [
        "// Code generated by tools/genbindings.py; DO NOT EDIT.",
        f"// Source: {header.name}",
        "",
        f"// Package {pkg} provides direct TinyGo bindings for {header.name}.",
        f"package {pkg}",
        "",
        'import "unsafe"',
        "",
        "var _ unsafe.Pointer",
    ]
    required_library = KERNEL_LIBRARIES.get(header.name)
    if required_library:
        marker = "pspsdk_go_require_" + required_library
        local_marker = "require" + "".join(
            part.capitalize() for part in required_library.split("_")
        )
        lines += [
            "",
            f"func init() {{ {local_marker}() }}",
            "",
            f"//go:linkname {local_marker} {marker}",
            f"func {local_marker}()",
        ]
    if enums:
        lines += ["", "const ("]
        for name, value in enums:
            lines.append(f"\t{name} = {value}")
        lines.append(")")
    for name, params, result in functions:
        plist = ", ".join(f"{n} {t}" for n, t in params)
        suffix = f" {result}" if result else ""
        lines += [
            "",
            f"//go:linkname {exported(name)} {name}",
            f"func {exported(name)}({plist}){suffix}",
        ]
    if skipped:
        lines += ["", "// Unsupported C declarations:"]
        lines += [f"// - {item}" for item in skipped]
    lines.append("")

    directory = OUT / pkg
    directory.mkdir(parents=True, exist_ok=True)
    (directory / "bindings_gen.go").write_text("\n".join(lines))
    return len(functions), len(enums), skipped


def main():
    report = []
    for header in target_headers():
        if header.name in HANDWRITTEN:
            report.append((header.name, "handwritten", 0, 0, []))
            continue
        try:
            funcs, enums, skipped = generate(header)
            report.append((header.name, package_name(header.name), funcs, enums, skipped))
        except Exception as exc:
            report.append((header.name, "ERROR", 0, 0, [str(exc)]))

    report_path = ROOT / "bindings-report.md"
    rows = [
        "# Generated PSPSDK bindings",
        "",
        "| Header | Package | Functions | Enum constants | Notes |",
        "| --- | --- | ---: | ---: | --- |",
    ]
    for header, pkg, funcs, enums, notes in report:
        rows.append(
            f"| `{header}` | `{pkg}` | {funcs} | {enums} | "
            + "; ".join(notes).replace("|", "\\|") + " |"
        )
    report_path.write_text("\n".join(rows) + "\n")

    errors = [r for r in report if r[1] == "ERROR"]
    print(f"processed {len(report)} headers; errors={len(errors)}")
    if errors:
        raise SystemExit(1)


if __name__ == "__main__":
    main()
