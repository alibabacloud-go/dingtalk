// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package amdp_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AmdpEmployeeDataPushHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AmdpEmployeeDataPushHeaders) String() string {
	return tea.Prettify(s)
}

func (s AmdpEmployeeDataPushHeaders) GoString() string {
	return s.String()
}

func (s *AmdpEmployeeDataPushHeaders) SetCommonHeaders(v map[string]*string) *AmdpEmployeeDataPushHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AmdpEmployeeDataPushHeaders) SetXAcsDingtalkAccessToken(v string) *AmdpEmployeeDataPushHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AmdpEmployeeDataPushRequest struct {
	Param []*AmdpEmployeeDataPushRequestParam `json:"param,omitempty" xml:"param,omitempty" type:"Repeated"`
}

func (s AmdpEmployeeDataPushRequest) String() string {
	return tea.Prettify(s)
}

func (s AmdpEmployeeDataPushRequest) GoString() string {
	return s.String()
}

func (s *AmdpEmployeeDataPushRequest) SetParam(v []*AmdpEmployeeDataPushRequestParam) *AmdpEmployeeDataPushRequest {
	s.Param = v
	return s
}

type AmdpEmployeeDataPushRequestParam struct {
	Avatar   *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	IsDelete *string `json:"isDelete,omitempty" xml:"isDelete,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	UnionId  *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId   *string `json:"userId,omitempty" xml:"userId,omitempty"`
	WorkNo   *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s AmdpEmployeeDataPushRequestParam) String() string {
	return tea.Prettify(s)
}

func (s AmdpEmployeeDataPushRequestParam) GoString() string {
	return s.String()
}

func (s *AmdpEmployeeDataPushRequestParam) SetAvatar(v string) *AmdpEmployeeDataPushRequestParam {
	s.Avatar = &v
	return s
}

func (s *AmdpEmployeeDataPushRequestParam) SetIsDelete(v string) *AmdpEmployeeDataPushRequestParam {
	s.IsDelete = &v
	return s
}

func (s *AmdpEmployeeDataPushRequestParam) SetName(v string) *AmdpEmployeeDataPushRequestParam {
	s.Name = &v
	return s
}

func (s *AmdpEmployeeDataPushRequestParam) SetUnionId(v string) *AmdpEmployeeDataPushRequestParam {
	s.UnionId = &v
	return s
}

func (s *AmdpEmployeeDataPushRequestParam) SetUserId(v string) *AmdpEmployeeDataPushRequestParam {
	s.UserId = &v
	return s
}

func (s *AmdpEmployeeDataPushRequestParam) SetWorkNo(v string) *AmdpEmployeeDataPushRequestParam {
	s.WorkNo = &v
	return s
}

type AmdpEmployeeDataPushResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AmdpEmployeeDataPushResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AmdpEmployeeDataPushResponseBody) GoString() string {
	return s.String()
}

func (s *AmdpEmployeeDataPushResponseBody) SetRequestId(v string) *AmdpEmployeeDataPushResponseBody {
	s.RequestId = &v
	return s
}

func (s *AmdpEmployeeDataPushResponseBody) SetResult(v bool) *AmdpEmployeeDataPushResponseBody {
	s.Result = &v
	return s
}

func (s *AmdpEmployeeDataPushResponseBody) SetStatus(v string) *AmdpEmployeeDataPushResponseBody {
	s.Status = &v
	return s
}

func (s *AmdpEmployeeDataPushResponseBody) SetSuccess(v bool) *AmdpEmployeeDataPushResponseBody {
	s.Success = &v
	return s
}

type AmdpEmployeeDataPushResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AmdpEmployeeDataPushResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AmdpEmployeeDataPushResponse) String() string {
	return tea.Prettify(s)
}

func (s AmdpEmployeeDataPushResponse) GoString() string {
	return s.String()
}

func (s *AmdpEmployeeDataPushResponse) SetHeaders(v map[string]*string) *AmdpEmployeeDataPushResponse {
	s.Headers = v
	return s
}

func (s *AmdpEmployeeDataPushResponse) SetStatusCode(v int32) *AmdpEmployeeDataPushResponse {
	s.StatusCode = &v
	return s
}

func (s *AmdpEmployeeDataPushResponse) SetBody(v *AmdpEmployeeDataPushResponseBody) *AmdpEmployeeDataPushResponse {
	s.Body = v
	return s
}

type AmdpJobPositionDataPushHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AmdpJobPositionDataPushHeaders) String() string {
	return tea.Prettify(s)
}

func (s AmdpJobPositionDataPushHeaders) GoString() string {
	return s.String()
}

func (s *AmdpJobPositionDataPushHeaders) SetCommonHeaders(v map[string]*string) *AmdpJobPositionDataPushHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AmdpJobPositionDataPushHeaders) SetXAcsDingtalkAccessToken(v string) *AmdpJobPositionDataPushHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AmdpJobPositionDataPushRequest struct {
	Param []*AmdpJobPositionDataPushRequestParam `json:"param,omitempty" xml:"param,omitempty" type:"Repeated"`
}

func (s AmdpJobPositionDataPushRequest) String() string {
	return tea.Prettify(s)
}

func (s AmdpJobPositionDataPushRequest) GoString() string {
	return s.String()
}

func (s *AmdpJobPositionDataPushRequest) SetParam(v []*AmdpJobPositionDataPushRequestParam) *AmdpJobPositionDataPushRequest {
	s.Param = v
	return s
}

type AmdpJobPositionDataPushRequestParam struct {
	DeptId       *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
	DeptLeader   *string `json:"deptLeader,omitempty" xml:"deptLeader,omitempty"`
	IsDelete     *string `json:"isDelete,omitempty" xml:"isDelete,omitempty"`
	LeaderDeptId *string `json:"leaderDeptId,omitempty" xml:"leaderDeptId,omitempty"`
	OrderNumber  *string `json:"orderNumber,omitempty" xml:"orderNumber,omitempty"`
	UserId       *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AmdpJobPositionDataPushRequestParam) String() string {
	return tea.Prettify(s)
}

func (s AmdpJobPositionDataPushRequestParam) GoString() string {
	return s.String()
}

func (s *AmdpJobPositionDataPushRequestParam) SetDeptId(v string) *AmdpJobPositionDataPushRequestParam {
	s.DeptId = &v
	return s
}

func (s *AmdpJobPositionDataPushRequestParam) SetDeptLeader(v string) *AmdpJobPositionDataPushRequestParam {
	s.DeptLeader = &v
	return s
}

func (s *AmdpJobPositionDataPushRequestParam) SetIsDelete(v string) *AmdpJobPositionDataPushRequestParam {
	s.IsDelete = &v
	return s
}

func (s *AmdpJobPositionDataPushRequestParam) SetLeaderDeptId(v string) *AmdpJobPositionDataPushRequestParam {
	s.LeaderDeptId = &v
	return s
}

func (s *AmdpJobPositionDataPushRequestParam) SetOrderNumber(v string) *AmdpJobPositionDataPushRequestParam {
	s.OrderNumber = &v
	return s
}

func (s *AmdpJobPositionDataPushRequestParam) SetUserId(v string) *AmdpJobPositionDataPushRequestParam {
	s.UserId = &v
	return s
}

type AmdpJobPositionDataPushResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AmdpJobPositionDataPushResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AmdpJobPositionDataPushResponseBody) GoString() string {
	return s.String()
}

func (s *AmdpJobPositionDataPushResponseBody) SetRequestId(v string) *AmdpJobPositionDataPushResponseBody {
	s.RequestId = &v
	return s
}

func (s *AmdpJobPositionDataPushResponseBody) SetResult(v bool) *AmdpJobPositionDataPushResponseBody {
	s.Result = &v
	return s
}

func (s *AmdpJobPositionDataPushResponseBody) SetStatus(v string) *AmdpJobPositionDataPushResponseBody {
	s.Status = &v
	return s
}

func (s *AmdpJobPositionDataPushResponseBody) SetSuccess(v bool) *AmdpJobPositionDataPushResponseBody {
	s.Success = &v
	return s
}

type AmdpJobPositionDataPushResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AmdpJobPositionDataPushResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AmdpJobPositionDataPushResponse) String() string {
	return tea.Prettify(s)
}

func (s AmdpJobPositionDataPushResponse) GoString() string {
	return s.String()
}

func (s *AmdpJobPositionDataPushResponse) SetHeaders(v map[string]*string) *AmdpJobPositionDataPushResponse {
	s.Headers = v
	return s
}

func (s *AmdpJobPositionDataPushResponse) SetStatusCode(v int32) *AmdpJobPositionDataPushResponse {
	s.StatusCode = &v
	return s
}

func (s *AmdpJobPositionDataPushResponse) SetBody(v *AmdpJobPositionDataPushResponseBody) *AmdpJobPositionDataPushResponse {
	s.Body = v
	return s
}

type AmdpOrganizationDataPushHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AmdpOrganizationDataPushHeaders) String() string {
	return tea.Prettify(s)
}

func (s AmdpOrganizationDataPushHeaders) GoString() string {
	return s.String()
}

func (s *AmdpOrganizationDataPushHeaders) SetCommonHeaders(v map[string]*string) *AmdpOrganizationDataPushHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AmdpOrganizationDataPushHeaders) SetXAcsDingtalkAccessToken(v string) *AmdpOrganizationDataPushHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AmdpOrganizationDataPushRequest struct {
	Param []*AmdpOrganizationDataPushRequestParam `json:"param,omitempty" xml:"param,omitempty" type:"Repeated"`
}

func (s AmdpOrganizationDataPushRequest) String() string {
	return tea.Prettify(s)
}

func (s AmdpOrganizationDataPushRequest) GoString() string {
	return s.String()
}

func (s *AmdpOrganizationDataPushRequest) SetParam(v []*AmdpOrganizationDataPushRequestParam) *AmdpOrganizationDataPushRequest {
	s.Param = v
	return s
}

type AmdpOrganizationDataPushRequestParam struct {
	DeptId            *string   `json:"deptId,omitempty" xml:"deptId,omitempty"`
	DeptManagerIdList []*string `json:"deptManagerIdList,omitempty" xml:"deptManagerIdList,omitempty" type:"Repeated"`
	DingTalkDeptId    *string   `json:"dingTalkDeptId,omitempty" xml:"dingTalkDeptId,omitempty"`
	DingTalkParentId  *string   `json:"dingTalkParentId,omitempty" xml:"dingTalkParentId,omitempty"`
	IsDelete          *string   `json:"isDelete,omitempty" xml:"isDelete,omitempty"`
	Name              *string   `json:"name,omitempty" xml:"name,omitempty"`
	ParentId          *string   `json:"parentId,omitempty" xml:"parentId,omitempty"`
}

func (s AmdpOrganizationDataPushRequestParam) String() string {
	return tea.Prettify(s)
}

func (s AmdpOrganizationDataPushRequestParam) GoString() string {
	return s.String()
}

func (s *AmdpOrganizationDataPushRequestParam) SetDeptId(v string) *AmdpOrganizationDataPushRequestParam {
	s.DeptId = &v
	return s
}

func (s *AmdpOrganizationDataPushRequestParam) SetDeptManagerIdList(v []*string) *AmdpOrganizationDataPushRequestParam {
	s.DeptManagerIdList = v
	return s
}

func (s *AmdpOrganizationDataPushRequestParam) SetDingTalkDeptId(v string) *AmdpOrganizationDataPushRequestParam {
	s.DingTalkDeptId = &v
	return s
}

func (s *AmdpOrganizationDataPushRequestParam) SetDingTalkParentId(v string) *AmdpOrganizationDataPushRequestParam {
	s.DingTalkParentId = &v
	return s
}

func (s *AmdpOrganizationDataPushRequestParam) SetIsDelete(v string) *AmdpOrganizationDataPushRequestParam {
	s.IsDelete = &v
	return s
}

func (s *AmdpOrganizationDataPushRequestParam) SetName(v string) *AmdpOrganizationDataPushRequestParam {
	s.Name = &v
	return s
}

func (s *AmdpOrganizationDataPushRequestParam) SetParentId(v string) *AmdpOrganizationDataPushRequestParam {
	s.ParentId = &v
	return s
}

type AmdpOrganizationDataPushResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AmdpOrganizationDataPushResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AmdpOrganizationDataPushResponseBody) GoString() string {
	return s.String()
}

func (s *AmdpOrganizationDataPushResponseBody) SetRequestId(v string) *AmdpOrganizationDataPushResponseBody {
	s.RequestId = &v
	return s
}

func (s *AmdpOrganizationDataPushResponseBody) SetResult(v bool) *AmdpOrganizationDataPushResponseBody {
	s.Result = &v
	return s
}

func (s *AmdpOrganizationDataPushResponseBody) SetStatus(v string) *AmdpOrganizationDataPushResponseBody {
	s.Status = &v
	return s
}

func (s *AmdpOrganizationDataPushResponseBody) SetSuccess(v bool) *AmdpOrganizationDataPushResponseBody {
	s.Success = &v
	return s
}

type AmdpOrganizationDataPushResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AmdpOrganizationDataPushResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AmdpOrganizationDataPushResponse) String() string {
	return tea.Prettify(s)
}

func (s AmdpOrganizationDataPushResponse) GoString() string {
	return s.String()
}

func (s *AmdpOrganizationDataPushResponse) SetHeaders(v map[string]*string) *AmdpOrganizationDataPushResponse {
	s.Headers = v
	return s
}

func (s *AmdpOrganizationDataPushResponse) SetStatusCode(v int32) *AmdpOrganizationDataPushResponse {
	s.StatusCode = &v
	return s
}

func (s *AmdpOrganizationDataPushResponse) SetBody(v *AmdpOrganizationDataPushResponseBody) *AmdpOrganizationDataPushResponse {
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
	interfaceSPI, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = interfaceSPI
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

func (client *Client) AmdpEmployeeDataPushWithOptions(request *AmdpEmployeeDataPushRequest, headers *AmdpEmployeeDataPushHeaders, runtime *util.RuntimeOptions) (_result *AmdpEmployeeDataPushResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Param)) {
		body["param"] = request.Param
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
	params := &openapi.Params{
		Action:      tea.String("AmdpEmployeeDataPush"),
		Version:     tea.String("amdp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/amdp/employees/datas/push"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AmdpEmployeeDataPushResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AmdpEmployeeDataPush(request *AmdpEmployeeDataPushRequest) (_result *AmdpEmployeeDataPushResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AmdpEmployeeDataPushHeaders{}
	_result = &AmdpEmployeeDataPushResponse{}
	_body, _err := client.AmdpEmployeeDataPushWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AmdpJobPositionDataPushWithOptions(request *AmdpJobPositionDataPushRequest, headers *AmdpJobPositionDataPushHeaders, runtime *util.RuntimeOptions) (_result *AmdpJobPositionDataPushResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Param)) {
		body["param"] = request.Param
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
	params := &openapi.Params{
		Action:      tea.String("AmdpJobPositionDataPush"),
		Version:     tea.String("amdp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/amdp/empJobs/datas/push"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AmdpJobPositionDataPushResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AmdpJobPositionDataPush(request *AmdpJobPositionDataPushRequest) (_result *AmdpJobPositionDataPushResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AmdpJobPositionDataPushHeaders{}
	_result = &AmdpJobPositionDataPushResponse{}
	_body, _err := client.AmdpJobPositionDataPushWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AmdpOrganizationDataPushWithOptions(request *AmdpOrganizationDataPushRequest, headers *AmdpOrganizationDataPushHeaders, runtime *util.RuntimeOptions) (_result *AmdpOrganizationDataPushResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Param)) {
		body["param"] = request.Param
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
	params := &openapi.Params{
		Action:      tea.String("AmdpOrganizationDataPush"),
		Version:     tea.String("amdp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/amdp/organizations/departments/datas/push"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AmdpOrganizationDataPushResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AmdpOrganizationDataPush(request *AmdpOrganizationDataPushRequest) (_result *AmdpOrganizationDataPushResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AmdpOrganizationDataPushHeaders{}
	_result = &AmdpOrganizationDataPushResponse{}
	_body, _err := client.AmdpOrganizationDataPushWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
