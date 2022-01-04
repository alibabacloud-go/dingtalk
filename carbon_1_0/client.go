// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package carbon_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type WriteAlibabaOrgCarbonHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s WriteAlibabaOrgCarbonHeaders) String() string {
	return tea.Prettify(s)
}

func (s WriteAlibabaOrgCarbonHeaders) GoString() string {
	return s.String()
}

func (s *WriteAlibabaOrgCarbonHeaders) SetCommonHeaders(v map[string]*string) *WriteAlibabaOrgCarbonHeaders {
	s.CommonHeaders = v
	return s
}

func (s *WriteAlibabaOrgCarbonHeaders) SetXAcsDingtalkAccessToken(v string) *WriteAlibabaOrgCarbonHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type WriteAlibabaOrgCarbonRequest struct {
	// 入参集
	OrgDetailsList []*WriteAlibabaOrgCarbonRequestOrgDetailsList `json:"orgDetailsList,omitempty" xml:"orgDetailsList,omitempty" type:"Repeated"`
}

func (s WriteAlibabaOrgCarbonRequest) String() string {
	return tea.Prettify(s)
}

func (s WriteAlibabaOrgCarbonRequest) GoString() string {
	return s.String()
}

func (s *WriteAlibabaOrgCarbonRequest) SetOrgDetailsList(v []*WriteAlibabaOrgCarbonRequestOrgDetailsList) *WriteAlibabaOrgCarbonRequest {
	s.OrgDetailsList = v
	return s
}

type WriteAlibabaOrgCarbonRequestOrgDetailsList struct {
	// 系统唯一id，生成格式：userId+日期20211126
	ActionId *string `json:"actionId,omitempty" xml:"actionId,omitempty"`
	// 行为发生时间
	ActionTime *string `json:"actionTime,omitempty" xml:"actionTime,omitempty"`
	// 碳能量行为类型，需要联系管理员添加
	ActionType *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// 碳能量数据
	CarbonAmount *string `json:"carbonAmount,omitempty" xml:"carbonAmount,omitempty"`
	// 钉钉组织id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 钉钉部门id
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 版本，默认为1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s WriteAlibabaOrgCarbonRequestOrgDetailsList) String() string {
	return tea.Prettify(s)
}

func (s WriteAlibabaOrgCarbonRequestOrgDetailsList) GoString() string {
	return s.String()
}

func (s *WriteAlibabaOrgCarbonRequestOrgDetailsList) SetActionId(v string) *WriteAlibabaOrgCarbonRequestOrgDetailsList {
	s.ActionId = &v
	return s
}

func (s *WriteAlibabaOrgCarbonRequestOrgDetailsList) SetActionTime(v string) *WriteAlibabaOrgCarbonRequestOrgDetailsList {
	s.ActionTime = &v
	return s
}

func (s *WriteAlibabaOrgCarbonRequestOrgDetailsList) SetActionType(v string) *WriteAlibabaOrgCarbonRequestOrgDetailsList {
	s.ActionType = &v
	return s
}

func (s *WriteAlibabaOrgCarbonRequestOrgDetailsList) SetCarbonAmount(v string) *WriteAlibabaOrgCarbonRequestOrgDetailsList {
	s.CarbonAmount = &v
	return s
}

func (s *WriteAlibabaOrgCarbonRequestOrgDetailsList) SetCorpId(v string) *WriteAlibabaOrgCarbonRequestOrgDetailsList {
	s.CorpId = &v
	return s
}

func (s *WriteAlibabaOrgCarbonRequestOrgDetailsList) SetDeptId(v int64) *WriteAlibabaOrgCarbonRequestOrgDetailsList {
	s.DeptId = &v
	return s
}

func (s *WriteAlibabaOrgCarbonRequestOrgDetailsList) SetVersion(v int32) *WriteAlibabaOrgCarbonRequestOrgDetailsList {
	s.Version = &v
	return s
}

