PROJECT_NAME=wallet
BUILD_VERSION=0.0.1

DOCKER_IMAGE=$(PROJECT_NAME):$(BUILD_VERSION)
GO_BUILD_ENV=CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on

compose_dev: docker
	docker compose up --build --force-recreate -d

build:
	$(GO_BUILD_ENV) go build -v -o $(PROJECT_NAME)-$(BUILD_VERSION).bin main.go

docker_prebuild: build
	mv $(PROJECT_NAME)-$(BUILD_VERSION).bin deploy/$(PROJECT_NAME).bin; \

docker_build:
	cd deploy; \
	docker build --rm -t $(DOCKER_IMAGE) .;

docker_postbuild:
	cd deploy; \
	rm -rf $(PROJECT_NAME).bin 2> /dev/null;\

docker: docker_prebuild docker_build docker_postbuild