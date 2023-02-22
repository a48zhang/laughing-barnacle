package qiniuStorage

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type QNConfig struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Bucket    string `json:"bucket_name"`
	Domain    string `json:"domain_name"`
}

var Conf QNConfig

func Connect() {
	file, err := os.Open("./Conf/qn.json")
	if err != nil {
		Conf = QNConfig{
			AccessKey: os.Getenv("access_key"),
			SecretKey: os.Getenv("secret_key"),
			Bucket:    os.Getenv("bucket_name"),
			Domain:    os.Getenv("domain_name"),
		}
		return
	} else {
		tmp, _ := io.ReadAll(file)
		err = json.Unmarshal(tmp, &Conf)
	}
	if err != nil {
		log.Fatal("Failed to connect to qiniu cloud. Error:" + err.Error())
	}
}
