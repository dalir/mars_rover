# Versioning
REPO_NAME := $(shell basename $(shell git config --get remote.origin.url) | cut -d. -f1)
GIT_SHA := $(shell git rev-parse --short HEAD)
DOCKER_IMAGE := ${REPO_NAME}:${GIT_SHA}

# Resources
CPU_QUOTA := -1
CPU_SET := 0
MEMORY := 8G

.PHONY: build docker-build docker-tag docker-push clean

# Compile the Go project
build:
	@echo "Building Go project..."
	go build -o ${REPO_NAME}
	@echo "Build completed for ${REPO_NAME}"

# Test the go project
test: build
	@echo "Testing Go project..."
	go test -v
	@echo "Testing completed for ${REPO_NAME}"

# Build the Docker image
docker-build:
	@echo "Building Docker image for ${REPO_NAME} with tag ${GIT_SHA}"
	docker build \
		--force-rm \
		--cpu-quota=${CPU_QUOTA} \
		--cpuset-cpus=${CPU_SET} \
		--memory=${MEMORY} \
		-t ${DOCKER_IMAGE} \
		.
	@make docker-tag

# Tag the Docker image with 'latest'
docker-tag:
	@echo "Tagging image ${DOCKER_IMAGE} as latest"
	docker tag ${DOCKER_IMAGE} ${REPO_NAME}:latest

# Push the image to a Docker registry
docker-push:
	@echo "Pushing ${DOCKER_IMAGE} and ${REPO_NAME}:latest to registry"
	docker push ${DOCKER_IMAGE}
	docker push ${REPO_NAME}:latest

# Clean up the built binary and Docker images
clean:
	@echo "Cleaning up binary and Docker images..."
	rm -f ${REPO_NAME}
	-docker rmi -f ${DOCKER_IMAGE} || true
	-docker rmi -f ${REPO_NAME}:latest || true
	@echo "Cleanup completed"
