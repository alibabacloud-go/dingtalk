// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package card_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AppendSpaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AppendSpaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceHeaders) GoString() string {
	return s.String()
}

func (s *AppendSpaceHeaders) SetCommonHeaders(v map[string]*string) *AppendSpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AppendSpaceHeaders) SetXAcsDingtalkAccessToken(v string) *AppendSpaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AppendSpaceRequest struct {
	// 协作场域信息
	CoFeedOpenSpaceModel *AppendSpaceRequestCoFeedOpenSpaceModel `json:"coFeedOpenSpaceModel,omitempty" xml:"coFeedOpenSpaceModel,omitempty" type:"Struct"`
	// IM群聊场域信息
	ImGroupOpenSpaceModel *AppendSpaceRequestImGroupOpenSpaceModel `json:"imGroupOpenSpaceModel,omitempty" xml:"imGroupOpenSpaceModel,omitempty" type:"Struct"`
	// IM单聊场域信息
	ImSingleOpenSpaceModel *AppendSpaceRequestImSingleOpenSpaceModel `json:"imSingleOpenSpaceModel,omitempty" xml:"imSingleOpenSpaceModel,omitempty" type:"Struct"`
	// 唯一标识一张卡片的外部Id
	OutTrackId *string `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	// 吊顶场域信息
	TopOpenSpaceModel *AppendSpaceRequestTopOpenSpaceModel `json:"topOpenSpaceModel,omitempty" xml:"topOpenSpaceModel,omitempty" type:"Struct"`
	// 工作台场域信息
	WorkBenchOpenSpaceModel *AppendSpaceRequestWorkBenchOpenSpaceModel `json:"workBenchOpenSpaceModel,omitempty" xml:"workBenchOpenSpaceModel,omitempty" type:"Struct"`
}

func (s AppendSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceRequest) GoString() string {
	return s.String()
}

func (s *AppendSpaceRequest) SetCoFeedOpenSpaceModel(v *AppendSpaceRequestCoFeedOpenSpaceModel) *AppendSpaceRequest {
	s.CoFeedOpenSpaceModel = v
	return s
}

func (s *AppendSpaceRequest) SetImGroupOpenSpaceModel(v *AppendSpaceRequestImGroupOpenSpaceModel) *AppendSpaceRequest {
	s.ImGroupOpenSpaceModel = v
	return s
}

func (s *AppendSpaceRequest) SetImSingleOpenSpaceModel(v *AppendSpaceRequestImSingleOpenSpaceModel) *AppendSpaceRequest {
	s.ImSingleOpenSpaceModel = v
	return s
}

func (s *AppendSpaceRequest) SetOutTrackId(v string) *AppendSpaceRequest {
	s.OutTrackId = &v
	return s
}

func (s *AppendSpaceRequest) SetTopOpenSpaceModel(v *AppendSpaceRequestTopOpenSpaceModel) *AppendSpaceRequest {
	s.TopOpenSpaceModel = v
	return s
}

func (s *AppendSpaceRequest) SetWorkBenchOpenSpaceModel(v *AppendSpaceRequestWorkBenchOpenSpaceModel) *AppendSpaceRequest {
	s.WorkBenchOpenSpaceModel = v
	return s
}

type AppendSpaceRequestCoFeedOpenSpaceModel struct {
	// 【必填】标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s AppendSpaceRequestCoFeedOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceRequestCoFeedOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *AppendSpaceRequestCoFeedOpenSpaceModel) SetTitle(v string) *AppendSpaceRequestCoFeedOpenSpaceModel {
	s.Title = &v
	return s
}

type AppendSpaceRequestImGroupOpenSpaceModel struct {
	// 支持国际化的LastMessage
	LastMessageI18n map[string]*string `json:"lastMessageI18n,omitempty" xml:"lastMessageI18n,omitempty"`
	// xpn信息
	Notification *AppendSpaceRequestImGroupOpenSpaceModelNotification `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	// 支持卡片消息可被搜索字段
	SearchSupport *AppendSpaceRequestImGroupOpenSpaceModelSearchSupport `json:"searchSupport,omitempty" xml:"searchSupport,omitempty" type:"Struct"`
	// 是否支持转发, 默认false
	SupportForward *bool `json:"supportForward,omitempty" xml:"supportForward,omitempty"`
}

func (s AppendSpaceRequestImGroupOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceRequestImGroupOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *AppendSpaceRequestImGroupOpenSpaceModel) SetLastMessageI18n(v map[string]*string) *AppendSpaceRequestImGroupOpenSpaceModel {
	s.LastMessageI18n = v
	return s
}

func (s *AppendSpaceRequestImGroupOpenSpaceModel) SetNotification(v *AppendSpaceRequestImGroupOpenSpaceModelNotification) *AppendSpaceRequestImGroupOpenSpaceModel {
	s.Notification = v
	return s
}

func (s *AppendSpaceRequestImGroupOpenSpaceModel) SetSearchSupport(v *AppendSpaceRequestImGroupOpenSpaceModelSearchSupport) *AppendSpaceRequestImGroupOpenSpaceModel {
	s.SearchSupport = v
	return s
}

func (s *AppendSpaceRequestImGroupOpenSpaceModel) SetSupportForward(v bool) *AppendSpaceRequestImGroupOpenSpaceModel {
	s.SupportForward = &v
	return s
}

type AppendSpaceRequestImGroupOpenSpaceModelNotification struct {
	AlertContent    *string `json:"alertContent,omitempty" xml:"alertContent,omitempty"`
	NotificationOff *bool   `json:"notificationOff,omitempty" xml:"notificationOff,omitempty"`
}

func (s AppendSpaceRequestImGroupOpenSpaceModelNotification) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceRequestImGroupOpenSpaceModelNotification) GoString() string {
	return s.String()
}

func (s *AppendSpaceRequestImGroupOpenSpaceModelNotification) SetAlertContent(v string) *AppendSpaceRequestImGroupOpenSpaceModelNotification {
	s.AlertContent = &v
	return s
}

func (s *AppendSpaceRequestImGroupOpenSpaceModelNotification) SetNotificationOff(v bool) *AppendSpaceRequestImGroupOpenSpaceModelNotification {
	s.NotificationOff = &v
	return s
}

type AppendSpaceRequestImGroupOpenSpaceModelSearchSupport struct {
	SearchDesc     *string `json:"searchDesc,omitempty" xml:"searchDesc,omitempty"`
	SearchIcon     *string `json:"searchIcon,omitempty" xml:"searchIcon,omitempty"`
	SearchTypeName *string `json:"searchTypeName,omitempty" xml:"searchTypeName,omitempty"`
}

func (s AppendSpaceRequestImGroupOpenSpaceModelSearchSupport) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceRequestImGroupOpenSpaceModelSearchSupport) GoString() string {
	return s.String()
}

func (s *AppendSpaceRequestImGroupOpenSpaceModelSearchSupport) SetSearchDesc(v string) *AppendSpaceRequestImGroupOpenSpaceModelSearchSupport {
	s.SearchDesc = &v
	return s
}

func (s *AppendSpaceRequestImGroupOpenSpaceModelSearchSupport) SetSearchIcon(v string) *AppendSpaceRequestImGroupOpenSpaceModelSearchSupport {
	s.SearchIcon = &v
	return s
}

func (s *AppendSpaceRequestImGroupOpenSpaceModelSearchSupport) SetSearchTypeName(v string) *AppendSpaceRequestImGroupOpenSpaceModelSearchSupport {
	s.SearchTypeName = &v
	return s
}

type AppendSpaceRequestImSingleOpenSpaceModel struct {
	// 支持国际化的LastMessage
	LastMessageI18n map[string]*string `json:"lastMessageI18n,omitempty" xml:"lastMessageI18n,omitempty"`
	// xpn信息
	Notification *AppendSpaceRequestImSingleOpenSpaceModelNotification `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	// 支持卡片消息可被搜索字段
	SearchSupport *AppendSpaceRequestImSingleOpenSpaceModelSearchSupport `json:"searchSupport,omitempty" xml:"searchSupport,omitempty" type:"Struct"`
	// 是否支持转发, 默认false
	SupportForward *bool `json:"supportForward,omitempty" xml:"supportForward,omitempty"`
}

func (s AppendSpaceRequestImSingleOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceRequestImSingleOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *AppendSpaceRequestImSingleOpenSpaceModel) SetLastMessageI18n(v map[string]*string) *AppendSpaceRequestImSingleOpenSpaceModel {
	s.LastMessageI18n = v
	return s
}

func (s *AppendSpaceRequestImSingleOpenSpaceModel) SetNotification(v *AppendSpaceRequestImSingleOpenSpaceModelNotification) *AppendSpaceRequestImSingleOpenSpaceModel {
	s.Notification = v
	return s
}

func (s *AppendSpaceRequestImSingleOpenSpaceModel) SetSearchSupport(v *AppendSpaceRequestImSingleOpenSpaceModelSearchSupport) *AppendSpaceRequestImSingleOpenSpaceModel {
	s.SearchSupport = v
	return s
}

func (s *AppendSpaceRequestImSingleOpenSpaceModel) SetSupportForward(v bool) *AppendSpaceRequestImSingleOpenSpaceModel {
	s.SupportForward = &v
	return s
}

type AppendSpaceRequestImSingleOpenSpaceModelNotification struct {
	AlertContent    *string `json:"alertContent,omitempty" xml:"alertContent,omitempty"`
	NotificationOff *bool   `json:"notificationOff,omitempty" xml:"notificationOff,omitempty"`
}

func (s AppendSpaceRequestImSingleOpenSpaceModelNotification) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceRequestImSingleOpenSpaceModelNotification) GoString() string {
	return s.String()
}

func (s *AppendSpaceRequestImSingleOpenSpaceModelNotification) SetAlertContent(v string) *AppendSpaceRequestImSingleOpenSpaceModelNotification {
	s.AlertContent = &v
	return s
}

func (s *AppendSpaceRequestImSingleOpenSpaceModelNotification) SetNotificationOff(v bool) *AppendSpaceRequestImSingleOpenSpaceModelNotification {
	s.NotificationOff = &v
	return s
}

type AppendSpaceRequestImSingleOpenSpaceModelSearchSupport struct {
	SearchDesc     *string `json:"searchDesc,omitempty" xml:"searchDesc,omitempty"`
	SearchIcon     *string `json:"searchIcon,omitempty" xml:"searchIcon,omitempty"`
	SearchTypeName *string `json:"searchTypeName,omitempty" xml:"searchTypeName,omitempty"`
}

func (s AppendSpaceRequestImSingleOpenSpaceModelSearchSupport) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceRequestImSingleOpenSpaceModelSearchSupport) GoString() string {
	return s.String()
}

func (s *AppendSpaceRequestImSingleOpenSpaceModelSearchSupport) SetSearchDesc(v string) *AppendSpaceRequestImSingleOpenSpaceModelSearchSupport {
	s.SearchDesc = &v
	return s
}

func (s *AppendSpaceRequestImSingleOpenSpaceModelSearchSupport) SetSearchIcon(v string) *AppendSpaceRequestImSingleOpenSpaceModelSearchSupport {
	s.SearchIcon = &v
	return s
}

func (s *AppendSpaceRequestImSingleOpenSpaceModelSearchSupport) SetSearchTypeName(v string) *AppendSpaceRequestImSingleOpenSpaceModelSearchSupport {
	s.SearchTypeName = &v
	return s
}

type AppendSpaceRequestTopOpenSpaceModel struct {
	// 【必填】场域类型 (IM: IM, IM_SINGLE: IM单聊, IM_GROUP: IM群聊, ONE_BOX: 群吊顶, COOPERATION_FEED: 协作, WORK_BENCH: 工作台)
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
}

func (s AppendSpaceRequestTopOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceRequestTopOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *AppendSpaceRequestTopOpenSpaceModel) SetSpaceType(v string) *AppendSpaceRequestTopOpenSpaceModel {
	s.SpaceType = &v
	return s
}

type AppendSpaceRequestWorkBenchOpenSpaceModel struct {
	// 【必填】场域类型 (IM: IM, IM_SINGLE: IM单聊, IM_GROUP: IM群聊, ONE_BOX: 群吊顶, COOPERATION_FEED: 协作, WORK_BENCH: 工作台)
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
}

func (s AppendSpaceRequestWorkBenchOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceRequestWorkBenchOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *AppendSpaceRequestWorkBenchOpenSpaceModel) SetSpaceType(v string) *AppendSpaceRequestWorkBenchOpenSpaceModel {
	s.SpaceType = &v
	return s
}

type AppendSpaceResponseBody struct {
	Result  *bool `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AppendSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *AppendSpaceResponseBody) SetResult(v bool) *AppendSpaceResponseBody {
	s.Result = &v
	return s
}

