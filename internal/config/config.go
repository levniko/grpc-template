package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	PostgreSQL PostgreSQLInstance
	Redis      RedisInstance
}

type PostgreSQLInstance struct {
	Master PostgreSQLConfig
	Slaves []PostgreSQLConfig
}

type PostgreSQLConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type RedisInstance struct {
	Master RedisConfig
	Slaves []RedisConfig
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

// New, belirtilen TOML dosyasından yapılandırma yükler.
func New(configPath string) (*Config, error) {
	var cfg Config
	if _, err := toml.DecodeFile(configPath, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
