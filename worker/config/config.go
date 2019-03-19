package config

import "time"

type Config struct {
	UpdateInterval time.Duration `yaml:"UpdateInterval"`
	Count          int           `yaml:"Count"`
	Expiration     time.Duration `yaml:"  Expiration"`
	LogToFile      bool          `yaml:"LogToFile"`
	LogFilename    string        `yaml:"LogFilename"`
}

func (c *Config) ConvertNanosecondToSeconds() {
	c.UpdateInterval = c.UpdateInterval * time.Second
}

func (c *Config) Default() {
	c.UpdateInterval = time.Second * 30
	c.Count = 2
	c.Expiration = 60
	c.LogToFile = false
	c.LogFilename = "worker.log"
}
