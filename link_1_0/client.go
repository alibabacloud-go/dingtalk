// This file is auto-generated, don't edit it. Thanks.
package link_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DetailUserIdPrivateDataMapValue struct {
	CardParamMap        map[string]interface{} `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
	CardMediaIdParamMap map[string]interface{} `json:"cardMediaIdParamMap,omitempty" xml:"cardMediaIdParamMap,omitempty"`
}

func (s DetailUserIdPrivateDataMapValue) String() string {
	return tea.Prettify(s)
}

func (s DetailUserIdPrivateDataMapValue) GoString() string {
	return s.String()
}

func (s *DetailUserIdPrivateDataMapValue) SetCardParamMap(v map[string]interface{}) *DetailUserIdPrivateDataMapValue {
	s.CardParamMap = v
	return s
}

func (s *DetailUserIdPrivateDataMapValue) SetCardMediaIdParamMap(v map[string]interface{}) *DetailUserIdPrivateDataMapValue {
	s.CardMediaIdParamMap = v
	return s
}

type ApplyFollowerAuthInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ApplyFollowerAuthInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s ApplyFollowerAuthInfoHeaders) GoString() string {
	return s.String()
}

func (s *ApplyFollowerAuthInfoHeaders) SetCommonHeaders(v map[string]*string) *ApplyFollowerAuthInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ApplyFollowerAuthInfoHeaders) SetXAcsDingtalkAccessToken(v string) *ApplyFollowerAuthInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ApplyFollowerAuthInfoRequest struct {
	AppAuthKey *string `json:"appAuthKey,omitempty" xml:"appAuthKey,omitempty"`
	// example:
	//
	// Contact.User.mobile
	FieldScope *string `json:"fieldScope,omitempty" xml:"fieldScope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sid22b31b4bf59ef2c783f7
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// idzb26bxl64vqx2keyi
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ApplyFollowerAuthInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyFollowerAuthInfoRequest) GoString() string {
	return s.String()
}

func (s *ApplyFollowerAuthInfoRequest) SetAppAuthKey(v string) *ApplyFollowerAuthInfoRequest {
	s.AppAuthKey = &v
	return s
}

func (s *ApplyFollowerAuthInfoRequest) SetFieldScope(v string) *ApplyFollowerAuthInfoRequest {
	s.FieldScope = &v
	return s
}

func (s *ApplyFollowerAuthInfoRequest) SetSessionId(v string) *ApplyFollowerAuthInfoRequest {
	s.SessionId = &v
	return s
}

func (s *ApplyFollowerAuthInfoRequest) SetUserId(v string) *ApplyFollowerAuthInfoRequest {
	s.UserId = &v
	return s
}

type ApplyFollowerAuthInfoResponseBody struct {
	// This parameter is required.
	Result *ApplyFollowerAuthInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ApplyFollowerAuthInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyFollowerAuthInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyFollowerAuthInfoResponseBody) SetResult(v *ApplyFollowerAuthInfoResponseBodyResult) *ApplyFollowerAuthInfoResponseBody {
	s.Result = v
	return s
}

type ApplyFollowerAuthInfoResponseBodyResult struct {
	// This parameter is required.
	OpenApplyId *string `json:"openApplyId,omitempty" xml:"openApplyId,omitempty"`
}

func (s ApplyFollowerAuthInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ApplyFollowerAuthInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ApplyFollowerAuthInfoResponseBodyResult) SetOpenApplyId(v string) *ApplyFollowerAuthInfoResponseBodyResult {
	s.OpenApplyId = &v
	return s
}

type ApplyFollowerAuthInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyFollowerAuthInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyFollowerAuthInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyFollowerAuthInfoResponse) GoString() string {
	return s.String()
}

func (s *ApplyFollowerAuthInfoResponse) SetHeaders(v map[string]*string) *ApplyFollowerAuthInfoResponse {
	s.Headers = v
	return s
}

func (s *ApplyFollowerAuthInfoResponse) SetStatusCode(v int32) *ApplyFollowerAuthInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyFollowerAuthInfoResponse) SetBody(v *ApplyFollowerAuthInfoResponseBody) *ApplyFollowerAuthInfoResponse {
	s.Body = v
	return s
}

type CallbackRegiesterHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CallbackRegiesterHeaders) String() string {
	return tea.Prettify(s)
}

func (s CallbackRegiesterHeaders) GoString() string {
	return s.String()
}

func (s *CallbackRegiesterHeaders) SetCommonHeaders(v map[string]*string) *CallbackRegiesterHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CallbackRegiesterHeaders) SetXAcsDingtalkAccessToken(v string) *CallbackRegiesterHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CallbackRegiesterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3b89df4dfaaccd5b8e1f9a2
	ApiSecret *string `json:"apiSecret,omitempty" xml:"apiSecret,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abc-123
	CallbackKey *string `json:"callbackKey,omitempty" xml:"callbackKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://www.your.com/callback
	CallbackUrl *string `json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// shortcut
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CallbackRegiesterRequest) String() string {
	return tea.Prettify(s)
}

func (s CallbackRegiesterRequest) GoString() string {
	return s.String()
}

func (s *CallbackRegiesterRequest) SetApiSecret(v string) *CallbackRegiesterRequest {
	s.ApiSecret = &v
	return s
}

func (s *CallbackRegiesterRequest) SetCallbackKey(v string) *CallbackRegiesterRequest {
	s.CallbackKey = &v
	return s
}

func (s *CallbackRegiesterRequest) SetCallbackUrl(v string) *CallbackRegiesterRequest {
	s.CallbackUrl = &v
	return s
}

func (s *CallbackRegiesterRequest) SetType(v string) *CallbackRegiesterRequest {
	s.Type = &v
	return s
}

type CallbackRegiesterResponseBody struct {
	// This parameter is required.
	Result *CallbackRegiesterResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CallbackRegiesterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CallbackRegiesterResponseBody) GoString() string {
	return s.String()
}

func (s *CallbackRegiesterResponseBody) SetResult(v *CallbackRegiesterResponseBodyResult) *CallbackRegiesterResponseBody {
	s.Result = v
	return s
}

type CallbackRegiesterResponseBodyResult struct {
	ApiSecret *string `json:"apiSecret,omitempty" xml:"apiSecret,omitempty"`
	// This parameter is required.
	CallbackUrl *string `json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`
}

func (s CallbackRegiesterResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CallbackRegiesterResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CallbackRegiesterResponseBodyResult) SetApiSecret(v string) *CallbackRegiesterResponseBodyResult {
	s.ApiSecret = &v
	return s
}

func (s *CallbackRegiesterResponseBodyResult) SetCallbackUrl(v string) *CallbackRegiesterResponseBodyResult {
	s.CallbackUrl = &v
	return s
}

type CallbackRegiesterResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CallbackRegiesterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CallbackRegiesterResponse) String() string {
	return tea.Prettify(s)
}

func (s CallbackRegiesterResponse) GoString() string {
	return s.String()
}

func (s *CallbackRegiesterResponse) SetHeaders(v map[string]*string) *CallbackRegiesterResponse {
	s.Headers = v
	return s
}

func (s *CallbackRegiesterResponse) SetStatusCode(v int32) *CallbackRegiesterResponse {
	s.StatusCode = &v
	return s
}

func (s *CallbackRegiesterResponse) SetBody(v *CallbackRegiesterResponseBody) *CallbackRegiesterResponse {
	s.Body = v
	return s
}

type CloseTopBoxInteractiveOTOMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CloseTopBoxInteractiveOTOMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s CloseTopBoxInteractiveOTOMessageHeaders) GoString() string {
	return s.String()
}

func (s *CloseTopBoxInteractiveOTOMessageHeaders) SetCommonHeaders(v map[string]*string) *CloseTopBoxInteractiveOTOMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CloseTopBoxInteractiveOTOMessageHeaders) SetXAcsDingtalkAccessToken(v string) *CloseTopBoxInteractiveOTOMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CloseTopBoxInteractiveOTOMessageRequest struct {
	// This parameter is required.
	Detail *CloseTopBoxInteractiveOTOMessageRequestDetail `json:"detail,omitempty" xml:"detail,omitempty" type:"Struct"`
}

func (s CloseTopBoxInteractiveOTOMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseTopBoxInteractiveOTOMessageRequest) GoString() string {
	return s.String()
}

func (s *CloseTopBoxInteractiveOTOMessageRequest) SetDetail(v *CloseTopBoxInteractiveOTOMessageRequestDetail) *CloseTopBoxInteractiveOTOMessageRequest {
	s.Detail = v
	return s
}

type CloseTopBoxInteractiveOTOMessageRequestDetail struct {
	// This parameter is required.
	//
	// example:
	//
	// service-card-20220824-001
	CardBizId *string `json:"cardBizId,omitempty" xml:"cardBizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3erkfi-42b0-4c83-bc56-ffhssde43
	CardTemplateId *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CloseTopBoxInteractiveOTOMessageRequestDetail) String() string {
	return tea.Prettify(s)
}

func (s CloseTopBoxInteractiveOTOMessageRequestDetail) GoString() string {
	return s.String()
}

func (s *CloseTopBoxInteractiveOTOMessageRequestDetail) SetCardBizId(v string) *CloseTopBoxInteractiveOTOMessageRequestDetail {
	s.CardBizId = &v
	return s
}

func (s *CloseTopBoxInteractiveOTOMessageRequestDetail) SetCardTemplateId(v string) *CloseTopBoxInteractiveOTOMessageRequestDetail {
	s.CardTemplateId = &v
	return s
}

func (s *CloseTopBoxInteractiveOTOMessageRequestDetail) SetUserId(v string) *CloseTopBoxInteractiveOTOMessageRequestDetail {
	s.UserId = &v
	return s
}

type CloseTopBoxInteractiveOTOMessageResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CloseTopBoxInteractiveOTOMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseTopBoxInteractiveOTOMessageResponseBody) GoString() string {
	return s.String()
}

func (s *CloseTopBoxInteractiveOTOMessageResponseBody) SetRequestId(v string) *CloseTopBoxInteractiveOTOMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseTopBoxInteractiveOTOMessageResponseBody) SetResult(v bool) *CloseTopBoxInteractiveOTOMessageResponseBody {
	s.Result = &v
	return s
}

type CloseTopBoxInteractiveOTOMessageResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseTopBoxInteractiveOTOMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseTopBoxInteractiveOTOMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseTopBoxInteractiveOTOMessageResponse) GoString() string {
	return s.String()
}

