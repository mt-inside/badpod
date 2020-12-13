.PHONY: lint run image image-push docker-run
.DEFAULT_GOAL := run


lint:
	golangci-lint run

run:
	go run -ldflags "$(shell build/ldflags.sh)" cmd/badpod.go

image:
	docker build -t mtinside/badpod:latest .
image-push: image
	docker push mtinside/badpod
docker-run: image
	docker run --rm --name badpod -p8080:8080 mtinside/badpod:latest
