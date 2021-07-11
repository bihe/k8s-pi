#vars
REPO=bihe
IMAGE_GOLANG=${REPO}/simpleapi

ifndef VERSION
$(error VERSION is not set)
endif

.PHONY: help build docker-build docker-run clean all docker-build-arm64

help:
	@echo "Makefile arguments:"
	@echo ""
	@echo "Makefile commands:"
	@echo "build"
	@echo "docker-build"
	@echo "docker-build-arm64"
	@echo "docker-run"
	@echo "clean"
	@echo "all"

.DEFAULT_GOAL := all

build:
ifeq ($(OS),Windows_NT)
	go build -ldflags="-s -w -X main.Version${VERSION}-DEV" -o simpleapi.exe ./...
else
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X main.Version=${VERSION}-DEV" -o simpleapi ./...
endif

docker-build:
	@docker build --pull --build-arg buildtime_variable_version=${VERSION} -t ${IMAGE_GOLANG}:${VERSION} ./

docker-build-arm64:
	@docker build --pull -t 192.168.1.111:32000/${IMAGE_GOLANG}:${VERSION}-arm64 --build-arg buildtime_variable_version=${VERSION} -f Dockerfile.arm64 ./
	@docker push 192.168.1.111:32000/${IMAGE_GOLANG}:${VERSION}-arm64

docker-run: docker-build
	@docker run -it -p 127.0.0.1:3000:3000 ${IMAGE_GOLANG}:${VERSION}

clean:
ifeq ($(OS),Windows_NT)
	pwsh -noprofile -command 'Remove-Item -Force -Recurse ./simpleapi.exe'
	docker rmi -f ${IMAGE_GOLANG}:${VERSION}
else
	rm -f ./simpleapi
	docker rmi -f ${IMAGE_GOLANG}:${VERSION}
endif

all: build