func (s *AppendSpaceResponseBody) SetSuccess(v bool) *AppendSpaceResponseBody {
	s.Success = &v
	return s
}

type AppendSpaceResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AppendSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AppendSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceResponse) GoString() string {
	return s.String()
}

func (s *AppendSpaceResponse) SetHeaders(v map[string]*string) *AppendSpaceResponse {
	s.Headers = v
	return s
}

func (s *AppendSpaceResponse) SetBody(v *AppendSpaceResponseBody) *AppendSpaceResponse {
	s.Body = v
	return s
}

type CreateAndDeliverHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateAndDeliverHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverHeaders) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverHeaders) SetCommonHeaders(v map[string]*string) *CreateAndDeliverHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateAndDeliverHeaders) SetXAcsDingtalkAccessToken(v string) *CreateAndDeliverHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateAndDeliverRequest struct {
	// 卡片回调时的路由 key
	CallbackRouteKey *string `json:"callbackRouteKey,omitempty" xml:"callbackRouteKey,omitempty"`
	// 卡片数据
	CardData *CreateAndDeliverRequestCardData `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	// 卡片内容模板ID
	CardTemplateId *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	// 协作投放参数
	CoFeedOpenDeliverModel *CreateAndDeliverRequestCoFeedOpenDeliverModel `json:"coFeedOpenDeliverModel,omitempty" xml:"coFeedOpenDeliverModel,omitempty" type:"Struct"`
	// 协作场域信息
	CoFeedOpenSpaceModel *CreateAndDeliverRequestCoFeedOpenSpaceModel `json:"coFeedOpenSpaceModel,omitempty" xml:"coFeedOpenSpaceModel,omitempty" type:"Struct"`
	// 群聊投放参数
	ImGroupOpenDeliverModel *CreateAndDeliverRequestImGroupOpenDeliverModel `json:"imGroupOpenDeliverModel,omitempty" xml:"imGroupOpenDeliverModel,omitempty" type:"Struct"`
	// IM群聊场域信息
	ImGroupOpenSpaceModel *CreateAndDeliverRequestImGroupOpenSpaceModel `json:"imGroupOpenSpaceModel,omitempty" xml:"imGroupOpenSpaceModel,omitempty" type:"Struct"`
	// 单聊场域投放参数
	ImSingleOpenDeliverModel *CreateAndDeliverRequestImSingleOpenDeliverModel `json:"imSingleOpenDeliverModel,omitempty" xml:"imSingleOpenDeliverModel,omitempty" type:"Struct"`
	// IM单聊场域信息
	ImSingleOpenSpaceModel *CreateAndDeliverRequestImSingleOpenSpaceModel `json:"imSingleOpenSpaceModel,omitempty" xml:"imSingleOpenSpaceModel,omitempty" type:"Struct"`
	// 动态数据源配置
	OpenDynamicDataConfig *CreateAndDeliverRequestOpenDynamicDataConfig `json:"openDynamicDataConfig,omitempty" xml:"openDynamicDataConfig,omitempty" type:"Struct"`
	// dt.card//spaceType.spaceId;spaceType.spaceId
	OpenSpaceId *string `json:"openSpaceId,omitempty" xml:"openSpaceId,omitempty"`
	// 外部业务标识符
	OutTrackId  *string                      `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	PrivateData map[string]*PrivateDataValue `json:"privateData,omitempty" xml:"privateData,omitempty"`
	// 吊顶投放参数
	TopOpenDeliverModel *CreateAndDeliverRequestTopOpenDeliverModel `json:"topOpenDeliverModel,omitempty" xml:"topOpenDeliverModel,omitempty" type:"Struct"`
	// 吊顶场域信息
	TopOpenSpaceModel *CreateAndDeliverRequestTopOpenSpaceModel `json:"topOpenSpaceModel,omitempty" xml:"topOpenSpaceModel,omitempty" type:"Struct"`
	// 卡片创建者 id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 工作台投放参数
	WorkBenchOpenDeliverModel *CreateAndDeliverRequestWorkBenchOpenDeliverModel `json:"workBenchOpenDeliverModel,omitempty" xml:"workBenchOpenDeliverModel,omitempty" type:"Struct"`
	// 工作台场域信息
	WorkBenchOpenSpaceModel *CreateAndDeliverRequestWorkBenchOpenSpaceModel `json:"workBenchOpenSpaceModel,omitempty" xml:"workBenchOpenSpaceModel,omitempty" type:"Struct"`
}

func (s CreateAndDeliverRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequest) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequest) SetCallbackRouteKey(v string) *CreateAndDeliverRequest {
	s.CallbackRouteKey = &v
	return s
}

func (s *CreateAndDeliverRequest) SetCardData(v *CreateAndDeliverRequestCardData) *CreateAndDeliverRequest {
	s.CardData = v
	return s
}

func (s *CreateAndDeliverRequest) SetCardTemplateId(v string) *CreateAndDeliverRequest {
	s.CardTemplateId = &v
	return s
}

func (s *CreateAndDeliverRequest) SetCoFeedOpenDeliverModel(v *CreateAndDeliverRequestCoFeedOpenDeliverModel) *CreateAndDeliverRequest {
	s.CoFeedOpenDeliverModel = v
	return s
}

func (s *CreateAndDeliverRequest) SetCoFeedOpenSpaceModel(v *CreateAndDeliverRequestCoFeedOpenSpaceModel) *CreateAndDeliverRequest {
	s.CoFeedOpenSpaceModel = v
	return s
}

func (s *CreateAndDeliverRequest) SetImGroupOpenDeliverModel(v *CreateAndDeliverRequestImGroupOpenDeliverModel) *CreateAndDeliverRequest {
	s.ImGroupOpenDeliverModel = v
	return s
}

func (s *CreateAndDeliverRequest) SetImGroupOpenSpaceModel(v *CreateAndDeliverRequestImGroupOpenSpaceModel) *CreateAndDeliverRequest {
	s.ImGroupOpenSpaceModel = v
	return s
}

func (s *CreateAndDeliverRequest) SetImSingleOpenDeliverModel(v *CreateAndDeliverRequestImSingleOpenDeliverModel) *CreateAndDeliverRequest {
	s.ImSingleOpenDeliverModel = v
	return s
}

func (s *CreateAndDeliverRequest) SetImSingleOpenSpaceModel(v *CreateAndDeliverRequestImSingleOpenSpaceModel) *CreateAndDeliverRequest {
	s.ImSingleOpenSpaceModel = v
	return s
}

func (s *CreateAndDeliverRequest) SetOpenDynamicDataConfig(v *CreateAndDeliverRequestOpenDynamicDataConfig) *CreateAndDeliverRequest {
	s.OpenDynamicDataConfig = v
	return s
}

func (s *CreateAndDeliverRequest) SetOpenSpaceId(v string) *CreateAndDeliverRequest {
	s.OpenSpaceId = &v
	return s
}

func (s *CreateAndDeliverRequest) SetOutTrackId(v string) *CreateAndDeliverRequest {
	s.OutTrackId = &v
	return s
}

func (s *CreateAndDeliverRequest) SetPrivateData(v map[string]*PrivateDataValue) *CreateAndDeliverRequest {
	s.PrivateData = v
	return s
}

func (s *CreateAndDeliverRequest) SetTopOpenDeliverModel(v *CreateAndDeliverRequestTopOpenDeliverModel) *CreateAndDeliverRequest {
	s.TopOpenDeliverModel = v
	return s
}

func (s *CreateAndDeliverRequest) SetTopOpenSpaceModel(v *CreateAndDeliverRequestTopOpenSpaceModel) *CreateAndDeliverRequest {
	s.TopOpenSpaceModel = v
	return s
}

func (s *CreateAndDeliverRequest) SetUserId(v string) *CreateAndDeliverRequest {
	s.UserId = &v
	return s
}

func (s *CreateAndDeliverRequest) SetWorkBenchOpenDeliverModel(v *CreateAndDeliverRequestWorkBenchOpenDeliverModel) *CreateAndDeliverRequest {
	s.WorkBenchOpenDeliverModel = v
	return s
}

func (s *CreateAndDeliverRequest) SetWorkBenchOpenSpaceModel(v *CreateAndDeliverRequestWorkBenchOpenSpaceModel) *CreateAndDeliverRequest {
	s.WorkBenchOpenSpaceModel = v
	return s
}

type CreateAndDeliverRequestCardData struct {
	// 卡片模板-文本内容参数
	CardParamMap map[string]*string `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
}

func (s CreateAndDeliverRequestCardData) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequestCardData) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequestCardData) SetCardParamMap(v map[string]*string) *CreateAndDeliverRequestCardData {
	s.CardParamMap = v
	return s
}

type CreateAndDeliverRequestCoFeedOpenDeliverModel struct {
	// 【必填】业务标识
	BizTag *string `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	// 【必填】协作场域下的排序时间
	GmtTimeLine *int64 `json:"gmtTimeLine,omitempty" xml:"gmtTimeLine,omitempty"`
}

func (s CreateAndDeliverRequestCoFeedOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequestCoFeedOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequestCoFeedOpenDeliverModel) SetBizTag(v string) *CreateAndDeliverRequestCoFeedOpenDeliverModel {
	s.BizTag = &v
	return s
}

func (s *CreateAndDeliverRequestCoFeedOpenDeliverModel) SetGmtTimeLine(v int64) *CreateAndDeliverRequestCoFeedOpenDeliverModel {
	s.GmtTimeLine = &v
	return s
}

type CreateAndDeliverRequestCoFeedOpenSpaceModel struct {
	// 【必填】标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateAndDeliverRequestCoFeedOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequestCoFeedOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequestCoFeedOpenSpaceModel) SetTitle(v string) *CreateAndDeliverRequestCoFeedOpenSpaceModel {
	s.Title = &v
	return s
}

type CreateAndDeliverRequestImGroupOpenDeliverModel struct {
	// 消息@人，
	AtUserIds map[string]*string `json:"atUserIds,omitempty" xml:"atUserIds,omitempty"`
	// 指定接收者
	Recipients []*string `json:"recipients,omitempty" xml:"recipients,omitempty" type:"Repeated"`
	// 机器人的code
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s CreateAndDeliverRequestImGroupOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequestImGroupOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequestImGroupOpenDeliverModel) SetAtUserIds(v map[string]*string) *CreateAndDeliverRequestImGroupOpenDeliverModel {
	s.AtUserIds = v
	return s
}