type WriteAlibabaOrgCarbonResponseBody struct {
	// 返回请求成功的数量
	Result *int32 `json:"result,omitempty" xml:"result,omitempty"`
	// 请求是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s WriteAlibabaOrgCarbonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WriteAlibabaOrgCarbonResponseBody) GoString() string {
	return s.String()
}

func (s *WriteAlibabaOrgCarbonResponseBody) SetResult(v int32) *WriteAlibabaOrgCarbonResponseBody {
	s.Result = &v
	return s
}

func (s *WriteAlibabaOrgCarbonResponseBody) SetSuccess(v bool) *WriteAlibabaOrgCarbonResponseBody {
	s.Success = &v
	return s
}

type WriteAlibabaOrgCarbonResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *WriteAlibabaOrgCarbonResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s WriteAlibabaOrgCarbonResponse) String() string {
	return tea.Prettify(s)
}

func (s WriteAlibabaOrgCarbonResponse) GoString() string {
	return s.String()
}

func (s *WriteAlibabaOrgCarbonResponse) SetHeaders(v map[string]*string) *WriteAlibabaOrgCarbonResponse {
	s.Headers = v
	return s
}

func (s *WriteAlibabaOrgCarbonResponse) SetBody(v *WriteAlibabaOrgCarbonResponseBody) *WriteAlibabaOrgCarbonResponse {
	s.Body = v
	return s
}

type WriteAlibabaUserCarbonHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s WriteAlibabaUserCarbonHeaders) String() string {
	return tea.Prettify(s)
}

func (s WriteAlibabaUserCarbonHeaders) GoString() string {
	return s.String()
}

func (s *WriteAlibabaUserCarbonHeaders) SetCommonHeaders(v map[string]*string) *WriteAlibabaUserCarbonHeaders {
	s.CommonHeaders = v
	return s
}

func (s *WriteAlibabaUserCarbonHeaders) SetXAcsDingtalkAccessToken(v string) *WriteAlibabaUserCarbonHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type WriteAlibabaUserCarbonRequest struct {
	// 入参集
	UserDetailsList []*WriteAlibabaUserCarbonRequestUserDetailsList `json:"userDetailsList,omitempty" xml:"userDetailsList,omitempty" type:"Repeated"`
}

func (s WriteAlibabaUserCarbonRequest) String() string {
	return tea.Prettify(s)
}

func (s WriteAlibabaUserCarbonRequest) GoString() string {
	return s.String()
}

func (s *WriteAlibabaUserCarbonRequest) SetUserDetailsList(v []*WriteAlibabaUserCarbonRequestUserDetailsList) *WriteAlibabaUserCarbonRequest {
	s.UserDetailsList = v
	return s
}

type WriteAlibabaUserCarbonRequestUserDetailsList struct {
	// 行为结束时间
	ActionEndTime *string `json:"actionEndTime,omitempty" xml:"actionEndTime,omitempty"`
	// 系统唯一id，生成格式：userId+日期20211126
	ActionId *string `json:"actionId,omitempty" xml:"actionId,omitempty"`
	// 行为起始时间
	ActionStartTime *string `json:"actionStartTime,omitempty" xml:"actionStartTime,omitempty"`
	// 碳能量行为类型，需要联系管理员添加
	ActionType *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// 碳能量数据
	CarbonAmount *string `json:"carbonAmount,omitempty" xml:"carbonAmount,omitempty"`
	// 钉钉组织id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 钉钉部门id
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 钉钉用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 版本，默认为1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s WriteAlibabaUserCarbonRequestUserDetailsList) String() string {
	return tea.Prettify(s)
}

func (s WriteAlibabaUserCarbonRequestUserDetailsList) GoString() string {
	return s.String()
}

func (s *WriteAlibabaUserCarbonRequestUserDetailsList) SetActionEndTime(v string) *WriteAlibabaUserCarbonRequestUserDetailsList {
	s.ActionEndTime = &v
	return s
}

