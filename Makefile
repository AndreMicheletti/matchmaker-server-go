
run:
	go run main.go

docker-build:
	docker build . -t go-practice

compose-local: docker-build
	docker-compose up