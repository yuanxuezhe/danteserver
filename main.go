package main

import (
	"danteserver/server/modules/gateway"
	"danteserver/server/modules/goods"
	"danteserver/server/modules/login"
	"danteserver/server/modules/register"
	_ "danteserver/server/util/pool"

	"gitee.com/yuanxuezhe/dante/core"
	_ "gitee.com/yuanxuezhe/dante/core/conf"
)

func main() {
	core.AddMod("Gateway", gateway.NewModule)
	core.AddMod("Register", register.NewModule)
	core.AddMod("Login", login.NewModule)
	core.AddMod("Goods", goods.NewModule)

	core.Run()
}
