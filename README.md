# User service example

Written using go.

# Calling endpoints locally

grpcurl -d '{"user_name": "gisli", "email": "gisli@yahoo.com", "password": "1234"}' -plaintext localhost:8080 UserService.AddUser
