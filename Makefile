set-windows-os:
	@GOOS=windows

build-windows: set-windows-os
	@go build -o bin/windows/op-url-tracker.exe cmd/main.go

run-win: build-windows
	@./bin/windows/op-url-tracker.exe

set-linux-os:
	@GOOS=linux

build-linux: set-linux-os
	@go build -o bin/linux/op-url-tracker cmd/main.go

run-linux: build-windows
	@./bin/linux/op-url-tracker
