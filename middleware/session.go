package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Session 初始化session
func Session(secret string,name string) gin.HandlerFunc {
	store := cookie.NewStore([]byte(secret))
	//Also set Secure: true if using SSL, you should though
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 7 * 86400, Path: "/"})
	return sessions.Sessions(name, store)
}
