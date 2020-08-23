up:
	docker-compose up -d

run: pack-migrations
	go run server/wellinformed.go

#destroys db
down:
	docker-compose down

restart: down up wait run

wait:
	sleep 1

gql:
	gqlgen generate

test:
	go test -v ./...

build-prod: pack-migrations
	GOOS=linux go build server/wellinformed.go

pack-migrations:
	go-bindata -prefix "database/migrations/" -pkg migrations -o database/migrations/bindata.go database/migrations/

