package pool

import (
	"database/sql"
	"fmt"
	"gitee.com/yuanxuezhe/dante/conf"
	"gitee.com/yuanxuezhe/dante/pool"
	"github.com/streadway/amqp"
)

var (
	MqPool  *pool.ConnPool
	SqlPool *pool.ConnPool
)

func InitConnpoll() {
	fmt.Println("初始化 Mysql 连接池，连接数：%d \n", conf.Conf.Pools["Mysql"].MaxOpenConns)
	SqlPool = &pool.ConnPool{
		MaxActive: conf.Conf.Pools["Mysql"].MaxOpenConns,
		Dial: func() (interface{}, error) {
			conn, err := sql.Open("mysql", conf.Conf.Pools["Mysql"].Url)
			return conn, err
		},
	}
	SqlPool.InitPool()
	fmt.Println("初始化 Rabbitmq 连接池，连接数：%d \n", conf.Conf.Pools["Rabbitmq"].MaxOpenConns)
	MqPool = &pool.ConnPool{
		MaxActive: conf.Conf.Pools["Rabbitmq"].MaxOpenConns,
		Dial: func() (interface{}, error) {
			conn, err := amqp.Dial(conf.Conf.Pools["Rabbitmq"].Url)
			return conn, err
		},
	}
	MqPool.InitPool()
}
