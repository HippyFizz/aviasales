package config

import (
	cacheConfig "aviasales/cache/config"
	aviasalesConfig "aviasales/modules/aviasales/config"
	webConfig "aviasales/web/config"
	workerConfig "aviasales/worker/config"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Configuration interface {
	Default()
}

type Config struct {
	Server   webConfig.Config       `yaml:"Server"`
	Redis    cacheConfig.Config     `yaml:"Redis"`
	Avisales aviasalesConfig.Config `yaml:"Aviasales"`
	Worker   workerConfig.Config    `yaml:"Worker"`
	Debug    bool                   `yaml:"Debug"`
}

var (
	conf   Config
	logger = log.New(os.Stdout, "config: ", log.LstdFlags)
)

func init() {
	conf = Config{}
}

func (c *Config) setDefaults() {
	c.Debug = true
	c.Server.Default()
	c.Redis.Default()
	c.Worker.Default()
	c.Avisales.Default()
}

func (c *Config) loadViaLocalFile(filepath string) {
	var (
		configData []byte
		err        error
	)
	configData, err = ioutil.ReadFile(filepath)

	if err != nil {
		configData, err = c.generate()
		err = ioutil.WriteFile(filepath, configData, os.ModePerm)
		if err != nil {
			logger.Print(fmt.Sprintf("Can't write file to filepath %s", filepath))
		}
	}

	err = yaml.Unmarshal(configData, c)

	if err != nil {
		panic(err)
	}
	logger.Print(fmt.Sprintf("Using local file configuration filepath: %s", filepath))

}

func (c *Config) generate() ([]byte, error) {
	c.setDefaults()

	result, err := yaml.Marshal(c)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func Manager(filename string) *Config {
	conf.loadViaLocalFile(filename)
	return &conf
}
