// This file is auto-generated, don't edit it. Thanks.
package yida_2_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BatchAddOrUpdateRoleMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchAddOrUpdateRoleMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchAddOrUpdateRoleMembersHeaders) GoString() string {
	return s.String()
}

func (s *BatchAddOrUpdateRoleMembersHeaders) SetCommonHeaders(v map[string]*string) *BatchAddOrUpdateRoleMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchAddOrUpdateRoleMembersHeaders) SetXAcsDingtalkAccessToken(v string) *BatchAddOrUpdateRoleMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchAddOrUpdateRoleMembersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dingxxxx
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"memberId":"5014533041684xx","manageScopes":"8360866xx,430181xx,429821xx"},{"memberId":"014329103xx","manageScopes":"all"}]
	MembersInfo *string `json:"membersInfo,omitempty" xml:"membersInfo,omitempty"`
	PageNumber  *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize    *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ROLE-71dc7859-17df-490d-a93d-eb5506e31f42
	RoleUuid *string `json:"roleUuid,omitempty" xml:"roleUuid,omitempty"`
	// This parameter is required.
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 501453
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchAddOrUpdateRoleMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchAddOrUpdateRoleMembersRequest) GoString() string {
	return s.String()
}

func (s *BatchAddOrUpdateRoleMembersRequest) SetCorpId(v string) *BatchAddOrUpdateRoleMembersRequest {
	s.CorpId = &v
	return s
}

func (s *BatchAddOrUpdateRoleMembersRequest) SetMembersInfo(v string) *BatchAddOrUpdateRoleMembersRequest {
	s.MembersInfo = &v
	return s
}

func (s *BatchAddOrUpdateRoleMembersRequest) SetPageNumber(v int32) *BatchAddOrUpdateRoleMembersRequest {
	s.PageNumber = &v
	return s
}

func (s *BatchAddOrUpdateRoleMembersRequest) SetPageSize(v int32) *BatchAddOrUpdateRoleMembersRequest {
	s.PageSize = &v
	return s
}

func (s *BatchAddOrUpdateRoleMembersRequest) SetRoleUuid(v string) *BatchAddOrUpdateRoleMembersRequest {
	s.RoleUuid = &v
	return s
}

func (s *BatchAddOrUpdateRoleMembersRequest) SetToken(v string) *BatchAddOrUpdateRoleMembersRequest {
	s.Token = &v
	return s
}

func (s *BatchAddOrUpdateRoleMembersRequest) SetUserId(v string) *BatchAddOrUpdateRoleMembersRequest {
	s.UserId = &v
	return s
}

type BatchAddOrUpdateRoleMembersResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchAddOrUpdateRoleMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchAddOrUpdateRoleMembersResponseBody) GoString() string {
	return s.String()
}

func (s *BatchAddOrUpdateRoleMembersResponseBody) SetSuccess(v bool) *BatchAddOrUpdateRoleMembersResponseBody {
	s.Success = &v
	return s
}

type BatchAddOrUpdateRoleMembersResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchAddOrUpdateRoleMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchAddOrUpdateRoleMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchAddOrUpdateRoleMembersResponse) GoString() string {
	return s.String()
}

func (s *BatchAddOrUpdateRoleMembersResponse) SetHeaders(v map[string]*string) *BatchAddOrUpdateRoleMembersResponse {
	s.Headers = v
	return s
}

func (s *BatchAddOrUpdateRoleMembersResponse) SetStatusCode(v int32) *BatchAddOrUpdateRoleMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchAddOrUpdateRoleMembersResponse) SetBody(v *BatchAddOrUpdateRoleMembersResponseBody) *BatchAddOrUpdateRoleMembersResponse {
	s.Body = v
	return s
}

type BatchDeleteRoleMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchDeleteRoleMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteRoleMembersHeaders) GoString() string {
	return s.String()
}

func (s *BatchDeleteRoleMembersHeaders) SetCommonHeaders(v map[string]*string) *BatchDeleteRoleMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchDeleteRoleMembersHeaders) SetXAcsDingtalkAccessToken(v string) *BatchDeleteRoleMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchDeleteRoleMembersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dingxxxx
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// This parameter is required.
	MemberIds  *string `json:"memberIds,omitempty" xml:"memberIds,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ROLE-71dc7859-17df-490d-a93d-eb5506e31f42
	RoleUuid *string `json:"roleUuid,omitempty" xml:"roleUuid,omitempty"`
	// This parameter is required.
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 501453
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchDeleteRoleMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteRoleMembersRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteRoleMembersRequest) SetCorpId(v string) *BatchDeleteRoleMembersRequest {
	s.CorpId = &v
	return s
}

func (s *BatchDeleteRoleMembersRequest) SetMemberIds(v string) *BatchDeleteRoleMembersRequest {
	s.MemberIds = &v
	return s
}

func (s *BatchDeleteRoleMembersRequest) SetPageNumber(v int32) *BatchDeleteRoleMembersRequest {
	s.PageNumber = &v
	return s
}

func (s *BatchDeleteRoleMembersRequest) SetPageSize(v int32) *BatchDeleteRoleMembersRequest {
	s.PageSize = &v
	return s
}

func (s *BatchDeleteRoleMembersRequest) SetRoleUuid(v string) *BatchDeleteRoleMembersRequest {
	s.RoleUuid = &v
	return s
}

func (s *BatchDeleteRoleMembersRequest) SetToken(v string) *BatchDeleteRoleMembersRequest {
	s.Token = &v
	return s
}

func (s *BatchDeleteRoleMembersRequest) SetUserId(v string) *BatchDeleteRoleMembersRequest {
	s.UserId = &v
	return s
}

type BatchDeleteRoleMembersResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchDeleteRoleMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteRoleMembersResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteRoleMembersResponseBody) SetSuccess(v bool) *BatchDeleteRoleMembersResponseBody {
	s.Success = &v
	return s
}

type BatchDeleteRoleMembersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteRoleMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteRoleMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteRoleMembersResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteRoleMembersResponse) SetHeaders(v map[string]*string) *BatchDeleteRoleMembersResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteRoleMembersResponse) SetStatusCode(v int32) *BatchDeleteRoleMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteRoleMembersResponse) SetBody(v *BatchDeleteRoleMembersResponseBody) *BatchDeleteRoleMembersResponse {
	s.Body = v
	return s
}

type CancelAgentTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CancelAgentTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s CancelAgentTaskHeaders) GoString() string {
	return s.String()
}

func (s *CancelAgentTaskHeaders) SetCommonHeaders(v map[string]*string) *CancelAgentTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CancelAgentTaskHeaders) SetXAcsDingtalkAccessToken(v string) *CancelAgentTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CancelAgentTaskRequest struct {
	AgentType *string `json:"agentType,omitempty" xml:"agentType,omitempty"`
	// This parameter is required.
	AgentUuid *string `json:"agentUuid,omitempty" xml:"agentUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dingxxxx
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// This parameter is required.
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 501453
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CancelAgentTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelAgentTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelAgentTaskRequest) SetAgentType(v string) *CancelAgentTaskRequest {
	s.AgentType = &v
	return s
}

func (s *CancelAgentTaskRequest) SetAgentUuid(v string) *CancelAgentTaskRequest {
	s.AgentUuid = &v
	return s
}

func (s *CancelAgentTaskRequest) SetCorpId(v string) *CancelAgentTaskRequest {
	s.CorpId = &v
	return s
}

func (s *CancelAgentTaskRequest) SetToken(v string) *CancelAgentTaskRequest {
	s.Token = &v
	return s
}

func (s *CancelAgentTaskRequest) SetUserId(v string) *CancelAgentTaskRequest {
	s.UserId = &v
	return s
}

type CancelAgentTaskResponseBody struct {
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CancelAgentTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelAgentTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelAgentTaskResponseBody) SetErrorCode(v string) *CancelAgentTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CancelAgentTaskResponseBody) SetErrorMsg(v string) *CancelAgentTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CancelAgentTaskResponseBody) SetResult(v bool) *CancelAgentTaskResponseBody {
	s.Result = &v
	return s
}

func (s *CancelAgentTaskResponseBody) SetSuccess(v bool) *CancelAgentTaskResponseBody {
	s.Success = &v
	return s
}

type CancelAgentTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelAgentTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelAgentTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelAgentTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelAgentTaskResponse) SetHeaders(v map[string]*string) *CancelAgentTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelAgentTaskResponse) SetStatusCode(v int32) *CancelAgentTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelAgentTaskResponse) SetBody(v *CancelAgentTaskResponseBody) *CancelAgentTaskResponse {
	s.Body = v
	return s
}

type CreateAgentTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateAgentTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateAgentTaskHeaders) GoString() string {
	return s.String()
}

func (s *CreateAgentTaskHeaders) SetCommonHeaders(v map[string]*string) *CreateAgentTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateAgentTaskHeaders) SetXAcsDingtalkAccessToken(v string) *CreateAgentTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateAgentTaskRequest struct {
	// example:
	//
	// EXECUTE
	AgentCategory *string `json:"agentCategory,omitempty" xml:"agentCategory,omitempty"`
	// example:
	//
	// ALL
	AgentRangeType *string `json:"agentRangeType,omitempty" xml:"agentRangeType,omitempty"`
	// example:
	//
	// [{\"appType\":\"APP_xxx\",\"formUuid\":\"FORM-xxx\"}]
	AgentRangeValue *string `json:"agentRangeValue,omitempty" xml:"agentRangeValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALL
	AgentType *string `json:"agentType,omitempty" xml:"agentType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10001
	AgentUserId *string `json:"agentUserId,omitempty" xml:"agentUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dingxxxx
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// 1761204600404
	EndTimestamp        *string `json:"endTimestamp,omitempty" xml:"endTimestamp,omitempty"`
	NeedNoticePrincipal *string `json:"needNoticePrincipal,omitempty" xml:"needNoticePrincipal,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10002
	PrincipalUserId *string `json:"principalUserId,omitempty" xml:"principalUserId,omitempty"`
	// example:
	//
	// 1761204600404
	StartTimestamp *string `json:"startTimestamp,omitempty" xml:"startTimestamp,omitempty"`
	// This parameter is required.
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 501453
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateAgentTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAgentTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentTaskRequest) SetAgentCategory(v string) *CreateAgentTaskRequest {
	s.AgentCategory = &v
	return s
}

func (s *CreateAgentTaskRequest) SetAgentRangeType(v string) *CreateAgentTaskRequest {
	s.AgentRangeType = &v
	return s
}

func (s *CreateAgentTaskRequest) SetAgentRangeValue(v string) *CreateAgentTaskRequest {
	s.AgentRangeValue = &v
	return s
}

func (s *CreateAgentTaskRequest) SetAgentType(v string) *CreateAgentTaskRequest {
	s.AgentType = &v
	return s
}

func (s *CreateAgentTaskRequest) SetAgentUserId(v string) *CreateAgentTaskRequest {
	s.AgentUserId = &v
	return s
}

func (s *CreateAgentTaskRequest) SetCorpId(v string) *CreateAgentTaskRequest {
	s.CorpId = &v
	return s
}

func (s *CreateAgentTaskRequest) SetEndTimestamp(v string) *CreateAgentTaskRequest {
	s.EndTimestamp = &v
	return s
}

func (s *CreateAgentTaskRequest) SetNeedNoticePrincipal(v string) *CreateAgentTaskRequest {
	s.NeedNoticePrincipal = &v
	return s
}

func (s *CreateAgentTaskRequest) SetPrincipalUserId(v string) *CreateAgentTaskRequest {
	s.PrincipalUserId = &v
	return s
}

func (s *CreateAgentTaskRequest) SetStartTimestamp(v string) *CreateAgentTaskRequest {
	s.StartTimestamp = &v
	return s
}

func (s *CreateAgentTaskRequest) SetToken(v string) *CreateAgentTaskRequest {
	s.Token = &v
	return s
}

func (s *CreateAgentTaskRequest) SetUserId(v string) *CreateAgentTaskRequest {
	s.UserId = &v
	return s
}

type CreateAgentTaskResponseBody struct {
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result    *string `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateAgentTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAgentTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentTaskResponseBody) SetErrorCode(v string) *CreateAgentTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateAgentTaskResponseBody) SetErrorMsg(v string) *CreateAgentTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateAgentTaskResponseBody) SetResult(v string) *CreateAgentTaskResponseBody {
	s.Result = &v
	return s
}

func (s *CreateAgentTaskResponseBody) SetSuccess(v bool) *CreateAgentTaskResponseBody {
	s.Success = &v
	return s
}

type CreateAgentTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAgentTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgentTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAgentTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateAgentTaskResponse) SetHeaders(v map[string]*string) *CreateAgentTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateAgentTaskResponse) SetStatusCode(v int32) *CreateAgentTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgentTaskResponse) SetBody(v *CreateAgentTaskResponseBody) *CreateAgentTaskResponse {
	s.Body = v
	return s
}

type CreateOrUpdateFormDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateOrUpdateFormDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateFormDataHeaders) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateFormDataHeaders) SetCommonHeaders(v map[string]*string) *CreateOrUpdateFormDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateOrUpdateFormDataHeaders) SetXAcsDingtalkAccessToken(v string) *CreateOrUpdateFormDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateOrUpdateFormDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_XCE0EVXS6DYG3YDYC5RD
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"countrySelectField_l0c1cwiu":[{"value":"US"}]}
	FormDataJson *string `json:"formDataJson,omitempty" xml:"formDataJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-GX866MC1NC1VOFF6WVQW33FD16E23L3CPMKVKA
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// false
	NoExecuteExpression *bool `json:"noExecuteExpression,omitempty" xml:"noExecuteExpression,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"key":"currentNodeName","value":"当前审批节点名称","type":"TEXT","operator":"like","componentName":"TextField"}]。详情参考 https://www.yuque.com/yida/support/agb8im#F4S8e
	SearchCondition *string `json:"searchCondition,omitempty" xml:"searchCondition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 09866181UTZVVD4R3DC955FNKIM52HVPU5WWK7
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// example:
	//
	// false
	UseAlias *bool `json:"useAlias,omitempty" xml:"useAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ding173982232112232
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateOrUpdateFormDataRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateFormDataRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateFormDataRequest) SetAppType(v string) *CreateOrUpdateFormDataRequest {
	s.AppType = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) SetFormDataJson(v string) *CreateOrUpdateFormDataRequest {
	s.FormDataJson = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) SetFormUuid(v string) *CreateOrUpdateFormDataRequest {
	s.FormUuid = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) SetNoExecuteExpression(v bool) *CreateOrUpdateFormDataRequest {
	s.NoExecuteExpression = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) SetSearchCondition(v string) *CreateOrUpdateFormDataRequest {
	s.SearchCondition = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) SetSystemToken(v string) *CreateOrUpdateFormDataRequest {
	s.SystemToken = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) SetUseAlias(v bool) *CreateOrUpdateFormDataRequest {
	s.UseAlias = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) SetUserId(v string) *CreateOrUpdateFormDataRequest {
	s.UserId = &v
	return s
}

type CreateOrUpdateFormDataResponseBody struct {
	Result []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s CreateOrUpdateFormDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateFormDataResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateFormDataResponseBody) SetResult(v []*string) *CreateOrUpdateFormDataResponseBody {
	s.Result = v
	return s
}

type CreateOrUpdateFormDataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrUpdateFormDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrUpdateFormDataResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateFormDataResponse) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateFormDataResponse) SetHeaders(v map[string]*string) *CreateOrUpdateFormDataResponse {
	s.Headers = v
	return s
}

func (s *CreateOrUpdateFormDataResponse) SetStatusCode(v int32) *CreateOrUpdateFormDataResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrUpdateFormDataResponse) SetBody(v *CreateOrUpdateFormDataResponseBody) *CreateOrUpdateFormDataResponse {
	s.Body = v
	return s
}

type DeleteMatrixDataByRowIdsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteMatrixDataByRowIdsHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteMatrixDataByRowIdsHeaders) GoString() string {
	return s.String()
}

func (s *DeleteMatrixDataByRowIdsHeaders) SetCommonHeaders(v map[string]*string) *DeleteMatrixDataByRowIdsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteMatrixDataByRowIdsHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteMatrixDataByRowIdsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteMatrixDataByRowIdsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dingxxxx
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MATRIX-C8I4J40EM81XLWZH61ZK
	MatrixId *string `json:"matrixId,omitempty" xml:"matrixId,omitempty"`
	// This parameter is required.
	RowIds *string `json:"rowIds,omitempty" xml:"rowIds,omitempty"`
	// This parameter is required.
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 501453
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeleteMatrixDataByRowIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMatrixDataByRowIdsRequest) GoString() string {
	return s.String()
}

func (s *DeleteMatrixDataByRowIdsRequest) SetCorpId(v string) *DeleteMatrixDataByRowIdsRequest {
	s.CorpId = &v
	return s
}

func (s *DeleteMatrixDataByRowIdsRequest) SetMatrixId(v string) *DeleteMatrixDataByRowIdsRequest {
	s.MatrixId = &v
	return s
}

func (s *DeleteMatrixDataByRowIdsRequest) SetRowIds(v string) *DeleteMatrixDataByRowIdsRequest {
	s.RowIds = &v
	return s
}

func (s *DeleteMatrixDataByRowIdsRequest) SetToken(v string) *DeleteMatrixDataByRowIdsRequest {
	s.Token = &v
	return s
}

func (s *DeleteMatrixDataByRowIdsRequest) SetUserId(v string) *DeleteMatrixDataByRowIdsRequest {
	s.UserId = &v
	return s
}

type DeleteMatrixDataByRowIdsResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteMatrixDataByRowIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMatrixDataByRowIdsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMatrixDataByRowIdsResponseBody) SetSuccess(v bool) *DeleteMatrixDataByRowIdsResponseBody {
	s.Success = &v
	return s
}

type DeleteMatrixDataByRowIdsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMatrixDataByRowIdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMatrixDataByRowIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMatrixDataByRowIdsResponse) GoString() string {
	return s.String()
}

func (s *DeleteMatrixDataByRowIdsResponse) SetHeaders(v map[string]*string) *DeleteMatrixDataByRowIdsResponse {
	s.Headers = v
	return s
}

func (s *DeleteMatrixDataByRowIdsResponse) SetStatusCode(v int32) *DeleteMatrixDataByRowIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMatrixDataByRowIdsResponse) SetBody(v *DeleteMatrixDataByRowIdsResponseBody) *DeleteMatrixDataByRowIdsResponse {
	s.Body = v
	return s
}

type GetAgentTasksHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAgentTasksHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAgentTasksHeaders) GoString() string {
	return s.String()
}

func (s *GetAgentTasksHeaders) SetCommonHeaders(v map[string]*string) *GetAgentTasksHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAgentTasksHeaders) SetXAcsDingtalkAccessToken(v string) *GetAgentTasksHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAgentTasksRequest struct {
	// example:
	//
	// Agent--XXXXX
	AgentUuid *string `json:"agentUuid,omitempty" xml:"agentUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dingxxxx
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// 10001
	Keywords *string `json:"keywords,omitempty" xml:"keywords,omitempty"`
	// example:
	//
	// 100
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// ALL
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 501453
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetAgentTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgentTasksRequest) GoString() string {
	return s.String()
}

func (s *GetAgentTasksRequest) SetAgentUuid(v string) *GetAgentTasksRequest {
	s.AgentUuid = &v
	return s
}

func (s *GetAgentTasksRequest) SetCorpId(v string) *GetAgentTasksRequest {
	s.CorpId = &v
	return s
}

func (s *GetAgentTasksRequest) SetKeywords(v string) *GetAgentTasksRequest {
	s.Keywords = &v
	return s
}

func (s *GetAgentTasksRequest) SetPageNumber(v int32) *GetAgentTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *GetAgentTasksRequest) SetPageSize(v int32) *GetAgentTasksRequest {
	s.PageSize = &v
	return s
}

func (s *GetAgentTasksRequest) SetStatus(v string) *GetAgentTasksRequest {
	s.Status = &v
	return s
}

func (s *GetAgentTasksRequest) SetToken(v string) *GetAgentTasksRequest {
	s.Token = &v
	return s
}

func (s *GetAgentTasksRequest) SetUserId(v string) *GetAgentTasksRequest {
	s.UserId = &v
	return s
}

type GetAgentTasksResponseBody struct {
	CurrentPage *int32                             `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	Limit       *int32                             `json:"limit,omitempty" xml:"limit,omitempty"`
	Start       *int32                             `json:"start,omitempty" xml:"start,omitempty"`
	TotalCount  *int32                             `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	Values      []*GetAgentTasksResponseBodyValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetAgentTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAgentTasksResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentTasksResponseBody) SetCurrentPage(v int32) *GetAgentTasksResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *GetAgentTasksResponseBody) SetLimit(v int32) *GetAgentTasksResponseBody {
	s.Limit = &v
	return s
}

func (s *GetAgentTasksResponseBody) SetStart(v int32) *GetAgentTasksResponseBody {
	s.Start = &v
	return s
}

func (s *GetAgentTasksResponseBody) SetTotalCount(v int32) *GetAgentTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetAgentTasksResponseBody) SetValues(v []*GetAgentTasksResponseBodyValues) *GetAgentTasksResponseBody {
	s.Values = v
	return s
}

