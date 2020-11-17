package main

// var (
// 	client = &redisClient{}
// )

// //type redisClient, declaration
// type redisClient struct {
// 	c *redis.Client
// }

// //type paramRedis, return data from redis
// type valRedis struct {
// 	Value interface{} `json:"value"`
// }

// func dbredis() {
// 	var err error
// 	client := initialize()

// 	//[]byte(`{ "nama": "asromi", "divisi": "IT" }`)

// 	key1 := "12345"
// 	value1 := &valRedis{Value: "asromi"}
// 	//value1 := []byte(`{ "nama": "asromi", "divisi": "IT" }`)

// 	expiration := time.Second * 20 //expiration logic; mPar.Expiration LIAR banget
// 	err = client.setKey(key1, value1, expiration)
// 	if err != nil {
// 		fmt.Println("redisClient.setKey:", err)
// 	}

// 	retRedis := &valRedis{}
// 	err = client.getKey("12345", retRedis)
// 	if err != nil {
// 		fmt.Println("redisClient.getKey:", err)
// 	}
// 	fmt.Println(retRedis.Value)

// 	err = client.delKey("12345")
// 	if err != nil {
// 		fmt.Println("redisClient.delKey:", err)
// 	}

// }

// //GetClient get the redis client
// func initialize() *redisClient {
// 	c := redis.NewClient(&redis.Options{
// 		Addr:     "127.0.0.1:6300",
// 		Password: "",
// 		DB:       0,
// 	})

// 	if err := c.Ping().Err(); err != nil {
// 		panic("Unable to connect to redis " + err.Error())
// 	}
// 	client.c = c
// 	return client
// }

// //GetKey get key redis
// func (client *redisClient) getKey(key string, src interface{}) error {
// 	val, err := client.c.Get(key).Result()
// 	if err == redis.Nil || err != nil {
// 		return err
// 	}
// 	err = json.Unmarshal([]byte(val), &src)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// //SetKey set key redis
// func (client *redisClient) setKey(key string, value interface{}, expiration time.Duration) error {
// 	cacheEntry, err := json.Marshal(value)
// 	if err != nil {
// 		return err
// 	}
// 	err = client.c.Set(key, cacheEntry, expiration).Err()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// //delKey delete key redis
// func (client *redisClient) delKey(key string) error {
// 	err := client.c.Del(key).Err()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// //isJSON : Validate JSON Format
// func isJSON(s string) bool {
// 	var js map[string]interface{}
// 	return json.Unmarshal([]byte(s), &js) == nil
// }
