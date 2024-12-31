package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Storage    Storage    `mapstructure:"storage"`
	HttpServer HttpServer `mapstructure:"http_server"`
	EcbApi     EcbApi     `mapstructure:"ecb_api"`
}

type Storage struct {
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	User       string `mapstructure:"user"`
	Password   string `mapstructure:"password"`
	Database   string `mapstructure:"database"`
	Collection string `mapstructure:"collection"`
}

type HttpServer struct {
	Address     string        `mapstructure:"address"`
	Timeout     time.Duration `mapstructure:"timeout"`
	IdleTimeout time.Duration `mapstructure:"idle_timeout"`
}

type EcbApi struct {
	Url      string        `mapstructure:"url"`
	PageSize int           `mapstructure:"pagesize"`
	Max      int           `mapstructure:"max"`
	Period   time.Duration `mapstructure:"period"`
	Name     string        `mapstructure:"name"`
}

func Read(configPath string) (Config, error) {
	if configPath == `` {
		return Config{}, errors.New(`config path is empty`)
	}

	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("error read config file: %s", err)
	}

	var conf Config

	if err := viper.Unmarshal(&conf); err != nil {
		return Config{}, fmt.Errorf("error unmarshal config file: %s", err)
	}

	return conf, nil
}
