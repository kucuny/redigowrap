package redis

import (
	"time"
	rg "github.com/garyburd/redigo/redis"
)

var DefaultRedisPoolConfig = RedisPoolConfig{
	maxIdle:     60,
	maxActive:   100,
	idleTimeout: 30,
}

type RedisPoolConfig struct {
	RedisConnConfig
	maxIdle     int
	maxActive   int
	idleTimeout time.Duration
}

type RedigoWrapPool struct {
	pool *rg.Pool
}

func NewRedisPoolConfig(config RedisConnConfig, maxIdle, maxActive int, idleTimeout time.Duration) RedisPoolConfig {
	return RedisPoolConfig{
		RedisConnConfig: config,
		maxIdle:         maxIdle,
		maxActive:       maxActive,
		idleTimeout:     idleTimeout,
	}
}

func NewRedigoWrapPool(poolConfig RedisPoolConfig) (*RedigoWrapPool, error) {
	dialer := func() (rg.Conn, error) {
		return newConn(poolConfig.RedisConnConfig)
	}

	tester := func(c rg.Conn, t time.Time) error {
		_, err := c.Do("PING")
		return err
	}

	config := getConnectionPoolConfig(poolConfig)

	pool := &rg.Pool{
		MaxIdle:      config.maxIdle,
		IdleTimeout:  config.idleTimeout,
		Dial:         dialer,
		TestOnBorrow: tester,
	}

	return &RedigoWrapPool{
		pool: pool,
	}, nil
}

func getConnectionPoolConfig(config RedisPoolConfig) RedisPoolConfig {
	if config.idleTimeout == 0 || config.maxIdle == 0 || config.maxActive == 0 {
		return DefaultRedisPoolConfig
	} else {
		return config
	}
}
