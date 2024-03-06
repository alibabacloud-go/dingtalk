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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AppendSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type AppendSpaceWithDelegateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AppendSpaceWithDelegateHeaders) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceWithDelegateHeaders) GoString() string {
	return s.String()
}

func (s *AppendSpaceWithDelegateHeaders) SetCommonHeaders(v map[string]*string) *AppendSpaceWithDelegateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AppendSpaceWithDelegateHeaders) SetXAcsDingtalkAccessToken(v string) *AppendSpaceWithDelegateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AppendSpaceWithDelegateRequest struct {
	CoFeedOpenSpaceModel  *AppendSpaceWithDelegateRequestCoFeedOpenSpaceModel  `json:"coFeedOpenSpaceModel,omitempty" xml:"coFeedOpenSpaceModel,omitempty" type:"Struct"`
	ImGroupOpenSpaceModel *AppendSpaceWithDelegateRequestImGroupOpenSpaceModel `json:"imGroupOpenSpaceModel,omitempty" xml:"imGroupOpenSpaceModel,omitempty" type:"Struct"`
	ImRobotOpenSpaceModel *AppendSpaceWithDelegateRequestImRobotOpenSpaceModel `json:"imRobotOpenSpaceModel,omitempty" xml:"imRobotOpenSpaceModel,omitempty" type:"Struct"`
	OutTrackId            *string                                              `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	TopOpenSpaceModel     *AppendSpaceWithDelegateRequestTopOpenSpaceModel     `json:"topOpenSpaceModel,omitempty" xml:"topOpenSpaceModel,omitempty" type:"Struct"`
}

func (s AppendSpaceWithDelegateRequest) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceWithDelegateRequest) GoString() string {
	return s.String()
}

func (s *AppendSpaceWithDelegateRequest) SetCoFeedOpenSpaceModel(v *AppendSpaceWithDelegateRequestCoFeedOpenSpaceModel) *AppendSpaceWithDelegateRequest {
	s.CoFeedOpenSpaceModel = v
	return s
}

func (s *AppendSpaceWithDelegateRequest) SetImGroupOpenSpaceModel(v *AppendSpaceWithDelegateRequestImGroupOpenSpaceModel) *AppendSpaceWithDelegateRequest {
	s.ImGroupOpenSpaceModel = v
	return s
}

func (s *AppendSpaceWithDelegateRequest) SetImRobotOpenSpaceModel(v *AppendSpaceWithDelegateRequestImRobotOpenSpaceModel) *AppendSpaceWithDelegateRequest {
	s.ImRobotOpenSpaceModel = v
	return s
}

func (s *AppendSpaceWithDelegateRequest) SetOutTrackId(v string) *AppendSpaceWithDelegateRequest {
	s.OutTrackId = &v
	return s
}

func (s *AppendSpaceWithDelegateRequest) SetTopOpenSpaceModel(v *AppendSpaceWithDelegateRequestTopOpenSpaceModel) *AppendSpaceWithDelegateRequest {
	s.TopOpenSpaceModel = v
	return s
}

type AppendSpaceWithDelegateRequestCoFeedOpenSpaceModel struct {
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s AppendSpaceWithDelegateRequestCoFeedOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceWithDelegateRequestCoFeedOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *AppendSpaceWithDelegateRequestCoFeedOpenSpaceModel) SetTitle(v string) *AppendSpaceWithDelegateRequestCoFeedOpenSpaceModel {
	s.Title = &v
	return s
}

type AppendSpaceWithDelegateRequestImGroupOpenSpaceModel struct {
	LastMessageI18n map[string]*string                                                `json:"lastMessageI18n,omitempty" xml:"lastMessageI18n,omitempty"`
	Notification    *AppendSpaceWithDelegateRequestImGroupOpenSpaceModelNotification  `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	SearchSupport   *AppendSpaceWithDelegateRequestImGroupOpenSpaceModelSearchSupport `json:"searchSupport,omitempty" xml:"searchSupport,omitempty" type:"Struct"`
	SupportForward  *bool                                                             `json:"supportForward,omitempty" xml:"supportForward,omitempty"`
}

func (s AppendSpaceWithDelegateRequestImGroupOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceWithDelegateRequestImGroupOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *AppendSpaceWithDelegateRequestImGroupOpenSpaceModel) SetLastMessageI18n(v map[string]*string) *AppendSpaceWithDelegateRequestImGroupOpenSpaceModel {
	s.LastMessageI18n = v
	return s
}

func (s *AppendSpaceWithDelegateRequestImGroupOpenSpaceModel) SetNotification(v *AppendSpaceWithDelegateRequestImGroupOpenSpaceModelNotification) *AppendSpaceWithDelegateRequestImGroupOpenSpaceModel {
	s.Notification = v
	return s
}

func (s *AppendSpaceWithDelegateRequestImGroupOpenSpaceModel) SetSearchSupport(v *AppendSpaceWithDelegateRequestImGroupOpenSpaceModelSearchSupport) *AppendSpaceWithDelegateRequestImGroupOpenSpaceModel {
	s.SearchSupport = v
	return s
}

func (s *AppendSpaceWithDelegateRequestImGroupOpenSpaceModel) SetSupportForward(v bool) *AppendSpaceWithDelegateRequestImGroupOpenSpaceModel {
	s.SupportForward = &v
	return s
}

type AppendSpaceWithDelegateRequestImGroupOpenSpaceModelNotification struct {
	AlertContent    *string `json:"alertContent,omitempty" xml:"alertContent,omitempty"`
	NotificationOff *bool   `json:"notificationOff,omitempty" xml:"notificationOff,omitempty"`
}

func (s AppendSpaceWithDelegateRequestImGroupOpenSpaceModelNotification) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceWithDelegateRequestImGroupOpenSpaceModelNotification) GoString() string {
	return s.String()
}

func (s *AppendSpaceWithDelegateRequestImGroupOpenSpaceModelNotification) SetAlertContent(v string) *AppendSpaceWithDelegateRequestImGroupOpenSpaceModelNotification {
	s.AlertContent = &v
	return s
}

func (s *AppendSpaceWithDelegateRequestImGroupOpenSpaceModelNotification) SetNotificationOff(v bool) *AppendSpaceWithDelegateRequestImGroupOpenSpaceModelNotification {
	s.NotificationOff = &v
	return s
}

type AppendSpaceWithDelegateRequestImGroupOpenSpaceModelSearchSupport struct {
	SearchDesc     *string `json:"searchDesc,omitempty" xml:"searchDesc,omitempty"`
	SearchIcon     *string `json:"searchIcon,omitempty" xml:"searchIcon,omitempty"`
	SearchTypeName *string `json:"searchTypeName,omitempty" xml:"searchTypeName,omitempty"`
}

func (s AppendSpaceWithDelegateRequestImGroupOpenSpaceModelSearchSupport) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceWithDelegateRequestImGroupOpenSpaceModelSearchSupport) GoString() string {
	return s.String()
}

func (s *AppendSpaceWithDelegateRequestImGroupOpenSpaceModelSearchSupport) SetSearchDesc(v string) *AppendSpaceWithDelegateRequestImGroupOpenSpaceModelSearchSupport {
	s.SearchDesc = &v
	return s
}

func (s *AppendSpaceWithDelegateRequestImGroupOpenSpaceModelSearchSupport) SetSearchIcon(v string) *AppendSpaceWithDelegateRequestImGroupOpenSpaceModelSearchSupport {
	s.SearchIcon = &v
	return s
}

func (s *AppendSpaceWithDelegateRequestImGroupOpenSpaceModelSearchSupport) SetSearchTypeName(v string) *AppendSpaceWithDelegateRequestImGroupOpenSpaceModelSearchSupport {
	s.SearchTypeName = &v
	return s
}

type AppendSpaceWithDelegateRequestImRobotOpenSpaceModel struct {
	LastMessageI18n map[string]*string                                                `json:"lastMessageI18n,omitempty" xml:"lastMessageI18n,omitempty"`
	Notification    *AppendSpaceWithDelegateRequestImRobotOpenSpaceModelNotification  `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	SearchSupport   *AppendSpaceWithDelegateRequestImRobotOpenSpaceModelSearchSupport `json:"searchSupport,omitempty" xml:"searchSupport,omitempty" type:"Struct"`
	SupportForward  *bool                                                             `json:"supportForward,omitempty" xml:"supportForward,omitempty"`
}

func (s AppendSpaceWithDelegateRequestImRobotOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceWithDelegateRequestImRobotOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *AppendSpaceWithDelegateRequestImRobotOpenSpaceModel) SetLastMessageI18n(v map[string]*string) *AppendSpaceWithDelegateRequestImRobotOpenSpaceModel {
	s.LastMessageI18n = v
	return s
}

func (s *AppendSpaceWithDelegateRequestImRobotOpenSpaceModel) SetNotification(v *AppendSpaceWithDelegateRequestImRobotOpenSpaceModelNotification) *AppendSpaceWithDelegateRequestImRobotOpenSpaceModel {
	s.Notification = v
	return s
}

func (s *AppendSpaceWithDelegateRequestImRobotOpenSpaceModel) SetSearchSupport(v *AppendSpaceWithDelegateRequestImRobotOpenSpaceModelSearchSupport) *AppendSpaceWithDelegateRequestImRobotOpenSpaceModel {
	s.SearchSupport = v
	return s
}

func (s *AppendSpaceWithDelegateRequestImRobotOpenSpaceModel) SetSupportForward(v bool) *AppendSpaceWithDelegateRequestImRobotOpenSpaceModel {
	s.SupportForward = &v
	return s
}

type AppendSpaceWithDelegateRequestImRobotOpenSpaceModelNotification struct {
	AlertContent    *string `json:"alertContent,omitempty" xml:"alertContent,omitempty"`
	NotificationOff *bool   `json:"notificationOff,omitempty" xml:"notificationOff,omitempty"`
}

func (s AppendSpaceWithDelegateRequestImRobotOpenSpaceModelNotification) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceWithDelegateRequestImRobotOpenSpaceModelNotification) GoString() string {
	return s.String()
}

func (s *AppendSpaceWithDelegateRequestImRobotOpenSpaceModelNotification) SetAlertContent(v string) *AppendSpaceWithDelegateRequestImRobotOpenSpaceModelNotification {
	s.AlertContent = &v
	return s
}

func (s *AppendSpaceWithDelegateRequestImRobotOpenSpaceModelNotification) SetNotificationOff(v bool) *AppendSpaceWithDelegateRequestImRobotOpenSpaceModelNotification {
	s.NotificationOff = &v
	return s
}

type AppendSpaceWithDelegateRequestImRobotOpenSpaceModelSearchSupport struct {
	SearchDesc     *string `json:"searchDesc,omitempty" xml:"searchDesc,omitempty"`
	SearchIcon     *string `json:"searchIcon,omitempty" xml:"searchIcon,omitempty"`
	SearchTypeName *string `json:"searchTypeName,omitempty" xml:"searchTypeName,omitempty"`
}

func (s AppendSpaceWithDelegateRequestImRobotOpenSpaceModelSearchSupport) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceWithDelegateRequestImRobotOpenSpaceModelSearchSupport) GoString() string {
	return s.String()
}

func (s *AppendSpaceWithDelegateRequestImRobotOpenSpaceModelSearchSupport) SetSearchDesc(v string) *AppendSpaceWithDelegateRequestImRobotOpenSpaceModelSearchSupport {
	s.SearchDesc = &v
	return s
}

func (s *AppendSpaceWithDelegateRequestImRobotOpenSpaceModelSearchSupport) SetSearchIcon(v string) *AppendSpaceWithDelegateRequestImRobotOpenSpaceModelSearchSupport {
	s.SearchIcon = &v
	return s
}

func (s *AppendSpaceWithDelegateRequestImRobotOpenSpaceModelSearchSupport) SetSearchTypeName(v string) *AppendSpaceWithDelegateRequestImRobotOpenSpaceModelSearchSupport {
	s.SearchTypeName = &v
	return s
}

type AppendSpaceWithDelegateRequestTopOpenSpaceModel struct {
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
}

func (s AppendSpaceWithDelegateRequestTopOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceWithDelegateRequestTopOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *AppendSpaceWithDelegateRequestTopOpenSpaceModel) SetSpaceType(v string) *AppendSpaceWithDelegateRequestTopOpenSpaceModel {
	s.SpaceType = &v
	return s
}

type AppendSpaceWithDelegateResponseBody struct {
	Result  *bool `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AppendSpaceWithDelegateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceWithDelegateResponseBody) GoString() string {
	return s.String()
}

func (s *AppendSpaceWithDelegateResponseBody) SetResult(v bool) *AppendSpaceWithDelegateResponseBody {
	s.Result = &v
	return s
}

func (s *AppendSpaceWithDelegateResponseBody) SetSuccess(v bool) *AppendSpaceWithDelegateResponseBody {
	s.Success = &v
	return s
}

type AppendSpaceWithDelegateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AppendSpaceWithDelegateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AppendSpaceWithDelegateResponse) String() string {
	return tea.Prettify(s)
}

func (s AppendSpaceWithDelegateResponse) GoString() string {
	return s.String()
}

func (s *AppendSpaceWithDelegateResponse) SetHeaders(v map[string]*string) *AppendSpaceWithDelegateResponse {
	s.Headers = v
	return s
}