type GetAgentTasksResponseBodyValues struct {
	AgentCategory       *string `json:"agentCategory,omitempty" xml:"agentCategory,omitempty"`
	AgentCreateGMT      *string `json:"agentCreateGMT,omitempty" xml:"agentCreateGMT,omitempty"`
	AgentEndGMT         *string `json:"agentEndGMT,omitempty" xml:"agentEndGMT,omitempty"`
	AgentName           *string `json:"agentName,omitempty" xml:"agentName,omitempty"`
	AgentRangeType      *string `json:"agentRangeType,omitempty" xml:"agentRangeType,omitempty"`
	AgentRangeValue     *string `json:"agentRangeValue,omitempty" xml:"agentRangeValue,omitempty"`
	AgentStartGMT       *string `json:"agentStartGMT,omitempty" xml:"agentStartGMT,omitempty"`
	AgentType           *string `json:"agentType,omitempty" xml:"agentType,omitempty"`
	AgentUserId         *string `json:"agentUserId,omitempty" xml:"agentUserId,omitempty"`
	AgentUuid           *string `json:"agentUuid,omitempty" xml:"agentUuid,omitempty"`
	Creator             *string `json:"creator,omitempty" xml:"creator,omitempty"`
	CreatorName         *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	Modifier            *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	NeedNoticePrincipal *string `json:"needNoticePrincipal,omitempty" xml:"needNoticePrincipal,omitempty"`
	PrincipalName       *string `json:"principalName,omitempty" xml:"principalName,omitempty"`
	PrincipalUserId     *string `json:"principalUserId,omitempty" xml:"principalUserId,omitempty"`
	Status              *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetAgentTasksResponseBodyValues) String() string {
	return tea.Prettify(s)
}

func (s GetAgentTasksResponseBodyValues) GoString() string {
	return s.String()
}

func (s *GetAgentTasksResponseBodyValues) SetAgentCategory(v string) *GetAgentTasksResponseBodyValues {
	s.AgentCategory = &v
	return s
}

func (s *GetAgentTasksResponseBodyValues) SetAgentCreateGMT(v string) *GetAgentTasksResponseBodyValues {
	s.AgentCreateGMT = &v
	return s
}

func (s *GetAgentTasksResponseBodyValues) SetAgentEndGMT(v string) *GetAgentTasksResponseBodyValues {
	s.AgentEndGMT = &v
	return s
}

func (s *GetAgentTasksResponseBodyValues) SetAgentName(v string) *GetAgentTasksResponseBodyValues {
	s.AgentName = &v
	return s
}

func (s *GetAgentTasksResponseBodyValues) SetAgentRangeType(v string) *GetAgentTasksResponseBodyValues {
	s.AgentRangeType = &v
	return s
}

func (s *GetAgentTasksResponseBodyValues) SetAgentRangeValue(v string) *GetAgentTasksResponseBodyValues {
	s.AgentRangeValue = &v
	return s
}

func (s *GetAgentTasksResponseBodyValues) SetAgentStartGMT(v string) *GetAgentTasksResponseBodyValues {
	s.AgentStartGMT = &v
	return s
}

func (s *GetAgentTasksResponseBodyValues) SetAgentType(v string) *GetAgentTasksResponseBodyValues {
	s.AgentType = &v
	return s
}

func (s *GetAgentTasksResponseBodyValues) SetAgentUserId(v string) *GetAgentTasksResponseBodyValues {
	s.AgentUserId = &v
	return s
}

func (s *GetAgentTasksResponseBodyValues) SetAgentUuid(v string) *GetAgentTasksResponseBodyValues {
	s.AgentUuid = &v
	return s
}

func (s *GetAgentTasksResponseBodyValues) SetCreator(v string) *GetAgentTasksResponseBodyValues {
	s.Creator = &v
	return s
}

func (s *GetAgentTasksResponseBodyValues) SetCreatorName(v string) *GetAgentTasksResponseBodyValues {
	s.CreatorName = &v
	return s
}

func (s *GetAgentTasksResponseBodyValues) SetModifier(v string) *GetAgentTasksResponseBodyValues {
	s.Modifier = &v
	return s
}

func (s *GetAgentTasksResponseBodyValues) SetNeedNoticePrincipal(v string) *GetAgentTasksResponseBodyValues {
	s.NeedNoticePrincipal = &v
	return s
}

func (s *GetAgentTasksResponseBodyValues) SetPrincipalName(v string) *GetAgentTasksResponseBodyValues {
	s.PrincipalName = &v
	return s
}

func (s *GetAgentTasksResponseBodyValues) SetPrincipalUserId(v string) *GetAgentTasksResponseBodyValues {
	s.PrincipalUserId = &v
	return s
}

func (s *GetAgentTasksResponseBodyValues) SetStatus(v string) *GetAgentTasksResponseBodyValues {
	s.Status = &v
	return s
}

type GetAgentTasksResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAgentTasksResponse) GoString() string {
	return s.String()
}

func (s *GetAgentTasksResponse) SetHeaders(v map[string]*string) *GetAgentTasksResponse {
	s.Headers = v
	return s
}

func (s *GetAgentTasksResponse) SetStatusCode(v int32) *GetAgentTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentTasksResponse) SetBody(v *GetAgentTasksResponseBody) *GetAgentTasksResponse {
	s.Body = v
	return s
}

type GetFormComponentAliasListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFormComponentAliasListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFormComponentAliasListHeaders) GoString() string {
	return s.String()
}

func (s *GetFormComponentAliasListHeaders) SetCommonHeaders(v map[string]*string) *GetFormComponentAliasListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFormComponentAliasListHeaders) SetXAcsDingtalkAccessToken(v string) *GetFormComponentAliasListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFormComponentAliasListRequest struct {
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxxx
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 未知
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetFormComponentAliasListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFormComponentAliasListRequest) GoString() string {
	return s.String()
}

func (s *GetFormComponentAliasListRequest) SetLanguage(v string) *GetFormComponentAliasListRequest {
	s.Language = &v
	return s
}

func (s *GetFormComponentAliasListRequest) SetSystemToken(v string) *GetFormComponentAliasListRequest {
	s.SystemToken = &v
	return s
}

func (s *GetFormComponentAliasListRequest) SetUserId(v string) *GetFormComponentAliasListRequest {
	s.UserId = &v
	return s
}

func (s *GetFormComponentAliasListRequest) SetVersion(v int64) *GetFormComponentAliasListRequest {
	s.Version = &v
	return s
}

type GetFormComponentAliasListResponseBody struct {
	Result []*GetFormComponentAliasListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetFormComponentAliasListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFormComponentAliasListResponseBody) GoString() string {
	return s.String()
}

func (s *GetFormComponentAliasListResponseBody) SetResult(v []*GetFormComponentAliasListResponseBodyResult) *GetFormComponentAliasListResponseBody {
	s.Result = v
	return s
}

type GetFormComponentAliasListResponseBodyResult struct {
	Alias   *string `json:"alias,omitempty" xml:"alias,omitempty"`
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
}

func (s GetFormComponentAliasListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFormComponentAliasListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFormComponentAliasListResponseBodyResult) SetAlias(v string) *GetFormComponentAliasListResponseBodyResult {
	s.Alias = &v
	return s
}

func (s *GetFormComponentAliasListResponseBodyResult) SetFieldId(v string) *GetFormComponentAliasListResponseBodyResult {
	s.FieldId = &v
	return s
}

type GetFormComponentAliasListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFormComponentAliasListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFormComponentAliasListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFormComponentAliasListResponse) GoString() string {
	return s.String()
}

func (s *GetFormComponentAliasListResponse) SetHeaders(v map[string]*string) *GetFormComponentAliasListResponse {
	s.Headers = v
	return s
}

func (s *GetFormComponentAliasListResponse) SetStatusCode(v int32) *GetFormComponentAliasListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFormComponentAliasListResponse) SetBody(v *GetFormComponentAliasListResponseBody) *GetFormComponentAliasListResponse {
	s.Body = v
	return s
}

type GetFormDataByIDHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFormDataByIDHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDHeaders) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDHeaders) SetCommonHeaders(v map[string]*string) *GetFormDataByIDHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFormDataByIDHeaders) SetXAcsDingtalkAccessToken(v string) *GetFormDataByIDHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFormDataByIDRequest struct {
	// example:
	//
	// APP_PBKT0MFBEBTDO8T7SLVP
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// example:
	//
	// FORM-AA28579F69644FC19A47FE267457E664ZVR1
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// hexxx
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// example:
	//
	// false
	UseAlias *bool   `json:"useAlias,omitempty" xml:"useAlias,omitempty"`
	UserId   *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetFormDataByIDRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDRequest) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDRequest) SetAppType(v string) *GetFormDataByIDRequest {
	s.AppType = &v
	return s
}

func (s *GetFormDataByIDRequest) SetFormUuid(v string) *GetFormDataByIDRequest {
	s.FormUuid = &v
	return s
}

func (s *GetFormDataByIDRequest) SetLanguage(v string) *GetFormDataByIDRequest {
	s.Language = &v
	return s
}

func (s *GetFormDataByIDRequest) SetSystemToken(v string) *GetFormDataByIDRequest {
	s.SystemToken = &v
	return s
}

func (s *GetFormDataByIDRequest) SetUseAlias(v bool) *GetFormDataByIDRequest {
	s.UseAlias = &v
	return s
}

func (s *GetFormDataByIDRequest) SetUserId(v string) *GetFormDataByIDRequest {
	s.UserId = &v
	return s
}

type GetFormDataByIDResponseBody struct {
	// example:
	//
	// {       "numberField_jcr0069o": 1,       "multiSelectField_jcr0069s": [         "选项三",         "选项二"       ],       "textareaField_jcr0069n": "duohang",       "employeeField_jcr0069x": [         "xxxx"       ],       "departmentField_jcr0069z": "xxxx",       "cascadeDate_jcr0069u": [         "1514736000000",         "1517328000000"       ],       "cascadeSelectField_jcr006a0": [         "part",         "part_b"       ],       "tableField_jcr006a1": [         {           "departmentField_jcr006ad": "xxxx",           "cascadeDate_jcr006aa": [             "1514736000000",             "1517328000000"           ],           "selectField_jcr006a6": "选项三",           "citySelectField_jcr006ac": [             "天津",             "天津市",             "河东区"           ],           "radioField_jcr006a5": "选项二",           "employeeField_jcr006ab": [             "xxxxxx",             "yyyyyy"           ],           "dateField_jcr006a9": 1517328000000,           "textField_jcr006a2": "明细下单行",           "textareaField_jcr006a3": "明细下多行",           "cascadeSelectField_jcr006ae": [             "product",             "product_a"           ],           "numberField_jcr006a4": 2,           "checkboxField_jcr006a7": [             "选项一",             "选项三",             "选项二"           ],           "multiSelectField_jcr006a8": [             "选项一",             "选项三",             "选项二"           ]         }       ],       "selectField_jcr0069q": "选项一",       "citySelectField_jcr0069y": [         "北京",         "北京市",         "东城区"       ],       "checkboxField_jcr0069r": [         "选项三",         "选项二"       ],       "textField_jcr0069m": "danhang",       "radioField_jcr0069p": "选项一",       "dateField_jcr0069t": 1516636800000     }
	FormData map[string]interface{} `json:"formData,omitempty" xml:"formData,omitempty"`
	// example:
	//
	// 33f6d221-17f8-42b7-836a-682b95a046c2
	FormInstId *string `json:"formInstId,omitempty" xml:"formInstId,omitempty"`
	// example:
	//
	// 2018-01-24 11:22:01
	ModifiedTimeGMT *string                                `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	Originator      *GetFormDataByIDResponseBodyOriginator `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
}

func (s GetFormDataByIDResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDResponseBody) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDResponseBody) SetFormData(v map[string]interface{}) *GetFormDataByIDResponseBody {
	s.FormData = v
	return s
}

func (s *GetFormDataByIDResponseBody) SetFormInstId(v string) *GetFormDataByIDResponseBody {
	s.FormInstId = &v
	return s
}

func (s *GetFormDataByIDResponseBody) SetModifiedTimeGMT(v string) *GetFormDataByIDResponseBody {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *GetFormDataByIDResponseBody) SetOriginator(v *GetFormDataByIDResponseBodyOriginator) *GetFormDataByIDResponseBody {
	s.Originator = v
	return s
}

type GetFormDataByIDResponseBodyOriginator struct {
	// example:
	//
	// 开发部
	DepartmentName *string                                    `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	Email          *string                                    `json:"email,omitempty" xml:"email,omitempty"`
	Name           *GetFormDataByIDResponseBodyOriginatorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	UserId         *string                                    `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetFormDataByIDResponseBodyOriginator) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDResponseBodyOriginator) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDResponseBodyOriginator) SetDepartmentName(v string) *GetFormDataByIDResponseBodyOriginator {
	s.DepartmentName = &v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginator) SetEmail(v string) *GetFormDataByIDResponseBodyOriginator {
	s.Email = &v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginator) SetName(v *GetFormDataByIDResponseBodyOriginatorName) *GetFormDataByIDResponseBodyOriginator {
	s.Name = v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginator) SetUserId(v string) *GetFormDataByIDResponseBodyOriginator {
	s.UserId = &v
	return s
}

type GetFormDataByIDResponseBodyOriginatorName struct {
	// example:
	//
	// 张三
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// example:
	//
	// ZhangSan
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	// example:
	//
	// i18n
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetFormDataByIDResponseBodyOriginatorName) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDResponseBodyOriginatorName) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDResponseBodyOriginatorName) SetNameInChinese(v string) *GetFormDataByIDResponseBodyOriginatorName {
	s.NameInChinese = &v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginatorName) SetNameInEnglish(v string) *GetFormDataByIDResponseBodyOriginatorName {
	s.NameInEnglish = &v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginatorName) SetType(v string) *GetFormDataByIDResponseBodyOriginatorName {
	s.Type = &v
	return s
}

type GetFormDataByIDResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFormDataByIDResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFormDataByIDResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDResponse) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDResponse) SetHeaders(v map[string]*string) *GetFormDataByIDResponse {
	s.Headers = v
	return s
}

func (s *GetFormDataByIDResponse) SetStatusCode(v int32) *GetFormDataByIDResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFormDataByIDResponse) SetBody(v *GetFormDataByIDResponseBody) *GetFormDataByIDResponse {
	s.Body = v
	return s
}

type GetInstanceByIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetInstanceByIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceByIdHeaders) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdHeaders) SetCommonHeaders(v map[string]*string) *GetInstanceByIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInstanceByIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetInstanceByIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetInstanceByIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_PBKT0MFBEBTDO8T7SLVP
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// example:
	//
	// FORM-ADFC8E8E5ADE4B2F8FC2316CFC42A55CJLWZ
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxyy
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// example:
	//
	// false
	UseAlias *bool `json:"useAlias,omitempty" xml:"useAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 未知
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetInstanceByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceByIdRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdRequest) SetAppType(v string) *GetInstanceByIdRequest {
	s.AppType = &v
	return s
}

func (s *GetInstanceByIdRequest) SetFormUuid(v string) *GetInstanceByIdRequest {
	s.FormUuid = &v
	return s
}

func (s *GetInstanceByIdRequest) SetLanguage(v string) *GetInstanceByIdRequest {
	s.Language = &v
	return s
}

func (s *GetInstanceByIdRequest) SetSystemToken(v string) *GetInstanceByIdRequest {
	s.SystemToken = &v
	return s
}

func (s *GetInstanceByIdRequest) SetUseAlias(v bool) *GetInstanceByIdRequest {
	s.UseAlias = &v
	return s
}

func (s *GetInstanceByIdRequest) SetUserId(v string) *GetInstanceByIdRequest {
	s.UserId = &v
	return s
}

type GetInstanceByIdResponseBody struct {
	ActionExecutor    []*GetInstanceByIdResponseBodyActionExecutor `json:"actionExecutor,omitempty" xml:"actionExecutor,omitempty" type:"Repeated"`
	ApprovedResult    *string                                      `json:"approvedResult,omitempty" xml:"approvedResult,omitempty"`
	CreateTimeGMT     *string                                      `json:"createTimeGMT,omitempty" xml:"createTimeGMT,omitempty"`
	Data              map[string]interface{}                       `json:"data,omitempty" xml:"data,omitempty"`
	FormUuid          *string                                      `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	InstanceStatus    *string                                      `json:"instanceStatus,omitempty" xml:"instanceStatus,omitempty"`
	ModifiedTimeGMT   *string                                      `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	Originator        *GetInstanceByIdResponseBodyOriginator       `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	ProcessCode       *string                                      `json:"processCode,omitempty" xml:"processCode,omitempty"`
	ProcessInstanceId *string                                      `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	Title             *string                                      `json:"title,omitempty" xml:"title,omitempty"`
	Version           *int64                                       `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetInstanceByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdResponseBody) SetActionExecutor(v []*GetInstanceByIdResponseBodyActionExecutor) *GetInstanceByIdResponseBody {
	s.ActionExecutor = v
	return s
}

func (s *GetInstanceByIdResponseBody) SetApprovedResult(v string) *GetInstanceByIdResponseBody {
	s.ApprovedResult = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetCreateTimeGMT(v string) *GetInstanceByIdResponseBody {
	s.CreateTimeGMT = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetData(v map[string]interface{}) *GetInstanceByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceByIdResponseBody) SetFormUuid(v string) *GetInstanceByIdResponseBody {
	s.FormUuid = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetInstanceStatus(v string) *GetInstanceByIdResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetModifiedTimeGMT(v string) *GetInstanceByIdResponseBody {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetOriginator(v *GetInstanceByIdResponseBodyOriginator) *GetInstanceByIdResponseBody {
	s.Originator = v
	return s
}

func (s *GetInstanceByIdResponseBody) SetProcessCode(v string) *GetInstanceByIdResponseBody {
	s.ProcessCode = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetProcessInstanceId(v string) *GetInstanceByIdResponseBody {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetTitle(v string) *GetInstanceByIdResponseBody {
	s.Title = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetVersion(v int64) *GetInstanceByIdResponseBody {
	s.Version = &v
	return s
}

type GetInstanceByIdResponseBodyActionExecutor struct {
	DeptName *string                                        `json:"deptName,omitempty" xml:"deptName,omitempty"`
	Email    *string                                        `json:"email,omitempty" xml:"email,omitempty"`
	Name     *GetInstanceByIdResponseBodyActionExecutorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	UserId   *string                                        `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetInstanceByIdResponseBodyActionExecutor) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceByIdResponseBodyActionExecutor) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdResponseBodyActionExecutor) SetDeptName(v string) *GetInstanceByIdResponseBodyActionExecutor {
	s.DeptName = &v
	return s
}

func (s *GetInstanceByIdResponseBodyActionExecutor) SetEmail(v string) *GetInstanceByIdResponseBodyActionExecutor {
	s.Email = &v
	return s
}

func (s *GetInstanceByIdResponseBodyActionExecutor) SetName(v *GetInstanceByIdResponseBodyActionExecutorName) *GetInstanceByIdResponseBodyActionExecutor {
	s.Name = v
	return s
}

func (s *GetInstanceByIdResponseBodyActionExecutor) SetUserId(v string) *GetInstanceByIdResponseBodyActionExecutor {
	s.UserId = &v
	return s
}

type GetInstanceByIdResponseBodyActionExecutorName struct {
	// example:
	//
	// 张三
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// example:
	//
	// ZhangSan
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	Type          *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetInstanceByIdResponseBodyActionExecutorName) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceByIdResponseBodyActionExecutorName) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdResponseBodyActionExecutorName) SetNameInChinese(v string) *GetInstanceByIdResponseBodyActionExecutorName {
	s.NameInChinese = &v
	return s
}

func (s *GetInstanceByIdResponseBodyActionExecutorName) SetNameInEnglish(v string) *GetInstanceByIdResponseBodyActionExecutorName {
	s.NameInEnglish = &v
	return s
}

func (s *GetInstanceByIdResponseBodyActionExecutorName) SetType(v string) *GetInstanceByIdResponseBodyActionExecutorName {
	s.Type = &v
	return s
}

type GetInstanceByIdResponseBodyOriginator struct {
	DeptName *string                                    `json:"deptName,omitempty" xml:"deptName,omitempty"`
	Email    *string                                    `json:"email,omitempty" xml:"email,omitempty"`
	Name     *GetInstanceByIdResponseBodyOriginatorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	UserId   *string                                    `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetInstanceByIdResponseBodyOriginator) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceByIdResponseBodyOriginator) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdResponseBodyOriginator) SetDeptName(v string) *GetInstanceByIdResponseBodyOriginator {
	s.DeptName = &v
	return s
}

func (s *GetInstanceByIdResponseBodyOriginator) SetEmail(v string) *GetInstanceByIdResponseBodyOriginator {
	s.Email = &v
	return s
}

func (s *GetInstanceByIdResponseBodyOriginator) SetName(v *GetInstanceByIdResponseBodyOriginatorName) *GetInstanceByIdResponseBodyOriginator {
	s.Name = v
	return s
}

func (s *GetInstanceByIdResponseBodyOriginator) SetUserId(v string) *GetInstanceByIdResponseBodyOriginator {
	s.UserId = &v
	return s
}

type GetInstanceByIdResponseBodyOriginatorName struct {
	// example:
	//
	// 张三
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// example:
	//
	// ZhangSan
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	Type          *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetInstanceByIdResponseBodyOriginatorName) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceByIdResponseBodyOriginatorName) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdResponseBodyOriginatorName) SetNameInChinese(v string) *GetInstanceByIdResponseBodyOriginatorName {
	s.NameInChinese = &v
	return s
}

func (s *GetInstanceByIdResponseBodyOriginatorName) SetNameInEnglish(v string) *GetInstanceByIdResponseBodyOriginatorName {
	s.NameInEnglish = &v
	return s
}

func (s *GetInstanceByIdResponseBodyOriginatorName) SetType(v string) *GetInstanceByIdResponseBodyOriginatorName {
	s.Type = &v
	return s
}

type GetInstanceByIdResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceByIdResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdResponse) SetHeaders(v map[string]*string) *GetInstanceByIdResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceByIdResponse) SetStatusCode(v int32) *GetInstanceByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceByIdResponse) SetBody(v *GetInstanceByIdResponseBody) *GetInstanceByIdResponse {
	s.Body = v
	return s
}

type GetInstanceIdListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetInstanceIdListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceIdListHeaders) GoString() string {
	return s.String()
}

func (s *GetInstanceIdListHeaders) SetCommonHeaders(v map[string]*string) *GetInstanceIdListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInstanceIdListHeaders) SetXAcsDingtalkAccessToken(v string) *GetInstanceIdListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetInstanceIdListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_PBKT0MFBEBTDO8T7SLVP
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// example:
	//
	// agree
	ApprovedResult *string `json:"approvedResult,omitempty" xml:"approvedResult,omitempty"`
	// example:
	//
	// 2018-01-01
	CreateFromTimeGMT *string `json:"createFromTimeGMT,omitempty" xml:"createFromTimeGMT,omitempty"`
	// example:
	//
	// 2018-02-01
	CreateToTimeGMT *string `json:"createToTimeGMT,omitempty" xml:"createToTimeGMT,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-EF6Y4G8WO2FN0SUB43TDQ3CGC3FMFQ1G9400RCJ3
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// RUNNING
	InstanceStatus *string `json:"instanceStatus,omitempty" xml:"instanceStatus,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// 2018-01-01
	ModifiedFromTimeGMT *string `json:"modifiedFromTimeGMT,omitempty" xml:"modifiedFromTimeGMT,omitempty"`
	// example:
	//
	// 2018-02-01
	ModifiedToTimeGMT *string `json:"modifiedToTimeGMT,omitempty" xml:"modifiedToTimeGMT,omitempty"`
	// example:
	//
	// ding123
	OriginatorId *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	// example:
	//
	// {"text_field":"123"}
	SearchFieldJson *string `json:"searchFieldJson,omitempty" xml:"searchFieldJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxxx
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// example:
	//
	// 2199132092
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// false
	UseAlias *bool `json:"useAlias,omitempty" xml:"useAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ding123
	UserId     *string `json:"userId,omitempty" xml:"userId,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetInstanceIdListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceIdListRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceIdListRequest) SetAppType(v string) *GetInstanceIdListRequest {
	s.AppType = &v
	return s
}

func (s *GetInstanceIdListRequest) SetApprovedResult(v string) *GetInstanceIdListRequest {
	s.ApprovedResult = &v
	return s
}

func (s *GetInstanceIdListRequest) SetCreateFromTimeGMT(v string) *GetInstanceIdListRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *GetInstanceIdListRequest) SetCreateToTimeGMT(v string) *GetInstanceIdListRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *GetInstanceIdListRequest) SetFormUuid(v string) *GetInstanceIdListRequest {
	s.FormUuid = &v
	return s
}

func (s *GetInstanceIdListRequest) SetInstanceStatus(v string) *GetInstanceIdListRequest {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstanceIdListRequest) SetLanguage(v string) *GetInstanceIdListRequest {
	s.Language = &v
	return s
}

func (s *GetInstanceIdListRequest) SetModifiedFromTimeGMT(v string) *GetInstanceIdListRequest {
	s.ModifiedFromTimeGMT = &v
	return s
}

func (s *GetInstanceIdListRequest) SetModifiedToTimeGMT(v string) *GetInstanceIdListRequest {
	s.ModifiedToTimeGMT = &v
	return s
}

func (s *GetInstanceIdListRequest) SetOriginatorId(v string) *GetInstanceIdListRequest {
	s.OriginatorId = &v
	return s
}

func (s *GetInstanceIdListRequest) SetSearchFieldJson(v string) *GetInstanceIdListRequest {
	s.SearchFieldJson = &v
	return s
}

func (s *GetInstanceIdListRequest) SetSystemToken(v string) *GetInstanceIdListRequest {
	s.SystemToken = &v
	return s
}

func (s *GetInstanceIdListRequest) SetTaskId(v string) *GetInstanceIdListRequest {
	s.TaskId = &v
	return s
}

func (s *GetInstanceIdListRequest) SetUseAlias(v bool) *GetInstanceIdListRequest {
	s.UseAlias = &v
	return s
}

func (s *GetInstanceIdListRequest) SetUserId(v string) *GetInstanceIdListRequest {
	s.UserId = &v
	return s
}

func (s *GetInstanceIdListRequest) SetPageNumber(v int32) *GetInstanceIdListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetInstanceIdListRequest) SetPageSize(v int32) *GetInstanceIdListRequest {
	s.PageSize = &v
	return s
}

type GetInstanceIdListResponseBody struct {
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetInstanceIdListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceIdListResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceIdListResponseBody) SetData(v []*string) *GetInstanceIdListResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceIdListResponseBody) SetPageNumber(v int64) *GetInstanceIdListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetInstanceIdListResponseBody) SetTotalCount(v int64) *GetInstanceIdListResponseBody {
	s.TotalCount = &v
	return s
}

type GetInstanceIdListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceIdListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceIdListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceIdListResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceIdListResponse) SetHeaders(v map[string]*string) *GetInstanceIdListResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceIdListResponse) SetStatusCode(v int32) *GetInstanceIdListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceIdListResponse) SetBody(v *GetInstanceIdListResponseBody) *GetInstanceIdListResponse {
	s.Body = v
	return s
}

type GetInstancesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetInstancesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesHeaders) GoString() string {
	return s.String()
}

func (s *GetInstancesHeaders) SetCommonHeaders(v map[string]*string) *GetInstancesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInstancesHeaders) SetXAcsDingtalkAccessToken(v string) *GetInstancesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetInstancesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_PBKT0MFBEBTDO8T7SLVP
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// example:
	//
	// agree
	ApprovedResult *string `json:"approvedResult,omitempty" xml:"approvedResult,omitempty"`
	// example:
	//
	// 2018-01-01
	CreateFromTimeGMT *string `json:"createFromTimeGMT,omitempty" xml:"createFromTimeGMT,omitempty"`
	// example:
	//
	// 2018-02-01
	CreateToTimeGMT *string `json:"createToTimeGMT,omitempty" xml:"createToTimeGMT,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-EF6Y4G8WO2FN0SUB43TDQ3CGC3FMFQ1G9400RCJ3
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// RUNNING
	InstanceStatus *string `json:"instanceStatus,omitempty" xml:"instanceStatus,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// 2018-01-01
	ModifiedFromTimeGMT *string `json:"modifiedFromTimeGMT,omitempty" xml:"modifiedFromTimeGMT,omitempty"`
	// example:
	//
	// 2018-02-01
	ModifiedToTimeGMT *string `json:"modifiedToTimeGMT,omitempty" xml:"modifiedToTimeGMT,omitempty"`
	// example:
	//
	// 例如按照创建时间升序再按照文本组件值升序排序则填{\"gmt_create\":\"+\",\"textField_1234\":\"+\"} ，详情参考 https://www.yuque.com/yida/support/agb8im#CQro8
	OrderConfigJson *string `json:"orderConfigJson,omitempty" xml:"orderConfigJson,omitempty"`
	// example:
	//
	// manager123
	OriginatorId *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	// example:
	//
	// 模式1：根据组件值模糊匹配，示例：{"textField_jcr0069m":"danhang","selectField_jcr0069q":"K"}     模式2: 采用数据管理的查询过滤条件，匹配功能更强大，示例：[{"key":"currentNodeName","value":"步凡","type":"TEXT","operator":"like","componentName":"TextField”}]，详情参考  https://www.yuque.com/yida/support/agb8im#F4S8e
	SearchFieldJson *string `json:"searchFieldJson,omitempty" xml:"searchFieldJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxx
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// example:
	//
	// 2199132092
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// false
	UseAlias *bool `json:"useAlias,omitempty" xml:"useAlias,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesRequest) GoString() string {
	return s.String()
}

func (s *GetInstancesRequest) SetAppType(v string) *GetInstancesRequest {
	s.AppType = &v
	return s
}

func (s *GetInstancesRequest) SetApprovedResult(v string) *GetInstancesRequest {
	s.ApprovedResult = &v
	return s
}

func (s *GetInstancesRequest) SetCreateFromTimeGMT(v string) *GetInstancesRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *GetInstancesRequest) SetCreateToTimeGMT(v string) *GetInstancesRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *GetInstancesRequest) SetFormUuid(v string) *GetInstancesRequest {
	s.FormUuid = &v
	return s
}

func (s *GetInstancesRequest) SetInstanceStatus(v string) *GetInstancesRequest {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstancesRequest) SetLanguage(v string) *GetInstancesRequest {
	s.Language = &v
	return s
}

func (s *GetInstancesRequest) SetModifiedFromTimeGMT(v string) *GetInstancesRequest {
	s.ModifiedFromTimeGMT = &v
	return s
}

func (s *GetInstancesRequest) SetModifiedToTimeGMT(v string) *GetInstancesRequest {
	s.ModifiedToTimeGMT = &v
	return s
}

func (s *GetInstancesRequest) SetOrderConfigJson(v string) *GetInstancesRequest {
	s.OrderConfigJson = &v
	return s
}

func (s *GetInstancesRequest) SetOriginatorId(v string) *GetInstancesRequest {
	s.OriginatorId = &v
	return s
}

func (s *GetInstancesRequest) SetSearchFieldJson(v string) *GetInstancesRequest {
	s.SearchFieldJson = &v
	return s
}

func (s *GetInstancesRequest) SetSystemToken(v string) *GetInstancesRequest {
	s.SystemToken = &v
	return s
}

func (s *GetInstancesRequest) SetTaskId(v string) *GetInstancesRequest {
	s.TaskId = &v
	return s
}

func (s *GetInstancesRequest) SetUseAlias(v bool) *GetInstancesRequest {
	s.UseAlias = &v
	return s
}

func (s *GetInstancesRequest) SetUserId(v string) *GetInstancesRequest {
	s.UserId = &v
	return s
}

func (s *GetInstancesRequest) SetPageNumber(v int32) *GetInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *GetInstancesRequest) SetPageSize(v int32) *GetInstancesRequest {
	s.PageSize = &v
	return s
}

type GetInstancesResponseBody struct {
	Data []*GetInstancesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstancesResponseBody) SetData(v []*GetInstancesResponseBodyData) *GetInstancesResponseBody {
	s.Data = v
	return s
}

func (s *GetInstancesResponseBody) SetPageNumber(v int64) *GetInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetInstancesResponseBody) SetTotalCount(v int64) *GetInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type GetInstancesResponseBodyData struct {
	ActionExecutor    []*GetInstancesResponseBodyDataActionExecutor `json:"actionExecutor,omitempty" xml:"actionExecutor,omitempty" type:"Repeated"`
	ApprovedResult    *string                                       `json:"approvedResult,omitempty" xml:"approvedResult,omitempty"`
	CreateTimeGMT     *string                                       `json:"createTimeGMT,omitempty" xml:"createTimeGMT,omitempty"`
	Data              map[string]interface{}                        `json:"data,omitempty" xml:"data,omitempty"`
	FormUuid          *string                                       `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	InstanceStatus    *string                                       `json:"instanceStatus,omitempty" xml:"instanceStatus,omitempty"`
	ModifiedTimeGMT   *string                                       `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	Originator        *GetInstancesResponseBodyDataOriginator       `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	ProcessCode       *string                                       `json:"processCode,omitempty" xml:"processCode,omitempty"`
	ProcessInstanceId *string                                       `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	Title             *string                                       `json:"title,omitempty" xml:"title,omitempty"`
	Version           *int64                                        `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstancesResponseBodyData) SetActionExecutor(v []*GetInstancesResponseBodyDataActionExecutor) *GetInstancesResponseBodyData {
	s.ActionExecutor = v
	return s
}

