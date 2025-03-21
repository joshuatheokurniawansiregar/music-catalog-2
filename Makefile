mock:
	go generate ./...
compose:
	docker-compose up
run:
	go run cmd/main.go
