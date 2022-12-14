package router

import (
    "github.com/gin-gonic/gin"
    "sdwan.ink/sdwan/api"
)


func InitRouter(group *gin.RouterGroup) (R gin.IRoutes){
    InitNode(group)
    InitSys(group)
    return group
}

func InitNode(group *gin.RouterGroup) (R gin.IRoutes){
    group = group.Group("node")
    {
        group.GET("test", apiNode.Test)
    }
    return group
}
func InitSys(group *gin.RouterGroup){
    group = group.Group("sys")
    {
        group.GET("test", apiNode.Test)
    }
}


var apiNode = api.WanApi.ApiNode