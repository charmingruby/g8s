DOCKER_IMAGE_NAME := g8s

################
# App          #
################
PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/g8s ./cmd/api/main.go

################
# Docker       #
################
.PHONY: docker-build
docker-build:
	docker build -t ${DOCKER_IMAGE_NAME} .

.PHONY: docker-tag
docker-tag:
	docker tag ${DOCKER_IMAGE_NAME} charmingruby/${DOCKER_IMAGE_NAME}:$(VERSION)

.PHONY: docker-push
docker-push:
	docker push charmingruby/${DOCKER_IMAGE_NAME}:$(VERSION)

.PHONY: docker-run
docker-run:
	docker run -p $(PORT):$(PORT) ${DOCKER_IMAGE_NAME}
