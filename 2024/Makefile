aoc2024:
	go build -o aoc2024

clean:
	rm aoc2024

format:
	gofumpt -l -w .

format-check:
	test -z $(shell gofumpt -l .)

lint:
	golangci-lint run

test:
	go test ./internal/...


.PHONY: clean format format-check lint
