package database


import (
	"log"
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/AndreMicheletti/matchmaker-server-go/internal/config"
)

var ctx = context.Background()

type Redis struct {
	client *redis.Client
}

func NewRedis(cfg config.Config) Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.REDIS_HOST,
		// Username: cfg.REDIS_USER,
		// Password: cfg.REDIS_PASS,
		Password: "",
		DB: 0,
	})

	redis := Redis{client: rdb}

	// You can also check the error directly
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Ping failed: %v", err)
	}
	log.Println("[REDIS] Successfully connected to Redis!")

	return redis
}

func (r *Redis) Client() *redis.Client {
	return r.client
}


func (r *Redis) Close() {
	r.client.Close()
}
