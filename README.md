# Oh Yeah API
Lemonade stand API written in Go utilizing httprouter & GORM.

## Running The API
The only significant dependency is a postgres server. The connection-string
cannot be changed per-environment and may need to be tweaked directly in `Db.go`

The default is `postgres://postgres:password@host.docker.internal:5432/postgres`.
You can spin up a docker container for this purpose with

```sh
docker run --name oh-ya-postgres -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres
```

## Get Started
```sh
git clone git@github.com:provinite/oh-ya
cd oh-ya
go run .
```