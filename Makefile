up:
	docker-compose up -d

run:
	go install
	rss

#destroys db
down:
	docker-compose down

restart: down up

gql:
	gqlgen generate