func (s *CloseTopBoxInteractiveOTOMessageResponse) SetHeaders(v map[string]*string) *CloseTopBoxInteractiveOTOMessageResponse {
	s.Headers = v
	return s
}

func (s *CloseTopBoxInteractiveOTOMessageResponse) SetStatusCode(v int32) *CloseTopBoxInteractiveOTOMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseTopBoxInteractiveOTOMessageResponse) SetBody(v *CloseTopBoxInteractiveOTOMessageResponseBody) *CloseTopBoxInteractiveOTOMessageResponse {
	s.Body = v
	return s
}

type GetFollowerAuthInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFollowerAuthInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFollowerAuthInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetFollowerAuthInfoHeaders) SetCommonHeaders(v map[string]*string) *GetFollowerAuthInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFollowerAuthInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetFollowerAuthInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFollowerAuthInfoRequest struct {
	// example:
	//
	// ding1234
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Rp3Rqcts7BE08y49Jr6iu6xW4iQ
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetFollowerAuthInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFollowerAuthInfoRequest) GoString() string {
	return s.String()
}

func (s *GetFollowerAuthInfoRequest) SetAccountId(v string) *GetFollowerAuthInfoRequest {
	s.AccountId = &v
	return s
}

func (s *GetFollowerAuthInfoRequest) SetUserId(v string) *GetFollowerAuthInfoRequest {
	s.UserId = &v
	return s
}

type GetFollowerAuthInfoResponseBody struct {
	// This parameter is required.
	Result *GetFollowerAuthInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetFollowerAuthInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFollowerAuthInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetFollowerAuthInfoResponseBody) SetResult(v *GetFollowerAuthInfoResponseBodyResult) *GetFollowerAuthInfoResponseBody {
	s.Result = v
	return s
}

type GetFollowerAuthInfoResponseBodyResult struct {
	AuthInfo *GetFollowerAuthInfoResponseBodyResultAuthInfo `json:"authInfo,omitempty" xml:"authInfo,omitempty" type:"Struct"`
}

func (s GetFollowerAuthInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFollowerAuthInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFollowerAuthInfoResponseBodyResult) SetAuthInfo(v *GetFollowerAuthInfoResponseBodyResultAuthInfo) *GetFollowerAuthInfoResponseBodyResult {
	s.AuthInfo = v
	return s
}

type GetFollowerAuthInfoResponseBodyResultAuthInfo struct {
	MainCorp *GetFollowerAuthInfoResponseBodyResultAuthInfoMainCorp `json:"mainCorp,omitempty" xml:"mainCorp,omitempty" type:"Struct"`
	Mobile   *GetFollowerAuthInfoResponseBodyResultAuthInfoMobile   `json:"mobile,omitempty" xml:"mobile,omitempty" type:"Struct"`
}

func (s GetFollowerAuthInfoResponseBodyResultAuthInfo) String() string {
	return tea.Prettify(s)
}

func (s GetFollowerAuthInfoResponseBodyResultAuthInfo) GoString() string {
	return s.String()
}

func (s *GetFollowerAuthInfoResponseBodyResultAuthInfo) SetMainCorp(v *GetFollowerAuthInfoResponseBodyResultAuthInfoMainCorp) *GetFollowerAuthInfoResponseBodyResultAuthInfo {
	s.MainCorp = v
	return s
}

func (s *GetFollowerAuthInfoResponseBodyResultAuthInfo) SetMobile(v *GetFollowerAuthInfoResponseBodyResultAuthInfoMobile) *GetFollowerAuthInfoResponseBodyResultAuthInfo {
	s.Mobile = v
	return s
}

type GetFollowerAuthInfoResponseBodyResultAuthInfoMainCorp struct {
	Authorized *bool   `json:"authorized,omitempty" xml:"authorized,omitempty"`
	CorpName   *string `json:"corpName,omitempty" xml:"corpName,omitempty"`
}

func (s GetFollowerAuthInfoResponseBodyResultAuthInfoMainCorp) String() string {
	return tea.Prettify(s)
}

func (s GetFollowerAuthInfoResponseBodyResultAuthInfoMainCorp) GoString() string {
	return s.String()
}

func (s *GetFollowerAuthInfoResponseBodyResultAuthInfoMainCorp) SetAuthorized(v bool) *GetFollowerAuthInfoResponseBodyResultAuthInfoMainCorp {
	s.Authorized = &v
	return s
}

func (s *GetFollowerAuthInfoResponseBodyResultAuthInfoMainCorp) SetCorpName(v string) *GetFollowerAuthInfoResponseBodyResultAuthInfoMainCorp {
	s.CorpName = &v
	return s
}

type GetFollowerAuthInfoResponseBodyResultAuthInfoMobile struct {
	Authorized *bool   `json:"authorized,omitempty" xml:"authorized,omitempty"`
	Mobile     *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	StateCode  *string `json:"stateCode,omitempty" xml:"stateCode,omitempty"`
}

func (s GetFollowerAuthInfoResponseBodyResultAuthInfoMobile) String() string {
	return tea.Prettify(s)
}

func (s GetFollowerAuthInfoResponseBodyResultAuthInfoMobile) GoString() string {
	return s.String()
}

func (s *GetFollowerAuthInfoResponseBodyResultAuthInfoMobile) SetAuthorized(v bool) *GetFollowerAuthInfoResponseBodyResultAuthInfoMobile {
	s.Authorized = &v
	return s
}

func (s *GetFollowerAuthInfoResponseBodyResultAuthInfoMobile) SetMobile(v string) *GetFollowerAuthInfoResponseBodyResultAuthInfoMobile {
	s.Mobile = &v
	return s
}

func (s *GetFollowerAuthInfoResponseBodyResultAuthInfoMobile) SetStateCode(v string) *GetFollowerAuthInfoResponseBodyResultAuthInfoMobile {
	s.StateCode = &v
	return s
}

type GetFollowerAuthInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFollowerAuthInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFollowerAuthInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFollowerAuthInfoResponse) GoString() string {
	return s.String()
}

func (s *GetFollowerAuthInfoResponse) SetHeaders(v map[string]*string) *GetFollowerAuthInfoResponse {
	s.Headers = v
	return s
}

func (s *GetFollowerAuthInfoResponse) SetStatusCode(v int32) *GetFollowerAuthInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFollowerAuthInfoResponse) SetBody(v *GetFollowerAuthInfoResponseBody) *GetFollowerAuthInfoResponse {
	s.Body = v
	return s
}

type GetFollowerInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFollowerInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFollowerInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetFollowerInfoHeaders) SetCommonHeaders(v map[string]*string) *GetFollowerInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFollowerInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetFollowerInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFollowerInfoRequest struct {
	// example:
	//
	// ding1234
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// UgIzXXo+Rp3Rqcts7BE08y49Jr6iu6xW4iQ
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// example:
	//
	// Rp3Rqcts7BE08y49Jr6iu6xW4iQ
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetFollowerInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFollowerInfoRequest) GoString() string {
	return s.String()
}

func (s *GetFollowerInfoRequest) SetAccountId(v string) *GetFollowerInfoRequest {
	s.AccountId = &v
	return s
}

func (s *GetFollowerInfoRequest) SetUnionId(v string) *GetFollowerInfoRequest {
	s.UnionId = &v
	return s
}

func (s *GetFollowerInfoRequest) SetUserId(v string) *GetFollowerInfoRequest {
	s.UserId = &v
	return s
}

type GetFollowerInfoResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	Result *GetFollowerInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetFollowerInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFollowerInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetFollowerInfoResponseBody) SetRequestId(v string) *GetFollowerInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFollowerInfoResponseBody) SetResult(v *GetFollowerInfoResponseBodyResult) *GetFollowerInfoResponseBody {
	s.Result = v
	return s
}

type GetFollowerInfoResponseBodyResult struct {
	User *GetFollowerInfoResponseBodyResultUser `json:"user,omitempty" xml:"user,omitempty" type:"Struct"`
}

func (s GetFollowerInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFollowerInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFollowerInfoResponseBodyResult) SetUser(v *GetFollowerInfoResponseBodyResultUser) *GetFollowerInfoResponseBodyResult {
	s.User = v
	return s
}

type GetFollowerInfoResponseBodyResultUser struct {
	// example:
	//
	// 小钉
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1661918406748
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// example:
	//
	// userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetFollowerInfoResponseBodyResultUser) String() string {
	return tea.Prettify(s)
}

func (s GetFollowerInfoResponseBodyResultUser) GoString() string {
	return s.String()
}

func (s *GetFollowerInfoResponseBodyResultUser) SetName(v string) *GetFollowerInfoResponseBodyResultUser {
	s.Name = &v
	return s
}

func (s *GetFollowerInfoResponseBodyResultUser) SetTimestamp(v string) *GetFollowerInfoResponseBodyResultUser {
	s.Timestamp = &v
	return s
}

func (s *GetFollowerInfoResponseBodyResultUser) SetUserId(v string) *GetFollowerInfoResponseBodyResultUser {
	s.UserId = &v
	return s
}

type GetFollowerInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFollowerInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFollowerInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFollowerInfoResponse) GoString() string {
	return s.String()
}

func (s *GetFollowerInfoResponse) SetHeaders(v map[string]*string) *GetFollowerInfoResponse {
	s.Headers = v
	return s
}

func (s *GetFollowerInfoResponse) SetStatusCode(v int32) *GetFollowerInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFollowerInfoResponse) SetBody(v *GetFollowerInfoResponseBody) *GetFollowerInfoResponse {
	s.Body = v
	return s
}

type GetPictureDownloadUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPictureDownloadUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPictureDownloadUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetPictureDownloadUrlHeaders) SetCommonHeaders(v map[string]*string) *GetPictureDownloadUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPictureDownloadUrlHeaders) SetXAcsDingtalkAccessToken(v string) *GetPictureDownloadUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPictureDownloadUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// UgIzXXo+Rp3Rqcts7BE08y49Jr6iu6xW4iQ
	DownloadCode *string `json:"downloadCode,omitempty" xml:"downloadCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sid001234
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s GetPictureDownloadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPictureDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetPictureDownloadUrlRequest) SetDownloadCode(v string) *GetPictureDownloadUrlRequest {
	s.DownloadCode = &v
	return s
}

func (s *GetPictureDownloadUrlRequest) SetSessionId(v string) *GetPictureDownloadUrlRequest {
	s.SessionId = &v
	return s
}

type GetPictureDownloadUrlResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	Result *GetPictureDownloadUrlResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetPictureDownloadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPictureDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetPictureDownloadUrlResponseBody) SetRequestId(v string) *GetPictureDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPictureDownloadUrlResponseBody) SetResult(v *GetPictureDownloadUrlResponseBodyResult) *GetPictureDownloadUrlResponseBody {
	s.Result = v
	return s
}

type GetPictureDownloadUrlResponseBodyResult struct {
	// This parameter is required.
	//
	// example:
	//
	// FOLLOWED
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetPictureDownloadUrlResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetPictureDownloadUrlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetPictureDownloadUrlResponseBodyResult) SetUrl(v string) *GetPictureDownloadUrlResponseBodyResult {
	s.Url = &v
	return s
}

type GetPictureDownloadUrlResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPictureDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPictureDownloadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPictureDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetPictureDownloadUrlResponse) SetHeaders(v map[string]*string) *GetPictureDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetPictureDownloadUrlResponse) SetStatusCode(v int32) *GetPictureDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPictureDownloadUrlResponse) SetBody(v *GetPictureDownloadUrlResponseBody) *GetPictureDownloadUrlResponse {
	s.Body = v
	return s
}

type GetUserFollowStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserFollowStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserFollowStatusHeaders) GoString() string {
	return s.String()
}

func (s *GetUserFollowStatusHeaders) SetCommonHeaders(v map[string]*string) *GetUserFollowStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserFollowStatusHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserFollowStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserFollowStatusRequest struct {
	// example:
	//
	// ding1234
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// UgIzXXo+Rp3Rqcts7BE08y49Jr6iu6xW4iQ
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// example:
	//
	// Rp3Rqcts7BE08y49Jr6iu6xW4iQ
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserFollowStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserFollowStatusRequest) GoString() string {
	return s.String()
}

func (s *GetUserFollowStatusRequest) SetAccountId(v string) *GetUserFollowStatusRequest {
	s.AccountId = &v
	return s
}

func (s *GetUserFollowStatusRequest) SetUnionId(v string) *GetUserFollowStatusRequest {
	s.UnionId = &v
	return s
}

func (s *GetUserFollowStatusRequest) SetUserId(v string) *GetUserFollowStatusRequest {
	s.UserId = &v
	return s
}

type GetUserFollowStatusResponseBody struct {
	// This parameter is required.
	Result *GetUserFollowStatusResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetUserFollowStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserFollowStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserFollowStatusResponseBody) SetResult(v *GetUserFollowStatusResponseBodyResult) *GetUserFollowStatusResponseBody {
	s.Result = v
	return s
}

type GetUserFollowStatusResponseBodyResult struct {
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetUserFollowStatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetUserFollowStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetUserFollowStatusResponseBodyResult) SetStatus(v string) *GetUserFollowStatusResponseBodyResult {
	s.Status = &v
	return s
}

type GetUserFollowStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserFollowStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserFollowStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserFollowStatusResponse) GoString() string {
	return s.String()
}

func (s *GetUserFollowStatusResponse) SetHeaders(v map[string]*string) *GetUserFollowStatusResponse {
	s.Headers = v
	return s
}

func (s *GetUserFollowStatusResponse) SetStatusCode(v int32) *GetUserFollowStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserFollowStatusResponse) SetBody(v *GetUserFollowStatusResponseBody) *GetUserFollowStatusResponse {
	s.Body = v
	return s
}

type ListAccountHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListAccountHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListAccountHeaders) GoString() string {
	return s.String()
}

func (s *ListAccountHeaders) SetCommonHeaders(v map[string]*string) *ListAccountHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAccountHeaders) SetXAcsDingtalkAccessToken(v string) *ListAccountHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListAccountResponseBody struct {
	Result []*ListAccountResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccountResponseBody) SetResult(v []*ListAccountResponseBodyResult) *ListAccountResponseBody {
	s.Result = v
	return s
}

type ListAccountResponseBodyResult struct {
	AccountId   *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
}

func (s ListAccountResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListAccountResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAccountResponseBodyResult) SetAccountId(v string) *ListAccountResponseBodyResult {
	s.AccountId = &v
	return s
}

func (s *ListAccountResponseBodyResult) SetAccountName(v string) *ListAccountResponseBodyResult {
	s.AccountName = &v
	return s
}

type ListAccountResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccountResponse) GoString() string {
	return s.String()
}

func (s *ListAccountResponse) SetHeaders(v map[string]*string) *ListAccountResponse {
	s.Headers = v
	return s
}

func (s *ListAccountResponse) SetStatusCode(v int32) *ListAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccountResponse) SetBody(v *ListAccountResponseBody) *ListAccountResponse {
	s.Body = v
	return s
}

type ListAccountInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListAccountInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListAccountInfoHeaders) GoString() string {
	return s.String()
}

func (s *ListAccountInfoHeaders) SetCommonHeaders(v map[string]*string) *ListAccountInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAccountInfoHeaders) SetXAcsDingtalkAccessToken(v string) *ListAccountInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListAccountInfoResponseBody struct {
	Result []*ListAccountInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListAccountInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccountInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccountInfoResponseBody) SetResult(v []*ListAccountInfoResponseBodyResult) *ListAccountInfoResponseBody {
	s.Result = v
	return s
}

type ListAccountInfoResponseBodyResult struct {
	AccountId   *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
}

func (s ListAccountInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListAccountInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAccountInfoResponseBodyResult) SetAccountId(v string) *ListAccountInfoResponseBodyResult {
	s.AccountId = &v
	return s
}

func (s *ListAccountInfoResponseBodyResult) SetAccountName(v string) *ListAccountInfoResponseBodyResult {
	s.AccountName = &v
	return s
}

type ListAccountInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccountInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccountInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccountInfoResponse) GoString() string {
	return s.String()
}

func (s *ListAccountInfoResponse) SetHeaders(v map[string]*string) *ListAccountInfoResponse {
	s.Headers = v
	return s
}

func (s *ListAccountInfoResponse) SetStatusCode(v int32) *ListAccountInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccountInfoResponse) SetBody(v *ListAccountInfoResponseBody) *ListAccountInfoResponse {
	s.Body = v
	return s
}

type ListFollowerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListFollowerHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListFollowerHeaders) GoString() string {
	return s.String()
}

func (s *ListFollowerHeaders) SetCommonHeaders(v map[string]*string) *ListFollowerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListFollowerHeaders) SetXAcsDingtalkAccessToken(v string) *ListFollowerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListFollowerRequest struct {
	// example:
	//
	// ding1234
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// Rp3Rqcts7BE08y49Jr6iu6xW4iQ
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListFollowerRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFollowerRequest) GoString() string {
	return s.String()
}

func (s *ListFollowerRequest) SetAccountId(v string) *ListFollowerRequest {
	s.AccountId = &v
	return s
}

func (s *ListFollowerRequest) SetMaxResults(v int32) *ListFollowerRequest {
	s.MaxResults = &v
	return s
}

func (s *ListFollowerRequest) SetNextToken(v string) *ListFollowerRequest {
	s.NextToken = &v
	return s
}

type ListFollowerResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	Result *ListFollowerResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ListFollowerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFollowerResponseBody) GoString() string {
	return s.String()
}

func (s *ListFollowerResponseBody) SetRequestId(v string) *ListFollowerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFollowerResponseBody) SetResult(v *ListFollowerResponseBodyResult) *ListFollowerResponseBody {
	s.Result = v
	return s
}

type ListFollowerResponseBodyResult struct {
	NextToken *string                                   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	UserList  []*ListFollowerResponseBodyResultUserList `json:"userList,omitempty" xml:"userList,omitempty" type:"Repeated"`
}

func (s ListFollowerResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListFollowerResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListFollowerResponseBodyResult) SetNextToken(v string) *ListFollowerResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *ListFollowerResponseBodyResult) SetUserList(v []*ListFollowerResponseBodyResultUserList) *ListFollowerResponseBodyResult {
	s.UserList = v
	return s
}

type ListFollowerResponseBodyResultUserList struct {
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Timestamp *int64  `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListFollowerResponseBodyResultUserList) String() string {
	return tea.Prettify(s)
}

func (s ListFollowerResponseBodyResultUserList) GoString() string {
	return s.String()
}

func (s *ListFollowerResponseBodyResultUserList) SetName(v string) *ListFollowerResponseBodyResultUserList {
	s.Name = &v
	return s
}

func (s *ListFollowerResponseBodyResultUserList) SetTimestamp(v int64) *ListFollowerResponseBodyResultUserList {
	s.Timestamp = &v
	return s
}

func (s *ListFollowerResponseBodyResultUserList) SetUserId(v string) *ListFollowerResponseBodyResultUserList {
	s.UserId = &v
	return s
}

type ListFollowerResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFollowerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFollowerResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFollowerResponse) GoString() string {
	return s.String()
}

func (s *ListFollowerResponse) SetHeaders(v map[string]*string) *ListFollowerResponse {
	s.Headers = v
	return s
}

func (s *ListFollowerResponse) SetStatusCode(v int32) *ListFollowerResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFollowerResponse) SetBody(v *ListFollowerResponseBody) *ListFollowerResponse {
	s.Body = v
	return s
}

type QueryUserFollowStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserFollowStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserFollowStatusHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserFollowStatusHeaders) SetCommonHeaders(v map[string]*string) *QueryUserFollowStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserFollowStatusHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserFollowStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserFollowStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ding1234
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// UgIzXXo+Rp3Rqcts7BE08y49Jr6iu6xW4iQ
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryUserFollowStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserFollowStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryUserFollowStatusRequest) SetAccountId(v string) *QueryUserFollowStatusRequest {
	s.AccountId = &v
	return s
}

func (s *QueryUserFollowStatusRequest) SetUnionId(v string) *QueryUserFollowStatusRequest {
	s.UnionId = &v
	return s
}

type QueryUserFollowStatusResponseBody struct {
	// This parameter is required.
	Result *QueryUserFollowStatusResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryUserFollowStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserFollowStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserFollowStatusResponseBody) SetResult(v *QueryUserFollowStatusResponseBodyResult) *QueryUserFollowStatusResponseBody {
	s.Result = v
	return s
}

