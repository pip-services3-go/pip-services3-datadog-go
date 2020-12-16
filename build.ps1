#!/usr/bin/env pwsh

Set-StrictMode -Version latest
$ErrorActionPreference = "Stop"

# Generate image and container names using the data in the "component.json" file
$component = Get-Content -Path "component.json" | ConvertFrom-Json
$buildImage="$($component.registry)/$($component.name):$($component.version)-$($component.build)-build"
$container=$component.name

# Remove build files
if (Test-Path "./dist") {
    Remove-Item -Recurse -Force -Path "./dist/*"
} else {
    New-Item -ItemType Directory -Force -Path "./dist"
}

# Build docker image
docker build -f docker/Dockerfile.build -t $buildImage .

# Create and copy compiled files, then destroy the container
docker create --name $container $buildImage
docker cp "$($container):/app/main" ./dist/main
docker rm $container

if (!(Test-Path "./dist")) {
    Write-Host "dist folder doesn't exist in root dir. Build failed. Watch logs above."
    exit 1
}
