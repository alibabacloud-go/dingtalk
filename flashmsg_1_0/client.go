// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package flashmsg_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type PrivateFieldMapValue struct {
	TipTitle     *string            `json:"tipTitle,omitempty" xml:"tipTitle,omitempty"`
	IsDingSend   *bool              `json:"isDingSend,omitempty" xml:"isDingSend,omitempty"`
	IsRead       *bool              `json:"isRead,omitempty" xml:"isRead,omitempty"`
	ButtonStatus *string            `json:"buttonStatus,omitempty" xml:"buttonStatus,omitempty"`
	Extension    map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
}

func (s PrivateFieldMapValue) String() string {
	return tea.Prettify(s)
}

func (s PrivateFieldMapValue) GoString() string {
	return s.String()
}

func (s *PrivateFieldMapValue) SetTipTitle(v string) *PrivateFieldMapValue {
	s.TipTitle = &v
	return s
}

func (s *PrivateFieldMapValue) SetIsDingSend(v bool) *PrivateFieldMapValue {
	s.IsDingSend = &v
	return s
}

func (s *PrivateFieldMapValue) SetIsRead(v bool) *PrivateFieldMapValue {
	s.IsRead = &v
	return s
}

func (s *PrivateFieldMapValue) SetButtonStatus(v string) *PrivateFieldMapValue {
	s.ButtonStatus = &v
	return s
}

func (s *PrivateFieldMapValue) SetExtension(v map[string]*string) *PrivateFieldMapValue {
	s.Extension = v
	return s
}

type GetBaseProfileListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetBaseProfileListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetBaseProfileListHeaders) GoString() string {
	return s.String()
}

func (s *GetBaseProfileListHeaders) SetCommonHeaders(v map[string]*string) *GetBaseProfileListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetBaseProfileListHeaders) SetXAcsDingtalkAccessToken(v string) *GetBaseProfileListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetBaseProfileListRequest struct {
	Body []*string `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s GetBaseProfileListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBaseProfileListRequest) GoString() string {
	return s.String()
}

func (s *GetBaseProfileListRequest) SetBody(v []*string) *GetBaseProfileListRequest {
	s.Body = v
	return s
}

type GetBaseProfileListResponseBody struct {
	Result  []*GetBaseProfileListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success *bool                                   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetBaseProfileListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBaseProfileListResponseBody) GoString() string {
	return s.String()
}

func (s *GetBaseProfileListResponseBody) SetResult(v []*GetBaseProfileListResponseBodyResult) *GetBaseProfileListResponseBody {
	s.Result = v
	return s
}

func (s *GetBaseProfileListResponseBody) SetSuccess(v bool) *GetBaseProfileListResponseBody {
	s.Success = &v
	return s
}

type GetBaseProfileListResponseBodyResult struct {
	AvatarMediaId *string `json:"avatarMediaId,omitempty" xml:"avatarMediaId,omitempty"`
	Nick          *string `json:"nick,omitempty" xml:"nick,omitempty"`
	NickPinyin    *string `json:"nickPinyin,omitempty" xml:"nickPinyin,omitempty"`
	UserId        *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetBaseProfileListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetBaseProfileListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetBaseProfileListResponseBodyResult) SetAvatarMediaId(v string) *GetBaseProfileListResponseBodyResult {
	s.AvatarMediaId = &v
	return s
}

func (s *GetBaseProfileListResponseBodyResult) SetNick(v string) *GetBaseProfileListResponseBodyResult {
	s.Nick = &v
	return s
}

func (s *GetBaseProfileListResponseBodyResult) SetNickPinyin(v string) *GetBaseProfileListResponseBodyResult {
	s.NickPinyin = &v
	return s
}

func (s *GetBaseProfileListResponseBodyResult) SetUserId(v string) *GetBaseProfileListResponseBodyResult {
	s.UserId = &v
	return s
}

type GetBaseProfileListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBaseProfileListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBaseProfileListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBaseProfileListResponse) GoString() string {
	return s.String()
}

func (s *GetBaseProfileListResponse) SetHeaders(v map[string]*string) *GetBaseProfileListResponse {
	s.Headers = v
	return s
}

func (s *GetBaseProfileListResponse) SetStatusCode(v int32) *GetBaseProfileListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBaseProfileListResponse) SetBody(v *GetBaseProfileListResponseBody) *GetBaseProfileListResponse {
	s.Body = v
	return s
}

type GetConversationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetConversationHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetConversationHeaders) GoString() string {
	return s.String()
}

func (s *GetConversationHeaders) SetCommonHeaders(v map[string]*string) *GetConversationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetConversationHeaders) SetXAcsDingtalkAccessToken(v string) *GetConversationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetConversationRequest struct {
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s GetConversationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConversationRequest) GoString() string {
	return s.String()
}

func (s *GetConversationRequest) SetOpenConversationId(v string) *GetConversationRequest {
	s.OpenConversationId = &v
	return s
}

type GetConversationResponseBody struct {
	Result  *GetConversationResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                              `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetConversationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConversationResponseBody) GoString() string {
	return s.String()
}

