// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package crm_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type QueryAllCustomerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllCustomerHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCustomerHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllCustomerHeaders) SetCommonHeaders(v map[string]*string) *QueryAllCustomerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllCustomerHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllCustomerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllCustomerRequest struct {
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	DingCorpId         *string `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	// 用户ID
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// 翻页size
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 分页游标，第一次调用传空或者null
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 数据类型
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s QueryAllCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCustomerRequest) GoString() string {
	return s.String()
}

func (s *QueryAllCustomerRequest) SetDingIsvOrgId(v int64) *QueryAllCustomerRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *QueryAllCustomerRequest) SetDingOrgId(v int64) *QueryAllCustomerRequest {
	s.DingOrgId = &v
	return s
}

func (s *QueryAllCustomerRequest) SetDingTokenGrantType(v int64) *QueryAllCustomerRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *QueryAllCustomerRequest) SetDingCorpId(v string) *QueryAllCustomerRequest {
	s.DingCorpId = &v
	return s
}

func (s *QueryAllCustomerRequest) SetDingSuiteKey(v string) *QueryAllCustomerRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *QueryAllCustomerRequest) SetOperatorUserId(v string) *QueryAllCustomerRequest {
	s.OperatorUserId = &v
	return s
}

func (s *QueryAllCustomerRequest) SetMaxResults(v int64) *QueryAllCustomerRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryAllCustomerRequest) SetNextToken(v string) *QueryAllCustomerRequest {
	s.NextToken = &v
	return s
}

func (s *QueryAllCustomerRequest) SetObjectType(v string) *QueryAllCustomerRequest {
	s.ObjectType = &v
	return s
}

type QueryAllCustomerResponseBody struct {
	// 分页结果
	Result *QueryAllCustomerResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryAllCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllCustomerResponseBody) SetResult(v *QueryAllCustomerResponseBodyResult) *QueryAllCustomerResponseBody {
	s.Result = v
	return s
}

type QueryAllCustomerResponseBodyResult struct {
	// 下一页的游标，为null则表示无数据
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 客户数据节点
	Values []*QueryAllCustomerResponseBodyResultValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
	// 分页大小
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
}

func (s QueryAllCustomerResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCustomerResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryAllCustomerResponseBodyResult) SetNextToken(v string) *QueryAllCustomerResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResult) SetValues(v []*QueryAllCustomerResponseBodyResultValues) *QueryAllCustomerResponseBodyResult {
	s.Values = v
	return s
}

func (s *QueryAllCustomerResponseBodyResult) SetMaxResults(v int64) *QueryAllCustomerResponseBodyResult {
	s.MaxResults = &v
	return s
}

type QueryAllCustomerResponseBodyResultValues struct {
	// 创建记录的用户昵称
	CreatorNick *string `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	// 记录修改时间
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// 创建记录的用户ID
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// 数据ID
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// 数据内容
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// 扩展数据内容
	ExtendData map[string]interface{} `json:"extendData,omitempty" xml:"extendData,omitempty"`
	// 记录创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 系统自动生成
	OrgId *int64 `json:"orgId,omitempty" xml:"orgId,omitempty"`
	// 数据类型
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// 数据权限信息
	Permission *QueryAllCustomerResponseBodyResultValuesPermission `json:"permission,omitempty" xml:"permission,omitempty" type:"Struct"`
	// 审批结果
	ProcessOutResult *string `json:"processOutResult,omitempty" xml:"processOutResult,omitempty"`
	// 审批状态
	ProcessInstanceStatus *string `json:"processInstanceStatus,omitempty" xml:"processInstanceStatus,omitempty"`
}

func (s QueryAllCustomerResponseBodyResultValues) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCustomerResponseBodyResultValues) GoString() string {
	return s.String()
}

func (s *QueryAllCustomerResponseBodyResultValues) SetCreatorNick(v string) *QueryAllCustomerResponseBodyResultValues {
	s.CreatorNick = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetModifyTime(v string) *QueryAllCustomerResponseBodyResultValues {
	s.ModifyTime = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetCreatorUserId(v string) *QueryAllCustomerResponseBodyResultValues {
	s.CreatorUserId = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetInstanceId(v string) *QueryAllCustomerResponseBodyResultValues {
	s.InstanceId = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetData(v map[string]interface{}) *QueryAllCustomerResponseBodyResultValues {
	s.Data = v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetExtendData(v map[string]interface{}) *QueryAllCustomerResponseBodyResultValues {
	s.ExtendData = v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetCreateTime(v string) *QueryAllCustomerResponseBodyResultValues {
	s.CreateTime = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetOrgId(v int64) *QueryAllCustomerResponseBodyResultValues {
	s.OrgId = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetObjectType(v string) *QueryAllCustomerResponseBodyResultValues {
	s.ObjectType = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetPermission(v *QueryAllCustomerResponseBodyResultValuesPermission) *QueryAllCustomerResponseBodyResultValues {
	s.Permission = v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetProcessOutResult(v string) *QueryAllCustomerResponseBodyResultValues {
	s.ProcessOutResult = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetProcessInstanceStatus(v string) *QueryAllCustomerResponseBodyResultValues {
	s.ProcessInstanceStatus = &v
	return s
}

type QueryAllCustomerResponseBodyResultValuesPermission struct {
	// 协同人用户ID列表
	ParticipantStaffIds []*string `json:"participantStaffIds,omitempty" xml:"participantStaffIds,omitempty" type:"Repeated"`
	// 负责人用户ID列表
	OwnerStaffIds []*string `json:"ownerStaffIds,omitempty" xml:"ownerStaffIds,omitempty" type:"Repeated"`
}

func (s QueryAllCustomerResponseBodyResultValuesPermission) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCustomerResponseBodyResultValuesPermission) GoString() string {
	return s.String()
}

func (s *QueryAllCustomerResponseBodyResultValuesPermission) SetParticipantStaffIds(v []*string) *QueryAllCustomerResponseBodyResultValuesPermission {
	s.ParticipantStaffIds = v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValuesPermission) SetOwnerStaffIds(v []*string) *QueryAllCustomerResponseBodyResultValuesPermission {
	s.OwnerStaffIds = v
	return s
}

type QueryAllCustomerResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAllCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCustomerResponse) GoString() string {
	return s.String()
}

func (s *QueryAllCustomerResponse) SetHeaders(v map[string]*string) *QueryAllCustomerResponse {
	s.Headers = v
	return s
}

func (s *QueryAllCustomerResponse) SetBody(v *QueryAllCustomerResponseBody) *QueryAllCustomerResponse {
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

func (client *Client) QueryAllCustomer(request *QueryAllCustomerRequest) (_result *QueryAllCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllCustomerHeaders{}
	_result = &QueryAllCustomerResponse{}
	_body, _err := client.QueryAllCustomerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllCustomerWithOptions(request *QueryAllCustomerRequest, headers *QueryAllCustomerHeaders, runtime *util.RuntimeOptions) (_result *QueryAllCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.DingCorpId)) {
		body["dingCorpId"] = request.DingCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectType)) {
		body["objectType"] = request.ObjectType
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &QueryAllCustomerResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAllCustomer"), tea.String("crm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/crm/customerInstances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
