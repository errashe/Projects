all:
	@go run *.go

clean:
	@rm -rf main main.exe

build: clean
	@go build -ldflags="-w -s" -o main *.go
	@strip -x main
	@upx -9 main