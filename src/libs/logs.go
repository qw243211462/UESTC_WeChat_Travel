package libs

import (
	"flag"
	"os"
	"io/ioutil"
	"log"
	"io"
)



var (
	logFileName = flag.String("log", "/Users/laiye/Documents/wangmaotong/laiye/UESTC_WeChat_Travel/test.log", "Log file name")
	Trace   *log.Logger // 记录所有日志
	Info    *log.Logger // 重要的信息
	Warning *log.Logger // 需要注意的信息
	Error   *log.Logger // 致命错误a
)


func LogInfo() {
	file, err := os.OpenFile("/Users/laiye/Documents/wangmaotong/laiye/UESTC_WeChat_Travel/test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Trace = log.New(ioutil.Discard, "TRACE: ", log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "Info: ", log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "Warning: ", log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stderr),  "Error", log.Ltime|log.Lshortfile)
}

func init() {
	LogInfo()
}