func (s *GetInstancesResponseBodyData) SetApprovedResult(v string) *GetInstancesResponseBodyData {
	s.ApprovedResult = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetCreateTimeGMT(v string) *GetInstancesResponseBodyData {
	s.CreateTimeGMT = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetData(v map[string]interface{}) *GetInstancesResponseBodyData {
	s.Data = v
	return s
}

func (s *GetInstancesResponseBodyData) SetFormUuid(v string) *GetInstancesResponseBodyData {
	s.FormUuid = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetInstanceStatus(v string) *GetInstancesResponseBodyData {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetModifiedTimeGMT(v string) *GetInstancesResponseBodyData {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetOriginator(v *GetInstancesResponseBodyDataOriginator) *GetInstancesResponseBodyData {
	s.Originator = v
	return s
}

func (s *GetInstancesResponseBodyData) SetProcessCode(v string) *GetInstancesResponseBodyData {
	s.ProcessCode = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetProcessInstanceId(v string) *GetInstancesResponseBodyData {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetTitle(v string) *GetInstancesResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetVersion(v int64) *GetInstancesResponseBodyData {
	s.Version = &v
	return s
}

type GetInstancesResponseBodyDataActionExecutor struct {
	DeptName *string                                         `json:"deptName,omitempty" xml:"deptName,omitempty"`
	Email    *string                                         `json:"email,omitempty" xml:"email,omitempty"`
	Name     *GetInstancesResponseBodyDataActionExecutorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	UserId   *string                                         `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetInstancesResponseBodyDataActionExecutor) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesResponseBodyDataActionExecutor) GoString() string {
	return s.String()
}

func (s *GetInstancesResponseBodyDataActionExecutor) SetDeptName(v string) *GetInstancesResponseBodyDataActionExecutor {
	s.DeptName = &v
	return s
}

func (s *GetInstancesResponseBodyDataActionExecutor) SetEmail(v string) *GetInstancesResponseBodyDataActionExecutor {
	s.Email = &v
	return s
}

func (s *GetInstancesResponseBodyDataActionExecutor) SetName(v *GetInstancesResponseBodyDataActionExecutorName) *GetInstancesResponseBodyDataActionExecutor {
	s.Name = v
	return s
}

func (s *GetInstancesResponseBodyDataActionExecutor) SetUserId(v string) *GetInstancesResponseBodyDataActionExecutor {
	s.UserId = &v
	return s
}

type GetInstancesResponseBodyDataActionExecutorName struct {
	// example:
	//
	// 张三
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// example:
	//
	// ZhangSan
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	Type          *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetInstancesResponseBodyDataActionExecutorName) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesResponseBodyDataActionExecutorName) GoString() string {
	return s.String()
}

func (s *GetInstancesResponseBodyDataActionExecutorName) SetNameInChinese(v string) *GetInstancesResponseBodyDataActionExecutorName {
	s.NameInChinese = &v
	return s
}

func (s *GetInstancesResponseBodyDataActionExecutorName) SetNameInEnglish(v string) *GetInstancesResponseBodyDataActionExecutorName {
	s.NameInEnglish = &v
	return s
}

func (s *GetInstancesResponseBodyDataActionExecutorName) SetType(v string) *GetInstancesResponseBodyDataActionExecutorName {
	s.Type = &v
	return s
}

type GetInstancesResponseBodyDataOriginator struct {
	DeptName *string                                     `json:"deptName,omitempty" xml:"deptName,omitempty"`
	Email    *string                                     `json:"email,omitempty" xml:"email,omitempty"`
	Name     *GetInstancesResponseBodyDataOriginatorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	UserId   *string                                     `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetInstancesResponseBodyDataOriginator) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesResponseBodyDataOriginator) GoString() string {
	return s.String()
}

func (s *GetInstancesResponseBodyDataOriginator) SetDeptName(v string) *GetInstancesResponseBodyDataOriginator {
	s.DeptName = &v
	return s
}

func (s *GetInstancesResponseBodyDataOriginator) SetEmail(v string) *GetInstancesResponseBodyDataOriginator {
	s.Email = &v
	return s
}

func (s *GetInstancesResponseBodyDataOriginator) SetName(v *GetInstancesResponseBodyDataOriginatorName) *GetInstancesResponseBodyDataOriginator {
	s.Name = v
	return s
}

func (s *GetInstancesResponseBodyDataOriginator) SetUserId(v string) *GetInstancesResponseBodyDataOriginator {
	s.UserId = &v
	return s
}

type GetInstancesResponseBodyDataOriginatorName struct {
	// example:
	//
	// 张三
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// example:
	//
	// ZhangSan
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	Type          *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetInstancesResponseBodyDataOriginatorName) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesResponseBodyDataOriginatorName) GoString() string {
	return s.String()
}

func (s *GetInstancesResponseBodyDataOriginatorName) SetNameInChinese(v string) *GetInstancesResponseBodyDataOriginatorName {
	s.NameInChinese = &v
	return s
}

func (s *GetInstancesResponseBodyDataOriginatorName) SetNameInEnglish(v string) *GetInstancesResponseBodyDataOriginatorName {
	s.NameInEnglish = &v
	return s
}

func (s *GetInstancesResponseBodyDataOriginatorName) SetType(v string) *GetInstancesResponseBodyDataOriginatorName {
	s.Type = &v
	return s
}

type GetInstancesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesResponse) GoString() string {
	return s.String()
}

func (s *GetInstancesResponse) SetHeaders(v map[string]*string) *GetInstancesResponse {
	s.Headers = v
	return s
}

func (s *GetInstancesResponse) SetStatusCode(v int32) *GetInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstancesResponse) SetBody(v *GetInstancesResponseBody) *GetInstancesResponse {
	s.Body = v
	return s
}

type GetMatrixDetailByIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetMatrixDetailByIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetMatrixDetailByIdHeaders) GoString() string {
	return s.String()
}

func (s *GetMatrixDetailByIdHeaders) SetCommonHeaders(v map[string]*string) *GetMatrixDetailByIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMatrixDetailByIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetMatrixDetailByIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetMatrixDetailByIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dingxxxx
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MATRIX-C8I4J40EM81XLWZH61ZK
	MatrixId *string `json:"matrixId,omitempty" xml:"matrixId,omitempty"`
	// example:
	//
	// 100
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 501453
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetMatrixDetailByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMatrixDetailByIdRequest) GoString() string {
	return s.String()
}

func (s *GetMatrixDetailByIdRequest) SetCorpId(v string) *GetMatrixDetailByIdRequest {
	s.CorpId = &v
	return s
}

func (s *GetMatrixDetailByIdRequest) SetMatrixId(v string) *GetMatrixDetailByIdRequest {
	s.MatrixId = &v
	return s
}

func (s *GetMatrixDetailByIdRequest) SetPageNumber(v int32) *GetMatrixDetailByIdRequest {
	s.PageNumber = &v
	return s
}

func (s *GetMatrixDetailByIdRequest) SetPageSize(v int32) *GetMatrixDetailByIdRequest {
	s.PageSize = &v
	return s
}

func (s *GetMatrixDetailByIdRequest) SetToken(v string) *GetMatrixDetailByIdRequest {
	s.Token = &v
	return s
}

func (s *GetMatrixDetailByIdRequest) SetUserId(v string) *GetMatrixDetailByIdRequest {
	s.UserId = &v
	return s
}

type GetMatrixDetailByIdResponseBody struct {
	Result  *GetMatrixDetailByIdResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetMatrixDetailByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMatrixDetailByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetMatrixDetailByIdResponseBody) SetResult(v *GetMatrixDetailByIdResponseBodyResult) *GetMatrixDetailByIdResponseBody {
	s.Result = v
	return s
}

func (s *GetMatrixDetailByIdResponseBody) SetSuccess(v bool) *GetMatrixDetailByIdResponseBody {
	s.Success = &v
	return s
}

type GetMatrixDetailByIdResponseBodyResult struct {
	Description   *GetMatrixDetailByIdResponseBodyResultDescription `json:"description,omitempty" xml:"description,omitempty" type:"Struct"`
	MatrixData    *GetMatrixDetailByIdResponseBodyResultMatrixData  `json:"matrixData,omitempty" xml:"matrixData,omitempty" type:"Struct"`
	MatrixId      *string                                           `json:"matrixId,omitempty" xml:"matrixId,omitempty"`
	MatrixTable   *GetMatrixDetailByIdResponseBodyResultMatrixTable `json:"matrixTable,omitempty" xml:"matrixTable,omitempty" type:"Struct"`
	Name          *GetMatrixDetailByIdResponseBodyResultName        `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	RowTotalCount *int32                                            `json:"rowTotalCount,omitempty" xml:"rowTotalCount,omitempty"`
}

func (s GetMatrixDetailByIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetMatrixDetailByIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetMatrixDetailByIdResponseBodyResult) SetDescription(v *GetMatrixDetailByIdResponseBodyResultDescription) *GetMatrixDetailByIdResponseBodyResult {
	s.Description = v
	return s
}

func (s *GetMatrixDetailByIdResponseBodyResult) SetMatrixData(v *GetMatrixDetailByIdResponseBodyResultMatrixData) *GetMatrixDetailByIdResponseBodyResult {
	s.MatrixData = v
	return s
}

func (s *GetMatrixDetailByIdResponseBodyResult) SetMatrixId(v string) *GetMatrixDetailByIdResponseBodyResult {
	s.MatrixId = &v
	return s
}

func (s *GetMatrixDetailByIdResponseBodyResult) SetMatrixTable(v *GetMatrixDetailByIdResponseBodyResultMatrixTable) *GetMatrixDetailByIdResponseBodyResult {
	s.MatrixTable = v
	return s
}

func (s *GetMatrixDetailByIdResponseBodyResult) SetName(v *GetMatrixDetailByIdResponseBodyResultName) *GetMatrixDetailByIdResponseBodyResult {
	s.Name = v
	return s
}

func (s *GetMatrixDetailByIdResponseBodyResult) SetRowTotalCount(v int32) *GetMatrixDetailByIdResponseBodyResult {
	s.RowTotalCount = &v
	return s
}

type GetMatrixDetailByIdResponseBodyResultDescription struct {
	EnUS *string `json:"en_US,omitempty" xml:"en_US,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	ZhCN *string `json:"zh_CN,omitempty" xml:"zh_CN,omitempty"`
}

func (s GetMatrixDetailByIdResponseBodyResultDescription) String() string {
	return tea.Prettify(s)
}

func (s GetMatrixDetailByIdResponseBodyResultDescription) GoString() string {
	return s.String()
}

func (s *GetMatrixDetailByIdResponseBodyResultDescription) SetEnUS(v string) *GetMatrixDetailByIdResponseBodyResultDescription {
	s.EnUS = &v
	return s
}

func (s *GetMatrixDetailByIdResponseBodyResultDescription) SetType(v string) *GetMatrixDetailByIdResponseBodyResultDescription {
	s.Type = &v
	return s
}

func (s *GetMatrixDetailByIdResponseBodyResultDescription) SetZhCN(v string) *GetMatrixDetailByIdResponseBodyResultDescription {
	s.ZhCN = &v
	return s
}

type GetMatrixDetailByIdResponseBodyResultMatrixData struct {
	CurrentPage *int32      `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	Data        interface{} `json:"data,omitempty" xml:"data,omitempty"`
	TotalCount  *int32      `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetMatrixDetailByIdResponseBodyResultMatrixData) String() string {
	return tea.Prettify(s)
}

func (s GetMatrixDetailByIdResponseBodyResultMatrixData) GoString() string {
	return s.String()
}

func (s *GetMatrixDetailByIdResponseBodyResultMatrixData) SetCurrentPage(v int32) *GetMatrixDetailByIdResponseBodyResultMatrixData {
	s.CurrentPage = &v
	return s
}

func (s *GetMatrixDetailByIdResponseBodyResultMatrixData) SetData(v interface{}) *GetMatrixDetailByIdResponseBodyResultMatrixData {
	s.Data = v
	return s
}

func (s *GetMatrixDetailByIdResponseBodyResultMatrixData) SetTotalCount(v int32) *GetMatrixDetailByIdResponseBodyResultMatrixData {
	s.TotalCount = &v
	return s
}

type GetMatrixDetailByIdResponseBodyResultMatrixTable struct {
	ConditionColumn []*GetMatrixDetailByIdResponseBodyResultMatrixTableConditionColumn `json:"conditionColumn,omitempty" xml:"conditionColumn,omitempty" type:"Repeated"`
	ResultColumn    []*GetMatrixDetailByIdResponseBodyResultMatrixTableResultColumn    `json:"resultColumn,omitempty" xml:"resultColumn,omitempty" type:"Repeated"`
}

func (s GetMatrixDetailByIdResponseBodyResultMatrixTable) String() string {
	return tea.Prettify(s)
}

func (s GetMatrixDetailByIdResponseBodyResultMatrixTable) GoString() string {
	return s.String()
}

func (s *GetMatrixDetailByIdResponseBodyResultMatrixTable) SetConditionColumn(v []*GetMatrixDetailByIdResponseBodyResultMatrixTableConditionColumn) *GetMatrixDetailByIdResponseBodyResultMatrixTable {
	s.ConditionColumn = v
	return s
}

func (s *GetMatrixDetailByIdResponseBodyResultMatrixTable) SetResultColumn(v []*GetMatrixDetailByIdResponseBodyResultMatrixTableResultColumn) *GetMatrixDetailByIdResponseBodyResultMatrixTable {
	s.ResultColumn = v
	return s
}

type GetMatrixDetailByIdResponseBodyResultMatrixTableConditionColumn struct {
	ColumnId      *string `json:"columnId,omitempty" xml:"columnId,omitempty"`
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetMatrixDetailByIdResponseBodyResultMatrixTableConditionColumn) String() string {
	return tea.Prettify(s)
}

func (s GetMatrixDetailByIdResponseBodyResultMatrixTableConditionColumn) GoString() string {
	return s.String()
}

func (s *GetMatrixDetailByIdResponseBodyResultMatrixTableConditionColumn) SetColumnId(v string) *GetMatrixDetailByIdResponseBodyResultMatrixTableConditionColumn {
	s.ColumnId = &v
	return s
}

func (s *GetMatrixDetailByIdResponseBodyResultMatrixTableConditionColumn) SetComponentType(v string) *GetMatrixDetailByIdResponseBodyResultMatrixTableConditionColumn {
	s.ComponentType = &v
	return s
}

func (s *GetMatrixDetailByIdResponseBodyResultMatrixTableConditionColumn) SetName(v string) *GetMatrixDetailByIdResponseBodyResultMatrixTableConditionColumn {
	s.Name = &v
	return s
}

type GetMatrixDetailByIdResponseBodyResultMatrixTableResultColumn struct {
	ColumnId      *string `json:"columnId,omitempty" xml:"columnId,omitempty"`
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetMatrixDetailByIdResponseBodyResultMatrixTableResultColumn) String() string {
	return tea.Prettify(s)
}

func (s GetMatrixDetailByIdResponseBodyResultMatrixTableResultColumn) GoString() string {
	return s.String()
}

func (s *GetMatrixDetailByIdResponseBodyResultMatrixTableResultColumn) SetColumnId(v string) *GetMatrixDetailByIdResponseBodyResultMatrixTableResultColumn {
	s.ColumnId = &v
	return s
}

func (s *GetMatrixDetailByIdResponseBodyResultMatrixTableResultColumn) SetComponentType(v string) *GetMatrixDetailByIdResponseBodyResultMatrixTableResultColumn {
	s.ComponentType = &v
	return s
}

func (s *GetMatrixDetailByIdResponseBodyResultMatrixTableResultColumn) SetName(v string) *GetMatrixDetailByIdResponseBodyResultMatrixTableResultColumn {
	s.Name = &v
	return s
}

type GetMatrixDetailByIdResponseBodyResultName struct {
	EnUS *string `json:"en_US,omitempty" xml:"en_US,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	ZhCN *string `json:"zh_CN,omitempty" xml:"zh_CN,omitempty"`
}

func (s GetMatrixDetailByIdResponseBodyResultName) String() string {
	return tea.Prettify(s)
}

func (s GetMatrixDetailByIdResponseBodyResultName) GoString() string {
	return s.String()
}

func (s *GetMatrixDetailByIdResponseBodyResultName) SetEnUS(v string) *GetMatrixDetailByIdResponseBodyResultName {
	s.EnUS = &v
	return s
}

func (s *GetMatrixDetailByIdResponseBodyResultName) SetType(v string) *GetMatrixDetailByIdResponseBodyResultName {
	s.Type = &v
	return s
}

func (s *GetMatrixDetailByIdResponseBodyResultName) SetZhCN(v string) *GetMatrixDetailByIdResponseBodyResultName {
	s.ZhCN = &v
	return s
}

type GetMatrixDetailByIdResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMatrixDetailByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMatrixDetailByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMatrixDetailByIdResponse) GoString() string {
	return s.String()
}

func (s *GetMatrixDetailByIdResponse) SetHeaders(v map[string]*string) *GetMatrixDetailByIdResponse {
	s.Headers = v
	return s
}

func (s *GetMatrixDetailByIdResponse) SetStatusCode(v int32) *GetMatrixDetailByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMatrixDetailByIdResponse) SetBody(v *GetMatrixDetailByIdResponseBody) *GetMatrixDetailByIdResponse {
	s.Body = v
	return s
}

type GetRoleDetailByIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetRoleDetailByIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetRoleDetailByIdHeaders) GoString() string {
	return s.String()
}

func (s *GetRoleDetailByIdHeaders) SetCommonHeaders(v map[string]*string) *GetRoleDetailByIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRoleDetailByIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetRoleDetailByIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetRoleDetailByIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dingxxxx
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// 100
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ROLE-71dc7859-17df-490d-a93d-eb5506e31f42
	RoleUuid *string `json:"roleUuid,omitempty" xml:"roleUuid,omitempty"`
	// This parameter is required.
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 501453
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetRoleDetailByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRoleDetailByIdRequest) GoString() string {
	return s.String()
}

func (s *GetRoleDetailByIdRequest) SetCorpId(v string) *GetRoleDetailByIdRequest {
	s.CorpId = &v
	return s
}

func (s *GetRoleDetailByIdRequest) SetPageNumber(v int32) *GetRoleDetailByIdRequest {
	s.PageNumber = &v
	return s
}

func (s *GetRoleDetailByIdRequest) SetPageSize(v int32) *GetRoleDetailByIdRequest {
	s.PageSize = &v
	return s
}

func (s *GetRoleDetailByIdRequest) SetRoleUuid(v string) *GetRoleDetailByIdRequest {
	s.RoleUuid = &v
	return s
}

func (s *GetRoleDetailByIdRequest) SetToken(v string) *GetRoleDetailByIdRequest {
	s.Token = &v
	return s
}

func (s *GetRoleDetailByIdRequest) SetUserId(v string) *GetRoleDetailByIdRequest {
	s.UserId = &v
	return s
}

type GetRoleDetailByIdResponseBody struct {
	Result  *GetRoleDetailByIdResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetRoleDetailByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRoleDetailByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoleDetailByIdResponseBody) SetResult(v *GetRoleDetailByIdResponseBodyResult) *GetRoleDetailByIdResponseBody {
	s.Result = v
	return s
}

func (s *GetRoleDetailByIdResponseBody) SetSuccess(v bool) *GetRoleDetailByIdResponseBody {
	s.Success = &v
	return s
}

type GetRoleDetailByIdResponseBodyResult struct {
	CanModifyOwners  interface{}                                 `json:"canModifyOwners,omitempty" xml:"canModifyOwners,omitempty"`
	Description      *string                                     `json:"description,omitempty" xml:"description,omitempty"`
	MemberTotalCount *int32                                      `json:"memberTotalCount,omitempty" xml:"memberTotalCount,omitempty"`
	Members          *GetRoleDetailByIdResponseBodyResultMembers `json:"members,omitempty" xml:"members,omitempty" type:"Struct"`
	Name             *string                                     `json:"name,omitempty" xml:"name,omitempty"`
	ParentUuid       *string                                     `json:"parentUuid,omitempty" xml:"parentUuid,omitempty"`
	RoleUuid         *string                                     `json:"roleUuid,omitempty" xml:"roleUuid,omitempty"`
}

func (s GetRoleDetailByIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetRoleDetailByIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRoleDetailByIdResponseBodyResult) SetCanModifyOwners(v interface{}) *GetRoleDetailByIdResponseBodyResult {
	s.CanModifyOwners = v
	return s
}

func (s *GetRoleDetailByIdResponseBodyResult) SetDescription(v string) *GetRoleDetailByIdResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetRoleDetailByIdResponseBodyResult) SetMemberTotalCount(v int32) *GetRoleDetailByIdResponseBodyResult {
	s.MemberTotalCount = &v
	return s
}

func (s *GetRoleDetailByIdResponseBodyResult) SetMembers(v *GetRoleDetailByIdResponseBodyResultMembers) *GetRoleDetailByIdResponseBodyResult {
	s.Members = v
	return s
}

func (s *GetRoleDetailByIdResponseBodyResult) SetName(v string) *GetRoleDetailByIdResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetRoleDetailByIdResponseBodyResult) SetParentUuid(v string) *GetRoleDetailByIdResponseBodyResult {
	s.ParentUuid = &v
	return s
}

func (s *GetRoleDetailByIdResponseBodyResult) SetRoleUuid(v string) *GetRoleDetailByIdResponseBodyResult {
	s.RoleUuid = &v
	return s
}

type GetRoleDetailByIdResponseBodyResultMembers struct {
	CurrentPage *int32      `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	Data        interface{} `json:"data,omitempty" xml:"data,omitempty"`
	TotalCount  *int32      `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetRoleDetailByIdResponseBodyResultMembers) String() string {
	return tea.Prettify(s)
}

func (s GetRoleDetailByIdResponseBodyResultMembers) GoString() string {
	return s.String()
}

func (s *GetRoleDetailByIdResponseBodyResultMembers) SetCurrentPage(v int32) *GetRoleDetailByIdResponseBodyResultMembers {
	s.CurrentPage = &v
	return s
}

func (s *GetRoleDetailByIdResponseBodyResultMembers) SetData(v interface{}) *GetRoleDetailByIdResponseBodyResultMembers {
	s.Data = v
	return s
}

func (s *GetRoleDetailByIdResponseBodyResultMembers) SetTotalCount(v int32) *GetRoleDetailByIdResponseBodyResultMembers {
	s.TotalCount = &v
	return s
}

type GetRoleDetailByIdResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRoleDetailByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRoleDetailByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRoleDetailByIdResponse) GoString() string {
	return s.String()
}

func (s *GetRoleDetailByIdResponse) SetHeaders(v map[string]*string) *GetRoleDetailByIdResponse {
	s.Headers = v
	return s
}

func (s *GetRoleDetailByIdResponse) SetStatusCode(v int32) *GetRoleDetailByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRoleDetailByIdResponse) SetBody(v *GetRoleDetailByIdResponseBody) *GetRoleDetailByIdResponse {
	s.Body = v
	return s
}

type SaveAndUpdateMatrixDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SaveAndUpdateMatrixDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveAndUpdateMatrixDataHeaders) GoString() string {
	return s.String()
}

func (s *SaveAndUpdateMatrixDataHeaders) SetCommonHeaders(v map[string]*string) *SaveAndUpdateMatrixDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveAndUpdateMatrixDataHeaders) SetXAcsDingtalkAccessToken(v string) *SaveAndUpdateMatrixDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SaveAndUpdateMatrixDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dingxxxx
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{ 	"column_xx": "deptId", 	"column_yy": ["userId"], 	"column_zz": "项目一", 	"rowId": "row_1748398062718" }, { 	"column_xx": "deptId", 	"column_yy": ["userId", "userId"], 	"column_zz": "项目二" }]
	DataJson *string `json:"dataJson,omitempty" xml:"dataJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MATRIX-C8I4J40EM81XLWZH61ZK
	MatrixId *string `json:"matrixId,omitempty" xml:"matrixId,omitempty"`
	// This parameter is required.
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 501453
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SaveAndUpdateMatrixDataRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveAndUpdateMatrixDataRequest) GoString() string {
	return s.String()
}

func (s *SaveAndUpdateMatrixDataRequest) SetCorpId(v string) *SaveAndUpdateMatrixDataRequest {
	s.CorpId = &v
	return s
}

func (s *SaveAndUpdateMatrixDataRequest) SetDataJson(v string) *SaveAndUpdateMatrixDataRequest {
	s.DataJson = &v
	return s
}

func (s *SaveAndUpdateMatrixDataRequest) SetMatrixId(v string) *SaveAndUpdateMatrixDataRequest {
	s.MatrixId = &v
	return s
}

func (s *SaveAndUpdateMatrixDataRequest) SetToken(v string) *SaveAndUpdateMatrixDataRequest {
	s.Token = &v
	return s
}

func (s *SaveAndUpdateMatrixDataRequest) SetUserId(v string) *SaveAndUpdateMatrixDataRequest {
	s.UserId = &v
	return s
}

type SaveAndUpdateMatrixDataResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SaveAndUpdateMatrixDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveAndUpdateMatrixDataResponseBody) GoString() string {
	return s.String()
}

func (s *SaveAndUpdateMatrixDataResponseBody) SetSuccess(v bool) *SaveAndUpdateMatrixDataResponseBody {
	s.Success = &v
	return s
}

type SaveAndUpdateMatrixDataResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveAndUpdateMatrixDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveAndUpdateMatrixDataResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveAndUpdateMatrixDataResponse) GoString() string {
	return s.String()
}

func (s *SaveAndUpdateMatrixDataResponse) SetHeaders(v map[string]*string) *SaveAndUpdateMatrixDataResponse {
	s.Headers = v
	return s
}

func (s *SaveAndUpdateMatrixDataResponse) SetStatusCode(v int32) *SaveAndUpdateMatrixDataResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveAndUpdateMatrixDataResponse) SetBody(v *SaveAndUpdateMatrixDataResponseBody) *SaveAndUpdateMatrixDataResponse {
	s.Body = v
	return s
}

type SaveFormDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SaveFormDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveFormDataHeaders) GoString() string {
	return s.String()
}

func (s *SaveFormDataHeaders) SetCommonHeaders(v map[string]*string) *SaveFormDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveFormDataHeaders) SetXAcsDingtalkAccessToken(v string) *SaveFormDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SaveFormDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_PBKT0MFBEBTDO8T7SLVP
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"textField_jcpm6agt": "单行","employeeField_jcos0sar": ["workno"]}
	FormDataJson *string `json:"formDataJson,omitempty" xml:"formDataJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-NJYJZELV8YZRDEI2N5IQ7L6VEDMR1VE9GMPCJB
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxx
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// example:
	//
	// false
	UseAlias *bool `json:"useAlias,omitempty" xml:"useAlias,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SaveFormDataRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveFormDataRequest) GoString() string {
	return s.String()
}

func (s *SaveFormDataRequest) SetAppType(v string) *SaveFormDataRequest {
	s.AppType = &v
	return s
}

func (s *SaveFormDataRequest) SetFormDataJson(v string) *SaveFormDataRequest {
	s.FormDataJson = &v
	return s
}

func (s *SaveFormDataRequest) SetFormUuid(v string) *SaveFormDataRequest {
	s.FormUuid = &v
	return s
}

func (s *SaveFormDataRequest) SetLanguage(v string) *SaveFormDataRequest {
	s.Language = &v
	return s
}

func (s *SaveFormDataRequest) SetSystemToken(v string) *SaveFormDataRequest {
	s.SystemToken = &v
	return s
}

func (s *SaveFormDataRequest) SetUseAlias(v bool) *SaveFormDataRequest {
	s.UseAlias = &v
	return s
}

func (s *SaveFormDataRequest) SetUserId(v string) *SaveFormDataRequest {
	s.UserId = &v
	return s
}

type SaveFormDataResponseBody struct {
	// example:
	//
	// FINST-XIA66E71N7HSLK7H4KOZ388EEOP03A46YAYRK1
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SaveFormDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveFormDataResponseBody) GoString() string {
	return s.String()
}

func (s *SaveFormDataResponseBody) SetResult(v string) *SaveFormDataResponseBody {
	s.Result = &v
	return s
}

type SaveFormDataResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveFormDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveFormDataResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveFormDataResponse) GoString() string {
	return s.String()
}

func (s *SaveFormDataResponse) SetHeaders(v map[string]*string) *SaveFormDataResponse {
	s.Headers = v
	return s
}

func (s *SaveFormDataResponse) SetStatusCode(v int32) *SaveFormDataResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveFormDataResponse) SetBody(v *SaveFormDataResponseBody) *SaveFormDataResponse {
	s.Body = v
	return s
}

type SearchFormDataIdListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchFormDataIdListHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataIdListHeaders) GoString() string {
	return s.String()
}

func (s *SearchFormDataIdListHeaders) SetCommonHeaders(v map[string]*string) *SearchFormDataIdListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchFormDataIdListHeaders) SetXAcsDingtalkAccessToken(v string) *SearchFormDataIdListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchFormDataIdListRequest struct {
	// example:
	//
	// 2018-01-01
	CreateFromTimeGMT *string `json:"createFromTimeGMT,omitempty" xml:"createFromTimeGMT,omitempty"`
	// example:
	//
	// 2018-02-01
	CreateToTimeGMT *string `json:"createToTimeGMT,omitempty" xml:"createToTimeGMT,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// 2018-01-01
	ModifiedFromTimeGMT *string `json:"modifiedFromTimeGMT,omitempty" xml:"modifiedFromTimeGMT,omitempty"`
	// example:
	//
	// 2018-02-01
	ModifiedToTimeGMT *string `json:"modifiedToTimeGMT,omitempty" xml:"modifiedToTimeGMT,omitempty"`
	// example:
	//
	// dign1234
	OriginatorId    *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	SearchFieldJson *string `json:"searchFieldJson,omitempty" xml:"searchFieldJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxx
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// example:
	//
	// false
	UseAlias *bool `json:"useAlias,omitempty" xml:"useAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ding1234
	UserId     *string `json:"userId,omitempty" xml:"userId,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s SearchFormDataIdListRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataIdListRequest) GoString() string {
	return s.String()
}

