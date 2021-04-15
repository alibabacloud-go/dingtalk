// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package contact_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetLatestDingIndexHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetLatestDingIndexHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetLatestDingIndexHeaders) GoString() string {
	return s.String()
}

func (s *GetLatestDingIndexHeaders) SetCommonHeaders(v map[string]*string) *GetLatestDingIndexHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetLatestDingIndexHeaders) SetXAcsDingtalkAccessToken(v string) *GetLatestDingIndexHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetLatestDingIndexResponseBody struct {
	// 日期
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
	// 钉钉指数
	IdxTotal *float32 `json:"idxTotal,omitempty" xml:"idxTotal,omitempty"`
	// 效率指数
	IdxEfficiency *float32 `json:"idxEfficiency,omitempty" xml:"idxEfficiency,omitempty"`
	// 绿色指数
	IdxCarbon *float32 `json:"idxCarbon,omitempty" xml:"idxCarbon,omitempty"`
	// 钉钉指数月均分
	IdxMonthlyAvg *float32 `json:"idxMonthlyAvg,omitempty" xml:"idxMonthlyAvg,omitempty"`
}

func (s GetLatestDingIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLatestDingIndexResponseBody) GoString() string {
	return s.String()
}

func (s *GetLatestDingIndexResponseBody) SetStatDate(v string) *GetLatestDingIndexResponseBody {
	s.StatDate = &v
	return s
}

func (s *GetLatestDingIndexResponseBody) SetIdxTotal(v float32) *GetLatestDingIndexResponseBody {
	s.IdxTotal = &v
	return s
}

func (s *GetLatestDingIndexResponseBody) SetIdxEfficiency(v float32) *GetLatestDingIndexResponseBody {
	s.IdxEfficiency = &v
	return s
}

func (s *GetLatestDingIndexResponseBody) SetIdxCarbon(v float32) *GetLatestDingIndexResponseBody {
	s.IdxCarbon = &v
	return s
}

func (s *GetLatestDingIndexResponseBody) SetIdxMonthlyAvg(v float32) *GetLatestDingIndexResponseBody {
	s.IdxMonthlyAvg = &v
	return s
}

type GetLatestDingIndexResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLatestDingIndexResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLatestDingIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLatestDingIndexResponse) GoString() string {
	return s.String()
}

func (s *GetLatestDingIndexResponse) SetHeaders(v map[string]*string) *GetLatestDingIndexResponse {
	s.Headers = v
	return s
}

func (s *GetLatestDingIndexResponse) SetBody(v *GetLatestDingIndexResponseBody) *GetLatestDingIndexResponse {
	s.Body = v
	return s
}

type GetUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserHeaders) GoString() string {
	return s.String()
}

func (s *GetUserHeaders) SetCommonHeaders(v map[string]*string) *GetUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserResponseBody struct {
	// 昵称
	Nick *string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 头像url
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// 手机号
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// openId
	OpenId *string `json:"openId,omitempty" xml:"openId,omitempty"`
	// unionId
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// 个人邮箱
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetNick(v string) *GetUserResponseBody {
	s.Nick = &v
	return s
}

func (s *GetUserResponseBody) SetAvatarUrl(v string) *GetUserResponseBody {
	s.AvatarUrl = &v
	return s
}

func (s *GetUserResponseBody) SetMobile(v string) *GetUserResponseBody {
	s.Mobile = &v
	return s
}

func (s *GetUserResponseBody) SetOpenId(v string) *GetUserResponseBody {
	s.OpenId = &v
	return s
}

func (s *GetUserResponseBody) SetUnionId(v string) *GetUserResponseBody {
	s.UnionId = &v
	return s
}

func (s *GetUserResponseBody) SetEmail(v string) *GetUserResponseBody {
	s.Email = &v
	return s
}

type GetUserResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetHeaders(v map[string]*string) *GetUserResponse {
	s.Headers = v
	return s
}

func (s *GetUserResponse) SetBody(v *GetUserResponseBody) *GetUserResponse {
	s.Body = v
	return s
}

type GetApplyInviteInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetApplyInviteInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetApplyInviteInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetApplyInviteInfoHeaders) SetCommonHeaders(v map[string]*string) *GetApplyInviteInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetApplyInviteInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetApplyInviteInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetApplyInviteInfoRequest struct {
	// 邀请者userId
	InviterUserId *string `json:"inviterUserId,omitempty" xml:"inviterUserId,omitempty"`
	// 获取部门邀请链接的部门ID
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s GetApplyInviteInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetApplyInviteInfoRequest) GoString() string {
	return s.String()
}

func (s *GetApplyInviteInfoRequest) SetInviterUserId(v string) *GetApplyInviteInfoRequest {
	s.InviterUserId = &v
	return s
}

