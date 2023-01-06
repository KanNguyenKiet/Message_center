GO_TOOLS = github.com/kyleconroy/sqlc/cmd/sqlc@v1.16.0 \
           google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 \
           google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 \
           github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.14.7

gen-protoc:
	protoc -I /home/kan/Documents/GolangServer -I vendor -I third_party/googleapis -I user_service/api/ --go_out=. --go-grpc_out=. --grpc-gateway_out=. ./user_service/api/*.proto

install_tools:
	@echo $(GO_TOOLS) | xargs -r -n1 go install