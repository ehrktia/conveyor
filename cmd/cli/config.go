package cli

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const defaultConfig = `
api:
    host: localhost
    port: "6286"
redis:
    host: localhost
    port: "6379"
etcd:
    host: http://etcd:2379
loki:
    host: http://loki:3100`

type Config struct {
	Api   *Httpconfig `yaml:"api"`
	Redis *Httpconfig `yaml:"redis"`
	Etcd  *HostConfig `yaml:"etcd"`
	Loki  *HostConfig `yaml:"loki"`
}

type HostConfig struct {
	Host string `yaml:"host"`
}

type Httpconfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func loadDefaultConfig() (Config, error) {
	c := &Config{}
	if err := yaml.Unmarshal([]byte(defaultConfig), c); err != nil {
		return Config{}, err
	}
	return *c, nil

}

func loadConfigFromFile(fileName string) (Config, error) {
	fileData, err := os.ReadFile(fileName)
	if err != nil && os.IsNotExist(err) {
		return Config{}, fmt.Errorf("%s-file not found", fileName)
	}
	c := &Config{}
	if err := yaml.Unmarshal(fileData, c); err != nil {
		return Config{}, err
	}
	return *c, nil
}
