.PHONY: format format-check vet test build check

format:
	gofmt -w $$(find . -name '*.go' -not -path './vendor/*')

format-check:
	@test -z "$$(gofmt -l $$(find . -name '*.go' -not -path './vendor/*'))"

vet:
	go vet ./...

test:
	go test -race ./...

build:
	go build ./cmd/app

check: format-check vet test build
