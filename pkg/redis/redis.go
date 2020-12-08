package redis

import (
	"github.com/gomodule/redigo/redis"
	"goblog/pkg/config"
	"time"
)

// 连接池的连接
var ConnPool *redis.Pool

// ConnectRedisPool 连接池
func ConnectRedisPool() {
	ConnPool = &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			// 连接redis客户端
			if conn, err = redis.Dial("tcp", config.GetString("database.redis.host")+":"+config.GetString("database.redis.port")); err != nil {
				return nil, err
			}
			// 使用密码
			if auth := config.GetString("database.redis.auth"); auth != "" {
				// 使用密码登录
				if _, err := conn.Do("AUTH", auth); err != nil {
					conn.Close()
					return nil, err
				}
			}
			// 连接指定库
			if _, err := conn.Do("SELECT", config.GetInt("database.redis.database")); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, err
		},
		TestOnBorrow:    nil,
		MaxIdle:         1,
		MaxActive:       10,
		IdleTimeout:     180 * time.Second,
		Wait:            true,
		MaxConnLifetime: 0,
	}
}
