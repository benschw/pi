package main

// Control a stepper motor (28BJY-48)
//
// Datasheet:
// http://www.raspberrypi-spy.co.uk/wp-content/uploads/2012/07/Stepper-Motor-28BJY-48-Datasheet.pdf
//
// this is a port of Matt Hawkins' example impl from
// http://www.raspberrypi-spy.co.uk/2012/07/stepper-motor-control-in-python/
// (link privides additional instructions for wiring your pi)

import (
	"flag"
	"github.com/benschw/pi/stepper"
	"os"
	"os/signal"
	"time"
)

func main() {
	stepWait := flag.Int("step-delay", 10, "milliseconds between steps")
	flag.Parse()

	st, err := stepper.NewStepper(17, 22, 23, 24)
	if err != nil {
		panic(err)
	}
	defer st.Close()

	// Start main loop
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	defer signal.Stop(quit)

	for {
		select {
		case <-time.After(time.Duration(*stepWait) * time.Millisecond):

			st.Step(stepper.CW)

		case <-quit:
			return
		}
	}

}
