package config

import (
	"ostium/pkg/repository"
	"ostium/pkg/wager"

	"github.com/spf13/viper"
)

type Config struct {
	WagerConfig wager.Config      `mapstructure:"wager"`
	DataBase    repository.Config `mapstructure:"database"`
}

func New(v *viper.Viper) (*Config, error) {
	cfg := &Config{}
	err := v.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, err
}
