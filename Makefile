.PHONY: all windows-x86_64 windows-x86 linux-x86_64 android-arm64

name := eulogist
output_suffix := build/

windows-x86_64.exe:
	GOOS=windows GOARCH=amd64 go build -o $(output_suffix)$@

windows-x86.exe:
	GOOS=windows GOARCH=386 go build -o $(output_suffix)$@

linux-x86_64:
	GOOS=linux GOARCH=amd64 go build -o $(output_suffix)$@

android-arm64:
	GOOS=android GOARCH=arm64 go build -o $(output_suffix)$@

windows-x86_64: windows-x86_64.exe
windows-x86: windows-x86.exe

all: windows-x86_64 windows-x86 linux-x86_64 android-arm64
