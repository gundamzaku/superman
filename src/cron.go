package main

import (
	"conn"
	"fmt"
	"log"
	"time"
	"os/exec"
	"os"
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
			//cmd := exec.Command("php", "/data/www/wei/script/test.php")
			cmd:= exec.Command("/bin/sh", "-c", `ps -ef | grep -v "grep" | grep "queue"`)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			err = cmd.Run()
			if err != nil {
				fmt.Println(err)
			}

			f, err := cmd.Output()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(f)
		}
		time.Sleep(5 * time.Second)
	}
}
