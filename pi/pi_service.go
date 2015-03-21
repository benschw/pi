package pi

import (
	"github.com/gorilla/mux"
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
	"log"
	"net/http"
)

var _ = log.Printf

type PiService struct {
	Bind string
}

func (s *PiService) Run() error {

	if err := embd.InitLED(); err != nil {
		return err
	}
	defer embd.CloseLED()

	led, err := embd.NewLED(0)
	if err != nil {
		return err
	}
	defer func() {
		led.Off()
		led.Close()
	}()

	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, os.Interrupt, os.Kill)
	// defer signal.Stop(quit)

	// route handlers
	blink := &BlinkResource{Led: led}

	// Configure Routes
	r := mux.NewRouter()

	r.HandleFunc("/toggle", blink.Toggle).Methods("POST")
	r.HandleFunc("/count-down", blink.CountDown).Methods("POST")

	http.Handle("/", r)

	// Start HTTP Server
	return http.ListenAndServe(s.Bind, nil)
}
