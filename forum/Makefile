all: run

delete:
	@rm -rf main

run: delete
	@go run src/*.go

build: delete
	@go build -ldflags="-w -s" -o main src/*.go

hardbuild: build
	@strip -x main
	@upx -9 main