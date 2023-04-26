// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package carbon_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetPersonalCarbonInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPersonalCarbonInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPersonalCarbonInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetPersonalCarbonInfoHeaders) SetCommonHeaders(v map[string]*string) *GetPersonalCarbonInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPersonalCarbonInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetPersonalCarbonInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPersonalCarbonInfoRequest struct {
	ActionType *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	UnionId    *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetPersonalCarbonInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPersonalCarbonInfoRequest) GoString() string {
	return s.String()
}

func (s *GetPersonalCarbonInfoRequest) SetActionType(v string) *GetPersonalCarbonInfoRequest {
	s.ActionType = &v
	return s
}

func (s *GetPersonalCarbonInfoRequest) SetUnionId(v string) *GetPersonalCarbonInfoRequest {
	s.UnionId = &v
	return s
}

type GetPersonalCarbonInfoResponseBody struct {
	Content              *string  `json:"content,omitempty" xml:"content,omitempty"`
	PersonalCarbonAmount *float64 `json:"personalCarbonAmount,omitempty" xml:"personalCarbonAmount,omitempty"`
}

func (s GetPersonalCarbonInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPersonalCarbonInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetPersonalCarbonInfoResponseBody) SetContent(v string) *GetPersonalCarbonInfoResponseBody {
	s.Content = &v
	return s
}

func (s *GetPersonalCarbonInfoResponseBody) SetPersonalCarbonAmount(v float64) *GetPersonalCarbonInfoResponseBody {
	s.PersonalCarbonAmount = &v
	return s
}

type GetPersonalCarbonInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPersonalCarbonInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPersonalCarbonInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPersonalCarbonInfoResponse) GoString() string {
	return s.String()
}

func (s *GetPersonalCarbonInfoResponse) SetHeaders(v map[string]*string) *GetPersonalCarbonInfoResponse {
	s.Headers = v
	return s
}

func (s *GetPersonalCarbonInfoResponse) SetStatusCode(v int32) *GetPersonalCarbonInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPersonalCarbonInfoResponse) SetBody(v *GetPersonalCarbonInfoResponseBody) *GetPersonalCarbonInfoResponse {
	s.Body = v
	return s
}

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
	ActionId     *string `json:"actionId,omitempty" xml:"actionId,omitempty"`
	ActionTime   *string `json:"actionTime,omitempty" xml:"actionTime,omitempty"`
	ActionType   *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	CarbonAmount *string `json:"carbonAmount,omitempty" xml:"carbonAmount,omitempty"`
	CorpId       *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeptId       *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	Version      *int32  `json:"version,omitempty" xml:"version,omitempty"`
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
	Result  *int32 `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool  `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *WriteAlibabaOrgCarbonResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *WriteAlibabaOrgCarbonResponse) SetStatusCode(v int32) *WriteAlibabaOrgCarbonResponse {
	s.StatusCode = &v
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
	ActionEndTime   *string `json:"actionEndTime,omitempty" xml:"actionEndTime,omitempty"`
	ActionId        *string `json:"actionId,omitempty" xml:"actionId,omitempty"`
	ActionStartTime *string `json:"actionStartTime,omitempty" xml:"actionStartTime,omitempty"`
	ActionType      *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	CarbonAmount    *string `json:"carbonAmount,omitempty" xml:"carbonAmount,omitempty"`
	CorpId          *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeptId          *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	UserId          *string `json:"userId,omitempty" xml:"userId,omitempty"`
	Version         *int32  `json:"version,omitempty" xml:"version,omitempty"`
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
	Result  *int32 `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool  `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *WriteAlibabaUserCarbonResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *WriteAlibabaUserCarbonResponse) SetStatusCode(v int32) *WriteAlibabaUserCarbonResponse {
	s.StatusCode = &v
	return s
}

func (s *WriteAlibabaUserCarbonResponse) SetBody(v *WriteAlibabaUserCarbonResponseBody) *WriteAlibabaUserCarbonResponse {
	s.Body = v
	return s
}

type WriteIsvStateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s WriteIsvStateHeaders) String() string {
	return tea.Prettify(s)
}

func (s WriteIsvStateHeaders) GoString() string {
	return s.String()
}

func (s *WriteIsvStateHeaders) SetCommonHeaders(v map[string]*string) *WriteIsvStateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *WriteIsvStateHeaders) SetXAcsDingtalkAccessToken(v string) *WriteIsvStateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type WriteIsvStateRequest struct {
	IsvName  *string `json:"isvName,omitempty" xml:"isvName,omitempty"`
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
}

func (s WriteIsvStateRequest) String() string {
	return tea.Prettify(s)
}

func (s WriteIsvStateRequest) GoString() string {
	return s.String()
}

func (s *WriteIsvStateRequest) SetIsvName(v string) *WriteIsvStateRequest {
	s.IsvName = &v
	return s
}

func (s *WriteIsvStateRequest) SetStatDate(v string) *WriteIsvStateRequest {
	s.StatDate = &v
	return s
}

type WriteIsvStateResponseBody struct {
	Result *int64 `json:"result,omitempty" xml:"result,omitempty"`
}

func (s WriteIsvStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WriteIsvStateResponseBody) GoString() string {
	return s.String()
}

func (s *WriteIsvStateResponseBody) SetResult(v int64) *WriteIsvStateResponseBody {
	s.Result = &v
	return s
}

type WriteIsvStateResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *WriteIsvStateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s WriteIsvStateResponse) String() string {
	return tea.Prettify(s)
}

func (s WriteIsvStateResponse) GoString() string {
	return s.String()
}

func (s *WriteIsvStateResponse) SetHeaders(v map[string]*string) *WriteIsvStateResponse {
	s.Headers = v
	return s
}

func (s *WriteIsvStateResponse) SetStatusCode(v int32) *WriteIsvStateResponse {
	s.StatusCode = &v
	return s
}

func (s *WriteIsvStateResponse) SetBody(v *WriteIsvStateResponseBody) *WriteIsvStateResponse {
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
	ActionId     *string `json:"actionId,omitempty" xml:"actionId,omitempty"`
	ActionTime   *string `json:"actionTime,omitempty" xml:"actionTime,omitempty"`
	ActionType   *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	CarbonAmount *string `json:"carbonAmount,omitempty" xml:"carbonAmount,omitempty"`
	CorpId       *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeptId       *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	Version      *int32  `json:"version,omitempty" xml:"version,omitempty"`
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
	Result  *int32 `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool  `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *WriteOrgCarbonResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *WriteOrgCarbonResponse) SetStatusCode(v int32) *WriteOrgCarbonResponse {
	s.StatusCode = &v
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
	ActionEndTime   *string `json:"actionEndTime,omitempty" xml:"actionEndTime,omitempty"`
	ActionId        *string `json:"actionId,omitempty" xml:"actionId,omitempty"`
	ActionStartTime *string `json:"actionStartTime,omitempty" xml:"actionStartTime,omitempty"`
	ActionType      *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	CarbonAmount    *string `json:"carbonAmount,omitempty" xml:"carbonAmount,omitempty"`
	CorpId          *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeptId          *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	UserId          *string `json:"userId,omitempty" xml:"userId,omitempty"`
	Version         *int32  `json:"version,omitempty" xml:"version,omitempty"`
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
	Result  *int32 `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool  `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *WriteUserCarbonResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *WriteUserCarbonResponse) SetStatusCode(v int32) *WriteUserCarbonResponse {
	s.StatusCode = &v
	return s
}

func (s *WriteUserCarbonResponse) SetBody(v *WriteUserCarbonResponseBody) *WriteUserCarbonResponse {
	s.Body = v
	return s
}

type WriteUserCarbonEnergyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s WriteUserCarbonEnergyHeaders) String() string {
	return tea.Prettify(s)
}

func (s WriteUserCarbonEnergyHeaders) GoString() string {
	return s.String()
}

