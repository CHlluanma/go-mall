.PHONY: build-server
build-server:
	@sh build.sh

run-server: build-server
	@sh output/bootstrap.sh

.PHONY: run-client
run-client:
	@go run cmd/client/*