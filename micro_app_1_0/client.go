// This file is auto-generated, don't edit it. Thanks.
package micro_app_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddAppRolesToMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddAppRolesToMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddAppRolesToMemberHeaders) GoString() string {
	return s.String()
}

func (s *AddAppRolesToMemberHeaders) SetCommonHeaders(v map[string]*string) *AddAppRolesToMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddAppRolesToMemberHeaders) SetXAcsDingtalkAccessToken(v string) *AddAppRolesToMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddAppRolesToMemberRequest struct {
	// This parameter is required.
	MemberId *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	// This parameter is required.
	MemberType *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// This parameter is required.
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// This parameter is required.
	RoleList []*AddAppRolesToMemberRequestRoleList `json:"roleList,omitempty" xml:"roleList,omitempty" type:"Repeated"`
}

func (s AddAppRolesToMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAppRolesToMemberRequest) GoString() string {
	return s.String()
}

func (s *AddAppRolesToMemberRequest) SetMemberId(v string) *AddAppRolesToMemberRequest {
	s.MemberId = &v
	return s
}

func (s *AddAppRolesToMemberRequest) SetMemberType(v string) *AddAppRolesToMemberRequest {
	s.MemberType = &v
	return s
}

func (s *AddAppRolesToMemberRequest) SetOpUserId(v string) *AddAppRolesToMemberRequest {
	s.OpUserId = &v
	return s
}

func (s *AddAppRolesToMemberRequest) SetRoleList(v []*AddAppRolesToMemberRequestRoleList) *AddAppRolesToMemberRequest {
	s.RoleList = v
	return s
}

type AddAppRolesToMemberRequestRoleList struct {
	// This parameter is required.
	RoleId *int64 `json:"roleId,omitempty" xml:"roleId,omitempty"`
	// This parameter is required.
	ScopeVersion *int64 `json:"scopeVersion,omitempty" xml:"scopeVersion,omitempty"`
}

func (s AddAppRolesToMemberRequestRoleList) String() string {
	return tea.Prettify(s)
}

func (s AddAppRolesToMemberRequestRoleList) GoString() string {
	return s.String()
}

func (s *AddAppRolesToMemberRequestRoleList) SetRoleId(v int64) *AddAppRolesToMemberRequestRoleList {
	s.RoleId = &v
	return s
}

func (s *AddAppRolesToMemberRequestRoleList) SetScopeVersion(v int64) *AddAppRolesToMemberRequestRoleList {
	s.ScopeVersion = &v
	return s
}

type AddAppRolesToMemberResponseBody struct {
	Result []*AddAppRolesToMemberResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s AddAppRolesToMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddAppRolesToMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddAppRolesToMemberResponseBody) SetResult(v []*AddAppRolesToMemberResponseBodyResult) *AddAppRolesToMemberResponseBody {
	s.Result = v
	return s
}

type AddAppRolesToMemberResponseBodyResult struct {
	// example:
	//
	// 123
	LatestScopeVersion *int64 `json:"latestScopeVersion,omitempty" xml:"latestScopeVersion,omitempty"`
	// example:
	//
	// 123
	RoleId *int64 `json:"roleId,omitempty" xml:"roleId,omitempty"`
	// example:
	//
	// userNoPrivilegeToManageApp
	SubErrorCode *string `json:"subErrorCode,omitempty" xml:"subErrorCode,omitempty"`
	// example:
	//
	// 传入的角色范围数据版本号不合法
	SubErrorMsg *string `json:"subErrorMsg,omitempty" xml:"subErrorMsg,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddAppRolesToMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddAppRolesToMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddAppRolesToMemberResponseBodyResult) SetLatestScopeVersion(v int64) *AddAppRolesToMemberResponseBodyResult {
	s.LatestScopeVersion = &v
	return s
}

func (s *AddAppRolesToMemberResponseBodyResult) SetRoleId(v int64) *AddAppRolesToMemberResponseBodyResult {
	s.RoleId = &v
	return s
}

func (s *AddAppRolesToMemberResponseBodyResult) SetSubErrorCode(v string) *AddAppRolesToMemberResponseBodyResult {
	s.SubErrorCode = &v
	return s
}

func (s *AddAppRolesToMemberResponseBodyResult) SetSubErrorMsg(v string) *AddAppRolesToMemberResponseBodyResult {
	s.SubErrorMsg = &v
	return s
}

func (s *AddAppRolesToMemberResponseBodyResult) SetSuccess(v bool) *AddAppRolesToMemberResponseBodyResult {
	s.Success = &v
	return s
}

type AddAppRolesToMemberResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAppRolesToMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAppRolesToMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAppRolesToMemberResponse) GoString() string {
	return s.String()
}

func (s *AddAppRolesToMemberResponse) SetHeaders(v map[string]*string) *AddAppRolesToMemberResponse {
	s.Headers = v
	return s
}

func (s *AddAppRolesToMemberResponse) SetStatusCode(v int32) *AddAppRolesToMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAppRolesToMemberResponse) SetBody(v *AddAppRolesToMemberResponseBody) *AddAppRolesToMemberResponse {
	s.Body = v
	return s
}

type AddAppToWorkBenchGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddAppToWorkBenchGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddAppToWorkBenchGroupHeaders) GoString() string {
	return s.String()
}

func (s *AddAppToWorkBenchGroupHeaders) SetCommonHeaders(v map[string]*string) *AddAppToWorkBenchGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddAppToWorkBenchGroupHeaders) SetXAcsDingtalkAccessToken(v string) *AddAppToWorkBenchGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddAppToWorkBenchGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// namexx
	ComponentId *string `json:"componentId,omitempty" xml:"componentId,omitempty"`
	// example:
	//
	// corpxxxx
	EcologicalCorpId *string `json:"ecologicalCorpId,omitempty" xml:"ecologicalCorpId,omitempty"`
	// example:
	//
	// xxxx
	OpUnionId *string `json:"opUnionId,omitempty" xml:"opUnionId,omitempty"`
}

func (s AddAppToWorkBenchGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAppToWorkBenchGroupRequest) GoString() string {
	return s.String()
}

func (s *AddAppToWorkBenchGroupRequest) SetComponentId(v string) *AddAppToWorkBenchGroupRequest {
	s.ComponentId = &v
	return s
}

func (s *AddAppToWorkBenchGroupRequest) SetEcologicalCorpId(v string) *AddAppToWorkBenchGroupRequest {
	s.EcologicalCorpId = &v
	return s
}

func (s *AddAppToWorkBenchGroupRequest) SetOpUnionId(v string) *AddAppToWorkBenchGroupRequest {
	s.OpUnionId = &v
	return s
}

type AddAppToWorkBenchGroupResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s AddAppToWorkBenchGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddAppToWorkBenchGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddAppToWorkBenchGroupResponseBody) SetResult(v bool) *AddAppToWorkBenchGroupResponseBody {
	s.Result = &v
	return s
}

type AddAppToWorkBenchGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAppToWorkBenchGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAppToWorkBenchGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAppToWorkBenchGroupResponse) GoString() string {
	return s.String()
}

func (s *AddAppToWorkBenchGroupResponse) SetHeaders(v map[string]*string) *AddAppToWorkBenchGroupResponse {
	s.Headers = v
	return s
}

func (s *AddAppToWorkBenchGroupResponse) SetStatusCode(v int32) *AddAppToWorkBenchGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAppToWorkBenchGroupResponse) SetBody(v *AddAppToWorkBenchGroupResponseBody) *AddAppToWorkBenchGroupResponse {
	s.Body = v
	return s
}

type AddMemberToAppRoleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddMemberToAppRoleHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddMemberToAppRoleHeaders) GoString() string {
	return s.String()
}

func (s *AddMemberToAppRoleHeaders) SetCommonHeaders(v map[string]*string) *AddMemberToAppRoleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddMemberToAppRoleHeaders) SetXAcsDingtalkAccessToken(v string) *AddMemberToAppRoleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddMemberToAppRoleRequest struct {
	DeptIdList []*int64 `json:"deptIdList,omitempty" xml:"deptIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// This parameter is required.
	ScopeVersion *int64    `json:"scopeVersion,omitempty" xml:"scopeVersion,omitempty"`
	UserIdList   []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s AddMemberToAppRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s AddMemberToAppRoleRequest) GoString() string {
	return s.String()
}

func (s *AddMemberToAppRoleRequest) SetDeptIdList(v []*int64) *AddMemberToAppRoleRequest {
	s.DeptIdList = v
	return s
}

func (s *AddMemberToAppRoleRequest) SetOpUserId(v string) *AddMemberToAppRoleRequest {
	s.OpUserId = &v
	return s
}

func (s *AddMemberToAppRoleRequest) SetScopeVersion(v int64) *AddMemberToAppRoleRequest {
	s.ScopeVersion = &v
	return s
}

func (s *AddMemberToAppRoleRequest) SetUserIdList(v []*string) *AddMemberToAppRoleRequest {
	s.UserIdList = v
	return s
}

type AddMemberToAppRoleResponseBody struct {
	// example:
	//
	// 123
	LatestScopeVersion *int64 `json:"latestScopeVersion,omitempty" xml:"latestScopeVersion,omitempty"`
}

func (s AddMemberToAppRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddMemberToAppRoleResponseBody) GoString() string {
	return s.String()
}

func (s *AddMemberToAppRoleResponseBody) SetLatestScopeVersion(v int64) *AddMemberToAppRoleResponseBody {
	s.LatestScopeVersion = &v
	return s
}

type AddMemberToAppRoleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMemberToAppRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMemberToAppRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s AddMemberToAppRoleResponse) GoString() string {
	return s.String()
}

func (s *AddMemberToAppRoleResponse) SetHeaders(v map[string]*string) *AddMemberToAppRoleResponse {
	s.Headers = v
	return s
}

func (s *AddMemberToAppRoleResponse) SetStatusCode(v int32) *AddMemberToAppRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMemberToAppRoleResponse) SetBody(v *AddMemberToAppRoleResponseBody) *AddMemberToAppRoleResponse {
	s.Body = v
	return s
}

type AnheiPResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s AnheiPResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AnheiPResponseBody) GoString() string {
	return s.String()
}

func (s *AnheiPResponseBody) SetResult(v string) *AnheiPResponseBody {
	s.Result = &v
	return s
}

type AnheiPResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AnheiPResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AnheiPResponse) String() string {
	return tea.Prettify(s)
}

func (s AnheiPResponse) GoString() string {
	return s.String()
}

func (s *AnheiPResponse) SetHeaders(v map[string]*string) *AnheiPResponse {
	s.Headers = v
	return s
}

func (s *AnheiPResponse) SetStatusCode(v int32) *AnheiPResponse {
	s.StatusCode = &v
	return s
}

func (s *AnheiPResponse) SetBody(v *AnheiPResponseBody) *AnheiPResponse {
	s.Body = v
	return s
}

type AnheiTest888ResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AnheiTest888ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AnheiTest888ResponseBody) GoString() string {
	return s.String()
}

func (s *AnheiTest888ResponseBody) SetRequestId(v string) *AnheiTest888ResponseBody {
	s.RequestId = &v
	return s
}

type AnheiTest888Response struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AnheiTest888ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AnheiTest888Response) String() string {
	return tea.Prettify(s)
}

func (s AnheiTest888Response) GoString() string {
	return s.String()
}

func (s *AnheiTest888Response) SetHeaders(v map[string]*string) *AnheiTest888Response {
	s.Headers = v
	return s
}

func (s *AnheiTest888Response) SetStatusCode(v int32) *AnheiTest888Response {
	s.StatusCode = &v
	return s
}

func (s *AnheiTest888Response) SetBody(v *AnheiTest888ResponseBody) *AnheiTest888Response {
	s.Body = v
	return s
}

type AnheiTestBResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AnheiTestBResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AnheiTestBResponseBody) GoString() string {
	return s.String()
}

func (s *AnheiTestBResponseBody) SetRequestId(v string) *AnheiTestBResponseBody {
	s.RequestId = &v
	return s
}

type AnheiTestBResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AnheiTestBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AnheiTestBResponse) String() string {
	return tea.Prettify(s)
}

func (s AnheiTestBResponse) GoString() string {
	return s.String()
}

func (s *AnheiTestBResponse) SetHeaders(v map[string]*string) *AnheiTestBResponse {
	s.Headers = v
	return s
}

func (s *AnheiTestBResponse) SetStatusCode(v int32) *AnheiTestBResponse {
	s.StatusCode = &v
	return s
}

func (s *AnheiTestBResponse) SetBody(v *AnheiTestBResponseBody) *AnheiTestBResponse {
	s.Body = v
	return s
}

type AnheiTestNineResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AnheiTestNineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AnheiTestNineResponseBody) GoString() string {
	return s.String()
}

func (s *AnheiTestNineResponseBody) SetRequestId(v string) *AnheiTestNineResponseBody {
	s.RequestId = &v
	return s
}

type AnheiTestNineResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AnheiTestNineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AnheiTestNineResponse) String() string {
	return tea.Prettify(s)
}

func (s AnheiTestNineResponse) GoString() string {
	return s.String()
}

func (s *AnheiTestNineResponse) SetHeaders(v map[string]*string) *AnheiTestNineResponse {
	s.Headers = v
	return s
}

func (s *AnheiTestNineResponse) SetStatusCode(v int32) *AnheiTestNineResponse {
	s.StatusCode = &v
	return s
}

func (s *AnheiTestNineResponse) SetBody(v *AnheiTestNineResponseBody) *AnheiTestNineResponse {
	s.Body = v
	return s
}

type AppStatusManagerTestRequest struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AppStatusManagerTestRequest) String() string {
	return tea.Prettify(s)
}

func (s AppStatusManagerTestRequest) GoString() string {
	return s.String()
}

func (s *AppStatusManagerTestRequest) SetRequestId(v string) *AppStatusManagerTestRequest {
	s.RequestId = &v
	return s
}

type AppStatusManagerTestResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AppStatusManagerTestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppStatusManagerTestResponseBody) GoString() string {
	return s.String()
}

func (s *AppStatusManagerTestResponseBody) SetRequestId(v string) *AppStatusManagerTestResponseBody {
	s.RequestId = &v
	return s
}

type AppStatusManagerTestResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AppStatusManagerTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AppStatusManagerTestResponse) String() string {
	return tea.Prettify(s)
}

func (s AppStatusManagerTestResponse) GoString() string {
	return s.String()
}

func (s *AppStatusManagerTestResponse) SetHeaders(v map[string]*string) *AppStatusManagerTestResponse {
	s.Headers = v
	return s
}

func (s *AppStatusManagerTestResponse) SetStatusCode(v int32) *AppStatusManagerTestResponse {
	s.StatusCode = &v
	return s
}

func (s *AppStatusManagerTestResponse) SetBody(v *AppStatusManagerTestResponseBody) *AppStatusManagerTestResponse {
	s.Body = v
	return s
}

type AyunTestResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AyunTestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AyunTestResponseBody) GoString() string {
	return s.String()
}

func (s *AyunTestResponseBody) SetRequestId(v string) *AyunTestResponseBody {
	s.RequestId = &v
	return s
}

type AyunTestResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AyunTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AyunTestResponse) String() string {
	return tea.Prettify(s)
}

func (s AyunTestResponse) GoString() string {
	return s.String()
}

func (s *AyunTestResponse) SetHeaders(v map[string]*string) *AyunTestResponse {
	s.Headers = v
	return s
}

func (s *AyunTestResponse) SetStatusCode(v int32) *AyunTestResponse {
	s.StatusCode = &v
	return s
}

func (s *AyunTestResponse) SetBody(v *AyunTestResponseBody) *AyunTestResponse {
	s.Body = v
	return s
}

type AyunTestOnlineResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AyunTestOnlineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AyunTestOnlineResponseBody) GoString() string {
	return s.String()
}

func (s *AyunTestOnlineResponseBody) SetRequestId(v string) *AyunTestOnlineResponseBody {
	s.RequestId = &v
	return s
}

type AyunTestOnlineResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AyunTestOnlineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AyunTestOnlineResponse) String() string {
	return tea.Prettify(s)
}

func (s AyunTestOnlineResponse) GoString() string {
	return s.String()
}

func (s *AyunTestOnlineResponse) SetHeaders(v map[string]*string) *AyunTestOnlineResponse {
	s.Headers = v
	return s
}

func (s *AyunTestOnlineResponse) SetStatusCode(v int32) *AyunTestOnlineResponse {
	s.StatusCode = &v
	return s
}

func (s *AyunTestOnlineResponse) SetBody(v *AyunTestOnlineResponseBody) *AyunTestOnlineResponse {
	s.Body = v
	return s
}

type CreateApaasAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateApaasAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateApaasAppHeaders) GoString() string {
	return s.String()
}

func (s *CreateApaasAppHeaders) SetCommonHeaders(v map[string]*string) *CreateApaasAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateApaasAppHeaders) SetXAcsDingtalkAccessToken(v string) *CreateApaasAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateApaasAppRequest struct {
	// This parameter is required.
	AppDesc *string `json:"appDesc,omitempty" xml:"appDesc,omitempty"`
	// This parameter is required.
	AppIcon *string `json:"appIcon,omitempty" xml:"appIcon,omitempty"`
	// This parameter is required.
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// This parameter is required.
	BizAppId         *string `json:"bizAppId,omitempty" xml:"bizAppId,omitempty"`
	HomepageEditLink *string `json:"homepageEditLink,omitempty" xml:"homepageEditLink,omitempty"`
	// This parameter is required.
	HomepageLink *string `json:"homepageLink,omitempty" xml:"homepageLink,omitempty"`
	IsShortCut   *int32  `json:"isShortCut,omitempty" xml:"isShortCut,omitempty"`
	OmpLink      *string `json:"ompLink,omitempty" xml:"ompLink,omitempty"`
	// This parameter is required.
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// This parameter is required.
	PcHomepageEditLink *string `json:"pcHomepageEditLink,omitempty" xml:"pcHomepageEditLink,omitempty"`
	// This parameter is required.
	PcHomepageLink *string `json:"pcHomepageLink,omitempty" xml:"pcHomepageLink,omitempty"`
	TemplateKey    *string `json:"templateKey,omitempty" xml:"templateKey,omitempty"`
}

func (s CreateApaasAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApaasAppRequest) GoString() string {
	return s.String()
}

func (s *CreateApaasAppRequest) SetAppDesc(v string) *CreateApaasAppRequest {
	s.AppDesc = &v
	return s
}

func (s *CreateApaasAppRequest) SetAppIcon(v string) *CreateApaasAppRequest {
	s.AppIcon = &v
	return s
}

func (s *CreateApaasAppRequest) SetAppName(v string) *CreateApaasAppRequest {
	s.AppName = &v
	return s
}

func (s *CreateApaasAppRequest) SetBizAppId(v string) *CreateApaasAppRequest {
	s.BizAppId = &v
	return s
}

