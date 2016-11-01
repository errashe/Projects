all: debug

debug:
	@go-bindata -debug -o src/bin.go data
	@go run src/*.go

release:
	@go-bindata -o src/bin.go data
	@go build -o main src/*.go
