package stepper

// Control a stepper motor (28BJY-48)
//
// Datasheet:
// http://www.raspberrypi-spy.co.uk/wp-content/uploads/2012/07/Stepper-Motor-28BJY-48-Datasheet.pdf

import (
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
	"log"
)

const (
	CW  = iota
	CCW = iota
)

type Stepper struct {
	Pins     []embd.DigitalPin
	Sequence [][]int
	Position int
}

func NewStepper(p1 int, p2 int, p3 int, p4 int) (*Stepper, error) {
	if err := embd.InitGPIO(); err != nil {
		return nil, err
	}

	// Physical pins 11,15,16,18
	//               17,22,23,24
	// GPIO17,GPIO22,GPIO23,GPIO24
	stepPinNums := []int{p1, p2, p3, p4}

	stepPins := make([]embd.DigitalPin, 4)

	for i, pinNum := range stepPinNums {
		pin, err := embd.NewDigitalPin(pinNum)
		if err != nil {
			return nil, err
		}
		if err := pin.SetDirection(embd.Out); err != nil {
			return nil, err
		}
		if err := pin.Write(embd.Low); err != nil {
			return nil, err
		}

		stepPins[i] = pin
	}

	// Define sequence as shown in manufacturers datasheet
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

	st := &Stepper{
		Pins:     stepPins,
		Sequence: seq,
		Position: 0,
	}
	return st, nil
}

func (s *Stepper) Close() {
	for _, pin := range s.Pins {
		if err := pin.SetDirection(embd.In); err != nil {
			panic(err)
		}

		pin.Close()
	}
	embd.CloseGPIO()
}

func (s *Stepper) Step() error {

	log.Printf("Position: %d\n", s.Position)

	// set position
	for i, pin := range s.Pins {
		if s.Sequence[s.Position][i] != 0 {
			log.Printf("Enable pin %d\n", i)
			if err := pin.Write(embd.High); err != nil {
				return err
			}
		} else {
			if err := pin.Write(embd.Low); err != nil {
				return err
			}

		}
	}

	// s.Position incremented by 1 or 2 for clockwise, -1 or -2 for counter-clockwise
	s.Position++
	if s.Position >= len(s.Sequence) {
		s.Position = 0
	}

	return nil
}
