package http

import (
	"fmt"
	"framework/cfgargs"
	"framework/net"
	"github.com/gin-gonic/gin"
)

type Route struct {
	method  string
	path    string
	handler gin.HandlerFunc
}

type NodeRoute struct {
	path   string
	routes []*Route
}

// VerRoute distinguish routes with version.
// type VerRoute struct {
// 	Version   string
// 	DevRoutes []DevRoute
// }

//Server ...
type Server struct {
	session *gin.Engine
	routers []*NodeRoute
}

// NewServer ...
func NewServer(cfg *cfgargs.SrvConfig) *Server {
	if cfg.HTTP.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	return &Server{
		session: gin.Default(),
		routers: []*NodeRoute{},
	}
}

func NewRoute(method, path string, handler gin.HandlerFunc) *Route {
	return &Route{
		method:  method,
		path:    path,
		handler: handler,
	}
}

func NewNodeRoute(path string, routers ...*Route) *NodeRoute {
	return &NodeRoute{
		path:   path,
		routes: routers,
	}
}
func (s *Server) AddNodeRoute(nodes ...*NodeRoute) {
	s.routers = append(s.routers, nodes...)
}

//Serve ...
func (s *Server) Serve(cfg *cfgargs.SrvConfig) error {
	if cfg.HTTP.Cors {
		s.session.Use(CORS())
	}
	if cfg.HTTP.Sign {
		s.session.Use(CheckSign(cfg))
	}
	s.mountRoutes()
	err := s.session.Run(fmt.Sprintf(":%v", cfg.HTTP.Port))
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Use(middlewares ...gin.HandlerFunc) {
	s.session.Use(middlewares...)
}

func (s *Server) mountRoutes() {
	router := s.session.Group("/")
	for _, node := range s.routers {
		group := router.Group(node.path)
		for _, route := range node.routes {
			methodMapper(group, route.method)(route.path, route.handler)
		}
	}
}

func methodMapper(group *gin.RouterGroup, method string) func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	switch method {
	case net.HTTP_METHOD_GET:
		return group.GET
	case net.HTTP_METHOD_POST:
		return group.POST
	case net.HTTP_METHOD_PUT:
		return group.PUT
	case net.HTTP_METHOD_DELETE:
		return group.DELETE
	case net.HTTP_METHOD_PATCH:
		return group.PATCH
	default:
		return group.Any
	}
}
