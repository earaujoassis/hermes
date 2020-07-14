package web

import (
    "strings"
    "net/http"
    "fmt"

    "github.com/gin-gonic/gin"

    "github.com/earaujoassis/hermes/server/config"
)

func SetupWeb() {
    router := gin.Default()
    exposeRoutes(router)
    router.Use(func(c *gin.Context) {
        defer func(c *gin.Context) {
            // TODO It is not displaying/logging the error
            if rec := recover(); rec != nil {
                if path := c.Request.URL.Path; strings.HasPrefix(path, "/api") {
                    c.JSON(http.StatusInternalServerError, H{
                        "_status":  "error",
                        "_message": "Bad server",
                        "error": "The server found an error; aborting",
                    })
                } else {
                    c.String(http.StatusInternalServerError, "error.internal")
                }
            }
        }(c)
        c.Next()
    })
    router.NoRoute(func(c *gin.Context) {
        if path := c.Request.URL.Path; strings.HasPrefix(path, "/api") {
            c.JSON(http.StatusNotFound, H{
                "_status":  "error",
                "_message": "Not found",
                "error": "Resource path not found",
            })
        } else {
            c.String(http.StatusNotFound, "error.not_found")
        }
    })
    router.Run(fmt.Sprintf(":%v", config.GetEnvVarDefault("PORT", "8080")))
}
