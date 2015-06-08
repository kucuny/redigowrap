package redigowrap

import (
	"git.cdnetworks.com/metric/redigowrap"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	REDIS_SERVER_ADDR string = "redis://:redis_user@localhost:6379/0"
)

func TestRedisConnection(t *testing.T) {
	assert := assert.New(t)

	var config = redis.ConnectionPoolConfig{
		MaxIdle:     60,
		MaxActive:   100,
		IdleTimeout: 30,
	}

	con, _ := redis.CreatePoolUri(REDIS_SERVER_ADDR, config)
	res, _ := con.Echo("aaa")
	assert.Equal("aaa", res)
}
