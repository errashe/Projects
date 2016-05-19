all: run

clean:
	@rm -rf main main.exe

run:
	@rm -rf users.db
	@go run *.go

build: clean
	@go build -ldflags="-s -w" -o main *.go
	@strip -x main
	@upx -9 main