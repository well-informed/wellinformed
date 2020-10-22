# wellinformed

## Install development dependencies

golang-migrate: Command line tool used to easily create versioned migration files

```
brew install golang-migrate
```

go-bindata: Binary packing tool so .sql migration files can be built into the deployed go binary
```
go get -u github.com/go-bindata/go-bindata/...
```

## Create a new migration file
In order to add new files to the database/migrations directory which are automatically named in conformance with the migration sequence numbers, run the below from the root dir of the project.
```
migrate create -ext sql -dir database/migrations -seq $MIGRATION_NAME
```

## Fix a Dirty Migration State
With golang-migrate installed on your machine and access to the postgres database run the following to clear out the dirty state so you can run migrations again.
```
migrate -path database/migrations -database "postgres://edyn:MPyDqCs4NCcCRe@edyn.c7xblzysdvfi.us-east-2.rds.amazonaws.com/edyn?sslmode=disable" force $LASTGOODVERSION
```

## Manually run DB Migration
This manually runs a migration so the error messages can be more easily checked in the case of a failed migration. This is equivalent to the migration step that is run internally when a new version is deployed through CI/CD

```
migrate -path database/migrations -database "postgres://edyn:MPyDqCs4NCcCRe@edyn.c7xblzysdvfi.us-east-2.rds.amazonaws.com/edyn?sslmode=disable" up
```