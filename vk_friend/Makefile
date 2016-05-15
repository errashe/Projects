all: build

parse:
	@go run *.go

search:
	@go run *.go search

build: clean
	@go build -ldflags="-w -s" -o main *.go
	@strip -x main
	@upx -9 main

linux: clean
	@GOOS=linux GOARCH=386 go build -ldflags="-w -s" -o main *.go
	@upx -9 main

clean:
	@rm -rf main main.exe