# SHELL := /bin/bash

.PHONY: tidy
tidy:
	@echo "Tidying up all modules..."
	@for dir in $(shell find . -name go.mod -type f -not -path './vendor/*' -print); do \
		echo "Tidying up ======>> $${dir}"; \
		cd $${dir%/*} && go mod tidy ;\
		echo "Tidying up done ======>> $${dir}"; \
		cd -> /dev/null; \
	done

.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --module github.com/CHlluanma/go-mall-kitex/demo/demo_proto --server_name demo_proto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server --module github.com/CHlluanma/go-mall-kitex/demo/demo_thrift --server_name demo_thrift --idl ../../idl/echo.thrift

.PHONY: demo-link-fix
demo-link-fix:
	@cd demo/demo_proto && golangci-lint run -E gofumpt --path-prefix=. --fix --timeout=5m

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/auth_page.proto --server_name frontend --module github.com/CHlluanma/go-mall-kitex/app/frontend -I ../../idl