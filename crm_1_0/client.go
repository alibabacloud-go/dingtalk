// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package crm_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetOfficialAccountContactsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOfficialAccountContactsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountContactsHeaders) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountContactsHeaders) SetCommonHeaders(v map[string]*string) *GetOfficialAccountContactsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOfficialAccountContactsHeaders) SetXAcsDingtalkAccessToken(v string) *GetOfficialAccountContactsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOfficialAccountContactsRequest struct {
	// 取数游标，第一次传0
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 分页大小，最大不超过10
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
}

func (s GetOfficialAccountContactsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountContactsRequest) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountContactsRequest) SetNextToken(v string) *GetOfficialAccountContactsRequest {
	s.NextToken = &v
	return s
}

func (s *GetOfficialAccountContactsRequest) SetMaxResults(v int64) *GetOfficialAccountContactsRequest {
	s.MaxResults = &v
	return s
}

type GetOfficialAccountContactsResponseBody struct {
	// 下一页的游标，为null则表示无数据
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 分页大小
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 客户数据节点
	Values []*GetOfficialAccountContactsResponseBodyValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetOfficialAccountContactsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountContactsResponseBody) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountContactsResponseBody) SetNextToken(v string) *GetOfficialAccountContactsResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetOfficialAccountContactsResponseBody) SetMaxResults(v int64) *GetOfficialAccountContactsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetOfficialAccountContactsResponseBody) SetValues(v []*GetOfficialAccountContactsResponseBodyValues) *GetOfficialAccountContactsResponseBody {
	s.Values = v
	return s
}

type GetOfficialAccountContactsResponseBodyValues struct {
	// 用户userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用户的联系人数据
	Contacts []*GetOfficialAccountContactsResponseBodyValuesContacts `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
}

func (s GetOfficialAccountContactsResponseBodyValues) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountContactsResponseBodyValues) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountContactsResponseBodyValues) SetUserId(v string) *GetOfficialAccountContactsResponseBodyValues {
	s.UserId = &v
	return s
}

func (s *GetOfficialAccountContactsResponseBodyValues) SetContacts(v []*GetOfficialAccountContactsResponseBodyValuesContacts) *GetOfficialAccountContactsResponseBodyValues {
	s.Contacts = v
	return s
}

type GetOfficialAccountContactsResponseBodyValuesContacts struct {
	// 创建记录的用户昵称
	CreatorNick *string `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	// 记录修改时间
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// 记录创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建记录的用户ID
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// 数据ID
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// 数据内容
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// 扩展数据内容
	ExtendData map[string]interface{} `json:"extendData,omitempty" xml:"extendData,omitempty"`
	// 数据权限信息
	Permission *GetOfficialAccountContactsResponseBodyValuesContactsPermission `json:"permission,omitempty" xml:"permission,omitempty" type:"Struct"`
}

func (s GetOfficialAccountContactsResponseBodyValuesContacts) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountContactsResponseBodyValuesContacts) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountContactsResponseBodyValuesContacts) SetCreatorNick(v string) *GetOfficialAccountContactsResponseBodyValuesContacts {
	s.CreatorNick = &v
	return s
}

func (s *GetOfficialAccountContactsResponseBodyValuesContacts) SetModifyTime(v string) *GetOfficialAccountContactsResponseBodyValuesContacts {
	s.ModifyTime = &v
	return s
}

func (s *GetOfficialAccountContactsResponseBodyValuesContacts) SetCreateTime(v string) *GetOfficialAccountContactsResponseBodyValuesContacts {
	s.CreateTime = &v
	return s
}

func (s *GetOfficialAccountContactsResponseBodyValuesContacts) SetCreatorUserId(v string) *GetOfficialAccountContactsResponseBodyValuesContacts {
	s.CreatorUserId = &v
	return s
}

func (s *GetOfficialAccountContactsResponseBodyValuesContacts) SetInstanceId(v string) *GetOfficialAccountContactsResponseBodyValuesContacts {
	s.InstanceId = &v
	return s
}

func (s *GetOfficialAccountContactsResponseBodyValuesContacts) SetData(v map[string]interface{}) *GetOfficialAccountContactsResponseBodyValuesContacts {
	s.Data = v
	return s
}

func (s *GetOfficialAccountContactsResponseBodyValuesContacts) SetExtendData(v map[string]interface{}) *GetOfficialAccountContactsResponseBodyValuesContacts {
	s.ExtendData = v
	return s
}

func (s *GetOfficialAccountContactsResponseBodyValuesContacts) SetPermission(v *GetOfficialAccountContactsResponseBodyValuesContactsPermission) *GetOfficialAccountContactsResponseBodyValuesContacts {
	s.Permission = v
	return s
}

type GetOfficialAccountContactsResponseBodyValuesContactsPermission struct {
	// 协同人用户ID列表
	ParticipantStaffIds []*string `json:"participantStaffIds,omitempty" xml:"participantStaffIds,omitempty" type:"Repeated"`
	// 负责人用户ID列表
	OwnerStaffIds []*string `json:"ownerStaffIds,omitempty" xml:"ownerStaffIds,omitempty" type:"Repeated"`
}

func (s GetOfficialAccountContactsResponseBodyValuesContactsPermission) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountContactsResponseBodyValuesContactsPermission) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountContactsResponseBodyValuesContactsPermission) SetParticipantStaffIds(v []*string) *GetOfficialAccountContactsResponseBodyValuesContactsPermission {
	s.ParticipantStaffIds = v
	return s
}

func (s *GetOfficialAccountContactsResponseBodyValuesContactsPermission) SetOwnerStaffIds(v []*string) *GetOfficialAccountContactsResponseBodyValuesContactsPermission {
	s.OwnerStaffIds = v
	return s
}

type GetOfficialAccountContactsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOfficialAccountContactsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOfficialAccountContactsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountContactsResponse) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountContactsResponse) SetHeaders(v map[string]*string) *GetOfficialAccountContactsResponse {
	s.Headers = v
	return s
}

func (s *GetOfficialAccountContactsResponse) SetBody(v *GetOfficialAccountContactsResponseBody) *GetOfficialAccountContactsResponse {
	s.Body = v
	return s
}

type ServiceWindowMessageBatchPushHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ServiceWindowMessageBatchPushHeaders) String() string {
	return tea.Prettify(s)
}

func (s ServiceWindowMessageBatchPushHeaders) GoString() string {
	return s.String()
}

func (s *ServiceWindowMessageBatchPushHeaders) SetCommonHeaders(v map[string]*string) *ServiceWindowMessageBatchPushHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ServiceWindowMessageBatchPushHeaders) SetXAcsDingtalkAccessToken(v string) *ServiceWindowMessageBatchPushHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ServiceWindowMessageBatchPushRequest struct {
	Detail             *ServiceWindowMessageBatchPushRequestDetail `json:"detail,omitempty" xml:"detail,omitempty" type:"Struct"`
	BizId              *string                                     `json:"bizId,omitempty" xml:"bizId,omitempty"`
	DingIsvOrgId       *int64                                      `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId          *int64                                      `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingTokenGrantType *int64                                      `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	DingSuiteKey       *string                                     `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
}

func (s ServiceWindowMessageBatchPushRequest) String() string {
	return tea.Prettify(s)
}

func (s ServiceWindowMessageBatchPushRequest) GoString() string {
	return s.String()
}

func (s *ServiceWindowMessageBatchPushRequest) SetDetail(v *ServiceWindowMessageBatchPushRequestDetail) *ServiceWindowMessageBatchPushRequest {
	s.Detail = v
	return s
}

func (s *ServiceWindowMessageBatchPushRequest) SetBizId(v string) *ServiceWindowMessageBatchPushRequest {
	s.BizId = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequest) SetDingIsvOrgId(v int64) *ServiceWindowMessageBatchPushRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequest) SetDingOrgId(v int64) *ServiceWindowMessageBatchPushRequest {
	s.DingOrgId = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequest) SetDingTokenGrantType(v int64) *ServiceWindowMessageBatchPushRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequest) SetDingSuiteKey(v string) *ServiceWindowMessageBatchPushRequest {
	s.DingSuiteKey = &v
	return s
}

