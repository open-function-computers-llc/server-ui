package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) checkSession(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, rp httprouter.Params) {
		session, err := s.sessions.Get(r, "ofco-server-session")
		if err != nil {
			// TODO: set up better error handling
			s.Log("ERROR WITH SESSION")
			return
		}

		sessionVal := session.Values["authenticated"]
		if authenticated, ok := sessionVal.(bool); !ok || !authenticated {
			s.Log("ERROR WITH SESSION VALUE - 'authenticated'", authenticated, sessionVal)
			http.Redirect(w, r, s.routePrefix+"/login", 302)
			return
		}

		next(w, r, rp)
	}
}
