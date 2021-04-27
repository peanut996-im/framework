package http

import (
	"framework/api"
	"framework/cfgargs"
	"framework/logger"
	"framework/net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckSign(cfg *cfgargs.SrvConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		j, err := net.ReadFromBody(c.Request.Body)

		if err != nil || len(j) == 0 {
			checkResult, err := api.CheckSignFromQueryParams(c.Request.URL.Query(), cfg.AppKey)
			if !checkResult || err != nil {
				logger.Debug("check sign with query string failed")
				c.AbortWithStatusJSON(http.StatusOK, net.SignInvaildResp)
			}
		} else {
			checkResult, err := api.CheckSignFromJsonString(j, cfg.AppKey)
			if !checkResult || err != nil {
				logger.Debug("check sign with json failed")
				c.AbortWithStatusJSON(http.StatusOK, net.SignInvaildResp)
			}
		}
		c.Next()
	}
}
