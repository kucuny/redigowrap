package redis

import (
	"fmt"
	"git.cdnetworks.com/metric/redigocon"
)

type Connection interface {
	RedisCommands
}

func CreateConnection(serverAddr, auth, db string) (Connection, error) {
	c, e := redigocon.Connect(serverAddr, auth, db)

	if e != nil {
		fmt.Println(e)

		return nil, e
	}

	con := &connection{c: c}

	return con, nil
}

func CreateConnectionUri(uri string) (Connection, error) {
	c, e := redigocon.ConnectUrl(uri)

	if e != nil {
		fmt.Println(e)

		return nil, e
	}

	con := &connection{c: c}

	return con, nil
}