func (s *AppendSpaceWithDelegateResponse) SetStatusCode(v int32) *AppendSpaceWithDelegateResponse {
	s.StatusCode = &v
	return s
}

func (s *AppendSpaceWithDelegateResponse) SetBody(v *AppendSpaceWithDelegateResponseBody) *AppendSpaceWithDelegateResponse {
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
	CallbackRouteKey         *string                                          `json:"callbackRouteKey,omitempty" xml:"callbackRouteKey,omitempty"`
	CallbackType             *string                                          `json:"callbackType,omitempty" xml:"callbackType,omitempty"`
	CardData                 *CreateAndDeliverRequestCardData                 `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	CardTemplateId           *string                                          `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	CoFeedOpenDeliverModel   *CreateAndDeliverRequestCoFeedOpenDeliverModel   `json:"coFeedOpenDeliverModel,omitempty" xml:"coFeedOpenDeliverModel,omitempty" type:"Struct"`
	CoFeedOpenSpaceModel     *CreateAndDeliverRequestCoFeedOpenSpaceModel     `json:"coFeedOpenSpaceModel,omitempty" xml:"coFeedOpenSpaceModel,omitempty" type:"Struct"`
	DocOpenDeliverModel      *CreateAndDeliverRequestDocOpenDeliverModel      `json:"docOpenDeliverModel,omitempty" xml:"docOpenDeliverModel,omitempty" type:"Struct"`
	ImGroupOpenDeliverModel  *CreateAndDeliverRequestImGroupOpenDeliverModel  `json:"imGroupOpenDeliverModel,omitempty" xml:"imGroupOpenDeliverModel,omitempty" type:"Struct"`
	ImGroupOpenSpaceModel    *CreateAndDeliverRequestImGroupOpenSpaceModel    `json:"imGroupOpenSpaceModel,omitempty" xml:"imGroupOpenSpaceModel,omitempty" type:"Struct"`
	ImRobotOpenDeliverModel  *CreateAndDeliverRequestImRobotOpenDeliverModel  `json:"imRobotOpenDeliverModel,omitempty" xml:"imRobotOpenDeliverModel,omitempty" type:"Struct"`
	ImRobotOpenSpaceModel    *CreateAndDeliverRequestImRobotOpenSpaceModel    `json:"imRobotOpenSpaceModel,omitempty" xml:"imRobotOpenSpaceModel,omitempty" type:"Struct"`
	ImSingleOpenDeliverModel *CreateAndDeliverRequestImSingleOpenDeliverModel `json:"imSingleOpenDeliverModel,omitempty" xml:"imSingleOpenDeliverModel,omitempty" type:"Struct"`
	ImSingleOpenSpaceModel   *CreateAndDeliverRequestImSingleOpenSpaceModel   `json:"imSingleOpenSpaceModel,omitempty" xml:"imSingleOpenSpaceModel,omitempty" type:"Struct"`
	OpenDynamicDataConfig    *CreateAndDeliverRequestOpenDynamicDataConfig    `json:"openDynamicDataConfig,omitempty" xml:"openDynamicDataConfig,omitempty" type:"Struct"`
	OpenSpaceId              *string                                          `json:"openSpaceId,omitempty" xml:"openSpaceId,omitempty"`
	OutTrackId               *string                                          `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	PrivateData              map[string]*PrivateDataValue                     `json:"privateData,omitempty" xml:"privateData,omitempty"`
	TopOpenDeliverModel      *CreateAndDeliverRequestTopOpenDeliverModel      `json:"topOpenDeliverModel,omitempty" xml:"topOpenDeliverModel,omitempty" type:"Struct"`
	TopOpenSpaceModel        *CreateAndDeliverRequestTopOpenSpaceModel        `json:"topOpenSpaceModel,omitempty" xml:"topOpenSpaceModel,omitempty" type:"Struct"`
	UserId                   *string                                          `json:"userId,omitempty" xml:"userId,omitempty"`
	UserIdType               *int32                                           `json:"userIdType,omitempty" xml:"userIdType,omitempty"`
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
	Extension  map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
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

func (s *CreateAndDeliverRequestImGroupOpenDeliverModel) SetExtension(v map[string]*string) *CreateAndDeliverRequestImGroupOpenDeliverModel {
	s.Extension = v
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
	Extension map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
	RobotCode *string            `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	SpaceType *string            `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
}

func (s CreateAndDeliverRequestImRobotOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverRequestImRobotOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverRequestImRobotOpenDeliverModel) SetExtension(v map[string]*string) *CreateAndDeliverRequestImRobotOpenDeliverModel {
	s.Extension = v
	return s
}

func (s *CreateAndDeliverRequestImRobotOpenDeliverModel) SetRobotCode(v string) *CreateAndDeliverRequestImRobotOpenDeliverModel {
	s.RobotCode = &v
	return s
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

type CreateAndDeliverRequestImSingleOpenDeliverModel struct {
	AtUserIds map[string]*string `json:"atUserIds,omitempty" xml:"atUserIds,omitempty"`
	Extension map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
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

func (s *CreateAndDeliverRequestImSingleOpenDeliverModel) SetExtension(v map[string]*string) *CreateAndDeliverRequestImSingleOpenDeliverModel {
	s.Extension = v
	return s
}

type CreateAndDeliverRequestImSingleOpenSpaceModel struct {
	LastMessageI18n map[string]*string                                          `json:"lastMessageI18n,omitempty" xml:"lastMessageI18n,omitempty"`
	Notification    *CreateAndDeliverRequestImSingleOpenSpaceModelNotification  `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	SearchSupport   *CreateAndDeliverRequestImSingleOpenSpaceModelSearchSupport `json:"searchSupport,omitempty" xml:"searchSupport,omitempty" type:"Struct"`
	SupportForward  *bool                                                       `json:"supportForward,omitempty" xml:"supportForward,omitempty"`
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
	CarrierId *string `json:"carrierId,omitempty" xml:"carrierId,omitempty"`
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

func (s *CreateAndDeliverResponseBodyResultDeliverResults) SetCarrierId(v string) *CreateAndDeliverResponseBodyResultDeliverResults {
	s.CarrierId = &v
	return s
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAndDeliverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CreateAndDeliverWithDelegateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateAndDeliverWithDelegateHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateHeaders) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateHeaders) SetCommonHeaders(v map[string]*string) *CreateAndDeliverWithDelegateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateAndDeliverWithDelegateHeaders) SetXAcsDingtalkAccessToken(v string) *CreateAndDeliverWithDelegateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateAndDeliverWithDelegateRequest struct {
	CallbackRouteKey         *string                                                      `json:"callbackRouteKey,omitempty" xml:"callbackRouteKey,omitempty"`
	CallbackType             *string                                                      `json:"callbackType,omitempty" xml:"callbackType,omitempty"`
	CardData                 *CreateAndDeliverWithDelegateRequestCardData                 `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	CardTemplateId           *string                                                      `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	CoFeedOpenDeliverModel   *CreateAndDeliverWithDelegateRequestCoFeedOpenDeliverModel   `json:"coFeedOpenDeliverModel,omitempty" xml:"coFeedOpenDeliverModel,omitempty" type:"Struct"`
	CoFeedOpenSpaceModel     *CreateAndDeliverWithDelegateRequestCoFeedOpenSpaceModel     `json:"coFeedOpenSpaceModel,omitempty" xml:"coFeedOpenSpaceModel,omitempty" type:"Struct"`
	DocOpenDeliverModel      *CreateAndDeliverWithDelegateRequestDocOpenDeliverModel      `json:"docOpenDeliverModel,omitempty" xml:"docOpenDeliverModel,omitempty" type:"Struct"`
	ImGroupOpenDeliverModel  *CreateAndDeliverWithDelegateRequestImGroupOpenDeliverModel  `json:"imGroupOpenDeliverModel,omitempty" xml:"imGroupOpenDeliverModel,omitempty" type:"Struct"`
	ImGroupOpenSpaceModel    *CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModel    `json:"imGroupOpenSpaceModel,omitempty" xml:"imGroupOpenSpaceModel,omitempty" type:"Struct"`
	ImRobotOpenDeliverModel  *CreateAndDeliverWithDelegateRequestImRobotOpenDeliverModel  `json:"imRobotOpenDeliverModel,omitempty" xml:"imRobotOpenDeliverModel,omitempty" type:"Struct"`
	ImRobotOpenSpaceModel    *CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModel    `json:"imRobotOpenSpaceModel,omitempty" xml:"imRobotOpenSpaceModel,omitempty" type:"Struct"`
	ImSingleOpenDeliverModel *CreateAndDeliverWithDelegateRequestImSingleOpenDeliverModel `json:"imSingleOpenDeliverModel,omitempty" xml:"imSingleOpenDeliverModel,omitempty" type:"Struct"`
	ImSingleOpenSpaceModel   *CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModel   `json:"imSingleOpenSpaceModel,omitempty" xml:"imSingleOpenSpaceModel,omitempty" type:"Struct"`
	OpenDynamicDataConfig    *CreateAndDeliverWithDelegateRequestOpenDynamicDataConfig    `json:"openDynamicDataConfig,omitempty" xml:"openDynamicDataConfig,omitempty" type:"Struct"`
	OpenSpaceId              *string                                                      `json:"openSpaceId,omitempty" xml:"openSpaceId,omitempty"`
	OutTrackId               *string                                                      `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	PrivateData              map[string]*PrivateDataValue                                 `json:"privateData,omitempty" xml:"privateData,omitempty"`
	TopOpenDeliverModel      *CreateAndDeliverWithDelegateRequestTopOpenDeliverModel      `json:"topOpenDeliverModel,omitempty" xml:"topOpenDeliverModel,omitempty" type:"Struct"`
	TopOpenSpaceModel        *CreateAndDeliverWithDelegateRequestTopOpenSpaceModel        `json:"topOpenSpaceModel,omitempty" xml:"topOpenSpaceModel,omitempty" type:"Struct"`
	UserId                   *string                                                      `json:"userId,omitempty" xml:"userId,omitempty"`
	UserIdType               *int32                                                       `json:"userIdType,omitempty" xml:"userIdType,omitempty"`
}

func (s CreateAndDeliverWithDelegateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateRequest) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateRequest) SetCallbackRouteKey(v string) *CreateAndDeliverWithDelegateRequest {
	s.CallbackRouteKey = &v
	return s
}

func (s *CreateAndDeliverWithDelegateRequest) SetCallbackType(v string) *CreateAndDeliverWithDelegateRequest {
	s.CallbackType = &v
	return s
}

func (s *CreateAndDeliverWithDelegateRequest) SetCardData(v *CreateAndDeliverWithDelegateRequestCardData) *CreateAndDeliverWithDelegateRequest {
	s.CardData = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequest) SetCardTemplateId(v string) *CreateAndDeliverWithDelegateRequest {
	s.CardTemplateId = &v
	return s
}

func (s *CreateAndDeliverWithDelegateRequest) SetCoFeedOpenDeliverModel(v *CreateAndDeliverWithDelegateRequestCoFeedOpenDeliverModel) *CreateAndDeliverWithDelegateRequest {
	s.CoFeedOpenDeliverModel = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequest) SetCoFeedOpenSpaceModel(v *CreateAndDeliverWithDelegateRequestCoFeedOpenSpaceModel) *CreateAndDeliverWithDelegateRequest {
	s.CoFeedOpenSpaceModel = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequest) SetDocOpenDeliverModel(v *CreateAndDeliverWithDelegateRequestDocOpenDeliverModel) *CreateAndDeliverWithDelegateRequest {
	s.DocOpenDeliverModel = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequest) SetImGroupOpenDeliverModel(v *CreateAndDeliverWithDelegateRequestImGroupOpenDeliverModel) *CreateAndDeliverWithDelegateRequest {
	s.ImGroupOpenDeliverModel = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequest) SetImGroupOpenSpaceModel(v *CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModel) *CreateAndDeliverWithDelegateRequest {
	s.ImGroupOpenSpaceModel = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequest) SetImRobotOpenDeliverModel(v *CreateAndDeliverWithDelegateRequestImRobotOpenDeliverModel) *CreateAndDeliverWithDelegateRequest {
	s.ImRobotOpenDeliverModel = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequest) SetImRobotOpenSpaceModel(v *CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModel) *CreateAndDeliverWithDelegateRequest {
	s.ImRobotOpenSpaceModel = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequest) SetImSingleOpenDeliverModel(v *CreateAndDeliverWithDelegateRequestImSingleOpenDeliverModel) *CreateAndDeliverWithDelegateRequest {
	s.ImSingleOpenDeliverModel = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequest) SetImSingleOpenSpaceModel(v *CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModel) *CreateAndDeliverWithDelegateRequest {
	s.ImSingleOpenSpaceModel = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequest) SetOpenDynamicDataConfig(v *CreateAndDeliverWithDelegateRequestOpenDynamicDataConfig) *CreateAndDeliverWithDelegateRequest {
	s.OpenDynamicDataConfig = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequest) SetOpenSpaceId(v string) *CreateAndDeliverWithDelegateRequest {
	s.OpenSpaceId = &v
	return s
}

func (s *CreateAndDeliverWithDelegateRequest) SetOutTrackId(v string) *CreateAndDeliverWithDelegateRequest {
	s.OutTrackId = &v
	return s
}

func (s *CreateAndDeliverWithDelegateRequest) SetPrivateData(v map[string]*PrivateDataValue) *CreateAndDeliverWithDelegateRequest {
	s.PrivateData = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequest) SetTopOpenDeliverModel(v *CreateAndDeliverWithDelegateRequestTopOpenDeliverModel) *CreateAndDeliverWithDelegateRequest {
	s.TopOpenDeliverModel = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequest) SetTopOpenSpaceModel(v *CreateAndDeliverWithDelegateRequestTopOpenSpaceModel) *CreateAndDeliverWithDelegateRequest {
	s.TopOpenSpaceModel = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequest) SetUserId(v string) *CreateAndDeliverWithDelegateRequest {
	s.UserId = &v
	return s
}

func (s *CreateAndDeliverWithDelegateRequest) SetUserIdType(v int32) *CreateAndDeliverWithDelegateRequest {
	s.UserIdType = &v
	return s
}

type CreateAndDeliverWithDelegateRequestCardData struct {
	CardParamMap map[string]*string `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
}

