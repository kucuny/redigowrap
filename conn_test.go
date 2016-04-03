package redis

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type RedigoWrapConnectionTest struct {
	suite.Suite
}

func (suite *RedigoWrapConnectionTest) SetupSuite() {
}

func (suite *RedigoWrapConnectionTest) TestRedisConnection() {
	connConfig := NewRedisConnConfig("localhost", 6379, "", 0)
	redis, err := NewRedigoWrap(connConfig)

	suite.Nil(err)
	suite.NotNil(redis)
}

func TestRedigoWrapConnectionTestSuite(t *testing.T) {
	suite.Run(t, new(RedigoWrapConnectionTest))
}