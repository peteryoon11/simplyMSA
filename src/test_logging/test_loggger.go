package main

import (
	"log"
	"os"
)

var myLogger *log.Logger

func main() {
	// 로그파일 오픈
	fpLog, err := os.OpenFile("./log/logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()

	myLogger = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	//....
	run()

	myLogger.Println("End of Program")
}

func run() {
	myLogger.Print("Test")
}
