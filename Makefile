all: deps build

build:
	go build
test: 
	go test
deps:
	go get -d -t
imports: 
	go get -v $(go list -f '{{range .Imports}}{{ . }} {{end}}')
