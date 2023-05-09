.PHONV: swag fmt genModel

swag:
	swag1.8 init --parseDependency -g ./router/server.go -o ./internal/swagger

fmt:
	swag1.8 fmt

genModel:
	go run cmd/gen/gen.go
