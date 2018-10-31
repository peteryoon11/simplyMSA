package main

import (
	"fmt"
	"time"
)

var exitChain <-chan time.Time

func main() {
	daemonmain()
}

func daemonmain() {
	var (
		module_run bool = false
		//	err               error
		logRotateChan     <-chan time.Time
		logRotateTimer    time.Duration
		logrotate_seconds int64 = 0

		worker5minTimer time.Duration
		worker5minChan  <-chan time.Time
	)

	var exit bool = false

	var logrotate_interval int64 = 86400 // 하루 단위로
	_, time_offset := time.Now().Zone()
	logrotate_seconds = (time.Now().Unix() + int64(time_offset)) % logrotate_interval
	logrotate_seconds = logrotate_interval - logrotate_seconds
	logrotate_seconds += 1

	logRotateTimer = time.Second * time.Duration(logrotate_seconds)
	logRotateChan = time.NewTimer(logRotateTimer).C
	worker5minTimer = time.Second * time.Duration(30)

	fmt.Println("  Interval timer = ", worker5minTimer)
	worker5minChan = time.NewTimer(worker5minTimer).C

	for {
		if true == exit || false == module_run {
			break
		}
		select {
		case <-logRotateChan:
			//	report_worker_logger.RotateLogFiles()
			logrotate_seconds = (time.Now().Unix() + int64(time_offset)) % logrotate_interval
			logrotate_seconds = logrotate_interval - logrotate_seconds
			logrotate_seconds += 1
			logRotateTimer = time.Second * time.Duration(logrotate_seconds)
			logRotateChan = time.NewTimer(logRotateTimer).C
			break
		case <-worker5minChan:

			worker5minChan = time.NewTimer(worker5minTimer).C
			fmt.Println("now called")
			break

		case <-exitChain:
			exit = true
			module_run = false

			fmt.Println("Process End")
			return
		}
	}

}
