// This file is auto-generated, don't edit it. Thanks.
package transcribe_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetTranscribeBriefHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTranscribeBriefHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTranscribeBriefHeaders) GoString() string {
	return s.String()
}

func (s *GetTranscribeBriefHeaders) SetCommonHeaders(v map[string]*string) *GetTranscribeBriefHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTranscribeBriefHeaders) SetXAcsDingtalkAccessToken(v string) *GetTranscribeBriefHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTranscribeBriefResponseBody struct {
	Data       *GetTranscribeBriefResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Success    *bool                               `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTranscribeBriefResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTranscribeBriefResponseBody) GoString() string {
	return s.String()
}

func (s *GetTranscribeBriefResponseBody) SetData(v *GetTranscribeBriefResponseBodyData) *GetTranscribeBriefResponseBody {
	s.Data = v
	return s
}

func (s *GetTranscribeBriefResponseBody) SetStatusCode(v int32) *GetTranscribeBriefResponseBody {
	s.StatusCode = &v
	return s
}

func (s *GetTranscribeBriefResponseBody) SetSuccess(v bool) *GetTranscribeBriefResponseBody {
	s.Success = &v
	return s
}

type GetTranscribeBriefResponseBodyData struct {
	BizType *int32 `json:"bizType,omitempty" xml:"bizType,omitempty"`
}

func (s GetTranscribeBriefResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTranscribeBriefResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTranscribeBriefResponseBodyData) SetBizType(v int32) *GetTranscribeBriefResponseBodyData {
	s.BizType = &v
	return s
}

type GetTranscribeBriefResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTranscribeBriefResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTranscribeBriefResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTranscribeBriefResponse) GoString() string {
	return s.String()
}

func (s *GetTranscribeBriefResponse) SetHeaders(v map[string]*string) *GetTranscribeBriefResponse {
	s.Headers = v
	return s
}

func (s *GetTranscribeBriefResponse) SetStatusCode(v int32) *GetTranscribeBriefResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTranscribeBriefResponse) SetBody(v *GetTranscribeBriefResponseBody) *GetTranscribeBriefResponse {
	s.Body = v
	return s
}

type RemovePermissionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RemovePermissionHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemovePermissionHeaders) GoString() string {
	return s.String()
}

func (s *RemovePermissionHeaders) SetCommonHeaders(v map[string]*string) *RemovePermissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemovePermissionHeaders) SetXAcsDingtalkAccessToken(v string) *RemovePermissionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RemovePermissionRequest struct {
	// This parameter is required.
	BizType *int32 `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	Members     []*RemovePermissionRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	TaskCreator *int64                            `json:"taskCreator,omitempty" xml:"taskCreator,omitempty"`
	TaskId      *int64                            `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s RemovePermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s RemovePermissionRequest) GoString() string {
	return s.String()
}

func (s *RemovePermissionRequest) SetBizType(v int32) *RemovePermissionRequest {
	s.BizType = &v
	return s
}

func (s *RemovePermissionRequest) SetMembers(v []*RemovePermissionRequestMembers) *RemovePermissionRequest {
	s.Members = v
	return s
}

func (s *RemovePermissionRequest) SetTaskCreator(v int64) *RemovePermissionRequest {
	s.TaskCreator = &v
	return s
}

func (s *RemovePermissionRequest) SetTaskId(v int64) *RemovePermissionRequest {
	s.TaskId = &v
	return s
}

type RemovePermissionRequestMembers struct {
	// This parameter is required.
	//
	// example:
	//
	// 533xxxxxx
	MemberId *int64 `json:"memberId,omitempty" xml:"memberId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// EDITOR
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// This parameter is required.
	PolicyType *string `json:"policyType,omitempty" xml:"policyType,omitempty"`
}

func (s RemovePermissionRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s RemovePermissionRequestMembers) GoString() string {
	return s.String()
}

func (s *RemovePermissionRequestMembers) SetMemberId(v int64) *RemovePermissionRequestMembers {
	s.MemberId = &v
	return s
}

func (s *RemovePermissionRequestMembers) SetMemberType(v string) *RemovePermissionRequestMembers {
	s.MemberType = &v
	return s
}

