package socketio

import (
	"fmt"
	"framework/cfgargs"
	"framework/logger"
	"net/http"
	"sync"

	sio "github.com/googollee/go-socket.io"
)

var nsp = "/"

type Server struct {
	srv                *sio.Server
	handlers           map[string]interface{}
	SocketIOToSessions map[string]*Session
	UIDSceneToSessions map[string]*Session
	sync.Mutex
}

func NewSIOHandlers() map[string]interface{} {
	return make(map[string]interface{})
}

func NewServer() *Server {
	s := &Server{
		handlers:           make(map[string]interface{}),
		SocketIOToSessions: make(map[string]*Session),
		UIDSceneToSessions: make(map[string]*Session),
	}
	return s
}

func (s *Server) Run(cfg *cfgargs.SrvConfig) error {
	srv, err := sio.NewServer(nil)
	if err != nil {
		return err
	}
	s.srv = srv

	defer s.srv.Close()
	go s.srv.Serve() //nolint: errcheck

	if cfg.HTTP.Cors {
		http.HandleFunc("/socket.io/", func(w http.ResponseWriter, r *http.Request) {
			allowHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"
			if origin := r.Header.Get("Origin"); origin != "" {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Vary", "Origin")
				w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, DELETE")
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				w.Header().Set("Access-Control-Allow-Headers", allowHeaders)
			}
			if r.Method == "OPTIONS" {
				return
			}
			r.Header.Del("Origin")
			s.srv.ServeHTTP(w, r)
		})
	} else {
		http.Handle("/socket.io/", s.srv)
	}
	addr := fmt.Sprintf(":%v", cfg.SocketIO.Port)
	logger.Info("Serving at %v...", addr)

	err = http.ListenAndServe(addr, nil)
	logger.Fatal("Serving at %v... err:%v", addr, err)
	return err
}

func (s *Server) OnConnect(f func(sio.Conn) error) {
	s.srv.OnConnect(nsp, f)
}

func (s *Server) OnDisconnect(f func(sio.Conn, string)) {
	s.srv.OnDisconnect(nsp, f)
}

func (s *Server) OnError(f func(sio.Conn, error)) {
	s.srv.OnError(nsp, f)
}

func (s *Server) MountHandlers(handlers map[string]func(sio.Conn, ...interface{})) {
	for k, v := range handlers {
		s.srv.OnEvent(nsp, k, v)
	}
}

func (s *Server) SocketIOToSession(c sio.Conn) *Session {
	s.Lock()
	si, ok := s.SocketIOToSessions[c.ID()]
	s.Unlock()
	if !ok {
		logger.Warn("session not found")
		return nil
	}
	return si
}

func (s *Server) UIDSceneToSession(uidScene string) *Session {
	s.Lock()
	si, ok := s.UIDSceneToSessions[uidScene]
	s.Unlock()
	if !ok {
		logger.Warn("session not found")
		return nil
	}
	return si
}

func (s *Server) DisconnectSession(conn sio.Conn) *Session {

	s.Lock()
	si, ok := s.SocketIOToSessions[conn.ID()]
	if ok || nil != si {
		delete(s.SocketIOToSessions, si.Conn.ID())
	} else {
		logger.Warn("Sessions.DisconnSession[%v] not found", SocketIOToString(conn))
	}

	if nil != si {
		siScene, ok := s.UIDSceneToSessions[si.UIDSceneString()]
		if ok || nil != siScene {
			logger.Info("Sessions.DisconnSession,UIDAndScene:v%", si.UIDSceneString())
			delete(s.UIDSceneToSessions, si.UIDSceneString())
		}
	}

	s.Unlock()
	return si
}