func (s *WriteAlibabaUserCarbonRequestUserDetailsList) SetActionId(v string) *WriteAlibabaUserCarbonRequestUserDetailsList {
	s.ActionId = &v
	return s
}

func (s *WriteAlibabaUserCarbonRequestUserDetailsList) SetActionStartTime(v string) *WriteAlibabaUserCarbonRequestUserDetailsList {
	s.ActionStartTime = &v
	return s
}

func (s *WriteAlibabaUserCarbonRequestUserDetailsList) SetActionType(v string) *WriteAlibabaUserCarbonRequestUserDetailsList {
	s.ActionType = &v
	return s
}

func (s *WriteAlibabaUserCarbonRequestUserDetailsList) SetCarbonAmount(v string) *WriteAlibabaUserCarbonRequestUserDetailsList {
	s.CarbonAmount = &v
	return s
}

func (s *WriteAlibabaUserCarbonRequestUserDetailsList) SetCorpId(v string) *WriteAlibabaUserCarbonRequestUserDetailsList {
	s.CorpId = &v
	return s
}

func (s *WriteAlibabaUserCarbonRequestUserDetailsList) SetDeptId(v int64) *WriteAlibabaUserCarbonRequestUserDetailsList {
	s.DeptId = &v
	return s
}

func (s *WriteAlibabaUserCarbonRequestUserDetailsList) SetUserId(v string) *WriteAlibabaUserCarbonRequestUserDetailsList {
	s.UserId = &v
	return s
}

func (s *WriteAlibabaUserCarbonRequestUserDetailsList) SetVersion(v int32) *WriteAlibabaUserCarbonRequestUserDetailsList {
	s.Version = &v
	return s
}

type WriteAlibabaUserCarbonResponseBody struct {
	// 返回请求成功个数
	Result *int32 `json:"result,omitempty" xml:"result,omitempty"`
	// 请求是否写入成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s WriteAlibabaUserCarbonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WriteAlibabaUserCarbonResponseBody) GoString() string {
	return s.String()
}

func (s *WriteAlibabaUserCarbonResponseBody) SetResult(v int32) *WriteAlibabaUserCarbonResponseBody {
	s.Result = &v
	return s
}

func (s *WriteAlibabaUserCarbonResponseBody) SetSuccess(v bool) *WriteAlibabaUserCarbonResponseBody {
	s.Success = &v
	return s
}

type WriteAlibabaUserCarbonResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *WriteAlibabaUserCarbonResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s WriteAlibabaUserCarbonResponse) String() string {
	return tea.Prettify(s)
}

func (s WriteAlibabaUserCarbonResponse) GoString() string {
	return s.String()
}

func (s *WriteAlibabaUserCarbonResponse) SetHeaders(v map[string]*string) *WriteAlibabaUserCarbonResponse {
	s.Headers = v
	return s
}

func (s *WriteAlibabaUserCarbonResponse) SetBody(v *WriteAlibabaUserCarbonResponseBody) *WriteAlibabaUserCarbonResponse {
	s.Body = v
	return s
}

type WriteOrgCarbonHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s WriteOrgCarbonHeaders) String() string {
	return tea.Prettify(s)
}

func (s WriteOrgCarbonHeaders) GoString() string {
	return s.String()
}

func (s *WriteOrgCarbonHeaders) SetCommonHeaders(v map[string]*string) *WriteOrgCarbonHeaders {
	s.CommonHeaders = v
	return s
}

func (s *WriteOrgCarbonHeaders) SetXAcsDingtalkAccessToken(v string) *WriteOrgCarbonHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type WriteOrgCarbonRequest struct {
	// 入参集
	OrgDetailsList []*WriteOrgCarbonRequestOrgDetailsList `json:"orgDetailsList,omitempty" xml:"orgDetailsList,omitempty" type:"Repeated"`
}

func (s WriteOrgCarbonRequest) String() string {
	return tea.Prettify(s)
}

func (s WriteOrgCarbonRequest) GoString() string {
	return s.String()
}

