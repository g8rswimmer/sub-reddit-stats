
gen_mocks:
	mockgen -source=internal/worker/runner.go -destination=internal/worker/mock_runner_test.go -package=worker
	mockgen -source=internal/manager/reddit.go -destination=internal/manager/mock_reddit_test.go -package=manager
	mockgen -source=internal/service/reddit.go -destination=internal/service/mock_reddit_test.go -package=service

init_db:
	mkdir db && touch db/sqlite-database.db

clean_db:
	rm -r db

migrate:
	go run cmd/migration/*.go -config=config.json

daemon:
	go run cmd/daemon/*.go -config=config.json

server:
	go run cmd/server/*.go -config=config.json

proto_gen_go:
	protoc -I . -I protos/ --go_out=./internal/proto --go-grpc_out=require_unimplemented_servers=false:./internal/proto $$(find protos/reddit -name "*.proto")

proto_gateway:
	protoc -I . -I protos/ --grpc-gateway_out=logtostderr=true:./internal/proto $$(find protos/reddit -name "*.proto")

proto_swagger:
	protoc -I . -I protos/ --openapiv2_out=json_names_for_fields=false,logtostderr=true,include_package_in_tags=true:./swagger $$(find protos/reddit -name "*.proto")