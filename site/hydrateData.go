package site

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func (s *Site) hydrateData() error {
	directory := os.Getenv("CONFIG_FILE_ROOT")
	status, err := ioutil.ReadFile(directory + "/" + s.Domain + "/settings.json")
	if err != nil {
		return err
	}
	var siteSettings settings
	err = json.Unmarshal(status, &siteSettings)
	if err != nil {
		return err
	}

	// now we hydrate the data from the json
	s.LockedStatus = siteSettings.SiteIsLocked
	s.UptimeURL = siteSettings.UptimeURI

	return nil
}
