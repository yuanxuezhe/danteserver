module danteserver

go 1.14

require (
	gitee.com/yuanxuezhe/dante v0.0.0-incompatible
	gitee.com/yuanxuezhe/ynet v1.1.6
	github.com/go-sql-driver/mysql v1.5.0
	github.com/goinggo/mapstructure v0.0.0-20140717182941-194205d9b4a9 // indirect
	github.com/streadway/amqp v1.0.0
)

replace (
	gitee.com/yuanxuezhe/dante => ../dante
	gitee.com/yuanxuezhe/ynet => ../ynet
)
