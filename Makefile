default: build

setup:
	go get

workdir:
	mkdir -p workdir

build:
	@go build

test: test-all

test-all:
	ginkgo -r