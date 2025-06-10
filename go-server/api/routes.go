package api

import "github.com/gin-gonic/gin"

func ApiRoutes(router *gin.Engine) {
	group := router.Group("/api")
	{
		group.POST("/embed", EmbedReq)
		group.GET("/search", SearchReq)
	}
}