func (s *CreateApaasAppRequest) SetHomepageEditLink(v string) *CreateApaasAppRequest {
	s.HomepageEditLink = &v
	return s
}

func (s *CreateApaasAppRequest) SetHomepageLink(v string) *CreateApaasAppRequest {
	s.HomepageLink = &v
	return s
}

func (s *CreateApaasAppRequest) SetIsShortCut(v int32) *CreateApaasAppRequest {
	s.IsShortCut = &v
	return s
}

func (s *CreateApaasAppRequest) SetOmpLink(v string) *CreateApaasAppRequest {
	s.OmpLink = &v
	return s
}

func (s *CreateApaasAppRequest) SetOpUserId(v string) *CreateApaasAppRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateApaasAppRequest) SetPcHomepageEditLink(v string) *CreateApaasAppRequest {
	s.PcHomepageEditLink = &v
	return s
}

func (s *CreateApaasAppRequest) SetPcHomepageLink(v string) *CreateApaasAppRequest {
	s.PcHomepageLink = &v
	return s
}

func (s *CreateApaasAppRequest) SetTemplateKey(v string) *CreateApaasAppRequest {
	s.TemplateKey = &v
	return s
}

type CreateApaasAppResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// This parameter is required.
	BizAppId *string `json:"bizAppId,omitempty" xml:"bizAppId,omitempty"`
}

func (s CreateApaasAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApaasAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApaasAppResponseBody) SetAgentId(v int64) *CreateApaasAppResponseBody {
	s.AgentId = &v
	return s
}

func (s *CreateApaasAppResponseBody) SetBizAppId(v string) *CreateApaasAppResponseBody {
	s.BizAppId = &v
	return s
}

type CreateApaasAppResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApaasAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApaasAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApaasAppResponse) GoString() string {
	return s.String()
}

func (s *CreateApaasAppResponse) SetHeaders(v map[string]*string) *CreateApaasAppResponse {
	s.Headers = v
	return s
}

func (s *CreateApaasAppResponse) SetStatusCode(v int32) *CreateApaasAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApaasAppResponse) SetBody(v *CreateApaasAppResponseBody) *CreateApaasAppResponse {
	s.Body = v
	return s
}

type CreateInnerAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateInnerAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateInnerAppHeaders) GoString() string {
	return s.String()
}

func (s *CreateInnerAppHeaders) SetCommonHeaders(v map[string]*string) *CreateInnerAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateInnerAppHeaders) SetXAcsDingtalkAccessToken(v string) *CreateInnerAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateInnerAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// descxxx
	Desc        *string `json:"desc,omitempty" xml:"desc,omitempty"`
	DevelopType *int32  `json:"developType,omitempty" xml:"developType,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	HomepageLink *string `json:"homepageLink,omitempty" xml:"homepageLink,omitempty"`
	// example:
	//
	// mediaxxx
	Icon        *string   `json:"icon,omitempty" xml:"icon,omitempty"`
	IpWhiteList []*string `json:"ipWhiteList,omitempty" xml:"ipWhiteList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// namexx
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	OmpLink *string `json:"ompLink,omitempty" xml:"ompLink,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	OpUnionId *string `json:"opUnionId,omitempty" xml:"opUnionId,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	PcHomepageLink *string `json:"pcHomepageLink,omitempty" xml:"pcHomepageLink,omitempty"`
	// example:
	//
	// BASE
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
}

func (s CreateInnerAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInnerAppRequest) GoString() string {
	return s.String()
}

func (s *CreateInnerAppRequest) SetDesc(v string) *CreateInnerAppRequest {
	s.Desc = &v
	return s
}

func (s *CreateInnerAppRequest) SetDevelopType(v int32) *CreateInnerAppRequest {
	s.DevelopType = &v
	return s
}

func (s *CreateInnerAppRequest) SetHomepageLink(v string) *CreateInnerAppRequest {
	s.HomepageLink = &v
	return s
}

func (s *CreateInnerAppRequest) SetIcon(v string) *CreateInnerAppRequest {
	s.Icon = &v
	return s
}

func (s *CreateInnerAppRequest) SetIpWhiteList(v []*string) *CreateInnerAppRequest {
	s.IpWhiteList = v
	return s
}

func (s *CreateInnerAppRequest) SetName(v string) *CreateInnerAppRequest {
	s.Name = &v
	return s
}

func (s *CreateInnerAppRequest) SetOmpLink(v string) *CreateInnerAppRequest {
	s.OmpLink = &v
	return s
}

func (s *CreateInnerAppRequest) SetOpUnionId(v string) *CreateInnerAppRequest {
	s.OpUnionId = &v
	return s
}

func (s *CreateInnerAppRequest) SetPcHomepageLink(v string) *CreateInnerAppRequest {
	s.PcHomepageLink = &v
	return s
}

func (s *CreateInnerAppRequest) SetScopeType(v string) *CreateInnerAppRequest {
	s.ScopeType = &v
	return s
}

type CreateInnerAppResponseBody struct {
	// example:
	//
	// 111
	AgentId   *int64  `json:"agentId,omitempty" xml:"agentId,omitempty"`
	AppKey    *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	AppSecret *string `json:"appSecret,omitempty" xml:"appSecret,omitempty"`
}

func (s CreateInnerAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInnerAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInnerAppResponseBody) SetAgentId(v int64) *CreateInnerAppResponseBody {
	s.AgentId = &v
	return s
}

func (s *CreateInnerAppResponseBody) SetAppKey(v string) *CreateInnerAppResponseBody {
	s.AppKey = &v
	return s
}

func (s *CreateInnerAppResponseBody) SetAppSecret(v string) *CreateInnerAppResponseBody {
	s.AppSecret = &v
	return s
}

type CreateInnerAppResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInnerAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInnerAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInnerAppResponse) GoString() string {
	return s.String()
}

func (s *CreateInnerAppResponse) SetHeaders(v map[string]*string) *CreateInnerAppResponse {
	s.Headers = v
	return s
}

func (s *CreateInnerAppResponse) SetStatusCode(v int32) *CreateInnerAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInnerAppResponse) SetBody(v *CreateInnerAppResponseBody) *CreateInnerAppResponse {
	s.Body = v
	return s
}

type DeleteAppRoleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteAppRoleHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppRoleHeaders) GoString() string {
	return s.String()
}

func (s *DeleteAppRoleHeaders) SetCommonHeaders(v map[string]*string) *DeleteAppRoleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteAppRoleHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteAppRoleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteAppRoleRequest struct {
	// This parameter is required.
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s DeleteAppRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppRoleRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppRoleRequest) SetOpUserId(v string) *DeleteAppRoleRequest {
	s.OpUserId = &v
	return s
}

type DeleteAppRoleResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteAppRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppRoleResponseBody) SetResult(v bool) *DeleteAppRoleResponseBody {
	s.Result = &v
	return s
}

type DeleteAppRoleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppRoleResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppRoleResponse) SetHeaders(v map[string]*string) *DeleteAppRoleResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppRoleResponse) SetStatusCode(v int32) *DeleteAppRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppRoleResponse) SetBody(v *DeleteAppRoleResponseBody) *DeleteAppRoleResponse {
	s.Body = v
	return s
}

type DeleteInnerAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteInnerAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteInnerAppHeaders) GoString() string {
	return s.String()
}

func (s *DeleteInnerAppHeaders) SetCommonHeaders(v map[string]*string) *DeleteInnerAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteInnerAppHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteInnerAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteInnerAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	OpUnionId *string `json:"opUnionId,omitempty" xml:"opUnionId,omitempty"`
}

func (s DeleteInnerAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInnerAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteInnerAppRequest) SetOpUnionId(v string) *DeleteInnerAppRequest {
	s.OpUnionId = &v
	return s
}

type DeleteInnerAppResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteInnerAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInnerAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInnerAppResponseBody) SetResult(v bool) *DeleteInnerAppResponseBody {
	s.Result = &v
	return s
}

type DeleteInnerAppResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInnerAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInnerAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInnerAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteInnerAppResponse) SetHeaders(v map[string]*string) *DeleteInnerAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteInnerAppResponse) SetStatusCode(v int32) *DeleteInnerAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInnerAppResponse) SetBody(v *DeleteInnerAppResponseBody) *DeleteInnerAppResponse {
	s.Body = v
	return s
}

type GetApaasAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetApaasAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetApaasAppHeaders) GoString() string {
	return s.String()
}

func (s *GetApaasAppHeaders) SetCommonHeaders(v map[string]*string) *GetApaasAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetApaasAppHeaders) SetXAcsDingtalkAccessToken(v string) *GetApaasAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetApaasAppResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// This parameter is required.
	BizAppId *string `json:"bizAppId,omitempty" xml:"bizAppId,omitempty"`
	// This parameter is required.
	PublishStatus *string `json:"publishStatus,omitempty" xml:"publishStatus,omitempty"`
}

func (s GetApaasAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApaasAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetApaasAppResponseBody) SetAgentId(v int64) *GetApaasAppResponseBody {
	s.AgentId = &v
	return s
}

func (s *GetApaasAppResponseBody) SetBizAppId(v string) *GetApaasAppResponseBody {
	s.BizAppId = &v
	return s
}

func (s *GetApaasAppResponseBody) SetPublishStatus(v string) *GetApaasAppResponseBody {
	s.PublishStatus = &v
	return s
}

type GetApaasAppResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApaasAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApaasAppResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApaasAppResponse) GoString() string {
	return s.String()
}

func (s *GetApaasAppResponse) SetHeaders(v map[string]*string) *GetApaasAppResponse {
	s.Headers = v
	return s
}

func (s *GetApaasAppResponse) SetStatusCode(v int32) *GetApaasAppResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApaasAppResponse) SetBody(v *GetApaasAppResponseBody) *GetApaasAppResponse {
	s.Body = v
	return s
}

type GetAppResourceUseInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAppResourceUseInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAppResourceUseInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetAppResourceUseInfoHeaders) SetCommonHeaders(v map[string]*string) *GetAppResourceUseInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAppResourceUseInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetAppResourceUseInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAppResourceUseInfoRequest struct {
	// example:
	//
	// api_count
	BenefitCode *string `json:"benefitCode,omitempty" xml:"benefitCode,omitempty"`
	// example:
	//
	// 202302
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// month
	PeriodType *string `json:"periodType,omitempty" xml:"periodType,omitempty"`
	// example:
	//
	// 202301
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetAppResourceUseInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppResourceUseInfoRequest) GoString() string {
	return s.String()
}

func (s *GetAppResourceUseInfoRequest) SetBenefitCode(v string) *GetAppResourceUseInfoRequest {
	s.BenefitCode = &v
	return s
}

func (s *GetAppResourceUseInfoRequest) SetEndTime(v string) *GetAppResourceUseInfoRequest {
	s.EndTime = &v
	return s
}

func (s *GetAppResourceUseInfoRequest) SetPeriodType(v string) *GetAppResourceUseInfoRequest {
	s.PeriodType = &v
	return s
}

func (s *GetAppResourceUseInfoRequest) SetStartTime(v string) *GetAppResourceUseInfoRequest {
	s.StartTime = &v
	return s
}

type GetAppResourceUseInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*GetAppResourceUseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s GetAppResourceUseInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppResourceUseInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAppResourceUseInfoResponse) SetHeaders(v map[string]*string) *GetAppResourceUseInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAppResourceUseInfoResponse) SetStatusCode(v int32) *GetAppResourceUseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppResourceUseInfoResponse) SetBody(v []*GetAppResourceUseInfoResponseBody) *GetAppResourceUseInfoResponse {
	s.Body = v
	return s
}

type GetAppResourceUseInfoResponseBody struct {
	// example:
	//
	// 202301
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// example:
	//
	// 8511
	UsedNum *int64 `json:"usedNum,omitempty" xml:"usedNum,omitempty"`
	// example:
	//
	// 10000
	QuotaNum *int64 `json:"quotaNum,omitempty" xml:"quotaNum,omitempty"`
}

func (s GetAppResourceUseInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppResourceUseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppResourceUseInfoResponseBody) SetPeriod(v string) *GetAppResourceUseInfoResponseBody {
	s.Period = &v
	return s
}

func (s *GetAppResourceUseInfoResponseBody) SetUsedNum(v int64) *GetAppResourceUseInfoResponseBody {
	s.UsedNum = &v
	return s
}

func (s *GetAppResourceUseInfoResponseBody) SetQuotaNum(v int64) *GetAppResourceUseInfoResponseBody {
	s.QuotaNum = &v
	return s
}

type GetAppRoleScopeByRoleIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAppRoleScopeByRoleIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAppRoleScopeByRoleIdHeaders) GoString() string {
	return s.String()
}

func (s *GetAppRoleScopeByRoleIdHeaders) SetCommonHeaders(v map[string]*string) *GetAppRoleScopeByRoleIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAppRoleScopeByRoleIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetAppRoleScopeByRoleIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAppRoleScopeByRoleIdResponseBody struct {
	// example:
	//
	// false
	CanManageRole *bool    `json:"canManageRole,omitempty" xml:"canManageRole,omitempty"`
	DeptIdList    []*int64 `json:"deptIdList,omitempty" xml:"deptIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 123
	RoleId *int64 `json:"roleId,omitempty" xml:"roleId,omitempty"`
	// example:
	//
	// 财务
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	// example:
	//
	// PART_VISIBLE
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	// example:
	//
	// 123
	ScopeVersion *string   `json:"scopeVersion,omitempty" xml:"scopeVersion,omitempty"`
	UserIdList   []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s GetAppRoleScopeByRoleIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppRoleScopeByRoleIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppRoleScopeByRoleIdResponseBody) SetCanManageRole(v bool) *GetAppRoleScopeByRoleIdResponseBody {
	s.CanManageRole = &v
	return s
}

func (s *GetAppRoleScopeByRoleIdResponseBody) SetDeptIdList(v []*int64) *GetAppRoleScopeByRoleIdResponseBody {
	s.DeptIdList = v
	return s
}

func (s *GetAppRoleScopeByRoleIdResponseBody) SetRoleId(v int64) *GetAppRoleScopeByRoleIdResponseBody {
	s.RoleId = &v
	return s
}

func (s *GetAppRoleScopeByRoleIdResponseBody) SetRoleName(v string) *GetAppRoleScopeByRoleIdResponseBody {
	s.RoleName = &v
	return s
}

func (s *GetAppRoleScopeByRoleIdResponseBody) SetScopeType(v string) *GetAppRoleScopeByRoleIdResponseBody {
	s.ScopeType = &v
	return s
}

func (s *GetAppRoleScopeByRoleIdResponseBody) SetScopeVersion(v string) *GetAppRoleScopeByRoleIdResponseBody {
	s.ScopeVersion = &v
	return s
}

func (s *GetAppRoleScopeByRoleIdResponseBody) SetUserIdList(v []*string) *GetAppRoleScopeByRoleIdResponseBody {
	s.UserIdList = v
	return s
}

type GetAppRoleScopeByRoleIdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppRoleScopeByRoleIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppRoleScopeByRoleIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppRoleScopeByRoleIdResponse) GoString() string {
	return s.String()
}

func (s *GetAppRoleScopeByRoleIdResponse) SetHeaders(v map[string]*string) *GetAppRoleScopeByRoleIdResponse {
	s.Headers = v
	return s
}

func (s *GetAppRoleScopeByRoleIdResponse) SetStatusCode(v int32) *GetAppRoleScopeByRoleIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppRoleScopeByRoleIdResponse) SetBody(v *GetAppRoleScopeByRoleIdResponseBody) *GetAppRoleScopeByRoleIdResponse {
	s.Body = v
	return s
}

type GetInnerAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetInnerAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetInnerAppHeaders) GoString() string {
	return s.String()
}

func (s *GetInnerAppHeaders) SetCommonHeaders(v map[string]*string) *GetInnerAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInnerAppHeaders) SetXAcsDingtalkAccessToken(v string) *GetInnerAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetInnerAppRequest struct {
	// example:
	//
	// corpxxxx
	EcologicalCorpId *string `json:"ecologicalCorpId,omitempty" xml:"ecologicalCorpId,omitempty"`
	// example:
	//
	// xxxx
	OpUnionId *string `json:"opUnionId,omitempty" xml:"opUnionId,omitempty"`
}

func (s GetInnerAppRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInnerAppRequest) GoString() string {
	return s.String()
}

func (s *GetInnerAppRequest) SetEcologicalCorpId(v string) *GetInnerAppRequest {
	s.EcologicalCorpId = &v
	return s
}

func (s *GetInnerAppRequest) SetOpUnionId(v string) *GetInnerAppRequest {
	s.OpUnionId = &v
	return s
}

type GetInnerAppResponseBody struct {
	// example:
	//
	// 1
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// aooxxx
	AppKey *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aaaxxxxx
	AppSecret *string `json:"appSecret,omitempty" xml:"appSecret,omitempty"`
	// example:
	//
	// desc
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	HomepageLink *string `json:"homepageLink,omitempty" xml:"homepageLink,omitempty"`
	// example:
	//
	// icon
	Icon        *string   `json:"icon,omitempty" xml:"icon,omitempty"`
	IpWhiteList []*string `json:"ipWhiteList,omitempty" xml:"ipWhiteList,omitempty" type:"Repeated"`
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	OmpLink *string `json:"ompLink,omitempty" xml:"ompLink,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	PcHomepageLink *string `json:"pcHomepageLink,omitempty" xml:"pcHomepageLink,omitempty"`
}

func (s GetInnerAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInnerAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetInnerAppResponseBody) SetAgentId(v int64) *GetInnerAppResponseBody {
	s.AgentId = &v
	return s
}

func (s *GetInnerAppResponseBody) SetAppKey(v string) *GetInnerAppResponseBody {
	s.AppKey = &v
	return s
}

func (s *GetInnerAppResponseBody) SetAppSecret(v string) *GetInnerAppResponseBody {
	s.AppSecret = &v
	return s
}

func (s *GetInnerAppResponseBody) SetDesc(v string) *GetInnerAppResponseBody {
	s.Desc = &v
	return s
}

func (s *GetInnerAppResponseBody) SetHomepageLink(v string) *GetInnerAppResponseBody {
	s.HomepageLink = &v
	return s
}

func (s *GetInnerAppResponseBody) SetIcon(v string) *GetInnerAppResponseBody {
	s.Icon = &v
	return s
}

func (s *GetInnerAppResponseBody) SetIpWhiteList(v []*string) *GetInnerAppResponseBody {
	s.IpWhiteList = v
	return s
}

func (s *GetInnerAppResponseBody) SetName(v string) *GetInnerAppResponseBody {
	s.Name = &v
	return s
}

func (s *GetInnerAppResponseBody) SetOmpLink(v string) *GetInnerAppResponseBody {
	s.OmpLink = &v
	return s
}

func (s *GetInnerAppResponseBody) SetPcHomepageLink(v string) *GetInnerAppResponseBody {
	s.PcHomepageLink = &v
	return s
}

type GetInnerAppResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInnerAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInnerAppResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInnerAppResponse) GoString() string {
	return s.String()
}

func (s *GetInnerAppResponse) SetHeaders(v map[string]*string) *GetInnerAppResponse {
	s.Headers = v
	return s
}

func (s *GetInnerAppResponse) SetStatusCode(v int32) *GetInnerAppResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInnerAppResponse) SetBody(v *GetInnerAppResponseBody) *GetInnerAppResponse {
	s.Body = v
	return s
}

type GetMicroAppScopeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetMicroAppScopeHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetMicroAppScopeHeaders) GoString() string {
	return s.String()
}

func (s *GetMicroAppScopeHeaders) SetCommonHeaders(v map[string]*string) *GetMicroAppScopeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMicroAppScopeHeaders) SetXAcsDingtalkAccessToken(v string) *GetMicroAppScopeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetMicroAppScopeResponseBody struct {
	// example:
	//
	// true
	Result *GetMicroAppScopeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetMicroAppScopeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMicroAppScopeResponseBody) GoString() string {
	return s.String()
}

func (s *GetMicroAppScopeResponseBody) SetResult(v *GetMicroAppScopeResponseBodyResult) *GetMicroAppScopeResponseBody {
	s.Result = v
	return s
}

type GetMicroAppScopeResponseBodyResult struct {
	// This parameter is required.
	DeptIds []*int64 `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	// This parameter is required.
	OnlyAdminVisible *bool `json:"onlyAdminVisible,omitempty" xml:"onlyAdminVisible,omitempty"`
	// This parameter is required.
	RoleIds []*int64 `json:"roleIds,omitempty" xml:"roleIds,omitempty" type:"Repeated"`
	// This parameter is required.
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s GetMicroAppScopeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetMicroAppScopeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetMicroAppScopeResponseBodyResult) SetDeptIds(v []*int64) *GetMicroAppScopeResponseBodyResult {
	s.DeptIds = v
	return s
}

func (s *GetMicroAppScopeResponseBodyResult) SetOnlyAdminVisible(v bool) *GetMicroAppScopeResponseBodyResult {
	s.OnlyAdminVisible = &v
	return s
}

func (s *GetMicroAppScopeResponseBodyResult) SetRoleIds(v []*int64) *GetMicroAppScopeResponseBodyResult {
	s.RoleIds = v
	return s
}

func (s *GetMicroAppScopeResponseBodyResult) SetUserIds(v []*string) *GetMicroAppScopeResponseBodyResult {
	s.UserIds = v
	return s
}

type GetMicroAppScopeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMicroAppScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMicroAppScopeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMicroAppScopeResponse) GoString() string {
	return s.String()
}

func (s *GetMicroAppScopeResponse) SetHeaders(v map[string]*string) *GetMicroAppScopeResponse {
	s.Headers = v
	return s
}

func (s *GetMicroAppScopeResponse) SetStatusCode(v int32) *GetMicroAppScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMicroAppScopeResponse) SetBody(v *GetMicroAppScopeResponseBody) *GetMicroAppScopeResponse {
	s.Body = v
	return s
}

type GetMicroAppUserAccessHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetMicroAppUserAccessHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetMicroAppUserAccessHeaders) GoString() string {
	return s.String()
}

func (s *GetMicroAppUserAccessHeaders) SetCommonHeaders(v map[string]*string) *GetMicroAppUserAccessHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMicroAppUserAccessHeaders) SetXAcsDingtalkAccessToken(v string) *GetMicroAppUserAccessHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetMicroAppUserAccessResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetMicroAppUserAccessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMicroAppUserAccessResponseBody) GoString() string {
	return s.String()
}

func (s *GetMicroAppUserAccessResponseBody) SetResult(v bool) *GetMicroAppUserAccessResponseBody {
	s.Result = &v
	return s
}

type GetMicroAppUserAccessResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMicroAppUserAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMicroAppUserAccessResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMicroAppUserAccessResponse) GoString() string {
	return s.String()
}

func (s *GetMicroAppUserAccessResponse) SetHeaders(v map[string]*string) *GetMicroAppUserAccessResponse {
	s.Headers = v
	return s
}

func (s *GetMicroAppUserAccessResponse) SetStatusCode(v int32) *GetMicroAppUserAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMicroAppUserAccessResponse) SetBody(v *GetMicroAppUserAccessResponseBody) *GetMicroAppUserAccessResponse {
	s.Body = v
	return s
}

type GetUserAppDevAccessHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserAppDevAccessHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserAppDevAccessHeaders) GoString() string {
	return s.String()
}

func (s *GetUserAppDevAccessHeaders) SetCommonHeaders(v map[string]*string) *GetUserAppDevAccessHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserAppDevAccessHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserAppDevAccessHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserAppDevAccessResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetUserAppDevAccessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserAppDevAccessResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserAppDevAccessResponseBody) SetResult(v bool) *GetUserAppDevAccessResponseBody {
	s.Result = &v
	return s
}

type GetUserAppDevAccessResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserAppDevAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserAppDevAccessResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserAppDevAccessResponse) GoString() string {
	return s.String()
}

func (s *GetUserAppDevAccessResponse) SetHeaders(v map[string]*string) *GetUserAppDevAccessResponse {
	s.Headers = v
	return s
}

func (s *GetUserAppDevAccessResponse) SetStatusCode(v int32) *GetUserAppDevAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserAppDevAccessResponse) SetBody(v *GetUserAppDevAccessResponseBody) *GetUserAppDevAccessResponse {
	s.Body = v
	return s
}

type GetVersionInfoRequest struct {
	UnifiedAppId *string `json:"unifiedAppId,omitempty" xml:"unifiedAppId,omitempty"`
	VersionId    *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s GetVersionInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVersionInfoRequest) GoString() string {
	return s.String()
}

func (s *GetVersionInfoRequest) SetUnifiedAppId(v string) *GetVersionInfoRequest {
	s.UnifiedAppId = &v
	return s
}

func (s *GetVersionInfoRequest) SetVersionId(v string) *GetVersionInfoRequest {
	s.VersionId = &v
	return s
}

type GetVersionInfoResponseBody struct {
	Arguments []*string `json:"arguments,omitempty" xml:"arguments,omitempty" type:"Repeated"`
	ErrorCode *string   `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string   `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result    *string   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetVersionInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVersionInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetVersionInfoResponseBody) SetArguments(v []*string) *GetVersionInfoResponseBody {
	s.Arguments = v
	return s
}

func (s *GetVersionInfoResponseBody) SetErrorCode(v string) *GetVersionInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetVersionInfoResponseBody) SetErrorMsg(v string) *GetVersionInfoResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetVersionInfoResponseBody) SetResult(v string) *GetVersionInfoResponseBody {
	s.Result = &v
	return s
}

func (s *GetVersionInfoResponseBody) SetSuccess(v bool) *GetVersionInfoResponseBody {
	s.Success = &v
	return s
}

type GetVersionInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVersionInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVersionInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVersionInfoResponse) GoString() string {
	return s.String()
}

func (s *GetVersionInfoResponse) SetHeaders(v map[string]*string) *GetVersionInfoResponse {
	s.Headers = v
	return s
}

func (s *GetVersionInfoResponse) SetStatusCode(v int32) *GetVersionInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVersionInfoResponse) SetBody(v *GetVersionInfoResponseBody) *GetVersionInfoResponse {
	s.Body = v
	return s
}

type IsOrgMicroAppVisibleByUserIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IsOrgMicroAppVisibleByUserIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s IsOrgMicroAppVisibleByUserIdHeaders) GoString() string {
	return s.String()
}

func (s *IsOrgMicroAppVisibleByUserIdHeaders) SetCommonHeaders(v map[string]*string) *IsOrgMicroAppVisibleByUserIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IsOrgMicroAppVisibleByUserIdHeaders) SetXAcsDingtalkAccessToken(v string) *IsOrgMicroAppVisibleByUserIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IsOrgMicroAppVisibleByUserIdRequest struct {
	// This parameter is required.
	AppId *int64 `json:"appId,omitempty" xml:"appId,omitempty"`
	// This parameter is required.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s IsOrgMicroAppVisibleByUserIdRequest) String() string {
	return tea.Prettify(s)
}

func (s IsOrgMicroAppVisibleByUserIdRequest) GoString() string {
	return s.String()
}

func (s *IsOrgMicroAppVisibleByUserIdRequest) SetAppId(v int64) *IsOrgMicroAppVisibleByUserIdRequest {
	s.AppId = &v
	return s
}

func (s *IsOrgMicroAppVisibleByUserIdRequest) SetId(v string) *IsOrgMicroAppVisibleByUserIdRequest {
	s.Id = &v
	return s
}

func (s *IsOrgMicroAppVisibleByUserIdRequest) SetType(v string) *IsOrgMicroAppVisibleByUserIdRequest {
	s.Type = &v
	return s
}

type IsOrgMicroAppVisibleByUserIdResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s IsOrgMicroAppVisibleByUserIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IsOrgMicroAppVisibleByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *IsOrgMicroAppVisibleByUserIdResponseBody) SetResult(v bool) *IsOrgMicroAppVisibleByUserIdResponseBody {
	s.Result = &v
	return s
}

type IsOrgMicroAppVisibleByUserIdResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IsOrgMicroAppVisibleByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IsOrgMicroAppVisibleByUserIdResponse) String() string {
	return tea.Prettify(s)
}

func (s IsOrgMicroAppVisibleByUserIdResponse) GoString() string {
	return s.String()
}

func (s *IsOrgMicroAppVisibleByUserIdResponse) SetHeaders(v map[string]*string) *IsOrgMicroAppVisibleByUserIdResponse {
	s.Headers = v
	return s
}

func (s *IsOrgMicroAppVisibleByUserIdResponse) SetStatusCode(v int32) *IsOrgMicroAppVisibleByUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *IsOrgMicroAppVisibleByUserIdResponse) SetBody(v *IsOrgMicroAppVisibleByUserIdResponseBody) *IsOrgMicroAppVisibleByUserIdResponse {
	s.Body = v
	return s
}

type ListAllAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListAllAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListAllAppHeaders) GoString() string {
	return s.String()
}

func (s *ListAllAppHeaders) SetCommonHeaders(v map[string]*string) *ListAllAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAllAppHeaders) SetXAcsDingtalkAccessToken(v string) *ListAllAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListAllAppResponseBody struct {
	AppList []*ListAllAppResponseBodyAppList `json:"appList,omitempty" xml:"appList,omitempty" type:"Repeated"`
}

func (s ListAllAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAllAppResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllAppResponseBody) SetAppList(v []*ListAllAppResponseBodyAppList) *ListAllAppResponseBody {
	s.AppList = v
	return s
}

type ListAllAppResponseBodyAppList struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111
	AppId *int64 `json:"appId,omitempty" xml:"appId,omitempty"`
	// This parameter is required.
	AppStatus *int32 `json:"appStatus,omitempty" xml:"appStatus,omitempty"`
	// example:
	//
	// desc
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	DevelopType *int32 `json:"developType,omitempty" xml:"developType,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	HomepageLink *string `json:"homepageLink,omitempty" xml:"homepageLink,omitempty"`
	// example:
	//
	// icon
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	OmpLink *string `json:"ompLink,omitempty" xml:"ompLink,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	PcHomepageLink *string `json:"pcHomepageLink,omitempty" xml:"pcHomepageLink,omitempty"`
}

func (s ListAllAppResponseBodyAppList) String() string {
	return tea.Prettify(s)
}

func (s ListAllAppResponseBodyAppList) GoString() string {
	return s.String()
}

func (s *ListAllAppResponseBodyAppList) SetAgentId(v int64) *ListAllAppResponseBodyAppList {
	s.AgentId = &v
	return s
}

func (s *ListAllAppResponseBodyAppList) SetAppId(v int64) *ListAllAppResponseBodyAppList {
	s.AppId = &v
	return s
}

func (s *ListAllAppResponseBodyAppList) SetAppStatus(v int32) *ListAllAppResponseBodyAppList {
	s.AppStatus = &v
	return s
}

func (s *ListAllAppResponseBodyAppList) SetDesc(v string) *ListAllAppResponseBodyAppList {
	s.Desc = &v
	return s
}

func (s *ListAllAppResponseBodyAppList) SetDevelopType(v int32) *ListAllAppResponseBodyAppList {
	s.DevelopType = &v
	return s
}

func (s *ListAllAppResponseBodyAppList) SetHomepageLink(v string) *ListAllAppResponseBodyAppList {
	s.HomepageLink = &v
	return s
}

func (s *ListAllAppResponseBodyAppList) SetIcon(v string) *ListAllAppResponseBodyAppList {
	s.Icon = &v
	return s
}

func (s *ListAllAppResponseBodyAppList) SetName(v string) *ListAllAppResponseBodyAppList {
	s.Name = &v
	return s
}

func (s *ListAllAppResponseBodyAppList) SetOmpLink(v string) *ListAllAppResponseBodyAppList {
	s.OmpLink = &v
	return s
}

func (s *ListAllAppResponseBodyAppList) SetPcHomepageLink(v string) *ListAllAppResponseBodyAppList {
	s.PcHomepageLink = &v
	return s
}

type ListAllAppResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllAppResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAllAppResponse) GoString() string {
	return s.String()
}

func (s *ListAllAppResponse) SetHeaders(v map[string]*string) *ListAllAppResponse {
	s.Headers = v
	return s
}

func (s *ListAllAppResponse) SetStatusCode(v int32) *ListAllAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllAppResponse) SetBody(v *ListAllAppResponseBody) *ListAllAppResponse {
	s.Body = v
	return s
}

type ListAllInnerAppsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListAllInnerAppsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListAllInnerAppsHeaders) GoString() string {
	return s.String()
}

func (s *ListAllInnerAppsHeaders) SetCommonHeaders(v map[string]*string) *ListAllInnerAppsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAllInnerAppsHeaders) SetXAcsDingtalkAccessToken(v string) *ListAllInnerAppsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListAllInnerAppsResponseBody struct {
	AppList []*ListAllInnerAppsResponseBodyAppList `json:"appList,omitempty" xml:"appList,omitempty" type:"Repeated"`
}

func (s ListAllInnerAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAllInnerAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllInnerAppsResponseBody) SetAppList(v []*ListAllInnerAppsResponseBodyAppList) *ListAllInnerAppsResponseBody {
	s.AppList = v
	return s
}

type ListAllInnerAppsResponseBodyAppList struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111
	AppId *int64 `json:"appId,omitempty" xml:"appId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	AppStatus   *int32                                            `json:"appStatus,omitempty" xml:"appStatus,omitempty"`
	CoolAppInfo []*ListAllInnerAppsResponseBodyAppListCoolAppInfo `json:"coolAppInfo,omitempty" xml:"coolAppInfo,omitempty" type:"Repeated"`
	// example:
	//
	// desc
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	DevelopType *int32 `json:"developType,omitempty" xml:"developType,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	HomepageLink *string `json:"homepageLink,omitempty" xml:"homepageLink,omitempty"`
	// example:
	//
	// icon
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	OmpLink *string `json:"ompLink,omitempty" xml:"ompLink,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	PcHomepageLink *string                                       `json:"pcHomepageLink,omitempty" xml:"pcHomepageLink,omitempty"`
	RobotInfo      *ListAllInnerAppsResponseBodyAppListRobotInfo `json:"robotInfo,omitempty" xml:"robotInfo,omitempty" type:"Struct"`
}

func (s ListAllInnerAppsResponseBodyAppList) String() string {
	return tea.Prettify(s)
}

func (s ListAllInnerAppsResponseBodyAppList) GoString() string {
	return s.String()
}

func (s *ListAllInnerAppsResponseBodyAppList) SetAgentId(v int64) *ListAllInnerAppsResponseBodyAppList {
	s.AgentId = &v
	return s
}

func (s *ListAllInnerAppsResponseBodyAppList) SetAppId(v int64) *ListAllInnerAppsResponseBodyAppList {
	s.AppId = &v
	return s
}

func (s *ListAllInnerAppsResponseBodyAppList) SetAppStatus(v int32) *ListAllInnerAppsResponseBodyAppList {
	s.AppStatus = &v
	return s
}

func (s *ListAllInnerAppsResponseBodyAppList) SetCoolAppInfo(v []*ListAllInnerAppsResponseBodyAppListCoolAppInfo) *ListAllInnerAppsResponseBodyAppList {
	s.CoolAppInfo = v
	return s
}

func (s *ListAllInnerAppsResponseBodyAppList) SetDesc(v string) *ListAllInnerAppsResponseBodyAppList {
	s.Desc = &v
	return s
}

func (s *ListAllInnerAppsResponseBodyAppList) SetDevelopType(v int32) *ListAllInnerAppsResponseBodyAppList {
	s.DevelopType = &v
	return s
}

func (s *ListAllInnerAppsResponseBodyAppList) SetHomepageLink(v string) *ListAllInnerAppsResponseBodyAppList {
	s.HomepageLink = &v
	return s
}

func (s *ListAllInnerAppsResponseBodyAppList) SetIcon(v string) *ListAllInnerAppsResponseBodyAppList {
	s.Icon = &v
	return s
}

func (s *ListAllInnerAppsResponseBodyAppList) SetName(v string) *ListAllInnerAppsResponseBodyAppList {
	s.Name = &v
	return s
}

func (s *ListAllInnerAppsResponseBodyAppList) SetOmpLink(v string) *ListAllInnerAppsResponseBodyAppList {
	s.OmpLink = &v
	return s
}

func (s *ListAllInnerAppsResponseBodyAppList) SetPcHomepageLink(v string) *ListAllInnerAppsResponseBodyAppList {
	s.PcHomepageLink = &v
	return s
}

func (s *ListAllInnerAppsResponseBodyAppList) SetRobotInfo(v *ListAllInnerAppsResponseBodyAppListRobotInfo) *ListAllInnerAppsResponseBodyAppList {
	s.RobotInfo = v
	return s
}

type ListAllInnerAppsResponseBodyAppListCoolAppInfo struct {
	CoolAppCode *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListAllInnerAppsResponseBodyAppListCoolAppInfo) String() string {
	return tea.Prettify(s)
}

func (s ListAllInnerAppsResponseBodyAppListCoolAppInfo) GoString() string {
	return s.String()
}

func (s *ListAllInnerAppsResponseBodyAppListCoolAppInfo) SetCoolAppCode(v string) *ListAllInnerAppsResponseBodyAppListCoolAppInfo {
	s.CoolAppCode = &v
	return s
}

func (s *ListAllInnerAppsResponseBodyAppListCoolAppInfo) SetName(v string) *ListAllInnerAppsResponseBodyAppListCoolAppInfo {
	s.Name = &v
	return s
}

type ListAllInnerAppsResponseBodyAppListRobotInfo struct {
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s ListAllInnerAppsResponseBodyAppListRobotInfo) String() string {
	return tea.Prettify(s)
}

