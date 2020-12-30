.PHONY: build clean deploy invoke

build:
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/main.go main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

invoke: clean build
	sls invoke local -f main --path lib/data.json