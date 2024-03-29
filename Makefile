PROGRAM_NAME=go-react-example
PUBLISHED_DIR=published

build: check-esbuild clean
	esbuild src/app.jsx --outfile=$(PUBLISHED_DIR)/app.js --bundle --loader:.js=jsx
	cp -r ./static/* $(PUBLISHED_DIR)/
	go build -o $(PROGRAM_NAME)

install-esbuild: check-npm
	@echo "Installing esbuild..."
	npm install -g esbuild

check-npm:
	@command -v npm >/dev/null 2>&1 || { echo >&2 "npm is required to install esbuild but it wasn't found. Aborting."; exit 1; }

check-esbuild:
	@command -v esbuild >/dev/null 2>&1 || { echo >&2 "esbuild is required but it's not installed. Install with 'make install-esbuild'"; exit 1; }

clean:
	rm -rf $(PUBLISHED_DIR)/*
	rm -f $(PROGRAM_NAME)

test: build 
	go test

.PHONY: build install-esbuild check-npm check-esbuild clean test