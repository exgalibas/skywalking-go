package redis

import "github.com/gomodule/redigo/redis"

//skywalking:native github.com/gomodule/redigo/redis poolConn
type nativePoolConn struct {
	c redis.Conn
}

//skywalking:native github.com/gomodule/redigo/redis activeConn
type nativeActiveConn struct {
	pc *nativePoolConn
}

//skywalking:native github.com/gomodule/redigo/redis errorConn
type nativeErrConn struct{}
