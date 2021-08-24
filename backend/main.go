package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func sqlOpen() {
	db, err = sql.Open("postgres", "port=5432 user=postgres password=123456 dbname=golang sslmode=disable")
	//port是数据库的端口号，默认是5432，如果改了，这里一定要自定义；
	//user就是你数据库的登录帐号;
	//dbname就是你在数据库里面建立的数据库的名字;
	//sslmode就是安全验证模式;
	checkErr(err)
}

func sqlInsert(c *gin.Context) {
	//插入数据
	stmt, err := db.Prepare("insert into comment(username, content, created) values ($1,$2,current_date) returning id")
	checkErr(err)

	res, err := stmt.Exec("a", "评论")
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("rows affect:", affect)
}

func sqlSelect() {
	//查询数据
	rows, err := db.Query("select * from comment")
	checkErr(err)

	println("-----------")
	for rows.Next() {
		var id int
		var content string
		err = rows.Scan(&id, &content)
		checkErr(err)
		fmt.Println("content = ", content)
	}
}

func sqlClose() {
	db.Close()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	sqlOpen()
	sqlSelect()
	sqlClose()
}