func (s CreateAndDeliverWithDelegateRequestCardData) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateRequestCardData) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateRequestCardData) SetCardParamMap(v map[string]*string) *CreateAndDeliverWithDelegateRequestCardData {
	s.CardParamMap = v
	return s
}

type CreateAndDeliverWithDelegateRequestCoFeedOpenDeliverModel struct {
	BizTag      *string `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	GmtTimeLine *int64  `json:"gmtTimeLine,omitempty" xml:"gmtTimeLine,omitempty"`
}

func (s CreateAndDeliverWithDelegateRequestCoFeedOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateRequestCoFeedOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateRequestCoFeedOpenDeliverModel) SetBizTag(v string) *CreateAndDeliverWithDelegateRequestCoFeedOpenDeliverModel {
	s.BizTag = &v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestCoFeedOpenDeliverModel) SetGmtTimeLine(v int64) *CreateAndDeliverWithDelegateRequestCoFeedOpenDeliverModel {
	s.GmtTimeLine = &v
	return s
}

type CreateAndDeliverWithDelegateRequestCoFeedOpenSpaceModel struct {
	CoolAppCode *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateAndDeliverWithDelegateRequestCoFeedOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateRequestCoFeedOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateRequestCoFeedOpenSpaceModel) SetCoolAppCode(v string) *CreateAndDeliverWithDelegateRequestCoFeedOpenSpaceModel {
	s.CoolAppCode = &v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestCoFeedOpenSpaceModel) SetTitle(v string) *CreateAndDeliverWithDelegateRequestCoFeedOpenSpaceModel {
	s.Title = &v
	return s
}

type CreateAndDeliverWithDelegateRequestDocOpenDeliverModel struct {
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateAndDeliverWithDelegateRequestDocOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateRequestDocOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateRequestDocOpenDeliverModel) SetUserId(v string) *CreateAndDeliverWithDelegateRequestDocOpenDeliverModel {
	s.UserId = &v
	return s
}

type CreateAndDeliverWithDelegateRequestImGroupOpenDeliverModel struct {
	AtUserIds  map[string]*string `json:"atUserIds,omitempty" xml:"atUserIds,omitempty"`
	Extension  map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
	Recipients []*string          `json:"recipients,omitempty" xml:"recipients,omitempty" type:"Repeated"`
	RobotCode  *string            `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s CreateAndDeliverWithDelegateRequestImGroupOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateRequestImGroupOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateRequestImGroupOpenDeliverModel) SetAtUserIds(v map[string]*string) *CreateAndDeliverWithDelegateRequestImGroupOpenDeliverModel {
	s.AtUserIds = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestImGroupOpenDeliverModel) SetExtension(v map[string]*string) *CreateAndDeliverWithDelegateRequestImGroupOpenDeliverModel {
	s.Extension = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestImGroupOpenDeliverModel) SetRecipients(v []*string) *CreateAndDeliverWithDelegateRequestImGroupOpenDeliverModel {
	s.Recipients = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestImGroupOpenDeliverModel) SetRobotCode(v string) *CreateAndDeliverWithDelegateRequestImGroupOpenDeliverModel {
	s.RobotCode = &v
	return s
}

type CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModel struct {
	LastMessageI18n map[string]*string                                                     `json:"lastMessageI18n,omitempty" xml:"lastMessageI18n,omitempty"`
	Notification    *CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModelNotification  `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	SearchSupport   *CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModelSearchSupport `json:"searchSupport,omitempty" xml:"searchSupport,omitempty" type:"Struct"`
	SupportForward  *bool                                                                  `json:"supportForward,omitempty" xml:"supportForward,omitempty"`
}

func (s CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModel) SetLastMessageI18n(v map[string]*string) *CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModel {
	s.LastMessageI18n = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModel) SetNotification(v *CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModelNotification) *CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModel {
	s.Notification = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModel) SetSearchSupport(v *CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModelSearchSupport) *CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModel {
	s.SearchSupport = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModel) SetSupportForward(v bool) *CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModel {
	s.SupportForward = &v
	return s
}

type CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModelNotification struct {
	AlertContent    *string `json:"alertContent,omitempty" xml:"alertContent,omitempty"`
	NotificationOff *bool   `json:"notificationOff,omitempty" xml:"notificationOff,omitempty"`
}

func (s CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModelNotification) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModelNotification) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModelNotification) SetAlertContent(v string) *CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModelNotification {
	s.AlertContent = &v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModelNotification) SetNotificationOff(v bool) *CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModelNotification {
	s.NotificationOff = &v
	return s
}

type CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModelSearchSupport struct {
	SearchDesc     *string `json:"searchDesc,omitempty" xml:"searchDesc,omitempty"`
	SearchIcon     *string `json:"searchIcon,omitempty" xml:"searchIcon,omitempty"`
	SearchTypeName *string `json:"searchTypeName,omitempty" xml:"searchTypeName,omitempty"`
}

func (s CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModelSearchSupport) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModelSearchSupport) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModelSearchSupport) SetSearchDesc(v string) *CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModelSearchSupport {
	s.SearchDesc = &v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModelSearchSupport) SetSearchIcon(v string) *CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModelSearchSupport {
	s.SearchIcon = &v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModelSearchSupport) SetSearchTypeName(v string) *CreateAndDeliverWithDelegateRequestImGroupOpenSpaceModelSearchSupport {
	s.SearchTypeName = &v
	return s
}

type CreateAndDeliverWithDelegateRequestImRobotOpenDeliverModel struct {
	Extension map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
	RobotCode *string            `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	SpaceType *string            `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
}

func (s CreateAndDeliverWithDelegateRequestImRobotOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateRequestImRobotOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateRequestImRobotOpenDeliverModel) SetExtension(v map[string]*string) *CreateAndDeliverWithDelegateRequestImRobotOpenDeliverModel {
	s.Extension = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestImRobotOpenDeliverModel) SetRobotCode(v string) *CreateAndDeliverWithDelegateRequestImRobotOpenDeliverModel {
	s.RobotCode = &v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestImRobotOpenDeliverModel) SetSpaceType(v string) *CreateAndDeliverWithDelegateRequestImRobotOpenDeliverModel {
	s.SpaceType = &v
	return s
}

type CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModel struct {
	LastMessageI18n map[string]*string                                                     `json:"lastMessageI18n,omitempty" xml:"lastMessageI18n,omitempty"`
	Notification    *CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModelNotification  `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	SearchSupport   *CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModelSearchSupport `json:"searchSupport,omitempty" xml:"searchSupport,omitempty" type:"Struct"`
	SupportForward  *bool                                                                  `json:"supportForward,omitempty" xml:"supportForward,omitempty"`
}

func (s CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModel) SetLastMessageI18n(v map[string]*string) *CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModel {
	s.LastMessageI18n = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModel) SetNotification(v *CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModelNotification) *CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModel {
	s.Notification = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModel) SetSearchSupport(v *CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModelSearchSupport) *CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModel {
	s.SearchSupport = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModel) SetSupportForward(v bool) *CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModel {
	s.SupportForward = &v
	return s
}

type CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModelNotification struct {
	AlertContent    *string `json:"alertContent,omitempty" xml:"alertContent,omitempty"`
	NotificationOff *bool   `json:"notificationOff,omitempty" xml:"notificationOff,omitempty"`
}

func (s CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModelNotification) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModelNotification) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModelNotification) SetAlertContent(v string) *CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModelNotification {
	s.AlertContent = &v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModelNotification) SetNotificationOff(v bool) *CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModelNotification {
	s.NotificationOff = &v
	return s
}

type CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModelSearchSupport struct {
	SearchDesc     *string `json:"searchDesc,omitempty" xml:"searchDesc,omitempty"`
	SearchIcon     *string `json:"searchIcon,omitempty" xml:"searchIcon,omitempty"`
	SearchTypeName *string `json:"searchTypeName,omitempty" xml:"searchTypeName,omitempty"`
}

func (s CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModelSearchSupport) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModelSearchSupport) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModelSearchSupport) SetSearchDesc(v string) *CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModelSearchSupport {
	s.SearchDesc = &v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModelSearchSupport) SetSearchIcon(v string) *CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModelSearchSupport {
	s.SearchIcon = &v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModelSearchSupport) SetSearchTypeName(v string) *CreateAndDeliverWithDelegateRequestImRobotOpenSpaceModelSearchSupport {
	s.SearchTypeName = &v
	return s
}

type CreateAndDeliverWithDelegateRequestImSingleOpenDeliverModel struct {
	AtUserIds map[string]*string `json:"atUserIds,omitempty" xml:"atUserIds,omitempty"`
	Extension map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
}

func (s CreateAndDeliverWithDelegateRequestImSingleOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateRequestImSingleOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateRequestImSingleOpenDeliverModel) SetAtUserIds(v map[string]*string) *CreateAndDeliverWithDelegateRequestImSingleOpenDeliverModel {
	s.AtUserIds = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestImSingleOpenDeliverModel) SetExtension(v map[string]*string) *CreateAndDeliverWithDelegateRequestImSingleOpenDeliverModel {
	s.Extension = v
	return s
}

type CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModel struct {
	LastMessageI18n map[string]*string                                                      `json:"lastMessageI18n,omitempty" xml:"lastMessageI18n,omitempty"`
	Notification    *CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModelNotification  `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	SearchSupport   *CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModelSearchSupport `json:"searchSupport,omitempty" xml:"searchSupport,omitempty" type:"Struct"`
	SupportForward  *bool                                                                   `json:"supportForward,omitempty" xml:"supportForward,omitempty"`
}

func (s CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModel) SetLastMessageI18n(v map[string]*string) *CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModel {
	s.LastMessageI18n = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModel) SetNotification(v *CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModelNotification) *CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModel {
	s.Notification = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModel) SetSearchSupport(v *CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModelSearchSupport) *CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModel {
	s.SearchSupport = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModel) SetSupportForward(v bool) *CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModel {
	s.SupportForward = &v
	return s
}

type CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModelNotification struct {
	AlertContent    *string `json:"alertContent,omitempty" xml:"alertContent,omitempty"`
	NotificationOff *bool   `json:"notificationOff,omitempty" xml:"notificationOff,omitempty"`
}

func (s CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModelNotification) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModelNotification) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModelNotification) SetAlertContent(v string) *CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModelNotification {
	s.AlertContent = &v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModelNotification) SetNotificationOff(v bool) *CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModelNotification {
	s.NotificationOff = &v
	return s
}

type CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModelSearchSupport struct {
	SearchDesc     *string `json:"searchDesc,omitempty" xml:"searchDesc,omitempty"`
	SearchIcon     *string `json:"searchIcon,omitempty" xml:"searchIcon,omitempty"`
	SearchTypeName *string `json:"searchTypeName,omitempty" xml:"searchTypeName,omitempty"`
}

func (s CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModelSearchSupport) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModelSearchSupport) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModelSearchSupport) SetSearchDesc(v string) *CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModelSearchSupport {
	s.SearchDesc = &v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModelSearchSupport) SetSearchIcon(v string) *CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModelSearchSupport {
	s.SearchIcon = &v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModelSearchSupport) SetSearchTypeName(v string) *CreateAndDeliverWithDelegateRequestImSingleOpenSpaceModelSearchSupport {
	s.SearchTypeName = &v
	return s
}

type CreateAndDeliverWithDelegateRequestOpenDynamicDataConfig struct {
	DynamicDataSourceConfigs []*CreateAndDeliverWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigs `json:"dynamicDataSourceConfigs,omitempty" xml:"dynamicDataSourceConfigs,omitempty" type:"Repeated"`
}

func (s CreateAndDeliverWithDelegateRequestOpenDynamicDataConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateRequestOpenDynamicDataConfig) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateRequestOpenDynamicDataConfig) SetDynamicDataSourceConfigs(v []*CreateAndDeliverWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigs) *CreateAndDeliverWithDelegateRequestOpenDynamicDataConfig {
	s.DynamicDataSourceConfigs = v
	return s
}

type CreateAndDeliverWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigs struct {
	ConstParams         map[string]*string                                                                          `json:"constParams,omitempty" xml:"constParams,omitempty"`
	DynamicDataSourceId *string                                                                                     `json:"dynamicDataSourceId,omitempty" xml:"dynamicDataSourceId,omitempty"`
	PullConfig          *CreateAndDeliverWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig `json:"pullConfig,omitempty" xml:"pullConfig,omitempty" type:"Struct"`
}

func (s CreateAndDeliverWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigs) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigs) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigs) SetConstParams(v map[string]*string) *CreateAndDeliverWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigs {
	s.ConstParams = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigs) SetDynamicDataSourceId(v string) *CreateAndDeliverWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigs {
	s.DynamicDataSourceId = &v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigs) SetPullConfig(v *CreateAndDeliverWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig) *CreateAndDeliverWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigs {
	s.PullConfig = v
	return s
}

type CreateAndDeliverWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig struct {
	Interval     *int32  `json:"interval,omitempty" xml:"interval,omitempty"`
	PullStrategy *string `json:"pullStrategy,omitempty" xml:"pullStrategy,omitempty"`
	TimeUnit     *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
}

func (s CreateAndDeliverWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig) SetInterval(v int32) *CreateAndDeliverWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig {
	s.Interval = &v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig) SetPullStrategy(v string) *CreateAndDeliverWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig {
	s.PullStrategy = &v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig) SetTimeUnit(v string) *CreateAndDeliverWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig {
	s.TimeUnit = &v
	return s
}

type CreateAndDeliverWithDelegateRequestTopOpenDeliverModel struct {
	ExpiredTimeMillis *int64    `json:"expiredTimeMillis,omitempty" xml:"expiredTimeMillis,omitempty"`
	Platforms         []*string `json:"platforms,omitempty" xml:"platforms,omitempty" type:"Repeated"`
	UserIds           []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s CreateAndDeliverWithDelegateRequestTopOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateRequestTopOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateRequestTopOpenDeliverModel) SetExpiredTimeMillis(v int64) *CreateAndDeliverWithDelegateRequestTopOpenDeliverModel {
	s.ExpiredTimeMillis = &v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestTopOpenDeliverModel) SetPlatforms(v []*string) *CreateAndDeliverWithDelegateRequestTopOpenDeliverModel {
	s.Platforms = v
	return s
}

func (s *CreateAndDeliverWithDelegateRequestTopOpenDeliverModel) SetUserIds(v []*string) *CreateAndDeliverWithDelegateRequestTopOpenDeliverModel {
	s.UserIds = v
	return s
}

type CreateAndDeliverWithDelegateRequestTopOpenSpaceModel struct {
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
}

func (s CreateAndDeliverWithDelegateRequestTopOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateRequestTopOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateRequestTopOpenSpaceModel) SetSpaceType(v string) *CreateAndDeliverWithDelegateRequestTopOpenSpaceModel {
	s.SpaceType = &v
	return s
}

type CreateAndDeliverWithDelegateResponseBody struct {
	Result  *CreateAndDeliverWithDelegateResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateAndDeliverWithDelegateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateResponseBody) SetResult(v *CreateAndDeliverWithDelegateResponseBodyResult) *CreateAndDeliverWithDelegateResponseBody {
	s.Result = v
	return s
}

func (s *CreateAndDeliverWithDelegateResponseBody) SetSuccess(v bool) *CreateAndDeliverWithDelegateResponseBody {
	s.Success = &v
	return s
}

type CreateAndDeliverWithDelegateResponseBodyResult struct {
	DeliverResults []*CreateAndDeliverWithDelegateResponseBodyResultDeliverResults `json:"deliverResults,omitempty" xml:"deliverResults,omitempty" type:"Repeated"`
	OutTrackId     *string                                                         `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
}

func (s CreateAndDeliverWithDelegateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateResponseBodyResult) SetDeliverResults(v []*CreateAndDeliverWithDelegateResponseBodyResultDeliverResults) *CreateAndDeliverWithDelegateResponseBodyResult {
	s.DeliverResults = v
	return s
}

func (s *CreateAndDeliverWithDelegateResponseBodyResult) SetOutTrackId(v string) *CreateAndDeliverWithDelegateResponseBodyResult {
	s.OutTrackId = &v
	return s
}

type CreateAndDeliverWithDelegateResponseBodyResultDeliverResults struct {
	CarrierId *string `json:"carrierId,omitempty" xml:"carrierId,omitempty"`
	ErrorMsg  *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	SpaceId   *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateAndDeliverWithDelegateResponseBodyResultDeliverResults) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateResponseBodyResultDeliverResults) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateResponseBodyResultDeliverResults) SetCarrierId(v string) *CreateAndDeliverWithDelegateResponseBodyResultDeliverResults {
	s.CarrierId = &v
	return s
}

func (s *CreateAndDeliverWithDelegateResponseBodyResultDeliverResults) SetErrorMsg(v string) *CreateAndDeliverWithDelegateResponseBodyResultDeliverResults {
	s.ErrorMsg = &v
	return s
}

func (s *CreateAndDeliverWithDelegateResponseBodyResultDeliverResults) SetSpaceId(v string) *CreateAndDeliverWithDelegateResponseBodyResultDeliverResults {
	s.SpaceId = &v
	return s
}

func (s *CreateAndDeliverWithDelegateResponseBodyResultDeliverResults) SetSpaceType(v string) *CreateAndDeliverWithDelegateResponseBodyResultDeliverResults {
	s.SpaceType = &v
	return s
}

func (s *CreateAndDeliverWithDelegateResponseBodyResultDeliverResults) SetSuccess(v bool) *CreateAndDeliverWithDelegateResponseBodyResultDeliverResults {
	s.Success = &v
	return s
}

type CreateAndDeliverWithDelegateResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAndDeliverWithDelegateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAndDeliverWithDelegateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAndDeliverWithDelegateResponse) GoString() string {
	return s.String()
}

func (s *CreateAndDeliverWithDelegateResponse) SetHeaders(v map[string]*string) *CreateAndDeliverWithDelegateResponse {
	s.Headers = v
	return s
}

func (s *CreateAndDeliverWithDelegateResponse) SetStatusCode(v int32) *CreateAndDeliverWithDelegateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAndDeliverWithDelegateResponse) SetBody(v *CreateAndDeliverWithDelegateResponseBody) *CreateAndDeliverWithDelegateResponse {
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
	CallbackRouteKey       *string                                  `json:"callbackRouteKey,omitempty" xml:"callbackRouteKey,omitempty"`
	CallbackType           *string                                  `json:"callbackType,omitempty" xml:"callbackType,omitempty"`
	CardData               *CreateCardRequestCardData               `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	CardTemplateId         *string                                  `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	CoFeedOpenSpaceModel   *CreateCardRequestCoFeedOpenSpaceModel   `json:"coFeedOpenSpaceModel,omitempty" xml:"coFeedOpenSpaceModel,omitempty" type:"Struct"`
	ImGroupOpenSpaceModel  *CreateCardRequestImGroupOpenSpaceModel  `json:"imGroupOpenSpaceModel,omitempty" xml:"imGroupOpenSpaceModel,omitempty" type:"Struct"`
	ImRobotOpenSpaceModel  *CreateCardRequestImRobotOpenSpaceModel  `json:"imRobotOpenSpaceModel,omitempty" xml:"imRobotOpenSpaceModel,omitempty" type:"Struct"`
	ImSingleOpenSpaceModel *CreateCardRequestImSingleOpenSpaceModel `json:"imSingleOpenSpaceModel,omitempty" xml:"imSingleOpenSpaceModel,omitempty" type:"Struct"`
	OpenDynamicDataConfig  *CreateCardRequestOpenDynamicDataConfig  `json:"openDynamicDataConfig,omitempty" xml:"openDynamicDataConfig,omitempty" type:"Struct"`
	OutTrackId             *string                                  `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	PrivateData            map[string]*PrivateDataValue             `json:"privateData,omitempty" xml:"privateData,omitempty"`
	TopOpenSpaceModel      *CreateCardRequestTopOpenSpaceModel      `json:"topOpenSpaceModel,omitempty" xml:"topOpenSpaceModel,omitempty" type:"Struct"`
	UserId                 *string                                  `json:"userId,omitempty" xml:"userId,omitempty"`
	UserIdType             *int32                                   `json:"userIdType,omitempty" xml:"userIdType,omitempty"`
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

type CreateCardRequestImSingleOpenSpaceModel struct {
	LastMessageI18n map[string]*string                                    `json:"lastMessageI18n,omitempty" xml:"lastMessageI18n,omitempty"`
	Notification    *CreateCardRequestImSingleOpenSpaceModelNotification  `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	SearchSupport   *CreateCardRequestImSingleOpenSpaceModelSearchSupport `json:"searchSupport,omitempty" xml:"searchSupport,omitempty" type:"Struct"`
	SupportForward  *bool                                                 `json:"supportForward,omitempty" xml:"supportForward,omitempty"`
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
	AlertContent    *string `json:"alertContent,omitempty" xml:"alertContent,omitempty"`
	NotificationOff *bool   `json:"notificationOff,omitempty" xml:"notificationOff,omitempty"`
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
	SearchDesc     *string `json:"searchDesc,omitempty" xml:"searchDesc,omitempty"`
	SearchIcon     *string `json:"searchIcon,omitempty" xml:"searchIcon,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CreateCardWithDelegateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateCardWithDelegateHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateCardWithDelegateHeaders) GoString() string {
	return s.String()
}

func (s *CreateCardWithDelegateHeaders) SetCommonHeaders(v map[string]*string) *CreateCardWithDelegateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateCardWithDelegateHeaders) SetXAcsDingtalkAccessToken(v string) *CreateCardWithDelegateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateCardWithDelegateRequest struct {
	CallbackRouteKey       *string                                              `json:"callbackRouteKey,omitempty" xml:"callbackRouteKey,omitempty"`
	CallbackType           *string                                              `json:"callbackType,omitempty" xml:"callbackType,omitempty"`
	CardData               *CreateCardWithDelegateRequestCardData               `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	CardTemplateId         *string                                              `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	CoFeedOpenSpaceModel   *CreateCardWithDelegateRequestCoFeedOpenSpaceModel   `json:"coFeedOpenSpaceModel,omitempty" xml:"coFeedOpenSpaceModel,omitempty" type:"Struct"`
	ImGroupOpenSpaceModel  *CreateCardWithDelegateRequestImGroupOpenSpaceModel  `json:"imGroupOpenSpaceModel,omitempty" xml:"imGroupOpenSpaceModel,omitempty" type:"Struct"`
	ImRobotOpenSpaceModel  *CreateCardWithDelegateRequestImRobotOpenSpaceModel  `json:"imRobotOpenSpaceModel,omitempty" xml:"imRobotOpenSpaceModel,omitempty" type:"Struct"`
	ImSingleOpenSpaceModel *CreateCardWithDelegateRequestImSingleOpenSpaceModel `json:"imSingleOpenSpaceModel,omitempty" xml:"imSingleOpenSpaceModel,omitempty" type:"Struct"`
	OpenDynamicDataConfig  *CreateCardWithDelegateRequestOpenDynamicDataConfig  `json:"openDynamicDataConfig,omitempty" xml:"openDynamicDataConfig,omitempty" type:"Struct"`
	OutTrackId             *string                                              `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	PrivateData            map[string]*PrivateDataValue                         `json:"privateData,omitempty" xml:"privateData,omitempty"`
	TopOpenSpaceModel      *CreateCardWithDelegateRequestTopOpenSpaceModel      `json:"topOpenSpaceModel,omitempty" xml:"topOpenSpaceModel,omitempty" type:"Struct"`
	UserId                 *string                                              `json:"userId,omitempty" xml:"userId,omitempty"`
	UserIdType             *int32                                               `json:"userIdType,omitempty" xml:"userIdType,omitempty"`
}

func (s CreateCardWithDelegateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCardWithDelegateRequest) GoString() string {
	return s.String()
}

func (s *CreateCardWithDelegateRequest) SetCallbackRouteKey(v string) *CreateCardWithDelegateRequest {
	s.CallbackRouteKey = &v
	return s
}

func (s *CreateCardWithDelegateRequest) SetCallbackType(v string) *CreateCardWithDelegateRequest {
	s.CallbackType = &v
	return s
}

func (s *CreateCardWithDelegateRequest) SetCardData(v *CreateCardWithDelegateRequestCardData) *CreateCardWithDelegateRequest {
	s.CardData = v
	return s
}

func (s *CreateCardWithDelegateRequest) SetCardTemplateId(v string) *CreateCardWithDelegateRequest {
	s.CardTemplateId = &v
	return s
}

func (s *CreateCardWithDelegateRequest) SetCoFeedOpenSpaceModel(v *CreateCardWithDelegateRequestCoFeedOpenSpaceModel) *CreateCardWithDelegateRequest {
	s.CoFeedOpenSpaceModel = v
	return s
}

func (s *CreateCardWithDelegateRequest) SetImGroupOpenSpaceModel(v *CreateCardWithDelegateRequestImGroupOpenSpaceModel) *CreateCardWithDelegateRequest {
	s.ImGroupOpenSpaceModel = v
	return s
}

func (s *CreateCardWithDelegateRequest) SetImRobotOpenSpaceModel(v *CreateCardWithDelegateRequestImRobotOpenSpaceModel) *CreateCardWithDelegateRequest {
	s.ImRobotOpenSpaceModel = v
	return s
}

