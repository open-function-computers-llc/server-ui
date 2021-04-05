package site

// LoadSiteByDomain - pass in a domain as a string and it should be found on
// this server. This will check that the folder exists in the standard hosting
// directory, that the unix user's home folder exists in the CONFIG_FILE_ROOT,
// and anything else that needs to happen to make sure the site is able to be
// maintaned by this application.
func LoadSiteByDomain(domain string) (Site, error) {
	s := Site{
		Domain: domain,
	}

	err := s.LoadStatus()

	return s, err
}
