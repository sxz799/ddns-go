package api

import (
	"ddns-go/model"
	"ddns-go/utils"
	"encoding/json"
	alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

// GetDomainRecordsByRRKeyWord 根据关键字获取解析记录
func GetDomainRecordsByRRKeyWord(key, domain string) (model.DomainRecords, error) {
	describeDomainRecordsRequest := &alidns20150109.DescribeDomainRecordsRequest{
		DomainName: tea.String(domain),
		RRKeyWord:  tea.String(key),
	}
	runtime := &util.RuntimeOptions{}
	resp, _err := utils.ApiClient.DescribeDomainRecordsWithOptions(describeDomainRecordsRequest, runtime)
	if _err != nil {
		return model.DomainRecords{}, _err
	} else {
		var OpenAPI model.OpenAPIResponse
		err := json.Unmarshal([]byte(*util.ToJSONString(resp.Body)), &OpenAPI)
		if err != nil {
			return model.DomainRecords{}, err
		} else {
			return OpenAPI.DomainRecords, nil
		}
	}

}

// GetDomainRecords 获取所有解析记录
func GetDomainRecords(domain string) (model.DomainRecords, error) {
	describeDomainRecordsRequest := &alidns20150109.DescribeDomainRecordsRequest{
		DomainName: tea.String(domain),
	}
	runtime := &util.RuntimeOptions{}
	resp, _err := utils.ApiClient.DescribeDomainRecordsWithOptions(describeDomainRecordsRequest, runtime)
	if _err != nil {
		return model.DomainRecords{}, _err
	} else {
		var OpenAPI model.OpenAPIResponse
		err := json.Unmarshal([]byte(*util.ToJSONString(resp.Body)), &OpenAPI)
		if err != nil {
			return model.DomainRecords{}, err
		} else {
			return OpenAPI.DomainRecords, nil
		}
	}

}

// AddDomainRecord 添加一条解析计量
func AddDomainRecord(domain, rr, tType, value string) (model.RecordResponse, error) {
	addDomainRecordRequest := &alidns20150109.AddDomainRecordRequest{
		DomainName: tea.String(domain),
		RR:         tea.String(rr),
		Type:       tea.String(tType),
		Value:      tea.String(value),
	}
	runtime := &util.RuntimeOptions{}
	resp, _err := utils.ApiClient.AddDomainRecordWithOptions(addDomainRecordRequest, runtime)
	if _err != nil {
		return model.RecordResponse{}, _err
	} else {
		var r model.RecordResponse
		err := json.Unmarshal([]byte(*util.ToJSONString(resp.Body)), &r)
		if err != nil {
			return model.RecordResponse{}, _err
		} else {
			return r, nil
		}
	}
}

// UpdateDomainRecord 更新一条解析记录
func UpdateDomainRecord(recordId, rr, tType, value string) (model.RecordResponse, error) {
	updateDomainRecordRequest := &alidns20150109.UpdateDomainRecordRequest{
		RecordId: tea.String(recordId),
		RR:       tea.String(rr),
		Type:     tea.String(tType),
		Value:    tea.String(value),
	}
	runtime := &util.RuntimeOptions{}
	resp, _err := utils.ApiClient.UpdateDomainRecordWithOptions(updateDomainRecordRequest, runtime)
	if _err != nil {
		return model.RecordResponse{}, _err
	} else {
		var r model.RecordResponse
		err := json.Unmarshal([]byte(*util.ToJSONString(resp.Body)), &r)
		if err != nil {
			return model.RecordResponse{}, _err
		} else {
			return r, nil
		}
	}
}

// DelDomainRecord 删除一条解析记录
func DelDomainRecord(recordId string) (model.RecordResponse, error) {
	deleteDomainRecordRequest := &alidns20150109.DeleteDomainRecordRequest{
		RecordId: tea.String(recordId),
	}
	runtime := &util.RuntimeOptions{}
	resp, _err := utils.ApiClient.DeleteDomainRecordWithOptions(deleteDomainRecordRequest, runtime)
	if _err != nil {
		return model.RecordResponse{}, _err
	} else {
		var r model.RecordResponse
		err := json.Unmarshal([]byte(*util.ToJSONString(resp.Body)), &r)
		if err != nil {
			return model.RecordResponse{}, _err
		} else {
			return r, nil
		}
	}
}
