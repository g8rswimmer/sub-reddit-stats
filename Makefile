
gen_mocks:
	mockgen -source=internal/worker/runner.go -destination=internal/worker/mock_runner_test.go -package=worker

init_db:
	mkdir db && touch db/sqlite-database.db

clean_db:
	rm -r db

migrate:
	go run cmd/migration/*.go

proto_gen_go:
	protoc -I . --go_out=./internal/proto --go-grpc_out=require_unimplemented_servers=false:./internal/proto $$(find protos -name "*.proto")