func (s *GetConversationResponseBody) SetResult(v *GetConversationResponseBodyResult) *GetConversationResponseBody {
	s.Result = v
	return s
}

func (s *GetConversationResponseBody) SetSuccess(v bool) *GetConversationResponseBody {
	s.Success = &v
	return s
}

type GetConversationResponseBodyResult struct {
	CorpId      *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	MemberCount *int32  `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetConversationResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetConversationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetConversationResponseBodyResult) SetCorpId(v string) *GetConversationResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *GetConversationResponseBodyResult) SetMemberCount(v int32) *GetConversationResponseBodyResult {
	s.MemberCount = &v
	return s
}

func (s *GetConversationResponseBodyResult) SetTitle(v string) *GetConversationResponseBodyResult {
	s.Title = &v
	return s
}

type GetConversationResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetConversationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConversationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConversationResponse) GoString() string {
	return s.String()
}

func (s *GetConversationResponse) SetHeaders(v map[string]*string) *GetConversationResponse {
	s.Headers = v
	return s
}

func (s *GetConversationResponse) SetStatusCode(v int32) *GetConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConversationResponse) SetBody(v *GetConversationResponseBody) *GetConversationResponse {
	s.Body = v
	return s
}

type GetMemberListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetMemberListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetMemberListHeaders) GoString() string {
	return s.String()
}

func (s *GetMemberListHeaders) SetCommonHeaders(v map[string]*string) *GetMemberListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMemberListHeaders) SetXAcsDingtalkAccessToken(v string) *GetMemberListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetMemberListRequest struct {
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	PageNumber         *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize           *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetMemberListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMemberListRequest) GoString() string {
	return s.String()
}

func (s *GetMemberListRequest) SetOpenConversationId(v string) *GetMemberListRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GetMemberListRequest) SetPageNumber(v int32) *GetMemberListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetMemberListRequest) SetPageSize(v int32) *GetMemberListRequest {
	s.PageSize = &v
	return s
}

type GetMemberListResponseBody struct {
	Result []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetMemberListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMemberListResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemberListResponseBody) SetResult(v []*string) *GetMemberListResponseBody {
	s.Result = v
	return s
}

type GetMemberListResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMemberListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMemberListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMemberListResponse) GoString() string {
	return s.String()
}

func (s *GetMemberListResponse) SetHeaders(v map[string]*string) *GetMemberListResponse {
	s.Headers = v
	return s
}

func (s *GetMemberListResponse) SetStatusCode(v int32) *GetMemberListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMemberListResponse) SetBody(v *GetMemberListResponseBody) *GetMemberListResponse {
	s.Body = v
	return s
}

type SendDingTipHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendDingTipHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendDingTipHeaders) GoString() string {
	return s.String()
}

func (s *SendDingTipHeaders) SetCommonHeaders(v map[string]*string) *SendDingTipHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendDingTipHeaders) SetXAcsDingtalkAccessToken(v string) *SendDingTipHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendDingTipRequest struct {
	Extension      map[string]*string      `json:"extension,omitempty" xml:"extension,omitempty"`
	Link           *SendDingTipRequestLink `json:"link,omitempty" xml:"link,omitempty" type:"Struct"`
	MessageId      *string                 `json:"messageId,omitempty" xml:"messageId,omitempty"`
	ReceiverUserId []*string               `json:"receiverUserId,omitempty" xml:"receiverUserId,omitempty" type:"Repeated"`
	TextContent    *string                 `json:"textContent,omitempty" xml:"textContent,omitempty"`
}

func (s SendDingTipRequest) String() string {
	return tea.Prettify(s)
}

func (s SendDingTipRequest) GoString() string {
	return s.String()
}

func (s *SendDingTipRequest) SetExtension(v map[string]*string) *SendDingTipRequest {
	s.Extension = v
	return s
}

func (s *SendDingTipRequest) SetLink(v *SendDingTipRequestLink) *SendDingTipRequest {
	s.Link = v
	return s
}

func (s *SendDingTipRequest) SetMessageId(v string) *SendDingTipRequest {
	s.MessageId = &v
	return s
}

func (s *SendDingTipRequest) SetReceiverUserId(v []*string) *SendDingTipRequest {
	s.ReceiverUserId = v
	return s
}

func (s *SendDingTipRequest) SetTextContent(v string) *SendDingTipRequest {
	s.TextContent = &v
	return s
}

type SendDingTipRequestLink struct {
	Extension  map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
	LinkUrl    *string            `json:"linkUrl,omitempty" xml:"linkUrl,omitempty"`
	PicMediaId *string            `json:"picMediaId,omitempty" xml:"picMediaId,omitempty"`
	Text       *string            `json:"text,omitempty" xml:"text,omitempty"`
}

func (s SendDingTipRequestLink) String() string {
	return tea.Prettify(s)
}

func (s SendDingTipRequestLink) GoString() string {
	return s.String()
}

func (s *SendDingTipRequestLink) SetExtension(v map[string]*string) *SendDingTipRequestLink {
	s.Extension = v
	return s
}

func (s *SendDingTipRequestLink) SetLinkUrl(v string) *SendDingTipRequestLink {
	s.LinkUrl = &v
	return s
}

func (s *SendDingTipRequestLink) SetPicMediaId(v string) *SendDingTipRequestLink {
	s.PicMediaId = &v
	return s
}

func (s *SendDingTipRequestLink) SetText(v string) *SendDingTipRequestLink {
	s.Text = &v
	return s
}

type SendDingTipResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SendDingTipResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendDingTipResponseBody) GoString() string {
	return s.String()
}

func (s *SendDingTipResponseBody) SetSuccess(v bool) *SendDingTipResponseBody {
	s.Success = &v
	return s
}

type SendDingTipResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendDingTipResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendDingTipResponse) String() string {
	return tea.Prettify(s)
}

func (s SendDingTipResponse) GoString() string {
	return s.String()
}

func (s *SendDingTipResponse) SetHeaders(v map[string]*string) *SendDingTipResponse {
	s.Headers = v
	return s
}

func (s *SendDingTipResponse) SetStatusCode(v int32) *SendDingTipResponse {
	s.StatusCode = &v
	return s
}

func (s *SendDingTipResponse) SetBody(v *SendDingTipResponseBody) *SendDingTipResponse {
	s.Body = v
	return s
}

type SendMessageTipHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendMessageTipHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendMessageTipHeaders) GoString() string {
	return s.String()
}

func (s *SendMessageTipHeaders) SetCommonHeaders(v map[string]*string) *SendMessageTipHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendMessageTipHeaders) SetXAcsDingtalkAccessToken(v string) *SendMessageTipHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendMessageTipRequest struct {
	DefaultView        *SendMessageTipRequestDefaultView `json:"defaultView,omitempty" xml:"defaultView,omitempty" type:"Struct"`
	MessageId          *string                           `json:"messageId,omitempty" xml:"messageId,omitempty"`
	OpenConversationId *string                           `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	PrivateFieldMap    map[string]*PrivateFieldMapValue  `json:"privateFieldMap,omitempty" xml:"privateFieldMap,omitempty"`
	PublicField        *SendMessageTipRequestPublicField `json:"publicField,omitempty" xml:"publicField,omitempty" type:"Struct"`
	ReceiverUserId     []*string                         `json:"receiverUserId,omitempty" xml:"receiverUserId,omitempty" type:"Repeated"`
}

