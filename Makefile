setup:
	echo "setup"

test:
	go test ./...

test-cover:
	go test -cover ./...

PROTO_FILES := $(shell find . -name '*.proto')

lint:
	docker run --volume "$(pwd):/workspace" --workdir /workspace bufbuild/buf lint

build: proto
	docker run --volume "$(pwd):/workspace" --workdir /workspace bufbuild/buf generate
