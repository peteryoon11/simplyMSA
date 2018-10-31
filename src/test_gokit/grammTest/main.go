package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var (
		httpAddr     = flag.String("http.addr", ":8000", "Address for HTTP (JSON) server")
		consulAddr   = flag.String("consul.addr", "", "Consul agent address")
		retryMax     = flag.Int("retry.max", 3, "per-request retries to different instances")
		retryTimeout = flag.Duration("retry.timeout", 500*time.Millisecond, "per-request timeout, including retries")
	)
	//flag.Parse()
	fmt.Println(httpAddr)
	fmt.Println(consulAddr)
	fmt.Println(retryMax)
	fmt.Println(retryTimeout)
}
