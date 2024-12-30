package server

type Config struct {
	Apps      map[string]AppConfig `yaml:"apps"`
	Url       string               `yaml:"url"`
	LoadError error
}
type AppConfig struct {
	ClientID     string   `yaml:"clientID"`
	ClientSecret string   `yaml:"clinetSecret"`
	Issuer       string   `yaml:"issuer"`
	RedirectURIs []string `yaml:"redirectURIs"`
}
