package web

import (
    "github.com/gin-gonic/gin"
)

func exposeRoutes(router *gin.Engine) {
    restApi := router.Group("/api")
    clientsRoutes := restApi.Group("/clients")
    clientsRoutes.POST("/", createClientHandler)
    clientsRoutes.GET("/", listClientsHandler)
}
