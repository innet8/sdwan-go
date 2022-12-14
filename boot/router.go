package boot

import (
    "sdwan.ink/sdwan/middleware"
    "github.com/gin-gonic/gin"
    "sdwan.ink/sdwan/router"
)

func Router() *gin.Engine {
    var r = gin.Default()
    r.Use(middleware.Cors())

    group := r.Group("api/v1")
    router.InitRouter(group)
//    InitRouter(group)
    return r
}
//
//func InitRouter(group *gin.RouterGroup){
////    node := router.WanRouter.Node
////    {
////        node.InitNode(group)
////    }
//    router.InitNode(group)
//}


