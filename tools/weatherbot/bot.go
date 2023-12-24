package bot

import (
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-mastodon"
	"github.com/peter-mount/piweather.center/store/api"
	"github.com/peter-mount/piweather.center/store/client"
	"github.com/peter-mount/piweather.center/util/config"
	_ "github.com/peter-mount/piweather.center/weather/measurement"
	"github.com/peter-mount/piweather.center/weather/state"
)

type Bot struct {
	ConfigManager  config.Manager `kernel:"inject"`
	Post           *string        `kernel:"flag,post,Post to issue"`
	Test           *bool          `kernel:"flag,test,If set then run in debug mode"`
	Host           *string        `kernel:"flag,host,Hostname of weathercenter,http://127.0.0.1:9001"`
	posts          map[string]*Post
	post           *Post
	result         *api.Result
	mastodonConfig mastodon.Config
	dbClient       client.Client
	client         mastodon.Client
	station        *state.Station
}

func (t *Bot) Start() error {
	// FIXME remove once refactoring is completed
	*t.Test = true

	t.dbClient.Url = *t.Host

	err := t.getPost()
	if err != nil {
		return err
	}

	query, err := ParsePost(t.post)
	if err != nil {
		return err
	}

	log.Println(query.Query)

	t.result, err = t.dbClient.Query(query.Query)
	if err != nil {
		return err
	}

	// Hack, only keep the last row in the results, as we sometimes get 2
	// rows
	for _, table := range t.result.Table {
		if len(table.Rows) > 1 {
			table.Rows = table.Rows[len(table.Rows)-1:]
		}
	}

	// Ensure cells have Value's populated
	t.result.Init()

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