func (s *WriteUserCarbonEnergyHeaders) SetCommonHeaders(v map[string]*string) *WriteUserCarbonEnergyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *WriteUserCarbonEnergyHeaders) SetXAcsDingtalkAccessToken(v string) *WriteUserCarbonEnergyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type WriteUserCarbonEnergyRequest struct {
	UserDetailsList []*WriteUserCarbonEnergyRequestUserDetailsList `json:"userDetailsList,omitempty" xml:"userDetailsList,omitempty" type:"Repeated"`
}

func (s WriteUserCarbonEnergyRequest) String() string {
	return tea.Prettify(s)
}

func (s WriteUserCarbonEnergyRequest) GoString() string {
	return s.String()
}

func (s *WriteUserCarbonEnergyRequest) SetUserDetailsList(v []*WriteUserCarbonEnergyRequestUserDetailsList) *WriteUserCarbonEnergyRequest {
	s.UserDetailsList = v
	return s
}

type WriteUserCarbonEnergyRequestUserDetailsList struct {
	ActionEndTime   *string `json:"actionEndTime,omitempty" xml:"actionEndTime,omitempty"`
	ActionId        *string `json:"actionId,omitempty" xml:"actionId,omitempty"`
	ActionStartTime *string `json:"actionStartTime,omitempty" xml:"actionStartTime,omitempty"`
	ActionType      *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	CarbonAmount    *string `json:"carbonAmount,omitempty" xml:"carbonAmount,omitempty"`
	CorpId          *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeptId          *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	UserId          *string `json:"userId,omitempty" xml:"userId,omitempty"`
	Version         *int32  `json:"version,omitempty" xml:"version,omitempty"`
}

func (s WriteUserCarbonEnergyRequestUserDetailsList) String() string {
	return tea.Prettify(s)
}

func (s WriteUserCarbonEnergyRequestUserDetailsList) GoString() string {
	return s.String()
}

func (s *WriteUserCarbonEnergyRequestUserDetailsList) SetActionEndTime(v string) *WriteUserCarbonEnergyRequestUserDetailsList {
	s.ActionEndTime = &v
	return s
}

func (s *WriteUserCarbonEnergyRequestUserDetailsList) SetActionId(v string) *WriteUserCarbonEnergyRequestUserDetailsList {
	s.ActionId = &v
	return s
}

func (s *WriteUserCarbonEnergyRequestUserDetailsList) SetActionStartTime(v string) *WriteUserCarbonEnergyRequestUserDetailsList {
	s.ActionStartTime = &v
	return s
}

func (s *WriteUserCarbonEnergyRequestUserDetailsList) SetActionType(v string) *WriteUserCarbonEnergyRequestUserDetailsList {
	s.ActionType = &v
	return s
}

func (s *WriteUserCarbonEnergyRequestUserDetailsList) SetCarbonAmount(v string) *WriteUserCarbonEnergyRequestUserDetailsList {
	s.CarbonAmount = &v
	return s
}

func (s *WriteUserCarbonEnergyRequestUserDetailsList) SetCorpId(v string) *WriteUserCarbonEnergyRequestUserDetailsList {
	s.CorpId = &v
	return s
}

func (s *WriteUserCarbonEnergyRequestUserDetailsList) SetDeptId(v int64) *WriteUserCarbonEnergyRequestUserDetailsList {
	s.DeptId = &v
	return s
}

func (s *WriteUserCarbonEnergyRequestUserDetailsList) SetUserId(v string) *WriteUserCarbonEnergyRequestUserDetailsList {
	s.UserId = &v
	return s
}

func (s *WriteUserCarbonEnergyRequestUserDetailsList) SetVersion(v int32) *WriteUserCarbonEnergyRequestUserDetailsList {
	s.Version = &v
	return s
}

