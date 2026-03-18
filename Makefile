build:
	go build -o twinflow ./cmd/twinflow

docker-build:
	docker build -t twinflow .

run-record:
	docker run --network=host -v $(pwd)/captures:/captures twinflow record --target http://localhost:9000

run-replay:
	docker run --network=host -v $(pwd)/captures:/captures twinflow replay --target http://localhost:9001