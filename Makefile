up:
	docker-compose up -d

run:
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

test:
	go test -v ./...

build-prod:
	GOOS=linux go build server/wellinformed.go
