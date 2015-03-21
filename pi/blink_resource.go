package pi

import (
	"fmt"
	"github.com/benschw/opin-go/rest"
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
	"net/http"
	"time"
)

type BlinkResource struct {
	Led embd.LED
}

func (r *BlinkResource) Toggle(res http.ResponseWriter, req *http.Request) {
	if err := r.Led.Toggle(); err != nil {
		rest.SetInternalServerErrorResponse(res, err)
		return
	}
	fmt.Println("Toggled")
}

func (r *BlinkResource) CountDown(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Countdown Started")

	for sleep := 20; sleep > 0; sleep-- {
		for i := 0; i < sleep; i++ {
			time.Sleep(100 * time.Millisecond)
		}
		if err := r.Led.Toggle(); err != nil {
			rest.SetInternalServerErrorResponse(res, err)
			return
		}
	}
	for sleep := 30; sleep > 0; sleep-- {
		time.Sleep(50 * time.Millisecond)
		if err := r.Led.Toggle(); err != nil {
			rest.SetInternalServerErrorResponse(res, err)
			return
		}
	}
	r.Led.On()
	time.Sleep(2000 * time.Millisecond)
	r.Led.Off()

	fmt.Println("Countdown Complete")
	return

}
