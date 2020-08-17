package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) handleSites() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		s.renderView("sites", w, nil)
	}
}
