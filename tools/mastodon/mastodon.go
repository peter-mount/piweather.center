package mastodon

type Mastodon struct {
	config *Config `kernel:"config,mastodon"`
}

type Config struct {
	ClientKey    string `yaml:"client_key"`
	ClientSecret string `yaml:"client_secret"`
	AccessToken  string `yaml:"access_token"`
}

type Toot struct {
	Status string `json:"status"`
}

// Toot publishes a toot
func (m *Mastodon) Toot(toot Toot) error {

}
