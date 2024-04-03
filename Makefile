.PHONY: build

build:
	set GOOS=linux
	set GOARCH=arm64
	set CGO_ENABLED=0
	go env -w GOFLAGS=-mod=mod
	go mod tidy
	GOARCH=arm64 GOOS=linux go build -tags lambda.norpc -o ./bin/comment-getall cmd/comment/getall/main.go
	zip -FS comment-getall.zip ./bin/comment-getall
	GOARCH=arm64 GOOS=linux go build -tags lambda.norpc -o ./bin/comment-store cmd/comment/store/main.go
	zip -FS comment-store.zip ./bin/comment-store
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
	@${GOPATH}/bin/golangci-lint run ./internal/... ./pkg/... ./cmd/...
	@echo "=> Running tests"
	@go test ./internal/... ./pkg/... -covermode=atomic -coverpkg=./... -count=1 -race -v