test:
	go test ./...

testv:
	go test -v ./...

testc:
	go test -cover ./...

build:
	go build -o dist/str

install: build
	cp dist/str $$HOME/go/bin
