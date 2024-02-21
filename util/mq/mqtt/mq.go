package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/peter-mount/go-kernel/v2/log"
	"sync"
	"time"
)

type MQ struct {
	// Url to connect to
	Url      string `yaml:"url"`
	ClientID string `yaml:"clientID"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Debug    bool   `yaml:"debug"`
	// ===== Internal
	mutex         sync.Mutex                // Mutex
	name          string                    // Name of broker in config
	client        mqtt.Client               // active Client
	subscriptions map[string]MessageHandler // Subscriptions
}

func (s *MQ) Log(f string, a ...interface{}) {
	if s.Debug {
		log.Printf("MQTT:"+s.name+":"+f, a...)
	}
}

func (s *MQ) connect() error {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(s.Url)
	opts.SetClientID(s.ClientID)
	opts.SetUsername(s.Username)
	opts.SetPassword(s.Password)
	opts.DefaultPublishHandler = s.publishHandler
	opts.OnConnect = s.onConnect
	opts.OnConnectionLost = s.onConnectionLost

	s.client = mqtt.NewClient(opts)

	// Now connect. This will return nil if successful, or an error if it
	// either failed to connect or timed out
	s.Log("Connecting...")
	return s.wait(s.client.Connect(), "connect")
}

func (s *MQ) wait(token mqtt.Token, action string) error {
	if token.WaitTimeout(time.Minute) {

		if token.Error() != nil {
			s.Log("%s failed %v", action, token.Error())
			return token.Error()
		}

		s.Log("%s successful", action)
		return nil
	}

	return fmt.Errorf("MQTT:%s: %s timed out", s.name, action)
}

func (s *MQ) onConnect(_ mqtt.Client) {
	s.Log("%s connected", s.ClientID)
}

func (s *MQ) onConnectionLost(_ mqtt.Client, err error) {
	s.Log("%s disconnected: %v", s.ClientID, err)
}
