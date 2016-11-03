package main

import (
	"fmt"
	"os"
	"log"
	"encoding/xml"
	"io/ioutil"
	"time"
	"regexp"
	"os/exec"
	"strings"
	"runtime"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"crons"`
	Version     string   `xml:"version,attr"`
	Svs         []cron `xml:"cron"`
	Description string   `xml:",innerxml"`
}
type cron struct {
	XMLName      xml.Name `xml:"cron"`
	CronId       string   `xml:"cronId"`
	CronDesc     string   `xml:"cronDesc"`
	CronBash     string   `xml:"cronBash"`
	CronPath     string   `xml:"cronPath"`
	CronName     string   `xml:"cronName"`
	CronInterval int64    `xml:"cronInterval"`
}

const TIMESLEEPINTERVAL = 3
// 先声明map
var mContainer map[string]int64
var buf []byte = nil

func main() {
	//产生一个MAP
	mContainer = make(map[string]int64)
	//要改成读文件
	args := os.Args
	show(1, "read file who's name is %s", args[1])
	inputFile, inputError := os.Open(args[1])//变量指向os.Open打开的文件时生成的文件句柄
	if inputError != nil {
		log.Fatalf("Read File error:%s", inputError)
	}
	xmlData, err := ioutil.ReadAll(inputFile)
	if (err != nil) {
		log.Fatalf("Read Xml error:%s", err)
	}
	v := Recurlyservers{}
	xml.Unmarshal(xmlData, &v)
	inputFile.Close()    //关闭文件
	if len(v.Svs) == 0 {
		log.Fatalf("No script need to run……")
	}
	for {
		show(1,"\n\n\nloop_____________________________")
		for i := 0; i < len(v.Svs); i++ {

			//CronId必须存在
			if v.Svs[i].CronId == "" {
				show(3, "CronId didn't null……")
				time.Sleep(TIMESLEEPINTERVAL * time.Second)
				continue
			}
			//CronId必须为字母数字
			reg, _ := regexp.Compile(`[^a-zA-Z0-9_]`)
			if len(reg.FindAllString(v.Svs[i].CronId, -1)) > 0 {
				show(3, "CronId must range [a-zA-Z0-9_]……")
				time.Sleep(TIMESLEEPINTERVAL * time.Second)
				continue
			}

			//存入MAP，记录时间间隔
			if (mContainer[v.Svs[i].CronId] == 0) {
				//初始化，执行时间
				show(1, "init runtime")
				mContainer[v.Svs[i].CronId] = time.Now().Unix()
			}

			if (mContainer[v.Svs[i].CronId] > time.Now().Unix()) {
				//还未到执行时间
				show(1, "runtime is not coming")
				time.Sleep(TIMESLEEPINTERVAL * time.Second)
				continue
			} else {
				//向后退移执行时间
				mContainer[v.Svs[i].CronId] += v.Svs[i].CronInterval
				show(2, "delay runtime to %d", mContainer[v.Svs[i].CronId])
			}

			show(1, "regular expression have done,pass")

			if runtime.GOOS != "windows" {
				cmd := exec.Command("/bin/sh", "-c", `ps -ef |grep -v "grep" |grep "` + v.Svs[i].CronName + `"`)
				cmd.Stderr = os.Stdout
				cmd.Stderr = os.Stderr

				buf, err := cmd.Output()
				if (err != nil) {
					show(4, "%s",err)
				}
				show(1, "Result: %s", buf)
			}

			//查找是否在进程中存在该程序
			s := byteString(buf)
			rs := strings.Contains(s, v.Svs[i].CronName)
			if (rs == true) {
				//此次不执行
				show(2, "The process is running now")
				time.Sleep(TIMESLEEPINTERVAL * time.Second)
				continue
			} else {
				if runtime.GOOS != "windows" {
					//继续执行下去
					show(1, "exec: %s", v.Svs[i].CronPath + v.Svs[i].CronName)
					runCmd := exec.Command(v.Svs[i].CronBash, v.Svs[i].CronPath + v.Svs[i].CronName,"&")
					runCmd.Stderr = os.Stdout
					runCmd.Stderr = os.Stderr
					buf, err := runCmd.Output()
					if (err != nil) {
						show(4,"%s",err)
					}
					fmt.Fprintf(os.Stdout, "Result: %s", buf)
				}
				time.Sleep(TIMESLEEPINTERVAL * time.Second)
			}
		}
	}
}

func byteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}

func show(status int, format string, a ...interface{}) {
	if (status == 1) {
		fmt.Print("[Info]    ")
	} else if (status == 2) {
		fmt.Print("[Notice]  ")
	} else if (status == 3) {
		fmt.Print("[Warning] ")
	} else if (status == 4) {
		fmt.Print("[Error]   ")
	}

	fmt.Fprintf(os.Stdout, format, a...)
	fmt.Println("")
}
