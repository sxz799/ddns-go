package utils

import (
	"errors"
	"fmt"
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/spf13/viper"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

var AliApiClient *alidns20150109.Client
var TencentClient *dnspod.Client

func init() {
	if _, err := os.Stat("conf/conf.yaml"); os.IsNotExist(err) {
		viper.SetConfigFile("conf.yaml")
	} else {
		viper.SetConfigFile("conf/conf.yaml")
	}
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
	case "ali":
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

func GetLocalIP() (string, error) {
	if viper.GetString("getPublicIPType") == "1" {
		return getLocalIPByInterface()
	} else {
		return getLocalIPByInternet()
	}
}
func getLocalIPByInterface() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		fmt.Println(iface.Name)
		if strings.ToLower(iface.Name) == "wan" {
			addrs, err := iface.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				if strings.Contains(addr.String(), "/") {
					addrStr := addr.String()[:strings.Index(addr.String(), "/")]
					if isPublicIP(net.ParseIP(addrStr)) {
						return addr.String(), nil
					}
				}

			}
		}
	}
	return "", errors.New("未获取到公网IP")
}
func getLocalIPByInternet() (string, error) {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	return string(ip), nil
}
func isPublicIP(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsLinkLocalMulticast() || ip.IsLinkLocalUnicast() {
		return false
	}
	if ipv4 := ip.To4(); ipv4 != nil {
		// 10.0.0.0/8
		if ipv4[0] == 10 {
			return false
		}
		// 172.16.0.0/12
		if ipv4[0] == 172 && ipv4[1] >= 16 && ipv4[1] <= 31 {
			return false
		}
		// 192.168.0.0/16
		if ipv4[0] == 192 && ipv4[1] == 168 {
			return false
		}
	}
	// 169.254.0.0/16
	if ip[0] == 169 && ip[1] == 254 {
		return false
	}
	// 240.0.0.0/4
	if ip[0] >= 240 {
		return false
	}
	// Otherwise, it's a public IP
	return true
}
