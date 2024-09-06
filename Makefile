build:
	go build -o blog-aggregator-cli

run:
	./blog-aggregator-cli

brun:
	go build -o blog-aggregator-cli && ./blog-aggregator-cli

test:
	go test ./... -v
