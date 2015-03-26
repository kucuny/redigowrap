package redis

import (
	rg "github.com/garyburd/redigo/redis"
)

type connection struct {
	c rg.Conn
	p *rg.Pool
}

type RedisCommands interface {
	// Redigo Function
	// Do(cmd string, args ...interface{}) (interface{}, error)

	// Cluster

	// Connection
	// Auth(password string) (string, error)
	Echo(message string) (string, error)
	Ping() (string, error)
	// Quit() (string, error)
	Select(index int) (string, error)
	Quit() (string, error)

	// Hashes

	// HyperLogLog

	// Keys
	// Del(keys ...string) (int, error)
	// Dump(key string) (int, error)
	// Exists(key string) (int, error)
	// Expire(key string, second int) (string, error)

	// Lists

	// Pub/Sub

	// Scripting

	// Server

	// Sets

	// Sorted Sets

	// Strings

	// Transactions
}
