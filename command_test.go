package redis

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

const (
	REDIS_SERVER_ADDR string = "redis://:passwd@localhost:6379/1"
)

type RedisCommandTestSuite struct {
	suite.Suite
	con Connection
}

func TestRedisCommandTestSuite(t *testing.T) {
	suite.Run(t, new(RedisCommandTestSuite))
}

func (suite *RedisCommandTestSuite) SetupTest() {
	assert := assert.New(suite.T())

	var err error
	suite.con, err = CreateConnectionUri(REDIS_SERVER_ADDR)

	assert.NoError(err)

	suite.con.FlushDB()
}

func (suite *RedisCommandTestSuite) TearDownTest() {
	suite.con.FlushDB()

	suite.con.Close()
}

func (suite *RedisCommandTestSuite) TestSortedSetCommands() {
	assert := assert.New(suite.T())

	isSuccess, err := suite.con.FlushDB()

	assert.NoError(err)
	assert.True(isSuccess)

	testKey := "test"
	testData := map[float64]string{
		1: "test1",
		2: "test2",
		3: "test3",
		4: "test4",
		5: "test5",
		6: "test6",
	}

	addResult, err := suite.con.ZAdd(testKey, testData)

	assert.NoError(err)
	assert.NotNil(addResult)

	zcardRes, err := suite.con.ZCard(testKey)

	assert.NoError(err)
	assert.Equal(6, zcardRes)

	zcountRes, err := suite.con.ZCount(testKey, "1", "2")

	assert.NoError(err)
	assert.Equal(2, zcountRes)

	zcountRes, err = suite.con.ZCount(testKey, "(1", "2")

	assert.NoError(err)
	assert.Equal(1, zcountRes)

	getResult, err := suite.con.ZRangeWithScores(testKey, 0, 5)

	assert.NoError(err)
	assert.Equal("test1", getResult[1])
	assert.Equal("test2", getResult[2])

	delResult, err := suite.con.ZRem(testKey, []string{"test1"})

	assert.NoError(err)
	assert.Equal(1, delResult)

	getResult, err = suite.con.ZRangeWithScores(testKey, 0, 5)

	assert.NoError(err)
	assert.Empty(getResult[1])
	assert.Equal("test2", getResult[2])

	delRangeResult, err := suite.con.ZRemRangeByScore(testKey, "(2", "3")

	assert.NoError(err)
	assert.Equal(1, delRangeResult)

	getResult, err = suite.con.ZRangeWithScores(testKey, 0, 5)

	assert.NoError(err)
	assert.NotEmpty(getResult[2])
	assert.Empty(getResult[3])
}