type QueryUserFollowStatusResponseBodyResult struct {
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryUserFollowStatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryUserFollowStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryUserFollowStatusResponseBodyResult) SetStatus(v string) *QueryUserFollowStatusResponseBodyResult {
	s.Status = &v
	return s
}

type QueryUserFollowStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserFollowStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserFollowStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserFollowStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryUserFollowStatusResponse) SetHeaders(v map[string]*string) *QueryUserFollowStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryUserFollowStatusResponse) SetStatusCode(v int32) *QueryUserFollowStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserFollowStatusResponse) SetBody(v *QueryUserFollowStatusResponseBody) *QueryUserFollowStatusResponse {
	s.Body = v
	return s
}

type SendAgentOTOMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendAgentOTOMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendAgentOTOMessageHeaders) GoString() string {
	return s.String()
}

func (s *SendAgentOTOMessageHeaders) SetCommonHeaders(v map[string]*string) *SendAgentOTOMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendAgentOTOMessageHeaders) SetXAcsDingtalkAccessToken(v string) *SendAgentOTOMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendAgentOTOMessageRequest struct {
	// This parameter is required.
	Detail *SendAgentOTOMessageRequestDetail `json:"detail,omitempty" xml:"detail,omitempty" type:"Struct"`
}

func (s SendAgentOTOMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendAgentOTOMessageRequest) GoString() string {
	return s.String()
}

func (s *SendAgentOTOMessageRequest) SetDetail(v *SendAgentOTOMessageRequestDetail) *SendAgentOTOMessageRequest {
	s.Detail = v
	return s
}

type SendAgentOTOMessageRequestDetail struct {
	// This parameter is required.
	MessageBody *SendAgentOTOMessageRequestDetailMessageBody `json:"messageBody,omitempty" xml:"messageBody,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// text
	MsgType *string `json:"msgType,omitempty" xml:"msgType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sid002b6dbb4f963e93e0
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user0001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234-5678-000
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s SendAgentOTOMessageRequestDetail) String() string {
	return tea.Prettify(s)
}

func (s SendAgentOTOMessageRequestDetail) GoString() string {
	return s.String()
}

func (s *SendAgentOTOMessageRequestDetail) SetMessageBody(v *SendAgentOTOMessageRequestDetailMessageBody) *SendAgentOTOMessageRequestDetail {
	s.MessageBody = v
	return s
}

func (s *SendAgentOTOMessageRequestDetail) SetMsgType(v string) *SendAgentOTOMessageRequestDetail {
	s.MsgType = &v
	return s
}

func (s *SendAgentOTOMessageRequestDetail) SetSessionId(v string) *SendAgentOTOMessageRequestDetail {
	s.SessionId = &v
	return s
}

func (s *SendAgentOTOMessageRequestDetail) SetUserId(v string) *SendAgentOTOMessageRequestDetail {
	s.UserId = &v
	return s
}

func (s *SendAgentOTOMessageRequestDetail) SetUuid(v string) *SendAgentOTOMessageRequestDetail {
	s.Uuid = &v
	return s
}

type SendAgentOTOMessageRequestDetailMessageBody struct {
	ActionCard         *SendAgentOTOMessageRequestDetailMessageBodyActionCard         `json:"actionCard,omitempty" xml:"actionCard,omitempty" type:"Struct"`
	Image              *SendAgentOTOMessageRequestDetailMessageBodyImage              `json:"image,omitempty" xml:"image,omitempty" type:"Struct"`
	InteractiveMessage *SendAgentOTOMessageRequestDetailMessageBodyInteractiveMessage `json:"interactiveMessage,omitempty" xml:"interactiveMessage,omitempty" type:"Struct"`
	Link               *SendAgentOTOMessageRequestDetailMessageBodyLink               `json:"link,omitempty" xml:"link,omitempty" type:"Struct"`
	Markdown           *SendAgentOTOMessageRequestDetailMessageBodyMarkdown           `json:"markdown,omitempty" xml:"markdown,omitempty" type:"Struct"`
	Text               *SendAgentOTOMessageRequestDetailMessageBodyText               `json:"text,omitempty" xml:"text,omitempty" type:"Struct"`
}

func (s SendAgentOTOMessageRequestDetailMessageBody) String() string {
	return tea.Prettify(s)
}

func (s SendAgentOTOMessageRequestDetailMessageBody) GoString() string {
	return s.String()
}

func (s *SendAgentOTOMessageRequestDetailMessageBody) SetActionCard(v *SendAgentOTOMessageRequestDetailMessageBodyActionCard) *SendAgentOTOMessageRequestDetailMessageBody {
	s.ActionCard = v
	return s
}

func (s *SendAgentOTOMessageRequestDetailMessageBody) SetImage(v *SendAgentOTOMessageRequestDetailMessageBodyImage) *SendAgentOTOMessageRequestDetailMessageBody {
	s.Image = v
	return s
}

func (s *SendAgentOTOMessageRequestDetailMessageBody) SetInteractiveMessage(v *SendAgentOTOMessageRequestDetailMessageBodyInteractiveMessage) *SendAgentOTOMessageRequestDetailMessageBody {
	s.InteractiveMessage = v
	return s
}

func (s *SendAgentOTOMessageRequestDetailMessageBody) SetLink(v *SendAgentOTOMessageRequestDetailMessageBodyLink) *SendAgentOTOMessageRequestDetailMessageBody {
	s.Link = v
	return s
}

func (s *SendAgentOTOMessageRequestDetailMessageBody) SetMarkdown(v *SendAgentOTOMessageRequestDetailMessageBodyMarkdown) *SendAgentOTOMessageRequestDetailMessageBody {
	s.Markdown = v
	return s
}

func (s *SendAgentOTOMessageRequestDetailMessageBody) SetText(v *SendAgentOTOMessageRequestDetailMessageBodyText) *SendAgentOTOMessageRequestDetailMessageBody {
	s.Text = v
	return s
}

