# Authentication

## Clean Architecture in Go with gin-gonic and Unit Testing

This is an example of implementation clean code architecture in Go with gin-gonic framework with some unit testing on the project.

## Run the project

Run dependency manager

```console
make dep
```

Run local deployment

```console
make run
```

Run local unit test

```console
make test
```

Run build project

```console
make build
```

## Core library

Library | Usage
-- | --
gin | Base framework
gorm | ORM library
postgres | Database
jwt-go | JWT authorization
go-sqlmock | Database mock
logrus | Logger library
viper | Config library

And others library are listed on `go.mod` file
