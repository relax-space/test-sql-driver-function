# test-sql-driver-function

## question
docker-entrypoint-initdb.d of `mysql docker` already supports `delimiter`, but `go-sql-driver` still does not support this syntax, resulting in failure to create sql functions

## solution

`multiStatements=true`

```
root:1234@tcp(127.0.0.1:3309)/fruit?charset=utf8&parseTime=True&loc=UTC&multiStatements=true
```

https://github.com/go-sql-driver/mysql/issues/1072

## start

1. test mysql(docker) docker-entrypoint-initdb.d 
``` shell
$ docker-compose up -d
$ docker logs -f mysql-fruit
```

2. test go-sql-driver
``` shell
$ go run .
```
output
``` shell
hello1 success {0xc0000c0000 0xc0000ba010}
hello2 success {0xc0000c0000 0xc0000ba040}
```