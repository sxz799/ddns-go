package main

import (
	"ddns-go/ddns"
	"ddns-go/utils"
	"fmt"
	"github.com/spf13/viper"
	"time"
)

func main() {
	var lastLocalIP string

	for {
		localIP, err := utils.GetLocalIP()
		if err != nil {
			fmt.Println("获取本机IP失败！错误信息：", err)
			time.Sleep(time.Minute * time.Duration(viper.GetInt("interval")))
			continue
		}
		err = ddns.InitRecord(localIP)
		if err != nil {
			fmt.Println("获取解析记录失败！错误信息：", err)
			time.Sleep(time.Minute * time.Duration(viper.GetInt("interval")))
			continue
		}

		if lastLocalIP == "" || lastLocalIP != localIP {
			err = ddns.UpdateDomainRecord(localIP)
			if err != nil {
				fmt.Println("更新解析记录失败！错误信息：", err)
				time.Sleep(time.Minute * time.Duration(viper.GetInt("interval")))
			} else {
				lastLocalIP = localIP
				fmt.Println("更新完成,当前解析地址为:", localIP)
			}
			time.Sleep(time.Minute * time.Duration(viper.GetInt("interval")))
		} else {
			fmt.Println("本地IP未发生变化，无需更新！")
			time.Sleep(time.Minute * time.Duration(viper.GetInt("interval")))
		}

	}

}