func (s *CreateAndDeliverRequestImGroupOpenDeliverModel) SetRecipients(v []*string) *CreateAndDeliverRequestImGroupOpenDeliverModel {
	s.Recipients = v
	return s
}

func (s *CreateAndDeliverRequestImGroupOpenDeliverModel) SetRobotCode(v string) *CreateAndDeliverRequestImGroupOpenDeliverModel {
	s.RobotCode = &v
	return s
}

type CreateAndDeliverRequestImGroupOpenSpaceModel struct {
	// 支持国际化的LastMessage
	LastMessageI18n map[string]*string `json:"lastMessageI18n,omitempty" xml:"lastMessageI18n,omitempty"`
	// 通知信息
	Notification *CreateAndDeliverRequestImGroupOpenSpaceModelNotification `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	// 支持卡片消息可被搜索字段
	SearchSupport *CreateAndDeliverRequestImGroupOpenSpaceModelSearchSupport `json:"searchSupport,omitempty" xml:"searchSupport,omitempty" type:"Struct"`
	// 是否支持转发, 默认false
	SupportForward *bool `json:"supportForward,omitempty" xml:"supportForward,omitempty"`
}

func (s CreateAndDeliverRequestImGroupOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequestImGroupOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequestImGroupOpenSpaceModel) SetLastMessageI18n(v map[string]*string) *CreateAndDeliverRequestImGroupOpenSpaceModel {
	s.LastMessageI18n = v
	return s
}

func (s *CreateAndDeliverRequestImGroupOpenSpaceModel) SetNotification(v *CreateAndDeliverRequestImGroupOpenSpaceModelNotification) *CreateAndDeliverRequestImGroupOpenSpaceModel {
	s.Notification = v
	return s
}

func (s *CreateAndDeliverRequestImGroupOpenSpaceModel) SetSearchSupport(v *CreateAndDeliverRequestImGroupOpenSpaceModelSearchSupport) *CreateAndDeliverRequestImGroupOpenSpaceModel {
	s.SearchSupport = v
	return s
}

func (s *CreateAndDeliverRequestImGroupOpenSpaceModel) SetSupportForward(v bool) *CreateAndDeliverRequestImGroupOpenSpaceModel {
	s.SupportForward = &v
	return s
}

type CreateAndDeliverRequestImGroupOpenSpaceModelNotification struct {
	AlertContent    *string `json:"alertContent,omitempty" xml:"alertContent,omitempty"`
	NotificationOff *bool   `json:"notificationOff,omitempty" xml:"notificationOff,omitempty"`
}

func (s CreateAndDeliverRequestImGroupOpenSpaceModelNotification) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequestImGroupOpenSpaceModelNotification) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequestImGroupOpenSpaceModelNotification) SetAlertContent(v string) *CreateAndDeliverRequestImGroupOpenSpaceModelNotification {
	s.AlertContent = &v
	return s
}

func (s *CreateAndDeliverRequestImGroupOpenSpaceModelNotification) SetNotificationOff(v bool) *CreateAndDeliverRequestImGroupOpenSpaceModelNotification {
	s.NotificationOff = &v
	return s
}

type CreateAndDeliverRequestImGroupOpenSpaceModelSearchSupport struct {
	SearchDesc     *string `json:"searchDesc,omitempty" xml:"searchDesc,omitempty"`
	SearchIcon     *string `json:"searchIcon,omitempty" xml:"searchIcon,omitempty"`
	SearchTypeName *string `json:"searchTypeName,omitempty" xml:"searchTypeName,omitempty"`
}

func (s CreateAndDeliverRequestImGroupOpenSpaceModelSearchSupport) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequestImGroupOpenSpaceModelSearchSupport) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequestImGroupOpenSpaceModelSearchSupport) SetSearchDesc(v string) *CreateAndDeliverRequestImGroupOpenSpaceModelSearchSupport {
	s.SearchDesc = &v
	return s
}

func (s *CreateAndDeliverRequestImGroupOpenSpaceModelSearchSupport) SetSearchIcon(v string) *CreateAndDeliverRequestImGroupOpenSpaceModelSearchSupport {
	s.SearchIcon = &v
	return s
}

func (s *CreateAndDeliverRequestImGroupOpenSpaceModelSearchSupport) SetSearchTypeName(v string) *CreateAndDeliverRequestImGroupOpenSpaceModelSearchSupport {
	s.SearchTypeName = &v
	return s
}

type CreateAndDeliverRequestImSingleOpenDeliverModel struct {
	// 消息@人，
	AtUserIds map[string]*string `json:"atUserIds,omitempty" xml:"atUserIds,omitempty"`
}

func (s CreateAndDeliverRequestImSingleOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequestImSingleOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequestImSingleOpenDeliverModel) SetAtUserIds(v map[string]*string) *CreateAndDeliverRequestImSingleOpenDeliverModel {
	s.AtUserIds = v
	return s
}

type CreateAndDeliverRequestImSingleOpenSpaceModel struct {
	// 支持国际化的LastMessage
	LastMessageI18n map[string]*string `json:"lastMessageI18n,omitempty" xml:"lastMessageI18n,omitempty"`
	// 通知信息
	Notification *CreateAndDeliverRequestImSingleOpenSpaceModelNotification `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	// 支持卡片消息可被搜索字段
	SearchSupport *CreateAndDeliverRequestImSingleOpenSpaceModelSearchSupport `json:"searchSupport,omitempty" xml:"searchSupport,omitempty" type:"Struct"`
	// 是否支持转发, 默认false
	SupportForward *bool `json:"supportForward,omitempty" xml:"supportForward,omitempty"`
}

func (s CreateAndDeliverRequestImSingleOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequestImSingleOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequestImSingleOpenSpaceModel) SetLastMessageI18n(v map[string]*string) *CreateAndDeliverRequestImSingleOpenSpaceModel {
	s.LastMessageI18n = v
	return s
}

func (s *CreateAndDeliverRequestImSingleOpenSpaceModel) SetNotification(v *CreateAndDeliverRequestImSingleOpenSpaceModelNotification) *CreateAndDeliverRequestImSingleOpenSpaceModel {
	s.Notification = v
	return s
}

func (s *CreateAndDeliverRequestImSingleOpenSpaceModel) SetSearchSupport(v *CreateAndDeliverRequestImSingleOpenSpaceModelSearchSupport) *CreateAndDeliverRequestImSingleOpenSpaceModel {
	s.SearchSupport = v
	return s
}

func (s *CreateAndDeliverRequestImSingleOpenSpaceModel) SetSupportForward(v bool) *CreateAndDeliverRequestImSingleOpenSpaceModel {
	s.SupportForward = &v
	return s
}

type CreateAndDeliverRequestImSingleOpenSpaceModelNotification struct {
	AlertContent    *string `json:"alertContent,omitempty" xml:"alertContent,omitempty"`
	NotificationOff *bool   `json:"notificationOff,omitempty" xml:"notificationOff,omitempty"`
}

func (s CreateAndDeliverRequestImSingleOpenSpaceModelNotification) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequestImSingleOpenSpaceModelNotification) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequestImSingleOpenSpaceModelNotification) SetAlertContent(v string) *CreateAndDeliverRequestImSingleOpenSpaceModelNotification {
	s.AlertContent = &v
	return s
}

func (s *CreateAndDeliverRequestImSingleOpenSpaceModelNotification) SetNotificationOff(v bool) *CreateAndDeliverRequestImSingleOpenSpaceModelNotification {
	s.NotificationOff = &v
	return s
}

type CreateAndDeliverRequestImSingleOpenSpaceModelSearchSupport struct {
	SearchDesc     *string `json:"searchDesc,omitempty" xml:"searchDesc,omitempty"`
	SearchIcon     *string `json:"searchIcon,omitempty" xml:"searchIcon,omitempty"`
	SearchTypeName *string `json:"searchTypeName,omitempty" xml:"searchTypeName,omitempty"`
}

func (s CreateAndDeliverRequestImSingleOpenSpaceModelSearchSupport) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequestImSingleOpenSpaceModelSearchSupport) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequestImSingleOpenSpaceModelSearchSupport) SetSearchDesc(v string) *CreateAndDeliverRequestImSingleOpenSpaceModelSearchSupport {
	s.SearchDesc = &v
	return s
}

func (s *CreateAndDeliverRequestImSingleOpenSpaceModelSearchSupport) SetSearchIcon(v string) *CreateAndDeliverRequestImSingleOpenSpaceModelSearchSupport {
	s.SearchIcon = &v
	return s
}

func (s *CreateAndDeliverRequestImSingleOpenSpaceModelSearchSupport) SetSearchTypeName(v string) *CreateAndDeliverRequestImSingleOpenSpaceModelSearchSupport {
	s.SearchTypeName = &v
	return s
}

type CreateAndDeliverRequestOpenDynamicDataConfig struct {
	// 动态数据替换关系,key是变量名, value是数据源的jsonPath相关配置
	DynamicDataMapping map[string]*OpenDynamicDataConfigDynamicDataMappingValue `json:"dynamicDataMapping,omitempty" xml:"dynamicDataMapping,omitempty"`
	// 动态数据映射类型 (REPLACE_WITHOUT_MAPPING: 直接将动态数据返回，无需根据 key mapping, MAPPING_BY_KEY: 根据创建时的 key 进行 mapping)
	DynamicDataMappingMethod *string `json:"dynamicDataMappingMethod,omitempty" xml:"dynamicDataMappingMethod,omitempty"`
	// 动态数据源配置列表
	DynamicDataSourceConfigs []*CreateAndDeliverRequestOpenDynamicDataConfigDynamicDataSourceConfigs `json:"dynamicDataSourceConfigs,omitempty" xml:"dynamicDataSourceConfigs,omitempty" type:"Repeated"`
}

func (s CreateAndDeliverRequestOpenDynamicDataConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequestOpenDynamicDataConfig) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequestOpenDynamicDataConfig) SetDynamicDataMapping(v map[string]*OpenDynamicDataConfigDynamicDataMappingValue) *CreateAndDeliverRequestOpenDynamicDataConfig {
	s.DynamicDataMapping = v
	return s
}

func (s *CreateAndDeliverRequestOpenDynamicDataConfig) SetDynamicDataMappingMethod(v string) *CreateAndDeliverRequestOpenDynamicDataConfig {
	s.DynamicDataMappingMethod = &v
	return s
}

func (s *CreateAndDeliverRequestOpenDynamicDataConfig) SetDynamicDataSourceConfigs(v []*CreateAndDeliverRequestOpenDynamicDataConfigDynamicDataSourceConfigs) *CreateAndDeliverRequestOpenDynamicDataConfig {
	s.DynamicDataSourceConfigs = v
	return s
}

type CreateAndDeliverRequestOpenDynamicDataConfigDynamicDataSourceConfigs struct {
	// 回调数据源的常量参数
	ConstParams map[string]*string `json:"constParams,omitempty" xml:"constParams,omitempty"`
	// 数据源配置id
	DynamicDataSourceId *string `json:"dynamicDataSourceId,omitempty" xml:"dynamicDataSourceId,omitempty"`
	// 数据源拉取配置
	PullConfig *CreateAndDeliverRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig `json:"pullConfig,omitempty" xml:"pullConfig,omitempty" type:"Struct"`
}

