all: server client
win: winserver winclient

server: server.go Makefile
	go build server.go

client: client.go Makefile
	go build client.go

winserver: server.go Makefile
	GOOS=windows GOARCH=386 go build server.go
winclient: client.go Makefile
	GOOS=windows GOARCH=386 go build client.go

clean:
	@rm -rf client client.exe server server.exe
	@echo "Deleted everyone executable"