.DEFAULT_GOAL := test

PROJECT?=github.com/krasio/gomate-web
APP?=gomate-web
PORT?=8080
REDIS_HOST_PORT?=9990
RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
GOOS?=linux
GOARCH?=amd64
CONTAINER_IMAGE?=docker.io/krasio/${APP}

clean:
	rm -f ${APP}

build: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
		-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o ${APP}

container: build
	docker build -t $(CONTAINER_IMAGE):$(RELEASE) .

push: container
	docker push $(CONTAINER_IMAGE):$(RELEASE)

run: container
	docker stop $(APP) || true && docker rm $(APP) || true
	docker stop gomate-redis || true && docker rm gomate-redis || true
	docker run --name gomate-redis -v ${PWD}/data:/data -p 127.0.0.1:${REDIS_HOST_PORT}:6379 -d redis redis-server --appendonly yes
	docker run --name ${APP} --link gomate-redis:gomate-redis -p ${PORT}:${PORT} --rm \
		-e "GOMATE_PORT=${PORT}" -e "GOMATE_REDIS_URL=redis://gomate-redis:6379/0"\
		$(APP):$(RELEASE)

stop:
	docker stop $(APP) || true
	docker stop gomate-redis || true

minikube: push
				for t in $(shell find ./kubernetes/gomate-web -type f -name "*.yaml"); do \
				cat $$t | \
								sed -E "s/\{\{\.Release\}\}/$(RELEASE)/g" | \
								sed -E "s/\{\{\.ServiceName\}\}/$(APP)/g" | \
								sed -E "s/\{\{\.Port\}\}/$(PORT)/g"; \
				echo ---; \
		done > kubernetes/gomate-web/tmp.yaml
				kubectl apply -f kubernetes/gomate-web/tmp.yaml

test:
	go test -v -race ./...
