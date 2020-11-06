package redis_adapter

import (
	"fmt"
	"time"

	"github.com/dollarkillerx/chronos/adapter"
	"github.com/dollarkillerx/chronos/utils"
	"github.com/gomodule/redigo/redis"
)

type RedisAdapter struct {
	redisPool *redis.Pool
}

func New(uri string, auth ...string) adapter.Adapter {
	output := &RedisAdapter{}
	var passwd string
	if len(auth) == 1 {
		passwd = auth[0]
	}
	output.redisPool = newRedisPool(uri, passwd)
	return output
}

func (r *RedisAdapter) AddRule(rules ...string) error {
	if len(rules) <= 1 {
		return fmt.Errorf("Rule less than 1")
	}
	var key []string
	key = append(rules[:len(rules)-1])
	val := rules[len(rules)-1]

	redis := r.redisPool.Get()
	defer redis.Close()

	_, err := redis.Do("set", utils.Combinations(key, ","), val)
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisAdapter) RemoveRule(rules ...string) error {
	if len(rules) <= 1 {
		return fmt.Errorf("Rule less than 1")
	}
	var key []string
	key = append(rules[:len(rules)-1])
	val := rules[len(rules)-1]

	redis := r.redisPool.Get()
	defer redis.Close()

	_, err := redis.Do("del", utils.Combinations(key, ","), val)
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisAdapter) GetRule(rules ...string) (rule string, err error) {
	if len(rules) <= 1 {
		return "", fmt.Errorf("Rule less than 1")
	}

	rd := r.redisPool.Get()
	defer rd.Close()

	return redis.String(rd.Do("get", utils.Combinations(rules, ",")))
}

// newRedisPool:创建redis连接池
func newRedisPool(host string, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     50,                // 池中的最大空闲连接数
		MaxActive:   30,                // 最大连接数
		IdleTimeout: 300 * time.Second, // 超时回收
		Dial: func() (conn redis.Conn, e error) {
			// 1. 打开连接
			dial, e := redis.Dial("tcp", host)
			if e != nil {
				fmt.Println(e.Error())
				return nil, e
			}
			// 2. 访问认证
			if password != "" {
				dial.Do("AUTH", password)
			}
			return dial, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error { // 定时检查连接是否可用
			// time.Since(t) 获取离现在过了多少时间
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}
