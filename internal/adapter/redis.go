package adapter

import (
	"context"
	"fmt"

	redis "github.com/go-redis/redis/v8"
	"github.com/porter-dev/porter/internal/config"
)

// NewRedisClient returns a new redis client instance
func NewRedisClient(conf *config.RedisConf) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		Username: conf.Username,
		Password: conf.Password,
		DB:       conf.DB,
	})

	_, err := client.Ping(context.Background()).Result()
	return client, err
}
