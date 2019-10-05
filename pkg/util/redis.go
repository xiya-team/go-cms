package util

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"sync"
	"time"
)

type Redis struct {
	Connect   string //连接字符串
	Db        int    //数据库
	Maxidle   int    //最大空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态
	Maxactive int    //最大的激活连接数，表示同时最多有N个连接
}

var (
	r    *Redis
	once sync.Once
	redisClient *redis.Pool
)
/**
 * 返回单例实例
 * @method New
 */
func New(connect string, db int, maxidle int, maxactive int) *Redis {
	once.Do(func() { //只执行一次
		r = &Redis{Connect: connect, Db: db, Maxidle:maxidle, Maxactive:maxactive}
        setPoll()
	})
	return r
}

/**
 * 公共方法
 */
/**
 * 设置连接池
 * @method setPoll
 */
func setPoll() {
	redisClient = &redis.Pool{
		MaxIdle:     r.Maxidle,             // idle的列表长度, 空闲的线程数
		MaxActive:   r.Maxactive,         // 线程池的最大连接数， 0表示没有限制
		Wait:        true,              // 当连接数已满，是否要阻塞等待获取连接。false表示不等待，直接返回错误。
		IdleTimeout: 180*time.Second,
		Dial: func() (redis.Conn, error) { // 创建链接
			c, err := redis.Dial("tcp", beego.AppConfig.String("redis_addr"))
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", beego.AppConfig.String("redis_password")); err != nil {
				c.Close()
				return nil, err
			}
			if _, err := c.Do("SELECT", beego.AppConfig.String("redis_index")); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error { //一个测试链接可用性
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

/**
 * 执行基本命令
 * @method func
 * @param  {[type]} n *Neo4j        [description]
 * @return {[type]}   [description]
 */
func (n *Redis) Do(cmd string, args ...interface{}) (interface{}, error) {
	conn := redisClient.Get()
	defer conn.Close()
	return conn.Do(cmd, args...)
}

/**
 * 设置键值对, ex单位是秒
 * @method func
 * @param  {[type]} n *Redis        [description]
 * @return {[type]}   [description]
 */
func (n *Redis) SetString(key string, value string, ex string) (interface{}, error) {
	conn := redisClient.Get()
	defer conn.Close()
	return conn.Do("SET", key, value, "EX", ex)
}
/**
 * 获取键的值
 * @method func
 * @param  {[type]} n *Redis        [description]
 * @return {[type]}   [description]
 */
func (n *Redis) GetString(key string) (string, error) {
  conn := redisClient.Get()
  defer conn.Close()
	value, err := redis.String(conn.Do("GET", key))
  return value, err
}

// del
func DelKey(key string) error {
	conn := redisClient.Get()
	defer conn.Close()
	_, err := conn.Do("DEL", key)
	return err
}

// lrange
func LRange(key string, start, stop int64) ([]string, error) {
	conn := redisClient.Get()
	defer conn.Close()
	return redis.Strings(conn.Do("LRANGE", key, start, stop))
}

// lpop
func LPop(key string) (string, error) {
	conn := redisClient.Get()
	defer conn.Close()
	return redis.String(conn.Do("LPOP", key))
}

// LPushAndTrimKey
func LPushAndTrimKey(key, value interface{}, size int64) error {
	conn := redisClient.Get()
	defer conn.Close()
	conn.Send("MULTI")
	conn.Send("LPUSH", key, value)
	conn.Send("LTRIM", key, size-2*size, -1)
	_, err := conn.Do("EXEC")
	return err
}

// RPushAndTrimKey
func RPushAndTrimKey(key, value interface{}, size int64) error {
	conn := redisClient.Get()
	defer conn.Close()
	conn.Send("MULTI")
	conn.Send("RPUSH", key, value)
	conn.Send("LTRIM", key, size-2*size, -1)
	_, err := conn.Do("EXEC")
	return err
	
}

// ExistsKey
func ExistsKey(key string) (bool, error) {
	conn := redisClient.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("EXISTS", key))
}

// ttl 返回剩余时间
func TTLKey(key string) (int64, error) {
	conn := redisClient.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("TTL", key))
}

// incr 自增
func Incr(key string) (int64, error) {
	conn := redisClient.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("INCR", key))
}

// Decr 自减
func Decr(key string) (int64, error) {
	conn := redisClient.Get()
	defer conn.Close()
	return redis.Int64(conn.Do("DECR", key))
}

// mset 批量写入 conn.Do("MSET", "ket1", "value1", "key2","value2")
func MsetKey(key_value ...interface{}) error {
	conn := redisClient.Get()
	defer conn.Close()
	_, err := conn.Do("MSET", key_value...)
	return err
}

// mget  批量读取 mget key1, key2, 返回map结构
func MgetKey(keys ...interface{}) map[interface{}]string {
	conn := redisClient.Get()
	defer conn.Close()
	values, _ := redis.Strings(conn.Do("MGET", keys...))
	resultMap := map[interface{}]string{}
	keyLen := len(keys)
	for i := 0; i < keyLen; i++ {
		resultMap[keys[i]] = values[i]
	}
	return resultMap
}