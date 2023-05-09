// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package card_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type PrivateDataValue struct {
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
	CoFeedOpenSpaceModel  *AppendSpaceRequestCoFeedOpenSpaceModel  `json:"coFeedOpenSpaceModel,omitempty" xml:"coFeedOpenSpaceModel,omitempty" type:"Struct"`
	ImGroupOpenSpaceModel *AppendSpaceRequestImGroupOpenSpaceModel `json:"imGroupOpenSpaceModel,omitempty" xml:"imGroupOpenSpaceModel,omitempty" type:"Struct"`
	ImRobotOpenSpaceModel *AppendSpaceRequestImRobotOpenSpaceModel `json:"imRobotOpenSpaceModel,omitempty" xml:"imRobotOpenSpaceModel,omitempty" type:"Struct"`
	OutTrackId            *string                                  `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	TopOpenSpaceModel     *AppendSpaceRequestTopOpenSpaceModel     `json:"topOpenSpaceModel,omitempty" xml:"topOpenSpaceModel,omitempty" type:"Struct"`
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

func (s *AppendSpaceRequest) SetImRobotOpenSpaceModel(v *AppendSpaceRequestImRobotOpenSpaceModel) *AppendSpaceRequest {
	s.ImRobotOpenSpaceModel = v
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

type AppendSpaceRequestCoFeedOpenSpaceModel struct {
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
	LastMessageI18n map[string]*string                                    `json:"lastMessageI18n,omitempty" xml:"lastMessageI18n,omitempty"`
	Notification    *AppendSpaceRequestImGroupOpenSpaceModelNotification  `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	SearchSupport   *AppendSpaceRequestImGroupOpenSpaceModelSearchSupport `json:"searchSupport,omitempty" xml:"searchSupport,omitempty" type:"Struct"`
	SupportForward  *bool                                                 `json:"supportForward,omitempty" xml:"supportForward,omitempty"`
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

