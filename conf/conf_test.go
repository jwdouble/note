package conf

import (
	"fmt"
	"os"
	"testing"

	"arctron.lib/conf"
	"github.com/spf13/viper"
)

func Test_conf(t *testing.T) {
	dsnDefault := "host=10.243.11.35 user=postgres password=p1ssw0rd dbname=arc-util port=5432 sslmode=disable TimeZone=Asia/Shanghai statement_cache_mode=describe"
	fmt.Println(conf.APP_PG_DSN.Value(dsnDefault))
}

func Test_ViperConf(t *testing.T) {
	// 方法1 直接设置文件详情
	//viper.SetConfigFile("../config.yaml")

	// 方法2 设置文件名称（不带类型），设置文件类型，添加文件路径
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("..")     // optionally look for config in the working directory
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	r2 := viper.Get("app.engine.log")
	r3 := viper.Get("test")
	fmt.Println(r2, r3)
}

func Test_viperEnv(t *testing.T) {
	viper.SetEnvPrefix("k8s") // will be uppercased automatically
	viper.BindEnv("ID2")

	os.Setenv("K8S_ID2", "13003") // typically done outside of the app

	id := viper.GetString("ID2") // 13
	fmt.Println(id)
}
