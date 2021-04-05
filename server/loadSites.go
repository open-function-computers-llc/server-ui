package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/open-function-computers-llc/server-ui/site"
)

func (s *Server) loadSites() error {
	type StatsShape struct {
		Uptime string `json:"uptime"`
		Discs  []struct {
			Size        string `json:"size"`
			MountPoint  string `json:"mountPoint"`
			Used        string `json:"used"`
			Free        string `json:"free"`
			UsedPercent string `json:"usedPercent"`
		} `json:"discs"`
		RAM struct {
			Mem struct {
				Available string `json:"available"`
				BuffCache string `json:"buff/cache"`
				Free      string `json:"free"`
				Shared    string `json:"shared"`
				Total     string `json:"total"`
				Used      string `json:"used"`
			} `json:"mem"`
			Swap struct {
				Free  string `json:"free"`
				Total string `json:"total"`
				Used  string `json:"used"`
			} `json:"swap"`
		} `json:"ram"`
		LoadAverages struct {
			OneMinute      string `json:"one-minute"`
			FiveMinutes    string `json:"five-minutes"`
			FifteenMinutes string `json:"fifteen-minutes"`
		} `json:"loadAverages"`
		AdditionalCommands []struct {
			Command string   `json:"command"`
			Output  []string `json:"output"`
		} `json:"additional-commands"`
	}
	statsJSON := "http://" + os.Getenv("STATS_HOST") + ":" + os.Getenv("STATS_PORT") + "/?token=" + os.Getenv("STATS_TOKEN")

	resp, err := http.Get(statsJSON)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	stats := StatsShape{}
	json.Unmarshal(jsonBytes, &stats)

	for _, domain := range stats.AdditionalCommands[0].Output {
		site, err := site.LoadSiteByDomain(domain)
		s.logger.Info(site.Domain + " loaded")
		if err != nil {
			return err
		}
		s.sites = append(s.sites, site)
	}
	return nil
}
