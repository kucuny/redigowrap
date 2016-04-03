package redis

import (
	"fmt"
	"github.com/kucuny/redigowrap/cmd"
	rg "github.com/garyburd/redigo/redis"
)

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

type RedigoWrap struct {
	cmd.RedisCommandConnection
	con rg.Conn
}

func NewRedigoWrap(config RedisConnConfig) (*RedigoWrap, error) {
	var err error

	con, err := newConn(config)

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

func newConn(config RedisConnConfig) (rg.Conn, error) {
	address := fmt.Sprintf("redis://%s:%d/%d", config.host, config.port, config.db)
	return rg.DialURL(address)
}
