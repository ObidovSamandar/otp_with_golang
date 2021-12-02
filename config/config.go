package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	PORT               string
	JWT_SECRET_KEY     string
	PG_HOST            string
	PG_PORT            string
	PG_USER            string
	PG_PASSWORD        string
	PG_DB              string
	SSlMode            string
	OTPDigit           int
	OTPPool            string
	MINIO_ENDPOINT     string
	MINIO_ACCESSKEY_ID string
	MINIO_SECRET_KEY   string
	MINIO_SSL_MODE     bool
	MINIO_BUCKET_NAME  string
}

func Load() Config {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found!")
	}

	config := Config{}

	config.PORT = cast.ToString(getOrReturnDefault("PORT", 8000))
	config.JWT_SECRET_KEY = cast.ToString(getOrReturnDefault("JWT_SECRET_KEY", ""))
	config.PG_HOST = cast.ToString(getOrReturnDefault("PG_HOST", ""))
	config.PG_PORT = cast.ToString(getOrReturnDefault("PG_PORT", ""))
	config.PG_PASSWORD = cast.ToString(getOrReturnDefault("PG_PASSWORD", ""))
	config.PG_DB = cast.ToString(getOrReturnDefault("PG_DB", ""))
	config.PG_USER = cast.ToString(getOrReturnDefault("PG_USER", ""))
	config.SSlMode = cast.ToString(getOrReturnDefault("SSLMode", "disable"))
	config.OTPDigit = cast.ToInt(getOrReturnDefault("OTP_DIGITS", 6))
	config.OTPPool = cast.ToString(getOrReturnDefault("OTP_POOL", "0123456789"))
	config.MINIO_ENDPOINT = cast.ToString(getOrReturnDefault("MINIO_ENDPOINT", ""))
	config.MINIO_ACCESSKEY_ID = cast.ToString(getOrReturnDefault("MINIO_ACCESSKEY_ID", ""))
	config.MINIO_SECRET_KEY = cast.ToString(getOrReturnDefault("MINIO_SECRET_KEY", ""))
	config.MINIO_SSL_MODE = cast.ToBool(getOrReturnDefault("MINIO_SSL_MODE", ""))
	config.MINIO_BUCKET_NAME = cast.ToString(getOrReturnDefault("MINIO_BUCKET_NAME", ""))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
