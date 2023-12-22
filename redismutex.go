package xtools

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisMutex struct {
	rdb      *redis.Client
	identity string
	lockname string
}

func NewRedisMutex(identity, lockname string) *RedisMutex {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "abc123", // no password set
		DB:       0,        // use default DB
	})
	return &RedisMutex{rdb, identity, lockname}
}

func (m *RedisMutex) TryLock() bool {
	ctx := context.Background()
	ok, err := m.rdb.SetNX(ctx, m.lockname, m.identity, time.Duration(3*time.Second)).Result()
	if err != nil {
		return false
	}
	return ok
}

func (m *RedisMutex) Unlock() bool {
	ctx := context.Background()
	identity, err := m.rdb.Get(ctx, m.lockname).Result()
	if err != nil {
		return false
	}
	if identity != m.identity {
		return false
	}
	_, err = m.rdb.Del(ctx, m.lockname).Result()
	return err == nil
}
