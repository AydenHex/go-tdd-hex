PROTO_DIR := src/todos/infrastructure/adapter/grpc/proto
GRPC_TARGET_DIR := src/todos/infrastructure/adapter/grpc
REST_GW_TARGET_DIR := src/todos/infrastructure/adapter/rest
REST_GW_OUT_FILE := todo.pb.gw.go

generate_proto:
	@protoc \
		-I $(GRPC_TARGET_DIR) \
		--go-grpc_out=$(GRPC_TARGET_DIR) \
		--grpc-gateway_out=logtostderr=true,import_path=todorest:$(REST_GW_TARGET_DIR) \
		--swagger_out=logtostderr=true:$(REST_GW_TARGET_DIR) \
		$(PROTO_DIR)/todo.proto
lint:
	golangci-lint run --build-tags test ./...

# https://github.com/golangci/golangci-lint
install-golangci-lint:
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(shell go env GOPATH)/bin v1.24.0


# https://github.com/psampaz/go-mod-outdated
outdated-list:
	go list -u -m -json all | go-mod-outdated -update -direct