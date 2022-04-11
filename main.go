package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jsmzr/bootstrap-config/config"
	"github.com/jsmzr/bootstrap-log/log"
	_ "github.com/jsmzr/bootstrap-plugin-config-yaml/yaml"
	_ "github.com/jsmzr/bootstrap-plugin-logrus/logrus"
	"github.com/jsmzr/bootstrap-plugin/plugin"
)

type GinProperties struct {
	Port        int    `default:"8080"`
	ContextPath string `default:"/"`
}

func main() {
	if err := plugin.PostProccess(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Info("gin-bootstrap 启动")
	var properties GinProperties
	if err := config.Resolve("boostrap.gin", &properties); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		log.Info("do something")
		ctx.JSON(200, gin.H{"message": "pong"})
	})
	// TODO gin.logger 与 logrus 的使用，out 的指定
	// TODO 其他配置项的设置
	r.Run(":" + strconv.Itoa(properties.Port))
}
