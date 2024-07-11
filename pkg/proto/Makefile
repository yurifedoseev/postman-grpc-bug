include ../../go.env.mk
PROTO_PATH := ../../proto

proto: godeps
	protoc \
	--go_out=. \
	--go-grpc_out=require_unimplemented_servers=false:. \
	--proto_path=$(PROTO_PATH) \
	$(shell find $(PROTO_PATH) -name *.proto)


godeps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest; \
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest