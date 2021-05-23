package logic

import (
	"encoding/json"
	"errors"
	"fmt"
	"framework/api"
	"framework/cfgargs"
	"framework/logger"
	"framework/net/http"
)

type LogicBrokerHttp struct {
	srv       *http.Server
	client    *http.Client
	panic     bool
	logicAddr string
}

func (l *LogicBrokerHttp) Send(event string, data interface{}) (interface{}, error) {
	logger.Info("logicBroker Event /%v Send.", event)
	url := l.logicAddr + "/" + event
	rawJson := ``
	switch event {
	case api.EventAuth:
		rawJson = fmt.Sprintf(`{ "token": "%v"}`, data.(string))
	case api.EventLoad:
		rawJson = fmt.Sprintf(`{ "uid": "%v"}`, data.(string))
	case api.EventAddFriend:
	case api.EventDeleteFriend:
	case api.EventCreateGroup:
	case api.EventJoinGroup:
	case api.EventLeaveGroup:
	default:
		return nil, nil
	}
	resp, body, errs := l.client.GetGoReq().Post(url).Send(rawJson).End()
	if len(errs) != 0 {
		for i, err := range errs {
			logger.Info("logicBroker Event /%v failed. errs[%v]: %v ", event, i, err)
		}
		return nil, errs[0]
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("LogicBroker Event /%v http failed code: %v", event, resp.StatusCode))
	}
	return json.RawMessage(body), nil
}

func (l *LogicBrokerHttp) Listen() {

}

func (l *LogicBrokerHttp) Register() {

}

func NewLogicBrokerHttp() *LogicBrokerHttp {
	return &LogicBrokerHttp{}
}
func (l *LogicBrokerHttp) Init(cfg *cfgargs.SrvConfig) {
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
	logicBroker = l
}

//Auth Authentication from the logic layer
func (l *LogicBrokerHttp) Auth(token string) (interface{}, error) {
	url := l.logicAddr + "/auth"
	resp, body, errs := l.client.GetGoReq().Post(url).Send(fmt.Sprintf(`{ "token": "%v"}`, token)).End()
	if len(errs) != 0 {
		for i, err := range errs {
			logger.Info("logicBroker Auth failed. errs[%v]: %v ", i, err)
		}
		return nil, errs[0]
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("LogicBroker Auth http failed code: %v", resp.StatusCode))
	}
	return json.RawMessage(body), nil
}

func (l *LogicBrokerHttp) GetUserInfo() {

}

func (l *LogicBrokerHttp) LoadInitData(uid string) (interface{}, error) {
	url := l.logicAddr + "/load"
	resp, body, errs := l.client.GetGoReq().Post(url).Send(fmt.Sprintf(`{ "user_id": "%v"}`, uid)).End()
	if len(errs) != 0 {
		for i, err := range errs {
			logger.Info("logicBroker Auth failed. errs[%v]: %v ", i, err)
		}
		return "", errs[0]
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("LogicBroker Auth http failed code: %v", resp.StatusCode))
	}
	// forwards to user need raw json
	return json.RawMessage(body), nil
}
func (l *LogicBrokerHttp) AddFriend(addFriendRequest interface{}) (interface{}, error) {
	url := l.logicAddr + "/addFriend"
	data, err := json.Marshal(addFriendRequest)
	if err != nil {
		return nil, err
	}
	resp, body, errs := l.client.GetGoReq().Post(url).Send(data).End()
	if len(errs) != 0 {
		for i, err := range errs {
			logger.Info("logicBroker AddFriend failed. errs[%v]: %v ", i, err)
		}
		return "", errs[0]
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("LogicBroker AddFriend http failed code: %v", resp.StatusCode))
	}
	return json.RawMessage(body), nil
}

func (l *LogicBrokerHttp) DeleteFriend(addFriendRequest interface{}) (interface{}, error) {
	url := l.logicAddr + "/addFriend"
	data, err := json.Marshal(addFriendRequest)
	if err != nil {
		return nil, err
	}
	resp, body, errs := l.client.GetGoReq().Post(url).Send(data).End()
	if len(errs) != 0 {
		for i, err := range errs {
			logger.Info("logicBroker AddFriend failed. errs[%v]: %v ", i, err)
		}
		return "", errs[0]
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("LogicBroker AddFriend http failed code: %v", resp.StatusCode))
	}
	return json.RawMessage(body), nil
}
