package config

import (
	"fmt"
	"os"
	"gopkg.in/yaml.v3"
)
// Config содержит параметры запуска сервера.
// Изучите config.yaml и добавьте поля самостоятельно.
type Config struct{
	ServerHost             string `yaml:"server_host"`
	ServerPort             int    `yaml:"server_port"`
	LogLevel               string `yaml:"log_level"`
}

// Load читает конфигурацию из файла config.yaml.
// Если файл не найден или поле не задано, применяются значения по умолчанию.
func Load() (*Config, error) {
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		
		return nil, fmt.Errorf("не удалось прочитать файл конфигурации: %w", err)
	}
	cfg := &Config{}
	
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, fmt.Errorf("не удалось распарсить YAML конфигурации: %w", err)
	}
	return cfg, nil
}
