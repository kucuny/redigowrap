package redis

import (
	rg "github.com/garyburd/redigo/redis"
)

type (
	connection struct {
		c rg.Conn
		p *rg.Pool
	}

	RedisCommands interface {
		// Redigo Function
		Do(cmd string, args ...interface{}) (interface{}, error)

		// Cluster

		// Connection
		// Auth(password string) (string, error)
		Echo(message string) (string, error)
		Ping() (string, error)
		Select(index int) (string, error)
		Quit() (string, error)

		// Hashes

		// HyperLogLog

		// Keys
		Del(key string) (bool, error)
		// Dump(key string) (int, error)
		Exists(key string) (bool, error)
		// Expire(key string, second int) (int, error)
		// Expireat(key string, timestamp int64) (int, error)
		// Keys(pattern string) ([]string, error)
		// // Migrate(host, port, key, destDB, timeout) (int)
		// Move(key, db string) (int, error)
		// // Object(subcommand string, args []string) (error)
		// Persist(key string) (int, error)
		// PExpire(key string, millisec int64) (int, error)
		// PExpireat(key string, millisec int64) (int, error)
		// PTTL(key string) (int, error)
		// RandomKey() (string, error)
		// Rename(key, newKey string) (int, error)
		// RenameNX(key, newKey string) (int, error)
		// Restore(key string, ttl int, serializedValue string) (int, error)
		// Scan(scanValue string) (map[string][]string, error)
		// // Sort(key string)
		// TTL(key string) (int, error)
		// Type(key string) (string, error)
		// Wait(numSlaves, timeout int) (int, error)

		// Lists

		// Pub/Sub

		// Scripting

		// Server

		// Sets

		// Sorted Sets

		// Strings
		Append(key, value string) (int, error)
		BitCount(key string) (int, error)
		BitCountRange(key string, start, end int) (int, error)
		BitOP(operation, destKey string, keys []interface{}) (int, error)
		BitPos(key string, start int) (int, error)
		BitPosRange(key string, start, end int) (int, error)
		Decr(key string) (int, error)
		DecrBy(key string, decrement int) (int, error)
		Get(key string) (string, error)
		GetBit(key string, offset int) (int, error)
		GetRange(key string, start, end int) (string, error)
		GetSet(key, value string) (string, error)
		Incr(key string) (int, error)
		IncrBy(key string, increment int) (int, error)
		IncrByFloat(key string, increment float64) (float64, error)
		MGet(keys []string) ([]string, error)
		MSet(keyValue map[string]string) (bool, error)
		MSetNX(keyValue map[string]string) (int, error)
		PSetEX(key, value string, millisec int64) (bool, error)
		Set(key, value string) (bool, error)
		SetBit(key, value string, offset int) (int, error)
		SetEX(key, value string, seconds int) (bool, error)
		SetNX(key, value string) (int, error)
		SetRange(key, value string, offset int) (int, error)
		StrLen(key string) (int, error)

		// Transactions
	}

	PoolCommands interface {
		GetConnection() (PoolConnection, error)
		ActiveCount() int
		Release()
		PoolClose()
	}
)
