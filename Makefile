VERSION=`git describe --tags`
BUILD=`date +%FT%T%z`
BINARY_NAME=telbot

# Default operation: test, lint, build container, run container
all: test lint run
	
# Run unit tests
test:
	vgo test -v

# Run gometalinter on the sources using the project configuration
lint:
	gometalinter.v2  ./...


# Compile an executable
build:
	vgo build -o $(BINARY_NAME)

# Cross-compile to linux, attempt to set version and build number while at it
build-linux:
	env GOOS=linux vgo build -o $(BINARY_NAME)


# Build and run the program in a container
run: build
	./$(BINARY_NAME)
