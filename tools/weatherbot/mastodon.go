package bot

import (
	"encoding/json"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/go-mastodon"
)

func (t *Bot) postToMastodon(status mastodon.PostStatus) error {
	if t.client == nil {
		if err := t.loginMastodon(); err != nil {
			return err
		}
	}

	result, err := t.client.Post(status)
	if err != nil {
		return err
	}

	log.Println(result)

	b, err := json.Marshal(result)
	if err != nil {
		return err
	}

	log.Println(string(b))
	return nil
}
