module danteserver

go 1.14

require (
	gitee.com/yuanxuezhe/dante v0.0.0-incompatible
	gitee.com/yuanxuezhe/ynet v1.1.6
)

replace (
	gitee.com/yuanxuezhe/dante => ../dante
	gitee.com/yuanxuezhe/ynet => ../ynet
)
