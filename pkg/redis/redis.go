package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

var (
	internalClient *Client
	once           sync.Once
	internalCtx    context.Context
)

// Init 初始化redis客户端
func Init(ctx context.Context, addr, password string, db int) *Client {
	once.Do(func() {
		internalCtx = ctx
		internalClient = New(addr, password, db)

	})
	return internalClient
}

// GetClient 获取redis客户端
func GetClient() *Client {
	return internalClient
}

// GetRedisClient 获取redis客户端
func GetRedisClient() *redis.Client {
	return internalClient.cli
}

// New 创建redis客户端实例
func New(addr, password string, db int) *Client {
	cli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	cmd := cli.Ping(internalCtx)
	if err := cmd.Err(); err != nil {
		panic(err)
	}

	return &Client{
		cli,
	}
}

// Client redis客户端
type Client struct {
	cli *redis.Client
}

// GetRedisClient 获取redis客户端
func (a *Client) GetRedisClient() *redis.Client {
	return a.cli
}

// Set 设定值
func (a *Client) Set(key, value string, expiration time.Duration) error {
	cmd := a.cli.Set(internalCtx, key, value, expiration)
	return cmd.Err()
}

// Get 获取值
func (a *Client) Get(key string) (string, error) {
	cmd := a.cli.Get(internalCtx, key)
	if err := cmd.Err(); err != nil {
		if err == redis.Nil {
			return "", nil
		}
	}
	return cmd.Val(), nil
}

// Exists 判断是否存在值
func (a *Client) Exists(key string) (bool, error) {
	cmd := a.cli.Exists(internalCtx, key)
	val, err := cmd.Result()
	if err != nil {
		return false, err
	}
	return val > 0, nil
}

// Del 删除key
func (a *Client) Del(key string) error {
	cmd := a.cli.Del(internalCtx, key)
	return cmd.Err()
}

// Incr 自增加值
func (a *Client) Incr(key string) error {
	cmd := a.cli.Incr(internalCtx, key)
	return cmd.Err()
}

// Close 关闭连接
func (a *Client) Close() {
	a.cli.Close()
}

// PSubscribe ...
func (a *Client) PSubscribe(channel string) *redis.PubSub {
	a.cli.ConfigSet(internalCtx, "set notify-keyspace-events", "Ex")
	return a.cli.PSubscribe(internalCtx, channel)
}
