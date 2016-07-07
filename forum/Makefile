all: run

delete:
	@rm -rf main main.exe

run: delete
	@go run *.go

build: delete
	@go build -ldflags="-w -s" -o main *.go

hardbuild: build
	@strip -x main
	@upx -9 main