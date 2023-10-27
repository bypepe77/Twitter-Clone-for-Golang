unit-test:
	go test -v -coverprofile=coverage.out ./...

run-linter:
	golangci-lint run

run:
	go run cmd/main.go

gen-mocks: 
	mockery