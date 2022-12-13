package boot

import (
    "sdwan.ink/sdwan/config"
    "sdwan.ink/sdwan/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Mysql() *gorm.DB {
	wanMysqlConfig := global.WanConfig.Mysql
    if wanMysqlConfig.Database == "" {
		return nil
	}
    mysqlConfig := config.MysqlConfig(wanMysqlConfig)
	gormConfig := config.GormConfig()
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gormConfig); err != nil {
		return nil
	} else {
		mysqlDb, _ := db.DB()
        mysqlDb.SetMaxIdleConns(wanMysqlConfig.MaxIdle)
        mysqlDb.SetMaxOpenConns(wanMysqlConfig.MaxOpen)
		return db
    }
}
