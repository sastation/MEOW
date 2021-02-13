#!/bin/bash

# don't forget to modify config.go for version info"

echo "Refresh chinaip data..."
go run china.ip/chinaip_gen.go

echo "Building win-amd64..."
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o meow-windows-amd64
echo "Building linux-amd64..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o meow-linux-amd64
echo "Building linux-arm..."
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o meow-linux-arm
