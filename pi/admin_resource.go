package pi

import (
	"github.com/benschw/opin-go/rest"
	"net/http"
)

type AdminResource struct {
}

func (r *AdminResource) Status(res http.ResponseWriter, req *http.Request) {

	if err := rest.SetOKResponse(res, "OK"); err != nil {
		rest.SetInternalServerErrorResponse(res, err)
		return
	}
}
