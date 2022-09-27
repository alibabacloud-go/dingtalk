// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package manufacturing_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type IndustrializeManufactureJobBookRequest struct {
	// 钉钉组织id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 扩展字段，用于增加自定义字段
	Extend *string `json:"extend,omitempty" xml:"extend,omitempty"`
	// 工单编号
	InstNo *string `json:"instNo,omitempty" xml:"instNo,omitempty"`
	// 是否是批量报工(取值[n,y])
	IsBatchJob *string `json:"isBatchJob,omitempty" xml:"isBatchJob,omitempty"`
	// 生产日期时间(到时分秒)
	ManufactureDate *string `json:"manufactureDate,omitempty" xml:"manufactureDate,omitempty"`
	// mes 系统唯一标识
	MesAppKey *string `json:"mesAppKey,omitempty" xml:"mesAppKey,omitempty"`
	// 制程英文名称
	ProcessEnName *string `json:"processEnName,omitempty" xml:"processEnName,omitempty"`
	// 制程名称
	ProcessName *string `json:"processName,omitempty" xml:"processName,omitempty"`
	// 产品唯一标识
	ProductCode *string `json:"productCode,omitempty" xml:"productCode,omitempty"`
	// 产品英文名称
	ProductEnName *string `json:"productEnName,omitempty" xml:"productEnName,omitempty"`
	// 产品名称，例如"双头螺柱001"
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// 产品规格
	ProductSpecification *string `json:"productSpecification,omitempty" xml:"productSpecification,omitempty"`
	// 合格数量
	QualifiedQuantity *string `json:"qualifiedQuantity,omitempty" xml:"qualifiedQuantity,omitempty"`
	// 可重工数量
	ReworkableQuantity *string `json:"reworkableQuantity,omitempty" xml:"reworkableQuantity,omitempty"`
	// 报废数量
	ScrappedQuantity *string `json:"scrappedQuantity,omitempty" xml:"scrappedQuantity,omitempty"`
	// 计件单价，单位：分
	UnitPrice *string `json:"unitPrice,omitempty" xml:"unitPrice,omitempty"`
	// 批量报工时多个工人userId以英文逗号分隔
	UserIdList *string `json:"userIdList,omitempty" xml:"userIdList,omitempty"`
	// 员工姓名
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// 批量报工时多个人名以英文逗号分隔
	UserNameList *string `json:"userNameList,omitempty" xml:"userNameList,omitempty"`
	// 随机串，唯一标识(用于幂等及更新)
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s IndustrializeManufactureJobBookRequest) String() string {
	return tea.Prettify(s)
}

func (s IndustrializeManufactureJobBookRequest) GoString() string {
	return s.String()
}

