package config

import (
	"strconv"
)

type RedisConfig struct {
	Host               string
	Port               string
	Password           string
	DB                 int
	TLSEnable          bool
	InsecureSkipVerify bool
}

func NewRedisConfig() RedisConfig {
	db, err := strconv.Atoi(getEnvOrDefault("REDIS_DB", "0"))
	if err != nil {
		panic(err)
	}

	return RedisConfig{
		Host:               getEnvOrDefault("REDIS_HOST", "localhost"),
		Port:               getEnvOrDefault("REDIS_PORT", "6379"),
		Password:           getEnvOrDefault("REDIS_PASSWORD", ""),
		DB:                 db,
		TLSEnable:          getEnvOrDefault("REDIS_TLS_ENABLE", "false") == "true",
		InsecureSkipVerify: getEnvOrDefault("REDIS_INSECURE_SKIP_VERIFY", "false") == "true",
	}
}
