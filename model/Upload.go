package model

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
	"myBlog/utils"
	"myBlog/utils/errmsg"
)

var Zone = utils.Zone
var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var ImgUrl = utils.QiniuSever

func UpLoadFile(file multipart.File, fileSize int64) (string, int) {
	/*putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	// UploadToken 方法用来进行上传凭证的生成
	// 该方法生成的过期时间是现对于现在的时间
	upToken := putPolicy.UploadToken(mac)
	cfg := setConfig()
	//可选配置
	putExtra := storage.PutExtra{}
	//表单验证
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", errmsg.ERROR
	}
	url := ImgUrl + ret.Key
	return url, errmsg.SUCCSE*/
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)
	region, _ := storage.GetRegion(AccessKey, Bucket)
	cfg := storage.Config{
		Region:        region,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := formUploader.Put(context.Background(), &ret, upToken, "key", file, fileSize, &putExtra)
	if err != nil {
		return "", errmsg.ERROR
	}
	url := ImgUrl + "/" + ret.Key
	return url, errmsg.SUCCSE
}

func setConfig() storage.Config {
	cfg := storage.Config{
		Zone:          selectZone(Zone),
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	return cfg
}

func selectZone(id int) *storage.Zone {
	switch id {
	case 1:
		return &storage.ZoneHuadong
	case 2:
		return &storage.ZoneHuabei
	case 3:
		return &storage.ZoneHuanan
	default:
		return &storage.ZoneHuadong
	}
}
