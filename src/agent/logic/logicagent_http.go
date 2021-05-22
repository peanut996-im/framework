package logic

import (
	"encoding/json"
	"errors"
	"fmt"
	"framework/cfgargs"
	"framework/logger"
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

//Auth Authentication from the logic layer
func (l *LogicAgentHttp) Auth(token string) (interface{}, error) {
	url := l.logicAddr + "/auth"
	resp, body, errs := l.client.GetGoReq().Post(url).Send(fmt.Sprintf(`{ "token": "%v"}`, token)).End()
	if len(errs) != 0 {
		for i, err := range errs {
			logger.Info("logicAgent Auth failed. errs[%v]: %v ", i, err)
		}
		return nil, errs[0]
	}
	if resp.StatusCode != 200 {
		logger.Debug("err here")
		return nil, errors.New(fmt.Sprintf("LogicAgent Auth http failed code: %v", resp.StatusCode))
	}

	//r := &api.BaseRepsonse{}
	//err := json.Unmarshal([]byte(body), r)
	//if r.Code != api.ERROR_CODE_OK {
	//	logger.Debug("err here")
	//	return nil, errors.New(fmt.Sprintf("logicAgent Auth failed. err: %v", r.Message))
	//}
	//if nil != err {
	//	logger.Debug("err here")
	//	logger.Info("logicAgent Auth failed. json format err: %v ", err)
	//	return nil, err
	//}
	// Data is model.User
	return json.RawMessage(body), nil
}

func (l *LogicAgentHttp) GetUserInfo() {

}

func (l *LogicAgentHttp) CheckToken() {

}

func (l *LogicAgentHttp) LoadInitData(uid string) (interface{}, error) {
	url := l.logicAddr + "/load"
	resp, body, errs := l.client.GetGoReq().Post(url).Send(fmt.Sprintf(`{ "user_id": "%v"}`, uid)).End()
	if resp.StatusCode != 200 {
		logger.Debug("err here")
		return nil, errors.New(fmt.Sprintf("LogicAgent Auth http failed code: %v", resp.StatusCode))
	}
	if len(errs) != 0 {
		for i, err := range errs {
			logger.Info("logicAgent Auth failed. errs[%v]: %v ", i, err)
		}
		logger.Debug("err here")
		return "", errs[0]
	}
	// forwards to user need raw json
	return json.RawMessage(body), nil
}