func (s *WriteOrgCarbonRequest) SetOrgDetailsList(v []*WriteOrgCarbonRequestOrgDetailsList) *WriteOrgCarbonRequest {
	s.OrgDetailsList = v
	return s
}

type WriteOrgCarbonRequestOrgDetailsList struct {
	// 系统唯一id，生成格式：userId+日期20211126
	ActionId *string `json:"actionId,omitempty" xml:"actionId,omitempty"`
	// 行为发生时间
	ActionTime *string `json:"actionTime,omitempty" xml:"actionTime,omitempty"`
	// 碳能量行为类型，需要联系管理员添加
	ActionType *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// 碳能量数据
	CarbonAmount *string `json:"carbonAmount,omitempty" xml:"carbonAmount,omitempty"`
	// 钉钉组织id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 钉钉部门id
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 版本，默认为1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s WriteOrgCarbonRequestOrgDetailsList) String() string {
	return tea.Prettify(s)
}

func (s WriteOrgCarbonRequestOrgDetailsList) GoString() string {
	return s.String()
}

func (s *WriteOrgCarbonRequestOrgDetailsList) SetActionId(v string) *WriteOrgCarbonRequestOrgDetailsList {
	s.ActionId = &v
	return s
}

func (s *WriteOrgCarbonRequestOrgDetailsList) SetActionTime(v string) *WriteOrgCarbonRequestOrgDetailsList {
	s.ActionTime = &v
	return s
}

func (s *WriteOrgCarbonRequestOrgDetailsList) SetActionType(v string) *WriteOrgCarbonRequestOrgDetailsList {
	s.ActionType = &v
	return s
}

func (s *WriteOrgCarbonRequestOrgDetailsList) SetCarbonAmount(v string) *WriteOrgCarbonRequestOrgDetailsList {
	s.CarbonAmount = &v
	return s
}

func (s *WriteOrgCarbonRequestOrgDetailsList) SetCorpId(v string) *WriteOrgCarbonRequestOrgDetailsList {
	s.CorpId = &v
	return s
}

func (s *WriteOrgCarbonRequestOrgDetailsList) SetDeptId(v int64) *WriteOrgCarbonRequestOrgDetailsList {
	s.DeptId = &v
	return s
}

func (s *WriteOrgCarbonRequestOrgDetailsList) SetVersion(v int32) *WriteOrgCarbonRequestOrgDetailsList {
	s.Version = &v
	return s
}

type WriteOrgCarbonResponseBody struct {
	// 请求成功返回的个数
	Result *int32 `json:"result,omitempty" xml:"result,omitempty"`
	// 请求是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s WriteOrgCarbonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WriteOrgCarbonResponseBody) GoString() string {
	return s.String()
}

func (s *WriteOrgCarbonResponseBody) SetResult(v int32) *WriteOrgCarbonResponseBody {
	s.Result = &v
	return s
}

func (s *WriteOrgCarbonResponseBody) SetSuccess(v bool) *WriteOrgCarbonResponseBody {
	s.Success = &v
	return s
}

type WriteOrgCarbonResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *WriteOrgCarbonResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s WriteOrgCarbonResponse) String() string {
	return tea.Prettify(s)
}

func (s WriteOrgCarbonResponse) GoString() string {
	return s.String()
}

func (s *WriteOrgCarbonResponse) SetHeaders(v map[string]*string) *WriteOrgCarbonResponse {
	s.Headers = v
	return s
}

func (s *WriteOrgCarbonResponse) SetBody(v *WriteOrgCarbonResponseBody) *WriteOrgCarbonResponse {
	s.Body = v
	return s
}

type WriteUserCarbonHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s WriteUserCarbonHeaders) String() string {
	return tea.Prettify(s)
}

func (s WriteUserCarbonHeaders) GoString() string {
	return s.String()
}

func (s *WriteUserCarbonHeaders) SetCommonHeaders(v map[string]*string) *WriteUserCarbonHeaders {
	s.CommonHeaders = v
	return s
}

