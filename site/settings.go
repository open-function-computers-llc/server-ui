package site

// this is the structure of the JSON file that we will use for configuration.
// make sure this matches the bash scripts for reading and writting settings.
type settings struct {
	Domain                    string   `json:"domain"`
	AlwaysUnlockedDirectories []string `json:"alwaysUnlockedDirectories"`
	CanonicalURI              string   `json:"canonicalURI"`
	SiteIsLocked              bool     `json:"siteIsLocked"`
	UptimeURI                 string   `json:"uptimeURI"`
	AlternateDomains          []string `json:"alternateDomains"`
}
