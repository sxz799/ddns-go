package utils

import (
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/spf13/viper"
	"log"
)

var ApiClient *alidns20150109.Client

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
	log.Println(accessKeyId)
	log.Println(accessKeySecret)
	ApiClient, err = createClient(&accessKeyId, &accessKeySecret)
	if err != nil {
		panic(err)
	}
}

func createClient(accessKeyId *string, accessKeySecret *string) (_result *alidns20150109.Client, _err error) {
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
