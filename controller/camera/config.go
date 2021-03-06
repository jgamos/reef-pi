package camera

import (
	"fmt"
	"github.com/reef-pi/reef-pi/controller/utils"
	"time"
)

const DefaulCaptureFlags = ""

type Config struct {
	Enable         bool          `json:"enable" yaml:"enable"`
	ImageDirectory string        `json:"image_directory" yaml:"image_directory"`
	CaptureFlags   string        `json:"capture_flags" yaml:"capture_flags"`
	TickInterval   time.Duration `json:"tick_interval" yaml:"tick_interval"`
	Upload         bool          `json:"upload" yaml:"upload"`
}

var Default = Config{
	ImageDirectory: "/var/lib/reef-pi/images",
	TickInterval:   120,
}

func loadConfig(store utils.Store) (Config, error) {
	var conf Config
	return conf, store.Get(Bucket, "config", &conf)
}

func saveConfig(store utils.Store, conf Config) error {
	if conf.TickInterval <= 0 {
		return fmt.Errorf("Tick Interval for camera controller must be greater than zero")
	}
	if conf.ImageDirectory == "" {
		return fmt.Errorf("Image directory cant not be empty")
	}
	return store.Update(Bucket, "config", conf)
}
