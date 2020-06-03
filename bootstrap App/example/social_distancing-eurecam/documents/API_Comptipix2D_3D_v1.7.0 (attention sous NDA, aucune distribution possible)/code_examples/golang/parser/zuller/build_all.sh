#!/bin/bash

# build for multiple arch
# NOTE: check OS/Arch available : https://golang.org/doc/install/source#environment
# if you need more OS/Arch, check GCCgo (for exemple if you need Solaris/Sparc) -> for this arch see http://ggolang.blogspot.fr/2015/05/gccgo-gcc510-cross-compile-for-sparc.html   +   https://hub.docker.com/r/craigbarrau/golang-solaris-sparc/

echo "Build zuller for Windows/OSX/Linux (amd64 and i386) + linux (arm and arm64) → use linux/arm if you intend to run on Raspberry Pi"

echo "  "
echo "--------- "
echo "BUILD    :"
echo "--------- "

# ------------------
# Windows build
# ------------------

mkdir -p build/windows
mkdir -p build/windows/i386
mkdir -p build/windows/amd64

echo "build windows/i386"
GOOS=windows GOARCH=386 go build -o zuller.exe zuller.go
mv zuller.exe build/windows/i386/

echo "build windows/amd64"
GOOS=windows GOARCH=amd64 go build -o zuller.exe zuller.go
mv zuller.exe build/windows/amd64/


# ------------------
# Darwin/OSX build
# ------------------

mkdir -p build/osx
mkdir -p build/osx/i386
mkdir -p build/osx/amd64

echo "build osx/i386"
GOOS=darwin GOARCH=386 go build -o zuller.app zuller.go
mv zuller.app build/osx/i386/

echo "build osx/amd64"
GOOS=darwin GOARCH=amd64 go build -o zuller.app zuller.go
mv zuller.app build/osx/amd64/


# ------------------
# Linux build
# ------------------

mkdir -p build/linux
mkdir -p build/linux/i386
mkdir -p build/linux/amd64

echo "build linux/i386"
GOARCH=386 go build -o zuller zuller.go
mv zuller build/linux/i386/

echo "build linux/amd64"
GOARCH=amd64 go build -o zuller zuller.go
mv zuller build/linux/amd64/


# ------------------
# Linux arm build
# ------------------

mkdir -p build/linux/arm
mkdir -p build/linux/arm64

echo "build linux/arm"
GOARCH=arm go build -o zuller zuller.go
mv zuller build/linux/arm/

echo "build linux/arm64"
GOARCH=arm64 go build -o zuller zuller.go
mv zuller build/linux/arm64/


# ------------------
# Compress results
# ------------------

echo "  "
echo "--------- "
echo "COMPRESS :"
echo "--------- "

7z a build/zuller_windows build/windows/ zuller.json -mx=9
7z a build/zuller_osx build/osx/ zuller.json -mx=9
7z a build/zuller_linux build/linux/ zuller.json -mx=9
7z a build/zuller_all build/linux/ build/osx/ build/windows/ zuller.json dst_rules.csv -mx=9
