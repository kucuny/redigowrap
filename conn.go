package redis

import (
	"fmt"
	rg "github.com/garyburd/redigo/redis"
	"github.com/kucuny/redigocon"
	"time"
)

var DefaultRedisPoolConfig = RedisPoolConfig{
	MaxIdle:     60,
	MaxActive:   100,
	IdleTimeout: 30,
}

type RedisConnConfig struct {
	host string
	port int
	auth string
	db   int
}

func NewRedisConnConfig(host string, port int, auth string, db int) RedisConnConfig {
	return RedisConnConfig{
		host: host,
		port: port,
		auth: auth,
		db:   db,
	}
}

type RedisPoolConfig struct {
	RedisConnConfig
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

func NewRedisPoolConfig(host string, port int, auth string, db int) RedisPoolConfig {

}

type RedigoWrap struct {
	con rg.Conn
}

type RedigoWrapPool struct {
	pool *rg.Pool
}

type RedigoWrapPubSub struct {
	pubsubCon *rg.PubSubConn
}

func NewRedigoWrap(config RedisConnConfig) (*RedigoWrap, error) {
	var err error

	address := fmt.Sprintf("redis://%s:%d/%d", config.host, config.port, config.db)

	fmt.Println(address)
	con, err := rg.DialURL(address)

	if err != nil {
		return nil, err
	}

	if config.auth != "" {
		_, err = con.Do("AUTH", config.auth)
	}

	if err != nil {
		return nil, err
	}

	_, err = con.Do("SELECT", config.db)

	if err != nil {
		return nil, err
	}

	return &RedigoWrap{
		con: con,
	}, nil
}

func NewRedigoWrapPool(uri string, poolConfig RedisPoolConfig) (PoolConnection, error) {
	dialer := func() (rg.Conn, error) {
		return redigocon.ConnectUrl(uri)
	}

	tester := func(c rg.Conn, t time.Time) error {
		_, err := c.Do("PING")
		return err
	}

	config := getConnectionPoolConfig(&poolConfig)

	pool := &rg.Pool{
		MaxIdle:      config.MaxIdle,
		IdleTimeout:  config.IdleTimeout,
		Dial:         dialer,
		TestOnBorrow: tester,
	}

	con := &connection{p: pool}

	return con, nil
}

func getConnectionPoolConfig(config *RedisPoolConfig) RedisPoolConfig {
	if config.IdleTimeout == 0 || config.MaxIdle == 0 || config.MaxActive == 0 {
		return DefaultRedisPoolConfig
	} else {
		return *config
	}
}