aoc2024:
	go build ./cmd/aoc2024

clean:
	rm aoc2024

format:
	gofumpt -w .

format-check:
	test -z $(shell gofumpt -l .)

lint:
	golangci-lint run

test:
	go test ./...


.PHONY: clean format format-check lint aoc2024
