.PHONY: proto run test

proto:
	bash proto.sh

run:
	go run main/main.go

mock_runner:
	go run experiments/main.go

test:
	docker-compose -f docker-compose.yml run bazel 