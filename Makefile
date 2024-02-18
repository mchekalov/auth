lint:
	./bin/golangci-lint run ./... --config .golangci.pipeline.yaml

build:
	go build -o ./bin -v ./...