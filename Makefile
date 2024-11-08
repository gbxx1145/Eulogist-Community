.PHONY: all windows-x86_64 windows-x86 linux-x86_64 linux-x86 android-aarch64

name := eulogist
output_suffix := build/$(name)

$(output_suffix)-windows-x86_64.exe:
	GOOS=windows GOARCH=amd64 go build -o $@

$(output_suffix)-windows-x86.exe:
	GOOS=windows GOARCH=386 go build -o $@

$(output_suffix)-linux-x86_64:
	GOOS=linux GOARCH=amd64 go build -o $@

$(output_suffix)-linux-x86:
	GOOS=linux GOARCH=386 go build -o $@

$(output_suffix)-android-aarch64:
	GOOS=android GOARCH=arm64 go build -o $@

$(output_suffix)-android-armreabi:
	GOOS=android GOARCH=arm go build -o $@

windows-x86_64: $(output_suffix)-windows-x86_64.exe
windows-x86: $(output_suffix)-windows-x86.exe
linux-x86_64: $(output_suffix)-linux-x86_64
linux-x86: $(output_suffix)-linux-x86
android-aarch64: $(output_suffix)-android-aarch64

all: windows-x86_64 windows-x86 linux-x86_64 linux-x86 android-aarch64