type ServiceWindowMessageBatchPushRequestDetail struct {
	MsgType      *string                                                `json:"msgType,omitempty" xml:"msgType,omitempty"`
	Uuid         *string                                                `json:"uuid,omitempty" xml:"uuid,omitempty"`
	BizRequestId *string                                                `json:"bizRequestId,omitempty" xml:"bizRequestId,omitempty"`
	UserIdList   []*string                                              `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
	MessageBody  *ServiceWindowMessageBatchPushRequestDetailMessageBody `json:"messageBody,omitempty" xml:"messageBody,omitempty" type:"Struct"`
	SendToAll    *bool                                                  `json:"sendToAll,omitempty" xml:"sendToAll,omitempty"`
}

func (s ServiceWindowMessageBatchPushRequestDetail) String() string {
	return tea.Prettify(s)
}

func (s ServiceWindowMessageBatchPushRequestDetail) GoString() string {
	return s.String()
}

func (s *ServiceWindowMessageBatchPushRequestDetail) SetMsgType(v string) *ServiceWindowMessageBatchPushRequestDetail {
	s.MsgType = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetail) SetUuid(v string) *ServiceWindowMessageBatchPushRequestDetail {
	s.Uuid = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetail) SetBizRequestId(v string) *ServiceWindowMessageBatchPushRequestDetail {
	s.BizRequestId = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetail) SetUserIdList(v []*string) *ServiceWindowMessageBatchPushRequestDetail {
	s.UserIdList = v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetail) SetMessageBody(v *ServiceWindowMessageBatchPushRequestDetailMessageBody) *ServiceWindowMessageBatchPushRequestDetail {
	s.MessageBody = v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetail) SetSendToAll(v bool) *ServiceWindowMessageBatchPushRequestDetail {
	s.SendToAll = &v
	return s
}

type ServiceWindowMessageBatchPushRequestDetailMessageBody struct {
	Text       *ServiceWindowMessageBatchPushRequestDetailMessageBodyText       `json:"text,omitempty" xml:"text,omitempty" type:"Struct"`
	Markdown   *ServiceWindowMessageBatchPushRequestDetailMessageBodyMarkdown   `json:"markdown,omitempty" xml:"markdown,omitempty" type:"Struct"`
	Link       *ServiceWindowMessageBatchPushRequestDetailMessageBodyLink       `json:"link,omitempty" xml:"link,omitempty" type:"Struct"`
	ActionCard *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard `json:"actionCard,omitempty" xml:"actionCard,omitempty" type:"Struct"`
}

func (s ServiceWindowMessageBatchPushRequestDetailMessageBody) String() string {
	return tea.Prettify(s)
}

func (s ServiceWindowMessageBatchPushRequestDetailMessageBody) GoString() string {
	return s.String()
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBody) SetText(v *ServiceWindowMessageBatchPushRequestDetailMessageBodyText) *ServiceWindowMessageBatchPushRequestDetailMessageBody {
	s.Text = v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBody) SetMarkdown(v *ServiceWindowMessageBatchPushRequestDetailMessageBodyMarkdown) *ServiceWindowMessageBatchPushRequestDetailMessageBody {
	s.Markdown = v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBody) SetLink(v *ServiceWindowMessageBatchPushRequestDetailMessageBodyLink) *ServiceWindowMessageBatchPushRequestDetailMessageBody {
	s.Link = v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBody) SetActionCard(v *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard) *ServiceWindowMessageBatchPushRequestDetailMessageBody {
	s.ActionCard = v
	return s
}

type ServiceWindowMessageBatchPushRequestDetailMessageBodyText struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s ServiceWindowMessageBatchPushRequestDetailMessageBodyText) String() string {
	return tea.Prettify(s)
}

func (s ServiceWindowMessageBatchPushRequestDetailMessageBodyText) GoString() string {
	return s.String()
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyText) SetContent(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyText {
	s.Content = &v
	return s
}

type ServiceWindowMessageBatchPushRequestDetailMessageBodyMarkdown struct {
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	Text  *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s ServiceWindowMessageBatchPushRequestDetailMessageBodyMarkdown) String() string {
	return tea.Prettify(s)
}

func (s ServiceWindowMessageBatchPushRequestDetailMessageBodyMarkdown) GoString() string {
	return s.String()
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyMarkdown) SetTitle(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyMarkdown {
	s.Title = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyMarkdown) SetText(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyMarkdown {
	s.Text = &v
	return s
}

type ServiceWindowMessageBatchPushRequestDetailMessageBodyLink struct {
	PicUrl     *string `json:"picUrl,omitempty" xml:"picUrl,omitempty"`
	MessageUrl *string `json:"messageUrl,omitempty" xml:"messageUrl,omitempty"`
	Title      *string `json:"title,omitempty" xml:"title,omitempty"`
	Text       *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s ServiceWindowMessageBatchPushRequestDetailMessageBodyLink) String() string {
	return tea.Prettify(s)
}

func (s ServiceWindowMessageBatchPushRequestDetailMessageBodyLink) GoString() string {
	return s.String()
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyLink) SetPicUrl(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyLink {
	s.PicUrl = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyLink) SetMessageUrl(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyLink {
	s.MessageUrl = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyLink) SetTitle(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyLink {
	s.Title = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyLink) SetText(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyLink {
	s.Text = &v
	return s
}

type ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard struct {
	ButtonOrientation *string                                                                      `json:"buttonOrientation,omitempty" xml:"buttonOrientation,omitempty"`
	SingleUrl         *string                                                                      `json:"singleUrl,omitempty" xml:"singleUrl,omitempty"`
	SingleTitle       *string                                                                      `json:"singleTitle,omitempty" xml:"singleTitle,omitempty"`
	Markdown          *string                                                                      `json:"markdown,omitempty" xml:"markdown,omitempty"`
	Title             *string                                                                      `json:"title,omitempty" xml:"title,omitempty"`
	ButtonList        []*ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCardButtonList `json:"buttonList,omitempty" xml:"buttonList,omitempty" type:"Repeated"`
}

func (s ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard) String() string {
	return tea.Prettify(s)
}

func (s ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard) GoString() string {
	return s.String()
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard) SetButtonOrientation(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard {
	s.ButtonOrientation = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard) SetSingleUrl(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard {
	s.SingleUrl = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard) SetSingleTitle(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard {
	s.SingleTitle = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard) SetMarkdown(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard {
	s.Markdown = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard) SetTitle(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard {
	s.Title = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard) SetButtonList(v []*ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCardButtonList) *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCard {
	s.ButtonList = v
	return s
}

type ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCardButtonList struct {
	Title     *string `json:"title,omitempty" xml:"title,omitempty"`
	ActionUrl *string `json:"actionUrl,omitempty" xml:"actionUrl,omitempty"`
}

func (s ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCardButtonList) String() string {
	return tea.Prettify(s)
}

func (s ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCardButtonList) GoString() string {
	return s.String()
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCardButtonList) SetTitle(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCardButtonList {
	s.Title = &v
	return s
}

func (s *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCardButtonList) SetActionUrl(v string) *ServiceWindowMessageBatchPushRequestDetailMessageBodyActionCardButtonList {
	s.ActionUrl = &v
	return s
}

type ServiceWindowMessageBatchPushResponseBody struct {
	// result
	Result    *ServiceWindowMessageBatchPushResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	RequestId *string                                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ServiceWindowMessageBatchPushResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ServiceWindowMessageBatchPushResponseBody) GoString() string {
	return s.String()
}

func (s *ServiceWindowMessageBatchPushResponseBody) SetResult(v *ServiceWindowMessageBatchPushResponseBodyResult) *ServiceWindowMessageBatchPushResponseBody {
	s.Result = v
	return s
}

func (s *ServiceWindowMessageBatchPushResponseBody) SetRequestId(v string) *ServiceWindowMessageBatchPushResponseBody {
	s.RequestId = &v
	return s
}

type ServiceWindowMessageBatchPushResponseBodyResult struct {
	OpenPushId *string `json:"openPushId,omitempty" xml:"openPushId,omitempty"`
}

func (s ServiceWindowMessageBatchPushResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ServiceWindowMessageBatchPushResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ServiceWindowMessageBatchPushResponseBodyResult) SetOpenPushId(v string) *ServiceWindowMessageBatchPushResponseBodyResult {
	s.OpenPushId = &v
	return s
}

type ServiceWindowMessageBatchPushResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ServiceWindowMessageBatchPushResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ServiceWindowMessageBatchPushResponse) String() string {
	return tea.Prettify(s)
}

func (s ServiceWindowMessageBatchPushResponse) GoString() string {
	return s.String()
}

func (s *ServiceWindowMessageBatchPushResponse) SetHeaders(v map[string]*string) *ServiceWindowMessageBatchPushResponse {
	s.Headers = v
	return s
}

func (s *ServiceWindowMessageBatchPushResponse) SetBody(v *ServiceWindowMessageBatchPushResponseBody) *ServiceWindowMessageBatchPushResponse {
	s.Body = v
	return s
}

type DeleteCrmFormInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteCrmFormInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteCrmFormInstanceHeaders) GoString() string {
	return s.String()
}

func (s *DeleteCrmFormInstanceHeaders) SetCommonHeaders(v map[string]*string) *DeleteCrmFormInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteCrmFormInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteCrmFormInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteCrmFormInstanceRequest struct {
	// 当前操作人id
	CurrentOperatorUserId *string `json:"currentOperatorUserId,omitempty" xml:"currentOperatorUserId,omitempty"`
	// 模版名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DeleteCrmFormInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCrmFormInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteCrmFormInstanceRequest) SetCurrentOperatorUserId(v string) *DeleteCrmFormInstanceRequest {
	s.CurrentOperatorUserId = &v
	return s
}

func (s *DeleteCrmFormInstanceRequest) SetName(v string) *DeleteCrmFormInstanceRequest {
	s.Name = &v
	return s
}

type DeleteCrmFormInstanceResponseBody struct {
	// 被删除的实例id
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s DeleteCrmFormInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCrmFormInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCrmFormInstanceResponseBody) SetInstanceId(v string) *DeleteCrmFormInstanceResponseBody {
	s.InstanceId = &v
	return s
}

type DeleteCrmFormInstanceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCrmFormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCrmFormInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCrmFormInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteCrmFormInstanceResponse) SetHeaders(v map[string]*string) *DeleteCrmFormInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteCrmFormInstanceResponse) SetBody(v *DeleteCrmFormInstanceResponseBody) *DeleteCrmFormInstanceResponse {
	s.Body = v
	return s
}

type BatchSendOfficialAccountOTOMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchSendOfficialAccountOTOMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOfficialAccountOTOMessageHeaders) GoString() string {
	return s.String()
}

func (s *BatchSendOfficialAccountOTOMessageHeaders) SetCommonHeaders(v map[string]*string) *BatchSendOfficialAccountOTOMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageHeaders) SetXAcsDingtalkAccessToken(v string) *BatchSendOfficialAccountOTOMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchSendOfficialAccountOTOMessageRequest struct {
	// 消息详情
	Detail *BatchSendOfficialAccountOTOMessageRequestDetail `json:"detail,omitempty" xml:"detail,omitempty" type:"Struct"`
	// 服务窗授权的调用方标识，可空
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// 服务窗帐号ID
	AccountId          *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
}

func (s BatchSendOfficialAccountOTOMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOfficialAccountOTOMessageRequest) GoString() string {
	return s.String()
}

func (s *BatchSendOfficialAccountOTOMessageRequest) SetDetail(v *BatchSendOfficialAccountOTOMessageRequestDetail) *BatchSendOfficialAccountOTOMessageRequest {
	s.Detail = v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequest) SetBizId(v string) *BatchSendOfficialAccountOTOMessageRequest {
	s.BizId = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequest) SetAccountId(v string) *BatchSendOfficialAccountOTOMessageRequest {
	s.AccountId = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequest) SetDingIsvOrgId(v int64) *BatchSendOfficialAccountOTOMessageRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequest) SetDingOrgId(v int64) *BatchSendOfficialAccountOTOMessageRequest {
	s.DingOrgId = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequest) SetDingTokenGrantType(v int64) *BatchSendOfficialAccountOTOMessageRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequest) SetDingSuiteKey(v string) *BatchSendOfficialAccountOTOMessageRequest {
	s.DingSuiteKey = &v
	return s
}

type BatchSendOfficialAccountOTOMessageRequestDetail struct {
	// 消息类型
	MsgType *string `json:"msgType,omitempty" xml:"msgType,omitempty"`
	// 消息请求唯一ID
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 业务请求标识，当一次业务请求需要多次调用发送API时可以设置此参数，方便后续跟踪处理。
	BizRequestId *string `json:"bizRequestId,omitempty" xml:"bizRequestId,omitempty"`
	// 消息接收人列表，最多支持1000人
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
	// 消息体
	MessageBody *BatchSendOfficialAccountOTOMessageRequestDetailMessageBody `json:"messageBody,omitempty" xml:"messageBody,omitempty" type:"Struct"`
	// 全员群发
	SendToAll *bool `json:"sendToAll,omitempty" xml:"sendToAll,omitempty"`
}

func (s BatchSendOfficialAccountOTOMessageRequestDetail) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOfficialAccountOTOMessageRequestDetail) GoString() string {
	return s.String()
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetail) SetMsgType(v string) *BatchSendOfficialAccountOTOMessageRequestDetail {
	s.MsgType = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetail) SetUuid(v string) *BatchSendOfficialAccountOTOMessageRequestDetail {
	s.Uuid = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetail) SetBizRequestId(v string) *BatchSendOfficialAccountOTOMessageRequestDetail {
	s.BizRequestId = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetail) SetUserIdList(v []*string) *BatchSendOfficialAccountOTOMessageRequestDetail {
	s.UserIdList = v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetail) SetMessageBody(v *BatchSendOfficialAccountOTOMessageRequestDetailMessageBody) *BatchSendOfficialAccountOTOMessageRequestDetail {
	s.MessageBody = v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetail) SetSendToAll(v bool) *BatchSendOfficialAccountOTOMessageRequestDetail {
	s.SendToAll = &v
	return s
}

type BatchSendOfficialAccountOTOMessageRequestDetailMessageBody struct {
	// 文本消息体  对于文本类型消息时必填
	Text *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyText `json:"text,omitempty" xml:"text,omitempty" type:"Struct"`
	// markdown消息，仅对消息类型为markdown时有效
	Markdown *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown `json:"markdown,omitempty" xml:"markdown,omitempty" type:"Struct"`
	// 链接消息类型
	Link *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyLink `json:"link,omitempty" xml:"link,omitempty" type:"Struct"`
	// 卡片消息
	ActionCard *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard `json:"actionCard,omitempty" xml:"actionCard,omitempty" type:"Struct"`
}

func (s BatchSendOfficialAccountOTOMessageRequestDetailMessageBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOfficialAccountOTOMessageRequestDetailMessageBody) GoString() string {
	return s.String()
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBody) SetText(v *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyText) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBody {
	s.Text = v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBody) SetMarkdown(v *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBody {
	s.Markdown = v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBody) SetLink(v *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyLink) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBody {
	s.Link = v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBody) SetActionCard(v *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBody {
	s.ActionCard = v
	return s
}

type BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyText struct {
	// 消息内容，建议500字符以内。
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyText) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyText) GoString() string {
	return s.String()
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyText) SetContent(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyText {
	s.Content = &v
	return s
}

type BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown struct {
	// 首屏会话透出的展示内容。
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// markdown格式的消息，建议500字符以内。
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown) GoString() string {
	return s.String()
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown) SetTitle(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown {
	s.Title = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown) SetText(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown {
	s.Text = &v
	return s
}

type BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyLink struct {
	// 图片地址
	PicUrl *string `json:"picUrl,omitempty" xml:"picUrl,omitempty"`
	// 消息点击链接地址，当发送消息为小程序时支持小程序跳转链接。
	MessageUrl *string `json:"messageUrl,omitempty" xml:"messageUrl,omitempty"`
	// 消息标题，建议100字符以内。
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 消息描述，建议500字符以内。
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyLink) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyLink) GoString() string {
	return s.String()
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyLink) SetPicUrl(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyLink {
	s.PicUrl = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyLink) SetMessageUrl(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyLink {
	s.MessageUrl = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyLink) SetTitle(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyLink {
	s.Title = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyLink) SetText(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyLink {
	s.Text = &v
	return s
}

type BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard struct {
	// 按钮排列方式： 0：竖直排列 1：横向排列 必须与buttonList同时设置。
	ButtonOrientation *string `json:"buttonOrientation,omitempty" xml:"buttonOrientation,omitempty"`
	// 消息点击链接地址，当发送消息为小程序时支持小程序跳转链接，最长500个字符。
	SingleUrl *string `json:"singleUrl,omitempty" xml:"singleUrl,omitempty"`
	// 使用整体跳转ActionCard样式时的标题。必须与singleUrl同时设置，最长20个字符。
	SingleTitle *string `json:"singleTitle,omitempty" xml:"singleTitle,omitempty"`
	// 消息内容，支持markdown，语法参考标准markdown语法。1000个字符以内。
	Markdown *string `json:"markdown,omitempty" xml:"markdown,omitempty"`
	// 透出到会话列表和通知的文案
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 使用独立跳转ActionCard样式时的按钮列表；必须与buttonOrientation同时设置，且长度不超过1000字符。
	ButtonList []*BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList `json:"buttonList,omitempty" xml:"buttonList,omitempty" type:"Repeated"`
}

func (s BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) GoString() string {
	return s.String()
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) SetButtonOrientation(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard {
	s.ButtonOrientation = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) SetSingleUrl(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard {
	s.SingleUrl = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) SetSingleTitle(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard {
	s.SingleTitle = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) SetMarkdown(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard {
	s.Markdown = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) SetTitle(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard {
	s.Title = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) SetButtonList(v []*BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard {
	s.ButtonList = v
	return s
}

type BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList struct {
	// 使用独立跳转ActionCard样式时的按钮的标题，最长20个字符。
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 使用独立跳转ActionCard样式时的跳转链接。
	ActionUrl *string `json:"actionUrl,omitempty" xml:"actionUrl,omitempty"`
}

func (s BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList) GoString() string {
	return s.String()
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList) SetTitle(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList {
	s.Title = &v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList) SetActionUrl(v string) *BatchSendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList {
	s.ActionUrl = &v
	return s
}

type BatchSendOfficialAccountOTOMessageResponseBody struct {
	// result
	Result *BatchSendOfficialAccountOTOMessageResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// 开放API
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s BatchSendOfficialAccountOTOMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOfficialAccountOTOMessageResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSendOfficialAccountOTOMessageResponseBody) SetResult(v *BatchSendOfficialAccountOTOMessageResponseBodyResult) *BatchSendOfficialAccountOTOMessageResponseBody {
	s.Result = v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageResponseBody) SetRequestId(v string) *BatchSendOfficialAccountOTOMessageResponseBody {
	s.RequestId = &v
	return s
}

type BatchSendOfficialAccountOTOMessageResponseBodyResult struct {
	OpenPushId *string `json:"openPushId,omitempty" xml:"openPushId,omitempty"`
}

func (s BatchSendOfficialAccountOTOMessageResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOfficialAccountOTOMessageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *BatchSendOfficialAccountOTOMessageResponseBodyResult) SetOpenPushId(v string) *BatchSendOfficialAccountOTOMessageResponseBodyResult {
	s.OpenPushId = &v
	return s
}

type BatchSendOfficialAccountOTOMessageResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchSendOfficialAccountOTOMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchSendOfficialAccountOTOMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchSendOfficialAccountOTOMessageResponse) GoString() string {
	return s.String()
}

func (s *BatchSendOfficialAccountOTOMessageResponse) SetHeaders(v map[string]*string) *BatchSendOfficialAccountOTOMessageResponse {
	s.Headers = v
	return s
}

func (s *BatchSendOfficialAccountOTOMessageResponse) SetBody(v *BatchSendOfficialAccountOTOMessageResponseBody) *BatchSendOfficialAccountOTOMessageResponse {
	s.Body = v
	return s
}

type GetOfficialAccountContactInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOfficialAccountContactInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountContactInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountContactInfoHeaders) SetCommonHeaders(v map[string]*string) *GetOfficialAccountContactInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOfficialAccountContactInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetOfficialAccountContactInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOfficialAccountContactInfoResponseBody struct {
	// 联系人主企业名称
	CorpName *string `json:"corpName,omitempty" xml:"corpName,omitempty"`
	// 手机号
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 手机号国家码
	StateCode *string `json:"stateCode,omitempty" xml:"stateCode,omitempty"`
	// 联系人的unionId
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// 已授权的字段
	AuthItems []*string `json:"authItems,omitempty" xml:"authItems,omitempty" type:"Repeated"`
	// 已授权的字段
	UserInfos []*string `json:"userInfos,omitempty" xml:"userInfos,omitempty" type:"Repeated"`
}

func (s GetOfficialAccountContactInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountContactInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountContactInfoResponseBody) SetCorpName(v string) *GetOfficialAccountContactInfoResponseBody {
	s.CorpName = &v
	return s
}

func (s *GetOfficialAccountContactInfoResponseBody) SetMobile(v string) *GetOfficialAccountContactInfoResponseBody {
	s.Mobile = &v
	return s
}

func (s *GetOfficialAccountContactInfoResponseBody) SetStateCode(v string) *GetOfficialAccountContactInfoResponseBody {
	s.StateCode = &v
	return s
}

func (s *GetOfficialAccountContactInfoResponseBody) SetUnionId(v string) *GetOfficialAccountContactInfoResponseBody {
	s.UnionId = &v
	return s
}

func (s *GetOfficialAccountContactInfoResponseBody) SetAuthItems(v []*string) *GetOfficialAccountContactInfoResponseBody {
	s.AuthItems = v
	return s
}

func (s *GetOfficialAccountContactInfoResponseBody) SetUserInfos(v []*string) *GetOfficialAccountContactInfoResponseBody {
	s.UserInfos = v
	return s
}

type GetOfficialAccountContactInfoResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOfficialAccountContactInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOfficialAccountContactInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountContactInfoResponse) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountContactInfoResponse) SetHeaders(v map[string]*string) *GetOfficialAccountContactInfoResponse {
	s.Headers = v
	return s
}

func (s *GetOfficialAccountContactInfoResponse) SetBody(v *GetOfficialAccountContactInfoResponseBody) *GetOfficialAccountContactInfoResponse {
	s.Body = v
	return s
}

type QueryAllCustomerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllCustomerHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCustomerHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllCustomerHeaders) SetCommonHeaders(v map[string]*string) *QueryAllCustomerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllCustomerHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllCustomerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllCustomerRequest struct {
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	DingCorpId         *string `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	// 用户ID
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// 翻页size
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 分页游标，第一次调用传空或者null
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 数据类型
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s QueryAllCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCustomerRequest) GoString() string {
	return s.String()
}