type SendAgentOTOMessageRequestDetailMessageBodyActionCard struct {
	ButtonList []*SendAgentOTOMessageRequestDetailMessageBodyActionCardButtonList `json:"buttonList,omitempty" xml:"buttonList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	ButtonOrientation *string `json:"buttonOrientation,omitempty" xml:"buttonOrientation,omitempty"`
	// example:
	//
	// **提示信息**
	Markdown *string `json:"markdown,omitempty" xml:"markdown,omitempty"`
	// example:
	//
	// 查看详情
	SingleTitle *string `json:"singleTitle,omitempty" xml:"singleTitle,omitempty"`
	// example:
	//
	// https://www.yourdomain.com
	SingleUrl *string `json:"singleUrl,omitempty" xml:"singleUrl,omitempty"`
	// example:
	//
	// 欢迎使用
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SendAgentOTOMessageRequestDetailMessageBodyActionCard) String() string {
	return tea.Prettify(s)
}

func (s SendAgentOTOMessageRequestDetailMessageBodyActionCard) GoString() string {
	return s.String()
}

func (s *SendAgentOTOMessageRequestDetailMessageBodyActionCard) SetButtonList(v []*SendAgentOTOMessageRequestDetailMessageBodyActionCardButtonList) *SendAgentOTOMessageRequestDetailMessageBodyActionCard {
	s.ButtonList = v
	return s
}

func (s *SendAgentOTOMessageRequestDetailMessageBodyActionCard) SetButtonOrientation(v string) *SendAgentOTOMessageRequestDetailMessageBodyActionCard {
	s.ButtonOrientation = &v
	return s
}

func (s *SendAgentOTOMessageRequestDetailMessageBodyActionCard) SetMarkdown(v string) *SendAgentOTOMessageRequestDetailMessageBodyActionCard {
	s.Markdown = &v
	return s
}

func (s *SendAgentOTOMessageRequestDetailMessageBodyActionCard) SetSingleTitle(v string) *SendAgentOTOMessageRequestDetailMessageBodyActionCard {
	s.SingleTitle = &v
	return s
}

func (s *SendAgentOTOMessageRequestDetailMessageBodyActionCard) SetSingleUrl(v string) *SendAgentOTOMessageRequestDetailMessageBodyActionCard {
	s.SingleUrl = &v
	return s
}

func (s *SendAgentOTOMessageRequestDetailMessageBodyActionCard) SetTitle(v string) *SendAgentOTOMessageRequestDetailMessageBodyActionCard {
	s.Title = &v
	return s
}

type SendAgentOTOMessageRequestDetailMessageBodyActionCardButtonList struct {
	// example:
	//
	// https://www.dingtalk.com/
	ActionUrl *string `json:"actionUrl,omitempty" xml:"actionUrl,omitempty"`
	// example:
	//
	// 查看详情
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SendAgentOTOMessageRequestDetailMessageBodyActionCardButtonList) String() string {
	return tea.Prettify(s)
}

func (s SendAgentOTOMessageRequestDetailMessageBodyActionCardButtonList) GoString() string {
	return s.String()
}

func (s *SendAgentOTOMessageRequestDetailMessageBodyActionCardButtonList) SetActionUrl(v string) *SendAgentOTOMessageRequestDetailMessageBodyActionCardButtonList {
	s.ActionUrl = &v
	return s
}

func (s *SendAgentOTOMessageRequestDetailMessageBodyActionCardButtonList) SetTitle(v string) *SendAgentOTOMessageRequestDetailMessageBodyActionCardButtonList {
	s.Title = &v
	return s
}

type SendAgentOTOMessageRequestDetailMessageBodyImage struct {
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
}

func (s SendAgentOTOMessageRequestDetailMessageBodyImage) String() string {
	return tea.Prettify(s)
}

func (s SendAgentOTOMessageRequestDetailMessageBodyImage) GoString() string {
	return s.String()
}

func (s *SendAgentOTOMessageRequestDetailMessageBodyImage) SetMediaId(v string) *SendAgentOTOMessageRequestDetailMessageBodyImage {
	s.MediaId = &v
	return s
}

type SendAgentOTOMessageRequestDetailMessageBodyInteractiveMessage struct {
	CallbackUrl    *string `json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`
	CardBizId      *string `json:"cardBizId,omitempty" xml:"cardBizId,omitempty"`
	CardData       *string `json:"cardData,omitempty" xml:"cardData,omitempty"`
	CardTemplateId *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
}

func (s SendAgentOTOMessageRequestDetailMessageBodyInteractiveMessage) String() string {
	return tea.Prettify(s)
}

func (s SendAgentOTOMessageRequestDetailMessageBodyInteractiveMessage) GoString() string {
	return s.String()
}

func (s *SendAgentOTOMessageRequestDetailMessageBodyInteractiveMessage) SetCallbackUrl(v string) *SendAgentOTOMessageRequestDetailMessageBodyInteractiveMessage {
	s.CallbackUrl = &v
	return s
}

func (s *SendAgentOTOMessageRequestDetailMessageBodyInteractiveMessage) SetCardBizId(v string) *SendAgentOTOMessageRequestDetailMessageBodyInteractiveMessage {
	s.CardBizId = &v
	return s
}

func (s *SendAgentOTOMessageRequestDetailMessageBodyInteractiveMessage) SetCardData(v string) *SendAgentOTOMessageRequestDetailMessageBodyInteractiveMessage {
	s.CardData = &v
	return s
}

func (s *SendAgentOTOMessageRequestDetailMessageBodyInteractiveMessage) SetCardTemplateId(v string) *SendAgentOTOMessageRequestDetailMessageBodyInteractiveMessage {
	s.CardTemplateId = &v
	return s
}

type SendAgentOTOMessageRequestDetailMessageBodyLink struct {
	// example:
	//
	// https://www.yourdomain.com
	MessageUrl *string `json:"messageUrl,omitempty" xml:"messageUrl,omitempty"`
	// example:
	//
	// @1234-456
	PicUrl *string `json:"picUrl,omitempty" xml:"picUrl,omitempty"`
	// example:
	//
	// 欢迎使用
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// example:
	//
	// 点击查看
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SendAgentOTOMessageRequestDetailMessageBodyLink) String() string {
	return tea.Prettify(s)
}

func (s SendAgentOTOMessageRequestDetailMessageBodyLink) GoString() string {
	return s.String()
}

func (s *SendAgentOTOMessageRequestDetailMessageBodyLink) SetMessageUrl(v string) *SendAgentOTOMessageRequestDetailMessageBodyLink {
	s.MessageUrl = &v
	return s
}

func (s *SendAgentOTOMessageRequestDetailMessageBodyLink) SetPicUrl(v string) *SendAgentOTOMessageRequestDetailMessageBodyLink {
	s.PicUrl = &v
	return s
}

func (s *SendAgentOTOMessageRequestDetailMessageBodyLink) SetText(v string) *SendAgentOTOMessageRequestDetailMessageBodyLink {
	s.Text = &v
	return s
}

func (s *SendAgentOTOMessageRequestDetailMessageBodyLink) SetTitle(v string) *SendAgentOTOMessageRequestDetailMessageBodyLink {
	s.Title = &v
	return s
}

type SendAgentOTOMessageRequestDetailMessageBodyMarkdown struct {
	// example:
	//
	// 欢迎使用。
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// example:
	//
	// 欢迎使用。
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SendAgentOTOMessageRequestDetailMessageBodyMarkdown) String() string {
	return tea.Prettify(s)
}

func (s SendAgentOTOMessageRequestDetailMessageBodyMarkdown) GoString() string {
	return s.String()
}

func (s *SendAgentOTOMessageRequestDetailMessageBodyMarkdown) SetText(v string) *SendAgentOTOMessageRequestDetailMessageBodyMarkdown {
	s.Text = &v
	return s
}

func (s *SendAgentOTOMessageRequestDetailMessageBodyMarkdown) SetTitle(v string) *SendAgentOTOMessageRequestDetailMessageBodyMarkdown {
	s.Title = &v
	return s
}

type SendAgentOTOMessageRequestDetailMessageBodyText struct {
	// example:
	//
	// 你好，欢迎使用服务窗。
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s SendAgentOTOMessageRequestDetailMessageBodyText) String() string {
	return tea.Prettify(s)
}

func (s SendAgentOTOMessageRequestDetailMessageBodyText) GoString() string {
	return s.String()
}

func (s *SendAgentOTOMessageRequestDetailMessageBodyText) SetContent(v string) *SendAgentOTOMessageRequestDetailMessageBodyText {
	s.Content = &v
	return s
}

type SendAgentOTOMessageResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	Result *SendAgentOTOMessageResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s SendAgentOTOMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendAgentOTOMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendAgentOTOMessageResponseBody) SetRequestId(v string) *SendAgentOTOMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendAgentOTOMessageResponseBody) SetResult(v *SendAgentOTOMessageResponseBodyResult) *SendAgentOTOMessageResponseBody {
	s.Result = v
	return s
}

type SendAgentOTOMessageResponseBodyResult struct {
	// This parameter is required.
	OpenPushId *string `json:"openPushId,omitempty" xml:"openPushId,omitempty"`
}

func (s SendAgentOTOMessageResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SendAgentOTOMessageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SendAgentOTOMessageResponseBodyResult) SetOpenPushId(v string) *SendAgentOTOMessageResponseBodyResult {
	s.OpenPushId = &v
	return s
}

type SendAgentOTOMessageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendAgentOTOMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendAgentOTOMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendAgentOTOMessageResponse) GoString() string {
	return s.String()
}

func (s *SendAgentOTOMessageResponse) SetHeaders(v map[string]*string) *SendAgentOTOMessageResponse {
	s.Headers = v
	return s
}

func (s *SendAgentOTOMessageResponse) SetStatusCode(v int32) *SendAgentOTOMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendAgentOTOMessageResponse) SetBody(v *SendAgentOTOMessageResponseBody) *SendAgentOTOMessageResponse {
	s.Body = v
	return s
}

type SendInteractiveOTOMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendInteractiveOTOMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendInteractiveOTOMessageHeaders) GoString() string {
	return s.String()
}

func (s *SendInteractiveOTOMessageHeaders) SetCommonHeaders(v map[string]*string) *SendInteractiveOTOMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendInteractiveOTOMessageHeaders) SetXAcsDingtalkAccessToken(v string) *SendInteractiveOTOMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendInteractiveOTOMessageRequest struct {
	// This parameter is required.
	Detail *SendInteractiveOTOMessageRequestDetail `json:"detail,omitempty" xml:"detail,omitempty" type:"Struct"`
}

func (s SendInteractiveOTOMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendInteractiveOTOMessageRequest) GoString() string {
	return s.String()
}

func (s *SendInteractiveOTOMessageRequest) SetDetail(v *SendInteractiveOTOMessageRequestDetail) *SendInteractiveOTOMessageRequest {
	s.Detail = v
	return s
}

type SendInteractiveOTOMessageRequestDetail struct {
	// example:
	//
	// https://www.youurl.com/callback/card
	CallbackUrl *string `json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// service-card-20220824-001
	CardBizId *string `json:"cardBizId,omitempty" xml:"cardBizId,omitempty"`
	// This parameter is required.
	CardData *string `json:"cardData,omitempty" xml:"cardData,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3erkfi-42b0-4c83-bc56-ffhssde43
	CardTemplateId *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user0001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// {"user001":""}
	UserIdPrivateDataMap *string `json:"userIdPrivateDataMap,omitempty" xml:"userIdPrivateDataMap,omitempty"`
}

func (s SendInteractiveOTOMessageRequestDetail) String() string {
	return tea.Prettify(s)
}

func (s SendInteractiveOTOMessageRequestDetail) GoString() string {
	return s.String()
}

func (s *SendInteractiveOTOMessageRequestDetail) SetCallbackUrl(v string) *SendInteractiveOTOMessageRequestDetail {
	s.CallbackUrl = &v
	return s
}

func (s *SendInteractiveOTOMessageRequestDetail) SetCardBizId(v string) *SendInteractiveOTOMessageRequestDetail {
	s.CardBizId = &v
	return s
}

func (s *SendInteractiveOTOMessageRequestDetail) SetCardData(v string) *SendInteractiveOTOMessageRequestDetail {
	s.CardData = &v
	return s
}

func (s *SendInteractiveOTOMessageRequestDetail) SetCardTemplateId(v string) *SendInteractiveOTOMessageRequestDetail {
	s.CardTemplateId = &v
	return s
}

func (s *SendInteractiveOTOMessageRequestDetail) SetUserId(v string) *SendInteractiveOTOMessageRequestDetail {
	s.UserId = &v
	return s
}

func (s *SendInteractiveOTOMessageRequestDetail) SetUserIdPrivateDataMap(v string) *SendInteractiveOTOMessageRequestDetail {
	s.UserIdPrivateDataMap = &v
	return s
}

type SendInteractiveOTOMessageResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	Result *SendInteractiveOTOMessageResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s SendInteractiveOTOMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendInteractiveOTOMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendInteractiveOTOMessageResponseBody) SetRequestId(v string) *SendInteractiveOTOMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendInteractiveOTOMessageResponseBody) SetResult(v *SendInteractiveOTOMessageResponseBodyResult) *SendInteractiveOTOMessageResponseBody {
	s.Result = v
	return s
}

type SendInteractiveOTOMessageResponseBodyResult struct {
	// This parameter is required.
	OpenPushId *string `json:"openPushId,omitempty" xml:"openPushId,omitempty"`
}

func (s SendInteractiveOTOMessageResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SendInteractiveOTOMessageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SendInteractiveOTOMessageResponseBodyResult) SetOpenPushId(v string) *SendInteractiveOTOMessageResponseBodyResult {
	s.OpenPushId = &v
	return s
}

type SendInteractiveOTOMessageResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendInteractiveOTOMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendInteractiveOTOMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendInteractiveOTOMessageResponse) GoString() string {
	return s.String()
}

func (s *SendInteractiveOTOMessageResponse) SetHeaders(v map[string]*string) *SendInteractiveOTOMessageResponse {
	s.Headers = v
	return s
}

func (s *SendInteractiveOTOMessageResponse) SetStatusCode(v int32) *SendInteractiveOTOMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendInteractiveOTOMessageResponse) SetBody(v *SendInteractiveOTOMessageResponseBody) *SendInteractiveOTOMessageResponse {
	s.Body = v
	return s
}

type SendTopBoxInteractiveOTOMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendTopBoxInteractiveOTOMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendTopBoxInteractiveOTOMessageHeaders) GoString() string {
	return s.String()
}

func (s *SendTopBoxInteractiveOTOMessageHeaders) SetCommonHeaders(v map[string]*string) *SendTopBoxInteractiveOTOMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendTopBoxInteractiveOTOMessageHeaders) SetXAcsDingtalkAccessToken(v string) *SendTopBoxInteractiveOTOMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendTopBoxInteractiveOTOMessageRequest struct {
	// This parameter is required.
	Detail *SendTopBoxInteractiveOTOMessageRequestDetail `json:"detail,omitempty" xml:"detail,omitempty" type:"Struct"`
}

func (s SendTopBoxInteractiveOTOMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendTopBoxInteractiveOTOMessageRequest) GoString() string {
	return s.String()
}

func (s *SendTopBoxInteractiveOTOMessageRequest) SetDetail(v *SendTopBoxInteractiveOTOMessageRequestDetail) *SendTopBoxInteractiveOTOMessageRequest {
	s.Detail = v
	return s
}

type SendTopBoxInteractiveOTOMessageRequestDetail struct {
	// example:
	//
	// https://www.youurl.com/callback/card
	CallbackUrl *string `json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// service-card-20220824-001
	CardBizId *string `json:"cardBizId,omitempty" xml:"cardBizId,omitempty"`
	// This parameter is required.
	CardData *SendTopBoxInteractiveOTOMessageRequestDetailCardData `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 3erkfi-42b0-4c83-bc56-ffhssde43
	CardTemplateId *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	// This parameter is required.
	ExpiredTime *int64 `json:"expiredTime,omitempty" xml:"expiredTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user001
	UserId               *string                                     `json:"userId,omitempty" xml:"userId,omitempty"`
	UserIdPrivateDataMap map[string]*DetailUserIdPrivateDataMapValue `json:"userIdPrivateDataMap,omitempty" xml:"userIdPrivateDataMap,omitempty"`
}

