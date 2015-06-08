package main

import (
	"fmt"
	"git.cdnetworks.com/metric/redigowrap/redis"
)

var redisServerAddr string = "redis://:redis_user@localhost:6379/0"

func main() {
	// Simple Connection
	con, _ := redis.CreateConnectionUri(redisServerAddr)
	defer con.Quit()
	fmt.Println(con.Echo("TESTEST"))
	fmt.Println(con.Ping())
	fmt.Println(con.Select(1))

	fmt.Println("MGET")
	fmt.Println(con.MGet([]string{"test", "key2"}))

	con.HSet("test:test1", "test1", "12341")
	con.HSet("test:test1", "test2", "12342")
	con.HSet("test:test1", "test3", "12343")
	fmt.Println(con.HMGet("testtest", []string{"test1", "test2", "test3"}))

	keyValue := map[string]string{
		"t1": "123",
		"t2": "456",
		"t3": "789",
	}
	con.HMSet("test:test2", keyValue)
	fmt.Println(con.HGetAll("test:test2"))

	// Pooled Connection
	// var config = redis.ConnectionPoolConfig{
	// 	MaxIdle:     60,
	// 	MaxActive:   100,
	// 	IdleTimeout: 30,
	// }

	// pool, _ := redis.CreatePoolUri(redisServerAddr, config)
	// fmt.Println(pool.Echo("TESTEST"))
	// fmt.Println(pool.ActiveCount())
	// fmt.Println(pool.Ping())
	// fmt.Println(pool.ActiveCount())
	// fmt.Println(pool.Do("GET", 1))

	// pool.PoolClose()
}
