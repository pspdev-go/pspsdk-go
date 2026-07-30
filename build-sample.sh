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

# Configure and build using the PSP CMake wrapper.
cd build
psp-cmake ..
make
cd ..

echo "Build complete. The resulting EBOOT.PBP can be found in the build directory."
