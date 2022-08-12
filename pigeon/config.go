package pigeon

import (
	"github.com/spf13/viper"
)

// BrokerConfig contains the required information for connecting to the MQTT broker
// Currently supports unencrypted, unauthorized broker only
type BrokerConfig struct {
	Url  string
	Port string
}

// DatabaseConfig contains the required information for connecting, writing, and querying a database.
// Currently supports InfluxDB Cloud
type DatabaseConfig struct {
	Url          string
	Organisation string
	TokenRead    string
	TokenWrite   string
}

// SiteConfig contains all devices on site and their value types
// Used to configure MQTT subscriptions
type SiteConfig struct {
	Name    string
	Devices []string
}

// Config is pigeons's full configuration
type Config struct {
	Broker   BrokerConfig
	Database DatabaseConfig
	Sites    []SiteConfig
}

// Override print method for readability
func (c *Config) Strings() {

}

// Reads in config from viper
// Sets up viper to watch for changes and update config accordingly
func (c *Config) ReadInConfig(cfgFile string) {
	if cfgFile == "" {
		// Use default config file /etc/ditto/config/pigeon.yaml
		c.ReadInConfig("/etc/ditto/config/pigeon.yaml")
	}
	viper.SetConfigFile(cfgFile)
	viper.ReadInConfig()
	viper.Unmarshal(&c)
	viper.WatchConfig()
}
