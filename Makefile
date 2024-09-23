api:
	goctl api go -api ./apps/app/api.api -dir ./apps/app/

order:
	goctl rpc protoc ./apps/order/rpc/order.proto --go_out=./apps/order/rpc/ --go-grpc_out=./apps/order/rpc --zrpc_out=./apps/order/rpc

product:
	goctl rpc protoc ./apps/product/rpc/product.proto --go_out=./apps/product/rpc/ --go-grpc_out=./apps/product/rpc --zrpc_out=./apps/product/rpc

etcd:
	etcd --listen-client-urls 'http://0.0.0.0:2379' --advertise-client-urls 'http://0.0.0.0:2379'

rpc:
	go run ./apps/order/rpc/order.go
	go run ./apps/product/rpc/product.go