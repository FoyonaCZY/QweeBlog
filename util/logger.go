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

func INFO(format string, v ...interface{}) {
	Println("INFO", fmt.Sprintf(format, v...))
}

func ERROR(format string, v ...interface{}) {
	Println("ERROR", fmt.Sprintf(format, v...))
}

func WARNING(format string, v ...interface{}) {
	Println("WARNING", fmt.Sprintf(format, v...))
}

func PANIC(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	Println("PANIC", msg)
	panic(msg)
}
