package main

import (
	"log"
	"menteslibres.net/gosexy/redis"
)

var client *redis.Client

func dbGet(key string) (value string) {
	s, _ := client.Get(key)
	return s
}

func dbSet(key string, value string) {
	client.Set(key, value)
}

func dbPop(key string) (value string) {
	client.Lpop(key)
}

func dbPush(key string, value string) {
	client.Lpush(key, value)
}

func dbConn() {

	var err error

	client = redis.New()

	err = client.Connect("127.0.0.1", 6379)

	if err != nil {
		log.Fatalf("Connect failed: %s\n", err.Error())
		return
	}

	log.Println("Connected to redis-server.")

}

func dbClose() {

	client.Quit()

}