func (s *QueryAllCustomerRequest) SetDingIsvOrgId(v int64) *QueryAllCustomerRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *QueryAllCustomerRequest) SetDingOrgId(v int64) *QueryAllCustomerRequest {
	s.DingOrgId = &v
	return s
}

func (s *QueryAllCustomerRequest) SetDingTokenGrantType(v int64) *QueryAllCustomerRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *QueryAllCustomerRequest) SetDingCorpId(v string) *QueryAllCustomerRequest {
	s.DingCorpId = &v
	return s
}

func (s *QueryAllCustomerRequest) SetDingSuiteKey(v string) *QueryAllCustomerRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *QueryAllCustomerRequest) SetOperatorUserId(v string) *QueryAllCustomerRequest {
	s.OperatorUserId = &v
	return s
}

func (s *QueryAllCustomerRequest) SetMaxResults(v int64) *QueryAllCustomerRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryAllCustomerRequest) SetNextToken(v string) *QueryAllCustomerRequest {
	s.NextToken = &v
	return s
}

func (s *QueryAllCustomerRequest) SetObjectType(v string) *QueryAllCustomerRequest {
	s.ObjectType = &v
	return s
}

type QueryAllCustomerResponseBody struct {
	// 分页结果
	Result *QueryAllCustomerResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryAllCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllCustomerResponseBody) SetResult(v *QueryAllCustomerResponseBodyResult) *QueryAllCustomerResponseBody {
	s.Result = v
	return s
}

type QueryAllCustomerResponseBodyResult struct {
	// 下一页的游标，为null则表示无数据
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 客户数据节点
	Values []*QueryAllCustomerResponseBodyResultValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
	// 分页大小
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
}

func (s QueryAllCustomerResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCustomerResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryAllCustomerResponseBodyResult) SetNextToken(v string) *QueryAllCustomerResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResult) SetValues(v []*QueryAllCustomerResponseBodyResultValues) *QueryAllCustomerResponseBodyResult {
	s.Values = v
	return s
}

func (s *QueryAllCustomerResponseBodyResult) SetMaxResults(v int64) *QueryAllCustomerResponseBodyResult {
	s.MaxResults = &v
	return s
}

