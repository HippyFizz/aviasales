package config

type Config struct {
	AuthSecret []byte `yaml:"AuthSecret"`
	Port       uint   `yaml:"Port"`
}

func (c *Config) Default() {
	c.AuthSecret = []byte("secret")
	c.Port = 1323
}
