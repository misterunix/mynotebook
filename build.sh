#!/bin/sh

GOOS=linux GOARCH=amd64 go build -o bin/mynotebook-linux-amd64
GOOS=windows GOARCH=amd64 go build -o bin/mynotebook-windows-amd64.exe
