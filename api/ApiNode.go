package api

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

type ApiNode struct{}

func (a *ApiNode) Test(c *gin.Context) {
    fmt.Println("接收到test了")

}
