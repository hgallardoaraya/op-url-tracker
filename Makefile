build:
	@go build -o bin/windows/op-url-tracker.exe cmd/main.go

run: build
	@./bin/windows/op-url-tracker.exe
