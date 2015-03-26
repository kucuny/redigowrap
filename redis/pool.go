package redis

import (
	rg "github.com/garyburd/redigo/redis"
	"github.com/kucuny/redigocon"
	"time"
)

type ConnectionPoolConfig struct {
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var DefaultConnectionPoolConfig = ConnectionPoolConfig{
	MaxIdle:     60,
	MaxActive:   100,
	IdleTimeout: 30,
}

func CreatePool(serverAddr, auth, db string, config ConnectionPoolConfig) (Connection, error) {
	dialer := func() (rg.Conn, error) {
		return redigocon.Connect(serverAddr, auth, db)
	}

	// config := getConnectionPoolConfig(poolConfig)

	pool := rg.NewPool(dialer, config.MaxIdle)
	pool.MaxActive = config.MaxActive
	pool.IdleTimeout = config.IdleTimeout

	con := &connection{p: pool}

	return con, nil
}

func CreatePoolUri(uri string, config ConnectionPoolConfig) (Connection, error) {
	dialer := func() (rg.Conn, error) {
		return redigocon.ConnectUrl(uri)
	}

	// config := getConnectionPoolConfig(poolConfig)

	pool := rg.NewPool(dialer, config.MaxIdle)
	pool.MaxActive = config.MaxActive
	pool.IdleTimeout = config.IdleTimeout

	con := &connection{p: pool}

	return con, nil
}

func (con *connection) GetConnection() (Connection, error) {
	c := con.p.Get()

	resCon := &connection{c: c}

	return resCon, nil
}

func (con *connection) ActiveCount() int {
	return con.p.ActiveCount()
}

func (con *connection) Release() {
	con.c.Close()
}
