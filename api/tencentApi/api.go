package tencentApi

import (
	"ddns-go/utils"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

// ListDomainRecords 获取所有解析记录
func ListDomainRecords(domain string) ([]*dnspod.RecordListItem, error) {
	request := dnspod.NewDescribeRecordListRequest()
	request.Domain = common.StringPtr(domain)
	response, err := utils.TencentClient.DescribeRecordList(request)
	if err != nil {
		return nil, err
	} else {
		return response.Response.RecordList, nil
	}

}

// AddDomainRecord 添加一条解析计量
func AddDomainRecord(domain, subDomain, tType, line, value string) (uint64, error) {
	request := dnspod.NewCreateRecordRequest()

	request.Domain = common.StringPtr(domain)
	request.SubDomain = common.StringPtr(subDomain)
	request.RecordType = common.StringPtr(tType)
	request.RecordLine = common.StringPtr(line)
	request.Value = common.StringPtr(value)

	// 返回的resp是一个CreateRecordResponse的实例，与请求对象对应
	response, err := utils.TencentClient.CreateRecord(request)
	if err != nil {
		return 0, err
	} else {
		return *response.Response.RecordId, nil
	}

}

// UpdateDomainRecord 更新一条解析记录
func UpdateDomainRecord(domain, subDomain, tType, line, value string, recordId uint64) error {
	request := dnspod.NewModifyRecordRequest()

	request.Domain = common.StringPtr(domain)
	request.SubDomain = common.StringPtr(subDomain)
	request.RecordType = common.StringPtr(tType)
	request.RecordLine = common.StringPtr(line)
	request.Value = common.StringPtr(value)
	request.RecordId = common.Uint64Ptr(recordId)

	// 返回的resp是一个ModifyRecordResponse的实例，与请求对象对应
	_, err := utils.TencentClient.ModifyRecord(request)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// DelDomainRecord 删除一条解析记录
func DelDomainRecord(domain string, recordId uint64) error {
	request := dnspod.NewDeleteRecordRequest()

	request.Domain = common.StringPtr(domain)
	request.RecordId = common.Uint64Ptr(recordId)

	// 返回的resp是一个DeleteRecordResponse的实例，与请求对象对应
	_, err := utils.TencentClient.DeleteRecord(request)
	if err != nil {
		return err
	} else {
		return nil
	}
}
