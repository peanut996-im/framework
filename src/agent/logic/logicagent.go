package logic

import (
	"framework/cfgargs"
)

var logicAgent LogicAgent

type LogicAgent interface {
	Init(cfg *cfgargs.SrvConfig)
	Auth(token string) (interface{}, error)
	GetUserInfo()
	CheckToken()
	LoadInitData(uid string) (interface{}, error)
}

func GetlastLogicAgent() LogicAgent {
	return logicAgent
}
