.PHONY: all delete proxy tester

define DELETE
	@rm -rf $1
endef

all: delete proxy tester

delete:
	$(call DELETE,proxy tester)

proxy:
	@go build -ldflags="-w -s" -o proxy proxy.go
	@strip -x proxy
	@upx -9 proxy

tester:
	@go build -ldflags="-w -s" -o tester tester.go
	@strip -x tester
	@upx -9 tester