package utils

import (
	"gopkg.in/ini.v1"
	"fmt"
)

var (
	AccessKey string
	SecretKey string
	Bucket string
	QiniuSever string
)

func init() {
	file,err :=ini.Load("/config/config.ini") //Load方法返回一个file指针
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径；",err)
	}
	LoadQiniu(file)
}

func LoadServer(file *ini.File) {

}

func LoadData(file *ini.File) {

}

func LoadQiniu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuSever = file.Section("qiniu").Key("QiniuSever").String()
}