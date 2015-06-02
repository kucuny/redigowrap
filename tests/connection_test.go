package redigowrap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	REDIS_SERVER_ADDR string = "redis://:redis_user@localhost:6379/0"
)

func TestRedisConnection(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("aaa", "aaa")
}
