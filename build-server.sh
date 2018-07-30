#!/bin/bash
echo "Checking bin..."
if [ ! -d "bin" ]; then
    mkdir bin
fi
echo "Checking config..."
if [ ! -d "config" ]; then
    mkdir config
fi
dep ensure
echo "Building config script..."
go build -o bin/choc-config cmd/choc-config.go
echo "Building startup script..."
go build -o bin/choc-up cmd/choc-start-server.go
echo "Done."
