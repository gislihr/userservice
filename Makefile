generate:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
	proto/user.proto

clean:
	rm -f proto/*.go

dev:
	PORT=8080 \
	PG_HOST=localhost \
	PG_USER=user \
	PG_PASSWORD=pass \
	PG_DATABASE=userservice \
	PG_SSL_MODE=disable \
	go run github.com/cosmtrek/air