func (s ListAllInnerAppsResponseBodyAppListRobotInfo) GoString() string {
	return s.String()
}

func (s *ListAllInnerAppsResponseBodyAppListRobotInfo) SetName(v string) *ListAllInnerAppsResponseBodyAppListRobotInfo {
	s.Name = &v
	return s
}

func (s *ListAllInnerAppsResponseBodyAppListRobotInfo) SetRobotCode(v string) *ListAllInnerAppsResponseBodyAppListRobotInfo {
	s.RobotCode = &v
	return s
}

type ListAllInnerAppsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllInnerAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllInnerAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAllInnerAppsResponse) GoString() string {
	return s.String()
}

func (s *ListAllInnerAppsResponse) SetHeaders(v map[string]*string) *ListAllInnerAppsResponse {
	s.Headers = v
	return s
}

func (s *ListAllInnerAppsResponse) SetStatusCode(v int32) *ListAllInnerAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllInnerAppsResponse) SetBody(v *ListAllInnerAppsResponseBody) *ListAllInnerAppsResponse {
	s.Body = v
	return s
}

type ListAppByClientIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListAppByClientIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListAppByClientIdHeaders) GoString() string {
	return s.String()
}

func (s *ListAppByClientIdHeaders) SetCommonHeaders(v map[string]*string) *ListAppByClientIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAppByClientIdHeaders) SetXAcsDingtalkAccessToken(v string) *ListAppByClientIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListAppByClientIdResponseBody struct {
	AgentId        *int64  `json:"agentId,omitempty" xml:"agentId,omitempty"`
	AppId          *int64  `json:"appId,omitempty" xml:"appId,omitempty"`
	AppStatus      *int32  `json:"appStatus,omitempty" xml:"appStatus,omitempty"`
	Desc           *string `json:"desc,omitempty" xml:"desc,omitempty"`
	DevelopType    *int32  `json:"developType,omitempty" xml:"developType,omitempty"`
	HomepageLink   *string `json:"homepageLink,omitempty" xml:"homepageLink,omitempty"`
	Icon           *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	OmpLink        *string `json:"ompLink,omitempty" xml:"ompLink,omitempty"`
	PcHomepageLink *string `json:"pcHomepageLink,omitempty" xml:"pcHomepageLink,omitempty"`
}

func (s ListAppByClientIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppByClientIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppByClientIdResponseBody) SetAgentId(v int64) *ListAppByClientIdResponseBody {
	s.AgentId = &v
	return s
}

func (s *ListAppByClientIdResponseBody) SetAppId(v int64) *ListAppByClientIdResponseBody {
	s.AppId = &v
	return s
}

func (s *ListAppByClientIdResponseBody) SetAppStatus(v int32) *ListAppByClientIdResponseBody {
	s.AppStatus = &v
	return s
}

func (s *ListAppByClientIdResponseBody) SetDesc(v string) *ListAppByClientIdResponseBody {
	s.Desc = &v
	return s
}

func (s *ListAppByClientIdResponseBody) SetDevelopType(v int32) *ListAppByClientIdResponseBody {
	s.DevelopType = &v
	return s
}

func (s *ListAppByClientIdResponseBody) SetHomepageLink(v string) *ListAppByClientIdResponseBody {
	s.HomepageLink = &v
	return s
}

func (s *ListAppByClientIdResponseBody) SetIcon(v string) *ListAppByClientIdResponseBody {
	s.Icon = &v
	return s
}

func (s *ListAppByClientIdResponseBody) SetName(v string) *ListAppByClientIdResponseBody {
	s.Name = &v
	return s
}

func (s *ListAppByClientIdResponseBody) SetOmpLink(v string) *ListAppByClientIdResponseBody {
	s.OmpLink = &v
	return s
}

func (s *ListAppByClientIdResponseBody) SetPcHomepageLink(v string) *ListAppByClientIdResponseBody {
	s.PcHomepageLink = &v
	return s
}

type ListAppByClientIdResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppByClientIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppByClientIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppByClientIdResponse) GoString() string {
	return s.String()
}

func (s *ListAppByClientIdResponse) SetHeaders(v map[string]*string) *ListAppByClientIdResponse {
	s.Headers = v
	return s
}

func (s *ListAppByClientIdResponse) SetStatusCode(v int32) *ListAppByClientIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppByClientIdResponse) SetBody(v *ListAppByClientIdResponseBody) *ListAppByClientIdResponse {
	s.Body = v
	return s
}

type ListAppRoleScopesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListAppRoleScopesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListAppRoleScopesHeaders) GoString() string {
	return s.String()
}

func (s *ListAppRoleScopesHeaders) SetCommonHeaders(v map[string]*string) *ListAppRoleScopesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAppRoleScopesHeaders) SetXAcsDingtalkAccessToken(v string) *ListAppRoleScopesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListAppRoleScopesRequest struct {
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Size      *int64 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListAppRoleScopesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppRoleScopesRequest) GoString() string {
	return s.String()
}

func (s *ListAppRoleScopesRequest) SetNextToken(v int64) *ListAppRoleScopesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAppRoleScopesRequest) SetSize(v int64) *ListAppRoleScopesRequest {
	s.Size = &v
	return s
}

type ListAppRoleScopesResponseBody struct {
	DataList []*ListAppRoleScopesResponseBodyDataList `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 0
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListAppRoleScopesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppRoleScopesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppRoleScopesResponseBody) SetDataList(v []*ListAppRoleScopesResponseBodyDataList) *ListAppRoleScopesResponseBody {
	s.DataList = v
	return s
}

func (s *ListAppRoleScopesResponseBody) SetHasMore(v bool) *ListAppRoleScopesResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListAppRoleScopesResponseBody) SetNextToken(v int64) *ListAppRoleScopesResponseBody {
	s.NextToken = &v
	return s
}

type ListAppRoleScopesResponseBodyDataList struct {
	// example:
	//
	// false
	CanManageRole *bool    `json:"canManageRole,omitempty" xml:"canManageRole,omitempty"`
	DeptIdList    []*int64 `json:"deptIdList,omitempty" xml:"deptIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 123
	RoleId *int64 `json:"roleId,omitempty" xml:"roleId,omitempty"`
	// example:
	//
	// 财务
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	// example:
	//
	// PART_VISIBLE
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	// example:
	//
	// 123
	ScopeVersion *int64    `json:"scopeVersion,omitempty" xml:"scopeVersion,omitempty"`
	UserIdList   []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s ListAppRoleScopesResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListAppRoleScopesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListAppRoleScopesResponseBodyDataList) SetCanManageRole(v bool) *ListAppRoleScopesResponseBodyDataList {
	s.CanManageRole = &v
	return s
}

func (s *ListAppRoleScopesResponseBodyDataList) SetDeptIdList(v []*int64) *ListAppRoleScopesResponseBodyDataList {
	s.DeptIdList = v
	return s
}

func (s *ListAppRoleScopesResponseBodyDataList) SetRoleId(v int64) *ListAppRoleScopesResponseBodyDataList {
	s.RoleId = &v
	return s
}

func (s *ListAppRoleScopesResponseBodyDataList) SetRoleName(v string) *ListAppRoleScopesResponseBodyDataList {
	s.RoleName = &v
	return s
}

func (s *ListAppRoleScopesResponseBodyDataList) SetScopeType(v string) *ListAppRoleScopesResponseBodyDataList {
	s.ScopeType = &v
	return s
}

func (s *ListAppRoleScopesResponseBodyDataList) SetScopeVersion(v int64) *ListAppRoleScopesResponseBodyDataList {
	s.ScopeVersion = &v
	return s
}

func (s *ListAppRoleScopesResponseBodyDataList) SetUserIdList(v []*string) *ListAppRoleScopesResponseBodyDataList {
	s.UserIdList = v
	return s
}

type ListAppRoleScopesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppRoleScopesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppRoleScopesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppRoleScopesResponse) GoString() string {
	return s.String()
}

func (s *ListAppRoleScopesResponse) SetHeaders(v map[string]*string) *ListAppRoleScopesResponse {
	s.Headers = v
	return s
}

func (s *ListAppRoleScopesResponse) SetStatusCode(v int32) *ListAppRoleScopesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppRoleScopesResponse) SetBody(v *ListAppRoleScopesResponseBody) *ListAppRoleScopesResponse {
	s.Body = v
	return s
}

type ListInnerAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListInnerAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListInnerAppHeaders) GoString() string {
	return s.String()
}

func (s *ListInnerAppHeaders) SetCommonHeaders(v map[string]*string) *ListInnerAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListInnerAppHeaders) SetXAcsDingtalkAccessToken(v string) *ListInnerAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListInnerAppRequest struct {
	// example:
	//
	// xxxx
	EcologicalCorpId *string `json:"ecologicalCorpId,omitempty" xml:"ecologicalCorpId,omitempty"`
}

func (s ListInnerAppRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInnerAppRequest) GoString() string {
	return s.String()
}

func (s *ListInnerAppRequest) SetEcologicalCorpId(v string) *ListInnerAppRequest {
	s.EcologicalCorpId = &v
	return s
}

type ListInnerAppResponseBody struct {
	AppList []*ListInnerAppResponseBodyAppList `json:"appList,omitempty" xml:"appList,omitempty" type:"Repeated"`
}

func (s ListInnerAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInnerAppResponseBody) GoString() string {
	return s.String()
}

func (s *ListInnerAppResponseBody) SetAppList(v []*ListInnerAppResponseBodyAppList) *ListInnerAppResponseBody {
	s.AppList = v
	return s
}

type ListInnerAppResponseBodyAppList struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// desc
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	HomepageLink *string `json:"homepageLink,omitempty" xml:"homepageLink,omitempty"`
	// example:
	//
	// icon
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	OmpLink *string `json:"ompLink,omitempty" xml:"ompLink,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	PcHomepageLink *string `json:"pcHomepageLink,omitempty" xml:"pcHomepageLink,omitempty"`
}

func (s ListInnerAppResponseBodyAppList) String() string {
	return tea.Prettify(s)
}

func (s ListInnerAppResponseBodyAppList) GoString() string {
	return s.String()
}

func (s *ListInnerAppResponseBodyAppList) SetAgentId(v int64) *ListInnerAppResponseBodyAppList {
	s.AgentId = &v
	return s
}

func (s *ListInnerAppResponseBodyAppList) SetDesc(v string) *ListInnerAppResponseBodyAppList {
	s.Desc = &v
	return s
}

func (s *ListInnerAppResponseBodyAppList) SetHomepageLink(v string) *ListInnerAppResponseBodyAppList {
	s.HomepageLink = &v
	return s
}

func (s *ListInnerAppResponseBodyAppList) SetIcon(v string) *ListInnerAppResponseBodyAppList {
	s.Icon = &v
	return s
}

func (s *ListInnerAppResponseBodyAppList) SetName(v string) *ListInnerAppResponseBodyAppList {
	s.Name = &v
	return s
}

func (s *ListInnerAppResponseBodyAppList) SetOmpLink(v string) *ListInnerAppResponseBodyAppList {
	s.OmpLink = &v
	return s
}

func (s *ListInnerAppResponseBodyAppList) SetPcHomepageLink(v string) *ListInnerAppResponseBodyAppList {
	s.PcHomepageLink = &v
	return s
}

type ListInnerAppResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInnerAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInnerAppResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInnerAppResponse) GoString() string {
	return s.String()
}

func (s *ListInnerAppResponse) SetHeaders(v map[string]*string) *ListInnerAppResponse {
	s.Headers = v
	return s
}

func (s *ListInnerAppResponse) SetStatusCode(v int32) *ListInnerAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInnerAppResponse) SetBody(v *ListInnerAppResponseBody) *ListInnerAppResponse {
	s.Body = v
	return s
}

type ListInnerAppVersionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListInnerAppVersionHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListInnerAppVersionHeaders) GoString() string {
	return s.String()
}

func (s *ListInnerAppVersionHeaders) SetCommonHeaders(v map[string]*string) *ListInnerAppVersionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListInnerAppVersionHeaders) SetXAcsDingtalkAccessToken(v string) *ListInnerAppVersionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListInnerAppVersionResponseBody struct {
	AppVersionList []*ListInnerAppVersionResponseBodyAppVersionList `json:"appVersionList,omitempty" xml:"appVersionList,omitempty" type:"Repeated"`
}

func (s ListInnerAppVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInnerAppVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ListInnerAppVersionResponseBody) SetAppVersionList(v []*ListInnerAppVersionResponseBodyAppVersionList) *ListInnerAppVersionResponseBody {
	s.AppVersionList = v
	return s
}

type ListInnerAppVersionResponseBodyAppVersionList struct {
	// This parameter is required.
	//
	// example:
	//
	// 0.0.1
	AppVersion *string `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	AppVersionId *int64 `json:"appVersionId,omitempty" xml:"appVersionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	AppVersionType *int32 `json:"appVersionType,omitempty" xml:"appVersionType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-01 00:00:00
	CreateTime   *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	EntranceLink *string `json:"entranceLink,omitempty" xml:"entranceLink,omitempty"`
	// This parameter is required.
	MiniAppId   *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	MiniAppOnPc *bool   `json:"miniAppOnPc,omitempty" xml:"miniAppOnPc,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-01 00:00:00
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
}

func (s ListInnerAppVersionResponseBodyAppVersionList) String() string {
	return tea.Prettify(s)
}

func (s ListInnerAppVersionResponseBodyAppVersionList) GoString() string {
	return s.String()
}

func (s *ListInnerAppVersionResponseBodyAppVersionList) SetAppVersion(v string) *ListInnerAppVersionResponseBodyAppVersionList {
	s.AppVersion = &v
	return s
}

func (s *ListInnerAppVersionResponseBodyAppVersionList) SetAppVersionId(v int64) *ListInnerAppVersionResponseBodyAppVersionList {
	s.AppVersionId = &v
	return s
}

func (s *ListInnerAppVersionResponseBodyAppVersionList) SetAppVersionType(v int32) *ListInnerAppVersionResponseBodyAppVersionList {
	s.AppVersionType = &v
	return s
}

func (s *ListInnerAppVersionResponseBodyAppVersionList) SetCreateTime(v string) *ListInnerAppVersionResponseBodyAppVersionList {
	s.CreateTime = &v
	return s
}

func (s *ListInnerAppVersionResponseBodyAppVersionList) SetEntranceLink(v string) *ListInnerAppVersionResponseBodyAppVersionList {
	s.EntranceLink = &v
	return s
}

func (s *ListInnerAppVersionResponseBodyAppVersionList) SetMiniAppId(v string) *ListInnerAppVersionResponseBodyAppVersionList {
	s.MiniAppId = &v
	return s
}

func (s *ListInnerAppVersionResponseBodyAppVersionList) SetMiniAppOnPc(v bool) *ListInnerAppVersionResponseBodyAppVersionList {
	s.MiniAppOnPc = &v
	return s
}

func (s *ListInnerAppVersionResponseBodyAppVersionList) SetModifyTime(v string) *ListInnerAppVersionResponseBodyAppVersionList {
	s.ModifyTime = &v
	return s
}

type ListInnerAppVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInnerAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInnerAppVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInnerAppVersionResponse) GoString() string {
	return s.String()
}

func (s *ListInnerAppVersionResponse) SetHeaders(v map[string]*string) *ListInnerAppVersionResponse {
	s.Headers = v
	return s
}

func (s *ListInnerAppVersionResponse) SetStatusCode(v int32) *ListInnerAppVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInnerAppVersionResponse) SetBody(v *ListInnerAppVersionResponseBody) *ListInnerAppVersionResponse {
	s.Body = v
	return s
}

type ListRoleInfoByUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListRoleInfoByUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListRoleInfoByUserHeaders) GoString() string {
	return s.String()
}

func (s *ListRoleInfoByUserHeaders) SetCommonHeaders(v map[string]*string) *ListRoleInfoByUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListRoleInfoByUserHeaders) SetXAcsDingtalkAccessToken(v string) *ListRoleInfoByUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListRoleInfoByUserResponseBody struct {
	Result []*ListRoleInfoByUserResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListRoleInfoByUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRoleInfoByUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListRoleInfoByUserResponseBody) SetResult(v []*ListRoleInfoByUserResponseBodyResult) *ListRoleInfoByUserResponseBody {
	s.Result = v
	return s
}

type ListRoleInfoByUserResponseBodyResult struct {
	// example:
	//
	// false
	CanManageRole *bool `json:"canManageRole,omitempty" xml:"canManageRole,omitempty"`
	// example:
	//
	// 123
	RoleId *int64 `json:"roleId,omitempty" xml:"roleId,omitempty"`
	// example:
	//
	// 财务
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
}

func (s ListRoleInfoByUserResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListRoleInfoByUserResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListRoleInfoByUserResponseBodyResult) SetCanManageRole(v bool) *ListRoleInfoByUserResponseBodyResult {
	s.CanManageRole = &v
	return s
}

func (s *ListRoleInfoByUserResponseBodyResult) SetRoleId(v int64) *ListRoleInfoByUserResponseBodyResult {
	s.RoleId = &v
	return s
}

func (s *ListRoleInfoByUserResponseBodyResult) SetRoleName(v string) *ListRoleInfoByUserResponseBodyResult {
	s.RoleName = &v
	return s
}

type ListRoleInfoByUserResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRoleInfoByUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRoleInfoByUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRoleInfoByUserResponse) GoString() string {
	return s.String()
}

func (s *ListRoleInfoByUserResponse) SetHeaders(v map[string]*string) *ListRoleInfoByUserResponse {
	s.Headers = v
	return s
}

func (s *ListRoleInfoByUserResponse) SetStatusCode(v int32) *ListRoleInfoByUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRoleInfoByUserResponse) SetBody(v *ListRoleInfoByUserResponseBody) *ListRoleInfoByUserResponse {
	s.Body = v
	return s
}

type ListUserVilebleAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListUserVilebleAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListUserVilebleAppHeaders) GoString() string {
	return s.String()
}

func (s *ListUserVilebleAppHeaders) SetCommonHeaders(v map[string]*string) *ListUserVilebleAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListUserVilebleAppHeaders) SetXAcsDingtalkAccessToken(v string) *ListUserVilebleAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListUserVilebleAppResponseBody struct {
	AppList []*ListUserVilebleAppResponseBodyAppList `json:"appList,omitempty" xml:"appList,omitempty" type:"Repeated"`
}

func (s ListUserVilebleAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserVilebleAppResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserVilebleAppResponseBody) SetAppList(v []*ListUserVilebleAppResponseBodyAppList) *ListUserVilebleAppResponseBody {
	s.AppList = v
	return s
}

type ListUserVilebleAppResponseBodyAppList struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111
	AppId *int64 `json:"appId,omitempty" xml:"appId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	AppStatus *int32 `json:"appStatus,omitempty" xml:"appStatus,omitempty"`
	// example:
	//
	// desc
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	DevelopType *int32 `json:"developType,omitempty" xml:"developType,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	HomepageLink *string                                      `json:"homepageLink,omitempty" xml:"homepageLink,omitempty"`
	I18n         []*ListUserVilebleAppResponseBodyAppListI18n `json:"i18n,omitempty" xml:"i18n,omitempty" type:"Repeated"`
	// example:
	//
	// icon
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	OmpLink *string `json:"ompLink,omitempty" xml:"ompLink,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	PcHomepageLink *string `json:"pcHomepageLink,omitempty" xml:"pcHomepageLink,omitempty"`
	// example:
	//
	// unifiedAppId
	UnifiedAppId *string `json:"unifiedAppId,omitempty" xml:"unifiedAppId,omitempty"`
}

