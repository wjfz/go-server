package cmd

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"wserver/pkg/rlogging"
	"wserver/routers"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve short comment",
	Long:  `serve long comment`,
	Run: func(cmd *cobra.Command, args []string) {
		port := viper.GetString("port")

		// 设置gin运行模式
		setGinMode()

		log.Println("Run gin on mode:", gin.Mode())
		log.Println("Listening and serving HTTP on:", port)

		// 初始化gin
		router := gin.New()
		router.Use(ginLogger())

		// 引入路由
		routers.InitRouter(router)

		// 启动http server服务
		server := &http.Server{
			Addr:           port,
			Handler:        router,
			ReadTimeout:    time.Second * 3,
			WriteTimeout:   time.Second * 3,
		}

		go func() {
			if err := server.ListenAndServe(); err != nil {
				log.Printf("Server Listen: %s\n", err)
			}
		}()

		// 平滑关闭http server
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt, os.Kill)
		<-quit

		log.Println("收到server停止信号，等待现有连接关闭……")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Println("Server Shutdown err:", err)
		}

		log.Println("Server 服务已停止")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func setGinMode() {
	switch viper.GetInt("debug") {
	case 1:
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
}

func ginLogger() gin.HandlerFunc {
	config := gin.LoggerConfig{
		Output: rlogging.NewRotateWriter("request", rlogging.RotateDay),
		Formatter: func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("%v | %-4s | %s | %3d | %13v | %15s %s\n",
				param.TimeStamp.Format("2006/01/02 15:04:05"),
				param.Method,
				param.Path,
				param.StatusCode,
				param.Latency,
				param.ClientIP,
				param.ErrorMessage,
			)
		},
	}
	return gin.LoggerWithConfig(config)
}