package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Port   string `env:"APP_PORT" envDefault:"8080"`
	Env    string `env:"APP_ENV" envDefault:"development"`
	Secret string `env:"APP_SECRET"`
	MySQL  MySQLConfig
	Redis  RedisConfig
}

type MySQLConfig struct {
	Host     string `env:"MYSQL_HOST" envDefault:"localhost"`
	Port     int    `env:"MYSQL_PORT" envDefault:"3306"`
	User     string `env:"MYSQL_USER" envDefault:"root"`
	Password string `env:"MYSQL_PASSWORD"`
	Database string `env:"MYSQL_DB" envDefault:"test"`
}

type RedisConfig struct {
	Host     string `env:"REDIS_HOST" envDefault:"localhost"`
	Port     int    `env:"REDIS_PORT" envDefault:"6379"`
	Password string `env:"REDIS_PASSWORD"`
	Database int    `env:"REDIS_DB" envDefault:"0"`
}

func InitConfig() *Config {
	var conf Config
	if err := env.Parse(&conf); err != nil {
		log.Fatal("MySQL config parse error:", err)
	}

	return &conf
}
