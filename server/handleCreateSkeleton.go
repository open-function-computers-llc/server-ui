package server

import (
	"net/http"
	"os"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) handleCreateSkeleton() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		err := r.ParseForm()
		if err != nil {
			s.LogError(err.Error())
			return
		}

		domain := r.Form.Get("domain")
		s.Log(domain)

		cmd := exec.Command(os.Getenv("TOOLSBASEDIR") + "ofco-update-modx-skeleton.sh")
		cmd.Env = append(os.Environ(),
			"MODX_LIVE_DOMAIN="+domain,
		)
		err = cmd.Run()
		if err != nil {
			s.LogError(err.Error() + domain)
			return
		}

		w.Write([]byte("success!"))
	}
}
