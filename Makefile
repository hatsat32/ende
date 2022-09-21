build:
	go build -ldflags="-s -w" -o build/ende

install:
	go install -ldflags="-s -w"

.PHONY: build
