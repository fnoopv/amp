package redis

import (
	"github.com/redis/rueidis"
	"goyave.dev/goyave/v5/util/errors"
)

// Client redis客户端
var Client rueidis.Client

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
