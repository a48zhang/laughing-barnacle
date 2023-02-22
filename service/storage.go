package service

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"main/model/qiniuStorage"
	"mime/multipart"
)

var Conf = qiniuStorage.Conf

// UploadPhoto
// Upload a photo to Qiniu Cloud Storage.
// Return: Url of the photo
func UploadPhoto(file *multipart.File, size int64, typ string) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope:        Conf.Bucket,
		SaveKey:      "shmily/" + typ + "/${year}_${mon}_${day}_${hour}_${min}_${sec}.jpg",
		ForceSaveKey: true,
	}

	mac := qbox.NewMac(Conf.AccessKey, Conf.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	formUploader := storage.NewFormUploader(&storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	})

	putExtra := new(storage.PutExtra)
	ret := new(storage.PutRet)

	if err := formUploader.Put(context.Background(), ret,
		upToken, "", *file, size, putExtra); err != nil {
		return "", err
	}

	return Conf.Domain + "/" + ret.Key, nil
}
