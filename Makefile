build:
	go build -ldflags="-s -w" -o build/ende

build:
	go install -ldflags="-s -w"

.PHONY: build