func (s ListUserVilebleAppResponseBodyAppList) String() string {
	return tea.Prettify(s)
}

func (s ListUserVilebleAppResponseBodyAppList) GoString() string {
	return s.String()
}

func (s *ListUserVilebleAppResponseBodyAppList) SetAgentId(v int64) *ListUserVilebleAppResponseBodyAppList {
	s.AgentId = &v
	return s
}

func (s *ListUserVilebleAppResponseBodyAppList) SetAppId(v int64) *ListUserVilebleAppResponseBodyAppList {
	s.AppId = &v
	return s
}

func (s *ListUserVilebleAppResponseBodyAppList) SetAppStatus(v int32) *ListUserVilebleAppResponseBodyAppList {
	s.AppStatus = &v
	return s
}

func (s *ListUserVilebleAppResponseBodyAppList) SetDesc(v string) *ListUserVilebleAppResponseBodyAppList {
	s.Desc = &v
	return s
}

func (s *ListUserVilebleAppResponseBodyAppList) SetDevelopType(v int32) *ListUserVilebleAppResponseBodyAppList {
	s.DevelopType = &v
	return s
}

func (s *ListUserVilebleAppResponseBodyAppList) SetHomepageLink(v string) *ListUserVilebleAppResponseBodyAppList {
	s.HomepageLink = &v
	return s
}

func (s *ListUserVilebleAppResponseBodyAppList) SetI18n(v []*ListUserVilebleAppResponseBodyAppListI18n) *ListUserVilebleAppResponseBodyAppList {
	s.I18n = v
	return s
}

func (s *ListUserVilebleAppResponseBodyAppList) SetIcon(v string) *ListUserVilebleAppResponseBodyAppList {
	s.Icon = &v
	return s
}

func (s *ListUserVilebleAppResponseBodyAppList) SetName(v string) *ListUserVilebleAppResponseBodyAppList {
	s.Name = &v
	return s
}

func (s *ListUserVilebleAppResponseBodyAppList) SetOmpLink(v string) *ListUserVilebleAppResponseBodyAppList {
	s.OmpLink = &v
	return s
}

func (s *ListUserVilebleAppResponseBodyAppList) SetPcHomepageLink(v string) *ListUserVilebleAppResponseBodyAppList {
	s.PcHomepageLink = &v
	return s
}

func (s *ListUserVilebleAppResponseBodyAppList) SetUnifiedAppId(v string) *ListUserVilebleAppResponseBodyAppList {
	s.UnifiedAppId = &v
	return s
}

type ListUserVilebleAppResponseBodyAppListI18n struct {
	Desc    *string `json:"desc,omitempty" xml:"desc,omitempty"`
	I18nKey *string `json:"i18n_key,omitempty" xml:"i18n_key,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListUserVilebleAppResponseBodyAppListI18n) String() string {
	return tea.Prettify(s)
}

func (s ListUserVilebleAppResponseBodyAppListI18n) GoString() string {
	return s.String()
}

func (s *ListUserVilebleAppResponseBodyAppListI18n) SetDesc(v string) *ListUserVilebleAppResponseBodyAppListI18n {
	s.Desc = &v
	return s
}

func (s *ListUserVilebleAppResponseBodyAppListI18n) SetI18nKey(v string) *ListUserVilebleAppResponseBodyAppListI18n {
	s.I18nKey = &v
	return s
}

func (s *ListUserVilebleAppResponseBodyAppListI18n) SetName(v string) *ListUserVilebleAppResponseBodyAppListI18n {
	s.Name = &v
	return s
}

type ListUserVilebleAppResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserVilebleAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserVilebleAppResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserVilebleAppResponse) GoString() string {
	return s.String()
}

func (s *ListUserVilebleAppResponse) SetHeaders(v map[string]*string) *ListUserVilebleAppResponse {
	s.Headers = v
	return s
}

func (s *ListUserVilebleAppResponse) SetStatusCode(v int32) *ListUserVilebleAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserVilebleAppResponse) SetBody(v *ListUserVilebleAppResponseBody) *ListUserVilebleAppResponse {
	s.Body = v
	return s
}

type PageInnerAppHistoryVersionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PageInnerAppHistoryVersionHeaders) String() string {
	return tea.Prettify(s)
}

func (s PageInnerAppHistoryVersionHeaders) GoString() string {
	return s.String()
}

func (s *PageInnerAppHistoryVersionHeaders) SetCommonHeaders(v map[string]*string) *PageInnerAppHistoryVersionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PageInnerAppHistoryVersionHeaders) SetXAcsDingtalkAccessToken(v string) *PageInnerAppHistoryVersionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PageInnerAppHistoryVersionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s PageInnerAppHistoryVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s PageInnerAppHistoryVersionRequest) GoString() string {
	return s.String()
}

func (s *PageInnerAppHistoryVersionRequest) SetPageNumber(v int32) *PageInnerAppHistoryVersionRequest {
	s.PageNumber = &v
	return s
}

func (s *PageInnerAppHistoryVersionRequest) SetPageSize(v int32) *PageInnerAppHistoryVersionRequest {
	s.PageSize = &v
	return s
}

type PageInnerAppHistoryVersionResponseBody struct {
	MiniAppVersionList []*PageInnerAppHistoryVersionResponseBodyMiniAppVersionList `json:"miniAppVersionList,omitempty" xml:"miniAppVersionList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s PageInnerAppHistoryVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageInnerAppHistoryVersionResponseBody) GoString() string {
	return s.String()
}

func (s *PageInnerAppHistoryVersionResponseBody) SetMiniAppVersionList(v []*PageInnerAppHistoryVersionResponseBodyMiniAppVersionList) *PageInnerAppHistoryVersionResponseBody {
	s.MiniAppVersionList = v
	return s
}

func (s *PageInnerAppHistoryVersionResponseBody) SetTotalCount(v int64) *PageInnerAppHistoryVersionResponseBody {
	s.TotalCount = &v
	return s
}

type PageInnerAppHistoryVersionResponseBodyMiniAppVersionList struct {
	// This parameter is required.
	//
	// example:
	//
	// 0.0.1
	AppVersion *string `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	AppVersionId *int64 `json:"appVersionId,omitempty" xml:"appVersionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	AppVersionType *int32 `json:"appVersionType,omitempty" xml:"appVersionType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-01 00:00:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 1
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	// This parameter is required.
	MiniAppOnPc *bool `json:"miniAppOnPc,omitempty" xml:"miniAppOnPc,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-01 00:00:00
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
}

func (s PageInnerAppHistoryVersionResponseBodyMiniAppVersionList) String() string {
	return tea.Prettify(s)
}

func (s PageInnerAppHistoryVersionResponseBodyMiniAppVersionList) GoString() string {
	return s.String()
}

func (s *PageInnerAppHistoryVersionResponseBodyMiniAppVersionList) SetAppVersion(v string) *PageInnerAppHistoryVersionResponseBodyMiniAppVersionList {
	s.AppVersion = &v
	return s
}

func (s *PageInnerAppHistoryVersionResponseBodyMiniAppVersionList) SetAppVersionId(v int64) *PageInnerAppHistoryVersionResponseBodyMiniAppVersionList {
	s.AppVersionId = &v
	return s
}

func (s *PageInnerAppHistoryVersionResponseBodyMiniAppVersionList) SetAppVersionType(v int32) *PageInnerAppHistoryVersionResponseBodyMiniAppVersionList {
	s.AppVersionType = &v
	return s
}

func (s *PageInnerAppHistoryVersionResponseBodyMiniAppVersionList) SetCreateTime(v string) *PageInnerAppHistoryVersionResponseBodyMiniAppVersionList {
	s.CreateTime = &v
	return s
}

func (s *PageInnerAppHistoryVersionResponseBodyMiniAppVersionList) SetMiniAppId(v string) *PageInnerAppHistoryVersionResponseBodyMiniAppVersionList {
	s.MiniAppId = &v
	return s
}

func (s *PageInnerAppHistoryVersionResponseBodyMiniAppVersionList) SetMiniAppOnPc(v bool) *PageInnerAppHistoryVersionResponseBodyMiniAppVersionList {
	s.MiniAppOnPc = &v
	return s
}

func (s *PageInnerAppHistoryVersionResponseBodyMiniAppVersionList) SetModifyTime(v string) *PageInnerAppHistoryVersionResponseBodyMiniAppVersionList {
	s.ModifyTime = &v
	return s
}

type PageInnerAppHistoryVersionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PageInnerAppHistoryVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PageInnerAppHistoryVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s PageInnerAppHistoryVersionResponse) GoString() string {
	return s.String()
}

func (s *PageInnerAppHistoryVersionResponse) SetHeaders(v map[string]*string) *PageInnerAppHistoryVersionResponse {
	s.Headers = v
	return s
}

func (s *PageInnerAppHistoryVersionResponse) SetStatusCode(v int32) *PageInnerAppHistoryVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *PageInnerAppHistoryVersionResponse) SetBody(v *PageInnerAppHistoryVersionResponseBody) *PageInnerAppHistoryVersionResponse {
	s.Body = v
	return s
}

type PublishInnerAppVersionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PublishInnerAppVersionHeaders) String() string {
	return tea.Prettify(s)
}

func (s PublishInnerAppVersionHeaders) GoString() string {
	return s.String()
}

func (s *PublishInnerAppVersionHeaders) SetCommonHeaders(v map[string]*string) *PublishInnerAppVersionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PublishInnerAppVersionHeaders) SetXAcsDingtalkAccessToken(v string) *PublishInnerAppVersionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PublishInnerAppVersionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	AppVersionId *int64 `json:"appVersionId,omitempty" xml:"appVersionId,omitempty"`
	MiniAppOnPc  *bool  `json:"miniAppOnPc,omitempty" xml:"miniAppOnPc,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	OpUnionId *string `json:"opUnionId,omitempty" xml:"opUnionId,omitempty"`
	// example:
	//
	// online
	PublishType *string `json:"publishType,omitempty" xml:"publishType,omitempty"`
}

func (s PublishInnerAppVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishInnerAppVersionRequest) GoString() string {
	return s.String()
}

func (s *PublishInnerAppVersionRequest) SetAppVersionId(v int64) *PublishInnerAppVersionRequest {
	s.AppVersionId = &v
	return s
}

func (s *PublishInnerAppVersionRequest) SetMiniAppOnPc(v bool) *PublishInnerAppVersionRequest {
	s.MiniAppOnPc = &v
	return s
}

func (s *PublishInnerAppVersionRequest) SetOpUnionId(v string) *PublishInnerAppVersionRequest {
	s.OpUnionId = &v
	return s
}

func (s *PublishInnerAppVersionRequest) SetPublishType(v string) *PublishInnerAppVersionRequest {
	s.PublishType = &v
	return s
}

type PublishInnerAppVersionResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s PublishInnerAppVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishInnerAppVersionResponseBody) GoString() string {
	return s.String()
}

func (s *PublishInnerAppVersionResponseBody) SetResult(v bool) *PublishInnerAppVersionResponseBody {
	s.Result = &v
	return s
}

type PublishInnerAppVersionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishInnerAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishInnerAppVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishInnerAppVersionResponse) GoString() string {
	return s.String()
}

func (s *PublishInnerAppVersionResponse) SetHeaders(v map[string]*string) *PublishInnerAppVersionResponse {
	s.Headers = v
	return s
}

func (s *PublishInnerAppVersionResponse) SetStatusCode(v int32) *PublishInnerAppVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishInnerAppVersionResponse) SetBody(v *PublishInnerAppVersionResponseBody) *PublishInnerAppVersionResponse {
	s.Body = v
	return s
}

type RebuildRoleScopeForAppRoleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RebuildRoleScopeForAppRoleHeaders) String() string {
	return tea.Prettify(s)
}

func (s RebuildRoleScopeForAppRoleHeaders) GoString() string {
	return s.String()
}

func (s *RebuildRoleScopeForAppRoleHeaders) SetCommonHeaders(v map[string]*string) *RebuildRoleScopeForAppRoleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RebuildRoleScopeForAppRoleHeaders) SetXAcsDingtalkAccessToken(v string) *RebuildRoleScopeForAppRoleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RebuildRoleScopeForAppRoleRequest struct {
	DeptIdList []*int64 `json:"deptIdList,omitempty" xml:"deptIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// This parameter is required.
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	// This parameter is required.
	ScopeVersion *int64    `json:"scopeVersion,omitempty" xml:"scopeVersion,omitempty"`
	UserIdList   []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s RebuildRoleScopeForAppRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s RebuildRoleScopeForAppRoleRequest) GoString() string {
	return s.String()
}

func (s *RebuildRoleScopeForAppRoleRequest) SetDeptIdList(v []*int64) *RebuildRoleScopeForAppRoleRequest {
	s.DeptIdList = v
	return s
}

func (s *RebuildRoleScopeForAppRoleRequest) SetOpUserId(v string) *RebuildRoleScopeForAppRoleRequest {
	s.OpUserId = &v
	return s
}

func (s *RebuildRoleScopeForAppRoleRequest) SetScopeType(v string) *RebuildRoleScopeForAppRoleRequest {
	s.ScopeType = &v
	return s
}

func (s *RebuildRoleScopeForAppRoleRequest) SetScopeVersion(v int64) *RebuildRoleScopeForAppRoleRequest {
	s.ScopeVersion = &v
	return s
}

func (s *RebuildRoleScopeForAppRoleRequest) SetUserIdList(v []*string) *RebuildRoleScopeForAppRoleRequest {
	s.UserIdList = v
	return s
}

type RebuildRoleScopeForAppRoleResponseBody struct {
	// example:
	//
	// 123
	LatestScopeVersion *int64 `json:"latestScopeVersion,omitempty" xml:"latestScopeVersion,omitempty"`
}

func (s RebuildRoleScopeForAppRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebuildRoleScopeForAppRoleResponseBody) GoString() string {
	return s.String()
}

func (s *RebuildRoleScopeForAppRoleResponseBody) SetLatestScopeVersion(v int64) *RebuildRoleScopeForAppRoleResponseBody {
	s.LatestScopeVersion = &v
	return s
}

type RebuildRoleScopeForAppRoleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebuildRoleScopeForAppRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebuildRoleScopeForAppRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s RebuildRoleScopeForAppRoleResponse) GoString() string {
	return s.String()
}

func (s *RebuildRoleScopeForAppRoleResponse) SetHeaders(v map[string]*string) *RebuildRoleScopeForAppRoleResponse {
	s.Headers = v
	return s
}

func (s *RebuildRoleScopeForAppRoleResponse) SetStatusCode(v int32) *RebuildRoleScopeForAppRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *RebuildRoleScopeForAppRoleResponse) SetBody(v *RebuildRoleScopeForAppRoleResponseBody) *RebuildRoleScopeForAppRoleResponse {
	s.Body = v
	return s
}

type RegisterCustomAppRoleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RegisterCustomAppRoleHeaders) String() string {
	return tea.Prettify(s)
}

func (s RegisterCustomAppRoleHeaders) GoString() string {
	return s.String()
}

func (s *RegisterCustomAppRoleHeaders) SetCommonHeaders(v map[string]*string) *RegisterCustomAppRoleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RegisterCustomAppRoleHeaders) SetXAcsDingtalkAccessToken(v string) *RegisterCustomAppRoleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RegisterCustomAppRoleRequest struct {
	// example:
	//
	// false
	CanManageRole *bool `json:"canManageRole,omitempty" xml:"canManageRole,omitempty"`
	// This parameter is required.
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// This parameter is required.
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
}

func (s RegisterCustomAppRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterCustomAppRoleRequest) GoString() string {
	return s.String()
}

func (s *RegisterCustomAppRoleRequest) SetCanManageRole(v bool) *RegisterCustomAppRoleRequest {
	s.CanManageRole = &v
	return s
}

func (s *RegisterCustomAppRoleRequest) SetOpUserId(v string) *RegisterCustomAppRoleRequest {
	s.OpUserId = &v
	return s
}

func (s *RegisterCustomAppRoleRequest) SetRoleName(v string) *RegisterCustomAppRoleRequest {
	s.RoleName = &v
	return s
}

type RegisterCustomAppRoleResponseBody struct {
	// example:
	//
	// 123
	RoleId *int64 `json:"roleId,omitempty" xml:"roleId,omitempty"`
	// example:
	//
	// 123123123
	ScopeVersion *int64 `json:"scopeVersion,omitempty" xml:"scopeVersion,omitempty"`
}

func (s RegisterCustomAppRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterCustomAppRoleResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterCustomAppRoleResponseBody) SetRoleId(v int64) *RegisterCustomAppRoleResponseBody {
	s.RoleId = &v
	return s
}

func (s *RegisterCustomAppRoleResponseBody) SetScopeVersion(v int64) *RegisterCustomAppRoleResponseBody {
	s.ScopeVersion = &v
	return s
}

type RegisterCustomAppRoleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterCustomAppRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterCustomAppRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterCustomAppRoleResponse) GoString() string {
	return s.String()
}

func (s *RegisterCustomAppRoleResponse) SetHeaders(v map[string]*string) *RegisterCustomAppRoleResponse {
	s.Headers = v
	return s
}

func (s *RegisterCustomAppRoleResponse) SetStatusCode(v int32) *RegisterCustomAppRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterCustomAppRoleResponse) SetBody(v *RegisterCustomAppRoleResponseBody) *RegisterCustomAppRoleResponse {
	s.Body = v
	return s
}

type RemoveApaasAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RemoveApaasAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveApaasAppHeaders) GoString() string {
	return s.String()
}

func (s *RemoveApaasAppHeaders) SetCommonHeaders(v map[string]*string) *RemoveApaasAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveApaasAppHeaders) SetXAcsDingtalkAccessToken(v string) *RemoveApaasAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RemoveApaasAppRequest struct {
	// This parameter is required.
	BizAppId *string `json:"bizAppId,omitempty" xml:"bizAppId,omitempty"`
	// This parameter is required.
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s RemoveApaasAppRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveApaasAppRequest) GoString() string {
	return s.String()
}

func (s *RemoveApaasAppRequest) SetBizAppId(v string) *RemoveApaasAppRequest {
	s.BizAppId = &v
	return s
}

func (s *RemoveApaasAppRequest) SetOpUserId(v string) *RemoveApaasAppRequest {
	s.OpUserId = &v
	return s
}

type RemoveApaasAppResponseBody struct {
	// This parameter is required.
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RemoveApaasAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveApaasAppResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveApaasAppResponseBody) SetResult(v bool) *RemoveApaasAppResponseBody {
	s.Result = &v
	return s
}

type RemoveApaasAppResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveApaasAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveApaasAppResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveApaasAppResponse) GoString() string {
	return s.String()
}

func (s *RemoveApaasAppResponse) SetHeaders(v map[string]*string) *RemoveApaasAppResponse {
	s.Headers = v
	return s
}

func (s *RemoveApaasAppResponse) SetStatusCode(v int32) *RemoveApaasAppResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveApaasAppResponse) SetBody(v *RemoveApaasAppResponseBody) *RemoveApaasAppResponse {
	s.Body = v
	return s
}

type RemoveMemberForAppRoleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RemoveMemberForAppRoleHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveMemberForAppRoleHeaders) GoString() string {
	return s.String()
}

func (s *RemoveMemberForAppRoleHeaders) SetCommonHeaders(v map[string]*string) *RemoveMemberForAppRoleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveMemberForAppRoleHeaders) SetXAcsDingtalkAccessToken(v string) *RemoveMemberForAppRoleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RemoveMemberForAppRoleRequest struct {
	DeptIdList []*int64 `json:"deptIdList,omitempty" xml:"deptIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// This parameter is required.
	ScopeVersion *int64    `json:"scopeVersion,omitempty" xml:"scopeVersion,omitempty"`
	UserIdList   []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s RemoveMemberForAppRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveMemberForAppRoleRequest) GoString() string {
	return s.String()
}

