package pi

import (
	"fmt"
	"github.com/benschw/opin-go/rest"
	"net/http"
	"time"
)

type BlinkResource struct {
	Ctrl *BlinkCtrl
}

func (r *BlinkResource) Toggle(res http.ResponseWriter, req *http.Request) {
	if err := r.Ctrl.Led.Toggle(); err != nil {
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
		if err := r.Ctrl.Led.Toggle(); err != nil {
			rest.SetInternalServerErrorResponse(res, err)
			return
		}
	}
	for sleep := 30; sleep > 0; sleep-- {
		time.Sleep(50 * time.Millisecond)
		if err := r.Ctrl.Led.Toggle(); err != nil {
			rest.SetInternalServerErrorResponse(res, err)
			return
		}
	}
	r.Ctrl.Led.On()
	time.Sleep(2000 * time.Millisecond)
	r.Ctrl.Led.Off()

	fmt.Println("Countdown Complete")
	return

}
