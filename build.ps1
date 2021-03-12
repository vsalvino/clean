New-Item -Type Directory -Force -Name dist

$origOs = $env:GOOS
$origArch = $env:GOARCH

# 64 bit for all
$env:GOARCH = "amd64"

$env:GOOS = "darwin"
$dir = Join-Path "dist" ${env:GOOS}-${env:GOARCH}
New-Item -Type Directory -Force -Name $dir
& go build -o $dir

$env:GOOS = "freebsd"
$dir = Join-Path "dist" ${env:GOOS}-${env:GOARCH}
New-Item -Type Directory -Force -Name $dir
& go build -o $dir

$env:GOOS = "linux"
$dir = Join-Path "dist" ${env:GOOS}-${env:GOARCH}
New-Item -Type Directory -Force -Name $dir
& go build -o $dir

$env:GOOS = "openbsd"
$dir = Join-Path "dist" ${env:GOOS}-${env:GOARCH}
New-Item -Type Directory -Force -Name $dir
& go build -o $dir

$env:GOOS = "netbsd"
$dir = Join-Path "dist" ${env:GOOS}-${env:GOARCH}
New-Item -Type Directory -Force -Name $dir
& go build -o $dir

$env:GOOS = "windows"
$dir = Join-Path "dist" ${env:GOOS}-${env:GOARCH}
New-Item -Type Directory -Force -Name $dir
& go build -o $dir

# Reset
$env:GOOS = $origOs
$env:GOARCH = $origArch
