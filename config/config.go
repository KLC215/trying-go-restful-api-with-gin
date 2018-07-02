package config

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// Initial configuration
	if err := c.initConfig(); err != nil {
		return err
	}

	// Monitoring configuration file changes and hot-reloading program
	c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" { // Parse specified configuration
		viper.SetConfigFile(c.Name)
	} else { // Parse default configuration
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yaml")
	// Matching environment variable
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APISERVER")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

// Hot-reloading configuration file
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}
