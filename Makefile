.PHONY: build
build:
	go build -o ./kubectl-yaml_writer *.go

test:
	go test -cover cmd/*.go

fmt:
	gofmt -w ./cmd/*.go

package:
	CGO_ENABLED=0 GOOS=darwin go build -o ./kubectl-yaml_writer *.go
	tar -czf kubectl-yaml_writer_darwin.tar.gz kubectl-yaml_writer

darwin:
	CGO_ENABLED=0 GOOS=darwin go build -o ./kubectl-yaml_writer *.go

linux:
	CGO_ENABLED=0 GOOS=linux go build -o ./kubectl-yaml_writer *.go
