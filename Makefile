.PHONY: build
build:
	#docker build -t ghcr.io/rebell81/ddnser:0.2 -t ghcr.io/rebell81/ddnser:latest --platform linux/arm64 --no-cache .
	docker buildx build  --platform linux/arm64/v8,linux/amd64 --push --tag ghcr.io/rebell81/ddnser:latest --tag ghcr.io/rebell81/ddnser:0.2 .
.PHONY: login
login:
	docker logout ghcr.io
	docker login ghcr.io

#docker buildx create --use --platform=linux/arm64,linux/amd64 --name multi-platform-builder
#docker buildx inspect --bootstrap