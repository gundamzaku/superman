package main

import (
	"conn"
	"fmt"
	"log"
	"time"
	"os/exec"
	"os"
	"io/ioutil"
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
			stdout, err := cmd.StdoutPipe()
			if err != nil {
				fmt.Println("StdoutPipe: " + err.Error())
				return
			}

			stderr, err := cmd.StderrPipe()
			if err != nil {
				fmt.Println("StderrPipe: ", err.Error())
				return
			}

			if err := cmd.Start(); err != nil {
				fmt.Println("Start: ", err.Error())
				return
			}

			bytesErr, err := ioutil.ReadAll(stderr)
			if err != nil {
				fmt.Println("ReadAll stderr: ", err.Error())
				return
			}

			if len(bytesErr) != 0 {
				fmt.Printf("stderr is not nil: %s", bytesErr)
				return
			}

			bytes, err := ioutil.ReadAll(stdout)
			if err != nil {
				fmt.Println("ReadAll stdout: ", err.Error())
				return
			}

			if err := cmd.Wait(); err != nil {
				fmt.Println("Wait: ", err.Error())
				return
			}

			fmt.Printf("stdout: %s", bytes)
		}
		time.Sleep(5 * time.Second)
	}
}
