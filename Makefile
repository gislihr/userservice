generate:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
    proto/user.proto

clean:
	rm -f proto/*.go

dev:
	go run github.com/cosmtrek/air