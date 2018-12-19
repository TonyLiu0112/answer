package activity

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

var (
	closeChan chan int
	rds       rdsInstance
)

type RedisInstance struct {
	maxIdle     int
	idleTimeout time.Duration
	server      string
	password    string
	_pool       *redis.Pool
}

type rdsInstance RedisInstance

func init() {
	log.Println("do init redis.")
	rds = rdsInstance{5, 240 * time.Second, "www.tony666.com:6379", "123456", nil}
	rds.newPool()
	go rds.closeListener()
}

func (instance *rdsInstance) newPool() {
	instance._pool = &redis.Pool{
		MaxIdle:     5,
		IdleTimeout: 240 * time.Second,
		Dial: func() (conn redis.Conn, e error) {
			conn, err := redis.Dial("tcp", instance.server, redis.DialPassword(instance.password))
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
}

func (instance *rdsInstance) getConn() redis.Conn {
	return rds._pool.Get()
}

func (instance *rdsInstance) get(key string) string {
	value, _ := redis.String(rds.getConn().Do("get", key))
	return value
}

func (instance *rdsInstance) set(key string, value string) bool {
	_, e := redis.String(rds.getConn().Do("set", key, value))
	return e == nil
}

func (instance *rdsInstance) exists(key string) bool {
	exists, _ := redis.Bool(rds.getConn().Do("exists", key))
	return exists
}

func (instance *rdsInstance) expire(key string, millisecond int64) bool {
	n, _ := rds.getConn().Do("expire", millisecond)
	return n == int64(1)
}

func (instance *rdsInstance) incr(key string) bool {
	_, err := rds.getConn().Do("incr", key)
	return err == nil
}

func (instance *rdsInstance) closeListener() {
	closeChan := make(chan int)
	for {
		select {
		case _, opened := <-closeChan:
			if !opened {
				rds._pool.Close()
			}
		}
	}
}
