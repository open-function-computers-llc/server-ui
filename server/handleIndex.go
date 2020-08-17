package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) handleIndex() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		s.renderView("dashboard", w, nil)
	}
}
