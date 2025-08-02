package utilities

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient  *redis.Client
	RedisContext context.Context
)

func ConnectToRedis() {
	RedisContext = context.Background()
	opt, _ := redis.ParseURL(os.Getenv("REDIS_URL"))
	RedisClient = redis.NewClient(opt)
	fmt.Println("Redis client connected", RedisClient)

	// RedisClient.Set(RedisContext, "foo", "bar", 0)
	// val := RedisClient.Get(RedisContext, "foo").Val()
	// print(val)
}
