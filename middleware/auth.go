package middleware

import (
	"bluebull/JWT/Token"
	"bluebull/respond"
	"github.com/gin-gonic/gin"
	"strings"
)

func TokenAuthMiddle() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定

		//判读是否有token
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			respond.Fail(c, respond.CodeEmptyAuth)
			c.Abort()
			return
		}

		//判断token格式是否正确
		//注意：sep那里是要有空格的  token是用空格分开的
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			respond.Fail(c, respond.CodeAuthFormatInvalid)
			c.Abort()
			return
		}

		//解析token
		msg, err := Token.ParseToken(parts[1])
		if err != nil {
			respond.Fail(c, respond.CodeTokenInvalid)
			c.Abort()
			return
		}

		//把token插入之后可以随时调用
		c.Set("department", msg.Department)
		c.Next()
	}
}
