package cmd

import (
	"fmt"
	"testing"

	"github.com/gomodule/redigo/redis"
)

var conn redis.Conn

func TestMain(m *testing.M) {
	conn := GetCacheConn()
	defer conn.Close()

    code := m.Run() 
	fmt.Printf("test result: %d", code)
	fmt.Println()
}

func TestAdd(t *testing.T) {
	k := "test"
	v := "val"

	AppendMember(conn, k, v)

	if MemberExists(conn, k, v) == false {
		t.Error("add test did not insert")
	}

	//clean up test insertion
	RemoveMember(k, v)	
}

func TestFalseAdd(t *testing.T) {
	k := "test"
	v := "val"

	AppendMember(conn, k, v)

	if MemberExists(conn, k, v) == false {
		t.Error("add test did not insert")
	}
}