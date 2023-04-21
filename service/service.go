package service

import (
	"ddns-go/api"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

func InitRecord() (string, error) {
	domain := viper.GetString("domain")
	rr := viper.GetString("rr")
	records, err := api.GetDomainRecordsByRRKeyWord(rr, domain)
	if err != nil {
		panic(err)
	}
	if len(records.Record) < 1 {
		record, err := api.AddDomainRecord(domain, rr, "A", "1.1.1.1")
		if err != nil {
			return "", err
		} else {
			return record.RecordID, nil
		}
	} else {
		for _, r := range records.Record {
			if r.RR == rr {
				return r.RecordID, nil
			}
		}
		return "", errors.New("未在存在的解析记录中找到指定内容")
	}
}
func GetLocalIP() (string, error) {
	if viper.GetString("getlocalIPTyoe") == "1" {
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
