package broker

import (
	"encoding/json"
	"errors"
	"fmt"
	"framework/api"
	"framework/cfgargs"
	"framework/logger"
	"framework/net/http"
	"framework/tool"
	"github.com/gin-gonic/gin"
	http2 "net/http"
)

type GateBrokerHttp struct {
	cfg       *cfgargs.SrvConfig
	srv       *http.Server
	client    *http.Client
	panic     bool
	logicAddr string
}

func (g *GateBrokerHttp) Send(event string, data interface{}) (interface{}, error) {
	//logger.Info("logicBroker Event /%v Send.", event)
	url := g.logicAddr + "/" + event
	rawJson := ``
	switch event {
	case api.EventAuth:
		rawJson = fmt.Sprintf(`{ "token": "%v"}`, data.(string))
	case api.EventLoad:
		rawJson = fmt.Sprintf(`{ "uid": "%v"}`, data.(string))
	default:
		bytes, err := json.Marshal(data)
		if err != nil {
			return nil, errors.New(fmt.Sprintf(api.UnmarshalJsonError, err))
		}
		rawJson = string(bytes)
	}
	resp, body, errs := g.client.GetGoReq().Post(url).Send(rawJson).End()
	if len(errs) != 0 {
		for i, err := range errs {
			logger.Info("GateBroker Event /%v failed. errs[%v]: %v ", event, i, err)
		}
		return nil, errs[0]
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("GateBroker Event /%v http failed code: %v", event, resp.StatusCode))
	}
	return json.RawMessage(body), nil
}

func (g *GateBrokerHttp) Listen() {
	g.srv.Run()
}

func (g *GateBrokerHttp) Register() error {
	ip, _ := tool.GetIp()
	addr := fmt.Sprintf("http://%v:%v", ip, g.cfg.HTTP.Port)
	data := make(map[string]interface{})
	data["ip"] = addr
	res, err := g.Send("gate/register", data)
	if err != nil {
		if g.panic {
			panic(err)
		} else {
			return err
		}
	}
	err = json.Unmarshal(res.(json.RawMessage), &data)
	if data["code"] != 0 {
		if g.panic {
			panic("GateBroker.Register failed")
		} else {
			logger.Debug("GateBroker.Register failed")
			return err
		}
	}
	return nil

}

func (g *GateBrokerHttp) Update(scenes interface{}) {
	ip, _ := tool.GetIp()
	addr := fmt.Sprintf("http://%v:%v", ip, g.cfg.HTTP.Port)
	data := make(map[string]interface{})
	data["ip"] = addr
	data["scene"] = scenes.([]string)
	res, err := g.Send("gate/update", data)
	if err != nil {
		if g.panic {
			panic(err)
		} else {
			return
		}
	}
	err = json.Unmarshal(res.(json.RawMessage), &data)
	if data["code"] != 0 {
		if g.panic {
			panic("GateBroker.Update failed")
		} else {
			logger.Debug("GateBroker.Update failed")
			return
		}
	}

}

func NewGateBrokerHttp() *GateBrokerHttp {
	return &GateBrokerHttp{}
}

func (g *GateBrokerHttp) Init(cfg *cfgargs.SrvConfig) {
	g.cfg = cfg
	g.srv = http.NewServer()
	g.srv.Init(cfg)
	g.client = http.NewClient()
	if cfg.Logic.Mode != "http" {
		logger.Warn("can't load logic configuration for http")
		if cfg.Logic.Panic {
			panic(errors.New("can't load logic configuration for http"))
		}
		return
	}
	addr := fmt.Sprintf("http://%v:%v", cfg.Logic.Host, cfg.Logic.Port)
	node := http.NewNodeRoute("logic", http.NewRoute(api.HTTPMethodPost, "ping", g.PingPong()))
	g.AddNodeRoute(node)
	g.logicAddr = addr
}

// AddNodeRoute Mount the route to the internal HTTP server.
func (g *GateBrokerHttp) AddNodeRoute(nodes ...*http.NodeRoute) {
	g.srv.AddNodeRoute(nodes...)
}

//Auth Authentication from the logic layer
func (g *GateBrokerHttp) Auth(token string) (interface{}, error) {
	url := g.logicAddr + "/auth"
	resp, body, errs := g.client.GetGoReq().Post(url).Send(fmt.Sprintf(`{ "token": "%v"}`, token)).End()
	if len(errs) != 0 {
		for i, err := range errs {
			logger.Info("logicBroker Auth failed. errs[%v]: %v ", i, err)
		}
		return nil, errs[0]
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("GateBroker Auth http failed code: %v", resp.StatusCode))
	}
	return json.RawMessage(body), nil
}

func (g *GateBrokerHttp) GetUserInfo() {
}

func (g *GateBrokerHttp) LoadInitData(uid string) (interface{}, error) {
	url := g.logicAddr + "/load"
	resp, body, errs := g.client.GetGoReq().Post(url).Send(fmt.Sprintf(`{ "user_id": "%v"}`, uid)).End()
	if len(errs) != 0 {
		for i, err := range errs {
			logger.Info("logicBroker Auth failed. errs[%v]: %v ", i, err)
		}
		return "", errs[0]
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("GateBroker Auth http failed code: %v", resp.StatusCode))
	}
	// forwards to user need raw json
	return json.RawMessage(body), nil
}

func (g *GateBrokerHttp) AddFriend(addFriendRequest interface{}) (interface{}, error) {
	url := g.logicAddr + "/addFriend"
	data, err := json.Marshal(addFriendRequest)
	if err != nil {
		return nil, err
	}
	resp, body, errs := g.client.GetGoReq().Post(url).Send(data).End()
	if len(errs) != 0 {
		for i, err := range errs {
			logger.Info("logicBroker AddFriend failed. errs[%v]: %v ", i, err)
		}
		return "", errs[0]
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("GateBroker AddFriend http failed code: %v", resp.StatusCode))
	}
	return json.RawMessage(body), nil
}

func (g *GateBrokerHttp) DeleteFriend(addFriendRequest interface{}) (interface{}, error) {
	url := g.logicAddr + "/addFriend"
	data, err := json.Marshal(addFriendRequest)
	if err != nil {
		return nil, err
	}
	resp, body, errs := g.client.GetGoReq().Post(url).Send(data).End()
	if len(errs) != 0 {
		for i, err := range errs {
			logger.Info("logicBroker AddFriend failed. errs[%v]: %v ", i, err)
		}
		return "", errs[0]
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("GateBroker AddFriend http failed code: %v", resp.StatusCode))
	}
	return json.RawMessage(body), nil
}

func (g *GateBrokerHttp) PingPong() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.String(http2.StatusOK, "Pong!")
	}
}
