package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Aws AWS
}

type AWS struct {
	AccessKeyID     string `env:"AWS_ACCESS_KEY_ID"`
	SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY"`
	Region          string `env:"AWS_REGION"`
	BucketName      string `env:"AWS_BUCKET_NAME"`
}

func New() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}
	return &cfg, nil
}
