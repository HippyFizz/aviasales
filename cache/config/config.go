package config

import (
	"crypto/tls"
	"time"
)

type Config struct {
	Network            string        `yaml:"Network"`
	Addr               string        `yaml:"Addr"`
	Password           string        `yaml:"Password"`
	DB                 int           `yaml:"DB"`
	MaxRetries         int           `yaml:"MaxRetries"`
	MinRetryBackoff    time.Duration `yaml:"MinRetryBackoff"`
	MaxRetryBackoff    time.Duration `yaml:"MaxRetryBackoff"`
	DialTimeout        time.Duration `yaml:"DialTimeout"`
	ReadTimeout        time.Duration `yaml:"ReadTimeout"`
	WriteTimeout       time.Duration `yaml:"WriteTimeout"`
	PoolSize           int           `yaml:"PoolSize"`
	MinIdleConns       int           `yaml:"MinIdleConns"`
	MaxConnAge         time.Duration `yaml:"MaxConnAge"`
	PoolTimeout        time.Duration `yaml:"PoolTimeout"`
	IdleTimeout        time.Duration `yaml:"IdleTimeout"`
	IdleCheckFrequency time.Duration `yaml:"IdleCheckFrequency"`
	TLSConfig          *tls.Config   `yaml:"TLSConfig"`
	Expiration         time.Duration `yaml:"Expiration"`
}

func (c *Config) ConvertNanosecondToSeconds() {
	c.Expiration = c.Expiration * time.Second
}

func (c *Config) Default() {
	c.Addr = "localhost:6379"
	c.DB = 0
	c.Expiration = 60
}
