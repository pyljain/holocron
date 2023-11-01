run: build
	./holocron proxy

build:
	go build

generate-proto-proxy:
	protoc --go_out=. --go_opt=paths=source_relative \
    	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
    	./pkg/proto/proxy.proto

generate-proto-lookup:
	protoc --go_out=. --go_opt=paths=source_relative \
    	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
    	./pkg/proto/lookup.proto

generate-proto: generate-proto-proxy generate-proto-lookup

run-nats:
	docker run -d \
		--name nats \
		-p 4222:4222 \
		-p 8222:8222 \
		nats --http_port 8222