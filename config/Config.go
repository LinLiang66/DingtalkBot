package config

import (
	"encoding/json"
)

type Config struct {
	AppName     string      `json:"app_name"`
	AppModel    string      `json:"app_model"`
	AppHost     string      `json:"app_host"`
	AppPort     string      `json:"app_port"`
	RedisConfig RedisConfig `json:"redis_config"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Addr     string `json:"addr"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

// 存储配置的全局对象
var cfg *Config = nil

func ChangeConfig(Jsonstr string) (*Config, error) {

	err := json.Unmarshal([]byte(Jsonstr), &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
