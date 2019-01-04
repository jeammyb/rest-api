default: build

IMAGE_FULL_NAME = gorestapi:latest

build:
	docker build -t ${IMAGE_FULL_NAME} .
run:
	docker run -d --rm --name gorestapi -p 8080:8080 ${IMAGE_FULL_NAME}
clean:
	docker stop gorestapi || true
	docker rm gorestapi || true
	docker rmi gorestapi:latest
