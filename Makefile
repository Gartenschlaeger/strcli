test:
	go test -v ./...

build:
	go build -o dist/str

install: build
	cp dist/str $$HOME/go/bin
