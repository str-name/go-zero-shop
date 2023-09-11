package qiniu

import (
	"bytes"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
	"zero-shop/common/xerr"
)

// getToken获取上传token
func getToken(accessKey, secretKey, bucket string) string {
	putPolicy := storage.PutPolicy{Scope: bucket}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}

// 获取上传的配置
func getCfg(zoneID string) storage.Config {
	cfg := storage.Config{}

	// 空间对应的机房
	zone, _ := storage.GetRegionByID(storage.RegionID(zoneID))
	cfg.Zone = &zone
	// 是否用https域名
	cfg.UseHTTPS = true
	// 上传是否用CDN加速
	cfg.UseCdnDomains = false
	return cfg
}

func getBucketManager(accessKey, secretKey, zoneID string) *storage.BucketManager {
	mac := qbox.NewMac(accessKey, secretKey)

	cfg := getCfg(zoneID)
	bucketManager := storage.NewBucketManager(mac, &cfg)
	return bucketManager
}

// UploadImage 上传用户头像图片
func UploadImage(ctx context.Context, data []byte, imageName, imageData, prefix, accessKey, secretKey, bucket, zoneID, cdn string) (filePath string, err error) {
	if accessKey == "" || secretKey == "" || bucket == "" {
		return "", xerr.NewErrMsg("请配置accessKey、secretKey和bucket")
	}

	upToken := getToken(accessKey, secretKey, bucket)
	cfg := getCfg(zoneID)

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	dataLen := int64(len(data))

	key := fmt.Sprintf("%s_%s_%s", prefix, imageName, imageData)
	err = formUploader.Put(ctx, &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		return "", xerr.NewErrMsg("七牛云上传图片上失败")
	}
	return fmt.Sprintf("%s%s", cdn, ret.Key), nil
}

func DeleteImage(key, accessKey, secretKey, bucket, zoneID string) (err error) {
	bucketManager := getBucketManager(accessKey, secretKey, zoneID)
	err = bucketManager.Delete(bucket, key)
	if err != nil {
		logx.Errorf("七牛云删除图片失败， imageName：%s, err: %+v", key, err)
		return xerr.NewErrMsg("七牛云删除图片失败")
	}
	return
}

func DeleteImageList(keyList string, accessKey, secretKey, bucket, zoneID string) (err error) {
	list := strings.Split(keyList, ",")
	for _, s := range list {
		kList := strings.Split(s, "/")
		key := kList[len(kList)-1]
		err = DeleteImage(key, accessKey, secretKey, bucket, zoneID)
		if err != nil {
			return err
		}
	}
	return
}
