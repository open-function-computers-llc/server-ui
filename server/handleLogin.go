package server

import (
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) handleLogin() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		err := r.ParseForm()
		if err != nil {
			s.LogError(err.Error())
			return
		}

		if r.Form.Get("username") != os.Getenv("APP_USERNAME") || r.Form.Get("password") != os.Getenv("APP_PASSWORD") {
			s.Log("bad creds!")
			http.Redirect(w, r, s.routePrefix+"/login", 302)
			return
		}
		session, _ := s.sessions.Get(r, "ofco-server-session")
		session.Values["authenticated"] = true

		err = session.Save(r, w)
		if err != nil {
			s.LogError(err.Error())
			return
		}

		http.Redirect(w, r, s.routePrefix+"/", 302)
	}
}

func (s *Server) showLoginForm() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		s.renderView("login", w, nil)
	}
}
