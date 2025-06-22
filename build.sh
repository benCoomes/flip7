#!/bin/bash

# Build script for Flip 7 Simulator
echo "Building Flip 7 Simulator..."

# Clean previous builds
rm -f flip7-simulator

# Build the application
go build -o flip7-simulator ./cmd/

if [ $? -eq 0 ]; then
    echo "✅ Build successful!"
    echo "Run with: ./flip7-simulator"
    echo "Or for help: ./flip7-simulator -help"
else
    echo "❌ Build failed!"
    exit 1
fi
