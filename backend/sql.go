package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func sqlOpen() {
	var err error
	db, err = sql.Open("postgres", "port=5432 user=postgres password=123456 dbname=golang sslmode=disable")
	//port是数据库的端口号，默认是5432，如果改了，这里一定要自定义；
	//user就是你数据库的登录帐号;
	//dbname就是你在数据库里面建立的数据库的名字;
	//sslmode就是安全验证模式;
	checkErr(err)
}

// func sqlInsert() {
// 	//插入数据
// 	stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid")
// 	checkErr(err)

// 	res, err := stmt.Exec("ficow", "软件开发部门", "2017-03-09")
// 	//这里的三个参数就是对应上面的$1,$2,$3了

// 	checkErr(err)

// 	affect, err := res.RowsAffected()
// 	checkErr(err)

// 	fmt.Println("rows affect:", affect)
// }
// func sqlDelete() {
// 	//删除数据
// 	stmt, err := db.Prepare("delete from userinfo where uid=$1")
// 	checkErr(err)

// 	res, err := stmt.Exec(1)
// 	checkErr(err)

// 	affect, err := res.RowsAffected()
// 	checkErr(err)

// 	fmt.Println("rows affect:", affect)
// }
// func sqlSelect() {
// 	//查询数据
// 	rows, err := db.Query("SELECT * FROM userinfo")
// 	checkErr(err)

// 	println("-----------")
// 	for rows.Next() {
// 		var uid int
// 		var username string
// 		var department string
// 		var created string
// 		err = rows.Scan(&uid, &username, &department, &created)
// 		checkErr(err)
// 		fmt.Println("uid = ", uid, "\nname = ", username, "\ndep = ", department, "\ncreated = ", created, "\n-----------")
// 	}
// }
// func sqlUpdate() {
// 	//更新数据
// 	stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
// 	checkErr(err)

// 	res, err := stmt.Exec("ficow", 1)
// 	checkErr(err)

// 	affect, err := res.RowsAffected()
// 	checkErr(err)

// 	fmt.Println("rows affect:", affect)
// }
func sqlClose() {
	db.Close()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// func sqlTest() {

// 	sep := "----------\n"
// 	sqlOpen()
// 	println(sep, "*sqlOpen")

// 	sqlSelect()
// 	println(sep, "*sqlSelect")

// 	sqlInsert()
// 	sqlSelect()
// 	println(sep, "*sqlInsert")

// 	sqlUpdate()
// 	sqlSelect()
// 	println(sep, "*sqlUpdate")

// 	sqlDelete()
// 	sqlSelect()
// 	println(sep, "*sqlDelete")

// 	sqlClose()
// 	println(sep, "*sqlClose")
// }

func main() {
	sqlOpen()
	sqlClose()

	r := gin.Default()

	r.GET("/go", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get",
		})
	})

	r.POST("/go", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "post",
		})
	})

	r.PUT("/go", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "put",
		})
	})

	r.DELETE("/go", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "delete",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