func (s *RemovePermissionRequestMembers) SetPolicyType(v string) *RemovePermissionRequestMembers {
	s.PolicyType = &v
	return s
}

type RemovePermissionResponseBody struct {
	Data       *RemovePermissionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Success    *bool                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RemovePermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemovePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *RemovePermissionResponseBody) SetData(v *RemovePermissionResponseBodyData) *RemovePermissionResponseBody {
	s.Data = v
	return s
}

func (s *RemovePermissionResponseBody) SetStatusCode(v int32) *RemovePermissionResponseBody {
	s.StatusCode = &v
	return s
}

func (s *RemovePermissionResponseBody) SetSuccess(v bool) *RemovePermissionResponseBody {
	s.Success = &v
	return s
}

type RemovePermissionResponseBodyData struct {
	FailNameList []*string `json:"failNameList,omitempty" xml:"failNameList,omitempty" type:"Repeated"`
}

func (s RemovePermissionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RemovePermissionResponseBodyData) GoString() string {
	return s.String()
}

func (s *RemovePermissionResponseBodyData) SetFailNameList(v []*string) *RemovePermissionResponseBodyData {
	s.FailNameList = v
	return s
}

type RemovePermissionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemovePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemovePermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s RemovePermissionResponse) GoString() string {
	return s.String()
}

func (s *RemovePermissionResponse) SetHeaders(v map[string]*string) *RemovePermissionResponse {
	s.Headers = v
	return s
}

func (s *RemovePermissionResponse) SetStatusCode(v int32) *RemovePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *RemovePermissionResponse) SetBody(v *RemovePermissionResponseBody) *RemovePermissionResponse {
	s.Body = v
	return s
}

type UpdatePermissionForUsersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdatePermissionForUsersHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdatePermissionForUsersHeaders) GoString() string {
	return s.String()
}

func (s *UpdatePermissionForUsersHeaders) SetCommonHeaders(v map[string]*string) *UpdatePermissionForUsersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdatePermissionForUsersHeaders) SetXAcsDingtalkAccessToken(v string) *UpdatePermissionForUsersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdatePermissionForUsersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	Members []*UpdatePermissionForUsersRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 533xxxxxx
	TaskCreator *int64 `json:"taskCreator,omitempty" xml:"taskCreator,omitempty"`
	OperatorUid *int64 `json:"operatorUid,omitempty" xml:"operatorUid,omitempty"`
}

func (s UpdatePermissionForUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePermissionForUsersRequest) GoString() string {
	return s.String()
}

func (s *UpdatePermissionForUsersRequest) SetBizType(v int32) *UpdatePermissionForUsersRequest {
	s.BizType = &v
	return s
}

func (s *UpdatePermissionForUsersRequest) SetMembers(v []*UpdatePermissionForUsersRequestMembers) *UpdatePermissionForUsersRequest {
	s.Members = v
	return s
}

func (s *UpdatePermissionForUsersRequest) SetTaskCreator(v int64) *UpdatePermissionForUsersRequest {
	s.TaskCreator = &v
	return s
}

func (s *UpdatePermissionForUsersRequest) SetOperatorUid(v int64) *UpdatePermissionForUsersRequest {
	s.OperatorUid = &v
	return s
}

type UpdatePermissionForUsersRequestMembers struct {
	// This parameter is required.
	MemberId *int64 `json:"memberId,omitempty" xml:"memberId,omitempty"`
	// This parameter is required.
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// EDITOR
	PolicyType *string `json:"policyType,omitempty" xml:"policyType,omitempty"`
}

func (s UpdatePermissionForUsersRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s UpdatePermissionForUsersRequestMembers) GoString() string {
	return s.String()
}

func (s *UpdatePermissionForUsersRequestMembers) SetMemberId(v int64) *UpdatePermissionForUsersRequestMembers {
	s.MemberId = &v
	return s
}

func (s *UpdatePermissionForUsersRequestMembers) SetMemberType(v string) *UpdatePermissionForUsersRequestMembers {
	s.MemberType = &v
	return s
}

func (s *UpdatePermissionForUsersRequestMembers) SetPolicyType(v string) *UpdatePermissionForUsersRequestMembers {
	s.PolicyType = &v
	return s
}