func (s CreateAndDeliverRequestOpenDynamicDataConfigDynamicDataSourceConfigs) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequestOpenDynamicDataConfigDynamicDataSourceConfigs) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequestOpenDynamicDataConfigDynamicDataSourceConfigs) SetConstParams(v map[string]*string) *CreateAndDeliverRequestOpenDynamicDataConfigDynamicDataSourceConfigs {
	s.ConstParams = v
	return s
}

func (s *CreateAndDeliverRequestOpenDynamicDataConfigDynamicDataSourceConfigs) SetDynamicDataSourceId(v string) *CreateAndDeliverRequestOpenDynamicDataConfigDynamicDataSourceConfigs {
	s.DynamicDataSourceId = &v
	return s
}

func (s *CreateAndDeliverRequestOpenDynamicDataConfigDynamicDataSourceConfigs) SetPullConfig(v *CreateAndDeliverRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig) *CreateAndDeliverRequestOpenDynamicDataConfigDynamicDataSourceConfigs {
	s.PullConfig = v
	return s
}

type CreateAndDeliverRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig struct {
	// 间隔
	Interval *int32 `json:"interval,omitempty" xml:"interval,omitempty"`
	// 拉取策略 (NONE: 不拉取,无动态数据, INTERVAL: 间隔拉取, ONCE: 只拉取一次)
	PullStrategy *string `json:"pullStrategy,omitempty" xml:"pullStrategy,omitempty"`
	// 间隔的时间单位 (SECONDS, MINUTES, HOURS, DAYS)
	TimeUnit *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
}

func (s CreateAndDeliverRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig) SetInterval(v int32) *CreateAndDeliverRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig {
	s.Interval = &v
	return s
}

func (s *CreateAndDeliverRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig) SetPullStrategy(v string) *CreateAndDeliverRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig {
	s.PullStrategy = &v
	return s
}

func (s *CreateAndDeliverRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig) SetTimeUnit(v string) *CreateAndDeliverRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig {
	s.TimeUnit = &v
	return s
}

type CreateAndDeliverRequestTopOpenDeliverModel struct {
	// 【必填】过期时间戳
	ExpiredTimeMillis *int64 `json:"expiredTimeMillis,omitempty" xml:"expiredTimeMillis,omitempty"`
	// 可以查看该吊顶卡片的设备
	Platforms []*string `json:"platforms,omitempty" xml:"platforms,omitempty" type:"Repeated"`
	// 可以查看该吊顶卡片的staffId
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s CreateAndDeliverRequestTopOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequestTopOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequestTopOpenDeliverModel) SetExpiredTimeMillis(v int64) *CreateAndDeliverRequestTopOpenDeliverModel {
	s.ExpiredTimeMillis = &v
	return s
}

func (s *CreateAndDeliverRequestTopOpenDeliverModel) SetPlatforms(v []*string) *CreateAndDeliverRequestTopOpenDeliverModel {
	s.Platforms = v
	return s
}

func (s *CreateAndDeliverRequestTopOpenDeliverModel) SetUserIds(v []*string) *CreateAndDeliverRequestTopOpenDeliverModel {
	s.UserIds = v
	return s
}

type CreateAndDeliverRequestTopOpenSpaceModel struct {
	// 【必填】场域类型 (IM: IM, IM_SINGLE: IM单聊, IM_GROUP: IM群聊, ONE_BOX: 群吊顶, COOPERATION_FEED: 协作, WORK_BENCH: 工作台)
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
}

func (s CreateAndDeliverRequestTopOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequestTopOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequestTopOpenSpaceModel) SetSpaceType(v string) *CreateAndDeliverRequestTopOpenSpaceModel {
	s.SpaceType = &v
	return s
}

type CreateAndDeliverRequestWorkBenchOpenDeliverModel struct {
	// 【必填】组件icon对应组件左上角的图标
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 【必填】卡片名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 【必填】卡片组件名
	PluginComponentName *string `json:"pluginComponentName,omitempty" xml:"pluginComponentName,omitempty"`
	// 【必填】卡片预览图
	PreviewUrl *string `json:"previewUrl,omitempty" xml:"previewUrl,omitempty"`
	// 【必填】保持和微应用名称相同
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// 【必填】操作者Id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateAndDeliverRequestWorkBenchOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequestWorkBenchOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequestWorkBenchOpenDeliverModel) SetIcon(v string) *CreateAndDeliverRequestWorkBenchOpenDeliverModel {
	s.Icon = &v
	return s
}

func (s *CreateAndDeliverRequestWorkBenchOpenDeliverModel) SetName(v string) *CreateAndDeliverRequestWorkBenchOpenDeliverModel {
	s.Name = &v
	return s
}

func (s *CreateAndDeliverRequestWorkBenchOpenDeliverModel) SetPluginComponentName(v string) *CreateAndDeliverRequestWorkBenchOpenDeliverModel {
	s.PluginComponentName = &v
	return s
}

func (s *CreateAndDeliverRequestWorkBenchOpenDeliverModel) SetPreviewUrl(v string) *CreateAndDeliverRequestWorkBenchOpenDeliverModel {
	s.PreviewUrl = &v
	return s
}

func (s *CreateAndDeliverRequestWorkBenchOpenDeliverModel) SetProjectName(v string) *CreateAndDeliverRequestWorkBenchOpenDeliverModel {
	s.ProjectName = &v
	return s
}

func (s *CreateAndDeliverRequestWorkBenchOpenDeliverModel) SetUserId(v string) *CreateAndDeliverRequestWorkBenchOpenDeliverModel {
	s.UserId = &v
	return s
}

type CreateAndDeliverRequestWorkBenchOpenSpaceModel struct {
	// 【必填】场域类型 (IM: IM, IM_SINGLE: IM单聊, IM_GROUP: IM群聊, ONE_BOX: 群吊顶, COOPERATION_FEED: 协作, WORK_BENCH: 工作台)
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
}

func (s CreateAndDeliverRequestWorkBenchOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequestWorkBenchOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequestWorkBenchOpenSpaceModel) SetSpaceType(v string) *CreateAndDeliverRequestWorkBenchOpenSpaceModel {
	s.SpaceType = &v
	return s
}

type CreateAndDeliverResponseBody struct {
	Result  *CreateAndDeliverResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                               `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateAndDeliverResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverResponseBody) SetResult(v *CreateAndDeliverResponseBodyResult) *CreateAndDeliverResponseBody {
	s.Result = v
	return s
}

func (s *CreateAndDeliverResponseBody) SetSuccess(v bool) *CreateAndDeliverResponseBody {
	s.Success = &v
	return s
}

type CreateAndDeliverResponseBodyResult struct {
	// 投放结果
	DeliverResults []*CreateAndDeliverResponseBodyResultDeliverResults `json:"deliverResults,omitempty" xml:"deliverResults,omitempty" type:"Repeated"`
	// 外部卡片实例Id
	OutTrackId *string `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
}

func (s CreateAndDeliverResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverResponseBodyResult) SetDeliverResults(v []*CreateAndDeliverResponseBodyResultDeliverResults) *CreateAndDeliverResponseBodyResult {
	s.DeliverResults = v
	return s
}

func (s *CreateAndDeliverResponseBodyResult) SetOutTrackId(v string) *CreateAndDeliverResponseBodyResult {
	s.OutTrackId = &v
	return s
}

type CreateAndDeliverResponseBodyResultDeliverResults struct {
	// 错误信息
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// 场域Id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 场域类型 (IM: IM类型，包括群聊和单聊，仅供返回结果使用, IM_SINGLE: IM单聊, IM_GROUP: IM群聊, ONE_BOX: 群吊顶, COOPERATION_FEED: 协作, WORK_BENCH: 工作台)
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	// 投放成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateAndDeliverResponseBodyResultDeliverResults) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverResponseBodyResultDeliverResults) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverResponseBodyResultDeliverResults) SetErrorMsg(v string) *CreateAndDeliverResponseBodyResultDeliverResults {
	s.ErrorMsg = &v
	return s
}

func (s *CreateAndDeliverResponseBodyResultDeliverResults) SetSpaceId(v string) *CreateAndDeliverResponseBodyResultDeliverResults {
	s.SpaceId = &v
	return s
}

func (s *CreateAndDeliverResponseBodyResultDeliverResults) SetSpaceType(v string) *CreateAndDeliverResponseBodyResultDeliverResults {
	s.SpaceType = &v
	return s
}

func (s *CreateAndDeliverResponseBodyResultDeliverResults) SetSuccess(v bool) *CreateAndDeliverResponseBodyResultDeliverResults {
	s.Success = &v
	return s
}

type CreateAndDeliverResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAndDeliverResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAndDeliverResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverResponse) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverResponse) SetHeaders(v map[string]*string) *CreateAndDeliverResponse {
	s.Headers = v
	return s
}

func (s *CreateAndDeliverResponse) SetBody(v *CreateAndDeliverResponseBody) *CreateAndDeliverResponse {
	s.Body = v
	return s
}

type CreateCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateCardHeaders) GoString() string {
	return s.String()
}

func (s *CreateCardHeaders) SetCommonHeaders(v map[string]*string) *CreateCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateCardHeaders) SetXAcsDingtalkAccessToken(v string) *CreateCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateCardRequest struct {
	// 卡片回调时的路由 Key，用于查询注册的 callbackUrl
	CallbackRouteKey *string `json:"callbackRouteKey,omitempty" xml:"callbackRouteKey,omitempty"`
	// 卡片数据
	CardData *CreateCardRequestCardData `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	// 卡片的模版 Id
	CardTemplateId *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	// 协作场域信息
	CoFeedOpenSpaceModel *CreateCardRequestCoFeedOpenSpaceModel `json:"coFeedOpenSpaceModel,omitempty" xml:"coFeedOpenSpaceModel,omitempty" type:"Struct"`
	// IM 群聊场域信息
	ImGroupOpenSpaceModel *CreateCardRequestImGroupOpenSpaceModel `json:"imGroupOpenSpaceModel,omitempty" xml:"imGroupOpenSpaceModel,omitempty" type:"Struct"`
	// IM 单聊场域信息
	ImSingleOpenSpaceModel *CreateCardRequestImSingleOpenSpaceModel `json:"imSingleOpenSpaceModel,omitempty" xml:"imSingleOpenSpaceModel,omitempty" type:"Struct"`
	// 动态数据源配置
	OpenDynamicDataConfig *CreateCardRequestOpenDynamicDataConfig `json:"openDynamicDataConfig,omitempty" xml:"openDynamicDataConfig,omitempty" type:"Struct"`
	// 唯一标示卡片的外部编码
	OutTrackId *string `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	// 用户的私有数据。
	// ● key：userId
	// ● value：用户私有数据（cardData）
	PrivateData map[string]*PrivateDataValue `json:"privateData,omitempty" xml:"privateData,omitempty"`
	// 吊顶场域信息
	TopOpenSpaceModel *CreateCardRequestTopOpenSpaceModel `json:"topOpenSpaceModel,omitempty" xml:"topOpenSpaceModel,omitempty" type:"Struct"`
	// 卡片创建者的 userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 工作台场域信息
	WorkBenchOpenSpaceModel *CreateCardRequestWorkBenchOpenSpaceModel `json:"workBenchOpenSpaceModel,omitempty" xml:"workBenchOpenSpaceModel,omitempty" type:"Struct"`
}

func (s CreateCardRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCardRequest) GoString() string {
	return s.String()
}

func (s *CreateCardRequest) SetCallbackRouteKey(v string) *CreateCardRequest {
	s.CallbackRouteKey = &v
	return s
}