func (s *WriteUserCarbonHeaders) SetXAcsDingtalkAccessToken(v string) *WriteUserCarbonHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type WriteUserCarbonRequest struct {
	// 入参集
	UserDetailsList []*WriteUserCarbonRequestUserDetailsList `json:"userDetailsList,omitempty" xml:"userDetailsList,omitempty" type:"Repeated"`
}

func (s WriteUserCarbonRequest) String() string {
	return tea.Prettify(s)
}

func (s WriteUserCarbonRequest) GoString() string {
	return s.String()
}

func (s *WriteUserCarbonRequest) SetUserDetailsList(v []*WriteUserCarbonRequestUserDetailsList) *WriteUserCarbonRequest {
	s.UserDetailsList = v
	return s
}

type WriteUserCarbonRequestUserDetailsList struct {
	// 行为结束时间
	ActionEndTime *string `json:"actionEndTime,omitempty" xml:"actionEndTime,omitempty"`
	// 系统唯一id，生成格式：userId+日期20211126
	ActionId *string `json:"actionId,omitempty" xml:"actionId,omitempty"`
	// 行为起始时间
	ActionStartTime *string `json:"actionStartTime,omitempty" xml:"actionStartTime,omitempty"`
	// 碳能量行为类型，需要联系管理员添加
	ActionType *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// 碳能量数据
	CarbonAmount *string `json:"carbonAmount,omitempty" xml:"carbonAmount,omitempty"`
	// 钉钉组织id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 钉钉部门id
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 钉钉用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 版本，默认为1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s WriteUserCarbonRequestUserDetailsList) String() string {
	return tea.Prettify(s)
}

func (s WriteUserCarbonRequestUserDetailsList) GoString() string {
	return s.String()
}

func (s *WriteUserCarbonRequestUserDetailsList) SetActionEndTime(v string) *WriteUserCarbonRequestUserDetailsList {
	s.ActionEndTime = &v
	return s
}

func (s *WriteUserCarbonRequestUserDetailsList) SetActionId(v string) *WriteUserCarbonRequestUserDetailsList {
	s.ActionId = &v
	return s
}

func (s *WriteUserCarbonRequestUserDetailsList) SetActionStartTime(v string) *WriteUserCarbonRequestUserDetailsList {
	s.ActionStartTime = &v
	return s
}

func (s *WriteUserCarbonRequestUserDetailsList) SetActionType(v string) *WriteUserCarbonRequestUserDetailsList {
	s.ActionType = &v
	return s
}

func (s *WriteUserCarbonRequestUserDetailsList) SetCarbonAmount(v string) *WriteUserCarbonRequestUserDetailsList {
	s.CarbonAmount = &v
	return s
}

func (s *WriteUserCarbonRequestUserDetailsList) SetCorpId(v string) *WriteUserCarbonRequestUserDetailsList {
	s.CorpId = &v
	return s
}

func (s *WriteUserCarbonRequestUserDetailsList) SetDeptId(v int64) *WriteUserCarbonRequestUserDetailsList {
	s.DeptId = &v
	return s
}

func (s *WriteUserCarbonRequestUserDetailsList) SetUserId(v string) *WriteUserCarbonRequestUserDetailsList {
	s.UserId = &v
	return s
}

func (s *WriteUserCarbonRequestUserDetailsList) SetVersion(v int32) *WriteUserCarbonRequestUserDetailsList {
	s.Version = &v
	return s
}

type WriteUserCarbonResponseBody struct {
	// 返回请求成功个数
	Result *int32 `json:"result,omitempty" xml:"result,omitempty"`
	// 请求是否写入成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s WriteUserCarbonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WriteUserCarbonResponseBody) GoString() string {
	return s.String()
}

func (s *WriteUserCarbonResponseBody) SetResult(v int32) *WriteUserCarbonResponseBody {
	s.Result = &v
	return s
}

