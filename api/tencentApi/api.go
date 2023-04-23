package tencentApi

import (
	"ddns-go/utils"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

// ListDomainRecords 获取所有解析记录
func ListDomainRecords(domain string) {
	request := dnspod.NewDescribeRecordListRequest()
	request.Domain = common.StringPtr(domain)
	response, err := utils.TencentClient.DescribeRecordList(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}

// AddDomainRecord 添加一条解析计量
func AddDomainRecord(domain, subDomail, tType, line, value string) {
	request := dnspod.NewCreateRecordRequest()
	request.Domain = common.StringPtr(domain)
	request.SubDomain = common.StringPtr(subDomail)
	request.RecordType = common.StringPtr(tType)
	request.RecordLine = common.StringPtr(line)
	request.Value = common.StringPtr(value)

	// 返回的resp是一个CreateRecordResponse的实例，与请求对象对应
	response, err := utils.TencentClient.CreateRecord(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}

// UpdateDomainRecord 更新一条解析记录
func UpdateDomainRecord(domain, tType, line, value string, recordId uint64) {
	request := dnspod.NewModifyRecordRequest()

	request.Domain = common.StringPtr(domain)
	request.RecordType = common.StringPtr(tType)
	request.RecordLine = common.StringPtr(line)
	request.Value = common.StringPtr(value)
	request.RecordId = common.Uint64Ptr(recordId)

	// 返回的resp是一个ModifyRecordResponse的实例，与请求对象对应
	response, err := utils.TencentClient.ModifyRecord(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}

// DelDomainRecord 删除一条解析记录
func DelDomainRecord(domain string, recordId uint64) {
	request := dnspod.NewDeleteRecordRequest()

	request.Domain = common.StringPtr(domain)
	request.RecordId = common.Uint64Ptr(recordId)

	// 返回的resp是一个DeleteRecordResponse的实例，与请求对象对应
	response, err := utils.TencentClient.DeleteRecord(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}
