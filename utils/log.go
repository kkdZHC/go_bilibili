package utils

import (
	"fmt"
	"io"
	"os"
	"time"
)

func CheckFile(Filename string) bool {
	exist := true
	_, err := os.Stat(Filename)
	if os.IsNotExist(err) {
		exist = false
		if err != nil {
			fmt.Println("not fount log")
		}
	}
	return exist
}

func Logfile(logType string, log string) {
	var file *os.File
	var err error

	filenames := "./file/logs/" + time.Now().Format("20060102") + ".log"

	if CheckFile(filenames) {
		file, err = os.OpenFile(filenames, os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println("文件存在已打开")
		}
	} else {
		file, err = os.Create(filenames)
		if err != nil {
			fmt.Println("文件创建失败")
		}
	}
	_, err = io.WriteString(file, logType+time.Now().Format("2006-01-02 15:04:05")+log+"\n") //写入文件
	if err != nil {
		fmt.Println(err)
	}
}
