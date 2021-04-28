package socketio

import (
	"fmt"
	"framework/api"
	"framework/logger"

	sio "github.com/googollee/go-socket.io"
)

type Session struct {
	Conn  sio.Conn
	token string
	uid   int64
	scene string
}

func (s *Session) SetScene(scene string) {
	s.scene = scene
}

func (s *Session) ID() string {
	return s.Conn.ID()
}

func SocketIOToString(c sio.Conn) string {
	if nil != c {
		id := c.ID()
		localAddr := c.LocalAddr()
		remoteAddr := c.RemoteAddr()
		return fmt.Sprintf("ID:%v addr.local:%v addr.remote:%v", id, localAddr, remoteAddr)
	}
	return "conn not found"
}

func (s *Session) UIDSceneString() string {
	return fmt.Sprintf("uid:%v_scene:%v", s.uid, s.scene)
}

func (s *Session) Auth(token string) (bool, error) {
	resp, err := api.CheckUserToken(token)
	if err != nil {
		logger.Info("check user token failed error: %v", err)
		return false, err
	}
	s.uid = resp.UID
	s.token = token
	return true, nil
}

func (s *Session) ToString() string {
	return SocketIOToString(s.Conn)
}
