.PHONY: build test run serve clean

build:
	go build ./cmd/cli ./cmd/http

test:
	go test ./...

run:
	go run ./cmd/cli

serve:
	go run ./cmd/http

clean:
	rm -f resume.html