func (s *SearchFormDataIdListRequest) SetCreateFromTimeGMT(v string) *SearchFormDataIdListRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetCreateToTimeGMT(v string) *SearchFormDataIdListRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetLanguage(v string) *SearchFormDataIdListRequest {
	s.Language = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetModifiedFromTimeGMT(v string) *SearchFormDataIdListRequest {
	s.ModifiedFromTimeGMT = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetModifiedToTimeGMT(v string) *SearchFormDataIdListRequest {
	s.ModifiedToTimeGMT = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetOriginatorId(v string) *SearchFormDataIdListRequest {
	s.OriginatorId = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetSearchFieldJson(v string) *SearchFormDataIdListRequest {
	s.SearchFieldJson = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetSystemToken(v string) *SearchFormDataIdListRequest {
	s.SystemToken = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetUseAlias(v bool) *SearchFormDataIdListRequest {
	s.UseAlias = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetUserId(v string) *SearchFormDataIdListRequest {
	s.UserId = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetPageNumber(v int32) *SearchFormDataIdListRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetPageSize(v int32) *SearchFormDataIdListRequest {
	s.PageSize = &v
	return s
}

type SearchFormDataIdListResponseBody struct {
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s SearchFormDataIdListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataIdListResponseBody) GoString() string {
	return s.String()
}

func (s *SearchFormDataIdListResponseBody) SetData(v []*string) *SearchFormDataIdListResponseBody {
	s.Data = v
	return s
}

func (s *SearchFormDataIdListResponseBody) SetPageNumber(v int64) *SearchFormDataIdListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchFormDataIdListResponseBody) SetTotalCount(v int64) *SearchFormDataIdListResponseBody {
	s.TotalCount = &v
	return s
}

type SearchFormDataIdListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchFormDataIdListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchFormDataIdListResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataIdListResponse) GoString() string {
	return s.String()
}

func (s *SearchFormDataIdListResponse) SetHeaders(v map[string]*string) *SearchFormDataIdListResponse {
	s.Headers = v
	return s
}

func (s *SearchFormDataIdListResponse) SetStatusCode(v int32) *SearchFormDataIdListResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchFormDataIdListResponse) SetBody(v *SearchFormDataIdListResponseBody) *SearchFormDataIdListResponse {
	s.Body = v
	return s
}

type SearchFormDataSecondGenerationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchFormDataSecondGenerationHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationHeaders) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationHeaders) SetCommonHeaders(v map[string]*string) *SearchFormDataSecondGenerationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchFormDataSecondGenerationHeaders) SetXAcsDingtalkAccessToken(v string) *SearchFormDataSecondGenerationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchFormDataSecondGenerationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_XCE0EVXS6DYG3YDYC5RD
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// example:
	//
	// 2021-05-01
	CreateFromTimeGMT *string `json:"createFromTimeGMT,omitempty" xml:"createFromTimeGMT,omitempty"`
	// example:
	//
	// 2021-05-01
	CreateToTimeGMT *string `json:"createToTimeGMT,omitempty" xml:"createToTimeGMT,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-GX866MC1NC1VOFF6WVQW33FD16E23L3CPMKVKA
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// 2021-05-01
	ModifiedFromTimeGMT *string `json:"modifiedFromTimeGMT,omitempty" xml:"modifiedFromTimeGMT,omitempty"`
	// example:
	//
	// 2021-09-10
	ModifiedToTimeGMT *string `json:"modifiedToTimeGMT,omitempty" xml:"modifiedToTimeGMT,omitempty"`
	// example:
	//
	// 例如按照创建时间升序按照文本组件值升序排序则填{\"gmt_create\":\"+\",\"textField_1234\":\"+\"}。详情参考 https://www.yuque.com/yida/support/agb8im#CQro8
	OrderConfigJson *string `json:"orderConfigJson,omitempty" xml:"orderConfigJson,omitempty"`
	// example:
	//
	// manager123
	OriginatorId *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// [{"key":"currentNodeName","value":"当前审批节点名称","type":"TEXT","operator":"like","componentName":"TextField"}]。详情参考 https://www.yuque.com/yida/support/agb8im#F4S8e
	SearchCondition *string `json:"searchCondition,omitempty" xml:"searchCondition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 09866181UTZVVD4R3DC955FNKIM52HVPU5WWK7
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// example:
	//
	// false
	UseAlias *bool `json:"useAlias,omitempty" xml:"useAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ding173982232112232
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchFormDataSecondGenerationRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationRequest) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationRequest) SetAppType(v string) *SearchFormDataSecondGenerationRequest {
	s.AppType = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetCreateFromTimeGMT(v string) *SearchFormDataSecondGenerationRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetCreateToTimeGMT(v string) *SearchFormDataSecondGenerationRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetFormUuid(v string) *SearchFormDataSecondGenerationRequest {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetModifiedFromTimeGMT(v string) *SearchFormDataSecondGenerationRequest {
	s.ModifiedFromTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetModifiedToTimeGMT(v string) *SearchFormDataSecondGenerationRequest {
	s.ModifiedToTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetOrderConfigJson(v string) *SearchFormDataSecondGenerationRequest {
	s.OrderConfigJson = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetOriginatorId(v string) *SearchFormDataSecondGenerationRequest {
	s.OriginatorId = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetPageNumber(v int32) *SearchFormDataSecondGenerationRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetPageSize(v int32) *SearchFormDataSecondGenerationRequest {
	s.PageSize = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetSearchCondition(v string) *SearchFormDataSecondGenerationRequest {
	s.SearchCondition = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetSystemToken(v string) *SearchFormDataSecondGenerationRequest {
	s.SystemToken = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetUseAlias(v bool) *SearchFormDataSecondGenerationRequest {
	s.UseAlias = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetUserId(v string) *SearchFormDataSecondGenerationRequest {
	s.UserId = &v
	return s
}

type SearchFormDataSecondGenerationResponseBody struct {
	Data []*SearchFormDataSecondGenerationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s SearchFormDataSecondGenerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponseBody) SetData(v []*SearchFormDataSecondGenerationResponseBodyData) *SearchFormDataSecondGenerationResponseBody {
	s.Data = v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBody) SetPageNumber(v int64) *SearchFormDataSecondGenerationResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBody) SetTotalCount(v int64) *SearchFormDataSecondGenerationResponseBody {
	s.TotalCount = &v
	return s
}

type SearchFormDataSecondGenerationResponseBodyData struct {
	// example:
	//
	// 2021-05-01
	CreateTimeGMT *string `json:"createTimeGMT,omitempty" xml:"createTimeGMT,omitempty"`
	// example:
	//
	// ding12345
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// example:
	//
	// {"addressField_l0c1cwiy_id":"\"海南省/469027/国营红岗农场/111\"","associationFormField_l0c1hdz4_id":"\"[{\\\"formType\\\":\\\"receipt\\\",\\\"formUuid\\\":\\\"FORM-QQ866JB1QW8YM5XZZZ64VQB61OGM1MLWE1C0LG\\\",\\\"instanceId\\\":\\\"FINST-CC666Y6198RY0LAN39XGND212MSX3TFT95S0LN31\\\",\\\"subTitle\\\":\\\"{\\\\\\\"type\\\\\\\":\\\\\\\"div\\\\\\\",\\\\\\\"props\\\\\\\":{\\\\\\\"children\\\\\\\":{\\\\\\\"type\\\\\\\":\\\\\\\"a\\\\\\\",\\\\\\\"props\\\\\\\":{\\\\\\\"children\\\\\\\":\\\\\\\"查看签名\\\\\\\",\\\\\\\"className\\\\\\\":\\\\\\\"inst-cell-item-link\\\\\\\",\\\\\\\"style\\\\\\\":{\\\\\\\"cursor\\\\\\\":\\\\\\\"pointer\\\\\\\",\\\\\\\"color\\\\\\\":\\\\\\\"#0068ff\\\\\\\"}}},\\\\\\\"className\\\\\\\":\\\\\\\"inst-cell-item\\\\\\\"}}\\\",\\\"appType\\\":\\\"APP_K6IGJJ6PFAARLPDSWKXQ\\\",\\\"title\\\":\\\"1\\\"}]\"","countrySelectField_l0c1cwiu_id":["PG"],"imageField_l0c1cwit":"[{\"previewUrl\":\"/ossFileHandle?appType=APP_K6IGJJ6PFAARLPDSWKXQ&fileName=APP_K6IGJJ6PFAARLPDSWKXQ_MTczMjU1NjYyMzg3MzI0NF8wUDk2NlQ2MVIzR1lHV1RaMjMxQ1A5U1Y1NU1NM1lMWVY1QzBMMQ$$.jpg&instId=&type=open&process=image/resize,m_fill,w_200,h_200,limit_0/quality,q_80\",\"size\":2610734,\"name\":\"Bts2Z0e6oxA.jpg\",\"downloadUrl\":\"/ossFileHandle?appType=APP_K6IGJJ6PFAARLPDSWKXQ&fileName=APP_K6IGJJ6PFAARLPDSWKXQ_MTczMjU1NjYyMzg3MzI0NF8wUDk2NlQ2MVIzR1lHV1RaMjMxQ1A5U1Y1NU1NM1lMWVY1QzBMMQ$$.jpg&instId=&type=download\",\"url\":\"/ossFileHandle?appType=APP_K6IGJJ6PFAARLPDSWKXQ&fileName=APP_K6IGJJ6PFAARLPDSWKXQ_MTczMjU1NjYyMzg3MzI0NF8wUDk2NlQ2MVIzR1lHV1RaMjMxQ1A5U1Y1NU1NM1lMWVY1QzBMMQ$$.jpg&instId=&type=download\"}]","rateField_l0c1cwis_value":"3","editorField_l0c1cwiz":"<div><strong>你好，这是测试</strong></div>\r\n<div><span style=\"font-family: kaiti;background-color: #ff0000;color: #ffff00;\"><em><strong>测试</strong></em></span></div>\r\n<div>&nbsp;</div>","rateField_l0c1cwis":3,"countrySelectField_l0c1cwiu":[],"attachmentField_l0ghkwv3":"[{\"downloadUrl\":\"/ossFileHandle?appType=APP_K6IGJJ6PFAARLPDSWKXQ&fileName=APP_K6IGJJ6PFAARLPDSWKXQ_MTczMjU1NjYyMzg3MzI0NF8wUDk2NlQ2MVIzR1lHV1RaMjMxQ1A5U1Y1NU1NM1lMWVY1QzBMMQ$$.jpg&instId=&type=download\",\"name\":\"Bts2Z0e6oxA.jpg\",\"previewUrl\":\"/ossFileHandle?appType=APP_K6IGJJ6PFAARLPDSWKXQ&fileName=APP_K6IGJJ6PFAARLPDSWKXQ_MTczMjU1NjYyMzg3MzI0NF8wUDk2NlQ2MVIzR1lHV1RaMjMxQ1A5U1Y1NU1NM1lMWVY1QzBMMQ$$.jpg&instId=&type=open&process=image/resize,m_fill,w_200,h_200,limit_0/quality,q_80\",\"size\":2610734,\"url\":\"/ossFileHandle?appType=APP_K6IGJJ6PFAARLPDSWKXQ&fileName=APP_K6IGJJ6PFAARLPDSWKXQ_MTczMjU1NjYyMzg3MzI0NF8wUDk2NlQ2MVIzR1lHV1RaMjMxQ1A5U1Y1NU1NM1lMWVY1QzBMMQ$$.jpg&instId=&type=download\"}]","addressField_l0c1cwiy":"{\"address\":\"111\",\"regionIds\":[460000,469027,469023401],\"regionText\":[{\"en_US\":\"hai+nan+sheng\",\"zh_CN\":\"海南省\"},{\"en_US\":\"cheng+mai+xian\",\"zh_CN\":\"澄迈县\"},{\"en_US\":\"guo+ying+hong+gang+nong+chang\",\"zh_CN\":\"国营红岗农场\"}]}"}
	FormData map[string]interface{} `json:"formData,omitempty" xml:"formData,omitempty"`
	// example:
	//
	// FINST-J8766S91O2UYN87ZX3XOF1MY8MBA2912BSV0L24
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// example:
	//
	// FORM-GX866MC1NC1VOFF6WVQW33FD16E23L3CPMKVKA
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// 12345
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// [{"componentName":"CountrySelectField","dateType":null,"fieldData":{"value":[{"text":{"en_US":"Papua New Guinea","zh_CN":"巴布亚新几内亚","type":"i18n"},"value":"PG"}]},"fieldDataUpdated":false,"fieldId":"countrySelectField_l0c1cwiu","format":null,"formatControls":null,"listNum":null,"options":[{"label":{"$ref":"$[0].fieldData.value[0].text"},"value":"PG"}],"rowId":null},{"componentName":"AssociationFormField","dateType":null,"fieldData":{"value":[{"formType":"receipt","formUuid":"FORM-QQ866JB1QW8YM5XZZZ64VQB61OGM1MLWE1C0LG","instanceId":"FINST-CC666Y6198RY0LAN39XGND212MSX3TFT95S0LN31","subTitle":"{\"type\":\"div\",\"props\":{\"children\":{\"type\":\"a\",\"props\":{\"children\":\"查看签名\",\"className\":\"inst-cell-item-link\",\"style\":{\"cursor\":\"pointer\",\"color\":\"#0068ff\"}}},\"className\":\"inst-cell-item\"}}","appType":"APP_K6IGJJ6PFAARLPDSWKXQ","title":"1"}]},"fieldDataUpdated":false,"fieldId":"associationFormField_l0c1hdz4","format":null,"formatControls":null,"listNum":null,"options":[],"rowId":null},{"componentName":"ImageField","dateType":null,"fieldData":{"value":[{"previewUrl":"/ossFileHandle?appType=APP_K6IGJJ6PFAARLPDSWKXQ&fileName=APP_K6IGJJ6PFAARLPDSWKXQ_MTczMjU1NjYyMzg3MzI0NF8wUDk2NlQ2MVIzR1lHV1RaMjMxQ1A5U1Y1NU1NM1lMWVY1QzBMMQ$$.jpg&instId=&type=open&process=image/resize,m_fill,w_200,h_200,limit_0/quality,q_80","size":2610734,"name":"Bts2Z0e6oxA.jpg","downloadUrl":"/ossFileHandle?appType=APP_K6IGJJ6PFAARLPDSWKXQ&fileName=APP_K6IGJJ6PFAARLPDSWKXQ_MTczMjU1NjYyMzg3MzI0NF8wUDk2NlQ2MVIzR1lHV1RaMjMxQ1A5U1Y1NU1NM1lMWVY1QzBMMQ$$.jpg&instId=&type=download","url":"/ossFileHandle?appType=APP_K6IGJJ6PFAARLPDSWKXQ&fileName=APP_K6IGJJ6PFAARLPDSWKXQ_MTczMjU1NjYyMzg3MzI0NF8wUDk2NlQ2MVIzR1lHV1RaMjMxQ1A5U1Y1NU1NM1lMWVY1QzBMMQ$$.jpg&instId=&type=download"}]},"fieldDataUpdated":false,"fieldId":"imageField_l0c1cwit","format":null,"formatControls":null,"listNum":null,"options":[],"rowId":null},{"componentName":"AddressField","dateType":null,"fieldData":{"value":{"address":"111","regionIds":[460000,469027,469023401],"regionText":[{"en_US":"hai+nan+sheng","zh_CN":"海南省"},{"en_US":"cheng+mai+xian","zh_CN":"澄迈县"},{"en_US":"guo+ying+hong+gang+nong+chang","zh_CN":"国营红岗农场"}]}},"fieldDataUpdated":false,"fieldId":"addressField_l0c1cwiy","format":null,"formatControls":null,"listNum":null,"options":[{"label":{"pureEn_US":"hai+nan+sheng","en_US":"hai+nan+sheng","zh_CN":"海南省","type":"i18n","key":null},"value":460000},{"label":{"pureEn_US":"cheng+mai+xian","en_US":"cheng+mai+xian","zh_CN":"澄迈县","type":"i18n","key":null},"value":469027},{"label":{"pureEn_US":"guo+ying+hong+gang+nong+chang","en_US":"guo+ying+hong+gang+nong+chang","zh_CN":"国营红岗农场","type":"i18n","key":null},"value":469023401}],"rowId":null},{"componentName":"EditorField","dateType":null,"fieldData":{"value":"<div><strong>你好，这是测试</strong></div>\r\n<div><span style=\"font-family: kaiti;background-color: #ff0000;color: #ffff00;\"><em><strong>测试</strong></em></span></div>\r\n<div>&nbsp;</div>"},"fieldDataUpdated":false,"fieldId":"editorField_l0c1cwiz","format":null,"formatControls":null,"listNum":null,"options":[],"rowId":null},{"componentName":"RateField","dateType":null,"fieldData":{"value":"3"},"fieldDataUpdated":false,"fieldId":"rateField_l0c1cwis","format":null,"formatControls":null,"listNum":null,"options":[],"rowId":null},{"componentName":"AttachmentField","dateType":null,"fieldData":{"value":[{"previewUrl":"/ossFileHandle?appType=APP_K6IGJJ6PFAARLPDSWKXQ&fileName=APP_K6IGJJ6PFAARLPDSWKXQ_MTczMjU1NjYyMzg3MzI0NF8wUDk2NlQ2MVIzR1lHV1RaMjMxQ1A5U1Y1NU1NM1lMWVY1QzBMMQ$$.jpg&instId=&type=open&process=image/resize,m_fill,w_200,h_200,limit_0/quality,q_80","size":2610734,"name":"Bts2Z0e6oxA.jpg","downloadUrl":"/ossFileHandle?appType=APP_K6IGJJ6PFAARLPDSWKXQ&fileName=APP_K6IGJJ6PFAARLPDSWKXQ_MTczMjU1NjYyMzg3MzI0NF8wUDk2NlQ2MVIzR1lHV1RaMjMxQ1A5U1Y1NU1NM1lMWVY1QzBMMQ$$.jpg&instId=&type=download","url":"/ossFileHandle?appType=APP_K6IGJJ6PFAARLPDSWKXQ&fileName=APP_K6IGJJ6PFAARLPDSWKXQ_MTczMjU1NjYyMzg3MzI0NF8wUDk2NlQ2MVIzR1lHV1RaMjMxQ1A5U1Y1NU1NM1lMWVY1QzBMMQ$$.jpg&instId=&type=download"}]},"fieldDataUpdated":false,"fieldId":"attachmentField_l0ghkwv3","format":null,"formatControls":null,"listNum":null,"options":[],"rowId":null}]
	InstanceValue *string `json:"instanceValue,omitempty" xml:"instanceValue,omitempty"`
	// example:
	//
	// 2021-05-01
	ModifiedTimeGMT *string `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	// example:
	//
	// manager123
	Modifier   *string                                                   `json:"modifier,omitempty" xml:"modifier,omitempty"`
	ModifyUser *SearchFormDataSecondGenerationResponseBodyDataModifyUser `json:"modifyUser,omitempty" xml:"modifyUser,omitempty" type:"Struct"`
	Originator *SearchFormDataSecondGenerationResponseBodyDataOriginator `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	// example:
	//
	// IMPORT-388664B1BAUVB3AYZE1RIUE88TDM1QI9WIOWK2
	Sequence *string `json:"sequence,omitempty" xml:"sequence,omitempty"`
	// example:
	//
	// YIDA909202202250027
	SerialNumber *string `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
	// example:
	//
	// 李四发起的请购单
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 1.0
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s SearchFormDataSecondGenerationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetCreateTimeGMT(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.CreateTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetCreatorUserId(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.CreatorUserId = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetFormData(v map[string]interface{}) *SearchFormDataSecondGenerationResponseBodyData {
	s.FormData = v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetFormInstanceId(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.FormInstanceId = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetFormUuid(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetId(v int64) *SearchFormDataSecondGenerationResponseBodyData {
	s.Id = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetInstanceValue(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.InstanceValue = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetModifiedTimeGMT(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetModifier(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.Modifier = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetModifyUser(v *SearchFormDataSecondGenerationResponseBodyDataModifyUser) *SearchFormDataSecondGenerationResponseBodyData {
	s.ModifyUser = v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetOriginator(v *SearchFormDataSecondGenerationResponseBodyDataOriginator) *SearchFormDataSecondGenerationResponseBodyData {
	s.Originator = v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetSequence(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.Sequence = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetSerialNumber(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.SerialNumber = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetTitle(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.Title = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetVersion(v int64) *SearchFormDataSecondGenerationResponseBodyData {
	s.Version = &v
	return s
}

type SearchFormDataSecondGenerationResponseBodyDataModifyUser struct {
	Name *SearchFormDataSecondGenerationResponseBodyDataModifyUserName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	// example:
	//
	// ding173982232112232
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchFormDataSecondGenerationResponseBodyDataModifyUser) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponseBodyDataModifyUser) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponseBodyDataModifyUser) SetName(v *SearchFormDataSecondGenerationResponseBodyDataModifyUserName) *SearchFormDataSecondGenerationResponseBodyDataModifyUser {
	s.Name = v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyDataModifyUser) SetUserId(v string) *SearchFormDataSecondGenerationResponseBodyDataModifyUser {
	s.UserId = &v
	return s
}

type SearchFormDataSecondGenerationResponseBodyDataModifyUserName struct {
	// example:
	//
	// 张三
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// example:
	//
	// ZhangSan
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
}

func (s SearchFormDataSecondGenerationResponseBodyDataModifyUserName) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponseBodyDataModifyUserName) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponseBodyDataModifyUserName) SetNameInChinese(v string) *SearchFormDataSecondGenerationResponseBodyDataModifyUserName {
	s.NameInChinese = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyDataModifyUserName) SetNameInEnglish(v string) *SearchFormDataSecondGenerationResponseBodyDataModifyUserName {
	s.NameInEnglish = &v
	return s
}

type SearchFormDataSecondGenerationResponseBodyDataOriginator struct {
	Name *SearchFormDataSecondGenerationResponseBodyDataOriginatorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	// example:
	//
	// ding173982232112232
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchFormDataSecondGenerationResponseBodyDataOriginator) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponseBodyDataOriginator) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponseBodyDataOriginator) SetName(v *SearchFormDataSecondGenerationResponseBodyDataOriginatorName) *SearchFormDataSecondGenerationResponseBodyDataOriginator {
	s.Name = v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyDataOriginator) SetUserId(v string) *SearchFormDataSecondGenerationResponseBodyDataOriginator {
	s.UserId = &v
	return s
}

type SearchFormDataSecondGenerationResponseBodyDataOriginatorName struct {
	// example:
	//
	// 张三
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// example:
	//
	// ZhangSan
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
}

func (s SearchFormDataSecondGenerationResponseBodyDataOriginatorName) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponseBodyDataOriginatorName) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponseBodyDataOriginatorName) SetNameInChinese(v string) *SearchFormDataSecondGenerationResponseBodyDataOriginatorName {
	s.NameInChinese = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyDataOriginatorName) SetNameInEnglish(v string) *SearchFormDataSecondGenerationResponseBodyDataOriginatorName {
	s.NameInEnglish = &v
	return s
}

type SearchFormDataSecondGenerationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchFormDataSecondGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchFormDataSecondGenerationResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponse) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponse) SetHeaders(v map[string]*string) *SearchFormDataSecondGenerationResponse {
	s.Headers = v
	return s
}

func (s *SearchFormDataSecondGenerationResponse) SetStatusCode(v int32) *SearchFormDataSecondGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponse) SetBody(v *SearchFormDataSecondGenerationResponseBody) *SearchFormDataSecondGenerationResponse {
	s.Body = v
	return s
}

type SearchFormDatasHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchFormDatasHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasHeaders) GoString() string {
	return s.String()
}

func (s *SearchFormDatasHeaders) SetCommonHeaders(v map[string]*string) *SearchFormDatasHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchFormDatasHeaders) SetXAcsDingtalkAccessToken(v string) *SearchFormDatasHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchFormDatasRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_PBKT0MFBEBTDO8T7SLVP
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// example:
	//
	// 2018-01-01
	CreateFromTimeGMT *string `json:"createFromTimeGMT,omitempty" xml:"createFromTimeGMT,omitempty"`
	// example:
	//
	// 2018-02-01
	CreateToTimeGMT *string `json:"createToTimeGMT,omitempty" xml:"createToTimeGMT,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// {"numberField_1ac":"+"}, 表示按照字段numberField_1ac升序排列
	DynamicOrder *string `json:"dynamicOrder,omitempty" xml:"dynamicOrder,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-EF6Y4G8WO2FN0SUB43TDQ3CGC3FMFQ1G9400RCJ3
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// 2018-01-01
	ModifiedFromTimeGMT *string `json:"modifiedFromTimeGMT,omitempty" xml:"modifiedFromTimeGMT,omitempty"`
	// example:
	//
	// 2018-02-01
	ModifiedToTimeGMT *string `json:"modifiedToTimeGMT,omitempty" xml:"modifiedToTimeGMT,omitempty"`
	OriginatorId      *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// {"textField_jcr0069m":"danhang","textareaField_jcr0069n":"duohang","numberField_jcr0069o":["1","10"],"radioField_jcr0069p":"选项一","selectField_jcr0069q":"选项一","checkboxField_jcr0069r":["选项二"],"multiSelectField_jcr0069s":["选项二","选项三"],"dateField_jcr0069t":[1514736000000,1517414399000],"cascadeDate_jcr0069u":[[1514736000000,1517414399000],[1514736000000,1517414399000]],"employeeField_jcr0069x":["xxxxx"],"citySelectField_jcr0069y":["110000","110100","110101"],"departmentField_jcr0069z":1123456,"cascadeSelectField_jcr006a0":["part","part_b"],"tableField_jcr006a1":"明细数据"}
	SearchFieldJson *string `json:"searchFieldJson,omitempty" xml:"searchFieldJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxx
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// example:
	//
	// false
	UseAlias *bool `json:"useAlias,omitempty" xml:"useAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 173112212211
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchFormDatasRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasRequest) GoString() string {
	return s.String()
}

func (s *SearchFormDatasRequest) SetAppType(v string) *SearchFormDatasRequest {
	s.AppType = &v
	return s
}

func (s *SearchFormDatasRequest) SetCreateFromTimeGMT(v string) *SearchFormDatasRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *SearchFormDatasRequest) SetCreateToTimeGMT(v string) *SearchFormDatasRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *SearchFormDatasRequest) SetCurrentPage(v int32) *SearchFormDatasRequest {
	s.CurrentPage = &v
	return s
}

