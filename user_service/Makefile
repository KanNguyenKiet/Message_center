run:
	go run cmd/*.go server

gen-proto:
	protoc -I ../third_party/googleapis -I ./api --go_out=. ./api/*.proto