package snogenerator

import (
	"danteserver/server/pool"
	"database/sql"
	"log"
	"sync"
)

var useridRW sync.RWMutex

func NewUserid() int {
	var userid int
	conn := pool.SqlPool.Get().(*sql.DB)
	useridRW.Lock()
	err := conn.QueryRow("SELECT max(userid) + 1  FROM userinfo").Scan(&userid)
	useridRW.Unlock()
	if err != nil {
		log.Fatal(err)
	}
	return userid
}
