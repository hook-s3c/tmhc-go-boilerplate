.PHONY: docs
init:
	godep go install
test:
	go test
docs:
	cd docs && make
build:
	go build
