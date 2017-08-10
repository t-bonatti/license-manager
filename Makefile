GOFILES = $$(find . -name '*.go' -not -path './vendor/*')
GOPACKAGES = $$(go list ./...  | grep -v /vendor/)

default: build

workdir:
	mkdir -p workdir

build:
	@go build

test: test-all

test-all:
	@go test -v $(GOPACKAGES)