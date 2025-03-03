package redis

import (
	"reflect"

	"github.com/redis/rueidis"
	"goyave.dev/goyave/v5/config"
	"goyave.dev/goyave/v5/util/errors"
)

// Client redis客户端
var Client rueidis.Client

// 注册配置项
func init() {
	config.Register("redis.host", config.Entry{
		Value:    "127.0.0.1",
		Type:     reflect.String,
		IsSlice:  false,
		Required: true,
	},
	)
	config.Register("redis.port", config.Entry{
		Value:    6379,
		Type:     reflect.Int,
		IsSlice:  false,
		Required: true,
	},
	)
}

// Initialize 初始化redis连接
func Initialize(address string) error {
	client, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{address},
	})
	if err != nil {
		return errors.New(err)
	}

	Client = client

	return nil
}