func (s SendTopBoxInteractiveOTOMessageRequestDetail) String() string {
	return tea.Prettify(s)
}

func (s SendTopBoxInteractiveOTOMessageRequestDetail) GoString() string {
	return s.String()
}

func (s *SendTopBoxInteractiveOTOMessageRequestDetail) SetCallbackUrl(v string) *SendTopBoxInteractiveOTOMessageRequestDetail {
	s.CallbackUrl = &v
	return s
}

func (s *SendTopBoxInteractiveOTOMessageRequestDetail) SetCardBizId(v string) *SendTopBoxInteractiveOTOMessageRequestDetail {
	s.CardBizId = &v
	return s
}

func (s *SendTopBoxInteractiveOTOMessageRequestDetail) SetCardData(v *SendTopBoxInteractiveOTOMessageRequestDetailCardData) *SendTopBoxInteractiveOTOMessageRequestDetail {
	s.CardData = v
	return s
}

func (s *SendTopBoxInteractiveOTOMessageRequestDetail) SetCardTemplateId(v string) *SendTopBoxInteractiveOTOMessageRequestDetail {
	s.CardTemplateId = &v
	return s
}

func (s *SendTopBoxInteractiveOTOMessageRequestDetail) SetExpiredTime(v int64) *SendTopBoxInteractiveOTOMessageRequestDetail {
	s.ExpiredTime = &v
	return s
}

func (s *SendTopBoxInteractiveOTOMessageRequestDetail) SetUserId(v string) *SendTopBoxInteractiveOTOMessageRequestDetail {
	s.UserId = &v
	return s
}

func (s *SendTopBoxInteractiveOTOMessageRequestDetail) SetUserIdPrivateDataMap(v map[string]*DetailUserIdPrivateDataMapValue) *SendTopBoxInteractiveOTOMessageRequestDetail {
	s.UserIdPrivateDataMap = v
	return s
}

type SendTopBoxInteractiveOTOMessageRequestDetailCardData struct {
	CardMediaIdParamMap map[string]interface{} `json:"cardMediaIdParamMap,omitempty" xml:"cardMediaIdParamMap,omitempty"`
	CardParamMap        map[string]interface{} `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
}

func (s SendTopBoxInteractiveOTOMessageRequestDetailCardData) String() string {
	return tea.Prettify(s)
}

func (s SendTopBoxInteractiveOTOMessageRequestDetailCardData) GoString() string {
	return s.String()
}

func (s *SendTopBoxInteractiveOTOMessageRequestDetailCardData) SetCardMediaIdParamMap(v map[string]interface{}) *SendTopBoxInteractiveOTOMessageRequestDetailCardData {
	s.CardMediaIdParamMap = v
	return s
}

func (s *SendTopBoxInteractiveOTOMessageRequestDetailCardData) SetCardParamMap(v map[string]interface{}) *SendTopBoxInteractiveOTOMessageRequestDetailCardData {
	s.CardParamMap = v
	return s
}

type SendTopBoxInteractiveOTOMessageResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *bool   `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SendTopBoxInteractiveOTOMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendTopBoxInteractiveOTOMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendTopBoxInteractiveOTOMessageResponseBody) SetRequestId(v string) *SendTopBoxInteractiveOTOMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendTopBoxInteractiveOTOMessageResponseBody) SetResult(v bool) *SendTopBoxInteractiveOTOMessageResponseBody {
	s.Result = &v
	return s
}

type SendTopBoxInteractiveOTOMessageResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendTopBoxInteractiveOTOMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendTopBoxInteractiveOTOMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendTopBoxInteractiveOTOMessageResponse) GoString() string {
	return s.String()
}

func (s *SendTopBoxInteractiveOTOMessageResponse) SetHeaders(v map[string]*string) *SendTopBoxInteractiveOTOMessageResponse {
	s.Headers = v
	return s
}

func (s *SendTopBoxInteractiveOTOMessageResponse) SetStatusCode(v int32) *SendTopBoxInteractiveOTOMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendTopBoxInteractiveOTOMessageResponse) SetBody(v *SendTopBoxInteractiveOTOMessageResponseBody) *SendTopBoxInteractiveOTOMessageResponse {
	s.Body = v
	return s
}

type UpdateInteractiveOTOMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateInteractiveOTOMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateInteractiveOTOMessageHeaders) GoString() string {
	return s.String()
}

func (s *UpdateInteractiveOTOMessageHeaders) SetCommonHeaders(v map[string]*string) *UpdateInteractiveOTOMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateInteractiveOTOMessageHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateInteractiveOTOMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateInteractiveOTOMessageRequest struct {
	// This parameter is required.
	Detail *UpdateInteractiveOTOMessageRequestDetail `json:"detail,omitempty" xml:"detail,omitempty" type:"Struct"`
}

func (s UpdateInteractiveOTOMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInteractiveOTOMessageRequest) GoString() string {
	return s.String()
}

func (s *UpdateInteractiveOTOMessageRequest) SetDetail(v *UpdateInteractiveOTOMessageRequestDetail) *UpdateInteractiveOTOMessageRequest {
	s.Detail = v
	return s
}

type UpdateInteractiveOTOMessageRequestDetail struct {
	// This parameter is required.
	//
	// example:
	//
	// service-card-202208240001
	CardBizId *string `json:"cardBizId,omitempty" xml:"cardBizId,omitempty"`
	// example:
	//
	// {"like":1}
	CardData      *string                                                `json:"cardData,omitempty" xml:"cardData,omitempty"`
	UpdateOptions *UpdateInteractiveOTOMessageRequestDetailUpdateOptions `json:"updateOptions,omitempty" xml:"updateOptions,omitempty" type:"Struct"`
	// example:
	//
	// {"userI":""}
	UserIdPrivateDataMap *string `json:"userIdPrivateDataMap,omitempty" xml:"userIdPrivateDataMap,omitempty"`
}

func (s UpdateInteractiveOTOMessageRequestDetail) String() string {
	return tea.Prettify(s)
}

func (s UpdateInteractiveOTOMessageRequestDetail) GoString() string {
	return s.String()
}

func (s *UpdateInteractiveOTOMessageRequestDetail) SetCardBizId(v string) *UpdateInteractiveOTOMessageRequestDetail {
	s.CardBizId = &v
	return s
}

func (s *UpdateInteractiveOTOMessageRequestDetail) SetCardData(v string) *UpdateInteractiveOTOMessageRequestDetail {
	s.CardData = &v
	return s
}

func (s *UpdateInteractiveOTOMessageRequestDetail) SetUpdateOptions(v *UpdateInteractiveOTOMessageRequestDetailUpdateOptions) *UpdateInteractiveOTOMessageRequestDetail {
	s.UpdateOptions = v
	return s
}

func (s *UpdateInteractiveOTOMessageRequestDetail) SetUserIdPrivateDataMap(v string) *UpdateInteractiveOTOMessageRequestDetail {
	s.UserIdPrivateDataMap = &v
	return s
}

type UpdateInteractiveOTOMessageRequestDetailUpdateOptions struct {
	UpdateCardDataByKey    *bool `json:"updateCardDataByKey,omitempty" xml:"updateCardDataByKey,omitempty"`
	UpdatePrivateDataByKey *bool `json:"updatePrivateDataByKey,omitempty" xml:"updatePrivateDataByKey,omitempty"`
}

func (s UpdateInteractiveOTOMessageRequestDetailUpdateOptions) String() string {
	return tea.Prettify(s)
}

func (s UpdateInteractiveOTOMessageRequestDetailUpdateOptions) GoString() string {
	return s.String()
}

func (s *UpdateInteractiveOTOMessageRequestDetailUpdateOptions) SetUpdateCardDataByKey(v bool) *UpdateInteractiveOTOMessageRequestDetailUpdateOptions {
	s.UpdateCardDataByKey = &v
	return s
}

func (s *UpdateInteractiveOTOMessageRequestDetailUpdateOptions) SetUpdatePrivateDataByKey(v bool) *UpdateInteractiveOTOMessageRequestDetailUpdateOptions {
	s.UpdatePrivateDataByKey = &v
	return s
}

type UpdateInteractiveOTOMessageResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	Result *UpdateInteractiveOTOMessageResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateInteractiveOTOMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInteractiveOTOMessageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInteractiveOTOMessageResponseBody) SetRequestId(v string) *UpdateInteractiveOTOMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInteractiveOTOMessageResponseBody) SetResult(v *UpdateInteractiveOTOMessageResponseBodyResult) *UpdateInteractiveOTOMessageResponseBody {
	s.Result = v
	return s
}

type UpdateInteractiveOTOMessageResponseBodyResult struct {
	// This parameter is required.
	OpenPushId *string `json:"openPushId,omitempty" xml:"openPushId,omitempty"`
}

func (s UpdateInteractiveOTOMessageResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateInteractiveOTOMessageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateInteractiveOTOMessageResponseBodyResult) SetOpenPushId(v string) *UpdateInteractiveOTOMessageResponseBodyResult {
	s.OpenPushId = &v
	return s
}

type UpdateInteractiveOTOMessageResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInteractiveOTOMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInteractiveOTOMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInteractiveOTOMessageResponse) GoString() string {
	return s.String()
}

func (s *UpdateInteractiveOTOMessageResponse) SetHeaders(v map[string]*string) *UpdateInteractiveOTOMessageResponse {
	s.Headers = v
	return s
}

func (s *UpdateInteractiveOTOMessageResponse) SetStatusCode(v int32) *UpdateInteractiveOTOMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInteractiveOTOMessageResponse) SetBody(v *UpdateInteractiveOTOMessageResponseBody) *UpdateInteractiveOTOMessageResponse {
	s.Body = v
	return s
}

type UpdateShortcutsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateShortcutsHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateShortcutsHeaders) GoString() string {
	return s.String()
}

func (s *UpdateShortcutsHeaders) SetCommonHeaders(v map[string]*string) *UpdateShortcutsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateShortcutsHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateShortcutsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateShortcutsRequest struct {
	Details []*UpdateShortcutsRequestDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// sid001234
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// idzb26bxl64vqx2keyi
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateShortcutsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateShortcutsRequest) GoString() string {
	return s.String()
}

func (s *UpdateShortcutsRequest) SetDetails(v []*UpdateShortcutsRequestDetails) *UpdateShortcutsRequest {
	s.Details = v
	return s
}

func (s *UpdateShortcutsRequest) SetSessionId(v string) *UpdateShortcutsRequest {
	s.SessionId = &v
	return s
}

func (s *UpdateShortcutsRequest) SetUserId(v string) *UpdateShortcutsRequest {
	s.UserId = &v
	return s
}

