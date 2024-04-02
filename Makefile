.PHONY: build

build:
	sam build

start:
	if [ -e ".aws-sam" ];then rm -rf ".aws-sam" ; fi
	make build
	sam local start-api --env-vars env.json

deploy:
	make build
	sam package --template-file template.yaml --output-template-file output.yaml
	sam deploy --template-file output.yaml

test:
	@echo "=> Running linter"
	@${GOPATH}/bin/golangci-lint run ./internal/... ./pkg/...
	@echo "=> Running tests"
	@go test ./internal/... ./pkg/... -covermode=atomic -coverpkg=./... -count=1 -race -v