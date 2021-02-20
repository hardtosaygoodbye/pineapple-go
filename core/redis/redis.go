package redis

import (
	"context"
	"fmt"
	"pineapple-go/config"
	"pineapple-go/core/log"
	"runtime"

	"github.com/go-redis/redis/v8"
)

// Client client instance to use redis
var Client *redis.Client

//ConnectRedis connect to redis
func ConnectRedis() {
	redis.SetLogger(&redisLogger{})
	Client = redis.NewClient(&redis.Options{
		Addr: config.Redis.Addr,
		DB:   config.Redis.DB, // use default DB
	})
	ctx := context.Background()
	_, err := Client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	log.InitLogger.Info("connnect to redis successful")
}

type redisLogger struct {
}

func (l *redisLogger) Printf(ctx context.Context, format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(3)
	message := fmt.Sprintf(format, v...)
	m := map[string]interface{}{
		"file":    fmt.Sprintf("%s:%d", file, line),
		"message": message,
	}
	log.Warn(ctx, "redis", m)
}
