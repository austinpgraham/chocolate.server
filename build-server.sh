#!/bin/bash
echo "Checking bin..."
if [ ! -d "$DIRECTORY" ]; then
    mkdir bin
fi
echo "Building config script..."
go build -o bin/choc-config cmd/choc-config.go
echo "Building startup script..."
go build -o bin/choc-up cmd/choc-start-server.go
echo "Done."