func (s *GetApplyInviteInfoRequest) SetDeptId(v int64) *GetApplyInviteInfoRequest {
	s.DeptId = &v
	return s
}

type GetApplyInviteInfoResponseBody struct {
	// 是否开启邀请
	InviteSwitch *bool `json:"inviteSwitch,omitempty" xml:"inviteSwitch,omitempty"`
	// 是否开启通过企业名称搜索申请
	SearchNameInvite *bool `json:"searchNameInvite,omitempty" xml:"searchNameInvite,omitempty"`
	// 是否开启通过团队号申请加入
	OrgApplyCodeInvite *bool `json:"orgApplyCodeInvite,omitempty" xml:"orgApplyCodeInvite,omitempty"`
	// 是否开启通过链接邀请加入
	LinkInvite *bool `json:"linkInvite,omitempty" xml:"linkInvite,omitempty"`
	// 邀请链接
	InviteUrl *string `json:"inviteUrl,omitempty" xml:"inviteUrl,omitempty"`
	// 仅部门邀请有效： 0-无需审核 1-有权限的子管理员
	AuditType *int64 `json:"auditType,omitempty" xml:"auditType,omitempty"`
	// 是否允许员工扫码加入部门，仅在无需审核的时候有效（已经在组织内的成员通过扫描部门二维码加入某个部门）
	EmpApplyJoinDept *bool `json:"empApplyJoinDept,omitempty" xml:"empApplyJoinDept,omitempty"`
}

func (s GetApplyInviteInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApplyInviteInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplyInviteInfoResponseBody) SetInviteSwitch(v bool) *GetApplyInviteInfoResponseBody {
	s.InviteSwitch = &v
	return s
}

func (s *GetApplyInviteInfoResponseBody) SetSearchNameInvite(v bool) *GetApplyInviteInfoResponseBody {
	s.SearchNameInvite = &v
	return s
}

func (s *GetApplyInviteInfoResponseBody) SetOrgApplyCodeInvite(v bool) *GetApplyInviteInfoResponseBody {
	s.OrgApplyCodeInvite = &v
	return s
}

func (s *GetApplyInviteInfoResponseBody) SetLinkInvite(v bool) *GetApplyInviteInfoResponseBody {
	s.LinkInvite = &v
	return s
}

func (s *GetApplyInviteInfoResponseBody) SetInviteUrl(v string) *GetApplyInviteInfoResponseBody {
	s.InviteUrl = &v
	return s
}

func (s *GetApplyInviteInfoResponseBody) SetAuditType(v int64) *GetApplyInviteInfoResponseBody {
	s.AuditType = &v
	return s
}

func (s *GetApplyInviteInfoResponseBody) SetEmpApplyJoinDept(v bool) *GetApplyInviteInfoResponseBody {
	s.EmpApplyJoinDept = &v
	return s
}

type GetApplyInviteInfoResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetApplyInviteInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetApplyInviteInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApplyInviteInfoResponse) GoString() string {
	return s.String()
}

func (s *GetApplyInviteInfoResponse) SetHeaders(v map[string]*string) *GetApplyInviteInfoResponse {
	s.Headers = v
	return s
}

func (s *GetApplyInviteInfoResponse) SetBody(v *GetApplyInviteInfoResponseBody) *GetApplyInviteInfoResponse {
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

func (client *Client) GetLatestDingIndex() (_result *GetLatestDingIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetLatestDingIndexHeaders{}
	_result = &GetLatestDingIndexResponse{}
	_body, _err := client.GetLatestDingIndexWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLatestDingIndexWithOptions(headers *GetLatestDingIndexHeaders, runtime *util.RuntimeOptions) (_result *GetLatestDingIndexResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &GetLatestDingIndexResponse{}
	_body, _err := client.DoROARequest(tea.String("GetLatestDingIndex"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/contact/dingIndexs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUser(unionId *string) (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserHeaders{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(unionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserWithOptions(unionId *string, headers *GetUserHeaders, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &GetUserResponse{}
	_body, _err := client.DoROARequest(tea.String("GetUser"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/contact/users/"+tea.StringValue(unionId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetApplyInviteInfo(request *GetApplyInviteInfoRequest) (_result *GetApplyInviteInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetApplyInviteInfoHeaders{}
	_result = &GetApplyInviteInfoResponse{}
	_body, _err := client.GetApplyInviteInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetApplyInviteInfoWithOptions(request *GetApplyInviteInfoRequest, headers *GetApplyInviteInfoHeaders, runtime *util.RuntimeOptions) (_result *GetApplyInviteInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InviterUserId)) {
		query["inviterUserId"] = request.InviterUserId
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
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
		Query:   openapiutil.Query(query),
	}
	_result = &GetApplyInviteInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetApplyInviteInfo"), tea.String("contact_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/contact/invites/infos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
