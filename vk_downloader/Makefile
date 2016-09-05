all: parser

parser:
	@go run parser.go

downloader:
	@go run downloader.go

build:
	@go build -ldflags="-w -s" -o parser parser.go
	@go build -ldflags="-w -s" -o downloader downloader.go
	@strip -x parser
	@strip -x downloader
	@upx -9 parser
	@upx -9 downloader