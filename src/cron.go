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
	//要改成读文件
	//db := conn.Conn()
	for {
		/*
		rows, err := db.Query("select id,title from cron")
		if err != nil {
			log.Fatalf("Connect table cron error: %s\n", err)
		}*/
		//var id int64
		//var title string

		//for rows.Next() {
			//数据
			//rows.Scan(&id, &title)
			//fmt.Println(id)
			//判断类型
			//执行PHP脚本
			//cmd := exec.Command("php", "/data/www/wei/script/test.php")
			//cmd:= exec.Command(`ps -ef | grep -v "grep" | grep "queue"`)
			cmd:= exec.Command("/bin/sh", "-c",`ps -ef |grep -v "grep" |grep "batch_qmfx_order"`)
			cmd.Stderr = os.Stdout
			cmd.Stderr = os.Stderr

			buf,err := cmd.Output()
			if(err != nil){
				fmt.Println(err)
			}
			fmt.Fprintf(os.Stdout, "Result: %s", buf)
			//需要确认是否有这个进程，有的话中断。

			//还要判断一下数据库里的数据是否匹配

			//执行可以执行的任务
			runCmd := exec.Command("php", "/data/www/wei/script/batch_qmfx_order.php")
			runCmd.Stderr = os.Stdout
			runCmd.Stderr = os.Stderr
			buf,err = runCmd.Output()
			if(err != nil){
				fmt.Println(err)
			}
			fmt.Fprintf(os.Stdout, "Result: %s", buf)
		//}
		time.Sleep(50 * time.Second)
	}
}