type AppendSpaceRequestImRobotOpenSpaceModel struct {
	LastMessageI18n map[string]*string                                    `json:"lastMessageI18n,omitempty" xml:"lastMessageI18n,omitempty"`
	Notification    *AppendSpaceRequestImRobotOpenSpaceModelNotification  `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	SearchSupport   *AppendSpaceRequestImRobotOpenSpaceModelSearchSupport `json:"searchSupport,omitempty" xml:"searchSupport,omitempty" type:"Struct"`
	SupportForward  *bool                                                 `json:"supportForward,omitempty" xml:"supportForward,omitempty"`
}

func (s AppendSpaceRequestImRobotOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceRequestImRobotOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *AppendSpaceRequestImRobotOpenSpaceModel) SetLastMessageI18n(v map[string]*string) *AppendSpaceRequestImRobotOpenSpaceModel {
	s.LastMessageI18n = v
	return s
}

func (s *AppendSpaceRequestImRobotOpenSpaceModel) SetNotification(v *AppendSpaceRequestImRobotOpenSpaceModelNotification) *AppendSpaceRequestImRobotOpenSpaceModel {
	s.Notification = v
	return s
}

func (s *AppendSpaceRequestImRobotOpenSpaceModel) SetSearchSupport(v *AppendSpaceRequestImRobotOpenSpaceModelSearchSupport) *AppendSpaceRequestImRobotOpenSpaceModel {
	s.SearchSupport = v
	return s
}

func (s *AppendSpaceRequestImRobotOpenSpaceModel) SetSupportForward(v bool) *AppendSpaceRequestImRobotOpenSpaceModel {
	s.SupportForward = &v
	return s
}

type AppendSpaceRequestImRobotOpenSpaceModelNotification struct {
	AlertContent    *string `json:"alertContent,omitempty" xml:"alertContent,omitempty"`
	NotificationOff *bool   `json:"notificationOff,omitempty" xml:"notificationOff,omitempty"`
}

func (s AppendSpaceRequestImRobotOpenSpaceModelNotification) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceRequestImRobotOpenSpaceModelNotification) GoString() string {
	return s.String()
}

func (s *AppendSpaceRequestImRobotOpenSpaceModelNotification) SetAlertContent(v string) *AppendSpaceRequestImRobotOpenSpaceModelNotification {
	s.AlertContent = &v
	return s
}

func (s *AppendSpaceRequestImRobotOpenSpaceModelNotification) SetNotificationOff(v bool) *AppendSpaceRequestImRobotOpenSpaceModelNotification {
	s.NotificationOff = &v
	return s
}

type AppendSpaceRequestImRobotOpenSpaceModelSearchSupport struct {
	SearchDesc     *string `json:"searchDesc,omitempty" xml:"searchDesc,omitempty"`
	SearchIcon     *string `json:"searchIcon,omitempty" xml:"searchIcon,omitempty"`
	SearchTypeName *string `json:"searchTypeName,omitempty" xml:"searchTypeName,omitempty"`
}

func (s AppendSpaceRequestImRobotOpenSpaceModelSearchSupport) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceRequestImRobotOpenSpaceModelSearchSupport) GoString() string {
	return s.String()
}

func (s *AppendSpaceRequestImRobotOpenSpaceModelSearchSupport) SetSearchDesc(v string) *AppendSpaceRequestImRobotOpenSpaceModelSearchSupport {
	s.SearchDesc = &v
	return s
}

func (s *AppendSpaceRequestImRobotOpenSpaceModelSearchSupport) SetSearchIcon(v string) *AppendSpaceRequestImRobotOpenSpaceModelSearchSupport {
	s.SearchIcon = &v
	return s
}

func (s *AppendSpaceRequestImRobotOpenSpaceModelSearchSupport) SetSearchTypeName(v string) *AppendSpaceRequestImRobotOpenSpaceModelSearchSupport {
	s.SearchTypeName = &v
	return s
}

type AppendSpaceRequestTopOpenSpaceModel struct {
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AppendSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AppendSpaceResponse) SetStatusCode(v int32) *AppendSpaceResponse {
	s.StatusCode = &v
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
	CallbackRouteKey        *string                                         `json:"callbackRouteKey,omitempty" xml:"callbackRouteKey,omitempty"`
	CallbackType            *string                                         `json:"callbackType,omitempty" xml:"callbackType,omitempty"`
	CardData                *CreateAndDeliverRequestCardData                `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	CardTemplateId          *string                                         `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	CoFeedOpenDeliverModel  *CreateAndDeliverRequestCoFeedOpenDeliverModel  `json:"coFeedOpenDeliverModel,omitempty" xml:"coFeedOpenDeliverModel,omitempty" type:"Struct"`
	CoFeedOpenSpaceModel    *CreateAndDeliverRequestCoFeedOpenSpaceModel    `json:"coFeedOpenSpaceModel,omitempty" xml:"coFeedOpenSpaceModel,omitempty" type:"Struct"`
	DocOpenDeliverModel     *CreateAndDeliverRequestDocOpenDeliverModel     `json:"docOpenDeliverModel,omitempty" xml:"docOpenDeliverModel,omitempty" type:"Struct"`
	ImGroupOpenDeliverModel *CreateAndDeliverRequestImGroupOpenDeliverModel `json:"imGroupOpenDeliverModel,omitempty" xml:"imGroupOpenDeliverModel,omitempty" type:"Struct"`
	ImGroupOpenSpaceModel   *CreateAndDeliverRequestImGroupOpenSpaceModel   `json:"imGroupOpenSpaceModel,omitempty" xml:"imGroupOpenSpaceModel,omitempty" type:"Struct"`
	ImRobotOpenDeliverModel *CreateAndDeliverRequestImRobotOpenDeliverModel `json:"imRobotOpenDeliverModel,omitempty" xml:"imRobotOpenDeliverModel,omitempty" type:"Struct"`
	ImRobotOpenSpaceModel   *CreateAndDeliverRequestImRobotOpenSpaceModel   `json:"imRobotOpenSpaceModel,omitempty" xml:"imRobotOpenSpaceModel,omitempty" type:"Struct"`
	OpenDynamicDataConfig   *CreateAndDeliverRequestOpenDynamicDataConfig   `json:"openDynamicDataConfig,omitempty" xml:"openDynamicDataConfig,omitempty" type:"Struct"`
	OpenSpaceId             *string                                         `json:"openSpaceId,omitempty" xml:"openSpaceId,omitempty"`
	OutTrackId              *string                                         `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	PrivateData             map[string]*PrivateDataValue                    `json:"privateData,omitempty" xml:"privateData,omitempty"`
	TopOpenDeliverModel     *CreateAndDeliverRequestTopOpenDeliverModel     `json:"topOpenDeliverModel,omitempty" xml:"topOpenDeliverModel,omitempty" type:"Struct"`
	TopOpenSpaceModel       *CreateAndDeliverRequestTopOpenSpaceModel       `json:"topOpenSpaceModel,omitempty" xml:"topOpenSpaceModel,omitempty" type:"Struct"`
	UserId                  *string                                         `json:"userId,omitempty" xml:"userId,omitempty"`
	UserIdType              *int32                                          `json:"userIdType,omitempty" xml:"userIdType,omitempty"`
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

func (s *CreateAndDeliverRequest) SetCallbackType(v string) *CreateAndDeliverRequest {
	s.CallbackType = &v
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

func (s *CreateAndDeliverRequest) SetDocOpenDeliverModel(v *CreateAndDeliverRequestDocOpenDeliverModel) *CreateAndDeliverRequest {
	s.DocOpenDeliverModel = v
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

func (s *CreateAndDeliverRequest) SetImRobotOpenDeliverModel(v *CreateAndDeliverRequestImRobotOpenDeliverModel) *CreateAndDeliverRequest {
	s.ImRobotOpenDeliverModel = v
	return s
}

func (s *CreateAndDeliverRequest) SetImRobotOpenSpaceModel(v *CreateAndDeliverRequestImRobotOpenSpaceModel) *CreateAndDeliverRequest {
	s.ImRobotOpenSpaceModel = v
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

func (s *CreateAndDeliverRequest) SetUserIdType(v int32) *CreateAndDeliverRequest {
	s.UserIdType = &v
	return s
}

type CreateAndDeliverRequestCardData struct {
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
	BizTag      *string `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	GmtTimeLine *int64  `json:"gmtTimeLine,omitempty" xml:"gmtTimeLine,omitempty"`
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
	CoolAppCode *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateAndDeliverRequestCoFeedOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequestCoFeedOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequestCoFeedOpenSpaceModel) SetCoolAppCode(v string) *CreateAndDeliverRequestCoFeedOpenSpaceModel {
	s.CoolAppCode = &v
	return s
}

func (s *CreateAndDeliverRequestCoFeedOpenSpaceModel) SetTitle(v string) *CreateAndDeliverRequestCoFeedOpenSpaceModel {
	s.Title = &v
	return s
}

type CreateAndDeliverRequestDocOpenDeliverModel struct {
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateAndDeliverRequestDocOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequestDocOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequestDocOpenDeliverModel) SetUserId(v string) *CreateAndDeliverRequestDocOpenDeliverModel {
	s.UserId = &v
	return s
}

type CreateAndDeliverRequestImGroupOpenDeliverModel struct {
	AtUserIds  map[string]*string `json:"atUserIds,omitempty" xml:"atUserIds,omitempty"`
	Recipients []*string          `json:"recipients,omitempty" xml:"recipients,omitempty" type:"Repeated"`
	RobotCode  *string            `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
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
	LastMessageI18n map[string]*string                                         `json:"lastMessageI18n,omitempty" xml:"lastMessageI18n,omitempty"`
	Notification    *CreateAndDeliverRequestImGroupOpenSpaceModelNotification  `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	SearchSupport   *CreateAndDeliverRequestImGroupOpenSpaceModelSearchSupport `json:"searchSupport,omitempty" xml:"searchSupport,omitempty" type:"Struct"`
	SupportForward  *bool                                                      `json:"supportForward,omitempty" xml:"supportForward,omitempty"`
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

type CreateAndDeliverRequestImRobotOpenDeliverModel struct {
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
}

func (s CreateAndDeliverRequestImRobotOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequestImRobotOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequestImRobotOpenDeliverModel) SetSpaceType(v string) *CreateAndDeliverRequestImRobotOpenDeliverModel {
	s.SpaceType = &v
	return s
}

type CreateAndDeliverRequestImRobotOpenSpaceModel struct {
	LastMessageI18n map[string]*string                                         `json:"lastMessageI18n,omitempty" xml:"lastMessageI18n,omitempty"`
	Notification    *CreateAndDeliverRequestImRobotOpenSpaceModelNotification  `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	SearchSupport   *CreateAndDeliverRequestImRobotOpenSpaceModelSearchSupport `json:"searchSupport,omitempty" xml:"searchSupport,omitempty" type:"Struct"`
	SupportForward  *bool                                                      `json:"supportForward,omitempty" xml:"supportForward,omitempty"`
}

func (s CreateAndDeliverRequestImRobotOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequestImRobotOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequestImRobotOpenSpaceModel) SetLastMessageI18n(v map[string]*string) *CreateAndDeliverRequestImRobotOpenSpaceModel {
	s.LastMessageI18n = v
	return s
}

func (s *CreateAndDeliverRequestImRobotOpenSpaceModel) SetNotification(v *CreateAndDeliverRequestImRobotOpenSpaceModelNotification) *CreateAndDeliverRequestImRobotOpenSpaceModel {
	s.Notification = v
	return s
}

func (s *CreateAndDeliverRequestImRobotOpenSpaceModel) SetSearchSupport(v *CreateAndDeliverRequestImRobotOpenSpaceModelSearchSupport) *CreateAndDeliverRequestImRobotOpenSpaceModel {
	s.SearchSupport = v
	return s
}

func (s *CreateAndDeliverRequestImRobotOpenSpaceModel) SetSupportForward(v bool) *CreateAndDeliverRequestImRobotOpenSpaceModel {
	s.SupportForward = &v
	return s
}

type CreateAndDeliverRequestImRobotOpenSpaceModelNotification struct {
	AlertContent    *string `json:"alertContent,omitempty" xml:"alertContent,omitempty"`
	NotificationOff *bool   `json:"notificationOff,omitempty" xml:"notificationOff,omitempty"`
}

func (s CreateAndDeliverRequestImRobotOpenSpaceModelNotification) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequestImRobotOpenSpaceModelNotification) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequestImRobotOpenSpaceModelNotification) SetAlertContent(v string) *CreateAndDeliverRequestImRobotOpenSpaceModelNotification {
	s.AlertContent = &v
	return s
}

func (s *CreateAndDeliverRequestImRobotOpenSpaceModelNotification) SetNotificationOff(v bool) *CreateAndDeliverRequestImRobotOpenSpaceModelNotification {
	s.NotificationOff = &v
	return s
}

type CreateAndDeliverRequestImRobotOpenSpaceModelSearchSupport struct {
	SearchDesc     *string `json:"searchDesc,omitempty" xml:"searchDesc,omitempty"`
	SearchIcon     *string `json:"searchIcon,omitempty" xml:"searchIcon,omitempty"`
	SearchTypeName *string `json:"searchTypeName,omitempty" xml:"searchTypeName,omitempty"`
}

func (s CreateAndDeliverRequestImRobotOpenSpaceModelSearchSupport) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequestImRobotOpenSpaceModelSearchSupport) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequestImRobotOpenSpaceModelSearchSupport) SetSearchDesc(v string) *CreateAndDeliverRequestImRobotOpenSpaceModelSearchSupport {
	s.SearchDesc = &v
	return s
}

