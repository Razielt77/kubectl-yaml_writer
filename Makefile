.PHONY: build
build:
	go build -o ./kyml *.go

test:
	go test -cover cmd/*.go

linux:
	CGO_ENABLED=0 GOOS=linux go build -o ./kyml *.go
