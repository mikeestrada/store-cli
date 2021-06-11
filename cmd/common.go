package cmd

import (
	"log"

	"github.com/gomodule/redigo/redis"
)
	
func GetCacheConn() redis.Conn {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
	return conn
}