package config

type PostgresConfig struct {
	Host     string
	Username string
	Password string
	DBName   string
	Port     string
	AppName  string
	SSLMode  string
	Timezone string
}

func NewPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     getEnvOrDefault("DB_HOST", "localhost"),
		Username: getEnvOrDefault("DB_USER", "postgres"),
		Password: getEnvOrDefault("DB_PASSWORD", "postgres"),
		DBName:   getEnvOrDefault("DB_NAME", "postgres"),
		Port:     getEnvOrDefault("DB_PORT", "5432"),
		AppName:  getEnvOrDefault("APP_NAME", "amethis-backend"),
		SSLMode:  getEnvOrDefault("DB_SSL_MODE", "disable"),
		Timezone: getEnvOrDefault("DB_TIMEZONE", "UTC"),
	}
}
