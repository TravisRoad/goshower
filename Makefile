.PHONY: build clean

build:
	@go build -o goshower main.go

clean:
	@rm goshower
	@docker rmi goshower

docker:
	@docker build . -t goshower:latest
