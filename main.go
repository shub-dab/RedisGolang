package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func yourRedisClient() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	return redisClient
}

func insertIt(key, value string) any {
	err := yourRedisClient().Set(key, value, 0).Err()

	if err != nil {
		return err
	}

	return "insertion successful"
}

func delete_(key string) any {
	deleted, err := yourRedisClient().Del(key).Result()

	if err != nil {
		log.Println(err)
	}

	return deleted
}

func insertInList(key string, list []string) any {
	insertResult, err := yourRedisClient().LPush(key, list).Result()
	if err != nil {
		fmt.Println(insertResult)
	}
	return err
}

func rangeList(key string) any {
	rangeResult, err := yourRedisClient().LRange(key, 0, -1).Result()

	if err != nil {
		return err
	}
	return rangeResult
}

func displayAll() {
	result, err := yourRedisClient().Keys("*").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

// "lrange", "Fruits" ,"0", "-1"

func unsupported() {
	result, err := yourRedisClient().Do("lrange", "Fruits" ,"0", "-1").Result()

	if err != nil {
		if err == redis.Nil {
			fmt.Println("key does not exist")
			return
		}
		panic(err)
	}
	fmt.Println(result)
}


func main() {
	response, err := yourRedisClient().Ping().Result()
	if err != nil {
		fmt.Println(response)
	}
	fmt.Println(response)

	// fmt.Println(insertIt("name1", "Aakif"))
	// fmt.Println(insertIt("name2", "Shwetank"))

	// fmt.Println(delete_("name2"))

	// fruits := []string{"Banana", "Grapes", "Mango"}
	// insertInList("Fruits", fruits)

	// fmt.Print(rangeList("Fruits"))
	// displayAll()

	unsupported()
}
