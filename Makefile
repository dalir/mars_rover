# Versioning
repo_name = $(shell basename $(shell git config --get remote.origin.url) | cut -d. -f1)
git_sha = $(shell git rev-parse --short HEAD)


.PHONY: docker-build


docker-build:
	docker build \
	--force-rm \
	--cpu-quota=-1 \
	--cpuset-cpus=0 \
	--memory=8G \
	-t ${repo_name}:${git_sha} \
	.


