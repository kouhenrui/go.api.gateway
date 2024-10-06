package config

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// RedisClient 封装 Redis 客户端
type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

// RedisConfig 配置结构体
type RedisConfig struct {
	Addr     string // Redis 服务器地址，格式为 "localhost:6379"
	Password string // Redis 服务器密码
	DB       int    // 使用的数据库编号，默认是 0
	PoolSize int    // 连接池大小
}

// NewRedisClient 初始化 Redis 客户端
func NewRedisClient(config *RedisConfig) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
		PoolSize: config.PoolSize,
	})

	// Ping 测试
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("连接 Redis 失败: %v", err))
	}

	return &RedisClient{
		client: rdb,
		ctx:    context.Background(),
	}
}

// Close 关闭 Redis 连接
//func (r *RedisClient) Close() {
//	r.client.Close()
//}

// Set 设置键值，带过期时间
func (r *RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	err := r.client.Set(r.ctx, key, value, expiration).Err()
	if err != nil {
		return fmt.Errorf("设置键 %s 失败: %v", key, err)
	}
	return nil
}

// Get 获取键的值
func (r *RedisClient) Get(key string) (string, error) {
	val, err := r.client.Get(r.ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("键 %s 不存在", key)
	} else if err != nil {
		return "", fmt.Errorf("获取键 %s 失败: %v", key, err)
	}
	return val, nil
}

// Del 删除键
func (r *RedisClient) Del(keys ...string) error {
	err := r.client.Del(r.ctx, keys...).Err()
	if err != nil {
		return fmt.Errorf("删除键 %v 失败: %v", keys, err)
	}
	return nil
}

// Exists 检查键是否存在
func (r *RedisClient) Exists(key string) (bool, error) {
	val, err := r.client.Exists(r.ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("检查键 %s 是否存在失败: %v", key, err)
	}
	return val == 1, nil
}

// Expire 设置键的过期时间
func (r *RedisClient) Expire(key string, expiration time.Duration) error {
	err := r.client.Expire(r.ctx, key, expiration).Err()
	if err != nil {
		return fmt.Errorf("设置键 %s 过期时间失败: %v", key, err)
	}
	return nil
}

// Incr 自增操作
func (r *RedisClient) Incr(key string) (int64, error) {
	val, err := r.client.Incr(r.ctx, key).Result()
	if err != nil {
		return 0, fmt.Errorf("键 %s 自增失败: %v", key, err)
	}
	return val, nil
}

// Decr 自减操作
func (r *RedisClient) Decr(key string) (int64, error) {
	val, err := r.client.Decr(r.ctx, key).Result()
	if err != nil {
		return 0, fmt.Errorf("键 %s 自减失败: %v", key, err)
	}
	return val, nil
}

// LPush 向列表左侧插入元素
func (r *RedisClient) LPush(key string, values ...interface{}) error {
	err := r.client.LPush(r.ctx, key, values...).Err()
	if err != nil {
		return fmt.Errorf("向列表 %s 插入元素失败: %v", key, err)
	}
	return nil
}

// RPush 向列表右侧插入元素
func (r *RedisClient) RPush(key string, values ...interface{}) error {
	err := r.client.RPush(r.ctx, key, values...).Err()
	if err != nil {
		return fmt.Errorf("向列表 %s 插入元素失败: %v", key, err)
	}
	return nil
}

// LPop 从列表左侧弹出元素
func (r *RedisClient) LPop(key string) (string, error) {
	val, err := r.client.LPop(r.ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("从列表 %s 弹出元素失败: %v", key, err)
	}
	return val, nil
}

// RPop 从列表右侧弹出元素
func (r *RedisClient) RPop(key string) (string, error) {
	val, err := r.client.RPop(r.ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("从列表 %s 弹出元素失败: %v", key, err)
	}
	return val, nil
}

// LRange 获取列表中指定范围的元素
func (r *RedisClient) LRange(key string, start, stop int64) ([]string, error) {
	vals, err := r.client.LRange(r.ctx, key, start, stop).Result()
	if err != nil {
		return nil, fmt.Errorf("获取列表 %s 的范围元素失败: %v", key, err)
	}
	return vals, nil
}