func (s *IndustrializeManufactureJobBookRequest) SetCorpId(v string) *IndustrializeManufactureJobBookRequest {
	s.CorpId = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetExtend(v string) *IndustrializeManufactureJobBookRequest {
	s.Extend = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetInstNo(v string) *IndustrializeManufactureJobBookRequest {
	s.InstNo = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetIsBatchJob(v string) *IndustrializeManufactureJobBookRequest {
	s.IsBatchJob = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetManufactureDate(v string) *IndustrializeManufactureJobBookRequest {
	s.ManufactureDate = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetMesAppKey(v string) *IndustrializeManufactureJobBookRequest {
	s.MesAppKey = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetProcessEnName(v string) *IndustrializeManufactureJobBookRequest {
	s.ProcessEnName = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetProcessName(v string) *IndustrializeManufactureJobBookRequest {
	s.ProcessName = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetProductCode(v string) *IndustrializeManufactureJobBookRequest {
	s.ProductCode = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetProductEnName(v string) *IndustrializeManufactureJobBookRequest {
	s.ProductEnName = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetProductName(v string) *IndustrializeManufactureJobBookRequest {
	s.ProductName = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetProductSpecification(v string) *IndustrializeManufactureJobBookRequest {
	s.ProductSpecification = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetQualifiedQuantity(v string) *IndustrializeManufactureJobBookRequest {
	s.QualifiedQuantity = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetReworkableQuantity(v string) *IndustrializeManufactureJobBookRequest {
	s.ReworkableQuantity = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetScrappedQuantity(v string) *IndustrializeManufactureJobBookRequest {
	s.ScrappedQuantity = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetUnitPrice(v string) *IndustrializeManufactureJobBookRequest {
	s.UnitPrice = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetUserIdList(v string) *IndustrializeManufactureJobBookRequest {
	s.UserIdList = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetUserName(v string) *IndustrializeManufactureJobBookRequest {
	s.UserName = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetUserNameList(v string) *IndustrializeManufactureJobBookRequest {
	s.UserNameList = &v
	return s
}

func (s *IndustrializeManufactureJobBookRequest) SetUuid(v string) *IndustrializeManufactureJobBookRequest {
	s.Uuid = &v
	return s
}

type IndustrializeManufactureJobBookResponseBody struct {
	// content
	Content *IndustrializeManufactureJobBookResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	// errorCode
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// errorLevel
	ErrorLevel *int32 `json:"errorLevel,omitempty" xml:"errorLevel,omitempty"`
	// errorMsg
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// httpCode
	HttpCode *string `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 此次报工记录的唯一标识
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s IndustrializeManufactureJobBookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndustrializeManufactureJobBookResponseBody) GoString() string {
	return s.String()
}

func (s *IndustrializeManufactureJobBookResponseBody) SetContent(v *IndustrializeManufactureJobBookResponseBodyContent) *IndustrializeManufactureJobBookResponseBody {
	s.Content = v
	return s
}

func (s *IndustrializeManufactureJobBookResponseBody) SetErrorCode(v string) *IndustrializeManufactureJobBookResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *IndustrializeManufactureJobBookResponseBody) SetErrorLevel(v int32) *IndustrializeManufactureJobBookResponseBody {
	s.ErrorLevel = &v
	return s
}

func (s *IndustrializeManufactureJobBookResponseBody) SetErrorMsg(v string) *IndustrializeManufactureJobBookResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *IndustrializeManufactureJobBookResponseBody) SetHttpCode(v string) *IndustrializeManufactureJobBookResponseBody {
	s.HttpCode = &v
	return s
}

func (s *IndustrializeManufactureJobBookResponseBody) SetSuccess(v bool) *IndustrializeManufactureJobBookResponseBody {
	s.Success = &v
	return s
}

func (s *IndustrializeManufactureJobBookResponseBody) SetUuid(v string) *IndustrializeManufactureJobBookResponseBody {
	s.Uuid = &v
	return s
}

type IndustrializeManufactureJobBookResponseBodyContent struct {
	// 影响行数
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// 新增记录的数据库id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s IndustrializeManufactureJobBookResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s IndustrializeManufactureJobBookResponseBodyContent) GoString() string {
	return s.String()
}

func (s *IndustrializeManufactureJobBookResponseBodyContent) SetCount(v int32) *IndustrializeManufactureJobBookResponseBodyContent {
	s.Count = &v
	return s
}

func (s *IndustrializeManufactureJobBookResponseBodyContent) SetId(v int64) *IndustrializeManufactureJobBookResponseBodyContent {
	s.Id = &v
	return s
}

type IndustrializeManufactureJobBookResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *IndustrializeManufactureJobBookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndustrializeManufactureJobBookResponse) String() string {
	return tea.Prettify(s)
}

func (s IndustrializeManufactureJobBookResponse) GoString() string {
	return s.String()
}

func (s *IndustrializeManufactureJobBookResponse) SetHeaders(v map[string]*string) *IndustrializeManufactureJobBookResponse {
	s.Headers = v
	return s
}

func (s *IndustrializeManufactureJobBookResponse) SetBody(v *IndustrializeManufactureJobBookResponseBody) *IndustrializeManufactureJobBookResponse {
	s.Body = v
	return s
}

type IndustrializeManufactureQueryJobsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IndustrializeManufactureQueryJobsHeaders) String() string {
	return tea.Prettify(s)
}

func (s IndustrializeManufactureQueryJobsHeaders) GoString() string {
	return s.String()
}

func (s *IndustrializeManufactureQueryJobsHeaders) SetCommonHeaders(v map[string]*string) *IndustrializeManufactureQueryJobsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IndustrializeManufactureQueryJobsHeaders) SetXAcsDingtalkAccessToken(v string) *IndustrializeManufactureQueryJobsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IndustrializeManufactureQueryJobsRequest struct {
	// 当前页序号(从1开始)
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// 工单编号
	InstNo *string `json:"instNo,omitempty" xml:"instNo,omitempty"`
	// 生产日期
	ManufactureDay *string `json:"manufactureDay,omitempty" xml:"manufactureDay,omitempty"`
	// MES系统唯一标识
	MesAppKey *string `json:"mesAppKey,omitempty" xml:"mesAppKey,omitempty"`
	// 每页显示记录条数
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 工序名称
	ProcessName *string `json:"processName,omitempty" xml:"processName,omitempty"`
	// 产品唯一标识
	ProductCode *string `json:"productCode,omitempty" xml:"productCode,omitempty"`
	// 产品中文名称
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// 产品规格
	ProductSpecification *string `json:"productSpecification,omitempty" xml:"productSpecification,omitempty"`
	// 报工合格数量
	QualifiedQuantity *string `json:"qualifiedQuantity,omitempty" xml:"qualifiedQuantity,omitempty"`
	// 计件单价，单位：分
	UnitPrice *string `json:"unitPrice,omitempty" xml:"unitPrice,omitempty"`
	// 员工钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 批量报工时多个人钉钉工号以英文逗号分隔
	UserIdList *string `json:"userIdList,omitempty" xml:"userIdList,omitempty"`
	// 员工姓名
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// 报工记录的唯一标识
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s IndustrializeManufactureQueryJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s IndustrializeManufactureQueryJobsRequest) GoString() string {
	return s.String()
}

func (s *IndustrializeManufactureQueryJobsRequest) SetCurrentPage(v int32) *IndustrializeManufactureQueryJobsRequest {
	s.CurrentPage = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsRequest) SetInstNo(v string) *IndustrializeManufactureQueryJobsRequest {
	s.InstNo = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsRequest) SetManufactureDay(v string) *IndustrializeManufactureQueryJobsRequest {
	s.ManufactureDay = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsRequest) SetMesAppKey(v string) *IndustrializeManufactureQueryJobsRequest {
	s.MesAppKey = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsRequest) SetPageSize(v int32) *IndustrializeManufactureQueryJobsRequest {
	s.PageSize = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsRequest) SetProcessName(v string) *IndustrializeManufactureQueryJobsRequest {
	s.ProcessName = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsRequest) SetProductCode(v string) *IndustrializeManufactureQueryJobsRequest {
	s.ProductCode = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsRequest) SetProductName(v string) *IndustrializeManufactureQueryJobsRequest {
	s.ProductName = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsRequest) SetProductSpecification(v string) *IndustrializeManufactureQueryJobsRequest {
	s.ProductSpecification = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsRequest) SetQualifiedQuantity(v string) *IndustrializeManufactureQueryJobsRequest {
	s.QualifiedQuantity = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsRequest) SetUnitPrice(v string) *IndustrializeManufactureQueryJobsRequest {
	s.UnitPrice = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsRequest) SetUserId(v string) *IndustrializeManufactureQueryJobsRequest {
	s.UserId = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsRequest) SetUserIdList(v string) *IndustrializeManufactureQueryJobsRequest {
	s.UserIdList = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsRequest) SetUserName(v string) *IndustrializeManufactureQueryJobsRequest {
	s.UserName = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsRequest) SetUuid(v string) *IndustrializeManufactureQueryJobsRequest {
	s.Uuid = &v
	return s
}

type IndustrializeManufactureQueryJobsResponseBody struct {
	// 查询的数据结果
	Content *IndustrializeManufactureQueryJobsResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	// httpCode
	HttpCode *string `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
}

func (s IndustrializeManufactureQueryJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndustrializeManufactureQueryJobsResponseBody) GoString() string {
	return s.String()
}

func (s *IndustrializeManufactureQueryJobsResponseBody) SetContent(v *IndustrializeManufactureQueryJobsResponseBodyContent) *IndustrializeManufactureQueryJobsResponseBody {
	s.Content = v
	return s
}

func (s *IndustrializeManufactureQueryJobsResponseBody) SetHttpCode(v string) *IndustrializeManufactureQueryJobsResponseBody {
	s.HttpCode = &v
	return s
}

type IndustrializeManufactureQueryJobsResponseBodyContent struct {
	// 组织id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 创建时间
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 修改时间
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 数据库id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 工单id
	InstNo *string `json:"instNo,omitempty" xml:"instNo,omitempty"`
	// 是否是批量报工，即一次报工由多个工人一起分担，取值[n,y],y表示是批量，批量时多个人名以英文逗号分隔
	IsBatchJob *string `json:"isBatchJob,omitempty" xml:"isBatchJob,omitempty"`
	// 生产日期时间(到时分秒),格式:2021-07-05 08:00:21
	ManufactureDate *string `json:"manufactureDate,omitempty" xml:"manufactureDate,omitempty"`
	// 生产日期(到天)
	ManufactureDay *string `json:"manufactureDay,omitempty" xml:"manufactureDay,omitempty"`
	// 分配给mes系统的appkey
	MesAppKey *string `json:"mesAppKey,omitempty" xml:"mesAppKey,omitempty"`
	// 工序名称
	ProcessName *string `json:"processName,omitempty" xml:"processName,omitempty"`
	// 合格数
	QualifiedQuantity *string `json:"qualifiedQuantity,omitempty" xml:"qualifiedQuantity,omitempty"`
	// 不合格数
	ScrappedQuantity *string `json:"scrappedQuantity,omitempty" xml:"scrappedQuantity,omitempty"`
	// 计件单价，单位：分
	UnitPrice *string `json:"unitPrice,omitempty" xml:"unitPrice,omitempty"`
	// 工人工号(isBatchJob=='n'时)
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 批量报工时多个人钉钉工号以英文逗号分隔
	UserIdList *string `json:"userIdList,omitempty" xml:"userIdList,omitempty"`
	// 批量报工时多个人名以英文逗号分隔
	UserNameList *string `json:"userNameList,omitempty" xml:"userNameList,omitempty"`
	// 报工记录的唯一标识
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s IndustrializeManufactureQueryJobsResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s IndustrializeManufactureQueryJobsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *IndustrializeManufactureQueryJobsResponseBodyContent) SetCorpId(v string) *IndustrializeManufactureQueryJobsResponseBodyContent {
	s.CorpId = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsResponseBodyContent) SetGmtCreate(v string) *IndustrializeManufactureQueryJobsResponseBodyContent {
	s.GmtCreate = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsResponseBodyContent) SetGmtModified(v string) *IndustrializeManufactureQueryJobsResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsResponseBodyContent) SetId(v int64) *IndustrializeManufactureQueryJobsResponseBodyContent {
	s.Id = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsResponseBodyContent) SetInstNo(v string) *IndustrializeManufactureQueryJobsResponseBodyContent {
	s.InstNo = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsResponseBodyContent) SetIsBatchJob(v string) *IndustrializeManufactureQueryJobsResponseBodyContent {
	s.IsBatchJob = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsResponseBodyContent) SetManufactureDate(v string) *IndustrializeManufactureQueryJobsResponseBodyContent {
	s.ManufactureDate = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsResponseBodyContent) SetManufactureDay(v string) *IndustrializeManufactureQueryJobsResponseBodyContent {
	s.ManufactureDay = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsResponseBodyContent) SetMesAppKey(v string) *IndustrializeManufactureQueryJobsResponseBodyContent {
	s.MesAppKey = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsResponseBodyContent) SetProcessName(v string) *IndustrializeManufactureQueryJobsResponseBodyContent {
	s.ProcessName = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsResponseBodyContent) SetQualifiedQuantity(v string) *IndustrializeManufactureQueryJobsResponseBodyContent {
	s.QualifiedQuantity = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsResponseBodyContent) SetScrappedQuantity(v string) *IndustrializeManufactureQueryJobsResponseBodyContent {
	s.ScrappedQuantity = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsResponseBodyContent) SetUnitPrice(v string) *IndustrializeManufactureQueryJobsResponseBodyContent {
	s.UnitPrice = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsResponseBodyContent) SetUserId(v string) *IndustrializeManufactureQueryJobsResponseBodyContent {
	s.UserId = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsResponseBodyContent) SetUserIdList(v string) *IndustrializeManufactureQueryJobsResponseBodyContent {
	s.UserIdList = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsResponseBodyContent) SetUserNameList(v string) *IndustrializeManufactureQueryJobsResponseBodyContent {
	s.UserNameList = &v
	return s
}

func (s *IndustrializeManufactureQueryJobsResponseBodyContent) SetUuid(v string) *IndustrializeManufactureQueryJobsResponseBodyContent {
	s.Uuid = &v
	return s
}

type IndustrializeManufactureQueryJobsResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *IndustrializeManufactureQueryJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndustrializeManufactureQueryJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s IndustrializeManufactureQueryJobsResponse) GoString() string {
	return s.String()
}

func (s *IndustrializeManufactureQueryJobsResponse) SetHeaders(v map[string]*string) *IndustrializeManufactureQueryJobsResponse {
	s.Headers = v
	return s
}

func (s *IndustrializeManufactureQueryJobsResponse) SetBody(v *IndustrializeManufactureQueryJobsResponseBody) *IndustrializeManufactureQueryJobsResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

func (client *Client) IndustrializeManufactureJobBook(userId *string, request *IndustrializeManufactureJobBookRequest) (_result *IndustrializeManufactureJobBookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &IndustrializeManufactureJobBookResponse{}
	_body, _err := client.IndustrializeManufactureJobBookWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndustrializeManufactureJobBookWithOptions(userId *string, request *IndustrializeManufactureJobBookRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *IndustrializeManufactureJobBookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Extend)) {
		body["extend"] = request.Extend
	}

	if !tea.BoolValue(util.IsUnset(request.InstNo)) {
		body["instNo"] = request.InstNo
	}

	if !tea.BoolValue(util.IsUnset(request.IsBatchJob)) {
		body["isBatchJob"] = request.IsBatchJob
	}

	if !tea.BoolValue(util.IsUnset(request.ManufactureDate)) {
		body["manufactureDate"] = request.ManufactureDate
	}

	if !tea.BoolValue(util.IsUnset(request.MesAppKey)) {
		body["mesAppKey"] = request.MesAppKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessEnName)) {
		body["processEnName"] = request.ProcessEnName
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessName)) {
		body["processName"] = request.ProcessName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["productCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductEnName)) {
		body["productEnName"] = request.ProductEnName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		body["productName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductSpecification)) {
		body["productSpecification"] = request.ProductSpecification
	}

	if !tea.BoolValue(util.IsUnset(request.QualifiedQuantity)) {
		body["qualifiedQuantity"] = request.QualifiedQuantity
	}

	if !tea.BoolValue(util.IsUnset(request.ReworkableQuantity)) {
		body["reworkableQuantity"] = request.ReworkableQuantity
	}

	if !tea.BoolValue(util.IsUnset(request.ScrappedQuantity)) {
		body["scrappedQuantity"] = request.ScrappedQuantity
	}

	if !tea.BoolValue(util.IsUnset(request.UnitPrice)) {
		body["unitPrice"] = request.UnitPrice
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdList)) {
		body["userIdList"] = request.UserIdList
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["userName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.UserNameList)) {
		body["userNameList"] = request.UserNameList
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &IndustrializeManufactureJobBookResponse{}
	_body, _err := client.DoROARequest(tea.String("IndustrializeManufactureJobBook"), tea.String("manufacturing_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/manufacturing/users/"+tea.StringValue(userId)+"/jobs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IndustrializeManufactureQueryJobs(request *IndustrializeManufactureQueryJobsRequest) (_result *IndustrializeManufactureQueryJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IndustrializeManufactureQueryJobsHeaders{}
	_result = &IndustrializeManufactureQueryJobsResponse{}
	_body, _err := client.IndustrializeManufactureQueryJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndustrializeManufactureQueryJobsWithOptions(request *IndustrializeManufactureQueryJobsRequest, headers *IndustrializeManufactureQueryJobsHeaders, runtime *util.RuntimeOptions) (_result *IndustrializeManufactureQueryJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["currentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.InstNo)) {
		body["instNo"] = request.InstNo
	}

	if !tea.BoolValue(util.IsUnset(request.ManufactureDay)) {
		body["manufactureDay"] = request.ManufactureDay
	}

	if !tea.BoolValue(util.IsUnset(request.MesAppKey)) {
		body["mesAppKey"] = request.MesAppKey
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessName)) {
		body["processName"] = request.ProcessName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["productCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		body["productName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductSpecification)) {
		body["productSpecification"] = request.ProductSpecification
	}

	if !tea.BoolValue(util.IsUnset(request.QualifiedQuantity)) {
		body["qualifiedQuantity"] = request.QualifiedQuantity
	}

	if !tea.BoolValue(util.IsUnset(request.UnitPrice)) {
		body["unitPrice"] = request.UnitPrice
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdList)) {
		body["userIdList"] = request.UserIdList
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["userName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["uuid"] = request.Uuid
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &IndustrializeManufactureQueryJobsResponse{}
	_body, _err := client.DoROARequest(tea.String("IndustrializeManufactureQueryJobs"), tea.String("manufacturing_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/manufacturing/users/jobs/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
