package main

import (
	"conn"
	"fmt"
	"log"
	"time"
	"os/exec"
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
			//执行PHP脚本
			f, err := exec.Command("ls", "/").Output()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(f)
		}
		time.Sleep(5 * time.Second)
	}
}
