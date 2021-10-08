REGISTRY=docker.io/louhwz/printenv
TAG=0.0.3

IMAGE=$(REGISTRY):$(TAG)

build:
	GOOS=linux GOARCH=amd64 go build -o main
	docker build -t $(IMAGE) .
	rm -rf main
push: build
	docker push $(IMAGE)