type QueryAllCustomerResponseBodyResultValues struct {
	// 创建记录的用户昵称
	CreatorNick *string `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	// 记录修改时间
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// 创建记录的用户ID
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// 数据ID
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// 数据内容
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// 扩展数据内容
	ExtendData map[string]interface{} `json:"extendData,omitempty" xml:"extendData,omitempty"`
	// 记录创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 系统自动生成
	OrgId *int64 `json:"orgId,omitempty" xml:"orgId,omitempty"`
	// 数据类型
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// 数据权限信息
	Permission *QueryAllCustomerResponseBodyResultValuesPermission `json:"permission,omitempty" xml:"permission,omitempty" type:"Struct"`
	// 审批结果
	ProcessOutResult *string `json:"processOutResult,omitempty" xml:"processOutResult,omitempty"`
	// 审批状态
	ProcessInstanceStatus *string `json:"processInstanceStatus,omitempty" xml:"processInstanceStatus,omitempty"`
}

func (s QueryAllCustomerResponseBodyResultValues) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCustomerResponseBodyResultValues) GoString() string {
	return s.String()
}

func (s *QueryAllCustomerResponseBodyResultValues) SetCreatorNick(v string) *QueryAllCustomerResponseBodyResultValues {
	s.CreatorNick = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetModifyTime(v string) *QueryAllCustomerResponseBodyResultValues {
	s.ModifyTime = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetCreatorUserId(v string) *QueryAllCustomerResponseBodyResultValues {
	s.CreatorUserId = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetInstanceId(v string) *QueryAllCustomerResponseBodyResultValues {
	s.InstanceId = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetData(v map[string]interface{}) *QueryAllCustomerResponseBodyResultValues {
	s.Data = v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetExtendData(v map[string]interface{}) *QueryAllCustomerResponseBodyResultValues {
	s.ExtendData = v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetCreateTime(v string) *QueryAllCustomerResponseBodyResultValues {
	s.CreateTime = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetOrgId(v int64) *QueryAllCustomerResponseBodyResultValues {
	s.OrgId = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetObjectType(v string) *QueryAllCustomerResponseBodyResultValues {
	s.ObjectType = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetPermission(v *QueryAllCustomerResponseBodyResultValuesPermission) *QueryAllCustomerResponseBodyResultValues {
	s.Permission = v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetProcessOutResult(v string) *QueryAllCustomerResponseBodyResultValues {
	s.ProcessOutResult = &v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValues) SetProcessInstanceStatus(v string) *QueryAllCustomerResponseBodyResultValues {
	s.ProcessInstanceStatus = &v
	return s
}

type QueryAllCustomerResponseBodyResultValuesPermission struct {
	// 协同人用户ID列表
	ParticipantStaffIds []*string `json:"participantStaffIds,omitempty" xml:"participantStaffIds,omitempty" type:"Repeated"`
	// 负责人用户ID列表
	OwnerStaffIds []*string `json:"ownerStaffIds,omitempty" xml:"ownerStaffIds,omitempty" type:"Repeated"`
}

func (s QueryAllCustomerResponseBodyResultValuesPermission) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCustomerResponseBodyResultValuesPermission) GoString() string {
	return s.String()
}

func (s *QueryAllCustomerResponseBodyResultValuesPermission) SetParticipantStaffIds(v []*string) *QueryAllCustomerResponseBodyResultValuesPermission {
	s.ParticipantStaffIds = v
	return s
}

func (s *QueryAllCustomerResponseBodyResultValuesPermission) SetOwnerStaffIds(v []*string) *QueryAllCustomerResponseBodyResultValuesPermission {
	s.OwnerStaffIds = v
	return s
}

type QueryAllCustomerResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAllCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCustomerResponse) GoString() string {
	return s.String()
}

func (s *QueryAllCustomerResponse) SetHeaders(v map[string]*string) *QueryAllCustomerResponse {
	s.Headers = v
	return s
}

func (s *QueryAllCustomerResponse) SetBody(v *QueryAllCustomerResponseBody) *QueryAllCustomerResponse {
	s.Body = v
	return s
}

type SendOfficialAccountOTOMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendOfficialAccountOTOMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountOTOMessageHeaders) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountOTOMessageHeaders) SetCommonHeaders(v map[string]*string) *SendOfficialAccountOTOMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendOfficialAccountOTOMessageHeaders) SetXAcsDingtalkAccessToken(v string) *SendOfficialAccountOTOMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendOfficialAccountOTOMessageRequest struct {
	// 消息详情
	Detail *SendOfficialAccountOTOMessageRequestDetail `json:"detail,omitempty" xml:"detail,omitempty" type:"Struct"`
	// API调用标识，可选参数
	BizId              *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	// 服务窗帐号ID
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s SendOfficialAccountOTOMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountOTOMessageRequest) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountOTOMessageRequest) SetDetail(v *SendOfficialAccountOTOMessageRequestDetail) *SendOfficialAccountOTOMessageRequest {
	s.Detail = v
	return s
}

func (s *SendOfficialAccountOTOMessageRequest) SetBizId(v string) *SendOfficialAccountOTOMessageRequest {
	s.BizId = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequest) SetDingTokenGrantType(v int64) *SendOfficialAccountOTOMessageRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequest) SetDingIsvOrgId(v int64) *SendOfficialAccountOTOMessageRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequest) SetDingOrgId(v int64) *SendOfficialAccountOTOMessageRequest {
	s.DingOrgId = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequest) SetDingSuiteKey(v string) *SendOfficialAccountOTOMessageRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequest) SetAccountId(v string) *SendOfficialAccountOTOMessageRequest {
	s.AccountId = &v
	return s
}

type SendOfficialAccountOTOMessageRequestDetail struct {
	// 消息类型
	MsgType *string `json:"msgType,omitempty" xml:"msgType,omitempty"`
	// 请求唯一 ID
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 消息接收人id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 消息体
	MessageBody *SendOfficialAccountOTOMessageRequestDetailMessageBody `json:"messageBody,omitempty" xml:"messageBody,omitempty" type:"Struct"`
}

func (s SendOfficialAccountOTOMessageRequestDetail) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountOTOMessageRequestDetail) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountOTOMessageRequestDetail) SetMsgType(v string) *SendOfficialAccountOTOMessageRequestDetail {
	s.MsgType = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetail) SetUuid(v string) *SendOfficialAccountOTOMessageRequestDetail {
	s.Uuid = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetail) SetUserId(v string) *SendOfficialAccountOTOMessageRequestDetail {
	s.UserId = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetail) SetMessageBody(v *SendOfficialAccountOTOMessageRequestDetailMessageBody) *SendOfficialAccountOTOMessageRequestDetail {
	s.MessageBody = v
	return s
}

type SendOfficialAccountOTOMessageRequestDetailMessageBody struct {
	// 文本消息体  对于文本类型消息时必填
	Text *SendOfficialAccountOTOMessageRequestDetailMessageBodyText `json:"text,omitempty" xml:"text,omitempty" type:"Struct"`
	// markdown消息，仅对消息类型为markdown时有效
	Markdown *SendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown `json:"markdown,omitempty" xml:"markdown,omitempty" type:"Struct"`
	// 链接消息类型
	Link *SendOfficialAccountOTOMessageRequestDetailMessageBodyLink `json:"link,omitempty" xml:"link,omitempty" type:"Struct"`
	// 卡片消息
	ActionCard *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard `json:"actionCard,omitempty" xml:"actionCard,omitempty" type:"Struct"`
}

func (s SendOfficialAccountOTOMessageRequestDetailMessageBody) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountOTOMessageRequestDetailMessageBody) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBody) SetText(v *SendOfficialAccountOTOMessageRequestDetailMessageBodyText) *SendOfficialAccountOTOMessageRequestDetailMessageBody {
	s.Text = v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBody) SetMarkdown(v *SendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown) *SendOfficialAccountOTOMessageRequestDetailMessageBody {
	s.Markdown = v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBody) SetLink(v *SendOfficialAccountOTOMessageRequestDetailMessageBodyLink) *SendOfficialAccountOTOMessageRequestDetailMessageBody {
	s.Link = v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBody) SetActionCard(v *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) *SendOfficialAccountOTOMessageRequestDetailMessageBody {
	s.ActionCard = v
	return s
}

type SendOfficialAccountOTOMessageRequestDetailMessageBodyText struct {
	// 消息内容，建议500字符以内。
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s SendOfficialAccountOTOMessageRequestDetailMessageBodyText) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountOTOMessageRequestDetailMessageBodyText) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyText) SetContent(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyText {
	s.Content = &v
	return s
}

type SendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown struct {
	// 首屏会话透出的展示内容。
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// markdown格式的消息，建议500字符以内。
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s SendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown) SetTitle(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown {
	s.Title = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown) SetText(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyMarkdown {
	s.Text = &v
	return s
}

type SendOfficialAccountOTOMessageRequestDetailMessageBodyLink struct {
	// 图片地址
	PicUrl *string `json:"picUrl,omitempty" xml:"picUrl,omitempty"`
	// 消息点击链接地址，当发送消息为小程序时支持小程序跳转链接。
	MessageUrl *string `json:"messageUrl,omitempty" xml:"messageUrl,omitempty"`
	// 消息标题，建议100字符以内。
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 消息描述，建议500字符以内。
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s SendOfficialAccountOTOMessageRequestDetailMessageBodyLink) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountOTOMessageRequestDetailMessageBodyLink) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyLink) SetPicUrl(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyLink {
	s.PicUrl = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyLink) SetMessageUrl(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyLink {
	s.MessageUrl = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyLink) SetTitle(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyLink {
	s.Title = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyLink) SetText(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyLink {
	s.Text = &v
	return s
}

type SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard struct {
	// 按钮排列方式： 0：竖直排列 1：横向排列 必须与buttonList同时设置。
	ButtonOrientation *string `json:"buttonOrientation,omitempty" xml:"buttonOrientation,omitempty"`
	// 消息点击链接地址，当发送消息为小程序时支持小程序跳转链接，最长500个字符。
	SingleUrl *string `json:"singleUrl,omitempty" xml:"singleUrl,omitempty"`
	// 使用整体跳转ActionCard样式时的标题。必须与singleUrl同时设置，最长20个字符。
	SingleTitle *string `json:"singleTitle,omitempty" xml:"singleTitle,omitempty"`
	// 消息内容，支持markdown，语法参考标准markdown语法。1000个字符以内。
	Markdown *string `json:"markdown,omitempty" xml:"markdown,omitempty"`
	// 透出到会话列表和通知的文案
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 使用独立跳转ActionCard样式时的按钮列表；必须与buttonOrientation同时设置，且长度不超过1000字符。
	ButtonList []*SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList `json:"buttonList,omitempty" xml:"buttonList,omitempty" type:"Repeated"`
}

func (s SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) SetButtonOrientation(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard {
	s.ButtonOrientation = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) SetSingleUrl(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard {
	s.SingleUrl = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) SetSingleTitle(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard {
	s.SingleTitle = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) SetMarkdown(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard {
	s.Markdown = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) SetTitle(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard {
	s.Title = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard) SetButtonList(v []*SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList) *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCard {
	s.ButtonList = v
	return s
}

type SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList struct {
	// 使用独立跳转ActionCard样式时的按钮的标题，最长20个字符。
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 使用独立跳转ActionCard样式时的跳转链接。
	ActionUrl *string `json:"actionUrl,omitempty" xml:"actionUrl,omitempty"`
}

func (s SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList) SetTitle(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList {
	s.Title = &v
	return s
}

func (s *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList) SetActionUrl(v string) *SendOfficialAccountOTOMessageRequestDetailMessageBodyActionCardButtonList {
	s.ActionUrl = &v
	return s
}

type SendOfficialAccountOTOMessageResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 推送结果
	Result *SendOfficialAccountOTOMessageResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s SendOfficialAccountOTOMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountOTOMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountOTOMessageResponseBody) SetRequestId(v string) *SendOfficialAccountOTOMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendOfficialAccountOTOMessageResponseBody) SetResult(v *SendOfficialAccountOTOMessageResponseBodyResult) *SendOfficialAccountOTOMessageResponseBody {
	s.Result = v
	return s
}

type SendOfficialAccountOTOMessageResponseBodyResult struct {
	// 推送ID
	OpenPushId *string `json:"openPushId,omitempty" xml:"openPushId,omitempty"`
}

func (s SendOfficialAccountOTOMessageResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountOTOMessageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountOTOMessageResponseBodyResult) SetOpenPushId(v string) *SendOfficialAccountOTOMessageResponseBodyResult {
	s.OpenPushId = &v
	return s
}

type SendOfficialAccountOTOMessageResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendOfficialAccountOTOMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendOfficialAccountOTOMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendOfficialAccountOTOMessageResponse) GoString() string {
	return s.String()
}

func (s *SendOfficialAccountOTOMessageResponse) SetHeaders(v map[string]*string) *SendOfficialAccountOTOMessageResponse {
	s.Headers = v
	return s
}

func (s *SendOfficialAccountOTOMessageResponse) SetBody(v *SendOfficialAccountOTOMessageResponseBody) *SendOfficialAccountOTOMessageResponse {
	s.Body = v
	return s
}

type GetOfficialAccountOTOMessageResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOfficialAccountOTOMessageResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountOTOMessageResultHeaders) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountOTOMessageResultHeaders) SetCommonHeaders(v map[string]*string) *GetOfficialAccountOTOMessageResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOfficialAccountOTOMessageResultHeaders) SetXAcsDingtalkAccessToken(v string) *GetOfficialAccountOTOMessageResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOfficialAccountOTOMessageResultRequest struct {
	// 推送ID
	OpenPushId *string `json:"openPushId,omitempty" xml:"openPushId,omitempty"`
	AccountId  *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetOfficialAccountOTOMessageResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountOTOMessageResultRequest) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountOTOMessageResultRequest) SetOpenPushId(v string) *GetOfficialAccountOTOMessageResultRequest {
	s.OpenPushId = &v
	return s
}

func (s *GetOfficialAccountOTOMessageResultRequest) SetAccountId(v string) *GetOfficialAccountOTOMessageResultRequest {
	s.AccountId = &v
	return s
}

type GetOfficialAccountOTOMessageResultResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 查询结果
	Result *GetOfficialAccountOTOMessageResultResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetOfficialAccountOTOMessageResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountOTOMessageResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountOTOMessageResultResponseBody) SetRequestId(v string) *GetOfficialAccountOTOMessageResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOfficialAccountOTOMessageResultResponseBody) SetResult(v *GetOfficialAccountOTOMessageResultResponseBodyResult) *GetOfficialAccountOTOMessageResultResponseBody {
	s.Result = v
	return s
}

type GetOfficialAccountOTOMessageResultResponseBodyResult struct {
	// 执行状态： 0：未开始  1：处理中  2：处理完毕
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 已读消息的userid列表
	ReadUserIdList []*string `json:"readUserIdList,omitempty" xml:"readUserIdList,omitempty" type:"Repeated"`
}

func (s GetOfficialAccountOTOMessageResultResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountOTOMessageResultResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountOTOMessageResultResponseBodyResult) SetStatus(v int64) *GetOfficialAccountOTOMessageResultResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetOfficialAccountOTOMessageResultResponseBodyResult) SetReadUserIdList(v []*string) *GetOfficialAccountOTOMessageResultResponseBodyResult {
	s.ReadUserIdList = v
	return s
}

type GetOfficialAccountOTOMessageResultResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOfficialAccountOTOMessageResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOfficialAccountOTOMessageResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOfficialAccountOTOMessageResultResponse) GoString() string {
	return s.String()
}

func (s *GetOfficialAccountOTOMessageResultResponse) SetHeaders(v map[string]*string) *GetOfficialAccountOTOMessageResultResponse {
	s.Headers = v
	return s
}

func (s *GetOfficialAccountOTOMessageResultResponse) SetBody(v *GetOfficialAccountOTOMessageResultResponseBody) *GetOfficialAccountOTOMessageResultResponse {
	s.Body = v
	return s
}

type AddCrmPersonalCustomerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddCrmPersonalCustomerHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddCrmPersonalCustomerHeaders) GoString() string {
	return s.String()
}

func (s *AddCrmPersonalCustomerHeaders) SetCommonHeaders(v map[string]*string) *AddCrmPersonalCustomerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddCrmPersonalCustomerHeaders) SetXAcsDingtalkAccessToken(v string) *AddCrmPersonalCustomerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddCrmPersonalCustomerRequest struct {
	// 记录创建人的用户ID
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// 记录创建人的昵称
	CreatorNick *string `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	// 数据内容
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// 扩展数据内容
	ExtendData map[string]interface{} `json:"extendData,omitempty" xml:"extendData,omitempty"`
	// 权限
	Permission *AddCrmPersonalCustomerRequestPermission `json:"permission,omitempty" xml:"permission,omitempty" type:"Struct"`
	// 跳过uk查重
	SkipDuplicateCheck *bool `json:"skipDuplicateCheck,omitempty" xml:"skipDuplicateCheck,omitempty"`
}

func (s AddCrmPersonalCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCrmPersonalCustomerRequest) GoString() string {
	return s.String()
}

func (s *AddCrmPersonalCustomerRequest) SetCreatorUserId(v string) *AddCrmPersonalCustomerRequest {
	s.CreatorUserId = &v
	return s
}

func (s *AddCrmPersonalCustomerRequest) SetCreatorNick(v string) *AddCrmPersonalCustomerRequest {
	s.CreatorNick = &v
	return s
}

func (s *AddCrmPersonalCustomerRequest) SetData(v map[string]interface{}) *AddCrmPersonalCustomerRequest {
	s.Data = v
	return s
}

func (s *AddCrmPersonalCustomerRequest) SetExtendData(v map[string]interface{}) *AddCrmPersonalCustomerRequest {
	s.ExtendData = v
	return s
}

func (s *AddCrmPersonalCustomerRequest) SetPermission(v *AddCrmPersonalCustomerRequestPermission) *AddCrmPersonalCustomerRequest {
	s.Permission = v
	return s
}

func (s *AddCrmPersonalCustomerRequest) SetSkipDuplicateCheck(v bool) *AddCrmPersonalCustomerRequest {
	s.SkipDuplicateCheck = &v
	return s
}

type AddCrmPersonalCustomerRequestPermission struct {
	// 负责人的用户ID
	OwnerStaffIds []*string `json:"ownerStaffIds,omitempty" xml:"ownerStaffIds,omitempty" type:"Repeated"`
	// 协同人的用户ID
	ParticipantStaffIds []*string `json:"participantStaffIds,omitempty" xml:"participantStaffIds,omitempty" type:"Repeated"`
}

func (s AddCrmPersonalCustomerRequestPermission) String() string {
	return tea.Prettify(s)
}

func (s AddCrmPersonalCustomerRequestPermission) GoString() string {
	return s.String()
}

func (s *AddCrmPersonalCustomerRequestPermission) SetOwnerStaffIds(v []*string) *AddCrmPersonalCustomerRequestPermission {
	s.OwnerStaffIds = v
	return s
}

func (s *AddCrmPersonalCustomerRequestPermission) SetParticipantStaffIds(v []*string) *AddCrmPersonalCustomerRequestPermission {
	s.ParticipantStaffIds = v
	return s
}

type AddCrmPersonalCustomerResponseBody struct {
	// 客户数据id
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s AddCrmPersonalCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCrmPersonalCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *AddCrmPersonalCustomerResponseBody) SetInstanceId(v string) *AddCrmPersonalCustomerResponseBody {
	s.InstanceId = &v
	return s
}

type AddCrmPersonalCustomerResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddCrmPersonalCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddCrmPersonalCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCrmPersonalCustomerResponse) GoString() string {
	return s.String()
}

