package register

import (
	"encoding/json"
	"net"

	"gitee.com/yuanxuezhe/dante/module"
	base "gitee.com/yuanxuezhe/dante/module/base"
	"gitee.com/yuanxuezhe/dante/module/register"
)

var NewModule = func() module.Module {
	mod := &Register{}
	mod.ModuleType = "Register"
	mod.ModuleVersion = "1.0.0"
	mod.Basemodule.DoWork = mod.DoWork
	return mod
}

type ConnSet map[net.Conn]struct{}

type Register struct {
	register.BaseRegister
}

func (m *Register) init() {
	m.Modules.SetCapacity(50)
	m.Modules.SetDescribe(m.ModuleId + "  " + "m.Modules  ")
}

func (m *Register) DoWork(buff []byte) ([]byte, error) {
	var err error
	// 解析消息体
	moduleInfo := base.ModuleInfo{}
	err = json.Unmarshal(buff, &moduleInfo) //转换成JSON返回的是byte[]

	if err != nil {
		return nil, err
	}
	var moduleInfos []base.ModuleInfo
	v, err := m.Modules.Get(moduleInfo.ModuleType)
	if err == nil {
		moduleInfos = v.([]base.ModuleInfo)
	}
	m.Modules.Set(moduleInfo.ModuleType, append(moduleInfos, moduleInfo))
	// 创建注册连接
	err = m.RegisterBeats()
	if err == nil {
		return nil, err
	}
	return []byte("Register successful:" + string(buff)), nil
}
