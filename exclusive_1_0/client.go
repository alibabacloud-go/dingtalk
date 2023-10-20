// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package exclusive_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddOrgHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddOrgHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddOrgHeaders) GoString() string {
	return s.String()
}

func (s *AddOrgHeaders) SetCommonHeaders(v map[string]*string) *AddOrgHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddOrgHeaders) SetXAcsDingtalkAccessToken(v string) *AddOrgHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddOrgRequest struct {
	MobileNum *string `json:"mobileNum,omitempty" xml:"mobileNum,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s AddOrgRequest) String() string {
	return tea.Prettify(s)
}

func (s AddOrgRequest) GoString() string {
	return s.String()
}

func (s *AddOrgRequest) SetMobileNum(v string) *AddOrgRequest {
	s.MobileNum = &v
	return s
}

func (s *AddOrgRequest) SetName(v string) *AddOrgRequest {
	s.Name = &v
	return s
}

type AddOrgResponseBody struct {
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s AddOrgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddOrgResponseBody) GoString() string {
	return s.String()
}

func (s *AddOrgResponseBody) SetCorpId(v string) *AddOrgResponseBody {
	s.CorpId = &v
	return s
}

type AddOrgResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddOrgResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddOrgResponse) String() string {
	return tea.Prettify(s)
}

func (s AddOrgResponse) GoString() string {
	return s.String()
}

func (s *AddOrgResponse) SetHeaders(v map[string]*string) *AddOrgResponse {
	s.Headers = v
	return s
}

func (s *AddOrgResponse) SetStatusCode(v int32) *AddOrgResponse {
	s.StatusCode = &v
	return s
}

func (s *AddOrgResponse) SetBody(v *AddOrgResponseBody) *AddOrgResponse {
	s.Body = v
	return s
}

type ApproveProcessCallbackHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ApproveProcessCallbackHeaders) String() string {
	return tea.Prettify(s)
}

func (s ApproveProcessCallbackHeaders) GoString() string {
	return s.String()
}

func (s *ApproveProcessCallbackHeaders) SetCommonHeaders(v map[string]*string) *ApproveProcessCallbackHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ApproveProcessCallbackHeaders) SetXAcsDingtalkAccessToken(v string) *ApproveProcessCallbackHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ApproveProcessCallbackRequest struct {
	AccessKeyId     *string                               `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	AccessKeySecret *string                               `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	Request         *ApproveProcessCallbackRequestRequest `json:"request,omitempty" xml:"request,omitempty" type:"Struct"`
	TargetCorpId    *string                               `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
}

func (s ApproveProcessCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s ApproveProcessCallbackRequest) GoString() string {
	return s.String()
}

func (s *ApproveProcessCallbackRequest) SetAccessKeyId(v string) *ApproveProcessCallbackRequest {
	s.AccessKeyId = &v
	return s
}

func (s *ApproveProcessCallbackRequest) SetAccessKeySecret(v string) *ApproveProcessCallbackRequest {
	s.AccessKeySecret = &v
	return s
}

func (s *ApproveProcessCallbackRequest) SetRequest(v *ApproveProcessCallbackRequestRequest) *ApproveProcessCallbackRequest {
	s.Request = v
	return s
}

func (s *ApproveProcessCallbackRequest) SetTargetCorpId(v string) *ApproveProcessCallbackRequest {
	s.TargetCorpId = &v
	return s
}

type ApproveProcessCallbackRequestRequest struct {
	ApproveResult     *string   `json:"approveResult,omitempty" xml:"approveResult,omitempty"`
	ApproveType       *string   `json:"approveType,omitempty" xml:"approveType,omitempty"`
	Approvers         []*string `json:"approvers,omitempty" xml:"approvers,omitempty" type:"Repeated"`
	CreateTime        *int64    `json:"createTime,omitempty" xml:"createTime,omitempty"`
	EventType         *string   `json:"eventType,omitempty" xml:"eventType,omitempty"`
	FinishTime        *int64    `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	Params            *string   `json:"params,omitempty" xml:"params,omitempty"`
	ProcessInstanceId *string   `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	Title             *string   `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ApproveProcessCallbackRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s ApproveProcessCallbackRequestRequest) GoString() string {
	return s.String()
}

func (s *ApproveProcessCallbackRequestRequest) SetApproveResult(v string) *ApproveProcessCallbackRequestRequest {
	s.ApproveResult = &v
	return s
}

func (s *ApproveProcessCallbackRequestRequest) SetApproveType(v string) *ApproveProcessCallbackRequestRequest {
	s.ApproveType = &v
	return s
}

func (s *ApproveProcessCallbackRequestRequest) SetApprovers(v []*string) *ApproveProcessCallbackRequestRequest {
	s.Approvers = v
	return s
}

func (s *ApproveProcessCallbackRequestRequest) SetCreateTime(v int64) *ApproveProcessCallbackRequestRequest {
	s.CreateTime = &v
	return s
}

func (s *ApproveProcessCallbackRequestRequest) SetEventType(v string) *ApproveProcessCallbackRequestRequest {
	s.EventType = &v
	return s
}

func (s *ApproveProcessCallbackRequestRequest) SetFinishTime(v int64) *ApproveProcessCallbackRequestRequest {
	s.FinishTime = &v
	return s
}

func (s *ApproveProcessCallbackRequestRequest) SetParams(v string) *ApproveProcessCallbackRequestRequest {
	s.Params = &v
	return s
}

func (s *ApproveProcessCallbackRequestRequest) SetProcessInstanceId(v string) *ApproveProcessCallbackRequestRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *ApproveProcessCallbackRequestRequest) SetTitle(v string) *ApproveProcessCallbackRequestRequest {
	s.Title = &v
	return s
}

type ApproveProcessCallbackResponseBody struct {
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ApproveProcessCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApproveProcessCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveProcessCallbackResponseBody) SetSuccess(v string) *ApproveProcessCallbackResponseBody {
	s.Success = &v
	return s
}

type ApproveProcessCallbackResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApproveProcessCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApproveProcessCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s ApproveProcessCallbackResponse) GoString() string {
	return s.String()
}

func (s *ApproveProcessCallbackResponse) SetHeaders(v map[string]*string) *ApproveProcessCallbackResponse {
	s.Headers = v
	return s
}

func (s *ApproveProcessCallbackResponse) SetStatusCode(v int32) *ApproveProcessCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *ApproveProcessCallbackResponse) SetBody(v *ApproveProcessCallbackResponseBody) *ApproveProcessCallbackResponse {
	s.Body = v
	return s
}

type BanOrOpenGroupWordsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BanOrOpenGroupWordsHeaders) String() string {
	return tea.Prettify(s)
}

func (s BanOrOpenGroupWordsHeaders) GoString() string {
	return s.String()
}

func (s *BanOrOpenGroupWordsHeaders) SetCommonHeaders(v map[string]*string) *BanOrOpenGroupWordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BanOrOpenGroupWordsHeaders) SetXAcsDingtalkAccessToken(v string) *BanOrOpenGroupWordsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BanOrOpenGroupWordsRequest struct {
	BanWordsType      *int32  `json:"banWordsType,omitempty" xml:"banWordsType,omitempty"`
	OpenConverationId *string `json:"openConverationId,omitempty" xml:"openConverationId,omitempty"`
}

func (s BanOrOpenGroupWordsRequest) String() string {
	return tea.Prettify(s)
}

func (s BanOrOpenGroupWordsRequest) GoString() string {
	return s.String()
}

func (s *BanOrOpenGroupWordsRequest) SetBanWordsType(v int32) *BanOrOpenGroupWordsRequest {
	s.BanWordsType = &v
	return s
}

func (s *BanOrOpenGroupWordsRequest) SetOpenConverationId(v string) *BanOrOpenGroupWordsRequest {
	s.OpenConverationId = &v
	return s
}

type BanOrOpenGroupWordsResponseBody struct {
	Cause *string `json:"cause,omitempty" xml:"cause,omitempty"`
	Code  *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s BanOrOpenGroupWordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BanOrOpenGroupWordsResponseBody) GoString() string {
	return s.String()
}

func (s *BanOrOpenGroupWordsResponseBody) SetCause(v string) *BanOrOpenGroupWordsResponseBody {
	s.Cause = &v
	return s
}

func (s *BanOrOpenGroupWordsResponseBody) SetCode(v string) *BanOrOpenGroupWordsResponseBody {
	s.Code = &v
	return s
}

type BanOrOpenGroupWordsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BanOrOpenGroupWordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BanOrOpenGroupWordsResponse) String() string {
	return tea.Prettify(s)
}

func (s BanOrOpenGroupWordsResponse) GoString() string {
	return s.String()
}

func (s *BanOrOpenGroupWordsResponse) SetHeaders(v map[string]*string) *BanOrOpenGroupWordsResponse {
	s.Headers = v
	return s
}

func (s *BanOrOpenGroupWordsResponse) SetStatusCode(v int32) *BanOrOpenGroupWordsResponse {
	s.StatusCode = &v
	return s
}

func (s *BanOrOpenGroupWordsResponse) SetBody(v *BanOrOpenGroupWordsResponseBody) *BanOrOpenGroupWordsResponse {
	s.Body = v
	return s
}

type CreateCategoryAndBindingGroupsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateCategoryAndBindingGroupsHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateCategoryAndBindingGroupsHeaders) GoString() string {
	return s.String()
}

func (s *CreateCategoryAndBindingGroupsHeaders) SetCommonHeaders(v map[string]*string) *CreateCategoryAndBindingGroupsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateCategoryAndBindingGroupsHeaders) SetXAcsDingtalkAccessToken(v string) *CreateCategoryAndBindingGroupsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateCategoryAndBindingGroupsRequest struct {
	CategoryName *string  `json:"categoryName,omitempty" xml:"categoryName,omitempty"`
	GroupIds     []*int64 `json:"groupIds,omitempty" xml:"groupIds,omitempty" type:"Repeated"`
}

func (s CreateCategoryAndBindingGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCategoryAndBindingGroupsRequest) GoString() string {
	return s.String()
}

func (s *CreateCategoryAndBindingGroupsRequest) SetCategoryName(v string) *CreateCategoryAndBindingGroupsRequest {
	s.CategoryName = &v
	return s
}

func (s *CreateCategoryAndBindingGroupsRequest) SetGroupIds(v []*int64) *CreateCategoryAndBindingGroupsRequest {
	s.GroupIds = v
	return s
}

type CreateCategoryAndBindingGroupsResponseBody struct {
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s CreateCategoryAndBindingGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCategoryAndBindingGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCategoryAndBindingGroupsResponseBody) SetId(v int64) *CreateCategoryAndBindingGroupsResponseBody {
	s.Id = &v
	return s
}

type CreateCategoryAndBindingGroupsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCategoryAndBindingGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCategoryAndBindingGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCategoryAndBindingGroupsResponse) GoString() string {
	return s.String()
}

func (s *CreateCategoryAndBindingGroupsResponse) SetHeaders(v map[string]*string) *CreateCategoryAndBindingGroupsResponse {
	s.Headers = v
	return s
}

func (s *CreateCategoryAndBindingGroupsResponse) SetStatusCode(v int32) *CreateCategoryAndBindingGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCategoryAndBindingGroupsResponse) SetBody(v *CreateCategoryAndBindingGroupsResponseBody) *CreateCategoryAndBindingGroupsResponse {
	s.Body = v
	return s
}

type CreateMessageCategoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateMessageCategoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateMessageCategoryHeaders) GoString() string {
	return s.String()
}

func (s *CreateMessageCategoryHeaders) SetCommonHeaders(v map[string]*string) *CreateMessageCategoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateMessageCategoryHeaders) SetXAcsDingtalkAccessToken(v string) *CreateMessageCategoryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateMessageCategoryRequest struct {
	CategoryName *string   `json:"categoryName,omitempty" xml:"categoryName,omitempty"`
	GroupIds     []*string `json:"groupIds,omitempty" xml:"groupIds,omitempty" type:"Repeated"`
}

func (s CreateMessageCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMessageCategoryRequest) GoString() string {
	return s.String()
}

func (s *CreateMessageCategoryRequest) SetCategoryName(v string) *CreateMessageCategoryRequest {
	s.CategoryName = &v
	return s
}

func (s *CreateMessageCategoryRequest) SetGroupIds(v []*string) *CreateMessageCategoryRequest {
	s.GroupIds = v
	return s
}

type CreateMessageCategoryResponseBody struct {
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s CreateMessageCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMessageCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMessageCategoryResponseBody) SetId(v int64) *CreateMessageCategoryResponseBody {
	s.Id = &v
	return s
}

type CreateMessageCategoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMessageCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMessageCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMessageCategoryResponse) GoString() string {
	return s.String()
}

func (s *CreateMessageCategoryResponse) SetHeaders(v map[string]*string) *CreateMessageCategoryResponse {
	s.Headers = v
	return s
}

func (s *CreateMessageCategoryResponse) SetStatusCode(v int32) *CreateMessageCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMessageCategoryResponse) SetBody(v *CreateMessageCategoryResponseBody) *CreateMessageCategoryResponse {
	s.Body = v
	return s
}

type CreateRuleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateRuleHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleHeaders) GoString() string {
	return s.String()
}

func (s *CreateRuleHeaders) SetCommonHeaders(v map[string]*string) *CreateRuleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateRuleHeaders) SetXAcsDingtalkAccessToken(v string) *CreateRuleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateRuleRequest struct {
	CustomPlan *CreateRuleRequestCustomPlan `json:"customPlan,omitempty" xml:"customPlan,omitempty" type:"Struct"`
}

func (s CreateRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRuleRequest) SetCustomPlan(v *CreateRuleRequestCustomPlan) *CreateRuleRequest {
	s.CustomPlan = v
	return s
}

type CreateRuleRequestCustomPlan struct {
	CurrentCategoryList  []*string `json:"currentCategoryList,omitempty" xml:"currentCategoryList,omitempty" type:"Repeated"`
	DeptIds              []*int64  `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	PlanName             *string   `json:"planName,omitempty" xml:"planName,omitempty"`
	UnSelectCategoryList []*string `json:"unSelectCategoryList,omitempty" xml:"unSelectCategoryList,omitempty" type:"Repeated"`
	UserIds              []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s CreateRuleRequestCustomPlan) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequestCustomPlan) GoString() string {
	return s.String()
}

func (s *CreateRuleRequestCustomPlan) SetCurrentCategoryList(v []*string) *CreateRuleRequestCustomPlan {
	s.CurrentCategoryList = v
	return s
}

func (s *CreateRuleRequestCustomPlan) SetDeptIds(v []*int64) *CreateRuleRequestCustomPlan {
	s.DeptIds = v
	return s
}

func (s *CreateRuleRequestCustomPlan) SetPlanName(v string) *CreateRuleRequestCustomPlan {
	s.PlanName = &v
	return s
}

func (s *CreateRuleRequestCustomPlan) SetUnSelectCategoryList(v []*string) *CreateRuleRequestCustomPlan {
	s.UnSelectCategoryList = v
	return s
}

func (s *CreateRuleRequestCustomPlan) SetUserIds(v []*string) *CreateRuleRequestCustomPlan {
	s.UserIds = v
	return s
}

type CreateRuleResponseBody struct {
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s CreateRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRuleResponseBody) SetId(v int64) *CreateRuleResponseBody {
	s.Id = &v
	return s
}

type CreateRuleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateRuleResponse) SetHeaders(v map[string]*string) *CreateRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateRuleResponse) SetStatusCode(v int32) *CreateRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRuleResponse) SetBody(v *CreateRuleResponseBody) *CreateRuleResponse {
	s.Body = v
	return s
}

type CreateTrustedDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateTrustedDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateTrustedDeviceHeaders) GoString() string {
	return s.String()
}

func (s *CreateTrustedDeviceHeaders) SetCommonHeaders(v map[string]*string) *CreateTrustedDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTrustedDeviceHeaders) SetXAcsDingtalkAccessToken(v string) *CreateTrustedDeviceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateTrustedDeviceRequest struct {
	Did        *string `json:"did,omitempty" xml:"did,omitempty"`
	MacAddress *string `json:"macAddress,omitempty" xml:"macAddress,omitempty"`
	Platform   *string `json:"platform,omitempty" xml:"platform,omitempty"`
	Status     *int32  `json:"status,omitempty" xml:"status,omitempty"`
	UserId     *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateTrustedDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTrustedDeviceRequest) GoString() string {
	return s.String()
}

func (s *CreateTrustedDeviceRequest) SetDid(v string) *CreateTrustedDeviceRequest {
	s.Did = &v
	return s
}

func (s *CreateTrustedDeviceRequest) SetMacAddress(v string) *CreateTrustedDeviceRequest {
	s.MacAddress = &v
	return s
}

func (s *CreateTrustedDeviceRequest) SetPlatform(v string) *CreateTrustedDeviceRequest {
	s.Platform = &v
	return s
}

func (s *CreateTrustedDeviceRequest) SetStatus(v int32) *CreateTrustedDeviceRequest {
	s.Status = &v
	return s
}

func (s *CreateTrustedDeviceRequest) SetUserId(v string) *CreateTrustedDeviceRequest {
	s.UserId = &v
	return s
}

type CreateTrustedDeviceResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateTrustedDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTrustedDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrustedDeviceResponseBody) SetSuccess(v bool) *CreateTrustedDeviceResponseBody {
	s.Success = &v
	return s
}

type CreateTrustedDeviceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTrustedDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTrustedDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTrustedDeviceResponse) GoString() string {
	return s.String()
}

func (s *CreateTrustedDeviceResponse) SetHeaders(v map[string]*string) *CreateTrustedDeviceResponse {
	s.Headers = v
	return s
}

func (s *CreateTrustedDeviceResponse) SetStatusCode(v int32) *CreateTrustedDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTrustedDeviceResponse) SetBody(v *CreateTrustedDeviceResponseBody) *CreateTrustedDeviceResponse {
	s.Body = v
	return s
}

type CreateTrustedDeviceBatchHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateTrustedDeviceBatchHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateTrustedDeviceBatchHeaders) GoString() string {
	return s.String()
}

func (s *CreateTrustedDeviceBatchHeaders) SetCommonHeaders(v map[string]*string) *CreateTrustedDeviceBatchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTrustedDeviceBatchHeaders) SetXAcsDingtalkAccessToken(v string) *CreateTrustedDeviceBatchHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateTrustedDeviceBatchRequest struct {
	MacAddressList []*string `json:"macAddressList,omitempty" xml:"macAddressList,omitempty" type:"Repeated"`
	Platform       *string   `json:"platform,omitempty" xml:"platform,omitempty"`
	UserId         *string   `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateTrustedDeviceBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTrustedDeviceBatchRequest) GoString() string {
	return s.String()
}

func (s *CreateTrustedDeviceBatchRequest) SetMacAddressList(v []*string) *CreateTrustedDeviceBatchRequest {
	s.MacAddressList = v
	return s
}

func (s *CreateTrustedDeviceBatchRequest) SetPlatform(v string) *CreateTrustedDeviceBatchRequest {
	s.Platform = &v
	return s
}

func (s *CreateTrustedDeviceBatchRequest) SetUserId(v string) *CreateTrustedDeviceBatchRequest {
	s.UserId = &v
	return s
}

type CreateTrustedDeviceBatchResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateTrustedDeviceBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTrustedDeviceBatchResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrustedDeviceBatchResponseBody) SetResult(v bool) *CreateTrustedDeviceBatchResponseBody {
	s.Result = &v
	return s
}

type CreateTrustedDeviceBatchResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTrustedDeviceBatchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTrustedDeviceBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTrustedDeviceBatchResponse) GoString() string {
	return s.String()
}

func (s *CreateTrustedDeviceBatchResponse) SetHeaders(v map[string]*string) *CreateTrustedDeviceBatchResponse {
	s.Headers = v
	return s
}

func (s *CreateTrustedDeviceBatchResponse) SetStatusCode(v int32) *CreateTrustedDeviceBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTrustedDeviceBatchResponse) SetBody(v *CreateTrustedDeviceBatchResponseBody) *CreateTrustedDeviceBatchResponse {
	s.Body = v
	return s
}

type DeleteAcrossCloudStroageConfigsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteAcrossCloudStroageConfigsHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteAcrossCloudStroageConfigsHeaders) GoString() string {
	return s.String()
}

func (s *DeleteAcrossCloudStroageConfigsHeaders) SetCommonHeaders(v map[string]*string) *DeleteAcrossCloudStroageConfigsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteAcrossCloudStroageConfigsHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteAcrossCloudStroageConfigsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteAcrossCloudStroageConfigsResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteAcrossCloudStroageConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAcrossCloudStroageConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAcrossCloudStroageConfigsResponseBody) SetResult(v bool) *DeleteAcrossCloudStroageConfigsResponseBody {
	s.Result = &v
	return s
}

type DeleteAcrossCloudStroageConfigsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAcrossCloudStroageConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAcrossCloudStroageConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAcrossCloudStroageConfigsResponse) GoString() string {
	return s.String()
}

func (s *DeleteAcrossCloudStroageConfigsResponse) SetHeaders(v map[string]*string) *DeleteAcrossCloudStroageConfigsResponse {
	s.Headers = v
	return s
}

func (s *DeleteAcrossCloudStroageConfigsResponse) SetStatusCode(v int32) *DeleteAcrossCloudStroageConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAcrossCloudStroageConfigsResponse) SetBody(v *DeleteAcrossCloudStroageConfigsResponseBody) *DeleteAcrossCloudStroageConfigsResponse {
	s.Body = v
	return s
}

type DeleteCommentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteCommentHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommentHeaders) GoString() string {
	return s.String()
}

func (s *DeleteCommentHeaders) SetCommonHeaders(v map[string]*string) *DeleteCommentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteCommentHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteCommentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteCommentResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *bool              `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommentResponse) GoString() string {
	return s.String()
}

func (s *DeleteCommentResponse) SetHeaders(v map[string]*string) *DeleteCommentResponse {
	s.Headers = v
	return s
}

func (s *DeleteCommentResponse) SetStatusCode(v int32) *DeleteCommentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCommentResponse) SetBody(v bool) *DeleteCommentResponse {
	s.Body = &v
	return s
}

type DeleteTrustedDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteTrustedDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrustedDeviceHeaders) GoString() string {
	return s.String()
}

func (s *DeleteTrustedDeviceHeaders) SetCommonHeaders(v map[string]*string) *DeleteTrustedDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteTrustedDeviceHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteTrustedDeviceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteTrustedDeviceRequest struct {
	KickOff    *bool   `json:"kickOff,omitempty" xml:"kickOff,omitempty"`
	MacAddress *string `json:"macAddress,omitempty" xml:"macAddress,omitempty"`
	UserId     *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeleteTrustedDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrustedDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteTrustedDeviceRequest) SetKickOff(v bool) *DeleteTrustedDeviceRequest {
	s.KickOff = &v
	return s
}

func (s *DeleteTrustedDeviceRequest) SetMacAddress(v string) *DeleteTrustedDeviceRequest {
	s.MacAddress = &v
	return s
}

func (s *DeleteTrustedDeviceRequest) SetUserId(v string) *DeleteTrustedDeviceRequest {
	s.UserId = &v
	return s
}

type DeleteTrustedDeviceResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteTrustedDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrustedDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTrustedDeviceResponseBody) SetSuccess(v bool) *DeleteTrustedDeviceResponseBody {
	s.Success = &v
	return s
}

type DeleteTrustedDeviceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTrustedDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTrustedDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrustedDeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteTrustedDeviceResponse) SetHeaders(v map[string]*string) *DeleteTrustedDeviceResponse {
	s.Headers = v
	return s
}

func (s *DeleteTrustedDeviceResponse) SetStatusCode(v int32) *DeleteTrustedDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTrustedDeviceResponse) SetBody(v *DeleteTrustedDeviceResponseBody) *DeleteTrustedDeviceResponse {
	s.Body = v
	return s
}

type DistributePartnerAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DistributePartnerAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s DistributePartnerAppHeaders) GoString() string {
	return s.String()
}

func (s *DistributePartnerAppHeaders) SetCommonHeaders(v map[string]*string) *DistributePartnerAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DistributePartnerAppHeaders) SetXAcsDingtalkAccessToken(v string) *DistributePartnerAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DistributePartnerAppRequest struct {
	AppId     *int64  `json:"appId,omitempty" xml:"appId,omitempty"`
	DeptId    *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
	Type      *int64  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DistributePartnerAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DistributePartnerAppRequest) GoString() string {
	return s.String()
}

func (s *DistributePartnerAppRequest) SetAppId(v int64) *DistributePartnerAppRequest {
	s.AppId = &v
	return s
}

func (s *DistributePartnerAppRequest) SetDeptId(v int64) *DistributePartnerAppRequest {
	s.DeptId = &v
	return s
}

func (s *DistributePartnerAppRequest) SetSubCorpId(v string) *DistributePartnerAppRequest {
	s.SubCorpId = &v
	return s
}

func (s *DistributePartnerAppRequest) SetType(v int64) *DistributePartnerAppRequest {
	s.Type = &v
	return s
}

type DistributePartnerAppResponseBody struct {
	InviteUrl *string `json:"inviteUrl,omitempty" xml:"inviteUrl,omitempty"`
}

func (s DistributePartnerAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DistributePartnerAppResponseBody) GoString() string {
	return s.String()
}

func (s *DistributePartnerAppResponseBody) SetInviteUrl(v string) *DistributePartnerAppResponseBody {
	s.InviteUrl = &v
	return s
}

type DistributePartnerAppResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DistributePartnerAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DistributePartnerAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DistributePartnerAppResponse) GoString() string {
	return s.String()
}

func (s *DistributePartnerAppResponse) SetHeaders(v map[string]*string) *DistributePartnerAppResponse {
	s.Headers = v
	return s
}

func (s *DistributePartnerAppResponse) SetStatusCode(v int32) *DistributePartnerAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DistributePartnerAppResponse) SetBody(v *DistributePartnerAppResponseBody) *DistributePartnerAppResponse {
	s.Body = v
	return s
}

type ExclusiveCreateDingPortalHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ExclusiveCreateDingPortalHeaders) String() string {
	return tea.Prettify(s)
}

func (s ExclusiveCreateDingPortalHeaders) GoString() string {
	return s.String()
}

func (s *ExclusiveCreateDingPortalHeaders) SetCommonHeaders(v map[string]*string) *ExclusiveCreateDingPortalHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ExclusiveCreateDingPortalHeaders) SetXAcsDingtalkAccessToken(v string) *ExclusiveCreateDingPortalHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ExclusiveCreateDingPortalRequest struct {
	DingPortalName  *string `json:"dingPortalName,omitempty" xml:"dingPortalName,omitempty"`
	TargetCorpId    *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
	TemplateAppUuid *string `json:"templateAppUuid,omitempty" xml:"templateAppUuid,omitempty"`
	TemplateCorpId  *string `json:"templateCorpId,omitempty" xml:"templateCorpId,omitempty"`
}

func (s ExclusiveCreateDingPortalRequest) String() string {
	return tea.Prettify(s)
}

func (s ExclusiveCreateDingPortalRequest) GoString() string {
	return s.String()
}

func (s *ExclusiveCreateDingPortalRequest) SetDingPortalName(v string) *ExclusiveCreateDingPortalRequest {
	s.DingPortalName = &v
	return s
}

func (s *ExclusiveCreateDingPortalRequest) SetTargetCorpId(v string) *ExclusiveCreateDingPortalRequest {
	s.TargetCorpId = &v
	return s
}

func (s *ExclusiveCreateDingPortalRequest) SetTemplateAppUuid(v string) *ExclusiveCreateDingPortalRequest {
	s.TemplateAppUuid = &v
	return s
}

func (s *ExclusiveCreateDingPortalRequest) SetTemplateCorpId(v string) *ExclusiveCreateDingPortalRequest {
	s.TemplateCorpId = &v
	return s
}

type ExclusiveCreateDingPortalResponseBody struct {
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExclusiveCreateDingPortalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExclusiveCreateDingPortalResponseBody) GoString() string {
	return s.String()
}

func (s *ExclusiveCreateDingPortalResponseBody) SetSuccess(v string) *ExclusiveCreateDingPortalResponseBody {
	s.Success = &v
	return s
}

type ExclusiveCreateDingPortalResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExclusiveCreateDingPortalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExclusiveCreateDingPortalResponse) String() string {
	return tea.Prettify(s)
}

func (s ExclusiveCreateDingPortalResponse) GoString() string {
	return s.String()
}

func (s *ExclusiveCreateDingPortalResponse) SetHeaders(v map[string]*string) *ExclusiveCreateDingPortalResponse {
	s.Headers = v
	return s
}

func (s *ExclusiveCreateDingPortalResponse) SetStatusCode(v int32) *ExclusiveCreateDingPortalResponse {
	s.StatusCode = &v
	return s
}

func (s *ExclusiveCreateDingPortalResponse) SetBody(v *ExclusiveCreateDingPortalResponseBody) *ExclusiveCreateDingPortalResponse {
	s.Body = v
	return s
}

type FileStorageActiveStorageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s FileStorageActiveStorageHeaders) String() string {
	return tea.Prettify(s)
}

func (s FileStorageActiveStorageHeaders) GoString() string {
	return s.String()
}

func (s *FileStorageActiveStorageHeaders) SetCommonHeaders(v map[string]*string) *FileStorageActiveStorageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FileStorageActiveStorageHeaders) SetXAcsDingtalkAccessToken(v string) *FileStorageActiveStorageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type FileStorageActiveStorageRequest struct {
	AccessKeyId     *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	Oss             *string `json:"oss,omitempty" xml:"oss,omitempty"`
	TargetCorpId    *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
}

func (s FileStorageActiveStorageRequest) String() string {
	return tea.Prettify(s)
}

func (s FileStorageActiveStorageRequest) GoString() string {
	return s.String()
}

func (s *FileStorageActiveStorageRequest) SetAccessKeyId(v string) *FileStorageActiveStorageRequest {
	s.AccessKeyId = &v
	return s
}

func (s *FileStorageActiveStorageRequest) SetAccessKeySecret(v string) *FileStorageActiveStorageRequest {
	s.AccessKeySecret = &v
	return s
}

func (s *FileStorageActiveStorageRequest) SetOss(v string) *FileStorageActiveStorageRequest {
	s.Oss = &v
	return s
}

func (s *FileStorageActiveStorageRequest) SetTargetCorpId(v string) *FileStorageActiveStorageRequest {
	s.TargetCorpId = &v
	return s
}

type FileStorageActiveStorageResponseBody struct {
	CreateDate            *string `json:"createDate,omitempty" xml:"createDate,omitempty"`
	FileStorageOpenStatus *int32  `json:"fileStorageOpenStatus,omitempty" xml:"fileStorageOpenStatus,omitempty"`
	StorageStatus         *int32  `json:"storageStatus,omitempty" xml:"storageStatus,omitempty"`
	UsedQuota             *int64  `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
}

func (s FileStorageActiveStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FileStorageActiveStorageResponseBody) GoString() string {
	return s.String()
}

func (s *FileStorageActiveStorageResponseBody) SetCreateDate(v string) *FileStorageActiveStorageResponseBody {
	s.CreateDate = &v
	return s
}

func (s *FileStorageActiveStorageResponseBody) SetFileStorageOpenStatus(v int32) *FileStorageActiveStorageResponseBody {
	s.FileStorageOpenStatus = &v
	return s
}

func (s *FileStorageActiveStorageResponseBody) SetStorageStatus(v int32) *FileStorageActiveStorageResponseBody {
	s.StorageStatus = &v
	return s
}

func (s *FileStorageActiveStorageResponseBody) SetUsedQuota(v int64) *FileStorageActiveStorageResponseBody {
	s.UsedQuota = &v
	return s
}

type FileStorageActiveStorageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FileStorageActiveStorageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FileStorageActiveStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s FileStorageActiveStorageResponse) GoString() string {
	return s.String()
}

func (s *FileStorageActiveStorageResponse) SetHeaders(v map[string]*string) *FileStorageActiveStorageResponse {
	s.Headers = v
	return s
}

func (s *FileStorageActiveStorageResponse) SetStatusCode(v int32) *FileStorageActiveStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *FileStorageActiveStorageResponse) SetBody(v *FileStorageActiveStorageResponseBody) *FileStorageActiveStorageResponse {
	s.Body = v
	return s
}

type FileStorageCheckConnectionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s FileStorageCheckConnectionHeaders) String() string {
	return tea.Prettify(s)
}

func (s FileStorageCheckConnectionHeaders) GoString() string {
	return s.String()
}

func (s *FileStorageCheckConnectionHeaders) SetCommonHeaders(v map[string]*string) *FileStorageCheckConnectionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FileStorageCheckConnectionHeaders) SetXAcsDingtalkAccessToken(v string) *FileStorageCheckConnectionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type FileStorageCheckConnectionRequest struct {
	AccessKeyId     *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	Oss             *string `json:"oss,omitempty" xml:"oss,omitempty"`
	TargetCorpId    *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
}

func (s FileStorageCheckConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s FileStorageCheckConnectionRequest) GoString() string {
	return s.String()
}

func (s *FileStorageCheckConnectionRequest) SetAccessKeyId(v string) *FileStorageCheckConnectionRequest {
	s.AccessKeyId = &v
	return s
}

func (s *FileStorageCheckConnectionRequest) SetAccessKeySecret(v string) *FileStorageCheckConnectionRequest {
	s.AccessKeySecret = &v
	return s
}

func (s *FileStorageCheckConnectionRequest) SetOss(v string) *FileStorageCheckConnectionRequest {
	s.Oss = &v
	return s
}

func (s *FileStorageCheckConnectionRequest) SetTargetCorpId(v string) *FileStorageCheckConnectionRequest {
	s.TargetCorpId = &v
	return s
}

type FileStorageCheckConnectionResponseBody struct {
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	CheckState  *int32  `json:"checkState,omitempty" xml:"checkState,omitempty"`
	Oss         *string `json:"oss,omitempty" xml:"oss,omitempty"`
}

func (s FileStorageCheckConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FileStorageCheckConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *FileStorageCheckConnectionResponseBody) SetAccessKeyId(v string) *FileStorageCheckConnectionResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *FileStorageCheckConnectionResponseBody) SetCheckState(v int32) *FileStorageCheckConnectionResponseBody {
	s.CheckState = &v
	return s
}

func (s *FileStorageCheckConnectionResponseBody) SetOss(v string) *FileStorageCheckConnectionResponseBody {
	s.Oss = &v
	return s
}

type FileStorageCheckConnectionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FileStorageCheckConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FileStorageCheckConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s FileStorageCheckConnectionResponse) GoString() string {
	return s.String()
}

func (s *FileStorageCheckConnectionResponse) SetHeaders(v map[string]*string) *FileStorageCheckConnectionResponse {
	s.Headers = v
	return s
}

func (s *FileStorageCheckConnectionResponse) SetStatusCode(v int32) *FileStorageCheckConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *FileStorageCheckConnectionResponse) SetBody(v *FileStorageCheckConnectionResponseBody) *FileStorageCheckConnectionResponse {
	s.Body = v
	return s
}

type FileStorageGetQuotaDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s FileStorageGetQuotaDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s FileStorageGetQuotaDataHeaders) GoString() string {
	return s.String()
}

func (s *FileStorageGetQuotaDataHeaders) SetCommonHeaders(v map[string]*string) *FileStorageGetQuotaDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FileStorageGetQuotaDataHeaders) SetXAcsDingtalkAccessToken(v string) *FileStorageGetQuotaDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type FileStorageGetQuotaDataRequest struct {
	EndTime      *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime    *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
	Type         *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FileStorageGetQuotaDataRequest) String() string {
	return tea.Prettify(s)
}

func (s FileStorageGetQuotaDataRequest) GoString() string {
	return s.String()
}

func (s *FileStorageGetQuotaDataRequest) SetEndTime(v string) *FileStorageGetQuotaDataRequest {
	s.EndTime = &v
	return s
}

func (s *FileStorageGetQuotaDataRequest) SetStartTime(v string) *FileStorageGetQuotaDataRequest {
	s.StartTime = &v
	return s
}

func (s *FileStorageGetQuotaDataRequest) SetTargetCorpId(v string) *FileStorageGetQuotaDataRequest {
	s.TargetCorpId = &v
	return s
}

func (s *FileStorageGetQuotaDataRequest) SetType(v string) *FileStorageGetQuotaDataRequest {
	s.Type = &v
	return s
}

type FileStorageGetQuotaDataResponseBody struct {
	QuotaModelList []*FileStorageGetQuotaDataResponseBodyQuotaModelList `json:"quotaModelList,omitempty" xml:"quotaModelList,omitempty" type:"Repeated"`
}

func (s FileStorageGetQuotaDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FileStorageGetQuotaDataResponseBody) GoString() string {
	return s.String()
}

func (s *FileStorageGetQuotaDataResponseBody) SetQuotaModelList(v []*FileStorageGetQuotaDataResponseBodyQuotaModelList) *FileStorageGetQuotaDataResponseBody {
	s.QuotaModelList = v
	return s
}

type FileStorageGetQuotaDataResponseBodyQuotaModelList struct {
	StatisticTime *string `json:"statisticTime,omitempty" xml:"statisticTime,omitempty"`
	UsedStorage   *int64  `json:"usedStorage,omitempty" xml:"usedStorage,omitempty"`
}

func (s FileStorageGetQuotaDataResponseBodyQuotaModelList) String() string {
	return tea.Prettify(s)
}

func (s FileStorageGetQuotaDataResponseBodyQuotaModelList) GoString() string {
	return s.String()
}

func (s *FileStorageGetQuotaDataResponseBodyQuotaModelList) SetStatisticTime(v string) *FileStorageGetQuotaDataResponseBodyQuotaModelList {
	s.StatisticTime = &v
	return s
}

func (s *FileStorageGetQuotaDataResponseBodyQuotaModelList) SetUsedStorage(v int64) *FileStorageGetQuotaDataResponseBodyQuotaModelList {
	s.UsedStorage = &v
	return s
}

type FileStorageGetQuotaDataResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FileStorageGetQuotaDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FileStorageGetQuotaDataResponse) String() string {
	return tea.Prettify(s)
}

func (s FileStorageGetQuotaDataResponse) GoString() string {
	return s.String()
}

func (s *FileStorageGetQuotaDataResponse) SetHeaders(v map[string]*string) *FileStorageGetQuotaDataResponse {
	s.Headers = v
	return s
}

func (s *FileStorageGetQuotaDataResponse) SetStatusCode(v int32) *FileStorageGetQuotaDataResponse {
	s.StatusCode = &v
	return s
}

func (s *FileStorageGetQuotaDataResponse) SetBody(v *FileStorageGetQuotaDataResponseBody) *FileStorageGetQuotaDataResponse {
	s.Body = v
	return s
}

type FileStorageGetStorageStateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s FileStorageGetStorageStateHeaders) String() string {
	return tea.Prettify(s)
}

func (s FileStorageGetStorageStateHeaders) GoString() string {
	return s.String()
}

func (s *FileStorageGetStorageStateHeaders) SetCommonHeaders(v map[string]*string) *FileStorageGetStorageStateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FileStorageGetStorageStateHeaders) SetXAcsDingtalkAccessToken(v string) *FileStorageGetStorageStateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type FileStorageGetStorageStateRequest struct {
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
}

func (s FileStorageGetStorageStateRequest) String() string {
	return tea.Prettify(s)
}

func (s FileStorageGetStorageStateRequest) GoString() string {
	return s.String()
}

func (s *FileStorageGetStorageStateRequest) SetTargetCorpId(v string) *FileStorageGetStorageStateRequest {
	s.TargetCorpId = &v
	return s
}

type FileStorageGetStorageStateResponseBody struct {
	AccessKeyId           *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	CreateDate            *string `json:"createDate,omitempty" xml:"createDate,omitempty"`
	FileStorageOpenStatus *int32  `json:"fileStorageOpenStatus,omitempty" xml:"fileStorageOpenStatus,omitempty"`
	Oss                   *string `json:"oss,omitempty" xml:"oss,omitempty"`
	StorageStatus         *int32  `json:"storageStatus,omitempty" xml:"storageStatus,omitempty"`
	UsedQuota             *int64  `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
}

func (s FileStorageGetStorageStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FileStorageGetStorageStateResponseBody) GoString() string {
	return s.String()
}

func (s *FileStorageGetStorageStateResponseBody) SetAccessKeyId(v string) *FileStorageGetStorageStateResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *FileStorageGetStorageStateResponseBody) SetCreateDate(v string) *FileStorageGetStorageStateResponseBody {
	s.CreateDate = &v
	return s
}

func (s *FileStorageGetStorageStateResponseBody) SetFileStorageOpenStatus(v int32) *FileStorageGetStorageStateResponseBody {
	s.FileStorageOpenStatus = &v
	return s
}

func (s *FileStorageGetStorageStateResponseBody) SetOss(v string) *FileStorageGetStorageStateResponseBody {
	s.Oss = &v
	return s
}

func (s *FileStorageGetStorageStateResponseBody) SetStorageStatus(v int32) *FileStorageGetStorageStateResponseBody {
	s.StorageStatus = &v
	return s
}

func (s *FileStorageGetStorageStateResponseBody) SetUsedQuota(v int64) *FileStorageGetStorageStateResponseBody {
	s.UsedQuota = &v
	return s
}

type FileStorageGetStorageStateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FileStorageGetStorageStateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FileStorageGetStorageStateResponse) String() string {
	return tea.Prettify(s)
}

func (s FileStorageGetStorageStateResponse) GoString() string {
	return s.String()
}

func (s *FileStorageGetStorageStateResponse) SetHeaders(v map[string]*string) *FileStorageGetStorageStateResponse {
	s.Headers = v
	return s
}

func (s *FileStorageGetStorageStateResponse) SetStatusCode(v int32) *FileStorageGetStorageStateResponse {
	s.StatusCode = &v
	return s
}

func (s *FileStorageGetStorageStateResponse) SetBody(v *FileStorageGetStorageStateResponseBody) *FileStorageGetStorageStateResponse {
	s.Body = v
	return s
}

type FileStorageUpdateStorageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s FileStorageUpdateStorageHeaders) String() string {
	return tea.Prettify(s)
}

func (s FileStorageUpdateStorageHeaders) GoString() string {
	return s.String()
}

func (s *FileStorageUpdateStorageHeaders) SetCommonHeaders(v map[string]*string) *FileStorageUpdateStorageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FileStorageUpdateStorageHeaders) SetXAcsDingtalkAccessToken(v string) *FileStorageUpdateStorageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type FileStorageUpdateStorageRequest struct {
	AccessKeyId     *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	TargetCorpId    *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
}

func (s FileStorageUpdateStorageRequest) String() string {
	return tea.Prettify(s)
}

func (s FileStorageUpdateStorageRequest) GoString() string {
	return s.String()
}

func (s *FileStorageUpdateStorageRequest) SetAccessKeyId(v string) *FileStorageUpdateStorageRequest {
	s.AccessKeyId = &v
	return s
}

func (s *FileStorageUpdateStorageRequest) SetAccessKeySecret(v string) *FileStorageUpdateStorageRequest {
	s.AccessKeySecret = &v
	return s
}

func (s *FileStorageUpdateStorageRequest) SetTargetCorpId(v string) *FileStorageUpdateStorageRequest {
	s.TargetCorpId = &v
	return s
}

type FileStorageUpdateStorageResponseBody struct {
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	Oss         *string `json:"oss,omitempty" xml:"oss,omitempty"`
}

func (s FileStorageUpdateStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FileStorageUpdateStorageResponseBody) GoString() string {
	return s.String()
}

func (s *FileStorageUpdateStorageResponseBody) SetAccessKeyId(v string) *FileStorageUpdateStorageResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *FileStorageUpdateStorageResponseBody) SetOss(v string) *FileStorageUpdateStorageResponseBody {
	s.Oss = &v
	return s
}

type FileStorageUpdateStorageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FileStorageUpdateStorageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FileStorageUpdateStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s FileStorageUpdateStorageResponse) GoString() string {
	return s.String()
}

func (s *FileStorageUpdateStorageResponse) SetHeaders(v map[string]*string) *FileStorageUpdateStorageResponse {
	s.Headers = v
	return s
}

func (s *FileStorageUpdateStorageResponse) SetStatusCode(v int32) *FileStorageUpdateStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *FileStorageUpdateStorageResponse) SetBody(v *FileStorageUpdateStorageResponseBody) *FileStorageUpdateStorageResponse {
	s.Body = v
	return s
}

type GenerateDarkWaterMarkHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GenerateDarkWaterMarkHeaders) String() string {
	return tea.Prettify(s)
}

func (s GenerateDarkWaterMarkHeaders) GoString() string {
	return s.String()
}

func (s *GenerateDarkWaterMarkHeaders) SetCommonHeaders(v map[string]*string) *GenerateDarkWaterMarkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GenerateDarkWaterMarkHeaders) SetXAcsDingtalkAccessToken(v string) *GenerateDarkWaterMarkHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GenerateDarkWaterMarkRequest struct {
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s GenerateDarkWaterMarkRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateDarkWaterMarkRequest) GoString() string {
	return s.String()
}

func (s *GenerateDarkWaterMarkRequest) SetUserIdList(v []*string) *GenerateDarkWaterMarkRequest {
	s.UserIdList = v
	return s
}

type GenerateDarkWaterMarkResponseBody struct {
	DarkWatermarkVOList []*GenerateDarkWaterMarkResponseBodyDarkWatermarkVOList `json:"darkWatermarkVOList,omitempty" xml:"darkWatermarkVOList,omitempty" type:"Repeated"`
}

func (s GenerateDarkWaterMarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateDarkWaterMarkResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateDarkWaterMarkResponseBody) SetDarkWatermarkVOList(v []*GenerateDarkWaterMarkResponseBodyDarkWatermarkVOList) *GenerateDarkWaterMarkResponseBody {
	s.DarkWatermarkVOList = v
	return s
}

type GenerateDarkWaterMarkResponseBodyDarkWatermarkVOList struct {
	DarkWatermark *string `json:"darkWatermark,omitempty" xml:"darkWatermark,omitempty"`
	UserId        *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GenerateDarkWaterMarkResponseBodyDarkWatermarkVOList) String() string {
	return tea.Prettify(s)
}

func (s GenerateDarkWaterMarkResponseBodyDarkWatermarkVOList) GoString() string {
	return s.String()
}

func (s *GenerateDarkWaterMarkResponseBodyDarkWatermarkVOList) SetDarkWatermark(v string) *GenerateDarkWaterMarkResponseBodyDarkWatermarkVOList {
	s.DarkWatermark = &v
	return s
}

func (s *GenerateDarkWaterMarkResponseBodyDarkWatermarkVOList) SetUserId(v string) *GenerateDarkWaterMarkResponseBodyDarkWatermarkVOList {
	s.UserId = &v
	return s
}

type GenerateDarkWaterMarkResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateDarkWaterMarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateDarkWaterMarkResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateDarkWaterMarkResponse) GoString() string {
	return s.String()
}

func (s *GenerateDarkWaterMarkResponse) SetHeaders(v map[string]*string) *GenerateDarkWaterMarkResponse {
	s.Headers = v
	return s
}

func (s *GenerateDarkWaterMarkResponse) SetStatusCode(v int32) *GenerateDarkWaterMarkResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateDarkWaterMarkResponse) SetBody(v *GenerateDarkWaterMarkResponseBody) *GenerateDarkWaterMarkResponse {
	s.Body = v
	return s
}

type GetAccountTransferListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAccountTransferListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAccountTransferListHeaders) GoString() string {
	return s.String()
}

func (s *GetAccountTransferListHeaders) SetCommonHeaders(v map[string]*string) *GetAccountTransferListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAccountTransferListHeaders) SetXAcsDingtalkAccessToken(v string) *GetAccountTransferListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAccountTransferListRequest struct {
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Status     *int64 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetAccountTransferListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccountTransferListRequest) GoString() string {
	return s.String()
}

func (s *GetAccountTransferListRequest) SetPageNumber(v int64) *GetAccountTransferListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetAccountTransferListRequest) SetPageSize(v int64) *GetAccountTransferListRequest {
	s.PageSize = &v
	return s
}

func (s *GetAccountTransferListRequest) SetStatus(v int64) *GetAccountTransferListRequest {
	s.Status = &v
	return s
}

type GetAccountTransferListResponseBody struct {
	ItemList   []*GetAccountTransferListResponseBodyItemList `json:"itemList,omitempty" xml:"itemList,omitempty" type:"Repeated"`
	TotalCount *int64                                        `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetAccountTransferListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccountTransferListResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountTransferListResponseBody) SetItemList(v []*GetAccountTransferListResponseBodyItemList) *GetAccountTransferListResponseBody {
	s.ItemList = v
	return s
}

func (s *GetAccountTransferListResponseBody) SetTotalCount(v int64) *GetAccountTransferListResponseBody {
	s.TotalCount = &v
	return s
}

type GetAccountTransferListResponseBodyItemList struct {
	DeptName *int64  `json:"deptName,omitempty" xml:"deptName,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId   *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetAccountTransferListResponseBodyItemList) String() string {
	return tea.Prettify(s)
}

func (s GetAccountTransferListResponseBodyItemList) GoString() string {
	return s.String()
}

func (s *GetAccountTransferListResponseBodyItemList) SetDeptName(v int64) *GetAccountTransferListResponseBodyItemList {
	s.DeptName = &v
	return s
}

func (s *GetAccountTransferListResponseBodyItemList) SetName(v string) *GetAccountTransferListResponseBodyItemList {
	s.Name = &v
	return s
}

func (s *GetAccountTransferListResponseBodyItemList) SetUserId(v string) *GetAccountTransferListResponseBodyItemList {
	s.UserId = &v
	return s
}

type GetAccountTransferListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAccountTransferListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAccountTransferListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountTransferListResponse) GoString() string {
	return s.String()
}

func (s *GetAccountTransferListResponse) SetHeaders(v map[string]*string) *GetAccountTransferListResponse {
	s.Headers = v
	return s
}

func (s *GetAccountTransferListResponse) SetStatusCode(v int32) *GetAccountTransferListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountTransferListResponse) SetBody(v *GetAccountTransferListResponseBody) *GetAccountTransferListResponse {
	s.Body = v
	return s
}

type GetActiveUserSummaryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetActiveUserSummaryHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetActiveUserSummaryHeaders) GoString() string {
	return s.String()
}

func (s *GetActiveUserSummaryHeaders) SetCommonHeaders(v map[string]*string) *GetActiveUserSummaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetActiveUserSummaryHeaders) SetXAcsDingtalkAccessToken(v string) *GetActiveUserSummaryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetActiveUserSummaryResponseBody struct {
	ActUsrCnt1m *string `json:"actUsrCnt1m,omitempty" xml:"actUsrCnt1m,omitempty"`
}

func (s GetActiveUserSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetActiveUserSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetActiveUserSummaryResponseBody) SetActUsrCnt1m(v string) *GetActiveUserSummaryResponseBody {
	s.ActUsrCnt1m = &v
	return s
}

type GetActiveUserSummaryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetActiveUserSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetActiveUserSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetActiveUserSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetActiveUserSummaryResponse) SetHeaders(v map[string]*string) *GetActiveUserSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetActiveUserSummaryResponse) SetStatusCode(v int32) *GetActiveUserSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetActiveUserSummaryResponse) SetBody(v *GetActiveUserSummaryResponseBody) *GetActiveUserSummaryResponse {
	s.Body = v
	return s
}

type GetAgentIdByRelatedAppIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAgentIdByRelatedAppIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAgentIdByRelatedAppIdHeaders) GoString() string {
	return s.String()
}

func (s *GetAgentIdByRelatedAppIdHeaders) SetCommonHeaders(v map[string]*string) *GetAgentIdByRelatedAppIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAgentIdByRelatedAppIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetAgentIdByRelatedAppIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAgentIdByRelatedAppIdRequest struct {
	AppId        *int64  `json:"appId,omitempty" xml:"appId,omitempty"`
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
}

func (s GetAgentIdByRelatedAppIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgentIdByRelatedAppIdRequest) GoString() string {
	return s.String()
}

func (s *GetAgentIdByRelatedAppIdRequest) SetAppId(v int64) *GetAgentIdByRelatedAppIdRequest {
	s.AppId = &v
	return s
}

func (s *GetAgentIdByRelatedAppIdRequest) SetTargetCorpId(v string) *GetAgentIdByRelatedAppIdRequest {
	s.TargetCorpId = &v
	return s
}

type GetAgentIdByRelatedAppIdResponseBody struct {
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
}

func (s GetAgentIdByRelatedAppIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAgentIdByRelatedAppIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentIdByRelatedAppIdResponseBody) SetAgentId(v int64) *GetAgentIdByRelatedAppIdResponseBody {
	s.AgentId = &v
	return s
}

type GetAgentIdByRelatedAppIdResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAgentIdByRelatedAppIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAgentIdByRelatedAppIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAgentIdByRelatedAppIdResponse) GoString() string {
	return s.String()
}

func (s *GetAgentIdByRelatedAppIdResponse) SetHeaders(v map[string]*string) *GetAgentIdByRelatedAppIdResponse {
	s.Headers = v
	return s
}

func (s *GetAgentIdByRelatedAppIdResponse) SetStatusCode(v int32) *GetAgentIdByRelatedAppIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentIdByRelatedAppIdResponse) SetBody(v *GetAgentIdByRelatedAppIdResponseBody) *GetAgentIdByRelatedAppIdResponse {
	s.Body = v
	return s
}

type GetAllLabelableDeptsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAllLabelableDeptsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAllLabelableDeptsHeaders) GoString() string {
	return s.String()
}

func (s *GetAllLabelableDeptsHeaders) SetCommonHeaders(v map[string]*string) *GetAllLabelableDeptsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAllLabelableDeptsHeaders) SetXAcsDingtalkAccessToken(v string) *GetAllLabelableDeptsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAllLabelableDeptsResponseBody struct {
	Data []*GetAllLabelableDeptsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s GetAllLabelableDeptsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAllLabelableDeptsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllLabelableDeptsResponseBody) SetData(v []*GetAllLabelableDeptsResponseBodyData) *GetAllLabelableDeptsResponseBody {
	s.Data = v
	return s
}

type GetAllLabelableDeptsResponseBodyData struct {
	DeptId               *string                                                   `json:"deptId,omitempty" xml:"deptId,omitempty"`
	DeptName             *string                                                   `json:"deptName,omitempty" xml:"deptName,omitempty"`
	MemberCount          *int64                                                    `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	PartnerLabelVOLevel1 *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel1 `json:"partnerLabelVOLevel1,omitempty" xml:"partnerLabelVOLevel1,omitempty" type:"Struct"`
	PartnerLabelVOLevel2 *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel2 `json:"partnerLabelVOLevel2,omitempty" xml:"partnerLabelVOLevel2,omitempty" type:"Struct"`
	PartnerLabelVOLevel3 *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel3 `json:"partnerLabelVOLevel3,omitempty" xml:"partnerLabelVOLevel3,omitempty" type:"Struct"`
	PartnerLabelVOLevel4 *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel4 `json:"partnerLabelVOLevel4,omitempty" xml:"partnerLabelVOLevel4,omitempty" type:"Struct"`
	PartnerLabelVOLevel5 *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel5 `json:"partnerLabelVOLevel5,omitempty" xml:"partnerLabelVOLevel5,omitempty" type:"Struct"`
	PartnerNum           *string                                                   `json:"partnerNum,omitempty" xml:"partnerNum,omitempty"`
	SuperDeptId          *string                                                   `json:"superDeptId,omitempty" xml:"superDeptId,omitempty"`
}

func (s GetAllLabelableDeptsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAllLabelableDeptsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAllLabelableDeptsResponseBodyData) SetDeptId(v string) *GetAllLabelableDeptsResponseBodyData {
	s.DeptId = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyData) SetDeptName(v string) *GetAllLabelableDeptsResponseBodyData {
	s.DeptName = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyData) SetMemberCount(v int64) *GetAllLabelableDeptsResponseBodyData {
	s.MemberCount = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyData) SetPartnerLabelVOLevel1(v *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel1) *GetAllLabelableDeptsResponseBodyData {
	s.PartnerLabelVOLevel1 = v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyData) SetPartnerLabelVOLevel2(v *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel2) *GetAllLabelableDeptsResponseBodyData {
	s.PartnerLabelVOLevel2 = v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyData) SetPartnerLabelVOLevel3(v *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel3) *GetAllLabelableDeptsResponseBodyData {
	s.PartnerLabelVOLevel3 = v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyData) SetPartnerLabelVOLevel4(v *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel4) *GetAllLabelableDeptsResponseBodyData {
	s.PartnerLabelVOLevel4 = v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyData) SetPartnerLabelVOLevel5(v *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel5) *GetAllLabelableDeptsResponseBodyData {
	s.PartnerLabelVOLevel5 = v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyData) SetPartnerNum(v string) *GetAllLabelableDeptsResponseBodyData {
	s.PartnerNum = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyData) SetSuperDeptId(v string) *GetAllLabelableDeptsResponseBodyData {
	s.SuperDeptId = &v
	return s
}

type GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel1 struct {
	LabelId   *int64  `json:"labelId,omitempty" xml:"labelId,omitempty"`
	LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty"`
	LevelNum  *int64  `json:"levelNum,omitempty" xml:"levelNum,omitempty"`
}

func (s GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel1) String() string {
	return tea.Prettify(s)
}

func (s GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel1) GoString() string {
	return s.String()
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel1) SetLabelId(v int64) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel1 {
	s.LabelId = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel1) SetLabelName(v string) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel1 {
	s.LabelName = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel1) SetLevelNum(v int64) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel1 {
	s.LevelNum = &v
	return s
}

type GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel2 struct {
	LabelId   *int64  `json:"labelId,omitempty" xml:"labelId,omitempty"`
	LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty"`
	LevelNum  *int64  `json:"levelNum,omitempty" xml:"levelNum,omitempty"`
}

func (s GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel2) String() string {
	return tea.Prettify(s)
}

func (s GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel2) GoString() string {
	return s.String()
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel2) SetLabelId(v int64) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel2 {
	s.LabelId = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel2) SetLabelName(v string) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel2 {
	s.LabelName = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel2) SetLevelNum(v int64) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel2 {
	s.LevelNum = &v
	return s
}

type GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel3 struct {
	LabelId   *int64  `json:"labelId,omitempty" xml:"labelId,omitempty"`
	LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty"`
	LevelNum  *int64  `json:"levelNum,omitempty" xml:"levelNum,omitempty"`
}

func (s GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel3) String() string {
	return tea.Prettify(s)
}

func (s GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel3) GoString() string {
	return s.String()
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel3) SetLabelId(v int64) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel3 {
	s.LabelId = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel3) SetLabelName(v string) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel3 {
	s.LabelName = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel3) SetLevelNum(v int64) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel3 {
	s.LevelNum = &v
	return s
}

type GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel4 struct {
	LabelId   *int64  `json:"labelId,omitempty" xml:"labelId,omitempty"`
	LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty"`
	LevelNum  *int64  `json:"levelNum,omitempty" xml:"levelNum,omitempty"`
}

func (s GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel4) String() string {
	return tea.Prettify(s)
}

func (s GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel4) GoString() string {
	return s.String()
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel4) SetLabelId(v int64) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel4 {
	s.LabelId = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel4) SetLabelName(v string) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel4 {
	s.LabelName = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel4) SetLevelNum(v int64) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel4 {
	s.LevelNum = &v
	return s
}

type GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel5 struct {
	LabelId   *int64  `json:"labelId,omitempty" xml:"labelId,omitempty"`
	LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty"`
	LevelNum  *int64  `json:"levelNum,omitempty" xml:"levelNum,omitempty"`
}

func (s GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel5) String() string {
	return tea.Prettify(s)
}

func (s GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel5) GoString() string {
	return s.String()
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel5) SetLabelId(v int64) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel5 {
	s.LabelId = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel5) SetLabelName(v string) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel5 {
	s.LabelName = &v
	return s
}

func (s *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel5) SetLevelNum(v int64) *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel5 {
	s.LevelNum = &v
	return s
}

type GetAllLabelableDeptsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAllLabelableDeptsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAllLabelableDeptsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllLabelableDeptsResponse) GoString() string {
	return s.String()
}

func (s *GetAllLabelableDeptsResponse) SetHeaders(v map[string]*string) *GetAllLabelableDeptsResponse {
	s.Headers = v
	return s
}

func (s *GetAllLabelableDeptsResponse) SetStatusCode(v int32) *GetAllLabelableDeptsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAllLabelableDeptsResponse) SetBody(v *GetAllLabelableDeptsResponseBody) *GetAllLabelableDeptsResponse {
	s.Body = v
	return s
}

type GetAppDispatchInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAppDispatchInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAppDispatchInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetAppDispatchInfoHeaders) SetCommonHeaders(v map[string]*string) *GetAppDispatchInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAppDispatchInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetAppDispatchInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAppDispatchInfoRequest struct {
	EndTime   *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetAppDispatchInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppDispatchInfoRequest) GoString() string {
	return s.String()
}

func (s *GetAppDispatchInfoRequest) SetEndTime(v int64) *GetAppDispatchInfoRequest {
	s.EndTime = &v
	return s
}

func (s *GetAppDispatchInfoRequest) SetStartTime(v int64) *GetAppDispatchInfoRequest {
	s.StartTime = &v
	return s
}

type GetAppDispatchInfoResponseBody struct {
	Android []*GetAppDispatchInfoResponseBodyAndroid `json:"android,omitempty" xml:"android,omitempty" type:"Repeated"`
	IOS     []*GetAppDispatchInfoResponseBodyIOS     `json:"iOS,omitempty" xml:"iOS,omitempty" type:"Repeated"`
	Mac     []*GetAppDispatchInfoResponseBodyMac     `json:"mac,omitempty" xml:"mac,omitempty" type:"Repeated"`
	Windows []*GetAppDispatchInfoResponseBodyWindows `json:"windows,omitempty" xml:"windows,omitempty" type:"Repeated"`
}

func (s GetAppDispatchInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppDispatchInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppDispatchInfoResponseBody) SetAndroid(v []*GetAppDispatchInfoResponseBodyAndroid) *GetAppDispatchInfoResponseBody {
	s.Android = v
	return s
}

func (s *GetAppDispatchInfoResponseBody) SetIOS(v []*GetAppDispatchInfoResponseBodyIOS) *GetAppDispatchInfoResponseBody {
	s.IOS = v
	return s
}

func (s *GetAppDispatchInfoResponseBody) SetMac(v []*GetAppDispatchInfoResponseBodyMac) *GetAppDispatchInfoResponseBody {
	s.Mac = v
	return s
}

func (s *GetAppDispatchInfoResponseBody) SetWindows(v []*GetAppDispatchInfoResponseBodyWindows) *GetAppDispatchInfoResponseBody {
	s.Windows = v
	return s
}

type GetAppDispatchInfoResponseBodyAndroid struct {
	BaseLineVersion *string `json:"baseLineVersion,omitempty" xml:"baseLineVersion,omitempty"`
	DownloadUrl     *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	InGray          *bool   `json:"inGray,omitempty" xml:"inGray,omitempty"`
	PackTime        *int64  `json:"packTime,omitempty" xml:"packTime,omitempty"`
	Platform        *string `json:"platform,omitempty" xml:"platform,omitempty"`
	Version         *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetAppDispatchInfoResponseBodyAndroid) String() string {
	return tea.Prettify(s)
}

func (s GetAppDispatchInfoResponseBodyAndroid) GoString() string {
	return s.String()
}

func (s *GetAppDispatchInfoResponseBodyAndroid) SetBaseLineVersion(v string) *GetAppDispatchInfoResponseBodyAndroid {
	s.BaseLineVersion = &v
	return s
}

func (s *GetAppDispatchInfoResponseBodyAndroid) SetDownloadUrl(v string) *GetAppDispatchInfoResponseBodyAndroid {
	s.DownloadUrl = &v
	return s
}

func (s *GetAppDispatchInfoResponseBodyAndroid) SetInGray(v bool) *GetAppDispatchInfoResponseBodyAndroid {
	s.InGray = &v
	return s
}

func (s *GetAppDispatchInfoResponseBodyAndroid) SetPackTime(v int64) *GetAppDispatchInfoResponseBodyAndroid {
	s.PackTime = &v
	return s
}

func (s *GetAppDispatchInfoResponseBodyAndroid) SetPlatform(v string) *GetAppDispatchInfoResponseBodyAndroid {
	s.Platform = &v
	return s
}

func (s *GetAppDispatchInfoResponseBodyAndroid) SetVersion(v string) *GetAppDispatchInfoResponseBodyAndroid {
	s.Version = &v
	return s
}

type GetAppDispatchInfoResponseBodyIOS struct {
	BaseLineVersion *string                               `json:"baseLineVersion,omitempty" xml:"baseLineVersion,omitempty"`
	DownloadUrl     *string                               `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	Ext             *GetAppDispatchInfoResponseBodyIOSExt `json:"ext,omitempty" xml:"ext,omitempty" type:"Struct"`
	InGray          *bool                                 `json:"inGray,omitempty" xml:"inGray,omitempty"`
	PackTime        *int64                                `json:"packTime,omitempty" xml:"packTime,omitempty"`
	Platform        *string                               `json:"platform,omitempty" xml:"platform,omitempty"`
	Version         *string                               `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetAppDispatchInfoResponseBodyIOS) String() string {
	return tea.Prettify(s)
}

func (s GetAppDispatchInfoResponseBodyIOS) GoString() string {
	return s.String()
}

func (s *GetAppDispatchInfoResponseBodyIOS) SetBaseLineVersion(v string) *GetAppDispatchInfoResponseBodyIOS {
	s.BaseLineVersion = &v
	return s
}

func (s *GetAppDispatchInfoResponseBodyIOS) SetDownloadUrl(v string) *GetAppDispatchInfoResponseBodyIOS {
	s.DownloadUrl = &v
	return s
}

func (s *GetAppDispatchInfoResponseBodyIOS) SetExt(v *GetAppDispatchInfoResponseBodyIOSExt) *GetAppDispatchInfoResponseBodyIOS {
	s.Ext = v
	return s
}

func (s *GetAppDispatchInfoResponseBodyIOS) SetInGray(v bool) *GetAppDispatchInfoResponseBodyIOS {
	s.InGray = &v
	return s
}

func (s *GetAppDispatchInfoResponseBodyIOS) SetPackTime(v int64) *GetAppDispatchInfoResponseBodyIOS {
	s.PackTime = &v
	return s
}

func (s *GetAppDispatchInfoResponseBodyIOS) SetPlatform(v string) *GetAppDispatchInfoResponseBodyIOS {
	s.Platform = &v
	return s
}

func (s *GetAppDispatchInfoResponseBodyIOS) SetVersion(v string) *GetAppDispatchInfoResponseBodyIOS {
	s.Version = &v
	return s
}

type GetAppDispatchInfoResponseBodyIOSExt struct {
	Plist *string `json:"plist,omitempty" xml:"plist,omitempty"`
}

func (s GetAppDispatchInfoResponseBodyIOSExt) String() string {
	return tea.Prettify(s)
}

func (s GetAppDispatchInfoResponseBodyIOSExt) GoString() string {
	return s.String()
}

func (s *GetAppDispatchInfoResponseBodyIOSExt) SetPlist(v string) *GetAppDispatchInfoResponseBodyIOSExt {
	s.Plist = &v
	return s
}

type GetAppDispatchInfoResponseBodyMac struct {
	BaseLineVersion *string `json:"baseLineVersion,omitempty" xml:"baseLineVersion,omitempty"`
	DownloadUrl     *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	InGray          *bool   `json:"inGray,omitempty" xml:"inGray,omitempty"`
	PackTime        *int64  `json:"packTime,omitempty" xml:"packTime,omitempty"`
	Platform        *string `json:"platform,omitempty" xml:"platform,omitempty"`
	Version         *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetAppDispatchInfoResponseBodyMac) String() string {
	return tea.Prettify(s)
}

func (s GetAppDispatchInfoResponseBodyMac) GoString() string {
	return s.String()
}

func (s *GetAppDispatchInfoResponseBodyMac) SetBaseLineVersion(v string) *GetAppDispatchInfoResponseBodyMac {
	s.BaseLineVersion = &v
	return s
}

func (s *GetAppDispatchInfoResponseBodyMac) SetDownloadUrl(v string) *GetAppDispatchInfoResponseBodyMac {
	s.DownloadUrl = &v
	return s
}

func (s *GetAppDispatchInfoResponseBodyMac) SetInGray(v bool) *GetAppDispatchInfoResponseBodyMac {
	s.InGray = &v
	return s
}

func (s *GetAppDispatchInfoResponseBodyMac) SetPackTime(v int64) *GetAppDispatchInfoResponseBodyMac {
	s.PackTime = &v
	return s
}

func (s *GetAppDispatchInfoResponseBodyMac) SetPlatform(v string) *GetAppDispatchInfoResponseBodyMac {
	s.Platform = &v
	return s
}

func (s *GetAppDispatchInfoResponseBodyMac) SetVersion(v string) *GetAppDispatchInfoResponseBodyMac {
	s.Version = &v
	return s
}

type GetAppDispatchInfoResponseBodyWindows struct {
	BaseLineVersion *string `json:"baseLineVersion,omitempty" xml:"baseLineVersion,omitempty"`
	DownloadUrl     *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	InGray          *bool   `json:"inGray,omitempty" xml:"inGray,omitempty"`
	PackTime        *int64  `json:"packTime,omitempty" xml:"packTime,omitempty"`
	Platform        *string `json:"platform,omitempty" xml:"platform,omitempty"`
	Version         *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetAppDispatchInfoResponseBodyWindows) String() string {
	return tea.Prettify(s)
}

func (s GetAppDispatchInfoResponseBodyWindows) GoString() string {
	return s.String()
}

func (s *GetAppDispatchInfoResponseBodyWindows) SetBaseLineVersion(v string) *GetAppDispatchInfoResponseBodyWindows {
	s.BaseLineVersion = &v
	return s
}

func (s *GetAppDispatchInfoResponseBodyWindows) SetDownloadUrl(v string) *GetAppDispatchInfoResponseBodyWindows {
	s.DownloadUrl = &v
	return s
}

func (s *GetAppDispatchInfoResponseBodyWindows) SetInGray(v bool) *GetAppDispatchInfoResponseBodyWindows {
	s.InGray = &v
	return s
}

func (s *GetAppDispatchInfoResponseBodyWindows) SetPackTime(v int64) *GetAppDispatchInfoResponseBodyWindows {
	s.PackTime = &v
	return s
}

func (s *GetAppDispatchInfoResponseBodyWindows) SetPlatform(v string) *GetAppDispatchInfoResponseBodyWindows {
	s.Platform = &v
	return s
}

func (s *GetAppDispatchInfoResponseBodyWindows) SetVersion(v string) *GetAppDispatchInfoResponseBodyWindows {
	s.Version = &v
	return s
}

type GetAppDispatchInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAppDispatchInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAppDispatchInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppDispatchInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAppDispatchInfoResponse) SetHeaders(v map[string]*string) *GetAppDispatchInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAppDispatchInfoResponse) SetStatusCode(v int32) *GetAppDispatchInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppDispatchInfoResponse) SetBody(v *GetAppDispatchInfoResponseBody) *GetAppDispatchInfoResponse {
	s.Body = v
	return s
}

type GetCalenderSummaryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCalenderSummaryHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCalenderSummaryHeaders) GoString() string {
	return s.String()
}

func (s *GetCalenderSummaryHeaders) SetCommonHeaders(v map[string]*string) *GetCalenderSummaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCalenderSummaryHeaders) SetXAcsDingtalkAccessToken(v string) *GetCalenderSummaryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCalenderSummaryResponseBody struct {
	CalendarCreateUserCnt *string `json:"calendarCreateUserCnt,omitempty" xml:"calendarCreateUserCnt,omitempty"`
	RecvCalendarUserCnt1d *string `json:"recvCalendarUserCnt1d,omitempty" xml:"recvCalendarUserCnt1d,omitempty"`
	UseCalendarUserCnt1d  *string `json:"useCalendarUserCnt1d,omitempty" xml:"useCalendarUserCnt1d,omitempty"`
}

func (s GetCalenderSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCalenderSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetCalenderSummaryResponseBody) SetCalendarCreateUserCnt(v string) *GetCalenderSummaryResponseBody {
	s.CalendarCreateUserCnt = &v
	return s
}

func (s *GetCalenderSummaryResponseBody) SetRecvCalendarUserCnt1d(v string) *GetCalenderSummaryResponseBody {
	s.RecvCalendarUserCnt1d = &v
	return s
}

func (s *GetCalenderSummaryResponseBody) SetUseCalendarUserCnt1d(v string) *GetCalenderSummaryResponseBody {
	s.UseCalendarUserCnt1d = &v
	return s
}

type GetCalenderSummaryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCalenderSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCalenderSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCalenderSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetCalenderSummaryResponse) SetHeaders(v map[string]*string) *GetCalenderSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetCalenderSummaryResponse) SetStatusCode(v int32) *GetCalenderSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCalenderSummaryResponse) SetBody(v *GetCalenderSummaryResponseBody) *GetCalenderSummaryResponse {
	s.Body = v
	return s
}

type GetCommentListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCommentListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCommentListHeaders) GoString() string {
	return s.String()
}

func (s *GetCommentListHeaders) SetCommonHeaders(v map[string]*string) *GetCommentListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCommentListHeaders) SetXAcsDingtalkAccessToken(v string) *GetCommentListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCommentListRequest struct {
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetCommentListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCommentListRequest) GoString() string {
	return s.String()
}

func (s *GetCommentListRequest) SetPageNumber(v int64) *GetCommentListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetCommentListRequest) SetPageSize(v int64) *GetCommentListRequest {
	s.PageSize = &v
	return s
}

type GetCommentListResponseBody struct {
	Data       []*GetCommentListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	TotalCount *int64                            `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetCommentListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCommentListResponseBody) GoString() string {
	return s.String()
}

func (s *GetCommentListResponseBody) SetData(v []*GetCommentListResponseBodyData) *GetCommentListResponseBody {
	s.Data = v
	return s
}

func (s *GetCommentListResponseBody) SetTotalCount(v int64) *GetCommentListResponseBody {
	s.TotalCount = &v
	return s
}

type GetCommentListResponseBodyData struct {
	CommentId       *string  `json:"commentId,omitempty" xml:"commentId,omitempty"`
	CommentTime     *float32 `json:"commentTime,omitempty" xml:"commentTime,omitempty"`
	CommentUserName *string  `json:"commentUserName,omitempty" xml:"commentUserName,omitempty"`
	Content         *string  `json:"content,omitempty" xml:"content,omitempty"`
}

func (s GetCommentListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCommentListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCommentListResponseBodyData) SetCommentId(v string) *GetCommentListResponseBodyData {
	s.CommentId = &v
	return s
}

func (s *GetCommentListResponseBodyData) SetCommentTime(v float32) *GetCommentListResponseBodyData {
	s.CommentTime = &v
	return s
}

func (s *GetCommentListResponseBodyData) SetCommentUserName(v string) *GetCommentListResponseBodyData {
	s.CommentUserName = &v
	return s
}

func (s *GetCommentListResponseBodyData) SetContent(v string) *GetCommentListResponseBodyData {
	s.Content = &v
	return s
}

type GetCommentListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCommentListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCommentListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCommentListResponse) GoString() string {
	return s.String()
}

func (s *GetCommentListResponse) SetHeaders(v map[string]*string) *GetCommentListResponse {
	s.Headers = v
	return s
}

func (s *GetCommentListResponse) SetStatusCode(v int32) *GetCommentListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCommentListResponse) SetBody(v *GetCommentListResponseBody) *GetCommentListResponse {
	s.Body = v
	return s
}

type GetConfBaseInfoByLogicalIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetConfBaseInfoByLogicalIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetConfBaseInfoByLogicalIdHeaders) GoString() string {
	return s.String()
}

func (s *GetConfBaseInfoByLogicalIdHeaders) SetCommonHeaders(v map[string]*string) *GetConfBaseInfoByLogicalIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetConfBaseInfoByLogicalIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetConfBaseInfoByLogicalIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetConfBaseInfoByLogicalIdRequest struct {
	LogicalConferenceId *string `json:"logicalConferenceId,omitempty" xml:"logicalConferenceId,omitempty"`
}

func (s GetConfBaseInfoByLogicalIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConfBaseInfoByLogicalIdRequest) GoString() string {
	return s.String()
}

func (s *GetConfBaseInfoByLogicalIdRequest) SetLogicalConferenceId(v string) *GetConfBaseInfoByLogicalIdRequest {
	s.LogicalConferenceId = &v
	return s
}

type GetConfBaseInfoByLogicalIdResponseBody struct {
	ConferenceId        *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	LogicalConferenceId *string `json:"logicalConferenceId,omitempty" xml:"logicalConferenceId,omitempty"`
	Nickname            *string `json:"nickname,omitempty" xml:"nickname,omitempty"`
	StartTime           *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Title               *string `json:"title,omitempty" xml:"title,omitempty"`
	UnionId             *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetConfBaseInfoByLogicalIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConfBaseInfoByLogicalIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetConfBaseInfoByLogicalIdResponseBody) SetConferenceId(v string) *GetConfBaseInfoByLogicalIdResponseBody {
	s.ConferenceId = &v
	return s
}

func (s *GetConfBaseInfoByLogicalIdResponseBody) SetLogicalConferenceId(v string) *GetConfBaseInfoByLogicalIdResponseBody {
	s.LogicalConferenceId = &v
	return s
}

func (s *GetConfBaseInfoByLogicalIdResponseBody) SetNickname(v string) *GetConfBaseInfoByLogicalIdResponseBody {
	s.Nickname = &v
	return s
}

func (s *GetConfBaseInfoByLogicalIdResponseBody) SetStartTime(v int64) *GetConfBaseInfoByLogicalIdResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetConfBaseInfoByLogicalIdResponseBody) SetTitle(v string) *GetConfBaseInfoByLogicalIdResponseBody {
	s.Title = &v
	return s
}

func (s *GetConfBaseInfoByLogicalIdResponseBody) SetUnionId(v string) *GetConfBaseInfoByLogicalIdResponseBody {
	s.UnionId = &v
	return s
}

type GetConfBaseInfoByLogicalIdResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetConfBaseInfoByLogicalIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConfBaseInfoByLogicalIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConfBaseInfoByLogicalIdResponse) GoString() string {
	return s.String()
}

func (s *GetConfBaseInfoByLogicalIdResponse) SetHeaders(v map[string]*string) *GetConfBaseInfoByLogicalIdResponse {
	s.Headers = v
	return s
}

func (s *GetConfBaseInfoByLogicalIdResponse) SetStatusCode(v int32) *GetConfBaseInfoByLogicalIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConfBaseInfoByLogicalIdResponse) SetBody(v *GetConfBaseInfoByLogicalIdResponseBody) *GetConfBaseInfoByLogicalIdResponse {
	s.Body = v
	return s
}

type GetConferenceDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetConferenceDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetConferenceDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetConferenceDetailHeaders) SetCommonHeaders(v map[string]*string) *GetConferenceDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetConferenceDetailHeaders) SetXAcsDingtalkAccessToken(v string) *GetConferenceDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetConferenceDetailResponseBody struct {
	AttendeeNum        *int64                                       `json:"attendeeNum,omitempty" xml:"attendeeNum,omitempty"`
	AttendeePercentage *string                                      `json:"attendeePercentage,omitempty" xml:"attendeePercentage,omitempty"`
	CallerId           *string                                      `json:"callerId,omitempty" xml:"callerId,omitempty"`
	CallerName         *string                                      `json:"callerName,omitempty" xml:"callerName,omitempty"`
	ConfStartTime      *float32                                     `json:"confStartTime,omitempty" xml:"confStartTime,omitempty"`
	ConferenceId       *string                                      `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	Duration           *float32                                     `json:"duration,omitempty" xml:"duration,omitempty"`
	MemberList         []*GetConferenceDetailResponseBodyMemberList `json:"memberList,omitempty" xml:"memberList,omitempty" type:"Repeated"`
	Title              *string                                      `json:"title,omitempty" xml:"title,omitempty"`
	TotalNum           *int64                                       `json:"totalNum,omitempty" xml:"totalNum,omitempty"`
}

func (s GetConferenceDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConferenceDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetConferenceDetailResponseBody) SetAttendeeNum(v int64) *GetConferenceDetailResponseBody {
	s.AttendeeNum = &v
	return s
}

func (s *GetConferenceDetailResponseBody) SetAttendeePercentage(v string) *GetConferenceDetailResponseBody {
	s.AttendeePercentage = &v
	return s
}

func (s *GetConferenceDetailResponseBody) SetCallerId(v string) *GetConferenceDetailResponseBody {
	s.CallerId = &v
	return s
}

func (s *GetConferenceDetailResponseBody) SetCallerName(v string) *GetConferenceDetailResponseBody {
	s.CallerName = &v
	return s
}

func (s *GetConferenceDetailResponseBody) SetConfStartTime(v float32) *GetConferenceDetailResponseBody {
	s.ConfStartTime = &v
	return s
}

func (s *GetConferenceDetailResponseBody) SetConferenceId(v string) *GetConferenceDetailResponseBody {
	s.ConferenceId = &v
	return s
}

func (s *GetConferenceDetailResponseBody) SetDuration(v float32) *GetConferenceDetailResponseBody {
	s.Duration = &v
	return s
}

func (s *GetConferenceDetailResponseBody) SetMemberList(v []*GetConferenceDetailResponseBodyMemberList) *GetConferenceDetailResponseBody {
	s.MemberList = v
	return s
}

func (s *GetConferenceDetailResponseBody) SetTitle(v string) *GetConferenceDetailResponseBody {
	s.Title = &v
	return s
}

func (s *GetConferenceDetailResponseBody) SetTotalNum(v int64) *GetConferenceDetailResponseBody {
	s.TotalNum = &v
	return s
}

type GetConferenceDetailResponseBodyMemberList struct {
	AttendDuration *float32 `json:"attendDuration,omitempty" xml:"attendDuration,omitempty"`
	Name           *string  `json:"name,omitempty" xml:"name,omitempty"`
	StaffId        *string  `json:"staffId,omitempty" xml:"staffId,omitempty"`
	UnionId        *string  `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetConferenceDetailResponseBodyMemberList) String() string {
	return tea.Prettify(s)
}

func (s GetConferenceDetailResponseBodyMemberList) GoString() string {
	return s.String()
}

func (s *GetConferenceDetailResponseBodyMemberList) SetAttendDuration(v float32) *GetConferenceDetailResponseBodyMemberList {
	s.AttendDuration = &v
	return s
}

func (s *GetConferenceDetailResponseBodyMemberList) SetName(v string) *GetConferenceDetailResponseBodyMemberList {
	s.Name = &v
	return s
}

func (s *GetConferenceDetailResponseBodyMemberList) SetStaffId(v string) *GetConferenceDetailResponseBodyMemberList {
	s.StaffId = &v
	return s
}

func (s *GetConferenceDetailResponseBodyMemberList) SetUnionId(v string) *GetConferenceDetailResponseBodyMemberList {
	s.UnionId = &v
	return s
}

type GetConferenceDetailResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetConferenceDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConferenceDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConferenceDetailResponse) GoString() string {
	return s.String()
}

func (s *GetConferenceDetailResponse) SetHeaders(v map[string]*string) *GetConferenceDetailResponse {
	s.Headers = v
	return s
}

func (s *GetConferenceDetailResponse) SetStatusCode(v int32) *GetConferenceDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConferenceDetailResponse) SetBody(v *GetConferenceDetailResponseBody) *GetConferenceDetailResponse {
	s.Body = v
	return s
}

type GetDingReportDeptSummaryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDingReportDeptSummaryHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDingReportDeptSummaryHeaders) GoString() string {
	return s.String()
}

func (s *GetDingReportDeptSummaryHeaders) SetCommonHeaders(v map[string]*string) *GetDingReportDeptSummaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDingReportDeptSummaryHeaders) SetXAcsDingtalkAccessToken(v string) *GetDingReportDeptSummaryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDingReportDeptSummaryRequest struct {
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetDingReportDeptSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDingReportDeptSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetDingReportDeptSummaryRequest) SetMaxResults(v int64) *GetDingReportDeptSummaryRequest {
	s.MaxResults = &v
	return s
}

func (s *GetDingReportDeptSummaryRequest) SetNextToken(v int64) *GetDingReportDeptSummaryRequest {
	s.NextToken = &v
	return s
}

type GetDingReportDeptSummaryResponseBody struct {
	Data      []*GetDingReportDeptSummaryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	HasMore   *bool                                       `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken *int64                                      `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetDingReportDeptSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDingReportDeptSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetDingReportDeptSummaryResponseBody) SetData(v []*GetDingReportDeptSummaryResponseBodyData) *GetDingReportDeptSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetDingReportDeptSummaryResponseBody) SetHasMore(v bool) *GetDingReportDeptSummaryResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetDingReportDeptSummaryResponseBody) SetNextToken(v int64) *GetDingReportDeptSummaryResponseBody {
	s.NextToken = &v
	return s
}

type GetDingReportDeptSummaryResponseBodyData struct {
	DeptId               *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
	DeptName             *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	DingReportSendCnt    *string `json:"dingReportSendCnt,omitempty" xml:"dingReportSendCnt,omitempty"`
	DingReportSendUsrCnt *string `json:"dingReportSendUsrCnt,omitempty" xml:"dingReportSendUsrCnt,omitempty"`
}

func (s GetDingReportDeptSummaryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDingReportDeptSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDingReportDeptSummaryResponseBodyData) SetDeptId(v string) *GetDingReportDeptSummaryResponseBodyData {
	s.DeptId = &v
	return s
}

func (s *GetDingReportDeptSummaryResponseBodyData) SetDeptName(v string) *GetDingReportDeptSummaryResponseBodyData {
	s.DeptName = &v
	return s
}

func (s *GetDingReportDeptSummaryResponseBodyData) SetDingReportSendCnt(v string) *GetDingReportDeptSummaryResponseBodyData {
	s.DingReportSendCnt = &v
	return s
}

func (s *GetDingReportDeptSummaryResponseBodyData) SetDingReportSendUsrCnt(v string) *GetDingReportDeptSummaryResponseBodyData {
	s.DingReportSendUsrCnt = &v
	return s
}

type GetDingReportDeptSummaryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDingReportDeptSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDingReportDeptSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDingReportDeptSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetDingReportDeptSummaryResponse) SetHeaders(v map[string]*string) *GetDingReportDeptSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetDingReportDeptSummaryResponse) SetStatusCode(v int32) *GetDingReportDeptSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDingReportDeptSummaryResponse) SetBody(v *GetDingReportDeptSummaryResponseBody) *GetDingReportDeptSummaryResponse {
	s.Body = v
	return s
}

type GetDingReportSummaryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDingReportSummaryHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDingReportSummaryHeaders) GoString() string {
	return s.String()
}

func (s *GetDingReportSummaryHeaders) SetCommonHeaders(v map[string]*string) *GetDingReportSummaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDingReportSummaryHeaders) SetXAcsDingtalkAccessToken(v string) *GetDingReportSummaryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDingReportSummaryResponseBody struct {
	ReportCommentUserCnt1d *string `json:"reportCommentUserCnt1d,omitempty" xml:"reportCommentUserCnt1d,omitempty"`
}

func (s GetDingReportSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDingReportSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetDingReportSummaryResponseBody) SetReportCommentUserCnt1d(v string) *GetDingReportSummaryResponseBody {
	s.ReportCommentUserCnt1d = &v
	return s
}

type GetDingReportSummaryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDingReportSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDingReportSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDingReportSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetDingReportSummaryResponse) SetHeaders(v map[string]*string) *GetDingReportSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetDingReportSummaryResponse) SetStatusCode(v int32) *GetDingReportSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDingReportSummaryResponse) SetBody(v *GetDingReportSummaryResponseBody) *GetDingReportSummaryResponse {
	s.Body = v
	return s
}

type GetDocCreatedDeptSummaryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDocCreatedDeptSummaryHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDocCreatedDeptSummaryHeaders) GoString() string {
	return s.String()
}

func (s *GetDocCreatedDeptSummaryHeaders) SetCommonHeaders(v map[string]*string) *GetDocCreatedDeptSummaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDocCreatedDeptSummaryHeaders) SetXAcsDingtalkAccessToken(v string) *GetDocCreatedDeptSummaryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDocCreatedDeptSummaryRequest struct {
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetDocCreatedDeptSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocCreatedDeptSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetDocCreatedDeptSummaryRequest) SetMaxResults(v int64) *GetDocCreatedDeptSummaryRequest {
	s.MaxResults = &v
	return s
}

func (s *GetDocCreatedDeptSummaryRequest) SetNextToken(v int64) *GetDocCreatedDeptSummaryRequest {
	s.NextToken = &v
	return s
}

type GetDocCreatedDeptSummaryResponseBody struct {
	Data      []*GetDocCreatedDeptSummaryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	HasMore   *bool                                       `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken *int64                                      `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetDocCreatedDeptSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocCreatedDeptSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocCreatedDeptSummaryResponseBody) SetData(v []*GetDocCreatedDeptSummaryResponseBodyData) *GetDocCreatedDeptSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetDocCreatedDeptSummaryResponseBody) SetHasMore(v bool) *GetDocCreatedDeptSummaryResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetDocCreatedDeptSummaryResponseBody) SetNextToken(v int64) *GetDocCreatedDeptSummaryResponseBody {
	s.NextToken = &v
	return s
}

type GetDocCreatedDeptSummaryResponseBodyData struct {
	CreateDocUserCnt1d *string `json:"createDocUserCnt1d,omitempty" xml:"createDocUserCnt1d,omitempty"`
	DeptId             *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
	DeptName           *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	DocCreatedCnt      *string `json:"docCreatedCnt,omitempty" xml:"docCreatedCnt,omitempty"`
}

func (s GetDocCreatedDeptSummaryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDocCreatedDeptSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocCreatedDeptSummaryResponseBodyData) SetCreateDocUserCnt1d(v string) *GetDocCreatedDeptSummaryResponseBodyData {
	s.CreateDocUserCnt1d = &v
	return s
}

func (s *GetDocCreatedDeptSummaryResponseBodyData) SetDeptId(v string) *GetDocCreatedDeptSummaryResponseBodyData {
	s.DeptId = &v
	return s
}

func (s *GetDocCreatedDeptSummaryResponseBodyData) SetDeptName(v string) *GetDocCreatedDeptSummaryResponseBodyData {
	s.DeptName = &v
	return s
}

func (s *GetDocCreatedDeptSummaryResponseBodyData) SetDocCreatedCnt(v string) *GetDocCreatedDeptSummaryResponseBodyData {
	s.DocCreatedCnt = &v
	return s
}

type GetDocCreatedDeptSummaryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDocCreatedDeptSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDocCreatedDeptSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocCreatedDeptSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetDocCreatedDeptSummaryResponse) SetHeaders(v map[string]*string) *GetDocCreatedDeptSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetDocCreatedDeptSummaryResponse) SetStatusCode(v int32) *GetDocCreatedDeptSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocCreatedDeptSummaryResponse) SetBody(v *GetDocCreatedDeptSummaryResponseBody) *GetDocCreatedDeptSummaryResponse {
	s.Body = v
	return s
}

type GetDocCreatedSummaryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDocCreatedSummaryHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDocCreatedSummaryHeaders) GoString() string {
	return s.String()
}

func (s *GetDocCreatedSummaryHeaders) SetCommonHeaders(v map[string]*string) *GetDocCreatedSummaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDocCreatedSummaryHeaders) SetXAcsDingtalkAccessToken(v string) *GetDocCreatedSummaryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDocCreatedSummaryResponseBody struct {
	DocCreateUserCnt1d *string `json:"docCreateUserCnt1d,omitempty" xml:"docCreateUserCnt1d,omitempty"`
	DocCreatedCnt      *string `json:"docCreatedCnt,omitempty" xml:"docCreatedCnt,omitempty"`
}

func (s GetDocCreatedSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocCreatedSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocCreatedSummaryResponseBody) SetDocCreateUserCnt1d(v string) *GetDocCreatedSummaryResponseBody {
	s.DocCreateUserCnt1d = &v
	return s
}

func (s *GetDocCreatedSummaryResponseBody) SetDocCreatedCnt(v string) *GetDocCreatedSummaryResponseBody {
	s.DocCreatedCnt = &v
	return s
}

type GetDocCreatedSummaryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDocCreatedSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDocCreatedSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocCreatedSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetDocCreatedSummaryResponse) SetHeaders(v map[string]*string) *GetDocCreatedSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetDocCreatedSummaryResponse) SetStatusCode(v int32) *GetDocCreatedSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocCreatedSummaryResponse) SetBody(v *GetDocCreatedSummaryResponseBody) *GetDocCreatedSummaryResponse {
	s.Body = v
	return s
}

type GetExclusiveAccountAllOrgListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetExclusiveAccountAllOrgListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetExclusiveAccountAllOrgListHeaders) GoString() string {
	return s.String()
}

func (s *GetExclusiveAccountAllOrgListHeaders) SetCommonHeaders(v map[string]*string) *GetExclusiveAccountAllOrgListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetExclusiveAccountAllOrgListHeaders) SetXAcsDingtalkAccessToken(v string) *GetExclusiveAccountAllOrgListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetExclusiveAccountAllOrgListRequest struct {
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetExclusiveAccountAllOrgListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExclusiveAccountAllOrgListRequest) GoString() string {
	return s.String()
}

func (s *GetExclusiveAccountAllOrgListRequest) SetUnionId(v string) *GetExclusiveAccountAllOrgListRequest {
	s.UnionId = &v
	return s
}

type GetExclusiveAccountAllOrgListResponseBody struct {
	OrgInfoList []*GetExclusiveAccountAllOrgListResponseBodyOrgInfoList `json:"orgInfoList,omitempty" xml:"orgInfoList,omitempty" type:"Repeated"`
}

func (s GetExclusiveAccountAllOrgListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExclusiveAccountAllOrgListResponseBody) GoString() string {
	return s.String()
}

func (s *GetExclusiveAccountAllOrgListResponseBody) SetOrgInfoList(v []*GetExclusiveAccountAllOrgListResponseBodyOrgInfoList) *GetExclusiveAccountAllOrgListResponseBody {
	s.OrgInfoList = v
	return s
}

type GetExclusiveAccountAllOrgListResponseBodyOrgInfoList struct {
	CorpId      *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	IsMainOrg   *bool   `json:"isMainOrg,omitempty" xml:"isMainOrg,omitempty"`
	LogoUrl     *string `json:"logoUrl,omitempty" xml:"logoUrl,omitempty"`
	OrgFullName *string `json:"orgFullName,omitempty" xml:"orgFullName,omitempty"`
	OrgName     *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
}

func (s GetExclusiveAccountAllOrgListResponseBodyOrgInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetExclusiveAccountAllOrgListResponseBodyOrgInfoList) GoString() string {
	return s.String()
}

func (s *GetExclusiveAccountAllOrgListResponseBodyOrgInfoList) SetCorpId(v string) *GetExclusiveAccountAllOrgListResponseBodyOrgInfoList {
	s.CorpId = &v
	return s
}

func (s *GetExclusiveAccountAllOrgListResponseBodyOrgInfoList) SetIsMainOrg(v bool) *GetExclusiveAccountAllOrgListResponseBodyOrgInfoList {
	s.IsMainOrg = &v
	return s
}

func (s *GetExclusiveAccountAllOrgListResponseBodyOrgInfoList) SetLogoUrl(v string) *GetExclusiveAccountAllOrgListResponseBodyOrgInfoList {
	s.LogoUrl = &v
	return s
}

func (s *GetExclusiveAccountAllOrgListResponseBodyOrgInfoList) SetOrgFullName(v string) *GetExclusiveAccountAllOrgListResponseBodyOrgInfoList {
	s.OrgFullName = &v
	return s
}

func (s *GetExclusiveAccountAllOrgListResponseBodyOrgInfoList) SetOrgName(v string) *GetExclusiveAccountAllOrgListResponseBodyOrgInfoList {
	s.OrgName = &v
	return s
}

type GetExclusiveAccountAllOrgListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetExclusiveAccountAllOrgListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetExclusiveAccountAllOrgListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExclusiveAccountAllOrgListResponse) GoString() string {
	return s.String()
}

func (s *GetExclusiveAccountAllOrgListResponse) SetHeaders(v map[string]*string) *GetExclusiveAccountAllOrgListResponse {
	s.Headers = v
	return s
}

func (s *GetExclusiveAccountAllOrgListResponse) SetStatusCode(v int32) *GetExclusiveAccountAllOrgListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExclusiveAccountAllOrgListResponse) SetBody(v *GetExclusiveAccountAllOrgListResponseBody) *GetExclusiveAccountAllOrgListResponse {
	s.Body = v
	return s
}

type GetGeneralFormCreatedDeptSummaryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetGeneralFormCreatedDeptSummaryHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetGeneralFormCreatedDeptSummaryHeaders) GoString() string {
	return s.String()
}

func (s *GetGeneralFormCreatedDeptSummaryHeaders) SetCommonHeaders(v map[string]*string) *GetGeneralFormCreatedDeptSummaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetGeneralFormCreatedDeptSummaryHeaders) SetXAcsDingtalkAccessToken(v string) *GetGeneralFormCreatedDeptSummaryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetGeneralFormCreatedDeptSummaryRequest struct {
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetGeneralFormCreatedDeptSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGeneralFormCreatedDeptSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetGeneralFormCreatedDeptSummaryRequest) SetMaxResults(v int64) *GetGeneralFormCreatedDeptSummaryRequest {
	s.MaxResults = &v
	return s
}

func (s *GetGeneralFormCreatedDeptSummaryRequest) SetNextToken(v int64) *GetGeneralFormCreatedDeptSummaryRequest {
	s.NextToken = &v
	return s
}

type GetGeneralFormCreatedDeptSummaryResponseBody struct {
	Data      []*GetGeneralFormCreatedDeptSummaryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	HasMore   *bool                                               `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken *int64                                              `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetGeneralFormCreatedDeptSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGeneralFormCreatedDeptSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetGeneralFormCreatedDeptSummaryResponseBody) SetData(v []*GetGeneralFormCreatedDeptSummaryResponseBodyData) *GetGeneralFormCreatedDeptSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetGeneralFormCreatedDeptSummaryResponseBody) SetHasMore(v bool) *GetGeneralFormCreatedDeptSummaryResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetGeneralFormCreatedDeptSummaryResponseBody) SetNextToken(v int64) *GetGeneralFormCreatedDeptSummaryResponseBody {
	s.NextToken = &v
	return s
}

type GetGeneralFormCreatedDeptSummaryResponseBodyData struct {
	DeptId                  *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
	DeptName                *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	GeneralFormCreateCnt1d  *string `json:"generalFormCreateCnt1d,omitempty" xml:"generalFormCreateCnt1d,omitempty"`
	UseGeneralFormUserCnt1d *string `json:"useGeneralFormUserCnt1d,omitempty" xml:"useGeneralFormUserCnt1d,omitempty"`
}

func (s GetGeneralFormCreatedDeptSummaryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetGeneralFormCreatedDeptSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGeneralFormCreatedDeptSummaryResponseBodyData) SetDeptId(v string) *GetGeneralFormCreatedDeptSummaryResponseBodyData {
	s.DeptId = &v
	return s
}

func (s *GetGeneralFormCreatedDeptSummaryResponseBodyData) SetDeptName(v string) *GetGeneralFormCreatedDeptSummaryResponseBodyData {
	s.DeptName = &v
	return s
}

func (s *GetGeneralFormCreatedDeptSummaryResponseBodyData) SetGeneralFormCreateCnt1d(v string) *GetGeneralFormCreatedDeptSummaryResponseBodyData {
	s.GeneralFormCreateCnt1d = &v
	return s
}

func (s *GetGeneralFormCreatedDeptSummaryResponseBodyData) SetUseGeneralFormUserCnt1d(v string) *GetGeneralFormCreatedDeptSummaryResponseBodyData {
	s.UseGeneralFormUserCnt1d = &v
	return s
}

type GetGeneralFormCreatedDeptSummaryResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetGeneralFormCreatedDeptSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGeneralFormCreatedDeptSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGeneralFormCreatedDeptSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetGeneralFormCreatedDeptSummaryResponse) SetHeaders(v map[string]*string) *GetGeneralFormCreatedDeptSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetGeneralFormCreatedDeptSummaryResponse) SetStatusCode(v int32) *GetGeneralFormCreatedDeptSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGeneralFormCreatedDeptSummaryResponse) SetBody(v *GetGeneralFormCreatedDeptSummaryResponseBody) *GetGeneralFormCreatedDeptSummaryResponse {
	s.Body = v
	return s
}

type GetGeneralFormCreatedSummaryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetGeneralFormCreatedSummaryHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetGeneralFormCreatedSummaryHeaders) GoString() string {
	return s.String()
}

func (s *GetGeneralFormCreatedSummaryHeaders) SetCommonHeaders(v map[string]*string) *GetGeneralFormCreatedSummaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetGeneralFormCreatedSummaryHeaders) SetXAcsDingtalkAccessToken(v string) *GetGeneralFormCreatedSummaryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetGeneralFormCreatedSummaryResponseBody struct {
	GeneralFormCreatedCnt   *string `json:"generalFormCreatedCnt,omitempty" xml:"generalFormCreatedCnt,omitempty"`
	UseGeneralFormUserCnt1d *string `json:"useGeneralFormUserCnt1d,omitempty" xml:"useGeneralFormUserCnt1d,omitempty"`
}

func (s GetGeneralFormCreatedSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGeneralFormCreatedSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetGeneralFormCreatedSummaryResponseBody) SetGeneralFormCreatedCnt(v string) *GetGeneralFormCreatedSummaryResponseBody {
	s.GeneralFormCreatedCnt = &v
	return s
}

func (s *GetGeneralFormCreatedSummaryResponseBody) SetUseGeneralFormUserCnt1d(v string) *GetGeneralFormCreatedSummaryResponseBody {
	s.UseGeneralFormUserCnt1d = &v
	return s
}

type GetGeneralFormCreatedSummaryResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetGeneralFormCreatedSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGeneralFormCreatedSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGeneralFormCreatedSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetGeneralFormCreatedSummaryResponse) SetHeaders(v map[string]*string) *GetGeneralFormCreatedSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetGeneralFormCreatedSummaryResponse) SetStatusCode(v int32) *GetGeneralFormCreatedSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGeneralFormCreatedSummaryResponse) SetBody(v *GetGeneralFormCreatedSummaryResponseBody) *GetGeneralFormCreatedSummaryResponse {
	s.Body = v
	return s
}

type GetGroupActiveInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetGroupActiveInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetGroupActiveInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetGroupActiveInfoHeaders) SetCommonHeaders(v map[string]*string) *GetGroupActiveInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetGroupActiveInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetGroupActiveInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetGroupActiveInfoRequest struct {
	DingGroupId *string `json:"dingGroupId,omitempty" xml:"dingGroupId,omitempty"`
	GroupType   *int64  `json:"groupType,omitempty" xml:"groupType,omitempty"`
	PageNumber  *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize    *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StatDate    *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s GetGroupActiveInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGroupActiveInfoRequest) GoString() string {
	return s.String()
}

func (s *GetGroupActiveInfoRequest) SetDingGroupId(v string) *GetGroupActiveInfoRequest {
	s.DingGroupId = &v
	return s
}

func (s *GetGroupActiveInfoRequest) SetGroupType(v int64) *GetGroupActiveInfoRequest {
	s.GroupType = &v
	return s
}

func (s *GetGroupActiveInfoRequest) SetPageNumber(v int64) *GetGroupActiveInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *GetGroupActiveInfoRequest) SetPageSize(v int64) *GetGroupActiveInfoRequest {
	s.PageSize = &v
	return s
}

func (s *GetGroupActiveInfoRequest) SetStatDate(v string) *GetGroupActiveInfoRequest {
	s.StatDate = &v
	return s
}

type GetGroupActiveInfoResponseBody struct {
	Data       []*GetGroupActiveInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	TotalCount *int64                                `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetGroupActiveInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGroupActiveInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupActiveInfoResponseBody) SetData(v []*GetGroupActiveInfoResponseBodyData) *GetGroupActiveInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetGroupActiveInfoResponseBody) SetTotalCount(v int64) *GetGroupActiveInfoResponseBody {
	s.TotalCount = &v
	return s
}

type GetGroupActiveInfoResponseBodyData struct {
	DingGroupId          *string `json:"dingGroupId,omitempty" xml:"dingGroupId,omitempty"`
	GroupCreateTime      *string `json:"groupCreateTime,omitempty" xml:"groupCreateTime,omitempty"`
	GroupCreateUserId    *string `json:"groupCreateUserId,omitempty" xml:"groupCreateUserId,omitempty"`
	GroupCreateUserName  *string `json:"groupCreateUserName,omitempty" xml:"groupCreateUserName,omitempty"`
	GroupName            *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	GroupType            *int64  `json:"groupType,omitempty" xml:"groupType,omitempty"`
	GroupUserCnt1d       *int32  `json:"groupUserCnt1d,omitempty" xml:"groupUserCnt1d,omitempty"`
	OpenConvUv1d         *int32  `json:"openConvUv1d,omitempty" xml:"openConvUv1d,omitempty"`
	SendMessageCnt1d     *int64  `json:"sendMessageCnt1d,omitempty" xml:"sendMessageCnt1d,omitempty"`
	SendMessageUserCnt1d *int64  `json:"sendMessageUserCnt1d,omitempty" xml:"sendMessageUserCnt1d,omitempty"`
	StatDate             *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s GetGroupActiveInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetGroupActiveInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGroupActiveInfoResponseBodyData) SetDingGroupId(v string) *GetGroupActiveInfoResponseBodyData {
	s.DingGroupId = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetGroupCreateTime(v string) *GetGroupActiveInfoResponseBodyData {
	s.GroupCreateTime = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetGroupCreateUserId(v string) *GetGroupActiveInfoResponseBodyData {
	s.GroupCreateUserId = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetGroupCreateUserName(v string) *GetGroupActiveInfoResponseBodyData {
	s.GroupCreateUserName = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetGroupName(v string) *GetGroupActiveInfoResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetGroupType(v int64) *GetGroupActiveInfoResponseBodyData {
	s.GroupType = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetGroupUserCnt1d(v int32) *GetGroupActiveInfoResponseBodyData {
	s.GroupUserCnt1d = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetOpenConvUv1d(v int32) *GetGroupActiveInfoResponseBodyData {
	s.OpenConvUv1d = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetSendMessageCnt1d(v int64) *GetGroupActiveInfoResponseBodyData {
	s.SendMessageCnt1d = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetSendMessageUserCnt1d(v int64) *GetGroupActiveInfoResponseBodyData {
	s.SendMessageUserCnt1d = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetStatDate(v string) *GetGroupActiveInfoResponseBodyData {
	s.StatDate = &v
	return s
}

type GetGroupActiveInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetGroupActiveInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGroupActiveInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGroupActiveInfoResponse) GoString() string {
	return s.String()
}

func (s *GetGroupActiveInfoResponse) SetHeaders(v map[string]*string) *GetGroupActiveInfoResponse {
	s.Headers = v
	return s
}

func (s *GetGroupActiveInfoResponse) SetStatusCode(v int32) *GetGroupActiveInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGroupActiveInfoResponse) SetBody(v *GetGroupActiveInfoResponseBody) *GetGroupActiveInfoResponse {
	s.Body = v
	return s
}

type GetInActiveUserListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetInActiveUserListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetInActiveUserListHeaders) GoString() string {
	return s.String()
}

func (s *GetInActiveUserListHeaders) SetCommonHeaders(v map[string]*string) *GetInActiveUserListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInActiveUserListHeaders) SetXAcsDingtalkAccessToken(v string) *GetInActiveUserListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetInActiveUserListRequest struct {
	DeptIds    []*string `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	PageNumber *int64    `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StatDate   *string   `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s GetInActiveUserListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInActiveUserListRequest) GoString() string {
	return s.String()
}

func (s *GetInActiveUserListRequest) SetDeptIds(v []*string) *GetInActiveUserListRequest {
	s.DeptIds = v
	return s
}

func (s *GetInActiveUserListRequest) SetPageNumber(v int64) *GetInActiveUserListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetInActiveUserListRequest) SetPageSize(v int64) *GetInActiveUserListRequest {
	s.PageSize = &v
	return s
}

func (s *GetInActiveUserListRequest) SetStatDate(v string) *GetInActiveUserListRequest {
	s.StatDate = &v
	return s
}

type GetInActiveUserListResponseBody struct {
	DataList []map[string]interface{}                   `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	MetaList []*GetInActiveUserListResponseBodyMetaList `json:"metaList,omitempty" xml:"metaList,omitempty" type:"Repeated"`
}

func (s GetInActiveUserListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInActiveUserListResponseBody) GoString() string {
	return s.String()
}

func (s *GetInActiveUserListResponseBody) SetDataList(v []map[string]interface{}) *GetInActiveUserListResponseBody {
	s.DataList = v
	return s
}

func (s *GetInActiveUserListResponseBody) SetMetaList(v []*GetInActiveUserListResponseBodyMetaList) *GetInActiveUserListResponseBody {
	s.MetaList = v
	return s
}

type GetInActiveUserListResponseBodyMetaList struct {
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	KpiId      *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	KpiName    *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	Period     *string `json:"period,omitempty" xml:"period,omitempty"`
	Unit       *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s GetInActiveUserListResponseBodyMetaList) String() string {
	return tea.Prettify(s)
}

func (s GetInActiveUserListResponseBodyMetaList) GoString() string {
	return s.String()
}

func (s *GetInActiveUserListResponseBodyMetaList) SetKpiCaliber(v string) *GetInActiveUserListResponseBodyMetaList {
	s.KpiCaliber = &v
	return s
}

func (s *GetInActiveUserListResponseBodyMetaList) SetKpiId(v string) *GetInActiveUserListResponseBodyMetaList {
	s.KpiId = &v
	return s
}

func (s *GetInActiveUserListResponseBodyMetaList) SetKpiName(v string) *GetInActiveUserListResponseBodyMetaList {
	s.KpiName = &v
	return s
}

func (s *GetInActiveUserListResponseBodyMetaList) SetPeriod(v string) *GetInActiveUserListResponseBodyMetaList {
	s.Period = &v
	return s
}

func (s *GetInActiveUserListResponseBodyMetaList) SetUnit(v string) *GetInActiveUserListResponseBodyMetaList {
	s.Unit = &v
	return s
}

type GetInActiveUserListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInActiveUserListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInActiveUserListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInActiveUserListResponse) GoString() string {
	return s.String()
}

func (s *GetInActiveUserListResponse) SetHeaders(v map[string]*string) *GetInActiveUserListResponse {
	s.Headers = v
	return s
}

func (s *GetInActiveUserListResponse) SetStatusCode(v int32) *GetInActiveUserListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInActiveUserListResponse) SetBody(v *GetInActiveUserListResponseBody) *GetInActiveUserListResponse {
	s.Body = v
	return s
}

type GetLastOrgAuthDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetLastOrgAuthDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetLastOrgAuthDataHeaders) GoString() string {
	return s.String()
}

func (s *GetLastOrgAuthDataHeaders) SetCommonHeaders(v map[string]*string) *GetLastOrgAuthDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetLastOrgAuthDataHeaders) SetXAcsDingtalkAccessToken(v string) *GetLastOrgAuthDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetLastOrgAuthDataRequest struct {
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
}

func (s GetLastOrgAuthDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLastOrgAuthDataRequest) GoString() string {
	return s.String()
}

func (s *GetLastOrgAuthDataRequest) SetTargetCorpId(v string) *GetLastOrgAuthDataRequest {
	s.TargetCorpId = &v
	return s
}

type GetLastOrgAuthDataResponseBody struct {
	AuthRemark *string `json:"authRemark,omitempty" xml:"authRemark,omitempty"`
	AuthStatus *int32  `json:"authStatus,omitempty" xml:"authStatus,omitempty"`
}

func (s GetLastOrgAuthDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLastOrgAuthDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetLastOrgAuthDataResponseBody) SetAuthRemark(v string) *GetLastOrgAuthDataResponseBody {
	s.AuthRemark = &v
	return s
}

func (s *GetLastOrgAuthDataResponseBody) SetAuthStatus(v int32) *GetLastOrgAuthDataResponseBody {
	s.AuthStatus = &v
	return s
}

type GetLastOrgAuthDataResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLastOrgAuthDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLastOrgAuthDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLastOrgAuthDataResponse) GoString() string {
	return s.String()
}

func (s *GetLastOrgAuthDataResponse) SetHeaders(v map[string]*string) *GetLastOrgAuthDataResponse {
	s.Headers = v
	return s
}

func (s *GetLastOrgAuthDataResponse) SetStatusCode(v int32) *GetLastOrgAuthDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLastOrgAuthDataResponse) SetBody(v *GetLastOrgAuthDataResponseBody) *GetLastOrgAuthDataResponse {
	s.Body = v
	return s
}

type GetOaOperatorLogListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOaOperatorLogListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOaOperatorLogListHeaders) GoString() string {
	return s.String()
}

func (s *GetOaOperatorLogListHeaders) SetCommonHeaders(v map[string]*string) *GetOaOperatorLogListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOaOperatorLogListHeaders) SetXAcsDingtalkAccessToken(v string) *GetOaOperatorLogListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOaOperatorLogListRequest struct {
	CategoryList []*string `json:"categoryList,omitempty" xml:"categoryList,omitempty" type:"Repeated"`
	EndTime      *int64    `json:"endTime,omitempty" xml:"endTime,omitempty"`
	OpUserId     *string   `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	PageNumber   *int64    `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize     *int32    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StartTime    *int64    `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetOaOperatorLogListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOaOperatorLogListRequest) GoString() string {
	return s.String()
}

func (s *GetOaOperatorLogListRequest) SetCategoryList(v []*string) *GetOaOperatorLogListRequest {
	s.CategoryList = v
	return s
}

func (s *GetOaOperatorLogListRequest) SetEndTime(v int64) *GetOaOperatorLogListRequest {
	s.EndTime = &v
	return s
}

func (s *GetOaOperatorLogListRequest) SetOpUserId(v string) *GetOaOperatorLogListRequest {
	s.OpUserId = &v
	return s
}

func (s *GetOaOperatorLogListRequest) SetPageNumber(v int64) *GetOaOperatorLogListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetOaOperatorLogListRequest) SetPageSize(v int32) *GetOaOperatorLogListRequest {
	s.PageSize = &v
	return s
}

func (s *GetOaOperatorLogListRequest) SetStartTime(v int64) *GetOaOperatorLogListRequest {
	s.StartTime = &v
	return s
}

type GetOaOperatorLogListResponseBody struct {
	Data      []*GetOaOperatorLogListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	ItemCount *int64                                  `json:"itemCount,omitempty" xml:"itemCount,omitempty"`
}

func (s GetOaOperatorLogListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOaOperatorLogListResponseBody) GoString() string {
	return s.String()
}

func (s *GetOaOperatorLogListResponseBody) SetData(v []*GetOaOperatorLogListResponseBodyData) *GetOaOperatorLogListResponseBody {
	s.Data = v
	return s
}

func (s *GetOaOperatorLogListResponseBody) SetItemCount(v int64) *GetOaOperatorLogListResponseBody {
	s.ItemCount = &v
	return s
}

type GetOaOperatorLogListResponseBodyData struct {
	Category1Name *string `json:"category1Name,omitempty" xml:"category1Name,omitempty"`
	Category2Name *string `json:"category2Name,omitempty" xml:"category2Name,omitempty"`
	Content       *string `json:"content,omitempty" xml:"content,omitempty"`
	OpName        *string `json:"opName,omitempty" xml:"opName,omitempty"`
	OpTime        *int64  `json:"opTime,omitempty" xml:"opTime,omitempty"`
	OpUserId      *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
}

func (s GetOaOperatorLogListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOaOperatorLogListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOaOperatorLogListResponseBodyData) SetCategory1Name(v string) *GetOaOperatorLogListResponseBodyData {
	s.Category1Name = &v
	return s
}

func (s *GetOaOperatorLogListResponseBodyData) SetCategory2Name(v string) *GetOaOperatorLogListResponseBodyData {
	s.Category2Name = &v
	return s
}

func (s *GetOaOperatorLogListResponseBodyData) SetContent(v string) *GetOaOperatorLogListResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetOaOperatorLogListResponseBodyData) SetOpName(v string) *GetOaOperatorLogListResponseBodyData {
	s.OpName = &v
	return s
}

func (s *GetOaOperatorLogListResponseBodyData) SetOpTime(v int64) *GetOaOperatorLogListResponseBodyData {
	s.OpTime = &v
	return s
}

func (s *GetOaOperatorLogListResponseBodyData) SetOpUserId(v string) *GetOaOperatorLogListResponseBodyData {
	s.OpUserId = &v
	return s
}

type GetOaOperatorLogListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOaOperatorLogListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOaOperatorLogListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOaOperatorLogListResponse) GoString() string {
	return s.String()
}

func (s *GetOaOperatorLogListResponse) SetHeaders(v map[string]*string) *GetOaOperatorLogListResponse {
	s.Headers = v
	return s
}

func (s *GetOaOperatorLogListResponse) SetStatusCode(v int32) *GetOaOperatorLogListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOaOperatorLogListResponse) SetBody(v *GetOaOperatorLogListResponseBody) *GetOaOperatorLogListResponse {
	s.Body = v
	return s
}

type GetOutGroupsByPageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOutGroupsByPageHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOutGroupsByPageHeaders) GoString() string {
	return s.String()
}

func (s *GetOutGroupsByPageHeaders) SetCommonHeaders(v map[string]*string) *GetOutGroupsByPageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOutGroupsByPageHeaders) SetXAcsDingtalkAccessToken(v string) *GetOutGroupsByPageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOutGroupsByPageRequest struct {
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetOutGroupsByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOutGroupsByPageRequest) GoString() string {
	return s.String()
}

func (s *GetOutGroupsByPageRequest) SetPageNumber(v int32) *GetOutGroupsByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *GetOutGroupsByPageRequest) SetPageSize(v int32) *GetOutGroupsByPageRequest {
	s.PageSize = &v
	return s
}

type GetOutGroupsByPageResponseBody struct {
	ResponseBody *GetOutGroupsByPageResponseBodyResponseBody `json:"responseBody,omitempty" xml:"responseBody,omitempty" type:"Struct"`
}

func (s GetOutGroupsByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOutGroupsByPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetOutGroupsByPageResponseBody) SetResponseBody(v *GetOutGroupsByPageResponseBodyResponseBody) *GetOutGroupsByPageResponseBody {
	s.ResponseBody = v
	return s
}

type GetOutGroupsByPageResponseBodyResponseBody struct {
	GroupList []*GetOutGroupsByPageResponseBodyResponseBodyGroupList `json:"groupList,omitempty" xml:"groupList,omitempty" type:"Repeated"`
	Total     *int32                                                 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetOutGroupsByPageResponseBodyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOutGroupsByPageResponseBodyResponseBody) GoString() string {
	return s.String()
}

func (s *GetOutGroupsByPageResponseBodyResponseBody) SetGroupList(v []*GetOutGroupsByPageResponseBodyResponseBodyGroupList) *GetOutGroupsByPageResponseBodyResponseBody {
	s.GroupList = v
	return s
}

func (s *GetOutGroupsByPageResponseBodyResponseBody) SetTotal(v int32) *GetOutGroupsByPageResponseBodyResponseBody {
	s.Total = &v
	return s
}

type GetOutGroupsByPageResponseBodyResponseBodyGroupList struct {
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s GetOutGroupsByPageResponseBodyResponseBodyGroupList) String() string {
	return tea.Prettify(s)
}

func (s GetOutGroupsByPageResponseBodyResponseBodyGroupList) GoString() string {
	return s.String()
}

func (s *GetOutGroupsByPageResponseBodyResponseBodyGroupList) SetOpenConversationId(v string) *GetOutGroupsByPageResponseBodyResponseBodyGroupList {
	s.OpenConversationId = &v
	return s
}

type GetOutGroupsByPageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOutGroupsByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOutGroupsByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOutGroupsByPageResponse) GoString() string {
	return s.String()
}

func (s *GetOutGroupsByPageResponse) SetHeaders(v map[string]*string) *GetOutGroupsByPageResponse {
	s.Headers = v
	return s
}

func (s *GetOutGroupsByPageResponse) SetStatusCode(v int32) *GetOutGroupsByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOutGroupsByPageResponse) SetBody(v *GetOutGroupsByPageResponseBody) *GetOutGroupsByPageResponse {
	s.Body = v
	return s
}

type GetOutsideAuditGroupMessageByPageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOutsideAuditGroupMessageByPageHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOutsideAuditGroupMessageByPageHeaders) GoString() string {
	return s.String()
}

func (s *GetOutsideAuditGroupMessageByPageHeaders) SetCommonHeaders(v map[string]*string) *GetOutsideAuditGroupMessageByPageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOutsideAuditGroupMessageByPageHeaders) SetXAcsDingtalkAccessToken(v string) *GetOutsideAuditGroupMessageByPageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOutsideAuditGroupMessageByPageRequest struct {
	MaxResults         *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken          *int64  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s GetOutsideAuditGroupMessageByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOutsideAuditGroupMessageByPageRequest) GoString() string {
	return s.String()
}

func (s *GetOutsideAuditGroupMessageByPageRequest) SetMaxResults(v int32) *GetOutsideAuditGroupMessageByPageRequest {
	s.MaxResults = &v
	return s
}

func (s *GetOutsideAuditGroupMessageByPageRequest) SetNextToken(v int64) *GetOutsideAuditGroupMessageByPageRequest {
	s.NextToken = &v
	return s
}

func (s *GetOutsideAuditGroupMessageByPageRequest) SetOpenConversationId(v string) *GetOutsideAuditGroupMessageByPageRequest {
	s.OpenConversationId = &v
	return s
}

type GetOutsideAuditGroupMessageByPageResponseBody struct {
	ResponseBody *GetOutsideAuditGroupMessageByPageResponseBodyResponseBody `json:"responseBody,omitempty" xml:"responseBody,omitempty" type:"Struct"`
}

func (s GetOutsideAuditGroupMessageByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOutsideAuditGroupMessageByPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetOutsideAuditGroupMessageByPageResponseBody) SetResponseBody(v *GetOutsideAuditGroupMessageByPageResponseBodyResponseBody) *GetOutsideAuditGroupMessageByPageResponseBody {
	s.ResponseBody = v
	return s
}

type GetOutsideAuditGroupMessageByPageResponseBodyResponseBody struct {
	MessageList []*GetOutsideAuditGroupMessageByPageResponseBodyResponseBodyMessageList `json:"messageList,omitempty" xml:"messageList,omitempty" type:"Repeated"`
	NextToken   *string                                                                 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetOutsideAuditGroupMessageByPageResponseBodyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOutsideAuditGroupMessageByPageResponseBodyResponseBody) GoString() string {
	return s.String()
}

func (s *GetOutsideAuditGroupMessageByPageResponseBodyResponseBody) SetMessageList(v []*GetOutsideAuditGroupMessageByPageResponseBodyResponseBodyMessageList) *GetOutsideAuditGroupMessageByPageResponseBodyResponseBody {
	s.MessageList = v
	return s
}

func (s *GetOutsideAuditGroupMessageByPageResponseBodyResponseBody) SetNextToken(v string) *GetOutsideAuditGroupMessageByPageResponseBodyResponseBody {
	s.NextToken = &v
	return s
}

type GetOutsideAuditGroupMessageByPageResponseBodyResponseBodyMessageList struct {
	Content            *string                                                                     `json:"content,omitempty" xml:"content,omitempty"`
	ContentType        *string                                                                     `json:"contentType,omitempty" xml:"contentType,omitempty"`
	CreateAt           *string                                                                     `json:"createAt,omitempty" xml:"createAt,omitempty"`
	OpenConversationId *string                                                                     `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	Sender             *GetOutsideAuditGroupMessageByPageResponseBodyResponseBodyMessageListSender `json:"sender,omitempty" xml:"sender,omitempty" type:"Struct"`
}

func (s GetOutsideAuditGroupMessageByPageResponseBodyResponseBodyMessageList) String() string {
	return tea.Prettify(s)
}

func (s GetOutsideAuditGroupMessageByPageResponseBodyResponseBodyMessageList) GoString() string {
	return s.String()
}

func (s *GetOutsideAuditGroupMessageByPageResponseBodyResponseBodyMessageList) SetContent(v string) *GetOutsideAuditGroupMessageByPageResponseBodyResponseBodyMessageList {
	s.Content = &v
	return s
}

func (s *GetOutsideAuditGroupMessageByPageResponseBodyResponseBodyMessageList) SetContentType(v string) *GetOutsideAuditGroupMessageByPageResponseBodyResponseBodyMessageList {
	s.ContentType = &v
	return s
}

func (s *GetOutsideAuditGroupMessageByPageResponseBodyResponseBodyMessageList) SetCreateAt(v string) *GetOutsideAuditGroupMessageByPageResponseBodyResponseBodyMessageList {
	s.CreateAt = &v
	return s
}

func (s *GetOutsideAuditGroupMessageByPageResponseBodyResponseBodyMessageList) SetOpenConversationId(v string) *GetOutsideAuditGroupMessageByPageResponseBodyResponseBodyMessageList {
	s.OpenConversationId = &v
	return s
}

func (s *GetOutsideAuditGroupMessageByPageResponseBodyResponseBodyMessageList) SetSender(v *GetOutsideAuditGroupMessageByPageResponseBodyResponseBodyMessageListSender) *GetOutsideAuditGroupMessageByPageResponseBodyResponseBodyMessageList {
	s.Sender = v
	return s
}

type GetOutsideAuditGroupMessageByPageResponseBodyResponseBodyMessageListSender struct {
	Id     *string `json:"id,omitempty" xml:"id,omitempty"`
	IdType *string `json:"idType,omitempty" xml:"idType,omitempty"`
	Type   *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetOutsideAuditGroupMessageByPageResponseBodyResponseBodyMessageListSender) String() string {
	return tea.Prettify(s)
}

func (s GetOutsideAuditGroupMessageByPageResponseBodyResponseBodyMessageListSender) GoString() string {
	return s.String()
}

func (s *GetOutsideAuditGroupMessageByPageResponseBodyResponseBodyMessageListSender) SetId(v string) *GetOutsideAuditGroupMessageByPageResponseBodyResponseBodyMessageListSender {
	s.Id = &v
	return s
}

func (s *GetOutsideAuditGroupMessageByPageResponseBodyResponseBodyMessageListSender) SetIdType(v string) *GetOutsideAuditGroupMessageByPageResponseBodyResponseBodyMessageListSender {
	s.IdType = &v
	return s
}

func (s *GetOutsideAuditGroupMessageByPageResponseBodyResponseBodyMessageListSender) SetType(v string) *GetOutsideAuditGroupMessageByPageResponseBodyResponseBodyMessageListSender {
	s.Type = &v
	return s
}

type GetOutsideAuditGroupMessageByPageResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOutsideAuditGroupMessageByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOutsideAuditGroupMessageByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOutsideAuditGroupMessageByPageResponse) GoString() string {
	return s.String()
}

func (s *GetOutsideAuditGroupMessageByPageResponse) SetHeaders(v map[string]*string) *GetOutsideAuditGroupMessageByPageResponse {
	s.Headers = v
	return s
}

func (s *GetOutsideAuditGroupMessageByPageResponse) SetStatusCode(v int32) *GetOutsideAuditGroupMessageByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOutsideAuditGroupMessageByPageResponse) SetBody(v *GetOutsideAuditGroupMessageByPageResponseBody) *GetOutsideAuditGroupMessageByPageResponse {
	s.Body = v
	return s
}

type GetPartnerTypeByParentIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPartnerTypeByParentIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPartnerTypeByParentIdHeaders) GoString() string {
	return s.String()
}

func (s *GetPartnerTypeByParentIdHeaders) SetCommonHeaders(v map[string]*string) *GetPartnerTypeByParentIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPartnerTypeByParentIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetPartnerTypeByParentIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPartnerTypeByParentIdResponseBody struct {
	Data []*GetPartnerTypeByParentIdResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s GetPartnerTypeByParentIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPartnerTypeByParentIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetPartnerTypeByParentIdResponseBody) SetData(v []*GetPartnerTypeByParentIdResponseBodyData) *GetPartnerTypeByParentIdResponseBody {
	s.Data = v
	return s
}

type GetPartnerTypeByParentIdResponseBodyData struct {
	LabelId  *string  `json:"labelId,omitempty" xml:"labelId,omitempty"`
	TypeId   *float32 `json:"typeId,omitempty" xml:"typeId,omitempty"`
	TypeName *string  `json:"typeName,omitempty" xml:"typeName,omitempty"`
}

func (s GetPartnerTypeByParentIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPartnerTypeByParentIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPartnerTypeByParentIdResponseBodyData) SetLabelId(v string) *GetPartnerTypeByParentIdResponseBodyData {
	s.LabelId = &v
	return s
}

func (s *GetPartnerTypeByParentIdResponseBodyData) SetTypeId(v float32) *GetPartnerTypeByParentIdResponseBodyData {
	s.TypeId = &v
	return s
}

func (s *GetPartnerTypeByParentIdResponseBodyData) SetTypeName(v string) *GetPartnerTypeByParentIdResponseBodyData {
	s.TypeName = &v
	return s
}

type GetPartnerTypeByParentIdResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPartnerTypeByParentIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPartnerTypeByParentIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPartnerTypeByParentIdResponse) GoString() string {
	return s.String()
}

func (s *GetPartnerTypeByParentIdResponse) SetHeaders(v map[string]*string) *GetPartnerTypeByParentIdResponse {
	s.Headers = v
	return s
}

func (s *GetPartnerTypeByParentIdResponse) SetStatusCode(v int32) *GetPartnerTypeByParentIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPartnerTypeByParentIdResponse) SetBody(v *GetPartnerTypeByParentIdResponseBody) *GetPartnerTypeByParentIdResponse {
	s.Body = v
	return s
}

type GetPublicDevicesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPublicDevicesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPublicDevicesHeaders) GoString() string {
	return s.String()
}

func (s *GetPublicDevicesHeaders) SetCommonHeaders(v map[string]*string) *GetPublicDevicesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPublicDevicesHeaders) SetXAcsDingtalkAccessToken(v string) *GetPublicDevicesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPublicDevicesRequest struct {
	EndTime    *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	MacAddress *string `json:"macAddress,omitempty" xml:"macAddress,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Platform   *string `json:"platform,omitempty" xml:"platform,omitempty"`
	StartTime  *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Title      *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetPublicDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPublicDevicesRequest) GoString() string {
	return s.String()
}

func (s *GetPublicDevicesRequest) SetEndTime(v int64) *GetPublicDevicesRequest {
	s.EndTime = &v
	return s
}

func (s *GetPublicDevicesRequest) SetMacAddress(v string) *GetPublicDevicesRequest {
	s.MacAddress = &v
	return s
}

func (s *GetPublicDevicesRequest) SetPageNumber(v int32) *GetPublicDevicesRequest {
	s.PageNumber = &v
	return s
}

func (s *GetPublicDevicesRequest) SetPageSize(v int32) *GetPublicDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *GetPublicDevicesRequest) SetPlatform(v string) *GetPublicDevicesRequest {
	s.Platform = &v
	return s
}

func (s *GetPublicDevicesRequest) SetStartTime(v int64) *GetPublicDevicesRequest {
	s.StartTime = &v
	return s
}

func (s *GetPublicDevicesRequest) SetTitle(v string) *GetPublicDevicesRequest {
	s.Title = &v
	return s
}

type GetPublicDevicesResponseBody struct {
	Data     []*GetPublicDevicesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	DataCnt  *int32                              `json:"dataCnt,omitempty" xml:"dataCnt,omitempty"`
	TotalCnt *int64                              `json:"totalCnt,omitempty" xml:"totalCnt,omitempty"`
}

func (s GetPublicDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPublicDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *GetPublicDevicesResponseBody) SetData(v []*GetPublicDevicesResponseBodyData) *GetPublicDevicesResponseBody {
	s.Data = v
	return s
}

func (s *GetPublicDevicesResponseBody) SetDataCnt(v int32) *GetPublicDevicesResponseBody {
	s.DataCnt = &v
	return s
}

func (s *GetPublicDevicesResponseBody) SetTotalCnt(v int64) *GetPublicDevicesResponseBody {
	s.TotalCnt = &v
	return s
}

type GetPublicDevicesResponseBodyData struct {
	DeviceDepts     []*GetPublicDevicesResponseBodyDataDeviceDepts  `json:"deviceDepts,omitempty" xml:"deviceDepts,omitempty" type:"Repeated"`
	DeviceRoles     []*GetPublicDevicesResponseBodyDataDeviceRoles  `json:"deviceRoles,omitempty" xml:"deviceRoles,omitempty" type:"Repeated"`
	DeviceScopeType *int32                                          `json:"deviceScopeType,omitempty" xml:"deviceScopeType,omitempty"`
	DeviceStaffs    []*GetPublicDevicesResponseBodyDataDeviceStaffs `json:"deviceStaffs,omitempty" xml:"deviceStaffs,omitempty" type:"Repeated"`
	GmtCreate       *int64                                          `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified     *int64                                          `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	MacAddress      *string                                         `json:"macAddress,omitempty" xml:"macAddress,omitempty"`
	Platform        *string                                         `json:"platform,omitempty" xml:"platform,omitempty"`
	Title           *string                                         `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetPublicDevicesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPublicDevicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPublicDevicesResponseBodyData) SetDeviceDepts(v []*GetPublicDevicesResponseBodyDataDeviceDepts) *GetPublicDevicesResponseBodyData {
	s.DeviceDepts = v
	return s
}

func (s *GetPublicDevicesResponseBodyData) SetDeviceRoles(v []*GetPublicDevicesResponseBodyDataDeviceRoles) *GetPublicDevicesResponseBodyData {
	s.DeviceRoles = v
	return s
}

func (s *GetPublicDevicesResponseBodyData) SetDeviceScopeType(v int32) *GetPublicDevicesResponseBodyData {
	s.DeviceScopeType = &v
	return s
}

func (s *GetPublicDevicesResponseBodyData) SetDeviceStaffs(v []*GetPublicDevicesResponseBodyDataDeviceStaffs) *GetPublicDevicesResponseBodyData {
	s.DeviceStaffs = v
	return s
}

func (s *GetPublicDevicesResponseBodyData) SetGmtCreate(v int64) *GetPublicDevicesResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetPublicDevicesResponseBodyData) SetGmtModified(v int64) *GetPublicDevicesResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetPublicDevicesResponseBodyData) SetMacAddress(v string) *GetPublicDevicesResponseBodyData {
	s.MacAddress = &v
	return s
}

func (s *GetPublicDevicesResponseBodyData) SetPlatform(v string) *GetPublicDevicesResponseBodyData {
	s.Platform = &v
	return s
}

func (s *GetPublicDevicesResponseBodyData) SetTitle(v string) *GetPublicDevicesResponseBodyData {
	s.Title = &v
	return s
}

type GetPublicDevicesResponseBodyDataDeviceDepts struct {
	Id   *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetPublicDevicesResponseBodyDataDeviceDepts) String() string {
	return tea.Prettify(s)
}

func (s GetPublicDevicesResponseBodyDataDeviceDepts) GoString() string {
	return s.String()
}

func (s *GetPublicDevicesResponseBodyDataDeviceDepts) SetId(v int64) *GetPublicDevicesResponseBodyDataDeviceDepts {
	s.Id = &v
	return s
}

func (s *GetPublicDevicesResponseBodyDataDeviceDepts) SetName(v string) *GetPublicDevicesResponseBodyDataDeviceDepts {
	s.Name = &v
	return s
}

type GetPublicDevicesResponseBodyDataDeviceRoles struct {
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	TagCode *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
}

func (s GetPublicDevicesResponseBodyDataDeviceRoles) String() string {
	return tea.Prettify(s)
}

func (s GetPublicDevicesResponseBodyDataDeviceRoles) GoString() string {
	return s.String()
}

func (s *GetPublicDevicesResponseBodyDataDeviceRoles) SetName(v string) *GetPublicDevicesResponseBodyDataDeviceRoles {
	s.Name = &v
	return s
}

func (s *GetPublicDevicesResponseBodyDataDeviceRoles) SetTagCode(v string) *GetPublicDevicesResponseBodyDataDeviceRoles {
	s.TagCode = &v
	return s
}

type GetPublicDevicesResponseBodyDataDeviceStaffs struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetPublicDevicesResponseBodyDataDeviceStaffs) String() string {
	return tea.Prettify(s)
}

func (s GetPublicDevicesResponseBodyDataDeviceStaffs) GoString() string {
	return s.String()
}

func (s *GetPublicDevicesResponseBodyDataDeviceStaffs) SetName(v string) *GetPublicDevicesResponseBodyDataDeviceStaffs {
	s.Name = &v
	return s
}

func (s *GetPublicDevicesResponseBodyDataDeviceStaffs) SetUserId(v string) *GetPublicDevicesResponseBodyDataDeviceStaffs {
	s.UserId = &v
	return s
}

type GetPublicDevicesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPublicDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPublicDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPublicDevicesResponse) GoString() string {
	return s.String()
}

func (s *GetPublicDevicesResponse) SetHeaders(v map[string]*string) *GetPublicDevicesResponse {
	s.Headers = v
	return s
}

func (s *GetPublicDevicesResponse) SetStatusCode(v int32) *GetPublicDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPublicDevicesResponse) SetBody(v *GetPublicDevicesResponseBody) *GetPublicDevicesResponse {
	s.Body = v
	return s
}

type GetPublisherSummaryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPublisherSummaryHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPublisherSummaryHeaders) GoString() string {
	return s.String()
}

func (s *GetPublisherSummaryHeaders) SetCommonHeaders(v map[string]*string) *GetPublisherSummaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPublisherSummaryHeaders) SetXAcsDingtalkAccessToken(v string) *GetPublisherSummaryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPublisherSummaryRequest struct {
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetPublisherSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPublisherSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetPublisherSummaryRequest) SetMaxResults(v int64) *GetPublisherSummaryRequest {
	s.MaxResults = &v
	return s
}

func (s *GetPublisherSummaryRequest) SetNextToken(v int64) *GetPublisherSummaryRequest {
	s.NextToken = &v
	return s
}

type GetPublisherSummaryResponseBody struct {
	Data                     []*GetPublisherSummaryResponseBodyData                   `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	HasMore                  *bool                                                    `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken                *int64                                                   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	PublisherArticleCntStd   *string                                                  `json:"publisherArticleCntStd,omitempty" xml:"publisherArticleCntStd,omitempty"`
	PublisherArticlePvCntStd *string                                                  `json:"publisherArticlePvCntStd,omitempty" xml:"publisherArticlePvCntStd,omitempty"`
	PublisherArticlePvTop5   []*GetPublisherSummaryResponseBodyPublisherArticlePvTop5 `json:"publisherArticlePvTop5,omitempty" xml:"publisherArticlePvTop5,omitempty" type:"Repeated"`
	PublisherCntStd          *string                                                  `json:"publisherCntStd,omitempty" xml:"publisherCntStd,omitempty"`
}

func (s GetPublisherSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPublisherSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetPublisherSummaryResponseBody) SetData(v []*GetPublisherSummaryResponseBodyData) *GetPublisherSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetPublisherSummaryResponseBody) SetHasMore(v bool) *GetPublisherSummaryResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetPublisherSummaryResponseBody) SetNextToken(v int64) *GetPublisherSummaryResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetPublisherSummaryResponseBody) SetPublisherArticleCntStd(v string) *GetPublisherSummaryResponseBody {
	s.PublisherArticleCntStd = &v
	return s
}

func (s *GetPublisherSummaryResponseBody) SetPublisherArticlePvCntStd(v string) *GetPublisherSummaryResponseBody {
	s.PublisherArticlePvCntStd = &v
	return s
}

func (s *GetPublisherSummaryResponseBody) SetPublisherArticlePvTop5(v []*GetPublisherSummaryResponseBodyPublisherArticlePvTop5) *GetPublisherSummaryResponseBody {
	s.PublisherArticlePvTop5 = v
	return s
}

func (s *GetPublisherSummaryResponseBody) SetPublisherCntStd(v string) *GetPublisherSummaryResponseBody {
	s.PublisherCntStd = &v
	return s
}

type GetPublisherSummaryResponseBodyData struct {
	PublisherArticleCntStd   *string `json:"publisherArticleCntStd,omitempty" xml:"publisherArticleCntStd,omitempty"`
	PublisherArticlePvCntStd *string `json:"publisherArticlePvCntStd,omitempty" xml:"publisherArticlePvCntStd,omitempty"`
	PublisherName            *string `json:"publisherName,omitempty" xml:"publisherName,omitempty"`
	UnionId                  *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetPublisherSummaryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPublisherSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPublisherSummaryResponseBodyData) SetPublisherArticleCntStd(v string) *GetPublisherSummaryResponseBodyData {
	s.PublisherArticleCntStd = &v
	return s
}

func (s *GetPublisherSummaryResponseBodyData) SetPublisherArticlePvCntStd(v string) *GetPublisherSummaryResponseBodyData {
	s.PublisherArticlePvCntStd = &v
	return s
}

func (s *GetPublisherSummaryResponseBodyData) SetPublisherName(v string) *GetPublisherSummaryResponseBodyData {
	s.PublisherName = &v
	return s
}

func (s *GetPublisherSummaryResponseBodyData) SetUnionId(v string) *GetPublisherSummaryResponseBodyData {
	s.UnionId = &v
	return s
}

type GetPublisherSummaryResponseBodyPublisherArticlePvTop5 struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetPublisherSummaryResponseBodyPublisherArticlePvTop5) String() string {
	return tea.Prettify(s)
}

func (s GetPublisherSummaryResponseBodyPublisherArticlePvTop5) GoString() string {
	return s.String()
}

func (s *GetPublisherSummaryResponseBodyPublisherArticlePvTop5) SetName(v string) *GetPublisherSummaryResponseBodyPublisherArticlePvTop5 {
	s.Name = &v
	return s
}

type GetPublisherSummaryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPublisherSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPublisherSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPublisherSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetPublisherSummaryResponse) SetHeaders(v map[string]*string) *GetPublisherSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetPublisherSummaryResponse) SetStatusCode(v int32) *GetPublisherSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPublisherSummaryResponse) SetBody(v *GetPublisherSummaryResponseBody) *GetPublisherSummaryResponse {
	s.Body = v
	return s
}

type GetRealPeopleRecordsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetRealPeopleRecordsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetRealPeopleRecordsHeaders) GoString() string {
	return s.String()
}

func (s *GetRealPeopleRecordsHeaders) SetCommonHeaders(v map[string]*string) *GetRealPeopleRecordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRealPeopleRecordsHeaders) SetXAcsDingtalkAccessToken(v string) *GetRealPeopleRecordsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetRealPeopleRecordsRequest struct {
	AgentId              *int64    `json:"agentId,omitempty" xml:"agentId,omitempty"`
	FromTime             *int64    `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	MaxResults           *int32    `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken            *int64    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	PersonIdentification *int32    `json:"personIdentification,omitempty" xml:"personIdentification,omitempty"`
	Scene                *int32    `json:"scene,omitempty" xml:"scene,omitempty"`
	ToTime               *int64    `json:"toTime,omitempty" xml:"toTime,omitempty"`
	UserIds              []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s GetRealPeopleRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRealPeopleRecordsRequest) GoString() string {
	return s.String()
}

func (s *GetRealPeopleRecordsRequest) SetAgentId(v int64) *GetRealPeopleRecordsRequest {
	s.AgentId = &v
	return s
}

func (s *GetRealPeopleRecordsRequest) SetFromTime(v int64) *GetRealPeopleRecordsRequest {
	s.FromTime = &v
	return s
}

func (s *GetRealPeopleRecordsRequest) SetMaxResults(v int32) *GetRealPeopleRecordsRequest {
	s.MaxResults = &v
	return s
}

func (s *GetRealPeopleRecordsRequest) SetNextToken(v int64) *GetRealPeopleRecordsRequest {
	s.NextToken = &v
	return s
}

func (s *GetRealPeopleRecordsRequest) SetPersonIdentification(v int32) *GetRealPeopleRecordsRequest {
	s.PersonIdentification = &v
	return s
}

func (s *GetRealPeopleRecordsRequest) SetScene(v int32) *GetRealPeopleRecordsRequest {
	s.Scene = &v
	return s
}

func (s *GetRealPeopleRecordsRequest) SetToTime(v int64) *GetRealPeopleRecordsRequest {
	s.ToTime = &v
	return s
}

func (s *GetRealPeopleRecordsRequest) SetUserIds(v []*string) *GetRealPeopleRecordsRequest {
	s.UserIds = v
	return s
}

type GetRealPeopleRecordsResponseBody struct {
	Data      []*GetRealPeopleRecordsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	NextToken *int64                                  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Total     *int32                                  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetRealPeopleRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRealPeopleRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *GetRealPeopleRecordsResponseBody) SetData(v []*GetRealPeopleRecordsResponseBodyData) *GetRealPeopleRecordsResponseBody {
	s.Data = v
	return s
}

func (s *GetRealPeopleRecordsResponseBody) SetNextToken(v int64) *GetRealPeopleRecordsResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetRealPeopleRecordsResponseBody) SetTotal(v int32) *GetRealPeopleRecordsResponseBody {
	s.Total = &v
	return s
}

type GetRealPeopleRecordsResponseBodyData struct {
	AgentId              *int64  `json:"agentId,omitempty" xml:"agentId,omitempty"`
	InvokeTime           *int64  `json:"invokeTime,omitempty" xml:"invokeTime,omitempty"`
	PersonIdentification *int32  `json:"personIdentification,omitempty" xml:"personIdentification,omitempty"`
	Platform             *int32  `json:"platform,omitempty" xml:"platform,omitempty"`
	Scene                *int32  `json:"scene,omitempty" xml:"scene,omitempty"`
	UserId               *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetRealPeopleRecordsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRealPeopleRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRealPeopleRecordsResponseBodyData) SetAgentId(v int64) *GetRealPeopleRecordsResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *GetRealPeopleRecordsResponseBodyData) SetInvokeTime(v int64) *GetRealPeopleRecordsResponseBodyData {
	s.InvokeTime = &v
	return s
}

func (s *GetRealPeopleRecordsResponseBodyData) SetPersonIdentification(v int32) *GetRealPeopleRecordsResponseBodyData {
	s.PersonIdentification = &v
	return s
}

func (s *GetRealPeopleRecordsResponseBodyData) SetPlatform(v int32) *GetRealPeopleRecordsResponseBodyData {
	s.Platform = &v
	return s
}

func (s *GetRealPeopleRecordsResponseBodyData) SetScene(v int32) *GetRealPeopleRecordsResponseBodyData {
	s.Scene = &v
	return s
}

func (s *GetRealPeopleRecordsResponseBodyData) SetUserId(v string) *GetRealPeopleRecordsResponseBodyData {
	s.UserId = &v
	return s
}

type GetRealPeopleRecordsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRealPeopleRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRealPeopleRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRealPeopleRecordsResponse) GoString() string {
	return s.String()
}

func (s *GetRealPeopleRecordsResponse) SetHeaders(v map[string]*string) *GetRealPeopleRecordsResponse {
	s.Headers = v
	return s
}

func (s *GetRealPeopleRecordsResponse) SetStatusCode(v int32) *GetRealPeopleRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRealPeopleRecordsResponse) SetBody(v *GetRealPeopleRecordsResponseBody) *GetRealPeopleRecordsResponse {
	s.Body = v
	return s
}

type GetRecognizeRecordsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetRecognizeRecordsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetRecognizeRecordsHeaders) GoString() string {
	return s.String()
}

func (s *GetRecognizeRecordsHeaders) SetCommonHeaders(v map[string]*string) *GetRecognizeRecordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRecognizeRecordsHeaders) SetXAcsDingtalkAccessToken(v string) *GetRecognizeRecordsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetRecognizeRecordsRequest struct {
	AgentId           *int64    `json:"agentId,omitempty" xml:"agentId,omitempty"`
	FaceCompareResult *int32    `json:"faceCompareResult,omitempty" xml:"faceCompareResult,omitempty"`
	FromTime          *int64    `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	MaxResults        *int32    `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken         *int64    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ToTime            *int64    `json:"toTime,omitempty" xml:"toTime,omitempty"`
	UserIds           []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s GetRecognizeRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRecognizeRecordsRequest) GoString() string {
	return s.String()
}

func (s *GetRecognizeRecordsRequest) SetAgentId(v int64) *GetRecognizeRecordsRequest {
	s.AgentId = &v
	return s
}

func (s *GetRecognizeRecordsRequest) SetFaceCompareResult(v int32) *GetRecognizeRecordsRequest {
	s.FaceCompareResult = &v
	return s
}

func (s *GetRecognizeRecordsRequest) SetFromTime(v int64) *GetRecognizeRecordsRequest {
	s.FromTime = &v
	return s
}

func (s *GetRecognizeRecordsRequest) SetMaxResults(v int32) *GetRecognizeRecordsRequest {
	s.MaxResults = &v
	return s
}

func (s *GetRecognizeRecordsRequest) SetNextToken(v int64) *GetRecognizeRecordsRequest {
	s.NextToken = &v
	return s
}

func (s *GetRecognizeRecordsRequest) SetToTime(v int64) *GetRecognizeRecordsRequest {
	s.ToTime = &v
	return s
}

func (s *GetRecognizeRecordsRequest) SetUserIds(v []*string) *GetRecognizeRecordsRequest {
	s.UserIds = v
	return s
}

type GetRecognizeRecordsResponseBody struct {
	Data      []*GetRecognizeRecordsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	NextToken *int64                                 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Total     *int32                                 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetRecognizeRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRecognizeRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecognizeRecordsResponseBody) SetData(v []*GetRecognizeRecordsResponseBodyData) *GetRecognizeRecordsResponseBody {
	s.Data = v
	return s
}

func (s *GetRecognizeRecordsResponseBody) SetNextToken(v int64) *GetRecognizeRecordsResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetRecognizeRecordsResponseBody) SetTotal(v int32) *GetRecognizeRecordsResponseBody {
	s.Total = &v
	return s
}

type GetRecognizeRecordsResponseBodyData struct {
	AgentId           *int64  `json:"agentId,omitempty" xml:"agentId,omitempty"`
	FaceCompareResult *int32  `json:"faceCompareResult,omitempty" xml:"faceCompareResult,omitempty"`
	InvokeTime        *int64  `json:"invokeTime,omitempty" xml:"invokeTime,omitempty"`
	Platform          *int32  `json:"platform,omitempty" xml:"platform,omitempty"`
	UserId            *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetRecognizeRecordsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRecognizeRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRecognizeRecordsResponseBodyData) SetAgentId(v int64) *GetRecognizeRecordsResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *GetRecognizeRecordsResponseBodyData) SetFaceCompareResult(v int32) *GetRecognizeRecordsResponseBodyData {
	s.FaceCompareResult = &v
	return s
}

func (s *GetRecognizeRecordsResponseBodyData) SetInvokeTime(v int64) *GetRecognizeRecordsResponseBodyData {
	s.InvokeTime = &v
	return s
}

func (s *GetRecognizeRecordsResponseBodyData) SetPlatform(v int32) *GetRecognizeRecordsResponseBodyData {
	s.Platform = &v
	return s
}

func (s *GetRecognizeRecordsResponseBodyData) SetUserId(v string) *GetRecognizeRecordsResponseBodyData {
	s.UserId = &v
	return s
}

type GetRecognizeRecordsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRecognizeRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRecognizeRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRecognizeRecordsResponse) GoString() string {
	return s.String()
}

func (s *GetRecognizeRecordsResponse) SetHeaders(v map[string]*string) *GetRecognizeRecordsResponse {
	s.Headers = v
	return s
}

func (s *GetRecognizeRecordsResponse) SetStatusCode(v int32) *GetRecognizeRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecognizeRecordsResponse) SetBody(v *GetRecognizeRecordsResponseBody) *GetRecognizeRecordsResponse {
	s.Body = v
	return s
}

type GetSignedDetailByPageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSignedDetailByPageHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSignedDetailByPageHeaders) GoString() string {
	return s.String()
}

func (s *GetSignedDetailByPageHeaders) SetCommonHeaders(v map[string]*string) *GetSignedDetailByPageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSignedDetailByPageHeaders) SetXAcsDingtalkAccessToken(v string) *GetSignedDetailByPageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSignedDetailByPageRequest struct {
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	SignStatus *int64 `json:"signStatus,omitempty" xml:"signStatus,omitempty"`
}

func (s GetSignedDetailByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSignedDetailByPageRequest) GoString() string {
	return s.String()
}

func (s *GetSignedDetailByPageRequest) SetPageNumber(v int64) *GetSignedDetailByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *GetSignedDetailByPageRequest) SetPageSize(v int64) *GetSignedDetailByPageRequest {
	s.PageSize = &v
	return s
}

func (s *GetSignedDetailByPageRequest) SetSignStatus(v int64) *GetSignedDetailByPageRequest {
	s.SignStatus = &v
	return s
}

type GetSignedDetailByPageResponseBody struct {
	AuditSignedDetailDTOList []*GetSignedDetailByPageResponseBodyAuditSignedDetailDTOList `json:"auditSignedDetailDTOList,omitempty" xml:"auditSignedDetailDTOList,omitempty" type:"Repeated"`
	CurrentPage              *int64                                                       `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	PageSize                 *int64                                                       `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total                    *int64                                                       `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetSignedDetailByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSignedDetailByPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetSignedDetailByPageResponseBody) SetAuditSignedDetailDTOList(v []*GetSignedDetailByPageResponseBodyAuditSignedDetailDTOList) *GetSignedDetailByPageResponseBody {
	s.AuditSignedDetailDTOList = v
	return s
}

func (s *GetSignedDetailByPageResponseBody) SetCurrentPage(v int64) *GetSignedDetailByPageResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *GetSignedDetailByPageResponseBody) SetPageSize(v int64) *GetSignedDetailByPageResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetSignedDetailByPageResponseBody) SetTotal(v int64) *GetSignedDetailByPageResponseBody {
	s.Total = &v
	return s
}

type GetSignedDetailByPageResponseBodyAuditSignedDetailDTOList struct {
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	Email    *string `json:"email,omitempty" xml:"email,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Phone    *string `json:"phone,omitempty" xml:"phone,omitempty"`
	Roles    *string `json:"roles,omitempty" xml:"roles,omitempty"`
	StaffId  *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
	Title    *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetSignedDetailByPageResponseBodyAuditSignedDetailDTOList) String() string {
	return tea.Prettify(s)
}

func (s GetSignedDetailByPageResponseBodyAuditSignedDetailDTOList) GoString() string {
	return s.String()
}

func (s *GetSignedDetailByPageResponseBodyAuditSignedDetailDTOList) SetDeptName(v string) *GetSignedDetailByPageResponseBodyAuditSignedDetailDTOList {
	s.DeptName = &v
	return s
}

func (s *GetSignedDetailByPageResponseBodyAuditSignedDetailDTOList) SetEmail(v string) *GetSignedDetailByPageResponseBodyAuditSignedDetailDTOList {
	s.Email = &v
	return s
}

func (s *GetSignedDetailByPageResponseBodyAuditSignedDetailDTOList) SetName(v string) *GetSignedDetailByPageResponseBodyAuditSignedDetailDTOList {
	s.Name = &v
	return s
}

func (s *GetSignedDetailByPageResponseBodyAuditSignedDetailDTOList) SetPhone(v string) *GetSignedDetailByPageResponseBodyAuditSignedDetailDTOList {
	s.Phone = &v
	return s
}

func (s *GetSignedDetailByPageResponseBodyAuditSignedDetailDTOList) SetRoles(v string) *GetSignedDetailByPageResponseBodyAuditSignedDetailDTOList {
	s.Roles = &v
	return s
}

func (s *GetSignedDetailByPageResponseBodyAuditSignedDetailDTOList) SetStaffId(v string) *GetSignedDetailByPageResponseBodyAuditSignedDetailDTOList {
	s.StaffId = &v
	return s
}

func (s *GetSignedDetailByPageResponseBodyAuditSignedDetailDTOList) SetTitle(v string) *GetSignedDetailByPageResponseBodyAuditSignedDetailDTOList {
	s.Title = &v
	return s
}

type GetSignedDetailByPageResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSignedDetailByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSignedDetailByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSignedDetailByPageResponse) GoString() string {
	return s.String()
}

func (s *GetSignedDetailByPageResponse) SetHeaders(v map[string]*string) *GetSignedDetailByPageResponse {
	s.Headers = v
	return s
}

func (s *GetSignedDetailByPageResponse) SetStatusCode(v int32) *GetSignedDetailByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSignedDetailByPageResponse) SetBody(v *GetSignedDetailByPageResponseBody) *GetSignedDetailByPageResponse {
	s.Body = v
	return s
}

type GetTrustDeviceListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTrustDeviceListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTrustDeviceListHeaders) GoString() string {
	return s.String()
}

func (s *GetTrustDeviceListHeaders) SetCommonHeaders(v map[string]*string) *GetTrustDeviceListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTrustDeviceListHeaders) SetXAcsDingtalkAccessToken(v string) *GetTrustDeviceListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTrustDeviceListRequest struct {
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s GetTrustDeviceListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTrustDeviceListRequest) GoString() string {
	return s.String()
}

func (s *GetTrustDeviceListRequest) SetUserIds(v []*string) *GetTrustDeviceListRequest {
	s.UserIds = v
	return s
}

type GetTrustDeviceListResponseBody struct {
	Data []*GetTrustDeviceListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s GetTrustDeviceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrustDeviceListResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrustDeviceListResponseBody) SetData(v []*GetTrustDeviceListResponseBodyData) *GetTrustDeviceListResponseBody {
	s.Data = v
	return s
}

type GetTrustDeviceListResponseBodyData struct {
	CreateTime *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	MacAddress *string `json:"macAddress,omitempty" xml:"macAddress,omitempty"`
	Model      *string `json:"model,omitempty" xml:"model,omitempty"`
	Platform   *string `json:"platform,omitempty" xml:"platform,omitempty"`
	Status     *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Title      *string `json:"title,omitempty" xml:"title,omitempty"`
	UserId     *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetTrustDeviceListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTrustDeviceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTrustDeviceListResponseBodyData) SetCreateTime(v int64) *GetTrustDeviceListResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetTrustDeviceListResponseBodyData) SetMacAddress(v string) *GetTrustDeviceListResponseBodyData {
	s.MacAddress = &v
	return s
}

func (s *GetTrustDeviceListResponseBodyData) SetModel(v string) *GetTrustDeviceListResponseBodyData {
	s.Model = &v
	return s
}

func (s *GetTrustDeviceListResponseBodyData) SetPlatform(v string) *GetTrustDeviceListResponseBodyData {
	s.Platform = &v
	return s
}

func (s *GetTrustDeviceListResponseBodyData) SetStatus(v int32) *GetTrustDeviceListResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTrustDeviceListResponseBodyData) SetTitle(v string) *GetTrustDeviceListResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetTrustDeviceListResponseBodyData) SetUserId(v string) *GetTrustDeviceListResponseBodyData {
	s.UserId = &v
	return s
}

type GetTrustDeviceListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTrustDeviceListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTrustDeviceListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrustDeviceListResponse) GoString() string {
	return s.String()
}

func (s *GetTrustDeviceListResponse) SetHeaders(v map[string]*string) *GetTrustDeviceListResponse {
	s.Headers = v
	return s
}

func (s *GetTrustDeviceListResponse) SetStatusCode(v int32) *GetTrustDeviceListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrustDeviceListResponse) SetBody(v *GetTrustDeviceListResponseBody) *GetTrustDeviceListResponse {
	s.Body = v
	return s
}

type GetUserAppVersionSummaryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserAppVersionSummaryHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserAppVersionSummaryHeaders) GoString() string {
	return s.String()
}

func (s *GetUserAppVersionSummaryHeaders) SetCommonHeaders(v map[string]*string) *GetUserAppVersionSummaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserAppVersionSummaryHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserAppVersionSummaryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserAppVersionSummaryRequest struct {
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetUserAppVersionSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserAppVersionSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetUserAppVersionSummaryRequest) SetMaxResults(v int64) *GetUserAppVersionSummaryRequest {
	s.MaxResults = &v
	return s
}

func (s *GetUserAppVersionSummaryRequest) SetNextToken(v int64) *GetUserAppVersionSummaryRequest {
	s.NextToken = &v
	return s
}

type GetUserAppVersionSummaryResponseBody struct {
	Data      []*GetUserAppVersionSummaryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	HasMore   *bool                                       `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken *int64                                      `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetUserAppVersionSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserAppVersionSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserAppVersionSummaryResponseBody) SetData(v []*GetUserAppVersionSummaryResponseBodyData) *GetUserAppVersionSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetUserAppVersionSummaryResponseBody) SetHasMore(v bool) *GetUserAppVersionSummaryResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetUserAppVersionSummaryResponseBody) SetNextToken(v int64) *GetUserAppVersionSummaryResponseBody {
	s.NextToken = &v
	return s
}

type GetUserAppVersionSummaryResponseBodyData struct {
	AppVersion *string  `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	Client     *string  `json:"client,omitempty" xml:"client,omitempty"`
	OrgName    *string  `json:"orgName,omitempty" xml:"orgName,omitempty"`
	StatDate   *string  `json:"statDate,omitempty" xml:"statDate,omitempty"`
	UserCnt    *float32 `json:"userCnt,omitempty" xml:"userCnt,omitempty"`
}

func (s GetUserAppVersionSummaryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUserAppVersionSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserAppVersionSummaryResponseBodyData) SetAppVersion(v string) *GetUserAppVersionSummaryResponseBodyData {
	s.AppVersion = &v
	return s
}

func (s *GetUserAppVersionSummaryResponseBodyData) SetClient(v string) *GetUserAppVersionSummaryResponseBodyData {
	s.Client = &v
	return s
}

func (s *GetUserAppVersionSummaryResponseBodyData) SetOrgName(v string) *GetUserAppVersionSummaryResponseBodyData {
	s.OrgName = &v
	return s
}

func (s *GetUserAppVersionSummaryResponseBodyData) SetStatDate(v string) *GetUserAppVersionSummaryResponseBodyData {
	s.StatDate = &v
	return s
}

func (s *GetUserAppVersionSummaryResponseBodyData) SetUserCnt(v float32) *GetUserAppVersionSummaryResponseBodyData {
	s.UserCnt = &v
	return s
}

type GetUserAppVersionSummaryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserAppVersionSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserAppVersionSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserAppVersionSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetUserAppVersionSummaryResponse) SetHeaders(v map[string]*string) *GetUserAppVersionSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetUserAppVersionSummaryResponse) SetStatusCode(v int32) *GetUserAppVersionSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserAppVersionSummaryResponse) SetBody(v *GetUserAppVersionSummaryResponseBody) *GetUserAppVersionSummaryResponse {
	s.Body = v
	return s
}

type GetUserFaceStateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserFaceStateHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserFaceStateHeaders) GoString() string {
	return s.String()
}

func (s *GetUserFaceStateHeaders) SetCommonHeaders(v map[string]*string) *GetUserFaceStateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserFaceStateHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserFaceStateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserFaceStateRequest struct {
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s GetUserFaceStateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserFaceStateRequest) GoString() string {
	return s.String()
}

func (s *GetUserFaceStateRequest) SetUserIds(v []*string) *GetUserFaceStateRequest {
	s.UserIds = v
	return s
}

type GetUserFaceStateResponseBody struct {
	Data []*GetUserFaceStateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s GetUserFaceStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserFaceStateResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserFaceStateResponseBody) SetData(v []*GetUserFaceStateResponseBodyData) *GetUserFaceStateResponseBody {
	s.Data = v
	return s
}

type GetUserFaceStateResponseBodyData struct {
	State  *int32  `json:"state,omitempty" xml:"state,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserFaceStateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUserFaceStateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserFaceStateResponseBodyData) SetState(v int32) *GetUserFaceStateResponseBodyData {
	s.State = &v
	return s
}

func (s *GetUserFaceStateResponseBodyData) SetUserId(v string) *GetUserFaceStateResponseBodyData {
	s.UserId = &v
	return s
}

type GetUserFaceStateResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserFaceStateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserFaceStateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserFaceStateResponse) GoString() string {
	return s.String()
}

func (s *GetUserFaceStateResponse) SetHeaders(v map[string]*string) *GetUserFaceStateResponse {
	s.Headers = v
	return s
}

func (s *GetUserFaceStateResponse) SetStatusCode(v int32) *GetUserFaceStateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserFaceStateResponse) SetBody(v *GetUserFaceStateResponseBody) *GetUserFaceStateResponse {
	s.Body = v
	return s
}

type GetUserRealPeopleStateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserRealPeopleStateHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserRealPeopleStateHeaders) GoString() string {
	return s.String()
}

func (s *GetUserRealPeopleStateHeaders) SetCommonHeaders(v map[string]*string) *GetUserRealPeopleStateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserRealPeopleStateHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserRealPeopleStateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserRealPeopleStateRequest struct {
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s GetUserRealPeopleStateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserRealPeopleStateRequest) GoString() string {
	return s.String()
}

func (s *GetUserRealPeopleStateRequest) SetUserIds(v []*string) *GetUserRealPeopleStateRequest {
	s.UserIds = v
	return s
}

type GetUserRealPeopleStateResponseBody struct {
	Data []*GetUserRealPeopleStateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s GetUserRealPeopleStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserRealPeopleStateResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserRealPeopleStateResponseBody) SetData(v []*GetUserRealPeopleStateResponseBodyData) *GetUserRealPeopleStateResponseBody {
	s.Data = v
	return s
}

type GetUserRealPeopleStateResponseBodyData struct {
	State  *int32  `json:"state,omitempty" xml:"state,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserRealPeopleStateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUserRealPeopleStateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserRealPeopleStateResponseBodyData) SetState(v int32) *GetUserRealPeopleStateResponseBodyData {
	s.State = &v
	return s
}

func (s *GetUserRealPeopleStateResponseBodyData) SetUserId(v string) *GetUserRealPeopleStateResponseBodyData {
	s.UserId = &v
	return s
}

type GetUserRealPeopleStateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserRealPeopleStateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserRealPeopleStateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserRealPeopleStateResponse) GoString() string {
	return s.String()
}

func (s *GetUserRealPeopleStateResponse) SetHeaders(v map[string]*string) *GetUserRealPeopleStateResponse {
	s.Headers = v
	return s
}

func (s *GetUserRealPeopleStateResponse) SetStatusCode(v int32) *GetUserRealPeopleStateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserRealPeopleStateResponse) SetBody(v *GetUserRealPeopleStateResponseBody) *GetUserRealPeopleStateResponse {
	s.Body = v
	return s
}

type GetUserStayLengthHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserStayLengthHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserStayLengthHeaders) GoString() string {
	return s.String()
}

func (s *GetUserStayLengthHeaders) SetCommonHeaders(v map[string]*string) *GetUserStayLengthHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserStayLengthHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserStayLengthHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserStayLengthRequest struct {
	PageNumber *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StatDate   *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s GetUserStayLengthRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserStayLengthRequest) GoString() string {
	return s.String()
}

func (s *GetUserStayLengthRequest) SetPageNumber(v int64) *GetUserStayLengthRequest {
	s.PageNumber = &v
	return s
}

func (s *GetUserStayLengthRequest) SetPageSize(v int64) *GetUserStayLengthRequest {
	s.PageSize = &v
	return s
}

func (s *GetUserStayLengthRequest) SetStatDate(v string) *GetUserStayLengthRequest {
	s.StatDate = &v
	return s
}

type GetUserStayLengthResponseBody struct {
	ItemList   []*GetUserStayLengthResponseBodyItemList `json:"itemList,omitempty" xml:"itemList,omitempty" type:"Repeated"`
	TotalCount *int64                                   `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetUserStayLengthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserStayLengthResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserStayLengthResponseBody) SetItemList(v []*GetUserStayLengthResponseBodyItemList) *GetUserStayLengthResponseBody {
	s.ItemList = v
	return s
}

func (s *GetUserStayLengthResponseBody) SetTotalCount(v int64) *GetUserStayLengthResponseBody {
	s.TotalCount = &v
	return s
}

type GetUserStayLengthResponseBodyItemList struct {
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
	StatDate         *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
	StayTimeLenApp1d *int64  `json:"stayTimeLenApp1d,omitempty" xml:"stayTimeLenApp1d,omitempty"`
	StayTimeLenPc1d  *int64  `json:"stayTimeLenPc1d,omitempty" xml:"stayTimeLenPc1d,omitempty"`
	UserId           *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserStayLengthResponseBodyItemList) String() string {
	return tea.Prettify(s)
}

func (s GetUserStayLengthResponseBodyItemList) GoString() string {
	return s.String()
}

func (s *GetUserStayLengthResponseBodyItemList) SetName(v string) *GetUserStayLengthResponseBodyItemList {
	s.Name = &v
	return s
}

func (s *GetUserStayLengthResponseBodyItemList) SetStatDate(v string) *GetUserStayLengthResponseBodyItemList {
	s.StatDate = &v
	return s
}

func (s *GetUserStayLengthResponseBodyItemList) SetStayTimeLenApp1d(v int64) *GetUserStayLengthResponseBodyItemList {
	s.StayTimeLenApp1d = &v
	return s
}

func (s *GetUserStayLengthResponseBodyItemList) SetStayTimeLenPc1d(v int64) *GetUserStayLengthResponseBodyItemList {
	s.StayTimeLenPc1d = &v
	return s
}

func (s *GetUserStayLengthResponseBodyItemList) SetUserId(v string) *GetUserStayLengthResponseBodyItemList {
	s.UserId = &v
	return s
}

type GetUserStayLengthResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserStayLengthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserStayLengthResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserStayLengthResponse) GoString() string {
	return s.String()
}

func (s *GetUserStayLengthResponse) SetHeaders(v map[string]*string) *GetUserStayLengthResponse {
	s.Headers = v
	return s
}

func (s *GetUserStayLengthResponse) SetStatusCode(v int32) *GetUserStayLengthResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserStayLengthResponse) SetBody(v *GetUserStayLengthResponseBody) *GetUserStayLengthResponse {
	s.Body = v
	return s
}

type ListAuditLogHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListAuditLogHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListAuditLogHeaders) GoString() string {
	return s.String()
}

func (s *ListAuditLogHeaders) SetCommonHeaders(v map[string]*string) *ListAuditLogHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAuditLogHeaders) SetXAcsDingtalkAccessToken(v string) *ListAuditLogHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListAuditLogRequest struct {
	EndDate       *int64 `json:"endDate,omitempty" xml:"endDate,omitempty"`
	NextBizId     *int64 `json:"nextBizId,omitempty" xml:"nextBizId,omitempty"`
	NextGmtCreate *int64 `json:"nextGmtCreate,omitempty" xml:"nextGmtCreate,omitempty"`
	PageSize      *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StartDate     *int64 `json:"startDate,omitempty" xml:"startDate,omitempty"`
}

func (s ListAuditLogRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAuditLogRequest) GoString() string {
	return s.String()
}

func (s *ListAuditLogRequest) SetEndDate(v int64) *ListAuditLogRequest {
	s.EndDate = &v
	return s
}

func (s *ListAuditLogRequest) SetNextBizId(v int64) *ListAuditLogRequest {
	s.NextBizId = &v
	return s
}

func (s *ListAuditLogRequest) SetNextGmtCreate(v int64) *ListAuditLogRequest {
	s.NextGmtCreate = &v
	return s
}

func (s *ListAuditLogRequest) SetPageSize(v int32) *ListAuditLogRequest {
	s.PageSize = &v
	return s
}

func (s *ListAuditLogRequest) SetStartDate(v int64) *ListAuditLogRequest {
	s.StartDate = &v
	return s
}

type ListAuditLogResponseBody struct {
	List []*ListAuditLogResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s ListAuditLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAuditLogResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuditLogResponseBody) SetList(v []*ListAuditLogResponseBodyList) *ListAuditLogResponseBody {
	s.List = v
	return s
}

type ListAuditLogResponseBodyList struct {
	Action                 *int32                                         `json:"action,omitempty" xml:"action,omitempty"`
	ActionView             *string                                        `json:"actionView,omitempty" xml:"actionView,omitempty"`
	BizId                  *string                                        `json:"bizId,omitempty" xml:"bizId,omitempty"`
	DocMemberList          []*ListAuditLogResponseBodyListDocMemberList   `json:"docMemberList,omitempty" xml:"docMemberList,omitempty" type:"Repeated"`
	DocMobileUrl           *string                                        `json:"docMobileUrl,omitempty" xml:"docMobileUrl,omitempty"`
	DocPcUrl               *string                                        `json:"docPcUrl,omitempty" xml:"docPcUrl,omitempty"`
	DocReceiverList        []*ListAuditLogResponseBodyListDocReceiverList `json:"docReceiverList,omitempty" xml:"docReceiverList,omitempty" type:"Repeated"`
	GmtCreate              *int64                                         `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified            *int64                                         `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	IpAddress              *string                                        `json:"ipAddress,omitempty" xml:"ipAddress,omitempty"`
	OperateModule          *int64                                         `json:"operateModule,omitempty" xml:"operateModule,omitempty"`
	OperateModuleView      *string                                        `json:"operateModuleView,omitempty" xml:"operateModuleView,omitempty"`
	OperatorName           *string                                        `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	OrgName                *string                                        `json:"orgName,omitempty" xml:"orgName,omitempty"`
	Platform               *int32                                         `json:"platform,omitempty" xml:"platform,omitempty"`
	PlatformView           *string                                        `json:"platformView,omitempty" xml:"platformView,omitempty"`
	PrevWorkSpaceId        *int64                                         `json:"prevWorkSpaceId,omitempty" xml:"prevWorkSpaceId,omitempty"`
	PrevWorkSpaceMobileUrl *string                                        `json:"prevWorkSpaceMobileUrl,omitempty" xml:"prevWorkSpaceMobileUrl,omitempty"`
	PrevWorkSpaceName      *string                                        `json:"prevWorkSpaceName,omitempty" xml:"prevWorkSpaceName,omitempty"`
	PrevWorkSpacePcUrl     *string                                        `json:"prevWorkSpacePcUrl,omitempty" xml:"prevWorkSpacePcUrl,omitempty"`
	RealName               *string                                        `json:"realName,omitempty" xml:"realName,omitempty"`
	ReceiverName           *string                                        `json:"receiverName,omitempty" xml:"receiverName,omitempty"`
	ReceiverType           *int32                                         `json:"receiverType,omitempty" xml:"receiverType,omitempty"`
	ReceiverTypeView       *string                                        `json:"receiverTypeView,omitempty" xml:"receiverTypeView,omitempty"`
	Resource               *string                                        `json:"resource,omitempty" xml:"resource,omitempty"`
	ResourceExtension      *string                                        `json:"resourceExtension,omitempty" xml:"resourceExtension,omitempty"`
	ResourceSize           *int64                                         `json:"resourceSize,omitempty" xml:"resourceSize,omitempty"`
	Status                 *int32                                         `json:"status,omitempty" xml:"status,omitempty"`
	TargetSpaceId          *int64                                         `json:"targetSpaceId,omitempty" xml:"targetSpaceId,omitempty"`
	UserId                 *string                                        `json:"userId,omitempty" xml:"userId,omitempty"`
	WorkSpaceId            *int64                                         `json:"workSpaceId,omitempty" xml:"workSpaceId,omitempty"`
	WorkSpaceMobileUrl     *string                                        `json:"workSpaceMobileUrl,omitempty" xml:"workSpaceMobileUrl,omitempty"`
	WorkSpaceName          *string                                        `json:"workSpaceName,omitempty" xml:"workSpaceName,omitempty"`
	WorkSpacePcUrl         *string                                        `json:"workSpacePcUrl,omitempty" xml:"workSpacePcUrl,omitempty"`
}

func (s ListAuditLogResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListAuditLogResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListAuditLogResponseBodyList) SetAction(v int32) *ListAuditLogResponseBodyList {
	s.Action = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetActionView(v string) *ListAuditLogResponseBodyList {
	s.ActionView = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetBizId(v string) *ListAuditLogResponseBodyList {
	s.BizId = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetDocMemberList(v []*ListAuditLogResponseBodyListDocMemberList) *ListAuditLogResponseBodyList {
	s.DocMemberList = v
	return s
}

func (s *ListAuditLogResponseBodyList) SetDocMobileUrl(v string) *ListAuditLogResponseBodyList {
	s.DocMobileUrl = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetDocPcUrl(v string) *ListAuditLogResponseBodyList {
	s.DocPcUrl = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetDocReceiverList(v []*ListAuditLogResponseBodyListDocReceiverList) *ListAuditLogResponseBodyList {
	s.DocReceiverList = v
	return s
}

func (s *ListAuditLogResponseBodyList) SetGmtCreate(v int64) *ListAuditLogResponseBodyList {
	s.GmtCreate = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetGmtModified(v int64) *ListAuditLogResponseBodyList {
	s.GmtModified = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetIpAddress(v string) *ListAuditLogResponseBodyList {
	s.IpAddress = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetOperateModule(v int64) *ListAuditLogResponseBodyList {
	s.OperateModule = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetOperateModuleView(v string) *ListAuditLogResponseBodyList {
	s.OperateModuleView = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetOperatorName(v string) *ListAuditLogResponseBodyList {
	s.OperatorName = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetOrgName(v string) *ListAuditLogResponseBodyList {
	s.OrgName = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetPlatform(v int32) *ListAuditLogResponseBodyList {
	s.Platform = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetPlatformView(v string) *ListAuditLogResponseBodyList {
	s.PlatformView = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetPrevWorkSpaceId(v int64) *ListAuditLogResponseBodyList {
	s.PrevWorkSpaceId = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetPrevWorkSpaceMobileUrl(v string) *ListAuditLogResponseBodyList {
	s.PrevWorkSpaceMobileUrl = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetPrevWorkSpaceName(v string) *ListAuditLogResponseBodyList {
	s.PrevWorkSpaceName = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetPrevWorkSpacePcUrl(v string) *ListAuditLogResponseBodyList {
	s.PrevWorkSpacePcUrl = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetRealName(v string) *ListAuditLogResponseBodyList {
	s.RealName = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetReceiverName(v string) *ListAuditLogResponseBodyList {
	s.ReceiverName = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetReceiverType(v int32) *ListAuditLogResponseBodyList {
	s.ReceiverType = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetReceiverTypeView(v string) *ListAuditLogResponseBodyList {
	s.ReceiverTypeView = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetResource(v string) *ListAuditLogResponseBodyList {
	s.Resource = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetResourceExtension(v string) *ListAuditLogResponseBodyList {
	s.ResourceExtension = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetResourceSize(v int64) *ListAuditLogResponseBodyList {
	s.ResourceSize = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetStatus(v int32) *ListAuditLogResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetTargetSpaceId(v int64) *ListAuditLogResponseBodyList {
	s.TargetSpaceId = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetUserId(v string) *ListAuditLogResponseBodyList {
	s.UserId = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetWorkSpaceId(v int64) *ListAuditLogResponseBodyList {
	s.WorkSpaceId = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetWorkSpaceMobileUrl(v string) *ListAuditLogResponseBodyList {
	s.WorkSpaceMobileUrl = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetWorkSpaceName(v string) *ListAuditLogResponseBodyList {
	s.WorkSpaceName = &v
	return s
}

func (s *ListAuditLogResponseBodyList) SetWorkSpacePcUrl(v string) *ListAuditLogResponseBodyList {
	s.WorkSpacePcUrl = &v
	return s
}

type ListAuditLogResponseBodyListDocMemberList struct {
	MemberName         *string `json:"memberName,omitempty" xml:"memberName,omitempty"`
	MemberType         *int32  `json:"memberType,omitempty" xml:"memberType,omitempty"`
	MemberTypeView     *string `json:"memberTypeView,omitempty" xml:"memberTypeView,omitempty"`
	PermissionRole     *int64  `json:"permissionRole,omitempty" xml:"permissionRole,omitempty"`
	PermissionRoleView *string `json:"permissionRoleView,omitempty" xml:"permissionRoleView,omitempty"`
}

func (s ListAuditLogResponseBodyListDocMemberList) String() string {
	return tea.Prettify(s)
}

func (s ListAuditLogResponseBodyListDocMemberList) GoString() string {
	return s.String()
}

func (s *ListAuditLogResponseBodyListDocMemberList) SetMemberName(v string) *ListAuditLogResponseBodyListDocMemberList {
	s.MemberName = &v
	return s
}

func (s *ListAuditLogResponseBodyListDocMemberList) SetMemberType(v int32) *ListAuditLogResponseBodyListDocMemberList {
	s.MemberType = &v
	return s
}

func (s *ListAuditLogResponseBodyListDocMemberList) SetMemberTypeView(v string) *ListAuditLogResponseBodyListDocMemberList {
	s.MemberTypeView = &v
	return s
}

func (s *ListAuditLogResponseBodyListDocMemberList) SetPermissionRole(v int64) *ListAuditLogResponseBodyListDocMemberList {
	s.PermissionRole = &v
	return s
}

func (s *ListAuditLogResponseBodyListDocMemberList) SetPermissionRoleView(v string) *ListAuditLogResponseBodyListDocMemberList {
	s.PermissionRoleView = &v
	return s
}

type ListAuditLogResponseBodyListDocReceiverList struct {
	ReceiverName     *string `json:"receiverName,omitempty" xml:"receiverName,omitempty"`
	ReceiverType     *int32  `json:"receiverType,omitempty" xml:"receiverType,omitempty"`
	ReceiverTypeView *string `json:"receiverTypeView,omitempty" xml:"receiverTypeView,omitempty"`
}

func (s ListAuditLogResponseBodyListDocReceiverList) String() string {
	return tea.Prettify(s)
}

func (s ListAuditLogResponseBodyListDocReceiverList) GoString() string {
	return s.String()
}

func (s *ListAuditLogResponseBodyListDocReceiverList) SetReceiverName(v string) *ListAuditLogResponseBodyListDocReceiverList {
	s.ReceiverName = &v
	return s
}

func (s *ListAuditLogResponseBodyListDocReceiverList) SetReceiverType(v int32) *ListAuditLogResponseBodyListDocReceiverList {
	s.ReceiverType = &v
	return s
}

func (s *ListAuditLogResponseBodyListDocReceiverList) SetReceiverTypeView(v string) *ListAuditLogResponseBodyListDocReceiverList {
	s.ReceiverTypeView = &v
	return s
}

type ListAuditLogResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAuditLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAuditLogResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAuditLogResponse) GoString() string {
	return s.String()
}

func (s *ListAuditLogResponse) SetHeaders(v map[string]*string) *ListAuditLogResponse {
	s.Headers = v
	return s
}

func (s *ListAuditLogResponse) SetStatusCode(v int32) *ListAuditLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuditLogResponse) SetBody(v *ListAuditLogResponseBody) *ListAuditLogResponse {
	s.Body = v
	return s
}

type ListCategorysHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListCategorysHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListCategorysHeaders) GoString() string {
	return s.String()
}

func (s *ListCategorysHeaders) SetCommonHeaders(v map[string]*string) *ListCategorysHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListCategorysHeaders) SetXAcsDingtalkAccessToken(v string) *ListCategorysHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListCategorysRequest struct {
	Body *ListCategorysRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s ListCategorysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCategorysRequest) GoString() string {
	return s.String()
}

func (s *ListCategorysRequest) SetBody(v *ListCategorysRequestBody) *ListCategorysRequest {
	s.Body = v
	return s
}

type ListCategorysRequestBody struct {
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListCategorysRequestBody) String() string {
	return tea.Prettify(s)
}

func (s ListCategorysRequestBody) GoString() string {
	return s.String()
}

func (s *ListCategorysRequestBody) SetStatus(v int64) *ListCategorysRequestBody {
	s.Status = &v
	return s
}

type ListCategorysShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCategorysShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCategorysShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCategorysShrinkRequest) SetBodyShrink(v string) *ListCategorysShrinkRequest {
	s.BodyShrink = &v
	return s
}

type ListCategorysResponseBody struct {
	DetailModelList []map[string]*string `json:"detailModelList,omitempty" xml:"detailModelList,omitempty" type:"Repeated"`
}

func (s ListCategorysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCategorysResponseBody) GoString() string {
	return s.String()
}

func (s *ListCategorysResponseBody) SetDetailModelList(v []map[string]*string) *ListCategorysResponseBody {
	s.DetailModelList = v
	return s
}

type ListCategorysResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCategorysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCategorysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCategorysResponse) GoString() string {
	return s.String()
}

func (s *ListCategorysResponse) SetHeaders(v map[string]*string) *ListCategorysResponse {
	s.Headers = v
	return s
}

func (s *ListCategorysResponse) SetStatusCode(v int32) *ListCategorysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCategorysResponse) SetBody(v *ListCategorysResponseBody) *ListCategorysResponse {
	s.Body = v
	return s
}

type ListJoinOrgInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListJoinOrgInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListJoinOrgInfoHeaders) GoString() string {
	return s.String()
}

func (s *ListJoinOrgInfoHeaders) SetCommonHeaders(v map[string]*string) *ListJoinOrgInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListJoinOrgInfoHeaders) SetXAcsDingtalkAccessToken(v string) *ListJoinOrgInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListJoinOrgInfoRequest struct {
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
}

func (s ListJoinOrgInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJoinOrgInfoRequest) GoString() string {
	return s.String()
}

func (s *ListJoinOrgInfoRequest) SetMobile(v string) *ListJoinOrgInfoRequest {
	s.Mobile = &v
	return s
}

type ListJoinOrgInfoResponseBody struct {
	OrgInfoList []*ListJoinOrgInfoResponseBodyOrgInfoList `json:"orgInfoList,omitempty" xml:"orgInfoList,omitempty" type:"Repeated"`
}

func (s ListJoinOrgInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJoinOrgInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListJoinOrgInfoResponseBody) SetOrgInfoList(v []*ListJoinOrgInfoResponseBodyOrgInfoList) *ListJoinOrgInfoResponseBody {
	s.OrgInfoList = v
	return s
}

type ListJoinOrgInfoResponseBodyOrgInfoList struct {
	CorpId      *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Domain      *string `json:"domain,omitempty" xml:"domain,omitempty"`
	OrgFullName *string `json:"orgFullName,omitempty" xml:"orgFullName,omitempty"`
	OrgName     *int64  `json:"orgName,omitempty" xml:"orgName,omitempty"`
}

func (s ListJoinOrgInfoResponseBodyOrgInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListJoinOrgInfoResponseBodyOrgInfoList) GoString() string {
	return s.String()
}

func (s *ListJoinOrgInfoResponseBodyOrgInfoList) SetCorpId(v string) *ListJoinOrgInfoResponseBodyOrgInfoList {
	s.CorpId = &v
	return s
}

func (s *ListJoinOrgInfoResponseBodyOrgInfoList) SetDomain(v string) *ListJoinOrgInfoResponseBodyOrgInfoList {
	s.Domain = &v
	return s
}

func (s *ListJoinOrgInfoResponseBodyOrgInfoList) SetOrgFullName(v string) *ListJoinOrgInfoResponseBodyOrgInfoList {
	s.OrgFullName = &v
	return s
}

func (s *ListJoinOrgInfoResponseBodyOrgInfoList) SetOrgName(v int64) *ListJoinOrgInfoResponseBodyOrgInfoList {
	s.OrgName = &v
	return s
}

type ListJoinOrgInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListJoinOrgInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListJoinOrgInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJoinOrgInfoResponse) GoString() string {
	return s.String()
}

func (s *ListJoinOrgInfoResponse) SetHeaders(v map[string]*string) *ListJoinOrgInfoResponse {
	s.Headers = v
	return s
}

func (s *ListJoinOrgInfoResponse) SetStatusCode(v int32) *ListJoinOrgInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJoinOrgInfoResponse) SetBody(v *ListJoinOrgInfoResponseBody) *ListJoinOrgInfoResponse {
	s.Body = v
	return s
}

type ListMiniAppAvailableVersionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListMiniAppAvailableVersionHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListMiniAppAvailableVersionHeaders) GoString() string {
	return s.String()
}

func (s *ListMiniAppAvailableVersionHeaders) SetCommonHeaders(v map[string]*string) *ListMiniAppAvailableVersionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListMiniAppAvailableVersionHeaders) SetXAcsDingtalkAccessToken(v string) *ListMiniAppAvailableVersionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListMiniAppAvailableVersionRequest struct {
	MiniAppId      *string  `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	PageNumber     *int32   `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	VersionTypeSet []*int32 `json:"versionTypeSet,omitempty" xml:"versionTypeSet,omitempty" type:"Repeated"`
}

func (s ListMiniAppAvailableVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMiniAppAvailableVersionRequest) GoString() string {
	return s.String()
}

func (s *ListMiniAppAvailableVersionRequest) SetMiniAppId(v string) *ListMiniAppAvailableVersionRequest {
	s.MiniAppId = &v
	return s
}

func (s *ListMiniAppAvailableVersionRequest) SetPageNumber(v int32) *ListMiniAppAvailableVersionRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMiniAppAvailableVersionRequest) SetPageSize(v int32) *ListMiniAppAvailableVersionRequest {
	s.PageSize = &v
	return s
}

func (s *ListMiniAppAvailableVersionRequest) SetVersionTypeSet(v []*int32) *ListMiniAppAvailableVersionRequest {
	s.VersionTypeSet = v
	return s
}

type ListMiniAppAvailableVersionResponseBody struct {
	List []*ListMiniAppAvailableVersionResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s ListMiniAppAvailableVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMiniAppAvailableVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ListMiniAppAvailableVersionResponseBody) SetList(v []*ListMiniAppAvailableVersionResponseBodyList) *ListMiniAppAvailableVersionResponseBody {
	s.List = v
	return s
}

type ListMiniAppAvailableVersionResponseBodyList struct {
	BuildStatus *int64  `json:"buildStatus,omitempty" xml:"buildStatus,omitempty"`
	Version     *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListMiniAppAvailableVersionResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListMiniAppAvailableVersionResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListMiniAppAvailableVersionResponseBodyList) SetBuildStatus(v int64) *ListMiniAppAvailableVersionResponseBodyList {
	s.BuildStatus = &v
	return s
}

func (s *ListMiniAppAvailableVersionResponseBodyList) SetVersion(v string) *ListMiniAppAvailableVersionResponseBodyList {
	s.Version = &v
	return s
}

type ListMiniAppAvailableVersionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMiniAppAvailableVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMiniAppAvailableVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMiniAppAvailableVersionResponse) GoString() string {
	return s.String()
}

func (s *ListMiniAppAvailableVersionResponse) SetHeaders(v map[string]*string) *ListMiniAppAvailableVersionResponse {
	s.Headers = v
	return s
}

func (s *ListMiniAppAvailableVersionResponse) SetStatusCode(v int32) *ListMiniAppAvailableVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMiniAppAvailableVersionResponse) SetBody(v *ListMiniAppAvailableVersionResponseBody) *ListMiniAppAvailableVersionResponse {
	s.Body = v
	return s
}

type ListMiniAppHistoryVersionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListMiniAppHistoryVersionHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListMiniAppHistoryVersionHeaders) GoString() string {
	return s.String()
}

func (s *ListMiniAppHistoryVersionHeaders) SetCommonHeaders(v map[string]*string) *ListMiniAppHistoryVersionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListMiniAppHistoryVersionHeaders) SetXAcsDingtalkAccessToken(v string) *ListMiniAppHistoryVersionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListMiniAppHistoryVersionRequest struct {
	MiniAppId  *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	PageNumber *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListMiniAppHistoryVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMiniAppHistoryVersionRequest) GoString() string {
	return s.String()
}

func (s *ListMiniAppHistoryVersionRequest) SetMiniAppId(v string) *ListMiniAppHistoryVersionRequest {
	s.MiniAppId = &v
	return s
}

func (s *ListMiniAppHistoryVersionRequest) SetPageNumber(v int64) *ListMiniAppHistoryVersionRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMiniAppHistoryVersionRequest) SetPageSize(v int64) *ListMiniAppHistoryVersionRequest {
	s.PageSize = &v
	return s
}

type ListMiniAppHistoryVersionResponseBody struct {
	List []*ListMiniAppHistoryVersionResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s ListMiniAppHistoryVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMiniAppHistoryVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ListMiniAppHistoryVersionResponseBody) SetList(v []*ListMiniAppHistoryVersionResponseBodyList) *ListMiniAppHistoryVersionResponseBody {
	s.List = v
	return s
}

type ListMiniAppHistoryVersionResponseBodyList struct {
	BuildStatus *int64  `json:"buildStatus,omitempty" xml:"buildStatus,omitempty"`
	H5Bundle    *string `json:"h5Bundle,omitempty" xml:"h5Bundle,omitempty"`
	PackageSize *string `json:"packageSize,omitempty" xml:"packageSize,omitempty"`
	PackageUrl  *string `json:"packageUrl,omitempty" xml:"packageUrl,omitempty"`
	Version     *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListMiniAppHistoryVersionResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListMiniAppHistoryVersionResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListMiniAppHistoryVersionResponseBodyList) SetBuildStatus(v int64) *ListMiniAppHistoryVersionResponseBodyList {
	s.BuildStatus = &v
	return s
}

func (s *ListMiniAppHistoryVersionResponseBodyList) SetH5Bundle(v string) *ListMiniAppHistoryVersionResponseBodyList {
	s.H5Bundle = &v
	return s
}

func (s *ListMiniAppHistoryVersionResponseBodyList) SetPackageSize(v string) *ListMiniAppHistoryVersionResponseBodyList {
	s.PackageSize = &v
	return s
}

func (s *ListMiniAppHistoryVersionResponseBodyList) SetPackageUrl(v string) *ListMiniAppHistoryVersionResponseBodyList {
	s.PackageUrl = &v
	return s
}

func (s *ListMiniAppHistoryVersionResponseBodyList) SetVersion(v string) *ListMiniAppHistoryVersionResponseBodyList {
	s.Version = &v
	return s
}

type ListMiniAppHistoryVersionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMiniAppHistoryVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMiniAppHistoryVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMiniAppHistoryVersionResponse) GoString() string {
	return s.String()
}

func (s *ListMiniAppHistoryVersionResponse) SetHeaders(v map[string]*string) *ListMiniAppHistoryVersionResponse {
	s.Headers = v
	return s
}

func (s *ListMiniAppHistoryVersionResponse) SetStatusCode(v int32) *ListMiniAppHistoryVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMiniAppHistoryVersionResponse) SetBody(v *ListMiniAppHistoryVersionResponseBody) *ListMiniAppHistoryVersionResponse {
	s.Body = v
	return s
}

type ListPartnerRolesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListPartnerRolesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListPartnerRolesHeaders) GoString() string {
	return s.String()
}

func (s *ListPartnerRolesHeaders) SetCommonHeaders(v map[string]*string) *ListPartnerRolesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListPartnerRolesHeaders) SetXAcsDingtalkAccessToken(v string) *ListPartnerRolesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListPartnerRolesResponseBody struct {
	List []*ListPartnerRolesResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s ListPartnerRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPartnerRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPartnerRolesResponseBody) SetList(v []*ListPartnerRolesResponseBodyList) *ListPartnerRolesResponseBody {
	s.List = v
	return s
}

type ListPartnerRolesResponseBodyList struct {
	Id           *int64                                          `json:"id,omitempty" xml:"id,omitempty"`
	IsNecessary  *int32                                          `json:"isNecessary,omitempty" xml:"isNecessary,omitempty"`
	Name         *string                                         `json:"name,omitempty" xml:"name,omitempty"`
	VisibleDepts []*ListPartnerRolesResponseBodyListVisibleDepts `json:"visibleDepts,omitempty" xml:"visibleDepts,omitempty" type:"Repeated"`
	VisibleUsers []*ListPartnerRolesResponseBodyListVisibleUsers `json:"visibleUsers,omitempty" xml:"visibleUsers,omitempty" type:"Repeated"`
	WarningDepts []*ListPartnerRolesResponseBodyListWarningDepts `json:"warningDepts,omitempty" xml:"warningDepts,omitempty" type:"Repeated"`
	WarningUsers []*ListPartnerRolesResponseBodyListWarningUsers `json:"warningUsers,omitempty" xml:"warningUsers,omitempty" type:"Repeated"`
}

func (s ListPartnerRolesResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListPartnerRolesResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListPartnerRolesResponseBodyList) SetId(v int64) *ListPartnerRolesResponseBodyList {
	s.Id = &v
	return s
}

func (s *ListPartnerRolesResponseBodyList) SetIsNecessary(v int32) *ListPartnerRolesResponseBodyList {
	s.IsNecessary = &v
	return s
}

func (s *ListPartnerRolesResponseBodyList) SetName(v string) *ListPartnerRolesResponseBodyList {
	s.Name = &v
	return s
}

func (s *ListPartnerRolesResponseBodyList) SetVisibleDepts(v []*ListPartnerRolesResponseBodyListVisibleDepts) *ListPartnerRolesResponseBodyList {
	s.VisibleDepts = v
	return s
}

func (s *ListPartnerRolesResponseBodyList) SetVisibleUsers(v []*ListPartnerRolesResponseBodyListVisibleUsers) *ListPartnerRolesResponseBodyList {
	s.VisibleUsers = v
	return s
}

func (s *ListPartnerRolesResponseBodyList) SetWarningDepts(v []*ListPartnerRolesResponseBodyListWarningDepts) *ListPartnerRolesResponseBodyList {
	s.WarningDepts = v
	return s
}

func (s *ListPartnerRolesResponseBodyList) SetWarningUsers(v []*ListPartnerRolesResponseBodyListWarningUsers) *ListPartnerRolesResponseBodyList {
	s.WarningUsers = v
	return s
}

type ListPartnerRolesResponseBodyListVisibleDepts struct {
	DeptId *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListPartnerRolesResponseBodyListVisibleDepts) String() string {
	return tea.Prettify(s)
}

func (s ListPartnerRolesResponseBodyListVisibleDepts) GoString() string {
	return s.String()
}

func (s *ListPartnerRolesResponseBodyListVisibleDepts) SetDeptId(v int64) *ListPartnerRolesResponseBodyListVisibleDepts {
	s.DeptId = &v
	return s
}

func (s *ListPartnerRolesResponseBodyListVisibleDepts) SetName(v string) *ListPartnerRolesResponseBodyListVisibleDepts {
	s.Name = &v
	return s
}

type ListPartnerRolesResponseBodyListVisibleUsers struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListPartnerRolesResponseBodyListVisibleUsers) String() string {
	return tea.Prettify(s)
}

func (s ListPartnerRolesResponseBodyListVisibleUsers) GoString() string {
	return s.String()
}

func (s *ListPartnerRolesResponseBodyListVisibleUsers) SetName(v string) *ListPartnerRolesResponseBodyListVisibleUsers {
	s.Name = &v
	return s
}

func (s *ListPartnerRolesResponseBodyListVisibleUsers) SetUserId(v string) *ListPartnerRolesResponseBodyListVisibleUsers {
	s.UserId = &v
	return s
}

type ListPartnerRolesResponseBodyListWarningDepts struct {
	DeptId *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListPartnerRolesResponseBodyListWarningDepts) String() string {
	return tea.Prettify(s)
}

func (s ListPartnerRolesResponseBodyListWarningDepts) GoString() string {
	return s.String()
}

func (s *ListPartnerRolesResponseBodyListWarningDepts) SetDeptId(v int64) *ListPartnerRolesResponseBodyListWarningDepts {
	s.DeptId = &v
	return s
}

func (s *ListPartnerRolesResponseBodyListWarningDepts) SetName(v string) *ListPartnerRolesResponseBodyListWarningDepts {
	s.Name = &v
	return s
}

type ListPartnerRolesResponseBodyListWarningUsers struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListPartnerRolesResponseBodyListWarningUsers) String() string {
	return tea.Prettify(s)
}

func (s ListPartnerRolesResponseBodyListWarningUsers) GoString() string {
	return s.String()
}

func (s *ListPartnerRolesResponseBodyListWarningUsers) SetName(v string) *ListPartnerRolesResponseBodyListWarningUsers {
	s.Name = &v
	return s
}

func (s *ListPartnerRolesResponseBodyListWarningUsers) SetUserId(v string) *ListPartnerRolesResponseBodyListWarningUsers {
	s.UserId = &v
	return s
}

type ListPartnerRolesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPartnerRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPartnerRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPartnerRolesResponse) GoString() string {
	return s.String()
}

func (s *ListPartnerRolesResponse) SetHeaders(v map[string]*string) *ListPartnerRolesResponse {
	s.Headers = v
	return s
}

func (s *ListPartnerRolesResponse) SetStatusCode(v int32) *ListPartnerRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPartnerRolesResponse) SetBody(v *ListPartnerRolesResponseBody) *ListPartnerRolesResponse {
	s.Body = v
	return s
}

type ListPunchScheduleByConditionWithPagingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListPunchScheduleByConditionWithPagingHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListPunchScheduleByConditionWithPagingHeaders) GoString() string {
	return s.String()
}

func (s *ListPunchScheduleByConditionWithPagingHeaders) SetCommonHeaders(v map[string]*string) *ListPunchScheduleByConditionWithPagingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListPunchScheduleByConditionWithPagingHeaders) SetXAcsDingtalkAccessToken(v string) *ListPunchScheduleByConditionWithPagingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListPunchScheduleByConditionWithPagingRequest struct {
	BizInstanceId     *string   `json:"bizInstanceId,omitempty" xml:"bizInstanceId,omitempty"`
	MaxResults        *int32    `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken         *int64    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ScheduleDateEnd   *string   `json:"scheduleDateEnd,omitempty" xml:"scheduleDateEnd,omitempty"`
	ScheduleDateStart *string   `json:"scheduleDateStart,omitempty" xml:"scheduleDateStart,omitempty"`
	UserIdList        []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s ListPunchScheduleByConditionWithPagingRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPunchScheduleByConditionWithPagingRequest) GoString() string {
	return s.String()
}

func (s *ListPunchScheduleByConditionWithPagingRequest) SetBizInstanceId(v string) *ListPunchScheduleByConditionWithPagingRequest {
	s.BizInstanceId = &v
	return s
}

func (s *ListPunchScheduleByConditionWithPagingRequest) SetMaxResults(v int32) *ListPunchScheduleByConditionWithPagingRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPunchScheduleByConditionWithPagingRequest) SetNextToken(v int64) *ListPunchScheduleByConditionWithPagingRequest {
	s.NextToken = &v
	return s
}

func (s *ListPunchScheduleByConditionWithPagingRequest) SetScheduleDateEnd(v string) *ListPunchScheduleByConditionWithPagingRequest {
	s.ScheduleDateEnd = &v
	return s
}

func (s *ListPunchScheduleByConditionWithPagingRequest) SetScheduleDateStart(v string) *ListPunchScheduleByConditionWithPagingRequest {
	s.ScheduleDateStart = &v
	return s
}

func (s *ListPunchScheduleByConditionWithPagingRequest) SetUserIdList(v []*string) *ListPunchScheduleByConditionWithPagingRequest {
	s.UserIdList = v
	return s
}

type ListPunchScheduleByConditionWithPagingResponseBody struct {
	HasMore   *bool                                                     `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List      []*ListPunchScheduleByConditionWithPagingResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextToken *int64                                                    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListPunchScheduleByConditionWithPagingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPunchScheduleByConditionWithPagingResponseBody) GoString() string {
	return s.String()
}

func (s *ListPunchScheduleByConditionWithPagingResponseBody) SetHasMore(v bool) *ListPunchScheduleByConditionWithPagingResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListPunchScheduleByConditionWithPagingResponseBody) SetList(v []*ListPunchScheduleByConditionWithPagingResponseBodyList) *ListPunchScheduleByConditionWithPagingResponseBody {
	s.List = v
	return s
}

func (s *ListPunchScheduleByConditionWithPagingResponseBody) SetNextToken(v int64) *ListPunchScheduleByConditionWithPagingResponseBody {
	s.NextToken = &v
	return s
}

type ListPunchScheduleByConditionWithPagingResponseBodyList struct {
	BizOuterId    *string `json:"bizOuterId,omitempty" xml:"bizOuterId,omitempty"`
	PositionName  *string `json:"positionName,omitempty" xml:"positionName,omitempty"`
	PunchSymbol   *string `json:"punchSymbol,omitempty" xml:"punchSymbol,omitempty"`
	UserId        *string `json:"userId,omitempty" xml:"userId,omitempty"`
	UserPunchTime *int64  `json:"userPunchTime,omitempty" xml:"userPunchTime,omitempty"`
}

func (s ListPunchScheduleByConditionWithPagingResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListPunchScheduleByConditionWithPagingResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListPunchScheduleByConditionWithPagingResponseBodyList) SetBizOuterId(v string) *ListPunchScheduleByConditionWithPagingResponseBodyList {
	s.BizOuterId = &v
	return s
}

func (s *ListPunchScheduleByConditionWithPagingResponseBodyList) SetPositionName(v string) *ListPunchScheduleByConditionWithPagingResponseBodyList {
	s.PositionName = &v
	return s
}

func (s *ListPunchScheduleByConditionWithPagingResponseBodyList) SetPunchSymbol(v string) *ListPunchScheduleByConditionWithPagingResponseBodyList {
	s.PunchSymbol = &v
	return s
}

func (s *ListPunchScheduleByConditionWithPagingResponseBodyList) SetUserId(v string) *ListPunchScheduleByConditionWithPagingResponseBodyList {
	s.UserId = &v
	return s
}

func (s *ListPunchScheduleByConditionWithPagingResponseBodyList) SetUserPunchTime(v int64) *ListPunchScheduleByConditionWithPagingResponseBodyList {
	s.UserPunchTime = &v
	return s
}

type ListPunchScheduleByConditionWithPagingResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPunchScheduleByConditionWithPagingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPunchScheduleByConditionWithPagingResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPunchScheduleByConditionWithPagingResponse) GoString() string {
	return s.String()
}

func (s *ListPunchScheduleByConditionWithPagingResponse) SetHeaders(v map[string]*string) *ListPunchScheduleByConditionWithPagingResponse {
	s.Headers = v
	return s
}

func (s *ListPunchScheduleByConditionWithPagingResponse) SetStatusCode(v int32) *ListPunchScheduleByConditionWithPagingResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPunchScheduleByConditionWithPagingResponse) SetBody(v *ListPunchScheduleByConditionWithPagingResponseBody) *ListPunchScheduleByConditionWithPagingResponse {
	s.Body = v
	return s
}

type ListRulesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListRulesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListRulesHeaders) GoString() string {
	return s.String()
}

func (s *ListRulesHeaders) SetCommonHeaders(v map[string]*string) *ListRulesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListRulesHeaders) SetXAcsDingtalkAccessToken(v string) *ListRulesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListRulesRequest struct {
	Body *ListRulesRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s ListRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRulesRequest) GoString() string {
	return s.String()
}

func (s *ListRulesRequest) SetBody(v *ListRulesRequestBody) *ListRulesRequest {
	s.Body = v
	return s
}

type ListRulesRequestBody struct {
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListRulesRequestBody) String() string {
	return tea.Prettify(s)
}

func (s ListRulesRequestBody) GoString() string {
	return s.String()
}

func (s *ListRulesRequestBody) SetStatus(v int64) *ListRulesRequestBody {
	s.Status = &v
	return s
}

type ListRulesShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRulesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListRulesShrinkRequest) SetBodyShrink(v string) *ListRulesShrinkRequest {
	s.BodyShrink = &v
	return s
}

type ListRulesResponseBody struct {
	DetailModelList []map[string]*string `json:"detailModelList,omitempty" xml:"detailModelList,omitempty" type:"Repeated"`
}

func (s ListRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBody) SetDetailModelList(v []map[string]*string) *ListRulesResponseBody {
	s.DetailModelList = v
	return s
}

type ListRulesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponse) GoString() string {
	return s.String()
}

func (s *ListRulesResponse) SetHeaders(v map[string]*string) *ListRulesResponse {
	s.Headers = v
	return s
}

func (s *ListRulesResponse) SetStatusCode(v int32) *ListRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRulesResponse) SetBody(v *ListRulesResponseBody) *ListRulesResponse {
	s.Body = v
	return s
}

type LogoutHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s LogoutHeaders) String() string {
	return tea.Prettify(s)
}

func (s LogoutHeaders) GoString() string {
	return s.String()
}

func (s *LogoutHeaders) SetCommonHeaders(v map[string]*string) *LogoutHeaders {
	s.CommonHeaders = v
	return s
}

func (s *LogoutHeaders) SetXAcsDingtalkAccessToken(v string) *LogoutHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type LogoutRequest struct {
	DeviceType *string `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
	UserId     *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s LogoutRequest) String() string {
	return tea.Prettify(s)
}

func (s LogoutRequest) GoString() string {
	return s.String()
}

func (s *LogoutRequest) SetDeviceType(v string) *LogoutRequest {
	s.DeviceType = &v
	return s
}

func (s *LogoutRequest) SetUserId(v string) *LogoutRequest {
	s.UserId = &v
	return s
}

type LogoutResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s LogoutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LogoutResponseBody) GoString() string {
	return s.String()
}

func (s *LogoutResponseBody) SetSuccess(v bool) *LogoutResponseBody {
	s.Success = &v
	return s
}

type LogoutResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LogoutResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LogoutResponse) String() string {
	return tea.Prettify(s)
}

func (s LogoutResponse) GoString() string {
	return s.String()
}

func (s *LogoutResponse) SetHeaders(v map[string]*string) *LogoutResponse {
	s.Headers = v
	return s
}

func (s *LogoutResponse) SetStatusCode(v int32) *LogoutResponse {
	s.StatusCode = &v
	return s
}

func (s *LogoutResponse) SetBody(v *LogoutResponseBody) *LogoutResponse {
	s.Body = v
	return s
}

type PublishFileChangeNoticeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PublishFileChangeNoticeHeaders) String() string {
	return tea.Prettify(s)
}

func (s PublishFileChangeNoticeHeaders) GoString() string {
	return s.String()
}

func (s *PublishFileChangeNoticeHeaders) SetCommonHeaders(v map[string]*string) *PublishFileChangeNoticeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PublishFileChangeNoticeHeaders) SetXAcsDingtalkAccessToken(v string) *PublishFileChangeNoticeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PublishFileChangeNoticeRequest struct {
	FileId          *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	OperateType     *string `json:"operateType,omitempty" xml:"operateType,omitempty"`
	OperatorUnionId *string `json:"operatorUnionId,omitempty" xml:"operatorUnionId,omitempty"`
	SpaceId         *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s PublishFileChangeNoticeRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishFileChangeNoticeRequest) GoString() string {
	return s.String()
}

func (s *PublishFileChangeNoticeRequest) SetFileId(v string) *PublishFileChangeNoticeRequest {
	s.FileId = &v
	return s
}

func (s *PublishFileChangeNoticeRequest) SetOperateType(v string) *PublishFileChangeNoticeRequest {
	s.OperateType = &v
	return s
}

func (s *PublishFileChangeNoticeRequest) SetOperatorUnionId(v string) *PublishFileChangeNoticeRequest {
	s.OperatorUnionId = &v
	return s
}

func (s *PublishFileChangeNoticeRequest) SetSpaceId(v string) *PublishFileChangeNoticeRequest {
	s.SpaceId = &v
	return s
}

type PublishFileChangeNoticeResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PublishFileChangeNoticeResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishFileChangeNoticeResponse) GoString() string {
	return s.String()
}

func (s *PublishFileChangeNoticeResponse) SetHeaders(v map[string]*string) *PublishFileChangeNoticeResponse {
	s.Headers = v
	return s
}

func (s *PublishFileChangeNoticeResponse) SetStatusCode(v int32) *PublishFileChangeNoticeResponse {
	s.StatusCode = &v
	return s
}

type PublishRuleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PublishRuleHeaders) String() string {
	return tea.Prettify(s)
}

func (s PublishRuleHeaders) GoString() string {
	return s.String()
}

func (s *PublishRuleHeaders) SetCommonHeaders(v map[string]*string) *PublishRuleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PublishRuleHeaders) SetXAcsDingtalkAccessToken(v string) *PublishRuleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PublishRuleRequest struct {
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PublishRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishRuleRequest) GoString() string {
	return s.String()
}

func (s *PublishRuleRequest) SetStatus(v int64) *PublishRuleRequest {
	s.Status = &v
	return s
}

type PublishRuleResponseBody struct {
	IsPublish *bool `json:"isPublish,omitempty" xml:"isPublish,omitempty"`
}

func (s PublishRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishRuleResponseBody) GoString() string {
	return s.String()
}

func (s *PublishRuleResponseBody) SetIsPublish(v bool) *PublishRuleResponseBody {
	s.IsPublish = &v
	return s
}

type PublishRuleResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PublishRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PublishRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishRuleResponse) GoString() string {
	return s.String()
}

func (s *PublishRuleResponse) SetHeaders(v map[string]*string) *PublishRuleResponse {
	s.Headers = v
	return s
}

func (s *PublishRuleResponse) SetStatusCode(v int32) *PublishRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishRuleResponse) SetBody(v *PublishRuleResponseBody) *PublishRuleResponse {
	s.Body = v
	return s
}

type PushBadgeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PushBadgeHeaders) String() string {
	return tea.Prettify(s)
}

func (s PushBadgeHeaders) GoString() string {
	return s.String()
}

func (s *PushBadgeHeaders) SetCommonHeaders(v map[string]*string) *PushBadgeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PushBadgeHeaders) SetXAcsDingtalkAccessToken(v string) *PushBadgeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PushBadgeRequest struct {
	AgentId    *string                       `json:"agentId,omitempty" xml:"agentId,omitempty"`
	BadgeItems []*PushBadgeRequestBadgeItems `json:"badgeItems,omitempty" xml:"badgeItems,omitempty" type:"Repeated"`
	PushType   *string                       `json:"pushType,omitempty" xml:"pushType,omitempty"`
}

func (s PushBadgeRequest) String() string {
	return tea.Prettify(s)
}

func (s PushBadgeRequest) GoString() string {
	return s.String()
}

func (s *PushBadgeRequest) SetAgentId(v string) *PushBadgeRequest {
	s.AgentId = &v
	return s
}

func (s *PushBadgeRequest) SetBadgeItems(v []*PushBadgeRequestBadgeItems) *PushBadgeRequest {
	s.BadgeItems = v
	return s
}

func (s *PushBadgeRequest) SetPushType(v string) *PushBadgeRequest {
	s.PushType = &v
	return s
}

type PushBadgeRequestBadgeItems struct {
	PushValue *string `json:"pushValue,omitempty" xml:"pushValue,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PushBadgeRequestBadgeItems) String() string {
	return tea.Prettify(s)
}

func (s PushBadgeRequestBadgeItems) GoString() string {
	return s.String()
}

func (s *PushBadgeRequestBadgeItems) SetPushValue(v string) *PushBadgeRequestBadgeItems {
	s.PushValue = &v
	return s
}

func (s *PushBadgeRequestBadgeItems) SetUserId(v string) *PushBadgeRequestBadgeItems {
	s.UserId = &v
	return s
}

type PushBadgeResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PushBadgeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushBadgeResponseBody) GoString() string {
	return s.String()
}

func (s *PushBadgeResponseBody) SetSuccess(v bool) *PushBadgeResponseBody {
	s.Success = &v
	return s
}

type PushBadgeResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PushBadgeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushBadgeResponse) String() string {
	return tea.Prettify(s)
}

func (s PushBadgeResponse) GoString() string {
	return s.String()
}

func (s *PushBadgeResponse) SetHeaders(v map[string]*string) *PushBadgeResponse {
	s.Headers = v
	return s
}

func (s *PushBadgeResponse) SetStatusCode(v int32) *PushBadgeResponse {
	s.StatusCode = &v
	return s
}

func (s *PushBadgeResponse) SetBody(v *PushBadgeResponseBody) *PushBadgeResponse {
	s.Body = v
	return s
}

type QueryAcrossCloudStroageConfigsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAcrossCloudStroageConfigsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAcrossCloudStroageConfigsHeaders) GoString() string {
	return s.String()
}

func (s *QueryAcrossCloudStroageConfigsHeaders) SetCommonHeaders(v map[string]*string) *QueryAcrossCloudStroageConfigsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAcrossCloudStroageConfigsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAcrossCloudStroageConfigsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAcrossCloudStroageConfigsRequest struct {
	TargetCloudType *int32  `json:"targetCloudType,omitempty" xml:"targetCloudType,omitempty"`
	TargetCorpId    *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
}

func (s QueryAcrossCloudStroageConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAcrossCloudStroageConfigsRequest) GoString() string {
	return s.String()
}

func (s *QueryAcrossCloudStroageConfigsRequest) SetTargetCloudType(v int32) *QueryAcrossCloudStroageConfigsRequest {
	s.TargetCloudType = &v
	return s
}

func (s *QueryAcrossCloudStroageConfigsRequest) SetTargetCorpId(v string) *QueryAcrossCloudStroageConfigsRequest {
	s.TargetCorpId = &v
	return s
}

type QueryAcrossCloudStroageConfigsResponseBody struct {
	AccessKeyId     *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	BucketName      *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	CloudType       *int32  `json:"cloudType,omitempty" xml:"cloudType,omitempty"`
	Endpoint        *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
}

func (s QueryAcrossCloudStroageConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAcrossCloudStroageConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAcrossCloudStroageConfigsResponseBody) SetAccessKeyId(v string) *QueryAcrossCloudStroageConfigsResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *QueryAcrossCloudStroageConfigsResponseBody) SetAccessKeySecret(v string) *QueryAcrossCloudStroageConfigsResponseBody {
	s.AccessKeySecret = &v
	return s
}

func (s *QueryAcrossCloudStroageConfigsResponseBody) SetBucketName(v string) *QueryAcrossCloudStroageConfigsResponseBody {
	s.BucketName = &v
	return s
}

func (s *QueryAcrossCloudStroageConfigsResponseBody) SetCloudType(v int32) *QueryAcrossCloudStroageConfigsResponseBody {
	s.CloudType = &v
	return s
}

func (s *QueryAcrossCloudStroageConfigsResponseBody) SetEndpoint(v string) *QueryAcrossCloudStroageConfigsResponseBody {
	s.Endpoint = &v
	return s
}

type QueryAcrossCloudStroageConfigsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryAcrossCloudStroageConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAcrossCloudStroageConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAcrossCloudStroageConfigsResponse) GoString() string {
	return s.String()
}

func (s *QueryAcrossCloudStroageConfigsResponse) SetHeaders(v map[string]*string) *QueryAcrossCloudStroageConfigsResponse {
	s.Headers = v
	return s
}

func (s *QueryAcrossCloudStroageConfigsResponse) SetStatusCode(v int32) *QueryAcrossCloudStroageConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAcrossCloudStroageConfigsResponse) SetBody(v *QueryAcrossCloudStroageConfigsResponseBody) *QueryAcrossCloudStroageConfigsResponse {
	s.Body = v
	return s
}

type QueryPartnerInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryPartnerInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryPartnerInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryPartnerInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryPartnerInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryPartnerInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryPartnerInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryPartnerInfoResponseBody struct {
	PartnerDeptList  []*QueryPartnerInfoResponseBodyPartnerDeptList  `json:"partnerDeptList,omitempty" xml:"partnerDeptList,omitempty" type:"Repeated"`
	PartnerLabelList []*QueryPartnerInfoResponseBodyPartnerLabelList `json:"partnerLabelList,omitempty" xml:"partnerLabelList,omitempty" type:"Repeated"`
	UserId           *string                                         `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryPartnerInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPartnerInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPartnerInfoResponseBody) SetPartnerDeptList(v []*QueryPartnerInfoResponseBodyPartnerDeptList) *QueryPartnerInfoResponseBody {
	s.PartnerDeptList = v
	return s
}

func (s *QueryPartnerInfoResponseBody) SetPartnerLabelList(v []*QueryPartnerInfoResponseBodyPartnerLabelList) *QueryPartnerInfoResponseBody {
	s.PartnerLabelList = v
	return s
}

func (s *QueryPartnerInfoResponseBody) SetUserId(v string) *QueryPartnerInfoResponseBody {
	s.UserId = &v
	return s
}

type QueryPartnerInfoResponseBodyPartnerDeptList struct {
	MemberCount             *int64                                                              `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	PartnerLabelModelLevel1 *QueryPartnerInfoResponseBodyPartnerDeptListPartnerLabelModelLevel1 `json:"partnerLabelModelLevel1,omitempty" xml:"partnerLabelModelLevel1,omitempty" type:"Struct"`
	PartnerNum              *string                                                             `json:"partnerNum,omitempty" xml:"partnerNum,omitempty"`
	Title                   *string                                                             `json:"title,omitempty" xml:"title,omitempty"`
	Value                   *string                                                             `json:"value,omitempty" xml:"value,omitempty"`
}

func (s QueryPartnerInfoResponseBodyPartnerDeptList) String() string {
	return tea.Prettify(s)
}

func (s QueryPartnerInfoResponseBodyPartnerDeptList) GoString() string {
	return s.String()
}

func (s *QueryPartnerInfoResponseBodyPartnerDeptList) SetMemberCount(v int64) *QueryPartnerInfoResponseBodyPartnerDeptList {
	s.MemberCount = &v
	return s
}

func (s *QueryPartnerInfoResponseBodyPartnerDeptList) SetPartnerLabelModelLevel1(v *QueryPartnerInfoResponseBodyPartnerDeptListPartnerLabelModelLevel1) *QueryPartnerInfoResponseBodyPartnerDeptList {
	s.PartnerLabelModelLevel1 = v
	return s
}

func (s *QueryPartnerInfoResponseBodyPartnerDeptList) SetPartnerNum(v string) *QueryPartnerInfoResponseBodyPartnerDeptList {
	s.PartnerNum = &v
	return s
}

func (s *QueryPartnerInfoResponseBodyPartnerDeptList) SetTitle(v string) *QueryPartnerInfoResponseBodyPartnerDeptList {
	s.Title = &v
	return s
}

func (s *QueryPartnerInfoResponseBodyPartnerDeptList) SetValue(v string) *QueryPartnerInfoResponseBodyPartnerDeptList {
	s.Value = &v
	return s
}

type QueryPartnerInfoResponseBodyPartnerDeptListPartnerLabelModelLevel1 struct {
	LabelId   *int64  `json:"labelId,omitempty" xml:"labelId,omitempty"`
	Labelname *string `json:"labelname,omitempty" xml:"labelname,omitempty"`
}

func (s QueryPartnerInfoResponseBodyPartnerDeptListPartnerLabelModelLevel1) String() string {
	return tea.Prettify(s)
}

func (s QueryPartnerInfoResponseBodyPartnerDeptListPartnerLabelModelLevel1) GoString() string {
	return s.String()
}

func (s *QueryPartnerInfoResponseBodyPartnerDeptListPartnerLabelModelLevel1) SetLabelId(v int64) *QueryPartnerInfoResponseBodyPartnerDeptListPartnerLabelModelLevel1 {
	s.LabelId = &v
	return s
}

func (s *QueryPartnerInfoResponseBodyPartnerDeptListPartnerLabelModelLevel1) SetLabelname(v string) *QueryPartnerInfoResponseBodyPartnerDeptListPartnerLabelModelLevel1 {
	s.Labelname = &v
	return s
}

type QueryPartnerInfoResponseBodyPartnerLabelList struct {
	Id   *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryPartnerInfoResponseBodyPartnerLabelList) String() string {
	return tea.Prettify(s)
}

func (s QueryPartnerInfoResponseBodyPartnerLabelList) GoString() string {
	return s.String()
}

func (s *QueryPartnerInfoResponseBodyPartnerLabelList) SetId(v int64) *QueryPartnerInfoResponseBodyPartnerLabelList {
	s.Id = &v
	return s
}

func (s *QueryPartnerInfoResponseBodyPartnerLabelList) SetName(v string) *QueryPartnerInfoResponseBodyPartnerLabelList {
	s.Name = &v
	return s
}

type QueryPartnerInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryPartnerInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPartnerInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPartnerInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryPartnerInfoResponse) SetHeaders(v map[string]*string) *QueryPartnerInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryPartnerInfoResponse) SetStatusCode(v int32) *QueryPartnerInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPartnerInfoResponse) SetBody(v *QueryPartnerInfoResponseBody) *QueryPartnerInfoResponse {
	s.Body = v
	return s
}

type QueryUserBehaviorHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserBehaviorHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserBehaviorHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserBehaviorHeaders) SetCommonHeaders(v map[string]*string) *QueryUserBehaviorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserBehaviorHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserBehaviorHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserBehaviorRequest struct {
	EndTime    *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	PageNumber *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Platform   *int32  `json:"platform,omitempty" xml:"platform,omitempty"`
	StartTime  *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Type       *int32  `json:"type,omitempty" xml:"type,omitempty"`
	UserId     *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryUserBehaviorRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserBehaviorRequest) GoString() string {
	return s.String()
}

func (s *QueryUserBehaviorRequest) SetEndTime(v int64) *QueryUserBehaviorRequest {
	s.EndTime = &v
	return s
}

func (s *QueryUserBehaviorRequest) SetPageNumber(v int64) *QueryUserBehaviorRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryUserBehaviorRequest) SetPageSize(v int32) *QueryUserBehaviorRequest {
	s.PageSize = &v
	return s
}

func (s *QueryUserBehaviorRequest) SetPlatform(v int32) *QueryUserBehaviorRequest {
	s.Platform = &v
	return s
}

func (s *QueryUserBehaviorRequest) SetStartTime(v int64) *QueryUserBehaviorRequest {
	s.StartTime = &v
	return s
}

func (s *QueryUserBehaviorRequest) SetType(v int32) *QueryUserBehaviorRequest {
	s.Type = &v
	return s
}

func (s *QueryUserBehaviorRequest) SetUserId(v string) *QueryUserBehaviorRequest {
	s.UserId = &v
	return s
}

type QueryUserBehaviorResponseBody struct {
	Data     []*QueryUserBehaviorResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	DataCnt  *int32                               `json:"dataCnt,omitempty" xml:"dataCnt,omitempty"`
	TotalCnt *int32                               `json:"totalCnt,omitempty" xml:"totalCnt,omitempty"`
}

func (s QueryUserBehaviorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserBehaviorResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserBehaviorResponseBody) SetData(v []*QueryUserBehaviorResponseBodyData) *QueryUserBehaviorResponseBody {
	s.Data = v
	return s
}

func (s *QueryUserBehaviorResponseBody) SetDataCnt(v int32) *QueryUserBehaviorResponseBody {
	s.DataCnt = &v
	return s
}

func (s *QueryUserBehaviorResponseBody) SetTotalCnt(v int32) *QueryUserBehaviorResponseBody {
	s.TotalCnt = &v
	return s
}

type QueryUserBehaviorResponseBodyData struct {
	PictureUrl *string `json:"pictureUrl,omitempty" xml:"pictureUrl,omitempty"`
	Platform   *int32  `json:"platform,omitempty" xml:"platform,omitempty"`
	Scene      *string `json:"scene,omitempty" xml:"scene,omitempty"`
	Time       *int64  `json:"time,omitempty" xml:"time,omitempty"`
	Type       *int32  `json:"type,omitempty" xml:"type,omitempty"`
	UserId     *string `json:"userId,omitempty" xml:"userId,omitempty"`
	UserName   *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s QueryUserBehaviorResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryUserBehaviorResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryUserBehaviorResponseBodyData) SetPictureUrl(v string) *QueryUserBehaviorResponseBodyData {
	s.PictureUrl = &v
	return s
}

func (s *QueryUserBehaviorResponseBodyData) SetPlatform(v int32) *QueryUserBehaviorResponseBodyData {
	s.Platform = &v
	return s
}

func (s *QueryUserBehaviorResponseBodyData) SetScene(v string) *QueryUserBehaviorResponseBodyData {
	s.Scene = &v
	return s
}

func (s *QueryUserBehaviorResponseBodyData) SetTime(v int64) *QueryUserBehaviorResponseBodyData {
	s.Time = &v
	return s
}

func (s *QueryUserBehaviorResponseBodyData) SetType(v int32) *QueryUserBehaviorResponseBodyData {
	s.Type = &v
	return s
}

func (s *QueryUserBehaviorResponseBodyData) SetUserId(v string) *QueryUserBehaviorResponseBodyData {
	s.UserId = &v
	return s
}

func (s *QueryUserBehaviorResponseBodyData) SetUserName(v string) *QueryUserBehaviorResponseBodyData {
	s.UserName = &v
	return s
}

type QueryUserBehaviorResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUserBehaviorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserBehaviorResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserBehaviorResponse) GoString() string {
	return s.String()
}

func (s *QueryUserBehaviorResponse) SetHeaders(v map[string]*string) *QueryUserBehaviorResponse {
	s.Headers = v
	return s
}

func (s *QueryUserBehaviorResponse) SetStatusCode(v int32) *QueryUserBehaviorResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserBehaviorResponse) SetBody(v *QueryUserBehaviorResponseBody) *QueryUserBehaviorResponse {
	s.Body = v
	return s
}

type RollbackMiniAppVersionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RollbackMiniAppVersionHeaders) String() string {
	return tea.Prettify(s)
}

func (s RollbackMiniAppVersionHeaders) GoString() string {
	return s.String()
}

func (s *RollbackMiniAppVersionHeaders) SetCommonHeaders(v map[string]*string) *RollbackMiniAppVersionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RollbackMiniAppVersionHeaders) SetXAcsDingtalkAccessToken(v string) *RollbackMiniAppVersionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RollbackMiniAppVersionRequest struct {
	MiniAppId       *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	RollbackVersion *string `json:"rollbackVersion,omitempty" xml:"rollbackVersion,omitempty"`
	TargetVersion   *string `json:"targetVersion,omitempty" xml:"targetVersion,omitempty"`
}

func (s RollbackMiniAppVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s RollbackMiniAppVersionRequest) GoString() string {
	return s.String()
}

func (s *RollbackMiniAppVersionRequest) SetMiniAppId(v string) *RollbackMiniAppVersionRequest {
	s.MiniAppId = &v
	return s
}

func (s *RollbackMiniAppVersionRequest) SetRollbackVersion(v string) *RollbackMiniAppVersionRequest {
	s.RollbackVersion = &v
	return s
}

func (s *RollbackMiniAppVersionRequest) SetTargetVersion(v string) *RollbackMiniAppVersionRequest {
	s.TargetVersion = &v
	return s
}

type RollbackMiniAppVersionResponseBody struct {
	Cause *string `json:"cause,omitempty" xml:"cause,omitempty"`
	Code  *int64  `json:"code,omitempty" xml:"code,omitempty"`
}

func (s RollbackMiniAppVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RollbackMiniAppVersionResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackMiniAppVersionResponseBody) SetCause(v string) *RollbackMiniAppVersionResponseBody {
	s.Cause = &v
	return s
}

func (s *RollbackMiniAppVersionResponseBody) SetCode(v int64) *RollbackMiniAppVersionResponseBody {
	s.Code = &v
	return s
}

type RollbackMiniAppVersionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RollbackMiniAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RollbackMiniAppVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s RollbackMiniAppVersionResponse) GoString() string {
	return s.String()
}

func (s *RollbackMiniAppVersionResponse) SetHeaders(v map[string]*string) *RollbackMiniAppVersionResponse {
	s.Headers = v
	return s
}

func (s *RollbackMiniAppVersionResponse) SetStatusCode(v int32) *RollbackMiniAppVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackMiniAppVersionResponse) SetBody(v *RollbackMiniAppVersionResponseBody) *RollbackMiniAppVersionResponse {
	s.Body = v
	return s
}

type SaveAcrossCloudStroageConfigsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SaveAcrossCloudStroageConfigsHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveAcrossCloudStroageConfigsHeaders) GoString() string {
	return s.String()
}

func (s *SaveAcrossCloudStroageConfigsHeaders) SetCommonHeaders(v map[string]*string) *SaveAcrossCloudStroageConfigsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveAcrossCloudStroageConfigsHeaders) SetXAcsDingtalkAccessToken(v string) *SaveAcrossCloudStroageConfigsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SaveAcrossCloudStroageConfigsRequest struct {
	AccessKeyId     *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	BucketName      *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	CloudType       *int32  `json:"cloudType,omitempty" xml:"cloudType,omitempty"`
	Endpoint        *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	TargetCorpId    *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
}

func (s SaveAcrossCloudStroageConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveAcrossCloudStroageConfigsRequest) GoString() string {
	return s.String()
}

func (s *SaveAcrossCloudStroageConfigsRequest) SetAccessKeyId(v string) *SaveAcrossCloudStroageConfigsRequest {
	s.AccessKeyId = &v
	return s
}

func (s *SaveAcrossCloudStroageConfigsRequest) SetAccessKeySecret(v string) *SaveAcrossCloudStroageConfigsRequest {
	s.AccessKeySecret = &v
	return s
}

func (s *SaveAcrossCloudStroageConfigsRequest) SetBucketName(v string) *SaveAcrossCloudStroageConfigsRequest {
	s.BucketName = &v
	return s
}

func (s *SaveAcrossCloudStroageConfigsRequest) SetCloudType(v int32) *SaveAcrossCloudStroageConfigsRequest {
	s.CloudType = &v
	return s
}

func (s *SaveAcrossCloudStroageConfigsRequest) SetEndpoint(v string) *SaveAcrossCloudStroageConfigsRequest {
	s.Endpoint = &v
	return s
}

func (s *SaveAcrossCloudStroageConfigsRequest) SetTargetCorpId(v string) *SaveAcrossCloudStroageConfigsRequest {
	s.TargetCorpId = &v
	return s
}

type SaveAcrossCloudStroageConfigsResponseBody struct {
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	Endpoint    *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	State       *int32  `json:"state,omitempty" xml:"state,omitempty"`
}

func (s SaveAcrossCloudStroageConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveAcrossCloudStroageConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *SaveAcrossCloudStroageConfigsResponseBody) SetAccessKeyId(v string) *SaveAcrossCloudStroageConfigsResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *SaveAcrossCloudStroageConfigsResponseBody) SetEndpoint(v string) *SaveAcrossCloudStroageConfigsResponseBody {
	s.Endpoint = &v
	return s
}

func (s *SaveAcrossCloudStroageConfigsResponseBody) SetState(v int32) *SaveAcrossCloudStroageConfigsResponseBody {
	s.State = &v
	return s
}

type SaveAcrossCloudStroageConfigsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveAcrossCloudStroageConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveAcrossCloudStroageConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveAcrossCloudStroageConfigsResponse) GoString() string {
	return s.String()
}

func (s *SaveAcrossCloudStroageConfigsResponse) SetHeaders(v map[string]*string) *SaveAcrossCloudStroageConfigsResponse {
	s.Headers = v
	return s
}

func (s *SaveAcrossCloudStroageConfigsResponse) SetStatusCode(v int32) *SaveAcrossCloudStroageConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveAcrossCloudStroageConfigsResponse) SetBody(v *SaveAcrossCloudStroageConfigsResponseBody) *SaveAcrossCloudStroageConfigsResponse {
	s.Body = v
	return s
}

type SaveAndSubmitAuthInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SaveAndSubmitAuthInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveAndSubmitAuthInfoHeaders) GoString() string {
	return s.String()
}

func (s *SaveAndSubmitAuthInfoHeaders) SetCommonHeaders(v map[string]*string) *SaveAndSubmitAuthInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveAndSubmitAuthInfoHeaders) SetXAcsDingtalkAccessToken(v string) *SaveAndSubmitAuthInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SaveAndSubmitAuthInfoRequest struct {
	ApplyRemark             *string `json:"applyRemark,omitempty" xml:"applyRemark,omitempty"`
	AuthorizeMediaId        *string `json:"authorizeMediaId,omitempty" xml:"authorizeMediaId,omitempty"`
	Industry                *string `json:"industry,omitempty" xml:"industry,omitempty"`
	LegalPerson             *string `json:"legalPerson,omitempty" xml:"legalPerson,omitempty"`
	LegalPersonIdCard       *string `json:"legalPersonIdCard,omitempty" xml:"legalPersonIdCard,omitempty"`
	LicenseMediaId          *string `json:"licenseMediaId,omitempty" xml:"licenseMediaId,omitempty"`
	LocCity                 *int64  `json:"locCity,omitempty" xml:"locCity,omitempty"`
	LocCityName             *string `json:"locCityName,omitempty" xml:"locCityName,omitempty"`
	LocProvince             *int64  `json:"locProvince,omitempty" xml:"locProvince,omitempty"`
	LocProvinceName         *string `json:"locProvinceName,omitempty" xml:"locProvinceName,omitempty"`
	MobileNum               *string `json:"mobileNum,omitempty" xml:"mobileNum,omitempty"`
	OrgName                 *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	OrganizationCode        *string `json:"organizationCode,omitempty" xml:"organizationCode,omitempty"`
	OrganizationCodeMediaId *string `json:"organizationCodeMediaId,omitempty" xml:"organizationCodeMediaId,omitempty"`
	RegistLocation          *string `json:"registLocation,omitempty" xml:"registLocation,omitempty"`
	RegistNum               *string `json:"registNum,omitempty" xml:"registNum,omitempty"`
	TargetCorpId            *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
	UnifiedSocialCredit     *string `json:"unifiedSocialCredit,omitempty" xml:"unifiedSocialCredit,omitempty"`
}

func (s SaveAndSubmitAuthInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveAndSubmitAuthInfoRequest) GoString() string {
	return s.String()
}

func (s *SaveAndSubmitAuthInfoRequest) SetApplyRemark(v string) *SaveAndSubmitAuthInfoRequest {
	s.ApplyRemark = &v
	return s
}

func (s *SaveAndSubmitAuthInfoRequest) SetAuthorizeMediaId(v string) *SaveAndSubmitAuthInfoRequest {
	s.AuthorizeMediaId = &v
	return s
}

func (s *SaveAndSubmitAuthInfoRequest) SetIndustry(v string) *SaveAndSubmitAuthInfoRequest {
	s.Industry = &v
	return s
}

func (s *SaveAndSubmitAuthInfoRequest) SetLegalPerson(v string) *SaveAndSubmitAuthInfoRequest {
	s.LegalPerson = &v
	return s
}

func (s *SaveAndSubmitAuthInfoRequest) SetLegalPersonIdCard(v string) *SaveAndSubmitAuthInfoRequest {
	s.LegalPersonIdCard = &v
	return s
}

func (s *SaveAndSubmitAuthInfoRequest) SetLicenseMediaId(v string) *SaveAndSubmitAuthInfoRequest {
	s.LicenseMediaId = &v
	return s
}

func (s *SaveAndSubmitAuthInfoRequest) SetLocCity(v int64) *SaveAndSubmitAuthInfoRequest {
	s.LocCity = &v
	return s
}

func (s *SaveAndSubmitAuthInfoRequest) SetLocCityName(v string) *SaveAndSubmitAuthInfoRequest {
	s.LocCityName = &v
	return s
}

func (s *SaveAndSubmitAuthInfoRequest) SetLocProvince(v int64) *SaveAndSubmitAuthInfoRequest {
	s.LocProvince = &v
	return s
}

func (s *SaveAndSubmitAuthInfoRequest) SetLocProvinceName(v string) *SaveAndSubmitAuthInfoRequest {
	s.LocProvinceName = &v
	return s
}

func (s *SaveAndSubmitAuthInfoRequest) SetMobileNum(v string) *SaveAndSubmitAuthInfoRequest {
	s.MobileNum = &v
	return s
}

func (s *SaveAndSubmitAuthInfoRequest) SetOrgName(v string) *SaveAndSubmitAuthInfoRequest {
	s.OrgName = &v
	return s
}

func (s *SaveAndSubmitAuthInfoRequest) SetOrganizationCode(v string) *SaveAndSubmitAuthInfoRequest {
	s.OrganizationCode = &v
	return s
}

func (s *SaveAndSubmitAuthInfoRequest) SetOrganizationCodeMediaId(v string) *SaveAndSubmitAuthInfoRequest {
	s.OrganizationCodeMediaId = &v
	return s
}

func (s *SaveAndSubmitAuthInfoRequest) SetRegistLocation(v string) *SaveAndSubmitAuthInfoRequest {
	s.RegistLocation = &v
	return s
}

func (s *SaveAndSubmitAuthInfoRequest) SetRegistNum(v string) *SaveAndSubmitAuthInfoRequest {
	s.RegistNum = &v
	return s
}

func (s *SaveAndSubmitAuthInfoRequest) SetTargetCorpId(v string) *SaveAndSubmitAuthInfoRequest {
	s.TargetCorpId = &v
	return s
}

func (s *SaveAndSubmitAuthInfoRequest) SetUnifiedSocialCredit(v string) *SaveAndSubmitAuthInfoRequest {
	s.UnifiedSocialCredit = &v
	return s
}

type SaveAndSubmitAuthInfoResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SaveAndSubmitAuthInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveAndSubmitAuthInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SaveAndSubmitAuthInfoResponseBody) SetSuccess(v bool) *SaveAndSubmitAuthInfoResponseBody {
	s.Success = &v
	return s
}

type SaveAndSubmitAuthInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveAndSubmitAuthInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveAndSubmitAuthInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveAndSubmitAuthInfoResponse) GoString() string {
	return s.String()
}

func (s *SaveAndSubmitAuthInfoResponse) SetHeaders(v map[string]*string) *SaveAndSubmitAuthInfoResponse {
	s.Headers = v
	return s
}

func (s *SaveAndSubmitAuthInfoResponse) SetStatusCode(v int32) *SaveAndSubmitAuthInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveAndSubmitAuthInfoResponse) SetBody(v *SaveAndSubmitAuthInfoResponseBody) *SaveAndSubmitAuthInfoResponse {
	s.Body = v
	return s
}

type SaveOpenTerminalInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SaveOpenTerminalInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveOpenTerminalInfoHeaders) GoString() string {
	return s.String()
}

func (s *SaveOpenTerminalInfoHeaders) SetCommonHeaders(v map[string]*string) *SaveOpenTerminalInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveOpenTerminalInfoHeaders) SetXAcsDingtalkAccessToken(v string) *SaveOpenTerminalInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SaveOpenTerminalInfoRequest struct {
	CorpId    *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	LogSource *string `json:"logSource,omitempty" xml:"logSource,omitempty"`
	LogType   *string `json:"logType,omitempty" xml:"logType,omitempty"`
	OpenExt   *string `json:"openExt,omitempty" xml:"openExt,omitempty"`
}

func (s SaveOpenTerminalInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveOpenTerminalInfoRequest) GoString() string {
	return s.String()
}

func (s *SaveOpenTerminalInfoRequest) SetCorpId(v string) *SaveOpenTerminalInfoRequest {
	s.CorpId = &v
	return s
}

func (s *SaveOpenTerminalInfoRequest) SetLogSource(v string) *SaveOpenTerminalInfoRequest {
	s.LogSource = &v
	return s
}

func (s *SaveOpenTerminalInfoRequest) SetLogType(v string) *SaveOpenTerminalInfoRequest {
	s.LogType = &v
	return s
}

func (s *SaveOpenTerminalInfoRequest) SetOpenExt(v string) *SaveOpenTerminalInfoRequest {
	s.OpenExt = &v
	return s
}

type SaveOpenTerminalInfoResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SaveOpenTerminalInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveOpenTerminalInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SaveOpenTerminalInfoResponseBody) SetSuccess(v bool) *SaveOpenTerminalInfoResponseBody {
	s.Success = &v
	return s
}

type SaveOpenTerminalInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveOpenTerminalInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveOpenTerminalInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveOpenTerminalInfoResponse) GoString() string {
	return s.String()
}

func (s *SaveOpenTerminalInfoResponse) SetHeaders(v map[string]*string) *SaveOpenTerminalInfoResponse {
	s.Headers = v
	return s
}

func (s *SaveOpenTerminalInfoResponse) SetStatusCode(v int32) *SaveOpenTerminalInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveOpenTerminalInfoResponse) SetBody(v *SaveOpenTerminalInfoResponseBody) *SaveOpenTerminalInfoResponse {
	s.Body = v
	return s
}

type SaveWhiteAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SaveWhiteAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveWhiteAppHeaders) GoString() string {
	return s.String()
}

func (s *SaveWhiteAppHeaders) SetCommonHeaders(v map[string]*string) *SaveWhiteAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveWhiteAppHeaders) SetXAcsDingtalkAccessToken(v string) *SaveWhiteAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SaveWhiteAppRequest struct {
	AgentIdList []*int64 `json:"agentIdList,omitempty" xml:"agentIdList,omitempty" type:"Repeated"`
	Operation   *string  `json:"operation,omitempty" xml:"operation,omitempty"`
}

func (s SaveWhiteAppRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveWhiteAppRequest) GoString() string {
	return s.String()
}

func (s *SaveWhiteAppRequest) SetAgentIdList(v []*int64) *SaveWhiteAppRequest {
	s.AgentIdList = v
	return s
}

func (s *SaveWhiteAppRequest) SetOperation(v string) *SaveWhiteAppRequest {
	s.Operation = &v
	return s
}

type SaveWhiteAppResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SaveWhiteAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveWhiteAppResponseBody) GoString() string {
	return s.String()
}

func (s *SaveWhiteAppResponseBody) SetSuccess(v bool) *SaveWhiteAppResponseBody {
	s.Success = &v
	return s
}

type SaveWhiteAppResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveWhiteAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveWhiteAppResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveWhiteAppResponse) GoString() string {
	return s.String()
}

func (s *SaveWhiteAppResponse) SetHeaders(v map[string]*string) *SaveWhiteAppResponse {
	s.Headers = v
	return s
}

func (s *SaveWhiteAppResponse) SetStatusCode(v int32) *SaveWhiteAppResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveWhiteAppResponse) SetBody(v *SaveWhiteAppResponseBody) *SaveWhiteAppResponse {
	s.Body = v
	return s
}

type SearchOrgInnerGroupInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchOrgInnerGroupInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchOrgInnerGroupInfoHeaders) GoString() string {
	return s.String()
}

func (s *SearchOrgInnerGroupInfoHeaders) SetCommonHeaders(v map[string]*string) *SearchOrgInnerGroupInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchOrgInnerGroupInfoHeaders) SetXAcsDingtalkAccessToken(v string) *SearchOrgInnerGroupInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchOrgInnerGroupInfoRequest struct {
	CreateTimeEnd          *int64  `json:"createTimeEnd,omitempty" xml:"createTimeEnd,omitempty"`
	CreateTimeStart        *int64  `json:"createTimeStart,omitempty" xml:"createTimeStart,omitempty"`
	GroupMembersCountEnd   *int32  `json:"groupMembersCountEnd,omitempty" xml:"groupMembersCountEnd,omitempty"`
	GroupMembersCountStart *int32  `json:"groupMembersCountStart,omitempty" xml:"groupMembersCountStart,omitempty"`
	GroupName              *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	GroupOwner             *string `json:"groupOwner,omitempty" xml:"groupOwner,omitempty"`
	LastActiveTimeEnd      *int64  `json:"lastActiveTimeEnd,omitempty" xml:"lastActiveTimeEnd,omitempty"`
	LastActiveTimeStart    *int64  `json:"lastActiveTimeStart,omitempty" xml:"lastActiveTimeStart,omitempty"`
	OperatorUserId         *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	PageSize               *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	PageStart              *int32  `json:"pageStart,omitempty" xml:"pageStart,omitempty"`
	SyncToDingpan          *int32  `json:"syncToDingpan,omitempty" xml:"syncToDingpan,omitempty"`
	Uuid                   *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s SearchOrgInnerGroupInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchOrgInnerGroupInfoRequest) GoString() string {
	return s.String()
}

func (s *SearchOrgInnerGroupInfoRequest) SetCreateTimeEnd(v int64) *SearchOrgInnerGroupInfoRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetCreateTimeStart(v int64) *SearchOrgInnerGroupInfoRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetGroupMembersCountEnd(v int32) *SearchOrgInnerGroupInfoRequest {
	s.GroupMembersCountEnd = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetGroupMembersCountStart(v int32) *SearchOrgInnerGroupInfoRequest {
	s.GroupMembersCountStart = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetGroupName(v string) *SearchOrgInnerGroupInfoRequest {
	s.GroupName = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetGroupOwner(v string) *SearchOrgInnerGroupInfoRequest {
	s.GroupOwner = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetLastActiveTimeEnd(v int64) *SearchOrgInnerGroupInfoRequest {
	s.LastActiveTimeEnd = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetLastActiveTimeStart(v int64) *SearchOrgInnerGroupInfoRequest {
	s.LastActiveTimeStart = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetOperatorUserId(v string) *SearchOrgInnerGroupInfoRequest {
	s.OperatorUserId = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetPageSize(v int32) *SearchOrgInnerGroupInfoRequest {
	s.PageSize = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetPageStart(v int32) *SearchOrgInnerGroupInfoRequest {
	s.PageStart = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetSyncToDingpan(v int32) *SearchOrgInnerGroupInfoRequest {
	s.SyncToDingpan = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetUuid(v string) *SearchOrgInnerGroupInfoRequest {
	s.Uuid = &v
	return s
}

type SearchOrgInnerGroupInfoResponseBody struct {
	ItemCount  *int32                                      `json:"itemCount,omitempty" xml:"itemCount,omitempty"`
	Items      []*SearchOrgInnerGroupInfoResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	TotalCount *int64                                      `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s SearchOrgInnerGroupInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchOrgInnerGroupInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SearchOrgInnerGroupInfoResponseBody) SetItemCount(v int32) *SearchOrgInnerGroupInfoResponseBody {
	s.ItemCount = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBody) SetItems(v []*SearchOrgInnerGroupInfoResponseBodyItems) *SearchOrgInnerGroupInfoResponseBody {
	s.Items = v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBody) SetTotalCount(v int64) *SearchOrgInnerGroupInfoResponseBody {
	s.TotalCount = &v
	return s
}

type SearchOrgInnerGroupInfoResponseBodyItems struct {
	GroupAdminsCount        *int32  `json:"groupAdminsCount,omitempty" xml:"groupAdminsCount,omitempty"`
	GroupCreateTime         *int64  `json:"groupCreateTime,omitempty" xml:"groupCreateTime,omitempty"`
	GroupLastActiveTime     *int64  `json:"groupLastActiveTime,omitempty" xml:"groupLastActiveTime,omitempty"`
	GroupLastActiveTimeShow *string `json:"groupLastActiveTimeShow,omitempty" xml:"groupLastActiveTimeShow,omitempty"`
	GroupMembersCount       *int32  `json:"groupMembersCount,omitempty" xml:"groupMembersCount,omitempty"`
	GroupName               *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	GroupOwner              *string `json:"groupOwner,omitempty" xml:"groupOwner,omitempty"`
	GroupOwnerUserId        *string `json:"groupOwnerUserId,omitempty" xml:"groupOwnerUserId,omitempty"`
	OpenConversationId      *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	Status                  *int32  `json:"status,omitempty" xml:"status,omitempty"`
	SyncToDingpan           *int32  `json:"syncToDingpan,omitempty" xml:"syncToDingpan,omitempty"`
	TemplateId              *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	TemplateName            *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	UsedQuota               *int64  `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
}

func (s SearchOrgInnerGroupInfoResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s SearchOrgInnerGroupInfoResponseBodyItems) GoString() string {
	return s.String()
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetGroupAdminsCount(v int32) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.GroupAdminsCount = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetGroupCreateTime(v int64) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.GroupCreateTime = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetGroupLastActiveTime(v int64) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.GroupLastActiveTime = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetGroupLastActiveTimeShow(v string) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.GroupLastActiveTimeShow = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetGroupMembersCount(v int32) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.GroupMembersCount = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetGroupName(v string) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.GroupName = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetGroupOwner(v string) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.GroupOwner = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetGroupOwnerUserId(v string) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.GroupOwnerUserId = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetOpenConversationId(v string) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.OpenConversationId = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetStatus(v int32) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.Status = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetSyncToDingpan(v int32) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.SyncToDingpan = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetTemplateId(v string) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.TemplateId = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetTemplateName(v string) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.TemplateName = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetUsedQuota(v int64) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.UsedQuota = &v
	return s
}

type SearchOrgInnerGroupInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchOrgInnerGroupInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchOrgInnerGroupInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchOrgInnerGroupInfoResponse) GoString() string {
	return s.String()
}

func (s *SearchOrgInnerGroupInfoResponse) SetHeaders(v map[string]*string) *SearchOrgInnerGroupInfoResponse {
	s.Headers = v
	return s
}

func (s *SearchOrgInnerGroupInfoResponse) SetStatusCode(v int32) *SearchOrgInnerGroupInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponse) SetBody(v *SearchOrgInnerGroupInfoResponseBody) *SearchOrgInnerGroupInfoResponse {
	s.Body = v
	return s
}

type SendAppDingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendAppDingHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendAppDingHeaders) GoString() string {
	return s.String()
}

func (s *SendAppDingHeaders) SetCommonHeaders(v map[string]*string) *SendAppDingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendAppDingHeaders) SetXAcsDingtalkAccessToken(v string) *SendAppDingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendAppDingRequest struct {
	Content *string   `json:"content,omitempty" xml:"content,omitempty"`
	Userids []*string `json:"userids,omitempty" xml:"userids,omitempty" type:"Repeated"`
}

func (s SendAppDingRequest) String() string {
	return tea.Prettify(s)
}

func (s SendAppDingRequest) GoString() string {
	return s.String()
}

func (s *SendAppDingRequest) SetContent(v string) *SendAppDingRequest {
	s.Content = &v
	return s
}

func (s *SendAppDingRequest) SetUserids(v []*string) *SendAppDingRequest {
	s.Userids = v
	return s
}

type SendAppDingResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s SendAppDingResponse) String() string {
	return tea.Prettify(s)
}

func (s SendAppDingResponse) GoString() string {
	return s.String()
}

func (s *SendAppDingResponse) SetHeaders(v map[string]*string) *SendAppDingResponse {
	s.Headers = v
	return s
}

func (s *SendAppDingResponse) SetStatusCode(v int32) *SendAppDingResponse {
	s.StatusCode = &v
	return s
}

type SendInvitationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendInvitationHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendInvitationHeaders) GoString() string {
	return s.String()
}

func (s *SendInvitationHeaders) SetCommonHeaders(v map[string]*string) *SendInvitationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendInvitationHeaders) SetXAcsDingtalkAccessToken(v string) *SendInvitationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendInvitationRequest struct {
	DeptId         *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
	OrgAlias       *string `json:"orgAlias,omitempty" xml:"orgAlias,omitempty"`
	PartnerLabelId *int64  `json:"partnerLabelId,omitempty" xml:"partnerLabelId,omitempty"`
	PartnerNum     *string `json:"partnerNum,omitempty" xml:"partnerNum,omitempty"`
	Phone          *string `json:"phone,omitempty" xml:"phone,omitempty"`
}

func (s SendInvitationRequest) String() string {
	return tea.Prettify(s)
}

func (s SendInvitationRequest) GoString() string {
	return s.String()
}

func (s *SendInvitationRequest) SetDeptId(v string) *SendInvitationRequest {
	s.DeptId = &v
	return s
}

func (s *SendInvitationRequest) SetOrgAlias(v string) *SendInvitationRequest {
	s.OrgAlias = &v
	return s
}

func (s *SendInvitationRequest) SetPartnerLabelId(v int64) *SendInvitationRequest {
	s.PartnerLabelId = &v
	return s
}

func (s *SendInvitationRequest) SetPartnerNum(v string) *SendInvitationRequest {
	s.PartnerNum = &v
	return s
}

func (s *SendInvitationRequest) SetPhone(v string) *SendInvitationRequest {
	s.Phone = &v
	return s
}

type SendInvitationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s SendInvitationResponse) String() string {
	return tea.Prettify(s)
}

func (s SendInvitationResponse) GoString() string {
	return s.String()
}

func (s *SendInvitationResponse) SetHeaders(v map[string]*string) *SendInvitationResponse {
	s.Headers = v
	return s
}

func (s *SendInvitationResponse) SetStatusCode(v int32) *SendInvitationResponse {
	s.StatusCode = &v
	return s
}

type SendPhoneDingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendPhoneDingHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendPhoneDingHeaders) GoString() string {
	return s.String()
}

func (s *SendPhoneDingHeaders) SetCommonHeaders(v map[string]*string) *SendPhoneDingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendPhoneDingHeaders) SetXAcsDingtalkAccessToken(v string) *SendPhoneDingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendPhoneDingRequest struct {
	Content *string   `json:"content,omitempty" xml:"content,omitempty"`
	Userids []*string `json:"userids,omitempty" xml:"userids,omitempty" type:"Repeated"`
}

func (s SendPhoneDingRequest) String() string {
	return tea.Prettify(s)
}

func (s SendPhoneDingRequest) GoString() string {
	return s.String()
}

func (s *SendPhoneDingRequest) SetContent(v string) *SendPhoneDingRequest {
	s.Content = &v
	return s
}

func (s *SendPhoneDingRequest) SetUserids(v []*string) *SendPhoneDingRequest {
	s.Userids = v
	return s
}

type SendPhoneDingResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SendPhoneDingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendPhoneDingResponseBody) GoString() string {
	return s.String()
}

func (s *SendPhoneDingResponseBody) SetSuccess(v bool) *SendPhoneDingResponseBody {
	s.Success = &v
	return s
}

type SendPhoneDingResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendPhoneDingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendPhoneDingResponse) String() string {
	return tea.Prettify(s)
}

func (s SendPhoneDingResponse) GoString() string {
	return s.String()
}

func (s *SendPhoneDingResponse) SetHeaders(v map[string]*string) *SendPhoneDingResponse {
	s.Headers = v
	return s
}

func (s *SendPhoneDingResponse) SetStatusCode(v int32) *SendPhoneDingResponse {
	s.StatusCode = &v
	return s
}

func (s *SendPhoneDingResponse) SetBody(v *SendPhoneDingResponseBody) *SendPhoneDingResponse {
	s.Body = v
	return s
}

type SetDeptPartnerTypeAndNumHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SetDeptPartnerTypeAndNumHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetDeptPartnerTypeAndNumHeaders) GoString() string {
	return s.String()
}

func (s *SetDeptPartnerTypeAndNumHeaders) SetCommonHeaders(v map[string]*string) *SetDeptPartnerTypeAndNumHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetDeptPartnerTypeAndNumHeaders) SetXAcsDingtalkAccessToken(v string) *SetDeptPartnerTypeAndNumHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SetDeptPartnerTypeAndNumRequest struct {
	DeptId     *string   `json:"deptId,omitempty" xml:"deptId,omitempty"`
	LabelIds   []*string `json:"labelIds,omitempty" xml:"labelIds,omitempty" type:"Repeated"`
	PartnerNum *string   `json:"partnerNum,omitempty" xml:"partnerNum,omitempty"`
}

func (s SetDeptPartnerTypeAndNumRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDeptPartnerTypeAndNumRequest) GoString() string {
	return s.String()
}

func (s *SetDeptPartnerTypeAndNumRequest) SetDeptId(v string) *SetDeptPartnerTypeAndNumRequest {
	s.DeptId = &v
	return s
}

func (s *SetDeptPartnerTypeAndNumRequest) SetLabelIds(v []*string) *SetDeptPartnerTypeAndNumRequest {
	s.LabelIds = v
	return s
}

func (s *SetDeptPartnerTypeAndNumRequest) SetPartnerNum(v string) *SetDeptPartnerTypeAndNumRequest {
	s.PartnerNum = &v
	return s
}

type SetDeptPartnerTypeAndNumResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s SetDeptPartnerTypeAndNumResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDeptPartnerTypeAndNumResponse) GoString() string {
	return s.String()
}

func (s *SetDeptPartnerTypeAndNumResponse) SetHeaders(v map[string]*string) *SetDeptPartnerTypeAndNumResponse {
	s.Headers = v
	return s
}

func (s *SetDeptPartnerTypeAndNumResponse) SetStatusCode(v int32) *SetDeptPartnerTypeAndNumResponse {
	s.StatusCode = &v
	return s
}

type UpdateCategoryNameHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateCategoryNameHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateCategoryNameHeaders) GoString() string {
	return s.String()
}

func (s *UpdateCategoryNameHeaders) SetCommonHeaders(v map[string]*string) *UpdateCategoryNameHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateCategoryNameHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateCategoryNameHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateCategoryNameRequest struct {
	CurrentCategoryName *string `json:"currentCategoryName,omitempty" xml:"currentCategoryName,omitempty"`
	TargetCategoryName  *string `json:"targetCategoryName,omitempty" xml:"targetCategoryName,omitempty"`
}

func (s UpdateCategoryNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCategoryNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateCategoryNameRequest) SetCurrentCategoryName(v string) *UpdateCategoryNameRequest {
	s.CurrentCategoryName = &v
	return s
}

func (s *UpdateCategoryNameRequest) SetTargetCategoryName(v string) *UpdateCategoryNameRequest {
	s.TargetCategoryName = &v
	return s
}

type UpdateCategoryNameResponseBody struct {
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateCategoryNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCategoryNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCategoryNameResponseBody) SetStatus(v int64) *UpdateCategoryNameResponseBody {
	s.Status = &v
	return s
}

type UpdateCategoryNameResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateCategoryNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCategoryNameResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCategoryNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateCategoryNameResponse) SetHeaders(v map[string]*string) *UpdateCategoryNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateCategoryNameResponse) SetStatusCode(v int32) *UpdateCategoryNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCategoryNameResponse) SetBody(v *UpdateCategoryNameResponseBody) *UpdateCategoryNameResponse {
	s.Body = v
	return s
}

type UpdateFileStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateFileStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileStatusHeaders) GoString() string {
	return s.String()
}

func (s *UpdateFileStatusHeaders) SetCommonHeaders(v map[string]*string) *UpdateFileStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateFileStatusHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateFileStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateFileStatusRequest struct {
	RequestIds []*string `json:"requestIds,omitempty" xml:"requestIds,omitempty" type:"Repeated"`
	Status     *int32    `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateFileStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileStatusRequest) SetRequestIds(v []*string) *UpdateFileStatusRequest {
	s.RequestIds = v
	return s
}

func (s *UpdateFileStatusRequest) SetStatus(v int32) *UpdateFileStatusRequest {
	s.Status = &v
	return s
}

type UpdateFileStatusResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateFileStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFileStatusResponseBody) SetSuccess(v bool) *UpdateFileStatusResponseBody {
	s.Success = &v
	return s
}

type UpdateFileStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateFileStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFileStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateFileStatusResponse) SetHeaders(v map[string]*string) *UpdateFileStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateFileStatusResponse) SetStatusCode(v int32) *UpdateFileStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFileStatusResponse) SetBody(v *UpdateFileStatusResponseBody) *UpdateFileStatusResponse {
	s.Body = v
	return s
}

type UpdateMiniAppVersionStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateMiniAppVersionStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateMiniAppVersionStatusHeaders) GoString() string {
	return s.String()
}

func (s *UpdateMiniAppVersionStatusHeaders) SetCommonHeaders(v map[string]*string) *UpdateMiniAppVersionStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateMiniAppVersionStatusHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateMiniAppVersionStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateMiniAppVersionStatusRequest struct {
	MiniAppId   *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	Version     *string `json:"version,omitempty" xml:"version,omitempty"`
	VersionType *int32  `json:"versionType,omitempty" xml:"versionType,omitempty"`
}

func (s UpdateMiniAppVersionStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMiniAppVersionStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateMiniAppVersionStatusRequest) SetMiniAppId(v string) *UpdateMiniAppVersionStatusRequest {
	s.MiniAppId = &v
	return s
}

func (s *UpdateMiniAppVersionStatusRequest) SetVersion(v string) *UpdateMiniAppVersionStatusRequest {
	s.Version = &v
	return s
}

func (s *UpdateMiniAppVersionStatusRequest) SetVersionType(v int32) *UpdateMiniAppVersionStatusRequest {
	s.VersionType = &v
	return s
}

type UpdateMiniAppVersionStatusResponseBody struct {
	Cause *string `json:"cause,omitempty" xml:"cause,omitempty"`
	Code  *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s UpdateMiniAppVersionStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMiniAppVersionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMiniAppVersionStatusResponseBody) SetCause(v string) *UpdateMiniAppVersionStatusResponseBody {
	s.Cause = &v
	return s
}

func (s *UpdateMiniAppVersionStatusResponseBody) SetCode(v string) *UpdateMiniAppVersionStatusResponseBody {
	s.Code = &v
	return s
}

type UpdateMiniAppVersionStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateMiniAppVersionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateMiniAppVersionStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMiniAppVersionStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateMiniAppVersionStatusResponse) SetHeaders(v map[string]*string) *UpdateMiniAppVersionStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateMiniAppVersionStatusResponse) SetStatusCode(v int32) *UpdateMiniAppVersionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMiniAppVersionStatusResponse) SetBody(v *UpdateMiniAppVersionStatusResponseBody) *UpdateMiniAppVersionStatusResponse {
	s.Body = v
	return s
}

type UpdatePartnerVisibilityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdatePartnerVisibilityHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdatePartnerVisibilityHeaders) GoString() string {
	return s.String()
}

func (s *UpdatePartnerVisibilityHeaders) SetCommonHeaders(v map[string]*string) *UpdatePartnerVisibilityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdatePartnerVisibilityHeaders) SetXAcsDingtalkAccessToken(v string) *UpdatePartnerVisibilityHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdatePartnerVisibilityRequest struct {
	DeptIds []*int64  `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	LabelId *int64    `json:"labelId,omitempty" xml:"labelId,omitempty"`
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s UpdatePartnerVisibilityRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePartnerVisibilityRequest) GoString() string {
	return s.String()
}

func (s *UpdatePartnerVisibilityRequest) SetDeptIds(v []*int64) *UpdatePartnerVisibilityRequest {
	s.DeptIds = v
	return s
}

func (s *UpdatePartnerVisibilityRequest) SetLabelId(v int64) *UpdatePartnerVisibilityRequest {
	s.LabelId = &v
	return s
}

func (s *UpdatePartnerVisibilityRequest) SetUserIds(v []*string) *UpdatePartnerVisibilityRequest {
	s.UserIds = v
	return s
}

type UpdatePartnerVisibilityResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *bool              `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePartnerVisibilityResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePartnerVisibilityResponse) GoString() string {
	return s.String()
}

func (s *UpdatePartnerVisibilityResponse) SetHeaders(v map[string]*string) *UpdatePartnerVisibilityResponse {
	s.Headers = v
	return s
}

func (s *UpdatePartnerVisibilityResponse) SetStatusCode(v int32) *UpdatePartnerVisibilityResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePartnerVisibilityResponse) SetBody(v bool) *UpdatePartnerVisibilityResponse {
	s.Body = &v
	return s
}

type UpdateRoleVisibilityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateRoleVisibilityHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleVisibilityHeaders) GoString() string {
	return s.String()
}

func (s *UpdateRoleVisibilityHeaders) SetCommonHeaders(v map[string]*string) *UpdateRoleVisibilityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateRoleVisibilityHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateRoleVisibilityHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateRoleVisibilityRequest struct {
	DeptIds []*int64  `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	LabelId *int64    `json:"labelId,omitempty" xml:"labelId,omitempty"`
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s UpdateRoleVisibilityRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleVisibilityRequest) GoString() string {
	return s.String()
}

func (s *UpdateRoleVisibilityRequest) SetDeptIds(v []*int64) *UpdateRoleVisibilityRequest {
	s.DeptIds = v
	return s
}

func (s *UpdateRoleVisibilityRequest) SetLabelId(v int64) *UpdateRoleVisibilityRequest {
	s.LabelId = &v
	return s
}

func (s *UpdateRoleVisibilityRequest) SetUserIds(v []*string) *UpdateRoleVisibilityRequest {
	s.UserIds = v
	return s
}

type UpdateRoleVisibilityResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *bool              `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRoleVisibilityResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleVisibilityResponse) GoString() string {
	return s.String()
}

func (s *UpdateRoleVisibilityResponse) SetHeaders(v map[string]*string) *UpdateRoleVisibilityResponse {
	s.Headers = v
	return s
}

func (s *UpdateRoleVisibilityResponse) SetStatusCode(v int32) *UpdateRoleVisibilityResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRoleVisibilityResponse) SetBody(v bool) *UpdateRoleVisibilityResponse {
	s.Body = &v
	return s
}

type UpdateStorageModeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateStorageModeHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateStorageModeHeaders) GoString() string {
	return s.String()
}

func (s *UpdateStorageModeHeaders) SetCommonHeaders(v map[string]*string) *UpdateStorageModeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateStorageModeHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateStorageModeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateStorageModeRequest struct {
	FileStorageMode *string `json:"fileStorageMode,omitempty" xml:"fileStorageMode,omitempty"`
	TargetCorpId    *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
}

func (s UpdateStorageModeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStorageModeRequest) GoString() string {
	return s.String()
}

func (s *UpdateStorageModeRequest) SetFileStorageMode(v string) *UpdateStorageModeRequest {
	s.FileStorageMode = &v
	return s
}

func (s *UpdateStorageModeRequest) SetTargetCorpId(v string) *UpdateStorageModeRequest {
	s.TargetCorpId = &v
	return s
}

type UpdateStorageModeResponseBody struct {
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
}

func (s UpdateStorageModeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateStorageModeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStorageModeResponseBody) SetTargetCorpId(v string) *UpdateStorageModeResponseBody {
	s.TargetCorpId = &v
	return s
}

type UpdateStorageModeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateStorageModeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateStorageModeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateStorageModeResponse) GoString() string {
	return s.String()
}

func (s *UpdateStorageModeResponse) SetHeaders(v map[string]*string) *UpdateStorageModeResponse {
	s.Headers = v
	return s
}

func (s *UpdateStorageModeResponse) SetStatusCode(v int32) *UpdateStorageModeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStorageModeResponse) SetBody(v *UpdateStorageModeResponseBody) *UpdateStorageModeResponse {
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

func (client *Client) AddOrgWithOptions(request *AddOrgRequest, headers *AddOrgHeaders, runtime *util.RuntimeOptions) (_result *AddOrgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MobileNum)) {
		body["mobileNum"] = request.MobileNum
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
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
		Action:      tea.String("AddOrg"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/orgnizations"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddOrgResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddOrg(request *AddOrgRequest) (_result *AddOrgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddOrgHeaders{}
	_result = &AddOrgResponse{}
	_body, _err := client.AddOrgWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApproveProcessCallbackWithOptions(request *ApproveProcessCallbackRequest, headers *ApproveProcessCallbackHeaders, runtime *util.RuntimeOptions) (_result *ApproveProcessCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKeyId)) {
		body["accessKeyId"] = request.AccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessKeySecret)) {
		body["accessKeySecret"] = request.AccessKeySecret
	}

	if !tea.BoolValue(util.IsUnset(request.Request)) {
		body["request"] = request.Request
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCorpId)) {
		body["targetCorpId"] = request.TargetCorpId
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
		Action:      tea.String("ApproveProcessCallback"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/approvalResults/callback"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ApproveProcessCallbackResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApproveProcessCallback(request *ApproveProcessCallbackRequest) (_result *ApproveProcessCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ApproveProcessCallbackHeaders{}
	_result = &ApproveProcessCallbackResponse{}
	_body, _err := client.ApproveProcessCallbackWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BanOrOpenGroupWordsWithOptions(request *BanOrOpenGroupWordsRequest, headers *BanOrOpenGroupWordsHeaders, runtime *util.RuntimeOptions) (_result *BanOrOpenGroupWordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BanWordsType)) {
		body["banWordsType"] = request.BanWordsType
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConverationId)) {
		body["openConverationId"] = request.OpenConverationId
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
		Action:      tea.String("BanOrOpenGroupWords"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/enterpriseSecurities/banOrOpenGroupWords"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BanOrOpenGroupWordsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BanOrOpenGroupWords(request *BanOrOpenGroupWordsRequest) (_result *BanOrOpenGroupWordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BanOrOpenGroupWordsHeaders{}
	_result = &BanOrOpenGroupWordsResponse{}
	_body, _err := client.BanOrOpenGroupWordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCategoryAndBindingGroupsWithOptions(request *CreateCategoryAndBindingGroupsRequest, headers *CreateCategoryAndBindingGroupsHeaders, runtime *util.RuntimeOptions) (_result *CreateCategoryAndBindingGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryName)) {
		body["categoryName"] = request.CategoryName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupIds)) {
		body["groupIds"] = request.GroupIds
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
		Action:      tea.String("CreateCategoryAndBindingGroups"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/messageCategories/categories/createAndBind"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCategoryAndBindingGroupsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCategoryAndBindingGroups(request *CreateCategoryAndBindingGroupsRequest) (_result *CreateCategoryAndBindingGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateCategoryAndBindingGroupsHeaders{}
	_result = &CreateCategoryAndBindingGroupsResponse{}
	_body, _err := client.CreateCategoryAndBindingGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMessageCategoryWithOptions(request *CreateMessageCategoryRequest, headers *CreateMessageCategoryHeaders, runtime *util.RuntimeOptions) (_result *CreateMessageCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryName)) {
		body["categoryName"] = request.CategoryName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupIds)) {
		body["groupIds"] = request.GroupIds
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
		Action:      tea.String("CreateMessageCategory"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/messageCategories/categories/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMessageCategoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMessageCategory(request *CreateMessageCategoryRequest) (_result *CreateMessageCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateMessageCategoryHeaders{}
	_result = &CreateMessageCategoryResponse{}
	_body, _err := client.CreateMessageCategoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRuleWithOptions(request *CreateRuleRequest, headers *CreateRuleHeaders, runtime *util.RuntimeOptions) (_result *CreateRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomPlan)) {
		body["customPlan"] = request.CustomPlan
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
		Action:      tea.String("CreateRule"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/messageCategories/rules"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRuleResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRule(request *CreateRuleRequest) (_result *CreateRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateRuleHeaders{}
	_result = &CreateRuleResponse{}
	_body, _err := client.CreateRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTrustedDeviceWithOptions(request *CreateTrustedDeviceRequest, headers *CreateTrustedDeviceHeaders, runtime *util.RuntimeOptions) (_result *CreateTrustedDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Did)) {
		body["did"] = request.Did
	}

	if !tea.BoolValue(util.IsUnset(request.MacAddress)) {
		body["macAddress"] = request.MacAddress
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		body["platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
		Action:      tea.String("CreateTrustedDevice"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/trustedDevices"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTrustedDeviceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTrustedDevice(request *CreateTrustedDeviceRequest) (_result *CreateTrustedDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateTrustedDeviceHeaders{}
	_result = &CreateTrustedDeviceResponse{}
	_body, _err := client.CreateTrustedDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTrustedDeviceBatchWithOptions(request *CreateTrustedDeviceBatchRequest, headers *CreateTrustedDeviceBatchHeaders, runtime *util.RuntimeOptions) (_result *CreateTrustedDeviceBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MacAddressList)) {
		body["macAddressList"] = request.MacAddressList
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		body["platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
		Action:      tea.String("CreateTrustedDeviceBatch"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/trusts/devices"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTrustedDeviceBatchResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTrustedDeviceBatch(request *CreateTrustedDeviceBatchRequest) (_result *CreateTrustedDeviceBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateTrustedDeviceBatchHeaders{}
	_result = &CreateTrustedDeviceBatchResponse{}
	_body, _err := client.CreateTrustedDeviceBatchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAcrossCloudStroageConfigsWithOptions(targetCorpId *string, headers *DeleteAcrossCloudStroageConfigsHeaders, runtime *util.RuntimeOptions) (_result *DeleteAcrossCloudStroageConfigsResponse, _err error) {
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
		Action:      tea.String("DeleteAcrossCloudStroageConfigs"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/fileStorages/acrossClouds/configurations/" + tea.StringValue(targetCorpId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAcrossCloudStroageConfigsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAcrossCloudStroageConfigs(targetCorpId *string) (_result *DeleteAcrossCloudStroageConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteAcrossCloudStroageConfigsHeaders{}
	_result = &DeleteAcrossCloudStroageConfigsResponse{}
	_body, _err := client.DeleteAcrossCloudStroageConfigsWithOptions(targetCorpId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCommentWithOptions(publisherId *string, commentId *string, headers *DeleteCommentHeaders, runtime *util.RuntimeOptions) (_result *DeleteCommentResponse, _err error) {
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
		Action:      tea.String("DeleteComment"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/publishers/" + tea.StringValue(publisherId) + "/comments/" + tea.StringValue(commentId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("boolean"),
	}
	_result = &DeleteCommentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteComment(publisherId *string, commentId *string) (_result *DeleteCommentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteCommentHeaders{}
	_result = &DeleteCommentResponse{}
	_body, _err := client.DeleteCommentWithOptions(publisherId, commentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTrustedDeviceWithOptions(request *DeleteTrustedDeviceRequest, headers *DeleteTrustedDeviceHeaders, runtime *util.RuntimeOptions) (_result *DeleteTrustedDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KickOff)) {
		body["kickOff"] = request.KickOff
	}

	if !tea.BoolValue(util.IsUnset(request.MacAddress)) {
		body["macAddress"] = request.MacAddress
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
		Action:      tea.String("DeleteTrustedDevice"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/trustedDevices/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTrustedDeviceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTrustedDevice(request *DeleteTrustedDeviceRequest) (_result *DeleteTrustedDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteTrustedDeviceHeaders{}
	_result = &DeleteTrustedDeviceResponse{}
	_body, _err := client.DeleteTrustedDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DistributePartnerAppWithOptions(request *DistributePartnerAppRequest, headers *DistributePartnerAppHeaders, runtime *util.RuntimeOptions) (_result *DistributePartnerAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.SubCorpId)) {
		body["subCorpId"] = request.SubCorpId
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
		Action:      tea.String("DistributePartnerApp"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/partners/applications/distribute"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DistributePartnerAppResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DistributePartnerApp(request *DistributePartnerAppRequest) (_result *DistributePartnerAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DistributePartnerAppHeaders{}
	_result = &DistributePartnerAppResponse{}
	_body, _err := client.DistributePartnerAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExclusiveCreateDingPortalWithOptions(request *ExclusiveCreateDingPortalRequest, headers *ExclusiveCreateDingPortalHeaders, runtime *util.RuntimeOptions) (_result *ExclusiveCreateDingPortalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingPortalName)) {
		body["dingPortalName"] = request.DingPortalName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCorpId)) {
		body["targetCorpId"] = request.TargetCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateAppUuid)) {
		body["templateAppUuid"] = request.TemplateAppUuid
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCorpId)) {
		body["templateCorpId"] = request.TemplateCorpId
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
		Action:      tea.String("ExclusiveCreateDingPortal"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/workbenches/templates/spread"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ExclusiveCreateDingPortalResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExclusiveCreateDingPortal(request *ExclusiveCreateDingPortalRequest) (_result *ExclusiveCreateDingPortalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ExclusiveCreateDingPortalHeaders{}
	_result = &ExclusiveCreateDingPortalResponse{}
	_body, _err := client.ExclusiveCreateDingPortalWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FileStorageActiveStorageWithOptions(request *FileStorageActiveStorageRequest, headers *FileStorageActiveStorageHeaders, runtime *util.RuntimeOptions) (_result *FileStorageActiveStorageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKeyId)) {
		body["accessKeyId"] = request.AccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessKeySecret)) {
		body["accessKeySecret"] = request.AccessKeySecret
	}

	if !tea.BoolValue(util.IsUnset(request.Oss)) {
		body["oss"] = request.Oss
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCorpId)) {
		body["targetCorpId"] = request.TargetCorpId
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
		Action:      tea.String("FileStorageActiveStorage"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/fileStorages/active"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &FileStorageActiveStorageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FileStorageActiveStorage(request *FileStorageActiveStorageRequest) (_result *FileStorageActiveStorageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &FileStorageActiveStorageHeaders{}
	_result = &FileStorageActiveStorageResponse{}
	_body, _err := client.FileStorageActiveStorageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FileStorageCheckConnectionWithOptions(request *FileStorageCheckConnectionRequest, headers *FileStorageCheckConnectionHeaders, runtime *util.RuntimeOptions) (_result *FileStorageCheckConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKeyId)) {
		body["accessKeyId"] = request.AccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessKeySecret)) {
		body["accessKeySecret"] = request.AccessKeySecret
	}

	if !tea.BoolValue(util.IsUnset(request.Oss)) {
		body["oss"] = request.Oss
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCorpId)) {
		body["targetCorpId"] = request.TargetCorpId
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
		Action:      tea.String("FileStorageCheckConnection"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/fileStorages/connections/check"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &FileStorageCheckConnectionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FileStorageCheckConnection(request *FileStorageCheckConnectionRequest) (_result *FileStorageCheckConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &FileStorageCheckConnectionHeaders{}
	_result = &FileStorageCheckConnectionResponse{}
	_body, _err := client.FileStorageCheckConnectionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FileStorageGetQuotaDataWithOptions(request *FileStorageGetQuotaDataRequest, headers *FileStorageGetQuotaDataHeaders, runtime *util.RuntimeOptions) (_result *FileStorageGetQuotaDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCorpId)) {
		query["targetCorpId"] = request.TargetCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
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
		Action:      tea.String("FileStorageGetQuotaData"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/fileStorages/quotaDatas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &FileStorageGetQuotaDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FileStorageGetQuotaData(request *FileStorageGetQuotaDataRequest) (_result *FileStorageGetQuotaDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &FileStorageGetQuotaDataHeaders{}
	_result = &FileStorageGetQuotaDataResponse{}
	_body, _err := client.FileStorageGetQuotaDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FileStorageGetStorageStateWithOptions(request *FileStorageGetStorageStateRequest, headers *FileStorageGetStorageStateHeaders, runtime *util.RuntimeOptions) (_result *FileStorageGetStorageStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TargetCorpId)) {
		query["targetCorpId"] = request.TargetCorpId
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
		Action:      tea.String("FileStorageGetStorageState"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/fileStorages/states"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &FileStorageGetStorageStateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FileStorageGetStorageState(request *FileStorageGetStorageStateRequest) (_result *FileStorageGetStorageStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &FileStorageGetStorageStateHeaders{}
	_result = &FileStorageGetStorageStateResponse{}
	_body, _err := client.FileStorageGetStorageStateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FileStorageUpdateStorageWithOptions(request *FileStorageUpdateStorageRequest, headers *FileStorageUpdateStorageHeaders, runtime *util.RuntimeOptions) (_result *FileStorageUpdateStorageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKeyId)) {
		body["accessKeyId"] = request.AccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessKeySecret)) {
		body["accessKeySecret"] = request.AccessKeySecret
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCorpId)) {
		body["targetCorpId"] = request.TargetCorpId
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
		Action:      tea.String("FileStorageUpdateStorage"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/fileStorages/configurations"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &FileStorageUpdateStorageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FileStorageUpdateStorage(request *FileStorageUpdateStorageRequest) (_result *FileStorageUpdateStorageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &FileStorageUpdateStorageHeaders{}
	_result = &FileStorageUpdateStorageResponse{}
	_body, _err := client.FileStorageUpdateStorageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateDarkWaterMarkWithOptions(request *GenerateDarkWaterMarkRequest, headers *GenerateDarkWaterMarkHeaders, runtime *util.RuntimeOptions) (_result *GenerateDarkWaterMarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
		Action:      tea.String("GenerateDarkWaterMark"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/enterpriseSecurities/darkWatermarks/generate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateDarkWaterMarkResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateDarkWaterMark(request *GenerateDarkWaterMarkRequest) (_result *GenerateDarkWaterMarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GenerateDarkWaterMarkHeaders{}
	_result = &GenerateDarkWaterMarkResponse{}
	_body, _err := client.GenerateDarkWaterMarkWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAccountTransferListWithOptions(request *GetAccountTransferListRequest, headers *GetAccountTransferListHeaders, runtime *util.RuntimeOptions) (_result *GetAccountTransferListResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
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
		Action:      tea.String("GetAccountTransferList"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/dataTransfer/accounts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccountTransferListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAccountTransferList(request *GetAccountTransferListRequest) (_result *GetAccountTransferListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAccountTransferListHeaders{}
	_result = &GetAccountTransferListResponse{}
	_body, _err := client.GetAccountTransferListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetActiveUserSummaryWithOptions(dataId *string, headers *GetActiveUserSummaryHeaders, runtime *util.RuntimeOptions) (_result *GetActiveUserSummaryResponse, _err error) {
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
		Action:      tea.String("GetActiveUserSummary"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/data/dau/org/" + tea.StringValue(dataId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetActiveUserSummaryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetActiveUserSummary(dataId *string) (_result *GetActiveUserSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetActiveUserSummaryHeaders{}
	_result = &GetActiveUserSummaryResponse{}
	_body, _err := client.GetActiveUserSummaryWithOptions(dataId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAgentIdByRelatedAppIdWithOptions(request *GetAgentIdByRelatedAppIdRequest, headers *GetAgentIdByRelatedAppIdHeaders, runtime *util.RuntimeOptions) (_result *GetAgentIdByRelatedAppIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCorpId)) {
		query["targetCorpId"] = request.TargetCorpId
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
		Action:      tea.String("GetAgentIdByRelatedAppId"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/exclusiveDesigns/agentId"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAgentIdByRelatedAppIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAgentIdByRelatedAppId(request *GetAgentIdByRelatedAppIdRequest) (_result *GetAgentIdByRelatedAppIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAgentIdByRelatedAppIdHeaders{}
	_result = &GetAgentIdByRelatedAppIdResponse{}
	_body, _err := client.GetAgentIdByRelatedAppIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAllLabelableDeptsWithOptions(headers *GetAllLabelableDeptsHeaders, runtime *util.RuntimeOptions) (_result *GetAllLabelableDeptsResponse, _err error) {
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
		Action:      tea.String("GetAllLabelableDepts"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/partnerDepartments"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAllLabelableDeptsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAllLabelableDepts() (_result *GetAllLabelableDeptsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAllLabelableDeptsHeaders{}
	_result = &GetAllLabelableDeptsResponse{}
	_body, _err := client.GetAllLabelableDeptsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppDispatchInfoWithOptions(request *GetAppDispatchInfoRequest, headers *GetAppDispatchInfoHeaders, runtime *util.RuntimeOptions) (_result *GetAppDispatchInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
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
		Action:      tea.String("GetAppDispatchInfo"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/apps/distributionInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppDispatchInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAppDispatchInfo(request *GetAppDispatchInfoRequest) (_result *GetAppDispatchInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAppDispatchInfoHeaders{}
	_result = &GetAppDispatchInfoResponse{}
	_body, _err := client.GetAppDispatchInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCalenderSummaryWithOptions(dataId *string, headers *GetCalenderSummaryHeaders, runtime *util.RuntimeOptions) (_result *GetCalenderSummaryResponse, _err error) {
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
		Action:      tea.String("GetCalenderSummary"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/data/calendar/org/" + tea.StringValue(dataId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCalenderSummaryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCalenderSummary(dataId *string) (_result *GetCalenderSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCalenderSummaryHeaders{}
	_result = &GetCalenderSummaryResponse{}
	_body, _err := client.GetCalenderSummaryWithOptions(dataId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCommentListWithOptions(publisherId *string, request *GetCommentListRequest, headers *GetCommentListHeaders, runtime *util.RuntimeOptions) (_result *GetCommentListResponse, _err error) {
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
		Action:      tea.String("GetCommentList"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/publishers/" + tea.StringValue(publisherId) + "/comments/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCommentListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCommentList(publisherId *string, request *GetCommentListRequest) (_result *GetCommentListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCommentListHeaders{}
	_result = &GetCommentListResponse{}
	_body, _err := client.GetCommentListWithOptions(publisherId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConfBaseInfoByLogicalIdWithOptions(request *GetConfBaseInfoByLogicalIdRequest, headers *GetConfBaseInfoByLogicalIdHeaders, runtime *util.RuntimeOptions) (_result *GetConfBaseInfoByLogicalIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogicalConferenceId)) {
		query["logicalConferenceId"] = request.LogicalConferenceId
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
		Action:      tea.String("GetConfBaseInfoByLogicalId"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/data/conferences"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConfBaseInfoByLogicalIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConfBaseInfoByLogicalId(request *GetConfBaseInfoByLogicalIdRequest) (_result *GetConfBaseInfoByLogicalIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetConfBaseInfoByLogicalIdHeaders{}
	_result = &GetConfBaseInfoByLogicalIdResponse{}
	_body, _err := client.GetConfBaseInfoByLogicalIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConferenceDetailWithOptions(conferenceId *string, headers *GetConferenceDetailHeaders, runtime *util.RuntimeOptions) (_result *GetConferenceDetailResponse, _err error) {
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
		Action:      tea.String("GetConferenceDetail"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/data/conferences/" + tea.StringValue(conferenceId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConferenceDetailResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConferenceDetail(conferenceId *string) (_result *GetConferenceDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetConferenceDetailHeaders{}
	_result = &GetConferenceDetailResponse{}
	_body, _err := client.GetConferenceDetailWithOptions(conferenceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDingReportDeptSummaryWithOptions(dataId *string, request *GetDingReportDeptSummaryRequest, headers *GetDingReportDeptSummaryHeaders, runtime *util.RuntimeOptions) (_result *GetDingReportDeptSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
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
		Action:      tea.String("GetDingReportDeptSummary"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/data/report/dept/" + tea.StringValue(dataId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDingReportDeptSummaryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDingReportDeptSummary(dataId *string, request *GetDingReportDeptSummaryRequest) (_result *GetDingReportDeptSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDingReportDeptSummaryHeaders{}
	_result = &GetDingReportDeptSummaryResponse{}
	_body, _err := client.GetDingReportDeptSummaryWithOptions(dataId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDingReportSummaryWithOptions(dataId *string, headers *GetDingReportSummaryHeaders, runtime *util.RuntimeOptions) (_result *GetDingReportSummaryResponse, _err error) {
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
		Action:      tea.String("GetDingReportSummary"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/datas/" + tea.StringValue(dataId) + "/reports/organizations"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDingReportSummaryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDingReportSummary(dataId *string) (_result *GetDingReportSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDingReportSummaryHeaders{}
	_result = &GetDingReportSummaryResponse{}
	_body, _err := client.GetDingReportSummaryWithOptions(dataId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDocCreatedDeptSummaryWithOptions(dataId *string, request *GetDocCreatedDeptSummaryRequest, headers *GetDocCreatedDeptSummaryHeaders, runtime *util.RuntimeOptions) (_result *GetDocCreatedDeptSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
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
		Action:      tea.String("GetDocCreatedDeptSummary"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/data/doc/dept/" + tea.StringValue(dataId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDocCreatedDeptSummaryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDocCreatedDeptSummary(dataId *string, request *GetDocCreatedDeptSummaryRequest) (_result *GetDocCreatedDeptSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDocCreatedDeptSummaryHeaders{}
	_result = &GetDocCreatedDeptSummaryResponse{}
	_body, _err := client.GetDocCreatedDeptSummaryWithOptions(dataId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDocCreatedSummaryWithOptions(dataId *string, headers *GetDocCreatedSummaryHeaders, runtime *util.RuntimeOptions) (_result *GetDocCreatedSummaryResponse, _err error) {
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
		Action:      tea.String("GetDocCreatedSummary"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/data/doc/org/" + tea.StringValue(dataId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDocCreatedSummaryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDocCreatedSummary(dataId *string) (_result *GetDocCreatedSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDocCreatedSummaryHeaders{}
	_result = &GetDocCreatedSummaryResponse{}
	_body, _err := client.GetDocCreatedSummaryWithOptions(dataId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetExclusiveAccountAllOrgListWithOptions(request *GetExclusiveAccountAllOrgListRequest, headers *GetExclusiveAccountAllOrgListHeaders, runtime *util.RuntimeOptions) (_result *GetExclusiveAccountAllOrgListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
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
		Action:      tea.String("GetExclusiveAccountAllOrgList"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/exclusiveAccounts/organizations"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExclusiveAccountAllOrgListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetExclusiveAccountAllOrgList(request *GetExclusiveAccountAllOrgListRequest) (_result *GetExclusiveAccountAllOrgListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetExclusiveAccountAllOrgListHeaders{}
	_result = &GetExclusiveAccountAllOrgListResponse{}
	_body, _err := client.GetExclusiveAccountAllOrgListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGeneralFormCreatedDeptSummaryWithOptions(dataId *string, request *GetGeneralFormCreatedDeptSummaryRequest, headers *GetGeneralFormCreatedDeptSummaryHeaders, runtime *util.RuntimeOptions) (_result *GetGeneralFormCreatedDeptSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
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
		Action:      tea.String("GetGeneralFormCreatedDeptSummary"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/data/generalForm/dept/" + tea.StringValue(dataId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGeneralFormCreatedDeptSummaryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGeneralFormCreatedDeptSummary(dataId *string, request *GetGeneralFormCreatedDeptSummaryRequest) (_result *GetGeneralFormCreatedDeptSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetGeneralFormCreatedDeptSummaryHeaders{}
	_result = &GetGeneralFormCreatedDeptSummaryResponse{}
	_body, _err := client.GetGeneralFormCreatedDeptSummaryWithOptions(dataId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGeneralFormCreatedSummaryWithOptions(dataId *string, headers *GetGeneralFormCreatedSummaryHeaders, runtime *util.RuntimeOptions) (_result *GetGeneralFormCreatedSummaryResponse, _err error) {
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
		Action:      tea.String("GetGeneralFormCreatedSummary"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/data/generalForm/org/" + tea.StringValue(dataId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGeneralFormCreatedSummaryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGeneralFormCreatedSummary(dataId *string) (_result *GetGeneralFormCreatedSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetGeneralFormCreatedSummaryHeaders{}
	_result = &GetGeneralFormCreatedSummaryResponse{}
	_body, _err := client.GetGeneralFormCreatedSummaryWithOptions(dataId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGroupActiveInfoWithOptions(request *GetGroupActiveInfoRequest, headers *GetGroupActiveInfoHeaders, runtime *util.RuntimeOptions) (_result *GetGroupActiveInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingGroupId)) {
		query["dingGroupId"] = request.DingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		query["groupType"] = request.GroupType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
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
		Action:      tea.String("GetGroupActiveInfo"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/data/activeGroups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGroupActiveInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGroupActiveInfo(request *GetGroupActiveInfoRequest) (_result *GetGroupActiveInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetGroupActiveInfoHeaders{}
	_result = &GetGroupActiveInfoResponse{}
	_body, _err := client.GetGroupActiveInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInActiveUserListWithOptions(request *GetInActiveUserListRequest, headers *GetInActiveUserListHeaders, runtime *util.RuntimeOptions) (_result *GetInActiveUserListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptIds)) {
		body["deptIds"] = request.DeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		body["statDate"] = request.StatDate
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
		Action:      tea.String("GetInActiveUserList"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/inactives/users/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInActiveUserListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInActiveUserList(request *GetInActiveUserListRequest) (_result *GetInActiveUserListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetInActiveUserListHeaders{}
	_result = &GetInActiveUserListResponse{}
	_body, _err := client.GetInActiveUserListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLastOrgAuthDataWithOptions(request *GetLastOrgAuthDataRequest, headers *GetLastOrgAuthDataHeaders, runtime *util.RuntimeOptions) (_result *GetLastOrgAuthDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TargetCorpId)) {
		query["targetCorpId"] = request.TargetCorpId
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
		Action:      tea.String("GetLastOrgAuthData"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/organizations/authInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLastOrgAuthDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLastOrgAuthData(request *GetLastOrgAuthDataRequest) (_result *GetLastOrgAuthDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetLastOrgAuthDataHeaders{}
	_result = &GetLastOrgAuthDataResponse{}
	_body, _err := client.GetLastOrgAuthDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOaOperatorLogListWithOptions(request *GetOaOperatorLogListRequest, headers *GetOaOperatorLogListHeaders, runtime *util.RuntimeOptions) (_result *GetOaOperatorLogListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryList)) {
		body["categoryList"] = request.CategoryList
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
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
		Action:      tea.String("GetOaOperatorLogList"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/oaOperatorLogs/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOaOperatorLogListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOaOperatorLogList(request *GetOaOperatorLogListRequest) (_result *GetOaOperatorLogListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOaOperatorLogListHeaders{}
	_result = &GetOaOperatorLogListResponse{}
	_body, _err := client.GetOaOperatorLogListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOutGroupsByPageWithOptions(request *GetOutGroupsByPageRequest, headers *GetOutGroupsByPageHeaders, runtime *util.RuntimeOptions) (_result *GetOutGroupsByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
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
		Action:      tea.String("GetOutGroupsByPage"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/audits/outsideGroups/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOutGroupsByPageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOutGroupsByPage(request *GetOutGroupsByPageRequest) (_result *GetOutGroupsByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOutGroupsByPageHeaders{}
	_result = &GetOutGroupsByPageResponse{}
	_body, _err := client.GetOutGroupsByPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOutsideAuditGroupMessageByPageWithOptions(request *GetOutsideAuditGroupMessageByPageRequest, headers *GetOutsideAuditGroupMessageByPageHeaders, runtime *util.RuntimeOptions) (_result *GetOutsideAuditGroupMessageByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
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
		Action:      tea.String("GetOutsideAuditGroupMessageByPage"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/audits/outsideGroups/messages/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOutsideAuditGroupMessageByPageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOutsideAuditGroupMessageByPage(request *GetOutsideAuditGroupMessageByPageRequest) (_result *GetOutsideAuditGroupMessageByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOutsideAuditGroupMessageByPageHeaders{}
	_result = &GetOutsideAuditGroupMessageByPageResponse{}
	_body, _err := client.GetOutsideAuditGroupMessageByPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPartnerTypeByParentIdWithOptions(parentId *string, headers *GetPartnerTypeByParentIdHeaders, runtime *util.RuntimeOptions) (_result *GetPartnerTypeByParentIdResponse, _err error) {
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
		Action:      tea.String("GetPartnerTypeByParentId"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/partnerLabels/" + tea.StringValue(parentId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPartnerTypeByParentIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPartnerTypeByParentId(parentId *string) (_result *GetPartnerTypeByParentIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPartnerTypeByParentIdHeaders{}
	_result = &GetPartnerTypeByParentIdResponse{}
	_body, _err := client.GetPartnerTypeByParentIdWithOptions(parentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPublicDevicesWithOptions(request *GetPublicDevicesRequest, headers *GetPublicDevicesHeaders, runtime *util.RuntimeOptions) (_result *GetPublicDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MacAddress)) {
		query["macAddress"] = request.MacAddress
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		query["platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["title"] = request.Title
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
		Action:      tea.String("GetPublicDevices"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/trusts/publicDevices"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPublicDevicesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPublicDevices(request *GetPublicDevicesRequest) (_result *GetPublicDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPublicDevicesHeaders{}
	_result = &GetPublicDevicesResponse{}
	_body, _err := client.GetPublicDevicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPublisherSummaryWithOptions(dataId *string, request *GetPublisherSummaryRequest, headers *GetPublisherSummaryHeaders, runtime *util.RuntimeOptions) (_result *GetPublisherSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
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
		Action:      tea.String("GetPublisherSummary"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/data/publisher/" + tea.StringValue(dataId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPublisherSummaryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPublisherSummary(dataId *string, request *GetPublisherSummaryRequest) (_result *GetPublisherSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPublisherSummaryHeaders{}
	_result = &GetPublisherSummaryResponse{}
	_body, _err := client.GetPublisherSummaryWithOptions(dataId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRealPeopleRecordsWithOptions(request *GetRealPeopleRecordsRequest, headers *GetRealPeopleRecordsHeaders, runtime *util.RuntimeOptions) (_result *GetRealPeopleRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		body["agentId"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.FromTime)) {
		body["fromTime"] = request.FromTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PersonIdentification)) {
		body["personIdentification"] = request.PersonIdentification
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.ToTime)) {
		body["toTime"] = request.ToTime
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
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
		Action:      tea.String("GetRealPeopleRecords"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/persons/identificationRecords/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRealPeopleRecordsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRealPeopleRecords(request *GetRealPeopleRecordsRequest) (_result *GetRealPeopleRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetRealPeopleRecordsHeaders{}
	_result = &GetRealPeopleRecordsResponse{}
	_body, _err := client.GetRealPeopleRecordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRecognizeRecordsWithOptions(request *GetRecognizeRecordsRequest, headers *GetRecognizeRecordsHeaders, runtime *util.RuntimeOptions) (_result *GetRecognizeRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		body["agentId"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.FaceCompareResult)) {
		body["faceCompareResult"] = request.FaceCompareResult
	}

	if !tea.BoolValue(util.IsUnset(request.FromTime)) {
		body["fromTime"] = request.FromTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ToTime)) {
		body["toTime"] = request.ToTime
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
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
		Action:      tea.String("GetRecognizeRecords"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/faces/recognizeRecords/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRecognizeRecordsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRecognizeRecords(request *GetRecognizeRecordsRequest) (_result *GetRecognizeRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetRecognizeRecordsHeaders{}
	_result = &GetRecognizeRecordsResponse{}
	_body, _err := client.GetRecognizeRecordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSignedDetailByPageWithOptions(request *GetSignedDetailByPageRequest, headers *GetSignedDetailByPageHeaders, runtime *util.RuntimeOptions) (_result *GetSignedDetailByPageResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SignStatus)) {
		query["signStatus"] = request.SignStatus
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
		Action:      tea.String("GetSignedDetailByPage"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/audits/users"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSignedDetailByPageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSignedDetailByPage(request *GetSignedDetailByPageRequest) (_result *GetSignedDetailByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSignedDetailByPageHeaders{}
	_result = &GetSignedDetailByPageResponse{}
	_body, _err := client.GetSignedDetailByPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTrustDeviceListWithOptions(request *GetTrustDeviceListRequest, headers *GetTrustDeviceListHeaders, runtime *util.RuntimeOptions) (_result *GetTrustDeviceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
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
		Action:      tea.String("GetTrustDeviceList"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/trustedDevices/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTrustDeviceListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTrustDeviceList(request *GetTrustDeviceListRequest) (_result *GetTrustDeviceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTrustDeviceListHeaders{}
	_result = &GetTrustDeviceListResponse{}
	_body, _err := client.GetTrustDeviceListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserAppVersionSummaryWithOptions(dataId *string, request *GetUserAppVersionSummaryRequest, headers *GetUserAppVersionSummaryHeaders, runtime *util.RuntimeOptions) (_result *GetUserAppVersionSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
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
		Action:      tea.String("GetUserAppVersionSummary"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/data/appVersion/org/" + tea.StringValue(dataId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserAppVersionSummaryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserAppVersionSummary(dataId *string, request *GetUserAppVersionSummaryRequest) (_result *GetUserAppVersionSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserAppVersionSummaryHeaders{}
	_result = &GetUserAppVersionSummaryResponse{}
	_body, _err := client.GetUserAppVersionSummaryWithOptions(dataId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserFaceStateWithOptions(request *GetUserFaceStateRequest, headers *GetUserFaceStateHeaders, runtime *util.RuntimeOptions) (_result *GetUserFaceStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
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
		Action:      tea.String("GetUserFaceState"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/faces/recognizeStates/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserFaceStateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserFaceState(request *GetUserFaceStateRequest) (_result *GetUserFaceStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserFaceStateHeaders{}
	_result = &GetUserFaceStateResponse{}
	_body, _err := client.GetUserFaceStateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserRealPeopleStateWithOptions(request *GetUserRealPeopleStateRequest, headers *GetUserRealPeopleStateHeaders, runtime *util.RuntimeOptions) (_result *GetUserRealPeopleStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
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
		Action:      tea.String("GetUserRealPeopleState"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/persons/identificationStates/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserRealPeopleStateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserRealPeopleState(request *GetUserRealPeopleStateRequest) (_result *GetUserRealPeopleStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserRealPeopleStateHeaders{}
	_result = &GetUserRealPeopleStateResponse{}
	_body, _err := client.GetUserRealPeopleStateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserStayLengthWithOptions(request *GetUserStayLengthRequest, headers *GetUserStayLengthHeaders, runtime *util.RuntimeOptions) (_result *GetUserStayLengthResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
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
		Action:      tea.String("GetUserStayLength"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/data/users/stayTimes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserStayLengthResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserStayLength(request *GetUserStayLengthRequest) (_result *GetUserStayLengthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserStayLengthHeaders{}
	_result = &GetUserStayLengthResponse{}
	_body, _err := client.GetUserStayLengthWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAuditLogWithOptions(request *ListAuditLogRequest, headers *ListAuditLogHeaders, runtime *util.RuntimeOptions) (_result *ListAuditLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["endDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.NextBizId)) {
		query["nextBizId"] = request.NextBizId
	}

	if !tea.BoolValue(util.IsUnset(request.NextGmtCreate)) {
		query["nextGmtCreate"] = request.NextGmtCreate
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["startDate"] = request.StartDate
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
		Action:      tea.String("ListAuditLog"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/fileAuditLogs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAuditLogResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAuditLog(request *ListAuditLogRequest) (_result *ListAuditLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListAuditLogHeaders{}
	_result = &ListAuditLogResponse{}
	_body, _err := client.ListAuditLogWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCategorysWithOptions(tmpReq *ListCategorysRequest, headers *ListCategorysHeaders, runtime *util.RuntimeOptions) (_result *ListCategorysResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListCategorysShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["body"] = request.BodyShrink
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
		Action:      tea.String("ListCategorys"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/messageCategories/categories/listCategories"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCategorysResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCategorys(request *ListCategorysRequest) (_result *ListCategorysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListCategorysHeaders{}
	_result = &ListCategorysResponse{}
	_body, _err := client.ListCategorysWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListJoinOrgInfoWithOptions(request *ListJoinOrgInfoRequest, headers *ListJoinOrgInfoHeaders, runtime *util.RuntimeOptions) (_result *ListJoinOrgInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["mobile"] = request.Mobile
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
		Action:      tea.String("ListJoinOrgInfo"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/exclusiveAccounts/organizations/infos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListJoinOrgInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListJoinOrgInfo(request *ListJoinOrgInfoRequest) (_result *ListJoinOrgInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListJoinOrgInfoHeaders{}
	_result = &ListJoinOrgInfoResponse{}
	_body, _err := client.ListJoinOrgInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMiniAppAvailableVersionWithOptions(request *ListMiniAppAvailableVersionRequest, headers *ListMiniAppAvailableVersionHeaders, runtime *util.RuntimeOptions) (_result *ListMiniAppAvailableVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		body["miniAppId"] = request.MiniAppId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.VersionTypeSet)) {
		body["versionTypeSet"] = request.VersionTypeSet
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
		Action:      tea.String("ListMiniAppAvailableVersion"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/miniApps/versions/availableLists"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMiniAppAvailableVersionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMiniAppAvailableVersion(request *ListMiniAppAvailableVersionRequest) (_result *ListMiniAppAvailableVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListMiniAppAvailableVersionHeaders{}
	_result = &ListMiniAppAvailableVersionResponse{}
	_body, _err := client.ListMiniAppAvailableVersionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMiniAppHistoryVersionWithOptions(request *ListMiniAppHistoryVersionRequest, headers *ListMiniAppHistoryVersionHeaders, runtime *util.RuntimeOptions) (_result *ListMiniAppHistoryVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		query["miniAppId"] = request.MiniAppId
	}

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
		Action:      tea.String("ListMiniAppHistoryVersion"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/miniApps/versions/historyLists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMiniAppHistoryVersionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMiniAppHistoryVersion(request *ListMiniAppHistoryVersionRequest) (_result *ListMiniAppHistoryVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListMiniAppHistoryVersionHeaders{}
	_result = &ListMiniAppHistoryVersionResponse{}
	_body, _err := client.ListMiniAppHistoryVersionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPartnerRolesWithOptions(parentId *string, headers *ListPartnerRolesHeaders, runtime *util.RuntimeOptions) (_result *ListPartnerRolesResponse, _err error) {
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
		Action:      tea.String("ListPartnerRoles"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/partners/roles/" + tea.StringValue(parentId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPartnerRolesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPartnerRoles(parentId *string) (_result *ListPartnerRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListPartnerRolesHeaders{}
	_result = &ListPartnerRolesResponse{}
	_body, _err := client.ListPartnerRolesWithOptions(parentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPunchScheduleByConditionWithPagingWithOptions(request *ListPunchScheduleByConditionWithPagingRequest, headers *ListPunchScheduleByConditionWithPagingHeaders, runtime *util.RuntimeOptions) (_result *ListPunchScheduleByConditionWithPagingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizInstanceId)) {
		body["bizInstanceId"] = request.BizInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleDateEnd)) {
		body["scheduleDateEnd"] = request.ScheduleDateEnd
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleDateStart)) {
		body["scheduleDateStart"] = request.ScheduleDateStart
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
		Action:      tea.String("ListPunchScheduleByConditionWithPaging"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/punchSchedules/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPunchScheduleByConditionWithPagingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPunchScheduleByConditionWithPaging(request *ListPunchScheduleByConditionWithPagingRequest) (_result *ListPunchScheduleByConditionWithPagingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListPunchScheduleByConditionWithPagingHeaders{}
	_result = &ListPunchScheduleByConditionWithPagingResponse{}
	_body, _err := client.ListPunchScheduleByConditionWithPagingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRulesWithOptions(tmpReq *ListRulesRequest, headers *ListRulesHeaders, runtime *util.RuntimeOptions) (_result *ListRulesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["body"] = request.BodyShrink
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
		Action:      tea.String("ListRules"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/messageCategories/rules/listRules"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRulesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRules(request *ListRulesRequest) (_result *ListRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListRulesHeaders{}
	_result = &ListRulesResponse{}
	_body, _err := client.ListRulesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LogoutWithOptions(request *LogoutRequest, headers *LogoutHeaders, runtime *util.RuntimeOptions) (_result *LogoutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceType)) {
		body["deviceType"] = request.DeviceType
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
		Action:      tea.String("Logout"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/users/logout"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &LogoutResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Logout(request *LogoutRequest) (_result *LogoutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &LogoutHeaders{}
	_result = &LogoutResponse{}
	_body, _err := client.LogoutWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishFileChangeNoticeWithOptions(request *PublishFileChangeNoticeRequest, headers *PublishFileChangeNoticeHeaders, runtime *util.RuntimeOptions) (_result *PublishFileChangeNoticeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["fileId"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.OperateType)) {
		body["operateType"] = request.OperateType
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUnionId)) {
		body["operatorUnionId"] = request.OperatorUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["spaceId"] = request.SpaceId
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
		Action:      tea.String("PublishFileChangeNotice"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/comments/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("none"),
	}
	_result = &PublishFileChangeNoticeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishFileChangeNotice(request *PublishFileChangeNoticeRequest) (_result *PublishFileChangeNoticeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PublishFileChangeNoticeHeaders{}
	_result = &PublishFileChangeNoticeResponse{}
	_body, _err := client.PublishFileChangeNoticeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishRuleWithOptions(request *PublishRuleRequest, headers *PublishRuleHeaders, runtime *util.RuntimeOptions) (_result *PublishRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
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
		Action:      tea.String("PublishRule"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/messageCategories/rules/publish"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishRuleResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishRule(request *PublishRuleRequest) (_result *PublishRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PublishRuleHeaders{}
	_result = &PublishRuleResponse{}
	_body, _err := client.PublishRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushBadgeWithOptions(request *PushBadgeRequest, headers *PushBadgeHeaders, runtime *util.RuntimeOptions) (_result *PushBadgeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		body["agentId"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.BadgeItems)) {
		body["badgeItems"] = request.BadgeItems
	}

	if !tea.BoolValue(util.IsUnset(request.PushType)) {
		body["pushType"] = request.PushType
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
		Action:      tea.String("PushBadge"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/exclusiveDesigns/redPoints/push"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PushBadgeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushBadge(request *PushBadgeRequest) (_result *PushBadgeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PushBadgeHeaders{}
	_result = &PushBadgeResponse{}
	_body, _err := client.PushBadgeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAcrossCloudStroageConfigsWithOptions(request *QueryAcrossCloudStroageConfigsRequest, headers *QueryAcrossCloudStroageConfigsHeaders, runtime *util.RuntimeOptions) (_result *QueryAcrossCloudStroageConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TargetCloudType)) {
		query["targetCloudType"] = request.TargetCloudType
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCorpId)) {
		query["targetCorpId"] = request.TargetCorpId
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
		Action:      tea.String("QueryAcrossCloudStroageConfigs"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/fileStorages/acrossClouds/configurations"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAcrossCloudStroageConfigsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAcrossCloudStroageConfigs(request *QueryAcrossCloudStroageConfigsRequest) (_result *QueryAcrossCloudStroageConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAcrossCloudStroageConfigsHeaders{}
	_result = &QueryAcrossCloudStroageConfigsResponse{}
	_body, _err := client.QueryAcrossCloudStroageConfigsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPartnerInfoWithOptions(userId *string, headers *QueryPartnerInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryPartnerInfoResponse, _err error) {
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
		Action:      tea.String("QueryPartnerInfo"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/partners/users/" + tea.StringValue(userId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPartnerInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPartnerInfo(userId *string) (_result *QueryPartnerInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryPartnerInfoHeaders{}
	_result = &QueryPartnerInfoResponse{}
	_body, _err := client.QueryPartnerInfoWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserBehaviorWithOptions(request *QueryUserBehaviorRequest, headers *QueryUserBehaviorHeaders, runtime *util.RuntimeOptions) (_result *QueryUserBehaviorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		body["platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
		Action:      tea.String("QueryUserBehavior"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/enterpriseSecurities/userBehaviors/screenshots/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserBehaviorResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserBehavior(request *QueryUserBehaviorRequest) (_result *QueryUserBehaviorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserBehaviorHeaders{}
	_result = &QueryUserBehaviorResponse{}
	_body, _err := client.QueryUserBehaviorWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RollbackMiniAppVersionWithOptions(request *RollbackMiniAppVersionRequest, headers *RollbackMiniAppVersionHeaders, runtime *util.RuntimeOptions) (_result *RollbackMiniAppVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		body["miniAppId"] = request.MiniAppId
	}

	if !tea.BoolValue(util.IsUnset(request.RollbackVersion)) {
		body["rollbackVersion"] = request.RollbackVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TargetVersion)) {
		body["targetVersion"] = request.TargetVersion
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
		Action:      tea.String("RollbackMiniAppVersion"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/miniApps/versions/rollback"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RollbackMiniAppVersionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RollbackMiniAppVersion(request *RollbackMiniAppVersionRequest) (_result *RollbackMiniAppVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RollbackMiniAppVersionHeaders{}
	_result = &RollbackMiniAppVersionResponse{}
	_body, _err := client.RollbackMiniAppVersionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveAcrossCloudStroageConfigsWithOptions(request *SaveAcrossCloudStroageConfigsRequest, headers *SaveAcrossCloudStroageConfigsHeaders, runtime *util.RuntimeOptions) (_result *SaveAcrossCloudStroageConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKeyId)) {
		body["accessKeyId"] = request.AccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessKeySecret)) {
		body["accessKeySecret"] = request.AccessKeySecret
	}

	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		body["bucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.CloudType)) {
		body["cloudType"] = request.CloudType
	}

	if !tea.BoolValue(util.IsUnset(request.Endpoint)) {
		body["endpoint"] = request.Endpoint
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCorpId)) {
		body["targetCorpId"] = request.TargetCorpId
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
		Action:      tea.String("SaveAcrossCloudStroageConfigs"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/fileStorages/acrossClouds/configurations"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveAcrossCloudStroageConfigsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveAcrossCloudStroageConfigs(request *SaveAcrossCloudStroageConfigsRequest) (_result *SaveAcrossCloudStroageConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveAcrossCloudStroageConfigsHeaders{}
	_result = &SaveAcrossCloudStroageConfigsResponse{}
	_body, _err := client.SaveAcrossCloudStroageConfigsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveAndSubmitAuthInfoWithOptions(request *SaveAndSubmitAuthInfoRequest, headers *SaveAndSubmitAuthInfoHeaders, runtime *util.RuntimeOptions) (_result *SaveAndSubmitAuthInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyRemark)) {
		body["applyRemark"] = request.ApplyRemark
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizeMediaId)) {
		body["authorizeMediaId"] = request.AuthorizeMediaId
	}

	if !tea.BoolValue(util.IsUnset(request.Industry)) {
		body["industry"] = request.Industry
	}

	if !tea.BoolValue(util.IsUnset(request.LegalPerson)) {
		body["legalPerson"] = request.LegalPerson
	}

	if !tea.BoolValue(util.IsUnset(request.LegalPersonIdCard)) {
		body["legalPersonIdCard"] = request.LegalPersonIdCard
	}

	if !tea.BoolValue(util.IsUnset(request.LicenseMediaId)) {
		body["licenseMediaId"] = request.LicenseMediaId
	}

	if !tea.BoolValue(util.IsUnset(request.LocCity)) {
		body["locCity"] = request.LocCity
	}

	if !tea.BoolValue(util.IsUnset(request.LocCityName)) {
		body["locCityName"] = request.LocCityName
	}

	if !tea.BoolValue(util.IsUnset(request.LocProvince)) {
		body["locProvince"] = request.LocProvince
	}

	if !tea.BoolValue(util.IsUnset(request.LocProvinceName)) {
		body["locProvinceName"] = request.LocProvinceName
	}

	if !tea.BoolValue(util.IsUnset(request.MobileNum)) {
		body["mobileNum"] = request.MobileNum
	}

	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		body["orgName"] = request.OrgName
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationCode)) {
		body["organizationCode"] = request.OrganizationCode
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationCodeMediaId)) {
		body["organizationCodeMediaId"] = request.OrganizationCodeMediaId
	}

	if !tea.BoolValue(util.IsUnset(request.RegistLocation)) {
		body["registLocation"] = request.RegistLocation
	}

	if !tea.BoolValue(util.IsUnset(request.RegistNum)) {
		body["registNum"] = request.RegistNum
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCorpId)) {
		body["targetCorpId"] = request.TargetCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.UnifiedSocialCredit)) {
		body["unifiedSocialCredit"] = request.UnifiedSocialCredit
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
		Action:      tea.String("SaveAndSubmitAuthInfo"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/ognizations/authInfos/saveAndSubmit"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveAndSubmitAuthInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveAndSubmitAuthInfo(request *SaveAndSubmitAuthInfoRequest) (_result *SaveAndSubmitAuthInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveAndSubmitAuthInfoHeaders{}
	_result = &SaveAndSubmitAuthInfoResponse{}
	_body, _err := client.SaveAndSubmitAuthInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveOpenTerminalInfoWithOptions(request *SaveOpenTerminalInfoRequest, headers *SaveOpenTerminalInfoHeaders, runtime *util.RuntimeOptions) (_result *SaveOpenTerminalInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.LogSource)) {
		body["logSource"] = request.LogSource
	}

	if !tea.BoolValue(util.IsUnset(request.LogType)) {
		body["logType"] = request.LogType
	}

	if !tea.BoolValue(util.IsUnset(request.OpenExt)) {
		body["openExt"] = request.OpenExt
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
		Action:      tea.String("SaveOpenTerminalInfo"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/externalLogs/terminalInfos/save"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveOpenTerminalInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveOpenTerminalInfo(request *SaveOpenTerminalInfoRequest) (_result *SaveOpenTerminalInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveOpenTerminalInfoHeaders{}
	_result = &SaveOpenTerminalInfoResponse{}
	_body, _err := client.SaveOpenTerminalInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveWhiteAppWithOptions(request *SaveWhiteAppRequest, headers *SaveWhiteAppHeaders, runtime *util.RuntimeOptions) (_result *SaveWhiteAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentIdList)) {
		body["agentIdList"] = request.AgentIdList
	}

	if !tea.BoolValue(util.IsUnset(request.Operation)) {
		body["operation"] = request.Operation
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
		Action:      tea.String("SaveWhiteApp"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/miniApps/whiteLists/save"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveWhiteAppResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveWhiteApp(request *SaveWhiteAppRequest) (_result *SaveWhiteAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveWhiteAppHeaders{}
	_result = &SaveWhiteAppResponse{}
	_body, _err := client.SaveWhiteAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchOrgInnerGroupInfoWithOptions(request *SearchOrgInnerGroupInfoRequest, headers *SearchOrgInnerGroupInfoHeaders, runtime *util.RuntimeOptions) (_result *SearchOrgInnerGroupInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateTimeEnd)) {
		query["createTimeEnd"] = request.CreateTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimeStart)) {
		query["createTimeStart"] = request.CreateTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.GroupMembersCountEnd)) {
		query["groupMembersCountEnd"] = request.GroupMembersCountEnd
	}

	if !tea.BoolValue(util.IsUnset(request.GroupMembersCountStart)) {
		query["groupMembersCountStart"] = request.GroupMembersCountStart
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupOwner)) {
		query["groupOwner"] = request.GroupOwner
	}

	if !tea.BoolValue(util.IsUnset(request.LastActiveTimeEnd)) {
		query["lastActiveTimeEnd"] = request.LastActiveTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.LastActiveTimeStart)) {
		query["lastActiveTimeStart"] = request.LastActiveTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		query["operatorUserId"] = request.OperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PageStart)) {
		query["pageStart"] = request.PageStart
	}

	if !tea.BoolValue(util.IsUnset(request.SyncToDingpan)) {
		query["syncToDingpan"] = request.SyncToDingpan
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["uuid"] = request.Uuid
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
		Action:      tea.String("SearchOrgInnerGroupInfo"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/securities/orgGroupInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchOrgInnerGroupInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchOrgInnerGroupInfo(request *SearchOrgInnerGroupInfoRequest) (_result *SearchOrgInnerGroupInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchOrgInnerGroupInfoHeaders{}
	_result = &SearchOrgInnerGroupInfoResponse{}
	_body, _err := client.SearchOrgInnerGroupInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendAppDingWithOptions(request *SendAppDingRequest, headers *SendAppDingHeaders, runtime *util.RuntimeOptions) (_result *SendAppDingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Userids)) {
		body["userids"] = request.Userids
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
		Action:      tea.String("SendAppDing"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/appDings/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("none"),
	}
	_result = &SendAppDingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendAppDing(request *SendAppDingRequest) (_result *SendAppDingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendAppDingHeaders{}
	_result = &SendAppDingResponse{}
	_body, _err := client.SendAppDingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendInvitationWithOptions(request *SendInvitationRequest, headers *SendInvitationHeaders, runtime *util.RuntimeOptions) (_result *SendInvitationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgAlias)) {
		body["orgAlias"] = request.OrgAlias
	}

	if !tea.BoolValue(util.IsUnset(request.PartnerLabelId)) {
		body["partnerLabelId"] = request.PartnerLabelId
	}

	if !tea.BoolValue(util.IsUnset(request.PartnerNum)) {
		body["partnerNum"] = request.PartnerNum
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		body["phone"] = request.Phone
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
		Action:      tea.String("SendInvitation"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/partnerDepartments/invitations/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("none"),
	}
	_result = &SendInvitationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendInvitation(request *SendInvitationRequest) (_result *SendInvitationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendInvitationHeaders{}
	_result = &SendInvitationResponse{}
	_body, _err := client.SendInvitationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendPhoneDingWithOptions(request *SendPhoneDingRequest, headers *SendPhoneDingHeaders, runtime *util.RuntimeOptions) (_result *SendPhoneDingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Userids)) {
		body["userids"] = request.Userids
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
		Action:      tea.String("SendPhoneDing"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/phoneDings/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendPhoneDingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendPhoneDing(request *SendPhoneDingRequest) (_result *SendPhoneDingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendPhoneDingHeaders{}
	_result = &SendPhoneDingResponse{}
	_body, _err := client.SendPhoneDingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDeptPartnerTypeAndNumWithOptions(request *SetDeptPartnerTypeAndNumRequest, headers *SetDeptPartnerTypeAndNumHeaders, runtime *util.RuntimeOptions) (_result *SetDeptPartnerTypeAndNumResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.LabelIds)) {
		body["labelIds"] = request.LabelIds
	}

	if !tea.BoolValue(util.IsUnset(request.PartnerNum)) {
		body["partnerNum"] = request.PartnerNum
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
		Action:      tea.String("SetDeptPartnerTypeAndNum"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/partnerDepartments"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDeptPartnerTypeAndNumResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDeptPartnerTypeAndNum(request *SetDeptPartnerTypeAndNumRequest) (_result *SetDeptPartnerTypeAndNumResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetDeptPartnerTypeAndNumHeaders{}
	_result = &SetDeptPartnerTypeAndNumResponse{}
	_body, _err := client.SetDeptPartnerTypeAndNumWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCategoryNameWithOptions(request *UpdateCategoryNameRequest, headers *UpdateCategoryNameHeaders, runtime *util.RuntimeOptions) (_result *UpdateCategoryNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentCategoryName)) {
		body["currentCategoryName"] = request.CurrentCategoryName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCategoryName)) {
		body["targetCategoryName"] = request.TargetCategoryName
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
		Action:      tea.String("UpdateCategoryName"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/messageCategories/categories/names"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCategoryNameResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCategoryName(request *UpdateCategoryNameRequest) (_result *UpdateCategoryNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateCategoryNameHeaders{}
	_result = &UpdateCategoryNameResponse{}
	_body, _err := client.UpdateCategoryNameWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFileStatusWithOptions(request *UpdateFileStatusRequest, headers *UpdateFileStatusHeaders, runtime *util.RuntimeOptions) (_result *UpdateFileStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RequestIds)) {
		body["requestIds"] = request.RequestIds
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
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
		Action:      tea.String("UpdateFileStatus"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/sending/files/status"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFileStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFileStatus(request *UpdateFileStatusRequest) (_result *UpdateFileStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateFileStatusHeaders{}
	_result = &UpdateFileStatusResponse{}
	_body, _err := client.UpdateFileStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMiniAppVersionStatusWithOptions(request *UpdateMiniAppVersionStatusRequest, headers *UpdateMiniAppVersionStatusHeaders, runtime *util.RuntimeOptions) (_result *UpdateMiniAppVersionStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MiniAppId)) {
		body["miniAppId"] = request.MiniAppId
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["version"] = request.Version
	}

	if !tea.BoolValue(util.IsUnset(request.VersionType)) {
		body["versionType"] = request.VersionType
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
		Action:      tea.String("UpdateMiniAppVersionStatus"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/miniApps/versions/versionStatus"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMiniAppVersionStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMiniAppVersionStatus(request *UpdateMiniAppVersionStatusRequest) (_result *UpdateMiniAppVersionStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateMiniAppVersionStatusHeaders{}
	_result = &UpdateMiniAppVersionStatusResponse{}
	_body, _err := client.UpdateMiniAppVersionStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePartnerVisibilityWithOptions(request *UpdatePartnerVisibilityRequest, headers *UpdatePartnerVisibilityHeaders, runtime *util.RuntimeOptions) (_result *UpdatePartnerVisibilityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptIds)) {
		body["deptIds"] = request.DeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.LabelId)) {
		body["labelId"] = request.LabelId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
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
		Action:      tea.String("UpdatePartnerVisibility"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/partnerDepartments/visibilityPartners"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("boolean"),
	}
	_result = &UpdatePartnerVisibilityResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePartnerVisibility(request *UpdatePartnerVisibilityRequest) (_result *UpdatePartnerVisibilityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdatePartnerVisibilityHeaders{}
	_result = &UpdatePartnerVisibilityResponse{}
	_body, _err := client.UpdatePartnerVisibilityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRoleVisibilityWithOptions(request *UpdateRoleVisibilityRequest, headers *UpdateRoleVisibilityHeaders, runtime *util.RuntimeOptions) (_result *UpdateRoleVisibilityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptIds)) {
		body["deptIds"] = request.DeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.LabelId)) {
		body["labelId"] = request.LabelId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
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
		Action:      tea.String("UpdateRoleVisibility"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/partnerDepartments/visibilityRoles"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("boolean"),
	}
	_result = &UpdateRoleVisibilityResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRoleVisibility(request *UpdateRoleVisibilityRequest) (_result *UpdateRoleVisibilityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateRoleVisibilityHeaders{}
	_result = &UpdateRoleVisibilityResponse{}
	_body, _err := client.UpdateRoleVisibilityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateStorageModeWithOptions(request *UpdateStorageModeRequest, headers *UpdateStorageModeHeaders, runtime *util.RuntimeOptions) (_result *UpdateStorageModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileStorageMode)) {
		body["fileStorageMode"] = request.FileStorageMode
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCorpId)) {
		body["targetCorpId"] = request.TargetCorpId
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
		Action:      tea.String("UpdateStorageMode"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/fileStorages/acrossClouds/storageModes"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateStorageModeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateStorageMode(request *UpdateStorageModeRequest) (_result *UpdateStorageModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateStorageModeHeaders{}
	_result = &UpdateStorageModeResponse{}
	_body, _err := client.UpdateStorageModeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
