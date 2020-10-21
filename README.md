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