func (s *CreateCardRequest) SetCardData(v *CreateCardRequestCardData) *CreateCardRequest {
	s.CardData = v
	return s
}

func (s *CreateCardRequest) SetCardTemplateId(v string) *CreateCardRequest {
	s.CardTemplateId = &v
	return s
}

func (s *CreateCardRequest) SetCoFeedOpenSpaceModel(v *CreateCardRequestCoFeedOpenSpaceModel) *CreateCardRequest {
	s.CoFeedOpenSpaceModel = v
	return s
}

func (s *CreateCardRequest) SetImGroupOpenSpaceModel(v *CreateCardRequestImGroupOpenSpaceModel) *CreateCardRequest {
	s.ImGroupOpenSpaceModel = v
	return s
}

func (s *CreateCardRequest) SetImSingleOpenSpaceModel(v *CreateCardRequestImSingleOpenSpaceModel) *CreateCardRequest {
	s.ImSingleOpenSpaceModel = v
	return s
}

func (s *CreateCardRequest) SetOpenDynamicDataConfig(v *CreateCardRequestOpenDynamicDataConfig) *CreateCardRequest {
	s.OpenDynamicDataConfig = v
	return s
}

func (s *CreateCardRequest) SetOutTrackId(v string) *CreateCardRequest {
	s.OutTrackId = &v
	return s
}

func (s *CreateCardRequest) SetPrivateData(v map[string]*PrivateDataValue) *CreateCardRequest {
	s.PrivateData = v
	return s
}

func (s *CreateCardRequest) SetTopOpenSpaceModel(v *CreateCardRequestTopOpenSpaceModel) *CreateCardRequest {
	s.TopOpenSpaceModel = v
	return s
}

func (s *CreateCardRequest) SetUserId(v string) *CreateCardRequest {
	s.UserId = &v
	return s
}

func (s *CreateCardRequest) SetWorkBenchOpenSpaceModel(v *CreateCardRequestWorkBenchOpenSpaceModel) *CreateCardRequest {
	s.WorkBenchOpenSpaceModel = v
	return s
}

type CreateCardRequestCardData struct {
	// 卡片模板内容替换参数，普通文本类型
	// ● key：参数名
	// ● value: 参数值
	CardParamMap map[string]*string `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
}

func (s CreateCardRequestCardData) String() string {
	return tea.Prettify(s)
}

func (s CreateCardRequestCardData) GoString() string {
	return s.String()
}

func (s *CreateCardRequestCardData) SetCardParamMap(v map[string]*string) *CreateCardRequestCardData {
	s.CardParamMap = v
	return s
}

type CreateCardRequestCoFeedOpenSpaceModel struct {
	// 卡片标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateCardRequestCoFeedOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s CreateCardRequestCoFeedOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *CreateCardRequestCoFeedOpenSpaceModel) SetTitle(v string) *CreateCardRequestCoFeedOpenSpaceModel {
	s.Title = &v
	return s
}

type CreateCardRequestImGroupOpenSpaceModel struct {
	// 支持国际化的LastMessage
	// key为语言枚举值，value为lastMessage内容
	// 目前支持的语言枚举值：
	// 简体中文: ZH_CN
	// 繁体中文: ZH_TW
	// 英文： EN_US
	// 日语: JA_JP
	// 越南语: VI_VN
	LastMessageI18n map[string]*string `json:"lastMessageI18n,omitempty" xml:"lastMessageI18n,omitempty"`
	// 卡片的通知属性信息
	Notification *CreateCardRequestImGroupOpenSpaceModelNotification `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	// 支持卡片消息可被搜索字段
	SearchSupport *CreateCardRequestImGroupOpenSpaceModelSearchSupport `json:"searchSupport,omitempty" xml:"searchSupport,omitempty" type:"Struct"`
	// 是否支持转发, 默认 false
	SupportForward *bool `json:"supportForward,omitempty" xml:"supportForward,omitempty"`
}

func (s CreateCardRequestImGroupOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s CreateCardRequestImGroupOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *CreateCardRequestImGroupOpenSpaceModel) SetLastMessageI18n(v map[string]*string) *CreateCardRequestImGroupOpenSpaceModel {
	s.LastMessageI18n = v
	return s
}

func (s *CreateCardRequestImGroupOpenSpaceModel) SetNotification(v *CreateCardRequestImGroupOpenSpaceModelNotification) *CreateCardRequestImGroupOpenSpaceModel {
	s.Notification = v
	return s
}

func (s *CreateCardRequestImGroupOpenSpaceModel) SetSearchSupport(v *CreateCardRequestImGroupOpenSpaceModelSearchSupport) *CreateCardRequestImGroupOpenSpaceModel {
	s.SearchSupport = v
	return s
}

func (s *CreateCardRequestImGroupOpenSpaceModel) SetSupportForward(v bool) *CreateCardRequestImGroupOpenSpaceModel {
	s.SupportForward = &v
	return s
}

type CreateCardRequestImGroupOpenSpaceModelNotification struct {
	// 【条件必填】通知内容
	// 若不填写则使用默认文案：如你收到1条新消息
	AlertContent *string `json:"alertContent,omitempty" xml:"alertContent,omitempty"`
	// 是否推送通知，默认为 false
	NotificationOff *bool `json:"notificationOff,omitempty" xml:"notificationOff,omitempty"`
}

func (s CreateCardRequestImGroupOpenSpaceModelNotification) String() string {
	return tea.Prettify(s)
}

func (s CreateCardRequestImGroupOpenSpaceModelNotification) GoString() string {
	return s.String()
}

func (s *CreateCardRequestImGroupOpenSpaceModelNotification) SetAlertContent(v string) *CreateCardRequestImGroupOpenSpaceModelNotification {
	s.AlertContent = &v
	return s
}

func (s *CreateCardRequestImGroupOpenSpaceModelNotification) SetNotificationOff(v bool) *CreateCardRequestImGroupOpenSpaceModelNotification {
	s.NotificationOff = &v
	return s
}

type CreateCardRequestImGroupOpenSpaceModelSearchSupport struct {
	//  【条件必填】供消息展示与搜索的字段
	//  【注意】最大限制200个字符，超过存储截断200
	SearchDesc *string `json:"searchDesc,omitempty" xml:"searchDesc,omitempty"`
	// 类型的icon，供搜索展示使用
	SearchIcon *string `json:"searchIcon,omitempty" xml:"searchIcon,omitempty"`
	// 卡片类型名
	SearchTypeName *string `json:"searchTypeName,omitempty" xml:"searchTypeName,omitempty"`
}

func (s CreateCardRequestImGroupOpenSpaceModelSearchSupport) String() string {
	return tea.Prettify(s)
}

func (s CreateCardRequestImGroupOpenSpaceModelSearchSupport) GoString() string {
	return s.String()
}

func (s *CreateCardRequestImGroupOpenSpaceModelSearchSupport) SetSearchDesc(v string) *CreateCardRequestImGroupOpenSpaceModelSearchSupport {
	s.SearchDesc = &v
	return s
}

func (s *CreateCardRequestImGroupOpenSpaceModelSearchSupport) SetSearchIcon(v string) *CreateCardRequestImGroupOpenSpaceModelSearchSupport {
	s.SearchIcon = &v
	return s
}

func (s *CreateCardRequestImGroupOpenSpaceModelSearchSupport) SetSearchTypeName(v string) *CreateCardRequestImGroupOpenSpaceModelSearchSupport {
	s.SearchTypeName = &v
	return s
}

