package logic

import (
	"framework/cfgargs"
)

var logicBroker LogicBroker

type LogicBroker interface {
	Init(cfg *cfgargs.SrvConfig)
	Send(event string, data interface{}) (interface{}, error)
	Listen()
	Register()
}

func GetlastLogicBroker() LogicBroker {
	return logicBroker
}