func (s *CreateCardWithDelegateRequest) SetImSingleOpenSpaceModel(v *CreateCardWithDelegateRequestImSingleOpenSpaceModel) *CreateCardWithDelegateRequest {
	s.ImSingleOpenSpaceModel = v
	return s
}

func (s *CreateCardWithDelegateRequest) SetOpenDynamicDataConfig(v *CreateCardWithDelegateRequestOpenDynamicDataConfig) *CreateCardWithDelegateRequest {
	s.OpenDynamicDataConfig = v
	return s
}

func (s *CreateCardWithDelegateRequest) SetOutTrackId(v string) *CreateCardWithDelegateRequest {
	s.OutTrackId = &v
	return s
}

func (s *CreateCardWithDelegateRequest) SetPrivateData(v map[string]*PrivateDataValue) *CreateCardWithDelegateRequest {
	s.PrivateData = v
	return s
}

func (s *CreateCardWithDelegateRequest) SetTopOpenSpaceModel(v *CreateCardWithDelegateRequestTopOpenSpaceModel) *CreateCardWithDelegateRequest {
	s.TopOpenSpaceModel = v
	return s
}

func (s *CreateCardWithDelegateRequest) SetUserId(v string) *CreateCardWithDelegateRequest {
	s.UserId = &v
	return s
}

func (s *CreateCardWithDelegateRequest) SetUserIdType(v int32) *CreateCardWithDelegateRequest {
	s.UserIdType = &v
	return s
}

type CreateCardWithDelegateRequestCardData struct {
	CardParamMap map[string]*string `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
}

func (s CreateCardWithDelegateRequestCardData) String() string {
	return tea.Prettify(s)
}

func (s CreateCardWithDelegateRequestCardData) GoString() string {
	return s.String()
}

func (s *CreateCardWithDelegateRequestCardData) SetCardParamMap(v map[string]*string) *CreateCardWithDelegateRequestCardData {
	s.CardParamMap = v
	return s
}

type CreateCardWithDelegateRequestCoFeedOpenSpaceModel struct {
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateCardWithDelegateRequestCoFeedOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s CreateCardWithDelegateRequestCoFeedOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *CreateCardWithDelegateRequestCoFeedOpenSpaceModel) SetTitle(v string) *CreateCardWithDelegateRequestCoFeedOpenSpaceModel {
	s.Title = &v
	return s
}

type CreateCardWithDelegateRequestImGroupOpenSpaceModel struct {
	LastMessageI18n map[string]*string                                               `json:"lastMessageI18n,omitempty" xml:"lastMessageI18n,omitempty"`
	Notification    *CreateCardWithDelegateRequestImGroupOpenSpaceModelNotification  `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	SearchSupport   *CreateCardWithDelegateRequestImGroupOpenSpaceModelSearchSupport `json:"searchSupport,omitempty" xml:"searchSupport,omitempty" type:"Struct"`
	SupportForward  *bool                                                            `json:"supportForward,omitempty" xml:"supportForward,omitempty"`
}

func (s CreateCardWithDelegateRequestImGroupOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s CreateCardWithDelegateRequestImGroupOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *CreateCardWithDelegateRequestImGroupOpenSpaceModel) SetLastMessageI18n(v map[string]*string) *CreateCardWithDelegateRequestImGroupOpenSpaceModel {
	s.LastMessageI18n = v
	return s
}

func (s *CreateCardWithDelegateRequestImGroupOpenSpaceModel) SetNotification(v *CreateCardWithDelegateRequestImGroupOpenSpaceModelNotification) *CreateCardWithDelegateRequestImGroupOpenSpaceModel {
	s.Notification = v
	return s
}

func (s *CreateCardWithDelegateRequestImGroupOpenSpaceModel) SetSearchSupport(v *CreateCardWithDelegateRequestImGroupOpenSpaceModelSearchSupport) *CreateCardWithDelegateRequestImGroupOpenSpaceModel {
	s.SearchSupport = v
	return s
}

func (s *CreateCardWithDelegateRequestImGroupOpenSpaceModel) SetSupportForward(v bool) *CreateCardWithDelegateRequestImGroupOpenSpaceModel {
	s.SupportForward = &v
	return s
}

type CreateCardWithDelegateRequestImGroupOpenSpaceModelNotification struct {
	AlertContent    *string `json:"alertContent,omitempty" xml:"alertContent,omitempty"`
	NotificationOff *bool   `json:"notificationOff,omitempty" xml:"notificationOff,omitempty"`
}

func (s CreateCardWithDelegateRequestImGroupOpenSpaceModelNotification) String() string {
	return tea.Prettify(s)
}

func (s CreateCardWithDelegateRequestImGroupOpenSpaceModelNotification) GoString() string {
	return s.String()
}

func (s *CreateCardWithDelegateRequestImGroupOpenSpaceModelNotification) SetAlertContent(v string) *CreateCardWithDelegateRequestImGroupOpenSpaceModelNotification {
	s.AlertContent = &v
	return s
}

func (s *CreateCardWithDelegateRequestImGroupOpenSpaceModelNotification) SetNotificationOff(v bool) *CreateCardWithDelegateRequestImGroupOpenSpaceModelNotification {
	s.NotificationOff = &v
	return s
}

type CreateCardWithDelegateRequestImGroupOpenSpaceModelSearchSupport struct {
	SearchDesc     *string `json:"searchDesc,omitempty" xml:"searchDesc,omitempty"`
	SearchIcon     *string `json:"searchIcon,omitempty" xml:"searchIcon,omitempty"`
	SearchTypeName *string `json:"searchTypeName,omitempty" xml:"searchTypeName,omitempty"`
}

func (s CreateCardWithDelegateRequestImGroupOpenSpaceModelSearchSupport) String() string {
	return tea.Prettify(s)
}

func (s CreateCardWithDelegateRequestImGroupOpenSpaceModelSearchSupport) GoString() string {
	return s.String()
}

func (s *CreateCardWithDelegateRequestImGroupOpenSpaceModelSearchSupport) SetSearchDesc(v string) *CreateCardWithDelegateRequestImGroupOpenSpaceModelSearchSupport {
	s.SearchDesc = &v
	return s
}

func (s *CreateCardWithDelegateRequestImGroupOpenSpaceModelSearchSupport) SetSearchIcon(v string) *CreateCardWithDelegateRequestImGroupOpenSpaceModelSearchSupport {
	s.SearchIcon = &v
	return s
}

func (s *CreateCardWithDelegateRequestImGroupOpenSpaceModelSearchSupport) SetSearchTypeName(v string) *CreateCardWithDelegateRequestImGroupOpenSpaceModelSearchSupport {
	s.SearchTypeName = &v
	return s
}

type CreateCardWithDelegateRequestImRobotOpenSpaceModel struct {
	LastMessageI18n map[string]*string                                               `json:"lastMessageI18n,omitempty" xml:"lastMessageI18n,omitempty"`
	Notification    *CreateCardWithDelegateRequestImRobotOpenSpaceModelNotification  `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	SearchSupport   *CreateCardWithDelegateRequestImRobotOpenSpaceModelSearchSupport `json:"searchSupport,omitempty" xml:"searchSupport,omitempty" type:"Struct"`
	SupportForward  *bool                                                            `json:"supportForward,omitempty" xml:"supportForward,omitempty"`
}

func (s CreateCardWithDelegateRequestImRobotOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s CreateCardWithDelegateRequestImRobotOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *CreateCardWithDelegateRequestImRobotOpenSpaceModel) SetLastMessageI18n(v map[string]*string) *CreateCardWithDelegateRequestImRobotOpenSpaceModel {
	s.LastMessageI18n = v
	return s
}

func (s *CreateCardWithDelegateRequestImRobotOpenSpaceModel) SetNotification(v *CreateCardWithDelegateRequestImRobotOpenSpaceModelNotification) *CreateCardWithDelegateRequestImRobotOpenSpaceModel {
	s.Notification = v
	return s
}

func (s *CreateCardWithDelegateRequestImRobotOpenSpaceModel) SetSearchSupport(v *CreateCardWithDelegateRequestImRobotOpenSpaceModelSearchSupport) *CreateCardWithDelegateRequestImRobotOpenSpaceModel {
	s.SearchSupport = v
	return s
}

func (s *CreateCardWithDelegateRequestImRobotOpenSpaceModel) SetSupportForward(v bool) *CreateCardWithDelegateRequestImRobotOpenSpaceModel {
	s.SupportForward = &v
	return s
}

type CreateCardWithDelegateRequestImRobotOpenSpaceModelNotification struct {
	AlertContent    *string `json:"alertContent,omitempty" xml:"alertContent,omitempty"`
	NotificationOff *bool   `json:"notificationOff,omitempty" xml:"notificationOff,omitempty"`
}

func (s CreateCardWithDelegateRequestImRobotOpenSpaceModelNotification) String() string {
	return tea.Prettify(s)
}

func (s CreateCardWithDelegateRequestImRobotOpenSpaceModelNotification) GoString() string {
	return s.String()
}

func (s *CreateCardWithDelegateRequestImRobotOpenSpaceModelNotification) SetAlertContent(v string) *CreateCardWithDelegateRequestImRobotOpenSpaceModelNotification {
	s.AlertContent = &v
	return s
}

func (s *CreateCardWithDelegateRequestImRobotOpenSpaceModelNotification) SetNotificationOff(v bool) *CreateCardWithDelegateRequestImRobotOpenSpaceModelNotification {
	s.NotificationOff = &v
	return s
}

type CreateCardWithDelegateRequestImRobotOpenSpaceModelSearchSupport struct {
	SearchDesc     *string `json:"searchDesc,omitempty" xml:"searchDesc,omitempty"`
	SearchIcon     *string `json:"searchIcon,omitempty" xml:"searchIcon,omitempty"`
	SearchTypeName *string `json:"searchTypeName,omitempty" xml:"searchTypeName,omitempty"`
}

func (s CreateCardWithDelegateRequestImRobotOpenSpaceModelSearchSupport) String() string {
	return tea.Prettify(s)
}

func (s CreateCardWithDelegateRequestImRobotOpenSpaceModelSearchSupport) GoString() string {
	return s.String()
}

func (s *CreateCardWithDelegateRequestImRobotOpenSpaceModelSearchSupport) SetSearchDesc(v string) *CreateCardWithDelegateRequestImRobotOpenSpaceModelSearchSupport {
	s.SearchDesc = &v
	return s
}

func (s *CreateCardWithDelegateRequestImRobotOpenSpaceModelSearchSupport) SetSearchIcon(v string) *CreateCardWithDelegateRequestImRobotOpenSpaceModelSearchSupport {
	s.SearchIcon = &v
	return s
}

func (s *CreateCardWithDelegateRequestImRobotOpenSpaceModelSearchSupport) SetSearchTypeName(v string) *CreateCardWithDelegateRequestImRobotOpenSpaceModelSearchSupport {
	s.SearchTypeName = &v
	return s
}

type CreateCardWithDelegateRequestImSingleOpenSpaceModel struct {
	LastMessageI18n map[string]*string                                                `json:"lastMessageI18n,omitempty" xml:"lastMessageI18n,omitempty"`
	Notification    *CreateCardWithDelegateRequestImSingleOpenSpaceModelNotification  `json:"notification,omitempty" xml:"notification,omitempty" type:"Struct"`
	SearchSupport   *CreateCardWithDelegateRequestImSingleOpenSpaceModelSearchSupport `json:"searchSupport,omitempty" xml:"searchSupport,omitempty" type:"Struct"`
	SupportForward  *bool                                                             `json:"supportForward,omitempty" xml:"supportForward,omitempty"`
}

func (s CreateCardWithDelegateRequestImSingleOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s CreateCardWithDelegateRequestImSingleOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *CreateCardWithDelegateRequestImSingleOpenSpaceModel) SetLastMessageI18n(v map[string]*string) *CreateCardWithDelegateRequestImSingleOpenSpaceModel {
	s.LastMessageI18n = v
	return s
}

func (s *CreateCardWithDelegateRequestImSingleOpenSpaceModel) SetNotification(v *CreateCardWithDelegateRequestImSingleOpenSpaceModelNotification) *CreateCardWithDelegateRequestImSingleOpenSpaceModel {
	s.Notification = v
	return s
}

func (s *CreateCardWithDelegateRequestImSingleOpenSpaceModel) SetSearchSupport(v *CreateCardWithDelegateRequestImSingleOpenSpaceModelSearchSupport) *CreateCardWithDelegateRequestImSingleOpenSpaceModel {
	s.SearchSupport = v
	return s
}

func (s *CreateCardWithDelegateRequestImSingleOpenSpaceModel) SetSupportForward(v bool) *CreateCardWithDelegateRequestImSingleOpenSpaceModel {
	s.SupportForward = &v
	return s
}

type CreateCardWithDelegateRequestImSingleOpenSpaceModelNotification struct {
	AlertContent    *string `json:"alertContent,omitempty" xml:"alertContent,omitempty"`
	NotificationOff *bool   `json:"notificationOff,omitempty" xml:"notificationOff,omitempty"`
}

func (s CreateCardWithDelegateRequestImSingleOpenSpaceModelNotification) String() string {
	return tea.Prettify(s)
}

func (s CreateCardWithDelegateRequestImSingleOpenSpaceModelNotification) GoString() string {
	return s.String()
}

func (s *CreateCardWithDelegateRequestImSingleOpenSpaceModelNotification) SetAlertContent(v string) *CreateCardWithDelegateRequestImSingleOpenSpaceModelNotification {
	s.AlertContent = &v
	return s
}

func (s *CreateCardWithDelegateRequestImSingleOpenSpaceModelNotification) SetNotificationOff(v bool) *CreateCardWithDelegateRequestImSingleOpenSpaceModelNotification {
	s.NotificationOff = &v
	return s
}

type CreateCardWithDelegateRequestImSingleOpenSpaceModelSearchSupport struct {
	SearchDesc     *string `json:"searchDesc,omitempty" xml:"searchDesc,omitempty"`
	SearchIcon     *string `json:"searchIcon,omitempty" xml:"searchIcon,omitempty"`
	SearchTypeName *string `json:"searchTypeName,omitempty" xml:"searchTypeName,omitempty"`
}

func (s CreateCardWithDelegateRequestImSingleOpenSpaceModelSearchSupport) String() string {
	return tea.Prettify(s)
}

func (s CreateCardWithDelegateRequestImSingleOpenSpaceModelSearchSupport) GoString() string {
	return s.String()
}

func (s *CreateCardWithDelegateRequestImSingleOpenSpaceModelSearchSupport) SetSearchDesc(v string) *CreateCardWithDelegateRequestImSingleOpenSpaceModelSearchSupport {
	s.SearchDesc = &v
	return s
}

func (s *CreateCardWithDelegateRequestImSingleOpenSpaceModelSearchSupport) SetSearchIcon(v string) *CreateCardWithDelegateRequestImSingleOpenSpaceModelSearchSupport {
	s.SearchIcon = &v
	return s
}

func (s *CreateCardWithDelegateRequestImSingleOpenSpaceModelSearchSupport) SetSearchTypeName(v string) *CreateCardWithDelegateRequestImSingleOpenSpaceModelSearchSupport {
	s.SearchTypeName = &v
	return s
}

type CreateCardWithDelegateRequestOpenDynamicDataConfig struct {
	DynamicDataSourceConfigs []*CreateCardWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigs `json:"dynamicDataSourceConfigs,omitempty" xml:"dynamicDataSourceConfigs,omitempty" type:"Repeated"`
}

func (s CreateCardWithDelegateRequestOpenDynamicDataConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateCardWithDelegateRequestOpenDynamicDataConfig) GoString() string {
	return s.String()
}

func (s *CreateCardWithDelegateRequestOpenDynamicDataConfig) SetDynamicDataSourceConfigs(v []*CreateCardWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigs) *CreateCardWithDelegateRequestOpenDynamicDataConfig {
	s.DynamicDataSourceConfigs = v
	return s
}

type CreateCardWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigs struct {
	ConstParams         map[string]*string                                                                    `json:"constParams,omitempty" xml:"constParams,omitempty"`
	DynamicDataSourceId *string                                                                               `json:"dynamicDataSourceId,omitempty" xml:"dynamicDataSourceId,omitempty"`
	PullConfig          *CreateCardWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig `json:"pullConfig,omitempty" xml:"pullConfig,omitempty" type:"Struct"`
}

func (s CreateCardWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigs) String() string {
	return tea.Prettify(s)
}

func (s CreateCardWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigs) GoString() string {
	return s.String()
}

func (s *CreateCardWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigs) SetConstParams(v map[string]*string) *CreateCardWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigs {
	s.ConstParams = v
	return s
}

func (s *CreateCardWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigs) SetDynamicDataSourceId(v string) *CreateCardWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigs {
	s.DynamicDataSourceId = &v
	return s
}

func (s *CreateCardWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigs) SetPullConfig(v *CreateCardWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig) *CreateCardWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigs {
	s.PullConfig = v
	return s
}

type CreateCardWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig struct {
	Interval     *int32  `json:"interval,omitempty" xml:"interval,omitempty"`
	PullStrategy *string `json:"pullStrategy,omitempty" xml:"pullStrategy,omitempty"`
	TimeUnit     *string `json:"timeUnit,omitempty" xml:"timeUnit,omitempty"`
}

func (s CreateCardWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateCardWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig) GoString() string {
	return s.String()
}

func (s *CreateCardWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig) SetInterval(v int32) *CreateCardWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig {
	s.Interval = &v
	return s
}

func (s *CreateCardWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig) SetPullStrategy(v string) *CreateCardWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig {
	s.PullStrategy = &v
	return s
}

func (s *CreateCardWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig) SetTimeUnit(v string) *CreateCardWithDelegateRequestOpenDynamicDataConfigDynamicDataSourceConfigsPullConfig {
	s.TimeUnit = &v
	return s
}

type CreateCardWithDelegateRequestTopOpenSpaceModel struct {
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
}

func (s CreateCardWithDelegateRequestTopOpenSpaceModel) String() string {
	return tea.Prettify(s)
}

func (s CreateCardWithDelegateRequestTopOpenSpaceModel) GoString() string {
	return s.String()
}

func (s *CreateCardWithDelegateRequestTopOpenSpaceModel) SetSpaceType(v string) *CreateCardWithDelegateRequestTopOpenSpaceModel {
	s.SpaceType = &v
	return s
}

type CreateCardWithDelegateResponseBody struct {
	Result  *string `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateCardWithDelegateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCardWithDelegateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCardWithDelegateResponseBody) SetResult(v string) *CreateCardWithDelegateResponseBody {
	s.Result = &v
	return s
}

