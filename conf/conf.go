package conf

import (
	"fmt"
	"go_bilibili/cache"
	"go_bilibili/model"
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/conf")
	fmt.Println("workDir: ", workDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	//数据库初始化
	cache.InitRedis()
	model.InitDB()
}
