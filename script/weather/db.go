package weather

import (
	"github.com/peter-mount/go-script/packages"
	"github.com/peter-mount/piweather.center/store/client"
)

func init() {
	packages.Register("weatherdb", &DB{})
}

type DB struct{}

func (_ DB) Connect(url string) *client.Client {
	return &client.Client{Url: url}
}
