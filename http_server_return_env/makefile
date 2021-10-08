REGISTRY=docker.io/louhwz/printenv
TAG=0.0.2

IMAGE=$(REGISTRY):$(TAG)

build:
	go build -o main
	docker build -t $(IMAGE) .
	rm -rf main
push: build
	docker push $(IMAGE)