.PHONY: Setup REST API

setup:
	go mod tidy

install: setup
	go install ./...

build:
	cd api && go build -o main main.go

deploy:
	cdk deploy

destroy:
	cdk destroy