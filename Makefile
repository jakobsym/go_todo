build:
	@go build -o ../../bin/go_todo

run: build
	@./bin/go_todo
test:
	@go test -v ./...