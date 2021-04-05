package site

// Site is a representation of a single website managed by a server
type Site struct {
	Domain          string
	CanonicalURL    string
	UptimeURL       string
	LockedStatus    bool
	OpenDirectories []string
}

// GetSite - make sure this domain is a valid one, and if so, return a new site
// instance after reading it's state from the JSON config file
func GetSite(url string) (Site, error) {
	s, err := verify(url)
	return s, err
}

func verify(url string) (Site, error) {
	s := Site{}
	return s, nil
}
