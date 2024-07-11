include go.env.mk
PROTO_PATH := ../../proto

run: build
	./cmd/backend

build: proto
	go build -C cmd -o backend

proto: godeps clean
	protoc \
	--go_out=. \
	--go-grpc_out=require_unimplemented_servers=false:. \
	--proto_path=$(PROTO_PATH) \
	$(shell find $(PROTO_PATH) -name *.proto)


godeps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest; \
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

clean:
	rm -rf pkg/proto/api