func (s *RemoveMemberForAppRoleRequest) SetDeptIdList(v []*int64) *RemoveMemberForAppRoleRequest {
	s.DeptIdList = v
	return s
}

func (s *RemoveMemberForAppRoleRequest) SetOpUserId(v string) *RemoveMemberForAppRoleRequest {
	s.OpUserId = &v
	return s
}

func (s *RemoveMemberForAppRoleRequest) SetScopeVersion(v int64) *RemoveMemberForAppRoleRequest {
	s.ScopeVersion = &v
	return s
}

func (s *RemoveMemberForAppRoleRequest) SetUserIdList(v []*string) *RemoveMemberForAppRoleRequest {
	s.UserIdList = v
	return s
}

type RemoveMemberForAppRoleResponseBody struct {
	// example:
	//
	// 123
	LatestScopeVersion *int64 `json:"latestScopeVersion,omitempty" xml:"latestScopeVersion,omitempty"`
}

func (s RemoveMemberForAppRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveMemberForAppRoleResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveMemberForAppRoleResponseBody) SetLatestScopeVersion(v int64) *RemoveMemberForAppRoleResponseBody {
	s.LatestScopeVersion = &v
	return s
}

type RemoveMemberForAppRoleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveMemberForAppRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveMemberForAppRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveMemberForAppRoleResponse) GoString() string {
	return s.String()
}

func (s *RemoveMemberForAppRoleResponse) SetHeaders(v map[string]*string) *RemoveMemberForAppRoleResponse {
	s.Headers = v
	return s
}

func (s *RemoveMemberForAppRoleResponse) SetStatusCode(v int32) *RemoveMemberForAppRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveMemberForAppRoleResponse) SetBody(v *RemoveMemberForAppRoleResponseBody) *RemoveMemberForAppRoleResponse {
	s.Body = v
	return s
}

type RollbackInnerAppVersionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RollbackInnerAppVersionHeaders) String() string {
	return tea.Prettify(s)
}

func (s RollbackInnerAppVersionHeaders) GoString() string {
	return s.String()
}

func (s *RollbackInnerAppVersionHeaders) SetCommonHeaders(v map[string]*string) *RollbackInnerAppVersionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RollbackInnerAppVersionHeaders) SetXAcsDingtalkAccessToken(v string) *RollbackInnerAppVersionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RollbackInnerAppVersionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	AppVersionId *int64 `json:"appVersionId,omitempty" xml:"appVersionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	OpUnionId *string `json:"opUnionId,omitempty" xml:"opUnionId,omitempty"`
}

func (s RollbackInnerAppVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s RollbackInnerAppVersionRequest) GoString() string {
	return s.String()
}

func (s *RollbackInnerAppVersionRequest) SetAppVersionId(v int64) *RollbackInnerAppVersionRequest {
	s.AppVersionId = &v
	return s
}

func (s *RollbackInnerAppVersionRequest) SetOpUnionId(v string) *RollbackInnerAppVersionRequest {
	s.OpUnionId = &v
	return s
}

type RollbackInnerAppVersionResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RollbackInnerAppVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RollbackInnerAppVersionResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackInnerAppVersionResponseBody) SetResult(v bool) *RollbackInnerAppVersionResponseBody {
	s.Result = &v
	return s
}

type RollbackInnerAppVersionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackInnerAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackInnerAppVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s RollbackInnerAppVersionResponse) GoString() string {
	return s.String()
}

func (s *RollbackInnerAppVersionResponse) SetHeaders(v map[string]*string) *RollbackInnerAppVersionResponse {
	s.Headers = v
	return s
}

func (s *RollbackInnerAppVersionResponse) SetStatusCode(v int32) *RollbackInnerAppVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackInnerAppVersionResponse) SetBody(v *RollbackInnerAppVersionResponseBody) *RollbackInnerAppVersionResponse {
	s.Body = v
	return s
}

type SetMicroAppScopeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SetMicroAppScopeHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetMicroAppScopeHeaders) GoString() string {
	return s.String()
}

func (s *SetMicroAppScopeHeaders) SetCommonHeaders(v map[string]*string) *SetMicroAppScopeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetMicroAppScopeHeaders) SetXAcsDingtalkAccessToken(v string) *SetMicroAppScopeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SetMicroAppScopeRequest struct {
	AddDeptIds       []*int64  `json:"addDeptIds,omitempty" xml:"addDeptIds,omitempty" type:"Repeated"`
	AddRoleIds       []*int64  `json:"addRoleIds,omitempty" xml:"addRoleIds,omitempty" type:"Repeated"`
	AddUserIds       []*string `json:"addUserIds,omitempty" xml:"addUserIds,omitempty" type:"Repeated"`
	DelDeptIds       []*int64  `json:"delDeptIds,omitempty" xml:"delDeptIds,omitempty" type:"Repeated"`
	DelRoleIds       []*int64  `json:"delRoleIds,omitempty" xml:"delRoleIds,omitempty" type:"Repeated"`
	DelUserIds       []*string `json:"delUserIds,omitempty" xml:"delUserIds,omitempty" type:"Repeated"`
	OnlyAdminVisible *bool     `json:"onlyAdminVisible,omitempty" xml:"onlyAdminVisible,omitempty"`
}

func (s SetMicroAppScopeRequest) String() string {
	return tea.Prettify(s)
}

func (s SetMicroAppScopeRequest) GoString() string {
	return s.String()
}

func (s *SetMicroAppScopeRequest) SetAddDeptIds(v []*int64) *SetMicroAppScopeRequest {
	s.AddDeptIds = v
	return s
}

func (s *SetMicroAppScopeRequest) SetAddRoleIds(v []*int64) *SetMicroAppScopeRequest {
	s.AddRoleIds = v
	return s
}

func (s *SetMicroAppScopeRequest) SetAddUserIds(v []*string) *SetMicroAppScopeRequest {
	s.AddUserIds = v
	return s
}

func (s *SetMicroAppScopeRequest) SetDelDeptIds(v []*int64) *SetMicroAppScopeRequest {
	s.DelDeptIds = v
	return s
}

func (s *SetMicroAppScopeRequest) SetDelRoleIds(v []*int64) *SetMicroAppScopeRequest {
	s.DelRoleIds = v
	return s
}

func (s *SetMicroAppScopeRequest) SetDelUserIds(v []*string) *SetMicroAppScopeRequest {
	s.DelUserIds = v
	return s
}

func (s *SetMicroAppScopeRequest) SetOnlyAdminVisible(v bool) *SetMicroAppScopeRequest {
	s.OnlyAdminVisible = &v
	return s
}

type SetMicroAppScopeResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SetMicroAppScopeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetMicroAppScopeResponseBody) GoString() string {
	return s.String()
}

func (s *SetMicroAppScopeResponseBody) SetResult(v bool) *SetMicroAppScopeResponseBody {
	s.Result = &v
	return s
}

type SetMicroAppScopeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetMicroAppScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetMicroAppScopeResponse) String() string {
	return tea.Prettify(s)
}

func (s SetMicroAppScopeResponse) GoString() string {
	return s.String()
}

func (s *SetMicroAppScopeResponse) SetHeaders(v map[string]*string) *SetMicroAppScopeResponse {
	s.Headers = v
	return s
}

func (s *SetMicroAppScopeResponse) SetStatusCode(v int32) *SetMicroAppScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetMicroAppScopeResponse) SetBody(v *SetMicroAppScopeResponseBody) *SetMicroAppScopeResponse {
	s.Body = v
	return s
}

type UpdateApaasAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateApaasAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateApaasAppHeaders) GoString() string {
	return s.String()
}

func (s *UpdateApaasAppHeaders) SetCommonHeaders(v map[string]*string) *UpdateApaasAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateApaasAppHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateApaasAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateApaasAppRequest struct {
	// This parameter is required.
	AppIcon *string `json:"appIcon,omitempty" xml:"appIcon,omitempty"`
	// This parameter is required.
	AppName   *string `json:"appName,omitempty" xml:"appName,omitempty"`
	AppStatus *int32  `json:"appStatus,omitempty" xml:"appStatus,omitempty"`
	// This parameter is required.
	BizAppId *string `json:"bizAppId,omitempty" xml:"bizAppId,omitempty"`
	// This parameter is required.
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s UpdateApaasAppRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApaasAppRequest) GoString() string {
	return s.String()
}

func (s *UpdateApaasAppRequest) SetAppIcon(v string) *UpdateApaasAppRequest {
	s.AppIcon = &v
	return s
}

func (s *UpdateApaasAppRequest) SetAppName(v string) *UpdateApaasAppRequest {
	s.AppName = &v
	return s
}

func (s *UpdateApaasAppRequest) SetAppStatus(v int32) *UpdateApaasAppRequest {
	s.AppStatus = &v
	return s
}

func (s *UpdateApaasAppRequest) SetBizAppId(v string) *UpdateApaasAppRequest {
	s.BizAppId = &v
	return s
}

func (s *UpdateApaasAppRequest) SetOpUserId(v string) *UpdateApaasAppRequest {
	s.OpUserId = &v
	return s
}

type UpdateApaasAppResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// This parameter is required.
	BizAppId *string `json:"bizAppId,omitempty" xml:"bizAppId,omitempty"`
}

func (s UpdateApaasAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateApaasAppResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApaasAppResponseBody) SetAgentId(v int64) *UpdateApaasAppResponseBody {
	s.AgentId = &v
	return s
}

func (s *UpdateApaasAppResponseBody) SetBizAppId(v string) *UpdateApaasAppResponseBody {
	s.BizAppId = &v
	return s
}

type UpdateApaasAppResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApaasAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApaasAppResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApaasAppResponse) GoString() string {
	return s.String()
}

func (s *UpdateApaasAppResponse) SetHeaders(v map[string]*string) *UpdateApaasAppResponse {
	s.Headers = v
	return s
}

func (s *UpdateApaasAppResponse) SetStatusCode(v int32) *UpdateApaasAppResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApaasAppResponse) SetBody(v *UpdateApaasAppResponseBody) *UpdateApaasAppResponse {
	s.Body = v
	return s
}

type UpdateAppRoleInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateAppRoleInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppRoleInfoHeaders) GoString() string {
	return s.String()
}

func (s *UpdateAppRoleInfoHeaders) SetCommonHeaders(v map[string]*string) *UpdateAppRoleInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateAppRoleInfoHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateAppRoleInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateAppRoleInfoRequest struct {
	CanManageRole *bool   `json:"canManageRole,omitempty" xml:"canManageRole,omitempty"`
	NewRoleName   *string `json:"newRoleName,omitempty" xml:"newRoleName,omitempty"`
	// This parameter is required.
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s UpdateAppRoleInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppRoleInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppRoleInfoRequest) SetCanManageRole(v bool) *UpdateAppRoleInfoRequest {
	s.CanManageRole = &v
	return s
}

func (s *UpdateAppRoleInfoRequest) SetNewRoleName(v string) *UpdateAppRoleInfoRequest {
	s.NewRoleName = &v
	return s
}

func (s *UpdateAppRoleInfoRequest) SetOpUserId(v string) *UpdateAppRoleInfoRequest {
	s.OpUserId = &v
	return s
}

type UpdateAppRoleInfoResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateAppRoleInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppRoleInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppRoleInfoResponseBody) SetResult(v bool) *UpdateAppRoleInfoResponseBody {
	s.Result = &v
	return s
}

type UpdateAppRoleInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppRoleInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppRoleInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppRoleInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppRoleInfoResponse) SetHeaders(v map[string]*string) *UpdateAppRoleInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppRoleInfoResponse) SetStatusCode(v int32) *UpdateAppRoleInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppRoleInfoResponse) SetBody(v *UpdateAppRoleInfoResponseBody) *UpdateAppRoleInfoResponse {
	s.Body = v
	return s
}

type UpdateInnerAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateInnerAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateInnerAppHeaders) GoString() string {
	return s.String()
}

func (s *UpdateInnerAppHeaders) SetCommonHeaders(v map[string]*string) *UpdateInnerAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateInnerAppHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateInnerAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateInnerAppRequest struct {
	// example:
	//
	// descxxx
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	HomepageLink *string `json:"homepageLink,omitempty" xml:"homepageLink,omitempty"`
	// example:
	//
	// mediaxxx
	Icon        *string   `json:"icon,omitempty" xml:"icon,omitempty"`
	IpWhiteList []*string `json:"ipWhiteList,omitempty" xml:"ipWhiteList,omitempty" type:"Repeated"`
	// example:
	//
	// namexx
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	OmpLink *string `json:"ompLink,omitempty" xml:"ompLink,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	OpUnionId *string `json:"opUnionId,omitempty" xml:"opUnionId,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	PcHomepageLink *string `json:"pcHomepageLink,omitempty" xml:"pcHomepageLink,omitempty"`
}

func (s UpdateInnerAppRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInnerAppRequest) GoString() string {
	return s.String()
}

func (s *UpdateInnerAppRequest) SetDesc(v string) *UpdateInnerAppRequest {
	s.Desc = &v
	return s
}

func (s *UpdateInnerAppRequest) SetHomepageLink(v string) *UpdateInnerAppRequest {
	s.HomepageLink = &v
	return s
}

func (s *UpdateInnerAppRequest) SetIcon(v string) *UpdateInnerAppRequest {
	s.Icon = &v
	return s
}

func (s *UpdateInnerAppRequest) SetIpWhiteList(v []*string) *UpdateInnerAppRequest {
	s.IpWhiteList = v
	return s
}

func (s *UpdateInnerAppRequest) SetName(v string) *UpdateInnerAppRequest {
	s.Name = &v
	return s
}

func (s *UpdateInnerAppRequest) SetOmpLink(v string) *UpdateInnerAppRequest {
	s.OmpLink = &v
	return s
}

func (s *UpdateInnerAppRequest) SetOpUnionId(v string) *UpdateInnerAppRequest {
	s.OpUnionId = &v
	return s
}

func (s *UpdateInnerAppRequest) SetPcHomepageLink(v string) *UpdateInnerAppRequest {
	s.PcHomepageLink = &v
	return s
}

type UpdateInnerAppResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateInnerAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInnerAppResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInnerAppResponseBody) SetResult(v bool) *UpdateInnerAppResponseBody {
	s.Result = &v
	return s
}

type UpdateInnerAppResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInnerAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInnerAppResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInnerAppResponse) GoString() string {
	return s.String()
}

func (s *UpdateInnerAppResponse) SetHeaders(v map[string]*string) *UpdateInnerAppResponse {
	s.Headers = v
	return s
}

func (s *UpdateInnerAppResponse) SetStatusCode(v int32) *UpdateInnerAppResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInnerAppResponse) SetBody(v *UpdateInnerAppResponseBody) *UpdateInnerAppResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

