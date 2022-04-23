package model

/* import (
	"Deskmate/utils"
	"mime/multipart"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var ImgUrl = utils.QiniuSever


func UploadFile(file multipart.File,fileSize int64) (string,int) {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone: &storage.ZoneHuadon,
		UseCdnDomains: false,
		UseHTTPS:false,
	}

	putExtra := storage.PutExtra{} //额外参数不配置

	formUploader := storage.NewFormUploader{&cfg}
	ret := storage.PutRet{}

	err :=formUploader.PutWithoutKey(context.Background(),&ret,upToken,file,fileSize,&putExtra)
	if err != nil{
		return "",errmsg.ERROR
	}
	url :=ImgUrl + ret.AccessKey
	return url,errmsg.SUCCSE
} */