package server

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) handleSiteStats() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// look up go access html file
		duration := params.ByName("duration")
		directory := os.Getenv("REPORT_DIR")
		if duration == "1" {
			report, err := ioutil.ReadFile(directory + "/" + params.ByName("url") + ".day.html")
			if err != nil {
				s.LogError(err.Error())
				return
			}
			w.Write(report)
			return
		}
		if duration == "30" {
			report, err := ioutil.ReadFile(directory + "/" + params.ByName("url") + ".month.html")
			if err != nil {
				s.LogError(err.Error())
				return
			}
			w.Write(report)
			return
		}
		if duration == "90" {
			report, err := ioutil.ReadFile(directory + "/" + params.ByName("url") + ".html")
			if err != nil {
				s.LogError(err.Error())
				return
			}
			w.Write(report)
			return
		}

		err := errors.New("Invalid duration requested for site analytics")
		s.LogError(err.Error())
		return
	}
}
