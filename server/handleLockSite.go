package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) handleLockSite() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		data := map[string]interface{}{
			"Action": "lock",
			"Domain": params.ByName("url"),
		}
		s.renderView("scriptRunner", w, &data)
	}
}
