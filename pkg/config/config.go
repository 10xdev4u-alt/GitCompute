package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	GitHubToken string `mapstructure:"github_token"`
	Owner       string `mapstructure:"owner"`
	Repo        string `mapstructure:"repo"`
	WorkflowID  string `mapstructure:"workflow_id"`
}

// LoadConfig loads the configuration from viper and validates required fields
func LoadConfig() (*Config, error) {
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if cfg.GitHubToken == "" {
		return nil, errors.New("github_token is required")
	}
	if cfg.Owner == "" {
		return nil, errors.New("owner is required")
	}
	if cfg.Repo == "" {
		return nil, errors.New("repo is required")
	}
	if cfg.WorkflowID == "" {
		// Default to worker.yml if not specified
		cfg.WorkflowID = "worker.yml"
	}

	return &cfg, nil
}
