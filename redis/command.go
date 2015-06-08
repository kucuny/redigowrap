package redis

import (
	rg "github.com/garyburd/redigo/redis"
)

func (con *connection) Do(command string, args ...interface{}) (interface{}, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.Do(command, args...)
	} else {
		return con.c.Do(command, args...)
	}
}

/*
	Connection
*/
func (con *connection) Auth(password string) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Bool(c.Auth(password))
	} else {
		res, err := rg.String(con.c.Do("AUTH", password))
		return getBool(res), err
	}
}

func (con *connection) Echo(message string) (string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.String(c.Echo(message))
	} else {
		return rg.String(con.c.Do("ECHO", message))
	}
}

func (con *connection) Ping() (string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.String(c.Ping())
	} else {
		return rg.String(con.c.Do("PING"))
	}
}

func (con *connection) Select(index int) (string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.String(c.Select(index))
	} else {
		return rg.String(con.c.Do("SELECT", index))
	}
}

func (con *connection) Quit() (string, error) {
	return rg.String(con.c.Do("QUIT"))
}

/*
	Hashes
*/
func (con *connection) HDel(hashKey string, fields []string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Int(c.HDel(hashKey, fields))
	} else {
		req := make([]interface{}, len(fields)+1)
		req[0] = hashKey
		for idx, val := range fields {
			req[idx+1] = val
		}

		return rg.Int(con.c.Do("HDEL", req...))
	}
}

func (con *connection) HExists(hashKey, field string) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Bool(c.HExists(hashKey, field))
	} else {
		return rg.Bool(con.c.Do("HEXISTS", hashKey, field))
	}
}

func (con *connection) HGet(hashKey, field string) (string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.String(c.HGet(hashKey, field))
	} else {
		return rg.String(con.c.Do("HGET", hashKey, field))
	}
}

func (con *connection) HGetAll(hashKey string) (map[string]string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.StringMap(c.HGetAll(hashKey))
	} else {
		return rg.StringMap(con.c.Do("HGETALL", hashKey))
	}
}

func (con *connection) HIncrBy(hashKey, field string, increment int) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Int(c.HIncrBy(hashKey, field, increment))
	} else {
		return rg.Int(con.c.Do("HINCRBY", hashKey, field, increment))
	}
}

func (con *connection) HIncrByFloat(hashKey, field string, increment float64) (float64, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Float64(c.HIncrByFloat(hashKey, field, increment))
	} else {
		return rg.Float64(con.c.Do("HINCRBYFLOAT", hashKey, field, increment))
	}
}

func (con *connection) HKeys(hashKey string) ([]string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Strings(c.HKeys(hashKey))
	} else {
		return rg.Strings(con.c.Do("HKEYS", hashKey))
	}
}

func (con *connection) HLen(hashKey string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Int(c.HLen(hashKey))
	} else {
		return rg.Int(con.c.Do("HLEN", hashKey))
	}
}

func (con *connection) HMGet(hashKey string, fields []string) ([]string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Strings(c.HMGet(hashKey, fields))
	} else {
		req := make([]interface{}, len(fields)+1)
		req[0] = hashKey
		for idx, val := range fields {
			req[idx+1] = val
		}
		return rg.Strings(con.c.Do("HMGet", req...))
	}
}

func (con *connection) HMSet(hashKey string, fieldValue map[string]string) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Bool(c.HMSet(hashKey, fieldValue))
	} else {
		req := make([]interface{}, len(fieldValue)*2+1)
		req[0] = hashKey
		idx := 1
		for name, value := range fieldValue {
			req[idx] = name
			idx++
			req[idx] = value
			idx++
		}

		res, err := rg.String(con.c.Do("HMSET", req...))
		isSuccess := getBool(res)

		return isSuccess, err
	}
}

// func HScan() ()

func (con *connection) HSet(hashKey, field, value string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Int(c.HSet(hashKey, field, value))
	} else {
		return rg.Int(con.c.Do("HSET", hashKey, field, value))
	}
}

func (con *connection) HSetNX(hashKey, field, value string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Int(c.HSetNX(hashKey, field, value))
	} else {
		return rg.Int(con.c.Do("HSETNX", hashKey, field, value))
	}
}

func (con *connection) HStrLen(hashKey, field string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Int(c.HStrLen(hashKey, field))
	} else {
		return rg.Int(con.c.Do("HSTRLEN", hashKey, field))
	}
}

func (con *connection) HVals(hashKey string) ([]string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Strings(c.HVals(hashKey))
	} else {
		return rg.Strings(con.c.Do("HVALS", hashKey))
	}
}

/*
	Keys
*/
func (con *connection) Del(keys []string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Int(c.Del(keys))
	} else {
		req := make([]interface{}, len(keys))
		for idx, val := range keys {
			req[idx] = val
		}
		return rg.Int(con.c.Do("DEL", req...))
	}
}

func (con *connection) Dump(key string) (string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.String(c.Exists(key))
	} else {
		return rg.String(con.c.Do("DUMP", key))
	}
}

