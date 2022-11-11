gen-protoc:
	protoc -I /home/kan/Documents/GolangServer -I vendor -I third_party/googleapis -I user_service/api/ --go_out=. --go-grpc_out=. --grpc-gateway_out=. ./user_service/api/*.proto
