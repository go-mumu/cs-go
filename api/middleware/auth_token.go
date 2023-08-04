/**
 * Created GoLand.
 * User: liyw<634482545@qq.com>
 * Date: 2023-07-28
 * File: login.go
 * Desc: auth token middleware
 */

package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-mumu/cs-go/library/log"
	"github.com/redis/go-redis/v9"
)

type Request struct {
	Token string `json:"token"`
}

type UserInfo struct {
	Mid       int64  `json:"mid"`
	XcxOpenId string `json:"xcx_openid"`
}

func AuthToken(client *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		if authMiniProgramToken(c, client) {
			c.Next()
		} else {
			c.AbortWithStatusJSON(505, gin.H{
				"message": "params token is invalid.",
			})
		}
	}
}

func authMiniProgramToken(c *gin.Context, client *redis.Client) bool {
	token, err := getMiniProgramToken(c)
	if token == "" || err != nil {
		return false
	}

	response := client.Get(c.Request.Context(), token)
	if response.Val() == "" || response.Err() != nil {
		return false
	}

	userInfo := new(UserInfo)
	if err = json.Unmarshal([]byte(response.Val()), &userInfo); err != nil {
		return false
	}

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