func (con *connection) Exists(key string) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Bool(c.Exists(key))
	} else {
		return rg.Bool(con.c.Do("EXISTS", key))
	}
}

func (con *connection) Expire(key string, seconds int) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Bool(c.Expire(key, seconds))
	} else {
		return rg.Bool(con.c.Do("EXPIRE", key, seconds))
	}
}

func (con *connection) Expireat(key string, timestamp int64) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Bool(c.Expireat(key, timestamp))
	} else {
		return rg.Bool(con.c.Do("EXPIREAT", key, timestamp))
	}
}

func (con *connection) Keys(pattern string) ([]string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Strings(c.Keys(pattern))
	} else {
		return rg.Strings(con.c.Do("KEYS", pattern))
	}
}

// Migrate()

func (con *connection) Move(key, db string) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()

		res, err := rg.String(c.Move(key, db))
		isSuccess := getBool(res)

		return isSuccess, err
	} else {
		res, err := rg.String(con.c.Do("MOVE", key, db))
		isSuccess := getBool(res)

		return isSuccess, err
	}
}

// Object()

func (con *connection) Persist(key string) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Bool(c.Persist(key))
	} else {
		return rg.Bool(con.c.Do("PERSIST", key))
	}
}

func (con *connection) PExpire(key string, millisec int64) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Bool(c.PExpire(key, millisec))
	} else {
		return rg.Bool(con.c.Do("PEXPIRE", key, millisec))
	}
}

func (con *connection) PExpireat(key string, millisecTimestamp int64) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Bool(c.PExpireat(key, millisecTimestamp))
	} else {
		return rg.Bool(con.c.Do("PEXPIREAT", key, millisecTimestamp))
	}
}

func (con *connection) PTTL(key string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Int(c.PTTL(key))
	} else {
		return rg.Int(con.c.Do("PTTL", key))
	}
}

func (con *connection) RandomKey() (string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.String(c.RandomKey())
	} else {
		return rg.String(con.c.Do("RANDOMKEY"))
	}
}

func (con *connection) Rename(key, newKey string) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()

		res, err := rg.String(c.Rename(key, newKey))
		isSuccess := getBool(res)

		return isSuccess, err
	} else {
		res, err := rg.String(con.c.Do("Rename", key, newKey))
		isSuccess := getBool(res)

		return isSuccess, err
	}
}

func (con *connection) RenameNX(key, newKey string) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()

		res, err := rg.String(c.RenameNX(key, newKey))
		isSuccess := getBool(res)

		return isSuccess, err
	} else {
		res, err := rg.String(con.c.Do("RenameNX", key, newKey))
		isSuccess := getBool(res)

		return isSuccess, err
	}
}

func (con *connection) Restore(key string, ttl int, serializedValue string) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()

		res, err := rg.String(c.Restore(key, ttl, serializedValue))
		isSuccess := getBool(res)

		return isSuccess, err
	} else {
		res, err := rg.String(con.c.Do("RESTORE", key, ttl, serializedValue))
		isSuccess := getBool(res)

		return isSuccess, err
	}
}

func (con *connection) RestoreWithReplace(key string, ttl int, serializedValue, replace string) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()

		res, err := rg.String(c.RestoreWithReplace(key, ttl, serializedValue, replace))
		isSuccess := getBool(res)

		return isSuccess, err
	} else {
		res, err := rg.String(con.c.Do("RESTORE", key, ttl, serializedValue, replace))
		isSuccess := getBool(res)

		return isSuccess, err
	}
}

// Scan()

func (con *connection) Sort(args ...interface{}) ([]string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Strings(c.Sort(args...))
	} else {
		return rg.Strings(con.c.Do("SORT", args...))
	}
}

func (con *connection) TTL(key string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Int(c.TTL(key))
	} else {
		return rg.Int(con.c.Do("TTL", key))
	}
}

func (con *connection) Type(key string) (string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.String(c.Type(key))
	} else {
		return rg.String(con.c.Do("TYPE", key))
	}
}

func (con *connection) Wait(numSlaves, ttl int) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Int(c.Wait(numSlaves, ttl))
	} else {
		return rg.Int(con.c.Do("Wait", numSlaves, ttl))
	}
}

/*
   Strings
*/
func (con *connection) Append(key, value string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Int(c.Append(key, value))
	} else {
		return rg.Int(con.c.Do("APPEND", key, value))
	}
}

func (con *connection) BitCount(key string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Int(c.BitCount(key))
	} else {
		return rg.Int(con.c.Do("BITCOUNT", key))
	}
}

func (con *connection) BitCountRange(key string, start, end int) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Int(c.BitCountRange(key, start, end))
	} else {
		return rg.Int(con.c.Do("BITCOUNTRANGE", key, start, end))
	}
}

func (con *connection) BitOP(key, destKey string, keys []interface{}) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Int(c.BitOP(key, destKey, keys))
	} else {
		return rg.Int(con.c.Do("BITOP", keys...))
	}
}

func (con *connection) BitPos(key string, start int) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Int(c.BitPos(key, start))
	} else {
		return rg.Int(con.c.Do("BITPOS", key, start))
	}
}

