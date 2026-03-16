package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/theOldZoom/gofm/internal/api"

	"go.yaml.in/yaml/v3"
)

func Save(cfg *Config) error {
	path, err := Path()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func ValidateAPIKey(apiKey string) error {
	if strings.TrimSpace(apiKey) == "" {
		return fmt.Errorf("API key is required")
	}

	return api.ValidateAPIKey(apiKey)
}

func ValidateUsername(username string, apiKey string) error {
	if strings.TrimSpace(username) == "" {
		return fmt.Errorf("username is required")
	}

	return api.ValidateUsername(username, apiKey)
}

func ValidationMessage(err error) string {
	var apiErr *api.APIError
	if errors.As(err, &apiErr) {
		return apiErr.Message
	}
	return err.Error()
}
