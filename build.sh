#!/bin/bash

# Paths
goOutput=release/bin
frontendPath=./frontend

# Script
(cd $frontendPath && npm run build)
templ generate
go build -o $goOutput