type UpdateShortcutsRequestDetails struct {
	// example:
	//
	// https://dingtalk.com
	ActionUrl *string `json:"actionUrl,omitempty" xml:"actionUrl,omitempty"`
	// example:
	//
	// 033bd94b1168d7e4f0d644c3c95e35bf
	CallbackKey *string `json:"callbackKey,omitempty" xml:"callbackKey,omitempty"`
	// example:
	//
	// e73e
	IconFont *string `json:"iconFont,omitempty" xml:"iconFont,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// @lADPDg7mWPzw0i_NArzNArw
	IconMediaId *string `json:"iconMediaId,omitempty" xml:"iconMediaId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test123456
	ShortcutId *string `json:"shortcutId,omitempty" xml:"shortcutId,omitempty"`
	// example:
	//
	// @lADPDg7mWPzw0i_NArzNArw
	SlideIconMediaId *string `json:"slideIconMediaId,omitempty" xml:"slideIconMediaId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s UpdateShortcutsRequestDetails) String() string {
	return tea.Prettify(s)
}

func (s UpdateShortcutsRequestDetails) GoString() string {
	return s.String()
}

func (s *UpdateShortcutsRequestDetails) SetActionUrl(v string) *UpdateShortcutsRequestDetails {
	s.ActionUrl = &v
	return s
}

func (s *UpdateShortcutsRequestDetails) SetCallbackKey(v string) *UpdateShortcutsRequestDetails {
	s.CallbackKey = &v
	return s
}

func (s *UpdateShortcutsRequestDetails) SetIconFont(v string) *UpdateShortcutsRequestDetails {
	s.IconFont = &v
	return s
}

func (s *UpdateShortcutsRequestDetails) SetIconMediaId(v string) *UpdateShortcutsRequestDetails {
	s.IconMediaId = &v
	return s
}

func (s *UpdateShortcutsRequestDetails) SetShortcutId(v string) *UpdateShortcutsRequestDetails {
	s.ShortcutId = &v
	return s
}

func (s *UpdateShortcutsRequestDetails) SetSlideIconMediaId(v string) *UpdateShortcutsRequestDetails {
	s.SlideIconMediaId = &v
	return s
}

func (s *UpdateShortcutsRequestDetails) SetTitle(v string) *UpdateShortcutsRequestDetails {
	s.Title = &v
	return s
}

type UpdateShortcutsResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateShortcutsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateShortcutsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateShortcutsResponseBody) SetResult(v bool) *UpdateShortcutsResponseBody {
	s.Result = &v
	return s
}

type UpdateShortcutsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateShortcutsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateShortcutsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateShortcutsResponse) GoString() string {
	return s.String()
}

func (s *UpdateShortcutsResponse) SetHeaders(v map[string]*string) *UpdateShortcutsResponse {
	s.Headers = v
	return s
}

func (s *UpdateShortcutsResponse) SetStatusCode(v int32) *UpdateShortcutsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateShortcutsResponse) SetBody(v *UpdateShortcutsResponseBody) *UpdateShortcutsResponse {
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
// 发送用户授权信息申请
//
// @param request - ApplyFollowerAuthInfoRequest
//
// @param headers - ApplyFollowerAuthInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyFollowerAuthInfoResponse
func (client *Client) ApplyFollowerAuthInfoWithOptions(request *ApplyFollowerAuthInfoRequest, headers *ApplyFollowerAuthInfoHeaders, runtime *util.RuntimeOptions) (_result *ApplyFollowerAuthInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppAuthKey)) {
		body["appAuthKey"] = request.AppAuthKey
	}

	if !tea.BoolValue(util.IsUnset(request.FieldScope)) {
		body["fieldScope"] = request.FieldScope
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["sessionId"] = request.SessionId
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
		Action:      tea.String("ApplyFollowerAuthInfo"),
		Version:     tea.String("link_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/link/followers/authInfos/apply"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyFollowerAuthInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送用户授权信息申请
//
// @param request - ApplyFollowerAuthInfoRequest
//
// @return ApplyFollowerAuthInfoResponse
func (client *Client) ApplyFollowerAuthInfo(request *ApplyFollowerAuthInfoRequest) (_result *ApplyFollowerAuthInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ApplyFollowerAuthInfoHeaders{}
	_result = &ApplyFollowerAuthInfoResponse{}
	_body, _err := client.ApplyFollowerAuthInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 注册服务窗消息回调服务
//
// @param request - CallbackRegiesterRequest
//
// @param headers - CallbackRegiesterHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CallbackRegiesterResponse
func (client *Client) CallbackRegiesterWithOptions(request *CallbackRegiesterRequest, headers *CallbackRegiesterHeaders, runtime *util.RuntimeOptions) (_result *CallbackRegiesterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiSecret)) {
		body["apiSecret"] = request.ApiSecret
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackKey)) {
		body["callbackKey"] = request.CallbackKey
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackUrl)) {
		body["callbackUrl"] = request.CallbackUrl
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
		Action:      tea.String("CallbackRegiester"),
		Version:     tea.String("link_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/link/callbacks/regiester"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CallbackRegiesterResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 注册服务窗消息回调服务
//
// @param request - CallbackRegiesterRequest
//
// @return CallbackRegiesterResponse
func (client *Client) CallbackRegiester(request *CallbackRegiesterRequest) (_result *CallbackRegiesterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CallbackRegiesterHeaders{}
	_result = &CallbackRegiesterResponse{}
	_body, _err := client.CallbackRegiesterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务窗吊顶卡片关闭接口
//
// @param request - CloseTopBoxInteractiveOTOMessageRequest
//
// @param headers - CloseTopBoxInteractiveOTOMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseTopBoxInteractiveOTOMessageResponse
func (client *Client) CloseTopBoxInteractiveOTOMessageWithOptions(request *CloseTopBoxInteractiveOTOMessageRequest, headers *CloseTopBoxInteractiveOTOMessageHeaders, runtime *util.RuntimeOptions) (_result *CloseTopBoxInteractiveOTOMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Detail)) {
		body["detail"] = request.Detail
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
		Action:      tea.String("CloseTopBoxInteractiveOTOMessage"),
		Version:     tea.String("link_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/link/oToMessages/topBoxes/close"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseTopBoxInteractiveOTOMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务窗吊顶卡片关闭接口
//
// @param request - CloseTopBoxInteractiveOTOMessageRequest
//
// @return CloseTopBoxInteractiveOTOMessageResponse
func (client *Client) CloseTopBoxInteractiveOTOMessage(request *CloseTopBoxInteractiveOTOMessageRequest) (_result *CloseTopBoxInteractiveOTOMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CloseTopBoxInteractiveOTOMessageHeaders{}
	_result = &CloseTopBoxInteractiveOTOMessageResponse{}
	_body, _err := client.CloseTopBoxInteractiveOTOMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户授权信息
//
// @param request - GetFollowerAuthInfoRequest
//
// @param headers - GetFollowerAuthInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFollowerAuthInfoResponse
func (client *Client) GetFollowerAuthInfoWithOptions(request *GetFollowerAuthInfoRequest, headers *GetFollowerAuthInfoHeaders, runtime *util.RuntimeOptions) (_result *GetFollowerAuthInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["accountId"] = request.AccountId
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
		Action:      tea.String("GetFollowerAuthInfo"),
		Version:     tea.String("link_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/link/followers/authInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFollowerAuthInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户授权信息
//
// @param request - GetFollowerAuthInfoRequest
//
// @return GetFollowerAuthInfoResponse
func (client *Client) GetFollowerAuthInfo(request *GetFollowerAuthInfoRequest) (_result *GetFollowerAuthInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFollowerAuthInfoHeaders{}
	_result = &GetFollowerAuthInfoResponse{}
	_body, _err := client.GetFollowerAuthInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取服务窗关注人信息
//
// @param request - GetFollowerInfoRequest
//
// @param headers - GetFollowerInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFollowerInfoResponse
func (client *Client) GetFollowerInfoWithOptions(request *GetFollowerInfoRequest, headers *GetFollowerInfoHeaders, runtime *util.RuntimeOptions) (_result *GetFollowerInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["accountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
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
		Action:      tea.String("GetFollowerInfo"),
		Version:     tea.String("link_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/link/followers/infos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFollowerInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取服务窗关注人信息
//
// @param request - GetFollowerInfoRequest
//
// @return GetFollowerInfoResponse
func (client *Client) GetFollowerInfo(request *GetFollowerInfoRequest) (_result *GetFollowerInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFollowerInfoHeaders{}
	_result = &GetFollowerInfoResponse{}
	_body, _err := client.GetFollowerInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务窗图片消息下载地址获取接口
//
// @param request - GetPictureDownloadUrlRequest
//
// @param headers - GetPictureDownloadUrlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPictureDownloadUrlResponse
func (client *Client) GetPictureDownloadUrlWithOptions(request *GetPictureDownloadUrlRequest, headers *GetPictureDownloadUrlHeaders, runtime *util.RuntimeOptions) (_result *GetPictureDownloadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DownloadCode)) {
		query["downloadCode"] = request.DownloadCode
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["sessionId"] = request.SessionId
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
		Action:      tea.String("GetPictureDownloadUrl"),
		Version:     tea.String("link_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/link/oToMessages/pictures/downloadUrls"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPictureDownloadUrlResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务窗图片消息下载地址获取接口
//
// @param request - GetPictureDownloadUrlRequest
//
// @return GetPictureDownloadUrlResponse
func (client *Client) GetPictureDownloadUrl(request *GetPictureDownloadUrlRequest) (_result *GetPictureDownloadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPictureDownloadUrlHeaders{}
	_result = &GetPictureDownloadUrlResponse{}
	_body, _err := client.GetPictureDownloadUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户关注状态
//
// @param request - GetUserFollowStatusRequest
//
// @param headers - GetUserFollowStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserFollowStatusResponse
func (client *Client) GetUserFollowStatusWithOptions(request *GetUserFollowStatusRequest, headers *GetUserFollowStatusHeaders, runtime *util.RuntimeOptions) (_result *GetUserFollowStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["accountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
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
		Action:      tea.String("GetUserFollowStatus"),
		Version:     tea.String("link_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/link/followers/statuses"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserFollowStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户关注状态
//
// @param request - GetUserFollowStatusRequest
//
// @return GetUserFollowStatusResponse
func (client *Client) GetUserFollowStatus(request *GetUserFollowStatusRequest) (_result *GetUserFollowStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserFollowStatusHeaders{}
	_result = &GetUserFollowStatusResponse{}
	_body, _err := client.GetUserFollowStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业下服务窗帐号列表
//
// @param headers - ListAccountHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccountResponse
func (client *Client) ListAccountWithOptions(headers *ListAccountHeaders, runtime *util.RuntimeOptions) (_result *ListAccountResponse, _err error) {
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
		Action:      tea.String("ListAccount"),
		Version:     tea.String("link_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/link/accounts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAccountResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业下服务窗帐号列表
//
// @return ListAccountResponse
func (client *Client) ListAccount() (_result *ListAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListAccountHeaders{}
	_result = &ListAccountResponse{}
	_body, _err := client.ListAccountWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 第三方企业应用查询服务窗帐号列表
//
// @param headers - ListAccountInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccountInfoResponse
func (client *Client) ListAccountInfoWithOptions(headers *ListAccountInfoHeaders, runtime *util.RuntimeOptions) (_result *ListAccountInfoResponse, _err error) {
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
		Action:      tea.String("ListAccountInfo"),
		Version:     tea.String("link_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/link/isv/accounts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAccountInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 第三方企业应用查询服务窗帐号列表
//
// @return ListAccountInfoResponse
func (client *Client) ListAccountInfo() (_result *ListAccountInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListAccountInfoHeaders{}
	_result = &ListAccountInfoResponse{}
	_body, _err := client.ListAccountInfoWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量获取服务窗关注人列表
//
// @param request - ListFollowerRequest
//
// @param headers - ListFollowerHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFollowerResponse
func (client *Client) ListFollowerWithOptions(request *ListFollowerRequest, headers *ListFollowerHeaders, runtime *util.RuntimeOptions) (_result *ListFollowerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["accountId"] = request.AccountId
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
		Action:      tea.String("ListFollower"),
		Version:     tea.String("link_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/link/followers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFollowerResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量获取服务窗关注人列表
//
// @param request - ListFollowerRequest
//
// @return ListFollowerResponse
func (client *Client) ListFollower(request *ListFollowerRequest) (_result *ListFollowerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListFollowerHeaders{}
	_result = &ListFollowerResponse{}
	_body, _err := client.ListFollowerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 第三方企业应用查询用户是否关注服务窗
//
// @param request - QueryUserFollowStatusRequest
//
// @param headers - QueryUserFollowStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserFollowStatusResponse
func (client *Client) QueryUserFollowStatusWithOptions(request *QueryUserFollowStatusRequest, headers *QueryUserFollowStatusHeaders, runtime *util.RuntimeOptions) (_result *QueryUserFollowStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["accountId"] = request.AccountId
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
		Action:      tea.String("QueryUserFollowStatus"),
		Version:     tea.String("link_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/link/isv/followers/statuses"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserFollowStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 第三方企业应用查询用户是否关注服务窗
//
// @param request - QueryUserFollowStatusRequest
//
// @return QueryUserFollowStatusResponse
func (client *Client) QueryUserFollowStatus(request *QueryUserFollowStatusRequest) (_result *QueryUserFollowStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserFollowStatusHeaders{}
	_result = &QueryUserFollowStatusResponse{}
	_body, _err := client.QueryUserFollowStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发送服务窗客服消息
//
// @param request - SendAgentOTOMessageRequest
//
// @param headers - SendAgentOTOMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendAgentOTOMessageResponse
func (client *Client) SendAgentOTOMessageWithOptions(request *SendAgentOTOMessageRequest, headers *SendAgentOTOMessageHeaders, runtime *util.RuntimeOptions) (_result *SendAgentOTOMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Detail)) {
		body["detail"] = request.Detail
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
		Action:      tea.String("SendAgentOTOMessage"),
		Version:     tea.String("link_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/link/oToMessages/agentMessages"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendAgentOTOMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送服务窗客服消息
//
// @param request - SendAgentOTOMessageRequest
//
// @return SendAgentOTOMessageResponse
func (client *Client) SendAgentOTOMessage(request *SendAgentOTOMessageRequest) (_result *SendAgentOTOMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendAgentOTOMessageHeaders{}
	_result = &SendAgentOTOMessageResponse{}
	_body, _err := client.SendAgentOTOMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务窗互动卡片单发接口
//
// @param request - SendInteractiveOTOMessageRequest
//
// @param headers - SendInteractiveOTOMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendInteractiveOTOMessageResponse
func (client *Client) SendInteractiveOTOMessageWithOptions(request *SendInteractiveOTOMessageRequest, headers *SendInteractiveOTOMessageHeaders, runtime *util.RuntimeOptions) (_result *SendInteractiveOTOMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Detail)) {
		body["detail"] = request.Detail
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
		Action:      tea.String("SendInteractiveOTOMessage"),
		Version:     tea.String("link_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/link/oToMessages/interactiveMessages"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendInteractiveOTOMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务窗互动卡片单发接口
//
// @param request - SendInteractiveOTOMessageRequest
//
// @return SendInteractiveOTOMessageResponse
func (client *Client) SendInteractiveOTOMessage(request *SendInteractiveOTOMessageRequest) (_result *SendInteractiveOTOMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendInteractiveOTOMessageHeaders{}
	_result = &SendInteractiveOTOMessageResponse{}
	_body, _err := client.SendInteractiveOTOMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务窗吊顶卡片发送接口
//
// @param request - SendTopBoxInteractiveOTOMessageRequest
//
// @param headers - SendTopBoxInteractiveOTOMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendTopBoxInteractiveOTOMessageResponse
func (client *Client) SendTopBoxInteractiveOTOMessageWithOptions(request *SendTopBoxInteractiveOTOMessageRequest, headers *SendTopBoxInteractiveOTOMessageHeaders, runtime *util.RuntimeOptions) (_result *SendTopBoxInteractiveOTOMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Detail)) {
		body["detail"] = request.Detail
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
		Action:      tea.String("SendTopBoxInteractiveOTOMessage"),
		Version:     tea.String("link_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/link/oToMessages/topBoxes/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendTopBoxInteractiveOTOMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务窗吊顶卡片发送接口
//
// @param request - SendTopBoxInteractiveOTOMessageRequest
//
// @return SendTopBoxInteractiveOTOMessageResponse
func (client *Client) SendTopBoxInteractiveOTOMessage(request *SendTopBoxInteractiveOTOMessageRequest) (_result *SendTopBoxInteractiveOTOMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendTopBoxInteractiveOTOMessageHeaders{}
	_result = &SendTopBoxInteractiveOTOMessageResponse{}
	_body, _err := client.SendTopBoxInteractiveOTOMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务窗互动卡片修改接口
//
// @param request - UpdateInteractiveOTOMessageRequest
//
// @param headers - UpdateInteractiveOTOMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInteractiveOTOMessageResponse
func (client *Client) UpdateInteractiveOTOMessageWithOptions(request *UpdateInteractiveOTOMessageRequest, headers *UpdateInteractiveOTOMessageHeaders, runtime *util.RuntimeOptions) (_result *UpdateInteractiveOTOMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Detail)) {
		body["detail"] = request.Detail
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
		Action:      tea.String("UpdateInteractiveOTOMessage"),
		Version:     tea.String("link_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/link/oToMessages/interactiveMessages"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInteractiveOTOMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务窗互动卡片修改接口
//
// @param request - UpdateInteractiveOTOMessageRequest
//
// @return UpdateInteractiveOTOMessageResponse
func (client *Client) UpdateInteractiveOTOMessage(request *UpdateInteractiveOTOMessageRequest) (_result *UpdateInteractiveOTOMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateInteractiveOTOMessageHeaders{}
	_result = &UpdateInteractiveOTOMessageResponse{}
	_body, _err := client.UpdateInteractiveOTOMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务窗会话窗口快捷栏配置接口
//
// @param request - UpdateShortcutsRequest
//
// @param headers - UpdateShortcutsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateShortcutsResponse
func (client *Client) UpdateShortcutsWithOptions(request *UpdateShortcutsRequest, headers *UpdateShortcutsHeaders, runtime *util.RuntimeOptions) (_result *UpdateShortcutsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Details)) {
		body["details"] = request.Details
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["sessionId"] = request.SessionId
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
		Action:      tea.String("UpdateShortcuts"),
		Version:     tea.String("link_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/link/shortcuts"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateShortcutsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务窗会话窗口快捷栏配置接口
//
// @param request - UpdateShortcutsRequest
//
// @return UpdateShortcutsResponse
func (client *Client) UpdateShortcuts(request *UpdateShortcutsRequest) (_result *UpdateShortcutsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateShortcutsHeaders{}
	_result = &UpdateShortcutsResponse{}
	_body, _err := client.UpdateShortcutsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
