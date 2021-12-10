package main

import (
	"danteserver/server/modules/gateway"
	"danteserver/server/modules/goods"
	"danteserver/server/modules/login"
	"danteserver/server/modules/register"

	"gitee.com/yuanxuezhe/dante"
	_ "gitee.com/yuanxuezhe/dante/conf"
)

func main() {
	dante.AddMod("Gateway", gateway.NewModule)
	dante.AddMod("Register", register.NewModule)
	dante.AddMod("Login", login.NewModule)
	dante.AddMod("Goods", goods.NewModule)

	dante.Run()
}
