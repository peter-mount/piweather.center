package config

import (
	"errors"
	"fmt"
	"github.com/peter-mount/go-kernel/v2/log"
	"github.com/peter-mount/piweather.center/station"
	"github.com/peter-mount/piweather.center/util/config"
	"path/filepath"
	"strings"
)

const (
	dashDir = "stations"
)

// Config checks that configuration files are syntactically correct.
type Config struct {
	Config   config.Manager    `kernel:"inject"`
	Stations *station.Stations `kernel:"inject"`
	Suffix   *string           `kernel:"flag,check,Check configuration"`
	Debug    *bool             `kernel:"flag,check-debug,Perform additional checks"`
	dashDir  string
}

func (r *Config) Run() error {
	if *r.Suffix != "" {
		return r.check(*r.Suffix)
	}
	return nil
}

func (r *Config) check(suffix string) error {

	r.dashDir = filepath.Join(r.Config.EtcDir(), dashDir)
	log.Printf("Directory %q", r.dashDir)

	loadOpt := loadOption(suffix)
	log.Printf("Load options: %04x", loadOpt)
	if loadOpt == station.NoOption {
		return fmt.Errorf("unknown suffix %q", suffix)
	}
	loadOpt = loadOpt | station.TestOption

	allFailed := false
	for _, suff := range suffixes(loadOpt) {
		failed, err := r.checkConfig(loadOpt, suff)
		if err != nil {
			return err
		}
		allFailed = allFailed || failed
	}

	if allFailed {
		return errors.New("no config found, might require -rootDir set")
	}

	return nil
}

func (r *Config) checkConfig(loadOpt station.LoadOption, suffix string) (bool, error) {
	log.Printf("Checking %q", suffix)
	stations, err := r.Stations.LoadDirectory(r.dashDir, suffix, loadOpt)

	if err == nil && *r.Debug {
		fmt.Println(stations.String())
	}

	if err != nil {
		s := err.Error()
		if strings.HasPrefix(s, "no config files found") {
			return true, nil
		}
		return true, err
	}

	return false, nil
}
