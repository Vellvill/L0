package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Db struct {
		Dsn            string `yaml:"dsn"`
		MigrationsPath string `yaml:"migrations_path"`
	} `yaml:"db"`
	Nats struct {
		Channel   string `yaml:"nats_channel"`
		ClusterID string `yaml:"cluster_id"`
		ClientID  string `yaml:"client_id"`
	} `yaml:"nats"`
}

func GetConfig() (*Config, error) {
	config := &Config{}
	err := cleanenv.ReadConfig("configs/config.yaml", config)
	if err != nil {
		_, err = cleanenv.GetDescription(config, nil)
		if err != nil {
			return nil, err
		}
	}
	return config, nil
}