func (s *WriteUserCarbonResponseBody) SetSuccess(v bool) *WriteUserCarbonResponseBody {
	s.Success = &v
	return s
}

type WriteUserCarbonResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *WriteUserCarbonResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s WriteUserCarbonResponse) String() string {
	return tea.Prettify(s)
}

func (s WriteUserCarbonResponse) GoString() string {
	return s.String()
}

func (s *WriteUserCarbonResponse) SetHeaders(v map[string]*string) *WriteUserCarbonResponse {
	s.Headers = v
	return s
}

func (s *WriteUserCarbonResponse) SetBody(v *WriteUserCarbonResponseBody) *WriteUserCarbonResponse {
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

func (client *Client) WriteAlibabaOrgCarbon(request *WriteAlibabaOrgCarbonRequest) (_result *WriteAlibabaOrgCarbonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &WriteAlibabaOrgCarbonHeaders{}
	_result = &WriteAlibabaOrgCarbonResponse{}
	_body, _err := client.WriteAlibabaOrgCarbonWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) WriteAlibabaOrgCarbonWithOptions(request *WriteAlibabaOrgCarbonRequest, headers *WriteAlibabaOrgCarbonHeaders, runtime *util.RuntimeOptions) (_result *WriteAlibabaOrgCarbonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgDetailsList)) {
		body["orgDetailsList"] = request.OrgDetailsList
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
	_result = &WriteAlibabaOrgCarbonResponse{}
	_body, _err := client.DoROARequest(tea.String("WriteAlibabaOrgCarbon"), tea.String("carbon_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/carbon/alibabaOrgDetails/write"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) WriteAlibabaUserCarbon(request *WriteAlibabaUserCarbonRequest) (_result *WriteAlibabaUserCarbonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &WriteAlibabaUserCarbonHeaders{}
	_result = &WriteAlibabaUserCarbonResponse{}
	_body, _err := client.WriteAlibabaUserCarbonWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) WriteAlibabaUserCarbonWithOptions(request *WriteAlibabaUserCarbonRequest, headers *WriteAlibabaUserCarbonHeaders, runtime *util.RuntimeOptions) (_result *WriteAlibabaUserCarbonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserDetailsList)) {
		body["userDetailsList"] = request.UserDetailsList
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
	_result = &WriteAlibabaUserCarbonResponse{}
	_body, _err := client.DoROARequest(tea.String("WriteAlibabaUserCarbon"), tea.String("carbon_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/carbon/alibabaUserDetails/write"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) WriteOrgCarbon(request *WriteOrgCarbonRequest) (_result *WriteOrgCarbonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &WriteOrgCarbonHeaders{}
	_result = &WriteOrgCarbonResponse{}
	_body, _err := client.WriteOrgCarbonWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) WriteOrgCarbonWithOptions(request *WriteOrgCarbonRequest, headers *WriteOrgCarbonHeaders, runtime *util.RuntimeOptions) (_result *WriteOrgCarbonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgDetailsList)) {
		body["orgDetailsList"] = request.OrgDetailsList
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
	_result = &WriteOrgCarbonResponse{}
	_body, _err := client.DoROARequest(tea.String("WriteOrgCarbon"), tea.String("carbon_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/carbon/orgDetails/write"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) WriteUserCarbon(request *WriteUserCarbonRequest) (_result *WriteUserCarbonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &WriteUserCarbonHeaders{}
	_result = &WriteUserCarbonResponse{}
	_body, _err := client.WriteUserCarbonWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) WriteUserCarbonWithOptions(request *WriteUserCarbonRequest, headers *WriteUserCarbonHeaders, runtime *util.RuntimeOptions) (_result *WriteUserCarbonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserDetailsList)) {
		body["userDetailsList"] = request.UserDetailsList
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
	_result = &WriteUserCarbonResponse{}
	_body, _err := client.DoROARequest(tea.String("WriteUserCarbon"), tea.String("carbon_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/carbon/userDetails/write"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
