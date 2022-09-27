VERSION=v0.0.1
build-image:
	docker buildx build --platform=linux/amd64 -f ./Dockerfile . -t liyaso/startpage:$(VERSION) --load