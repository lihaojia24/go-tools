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
	ticker   *time.Ticker
}

func NewRedisMutex(identity, lockname string) *RedisMutex {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "abc123",
		DB:       0,
	})
	return &RedisMutex{rdb, identity, lockname, nil}
}

func (m *RedisMutex) TryLock() bool {
	ctx := context.Background()
	ok, err := m.rdb.SetNX(ctx, m.lockname, m.identity, time.Duration(3*time.Second)).Result()
	if err != nil {
		return false
	}
	m.ticker = time.NewTicker(time.Duration(time.Second * 2))
	go func() {
		for {
			_, ok := <-m.ticker.C
			if ok {
				m.rdb.SetEx(ctx, m.lockname, m.identity, time.Duration(3*time.Second))
			} else {
				break
			}
		}
	}()
	return ok
}

func (m *RedisMutex) Lock() {
	ctx := context.Background()
	for {
		ok, err := m.rdb.SetNX(ctx, m.lockname, m.identity, time.Duration(3*time.Second)).Result()
		if err == nil && ok {
			break
		}
		time.Sleep(time.Second)
	}
	m.ticker = time.NewTicker(time.Duration(time.Second * 2))
	go func() {
		for {
			_, ok := <-m.ticker.C
			if ok {
				m.rdb.SetEx(ctx, m.lockname, m.identity, time.Duration(3*time.Second))
			} else {
				break
			}
		}
	}()
}

func (m *RedisMutex) Unlock() bool {
	ctx := context.Background()
	identity, err := m.rdb.Get(ctx, m.lockname).Result()
	if err != nil {
		return false
	}
	if m.ticker != nil {
		m.ticker.Stop()
	}
	if identity != m.identity {
		return false
	}
	_, err = m.rdb.Del(ctx, m.lockname).Result()
	return err == nil
}
