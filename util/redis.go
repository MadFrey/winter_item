package util

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

var rdb *redis.Client

func InitClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func StoreVerCode(code string) {
	rdb.Set(code, fmt.Sprintf("%s", code), 60)
}

func VerifyCode(code string) bool {
	realCode, err := rdb.Get(code).Result()
	if err != nil {
		log.Println(err)
		return false
	}
	if code != realCode {
		return false
	}
	return true
}
