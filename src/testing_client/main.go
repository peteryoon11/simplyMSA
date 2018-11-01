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

		IntervalValue     time.Duration
		IntervalValueChan <-chan time.Time
	)

	var exit bool = false

	IntervalValue = time.Second * time.Duration(5)

	fmt.Println("  Interval timer = ", IntervalValue)
	IntervalValueChan = time.NewTimer(IntervalValue).C
	module_run = true
	for {
		if true == exit || false == module_run {
			break
		}
		select {

		case <-IntervalValueChan:

			IntervalValueChan = time.NewTimer(IntervalValue).C
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
