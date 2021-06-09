// Package gate
// @Title  logicbroker_http.go
// @Description
// @Author  peanut996
// @Update  peanut996  2021/5/31 11:40
package broker

import (
	"encoding/json"
	"errors"
	"fmt"
	"framework/cfgargs"
	"framework/logger"
	"framework/net/http"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

type LogicBrokerHttp struct {
	srv          *http.Server
	client       *http.Client
	gateAddr     string
	gateToScenes map[string][]string
	sceneToGate  map[string]string
	sync.Mutex
}

func NewLogicBrokerHttp() *LogicBrokerHttp {
	return &LogicBrokerHttp{}
}

func (l *LogicBrokerHttp) Init(cfg *cfgargs.SrvConfig) {
	l.srv = http.NewServer()
	l.srv.Init(cfg)
	l.client = http.NewClient()
	l.gateAddr = fmt.Sprintf("http://%v:%v", cfg.Gate.Host, cfg.Gate.Port)
	if cfg.Gate.Mode != "http" {
		logger.Warn("can't load gate configuration for http")
		if cfg.Gate.Panic {
			panic(errors.New("can't load gate configuration for http"))
		}
		return
	}
}

func (l *LogicBrokerHttp) Listen() {
	//启动http server
	l.srv.Run()
}

//Invoke Push data to gate
func (l *LogicBrokerHttp) Invoke(packet interface{}) (interface{}, error) {
	// FIXME With target, event
	// default http address
	addr := l.gateAddr + "/"
	resp, body, errs := l.client.GetGoReq().Post(addr).Send(packet).End()
	if len(errs) != 0 {
		for i, err := range errs {
			logger.Info("LogicBroker Event /%v failed. errs[%v]: %v ", i, err)
		}
		return nil, errs[0]
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("LogicBroker Event /%v http failed code: %v", resp.StatusCode))
	}
	return json.RawMessage(body), nil
}

// InvokeTarget Send commands to the gate instance
func (l *LogicBrokerHttp) InvokeTarget(target, event string, data interface{}) {
	l.Lock()
	addr, ok := l.sceneToGate[target]
	if ok {
		go l.client.GetGoReq().Post(addr + event).Send(data).End()
	}
}

// AddNodeRoute Mount the route to the internal HTTP server.
func (l *LogicBrokerHttp) AddNodeRoute(nodes ...*http.NodeRoute) {
	l.srv.AddNodeRoute(nodes...)
}

func (l *LogicBrokerHttp) handleRegister(c *gin.Context) {
	data := make(map[string]interface{})
	c.BindJSON(&data)
	l.Unlock()
	_, ok := l.gateToScenes[data["ip"].(string)]
	if !ok {
		l.gateToScenes[data["ip"].(string)] = []string{"1", "2"}
	}
	l.Lock()
}

func (l *LogicBrokerHttp) handleUpdate(c *gin.Context) {
	data := make(map[string]interface{})
	c.BindJSON(&data)
	l.Unlock()
	scenes, ok := l.gateToScenes[data["ip"].(string)]
	if !ok {
		l.gateToScenes[data["ip"].(string)] = data["scenes"].([]string)
	} else {
		scenes = data["scenes"].([]string)
		l.gateToScenes[data["ip"].(string)] = scenes
		for _, scene := range scenes {
			l.sceneToGate[scene] = data["ip"].(string)
		}
	}
	l.Lock()
}

func (l *LogicBrokerHttp) loopDoKeepAlive() {
	for {
		for addr := range l.gateToScenes {
			go func() {
				isAlive := false
				resp, _, err := l.client.GetGoReq().Post(fmt.Sprintf("%v:%v", addr, "logic/ping")).End()
				if err != nil || resp.StatusCode != 200 {
					// start retry
					for i := 0; i < 3; i++ {
						resp, _, err := l.client.GetGoReq().Post(fmt.Sprintf("%v:%v", addr, "logic/ping")).End()
						if err == nil && resp.StatusCode == 200 {
							isAlive = true
							break
						}
					}
				} else {
					isAlive = true
				}
				if !isAlive {
					l.Lock()
					scenes, ok := l.gateToScenes[addr]
					if ok {
						for _, scene := range scenes {
							delete(l.sceneToGate, scene)
						}
					}
					delete(l.gateToScenes, addr)

				}

			}()
		}
		time.After(500 * time.Millisecond)
	}

}
