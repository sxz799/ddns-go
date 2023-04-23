package aliApi

import (
	"ddns-go/model/aliModel"
	"ddns-go/utils"
	"encoding/json"
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

// GetDomainRecordsByRRKeyWord 根据关键字获取解析记录
func GetDomainRecordsByRRKeyWord(key, domain string) ([]aliModel.Record, error) {
	describeDomainRecordsRequest := &alidns20150109.DescribeDomainRecordsRequest{
		DomainName: tea.String(domain),
		RRKeyWord:  tea.String(key),
	}
	runtime := &util.RuntimeOptions{}
	resp, _err := utils.AliApiClient.DescribeDomainRecordsWithOptions(describeDomainRecordsRequest, runtime)
	if _err != nil {
		return nil, _err
	} else {
		var aliResp aliModel.AliAPIResponse
		err := json.Unmarshal([]byte(*util.ToJSONString(resp.Body)), &aliResp)
		if err != nil {
			return nil, err
		} else {
			return aliResp.DomainRecords.Record, nil
		}
	}

}

// ListDomainRecords 获取所有解析记录
func ListDomainRecords(domain string) ([]aliModel.Record, error) {
	describeDomainRecordsRequest := &alidns20150109.DescribeDomainRecordsRequest{
		DomainName: tea.String(domain),
	}
	runtime := &util.RuntimeOptions{}
	resp, _err := utils.AliApiClient.DescribeDomainRecordsWithOptions(describeDomainRecordsRequest, runtime)
	if _err != nil {
		return nil, _err
	} else {
		var aliResp aliModel.AliAPIResponse
		err := json.Unmarshal([]byte(*util.ToJSONString(resp.Body)), &aliResp)
		if err != nil {
			return nil, err
		} else {
			return aliResp.DomainRecords.Record, nil
		}
	}

}

// AddDomainRecord 添加一条解析计量
func AddDomainRecord(domain, rr, tType, value string) (string, error) {
	addDomainRecordRequest := &alidns20150109.AddDomainRecordRequest{
		DomainName: tea.String(domain),
		RR:         tea.String(rr),
		Type:       tea.String(tType),
		Value:      tea.String(value),
	}
	runtime := &util.RuntimeOptions{}
	resp, _err := utils.AliApiClient.AddDomainRecordWithOptions(addDomainRecordRequest, runtime)
	if _err != nil {
		return "", _err
	} else {
		var r aliModel.RecordResponse
		err := json.Unmarshal([]byte(*util.ToJSONString(resp.Body)), &r)
		if err != nil {
			return "", _err
		} else {
			return r.RecordID, nil
		}
	}
}

// UpdateDomainRecord 更新一条解析记录
func UpdateDomainRecord(recordId, rr, tType, value string) error {
	updateDomainRecordRequest := &alidns20150109.UpdateDomainRecordRequest{
		RecordId: tea.String(recordId),
		RR:       tea.String(rr),
		Type:     tea.String(tType),
		Value:    tea.String(value),
	}
	runtime := &util.RuntimeOptions{}
	resp, _err := utils.AliApiClient.UpdateDomainRecordWithOptions(updateDomainRecordRequest, runtime)
	if _err != nil {
		return _err
	} else {
		var r aliModel.RecordResponse
		err := json.Unmarshal([]byte(*util.ToJSONString(resp.Body)), &r)
		if err != nil {
			return _err
		} else {
			return nil
		}
	}
}

// DelDomainRecord 删除一条解析记录
func DelDomainRecord(recordId string) error {
	deleteDomainRecordRequest := &alidns20150109.DeleteDomainRecordRequest{
		RecordId: tea.String(recordId),
	}
	runtime := &util.RuntimeOptions{}
	resp, _err := utils.AliApiClient.DeleteDomainRecordWithOptions(deleteDomainRecordRequest, runtime)
	if _err != nil {
		return _err
	} else {
		var r aliModel.RecordResponse
		err := json.Unmarshal([]byte(*util.ToJSONString(resp.Body)), &r)
		if err != nil {
			return _err
		} else {
			return nil
		}
	}
}