type CreateCardRequestImSingleOpenSpaceModel struct {
	// 支持国际化的LastMessage
	// key为语言枚举值，value为lastMessage内容
	// 目前支持的语言枚举值：
	// 简体中文: ZH_CN
	// 繁体中文: ZH_TW
	// 英文： EN_US
	// 日语: JA_JP
	// 越南语: VI_VN
	LastMessageI18n map[string]*string `json:"lastMessageI18n,omitempty" xml:"lastMessageI18n,omitempty"`
	// 卡片的通知属性信息
	Notification *CreateCardRequestImSingleOpenSpaceModelNotification `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	// 支持卡片消息可被搜索字段
	SearchSupport *CreateCardRequestImSingleOpenSpaceModelSearchSupport `json:"searchSupport,omitempty" xml:"searchSupport,omitempty" type:"Struct"`
	// 是否支持转发, 默认 false
	SupportForward *bool `json:"supportForward,omitempty" xml:"supportForward,omitempty"`
}

func (s CreateCardRequestImSingleOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s CreateCardRequestImSingleOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *CreateCardRequestImSingleOpenSpaceModel) SetLastMessageI18n(v map[string]*string) *CreateCardRequestImSingleOpenSpaceModel {
	s.LastMessageI18n = v
	return s
}

func (s *CreateCardRequestImSingleOpenSpaceModel) SetNotification(v *CreateCardRequestImSingleOpenSpaceModelNotification) *CreateCardRequestImSingleOpenSpaceModel {
	s.Notification = v
	return s
}

func (s *CreateCardRequestImSingleOpenSpaceModel) SetSearchSupport(v *CreateCardRequestImSingleOpenSpaceModelSearchSupport) *CreateCardRequestImSingleOpenSpaceModel {
	s.SearchSupport = v
	return s
}

func (s *CreateCardRequestImSingleOpenSpaceModel) SetSupportForward(v bool) *CreateCardRequestImSingleOpenSpaceModel {
	s.SupportForward = &v
	return s
}

type CreateCardRequestImSingleOpenSpaceModelNotification struct {
	// 【条件必填】通知内容
	// 若不填写则使用默认文案：如你收到1条新消息
	AlertContent *string `json:"alertContent,omitempty" xml:"alertContent,omitempty"`
	// 是否推送通知，默认为 false
	NotificationOff *bool `json:"notificationOff,omitempty" xml:"notificationOff,omitempty"`
}

func (s CreateCardRequestImSingleOpenSpaceModelNotification) String() string {
	return tea.Prettify(s)
}

func (s CreateCardRequestImSingleOpenSpaceModelNotification) GoString() string {
	return s.String()
}

func (s *CreateCardRequestImSingleOpenSpaceModelNotification) SetAlertContent(v string) *CreateCardRequestImSingleOpenSpaceModelNotification {
	s.AlertContent = &v
	return s
}

func (s *CreateCardRequestImSingleOpenSpaceModelNotification) SetNotificationOff(v bool) *CreateCardRequestImSingleOpenSpaceModelNotification {
	s.NotificationOff = &v
	return s
}

type CreateCardRequestImSingleOpenSpaceModelSearchSupport struct {
	// 【条件必填】供消息展示与搜索的字段
	//  【注意】最大限制200个字符，超过存储截断200
	SearchDesc *string `json:"searchDesc,omitempty" xml:"searchDesc,omitempty"`
	// 类型的icon，供搜索展示使用
	SearchIcon *string `json:"searchIcon,omitempty" xml:"searchIcon,omitempty"`
	// 卡片类型名
	SearchTypeName *string `json:"searchTypeName,omitempty" xml:"searchTypeName,omitempty"`
}

func (s CreateCardRequestImSingleOpenSpaceModelSearchSupport) String() string {
	return tea.Prettify(s)
}

func (s CreateCardRequestImSingleOpenSpaceModelSearchSupport) GoString() string {
	return s.String()
}

func (s *CreateCardRequestImSingleOpenSpaceModelSearchSupport) SetSearchDesc(v string) *CreateCardRequestImSingleOpenSpaceModelSearchSupport {
	s.SearchDesc = &v
	return s
}

func (s *CreateCardRequestImSingleOpenSpaceModelSearchSupport) SetSearchIcon(v string) *CreateCardRequestImSingleOpenSpaceModelSearchSupport {
	s.SearchIcon = &v
	return s
}

func (s *CreateCardRequestImSingleOpenSpaceModelSearchSupport) SetSearchTypeName(v string) *CreateCardRequestImSingleOpenSpaceModelSearchSupport {
	s.SearchTypeName = &v
	return s
}

type CreateCardRequestOpenDynamicDataConfig struct {
	// 动态数据替换关系
	// ● key：变量名
	// ● value：数据映射的配置
	DynamicDataMapping map[string]*OpenDynamicDataConfigDynamicDataMappingValue `json:"dynamicDataMapping,omitempty" xml:"dynamicDataMapping,omitempty"`
	// 动态数据映射方法，可选值
	// ● REPLACE_WITHOUT_MAPPING：直接返回动态数据
	// ● MAPPING_BY_KEY：根据指定的规则，将返回数据映射到卡片数据上
	// 默认值：REPLACE_WITHOUT_MAPPING
	DynamicDataMappingMethod *string `json:"dynamicDataMappingMethod,omitempty" xml:"dynamicDataMappingMethod,omitempty"`
	// 动态数据源配置列表
	DynamicDataSourceConfigs []*CreateCardRequestOpenDynamicDataConfigDynamicDataSourceConfigs `json:"dynamicDataSourceConfigs,omitempty" xml:"dynamicDataSourceConfigs,omitempty" type:"Repeated"`
}

func (s CreateCardRequestOpenDynamicDataConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateCardRequestOpenDynamicDataConfig) GoString() string {
	return s.String()
}

func (s *CreateCardRequestOpenDynamicDataConfig) SetDynamicDataMapping(v map[string]*OpenDynamicDataConfigDynamicDataMappingValue) *CreateCardRequestOpenDynamicDataConfig {
	s.DynamicDataMapping = v
	return s
}

func (s *CreateCardRequestOpenDynamicDataConfig) SetDynamicDataMappingMethod(v string) *CreateCardRequestOpenDynamicDataConfig {
	s.DynamicDataMappingMethod = &v
	return s
}

func (s *CreateCardRequestOpenDynamicDataConfig) SetDynamicDataSourceConfigs(v []*CreateCardRequestOpenDynamicDataConfigDynamicDataSourceConfigs) *CreateCardRequestOpenDynamicDataConfig {
	s.DynamicDataSourceConfigs = v
	return s
}

type CreateCardRequestOpenDynamicDataConfigDynamicDataSourceConfigs struct {
	// 回调数据源的常量参数
	ConstParams map[string]*string `json:"constParams,omitempty" xml:"constParams,omitempty"`
	// 【条件必填】数据源的唯一 ID
	DynamicDataSourceId *string `json:"dynamicDataSourceId,omitempty" xml:"dynamicDataSourceId,omitempty"`
	// 【条件必填】数据源拉取配置
	PullConfig *CreateCardRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig `json:"pullConfig,omitempty" xml:"pullConfig,omitempty" type:"Struct"`
}

func (s CreateCardRequestOpenDynamicDataConfigDynamicDataSourceConfigs) String() string {
	return tea.Prettify(s)
}

func (s CreateCardRequestOpenDynamicDataConfigDynamicDataSourceConfigs) GoString() string {
	return s.String()
}

func (s *CreateCardRequestOpenDynamicDataConfigDynamicDataSourceConfigs) SetConstParams(v map[string]*string) *CreateCardRequestOpenDynamicDataConfigDynamicDataSourceConfigs {
	s.ConstParams = v
	return s
}

func (s *CreateCardRequestOpenDynamicDataConfigDynamicDataSourceConfigs) SetDynamicDataSourceId(v string) *CreateCardRequestOpenDynamicDataConfigDynamicDataSourceConfigs {
	s.DynamicDataSourceId = &v
	return s
}

func (s *CreateCardRequestOpenDynamicDataConfigDynamicDataSourceConfigs) SetPullConfig(v *CreateCardRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig) *CreateCardRequestOpenDynamicDataConfigDynamicDataSourceConfigs {
	s.PullConfig = v
	return s
}

type CreateCardRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig struct {
	// 拉取的间隔时间，只在将 pullStrategy 设置为 INTERVAL 的时候生效
	Interval *int32 `json:"interval,omitempty" xml:"interval,omitempty"`
	// 【条件必填】拉取策略，可选值：
	// ● NONE：不拉取，无动态数据
	// ● INTERVAL：间隔拉取
	// ● ONCE：只拉取一次
	PullStrategy *string `json:"pullStrategy,omitempty" xml:"pullStrategy,omitempty"`
	// 拉取的间隔时间的单位，只在将 pullStrategy 设置为 INTERVAL 的时候生效。 可选值：
	// ● SECONDS
	// ● MINUTES
	// ● HOURS
	// ● DAYS
	TimeUnit *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
}

func (s CreateCardRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateCardRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig) GoString() string {
	return s.String()
}

func (s *CreateCardRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig) SetInterval(v int32) *CreateCardRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig {
	s.Interval = &v
	return s
}

func (s *CreateCardRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig) SetPullStrategy(v string) *CreateCardRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig {
	s.PullStrategy = &v
	return s
}

func (s *CreateCardRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig) SetTimeUnit(v string) *CreateCardRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig {
	s.TimeUnit = &v
	return s
}

type CreateCardRequestTopOpenSpaceModel struct {
	// 吊顶无其他场域属性，通过增加spaeType使卡片支持吊顶场域；吊顶对应spaceType为: ONE_BOX
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
}

func (s CreateCardRequestTopOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s CreateCardRequestTopOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *CreateCardRequestTopOpenSpaceModel) SetSpaceType(v string) *CreateCardRequestTopOpenSpaceModel {
	s.SpaceType = &v
	return s
}

type CreateCardRequestWorkBenchOpenSpaceModel struct {
	// 工作台无其他场域属性，通过增加spaeType使卡片支持工作台场域;工作台对应spaceType为: WORK_BENCH
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
}

func (s CreateCardRequestWorkBenchOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s CreateCardRequestWorkBenchOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *CreateCardRequestWorkBenchOpenSpaceModel) SetSpaceType(v string) *CreateCardRequestWorkBenchOpenSpaceModel {
	s.SpaceType = &v
	return s
}

type CreateCardResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCardResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCardResponseBody) SetResult(v string) *CreateCardResponseBody {
	s.Result = &v
	return s
}

func (s *CreateCardResponseBody) SetSuccess(v bool) *CreateCardResponseBody {
	s.Success = &v
	return s
}

type CreateCardResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCardResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCardResponse) GoString() string {
	return s.String()
}

func (s *CreateCardResponse) SetHeaders(v map[string]*string) *CreateCardResponse {
	s.Headers = v
	return s
}

func (s *CreateCardResponse) SetBody(v *CreateCardResponseBody) *CreateCardResponse {
	s.Body = v
	return s
}

type DeliverCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeliverCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeliverCardHeaders) GoString() string {
	return s.String()
}

func (s *DeliverCardHeaders) SetCommonHeaders(v map[string]*string) *DeliverCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeliverCardHeaders) SetXAcsDingtalkAccessToken(v string) *DeliverCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeliverCardRequest struct {
	// 协作投放参数
	CoFeedOpenDeliverModel *DeliverCardRequestCoFeedOpenDeliverModel `json:"coFeedOpenDeliverModel,omitempty" xml:"coFeedOpenDeliverModel,omitempty" type:"Struct"`
	// 群聊投放参数
	ImGroupOpenDeliverModel *DeliverCardRequestImGroupOpenDeliverModel `json:"imGroupOpenDeliverModel,omitempty" xml:"imGroupOpenDeliverModel,omitempty" type:"Struct"`
	// 单聊场域投放参数
	ImSingleOpenDeliverModel *DeliverCardRequestImSingleOpenDeliverModel `json:"imSingleOpenDeliverModel,omitempty" xml:"imSingleOpenDeliverModel,omitempty" type:"Struct"`
	// dt.card//spaceType.spaceId;spaceType.spaceId
	OpenSpaceId *string `json:"openSpaceId,omitempty" xml:"openSpaceId,omitempty"`
	// 外部卡片实例Id
	OutTrackId *string `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	// 吊顶投放参数
	TopOpenDeliverModel *DeliverCardRequestTopOpenDeliverModel `json:"topOpenDeliverModel,omitempty" xml:"topOpenDeliverModel,omitempty" type:"Struct"`
	// 工作台投放参数
	WorkBenchOpenDeliverModel *DeliverCardRequestWorkBenchOpenDeliverModel `json:"workBenchOpenDeliverModel,omitempty" xml:"workBenchOpenDeliverModel,omitempty" type:"Struct"`
}

func (s DeliverCardRequest) String() string {
	return tea.Prettify(s)
}

func (s DeliverCardRequest) GoString() string {
	return s.String()
}

func (s *DeliverCardRequest) SetCoFeedOpenDeliverModel(v *DeliverCardRequestCoFeedOpenDeliverModel) *DeliverCardRequest {
	s.CoFeedOpenDeliverModel = v
	return s
}

func (s *DeliverCardRequest) SetImGroupOpenDeliverModel(v *DeliverCardRequestImGroupOpenDeliverModel) *DeliverCardRequest {
	s.ImGroupOpenDeliverModel = v
	return s
}

func (s *DeliverCardRequest) SetImSingleOpenDeliverModel(v *DeliverCardRequestImSingleOpenDeliverModel) *DeliverCardRequest {
	s.ImSingleOpenDeliverModel = v
	return s
}

func (s *DeliverCardRequest) SetOpenSpaceId(v string) *DeliverCardRequest {
	s.OpenSpaceId = &v
	return s
}

func (s *DeliverCardRequest) SetOutTrackId(v string) *DeliverCardRequest {
	s.OutTrackId = &v
	return s
}

func (s *DeliverCardRequest) SetTopOpenDeliverModel(v *DeliverCardRequestTopOpenDeliverModel) *DeliverCardRequest {
	s.TopOpenDeliverModel = v
	return s
}

func (s *DeliverCardRequest) SetWorkBenchOpenDeliverModel(v *DeliverCardRequestWorkBenchOpenDeliverModel) *DeliverCardRequest {
	s.WorkBenchOpenDeliverModel = v
	return s
}

