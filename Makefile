.PHONY: build
build:
	go build
	docker build -t ghcr.io/rebell81/ddnser:0.2 -t ghcr.io/rebell81/ddnser:latest --platform linux/arm64 --no-cache .
.PHONY: push
push:
	docker push ghcr.io/rebell81/ddnser --all-tags

.PHONY: login
login:
	docker logout ghcr.io
	docker login ghcr.io

.PHONY: build-push
build-push: build push