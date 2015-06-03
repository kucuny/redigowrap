package main

import (
	"fmt"
	"github.com/kucuny/redigowrap/redis"
)

var redisServerAddr string = "redis://:redis_user@localhost:6379/0"

func main() {
	// Simple Connection
	con, _ := redis.CreateConnectionUri(redisServerAddr)
	fmt.Println(con.Echo("TESTEST"))
	fmt.Println(con.Ping())
	fmt.Println(con.Select(1))
	res, err := con.Quit()
	fmt.Println(res, err)

	fmt.Println("MGET")
	fmt.Println(con.MGet("test", "key2"))

	// Pooled Connection
	var config = redis.ConnectionPoolConfig{
		MaxIdle:     60,
		MaxActive:   100,
		IdleTimeout: 30,
	}

	pool, _ := redis.CreatePoolUri(redisServerAddr, config)
	fmt.Println(pool.Echo("TESTEST"))
	fmt.Println(pool.ActiveCount())
	fmt.Println(pool.Ping())
	fmt.Println(pool.ActiveCount())
	fmt.Println(pool.Do("GET", 1))

	pool.PoolClose()
}
