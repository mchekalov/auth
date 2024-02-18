package main

import "fmt"

// This package was created to solve error when running comand
// ./bin/golangci-lint run ./... --config .golangci.pipeline.yaml
// ERRO [linters_context] typechecking error: pattern ./...: directory prefix . does
// not contain modules listed in go.work or their selected dependencies
// See more: https://github.com/golangci/golangci-lint/issues/2654
func main() {
	fmt.Printf("Hello, Tester")
}
