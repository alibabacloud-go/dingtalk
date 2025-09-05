// This file is auto-generated, don't edit it. Thanks.
package exclusive_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ConversationCategoryModel struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CategoryId *int64 `json:"categoryId,omitempty" xml:"categoryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 分组
	CategoryName *string                      `json:"categoryName,omitempty" xml:"categoryName,omitempty"`
	Children     []*ConversationCategoryModel `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	LevelNum *int32 `json:"levelNum,omitempty" xml:"levelNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
}

func (s ConversationCategoryModel) String() string {
	return tea.Prettify(s)
}

func (s ConversationCategoryModel) GoString() string {
	return s.String()
}

func (s *ConversationCategoryModel) SetCategoryId(v int64) *ConversationCategoryModel {
	s.CategoryId = &v
	return s
}

func (s *ConversationCategoryModel) SetCategoryName(v string) *ConversationCategoryModel {
	s.CategoryName = &v
	return s
}

func (s *ConversationCategoryModel) SetChildren(v []*ConversationCategoryModel) *ConversationCategoryModel {
	s.Children = v
	return s
}

func (s *ConversationCategoryModel) SetLevelNum(v int32) *ConversationCategoryModel {
	s.LevelNum = &v
	return s
}

func (s *ConversationCategoryModel) SetOrder(v int32) *ConversationCategoryModel {
	s.Order = &v
	return s
}

type AddCustomSignConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddCustomSignConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddCustomSignConfigHeaders) GoString() string {
	return s.String()
}

func (s *AddCustomSignConfigHeaders) SetCommonHeaders(v map[string]*string) *AddCustomSignConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddCustomSignConfigHeaders) SetXAcsDingtalkAccessToken(v string) *AddCustomSignConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddCustomSignConfigRequest struct {
	AllEffect   *bool `json:"allEffect,omitempty" xml:"allEffect,omitempty"`
	CanDownload *bool `json:"canDownload,omitempty" xml:"canDownload,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx协议
	ProtocolName *string   `json:"protocolName,omitempty" xml:"protocolName,omitempty"`
	PushDeptIds  []*string `json:"pushDeptIds,omitempty" xml:"pushDeptIds,omitempty" type:"Repeated"`
	PushStaffIds []*string `json:"pushStaffIds,omitempty" xml:"pushStaffIds,omitempty" type:"Repeated"`
	// This parameter is required.
	SignTermFiles []*AddCustomSignConfigRequestSignTermFiles `json:"signTermFiles,omitempty" xml:"signTermFiles,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// xxx协议说明
	TermMessage    *string   `json:"termMessage,omitempty" xml:"termMessage,omitempty"`
	UnpushDeptIds  []*string `json:"unpushDeptIds,omitempty" xml:"unpushDeptIds,omitempty" type:"Repeated"`
	UnpushStaffIds []*string `json:"unpushStaffIds,omitempty" xml:"unpushStaffIds,omitempty" type:"Repeated"`
}

func (s AddCustomSignConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCustomSignConfigRequest) GoString() string {
	return s.String()
}

func (s *AddCustomSignConfigRequest) SetAllEffect(v bool) *AddCustomSignConfigRequest {
	s.AllEffect = &v
	return s
}

func (s *AddCustomSignConfigRequest) SetCanDownload(v bool) *AddCustomSignConfigRequest {
	s.CanDownload = &v
	return s
}

func (s *AddCustomSignConfigRequest) SetProtocolName(v string) *AddCustomSignConfigRequest {
	s.ProtocolName = &v
	return s
}

func (s *AddCustomSignConfigRequest) SetPushDeptIds(v []*string) *AddCustomSignConfigRequest {
	s.PushDeptIds = v
	return s
}

func (s *AddCustomSignConfigRequest) SetPushStaffIds(v []*string) *AddCustomSignConfigRequest {
	s.PushStaffIds = v
	return s
}

func (s *AddCustomSignConfigRequest) SetSignTermFiles(v []*AddCustomSignConfigRequestSignTermFiles) *AddCustomSignConfigRequest {
	s.SignTermFiles = v
	return s
}

func (s *AddCustomSignConfigRequest) SetTermMessage(v string) *AddCustomSignConfigRequest {
	s.TermMessage = &v
	return s
}

func (s *AddCustomSignConfigRequest) SetUnpushDeptIds(v []*string) *AddCustomSignConfigRequest {
	s.UnpushDeptIds = v
	return s
}

func (s *AddCustomSignConfigRequest) SetUnpushStaffIds(v []*string) *AddCustomSignConfigRequest {
	s.UnpushStaffIds = v
	return s
}

type AddCustomSignConfigRequestSignTermFiles struct {
	// This parameter is required.
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// This parameter is required.
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
}

func (s AddCustomSignConfigRequestSignTermFiles) String() string {
	return tea.Prettify(s)
}

func (s AddCustomSignConfigRequestSignTermFiles) GoString() string {
	return s.String()
}

func (s *AddCustomSignConfigRequestSignTermFiles) SetFileName(v string) *AddCustomSignConfigRequestSignTermFiles {
	s.FileName = &v
	return s
}

func (s *AddCustomSignConfigRequestSignTermFiles) SetMediaId(v string) *AddCustomSignConfigRequestSignTermFiles {
	s.MediaId = &v
	return s
}

type AddCustomSignConfigResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddCustomSignConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCustomSignConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddCustomSignConfigResponseBody) SetSuccess(v bool) *AddCustomSignConfigResponseBody {
	s.Success = &v
	return s
}

type AddCustomSignConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCustomSignConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCustomSignConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCustomSignConfigResponse) GoString() string {
	return s.String()
}

func (s *AddCustomSignConfigResponse) SetHeaders(v map[string]*string) *AddCustomSignConfigResponse {
	s.Headers = v
	return s
}

func (s *AddCustomSignConfigResponse) SetStatusCode(v int32) *AddCustomSignConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCustomSignConfigResponse) SetBody(v *AddCustomSignConfigResponseBody) *AddCustomSignConfigResponse {
	s.Body = v
	return s
}

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
	City         *string `json:"city,omitempty" xml:"city,omitempty"`
	Industry     *string `json:"industry,omitempty" xml:"industry,omitempty"`
	IndustryCode *int32  `json:"industryCode,omitempty" xml:"industryCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 15800000000
	MobileNum *string `json:"mobileNum,omitempty" xml:"mobileNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试组织
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Province *string `json:"province,omitempty" xml:"province,omitempty"`
}

func (s AddOrgRequest) String() string {
	return tea.Prettify(s)
}

func (s AddOrgRequest) GoString() string {
	return s.String()
}

func (s *AddOrgRequest) SetCity(v string) *AddOrgRequest {
	s.City = &v
	return s
}

func (s *AddOrgRequest) SetIndustry(v string) *AddOrgRequest {
	s.Industry = &v
	return s
}

func (s *AddOrgRequest) SetIndustryCode(v int32) *AddOrgRequest {
	s.IndustryCode = &v
	return s
}

func (s *AddOrgRequest) SetMobileNum(v string) *AddOrgRequest {
	s.MobileNum = &v
	return s
}

func (s *AddOrgRequest) SetName(v string) *AddOrgRequest {
	s.Name = &v
	return s
}

func (s *AddOrgRequest) SetProvince(v string) *AddOrgRequest {
	s.Province = &v
	return s
}

type AddOrgResponseBody struct {
	// example:
	//
	// dinge4a454fa5f32aba6f5bf40edxxxxxxxx
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
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddOrgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// sampleKeyId1234
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sampleSecretId1234
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	// This parameter is required.
	Request *ApproveProcessCallbackRequestRequest `json:"request,omitempty" xml:"request,omitempty" type:"Struct"`
	// example:
	//
	// dingxxxxxxxxxxxx
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// agree
	ApproveResult *string `json:"approveResult,omitempty" xml:"approveResult,omitempty"`
	// This parameter is required.
	ApproveType *string   `json:"approveType,omitempty" xml:"approveType,omitempty"`
	Approvers   []*string `json:"approvers,omitempty" xml:"approvers,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1495592259000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// approve_open_group_expansion
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1495592259000
	FinishTime *int64  `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	Params     *string `json:"params,omitempty" xml:"params,omitempty"`
	// This parameter is required.
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	Title             *string `json:"title,omitempty" xml:"title,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApproveProcessCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	BanWordsType *int32 `json:"banWordsType,omitempty" xml:"banWordsType,omitempty"`
	// This parameter is required.
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
	// example:
	//
	// 成功
	Cause *string `json:"cause,omitempty" xml:"cause,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BanOrOpenGroupWordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type BusinessEventUpdateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BusinessEventUpdateHeaders) String() string {
	return tea.Prettify(s)
}

func (s BusinessEventUpdateHeaders) GoString() string {
	return s.String()
}

func (s *BusinessEventUpdateHeaders) SetCommonHeaders(v map[string]*string) *BusinessEventUpdateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BusinessEventUpdateHeaders) SetXAcsDingtalkAccessToken(v string) *BusinessEventUpdateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BusinessEventUpdateRequest struct {
	// This parameter is required.
	BusinessData map[string]interface{} `json:"businessData,omitempty" xml:"businessData,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	EventType *int32 `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cidxxx
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	UpdateByKey *bool `json:"updateByKey,omitempty" xml:"updateByKey,omitempty"`
}

func (s BusinessEventUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s BusinessEventUpdateRequest) GoString() string {
	return s.String()
}

func (s *BusinessEventUpdateRequest) SetBusinessData(v map[string]interface{}) *BusinessEventUpdateRequest {
	s.BusinessData = v
	return s
}

func (s *BusinessEventUpdateRequest) SetEventType(v int32) *BusinessEventUpdateRequest {
	s.EventType = &v
	return s
}

func (s *BusinessEventUpdateRequest) SetOpenConversationId(v string) *BusinessEventUpdateRequest {
	s.OpenConversationId = &v
	return s
}

func (s *BusinessEventUpdateRequest) SetUpdateByKey(v bool) *BusinessEventUpdateRequest {
	s.UpdateByKey = &v
	return s
}

type BusinessEventUpdateResponseBody struct {
	Result  *bool `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BusinessEventUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BusinessEventUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *BusinessEventUpdateResponseBody) SetResult(v bool) *BusinessEventUpdateResponseBody {
	s.Result = &v
	return s
}

func (s *BusinessEventUpdateResponseBody) SetSuccess(v bool) *BusinessEventUpdateResponseBody {
	s.Success = &v
	return s
}

type BusinessEventUpdateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BusinessEventUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BusinessEventUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s BusinessEventUpdateResponse) GoString() string {
	return s.String()
}

func (s *BusinessEventUpdateResponse) SetHeaders(v map[string]*string) *BusinessEventUpdateResponse {
	s.Headers = v
	return s
}

func (s *BusinessEventUpdateResponse) SetStatusCode(v int32) *BusinessEventUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *BusinessEventUpdateResponse) SetBody(v *BusinessEventUpdateResponseBody) *BusinessEventUpdateResponse {
	s.Body = v
	return s
}

type CheckControlHitStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CheckControlHitStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s CheckControlHitStatusHeaders) GoString() string {
	return s.String()
}

func (s *CheckControlHitStatusHeaders) SetCommonHeaders(v map[string]*string) *CheckControlHitStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CheckControlHitStatusHeaders) SetXAcsDingtalkAccessToken(v string) *CheckControlHitStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CheckControlHitStatusRequest struct {
	NeedMissedFunction *bool `json:"needMissedFunction,omitempty" xml:"needMissedFunction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// staffId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CheckControlHitStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckControlHitStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckControlHitStatusRequest) SetNeedMissedFunction(v bool) *CheckControlHitStatusRequest {
	s.NeedMissedFunction = &v
	return s
}

func (s *CheckControlHitStatusRequest) SetUserId(v string) *CheckControlHitStatusRequest {
	s.UserId = &v
	return s
}

type CheckControlHitStatusResponseBody struct {
	Result *CheckControlHitStatusResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// This parameter is required.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CheckControlHitStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckControlHitStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckControlHitStatusResponseBody) SetResult(v *CheckControlHitStatusResponseBodyResult) *CheckControlHitStatusResponseBody {
	s.Result = v
	return s
}

func (s *CheckControlHitStatusResponseBody) SetSuccess(v bool) *CheckControlHitStatusResponseBody {
	s.Success = &v
	return s
}

type CheckControlHitStatusResponseBodyResult struct {
	ControlList   []*string `json:"controlList,omitempty" xml:"controlList,omitempty" type:"Repeated"`
	ControlStatus *int32    `json:"controlStatus,omitempty" xml:"controlStatus,omitempty"`
	Reason        *string   `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s CheckControlHitStatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CheckControlHitStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CheckControlHitStatusResponseBodyResult) SetControlList(v []*string) *CheckControlHitStatusResponseBodyResult {
	s.ControlList = v
	return s
}

func (s *CheckControlHitStatusResponseBodyResult) SetControlStatus(v int32) *CheckControlHitStatusResponseBodyResult {
	s.ControlStatus = &v
	return s
}

func (s *CheckControlHitStatusResponseBodyResult) SetReason(v string) *CheckControlHitStatusResponseBodyResult {
	s.Reason = &v
	return s
}

type CheckControlHitStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckControlHitStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckControlHitStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckControlHitStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckControlHitStatusResponse) SetHeaders(v map[string]*string) *CheckControlHitStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckControlHitStatusResponse) SetStatusCode(v int32) *CheckControlHitStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckControlHitStatusResponse) SetBody(v *CheckControlHitStatusResponseBody) *CheckControlHitStatusResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// 1
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCategoryAndBindingGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CreateDlpTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateDlpTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateDlpTaskHeaders) GoString() string {
	return s.String()
}

func (s *CreateDlpTaskHeaders) SetCommonHeaders(v map[string]*string) *CreateDlpTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateDlpTaskHeaders) SetXAcsDingtalkAccessToken(v string) *CreateDlpTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateDlpTaskRequest struct {
	// This parameter is required.
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// This parameter is required.
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateDlpTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDlpTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDlpTaskRequest) SetDentryId(v string) *CreateDlpTaskRequest {
	s.DentryId = &v
	return s
}

func (s *CreateDlpTaskRequest) SetSpaceId(v string) *CreateDlpTaskRequest {
	s.SpaceId = &v
	return s
}

func (s *CreateDlpTaskRequest) SetUserId(v string) *CreateDlpTaskRequest {
	s.UserId = &v
	return s
}

type CreateDlpTaskResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateDlpTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDlpTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDlpTaskResponseBody) SetRequestId(v string) *CreateDlpTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateDlpTaskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDlpTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDlpTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDlpTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDlpTaskResponse) SetHeaders(v map[string]*string) *CreateDlpTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDlpTaskResponse) SetStatusCode(v int32) *CreateDlpTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDlpTaskResponse) SetBody(v *CreateDlpTaskResponseBody) *CreateDlpTaskResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// 1
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMessageCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	CurrentCategoryList []*string `json:"currentCategoryList,omitempty" xml:"currentCategoryList,omitempty" type:"Repeated"`
	DeptIds             []*int64  `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	// example:
	//
	// test
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
	// example:
	//
	// 1
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// CV11z5d2bdbb2260d1576000b4a9955fa
	Did *string `json:"did,omitempty" xml:"did,omitempty"`
	// example:
	//
	// 88:92:5a:1f:e8:24
	MacAddress *string `json:"macAddress,omitempty" xml:"macAddress,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Mac
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// example:
	//
	// 11-22-33-44-55
	SerialNumber *string `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
	// example:
	//
	// 2
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 设备名称
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 65224157501039784
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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

func (s *CreateTrustedDeviceRequest) SetSerialNumber(v string) *CreateTrustedDeviceRequest {
	s.SerialNumber = &v
	return s
}

func (s *CreateTrustedDeviceRequest) SetStatus(v int32) *CreateTrustedDeviceRequest {
	s.Status = &v
	return s
}

func (s *CreateTrustedDeviceRequest) SetTitle(v string) *CreateTrustedDeviceRequest {
	s.Title = &v
	return s
}

func (s *CreateTrustedDeviceRequest) SetUserId(v string) *CreateTrustedDeviceRequest {
	s.UserId = &v
	return s
}

type CreateTrustedDeviceResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// true
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTrustedDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	DetailList     []*CreateTrustedDeviceBatchRequestDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" type:"Repeated"`
	MacAddressList []*string                                    `json:"macAddressList,omitempty" xml:"macAddressList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// Win
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateTrustedDeviceBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTrustedDeviceBatchRequest) GoString() string {
	return s.String()
}

func (s *CreateTrustedDeviceBatchRequest) SetDetailList(v []*CreateTrustedDeviceBatchRequestDetailList) *CreateTrustedDeviceBatchRequest {
	s.DetailList = v
	return s
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

type CreateTrustedDeviceBatchRequestDetailList struct {
	MacAddress   *string `json:"macAddress,omitempty" xml:"macAddress,omitempty"`
	SerialNumber *string `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
	Title        *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateTrustedDeviceBatchRequestDetailList) String() string {
	return tea.Prettify(s)
}

func (s CreateTrustedDeviceBatchRequestDetailList) GoString() string {
	return s.String()
}

func (s *CreateTrustedDeviceBatchRequestDetailList) SetMacAddress(v string) *CreateTrustedDeviceBatchRequestDetailList {
	s.MacAddress = &v
	return s
}

func (s *CreateTrustedDeviceBatchRequestDetailList) SetSerialNumber(v string) *CreateTrustedDeviceBatchRequestDetailList {
	s.SerialNumber = &v
	return s
}

func (s *CreateTrustedDeviceBatchRequestDetailList) SetTitle(v string) *CreateTrustedDeviceBatchRequestDetailList {
	s.Title = &v
	return s
}

type CreateTrustedDeviceBatchResponseBody struct {
	// example:
	//
	// true
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTrustedDeviceBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CreateVirusScanTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateVirusScanTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateVirusScanTaskHeaders) GoString() string {
	return s.String()
}

func (s *CreateVirusScanTaskHeaders) SetCommonHeaders(v map[string]*string) *CreateVirusScanTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateVirusScanTaskHeaders) SetXAcsDingtalkAccessToken(v string) *CreateVirusScanTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateVirusScanTaskRequest struct {
	DentryId    *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	FileMd5     *string `json:"fileMd5,omitempty" xml:"fileMd5,omitempty"`
	FileName    *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileSize    *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// This parameter is required.
	Source  *int32  `json:"source,omitempty" xml:"source,omitempty"`
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateVirusScanTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVirusScanTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVirusScanTaskRequest) SetDentryId(v string) *CreateVirusScanTaskRequest {
	s.DentryId = &v
	return s
}

func (s *CreateVirusScanTaskRequest) SetDownloadUrl(v string) *CreateVirusScanTaskRequest {
	s.DownloadUrl = &v
	return s
}

func (s *CreateVirusScanTaskRequest) SetFileMd5(v string) *CreateVirusScanTaskRequest {
	s.FileMd5 = &v
	return s
}

func (s *CreateVirusScanTaskRequest) SetFileName(v string) *CreateVirusScanTaskRequest {
	s.FileName = &v
	return s
}

func (s *CreateVirusScanTaskRequest) SetFileSize(v int64) *CreateVirusScanTaskRequest {
	s.FileSize = &v
	return s
}

func (s *CreateVirusScanTaskRequest) SetSource(v int32) *CreateVirusScanTaskRequest {
	s.Source = &v
	return s
}

func (s *CreateVirusScanTaskRequest) SetSpaceId(v string) *CreateVirusScanTaskRequest {
	s.SpaceId = &v
	return s
}

func (s *CreateVirusScanTaskRequest) SetUserId(v string) *CreateVirusScanTaskRequest {
	s.UserId = &v
	return s
}

type CreateVirusScanTaskResponseBody struct {
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateVirusScanTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVirusScanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVirusScanTaskResponseBody) SetTaskId(v string) *CreateVirusScanTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateVirusScanTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVirusScanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVirusScanTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVirusScanTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateVirusScanTaskResponse) SetHeaders(v map[string]*string) *CreateVirusScanTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateVirusScanTaskResponse) SetStatusCode(v int32) *CreateVirusScanTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVirusScanTaskResponse) SetBody(v *CreateVirusScanTaskResponseBody) *CreateVirusScanTaskResponse {
	s.Body = v
	return s
}

type DataSyncHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DataSyncHeaders) String() string {
	return tea.Prettify(s)
}

func (s DataSyncHeaders) GoString() string {
	return s.String()
}

func (s *DataSyncHeaders) SetCommonHeaders(v map[string]*string) *DataSyncHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DataSyncHeaders) SetXAcsDingtalkAccessToken(v string) *DataSyncHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DataSyncRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sql语句
	Sql *string `json:"sql,omitempty" xml:"sql,omitempty"`
}

func (s DataSyncRequest) String() string {
	return tea.Prettify(s)
}

func (s DataSyncRequest) GoString() string {
	return s.String()
}

func (s *DataSyncRequest) SetSql(v string) *DataSyncRequest {
	s.Sql = &v
	return s
}

type DataSyncResponseBody struct {
	DataList []map[string]interface{} `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	RowsAffected *int32 `json:"rowsAffected,omitempty" xml:"rowsAffected,omitempty"`
}

func (s DataSyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DataSyncResponseBody) GoString() string {
	return s.String()
}

func (s *DataSyncResponseBody) SetDataList(v []map[string]interface{}) *DataSyncResponseBody {
	s.DataList = v
	return s
}

func (s *DataSyncResponseBody) SetRowsAffected(v int32) *DataSyncResponseBody {
	s.RowsAffected = &v
	return s
}

type DataSyncResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DataSyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DataSyncResponse) String() string {
	return tea.Prettify(s)
}

func (s DataSyncResponse) GoString() string {
	return s.String()
}

func (s *DataSyncResponse) SetHeaders(v map[string]*string) *DataSyncResponse {
	s.Headers = v
	return s
}

func (s *DataSyncResponse) SetStatusCode(v int32) *DataSyncResponse {
	s.StatusCode = &v
	return s
}

func (s *DataSyncResponse) SetBody(v *DataSyncResponseBody) *DataSyncResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// true
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAcrossCloudStroageConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *bool              `json:"body,omitempty" xml:"body,omitempty"`
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
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	KickOff *bool `json:"kickOff,omitempty" xml:"kickOff,omitempty"`
	// example:
	//
	// 88:92:5a:1f:e8:24
	MacAddress *string `json:"macAddress,omitempty" xml:"macAddress,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0119194439361061403
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeleteTrustedDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrustedDeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteTrustedDeviceRequest) SetId(v int64) *DeleteTrustedDeviceRequest {
	s.Id = &v
	return s
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
	// This parameter is required.
	//
	// example:
	//
	// true
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTrustedDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	AppId  *int64 `json:"appId,omitempty" xml:"appId,omitempty"`
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// This parameter is required.
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
	// This parameter is required.
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DistributePartnerAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type EditSecurityConfigMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EditSecurityConfigMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s EditSecurityConfigMemberHeaders) GoString() string {
	return s.String()
}

func (s *EditSecurityConfigMemberHeaders) SetCommonHeaders(v map[string]*string) *EditSecurityConfigMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EditSecurityConfigMemberHeaders) SetXAcsDingtalkAccessToken(v string) *EditSecurityConfigMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EditSecurityConfigMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ctrl.xxx
	ConfigKey *string `json:"configKey,omitempty" xml:"configKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// add
	OperateType *string `json:"operateType,omitempty" xml:"operateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// staffxxx
	OperateUserId *string `json:"operateUserId,omitempty" xml:"operateUserId,omitempty"`
	// This parameter is required.
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s EditSecurityConfigMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s EditSecurityConfigMemberRequest) GoString() string {
	return s.String()
}

func (s *EditSecurityConfigMemberRequest) SetConfigKey(v string) *EditSecurityConfigMemberRequest {
	s.ConfigKey = &v
	return s
}

func (s *EditSecurityConfigMemberRequest) SetOperateType(v string) *EditSecurityConfigMemberRequest {
	s.OperateType = &v
	return s
}

func (s *EditSecurityConfigMemberRequest) SetOperateUserId(v string) *EditSecurityConfigMemberRequest {
	s.OperateUserId = &v
	return s
}

func (s *EditSecurityConfigMemberRequest) SetUserIds(v []*string) *EditSecurityConfigMemberRequest {
	s.UserIds = v
	return s
}

type EditSecurityConfigMemberResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s EditSecurityConfigMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditSecurityConfigMemberResponseBody) GoString() string {
	return s.String()
}

func (s *EditSecurityConfigMemberResponseBody) SetSuccess(v bool) *EditSecurityConfigMemberResponseBody {
	s.Success = &v
	return s
}

type EditSecurityConfigMemberResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EditSecurityConfigMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditSecurityConfigMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s EditSecurityConfigMemberResponse) GoString() string {
	return s.String()
}

func (s *EditSecurityConfigMemberResponse) SetHeaders(v map[string]*string) *EditSecurityConfigMemberResponse {
	s.Headers = v
	return s
}

func (s *EditSecurityConfigMemberResponse) SetStatusCode(v int32) *EditSecurityConfigMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *EditSecurityConfigMemberResponse) SetBody(v *EditSecurityConfigMemberResponseBody) *EditSecurityConfigMemberResponse {
	s.Body = v
	return s
}

type ExchangeMainAdminHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ExchangeMainAdminHeaders) String() string {
	return tea.Prettify(s)
}

func (s ExchangeMainAdminHeaders) GoString() string {
	return s.String()
}

func (s *ExchangeMainAdminHeaders) SetCommonHeaders(v map[string]*string) *ExchangeMainAdminHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ExchangeMainAdminHeaders) SetXAcsDingtalkAccessToken(v string) *ExchangeMainAdminHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ExchangeMainAdminRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 509999
	NewAdminUserId *string `json:"newAdminUserId,omitempty" xml:"newAdminUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4045678
	OldAdminUserId *string `json:"oldAdminUserId,omitempty" xml:"oldAdminUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dingxxxxxxxxxxxxx
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
}

func (s ExchangeMainAdminRequest) String() string {
	return tea.Prettify(s)
}

func (s ExchangeMainAdminRequest) GoString() string {
	return s.String()
}

func (s *ExchangeMainAdminRequest) SetNewAdminUserId(v string) *ExchangeMainAdminRequest {
	s.NewAdminUserId = &v
	return s
}

func (s *ExchangeMainAdminRequest) SetOldAdminUserId(v string) *ExchangeMainAdminRequest {
	s.OldAdminUserId = &v
	return s
}

func (s *ExchangeMainAdminRequest) SetTargetCorpId(v string) *ExchangeMainAdminRequest {
	s.TargetCorpId = &v
	return s
}

type ExchangeMainAdminResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExchangeMainAdminResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExchangeMainAdminResponseBody) GoString() string {
	return s.String()
}

func (s *ExchangeMainAdminResponseBody) SetSuccess(v bool) *ExchangeMainAdminResponseBody {
	s.Success = &v
	return s
}

type ExchangeMainAdminResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExchangeMainAdminResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExchangeMainAdminResponse) String() string {
	return tea.Prettify(s)
}

func (s ExchangeMainAdminResponse) GoString() string {
	return s.String()
}

func (s *ExchangeMainAdminResponse) SetHeaders(v map[string]*string) *ExchangeMainAdminResponse {
	s.Headers = v
	return s
}

func (s *ExchangeMainAdminResponse) SetStatusCode(v int32) *ExchangeMainAdminResponse {
	s.StatusCode = &v
	return s
}

func (s *ExchangeMainAdminResponse) SetBody(v *ExchangeMainAdminResponseBody) *ExchangeMainAdminResponse {
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
	// example:
	//
	// XX工作台
	DingPortalName *string `json:"dingPortalName,omitempty" xml:"dingPortalName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dingxxxxxxxxxxxx
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TPL_APP-C97B75277B144ED7AEFE7XXXXXXXX6BA
	TemplateAppUuid *string `json:"templateAppUuid,omitempty" xml:"templateAppUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dingxxxxxxxxxxxx
	TemplateCorpId *string `json:"templateCorpId,omitempty" xml:"templateCorpId,omitempty"`
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
	// example:
	//
	// true
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExclusiveCreateDingPortalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// sampleKeyId1234
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sampleSecretId1234
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://oss.aliyuncs.com/bucket-name/
	Oss *string `json:"oss,omitempty" xml:"oss,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dingxxxxxxxxxxxx
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FileStorageActiveStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// sampleKeyId1234
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sampleSecretId1234
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://oss.aliyuncs.com/bucket-name/
	Oss *string `json:"oss,omitempty" xml:"oss,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dingxxxxxxxxxxxx
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
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
	// example:
	//
	// sampleKeyId1234
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	CheckState *int32 `json:"checkState,omitempty" xml:"checkState,omitempty"`
	// example:
	//
	// https://oss-cn-test.aliyuncs.com\
	Oss *string `json:"oss,omitempty" xml:"oss,omitempty"`
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FileStorageCheckConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 2022-04
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-04
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ding77b8cac4e026cc123xxxxxxxxeb6378f
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 2022-04
	StatisticTime *string `json:"statisticTime,omitempty" xml:"statisticTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 14000
	UsedStorage *int64 `json:"usedStorage,omitempty" xml:"usedStorage,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FileStorageGetQuotaDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// ding77b8cac4e026cc123xxxxxxxxeb6378f
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FileStorageGetStorageStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// sampleKeyId1234
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sampleSecretId1234
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dingxxxxxxxxxxxx
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
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
	// example:
	//
	// sampleKeyId1234
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// example:
	//
	// https://oss-cn-test.aliyuncs.com\
	Oss *string `json:"oss,omitempty" xml:"oss,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FileStorageUpdateStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
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
	// example:
	//
	// 200
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
	// This parameter is required.
	//
	// example:
	//
	// https://down-cdn.dingtalk.com/ddmedia/iAEKAqRmaWxlAwYEzh55BdsFzlFx040G2gAhhAGkC1HdIgKqLZPt3bUH2pAeUAPPAAABgRPQ5KgEzTBIBwAIAA.file
	DarkWatermark *string `json:"darkWatermark,omitempty" xml:"darkWatermark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0138350100401024915916
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateDarkWaterMarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
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
	ItemList []*GetAccountTransferListResponseBodyItemList `json:"itemList,omitempty" xml:"itemList,omitempty" type:"Repeated"`
	// example:
	//
	// 1000
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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
	// example:
	//
	// 财务部
	DeptName *int64 `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// example:
	//
	// 小张
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 123***
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountTransferListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 20
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetActiveUserSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 10000
	AppId *int64 `json:"appId,omitempty" xml:"appId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ding24f2f5ccxxxxxxxx
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
	// example:
	//
	// 3300000
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentIdByRelatedAppIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// 1234
	DeptId *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 管理部
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	MemberCount *int64 `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	// This parameter is required.
	PartnerLabelVOLevel1 *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel1 `json:"partnerLabelVOLevel1,omitempty" xml:"partnerLabelVOLevel1,omitempty" type:"Struct"`
	// This parameter is required.
	PartnerLabelVOLevel2 *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel2 `json:"partnerLabelVOLevel2,omitempty" xml:"partnerLabelVOLevel2,omitempty" type:"Struct"`
	// This parameter is required.
	PartnerLabelVOLevel3 *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel3 `json:"partnerLabelVOLevel3,omitempty" xml:"partnerLabelVOLevel3,omitempty" type:"Struct"`
	// This parameter is required.
	PartnerLabelVOLevel4 *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel4 `json:"partnerLabelVOLevel4,omitempty" xml:"partnerLabelVOLevel4,omitempty" type:"Struct"`
	// This parameter is required.
	PartnerLabelVOLevel5 *GetAllLabelableDeptsResponseBodyDataPartnerLabelVOLevel5 `json:"partnerLabelVOLevel5,omitempty" xml:"partnerLabelVOLevel5,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 1234
	PartnerNum *string `json:"partnerNum,omitempty" xml:"partnerNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234
	SuperDeptId *string `json:"superDeptId,omitempty" xml:"superDeptId,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 1234
	LabelId *int64 `json:"labelId,omitempty" xml:"labelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 一级供应商
	LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	LevelNum *int64 `json:"levelNum,omitempty" xml:"levelNum,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 1234
	LabelId *int64 `json:"labelId,omitempty" xml:"labelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 二级供应商
	LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	LevelNum *int64 `json:"levelNum,omitempty" xml:"levelNum,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 1234
	LabelId *int64 `json:"labelId,omitempty" xml:"labelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 三级供应商
	LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	LevelNum *int64 `json:"levelNum,omitempty" xml:"levelNum,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 1234
	LabelId *int64 `json:"labelId,omitempty" xml:"labelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 四级供应商
	LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	LevelNum *int64 `json:"levelNum,omitempty" xml:"levelNum,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 1234
	LabelId *int64 `json:"labelId,omitempty" xml:"labelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 五级供应商
	LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5
	LevelNum *int64 `json:"levelNum,omitempty" xml:"levelNum,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAllLabelableDeptsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 1655709383307
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 1655709383307
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppDispatchInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 20
	CalendarCreateUserCnt *string `json:"calendarCreateUserCnt,omitempty" xml:"calendarCreateUserCnt,omitempty"`
	// example:
	//
	// 20
	RecvCalendarUserCnt1d *string `json:"recvCalendarUserCnt1d,omitempty" xml:"recvCalendarUserCnt1d,omitempty"`
	// example:
	//
	// 20
	UseCalendarUserCnt1d *string `json:"useCalendarUserCnt1d,omitempty" xml:"useCalendarUserCnt1d,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCalenderSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetCidsByBotCodeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCidsByBotCodeHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCidsByBotCodeHeaders) GoString() string {
	return s.String()
}

func (s *GetCidsByBotCodeHeaders) SetCommonHeaders(v map[string]*string) *GetCidsByBotCodeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCidsByBotCodeHeaders) SetXAcsDingtalkAccessToken(v string) *GetCidsByBotCodeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCidsByBotCodeRequest struct {
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RobotCode  *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s GetCidsByBotCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCidsByBotCodeRequest) GoString() string {
	return s.String()
}

func (s *GetCidsByBotCodeRequest) SetPageNumber(v int32) *GetCidsByBotCodeRequest {
	s.PageNumber = &v
	return s
}

func (s *GetCidsByBotCodeRequest) SetPageSize(v int32) *GetCidsByBotCodeRequest {
	s.PageSize = &v
	return s
}

func (s *GetCidsByBotCodeRequest) SetRobotCode(v string) *GetCidsByBotCodeRequest {
	s.RobotCode = &v
	return s
}

type GetCidsByBotCodeResponseBody struct {
	GroupInfos []*GetCidsByBotCodeResponseBodyGroupInfos `json:"groupInfos,omitempty" xml:"groupInfos,omitempty" type:"Repeated"`
	HasMore    *bool                                     `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
}

func (s GetCidsByBotCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCidsByBotCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetCidsByBotCodeResponseBody) SetGroupInfos(v []*GetCidsByBotCodeResponseBodyGroupInfos) *GetCidsByBotCodeResponseBody {
	s.GroupInfos = v
	return s
}

func (s *GetCidsByBotCodeResponseBody) SetHasMore(v bool) *GetCidsByBotCodeResponseBody {
	s.HasMore = &v
	return s
}

type GetCidsByBotCodeResponseBodyGroupInfos struct {
	BotCreator            *string `json:"botCreator,omitempty" xml:"botCreator,omitempty"`
	BotCreatorIsOrgMember *bool   `json:"botCreatorIsOrgMember,omitempty" xml:"botCreatorIsOrgMember,omitempty"`
	OpenConversationId    *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s GetCidsByBotCodeResponseBodyGroupInfos) String() string {
	return tea.Prettify(s)
}

func (s GetCidsByBotCodeResponseBodyGroupInfos) GoString() string {
	return s.String()
}

func (s *GetCidsByBotCodeResponseBodyGroupInfos) SetBotCreator(v string) *GetCidsByBotCodeResponseBodyGroupInfos {
	s.BotCreator = &v
	return s
}

func (s *GetCidsByBotCodeResponseBodyGroupInfos) SetBotCreatorIsOrgMember(v bool) *GetCidsByBotCodeResponseBodyGroupInfos {
	s.BotCreatorIsOrgMember = &v
	return s
}

func (s *GetCidsByBotCodeResponseBodyGroupInfos) SetOpenConversationId(v string) *GetCidsByBotCodeResponseBodyGroupInfos {
	s.OpenConversationId = &v
	return s
}

type GetCidsByBotCodeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCidsByBotCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCidsByBotCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCidsByBotCodeResponse) GoString() string {
	return s.String()
}

func (s *GetCidsByBotCodeResponse) SetHeaders(v map[string]*string) *GetCidsByBotCodeResponse {
	s.Headers = v
	return s
}

func (s *GetCidsByBotCodeResponse) SetStatusCode(v int32) *GetCidsByBotCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCidsByBotCodeResponse) SetBody(v *GetCidsByBotCodeResponseBody) *GetCidsByBotCodeResponse {
	s.Body = v
	return s
}

type GetClassTagHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetClassTagHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetClassTagHeaders) GoString() string {
	return s.String()
}

func (s *GetClassTagHeaders) SetCommonHeaders(v map[string]*string) *GetClassTagHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetClassTagHeaders) SetXAcsDingtalkAccessToken(v string) *GetClassTagHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetClassTagRequest struct {
	// This parameter is required.
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// This parameter is required.
	TagCode *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
}

func (s GetClassTagRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClassTagRequest) GoString() string {
	return s.String()
}

func (s *GetClassTagRequest) SetEntityId(v string) *GetClassTagRequest {
	s.EntityId = &v
	return s
}

func (s *GetClassTagRequest) SetTagCode(v string) *GetClassTagRequest {
	s.TagCode = &v
	return s
}

type GetClassTagResponseBody struct {
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// example:
	//
	// 1
	DataType *int32 `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 描述信息
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	InnerDownload *string `json:"innerDownload,omitempty" xml:"innerDownload,omitempty"`
	// example:
	//
	// 1
	InnerTransfer *string `json:"innerTransfer,omitempty" xml:"innerTransfer,omitempty"`
	// example:
	//
	// 张三
	ModifierName *string `json:"modifierName,omitempty" xml:"modifierName,omitempty"`
	// example:
	//
	// 标签名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	OutOp *string `json:"outOp,omitempty" xml:"outOp,omitempty"`
	// example:
	//
	// 1
	Rank *int32 `json:"rank,omitempty" xml:"rank,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1735023822867
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetClassTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClassTagResponseBody) GoString() string {
	return s.String()
}

func (s *GetClassTagResponseBody) SetCreatorName(v string) *GetClassTagResponseBody {
	s.CreatorName = &v
	return s
}

func (s *GetClassTagResponseBody) SetDataType(v int32) *GetClassTagResponseBody {
	s.DataType = &v
	return s
}

func (s *GetClassTagResponseBody) SetDescription(v string) *GetClassTagResponseBody {
	s.Description = &v
	return s
}

func (s *GetClassTagResponseBody) SetInnerDownload(v string) *GetClassTagResponseBody {
	s.InnerDownload = &v
	return s
}

func (s *GetClassTagResponseBody) SetInnerTransfer(v string) *GetClassTagResponseBody {
	s.InnerTransfer = &v
	return s
}

func (s *GetClassTagResponseBody) SetModifierName(v string) *GetClassTagResponseBody {
	s.ModifierName = &v
	return s
}

func (s *GetClassTagResponseBody) SetName(v string) *GetClassTagResponseBody {
	s.Name = &v
	return s
}

func (s *GetClassTagResponseBody) SetOutOp(v string) *GetClassTagResponseBody {
	s.OutOp = &v
	return s
}

func (s *GetClassTagResponseBody) SetRank(v int32) *GetClassTagResponseBody {
	s.Rank = &v
	return s
}

func (s *GetClassTagResponseBody) SetStatus(v int32) *GetClassTagResponseBody {
	s.Status = &v
	return s
}

func (s *GetClassTagResponseBody) SetUpdateTime(v int64) *GetClassTagResponseBody {
	s.UpdateTime = &v
	return s
}

type GetClassTagResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClassTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClassTagResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClassTagResponse) GoString() string {
	return s.String()
}

func (s *GetClassTagResponse) SetHeaders(v map[string]*string) *GetClassTagResponse {
	s.Headers = v
	return s
}

func (s *GetClassTagResponse) SetStatusCode(v int32) *GetClassTagResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClassTagResponse) SetBody(v *GetClassTagResponseBody) *GetClassTagResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
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
	// This parameter is required.
	Data []*GetCommentListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// This parameter is required.
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCommentListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConfBaseInfoByLogicalIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	AttendDuration *float32 `json:"attendDuration,omitempty" xml:"attendDuration,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	StaffId *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
	// This parameter is required.
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConferenceDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetConversationCategoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetConversationCategoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetConversationCategoryHeaders) GoString() string {
	return s.String()
}

func (s *GetConversationCategoryHeaders) SetCommonHeaders(v map[string]*string) *GetConversationCategoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetConversationCategoryHeaders) SetXAcsDingtalkAccessToken(v string) *GetConversationCategoryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetConversationCategoryResponseBody struct {
	Result  []*ConversationCategoryModel `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success *bool                        `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetConversationCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConversationCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetConversationCategoryResponseBody) SetResult(v []*ConversationCategoryModel) *GetConversationCategoryResponseBody {
	s.Result = v
	return s
}

func (s *GetConversationCategoryResponseBody) SetSuccess(v bool) *GetConversationCategoryResponseBody {
	s.Success = &v
	return s
}

type GetConversationCategoryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConversationCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConversationCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConversationCategoryResponse) GoString() string {
	return s.String()
}

func (s *GetConversationCategoryResponse) SetHeaders(v map[string]*string) *GetConversationCategoryResponse {
	s.Headers = v
	return s
}

func (s *GetConversationCategoryResponse) SetStatusCode(v int32) *GetConversationCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConversationCategoryResponse) SetBody(v *GetConversationCategoryResponseBody) *GetConversationCategoryResponse {
	s.Body = v
	return s
}

type GetConversationDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetConversationDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetConversationDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetConversationDetailHeaders) SetCommonHeaders(v map[string]*string) *GetConversationDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetConversationDetailHeaders) SetXAcsDingtalkAccessToken(v string) *GetConversationDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetConversationDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cid123xxx
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s GetConversationDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConversationDetailRequest) GoString() string {
	return s.String()
}

func (s *GetConversationDetailRequest) SetOpenConversationId(v string) *GetConversationDetailRequest {
	s.OpenConversationId = &v
	return s
}

type GetConversationDetailResponseBody struct {
	Result  *GetConversationDetailResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                    `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetConversationDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConversationDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetConversationDetailResponseBody) SetResult(v *GetConversationDetailResponseBodyResult) *GetConversationDetailResponseBody {
	s.Result = v
	return s
}

func (s *GetConversationDetailResponseBody) SetSuccess(v bool) *GetConversationDetailResponseBody {
	s.Success = &v
	return s
}

type GetConversationDetailResponseBodyResult struct {
	// example:
	//
	// -1
	CategoryId *int64 `json:"categoryId,omitempty" xml:"categoryId,omitempty"`
	// example:
	//
	// categoryName
	CategoryName *string `json:"categoryName,omitempty" xml:"categoryName,omitempty"`
	GroupCode    *string `json:"groupCode,omitempty" xml:"groupCode,omitempty"`
	// example:
	//
	// 5
	GroupMembersCnt *int32 `json:"groupMembersCnt,omitempty" xml:"groupMembersCnt,omitempty"`
	// example:
	//
	// 5
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// example:
	//
	// groupOwnerName
	GroupOwnerName *string `json:"groupOwnerName,omitempty" xml:"groupOwnerName,omitempty"`
	// example:
	//
	// groupOwnerUserId
	GroupOwnerUserId *string `json:"groupOwnerUserId,omitempty" xml:"groupOwnerUserId,omitempty"`
	IsKpConversation *bool   `json:"isKpConversation,omitempty" xml:"isKpConversation,omitempty"`
	// example:
	//
	// 1
	ManageSign           *int32                                                         `json:"manageSign,omitempty" xml:"manageSign,omitempty"`
	MultipleCategoryList []*GetConversationDetailResponseBodyResultMultipleCategoryList `json:"multipleCategoryList,omitempty" xml:"multipleCategoryList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cidxxx
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// 1
	Order  *int32 `json:"order,omitempty" xml:"order,omitempty"`
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetConversationDetailResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetConversationDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetConversationDetailResponseBodyResult) SetCategoryId(v int64) *GetConversationDetailResponseBodyResult {
	s.CategoryId = &v
	return s
}

func (s *GetConversationDetailResponseBodyResult) SetCategoryName(v string) *GetConversationDetailResponseBodyResult {
	s.CategoryName = &v
	return s
}

func (s *GetConversationDetailResponseBodyResult) SetGroupCode(v string) *GetConversationDetailResponseBodyResult {
	s.GroupCode = &v
	return s
}

func (s *GetConversationDetailResponseBodyResult) SetGroupMembersCnt(v int32) *GetConversationDetailResponseBodyResult {
	s.GroupMembersCnt = &v
	return s
}

func (s *GetConversationDetailResponseBodyResult) SetGroupName(v string) *GetConversationDetailResponseBodyResult {
	s.GroupName = &v
	return s
}

func (s *GetConversationDetailResponseBodyResult) SetGroupOwnerName(v string) *GetConversationDetailResponseBodyResult {
	s.GroupOwnerName = &v
	return s
}

func (s *GetConversationDetailResponseBodyResult) SetGroupOwnerUserId(v string) *GetConversationDetailResponseBodyResult {
	s.GroupOwnerUserId = &v
	return s
}

func (s *GetConversationDetailResponseBodyResult) SetIsKpConversation(v bool) *GetConversationDetailResponseBodyResult {
	s.IsKpConversation = &v
	return s
}

func (s *GetConversationDetailResponseBodyResult) SetManageSign(v int32) *GetConversationDetailResponseBodyResult {
	s.ManageSign = &v
	return s
}

func (s *GetConversationDetailResponseBodyResult) SetMultipleCategoryList(v []*GetConversationDetailResponseBodyResultMultipleCategoryList) *GetConversationDetailResponseBodyResult {
	s.MultipleCategoryList = v
	return s
}

func (s *GetConversationDetailResponseBodyResult) SetOpenConversationId(v string) *GetConversationDetailResponseBodyResult {
	s.OpenConversationId = &v
	return s
}

func (s *GetConversationDetailResponseBodyResult) SetOrder(v int32) *GetConversationDetailResponseBodyResult {
	s.Order = &v
	return s
}

func (s *GetConversationDetailResponseBodyResult) SetStatus(v int32) *GetConversationDetailResponseBodyResult {
	s.Status = &v
	return s
}

type GetConversationDetailResponseBodyResultMultipleCategoryList struct {
	// example:
	//
	// 0
	CategoryId *int64 `json:"categoryId,omitempty" xml:"categoryId,omitempty"`
	// example:
	//
	// 全部
	CategoryName *string `json:"categoryName,omitempty" xml:"categoryName,omitempty"`
	Order        *int32  `json:"order,omitempty" xml:"order,omitempty"`
}

func (s GetConversationDetailResponseBodyResultMultipleCategoryList) String() string {
	return tea.Prettify(s)
}

func (s GetConversationDetailResponseBodyResultMultipleCategoryList) GoString() string {
	return s.String()
}

func (s *GetConversationDetailResponseBodyResultMultipleCategoryList) SetCategoryId(v int64) *GetConversationDetailResponseBodyResultMultipleCategoryList {
	s.CategoryId = &v
	return s
}

func (s *GetConversationDetailResponseBodyResultMultipleCategoryList) SetCategoryName(v string) *GetConversationDetailResponseBodyResultMultipleCategoryList {
	s.CategoryName = &v
	return s
}

func (s *GetConversationDetailResponseBodyResultMultipleCategoryList) SetOrder(v int32) *GetConversationDetailResponseBodyResultMultipleCategoryList {
	s.Order = &v
	return s
}

type GetConversationDetailResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConversationDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConversationDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConversationDetailResponse) GoString() string {
	return s.String()
}

func (s *GetConversationDetailResponse) SetHeaders(v map[string]*string) *GetConversationDetailResponse {
	s.Headers = v
	return s
}

func (s *GetConversationDetailResponse) SetStatusCode(v int32) *GetConversationDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConversationDetailResponse) SetBody(v *GetConversationDetailResponseBody) *GetConversationDetailResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	Data []*GetDingReportDeptSummaryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// false
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 2
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	// example:
	//
	// 123
	DeptId *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// example:
	//
	// 部门A
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// example:
	//
	// 100
	DingReportSendCnt *string `json:"dingReportSendCnt,omitempty" xml:"dingReportSendCnt,omitempty"`
	// example:
	//
	// 100
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDingReportDeptSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 100
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDingReportSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	Data []*GetDocCreatedDeptSummaryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// false
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 2
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	// example:
	//
	// 100
	CreateDocUserCnt1d *string `json:"createDocUserCnt1d,omitempty" xml:"createDocUserCnt1d,omitempty"`
	// example:
	//
	// 123
	DeptId *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// example:
	//
	// 部门A
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// example:
	//
	// 100
	DocCreatedCnt *string `json:"docCreatedCnt,omitempty" xml:"docCreatedCnt,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocCreatedDeptSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 100
	DocCreateUserCnt1d *string `json:"docCreateUserCnt1d,omitempty" xml:"docCreateUserCnt1d,omitempty"`
	// example:
	//
	// 100
	DocCreatedCnt *string `json:"docCreatedCnt,omitempty" xml:"docCreatedCnt,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocCreatedSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// TBXK65QHiiorHvxxxxxxxxP6giEiE
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
	// example:
	//
	// ding32xxxxxxxxe0105d
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// true
	IsMainOrg *bool `json:"isMainOrg,omitempty" xml:"isMainOrg,omitempty"`
	// example:
	//
	// http://xxx.dingtalk.com/xxx.jpg
	LogoUrl *string `json:"logoUrl,omitempty" xml:"logoUrl,omitempty"`
	// example:
	//
	// 钉钉(中国)信息技术有限公司
	OrgFullName *string `json:"orgFullName,omitempty" xml:"orgFullName,omitempty"`
	// example:
	//
	// 钉钉
	OrgName *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExclusiveAccountAllOrgListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	Data []*GetGeneralFormCreatedDeptSummaryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// false
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 2
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	// example:
	//
	// 123
	DeptId *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// example:
	//
	// 部门A
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// example:
	//
	// 100
	GeneralFormCreateCnt1d *string `json:"generalFormCreateCnt1d,omitempty" xml:"generalFormCreateCnt1d,omitempty"`
	// example:
	//
	// 100
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
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGeneralFormCreatedDeptSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 100
	GeneralFormCreatedCnt *string `json:"generalFormCreatedCnt,omitempty" xml:"generalFormCreatedCnt,omitempty"`
	// example:
	//
	// 100
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGeneralFormCreatedSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// cidV3xxxrSuxxxxxxnB8o8gJw==
	DingGroupId *string `json:"dingGroupId,omitempty" xml:"dingGroupId,omitempty"`
	// example:
	//
	// 1
	GroupType *int64 `json:"groupType,omitempty" xml:"groupType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20200305
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
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
	// This parameter is required.
	Data []*GetGroupActiveInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// This parameter is required.
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGroupActiveInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetGroupInfoByCidHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetGroupInfoByCidHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetGroupInfoByCidHeaders) GoString() string {
	return s.String()
}

func (s *GetGroupInfoByCidHeaders) SetCommonHeaders(v map[string]*string) *GetGroupInfoByCidHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetGroupInfoByCidHeaders) SetXAcsDingtalkAccessToken(v string) *GetGroupInfoByCidHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetGroupInfoByCidRequest struct {
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s GetGroupInfoByCidRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGroupInfoByCidRequest) GoString() string {
	return s.String()
}

func (s *GetGroupInfoByCidRequest) SetOpenConversationId(v string) *GetGroupInfoByCidRequest {
	s.OpenConversationId = &v
	return s
}

type GetGroupInfoByCidResponseBody struct {
	GroupInfo *GetGroupInfoByCidResponseBodyGroupInfo `json:"groupInfo,omitempty" xml:"groupInfo,omitempty" type:"Struct"`
}

func (s GetGroupInfoByCidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGroupInfoByCidResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupInfoByCidResponseBody) SetGroupInfo(v *GetGroupInfoByCidResponseBodyGroupInfo) *GetGroupInfoByCidResponseBody {
	s.GroupInfo = v
	return s
}

type GetGroupInfoByCidResponseBodyGroupInfo struct {
	AllOrgMember       *bool   `json:"allOrgMember,omitempty" xml:"allOrgMember,omitempty"`
	GroupName          *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	GroupNumber        *int64  `json:"groupNumber,omitempty" xml:"groupNumber,omitempty"`
	GroupOrganization  *string `json:"groupOrganization,omitempty" xml:"groupOrganization,omitempty"`
	JoinGroupUrl       *string `json:"joinGroupUrl,omitempty" xml:"joinGroupUrl,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s GetGroupInfoByCidResponseBodyGroupInfo) String() string {
	return tea.Prettify(s)
}

func (s GetGroupInfoByCidResponseBodyGroupInfo) GoString() string {
	return s.String()
}

func (s *GetGroupInfoByCidResponseBodyGroupInfo) SetAllOrgMember(v bool) *GetGroupInfoByCidResponseBodyGroupInfo {
	s.AllOrgMember = &v
	return s
}

func (s *GetGroupInfoByCidResponseBodyGroupInfo) SetGroupName(v string) *GetGroupInfoByCidResponseBodyGroupInfo {
	s.GroupName = &v
	return s
}

func (s *GetGroupInfoByCidResponseBodyGroupInfo) SetGroupNumber(v int64) *GetGroupInfoByCidResponseBodyGroupInfo {
	s.GroupNumber = &v
	return s
}

func (s *GetGroupInfoByCidResponseBodyGroupInfo) SetGroupOrganization(v string) *GetGroupInfoByCidResponseBodyGroupInfo {
	s.GroupOrganization = &v
	return s
}

func (s *GetGroupInfoByCidResponseBodyGroupInfo) SetJoinGroupUrl(v string) *GetGroupInfoByCidResponseBodyGroupInfo {
	s.JoinGroupUrl = &v
	return s
}

func (s *GetGroupInfoByCidResponseBodyGroupInfo) SetOpenConversationId(v string) *GetGroupInfoByCidResponseBodyGroupInfo {
	s.OpenConversationId = &v
	return s
}

type GetGroupInfoByCidResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGroupInfoByCidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGroupInfoByCidResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGroupInfoByCidResponse) GoString() string {
	return s.String()
}

func (s *GetGroupInfoByCidResponse) SetHeaders(v map[string]*string) *GetGroupInfoByCidResponse {
	s.Headers = v
	return s
}

func (s *GetGroupInfoByCidResponse) SetStatusCode(v int32) *GetGroupInfoByCidResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGroupInfoByCidResponse) SetBody(v *GetGroupInfoByCidResponseBody) *GetGroupInfoByCidResponse {
	s.Body = v
	return s
}

type GetGroupOrgByCidHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetGroupOrgByCidHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetGroupOrgByCidHeaders) GoString() string {
	return s.String()
}

func (s *GetGroupOrgByCidHeaders) SetCommonHeaders(v map[string]*string) *GetGroupOrgByCidHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetGroupOrgByCidHeaders) SetXAcsDingtalkAccessToken(v string) *GetGroupOrgByCidHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetGroupOrgByCidRequest struct {
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s GetGroupOrgByCidRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGroupOrgByCidRequest) GoString() string {
	return s.String()
}

func (s *GetGroupOrgByCidRequest) SetOpenConversationId(v string) *GetGroupOrgByCidRequest {
	s.OpenConversationId = &v
	return s
}

type GetGroupOrgByCidResponseBody struct {
	GroupOrganization *string `json:"groupOrganization,omitempty" xml:"groupOrganization,omitempty"`
}

func (s GetGroupOrgByCidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGroupOrgByCidResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupOrgByCidResponseBody) SetGroupOrganization(v string) *GetGroupOrgByCidResponseBody {
	s.GroupOrganization = &v
	return s
}

type GetGroupOrgByCidResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGroupOrgByCidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGroupOrgByCidResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGroupOrgByCidResponse) GoString() string {
	return s.String()
}

func (s *GetGroupOrgByCidResponse) SetHeaders(v map[string]*string) *GetGroupOrgByCidResponse {
	s.Headers = v
	return s
}

func (s *GetGroupOrgByCidResponse) SetStatusCode(v int32) *GetGroupOrgByCidResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGroupOrgByCidResponse) SetBody(v *GetGroupOrgByCidResponseBody) *GetGroupOrgByCidResponse {
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
	// if can be null:
	// true
	DeptIds []*string `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	// This parameter is required.
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
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
	// This parameter is required.
	KpiCaliber *string `json:"kpiCaliber,omitempty" xml:"kpiCaliber,omitempty"`
	// This parameter is required.
	KpiId *string `json:"kpiId,omitempty" xml:"kpiId,omitempty"`
	// This parameter is required.
	KpiName *string `json:"kpiName,omitempty" xml:"kpiName,omitempty"`
	// This parameter is required.
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInActiveUserListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// ding77b8cac4e026cc123xxxxxxxxeb6378f
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
	// example:
	//
	// 营业执照照片不清晰
	AuthRemark *string `json:"authRemark,omitempty" xml:"authRemark,omitempty"`
	// example:
	//
	// 2
	AuthStatus *int32 `json:"authStatus,omitempty" xml:"authStatus,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLastOrgAuthDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetMsgConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetMsgConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetMsgConfigHeaders) GoString() string {
	return s.String()
}

func (s *GetMsgConfigHeaders) SetCommonHeaders(v map[string]*string) *GetMsgConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMsgConfigHeaders) SetXAcsDingtalkAccessToken(v string) *GetMsgConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetMsgConfigRequest struct {
	GroupTopic       *string                               `json:"groupTopic,omitempty" xml:"groupTopic,omitempty"`
	GroupType        *string                               `json:"groupType,omitempty" xml:"groupType,omitempty"`
	ListDynamicAttr  []*GetMsgConfigRequestListDynamicAttr `json:"listDynamicAttr,omitempty" xml:"listDynamicAttr,omitempty" type:"Repeated"`
	ListEmployeeCode []*string                             `json:"listEmployeeCode,omitempty" xml:"listEmployeeCode,omitempty" type:"Repeated"`
	ListUnitId       []*int64                              `json:"listUnitId,omitempty" xml:"listUnitId,omitempty" type:"Repeated"`
	OwnerJobNo       *string                               `json:"ownerJobNo,omitempty" xml:"ownerJobNo,omitempty"`
	RuleBusinessCode *string                               `json:"ruleBusinessCode,omitempty" xml:"ruleBusinessCode,omitempty"`
	RuleCategory     *int32                                `json:"ruleCategory,omitempty" xml:"ruleCategory,omitempty"`
	RuleCode         *string                               `json:"ruleCode,omitempty" xml:"ruleCode,omitempty"`
	// This parameter is required.
	SecretKey *string `json:"secretKey,omitempty" xml:"secretKey,omitempty"`
	SysCode   *string `json:"sysCode,omitempty" xml:"sysCode,omitempty"`
}

func (s GetMsgConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMsgConfigRequest) GoString() string {
	return s.String()
}

func (s *GetMsgConfigRequest) SetGroupTopic(v string) *GetMsgConfigRequest {
	s.GroupTopic = &v
	return s
}

func (s *GetMsgConfigRequest) SetGroupType(v string) *GetMsgConfigRequest {
	s.GroupType = &v
	return s
}

func (s *GetMsgConfigRequest) SetListDynamicAttr(v []*GetMsgConfigRequestListDynamicAttr) *GetMsgConfigRequest {
	s.ListDynamicAttr = v
	return s
}

func (s *GetMsgConfigRequest) SetListEmployeeCode(v []*string) *GetMsgConfigRequest {
	s.ListEmployeeCode = v
	return s
}

func (s *GetMsgConfigRequest) SetListUnitId(v []*int64) *GetMsgConfigRequest {
	s.ListUnitId = v
	return s
}

func (s *GetMsgConfigRequest) SetOwnerJobNo(v string) *GetMsgConfigRequest {
	s.OwnerJobNo = &v
	return s
}

func (s *GetMsgConfigRequest) SetRuleBusinessCode(v string) *GetMsgConfigRequest {
	s.RuleBusinessCode = &v
	return s
}

func (s *GetMsgConfigRequest) SetRuleCategory(v int32) *GetMsgConfigRequest {
	s.RuleCategory = &v
	return s
}

func (s *GetMsgConfigRequest) SetRuleCode(v string) *GetMsgConfigRequest {
	s.RuleCode = &v
	return s
}

func (s *GetMsgConfigRequest) SetSecretKey(v string) *GetMsgConfigRequest {
	s.SecretKey = &v
	return s
}

func (s *GetMsgConfigRequest) SetSysCode(v string) *GetMsgConfigRequest {
	s.SysCode = &v
	return s
}

type GetMsgConfigRequestListDynamicAttr struct {
	AttrCode            *string   `json:"attrCode,omitempty" xml:"attrCode,omitempty"`
	ListAttrOptionsCode []*string `json:"listAttrOptionsCode,omitempty" xml:"listAttrOptionsCode,omitempty" type:"Repeated"`
}

func (s GetMsgConfigRequestListDynamicAttr) String() string {
	return tea.Prettify(s)
}

func (s GetMsgConfigRequestListDynamicAttr) GoString() string {
	return s.String()
}

func (s *GetMsgConfigRequestListDynamicAttr) SetAttrCode(v string) *GetMsgConfigRequestListDynamicAttr {
	s.AttrCode = &v
	return s
}

func (s *GetMsgConfigRequestListDynamicAttr) SetListAttrOptionsCode(v []*string) *GetMsgConfigRequestListDynamicAttr {
	s.ListAttrOptionsCode = v
	return s
}

type GetMsgConfigResponseBody struct {
	Code    *int32                        `json:"code,omitempty" xml:"code,omitempty"`
	Data    *GetMsgConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                       `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetMsgConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMsgConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetMsgConfigResponseBody) SetCode(v int32) *GetMsgConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetMsgConfigResponseBody) SetData(v *GetMsgConfigResponseBodyData) *GetMsgConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetMsgConfigResponseBody) SetMessage(v string) *GetMsgConfigResponseBody {
	s.Message = &v
	return s
}

type GetMsgConfigResponseBodyData struct {
	GroupAttributes    []*GetMsgConfigResponseBodyDataGroupAttributes    `json:"groupAttributes,omitempty" xml:"groupAttributes,omitempty" type:"Repeated"`
	MsgConfigs         *GetMsgConfigResponseBodyDataMsgConfigs           `json:"msgConfigs,omitempty" xml:"msgConfigs,omitempty" type:"Struct"`
	ReceiverAttributes []*GetMsgConfigResponseBodyDataReceiverAttributes `json:"receiverAttributes,omitempty" xml:"receiverAttributes,omitempty" type:"Repeated"`
	UnitAttributes     []*GetMsgConfigResponseBodyDataUnitAttributes     `json:"unitAttributes,omitempty" xml:"unitAttributes,omitempty" type:"Repeated"`
}

func (s GetMsgConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMsgConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMsgConfigResponseBodyData) SetGroupAttributes(v []*GetMsgConfigResponseBodyDataGroupAttributes) *GetMsgConfigResponseBodyData {
	s.GroupAttributes = v
	return s
}

func (s *GetMsgConfigResponseBodyData) SetMsgConfigs(v *GetMsgConfigResponseBodyDataMsgConfigs) *GetMsgConfigResponseBodyData {
	s.MsgConfigs = v
	return s
}

func (s *GetMsgConfigResponseBodyData) SetReceiverAttributes(v []*GetMsgConfigResponseBodyDataReceiverAttributes) *GetMsgConfigResponseBodyData {
	s.ReceiverAttributes = v
	return s
}

func (s *GetMsgConfigResponseBodyData) SetUnitAttributes(v []*GetMsgConfigResponseBodyDataUnitAttributes) *GetMsgConfigResponseBodyData {
	s.UnitAttributes = v
	return s
}

type GetMsgConfigResponseBodyDataGroupAttributes struct {
	ConfigGroupId      *int64                                                        `json:"configGroupId,omitempty" xml:"configGroupId,omitempty"`
	CorpId             *string                                                       `json:"corpId,omitempty" xml:"corpId,omitempty"`
	GroupTopic         *string                                                       `json:"groupTopic,omitempty" xml:"groupTopic,omitempty"`
	GroupType          *string                                                       `json:"groupType,omitempty" xml:"groupType,omitempty"`
	ListDynamicAttr    []*GetMsgConfigResponseBodyDataGroupAttributesListDynamicAttr `json:"listDynamicAttr,omitempty" xml:"listDynamicAttr,omitempty" type:"Repeated"`
	ListReceiver       []*GetMsgConfigResponseBodyDataGroupAttributesListReceiver    `json:"listReceiver,omitempty" xml:"listReceiver,omitempty" type:"Repeated"`
	OpenConversationId *string                                                       `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	OwnerJobNo         *string                                                       `json:"ownerJobNo,omitempty" xml:"ownerJobNo,omitempty"`
	SubRuleCode        *string                                                       `json:"subRuleCode,omitempty" xml:"subRuleCode,omitempty"`
}

func (s GetMsgConfigResponseBodyDataGroupAttributes) String() string {
	return tea.Prettify(s)
}

func (s GetMsgConfigResponseBodyDataGroupAttributes) GoString() string {
	return s.String()
}

func (s *GetMsgConfigResponseBodyDataGroupAttributes) SetConfigGroupId(v int64) *GetMsgConfigResponseBodyDataGroupAttributes {
	s.ConfigGroupId = &v
	return s
}

func (s *GetMsgConfigResponseBodyDataGroupAttributes) SetCorpId(v string) *GetMsgConfigResponseBodyDataGroupAttributes {
	s.CorpId = &v
	return s
}

func (s *GetMsgConfigResponseBodyDataGroupAttributes) SetGroupTopic(v string) *GetMsgConfigResponseBodyDataGroupAttributes {
	s.GroupTopic = &v
	return s
}

func (s *GetMsgConfigResponseBodyDataGroupAttributes) SetGroupType(v string) *GetMsgConfigResponseBodyDataGroupAttributes {
	s.GroupType = &v
	return s
}

func (s *GetMsgConfigResponseBodyDataGroupAttributes) SetListDynamicAttr(v []*GetMsgConfigResponseBodyDataGroupAttributesListDynamicAttr) *GetMsgConfigResponseBodyDataGroupAttributes {
	s.ListDynamicAttr = v
	return s
}

func (s *GetMsgConfigResponseBodyDataGroupAttributes) SetListReceiver(v []*GetMsgConfigResponseBodyDataGroupAttributesListReceiver) *GetMsgConfigResponseBodyDataGroupAttributes {
	s.ListReceiver = v
	return s
}

func (s *GetMsgConfigResponseBodyDataGroupAttributes) SetOpenConversationId(v string) *GetMsgConfigResponseBodyDataGroupAttributes {
	s.OpenConversationId = &v
	return s
}

func (s *GetMsgConfigResponseBodyDataGroupAttributes) SetOwnerJobNo(v string) *GetMsgConfigResponseBodyDataGroupAttributes {
	s.OwnerJobNo = &v
	return s
}

func (s *GetMsgConfigResponseBodyDataGroupAttributes) SetSubRuleCode(v string) *GetMsgConfigResponseBodyDataGroupAttributes {
	s.SubRuleCode = &v
	return s
}

type GetMsgConfigResponseBodyDataGroupAttributesListDynamicAttr struct {
	AttrCode            *string   `json:"attrCode,omitempty" xml:"attrCode,omitempty"`
	ListAttrOptionsCode []*string `json:"listAttrOptionsCode,omitempty" xml:"listAttrOptionsCode,omitempty" type:"Repeated"`
}

func (s GetMsgConfigResponseBodyDataGroupAttributesListDynamicAttr) String() string {
	return tea.Prettify(s)
}

func (s GetMsgConfigResponseBodyDataGroupAttributesListDynamicAttr) GoString() string {
	return s.String()
}

func (s *GetMsgConfigResponseBodyDataGroupAttributesListDynamicAttr) SetAttrCode(v string) *GetMsgConfigResponseBodyDataGroupAttributesListDynamicAttr {
	s.AttrCode = &v
	return s
}

func (s *GetMsgConfigResponseBodyDataGroupAttributesListDynamicAttr) SetListAttrOptionsCode(v []*string) *GetMsgConfigResponseBodyDataGroupAttributesListDynamicAttr {
	s.ListAttrOptionsCode = v
	return s
}

type GetMsgConfigResponseBodyDataGroupAttributesListReceiver struct {
	EmployeeCode *string `json:"employeeCode,omitempty" xml:"employeeCode,omitempty"`
	EmployeeName *string `json:"employeeName,omitempty" xml:"employeeName,omitempty"`
}

func (s GetMsgConfigResponseBodyDataGroupAttributesListReceiver) String() string {
	return tea.Prettify(s)
}

func (s GetMsgConfigResponseBodyDataGroupAttributesListReceiver) GoString() string {
	return s.String()
}

func (s *GetMsgConfigResponseBodyDataGroupAttributesListReceiver) SetEmployeeCode(v string) *GetMsgConfigResponseBodyDataGroupAttributesListReceiver {
	s.EmployeeCode = &v
	return s
}

func (s *GetMsgConfigResponseBodyDataGroupAttributesListReceiver) SetEmployeeName(v string) *GetMsgConfigResponseBodyDataGroupAttributesListReceiver {
	s.EmployeeName = &v
	return s
}

type GetMsgConfigResponseBodyDataMsgConfigs struct {
	CardId               *string `json:"cardId,omitempty" xml:"cardId,omitempty"`
	CorpId               *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	CustomParameters     *string `json:"customParameters,omitempty" xml:"customParameters,omitempty"`
	MsgContentConsisFlag *int32  `json:"msgContentConsisFlag,omitempty" xml:"msgContentConsisFlag,omitempty"`
	MsgId                *string `json:"msgId,omitempty" xml:"msgId,omitempty"`
	RobotCode            *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	RuleBusinessCode     *string `json:"ruleBusinessCode,omitempty" xml:"ruleBusinessCode,omitempty"`
	RuleCategory         *int32  `json:"ruleCategory,omitempty" xml:"ruleCategory,omitempty"`
	RuleCode             *string `json:"ruleCode,omitempty" xml:"ruleCode,omitempty"`
	RuleName             *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	SubRuleCode          *string `json:"subRuleCode,omitempty" xml:"subRuleCode,omitempty"`
	SystemCode           *string `json:"systemCode,omitempty" xml:"systemCode,omitempty"`
	TaskBatchNo          *string `json:"taskBatchNo,omitempty" xml:"taskBatchNo,omitempty"`
}

func (s GetMsgConfigResponseBodyDataMsgConfigs) String() string {
	return tea.Prettify(s)
}

func (s GetMsgConfigResponseBodyDataMsgConfigs) GoString() string {
	return s.String()
}

func (s *GetMsgConfigResponseBodyDataMsgConfigs) SetCardId(v string) *GetMsgConfigResponseBodyDataMsgConfigs {
	s.CardId = &v
	return s
}

func (s *GetMsgConfigResponseBodyDataMsgConfigs) SetCorpId(v string) *GetMsgConfigResponseBodyDataMsgConfigs {
	s.CorpId = &v
	return s
}

func (s *GetMsgConfigResponseBodyDataMsgConfigs) SetCustomParameters(v string) *GetMsgConfigResponseBodyDataMsgConfigs {
	s.CustomParameters = &v
	return s
}

func (s *GetMsgConfigResponseBodyDataMsgConfigs) SetMsgContentConsisFlag(v int32) *GetMsgConfigResponseBodyDataMsgConfigs {
	s.MsgContentConsisFlag = &v
	return s
}

func (s *GetMsgConfigResponseBodyDataMsgConfigs) SetMsgId(v string) *GetMsgConfigResponseBodyDataMsgConfigs {
	s.MsgId = &v
	return s
}

func (s *GetMsgConfigResponseBodyDataMsgConfigs) SetRobotCode(v string) *GetMsgConfigResponseBodyDataMsgConfigs {
	s.RobotCode = &v
	return s
}

func (s *GetMsgConfigResponseBodyDataMsgConfigs) SetRuleBusinessCode(v string) *GetMsgConfigResponseBodyDataMsgConfigs {
	s.RuleBusinessCode = &v
	return s
}

func (s *GetMsgConfigResponseBodyDataMsgConfigs) SetRuleCategory(v int32) *GetMsgConfigResponseBodyDataMsgConfigs {
	s.RuleCategory = &v
	return s
}

func (s *GetMsgConfigResponseBodyDataMsgConfigs) SetRuleCode(v string) *GetMsgConfigResponseBodyDataMsgConfigs {
	s.RuleCode = &v
	return s
}

func (s *GetMsgConfigResponseBodyDataMsgConfigs) SetRuleName(v string) *GetMsgConfigResponseBodyDataMsgConfigs {
	s.RuleName = &v
	return s
}

func (s *GetMsgConfigResponseBodyDataMsgConfigs) SetSubRuleCode(v string) *GetMsgConfigResponseBodyDataMsgConfigs {
	s.SubRuleCode = &v
	return s
}

func (s *GetMsgConfigResponseBodyDataMsgConfigs) SetSystemCode(v string) *GetMsgConfigResponseBodyDataMsgConfigs {
	s.SystemCode = &v
	return s
}

func (s *GetMsgConfigResponseBodyDataMsgConfigs) SetTaskBatchNo(v string) *GetMsgConfigResponseBodyDataMsgConfigs {
	s.TaskBatchNo = &v
	return s
}

type GetMsgConfigResponseBodyDataReceiverAttributes struct {
	EmployeeCode *string `json:"employeeCode,omitempty" xml:"employeeCode,omitempty"`
}

func (s GetMsgConfigResponseBodyDataReceiverAttributes) String() string {
	return tea.Prettify(s)
}

func (s GetMsgConfigResponseBodyDataReceiverAttributes) GoString() string {
	return s.String()
}

func (s *GetMsgConfigResponseBodyDataReceiverAttributes) SetEmployeeCode(v string) *GetMsgConfigResponseBodyDataReceiverAttributes {
	s.EmployeeCode = &v
	return s
}

type GetMsgConfigResponseBodyDataUnitAttributes struct {
	UnitId *int64 `json:"unitId,omitempty" xml:"unitId,omitempty"`
}

func (s GetMsgConfigResponseBodyDataUnitAttributes) String() string {
	return tea.Prettify(s)
}

func (s GetMsgConfigResponseBodyDataUnitAttributes) GoString() string {
	return s.String()
}

func (s *GetMsgConfigResponseBodyDataUnitAttributes) SetUnitId(v int64) *GetMsgConfigResponseBodyDataUnitAttributes {
	s.UnitId = &v
	return s
}

type GetMsgConfigResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMsgConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMsgConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMsgConfigResponse) GoString() string {
	return s.String()
}

func (s *GetMsgConfigResponse) SetHeaders(v map[string]*string) *GetMsgConfigResponse {
	s.Headers = v
	return s
}

func (s *GetMsgConfigResponse) SetStatusCode(v int32) *GetMsgConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMsgConfigResponse) SetBody(v *GetMsgConfigResponseBody) *GetMsgConfigResponse {
	s.Body = v
	return s
}

type GetMsgLocationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetMsgLocationHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetMsgLocationHeaders) GoString() string {
	return s.String()
}

func (s *GetMsgLocationHeaders) SetCommonHeaders(v map[string]*string) *GetMsgLocationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMsgLocationHeaders) SetXAcsDingtalkAccessToken(v string) *GetMsgLocationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetMsgLocationRequest struct {
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	OpenMsgId *string `json:"openMsgId,omitempty" xml:"openMsgId,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetMsgLocationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMsgLocationRequest) GoString() string {
	return s.String()
}

func (s *GetMsgLocationRequest) SetOpenConversationId(v string) *GetMsgLocationRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GetMsgLocationRequest) SetOpenMsgId(v string) *GetMsgLocationRequest {
	s.OpenMsgId = &v
	return s
}

func (s *GetMsgLocationRequest) SetUserId(v string) *GetMsgLocationRequest {
	s.UserId = &v
	return s
}

type GetMsgLocationResponseBody struct {
	MsgLocationUrl *string `json:"msgLocationUrl,omitempty" xml:"msgLocationUrl,omitempty"`
}

func (s GetMsgLocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMsgLocationResponseBody) GoString() string {
	return s.String()
}

func (s *GetMsgLocationResponseBody) SetMsgLocationUrl(v string) *GetMsgLocationResponseBody {
	s.MsgLocationUrl = &v
	return s
}

type GetMsgLocationResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMsgLocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMsgLocationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMsgLocationResponse) GoString() string {
	return s.String()
}

func (s *GetMsgLocationResponse) SetHeaders(v map[string]*string) *GetMsgLocationResponse {
	s.Headers = v
	return s
}

func (s *GetMsgLocationResponse) SetStatusCode(v int32) *GetMsgLocationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMsgLocationResponse) SetBody(v *GetMsgLocationResponseBody) *GetMsgLocationResponse {
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
	// This parameter is required.
	EndTime  *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
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
	// This parameter is required.
	Data []*GetOaOperatorLogListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// This parameter is required.
	ItemCount *int64 `json:"itemCount,omitempty" xml:"itemCount,omitempty"`
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
	// This parameter is required.
	Category1Name *string `json:"category1Name,omitempty" xml:"category1Name,omitempty"`
	// This parameter is required.
	Category2Name *string `json:"category2Name,omitempty" xml:"category2Name,omitempty"`
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	OpName *string `json:"opName,omitempty" xml:"opName,omitempty"`
	// This parameter is required.
	OpTime *int64 `json:"opTime,omitempty" xml:"opTime,omitempty"`
	// This parameter is required.
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOaOperatorLogListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
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
	// This parameter is required.
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
	// example:
	//
	// 10
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
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
	// example:
	//
	// {   "text": {     "content": "这是一段文本"   } }
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOutGroupsByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1680493837426
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
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
	// This parameter is required.
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
	// example:
	//
	// 1680493837428
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	// example:
	//
	// {   "text": {     "content": "这是一段文本"   } }
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// text/audio/video
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// example:
	//
	// 2022-07-05 15:43:03
	CreateAt *string `json:"createAt,omitempty" xml:"createAt,omitempty"`
	// example:
	//
	// cid123456
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
	// example:
	//
	// user1234
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// userId
	IdType *string `json:"idType,omitempty" xml:"idType,omitempty"`
	// example:
	//
	// user
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOutsideAuditGroupMessageByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// 1234
	LabelId *string `json:"labelId,omitempty" xml:"labelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 目前无意义
	TypeId *float32 `json:"typeId,omitempty" xml:"typeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 经销商
	TypeName *string `json:"typeName,omitempty" xml:"typeName,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPartnerTypeByParentIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetPrivateStoreCapacityUsageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPrivateStoreCapacityUsageHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateStoreCapacityUsageHeaders) GoString() string {
	return s.String()
}

func (s *GetPrivateStoreCapacityUsageHeaders) SetCommonHeaders(v map[string]*string) *GetPrivateStoreCapacityUsageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPrivateStoreCapacityUsageHeaders) SetXAcsDingtalkAccessToken(v string) *GetPrivateStoreCapacityUsageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPrivateStoreCapacityUsageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ding1234
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
}

func (s GetPrivateStoreCapacityUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateStoreCapacityUsageRequest) GoString() string {
	return s.String()
}

func (s *GetPrivateStoreCapacityUsageRequest) SetTargetCorpId(v string) *GetPrivateStoreCapacityUsageRequest {
	s.TargetCorpId = &v
	return s
}

type GetPrivateStoreCapacityUsageResponseBody struct {
	UsedSize *int64 `json:"usedSize,omitempty" xml:"usedSize,omitempty"`
}

func (s GetPrivateStoreCapacityUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateStoreCapacityUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetPrivateStoreCapacityUsageResponseBody) SetUsedSize(v int64) *GetPrivateStoreCapacityUsageResponseBody {
	s.UsedSize = &v
	return s
}

type GetPrivateStoreCapacityUsageResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPrivateStoreCapacityUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPrivateStoreCapacityUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateStoreCapacityUsageResponse) GoString() string {
	return s.String()
}

func (s *GetPrivateStoreCapacityUsageResponse) SetHeaders(v map[string]*string) *GetPrivateStoreCapacityUsageResponse {
	s.Headers = v
	return s
}

func (s *GetPrivateStoreCapacityUsageResponse) SetStatusCode(v int32) *GetPrivateStoreCapacityUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPrivateStoreCapacityUsageResponse) SetBody(v *GetPrivateStoreCapacityUsageResponseBody) *GetPrivateStoreCapacityUsageResponse {
	s.Body = v
	return s
}

type GetPrivateStoreFileInfosByPageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPrivateStoreFileInfosByPageHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateStoreFileInfosByPageHeaders) GoString() string {
	return s.String()
}

func (s *GetPrivateStoreFileInfosByPageHeaders) SetCommonHeaders(v map[string]*string) *GetPrivateStoreFileInfosByPageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPrivateStoreFileInfosByPageHeaders) SetXAcsDingtalkAccessToken(v string) *GetPrivateStoreFileInfosByPageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPrivateStoreFileInfosByPageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1680493837426
	FileCreateTime *int64  `json:"fileCreateTime,omitempty" xml:"fileCreateTime,omitempty"`
	FileStatus     *string `json:"fileStatus,omitempty" xml:"fileStatus,omitempty"`
	MaxResults     *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken      *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	Order *string `json:"order,omitempty" xml:"order,omitempty"`
	// This parameter is required.
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
}

func (s GetPrivateStoreFileInfosByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateStoreFileInfosByPageRequest) GoString() string {
	return s.String()
}

func (s *GetPrivateStoreFileInfosByPageRequest) SetFileCreateTime(v int64) *GetPrivateStoreFileInfosByPageRequest {
	s.FileCreateTime = &v
	return s
}

func (s *GetPrivateStoreFileInfosByPageRequest) SetFileStatus(v string) *GetPrivateStoreFileInfosByPageRequest {
	s.FileStatus = &v
	return s
}

func (s *GetPrivateStoreFileInfosByPageRequest) SetMaxResults(v int32) *GetPrivateStoreFileInfosByPageRequest {
	s.MaxResults = &v
	return s
}

func (s *GetPrivateStoreFileInfosByPageRequest) SetNextToken(v string) *GetPrivateStoreFileInfosByPageRequest {
	s.NextToken = &v
	return s
}

func (s *GetPrivateStoreFileInfosByPageRequest) SetOrder(v string) *GetPrivateStoreFileInfosByPageRequest {
	s.Order = &v
	return s
}

func (s *GetPrivateStoreFileInfosByPageRequest) SetTargetCorpId(v string) *GetPrivateStoreFileInfosByPageRequest {
	s.TargetCorpId = &v
	return s
}

type GetPrivateStoreFileInfosByPageResponseBody struct {
	FileInfos []*GetPrivateStoreFileInfosByPageResponseBodyFileInfos `json:"fileInfos,omitempty" xml:"fileInfos,omitempty" type:"Repeated"`
	NextToken *string                                                `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetPrivateStoreFileInfosByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateStoreFileInfosByPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetPrivateStoreFileInfosByPageResponseBody) SetFileInfos(v []*GetPrivateStoreFileInfosByPageResponseBodyFileInfos) *GetPrivateStoreFileInfosByPageResponseBody {
	s.FileInfos = v
	return s
}

func (s *GetPrivateStoreFileInfosByPageResponseBody) SetNextToken(v string) *GetPrivateStoreFileInfosByPageResponseBody {
	s.NextToken = &v
	return s
}

type GetPrivateStoreFileInfosByPageResponseBodyFileInfos struct {
	DentryId       *int64  `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	FileCreateTime *int64  `json:"fileCreateTime,omitempty" xml:"fileCreateTime,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FilePath       *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	FileSize       *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	SpaceId        *int64  `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	Status         *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetPrivateStoreFileInfosByPageResponseBodyFileInfos) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateStoreFileInfosByPageResponseBodyFileInfos) GoString() string {
	return s.String()
}

func (s *GetPrivateStoreFileInfosByPageResponseBodyFileInfos) SetDentryId(v int64) *GetPrivateStoreFileInfosByPageResponseBodyFileInfos {
	s.DentryId = &v
	return s
}

func (s *GetPrivateStoreFileInfosByPageResponseBodyFileInfos) SetFileCreateTime(v int64) *GetPrivateStoreFileInfosByPageResponseBodyFileInfos {
	s.FileCreateTime = &v
	return s
}

func (s *GetPrivateStoreFileInfosByPageResponseBodyFileInfos) SetFileName(v string) *GetPrivateStoreFileInfosByPageResponseBodyFileInfos {
	s.FileName = &v
	return s
}

func (s *GetPrivateStoreFileInfosByPageResponseBodyFileInfos) SetFilePath(v string) *GetPrivateStoreFileInfosByPageResponseBodyFileInfos {
	s.FilePath = &v
	return s
}

func (s *GetPrivateStoreFileInfosByPageResponseBodyFileInfos) SetFileSize(v int64) *GetPrivateStoreFileInfosByPageResponseBodyFileInfos {
	s.FileSize = &v
	return s
}

func (s *GetPrivateStoreFileInfosByPageResponseBodyFileInfos) SetSpaceId(v int64) *GetPrivateStoreFileInfosByPageResponseBodyFileInfos {
	s.SpaceId = &v
	return s
}

func (s *GetPrivateStoreFileInfosByPageResponseBodyFileInfos) SetStatus(v string) *GetPrivateStoreFileInfosByPageResponseBodyFileInfos {
	s.Status = &v
	return s
}

type GetPrivateStoreFileInfosByPageResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPrivateStoreFileInfosByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPrivateStoreFileInfosByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateStoreFileInfosByPageResponse) GoString() string {
	return s.String()
}

func (s *GetPrivateStoreFileInfosByPageResponse) SetHeaders(v map[string]*string) *GetPrivateStoreFileInfosByPageResponse {
	s.Headers = v
	return s
}

func (s *GetPrivateStoreFileInfosByPageResponse) SetStatusCode(v int32) *GetPrivateStoreFileInfosByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPrivateStoreFileInfosByPageResponse) SetBody(v *GetPrivateStoreFileInfosByPageResponseBody) *GetPrivateStoreFileInfosByPageResponse {
	s.Body = v
	return s
}

type GetPrivateStoreFilePathHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPrivateStoreFilePathHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateStoreFilePathHeaders) GoString() string {
	return s.String()
}

func (s *GetPrivateStoreFilePathHeaders) SetCommonHeaders(v map[string]*string) *GetPrivateStoreFilePathHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPrivateStoreFilePathHeaders) SetXAcsDingtalkAccessToken(v string) *GetPrivateStoreFilePathHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPrivateStoreFilePathRequest struct {
	// This parameter is required.
	DentryId *int64 `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// This parameter is required.
	SpaceId *int64 `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s GetPrivateStoreFilePathRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateStoreFilePathRequest) GoString() string {
	return s.String()
}

func (s *GetPrivateStoreFilePathRequest) SetDentryId(v int64) *GetPrivateStoreFilePathRequest {
	s.DentryId = &v
	return s
}

func (s *GetPrivateStoreFilePathRequest) SetSpaceId(v int64) *GetPrivateStoreFilePathRequest {
	s.SpaceId = &v
	return s
}

type GetPrivateStoreFilePathResponseBody struct {
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
}

func (s GetPrivateStoreFilePathResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateStoreFilePathResponseBody) GoString() string {
	return s.String()
}

func (s *GetPrivateStoreFilePathResponseBody) SetFilePath(v string) *GetPrivateStoreFilePathResponseBody {
	s.FilePath = &v
	return s
}

type GetPrivateStoreFilePathResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPrivateStoreFilePathResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPrivateStoreFilePathResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateStoreFilePathResponse) GoString() string {
	return s.String()
}

func (s *GetPrivateStoreFilePathResponse) SetHeaders(v map[string]*string) *GetPrivateStoreFilePathResponse {
	s.Headers = v
	return s
}

func (s *GetPrivateStoreFilePathResponse) SetStatusCode(v int32) *GetPrivateStoreFilePathResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPrivateStoreFilePathResponse) SetBody(v *GetPrivateStoreFilePathResponseBody) *GetPrivateStoreFilePathResponse {
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
	// example:
	//
	// 1671767361000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 88:66:5a:07:2b:04
	MacAddress *string `json:"macAddress,omitempty" xml:"macAddress,omitempty"`
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
	// Mac
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// example:
	//
	// 11-22-33-44
	SerialNumber *string `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
	// example:
	//
	// 1671767361000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 这是标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
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

func (s *GetPublicDevicesRequest) SetSerialNumber(v string) *GetPublicDevicesRequest {
	s.SerialNumber = &v
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
	Data []*GetPublicDevicesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	DataCnt *int32 `json:"dataCnt,omitempty" xml:"dataCnt,omitempty"`
	// example:
	//
	// 100
	TotalCnt *int64 `json:"totalCnt,omitempty" xml:"totalCnt,omitempty"`
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
	DeviceDepts []*GetPublicDevicesResponseBodyDataDeviceDepts `json:"deviceDepts,omitempty" xml:"deviceDepts,omitempty" type:"Repeated"`
	DeviceRoles []*GetPublicDevicesResponseBodyDataDeviceRoles `json:"deviceRoles,omitempty" xml:"deviceRoles,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	DeviceScopeType *int32                                          `json:"deviceScopeType,omitempty" xml:"deviceScopeType,omitempty"`
	DeviceStaffs    []*GetPublicDevicesResponseBodyDataDeviceStaffs `json:"deviceStaffs,omitempty" xml:"deviceStaffs,omitempty" type:"Repeated"`
	// example:
	//
	// 1671767361000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1671767361000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 88:66:5a:07:2b:04
	MacAddress *string `json:"macAddress,omitempty" xml:"macAddress,omitempty"`
	// example:
	//
	// Mac
	Platform     *string `json:"platform,omitempty" xml:"platform,omitempty"`
	SerialNumber *string `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
	// example:
	//
	// 这是标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
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

func (s *GetPublicDevicesResponseBodyData) SetSerialNumber(v string) *GetPublicDevicesResponseBodyData {
	s.SerialNumber = &v
	return s
}

func (s *GetPublicDevicesResponseBodyData) SetTitle(v string) *GetPublicDevicesResponseBodyData {
	s.Title = &v
	return s
}

type GetPublicDevicesResponseBodyDataDeviceDepts struct {
	// example:
	//
	// 123
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 测试部门
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
	// example:
	//
	// 测试角色
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 123
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
	// example:
	//
	// 张三
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 123
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPublicDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	Data []*GetPublisherSummaryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// false
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 2
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 100
	PublisherArticleCntStd *string `json:"publisherArticleCntStd,omitempty" xml:"publisherArticleCntStd,omitempty"`
	// example:
	//
	// 100
	PublisherArticlePvCntStd *string `json:"publisherArticlePvCntStd,omitempty" xml:"publisherArticlePvCntStd,omitempty"`
	// example:
	//
	// 阅读量最高文章，阅读量第二高文章
	PublisherArticlePvTop5 []*GetPublisherSummaryResponseBodyPublisherArticlePvTop5 `json:"publisherArticlePvTop5,omitempty" xml:"publisherArticlePvTop5,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	PublisherCntStd *string `json:"publisherCntStd,omitempty" xml:"publisherCntStd,omitempty"`
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
	// example:
	//
	// 100
	PublisherArticleCntStd *string `json:"publisherArticleCntStd,omitempty" xml:"publisherArticleCntStd,omitempty"`
	// example:
	//
	// 100
	PublisherArticlePvCntStd *string `json:"publisherArticlePvCntStd,omitempty" xml:"publisherArticlePvCntStd,omitempty"`
	// example:
	//
	// 服务窗1
	PublisherName *string `json:"publisherName,omitempty" xml:"publisherName,omitempty"`
	// example:
	//
	// 123
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
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
	// example:
	//
	// 文章1
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPublisherSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 123333
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// 1667000000
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1
	PersonIdentification *int32 `json:"personIdentification,omitempty" xml:"personIdentification,omitempty"`
	// example:
	//
	// 1
	Scene *int32 `json:"scene,omitempty" xml:"scene,omitempty"`
	// example:
	//
	// 1669000000
	ToTime  *int64    `json:"toTime,omitempty" xml:"toTime,omitempty"`
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
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
	Data []*GetRealPeopleRecordsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 20
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
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
	// example:
	//
	// agentId
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// 166700000
	InvokeTime *int64 `json:"invokeTime,omitempty" xml:"invokeTime,omitempty"`
	// example:
	//
	// 1
	PersonIdentification *int32 `json:"personIdentification,omitempty" xml:"personIdentification,omitempty"`
	// example:
	//
	// 1
	Platform *int32 `json:"platform,omitempty" xml:"platform,omitempty"`
	// example:
	//
	// 1
	Scene *int32 `json:"scene,omitempty" xml:"scene,omitempty"`
	// example:
	//
	// 1234
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRealPeopleRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 123333
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// 1
	FaceCompareResult *int32 `json:"faceCompareResult,omitempty" xml:"faceCompareResult,omitempty"`
	// example:
	//
	// 1667000000
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1669000000
	ToTime  *int64    `json:"toTime,omitempty" xml:"toTime,omitempty"`
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
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
	Data []*GetRecognizeRecordsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 20
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
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
	// example:
	//
	// agentId
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// 1
	FaceCompareResult *int32 `json:"faceCompareResult,omitempty" xml:"faceCompareResult,omitempty"`
	// example:
	//
	// 166700000
	InvokeTime *int64 `json:"invokeTime,omitempty" xml:"invokeTime,omitempty"`
	// example:
	//
	// 1
	Platform *int32 `json:"platform,omitempty" xml:"platform,omitempty"`
	// example:
	//
	// 1234
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRecognizeRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetRobotInfoByCodeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetRobotInfoByCodeHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetRobotInfoByCodeHeaders) GoString() string {
	return s.String()
}

func (s *GetRobotInfoByCodeHeaders) SetCommonHeaders(v map[string]*string) *GetRobotInfoByCodeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRobotInfoByCodeHeaders) SetXAcsDingtalkAccessToken(v string) *GetRobotInfoByCodeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetRobotInfoByCodeRequest struct {
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s GetRobotInfoByCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRobotInfoByCodeRequest) GoString() string {
	return s.String()
}

func (s *GetRobotInfoByCodeRequest) SetRobotCode(v string) *GetRobotInfoByCodeRequest {
	s.RobotCode = &v
	return s
}

type GetRobotInfoByCodeResponseBody struct {
	RobotInfoVO *GetRobotInfoByCodeResponseBodyRobotInfoVO `json:"robotInfoVO,omitempty" xml:"robotInfoVO,omitempty" type:"Struct"`
}

func (s GetRobotInfoByCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRobotInfoByCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetRobotInfoByCodeResponseBody) SetRobotInfoVO(v *GetRobotInfoByCodeResponseBodyRobotInfoVO) *GetRobotInfoByCodeResponseBody {
	s.RobotInfoVO = v
	return s
}

type GetRobotInfoByCodeResponseBodyRobotInfoVO struct {
	AgentId           *int64  `json:"agentId,omitempty" xml:"agentId,omitempty"`
	Brief             *string `json:"brief,omitempty" xml:"brief,omitempty"`
	Description       *string `json:"description,omitempty" xml:"description,omitempty"`
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
	RobotOrganization *int64  `json:"robotOrganization,omitempty" xml:"robotOrganization,omitempty"`
}

func (s GetRobotInfoByCodeResponseBodyRobotInfoVO) String() string {
	return tea.Prettify(s)
}

func (s GetRobotInfoByCodeResponseBodyRobotInfoVO) GoString() string {
	return s.String()
}

func (s *GetRobotInfoByCodeResponseBodyRobotInfoVO) SetAgentId(v int64) *GetRobotInfoByCodeResponseBodyRobotInfoVO {
	s.AgentId = &v
	return s
}

func (s *GetRobotInfoByCodeResponseBodyRobotInfoVO) SetBrief(v string) *GetRobotInfoByCodeResponseBodyRobotInfoVO {
	s.Brief = &v
	return s
}

func (s *GetRobotInfoByCodeResponseBodyRobotInfoVO) SetDescription(v string) *GetRobotInfoByCodeResponseBodyRobotInfoVO {
	s.Description = &v
	return s
}

func (s *GetRobotInfoByCodeResponseBodyRobotInfoVO) SetName(v string) *GetRobotInfoByCodeResponseBodyRobotInfoVO {
	s.Name = &v
	return s
}

func (s *GetRobotInfoByCodeResponseBodyRobotInfoVO) SetRobotOrganization(v int64) *GetRobotInfoByCodeResponseBodyRobotInfoVO {
	s.RobotOrganization = &v
	return s
}

type GetRobotInfoByCodeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRobotInfoByCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRobotInfoByCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRobotInfoByCodeResponse) GoString() string {
	return s.String()
}

func (s *GetRobotInfoByCodeResponse) SetHeaders(v map[string]*string) *GetRobotInfoByCodeResponse {
	s.Headers = v
	return s
}

func (s *GetRobotInfoByCodeResponse) SetStatusCode(v int32) *GetRobotInfoByCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRobotInfoByCodeResponse) SetBody(v *GetRobotInfoByCodeResponseBody) *GetRobotInfoByCodeResponse {
	s.Body = v
	return s
}

type GetSecurityConfigMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSecurityConfigMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityConfigMemberHeaders) GoString() string {
	return s.String()
}

func (s *GetSecurityConfigMemberHeaders) SetCommonHeaders(v map[string]*string) *GetSecurityConfigMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSecurityConfigMemberHeaders) SetXAcsDingtalkAccessToken(v string) *GetSecurityConfigMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSecurityConfigMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ctrl.xxx
	ConfigKey *string `json:"configKey,omitempty" xml:"configKey,omitempty"`
	// This parameter is required.
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	NextToken *float32 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetSecurityConfigMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityConfigMemberRequest) GoString() string {
	return s.String()
}

func (s *GetSecurityConfigMemberRequest) SetConfigKey(v string) *GetSecurityConfigMemberRequest {
	s.ConfigKey = &v
	return s
}

func (s *GetSecurityConfigMemberRequest) SetMaxResults(v int32) *GetSecurityConfigMemberRequest {
	s.MaxResults = &v
	return s
}

func (s *GetSecurityConfigMemberRequest) SetNextToken(v float32) *GetSecurityConfigMemberRequest {
	s.NextToken = &v
	return s
}

type GetSecurityConfigMemberResponseBody struct {
	Result  *GetSecurityConfigMemberResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetSecurityConfigMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityConfigMemberResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecurityConfigMemberResponseBody) SetResult(v *GetSecurityConfigMemberResponseBodyResult) *GetSecurityConfigMemberResponseBody {
	s.Result = v
	return s
}

func (s *GetSecurityConfigMemberResponseBody) SetSuccess(v bool) *GetSecurityConfigMemberResponseBody {
	s.Success = &v
	return s
}

type GetSecurityConfigMemberResponseBodyResult struct {
	HasNext   *bool                                                 `json:"hasNext,omitempty" xml:"hasNext,omitempty"`
	NextToken *float32                                              `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ScopeType *int32                                                `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	UserInfos []*GetSecurityConfigMemberResponseBodyResultUserInfos `json:"userInfos,omitempty" xml:"userInfos,omitempty" type:"Repeated"`
}

func (s GetSecurityConfigMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityConfigMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSecurityConfigMemberResponseBodyResult) SetHasNext(v bool) *GetSecurityConfigMemberResponseBodyResult {
	s.HasNext = &v
	return s
}

func (s *GetSecurityConfigMemberResponseBodyResult) SetNextToken(v float32) *GetSecurityConfigMemberResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *GetSecurityConfigMemberResponseBodyResult) SetScopeType(v int32) *GetSecurityConfigMemberResponseBodyResult {
	s.ScopeType = &v
	return s
}

func (s *GetSecurityConfigMemberResponseBodyResult) SetUserInfos(v []*GetSecurityConfigMemberResponseBodyResultUserInfos) *GetSecurityConfigMemberResponseBodyResult {
	s.UserInfos = v
	return s
}

type GetSecurityConfigMemberResponseBodyResultUserInfos struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetSecurityConfigMemberResponseBodyResultUserInfos) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityConfigMemberResponseBodyResultUserInfos) GoString() string {
	return s.String()
}

func (s *GetSecurityConfigMemberResponseBodyResultUserInfos) SetName(v string) *GetSecurityConfigMemberResponseBodyResultUserInfos {
	s.Name = &v
	return s
}

func (s *GetSecurityConfigMemberResponseBodyResultUserInfos) SetUserId(v string) *GetSecurityConfigMemberResponseBodyResultUserInfos {
	s.UserId = &v
	return s
}

type GetSecurityConfigMemberResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecurityConfigMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecurityConfigMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityConfigMemberResponse) GoString() string {
	return s.String()
}

func (s *GetSecurityConfigMemberResponse) SetHeaders(v map[string]*string) *GetSecurityConfigMemberResponse {
	s.Headers = v
	return s
}

func (s *GetSecurityConfigMemberResponse) SetStatusCode(v int32) *GetSecurityConfigMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecurityConfigMemberResponse) SetBody(v *GetSecurityConfigMemberResponseBody) *GetSecurityConfigMemberResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
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
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// 50
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1000
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
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
	// example:
	//
	// 部门1
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// example:
	//
	// **@**.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 小张
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// ***
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// example:
	//
	// 主管理员
	Roles *string `json:"roles,omitempty" xml:"roles,omitempty"`
	// example:
	//
	// 123***
	StaffId *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
	// example:
	//
	// 经理
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSignedDetailByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 1721718854814
	GmtCreateEnd *int64 `json:"gmtCreateEnd,omitempty" xml:"gmtCreateEnd,omitempty"`
	// example:
	//
	// 1721718854814
	GmtCreateStart *int64 `json:"gmtCreateStart,omitempty" xml:"gmtCreateStart,omitempty"`
	// example:
	//
	// 1721718854814
	GmtModifiedEnd *int64 `json:"gmtModifiedEnd,omitempty" xml:"gmtModifiedEnd,omitempty"`
	// example:
	//
	// 1721718854814
	GmtModifiedStart *int64 `json:"gmtModifiedStart,omitempty" xml:"gmtModifiedStart,omitempty"`
	// example:
	//
	// xx:xx:xx:xx:xx:xx
	MacAddress *string `json:"macAddress,omitempty" xml:"macAddress,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 500
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// Android
	Platform     *string   `json:"platform,omitempty" xml:"platform,omitempty"`
	SerialNumber *string   `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
	Status       *int32    `json:"status,omitempty" xml:"status,omitempty"`
	UserIds      []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s GetTrustDeviceListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTrustDeviceListRequest) GoString() string {
	return s.String()
}

func (s *GetTrustDeviceListRequest) SetGmtCreateEnd(v int64) *GetTrustDeviceListRequest {
	s.GmtCreateEnd = &v
	return s
}

func (s *GetTrustDeviceListRequest) SetGmtCreateStart(v int64) *GetTrustDeviceListRequest {
	s.GmtCreateStart = &v
	return s
}

func (s *GetTrustDeviceListRequest) SetGmtModifiedEnd(v int64) *GetTrustDeviceListRequest {
	s.GmtModifiedEnd = &v
	return s
}

func (s *GetTrustDeviceListRequest) SetGmtModifiedStart(v int64) *GetTrustDeviceListRequest {
	s.GmtModifiedStart = &v
	return s
}

func (s *GetTrustDeviceListRequest) SetMacAddress(v string) *GetTrustDeviceListRequest {
	s.MacAddress = &v
	return s
}

func (s *GetTrustDeviceListRequest) SetPageNumber(v int64) *GetTrustDeviceListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetTrustDeviceListRequest) SetPageSize(v int64) *GetTrustDeviceListRequest {
	s.PageSize = &v
	return s
}

func (s *GetTrustDeviceListRequest) SetPlatform(v string) *GetTrustDeviceListRequest {
	s.Platform = &v
	return s
}

func (s *GetTrustDeviceListRequest) SetSerialNumber(v string) *GetTrustDeviceListRequest {
	s.SerialNumber = &v
	return s
}

func (s *GetTrustDeviceListRequest) SetStatus(v int32) *GetTrustDeviceListRequest {
	s.Status = &v
	return s
}

func (s *GetTrustDeviceListRequest) SetUserIds(v []*string) *GetTrustDeviceListRequest {
	s.UserIds = v
	return s
}

type GetTrustDeviceListResponseBody struct {
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// This parameter is required.
	Data     []*GetTrustDeviceListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	PageSize *int64                                `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int64                                `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetTrustDeviceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrustDeviceListResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrustDeviceListResponseBody) SetCurrentPage(v int64) *GetTrustDeviceListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *GetTrustDeviceListResponseBody) SetData(v []*GetTrustDeviceListResponseBodyData) *GetTrustDeviceListResponseBody {
	s.Data = v
	return s
}

func (s *GetTrustDeviceListResponseBody) SetPageSize(v int64) *GetTrustDeviceListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetTrustDeviceListResponseBody) SetTotal(v int64) *GetTrustDeviceListResponseBody {
	s.Total = &v
	return s
}

type GetTrustDeviceListResponseBodyData struct {
	// example:
	//
	// 1628650483
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Id         *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 88:92:5a:1f:e8:24
	MacAddress   *string `json:"macAddress,omitempty" xml:"macAddress,omitempty"`
	Model        *string `json:"model,omitempty" xml:"model,omitempty"`
	ModifiedTime *int64  `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// example:
	//
	// Mac
	Platform     *string `json:"platform,omitempty" xml:"platform,omitempty"`
	SerialNumber *string `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
	// example:
	//
	// 2
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 设备标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 65224157501039784
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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

func (s *GetTrustDeviceListResponseBodyData) SetId(v int64) *GetTrustDeviceListResponseBodyData {
	s.Id = &v
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

func (s *GetTrustDeviceListResponseBodyData) SetModifiedTime(v int64) *GetTrustDeviceListResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *GetTrustDeviceListResponseBodyData) SetPlatform(v string) *GetTrustDeviceListResponseBodyData {
	s.Platform = &v
	return s
}

func (s *GetTrustDeviceListResponseBodyData) SetSerialNumber(v string) *GetTrustDeviceListResponseBodyData {
	s.SerialNumber = &v
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrustDeviceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	Data []*GetUserAppVersionSummaryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// false
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 10
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	// example:
	//
	// 6.0
	AppVersion *string `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	// example:
	//
	// iOS
	Client *string `json:"client,omitempty" xml:"client,omitempty"`
	// example:
	//
	// 组织1
	OrgName *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	// example:
	//
	// 20210808
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
	// example:
	//
	// 10
	UserCnt *float32 `json:"userCnt,omitempty" xml:"userCnt,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserAppVersionSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
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
	// example:
	//
	// 1
	State *int32 `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// userId
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserFaceStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
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
	// example:
	//
	// 1
	State *int32 `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// userId
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserRealPeopleStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20220501
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
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
	ItemList []*GetUserStayLengthResponseBodyItemList `json:"itemList,omitempty" xml:"itemList,omitempty" type:"Repeated"`
	// example:
	//
	// 1000
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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
	// example:
	//
	// 小张
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 20220501
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
	// example:
	//
	// 1000
	StayTimeLenApp1d *int64 `json:"stayTimeLenApp1d,omitempty" xml:"stayTimeLenApp1d,omitempty"`
	// example:
	//
	// 2000
	StayTimeLenPc1d *int64 `json:"stayTimeLenPc1d,omitempty" xml:"stayTimeLenPc1d,omitempty"`
	// example:
	//
	// 123***
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserStayLengthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetVirusScanResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetVirusScanResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetVirusScanResultHeaders) GoString() string {
	return s.String()
}

func (s *GetVirusScanResultHeaders) SetCommonHeaders(v map[string]*string) *GetVirusScanResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetVirusScanResultHeaders) SetXAcsDingtalkAccessToken(v string) *GetVirusScanResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetVirusScanResultRequest struct {
	// This parameter is required.
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetVirusScanResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVirusScanResultRequest) GoString() string {
	return s.String()
}

func (s *GetVirusScanResultRequest) SetTaskId(v string) *GetVirusScanResultRequest {
	s.TaskId = &v
	return s
}

type GetVirusScanResultResponseBody struct {
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	Status *int32  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetVirusScanResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVirusScanResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetVirusScanResultResponseBody) SetReason(v string) *GetVirusScanResultResponseBody {
	s.Reason = &v
	return s
}

func (s *GetVirusScanResultResponseBody) SetStatus(v int32) *GetVirusScanResultResponseBody {
	s.Status = &v
	return s
}

type GetVirusScanResultResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVirusScanResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVirusScanResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVirusScanResultResponse) GoString() string {
	return s.String()
}

func (s *GetVirusScanResultResponse) SetHeaders(v map[string]*string) *GetVirusScanResultResponse {
	s.Headers = v
	return s
}

func (s *GetVirusScanResultResponse) SetStatusCode(v int32) *GetVirusScanResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVirusScanResultResponse) SetBody(v *GetVirusScanResultResponseBody) *GetVirusScanResultResponse {
	s.Body = v
	return s
}

type GroupQueryByAttrHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GroupQueryByAttrHeaders) String() string {
	return tea.Prettify(s)
}

func (s GroupQueryByAttrHeaders) GoString() string {
	return s.String()
}

func (s *GroupQueryByAttrHeaders) SetCommonHeaders(v map[string]*string) *GroupQueryByAttrHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GroupQueryByAttrHeaders) SetXAcsDingtalkAccessToken(v string) *GroupQueryByAttrHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GroupQueryByAttrRequest struct {
	// This parameter is required.
	CorpId          *string                                   `json:"corpId,omitempty" xml:"corpId,omitempty"`
	GroupTopic      *string                                   `json:"groupTopic,omitempty" xml:"groupTopic,omitempty"`
	GroupType       *string                                   `json:"groupType,omitempty" xml:"groupType,omitempty"`
	ListDynamicAttr []*GroupQueryByAttrRequestListDynamicAttr `json:"listDynamicAttr,omitempty" xml:"listDynamicAttr,omitempty" type:"Repeated"`
	PageIndex       *int32                                    `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	PageSize        *int32                                    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	SecretKey *string `json:"secretKey,omitempty" xml:"secretKey,omitempty"`
}

func (s GroupQueryByAttrRequest) String() string {
	return tea.Prettify(s)
}

func (s GroupQueryByAttrRequest) GoString() string {
	return s.String()
}

func (s *GroupQueryByAttrRequest) SetCorpId(v string) *GroupQueryByAttrRequest {
	s.CorpId = &v
	return s
}

func (s *GroupQueryByAttrRequest) SetGroupTopic(v string) *GroupQueryByAttrRequest {
	s.GroupTopic = &v
	return s
}

func (s *GroupQueryByAttrRequest) SetGroupType(v string) *GroupQueryByAttrRequest {
	s.GroupType = &v
	return s
}

func (s *GroupQueryByAttrRequest) SetListDynamicAttr(v []*GroupQueryByAttrRequestListDynamicAttr) *GroupQueryByAttrRequest {
	s.ListDynamicAttr = v
	return s
}

func (s *GroupQueryByAttrRequest) SetPageIndex(v int32) *GroupQueryByAttrRequest {
	s.PageIndex = &v
	return s
}

func (s *GroupQueryByAttrRequest) SetPageSize(v int32) *GroupQueryByAttrRequest {
	s.PageSize = &v
	return s
}

func (s *GroupQueryByAttrRequest) SetSecretKey(v string) *GroupQueryByAttrRequest {
	s.SecretKey = &v
	return s
}

type GroupQueryByAttrRequestListDynamicAttr struct {
	// This parameter is required.
	AttrCode *string `json:"attrCode,omitempty" xml:"attrCode,omitempty"`
	// This parameter is required.
	ListAttrOptionsCode []*string `json:"listAttrOptionsCode,omitempty" xml:"listAttrOptionsCode,omitempty" type:"Repeated"`
}

func (s GroupQueryByAttrRequestListDynamicAttr) String() string {
	return tea.Prettify(s)
}

func (s GroupQueryByAttrRequestListDynamicAttr) GoString() string {
	return s.String()
}

func (s *GroupQueryByAttrRequestListDynamicAttr) SetAttrCode(v string) *GroupQueryByAttrRequestListDynamicAttr {
	s.AttrCode = &v
	return s
}

func (s *GroupQueryByAttrRequestListDynamicAttr) SetListAttrOptionsCode(v []*string) *GroupQueryByAttrRequestListDynamicAttr {
	s.ListAttrOptionsCode = v
	return s
}

type GroupQueryByAttrResponseBody struct {
	Code    *int32                            `json:"code,omitempty" xml:"code,omitempty"`
	Data    *GroupQueryByAttrResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                           `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GroupQueryByAttrResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GroupQueryByAttrResponseBody) GoString() string {
	return s.String()
}

func (s *GroupQueryByAttrResponseBody) SetCode(v int32) *GroupQueryByAttrResponseBody {
	s.Code = &v
	return s
}

func (s *GroupQueryByAttrResponseBody) SetData(v *GroupQueryByAttrResponseBodyData) *GroupQueryByAttrResponseBody {
	s.Data = v
	return s
}

func (s *GroupQueryByAttrResponseBody) SetMessage(v string) *GroupQueryByAttrResponseBody {
	s.Message = &v
	return s
}

type GroupQueryByAttrResponseBodyData struct {
	Counts    *int64                                  `json:"counts,omitempty" xml:"counts,omitempty"`
	List      []*GroupQueryByAttrResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageIndex *int64                                  `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	PageSize  *int64                                  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GroupQueryByAttrResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GroupQueryByAttrResponseBodyData) GoString() string {
	return s.String()
}

func (s *GroupQueryByAttrResponseBodyData) SetCounts(v int64) *GroupQueryByAttrResponseBodyData {
	s.Counts = &v
	return s
}

func (s *GroupQueryByAttrResponseBodyData) SetList(v []*GroupQueryByAttrResponseBodyDataList) *GroupQueryByAttrResponseBodyData {
	s.List = v
	return s
}

func (s *GroupQueryByAttrResponseBodyData) SetPageIndex(v int64) *GroupQueryByAttrResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *GroupQueryByAttrResponseBodyData) SetPageSize(v int64) *GroupQueryByAttrResponseBodyData {
	s.PageSize = &v
	return s
}

type GroupQueryByAttrResponseBodyDataList struct {
	GroupMemberCount   *int32  `json:"groupMemberCount,omitempty" xml:"groupMemberCount,omitempty"`
	GroupName          *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	OwnerJobNo         *string `json:"ownerJobNo,omitempty" xml:"ownerJobNo,omitempty"`
	OwnerUserName      *string `json:"ownerUserName,omitempty" xml:"ownerUserName,omitempty"`
}

func (s GroupQueryByAttrResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GroupQueryByAttrResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GroupQueryByAttrResponseBodyDataList) SetGroupMemberCount(v int32) *GroupQueryByAttrResponseBodyDataList {
	s.GroupMemberCount = &v
	return s
}

func (s *GroupQueryByAttrResponseBodyDataList) SetGroupName(v string) *GroupQueryByAttrResponseBodyDataList {
	s.GroupName = &v
	return s
}

func (s *GroupQueryByAttrResponseBodyDataList) SetOpenConversationId(v string) *GroupQueryByAttrResponseBodyDataList {
	s.OpenConversationId = &v
	return s
}

func (s *GroupQueryByAttrResponseBodyDataList) SetOwnerJobNo(v string) *GroupQueryByAttrResponseBodyDataList {
	s.OwnerJobNo = &v
	return s
}

func (s *GroupQueryByAttrResponseBodyDataList) SetOwnerUserName(v string) *GroupQueryByAttrResponseBodyDataList {
	s.OwnerUserName = &v
	return s
}

type GroupQueryByAttrResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GroupQueryByAttrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GroupQueryByAttrResponse) String() string {
	return tea.Prettify(s)
}

func (s GroupQueryByAttrResponse) GoString() string {
	return s.String()
}

func (s *GroupQueryByAttrResponse) SetHeaders(v map[string]*string) *GroupQueryByAttrResponse {
	s.Headers = v
	return s
}

func (s *GroupQueryByAttrResponse) SetStatusCode(v int32) *GroupQueryByAttrResponse {
	s.StatusCode = &v
	return s
}

func (s *GroupQueryByAttrResponse) SetBody(v *GroupQueryByAttrResponseBody) *GroupQueryByAttrResponse {
	s.Body = v
	return s
}

type GroupQueryByOpenIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GroupQueryByOpenIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GroupQueryByOpenIdHeaders) GoString() string {
	return s.String()
}

func (s *GroupQueryByOpenIdHeaders) SetCommonHeaders(v map[string]*string) *GroupQueryByOpenIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GroupQueryByOpenIdHeaders) SetXAcsDingtalkAccessToken(v string) *GroupQueryByOpenIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GroupQueryByOpenIdRequest struct {
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	SecretKey *string `json:"secretKey,omitempty" xml:"secretKey,omitempty"`
}

func (s GroupQueryByOpenIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GroupQueryByOpenIdRequest) GoString() string {
	return s.String()
}

func (s *GroupQueryByOpenIdRequest) SetOpenConversationId(v string) *GroupQueryByOpenIdRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GroupQueryByOpenIdRequest) SetSecretKey(v string) *GroupQueryByOpenIdRequest {
	s.SecretKey = &v
	return s
}

type GroupQueryByOpenIdResponseBody struct {
	Code    *int32                              `json:"code,omitempty" xml:"code,omitempty"`
	Data    *GroupQueryByOpenIdResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                             `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GroupQueryByOpenIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GroupQueryByOpenIdResponseBody) GoString() string {
	return s.String()
}

func (s *GroupQueryByOpenIdResponseBody) SetCode(v int32) *GroupQueryByOpenIdResponseBody {
	s.Code = &v
	return s
}

func (s *GroupQueryByOpenIdResponseBody) SetData(v *GroupQueryByOpenIdResponseBodyData) *GroupQueryByOpenIdResponseBody {
	s.Data = v
	return s
}

func (s *GroupQueryByOpenIdResponseBody) SetMessage(v string) *GroupQueryByOpenIdResponseBody {
	s.Message = &v
	return s
}

type GroupQueryByOpenIdResponseBodyData struct {
	GroupName            *string                                                   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	GroupTemplateId      *string                                                   `json:"groupTemplateId,omitempty" xml:"groupTemplateId,omitempty"`
	GroupTemplateName    *string                                                   `json:"groupTemplateName,omitempty" xml:"groupTemplateName,omitempty"`
	GroupTopic           *string                                                   `json:"groupTopic,omitempty" xml:"groupTopic,omitempty"`
	GroupType            *string                                                   `json:"groupType,omitempty" xml:"groupType,omitempty"`
	Id                   *int64                                                    `json:"id,omitempty" xml:"id,omitempty"`
	ListGroupDynamicAttr []*GroupQueryByOpenIdResponseBodyDataListGroupDynamicAttr `json:"listGroupDynamicAttr,omitempty" xml:"listGroupDynamicAttr,omitempty" type:"Repeated"`
	OpenConversationId   *string                                                   `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s GroupQueryByOpenIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GroupQueryByOpenIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GroupQueryByOpenIdResponseBodyData) SetGroupName(v string) *GroupQueryByOpenIdResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *GroupQueryByOpenIdResponseBodyData) SetGroupTemplateId(v string) *GroupQueryByOpenIdResponseBodyData {
	s.GroupTemplateId = &v
	return s
}

func (s *GroupQueryByOpenIdResponseBodyData) SetGroupTemplateName(v string) *GroupQueryByOpenIdResponseBodyData {
	s.GroupTemplateName = &v
	return s
}

func (s *GroupQueryByOpenIdResponseBodyData) SetGroupTopic(v string) *GroupQueryByOpenIdResponseBodyData {
	s.GroupTopic = &v
	return s
}

func (s *GroupQueryByOpenIdResponseBodyData) SetGroupType(v string) *GroupQueryByOpenIdResponseBodyData {
	s.GroupType = &v
	return s
}

func (s *GroupQueryByOpenIdResponseBodyData) SetId(v int64) *GroupQueryByOpenIdResponseBodyData {
	s.Id = &v
	return s
}

func (s *GroupQueryByOpenIdResponseBodyData) SetListGroupDynamicAttr(v []*GroupQueryByOpenIdResponseBodyDataListGroupDynamicAttr) *GroupQueryByOpenIdResponseBodyData {
	s.ListGroupDynamicAttr = v
	return s
}

func (s *GroupQueryByOpenIdResponseBodyData) SetOpenConversationId(v string) *GroupQueryByOpenIdResponseBodyData {
	s.OpenConversationId = &v
	return s
}

type GroupQueryByOpenIdResponseBodyDataListGroupDynamicAttr struct {
	AttrCode            *string   `json:"attrCode,omitempty" xml:"attrCode,omitempty"`
	ListAttrOptionsCode []*string `json:"listAttrOptionsCode,omitempty" xml:"listAttrOptionsCode,omitempty" type:"Repeated"`
}

func (s GroupQueryByOpenIdResponseBodyDataListGroupDynamicAttr) String() string {
	return tea.Prettify(s)
}

func (s GroupQueryByOpenIdResponseBodyDataListGroupDynamicAttr) GoString() string {
	return s.String()
}

func (s *GroupQueryByOpenIdResponseBodyDataListGroupDynamicAttr) SetAttrCode(v string) *GroupQueryByOpenIdResponseBodyDataListGroupDynamicAttr {
	s.AttrCode = &v
	return s
}

func (s *GroupQueryByOpenIdResponseBodyDataListGroupDynamicAttr) SetListAttrOptionsCode(v []*string) *GroupQueryByOpenIdResponseBodyDataListGroupDynamicAttr {
	s.ListAttrOptionsCode = v
	return s
}

type GroupQueryByOpenIdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GroupQueryByOpenIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GroupQueryByOpenIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GroupQueryByOpenIdResponse) GoString() string {
	return s.String()
}

func (s *GroupQueryByOpenIdResponse) SetHeaders(v map[string]*string) *GroupQueryByOpenIdResponse {
	s.Headers = v
	return s
}

func (s *GroupQueryByOpenIdResponse) SetStatusCode(v int32) *GroupQueryByOpenIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GroupQueryByOpenIdResponse) SetBody(v *GroupQueryByOpenIdResponseBody) *GroupQueryByOpenIdResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// 1577945731837
	EndDate *int64 `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// example:
	//
	// 6406817113
	NextBizId *int64 `json:"nextBizId,omitempty" xml:"nextBizId,omitempty"`
	// example:
	//
	// 1577340931837
	NextGmtCreate *int64 `json:"nextGmtCreate,omitempty" xml:"nextGmtCreate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 500
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1577340931837
	StartDate *int64 `json:"startDate,omitempty" xml:"startDate,omitempty"`
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
	// example:
	//
	// 0
	Action *int32 `json:"action,omitempty" xml:"action,omitempty"`
	// example:
	//
	// 企业群
	ActionView *string `json:"actionView,omitempty" xml:"actionView,omitempty"`
	// example:
	//
	// 11258620701
	BizId           *string                                        `json:"bizId,omitempty" xml:"bizId,omitempty"`
	DocMemberList   []*ListAuditLogResponseBodyListDocMemberList   `json:"docMemberList,omitempty" xml:"docMemberList,omitempty" type:"Repeated"`
	DocMobileUrl    *string                                        `json:"docMobileUrl,omitempty" xml:"docMobileUrl,omitempty"`
	DocPcUrl        *string                                        `json:"docPcUrl,omitempty" xml:"docPcUrl,omitempty"`
	DocReceiverList []*ListAuditLogResponseBodyListDocReceiverList `json:"docReceiverList,omitempty" xml:"docReceiverList,omitempty" type:"Repeated"`
	// example:
	//
	// 1577601221260
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1577601221260
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 1.1.1.1
	IpAddress *string `json:"ipAddress,omitempty" xml:"ipAddress,omitempty"`
	// example:
	//
	// 2
	OperateModule *int64 `json:"operateModule,omitempty" xml:"operateModule,omitempty"`
	// example:
	//
	// 企业群
	OperateModuleView *string `json:"operateModuleView,omitempty" xml:"operateModuleView,omitempty"`
	// example:
	//
	// 测试
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	// example:
	//
	// 水果公司
	OrgName *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	// example:
	//
	// 11
	Platform *int32 `json:"platform,omitempty" xml:"platform,omitempty"`
	// example:
	//
	// WIN
	PlatformView           *string `json:"platformView,omitempty" xml:"platformView,omitempty"`
	PrevWorkSpaceId        *int64  `json:"prevWorkSpaceId,omitempty" xml:"prevWorkSpaceId,omitempty"`
	PrevWorkSpaceMobileUrl *string `json:"prevWorkSpaceMobileUrl,omitempty" xml:"prevWorkSpaceMobileUrl,omitempty"`
	PrevWorkSpaceName      *string `json:"prevWorkSpaceName,omitempty" xml:"prevWorkSpaceName,omitempty"`
	PrevWorkSpacePcUrl     *string `json:"prevWorkSpacePcUrl,omitempty" xml:"prevWorkSpacePcUrl,omitempty"`
	// example:
	//
	// 张三
	RealName *string `json:"realName,omitempty" xml:"realName,omitempty"`
	// example:
	//
	// 总经理办公室
	ReceiverName *string `json:"receiverName,omitempty" xml:"receiverName,omitempty"`
	// example:
	//
	// 0
	ReceiverType *int32 `json:"receiverType,omitempty" xml:"receiverType,omitempty"`
	// example:
	//
	// 单聊
	ReceiverTypeView *string `json:"receiverTypeView,omitempty" xml:"receiverTypeView,omitempty"`
	// example:
	//
	// 文件名
	Resource *string `json:"resource,omitempty" xml:"resource,omitempty"`
	// example:
	//
	// doc
	ResourceExtension *string `json:"resourceExtension,omitempty" xml:"resourceExtension,omitempty"`
	// example:
	//
	// 1024
	ResourceSize *int64 `json:"resourceSize,omitempty" xml:"resourceSize,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 11258620
	TargetSpaceId *int64 `json:"targetSpaceId,omitempty" xml:"targetSpaceId,omitempty"`
	// example:
	//
	// 123
	UserId             *string `json:"userId,omitempty" xml:"userId,omitempty"`
	WorkSpaceId        *int64  `json:"workSpaceId,omitempty" xml:"workSpaceId,omitempty"`
	WorkSpaceMobileUrl *string `json:"workSpaceMobileUrl,omitempty" xml:"workSpaceMobileUrl,omitempty"`
	WorkSpaceName      *string `json:"workSpaceName,omitempty" xml:"workSpaceName,omitempty"`
	WorkSpacePcUrl     *string `json:"workSpacePcUrl,omitempty" xml:"workSpacePcUrl,omitempty"`
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
	// example:
	//
	// 张三
	MemberName *string `json:"memberName,omitempty" xml:"memberName,omitempty"`
	// example:
	//
	// 0
	MemberType *int32 `json:"memberType,omitempty" xml:"memberType,omitempty"`
	// example:
	//
	// 部门
	MemberTypeView *string `json:"memberTypeView,omitempty" xml:"memberTypeView,omitempty"`
	// example:
	//
	// 1
	PermissionRole *int64 `json:"permissionRole,omitempty" xml:"permissionRole,omitempty"`
	// example:
	//
	// 阅读者（可查看\下载）
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
	// example:
	//
	// 张三
	ReceiverName *string `json:"receiverName,omitempty" xml:"receiverName,omitempty"`
	// example:
	//
	// 1
	ReceiverType *int32 `json:"receiverType,omitempty" xml:"receiverType,omitempty"`
	// example:
	//
	// 单聊
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuditLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListByCodesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListByCodesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListByCodesHeaders) GoString() string {
	return s.String()
}

func (s *ListByCodesHeaders) SetCommonHeaders(v map[string]*string) *ListByCodesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListByCodesHeaders) SetXAcsDingtalkAccessToken(v string) *ListByCodesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListByCodesRequest struct {
	Body []*string `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s ListByCodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListByCodesRequest) GoString() string {
	return s.String()
}

func (s *ListByCodesRequest) SetBody(v []*string) *ListByCodesRequest {
	s.Body = v
	return s
}

type ListByCodesResponseBody struct {
	RobotInfoList []*ListByCodesResponseBodyRobotInfoList `json:"robotInfoList,omitempty" xml:"robotInfoList,omitempty" type:"Repeated"`
}

func (s ListByCodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListByCodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListByCodesResponseBody) SetRobotInfoList(v []*ListByCodesResponseBodyRobotInfoList) *ListByCodesResponseBody {
	s.RobotInfoList = v
	return s
}

type ListByCodesResponseBodyRobotInfoList struct {
	Brief          *string `json:"brief,omitempty" xml:"brief,omitempty"`
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	CreateAt       *int64  `json:"createAt,omitempty" xml:"createAt,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	Dev            *string `json:"dev,omitempty" xml:"dev,omitempty"`
	Icon           *string `json:"icon,omitempty" xml:"icon,omitempty"`
	ModifiedAt     *int64  `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	OutgoingToken  *string `json:"outgoingToken,omitempty" xml:"outgoingToken,omitempty"`
	OutgoingUrl    *string `json:"outgoingUrl,omitempty" xml:"outgoingUrl,omitempty"`
	PreviewMediaId *string `json:"previewMediaId,omitempty" xml:"previewMediaId,omitempty"`
	SourceUrl      *string `json:"sourceUrl,omitempty" xml:"sourceUrl,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListByCodesResponseBodyRobotInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListByCodesResponseBodyRobotInfoList) GoString() string {
	return s.String()
}

func (s *ListByCodesResponseBodyRobotInfoList) SetBrief(v string) *ListByCodesResponseBodyRobotInfoList {
	s.Brief = &v
	return s
}

func (s *ListByCodesResponseBodyRobotInfoList) SetCode(v string) *ListByCodesResponseBodyRobotInfoList {
	s.Code = &v
	return s
}

func (s *ListByCodesResponseBodyRobotInfoList) SetCreateAt(v int64) *ListByCodesResponseBodyRobotInfoList {
	s.CreateAt = &v
	return s
}

func (s *ListByCodesResponseBodyRobotInfoList) SetDescription(v string) *ListByCodesResponseBodyRobotInfoList {
	s.Description = &v
	return s
}

func (s *ListByCodesResponseBodyRobotInfoList) SetDev(v string) *ListByCodesResponseBodyRobotInfoList {
	s.Dev = &v
	return s
}

func (s *ListByCodesResponseBodyRobotInfoList) SetIcon(v string) *ListByCodesResponseBodyRobotInfoList {
	s.Icon = &v
	return s
}

func (s *ListByCodesResponseBodyRobotInfoList) SetModifiedAt(v int64) *ListByCodesResponseBodyRobotInfoList {
	s.ModifiedAt = &v
	return s
}

func (s *ListByCodesResponseBodyRobotInfoList) SetName(v string) *ListByCodesResponseBodyRobotInfoList {
	s.Name = &v
	return s
}

func (s *ListByCodesResponseBodyRobotInfoList) SetOutgoingToken(v string) *ListByCodesResponseBodyRobotInfoList {
	s.OutgoingToken = &v
	return s
}

func (s *ListByCodesResponseBodyRobotInfoList) SetOutgoingUrl(v string) *ListByCodesResponseBodyRobotInfoList {
	s.OutgoingUrl = &v
	return s
}

func (s *ListByCodesResponseBodyRobotInfoList) SetPreviewMediaId(v string) *ListByCodesResponseBodyRobotInfoList {
	s.PreviewMediaId = &v
	return s
}

func (s *ListByCodesResponseBodyRobotInfoList) SetSourceUrl(v string) *ListByCodesResponseBodyRobotInfoList {
	s.SourceUrl = &v
	return s
}

func (s *ListByCodesResponseBodyRobotInfoList) SetStatus(v int32) *ListByCodesResponseBodyRobotInfoList {
	s.Status = &v
	return s
}

type ListByCodesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListByCodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListByCodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListByCodesResponse) GoString() string {
	return s.String()
}

func (s *ListByCodesResponse) SetHeaders(v map[string]*string) *ListByCodesResponse {
	s.Headers = v
	return s
}

func (s *ListByCodesResponse) SetStatusCode(v int32) *ListByCodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListByCodesResponse) SetBody(v *ListByCodesResponseBody) *ListByCodesResponse {
	s.Body = v
	return s
}

type ListByPluginIdsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListByPluginIdsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListByPluginIdsHeaders) GoString() string {
	return s.String()
}

func (s *ListByPluginIdsHeaders) SetCommonHeaders(v map[string]*string) *ListByPluginIdsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListByPluginIdsHeaders) SetXAcsDingtalkAccessToken(v string) *ListByPluginIdsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListByPluginIdsRequest struct {
	Body []*string `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s ListByPluginIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListByPluginIdsRequest) GoString() string {
	return s.String()
}

func (s *ListByPluginIdsRequest) SetBody(v []*string) *ListByPluginIdsRequest {
	s.Body = v
	return s
}

type ListByPluginIdsResponseBody struct {
	PluginInfoList []*ListByPluginIdsResponseBodyPluginInfoList `json:"pluginInfoList,omitempty" xml:"pluginInfoList,omitempty" type:"Repeated"`
}

func (s ListByPluginIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListByPluginIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListByPluginIdsResponseBody) SetPluginInfoList(v []*ListByPluginIdsResponseBodyPluginInfoList) *ListByPluginIdsResponseBody {
	s.PluginInfoList = v
	return s
}

type ListByPluginIdsResponseBodyPluginInfoList struct {
	AppId      *string `json:"appId,omitempty" xml:"appId,omitempty"`
	CreateAt   *int64  `json:"createAt,omitempty" xml:"createAt,omitempty"`
	Desc       *string `json:"desc,omitempty" xml:"desc,omitempty"`
	Icons      *string `json:"icons,omitempty" xml:"icons,omitempty"`
	ModifiedAt *int64  `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	PcUrl      *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	PluginId   *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	Status     *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url        *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListByPluginIdsResponseBodyPluginInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListByPluginIdsResponseBodyPluginInfoList) GoString() string {
	return s.String()
}

func (s *ListByPluginIdsResponseBodyPluginInfoList) SetAppId(v string) *ListByPluginIdsResponseBodyPluginInfoList {
	s.AppId = &v
	return s
}

func (s *ListByPluginIdsResponseBodyPluginInfoList) SetCreateAt(v int64) *ListByPluginIdsResponseBodyPluginInfoList {
	s.CreateAt = &v
	return s
}

func (s *ListByPluginIdsResponseBodyPluginInfoList) SetDesc(v string) *ListByPluginIdsResponseBodyPluginInfoList {
	s.Desc = &v
	return s
}

func (s *ListByPluginIdsResponseBodyPluginInfoList) SetIcons(v string) *ListByPluginIdsResponseBodyPluginInfoList {
	s.Icons = &v
	return s
}

func (s *ListByPluginIdsResponseBodyPluginInfoList) SetModifiedAt(v int64) *ListByPluginIdsResponseBodyPluginInfoList {
	s.ModifiedAt = &v
	return s
}

func (s *ListByPluginIdsResponseBodyPluginInfoList) SetName(v string) *ListByPluginIdsResponseBodyPluginInfoList {
	s.Name = &v
	return s
}

func (s *ListByPluginIdsResponseBodyPluginInfoList) SetPcUrl(v string) *ListByPluginIdsResponseBodyPluginInfoList {
	s.PcUrl = &v
	return s
}

func (s *ListByPluginIdsResponseBodyPluginInfoList) SetPluginId(v string) *ListByPluginIdsResponseBodyPluginInfoList {
	s.PluginId = &v
	return s
}

func (s *ListByPluginIdsResponseBodyPluginInfoList) SetStatus(v int32) *ListByPluginIdsResponseBodyPluginInfoList {
	s.Status = &v
	return s
}

func (s *ListByPluginIdsResponseBodyPluginInfoList) SetUrl(v string) *ListByPluginIdsResponseBodyPluginInfoList {
	s.Url = &v
	return s
}

type ListByPluginIdsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListByPluginIdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListByPluginIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListByPluginIdsResponse) GoString() string {
	return s.String()
}

func (s *ListByPluginIdsResponse) SetHeaders(v map[string]*string) *ListByPluginIdsResponse {
	s.Headers = v
	return s
}

func (s *ListByPluginIdsResponse) SetStatusCode(v int32) *ListByPluginIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListByPluginIdsResponse) SetBody(v *ListByPluginIdsResponseBody) *ListByPluginIdsResponse {
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
	// This parameter is required.
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
	// example:
	//
	// 1
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
	// This parameter is required.
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCategorysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 15800000000
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
	// This parameter is required.
	//
	// example:
	//
	// ding32xxxxxxxxe0105d
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testCode
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 钉钉(中国)信息技术有限公司
	OrgFullName *string `json:"orgFullName,omitempty" xml:"orgFullName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 钉钉
	OrgName *int64 `json:"orgName,omitempty" xml:"orgName,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJoinOrgInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 5000003
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	// This parameter is required.
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// 1
	BuildStatus *int64 `json:"buildStatus,omitempty" xml:"buildStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.0.5
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMiniAppAvailableVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 500000003
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 0-打包中 ，1-成功，2-失败
	BuildStatus *int64 `json:"buildStatus,omitempty" xml:"buildStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://xxx.con/download/id
	H5Bundle *string `json:"h5Bundle,omitempty" xml:"h5Bundle,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5000
	PackageSize *string `json:"packageSize,omitempty" xml:"packageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://xxx.con/download/id
	PackageUrl *string `json:"packageUrl,omitempty" xml:"packageUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.0.5
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMiniAppHistoryVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
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
	// example:
	//
	// 123
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	IsNecessary *int32 `json:"isNecessary,omitempty" xml:"isNecessary,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 供应商
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	VisibleDepts []*ListPartnerRolesResponseBodyListVisibleDepts `json:"visibleDepts,omitempty" xml:"visibleDepts,omitempty" type:"Repeated"`
	// This parameter is required.
	VisibleUsers []*ListPartnerRolesResponseBodyListVisibleUsers `json:"visibleUsers,omitempty" xml:"visibleUsers,omitempty" type:"Repeated"`
	// This parameter is required.
	WarningDepts []*ListPartnerRolesResponseBodyListWarningDepts `json:"warningDepts,omitempty" xml:"warningDepts,omitempty" type:"Repeated"`
	// This parameter is required.
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
	// This parameter is required.
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
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
	// This parameter is required.
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPartnerRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 2aa6736c715944329xxxxxxxxd38be41
	BizInstanceId *string `json:"bizInstanceId,omitempty" xml:"bizInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 2022-03-13
	ScheduleDateEnd *string `json:"scheduleDateEnd,omitempty" xml:"scheduleDateEnd,omitempty"`
	// example:
	//
	// 2022-03-13
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
	// example:
	//
	// false
	HasMore *bool                                                     `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List    []*ListPunchScheduleByConditionWithPagingResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// example:
	//
	// 10000
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	// example:
	//
	// be0d84e04316488cxxxxxxxx129522b0
	BizOuterId *string `json:"bizOuterId,omitempty" xml:"bizOuterId,omitempty"`
	// example:
	//
	// 测试打卡机
	PositionName *string `json:"positionName,omitempty" xml:"positionName,omitempty"`
	// example:
	//
	// checkIn
	PunchSymbol *string `json:"punchSymbol,omitempty" xml:"punchSymbol,omitempty"`
	// example:
	//
	// 200003
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// 1647333408000
	UserPunchTime *int64 `json:"userPunchTime,omitempty" xml:"userPunchTime,omitempty"`
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
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPunchScheduleByConditionWithPagingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
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
	// example:
	//
	// 1
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
	// This parameter is required.
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 1
	DeviceType *string `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a523641
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LogoutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type OpenBenefitPackageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s OpenBenefitPackageHeaders) String() string {
	return tea.Prettify(s)
}

func (s OpenBenefitPackageHeaders) GoString() string {
	return s.String()
}

func (s *OpenBenefitPackageHeaders) SetCommonHeaders(v map[string]*string) *OpenBenefitPackageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OpenBenefitPackageHeaders) SetXAcsDingtalkAccessToken(v string) *OpenBenefitPackageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type OpenBenefitPackageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	BenefitPackage *int32 `json:"benefitPackage,omitempty" xml:"benefitPackage,omitempty"`
	// This parameter is required.
	EndDate *int64 `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// This parameter is required.
	StartDate *int64 `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dingxxxxx
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
}

func (s OpenBenefitPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenBenefitPackageRequest) GoString() string {
	return s.String()
}

func (s *OpenBenefitPackageRequest) SetBenefitPackage(v int32) *OpenBenefitPackageRequest {
	s.BenefitPackage = &v
	return s
}

func (s *OpenBenefitPackageRequest) SetEndDate(v int64) *OpenBenefitPackageRequest {
	s.EndDate = &v
	return s
}

func (s *OpenBenefitPackageRequest) SetStartDate(v int64) *OpenBenefitPackageRequest {
	s.StartDate = &v
	return s
}

func (s *OpenBenefitPackageRequest) SetTargetCorpId(v string) *OpenBenefitPackageRequest {
	s.TargetCorpId = &v
	return s
}

type OpenBenefitPackageResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OpenBenefitPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenBenefitPackageResponseBody) GoString() string {
	return s.String()
}

func (s *OpenBenefitPackageResponseBody) SetSuccess(v bool) *OpenBenefitPackageResponseBody {
	s.Success = &v
	return s
}

type OpenBenefitPackageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenBenefitPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenBenefitPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenBenefitPackageResponse) GoString() string {
	return s.String()
}

func (s *OpenBenefitPackageResponse) SetHeaders(v map[string]*string) *OpenBenefitPackageResponse {
	s.Headers = v
	return s
}

func (s *OpenBenefitPackageResponse) SetStatusCode(v int32) *OpenBenefitPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenBenefitPackageResponse) SetBody(v *OpenBenefitPackageResponseBody) *OpenBenefitPackageResponse {
	s.Body = v
	return s
}

type OpportunitySearchHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s OpportunitySearchHeaders) String() string {
	return tea.Prettify(s)
}

func (s OpportunitySearchHeaders) GoString() string {
	return s.String()
}

func (s *OpportunitySearchHeaders) SetCommonHeaders(v map[string]*string) *OpportunitySearchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OpportunitySearchHeaders) SetXAcsDingtalkAccessToken(v string) *OpportunitySearchHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type OpportunitySearchRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 测试组织
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
}

func (s OpportunitySearchRequest) String() string {
	return tea.Prettify(s)
}

func (s OpportunitySearchRequest) GoString() string {
	return s.String()
}

func (s *OpportunitySearchRequest) SetTargetCorpId(v string) *OpportunitySearchRequest {
	s.TargetCorpId = &v
	return s
}

type OpportunitySearchResponseBody struct {
	OpportunityExist *bool `json:"opportunityExist,omitempty" xml:"opportunityExist,omitempty"`
}

func (s OpportunitySearchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpportunitySearchResponseBody) GoString() string {
	return s.String()
}

func (s *OpportunitySearchResponseBody) SetOpportunityExist(v bool) *OpportunitySearchResponseBody {
	s.OpportunityExist = &v
	return s
}

type OpportunitySearchResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpportunitySearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpportunitySearchResponse) String() string {
	return tea.Prettify(s)
}

func (s OpportunitySearchResponse) GoString() string {
	return s.String()
}

func (s *OpportunitySearchResponse) SetHeaders(v map[string]*string) *OpportunitySearchResponse {
	s.Headers = v
	return s
}

func (s *OpportunitySearchResponse) SetStatusCode(v int32) *OpportunitySearchResponse {
	s.StatusCode = &v
	return s
}

func (s *OpportunitySearchResponse) SetBody(v *OpportunitySearchResponseBody) *OpportunitySearchResponse {
	s.Body = v
	return s
}

type PreventCheatingCheckRiskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PreventCheatingCheckRiskHeaders) String() string {
	return tea.Prettify(s)
}

func (s PreventCheatingCheckRiskHeaders) GoString() string {
	return s.String()
}

func (s *PreventCheatingCheckRiskHeaders) SetCommonHeaders(v map[string]*string) *PreventCheatingCheckRiskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PreventCheatingCheckRiskHeaders) SetXAcsDingtalkAccessToken(v string) *PreventCheatingCheckRiskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PreventCheatingCheckRiskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 6.3.30
	ClientVer *string `json:"clientVer,omitempty" xml:"clientVer,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// iOS
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 15
	PlatformVer *string `json:"platformVer,omitempty" xml:"platformVer,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"lbsWuaToken": "lbsWua","ddSec":"ddSec"}
	Sec *string `json:"sec,omitempty" xml:"sec,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user1
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PreventCheatingCheckRiskRequest) String() string {
	return tea.Prettify(s)
}

func (s PreventCheatingCheckRiskRequest) GoString() string {
	return s.String()
}

func (s *PreventCheatingCheckRiskRequest) SetClientVer(v string) *PreventCheatingCheckRiskRequest {
	s.ClientVer = &v
	return s
}

func (s *PreventCheatingCheckRiskRequest) SetPlatform(v string) *PreventCheatingCheckRiskRequest {
	s.Platform = &v
	return s
}

func (s *PreventCheatingCheckRiskRequest) SetPlatformVer(v string) *PreventCheatingCheckRiskRequest {
	s.PlatformVer = &v
	return s
}

func (s *PreventCheatingCheckRiskRequest) SetSec(v string) *PreventCheatingCheckRiskRequest {
	s.Sec = &v
	return s
}

func (s *PreventCheatingCheckRiskRequest) SetUserId(v string) *PreventCheatingCheckRiskRequest {
	s.UserId = &v
	return s
}

type PreventCheatingCheckRiskResponseBody struct {
	Result *PreventCheatingCheckRiskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s PreventCheatingCheckRiskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PreventCheatingCheckRiskResponseBody) GoString() string {
	return s.String()
}

func (s *PreventCheatingCheckRiskResponseBody) SetResult(v *PreventCheatingCheckRiskResponseBodyResult) *PreventCheatingCheckRiskResponseBody {
	s.Result = v
	return s
}

type PreventCheatingCheckRiskResponseBodyResult struct {
	// This parameter is required.
	HasRisk  *bool              `json:"hasRisk,omitempty" xml:"hasRisk,omitempty"`
	RiskInfo map[string]*string `json:"riskInfo,omitempty" xml:"riskInfo,omitempty"`
}

func (s PreventCheatingCheckRiskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PreventCheatingCheckRiskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PreventCheatingCheckRiskResponseBodyResult) SetHasRisk(v bool) *PreventCheatingCheckRiskResponseBodyResult {
	s.HasRisk = &v
	return s
}

func (s *PreventCheatingCheckRiskResponseBodyResult) SetRiskInfo(v map[string]*string) *PreventCheatingCheckRiskResponseBodyResult {
	s.RiskInfo = v
	return s
}

type PreventCheatingCheckRiskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreventCheatingCheckRiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreventCheatingCheckRiskResponse) String() string {
	return tea.Prettify(s)
}

func (s PreventCheatingCheckRiskResponse) GoString() string {
	return s.String()
}

func (s *PreventCheatingCheckRiskResponse) SetHeaders(v map[string]*string) *PreventCheatingCheckRiskResponse {
	s.Headers = v
	return s
}

func (s *PreventCheatingCheckRiskResponse) SetStatusCode(v int32) *PreventCheatingCheckRiskResponse {
	s.StatusCode = &v
	return s
}

func (s *PreventCheatingCheckRiskResponse) SetBody(v *PreventCheatingCheckRiskResponseBody) *PreventCheatingCheckRiskResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// example:
	//
	// 1
	OperateType *string `json:"operateType,omitempty" xml:"operateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	OperatorUnionId *string `json:"operatorUnionId,omitempty" xml:"operatorUnionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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
	// example:
	//
	// 1
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
	// This parameter is required.
	//
	// example:
	//
	// true
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 110000000
	AgentId    *string                       `json:"agentId,omitempty" xml:"agentId,omitempty"`
	BadgeItems []*PushBadgeRequestBadgeItems `json:"badgeItems,omitempty" xml:"badgeItems,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	PushType *string `json:"pushType,omitempty" xml:"pushType,omitempty"`
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
	// example:
	//
	// 1
	PushValue *string `json:"pushValue,omitempty" xml:"pushValue,omitempty"`
	// example:
	//
	// 12345
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	// example:
	//
	// true
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushBadgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 1
	TargetCloudType *int32 `json:"targetCloudType,omitempty" xml:"targetCloudType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ding77b8cac4e026cc123xxxxxxxxeb6378f
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// sampleKeyId1234
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sampleSecretId1234
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxxbucket
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CloudType *int32 `json:"cloudType,omitempty" xml:"cloudType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://oss-cn-test.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAcrossCloudStroageConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type QueryChannelStaffInfoByMobileHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryChannelStaffInfoByMobileHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryChannelStaffInfoByMobileHeaders) GoString() string {
	return s.String()
}

func (s *QueryChannelStaffInfoByMobileHeaders) SetCommonHeaders(v map[string]*string) *QueryChannelStaffInfoByMobileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryChannelStaffInfoByMobileHeaders) SetXAcsDingtalkAccessToken(v string) *QueryChannelStaffInfoByMobileHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryChannelStaffInfoByMobileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 158xxxxxxxx
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dingxxxxxxx
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
}

func (s QueryChannelStaffInfoByMobileRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryChannelStaffInfoByMobileRequest) GoString() string {
	return s.String()
}

func (s *QueryChannelStaffInfoByMobileRequest) SetMobile(v string) *QueryChannelStaffInfoByMobileRequest {
	s.Mobile = &v
	return s
}

func (s *QueryChannelStaffInfoByMobileRequest) SetTargetCorpId(v string) *QueryChannelStaffInfoByMobileRequest {
	s.TargetCorpId = &v
	return s
}

type QueryChannelStaffInfoByMobileResponseBody struct {
	EmpInfo                     *QueryChannelStaffInfoByMobileResponseBodyEmpInfo                       `json:"empInfo,omitempty" xml:"empInfo,omitempty" type:"Struct"`
	ExclusiveAccountEmpInfoList []*QueryChannelStaffInfoByMobileResponseBodyExclusiveAccountEmpInfoList `json:"exclusiveAccountEmpInfoList,omitempty" xml:"exclusiveAccountEmpInfoList,omitempty" type:"Repeated"`
}

func (s QueryChannelStaffInfoByMobileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryChannelStaffInfoByMobileResponseBody) GoString() string {
	return s.String()
}

func (s *QueryChannelStaffInfoByMobileResponseBody) SetEmpInfo(v *QueryChannelStaffInfoByMobileResponseBodyEmpInfo) *QueryChannelStaffInfoByMobileResponseBody {
	s.EmpInfo = v
	return s
}

func (s *QueryChannelStaffInfoByMobileResponseBody) SetExclusiveAccountEmpInfoList(v []*QueryChannelStaffInfoByMobileResponseBodyExclusiveAccountEmpInfoList) *QueryChannelStaffInfoByMobileResponseBody {
	s.ExclusiveAccountEmpInfoList = v
	return s
}

type QueryChannelStaffInfoByMobileResponseBodyEmpInfo struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryChannelStaffInfoByMobileResponseBodyEmpInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryChannelStaffInfoByMobileResponseBodyEmpInfo) GoString() string {
	return s.String()
}

func (s *QueryChannelStaffInfoByMobileResponseBodyEmpInfo) SetName(v string) *QueryChannelStaffInfoByMobileResponseBodyEmpInfo {
	s.Name = &v
	return s
}

func (s *QueryChannelStaffInfoByMobileResponseBodyEmpInfo) SetUserId(v string) *QueryChannelStaffInfoByMobileResponseBodyEmpInfo {
	s.UserId = &v
	return s
}

type QueryChannelStaffInfoByMobileResponseBodyExclusiveAccountEmpInfoList struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryChannelStaffInfoByMobileResponseBodyExclusiveAccountEmpInfoList) String() string {
	return tea.Prettify(s)
}

func (s QueryChannelStaffInfoByMobileResponseBodyExclusiveAccountEmpInfoList) GoString() string {
	return s.String()
}

func (s *QueryChannelStaffInfoByMobileResponseBodyExclusiveAccountEmpInfoList) SetName(v string) *QueryChannelStaffInfoByMobileResponseBodyExclusiveAccountEmpInfoList {
	s.Name = &v
	return s
}

func (s *QueryChannelStaffInfoByMobileResponseBodyExclusiveAccountEmpInfoList) SetUserId(v string) *QueryChannelStaffInfoByMobileResponseBodyExclusiveAccountEmpInfoList {
	s.UserId = &v
	return s
}

type QueryChannelStaffInfoByMobileResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryChannelStaffInfoByMobileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryChannelStaffInfoByMobileResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryChannelStaffInfoByMobileResponse) GoString() string {
	return s.String()
}

func (s *QueryChannelStaffInfoByMobileResponse) SetHeaders(v map[string]*string) *QueryChannelStaffInfoByMobileResponse {
	s.Headers = v
	return s
}

func (s *QueryChannelStaffInfoByMobileResponse) SetStatusCode(v int32) *QueryChannelStaffInfoByMobileResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryChannelStaffInfoByMobileResponse) SetBody(v *QueryChannelStaffInfoByMobileResponseBody) *QueryChannelStaffInfoByMobileResponse {
	s.Body = v
	return s
}

type QueryConversationPageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryConversationPageHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryConversationPageHeaders) GoString() string {
	return s.String()
}

func (s *QueryConversationPageHeaders) SetCommonHeaders(v map[string]*string) *QueryConversationPageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryConversationPageHeaders) SetXAcsDingtalkAccessToken(v string) *QueryConversationPageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryConversationPageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// -1
	CategoryId *int64 `json:"categoryId,omitempty" xml:"categoryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryConversationPageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryConversationPageRequest) GoString() string {
	return s.String()
}

func (s *QueryConversationPageRequest) SetCategoryId(v int64) *QueryConversationPageRequest {
	s.CategoryId = &v
	return s
}

func (s *QueryConversationPageRequest) SetMaxResults(v int32) *QueryConversationPageRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryConversationPageRequest) SetNextToken(v string) *QueryConversationPageRequest {
	s.NextToken = &v
	return s
}

type QueryConversationPageResponseBody struct {
	Result  *QueryConversationPageResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                    `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryConversationPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryConversationPageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryConversationPageResponseBody) SetResult(v *QueryConversationPageResponseBodyResult) *QueryConversationPageResponseBody {
	s.Result = v
	return s
}

func (s *QueryConversationPageResponseBody) SetSuccess(v bool) *QueryConversationPageResponseBody {
	s.Success = &v
	return s
}

type QueryConversationPageResponseBodyResult struct {
	Data []*QueryConversationPageResponseBodyResultData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 10
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 999
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryConversationPageResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryConversationPageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryConversationPageResponseBodyResult) SetData(v []*QueryConversationPageResponseBodyResultData) *QueryConversationPageResponseBodyResult {
	s.Data = v
	return s
}

func (s *QueryConversationPageResponseBodyResult) SetMaxResults(v int32) *QueryConversationPageResponseBodyResult {
	s.MaxResults = &v
	return s
}

func (s *QueryConversationPageResponseBodyResult) SetNextToken(v string) *QueryConversationPageResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *QueryConversationPageResponseBodyResult) SetTotalCount(v int32) *QueryConversationPageResponseBodyResult {
	s.TotalCount = &v
	return s
}

type QueryConversationPageResponseBodyResultData struct {
	CategoryId *int64 `json:"categoryId,omitempty" xml:"categoryId,omitempty"`
	// example:
	//
	// 未分组
	CategoryName *string `json:"categoryName,omitempty" xml:"categoryName,omitempty"`
	// example:
	//
	// S24A01123
	GroupCode       *string `json:"groupCode,omitempty" xml:"groupCode,omitempty"`
	GroupMembersCnt *int32  `json:"groupMembersCnt,omitempty" xml:"groupMembersCnt,omitempty"`
	// example:
	//
	// 群聊
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// example:
	//
	// ownername
	GroupOwnerName *string `json:"groupOwnerName,omitempty" xml:"groupOwnerName,omitempty"`
	// example:
	//
	// useridxxx
	GroupOwnerUserId *string `json:"groupOwnerUserId,omitempty" xml:"groupOwnerUserId,omitempty"`
	IsKpConversation *bool   `json:"isKpConversation,omitempty" xml:"isKpConversation,omitempty"`
	// example:
	//
	// 1
	ManageSign *int32 `json:"manageSign,omitempty" xml:"manageSign,omitempty"`
	// example:
	//
	// cid123xxxxxx
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	Order              *int32  `json:"order,omitempty" xml:"order,omitempty"`
	Status             *int32  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryConversationPageResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s QueryConversationPageResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *QueryConversationPageResponseBodyResultData) SetCategoryId(v int64) *QueryConversationPageResponseBodyResultData {
	s.CategoryId = &v
	return s
}

func (s *QueryConversationPageResponseBodyResultData) SetCategoryName(v string) *QueryConversationPageResponseBodyResultData {
	s.CategoryName = &v
	return s
}

func (s *QueryConversationPageResponseBodyResultData) SetGroupCode(v string) *QueryConversationPageResponseBodyResultData {
	s.GroupCode = &v
	return s
}

func (s *QueryConversationPageResponseBodyResultData) SetGroupMembersCnt(v int32) *QueryConversationPageResponseBodyResultData {
	s.GroupMembersCnt = &v
	return s
}

func (s *QueryConversationPageResponseBodyResultData) SetGroupName(v string) *QueryConversationPageResponseBodyResultData {
	s.GroupName = &v
	return s
}

func (s *QueryConversationPageResponseBodyResultData) SetGroupOwnerName(v string) *QueryConversationPageResponseBodyResultData {
	s.GroupOwnerName = &v
	return s
}

func (s *QueryConversationPageResponseBodyResultData) SetGroupOwnerUserId(v string) *QueryConversationPageResponseBodyResultData {
	s.GroupOwnerUserId = &v
	return s
}

func (s *QueryConversationPageResponseBodyResultData) SetIsKpConversation(v bool) *QueryConversationPageResponseBodyResultData {
	s.IsKpConversation = &v
	return s
}

func (s *QueryConversationPageResponseBodyResultData) SetManageSign(v int32) *QueryConversationPageResponseBodyResultData {
	s.ManageSign = &v
	return s
}

func (s *QueryConversationPageResponseBodyResultData) SetOpenConversationId(v string) *QueryConversationPageResponseBodyResultData {
	s.OpenConversationId = &v
	return s
}

func (s *QueryConversationPageResponseBodyResultData) SetOrder(v int32) *QueryConversationPageResponseBodyResultData {
	s.Order = &v
	return s
}

func (s *QueryConversationPageResponseBodyResultData) SetStatus(v int32) *QueryConversationPageResponseBodyResultData {
	s.Status = &v
	return s
}

type QueryConversationPageResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryConversationPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryConversationPageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryConversationPageResponse) GoString() string {
	return s.String()
}

func (s *QueryConversationPageResponse) SetHeaders(v map[string]*string) *QueryConversationPageResponse {
	s.Headers = v
	return s
}

func (s *QueryConversationPageResponse) SetStatusCode(v int32) *QueryConversationPageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryConversationPageResponse) SetBody(v *QueryConversationPageResponseBody) *QueryConversationPageResponse {
	s.Body = v
	return s
}

type QueryExclusiveBenefitsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryExclusiveBenefitsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryExclusiveBenefitsHeaders) GoString() string {
	return s.String()
}

func (s *QueryExclusiveBenefitsHeaders) SetCommonHeaders(v map[string]*string) *QueryExclusiveBenefitsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryExclusiveBenefitsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryExclusiveBenefitsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryExclusiveBenefitsResponseBody struct {
	BenefitsList []*string `json:"benefitsList,omitempty" xml:"benefitsList,omitempty" type:"Repeated"`
}

func (s QueryExclusiveBenefitsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryExclusiveBenefitsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryExclusiveBenefitsResponseBody) SetBenefitsList(v []*string) *QueryExclusiveBenefitsResponseBody {
	s.BenefitsList = v
	return s
}

type QueryExclusiveBenefitsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryExclusiveBenefitsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryExclusiveBenefitsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryExclusiveBenefitsResponse) GoString() string {
	return s.String()
}

func (s *QueryExclusiveBenefitsResponse) SetHeaders(v map[string]*string) *QueryExclusiveBenefitsResponse {
	s.Headers = v
	return s
}

func (s *QueryExclusiveBenefitsResponse) SetStatusCode(v int32) *QueryExclusiveBenefitsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryExclusiveBenefitsResponse) SetBody(v *QueryExclusiveBenefitsResponseBody) *QueryExclusiveBenefitsResponse {
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
	// This parameter is required.
	MemberCount             *int64                                                              `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	PartnerLabelModelLevel1 *QueryPartnerInfoResponseBodyPartnerDeptListPartnerLabelModelLevel1 `json:"partnerLabelModelLevel1,omitempty" xml:"partnerLabelModelLevel1,omitempty" type:"Struct"`
	PartnerNum              *string                                                             `json:"partnerNum,omitempty" xml:"partnerNum,omitempty"`
	// This parameter is required.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPartnerInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type QueryTemplateInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryTemplateInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryTemplateInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryTemplateInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryTemplateInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryTemplateInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryTemplateInfoResponseBody struct {
	AbilitySwitch          *int64                                               `json:"abilitySwitch,omitempty" xml:"abilitySwitch,omitempty"`
	AppInfo                *QueryTemplateInfoResponseBodyAppInfo                `json:"appInfo,omitempty" xml:"appInfo,omitempty" type:"Struct"`
	ConversationScope      []*string                                            `json:"conversationScope,omitempty" xml:"conversationScope,omitempty" type:"Repeated"`
	CreateAt               *int64                                               `json:"createAt,omitempty" xml:"createAt,omitempty"`
	Description            *string                                              `json:"description,omitempty" xml:"description,omitempty"`
	GrayConversationIds    []*string                                            `json:"grayConversationIds,omitempty" xml:"grayConversationIds,omitempty" type:"Repeated"`
	GrayInfo               *QueryTemplateInfoResponseBodyGrayInfo               `json:"grayInfo,omitempty" xml:"grayInfo,omitempty" type:"Struct"`
	GrayTemplateId         *string                                              `json:"grayTemplateId,omitempty" xml:"grayTemplateId,omitempty"`
	GroupSettingList       []*QueryTemplateInfoResponseBodyGroupSettingList     `json:"groupSettingList,omitempty" xml:"groupSettingList,omitempty" type:"Repeated"`
	IconMediaId            *string                                              `json:"iconMediaId,omitempty" xml:"iconMediaId,omitempty"`
	ModifiedAt             *int64                                               `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
	ModifyOrderId          *int64                                               `json:"modifyOrderId,omitempty" xml:"modifyOrderId,omitempty"`
	ModifyStatus           *int64                                               `json:"modifyStatus,omitempty" xml:"modifyStatus,omitempty"`
	ParentTemplateDetailVO *QueryTemplateInfoResponseBodyParentTemplateDetailVO `json:"parentTemplateDetailVO,omitempty" xml:"parentTemplateDetailVO,omitempty" type:"Struct"`
	PublishSubState        *string                                              `json:"publishSubState,omitempty" xml:"publishSubState,omitempty"`
	RobotTemplateList      []*string                                            `json:"robotTemplateList,omitempty" xml:"robotTemplateList,omitempty" type:"Repeated"`
	Status                 *int32                                               `json:"status,omitempty" xml:"status,omitempty"`
	TemplateId             *string                                              `json:"templateId,omitempty" xml:"templateId,omitempty"`
	TemplateIntroduction   *QueryTemplateInfoResponseBodyTemplateIntroduction   `json:"templateIntroduction,omitempty" xml:"templateIntroduction,omitempty" type:"Struct"`
	TemplateName           *string                                              `json:"templateName,omitempty" xml:"templateName,omitempty"`
	TemplateType           *int32                                               `json:"templateType,omitempty" xml:"templateType,omitempty"`
	TemplateVisibility     *QueryTemplateInfoResponseBodyTemplateVisibility     `json:"templateVisibility,omitempty" xml:"templateVisibility,omitempty" type:"Struct"`
	ToolbarPluginList      []*string                                            `json:"toolbarPluginList,omitempty" xml:"toolbarPluginList,omitempty" type:"Repeated"`
	Version                *int64                                               `json:"version,omitempty" xml:"version,omitempty"`
}

func (s QueryTemplateInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTemplateInfoResponseBody) SetAbilitySwitch(v int64) *QueryTemplateInfoResponseBody {
	s.AbilitySwitch = &v
	return s
}

func (s *QueryTemplateInfoResponseBody) SetAppInfo(v *QueryTemplateInfoResponseBodyAppInfo) *QueryTemplateInfoResponseBody {
	s.AppInfo = v
	return s
}

func (s *QueryTemplateInfoResponseBody) SetConversationScope(v []*string) *QueryTemplateInfoResponseBody {
	s.ConversationScope = v
	return s
}

func (s *QueryTemplateInfoResponseBody) SetCreateAt(v int64) *QueryTemplateInfoResponseBody {
	s.CreateAt = &v
	return s
}

func (s *QueryTemplateInfoResponseBody) SetDescription(v string) *QueryTemplateInfoResponseBody {
	s.Description = &v
	return s
}

func (s *QueryTemplateInfoResponseBody) SetGrayConversationIds(v []*string) *QueryTemplateInfoResponseBody {
	s.GrayConversationIds = v
	return s
}

func (s *QueryTemplateInfoResponseBody) SetGrayInfo(v *QueryTemplateInfoResponseBodyGrayInfo) *QueryTemplateInfoResponseBody {
	s.GrayInfo = v
	return s
}

func (s *QueryTemplateInfoResponseBody) SetGrayTemplateId(v string) *QueryTemplateInfoResponseBody {
	s.GrayTemplateId = &v
	return s
}

func (s *QueryTemplateInfoResponseBody) SetGroupSettingList(v []*QueryTemplateInfoResponseBodyGroupSettingList) *QueryTemplateInfoResponseBody {
	s.GroupSettingList = v
	return s
}

func (s *QueryTemplateInfoResponseBody) SetIconMediaId(v string) *QueryTemplateInfoResponseBody {
	s.IconMediaId = &v
	return s
}

func (s *QueryTemplateInfoResponseBody) SetModifiedAt(v int64) *QueryTemplateInfoResponseBody {
	s.ModifiedAt = &v
	return s
}

func (s *QueryTemplateInfoResponseBody) SetModifyOrderId(v int64) *QueryTemplateInfoResponseBody {
	s.ModifyOrderId = &v
	return s
}

func (s *QueryTemplateInfoResponseBody) SetModifyStatus(v int64) *QueryTemplateInfoResponseBody {
	s.ModifyStatus = &v
	return s
}

func (s *QueryTemplateInfoResponseBody) SetParentTemplateDetailVO(v *QueryTemplateInfoResponseBodyParentTemplateDetailVO) *QueryTemplateInfoResponseBody {
	s.ParentTemplateDetailVO = v
	return s
}

func (s *QueryTemplateInfoResponseBody) SetPublishSubState(v string) *QueryTemplateInfoResponseBody {
	s.PublishSubState = &v
	return s
}

func (s *QueryTemplateInfoResponseBody) SetRobotTemplateList(v []*string) *QueryTemplateInfoResponseBody {
	s.RobotTemplateList = v
	return s
}

func (s *QueryTemplateInfoResponseBody) SetStatus(v int32) *QueryTemplateInfoResponseBody {
	s.Status = &v
	return s
}

func (s *QueryTemplateInfoResponseBody) SetTemplateId(v string) *QueryTemplateInfoResponseBody {
	s.TemplateId = &v
	return s
}

func (s *QueryTemplateInfoResponseBody) SetTemplateIntroduction(v *QueryTemplateInfoResponseBodyTemplateIntroduction) *QueryTemplateInfoResponseBody {
	s.TemplateIntroduction = v
	return s
}

func (s *QueryTemplateInfoResponseBody) SetTemplateName(v string) *QueryTemplateInfoResponseBody {
	s.TemplateName = &v
	return s
}

func (s *QueryTemplateInfoResponseBody) SetTemplateType(v int32) *QueryTemplateInfoResponseBody {
	s.TemplateType = &v
	return s
}

func (s *QueryTemplateInfoResponseBody) SetTemplateVisibility(v *QueryTemplateInfoResponseBodyTemplateVisibility) *QueryTemplateInfoResponseBody {
	s.TemplateVisibility = v
	return s
}

func (s *QueryTemplateInfoResponseBody) SetToolbarPluginList(v []*string) *QueryTemplateInfoResponseBody {
	s.ToolbarPluginList = v
	return s
}

func (s *QueryTemplateInfoResponseBody) SetVersion(v int64) *QueryTemplateInfoResponseBody {
	s.Version = &v
	return s
}

type QueryTemplateInfoResponseBodyAppInfo struct {
	AppIcon *string `json:"appIcon,omitempty" xml:"appIcon,omitempty"`
	AppId   *string `json:"appId,omitempty" xml:"appId,omitempty"`
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	CorpId  *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s QueryTemplateInfoResponseBodyAppInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateInfoResponseBodyAppInfo) GoString() string {
	return s.String()
}

func (s *QueryTemplateInfoResponseBodyAppInfo) SetAppIcon(v string) *QueryTemplateInfoResponseBodyAppInfo {
	s.AppIcon = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyAppInfo) SetAppId(v string) *QueryTemplateInfoResponseBodyAppInfo {
	s.AppId = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyAppInfo) SetAppName(v string) *QueryTemplateInfoResponseBodyAppInfo {
	s.AppName = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyAppInfo) SetCorpId(v string) *QueryTemplateInfoResponseBodyAppInfo {
	s.CorpId = &v
	return s
}

type QueryTemplateInfoResponseBodyGrayInfo struct {
	TenThousandPercent *int32    `json:"tenThousandPercent,omitempty" xml:"tenThousandPercent,omitempty"`
	WhiteSet           []*string `json:"whiteSet,omitempty" xml:"whiteSet,omitempty" type:"Repeated"`
}

func (s QueryTemplateInfoResponseBodyGrayInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateInfoResponseBodyGrayInfo) GoString() string {
	return s.String()
}

func (s *QueryTemplateInfoResponseBodyGrayInfo) SetTenThousandPercent(v int32) *QueryTemplateInfoResponseBodyGrayInfo {
	s.TenThousandPercent = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyGrayInfo) SetWhiteSet(v []*string) *QueryTemplateInfoResponseBodyGrayInfo {
	s.WhiteSet = v
	return s
}

type QueryTemplateInfoResponseBodyGroupSettingList struct {
	Desc  *string `json:"desc,omitempty" xml:"desc,omitempty"`
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	State *bool   `json:"state,omitempty" xml:"state,omitempty"`
}

func (s QueryTemplateInfoResponseBodyGroupSettingList) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateInfoResponseBodyGroupSettingList) GoString() string {
	return s.String()
}

func (s *QueryTemplateInfoResponseBodyGroupSettingList) SetDesc(v string) *QueryTemplateInfoResponseBodyGroupSettingList {
	s.Desc = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyGroupSettingList) SetName(v string) *QueryTemplateInfoResponseBodyGroupSettingList {
	s.Name = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyGroupSettingList) SetState(v bool) *QueryTemplateInfoResponseBodyGroupSettingList {
	s.State = &v
	return s
}

type QueryTemplateInfoResponseBodyParentTemplateDetailVO struct {
	RobotTemplateList []*QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList `json:"robotTemplateList,omitempty" xml:"robotTemplateList,omitempty" type:"Repeated"`
	TemplateId        *string                                                                 `json:"templateId,omitempty" xml:"templateId,omitempty"`
	ToolbarPluginList []*QueryTemplateInfoResponseBodyParentTemplateDetailVOToolbarPluginList `json:"toolbarPluginList,omitempty" xml:"toolbarPluginList,omitempty" type:"Repeated"`
}

func (s QueryTemplateInfoResponseBodyParentTemplateDetailVO) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateInfoResponseBodyParentTemplateDetailVO) GoString() string {
	return s.String()
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVO) SetRobotTemplateList(v []*QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList) *QueryTemplateInfoResponseBodyParentTemplateDetailVO {
	s.RobotTemplateList = v
	return s
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVO) SetTemplateId(v string) *QueryTemplateInfoResponseBodyParentTemplateDetailVO {
	s.TemplateId = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVO) SetToolbarPluginList(v []*QueryTemplateInfoResponseBodyParentTemplateDetailVOToolbarPluginList) *QueryTemplateInfoResponseBodyParentTemplateDetailVO {
	s.ToolbarPluginList = v
	return s
}

type QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList struct {
	Brief           *string `json:"brief,omitempty" xml:"brief,omitempty"`
	Code            *string `json:"code,omitempty" xml:"code,omitempty"`
	CorpId          *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	CreateAt        *int64  `json:"createAt,omitempty" xml:"createAt,omitempty"`
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	Dev             *string `json:"dev,omitempty" xml:"dev,omitempty"`
	GroupTemplateId *string `json:"groupTemplateId,omitempty" xml:"groupTemplateId,omitempty"`
	Icon            *string `json:"icon,omitempty" xml:"icon,omitempty"`
	ModifiedAt      *int64  `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	OutgoingToken   *string `json:"outgoingToken,omitempty" xml:"outgoingToken,omitempty"`
	OutgoingUrl     *string `json:"outgoingUrl,omitempty" xml:"outgoingUrl,omitempty"`
	PreviewMediaId  *string `json:"previewMediaId,omitempty" xml:"previewMediaId,omitempty"`
	SourceUrl       *string `json:"sourceUrl,omitempty" xml:"sourceUrl,omitempty"`
	Status          *int32  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList) GoString() string {
	return s.String()
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList) SetBrief(v string) *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList {
	s.Brief = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList) SetCode(v string) *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList {
	s.Code = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList) SetCorpId(v string) *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList {
	s.CorpId = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList) SetCreateAt(v int64) *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList {
	s.CreateAt = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList) SetDescription(v string) *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList {
	s.Description = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList) SetDev(v string) *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList {
	s.Dev = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList) SetGroupTemplateId(v string) *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList {
	s.GroupTemplateId = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList) SetIcon(v string) *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList {
	s.Icon = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList) SetModifiedAt(v int64) *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList {
	s.ModifiedAt = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList) SetName(v string) *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList {
	s.Name = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList) SetOutgoingToken(v string) *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList {
	s.OutgoingToken = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList) SetOutgoingUrl(v string) *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList {
	s.OutgoingUrl = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList) SetPreviewMediaId(v string) *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList {
	s.PreviewMediaId = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList) SetSourceUrl(v string) *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList {
	s.SourceUrl = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList) SetStatus(v int32) *QueryTemplateInfoResponseBodyParentTemplateDetailVORobotTemplateList {
	s.Status = &v
	return s
}

type QueryTemplateInfoResponseBodyParentTemplateDetailVOToolbarPluginList struct {
	AppId      *string `json:"appId,omitempty" xml:"appId,omitempty"`
	CreateAt   *int64  `json:"createAt,omitempty" xml:"createAt,omitempty"`
	Desc       *string `json:"desc,omitempty" xml:"desc,omitempty"`
	Icons      *string `json:"icons,omitempty" xml:"icons,omitempty"`
	ModifiedAt *int64  `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	PcUrl      *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	PluginId   *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	Status     *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url        *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s QueryTemplateInfoResponseBodyParentTemplateDetailVOToolbarPluginList) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateInfoResponseBodyParentTemplateDetailVOToolbarPluginList) GoString() string {
	return s.String()
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVOToolbarPluginList) SetAppId(v string) *QueryTemplateInfoResponseBodyParentTemplateDetailVOToolbarPluginList {
	s.AppId = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVOToolbarPluginList) SetCreateAt(v int64) *QueryTemplateInfoResponseBodyParentTemplateDetailVOToolbarPluginList {
	s.CreateAt = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVOToolbarPluginList) SetDesc(v string) *QueryTemplateInfoResponseBodyParentTemplateDetailVOToolbarPluginList {
	s.Desc = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVOToolbarPluginList) SetIcons(v string) *QueryTemplateInfoResponseBodyParentTemplateDetailVOToolbarPluginList {
	s.Icons = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVOToolbarPluginList) SetModifiedAt(v int64) *QueryTemplateInfoResponseBodyParentTemplateDetailVOToolbarPluginList {
	s.ModifiedAt = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVOToolbarPluginList) SetName(v string) *QueryTemplateInfoResponseBodyParentTemplateDetailVOToolbarPluginList {
	s.Name = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVOToolbarPluginList) SetPcUrl(v string) *QueryTemplateInfoResponseBodyParentTemplateDetailVOToolbarPluginList {
	s.PcUrl = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVOToolbarPluginList) SetPluginId(v string) *QueryTemplateInfoResponseBodyParentTemplateDetailVOToolbarPluginList {
	s.PluginId = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVOToolbarPluginList) SetStatus(v int32) *QueryTemplateInfoResponseBodyParentTemplateDetailVOToolbarPluginList {
	s.Status = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyParentTemplateDetailVOToolbarPluginList) SetUrl(v string) *QueryTemplateInfoResponseBodyParentTemplateDetailVOToolbarPluginList {
	s.Url = &v
	return s
}

type QueryTemplateInfoResponseBodyTemplateIntroduction struct {
	Banner *string `json:"banner,omitempty" xml:"banner,omitempty"`
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	Title  *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryTemplateInfoResponseBodyTemplateIntroduction) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateInfoResponseBodyTemplateIntroduction) GoString() string {
	return s.String()
}

func (s *QueryTemplateInfoResponseBodyTemplateIntroduction) SetBanner(v string) *QueryTemplateInfoResponseBodyTemplateIntroduction {
	s.Banner = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyTemplateIntroduction) SetDetail(v string) *QueryTemplateInfoResponseBodyTemplateIntroduction {
	s.Detail = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyTemplateIntroduction) SetTitle(v string) *QueryTemplateInfoResponseBodyTemplateIntroduction {
	s.Title = &v
	return s
}

type QueryTemplateInfoResponseBodyTemplateVisibility struct {
	CorpId  *string                                                   `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeptIds []*QueryTemplateInfoResponseBodyTemplateVisibilityDeptIds `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	RoleIds []*string                                                 `json:"roleIds,omitempty" xml:"roleIds,omitempty" type:"Repeated"`
	UserIds []*QueryTemplateInfoResponseBodyTemplateVisibilityUserIds `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s QueryTemplateInfoResponseBodyTemplateVisibility) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateInfoResponseBodyTemplateVisibility) GoString() string {
	return s.String()
}

func (s *QueryTemplateInfoResponseBodyTemplateVisibility) SetCorpId(v string) *QueryTemplateInfoResponseBodyTemplateVisibility {
	s.CorpId = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyTemplateVisibility) SetDeptIds(v []*QueryTemplateInfoResponseBodyTemplateVisibilityDeptIds) *QueryTemplateInfoResponseBodyTemplateVisibility {
	s.DeptIds = v
	return s
}

func (s *QueryTemplateInfoResponseBodyTemplateVisibility) SetRoleIds(v []*string) *QueryTemplateInfoResponseBodyTemplateVisibility {
	s.RoleIds = v
	return s
}

func (s *QueryTemplateInfoResponseBodyTemplateVisibility) SetUserIds(v []*QueryTemplateInfoResponseBodyTemplateVisibilityUserIds) *QueryTemplateInfoResponseBodyTemplateVisibility {
	s.UserIds = v
	return s
}

type QueryTemplateInfoResponseBodyTemplateVisibilityDeptIds struct {
	DeptId   *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
}

func (s QueryTemplateInfoResponseBodyTemplateVisibilityDeptIds) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateInfoResponseBodyTemplateVisibilityDeptIds) GoString() string {
	return s.String()
}

func (s *QueryTemplateInfoResponseBodyTemplateVisibilityDeptIds) SetDeptId(v string) *QueryTemplateInfoResponseBodyTemplateVisibilityDeptIds {
	s.DeptId = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyTemplateVisibilityDeptIds) SetDeptName(v string) *QueryTemplateInfoResponseBodyTemplateVisibilityDeptIds {
	s.DeptName = &v
	return s
}

type QueryTemplateInfoResponseBodyTemplateVisibilityUserIds struct {
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	Nick   *string `json:"nick,omitempty" xml:"nick,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryTemplateInfoResponseBodyTemplateVisibilityUserIds) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateInfoResponseBodyTemplateVisibilityUserIds) GoString() string {
	return s.String()
}

func (s *QueryTemplateInfoResponseBodyTemplateVisibilityUserIds) SetAvatar(v string) *QueryTemplateInfoResponseBodyTemplateVisibilityUserIds {
	s.Avatar = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyTemplateVisibilityUserIds) SetNick(v string) *QueryTemplateInfoResponseBodyTemplateVisibilityUserIds {
	s.Nick = &v
	return s
}

func (s *QueryTemplateInfoResponseBodyTemplateVisibilityUserIds) SetUserId(v string) *QueryTemplateInfoResponseBodyTemplateVisibilityUserIds {
	s.UserId = &v
	return s
}

type QueryTemplateInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTemplateInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTemplateInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTemplateInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryTemplateInfoResponse) SetHeaders(v map[string]*string) *QueryTemplateInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryTemplateInfoResponse) SetStatusCode(v int32) *QueryTemplateInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTemplateInfoResponse) SetBody(v *QueryTemplateInfoResponseBody) *QueryTemplateInfoResponse {
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
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	Platform  *int32 `json:"platform,omitempty" xml:"platform,omitempty"`
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// This parameter is required.
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 12034-1233
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserBehaviorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 5000000003
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	// example:
	//
	// 0.0.5
	RollbackVersion *string `json:"rollbackVersion,omitempty" xml:"rollbackVersion,omitempty"`
	// example:
	//
	// 0.0.4
	TargetVersion *string `json:"targetVersion,omitempty" xml:"targetVersion,omitempty"`
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
	// example:
	//
	// 成功
	Cause *string `json:"cause,omitempty" xml:"cause,omitempty"`
	// example:
	//
	// 200
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackMiniAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type RuleBatchReceiverHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RuleBatchReceiverHeaders) String() string {
	return tea.Prettify(s)
}

func (s RuleBatchReceiverHeaders) GoString() string {
	return s.String()
}

func (s *RuleBatchReceiverHeaders) SetCommonHeaders(v map[string]*string) *RuleBatchReceiverHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RuleBatchReceiverHeaders) SetXAcsDingtalkAccessToken(v string) *RuleBatchReceiverHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RuleBatchReceiverRequest struct {
	BatchNo         *string                         `json:"batchNo,omitempty" xml:"batchNo,omitempty"`
	CardOptions     *string                         `json:"cardOptions,omitempty" xml:"cardOptions,omitempty"`
	Data            []*RuleBatchReceiverRequestData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	RuleCode        *string                         `json:"ruleCode,omitempty" xml:"ruleCode,omitempty"`
	SecretKey       *string                         `json:"secretKey,omitempty" xml:"secretKey,omitempty"`
	SpecialStrategy *bool                           `json:"specialStrategy,omitempty" xml:"specialStrategy,omitempty"`
	TaskBatchNo     *string                         `json:"taskBatchNo,omitempty" xml:"taskBatchNo,omitempty"`
}

func (s RuleBatchReceiverRequest) String() string {
	return tea.Prettify(s)
}

func (s RuleBatchReceiverRequest) GoString() string {
	return s.String()
}

func (s *RuleBatchReceiverRequest) SetBatchNo(v string) *RuleBatchReceiverRequest {
	s.BatchNo = &v
	return s
}

func (s *RuleBatchReceiverRequest) SetCardOptions(v string) *RuleBatchReceiverRequest {
	s.CardOptions = &v
	return s
}

func (s *RuleBatchReceiverRequest) SetData(v []*RuleBatchReceiverRequestData) *RuleBatchReceiverRequest {
	s.Data = v
	return s
}

func (s *RuleBatchReceiverRequest) SetRuleCode(v string) *RuleBatchReceiverRequest {
	s.RuleCode = &v
	return s
}

func (s *RuleBatchReceiverRequest) SetSecretKey(v string) *RuleBatchReceiverRequest {
	s.SecretKey = &v
	return s
}

func (s *RuleBatchReceiverRequest) SetSpecialStrategy(v bool) *RuleBatchReceiverRequest {
	s.SpecialStrategy = &v
	return s
}

func (s *RuleBatchReceiverRequest) SetTaskBatchNo(v string) *RuleBatchReceiverRequest {
	s.TaskBatchNo = &v
	return s
}

type RuleBatchReceiverRequestData struct {
	AtAccount       *string                            `json:"atAccount,omitempty" xml:"atAccount,omitempty"`
	Attrs           *RuleBatchReceiverRequestDataAttrs `json:"attrs,omitempty" xml:"attrs,omitempty" type:"Struct"`
	CallbackUrl     *string                            `json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`
	CardCallbackUrl *string                            `json:"cardCallbackUrl,omitempty" xml:"cardCallbackUrl,omitempty"`
	Content         map[string]map[string]interface{}  `json:"content,omitempty" xml:"content,omitempty"`
	IsAtAll         *bool                              `json:"isAtAll,omitempty" xml:"isAtAll,omitempty"`
	ReceiverAccount *string                            `json:"receiverAccount,omitempty" xml:"receiverAccount,omitempty"`
	ReceiverType    *int32                             `json:"receiverType,omitempty" xml:"receiverType,omitempty"`
	SerialNumber    *string                            `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
}

func (s RuleBatchReceiverRequestData) String() string {
	return tea.Prettify(s)
}

func (s RuleBatchReceiverRequestData) GoString() string {
	return s.String()
}

func (s *RuleBatchReceiverRequestData) SetAtAccount(v string) *RuleBatchReceiverRequestData {
	s.AtAccount = &v
	return s
}

func (s *RuleBatchReceiverRequestData) SetAttrs(v *RuleBatchReceiverRequestDataAttrs) *RuleBatchReceiverRequestData {
	s.Attrs = v
	return s
}

func (s *RuleBatchReceiverRequestData) SetCallbackUrl(v string) *RuleBatchReceiverRequestData {
	s.CallbackUrl = &v
	return s
}

func (s *RuleBatchReceiverRequestData) SetCardCallbackUrl(v string) *RuleBatchReceiverRequestData {
	s.CardCallbackUrl = &v
	return s
}

func (s *RuleBatchReceiverRequestData) SetContent(v map[string]map[string]interface{}) *RuleBatchReceiverRequestData {
	s.Content = v
	return s
}

func (s *RuleBatchReceiverRequestData) SetIsAtAll(v bool) *RuleBatchReceiverRequestData {
	s.IsAtAll = &v
	return s
}

func (s *RuleBatchReceiverRequestData) SetReceiverAccount(v string) *RuleBatchReceiverRequestData {
	s.ReceiverAccount = &v
	return s
}

func (s *RuleBatchReceiverRequestData) SetReceiverType(v int32) *RuleBatchReceiverRequestData {
	s.ReceiverType = &v
	return s
}

func (s *RuleBatchReceiverRequestData) SetSerialNumber(v string) *RuleBatchReceiverRequestData {
	s.SerialNumber = &v
	return s
}

type RuleBatchReceiverRequestDataAttrs struct {
	ListUnitId []*int64 `json:"listUnitId,omitempty" xml:"listUnitId,omitempty" type:"Repeated"`
}

func (s RuleBatchReceiverRequestDataAttrs) String() string {
	return tea.Prettify(s)
}

func (s RuleBatchReceiverRequestDataAttrs) GoString() string {
	return s.String()
}

func (s *RuleBatchReceiverRequestDataAttrs) SetListUnitId(v []*int64) *RuleBatchReceiverRequestDataAttrs {
	s.ListUnitId = v
	return s
}

type RuleBatchReceiverResponseBody struct {
	Code  *int32                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data  []*RuleBatchReceiverResponseBodyData   `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Msg   *string                                `json:"msg,omitempty" xml:"msg,omitempty"`
	MsgId *string                                `json:"msgId,omitempty" xml:"msgId,omitempty"`
	Rows  [][]*RuleBatchReceiverResponseBodyRows `json:"rows,omitempty" xml:"rows,omitempty" type:"Repeated"`
}

func (s RuleBatchReceiverResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RuleBatchReceiverResponseBody) GoString() string {
	return s.String()
}

func (s *RuleBatchReceiverResponseBody) SetCode(v int32) *RuleBatchReceiverResponseBody {
	s.Code = &v
	return s
}

func (s *RuleBatchReceiverResponseBody) SetData(v []*RuleBatchReceiverResponseBodyData) *RuleBatchReceiverResponseBody {
	s.Data = v
	return s
}

func (s *RuleBatchReceiverResponseBody) SetMsg(v string) *RuleBatchReceiverResponseBody {
	s.Msg = &v
	return s
}

func (s *RuleBatchReceiverResponseBody) SetMsgId(v string) *RuleBatchReceiverResponseBody {
	s.MsgId = &v
	return s
}

func (s *RuleBatchReceiverResponseBody) SetRows(v [][]*RuleBatchReceiverResponseBodyRows) *RuleBatchReceiverResponseBody {
	s.Rows = v
	return s
}

type RuleBatchReceiverResponseBodyData struct {
	MsgId        *string `json:"msgId,omitempty" xml:"msgId,omitempty"`
	SerialNumber *string `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
}

func (s RuleBatchReceiverResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RuleBatchReceiverResponseBodyData) GoString() string {
	return s.String()
}

func (s *RuleBatchReceiverResponseBodyData) SetMsgId(v string) *RuleBatchReceiverResponseBodyData {
	s.MsgId = &v
	return s
}

func (s *RuleBatchReceiverResponseBodyData) SetSerialNumber(v string) *RuleBatchReceiverResponseBodyData {
	s.SerialNumber = &v
	return s
}

type RuleBatchReceiverResponseBodyRows struct {
	SerialNumber *string `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
	MsgId        *string `json:"msgId,omitempty" xml:"msgId,omitempty"`
}

func (s RuleBatchReceiverResponseBodyRows) String() string {
	return tea.Prettify(s)
}

func (s RuleBatchReceiverResponseBodyRows) GoString() string {
	return s.String()
}

func (s *RuleBatchReceiverResponseBodyRows) SetSerialNumber(v string) *RuleBatchReceiverResponseBodyRows {
	s.SerialNumber = &v
	return s
}

func (s *RuleBatchReceiverResponseBodyRows) SetMsgId(v string) *RuleBatchReceiverResponseBodyRows {
	s.MsgId = &v
	return s
}

type RuleBatchReceiverResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RuleBatchReceiverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RuleBatchReceiverResponse) String() string {
	return tea.Prettify(s)
}

func (s RuleBatchReceiverResponse) GoString() string {
	return s.String()
}

func (s *RuleBatchReceiverResponse) SetHeaders(v map[string]*string) *RuleBatchReceiverResponse {
	s.Headers = v
	return s
}

func (s *RuleBatchReceiverResponse) SetStatusCode(v int32) *RuleBatchReceiverResponse {
	s.StatusCode = &v
	return s
}

func (s *RuleBatchReceiverResponse) SetBody(v *RuleBatchReceiverResponseBody) *RuleBatchReceiverResponse {
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
	// example:
	//
	// sampleKeyId1234
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// example:
	//
	// sampleSecretId1234
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	// example:
	//
	// xxxxbucket
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CloudType *int32 `json:"cloudType,omitempty" xml:"cloudType,omitempty"`
	// example:
	//
	// https://oss-cn-test.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dingxxxxxxxxxxxx
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
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
	// example:
	//
	// sampleKeyId1234
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://oss-cn-test.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	State *int32 `json:"state,omitempty" xml:"state,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveAcrossCloudStroageConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// XXX组织申请高级认证
	ApplyRemark *string `json:"applyRemark,omitempty" xml:"applyRemark,omitempty"`
	// example:
	//
	// @lQLxxxxxxxxVvjg8zImwm6t1BYIUNv0Cas0x7UA-AA
	AuthorizeMediaId *string `json:"authorizeMediaId,omitempty" xml:"authorizeMediaId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 计算机
	Industry *string `json:"industry,omitempty" xml:"industry,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 钉三多
	LegalPerson *string `json:"legalPerson,omitempty" xml:"legalPerson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3301XX1997XXXXXXXXX
	LegalPersonIdCard *string `json:"legalPersonIdCard,omitempty" xml:"legalPersonIdCard,omitempty"`
	// example:
	//
	// @lQLxxxxxxxxVvjg8zImwm6t1BYIUNv0Cas0x7UA-AA
	LicenseMediaId *string `json:"licenseMediaId,omitempty" xml:"licenseMediaId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 330100
	LocCity *int64 `json:"locCity,omitempty" xml:"locCity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 杭州
	LocCityName *string `json:"locCityName,omitempty" xml:"locCityName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 330000
	LocProvince *int64 `json:"locProvince,omitempty" xml:"locProvince,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 浙江
	LocProvinceName *string `json:"locProvinceName,omitempty" xml:"locProvinceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 15869110714
	MobileNum *string `json:"mobileNum,omitempty" xml:"mobileNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试组织
	OrgName *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	// example:
	//
	// 11111111-1
	OrganizationCode *string `json:"organizationCode,omitempty" xml:"organizationCode,omitempty"`
	// example:
	//
	// @lQLxxxxxxxxVvjg8zImwm6t1BYIUNv0Cas0x7UA-AA
	OrganizationCodeMediaId *string `json:"organizationCodeMediaId,omitempty" xml:"organizationCodeMediaId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 余杭区文一西路XX号
	RegistLocation *string `json:"registLocation,omitempty" xml:"registLocation,omitempty"`
	// example:
	//
	// 1111111111111111
	RegistNum *string `json:"registNum,omitempty" xml:"registNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dingxxxxxxxxxxxx
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
	// example:
	//
	// 9111111XX2957XX4X
	UnifiedSocialCredit *string `json:"unifiedSocialCredit,omitempty" xml:"unifiedSocialCredit,omitempty"`
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
	// example:
	//
	// true
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveAndSubmitAuthInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// dingf8d907412a586
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 云枢
	LogSource *string `json:"logSource,omitempty" xml:"logSource,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// terminalInfo
	LogType *string `json:"logType,omitempty" xml:"logType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"date":"2023-05-10","macAddress":"34-2E-B7-AF-EA-JF","devSn":"68D1F0-B76A-5CC9-BCFC-BD7548BA","staffId":"05166141678164"}]
	OpenExt *string `json:"openExt,omitempty" xml:"openExt,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveOpenTerminalInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type SaveStorageFunctionSwitchHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SaveStorageFunctionSwitchHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveStorageFunctionSwitchHeaders) GoString() string {
	return s.String()
}

func (s *SaveStorageFunctionSwitchHeaders) SetCommonHeaders(v map[string]*string) *SaveStorageFunctionSwitchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveStorageFunctionSwitchHeaders) SetXAcsDingtalkAccessToken(v string) *SaveStorageFunctionSwitchHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SaveStorageFunctionSwitchRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 15800000000
	FunctionList []*SaveStorageFunctionSwitchRequestFunctionList `json:"functionList,omitempty" xml:"functionList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// dingxxxxxx
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
}

func (s SaveStorageFunctionSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveStorageFunctionSwitchRequest) GoString() string {
	return s.String()
}

func (s *SaveStorageFunctionSwitchRequest) SetFunctionList(v []*SaveStorageFunctionSwitchRequestFunctionList) *SaveStorageFunctionSwitchRequest {
	s.FunctionList = v
	return s
}

func (s *SaveStorageFunctionSwitchRequest) SetTargetCorpId(v string) *SaveStorageFunctionSwitchRequest {
	s.TargetCorpId = &v
	return s
}

type SaveStorageFunctionSwitchRequestFunctionList struct {
	// This parameter is required.
	FunctionKey *string `json:"functionKey,omitempty" xml:"functionKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	FunctionValue *string `json:"functionValue,omitempty" xml:"functionValue,omitempty"`
}

func (s SaveStorageFunctionSwitchRequestFunctionList) String() string {
	return tea.Prettify(s)
}

func (s SaveStorageFunctionSwitchRequestFunctionList) GoString() string {
	return s.String()
}

func (s *SaveStorageFunctionSwitchRequestFunctionList) SetFunctionKey(v string) *SaveStorageFunctionSwitchRequestFunctionList {
	s.FunctionKey = &v
	return s
}

func (s *SaveStorageFunctionSwitchRequestFunctionList) SetFunctionValue(v string) *SaveStorageFunctionSwitchRequestFunctionList {
	s.FunctionValue = &v
	return s
}

type SaveStorageFunctionSwitchResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SaveStorageFunctionSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveStorageFunctionSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *SaveStorageFunctionSwitchResponseBody) SetSuccess(v bool) *SaveStorageFunctionSwitchResponseBody {
	s.Success = &v
	return s
}

type SaveStorageFunctionSwitchResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveStorageFunctionSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveStorageFunctionSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveStorageFunctionSwitchResponse) GoString() string {
	return s.String()
}

func (s *SaveStorageFunctionSwitchResponse) SetHeaders(v map[string]*string) *SaveStorageFunctionSwitchResponse {
	s.Headers = v
	return s
}

func (s *SaveStorageFunctionSwitchResponse) SetStatusCode(v int32) *SaveStorageFunctionSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveStorageFunctionSwitchResponse) SetBody(v *SaveStorageFunctionSwitchResponseBody) *SaveStorageFunctionSwitchResponse {
	s.Body = v
	return s
}

type SaveStorageSwitchHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SaveStorageSwitchHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveStorageSwitchHeaders) GoString() string {
	return s.String()
}

func (s *SaveStorageSwitchHeaders) SetCommonHeaders(v map[string]*string) *SaveStorageSwitchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveStorageSwitchHeaders) SetXAcsDingtalkAccessToken(v string) *SaveStorageSwitchHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SaveStorageSwitchRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0-关闭，1-开启
	OpenStorage *int32 `json:"openStorage,omitempty" xml:"openStorage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dingxxxxxxxxx
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
}

func (s SaveStorageSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveStorageSwitchRequest) GoString() string {
	return s.String()
}

func (s *SaveStorageSwitchRequest) SetOpenStorage(v int32) *SaveStorageSwitchRequest {
	s.OpenStorage = &v
	return s
}

func (s *SaveStorageSwitchRequest) SetTargetCorpId(v string) *SaveStorageSwitchRequest {
	s.TargetCorpId = &v
	return s
}

type SaveStorageSwitchResponseBody struct {
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SaveStorageSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveStorageSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *SaveStorageSwitchResponseBody) SetSuccess(v bool) *SaveStorageSwitchResponseBody {
	s.Success = &v
	return s
}

type SaveStorageSwitchResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveStorageSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveStorageSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveStorageSwitchResponse) GoString() string {
	return s.String()
}

func (s *SaveStorageSwitchResponse) SetHeaders(v map[string]*string) *SaveStorageSwitchResponse {
	s.Headers = v
	return s
}

func (s *SaveStorageSwitchResponse) SetStatusCode(v int32) *SaveStorageSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveStorageSwitchResponse) SetBody(v *SaveStorageSwitchResponseBody) *SaveStorageSwitchResponse {
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
	// Deprecated
	AgentIdList []*int64 `json:"agentIdList,omitempty" xml:"agentIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// {"openShareControl":[123],"openClipboardPaste":[123]}
	AgentIdMap *string `json:"agentIdMap,omitempty" xml:"agentIdMap,omitempty"`
	// Deprecated
	//
	// example:
	//
	// add
	Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
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

func (s *SaveWhiteAppRequest) SetAgentIdMap(v string) *SaveWhiteAppRequest {
	s.AgentIdMap = &v
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveWhiteAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 创建时间查询最大时间戳
	CreateTimeEnd *int64 `json:"createTimeEnd,omitempty" xml:"createTimeEnd,omitempty"`
	// example:
	//
	// 创建时间查询最小时间戳
	CreateTimeStart *int64 `json:"createTimeStart,omitempty" xml:"createTimeStart,omitempty"`
	// example:
	//
	// 群人数范围最大值，例如100
	GroupMembersCountEnd *int32 `json:"groupMembersCountEnd,omitempty" xml:"groupMembersCountEnd,omitempty"`
	// example:
	//
	// 群人数范围最小值，例如1
	GroupMembersCountStart *int32 `json:"groupMembersCountStart,omitempty" xml:"groupMembersCountStart,omitempty"`
	// example:
	//
	// 群名称
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// example:
	//
	// 群主userId
	GroupOwner *string `json:"groupOwner,omitempty" xml:"groupOwner,omitempty"`
	// example:
	//
	// 最后一次活跃时间戳最大值
	LastActiveTimeEnd *int64 `json:"lastActiveTimeEnd,omitempty" xml:"lastActiveTimeEnd,omitempty"`
	// example:
	//
	// 最后一次活跃时间戳最小值
	LastActiveTimeStart *int64 `json:"lastActiveTimeStart,omitempty" xml:"lastActiveTimeStart,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 当前查询人的userId
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 分页大小，最大不能超过100
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 分页号，从1开始
	PageStart *int32 `json:"pageStart,omitempty" xml:"pageStart,omitempty"`
	// example:
	//
	// 是否同步到钉盘 0不同步 1同步
	SyncToDingpan *int32 `json:"syncToDingpan,omitempty" xml:"syncToDingpan,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 每次查询唯一标识，保证每次分页查询时该值不变
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
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
	// This parameter is required.
	ItemCount *int32 `json:"itemCount,omitempty" xml:"itemCount,omitempty"`
	// This parameter is required.
	Items []*SearchOrgInnerGroupInfoResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// This parameter is required.
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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
	Extensions map[string]*string `json:"extensions,omitempty" xml:"extensions,omitempty"`
	// This parameter is required.
	GroupAdminsCount *int32 `json:"groupAdminsCount,omitempty" xml:"groupAdminsCount,omitempty"`
	// This parameter is required.
	GroupCreateTime *int64 `json:"groupCreateTime,omitempty" xml:"groupCreateTime,omitempty"`
	// This parameter is required.
	GroupLastActiveTime *int64 `json:"groupLastActiveTime,omitempty" xml:"groupLastActiveTime,omitempty"`
	// This parameter is required.
	GroupLastActiveTimeShow *string `json:"groupLastActiveTimeShow,omitempty" xml:"groupLastActiveTimeShow,omitempty"`
	// This parameter is required.
	GroupMembersCount *int32 `json:"groupMembersCount,omitempty" xml:"groupMembersCount,omitempty"`
	// This parameter is required.
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// This parameter is required.
	GroupOwner *string `json:"groupOwner,omitempty" xml:"groupOwner,omitempty"`
	// This parameter is required.
	GroupOwnerUserId *string `json:"groupOwnerUserId,omitempty" xml:"groupOwnerUserId,omitempty"`
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	SyncToDingpan *int32 `json:"syncToDingpan,omitempty" xml:"syncToDingpan,omitempty"`
	// This parameter is required.
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// This parameter is required.
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// This parameter is required.
	UsedQuota *int64 `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
}

func (s SearchOrgInnerGroupInfoResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s SearchOrgInnerGroupInfoResponseBodyItems) GoString() string {
	return s.String()
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetExtensions(v map[string]*string) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.Extensions = v
	return s
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchOrgInnerGroupInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 开会
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 1234
	DeptId *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 钉开放组织别名
	OrgAlias *string `json:"orgAlias,omitempty" xml:"orgAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234
	PartnerLabelId *int64 `json:"partnerLabelId,omitempty" xml:"partnerLabelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234
	PartnerNum *string `json:"partnerNum,omitempty" xml:"partnerNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 133XXXXXX57
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 开会
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
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
	// example:
	//
	// true
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendPhoneDingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type SetConversationCategoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SetConversationCategoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetConversationCategoryHeaders) GoString() string {
	return s.String()
}

func (s *SetConversationCategoryHeaders) SetCommonHeaders(v map[string]*string) *SetConversationCategoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetConversationCategoryHeaders) SetXAcsDingtalkAccessToken(v string) *SetConversationCategoryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SetConversationCategoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// -1
	CategoryId *int64 `json:"categoryId,omitempty" xml:"categoryId,omitempty"`
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s SetConversationCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s SetConversationCategoryRequest) GoString() string {
	return s.String()
}

func (s *SetConversationCategoryRequest) SetCategoryId(v int64) *SetConversationCategoryRequest {
	s.CategoryId = &v
	return s
}

func (s *SetConversationCategoryRequest) SetOpenConversationId(v string) *SetConversationCategoryRequest {
	s.OpenConversationId = &v
	return s
}

type SetConversationCategoryResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SetConversationCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetConversationCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *SetConversationCategoryResponseBody) SetResult(v string) *SetConversationCategoryResponseBody {
	s.Result = &v
	return s
}

func (s *SetConversationCategoryResponseBody) SetSuccess(v bool) *SetConversationCategoryResponseBody {
	s.Success = &v
	return s
}

type SetConversationCategoryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetConversationCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetConversationCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s SetConversationCategoryResponse) GoString() string {
	return s.String()
}

func (s *SetConversationCategoryResponse) SetHeaders(v map[string]*string) *SetConversationCategoryResponse {
	s.Headers = v
	return s
}

func (s *SetConversationCategoryResponse) SetStatusCode(v int32) *SetConversationCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *SetConversationCategoryResponse) SetBody(v *SetConversationCategoryResponseBody) *SetConversationCategoryResponse {
	s.Body = v
	return s
}

type SetConversationSubtitleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SetConversationSubtitleHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetConversationSubtitleHeaders) GoString() string {
	return s.String()
}

func (s *SetConversationSubtitleHeaders) SetCommonHeaders(v map[string]*string) *SetConversationSubtitleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetConversationSubtitleHeaders) SetXAcsDingtalkAccessToken(v string) *SetConversationSubtitleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SetConversationSubtitleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cidxxxx
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// 副标题
	Subtitle *string `json:"subtitle,omitempty" xml:"subtitle,omitempty"`
	// example:
	//
	// #0000FF
	SubtitleColor *string `json:"subtitleColor,omitempty" xml:"subtitleColor,omitempty"`
}

func (s SetConversationSubtitleRequest) String() string {
	return tea.Prettify(s)
}

func (s SetConversationSubtitleRequest) GoString() string {
	return s.String()
}

func (s *SetConversationSubtitleRequest) SetOpenConversationId(v string) *SetConversationSubtitleRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SetConversationSubtitleRequest) SetSubtitle(v string) *SetConversationSubtitleRequest {
	s.Subtitle = &v
	return s
}

func (s *SetConversationSubtitleRequest) SetSubtitleColor(v string) *SetConversationSubtitleRequest {
	s.SubtitleColor = &v
	return s
}

type SetConversationSubtitleResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SetConversationSubtitleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetConversationSubtitleResponseBody) GoString() string {
	return s.String()
}

func (s *SetConversationSubtitleResponseBody) SetSuccess(v bool) *SetConversationSubtitleResponseBody {
	s.Success = &v
	return s
}

type SetConversationSubtitleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetConversationSubtitleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetConversationSubtitleResponse) String() string {
	return tea.Prettify(s)
}

func (s SetConversationSubtitleResponse) GoString() string {
	return s.String()
}

func (s *SetConversationSubtitleResponse) SetHeaders(v map[string]*string) *SetConversationSubtitleResponse {
	s.Headers = v
	return s
}

func (s *SetConversationSubtitleResponse) SetStatusCode(v int32) *SetConversationSubtitleResponse {
	s.StatusCode = &v
	return s
}

func (s *SetConversationSubtitleResponse) SetBody(v *SetConversationSubtitleResponseBody) *SetConversationSubtitleResponse {
	s.Body = v
	return s
}

type SetConversationTopCategoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SetConversationTopCategoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetConversationTopCategoryHeaders) GoString() string {
	return s.String()
}

func (s *SetConversationTopCategoryHeaders) SetCommonHeaders(v map[string]*string) *SetConversationTopCategoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetConversationTopCategoryHeaders) SetXAcsDingtalkAccessToken(v string) *SetConversationTopCategoryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SetConversationTopCategoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cidxxxx
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// if can be null:
	// true
	SetCategoryList []*SetConversationTopCategoryRequestSetCategoryList `json:"setCategoryList,omitempty" xml:"setCategoryList,omitempty" type:"Repeated"`
	// if can be null:
	// true
	//
	// example:
	//
	// 10
	Sign *int32 `json:"sign,omitempty" xml:"sign,omitempty"`
}

func (s SetConversationTopCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s SetConversationTopCategoryRequest) GoString() string {
	return s.String()
}

func (s *SetConversationTopCategoryRequest) SetOpenConversationId(v string) *SetConversationTopCategoryRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SetConversationTopCategoryRequest) SetSetCategoryList(v []*SetConversationTopCategoryRequestSetCategoryList) *SetConversationTopCategoryRequest {
	s.SetCategoryList = v
	return s
}

func (s *SetConversationTopCategoryRequest) SetSign(v int32) *SetConversationTopCategoryRequest {
	s.Sign = &v
	return s
}

type SetConversationTopCategoryRequestSetCategoryList struct {
	// This parameter is required.
	//
	// example:
	//
	// 0
	CategoryId *int64 `json:"categoryId,omitempty" xml:"categoryId,omitempty"`
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
}

func (s SetConversationTopCategoryRequestSetCategoryList) String() string {
	return tea.Prettify(s)
}

func (s SetConversationTopCategoryRequestSetCategoryList) GoString() string {
	return s.String()
}

func (s *SetConversationTopCategoryRequestSetCategoryList) SetCategoryId(v int64) *SetConversationTopCategoryRequestSetCategoryList {
	s.CategoryId = &v
	return s
}

func (s *SetConversationTopCategoryRequestSetCategoryList) SetOrder(v int32) *SetConversationTopCategoryRequestSetCategoryList {
	s.Order = &v
	return s
}

type SetConversationTopCategoryResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SetConversationTopCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetConversationTopCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *SetConversationTopCategoryResponseBody) SetSuccess(v bool) *SetConversationTopCategoryResponseBody {
	s.Success = &v
	return s
}

type SetConversationTopCategoryResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetConversationTopCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetConversationTopCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s SetConversationTopCategoryResponse) GoString() string {
	return s.String()
}

func (s *SetConversationTopCategoryResponse) SetHeaders(v map[string]*string) *SetConversationTopCategoryResponse {
	s.Headers = v
	return s
}

func (s *SetConversationTopCategoryResponse) SetStatusCode(v int32) *SetConversationTopCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *SetConversationTopCategoryResponse) SetBody(v *SetConversationTopCategoryResponseBody) *SetConversationTopCategoryResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// 1234
	DeptId   *string   `json:"deptId,omitempty" xml:"deptId,omitempty"`
	LabelIds []*string `json:"labelIds,omitempty" xml:"labelIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1234
	PartnerNum *string `json:"partnerNum,omitempty" xml:"partnerNum,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
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

type SetOrgTopConversationCategoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SetOrgTopConversationCategoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetOrgTopConversationCategoryHeaders) GoString() string {
	return s.String()
}

func (s *SetOrgTopConversationCategoryHeaders) SetCommonHeaders(v map[string]*string) *SetOrgTopConversationCategoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetOrgTopConversationCategoryHeaders) SetXAcsDingtalkAccessToken(v string) *SetOrgTopConversationCategoryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SetOrgTopConversationCategoryRequest struct {
	Body []*SetOrgTopConversationCategoryRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s SetOrgTopConversationCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s SetOrgTopConversationCategoryRequest) GoString() string {
	return s.String()
}

func (s *SetOrgTopConversationCategoryRequest) SetBody(v []*SetOrgTopConversationCategoryRequestBody) *SetOrgTopConversationCategoryRequest {
	s.Body = v
	return s
}

type SetOrgTopConversationCategoryRequestBody struct {
	// example:
	//
	// 123111
	CategoryId *int64 `json:"categoryId,omitempty" xml:"categoryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 重要保障
	CategoryName *string `json:"categoryName,omitempty" xml:"categoryName,omitempty"`
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
}

func (s SetOrgTopConversationCategoryRequestBody) String() string {
	return tea.Prettify(s)
}

func (s SetOrgTopConversationCategoryRequestBody) GoString() string {
	return s.String()
}

func (s *SetOrgTopConversationCategoryRequestBody) SetCategoryId(v int64) *SetOrgTopConversationCategoryRequestBody {
	s.CategoryId = &v
	return s
}

func (s *SetOrgTopConversationCategoryRequestBody) SetCategoryName(v string) *SetOrgTopConversationCategoryRequestBody {
	s.CategoryName = &v
	return s
}

func (s *SetOrgTopConversationCategoryRequestBody) SetOrder(v int32) *SetOrgTopConversationCategoryRequestBody {
	s.Order = &v
	return s
}

type SetOrgTopConversationCategoryResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SetOrgTopConversationCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetOrgTopConversationCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *SetOrgTopConversationCategoryResponseBody) SetResult(v string) *SetOrgTopConversationCategoryResponseBody {
	s.Result = &v
	return s
}

func (s *SetOrgTopConversationCategoryResponseBody) SetSuccess(v bool) *SetOrgTopConversationCategoryResponseBody {
	s.Success = &v
	return s
}

type SetOrgTopConversationCategoryResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetOrgTopConversationCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetOrgTopConversationCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s SetOrgTopConversationCategoryResponse) GoString() string {
	return s.String()
}

func (s *SetOrgTopConversationCategoryResponse) SetHeaders(v map[string]*string) *SetOrgTopConversationCategoryResponse {
	s.Headers = v
	return s
}

func (s *SetOrgTopConversationCategoryResponse) SetStatusCode(v int32) *SetOrgTopConversationCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *SetOrgTopConversationCategoryResponse) SetBody(v *SetOrgTopConversationCategoryResponseBody) *SetOrgTopConversationCategoryResponse {
	s.Body = v
	return s
}

type SpecialRuleBatchReceiverHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SpecialRuleBatchReceiverHeaders) String() string {
	return tea.Prettify(s)
}

func (s SpecialRuleBatchReceiverHeaders) GoString() string {
	return s.String()
}

func (s *SpecialRuleBatchReceiverHeaders) SetCommonHeaders(v map[string]*string) *SpecialRuleBatchReceiverHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SpecialRuleBatchReceiverHeaders) SetXAcsDingtalkAccessToken(v string) *SpecialRuleBatchReceiverHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SpecialRuleBatchReceiverRequest struct {
	BatchNo         *string                                `json:"batchNo,omitempty" xml:"batchNo,omitempty"`
	CardOptions     *string                                `json:"cardOptions,omitempty" xml:"cardOptions,omitempty"`
	Data            []*SpecialRuleBatchReceiverRequestData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	RuleCode        *string                                `json:"ruleCode,omitempty" xml:"ruleCode,omitempty"`
	SecretKey       *string                                `json:"secretKey,omitempty" xml:"secretKey,omitempty"`
	SpecialStrategy *bool                                  `json:"specialStrategy,omitempty" xml:"specialStrategy,omitempty"`
	TaskBatchNo     *string                                `json:"taskBatchNo,omitempty" xml:"taskBatchNo,omitempty"`
}

func (s SpecialRuleBatchReceiverRequest) String() string {
	return tea.Prettify(s)
}

func (s SpecialRuleBatchReceiverRequest) GoString() string {
	return s.String()
}

func (s *SpecialRuleBatchReceiverRequest) SetBatchNo(v string) *SpecialRuleBatchReceiverRequest {
	s.BatchNo = &v
	return s
}

func (s *SpecialRuleBatchReceiverRequest) SetCardOptions(v string) *SpecialRuleBatchReceiverRequest {
	s.CardOptions = &v
	return s
}

func (s *SpecialRuleBatchReceiverRequest) SetData(v []*SpecialRuleBatchReceiverRequestData) *SpecialRuleBatchReceiverRequest {
	s.Data = v
	return s
}

func (s *SpecialRuleBatchReceiverRequest) SetRuleCode(v string) *SpecialRuleBatchReceiverRequest {
	s.RuleCode = &v
	return s
}

func (s *SpecialRuleBatchReceiverRequest) SetSecretKey(v string) *SpecialRuleBatchReceiverRequest {
	s.SecretKey = &v
	return s
}

func (s *SpecialRuleBatchReceiverRequest) SetSpecialStrategy(v bool) *SpecialRuleBatchReceiverRequest {
	s.SpecialStrategy = &v
	return s
}

func (s *SpecialRuleBatchReceiverRequest) SetTaskBatchNo(v string) *SpecialRuleBatchReceiverRequest {
	s.TaskBatchNo = &v
	return s
}

type SpecialRuleBatchReceiverRequestData struct {
	AtAccount       *string                                   `json:"atAccount,omitempty" xml:"atAccount,omitempty"`
	Attrs           *SpecialRuleBatchReceiverRequestDataAttrs `json:"attrs,omitempty" xml:"attrs,omitempty" type:"Struct"`
	CallbackUrl     *string                                   `json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`
	CardCallbackUrl *string                                   `json:"cardCallbackUrl,omitempty" xml:"cardCallbackUrl,omitempty"`
	Content         map[string]map[string]interface{}         `json:"content,omitempty" xml:"content,omitempty"`
	IsAtAll         *bool                                     `json:"isAtAll,omitempty" xml:"isAtAll,omitempty"`
	PrivateContent  map[string]map[string]interface{}         `json:"privateContent,omitempty" xml:"privateContent,omitempty"`
	ReceiverAccount *string                                   `json:"receiverAccount,omitempty" xml:"receiverAccount,omitempty"`
	ReceiverType    *int32                                    `json:"receiverType,omitempty" xml:"receiverType,omitempty"`
	SerialNumber    *string                                   `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
}

func (s SpecialRuleBatchReceiverRequestData) String() string {
	return tea.Prettify(s)
}

func (s SpecialRuleBatchReceiverRequestData) GoString() string {
	return s.String()
}

func (s *SpecialRuleBatchReceiverRequestData) SetAtAccount(v string) *SpecialRuleBatchReceiverRequestData {
	s.AtAccount = &v
	return s
}

func (s *SpecialRuleBatchReceiverRequestData) SetAttrs(v *SpecialRuleBatchReceiverRequestDataAttrs) *SpecialRuleBatchReceiverRequestData {
	s.Attrs = v
	return s
}

func (s *SpecialRuleBatchReceiverRequestData) SetCallbackUrl(v string) *SpecialRuleBatchReceiverRequestData {
	s.CallbackUrl = &v
	return s
}

func (s *SpecialRuleBatchReceiverRequestData) SetCardCallbackUrl(v string) *SpecialRuleBatchReceiverRequestData {
	s.CardCallbackUrl = &v
	return s
}

func (s *SpecialRuleBatchReceiverRequestData) SetContent(v map[string]map[string]interface{}) *SpecialRuleBatchReceiverRequestData {
	s.Content = v
	return s
}

func (s *SpecialRuleBatchReceiverRequestData) SetIsAtAll(v bool) *SpecialRuleBatchReceiverRequestData {
	s.IsAtAll = &v
	return s
}

func (s *SpecialRuleBatchReceiverRequestData) SetPrivateContent(v map[string]map[string]interface{}) *SpecialRuleBatchReceiverRequestData {
	s.PrivateContent = v
	return s
}

func (s *SpecialRuleBatchReceiverRequestData) SetReceiverAccount(v string) *SpecialRuleBatchReceiverRequestData {
	s.ReceiverAccount = &v
	return s
}

func (s *SpecialRuleBatchReceiverRequestData) SetReceiverType(v int32) *SpecialRuleBatchReceiverRequestData {
	s.ReceiverType = &v
	return s
}

func (s *SpecialRuleBatchReceiverRequestData) SetSerialNumber(v string) *SpecialRuleBatchReceiverRequestData {
	s.SerialNumber = &v
	return s
}

type SpecialRuleBatchReceiverRequestDataAttrs struct {
	ListUnitId []*int64 `json:"listUnitId,omitempty" xml:"listUnitId,omitempty" type:"Repeated"`
}

func (s SpecialRuleBatchReceiverRequestDataAttrs) String() string {
	return tea.Prettify(s)
}

func (s SpecialRuleBatchReceiverRequestDataAttrs) GoString() string {
	return s.String()
}

func (s *SpecialRuleBatchReceiverRequestDataAttrs) SetListUnitId(v []*int64) *SpecialRuleBatchReceiverRequestDataAttrs {
	s.ListUnitId = v
	return s
}

type SpecialRuleBatchReceiverResponseBody struct {
	Code  *int32                                        `json:"code,omitempty" xml:"code,omitempty"`
	Data  []*SpecialRuleBatchReceiverResponseBodyData   `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Msg   *string                                       `json:"msg,omitempty" xml:"msg,omitempty"`
	MsgId *string                                       `json:"msgId,omitempty" xml:"msgId,omitempty"`
	Rows  [][]*SpecialRuleBatchReceiverResponseBodyRows `json:"rows,omitempty" xml:"rows,omitempty" type:"Repeated"`
}

func (s SpecialRuleBatchReceiverResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SpecialRuleBatchReceiverResponseBody) GoString() string {
	return s.String()
}

func (s *SpecialRuleBatchReceiverResponseBody) SetCode(v int32) *SpecialRuleBatchReceiverResponseBody {
	s.Code = &v
	return s
}

func (s *SpecialRuleBatchReceiverResponseBody) SetData(v []*SpecialRuleBatchReceiverResponseBodyData) *SpecialRuleBatchReceiverResponseBody {
	s.Data = v
	return s
}

func (s *SpecialRuleBatchReceiverResponseBody) SetMsg(v string) *SpecialRuleBatchReceiverResponseBody {
	s.Msg = &v
	return s
}

func (s *SpecialRuleBatchReceiverResponseBody) SetMsgId(v string) *SpecialRuleBatchReceiverResponseBody {
	s.MsgId = &v
	return s
}

func (s *SpecialRuleBatchReceiverResponseBody) SetRows(v [][]*SpecialRuleBatchReceiverResponseBodyRows) *SpecialRuleBatchReceiverResponseBody {
	s.Rows = v
	return s
}

type SpecialRuleBatchReceiverResponseBodyData struct {
	MsgId        *string `json:"msgId,omitempty" xml:"msgId,omitempty"`
	SerialNumber *string `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
}

func (s SpecialRuleBatchReceiverResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SpecialRuleBatchReceiverResponseBodyData) GoString() string {
	return s.String()
}

func (s *SpecialRuleBatchReceiverResponseBodyData) SetMsgId(v string) *SpecialRuleBatchReceiverResponseBodyData {
	s.MsgId = &v
	return s
}

func (s *SpecialRuleBatchReceiverResponseBodyData) SetSerialNumber(v string) *SpecialRuleBatchReceiverResponseBodyData {
	s.SerialNumber = &v
	return s
}

type SpecialRuleBatchReceiverResponseBodyRows struct {
	SerialNumber *string `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
	MsgId        *string `json:"msgId,omitempty" xml:"msgId,omitempty"`
}

func (s SpecialRuleBatchReceiverResponseBodyRows) String() string {
	return tea.Prettify(s)
}

func (s SpecialRuleBatchReceiverResponseBodyRows) GoString() string {
	return s.String()
}

func (s *SpecialRuleBatchReceiverResponseBodyRows) SetSerialNumber(v string) *SpecialRuleBatchReceiverResponseBodyRows {
	s.SerialNumber = &v
	return s
}

func (s *SpecialRuleBatchReceiverResponseBodyRows) SetMsgId(v string) *SpecialRuleBatchReceiverResponseBodyRows {
	s.MsgId = &v
	return s
}

type SpecialRuleBatchReceiverResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SpecialRuleBatchReceiverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SpecialRuleBatchReceiverResponse) String() string {
	return tea.Prettify(s)
}

func (s SpecialRuleBatchReceiverResponse) GoString() string {
	return s.String()
}

func (s *SpecialRuleBatchReceiverResponse) SetHeaders(v map[string]*string) *SpecialRuleBatchReceiverResponse {
	s.Headers = v
	return s
}

func (s *SpecialRuleBatchReceiverResponse) SetStatusCode(v int32) *SpecialRuleBatchReceiverResponse {
	s.StatusCode = &v
	return s
}

func (s *SpecialRuleBatchReceiverResponse) SetBody(v *SpecialRuleBatchReceiverResponseBody) *SpecialRuleBatchReceiverResponse {
	s.Body = v
	return s
}

type TaskInfoAddDelTaskPersonHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TaskInfoAddDelTaskPersonHeaders) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoAddDelTaskPersonHeaders) GoString() string {
	return s.String()
}

func (s *TaskInfoAddDelTaskPersonHeaders) SetCommonHeaders(v map[string]*string) *TaskInfoAddDelTaskPersonHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TaskInfoAddDelTaskPersonHeaders) SetXAcsDingtalkAccessToken(v string) *TaskInfoAddDelTaskPersonHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TaskInfoAddDelTaskPersonRequest struct {
	// This parameter is required.
	OperateType *int32 `json:"operateType,omitempty" xml:"operateType,omitempty"`
	// This parameter is required.
	OperatorAccount *string `json:"operatorAccount,omitempty" xml:"operatorAccount,omitempty"`
	// This parameter is required.
	OutTaskId *string `json:"outTaskId,omitempty" xml:"outTaskId,omitempty"`
	// This parameter is required.
	ProjId *string `json:"projId,omitempty" xml:"projId,omitempty"`
	// This parameter is required.
	SecretKey *string `json:"secretKey,omitempty" xml:"secretKey,omitempty"`
	// This parameter is required.
	TaskExecutePersonDTOS []*TaskInfoAddDelTaskPersonRequestTaskExecutePersonDTOS `json:"taskExecutePersonDTOS,omitempty" xml:"taskExecutePersonDTOS,omitempty" type:"Repeated"`
}

func (s TaskInfoAddDelTaskPersonRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoAddDelTaskPersonRequest) GoString() string {
	return s.String()
}

func (s *TaskInfoAddDelTaskPersonRequest) SetOperateType(v int32) *TaskInfoAddDelTaskPersonRequest {
	s.OperateType = &v
	return s
}

func (s *TaskInfoAddDelTaskPersonRequest) SetOperatorAccount(v string) *TaskInfoAddDelTaskPersonRequest {
	s.OperatorAccount = &v
	return s
}

func (s *TaskInfoAddDelTaskPersonRequest) SetOutTaskId(v string) *TaskInfoAddDelTaskPersonRequest {
	s.OutTaskId = &v
	return s
}

func (s *TaskInfoAddDelTaskPersonRequest) SetProjId(v string) *TaskInfoAddDelTaskPersonRequest {
	s.ProjId = &v
	return s
}

func (s *TaskInfoAddDelTaskPersonRequest) SetSecretKey(v string) *TaskInfoAddDelTaskPersonRequest {
	s.SecretKey = &v
	return s
}

func (s *TaskInfoAddDelTaskPersonRequest) SetTaskExecutePersonDTOS(v []*TaskInfoAddDelTaskPersonRequestTaskExecutePersonDTOS) *TaskInfoAddDelTaskPersonRequest {
	s.TaskExecutePersonDTOS = v
	return s
}

type TaskInfoAddDelTaskPersonRequestTaskExecutePersonDTOS struct {
	EmployeeCode *string `json:"employeeCode,omitempty" xml:"employeeCode,omitempty"`
	PersonType   *int32  `json:"personType,omitempty" xml:"personType,omitempty"`
}

func (s TaskInfoAddDelTaskPersonRequestTaskExecutePersonDTOS) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoAddDelTaskPersonRequestTaskExecutePersonDTOS) GoString() string {
	return s.String()
}

func (s *TaskInfoAddDelTaskPersonRequestTaskExecutePersonDTOS) SetEmployeeCode(v string) *TaskInfoAddDelTaskPersonRequestTaskExecutePersonDTOS {
	s.EmployeeCode = &v
	return s
}

func (s *TaskInfoAddDelTaskPersonRequestTaskExecutePersonDTOS) SetPersonType(v int32) *TaskInfoAddDelTaskPersonRequestTaskExecutePersonDTOS {
	s.PersonType = &v
	return s
}

type TaskInfoAddDelTaskPersonResponseBody struct {
	Code    *int32      `json:"code,omitempty" xml:"code,omitempty"`
	Data    interface{} `json:"data,omitempty" xml:"data,omitempty"`
	Message *string     `json:"message,omitempty" xml:"message,omitempty"`
}

func (s TaskInfoAddDelTaskPersonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoAddDelTaskPersonResponseBody) GoString() string {
	return s.String()
}

func (s *TaskInfoAddDelTaskPersonResponseBody) SetCode(v int32) *TaskInfoAddDelTaskPersonResponseBody {
	s.Code = &v
	return s
}

func (s *TaskInfoAddDelTaskPersonResponseBody) SetData(v interface{}) *TaskInfoAddDelTaskPersonResponseBody {
	s.Data = v
	return s
}

func (s *TaskInfoAddDelTaskPersonResponseBody) SetMessage(v string) *TaskInfoAddDelTaskPersonResponseBody {
	s.Message = &v
	return s
}

type TaskInfoAddDelTaskPersonResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TaskInfoAddDelTaskPersonResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TaskInfoAddDelTaskPersonResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoAddDelTaskPersonResponse) GoString() string {
	return s.String()
}

func (s *TaskInfoAddDelTaskPersonResponse) SetHeaders(v map[string]*string) *TaskInfoAddDelTaskPersonResponse {
	s.Headers = v
	return s
}

func (s *TaskInfoAddDelTaskPersonResponse) SetStatusCode(v int32) *TaskInfoAddDelTaskPersonResponse {
	s.StatusCode = &v
	return s
}

func (s *TaskInfoAddDelTaskPersonResponse) SetBody(v *TaskInfoAddDelTaskPersonResponseBody) *TaskInfoAddDelTaskPersonResponse {
	s.Body = v
	return s
}

type TaskInfoCancelOrDelTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TaskInfoCancelOrDelTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoCancelOrDelTaskHeaders) GoString() string {
	return s.String()
}

func (s *TaskInfoCancelOrDelTaskHeaders) SetCommonHeaders(v map[string]*string) *TaskInfoCancelOrDelTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TaskInfoCancelOrDelTaskHeaders) SetXAcsDingtalkAccessToken(v string) *TaskInfoCancelOrDelTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TaskInfoCancelOrDelTaskRequest struct {
	CardDTO *TaskInfoCancelOrDelTaskRequestCardDTO `json:"cardDTO,omitempty" xml:"cardDTO,omitempty" type:"Struct"`
	// This parameter is required.
	OperatorAccount *string `json:"operatorAccount,omitempty" xml:"operatorAccount,omitempty"`
	OutTaskId       *string `json:"outTaskId,omitempty" xml:"outTaskId,omitempty"`
	// This parameter is required.
	ProjId *string `json:"projId,omitempty" xml:"projId,omitempty"`
	// This parameter is required.
	SecretKey             *string                                                `json:"secretKey,omitempty" xml:"secretKey,omitempty"`
	SendMsgFlag           *int32                                                 `json:"sendMsgFlag,omitempty" xml:"sendMsgFlag,omitempty"`
	TaskExecutePersonDTOS []*TaskInfoCancelOrDelTaskRequestTaskExecutePersonDTOS `json:"taskExecutePersonDTOS,omitempty" xml:"taskExecutePersonDTOS,omitempty" type:"Repeated"`
}

func (s TaskInfoCancelOrDelTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoCancelOrDelTaskRequest) GoString() string {
	return s.String()
}

func (s *TaskInfoCancelOrDelTaskRequest) SetCardDTO(v *TaskInfoCancelOrDelTaskRequestCardDTO) *TaskInfoCancelOrDelTaskRequest {
	s.CardDTO = v
	return s
}

func (s *TaskInfoCancelOrDelTaskRequest) SetOperatorAccount(v string) *TaskInfoCancelOrDelTaskRequest {
	s.OperatorAccount = &v
	return s
}

func (s *TaskInfoCancelOrDelTaskRequest) SetOutTaskId(v string) *TaskInfoCancelOrDelTaskRequest {
	s.OutTaskId = &v
	return s
}

func (s *TaskInfoCancelOrDelTaskRequest) SetProjId(v string) *TaskInfoCancelOrDelTaskRequest {
	s.ProjId = &v
	return s
}

func (s *TaskInfoCancelOrDelTaskRequest) SetSecretKey(v string) *TaskInfoCancelOrDelTaskRequest {
	s.SecretKey = &v
	return s
}

func (s *TaskInfoCancelOrDelTaskRequest) SetSendMsgFlag(v int32) *TaskInfoCancelOrDelTaskRequest {
	s.SendMsgFlag = &v
	return s
}

func (s *TaskInfoCancelOrDelTaskRequest) SetTaskExecutePersonDTOS(v []*TaskInfoCancelOrDelTaskRequestTaskExecutePersonDTOS) *TaskInfoCancelOrDelTaskRequest {
	s.TaskExecutePersonDTOS = v
	return s
}

type TaskInfoCancelOrDelTaskRequestCardDTO struct {
	AtAccount       *string     `json:"atAccount,omitempty" xml:"atAccount,omitempty"`
	CardCallbackUrl *string     `json:"cardCallbackUrl,omitempty" xml:"cardCallbackUrl,omitempty"`
	Content         interface{} `json:"content,omitempty" xml:"content,omitempty"`
	IsAtAll         *bool       `json:"isAtAll,omitempty" xml:"isAtAll,omitempty"`
	ReceiverAccount *string     `json:"receiverAccount,omitempty" xml:"receiverAccount,omitempty"`
}

func (s TaskInfoCancelOrDelTaskRequestCardDTO) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoCancelOrDelTaskRequestCardDTO) GoString() string {
	return s.String()
}

func (s *TaskInfoCancelOrDelTaskRequestCardDTO) SetAtAccount(v string) *TaskInfoCancelOrDelTaskRequestCardDTO {
	s.AtAccount = &v
	return s
}

func (s *TaskInfoCancelOrDelTaskRequestCardDTO) SetCardCallbackUrl(v string) *TaskInfoCancelOrDelTaskRequestCardDTO {
	s.CardCallbackUrl = &v
	return s
}

func (s *TaskInfoCancelOrDelTaskRequestCardDTO) SetContent(v interface{}) *TaskInfoCancelOrDelTaskRequestCardDTO {
	s.Content = v
	return s
}

func (s *TaskInfoCancelOrDelTaskRequestCardDTO) SetIsAtAll(v bool) *TaskInfoCancelOrDelTaskRequestCardDTO {
	s.IsAtAll = &v
	return s
}

func (s *TaskInfoCancelOrDelTaskRequestCardDTO) SetReceiverAccount(v string) *TaskInfoCancelOrDelTaskRequestCardDTO {
	s.ReceiverAccount = &v
	return s
}

type TaskInfoCancelOrDelTaskRequestTaskExecutePersonDTOS struct {
	EmployeeCode *string `json:"employeeCode,omitempty" xml:"employeeCode,omitempty"`
	PersonType   *int32  `json:"personType,omitempty" xml:"personType,omitempty"`
}

func (s TaskInfoCancelOrDelTaskRequestTaskExecutePersonDTOS) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoCancelOrDelTaskRequestTaskExecutePersonDTOS) GoString() string {
	return s.String()
}

func (s *TaskInfoCancelOrDelTaskRequestTaskExecutePersonDTOS) SetEmployeeCode(v string) *TaskInfoCancelOrDelTaskRequestTaskExecutePersonDTOS {
	s.EmployeeCode = &v
	return s
}

func (s *TaskInfoCancelOrDelTaskRequestTaskExecutePersonDTOS) SetPersonType(v int32) *TaskInfoCancelOrDelTaskRequestTaskExecutePersonDTOS {
	s.PersonType = &v
	return s
}

type TaskInfoCancelOrDelTaskResponseBody struct {
	Code    *int32      `json:"code,omitempty" xml:"code,omitempty"`
	Data    interface{} `json:"data,omitempty" xml:"data,omitempty"`
	Message *string     `json:"message,omitempty" xml:"message,omitempty"`
}

func (s TaskInfoCancelOrDelTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoCancelOrDelTaskResponseBody) GoString() string {
	return s.String()
}

func (s *TaskInfoCancelOrDelTaskResponseBody) SetCode(v int32) *TaskInfoCancelOrDelTaskResponseBody {
	s.Code = &v
	return s
}

func (s *TaskInfoCancelOrDelTaskResponseBody) SetData(v interface{}) *TaskInfoCancelOrDelTaskResponseBody {
	s.Data = v
	return s
}

func (s *TaskInfoCancelOrDelTaskResponseBody) SetMessage(v string) *TaskInfoCancelOrDelTaskResponseBody {
	s.Message = &v
	return s
}

type TaskInfoCancelOrDelTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TaskInfoCancelOrDelTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TaskInfoCancelOrDelTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoCancelOrDelTaskResponse) GoString() string {
	return s.String()
}

func (s *TaskInfoCancelOrDelTaskResponse) SetHeaders(v map[string]*string) *TaskInfoCancelOrDelTaskResponse {
	s.Headers = v
	return s
}

func (s *TaskInfoCancelOrDelTaskResponse) SetStatusCode(v int32) *TaskInfoCancelOrDelTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *TaskInfoCancelOrDelTaskResponse) SetBody(v *TaskInfoCancelOrDelTaskResponseBody) *TaskInfoCancelOrDelTaskResponse {
	s.Body = v
	return s
}

type TaskInfoCreateAndStartTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TaskInfoCreateAndStartTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoCreateAndStartTaskHeaders) GoString() string {
	return s.String()
}

func (s *TaskInfoCreateAndStartTaskHeaders) SetCommonHeaders(v map[string]*string) *TaskInfoCreateAndStartTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TaskInfoCreateAndStartTaskHeaders) SetXAcsDingtalkAccessToken(v string) *TaskInfoCreateAndStartTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TaskInfoCreateAndStartTaskRequest struct {
	Attr                *TaskInfoCreateAndStartTaskRequestAttr       `json:"attr,omitempty" xml:"attr,omitempty" type:"Struct"`
	BacklogDTO          *TaskInfoCreateAndStartTaskRequestBacklogDTO `json:"backlogDTO,omitempty" xml:"backlogDTO,omitempty" type:"Struct"`
	BacklogGenerateFlag *int32                                       `json:"backlogGenerateFlag,omitempty" xml:"backlogGenerateFlag,omitempty"`
	BusinessCode        *string                                      `json:"businessCode,omitempty" xml:"businessCode,omitempty"`
	CanceldelTaskCardId *string                                      `json:"canceldelTaskCardId,omitempty" xml:"canceldelTaskCardId,omitempty"`
	CardDTO             *TaskInfoCreateAndStartTaskRequestCardDTO    `json:"cardDTO,omitempty" xml:"cardDTO,omitempty" type:"Struct"`
	// This parameter is required.
	CustomFlag       *int32                                      `json:"customFlag,omitempty" xml:"customFlag,omitempty"`
	DetailUrl        *TaskInfoCreateAndStartTaskRequestDetailUrl `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	FinishTaskCardId *string                                     `json:"finishTaskCardId,omitempty" xml:"finishTaskCardId,omitempty"`
	// This parameter is required.
	OperatorAccount *string `json:"operatorAccount,omitempty" xml:"operatorAccount,omitempty"`
	// This parameter is required.
	OutTaskId *string `json:"outTaskId,omitempty" xml:"outTaskId,omitempty"`
	// This parameter is required.
	ProjId    *string `json:"projId,omitempty" xml:"projId,omitempty"`
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	// This parameter is required.
	SecretKey             *string                                                   `json:"secretKey,omitempty" xml:"secretKey,omitempty"`
	SendMsgFlag           *int32                                                    `json:"sendMsgFlag,omitempty" xml:"sendMsgFlag,omitempty"`
	Sort                  *int32                                                    `json:"sort,omitempty" xml:"sort,omitempty"`
	StartTaskCardId       *string                                                   `json:"startTaskCardId,omitempty" xml:"startTaskCardId,omitempty"`
	State                 *int32                                                    `json:"state,omitempty" xml:"state,omitempty"`
	TaskContent           *string                                                   `json:"taskContent,omitempty" xml:"taskContent,omitempty"`
	TaskEndTime           *int64                                                    `json:"taskEndTime,omitempty" xml:"taskEndTime,omitempty"`
	TaskExecutePersonDTOS []*TaskInfoCreateAndStartTaskRequestTaskExecutePersonDTOS `json:"taskExecutePersonDTOS,omitempty" xml:"taskExecutePersonDTOS,omitempty" type:"Repeated"`
	TaskGroupDTOList      []*TaskInfoCreateAndStartTaskRequestTaskGroupDTOList      `json:"taskGroupDTOList,omitempty" xml:"taskGroupDTOList,omitempty" type:"Repeated"`
	// This parameter is required.
	TaskSystem    *string `json:"taskSystem,omitempty" xml:"taskSystem,omitempty"`
	TaskTemplCode *string `json:"taskTemplCode,omitempty" xml:"taskTemplCode,omitempty"`
	// This parameter is required.
	TaskTitle *string `json:"taskTitle,omitempty" xml:"taskTitle,omitempty"`
	// This parameter is required.
	TaskType         *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	TaskUrlMobile    *string `json:"taskUrlMobile,omitempty" xml:"taskUrlMobile,omitempty"`
	TaskUrlPc        *string `json:"taskUrlPc,omitempty" xml:"taskUrlPc,omitempty"`
	UpdateTaskCardId *string `json:"updateTaskCardId,omitempty" xml:"updateTaskCardId,omitempty"`
}

func (s TaskInfoCreateAndStartTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoCreateAndStartTaskRequest) GoString() string {
	return s.String()
}

func (s *TaskInfoCreateAndStartTaskRequest) SetAttr(v *TaskInfoCreateAndStartTaskRequestAttr) *TaskInfoCreateAndStartTaskRequest {
	s.Attr = v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetBacklogDTO(v *TaskInfoCreateAndStartTaskRequestBacklogDTO) *TaskInfoCreateAndStartTaskRequest {
	s.BacklogDTO = v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetBacklogGenerateFlag(v int32) *TaskInfoCreateAndStartTaskRequest {
	s.BacklogGenerateFlag = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetBusinessCode(v string) *TaskInfoCreateAndStartTaskRequest {
	s.BusinessCode = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetCanceldelTaskCardId(v string) *TaskInfoCreateAndStartTaskRequest {
	s.CanceldelTaskCardId = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetCardDTO(v *TaskInfoCreateAndStartTaskRequestCardDTO) *TaskInfoCreateAndStartTaskRequest {
	s.CardDTO = v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetCustomFlag(v int32) *TaskInfoCreateAndStartTaskRequest {
	s.CustomFlag = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetDetailUrl(v *TaskInfoCreateAndStartTaskRequestDetailUrl) *TaskInfoCreateAndStartTaskRequest {
	s.DetailUrl = v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetFinishTaskCardId(v string) *TaskInfoCreateAndStartTaskRequest {
	s.FinishTaskCardId = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetOperatorAccount(v string) *TaskInfoCreateAndStartTaskRequest {
	s.OperatorAccount = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetOutTaskId(v string) *TaskInfoCreateAndStartTaskRequest {
	s.OutTaskId = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetProjId(v string) *TaskInfoCreateAndStartTaskRequest {
	s.ProjId = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetRobotCode(v string) *TaskInfoCreateAndStartTaskRequest {
	s.RobotCode = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetSecretKey(v string) *TaskInfoCreateAndStartTaskRequest {
	s.SecretKey = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetSendMsgFlag(v int32) *TaskInfoCreateAndStartTaskRequest {
	s.SendMsgFlag = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetSort(v int32) *TaskInfoCreateAndStartTaskRequest {
	s.Sort = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetStartTaskCardId(v string) *TaskInfoCreateAndStartTaskRequest {
	s.StartTaskCardId = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetState(v int32) *TaskInfoCreateAndStartTaskRequest {
	s.State = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetTaskContent(v string) *TaskInfoCreateAndStartTaskRequest {
	s.TaskContent = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetTaskEndTime(v int64) *TaskInfoCreateAndStartTaskRequest {
	s.TaskEndTime = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetTaskExecutePersonDTOS(v []*TaskInfoCreateAndStartTaskRequestTaskExecutePersonDTOS) *TaskInfoCreateAndStartTaskRequest {
	s.TaskExecutePersonDTOS = v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetTaskGroupDTOList(v []*TaskInfoCreateAndStartTaskRequestTaskGroupDTOList) *TaskInfoCreateAndStartTaskRequest {
	s.TaskGroupDTOList = v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetTaskSystem(v string) *TaskInfoCreateAndStartTaskRequest {
	s.TaskSystem = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetTaskTemplCode(v string) *TaskInfoCreateAndStartTaskRequest {
	s.TaskTemplCode = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetTaskTitle(v string) *TaskInfoCreateAndStartTaskRequest {
	s.TaskTitle = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetTaskType(v string) *TaskInfoCreateAndStartTaskRequest {
	s.TaskType = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetTaskUrlMobile(v string) *TaskInfoCreateAndStartTaskRequest {
	s.TaskUrlMobile = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetTaskUrlPc(v string) *TaskInfoCreateAndStartTaskRequest {
	s.TaskUrlPc = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequest) SetUpdateTaskCardId(v string) *TaskInfoCreateAndStartTaskRequest {
	s.UpdateTaskCardId = &v
	return s
}

type TaskInfoCreateAndStartTaskRequestAttr struct {
	ListTaskDynamicAttr []*TaskInfoCreateAndStartTaskRequestAttrListTaskDynamicAttr `json:"listTaskDynamicAttr,omitempty" xml:"listTaskDynamicAttr,omitempty" type:"Repeated"`
}

func (s TaskInfoCreateAndStartTaskRequestAttr) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoCreateAndStartTaskRequestAttr) GoString() string {
	return s.String()
}

func (s *TaskInfoCreateAndStartTaskRequestAttr) SetListTaskDynamicAttr(v []*TaskInfoCreateAndStartTaskRequestAttrListTaskDynamicAttr) *TaskInfoCreateAndStartTaskRequestAttr {
	s.ListTaskDynamicAttr = v
	return s
}

type TaskInfoCreateAndStartTaskRequestAttrListTaskDynamicAttr struct {
	AttrCode            *string   `json:"attrCode,omitempty" xml:"attrCode,omitempty"`
	ListAttrOptionsCode []*string `json:"listAttrOptionsCode,omitempty" xml:"listAttrOptionsCode,omitempty" type:"Repeated"`
}

func (s TaskInfoCreateAndStartTaskRequestAttrListTaskDynamicAttr) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoCreateAndStartTaskRequestAttrListTaskDynamicAttr) GoString() string {
	return s.String()
}

func (s *TaskInfoCreateAndStartTaskRequestAttrListTaskDynamicAttr) SetAttrCode(v string) *TaskInfoCreateAndStartTaskRequestAttrListTaskDynamicAttr {
	s.AttrCode = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequestAttrListTaskDynamicAttr) SetListAttrOptionsCode(v []*string) *TaskInfoCreateAndStartTaskRequestAttrListTaskDynamicAttr {
	s.ListAttrOptionsCode = v
	return s
}

type TaskInfoCreateAndStartTaskRequestBacklogDTO struct {
	Content *TaskInfoCreateAndStartTaskRequestBacklogDTOContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
}

func (s TaskInfoCreateAndStartTaskRequestBacklogDTO) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoCreateAndStartTaskRequestBacklogDTO) GoString() string {
	return s.String()
}

func (s *TaskInfoCreateAndStartTaskRequestBacklogDTO) SetContent(v *TaskInfoCreateAndStartTaskRequestBacklogDTOContent) *TaskInfoCreateAndStartTaskRequestBacklogDTO {
	s.Content = v
	return s
}

type TaskInfoCreateAndStartTaskRequestBacklogDTOContent struct {
	IsOnlyShowExecutor *bool  `json:"isOnlyShowExecutor,omitempty" xml:"isOnlyShowExecutor,omitempty"`
	Priority           *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
}

func (s TaskInfoCreateAndStartTaskRequestBacklogDTOContent) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoCreateAndStartTaskRequestBacklogDTOContent) GoString() string {
	return s.String()
}

func (s *TaskInfoCreateAndStartTaskRequestBacklogDTOContent) SetIsOnlyShowExecutor(v bool) *TaskInfoCreateAndStartTaskRequestBacklogDTOContent {
	s.IsOnlyShowExecutor = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequestBacklogDTOContent) SetPriority(v int32) *TaskInfoCreateAndStartTaskRequestBacklogDTOContent {
	s.Priority = &v
	return s
}

type TaskInfoCreateAndStartTaskRequestCardDTO struct {
	AtAccount       *string     `json:"atAccount,omitempty" xml:"atAccount,omitempty"`
	CardCallbackUrl *string     `json:"cardCallbackUrl,omitempty" xml:"cardCallbackUrl,omitempty"`
	Content         interface{} `json:"content,omitempty" xml:"content,omitempty"`
	IsAtAll         *bool       `json:"isAtAll,omitempty" xml:"isAtAll,omitempty"`
	ReceiverAccount *string     `json:"receiverAccount,omitempty" xml:"receiverAccount,omitempty"`
}

func (s TaskInfoCreateAndStartTaskRequestCardDTO) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoCreateAndStartTaskRequestCardDTO) GoString() string {
	return s.String()
}

func (s *TaskInfoCreateAndStartTaskRequestCardDTO) SetAtAccount(v string) *TaskInfoCreateAndStartTaskRequestCardDTO {
	s.AtAccount = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequestCardDTO) SetCardCallbackUrl(v string) *TaskInfoCreateAndStartTaskRequestCardDTO {
	s.CardCallbackUrl = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequestCardDTO) SetContent(v interface{}) *TaskInfoCreateAndStartTaskRequestCardDTO {
	s.Content = v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequestCardDTO) SetIsAtAll(v bool) *TaskInfoCreateAndStartTaskRequestCardDTO {
	s.IsAtAll = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequestCardDTO) SetReceiverAccount(v string) *TaskInfoCreateAndStartTaskRequestCardDTO {
	s.ReceiverAccount = &v
	return s
}

type TaskInfoCreateAndStartTaskRequestDetailUrl struct {
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
	PcUrl  *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
}

func (s TaskInfoCreateAndStartTaskRequestDetailUrl) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoCreateAndStartTaskRequestDetailUrl) GoString() string {
	return s.String()
}

func (s *TaskInfoCreateAndStartTaskRequestDetailUrl) SetAppUrl(v string) *TaskInfoCreateAndStartTaskRequestDetailUrl {
	s.AppUrl = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequestDetailUrl) SetPcUrl(v string) *TaskInfoCreateAndStartTaskRequestDetailUrl {
	s.PcUrl = &v
	return s
}

type TaskInfoCreateAndStartTaskRequestTaskExecutePersonDTOS struct {
	EmployeeCode *string `json:"employeeCode,omitempty" xml:"employeeCode,omitempty"`
	PersonType   *int32  `json:"personType,omitempty" xml:"personType,omitempty"`
}

func (s TaskInfoCreateAndStartTaskRequestTaskExecutePersonDTOS) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoCreateAndStartTaskRequestTaskExecutePersonDTOS) GoString() string {
	return s.String()
}

func (s *TaskInfoCreateAndStartTaskRequestTaskExecutePersonDTOS) SetEmployeeCode(v string) *TaskInfoCreateAndStartTaskRequestTaskExecutePersonDTOS {
	s.EmployeeCode = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskRequestTaskExecutePersonDTOS) SetPersonType(v int32) *TaskInfoCreateAndStartTaskRequestTaskExecutePersonDTOS {
	s.PersonType = &v
	return s
}

type TaskInfoCreateAndStartTaskRequestTaskGroupDTOList struct {
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s TaskInfoCreateAndStartTaskRequestTaskGroupDTOList) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoCreateAndStartTaskRequestTaskGroupDTOList) GoString() string {
	return s.String()
}

func (s *TaskInfoCreateAndStartTaskRequestTaskGroupDTOList) SetOpenConversationId(v string) *TaskInfoCreateAndStartTaskRequestTaskGroupDTOList {
	s.OpenConversationId = &v
	return s
}

type TaskInfoCreateAndStartTaskResponseBody struct {
	Code    *int32                                      `json:"code,omitempty" xml:"code,omitempty"`
	Data    *TaskInfoCreateAndStartTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                                     `json:"message,omitempty" xml:"message,omitempty"`
}

func (s TaskInfoCreateAndStartTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoCreateAndStartTaskResponseBody) GoString() string {
	return s.String()
}

func (s *TaskInfoCreateAndStartTaskResponseBody) SetCode(v int32) *TaskInfoCreateAndStartTaskResponseBody {
	s.Code = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskResponseBody) SetData(v *TaskInfoCreateAndStartTaskResponseBodyData) *TaskInfoCreateAndStartTaskResponseBody {
	s.Data = v
	return s
}

func (s *TaskInfoCreateAndStartTaskResponseBody) SetMessage(v string) *TaskInfoCreateAndStartTaskResponseBody {
	s.Message = &v
	return s
}

type TaskInfoCreateAndStartTaskResponseBodyData struct {
	GroupVoList []*TaskInfoCreateAndStartTaskResponseBodyDataGroupVoList `json:"groupVoList,omitempty" xml:"groupVoList,omitempty" type:"Repeated"`
	TaskId      *string                                                  `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s TaskInfoCreateAndStartTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoCreateAndStartTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *TaskInfoCreateAndStartTaskResponseBodyData) SetGroupVoList(v []*TaskInfoCreateAndStartTaskResponseBodyDataGroupVoList) *TaskInfoCreateAndStartTaskResponseBodyData {
	s.GroupVoList = v
	return s
}

func (s *TaskInfoCreateAndStartTaskResponseBodyData) SetTaskId(v string) *TaskInfoCreateAndStartTaskResponseBodyData {
	s.TaskId = &v
	return s
}

type TaskInfoCreateAndStartTaskResponseBodyDataGroupVoList struct {
	CorpId             *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s TaskInfoCreateAndStartTaskResponseBodyDataGroupVoList) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoCreateAndStartTaskResponseBodyDataGroupVoList) GoString() string {
	return s.String()
}

func (s *TaskInfoCreateAndStartTaskResponseBodyDataGroupVoList) SetCorpId(v string) *TaskInfoCreateAndStartTaskResponseBodyDataGroupVoList {
	s.CorpId = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskResponseBodyDataGroupVoList) SetOpenConversationId(v string) *TaskInfoCreateAndStartTaskResponseBodyDataGroupVoList {
	s.OpenConversationId = &v
	return s
}

type TaskInfoCreateAndStartTaskResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TaskInfoCreateAndStartTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TaskInfoCreateAndStartTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoCreateAndStartTaskResponse) GoString() string {
	return s.String()
}

func (s *TaskInfoCreateAndStartTaskResponse) SetHeaders(v map[string]*string) *TaskInfoCreateAndStartTaskResponse {
	s.Headers = v
	return s
}

func (s *TaskInfoCreateAndStartTaskResponse) SetStatusCode(v int32) *TaskInfoCreateAndStartTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *TaskInfoCreateAndStartTaskResponse) SetBody(v *TaskInfoCreateAndStartTaskResponseBody) *TaskInfoCreateAndStartTaskResponse {
	s.Body = v
	return s
}

type TaskInfoFinishTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TaskInfoFinishTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoFinishTaskHeaders) GoString() string {
	return s.String()
}

func (s *TaskInfoFinishTaskHeaders) SetCommonHeaders(v map[string]*string) *TaskInfoFinishTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TaskInfoFinishTaskHeaders) SetXAcsDingtalkAccessToken(v string) *TaskInfoFinishTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TaskInfoFinishTaskRequest struct {
	CardDTO *TaskInfoFinishTaskRequestCardDTO `json:"cardDTO,omitempty" xml:"cardDTO,omitempty" type:"Struct"`
	// This parameter is required.
	OperatorAccount *string `json:"operatorAccount,omitempty" xml:"operatorAccount,omitempty"`
	OutTaskId       *string `json:"outTaskId,omitempty" xml:"outTaskId,omitempty"`
	// This parameter is required.
	ProjId *string `json:"projId,omitempty" xml:"projId,omitempty"`
	// This parameter is required.
	SecretKey             *string                                           `json:"secretKey,omitempty" xml:"secretKey,omitempty"`
	SendMsgFlag           *int32                                            `json:"sendMsgFlag,omitempty" xml:"sendMsgFlag,omitempty"`
	TaskExecutePersonDTOS []*TaskInfoFinishTaskRequestTaskExecutePersonDTOS `json:"taskExecutePersonDTOS,omitempty" xml:"taskExecutePersonDTOS,omitempty" type:"Repeated"`
}

func (s TaskInfoFinishTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoFinishTaskRequest) GoString() string {
	return s.String()
}

func (s *TaskInfoFinishTaskRequest) SetCardDTO(v *TaskInfoFinishTaskRequestCardDTO) *TaskInfoFinishTaskRequest {
	s.CardDTO = v
	return s
}

func (s *TaskInfoFinishTaskRequest) SetOperatorAccount(v string) *TaskInfoFinishTaskRequest {
	s.OperatorAccount = &v
	return s
}

func (s *TaskInfoFinishTaskRequest) SetOutTaskId(v string) *TaskInfoFinishTaskRequest {
	s.OutTaskId = &v
	return s
}

func (s *TaskInfoFinishTaskRequest) SetProjId(v string) *TaskInfoFinishTaskRequest {
	s.ProjId = &v
	return s
}

func (s *TaskInfoFinishTaskRequest) SetSecretKey(v string) *TaskInfoFinishTaskRequest {
	s.SecretKey = &v
	return s
}

func (s *TaskInfoFinishTaskRequest) SetSendMsgFlag(v int32) *TaskInfoFinishTaskRequest {
	s.SendMsgFlag = &v
	return s
}

func (s *TaskInfoFinishTaskRequest) SetTaskExecutePersonDTOS(v []*TaskInfoFinishTaskRequestTaskExecutePersonDTOS) *TaskInfoFinishTaskRequest {
	s.TaskExecutePersonDTOS = v
	return s
}

type TaskInfoFinishTaskRequestCardDTO struct {
	AtAccount       *string     `json:"atAccount,omitempty" xml:"atAccount,omitempty"`
	CardCallbackUrl *string     `json:"cardCallbackUrl,omitempty" xml:"cardCallbackUrl,omitempty"`
	Content         interface{} `json:"content,omitempty" xml:"content,omitempty"`
	IsAtAll         *bool       `json:"isAtAll,omitempty" xml:"isAtAll,omitempty"`
	ReceiverAccount *string     `json:"receiverAccount,omitempty" xml:"receiverAccount,omitempty"`
}

func (s TaskInfoFinishTaskRequestCardDTO) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoFinishTaskRequestCardDTO) GoString() string {
	return s.String()
}

func (s *TaskInfoFinishTaskRequestCardDTO) SetAtAccount(v string) *TaskInfoFinishTaskRequestCardDTO {
	s.AtAccount = &v
	return s
}

func (s *TaskInfoFinishTaskRequestCardDTO) SetCardCallbackUrl(v string) *TaskInfoFinishTaskRequestCardDTO {
	s.CardCallbackUrl = &v
	return s
}

func (s *TaskInfoFinishTaskRequestCardDTO) SetContent(v interface{}) *TaskInfoFinishTaskRequestCardDTO {
	s.Content = v
	return s
}

func (s *TaskInfoFinishTaskRequestCardDTO) SetIsAtAll(v bool) *TaskInfoFinishTaskRequestCardDTO {
	s.IsAtAll = &v
	return s
}

func (s *TaskInfoFinishTaskRequestCardDTO) SetReceiverAccount(v string) *TaskInfoFinishTaskRequestCardDTO {
	s.ReceiverAccount = &v
	return s
}

type TaskInfoFinishTaskRequestTaskExecutePersonDTOS struct {
	EmployeeCode *string `json:"employeeCode,omitempty" xml:"employeeCode,omitempty"`
	PersonType   *int32  `json:"personType,omitempty" xml:"personType,omitempty"`
}

func (s TaskInfoFinishTaskRequestTaskExecutePersonDTOS) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoFinishTaskRequestTaskExecutePersonDTOS) GoString() string {
	return s.String()
}

func (s *TaskInfoFinishTaskRequestTaskExecutePersonDTOS) SetEmployeeCode(v string) *TaskInfoFinishTaskRequestTaskExecutePersonDTOS {
	s.EmployeeCode = &v
	return s
}

func (s *TaskInfoFinishTaskRequestTaskExecutePersonDTOS) SetPersonType(v int32) *TaskInfoFinishTaskRequestTaskExecutePersonDTOS {
	s.PersonType = &v
	return s
}

type TaskInfoFinishTaskResponseBody struct {
	Code    *int32      `json:"code,omitempty" xml:"code,omitempty"`
	Data    interface{} `json:"data,omitempty" xml:"data,omitempty"`
	Message *string     `json:"message,omitempty" xml:"message,omitempty"`
}

func (s TaskInfoFinishTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoFinishTaskResponseBody) GoString() string {
	return s.String()
}

func (s *TaskInfoFinishTaskResponseBody) SetCode(v int32) *TaskInfoFinishTaskResponseBody {
	s.Code = &v
	return s
}

func (s *TaskInfoFinishTaskResponseBody) SetData(v interface{}) *TaskInfoFinishTaskResponseBody {
	s.Data = v
	return s
}

func (s *TaskInfoFinishTaskResponseBody) SetMessage(v string) *TaskInfoFinishTaskResponseBody {
	s.Message = &v
	return s
}

type TaskInfoFinishTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TaskInfoFinishTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TaskInfoFinishTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoFinishTaskResponse) GoString() string {
	return s.String()
}

func (s *TaskInfoFinishTaskResponse) SetHeaders(v map[string]*string) *TaskInfoFinishTaskResponse {
	s.Headers = v
	return s
}

func (s *TaskInfoFinishTaskResponse) SetStatusCode(v int32) *TaskInfoFinishTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *TaskInfoFinishTaskResponse) SetBody(v *TaskInfoFinishTaskResponseBody) *TaskInfoFinishTaskResponse {
	s.Body = v
	return s
}

type TaskInfoUpdateTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TaskInfoUpdateTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoUpdateTaskHeaders) GoString() string {
	return s.String()
}

func (s *TaskInfoUpdateTaskHeaders) SetCommonHeaders(v map[string]*string) *TaskInfoUpdateTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TaskInfoUpdateTaskHeaders) SetXAcsDingtalkAccessToken(v string) *TaskInfoUpdateTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TaskInfoUpdateTaskRequest struct {
	Attr                   *TaskInfoUpdateTaskRequestAttr      `json:"attr,omitempty" xml:"attr,omitempty" type:"Struct"`
	CanceldelTaskCardId    *string                             `json:"canceldelTaskCardId,omitempty" xml:"canceldelTaskCardId,omitempty"`
	CardDTO                *TaskInfoUpdateTaskRequestCardDTO   `json:"cardDTO,omitempty" xml:"cardDTO,omitempty" type:"Struct"`
	DetailUrl              *TaskInfoUpdateTaskRequestDetailUrl `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	FinishTaskCardId       *string                             `json:"finishTaskCardId,omitempty" xml:"finishTaskCardId,omitempty"`
	ListOpenConversationId []*string                           `json:"listOpenConversationId,omitempty" xml:"listOpenConversationId,omitempty" type:"Repeated"`
	OperateType            *int32                              `json:"operateType,omitempty" xml:"operateType,omitempty"`
	// This parameter is required.
	OperatorAccount *string `json:"operatorAccount,omitempty" xml:"operatorAccount,omitempty"`
	OutTaskId       *string `json:"outTaskId,omitempty" xml:"outTaskId,omitempty"`
	// This parameter is required.
	ProjId *string `json:"projId,omitempty" xml:"projId,omitempty"`
	// This parameter is required.
	SecretKey             *string                                           `json:"secretKey,omitempty" xml:"secretKey,omitempty"`
	SendMsgFlag           *int32                                            `json:"sendMsgFlag,omitempty" xml:"sendMsgFlag,omitempty"`
	StartTaskCardId       *string                                           `json:"startTaskCardId,omitempty" xml:"startTaskCardId,omitempty"`
	TaskContent           *string                                           `json:"taskContent,omitempty" xml:"taskContent,omitempty"`
	TaskEndTime           *int64                                            `json:"taskEndTime,omitempty" xml:"taskEndTime,omitempty"`
	TaskExecutePersonDTOS []*TaskInfoUpdateTaskRequestTaskExecutePersonDTOS `json:"taskExecutePersonDTOS,omitempty" xml:"taskExecutePersonDTOS,omitempty" type:"Repeated"`
	TaskTitle             *string                                           `json:"taskTitle,omitempty" xml:"taskTitle,omitempty"`
	TaskUrlMobile         *string                                           `json:"taskUrlMobile,omitempty" xml:"taskUrlMobile,omitempty"`
	TaskUrlPc             *string                                           `json:"taskUrlPc,omitempty" xml:"taskUrlPc,omitempty"`
	UpdateTaskCardId      *string                                           `json:"updateTaskCardId,omitempty" xml:"updateTaskCardId,omitempty"`
}

func (s TaskInfoUpdateTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoUpdateTaskRequest) GoString() string {
	return s.String()
}

func (s *TaskInfoUpdateTaskRequest) SetAttr(v *TaskInfoUpdateTaskRequestAttr) *TaskInfoUpdateTaskRequest {
	s.Attr = v
	return s
}

func (s *TaskInfoUpdateTaskRequest) SetCanceldelTaskCardId(v string) *TaskInfoUpdateTaskRequest {
	s.CanceldelTaskCardId = &v
	return s
}

func (s *TaskInfoUpdateTaskRequest) SetCardDTO(v *TaskInfoUpdateTaskRequestCardDTO) *TaskInfoUpdateTaskRequest {
	s.CardDTO = v
	return s
}

func (s *TaskInfoUpdateTaskRequest) SetDetailUrl(v *TaskInfoUpdateTaskRequestDetailUrl) *TaskInfoUpdateTaskRequest {
	s.DetailUrl = v
	return s
}

func (s *TaskInfoUpdateTaskRequest) SetFinishTaskCardId(v string) *TaskInfoUpdateTaskRequest {
	s.FinishTaskCardId = &v
	return s
}

func (s *TaskInfoUpdateTaskRequest) SetListOpenConversationId(v []*string) *TaskInfoUpdateTaskRequest {
	s.ListOpenConversationId = v
	return s
}

func (s *TaskInfoUpdateTaskRequest) SetOperateType(v int32) *TaskInfoUpdateTaskRequest {
	s.OperateType = &v
	return s
}

func (s *TaskInfoUpdateTaskRequest) SetOperatorAccount(v string) *TaskInfoUpdateTaskRequest {
	s.OperatorAccount = &v
	return s
}

func (s *TaskInfoUpdateTaskRequest) SetOutTaskId(v string) *TaskInfoUpdateTaskRequest {
	s.OutTaskId = &v
	return s
}

func (s *TaskInfoUpdateTaskRequest) SetProjId(v string) *TaskInfoUpdateTaskRequest {
	s.ProjId = &v
	return s
}

func (s *TaskInfoUpdateTaskRequest) SetSecretKey(v string) *TaskInfoUpdateTaskRequest {
	s.SecretKey = &v
	return s
}

func (s *TaskInfoUpdateTaskRequest) SetSendMsgFlag(v int32) *TaskInfoUpdateTaskRequest {
	s.SendMsgFlag = &v
	return s
}

func (s *TaskInfoUpdateTaskRequest) SetStartTaskCardId(v string) *TaskInfoUpdateTaskRequest {
	s.StartTaskCardId = &v
	return s
}

func (s *TaskInfoUpdateTaskRequest) SetTaskContent(v string) *TaskInfoUpdateTaskRequest {
	s.TaskContent = &v
	return s
}

func (s *TaskInfoUpdateTaskRequest) SetTaskEndTime(v int64) *TaskInfoUpdateTaskRequest {
	s.TaskEndTime = &v
	return s
}

func (s *TaskInfoUpdateTaskRequest) SetTaskExecutePersonDTOS(v []*TaskInfoUpdateTaskRequestTaskExecutePersonDTOS) *TaskInfoUpdateTaskRequest {
	s.TaskExecutePersonDTOS = v
	return s
}

func (s *TaskInfoUpdateTaskRequest) SetTaskTitle(v string) *TaskInfoUpdateTaskRequest {
	s.TaskTitle = &v
	return s
}

func (s *TaskInfoUpdateTaskRequest) SetTaskUrlMobile(v string) *TaskInfoUpdateTaskRequest {
	s.TaskUrlMobile = &v
	return s
}

func (s *TaskInfoUpdateTaskRequest) SetTaskUrlPc(v string) *TaskInfoUpdateTaskRequest {
	s.TaskUrlPc = &v
	return s
}

func (s *TaskInfoUpdateTaskRequest) SetUpdateTaskCardId(v string) *TaskInfoUpdateTaskRequest {
	s.UpdateTaskCardId = &v
	return s
}

type TaskInfoUpdateTaskRequestAttr struct {
	ListTaskDynamicAttr []*TaskInfoUpdateTaskRequestAttrListTaskDynamicAttr `json:"listTaskDynamicAttr,omitempty" xml:"listTaskDynamicAttr,omitempty" type:"Repeated"`
}

func (s TaskInfoUpdateTaskRequestAttr) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoUpdateTaskRequestAttr) GoString() string {
	return s.String()
}

func (s *TaskInfoUpdateTaskRequestAttr) SetListTaskDynamicAttr(v []*TaskInfoUpdateTaskRequestAttrListTaskDynamicAttr) *TaskInfoUpdateTaskRequestAttr {
	s.ListTaskDynamicAttr = v
	return s
}

type TaskInfoUpdateTaskRequestAttrListTaskDynamicAttr struct {
	AttrCode            *string   `json:"attrCode,omitempty" xml:"attrCode,omitempty"`
	ListAttrOptionsCode []*string `json:"listAttrOptionsCode,omitempty" xml:"listAttrOptionsCode,omitempty" type:"Repeated"`
}

func (s TaskInfoUpdateTaskRequestAttrListTaskDynamicAttr) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoUpdateTaskRequestAttrListTaskDynamicAttr) GoString() string {
	return s.String()
}

func (s *TaskInfoUpdateTaskRequestAttrListTaskDynamicAttr) SetAttrCode(v string) *TaskInfoUpdateTaskRequestAttrListTaskDynamicAttr {
	s.AttrCode = &v
	return s
}

func (s *TaskInfoUpdateTaskRequestAttrListTaskDynamicAttr) SetListAttrOptionsCode(v []*string) *TaskInfoUpdateTaskRequestAttrListTaskDynamicAttr {
	s.ListAttrOptionsCode = v
	return s
}

type TaskInfoUpdateTaskRequestCardDTO struct {
	AtAccount       *string     `json:"atAccount,omitempty" xml:"atAccount,omitempty"`
	CardCallbackUrl *string     `json:"cardCallbackUrl,omitempty" xml:"cardCallbackUrl,omitempty"`
	Content         interface{} `json:"content,omitempty" xml:"content,omitempty"`
	IsAtAll         *bool       `json:"isAtAll,omitempty" xml:"isAtAll,omitempty"`
	ReceiverAccount *string     `json:"receiverAccount,omitempty" xml:"receiverAccount,omitempty"`
}

func (s TaskInfoUpdateTaskRequestCardDTO) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoUpdateTaskRequestCardDTO) GoString() string {
	return s.String()
}

func (s *TaskInfoUpdateTaskRequestCardDTO) SetAtAccount(v string) *TaskInfoUpdateTaskRequestCardDTO {
	s.AtAccount = &v
	return s
}

func (s *TaskInfoUpdateTaskRequestCardDTO) SetCardCallbackUrl(v string) *TaskInfoUpdateTaskRequestCardDTO {
	s.CardCallbackUrl = &v
	return s
}

func (s *TaskInfoUpdateTaskRequestCardDTO) SetContent(v interface{}) *TaskInfoUpdateTaskRequestCardDTO {
	s.Content = v
	return s
}

func (s *TaskInfoUpdateTaskRequestCardDTO) SetIsAtAll(v bool) *TaskInfoUpdateTaskRequestCardDTO {
	s.IsAtAll = &v
	return s
}

func (s *TaskInfoUpdateTaskRequestCardDTO) SetReceiverAccount(v string) *TaskInfoUpdateTaskRequestCardDTO {
	s.ReceiverAccount = &v
	return s
}

type TaskInfoUpdateTaskRequestDetailUrl struct {
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
	PcUrl  *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
}

func (s TaskInfoUpdateTaskRequestDetailUrl) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoUpdateTaskRequestDetailUrl) GoString() string {
	return s.String()
}

func (s *TaskInfoUpdateTaskRequestDetailUrl) SetAppUrl(v string) *TaskInfoUpdateTaskRequestDetailUrl {
	s.AppUrl = &v
	return s
}

func (s *TaskInfoUpdateTaskRequestDetailUrl) SetPcUrl(v string) *TaskInfoUpdateTaskRequestDetailUrl {
	s.PcUrl = &v
	return s
}

type TaskInfoUpdateTaskRequestTaskExecutePersonDTOS struct {
	EmployeeCode *string `json:"employeeCode,omitempty" xml:"employeeCode,omitempty"`
	PersonType   *int32  `json:"personType,omitempty" xml:"personType,omitempty"`
}

func (s TaskInfoUpdateTaskRequestTaskExecutePersonDTOS) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoUpdateTaskRequestTaskExecutePersonDTOS) GoString() string {
	return s.String()
}

func (s *TaskInfoUpdateTaskRequestTaskExecutePersonDTOS) SetEmployeeCode(v string) *TaskInfoUpdateTaskRequestTaskExecutePersonDTOS {
	s.EmployeeCode = &v
	return s
}

func (s *TaskInfoUpdateTaskRequestTaskExecutePersonDTOS) SetPersonType(v int32) *TaskInfoUpdateTaskRequestTaskExecutePersonDTOS {
	s.PersonType = &v
	return s
}

type TaskInfoUpdateTaskResponseBody struct {
	Code    *int32                              `json:"code,omitempty" xml:"code,omitempty"`
	Data    *TaskInfoUpdateTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                             `json:"message,omitempty" xml:"message,omitempty"`
}

func (s TaskInfoUpdateTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoUpdateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *TaskInfoUpdateTaskResponseBody) SetCode(v int32) *TaskInfoUpdateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *TaskInfoUpdateTaskResponseBody) SetData(v *TaskInfoUpdateTaskResponseBodyData) *TaskInfoUpdateTaskResponseBody {
	s.Data = v
	return s
}

func (s *TaskInfoUpdateTaskResponseBody) SetMessage(v string) *TaskInfoUpdateTaskResponseBody {
	s.Message = &v
	return s
}

type TaskInfoUpdateTaskResponseBodyData struct {
	GroupVoList []*TaskInfoUpdateTaskResponseBodyDataGroupVoList `json:"groupVoList,omitempty" xml:"groupVoList,omitempty" type:"Repeated"`
	TaskId      *string                                          `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s TaskInfoUpdateTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoUpdateTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *TaskInfoUpdateTaskResponseBodyData) SetGroupVoList(v []*TaskInfoUpdateTaskResponseBodyDataGroupVoList) *TaskInfoUpdateTaskResponseBodyData {
	s.GroupVoList = v
	return s
}

func (s *TaskInfoUpdateTaskResponseBodyData) SetTaskId(v string) *TaskInfoUpdateTaskResponseBodyData {
	s.TaskId = &v
	return s
}

type TaskInfoUpdateTaskResponseBodyDataGroupVoList struct {
	CorpId             *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s TaskInfoUpdateTaskResponseBodyDataGroupVoList) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoUpdateTaskResponseBodyDataGroupVoList) GoString() string {
	return s.String()
}

func (s *TaskInfoUpdateTaskResponseBodyDataGroupVoList) SetCorpId(v string) *TaskInfoUpdateTaskResponseBodyDataGroupVoList {
	s.CorpId = &v
	return s
}

func (s *TaskInfoUpdateTaskResponseBodyDataGroupVoList) SetOpenConversationId(v string) *TaskInfoUpdateTaskResponseBodyDataGroupVoList {
	s.OpenConversationId = &v
	return s
}

type TaskInfoUpdateTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TaskInfoUpdateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TaskInfoUpdateTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s TaskInfoUpdateTaskResponse) GoString() string {
	return s.String()
}

func (s *TaskInfoUpdateTaskResponse) SetHeaders(v map[string]*string) *TaskInfoUpdateTaskResponse {
	s.Headers = v
	return s
}

func (s *TaskInfoUpdateTaskResponse) SetStatusCode(v int32) *TaskInfoUpdateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *TaskInfoUpdateTaskResponse) SetBody(v *TaskInfoUpdateTaskResponseBody) *TaskInfoUpdateTaskResponse {
	s.Body = v
	return s
}

type TransferExclusiveAccountOrgHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TransferExclusiveAccountOrgHeaders) String() string {
	return tea.Prettify(s)
}

func (s TransferExclusiveAccountOrgHeaders) GoString() string {
	return s.String()
}

func (s *TransferExclusiveAccountOrgHeaders) SetCommonHeaders(v map[string]*string) *TransferExclusiveAccountOrgHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TransferExclusiveAccountOrgHeaders) SetXAcsDingtalkAccessToken(v string) *TransferExclusiveAccountOrgHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TransferExclusiveAccountOrgRequest struct {
	IsSettingMainOrg *bool   `json:"isSettingMainOrg,omitempty" xml:"isSettingMainOrg,omitempty"`
	TargetCorpId     *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
	// This parameter is required.
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s TransferExclusiveAccountOrgRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferExclusiveAccountOrgRequest) GoString() string {
	return s.String()
}

func (s *TransferExclusiveAccountOrgRequest) SetIsSettingMainOrg(v bool) *TransferExclusiveAccountOrgRequest {
	s.IsSettingMainOrg = &v
	return s
}

func (s *TransferExclusiveAccountOrgRequest) SetTargetCorpId(v string) *TransferExclusiveAccountOrgRequest {
	s.TargetCorpId = &v
	return s
}

func (s *TransferExclusiveAccountOrgRequest) SetUserIds(v []*string) *TransferExclusiveAccountOrgRequest {
	s.UserIds = v
	return s
}

type TransferExclusiveAccountOrgResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TransferExclusiveAccountOrgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferExclusiveAccountOrgResponseBody) GoString() string {
	return s.String()
}

func (s *TransferExclusiveAccountOrgResponseBody) SetSuccess(v bool) *TransferExclusiveAccountOrgResponseBody {
	s.Success = &v
	return s
}

type TransferExclusiveAccountOrgResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferExclusiveAccountOrgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferExclusiveAccountOrgResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferExclusiveAccountOrgResponse) GoString() string {
	return s.String()
}

func (s *TransferExclusiveAccountOrgResponse) SetHeaders(v map[string]*string) *TransferExclusiveAccountOrgResponse {
	s.Headers = v
	return s
}

func (s *TransferExclusiveAccountOrgResponse) SetStatusCode(v int32) *TransferExclusiveAccountOrgResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferExclusiveAccountOrgResponse) SetBody(v *TransferExclusiveAccountOrgResponseBody) *TransferExclusiveAccountOrgResponse {
	s.Body = v
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
	// example:
	//
	// demo
	CurrentCategoryName *string `json:"currentCategoryName,omitempty" xml:"currentCategoryName,omitempty"`
	// example:
	//
	// demo01
	TargetCategoryName *string `json:"targetCategoryName,omitempty" xml:"targetCategoryName,omitempty"`
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
	// example:
	//
	// 1
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCategoryNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type UpdateConversationTypeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateConversationTypeHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateConversationTypeHeaders) GoString() string {
	return s.String()
}

func (s *UpdateConversationTypeHeaders) SetCommonHeaders(v map[string]*string) *UpdateConversationTypeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateConversationTypeHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateConversationTypeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateConversationTypeRequest struct {
	// This parameter is required.
	ManageSign *int32 `json:"manageSign,omitempty" xml:"manageSign,omitempty"`
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s UpdateConversationTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConversationTypeRequest) GoString() string {
	return s.String()
}

func (s *UpdateConversationTypeRequest) SetManageSign(v int32) *UpdateConversationTypeRequest {
	s.ManageSign = &v
	return s
}

func (s *UpdateConversationTypeRequest) SetOpenConversationId(v string) *UpdateConversationTypeRequest {
	s.OpenConversationId = &v
	return s
}

type UpdateConversationTypeResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateConversationTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateConversationTypeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConversationTypeResponseBody) SetResult(v string) *UpdateConversationTypeResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateConversationTypeResponseBody) SetSuccess(v bool) *UpdateConversationTypeResponseBody {
	s.Success = &v
	return s
}

type UpdateConversationTypeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConversationTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConversationTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateConversationTypeResponse) GoString() string {
	return s.String()
}

func (s *UpdateConversationTypeResponse) SetHeaders(v map[string]*string) *UpdateConversationTypeResponse {
	s.Headers = v
	return s
}

func (s *UpdateConversationTypeResponse) SetStatusCode(v int32) *UpdateConversationTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConversationTypeResponse) SetBody(v *UpdateConversationTypeResponseBody) *UpdateConversationTypeResponse {
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
	// This parameter is required.
	RequestIds []*string `json:"requestIds,omitempty" xml:"requestIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1-检测通过，2-检测失败
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// true
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFileStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 500000003
	MiniAppId *string `json:"miniAppId,omitempty" xml:"miniAppId,omitempty"`
	// example:
	//
	// 0.0.5
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0-开发版，1-灰度版，2-发布版，3-体验版
	VersionType *int32 `json:"versionType,omitempty" xml:"versionType,omitempty"`
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
	// example:
	//
	// 成功
	Cause *string `json:"cause,omitempty" xml:"cause,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMiniAppVersionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 0.0.5
	DeptIds []*int64 `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1312312
	LabelId *int64 `json:"labelId,omitempty" xml:"labelId,omitempty"`
	// example:
	//
	// 500000003
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *bool              `json:"body,omitempty" xml:"body,omitempty"`
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

type UpdateRealmLicenseHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateRealmLicenseHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateRealmLicenseHeaders) GoString() string {
	return s.String()
}

func (s *UpdateRealmLicenseHeaders) SetCommonHeaders(v map[string]*string) *UpdateRealmLicenseHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateRealmLicenseHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateRealmLicenseHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateRealmLicenseRequest struct {
	DetailList []*UpdateRealmLicenseRequestDetailList `json:"detailList,omitempty" xml:"detailList,omitempty" type:"Repeated"`
}

func (s UpdateRealmLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRealmLicenseRequest) GoString() string {
	return s.String()
}

func (s *UpdateRealmLicenseRequest) SetDetailList(v []*UpdateRealmLicenseRequestDetailList) *UpdateRealmLicenseRequest {
	s.DetailList = v
	return s
}

type UpdateRealmLicenseRequestDetailList struct {
	// example:
	//
	// 1001
	LicenseType *int32 `json:"licenseType,omitempty" xml:"licenseType,omitempty"`
	// example:
	//
	// 123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateRealmLicenseRequestDetailList) String() string {
	return tea.Prettify(s)
}

func (s UpdateRealmLicenseRequestDetailList) GoString() string {
	return s.String()
}

func (s *UpdateRealmLicenseRequestDetailList) SetLicenseType(v int32) *UpdateRealmLicenseRequestDetailList {
	s.LicenseType = &v
	return s
}

func (s *UpdateRealmLicenseRequestDetailList) SetUserId(v string) *UpdateRealmLicenseRequestDetailList {
	s.UserId = &v
	return s
}

type UpdateRealmLicenseResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateRealmLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRealmLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRealmLicenseResponseBody) SetResult(v bool) *UpdateRealmLicenseResponseBody {
	s.Result = &v
	return s
}

type UpdateRealmLicenseResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRealmLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRealmLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRealmLicenseResponse) GoString() string {
	return s.String()
}

func (s *UpdateRealmLicenseResponse) SetHeaders(v map[string]*string) *UpdateRealmLicenseResponse {
	s.Headers = v
	return s
}

func (s *UpdateRealmLicenseResponse) SetStatusCode(v int32) *UpdateRealmLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRealmLicenseResponse) SetBody(v *UpdateRealmLicenseResponseBody) *UpdateRealmLicenseResponse {
	s.Body = v
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
	// example:
	//
	// 0.0.5
	DeptIds []*int64 `json:"deptIds,omitempty" xml:"deptIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1312312
	LabelId *int64 `json:"labelId,omitempty" xml:"labelId,omitempty"`
	// example:
	//
	// 500000003
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *bool              `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 0
	FileStorageMode *string `json:"fileStorageMode,omitempty" xml:"fileStorageMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dingxxxxxxxxxxxx
	TargetCorpId *string `json:"targetCorpId,omitempty" xml:"targetCorpId,omitempty"`
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
	// example:
	//
	// ding66b0e9d003fc91ed35c2f4xxxxxxxxxx
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStorageModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type UpdateTrustedDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateTrustedDeviceHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrustedDeviceHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTrustedDeviceHeaders) SetCommonHeaders(v map[string]*string) *UpdateTrustedDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTrustedDeviceHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateTrustedDeviceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateTrustedDeviceRequest struct {
	Status *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Title  *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s UpdateTrustedDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrustedDeviceRequest) GoString() string {
	return s.String()
}

func (s *UpdateTrustedDeviceRequest) SetStatus(v int32) *UpdateTrustedDeviceRequest {
	s.Status = &v
	return s
}

func (s *UpdateTrustedDeviceRequest) SetTitle(v string) *UpdateTrustedDeviceRequest {
	s.Title = &v
	return s
}

type UpdateTrustedDeviceResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateTrustedDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrustedDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTrustedDeviceResponseBody) SetSuccess(v bool) *UpdateTrustedDeviceResponseBody {
	s.Success = &v
	return s
}

type UpdateTrustedDeviceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTrustedDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTrustedDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTrustedDeviceResponse) GoString() string {
	return s.String()
}

func (s *UpdateTrustedDeviceResponse) SetHeaders(v map[string]*string) *UpdateTrustedDeviceResponse {
	s.Headers = v
	return s
}

func (s *UpdateTrustedDeviceResponse) SetStatusCode(v int32) *UpdateTrustedDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTrustedDeviceResponse) SetBody(v *UpdateTrustedDeviceResponseBody) *UpdateTrustedDeviceResponse {
	s.Body = v
	return s
}

type UpdateVoiceMsgCtrlStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateVoiceMsgCtrlStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateVoiceMsgCtrlStatusHeaders) GoString() string {
	return s.String()
}

func (s *UpdateVoiceMsgCtrlStatusHeaders) SetCommonHeaders(v map[string]*string) *UpdateVoiceMsgCtrlStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateVoiceMsgCtrlStatusHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateVoiceMsgCtrlStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateVoiceMsgCtrlStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1-检测通过，2-检测失败
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	VoiceMsgCtrlInfo *UpdateVoiceMsgCtrlStatusRequestVoiceMsgCtrlInfo `json:"voiceMsgCtrlInfo,omitempty" xml:"voiceMsgCtrlInfo,omitempty" type:"Struct"`
}

func (s UpdateVoiceMsgCtrlStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVoiceMsgCtrlStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateVoiceMsgCtrlStatusRequest) SetStatus(v int32) *UpdateVoiceMsgCtrlStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateVoiceMsgCtrlStatusRequest) SetVoiceMsgCtrlInfo(v *UpdateVoiceMsgCtrlStatusRequestVoiceMsgCtrlInfo) *UpdateVoiceMsgCtrlStatusRequest {
	s.VoiceMsgCtrlInfo = v
	return s
}

type UpdateVoiceMsgCtrlStatusRequestVoiceMsgCtrlInfo struct {
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	OpenMsgId *string `json:"openMsgId,omitempty" xml:"openMsgId,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateVoiceMsgCtrlStatusRequestVoiceMsgCtrlInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateVoiceMsgCtrlStatusRequestVoiceMsgCtrlInfo) GoString() string {
	return s.String()
}

func (s *UpdateVoiceMsgCtrlStatusRequestVoiceMsgCtrlInfo) SetOpenConversationId(v string) *UpdateVoiceMsgCtrlStatusRequestVoiceMsgCtrlInfo {
	s.OpenConversationId = &v
	return s
}

func (s *UpdateVoiceMsgCtrlStatusRequestVoiceMsgCtrlInfo) SetOpenMsgId(v string) *UpdateVoiceMsgCtrlStatusRequestVoiceMsgCtrlInfo {
	s.OpenMsgId = &v
	return s
}

func (s *UpdateVoiceMsgCtrlStatusRequestVoiceMsgCtrlInfo) SetUserId(v string) *UpdateVoiceMsgCtrlStatusRequestVoiceMsgCtrlInfo {
	s.UserId = &v
	return s
}

type UpdateVoiceMsgCtrlStatusResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateVoiceMsgCtrlStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVoiceMsgCtrlStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVoiceMsgCtrlStatusResponseBody) SetSuccess(v bool) *UpdateVoiceMsgCtrlStatusResponseBody {
	s.Success = &v
	return s
}

type UpdateVoiceMsgCtrlStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVoiceMsgCtrlStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVoiceMsgCtrlStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVoiceMsgCtrlStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateVoiceMsgCtrlStatusResponse) SetHeaders(v map[string]*string) *UpdateVoiceMsgCtrlStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateVoiceMsgCtrlStatusResponse) SetStatusCode(v int32) *UpdateVoiceMsgCtrlStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVoiceMsgCtrlStatusResponse) SetBody(v *UpdateVoiceMsgCtrlStatusResponseBody) *UpdateVoiceMsgCtrlStatusResponse {
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
// 添加自主协议
//
// @param request - AddCustomSignConfigRequest
//
// @param headers - AddCustomSignConfigHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCustomSignConfigResponse
func (client *Client) AddCustomSignConfigWithOptions(request *AddCustomSignConfigRequest, headers *AddCustomSignConfigHeaders, runtime *util.RuntimeOptions) (_result *AddCustomSignConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllEffect)) {
		body["allEffect"] = request.AllEffect
	}

	if !tea.BoolValue(util.IsUnset(request.CanDownload)) {
		body["canDownload"] = request.CanDownload
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolName)) {
		body["protocolName"] = request.ProtocolName
	}

	if !tea.BoolValue(util.IsUnset(request.PushDeptIds)) {
		body["pushDeptIds"] = request.PushDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.PushStaffIds)) {
		body["pushStaffIds"] = request.PushStaffIds
	}

	if !tea.BoolValue(util.IsUnset(request.SignTermFiles)) {
		body["signTermFiles"] = request.SignTermFiles
	}

	if !tea.BoolValue(util.IsUnset(request.TermMessage)) {
		body["termMessage"] = request.TermMessage
	}

	if !tea.BoolValue(util.IsUnset(request.UnpushDeptIds)) {
		body["unpushDeptIds"] = request.UnpushDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.UnpushStaffIds)) {
		body["unpushStaffIds"] = request.UnpushStaffIds
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
		Action:      tea.String("AddCustomSignConfig"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/sign/addCustomSignConfig"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddCustomSignConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加自主协议
//
// @param request - AddCustomSignConfigRequest
//
// @return AddCustomSignConfigResponse
func (client *Client) AddCustomSignConfig(request *AddCustomSignConfigRequest) (_result *AddCustomSignConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddCustomSignConfigHeaders{}
	_result = &AddCustomSignConfigResponse{}
	_body, _err := client.AddCustomSignConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增企业
//
// @param request - AddOrgRequest
//
// @param headers - AddOrgHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddOrgResponse
func (client *Client) AddOrgWithOptions(request *AddOrgRequest, headers *AddOrgHeaders, runtime *util.RuntimeOptions) (_result *AddOrgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.City)) {
		body["city"] = request.City
	}

	if !tea.BoolValue(util.IsUnset(request.Industry)) {
		body["industry"] = request.Industry
	}

	if !tea.BoolValue(util.IsUnset(request.IndustryCode)) {
		body["industryCode"] = request.IndustryCode
	}

	if !tea.BoolValue(util.IsUnset(request.MobileNum)) {
		body["mobileNum"] = request.MobileNum
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Province)) {
		body["province"] = request.Province
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

// Summary:
//
// 新增企业
//
// @param request - AddOrgRequest
//
// @return AddOrgResponse
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

// Summary:
//
// 专属审批结果回调
//
// @param request - ApproveProcessCallbackRequest
//
// @param headers - ApproveProcessCallbackHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApproveProcessCallbackResponse
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

// Summary:
//
// 专属审批结果回调
//
// @param request - ApproveProcessCallbackRequest
//
// @return ApproveProcessCallbackResponse
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

// Summary:
//
// 群禁言或解禁
//
// @param request - BanOrOpenGroupWordsRequest
//
// @param headers - BanOrOpenGroupWordsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BanOrOpenGroupWordsResponse
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

// Summary:
//
// 群禁言或解禁
//
// @param request - BanOrOpenGroupWordsRequest
//
// @return BanOrOpenGroupWordsResponse
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

// Summary:
//
// 业务融合群业务事件变更
//
// @param request - BusinessEventUpdateRequest
//
// @param headers - BusinessEventUpdateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BusinessEventUpdateResponse
func (client *Client) BusinessEventUpdateWithOptions(request *BusinessEventUpdateRequest, headers *BusinessEventUpdateHeaders, runtime *util.RuntimeOptions) (_result *BusinessEventUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessData)) {
		body["businessData"] = request.BusinessData
	}

	if !tea.BoolValue(util.IsUnset(request.EventType)) {
		body["eventType"] = request.EventType
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateByKey)) {
		body["updateByKey"] = request.UpdateByKey
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
		Action:      tea.String("BusinessEventUpdate"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/businessGroups/events"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BusinessEventUpdateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 业务融合群业务事件变更
//
// @param request - BusinessEventUpdateRequest
//
// @return BusinessEventUpdateResponse
func (client *Client) BusinessEventUpdate(request *BusinessEventUpdateRequest) (_result *BusinessEventUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BusinessEventUpdateHeaders{}
	_result = &BusinessEventUpdateResponse{}
	_body, _err := client.BusinessEventUpdateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 检测安全管控功能命中状态
//
// @param request - CheckControlHitStatusRequest
//
// @param headers - CheckControlHitStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckControlHitStatusResponse
func (client *Client) CheckControlHitStatusWithOptions(request *CheckControlHitStatusRequest, headers *CheckControlHitStatusHeaders, runtime *util.RuntimeOptions) (_result *CheckControlHitStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedMissedFunction)) {
		query["needMissedFunction"] = request.NeedMissedFunction
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
		Action:      tea.String("CheckControlHitStatus"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/soc/functionHitStatuses/check"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckControlHitStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检测安全管控功能命中状态
//
// @param request - CheckControlHitStatusRequest
//
// @return CheckControlHitStatusResponse
func (client *Client) CheckControlHitStatus(request *CheckControlHitStatusRequest) (_result *CheckControlHitStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CheckControlHitStatusHeaders{}
	_result = &CheckControlHitStatusResponse{}
	_body, _err := client.CheckControlHitStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建分组并绑定会话
//
// @param request - CreateCategoryAndBindingGroupsRequest
//
// @param headers - CreateCategoryAndBindingGroupsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCategoryAndBindingGroupsResponse
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

// Summary:
//
// 创建分组并绑定会话
//
// @param request - CreateCategoryAndBindingGroupsRequest
//
// @return CreateCategoryAndBindingGroupsResponse
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

// Summary:
//
// 创建文件检测任务
//
// @param request - CreateDlpTaskRequest
//
// @param headers - CreateDlpTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDlpTaskResponse
func (client *Client) CreateDlpTaskWithOptions(request *CreateDlpTaskRequest, headers *CreateDlpTaskHeaders, runtime *util.RuntimeOptions) (_result *CreateDlpTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DentryId)) {
		body["dentryId"] = request.DentryId
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["spaceId"] = request.SpaceId
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
		Action:      tea.String("CreateDlpTask"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/dlpTasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDlpTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建文件检测任务
//
// @param request - CreateDlpTaskRequest
//
// @return CreateDlpTaskResponse
func (client *Client) CreateDlpTask(request *CreateDlpTaskRequest) (_result *CreateDlpTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateDlpTaskHeaders{}
	_result = &CreateDlpTaskResponse{}
	_body, _err := client.CreateDlpTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建分组并绑定会话
//
// @param request - CreateMessageCategoryRequest
//
// @param headers - CreateMessageCategoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMessageCategoryResponse
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

// Summary:
//
// 创建分组并绑定会话
//
// @param request - CreateMessageCategoryRequest
//
// @return CreateMessageCategoryResponse
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

// Summary:
//
// 创建规则
//
// @param request - CreateRuleRequest
//
// @param headers - CreateRuleHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRuleResponse
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

// Summary:
//
// 创建规则
//
// @param request - CreateRuleRequest
//
// @return CreateRuleResponse
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

// Summary:
//
// 存入可信设备信息
//
// @param request - CreateTrustedDeviceRequest
//
// @param headers - CreateTrustedDeviceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTrustedDeviceResponse
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

	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		body["serialNumber"] = request.SerialNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
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

// Summary:
//
// 存入可信设备信息
//
// @param request - CreateTrustedDeviceRequest
//
// @return CreateTrustedDeviceResponse
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

// Summary:
//
// 批量新增可信设备
//
// @param request - CreateTrustedDeviceBatchRequest
//
// @param headers - CreateTrustedDeviceBatchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTrustedDeviceBatchResponse
func (client *Client) CreateTrustedDeviceBatchWithOptions(request *CreateTrustedDeviceBatchRequest, headers *CreateTrustedDeviceBatchHeaders, runtime *util.RuntimeOptions) (_result *CreateTrustedDeviceBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DetailList)) {
		body["detailList"] = request.DetailList
	}

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

// Summary:
//
// 批量新增可信设备
//
// @param request - CreateTrustedDeviceBatchRequest
//
// @return CreateTrustedDeviceBatchResponse
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

// Summary:
//
// 触发文件病毒扫描任务
//
// @param request - CreateVirusScanTaskRequest
//
// @param headers - CreateVirusScanTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVirusScanTaskResponse
func (client *Client) CreateVirusScanTaskWithOptions(request *CreateVirusScanTaskRequest, headers *CreateVirusScanTaskHeaders, runtime *util.RuntimeOptions) (_result *CreateVirusScanTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DentryId)) {
		body["dentryId"] = request.DentryId
	}

	if !tea.BoolValue(util.IsUnset(request.DownloadUrl)) {
		body["downloadUrl"] = request.DownloadUrl
	}

	if !tea.BoolValue(util.IsUnset(request.FileMd5)) {
		body["fileMd5"] = request.FileMd5
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileSize)) {
		body["fileSize"] = request.FileSize
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["spaceId"] = request.SpaceId
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
		Action:      tea.String("CreateVirusScanTask"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/virusScanTasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVirusScanTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 触发文件病毒扫描任务
//
// @param request - CreateVirusScanTaskRequest
//
// @return CreateVirusScanTaskResponse
func (client *Client) CreateVirusScanTask(request *CreateVirusScanTaskRequest) (_result *CreateVirusScanTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateVirusScanTaskHeaders{}
	_result = &CreateVirusScanTaskResponse{}
	_body, _err := client.CreateVirusScanTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 为应用同步数据到专属存储
//
// @param request - DataSyncRequest
//
// @param headers - DataSyncHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DataSyncResponse
func (client *Client) DataSyncWithOptions(request *DataSyncRequest, headers *DataSyncHeaders, runtime *util.RuntimeOptions) (_result *DataSyncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Sql)) {
		body["sql"] = request.Sql
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
		Action:      tea.String("DataSync"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/datas/sync"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DataSyncResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 为应用同步数据到专属存储
//
// @param request - DataSyncRequest
//
// @return DataSyncResponse
func (client *Client) DataSync(request *DataSyncRequest) (_result *DataSyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DataSyncHeaders{}
	_result = &DataSyncResponse{}
	_body, _err := client.DataSyncWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除跨云存储配置
//
// @param headers - DeleteAcrossCloudStroageConfigsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAcrossCloudStroageConfigsResponse
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

// Summary:
//
// 删除跨云存储配置
//
// @return DeleteAcrossCloudStroageConfigsResponse
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

// Summary:
//
// 删除评论
//
// @param headers - DeleteCommentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCommentResponse
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

// Summary:
//
// 删除评论
//
// @return DeleteCommentResponse
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

// Summary:
//
// 删除指定可信设备
//
// @param request - DeleteTrustedDeviceRequest
//
// @param headers - DeleteTrustedDeviceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTrustedDeviceResponse
func (client *Client) DeleteTrustedDeviceWithOptions(request *DeleteTrustedDeviceRequest, headers *DeleteTrustedDeviceHeaders, runtime *util.RuntimeOptions) (_result *DeleteTrustedDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

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

// Summary:
//
// 删除指定可信设备
//
// @param request - DeleteTrustedDeviceRequest
//
// @return DeleteTrustedDeviceResponse
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

// Summary:
//
// 分发伙伴应用
//
// @param request - DistributePartnerAppRequest
//
// @param headers - DistributePartnerAppHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DistributePartnerAppResponse
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

// Summary:
//
// 分发伙伴应用
//
// @param request - DistributePartnerAppRequest
//
// @return DistributePartnerAppResponse
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

// Summary:
//
// 编辑安全卡片管控成员
//
// @param request - EditSecurityConfigMemberRequest
//
// @param headers - EditSecurityConfigMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditSecurityConfigMemberResponse
func (client *Client) EditSecurityConfigMemberWithOptions(request *EditSecurityConfigMemberRequest, headers *EditSecurityConfigMemberHeaders, runtime *util.RuntimeOptions) (_result *EditSecurityConfigMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigKey)) {
		body["configKey"] = request.ConfigKey
	}

	if !tea.BoolValue(util.IsUnset(request.OperateType)) {
		body["operateType"] = request.OperateType
	}

	if !tea.BoolValue(util.IsUnset(request.OperateUserId)) {
		body["operateUserId"] = request.OperateUserId
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
		Action:      tea.String("EditSecurityConfigMember"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/securities/configs/members"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &EditSecurityConfigMemberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑安全卡片管控成员
//
// @param request - EditSecurityConfigMemberRequest
//
// @return EditSecurityConfigMemberResponse
func (client *Client) EditSecurityConfigMember(request *EditSecurityConfigMemberRequest) (_result *EditSecurityConfigMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EditSecurityConfigMemberHeaders{}
	_result = &EditSecurityConfigMemberResponse{}
	_body, _err := client.EditSecurityConfigMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更换组织主管理员
//
// @param request - ExchangeMainAdminRequest
//
// @param headers - ExchangeMainAdminHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExchangeMainAdminResponse
func (client *Client) ExchangeMainAdminWithOptions(request *ExchangeMainAdminRequest, headers *ExchangeMainAdminHeaders, runtime *util.RuntimeOptions) (_result *ExchangeMainAdminResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewAdminUserId)) {
		body["newAdminUserId"] = request.NewAdminUserId
	}

	if !tea.BoolValue(util.IsUnset(request.OldAdminUserId)) {
		body["oldAdminUserId"] = request.OldAdminUserId
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
		Action:      tea.String("ExchangeMainAdmin"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/orgnizations/mainAdministrators/exchange"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ExchangeMainAdminResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更换组织主管理员
//
// @param request - ExchangeMainAdminRequest
//
// @return ExchangeMainAdminResponse
func (client *Client) ExchangeMainAdmin(request *ExchangeMainAdminRequest) (_result *ExchangeMainAdminResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ExchangeMainAdminHeaders{}
	_result = &ExchangeMainAdminResponse{}
	_body, _err := client.ExchangeMainAdminWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分发工作台模版
//
// @param request - ExclusiveCreateDingPortalRequest
//
// @param headers - ExclusiveCreateDingPortalHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExclusiveCreateDingPortalResponse
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

// Summary:
//
// 分发工作台模版
//
// @param request - ExclusiveCreateDingPortalRequest
//
// @return ExclusiveCreateDingPortalResponse
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

// Summary:
//
// 专属文件第一次设置，激活配置
//
// @param request - FileStorageActiveStorageRequest
//
// @param headers - FileStorageActiveStorageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FileStorageActiveStorageResponse
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

// Summary:
//
// 专属文件第一次设置，激活配置
//
// @param request - FileStorageActiveStorageRequest
//
// @return FileStorageActiveStorageResponse
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

// Summary:
//
// 检查专属存储OSS连接
//
// @param request - FileStorageCheckConnectionRequest
//
// @param headers - FileStorageCheckConnectionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FileStorageCheckConnectionResponse
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

// Summary:
//
// 检查专属存储OSS连接
//
// @param request - FileStorageCheckConnectionRequest
//
// @return FileStorageCheckConnectionResponse
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

// Summary:
//
// 专属文件存储获取存储情况(按时间区间)
//
// @param request - FileStorageGetQuotaDataRequest
//
// @param headers - FileStorageGetQuotaDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FileStorageGetQuotaDataResponse
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

// Summary:
//
// 专属文件存储获取存储情况(按时间区间)
//
// @param request - FileStorageGetQuotaDataRequest
//
// @return FileStorageGetQuotaDataResponse
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

// Summary:
//
// 专属文件存储获取存储情况和配置
//
// @param request - FileStorageGetStorageStateRequest
//
// @param headers - FileStorageGetStorageStateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FileStorageGetStorageStateResponse
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

// Summary:
//
// 专属文件存储获取存储情况和配置
//
// @param request - FileStorageGetStorageStateRequest
//
// @return FileStorageGetStorageStateResponse
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

// Summary:
//
// 更新文件专属存储配置
//
// @param request - FileStorageUpdateStorageRequest
//
// @param headers - FileStorageUpdateStorageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FileStorageUpdateStorageResponse
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

// Summary:
//
// 更新文件专属存储配置
//
// @param request - FileStorageUpdateStorageRequest
//
// @return FileStorageUpdateStorageResponse
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

// Summary:
//
// 生成暗水印
//
// @param request - GenerateDarkWaterMarkRequest
//
// @param headers - GenerateDarkWaterMarkHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateDarkWaterMarkResponse
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

// Summary:
//
// 生成暗水印
//
// @param request - GenerateDarkWaterMarkRequest
//
// @return GenerateDarkWaterMarkResponse
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

// Summary:
//
// 获取专属钉钉账号数据迁移结果
//
// @param request - GetAccountTransferListRequest
//
// @param headers - GetAccountTransferListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountTransferListResponse
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

// Summary:
//
// 获取专属钉钉账号数据迁移结果
//
// @param request - GetAccountTransferListRequest
//
// @return GetAccountTransferListResponse
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

// Summary:
//
// 获得组织维度的用户dau
//
// @param headers - GetActiveUserSummaryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetActiveUserSummaryResponse
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

// Summary:
//
// 获得组织维度的用户dau
//
// @return GetActiveUserSummaryResponse
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

// Summary:
//
// 根据AppId获取微应用在该组织下的agentId
//
// @param request - GetAgentIdByRelatedAppIdRequest
//
// @param headers - GetAgentIdByRelatedAppIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentIdByRelatedAppIdResponse
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

// Summary:
//
// 根据AppId获取微应用在该组织下的agentId
//
// @param request - GetAgentIdByRelatedAppIdRequest
//
// @return GetAgentIdByRelatedAppIdResponse
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

// Summary:
//
// 伙伴钉可打标签部门查询
//
// @param headers - GetAllLabelableDeptsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAllLabelableDeptsResponse
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

// Summary:
//
// 伙伴钉可打标签部门查询
//
// @return GetAllLabelableDeptsResponse
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

// Summary:
//
// 获得app分发信息
//
// @param request - GetAppDispatchInfoRequest
//
// @param headers - GetAppDispatchInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppDispatchInfoResponse
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

// Summary:
//
// 获得app分发信息
//
// @param request - GetAppDispatchInfoRequest
//
// @return GetAppDispatchInfoResponse
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

// Summary:
//
// 获得组织维度日程相关信息
//
// @param headers - GetCalenderSummaryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCalenderSummaryResponse
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

// Summary:
//
// 获得组织维度日程相关信息
//
// @return GetCalenderSummaryResponse
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

// Summary:
//
// 根据机器人code获取群openConversationId列表
//
// @param request - GetCidsByBotCodeRequest
//
// @param headers - GetCidsByBotCodeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCidsByBotCodeResponse
func (client *Client) GetCidsByBotCodeWithOptions(request *GetCidsByBotCodeRequest, headers *GetCidsByBotCodeHeaders, runtime *util.RuntimeOptions) (_result *GetCidsByBotCodeResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		query["robotCode"] = request.RobotCode
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
		Action:      tea.String("GetCidsByBotCode"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/groups/openConversationIds"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCidsByBotCodeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据机器人code获取群openConversationId列表
//
// @param request - GetCidsByBotCodeRequest
//
// @return GetCidsByBotCodeResponse
func (client *Client) GetCidsByBotCode(request *GetCidsByBotCodeRequest) (_result *GetCidsByBotCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCidsByBotCodeHeaders{}
	_result = &GetCidsByBotCodeResponse{}
	_body, _err := client.GetCidsByBotCodeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取密级标签
//
// @param request - GetClassTagRequest
//
// @param headers - GetClassTagHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClassTagResponse
func (client *Client) GetClassTagWithOptions(request *GetClassTagRequest, headers *GetClassTagHeaders, runtime *util.RuntimeOptions) (_result *GetClassTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		query["entityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.TagCode)) {
		query["tagCode"] = request.TagCode
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
		Action:      tea.String("GetClassTag"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/classes/entities/tags"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetClassTagResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取密级标签
//
// @param request - GetClassTagRequest
//
// @return GetClassTagResponse
func (client *Client) GetClassTag(request *GetClassTagRequest) (_result *GetClassTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetClassTagHeaders{}
	_result = &GetClassTagResponse{}
	_body, _err := client.GetClassTagWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取发布号的评论列表
//
// @param request - GetCommentListRequest
//
// @param headers - GetCommentListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCommentListResponse
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

// Summary:
//
// 获取发布号的评论列表
//
// @param request - GetCommentListRequest
//
// @return GetCommentListResponse
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

// Summary:
//
// 根据逻辑会议id获取会议基本信息
//
// @param request - GetConfBaseInfoByLogicalIdRequest
//
// @param headers - GetConfBaseInfoByLogicalIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConfBaseInfoByLogicalIdResponse
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

// Summary:
//
// 根据逻辑会议id获取会议基本信息
//
// @param request - GetConfBaseInfoByLogicalIdRequest
//
// @return GetConfBaseInfoByLogicalIdResponse
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

// Summary:
//
// 获取视频会议明细
//
// @param headers - GetConferenceDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConferenceDetailResponse
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

// Summary:
//
// 获取视频会议明细
//
// @return GetConferenceDetailResponse
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

// Summary:
//
// 获取会话分组数据
//
// @param headers - GetConversationCategoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConversationCategoryResponse
func (client *Client) GetConversationCategoryWithOptions(headers *GetConversationCategoryHeaders, runtime *util.RuntimeOptions) (_result *GetConversationCategoryResponse, _err error) {
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
		Action:      tea.String("GetConversationCategory"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/conversationCategories"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConversationCategoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取会话分组数据
//
// @return GetConversationCategoryResponse
func (client *Client) GetConversationCategory() (_result *GetConversationCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetConversationCategoryHeaders{}
	_result = &GetConversationCategoryResponse{}
	_body, _err := client.GetConversationCategoryWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取会话分组详情
//
// @param request - GetConversationDetailRequest
//
// @param headers - GetConversationDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConversationDetailResponse
func (client *Client) GetConversationDetailWithOptions(request *GetConversationDetailRequest, headers *GetConversationDetailHeaders, runtime *util.RuntimeOptions) (_result *GetConversationDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
		Action:      tea.String("GetConversationDetail"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/categories/conversations/details/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConversationDetailResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取会话分组详情
//
// @param request - GetConversationDetailRequest
//
// @return GetConversationDetailResponse
func (client *Client) GetConversationDetail(request *GetConversationDetailRequest) (_result *GetConversationDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetConversationDetailHeaders{}
	_result = &GetConversationDetailResponse{}
	_body, _err := client.GetConversationDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取部门维度发布日志信息
//
// @param request - GetDingReportDeptSummaryRequest
//
// @param headers - GetDingReportDeptSummaryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDingReportDeptSummaryResponse
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

// Summary:
//
// 获取部门维度发布日志信息
//
// @param request - GetDingReportDeptSummaryRequest
//
// @return GetDingReportDeptSummaryResponse
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

// Summary:
//
// 获取组织维度发布日志信息
//
// @param headers - GetDingReportSummaryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDingReportSummaryResponse
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

// Summary:
//
// 获取组织维度发布日志信息
//
// @return GetDingReportSummaryResponse
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

// Summary:
//
// 获得部门维度用户创建文档数和创建文档人数
//
// @param request - GetDocCreatedDeptSummaryRequest
//
// @param headers - GetDocCreatedDeptSummaryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocCreatedDeptSummaryResponse
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

// Summary:
//
// 获得部门维度用户创建文档数和创建文档人数
//
// @param request - GetDocCreatedDeptSummaryRequest
//
// @return GetDocCreatedDeptSummaryResponse
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

// Summary:
//
// 获取组织维度用户创建文档数和创建文档人数
//
// @param headers - GetDocCreatedSummaryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocCreatedSummaryResponse
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

// Summary:
//
// 获取组织维度用户创建文档数和创建文档人数
//
// @return GetDocCreatedSummaryResponse
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

// Summary:
//
// 获取专属账号所有组织列表
//
// @param request - GetExclusiveAccountAllOrgListRequest
//
// @param headers - GetExclusiveAccountAllOrgListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExclusiveAccountAllOrgListResponse
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

// Summary:
//
// 获取专属账号所有组织列表
//
// @param request - GetExclusiveAccountAllOrgListRequest
//
// @return GetExclusiveAccountAllOrgListResponse
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

// Summary:
//
// 获取部门维度发布智能填表数量和使用智能填表人数
//
// @param request - GetGeneralFormCreatedDeptSummaryRequest
//
// @param headers - GetGeneralFormCreatedDeptSummaryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGeneralFormCreatedDeptSummaryResponse
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

// Summary:
//
// 获取部门维度发布智能填表数量和使用智能填表人数
//
// @param request - GetGeneralFormCreatedDeptSummaryRequest
//
// @return GetGeneralFormCreatedDeptSummaryResponse
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

// Summary:
//
// 获取组织维度发布智能填表数量和使用智能填表人数
//
// @param headers - GetGeneralFormCreatedSummaryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGeneralFormCreatedSummaryResponse
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

// Summary:
//
// 获取组织维度发布智能填表数量和使用智能填表人数
//
// @return GetGeneralFormCreatedSummaryResponse
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

// Summary:
//
// 获取群活跃明细
//
// @param request - GetGroupActiveInfoRequest
//
// @param headers - GetGroupActiveInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupActiveInfoResponse
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

// Summary:
//
// 获取群活跃明细
//
// @param request - GetGroupActiveInfoRequest
//
// @return GetGroupActiveInfoResponse
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

// Summary:
//
// 根据群会话id获取群相关信息
//
// @param request - GetGroupInfoByCidRequest
//
// @param headers - GetGroupInfoByCidHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupInfoByCidResponse
func (client *Client) GetGroupInfoByCidWithOptions(request *GetGroupInfoByCidRequest, headers *GetGroupInfoByCidHeaders, runtime *util.RuntimeOptions) (_result *GetGroupInfoByCidResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		query["openConversationId"] = request.OpenConversationId
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
		Action:      tea.String("GetGroupInfoByCid"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/groups/infos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGroupInfoByCidResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据群会话id获取群相关信息
//
// @param request - GetGroupInfoByCidRequest
//
// @return GetGroupInfoByCidResponse
func (client *Client) GetGroupInfoByCid(request *GetGroupInfoByCidRequest) (_result *GetGroupInfoByCidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetGroupInfoByCidHeaders{}
	_result = &GetGroupInfoByCidResponse{}
	_body, _err := client.GetGroupInfoByCidWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据群会话id获取组织cropid
//
// @param request - GetGroupOrgByCidRequest
//
// @param headers - GetGroupOrgByCidHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupOrgByCidResponse
func (client *Client) GetGroupOrgByCidWithOptions(request *GetGroupOrgByCidRequest, headers *GetGroupOrgByCidHeaders, runtime *util.RuntimeOptions) (_result *GetGroupOrgByCidResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		query["openConversationId"] = request.OpenConversationId
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
		Action:      tea.String("GetGroupOrgByCid"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/groups/organizations/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGroupOrgByCidResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据群会话id获取组织cropid
//
// @param request - GetGroupOrgByCidRequest
//
// @return GetGroupOrgByCidResponse
func (client *Client) GetGroupOrgByCid(request *GetGroupOrgByCidRequest) (_result *GetGroupOrgByCidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetGroupOrgByCidHeaders{}
	_result = &GetGroupOrgByCidResponse{}
	_body, _err := client.GetGroupOrgByCidWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取未活跃用户登陆统计明细
//
// @param request - GetInActiveUserListRequest
//
// @param headers - GetInActiveUserListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInActiveUserListResponse
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

// Summary:
//
// 获取未活跃用户登陆统计明细
//
// @param request - GetInActiveUserListRequest
//
// @return GetInActiveUserListResponse
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

// Summary:
//
// 获取最后一次验证通过的企业认证信息
//
// @param request - GetLastOrgAuthDataRequest
//
// @param headers - GetLastOrgAuthDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLastOrgAuthDataResponse
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

// Summary:
//
// 获取最后一次验证通过的企业认证信息
//
// @param request - GetLastOrgAuthDataRequest
//
// @return GetLastOrgAuthDataResponse
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

// Summary:
//
// 消息规则配置和群属性接口
//
// @param request - GetMsgConfigRequest
//
// @param headers - GetMsgConfigHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMsgConfigResponse
func (client *Client) GetMsgConfigWithOptions(request *GetMsgConfigRequest, headers *GetMsgConfigHeaders, runtime *util.RuntimeOptions) (_result *GetMsgConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupTopic)) {
		body["groupTopic"] = request.GroupTopic
	}

	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		body["groupType"] = request.GroupType
	}

	if !tea.BoolValue(util.IsUnset(request.ListDynamicAttr)) {
		body["listDynamicAttr"] = request.ListDynamicAttr
	}

	if !tea.BoolValue(util.IsUnset(request.ListEmployeeCode)) {
		body["listEmployeeCode"] = request.ListEmployeeCode
	}

	if !tea.BoolValue(util.IsUnset(request.ListUnitId)) {
		body["listUnitId"] = request.ListUnitId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerJobNo)) {
		body["ownerJobNo"] = request.OwnerJobNo
	}

	if !tea.BoolValue(util.IsUnset(request.RuleBusinessCode)) {
		body["ruleBusinessCode"] = request.RuleBusinessCode
	}

	if !tea.BoolValue(util.IsUnset(request.RuleCategory)) {
		body["ruleCategory"] = request.RuleCategory
	}

	if !tea.BoolValue(util.IsUnset(request.RuleCode)) {
		body["ruleCode"] = request.RuleCode
	}

	if !tea.BoolValue(util.IsUnset(request.SecretKey)) {
		body["secretKey"] = request.SecretKey
	}

	if !tea.BoolValue(util.IsUnset(request.SysCode)) {
		body["sysCode"] = request.SysCode
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
		Action:      tea.String("GetMsgConfig"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/portals/msgConfigs/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMsgConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 消息规则配置和群属性接口
//
// @param request - GetMsgConfigRequest
//
// @return GetMsgConfigResponse
func (client *Client) GetMsgConfig(request *GetMsgConfigRequest) (_result *GetMsgConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetMsgConfigHeaders{}
	_result = &GetMsgConfigResponse{}
	_body, _err := client.GetMsgConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取消息定位链接
//
// @param request - GetMsgLocationRequest
//
// @param headers - GetMsgLocationHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMsgLocationResponse
func (client *Client) GetMsgLocationWithOptions(request *GetMsgLocationRequest, headers *GetMsgLocationHeaders, runtime *util.RuntimeOptions) (_result *GetMsgLocationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenMsgId)) {
		body["openMsgId"] = request.OpenMsgId
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
		Action:      tea.String("GetMsgLocation"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/messageLocations/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMsgLocationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取消息定位链接
//
// @param request - GetMsgLocationRequest
//
// @return GetMsgLocationResponse
func (client *Client) GetMsgLocation(request *GetMsgLocationRequest) (_result *GetMsgLocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetMsgLocationHeaders{}
	_result = &GetMsgLocationResponse{}
	_body, _err := client.GetMsgLocationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取oa后台操作日志记录
//
// @param request - GetOaOperatorLogListRequest
//
// @param headers - GetOaOperatorLogListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOaOperatorLogListResponse
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

// Summary:
//
// 获取oa后台操作日志记录
//
// @param request - GetOaOperatorLogListRequest
//
// @return GetOaOperatorLogListResponse
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

// Summary:
//
// 获取企业的外部审计群列表
//
// @param request - GetOutGroupsByPageRequest
//
// @param headers - GetOutGroupsByPageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOutGroupsByPageResponse
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

// Summary:
//
// 获取企业的外部审计群列表
//
// @param request - GetOutGroupsByPageRequest
//
// @return GetOutGroupsByPageResponse
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

// Summary:
//
// 获取外部审计群消息记录
//
// @param request - GetOutsideAuditGroupMessageByPageRequest
//
// @param headers - GetOutsideAuditGroupMessageByPageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOutsideAuditGroupMessageByPageResponse
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

// Summary:
//
// 获取外部审计群消息记录
//
// @param request - GetOutsideAuditGroupMessageByPageRequest
//
// @return GetOutsideAuditGroupMessageByPageResponse
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

// Summary:
//
// 伙伴钉根据父标签查询子标签
//
// @param headers - GetPartnerTypeByParentIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPartnerTypeByParentIdResponse
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

// Summary:
//
// 伙伴钉根据父标签查询子标签
//
// @return GetPartnerTypeByParentIdResponse
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

// Summary:
//
// 获取专属存储容量信息
//
// @param request - GetPrivateStoreCapacityUsageRequest
//
// @param headers - GetPrivateStoreCapacityUsageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrivateStoreCapacityUsageResponse
func (client *Client) GetPrivateStoreCapacityUsageWithOptions(request *GetPrivateStoreCapacityUsageRequest, headers *GetPrivateStoreCapacityUsageHeaders, runtime *util.RuntimeOptions) (_result *GetPrivateStoreCapacityUsageResponse, _err error) {
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
		Action:      tea.String("GetPrivateStoreCapacityUsage"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/privateStores/capacityUsages/infos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPrivateStoreCapacityUsageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取专属存储容量信息
//
// @param request - GetPrivateStoreCapacityUsageRequest
//
// @return GetPrivateStoreCapacityUsageResponse
func (client *Client) GetPrivateStoreCapacityUsage(request *GetPrivateStoreCapacityUsageRequest) (_result *GetPrivateStoreCapacityUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPrivateStoreCapacityUsageHeaders{}
	_result = &GetPrivateStoreCapacityUsageResponse{}
	_body, _err := client.GetPrivateStoreCapacityUsageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页获取专属存储文件信息
//
// @param request - GetPrivateStoreFileInfosByPageRequest
//
// @param headers - GetPrivateStoreFileInfosByPageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrivateStoreFileInfosByPageResponse
func (client *Client) GetPrivateStoreFileInfosByPageWithOptions(request *GetPrivateStoreFileInfosByPageRequest, headers *GetPrivateStoreFileInfosByPageHeaders, runtime *util.RuntimeOptions) (_result *GetPrivateStoreFileInfosByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileCreateTime)) {
		body["fileCreateTime"] = request.FileCreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.FileStatus)) {
		body["fileStatus"] = request.FileStatus
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["order"] = request.Order
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
		Action:      tea.String("GetPrivateStoreFileInfosByPage"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/privateStores/fileInfos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPrivateStoreFileInfosByPageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页获取专属存储文件信息
//
// @param request - GetPrivateStoreFileInfosByPageRequest
//
// @return GetPrivateStoreFileInfosByPageResponse
func (client *Client) GetPrivateStoreFileInfosByPage(request *GetPrivateStoreFileInfosByPageRequest) (_result *GetPrivateStoreFileInfosByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPrivateStoreFileInfosByPageHeaders{}
	_result = &GetPrivateStoreFileInfosByPageResponse{}
	_body, _err := client.GetPrivateStoreFileInfosByPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取专属存储文件路径
//
// @param request - GetPrivateStoreFilePathRequest
//
// @param headers - GetPrivateStoreFilePathHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrivateStoreFilePathResponse
func (client *Client) GetPrivateStoreFilePathWithOptions(request *GetPrivateStoreFilePathRequest, headers *GetPrivateStoreFilePathHeaders, runtime *util.RuntimeOptions) (_result *GetPrivateStoreFilePathResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DentryId)) {
		body["dentryId"] = request.DentryId
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
		Action:      tea.String("GetPrivateStoreFilePath"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/privateStores/filePaths/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPrivateStoreFilePathResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取专属存储文件路径
//
// @param request - GetPrivateStoreFilePathRequest
//
// @return GetPrivateStoreFilePathResponse
func (client *Client) GetPrivateStoreFilePath(request *GetPrivateStoreFilePathRequest) (_result *GetPrivateStoreFilePathResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPrivateStoreFilePathHeaders{}
	_result = &GetPrivateStoreFilePathResponse{}
	_body, _err := client.GetPrivateStoreFilePathWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取公共设备列表。
//
// @param request - GetPublicDevicesRequest
//
// @param headers - GetPublicDevicesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPublicDevicesResponse
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

	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		query["serialNumber"] = request.SerialNumber
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

// Summary:
//
// 获取公共设备列表。
//
// @param request - GetPublicDevicesRequest
//
// @return GetPublicDevicesResponse
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

// Summary:
//
// 获取互动服务窗相关数据
//
// @param request - GetPublisherSummaryRequest
//
// @param headers - GetPublisherSummaryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPublisherSummaryResponse
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

// Summary:
//
// 获取互动服务窗相关数据
//
// @param request - GetPublisherSummaryRequest
//
// @return GetPublisherSummaryResponse
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

// Summary:
//
// 获取实人认证接口调用记录
//
// @param request - GetRealPeopleRecordsRequest
//
// @param headers - GetRealPeopleRecordsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRealPeopleRecordsResponse
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

// Summary:
//
// 获取实人认证接口调用记录
//
// @param request - GetRealPeopleRecordsRequest
//
// @return GetRealPeopleRecordsResponse
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

// Summary:
//
// 获取人脸对比接口调用记录
//
// @param request - GetRecognizeRecordsRequest
//
// @param headers - GetRecognizeRecordsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRecognizeRecordsResponse
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

// Summary:
//
// 获取人脸对比接口调用记录
//
// @param request - GetRecognizeRecordsRequest
//
// @return GetRecognizeRecordsResponse
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

// Summary:
//
// 根据机器人标识查询机器人信息
//
// @param request - GetRobotInfoByCodeRequest
//
// @param headers - GetRobotInfoByCodeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRobotInfoByCodeResponse
func (client *Client) GetRobotInfoByCodeWithOptions(request *GetRobotInfoByCodeRequest, headers *GetRobotInfoByCodeHeaders, runtime *util.RuntimeOptions) (_result *GetRobotInfoByCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		query["robotCode"] = request.RobotCode
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
		Action:      tea.String("GetRobotInfoByCode"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/robots/infos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRobotInfoByCodeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据机器人标识查询机器人信息
//
// @param request - GetRobotInfoByCodeRequest
//
// @return GetRobotInfoByCodeResponse
func (client *Client) GetRobotInfoByCode(request *GetRobotInfoByCodeRequest) (_result *GetRobotInfoByCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetRobotInfoByCodeHeaders{}
	_result = &GetRobotInfoByCodeResponse{}
	_body, _err := client.GetRobotInfoByCodeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取安全管控卡片成员
//
// @param request - GetSecurityConfigMemberRequest
//
// @param headers - GetSecurityConfigMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecurityConfigMemberResponse
func (client *Client) GetSecurityConfigMemberWithOptions(request *GetSecurityConfigMemberRequest, headers *GetSecurityConfigMemberHeaders, runtime *util.RuntimeOptions) (_result *GetSecurityConfigMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigKey)) {
		body["configKey"] = request.ConfigKey
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
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
		Action:      tea.String("GetSecurityConfigMember"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/securities/configs/members/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSecurityConfigMemberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取安全管控卡片成员
//
// @param request - GetSecurityConfigMemberRequest
//
// @return GetSecurityConfigMemberResponse
func (client *Client) GetSecurityConfigMember(request *GetSecurityConfigMemberRequest) (_result *GetSecurityConfigMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSecurityConfigMemberHeaders{}
	_result = &GetSecurityConfigMemberResponse{}
	_body, _err := client.GetSecurityConfigMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取审计协议签署人员信息
//
// @param request - GetSignedDetailByPageRequest
//
// @param headers - GetSignedDetailByPageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSignedDetailByPageResponse
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

// Summary:
//
// 获取审计协议签署人员信息
//
// @param request - GetSignedDetailByPageRequest
//
// @return GetSignedDetailByPageResponse
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

// Summary:
//
// 获取多个可信设备信息，包括mac地址、staffId、platform
//
// @param request - GetTrustDeviceListRequest
//
// @param headers - GetTrustDeviceListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrustDeviceListResponse
func (client *Client) GetTrustDeviceListWithOptions(request *GetTrustDeviceListRequest, headers *GetTrustDeviceListHeaders, runtime *util.RuntimeOptions) (_result *GetTrustDeviceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GmtCreateEnd)) {
		body["gmtCreateEnd"] = request.GmtCreateEnd
	}

	if !tea.BoolValue(util.IsUnset(request.GmtCreateStart)) {
		body["gmtCreateStart"] = request.GmtCreateStart
	}

	if !tea.BoolValue(util.IsUnset(request.GmtModifiedEnd)) {
		body["gmtModifiedEnd"] = request.GmtModifiedEnd
	}

	if !tea.BoolValue(util.IsUnset(request.GmtModifiedStart)) {
		body["gmtModifiedStart"] = request.GmtModifiedStart
	}

	if !tea.BoolValue(util.IsUnset(request.MacAddress)) {
		body["macAddress"] = request.MacAddress
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

	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		body["serialNumber"] = request.SerialNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
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

// Summary:
//
// 获取多个可信设备信息，包括mac地址、staffId、platform
//
// @param request - GetTrustDeviceListRequest
//
// @return GetTrustDeviceListResponse
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

// Summary:
//
// 获得组织维度用户版本分布情况
//
// @param request - GetUserAppVersionSummaryRequest
//
// @param headers - GetUserAppVersionSummaryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserAppVersionSummaryResponse
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

// Summary:
//
// 获得组织维度用户版本分布情况
//
// @param request - GetUserAppVersionSummaryRequest
//
// @return GetUserAppVersionSummaryResponse
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

// Summary:
//
// 人脸录入状态查询
//
// @param request - GetUserFaceStateRequest
//
// @param headers - GetUserFaceStateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserFaceStateResponse
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

// Summary:
//
// 人脸录入状态查询
//
// @param request - GetUserFaceStateRequest
//
// @return GetUserFaceStateResponse
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

// Summary:
//
// 实人认证状态查询
//
// @param request - GetUserRealPeopleStateRequest
//
// @param headers - GetUserRealPeopleStateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserRealPeopleStateResponse
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

// Summary:
//
// 实人认证状态查询
//
// @param request - GetUserRealPeopleStateRequest
//
// @return GetUserRealPeopleStateResponse
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

// Summary:
//
// 获取用户停留时长
//
// @param request - GetUserStayLengthRequest
//
// @param headers - GetUserStayLengthHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserStayLengthResponse
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

// Summary:
//
// 获取用户停留时长
//
// @param request - GetUserStayLengthRequest
//
// @return GetUserStayLengthResponse
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

// Summary:
//
// 获取文件病毒扫描结果
//
// @param request - GetVirusScanResultRequest
//
// @param headers - GetVirusScanResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVirusScanResultResponse
func (client *Client) GetVirusScanResultWithOptions(request *GetVirusScanResultRequest, headers *GetVirusScanResultHeaders, runtime *util.RuntimeOptions) (_result *GetVirusScanResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
		Action:      tea.String("GetVirusScanResult"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/virusScanTasks/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVirusScanResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件病毒扫描结果
//
// @param request - GetVirusScanResultRequest
//
// @return GetVirusScanResultResponse
func (client *Client) GetVirusScanResult(request *GetVirusScanResultRequest) (_result *GetVirusScanResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetVirusScanResultHeaders{}
	_result = &GetVirusScanResultResponse{}
	_body, _err := client.GetVirusScanResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据群属性查询群ID
//
// @param request - GroupQueryByAttrRequest
//
// @param headers - GroupQueryByAttrHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GroupQueryByAttrResponse
func (client *Client) GroupQueryByAttrWithOptions(request *GroupQueryByAttrRequest, headers *GroupQueryByAttrHeaders, runtime *util.RuntimeOptions) (_result *GroupQueryByAttrResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupTopic)) {
		body["groupTopic"] = request.GroupTopic
	}

	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		body["groupType"] = request.GroupType
	}

	if !tea.BoolValue(util.IsUnset(request.ListDynamicAttr)) {
		body["listDynamicAttr"] = request.ListDynamicAttr
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		body["pageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecretKey)) {
		body["secretKey"] = request.SecretKey
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
		Action:      tea.String("GroupQueryByAttr"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/portals/groups/queryGroup"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GroupQueryByAttrResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据群属性查询群ID
//
// @param request - GroupQueryByAttrRequest
//
// @return GroupQueryByAttrResponse
func (client *Client) GroupQueryByAttr(request *GroupQueryByAttrRequest) (_result *GroupQueryByAttrResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GroupQueryByAttrHeaders{}
	_result = &GroupQueryByAttrResponse{}
	_body, _err := client.GroupQueryByAttrWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据群ID查询群属性
//
// @param request - GroupQueryByOpenIdRequest
//
// @param headers - GroupQueryByOpenIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GroupQueryByOpenIdResponse
func (client *Client) GroupQueryByOpenIdWithOptions(request *GroupQueryByOpenIdRequest, headers *GroupQueryByOpenIdHeaders, runtime *util.RuntimeOptions) (_result *GroupQueryByOpenIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretKey)) {
		body["secretKey"] = request.SecretKey
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
		Action:      tea.String("GroupQueryByOpenId"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/portals/groups/getGroupByOpenConversationId"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GroupQueryByOpenIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据群ID查询群属性
//
// @param request - GroupQueryByOpenIdRequest
//
// @return GroupQueryByOpenIdResponse
func (client *Client) GroupQueryByOpenId(request *GroupQueryByOpenIdRequest) (_result *GroupQueryByOpenIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GroupQueryByOpenIdHeaders{}
	_result = &GroupQueryByOpenIdResponse{}
	_body, _err := client.GroupQueryByOpenIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业文件审计日志
//
// @param request - ListAuditLogRequest
//
// @param headers - ListAuditLogHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuditLogResponse
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

// Summary:
//
// 获取企业文件审计日志
//
// @param request - ListAuditLogRequest
//
// @return ListAuditLogResponse
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

// Summary:
//
// 根据机器人code列表查询机器人信息
//
// @param request - ListByCodesRequest
//
// @param headers - ListByCodesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListByCodesResponse
func (client *Client) ListByCodesWithOptions(request *ListByCodesRequest, headers *ListByCodesHeaders, runtime *util.RuntimeOptions) (_result *ListByCodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
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
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("ListByCodes"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/sceneGroups/robotInfos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListByCodesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据机器人code列表查询机器人信息
//
// @param request - ListByCodesRequest
//
// @return ListByCodesResponse
func (client *Client) ListByCodes(request *ListByCodesRequest) (_result *ListByCodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListByCodesHeaders{}
	_result = &ListByCodesResponse{}
	_body, _err := client.ListByCodesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据插件id列表查询插件信息
//
// @param request - ListByPluginIdsRequest
//
// @param headers - ListByPluginIdsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListByPluginIdsResponse
func (client *Client) ListByPluginIdsWithOptions(request *ListByPluginIdsRequest, headers *ListByPluginIdsHeaders, runtime *util.RuntimeOptions) (_result *ListByPluginIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
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
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("ListByPluginIds"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/sceneGroups/pluginInfos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListByPluginIdsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据插件id列表查询插件信息
//
// @param request - ListByPluginIdsRequest
//
// @return ListByPluginIdsResponse
func (client *Client) ListByPluginIds(request *ListByPluginIdsRequest) (_result *ListByPluginIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListByPluginIdsHeaders{}
	_result = &ListByPluginIdsResponse{}
	_body, _err := client.ListByPluginIdsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询分组列表
//
// @param tmpReq - ListCategorysRequest
//
// @param headers - ListCategorysHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCategorysResponse
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

// Summary:
//
// 查询分组列表
//
// @param request - ListCategorysRequest
//
// @return ListCategorysResponse
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

// Summary:
//
// 通过手机号获取已加入的属于渠道组织的列表信息
//
// @param request - ListJoinOrgInfoRequest
//
// @param headers - ListJoinOrgInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJoinOrgInfoResponse
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

// Summary:
//
// 通过手机号获取已加入的属于渠道组织的列表信息
//
// @param request - ListJoinOrgInfoRequest
//
// @return ListJoinOrgInfoResponse
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

// Summary:
//
// 获取小程序版本列表
//
// @param request - ListMiniAppAvailableVersionRequest
//
// @param headers - ListMiniAppAvailableVersionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMiniAppAvailableVersionResponse
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

// Summary:
//
// 获取小程序版本列表
//
// @param request - ListMiniAppAvailableVersionRequest
//
// @return ListMiniAppAvailableVersionResponse
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

// Summary:
//
// 获取小程序历史版本列表
//
// @param request - ListMiniAppHistoryVersionRequest
//
// @param headers - ListMiniAppHistoryVersionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMiniAppHistoryVersionResponse
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

// Summary:
//
// 获取小程序历史版本列表
//
// @param request - ListMiniAppHistoryVersionRequest
//
// @return ListMiniAppHistoryVersionResponse
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

// Summary:
//
// 查询伙伴角色
//
// @param headers - ListPartnerRolesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPartnerRolesResponse
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

// Summary:
//
// 查询伙伴角色
//
// @return ListPartnerRolesResponse
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

// Summary:
//
// 获取巡点信息列表
//
// @param request - ListPunchScheduleByConditionWithPagingRequest
//
// @param headers - ListPunchScheduleByConditionWithPagingHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPunchScheduleByConditionWithPagingResponse
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

// Summary:
//
// 获取巡点信息列表
//
// @param request - ListPunchScheduleByConditionWithPagingRequest
//
// @return ListPunchScheduleByConditionWithPagingResponse
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

// Summary:
//
// 查询规则列表
//
// @param tmpReq - ListRulesRequest
//
// @param headers - ListRulesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRulesResponse
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

// Summary:
//
// 查询规则列表
//
// @param request - ListRulesRequest
//
// @return ListRulesResponse
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

// Summary:
//
// 指定用户强制下线
//
// @param request - LogoutRequest
//
// @param headers - LogoutHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LogoutResponse
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

// Summary:
//
// 指定用户强制下线
//
// @param request - LogoutRequest
//
// @return LogoutResponse
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

// Summary:
//
// 购买权益包
//
// @param request - OpenBenefitPackageRequest
//
// @param headers - OpenBenefitPackageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenBenefitPackageResponse
func (client *Client) OpenBenefitPackageWithOptions(request *OpenBenefitPackageRequest, headers *OpenBenefitPackageHeaders, runtime *util.RuntimeOptions) (_result *OpenBenefitPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BenefitPackage)) {
		body["benefitPackage"] = request.BenefitPackage
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["endDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["startDate"] = request.StartDate
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
		Action:      tea.String("OpenBenefitPackage"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/benefitPackages/purchase"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenBenefitPackageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 购买权益包
//
// @param request - OpenBenefitPackageRequest
//
// @return OpenBenefitPackageResponse
func (client *Client) OpenBenefitPackage(request *OpenBenefitPackageRequest) (_result *OpenBenefitPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OpenBenefitPackageHeaders{}
	_result = &OpenBenefitPackageResponse{}
	_body, _err := client.OpenBenefitPackageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 商机冲突检测
//
// @param request - OpportunitySearchRequest
//
// @param headers - OpportunitySearchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpportunitySearchResponse
func (client *Client) OpportunitySearchWithOptions(request *OpportunitySearchRequest, headers *OpportunitySearchHeaders, runtime *util.RuntimeOptions) (_result *OpportunitySearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
		Action:      tea.String("OpportunitySearch"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/opportunities/check"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &OpportunitySearchResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 商机冲突检测
//
// @param request - OpportunitySearchRequest
//
// @return OpportunitySearchResponse
func (client *Client) OpportunitySearch(request *OpportunitySearchRequest) (_result *OpportunitySearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OpportunitySearchHeaders{}
	_result = &OpportunitySearchResponse{}
	_body, _err := client.OpportunitySearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 防作弊风险检测
//
// @param request - PreventCheatingCheckRiskRequest
//
// @param headers - PreventCheatingCheckRiskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreventCheatingCheckRiskResponse
func (client *Client) PreventCheatingCheckRiskWithOptions(request *PreventCheatingCheckRiskRequest, headers *PreventCheatingCheckRiskHeaders, runtime *util.RuntimeOptions) (_result *PreventCheatingCheckRiskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientVer)) {
		body["clientVer"] = request.ClientVer
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		body["platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.PlatformVer)) {
		body["platformVer"] = request.PlatformVer
	}

	if !tea.BoolValue(util.IsUnset(request.Sec)) {
		body["sec"] = request.Sec
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
		Action:      tea.String("PreventCheatingCheckRisk"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/preventCheats/risks/check"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PreventCheatingCheckRiskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 防作弊风险检测
//
// @param request - PreventCheatingCheckRiskRequest
//
// @return PreventCheatingCheckRiskResponse
func (client *Client) PreventCheatingCheckRisk(request *PreventCheatingCheckRiskRequest) (_result *PreventCheatingCheckRiskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PreventCheatingCheckRiskHeaders{}
	_result = &PreventCheatingCheckRiskResponse{}
	_body, _err := client.PreventCheatingCheckRiskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发送文件更改的评论
//
// @param request - PublishFileChangeNoticeRequest
//
// @param headers - PublishFileChangeNoticeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishFileChangeNoticeResponse
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

// Summary:
//
// 发送文件更改的评论
//
// @param request - PublishFileChangeNoticeRequest
//
// @return PublishFileChangeNoticeResponse
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

// Summary:
//
// 发布规则
//
// @param request - PublishRuleRequest
//
// @param headers - PublishRuleHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishRuleResponse
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

// Summary:
//
// 发布规则
//
// @param request - PublishRuleRequest
//
// @return PublishRuleResponse
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

// Summary:
//
// 推送专属设计中自建/第三方应用的小红点
//
// @param request - PushBadgeRequest
//
// @param headers - PushBadgeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushBadgeResponse
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

// Summary:
//
// 推送专属设计中自建/第三方应用的小红点
//
// @param request - PushBadgeRequest
//
// @return PushBadgeResponse
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

// Summary:
//
// 查询跨云存储配置
//
// @param request - QueryAcrossCloudStroageConfigsRequest
//
// @param headers - QueryAcrossCloudStroageConfigsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAcrossCloudStroageConfigsResponse
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

// Summary:
//
// 查询跨云存储配置
//
// @param request - QueryAcrossCloudStroageConfigsRequest
//
// @return QueryAcrossCloudStroageConfigsResponse
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

// Summary:
//
// 根据手机号查询渠道组织中的员工信息
//
// @param request - QueryChannelStaffInfoByMobileRequest
//
// @param headers - QueryChannelStaffInfoByMobileHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryChannelStaffInfoByMobileResponse
func (client *Client) QueryChannelStaffInfoByMobileWithOptions(request *QueryChannelStaffInfoByMobileRequest, headers *QueryChannelStaffInfoByMobileHeaders, runtime *util.RuntimeOptions) (_result *QueryChannelStaffInfoByMobileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["mobile"] = request.Mobile
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
		Action:      tea.String("QueryChannelStaffInfoByMobile"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/channelOrganizations/users"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryChannelStaffInfoByMobileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据手机号查询渠道组织中的员工信息
//
// @param request - QueryChannelStaffInfoByMobileRequest
//
// @return QueryChannelStaffInfoByMobileResponse
func (client *Client) QueryChannelStaffInfoByMobile(request *QueryChannelStaffInfoByMobileRequest) (_result *QueryChannelStaffInfoByMobileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryChannelStaffInfoByMobileHeaders{}
	_result = &QueryChannelStaffInfoByMobileResponse{}
	_body, _err := client.QueryChannelStaffInfoByMobileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取分组下会话列表
//
// @param request - QueryConversationPageRequest
//
// @param headers - QueryConversationPageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryConversationPageResponse
func (client *Client) QueryConversationPageWithOptions(request *QueryConversationPageRequest, headers *QueryConversationPageHeaders, runtime *util.RuntimeOptions) (_result *QueryConversationPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		query["categoryId"] = request.CategoryId
	}

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
		Action:      tea.String("QueryConversationPage"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/categories/conversations"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryConversationPageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取分组下会话列表
//
// @param request - QueryConversationPageRequest
//
// @return QueryConversationPageResponse
func (client *Client) QueryConversationPage(request *QueryConversationPageRequest) (_result *QueryConversationPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryConversationPageHeaders{}
	_result = &QueryConversationPageResponse{}
	_body, _err := client.QueryConversationPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询专属版权益
//
// @param headers - QueryExclusiveBenefitsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryExclusiveBenefitsResponse
func (client *Client) QueryExclusiveBenefitsWithOptions(headers *QueryExclusiveBenefitsHeaders, runtime *util.RuntimeOptions) (_result *QueryExclusiveBenefitsResponse, _err error) {
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
		Action:      tea.String("QueryExclusiveBenefits"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/benefits"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryExclusiveBenefitsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询专属版权益
//
// @return QueryExclusiveBenefitsResponse
func (client *Client) QueryExclusiveBenefits() (_result *QueryExclusiveBenefitsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryExclusiveBenefitsHeaders{}
	_result = &QueryExclusiveBenefitsResponse{}
	_body, _err := client.QueryExclusiveBenefitsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 伙伴钉根据uid查询人员的标签信息
//
// @param headers - QueryPartnerInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPartnerInfoResponse
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

// Summary:
//
// 伙伴钉根据uid查询人员的标签信息
//
// @return QueryPartnerInfoResponse
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

// Summary:
//
// 根据templateId查询模版信息
//
// @param headers - QueryTemplateInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTemplateInfoResponse
func (client *Client) QueryTemplateInfoWithOptions(templateId *string, headers *QueryTemplateInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryTemplateInfoResponse, _err error) {
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
		Action:      tea.String("QueryTemplateInfo"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/sceneGroups/templates/" + tea.StringValue(templateId) + "/infos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTemplateInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据templateId查询模版信息
//
// @return QueryTemplateInfoResponse
func (client *Client) QueryTemplateInfo(templateId *string) (_result *QueryTemplateInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryTemplateInfoHeaders{}
	_result = &QueryTemplateInfoResponse{}
	_body, _err := client.QueryTemplateInfoWithOptions(templateId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户截屏操作记录
//
// @param request - QueryUserBehaviorRequest
//
// @param headers - QueryUserBehaviorHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserBehaviorResponse
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

// Summary:
//
// 获取用户截屏操作记录
//
// @param request - QueryUserBehaviorRequest
//
// @return QueryUserBehaviorResponse
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

// Summary:
//
// 小程序版本回滚
//
// @param request - RollbackMiniAppVersionRequest
//
// @param headers - RollbackMiniAppVersionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackMiniAppVersionResponse
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

// Summary:
//
// 小程序版本回滚
//
// @param request - RollbackMiniAppVersionRequest
//
// @return RollbackMiniAppVersionResponse
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

// Summary:
//
// 按规则批量发消息
//
// @param request - RuleBatchReceiverRequest
//
// @param headers - RuleBatchReceiverHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RuleBatchReceiverResponse
func (client *Client) RuleBatchReceiverWithOptions(request *RuleBatchReceiverRequest, headers *RuleBatchReceiverHeaders, runtime *util.RuntimeOptions) (_result *RuleBatchReceiverResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BatchNo)) {
		body["batchNo"] = request.BatchNo
	}

	if !tea.BoolValue(util.IsUnset(request.CardOptions)) {
		body["cardOptions"] = request.CardOptions
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.RuleCode)) {
		body["ruleCode"] = request.RuleCode
	}

	if !tea.BoolValue(util.IsUnset(request.SecretKey)) {
		body["secretKey"] = request.SecretKey
	}

	if !tea.BoolValue(util.IsUnset(request.SpecialStrategy)) {
		body["specialStrategy"] = request.SpecialStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.TaskBatchNo)) {
		body["taskBatchNo"] = request.TaskBatchNo
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
		Action:      tea.String("RuleBatchReceiver"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/dmc/rules/messages/batchSend"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RuleBatchReceiverResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 按规则批量发消息
//
// @param request - RuleBatchReceiverRequest
//
// @return RuleBatchReceiverResponse
func (client *Client) RuleBatchReceiver(request *RuleBatchReceiverRequest) (_result *RuleBatchReceiverResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RuleBatchReceiverHeaders{}
	_result = &RuleBatchReceiverResponse{}
	_body, _err := client.RuleBatchReceiverWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增跨云存储配置
//
// @param request - SaveAcrossCloudStroageConfigsRequest
//
// @param headers - SaveAcrossCloudStroageConfigsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveAcrossCloudStroageConfigsResponse
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

// Summary:
//
// 新增跨云存储配置
//
// @param request - SaveAcrossCloudStroageConfigsRequest
//
// @return SaveAcrossCloudStroageConfigsResponse
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

// Summary:
//
// 保存并提交认证信息
//
// @param request - SaveAndSubmitAuthInfoRequest
//
// @param headers - SaveAndSubmitAuthInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveAndSubmitAuthInfoResponse
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

// Summary:
//
// 保存并提交认证信息
//
// @param request - SaveAndSubmitAuthInfoRequest
//
// @return SaveAndSubmitAuthInfoResponse
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

// Summary:
//
// 亿格云能力结合
//
// @param request - SaveOpenTerminalInfoRequest
//
// @param headers - SaveOpenTerminalInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveOpenTerminalInfoResponse
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

// Summary:
//
// 亿格云能力结合
//
// @param request - SaveOpenTerminalInfoRequest
//
// @return SaveOpenTerminalInfoResponse
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

// Summary:
//
// 保存专属文件存储的功能项
//
// @param request - SaveStorageFunctionSwitchRequest
//
// @param headers - SaveStorageFunctionSwitchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveStorageFunctionSwitchResponse
func (client *Client) SaveStorageFunctionSwitchWithOptions(request *SaveStorageFunctionSwitchRequest, headers *SaveStorageFunctionSwitchHeaders, runtime *util.RuntimeOptions) (_result *SaveStorageFunctionSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FunctionList)) {
		body["functionList"] = request.FunctionList
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
		Action:      tea.String("SaveStorageFunctionSwitch"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/storages/functions/save"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveStorageFunctionSwitchResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存专属文件存储的功能项
//
// @param request - SaveStorageFunctionSwitchRequest
//
// @return SaveStorageFunctionSwitchResponse
func (client *Client) SaveStorageFunctionSwitch(request *SaveStorageFunctionSwitchRequest) (_result *SaveStorageFunctionSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveStorageFunctionSwitchHeaders{}
	_result = &SaveStorageFunctionSwitchResponse{}
	_body, _err := client.SaveStorageFunctionSwitchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 保存专属文件存储整体开关
//
// @param request - SaveStorageSwitchRequest
//
// @param headers - SaveStorageSwitchHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveStorageSwitchResponse
func (client *Client) SaveStorageSwitchWithOptions(request *SaveStorageSwitchRequest, headers *SaveStorageSwitchHeaders, runtime *util.RuntimeOptions) (_result *SaveStorageSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenStorage)) {
		body["openStorage"] = request.OpenStorage
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
		Action:      tea.String("SaveStorageSwitch"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/storages/switches/save"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveStorageSwitchResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存专属文件存储整体开关
//
// @param request - SaveStorageSwitchRequest
//
// @return SaveStorageSwitchResponse
func (client *Client) SaveStorageSwitch(request *SaveStorageSwitchRequest) (_result *SaveStorageSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveStorageSwitchHeaders{}
	_result = &SaveStorageSwitchResponse{}
	_body, _err := client.SaveStorageSwitchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用于提供mdm微应用白名单配置能力
//
// @param request - SaveWhiteAppRequest
//
// @param headers - SaveWhiteAppHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveWhiteAppResponse
func (client *Client) SaveWhiteAppWithOptions(request *SaveWhiteAppRequest, headers *SaveWhiteAppHeaders, runtime *util.RuntimeOptions) (_result *SaveWhiteAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentIdList)) {
		body["agentIdList"] = request.AgentIdList
	}

	if !tea.BoolValue(util.IsUnset(request.AgentIdMap)) {
		body["agentIdMap"] = request.AgentIdMap
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

// Summary:
//
// 用于提供mdm微应用白名单配置能力
//
// @param request - SaveWhiteAppRequest
//
// @return SaveWhiteAppResponse
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

// Summary:
//
// 查询企业内部群信息
//
// @param request - SearchOrgInnerGroupInfoRequest
//
// @param headers - SearchOrgInnerGroupInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchOrgInnerGroupInfoResponse
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

// Summary:
//
// 查询企业内部群信息
//
// @param request - SearchOrgInnerGroupInfoRequest
//
// @return SearchOrgInnerGroupInfoResponse
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

// Summary:
//
// 通过接口发送应用内DING
//
// @param request - SendAppDingRequest
//
// @param headers - SendAppDingHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendAppDingResponse
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

// Summary:
//
// 通过接口发送应用内DING
//
// @param request - SendAppDingRequest
//
// @return SendAppDingResponse
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

// Summary:
//
// 发送邀请函
//
// @param request - SendInvitationRequest
//
// @param headers - SendInvitationHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendInvitationResponse
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

// Summary:
//
// 发送邀请函
//
// @param request - SendInvitationRequest
//
// @return SendInvitationResponse
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

// Summary:
//
// 通过接口发送电话DING
//
// @param request - SendPhoneDingRequest
//
// @param headers - SendPhoneDingHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendPhoneDingResponse
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

// Summary:
//
// 通过接口发送电话DING
//
// @param request - SendPhoneDingRequest
//
// @return SendPhoneDingResponse
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

// Summary:
//
// 设置会话所属分组
//
// @param request - SetConversationCategoryRequest
//
// @param headers - SetConversationCategoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetConversationCategoryResponse
func (client *Client) SetConversationCategoryWithOptions(request *SetConversationCategoryRequest, headers *SetConversationCategoryHeaders, runtime *util.RuntimeOptions) (_result *SetConversationCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		body["categoryId"] = request.CategoryId
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
		Action:      tea.String("SetConversationCategory"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/conversationCategories/set"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SetConversationCategoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置会话所属分组
//
// @param request - SetConversationCategoryRequest
//
// @return SetConversationCategoryResponse
func (client *Client) SetConversationCategory(request *SetConversationCategoryRequest) (_result *SetConversationCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetConversationCategoryHeaders{}
	_result = &SetConversationCategoryResponse{}
	_body, _err := client.SetConversationCategoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置会话副标题
//
// @param request - SetConversationSubtitleRequest
//
// @param headers - SetConversationSubtitleHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetConversationSubtitleResponse
func (client *Client) SetConversationSubtitleWithOptions(request *SetConversationSubtitleRequest, headers *SetConversationSubtitleHeaders, runtime *util.RuntimeOptions) (_result *SetConversationSubtitleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Subtitle)) {
		body["subtitle"] = request.Subtitle
	}

	if !tea.BoolValue(util.IsUnset(request.SubtitleColor)) {
		body["subtitleColor"] = request.SubtitleColor
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
		Action:      tea.String("SetConversationSubtitle"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/conversations/subtitles/set"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SetConversationSubtitleResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置会话副标题
//
// @param request - SetConversationSubtitleRequest
//
// @return SetConversationSubtitleResponse
func (client *Client) SetConversationSubtitle(request *SetConversationSubtitleRequest) (_result *SetConversationSubtitleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetConversationSubtitleHeaders{}
	_result = &SetConversationSubtitleResponse{}
	_body, _err := client.SetConversationSubtitleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置会话所属顶部分组
//
// @param request - SetConversationTopCategoryRequest
//
// @param headers - SetConversationTopCategoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetConversationTopCategoryResponse
func (client *Client) SetConversationTopCategoryWithOptions(request *SetConversationTopCategoryRequest, headers *SetConversationTopCategoryHeaders, runtime *util.RuntimeOptions) (_result *SetConversationTopCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.SetCategoryList)) {
		body["setCategoryList"] = request.SetCategoryList
	}

	if !tea.BoolValue(util.IsUnset(request.Sign)) {
		body["sign"] = request.Sign
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
		Action:      tea.String("SetConversationTopCategory"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/conversations/topCategories/set"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SetConversationTopCategoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置会话所属顶部分组
//
// @param request - SetConversationTopCategoryRequest
//
// @return SetConversationTopCategoryResponse
func (client *Client) SetConversationTopCategory(request *SetConversationTopCategoryRequest) (_result *SetConversationTopCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetConversationTopCategoryHeaders{}
	_result = &SetConversationTopCategoryResponse{}
	_body, _err := client.SetConversationTopCategoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 伙伴钉设置部门伙伴编码和伙伴类型
//
// @param request - SetDeptPartnerTypeAndNumRequest
//
// @param headers - SetDeptPartnerTypeAndNumHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDeptPartnerTypeAndNumResponse
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

// Summary:
//
// 伙伴钉设置部门伙伴编码和伙伴类型
//
// @param request - SetDeptPartnerTypeAndNumRequest
//
// @return SetDeptPartnerTypeAndNumResponse
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

// Summary:
//
// 设置企业全局顶部会话分组
//
// @param request - SetOrgTopConversationCategoryRequest
//
// @param headers - SetOrgTopConversationCategoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetOrgTopConversationCategoryResponse
func (client *Client) SetOrgTopConversationCategoryWithOptions(request *SetOrgTopConversationCategoryRequest, headers *SetOrgTopConversationCategoryHeaders, runtime *util.RuntimeOptions) (_result *SetOrgTopConversationCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
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
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("SetOrgTopConversationCategory"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/topConversations/categories/set"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SetOrgTopConversationCategoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置企业全局顶部会话分组
//
// @param request - SetOrgTopConversationCategoryRequest
//
// @return SetOrgTopConversationCategoryResponse
func (client *Client) SetOrgTopConversationCategory(request *SetOrgTopConversationCategoryRequest) (_result *SetOrgTopConversationCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetOrgTopConversationCategoryHeaders{}
	_result = &SetOrgTopConversationCategoryResponse{}
	_body, _err := client.SetOrgTopConversationCategoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 千人千面按规则批量发消息
//
// @param request - SpecialRuleBatchReceiverRequest
//
// @param headers - SpecialRuleBatchReceiverHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SpecialRuleBatchReceiverResponse
func (client *Client) SpecialRuleBatchReceiverWithOptions(request *SpecialRuleBatchReceiverRequest, headers *SpecialRuleBatchReceiverHeaders, runtime *util.RuntimeOptions) (_result *SpecialRuleBatchReceiverResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BatchNo)) {
		body["batchNo"] = request.BatchNo
	}

	if !tea.BoolValue(util.IsUnset(request.CardOptions)) {
		body["cardOptions"] = request.CardOptions
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.RuleCode)) {
		body["ruleCode"] = request.RuleCode
	}

	if !tea.BoolValue(util.IsUnset(request.SecretKey)) {
		body["secretKey"] = request.SecretKey
	}

	if !tea.BoolValue(util.IsUnset(request.SpecialStrategy)) {
		body["specialStrategy"] = request.SpecialStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.TaskBatchNo)) {
		body["taskBatchNo"] = request.TaskBatchNo
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
		Action:      tea.String("SpecialRuleBatchReceiver"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/dmc/rules/specialMessages/batchSend"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SpecialRuleBatchReceiverResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 千人千面按规则批量发消息
//
// @param request - SpecialRuleBatchReceiverRequest
//
// @return SpecialRuleBatchReceiverResponse
func (client *Client) SpecialRuleBatchReceiver(request *SpecialRuleBatchReceiverRequest) (_result *SpecialRuleBatchReceiverResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SpecialRuleBatchReceiverHeaders{}
	_result = &SpecialRuleBatchReceiverResponse{}
	_body, _err := client.SpecialRuleBatchReceiverWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 增加/删除任务人员
//
// @param request - TaskInfoAddDelTaskPersonRequest
//
// @param headers - TaskInfoAddDelTaskPersonHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TaskInfoAddDelTaskPersonResponse
func (client *Client) TaskInfoAddDelTaskPersonWithOptions(request *TaskInfoAddDelTaskPersonRequest, headers *TaskInfoAddDelTaskPersonHeaders, runtime *util.RuntimeOptions) (_result *TaskInfoAddDelTaskPersonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperateType)) {
		body["operateType"] = request.OperateType
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorAccount)) {
		body["operatorAccount"] = request.OperatorAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OutTaskId)) {
		body["outTaskId"] = request.OutTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjId)) {
		body["projId"] = request.ProjId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretKey)) {
		body["secretKey"] = request.SecretKey
	}

	if !tea.BoolValue(util.IsUnset(request.TaskExecutePersonDTOS)) {
		body["taskExecutePersonDTOS"] = request.TaskExecutePersonDTOS
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
		Action:      tea.String("TaskInfoAddDelTaskPerson"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/taskCenters/taskInfos/addDelTaskPerson"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TaskInfoAddDelTaskPersonResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 增加/删除任务人员
//
// @param request - TaskInfoAddDelTaskPersonRequest
//
// @return TaskInfoAddDelTaskPersonResponse
func (client *Client) TaskInfoAddDelTaskPerson(request *TaskInfoAddDelTaskPersonRequest) (_result *TaskInfoAddDelTaskPersonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TaskInfoAddDelTaskPersonHeaders{}
	_result = &TaskInfoAddDelTaskPersonResponse{}
	_body, _err := client.TaskInfoAddDelTaskPersonWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除任务
//
// @param request - TaskInfoCancelOrDelTaskRequest
//
// @param headers - TaskInfoCancelOrDelTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TaskInfoCancelOrDelTaskResponse
func (client *Client) TaskInfoCancelOrDelTaskWithOptions(request *TaskInfoCancelOrDelTaskRequest, headers *TaskInfoCancelOrDelTaskHeaders, runtime *util.RuntimeOptions) (_result *TaskInfoCancelOrDelTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CardDTO)) {
		body["cardDTO"] = request.CardDTO
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorAccount)) {
		body["operatorAccount"] = request.OperatorAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OutTaskId)) {
		body["outTaskId"] = request.OutTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjId)) {
		body["projId"] = request.ProjId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretKey)) {
		body["secretKey"] = request.SecretKey
	}

	if !tea.BoolValue(util.IsUnset(request.SendMsgFlag)) {
		body["sendMsgFlag"] = request.SendMsgFlag
	}

	if !tea.BoolValue(util.IsUnset(request.TaskExecutePersonDTOS)) {
		body["taskExecutePersonDTOS"] = request.TaskExecutePersonDTOS
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
		Action:      tea.String("TaskInfoCancelOrDelTask"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/taskCenters/taskInfos/cancelOrDelTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TaskInfoCancelOrDelTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除任务
//
// @param request - TaskInfoCancelOrDelTaskRequest
//
// @return TaskInfoCancelOrDelTaskResponse
func (client *Client) TaskInfoCancelOrDelTask(request *TaskInfoCancelOrDelTaskRequest) (_result *TaskInfoCancelOrDelTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TaskInfoCancelOrDelTaskHeaders{}
	_result = &TaskInfoCancelOrDelTaskResponse{}
	_body, _err := client.TaskInfoCancelOrDelTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建并启动任务
//
// @param request - TaskInfoCreateAndStartTaskRequest
//
// @param headers - TaskInfoCreateAndStartTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TaskInfoCreateAndStartTaskResponse
func (client *Client) TaskInfoCreateAndStartTaskWithOptions(request *TaskInfoCreateAndStartTaskRequest, headers *TaskInfoCreateAndStartTaskHeaders, runtime *util.RuntimeOptions) (_result *TaskInfoCreateAndStartTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Attr)) {
		body["attr"] = request.Attr
	}

	if !tea.BoolValue(util.IsUnset(request.BacklogDTO)) {
		body["backlogDTO"] = request.BacklogDTO
	}

	if !tea.BoolValue(util.IsUnset(request.BacklogGenerateFlag)) {
		body["backlogGenerateFlag"] = request.BacklogGenerateFlag
	}

	if !tea.BoolValue(util.IsUnset(request.BusinessCode)) {
		body["businessCode"] = request.BusinessCode
	}

	if !tea.BoolValue(util.IsUnset(request.CanceldelTaskCardId)) {
		body["canceldelTaskCardId"] = request.CanceldelTaskCardId
	}

	if !tea.BoolValue(util.IsUnset(request.CardDTO)) {
		body["cardDTO"] = request.CardDTO
	}

	if !tea.BoolValue(util.IsUnset(request.CustomFlag)) {
		body["customFlag"] = request.CustomFlag
	}

	if !tea.BoolValue(util.IsUnset(request.DetailUrl)) {
		body["detailUrl"] = request.DetailUrl
	}

	if !tea.BoolValue(util.IsUnset(request.FinishTaskCardId)) {
		body["finishTaskCardId"] = request.FinishTaskCardId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorAccount)) {
		body["operatorAccount"] = request.OperatorAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OutTaskId)) {
		body["outTaskId"] = request.OutTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjId)) {
		body["projId"] = request.ProjId
	}

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
	}

	if !tea.BoolValue(util.IsUnset(request.SecretKey)) {
		body["secretKey"] = request.SecretKey
	}

	if !tea.BoolValue(util.IsUnset(request.SendMsgFlag)) {
		body["sendMsgFlag"] = request.SendMsgFlag
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		body["sort"] = request.Sort
	}

	if !tea.BoolValue(util.IsUnset(request.StartTaskCardId)) {
		body["startTaskCardId"] = request.StartTaskCardId
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		body["state"] = request.State
	}

	if !tea.BoolValue(util.IsUnset(request.TaskContent)) {
		body["taskContent"] = request.TaskContent
	}

	if !tea.BoolValue(util.IsUnset(request.TaskEndTime)) {
		body["taskEndTime"] = request.TaskEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskExecutePersonDTOS)) {
		body["taskExecutePersonDTOS"] = request.TaskExecutePersonDTOS
	}

	if !tea.BoolValue(util.IsUnset(request.TaskGroupDTOList)) {
		body["taskGroupDTOList"] = request.TaskGroupDTOList
	}

	if !tea.BoolValue(util.IsUnset(request.TaskSystem)) {
		body["taskSystem"] = request.TaskSystem
	}

	if !tea.BoolValue(util.IsUnset(request.TaskTemplCode)) {
		body["taskTemplCode"] = request.TaskTemplCode
	}

	if !tea.BoolValue(util.IsUnset(request.TaskTitle)) {
		body["taskTitle"] = request.TaskTitle
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		body["taskType"] = request.TaskType
	}

	if !tea.BoolValue(util.IsUnset(request.TaskUrlMobile)) {
		body["taskUrlMobile"] = request.TaskUrlMobile
	}

	if !tea.BoolValue(util.IsUnset(request.TaskUrlPc)) {
		body["taskUrlPc"] = request.TaskUrlPc
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateTaskCardId)) {
		body["updateTaskCardId"] = request.UpdateTaskCardId
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
		Action:      tea.String("TaskInfoCreateAndStartTask"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/taskCenters/taskInfos/createAndStart"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TaskInfoCreateAndStartTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建并启动任务
//
// @param request - TaskInfoCreateAndStartTaskRequest
//
// @return TaskInfoCreateAndStartTaskResponse
func (client *Client) TaskInfoCreateAndStartTask(request *TaskInfoCreateAndStartTaskRequest) (_result *TaskInfoCreateAndStartTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TaskInfoCreateAndStartTaskHeaders{}
	_result = &TaskInfoCreateAndStartTaskResponse{}
	_body, _err := client.TaskInfoCreateAndStartTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 完成任务
//
// @param request - TaskInfoFinishTaskRequest
//
// @param headers - TaskInfoFinishTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TaskInfoFinishTaskResponse
func (client *Client) TaskInfoFinishTaskWithOptions(request *TaskInfoFinishTaskRequest, headers *TaskInfoFinishTaskHeaders, runtime *util.RuntimeOptions) (_result *TaskInfoFinishTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CardDTO)) {
		body["cardDTO"] = request.CardDTO
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorAccount)) {
		body["operatorAccount"] = request.OperatorAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OutTaskId)) {
		body["outTaskId"] = request.OutTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjId)) {
		body["projId"] = request.ProjId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretKey)) {
		body["secretKey"] = request.SecretKey
	}

	if !tea.BoolValue(util.IsUnset(request.SendMsgFlag)) {
		body["sendMsgFlag"] = request.SendMsgFlag
	}

	if !tea.BoolValue(util.IsUnset(request.TaskExecutePersonDTOS)) {
		body["taskExecutePersonDTOS"] = request.TaskExecutePersonDTOS
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
		Action:      tea.String("TaskInfoFinishTask"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/taskCenters/taskInfos/finishTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TaskInfoFinishTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 完成任务
//
// @param request - TaskInfoFinishTaskRequest
//
// @return TaskInfoFinishTaskResponse
func (client *Client) TaskInfoFinishTask(request *TaskInfoFinishTaskRequest) (_result *TaskInfoFinishTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TaskInfoFinishTaskHeaders{}
	_result = &TaskInfoFinishTaskResponse{}
	_body, _err := client.TaskInfoFinishTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新任务
//
// @param request - TaskInfoUpdateTaskRequest
//
// @param headers - TaskInfoUpdateTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TaskInfoUpdateTaskResponse
func (client *Client) TaskInfoUpdateTaskWithOptions(request *TaskInfoUpdateTaskRequest, headers *TaskInfoUpdateTaskHeaders, runtime *util.RuntimeOptions) (_result *TaskInfoUpdateTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Attr)) {
		body["attr"] = request.Attr
	}

	if !tea.BoolValue(util.IsUnset(request.CanceldelTaskCardId)) {
		body["canceldelTaskCardId"] = request.CanceldelTaskCardId
	}

	if !tea.BoolValue(util.IsUnset(request.CardDTO)) {
		body["cardDTO"] = request.CardDTO
	}

	if !tea.BoolValue(util.IsUnset(request.DetailUrl)) {
		body["detailUrl"] = request.DetailUrl
	}

	if !tea.BoolValue(util.IsUnset(request.FinishTaskCardId)) {
		body["finishTaskCardId"] = request.FinishTaskCardId
	}

	if !tea.BoolValue(util.IsUnset(request.ListOpenConversationId)) {
		body["listOpenConversationId"] = request.ListOpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OperateType)) {
		body["operateType"] = request.OperateType
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorAccount)) {
		body["operatorAccount"] = request.OperatorAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OutTaskId)) {
		body["outTaskId"] = request.OutTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjId)) {
		body["projId"] = request.ProjId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretKey)) {
		body["secretKey"] = request.SecretKey
	}

	if !tea.BoolValue(util.IsUnset(request.SendMsgFlag)) {
		body["sendMsgFlag"] = request.SendMsgFlag
	}

	if !tea.BoolValue(util.IsUnset(request.StartTaskCardId)) {
		body["startTaskCardId"] = request.StartTaskCardId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskContent)) {
		body["taskContent"] = request.TaskContent
	}

	if !tea.BoolValue(util.IsUnset(request.TaskEndTime)) {
		body["taskEndTime"] = request.TaskEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskExecutePersonDTOS)) {
		body["taskExecutePersonDTOS"] = request.TaskExecutePersonDTOS
	}

	if !tea.BoolValue(util.IsUnset(request.TaskTitle)) {
		body["taskTitle"] = request.TaskTitle
	}

	if !tea.BoolValue(util.IsUnset(request.TaskUrlMobile)) {
		body["taskUrlMobile"] = request.TaskUrlMobile
	}

	if !tea.BoolValue(util.IsUnset(request.TaskUrlPc)) {
		body["taskUrlPc"] = request.TaskUrlPc
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateTaskCardId)) {
		body["updateTaskCardId"] = request.UpdateTaskCardId
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
		Action:      tea.String("TaskInfoUpdateTask"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/taskCenters/taskInfos/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TaskInfoUpdateTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新任务
//
// @param request - TaskInfoUpdateTaskRequest
//
// @return TaskInfoUpdateTaskResponse
func (client *Client) TaskInfoUpdateTask(request *TaskInfoUpdateTaskRequest) (_result *TaskInfoUpdateTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TaskInfoUpdateTaskHeaders{}
	_result = &TaskInfoUpdateTaskResponse{}
	_body, _err := client.TaskInfoUpdateTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 切换组织归属
//
// @param request - TransferExclusiveAccountOrgRequest
//
// @param headers - TransferExclusiveAccountOrgHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferExclusiveAccountOrgResponse
func (client *Client) TransferExclusiveAccountOrgWithOptions(request *TransferExclusiveAccountOrgRequest, headers *TransferExclusiveAccountOrgHeaders, runtime *util.RuntimeOptions) (_result *TransferExclusiveAccountOrgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsSettingMainOrg)) {
		body["isSettingMainOrg"] = request.IsSettingMainOrg
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCorpId)) {
		body["targetCorpId"] = request.TargetCorpId
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
		Action:      tea.String("TransferExclusiveAccountOrg"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/organizations/transfer"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TransferExclusiveAccountOrgResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 切换组织归属
//
// @param request - TransferExclusiveAccountOrgRequest
//
// @return TransferExclusiveAccountOrgResponse
func (client *Client) TransferExclusiveAccountOrg(request *TransferExclusiveAccountOrgRequest) (_result *TransferExclusiveAccountOrgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TransferExclusiveAccountOrgHeaders{}
	_result = &TransferExclusiveAccountOrgResponse{}
	_body, _err := client.TransferExclusiveAccountOrgWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更改分组名称
//
// @param request - UpdateCategoryNameRequest
//
// @param headers - UpdateCategoryNameHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCategoryNameResponse
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

// Summary:
//
// 更改分组名称
//
// @param request - UpdateCategoryNameRequest
//
// @return UpdateCategoryNameResponse
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

// Summary:
//
// 变更群聊类型
//
// @param request - UpdateConversationTypeRequest
//
// @param headers - UpdateConversationTypeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConversationTypeResponse
func (client *Client) UpdateConversationTypeWithOptions(request *UpdateConversationTypeRequest, headers *UpdateConversationTypeHeaders, runtime *util.RuntimeOptions) (_result *UpdateConversationTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ManageSign)) {
		body["manageSign"] = request.ManageSign
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
		Action:      tea.String("UpdateConversationType"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/conversationTypes"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateConversationTypeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 变更群聊类型
//
// @param request - UpdateConversationTypeRequest
//
// @return UpdateConversationTypeResponse
func (client *Client) UpdateConversationType(request *UpdateConversationTypeRequest) (_result *UpdateConversationTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateConversationTypeHeaders{}
	_result = &UpdateConversationTypeResponse{}
	_body, _err := client.UpdateConversationTypeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新发送文件的检测状态
//
// @param request - UpdateFileStatusRequest
//
// @param headers - UpdateFileStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFileStatusResponse
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

// Summary:
//
// 更新发送文件的检测状态
//
// @param request - UpdateFileStatusRequest
//
// @return UpdateFileStatusResponse
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

// Summary:
//
// 发布版本
//
// @param request - UpdateMiniAppVersionStatusRequest
//
// @param headers - UpdateMiniAppVersionStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMiniAppVersionStatusResponse
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

// Summary:
//
// 发布版本
//
// @param request - UpdateMiniAppVersionStatusRequest
//
// @return UpdateMiniAppVersionStatusResponse
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

// Summary:
//
// 修改伙伴类型可见性
//
// @param request - UpdatePartnerVisibilityRequest
//
// @param headers - UpdatePartnerVisibilityHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePartnerVisibilityResponse
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

// Summary:
//
// 修改伙伴类型可见性
//
// @param request - UpdatePartnerVisibilityRequest
//
// @return UpdatePartnerVisibilityResponse
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

// Summary:
//
// 专属一线版-企业修改员工license
//
// @param request - UpdateRealmLicenseRequest
//
// @param headers - UpdateRealmLicenseHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRealmLicenseResponse
func (client *Client) UpdateRealmLicenseWithOptions(request *UpdateRealmLicenseRequest, headers *UpdateRealmLicenseHeaders, runtime *util.RuntimeOptions) (_result *UpdateRealmLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DetailList)) {
		body["detailList"] = request.DetailList
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
		Action:      tea.String("UpdateRealmLicense"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/frontLines/licenses"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRealmLicenseResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 专属一线版-企业修改员工license
//
// @param request - UpdateRealmLicenseRequest
//
// @return UpdateRealmLicenseResponse
func (client *Client) UpdateRealmLicense(request *UpdateRealmLicenseRequest) (_result *UpdateRealmLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateRealmLicenseHeaders{}
	_result = &UpdateRealmLicenseResponse{}
	_body, _err := client.UpdateRealmLicenseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改角色可见性
//
// @param request - UpdateRoleVisibilityRequest
//
// @param headers - UpdateRoleVisibilityHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRoleVisibilityResponse
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

// Summary:
//
// 修改角色可见性
//
// @param request - UpdateRoleVisibilityRequest
//
// @return UpdateRoleVisibilityResponse
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

// Summary:
//
// 更新组织专属存储模式
//
// @param request - UpdateStorageModeRequest
//
// @param headers - UpdateStorageModeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStorageModeResponse
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

// Summary:
//
// 更新组织专属存储模式
//
// @param request - UpdateStorageModeRequest
//
// @return UpdateStorageModeResponse
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

// Summary:
//
// 通过设备编号修改设备信息。
//
// @param request - UpdateTrustedDeviceRequest
//
// @param headers - UpdateTrustedDeviceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTrustedDeviceResponse
func (client *Client) UpdateTrustedDeviceWithOptions(deviceId *string, request *UpdateTrustedDeviceRequest, headers *UpdateTrustedDeviceHeaders, runtime *util.RuntimeOptions) (_result *UpdateTrustedDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
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
		Action:      tea.String("UpdateTrustedDevice"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/trustedDevices/" + tea.StringValue(deviceId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTrustedDeviceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过设备编号修改设备信息。
//
// @param request - UpdateTrustedDeviceRequest
//
// @return UpdateTrustedDeviceResponse
func (client *Client) UpdateTrustedDevice(deviceId *string, request *UpdateTrustedDeviceRequest) (_result *UpdateTrustedDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTrustedDeviceHeaders{}
	_result = &UpdateTrustedDeviceResponse{}
	_body, _err := client.UpdateTrustedDeviceWithOptions(deviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 允许三方调用该API，决定对应的语音消息管控状态
//
// @param request - UpdateVoiceMsgCtrlStatusRequest
//
// @param headers - UpdateVoiceMsgCtrlStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVoiceMsgCtrlStatusResponse
func (client *Client) UpdateVoiceMsgCtrlStatusWithOptions(request *UpdateVoiceMsgCtrlStatusRequest, headers *UpdateVoiceMsgCtrlStatusHeaders, runtime *util.RuntimeOptions) (_result *UpdateVoiceMsgCtrlStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.VoiceMsgCtrlInfo)) {
		body["voiceMsgCtrlInfo"] = request.VoiceMsgCtrlInfo
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
		Action:      tea.String("UpdateVoiceMsgCtrlStatus"),
		Version:     tea.String("exclusive_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/exclusive/voiceMessages/ctrlStatuses"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateVoiceMsgCtrlStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 允许三方调用该API，决定对应的语音消息管控状态
//
// @param request - UpdateVoiceMsgCtrlStatusRequest
//
// @return UpdateVoiceMsgCtrlStatusResponse
func (client *Client) UpdateVoiceMsgCtrlStatus(request *UpdateVoiceMsgCtrlStatusRequest) (_result *UpdateVoiceMsgCtrlStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateVoiceMsgCtrlStatusHeaders{}
	_result = &UpdateVoiceMsgCtrlStatusResponse{}
	_body, _err := client.UpdateVoiceMsgCtrlStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
