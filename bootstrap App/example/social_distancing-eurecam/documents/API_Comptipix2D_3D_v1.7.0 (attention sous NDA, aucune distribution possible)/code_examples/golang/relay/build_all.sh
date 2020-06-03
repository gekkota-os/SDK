#!/bin/bash

echo "Build relay for Windows/OSX/Linux (amd64 and i386) + linux (arm and arm64) → use linux/arm if you intend to run on Raspberry Pi"

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
GOOS=windows GOARCH=386 go build
mv relay.exe build/windows/i386/

echo "build windows/amd64"
GOOS=windows GOARCH=amd64 go build
mv relay.exe build/windows/amd64/


# ------------------
# Darwin/OSX build
# ------------------

mkdir -p build/osx
mkdir -p build/osx/i386
mkdir -p build/osx/amd64

echo "build osx/i386"
GOOS=darwin GOARCH=386 go build -o relay.app relay_serv.go
mv relay.app build/osx/i386/

echo "build osx/amd64"
GOOS=darwin GOARCH=amd64 go build -o relay.app relay_serv.go
mv relay.app build/osx/amd64/


# ------------------
# Linux build
# ------------------

mkdir -p build/linux
mkdir -p build/linux/i386
mkdir -p build/linux/amd64
mkdir -p build/linux/arm
mkdir -p build/linux/arm64

echo "build linux/i386"
GOARCH=386 go build
mv relay build/linux/i386/

echo "build linux/amd64"
GOARCH=amd64 go build
mv relay build/linux/amd64/

echo "build linux/arm"
GOARCH=arm go build
mv relay build/linux/arm/

echo "build linux/arm64"
GOARCH=arm64 go build
mv relay build/linux/arm64/


# ------------------
# Compress results
# ------------------

echo "  "
echo "--------- "
echo "COMPRESS :"
echo "--------- "

7z a relay_windows build/windows/ relay_web/ doc_relay.svg READ_ME.txt relay.json -mx=9
7z a relay_osx build/osx/ relay_web/ doc_relay.svg READ_ME.txt relay.json -mx=9
7z a relay_linux build/linux/ relay_web/ doc_relay.svg READ_ME.txt relay.json -mx=9
7z a relay_all build/linux/ build/osx/ build/windows/ relay_web/ doc_relay.svg READ_ME.txt relay.json -mx=9
