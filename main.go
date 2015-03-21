package main

import (
	"flag"
	"github.com/benschw/pi/pi"
)

func main() {
	flag.Parse()

	svc := &pi.PiService{Bind: "0.0.0.0:8080"}

	if err := svc.Run(); err != nil {
		panic(err)
	}

}
