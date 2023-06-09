package bot

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-mastodon"
	"github.com/peter-mount/piweather.center/io"
	_ "github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/peter-mount/piweather.center/weather/state"
	"os"
	"path"
	"path/filepath"
)

type Bot struct {
	RootDir        *string `kernel:"flag,rootDir,Location of config files"`
	Post           *string `kernel:"flag,post,Post to issue"`
	Test           *bool   `kernel:"flag,test,If set then run in debug mode"`
	Host           *string `kernel:"flag,host,Hostname of weathercenter,http://127.0.0.1:8080"`
	posts          map[string]*Post
	post           *Post
	station        *state.Station
	mastodonConfig mastodon.Config
	client         mastodon.Client
}

func (t *Bot) Start() error {
	// Path to lib directory for data lookup
	if *t.RootDir == "" {
		*t.RootDir = path.Join(filepath.Dir(os.Args[0]), "../etc")
	}

	err := t.getPost()
	if err != nil {
		return err
	}

	err = t.getCurrentState()
	if err != nil {
		return err
	}

	err = t.createPostText()
	if err != nil {
		return err
	}

	return nil
}

func (t *Bot) loginMastodon() error {

	if err := io.NewReader().
		Yaml(&t.mastodonConfig).
		Open(filepath.Join(*t.RootDir, "mastodon.yaml")); err != nil {
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
