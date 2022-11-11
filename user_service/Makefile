run:
	go run cmd/*.go server

gen-proto:
	protoc --proto_path=../third_party/googleapis/google/api --go_out=. ./api/*.proto