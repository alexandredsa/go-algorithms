.PHONY: test coverage

test:
	go test -coverprofile=coverage.out ./...

coverage:
	go tool cover -html=coverage.out -o coverage.html
