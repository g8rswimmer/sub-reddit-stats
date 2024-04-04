
gen_mocks:
	mockgen -source=internal/worker/runner.go -destination=internal/worker/mock_runner_test.go -package=worker