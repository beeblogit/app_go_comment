.PHONY: build

build:
	sam build

start:
	make build
	sam local start-api

deploy:
	make build
	sam deploy --guided