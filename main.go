package main

import (
	"danteserver/server/modules/gateway"
	"danteserver/server/modules/goods"
	"danteserver/server/modules/login"
	"danteserver/server/modules/register"

	"gitee.com/yuanxuezhe/dante"
	_ "gitee.com/yuanxuezhe/dante/conf"
	"gitee.com/yuanxuezhe/dante/log"
)

func main() {
	log.LogPrint(&log.LogStruct{log.LEVEL_DEBUG,"hhhhhhhhhhhhhDDD %s%s%s\n", [] interface {}{"xxxxx", "11", "22"},})
	log.LogPrint(&log.LogStruct{log.LEVEL_ERROR,"hhhhhhhhhhhhhEEE %s%d%d\n", [] interface {}{"xxxxx", 11, 22},})
	dante.AddMod("Gateway", gateway.NewModule)
	dante.AddMod("Register", register.NewModule)
	dante.AddMod("Login", login.NewModule)
	dante.AddMod("Goods", goods.NewModule)

	dante.Run()
}