func (s SendMessageTipRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageTipRequest) GoString() string {
	return s.String()
}

func (s *SendMessageTipRequest) SetDefaultView(v *SendMessageTipRequestDefaultView) *SendMessageTipRequest {
	s.DefaultView = v
	return s
}

func (s *SendMessageTipRequest) SetMessageId(v string) *SendMessageTipRequest {
	s.MessageId = &v
	return s
}

func (s *SendMessageTipRequest) SetOpenConversationId(v string) *SendMessageTipRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SendMessageTipRequest) SetPrivateFieldMap(v map[string]*PrivateFieldMapValue) *SendMessageTipRequest {
	s.PrivateFieldMap = v
	return s
}

func (s *SendMessageTipRequest) SetPublicField(v *SendMessageTipRequestPublicField) *SendMessageTipRequest {
	s.PublicField = v
	return s
}

func (s *SendMessageTipRequest) SetReceiverUserId(v []*string) *SendMessageTipRequest {
	s.ReceiverUserId = v
	return s
}

type SendMessageTipRequestDefaultView struct {
	ActionName         *string `json:"actionName,omitempty" xml:"actionName,omitempty"`
	ActionUrl          *string `json:"actionUrl,omitempty" xml:"actionUrl,omitempty"`
	AuthCode           *string `json:"authCode,omitempty" xml:"authCode,omitempty"`
	AuthMediaId        *string `json:"authMediaId,omitempty" xml:"authMediaId,omitempty"`
	CardTitle          *string `json:"cardTitle,omitempty" xml:"cardTitle,omitempty"`
	CardTitleColor     *string `json:"cardTitleColor,omitempty" xml:"cardTitleColor,omitempty"`
	Desc               *string `json:"desc,omitempty" xml:"desc,omitempty"`
	MediaId            *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	NeedShowUpdateTail *string `json:"needShowUpdateTail,omitempty" xml:"needShowUpdateTail,omitempty"`
	Summary            *string `json:"summary,omitempty" xml:"summary,omitempty"`
	Title              *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SendMessageTipRequestDefaultView) String() string {
	return tea.Prettify(s)
}

