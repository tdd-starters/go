.PHONY: all
all: test vet build

.PHONY: all
test:
	go test ./...

.PHONY: build
build:
	go build ./...
	go list ./... | grep /cmd/ | while read p; do \
		go build $$p ; \
	done

.PHONY: format
format:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: clean
clean:
	go clean ./...
