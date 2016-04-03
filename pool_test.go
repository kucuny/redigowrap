package redis

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type RedigoWrapPoolTest struct {
	suite.Suite
}

func (suite *RedigoWrapPoolTest) SetupSuite() {
}

func (suite *RedigoWrapPoolTest) TestRedisPoolConnection() {
	connConfig := NewRedisConnConfig("localhost", 6379, "", 0)
	poolConfig := NewRedisPoolConfig(connConfig, 0, 0, 0)

	pool, err := NewRedigoWrapPool(poolConfig)

	suite.Nil(err)
	suite.NotNil(pool)
}

func TestRedigoWrapPoolConnectionTestSuite(t *testing.T) {
	suite.Run(t, new(RedigoWrapPoolTest))
}