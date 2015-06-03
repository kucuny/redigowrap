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

func (con *connection) Del(key string) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return rg.Bool(c.Del(key))
	} else {
		return rg.Bool(con.c.Do("DEL", key))
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
