package main

import (
	"ddns-go/service"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	domain := viper.GetString("domain")
	rr := viper.GetString("rr")
	records, err := service.GetDomainRecordsByRRKeyWord(rr, domain)
	if err != nil {
		panic(err)
	}
	for _, r := range records.Record {
		//service.UpdateDomainRecord(r.RecordID, rr, "A", "89.88.88.88")
		fmt.Println(r.RecordID, r.RR, r.Value, r.Type)
	}
	//service.AddDomainRecord(domain, "test2", "A", "66.99.55.66")

}