type DeliverCardRequestCoFeedOpenDeliverModel struct {
	// 【必填】业务标识
	BizTag *string `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	// 【必填】协作场域下的排序时间
	GmtTimeLine *int64 `json:"gmtTimeLine,omitempty" xml:"gmtTimeLine,omitempty"`
}

func (s DeliverCardRequestCoFeedOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s DeliverCardRequestCoFeedOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *DeliverCardRequestCoFeedOpenDeliverModel) SetBizTag(v string) *DeliverCardRequestCoFeedOpenDeliverModel {
	s.BizTag = &v
	return s
}

func (s *DeliverCardRequestCoFeedOpenDeliverModel) SetGmtTimeLine(v int64) *DeliverCardRequestCoFeedOpenDeliverModel {
	s.GmtTimeLine = &v
	return s
}

type DeliverCardRequestImGroupOpenDeliverModel struct {
	// 消息@人，
	AtUserIds map[string]*string `json:"atUserIds,omitempty" xml:"atUserIds,omitempty"`
	// 指定接收者
	Recipients []*string `json:"recipients,omitempty" xml:"recipients,omitempty" type:"Repeated"`
	// 机器人的code
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s DeliverCardRequestImGroupOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s DeliverCardRequestImGroupOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *DeliverCardRequestImGroupOpenDeliverModel) SetAtUserIds(v map[string]*string) *DeliverCardRequestImGroupOpenDeliverModel {
	s.AtUserIds = v
	return s
}

func (s *DeliverCardRequestImGroupOpenDeliverModel) SetRecipients(v []*string) *DeliverCardRequestImGroupOpenDeliverModel {
	s.Recipients = v
	return s
}

func (s *DeliverCardRequestImGroupOpenDeliverModel) SetRobotCode(v string) *DeliverCardRequestImGroupOpenDeliverModel {
	s.RobotCode = &v
	return s
}

type DeliverCardRequestImSingleOpenDeliverModel struct {
	// 消息@人，
	AtUserIds map[string]*string `json:"atUserIds,omitempty" xml:"atUserIds,omitempty"`
}

func (s DeliverCardRequestImSingleOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s DeliverCardRequestImSingleOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *DeliverCardRequestImSingleOpenDeliverModel) SetAtUserIds(v map[string]*string) *DeliverCardRequestImSingleOpenDeliverModel {
	s.AtUserIds = v
	return s
}

type DeliverCardRequestTopOpenDeliverModel struct {
	// 【必填】过期时间戳
	ExpiredTimeMills *int64 `json:"expiredTimeMills,omitempty" xml:"expiredTimeMills,omitempty"`
	// 可以查看该吊顶卡片的设备
	Platforms []*string `json:"platforms,omitempty" xml:"platforms,omitempty" type:"Repeated"`
	// 可以查看该吊顶卡片的staffId
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s DeliverCardRequestTopOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s DeliverCardRequestTopOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *DeliverCardRequestTopOpenDeliverModel) SetExpiredTimeMills(v int64) *DeliverCardRequestTopOpenDeliverModel {
	s.ExpiredTimeMills = &v
	return s
}

func (s *DeliverCardRequestTopOpenDeliverModel) SetPlatforms(v []*string) *DeliverCardRequestTopOpenDeliverModel {
	s.Platforms = v
	return s
}

func (s *DeliverCardRequestTopOpenDeliverModel) SetUserIds(v []*string) *DeliverCardRequestTopOpenDeliverModel {
	s.UserIds = v
	return s
}

type DeliverCardRequestWorkBenchOpenDeliverModel struct {
	// 【必填】组件icon对应组件左上角的图标
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 【必填】卡片名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 【必填】卡片组件名
	PluginComponentName *string `json:"pluginComponentName,omitempty" xml:"pluginComponentName,omitempty"`
	// 【必填】卡片预览图
	PreviewUrl *string `json:"previewUrl,omitempty" xml:"previewUrl,omitempty"`
	// 【必填】保持和微应用名称相同
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// 【必填】操作者Id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeliverCardRequestWorkBenchOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s DeliverCardRequestWorkBenchOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *DeliverCardRequestWorkBenchOpenDeliverModel) SetIcon(v string) *DeliverCardRequestWorkBenchOpenDeliverModel {
	s.Icon = &v
	return s
}

func (s *DeliverCardRequestWorkBenchOpenDeliverModel) SetName(v string) *DeliverCardRequestWorkBenchOpenDeliverModel {
	s.Name = &v
	return s
}

func (s *DeliverCardRequestWorkBenchOpenDeliverModel) SetPluginComponentName(v string) *DeliverCardRequestWorkBenchOpenDeliverModel {
	s.PluginComponentName = &v
	return s
}

func (s *DeliverCardRequestWorkBenchOpenDeliverModel) SetPreviewUrl(v string) *DeliverCardRequestWorkBenchOpenDeliverModel {
	s.PreviewUrl = &v
	return s
}

func (s *DeliverCardRequestWorkBenchOpenDeliverModel) SetProjectName(v string) *DeliverCardRequestWorkBenchOpenDeliverModel {
	s.ProjectName = &v
	return s
}

func (s *DeliverCardRequestWorkBenchOpenDeliverModel) SetUserId(v string) *DeliverCardRequestWorkBenchOpenDeliverModel {
	s.UserId = &v
	return s
}

type DeliverCardResponseBody struct {
	Result  []*DeliverCardResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success *bool                            `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeliverCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeliverCardResponseBody) GoString() string {
	return s.String()
}

func (s *DeliverCardResponseBody) SetResult(v []*DeliverCardResponseBodyResult) *DeliverCardResponseBody {
	s.Result = v
	return s
}

func (s *DeliverCardResponseBody) SetSuccess(v bool) *DeliverCardResponseBody {
	s.Success = &v
	return s
}

type DeliverCardResponseBodyResult struct {
	// 场域Id
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// 场域类型 (IM: IM, IM_SINGLE: IM单聊, IM_GROUP: IM群聊, ONE_BOX: 群吊顶, COOPERATION_FEED: 协作, WORK_BENCH: 工作台)
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	// 投放成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeliverCardResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeliverCardResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeliverCardResponseBodyResult) SetSpaceId(v string) *DeliverCardResponseBodyResult {
	s.SpaceId = &v
	return s
}

func (s *DeliverCardResponseBodyResult) SetSpaceType(v string) *DeliverCardResponseBodyResult {
	s.SpaceType = &v
	return s
}

func (s *DeliverCardResponseBodyResult) SetSuccess(v bool) *DeliverCardResponseBodyResult {
	s.Success = &v
	return s
}

type DeliverCardResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeliverCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeliverCardResponse) String() string {
	return tea.Prettify(s)
}

func (s DeliverCardResponse) GoString() string {
	return s.String()
}

func (s *DeliverCardResponse) SetHeaders(v map[string]*string) *DeliverCardResponse {
	s.Headers = v
	return s
}

func (s *DeliverCardResponse) SetBody(v *DeliverCardResponseBody) *DeliverCardResponse {
	s.Body = v
	return s
}

type RegisterCallbackHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RegisterCallbackHeaders) String() string {
	return tea.Prettify(s)
}

func (s RegisterCallbackHeaders) GoString() string {
	return s.String()
}

func (s *RegisterCallbackHeaders) SetCommonHeaders(v map[string]*string) *RegisterCallbackHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RegisterCallbackHeaders) SetXAcsDingtalkAccessToken(v string) *RegisterCallbackHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RegisterCallbackRequest struct {
	// 加密密钥用于校验来源
	ApiSecret *string `json:"apiSecret,omitempty" xml:"apiSecret,omitempty"`
	// callbackUrl 的路由 Key，一个 callbackRouteKey 可以映射一个 callbackUrl
	CallbackRouteKey *string `json:"callbackRouteKey,omitempty" xml:"callbackRouteKey,omitempty"`
	// 注册的回调 URL
	CallbackUrl *string `json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`
	// 是否强制覆盖更新，默认 false
	ForceUpdate *bool `json:"forceUpdate,omitempty" xml:"forceUpdate,omitempty"`
}

func (s RegisterCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterCallbackRequest) GoString() string {
	return s.String()
}

func (s *RegisterCallbackRequest) SetApiSecret(v string) *RegisterCallbackRequest {
	s.ApiSecret = &v
	return s
}

func (s *RegisterCallbackRequest) SetCallbackRouteKey(v string) *RegisterCallbackRequest {
	s.CallbackRouteKey = &v
	return s
}

func (s *RegisterCallbackRequest) SetCallbackUrl(v string) *RegisterCallbackRequest {
	s.CallbackUrl = &v
	return s
}

func (s *RegisterCallbackRequest) SetForceUpdate(v bool) *RegisterCallbackRequest {
	s.ForceUpdate = &v
	return s
}

type RegisterCallbackResponseBody struct {
	Result  *RegisterCallbackResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                               `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RegisterCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterCallbackResponseBody) SetResult(v *RegisterCallbackResponseBodyResult) *RegisterCallbackResponseBody {
	s.Result = v
	return s
}

func (s *RegisterCallbackResponseBody) SetSuccess(v bool) *RegisterCallbackResponseBody {
	s.Success = &v
	return s
}

type RegisterCallbackResponseBodyResult struct {
	// api 签名密钥
	ApiSecret *string `json:"apiSecret,omitempty" xml:"apiSecret,omitempty"`
	// ISV 接受动态卡片点击的回调地址
	CallbackUrl *string `json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`
}

func (s RegisterCallbackResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s RegisterCallbackResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RegisterCallbackResponseBodyResult) SetApiSecret(v string) *RegisterCallbackResponseBodyResult {
	s.ApiSecret = &v
	return s
}

func (s *RegisterCallbackResponseBodyResult) SetCallbackUrl(v string) *RegisterCallbackResponseBodyResult {
	s.CallbackUrl = &v
	return s
}

type RegisterCallbackResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterCallbackResponse) GoString() string {
	return s.String()
}

func (s *RegisterCallbackResponse) SetHeaders(v map[string]*string) *RegisterCallbackResponse {
	s.Headers = v
	return s
}

func (s *RegisterCallbackResponse) SetBody(v *RegisterCallbackResponseBody) *RegisterCallbackResponse {
	s.Body = v
	return s
}

type UpdateCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateCardHeaders) GoString() string {
	return s.String()
}

func (s *UpdateCardHeaders) SetCommonHeaders(v map[string]*string) *UpdateCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateCardHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateCardRequest struct {
	// 卡片数据
	CardData *UpdateCardRequestCardData `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	// 卡片更新选项
	CardUpdateOptions *UpdateCardRequestCardUpdateOptions `json:"cardUpdateOptions,omitempty" xml:"cardUpdateOptions,omitempty" type:"Struct"`
	// 【必填】外部卡片实例Id
	OutTrackId *string `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	// 用户的私有数据。
	// ● key：userId
	// ● value：用户私有数据（cardData）
	PrivateData map[string]*PrivateDataValue `json:"privateData,omitempty" xml:"privateData,omitempty"`
}

func (s UpdateCardRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCardRequest) GoString() string {
	return s.String()
}

func (s *UpdateCardRequest) SetCardData(v *UpdateCardRequestCardData) *UpdateCardRequest {
	s.CardData = v
	return s
}

func (s *UpdateCardRequest) SetCardUpdateOptions(v *UpdateCardRequestCardUpdateOptions) *UpdateCardRequest {
	s.CardUpdateOptions = v
	return s
}

func (s *UpdateCardRequest) SetOutTrackId(v string) *UpdateCardRequest {
	s.OutTrackId = &v
	return s
}

func (s *UpdateCardRequest) SetPrivateData(v map[string]*PrivateDataValue) *UpdateCardRequest {
	s.PrivateData = v
	return s
}

type UpdateCardRequestCardData struct {
	// 卡片模板内容替换参数，普通文本类型
	// ● key：参数名
	// ● value: 参数值
	CardParamMap map[string]*string `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
}

func (s UpdateCardRequestCardData) String() string {
	return tea.Prettify(s)
}

func (s UpdateCardRequestCardData) GoString() string {
	return s.String()
}

func (s *UpdateCardRequestCardData) SetCardParamMap(v map[string]*string) *UpdateCardRequestCardData {
	s.CardParamMap = v
	return s
}

type UpdateCardRequestCardUpdateOptions struct {
	// 按 key 更新 cardData 数据，不填默认覆盖更新。
	UpdateCardDataByKey *bool `json:"updateCardDataByKey,omitempty" xml:"updateCardDataByKey,omitempty"`
	// 【可选】按key更新privateData用户数据，不填默认覆盖更新。
	UpdatePrivateDataByKey *bool `json:"updatePrivateDataByKey,omitempty" xml:"updatePrivateDataByKey,omitempty"`
}

