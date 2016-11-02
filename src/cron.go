package main

import (
	"fmt"
	"os"
	"log"
	"encoding/xml"
	"io/ioutil"
	"strings"
	"time"
	"regexp"
	"os/exec"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"crons"`
	Version     string   `xml:"version,attr"`
	Svs         []cron `xml:"cron"`
	Description string   `xml:",innerxml"`
}
type cron struct {
	XMLName    xml.Name `xml:"cron"`
	CronId string   `xml:"cronId"`
	CronDesc string   `xml:"cronDesc"`
	CronBash string   `xml:"cronBash"`
	CronPath   string   `xml:"cronPath"`
	CronName   string   `xml:"cronName"`
	CronInterval	int		`xml:"cronInterval"`
}
var commondString string

const TIMESLEEPINTERVAL  = 5
// 先声明map
var mContainer map[string]int32

func main() {
	//产生一个MAP
	mContainer = make(map[string]int32)
	//要改成读文件
	args := os.Args
	fmt.Println("read file who's name is %s",args[1])
	inputFile, inputError := os.Open(args[1])//变量指向os.Open打开的文件时生成的文件句柄
	if inputError!=nil {
		log.Fatalf("Read File error:%s",inputError)
	}
	xmlData, err := ioutil.ReadAll(inputFile)
	if(err!=nil){
		log.Fatalf("Read Xml error:%s",err)
	}
	v := Recurlyservers{}
	xml.Unmarshal(xmlData,&v)
	inputFile.Close()	//关闭文件
	if len(v.Svs) == 0{
		log.Fatalf("No script need to run……")
	}

	for i:=0; i<len(v.Svs);i++  {

		//CronId必须存在
		if v.Svs[i].CronId == ""{
			fmt.Println("CronId didn't null……")
			time.Sleep(TIMESLEEPINTERVAL * time.Second)
			continue
		}
		//CronId必须为字母数字
		fmt.Println(v.Svs[i].CronId)
		reg, _ := regexp.Compile(`[^a-zA-Z0-9_]`)
		if len(reg.FindAllString(v.Svs[i].CronId,-1)) >0 {
			fmt.Println("CronId must range [a-zA-Z0-9_]……")
			time.Sleep(TIMESLEEPINTERVAL * time.Second)
			continue
		}

		//**********
		cmd:= exec.Command("/bin/sh", "-c",`ps -ef |grep -v "grep" |grep "`+v.Svs[i].CronName+`"`)
		cmd.Stderr = os.Stdout
		cmd.Stderr = os.Stderr

		buf,err := cmd.Output()
		if(err != nil){
			fmt.Println(err)
		}
		fmt.Fprintf(os.Stdout, "Result: %s", buf)
		//**********

		//查找是否在进程中存在该程序
		/*
		rs := strings.Contains(buf,v.Svs[i].CronName)
		if(rs == true){
			//此次不执行
			fmt.Println("The process is running now")
			time.Sleep(TIMESLEEPINTERVAL * time.Second)
			continue
		}else{
			//继续执行下去
			runCmd := exec.Command(v.Svs[i].CronBash, v.Svs[i].CronPath+v.Svs[i].CronName)
			runCmd.Stderr = os.Stdout
			runCmd.Stderr = os.Stderr
			buf,err := runCmd.Output()
			if(err != nil){
				fmt.Println(err)
			}
			fmt.Fprintf(os.Stdout, "Result: %s", buf)
			time.Sleep(50 * time.Second)
		}*/
	}

	/*
	for {
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
		time.Sleep(50 * time.Second)
	}*/
}
