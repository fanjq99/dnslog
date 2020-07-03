package api

import (
	"time"

	"github.com/gin-gonic/gin"
)

func (h *HttpServer) Index(c *gin.Context) {
	body := `<h1>d4f800167a6e317f35454ed9024ebd420</h1>">`
	c.Data(200, "text/html; charset=utf-8", []byte(body))
}

func (h *HttpServer) Status(c *gin.Context) {
	c.Data(200, "text/html; charset=utf-8", []byte(""))
}

func (h *HttpServer) PHPInclude(c *gin.Context) {
	body := `<?php  echo md5("websec");`
	c.Data(200, "text/html; charset=utf-8", []byte(body))
}

func (h *HttpServer) Xss(c *gin.Context) {
	body := `<script>prompt(98589956)</script>`
	c.Data(200, "text/html; charset=utf-8", []byte(body))
}

func (h *HttpServer) VulVerifyHttp(c *gin.Context) {
	_, exist := c.GetQuery("verify")
	rmd := c.Query("rmd")
	rmdLen := len(rmd)
	if rmdLen > 1 && rmdLen < 100 {
		if exist {
			responseContent := "is fail"
			res, err := h.redisClient.Get(rmd).Result()
			if err == nil && res != "" {
				responseContent = "Vulnerabilities exist " + res
			}
			c.Data(200, "text/html; charset=utf-8", []byte(responseContent))
		} else {
			h.redisClient.Set(rmd, c.ClientIP(), time.Duration(h.saveTime)*time.Second)
			c.Data(200, "text/html; charset=utf-8", []byte("success"))
		}
	} else {
		c.Data(200, "text/html; charset=utf-8", []byte("this is an illegal request"))
	}
}

func (h *HttpServer) VulVerifyDNS(c *gin.Context) {
	rmd := c.Query("rmd")
	rmdLen := len(rmd)
	if rmdLen > 1 && rmdLen < 100 {
		responseContent := "is fail"
		res, err := h.redisClient.Get(rmd).Result()
		if err == nil && res != "" {
			responseContent = "Vulnerabilities exist " + res
		}
		c.Data(200, "text/html; charset=utf-8", []byte(responseContent))
	} else {
		c.Data(200, "text/html; charset=utf-8", []byte("this is an illegal request"))
	}
}
