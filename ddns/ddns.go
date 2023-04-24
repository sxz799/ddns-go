package ddns

import (
	"ddns-go/api/aliApi"
	"ddns-go/api/tencentApi"
	"github.com/spf13/viper"
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
