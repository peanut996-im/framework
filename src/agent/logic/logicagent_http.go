package logic

import (
	"encoding/json"
	"errors"
	"fmt"
	"framework/api"
	"framework/cfgargs"
	"framework/logger"
	"framework/net"
	"framework/net/http"
)

type LogicAgentHttp struct {
	srv       *http.Server
	client    *http.Client
	panic     bool
	logicAddr string
}

func NewLogicAgentHttp() *LogicAgentHttp {
	return &LogicAgentHttp{}
}
func (l *LogicAgentHttp) Init(cfg *cfgargs.SrvConfig) {
	l.srv = http.NewServer(cfg)
	l.client = http.NewClient()
	if cfg.Logic.Mode != "http" {
		logger.Warn("can't load logic configuration for http")
		if l.panic {
			panic(errors.New("can't load logic configuration for http"))
		}
		return
	}
	addr := fmt.Sprintf("http://%v:%v", cfg.Logic.Host, cfg.Logic.Port)
	l.logicAddr = addr
	// global
	logicAgent = l
}
func (l *LogicAgentHttp) Auth(token string) (*api.User, error) {
	url := l.logicAddr + "/auth"
	resp, body, errs := l.client.GetGoReq().Post(url).Send(fmt.Sprintf(`{ "token": "%v"}`, token)).End()
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("LogicAgent Auth http failed code: %v", resp.StatusCode))
	}
	if len(errs) != 0 {
		for i, err := range errs {
			logger.Info("logicAgent Auth failed. errs[%v]: %v ", i, err)
		}
		return nil, errs[0]
	}
	user := &api.User{}
	j := make(map[string]interface{})
	err := json.Unmarshal([]byte(body), &j)
	if j["code"] != net.ERROR_CODE_OK {
		return nil, errors.New(fmt.Sprintf("logicAgent Auth failed. err: %v", j["message"]))
	}
	err = json.Unmarshal([]byte(j["data"].(string)), user)
	if nil != err {
		return nil, err
	}
	return user, nil
}

func (l *LogicAgentHttp) GetUserInfo() {

}

func (l *LogicAgentHttp) CheckToken() {

}
