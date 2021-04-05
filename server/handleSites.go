package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) handleSites() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		sites := []string{}

		for _, site := range s.sites {
			sites = append(sites, site.Domain)
		}

		data := map[string]interface{}{
			"Sites": sites,
		}

		s.renderView("sites", w, &data)
	}
}
