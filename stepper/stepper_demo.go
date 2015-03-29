package main

import (
	"flag"
	"fmt"
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
	"os"
	"os/signal"
	"time"
)

func main() {
	stepWait := flag.Int("step-wait", 10, "milliseconds between steps")
	flag.Parse()

	if err := embd.InitGPIO(); err != nil {
		panic(err)
	}
	defer embd.CloseGPIO()

	// Physical pins 11,15,16,18
	// GPIO17,GPIO22,GPIO23,GPIO24
	stepPinNums := []int{17, 22, 23, 24}

	stepPins := make([]embd.DigitalPin, 4)

	for i, pinNum := range stepPinNums {
		pin, err := embd.NewDigitalPin(pinNum)
		if err != nil {
			panic(err)
		}
		defer pin.Close()
		if err := pin.SetDirection(embd.Out); err != nil {
			panic(err)
		}
		if err := pin.Write(embd.Low); err != nil {
			panic(err)
		}
		defer func() {
			if err := pin.SetDirection(embd.In); err != nil {
				panic(err)
			}
		}()

		stepPins[i] = pin
	}

	// Define advanced sequence
	// as shown in manufacturers datasheet
	seq := [][]int{
		[]int{1, 0, 0, 0},
		[]int{1, 1, 0, 0},
		[]int{0, 1, 0, 0},
		[]int{0, 1, 1, 0},
		[]int{0, 0, 1, 0},
		[]int{0, 0, 1, 1},
		[]int{0, 0, 0, 1},
		[]int{1, 0, 0, 1},
	}

	stepCount := len(seq) - 1
	stepDir := 2 // Set to 1 or 2 for clockwise, -1 or -2 for counter-clockwise

	// Initialise variables
	stepCounter := 0

	// Start main loop
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	defer signal.Stop(quit)

	for {
		select {
		case <-time.After(time.Duration(*stepWait) * time.Millisecond):
			for i, pin := range stepPins {
				if seq[stepCounter][i] != 0 {
					fmt.Printf("Enable pin %d, step %d\n", i, stepCounter)
					if err := pin.Write(embd.High); err != nil {
						panic(err)
					}
				} else {
					if err := pin.Write(embd.Low); err != nil {
						panic(err)
					}

				}
			}
			stepCounter += stepDir

			// If we reach the end of the sequence start again
			if stepCounter >= stepCount {
				stepCounter = 0
			} else if stepCounter < 0 {
				stepCounter = stepCount
			}
		case <-quit:
			return
		}
	}

}
