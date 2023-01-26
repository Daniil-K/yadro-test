.PHONY:

build:
	go build -o ./.bin/test cmd/main.go

run: build
	./.bin/test
