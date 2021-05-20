package logic

import (
	"framework/api"
	"framework/cfgargs"
)

var logicAgent LogicAgent

type LogicAgent interface {
	Init(cfg *cfgargs.SrvConfig)
	Auth(token string) (*api.User, error)
	GetUserInfo()
	CheckToken()
}

func GetlastLogicAgent() LogicAgent {
	return logicAgent
}
