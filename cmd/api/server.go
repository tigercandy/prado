package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/tigercandy/prado/global"
	"github.com/tigercandy/prado/global/orm"
	"github.com/tigercandy/prado/initialize/config"
	"github.com/tigercandy/prado/initialize/database"
	"github.com/tigercandy/prado/internal/pkg/logger"
	"github.com/tigercandy/prado/internal/router"
	"github.com/tigercandy/prado/pkg/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	conf     string
	port     string
	StartCmd = &cobra.Command{
		Use:     "server",
		Short:   "Start server",
		Example: "prado server configs/config.yaml",
		PreRun: func(cmd *cobra.Command, args []string) {
			usage()
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func usage() {
	usageStr := `starting server`
	log.Printf("%s\n", usageStr)
}

func setup() {
	// 初始化配置文件
	config.Setup(conf)
	// 初始化日志
	global.App.Log = logger.Init()
	// 初始化数据库
	orm.Eloquent = database.Setup()
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&conf, "config", "c", "configs/config.yaml", "Start server with input config file")
	StartCmd.PersistentFlags().StringVarP(&port, "port", "p", "8002", "Server listen port")
}

func run() error {
	if global.App.Config.App.Env == string(utils.ModeProd) {
		gin.SetMode(gin.ReleaseMode)
	}
	r := router.InitRouter()

	// 程序运行结束前关闭数据库连接
	defer func() {
		if orm.Eloquent != nil {
			db, _ := orm.Eloquent.DB()
			_ = db.Close()
		}
	}()

	srv := &http.Server{
		Addr:    global.App.Config.App.Host + ":" + global.App.Config.App.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
		}
	}()
	// 优雅关闭服务
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Printf("%s Enter Ctrl+C Shutdown server... \r\n", utils.GetCurrentTimeStr())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown:", err)
	}
	logger.Info("Server exiting...")
	return nil
}
