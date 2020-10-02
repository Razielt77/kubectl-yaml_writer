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

packagedarwin:
	CGO_ENABLED=0 GOOS=darwin go build -o ./kubectl-yaml_writer *.go
	mkdir temp
	mv ./kubectl-yaml_writer ./temp/
	cp ./LICENSE ./temp/
	cd ./temp && tar -czf ../kubectl-yaml_writer_darwin.tar.gz ./
	rm -rf temp

packagelinux:
	CGO_ENABLED=0 GOOS=linux go build -o ./kubectl-yaml_writer *.go
	mkdir temp
	mv ./kubectl-yaml_writer ./temp/
	cp ./LICENSE ./temp/
	cd ./temp && tar -czf ../kubectl-yaml_writer_linux.tar.gz ./
	rm -rf temp

darwin:
	CGO_ENABLED=0 GOOS=darwin go build -o ./kubectl-yaml_writer *.go

linux:
	CGO_ENABLED=0 GOOS=linux go build -o ./kubectl-yaml_writer *.go

clean:
	rm kubectl-yaml_writer *.tar.gz
