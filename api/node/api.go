package node

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

type ApiApi struct{}

func (a *ApiApi) GetApiList(c *gin.Context) {
    fmt.Print("tt")
}