func (s *CreateCardWithDelegateResponseBody) SetSuccess(v bool) *CreateCardWithDelegateResponseBody {
	s.Success = &v
	return s
}

type CreateCardWithDelegateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCardWithDelegateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCardWithDelegateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCardWithDelegateResponse) GoString() string {
	return s.String()
}

func (s *CreateCardWithDelegateResponse) SetHeaders(v map[string]*string) *CreateCardWithDelegateResponse {
	s.Headers = v
	return s
}

func (s *CreateCardWithDelegateResponse) SetStatusCode(v int32) *CreateCardWithDelegateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCardWithDelegateResponse) SetBody(v *CreateCardWithDelegateResponseBody) *CreateCardWithDelegateResponse {
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
	CoFeedOpenDeliverModel   *DeliverCardRequestCoFeedOpenDeliverModel   `json:"coFeedOpenDeliverModel,omitempty" xml:"coFeedOpenDeliverModel,omitempty" type:"Struct"`
	DocOpenDeliverModel      *DeliverCardRequestDocOpenDeliverModel      `json:"docOpenDeliverModel,omitempty" xml:"docOpenDeliverModel,omitempty" type:"Struct"`
	ImGroupOpenDeliverModel  *DeliverCardRequestImGroupOpenDeliverModel  `json:"imGroupOpenDeliverModel,omitempty" xml:"imGroupOpenDeliverModel,omitempty" type:"Struct"`
	ImRobotOpenDeliverModel  *DeliverCardRequestImRobotOpenDeliverModel  `json:"imRobotOpenDeliverModel,omitempty" xml:"imRobotOpenDeliverModel,omitempty" type:"Struct"`
	ImSingleOpenDeliverModel *DeliverCardRequestImSingleOpenDeliverModel `json:"imSingleOpenDeliverModel,omitempty" xml:"imSingleOpenDeliverModel,omitempty" type:"Struct"`
	OpenSpaceId              *string                                     `json:"openSpaceId,omitempty" xml:"openSpaceId,omitempty"`
	OutTrackId               *string                                     `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	TopOpenDeliverModel      *DeliverCardRequestTopOpenDeliverModel      `json:"topOpenDeliverModel,omitempty" xml:"topOpenDeliverModel,omitempty" type:"Struct"`
	UserIdType               *int32                                      `json:"userIdType,omitempty" xml:"userIdType,omitempty"`
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
	Extension  map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
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

func (s *DeliverCardRequestImGroupOpenDeliverModel) SetExtension(v map[string]*string) *DeliverCardRequestImGroupOpenDeliverModel {
	s.Extension = v
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
	Extension map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
	RobotCode *string            `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	SpaceType *string            `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
}

func (s DeliverCardRequestImRobotOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s DeliverCardRequestImRobotOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *DeliverCardRequestImRobotOpenDeliverModel) SetExtension(v map[string]*string) *DeliverCardRequestImRobotOpenDeliverModel {
	s.Extension = v
	return s
}

func (s *DeliverCardRequestImRobotOpenDeliverModel) SetRobotCode(v string) *DeliverCardRequestImRobotOpenDeliverModel {
	s.RobotCode = &v
	return s
}

func (s *DeliverCardRequestImRobotOpenDeliverModel) SetSpaceType(v string) *DeliverCardRequestImRobotOpenDeliverModel {
	s.SpaceType = &v
	return s
}

type DeliverCardRequestImSingleOpenDeliverModel struct {
	AtUserIds map[string]*string `json:"atUserIds,omitempty" xml:"atUserIds,omitempty"`
	Extension map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
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

func (s *DeliverCardRequestImSingleOpenDeliverModel) SetExtension(v map[string]*string) *DeliverCardRequestImSingleOpenDeliverModel {
	s.Extension = v
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
	CarrierId *string `json:"carrierId,omitempty" xml:"carrierId,omitempty"`
	ErrorMsg  *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
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

func (s *DeliverCardResponseBodyResult) SetCarrierId(v string) *DeliverCardResponseBodyResult {
	s.CarrierId = &v
	return s
}

func (s *DeliverCardResponseBodyResult) SetErrorMsg(v string) *DeliverCardResponseBodyResult {
	s.ErrorMsg = &v
	return s
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeliverCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DeliverCardWithDelegateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeliverCardWithDelegateHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeliverCardWithDelegateHeaders) GoString() string {
	return s.String()
}

func (s *DeliverCardWithDelegateHeaders) SetCommonHeaders(v map[string]*string) *DeliverCardWithDelegateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeliverCardWithDelegateHeaders) SetXAcsDingtalkAccessToken(v string) *DeliverCardWithDelegateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeliverCardWithDelegateRequest struct {
	CoFeedOpenDeliverModel   *DeliverCardWithDelegateRequestCoFeedOpenDeliverModel   `json:"coFeedOpenDeliverModel,omitempty" xml:"coFeedOpenDeliverModel,omitempty" type:"Struct"`
	DocOpenDeliverModel      *DeliverCardWithDelegateRequestDocOpenDeliverModel      `json:"docOpenDeliverModel,omitempty" xml:"docOpenDeliverModel,omitempty" type:"Struct"`
	ImGroupOpenDeliverModel  *DeliverCardWithDelegateRequestImGroupOpenDeliverModel  `json:"imGroupOpenDeliverModel,omitempty" xml:"imGroupOpenDeliverModel,omitempty" type:"Struct"`
	ImRobotOpenDeliverModel  *DeliverCardWithDelegateRequestImRobotOpenDeliverModel  `json:"imRobotOpenDeliverModel,omitempty" xml:"imRobotOpenDeliverModel,omitempty" type:"Struct"`
	ImSingleOpenDeliverModel *DeliverCardWithDelegateRequestImSingleOpenDeliverModel `json:"imSingleOpenDeliverModel,omitempty" xml:"imSingleOpenDeliverModel,omitempty" type:"Struct"`
	OpenSpaceId              *string                                                 `json:"openSpaceId,omitempty" xml:"openSpaceId,omitempty"`
	OutTrackId               *string                                                 `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	TopOpenDeliverModel      *DeliverCardWithDelegateRequestTopOpenDeliverModel      `json:"topOpenDeliverModel,omitempty" xml:"topOpenDeliverModel,omitempty" type:"Struct"`
	UserIdType               *int32                                                  `json:"userIdType,omitempty" xml:"userIdType,omitempty"`
}

func (s DeliverCardWithDelegateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeliverCardWithDelegateRequest) GoString() string {
	return s.String()
}

func (s *DeliverCardWithDelegateRequest) SetCoFeedOpenDeliverModel(v *DeliverCardWithDelegateRequestCoFeedOpenDeliverModel) *DeliverCardWithDelegateRequest {
	s.CoFeedOpenDeliverModel = v
	return s
}

func (s *DeliverCardWithDelegateRequest) SetDocOpenDeliverModel(v *DeliverCardWithDelegateRequestDocOpenDeliverModel) *DeliverCardWithDelegateRequest {
	s.DocOpenDeliverModel = v
	return s
}

func (s *DeliverCardWithDelegateRequest) SetImGroupOpenDeliverModel(v *DeliverCardWithDelegateRequestImGroupOpenDeliverModel) *DeliverCardWithDelegateRequest {
	s.ImGroupOpenDeliverModel = v
	return s
}

func (s *DeliverCardWithDelegateRequest) SetImRobotOpenDeliverModel(v *DeliverCardWithDelegateRequestImRobotOpenDeliverModel) *DeliverCardWithDelegateRequest {
	s.ImRobotOpenDeliverModel = v
	return s
}

func (s *DeliverCardWithDelegateRequest) SetImSingleOpenDeliverModel(v *DeliverCardWithDelegateRequestImSingleOpenDeliverModel) *DeliverCardWithDelegateRequest {
	s.ImSingleOpenDeliverModel = v
	return s
}

