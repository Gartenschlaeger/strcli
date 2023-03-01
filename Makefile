test:
	go test -v ./...

run-field:
	echo "Das ist ein Test" | go run cmd/app/main.go field
