package node

import "github.com/gin-gonic/gin"

type NodeApi struct{}

func (r *NodeApi) InitRouterNodeApi(privateGroup *gin.RouterGroup) (R gin.IRoutes) {
    privateGroup = privateGroup.Group("api")
//    {
//        privateGroup.POST("get-api-list", apiPrivate.ApiApi.GetApiList)
//        privateGroup.POST("edit-api", apiPrivate.ApiApi.EditApi)
//        privateGroup.POST("add-api", apiPrivate.ApiApi.AddApi)
//        privateGroup.POST("delete-api-by-id", apiPrivate.ApiApi.DeleteApiById)
//        privateGroup.POST("query-api-by-id", apiPrivate.ApiApi.QueryApiById)
//    }
//    return privateGroup
}