func (s *CreateAndDeliverRequestImRobotOpenSpaceModelSearchSupport) SetSearchIcon(v string) *CreateAndDeliverRequestImRobotOpenSpaceModelSearchSupport {
	s.SearchIcon = &v
	return s
}

func (s *CreateAndDeliverRequestImRobotOpenSpaceModelSearchSupport) SetSearchTypeName(v string) *CreateAndDeliverRequestImRobotOpenSpaceModelSearchSupport {
	s.SearchTypeName = &v
	return s
}

type CreateAndDeliverRequestOpenDynamicDataConfig struct {
	DynamicDataSourceConfigs []*CreateAndDeliverRequestOpenDynamicDataConfigDynamicDataSourceConfigs `json:"dynamicDataSourceConfigs,omitempty" xml:"dynamicDataSourceConfigs,omitempty" type:"Repeated"`
}

func (s CreateAndDeliverRequestOpenDynamicDataConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequestOpenDynamicDataConfig) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequestOpenDynamicDataConfig) SetDynamicDataSourceConfigs(v []*CreateAndDeliverRequestOpenDynamicDataConfigDynamicDataSourceConfigs) *CreateAndDeliverRequestOpenDynamicDataConfig {
	s.DynamicDataSourceConfigs = v
	return s
}

type CreateAndDeliverRequestOpenDynamicDataConfigDynamicDataSourceConfigs struct {
	ConstParams         map[string]*string                                                              `json:"constParams,omitempty" xml:"constParams,omitempty"`
	DynamicDataSourceId *string                                                                         `json:"dynamicDataSourceId,omitempty" xml:"dynamicDataSourceId,omitempty"`
	PullConfig          *CreateAndDeliverRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig `json:"pullConfig,omitempty" xml:"pullConfig,omitempty" type:"Struct"`
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
	Interval     *int32  `json:"interval,omitempty" xml:"interval,omitempty"`
	PullStrategy *string `json:"pullStrategy,omitempty" xml:"pullStrategy,omitempty"`
	TimeUnit     *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
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
	ExpiredTimeMillis *int64    `json:"expiredTimeMillis,omitempty" xml:"expiredTimeMillis,omitempty"`
	Platforms         []*string `json:"platforms,omitempty" xml:"platforms,omitempty" type:"Repeated"`
	UserIds           []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
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
	DeliverResults []*CreateAndDeliverResponseBodyResultDeliverResults `json:"deliverResults,omitempty" xml:"deliverResults,omitempty" type:"Repeated"`
	OutTrackId     *string                                             `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
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
	ErrorMsg  *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	SpaceId   *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAndDeliverResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateAndDeliverResponse) SetStatusCode(v int32) *CreateAndDeliverResponse {
	s.StatusCode = &v
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
	CallbackRouteKey      *string                                 `json:"callbackRouteKey,omitempty" xml:"callbackRouteKey,omitempty"`
	CallbackType          *string                                 `json:"callbackType,omitempty" xml:"callbackType,omitempty"`
	CardData              *CreateCardRequestCardData              `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	CardTemplateId        *string                                 `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	CoFeedOpenSpaceModel  *CreateCardRequestCoFeedOpenSpaceModel  `json:"coFeedOpenSpaceModel,omitempty" xml:"coFeedOpenSpaceModel,omitempty" type:"Struct"`
	ImGroupOpenSpaceModel *CreateCardRequestImGroupOpenSpaceModel `json:"imGroupOpenSpaceModel,omitempty" xml:"imGroupOpenSpaceModel,omitempty" type:"Struct"`
	ImRobotOpenSpaceModel *CreateCardRequestImRobotOpenSpaceModel `json:"imRobotOpenSpaceModel,omitempty" xml:"imRobotOpenSpaceModel,omitempty" type:"Struct"`
	OpenDynamicDataConfig *CreateCardRequestOpenDynamicDataConfig `json:"openDynamicDataConfig,omitempty" xml:"openDynamicDataConfig,omitempty" type:"Struct"`
	OutTrackId            *string                                 `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	PrivateData           map[string]*PrivateDataValue            `json:"privateData,omitempty" xml:"privateData,omitempty"`
	TopOpenSpaceModel     *CreateCardRequestTopOpenSpaceModel     `json:"topOpenSpaceModel,omitempty" xml:"topOpenSpaceModel,omitempty" type:"Struct"`
	UserId                *string                                 `json:"userId,omitempty" xml:"userId,omitempty"`
	UserIdType            *int32                                  `json:"userIdType,omitempty" xml:"userIdType,omitempty"`
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

func (s *CreateCardRequest) SetCallbackType(v string) *CreateCardRequest {
	s.CallbackType = &v
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

func (s *CreateCardRequest) SetImRobotOpenSpaceModel(v *CreateCardRequestImRobotOpenSpaceModel) *CreateCardRequest {
	s.ImRobotOpenSpaceModel = v
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

func (s *CreateCardRequest) SetUserIdType(v int32) *CreateCardRequest {
	s.UserIdType = &v
	return s
}

type CreateCardRequestCardData struct {
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
	LastMessageI18n map[string]*string                                   `json:"lastMessageI18n,omitempty" xml:"lastMessageI18n,omitempty"`
	Notification    *CreateCardRequestImGroupOpenSpaceModelNotification  `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	SearchSupport   *CreateCardRequestImGroupOpenSpaceModelSearchSupport `json:"searchSupport,omitempty" xml:"searchSupport,omitempty" type:"Struct"`
	SupportForward  *bool                                                `json:"supportForward,omitempty" xml:"supportForward,omitempty"`
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
	AlertContent    *string `json:"alertContent,omitempty" xml:"alertContent,omitempty"`
	NotificationOff *bool   `json:"notificationOff,omitempty" xml:"notificationOff,omitempty"`
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
	SearchDesc     *string `json:"searchDesc,omitempty" xml:"searchDesc,omitempty"`
	SearchIcon     *string `json:"searchIcon,omitempty" xml:"searchIcon,omitempty"`
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

type CreateCardRequestImRobotOpenSpaceModel struct {
	LastMessageI18n map[string]*string                                   `json:"lastMessageI18n,omitempty" xml:"lastMessageI18n,omitempty"`
	Notification    *CreateCardRequestImRobotOpenSpaceModelNotification  `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	SearchSupport   *CreateCardRequestImRobotOpenSpaceModelSearchSupport `json:"searchSupport,omitempty" xml:"searchSupport,omitempty" type:"Struct"`
	SupportForward  *bool                                                `json:"supportForward,omitempty" xml:"supportForward,omitempty"`
}

func (s CreateCardRequestImRobotOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s CreateCardRequestImRobotOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *CreateCardRequestImRobotOpenSpaceModel) SetLastMessageI18n(v map[string]*string) *CreateCardRequestImRobotOpenSpaceModel {
	s.LastMessageI18n = v
	return s
}

func (s *CreateCardRequestImRobotOpenSpaceModel) SetNotification(v *CreateCardRequestImRobotOpenSpaceModelNotification) *CreateCardRequestImRobotOpenSpaceModel {
	s.Notification = v
	return s
}

func (s *CreateCardRequestImRobotOpenSpaceModel) SetSearchSupport(v *CreateCardRequestImRobotOpenSpaceModelSearchSupport) *CreateCardRequestImRobotOpenSpaceModel {
	s.SearchSupport = v
	return s
}

func (s *CreateCardRequestImRobotOpenSpaceModel) SetSupportForward(v bool) *CreateCardRequestImRobotOpenSpaceModel {
	s.SupportForward = &v
	return s
}

type CreateCardRequestImRobotOpenSpaceModelNotification struct {
	AlertContent    *string `json:"alertContent,omitempty" xml:"alertContent,omitempty"`
	NotificationOff *bool   `json:"notificationOff,omitempty" xml:"notificationOff,omitempty"`
}

func (s CreateCardRequestImRobotOpenSpaceModelNotification) String() string {
	return tea.Prettify(s)
}

func (s CreateCardRequestImRobotOpenSpaceModelNotification) GoString() string {
	return s.String()
}

func (s *CreateCardRequestImRobotOpenSpaceModelNotification) SetAlertContent(v string) *CreateCardRequestImRobotOpenSpaceModelNotification {
	s.AlertContent = &v
	return s
}

func (s *CreateCardRequestImRobotOpenSpaceModelNotification) SetNotificationOff(v bool) *CreateCardRequestImRobotOpenSpaceModelNotification {
	s.NotificationOff = &v
	return s
}

type CreateCardRequestImRobotOpenSpaceModelSearchSupport struct {
	SearchDesc     *string `json:"searchDesc,omitempty" xml:"searchDesc,omitempty"`
	SearchIcon     *string `json:"searchIcon,omitempty" xml:"searchIcon,omitempty"`
	SearchTypeName *string `json:"searchTypeName,omitempty" xml:"searchTypeName,omitempty"`
}

func (s CreateCardRequestImRobotOpenSpaceModelSearchSupport) String() string {
	return tea.Prettify(s)
}

func (s CreateCardRequestImRobotOpenSpaceModelSearchSupport) GoString() string {
	return s.String()
}

func (s *CreateCardRequestImRobotOpenSpaceModelSearchSupport) SetSearchDesc(v string) *CreateCardRequestImRobotOpenSpaceModelSearchSupport {
	s.SearchDesc = &v
	return s
}

func (s *CreateCardRequestImRobotOpenSpaceModelSearchSupport) SetSearchIcon(v string) *CreateCardRequestImRobotOpenSpaceModelSearchSupport {
	s.SearchIcon = &v
	return s
}

func (s *CreateCardRequestImRobotOpenSpaceModelSearchSupport) SetSearchTypeName(v string) *CreateCardRequestImRobotOpenSpaceModelSearchSupport {
	s.SearchTypeName = &v
	return s
}

type CreateCardRequestOpenDynamicDataConfig struct {
	DynamicDataSourceConfigs []*CreateCardRequestOpenDynamicDataConfigDynamicDataSourceConfigs `json:"dynamicDataSourceConfigs,omitempty" xml:"dynamicDataSourceConfigs,omitempty" type:"Repeated"`
}

func (s CreateCardRequestOpenDynamicDataConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateCardRequestOpenDynamicDataConfig) GoString() string {
	return s.String()
}

func (s *CreateCardRequestOpenDynamicDataConfig) SetDynamicDataSourceConfigs(v []*CreateCardRequestOpenDynamicDataConfigDynamicDataSourceConfigs) *CreateCardRequestOpenDynamicDataConfig {
	s.DynamicDataSourceConfigs = v
	return s
}

type CreateCardRequestOpenDynamicDataConfigDynamicDataSourceConfigs struct {
	ConstParams         map[string]*string                                                        `json:"constParams,omitempty" xml:"constParams,omitempty"`
	DynamicDataSourceId *string                                                                   `json:"dynamicDataSourceId,omitempty" xml:"dynamicDataSourceId,omitempty"`
	PullConfig          *CreateCardRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig `json:"pullConfig,omitempty" xml:"pullConfig,omitempty" type:"Struct"`
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
	Interval     *int32  `json:"interval,omitempty" xml:"interval,omitempty"`
	PullStrategy *string `json:"pullStrategy,omitempty" xml:"pullStrategy,omitempty"`
	TimeUnit     *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateCardResponse) SetStatusCode(v int32) *CreateCardResponse {
	s.StatusCode = &v
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
	CoFeedOpenDeliverModel  *DeliverCardRequestCoFeedOpenDeliverModel  `json:"coFeedOpenDeliverModel,omitempty" xml:"coFeedOpenDeliverModel,omitempty" type:"Struct"`
	DocOpenDeliverModel     *DeliverCardRequestDocOpenDeliverModel     `json:"docOpenDeliverModel,omitempty" xml:"docOpenDeliverModel,omitempty" type:"Struct"`
	ImGroupOpenDeliverModel *DeliverCardRequestImGroupOpenDeliverModel `json:"imGroupOpenDeliverModel,omitempty" xml:"imGroupOpenDeliverModel,omitempty" type:"Struct"`
	ImRobotOpenDeliverModel *DeliverCardRequestImRobotOpenDeliverModel `json:"imRobotOpenDeliverModel,omitempty" xml:"imRobotOpenDeliverModel,omitempty" type:"Struct"`
	OpenSpaceId             *string                                    `json:"openSpaceId,omitempty" xml:"openSpaceId,omitempty"`
	OutTrackId              *string                                    `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	TopOpenDeliverModel     *DeliverCardRequestTopOpenDeliverModel     `json:"topOpenDeliverModel,omitempty" xml:"topOpenDeliverModel,omitempty" type:"Struct"`
	UserIdType              *int32                                     `json:"userIdType,omitempty" xml:"userIdType,omitempty"`
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

func (s *DeliverCardRequest) SetDocOpenDeliverModel(v *DeliverCardRequestDocOpenDeliverModel) *DeliverCardRequest {
	s.DocOpenDeliverModel = v
	return s
}

func (s *DeliverCardRequest) SetImGroupOpenDeliverModel(v *DeliverCardRequestImGroupOpenDeliverModel) *DeliverCardRequest {
	s.ImGroupOpenDeliverModel = v
	return s
}

func (s *DeliverCardRequest) SetImRobotOpenDeliverModel(v *DeliverCardRequestImRobotOpenDeliverModel) *DeliverCardRequest {
	s.ImRobotOpenDeliverModel = v
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

func (s *DeliverCardRequest) SetUserIdType(v int32) *DeliverCardRequest {
	s.UserIdType = &v
	return s
}

type DeliverCardRequestCoFeedOpenDeliverModel struct {
	BizTag      *string `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	GmtTimeLine *int64  `json:"gmtTimeLine,omitempty" xml:"gmtTimeLine,omitempty"`
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

type DeliverCardRequestDocOpenDeliverModel struct {
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeliverCardRequestDocOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s DeliverCardRequestDocOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *DeliverCardRequestDocOpenDeliverModel) SetUserId(v string) *DeliverCardRequestDocOpenDeliverModel {
	s.UserId = &v
	return s
}

type DeliverCardRequestImGroupOpenDeliverModel struct {
	AtUserIds  map[string]*string `json:"atUserIds,omitempty" xml:"atUserIds,omitempty"`
	Recipients []*string          `json:"recipients,omitempty" xml:"recipients,omitempty" type:"Repeated"`
	RobotCode  *string            `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
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

type DeliverCardRequestImRobotOpenDeliverModel struct {
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
}

func (s DeliverCardRequestImRobotOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s DeliverCardRequestImRobotOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *DeliverCardRequestImRobotOpenDeliverModel) SetSpaceType(v string) *DeliverCardRequestImRobotOpenDeliverModel {
	s.SpaceType = &v
	return s
}

type DeliverCardRequestTopOpenDeliverModel struct {
	ExpiredTimeMillis *int64    `json:"expiredTimeMillis,omitempty" xml:"expiredTimeMillis,omitempty"`
	Platforms         []*string `json:"platforms,omitempty" xml:"platforms,omitempty" type:"Repeated"`
	UserIds           []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s DeliverCardRequestTopOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s DeliverCardRequestTopOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *DeliverCardRequestTopOpenDeliverModel) SetExpiredTimeMillis(v int64) *DeliverCardRequestTopOpenDeliverModel {
	s.ExpiredTimeMillis = &v
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
	SpaceId   *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeliverCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeliverCardResponse) SetStatusCode(v int32) *DeliverCardResponse {
	s.StatusCode = &v
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
	ApiSecret        *string `json:"apiSecret,omitempty" xml:"apiSecret,omitempty"`
	CallbackRouteKey *string `json:"callbackRouteKey,omitempty" xml:"callbackRouteKey,omitempty"`
	CallbackUrl      *string `json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`
	ForceUpdate      *bool   `json:"forceUpdate,omitempty" xml:"forceUpdate,omitempty"`
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
	ApiSecret   *string `json:"apiSecret,omitempty" xml:"apiSecret,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RegisterCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RegisterCallbackResponse) SetStatusCode(v int32) *RegisterCallbackResponse {
	s.StatusCode = &v
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
	CardData          *UpdateCardRequestCardData          `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	CardUpdateOptions *UpdateCardRequestCardUpdateOptions `json:"cardUpdateOptions,omitempty" xml:"cardUpdateOptions,omitempty" type:"Struct"`
	OutTrackId        *string                             `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	PrivateData       map[string]*PrivateDataValue        `json:"privateData,omitempty" xml:"privateData,omitempty"`
	UserIdType        *int32                              `json:"userIdType,omitempty" xml:"userIdType,omitempty"`
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

func (s *UpdateCardRequest) SetUserIdType(v int32) *UpdateCardRequest {
	s.UserIdType = &v
	return s
}

type UpdateCardRequestCardData struct {
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
	UpdateCardDataByKey    *bool `json:"updateCardDataByKey,omitempty" xml:"updateCardDataByKey,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateCardResponse) SetStatusCode(v int32) *UpdateCardResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCardResponse) SetBody(v *UpdateCardResponseBody) *UpdateCardResponse {
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

func (client *Client) AppendSpaceWithOptions(request *AppendSpaceRequest, headers *AppendSpaceHeaders, runtime *util.RuntimeOptions) (_result *AppendSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CoFeedOpenSpaceModel)) {
		body["coFeedOpenSpaceModel"] = request.CoFeedOpenSpaceModel
	}

	if !tea.BoolValue(util.IsUnset(request.ImGroupOpenSpaceModel)) {
		body["imGroupOpenSpaceModel"] = request.ImGroupOpenSpaceModel
	}

	if !tea.BoolValue(util.IsUnset(request.ImRobotOpenSpaceModel)) {
		body["imRobotOpenSpaceModel"] = request.ImRobotOpenSpaceModel
	}

	if !tea.BoolValue(util.IsUnset(request.OutTrackId)) {
		body["outTrackId"] = request.OutTrackId
	}

	if !tea.BoolValue(util.IsUnset(request.TopOpenSpaceModel)) {
		body["topOpenSpaceModel"] = request.TopOpenSpaceModel
	}

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
		Action:      tea.String("AppendSpace"),
		Version:     tea.String("card_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/card/instances/spaces"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AppendSpaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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

func (client *Client) CreateAndDeliverWithOptions(request *CreateAndDeliverRequest, headers *CreateAndDeliverHeaders, runtime *util.RuntimeOptions) (_result *CreateAndDeliverResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallbackRouteKey)) {
		body["callbackRouteKey"] = request.CallbackRouteKey
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackType)) {
		body["callbackType"] = request.CallbackType
	}

	if !tea.BoolValue(util.IsUnset(request.CardData)) {
		body["cardData"] = request.CardData
	}

	if !tea.BoolValue(util.IsUnset(request.CardTemplateId)) {
		body["cardTemplateId"] = request.CardTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.CoFeedOpenDeliverModel)) {
		body["coFeedOpenDeliverModel"] = request.CoFeedOpenDeliverModel
	}

	if !tea.BoolValue(util.IsUnset(request.CoFeedOpenSpaceModel)) {
		body["coFeedOpenSpaceModel"] = request.CoFeedOpenSpaceModel
	}

	if !tea.BoolValue(util.IsUnset(request.DocOpenDeliverModel)) {
		body["docOpenDeliverModel"] = request.DocOpenDeliverModel
	}

	if !tea.BoolValue(util.IsUnset(request.ImGroupOpenDeliverModel)) {
		body["imGroupOpenDeliverModel"] = request.ImGroupOpenDeliverModel
	}

	if !tea.BoolValue(util.IsUnset(request.ImGroupOpenSpaceModel)) {
		body["imGroupOpenSpaceModel"] = request.ImGroupOpenSpaceModel
	}

	if !tea.BoolValue(util.IsUnset(request.ImRobotOpenDeliverModel)) {
		body["imRobotOpenDeliverModel"] = request.ImRobotOpenDeliverModel
	}

	if !tea.BoolValue(util.IsUnset(request.ImRobotOpenSpaceModel)) {
		body["imRobotOpenSpaceModel"] = request.ImRobotOpenSpaceModel
	}

	if !tea.BoolValue(util.IsUnset(request.OpenDynamicDataConfig)) {
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

	if !tea.BoolValue(util.IsUnset(request.TopOpenDeliverModel)) {
		body["topOpenDeliverModel"] = request.TopOpenDeliverModel
	}

	if !tea.BoolValue(util.IsUnset(request.TopOpenSpaceModel)) {
		body["topOpenSpaceModel"] = request.TopOpenSpaceModel
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdType)) {
		body["userIdType"] = request.UserIdType
	}

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
		Action:      tea.String("CreateAndDeliver"),
		Version:     tea.String("card_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/card/instances/createAndDeliver"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAndDeliverResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) CreateCardWithOptions(request *CreateCardRequest, headers *CreateCardHeaders, runtime *util.RuntimeOptions) (_result *CreateCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallbackRouteKey)) {
		body["callbackRouteKey"] = request.CallbackRouteKey
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackType)) {
		body["callbackType"] = request.CallbackType
	}

	if !tea.BoolValue(util.IsUnset(request.CardData)) {
		body["cardData"] = request.CardData
	}

	if !tea.BoolValue(util.IsUnset(request.CardTemplateId)) {
		body["cardTemplateId"] = request.CardTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.CoFeedOpenSpaceModel)) {
		body["coFeedOpenSpaceModel"] = request.CoFeedOpenSpaceModel
	}

	if !tea.BoolValue(util.IsUnset(request.ImGroupOpenSpaceModel)) {
		body["imGroupOpenSpaceModel"] = request.ImGroupOpenSpaceModel
	}

	if !tea.BoolValue(util.IsUnset(request.ImRobotOpenSpaceModel)) {
		body["imRobotOpenSpaceModel"] = request.ImRobotOpenSpaceModel
	}

	if !tea.BoolValue(util.IsUnset(request.OpenDynamicDataConfig)) {
		body["openDynamicDataConfig"] = request.OpenDynamicDataConfig
	}

	if !tea.BoolValue(util.IsUnset(request.OutTrackId)) {
		body["outTrackId"] = request.OutTrackId
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateData)) {
		body["privateData"] = request.PrivateData
	}

	if !tea.BoolValue(util.IsUnset(request.TopOpenSpaceModel)) {
		body["topOpenSpaceModel"] = request.TopOpenSpaceModel
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdType)) {
		body["userIdType"] = request.UserIdType
	}

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
		Action:      tea.String("CreateCard"),
		Version:     tea.String("card_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/card/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCardResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) DeliverCardWithOptions(request *DeliverCardRequest, headers *DeliverCardHeaders, runtime *util.RuntimeOptions) (_result *DeliverCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CoFeedOpenDeliverModel)) {
		body["coFeedOpenDeliverModel"] = request.CoFeedOpenDeliverModel
	}

	if !tea.BoolValue(util.IsUnset(request.DocOpenDeliverModel)) {
		body["docOpenDeliverModel"] = request.DocOpenDeliverModel
	}

	if !tea.BoolValue(util.IsUnset(request.ImGroupOpenDeliverModel)) {
		body["imGroupOpenDeliverModel"] = request.ImGroupOpenDeliverModel
	}

	if !tea.BoolValue(util.IsUnset(request.ImRobotOpenDeliverModel)) {
		body["imRobotOpenDeliverModel"] = request.ImRobotOpenDeliverModel
	}

	if !tea.BoolValue(util.IsUnset(request.OpenSpaceId)) {
		body["openSpaceId"] = request.OpenSpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.OutTrackId)) {
		body["outTrackId"] = request.OutTrackId
	}

	if !tea.BoolValue(util.IsUnset(request.TopOpenDeliverModel)) {
		body["topOpenDeliverModel"] = request.TopOpenDeliverModel
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdType)) {
		body["userIdType"] = request.UserIdType
	}

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
		Action:      tea.String("DeliverCard"),
		Version:     tea.String("card_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/card/instances/deliver"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeliverCardResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("RegisterCallback"),
		Version:     tea.String("card_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/card/callbacks/register"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RegisterCallbackResponse{}
	_body, _err := client.Execute(params, req, runtime)
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

func (client *Client) UpdateCardWithOptions(request *UpdateCardRequest, headers *UpdateCardHeaders, runtime *util.RuntimeOptions) (_result *UpdateCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CardData)) {
		body["cardData"] = request.CardData
	}

	if !tea.BoolValue(util.IsUnset(request.CardUpdateOptions)) {
		body["cardUpdateOptions"] = request.CardUpdateOptions
	}

	if !tea.BoolValue(util.IsUnset(request.OutTrackId)) {
		body["outTrackId"] = request.OutTrackId
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateData)) {
		body["privateData"] = request.PrivateData
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdType)) {
		body["userIdType"] = request.UserIdType
	}

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
		Action:      tea.String("UpdateCard"),
		Version:     tea.String("card_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/card/instances"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCardResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