type WriteUserCarbonEnergyResponseBody struct {
	Result  *int32 `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s WriteUserCarbonEnergyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WriteUserCarbonEnergyResponseBody) GoString() string {
	return s.String()
}

func (s *WriteUserCarbonEnergyResponseBody) SetResult(v int32) *WriteUserCarbonEnergyResponseBody {
	s.Result = &v
	return s
}

func (s *WriteUserCarbonEnergyResponseBody) SetSuccess(v bool) *WriteUserCarbonEnergyResponseBody {
	s.Success = &v
	return s
}

type WriteUserCarbonEnergyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *WriteUserCarbonEnergyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s WriteUserCarbonEnergyResponse) String() string {
	return tea.Prettify(s)
}

func (s WriteUserCarbonEnergyResponse) GoString() string {
	return s.String()
}

func (s *WriteUserCarbonEnergyResponse) SetHeaders(v map[string]*string) *WriteUserCarbonEnergyResponse {
	s.Headers = v
	return s
}

func (s *WriteUserCarbonEnergyResponse) SetStatusCode(v int32) *WriteUserCarbonEnergyResponse {
	s.StatusCode = &v
	return s
}

func (s *WriteUserCarbonEnergyResponse) SetBody(v *WriteUserCarbonEnergyResponseBody) *WriteUserCarbonEnergyResponse {
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

func (client *Client) GetPersonalCarbonInfoWithOptions(request *GetPersonalCarbonInfoRequest, headers *GetPersonalCarbonInfoHeaders, runtime *util.RuntimeOptions) (_result *GetPersonalCarbonInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionType)) {
		query["actionType"] = request.ActionType
	}

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
		Action:      tea.String("GetPersonalCarbonInfo"),
		Version:     tea.String("carbon_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/carbon/personals/infos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPersonalCarbonInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPersonalCarbonInfo(request *GetPersonalCarbonInfoRequest) (_result *GetPersonalCarbonInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPersonalCarbonInfoHeaders{}
	_result = &GetPersonalCarbonInfoResponse{}
	_body, _err := client.GetPersonalCarbonInfoWithOptions(request, headers, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("WriteAlibabaOrgCarbon"),
		Version:     tea.String("carbon_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/carbon/alibabaOrgDetails/write"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &WriteAlibabaOrgCarbonResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
	params := &openapi.Params{
		Action:      tea.String("WriteAlibabaUserCarbon"),
		Version:     tea.String("carbon_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/carbon/alibabaUserDetails/write"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &WriteAlibabaUserCarbonResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) WriteIsvStateWithOptions(request *WriteIsvStateRequest, headers *WriteIsvStateHeaders, runtime *util.RuntimeOptions) (_result *WriteIsvStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsvName)) {
		query["isvName"] = request.IsvName
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
		Action:      tea.String("WriteIsvState"),
		Version:     tea.String("carbon_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/carbon/datas/states/write"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &WriteIsvStateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) WriteIsvState(request *WriteIsvStateRequest) (_result *WriteIsvStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &WriteIsvStateHeaders{}
	_result = &WriteIsvStateResponse{}
	_body, _err := client.WriteIsvStateWithOptions(request, headers, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("WriteOrgCarbon"),
		Version:     tea.String("carbon_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/carbon/orgDetails/write"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &WriteOrgCarbonResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("WriteUserCarbon"),
		Version:     tea.String("carbon_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/carbon/userDetails/write"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &WriteUserCarbonResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) WriteUserCarbonEnergyWithOptions(request *WriteUserCarbonEnergyRequest, headers *WriteUserCarbonEnergyHeaders, runtime *util.RuntimeOptions) (_result *WriteUserCarbonEnergyResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("WriteUserCarbonEnergy"),
		Version:     tea.String("carbon_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/carbon/userDetails/energies/write"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &WriteUserCarbonEnergyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) WriteUserCarbonEnergy(request *WriteUserCarbonEnergyRequest) (_result *WriteUserCarbonEnergyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &WriteUserCarbonEnergyHeaders{}
	_result = &WriteUserCarbonEnergyResponse{}
	_body, _err := client.WriteUserCarbonEnergyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
