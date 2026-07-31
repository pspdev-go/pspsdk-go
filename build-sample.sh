#!/usr/bin/env bash
set -eu

export GOTOOLCHAIN=local

echo "Using Go binary: $(which go)"
echo "Using Go version: $(GOTOOLCHAIN=local go version)"
echo "Using TinyGo binary: $(which tinygo)"
echo "Using TinyGo version: $(GOTOOLCHAIN=local tinygo version)"

rm -rf build
mkdir -p build

# Build the Go side into a relocatable object file.
cd sample
GOMIPS=softfloat tinygo build \
  -scheduler=none \
  -gc=psp \
  -target ../psp.json \
  -o ../build/goexports.o \
  .
cd ../

# Resolve the external symbols that survived TinyGo dead-code elimination to
# the PSPSDK archives that define them. Kprintf is required by bridge/printf.c.
python3 tools/resolve_pspsdk_libs.py \
  --output build/pspsdk-libraries.cmake \
  --require Kprintf \
  --require fdprintf \
  --require pspDebugScreenKprintf \
  build/goexports.o

# Configure and build using the PSP CMake wrapper.
cd build
psp_kernel_mode="${PSP_KERNEL_MODE:-OFF}"
psp-cmake -DPSP_KERNEL_MODE="$psp_kernel_mode" ..
make
cd ..

echo "Build complete. The resulting EBOOT.PBP can be found in the build directory."
