#!/bin/sh

GOOS=linux GOARCH=amd64 go build -gcflags="all=-N -l" -o linux-amd64/mynotebook-debug-amd64
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o linux-amd64/mynotebook-amd64-min
GOOS=windows GOARCH=amd64 go build -gcflags="all=-N -l" -o windows-amd64/mynotebook-debug-amd64.exe
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o windows-amd64/mynotebook-amd64.exe


gomarkdoc -u -o ../../../docs/main.md main.go
gomarkdoc -u -o ../../../docs/common.md ../../../internal/common
gomarkdoc -u -o ../../../docs/pdf.md ../../../internal/pdf 

