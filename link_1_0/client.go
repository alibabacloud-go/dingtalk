// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package link_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	UnionId   *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 响应结果
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
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
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

func (s *GetFollowerInfoResponseBodyResultUser) SetStatus(v string) *GetFollowerInfoResponseBodyResultUser {
	s.Status = &v
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFollowerInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DownloadCode *string `json:"downloadCode,omitempty" xml:"downloadCode,omitempty"`
	SessionId    *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
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
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 响应结果
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
	// 关注状态
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPictureDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetPictureDownloadUrlResponse) SetBody(v *GetPictureDownloadUrlResponseBody) *GetPictureDownloadUrlResponse {
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
	// 服务窗帐号ID，用于非服务窗自建应用场景下指定服务窗帐号。
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// 分页查询页大小。
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 分页查询下一页token,首页查询此字段可空，其它页查询时需要将此值设置炎上一次接口调用的token
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
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 响应结果
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
	// 下一页查询位置
	// 当此返回值为空时，则说明全部数据查询完成。
	// 当此返回值不为空时，可以将此值设置为下一次查询的参数。
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 用户列表
	UserList []*ListFollowerResponseBodyResultUserList `json:"userList,omitempty" xml:"userList,omitempty" type:"Repeated"`
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
	// 关注者昵称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 关注时间
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// 关注者userId，可用于消息推送等场景。
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFollowerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListFollowerResponse) SetBody(v *ListFollowerResponseBody) *ListFollowerResponse {
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
	// 消息详情
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
	// 消息体
	MessageBody *SendAgentOTOMessageRequestDetailMessageBody `json:"messageBody,omitempty" xml:"messageBody,omitempty" type:"Struct"`
	// 消息类型
	MsgType   *string `json:"msgType,omitempty" xml:"msgType,omitempty"`
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// 消息接收人id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 请求唯一 ID
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
	// 卡片消息
	ActionCard *SendAgentOTOMessageRequestDetailMessageBodyActionCard `json:"actionCard,omitempty" xml:"actionCard,omitempty" type:"Struct"`
	// 链接消息类型
	Link *SendAgentOTOMessageRequestDetailMessageBodyLink `json:"link,omitempty" xml:"link,omitempty" type:"Struct"`
	// markdown消息，仅对消息类型为markdown时有效
	Markdown *SendAgentOTOMessageRequestDetailMessageBodyMarkdown `json:"markdown,omitempty" xml:"markdown,omitempty" type:"Struct"`
	// 文本消息体  对于文本类型消息时必填
	Text *SendAgentOTOMessageRequestDetailMessageBodyText `json:"text,omitempty" xml:"text,omitempty" type:"Struct"`
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
	// 使用独立跳转ActionCard样式时的按钮列表；必须与buttonOrientation同时设置，且长度不超过1000字符。
	ButtonList []*SendAgentOTOMessageRequestDetailMessageBodyActionCardButtonList `json:"buttonList,omitempty" xml:"buttonList,omitempty" type:"Repeated"`
	// 按钮排列方式： 0：竖直排列 1：横向排列 必须与buttonList同时设置。
	ButtonOrientation *string `json:"buttonOrientation,omitempty" xml:"buttonOrientation,omitempty"`
	// 消息内容，支持markdown，语法参考标准markdown语法。1000个字符以内。
	Markdown *string `json:"markdown,omitempty" xml:"markdown,omitempty"`
	// 使用整体跳转ActionCard样式时的标题。必须与singleUrl同时设置，最长20个字符。
	SingleTitle *string `json:"singleTitle,omitempty" xml:"singleTitle,omitempty"`
	// 消息点击链接地址，当发送消息为小程序时支持小程序跳转链接，最长500个字符。
	SingleUrl *string `json:"singleUrl,omitempty" xml:"singleUrl,omitempty"`
	// 透出到会话列表和通知的文案
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
	// 使用独立跳转ActionCard样式时的跳转链接。
	ActionUrl *string `json:"actionUrl,omitempty" xml:"actionUrl,omitempty"`
	// 使用独立跳转ActionCard样式时的按钮的标题，最长20个字符。
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

type SendAgentOTOMessageRequestDetailMessageBodyLink struct {
	// 消息点击链接地址，当发送消息为小程序时支持小程序跳转链接。
	MessageUrl *string `json:"messageUrl,omitempty" xml:"messageUrl,omitempty"`
	// 图片地址
	PicUrl *string `json:"picUrl,omitempty" xml:"picUrl,omitempty"`
	// 消息描述，建议500字符以内。
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// 消息标题，建议100字符以内。
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
	// markdown格式的消息，建议500字符以内。
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// 首屏会话透出的展示内容。
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
	// 消息内容，建议500字符以内。
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
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 推送结果
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
	// 推送ID
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendAgentOTOMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// 消息详情
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
	CallbackUrl    *string `json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`
	CardBizId      *string `json:"cardBizId,omitempty" xml:"cardBizId,omitempty"`
	CardData       *string `json:"cardData,omitempty" xml:"cardData,omitempty"`
	CardTemplateId *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	// 消息接收人id
	UserId               *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 推送结果
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
	// 推送ID
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendInteractiveOTOMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SendInteractiveOTOMessageResponse) SetBody(v *SendInteractiveOTOMessageResponseBody) *SendInteractiveOTOMessageResponse {
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
	// 消息详情
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
	CardBizId            *string                                                `json:"cardBizId,omitempty" xml:"cardBizId,omitempty"`
	CardData             *string                                                `json:"cardData,omitempty" xml:"cardData,omitempty"`
	UpdateOptions        *UpdateInteractiveOTOMessageRequestDetailUpdateOptions `json:"updateOptions,omitempty" xml:"updateOptions,omitempty" type:"Struct"`
	UserIdPrivateDataMap *string                                                `json:"userIdPrivateDataMap,omitempty" xml:"userIdPrivateDataMap,omitempty"`
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
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 推送结果
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
	// 推送ID
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
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateInteractiveOTOMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateInteractiveOTOMessageResponse) SetBody(v *UpdateInteractiveOTOMessageResponseBody) *UpdateInteractiveOTOMessageResponse {
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
	_result = &GetFollowerInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetFollowerInfo"), tea.String("link_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/link/followers/infos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetPictureDownloadUrlResponse{}
	_body, _err := client.DoROARequest(tea.String("GetPictureDownloadUrl"), tea.String("link_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/link/oToMessages/pictures/downloadUrls"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListFollowerResponse{}
	_body, _err := client.DoROARequest(tea.String("ListFollower"), tea.String("link_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/link/followers"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) SendAgentOTOMessageWithOptions(request *SendAgentOTOMessageRequest, headers *SendAgentOTOMessageHeaders, runtime *util.RuntimeOptions) (_result *SendAgentOTOMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Detail))) {
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
	_result = &SendAgentOTOMessageResponse{}
	_body, _err := client.DoROARequest(tea.String("SendAgentOTOMessage"), tea.String("link_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/link/oToMessages/agentMessages"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) SendInteractiveOTOMessageWithOptions(request *SendInteractiveOTOMessageRequest, headers *SendInteractiveOTOMessageHeaders, runtime *util.RuntimeOptions) (_result *SendInteractiveOTOMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Detail))) {
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
	_result = &SendInteractiveOTOMessageResponse{}
	_body, _err := client.DoROARequest(tea.String("SendInteractiveOTOMessage"), tea.String("link_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/link/oToMessages/interactiveMessages"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) UpdateInteractiveOTOMessageWithOptions(request *UpdateInteractiveOTOMessageRequest, headers *UpdateInteractiveOTOMessageHeaders, runtime *util.RuntimeOptions) (_result *UpdateInteractiveOTOMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Detail))) {
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
	_result = &UpdateInteractiveOTOMessageResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateInteractiveOTOMessage"), tea.String("link_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/link/oToMessages/interactiveMessages"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
