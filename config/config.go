package config

import (
	"time"

	"github.com/spf13/viper"
)

// Provider defines a set of read-only methods for accessing the application
// configuration params as defined in one of the config files.
type Provider interface {
	ConfigFileUsed() string
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	InConfig(key string) bool
	IsSet(key string) bool
}

var defaultConfig *viper.Viper

// Config returns a default config providers
func Config() Provider {
	return defaultConfig
}

// LoadConfigProvider returns a configured viper instance
func LoadConfigProvider(appName string) Provider {
	return readViperConfig(appName)
}

func init() {
	defaultConfig = readViperConfig("TODO_SERVE")
}

// 重新读取配置文件
func ReadViperConfigFromFile(configPath string) error {
	defaultConfig.SetConfigFile(configPath)
	return defaultConfig.ReadInConfig()
}

func readViperConfig(appName string) *viper.Viper {
	v := viper.New()
	v.SetEnvPrefix(appName)
	v.AutomaticEnv()

	// global defaults

	v.SetDefault("json_logs", false)
	v.SetDefault("loglevel", "debug")
	v.SetDefault("db.user", "root")
	v.SetDefault("db.password", "root")
	v.SetDefault("db.addr", "127.0.0.1:3306")
	v.SetDefault("db.name", "test")

	v.SetDefault("http.mode", "release")
	v.SetDefault("http.port", 8080)
	v.SetDefault("http.read_timeout", 30)
	v.SetDefault("http.write_timeout", 30)

	return v
}
