package server

import (
	"net/http"
	"os"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) handleCreate() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		err := r.ParseForm()
		if err != nil {
			s.LogError(err.Error())
			return
		}

		domain := r.Form.Get("domain")
		skeleton := r.Form.Get("skeleton")
		s.Log(domain)

		cmd := exec.Command(os.Getenv("TOOLSBASEDIR") + "ofco-deploy-modx-apache.sh")
		cmd.Env = append(os.Environ(),
			"SKELETON_DOMAIN="+skeleton,
			"DOMAIN="+domain,
		)
		err = cmd.Run()
		if err != nil {
			s.LogError(err.Error())
			return
		}

		w.Write([]byte("success!"))
	}
}
