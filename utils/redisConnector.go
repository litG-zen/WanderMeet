package utils

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	redisClient *RedisClient
	redisOnce   sync.Once
	REDIS_URL   = os.Getenv("REDIS_URL")
)

type RedisClient struct {
	client *redis.Client
	ctx    context.Context
	mutex  sync.RWMutex
}

func (r *RedisClient) initialize() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	fmt.Printf("REDIS_URL :%v", REDIS_URL)
	if REDIS_URL == "" {
		return fmt.Errorf("REDIS URL is not defined")
	}

	opt, err := redis.ParseURL(REDIS_URL)
	if err != nil {
		return fmt.Errorf("Failed to parse REDIS URL : %v", REDIS_URL)

	}

	r.client = redis.NewClient(opt)

	// Test connection
	_, err = r.client.Ping(r.ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to redis instance %v", err)
	}

	return nil
}

func NewRedisClient() (*RedisClient, error) {
	var initErr error
	redisOnce.Do(func() {
		redisClient = &RedisClient{ctx: context.Background()}
		initErr = redisClient.initialize()
	})

	if initErr != nil {
		return nil, initErr
	}

	fmt.Printf("\n****Redis connection Initiated****\n")

	return redisClient, nil
}

func GetRedisInstance() (*RedisClient, error) {
	if redisClient == nil {
		return NewRedisClient()
	}
	return redisClient, nil
}

func CloseRedisConnection() {
	redisClient.client.Close()
	fmt.Printf("\n****Redis connection Closed****\n")

}

func AddUserLocation(lat float32, long float32, key string) error {
	//rds := GetRedisInstance()
	//rds.client.
	return nil
}