func (s UpdateCardRequestCardUpdateOptions) String() string {
	return tea.Prettify(s)
}

func (s UpdateCardRequestCardUpdateOptions) GoString() string {
	return s.String()
}

func (s *UpdateCardRequestCardUpdateOptions) SetUpdateCardDataByKey(v bool) *UpdateCardRequestCardUpdateOptions {
	s.UpdateCardDataByKey = &v
	return s
}

func (s *UpdateCardRequestCardUpdateOptions) SetUpdatePrivateDataByKey(v bool) *UpdateCardRequestCardUpdateOptions {
	s.UpdatePrivateDataByKey = &v
	return s
}

type UpdateCardResponseBody struct {
	Result  *bool `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCardResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCardResponseBody) SetResult(v bool) *UpdateCardResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateCardResponseBody) SetSuccess(v bool) *UpdateCardResponseBody {
	s.Success = &v
	return s
}

type UpdateCardResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCardResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCardResponse) GoString() string {
	return s.String()
}

func (s *UpdateCardResponse) SetHeaders(v map[string]*string) *UpdateCardResponse {
	s.Headers = v
	return s
}

func (s *UpdateCardResponse) SetBody(v *UpdateCardResponseBody) *UpdateCardResponse {
	s.Body = v
	return s
}

type OpenDynamicDataConfigDynamicDataMappingValue struct {
	// jsonPath
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// 值的类型 (STRING: String, ARRAY: 数组, OBJECT: 对象, CHART: 图表, TABLE: 表格, INDICATOR: 指标卡)
	DynamicDataValueType *string `json:"dynamicDataValueType,omitempty" xml:"dynamicDataValueType,omitempty"`
}

func (s OpenDynamicDataConfigDynamicDataMappingValue) String() string {
	return tea.Prettify(s)
}

func (s OpenDynamicDataConfigDynamicDataMappingValue) GoString() string {
	return s.String()
}

func (s *OpenDynamicDataConfigDynamicDataMappingValue) SetPath(v string) *OpenDynamicDataConfigDynamicDataMappingValue {
	s.Path = &v
	return s
}

func (s *OpenDynamicDataConfigDynamicDataMappingValue) SetDynamicDataValueType(v string) *OpenDynamicDataConfigDynamicDataMappingValue {
	s.DynamicDataValueType = &v
	return s
}

type PrivateDataValue struct {
	// 卡片模板-文本内容参数
	CardParamMap map[string]*string `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
}

func (s PrivateDataValue) String() string {
	return tea.Prettify(s)
}

func (s PrivateDataValue) GoString() string {
	return s.String()
}

func (s *PrivateDataValue) SetCardParamMap(v map[string]*string) *PrivateDataValue {
	s.CardParamMap = v
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

func (client *Client) AppendSpace(request *AppendSpaceRequest) (_result *AppendSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AppendSpaceHeaders{}
	_result = &AppendSpaceResponse{}
	_body, _err := client.AppendSpaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AppendSpaceWithOptions(request *AppendSpaceRequest, headers *AppendSpaceHeaders, runtime *util.RuntimeOptions) (_result *AppendSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CoFeedOpenSpaceModel))) {
		body["coFeedOpenSpaceModel"] = request.CoFeedOpenSpaceModel
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.ImGroupOpenSpaceModel))) {
		body["imGroupOpenSpaceModel"] = request.ImGroupOpenSpaceModel
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.ImSingleOpenSpaceModel))) {
		body["imSingleOpenSpaceModel"] = request.ImSingleOpenSpaceModel
	}

	if !tea.BoolValue(util.IsUnset(request.OutTrackId)) {
		body["outTrackId"] = request.OutTrackId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TopOpenSpaceModel))) {
		body["topOpenSpaceModel"] = request.TopOpenSpaceModel
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.WorkBenchOpenSpaceModel))) {
		body["workBenchOpenSpaceModel"] = request.WorkBenchOpenSpaceModel
	}

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
	_result = &AppendSpaceResponse{}
	_body, _err := client.DoROARequest(tea.String("AppendSpace"), tea.String("card_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/card/instances/spaces"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAndDeliver(request *CreateAndDeliverRequest) (_result *CreateAndDeliverResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateAndDeliverHeaders{}
	_result = &CreateAndDeliverResponse{}
	_body, _err := client.CreateAndDeliverWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAndDeliverWithOptions(request *CreateAndDeliverRequest, headers *CreateAndDeliverHeaders, runtime *util.RuntimeOptions) (_result *CreateAndDeliverResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallbackRouteKey)) {
		body["callbackRouteKey"] = request.CallbackRouteKey
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CardData))) {
		body["cardData"] = request.CardData
	}

	if !tea.BoolValue(util.IsUnset(request.CardTemplateId)) {
		body["cardTemplateId"] = request.CardTemplateId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CoFeedOpenDeliverModel))) {
		body["coFeedOpenDeliverModel"] = request.CoFeedOpenDeliverModel
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CoFeedOpenSpaceModel))) {
		body["coFeedOpenSpaceModel"] = request.CoFeedOpenSpaceModel
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.ImGroupOpenDeliverModel))) {
		body["imGroupOpenDeliverModel"] = request.ImGroupOpenDeliverModel
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.ImGroupOpenSpaceModel))) {
		body["imGroupOpenSpaceModel"] = request.ImGroupOpenSpaceModel
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.ImSingleOpenDeliverModel))) {
		body["imSingleOpenDeliverModel"] = request.ImSingleOpenDeliverModel
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.ImSingleOpenSpaceModel))) {
		body["imSingleOpenSpaceModel"] = request.ImSingleOpenSpaceModel
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.OpenDynamicDataConfig))) {
		body["openDynamicDataConfig"] = request.OpenDynamicDataConfig
	}

	if !tea.BoolValue(util.IsUnset(request.OpenSpaceId)) {
		body["openSpaceId"] = request.OpenSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.OutTrackId)) {
		body["outTrackId"] = request.OutTrackId
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateData)) {
		body["privateData"] = request.PrivateData
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TopOpenDeliverModel))) {
		body["topOpenDeliverModel"] = request.TopOpenDeliverModel
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TopOpenSpaceModel))) {
		body["topOpenSpaceModel"] = request.TopOpenSpaceModel
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.WorkBenchOpenDeliverModel))) {
		body["workBenchOpenDeliverModel"] = request.WorkBenchOpenDeliverModel
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.WorkBenchOpenSpaceModel))) {
		body["workBenchOpenSpaceModel"] = request.WorkBenchOpenSpaceModel
	}

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
	_result = &CreateAndDeliverResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateAndDeliver"), tea.String("card_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/card/instances/createAndDeliver"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCard(request *CreateCardRequest) (_result *CreateCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateCardHeaders{}
	_result = &CreateCardResponse{}
	_body, _err := client.CreateCardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCardWithOptions(request *CreateCardRequest, headers *CreateCardHeaders, runtime *util.RuntimeOptions) (_result *CreateCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallbackRouteKey)) {
		body["callbackRouteKey"] = request.CallbackRouteKey
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CardData))) {
		body["cardData"] = request.CardData
	}

	if !tea.BoolValue(util.IsUnset(request.CardTemplateId)) {
		body["cardTemplateId"] = request.CardTemplateId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CoFeedOpenSpaceModel))) {
		body["coFeedOpenSpaceModel"] = request.CoFeedOpenSpaceModel
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.ImGroupOpenSpaceModel))) {
		body["imGroupOpenSpaceModel"] = request.ImGroupOpenSpaceModel
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.ImSingleOpenSpaceModel))) {
		body["imSingleOpenSpaceModel"] = request.ImSingleOpenSpaceModel
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.OpenDynamicDataConfig))) {
		body["openDynamicDataConfig"] = request.OpenDynamicDataConfig
	}

	if !tea.BoolValue(util.IsUnset(request.OutTrackId)) {
		body["outTrackId"] = request.OutTrackId
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateData)) {
		body["privateData"] = request.PrivateData
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TopOpenSpaceModel))) {
		body["topOpenSpaceModel"] = request.TopOpenSpaceModel
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.WorkBenchOpenSpaceModel))) {
		body["workBenchOpenSpaceModel"] = request.WorkBenchOpenSpaceModel
	}

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
	_result = &CreateCardResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateCard"), tea.String("card_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/card/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeliverCard(request *DeliverCardRequest) (_result *DeliverCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeliverCardHeaders{}
	_result = &DeliverCardResponse{}
	_body, _err := client.DeliverCardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeliverCardWithOptions(request *DeliverCardRequest, headers *DeliverCardHeaders, runtime *util.RuntimeOptions) (_result *DeliverCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CoFeedOpenDeliverModel))) {
		body["coFeedOpenDeliverModel"] = request.CoFeedOpenDeliverModel
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.ImGroupOpenDeliverModel))) {
		body["imGroupOpenDeliverModel"] = request.ImGroupOpenDeliverModel
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.ImSingleOpenDeliverModel))) {
		body["imSingleOpenDeliverModel"] = request.ImSingleOpenDeliverModel
	}

	if !tea.BoolValue(util.IsUnset(request.OpenSpaceId)) {
		body["openSpaceId"] = request.OpenSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.OutTrackId)) {
		body["outTrackId"] = request.OutTrackId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TopOpenDeliverModel))) {
		body["topOpenDeliverModel"] = request.TopOpenDeliverModel
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.WorkBenchOpenDeliverModel))) {
		body["workBenchOpenDeliverModel"] = request.WorkBenchOpenDeliverModel
	}

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
	_result = &DeliverCardResponse{}
	_body, _err := client.DoROARequest(tea.String("DeliverCard"), tea.String("card_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/card/instances/deliver"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterCallback(request *RegisterCallbackRequest) (_result *RegisterCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RegisterCallbackHeaders{}
	_result = &RegisterCallbackResponse{}
	_body, _err := client.RegisterCallbackWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterCallbackWithOptions(request *RegisterCallbackRequest, headers *RegisterCallbackHeaders, runtime *util.RuntimeOptions) (_result *RegisterCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiSecret)) {
		body["apiSecret"] = request.ApiSecret
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackRouteKey)) {
		body["callbackRouteKey"] = request.CallbackRouteKey
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackUrl)) {
		body["callbackUrl"] = request.CallbackUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ForceUpdate)) {
		body["forceUpdate"] = request.ForceUpdate
	}

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
	_result = &RegisterCallbackResponse{}
	_body, _err := client.DoROARequest(tea.String("RegisterCallback"), tea.String("card_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/card/callbacks/register"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCard(request *UpdateCardRequest) (_result *UpdateCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateCardHeaders{}
	_result = &UpdateCardResponse{}
	_body, _err := client.UpdateCardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCardWithOptions(request *UpdateCardRequest, headers *UpdateCardHeaders, runtime *util.RuntimeOptions) (_result *UpdateCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CardData))) {
		body["cardData"] = request.CardData
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CardUpdateOptions))) {
		body["cardUpdateOptions"] = request.CardUpdateOptions
	}

	if !tea.BoolValue(util.IsUnset(request.OutTrackId)) {
		body["outTrackId"] = request.OutTrackId
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateData)) {
		body["privateData"] = request.PrivateData
	}

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
	_result = &UpdateCardResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateCard"), tea.String("card_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/card/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
