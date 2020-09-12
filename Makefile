.PHONY: build
build:
	go build -o ./kubectl-yaml_writer *.go

test:
	go test -cover cmd/*.go

fmt:
	gofmt -w ./cmd/*.go

linux:
	CGO_ENABLED=0 GOOS=linux go build -o ./kubectl-yaml_writer *.go
