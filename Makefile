
docker-deploy: docker-build docker-run

docker-build:
	docker build -t newsletter .

docker-run:
	docker run --rm newsletter
