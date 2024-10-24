run:
	@go run main.go

build-swagger:
	@go run github.com/swaggo/swag/cmd/swag@latest init -g main.go -o docs