func (s SendMessageTipRequestDefaultView) GoString() string {
	return s.String()
}

func (s *SendMessageTipRequestDefaultView) SetActionName(v string) *SendMessageTipRequestDefaultView {
	s.ActionName = &v
	return s
}

func (s *SendMessageTipRequestDefaultView) SetActionUrl(v string) *SendMessageTipRequestDefaultView {
	s.ActionUrl = &v
	return s
}

func (s *SendMessageTipRequestDefaultView) SetAuthCode(v string) *SendMessageTipRequestDefaultView {
	s.AuthCode = &v
	return s
}

func (s *SendMessageTipRequestDefaultView) SetAuthMediaId(v string) *SendMessageTipRequestDefaultView {
	s.AuthMediaId = &v
	return s
}

func (s *SendMessageTipRequestDefaultView) SetCardTitle(v string) *SendMessageTipRequestDefaultView {
	s.CardTitle = &v
	return s
}

func (s *SendMessageTipRequestDefaultView) SetCardTitleColor(v string) *SendMessageTipRequestDefaultView {
	s.CardTitleColor = &v
	return s
}

func (s *SendMessageTipRequestDefaultView) SetDesc(v string) *SendMessageTipRequestDefaultView {
	s.Desc = &v
	return s
}

func (s *SendMessageTipRequestDefaultView) SetMediaId(v string) *SendMessageTipRequestDefaultView {
	s.MediaId = &v
	return s
}

