package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) checkSession(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, rp httprouter.Params) {
		session, err := s.sessions.Get(r, "admin-users")
		if err != nil {
			// TODO: set up better error handling
			s.Log("ERROR WITH SESSION")
			return
		}

		user := ""
		sessionVal := session.Values["user"]
		if u, ok := sessionVal.(string); ok {
			user = u
		} else {
			s.Log("ERROR WITH SESSION VALUE - 'user'")
			http.Redirect(w, r, s.routePrefix+"/login", 302)
			return
		}
		s.Log(user)
		s.Log("hit the middleware")

		next(w, r, rp)
	}
}
