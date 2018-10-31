package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kardianos/service"
)

var exitChain <-chan time.Time

type program struct {
}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	daemonmain()
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	exitChain = time.NewTimer(time.Second * 2).C
	<-time.After(time.Second * 9)
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "msa testing client",
		DisplayName: "msa testing client",
		Description: "msa testing client",
	}

	prg := &program{}

	s, err := service.New(prg, svcConfig)
	if err != nil {
		fmt.Println(err)

	}
	if len(os.Args) > 1 {

		if strings.EqualFold(os.Args[1], "-V") {
			fmt.Println("")
			return
		}

	}

	err = s.Run()
	if err != nil {
		fmt.Println(err)

	}
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
	worker5minTimer = time.Second * time.Duration(5)

	fmt.Println("  Interval timer = ", worker5minTimer)
	worker5minChan = time.NewTimer(worker5minTimer).C
	module_run = true
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
