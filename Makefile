.PHONY: docs
init:
	godep go install
test:
	# This runs all of the tests, on both Python 2 and Python 3.
	go test
docs:
	cd docs && make
build:
	go build
