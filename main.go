package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gover-server/common"
	"gover-server/router"
	"gover-server/service"
	"log"
	"os"
)

func main() {

	//初始化读取配置
	initConfig()
	// 初始化日志文件
	initLog()
	db := common.InitDB()
	defer db.Close()

	// 初始化管理员
	service.InitAdmin()

	r := gin.Default()
	r = router.CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		//取代默认端口
		panic(r.Run(":" + port))
	}
	panic(r.Run())

}

// 数据库的配置文件
func initConfig() {
	workDir, _ := os.Getwd()
	//设置要读取的文件名
	viper.SetConfigName("application")
	//设置读取文件的类型
	viper.SetConfigType("yml")
	//设置文件的路径
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("读取配置文件失败")
	}
}

// 初始化日志文件
func initLog() {
	file := viper.GetString("logfile.path")
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil && os.IsNotExist(err) {
		logFile, _ = os.Create(file)
	} else if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[DEBUG]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	return
}
