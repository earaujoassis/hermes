package web

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func listClientsHandler(c *gin.Context) {
    c.JSON(http.StatusOK, H{
        "_status":  "success",
        "clients": nil,
    })
}

func createClientHandler(c *gin.Context) {

}
