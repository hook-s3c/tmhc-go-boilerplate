.PHONY: docs
init:
	godep go install
test:
	cd src && goconvey -host="0.0.0.0"
docs:
	cd docs && make
build:
	cd src && go build -o dist/main
run:
	cd src && go run main.go
