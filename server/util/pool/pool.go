package pool

import (
	"database/sql"
	"gitee.com/yuanxuezhe/dante/conf"
	//"gitee.com/yuanxuezhe/ynet/yconnpool"
)

var MysqlDb *sql.DB

func init() {
	var err error
	MysqlDb, err = sql.Open("mysql", conf.Conf.Mysql.Url)
	if err != nil {
		panic(err)
	}

	MysqlDb.SetMaxOpenConns(conf.Conf.Mysql.MaxOpenConns)
	MysqlDb.SetMaxIdleConns(conf.Conf.Mysql.MaxIdleConns)
}