func (s *AddCrmPersonalCustomerResponse) SetHeaders(v map[string]*string) *AddCrmPersonalCustomerResponse {
	s.Headers = v
	return s
}

func (s *AddCrmPersonalCustomerResponse) SetBody(v *AddCrmPersonalCustomerResponseBody) *AddCrmPersonalCustomerResponse {
	s.Body = v
	return s
}

type RecallOfficialAccountOTOMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RecallOfficialAccountOTOMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s RecallOfficialAccountOTOMessageHeaders) GoString() string {
	return s.String()
}

func (s *RecallOfficialAccountOTOMessageHeaders) SetCommonHeaders(v map[string]*string) *RecallOfficialAccountOTOMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RecallOfficialAccountOTOMessageHeaders) SetXAcsDingtalkAccessToken(v string) *RecallOfficialAccountOTOMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RecallOfficialAccountOTOMessageRequest struct {
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	// 帐号ID 可空
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// 消息推送时返回的ID
	OpenPushId *string `json:"openPushId,omitempty" xml:"openPushId,omitempty"`
}

func (s RecallOfficialAccountOTOMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s RecallOfficialAccountOTOMessageRequest) GoString() string {
	return s.String()
}

func (s *RecallOfficialAccountOTOMessageRequest) SetDingSuiteKey(v string) *RecallOfficialAccountOTOMessageRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *RecallOfficialAccountOTOMessageRequest) SetDingOrgId(v int64) *RecallOfficialAccountOTOMessageRequest {
	s.DingOrgId = &v
	return s
}

func (s *RecallOfficialAccountOTOMessageRequest) SetDingIsvOrgId(v int64) *RecallOfficialAccountOTOMessageRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *RecallOfficialAccountOTOMessageRequest) SetDingTokenGrantType(v int64) *RecallOfficialAccountOTOMessageRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *RecallOfficialAccountOTOMessageRequest) SetAccountId(v string) *RecallOfficialAccountOTOMessageRequest {
	s.AccountId = &v
	return s
}

func (s *RecallOfficialAccountOTOMessageRequest) SetOpenPushId(v string) *RecallOfficialAccountOTOMessageRequest {
	s.OpenPushId = &v
	return s
}

type RecallOfficialAccountOTOMessageResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RecallOfficialAccountOTOMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecallOfficialAccountOTOMessageResponseBody) GoString() string {
	return s.String()
}

func (s *RecallOfficialAccountOTOMessageResponseBody) SetRequestId(v string) *RecallOfficialAccountOTOMessageResponseBody {
	s.RequestId = &v
	return s
}

type RecallOfficialAccountOTOMessageResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecallOfficialAccountOTOMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecallOfficialAccountOTOMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s RecallOfficialAccountOTOMessageResponse) GoString() string {
	return s.String()
}

func (s *RecallOfficialAccountOTOMessageResponse) SetHeaders(v map[string]*string) *RecallOfficialAccountOTOMessageResponse {
	s.Headers = v
	return s
}

func (s *RecallOfficialAccountOTOMessageResponse) SetBody(v *RecallOfficialAccountOTOMessageResponseBody) *RecallOfficialAccountOTOMessageResponse {
	s.Body = v
	return s
}

type DescribeCrmPersonalCustomerObjectMetaHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DescribeCrmPersonalCustomerObjectMetaHeaders) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrmPersonalCustomerObjectMetaHeaders) GoString() string {
	return s.String()
}

func (s *DescribeCrmPersonalCustomerObjectMetaHeaders) SetCommonHeaders(v map[string]*string) *DescribeCrmPersonalCustomerObjectMetaHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaHeaders) SetXAcsDingtalkAccessToken(v string) *DescribeCrmPersonalCustomerObjectMetaHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DescribeCrmPersonalCustomerObjectMetaResponseBody struct {
	// 对象名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 是否自定义对象
	Customized *bool `json:"customized,omitempty" xml:"customized,omitempty"`
	// 字段列表
	Fields []*DescribeCrmPersonalCustomerObjectMetaResponseBodyFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
}

func (s DescribeCrmPersonalCustomerObjectMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrmPersonalCustomerObjectMetaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBody) SetName(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBody) SetCustomized(v bool) *DescribeCrmPersonalCustomerObjectMetaResponseBody {
	s.Customized = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBody) SetFields(v []*DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) *DescribeCrmPersonalCustomerObjectMetaResponseBody {
	s.Fields = v
	return s
}

type DescribeCrmPersonalCustomerObjectMetaResponseBodyFields struct {
	Name                *string                                                                       `json:"name,omitempty" xml:"name,omitempty"`
	Customized          *bool                                                                         `json:"customized,omitempty" xml:"customized,omitempty"`
	Label               *string                                                                       `json:"label,omitempty" xml:"label,omitempty"`
	Type                *string                                                                       `json:"type,omitempty" xml:"type,omitempty"`
	Nillable            *bool                                                                         `json:"nillable,omitempty" xml:"nillable,omitempty"`
	Format              *string                                                                       `json:"format,omitempty" xml:"format,omitempty"`
	Unit                *string                                                                       `json:"unit,omitempty" xml:"unit,omitempty"`
	SelectOptions       []*DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsSelectOptions       `json:"selectOptions,omitempty" xml:"selectOptions,omitempty" type:"Repeated"`
	Quote               *bool                                                                         `json:"quote,omitempty" xml:"quote,omitempty"`
	ReferenceTo         *string                                                                       `json:"referenceTo,omitempty" xml:"referenceTo,omitempty"`
	ReferenceFields     []*DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields     `json:"referenceFields,omitempty" xml:"referenceFields,omitempty" type:"Repeated"`
	RollUpSummaryFields []*DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsRollUpSummaryFields `json:"rollUpSummaryFields,omitempty" xml:"rollUpSummaryFields,omitempty" type:"Repeated"`
}