func (s *DeliverCardWithDelegateRequest) SetOpenSpaceId(v string) *DeliverCardWithDelegateRequest {
	s.OpenSpaceId = &v
	return s
}

func (s *DeliverCardWithDelegateRequest) SetOutTrackId(v string) *DeliverCardWithDelegateRequest {
	s.OutTrackId = &v
	return s
}

func (s *DeliverCardWithDelegateRequest) SetTopOpenDeliverModel(v *DeliverCardWithDelegateRequestTopOpenDeliverModel) *DeliverCardWithDelegateRequest {
	s.TopOpenDeliverModel = v
	return s
}

func (s *DeliverCardWithDelegateRequest) SetUserIdType(v int32) *DeliverCardWithDelegateRequest {
	s.UserIdType = &v
	return s
}

type DeliverCardWithDelegateRequestCoFeedOpenDeliverModel struct {
	BizTag      *string `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	GmtTimeLine *int64  `json:"gmtTimeLine,omitempty" xml:"gmtTimeLine,omitempty"`
}

func (s DeliverCardWithDelegateRequestCoFeedOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s DeliverCardWithDelegateRequestCoFeedOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *DeliverCardWithDelegateRequestCoFeedOpenDeliverModel) SetBizTag(v string) *DeliverCardWithDelegateRequestCoFeedOpenDeliverModel {
	s.BizTag = &v
	return s
}

func (s *DeliverCardWithDelegateRequestCoFeedOpenDeliverModel) SetGmtTimeLine(v int64) *DeliverCardWithDelegateRequestCoFeedOpenDeliverModel {
	s.GmtTimeLine = &v
	return s
}

type DeliverCardWithDelegateRequestDocOpenDeliverModel struct {
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeliverCardWithDelegateRequestDocOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s DeliverCardWithDelegateRequestDocOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *DeliverCardWithDelegateRequestDocOpenDeliverModel) SetUserId(v string) *DeliverCardWithDelegateRequestDocOpenDeliverModel {
	s.UserId = &v
	return s
}

type DeliverCardWithDelegateRequestImGroupOpenDeliverModel struct {
	AtUserIds  map[string]*string `json:"atUserIds,omitempty" xml:"atUserIds,omitempty"`
	Extension  map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
	Recipients []*string          `json:"recipients,omitempty" xml:"recipients,omitempty" type:"Repeated"`
	RobotCode  *string            `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s DeliverCardWithDelegateRequestImGroupOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s DeliverCardWithDelegateRequestImGroupOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *DeliverCardWithDelegateRequestImGroupOpenDeliverModel) SetAtUserIds(v map[string]*string) *DeliverCardWithDelegateRequestImGroupOpenDeliverModel {
	s.AtUserIds = v
	return s
}

func (s *DeliverCardWithDelegateRequestImGroupOpenDeliverModel) SetExtension(v map[string]*string) *DeliverCardWithDelegateRequestImGroupOpenDeliverModel {
	s.Extension = v
	return s
}

func (s *DeliverCardWithDelegateRequestImGroupOpenDeliverModel) SetRecipients(v []*string) *DeliverCardWithDelegateRequestImGroupOpenDeliverModel {
	s.Recipients = v
	return s
}

func (s *DeliverCardWithDelegateRequestImGroupOpenDeliverModel) SetRobotCode(v string) *DeliverCardWithDelegateRequestImGroupOpenDeliverModel {
	s.RobotCode = &v
	return s
}

type DeliverCardWithDelegateRequestImRobotOpenDeliverModel struct {
	Extension map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
	RobotCode *string            `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	SpaceType *string            `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
}

func (s DeliverCardWithDelegateRequestImRobotOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s DeliverCardWithDelegateRequestImRobotOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *DeliverCardWithDelegateRequestImRobotOpenDeliverModel) SetExtension(v map[string]*string) *DeliverCardWithDelegateRequestImRobotOpenDeliverModel {
	s.Extension = v
	return s
}

func (s *DeliverCardWithDelegateRequestImRobotOpenDeliverModel) SetRobotCode(v string) *DeliverCardWithDelegateRequestImRobotOpenDeliverModel {
	s.RobotCode = &v
	return s
}

func (s *DeliverCardWithDelegateRequestImRobotOpenDeliverModel) SetSpaceType(v string) *DeliverCardWithDelegateRequestImRobotOpenDeliverModel {
	s.SpaceType = &v
	return s
}

type DeliverCardWithDelegateRequestImSingleOpenDeliverModel struct {
	AtUserIds map[string]*string `json:"atUserIds,omitempty" xml:"atUserIds,omitempty"`
	Extension map[string]*string `json:"extension,omitempty" xml:"extension,omitempty"`
}

func (s DeliverCardWithDelegateRequestImSingleOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s DeliverCardWithDelegateRequestImSingleOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *DeliverCardWithDelegateRequestImSingleOpenDeliverModel) SetAtUserIds(v map[string]*string) *DeliverCardWithDelegateRequestImSingleOpenDeliverModel {
	s.AtUserIds = v
	return s
}

func (s *DeliverCardWithDelegateRequestImSingleOpenDeliverModel) SetExtension(v map[string]*string) *DeliverCardWithDelegateRequestImSingleOpenDeliverModel {
	s.Extension = v
	return s
}

type DeliverCardWithDelegateRequestTopOpenDeliverModel struct {
	ExpiredTimeMillis *int64    `json:"expiredTimeMillis,omitempty" xml:"expiredTimeMillis,omitempty"`
	Platforms         []*string `json:"platforms,omitempty" xml:"platforms,omitempty" type:"Repeated"`
	UserIds           []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s DeliverCardWithDelegateRequestTopOpenDeliverModel) String() string {
	return tea.Prettify(s)
}

func (s DeliverCardWithDelegateRequestTopOpenDeliverModel) GoString() string {
	return s.String()
}

func (s *DeliverCardWithDelegateRequestTopOpenDeliverModel) SetExpiredTimeMillis(v int64) *DeliverCardWithDelegateRequestTopOpenDeliverModel {
	s.ExpiredTimeMillis = &v
	return s
}

func (s *DeliverCardWithDelegateRequestTopOpenDeliverModel) SetPlatforms(v []*string) *DeliverCardWithDelegateRequestTopOpenDeliverModel {
	s.Platforms = v
	return s
}

func (s *DeliverCardWithDelegateRequestTopOpenDeliverModel) SetUserIds(v []*string) *DeliverCardWithDelegateRequestTopOpenDeliverModel {
	s.UserIds = v
	return s
}

type DeliverCardWithDelegateResponseBody struct {
	Result  []*DeliverCardWithDelegateResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success *bool                                        `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeliverCardWithDelegateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeliverCardWithDelegateResponseBody) GoString() string {
	return s.String()
}

func (s *DeliverCardWithDelegateResponseBody) SetResult(v []*DeliverCardWithDelegateResponseBodyResult) *DeliverCardWithDelegateResponseBody {
	s.Result = v
	return s
}

func (s *DeliverCardWithDelegateResponseBody) SetSuccess(v bool) *DeliverCardWithDelegateResponseBody {
	s.Success = &v
	return s
}

type DeliverCardWithDelegateResponseBodyResult struct {
	CarrierId *string `json:"carrierId,omitempty" xml:"carrierId,omitempty"`
	ErrorMsg  *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	SpaceId   *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeliverCardWithDelegateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeliverCardWithDelegateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeliverCardWithDelegateResponseBodyResult) SetCarrierId(v string) *DeliverCardWithDelegateResponseBodyResult {
	s.CarrierId = &v
	return s
}

func (s *DeliverCardWithDelegateResponseBodyResult) SetErrorMsg(v string) *DeliverCardWithDelegateResponseBodyResult {
	s.ErrorMsg = &v
	return s
}

func (s *DeliverCardWithDelegateResponseBodyResult) SetSpaceId(v string) *DeliverCardWithDelegateResponseBodyResult {
	s.SpaceId = &v
	return s
}

func (s *DeliverCardWithDelegateResponseBodyResult) SetSpaceType(v string) *DeliverCardWithDelegateResponseBodyResult {
	s.SpaceType = &v
	return s
}

func (s *DeliverCardWithDelegateResponseBodyResult) SetSuccess(v bool) *DeliverCardWithDelegateResponseBodyResult {
	s.Success = &v
	return s
}

type DeliverCardWithDelegateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeliverCardWithDelegateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeliverCardWithDelegateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeliverCardWithDelegateResponse) GoString() string {
	return s.String()
}

func (s *DeliverCardWithDelegateResponse) SetHeaders(v map[string]*string) *DeliverCardWithDelegateResponse {
	s.Headers = v
	return s
}

func (s *DeliverCardWithDelegateResponse) SetStatusCode(v int32) *DeliverCardWithDelegateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeliverCardWithDelegateResponse) SetBody(v *DeliverCardWithDelegateResponseBody) *DeliverCardWithDelegateResponse {
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type RegisterCallbackWithDelegateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RegisterCallbackWithDelegateHeaders) String() string {
	return tea.Prettify(s)
}

func (s RegisterCallbackWithDelegateHeaders) GoString() string {
	return s.String()
}

func (s *RegisterCallbackWithDelegateHeaders) SetCommonHeaders(v map[string]*string) *RegisterCallbackWithDelegateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RegisterCallbackWithDelegateHeaders) SetXAcsDingtalkAccessToken(v string) *RegisterCallbackWithDelegateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RegisterCallbackWithDelegateRequest struct {
	ApiSecret        *string `json:"apiSecret,omitempty" xml:"apiSecret,omitempty"`
	CallbackRouteKey *string `json:"callbackRouteKey,omitempty" xml:"callbackRouteKey,omitempty"`
	CallbackUrl      *string `json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`
	ForceUpdate      *bool   `json:"forceUpdate,omitempty" xml:"forceUpdate,omitempty"`
}

func (s RegisterCallbackWithDelegateRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterCallbackWithDelegateRequest) GoString() string {
	return s.String()
}

func (s *RegisterCallbackWithDelegateRequest) SetApiSecret(v string) *RegisterCallbackWithDelegateRequest {
	s.ApiSecret = &v
	return s
}

func (s *RegisterCallbackWithDelegateRequest) SetCallbackRouteKey(v string) *RegisterCallbackWithDelegateRequest {
	s.CallbackRouteKey = &v
	return s
}

func (s *RegisterCallbackWithDelegateRequest) SetCallbackUrl(v string) *RegisterCallbackWithDelegateRequest {
	s.CallbackUrl = &v
	return s
}

func (s *RegisterCallbackWithDelegateRequest) SetForceUpdate(v bool) *RegisterCallbackWithDelegateRequest {
	s.ForceUpdate = &v
	return s
}

type RegisterCallbackWithDelegateResponseBody struct {
	Result  *RegisterCallbackWithDelegateResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RegisterCallbackWithDelegateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterCallbackWithDelegateResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterCallbackWithDelegateResponseBody) SetResult(v *RegisterCallbackWithDelegateResponseBodyResult) *RegisterCallbackWithDelegateResponseBody {
	s.Result = v
	return s
}

func (s *RegisterCallbackWithDelegateResponseBody) SetSuccess(v bool) *RegisterCallbackWithDelegateResponseBody {
	s.Success = &v
	return s
}

type RegisterCallbackWithDelegateResponseBodyResult struct {
	ApiSecret   *string `json:"apiSecret,omitempty" xml:"apiSecret,omitempty"`
	CallbackUrl *string `json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`
}

func (s RegisterCallbackWithDelegateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s RegisterCallbackWithDelegateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RegisterCallbackWithDelegateResponseBodyResult) SetApiSecret(v string) *RegisterCallbackWithDelegateResponseBodyResult {
	s.ApiSecret = &v
	return s
}

func (s *RegisterCallbackWithDelegateResponseBodyResult) SetCallbackUrl(v string) *RegisterCallbackWithDelegateResponseBodyResult {
	s.CallbackUrl = &v
	return s
}

type RegisterCallbackWithDelegateResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterCallbackWithDelegateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterCallbackWithDelegateResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterCallbackWithDelegateResponse) GoString() string {
	return s.String()
}

func (s *RegisterCallbackWithDelegateResponse) SetHeaders(v map[string]*string) *RegisterCallbackWithDelegateResponse {
	s.Headers = v
	return s
}

func (s *RegisterCallbackWithDelegateResponse) SetStatusCode(v int32) *RegisterCallbackWithDelegateResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterCallbackWithDelegateResponse) SetBody(v *RegisterCallbackWithDelegateResponseBody) *RegisterCallbackWithDelegateResponse {
	s.Body = v
	return s
}

type StreamingUpdateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s StreamingUpdateHeaders) String() string {
	return tea.Prettify(s)
}

func (s StreamingUpdateHeaders) GoString() string {
	return s.String()
}

func (s *StreamingUpdateHeaders) SetCommonHeaders(v map[string]*string) *StreamingUpdateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StreamingUpdateHeaders) SetXAcsDingtalkAccessToken(v string) *StreamingUpdateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type StreamingUpdateRequest struct {
	Content    *string `json:"content,omitempty" xml:"content,omitempty"`
	Guid       *string `json:"guid,omitempty" xml:"guid,omitempty"`
	IsError    *bool   `json:"isError,omitempty" xml:"isError,omitempty"`
	IsFinalize *bool   `json:"isFinalize,omitempty" xml:"isFinalize,omitempty"`
	IsFull     *bool   `json:"isFull,omitempty" xml:"isFull,omitempty"`
	Key        *string `json:"key,omitempty" xml:"key,omitempty"`
	OutTrackId *string `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
}

