#!/bin/bash
# Windows
#64bit
GOOS=windows GOARCH=amd64 go build -o LabNote-windows-amd64.exe

# Linux
# 64-bit
GOOS=linux GOARCH=amd64 go build -o LabNote-linux-amd64

# MAC os
# 64-bit
GOOS=darwin GOARCH=amd64 go build -o LabNote-amd64-darwin
