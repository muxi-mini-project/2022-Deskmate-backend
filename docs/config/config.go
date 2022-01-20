package config //按wz的格式来的，是用来配置config.yaml的文件

import (
	"log"

	"github.com/spf13/viper"
)

func ConfigInit() {
	viper.SetConfigFile("./conf/config.yaml") //这三行是来指定配置文件的路径
	viper.SetConfigName("config")
	viper.AddConfigPath("./conf") //解析默认的配置文件
	err := viper.ReadInConfig()   //按照路径读取yaml文件，之后可以直接viper.Get(mysql.host)这样从yaml中提取信息

	if err != nil {
		log.Fatal(err)
	}

}
