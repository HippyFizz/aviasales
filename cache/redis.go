package cache

import (
	"aviasales/cache/config"
	"fmt"
	"github.com/go-redis/redis"
)

var client *redis.Client

func newClient(config *config.Config) error {
	client = redis.NewClient(&redis.Options{
		Network:            config.Network,
		Addr:               config.Addr,
		Password:           config.Password,
		DB:                 config.DB,
		MaxRetries:         config.MaxRetries,
		MinRetryBackoff:    config.MinRetryBackoff,
		MaxRetryBackoff:    config.MaxRetryBackoff,
		DialTimeout:        config.DialTimeout,
		ReadTimeout:        config.ReadTimeout,
		WriteTimeout:       config.WriteTimeout,
		PoolSize:           config.PoolSize,
		MinIdleConns:       config.MinIdleConns,
		MaxConnAge:         config.MaxConnAge,
		PoolTimeout:        config.PoolTimeout,
		IdleTimeout:        config.IdleTimeout,
		IdleCheckFrequency: config.IdleCheckFrequency,
		TLSConfig:          config.TLSConfig,
	})

	pong, err := client.Ping().Result()

	if err != nil {
		return err
	}
	fmt.Println(pong, err)
	return nil
}

func Manager(config *config.Config) *redis.Client {
	config.ConvertNanosecondToSeconds()
	if err := newClient(config); err != nil {
		panic(err)
	}
	return client
}
