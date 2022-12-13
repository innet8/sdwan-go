package boot

import (
    "sdwan.ink/sdwan/middleware"
    "github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
    var r = gin.Default()
    r.Use(middleware.Cors())
    return r
}


