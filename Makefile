tidy:
	go mod tidy
	go mod download
.PHONY: tidy

run-api: tidy
	go run .
.PHONY: run-api