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
		fmt.Println(r)
	}

}
