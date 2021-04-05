package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/open-function-computers-llc/server-ui/site"
)

func (s *Server) handleSiteDetails() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		var site site.Site
		for _, loadedSite := range s.sites {
			if loadedSite.Domain == params.ByName("url") {
				site = loadedSite
			}
		}

		// get a fresh copy of the status from the JSON file
		site.LoadStatus()

		data := map[string]interface{}{
			"Domain": site.Domain,
			"Locked": site.LockedStatus,
		}

		s.renderView("details", w, &data)
	}
}