func (s *SearchFormDatasRequest) SetDynamicOrder(v string) *SearchFormDatasRequest {
	s.DynamicOrder = &v
	return s
}

func (s *SearchFormDatasRequest) SetFormUuid(v string) *SearchFormDatasRequest {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDatasRequest) SetLanguage(v string) *SearchFormDatasRequest {
	s.Language = &v
	return s
}

func (s *SearchFormDatasRequest) SetModifiedFromTimeGMT(v string) *SearchFormDatasRequest {
	s.ModifiedFromTimeGMT = &v
	return s
}

func (s *SearchFormDatasRequest) SetModifiedToTimeGMT(v string) *SearchFormDatasRequest {
	s.ModifiedToTimeGMT = &v
	return s
}

func (s *SearchFormDatasRequest) SetOriginatorId(v string) *SearchFormDatasRequest {
	s.OriginatorId = &v
	return s
}

func (s *SearchFormDatasRequest) SetPageSize(v int32) *SearchFormDatasRequest {
	s.PageSize = &v
	return s
}

func (s *SearchFormDatasRequest) SetSearchFieldJson(v string) *SearchFormDatasRequest {
	s.SearchFieldJson = &v
	return s
}

func (s *SearchFormDatasRequest) SetSystemToken(v string) *SearchFormDatasRequest {
	s.SystemToken = &v
	return s
}

func (s *SearchFormDatasRequest) SetUseAlias(v bool) *SearchFormDatasRequest {
	s.UseAlias = &v
	return s
}

func (s *SearchFormDatasRequest) SetUserId(v string) *SearchFormDatasRequest {
	s.UserId = &v
	return s
}

type SearchFormDatasResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int64                             `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	Data        []*SearchFormDatasResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s SearchFormDatasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponseBody) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBody) SetCurrentPage(v int64) *SearchFormDatasResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *SearchFormDatasResponseBody) SetData(v []*SearchFormDatasResponseBodyData) *SearchFormDatasResponseBody {
	s.Data = v
	return s
}

func (s *SearchFormDatasResponseBody) SetTotalCount(v int64) *SearchFormDatasResponseBody {
	s.TotalCount = &v
	return s
}

type SearchFormDatasResponseBodyData struct {
	// example:
	//
	// 2018-01-24 11:22:01
	CreatedTimeGMT *string `json:"createdTimeGMT,omitempty" xml:"createdTimeGMT,omitempty"`
	// example:
	//
	// 1731234567
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// example:
	//
	// 1002
	DataId *int64 `json:"dataId,omitempty" xml:"dataId,omitempty"`
	// example:
	//
	// {"numberField_jcr0069o":1,"multiSelectField_jcr0069s":["选项三","选项二"],"textareaField_jcr0069n":"duohang","employeeField_jcr0069x":["xxxx"],"departmentField_jcr0069z":"xxxx","cascadeDate_jcr0069u":["1514736000000","1517328000000"],"cascadeSelectField_jcr006a0":["part","part_b"],"tableField_jcr006a1":[{"departmentField_jcr006ad":"xxxx","cascadeDate_jcr006aa":["1514736000000","1517328000000"],"selectField_jcr006a6":"选项三","citySelectField_jcr006ac":["天津","天津市","河东区"],"radioField_jcr006a5":"选项二","employeeField_jcr006ab":["xxxxxx","yyyyyy"],"dateField_jcr006a9":1517328000000,"textField_jcr006a2":"明细下单行","textareaField_jcr006a3":"明细下多行","cascadeSelectField_jcr006ae":["product","product_a"],"numberField_jcr006a4":2,"checkboxField_jcr006a7":["选项一","选项三","选项二"],"multiSelectField_jcr006a8":["选项一","选项三","选项二"]}],"selectField_jcr0069q":"选项一","citySelectField_jcr0069y":["北京","北京市","东城区"],"checkboxField_jcr0069r":["选项三","选项二"],"textField_jcr0069m":"danhang","radioField_jcr0069p":"选项一","dateField_jcr0069t":1516636800000}
	FormData map[string]interface{} `json:"formData,omitempty" xml:"formData,omitempty"`
	// example:
	//
	// FINST-BNKJDRF
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// example:
	//
	// FORM-EF6Y93URN24F1SCX15VA2P918LPEIJ2H3UFORCJ1
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// {"textField":"124"}
	InstanceValue *string `json:"instanceValue,omitempty" xml:"instanceValue,omitempty"`
	ModelUuid     *string `json:"modelUuid,omitempty" xml:"modelUuid,omitempty"`
	// example:
	//
	// 2018-01-24 11:22:01
	ModifiedTimeGMT *string `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	// example:
	//
	// 1731234567
	ModifierUserId *string                                    `json:"modifierUserId,omitempty" xml:"modifierUserId,omitempty"`
	ModifyUser     *SearchFormDatasResponseBodyDataModifyUser `json:"modifyUser,omitempty" xml:"modifyUser,omitempty" type:"Struct"`
	Originator     *SearchFormDatasResponseBodyDataOriginator `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	// example:
	//
	// Squence-XXX
	Sequence *string `json:"sequence,omitempty" xml:"sequence,omitempty"`
	// example:
	//
	// 1234
	SerialNo *string `json:"serialNo,omitempty" xml:"serialNo,omitempty"`
	// example:
	//
	// 张三发起的表单
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 3
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s SearchFormDatasResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBodyData) SetCreatedTimeGMT(v string) *SearchFormDatasResponseBodyData {
	s.CreatedTimeGMT = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetCreatorUserId(v string) *SearchFormDatasResponseBodyData {
	s.CreatorUserId = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetDataId(v int64) *SearchFormDatasResponseBodyData {
	s.DataId = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetFormData(v map[string]interface{}) *SearchFormDatasResponseBodyData {
	s.FormData = v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetFormInstanceId(v string) *SearchFormDatasResponseBodyData {
	s.FormInstanceId = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetFormUuid(v string) *SearchFormDatasResponseBodyData {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetInstanceValue(v string) *SearchFormDatasResponseBodyData {
	s.InstanceValue = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetModelUuid(v string) *SearchFormDatasResponseBodyData {
	s.ModelUuid = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetModifiedTimeGMT(v string) *SearchFormDatasResponseBodyData {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetModifierUserId(v string) *SearchFormDatasResponseBodyData {
	s.ModifierUserId = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetModifyUser(v *SearchFormDatasResponseBodyDataModifyUser) *SearchFormDatasResponseBodyData {
	s.ModifyUser = v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetOriginator(v *SearchFormDatasResponseBodyDataOriginator) *SearchFormDatasResponseBodyData {
	s.Originator = v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetSequence(v string) *SearchFormDatasResponseBodyData {
	s.Sequence = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetSerialNo(v string) *SearchFormDatasResponseBodyData {
	s.SerialNo = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetTitle(v string) *SearchFormDatasResponseBodyData {
	s.Title = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetVersion(v int64) *SearchFormDatasResponseBodyData {
	s.Version = &v
	return s
}

type SearchFormDatasResponseBodyDataModifyUser struct {
	UserId   *string                                            `json:"userId,omitempty" xml:"userId,omitempty"`
	UserName *SearchFormDatasResponseBodyDataModifyUserUserName `json:"userName,omitempty" xml:"userName,omitempty" type:"Struct"`
}

func (s SearchFormDatasResponseBodyDataModifyUser) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponseBodyDataModifyUser) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBodyDataModifyUser) SetUserId(v string) *SearchFormDatasResponseBodyDataModifyUser {
	s.UserId = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataModifyUser) SetUserName(v *SearchFormDatasResponseBodyDataModifyUserUserName) *SearchFormDatasResponseBodyDataModifyUser {
	s.UserName = v
	return s
}

type SearchFormDatasResponseBodyDataModifyUserUserName struct {
	// example:
	//
	// 张三
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// example:
	//
	// ZhangSan
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	// example:
	//
	// i18n
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SearchFormDatasResponseBodyDataModifyUserUserName) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponseBodyDataModifyUserUserName) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBodyDataModifyUserUserName) SetNameInChinese(v string) *SearchFormDatasResponseBodyDataModifyUserUserName {
	s.NameInChinese = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataModifyUserUserName) SetNameInEnglish(v string) *SearchFormDatasResponseBodyDataModifyUserUserName {
	s.NameInEnglish = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataModifyUserUserName) SetType(v string) *SearchFormDatasResponseBodyDataModifyUserUserName {
	s.Type = &v
	return s
}

type SearchFormDatasResponseBodyDataOriginator struct {
	UserId   *string                                            `json:"userId,omitempty" xml:"userId,omitempty"`
	UserName *SearchFormDatasResponseBodyDataOriginatorUserName `json:"userName,omitempty" xml:"userName,omitempty" type:"Struct"`
}

func (s SearchFormDatasResponseBodyDataOriginator) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponseBodyDataOriginator) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBodyDataOriginator) SetUserId(v string) *SearchFormDatasResponseBodyDataOriginator {
	s.UserId = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataOriginator) SetUserName(v *SearchFormDatasResponseBodyDataOriginatorUserName) *SearchFormDatasResponseBodyDataOriginator {
	s.UserName = v
	return s
}

type SearchFormDatasResponseBodyDataOriginatorUserName struct {
	// example:
	//
	// 张三
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// example:
	//
	// ZhangSan
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	// example:
	//
	// i18n
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SearchFormDatasResponseBodyDataOriginatorUserName) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponseBodyDataOriginatorUserName) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBodyDataOriginatorUserName) SetNameInChinese(v string) *SearchFormDatasResponseBodyDataOriginatorUserName {
	s.NameInChinese = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataOriginatorUserName) SetNameInEnglish(v string) *SearchFormDatasResponseBodyDataOriginatorUserName {
	s.NameInEnglish = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataOriginatorUserName) SetType(v string) *SearchFormDatasResponseBodyDataOriginatorUserName {
	s.Type = &v
	return s
}

type SearchFormDatasResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchFormDatasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchFormDatasResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponse) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponse) SetHeaders(v map[string]*string) *SearchFormDatasResponse {
	s.Headers = v
	return s
}

func (s *SearchFormDatasResponse) SetStatusCode(v int32) *SearchFormDatasResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchFormDatasResponse) SetBody(v *SearchFormDatasResponseBody) *SearchFormDatasResponse {
	s.Body = v
	return s
}

type StartInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s StartInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceHeaders) GoString() string {
	return s.String()
}

func (s *StartInstanceHeaders) SetCommonHeaders(v map[string]*string) *StartInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *StartInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type StartInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_PBKT0MFBEBTDO8T7SLVP
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// example:
	//
	// 18295
	DepartmentId *string `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"textField_jcpm6agt": "单行","employeeField_jcos0sar": ["workno"]}
	FormDataJson *string `json:"formDataJson,omitempty" xml:"formDataJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-NJYJZELV8YZRDEI2N5IQ7L6VEDMR1VE9GMPCJB
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// TPROC--EF6Y4G8WO2FN0SUB43TDQ3CGC3FMFQ1G9400RCJ4
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// example:
	//
	// [{ 	"key": "__optionalApproval_node_ocltdztr2b1", 	"value": ["5014533041684350"] }, { 	"key": "__optionalApproval_node_ocltdztr2b3", 	"value": ["5014533041684350", "01536610064226180419"] }, { 	"key": "__optionalApproval_node_oclte07cwn1", 	"value": ["01432910392321237660"] }]
	ProcessData *string `json:"processData,omitempty" xml:"processData,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxx
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// example:
	//
	// false
	UseAlias *bool `json:"useAlias,omitempty" xml:"useAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1731234567
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s StartInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartInstanceRequest) SetAppType(v string) *StartInstanceRequest {
	s.AppType = &v
	return s
}

func (s *StartInstanceRequest) SetDepartmentId(v string) *StartInstanceRequest {
	s.DepartmentId = &v
	return s
}

func (s *StartInstanceRequest) SetFormDataJson(v string) *StartInstanceRequest {
	s.FormDataJson = &v
	return s
}

func (s *StartInstanceRequest) SetFormUuid(v string) *StartInstanceRequest {
	s.FormUuid = &v
	return s
}

func (s *StartInstanceRequest) SetLanguage(v string) *StartInstanceRequest {
	s.Language = &v
	return s
}

func (s *StartInstanceRequest) SetProcessCode(v string) *StartInstanceRequest {
	s.ProcessCode = &v
	return s
}

func (s *StartInstanceRequest) SetProcessData(v string) *StartInstanceRequest {
	s.ProcessData = &v
	return s
}

func (s *StartInstanceRequest) SetSystemToken(v string) *StartInstanceRequest {
	s.SystemToken = &v
	return s
}

func (s *StartInstanceRequest) SetUseAlias(v bool) *StartInstanceRequest {
	s.UseAlias = &v
	return s
}

func (s *StartInstanceRequest) SetUserId(v string) *StartInstanceRequest {
	s.UserId = &v
	return s
}

type StartInstanceResponseBody struct {
	// example:
	//
	// f30233fb-72e1-4af4-8cb8-c7e0ea9ee530
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s StartInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstanceResponseBody) SetResult(v string) *StartInstanceResponseBody {
	s.Result = &v
	return s
}

type StartInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartInstanceResponse) SetHeaders(v map[string]*string) *StartInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartInstanceResponse) SetStatusCode(v int32) *StartInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartInstanceResponse) SetBody(v *StartInstanceResponseBody) *StartInstanceResponse {
	s.Body = v
	return s
}

type UpdateAgentTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateAgentTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateAgentTaskHeaders) GoString() string {
	return s.String()
}

func (s *UpdateAgentTaskHeaders) SetCommonHeaders(v map[string]*string) *UpdateAgentTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateAgentTaskHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateAgentTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateAgentTaskRequest struct {
	// example:
	//
	// ALL
	AgentRangeType *string `json:"agentRangeType,omitempty" xml:"agentRangeType,omitempty"`
	// example:
	//
	// [{\"appType\":\"APP_xxx\",\"formUuid\":\"FORM-xxx\"}]
	AgentRangeValue *string `json:"agentRangeValue,omitempty" xml:"agentRangeValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10001
	AgentUserId *string `json:"agentUserId,omitempty" xml:"agentUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Agent--xxxxx
	AgentUuid *string `json:"agentUuid,omitempty" xml:"agentUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dingxxxx
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// 1761204600404
	EndTimestamp        *string `json:"endTimestamp,omitempty" xml:"endTimestamp,omitempty"`
	NeedNoticePrincipal *string `json:"needNoticePrincipal,omitempty" xml:"needNoticePrincipal,omitempty"`
	// example:
	//
	// 1761204600404
	StartTimestamp *string `json:"startTimestamp,omitempty" xml:"startTimestamp,omitempty"`
	// This parameter is required.
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 501453
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateAgentTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAgentTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentTaskRequest) SetAgentRangeType(v string) *UpdateAgentTaskRequest {
	s.AgentRangeType = &v
	return s
}

func (s *UpdateAgentTaskRequest) SetAgentRangeValue(v string) *UpdateAgentTaskRequest {
	s.AgentRangeValue = &v
	return s
}

func (s *UpdateAgentTaskRequest) SetAgentUserId(v string) *UpdateAgentTaskRequest {
	s.AgentUserId = &v
	return s
}

func (s *UpdateAgentTaskRequest) SetAgentUuid(v string) *UpdateAgentTaskRequest {
	s.AgentUuid = &v
	return s
}

func (s *UpdateAgentTaskRequest) SetCorpId(v string) *UpdateAgentTaskRequest {
	s.CorpId = &v
	return s
}

func (s *UpdateAgentTaskRequest) SetEndTimestamp(v string) *UpdateAgentTaskRequest {
	s.EndTimestamp = &v
	return s
}

func (s *UpdateAgentTaskRequest) SetNeedNoticePrincipal(v string) *UpdateAgentTaskRequest {
	s.NeedNoticePrincipal = &v
	return s
}

func (s *UpdateAgentTaskRequest) SetStartTimestamp(v string) *UpdateAgentTaskRequest {
	s.StartTimestamp = &v
	return s
}

func (s *UpdateAgentTaskRequest) SetToken(v string) *UpdateAgentTaskRequest {
	s.Token = &v
	return s
}

func (s *UpdateAgentTaskRequest) SetUserId(v string) *UpdateAgentTaskRequest {
	s.UserId = &v
	return s
}