func (s *SendMessageTipRequestDefaultView) SetNeedShowUpdateTail(v string) *SendMessageTipRequestDefaultView {
	s.NeedShowUpdateTail = &v
	return s
}

func (s *SendMessageTipRequestDefaultView) SetSummary(v string) *SendMessageTipRequestDefaultView {
	s.Summary = &v
	return s
}

func (s *SendMessageTipRequestDefaultView) SetTitle(v string) *SendMessageTipRequestDefaultView {
	s.Title = &v
	return s
}

type SendMessageTipRequestPublicField struct {
	DetailUrl        *string            `json:"detailUrl,omitempty" xml:"detailUrl,omitempty"`
	DurationDesc     *string            `json:"durationDesc,omitempty" xml:"durationDesc,omitempty"`
	Extension        map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
	IsExpired        *bool              `json:"isExpired,omitempty" xml:"isExpired,omitempty"`
	ReadActionUrl    *string            `json:"readActionUrl,omitempty" xml:"readActionUrl,omitempty"`
	ReadProgressDesc *string            `json:"readProgressDesc,omitempty" xml:"readProgressDesc,omitempty"`
}

func (s SendMessageTipRequestPublicField) String() string {
	return tea.Prettify(s)
}

func (s SendMessageTipRequestPublicField) GoString() string {
	return s.String()
}

func (s *SendMessageTipRequestPublicField) SetDetailUrl(v string) *SendMessageTipRequestPublicField {
	s.DetailUrl = &v
	return s
}

func (s *SendMessageTipRequestPublicField) SetDurationDesc(v string) *SendMessageTipRequestPublicField {
	s.DurationDesc = &v
	return s
}

func (s *SendMessageTipRequestPublicField) SetExtension(v map[string]*string) *SendMessageTipRequestPublicField {
	s.Extension = v
	return s
}

func (s *SendMessageTipRequestPublicField) SetIsExpired(v bool) *SendMessageTipRequestPublicField {
	s.IsExpired = &v
	return s
}

func (s *SendMessageTipRequestPublicField) SetReadActionUrl(v string) *SendMessageTipRequestPublicField {
	s.ReadActionUrl = &v
	return s
}

func (s *SendMessageTipRequestPublicField) SetReadProgressDesc(v string) *SendMessageTipRequestPublicField {
	s.ReadProgressDesc = &v
	return s
}

type SendMessageTipResponseBody struct {
	Result  *SendMessageTipResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SendMessageTipResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMessageTipResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageTipResponseBody) SetResult(v *SendMessageTipResponseBodyResult) *SendMessageTipResponseBody {
	s.Result = v
	return s
}

func (s *SendMessageTipResponseBody) SetSuccess(v bool) *SendMessageTipResponseBody {
	s.Success = &v
	return s
}

type SendMessageTipResponseBodyResult struct {
	BizId          *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	CardInstanceId *string `json:"cardInstanceId,omitempty" xml:"cardInstanceId,omitempty"`
}

func (s SendMessageTipResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SendMessageTipResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SendMessageTipResponseBodyResult) SetBizId(v string) *SendMessageTipResponseBodyResult {
	s.BizId = &v
	return s
}

func (s *SendMessageTipResponseBodyResult) SetCardInstanceId(v string) *SendMessageTipResponseBodyResult {
	s.CardInstanceId = &v
	return s
}

type SendMessageTipResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendMessageTipResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendMessageTipResponse) String() string {
	return tea.Prettify(s)
}

func (s SendMessageTipResponse) GoString() string {
	return s.String()
}

func (s *SendMessageTipResponse) SetHeaders(v map[string]*string) *SendMessageTipResponse {
	s.Headers = v
	return s
}

