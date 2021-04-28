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
		b, err := c.GetRawData()
		if err != nil {
			logger.Error("get raw data err: %v", err)
			c.AbortWithStatusJSON(http.StatusOK, net.NewHttpInnerErrorResp(err))
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(b))

		body := string(b)

		if err != nil || len(body) == 0 {
			// query string
			checkResult, err := api.CheckSignFromQueryParams(c.Request.URL.Query(), cfg.AppKey)
			if !checkResult || err != nil {
				logger.Debug("check sign with query string failed")
				if !cfg.HTTP.Release {
					sign, err := api.MakeSignWithQueryParams(c.Request.URL.Query(), cfg.AppKey)
					if err == nil {
						c.AbortWithStatusJSON(http.StatusOK, net.NewBaseResponse(net.ERROR_SIGN_INVAILD, gin.H{"sign": sign}))
						return
					}
				}
				c.AbortWithStatusJSON(http.StatusOK, net.SignInvaildResp)
			}
		} else {
			// json
			checkResult, err := api.CheckSignFromJsonString(body, cfg.AppKey)
			if !checkResult || err != nil {
				logger.Debug("check sign with json failed")
				if !cfg.HTTP.Release {
					sign, err := api.MakeSignWithJsonString(body, cfg.AppKey)
					if err == nil {
						c.AbortWithStatusJSON(http.StatusOK, net.NewBaseResponse(net.ERROR_SIGN_INVAILD, gin.H{"sign": sign}))
					}
				}
				c.AbortWithStatusJSON(http.StatusOK, net.SignInvaildResp)
			}
		}
		c.Next()
	}
}
