#!/bin/bash

echo "Build parser for Windows/OSX/Linux (amd64 and i386) + linux (arm and arm64) → use linux/arm if you intend to run on Raspberry Pi"

echo "  "
echo "--------- "
echo "BUILD    :"
echo "--------- "

cd src

# ------------------
# Windows build
# ------------------

mkdir -p ../build/windows
mkdir -p ../build/windows/i386
mkdir -p ../build/windows/amd64

echo "build windows/i386"
GOOS=windows GOARCH=386 go build -tags netgo -o ../parser.exe
mv ../parser.exe ../build/windows/i386/

echo "build windows/amd64"
GOOS=windows GOARCH=amd64 go build -tags netgo -o ../parser.exe
mv ../parser.exe ../build/windows/amd64/


# ------------------
# Darwin/OSX build
# ------------------

mkdir -p ../build/osx
mkdir -p ../build/osx/i386
mkdir -p ../build/osx/amd64

echo "build osx/i386"
GOOS=darwin GOARCH=386 go build -tags netgo -o ../parser.app
mv ../parser.app ../build/osx/i386/

echo "build osx/amd64"
GOOS=darwin GOARCH=amd64 go build -tags netgo -o ../parser.app
mv ../parser.app ../build/osx/amd64/


# ------------------
# Linux build
# ------------------

mkdir -p ../build/linux
mkdir -p ../build/linux/i386
mkdir -p ../build/linux/amd64
mkdir -p ../build/linux/arm
mkdir -p ../build/linux/arm64

echo "build linux/i386"
GOARCH=386 go build -tags netgo -o ../parser
mv ../parser ../build/linux/i386/

echo "build linux/amd64"
GOARCH=amd64 go build -tags netgo -o ../parser
mv ../parser ../build/linux/amd64/

echo "build linux/arm"
GOARCH=arm go build -tags netgo -o ../parser
mv ../parser ../build/linux/arm/

echo "build linux/arm64"
GOARCH=arm64 go build -tags netgo -o ../parser
mv ../parser ../build/linux/arm64/


# ------------------
# Compress results
# ------------------

echo "  "
echo "--------- "
echo "COMPRESS :"
echo "--------- "

cd ..

7z a build/parser_windows build/windows/ doc-files/ parser.json -mx=9
7z a build/parser_osx build/osx/ doc-files/ parser.json -mx=9
7z a build/parser_linux build/linux/ doc-files/ parser.json -mx=9
7z a build/parser_all build/linux/ build/osx/ build/windows/ doc-files/ parser.json -mx=9
