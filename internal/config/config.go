package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Task TaskConfig `yaml:"task"`
}

type TaskConfig struct {
	MaxTasksAmount int `yaml:"max_concurrent_tasks" env:"MAX_CONCURRENT_TASKS" env-default:"3"`
}

func Load() *Config {
	cfg := &Config{}
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		panic(err)
	}

	fmt.Printf("Raw config data:\n%s\n", string(data))

	if err := yaml.Unmarshal(data, cfg); err != nil {
		fmt.Printf("Error unmarshaling config: %v\n", err)
		panic(err)
	}

	fmt.Printf("Parsed config: %+v\n", cfg)
	return cfg
}
