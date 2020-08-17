package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) handleLogin() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		err := r.ParseForm()
		if err != nil {
			s.LogError(err.Error())
			return
		}

		s.Log(r.Form.Get("username"))
		s.Log(r.Form.Get("password"))
		if r.Form.Get("username") != "admin" || r.Form.Get("password") != "password" {
			s.Log("bad creds!")
			http.Redirect(w, r, s.routePrefix+"/login", 302)
			return
		}
		s.Log("time to log in!")
	}
}

func (s *Server) showLoginForm() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		s.renderView("login", w, nil)
	}
}
