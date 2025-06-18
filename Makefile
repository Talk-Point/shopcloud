setup:
	echo "setup"

sync-proto:
	./scripts/sync-proto.sh

test:
	go test ./...

test-cover:
	go test -cover ./...

PROTO_FILES := $(shell find . -name '*.proto')

lint:
	docker run --volume "$(CURDIR):/workspace" --workdir /workspace bufbuild/buf lint

build: proto
	docker run --volume "$(CURDIR):/workspace" --workdir /workspace bufbuild/buf generate
