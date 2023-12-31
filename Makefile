.PHONY: build clean generate

build: generate
	@go build -o goshower main.go

clean:
	@rm goshower
	@docker rmi goshower

generate:
	@go generate ./global
	@swag init

docker:
	@docker build . -t goshower:latest
