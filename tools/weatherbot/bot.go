package bot

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-mastodon"
	"github.com/peter-mount/piweather.center/util/config"
	_ "github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/peter-mount/piweather.center/weather/state"
)

type Bot struct {
	ConfigManager  config.Manager `kernel:"inject"`
	Post           *string        `kernel:"flag,post,Post to issue"`
	Test           *bool          `kernel:"flag,test,If set then run in debug mode"`
	Host           *string        `kernel:"flag,host,Hostname of weathercenter,http://127.0.0.1:8080"`
	posts          map[string]*Post
	post           *Post
	station        *state.Station
	mastodonConfig mastodon.Config
	client         mastodon.Client
}

func (t *Bot) Start() error {
	err := t.getPost()
	if err != nil {
		return err
	}

	err = t.getCurrentState()
	if err != nil {
		return err
	}

	err = t.postText()
	if err != nil {
		return err
	}

	return nil
}

func (t *Bot) loginMastodon() error {

	if err := t.ConfigManager.ReadYaml("mastodon.yaml", &t.mastodonConfig); err != nil {
		return err
	}
	t.client = t.mastodonConfig.Client()

	log.Printf("Validating connection to %s", t.mastodonConfig.Server)
	app, err := t.client.VerifyCredentials()
	if err != nil {
		return err
	}
	log.Printf("Logged %s into %s", app.Name, t.mastodonConfig.Server)

	return nil
}
