.PHONV: swag fmt

swag:
	swag1.8 init --parseDependency -g ./router/server.go -o ./internal/swagger

fmt:
	swag1.8 fmt
