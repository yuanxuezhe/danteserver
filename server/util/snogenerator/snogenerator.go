package snogenerator

import (
	"danteserver/server/util/pool"
	"log"
	"sync"
)

var useridRW sync.RWMutex

func NewUserid() int {
	var userid int
	useridRW.Lock()
	err := pool.MysqlDb.QueryRow("SELECT max(userid) + 1  FROM userinfo").Scan(&userid)
	useridRW.Unlock()
	if err != nil {
		log.Fatal(err)
	}
	return userid
}
