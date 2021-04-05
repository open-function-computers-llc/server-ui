package server

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) handleStats() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		statsJSON := "http://" + os.Getenv("STATS_HOST") + ":" + os.Getenv("STATS_PORT") + "/?token=" + os.Getenv("STATS_TOKEN")

		resp, err := http.Get(statsJSON)
		if err != nil {
			s.LogError(err.Error())
			io.WriteString(w, err.Error())
			return
		}
		defer resp.Body.Close()

		jsonBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			s.LogError(err.Error())
			io.WriteString(w, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	}
}
