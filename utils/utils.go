package utils

import (
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/spf13/viper"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	"log"
	"os"
)

var AliApiClient *alidns20150109.Client
var TencentClient *dnspod.Client

func init() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic("viper load fail ...")
	}
	accessKeyId := viper.GetString("key.accessKeyId")
	accessKeySecret := viper.GetString("key.accessKeySecret")
	if accessKeyId == "" || accessKeySecret == "" {
		log.Println("请填写密钥信息！")
		os.Exit(0)
	}
	server := viper.GetString("server")
	switch server {
	case "aliModel":
		AliApiClient, err = createAliClient(&accessKeyId, &accessKeySecret)
	case "tencent":
		TencentClient, err = createTencentClient(accessKeyId, accessKeySecret)
	}
	if err != nil {
		panic(err)
	}
}

func createAliClient(accessKeyId *string, accessKeySecret *string) (_result *alidns20150109.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("alidns.cn-hangzhou.aliyuncs.com")
	_result = &alidns20150109.Client{}
	_result, _err = alidns20150109.NewClient(config)
	return _result, _err
}

func createTencentClient(accessKeyId string, accessKeySecret string) (*dnspod.Client, error) {
	credential := common.NewCredential(
		accessKeyId,
		accessKeySecret,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	client, err := dnspod.NewClient(credential, "", cpf)
	return client, err
}
