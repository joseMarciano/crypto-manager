package config

import (
	"strings"

	"github.com/spf13/viper"
)

type (
	Configuration struct {
		Database Database
		Nats     Nats
		Server   Server
	}

	Loader struct {
		loaders []loaderConfig
	}

	properties map[string]string

	loaderConfig func(*Configuration, properties)
)

func NewLoader() *Loader {
	return &Loader{
		loaders: []loaderConfig{databaseLoader, natsLoader, serverLoader},
	}
}

func (l Loader) Load() (Configuration, error) {
	cfg := new(Configuration)
	props, err := l.localProps()
	if err != nil {
		return Configuration{}, err
	}

	for _, loader := range l.loaders {
		loader(cfg, props)
	}

	return *cfg, nil
}

func (Loader) localProps() (properties, error) {
	viper.SetConfigName("env_local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	props := make(properties)
	for _, key := range viper.AllKeys() {
		localPropValue := viper.GetString(key)
		if localPropValue != "" {
			props[strings.ToUpper(key)] = localPropValue
		}
	}

	return props, nil
}
