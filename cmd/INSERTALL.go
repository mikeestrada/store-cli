package cmd

import (
	"fmt"
)

// InsertAll inserts all given members for a key
func InsertAll(k string, args []string) {
	conn := GetCacheConn()

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