package boot

import (
	"context"
	"fmt"
	"sdwan.ink/sdwan/global"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Boot() {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", global.WanConfig.System.Port),
		Handler: Router(),
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("【sdwan】监听：%s\n", err)
		}
	}()
	logo(fmt.Sprintf(":%d", global.WanConfig.System.Port))
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("【sdwan】开始退出...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("【sdwan】强制退出：", err)
	}

    log.Println("【sdwan】退出完成！")
}

var logoText = "  ________.__                 ________                                                 _____       .___      .__        \n /  _____/|__| ____           \\_____  \\  __ _______    ___________ _______            /  _  \\    __| _/_____ |__| ____  \n/   \\  ___|  |/    \\   ______  /  / \\  \\|  |  \\__  \\  /  ___/\\__  \\\\_  __ \\  ______  /  /_\\  \\  / __ |/     \\|  |/    \\ \n\\    \\_\\  \\  |   |  \\ /_____/ /   \\_/.  \\  |  // __ \\_\\___ \\  / __ \\|  | \\/ /_____/ /    |    \\/ /_/ |  Y Y  \\  |   |  \\\n \\______  /__|___|  /         \\_____\\ \\_/____/(____  /____  >(____  /__|            \\____|__  /\\____ |__|_|  /__|___|  /\n        \\/        \\/                 \\__>          \\/     \\/      \\/                        \\/      \\/     \\/        \\/"

func logo(port string) {
	fmt.Println(logoText)
	fmt.Println("欢迎使用sdwan")
	fmt.Println("系统已启动,开始监听 " + port + " 端口...")
    global.WanLogger.Info("欢迎使用sdwan,系统已启动,开始监听 " + port + " 端口...")
}
