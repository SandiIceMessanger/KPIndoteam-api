package config

import "os"

var (
	ServerPort      = GetEnv("SERVER_PORT", "9000")
	MongoUrl        = GetEnv("MONGODB_URL", "mongodb://chatnews:chatnews2022!!@172.16.255.11:27017")
	MongoDatabase   = GetEnv("MONGODB_DATABASE", "kpi_dev")
	JWTSecret       = GetEnv("JWT_SECRET", "bermaslaah")
	JWTExpirationMs = GetEnv("JWT_EXPIRATION_MS", "86400000")
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
