package main

import (
	"fmt"
	"git.cdnetworks.com/metric/redigowrap"
)

const (
	REDIS_SERVER_ADDR string = "redis://:passwd@localhost:6379/1"
)

func main() {
	con, _ := redis.CreateConnectionUri(REDIS_SERVER_ADDR)

	con.FlushDB()

	testKey := "test"
	testData := map[float64]string{
		1: "test 1",
		2: "test 2",
		3: "test TearDownSuite3",
	}

	res, err := con.ZAdd(testKey, testData)

	fmt.Println(err)
	fmt.Println(res)

	result, err := con.ZRangeWithScores(testKey, 0, 5)

	fmt.Println(err)
	fmt.Println(result)
}
