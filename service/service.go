package service

import (
	"ddns-go/api/aliApi"
	"ddns-go/api/tencentApi"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

var domain, subDomain, server string
var aliRecord string
var tencentRecord uint64

func InitRecord(value string) error {
	if aliRecord != "" || tencentRecord > 0 {
		return nil
	}
	domain = viper.GetString("domain")
	subDomain = viper.GetString("subDomain")
	server = viper.GetString("server")
	switch server {
	case "ali":
		records, err := aliApi.GetDomainRecordsByRRKeyWord(subDomain, domain)
		if err != nil {
			return err
		}
		for _, r := range records {
			if r.RR == subDomain {
				aliRecord = r.RecordID
				return nil
			}
		}
		aliRecord, err = aliApi.AddDomainRecord(domain, subDomain, "A", value)
		if err != nil {
			return err
		} else {
			return nil
		}

	case "tencent":
		records, err := tencentApi.ListDomainRecords(domain)
		if err != nil {
			return err
		}
		for _, r := range records {
			if *r.Name == subDomain {
				tencentRecord = *r.RecordId
				return nil
			}
		}
		tencentRecord, err = tencentApi.AddDomainRecord(domain, subDomain, "A", "默认", value)
		if err != nil {
			return err
		} else {
			return nil
		}

	}
	panic("配置有误")

}

func UpdateDomainRecord(value string) error {
	switch server {
	case "ali":
		err := aliApi.UpdateDomainRecord(aliRecord, subDomain, "A", value)
		if err != nil {
			return err
		}

	case "tencent":
		err := tencentApi.UpdateDomainRecord(domain, subDomain, "A", "默认", value, tencentRecord)
		if err != nil {
			return err
		}
	}
	return nil
}
func GetLocalIP() (string, error) {
	if viper.GetString("getPublicIPType") == "1" {
		return GetLocalIPByInterface()
	} else {
		return GetLocalIPByInternet()
	}
}
func GetLocalIPByInterface() (string, error) {
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
func GetLocalIPByInternet() (string, error) {
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
