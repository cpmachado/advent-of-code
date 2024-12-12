aoc2024:
	go build -o aoc2024

clean:
	rm aoc2024

format:
	gofumpt -l -w .

format-check:
	exit $(shell gofumpt -l . | wc -l)

lint:
	golangci-lint run

test:
	go test ./...


.PHONY: clean format format-check lint
