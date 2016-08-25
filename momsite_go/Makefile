all: run

run:
	@go run *.go

build: clear
	@go build -o main *.go
	@strip -x main
	@upx -9 main

clear:
	@rm -rf main