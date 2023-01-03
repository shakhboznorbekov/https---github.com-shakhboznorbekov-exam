

go:
	go run cmd/main.go

swag-init:
	swag init -g api/api.go -o api/docs

migration-up:
	migrate -path ./migrations/postgres/ -database 'postgres://admin:korzinka@localhost:5432/users?sslmode=disable' up

migration-down:
	migrate -path ./migrations/postgres/ -database 'postgres://admin:korzinka@localhost:5432/users?sslmode=disable' down