func (s DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) GoString() string {
	return s.String()
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) SetName(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields {
	s.Name = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) SetCustomized(v bool) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields {
	s.Customized = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) SetLabel(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields {
	s.Label = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) SetType(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields {
	s.Type = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) SetNillable(v bool) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields {
	s.Nillable = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) SetFormat(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields {
	s.Format = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) SetUnit(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields {
	s.Unit = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) SetSelectOptions(v []*DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsSelectOptions) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields {
	s.SelectOptions = v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) SetQuote(v bool) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields {
	s.Quote = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) SetReferenceTo(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields {
	s.ReferenceTo = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) SetReferenceFields(v []*DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields {
	s.ReferenceFields = v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields) SetRollUpSummaryFields(v []*DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsRollUpSummaryFields) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFields {
	s.RollUpSummaryFields = v
	return s
}

type DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsSelectOptions struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsSelectOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsSelectOptions) GoString() string {
	return s.String()
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsSelectOptions) SetKey(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsSelectOptions {
	s.Key = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsSelectOptions) SetValue(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsSelectOptions {
	s.Value = &v
	return s
}

type DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields struct {
	Label         *string                                                                                `json:"label,omitempty" xml:"label,omitempty"`
	Type          *string                                                                                `json:"type,omitempty" xml:"type,omitempty"`
	Nillable      *bool                                                                                  `json:"nillable,omitempty" xml:"nillable,omitempty"`
	Unit          *string                                                                                `json:"unit,omitempty" xml:"unit,omitempty"`
	Format        *string                                                                                `json:"format,omitempty" xml:"format,omitempty"`
	SelectOptions []*DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFieldsSelectOptions `json:"selectOptions,omitempty" xml:"selectOptions,omitempty" type:"Repeated"`
	Name          *string                                                                                `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields) GoString() string {
	return s.String()
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields) SetLabel(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields {
	s.Label = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields) SetType(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields {
	s.Type = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields) SetNillable(v bool) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields {
	s.Nillable = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields) SetUnit(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields {
	s.Unit = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields) SetFormat(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields {
	s.Format = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields) SetSelectOptions(v []*DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFieldsSelectOptions) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields {
	s.SelectOptions = v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields) SetName(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFields {
	s.Name = &v
	return s
}

type DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFieldsSelectOptions struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFieldsSelectOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFieldsSelectOptions) GoString() string {
	return s.String()
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFieldsSelectOptions) SetKey(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFieldsSelectOptions {
	s.Key = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFieldsSelectOptions) SetValue(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsReferenceFieldsSelectOptions {
	s.Value = &v
	return s
}

type DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsRollUpSummaryFields struct {
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	Aggregator *string `json:"aggregator,omitempty" xml:"aggregator,omitempty"`
}

func (s DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsRollUpSummaryFields) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsRollUpSummaryFields) GoString() string {
	return s.String()
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsRollUpSummaryFields) SetName(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsRollUpSummaryFields {
	s.Name = &v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsRollUpSummaryFields) SetAggregator(v string) *DescribeCrmPersonalCustomerObjectMetaResponseBodyFieldsRollUpSummaryFields {
	s.Aggregator = &v
	return s
}

type DescribeCrmPersonalCustomerObjectMetaResponse struct {
	Headers map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCrmPersonalCustomerObjectMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCrmPersonalCustomerObjectMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCrmPersonalCustomerObjectMetaResponse) GoString() string {
	return s.String()
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponse) SetHeaders(v map[string]*string) *DescribeCrmPersonalCustomerObjectMetaResponse {
	s.Headers = v
	return s
}

func (s *DescribeCrmPersonalCustomerObjectMetaResponse) SetBody(v *DescribeCrmPersonalCustomerObjectMetaResponseBody) *DescribeCrmPersonalCustomerObjectMetaResponse {
	s.Body = v
	return s
}

type DeleteCrmPersonalCustomerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteCrmPersonalCustomerHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteCrmPersonalCustomerHeaders) GoString() string {
	return s.String()
}

func (s *DeleteCrmPersonalCustomerHeaders) SetCommonHeaders(v map[string]*string) *DeleteCrmPersonalCustomerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteCrmPersonalCustomerHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteCrmPersonalCustomerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteCrmPersonalCustomerRequest struct {
	// 操作人用户ID
	CurrentOperatorUserId *string `json:"currentOperatorUserId,omitempty" xml:"currentOperatorUserId,omitempty"`
}

func (s DeleteCrmPersonalCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCrmPersonalCustomerRequest) GoString() string {
	return s.String()
}

func (s *DeleteCrmPersonalCustomerRequest) SetCurrentOperatorUserId(v string) *DeleteCrmPersonalCustomerRequest {
	s.CurrentOperatorUserId = &v
	return s
}

type DeleteCrmPersonalCustomerResponseBody struct {
	// 客户数据id
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s DeleteCrmPersonalCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCrmPersonalCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCrmPersonalCustomerResponseBody) SetInstanceId(v string) *DeleteCrmPersonalCustomerResponseBody {
	s.InstanceId = &v
	return s
}

type DeleteCrmPersonalCustomerResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCrmPersonalCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCrmPersonalCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCrmPersonalCustomerResponse) GoString() string {
	return s.String()
}

func (s *DeleteCrmPersonalCustomerResponse) SetHeaders(v map[string]*string) *DeleteCrmPersonalCustomerResponse {
	s.Headers = v
	return s
}

func (s *DeleteCrmPersonalCustomerResponse) SetBody(v *DeleteCrmPersonalCustomerResponseBody) *DeleteCrmPersonalCustomerResponse {
	s.Body = v
	return s
}

type UpdateCrmPersonalCustomerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateCrmPersonalCustomerHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateCrmPersonalCustomerHeaders) GoString() string {
	return s.String()
}

func (s *UpdateCrmPersonalCustomerHeaders) SetCommonHeaders(v map[string]*string) *UpdateCrmPersonalCustomerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateCrmPersonalCustomerHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateCrmPersonalCustomerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateCrmPersonalCustomerRequest struct {
	InstanceId     *string                                     `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	ModifierUserId *string                                     `json:"modifierUserId,omitempty" xml:"modifierUserId,omitempty"`
	ModifierNick   *string                                     `json:"modifierNick,omitempty" xml:"modifierNick,omitempty"`
	Data           map[string]interface{}                      `json:"data,omitempty" xml:"data,omitempty"`
	ExtendData     map[string]interface{}                      `json:"extendData,omitempty" xml:"extendData,omitempty"`
	Permission     *UpdateCrmPersonalCustomerRequestPermission `json:"permission,omitempty" xml:"permission,omitempty" type:"Struct"`
	// 跳过uk查重
	SkipDuplicateCheck *bool `json:"skipDuplicateCheck,omitempty" xml:"skipDuplicateCheck,omitempty"`
}

func (s UpdateCrmPersonalCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCrmPersonalCustomerRequest) GoString() string {
	return s.String()
}

func (s *UpdateCrmPersonalCustomerRequest) SetInstanceId(v string) *UpdateCrmPersonalCustomerRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCrmPersonalCustomerRequest) SetModifierUserId(v string) *UpdateCrmPersonalCustomerRequest {
	s.ModifierUserId = &v
	return s
}

func (s *UpdateCrmPersonalCustomerRequest) SetModifierNick(v string) *UpdateCrmPersonalCustomerRequest {
	s.ModifierNick = &v
	return s
}

func (s *UpdateCrmPersonalCustomerRequest) SetData(v map[string]interface{}) *UpdateCrmPersonalCustomerRequest {
	s.Data = v
	return s
}

func (s *UpdateCrmPersonalCustomerRequest) SetExtendData(v map[string]interface{}) *UpdateCrmPersonalCustomerRequest {
	s.ExtendData = v
	return s
}

func (s *UpdateCrmPersonalCustomerRequest) SetPermission(v *UpdateCrmPersonalCustomerRequestPermission) *UpdateCrmPersonalCustomerRequest {
	s.Permission = v
	return s
}

func (s *UpdateCrmPersonalCustomerRequest) SetSkipDuplicateCheck(v bool) *UpdateCrmPersonalCustomerRequest {
	s.SkipDuplicateCheck = &v
	return s
}

type UpdateCrmPersonalCustomerRequestPermission struct {
	OwnerStaffIds       []*string `json:"ownerStaffIds,omitempty" xml:"ownerStaffIds,omitempty" type:"Repeated"`
	ParticipantStaffIds []*string `json:"participantStaffIds,omitempty" xml:"participantStaffIds,omitempty" type:"Repeated"`
}

func (s UpdateCrmPersonalCustomerRequestPermission) String() string {
	return tea.Prettify(s)
}

func (s UpdateCrmPersonalCustomerRequestPermission) GoString() string {
	return s.String()
}

func (s *UpdateCrmPersonalCustomerRequestPermission) SetOwnerStaffIds(v []*string) *UpdateCrmPersonalCustomerRequestPermission {
	s.OwnerStaffIds = v
	return s
}

func (s *UpdateCrmPersonalCustomerRequestPermission) SetParticipantStaffIds(v []*string) *UpdateCrmPersonalCustomerRequestPermission {
	s.ParticipantStaffIds = v
	return s
}

type UpdateCrmPersonalCustomerResponseBody struct {
	// 客户数据id
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s UpdateCrmPersonalCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCrmPersonalCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCrmPersonalCustomerResponseBody) SetInstanceId(v string) *UpdateCrmPersonalCustomerResponseBody {
	s.InstanceId = &v
	return s
}

type UpdateCrmPersonalCustomerResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCrmPersonalCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCrmPersonalCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCrmPersonalCustomerResponse) GoString() string {
	return s.String()
}

func (s *UpdateCrmPersonalCustomerResponse) SetHeaders(v map[string]*string) *UpdateCrmPersonalCustomerResponse {
	s.Headers = v
	return s
}

func (s *UpdateCrmPersonalCustomerResponse) SetBody(v *UpdateCrmPersonalCustomerResponseBody) *UpdateCrmPersonalCustomerResponse {
	s.Body = v
	return s
}

type QueryCrmPersonalCustomerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCrmPersonalCustomerHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCrmPersonalCustomerHeaders) GoString() string {
	return s.String()
}

func (s *QueryCrmPersonalCustomerHeaders) SetCommonHeaders(v map[string]*string) *QueryCrmPersonalCustomerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCrmPersonalCustomerHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCrmPersonalCustomerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCrmPersonalCustomerRequest struct {
	// 用户ID
	CurrentOperatorUserId *string `json:"currentOperatorUserId,omitempty" xml:"currentOperatorUserId,omitempty"`
	// 分页页码
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 分页条数
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 查询条件
	QueryDsl *string `json:"queryDsl,omitempty" xml:"queryDsl,omitempty"`
}

func (s QueryCrmPersonalCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCrmPersonalCustomerRequest) GoString() string {
	return s.String()
}

func (s *QueryCrmPersonalCustomerRequest) SetCurrentOperatorUserId(v string) *QueryCrmPersonalCustomerRequest {
	s.CurrentOperatorUserId = &v
	return s
}

func (s *QueryCrmPersonalCustomerRequest) SetNextToken(v string) *QueryCrmPersonalCustomerRequest {
	s.NextToken = &v
	return s
}

func (s *QueryCrmPersonalCustomerRequest) SetMaxResults(v int32) *QueryCrmPersonalCustomerRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryCrmPersonalCustomerRequest) SetQueryDsl(v string) *QueryCrmPersonalCustomerRequest {
	s.QueryDsl = &v
	return s
}

type QueryCrmPersonalCustomerResponseBody struct {
	Values     []*QueryCrmPersonalCustomerResponseBodyValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
	HasMore    *bool                                         `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken  *string                                       `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	MaxResults *int32                                        `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
}

func (s QueryCrmPersonalCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCrmPersonalCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCrmPersonalCustomerResponseBody) SetValues(v []*QueryCrmPersonalCustomerResponseBodyValues) *QueryCrmPersonalCustomerResponseBody {
	s.Values = v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBody) SetHasMore(v bool) *QueryCrmPersonalCustomerResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBody) SetNextToken(v string) *QueryCrmPersonalCustomerResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBody) SetMaxResults(v int32) *QueryCrmPersonalCustomerResponseBody {
	s.MaxResults = &v
	return s
}

type QueryCrmPersonalCustomerResponseBodyValues struct {
	// 数据ID
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// 数据类型
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// 创建记录的用户ID
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// 创建记录的用户昵称
	CreatorNick *string `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	// 数据内容
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// 扩展数据内容
	ExtendData map[string]interface{} `json:"extendData,omitempty" xml:"extendData,omitempty"`
	// 数据权限信息
	Permission *QueryCrmPersonalCustomerResponseBodyValuesPermission `json:"permission,omitempty" xml:"permission,omitempty" type:"Struct"`
	// 审批结果
	ProcOutResult *string `json:"procOutResult,omitempty" xml:"procOutResult,omitempty"`
	// 审批状态
	ProcInstStatus *string `json:"procInstStatus,omitempty" xml:"procInstStatus,omitempty"`
	// 记录创建时间
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 记录修改时间
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
}

func (s QueryCrmPersonalCustomerResponseBodyValues) String() string {
	return tea.Prettify(s)
}

func (s QueryCrmPersonalCustomerResponseBodyValues) GoString() string {
	return s.String()
}

func (s *QueryCrmPersonalCustomerResponseBodyValues) SetInstanceId(v string) *QueryCrmPersonalCustomerResponseBodyValues {
	s.InstanceId = &v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBodyValues) SetObjectType(v string) *QueryCrmPersonalCustomerResponseBodyValues {
	s.ObjectType = &v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBodyValues) SetCreatorUserId(v string) *QueryCrmPersonalCustomerResponseBodyValues {
	s.CreatorUserId = &v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBodyValues) SetCreatorNick(v string) *QueryCrmPersonalCustomerResponseBodyValues {
	s.CreatorNick = &v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBodyValues) SetData(v map[string]interface{}) *QueryCrmPersonalCustomerResponseBodyValues {
	s.Data = v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBodyValues) SetExtendData(v map[string]interface{}) *QueryCrmPersonalCustomerResponseBodyValues {
	s.ExtendData = v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBodyValues) SetPermission(v *QueryCrmPersonalCustomerResponseBodyValuesPermission) *QueryCrmPersonalCustomerResponseBodyValues {
	s.Permission = v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBodyValues) SetProcOutResult(v string) *QueryCrmPersonalCustomerResponseBodyValues {
	s.ProcOutResult = &v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBodyValues) SetProcInstStatus(v string) *QueryCrmPersonalCustomerResponseBodyValues {
	s.ProcInstStatus = &v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBodyValues) SetGmtCreate(v string) *QueryCrmPersonalCustomerResponseBodyValues {
	s.GmtCreate = &v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBodyValues) SetGmtModified(v string) *QueryCrmPersonalCustomerResponseBodyValues {
	s.GmtModified = &v
	return s
}

type QueryCrmPersonalCustomerResponseBodyValuesPermission struct {
	// 负责人用户ID列表
	OwnerStaffIds []*string `json:"ownerStaffIds,omitempty" xml:"ownerStaffIds,omitempty" type:"Repeated"`
	// 协同人用户ID列表
	ParticipantStaffIds []*string `json:"participantStaffIds,omitempty" xml:"participantStaffIds,omitempty" type:"Repeated"`
}

func (s QueryCrmPersonalCustomerResponseBodyValuesPermission) String() string {
	return tea.Prettify(s)
}

func (s QueryCrmPersonalCustomerResponseBodyValuesPermission) GoString() string {
	return s.String()
}

func (s *QueryCrmPersonalCustomerResponseBodyValuesPermission) SetOwnerStaffIds(v []*string) *QueryCrmPersonalCustomerResponseBodyValuesPermission {
	s.OwnerStaffIds = v
	return s
}

func (s *QueryCrmPersonalCustomerResponseBodyValuesPermission) SetParticipantStaffIds(v []*string) *QueryCrmPersonalCustomerResponseBodyValuesPermission {
	s.ParticipantStaffIds = v
	return s
}

type QueryCrmPersonalCustomerResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCrmPersonalCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCrmPersonalCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCrmPersonalCustomerResponse) GoString() string {
	return s.String()
}

func (s *QueryCrmPersonalCustomerResponse) SetHeaders(v map[string]*string) *QueryCrmPersonalCustomerResponse {
	s.Headers = v
	return s
}

func (s *QueryCrmPersonalCustomerResponse) SetBody(v *QueryCrmPersonalCustomerResponseBody) *QueryCrmPersonalCustomerResponse {
	s.Body = v
	return s
}

type ListCrmPersonalCustomersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListCrmPersonalCustomersHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListCrmPersonalCustomersHeaders) GoString() string {
	return s.String()
}

func (s *ListCrmPersonalCustomersHeaders) SetCommonHeaders(v map[string]*string) *ListCrmPersonalCustomersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListCrmPersonalCustomersHeaders) SetXAcsDingtalkAccessToken(v string) *ListCrmPersonalCustomersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListCrmPersonalCustomersRequest struct {
	// 操作人用户ID
	CurrentOperatorUserId *string `json:"currentOperatorUserId,omitempty" xml:"currentOperatorUserId,omitempty"`
	// 数据客户列表
	Body []*string `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s ListCrmPersonalCustomersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCrmPersonalCustomersRequest) GoString() string {
	return s.String()
}

func (s *ListCrmPersonalCustomersRequest) SetCurrentOperatorUserId(v string) *ListCrmPersonalCustomersRequest {
	s.CurrentOperatorUserId = &v
	return s
}

func (s *ListCrmPersonalCustomersRequest) SetBody(v []*string) *ListCrmPersonalCustomersRequest {
	s.Body = v
	return s
}

type ListCrmPersonalCustomersResponseBody struct {
	Result []*ListCrmPersonalCustomersResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListCrmPersonalCustomersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCrmPersonalCustomersResponseBody) GoString() string {
	return s.String()
}

func (s *ListCrmPersonalCustomersResponseBody) SetResult(v []*ListCrmPersonalCustomersResponseBodyResult) *ListCrmPersonalCustomersResponseBody {
	s.Result = v
	return s
}

type ListCrmPersonalCustomersResponseBodyResult struct {
	OrgId          *int64                                                `json:"orgId,omitempty" xml:"orgId,omitempty"`
	InstanceId     *string                                               `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	ObjectType     *string                                               `json:"objectType,omitempty" xml:"objectType,omitempty"`
	CreatorUserId  *string                                               `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	CreatorNick    *string                                               `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	Data           map[string]interface{}                                `json:"data,omitempty" xml:"data,omitempty"`
	ExtendData     map[string]interface{}                                `json:"extendData,omitempty" xml:"extendData,omitempty"`
	Permission     *ListCrmPersonalCustomersResponseBodyResultPermission `json:"permission,omitempty" xml:"permission,omitempty" type:"Struct"`
	AppUuid        *string                                               `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	FormCode       *string                                               `json:"formCode,omitempty" xml:"formCode,omitempty"`
	ProcOutResult  *string                                               `json:"procOutResult,omitempty" xml:"procOutResult,omitempty"`
	ProcInstStatus *string                                               `json:"procInstStatus,omitempty" xml:"procInstStatus,omitempty"`
	GmtCreate      *string                                               `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string                                               `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
}

func (s ListCrmPersonalCustomersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListCrmPersonalCustomersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetOrgId(v int64) *ListCrmPersonalCustomersResponseBodyResult {
	s.OrgId = &v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetInstanceId(v string) *ListCrmPersonalCustomersResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetObjectType(v string) *ListCrmPersonalCustomersResponseBodyResult {
	s.ObjectType = &v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetCreatorUserId(v string) *ListCrmPersonalCustomersResponseBodyResult {
	s.CreatorUserId = &v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetCreatorNick(v string) *ListCrmPersonalCustomersResponseBodyResult {
	s.CreatorNick = &v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetData(v map[string]interface{}) *ListCrmPersonalCustomersResponseBodyResult {
	s.Data = v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetExtendData(v map[string]interface{}) *ListCrmPersonalCustomersResponseBodyResult {
	s.ExtendData = v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetPermission(v *ListCrmPersonalCustomersResponseBodyResultPermission) *ListCrmPersonalCustomersResponseBodyResult {
	s.Permission = v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetAppUuid(v string) *ListCrmPersonalCustomersResponseBodyResult {
	s.AppUuid = &v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetFormCode(v string) *ListCrmPersonalCustomersResponseBodyResult {
	s.FormCode = &v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetProcOutResult(v string) *ListCrmPersonalCustomersResponseBodyResult {
	s.ProcOutResult = &v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetProcInstStatus(v string) *ListCrmPersonalCustomersResponseBodyResult {
	s.ProcInstStatus = &v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetGmtCreate(v string) *ListCrmPersonalCustomersResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResult) SetGmtModified(v string) *ListCrmPersonalCustomersResponseBodyResult {
	s.GmtModified = &v
	return s
}

type ListCrmPersonalCustomersResponseBodyResultPermission struct {
	OwnerStaffIds       []*string `json:"ownerStaffIds,omitempty" xml:"ownerStaffIds,omitempty" type:"Repeated"`
	ParticipantStaffIds []*string `json:"participantStaffIds,omitempty" xml:"participantStaffIds,omitempty" type:"Repeated"`
}

func (s ListCrmPersonalCustomersResponseBodyResultPermission) String() string {
	return tea.Prettify(s)
}

func (s ListCrmPersonalCustomersResponseBodyResultPermission) GoString() string {
	return s.String()
}

func (s *ListCrmPersonalCustomersResponseBodyResultPermission) SetOwnerStaffIds(v []*string) *ListCrmPersonalCustomersResponseBodyResultPermission {
	s.OwnerStaffIds = v
	return s
}

func (s *ListCrmPersonalCustomersResponseBodyResultPermission) SetParticipantStaffIds(v []*string) *ListCrmPersonalCustomersResponseBodyResultPermission {
	s.ParticipantStaffIds = v
	return s
}

type ListCrmPersonalCustomersResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCrmPersonalCustomersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCrmPersonalCustomersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCrmPersonalCustomersResponse) GoString() string {
	return s.String()
}

func (s *ListCrmPersonalCustomersResponse) SetHeaders(v map[string]*string) *ListCrmPersonalCustomersResponse {
	s.Headers = v
	return s
}

func (s *ListCrmPersonalCustomersResponse) SetBody(v *ListCrmPersonalCustomersResponseBody) *ListCrmPersonalCustomersResponse {
	s.Body = v
	return s
}

type CreateCustomerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateCustomerHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerHeaders) GoString() string {
	return s.String()
}

func (s *CreateCustomerHeaders) SetCommonHeaders(v map[string]*string) *CreateCustomerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateCustomerHeaders) SetXAcsDingtalkAccessToken(v string) *CreateCustomerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateCustomerRequest struct {
	// 写入客户类型：个人客户crm_customer_personal; 企业客户crm_customer
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// 已存在客户时，添加联系人，可以传入客户的instanceId用作关联绑定
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// 创建人的userId
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// 客户实例数据（表单数据）
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// 客户实例扩展数据
	ExtendData map[string]interface{} `json:"extendData,omitempty" xml:"extendData,omitempty"`
	// 关联联系人数据
	Contacts []*CreateCustomerRequestContacts `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	// 权限
	Permission *CreateCustomerRequestPermission `json:"permission,omitempty" xml:"permission,omitempty" type:"Struct"`
	// 保存配置项
	SaveOption *CreateCustomerRequestSaveOption `json:"saveOption,omitempty" xml:"saveOption,omitempty" type:"Struct"`
}

func (s CreateCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomerRequest) SetObjectType(v string) *CreateCustomerRequest {
	s.ObjectType = &v
	return s
}

func (s *CreateCustomerRequest) SetInstanceId(v string) *CreateCustomerRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCustomerRequest) SetCreatorUserId(v string) *CreateCustomerRequest {
	s.CreatorUserId = &v
	return s
}