func (s StreamingUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s StreamingUpdateRequest) GoString() string {
	return s.String()
}

func (s *StreamingUpdateRequest) SetContent(v string) *StreamingUpdateRequest {
	s.Content = &v
	return s
}

func (s *StreamingUpdateRequest) SetGuid(v string) *StreamingUpdateRequest {
	s.Guid = &v
	return s
}

func (s *StreamingUpdateRequest) SetIsError(v bool) *StreamingUpdateRequest {
	s.IsError = &v
	return s
}

func (s *StreamingUpdateRequest) SetIsFinalize(v bool) *StreamingUpdateRequest {
	s.IsFinalize = &v
	return s
}

func (s *StreamingUpdateRequest) SetIsFull(v bool) *StreamingUpdateRequest {
	s.IsFull = &v
	return s
}

func (s *StreamingUpdateRequest) SetKey(v string) *StreamingUpdateRequest {
	s.Key = &v
	return s
}

func (s *StreamingUpdateRequest) SetOutTrackId(v string) *StreamingUpdateRequest {
	s.OutTrackId = &v
	return s
}

type StreamingUpdateResponseBody struct {
	Result  *bool `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StreamingUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StreamingUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *StreamingUpdateResponseBody) SetResult(v bool) *StreamingUpdateResponseBody {
	s.Result = &v
	return s
}

func (s *StreamingUpdateResponseBody) SetSuccess(v bool) *StreamingUpdateResponseBody {
	s.Success = &v
	return s
}

type StreamingUpdateResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StreamingUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StreamingUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s StreamingUpdateResponse) GoString() string {
	return s.String()
}

func (s *StreamingUpdateResponse) SetHeaders(v map[string]*string) *StreamingUpdateResponse {
	s.Headers = v
	return s
}

func (s *StreamingUpdateResponse) SetStatusCode(v int32) *StreamingUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *StreamingUpdateResponse) SetBody(v *StreamingUpdateResponseBody) *StreamingUpdateResponse {
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type UpdateCardWithDelegateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateCardWithDelegateHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateCardWithDelegateHeaders) GoString() string {
	return s.String()
}

func (s *UpdateCardWithDelegateHeaders) SetCommonHeaders(v map[string]*string) *UpdateCardWithDelegateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateCardWithDelegateHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateCardWithDelegateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateCardWithDelegateRequest struct {
	CardData          *UpdateCardWithDelegateRequestCardData          `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	CardUpdateOptions *UpdateCardWithDelegateRequestCardUpdateOptions `json:"cardUpdateOptions,omitempty" xml:"cardUpdateOptions,omitempty" type:"Struct"`
	OutTrackId        *string                                         `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	PrivateData       map[string]*PrivateDataValue                    `json:"privateData,omitempty" xml:"privateData,omitempty"`
	UserIdType        *int32                                          `json:"userIdType,omitempty" xml:"userIdType,omitempty"`
}

func (s UpdateCardWithDelegateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCardWithDelegateRequest) GoString() string {
	return s.String()
}

func (s *UpdateCardWithDelegateRequest) SetCardData(v *UpdateCardWithDelegateRequestCardData) *UpdateCardWithDelegateRequest {
	s.CardData = v
	return s
}

func (s *UpdateCardWithDelegateRequest) SetCardUpdateOptions(v *UpdateCardWithDelegateRequestCardUpdateOptions) *UpdateCardWithDelegateRequest {
	s.CardUpdateOptions = v
	return s
}

func (s *UpdateCardWithDelegateRequest) SetOutTrackId(v string) *UpdateCardWithDelegateRequest {
	s.OutTrackId = &v
	return s
}

func (s *UpdateCardWithDelegateRequest) SetPrivateData(v map[string]*PrivateDataValue) *UpdateCardWithDelegateRequest {
	s.PrivateData = v
	return s
}

func (s *UpdateCardWithDelegateRequest) SetUserIdType(v int32) *UpdateCardWithDelegateRequest {
	s.UserIdType = &v
	return s
}

type UpdateCardWithDelegateRequestCardData struct {
	CardParamMap map[string]*string `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
}

func (s UpdateCardWithDelegateRequestCardData) String() string {
	return tea.Prettify(s)
}

func (s UpdateCardWithDelegateRequestCardData) GoString() string {
	return s.String()
}

func (s *UpdateCardWithDelegateRequestCardData) SetCardParamMap(v map[string]*string) *UpdateCardWithDelegateRequestCardData {
	s.CardParamMap = v
	return s
}

type UpdateCardWithDelegateRequestCardUpdateOptions struct {
	UpdateCardDataByKey    *bool `json:"updateCardDataByKey,omitempty" xml:"updateCardDataByKey,omitempty"`
	UpdatePrivateDataByKey *bool `json:"updatePrivateDataByKey,omitempty" xml:"updatePrivateDataByKey,omitempty"`
}

func (s UpdateCardWithDelegateRequestCardUpdateOptions) String() string {
	return tea.Prettify(s)
}

func (s UpdateCardWithDelegateRequestCardUpdateOptions) GoString() string {
	return s.String()
}

func (s *UpdateCardWithDelegateRequestCardUpdateOptions) SetUpdateCardDataByKey(v bool) *UpdateCardWithDelegateRequestCardUpdateOptions {
	s.UpdateCardDataByKey = &v
	return s
}

func (s *UpdateCardWithDelegateRequestCardUpdateOptions) SetUpdatePrivateDataByKey(v bool) *UpdateCardWithDelegateRequestCardUpdateOptions {
	s.UpdatePrivateDataByKey = &v
	return s
}

type UpdateCardWithDelegateResponseBody struct {
	Result  *bool `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateCardWithDelegateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCardWithDelegateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCardWithDelegateResponseBody) SetResult(v bool) *UpdateCardWithDelegateResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateCardWithDelegateResponseBody) SetSuccess(v bool) *UpdateCardWithDelegateResponseBody {
	s.Success = &v
	return s
}

type UpdateCardWithDelegateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCardWithDelegateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCardWithDelegateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCardWithDelegateResponse) GoString() string {
	return s.String()
}

func (s *UpdateCardWithDelegateResponse) SetHeaders(v map[string]*string) *UpdateCardWithDelegateResponse {
	s.Headers = v
	return s
}

func (s *UpdateCardWithDelegateResponse) SetStatusCode(v int32) *UpdateCardWithDelegateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCardWithDelegateResponse) SetBody(v *UpdateCardWithDelegateResponseBody) *UpdateCardWithDelegateResponse {
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

func (client *Client) AppendSpaceWithDelegateWithOptions(request *AppendSpaceWithDelegateRequest, headers *AppendSpaceWithDelegateHeaders, runtime *util.RuntimeOptions) (_result *AppendSpaceWithDelegateResponse, _err error) {
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
		Action:      tea.String("AppendSpaceWithDelegate"),
		Version:     tea.String("card_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/card/me/instances/spaces"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AppendSpaceWithDelegateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AppendSpaceWithDelegate(request *AppendSpaceWithDelegateRequest) (_result *AppendSpaceWithDelegateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AppendSpaceWithDelegateHeaders{}
	_result = &AppendSpaceWithDelegateResponse{}
	_body, _err := client.AppendSpaceWithDelegateWithOptions(request, headers, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.ImSingleOpenDeliverModel)) {
		body["imSingleOpenDeliverModel"] = request.ImSingleOpenDeliverModel
	}

	if !tea.BoolValue(util.IsUnset(request.ImSingleOpenSpaceModel)) {
		body["imSingleOpenSpaceModel"] = request.ImSingleOpenSpaceModel
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

func (client *Client) CreateAndDeliverWithDelegateWithOptions(request *CreateAndDeliverWithDelegateRequest, headers *CreateAndDeliverWithDelegateHeaders, runtime *util.RuntimeOptions) (_result *CreateAndDeliverWithDelegateResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ImSingleOpenDeliverModel)) {
		body["imSingleOpenDeliverModel"] = request.ImSingleOpenDeliverModel
	}

	if !tea.BoolValue(util.IsUnset(request.ImSingleOpenSpaceModel)) {
		body["imSingleOpenSpaceModel"] = request.ImSingleOpenSpaceModel
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
		Action:      tea.String("CreateAndDeliverWithDelegate"),
		Version:     tea.String("card_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/card/me/instances/createAndDeliver"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAndDeliverWithDelegateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAndDeliverWithDelegate(request *CreateAndDeliverWithDelegateRequest) (_result *CreateAndDeliverWithDelegateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateAndDeliverWithDelegateHeaders{}
	_result = &CreateAndDeliverWithDelegateResponse{}
	_body, _err := client.CreateAndDeliverWithDelegateWithOptions(request, headers, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.ImSingleOpenSpaceModel)) {
		body["imSingleOpenSpaceModel"] = request.ImSingleOpenSpaceModel
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

func (client *Client) CreateCardWithDelegateWithOptions(request *CreateCardWithDelegateRequest, headers *CreateCardWithDelegateHeaders, runtime *util.RuntimeOptions) (_result *CreateCardWithDelegateResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ImSingleOpenSpaceModel)) {
		body["imSingleOpenSpaceModel"] = request.ImSingleOpenSpaceModel
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
		Action:      tea.String("CreateCardWithDelegate"),
		Version:     tea.String("card_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/card/me/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCardWithDelegateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCardWithDelegate(request *CreateCardWithDelegateRequest) (_result *CreateCardWithDelegateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateCardWithDelegateHeaders{}
	_result = &CreateCardWithDelegateResponse{}
	_body, _err := client.CreateCardWithDelegateWithOptions(request, headers, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.ImSingleOpenDeliverModel)) {
		body["imSingleOpenDeliverModel"] = request.ImSingleOpenDeliverModel
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

func (client *Client) DeliverCardWithDelegateWithOptions(request *DeliverCardWithDelegateRequest, headers *DeliverCardWithDelegateHeaders, runtime *util.RuntimeOptions) (_result *DeliverCardWithDelegateResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ImSingleOpenDeliverModel)) {
		body["imSingleOpenDeliverModel"] = request.ImSingleOpenDeliverModel
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
		Action:      tea.String("DeliverCardWithDelegate"),
		Version:     tea.String("card_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/card/me/instances/deliver"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeliverCardWithDelegateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeliverCardWithDelegate(request *DeliverCardWithDelegateRequest) (_result *DeliverCardWithDelegateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeliverCardWithDelegateHeaders{}
	_result = &DeliverCardWithDelegateResponse{}
	_body, _err := client.DeliverCardWithDelegateWithOptions(request, headers, runtime)
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

func (client *Client) RegisterCallbackWithDelegateWithOptions(request *RegisterCallbackWithDelegateRequest, headers *RegisterCallbackWithDelegateHeaders, runtime *util.RuntimeOptions) (_result *RegisterCallbackWithDelegateResponse, _err error) {
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
		Action:      tea.String("RegisterCallbackWithDelegate"),
		Version:     tea.String("card_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/card/me/callbacks/register"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RegisterCallbackWithDelegateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterCallbackWithDelegate(request *RegisterCallbackWithDelegateRequest) (_result *RegisterCallbackWithDelegateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RegisterCallbackWithDelegateHeaders{}
	_result = &RegisterCallbackWithDelegateResponse{}
	_body, _err := client.RegisterCallbackWithDelegateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StreamingUpdateWithOptions(request *StreamingUpdateRequest, headers *StreamingUpdateHeaders, runtime *util.RuntimeOptions) (_result *StreamingUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Guid)) {
		body["guid"] = request.Guid
	}

	if !tea.BoolValue(util.IsUnset(request.IsError)) {
		body["isError"] = request.IsError
	}

	if !tea.BoolValue(util.IsUnset(request.IsFinalize)) {
		body["isFinalize"] = request.IsFinalize
	}

	if !tea.BoolValue(util.IsUnset(request.IsFull)) {
		body["isFull"] = request.IsFull
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		body["key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.OutTrackId)) {
		body["outTrackId"] = request.OutTrackId
	}

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
		Action:      tea.String("StreamingUpdate"),
		Version:     tea.String("card_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/card/streaming"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &StreamingUpdateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StreamingUpdate(request *StreamingUpdateRequest) (_result *StreamingUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StreamingUpdateHeaders{}
	_result = &StreamingUpdateResponse{}
	_body, _err := client.StreamingUpdateWithOptions(request, headers, runtime)
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

func (client *Client) UpdateCardWithDelegateWithOptions(request *UpdateCardWithDelegateRequest, headers *UpdateCardWithDelegateHeaders, runtime *util.RuntimeOptions) (_result *UpdateCardWithDelegateResponse, _err error) {
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
		Action:      tea.String("UpdateCardWithDelegate"),
		Version:     tea.String("card_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/card/me/instances"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCardWithDelegateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCardWithDelegate(request *UpdateCardWithDelegateRequest) (_result *UpdateCardWithDelegateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateCardWithDelegateHeaders{}
	_result = &UpdateCardWithDelegateResponse{}
	_body, _err := client.UpdateCardWithDelegateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