type UpdateAgentTaskResponseBody struct {
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateAgentTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAgentTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgentTaskResponseBody) SetErrorCode(v string) *UpdateAgentTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateAgentTaskResponseBody) SetErrorMsg(v string) *UpdateAgentTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateAgentTaskResponseBody) SetResult(v bool) *UpdateAgentTaskResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateAgentTaskResponseBody) SetSuccess(v bool) *UpdateAgentTaskResponseBody {
	s.Success = &v
	return s
}

type UpdateAgentTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAgentTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAgentTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAgentTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateAgentTaskResponse) SetHeaders(v map[string]*string) *UpdateAgentTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateAgentTaskResponse) SetStatusCode(v int32) *UpdateAgentTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAgentTaskResponse) SetBody(v *UpdateAgentTaskResponseBody) *UpdateAgentTaskResponse {
	s.Body = v
	return s
}

type UpdateFormDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateFormDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateFormDataHeaders) GoString() string {
	return s.String()
}

func (s *UpdateFormDataHeaders) SetCommonHeaders(v map[string]*string) *UpdateFormDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateFormDataHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateFormDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateFormDataRequest struct {
	// example:
	//
	// APP_PBKT0MFBEBTDO8T7SLVP
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 33f6d221-17f8-42b7-836a-682b95a046c2
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// example:
	//
	// FORM-AA28579F69644FC19A47FE267457E664ZVR1
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxx
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"textField_jcr0069m":"danhang","textareaField_jcr0069n":"duohang","numberField_jcr0069o":1,"radioField_jcr0069p":"选项一","selectField_jcr0069q":"选项一","checkboxField_jcr0069r":["选项二","选项三"],"multiSelectField_jcr0069s":["选项二","选项三"],"dateField_jcr0069t":1516636800000,"cascadeDate_jcr0069u":["1514736000000","1517328000000"],"employeeField_jcr0069x":["xxxxx"],"citySelectField_jcr0069y":["110000","110100","110101"],"departmentField_jcr0069z":1123456,"cascadeSelectField_jcr006a0":["part","part_b"],{"attachmentField_jna1lvyb":[{"downloadUrl":"https://www.aliwork.com/fileHandle?appType=default_tianshu_app&fileName=edd07ca9-1d2e-44b5-98fe-c1e16202f90d.txt&instId=&type=download","name":"test.txt","previewUrl":"https://www.aliwork.com/inst/preview?appType=default_tianshu_app&fileName=test.txt&fileSize=4&downloadUrl=edd07ca9-1d2e-44b5-98fe-c1e16202f90d.txt","url":"https://www.aliwork.com/fileHandle?appType=default_tianshu_app&fileName=edd07ca9-1d2e-44b5-98fe-c1e16202f90d.txt&instId=&type=download","ext":"txt"}]},"tableField_jcr006a1":[{"cascadeDate_jcr006aa":["1514736000000","1517328000000"],"cascadeSelectField_jcr006ae":["product","product_a"],"checkboxField_jcr006a7":["选项一","选项二","选项三"],"citySelectField_jcr006ac":["120000","120100","120102"],"dateField_jcr006a9":1517328000000,"departmentField_jcr006ad":1123456,"employeeField_jcr006ab":["yyyyy","xxxxx"],"multiSelectField_jcr006a8":["选项一","选项二","选项三"],"numberField_jcr006a4":2,"radioField_jcr006a5":"选项二","selectField_jcr006a6":"选项三","textField_jcr006a2":"明细下单行","textareaField_jcr006a3":"明细下多行"}]}
	UpdateFormDataJson *string `json:"updateFormDataJson,omitempty" xml:"updateFormDataJson,omitempty"`
	// example:
	//
	// false
	UseAlias *bool `json:"useAlias,omitempty" xml:"useAlias,omitempty"`
	// example:
	//
	// false
	UseLatestVersion *bool `json:"useLatestVersion,omitempty" xml:"useLatestVersion,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateFormDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFormDataRequest) GoString() string {
	return s.String()
}

func (s *UpdateFormDataRequest) SetAppType(v string) *UpdateFormDataRequest {
	s.AppType = &v
	return s
}

func (s *UpdateFormDataRequest) SetFormInstanceId(v string) *UpdateFormDataRequest {
	s.FormInstanceId = &v
	return s
}

func (s *UpdateFormDataRequest) SetFormUuid(v string) *UpdateFormDataRequest {
	s.FormUuid = &v
	return s
}

func (s *UpdateFormDataRequest) SetLanguage(v string) *UpdateFormDataRequest {
	s.Language = &v
	return s
}

func (s *UpdateFormDataRequest) SetSystemToken(v string) *UpdateFormDataRequest {
	s.SystemToken = &v
	return s
}

func (s *UpdateFormDataRequest) SetUpdateFormDataJson(v string) *UpdateFormDataRequest {
	s.UpdateFormDataJson = &v
	return s
}

func (s *UpdateFormDataRequest) SetUseAlias(v bool) *UpdateFormDataRequest {
	s.UseAlias = &v
	return s
}

func (s *UpdateFormDataRequest) SetUseLatestVersion(v bool) *UpdateFormDataRequest {
	s.UseLatestVersion = &v
	return s
}

func (s *UpdateFormDataRequest) SetUserId(v string) *UpdateFormDataRequest {
	s.UserId = &v
	return s
}

type UpdateFormDataResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateFormDataResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFormDataResponse) GoString() string {
	return s.String()
}

func (s *UpdateFormDataResponse) SetHeaders(v map[string]*string) *UpdateFormDataResponse {
	s.Headers = v
	return s
}

func (s *UpdateFormDataResponse) SetStatusCode(v int32) *UpdateFormDataResponse {
	s.StatusCode = &v
	return s
}

type UpdateSubTableHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateSubTableHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubTableHeaders) GoString() string {
	return s.String()
}

func (s *UpdateSubTableHeaders) SetCommonHeaders(v map[string]*string) *UpdateSubTableHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateSubTableHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateSubTableHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateSubTableRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_XCE0EVXS6DYG3YDYC5RD
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FINST-J8766S91O2UYN87ZX3XOF1MY8MBA2912BSV0L24
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	Language       *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// true
	NoExecuteExpression *bool `json:"noExecuteExpression,omitempty" xml:"noExecuteExpression,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 09866181UTZVVD4R3DC955FNKIM52HVPU5WWK7
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// This parameter is required.
	TableFieldIds *string `json:"tableFieldIds,omitempty" xml:"tableFieldIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"countrySelectField_l0c1cwiu":[{"value":"US"}],"addressField_l0c1cwiy":{"address":"111","regionIds":[460000,469027,469023401],"regionText":[{"en_US":"hai+nan+sheng","zh_CN":"海南省"},{"en_US":"cheng+mai+xian","zh_CN":"澄迈县"},{"en_US":"guo+ying+hong+gang+nong+chang","zh_CN":"国营红岗农场"}]}}
	UpdateFormDataJson *string `json:"updateFormDataJson,omitempty" xml:"updateFormDataJson,omitempty"`
	UseAlias           *bool   `json:"useAlias,omitempty" xml:"useAlias,omitempty"`
	// example:
	//
	// false
	UseLatestFormSchemaVersion *bool `json:"useLatestFormSchemaVersion,omitempty" xml:"useLatestFormSchemaVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ding173982232112232
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateSubTableRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubTableRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubTableRequest) SetAppType(v string) *UpdateSubTableRequest {
	s.AppType = &v
	return s
}

func (s *UpdateSubTableRequest) SetFormInstanceId(v string) *UpdateSubTableRequest {
	s.FormInstanceId = &v
	return s
}

func (s *UpdateSubTableRequest) SetLanguage(v string) *UpdateSubTableRequest {
	s.Language = &v
	return s
}

func (s *UpdateSubTableRequest) SetNoExecuteExpression(v bool) *UpdateSubTableRequest {
	s.NoExecuteExpression = &v
	return s
}

func (s *UpdateSubTableRequest) SetSystemToken(v string) *UpdateSubTableRequest {
	s.SystemToken = &v
	return s
}

func (s *UpdateSubTableRequest) SetTableFieldIds(v string) *UpdateSubTableRequest {
	s.TableFieldIds = &v
	return s
}

func (s *UpdateSubTableRequest) SetUpdateFormDataJson(v string) *UpdateSubTableRequest {
	s.UpdateFormDataJson = &v
	return s
}

func (s *UpdateSubTableRequest) SetUseAlias(v bool) *UpdateSubTableRequest {
	s.UseAlias = &v
	return s
}

func (s *UpdateSubTableRequest) SetUseLatestFormSchemaVersion(v bool) *UpdateSubTableRequest {
	s.UseLatestFormSchemaVersion = &v
	return s
}

func (s *UpdateSubTableRequest) SetUserId(v string) *UpdateSubTableRequest {
	s.UserId = &v
	return s
}

type UpdateSubTableResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateSubTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubTableResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSubTableResponseBody) SetResult(v bool) *UpdateSubTableResponseBody {
	s.Result = &v
	return s
}

type UpdateSubTableResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSubTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSubTableResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubTableResponse) GoString() string {
	return s.String()
}

func (s *UpdateSubTableResponse) SetHeaders(v map[string]*string) *UpdateSubTableResponse {
	s.Headers = v
	return s
}

func (s *UpdateSubTableResponse) SetStatusCode(v int32) *UpdateSubTableResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSubTableResponse) SetBody(v *UpdateSubTableResponseBody) *UpdateSubTableResponse {
	s.Body = v
	return s
}

type UpdateSubTableByRowIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateSubTableByRowIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubTableByRowIdHeaders) GoString() string {
	return s.String()
}

func (s *UpdateSubTableByRowIdHeaders) SetCommonHeaders(v map[string]*string) *UpdateSubTableByRowIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateSubTableByRowIdHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateSubTableByRowIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateSubTableByRowIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_XCE0EVXS6DYG3YDYC5RD
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FINST-J8766S91O2UYN87ZX3XOF1MY8MBA2912BSV0L24
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// if can be null:
	// false
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 09866181UTZVVD4R3DC955FNKIM52HVPU5WWK7
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// This parameter is required.
	TableFieldId *string `json:"tableFieldId,omitempty" xml:"tableFieldId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"textField_md2x1jow":"更新子表通过rowId","textareaField_md2x1jox":"更新子表通过rowId","rowId":"xxxxxxxxxxxxxxxx"},{"textField_md2x1jow":"更新子表通过rowId","textareaField_md2x1jox":"更新子表通过rowId","rowId":"xxxxxxxxxxxxxxxx"}]
	UpdateSubTableDataJson *string `json:"updateSubTableDataJson,omitempty" xml:"updateSubTableDataJson,omitempty"`
	UseAlias               *bool   `json:"useAlias,omitempty" xml:"useAlias,omitempty"`
	// example:
	//
	// false
	UseLatestFormSchemaVersion *bool `json:"useLatestFormSchemaVersion,omitempty" xml:"useLatestFormSchemaVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ding173982232112232
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateSubTableByRowIdRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubTableByRowIdRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubTableByRowIdRequest) SetAppType(v string) *UpdateSubTableByRowIdRequest {
	s.AppType = &v
	return s
}

func (s *UpdateSubTableByRowIdRequest) SetFormInstanceId(v string) *UpdateSubTableByRowIdRequest {
	s.FormInstanceId = &v
	return s
}

func (s *UpdateSubTableByRowIdRequest) SetFormUuid(v string) *UpdateSubTableByRowIdRequest {
	s.FormUuid = &v
	return s
}

func (s *UpdateSubTableByRowIdRequest) SetSystemToken(v string) *UpdateSubTableByRowIdRequest {
	s.SystemToken = &v
	return s
}

func (s *UpdateSubTableByRowIdRequest) SetTableFieldId(v string) *UpdateSubTableByRowIdRequest {
	s.TableFieldId = &v
	return s
}

func (s *UpdateSubTableByRowIdRequest) SetUpdateSubTableDataJson(v string) *UpdateSubTableByRowIdRequest {
	s.UpdateSubTableDataJson = &v
	return s
}

func (s *UpdateSubTableByRowIdRequest) SetUseAlias(v bool) *UpdateSubTableByRowIdRequest {
	s.UseAlias = &v
	return s
}

func (s *UpdateSubTableByRowIdRequest) SetUseLatestFormSchemaVersion(v bool) *UpdateSubTableByRowIdRequest {
	s.UseLatestFormSchemaVersion = &v
	return s
}

func (s *UpdateSubTableByRowIdRequest) SetUserId(v string) *UpdateSubTableByRowIdRequest {
	s.UserId = &v
	return s
}

type UpdateSubTableByRowIdResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateSubTableByRowIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubTableByRowIdResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSubTableByRowIdResponseBody) SetResult(v bool) *UpdateSubTableByRowIdResponseBody {
	s.Result = &v
	return s
}

type UpdateSubTableByRowIdResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSubTableByRowIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSubTableByRowIdResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSubTableByRowIdResponse) GoString() string {
	return s.String()
}

func (s *UpdateSubTableByRowIdResponse) SetHeaders(v map[string]*string) *UpdateSubTableByRowIdResponse {
	s.Headers = v
	return s
}

