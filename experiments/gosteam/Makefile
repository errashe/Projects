all:
	@go run *.go

clean:
	@rm -rf main main.exe

define buildProject
	$(1) go build -ldflags "-s -w" -o main *.go
	strip -x main
	upx -9 main
endef

build: clean
	$(call buildProject)

buildl: clean
	$(call buildProject, GOOS=linux GOARCH=386)