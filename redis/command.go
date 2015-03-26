package redis

import (
	rg "github.com/garyburd/redigo/redis"
)

func (con *connection) Do(command string, args ...interface{}) (interface{}, error) {
	return con.c.Do(command, args...)
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
