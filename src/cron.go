package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"time"
	"regexp"
	"runtime"
	"os/exec"
	"strings"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"crons"`
	Version     string   `xml:"version,attr"`
	Cron        []cron `xml:"cron"`
	Description string   `xml:",innerxml"`
}
type cron struct {
	XMLName      xml.Name `xml:"cron"`
	CronId       string   `xml:"cronId"`
	CronDesc     string   `xml:"cronDesc"`
	CronBash     string   `xml:"cronBash"`
	CronPath     string   `xml:"cronPath"`
	CronName     string   `xml:"cronName"`
	CronParam    string   `xml:"cronParam"`
	CronInterval int64    `xml:"cronInterval"`
}
const TIMESLEEPINTERVAL = 3
/*
声明一个脚本类
 */
type Cron struct {
	mContainer map[string]int64
	buf []byte
	view Recurlyservers
}

/*
打开文件
*/
func (cron *Cron) OpenFile() {

	//接受外部参数
	args := os.Args
	cron.Show(1, "read file who's name is %s", args[1])

	inputFile, inputError := os.Open(args[1])//变量指向os.Open打开的文件时生成的文件句柄
	if inputError != nil {
		log.Fatalf("Read File error:%s", inputError)
	}
	xmlData, err := ioutil.ReadAll(inputFile)
	if (err != nil) {
		log.Fatalf("Read Xml error:%s", err)
	}
	//解析XML
	xml.Unmarshal(xmlData, &cron.view)
	inputFile.Close()    //关闭文件
	if len(cron.view.Cron) == 0 {
		log.Fatalf("No script need to run……")
	}
	//return cron.view
}

/*
打开文件
*/
func (cron *Cron) Show(status int, format string, a ...interface{}) {

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
func (cron *Cron)ByteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}

//启动
func (cron *Cron) Run() {

	var err bool

	for i := 0; i < len(cron.view.Cron); i++ {
		err = cron.CheckCronId(i)
		if(err == false){
			continue
		}
		//分配时间间隔
		err = cron.SetInterVal(i,cron.view.Cron[i].CronId)

		if(err == false){
			continue
		}
		cron.Show(1, "regular expression have done,pass")

		if runtime.GOOS != "windows" {
			//查看进程
			err = cron.CheckProcess(i)
			if(err == false){
				continue
			}
			cron.RunProcess(i)
		}
	}
}
func (cron *Cron) RunProcess(i int){
	//继续执行下去
	cron.Show(1, "exec: %s %s", cron.view.Cron[i].CronPath + cron.view.Cron[i].CronName,cron.view.Cron[i].CronParam)
	runCmd := exec.Command(cron.view.Cron[i].CronBash, cron.view.Cron[i].CronPath + cron.view.Cron[i].CronName,cron.view.Cron[i].CronParam,">>/tmp/cron.txt &")

	runCmd.Stderr = os.Stdout
	runCmd.Stderr = os.Stderr
	runCmd.Start()	//这个是表示不阻塞的，还有阻塞的，用run()
	buf, err := runCmd.Output()
	if (err != nil) {
		cron.Show(4,"%s",err)
	}
	//runCmd.Wait()
	cron.Show(1,"Result: %s", buf)
	//fmt.Fprintf(os.Stdout, "Result: %s", buf)
}

func (cron *Cron) CheckProcess(i int) bool{

	cmd := exec.Command("/bin/sh", "-c", `ps -ef |grep -v "grep" |grep "` + cron.view.Cron[i].CronName + ` `+cron.view.Cron[i].CronParam+`"`)
	cmd.Stderr = os.Stdout
	cmd.Stderr = os.Stderr
	buf, err := cmd.Output()
	if (err != nil) {
		cron.Show(4, "%s",err)
	}

	s := cron.ByteString(buf)
	rs := strings.Contains(s, cron.view.Cron[i].CronName)
	if (rs == true) {
		//此次不执行
		cron.Show(2, "The process is running now")
		time.Sleep(TIMESLEEPINTERVAL * time.Second)
		return false
	}
	return true
}

func (cron *Cron) SetInterVal(i int,cronId string) bool {
	//存入MAP，记录时间间隔
	if (cron.mContainer[cronId] == 0) {
		//初始化，执行时间
		cron.Show(1, "init runtime")
		cron.mContainer[cronId] = time.Now().Unix()
	}

	if (cron.mContainer[cronId] > time.Now().Unix()) {
		//还未到执行时间
		cron.Show(1, "runtime is not coming")
		time.Sleep(TIMESLEEPINTERVAL * time.Second)
		return false
	} else {
		//向后退移执行时间
		cron.mContainer[cronId] += cron.view.Cron[i].CronInterval
		cron.Show(2, "delay runtime to %d", cron.mContainer[cronId])
	}
	return true
}

//启动
func (cron *Cron) CheckCronId(i int) bool {
	//CronId必须存在
	if cron.view.Cron[i].CronId == "" {
		cron.Show(3, "CronId didn't null……")
		time.Sleep(TIMESLEEPINTERVAL * time.Second)
		return false
	}
	//CronId必须为字母数字
	reg, _ := regexp.Compile(`[^a-zA-Z0-9_]`)
	if len(reg.FindAllString(cron.view.Cron[i].CronId, -1)) > 0 {
		cron.Show(3, "CronId must range [a-zA-Z0-9_]……")
		time.Sleep(TIMESLEEPINTERVAL * time.Second)
		return false
	}
	return true
}

func main()  {
	//构造
	cron := Cron{make(map[string]int64),nil,Recurlyservers{}}
	cron.OpenFile()

	for{
		cron.Show(1,"loop %d",time.Now().Unix())
		cron.Run()
		time.Sleep(TIMESLEEPINTERVAL * time.Second)
	}
}