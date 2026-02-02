
run:
	go run cmd/matchmaker/main.go

docker-build:
	docker build . -t go-practice

compose-local:
	docker-compose up redis

compose:
	docker-compose up --build
