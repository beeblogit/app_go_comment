.PHONY: build

build:
	sam build

start:
	make build
	sam local start-api
