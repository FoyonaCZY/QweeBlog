package util

import (
	"fmt"
	"time"
)

func InitLogger() {
	fmt.Println("日志服务已启动")
	fmt.Println()
}

func Println(prefix, msg string) {
	fmt.Printf("[%s][%s] %s\n",
		prefix,
		time.Now().Format("2006-01-02 15:04:05"),
		msg)
}

func Info(format string, v ...interface{}) {
	Println("Info", fmt.Sprintf(format, v...))
}

func Error(format string, v ...interface{}) {
	Println("Error", fmt.Sprintf(format, v...))
}

func Warning(format string, v ...interface{}) {
	Println("Warning", fmt.Sprintf(format, v...))
}

func Panic(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	Println("Panic", msg)
	panic(msg)
}
