setup:
	echo "setup"

test:
	go test ./...

test-cover:
	go test -cover ./...

PROTO_FILES := $(shell find . -name '*.proto')

lint:
	buf lint

build: proto
	buf generate
