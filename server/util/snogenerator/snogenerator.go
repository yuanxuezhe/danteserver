package snogenerator

import (
	"gitee.com/yuanxuezhe/dante/db/mysql"
	"log"
	"sync"
)

var useridRW sync.RWMutex

func NewUserid() int {
	var userid int
	conn := mysql.GetMysqlDB()
	useridRW.Lock()
	err := conn.QueryRow("SELECT max(userid) + 1  FROM userinfo").Scan(&userid)
	useridRW.Unlock()
	if err != nil {
		log.Fatal(err)
	}
	return userid
}
