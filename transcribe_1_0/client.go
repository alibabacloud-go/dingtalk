// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package transcribe_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
	// 闪记任务的类型。枚举值，从任务详情中获取。
	BizType *int32 `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 被授权的用户信息列表
	Members []*UpdatePermissionForUsersRequestMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// 要操作的闪记任务的创建者userId。
	TaskCreator *int64 `json:"taskCreator,omitempty" xml:"taskCreator,omitempty"`
	// 闪记任务的特定标识。是一个不定长的字符串。
	Uuid        *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	OperatorUid *int64  `json:"operatorUid,omitempty" xml:"operatorUid,omitempty"`
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

func (s *UpdatePermissionForUsersRequest) SetUuid(v string) *UpdatePermissionForUsersRequest {
	s.Uuid = &v
	return s
}

func (s *UpdatePermissionForUsersRequest) SetOperatorUid(v int64) *UpdatePermissionForUsersRequest {
	s.OperatorUid = &v
	return s
}

type UpdatePermissionForUsersRequestMembers struct {
	MemberId *int64 `json:"memberId,omitempty" xml:"memberId,omitempty"`
	// 要赋予用户的权限名称。该字段表示要授予特定用户的权限名称，由开发者填写。
	//
	// EDITOR：可编辑权限
	//
	// READ_DOWNLOAD：仅可查看、下载
	//
	// READ：只读权限
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdatePermissionForUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

func (client *Client) UpdatePermissionForUsers(taskId *string, request *UpdatePermissionForUsersRequest) (_result *UpdatePermissionForUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdatePermissionForUsersHeaders{}
	_result = &UpdatePermissionForUsersResponse{}
	_body, _err := client.UpdatePermissionForUsersWithOptions(taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePermissionForUsersWithOptions(taskId *string, request *UpdatePermissionForUsersRequest, headers *UpdatePermissionForUsersHeaders, runtime *util.RuntimeOptions) (_result *UpdatePermissionForUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	taskId = openapiutil.GetEncodeParam(taskId)
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
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdatePermissionForUsersResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdatePermissionForUsers"), tea.String("transcribe_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/transcribe/tasks/"+tea.StringValue(taskId)+"/permissions"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
