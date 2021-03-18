#!/usr/bin/env pwsh

Remove-Item -Recurse -Force -ErrorAction Ignore -Path dist
New-Item -Type Directory -Force -Name dist | Out-Null

$origOs = $env:GOOS
$origArch = $env:GOARCH

# Get all supported distributions.
$goDists = & go tool dist list -json | ConvertFrom-Json

foreach ($dist in $goDists) {
    # Skip android, ios, js builds.
    if (@("android", "ios", "js").Contains($dist.GOOS)) { continue }

    # Set output path.
    $out = Join-Path "dist" "clean-$($dist.GOOS)-$($dist.GOARCH)"
    if ($dist.GOOS -eq "windows") { $out += ".exe" }

    # Set env variables for go build and compile.
    $env:GOOS = $dist.GOOS
    $env:GOARCH = $dist.GOARCH
    Write-Output "Building $out"
    & go build -o $out
}

# Reset.
$env:GOOS = $origOs
$env:GOARCH = $origArch
