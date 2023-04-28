package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type Cache interface {
	Has(string) (bool, error)
	Get(string) (any, error)
	Set(string, any, ...int) error
	Forget(string) error
	EmptyByMatch(string) error
	Empty() error
}

type RedisCache struct {
	Conn   *redis.Pool
	Prefix string
}

type Entry map[string]any

func (c *RedisCache) Has(str string) (bool, error) {
	key := fmt.Sprintf("%s:%s", c.Prefix, str)
	conn := c.Conn.Get()
	defer conn.Close()

	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false, err
	}

	return ok, nil
}

func (c *RedisCache) Get(str string) (any, error) {
	return "", nil
}

func (c *RedisCache) Set(str string, data any, ttl ...int) error {
	return nil
}

func (c *RedisCache) Forget(str string) error {
	return nil
}

func (c *RedisCache) EmptyByMatch(str string) error {
	return nil
}

func (c *RedisCache) Empty() error {
	return nil
}
