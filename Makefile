VERSION=v0.0.1
build-image:
	docker buildx build --platform=linux/amd64 -f ./Dockerfile . -t registry.cn-hangzhou.aliyuncs.com/liyaso/startpage:$(VERSION) --load