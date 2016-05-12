all:
	@go run *.go

build: clean
	@go build -o main -ldflags="-s -w" *.go
	@strip -x main
	@upx -9 main

cross: clean
	@GOOS=windows GOARCH=386 go build -o main.exe -ldflags="-s -w" *.go
	@upx -9 main.exe

clean:
	@rm -rf main main.exe