func (s *CreateCustomerRequest) SetData(v map[string]interface{}) *CreateCustomerRequest {
	s.Data = v
	return s
}

func (s *CreateCustomerRequest) SetExtendData(v map[string]interface{}) *CreateCustomerRequest {
	s.ExtendData = v
	return s
}

func (s *CreateCustomerRequest) SetContacts(v []*CreateCustomerRequestContacts) *CreateCustomerRequest {
	s.Contacts = v
	return s
}

func (s *CreateCustomerRequest) SetPermission(v *CreateCustomerRequestPermission) *CreateCustomerRequest {
	s.Permission = v
	return s
}

func (s *CreateCustomerRequest) SetSaveOption(v *CreateCustomerRequestSaveOption) *CreateCustomerRequest {
	s.SaveOption = v
	return s
}

type CreateCustomerRequestContacts struct {
	// 联系人表单数据
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// 联系人扩展数据
	ExtendData map[string]interface{} `json:"extendData,omitempty" xml:"extendData,omitempty"`
}

func (s CreateCustomerRequestContacts) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerRequestContacts) GoString() string {
	return s.String()
}

func (s *CreateCustomerRequestContacts) SetData(v map[string]interface{}) *CreateCustomerRequestContacts {
	s.Data = v
	return s
}

func (s *CreateCustomerRequestContacts) SetExtendData(v map[string]interface{}) *CreateCustomerRequestContacts {
	s.ExtendData = v
	return s
}

type CreateCustomerRequestPermission struct {
	// 负责人
	OwnerStaffIds []*string `json:"ownerStaffIds,omitempty" xml:"ownerStaffIds,omitempty" type:"Repeated"`
	// 协同人
	ParticipantStaffIds []*string `json:"participantStaffIds,omitempty" xml:"participantStaffIds,omitempty" type:"Repeated"`
}

func (s CreateCustomerRequestPermission) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerRequestPermission) GoString() string {
	return s.String()
}

func (s *CreateCustomerRequestPermission) SetOwnerStaffIds(v []*string) *CreateCustomerRequestPermission {
	s.OwnerStaffIds = v
	return s
}

func (s *CreateCustomerRequestPermission) SetParticipantStaffIds(v []*string) *CreateCustomerRequestPermission {
	s.ParticipantStaffIds = v
	return s
}

type CreateCustomerRequestSaveOption struct {
	// 关注配置：0 不处理， 1 自动关注（需要单独申请白名单）
	SubscribePolicy *int64 `json:"subscribePolicy,omitempty" xml:"subscribePolicy,omitempty"`
	// 保存联系人失败时是否阻断
	ThrowExceptionWhileSavingContactFailed *bool `json:"throwExceptionWhileSavingContactFailed,omitempty" xml:"throwExceptionWhileSavingContactFailed,omitempty"`
	// 客户已存在时的处理策略：APPEND_CONTACT_FORCE 直接追加联系人； REJECT 返回失败
	CustomerExistedPolicy *string `json:"customerExistedPolicy,omitempty" xml:"customerExistedPolicy,omitempty"`
	// 跳过uk查重
	SkipDuplicateCheck *bool `json:"skipDuplicateCheck,omitempty" xml:"skipDuplicateCheck,omitempty"`
}

func (s CreateCustomerRequestSaveOption) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerRequestSaveOption) GoString() string {
	return s.String()
}

func (s *CreateCustomerRequestSaveOption) SetSubscribePolicy(v int64) *CreateCustomerRequestSaveOption {
	s.SubscribePolicy = &v
	return s
}

func (s *CreateCustomerRequestSaveOption) SetThrowExceptionWhileSavingContactFailed(v bool) *CreateCustomerRequestSaveOption {
	s.ThrowExceptionWhileSavingContactFailed = &v
	return s
}

func (s *CreateCustomerRequestSaveOption) SetCustomerExistedPolicy(v string) *CreateCustomerRequestSaveOption {
	s.CustomerExistedPolicy = &v
	return s
}

func (s *CreateCustomerRequestSaveOption) SetSkipDuplicateCheck(v bool) *CreateCustomerRequestSaveOption {
	s.SkipDuplicateCheck = &v
	return s
}

