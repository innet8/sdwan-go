package main
import (
    "sdwan.ink/sdwan/global"
    "sdwan.ink/sdwan/boot"
)

func main(){
    global.WanViper = boot.Viper()
    global.WanLogger = boot.Zap()
	global.WanDb = boot.Mysql()
    if global.WanDb != nil {
        boot.Migrate(global.WanDb)
        db, _ := global.WanDb.DB()
		defer db.Close()
	}
    boot.Boot()
}