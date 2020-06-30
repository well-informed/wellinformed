up:
	docker-compose up -d

run:
	go run server/server.go

#destroys db
down:
	docker-compose down

restart: down up wait run

wait:
	sleep 1

gql:
	gqlgen generate