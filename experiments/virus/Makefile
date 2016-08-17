all:
	@go build -ldflags="-s -w" -o main *.go
	@strip -x main
	@upx -9 main
	ls -la