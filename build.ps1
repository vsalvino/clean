Remove-Item -Recurse -Force -ErrorAction Ignore -Path dist
New-Item -Type Directory -Force -Name dist | Out-Null

$origOs = $env:GOOS
$origArch = $env:GOARCH

# 64 bit for all
$env:GOARCH = "amd64"

$env:GOOS = "darwin"
$out = Join-Path "dist" "clean-${env:GOOS}-${env:GOARCH}"
& go build -o $out

$env:GOOS = "freebsd"
$out = Join-Path "dist" "clean-${env:GOOS}-${env:GOARCH}"
& go build -o $out

$env:GOOS = "linux"
$out = Join-Path "dist" "clean-${env:GOOS}-${env:GOARCH}"
& go build -o $out

$env:GOOS = "netbsd"
$out = Join-Path "dist" "clean-${env:GOOS}-${env:GOARCH}"
& go build -o $out

$env:GOOS = "openbsd"
$out = Join-Path "dist" "clean-${env:GOOS}-${env:GOARCH}"
& go build -o $out

$env:GOOS = "windows"
$out = Join-Path "dist" "clean-${env:GOOS}-${env:GOARCH}.exe"
& go build -o $out

# Reset
$env:GOOS = $origOs
$env:GOARCH = $origArch
