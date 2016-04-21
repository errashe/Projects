all: server client

server: server.go Makefile
	go build server.go

client: client.go Makefile
	go build client.go

win: client.go server.go Makefile
	@GOOS=windows GOARCH=386 go build client.go
	@GOOS=windows GOARCH=386 go build server.go
	@echo "Compiled!"

cpp: echo_client.cpp Makefile
	g++ echo_client.cpp -o echo