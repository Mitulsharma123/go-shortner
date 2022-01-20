package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var (						// declaration for storeservice and redis contect
	storeService = &StorageService{}
	ctx          = context.Background()
)

const CacheDuration = 6 * time.Hour		//settting the cache duration 

type StorageService struct {			// wrapper struct around redis client 
	redisClient *redis.Client
}

func InitializeStore() *StorageService {		//store service initilization 
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}



