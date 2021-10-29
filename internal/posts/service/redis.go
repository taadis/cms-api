package service

import "github.com/go-redis/redis/v8"

func getRdb() *redis.Client {
	opts := &redis.Options{}
	opts.Addr = "localhost:6379"
	opts.Password = "" // no password
	opts.DB = 0        // use default DB
	rdb := redis.NewClient(opts)
	return rdb
}