func (s *UpdateSubTableByRowIdResponse) SetStatusCode(v int32) *UpdateSubTableByRowIdResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSubTableByRowIdResponse) SetBody(v *UpdateSubTableByRowIdResponseBody) *UpdateSubTableByRowIdResponse {
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
// 批量新增/更新角色成员
//
// @param request - BatchAddOrUpdateRoleMembersRequest
//
// @param headers - BatchAddOrUpdateRoleMembersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchAddOrUpdateRoleMembersResponse
func (client *Client) BatchAddOrUpdateRoleMembersWithOptions(request *BatchAddOrUpdateRoleMembersRequest, headers *BatchAddOrUpdateRoleMembersHeaders, runtime *util.RuntimeOptions) (_result *BatchAddOrUpdateRoleMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.MembersInfo)) {
		body["membersInfo"] = request.MembersInfo
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RoleUuid)) {
		body["roleUuid"] = request.RoleUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		body["token"] = request.Token
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
		Action:      tea.String("BatchAddOrUpdateRoleMembers"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/resources/roles/upsert"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchAddOrUpdateRoleMembersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量新增/更新角色成员
//
// @param request - BatchAddOrUpdateRoleMembersRequest
//
// @return BatchAddOrUpdateRoleMembersResponse
func (client *Client) BatchAddOrUpdateRoleMembers(request *BatchAddOrUpdateRoleMembersRequest) (_result *BatchAddOrUpdateRoleMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchAddOrUpdateRoleMembersHeaders{}
	_result = &BatchAddOrUpdateRoleMembersResponse{}
	_body, _err := client.BatchAddOrUpdateRoleMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量删除角色成员
//
// @param request - BatchDeleteRoleMembersRequest
//
// @param headers - BatchDeleteRoleMembersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteRoleMembersResponse
func (client *Client) BatchDeleteRoleMembersWithOptions(request *BatchDeleteRoleMembersRequest, headers *BatchDeleteRoleMembersHeaders, runtime *util.RuntimeOptions) (_result *BatchDeleteRoleMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberIds)) {
		body["memberIds"] = request.MemberIds
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RoleUuid)) {
		body["roleUuid"] = request.RoleUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		body["token"] = request.Token
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
		Action:      tea.String("BatchDeleteRoleMembers"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/resources/roles/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchDeleteRoleMembersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除角色成员
//
// @param request - BatchDeleteRoleMembersRequest
//
// @return BatchDeleteRoleMembersResponse
func (client *Client) BatchDeleteRoleMembers(request *BatchDeleteRoleMembersRequest) (_result *BatchDeleteRoleMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchDeleteRoleMembersHeaders{}
	_result = &BatchDeleteRoleMembersResponse{}
	_body, _err := client.BatchDeleteRoleMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 撤销代理关系
//
// @param request - CancelAgentTaskRequest
//
// @param headers - CancelAgentTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelAgentTaskResponse
func (client *Client) CancelAgentTaskWithOptions(request *CancelAgentTaskRequest, headers *CancelAgentTaskHeaders, runtime *util.RuntimeOptions) (_result *CancelAgentTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentType)) {
		query["agentType"] = request.AgentType
	}

	if !tea.BoolValue(util.IsUnset(request.AgentUuid)) {
		query["agentUuid"] = request.AgentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
		Action:      tea.String("CancelAgentTask"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/resources/agents/cancel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelAgentTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 撤销代理关系
//
// @param request - CancelAgentTaskRequest
//
// @return CancelAgentTaskResponse
func (client *Client) CancelAgentTask(request *CancelAgentTaskRequest) (_result *CancelAgentTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CancelAgentTaskHeaders{}
	_result = &CancelAgentTaskResponse{}
	_body, _err := client.CancelAgentTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建代理关系
//
// @param request - CreateAgentTaskRequest
//
// @param headers - CreateAgentTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgentTaskResponse
func (client *Client) CreateAgentTaskWithOptions(request *CreateAgentTaskRequest, headers *CreateAgentTaskHeaders, runtime *util.RuntimeOptions) (_result *CreateAgentTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentCategory)) {
		query["agentCategory"] = request.AgentCategory
	}

	if !tea.BoolValue(util.IsUnset(request.AgentRangeType)) {
		query["agentRangeType"] = request.AgentRangeType
	}

	if !tea.BoolValue(util.IsUnset(request.AgentRangeValue)) {
		query["agentRangeValue"] = request.AgentRangeValue
	}

	if !tea.BoolValue(util.IsUnset(request.AgentType)) {
		query["agentType"] = request.AgentType
	}

	if !tea.BoolValue(util.IsUnset(request.AgentUserId)) {
		query["agentUserId"] = request.AgentUserId
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["endTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.NeedNoticePrincipal)) {
		query["needNoticePrincipal"] = request.NeedNoticePrincipal
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalUserId)) {
		query["principalUserId"] = request.PrincipalUserId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["startTimestamp"] = request.StartTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
		Action:      tea.String("CreateAgentTask"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/resources/agents/insert"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAgentTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建代理关系
//
// @param request - CreateAgentTaskRequest
//
// @return CreateAgentTaskResponse
func (client *Client) CreateAgentTask(request *CreateAgentTaskRequest) (_result *CreateAgentTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateAgentTaskHeaders{}
	_result = &CreateAgentTaskResponse{}
	_body, _err := client.CreateAgentTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增或更新表单实例
//
// @param request - CreateOrUpdateFormDataRequest
//
// @param headers - CreateOrUpdateFormDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrUpdateFormDataResponse
func (client *Client) CreateOrUpdateFormDataWithOptions(request *CreateOrUpdateFormDataRequest, headers *CreateOrUpdateFormDataHeaders, runtime *util.RuntimeOptions) (_result *CreateOrUpdateFormDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.FormDataJson)) {
		body["formDataJson"] = request.FormDataJson
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.NoExecuteExpression)) {
		body["noExecuteExpression"] = request.NoExecuteExpression
	}

	if !tea.BoolValue(util.IsUnset(request.SearchCondition)) {
		body["searchCondition"] = request.SearchCondition
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UseAlias)) {
		body["useAlias"] = request.UseAlias
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
		Action:      tea.String("CreateOrUpdateFormData"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/instances/insertOrUpdate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOrUpdateFormDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增或更新表单实例
//
// @param request - CreateOrUpdateFormDataRequest
//
// @return CreateOrUpdateFormDataResponse
func (client *Client) CreateOrUpdateFormData(request *CreateOrUpdateFormDataRequest) (_result *CreateOrUpdateFormDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateOrUpdateFormDataHeaders{}
	_result = &CreateOrUpdateFormDataResponse{}
	_body, _err := client.CreateOrUpdateFormDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量删除矩阵明细数据
//
// @param request - DeleteMatrixDataByRowIdsRequest
//
// @param headers - DeleteMatrixDataByRowIdsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMatrixDataByRowIdsResponse
func (client *Client) DeleteMatrixDataByRowIdsWithOptions(request *DeleteMatrixDataByRowIdsRequest, headers *DeleteMatrixDataByRowIdsHeaders, runtime *util.RuntimeOptions) (_result *DeleteMatrixDataByRowIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.MatrixId)) {
		body["matrixId"] = request.MatrixId
	}

	if !tea.BoolValue(util.IsUnset(request.RowIds)) {
		body["rowIds"] = request.RowIds
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		body["token"] = request.Token
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
		Action:      tea.String("DeleteMatrixDataByRowIds"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/resources/matrices/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMatrixDataByRowIdsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除矩阵明细数据
//
// @param request - DeleteMatrixDataByRowIdsRequest
//
// @return DeleteMatrixDataByRowIdsResponse
func (client *Client) DeleteMatrixDataByRowIds(request *DeleteMatrixDataByRowIdsRequest) (_result *DeleteMatrixDataByRowIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteMatrixDataByRowIdsHeaders{}
	_result = &DeleteMatrixDataByRowIdsResponse{}
	_body, _err := client.DeleteMatrixDataByRowIdsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取代理列表
//
// @param request - GetAgentTasksRequest
//
// @param headers - GetAgentTasksHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentTasksResponse
func (client *Client) GetAgentTasksWithOptions(request *GetAgentTasksRequest, headers *GetAgentTasksHeaders, runtime *util.RuntimeOptions) (_result *GetAgentTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentUuid)) {
		query["agentUuid"] = request.AgentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Keywords)) {
		query["keywords"] = request.Keywords
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
		Action:      tea.String("GetAgentTasks"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/resources/agents"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAgentTasksResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取代理列表
//
// @param request - GetAgentTasksRequest
//
// @return GetAgentTasksResponse
func (client *Client) GetAgentTasks(request *GetAgentTasksRequest) (_result *GetAgentTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAgentTasksHeaders{}
	_result = &GetAgentTasksResponse{}
	_body, _err := client.GetAgentTasksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取表单组件别名列表
//
// @param request - GetFormComponentAliasListRequest
//
// @param headers - GetFormComponentAliasListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFormComponentAliasListResponse
func (client *Client) GetFormComponentAliasListWithOptions(appType *string, formUuid *string, request *GetFormComponentAliasListRequest, headers *GetFormComponentAliasListHeaders, runtime *util.RuntimeOptions) (_result *GetFormComponentAliasListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		query["version"] = request.Version
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
		Action:      tea.String("GetFormComponentAliasList"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/component/alias/" + tea.StringValue(appType) + "/" + tea.StringValue(formUuid)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFormComponentAliasListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取表单组件别名列表
//
// @param request - GetFormComponentAliasListRequest
//
// @return GetFormComponentAliasListResponse
func (client *Client) GetFormComponentAliasList(appType *string, formUuid *string, request *GetFormComponentAliasListRequest) (_result *GetFormComponentAliasListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFormComponentAliasListHeaders{}
	_result = &GetFormComponentAliasListResponse{}
	_body, _err := client.GetFormComponentAliasListWithOptions(appType, formUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据表单 ID 查询实例详情
//
// @param request - GetFormDataByIDRequest
//
// @param headers - GetFormDataByIDHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFormDataByIDResponse
func (client *Client) GetFormDataByIDWithOptions(id *string, request *GetFormDataByIDRequest, headers *GetFormDataByIDHeaders, runtime *util.RuntimeOptions) (_result *GetFormDataByIDResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		query["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UseAlias)) {
		query["useAlias"] = request.UseAlias
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
		Action:      tea.String("GetFormDataByID"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/instances/" + tea.StringValue(id)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFormDataByIDResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据表单 ID 查询实例详情
//
// @param request - GetFormDataByIDRequest
//
// @return GetFormDataByIDResponse
func (client *Client) GetFormDataByID(id *string, request *GetFormDataByIDRequest) (_result *GetFormDataByIDResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFormDataByIDHeaders{}
	_result = &GetFormDataByIDResponse{}
	_body, _err := client.GetFormDataByIDWithOptions(id, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据实例 ID 获取流程实例详情
//
// @param request - GetInstanceByIdRequest
//
// @param headers - GetInstanceByIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceByIdResponse
func (client *Client) GetInstanceByIdWithOptions(id *string, request *GetInstanceByIdRequest, headers *GetInstanceByIdHeaders, runtime *util.RuntimeOptions) (_result *GetInstanceByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		query["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UseAlias)) {
		query["useAlias"] = request.UseAlias
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
		Action:      tea.String("GetInstanceById"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/processes/instancesInfos/" + tea.StringValue(id)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceByIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据实例 ID 获取流程实例详情
//
// @param request - GetInstanceByIdRequest
//
// @return GetInstanceByIdResponse
func (client *Client) GetInstanceById(id *string, request *GetInstanceByIdRequest) (_result *GetInstanceByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetInstanceByIdHeaders{}
	_result = &GetInstanceByIdResponse{}
	_body, _err := client.GetInstanceByIdWithOptions(id, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据条件搜索流程实例 ID
//
// @param request - GetInstanceIdListRequest
//
// @param headers - GetInstanceIdListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceIdListResponse
func (client *Client) GetInstanceIdListWithOptions(request *GetInstanceIdListRequest, headers *GetInstanceIdListHeaders, runtime *util.RuntimeOptions) (_result *GetInstanceIdListResponse, _err error) {
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

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.ApprovedResult)) {
		body["approvedResult"] = request.ApprovedResult
	}

	if !tea.BoolValue(util.IsUnset(request.CreateFromTimeGMT)) {
		body["createFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CreateToTimeGMT)) {
		body["createToTimeGMT"] = request.CreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceStatus)) {
		body["instanceStatus"] = request.InstanceStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedFromTimeGMT)) {
		body["modifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedToTimeGMT)) {
		body["modifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.OriginatorId)) {
		body["originatorId"] = request.OriginatorId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchFieldJson)) {
		body["searchFieldJson"] = request.SearchFieldJson
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.UseAlias)) {
		body["useAlias"] = request.UseAlias
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
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceIdList"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/processes/instanceIds"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceIdListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据条件搜索流程实例 ID
//
// @param request - GetInstanceIdListRequest
//
// @return GetInstanceIdListResponse
func (client *Client) GetInstanceIdList(request *GetInstanceIdListRequest) (_result *GetInstanceIdListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetInstanceIdListHeaders{}
	_result = &GetInstanceIdListResponse{}
	_body, _err := client.GetInstanceIdListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据搜索条件获取流程表单实例详情
//
// @param request - GetInstancesRequest
//
// @param headers - GetInstancesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstancesResponse
func (client *Client) GetInstancesWithOptions(request *GetInstancesRequest, headers *GetInstancesHeaders, runtime *util.RuntimeOptions) (_result *GetInstancesResponse, _err error) {
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

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.ApprovedResult)) {
		body["approvedResult"] = request.ApprovedResult
	}

	if !tea.BoolValue(util.IsUnset(request.CreateFromTimeGMT)) {
		body["createFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CreateToTimeGMT)) {
		body["createToTimeGMT"] = request.CreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceStatus)) {
		body["instanceStatus"] = request.InstanceStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedFromTimeGMT)) {
		body["modifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedToTimeGMT)) {
		body["modifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.OrderConfigJson)) {
		body["orderConfigJson"] = request.OrderConfigJson
	}

	if !tea.BoolValue(util.IsUnset(request.OriginatorId)) {
		body["originatorId"] = request.OriginatorId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchFieldJson)) {
		body["searchFieldJson"] = request.SearchFieldJson
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.UseAlias)) {
		body["useAlias"] = request.UseAlias
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
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstances"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/processes/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstancesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据搜索条件获取流程表单实例详情
//
// @param request - GetInstancesRequest
//
// @return GetInstancesResponse
func (client *Client) GetInstances(request *GetInstancesRequest) (_result *GetInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetInstancesHeaders{}
	_result = &GetInstancesResponse{}
	_body, _err := client.GetInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取权限矩阵数据详情
//
// @param request - GetMatrixDetailByIdRequest
//
// @param headers - GetMatrixDetailByIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMatrixDetailByIdResponse
func (client *Client) GetMatrixDetailByIdWithOptions(request *GetMatrixDetailByIdRequest, headers *GetMatrixDetailByIdHeaders, runtime *util.RuntimeOptions) (_result *GetMatrixDetailByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.MatrixId)) {
		query["matrixId"] = request.MatrixId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
		Action:      tea.String("GetMatrixDetailById"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/resources/matrices"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMatrixDetailByIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取权限矩阵数据详情
//
// @param request - GetMatrixDetailByIdRequest
//
// @return GetMatrixDetailByIdResponse
func (client *Client) GetMatrixDetailById(request *GetMatrixDetailByIdRequest) (_result *GetMatrixDetailByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetMatrixDetailByIdHeaders{}
	_result = &GetMatrixDetailByIdResponse{}
	_body, _err := client.GetMatrixDetailByIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取角色详情和成员列表
//
// @param request - GetRoleDetailByIdRequest
//
// @param headers - GetRoleDetailByIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRoleDetailByIdResponse
func (client *Client) GetRoleDetailByIdWithOptions(request *GetRoleDetailByIdRequest, headers *GetRoleDetailByIdHeaders, runtime *util.RuntimeOptions) (_result *GetRoleDetailByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RoleUuid)) {
		query["roleUuid"] = request.RoleUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
		Action:      tea.String("GetRoleDetailById"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/resources/roles"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRoleDetailByIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取角色详情和成员列表
//
// @param request - GetRoleDetailByIdRequest
//
// @return GetRoleDetailByIdResponse
func (client *Client) GetRoleDetailById(request *GetRoleDetailByIdRequest) (_result *GetRoleDetailByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetRoleDetailByIdHeaders{}
	_result = &GetRoleDetailByIdResponse{}
	_body, _err := client.GetRoleDetailByIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改/新增矩阵明细数据
//
// @param request - SaveAndUpdateMatrixDataRequest
//
// @param headers - SaveAndUpdateMatrixDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveAndUpdateMatrixDataResponse
func (client *Client) SaveAndUpdateMatrixDataWithOptions(request *SaveAndUpdateMatrixDataRequest, headers *SaveAndUpdateMatrixDataHeaders, runtime *util.RuntimeOptions) (_result *SaveAndUpdateMatrixDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DataJson)) {
		body["dataJson"] = request.DataJson
	}

	if !tea.BoolValue(util.IsUnset(request.MatrixId)) {
		body["matrixId"] = request.MatrixId
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		body["token"] = request.Token
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
		Action:      tea.String("SaveAndUpdateMatrixData"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/resources/matrices/upsert"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveAndUpdateMatrixDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改/新增矩阵明细数据
//
// @param request - SaveAndUpdateMatrixDataRequest
//
// @return SaveAndUpdateMatrixDataResponse
func (client *Client) SaveAndUpdateMatrixData(request *SaveAndUpdateMatrixDataRequest) (_result *SaveAndUpdateMatrixDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveAndUpdateMatrixDataHeaders{}
	_result = &SaveAndUpdateMatrixDataResponse{}
	_body, _err := client.SaveAndUpdateMatrixDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增表单实例
//
// @param request - SaveFormDataRequest
//
// @param headers - SaveFormDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveFormDataResponse
func (client *Client) SaveFormDataWithOptions(request *SaveFormDataRequest, headers *SaveFormDataHeaders, runtime *util.RuntimeOptions) (_result *SaveFormDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.FormDataJson)) {
		body["formDataJson"] = request.FormDataJson
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UseAlias)) {
		body["useAlias"] = request.UseAlias
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
		Action:      tea.String("SaveFormData"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveFormDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增表单实例
//
// @param request - SaveFormDataRequest
//
// @return SaveFormDataResponse
func (client *Client) SaveFormData(request *SaveFormDataRequest) (_result *SaveFormDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveFormDataHeaders{}
	_result = &SaveFormDataResponse{}
	_body, _err := client.SaveFormDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据条件搜索表单实例 ID 列表
//
// @param request - SearchFormDataIdListRequest
//
// @param headers - SearchFormDataIdListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchFormDataIdListResponse
func (client *Client) SearchFormDataIdListWithOptions(appType *string, formUuid *string, request *SearchFormDataIdListRequest, headers *SearchFormDataIdListHeaders, runtime *util.RuntimeOptions) (_result *SearchFormDataIdListResponse, _err error) {
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

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateFromTimeGMT)) {
		body["createFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CreateToTimeGMT)) {
		body["createToTimeGMT"] = request.CreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedFromTimeGMT)) {
		body["modifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedToTimeGMT)) {
		body["modifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.OriginatorId)) {
		body["originatorId"] = request.OriginatorId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchFieldJson)) {
		body["searchFieldJson"] = request.SearchFieldJson
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UseAlias)) {
		body["useAlias"] = request.UseAlias
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
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchFormDataIdList"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/instances/ids/" + tea.StringValue(appType) + "/" + tea.StringValue(formUuid)),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchFormDataIdListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据条件搜索表单实例 ID 列表
//
// @param request - SearchFormDataIdListRequest
//
// @return SearchFormDataIdListResponse
func (client *Client) SearchFormDataIdList(appType *string, formUuid *string, request *SearchFormDataIdListRequest) (_result *SearchFormDataIdListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchFormDataIdListHeaders{}
	_result = &SearchFormDataIdListResponse{}
	_body, _err := client.SearchFormDataIdListWithOptions(appType, formUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过高级检索条件查询表单实例
//
// @param request - SearchFormDataSecondGenerationRequest
//
// @param headers - SearchFormDataSecondGenerationHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchFormDataSecondGenerationResponse
func (client *Client) SearchFormDataSecondGenerationWithOptions(request *SearchFormDataSecondGenerationRequest, headers *SearchFormDataSecondGenerationHeaders, runtime *util.RuntimeOptions) (_result *SearchFormDataSecondGenerationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.CreateFromTimeGMT)) {
		body["createFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CreateToTimeGMT)) {
		body["createToTimeGMT"] = request.CreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedFromTimeGMT)) {
		body["modifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedToTimeGMT)) {
		body["modifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.OrderConfigJson)) {
		body["orderConfigJson"] = request.OrderConfigJson
	}

	if !tea.BoolValue(util.IsUnset(request.OriginatorId)) {
		body["originatorId"] = request.OriginatorId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchCondition)) {
		body["searchCondition"] = request.SearchCondition
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UseAlias)) {
		body["useAlias"] = request.UseAlias
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
		Action:      tea.String("SearchFormDataSecondGeneration"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/instances/advances/queryAll"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchFormDataSecondGenerationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过高级检索条件查询表单实例
//
// @param request - SearchFormDataSecondGenerationRequest
//
// @return SearchFormDataSecondGenerationResponse
func (client *Client) SearchFormDataSecondGeneration(request *SearchFormDataSecondGenerationRequest) (_result *SearchFormDataSecondGenerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchFormDataSecondGenerationHeaders{}
	_result = &SearchFormDataSecondGenerationResponse{}
	_body, _err := client.SearchFormDataSecondGenerationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据条件搜索表单实例详情列表
//
// @param request - SearchFormDatasRequest
//
// @param headers - SearchFormDatasHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchFormDatasResponse
func (client *Client) SearchFormDatasWithOptions(request *SearchFormDatasRequest, headers *SearchFormDatasHeaders, runtime *util.RuntimeOptions) (_result *SearchFormDatasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.CreateFromTimeGMT)) {
		body["createFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CreateToTimeGMT)) {
		body["createToTimeGMT"] = request.CreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["currentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DynamicOrder)) {
		body["dynamicOrder"] = request.DynamicOrder
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedFromTimeGMT)) {
		body["modifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedToTimeGMT)) {
		body["modifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.OriginatorId)) {
		body["originatorId"] = request.OriginatorId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchFieldJson)) {
		body["searchFieldJson"] = request.SearchFieldJson
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UseAlias)) {
		body["useAlias"] = request.UseAlias
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
		Action:      tea.String("SearchFormDatas"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/instances/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchFormDatasResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据条件搜索表单实例详情列表
//
// @param request - SearchFormDatasRequest
//
// @return SearchFormDatasResponse
func (client *Client) SearchFormDatas(request *SearchFormDatasRequest) (_result *SearchFormDatasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchFormDatasHeaders{}
	_result = &SearchFormDatasResponse{}
	_body, _err := client.SearchFormDatasWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发起新的流程实例
//
// @param request - StartInstanceRequest
//
// @param headers - StartInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartInstanceResponse
func (client *Client) StartInstanceWithOptions(request *StartInstanceRequest, headers *StartInstanceHeaders, runtime *util.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		body["departmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.FormDataJson)) {
		body["formDataJson"] = request.FormDataJson
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		body["processCode"] = request.ProcessCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessData)) {
		body["processData"] = request.ProcessData
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UseAlias)) {
		body["useAlias"] = request.UseAlias
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
		Action:      tea.String("StartInstance"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/processes/instances/start"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &StartInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发起新的流程实例
//
// @param request - StartInstanceRequest
//
// @return StartInstanceResponse
func (client *Client) StartInstance(request *StartInstanceRequest) (_result *StartInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StartInstanceHeaders{}
	_result = &StartInstanceResponse{}
	_body, _err := client.StartInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改代理信息
//
// @param request - UpdateAgentTaskRequest
//
// @param headers - UpdateAgentTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAgentTaskResponse
func (client *Client) UpdateAgentTaskWithOptions(request *UpdateAgentTaskRequest, headers *UpdateAgentTaskHeaders, runtime *util.RuntimeOptions) (_result *UpdateAgentTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentRangeType)) {
		query["agentRangeType"] = request.AgentRangeType
	}

	if !tea.BoolValue(util.IsUnset(request.AgentRangeValue)) {
		query["agentRangeValue"] = request.AgentRangeValue
	}

	if !tea.BoolValue(util.IsUnset(request.AgentUserId)) {
		query["agentUserId"] = request.AgentUserId
	}

	if !tea.BoolValue(util.IsUnset(request.AgentUuid)) {
		query["agentUuid"] = request.AgentUuid
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		query["endTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.NeedNoticePrincipal)) {
		query["needNoticePrincipal"] = request.NeedNoticePrincipal
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		query["startTimestamp"] = request.StartTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
		Action:      tea.String("UpdateAgentTask"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/resources/agents/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAgentTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改代理信息
//
// @param request - UpdateAgentTaskRequest
//
// @return UpdateAgentTaskResponse
func (client *Client) UpdateAgentTask(request *UpdateAgentTaskRequest) (_result *UpdateAgentTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateAgentTaskHeaders{}
	_result = &UpdateAgentTaskResponse{}
	_body, _err := client.UpdateAgentTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新表单实例
//
// @param request - UpdateFormDataRequest
//
// @param headers - UpdateFormDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFormDataResponse
func (client *Client) UpdateFormDataWithOptions(request *UpdateFormDataRequest, headers *UpdateFormDataHeaders, runtime *util.RuntimeOptions) (_result *UpdateFormDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.FormInstanceId)) {
		body["formInstanceId"] = request.FormInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateFormDataJson)) {
		body["updateFormDataJson"] = request.UpdateFormDataJson
	}

	if !tea.BoolValue(util.IsUnset(request.UseAlias)) {
		body["useAlias"] = request.UseAlias
	}

	if !tea.BoolValue(util.IsUnset(request.UseLatestVersion)) {
		body["useLatestVersion"] = request.UseLatestVersion
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
		Action:      tea.String("UpdateFormData"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/instances"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateFormDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新表单实例
//
// @param request - UpdateFormDataRequest
//
// @return UpdateFormDataResponse
func (client *Client) UpdateFormData(request *UpdateFormDataRequest) (_result *UpdateFormDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateFormDataHeaders{}
	_result = &UpdateFormDataResponse{}
	_body, _err := client.UpdateFormDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新宜搭子表单
//
// @param request - UpdateSubTableRequest
//
// @param headers - UpdateSubTableHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSubTableResponse
func (client *Client) UpdateSubTableWithOptions(request *UpdateSubTableRequest, headers *UpdateSubTableHeaders, runtime *util.RuntimeOptions) (_result *UpdateSubTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.FormInstanceId)) {
		body["formInstanceId"] = request.FormInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.NoExecuteExpression)) {
		body["noExecuteExpression"] = request.NoExecuteExpression
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.TableFieldIds)) {
		body["tableFieldIds"] = request.TableFieldIds
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateFormDataJson)) {
		body["updateFormDataJson"] = request.UpdateFormDataJson
	}

	if !tea.BoolValue(util.IsUnset(request.UseAlias)) {
		body["useAlias"] = request.UseAlias
	}

	if !tea.BoolValue(util.IsUnset(request.UseLatestFormSchemaVersion)) {
		body["useLatestFormSchemaVersion"] = request.UseLatestFormSchemaVersion
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
		Action:      tea.String("UpdateSubTable"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/updateSubTable"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSubTableResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新宜搭子表单
//
// @param request - UpdateSubTableRequest
//
// @return UpdateSubTableResponse
func (client *Client) UpdateSubTable(request *UpdateSubTableRequest) (_result *UpdateSubTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateSubTableHeaders{}
	_result = &UpdateSubTableResponse{}
	_body, _err := client.UpdateSubTableWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过指定rowId更新宜搭子表单数据
//
// @param request - UpdateSubTableByRowIdRequest
//
// @param headers - UpdateSubTableByRowIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSubTableByRowIdResponse
func (client *Client) UpdateSubTableByRowIdWithOptions(request *UpdateSubTableByRowIdRequest, headers *UpdateSubTableByRowIdHeaders, runtime *util.RuntimeOptions) (_result *UpdateSubTableByRowIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.FormInstanceId)) {
		body["formInstanceId"] = request.FormInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.TableFieldId)) {
		body["tableFieldId"] = request.TableFieldId
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateSubTableDataJson)) {
		body["updateSubTableDataJson"] = request.UpdateSubTableDataJson
	}

	if !tea.BoolValue(util.IsUnset(request.UseAlias)) {
		body["useAlias"] = request.UseAlias
	}

	if !tea.BoolValue(util.IsUnset(request.UseLatestFormSchemaVersion)) {
		body["useLatestFormSchemaVersion"] = request.UseLatestFormSchemaVersion
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
		Action:      tea.String("UpdateSubTableByRowId"),
		Version:     tea.String("yida_2.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v2.0/yida/forms/updateSubTableByRowId"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSubTableByRowIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过指定rowId更新宜搭子表单数据
//
// @param request - UpdateSubTableByRowIdRequest
//
// @return UpdateSubTableByRowIdResponse
func (client *Client) UpdateSubTableByRowId(request *UpdateSubTableByRowIdRequest) (_result *UpdateSubTableByRowIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateSubTableByRowIdHeaders{}
	_result = &UpdateSubTableByRowIdResponse{}
	_body, _err := client.UpdateSubTableByRowIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
