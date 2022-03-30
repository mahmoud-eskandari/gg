build:
	go build -v ./...

test:
	go test -race -v ./...
	
watch-test:
	reflex -t 50ms -s -- sh -c 'gotest -race -v ./...'

bench:
	go test -benchmem -count 3 -bench ./...

watch-benchmark:
	reflex -t 50ms -s -- sh -c 'go test -benchmem -count 3 -bench ./...'

coverage:
	go test -v -coverprofile cover.out .
	go tool cover -html=cover.out -o cover.html

tools:
	go install github.com/cespare/reflex@latest
	go install github.com/rakyll/gotest@latest
	go install github.com/psampaz/go-mod-outdated@latest
	go install github.com/jondot/goweight@latest
	go get -t -u golang.org/x/tools/cmd/cover
	go mod tidy

lint:
	go fmt

audit: tools
	go mod tidy

outdated: tools
	go mod tidy
	go list -u -m -json all | go-mod-outdated -update -direct
