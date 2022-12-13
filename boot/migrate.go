package boot

import (
    "fmt"
    "sdwan.ink/sdwan/global"
    "go.uber.org/zap"
    "gorm.io/gorm"
    "os"
)

func Migrate(db *gorm.DB) {
    err := db.AutoMigrate()
    if err != nil {
        fmt.Println("迁移 sdwan 数据库失败！", zap.Any("error:", err))
        global.WanLogger.Error("迁移 sdwan 数据库失败！", zap.Any("error:", err))
        os.Exit(0)
    }
    fmt.Println("迁移 sdwan 数据库成功！")
    global.WanLogger.Info("迁移 sdwan 数据库成功！")
}
