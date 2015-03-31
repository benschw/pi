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
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	stepWait := flag.Int("step-delay", 10, "milliseconds between steps")
	steps := flag.Int("steps", 0, "number of steps, 0 for infinite")
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

	// Start main loop
	ticker := time.NewTicker(time.Duration(*stepDelay) * time.Millisecond)

	i := *steps
	for {
		select {
		case <-ticker.C:
			if i == 0 {
				ticker.Stop()
				return
			}
			i--
			st.Step()

		case <-quit:
			ticker.Stop()
			return
		}
	}

}
