NAME := k8sevents

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${NAME}
docker-build:
	docker build --rm -t quay.io/arschles/${NAME}

docker-push:
	docker push quay.io/arschles/${NAME}
