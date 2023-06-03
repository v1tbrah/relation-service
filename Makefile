up:
	docker-compose up -d

down:
	docker-compose down

protogen:
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
	cd rpbapi && \
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative relation-service.proto

mockgen:
	go install github.com/vektra/mockery/v2@v2.24.0
	go generate ./...
