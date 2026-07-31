#!/usr/bin/env python3
"""Resolve PSP object undefined symbols to the required PSPSDK archives."""

import argparse
import os
import re
import subprocess
from collections import defaultdict, deque
from pathlib import Path

TOOLCHAIN_SYMBOLS = {
    "abort", "calloc", "free", "malloc", "memalign", "memchr", "memcmp",
    "memcpy", "memmove", "memset", "printf", "putchar", "puts", "realloc",
    "snprintf", "sprintf", "strcat", "strchr", "strcmp", "strcpy", "strlen",
    "strncmp", "strncpy", "strrchr", "strstr", "vsnprintf",
}


def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument("--output", required=True, type=Path)
    parser.add_argument("--require", action="append", default=[])
    parser.add_argument("--verbose", action="store_true")
    parser.add_argument("objects", nargs="+", type=Path)
    return parser.parse_args()


def nm_undefined(nm: Path, objects):
    result = subprocess.run(
        [str(nm), "-u", *map(str, objects)],
        check=True, capture_output=True, text=True,
    )
    symbols = set()
    for line in result.stdout.splitlines():
        match = re.search(r"\bU\s+(\S+)$", line)
        if match:
            symbols.add(match.group(1))
    return symbols


def library_penalty(path: Path):
    name = path.stem.removeprefix("lib")
    penalty = len(name)
    if "driver" in name:
        penalty += 1000
    if "_660" in name or name.endswith("kernel"):
        penalty += 500
    if name == "pspuser":
        penalty += 100
    return penalty


def scan_archives(nm: Path, archives):
    result = subprocess.run(
        [str(nm), "-A", "-g", *map(str, archives)],
        check=True, capture_output=True, text=True,
    )
    definitions = defaultdict(list)
    undefined = defaultdict(set)
    pattern = re.compile(r"^(.+?\.a):([^:]+):(?:[0-9A-Fa-f]+)?\s*([A-Za-zU])\s+(\S+)$")
    for line in result.stdout.splitlines():
        match = pattern.match(line)
        if not match:
            continue
        archive, member, kind, symbol = match.groups()
        key = (Path(archive), member)
        if kind == "U":
            undefined[key].add(symbol)
        else:
            definitions[symbol].append(key)
    return definitions, undefined


def choose_definition(symbol, candidates, selected):
    def score(item):
        archive, _ = item
        name = archive.stem.removeprefix("lib")
        affinity = 0
        normalized_symbol = symbol.lower().replace("_", "")
        normalized_name = name.lower().replace("_", "")
        if normalized_name.startswith("psp"):
            token = normalized_name[3:]
            if token and token in normalized_symbol:
                affinity = -200
        already_selected = -10000 if archive in selected else 0
        return (already_selected + library_penalty(archive) + affinity, archive.name)

    return min(candidates, key=score)


def cmake_name(archive: Path):
    return archive.stem.removeprefix("lib")


def main():
    args = parse_args()
    pspdev = Path(os.environ["PSPDEV"])
    nm = pspdev / "bin/psp-nm"
    lib_dir = pspdev / "psp/sdk/lib"
    archives = sorted(lib_dir.glob("*.a"))

    definitions, member_undefined = scan_archives(nm, archives)
    pending = deque(sorted(nm_undefined(nm, args.objects) | set(args.require)))
    visited_symbols = set()
    visited_members = set()
    selected = set()
    unresolved_psp = set()

    while pending:
        symbol = pending.popleft()
        if symbol in visited_symbols:
            continue
        visited_symbols.add(symbol)
        marker_prefix = "pspsdk_go_require_"
        if symbol.startswith(marker_prefix):
            requested_name = symbol[len(marker_prefix):]
            requested = lib_dir / f"lib{requested_name}.a"
            if not requested.exists():
                raise SystemExit(
                    f"Required PSPSDK library does not exist: {requested}"
                )
            selected.add(requested)
            if args.verbose:
                print(f"{symbol} -> {requested.name} (package requirement)")
            continue
        if symbol.startswith("pspsdk_go_"):
            # Project-local bridge or helper symbol, resolved by a CMake target.
            continue
        if symbol.startswith("__") or symbol in TOOLCHAIN_SYMBOLS:
            continue
        candidates = definitions.get(symbol)
        if not candidates:
            if symbol.startswith(("sce", "psp", "Kprintf")):
                unresolved_psp.add(symbol)
            continue
        archive, member = choose_definition(symbol, candidates, selected)
        if args.verbose:
            print(f"{symbol} -> {archive.name}:{member}")
        selected.add(archive)
        key = (archive, member)
        if key in visited_members:
            continue
        visited_members.add(key)
        pending.extend(member_undefined.get(key, ()))

    if unresolved_psp:
        names = ", ".join(sorted(unresolved_psp))
        raise SystemExit(f"PSPSDK libraries not found for: {names}")

    names = sorted((cmake_name(path) for path in selected))
    args.output.parent.mkdir(parents=True, exist_ok=True)
    content = [
        "# Generated by tools/resolve_pspsdk_libs.py; DO NOT EDIT.",
        "set(PSPSDK_GO_LIBRARIES",
        *[f"    {name}" for name in names],
        ")",
        "",
    ]
    args.output.write_text("\n".join(content))
    print("Resolved PSPSDK libraries: " + (" ".join(names) if names else "(none)"))


if __name__ == "__main__":
    main()
