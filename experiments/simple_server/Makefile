all: server client

server: server.go
	go build server.go

client: client.go
	go build client.go

win:
	@GOOS=windows GOARCH=386 go build client.go
	@GOOS=windows GOARCH=386 go build server.go
	@echo "Compiled win!"

unix:
	@GOOS=linux GOARCH=386 go build -o clientu client.go
	@GOOS=linux GOARCH=386 go build -o serveru server.go
	@echo "Compiled unix!"

clean:
	@rm -rf client clientu client.exe
	@rm -rf server serveru server.exe