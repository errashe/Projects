all:
	# @rm -rf dev.db
	@go-bindata --debug data
	@go run *.go

build:
	@go-bindata data
	@go build -o main *.go