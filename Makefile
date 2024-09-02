BINARY_NAME=tasks
MAIN_FILE=main.go

.PHONY: build
build:
	go build -o $(BINARY_NAME) $(MAIN_FILE)

.PHONY: install
install:
	go install

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -f $(BINARY_NAME)

.PHONY: run
run:
	go run $(MAIN_FILE)
