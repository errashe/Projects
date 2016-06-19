all:
	@go run *.go

clean:
	@rm -rf main main.exe

build:
	@go build -ldflags="-s -w" -o main *.go
	@strip -x main
	@upx -9 main