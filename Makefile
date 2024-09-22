
docker-deploy: docker-build docker-run
run: tw run-local

docker-build:
	docker build -t newsletter .

docker-run:
	docker run --rm newsletter

run-local:
	cd ./cmd; go run .;

tw:
	cd ./internal/views; \
	templ generate

push:
	git add .;
	git commit -m "quick commit";
	git push;
