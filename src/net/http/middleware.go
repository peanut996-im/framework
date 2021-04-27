package http

import (
	"bytes"
	"framework/api"
	"framework/cfgargs"
	"framework/logger"
	"framework/net"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckSign(cfg *cfgargs.SrvConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// j, err := net.ReadStringFromBody(c.Request.Body)
		b, err := c.GetRawData()
		if err != nil {
			logger.Error("get raw data err: %v", err)
			c.AbortWithStatusJSON(http.StatusOK, net.NewHttpInnerErrorResp(err))
		}
		// write data to body
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(b))

		body := string(b)

		logger.Debug("get body : %v", body)

		if err != nil || len(body) == 0 {
			checkResult, err := api.CheckSignFromQueryParams(c.Request.URL.Query(), cfg.AppKey)
			if !checkResult || err != nil {
				logger.Debug("check sign with query string failed")
				c.AbortWithStatusJSON(http.StatusOK, net.SignInvaildResp)
			}
		} else {
			checkResult, err := api.CheckSignFromJsonString(body, cfg.AppKey)
			if !checkResult || err != nil {
				logger.Debug("check sign with json failed")
				c.AbortWithStatusJSON(http.StatusOK, net.SignInvaildResp)
			}
		}
		c.Next()
	}
}