func (s *SendMessageTipResponse) SetStatusCode(v int32) *SendMessageTipResponse {
	s.StatusCode = &v
	return s
}

func (s *SendMessageTipResponse) SetBody(v *SendMessageTipResponseBody) *SendMessageTipResponse {
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

func (client *Client) GetBaseProfileListWithOptions(request *GetBaseProfileListRequest, headers *GetBaseProfileListHeaders, runtime *util.RuntimeOptions) (_result *GetBaseProfileListResponse, _err error) {
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
		Action:      tea.String("GetBaseProfileList"),
		Version:     tea.String("flashmsg_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/flashmsg/users/baseInfos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBaseProfileListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBaseProfileList(request *GetBaseProfileListRequest) (_result *GetBaseProfileListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetBaseProfileListHeaders{}
	_result = &GetBaseProfileListResponse{}
	_body, _err := client.GetBaseProfileListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConversationWithOptions(request *GetConversationRequest, headers *GetConversationHeaders, runtime *util.RuntimeOptions) (_result *GetConversationResponse, _err error) {
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
		Action:      tea.String("GetConversation"),
		Version:     tea.String("flashmsg_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/flashmsg/conversations/infos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConversationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConversation(request *GetConversationRequest) (_result *GetConversationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetConversationHeaders{}
	_result = &GetConversationResponse{}
	_body, _err := client.GetConversationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMemberListWithOptions(request *GetMemberListRequest, headers *GetMemberListHeaders, runtime *util.RuntimeOptions) (_result *GetMemberListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		query["openConversationId"] = request.OpenConversationId
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
		Action:      tea.String("GetMemberList"),
		Version:     tea.String("flashmsg_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/flashmsg/conversations/memberIdLists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMemberListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMemberList(request *GetMemberListRequest) (_result *GetMemberListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetMemberListHeaders{}
	_result = &GetMemberListResponse{}
	_body, _err := client.GetMemberListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendDingTipWithOptions(request *SendDingTipRequest, headers *SendDingTipHeaders, runtime *util.RuntimeOptions) (_result *SendDingTipResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Extension)) {
		body["extension"] = request.Extension
	}

	if !tea.BoolValue(util.IsUnset(request.Link)) {
		body["link"] = request.Link
	}

	if !tea.BoolValue(util.IsUnset(request.MessageId)) {
		body["messageId"] = request.MessageId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUserId)) {
		body["receiverUserId"] = request.ReceiverUserId
	}

	if !tea.BoolValue(util.IsUnset(request.TextContent)) {
		body["textContent"] = request.TextContent
	}

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
		Action:      tea.String("SendDingTip"),
		Version:     tea.String("flashmsg_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/flashmsg/ding/messages/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendDingTipResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendDingTip(request *SendDingTipRequest) (_result *SendDingTipResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendDingTipHeaders{}
	_result = &SendDingTipResponse{}
	_body, _err := client.SendDingTipWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendMessageTipWithOptions(request *SendMessageTipRequest, headers *SendMessageTipHeaders, runtime *util.RuntimeOptions) (_result *SendMessageTipResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefaultView)) {
		body["defaultView"] = request.DefaultView
	}

	if !tea.BoolValue(util.IsUnset(request.MessageId)) {
		body["messageId"] = request.MessageId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateFieldMap)) {
		body["privateFieldMap"] = request.PrivateFieldMap
	}

	if !tea.BoolValue(util.IsUnset(request.PublicField)) {
		body["publicField"] = request.PublicField
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUserId)) {
		body["receiverUserId"] = request.ReceiverUserId
	}

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
		Action:      tea.String("SendMessageTip"),
		Version:     tea.String("flashmsg_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/flashmsg/messages/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendMessageTipResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendMessageTip(request *SendMessageTipRequest) (_result *SendMessageTipResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendMessageTipHeaders{}
	_result = &SendMessageTipResponse{}
	_body, _err := client.SendMessageTipWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
