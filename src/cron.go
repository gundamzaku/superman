package main

import (
	"conn"
	"fmt"
	"log"
	"time"
)

func main() {
	db := conn.Conn()
	for {
		rows, err := db.Query("select id,title from cron")
		if err != nil {
			log.Fatalf("Connect table cron error: %s\n", err)
		}
		var id int64
		var title string

		for rows.Next() {
			//数据
			rows.Scan(&id, &title)
			fmt.Println(id)
			//判断类型
		}
		time.Sleep(2 * time.Second)
	}
}
