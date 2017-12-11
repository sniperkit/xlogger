package config

import (
	"farm.e-pedion.com/repo/cheetah/client/mongo"
	"farm.e-pedion.com/repo/cheetah/logger"
	"fmt"
	"github.com/spf13/viper"
	path "path/filepath"
	"strings"
)

var (
	configuration *Configuration
)

// Configuration holds all the system parameters
type Configuration struct {
	Bind   string               `json:"bind" mapstructure:"bind"`
	Logger logger.Configuration `json:"logger" mapstructure:"logger"`
	Mongo  mongo.Configuration  `json:"mongo" mapstructure:"mongo"`
}

// Setup cofigures the package
func Setup(configPath string) error {
	ext := path.Ext(configPath)
	name := strings.TrimSuffix(path.Base(configPath), ext)
	path := path.Dir(configPath)
	fmt.Printf("LoadingConfig[Path=%s Name=%s Ext=%s Dir=%s]\n", configPath, name, ext, path)

	viper.SetConfigName(name)
	viper.SetConfigType(ext[1:])
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("config.SetupErr[Path=%s Err=%s]", configPath, err.Error())
	}

	if err := viper.Unmarshal(&configuration); err != nil {
		return fmt.Errorf("config.DecodeErr[Path=%s Err=%s]", configPath, err.Error())
	}
	fmt.Printf("config.Loaded[Instance=%+v]\n", configuration)
	return nil
}

// Get returns the currently setuped configuration
func Get() Configuration {
	return *configuration
}