// Summary:
//
// 给指定成员添加角色
//
// @param request - AddAppRolesToMemberRequest
//
// @param headers - AddAppRolesToMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAppRolesToMemberResponse
func (client *Client) AddAppRolesToMemberWithOptions(agentId *string, request *AddAppRolesToMemberRequest, headers *AddAppRolesToMemberHeaders, runtime *util.RuntimeOptions) (_result *AddAppRolesToMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		body["memberId"] = request.MemberId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberType)) {
		body["memberType"] = request.MemberType
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleList)) {
		body["roleList"] = request.RoleList
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
		Action:      tea.String("AddAppRolesToMember"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/apps/" + tea.StringValue(agentId) + "/members/roles"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddAppRolesToMemberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 给指定成员添加角色
//
// @param request - AddAppRolesToMemberRequest
//
// @return AddAppRolesToMemberResponse
func (client *Client) AddAppRolesToMember(agentId *string, request *AddAppRolesToMemberRequest) (_result *AddAppRolesToMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddAppRolesToMemberHeaders{}
	_result = &AddAppRolesToMemberResponse{}
	_body, _err := client.AddAppRolesToMemberWithOptions(agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加应用到工作台分组
//
// @param request - AddAppToWorkBenchGroupRequest
//
// @param headers - AddAppToWorkBenchGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAppToWorkBenchGroupResponse
func (client *Client) AddAppToWorkBenchGroupWithOptions(agentId *string, request *AddAppToWorkBenchGroupRequest, headers *AddAppToWorkBenchGroupHeaders, runtime *util.RuntimeOptions) (_result *AddAppToWorkBenchGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ComponentId)) {
		body["componentId"] = request.ComponentId
	}

	if !tea.BoolValue(util.IsUnset(request.EcologicalCorpId)) {
		body["ecologicalCorpId"] = request.EcologicalCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.OpUnionId)) {
		body["opUnionId"] = request.OpUnionId
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
		Action:      tea.String("AddAppToWorkBenchGroup"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/apps/" + tea.StringValue(agentId) + "/addToWorkBenchGroup"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddAppToWorkBenchGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加应用到工作台分组
//
// @param request - AddAppToWorkBenchGroupRequest
//
// @return AddAppToWorkBenchGroupResponse
func (client *Client) AddAppToWorkBenchGroup(agentId *string, request *AddAppToWorkBenchGroupRequest) (_result *AddAppToWorkBenchGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddAppToWorkBenchGroupHeaders{}
	_result = &AddAppToWorkBenchGroupResponse{}
	_body, _err := client.AddAppToWorkBenchGroupWithOptions(agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 给指定角色添加人员
//
// @param request - AddMemberToAppRoleRequest
//
// @param headers - AddMemberToAppRoleHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMemberToAppRoleResponse
func (client *Client) AddMemberToAppRoleWithOptions(agentId *string, roleId *string, request *AddMemberToAppRoleRequest, headers *AddMemberToAppRoleHeaders, runtime *util.RuntimeOptions) (_result *AddMemberToAppRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptIdList)) {
		body["deptIdList"] = request.DeptIdList
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeVersion)) {
		body["scopeVersion"] = request.ScopeVersion
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdList)) {
		body["userIdList"] = request.UserIdList
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
		Action:      tea.String("AddMemberToAppRole"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/apps/" + tea.StringValue(agentId) + "/roles/" + tea.StringValue(roleId) + "/members"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddMemberToAppRoleResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 给指定角色添加人员
//
// @param request - AddMemberToAppRoleRequest
//
// @return AddMemberToAppRoleResponse
func (client *Client) AddMemberToAppRole(agentId *string, roleId *string, request *AddMemberToAppRoleRequest) (_result *AddMemberToAppRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddMemberToAppRoleHeaders{}
	_result = &AddMemberToAppRoleResponse{}
	_body, _err := client.AddMemberToAppRoleWithOptions(agentId, roleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # AnheiP
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AnheiPResponse
func (client *Client) AnheiPWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *AnheiPResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("AnheiP"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/anheiP"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AnheiPResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # AnheiP
//
// @return AnheiPResponse
func (client *Client) AnheiP() (_result *AnheiPResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AnheiPResponse{}
	_body, _err := client.AnheiPWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # AnheiTest888
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AnheiTest888Response
func (client *Client) AnheiTest888WithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *AnheiTest888Response, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("AnheiTest888"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/anheiTest888"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AnheiTest888Response{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # AnheiTest888
//
// @return AnheiTest888Response
func (client *Client) AnheiTest888() (_result *AnheiTest888Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AnheiTest888Response{}
	_body, _err := client.AnheiTest888WithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # AnheiTestB
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AnheiTestBResponse
func (client *Client) AnheiTestBWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *AnheiTestBResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("AnheiTestB"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/anheiTestB"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AnheiTestBResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # AnheiTestB
//
// @return AnheiTestBResponse
func (client *Client) AnheiTestB() (_result *AnheiTestBResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AnheiTestBResponse{}
	_body, _err := client.AnheiTestBWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 暗黑测试
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AnheiTestNineResponse
func (client *Client) AnheiTestNineWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *AnheiTestNineResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("AnheiTestNine"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/anheiTestNine"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AnheiTestNineResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 暗黑测试
//
// @return AnheiTestNineResponse
func (client *Client) AnheiTestNine() (_result *AnheiTestNineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AnheiTestNineResponse{}
	_body, _err := client.AnheiTestNineWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 应用状态管理测试
//
// @param request - AppStatusManagerTestRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AppStatusManagerTestResponse
func (client *Client) AppStatusManagerTestWithOptions(request *AppStatusManagerTestRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AppStatusManagerTestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["requestId"] = request.RequestId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AppStatusManagerTest"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/manager/test"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AppStatusManagerTestResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 应用状态管理测试
//
// @param request - AppStatusManagerTestRequest
//
// @return AppStatusManagerTestResponse
func (client *Client) AppStatusManagerTest(request *AppStatusManagerTestRequest) (_result *AppStatusManagerTestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AppStatusManagerTestResponse{}
	_body, _err := client.AppStatusManagerTestWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 能力开放中心录入测试数据
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AyunTestResponse
func (client *Client) AyunTestWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *AyunTestResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("AyunTest"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/ayun/test"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AyunTestResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 能力开放中心录入测试数据
//
// @return AyunTestResponse
func (client *Client) AyunTest() (_result *AyunTestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AyunTestResponse{}
	_body, _err := client.AyunTestWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// openAPI录入上线后的测试
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AyunTestOnlineResponse
func (client *Client) AyunTestOnlineWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *AyunTestOnlineResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("AyunTestOnline"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/ayunTest"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AyunTestOnlineResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// openAPI录入上线后的测试
//
// @return AyunTestOnlineResponse
func (client *Client) AyunTestOnline() (_result *AyunTestOnlineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AyunTestOnlineResponse{}
	_body, _err := client.AyunTestOnlineWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建Apaas应用
//
// @param request - CreateApaasAppRequest
//
// @param headers - CreateApaasAppHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApaasAppResponse
func (client *Client) CreateApaasAppWithOptions(request *CreateApaasAppRequest, headers *CreateApaasAppHeaders, runtime *util.RuntimeOptions) (_result *CreateApaasAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppDesc)) {
		body["appDesc"] = request.AppDesc
	}

	if !tea.BoolValue(util.IsUnset(request.AppIcon)) {
		body["appIcon"] = request.AppIcon
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["appName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.BizAppId)) {
		body["bizAppId"] = request.BizAppId
	}

	if !tea.BoolValue(util.IsUnset(request.HomepageEditLink)) {
		body["homepageEditLink"] = request.HomepageEditLink
	}

	if !tea.BoolValue(util.IsUnset(request.HomepageLink)) {
		body["homepageLink"] = request.HomepageLink
	}

	if !tea.BoolValue(util.IsUnset(request.IsShortCut)) {
		body["isShortCut"] = request.IsShortCut
	}

	if !tea.BoolValue(util.IsUnset(request.OmpLink)) {
		body["ompLink"] = request.OmpLink
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.PcHomepageEditLink)) {
		body["pcHomepageEditLink"] = request.PcHomepageEditLink
	}

	if !tea.BoolValue(util.IsUnset(request.PcHomepageLink)) {
		body["pcHomepageLink"] = request.PcHomepageLink
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateKey)) {
		body["templateKey"] = request.TemplateKey
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
		Action:      tea.String("CreateApaasApp"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/apaasApps"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateApaasAppResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Apaas应用
//
// @param request - CreateApaasAppRequest
//
// @return CreateApaasAppResponse
func (client *Client) CreateApaasApp(request *CreateApaasAppRequest) (_result *CreateApaasAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateApaasAppHeaders{}
	_result = &CreateApaasAppResponse{}
	_body, _err := client.CreateApaasAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建企业内部应用
//
// @param request - CreateInnerAppRequest
//
// @param headers - CreateInnerAppHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInnerAppResponse
func (client *Client) CreateInnerAppWithOptions(request *CreateInnerAppRequest, headers *CreateInnerAppHeaders, runtime *util.RuntimeOptions) (_result *CreateInnerAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Desc)) {
		body["desc"] = request.Desc
	}

	if !tea.BoolValue(util.IsUnset(request.DevelopType)) {
		body["developType"] = request.DevelopType
	}

	if !tea.BoolValue(util.IsUnset(request.HomepageLink)) {
		body["homepageLink"] = request.HomepageLink
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		body["icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.IpWhiteList)) {
		body["ipWhiteList"] = request.IpWhiteList
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OmpLink)) {
		body["ompLink"] = request.OmpLink
	}

	if !tea.BoolValue(util.IsUnset(request.OpUnionId)) {
		body["opUnionId"] = request.OpUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.PcHomepageLink)) {
		body["pcHomepageLink"] = request.PcHomepageLink
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeType)) {
		body["scopeType"] = request.ScopeType
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
		Action:      tea.String("CreateInnerApp"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/apps"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInnerAppResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建企业内部应用
//
// @param request - CreateInnerAppRequest
//
// @return CreateInnerAppResponse
func (client *Client) CreateInnerApp(request *CreateInnerAppRequest) (_result *CreateInnerAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateInnerAppHeaders{}
	_result = &CreateInnerAppResponse{}
	_body, _err := client.CreateInnerAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除应用角色
//
// @param request - DeleteAppRoleRequest
//
// @param headers - DeleteAppRoleHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppRoleResponse
func (client *Client) DeleteAppRoleWithOptions(agentId *string, roleId *string, request *DeleteAppRoleRequest, headers *DeleteAppRoleHeaders, runtime *util.RuntimeOptions) (_result *DeleteAppRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		query["opUserId"] = request.OpUserId
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
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAppRole"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/apps/" + tea.StringValue(agentId) + "/roles/" + tea.StringValue(roleId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAppRoleResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除应用角色
//
// @param request - DeleteAppRoleRequest
//
// @return DeleteAppRoleResponse
func (client *Client) DeleteAppRole(agentId *string, roleId *string, request *DeleteAppRoleRequest) (_result *DeleteAppRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteAppRoleHeaders{}
	_result = &DeleteAppRoleResponse{}
	_body, _err := client.DeleteAppRoleWithOptions(agentId, roleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除企业内部应用
//
// @param request - DeleteInnerAppRequest
//
// @param headers - DeleteInnerAppHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInnerAppResponse
func (client *Client) DeleteInnerAppWithOptions(agentId *string, request *DeleteInnerAppRequest, headers *DeleteInnerAppHeaders, runtime *util.RuntimeOptions) (_result *DeleteInnerAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUnionId)) {
		query["opUnionId"] = request.OpUnionId
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
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInnerApp"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/apps/" + tea.StringValue(agentId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInnerAppResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除企业内部应用
//
// @param request - DeleteInnerAppRequest
//
// @return DeleteInnerAppResponse
func (client *Client) DeleteInnerApp(agentId *string, request *DeleteInnerAppRequest) (_result *DeleteInnerAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteInnerAppHeaders{}
	_result = &DeleteInnerAppResponse{}
	_body, _err := client.DeleteInnerAppWithOptions(agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询Apaas应用
//
// @param headers - GetApaasAppHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApaasAppResponse
func (client *Client) GetApaasAppWithOptions(bizAppId *string, headers *GetApaasAppHeaders, runtime *util.RuntimeOptions) (_result *GetApaasAppResponse, _err error) {
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
		Action:      tea.String("GetApaasApp"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/apaasApps/" + tea.StringValue(bizAppId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetApaasAppResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Apaas应用
//
// @return GetApaasAppResponse
func (client *Client) GetApaasApp(bizAppId *string) (_result *GetApaasAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetApaasAppHeaders{}
	_result = &GetApaasAppResponse{}
	_body, _err := client.GetApaasAppWithOptions(bizAppId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取应用资源用量信息
//
// @param request - GetAppResourceUseInfoRequest
//
// @param headers - GetAppResourceUseInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppResourceUseInfoResponse
func (client *Client) GetAppResourceUseInfoWithOptions(request *GetAppResourceUseInfoRequest, headers *GetAppResourceUseInfoHeaders, runtime *util.RuntimeOptions) (_result *GetAppResourceUseInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BenefitCode)) {
		query["benefitCode"] = request.BenefitCode
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodType)) {
		query["periodType"] = request.PeriodType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
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
	}
	params := &openapi.Params{
		Action:      tea.String("GetAppResourceUseInfo"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/resources/useInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("array"),
	}
	_result = &GetAppResourceUseInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取应用资源用量信息
//
// @param request - GetAppResourceUseInfoRequest
//
// @return GetAppResourceUseInfoResponse
func (client *Client) GetAppResourceUseInfo(request *GetAppResourceUseInfoRequest) (_result *GetAppResourceUseInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAppResourceUseInfoHeaders{}
	_result = &GetAppResourceUseInfoResponse{}
	_body, _err := client.GetAppResourceUseInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询指定角色的角色范围
//
// @param headers - GetAppRoleScopeByRoleIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppRoleScopeByRoleIdResponse
func (client *Client) GetAppRoleScopeByRoleIdWithOptions(agentId *string, roleId *string, headers *GetAppRoleScopeByRoleIdHeaders, runtime *util.RuntimeOptions) (_result *GetAppRoleScopeByRoleIdResponse, _err error) {
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
		Action:      tea.String("GetAppRoleScopeByRoleId"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/apps/" + tea.StringValue(agentId) + "/roles/" + tea.StringValue(roleId) + "/scopes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppRoleScopeByRoleIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询指定角色的角色范围
//
// @return GetAppRoleScopeByRoleIdResponse
func (client *Client) GetAppRoleScopeByRoleId(agentId *string, roleId *string) (_result *GetAppRoleScopeByRoleIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAppRoleScopeByRoleIdHeaders{}
	_result = &GetAppRoleScopeByRoleIdResponse{}
	_body, _err := client.GetAppRoleScopeByRoleIdWithOptions(agentId, roleId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业内部H5应用
//
// @param request - GetInnerAppRequest
//
// @param headers - GetInnerAppHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInnerAppResponse
func (client *Client) GetInnerAppWithOptions(agentId *string, request *GetInnerAppRequest, headers *GetInnerAppHeaders, runtime *util.RuntimeOptions) (_result *GetInnerAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EcologicalCorpId)) {
		query["ecologicalCorpId"] = request.EcologicalCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.OpUnionId)) {
		query["opUnionId"] = request.OpUnionId
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
	}
	params := &openapi.Params{
		Action:      tea.String("GetInnerApp"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/apps/" + tea.StringValue(agentId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInnerAppResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业内部H5应用
//
// @param request - GetInnerAppRequest
//
// @return GetInnerAppResponse
func (client *Client) GetInnerApp(agentId *string, request *GetInnerAppRequest) (_result *GetInnerAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetInnerAppHeaders{}
	_result = &GetInnerAppResponse{}
	_body, _err := client.GetInnerAppWithOptions(agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取应用可见范围
//
// @param headers - GetMicroAppScopeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMicroAppScopeResponse
func (client *Client) GetMicroAppScopeWithOptions(agentId *string, headers *GetMicroAppScopeHeaders, runtime *util.RuntimeOptions) (_result *GetMicroAppScopeResponse, _err error) {
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
		Action:      tea.String("GetMicroAppScope"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/apps/" + tea.StringValue(agentId) + "/scopes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMicroAppScopeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取应用可见范围
//
// @return GetMicroAppScopeResponse
func (client *Client) GetMicroAppScope(agentId *string) (_result *GetMicroAppScopeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetMicroAppScopeHeaders{}
	_result = &GetMicroAppScopeResponse{}
	_body, _err := client.GetMicroAppScopeWithOptions(agentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户对应用的管理权限
//
// @param headers - GetMicroAppUserAccessHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMicroAppUserAccessResponse
func (client *Client) GetMicroAppUserAccessWithOptions(agentId *string, userId *string, headers *GetMicroAppUserAccessHeaders, runtime *util.RuntimeOptions) (_result *GetMicroAppUserAccessResponse, _err error) {
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
		Action:      tea.String("GetMicroAppUserAccess"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/apps/" + tea.StringValue(agentId) + "/users/" + tea.StringValue(userId) + "/adminAccess"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMicroAppUserAccessResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户对应用的管理权限
//
// @return GetMicroAppUserAccessResponse
func (client *Client) GetMicroAppUserAccess(agentId *string, userId *string) (_result *GetMicroAppUserAccessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetMicroAppUserAccessHeaders{}
	_result = &GetMicroAppUserAccessResponse{}
	_body, _err := client.GetMicroAppUserAccessWithOptions(agentId, userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用户是否拥有开发者权限
//
// @param headers - GetUserAppDevAccessHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserAppDevAccessResponse
func (client *Client) GetUserAppDevAccessWithOptions(userId *string, headers *GetUserAppDevAccessHeaders, runtime *util.RuntimeOptions) (_result *GetUserAppDevAccessResponse, _err error) {
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
		Action:      tea.String("GetUserAppDevAccess"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/users/" + tea.StringValue(userId) + "/devAccesses"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserAppDevAccessResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用户是否拥有开发者权限
//
// @return GetUserAppDevAccessResponse
func (client *Client) GetUserAppDevAccess(userId *string) (_result *GetUserAppDevAccessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserAppDevAccessHeaders{}
	_result = &GetUserAppDevAccessResponse{}
	_body, _err := client.GetUserAppDevAccessWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取版本-仅用于测试
//
// @param request - GetVersionInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVersionInfoResponse
func (client *Client) GetVersionInfoWithOptions(request *GetVersionInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetVersionInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnifiedAppId)) {
		query["unifiedAppId"] = request.UnifiedAppId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["versionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVersionInfo"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/getVersionInfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVersionInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取版本-仅用于测试
//
// @param request - GetVersionInfoRequest
//
// @return GetVersionInfoResponse
func (client *Client) GetVersionInfo(request *GetVersionInfoRequest) (_result *GetVersionInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetVersionInfoResponse{}
	_body, _err := client.GetVersionInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 判断人员是否在应用的可见范围内
//
// @param request - IsOrgMicroAppVisibleByUserIdRequest
//
// @param headers - IsOrgMicroAppVisibleByUserIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IsOrgMicroAppVisibleByUserIdResponse
func (client *Client) IsOrgMicroAppVisibleByUserIdWithOptions(request *IsOrgMicroAppVisibleByUserIdRequest, headers *IsOrgMicroAppVisibleByUserIdHeaders, runtime *util.RuntimeOptions) (_result *IsOrgMicroAppVisibleByUserIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
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
		Action:      tea.String("IsOrgMicroAppVisibleByUserId"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/visibleScopes/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &IsOrgMicroAppVisibleByUserIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 判断人员是否在应用的可见范围内
//
// @param request - IsOrgMicroAppVisibleByUserIdRequest
//
// @return IsOrgMicroAppVisibleByUserIdResponse
func (client *Client) IsOrgMicroAppVisibleByUserId(request *IsOrgMicroAppVisibleByUserIdRequest) (_result *IsOrgMicroAppVisibleByUserIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IsOrgMicroAppVisibleByUserIdHeaders{}
	_result = &IsOrgMicroAppVisibleByUserIdResponse{}
	_body, _err := client.IsOrgMicroAppVisibleByUserIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业所有应用列表
//
// @param headers - ListAllAppHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllAppResponse
func (client *Client) ListAllAppWithOptions(headers *ListAllAppHeaders, runtime *util.RuntimeOptions) (_result *ListAllAppResponse, _err error) {
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
		Action:      tea.String("ListAllApp"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/allApps"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAllAppResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业所有应用列表
//
// @return ListAllAppResponse
func (client *Client) ListAllApp() (_result *ListAllAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListAllAppHeaders{}
	_result = &ListAllAppResponse{}
	_body, _err := client.ListAllAppWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业所有内部应用列表
//
// @param headers - ListAllInnerAppsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllInnerAppsResponse
func (client *Client) ListAllInnerAppsWithOptions(headers *ListAllInnerAppsHeaders, runtime *util.RuntimeOptions) (_result *ListAllInnerAppsResponse, _err error) {
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
		Action:      tea.String("ListAllInnerApps"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/allInnerApps"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAllInnerAppsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业所有内部应用列表
//
// @return ListAllInnerAppsResponse
func (client *Client) ListAllInnerApps() (_result *ListAllInnerAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListAllInnerAppsHeaders{}
	_result = &ListAllInnerAppsResponse{}
	_body, _err := client.ListAllInnerAppsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取单个企业应用信息
//
// @param headers - ListAppByClientIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppByClientIdResponse
func (client *Client) ListAppByClientIdWithOptions(headers *ListAppByClientIdHeaders, runtime *util.RuntimeOptions) (_result *ListAppByClientIdResponse, _err error) {
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
		Action:      tea.String("ListAppByClientId"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/app/detail"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppByClientIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取单个企业应用信息
//
// @return ListAppByClientIdResponse
func (client *Client) ListAppByClientId() (_result *ListAppByClientIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListAppByClientIdHeaders{}
	_result = &ListAppByClientIdResponse{}
	_body, _err := client.ListAppByClientIdWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业应用的角色完整信息
//
// @param request - ListAppRoleScopesRequest
//
// @param headers - ListAppRoleScopesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppRoleScopesResponse
func (client *Client) ListAppRoleScopesWithOptions(agentId *string, request *ListAppRoleScopesRequest, headers *ListAppRoleScopesHeaders, runtime *util.RuntimeOptions) (_result *ListAppRoleScopesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
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
	}
	params := &openapi.Params{
		Action:      tea.String("ListAppRoleScopes"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/apps/" + tea.StringValue(agentId) + "/roles"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppRoleScopesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业应用的角色完整信息
//
// @param request - ListAppRoleScopesRequest
//
// @return ListAppRoleScopesResponse
func (client *Client) ListAppRoleScopes(agentId *string, request *ListAppRoleScopesRequest) (_result *ListAppRoleScopesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListAppRoleScopesHeaders{}
	_result = &ListAppRoleScopesResponse{}
	_body, _err := client.ListAppRoleScopesWithOptions(agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出企业内部H5应用
//
// @param request - ListInnerAppRequest
//
// @param headers - ListInnerAppHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInnerAppResponse
func (client *Client) ListInnerAppWithOptions(request *ListInnerAppRequest, headers *ListInnerAppHeaders, runtime *util.RuntimeOptions) (_result *ListInnerAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EcologicalCorpId)) {
		query["ecologicalCorpId"] = request.EcologicalCorpId
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
	}
	params := &openapi.Params{
		Action:      tea.String("ListInnerApp"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/apps"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInnerAppResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出企业内部H5应用
//
// @param request - ListInnerAppRequest
//
// @return ListInnerAppResponse
func (client *Client) ListInnerApp(request *ListInnerAppRequest) (_result *ListInnerAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListInnerAppHeaders{}
	_result = &ListInnerAppResponse{}
	_body, _err := client.ListInnerAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业内部小程序的版本列表
//
// @param headers - ListInnerAppVersionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInnerAppVersionResponse
func (client *Client) ListInnerAppVersionWithOptions(agentId *string, headers *ListInnerAppVersionHeaders, runtime *util.RuntimeOptions) (_result *ListInnerAppVersionResponse, _err error) {
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
		Action:      tea.String("ListInnerAppVersion"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/innerMiniApps/" + tea.StringValue(agentId) + "/versions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInnerAppVersionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业内部小程序的版本列表
//
// @return ListInnerAppVersionResponse
func (client *Client) ListInnerAppVersion(agentId *string) (_result *ListInnerAppVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListInnerAppVersionHeaders{}
	_result = &ListInnerAppVersionResponse{}
	_body, _err := client.ListInnerAppVersionWithOptions(agentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户在应用中的角色信息列表
//
// @param headers - ListRoleInfoByUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRoleInfoByUserResponse
func (client *Client) ListRoleInfoByUserWithOptions(agentId *string, userId *string, headers *ListRoleInfoByUserHeaders, runtime *util.RuntimeOptions) (_result *ListRoleInfoByUserResponse, _err error) {
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
		Action:      tea.String("ListRoleInfoByUser"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/apps/" + tea.StringValue(agentId) + "/users/" + tea.StringValue(userId) + "/roles"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRoleInfoByUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户在应用中的角色信息列表
//
// @return ListRoleInfoByUserResponse
func (client *Client) ListRoleInfoByUser(agentId *string, userId *string) (_result *ListRoleInfoByUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListRoleInfoByUserHeaders{}
	_result = &ListRoleInfoByUserResponse{}
	_body, _err := client.ListRoleInfoByUserWithOptions(agentId, userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出用户可见的企业应用
//
// @param headers - ListUserVilebleAppHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserVilebleAppResponse
func (client *Client) ListUserVilebleAppWithOptions(userId *string, headers *ListUserVilebleAppHeaders, runtime *util.RuntimeOptions) (_result *ListUserVilebleAppResponse, _err error) {
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
		Action:      tea.String("ListUserVilebleApp"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/users/" + tea.StringValue(userId) + "/apps"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserVilebleAppResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出用户可见的企业应用
//
// @return ListUserVilebleAppResponse
func (client *Client) ListUserVilebleApp(userId *string) (_result *ListUserVilebleAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListUserVilebleAppHeaders{}
	_result = &ListUserVilebleAppResponse{}
	_body, _err := client.ListUserVilebleAppWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业内部小程序历史版本列表
//
// @param request - PageInnerAppHistoryVersionRequest
//
// @param headers - PageInnerAppHistoryVersionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PageInnerAppHistoryVersionResponse
func (client *Client) PageInnerAppHistoryVersionWithOptions(agentId *string, request *PageInnerAppHistoryVersionRequest, headers *PageInnerAppHistoryVersionHeaders, runtime *util.RuntimeOptions) (_result *PageInnerAppHistoryVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
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
	}
	params := &openapi.Params{
		Action:      tea.String("PageInnerAppHistoryVersion"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/innerMiniApps/" + tea.StringValue(agentId) + "/historyVersions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PageInnerAppHistoryVersionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业内部小程序历史版本列表
//
// @param request - PageInnerAppHistoryVersionRequest
//
// @return PageInnerAppHistoryVersionResponse
func (client *Client) PageInnerAppHistoryVersion(agentId *string, request *PageInnerAppHistoryVersionRequest) (_result *PageInnerAppHistoryVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PageInnerAppHistoryVersionHeaders{}
	_result = &PageInnerAppHistoryVersionResponse{}
	_body, _err := client.PageInnerAppHistoryVersionWithOptions(agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发布企业内部小程序版本
//
// @param request - PublishInnerAppVersionRequest
//
// @param headers - PublishInnerAppVersionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishInnerAppVersionResponse
func (client *Client) PublishInnerAppVersionWithOptions(agentId *string, request *PublishInnerAppVersionRequest, headers *PublishInnerAppVersionHeaders, runtime *util.RuntimeOptions) (_result *PublishInnerAppVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersionId)) {
		body["appVersionId"] = request.AppVersionId
	}

	if !tea.BoolValue(util.IsUnset(request.MiniAppOnPc)) {
		body["miniAppOnPc"] = request.MiniAppOnPc
	}

	if !tea.BoolValue(util.IsUnset(request.OpUnionId)) {
		body["opUnionId"] = request.OpUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.PublishType)) {
		body["publishType"] = request.PublishType
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
		Action:      tea.String("PublishInnerAppVersion"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/innerMiniApps/" + tea.StringValue(agentId) + "/versions/publish"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishInnerAppVersionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发布企业内部小程序版本
//
// @param request - PublishInnerAppVersionRequest
//
// @return PublishInnerAppVersionResponse
func (client *Client) PublishInnerAppVersion(agentId *string, request *PublishInnerAppVersionRequest) (_result *PublishInnerAppVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PublishInnerAppVersionHeaders{}
	_result = &PublishInnerAppVersionResponse{}
	_body, _err := client.PublishInnerAppVersionWithOptions(agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重设角色范围
//
// @param request - RebuildRoleScopeForAppRoleRequest
//
// @param headers - RebuildRoleScopeForAppRoleHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebuildRoleScopeForAppRoleResponse
func (client *Client) RebuildRoleScopeForAppRoleWithOptions(agentId *string, roleId *string, request *RebuildRoleScopeForAppRoleRequest, headers *RebuildRoleScopeForAppRoleHeaders, runtime *util.RuntimeOptions) (_result *RebuildRoleScopeForAppRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptIdList)) {
		body["deptIdList"] = request.DeptIdList
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeType)) {
		body["scopeType"] = request.ScopeType
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeVersion)) {
		body["scopeVersion"] = request.ScopeVersion
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdList)) {
		body["userIdList"] = request.UserIdList
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
		Action:      tea.String("RebuildRoleScopeForAppRole"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/apps/" + tea.StringValue(agentId) + "/roles/" + tea.StringValue(roleId) + "/scopes/rebuild"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RebuildRoleScopeForAppRoleResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重设角色范围
//
// @param request - RebuildRoleScopeForAppRoleRequest
//
// @return RebuildRoleScopeForAppRoleResponse
func (client *Client) RebuildRoleScopeForAppRole(agentId *string, roleId *string, request *RebuildRoleScopeForAppRoleRequest) (_result *RebuildRoleScopeForAppRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RebuildRoleScopeForAppRoleHeaders{}
	_result = &RebuildRoleScopeForAppRoleResponse{}
	_body, _err := client.RebuildRoleScopeForAppRoleWithOptions(agentId, roleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 注册自定义应用角色
//
// @param request - RegisterCustomAppRoleRequest
//
// @param headers - RegisterCustomAppRoleHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterCustomAppRoleResponse
func (client *Client) RegisterCustomAppRoleWithOptions(agentId *string, request *RegisterCustomAppRoleRequest, headers *RegisterCustomAppRoleHeaders, runtime *util.RuntimeOptions) (_result *RegisterCustomAppRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CanManageRole)) {
		body["canManageRole"] = request.CanManageRole
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		body["roleName"] = request.RoleName
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
		Action:      tea.String("RegisterCustomAppRole"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/apps/" + tea.StringValue(agentId) + "/roles"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RegisterCustomAppRoleResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 注册自定义应用角色
//
// @param request - RegisterCustomAppRoleRequest
//
// @return RegisterCustomAppRoleResponse
func (client *Client) RegisterCustomAppRole(agentId *string, request *RegisterCustomAppRoleRequest) (_result *RegisterCustomAppRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RegisterCustomAppRoleHeaders{}
	_result = &RegisterCustomAppRoleResponse{}
	_body, _err := client.RegisterCustomAppRoleWithOptions(agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除apaas应用
//
// @param request - RemoveApaasAppRequest
//
// @param headers - RemoveApaasAppHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveApaasAppResponse
func (client *Client) RemoveApaasAppWithOptions(request *RemoveApaasAppRequest, headers *RemoveApaasAppHeaders, runtime *util.RuntimeOptions) (_result *RemoveApaasAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizAppId)) {
		body["bizAppId"] = request.BizAppId
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
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
		Action:      tea.String("RemoveApaasApp"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/apaasApps/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveApaasAppResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除apaas应用
//
// @param request - RemoveApaasAppRequest
//
// @return RemoveApaasAppResponse
func (client *Client) RemoveApaasApp(request *RemoveApaasAppRequest) (_result *RemoveApaasAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveApaasAppHeaders{}
	_result = &RemoveApaasAppResponse{}
	_body, _err := client.RemoveApaasAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定角色下的成员
//
// @param request - RemoveMemberForAppRoleRequest
//
// @param headers - RemoveMemberForAppRoleHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveMemberForAppRoleResponse
func (client *Client) RemoveMemberForAppRoleWithOptions(agentId *string, roleId *string, request *RemoveMemberForAppRoleRequest, headers *RemoveMemberForAppRoleHeaders, runtime *util.RuntimeOptions) (_result *RemoveMemberForAppRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptIdList)) {
		body["deptIdList"] = request.DeptIdList
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeVersion)) {
		body["scopeVersion"] = request.ScopeVersion
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdList)) {
		body["userIdList"] = request.UserIdList
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
		Action:      tea.String("RemoveMemberForAppRole"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/apps/" + tea.StringValue(agentId) + "/roles/" + tea.StringValue(roleId) + "/members/batchRemove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveMemberForAppRoleResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定角色下的成员
//
// @param request - RemoveMemberForAppRoleRequest
//
// @return RemoveMemberForAppRoleResponse
func (client *Client) RemoveMemberForAppRole(agentId *string, roleId *string, request *RemoveMemberForAppRoleRequest) (_result *RemoveMemberForAppRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveMemberForAppRoleHeaders{}
	_result = &RemoveMemberForAppRoleResponse{}
	_body, _err := client.RemoveMemberForAppRoleWithOptions(agentId, roleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 回滚企业内部小程序版本
//
// @param request - RollbackInnerAppVersionRequest
//
// @param headers - RollbackInnerAppVersionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackInnerAppVersionResponse
func (client *Client) RollbackInnerAppVersionWithOptions(agentId *string, request *RollbackInnerAppVersionRequest, headers *RollbackInnerAppVersionHeaders, runtime *util.RuntimeOptions) (_result *RollbackInnerAppVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersionId)) {
		body["appVersionId"] = request.AppVersionId
	}

	if !tea.BoolValue(util.IsUnset(request.OpUnionId)) {
		body["opUnionId"] = request.OpUnionId
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
		Action:      tea.String("RollbackInnerAppVersion"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/innerMiniApps/" + tea.StringValue(agentId) + "/versions/rollback"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RollbackInnerAppVersionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 回滚企业内部小程序版本
//
// @param request - RollbackInnerAppVersionRequest
//
// @return RollbackInnerAppVersionResponse
func (client *Client) RollbackInnerAppVersion(agentId *string, request *RollbackInnerAppVersionRequest) (_result *RollbackInnerAppVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RollbackInnerAppVersionHeaders{}
	_result = &RollbackInnerAppVersionResponse{}
	_body, _err := client.RollbackInnerAppVersionWithOptions(agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置应用可见范围
//
// @param request - SetMicroAppScopeRequest
//
// @param headers - SetMicroAppScopeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetMicroAppScopeResponse
func (client *Client) SetMicroAppScopeWithOptions(agentId *string, request *SetMicroAppScopeRequest, headers *SetMicroAppScopeHeaders, runtime *util.RuntimeOptions) (_result *SetMicroAppScopeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddDeptIds)) {
		body["addDeptIds"] = request.AddDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.AddRoleIds)) {
		body["addRoleIds"] = request.AddRoleIds
	}

	if !tea.BoolValue(util.IsUnset(request.AddUserIds)) {
		body["addUserIds"] = request.AddUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.DelDeptIds)) {
		body["delDeptIds"] = request.DelDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.DelRoleIds)) {
		body["delRoleIds"] = request.DelRoleIds
	}

	if !tea.BoolValue(util.IsUnset(request.DelUserIds)) {
		body["delUserIds"] = request.DelUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.OnlyAdminVisible)) {
		body["onlyAdminVisible"] = request.OnlyAdminVisible
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
		Action:      tea.String("SetMicroAppScope"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/apps/" + tea.StringValue(agentId) + "/scopes"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SetMicroAppScopeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置应用可见范围
//
// @param request - SetMicroAppScopeRequest
//
// @return SetMicroAppScopeResponse
func (client *Client) SetMicroAppScope(agentId *string, request *SetMicroAppScopeRequest) (_result *SetMicroAppScopeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetMicroAppScopeHeaders{}
	_result = &SetMicroAppScopeResponse{}
	_body, _err := client.SetMicroAppScopeWithOptions(agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新apaas应用
//
// @param request - UpdateApaasAppRequest
//
// @param headers - UpdateApaasAppHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApaasAppResponse
func (client *Client) UpdateApaasAppWithOptions(request *UpdateApaasAppRequest, headers *UpdateApaasAppHeaders, runtime *util.RuntimeOptions) (_result *UpdateApaasAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppIcon)) {
		body["appIcon"] = request.AppIcon
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["appName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AppStatus)) {
		body["appStatus"] = request.AppStatus
	}

	if !tea.BoolValue(util.IsUnset(request.BizAppId)) {
		body["bizAppId"] = request.BizAppId
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
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
		Action:      tea.String("UpdateApaasApp"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/apaasApps"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateApaasAppResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新apaas应用
//
// @param request - UpdateApaasAppRequest
//
// @return UpdateApaasAppResponse
func (client *Client) UpdateApaasApp(request *UpdateApaasAppRequest) (_result *UpdateApaasAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateApaasAppHeaders{}
	_result = &UpdateApaasAppResponse{}
	_body, _err := client.UpdateApaasAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新应用角色信息
//
// @param request - UpdateAppRoleInfoRequest
//
// @param headers - UpdateAppRoleInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAppRoleInfoResponse
func (client *Client) UpdateAppRoleInfoWithOptions(agentId *string, roleId *string, request *UpdateAppRoleInfoRequest, headers *UpdateAppRoleInfoHeaders, runtime *util.RuntimeOptions) (_result *UpdateAppRoleInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CanManageRole)) {
		body["canManageRole"] = request.CanManageRole
	}

	if !tea.BoolValue(util.IsUnset(request.NewRoleName)) {
		body["newRoleName"] = request.NewRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
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
		Action:      tea.String("UpdateAppRoleInfo"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/apps/" + tea.StringValue(agentId) + "/roles/" + tea.StringValue(roleId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAppRoleInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新应用角色信息
//
// @param request - UpdateAppRoleInfoRequest
//
// @return UpdateAppRoleInfoResponse
func (client *Client) UpdateAppRoleInfo(agentId *string, roleId *string, request *UpdateAppRoleInfoRequest) (_result *UpdateAppRoleInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateAppRoleInfoHeaders{}
	_result = &UpdateAppRoleInfoResponse{}
	_body, _err := client.UpdateAppRoleInfoWithOptions(agentId, roleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新企业内部应用
//
// @param request - UpdateInnerAppRequest
//
// @param headers - UpdateInnerAppHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInnerAppResponse
func (client *Client) UpdateInnerAppWithOptions(agentId *string, request *UpdateInnerAppRequest, headers *UpdateInnerAppHeaders, runtime *util.RuntimeOptions) (_result *UpdateInnerAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Desc)) {
		body["desc"] = request.Desc
	}

	if !tea.BoolValue(util.IsUnset(request.HomepageLink)) {
		body["homepageLink"] = request.HomepageLink
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		body["icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.IpWhiteList)) {
		body["ipWhiteList"] = request.IpWhiteList
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OmpLink)) {
		body["ompLink"] = request.OmpLink
	}

	if !tea.BoolValue(util.IsUnset(request.OpUnionId)) {
		body["opUnionId"] = request.OpUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.PcHomepageLink)) {
		body["pcHomepageLink"] = request.PcHomepageLink
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
		Action:      tea.String("UpdateInnerApp"),
		Version:     tea.String("microApp_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/microApp/apps/" + tea.StringValue(agentId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInnerAppResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新企业内部应用
//
// @param request - UpdateInnerAppRequest
//
// @return UpdateInnerAppResponse
func (client *Client) UpdateInnerApp(agentId *string, request *UpdateInnerAppRequest) (_result *UpdateInnerAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateInnerAppHeaders{}
	_result = &UpdateInnerAppResponse{}
	_body, _err := client.UpdateInnerAppWithOptions(agentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
