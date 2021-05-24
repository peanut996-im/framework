package logic

import (
	"framework/cfgargs"
)

var logicBroker LogicBroker

type LogicBroker interface {
	Init(*cfgargs.SrvConfig)
	Send(string, interface{}) (interface{}, error)
	Listen(interface{})
	Register()
}

func GetlastLogicBroker() LogicBroker {
	return logicBroker
}
