package conn

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Conn() *sql.DB {
	db, err := sql.Open("mysql", "liudan:liudan123@tcp(127.0.0.1:3306)/weidaogou")
	//连接不上，报错
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
		//return false
	}
	//最后关闭数据库
	//defer db.Close()
	return db
}
