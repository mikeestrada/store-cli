package cmd

import "testing"

func TestAdd(t *testing.T) {
	conn := GetCacheConn()
	defer conn.Close()

	k := "test"
	v := "val"

	AppendMember(k, v)

	if MemberExists(conn, k, v) == false {
		t.Error("add test did not insert")
	}
	
}