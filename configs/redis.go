package configs

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func (c *Config) LoadRedisClient() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", c.Redis.Host, c.Redis.Port),
		Password: c.Redis.Password,
		DB:       0,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
