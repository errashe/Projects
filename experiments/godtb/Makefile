clean:
	@rm -rf main main.exe

all:
	@go run *.go

build:
	@go build -ldflags="-s -w" -o main *.go
	@strip -x main
	@upx -9 main