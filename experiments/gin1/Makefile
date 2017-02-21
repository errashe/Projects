FILES = main.go funcs.go

all:
	@go run $(FILES)

build:
	@go build -o main $(FILES)

build_linux:
	@GOOS=linux GOARCH=amd64 go build -o main $(FILES)