type UpdatePermissionForUsersResponseBody struct {
	IsSuccess *bool `json:"isSuccess,omitempty" xml:"isSuccess,omitempty"`
}

func (s UpdatePermissionForUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePermissionForUsersResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePermissionForUsersResponseBody) SetIsSuccess(v bool) *UpdatePermissionForUsersResponseBody {
	s.IsSuccess = &v
	return s
}

type UpdatePermissionForUsersResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePermissionForUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePermissionForUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePermissionForUsersResponse) GoString() string {
	return s.String()
}

func (s *UpdatePermissionForUsersResponse) SetHeaders(v map[string]*string) *UpdatePermissionForUsersResponse {
	s.Headers = v
	return s
}

func (s *UpdatePermissionForUsersResponse) SetStatusCode(v int32) *UpdatePermissionForUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePermissionForUsersResponse) SetBody(v *UpdatePermissionForUsersResponseBody) *UpdatePermissionForUsersResponse {
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
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

// Summary:
//
// 获取闪记任务的概要信息
//
// @param headers - GetTranscribeBriefHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTranscribeBriefResponse
func (client *Client) GetTranscribeBriefWithOptions(taskUuid *string, headers *GetTranscribeBriefHeaders, runtime *util.RuntimeOptions) (_result *GetTranscribeBriefResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTranscribeBrief"),
		Version:     tea.String("transcribe_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/transcribe/tasks/" + tea.StringValue(taskUuid) + "/briefInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTranscribeBriefResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取闪记任务的概要信息
//
// @return GetTranscribeBriefResponse
func (client *Client) GetTranscribeBrief(taskUuid *string) (_result *GetTranscribeBriefResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTranscribeBriefHeaders{}
	_result = &GetTranscribeBriefResponse{}
	_body, _err := client.GetTranscribeBriefWithOptions(taskUuid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移除指定用户对闪记任务的权限
//
// @param request - RemovePermissionRequest
//
// @param headers - RemovePermissionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemovePermissionResponse
func (client *Client) RemovePermissionWithOptions(taskUuid *string, request *RemovePermissionRequest, headers *RemovePermissionHeaders, runtime *util.RuntimeOptions) (_result *RemovePermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.TaskCreator)) {
		body["taskCreator"] = request.TaskCreator
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
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
		Action:      tea.String("RemovePermission"),
		Version:     tea.String("transcribe_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/transcribe/tasks/" + tea.StringValue(taskUuid) + "/permissions/remove"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RemovePermissionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移除指定用户对闪记任务的权限
//
// @param request - RemovePermissionRequest
//
// @return RemovePermissionResponse
func (client *Client) RemovePermission(taskUuid *string, request *RemovePermissionRequest) (_result *RemovePermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemovePermissionHeaders{}
	_result = &RemovePermissionResponse{}
	_body, _err := client.RemovePermissionWithOptions(taskUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 针对指定的闪记，修改或者授予指定用户权限
//
// @param request - UpdatePermissionForUsersRequest
//
// @param headers - UpdatePermissionForUsersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePermissionForUsersResponse
func (client *Client) UpdatePermissionForUsersWithOptions(taskUuid *string, request *UpdatePermissionForUsersRequest, headers *UpdatePermissionForUsersHeaders, runtime *util.RuntimeOptions) (_result *UpdatePermissionForUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorUid)) {
		query["operatorUid"] = request.OperatorUid
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.TaskCreator)) {
		body["taskCreator"] = request.TaskCreator
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
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePermissionForUsers"),
		Version:     tea.String("transcribe_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/transcribe/tasks/" + tea.StringValue(taskUuid) + "/permissions"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePermissionForUsersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 针对指定的闪记，修改或者授予指定用户权限
//
// @param request - UpdatePermissionForUsersRequest
//
// @return UpdatePermissionForUsersResponse
func (client *Client) UpdatePermissionForUsers(taskUuid *string, request *UpdatePermissionForUsersRequest) (_result *UpdatePermissionForUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdatePermissionForUsersHeaders{}
	_result = &UpdatePermissionForUsersResponse{}
	_body, _err := client.UpdatePermissionForUsersWithOptions(taskUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
