package pi

import (
	"github.com/benschw/opin-go/ophttp"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var _ = log.Printf

type PiService struct {
	Bind string
}

func (s *PiService) Run() error {
	defer log.Println("Exiting")

	bCtrl, err := NewBlinkCtrl()
	if err != nil {
		return err
	}
	defer bCtrl.Close()

	// route handlers
	admin := &AdminResource{}
	blink := &BlinkResource{Ctrl: bCtrl}

	// Configure Routes
	r := mux.NewRouter()

	r.HandleFunc("/admin/status", admin.Status).Methods("GET")

	r.HandleFunc("/blink/toggle", blink.Toggle).Methods("POST")
	r.HandleFunc("/blink/count-down", blink.CountDown).Methods("POST")

	http.Handle("/", r)

	// Start HTTP Server
	return ophttp.StartServer(s.Bind)
}
