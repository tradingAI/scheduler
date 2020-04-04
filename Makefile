.PHONY: proto run test

proto:
	bash proto.sh

run:
	go run main/main.go

test:
	docker-compose -f docker-compose.yml run bazel 