type CreateCustomerResponseBody struct {
	// 客户实例id
	CustomerInstanceId *string `json:"customerInstanceId,omitempty" xml:"customerInstanceId,omitempty"`
	// 保存客户类型
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// 联系人保存结果
	Contacts []*CreateCustomerResponseBodyContacts `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
}

func (s CreateCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomerResponseBody) SetCustomerInstanceId(v string) *CreateCustomerResponseBody {
	s.CustomerInstanceId = &v
	return s
}

func (s *CreateCustomerResponseBody) SetObjectType(v string) *CreateCustomerResponseBody {
	s.ObjectType = &v
	return s
}

func (s *CreateCustomerResponseBody) SetContacts(v []*CreateCustomerResponseBodyContacts) *CreateCustomerResponseBody {
	s.Contacts = v
	return s
}

type CreateCustomerResponseBodyContacts struct {
	// 联系人实例id
	ContactInstanceId *string `json:"contactInstanceId,omitempty" xml:"contactInstanceId,omitempty"`
}

func (s CreateCustomerResponseBodyContacts) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerResponseBodyContacts) GoString() string {
	return s.String()
}

func (s *CreateCustomerResponseBodyContacts) SetContactInstanceId(v string) *CreateCustomerResponseBodyContacts {
	s.ContactInstanceId = &v
	return s
}

type CreateCustomerResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomerResponse) SetHeaders(v map[string]*string) *CreateCustomerResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomerResponse) SetBody(v *CreateCustomerResponseBody) *CreateCustomerResponse {
	s.Body = v
	return s
}

type QueryAllTracksHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllTracksHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllTracksHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllTracksHeaders) SetCommonHeaders(v map[string]*string) *QueryAllTracksHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllTracksHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllTracksHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllTracksRequest struct {
	// 分页游标
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 分页size
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 排序
	Order *string `json:"order,omitempty" xml:"order,omitempty"`
}

func (s QueryAllTracksRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllTracksRequest) GoString() string {
	return s.String()
}

func (s *QueryAllTracksRequest) SetNextToken(v string) *QueryAllTracksRequest {
	s.NextToken = &v
	return s
}

func (s *QueryAllTracksRequest) SetMaxResults(v int32) *QueryAllTracksRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryAllTracksRequest) SetOrder(v string) *QueryAllTracksRequest {
	s.Order = &v
	return s
}

type QueryAllTracksResponseBody struct {
	// 客户动态分页数据
	Values []*QueryAllTracksResponseBodyValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
	// 是否还有数据
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 下页翻页起始游标
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 翻页size
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
}

func (s QueryAllTracksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllTracksResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllTracksResponseBody) SetValues(v []*QueryAllTracksResponseBodyValues) *QueryAllTracksResponseBody {
	s.Values = v
	return s
}

func (s *QueryAllTracksResponseBody) SetHasMore(v bool) *QueryAllTracksResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryAllTracksResponseBody) SetNextToken(v string) *QueryAllTracksResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryAllTracksResponseBody) SetMaxResults(v int32) *QueryAllTracksResponseBody {
	s.MaxResults = &v
	return s
}

type QueryAllTracksResponseBodyValues struct {
	// 企业id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 客户id
	CustomerId *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	// 动态类型
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// 动态子类型
	SubType *int32 `json:"subType,omitempty" xml:"subType,omitempty"`
	// 创建时间
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 创建人userId
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 动态外键
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
}

func (s QueryAllTracksResponseBodyValues) String() string {
	return tea.Prettify(s)
}

func (s QueryAllTracksResponseBodyValues) GoString() string {
	return s.String()
}

func (s *QueryAllTracksResponseBodyValues) SetCorpId(v string) *QueryAllTracksResponseBodyValues {
	s.CorpId = &v
	return s
}

func (s *QueryAllTracksResponseBodyValues) SetCustomerId(v string) *QueryAllTracksResponseBodyValues {
	s.CustomerId = &v
	return s
}

func (s *QueryAllTracksResponseBodyValues) SetType(v int32) *QueryAllTracksResponseBodyValues {
	s.Type = &v
	return s
}

func (s *QueryAllTracksResponseBodyValues) SetSubType(v int32) *QueryAllTracksResponseBodyValues {
	s.SubType = &v
	return s
}

func (s *QueryAllTracksResponseBodyValues) SetGmtCreate(v int64) *QueryAllTracksResponseBodyValues {
	s.GmtCreate = &v
	return s
}

func (s *QueryAllTracksResponseBodyValues) SetCreator(v string) *QueryAllTracksResponseBodyValues {
	s.Creator = &v
	return s
}

func (s *QueryAllTracksResponseBodyValues) SetBizId(v string) *QueryAllTracksResponseBodyValues {
	s.BizId = &v
	return s
}

type QueryAllTracksResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAllTracksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllTracksResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllTracksResponse) GoString() string {
	return s.String()
}

func (s *QueryAllTracksResponse) SetHeaders(v map[string]*string) *QueryAllTracksResponse {
	s.Headers = v
	return s
}

func (s *QueryAllTracksResponse) SetBody(v *QueryAllTracksResponseBody) *QueryAllTracksResponse {
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

func (client *Client) GetOfficialAccountContacts(request *GetOfficialAccountContactsRequest) (_result *GetOfficialAccountContactsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOfficialAccountContactsHeaders{}
	_result = &GetOfficialAccountContactsResponse{}
	_body, _err := client.GetOfficialAccountContactsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOfficialAccountContactsWithOptions(request *GetOfficialAccountContactsRequest, headers *GetOfficialAccountContactsHeaders, runtime *util.RuntimeOptions) (_result *GetOfficialAccountContactsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
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
	_result = &GetOfficialAccountContactsResponse{}
	_body, _err := client.DoROARequest(tea.String("GetOfficialAccountContacts"), tea.String("crm_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/crm/officialAccounts/contacts"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ServiceWindowMessageBatchPush(request *ServiceWindowMessageBatchPushRequest) (_result *ServiceWindowMessageBatchPushResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ServiceWindowMessageBatchPushHeaders{}
	_result = &ServiceWindowMessageBatchPushResponse{}
	_body, _err := client.ServiceWindowMessageBatchPushWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ServiceWindowMessageBatchPushWithOptions(request *ServiceWindowMessageBatchPushRequest, headers *ServiceWindowMessageBatchPushHeaders, runtime *util.RuntimeOptions) (_result *ServiceWindowMessageBatchPushResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Detail))) {
		body["detail"] = request.Detail
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &ServiceWindowMessageBatchPushResponse{}
	_body, _err := client.DoROARequest(tea.String("ServiceWindowMessageBatchPush"), tea.String("crm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/crm/messages/batchSend"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCrmFormInstance(instanceId *string, request *DeleteCrmFormInstanceRequest) (_result *DeleteCrmFormInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteCrmFormInstanceHeaders{}
	_result = &DeleteCrmFormInstanceResponse{}
	_body, _err := client.DeleteCrmFormInstanceWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCrmFormInstanceWithOptions(instanceId *string, request *DeleteCrmFormInstanceRequest, headers *DeleteCrmFormInstanceHeaders, runtime *util.RuntimeOptions) (_result *DeleteCrmFormInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentOperatorUserId)) {
		query["currentOperatorUserId"] = request.CurrentOperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
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
	_result = &DeleteCrmFormInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteCrmFormInstance"), tea.String("crm_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/crm/formInstances/"+tea.StringValue(instanceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchSendOfficialAccountOTOMessage(request *BatchSendOfficialAccountOTOMessageRequest) (_result *BatchSendOfficialAccountOTOMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchSendOfficialAccountOTOMessageHeaders{}
	_result = &BatchSendOfficialAccountOTOMessageResponse{}
	_body, _err := client.BatchSendOfficialAccountOTOMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchSendOfficialAccountOTOMessageWithOptions(request *BatchSendOfficialAccountOTOMessageRequest, headers *BatchSendOfficialAccountOTOMessageHeaders, runtime *util.RuntimeOptions) (_result *BatchSendOfficialAccountOTOMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Detail))) {
		body["detail"] = request.Detail
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["accountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &BatchSendOfficialAccountOTOMessageResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchSendOfficialAccountOTOMessage"), tea.String("crm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/crm/officialAccounts/oToMessages/batchSend"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOfficialAccountContactInfo(userId *string) (_result *GetOfficialAccountContactInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOfficialAccountContactInfoHeaders{}
	_result = &GetOfficialAccountContactInfoResponse{}
	_body, _err := client.GetOfficialAccountContactInfoWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOfficialAccountContactInfoWithOptions(userId *string, headers *GetOfficialAccountContactInfoHeaders, runtime *util.RuntimeOptions) (_result *GetOfficialAccountContactInfoResponse, _err error) {
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
	_result = &GetOfficialAccountContactInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetOfficialAccountContactInfo"), tea.String("crm_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/crm/officialAccounts/contacts/"+tea.StringValue(userId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllCustomer(request *QueryAllCustomerRequest) (_result *QueryAllCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllCustomerHeaders{}
	_result = &QueryAllCustomerResponse{}
	_body, _err := client.QueryAllCustomerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllCustomerWithOptions(request *QueryAllCustomerRequest, headers *QueryAllCustomerHeaders, runtime *util.RuntimeOptions) (_result *QueryAllCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.DingCorpId)) {
		body["dingCorpId"] = request.DingCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectType)) {
		body["objectType"] = request.ObjectType
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &QueryAllCustomerResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAllCustomer"), tea.String("crm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/crm/customerInstances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendOfficialAccountOTOMessage(request *SendOfficialAccountOTOMessageRequest) (_result *SendOfficialAccountOTOMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendOfficialAccountOTOMessageHeaders{}
	_result = &SendOfficialAccountOTOMessageResponse{}
	_body, _err := client.SendOfficialAccountOTOMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendOfficialAccountOTOMessageWithOptions(request *SendOfficialAccountOTOMessageRequest, headers *SendOfficialAccountOTOMessageHeaders, runtime *util.RuntimeOptions) (_result *SendOfficialAccountOTOMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Detail))) {
		body["detail"] = request.Detail
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["accountId"] = request.AccountId
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &SendOfficialAccountOTOMessageResponse{}
	_body, _err := client.DoROARequest(tea.String("SendOfficialAccountOTOMessage"), tea.String("crm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/crm/officialAccounts/oToMessages/send"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOfficialAccountOTOMessageResult(request *GetOfficialAccountOTOMessageResultRequest) (_result *GetOfficialAccountOTOMessageResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOfficialAccountOTOMessageResultHeaders{}
	_result = &GetOfficialAccountOTOMessageResultResponse{}
	_body, _err := client.GetOfficialAccountOTOMessageResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOfficialAccountOTOMessageResultWithOptions(request *GetOfficialAccountOTOMessageResultRequest, headers *GetOfficialAccountOTOMessageResultHeaders, runtime *util.RuntimeOptions) (_result *GetOfficialAccountOTOMessageResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenPushId)) {
		query["openPushId"] = request.OpenPushId
	}

	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["accountId"] = request.AccountId
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
	_result = &GetOfficialAccountOTOMessageResultResponse{}
	_body, _err := client.DoROARequest(tea.String("GetOfficialAccountOTOMessageResult"), tea.String("crm_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/crm/officialAccounts/oToMessages/sendResults"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddCrmPersonalCustomer(request *AddCrmPersonalCustomerRequest) (_result *AddCrmPersonalCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddCrmPersonalCustomerHeaders{}
	_result = &AddCrmPersonalCustomerResponse{}
	_body, _err := client.AddCrmPersonalCustomerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddCrmPersonalCustomerWithOptions(request *AddCrmPersonalCustomerRequest, headers *AddCrmPersonalCustomerHeaders, runtime *util.RuntimeOptions) (_result *AddCrmPersonalCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreatorUserId)) {
		body["creatorUserId"] = request.CreatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorNick)) {
		body["creatorNick"] = request.CreatorNick
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendData)) {
		body["extendData"] = request.ExtendData
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Permission))) {
		body["permission"] = request.Permission
	}

	if !tea.BoolValue(util.IsUnset(request.SkipDuplicateCheck)) {
		body["skipDuplicateCheck"] = request.SkipDuplicateCheck
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &AddCrmPersonalCustomerResponse{}
	_body, _err := client.DoROARequest(tea.String("AddCrmPersonalCustomer"), tea.String("crm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/crm/personalCustomers"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecallOfficialAccountOTOMessage(request *RecallOfficialAccountOTOMessageRequest) (_result *RecallOfficialAccountOTOMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RecallOfficialAccountOTOMessageHeaders{}
	_result = &RecallOfficialAccountOTOMessageResponse{}
	_body, _err := client.RecallOfficialAccountOTOMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecallOfficialAccountOTOMessageWithOptions(request *RecallOfficialAccountOTOMessageRequest, headers *RecallOfficialAccountOTOMessageHeaders, runtime *util.RuntimeOptions) (_result *RecallOfficialAccountOTOMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["accountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenPushId)) {
		body["openPushId"] = request.OpenPushId
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &RecallOfficialAccountOTOMessageResponse{}
	_body, _err := client.DoROARequest(tea.String("RecallOfficialAccountOTOMessage"), tea.String("crm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/crm/officialAccounts/oToMessages/recall"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCrmPersonalCustomerObjectMeta() (_result *DescribeCrmPersonalCustomerObjectMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DescribeCrmPersonalCustomerObjectMetaHeaders{}
	_result = &DescribeCrmPersonalCustomerObjectMetaResponse{}
	_body, _err := client.DescribeCrmPersonalCustomerObjectMetaWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCrmPersonalCustomerObjectMetaWithOptions(headers *DescribeCrmPersonalCustomerObjectMetaHeaders, runtime *util.RuntimeOptions) (_result *DescribeCrmPersonalCustomerObjectMetaResponse, _err error) {
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
	_result = &DescribeCrmPersonalCustomerObjectMetaResponse{}
	_body, _err := client.DoROARequest(tea.String("DescribeCrmPersonalCustomerObjectMeta"), tea.String("crm_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/crm/personalCustomers/objectMeta"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCrmPersonalCustomer(dataId *string, request *DeleteCrmPersonalCustomerRequest) (_result *DeleteCrmPersonalCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteCrmPersonalCustomerHeaders{}
	_result = &DeleteCrmPersonalCustomerResponse{}
	_body, _err := client.DeleteCrmPersonalCustomerWithOptions(dataId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCrmPersonalCustomerWithOptions(dataId *string, request *DeleteCrmPersonalCustomerRequest, headers *DeleteCrmPersonalCustomerHeaders, runtime *util.RuntimeOptions) (_result *DeleteCrmPersonalCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentOperatorUserId)) {
		query["currentOperatorUserId"] = request.CurrentOperatorUserId
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
	_result = &DeleteCrmPersonalCustomerResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteCrmPersonalCustomer"), tea.String("crm_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/crm/personalCustomers/"+tea.StringValue(dataId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCrmPersonalCustomer(request *UpdateCrmPersonalCustomerRequest) (_result *UpdateCrmPersonalCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateCrmPersonalCustomerHeaders{}
	_result = &UpdateCrmPersonalCustomerResponse{}
	_body, _err := client.UpdateCrmPersonalCustomerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCrmPersonalCustomerWithOptions(request *UpdateCrmPersonalCustomerRequest, headers *UpdateCrmPersonalCustomerHeaders, runtime *util.RuntimeOptions) (_result *UpdateCrmPersonalCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ModifierUserId)) {
		body["modifierUserId"] = request.ModifierUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ModifierNick)) {
		body["modifierNick"] = request.ModifierNick
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendData)) {
		body["extendData"] = request.ExtendData
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Permission))) {
		body["permission"] = request.Permission
	}

	if !tea.BoolValue(util.IsUnset(request.SkipDuplicateCheck)) {
		body["skipDuplicateCheck"] = request.SkipDuplicateCheck
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateCrmPersonalCustomerResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateCrmPersonalCustomer"), tea.String("crm_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/crm/personalCustomers"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCrmPersonalCustomer(request *QueryCrmPersonalCustomerRequest) (_result *QueryCrmPersonalCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCrmPersonalCustomerHeaders{}
	_result = &QueryCrmPersonalCustomerResponse{}
	_body, _err := client.QueryCrmPersonalCustomerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCrmPersonalCustomerWithOptions(request *QueryCrmPersonalCustomerRequest, headers *QueryCrmPersonalCustomerHeaders, runtime *util.RuntimeOptions) (_result *QueryCrmPersonalCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentOperatorUserId)) {
		query["currentOperatorUserId"] = request.CurrentOperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.QueryDsl)) {
		query["queryDsl"] = request.QueryDsl
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
	_result = &QueryCrmPersonalCustomerResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryCrmPersonalCustomer"), tea.String("crm_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/crm/personalCustomers"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCrmPersonalCustomers(request *ListCrmPersonalCustomersRequest) (_result *ListCrmPersonalCustomersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListCrmPersonalCustomersHeaders{}
	_result = &ListCrmPersonalCustomersResponse{}
	_body, _err := client.ListCrmPersonalCustomersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCrmPersonalCustomersWithOptions(request *ListCrmPersonalCustomersRequest, headers *ListCrmPersonalCustomersHeaders, runtime *util.RuntimeOptions) (_result *ListCrmPersonalCustomersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentOperatorUserId)) {
		query["currentOperatorUserId"] = request.CurrentOperatorUserId
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
		Body:    request.Body,
	}
	_result = &ListCrmPersonalCustomersResponse{}
	_body, _err := client.DoROARequest(tea.String("ListCrmPersonalCustomers"), tea.String("crm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/crm/personalCustomers/batchQuery"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCustomer(request *CreateCustomerRequest) (_result *CreateCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateCustomerHeaders{}
	_result = &CreateCustomerResponse{}
	_body, _err := client.CreateCustomerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCustomerWithOptions(request *CreateCustomerRequest, headers *CreateCustomerHeaders, runtime *util.RuntimeOptions) (_result *CreateCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ObjectType)) {
		body["objectType"] = request.ObjectType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorUserId)) {
		body["creatorUserId"] = request.CreatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendData)) {
		body["extendData"] = request.ExtendData
	}

	if !tea.BoolValue(util.IsUnset(request.Contacts)) {
		body["contacts"] = request.Contacts
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Permission))) {
		body["permission"] = request.Permission
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SaveOption))) {
		body["saveOption"] = request.SaveOption
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
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateCustomerResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateCustomer"), tea.String("crm_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/crm/customers"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllTracks(request *QueryAllTracksRequest) (_result *QueryAllTracksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllTracksHeaders{}
	_result = &QueryAllTracksResponse{}
	_body, _err := client.QueryAllTracksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllTracksWithOptions(request *QueryAllTracksRequest, headers *QueryAllTracksHeaders, runtime *util.RuntimeOptions) (_result *QueryAllTracksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["order"] = request.Order
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
	_result = &QueryAllTracksResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAllTracks"), tea.String("crm_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/crm/customers/tracks"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
