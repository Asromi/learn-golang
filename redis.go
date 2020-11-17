package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var (
	client = &redisClient{}
)

//type redisClient, declaration
type redisClient struct {
	c *redis.Client
}

func dbredis() {
	var err error
	client := initialize()

	//[]byte(`{ "nama": "asromi", "divisi": "IT" }`)

	key1 := "apiCOREtemplateeventsMenu"                     //string dataType
	value1 := []byte(`{ "dbCache": "REDIS", "crud": "R" }`) //interface{} dataType

	expiration := time.Minute * 30 //expiration logic; mPar.Expiration LIAR banget
	err = client.setKey(key1, value1, expiration)
	if err != nil {
		fmt.Println("redisClient.setKey:", err)
	}

	retRedis, err := client.getKey(key1)
	if err != nil {
		fmt.Println("redisClient.getKey:", err)
	}
	fmt.Println(retRedis)

	// err = client.delKey(key1)
	// if err != nil {
	// 	fmt.Println("redisClient.delKey:", err)
	// }

}

//GetClient get the redis client
func initialize() *redisClient {
	c := redis.NewClient(&redis.Options{
		Addr:       "localhost:6300", // "192.168.20.75:6300", //
		Password:   "",
		DB:         0,
		MaxRetries: 3,
	})
	fmt.Println(c)

	if err := c.Ping().Err(); err != nil {
		panic("Unable to connect to redis " + err.Error())
	}

	// pong, err := rdb.Ping(ctx).Result()
	// fmt.Println(pong, err)

	client.c = c
	return client
}

//GetKey get key redis
func (client *redisClient) getKey(key string) (interface{}, error) {
	val, err := client.c.Get(key).Result()
	if err == redis.Nil || err != nil {
		return nil, err
	}
	return val, nil
}

//SetKey set key redis
func (client *redisClient) setKey(key string, value interface{}, expiration time.Duration) error {
	err := client.c.Set(key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

//delKey delete key redis
func (client *redisClient) delKey(key string) error {
	err := client.c.Del(key).Err()
	if err != nil {
		return err
	}
	return nil
}

//isJSON : Validate JSON Format
func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}
