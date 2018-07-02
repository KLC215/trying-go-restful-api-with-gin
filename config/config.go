package config

import (
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
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
	// Initial logging package
	c.initLog()

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

func (c *Config) initLog() {
	passLagerCfg := log.PassLagerCfg{
		// file: save log to LoggerFile
		// stdout: display standard output
		Writers: viper.GetString("log.writers"),
		// DEBUG, INFO, WARN, ERROR, FATAL
		LoggerLevel: viper.GetString("log.logger_level"),
		// Logger file name
		LoggerFile: viper.GetString("log.logger_file"),
		// true: json format
		// false: plain text
		LogFormatText: viper.GetBool("log.log_format_text"),
		// daliy: save by day
		// size: save by size
		RollingPolicy: viper.GetString("log.rolling_policy"),
		// pair with RollingPolicy: daliy
		LogRotateDate: viper.GetInt("log.log_rotate_date"),
		// pair with RollingPolicy: size
		LogRotateSize: viper.GetInt("log.log_rotate_size"),
		// Backup zip file count
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}

	log.InitWithConfig(&passLagerCfg)
}

// Hot-reloading configuration file
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
	})
}