// HSet 设置哈希表中的字段值
func (r *RedisClient) HSet(key, field string, value interface{}) error {
	err := r.client.HSet(r.ctx, key, field, value).Err()
	if err != nil {
		return fmt.Errorf("设置哈希表 %s 字段 %s 失败: %v", key, field, err)
	}
	return nil
}

// HGet 获取哈希表中的字段值
func (r *RedisClient) HGet(key, field string) (string, error) {
	val, err := r.client.HGet(r.ctx, key, field).Result()
	if err != nil {
		return "", fmt.Errorf("获取哈希表 %s 字段 %s 失败: %v", key, field, err)
	}
	return val, nil
}

// HDel 删除哈希表中的字段
func (r *RedisClient) HDel(key string, fields ...string) error {
	err := r.client.HDel(r.ctx, key, fields...).Err()
	if err != nil {
		return fmt.Errorf("删除哈希表 %s 的字段 %v 失败: %v", key, fields, err)
	}
	return nil
}

// HGetAll 获取哈希表中的所有字段和值
func (r *RedisClient) HGetAll(key string) (map[string]string, error) {
	vals, err := r.client.HGetAll(r.ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("获取哈希表 %s 的所有字段和值失败: %v", key, err)
	}
	return vals, nil
}

// SAdd 向集合中添加元素
func (r *RedisClient) SAdd(key string, members ...interface{}) error {
	err := r.client.SAdd(r.ctx, key, members...).Err()
	if err != nil {
		return fmt.Errorf("向集合 %s 添加元素失败: %v", key, err)
	}
	return nil
}

// SMembers 获取集合中的所有元素
func (r *RedisClient) SMembers(key string) ([]string, error) {
	vals, err := r.client.SMembers(r.ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("获取集合 %s 的元素失败: %v", key, err)
	}
	return vals, nil
}

// SRem 从集合中移除元素
func (r *RedisClient) SRem(key string, members ...interface{}) error {
	err := r.client.SRem(r.ctx, key, members...).Err()
	if err != nil {
		return fmt.Errorf("从集合 %s 移除元素失败: %v", key, err)
	}
	return nil
}

// TTL 获取键的过期时间
func (r *RedisClient) TTL(key string) (time.Duration, error) {
	ttl, err := r.client.TTL(r.ctx, key).Result()
	if err != nil {
		return 0, fmt.Errorf("获取键 %s 的过期时间失败: %v", key, err)
	}
	return ttl, nil
}

// Persist 移除键的过期时间
func (r *RedisClient) Persist(key string) error {
	err := r.client.Persist(r.ctx, key).Err()
	if err != nil {
		return fmt.Errorf("移除键 %s 的过期时间失败: %v", key, err)
	}
	return nil
}

// SetBit 设置位操作
func (r *RedisClient) SetBit(key string, offset int64, value int) error {
	err := r.client.SetBit(r.ctx, key, offset, value).Err()
	if err != nil {
		return fmt.Errorf("设置键 %s 的位 %d 失败: %v", key, offset, err)
	}
	return nil
}

// GetBit 获取位操作
func (r *RedisClient) GetBit(key string, offset int64) (int64, error) {
	val, err := r.client.GetBit(r.ctx, key, offset).Result()
	if err != nil {
		return 0, fmt.Errorf("获取键 %s 的位 %d 失败: %v", key, offset, err)
	}
	return val, nil
}

// 流量监控
// XAdd 向 Redis Stream 中添加数据
func (r *RedisClient) XAdd(stream, id string, values map[string]interface{}) error {
	args := &redis.XAddArgs{
		Stream: stream,
		ID:     id,
		Values: values,
	}
	_, err := r.client.XAdd(r.ctx, args).Result()
	if err != nil {
		return fmt.Errorf("向 Stream %s 添加数据失败: %v", stream, err)
	}
	return nil
}

// XRead 从 Redis Stream 中读取数据
func (r *RedisClient) XRead(stream string, count int64, block time.Duration) ([]redis.XMessage, error) {
	args := &redis.XReadArgs{
		Streams: []string{stream, "0"},
		Count:   count,
		Block:   block,
	}
	result, err := r.client.XRead(r.ctx, args).Result()
	if err != nil {
		return nil, fmt.Errorf("从 Stream %s 读取数据失败: %v", stream, err)
	}
	return result[0].Messages, nil
}
