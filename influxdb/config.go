package influxdb

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"sync"
)

type Client interface {
	WriteAPIBlocking() api.WriteAPIBlocking
}
type Config struct {
	Org              string `yaml:"org"`    // Organisation
	Bucket           string `yaml:"bucket"` // InfluxDB bucket
	Url              string `yaml:"url"`    // Url to influxdb
	Token            string `yaml:"token"`  // Authentication token
	mutex            sync.Mutex
	client           influxdb2.Client
	writeAPIBlocking api.WriteAPIBlocking
}

func (c *Config) getClient() influxdb2.Client {
	if c.client == nil {
		c.client = influxdb2.NewClient(c.Url, c.Token)
	}
	return c.client
}

func (c *Config) WriteAPIBlocking() api.WriteAPIBlocking {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.writeAPIBlocking == nil {
		c.writeAPIBlocking = c.getClient().WriteAPIBlocking(c.Org, c.Bucket)
	}
	return c.writeAPIBlocking
}
