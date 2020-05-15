package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) handleIndex() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Welcome to the OfCO server maintainer")
	}
}
