package server

import "github.com/gin-gonic/gin"

type RouteGroup interface {
	RegisterRoutes(router *gin.Engine)
}
