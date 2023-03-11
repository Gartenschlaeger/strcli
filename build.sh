#!/bin/sh

echo "Building for macOS"
GOOS=darwin GOARCH=amd64 go build -o dist/darwin-amd64/str
zip -j dist/darwin-amd64.zip dist/darwin-amd64/**

echo "Building for Linux"
GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/str
zip -j dist/linux-amd64.zip dist/linux-amd64/**

echo "Building for Windows"
GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/str.exe
zip -j dist/windows-amd64.zip dist/windows-amd64/**
