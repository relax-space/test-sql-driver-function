# test-sql-driver-function

## question
docker-entrypoint-initdb.d of `mysql docker` already supports `delimiter`, but `go-sql-driver` still does not support this syntax, resulting in failure to create sql functions

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
PS D:\source\gopath\src\test-sql-driver-function> go run .
hello1 success {0xc0000aa100 0xc00001e240}
panic: hello2 failure,Error 1064: You have an error in your SQL syntax; check the 
manual that corresponds to your MySQL server version for the right syntax to use near 'DELIMITER $$
CREATE FUNCTION hello2(labelIds varchar(500))
RETURNS varchar(1000' at line 3

goroutine 1 [running]:
main.main()
        D:/source/gopath/src/test-sql-driver-function/main.go:40 +0x2d1
exit status 2
```