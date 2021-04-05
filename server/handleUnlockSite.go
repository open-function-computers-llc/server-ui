package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) handleUnlockSite() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		data := map[string]interface{}{
			"Action": "unlock",
			"Domain": params.ByName("url"),
		}
		s.renderView("scriptRunner", w, &data)
	}
}
