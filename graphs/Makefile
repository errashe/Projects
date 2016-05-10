all:
	@go run *.go

clean:
	@rm -rf main main.exe
	
build: clean
	@go build -o main -ldflags "-s -w" *.go
	@strip -x main

cross: clean build
	@GOOS=windows GOARCH=386 go build -o main.exe -ldflags "-s -w" *.go