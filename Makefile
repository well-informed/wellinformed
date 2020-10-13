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

gen:
	go generate ./pagination
	gqlgen generate

test: pack-migrations up
	go test -v ./...

build-prod: pack-migrations gen
	GOOS=linux go build server/wellinformed.go

pack-migrations:
	go-bindata -prefix "database/migrations/" -pkg migrations -o database/migrations/bindata.go database/migrations/

