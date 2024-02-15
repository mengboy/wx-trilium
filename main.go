package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mengboy/wx-trilium/wxmsg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	config   string
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "start service",
		Run:   start,
	}
)

var RootCmd = &cobra.Command{
	Use:   "wx-trilium",
	Short: "service",
}

func init() {
	startCmd.PersistentFlags().StringVarP(&config, "config", "c", "", "加载指定配置文件")
	_ = viper.BindPFlags(startCmd.PersistentFlags())
	RootCmd.AddCommand(startCmd)
}

func parseConfig() {
	if config == "" {
		config = fmt.Sprintf("%s/conf.toml", os.Getenv("WX_SUB_PATH"))
	}
	viper.SetConfigFile(config)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("parse config file fail: %s \n You can set env KOI_PATH to load config file", err))
	}
}

func start(_ *cobra.Command, _ []string) {
	parseConfig()
	gin.SetMode(viper.GetString("app.runmod"))
	r := gin.Default()
	r.GET("/wxbg", wxmsg.CheckSignature)
	r.POST("/wxbg", wxmsg.Receive)
	r.Run("0.0.0.0:1234")
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
