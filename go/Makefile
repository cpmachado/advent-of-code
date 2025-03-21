aoc:
	go build ./cmd/aoc-go

clean:
	rm aoc-go

format:
	gofumpt -w .

format-check:
	test -z $(shell gofumpt -l .)

lint:
	golangci-lint run

test:
	go test ./y2024/...


.PHONY: clean format format-check lint aoc-go
