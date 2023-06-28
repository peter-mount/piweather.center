package comqtt

import (
	rv8 "github.com/go-redis/redis/v8"
	"github.com/peter-mount/go-build/version"
	"github.com/peter-mount/go-kernel/v2/log"
	configManager "github.com/peter-mount/piweather.center/util/config"
	colog "github.com/wind-c/comqtt/v2/cluster/log/zero"
	"github.com/wind-c/comqtt/v2/config"
	"github.com/wind-c/comqtt/v2/mqtt"
	"github.com/wind-c/comqtt/v2/mqtt/hooks/auth"
	"github.com/wind-c/comqtt/v2/mqtt/hooks/storage/badger"
	"github.com/wind-c/comqtt/v2/mqtt/hooks/storage/bolt"
	"github.com/wind-c/comqtt/v2/mqtt/hooks/storage/redis"
	"go.etcd.io/bbolt"
	"os"
	"time"
)

// CoMQTT provides a MQTT Broker
type CoMQTT struct {
	configManager configManager.Manager `kernel:"inject"`
	config        config.Config
	server        *mqtt.Server
}

func (s *CoMQTT) Start() error {
	if log.IsVerbose() {
		log.Println(version.Version)
	}

	conf := &s.config
	cm := s.configManager

	if err := s.configManager.ReadYaml("comqtt.yaml", conf); err != nil {
		return err
	}

	// Ensure that any local paths in the config point to
	// the distribution's "etc" directory
	cm.FixPath(&conf.Auth.ConfPath)
	cm.FixPath(&conf.Auth.BlacklistPath)

	cm.FixPath(&conf.Mqtt.Tls.CACert)
	cm.FixPath(&conf.Mqtt.Tls.ServerCert)
	cm.FixPath(&conf.Mqtt.Tls.ServerKey)

	// TODO these should be in a log directory not etc
	cm.FixPath(&conf.Log.InfoFile)
	cm.FixPath(&conf.Log.ErrorFile)
	cm.FixPath(&conf.Log.ThirdpartyFile)

	// Init log
	if hn, err := os.Hostname(); err == nil {
		s.config.Log.NodeName = hn
	}
	s.config.Mqtt.Options.Logger = colog.Init(s.config.Log)

	// Setup server
	s.server = mqtt.New(&s.config.Mqtt.Options)

	return nil
}

func (s *CoMQTT) Run() error {

	err := s.server.Serve()

	if err == nil {
		err = s.initStorage()
	}

	if err == nil {
		err = s.initAuth()
	}

	if err == nil {
		err = s.initBridge()
	}

	if err != nil {
		return err
	}

	// The Kernel manages signal handling, so as CoMQTT runs in the background
	// we need to keep the boot thread here forever otherwise we will just exit
	done := make(chan bool, 1)
	<-done

	return nil
}

func (s *CoMQTT) Stop() {
	if s.server != nil {
		s.server.Log.Info().Msg("Stopping server")
		_ = s.server.Close()
		s.server = nil
	}
}

func (s *CoMQTT) initAuth() error {
	server := s.server
	conf := s.config

	if conf.Auth.Way == config.AuthModeAnonymous {
		return server.AddHook(new(auth.AllowHook), nil)
	}

	/* FIXME implement
	 if conf.Auth.Way == config.AuthModeUsername || conf.Auth.Way == config.AuthModeClientid {
		switch conf.Auth.Datasource {
		case config.AuthDSRedis:
			opts := rauth.Options{}
			onError(plugin.LoadYaml(conf.Auth.ConfPath, &opts), logMsg)
			onError(server.AddHook(new(rauth.Auth), &opts), logMsg)
		case config.AuthDSMysql:
			opts := mauth.Options{}
			onError(plugin.LoadYaml(conf.Auth.ConfPath, &opts), logMsg)
			onError(server.AddHook(new(mauth.Auth), &opts), logMsg)
		case config.AuthDSPostgresql:
			opts := pauth.Options{}
			onError(plugin.LoadYaml(conf.Auth.ConfPath, &opts), logMsg)
			onError(server.AddHook(new(pauth.Auth), &opts), logMsg)
		case config.AuthDSHttp:
			opts := hauth.Options{}
			onError(plugin.LoadYaml(conf.Auth.ConfPath, &opts), logMsg)
			onError(server.AddHook(new(hauth.Auth), &opts), logMsg)
		}*/

	return config.ErrAuthWay
}

func (s *CoMQTT) initStorage() error {
	server := s.server
	conf := s.config

	switch conf.StorageWay {
	case config.StorageWayBolt:
		return server.AddHook(new(bolt.Hook), &bolt.Options{
			Path: conf.StoragePath,
			Options: &bbolt.Options{
				Timeout: 500 * time.Millisecond,
			},
		})

	case config.StorageWayBadger:
		return server.AddHook(new(badger.Hook), &badger.Options{
			Path: conf.StoragePath,
		})

	case config.StorageWayRedis:
		return server.AddHook(new(redis.Hook), &redis.Options{
			HPrefix: conf.Redis.HPrefix,
			Options: &rv8.Options{
				Addr:     conf.Redis.Options.Addr,
				DB:       conf.Redis.Options.DB,
				Password: conf.Redis.Options.Password,
			},
		})

	}

	return nil
}

func (s *CoMQTT) initBridge() error {
	conf := s.config

	switch conf.BridgeWay {
	case config.BridgeWayNone:
		// Do nothing

	case config.BridgeWayKafka:
		// FIXME implement
		//opts := cokafka.Options{}
		//onError(plugin.LoadYaml(conf.BridgePath, &opts), logMsg)
		//onError(s.server.AddHook(new(cokafka.Bridge), &opts), logMsg)
	}

	return nil
}