func (con *connection) BitPosRange(key string, start, end int) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Int(c.BitPosRange(key, start, end))
	} else {
		return rg.Int(con.c.Do("BITPOS", key, start, end))
	}
}

func (con *connection) Decr(key string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Int(c.Decr(key))
	} else {
		return rg.Int(con.c.Do("DECR", key))
	}
}

func (con *connection) DecrBy(key string, decrement int) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Int(c.DecrBy(key, decrement))
	} else {
		return rg.Int(con.c.Do("DECRBY", key, decrement))
	}
}

func (con *connection) Get(key string) (string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.String(c.Get(key))
	} else {
		return rg.String(con.c.Do("GET", key))
	}
}

func (con *connection) GetBit(key string, offset int) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Int(c.GetBit(key, offset))
	} else {
		return rg.Int(con.c.Do("GETBIT", key, offset))
	}
}

func (con *connection) GetRange(key string, start, end int) (string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.String(c.GetRange(key, start, end))
	} else {
		return rg.String(con.c.Do("GETRANGE", key, start, end))
	}
}

func (con *connection) GetSet(key, value string) (string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.String(c.GetSet(key, value))
	} else {
		return rg.String(con.c.Do("GETSET", key, value))
	}
}

func (con *connection) Incr(key string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Int(c.Incr(key))
	} else {
		return rg.Int(con.c.Do("INCR", key))
	}
}

func (con *connection) IncrBy(key string, increment int) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Int(c.IncrBy(key, increment))
	} else {
		return rg.Int(con.c.Do("INCRBY", key, increment))
	}
}

func (con *connection) IncrByFloat(key string, increment float64) (float64, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Float64(c.IncrByFloat(key, increment))
	} else {
		return rg.Float64(con.c.Do("INCRBYFLOAT", key, increment))
	}
}

func (con *connection) MGet(keys []string) ([]string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Strings(c.MGet(keys))
	} else {
		req := make([]interface{}, len(keys))
		for idx, val := range keys {
			req[idx] = val
		}
		return rg.Strings(con.c.Do("MGET", req...))
	}
}

func (con *connection) MSet(keyValue map[string]string) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()

		res, err := rg.String(c.MSet(keyValue))
		isSuccess := getBool(res)

		return isSuccess, err
	} else {
		req := make([]interface{}, len(keyValue)*2)
		idx := 0
		for name, value := range keyValue {
			req[idx] = name
			idx++
			req[idx] = value
			idx++
		}

		res, err := rg.String(con.c.Do("MSET", req...))
		isSuccess := getBool(res)

		return isSuccess, err
	}
}

func (con *connection) MSetNX(keyValue map[string]string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Int(c.MSetNX(keyValue))
	} else {
		req := make([]interface{}, len(keyValue)*2)
		idx := 0
		for name, value := range keyValue {
			req[idx] = name
			idx++
			req[idx] = value
			idx++
		}

		return rg.Int(con.c.Do("MSETNX", req...))
	}
}

func (con *connection) PSetEX(key, value string, millisec int64) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer con.Release()

		res, err := rg.String(c.PSetEX(key, value, millisec))
		isSuccess := getBool(res)

		return isSuccess, err
	} else {
		res, err := rg.String(con.c.Do("PSETEX", key, millisec, value))
		isSuccess := getBool(res)

		return isSuccess, err
	}
}

func (con *connection) Set(key, value string) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()

		res, err := rg.String(c.Set(key, value))
		isSuccess := getBool(res)

		return isSuccess, err
	} else {
		res, err := rg.String(con.c.Do("SET", key, value))
		isSuccess := getBool(res)

		return isSuccess, err
	}
}

func (con *connection) SetFloat64(key string, value float64) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()

		res, err := rg.String(c.SetFloat64(key, value))
		isSuccess := getBool(res)

		return isSuccess, err
	} else {
		res, err := rg.String(con.c.Do("SET", key, value))
		isSuccess := getBool(res)

		return isSuccess, err
	}
}

func (con *connection) SetBit(key, value string, offset int) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Int(c.SetBit(key, value, offset))
	} else {
		return rg.Int(con.c.Do("SETBIT", key, value, offset))
	}
}

func (con *connection) SetEX(key, value string, seconds int) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()

		res, err := rg.String(c.SetEX(key, value, seconds))
		isSuccess := getBool(res)

		return isSuccess, err
	} else {
		res, err := rg.String(con.c.Do("SETEX", key, seconds, value))
		isSuccess := getBool(res)

		return isSuccess, err
	}
}

func (con *connection) SetNX(key, value string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Int(c.SetNX(key, value))
	} else {
		return rg.Int(con.c.Do("SETNX", key, value))
	}
}

func (con *connection) SetRange(key, value string, offset int) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Int(c.SetRange(key, value, offset))
	} else {
		return rg.Int(con.c.Do("SETRANGE", key, offset, value))
	}
}

func (con *connection) StrLen(key string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Int(c.StrLen(key))
	} else {
		return rg.Int(con.c.Do("STRLEN", key))
	}
}
