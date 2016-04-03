package redis

import (
	rg "github.com/garyburd/redigo/redis"
)

type RedigoWrapPubSub struct {
	pubsubCon *rg.PubSubConn
}
