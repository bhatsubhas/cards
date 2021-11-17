build:
	go build
clean:
	go clean
	touch coverage.txt && rm coverage.txt
coverage:
	go test -coverprofile=coverage.txt ./...
all: clean coverage build
