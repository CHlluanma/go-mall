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
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/order_page.proto --server_name frontend --module github.com/CHlluanma/go-mall-kitex/app/frontend -I ../../idl


.PHONY: gen-user
gen-user: gen-user-client gen-user-server

.PHONY: gen-user-client
gen-user-client:
	@cd rpc_gen && cwgo client --type RPC --I ../idl --idl ../idl/rpc/user.proto --module github.com/CHlluanma/go-mall-kitex/rpc_gen --server_name user

.PHONY: gen-user-server
gen-user-server:
	@cd app/user && cwgo server --type RPC --I ../../idl --idl ../../idl/rpc/user.proto --module github.com/CHlluanma/go-mall-kitex/app/user --server_name user --pass "-use github.com/CHlluanma/go-mall-kitex/rpc_gen/kitex_gen"


.PHONY: gen-product
gen-checkout: gen-product-client gen-product-server

.PHONY: gen-product-client
gen-product-client:
	@cd rpc_gen && cwgo client --type RPC --I ../idl --idl ../idl/rpc/product.proto --module github.com/CHlluanma/go-mall-kitex/rpc_gen --server_name product

.PHONY: gen-product-server
gen-product-server:
	@cd app/product && cwgo server --type RPC --I ../../idl --idl ../../idl/rpc/product.proto --module github.com/CHlluanma/go-mall-kitex/app/product --server_name product --pass "-use github.com/CHlluanma/go-mall-kitex/rpc_gen/kitex_gen"


.PHONY: gen-cart
gen-checkout: gen-cart-client gen-cart-server

.PHONY: gen-cart-client
gen-cart-client:
	@cd rpc_gen && cwgo client --type RPC --I ../idl --idl ../idl/rpc/cart.proto --module github.com/CHlluanma/go-mall-kitex/rpc_gen --server_name cart

.PHONY: gen-cart-server
gen-cart-server:
	@cd app/cart && cwgo server --type RPC --I ../../idl --idl ../../idl/rpc/cart.proto --module github.com/CHlluanma/go-mall-kitex/app/cart --server_name cart --pass "-use github.com/CHlluanma/go-mall-kitex/rpc_gen/kitex_gen"


.PHONY: gen-payment
gen-payment: gen-payment-client gen-payment-server

.PHONY: gen-payment-client
gen-payment-client:
	@cd rpc_gen && cwgo client --type RPC --I ../idl --idl ../idl/rpc/payment.proto --module github.com/CHlluanma/go-mall-kitex/rpc_gen --server_name payment

.PHONY: gen-payment-server
gen-payment-server:
	@cd app/payment && cwgo server --type RPC --I ../../idl --idl ../../idl/rpc/payment.proto --module github.com/CHlluanma/go-mall-kitex/app/payment --server_name payment --pass "-use github.com/CHlluanma/go-mall-kitex/rpc_gen/kitex_gen"


.PHONY: gen-checkout
gen-checkout: gen-checkout-client gen-checkout-server

.PHONY: gen-checkout-client
gen-checkout-client:
	@cd rpc_gen && cwgo client --type RPC --I ../idl --idl ../idl/rpc/checkout.proto --module github.com/CHlluanma/go-mall-kitex/rpc_gen --server_name checkout

.PHONY: gen-checkout-server
gen-checkout-server:
	@cd app/checkout && cwgo server --type RPC --I ../../idl --idl ../../idl/rpc/checkout.proto --module github.com/CHlluanma/go-mall-kitex/app/checkout --server_name checkout --pass "-use github.com/CHlluanma/go-mall-kitex/rpc_gen/kitex_gen"


.PHONY: gen-order
gen-order: gen-order-client gen-order-server

.PHONY: gen-order-client
gen-order-client:
	@cd rpc_gen && cwgo client --type RPC --I ../idl --idl ../idl/rpc/order.proto --module github.com/CHlluanma/go-mall-kitex/rpc_gen --server_name order

.PHONY: gen-order-server
gen-order-server:
	@cd app/order && cwgo server --type RPC --I ../../idl --idl ../../idl/rpc/order.proto --module github.com/CHlluanma/go-mall-kitex/app/order --server_name order --pass "-use github.com/CHlluanma/go-mall-kitex/rpc_gen/kitex_gen"