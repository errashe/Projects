all:
	@go run *.go

build: clean
	@go build -ldflags="-w -s" -o main *.go
	@strip -x main
	@upx -9 main

clean:
	@rm -rf main