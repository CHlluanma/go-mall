.PHONY: build-server
build-server:
	@sh build.sh
	@echo "build-server done"

run-server: build-server
	@sh output/bootstrap.sh

.PHONY: run-client
run-client:
	@go run cmd/client/*

.PHONY: run-dbop
run-dbop:
	@go run cmd/dbop/*