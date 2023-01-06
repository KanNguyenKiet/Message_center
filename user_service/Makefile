run:
	go run cmd/*.go server

gen-proto:
	protoc -I ../third_party/googleapis -I ./api --go_out=. --go-grpc_out=. \
	--grpc-gateway_out ./api \
	--grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
	./api/*.proto