package pi

import (
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
)

type BlinkCtrl struct {
	Led embd.LED
}

func NewBlinkCtrl() (*BlinkCtrl, error) {
	if err := embd.InitLED(); err != nil {
		return nil, err
	}
	led, err := embd.NewLED(0)
	if err != nil {
		return nil, err
	}

	c := &BlinkCtrl{
		Led: led,
	}

	return c, nil
}

func (b *BlinkCtrl) Close() error {
	embd.CloseLED()
	b.Led.Off()
	b.Led.Close()
	return nil
}
