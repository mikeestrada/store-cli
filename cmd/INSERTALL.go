package cmd

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// InsertAll inserts all given members for a key
func InsertAll(r redis.Conn, k string, args []string) {
	var conn redis.Conn
	if r == nil {
		conn = GetCacheConn()
		defer conn.Close()
	} else {
		conn = r
	}

	_, err := conn.Do("SET", k, sliceToString(args))
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	defer conn.Close()
}

// sliceToString converts slice into comma-delimited string
func sliceToString(slice []string) string {
	var s string
	for i := 0; i < len(slice); i++ {
		s += slice[i] + ","
	}
	return s
}