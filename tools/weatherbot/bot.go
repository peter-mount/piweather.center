package bot

import (
	"errors"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-mastodon"
	"github.com/peter-mount/piweather.center/io"
	"github.com/peter-mount/piweather.center/weather/state"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
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
	cleanup        *regexp.Regexp
}

func (t *Bot) Start() error {
	// Regex to replace errors when data is unavailable with "N/A"
	r, err := regexp.Compile("(%!f\\(string=\\W*?\\))")
	if err != nil {
		return err
	}
	t.cleanup = r

	// Path to lib directory for data lookup
	if *t.RootDir == "" {
		*t.RootDir = path.Join(filepath.Dir(os.Args[0]), "../etc")
	}

	err = t.getPost()
	if err != nil {
		return err
	}

	err = t.getCurrentState()
	if err != nil {
		return err
	}

	return nil
}

func (t *Bot) getPost() error {
	t.posts = make(map[string]*Post)
	if err := io.NewReader().
		Yaml(&t.posts).
		Open(filepath.Join(*t.RootDir, "weatherbot.yaml")); err != nil {
		return err
	}

	// Lookup post, show available posts & exit if not found
	t.post = t.posts[*t.Post]
	if *t.Post == "" || t.post == nil {
		a := append([]string{}, "Available posts:")
		for k, e := range t.posts {
			a = append(a, fmt.Sprintf("%s: %s", k, e.Name))
		}
		return errors.New(strings.Join(a, "\n"))
	}

	return nil
}

func (t *Bot) getCurrentState() error {
	// Get current state for the station for this post
	stn, err := state.New(*t.Host).GetState(t.post.StationId)
	if err != nil {
		return err
	}
	if stn == nil {
		return fmt.Errorf("StationId %q does not exist", t.post.StationId)
	}

	log.Printf("Station %q data at %v", stn.ID, stn.Meta.Time)
	t.station = stn
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
