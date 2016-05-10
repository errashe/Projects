all: run

run:
	@go run *.go

build:
	@go build -o main -ldflags "-s -w" *.go
	@strip -x main

cross:
	@GOOS=windows GOARCH=386 go build -o main.exe -ldflags "-s -w" *.go

test:
	@go run *.go "музеи кургана" "памятные места кургана" "церкви кургана"

clean:
	@rm -rf main main.exe