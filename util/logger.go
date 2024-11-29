package util

import (
	"fmt"
	"time"
)

func InitLogger() {
	fmt.Println("日志服务已启动")
	fmt.Println(
		"\n" +
			"  ___                    ____  _             \n" +
			" / _ \\__      _____  ___| __ )| | ___   __ _ \n" +
			"| | | \\ \\ /\\ / / _ \\/ _ \\  _ \\| |/ _ \\ / _` |\n" +
			"| |_| |\\ V  V /  __/  __/ |_) | | (_) | (_| |\n" +
			" \\__\\_\\ \\_/\\_/ \\___|\\___|____/|_|\\___/ \\__, |\n" +
			"                                       |___/ ")
	fmt.Println()
}

func Println(prefix, msg string) {
	fmt.Printf("[%s][%s] %s\n",
		prefix,
		time.Now().Format("2006-01-02 15:04:05"),
		msg)
}

func Info(format string, v ...any) {
	Println("Info", fmt.Sprintf(format, v...))
}

func Error(format string, v ...any) {
	Println("Error", fmt.Sprintf(format, v...))
}

func Warning(format string, v ...any) {
	Println("Warning", fmt.Sprintf(format, v...))
}

func Panic(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	Println("Panic", msg)
	panic(msg)
}
