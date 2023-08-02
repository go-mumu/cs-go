/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-28
 * File: login.go
 * Desc: 验证登陆
 */

package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-mumu/cs-go/api/container"
	"github.com/go-mumu/cs-go/log"
)

type Request struct {
	Token string `json:"token"`
}

type UserInfo struct {
	Mid       int64  `json:"mid"`
	XcxOpenId string `json:"xcx_openid"`
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		if authMiniProgramToken(c) {
			c.Next()
		} else {
			c.AbortWithStatusJSON(505, gin.H{
				"message": "params token is invalid.",
			})
		}
	}
}

func authMiniProgramToken(c *gin.Context) bool {
	token, err := getMiniProgramToken(c)
	if token == "" || err != nil {
		return false
	}

	response := container.Redis().Get(c.Request.Context(), token)
	if response.Val() == "" || response.Err() != nil {
		return false
	}

	userInfo := new(UserInfo)
	if err = json.Unmarshal([]byte(response.Val()), &userInfo); err != nil {
		return false
	}

	log.Log.Info("log", "user info", userInfo, "redis val", response.Val())

	c.Set("mid", userInfo.Mid)
	c.Set("xcx_openid", userInfo.XcxOpenId)

	return true
}

// 获取小程序token
func getMiniProgramToken(c *gin.Context) (string, error) {
	request := new(Request)

	if err := c.ShouldBindJSON(request); err != nil {
		log.Log.Warn("param is invalid.")
		return "", err
	}

	return request.Token, nil
}
