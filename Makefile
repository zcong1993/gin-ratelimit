install:
	@go get -u -v github.com/golang/dep/cmd/dep
	@dep ensure

test:
	@go test ./...

test.cov:
	@go test ./... -coverprofile=coverage.txt -covermode=atomic

.PHONY: install test test.cov
