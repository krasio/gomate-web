.DEFAULT_GOAL := test

APP?=gomate-web

clean:
	rm -f ${APP}

build: clean
	go build -o ${APP}

run: build
	./${APP}

test:
	go test -v -race ./...
