package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3309)/fruit?charset=utf8&parseTime=True&loc=UTC&multiStatements=true")
	if err != nil {
		panic(err)
	}
	sql := `
DROP FUNCTION IF EXISTS hello1;
CREATE FUNCTION hello1 (s CHAR(20))
RETURNS CHAR(50) DETERMINISTIC
RETURN CONCAT('Hello1, ',s,'!');
`
	rows, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	fmt.Println("hello1 success", rows) //success

	sql = `
DROP FUNCTION IF EXISTS hello2;

CREATE FUNCTION hello2(labelIds varchar(500)) 
RETURNS varchar(1000)
BEGIN
    RETURN CONCAT('Hello, ',s,'!');
END
`
	rows, err = db.Exec(sql)
	if err != nil { //failure
		err = fmt.Errorf("hello2 failure,%v", err)
		panic(err)
	}
	fmt.Println("hello2 success", rows) //success
}
