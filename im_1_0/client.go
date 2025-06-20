// This file is auto-generated, don't edit it. Thanks.
package im_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type UserListValue struct {
	JoinTime   *int64 `json:"joinTime,omitempty" xml:"joinTime,omitempty"`
	ModifyTime *int64 `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	Mute       *bool  `json:"mute,omitempty" xml:"mute,omitempty"`
	TopRank    *bool  `json:"topRank,omitempty" xml:"topRank,omitempty"`
	Visible    *bool  `json:"visible,omitempty" xml:"visible,omitempty"`
}

func (s UserListValue) String() string {
	return tea.Prettify(s)
}

func (s UserListValue) GoString() string {
	return s.String()
}

func (s *UserListValue) SetJoinTime(v int64) *UserListValue {
	s.JoinTime = &v
	return s
}

func (s *UserListValue) SetModifyTime(v int64) *UserListValue {
	s.ModifyTime = &v
	return s
}

func (s *UserListValue) SetMute(v bool) *UserListValue {
	s.Mute = &v
	return s
}

func (s *UserListValue) SetTopRank(v bool) *UserListValue {
	s.TopRank = &v
	return s
}

func (s *UserListValue) SetVisible(v bool) *UserListValue {
	s.Visible = &v
	return s
}

type PrivateDataValue struct {
	CardParamMap        map[string]*string `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
	CardMediaIdParamMap map[string]*string `json:"cardMediaIdParamMap,omitempty" xml:"cardMediaIdParamMap,omitempty"`
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

func (s *PrivateDataValue) SetCardMediaIdParamMap(v map[string]*string) *PrivateDataValue {
	s.CardMediaIdParamMap = v
	return s
}

type AddOrgTextEmotionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddOrgTextEmotionHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddOrgTextEmotionHeaders) GoString() string {
	return s.String()
}

func (s *AddOrgTextEmotionHeaders) SetCommonHeaders(v map[string]*string) *AddOrgTextEmotionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddOrgTextEmotionHeaders) SetXAcsDingtalkAccessToken(v string) *AddOrgTextEmotionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddOrgTextEmotionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// @123xxx
	BackgroundMediaId *string `json:"backgroundMediaId,omitempty" xml:"backgroundMediaId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// @345xxx
	BackgroundMediaIdForPanel *string `json:"backgroundMediaIdForPanel,omitempty" xml:"backgroundMediaIdForPanel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// -1
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 企业表情1
	EmotionName *string `json:"emotionName,omitempty" xml:"emotionName,omitempty"`
}

func (s AddOrgTextEmotionRequest) String() string {
	return tea.Prettify(s)
}

func (s AddOrgTextEmotionRequest) GoString() string {
	return s.String()
}

func (s *AddOrgTextEmotionRequest) SetBackgroundMediaId(v string) *AddOrgTextEmotionRequest {
	s.BackgroundMediaId = &v
	return s
}

func (s *AddOrgTextEmotionRequest) SetBackgroundMediaIdForPanel(v string) *AddOrgTextEmotionRequest {
	s.BackgroundMediaIdForPanel = &v
	return s
}

func (s *AddOrgTextEmotionRequest) SetDeptId(v int64) *AddOrgTextEmotionRequest {
	s.DeptId = &v
	return s
}

func (s *AddOrgTextEmotionRequest) SetEmotionName(v string) *AddOrgTextEmotionRequest {
	s.EmotionName = &v
	return s
}

type AddOrgTextEmotionResponseBody struct {
	Result  *AddOrgTextEmotionResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddOrgTextEmotionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddOrgTextEmotionResponseBody) GoString() string {
	return s.String()
}

func (s *AddOrgTextEmotionResponseBody) SetResult(v *AddOrgTextEmotionResponseBodyResult) *AddOrgTextEmotionResponseBody {
	s.Result = v
	return s
}

func (s *AddOrgTextEmotionResponseBody) SetSuccess(v bool) *AddOrgTextEmotionResponseBody {
	s.Success = &v
	return s
}

type AddOrgTextEmotionResponseBodyResult struct {
	// example:
	//
	// corp_123456
	EmotionId *string `json:"emotionId,omitempty" xml:"emotionId,omitempty"`
}

func (s AddOrgTextEmotionResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s AddOrgTextEmotionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddOrgTextEmotionResponseBodyResult) SetEmotionId(v string) *AddOrgTextEmotionResponseBodyResult {
	s.EmotionId = &v
	return s
}

type AddOrgTextEmotionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddOrgTextEmotionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddOrgTextEmotionResponse) String() string {
	return tea.Prettify(s)
}

func (s AddOrgTextEmotionResponse) GoString() string {
	return s.String()
}

func (s *AddOrgTextEmotionResponse) SetHeaders(v map[string]*string) *AddOrgTextEmotionResponse {
	s.Headers = v
	return s
}

func (s *AddOrgTextEmotionResponse) SetStatusCode(v int32) *AddOrgTextEmotionResponse {
	s.StatusCode = &v
	return s
}

func (s *AddOrgTextEmotionResponse) SetBody(v *AddOrgTextEmotionResponseBody) *AddOrgTextEmotionResponse {
	s.Body = v
	return s
}

type AddRobotToConversationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddRobotToConversationHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddRobotToConversationHeaders) GoString() string {
	return s.String()
}

func (s *AddRobotToConversationHeaders) SetCommonHeaders(v map[string]*string) *AddRobotToConversationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddRobotToConversationHeaders) SetXAcsDingtalkAccessToken(v string) *AddRobotToConversationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddRobotToConversationRequest struct {
	// example:
	//
	// @lALPDe7s26Bre
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// 小加
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cid123cd
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s AddRobotToConversationRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRobotToConversationRequest) GoString() string {
	return s.String()
}

func (s *AddRobotToConversationRequest) SetIcon(v string) *AddRobotToConversationRequest {
	s.Icon = &v
	return s
}

func (s *AddRobotToConversationRequest) SetName(v string) *AddRobotToConversationRequest {
	s.Name = &v
	return s
}

func (s *AddRobotToConversationRequest) SetOpenConversationId(v string) *AddRobotToConversationRequest {
	s.OpenConversationId = &v
	return s
}

func (s *AddRobotToConversationRequest) SetRobotCode(v string) *AddRobotToConversationRequest {
	s.RobotCode = &v
	return s
}

type AddRobotToConversationResponseBody struct {
	ChatBotUserId *string `json:"chatBotUserId,omitempty" xml:"chatBotUserId,omitempty"`
}

func (s AddRobotToConversationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddRobotToConversationResponseBody) GoString() string {
	return s.String()
}

func (s *AddRobotToConversationResponseBody) SetChatBotUserId(v string) *AddRobotToConversationResponseBody {
	s.ChatBotUserId = &v
	return s
}

type AddRobotToConversationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRobotToConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRobotToConversationResponse) String() string {
	return tea.Prettify(s)
}

func (s AddRobotToConversationResponse) GoString() string {
	return s.String()
}

func (s *AddRobotToConversationResponse) SetHeaders(v map[string]*string) *AddRobotToConversationResponse {
	s.Headers = v
	return s
}

func (s *AddRobotToConversationResponse) SetStatusCode(v int32) *AddRobotToConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRobotToConversationResponse) SetBody(v *AddRobotToConversationResponseBody) *AddRobotToConversationResponse {
	s.Body = v
	return s
}

type AddUnfurlingRegisterHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddUnfurlingRegisterHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddUnfurlingRegisterHeaders) GoString() string {
	return s.String()
}

func (s *AddUnfurlingRegisterHeaders) SetCommonHeaders(v map[string]*string) *AddUnfurlingRegisterHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddUnfurlingRegisterHeaders) SetXAcsDingtalkAccessToken(v string) *AddUnfurlingRegisterHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddUnfurlingRegisterRequest struct {
	// This parameter is required.
	ApiSecret *string `json:"apiSecret,omitempty" xml:"apiSecret,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3102xxxxxxx
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://xxx.xxx.com/api/dingtalk/link_unfurling
	CallbackUrl *string `json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// d7b9xxx-xxx-xxxx-xxxx-xxxxxxx.schema
	CardTemplateId *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// www.dingtalk.com
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// 规则描述
	RuleDesc *string `json:"ruleDesc,omitempty" xml:"ruleDesc,omitempty"`
	// This parameter is required.
	RuleMatchType *int32 `json:"ruleMatchType,omitempty" xml:"ruleMatchType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 37xxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AddUnfurlingRegisterRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUnfurlingRegisterRequest) GoString() string {
	return s.String()
}

func (s *AddUnfurlingRegisterRequest) SetApiSecret(v string) *AddUnfurlingRegisterRequest {
	s.ApiSecret = &v
	return s
}

func (s *AddUnfurlingRegisterRequest) SetAppId(v string) *AddUnfurlingRegisterRequest {
	s.AppId = &v
	return s
}

func (s *AddUnfurlingRegisterRequest) SetCallbackUrl(v string) *AddUnfurlingRegisterRequest {
	s.CallbackUrl = &v
	return s
}

func (s *AddUnfurlingRegisterRequest) SetCardTemplateId(v string) *AddUnfurlingRegisterRequest {
	s.CardTemplateId = &v
	return s
}

func (s *AddUnfurlingRegisterRequest) SetDomain(v string) *AddUnfurlingRegisterRequest {
	s.Domain = &v
	return s
}

func (s *AddUnfurlingRegisterRequest) SetPath(v string) *AddUnfurlingRegisterRequest {
	s.Path = &v
	return s
}

func (s *AddUnfurlingRegisterRequest) SetRuleDesc(v string) *AddUnfurlingRegisterRequest {
	s.RuleDesc = &v
	return s
}

func (s *AddUnfurlingRegisterRequest) SetRuleMatchType(v int32) *AddUnfurlingRegisterRequest {
	s.RuleMatchType = &v
	return s
}

func (s *AddUnfurlingRegisterRequest) SetUserId(v string) *AddUnfurlingRegisterRequest {
	s.UserId = &v
	return s
}

type AddUnfurlingRegisterResponseBody struct {
	// example:
	//
	// 1
	Id      *int64 `json:"id,omitempty" xml:"id,omitempty"`
	Success *bool  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddUnfurlingRegisterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUnfurlingRegisterResponseBody) GoString() string {
	return s.String()
}

func (s *AddUnfurlingRegisterResponseBody) SetId(v int64) *AddUnfurlingRegisterResponseBody {
	s.Id = &v
	return s
}

func (s *AddUnfurlingRegisterResponseBody) SetSuccess(v bool) *AddUnfurlingRegisterResponseBody {
	s.Success = &v
	return s
}

type AddUnfurlingRegisterResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUnfurlingRegisterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUnfurlingRegisterResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUnfurlingRegisterResponse) GoString() string {
	return s.String()
}

func (s *AddUnfurlingRegisterResponse) SetHeaders(v map[string]*string) *AddUnfurlingRegisterResponse {
	s.Headers = v
	return s
}

func (s *AddUnfurlingRegisterResponse) SetStatusCode(v int32) *AddUnfurlingRegisterResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUnfurlingRegisterResponse) SetBody(v *AddUnfurlingRegisterResponseBody) *AddUnfurlingRegisterResponse {
	s.Body = v
	return s
}

type AutoOpenDingTalkConnectHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AutoOpenDingTalkConnectHeaders) String() string {
	return tea.Prettify(s)
}

func (s AutoOpenDingTalkConnectHeaders) GoString() string {
	return s.String()
}

func (s *AutoOpenDingTalkConnectHeaders) SetCommonHeaders(v map[string]*string) *AutoOpenDingTalkConnectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AutoOpenDingTalkConnectHeaders) SetXAcsDingtalkAccessToken(v string) *AutoOpenDingTalkConnectHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AutoOpenDingTalkConnectResponseBody struct {
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s AutoOpenDingTalkConnectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AutoOpenDingTalkConnectResponseBody) GoString() string {
	return s.String()
}

func (s *AutoOpenDingTalkConnectResponseBody) SetMessage(v string) *AutoOpenDingTalkConnectResponseBody {
	s.Message = &v
	return s
}

type AutoOpenDingTalkConnectResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AutoOpenDingTalkConnectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AutoOpenDingTalkConnectResponse) String() string {
	return tea.Prettify(s)
}

func (s AutoOpenDingTalkConnectResponse) GoString() string {
	return s.String()
}

func (s *AutoOpenDingTalkConnectResponse) SetHeaders(v map[string]*string) *AutoOpenDingTalkConnectResponse {
	s.Headers = v
	return s
}

func (s *AutoOpenDingTalkConnectResponse) SetStatusCode(v int32) *AutoOpenDingTalkConnectResponse {
	s.StatusCode = &v
	return s
}

func (s *AutoOpenDingTalkConnectResponse) SetBody(v *AutoOpenDingTalkConnectResponseBody) *AutoOpenDingTalkConnectResponse {
	s.Body = v
	return s
}

type BatchQueryFamilySchoolMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchQueryFamilySchoolMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryFamilySchoolMessageHeaders) GoString() string {
	return s.String()
}

func (s *BatchQueryFamilySchoolMessageHeaders) SetCommonHeaders(v map[string]*string) *BatchQueryFamilySchoolMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchQueryFamilySchoolMessageHeaders) SetXAcsDingtalkAccessToken(v string) *BatchQueryFamilySchoolMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchQueryFamilySchoolMessageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cidxxxx
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	OpenMessageIds []*string `json:"openMessageIds,omitempty" xml:"openMessageIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s BatchQueryFamilySchoolMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryFamilySchoolMessageRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryFamilySchoolMessageRequest) SetOpenConversationId(v string) *BatchQueryFamilySchoolMessageRequest {
	s.OpenConversationId = &v
	return s
}

func (s *BatchQueryFamilySchoolMessageRequest) SetOpenMessageIds(v []*string) *BatchQueryFamilySchoolMessageRequest {
	s.OpenMessageIds = v
	return s
}

func (s *BatchQueryFamilySchoolMessageRequest) SetUnionId(v string) *BatchQueryFamilySchoolMessageRequest {
	s.UnionId = &v
	return s
}

type BatchQueryFamilySchoolMessageResponseBody struct {
	Messages []*BatchQueryFamilySchoolMessageResponseBodyMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
}

func (s BatchQueryFamilySchoolMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryFamilySchoolMessageResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQueryFamilySchoolMessageResponseBody) SetMessages(v []*BatchQueryFamilySchoolMessageResponseBodyMessages) *BatchQueryFamilySchoolMessageResponseBody {
	s.Messages = v
	return s
}

type BatchQueryFamilySchoolMessageResponseBodyMessages struct {
	ContentType *int32                                                          `json:"contentType,omitempty" xml:"contentType,omitempty"`
	CreateAt    *int64                                                          `json:"createAt,omitempty" xml:"createAt,omitempty"`
	MediaModels []*BatchQueryFamilySchoolMessageResponseBodyMessagesMediaModels `json:"mediaModels,omitempty" xml:"mediaModels,omitempty" type:"Repeated"`
	// example:
	//
	// msgxxx
	OpenMsgId *string `json:"openMsgId,omitempty" xml:"openMsgId,omitempty"`
}

func (s BatchQueryFamilySchoolMessageResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryFamilySchoolMessageResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *BatchQueryFamilySchoolMessageResponseBodyMessages) SetContentType(v int32) *BatchQueryFamilySchoolMessageResponseBodyMessages {
	s.ContentType = &v
	return s
}

func (s *BatchQueryFamilySchoolMessageResponseBodyMessages) SetCreateAt(v int64) *BatchQueryFamilySchoolMessageResponseBodyMessages {
	s.CreateAt = &v
	return s
}

func (s *BatchQueryFamilySchoolMessageResponseBodyMessages) SetMediaModels(v []*BatchQueryFamilySchoolMessageResponseBodyMessagesMediaModels) *BatchQueryFamilySchoolMessageResponseBodyMessages {
	s.MediaModels = v
	return s
}

func (s *BatchQueryFamilySchoolMessageResponseBodyMessages) SetOpenMsgId(v string) *BatchQueryFamilySchoolMessageResponseBodyMessages {
	s.OpenMsgId = &v
	return s
}

type BatchQueryFamilySchoolMessageResponseBodyMessagesMediaModels struct {
	// example:
	//
	// aa.png
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// png
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// @12xxx34
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	// example:
	//
	// 1234
	Size *string `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// https://wukong-xxxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// example:
	//
	// @12xx34
	VideoPicMediaId *string `json:"videoPicMediaId,omitempty" xml:"videoPicMediaId,omitempty"`
}

func (s BatchQueryFamilySchoolMessageResponseBodyMessagesMediaModels) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryFamilySchoolMessageResponseBodyMessagesMediaModels) GoString() string {
	return s.String()
}

func (s *BatchQueryFamilySchoolMessageResponseBodyMessagesMediaModels) SetFileName(v string) *BatchQueryFamilySchoolMessageResponseBodyMessagesMediaModels {
	s.FileName = &v
	return s
}

func (s *BatchQueryFamilySchoolMessageResponseBodyMessagesMediaModels) SetFileType(v string) *BatchQueryFamilySchoolMessageResponseBodyMessagesMediaModels {
	s.FileType = &v
	return s
}

func (s *BatchQueryFamilySchoolMessageResponseBodyMessagesMediaModels) SetMediaId(v string) *BatchQueryFamilySchoolMessageResponseBodyMessagesMediaModels {
	s.MediaId = &v
	return s
}

func (s *BatchQueryFamilySchoolMessageResponseBodyMessagesMediaModels) SetSize(v string) *BatchQueryFamilySchoolMessageResponseBodyMessagesMediaModels {
	s.Size = &v
	return s
}

func (s *BatchQueryFamilySchoolMessageResponseBodyMessagesMediaModels) SetUrl(v string) *BatchQueryFamilySchoolMessageResponseBodyMessagesMediaModels {
	s.Url = &v
	return s
}

func (s *BatchQueryFamilySchoolMessageResponseBodyMessagesMediaModels) SetVideoPicMediaId(v string) *BatchQueryFamilySchoolMessageResponseBodyMessagesMediaModels {
	s.VideoPicMediaId = &v
	return s
}

type BatchQueryFamilySchoolMessageResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchQueryFamilySchoolMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchQueryFamilySchoolMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryFamilySchoolMessageResponse) GoString() string {
	return s.String()
}

func (s *BatchQueryFamilySchoolMessageResponse) SetHeaders(v map[string]*string) *BatchQueryFamilySchoolMessageResponse {
	s.Headers = v
	return s
}

func (s *BatchQueryFamilySchoolMessageResponse) SetStatusCode(v int32) *BatchQueryFamilySchoolMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchQueryFamilySchoolMessageResponse) SetBody(v *BatchQueryFamilySchoolMessageResponseBody) *BatchQueryFamilySchoolMessageResponse {
	s.Body = v
	return s
}

type BatchQueryGroupMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchQueryGroupMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryGroupMemberHeaders) GoString() string {
	return s.String()
}

func (s *BatchQueryGroupMemberHeaders) SetCommonHeaders(v map[string]*string) *BatchQueryGroupMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchQueryGroupMemberHeaders) SetXAcsDingtalkAccessToken(v string) *BatchQueryGroupMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchQueryGroupMemberRequest struct {
	CoolAppCode *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 200
	MaxResults *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cidXXXXXXX
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s BatchQueryGroupMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryGroupMemberRequest) SetCoolAppCode(v string) *BatchQueryGroupMemberRequest {
	s.CoolAppCode = &v
	return s
}

func (s *BatchQueryGroupMemberRequest) SetMaxResults(v int64) *BatchQueryGroupMemberRequest {
	s.MaxResults = &v
	return s
}

func (s *BatchQueryGroupMemberRequest) SetNextToken(v string) *BatchQueryGroupMemberRequest {
	s.NextToken = &v
	return s
}

func (s *BatchQueryGroupMemberRequest) SetOpenConversationId(v string) *BatchQueryGroupMemberRequest {
	s.OpenConversationId = &v
	return s
}

type BatchQueryGroupMemberResponseBody struct {
	// example:
	//
	// false
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// cidXXXXXXXXX==
	MemberUserIds []*string `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// 92233720368
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchQueryGroupMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *BatchQueryGroupMemberResponseBody) SetHasMore(v bool) *BatchQueryGroupMemberResponseBody {
	s.HasMore = &v
	return s
}

func (s *BatchQueryGroupMemberResponseBody) SetMemberUserIds(v []*string) *BatchQueryGroupMemberResponseBody {
	s.MemberUserIds = v
	return s
}

func (s *BatchQueryGroupMemberResponseBody) SetNextToken(v string) *BatchQueryGroupMemberResponseBody {
	s.NextToken = &v
	return s
}

func (s *BatchQueryGroupMemberResponseBody) SetSuccess(v bool) *BatchQueryGroupMemberResponseBody {
	s.Success = &v
	return s
}

type BatchQueryGroupMemberResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchQueryGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchQueryGroupMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchQueryGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *BatchQueryGroupMemberResponse) SetHeaders(v map[string]*string) *BatchQueryGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *BatchQueryGroupMemberResponse) SetStatusCode(v int32) *BatchQueryGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchQueryGroupMemberResponse) SetBody(v *BatchQueryGroupMemberResponseBody) *BatchQueryGroupMemberResponse {
	s.Body = v
	return s
}

type CardTemplateBuildActionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CardTemplateBuildActionHeaders) String() string {
	return tea.Prettify(s)
}

func (s CardTemplateBuildActionHeaders) GoString() string {
	return s.String()
}

func (s *CardTemplateBuildActionHeaders) SetCommonHeaders(v map[string]*string) *CardTemplateBuildActionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CardTemplateBuildActionHeaders) SetXAcsDingtalkAccessToken(v string) *CardTemplateBuildActionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CardTemplateBuildActionRequest struct {
	// This parameter is required.
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// merge
	CardTemplateJson *string `json:"cardTemplateJson,omitempty" xml:"cardTemplateJson,omitempty"`
}

func (s CardTemplateBuildActionRequest) String() string {
	return tea.Prettify(s)
}

func (s CardTemplateBuildActionRequest) GoString() string {
	return s.String()
}

func (s *CardTemplateBuildActionRequest) SetAction(v string) *CardTemplateBuildActionRequest {
	s.Action = &v
	return s
}

func (s *CardTemplateBuildActionRequest) SetCardTemplateJson(v string) *CardTemplateBuildActionRequest {
	s.CardTemplateJson = &v
	return s
}

type CardTemplateBuildActionResponseBody struct {
	// example:
	//
	// {"xxx":"xxx"}
	CardTemplateJson *string `json:"cardTemplateJson,omitempty" xml:"cardTemplateJson,omitempty"`
}

func (s CardTemplateBuildActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CardTemplateBuildActionResponseBody) GoString() string {
	return s.String()
}

func (s *CardTemplateBuildActionResponseBody) SetCardTemplateJson(v string) *CardTemplateBuildActionResponseBody {
	s.CardTemplateJson = &v
	return s
}

type CardTemplateBuildActionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CardTemplateBuildActionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CardTemplateBuildActionResponse) String() string {
	return tea.Prettify(s)
}

func (s CardTemplateBuildActionResponse) GoString() string {
	return s.String()
}

func (s *CardTemplateBuildActionResponse) SetHeaders(v map[string]*string) *CardTemplateBuildActionResponse {
	s.Headers = v
	return s
}

func (s *CardTemplateBuildActionResponse) SetStatusCode(v int32) *CardTemplateBuildActionResponse {
	s.StatusCode = &v
	return s
}

func (s *CardTemplateBuildActionResponse) SetBody(v *CardTemplateBuildActionResponseBody) *CardTemplateBuildActionResponse {
	s.Body = v
	return s
}

type ChangeGroupOwnerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ChangeGroupOwnerHeaders) String() string {
	return tea.Prettify(s)
}

func (s ChangeGroupOwnerHeaders) GoString() string {
	return s.String()
}

func (s *ChangeGroupOwnerHeaders) SetCommonHeaders(v map[string]*string) *ChangeGroupOwnerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ChangeGroupOwnerHeaders) SetXAcsDingtalkAccessToken(v string) *ChangeGroupOwnerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ChangeGroupOwnerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 14da****2760
	GroupOwnerId *string `json:"groupOwnerId,omitempty" xml:"groupOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	GroupOwnerType *int32 `json:"groupOwnerType,omitempty" xml:"groupOwnerType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 14da****2760
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s ChangeGroupOwnerRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeGroupOwnerRequest) GoString() string {
	return s.String()
}

func (s *ChangeGroupOwnerRequest) SetGroupOwnerId(v string) *ChangeGroupOwnerRequest {
	s.GroupOwnerId = &v
	return s
}

func (s *ChangeGroupOwnerRequest) SetGroupOwnerType(v int32) *ChangeGroupOwnerRequest {
	s.GroupOwnerType = &v
	return s
}

func (s *ChangeGroupOwnerRequest) SetOpenConversationId(v string) *ChangeGroupOwnerRequest {
	s.OpenConversationId = &v
	return s
}

type ChangeGroupOwnerResponseBody struct {
	NewGroupOwnerId   *string `json:"newGroupOwnerId,omitempty" xml:"newGroupOwnerId,omitempty"`
	NewGroupOwnerType *int32  `json:"newGroupOwnerType,omitempty" xml:"newGroupOwnerType,omitempty"`
}

func (s ChangeGroupOwnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeGroupOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeGroupOwnerResponseBody) SetNewGroupOwnerId(v string) *ChangeGroupOwnerResponseBody {
	s.NewGroupOwnerId = &v
	return s
}

func (s *ChangeGroupOwnerResponseBody) SetNewGroupOwnerType(v int32) *ChangeGroupOwnerResponseBody {
	s.NewGroupOwnerType = &v
	return s
}

type ChangeGroupOwnerResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeGroupOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeGroupOwnerResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeGroupOwnerResponse) GoString() string {
	return s.String()
}

func (s *ChangeGroupOwnerResponse) SetHeaders(v map[string]*string) *ChangeGroupOwnerResponse {
	s.Headers = v
	return s
}

func (s *ChangeGroupOwnerResponse) SetStatusCode(v int32) *ChangeGroupOwnerResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeGroupOwnerResponse) SetBody(v *ChangeGroupOwnerResponseBody) *ChangeGroupOwnerResponse {
	s.Body = v
	return s
}

type ChatIdToOpenConversationIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ChatIdToOpenConversationIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s ChatIdToOpenConversationIdHeaders) GoString() string {
	return s.String()
}

func (s *ChatIdToOpenConversationIdHeaders) SetCommonHeaders(v map[string]*string) *ChatIdToOpenConversationIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ChatIdToOpenConversationIdHeaders) SetXAcsDingtalkAccessToken(v string) *ChatIdToOpenConversationIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ChatIdToOpenConversationIdResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// cidl1B8RVUFmkO50OC9uEbySQ==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s ChatIdToOpenConversationIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChatIdToOpenConversationIdResponseBody) GoString() string {
	return s.String()
}

func (s *ChatIdToOpenConversationIdResponseBody) SetOpenConversationId(v string) *ChatIdToOpenConversationIdResponseBody {
	s.OpenConversationId = &v
	return s
}

type ChatIdToOpenConversationIdResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatIdToOpenConversationIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatIdToOpenConversationIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ChatIdToOpenConversationIdResponse) GoString() string {
	return s.String()
}

func (s *ChatIdToOpenConversationIdResponse) SetHeaders(v map[string]*string) *ChatIdToOpenConversationIdResponse {
	s.Headers = v
	return s
}

func (s *ChatIdToOpenConversationIdResponse) SetStatusCode(v int32) *ChatIdToOpenConversationIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatIdToOpenConversationIdResponse) SetBody(v *ChatIdToOpenConversationIdResponseBody) *ChatIdToOpenConversationIdResponse {
	s.Body = v
	return s
}

type ChatSubAdminUpdateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ChatSubAdminUpdateHeaders) String() string {
	return tea.Prettify(s)
}

func (s ChatSubAdminUpdateHeaders) GoString() string {
	return s.String()
}

func (s *ChatSubAdminUpdateHeaders) SetCommonHeaders(v map[string]*string) *ChatSubAdminUpdateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ChatSubAdminUpdateHeaders) SetXAcsDingtalkAccessToken(v string) *ChatSubAdminUpdateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ChatSubAdminUpdateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cidVwhmrlxsR3sL3+JdH1LjUA==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	Role *int32 `json:"role,omitempty" xml:"role,omitempty"`
	// This parameter is required.
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s ChatSubAdminUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s ChatSubAdminUpdateRequest) GoString() string {
	return s.String()
}

func (s *ChatSubAdminUpdateRequest) SetOpenConversationId(v string) *ChatSubAdminUpdateRequest {
	s.OpenConversationId = &v
	return s
}

func (s *ChatSubAdminUpdateRequest) SetRole(v int32) *ChatSubAdminUpdateRequest {
	s.Role = &v
	return s
}

func (s *ChatSubAdminUpdateRequest) SetUserIds(v []*string) *ChatSubAdminUpdateRequest {
	s.UserIds = v
	return s
}

type ChatSubAdminUpdateResponseBody struct {
	// This parameter is required.
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ChatSubAdminUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChatSubAdminUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *ChatSubAdminUpdateResponseBody) SetSuccess(v string) *ChatSubAdminUpdateResponseBody {
	s.Success = &v
	return s
}

type ChatSubAdminUpdateResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatSubAdminUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatSubAdminUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s ChatSubAdminUpdateResponse) GoString() string {
	return s.String()
}

func (s *ChatSubAdminUpdateResponse) SetHeaders(v map[string]*string) *ChatSubAdminUpdateResponse {
	s.Headers = v
	return s
}

func (s *ChatSubAdminUpdateResponse) SetStatusCode(v int32) *ChatSubAdminUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatSubAdminUpdateResponse) SetBody(v *ChatSubAdminUpdateResponseBody) *ChatSubAdminUpdateResponse {
	s.Body = v
	return s
}

type CheckUserIsGroupMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CheckUserIsGroupMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s CheckUserIsGroupMemberHeaders) GoString() string {
	return s.String()
}

func (s *CheckUserIsGroupMemberHeaders) SetCommonHeaders(v map[string]*string) *CheckUserIsGroupMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CheckUserIsGroupMemberHeaders) SetXAcsDingtalkAccessToken(v string) *CheckUserIsGroupMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CheckUserIsGroupMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cidD2y*****==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 015*****
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CheckUserIsGroupMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckUserIsGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *CheckUserIsGroupMemberRequest) SetOpenConversationId(v string) *CheckUserIsGroupMemberRequest {
	s.OpenConversationId = &v
	return s
}

func (s *CheckUserIsGroupMemberRequest) SetUserId(v string) *CheckUserIsGroupMemberRequest {
	s.UserId = &v
	return s
}

type CheckUserIsGroupMemberResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CheckUserIsGroupMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckUserIsGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUserIsGroupMemberResponseBody) SetResult(v bool) *CheckUserIsGroupMemberResponseBody {
	s.Result = &v
	return s
}

type CheckUserIsGroupMemberResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckUserIsGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckUserIsGroupMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckUserIsGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *CheckUserIsGroupMemberResponse) SetHeaders(v map[string]*string) *CheckUserIsGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *CheckUserIsGroupMemberResponse) SetStatusCode(v int32) *CheckUserIsGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckUserIsGroupMemberResponse) SetBody(v *CheckUserIsGroupMemberResponseBody) *CheckUserIsGroupMemberResponse {
	s.Body = v
	return s
}

type CopyUnfurlingRegisterHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CopyUnfurlingRegisterHeaders) String() string {
	return tea.Prettify(s)
}

func (s CopyUnfurlingRegisterHeaders) GoString() string {
	return s.String()
}

func (s *CopyUnfurlingRegisterHeaders) SetCommonHeaders(v map[string]*string) *CopyUnfurlingRegisterHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CopyUnfurlingRegisterHeaders) SetXAcsDingtalkAccessToken(v string) *CopyUnfurlingRegisterHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CopyUnfurlingRegisterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3102xxxxxxx
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// This parameter is required.
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 37xxxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CopyUnfurlingRegisterRequest) String() string {
	return tea.Prettify(s)
}

func (s CopyUnfurlingRegisterRequest) GoString() string {
	return s.String()
}

func (s *CopyUnfurlingRegisterRequest) SetAppId(v string) *CopyUnfurlingRegisterRequest {
	s.AppId = &v
	return s
}

func (s *CopyUnfurlingRegisterRequest) SetId(v int64) *CopyUnfurlingRegisterRequest {
	s.Id = &v
	return s
}

func (s *CopyUnfurlingRegisterRequest) SetUserId(v string) *CopyUnfurlingRegisterRequest {
	s.UserId = &v
	return s
}

type CopyUnfurlingRegisterResponseBody struct {
	Result  *int64 `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CopyUnfurlingRegisterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CopyUnfurlingRegisterResponseBody) GoString() string {
	return s.String()
}

func (s *CopyUnfurlingRegisterResponseBody) SetResult(v int64) *CopyUnfurlingRegisterResponseBody {
	s.Result = &v
	return s
}

func (s *CopyUnfurlingRegisterResponseBody) SetSuccess(v bool) *CopyUnfurlingRegisterResponseBody {
	s.Success = &v
	return s
}

type CopyUnfurlingRegisterResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyUnfurlingRegisterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyUnfurlingRegisterResponse) String() string {
	return tea.Prettify(s)
}

func (s CopyUnfurlingRegisterResponse) GoString() string {
	return s.String()
}

func (s *CopyUnfurlingRegisterResponse) SetHeaders(v map[string]*string) *CopyUnfurlingRegisterResponse {
	s.Headers = v
	return s
}

func (s *CopyUnfurlingRegisterResponse) SetStatusCode(v int32) *CopyUnfurlingRegisterResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyUnfurlingRegisterResponse) SetBody(v *CopyUnfurlingRegisterResponseBody) *CopyUnfurlingRegisterResponse {
	s.Body = v
	return s
}

type CountOpenMsgSceneGroupsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CountOpenMsgSceneGroupsHeaders) String() string {
	return tea.Prettify(s)
}

func (s CountOpenMsgSceneGroupsHeaders) GoString() string {
	return s.String()
}

func (s *CountOpenMsgSceneGroupsHeaders) SetCommonHeaders(v map[string]*string) *CountOpenMsgSceneGroupsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CountOpenMsgSceneGroupsHeaders) SetXAcsDingtalkAccessToken(v string) *CountOpenMsgSceneGroupsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CountOpenMsgSceneGroupsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// f6xxxxx
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s CountOpenMsgSceneGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s CountOpenMsgSceneGroupsRequest) GoString() string {
	return s.String()
}

func (s *CountOpenMsgSceneGroupsRequest) SetTemplateId(v string) *CountOpenMsgSceneGroupsRequest {
	s.TemplateId = &v
	return s
}

type CountOpenMsgSceneGroupsResponseBody struct {
	Result  *CountOpenMsgSceneGroupsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CountOpenMsgSceneGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountOpenMsgSceneGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *CountOpenMsgSceneGroupsResponseBody) SetResult(v *CountOpenMsgSceneGroupsResponseBodyResult) *CountOpenMsgSceneGroupsResponseBody {
	s.Result = v
	return s
}

func (s *CountOpenMsgSceneGroupsResponseBody) SetSuccess(v bool) *CountOpenMsgSceneGroupsResponseBody {
	s.Success = &v
	return s
}

type CountOpenMsgSceneGroupsResponseBodyResult struct {
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
}

func (s CountOpenMsgSceneGroupsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CountOpenMsgSceneGroupsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CountOpenMsgSceneGroupsResponseBodyResult) SetCount(v int64) *CountOpenMsgSceneGroupsResponseBodyResult {
	s.Count = &v
	return s
}

type CountOpenMsgSceneGroupsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CountOpenMsgSceneGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CountOpenMsgSceneGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s CountOpenMsgSceneGroupsResponse) GoString() string {
	return s.String()
}

func (s *CountOpenMsgSceneGroupsResponse) SetHeaders(v map[string]*string) *CountOpenMsgSceneGroupsResponse {
	s.Headers = v
	return s
}

func (s *CountOpenMsgSceneGroupsResponse) SetStatusCode(v int32) *CountOpenMsgSceneGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *CountOpenMsgSceneGroupsResponse) SetBody(v *CountOpenMsgSceneGroupsResponseBody) *CountOpenMsgSceneGroupsResponse {
	s.Body = v
	return s
}

type CountOrgMessageOpenSceneGroupsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CountOrgMessageOpenSceneGroupsHeaders) String() string {
	return tea.Prettify(s)
}

func (s CountOrgMessageOpenSceneGroupsHeaders) GoString() string {
	return s.String()
}

func (s *CountOrgMessageOpenSceneGroupsHeaders) SetCommonHeaders(v map[string]*string) *CountOrgMessageOpenSceneGroupsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CountOrgMessageOpenSceneGroupsHeaders) SetXAcsDingtalkAccessToken(v string) *CountOrgMessageOpenSceneGroupsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CountOrgMessageOpenSceneGroupsResponseBody struct {
	Count   *int32 `json:"count,omitempty" xml:"count,omitempty"`
	Success *bool  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CountOrgMessageOpenSceneGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountOrgMessageOpenSceneGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *CountOrgMessageOpenSceneGroupsResponseBody) SetCount(v int32) *CountOrgMessageOpenSceneGroupsResponseBody {
	s.Count = &v
	return s
}

func (s *CountOrgMessageOpenSceneGroupsResponseBody) SetSuccess(v bool) *CountOrgMessageOpenSceneGroupsResponseBody {
	s.Success = &v
	return s
}

type CountOrgMessageOpenSceneGroupsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CountOrgMessageOpenSceneGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CountOrgMessageOpenSceneGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s CountOrgMessageOpenSceneGroupsResponse) GoString() string {
	return s.String()
}

func (s *CountOrgMessageOpenSceneGroupsResponse) SetHeaders(v map[string]*string) *CountOrgMessageOpenSceneGroupsResponse {
	s.Headers = v
	return s
}

func (s *CountOrgMessageOpenSceneGroupsResponse) SetStatusCode(v int32) *CountOrgMessageOpenSceneGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *CountOrgMessageOpenSceneGroupsResponse) SetBody(v *CountOrgMessageOpenSceneGroupsResponseBody) *CountOrgMessageOpenSceneGroupsResponse {
	s.Body = v
	return s
}

type CountSceneGroupsByTemplateIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CountSceneGroupsByTemplateIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s CountSceneGroupsByTemplateIdHeaders) GoString() string {
	return s.String()
}

func (s *CountSceneGroupsByTemplateIdHeaders) SetCommonHeaders(v map[string]*string) *CountSceneGroupsByTemplateIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CountSceneGroupsByTemplateIdHeaders) SetXAcsDingtalkAccessToken(v string) *CountSceneGroupsByTemplateIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CountSceneGroupsByTemplateIdResponseBody struct {
	Count   *int32 `json:"count,omitempty" xml:"count,omitempty"`
	Success *bool  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CountSceneGroupsByTemplateIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CountSceneGroupsByTemplateIdResponseBody) GoString() string {
	return s.String()
}

func (s *CountSceneGroupsByTemplateIdResponseBody) SetCount(v int32) *CountSceneGroupsByTemplateIdResponseBody {
	s.Count = &v
	return s
}

func (s *CountSceneGroupsByTemplateIdResponseBody) SetSuccess(v bool) *CountSceneGroupsByTemplateIdResponseBody {
	s.Success = &v
	return s
}

type CountSceneGroupsByTemplateIdResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CountSceneGroupsByTemplateIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CountSceneGroupsByTemplateIdResponse) String() string {
	return tea.Prettify(s)
}

func (s CountSceneGroupsByTemplateIdResponse) GoString() string {
	return s.String()
}

func (s *CountSceneGroupsByTemplateIdResponse) SetHeaders(v map[string]*string) *CountSceneGroupsByTemplateIdResponse {
	s.Headers = v
	return s
}

func (s *CountSceneGroupsByTemplateIdResponse) SetStatusCode(v int32) *CountSceneGroupsByTemplateIdResponse {
	s.StatusCode = &v
	return s
}

func (s *CountSceneGroupsByTemplateIdResponse) SetBody(v *CountSceneGroupsByTemplateIdResponseBody) *CountSceneGroupsByTemplateIdResponse {
	s.Body = v
	return s
}

type CreateCoupleGroupConversationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateCoupleGroupConversationHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateCoupleGroupConversationHeaders) GoString() string {
	return s.String()
}

func (s *CreateCoupleGroupConversationHeaders) SetCommonHeaders(v map[string]*string) *CreateCoupleGroupConversationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateCoupleGroupConversationHeaders) SetXAcsDingtalkAccessToken(v string) *CreateCoupleGroupConversationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateCoupleGroupConversationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1107****2121
	AppUserId *string `json:"appUserId,omitempty" xml:"appUserId,omitempty"`
	// example:
	//
	// http://***.png
	GroupAvatar *string `json:"groupAvatar,omitempty" xml:"groupAvatar,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 客户群
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1107****2120
	GroupOwnerId *string `json:"groupOwnerId,omitempty" xml:"groupOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8d42****nkld
	GroupTemplateId *string `json:"groupTemplateId,omitempty" xml:"groupTemplateId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1107****2120
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
}

func (s CreateCoupleGroupConversationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCoupleGroupConversationRequest) GoString() string {
	return s.String()
}

func (s *CreateCoupleGroupConversationRequest) SetAppUserId(v string) *CreateCoupleGroupConversationRequest {
	s.AppUserId = &v
	return s
}

func (s *CreateCoupleGroupConversationRequest) SetGroupAvatar(v string) *CreateCoupleGroupConversationRequest {
	s.GroupAvatar = &v
	return s
}

func (s *CreateCoupleGroupConversationRequest) SetGroupName(v string) *CreateCoupleGroupConversationRequest {
	s.GroupName = &v
	return s
}

func (s *CreateCoupleGroupConversationRequest) SetGroupOwnerId(v string) *CreateCoupleGroupConversationRequest {
	s.GroupOwnerId = &v
	return s
}

func (s *CreateCoupleGroupConversationRequest) SetGroupTemplateId(v string) *CreateCoupleGroupConversationRequest {
	s.GroupTemplateId = &v
	return s
}

func (s *CreateCoupleGroupConversationRequest) SetOperatorId(v string) *CreateCoupleGroupConversationRequest {
	s.OperatorId = &v
	return s
}

type CreateCoupleGroupConversationResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// cid****8Q==
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 14da****2760
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s CreateCoupleGroupConversationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCoupleGroupConversationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCoupleGroupConversationResponseBody) SetConversationId(v string) *CreateCoupleGroupConversationResponseBody {
	s.ConversationId = &v
	return s
}

func (s *CreateCoupleGroupConversationResponseBody) SetOpenConversationId(v string) *CreateCoupleGroupConversationResponseBody {
	s.OpenConversationId = &v
	return s
}

type CreateCoupleGroupConversationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCoupleGroupConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCoupleGroupConversationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCoupleGroupConversationResponse) GoString() string {
	return s.String()
}

func (s *CreateCoupleGroupConversationResponse) SetHeaders(v map[string]*string) *CreateCoupleGroupConversationResponse {
	s.Headers = v
	return s
}

func (s *CreateCoupleGroupConversationResponse) SetStatusCode(v int32) *CreateCoupleGroupConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCoupleGroupConversationResponse) SetBody(v *CreateCoupleGroupConversationResponseBody) *CreateCoupleGroupConversationResponse {
	s.Body = v
	return s
}

type CreateGroupConversationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateGroupConversationHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupConversationHeaders) GoString() string {
	return s.String()
}

func (s *CreateGroupConversationHeaders) SetCommonHeaders(v map[string]*string) *CreateGroupConversationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateGroupConversationHeaders) SetXAcsDingtalkAccessToken(v string) *CreateGroupConversationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateGroupConversationRequest struct {
	AppUserIds []*string `json:"appUserIds,omitempty" xml:"appUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// http://***.png
	GroupAvatar *string `json:"groupAvatar,omitempty" xml:"groupAvatar,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 客户群
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1745****8777
	GroupOwnerId *string `json:"groupOwnerId,omitempty" xml:"groupOwnerId,omitempty"`
	// example:
	//
	// 3
	GroupOwnerType *int32 `json:"groupOwnerType,omitempty" xml:"groupOwnerType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8d42****nkld
	GroupTemplateId *string `json:"groupTemplateId,omitempty" xml:"groupTemplateId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1745****8777
	OperatorId *string   `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	UserIds    []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s CreateGroupConversationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupConversationRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupConversationRequest) SetAppUserIds(v []*string) *CreateGroupConversationRequest {
	s.AppUserIds = v
	return s
}

func (s *CreateGroupConversationRequest) SetGroupAvatar(v string) *CreateGroupConversationRequest {
	s.GroupAvatar = &v
	return s
}

func (s *CreateGroupConversationRequest) SetGroupName(v string) *CreateGroupConversationRequest {
	s.GroupName = &v
	return s
}

func (s *CreateGroupConversationRequest) SetGroupOwnerId(v string) *CreateGroupConversationRequest {
	s.GroupOwnerId = &v
	return s
}

func (s *CreateGroupConversationRequest) SetGroupOwnerType(v int32) *CreateGroupConversationRequest {
	s.GroupOwnerType = &v
	return s
}

func (s *CreateGroupConversationRequest) SetGroupTemplateId(v string) *CreateGroupConversationRequest {
	s.GroupTemplateId = &v
	return s
}

func (s *CreateGroupConversationRequest) SetOperatorId(v string) *CreateGroupConversationRequest {
	s.OperatorId = &v
	return s
}

func (s *CreateGroupConversationRequest) SetUserIds(v []*string) *CreateGroupConversationRequest {
	s.UserIds = v
	return s
}

type CreateGroupConversationResponseBody struct {
	// This parameter is required.
	AppUserIds []*string `json:"appUserIds,omitempty" xml:"appUserIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cidpZ****Vcp4g==
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 14da****2760
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s CreateGroupConversationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupConversationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupConversationResponseBody) SetAppUserIds(v []*string) *CreateGroupConversationResponseBody {
	s.AppUserIds = v
	return s
}

func (s *CreateGroupConversationResponseBody) SetConversationId(v string) *CreateGroupConversationResponseBody {
	s.ConversationId = &v
	return s
}

func (s *CreateGroupConversationResponseBody) SetOpenConversationId(v string) *CreateGroupConversationResponseBody {
	s.OpenConversationId = &v
	return s
}

func (s *CreateGroupConversationResponseBody) SetUserIds(v []*string) *CreateGroupConversationResponseBody {
	s.UserIds = v
	return s
}

type CreateGroupConversationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGroupConversationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupConversationResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupConversationResponse) SetHeaders(v map[string]*string) *CreateGroupConversationResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupConversationResponse) SetStatusCode(v int32) *CreateGroupConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupConversationResponse) SetBody(v *CreateGroupConversationResponseBody) *CreateGroupConversationResponse {
	s.Body = v
	return s
}

type CreateInterconnectionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateInterconnectionHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateInterconnectionHeaders) GoString() string {
	return s.String()
}

func (s *CreateInterconnectionHeaders) SetCommonHeaders(v map[string]*string) *CreateInterconnectionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateInterconnectionHeaders) SetXAcsDingtalkAccessToken(v string) *CreateInterconnectionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateInterconnectionRequest struct {
	// This parameter is required.
	Interconnections []*CreateInterconnectionRequestInterconnections `json:"interconnections,omitempty" xml:"interconnections,omitempty" type:"Repeated"`
}

func (s CreateInterconnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInterconnectionRequest) GoString() string {
	return s.String()
}

func (s *CreateInterconnectionRequest) SetInterconnections(v []*CreateInterconnectionRequestInterconnections) *CreateInterconnectionRequest {
	s.Interconnections = v
	return s
}

type CreateInterconnectionRequestInterconnections struct {
	// example:
	//
	// http://***.png
	AppUserAvatar *string `json:"appUserAvatar,omitempty" xml:"appUserAvatar,omitempty"`
	// example:
	//
	// 1
	AppUserAvatarMediaType *int32 `json:"appUserAvatarMediaType,omitempty" xml:"appUserAvatarMediaType,omitempty"`
	// example:
	//
	// 认真工作,快乐生活
	AppUserDynamics *string `json:"appUserDynamics,omitempty" xml:"appUserDynamics,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1107****2120
	AppUserId *string `json:"appUserId,omitempty" xml:"appUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 188****8655
	AppUserMobile *string `json:"appUserMobile,omitempty" xml:"appUserMobile,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Foo
	AppUserName *string `json:"appUserName,omitempty" xml:"appUserName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ChannelCode *string `json:"channelCode,omitempty" xml:"channelCode,omitempty"`
	// example:
	//
	// 1745****8777
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateInterconnectionRequestInterconnections) String() string {
	return tea.Prettify(s)
}

func (s CreateInterconnectionRequestInterconnections) GoString() string {
	return s.String()
}

func (s *CreateInterconnectionRequestInterconnections) SetAppUserAvatar(v string) *CreateInterconnectionRequestInterconnections {
	s.AppUserAvatar = &v
	return s
}

func (s *CreateInterconnectionRequestInterconnections) SetAppUserAvatarMediaType(v int32) *CreateInterconnectionRequestInterconnections {
	s.AppUserAvatarMediaType = &v
	return s
}

func (s *CreateInterconnectionRequestInterconnections) SetAppUserDynamics(v string) *CreateInterconnectionRequestInterconnections {
	s.AppUserDynamics = &v
	return s
}

func (s *CreateInterconnectionRequestInterconnections) SetAppUserId(v string) *CreateInterconnectionRequestInterconnections {
	s.AppUserId = &v
	return s
}

func (s *CreateInterconnectionRequestInterconnections) SetAppUserMobile(v string) *CreateInterconnectionRequestInterconnections {
	s.AppUserMobile = &v
	return s
}

func (s *CreateInterconnectionRequestInterconnections) SetAppUserName(v string) *CreateInterconnectionRequestInterconnections {
	s.AppUserName = &v
	return s
}

func (s *CreateInterconnectionRequestInterconnections) SetChannelCode(v string) *CreateInterconnectionRequestInterconnections {
	s.ChannelCode = &v
	return s
}

func (s *CreateInterconnectionRequestInterconnections) SetUserId(v string) *CreateInterconnectionRequestInterconnections {
	s.UserId = &v
	return s
}

type CreateInterconnectionResponseBody struct {
	Results []*CreateInterconnectionResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s CreateInterconnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInterconnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInterconnectionResponseBody) SetResults(v []*CreateInterconnectionResponseBodyResults) *CreateInterconnectionResponseBody {
	s.Results = v
	return s
}

type CreateInterconnectionResponseBodyResults struct {
	AppUserId *string `json:"appUserId,omitempty" xml:"appUserId,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateInterconnectionResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s CreateInterconnectionResponseBodyResults) GoString() string {
	return s.String()
}

func (s *CreateInterconnectionResponseBodyResults) SetAppUserId(v string) *CreateInterconnectionResponseBodyResults {
	s.AppUserId = &v
	return s
}

func (s *CreateInterconnectionResponseBodyResults) SetMessage(v string) *CreateInterconnectionResponseBodyResults {
	s.Message = &v
	return s
}

func (s *CreateInterconnectionResponseBodyResults) SetUserId(v string) *CreateInterconnectionResponseBodyResults {
	s.UserId = &v
	return s
}

type CreateInterconnectionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInterconnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInterconnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInterconnectionResponse) GoString() string {
	return s.String()
}

func (s *CreateInterconnectionResponse) SetHeaders(v map[string]*string) *CreateInterconnectionResponse {
	s.Headers = v
	return s
}

func (s *CreateInterconnectionResponse) SetStatusCode(v int32) *CreateInterconnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInterconnectionResponse) SetBody(v *CreateInterconnectionResponseBody) *CreateInterconnectionResponse {
	s.Body = v
	return s
}

type CreateSceneGroupConversationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateSceneGroupConversationHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateSceneGroupConversationHeaders) GoString() string {
	return s.String()
}

func (s *CreateSceneGroupConversationHeaders) SetCommonHeaders(v map[string]*string) *CreateSceneGroupConversationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSceneGroupConversationHeaders) SetXAcsDingtalkAccessToken(v string) *CreateSceneGroupConversationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateSceneGroupConversationRequest struct {
	Features map[string]*string `json:"features,omitempty" xml:"features,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 客户群
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1107****2120
	GroupOwnerId *string `json:"groupOwnerId,omitempty" xml:"groupOwnerId,omitempty"`
	// example:
	//
	// http://***.png
	Icon              *string                                               `json:"icon,omitempty" xml:"icon,omitempty"`
	ManagementOptions *CreateSceneGroupConversationRequestManagementOptions `json:"managementOptions,omitempty" xml:"managementOptions,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 8d42****nkld
	TemplateId *string   `json:"templateId,omitempty" xml:"templateId,omitempty"`
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
	// example:
	//
	// asdazxc
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s CreateSceneGroupConversationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSceneGroupConversationRequest) GoString() string {
	return s.String()
}

func (s *CreateSceneGroupConversationRequest) SetFeatures(v map[string]*string) *CreateSceneGroupConversationRequest {
	s.Features = v
	return s
}

func (s *CreateSceneGroupConversationRequest) SetGroupName(v string) *CreateSceneGroupConversationRequest {
	s.GroupName = &v
	return s
}

func (s *CreateSceneGroupConversationRequest) SetGroupOwnerId(v string) *CreateSceneGroupConversationRequest {
	s.GroupOwnerId = &v
	return s
}

func (s *CreateSceneGroupConversationRequest) SetIcon(v string) *CreateSceneGroupConversationRequest {
	s.Icon = &v
	return s
}

func (s *CreateSceneGroupConversationRequest) SetManagementOptions(v *CreateSceneGroupConversationRequestManagementOptions) *CreateSceneGroupConversationRequest {
	s.ManagementOptions = v
	return s
}

func (s *CreateSceneGroupConversationRequest) SetTemplateId(v string) *CreateSceneGroupConversationRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateSceneGroupConversationRequest) SetUserIdList(v []*string) *CreateSceneGroupConversationRequest {
	s.UserIdList = v
	return s
}

func (s *CreateSceneGroupConversationRequest) SetUuid(v string) *CreateSceneGroupConversationRequest {
	s.Uuid = &v
	return s
}

type CreateSceneGroupConversationRequestManagementOptions struct {
	ChatBannedType      *int32 `json:"chatBannedType,omitempty" xml:"chatBannedType,omitempty"`
	ManagementType      *int32 `json:"managementType,omitempty" xml:"managementType,omitempty"`
	MentionAllAuthority *int32 `json:"mentionAllAuthority,omitempty" xml:"mentionAllAuthority,omitempty"`
	Searchable          *int32 `json:"searchable,omitempty" xml:"searchable,omitempty"`
	ShowHistoryType     *int32 `json:"showHistoryType,omitempty" xml:"showHistoryType,omitempty"`
	ValidationType      *int32 `json:"validationType,omitempty" xml:"validationType,omitempty"`
}

func (s CreateSceneGroupConversationRequestManagementOptions) String() string {
	return tea.Prettify(s)
}

func (s CreateSceneGroupConversationRequestManagementOptions) GoString() string {
	return s.String()
}

func (s *CreateSceneGroupConversationRequestManagementOptions) SetChatBannedType(v int32) *CreateSceneGroupConversationRequestManagementOptions {
	s.ChatBannedType = &v
	return s
}

func (s *CreateSceneGroupConversationRequestManagementOptions) SetManagementType(v int32) *CreateSceneGroupConversationRequestManagementOptions {
	s.ManagementType = &v
	return s
}

func (s *CreateSceneGroupConversationRequestManagementOptions) SetMentionAllAuthority(v int32) *CreateSceneGroupConversationRequestManagementOptions {
	s.MentionAllAuthority = &v
	return s
}

func (s *CreateSceneGroupConversationRequestManagementOptions) SetSearchable(v int32) *CreateSceneGroupConversationRequestManagementOptions {
	s.Searchable = &v
	return s
}

func (s *CreateSceneGroupConversationRequestManagementOptions) SetShowHistoryType(v int32) *CreateSceneGroupConversationRequestManagementOptions {
	s.ShowHistoryType = &v
	return s
}

func (s *CreateSceneGroupConversationRequestManagementOptions) SetValidationType(v int32) *CreateSceneGroupConversationRequestManagementOptions {
	s.ValidationType = &v
	return s
}

type CreateSceneGroupConversationResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// cidxxxxxx==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s CreateSceneGroupConversationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSceneGroupConversationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSceneGroupConversationResponseBody) SetOpenConversationId(v string) *CreateSceneGroupConversationResponseBody {
	s.OpenConversationId = &v
	return s
}

type CreateSceneGroupConversationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSceneGroupConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSceneGroupConversationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSceneGroupConversationResponse) GoString() string {
	return s.String()
}

func (s *CreateSceneGroupConversationResponse) SetHeaders(v map[string]*string) *CreateSceneGroupConversationResponse {
	s.Headers = v
	return s
}

func (s *CreateSceneGroupConversationResponse) SetStatusCode(v int32) *CreateSceneGroupConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSceneGroupConversationResponse) SetBody(v *CreateSceneGroupConversationResponseBody) *CreateSceneGroupConversationResponse {
	s.Body = v
	return s
}

type CreateStoreGroupConversationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateStoreGroupConversationHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateStoreGroupConversationHeaders) GoString() string {
	return s.String()
}

func (s *CreateStoreGroupConversationHeaders) SetCommonHeaders(v map[string]*string) *CreateStoreGroupConversationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateStoreGroupConversationHeaders) SetXAcsDingtalkAccessToken(v string) *CreateStoreGroupConversationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateStoreGroupConversationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1107****2120
	AppUserId *string `json:"appUserId,omitempty" xml:"appUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// store1
	BusinessUniqueKey *string `json:"businessUniqueKey,omitempty" xml:"businessUniqueKey,omitempty"`
	// example:
	//
	// http://***.png
	GroupAvatar *string `json:"groupAvatar,omitempty" xml:"groupAvatar,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 客户群
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8d42****nkld
	GroupTemplateId *string `json:"groupTemplateId,omitempty" xml:"groupTemplateId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1107****2120
	OperatorId *string   `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	UserIds    []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s CreateStoreGroupConversationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStoreGroupConversationRequest) GoString() string {
	return s.String()
}

func (s *CreateStoreGroupConversationRequest) SetAppUserId(v string) *CreateStoreGroupConversationRequest {
	s.AppUserId = &v
	return s
}

func (s *CreateStoreGroupConversationRequest) SetBusinessUniqueKey(v string) *CreateStoreGroupConversationRequest {
	s.BusinessUniqueKey = &v
	return s
}

func (s *CreateStoreGroupConversationRequest) SetGroupAvatar(v string) *CreateStoreGroupConversationRequest {
	s.GroupAvatar = &v
	return s
}

func (s *CreateStoreGroupConversationRequest) SetGroupName(v string) *CreateStoreGroupConversationRequest {
	s.GroupName = &v
	return s
}

func (s *CreateStoreGroupConversationRequest) SetGroupTemplateId(v string) *CreateStoreGroupConversationRequest {
	s.GroupTemplateId = &v
	return s
}

func (s *CreateStoreGroupConversationRequest) SetOperatorId(v string) *CreateStoreGroupConversationRequest {
	s.OperatorId = &v
	return s
}

func (s *CreateStoreGroupConversationRequest) SetUserIds(v []*string) *CreateStoreGroupConversationRequest {
	s.UserIds = v
	return s
}

type CreateStoreGroupConversationResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// cid****8Q==
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 14da****2760
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s CreateStoreGroupConversationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateStoreGroupConversationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStoreGroupConversationResponseBody) SetConversationId(v string) *CreateStoreGroupConversationResponseBody {
	s.ConversationId = &v
	return s
}

func (s *CreateStoreGroupConversationResponseBody) SetOpenConversationId(v string) *CreateStoreGroupConversationResponseBody {
	s.OpenConversationId = &v
	return s
}

type CreateStoreGroupConversationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStoreGroupConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStoreGroupConversationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStoreGroupConversationResponse) GoString() string {
	return s.String()
}

func (s *CreateStoreGroupConversationResponse) SetHeaders(v map[string]*string) *CreateStoreGroupConversationResponse {
	s.Headers = v
	return s
}

func (s *CreateStoreGroupConversationResponse) SetStatusCode(v int32) *CreateStoreGroupConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStoreGroupConversationResponse) SetBody(v *CreateStoreGroupConversationResponseBody) *CreateStoreGroupConversationResponse {
	s.Body = v
	return s
}

type DebugUnfurlingRegisterHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DebugUnfurlingRegisterHeaders) String() string {
	return tea.Prettify(s)
}

func (s DebugUnfurlingRegisterHeaders) GoString() string {
	return s.String()
}

func (s *DebugUnfurlingRegisterHeaders) SetCommonHeaders(v map[string]*string) *DebugUnfurlingRegisterHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DebugUnfurlingRegisterHeaders) SetXAcsDingtalkAccessToken(v string) *DebugUnfurlingRegisterHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DebugUnfurlingRegisterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3102xxxxxxx
	AppId           *string   `json:"appId,omitempty" xml:"appId,omitempty"`
	GrayGroupIdList []*string `json:"grayGroupIdList,omitempty" xml:"grayGroupIdList,omitempty" type:"Repeated"`
	GrayUserIdList  []*string `json:"grayUserIdList,omitempty" xml:"grayUserIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 37xxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DebugUnfurlingRegisterRequest) String() string {
	return tea.Prettify(s)
}

func (s DebugUnfurlingRegisterRequest) GoString() string {
	return s.String()
}

func (s *DebugUnfurlingRegisterRequest) SetAppId(v string) *DebugUnfurlingRegisterRequest {
	s.AppId = &v
	return s
}

func (s *DebugUnfurlingRegisterRequest) SetGrayGroupIdList(v []*string) *DebugUnfurlingRegisterRequest {
	s.GrayGroupIdList = v
	return s
}

func (s *DebugUnfurlingRegisterRequest) SetGrayUserIdList(v []*string) *DebugUnfurlingRegisterRequest {
	s.GrayUserIdList = v
	return s
}

func (s *DebugUnfurlingRegisterRequest) SetId(v int64) *DebugUnfurlingRegisterRequest {
	s.Id = &v
	return s
}

func (s *DebugUnfurlingRegisterRequest) SetUserId(v string) *DebugUnfurlingRegisterRequest {
	s.UserId = &v
	return s
}

type DebugUnfurlingRegisterResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DebugUnfurlingRegisterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DebugUnfurlingRegisterResponseBody) GoString() string {
	return s.String()
}

func (s *DebugUnfurlingRegisterResponseBody) SetSuccess(v bool) *DebugUnfurlingRegisterResponseBody {
	s.Success = &v
	return s
}

type DebugUnfurlingRegisterResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DebugUnfurlingRegisterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DebugUnfurlingRegisterResponse) String() string {
	return tea.Prettify(s)
}

func (s DebugUnfurlingRegisterResponse) GoString() string {
	return s.String()
}

func (s *DebugUnfurlingRegisterResponse) SetHeaders(v map[string]*string) *DebugUnfurlingRegisterResponse {
	s.Headers = v
	return s
}

func (s *DebugUnfurlingRegisterResponse) SetStatusCode(v int32) *DebugUnfurlingRegisterResponse {
	s.StatusCode = &v
	return s
}

func (s *DebugUnfurlingRegisterResponse) SetBody(v *DebugUnfurlingRegisterResponseBody) *DebugUnfurlingRegisterResponse {
	s.Body = v
	return s
}

type DeleteOrgTextEmotionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteOrgTextEmotionHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrgTextEmotionHeaders) GoString() string {
	return s.String()
}

func (s *DeleteOrgTextEmotionHeaders) SetCommonHeaders(v map[string]*string) *DeleteOrgTextEmotionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteOrgTextEmotionHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteOrgTextEmotionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteOrgTextEmotionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// -1
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// This parameter is required.
	EmotionIds []*string `json:"emotionIds,omitempty" xml:"emotionIds,omitempty" type:"Repeated"`
}

func (s DeleteOrgTextEmotionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrgTextEmotionRequest) GoString() string {
	return s.String()
}

func (s *DeleteOrgTextEmotionRequest) SetDeptId(v int64) *DeleteOrgTextEmotionRequest {
	s.DeptId = &v
	return s
}

func (s *DeleteOrgTextEmotionRequest) SetEmotionIds(v []*string) *DeleteOrgTextEmotionRequest {
	s.EmotionIds = v
	return s
}

type DeleteOrgTextEmotionResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteOrgTextEmotionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrgTextEmotionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOrgTextEmotionResponseBody) SetSuccess(v bool) *DeleteOrgTextEmotionResponseBody {
	s.Success = &v
	return s
}

type DeleteOrgTextEmotionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOrgTextEmotionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOrgTextEmotionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrgTextEmotionResponse) GoString() string {
	return s.String()
}

func (s *DeleteOrgTextEmotionResponse) SetHeaders(v map[string]*string) *DeleteOrgTextEmotionResponse {
	s.Headers = v
	return s
}

func (s *DeleteOrgTextEmotionResponse) SetStatusCode(v int32) *DeleteOrgTextEmotionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOrgTextEmotionResponse) SetBody(v *DeleteOrgTextEmotionResponseBody) *DeleteOrgTextEmotionResponse {
	s.Body = v
	return s
}

type DismissGroupConversationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DismissGroupConversationHeaders) String() string {
	return tea.Prettify(s)
}

func (s DismissGroupConversationHeaders) GoString() string {
	return s.String()
}

func (s *DismissGroupConversationHeaders) SetCommonHeaders(v map[string]*string) *DismissGroupConversationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DismissGroupConversationHeaders) SetXAcsDingtalkAccessToken(v string) *DismissGroupConversationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DismissGroupConversationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 14da****2760
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s DismissGroupConversationRequest) String() string {
	return tea.Prettify(s)
}

func (s DismissGroupConversationRequest) GoString() string {
	return s.String()
}

func (s *DismissGroupConversationRequest) SetOpenConversationId(v string) *DismissGroupConversationRequest {
	s.OpenConversationId = &v
	return s
}

type DismissGroupConversationResponseBody struct {
	// example:
	//
	// 14da****2760
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s DismissGroupConversationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DismissGroupConversationResponseBody) GoString() string {
	return s.String()
}

func (s *DismissGroupConversationResponseBody) SetOpenConversationId(v string) *DismissGroupConversationResponseBody {
	s.OpenConversationId = &v
	return s
}

type DismissGroupConversationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DismissGroupConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DismissGroupConversationResponse) String() string {
	return tea.Prettify(s)
}

func (s DismissGroupConversationResponse) GoString() string {
	return s.String()
}

func (s *DismissGroupConversationResponse) SetHeaders(v map[string]*string) *DismissGroupConversationResponse {
	s.Headers = v
	return s
}

func (s *DismissGroupConversationResponse) SetStatusCode(v int32) *DismissGroupConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *DismissGroupConversationResponse) SetBody(v *DismissGroupConversationResponseBody) *DismissGroupConversationResponse {
	s.Body = v
	return s
}

type FreezeGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s FreezeGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s FreezeGroupHeaders) GoString() string {
	return s.String()
}

func (s *FreezeGroupHeaders) SetCommonHeaders(v map[string]*string) *FreezeGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FreezeGroupHeaders) SetXAcsDingtalkAccessToken(v string) *FreezeGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type FreezeGroupRequest struct {
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s FreezeGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s FreezeGroupRequest) GoString() string {
	return s.String()
}

func (s *FreezeGroupRequest) SetOpenConversationId(v string) *FreezeGroupRequest {
	s.OpenConversationId = &v
	return s
}

type FreezeGroupResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s FreezeGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FreezeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *FreezeGroupResponseBody) SetSuccess(v bool) *FreezeGroupResponseBody {
	s.Success = &v
	return s
}

type FreezeGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FreezeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FreezeGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s FreezeGroupResponse) GoString() string {
	return s.String()
}

func (s *FreezeGroupResponse) SetHeaders(v map[string]*string) *FreezeGroupResponse {
	s.Headers = v
	return s
}

func (s *FreezeGroupResponse) SetStatusCode(v int32) *FreezeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *FreezeGroupResponse) SetBody(v *FreezeGroupResponseBody) *FreezeGroupResponse {
	s.Body = v
	return s
}

type GetConversationUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetConversationUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetConversationUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetConversationUrlHeaders) SetCommonHeaders(v map[string]*string) *GetConversationUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetConversationUrlHeaders) SetXAcsDingtalkAccessToken(v string) *GetConversationUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetConversationUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1107****2120
	AppUserId *string `json:"appUserId,omitempty" xml:"appUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oK4e****qER2
	ChannelCode *string `json:"channelCode,omitempty" xml:"channelCode,omitempty"`
	// example:
	//
	// 123**789
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// example:
	//
	// f67b****8a0f
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// 1745****8777
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetConversationUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConversationUrlRequest) GoString() string {
	return s.String()
}

func (s *GetConversationUrlRequest) SetAppUserId(v string) *GetConversationUrlRequest {
	s.AppUserId = &v
	return s
}

func (s *GetConversationUrlRequest) SetChannelCode(v string) *GetConversationUrlRequest {
	s.ChannelCode = &v
	return s
}

func (s *GetConversationUrlRequest) SetDeviceId(v string) *GetConversationUrlRequest {
	s.DeviceId = &v
	return s
}

func (s *GetConversationUrlRequest) SetOpenConversationId(v string) *GetConversationUrlRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GetConversationUrlRequest) SetUserId(v string) *GetConversationUrlRequest {
	s.UserId = &v
	return s
}

type GetConversationUrlResponseBody struct {
	// This parameter is required.
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetConversationUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConversationUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetConversationUrlResponseBody) SetUrl(v string) *GetConversationUrlResponseBody {
	s.Url = &v
	return s
}

type GetConversationUrlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConversationUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConversationUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConversationUrlResponse) GoString() string {
	return s.String()
}

func (s *GetConversationUrlResponse) SetHeaders(v map[string]*string) *GetConversationUrlResponse {
	s.Headers = v
	return s
}

func (s *GetConversationUrlResponse) SetStatusCode(v int32) *GetConversationUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConversationUrlResponse) SetBody(v *GetConversationUrlResponseBody) *GetConversationUrlResponse {
	s.Body = v
	return s
}

type GetFamilySchoolConversationMsgHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFamilySchoolConversationMsgHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFamilySchoolConversationMsgHeaders) GoString() string {
	return s.String()
}

func (s *GetFamilySchoolConversationMsgHeaders) SetCommonHeaders(v map[string]*string) *GetFamilySchoolConversationMsgHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFamilySchoolConversationMsgHeaders) SetXAcsDingtalkAccessToken(v string) *GetFamilySchoolConversationMsgHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFamilySchoolConversationMsgRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	MsgTypes []*int32 `json:"msgTypes,omitempty" xml:"msgTypes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1666671122000
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cidxxxx
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetFamilySchoolConversationMsgRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFamilySchoolConversationMsgRequest) GoString() string {
	return s.String()
}

func (s *GetFamilySchoolConversationMsgRequest) SetMaxResults(v int32) *GetFamilySchoolConversationMsgRequest {
	s.MaxResults = &v
	return s
}

func (s *GetFamilySchoolConversationMsgRequest) SetMsgTypes(v []*int32) *GetFamilySchoolConversationMsgRequest {
	s.MsgTypes = v
	return s
}

func (s *GetFamilySchoolConversationMsgRequest) SetNextToken(v int64) *GetFamilySchoolConversationMsgRequest {
	s.NextToken = &v
	return s
}

func (s *GetFamilySchoolConversationMsgRequest) SetOpenConversationId(v string) *GetFamilySchoolConversationMsgRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GetFamilySchoolConversationMsgRequest) SetUnionId(v string) *GetFamilySchoolConversationMsgRequest {
	s.UnionId = &v
	return s
}

type GetFamilySchoolConversationMsgResponseBody struct {
	// example:
	//
	// corp123
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// false
	HasMore  *string                                               `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	Messages []*GetFamilySchoolConversationMsgResponseBodyMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	// example:
	//
	// 1666671122000
	NextToken          *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s GetFamilySchoolConversationMsgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFamilySchoolConversationMsgResponseBody) GoString() string {
	return s.String()
}

func (s *GetFamilySchoolConversationMsgResponseBody) SetCorpId(v string) *GetFamilySchoolConversationMsgResponseBody {
	s.CorpId = &v
	return s
}

func (s *GetFamilySchoolConversationMsgResponseBody) SetHasMore(v string) *GetFamilySchoolConversationMsgResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetFamilySchoolConversationMsgResponseBody) SetMessages(v []*GetFamilySchoolConversationMsgResponseBodyMessages) *GetFamilySchoolConversationMsgResponseBody {
	s.Messages = v
	return s
}

func (s *GetFamilySchoolConversationMsgResponseBody) SetNextToken(v string) *GetFamilySchoolConversationMsgResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetFamilySchoolConversationMsgResponseBody) SetOpenConversationId(v string) *GetFamilySchoolConversationMsgResponseBody {
	s.OpenConversationId = &v
	return s
}

type GetFamilySchoolConversationMsgResponseBodyMessages struct {
	ContentType *int32                                                           `json:"contentType,omitempty" xml:"contentType,omitempty"`
	CreateAt    *int64                                                           `json:"createAt,omitempty" xml:"createAt,omitempty"`
	MediaModels []*GetFamilySchoolConversationMsgResponseBodyMessagesMediaModels `json:"mediaModels,omitempty" xml:"mediaModels,omitempty" type:"Repeated"`
	// example:
	//
	// msgxxx
	OpenMsgId *string `json:"openMsgId,omitempty" xml:"openMsgId,omitempty"`
}

func (s GetFamilySchoolConversationMsgResponseBodyMessages) String() string {
	return tea.Prettify(s)
}

func (s GetFamilySchoolConversationMsgResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *GetFamilySchoolConversationMsgResponseBodyMessages) SetContentType(v int32) *GetFamilySchoolConversationMsgResponseBodyMessages {
	s.ContentType = &v
	return s
}

func (s *GetFamilySchoolConversationMsgResponseBodyMessages) SetCreateAt(v int64) *GetFamilySchoolConversationMsgResponseBodyMessages {
	s.CreateAt = &v
	return s
}

func (s *GetFamilySchoolConversationMsgResponseBodyMessages) SetMediaModels(v []*GetFamilySchoolConversationMsgResponseBodyMessagesMediaModels) *GetFamilySchoolConversationMsgResponseBodyMessages {
	s.MediaModels = v
	return s
}

func (s *GetFamilySchoolConversationMsgResponseBodyMessages) SetOpenMsgId(v string) *GetFamilySchoolConversationMsgResponseBodyMessages {
	s.OpenMsgId = &v
	return s
}

type GetFamilySchoolConversationMsgResponseBodyMessagesMediaModels struct {
	// example:
	//
	// aa.png
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// png
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// @12xxx34
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	// example:
	//
	// 1234
	Size *string `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// https://wukong-xxxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// example:
	//
	// @12xx34
	VideoPicMediaId *string `json:"videoPicMediaId,omitempty" xml:"videoPicMediaId,omitempty"`
}

func (s GetFamilySchoolConversationMsgResponseBodyMessagesMediaModels) String() string {
	return tea.Prettify(s)
}

func (s GetFamilySchoolConversationMsgResponseBodyMessagesMediaModels) GoString() string {
	return s.String()
}

func (s *GetFamilySchoolConversationMsgResponseBodyMessagesMediaModels) SetFileName(v string) *GetFamilySchoolConversationMsgResponseBodyMessagesMediaModels {
	s.FileName = &v
	return s
}

func (s *GetFamilySchoolConversationMsgResponseBodyMessagesMediaModels) SetFileType(v string) *GetFamilySchoolConversationMsgResponseBodyMessagesMediaModels {
	s.FileType = &v
	return s
}

func (s *GetFamilySchoolConversationMsgResponseBodyMessagesMediaModels) SetMediaId(v string) *GetFamilySchoolConversationMsgResponseBodyMessagesMediaModels {
	s.MediaId = &v
	return s
}

func (s *GetFamilySchoolConversationMsgResponseBodyMessagesMediaModels) SetSize(v string) *GetFamilySchoolConversationMsgResponseBodyMessagesMediaModels {
	s.Size = &v
	return s
}

func (s *GetFamilySchoolConversationMsgResponseBodyMessagesMediaModels) SetUrl(v string) *GetFamilySchoolConversationMsgResponseBodyMessagesMediaModels {
	s.Url = &v
	return s
}

func (s *GetFamilySchoolConversationMsgResponseBodyMessagesMediaModels) SetVideoPicMediaId(v string) *GetFamilySchoolConversationMsgResponseBodyMessagesMediaModels {
	s.VideoPicMediaId = &v
	return s
}

type GetFamilySchoolConversationMsgResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFamilySchoolConversationMsgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFamilySchoolConversationMsgResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFamilySchoolConversationMsgResponse) GoString() string {
	return s.String()
}

func (s *GetFamilySchoolConversationMsgResponse) SetHeaders(v map[string]*string) *GetFamilySchoolConversationMsgResponse {
	s.Headers = v
	return s
}

func (s *GetFamilySchoolConversationMsgResponse) SetStatusCode(v int32) *GetFamilySchoolConversationMsgResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFamilySchoolConversationMsgResponse) SetBody(v *GetFamilySchoolConversationMsgResponseBody) *GetFamilySchoolConversationMsgResponse {
	s.Body = v
	return s
}

type GetFamilySchoolConversationsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFamilySchoolConversationsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFamilySchoolConversationsHeaders) GoString() string {
	return s.String()
}

func (s *GetFamilySchoolConversationsHeaders) SetCommonHeaders(v map[string]*string) *GetFamilySchoolConversationsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFamilySchoolConversationsHeaders) SetXAcsDingtalkAccessToken(v string) *GetFamilySchoolConversationsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFamilySchoolConversationsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s GetFamilySchoolConversationsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFamilySchoolConversationsRequest) GoString() string {
	return s.String()
}

func (s *GetFamilySchoolConversationsRequest) SetMaxResults(v int32) *GetFamilySchoolConversationsRequest {
	s.MaxResults = &v
	return s
}

func (s *GetFamilySchoolConversationsRequest) SetNextToken(v int64) *GetFamilySchoolConversationsRequest {
	s.NextToken = &v
	return s
}

func (s *GetFamilySchoolConversationsRequest) SetUnionId(v string) *GetFamilySchoolConversationsRequest {
	s.UnionId = &v
	return s
}

type GetFamilySchoolConversationsResponseBody struct {
	GroupInfoList []*GetFamilySchoolConversationsResponseBodyGroupInfoList `json:"groupInfoList,omitempty" xml:"groupInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// false
	HasMore *string `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 1666671122000
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetFamilySchoolConversationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFamilySchoolConversationsResponseBody) GoString() string {
	return s.String()
}

func (s *GetFamilySchoolConversationsResponseBody) SetGroupInfoList(v []*GetFamilySchoolConversationsResponseBodyGroupInfoList) *GetFamilySchoolConversationsResponseBody {
	s.GroupInfoList = v
	return s
}

func (s *GetFamilySchoolConversationsResponseBody) SetHasMore(v string) *GetFamilySchoolConversationsResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetFamilySchoolConversationsResponseBody) SetNextToken(v string) *GetFamilySchoolConversationsResponseBody {
	s.NextToken = &v
	return s
}

type GetFamilySchoolConversationsResponseBodyGroupInfoList struct {
	// example:
	//
	// corp123
	CorpId        *string   `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeptNameChain []*string `json:"deptNameChain,omitempty" xml:"deptNameChain,omitempty" type:"Repeated"`
	// example:
	//
	// 小王的家校群
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// example:
	//
	// 2
	GroupType     *string `json:"groupType,omitempty" xml:"groupType,omitempty"`
	JoinGroupTime *int64  `json:"joinGroupTime,omitempty" xml:"joinGroupTime,omitempty"`
	// example:
	//
	// cidxxx
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s GetFamilySchoolConversationsResponseBodyGroupInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetFamilySchoolConversationsResponseBodyGroupInfoList) GoString() string {
	return s.String()
}

func (s *GetFamilySchoolConversationsResponseBodyGroupInfoList) SetCorpId(v string) *GetFamilySchoolConversationsResponseBodyGroupInfoList {
	s.CorpId = &v
	return s
}

func (s *GetFamilySchoolConversationsResponseBodyGroupInfoList) SetDeptNameChain(v []*string) *GetFamilySchoolConversationsResponseBodyGroupInfoList {
	s.DeptNameChain = v
	return s
}

func (s *GetFamilySchoolConversationsResponseBodyGroupInfoList) SetGroupName(v string) *GetFamilySchoolConversationsResponseBodyGroupInfoList {
	s.GroupName = &v
	return s
}

func (s *GetFamilySchoolConversationsResponseBodyGroupInfoList) SetGroupType(v string) *GetFamilySchoolConversationsResponseBodyGroupInfoList {
	s.GroupType = &v
	return s
}

func (s *GetFamilySchoolConversationsResponseBodyGroupInfoList) SetJoinGroupTime(v int64) *GetFamilySchoolConversationsResponseBodyGroupInfoList {
	s.JoinGroupTime = &v
	return s
}

func (s *GetFamilySchoolConversationsResponseBodyGroupInfoList) SetOpenConversationId(v string) *GetFamilySchoolConversationsResponseBodyGroupInfoList {
	s.OpenConversationId = &v
	return s
}

type GetFamilySchoolConversationsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFamilySchoolConversationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFamilySchoolConversationsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFamilySchoolConversationsResponse) GoString() string {
	return s.String()
}

func (s *GetFamilySchoolConversationsResponse) SetHeaders(v map[string]*string) *GetFamilySchoolConversationsResponse {
	s.Headers = v
	return s
}

func (s *GetFamilySchoolConversationsResponse) SetStatusCode(v int32) *GetFamilySchoolConversationsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFamilySchoolConversationsResponse) SetBody(v *GetFamilySchoolConversationsResponseBody) *GetFamilySchoolConversationsResponse {
	s.Body = v
	return s
}

type GetInnerGroupMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetInnerGroupMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetInnerGroupMembersHeaders) GoString() string {
	return s.String()
}

func (s *GetInnerGroupMembersHeaders) SetCommonHeaders(v map[string]*string) *GetInnerGroupMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInnerGroupMembersHeaders) SetXAcsDingtalkAccessToken(v string) *GetInnerGroupMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetInnerGroupMembersRequest struct {
	// This parameter is required.
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// UZr*****
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cid1e*****==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 015*****
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetInnerGroupMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInnerGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *GetInnerGroupMembersRequest) SetMaxResults(v int32) *GetInnerGroupMembersRequest {
	s.MaxResults = &v
	return s
}

func (s *GetInnerGroupMembersRequest) SetNextToken(v string) *GetInnerGroupMembersRequest {
	s.NextToken = &v
	return s
}

func (s *GetInnerGroupMembersRequest) SetOpenConversationId(v string) *GetInnerGroupMembersRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GetInnerGroupMembersRequest) SetUserId(v string) *GetInnerGroupMembersRequest {
	s.UserId = &v
	return s
}

type GetInnerGroupMembersResponseBody struct {
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// UZr*****
	NextToken *string   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	UserIds   []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s GetInnerGroupMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInnerGroupMembersResponseBody) GoString() string {
	return s.String()
}

func (s *GetInnerGroupMembersResponseBody) SetHasMore(v bool) *GetInnerGroupMembersResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetInnerGroupMembersResponseBody) SetNextToken(v string) *GetInnerGroupMembersResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetInnerGroupMembersResponseBody) SetUserIds(v []*string) *GetInnerGroupMembersResponseBody {
	s.UserIds = v
	return s
}

type GetInnerGroupMembersResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInnerGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInnerGroupMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInnerGroupMembersResponse) GoString() string {
	return s.String()
}

func (s *GetInnerGroupMembersResponse) SetHeaders(v map[string]*string) *GetInnerGroupMembersResponse {
	s.Headers = v
	return s
}

func (s *GetInnerGroupMembersResponse) SetStatusCode(v int32) *GetInnerGroupMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInnerGroupMembersResponse) SetBody(v *GetInnerGroupMembersResponseBody) *GetInnerGroupMembersResponse {
	s.Body = v
	return s
}

type GetInterconnectionUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetInterconnectionUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetInterconnectionUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetInterconnectionUrlHeaders) SetCommonHeaders(v map[string]*string) *GetInterconnectionUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInterconnectionUrlHeaders) SetXAcsDingtalkAccessToken(v string) *GetInterconnectionUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetInterconnectionUrlRequest struct {
	AppUserAvatar     *string `json:"appUserAvatar,omitempty" xml:"appUserAvatar,omitempty"`
	AppUserAvatarType *int32  `json:"appUserAvatarType,omitempty" xml:"appUserAvatarType,omitempty"`
	// This parameter is required.
	AppUserId *string `json:"appUserId,omitempty" xml:"appUserId,omitempty"`
	// This parameter is required.
	AppUserMobileNumber *string `json:"appUserMobileNumber,omitempty" xml:"appUserMobileNumber,omitempty"`
	// This parameter is required.
	AppUserName *string `json:"appUserName,omitempty" xml:"appUserName,omitempty"`
	// This parameter is required.
	MsgPageType *int32  `json:"msgPageType,omitempty" xml:"msgPageType,omitempty"`
	QrCode      *string `json:"qrCode,omitempty" xml:"qrCode,omitempty"`
	// This parameter is required.
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	// This parameter is required.
	SourceCode *string `json:"sourceCode,omitempty" xml:"sourceCode,omitempty"`
	// This parameter is required.
	SourceType *int32  `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	UserId     *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetInterconnectionUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInterconnectionUrlRequest) GoString() string {
	return s.String()
}

func (s *GetInterconnectionUrlRequest) SetAppUserAvatar(v string) *GetInterconnectionUrlRequest {
	s.AppUserAvatar = &v
	return s
}

func (s *GetInterconnectionUrlRequest) SetAppUserAvatarType(v int32) *GetInterconnectionUrlRequest {
	s.AppUserAvatarType = &v
	return s
}

func (s *GetInterconnectionUrlRequest) SetAppUserId(v string) *GetInterconnectionUrlRequest {
	s.AppUserId = &v
	return s
}

func (s *GetInterconnectionUrlRequest) SetAppUserMobileNumber(v string) *GetInterconnectionUrlRequest {
	s.AppUserMobileNumber = &v
	return s
}

func (s *GetInterconnectionUrlRequest) SetAppUserName(v string) *GetInterconnectionUrlRequest {
	s.AppUserName = &v
	return s
}

func (s *GetInterconnectionUrlRequest) SetMsgPageType(v int32) *GetInterconnectionUrlRequest {
	s.MsgPageType = &v
	return s
}

func (s *GetInterconnectionUrlRequest) SetQrCode(v string) *GetInterconnectionUrlRequest {
	s.QrCode = &v
	return s
}

func (s *GetInterconnectionUrlRequest) SetSignature(v string) *GetInterconnectionUrlRequest {
	s.Signature = &v
	return s
}

func (s *GetInterconnectionUrlRequest) SetSourceCode(v string) *GetInterconnectionUrlRequest {
	s.SourceCode = &v
	return s
}

func (s *GetInterconnectionUrlRequest) SetSourceType(v int32) *GetInterconnectionUrlRequest {
	s.SourceType = &v
	return s
}

func (s *GetInterconnectionUrlRequest) SetUserId(v string) *GetInterconnectionUrlRequest {
	s.UserId = &v
	return s
}

type GetInterconnectionUrlResponseBody struct {
	// This parameter is required.
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetInterconnectionUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInterconnectionUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetInterconnectionUrlResponseBody) SetUrl(v string) *GetInterconnectionUrlResponseBody {
	s.Url = &v
	return s
}

type GetInterconnectionUrlResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInterconnectionUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInterconnectionUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInterconnectionUrlResponse) GoString() string {
	return s.String()
}

func (s *GetInterconnectionUrlResponse) SetHeaders(v map[string]*string) *GetInterconnectionUrlResponse {
	s.Headers = v
	return s
}

func (s *GetInterconnectionUrlResponse) SetStatusCode(v int32) *GetInterconnectionUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInterconnectionUrlResponse) SetBody(v *GetInterconnectionUrlResponseBody) *GetInterconnectionUrlResponse {
	s.Body = v
	return s
}

type GetNewestInnerGroupsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetNewestInnerGroupsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetNewestInnerGroupsHeaders) GoString() string {
	return s.String()
}

func (s *GetNewestInnerGroupsHeaders) SetCommonHeaders(v map[string]*string) *GetNewestInnerGroupsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetNewestInnerGroupsHeaders) SetXAcsDingtalkAccessToken(v string) *GetNewestInnerGroupsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetNewestInnerGroupsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 015*****
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetNewestInnerGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNewestInnerGroupsRequest) GoString() string {
	return s.String()
}

func (s *GetNewestInnerGroupsRequest) SetUserId(v string) *GetNewestInnerGroupsRequest {
	s.UserId = &v
	return s
}

type GetNewestInnerGroupsResponseBody struct {
	GroupInfos []*GetNewestInnerGroupsResponseBodyGroupInfos `json:"groupInfos,omitempty" xml:"groupInfos,omitempty" type:"Repeated"`
}

func (s GetNewestInnerGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNewestInnerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *GetNewestInnerGroupsResponseBody) SetGroupInfos(v []*GetNewestInnerGroupsResponseBodyGroupInfos) *GetNewestInnerGroupsResponseBody {
	s.GroupInfos = v
	return s
}

type GetNewestInnerGroupsResponseBodyGroupInfos struct {
	// example:
	//
	// @lADOADma*****QKA
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// 10
	MemberAmount *string `json:"memberAmount,omitempty" xml:"memberAmount,omitempty"`
	// example:
	//
	// cid1e*****==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// 测试群名称
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetNewestInnerGroupsResponseBodyGroupInfos) String() string {
	return tea.Prettify(s)
}

func (s GetNewestInnerGroupsResponseBodyGroupInfos) GoString() string {
	return s.String()
}

func (s *GetNewestInnerGroupsResponseBodyGroupInfos) SetIcon(v string) *GetNewestInnerGroupsResponseBodyGroupInfos {
	s.Icon = &v
	return s
}

func (s *GetNewestInnerGroupsResponseBodyGroupInfos) SetMemberAmount(v string) *GetNewestInnerGroupsResponseBodyGroupInfos {
	s.MemberAmount = &v
	return s
}

func (s *GetNewestInnerGroupsResponseBodyGroupInfos) SetOpenConversationId(v string) *GetNewestInnerGroupsResponseBodyGroupInfos {
	s.OpenConversationId = &v
	return s
}

func (s *GetNewestInnerGroupsResponseBodyGroupInfos) SetTitle(v string) *GetNewestInnerGroupsResponseBodyGroupInfos {
	s.Title = &v
	return s
}

type GetNewestInnerGroupsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNewestInnerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNewestInnerGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNewestInnerGroupsResponse) GoString() string {
	return s.String()
}

func (s *GetNewestInnerGroupsResponse) SetHeaders(v map[string]*string) *GetNewestInnerGroupsResponse {
	s.Headers = v
	return s
}

func (s *GetNewestInnerGroupsResponse) SetStatusCode(v int32) *GetNewestInnerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNewestInnerGroupsResponse) SetBody(v *GetNewestInnerGroupsResponseBody) *GetNewestInnerGroupsResponse {
	s.Body = v
	return s
}

type GetSceneGroupInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSceneGroupInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSceneGroupInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetSceneGroupInfoHeaders) SetCommonHeaders(v map[string]*string) *GetSceneGroupInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSceneGroupInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetSceneGroupInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSceneGroupInfoRequest struct {
	CoolAppCode *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cidXXXXXXX
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s GetSceneGroupInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSceneGroupInfoRequest) GoString() string {
	return s.String()
}

func (s *GetSceneGroupInfoRequest) SetCoolAppCode(v string) *GetSceneGroupInfoRequest {
	s.CoolAppCode = &v
	return s
}

func (s *GetSceneGroupInfoRequest) SetOpenConversationId(v string) *GetSceneGroupInfoRequest {
	s.OpenConversationId = &v
	return s
}

type GetSceneGroupInfoResponseBody struct {
	GroupUrl *string `json:"groupUrl,omitempty" xml:"groupUrl,omitempty"`
	Icon     *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// cidXXXXXXXXX==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	OwnerUserId        *string `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	Status             *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Success            *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TemplateId         *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	Title              *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetSceneGroupInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSceneGroupInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetSceneGroupInfoResponseBody) SetGroupUrl(v string) *GetSceneGroupInfoResponseBody {
	s.GroupUrl = &v
	return s
}

func (s *GetSceneGroupInfoResponseBody) SetIcon(v string) *GetSceneGroupInfoResponseBody {
	s.Icon = &v
	return s
}

func (s *GetSceneGroupInfoResponseBody) SetOpenConversationId(v string) *GetSceneGroupInfoResponseBody {
	s.OpenConversationId = &v
	return s
}

func (s *GetSceneGroupInfoResponseBody) SetOwnerUserId(v string) *GetSceneGroupInfoResponseBody {
	s.OwnerUserId = &v
	return s
}

func (s *GetSceneGroupInfoResponseBody) SetStatus(v int32) *GetSceneGroupInfoResponseBody {
	s.Status = &v
	return s
}

func (s *GetSceneGroupInfoResponseBody) SetSuccess(v bool) *GetSceneGroupInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetSceneGroupInfoResponseBody) SetTemplateId(v string) *GetSceneGroupInfoResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetSceneGroupInfoResponseBody) SetTitle(v string) *GetSceneGroupInfoResponseBody {
	s.Title = &v
	return s
}

type GetSceneGroupInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSceneGroupInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSceneGroupInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSceneGroupInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSceneGroupInfoResponse) SetHeaders(v map[string]*string) *GetSceneGroupInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSceneGroupInfoResponse) SetStatusCode(v int32) *GetSceneGroupInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSceneGroupInfoResponse) SetBody(v *GetSceneGroupInfoResponseBody) *GetSceneGroupInfoResponse {
	s.Body = v
	return s
}

type GetSceneGroupMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSceneGroupMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSceneGroupMembersHeaders) GoString() string {
	return s.String()
}

func (s *GetSceneGroupMembersHeaders) SetCommonHeaders(v map[string]*string) *GetSceneGroupMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSceneGroupMembersHeaders) SetXAcsDingtalkAccessToken(v string) *GetSceneGroupMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSceneGroupMembersRequest struct {
	// This parameter is required.
	CoolAppCode *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	Cursor *string `json:"cursor,omitempty" xml:"cursor,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cidXXXXXXX
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 200
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s GetSceneGroupMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSceneGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *GetSceneGroupMembersRequest) SetCoolAppCode(v string) *GetSceneGroupMembersRequest {
	s.CoolAppCode = &v
	return s
}

func (s *GetSceneGroupMembersRequest) SetCursor(v string) *GetSceneGroupMembersRequest {
	s.Cursor = &v
	return s
}

func (s *GetSceneGroupMembersRequest) SetOpenConversationId(v string) *GetSceneGroupMembersRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GetSceneGroupMembersRequest) SetSize(v int64) *GetSceneGroupMembersRequest {
	s.Size = &v
	return s
}

type GetSceneGroupMembersResponseBody struct {
	// example:
	//
	// false
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// cidXXXXXXXXX==
	MemberUserIds []*string `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// 92233720368
	NextCursor *string `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	Success    *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetSceneGroupMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSceneGroupMembersResponseBody) GoString() string {
	return s.String()
}

func (s *GetSceneGroupMembersResponseBody) SetHasMore(v bool) *GetSceneGroupMembersResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetSceneGroupMembersResponseBody) SetMemberUserIds(v []*string) *GetSceneGroupMembersResponseBody {
	s.MemberUserIds = v
	return s
}

func (s *GetSceneGroupMembersResponseBody) SetNextCursor(v string) *GetSceneGroupMembersResponseBody {
	s.NextCursor = &v
	return s
}

func (s *GetSceneGroupMembersResponseBody) SetSuccess(v bool) *GetSceneGroupMembersResponseBody {
	s.Success = &v
	return s
}

type GetSceneGroupMembersResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSceneGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSceneGroupMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSceneGroupMembersResponse) GoString() string {
	return s.String()
}

func (s *GetSceneGroupMembersResponse) SetHeaders(v map[string]*string) *GetSceneGroupMembersResponse {
	s.Headers = v
	return s
}

func (s *GetSceneGroupMembersResponse) SetStatusCode(v int32) *GetSceneGroupMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSceneGroupMembersResponse) SetBody(v *GetSceneGroupMembersResponseBody) *GetSceneGroupMembersResponse {
	s.Body = v
	return s
}

type GetSceneGroupTemplateMessageOpenStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSceneGroupTemplateMessageOpenStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSceneGroupTemplateMessageOpenStatusHeaders) GoString() string {
	return s.String()
}

func (s *GetSceneGroupTemplateMessageOpenStatusHeaders) SetCommonHeaders(v map[string]*string) *GetSceneGroupTemplateMessageOpenStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSceneGroupTemplateMessageOpenStatusHeaders) SetXAcsDingtalkAccessToken(v string) *GetSceneGroupTemplateMessageOpenStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSceneGroupTemplateMessageOpenStatusResponseBody struct {
	Status  *int32 `json:"status,omitempty" xml:"status,omitempty"`
	Success *bool  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetSceneGroupTemplateMessageOpenStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSceneGroupTemplateMessageOpenStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetSceneGroupTemplateMessageOpenStatusResponseBody) SetStatus(v int32) *GetSceneGroupTemplateMessageOpenStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetSceneGroupTemplateMessageOpenStatusResponseBody) SetSuccess(v bool) *GetSceneGroupTemplateMessageOpenStatusResponseBody {
	s.Success = &v
	return s
}

type GetSceneGroupTemplateMessageOpenStatusResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSceneGroupTemplateMessageOpenStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSceneGroupTemplateMessageOpenStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSceneGroupTemplateMessageOpenStatusResponse) GoString() string {
	return s.String()
}

func (s *GetSceneGroupTemplateMessageOpenStatusResponse) SetHeaders(v map[string]*string) *GetSceneGroupTemplateMessageOpenStatusResponse {
	s.Headers = v
	return s
}

func (s *GetSceneGroupTemplateMessageOpenStatusResponse) SetStatusCode(v int32) *GetSceneGroupTemplateMessageOpenStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSceneGroupTemplateMessageOpenStatusResponse) SetBody(v *GetSceneGroupTemplateMessageOpenStatusResponseBody) *GetSceneGroupTemplateMessageOpenStatusResponse {
	s.Body = v
	return s
}

type GetSingleChatOpenConversationIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSingleChatOpenConversationIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSingleChatOpenConversationIdHeaders) GoString() string {
	return s.String()
}

func (s *GetSingleChatOpenConversationIdHeaders) SetCommonHeaders(v map[string]*string) *GetSingleChatOpenConversationIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSingleChatOpenConversationIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetSingleChatOpenConversationIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSingleChatOpenConversationIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 022*****2134
	UserId1 *string `json:"userId1,omitempty" xml:"userId1,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 072*****1243
	UserId2 *string `json:"userId2,omitempty" xml:"userId2,omitempty"`
}

func (s GetSingleChatOpenConversationIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSingleChatOpenConversationIdRequest) GoString() string {
	return s.String()
}

func (s *GetSingleChatOpenConversationIdRequest) SetUserId1(v string) *GetSingleChatOpenConversationIdRequest {
	s.UserId1 = &v
	return s
}

func (s *GetSingleChatOpenConversationIdRequest) SetUserId2(v string) *GetSingleChatOpenConversationIdRequest {
	s.UserId2 = &v
	return s
}

type GetSingleChatOpenConversationIdResponseBody struct {
	Result  *GetSingleChatOpenConversationIdResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *string                                            `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetSingleChatOpenConversationIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSingleChatOpenConversationIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetSingleChatOpenConversationIdResponseBody) SetResult(v *GetSingleChatOpenConversationIdResponseBodyResult) *GetSingleChatOpenConversationIdResponseBody {
	s.Result = v
	return s
}

func (s *GetSingleChatOpenConversationIdResponseBody) SetSuccess(v string) *GetSingleChatOpenConversationIdResponseBody {
	s.Success = &v
	return s
}

type GetSingleChatOpenConversationIdResponseBodyResult struct {
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s GetSingleChatOpenConversationIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetSingleChatOpenConversationIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSingleChatOpenConversationIdResponseBodyResult) SetOpenConversationId(v string) *GetSingleChatOpenConversationIdResponseBodyResult {
	s.OpenConversationId = &v
	return s
}

type GetSingleChatOpenConversationIdResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSingleChatOpenConversationIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSingleChatOpenConversationIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSingleChatOpenConversationIdResponse) GoString() string {
	return s.String()
}

func (s *GetSingleChatOpenConversationIdResponse) SetHeaders(v map[string]*string) *GetSingleChatOpenConversationIdResponse {
	s.Headers = v
	return s
}

func (s *GetSingleChatOpenConversationIdResponse) SetStatusCode(v int32) *GetSingleChatOpenConversationIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSingleChatOpenConversationIdResponse) SetBody(v *GetSingleChatOpenConversationIdResponseBody) *GetSingleChatOpenConversationIdResponse {
	s.Body = v
	return s
}

type GetSuperAdminOpenSceneGroupInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSuperAdminOpenSceneGroupInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSuperAdminOpenSceneGroupInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetSuperAdminOpenSceneGroupInfoHeaders) SetCommonHeaders(v map[string]*string) *GetSuperAdminOpenSceneGroupInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSuperAdminOpenSceneGroupInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetSuperAdminOpenSceneGroupInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSuperAdminOpenSceneGroupInfoRequest struct {
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s GetSuperAdminOpenSceneGroupInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSuperAdminOpenSceneGroupInfoRequest) GoString() string {
	return s.String()
}

func (s *GetSuperAdminOpenSceneGroupInfoRequest) SetOpenConversationId(v string) *GetSuperAdminOpenSceneGroupInfoRequest {
	s.OpenConversationId = &v
	return s
}

type GetSuperAdminOpenSceneGroupInfoResponseBody struct {
	GroupUrl           *string                                                       `json:"groupUrl,omitempty" xml:"groupUrl,omitempty"`
	Icon               *string                                                       `json:"icon,omitempty" xml:"icon,omitempty"`
	ManagementOptions  *GetSuperAdminOpenSceneGroupInfoResponseBodyManagementOptions `json:"managementOptions,omitempty" xml:"managementOptions,omitempty" type:"Struct"`
	OpenConversationId *string                                                       `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	OwnerUserId        *string                                                       `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	SubAdminUserIds    []*string                                                     `json:"subAdminUserIds,omitempty" xml:"subAdminUserIds,omitempty" type:"Repeated"`
	Sucess             *bool                                                         `json:"sucess,omitempty" xml:"sucess,omitempty"`
	TemplateId         *string                                                       `json:"templateId,omitempty" xml:"templateId,omitempty"`
	Title              *string                                                       `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetSuperAdminOpenSceneGroupInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSuperAdminOpenSceneGroupInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetSuperAdminOpenSceneGroupInfoResponseBody) SetGroupUrl(v string) *GetSuperAdminOpenSceneGroupInfoResponseBody {
	s.GroupUrl = &v
	return s
}

func (s *GetSuperAdminOpenSceneGroupInfoResponseBody) SetIcon(v string) *GetSuperAdminOpenSceneGroupInfoResponseBody {
	s.Icon = &v
	return s
}

func (s *GetSuperAdminOpenSceneGroupInfoResponseBody) SetManagementOptions(v *GetSuperAdminOpenSceneGroupInfoResponseBodyManagementOptions) *GetSuperAdminOpenSceneGroupInfoResponseBody {
	s.ManagementOptions = v
	return s
}

func (s *GetSuperAdminOpenSceneGroupInfoResponseBody) SetOpenConversationId(v string) *GetSuperAdminOpenSceneGroupInfoResponseBody {
	s.OpenConversationId = &v
	return s
}

func (s *GetSuperAdminOpenSceneGroupInfoResponseBody) SetOwnerUserId(v string) *GetSuperAdminOpenSceneGroupInfoResponseBody {
	s.OwnerUserId = &v
	return s
}

func (s *GetSuperAdminOpenSceneGroupInfoResponseBody) SetSubAdminUserIds(v []*string) *GetSuperAdminOpenSceneGroupInfoResponseBody {
	s.SubAdminUserIds = v
	return s
}

func (s *GetSuperAdminOpenSceneGroupInfoResponseBody) SetSucess(v bool) *GetSuperAdminOpenSceneGroupInfoResponseBody {
	s.Sucess = &v
	return s
}

func (s *GetSuperAdminOpenSceneGroupInfoResponseBody) SetTemplateId(v string) *GetSuperAdminOpenSceneGroupInfoResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetSuperAdminOpenSceneGroupInfoResponseBody) SetTitle(v string) *GetSuperAdminOpenSceneGroupInfoResponseBody {
	s.Title = &v
	return s
}

type GetSuperAdminOpenSceneGroupInfoResponseBodyManagementOptions struct {
	ChatBannedType      *string `json:"chatBannedType,omitempty" xml:"chatBannedType,omitempty"`
	ManagementType      *string `json:"managementType,omitempty" xml:"managementType,omitempty"`
	MentionAllAuthority *string `json:"mentionAllAuthority,omitempty" xml:"mentionAllAuthority,omitempty"`
	Searchable          *string `json:"searchable,omitempty" xml:"searchable,omitempty"`
	ShowHistoryType     *string `json:"showHistoryType,omitempty" xml:"showHistoryType,omitempty"`
	ValidationType      *string `json:"validationType,omitempty" xml:"validationType,omitempty"`
}

func (s GetSuperAdminOpenSceneGroupInfoResponseBodyManagementOptions) String() string {
	return tea.Prettify(s)
}

func (s GetSuperAdminOpenSceneGroupInfoResponseBodyManagementOptions) GoString() string {
	return s.String()
}

func (s *GetSuperAdminOpenSceneGroupInfoResponseBodyManagementOptions) SetChatBannedType(v string) *GetSuperAdminOpenSceneGroupInfoResponseBodyManagementOptions {
	s.ChatBannedType = &v
	return s
}

func (s *GetSuperAdminOpenSceneGroupInfoResponseBodyManagementOptions) SetManagementType(v string) *GetSuperAdminOpenSceneGroupInfoResponseBodyManagementOptions {
	s.ManagementType = &v
	return s
}

func (s *GetSuperAdminOpenSceneGroupInfoResponseBodyManagementOptions) SetMentionAllAuthority(v string) *GetSuperAdminOpenSceneGroupInfoResponseBodyManagementOptions {
	s.MentionAllAuthority = &v
	return s
}

func (s *GetSuperAdminOpenSceneGroupInfoResponseBodyManagementOptions) SetSearchable(v string) *GetSuperAdminOpenSceneGroupInfoResponseBodyManagementOptions {
	s.Searchable = &v
	return s
}

func (s *GetSuperAdminOpenSceneGroupInfoResponseBodyManagementOptions) SetShowHistoryType(v string) *GetSuperAdminOpenSceneGroupInfoResponseBodyManagementOptions {
	s.ShowHistoryType = &v
	return s
}

func (s *GetSuperAdminOpenSceneGroupInfoResponseBodyManagementOptions) SetValidationType(v string) *GetSuperAdminOpenSceneGroupInfoResponseBodyManagementOptions {
	s.ValidationType = &v
	return s
}

type GetSuperAdminOpenSceneGroupInfoResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSuperAdminOpenSceneGroupInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSuperAdminOpenSceneGroupInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSuperAdminOpenSceneGroupInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSuperAdminOpenSceneGroupInfoResponse) SetHeaders(v map[string]*string) *GetSuperAdminOpenSceneGroupInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSuperAdminOpenSceneGroupInfoResponse) SetStatusCode(v int32) *GetSuperAdminOpenSceneGroupInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSuperAdminOpenSceneGroupInfoResponse) SetBody(v *GetSuperAdminOpenSceneGroupInfoResponseBody) *GetSuperAdminOpenSceneGroupInfoResponse {
	s.Body = v
	return s
}

type GroupBanWordsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GroupBanWordsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GroupBanWordsHeaders) GoString() string {
	return s.String()
}

func (s *GroupBanWordsHeaders) SetCommonHeaders(v map[string]*string) *GroupBanWordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GroupBanWordsHeaders) SetXAcsDingtalkAccessToken(v string) *GroupBanWordsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GroupBanWordsRequest struct {
	// example:
	//
	// 1
	BanWordsMode *int32 `json:"banWordsMode,omitempty" xml:"banWordsMode,omitempty"`
	// example:
	//
	// cidnvcxzklxv
	OpenConversationId *string                `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	Options            map[string]interface{} `json:"options,omitempty" xml:"options,omitempty"`
}

func (s GroupBanWordsRequest) String() string {
	return tea.Prettify(s)
}

func (s GroupBanWordsRequest) GoString() string {
	return s.String()
}

func (s *GroupBanWordsRequest) SetBanWordsMode(v int32) *GroupBanWordsRequest {
	s.BanWordsMode = &v
	return s
}

func (s *GroupBanWordsRequest) SetOpenConversationId(v string) *GroupBanWordsRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GroupBanWordsRequest) SetOptions(v map[string]interface{}) *GroupBanWordsRequest {
	s.Options = v
	return s
}

type GroupBanWordsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s GroupBanWordsResponse) String() string {
	return tea.Prettify(s)
}

func (s GroupBanWordsResponse) GoString() string {
	return s.String()
}

func (s *GroupBanWordsResponse) SetHeaders(v map[string]*string) *GroupBanWordsResponse {
	s.Headers = v
	return s
}

func (s *GroupBanWordsResponse) SetStatusCode(v int32) *GroupBanWordsResponse {
	s.StatusCode = &v
	return s
}

type GroupCapacityInquiryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GroupCapacityInquiryHeaders) String() string {
	return tea.Prettify(s)
}

func (s GroupCapacityInquiryHeaders) GoString() string {
	return s.String()
}

func (s *GroupCapacityInquiryHeaders) SetCommonHeaders(v map[string]*string) *GroupCapacityInquiryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GroupCapacityInquiryHeaders) SetXAcsDingtalkAccessToken(v string) *GroupCapacityInquiryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GroupCapacityInquiryRequest struct {
	// example:
	//
	// 1Y
	EffectiveDuration *string `json:"effectiveDuration,omitempty" xml:"effectiveDuration,omitempty"`
	// example:
	//
	// ciddmslasdfxcvbxcvgidnxsd==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// 5782431748978293
	Operator *string                `json:"operator,omitempty" xml:"operator,omitempty"`
	Options  map[string]interface{} `json:"options,omitempty" xml:"options,omitempty"`
	// example:
	//
	// 2000
	TargetCapacity *int32 `json:"targetCapacity,omitempty" xml:"targetCapacity,omitempty"`
}

func (s GroupCapacityInquiryRequest) String() string {
	return tea.Prettify(s)
}

func (s GroupCapacityInquiryRequest) GoString() string {
	return s.String()
}

func (s *GroupCapacityInquiryRequest) SetEffectiveDuration(v string) *GroupCapacityInquiryRequest {
	s.EffectiveDuration = &v
	return s
}

func (s *GroupCapacityInquiryRequest) SetOpenConversationId(v string) *GroupCapacityInquiryRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GroupCapacityInquiryRequest) SetOperator(v string) *GroupCapacityInquiryRequest {
	s.Operator = &v
	return s
}

func (s *GroupCapacityInquiryRequest) SetOptions(v map[string]interface{}) *GroupCapacityInquiryRequest {
	s.Options = v
	return s
}

func (s *GroupCapacityInquiryRequest) SetTargetCapacity(v int32) *GroupCapacityInquiryRequest {
	s.TargetCapacity = &v
	return s
}

type GroupCapacityInquiryResponseBody struct {
	// example:
	//
	// 85000
	ActualPrice *int64 `json:"actualPrice,omitempty" xml:"actualPrice,omitempty"`
	// example:
	//
	// 1652183395772
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 500
	CurrentCapacity *int32 `json:"currentCapacity,omitempty" xml:"currentCapacity,omitempty"`
	// example:
	//
	// 1652183395772
	CurrentEffectUntil *int64 `json:"currentEffectUntil,omitempty" xml:"currentEffectUntil,omitempty"`
	// example:
	//
	// 85
	Discount *int32                 `json:"discount,omitempty" xml:"discount,omitempty"`
	ExtInfo  map[string]interface{} `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// example:
	//
	// 678912390478123
	GroupOwner *string `json:"groupOwner,omitempty" xml:"groupOwner,omitempty"`
	// example:
	//
	// 今天吃肘子群
	GroupTitle *string `json:"groupTitle,omitempty" xml:"groupTitle,omitempty"`
	// example:
	//
	// 10000
	MarkedPrice *int64 `json:"markedPrice,omitempty" xml:"markedPrice,omitempty"`
	// example:
	//
	// 500
	MemberCount *int32 `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	// example:
	//
	// cidoondswfakscdviouhao==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// 32453245234523425
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// example:
	//
	// 10000
	TargetCapacity *int32 `json:"targetCapacity,omitempty" xml:"targetCapacity,omitempty"`
	// example:
	//
	// 1652183395772
	TargetEffectUntil *int64 `json:"targetEffectUntil,omitempty" xml:"targetEffectUntil,omitempty"`
	// example:
	//
	// jklasdhjfasdjkfkh421jk5bb243b523
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s GroupCapacityInquiryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GroupCapacityInquiryResponseBody) GoString() string {
	return s.String()
}

func (s *GroupCapacityInquiryResponseBody) SetActualPrice(v int64) *GroupCapacityInquiryResponseBody {
	s.ActualPrice = &v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetCreatedAt(v int64) *GroupCapacityInquiryResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetCurrentCapacity(v int32) *GroupCapacityInquiryResponseBody {
	s.CurrentCapacity = &v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetCurrentEffectUntil(v int64) *GroupCapacityInquiryResponseBody {
	s.CurrentEffectUntil = &v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetDiscount(v int32) *GroupCapacityInquiryResponseBody {
	s.Discount = &v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetExtInfo(v map[string]interface{}) *GroupCapacityInquiryResponseBody {
	s.ExtInfo = v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetGroupOwner(v string) *GroupCapacityInquiryResponseBody {
	s.GroupOwner = &v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetGroupTitle(v string) *GroupCapacityInquiryResponseBody {
	s.GroupTitle = &v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetMarkedPrice(v int64) *GroupCapacityInquiryResponseBody {
	s.MarkedPrice = &v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetMemberCount(v int32) *GroupCapacityInquiryResponseBody {
	s.MemberCount = &v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetOpenConversationId(v string) *GroupCapacityInquiryResponseBody {
	s.OpenConversationId = &v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetOperator(v string) *GroupCapacityInquiryResponseBody {
	s.Operator = &v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetTargetCapacity(v int32) *GroupCapacityInquiryResponseBody {
	s.TargetCapacity = &v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetTargetEffectUntil(v int64) *GroupCapacityInquiryResponseBody {
	s.TargetEffectUntil = &v
	return s
}

func (s *GroupCapacityInquiryResponseBody) SetToken(v string) *GroupCapacityInquiryResponseBody {
	s.Token = &v
	return s
}

type GroupCapacityInquiryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GroupCapacityInquiryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GroupCapacityInquiryResponse) String() string {
	return tea.Prettify(s)
}

func (s GroupCapacityInquiryResponse) GoString() string {
	return s.String()
}

func (s *GroupCapacityInquiryResponse) SetHeaders(v map[string]*string) *GroupCapacityInquiryResponse {
	s.Headers = v
	return s
}

func (s *GroupCapacityInquiryResponse) SetStatusCode(v int32) *GroupCapacityInquiryResponse {
	s.StatusCode = &v
	return s
}

func (s *GroupCapacityInquiryResponse) SetBody(v *GroupCapacityInquiryResponseBody) *GroupCapacityInquiryResponse {
	s.Body = v
	return s
}

type GroupCapacityOrderConfirmHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GroupCapacityOrderConfirmHeaders) String() string {
	return tea.Prettify(s)
}

func (s GroupCapacityOrderConfirmHeaders) GoString() string {
	return s.String()
}

func (s *GroupCapacityOrderConfirmHeaders) SetCommonHeaders(v map[string]*string) *GroupCapacityOrderConfirmHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GroupCapacityOrderConfirmHeaders) SetXAcsDingtalkAccessToken(v string) *GroupCapacityOrderConfirmHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GroupCapacityOrderConfirmRequest struct {
	// example:
	//
	// 066224
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// example:
	//
	// FAKE:0-28937rufhjdkslnawdkjsfk
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s GroupCapacityOrderConfirmRequest) String() string {
	return tea.Prettify(s)
}

func (s GroupCapacityOrderConfirmRequest) GoString() string {
	return s.String()
}

func (s *GroupCapacityOrderConfirmRequest) SetOperator(v string) *GroupCapacityOrderConfirmRequest {
	s.Operator = &v
	return s
}

func (s *GroupCapacityOrderConfirmRequest) SetOrderId(v string) *GroupCapacityOrderConfirmRequest {
	s.OrderId = &v
	return s
}

type GroupCapacityOrderConfirmResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GroupCapacityOrderConfirmResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GroupCapacityOrderConfirmResponseBody) GoString() string {
	return s.String()
}

func (s *GroupCapacityOrderConfirmResponseBody) SetSuccess(v bool) *GroupCapacityOrderConfirmResponseBody {
	s.Success = &v
	return s
}

type GroupCapacityOrderConfirmResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GroupCapacityOrderConfirmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GroupCapacityOrderConfirmResponse) String() string {
	return tea.Prettify(s)
}

func (s GroupCapacityOrderConfirmResponse) GoString() string {
	return s.String()
}

func (s *GroupCapacityOrderConfirmResponse) SetHeaders(v map[string]*string) *GroupCapacityOrderConfirmResponse {
	s.Headers = v
	return s
}

func (s *GroupCapacityOrderConfirmResponse) SetStatusCode(v int32) *GroupCapacityOrderConfirmResponse {
	s.StatusCode = &v
	return s
}

func (s *GroupCapacityOrderConfirmResponse) SetBody(v *GroupCapacityOrderConfirmResponseBody) *GroupCapacityOrderConfirmResponse {
	s.Body = v
	return s
}

type GroupCapacityOrderPlaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GroupCapacityOrderPlaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s GroupCapacityOrderPlaceHeaders) GoString() string {
	return s.String()
}

func (s *GroupCapacityOrderPlaceHeaders) SetCommonHeaders(v map[string]*string) *GroupCapacityOrderPlaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GroupCapacityOrderPlaceHeaders) SetXAcsDingtalkAccessToken(v string) *GroupCapacityOrderPlaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GroupCapacityOrderPlaceRequest struct {
	// example:
	//
	// 123
	ActualPrice *int64 `json:"actualPrice,omitempty" xml:"actualPrice,omitempty"`
	// example:
	//
	// 500
	CurrentCapacity *int32 `json:"currentCapacity,omitempty" xml:"currentCapacity,omitempty"`
	// example:
	//
	// 1651751906
	CurrentEffectUntil *int64 `json:"currentEffectUntil,omitempty" xml:"currentEffectUntil,omitempty"`
	// example:
	//
	// 85
	Discount *int32                 `json:"discount,omitempty" xml:"discount,omitempty"`
	ExtInfo  map[string]interface{} `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// example:
	//
	// 123
	MarkedPrice *int64 `json:"markedPrice,omitempty" xml:"markedPrice,omitempty"`
	// example:
	//
	// ciddmslidnxsd==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// 531781123123
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// example:
	//
	// 1000
	TargetCapacity *int32 `json:"targetCapacity,omitempty" xml:"targetCapacity,omitempty"`
	// example:
	//
	// 1651751906
	TargetEffectUntil *int64 `json:"targetEffectUntil,omitempty" xml:"targetEffectUntil,omitempty"`
	// example:
	//
	// dfsafsd
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s GroupCapacityOrderPlaceRequest) String() string {
	return tea.Prettify(s)
}

func (s GroupCapacityOrderPlaceRequest) GoString() string {
	return s.String()
}

func (s *GroupCapacityOrderPlaceRequest) SetActualPrice(v int64) *GroupCapacityOrderPlaceRequest {
	s.ActualPrice = &v
	return s
}

func (s *GroupCapacityOrderPlaceRequest) SetCurrentCapacity(v int32) *GroupCapacityOrderPlaceRequest {
	s.CurrentCapacity = &v
	return s
}

func (s *GroupCapacityOrderPlaceRequest) SetCurrentEffectUntil(v int64) *GroupCapacityOrderPlaceRequest {
	s.CurrentEffectUntil = &v
	return s
}

func (s *GroupCapacityOrderPlaceRequest) SetDiscount(v int32) *GroupCapacityOrderPlaceRequest {
	s.Discount = &v
	return s
}

func (s *GroupCapacityOrderPlaceRequest) SetExtInfo(v map[string]interface{}) *GroupCapacityOrderPlaceRequest {
	s.ExtInfo = v
	return s
}

func (s *GroupCapacityOrderPlaceRequest) SetMarkedPrice(v int64) *GroupCapacityOrderPlaceRequest {
	s.MarkedPrice = &v
	return s
}

func (s *GroupCapacityOrderPlaceRequest) SetOpenConversationId(v string) *GroupCapacityOrderPlaceRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GroupCapacityOrderPlaceRequest) SetOperator(v string) *GroupCapacityOrderPlaceRequest {
	s.Operator = &v
	return s
}

func (s *GroupCapacityOrderPlaceRequest) SetTargetCapacity(v int32) *GroupCapacityOrderPlaceRequest {
	s.TargetCapacity = &v
	return s
}

func (s *GroupCapacityOrderPlaceRequest) SetTargetEffectUntil(v int64) *GroupCapacityOrderPlaceRequest {
	s.TargetEffectUntil = &v
	return s
}

func (s *GroupCapacityOrderPlaceRequest) SetToken(v string) *GroupCapacityOrderPlaceRequest {
	s.Token = &v
	return s
}

type GroupCapacityOrderPlaceResponseBody struct {
	// example:
	//
	// 85000
	ActualPrice *int64 `json:"actualPrice,omitempty" xml:"actualPrice,omitempty"`
	// example:
	//
	// 500
	CurrentCapacity *int32 `json:"currentCapacity,omitempty" xml:"currentCapacity,omitempty"`
	// example:
	//
	// 1652669110553
	CurrentEffectUntil *int64 `json:"currentEffectUntil,omitempty" xml:"currentEffectUntil,omitempty"`
	// example:
	//
	// 85
	Discount *int32             `json:"discount,omitempty" xml:"discount,omitempty"`
	ExtInfo  map[string]*string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// example:
	//
	// 10000
	MarkedPrice *int64 `json:"markedPrice,omitempty" xml:"markedPrice,omitempty"`
	// example:
	//
	// ciddfasvc
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// 033333
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// example:
	//
	// 12389023745345500
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// example:
	//
	// 10000
	TargetCapacity *int32 `json:"targetCapacity,omitempty" xml:"targetCapacity,omitempty"`
	// example:
	//
	// 1652669110553
	TargetEffectUntil *int64 `json:"targetEffectUntil,omitempty" xml:"targetEffectUntil,omitempty"`
	// example:
	//
	// 90ji34ontgrefv98u0ijo3q4awefg90rej
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s GroupCapacityOrderPlaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GroupCapacityOrderPlaceResponseBody) GoString() string {
	return s.String()
}

func (s *GroupCapacityOrderPlaceResponseBody) SetActualPrice(v int64) *GroupCapacityOrderPlaceResponseBody {
	s.ActualPrice = &v
	return s
}

func (s *GroupCapacityOrderPlaceResponseBody) SetCurrentCapacity(v int32) *GroupCapacityOrderPlaceResponseBody {
	s.CurrentCapacity = &v
	return s
}

func (s *GroupCapacityOrderPlaceResponseBody) SetCurrentEffectUntil(v int64) *GroupCapacityOrderPlaceResponseBody {
	s.CurrentEffectUntil = &v
	return s
}

func (s *GroupCapacityOrderPlaceResponseBody) SetDiscount(v int32) *GroupCapacityOrderPlaceResponseBody {
	s.Discount = &v
	return s
}

func (s *GroupCapacityOrderPlaceResponseBody) SetExtInfo(v map[string]*string) *GroupCapacityOrderPlaceResponseBody {
	s.ExtInfo = v
	return s
}

func (s *GroupCapacityOrderPlaceResponseBody) SetMarkedPrice(v int64) *GroupCapacityOrderPlaceResponseBody {
	s.MarkedPrice = &v
	return s
}

func (s *GroupCapacityOrderPlaceResponseBody) SetOpenConversationId(v string) *GroupCapacityOrderPlaceResponseBody {
	s.OpenConversationId = &v
	return s
}

func (s *GroupCapacityOrderPlaceResponseBody) SetOperator(v string) *GroupCapacityOrderPlaceResponseBody {
	s.Operator = &v
	return s
}

func (s *GroupCapacityOrderPlaceResponseBody) SetOrderId(v string) *GroupCapacityOrderPlaceResponseBody {
	s.OrderId = &v
	return s
}

func (s *GroupCapacityOrderPlaceResponseBody) SetTargetCapacity(v int32) *GroupCapacityOrderPlaceResponseBody {
	s.TargetCapacity = &v
	return s
}

func (s *GroupCapacityOrderPlaceResponseBody) SetTargetEffectUntil(v int64) *GroupCapacityOrderPlaceResponseBody {
	s.TargetEffectUntil = &v
	return s
}

func (s *GroupCapacityOrderPlaceResponseBody) SetToken(v string) *GroupCapacityOrderPlaceResponseBody {
	s.Token = &v
	return s
}

type GroupCapacityOrderPlaceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GroupCapacityOrderPlaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GroupCapacityOrderPlaceResponse) String() string {
	return tea.Prettify(s)
}

func (s GroupCapacityOrderPlaceResponse) GoString() string {
	return s.String()
}

func (s *GroupCapacityOrderPlaceResponse) SetHeaders(v map[string]*string) *GroupCapacityOrderPlaceResponse {
	s.Headers = v
	return s
}

func (s *GroupCapacityOrderPlaceResponse) SetStatusCode(v int32) *GroupCapacityOrderPlaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GroupCapacityOrderPlaceResponse) SetBody(v *GroupCapacityOrderPlaceResponseBody) *GroupCapacityOrderPlaceResponse {
	s.Body = v
	return s
}

type GroupManageQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GroupManageQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s GroupManageQueryHeaders) GoString() string {
	return s.String()
}

func (s *GroupManageQueryHeaders) SetCommonHeaders(v map[string]*string) *GroupManageQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GroupManageQueryHeaders) SetXAcsDingtalkAccessToken(v string) *GroupManageQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GroupManageQueryRequest struct {
	// example:
	//
	// 1652183395772
	CreatedAfter *int64 `json:"createdAfter,omitempty" xml:"createdAfter,omitempty"`
	// example:
	//
	// 53364321
	GroupId            *string   `json:"groupId,omitempty" xml:"groupId,omitempty"`
	GroupMemberSamples []*string `json:"groupMemberSamples,omitempty" xml:"groupMemberSamples,omitempty" type:"Repeated"`
	// example:
	//
	// 4122134
	GroupOwner         *string   `json:"groupOwner,omitempty" xml:"groupOwner,omitempty"`
	GroupTitleKeywords []*string `json:"groupTitleKeywords,omitempty" xml:"groupTitleKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// https://h5.dingtalk.com/circle/healthCheckin.html?dtaction=os&corpId=ding91766asjkldhfkjklasdjkfjkhajksdjkfhjkla811&5fd5e=db2ed&cbdbhh=qwertyuiop
	GroupUrl *string `json:"groupUrl,omitempty" xml:"groupUrl,omitempty"`
	// example:
	//
	// 500
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 1
	MembersOver *int32 `json:"membersOver,omitempty" xml:"membersOver,omitempty"`
	// example:
	//
	// 500
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// cidnvcxzklxv23jhkg412hj==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s GroupManageQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s GroupManageQueryRequest) GoString() string {
	return s.String()
}

func (s *GroupManageQueryRequest) SetCreatedAfter(v int64) *GroupManageQueryRequest {
	s.CreatedAfter = &v
	return s
}

func (s *GroupManageQueryRequest) SetGroupId(v string) *GroupManageQueryRequest {
	s.GroupId = &v
	return s
}

func (s *GroupManageQueryRequest) SetGroupMemberSamples(v []*string) *GroupManageQueryRequest {
	s.GroupMemberSamples = v
	return s
}

func (s *GroupManageQueryRequest) SetGroupOwner(v string) *GroupManageQueryRequest {
	s.GroupOwner = &v
	return s
}

func (s *GroupManageQueryRequest) SetGroupTitleKeywords(v []*string) *GroupManageQueryRequest {
	s.GroupTitleKeywords = v
	return s
}

func (s *GroupManageQueryRequest) SetGroupUrl(v string) *GroupManageQueryRequest {
	s.GroupUrl = &v
	return s
}

func (s *GroupManageQueryRequest) SetMaxResults(v int32) *GroupManageQueryRequest {
	s.MaxResults = &v
	return s
}

func (s *GroupManageQueryRequest) SetMembersOver(v int32) *GroupManageQueryRequest {
	s.MembersOver = &v
	return s
}

func (s *GroupManageQueryRequest) SetNextToken(v string) *GroupManageQueryRequest {
	s.NextToken = &v
	return s
}

func (s *GroupManageQueryRequest) SetOpenConversationId(v string) *GroupManageQueryRequest {
	s.OpenConversationId = &v
	return s
}

type GroupManageQueryResponseBody struct {
	GroupInfoList []*GroupManageQueryResponseBodyGroupInfoList `json:"groupInfoList,omitempty" xml:"groupInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// example:
	//
	// 500
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GroupManageQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GroupManageQueryResponseBody) GoString() string {
	return s.String()
}

func (s *GroupManageQueryResponseBody) SetGroupInfoList(v []*GroupManageQueryResponseBodyGroupInfoList) *GroupManageQueryResponseBody {
	s.GroupInfoList = v
	return s
}

func (s *GroupManageQueryResponseBody) SetHasMore(v bool) *GroupManageQueryResponseBody {
	s.HasMore = &v
	return s
}

func (s *GroupManageQueryResponseBody) SetNextToken(v string) *GroupManageQueryResponseBody {
	s.NextToken = &v
	return s
}

type GroupManageQueryResponseBodyGroupInfoList struct {
	// example:
	//
	// 0
	BanWordsMode *int32 `json:"banWordsMode,omitempty" xml:"banWordsMode,omitempty"`
	// example:
	//
	// 1000
	Capacity *int32 `json:"capacity,omitempty" xml:"capacity,omitempty"`
	// example:
	//
	// 1652183395772
	CreatedAt      *int64                 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	ExtInfo        map[string]interface{} `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	GroupAdminList []*string              `json:"groupAdminList,omitempty" xml:"groupAdminList,omitempty" type:"Repeated"`
	// example:
	//
	// 574892167781263748
	GroupOwner *string `json:"groupOwner,omitempty" xml:"groupOwner,omitempty"`
	// example:
	//
	// 今天吃肘子群
	GroupTitle *string `json:"groupTitle,omitempty" xml:"groupTitle,omitempty"`
	// example:
	//
	// 500
	MemberCount *int32 `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	// example:
	//
	// cidnvcxzklxv23jhkg412hj==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// INNER
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GroupManageQueryResponseBodyGroupInfoList) String() string {
	return tea.Prettify(s)
}

func (s GroupManageQueryResponseBodyGroupInfoList) GoString() string {
	return s.String()
}

func (s *GroupManageQueryResponseBodyGroupInfoList) SetBanWordsMode(v int32) *GroupManageQueryResponseBodyGroupInfoList {
	s.BanWordsMode = &v
	return s
}

func (s *GroupManageQueryResponseBodyGroupInfoList) SetCapacity(v int32) *GroupManageQueryResponseBodyGroupInfoList {
	s.Capacity = &v
	return s
}

func (s *GroupManageQueryResponseBodyGroupInfoList) SetCreatedAt(v int64) *GroupManageQueryResponseBodyGroupInfoList {
	s.CreatedAt = &v
	return s
}

func (s *GroupManageQueryResponseBodyGroupInfoList) SetExtInfo(v map[string]interface{}) *GroupManageQueryResponseBodyGroupInfoList {
	s.ExtInfo = v
	return s
}

func (s *GroupManageQueryResponseBodyGroupInfoList) SetGroupAdminList(v []*string) *GroupManageQueryResponseBodyGroupInfoList {
	s.GroupAdminList = v
	return s
}

func (s *GroupManageQueryResponseBodyGroupInfoList) SetGroupOwner(v string) *GroupManageQueryResponseBodyGroupInfoList {
	s.GroupOwner = &v
	return s
}

func (s *GroupManageQueryResponseBodyGroupInfoList) SetGroupTitle(v string) *GroupManageQueryResponseBodyGroupInfoList {
	s.GroupTitle = &v
	return s
}

func (s *GroupManageQueryResponseBodyGroupInfoList) SetMemberCount(v int32) *GroupManageQueryResponseBodyGroupInfoList {
	s.MemberCount = &v
	return s
}

func (s *GroupManageQueryResponseBodyGroupInfoList) SetOpenConversationId(v string) *GroupManageQueryResponseBodyGroupInfoList {
	s.OpenConversationId = &v
	return s
}

func (s *GroupManageQueryResponseBodyGroupInfoList) SetType(v string) *GroupManageQueryResponseBodyGroupInfoList {
	s.Type = &v
	return s
}

type GroupManageQueryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GroupManageQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GroupManageQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s GroupManageQueryResponse) GoString() string {
	return s.String()
}

func (s *GroupManageQueryResponse) SetHeaders(v map[string]*string) *GroupManageQueryResponse {
	s.Headers = v
	return s
}

func (s *GroupManageQueryResponse) SetStatusCode(v int32) *GroupManageQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *GroupManageQueryResponse) SetBody(v *GroupManageQueryResponseBody) *GroupManageQueryResponse {
	s.Body = v
	return s
}

type GroupManageReduceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GroupManageReduceHeaders) String() string {
	return tea.Prettify(s)
}

func (s GroupManageReduceHeaders) GoString() string {
	return s.String()
}

func (s *GroupManageReduceHeaders) SetCommonHeaders(v map[string]*string) *GroupManageReduceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GroupManageReduceHeaders) SetXAcsDingtalkAccessToken(v string) *GroupManageReduceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GroupManageReduceRequest struct {
	// example:
	//
	// 200
	CapacityLimit *int32 `json:"capacityLimit,omitempty" xml:"capacityLimit,omitempty"`
	// example:
	//
	// cidnvcxzklxv
	OpenConversationId *string                `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	Options            map[string]interface{} `json:"options,omitempty" xml:"options,omitempty"`
}

func (s GroupManageReduceRequest) String() string {
	return tea.Prettify(s)
}

func (s GroupManageReduceRequest) GoString() string {
	return s.String()
}

func (s *GroupManageReduceRequest) SetCapacityLimit(v int32) *GroupManageReduceRequest {
	s.CapacityLimit = &v
	return s
}

func (s *GroupManageReduceRequest) SetOpenConversationId(v string) *GroupManageReduceRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GroupManageReduceRequest) SetOptions(v map[string]interface{}) *GroupManageReduceRequest {
	s.Options = v
	return s
}

type GroupManageReduceResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s GroupManageReduceResponse) String() string {
	return tea.Prettify(s)
}

func (s GroupManageReduceResponse) GoString() string {
	return s.String()
}

func (s *GroupManageReduceResponse) SetHeaders(v map[string]*string) *GroupManageReduceResponse {
	s.Headers = v
	return s
}

func (s *GroupManageReduceResponse) SetStatusCode(v int32) *GroupManageReduceResponse {
	s.StatusCode = &v
	return s
}

type ImportGroupChatHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ImportGroupChatHeaders) String() string {
	return tea.Prettify(s)
}

func (s ImportGroupChatHeaders) GoString() string {
	return s.String()
}

func (s *ImportGroupChatHeaders) SetCommonHeaders(v map[string]*string) *ImportGroupChatHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ImportGroupChatHeaders) SetXAcsDingtalkAccessToken(v string) *ImportGroupChatHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ImportGroupChatRequest struct {
	AdminIds []*string `json:"adminIds,omitempty" xml:"adminIds,omitempty" type:"Repeated"`
	CreateAt *int64    `json:"createAt,omitempty" xml:"createAt,omitempty"`
	// example:
	//
	// @lADOADma*****QKA
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// axcf*-*****-*****-23da*
	ImportUuid *string `json:"importUuid,omitempty" xml:"importUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1107****2120
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// example:
	//
	// c354***-***-***-b4ea-6f1ab***65
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例群名称
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	UserList map[string]*UserListValue `json:"userList,omitempty" xml:"userList,omitempty"`
}

func (s ImportGroupChatRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportGroupChatRequest) GoString() string {
	return s.String()
}

func (s *ImportGroupChatRequest) SetAdminIds(v []*string) *ImportGroupChatRequest {
	s.AdminIds = v
	return s
}

func (s *ImportGroupChatRequest) SetCreateAt(v int64) *ImportGroupChatRequest {
	s.CreateAt = &v
	return s
}

func (s *ImportGroupChatRequest) SetIcon(v string) *ImportGroupChatRequest {
	s.Icon = &v
	return s
}

func (s *ImportGroupChatRequest) SetImportUuid(v string) *ImportGroupChatRequest {
	s.ImportUuid = &v
	return s
}

func (s *ImportGroupChatRequest) SetOwner(v string) *ImportGroupChatRequest {
	s.Owner = &v
	return s
}

func (s *ImportGroupChatRequest) SetTemplateId(v string) *ImportGroupChatRequest {
	s.TemplateId = &v
	return s
}

func (s *ImportGroupChatRequest) SetTitle(v string) *ImportGroupChatRequest {
	s.Title = &v
	return s
}

func (s *ImportGroupChatRequest) SetUserList(v map[string]*UserListValue) *ImportGroupChatRequest {
	s.UserList = v
	return s
}

type ImportGroupChatResponseBody struct {
	Result  *ImportGroupChatResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *string                            `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ImportGroupChatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportGroupChatResponseBody) GoString() string {
	return s.String()
}

func (s *ImportGroupChatResponseBody) SetResult(v *ImportGroupChatResponseBodyResult) *ImportGroupChatResponseBody {
	s.Result = v
	return s
}

func (s *ImportGroupChatResponseBody) SetSuccess(v string) *ImportGroupChatResponseBody {
	s.Success = &v
	return s
}

type ImportGroupChatResponseBodyResult struct {
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s ImportGroupChatResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ImportGroupChatResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ImportGroupChatResponseBodyResult) SetOpenConversationId(v string) *ImportGroupChatResponseBodyResult {
	s.OpenConversationId = &v
	return s
}

type ImportGroupChatResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportGroupChatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportGroupChatResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportGroupChatResponse) GoString() string {
	return s.String()
}

func (s *ImportGroupChatResponse) SetHeaders(v map[string]*string) *ImportGroupChatResponse {
	s.Headers = v
	return s
}

func (s *ImportGroupChatResponse) SetStatusCode(v int32) *ImportGroupChatResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportGroupChatResponse) SetBody(v *ImportGroupChatResponseBody) *ImportGroupChatResponse {
	s.Body = v
	return s
}

type ImportMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ImportMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s ImportMessageHeaders) GoString() string {
	return s.String()
}

func (s *ImportMessageHeaders) SetCommonHeaders(v map[string]*string) *ImportMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ImportMessageHeaders) SetXAcsDingtalkAccessToken(v string) *ImportMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ImportMessageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {"content":"月会通知<@all> ","at":{"atUserIds":[],"isAtAll":true}}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// axcf*-*****-*****-23da*
	ImportUuid           *string `json:"importUuid,omitempty" xml:"importUuid,omitempty"`
	MsgReadStatusSetting *bool   `json:"msgReadStatusSetting,omitempty" xml:"msgReadStatusSetting,omitempty"`
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
	// cidt*****Xa4K10w==
	OpenConversationId *string   `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	Receivers          []*string `json:"receivers,omitempty" xml:"receivers,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 013*****21312
	SenderId *string `json:"senderId,omitempty" xml:"senderId,omitempty"`
}

func (s ImportMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportMessageRequest) GoString() string {
	return s.String()
}

func (s *ImportMessageRequest) SetContent(v string) *ImportMessageRequest {
	s.Content = &v
	return s
}

func (s *ImportMessageRequest) SetCreateTime(v int64) *ImportMessageRequest {
	s.CreateTime = &v
	return s
}

func (s *ImportMessageRequest) SetImportUuid(v string) *ImportMessageRequest {
	s.ImportUuid = &v
	return s
}

func (s *ImportMessageRequest) SetMsgReadStatusSetting(v bool) *ImportMessageRequest {
	s.MsgReadStatusSetting = &v
	return s
}

func (s *ImportMessageRequest) SetMsgType(v string) *ImportMessageRequest {
	s.MsgType = &v
	return s
}

func (s *ImportMessageRequest) SetOpenConversationId(v string) *ImportMessageRequest {
	s.OpenConversationId = &v
	return s
}

func (s *ImportMessageRequest) SetReceivers(v []*string) *ImportMessageRequest {
	s.Receivers = v
	return s
}

func (s *ImportMessageRequest) SetSenderId(v string) *ImportMessageRequest {
	s.SenderId = &v
	return s
}

type ImportMessageResponseBody struct {
	Result  *ImportMessageResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *string                          `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ImportMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportMessageResponseBody) GoString() string {
	return s.String()
}

func (s *ImportMessageResponseBody) SetResult(v *ImportMessageResponseBodyResult) *ImportMessageResponseBody {
	s.Result = v
	return s
}

func (s *ImportMessageResponseBody) SetSuccess(v string) *ImportMessageResponseBody {
	s.Success = &v
	return s
}

type ImportMessageResponseBodyResult struct {
	OpenTaskId *string `json:"openTaskId,omitempty" xml:"openTaskId,omitempty"`
}

func (s ImportMessageResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ImportMessageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ImportMessageResponseBodyResult) SetOpenTaskId(v string) *ImportMessageResponseBodyResult {
	s.OpenTaskId = &v
	return s
}

type ImportMessageResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportMessageResponse) GoString() string {
	return s.String()
}

func (s *ImportMessageResponse) SetHeaders(v map[string]*string) *ImportMessageResponse {
	s.Headers = v
	return s
}

func (s *ImportMessageResponse) SetStatusCode(v int32) *ImportMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportMessageResponse) SetBody(v *ImportMessageResponseBody) *ImportMessageResponse {
	s.Body = v
	return s
}

type InstallRobotToOrgHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InstallRobotToOrgHeaders) String() string {
	return tea.Prettify(s)
}

func (s InstallRobotToOrgHeaders) GoString() string {
	return s.String()
}

func (s *InstallRobotToOrgHeaders) SetCommonHeaders(v map[string]*string) *InstallRobotToOrgHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InstallRobotToOrgHeaders) SetXAcsDingtalkAccessToken(v string) *InstallRobotToOrgHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InstallRobotToOrgRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 这是小丁
	Brief *string `json:"brief,omitempty" xml:"brief,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 我是小丁
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// @lALPDe7s26Bre
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 小丁
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	OutgoingToken *string `json:"outgoingToken,omitempty" xml:"outgoingToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://*.com
	OutgoingUrl *string `json:"outgoingUrl,omitempty" xml:"outgoingUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// @lALPDe7s26Bre
	PreviewMediaId *string `json:"previewMediaId,omitempty" xml:"previewMediaId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s InstallRobotToOrgRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallRobotToOrgRequest) GoString() string {
	return s.String()
}

func (s *InstallRobotToOrgRequest) SetBrief(v string) *InstallRobotToOrgRequest {
	s.Brief = &v
	return s
}

func (s *InstallRobotToOrgRequest) SetDescription(v string) *InstallRobotToOrgRequest {
	s.Description = &v
	return s
}

func (s *InstallRobotToOrgRequest) SetIcon(v string) *InstallRobotToOrgRequest {
	s.Icon = &v
	return s
}

func (s *InstallRobotToOrgRequest) SetName(v string) *InstallRobotToOrgRequest {
	s.Name = &v
	return s
}

func (s *InstallRobotToOrgRequest) SetOutgoingToken(v string) *InstallRobotToOrgRequest {
	s.OutgoingToken = &v
	return s
}

func (s *InstallRobotToOrgRequest) SetOutgoingUrl(v string) *InstallRobotToOrgRequest {
	s.OutgoingUrl = &v
	return s
}

func (s *InstallRobotToOrgRequest) SetPreviewMediaId(v string) *InstallRobotToOrgRequest {
	s.PreviewMediaId = &v
	return s
}

func (s *InstallRobotToOrgRequest) SetRobotCode(v string) *InstallRobotToOrgRequest {
	s.RobotCode = &v
	return s
}

type InstallRobotToOrgResponseBody struct {
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s InstallRobotToOrgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallRobotToOrgResponseBody) GoString() string {
	return s.String()
}

func (s *InstallRobotToOrgResponseBody) SetRobotCode(v string) *InstallRobotToOrgResponseBody {
	s.RobotCode = &v
	return s
}

type InstallRobotToOrgResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallRobotToOrgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallRobotToOrgResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallRobotToOrgResponse) GoString() string {
	return s.String()
}

func (s *InstallRobotToOrgResponse) SetHeaders(v map[string]*string) *InstallRobotToOrgResponse {
	s.Headers = v
	return s
}

func (s *InstallRobotToOrgResponse) SetStatusCode(v int32) *InstallRobotToOrgResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallRobotToOrgResponse) SetBody(v *InstallRobotToOrgResponseBody) *InstallRobotToOrgResponse {
	s.Body = v
	return s
}

type InteractiveCardCreateInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InteractiveCardCreateInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s InteractiveCardCreateInstanceHeaders) GoString() string {
	return s.String()
}

func (s *InteractiveCardCreateInstanceHeaders) SetCommonHeaders(v map[string]*string) *InteractiveCardCreateInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InteractiveCardCreateInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *InteractiveCardCreateInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InteractiveCardCreateInstanceRequest struct {
	CallbackRouteKey *string `json:"callbackRouteKey,omitempty" xml:"callbackRouteKey,omitempty"`
	// This parameter is required.
	CardData *InteractiveCardCreateInstanceRequestCardData `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	// This parameter is required.
	CardTemplateId *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	ChatBotId      *string `json:"chatBotId,omitempty" xml:"chatBotId,omitempty"`
	// This parameter is required.
	ConversationType   *int32  `json:"conversationType,omitempty" xml:"conversationType,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	OutTrackId         *string                      `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	PrivateData        map[string]*PrivateDataValue `json:"privateData,omitempty" xml:"privateData,omitempty"`
	PullStrategy       *bool                        `json:"pullStrategy,omitempty" xml:"pullStrategy,omitempty"`
	ReceiverUserIdList []*string                    `json:"receiverUserIdList,omitempty" xml:"receiverUserIdList,omitempty" type:"Repeated"`
	RobotCode          *string                      `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	UserIdType         *int32                       `json:"userIdType,omitempty" xml:"userIdType,omitempty"`
}

func (s InteractiveCardCreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s InteractiveCardCreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *InteractiveCardCreateInstanceRequest) SetCallbackRouteKey(v string) *InteractiveCardCreateInstanceRequest {
	s.CallbackRouteKey = &v
	return s
}

func (s *InteractiveCardCreateInstanceRequest) SetCardData(v *InteractiveCardCreateInstanceRequestCardData) *InteractiveCardCreateInstanceRequest {
	s.CardData = v
	return s
}

func (s *InteractiveCardCreateInstanceRequest) SetCardTemplateId(v string) *InteractiveCardCreateInstanceRequest {
	s.CardTemplateId = &v
	return s
}

func (s *InteractiveCardCreateInstanceRequest) SetChatBotId(v string) *InteractiveCardCreateInstanceRequest {
	s.ChatBotId = &v
	return s
}

func (s *InteractiveCardCreateInstanceRequest) SetConversationType(v int32) *InteractiveCardCreateInstanceRequest {
	s.ConversationType = &v
	return s
}

func (s *InteractiveCardCreateInstanceRequest) SetOpenConversationId(v string) *InteractiveCardCreateInstanceRequest {
	s.OpenConversationId = &v
	return s
}

func (s *InteractiveCardCreateInstanceRequest) SetOutTrackId(v string) *InteractiveCardCreateInstanceRequest {
	s.OutTrackId = &v
	return s
}

func (s *InteractiveCardCreateInstanceRequest) SetPrivateData(v map[string]*PrivateDataValue) *InteractiveCardCreateInstanceRequest {
	s.PrivateData = v
	return s
}

func (s *InteractiveCardCreateInstanceRequest) SetPullStrategy(v bool) *InteractiveCardCreateInstanceRequest {
	s.PullStrategy = &v
	return s
}

func (s *InteractiveCardCreateInstanceRequest) SetReceiverUserIdList(v []*string) *InteractiveCardCreateInstanceRequest {
	s.ReceiverUserIdList = v
	return s
}

func (s *InteractiveCardCreateInstanceRequest) SetRobotCode(v string) *InteractiveCardCreateInstanceRequest {
	s.RobotCode = &v
	return s
}

func (s *InteractiveCardCreateInstanceRequest) SetUserIdType(v int32) *InteractiveCardCreateInstanceRequest {
	s.UserIdType = &v
	return s
}

type InteractiveCardCreateInstanceRequestCardData struct {
	CardMediaIdParamMap map[string]*string `json:"cardMediaIdParamMap,omitempty" xml:"cardMediaIdParamMap,omitempty"`
	CardParamMap        map[string]*string `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
}

func (s InteractiveCardCreateInstanceRequestCardData) String() string {
	return tea.Prettify(s)
}

func (s InteractiveCardCreateInstanceRequestCardData) GoString() string {
	return s.String()
}

func (s *InteractiveCardCreateInstanceRequestCardData) SetCardMediaIdParamMap(v map[string]*string) *InteractiveCardCreateInstanceRequestCardData {
	s.CardMediaIdParamMap = v
	return s
}

func (s *InteractiveCardCreateInstanceRequestCardData) SetCardParamMap(v map[string]*string) *InteractiveCardCreateInstanceRequestCardData {
	s.CardParamMap = v
	return s
}

type InteractiveCardCreateInstanceResponseBody struct {
	// example:
	//
	// xxxxxx
	ProcessQueryKey *string `json:"processQueryKey,omitempty" xml:"processQueryKey,omitempty"`
}

func (s InteractiveCardCreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InteractiveCardCreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *InteractiveCardCreateInstanceResponseBody) SetProcessQueryKey(v string) *InteractiveCardCreateInstanceResponseBody {
	s.ProcessQueryKey = &v
	return s
}

type InteractiveCardCreateInstanceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InteractiveCardCreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InteractiveCardCreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s InteractiveCardCreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *InteractiveCardCreateInstanceResponse) SetHeaders(v map[string]*string) *InteractiveCardCreateInstanceResponse {
	s.Headers = v
	return s
}

func (s *InteractiveCardCreateInstanceResponse) SetStatusCode(v int32) *InteractiveCardCreateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *InteractiveCardCreateInstanceResponse) SetBody(v *InteractiveCardCreateInstanceResponseBody) *InteractiveCardCreateInstanceResponse {
	s.Body = v
	return s
}

type ListGroupTemplatesByOrgIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListGroupTemplatesByOrgIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListGroupTemplatesByOrgIdHeaders) GoString() string {
	return s.String()
}

func (s *ListGroupTemplatesByOrgIdHeaders) SetCommonHeaders(v map[string]*string) *ListGroupTemplatesByOrgIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListGroupTemplatesByOrgIdHeaders) SetXAcsDingtalkAccessToken(v string) *ListGroupTemplatesByOrgIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListGroupTemplatesByOrgIdRequest struct {
	// This parameter is required.
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListGroupTemplatesByOrgIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupTemplatesByOrgIdRequest) GoString() string {
	return s.String()
}

func (s *ListGroupTemplatesByOrgIdRequest) SetPageNumber(v int32) *ListGroupTemplatesByOrgIdRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGroupTemplatesByOrgIdRequest) SetPageSize(v int32) *ListGroupTemplatesByOrgIdRequest {
	s.PageSize = &v
	return s
}

type ListGroupTemplatesByOrgIdResponseBody struct {
	Count                  *int32                                                         `json:"count,omitempty" xml:"count,omitempty"`
	SceneGroupDetailModels []*ListGroupTemplatesByOrgIdResponseBodySceneGroupDetailModels `json:"sceneGroupDetailModels,omitempty" xml:"sceneGroupDetailModels,omitempty" type:"Repeated"`
	Success                *bool                                                          `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListGroupTemplatesByOrgIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupTemplatesByOrgIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupTemplatesByOrgIdResponseBody) SetCount(v int32) *ListGroupTemplatesByOrgIdResponseBody {
	s.Count = &v
	return s
}

func (s *ListGroupTemplatesByOrgIdResponseBody) SetSceneGroupDetailModels(v []*ListGroupTemplatesByOrgIdResponseBodySceneGroupDetailModels) *ListGroupTemplatesByOrgIdResponseBody {
	s.SceneGroupDetailModels = v
	return s
}

func (s *ListGroupTemplatesByOrgIdResponseBody) SetSuccess(v bool) *ListGroupTemplatesByOrgIdResponseBody {
	s.Success = &v
	return s
}

type ListGroupTemplatesByOrgIdResponseBodySceneGroupDetailModels struct {
	Description  *string `json:"description,omitempty" xml:"description,omitempty"`
	GmtCreate    *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	Icon         *string `json:"icon,omitempty" xml:"icon,omitempty"`
	MsgOpen      *bool   `json:"msgOpen,omitempty" xml:"msgOpen,omitempty"`
	TemplateId   *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
}

func (s ListGroupTemplatesByOrgIdResponseBodySceneGroupDetailModels) String() string {
	return tea.Prettify(s)
}

func (s ListGroupTemplatesByOrgIdResponseBodySceneGroupDetailModels) GoString() string {
	return s.String()
}

func (s *ListGroupTemplatesByOrgIdResponseBodySceneGroupDetailModels) SetDescription(v string) *ListGroupTemplatesByOrgIdResponseBodySceneGroupDetailModels {
	s.Description = &v
	return s
}

func (s *ListGroupTemplatesByOrgIdResponseBodySceneGroupDetailModels) SetGmtCreate(v string) *ListGroupTemplatesByOrgIdResponseBodySceneGroupDetailModels {
	s.GmtCreate = &v
	return s
}

func (s *ListGroupTemplatesByOrgIdResponseBodySceneGroupDetailModels) SetIcon(v string) *ListGroupTemplatesByOrgIdResponseBodySceneGroupDetailModels {
	s.Icon = &v
	return s
}

func (s *ListGroupTemplatesByOrgIdResponseBodySceneGroupDetailModels) SetMsgOpen(v bool) *ListGroupTemplatesByOrgIdResponseBodySceneGroupDetailModels {
	s.MsgOpen = &v
	return s
}

func (s *ListGroupTemplatesByOrgIdResponseBodySceneGroupDetailModels) SetTemplateId(v string) *ListGroupTemplatesByOrgIdResponseBodySceneGroupDetailModels {
	s.TemplateId = &v
	return s
}

func (s *ListGroupTemplatesByOrgIdResponseBodySceneGroupDetailModels) SetTemplateName(v string) *ListGroupTemplatesByOrgIdResponseBodySceneGroupDetailModels {
	s.TemplateName = &v
	return s
}

type ListGroupTemplatesByOrgIdResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupTemplatesByOrgIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupTemplatesByOrgIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupTemplatesByOrgIdResponse) GoString() string {
	return s.String()
}

func (s *ListGroupTemplatesByOrgIdResponse) SetHeaders(v map[string]*string) *ListGroupTemplatesByOrgIdResponse {
	s.Headers = v
	return s
}

func (s *ListGroupTemplatesByOrgIdResponse) SetStatusCode(v int32) *ListGroupTemplatesByOrgIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupTemplatesByOrgIdResponse) SetBody(v *ListGroupTemplatesByOrgIdResponseBody) *ListGroupTemplatesByOrgIdResponse {
	s.Body = v
	return s
}

type ListOrgTextEmotionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListOrgTextEmotionHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListOrgTextEmotionHeaders) GoString() string {
	return s.String()
}

func (s *ListOrgTextEmotionHeaders) SetCommonHeaders(v map[string]*string) *ListOrgTextEmotionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListOrgTextEmotionHeaders) SetXAcsDingtalkAccessToken(v string) *ListOrgTextEmotionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListOrgTextEmotionResponseBody struct {
	Result  *ListOrgTextEmotionResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                 `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListOrgTextEmotionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrgTextEmotionResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrgTextEmotionResponseBody) SetResult(v *ListOrgTextEmotionResponseBodyResult) *ListOrgTextEmotionResponseBody {
	s.Result = v
	return s
}

func (s *ListOrgTextEmotionResponseBody) SetSuccess(v bool) *ListOrgTextEmotionResponseBody {
	s.Success = &v
	return s
}

type ListOrgTextEmotionResponseBodyResult struct {
	Emotions []*ListOrgTextEmotionResponseBodyResultEmotions `json:"emotions,omitempty" xml:"emotions,omitempty" type:"Repeated"`
}

func (s ListOrgTextEmotionResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListOrgTextEmotionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListOrgTextEmotionResponseBodyResult) SetEmotions(v []*ListOrgTextEmotionResponseBodyResultEmotions) *ListOrgTextEmotionResponseBodyResult {
	s.Emotions = v
	return s
}

type ListOrgTextEmotionResponseBodyResultEmotions struct {
	// example:
	//
	// @234xxx
	BackgroundMediaId *string `json:"backgroundMediaId,omitempty" xml:"backgroundMediaId,omitempty"`
	// example:
	//
	// @123xxx
	BackgroundMediaIdForPanel *string `json:"backgroundMediaIdForPanel,omitempty" xml:"backgroundMediaIdForPanel,omitempty"`
	// example:
	//
	// -1
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// example:
	//
	// corp_131xxx
	EmotionId *string `json:"emotionId,omitempty" xml:"emotionId,omitempty"`
	// example:
	//
	// 企业表情1
	EmotionName *string `json:"emotionName,omitempty" xml:"emotionName,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListOrgTextEmotionResponseBodyResultEmotions) String() string {
	return tea.Prettify(s)
}

func (s ListOrgTextEmotionResponseBodyResultEmotions) GoString() string {
	return s.String()
}

func (s *ListOrgTextEmotionResponseBodyResultEmotions) SetBackgroundMediaId(v string) *ListOrgTextEmotionResponseBodyResultEmotions {
	s.BackgroundMediaId = &v
	return s
}

func (s *ListOrgTextEmotionResponseBodyResultEmotions) SetBackgroundMediaIdForPanel(v string) *ListOrgTextEmotionResponseBodyResultEmotions {
	s.BackgroundMediaIdForPanel = &v
	return s
}

func (s *ListOrgTextEmotionResponseBodyResultEmotions) SetDeptId(v int64) *ListOrgTextEmotionResponseBodyResultEmotions {
	s.DeptId = &v
	return s
}

func (s *ListOrgTextEmotionResponseBodyResultEmotions) SetEmotionId(v string) *ListOrgTextEmotionResponseBodyResultEmotions {
	s.EmotionId = &v
	return s
}

func (s *ListOrgTextEmotionResponseBodyResultEmotions) SetEmotionName(v string) *ListOrgTextEmotionResponseBodyResultEmotions {
	s.EmotionName = &v
	return s
}

func (s *ListOrgTextEmotionResponseBodyResultEmotions) SetStatus(v int32) *ListOrgTextEmotionResponseBodyResultEmotions {
	s.Status = &v
	return s
}

type ListOrgTextEmotionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOrgTextEmotionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOrgTextEmotionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrgTextEmotionResponse) GoString() string {
	return s.String()
}

func (s *ListOrgTextEmotionResponse) SetHeaders(v map[string]*string) *ListOrgTextEmotionResponse {
	s.Headers = v
	return s
}

func (s *ListOrgTextEmotionResponse) SetStatusCode(v int32) *ListOrgTextEmotionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrgTextEmotionResponse) SetBody(v *ListOrgTextEmotionResponseBody) *ListOrgTextEmotionResponse {
	s.Body = v
	return s
}

type ListSceneGroupsByTemplateIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListSceneGroupsByTemplateIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListSceneGroupsByTemplateIdHeaders) GoString() string {
	return s.String()
}

func (s *ListSceneGroupsByTemplateIdHeaders) SetCommonHeaders(v map[string]*string) *ListSceneGroupsByTemplateIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSceneGroupsByTemplateIdHeaders) SetXAcsDingtalkAccessToken(v string) *ListSceneGroupsByTemplateIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListSceneGroupsByTemplateIdRequest struct {
	// This parameter is required.
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListSceneGroupsByTemplateIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSceneGroupsByTemplateIdRequest) GoString() string {
	return s.String()
}

func (s *ListSceneGroupsByTemplateIdRequest) SetPageNumber(v int32) *ListSceneGroupsByTemplateIdRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSceneGroupsByTemplateIdRequest) SetPageSize(v int32) *ListSceneGroupsByTemplateIdRequest {
	s.PageSize = &v
	return s
}

type ListSceneGroupsByTemplateIdResponseBody struct {
	OpenConversationIds []*string `json:"openConversationIds,omitempty" xml:"openConversationIds,omitempty" type:"Repeated"`
	Success             *bool     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListSceneGroupsByTemplateIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSceneGroupsByTemplateIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListSceneGroupsByTemplateIdResponseBody) SetOpenConversationIds(v []*string) *ListSceneGroupsByTemplateIdResponseBody {
	s.OpenConversationIds = v
	return s
}

func (s *ListSceneGroupsByTemplateIdResponseBody) SetSuccess(v bool) *ListSceneGroupsByTemplateIdResponseBody {
	s.Success = &v
	return s
}

type ListSceneGroupsByTemplateIdResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSceneGroupsByTemplateIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSceneGroupsByTemplateIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSceneGroupsByTemplateIdResponse) GoString() string {
	return s.String()
}

func (s *ListSceneGroupsByTemplateIdResponse) SetHeaders(v map[string]*string) *ListSceneGroupsByTemplateIdResponse {
	s.Headers = v
	return s
}

func (s *ListSceneGroupsByTemplateIdResponse) SetStatusCode(v int32) *ListSceneGroupsByTemplateIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSceneGroupsByTemplateIdResponse) SetBody(v *ListSceneGroupsByTemplateIdResponseBody) *ListSceneGroupsByTemplateIdResponse {
	s.Body = v
	return s
}

type OfflineUnfurlingRegisterHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s OfflineUnfurlingRegisterHeaders) String() string {
	return tea.Prettify(s)
}

func (s OfflineUnfurlingRegisterHeaders) GoString() string {
	return s.String()
}

func (s *OfflineUnfurlingRegisterHeaders) SetCommonHeaders(v map[string]*string) *OfflineUnfurlingRegisterHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OfflineUnfurlingRegisterHeaders) SetXAcsDingtalkAccessToken(v string) *OfflineUnfurlingRegisterHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type OfflineUnfurlingRegisterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3102xxxxxxx
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 37xxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s OfflineUnfurlingRegisterRequest) String() string {
	return tea.Prettify(s)
}

func (s OfflineUnfurlingRegisterRequest) GoString() string {
	return s.String()
}

func (s *OfflineUnfurlingRegisterRequest) SetAppId(v string) *OfflineUnfurlingRegisterRequest {
	s.AppId = &v
	return s
}

func (s *OfflineUnfurlingRegisterRequest) SetId(v int64) *OfflineUnfurlingRegisterRequest {
	s.Id = &v
	return s
}

func (s *OfflineUnfurlingRegisterRequest) SetUserId(v string) *OfflineUnfurlingRegisterRequest {
	s.UserId = &v
	return s
}

type OfflineUnfurlingRegisterResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OfflineUnfurlingRegisterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OfflineUnfurlingRegisterResponseBody) GoString() string {
	return s.String()
}

func (s *OfflineUnfurlingRegisterResponseBody) SetSuccess(v bool) *OfflineUnfurlingRegisterResponseBody {
	s.Success = &v
	return s
}

type OfflineUnfurlingRegisterResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OfflineUnfurlingRegisterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OfflineUnfurlingRegisterResponse) String() string {
	return tea.Prettify(s)
}

func (s OfflineUnfurlingRegisterResponse) GoString() string {
	return s.String()
}

func (s *OfflineUnfurlingRegisterResponse) SetHeaders(v map[string]*string) *OfflineUnfurlingRegisterResponse {
	s.Headers = v
	return s
}

func (s *OfflineUnfurlingRegisterResponse) SetStatusCode(v int32) *OfflineUnfurlingRegisterResponse {
	s.StatusCode = &v
	return s
}

func (s *OfflineUnfurlingRegisterResponse) SetBody(v *OfflineUnfurlingRegisterResponseBody) *OfflineUnfurlingRegisterResponse {
	s.Body = v
	return s
}

type OpenGroupRoleAddHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s OpenGroupRoleAddHeaders) String() string {
	return tea.Prettify(s)
}

func (s OpenGroupRoleAddHeaders) GoString() string {
	return s.String()
}

func (s *OpenGroupRoleAddHeaders) SetCommonHeaders(v map[string]*string) *OpenGroupRoleAddHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OpenGroupRoleAddHeaders) SetXAcsDingtalkAccessToken(v string) *OpenGroupRoleAddHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type OpenGroupRoleAddRequest struct {
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s OpenGroupRoleAddRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenGroupRoleAddRequest) GoString() string {
	return s.String()
}

func (s *OpenGroupRoleAddRequest) SetOpenConversationId(v string) *OpenGroupRoleAddRequest {
	s.OpenConversationId = &v
	return s
}

func (s *OpenGroupRoleAddRequest) SetRoleName(v string) *OpenGroupRoleAddRequest {
	s.RoleName = &v
	return s
}

func (s *OpenGroupRoleAddRequest) SetUserId(v string) *OpenGroupRoleAddRequest {
	s.UserId = &v
	return s
}

type OpenGroupRoleAddResponseBody struct {
	Result  *OpenGroupRoleAddResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                               `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OpenGroupRoleAddResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenGroupRoleAddResponseBody) GoString() string {
	return s.String()
}

func (s *OpenGroupRoleAddResponseBody) SetResult(v *OpenGroupRoleAddResponseBodyResult) *OpenGroupRoleAddResponseBody {
	s.Result = v
	return s
}

func (s *OpenGroupRoleAddResponseBody) SetSuccess(v bool) *OpenGroupRoleAddResponseBody {
	s.Success = &v
	return s
}

type OpenGroupRoleAddResponseBodyResult struct {
	OpenRoleId *string `json:"openRoleId,omitempty" xml:"openRoleId,omitempty"`
}

func (s OpenGroupRoleAddResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s OpenGroupRoleAddResponseBodyResult) GoString() string {
	return s.String()
}

func (s *OpenGroupRoleAddResponseBodyResult) SetOpenRoleId(v string) *OpenGroupRoleAddResponseBodyResult {
	s.OpenRoleId = &v
	return s
}

type OpenGroupRoleAddResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenGroupRoleAddResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenGroupRoleAddResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenGroupRoleAddResponse) GoString() string {
	return s.String()
}

func (s *OpenGroupRoleAddResponse) SetHeaders(v map[string]*string) *OpenGroupRoleAddResponse {
	s.Headers = v
	return s
}

func (s *OpenGroupRoleAddResponse) SetStatusCode(v int32) *OpenGroupRoleAddResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenGroupRoleAddResponse) SetBody(v *OpenGroupRoleAddResponseBody) *OpenGroupRoleAddResponse {
	s.Body = v
	return s
}

type OpenGroupRoleQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s OpenGroupRoleQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s OpenGroupRoleQueryHeaders) GoString() string {
	return s.String()
}

func (s *OpenGroupRoleQueryHeaders) SetCommonHeaders(v map[string]*string) *OpenGroupRoleQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OpenGroupRoleQueryHeaders) SetXAcsDingtalkAccessToken(v string) *OpenGroupRoleQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type OpenGroupRoleQueryRequest struct {
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s OpenGroupRoleQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenGroupRoleQueryRequest) GoString() string {
	return s.String()
}

func (s *OpenGroupRoleQueryRequest) SetOpenConversationId(v string) *OpenGroupRoleQueryRequest {
	s.OpenConversationId = &v
	return s
}

func (s *OpenGroupRoleQueryRequest) SetUserId(v string) *OpenGroupRoleQueryRequest {
	s.UserId = &v
	return s
}

type OpenGroupRoleQueryResponseBody struct {
	Result  *OpenGroupRoleQueryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                 `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OpenGroupRoleQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenGroupRoleQueryResponseBody) GoString() string {
	return s.String()
}

func (s *OpenGroupRoleQueryResponseBody) SetResult(v *OpenGroupRoleQueryResponseBodyResult) *OpenGroupRoleQueryResponseBody {
	s.Result = v
	return s
}

func (s *OpenGroupRoleQueryResponseBody) SetSuccess(v bool) *OpenGroupRoleQueryResponseBody {
	s.Success = &v
	return s
}

type OpenGroupRoleQueryResponseBodyResult struct {
	GroupRoles []*OpenGroupRoleQueryResponseBodyResultGroupRoles `json:"groupRoles,omitempty" xml:"groupRoles,omitempty" type:"Repeated"`
}

func (s OpenGroupRoleQueryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s OpenGroupRoleQueryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *OpenGroupRoleQueryResponseBodyResult) SetGroupRoles(v []*OpenGroupRoleQueryResponseBodyResultGroupRoles) *OpenGroupRoleQueryResponseBodyResult {
	s.GroupRoles = v
	return s
}

type OpenGroupRoleQueryResponseBodyResultGroupRoles struct {
	OpenRoleId *string `json:"openRoleId,omitempty" xml:"openRoleId,omitempty"`
	RoleName   *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
}

func (s OpenGroupRoleQueryResponseBodyResultGroupRoles) String() string {
	return tea.Prettify(s)
}

func (s OpenGroupRoleQueryResponseBodyResultGroupRoles) GoString() string {
	return s.String()
}

func (s *OpenGroupRoleQueryResponseBodyResultGroupRoles) SetOpenRoleId(v string) *OpenGroupRoleQueryResponseBodyResultGroupRoles {
	s.OpenRoleId = &v
	return s
}

func (s *OpenGroupRoleQueryResponseBodyResultGroupRoles) SetRoleName(v string) *OpenGroupRoleQueryResponseBodyResultGroupRoles {
	s.RoleName = &v
	return s
}

type OpenGroupRoleQueryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenGroupRoleQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenGroupRoleQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenGroupRoleQueryResponse) GoString() string {
	return s.String()
}

func (s *OpenGroupRoleQueryResponse) SetHeaders(v map[string]*string) *OpenGroupRoleQueryResponse {
	s.Headers = v
	return s
}

func (s *OpenGroupRoleQueryResponse) SetStatusCode(v int32) *OpenGroupRoleQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenGroupRoleQueryResponse) SetBody(v *OpenGroupRoleQueryResponseBody) *OpenGroupRoleQueryResponse {
	s.Body = v
	return s
}

type OpenGroupRoleRemoveHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s OpenGroupRoleRemoveHeaders) String() string {
	return tea.Prettify(s)
}

func (s OpenGroupRoleRemoveHeaders) GoString() string {
	return s.String()
}

func (s *OpenGroupRoleRemoveHeaders) SetCommonHeaders(v map[string]*string) *OpenGroupRoleRemoveHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OpenGroupRoleRemoveHeaders) SetXAcsDingtalkAccessToken(v string) *OpenGroupRoleRemoveHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type OpenGroupRoleRemoveRequest struct {
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	OpenRoleId *string `json:"openRoleId,omitempty" xml:"openRoleId,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s OpenGroupRoleRemoveRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenGroupRoleRemoveRequest) GoString() string {
	return s.String()
}

func (s *OpenGroupRoleRemoveRequest) SetOpenConversationId(v string) *OpenGroupRoleRemoveRequest {
	s.OpenConversationId = &v
	return s
}

func (s *OpenGroupRoleRemoveRequest) SetOpenRoleId(v string) *OpenGroupRoleRemoveRequest {
	s.OpenRoleId = &v
	return s
}

func (s *OpenGroupRoleRemoveRequest) SetUserId(v string) *OpenGroupRoleRemoveRequest {
	s.UserId = &v
	return s
}

type OpenGroupRoleRemoveResponseBody struct {
	Result  *bool `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OpenGroupRoleRemoveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenGroupRoleRemoveResponseBody) GoString() string {
	return s.String()
}

func (s *OpenGroupRoleRemoveResponseBody) SetResult(v bool) *OpenGroupRoleRemoveResponseBody {
	s.Result = &v
	return s
}

func (s *OpenGroupRoleRemoveResponseBody) SetSuccess(v bool) *OpenGroupRoleRemoveResponseBody {
	s.Success = &v
	return s
}

type OpenGroupRoleRemoveResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenGroupRoleRemoveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenGroupRoleRemoveResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenGroupRoleRemoveResponse) GoString() string {
	return s.String()
}

func (s *OpenGroupRoleRemoveResponse) SetHeaders(v map[string]*string) *OpenGroupRoleRemoveResponse {
	s.Headers = v
	return s
}

func (s *OpenGroupRoleRemoveResponse) SetStatusCode(v int32) *OpenGroupRoleRemoveResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenGroupRoleRemoveResponse) SetBody(v *OpenGroupRoleRemoveResponseBody) *OpenGroupRoleRemoveResponse {
	s.Body = v
	return s
}

type OpenGroupRoleUpdateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s OpenGroupRoleUpdateHeaders) String() string {
	return tea.Prettify(s)
}

func (s OpenGroupRoleUpdateHeaders) GoString() string {
	return s.String()
}

func (s *OpenGroupRoleUpdateHeaders) SetCommonHeaders(v map[string]*string) *OpenGroupRoleUpdateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OpenGroupRoleUpdateHeaders) SetXAcsDingtalkAccessToken(v string) *OpenGroupRoleUpdateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type OpenGroupRoleUpdateRequest struct {
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	OpenRoleId *string `json:"openRoleId,omitempty" xml:"openRoleId,omitempty"`
	// This parameter is required.
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s OpenGroupRoleUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenGroupRoleUpdateRequest) GoString() string {
	return s.String()
}

func (s *OpenGroupRoleUpdateRequest) SetOpenConversationId(v string) *OpenGroupRoleUpdateRequest {
	s.OpenConversationId = &v
	return s
}

func (s *OpenGroupRoleUpdateRequest) SetOpenRoleId(v string) *OpenGroupRoleUpdateRequest {
	s.OpenRoleId = &v
	return s
}

func (s *OpenGroupRoleUpdateRequest) SetRoleName(v string) *OpenGroupRoleUpdateRequest {
	s.RoleName = &v
	return s
}

func (s *OpenGroupRoleUpdateRequest) SetUserId(v string) *OpenGroupRoleUpdateRequest {
	s.UserId = &v
	return s
}

type OpenGroupRoleUpdateResponseBody struct {
	Result  *bool `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OpenGroupRoleUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenGroupRoleUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *OpenGroupRoleUpdateResponseBody) SetResult(v bool) *OpenGroupRoleUpdateResponseBody {
	s.Result = &v
	return s
}

func (s *OpenGroupRoleUpdateResponseBody) SetSuccess(v bool) *OpenGroupRoleUpdateResponseBody {
	s.Success = &v
	return s
}

type OpenGroupRoleUpdateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenGroupRoleUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenGroupRoleUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenGroupRoleUpdateResponse) GoString() string {
	return s.String()
}

func (s *OpenGroupRoleUpdateResponse) SetHeaders(v map[string]*string) *OpenGroupRoleUpdateResponse {
	s.Headers = v
	return s
}

func (s *OpenGroupRoleUpdateResponse) SetStatusCode(v int32) *OpenGroupRoleUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenGroupRoleUpdateResponse) SetBody(v *OpenGroupRoleUpdateResponseBody) *OpenGroupRoleUpdateResponse {
	s.Body = v
	return s
}

type OpenGroupUserRoleQueryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s OpenGroupUserRoleQueryHeaders) String() string {
	return tea.Prettify(s)
}

func (s OpenGroupUserRoleQueryHeaders) GoString() string {
	return s.String()
}

func (s *OpenGroupUserRoleQueryHeaders) SetCommonHeaders(v map[string]*string) *OpenGroupUserRoleQueryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OpenGroupUserRoleQueryHeaders) SetXAcsDingtalkAccessToken(v string) *OpenGroupUserRoleQueryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type OpenGroupUserRoleQueryRequest struct {
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// This parameter is required.
	ViewedUserId *string `json:"viewedUserId,omitempty" xml:"viewedUserId,omitempty"`
}

func (s OpenGroupUserRoleQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenGroupUserRoleQueryRequest) GoString() string {
	return s.String()
}

func (s *OpenGroupUserRoleQueryRequest) SetOpenConversationId(v string) *OpenGroupUserRoleQueryRequest {
	s.OpenConversationId = &v
	return s
}

func (s *OpenGroupUserRoleQueryRequest) SetUserId(v string) *OpenGroupUserRoleQueryRequest {
	s.UserId = &v
	return s
}

func (s *OpenGroupUserRoleQueryRequest) SetViewedUserId(v string) *OpenGroupUserRoleQueryRequest {
	s.ViewedUserId = &v
	return s
}

type OpenGroupUserRoleQueryResponseBody struct {
	Result  *OpenGroupUserRoleQueryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OpenGroupUserRoleQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenGroupUserRoleQueryResponseBody) GoString() string {
	return s.String()
}

func (s *OpenGroupUserRoleQueryResponseBody) SetResult(v *OpenGroupUserRoleQueryResponseBodyResult) *OpenGroupUserRoleQueryResponseBody {
	s.Result = v
	return s
}

func (s *OpenGroupUserRoleQueryResponseBody) SetSuccess(v bool) *OpenGroupUserRoleQueryResponseBody {
	s.Success = &v
	return s
}

type OpenGroupUserRoleQueryResponseBodyResult struct {
	GroupRoles []*OpenGroupUserRoleQueryResponseBodyResultGroupRoles `json:"groupRoles,omitempty" xml:"groupRoles,omitempty" type:"Repeated"`
}

func (s OpenGroupUserRoleQueryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s OpenGroupUserRoleQueryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *OpenGroupUserRoleQueryResponseBodyResult) SetGroupRoles(v []*OpenGroupUserRoleQueryResponseBodyResultGroupRoles) *OpenGroupUserRoleQueryResponseBodyResult {
	s.GroupRoles = v
	return s
}

type OpenGroupUserRoleQueryResponseBodyResultGroupRoles struct {
	// example:
	//
	// rolexxxxxxx
	OpenRoleId *string `json:"openRoleId,omitempty" xml:"openRoleId,omitempty"`
	// example:
	//
	// 小美
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
}

func (s OpenGroupUserRoleQueryResponseBodyResultGroupRoles) String() string {
	return tea.Prettify(s)
}

func (s OpenGroupUserRoleQueryResponseBodyResultGroupRoles) GoString() string {
	return s.String()
}

func (s *OpenGroupUserRoleQueryResponseBodyResultGroupRoles) SetOpenRoleId(v string) *OpenGroupUserRoleQueryResponseBodyResultGroupRoles {
	s.OpenRoleId = &v
	return s
}

func (s *OpenGroupUserRoleQueryResponseBodyResultGroupRoles) SetRoleName(v string) *OpenGroupUserRoleQueryResponseBodyResultGroupRoles {
	s.RoleName = &v
	return s
}

type OpenGroupUserRoleQueryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenGroupUserRoleQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenGroupUserRoleQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenGroupUserRoleQueryResponse) GoString() string {
	return s.String()
}

func (s *OpenGroupUserRoleQueryResponse) SetHeaders(v map[string]*string) *OpenGroupUserRoleQueryResponse {
	s.Headers = v
	return s
}

func (s *OpenGroupUserRoleQueryResponse) SetStatusCode(v int32) *OpenGroupUserRoleQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenGroupUserRoleQueryResponse) SetBody(v *OpenGroupUserRoleQueryResponseBody) *OpenGroupUserRoleQueryResponse {
	s.Body = v
	return s
}

type OpenInnerGroupTransferToDeptGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s OpenInnerGroupTransferToDeptGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s OpenInnerGroupTransferToDeptGroupHeaders) GoString() string {
	return s.String()
}

func (s *OpenInnerGroupTransferToDeptGroupHeaders) SetCommonHeaders(v map[string]*string) *OpenInnerGroupTransferToDeptGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OpenInnerGroupTransferToDeptGroupHeaders) SetXAcsDingtalkAccessToken(v string) *OpenInnerGroupTransferToDeptGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type OpenInnerGroupTransferToDeptGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 100
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cidD2y*****==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s OpenInnerGroupTransferToDeptGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenInnerGroupTransferToDeptGroupRequest) GoString() string {
	return s.String()
}

func (s *OpenInnerGroupTransferToDeptGroupRequest) SetDeptId(v int64) *OpenInnerGroupTransferToDeptGroupRequest {
	s.DeptId = &v
	return s
}

func (s *OpenInnerGroupTransferToDeptGroupRequest) SetOpenConversationId(v string) *OpenInnerGroupTransferToDeptGroupRequest {
	s.OpenConversationId = &v
	return s
}

type OpenInnerGroupTransferToDeptGroupResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OpenInnerGroupTransferToDeptGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenInnerGroupTransferToDeptGroupResponseBody) GoString() string {
	return s.String()
}

func (s *OpenInnerGroupTransferToDeptGroupResponseBody) SetSuccess(v bool) *OpenInnerGroupTransferToDeptGroupResponseBody {
	s.Success = &v
	return s
}

type OpenInnerGroupTransferToDeptGroupResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenInnerGroupTransferToDeptGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenInnerGroupTransferToDeptGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenInnerGroupTransferToDeptGroupResponse) GoString() string {
	return s.String()
}

func (s *OpenInnerGroupTransferToDeptGroupResponse) SetHeaders(v map[string]*string) *OpenInnerGroupTransferToDeptGroupResponse {
	s.Headers = v
	return s
}

func (s *OpenInnerGroupTransferToDeptGroupResponse) SetStatusCode(v int32) *OpenInnerGroupTransferToDeptGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenInnerGroupTransferToDeptGroupResponse) SetBody(v *OpenInnerGroupTransferToDeptGroupResponseBody) *OpenInnerGroupTransferToDeptGroupResponse {
	s.Body = v
	return s
}

type OpenSearchGroupListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s OpenSearchGroupListHeaders) String() string {
	return tea.Prettify(s)
}

func (s OpenSearchGroupListHeaders) GoString() string {
	return s.String()
}

func (s *OpenSearchGroupListHeaders) SetCommonHeaders(v map[string]*string) *OpenSearchGroupListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OpenSearchGroupListHeaders) SetXAcsDingtalkAccessToken(v string) *OpenSearchGroupListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type OpenSearchGroupListRequest struct {
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s OpenSearchGroupListRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenSearchGroupListRequest) GoString() string {
	return s.String()
}

func (s *OpenSearchGroupListRequest) SetKeyword(v string) *OpenSearchGroupListRequest {
	s.Keyword = &v
	return s
}

func (s *OpenSearchGroupListRequest) SetUserId(v string) *OpenSearchGroupListRequest {
	s.UserId = &v
	return s
}

type OpenSearchGroupListResponseBody struct {
	Result  *OpenSearchGroupListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OpenSearchGroupListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenSearchGroupListResponseBody) GoString() string {
	return s.String()
}

func (s *OpenSearchGroupListResponseBody) SetResult(v *OpenSearchGroupListResponseBodyResult) *OpenSearchGroupListResponseBody {
	s.Result = v
	return s
}

func (s *OpenSearchGroupListResponseBody) SetSuccess(v bool) *OpenSearchGroupListResponseBody {
	s.Success = &v
	return s
}

type OpenSearchGroupListResponseBodyResult struct {
	GroupList []*OpenSearchGroupListResponseBodyResultGroupList `json:"groupList,omitempty" xml:"groupList,omitempty" type:"Repeated"`
}

func (s OpenSearchGroupListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s OpenSearchGroupListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *OpenSearchGroupListResponseBodyResult) SetGroupList(v []*OpenSearchGroupListResponseBodyResultGroupList) *OpenSearchGroupListResponseBodyResult {
	s.GroupList = v
	return s
}

type OpenSearchGroupListResponseBodyResultGroupList struct {
	Icon               *string `json:"icon,omitempty" xml:"icon,omitempty"`
	MemberCount        *int32  `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	Tag                *string `json:"tag,omitempty" xml:"tag,omitempty"`
	Title              *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s OpenSearchGroupListResponseBodyResultGroupList) String() string {
	return tea.Prettify(s)
}

func (s OpenSearchGroupListResponseBodyResultGroupList) GoString() string {
	return s.String()
}

func (s *OpenSearchGroupListResponseBodyResultGroupList) SetIcon(v string) *OpenSearchGroupListResponseBodyResultGroupList {
	s.Icon = &v
	return s
}

func (s *OpenSearchGroupListResponseBodyResultGroupList) SetMemberCount(v int32) *OpenSearchGroupListResponseBodyResultGroupList {
	s.MemberCount = &v
	return s
}

func (s *OpenSearchGroupListResponseBodyResultGroupList) SetOpenConversationId(v string) *OpenSearchGroupListResponseBodyResultGroupList {
	s.OpenConversationId = &v
	return s
}

func (s *OpenSearchGroupListResponseBodyResultGroupList) SetTag(v string) *OpenSearchGroupListResponseBodyResultGroupList {
	s.Tag = &v
	return s
}

func (s *OpenSearchGroupListResponseBodyResultGroupList) SetTitle(v string) *OpenSearchGroupListResponseBodyResultGroupList {
	s.Title = &v
	return s
}

type OpenSearchGroupListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenSearchGroupListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenSearchGroupListResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenSearchGroupListResponse) GoString() string {
	return s.String()
}

func (s *OpenSearchGroupListResponse) SetHeaders(v map[string]*string) *OpenSearchGroupListResponse {
	s.Headers = v
	return s
}

func (s *OpenSearchGroupListResponse) SetStatusCode(v int32) *OpenSearchGroupListResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenSearchGroupListResponse) SetBody(v *OpenSearchGroupListResponseBody) *OpenSearchGroupListResponse {
	s.Body = v
	return s
}

type OpenUserSendCardMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s OpenUserSendCardMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s OpenUserSendCardMessageHeaders) GoString() string {
	return s.String()
}

func (s *OpenUserSendCardMessageHeaders) SetCommonHeaders(v map[string]*string) *OpenUserSendCardMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OpenUserSendCardMessageHeaders) SetXAcsDingtalkAccessToken(v string) *OpenUserSendCardMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type OpenUserSendCardMessageRequest struct {
	// This parameter is required.
	CardContent        *OpenUserSendCardMessageRequestCardContent `json:"cardContent,omitempty" xml:"cardContent,omitempty" type:"Struct"`
	OpenConversationId *string                                    `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	ReceiveUserId      *string                                    `json:"receiveUserId,omitempty" xml:"receiveUserId,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s OpenUserSendCardMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenUserSendCardMessageRequest) GoString() string {
	return s.String()
}

func (s *OpenUserSendCardMessageRequest) SetCardContent(v *OpenUserSendCardMessageRequestCardContent) *OpenUserSendCardMessageRequest {
	s.CardContent = v
	return s
}

func (s *OpenUserSendCardMessageRequest) SetOpenConversationId(v string) *OpenUserSendCardMessageRequest {
	s.OpenConversationId = &v
	return s
}

func (s *OpenUserSendCardMessageRequest) SetReceiveUserId(v string) *OpenUserSendCardMessageRequest {
	s.ReceiveUserId = &v
	return s
}

func (s *OpenUserSendCardMessageRequest) SetUserId(v string) *OpenUserSendCardMessageRequest {
	s.UserId = &v
	return s
}

type OpenUserSendCardMessageRequestCardContent struct {
	// This parameter is required.
	LastMessage *string `json:"lastMessage,omitempty" xml:"lastMessage,omitempty"`
	// This parameter is required.
	OutTrackId *string `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
}

func (s OpenUserSendCardMessageRequestCardContent) String() string {
	return tea.Prettify(s)
}

func (s OpenUserSendCardMessageRequestCardContent) GoString() string {
	return s.String()
}

func (s *OpenUserSendCardMessageRequestCardContent) SetLastMessage(v string) *OpenUserSendCardMessageRequestCardContent {
	s.LastMessage = &v
	return s
}

func (s *OpenUserSendCardMessageRequestCardContent) SetOutTrackId(v string) *OpenUserSendCardMessageRequestCardContent {
	s.OutTrackId = &v
	return s
}

type OpenUserSendCardMessageResponseBody struct {
	Result  *OpenUserSendCardMessageResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OpenUserSendCardMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenUserSendCardMessageResponseBody) GoString() string {
	return s.String()
}

func (s *OpenUserSendCardMessageResponseBody) SetResult(v *OpenUserSendCardMessageResponseBodyResult) *OpenUserSendCardMessageResponseBody {
	s.Result = v
	return s
}

func (s *OpenUserSendCardMessageResponseBody) SetSuccess(v bool) *OpenUserSendCardMessageResponseBody {
	s.Success = &v
	return s
}

type OpenUserSendCardMessageResponseBodyResult struct {
	OpenTaskId *string `json:"openTaskId,omitempty" xml:"openTaskId,omitempty"`
}

func (s OpenUserSendCardMessageResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s OpenUserSendCardMessageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *OpenUserSendCardMessageResponseBodyResult) SetOpenTaskId(v string) *OpenUserSendCardMessageResponseBodyResult {
	s.OpenTaskId = &v
	return s
}

type OpenUserSendCardMessageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenUserSendCardMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenUserSendCardMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenUserSendCardMessageResponse) GoString() string {
	return s.String()
}

func (s *OpenUserSendCardMessageResponse) SetHeaders(v map[string]*string) *OpenUserSendCardMessageResponse {
	s.Headers = v
	return s
}

func (s *OpenUserSendCardMessageResponse) SetStatusCode(v int32) *OpenUserSendCardMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenUserSendCardMessageResponse) SetBody(v *OpenUserSendCardMessageResponseBody) *OpenUserSendCardMessageResponse {
	s.Body = v
	return s
}

type PersonalSendCardMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PersonalSendCardMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s PersonalSendCardMessageHeaders) GoString() string {
	return s.String()
}

func (s *PersonalSendCardMessageHeaders) SetCommonHeaders(v map[string]*string) *PersonalSendCardMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PersonalSendCardMessageHeaders) SetXAcsDingtalkAccessToken(v string) *PersonalSendCardMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PersonalSendCardMessageRequest struct {
	AtUserIds          []*string                                  `json:"atUserIds,omitempty" xml:"atUserIds,omitempty" type:"Repeated"`
	CardContent        *PersonalSendCardMessageRequestCardContent `json:"cardContent,omitempty" xml:"cardContent,omitempty" type:"Struct"`
	OpenConversationId *string                                    `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	ReceiveUserId      *string                                    `json:"receiveUserId,omitempty" xml:"receiveUserId,omitempty"`
}

func (s PersonalSendCardMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s PersonalSendCardMessageRequest) GoString() string {
	return s.String()
}

func (s *PersonalSendCardMessageRequest) SetAtUserIds(v []*string) *PersonalSendCardMessageRequest {
	s.AtUserIds = v
	return s
}

func (s *PersonalSendCardMessageRequest) SetCardContent(v *PersonalSendCardMessageRequestCardContent) *PersonalSendCardMessageRequest {
	s.CardContent = v
	return s
}

func (s *PersonalSendCardMessageRequest) SetOpenConversationId(v string) *PersonalSendCardMessageRequest {
	s.OpenConversationId = &v
	return s
}

func (s *PersonalSendCardMessageRequest) SetReceiveUserId(v string) *PersonalSendCardMessageRequest {
	s.ReceiveUserId = &v
	return s
}

type PersonalSendCardMessageRequestCardContent struct {
	LastMessage *string `json:"lastMessage,omitempty" xml:"lastMessage,omitempty"`
	OutTrackId  *string `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
}

func (s PersonalSendCardMessageRequestCardContent) String() string {
	return tea.Prettify(s)
}

func (s PersonalSendCardMessageRequestCardContent) GoString() string {
	return s.String()
}

func (s *PersonalSendCardMessageRequestCardContent) SetLastMessage(v string) *PersonalSendCardMessageRequestCardContent {
	s.LastMessage = &v
	return s
}

func (s *PersonalSendCardMessageRequestCardContent) SetOutTrackId(v string) *PersonalSendCardMessageRequestCardContent {
	s.OutTrackId = &v
	return s
}

type PersonalSendCardMessageResponseBody struct {
	Result  *PersonalSendCardMessageResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PersonalSendCardMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PersonalSendCardMessageResponseBody) GoString() string {
	return s.String()
}

func (s *PersonalSendCardMessageResponseBody) SetResult(v *PersonalSendCardMessageResponseBodyResult) *PersonalSendCardMessageResponseBody {
	s.Result = v
	return s
}

func (s *PersonalSendCardMessageResponseBody) SetSuccess(v bool) *PersonalSendCardMessageResponseBody {
	s.Success = &v
	return s
}

type PersonalSendCardMessageResponseBodyResult struct {
	OpenTaskId *string `json:"openTaskId,omitempty" xml:"openTaskId,omitempty"`
}

func (s PersonalSendCardMessageResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PersonalSendCardMessageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PersonalSendCardMessageResponseBodyResult) SetOpenTaskId(v string) *PersonalSendCardMessageResponseBodyResult {
	s.OpenTaskId = &v
	return s
}

type PersonalSendCardMessageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PersonalSendCardMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PersonalSendCardMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s PersonalSendCardMessageResponse) GoString() string {
	return s.String()
}

func (s *PersonalSendCardMessageResponse) SetHeaders(v map[string]*string) *PersonalSendCardMessageResponse {
	s.Headers = v
	return s
}

func (s *PersonalSendCardMessageResponse) SetStatusCode(v int32) *PersonalSendCardMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *PersonalSendCardMessageResponse) SetBody(v *PersonalSendCardMessageResponseBody) *PersonalSendCardMessageResponse {
	s.Body = v
	return s
}

type QueryGroupInfoByMemberAuthHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryGroupInfoByMemberAuthHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoByMemberAuthHeaders) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoByMemberAuthHeaders) SetCommonHeaders(v map[string]*string) *QueryGroupInfoByMemberAuthHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGroupInfoByMemberAuthHeaders) SetXAcsDingtalkAccessToken(v string) *QueryGroupInfoByMemberAuthHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryGroupInfoByMemberAuthRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// COOLAPP-XXX
	CoolAppCode *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cidXXX
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s QueryGroupInfoByMemberAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoByMemberAuthRequest) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoByMemberAuthRequest) SetCoolAppCode(v string) *QueryGroupInfoByMemberAuthRequest {
	s.CoolAppCode = &v
	return s
}

func (s *QueryGroupInfoByMemberAuthRequest) SetOpenConversationId(v string) *QueryGroupInfoByMemberAuthRequest {
	s.OpenConversationId = &v
	return s
}

type QueryGroupInfoByMemberAuthResponseBody struct {
	// example:
	//
	// 99
	MemberCount *int32 `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
}

func (s QueryGroupInfoByMemberAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoByMemberAuthResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoByMemberAuthResponseBody) SetMemberCount(v int32) *QueryGroupInfoByMemberAuthResponseBody {
	s.MemberCount = &v
	return s
}

type QueryGroupInfoByMemberAuthResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGroupInfoByMemberAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGroupInfoByMemberAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoByMemberAuthResponse) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoByMemberAuthResponse) SetHeaders(v map[string]*string) *QueryGroupInfoByMemberAuthResponse {
	s.Headers = v
	return s
}

func (s *QueryGroupInfoByMemberAuthResponse) SetStatusCode(v int32) *QueryGroupInfoByMemberAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGroupInfoByMemberAuthResponse) SetBody(v *QueryGroupInfoByMemberAuthResponseBody) *QueryGroupInfoByMemberAuthResponse {
	s.Body = v
	return s
}

type QueryGroupMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryGroupMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMemberHeaders) GoString() string {
	return s.String()
}

func (s *QueryGroupMemberHeaders) SetCommonHeaders(v map[string]*string) *QueryGroupMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGroupMemberHeaders) SetXAcsDingtalkAccessToken(v string) *QueryGroupMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryGroupMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 14da****2760
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s QueryGroupMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *QueryGroupMemberRequest) SetOpenConversationId(v string) *QueryGroupMemberRequest {
	s.OpenConversationId = &v
	return s
}

type QueryGroupMemberResponseBody struct {
	// This parameter is required.
	GroupMembers []*QueryGroupMemberResponseBodyGroupMembers `json:"groupMembers,omitempty" xml:"groupMembers,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 14da****2760
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s QueryGroupMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGroupMemberResponseBody) SetGroupMembers(v []*QueryGroupMemberResponseBodyGroupMembers) *QueryGroupMemberResponseBody {
	s.GroupMembers = v
	return s
}

func (s *QueryGroupMemberResponseBody) SetOpenConversationId(v string) *QueryGroupMemberResponseBody {
	s.OpenConversationId = &v
	return s
}

type QueryGroupMemberResponseBodyGroupMembers struct {
	// example:
	//
	// http://****.png
	GroupMemberAvatar *string `json:"groupMemberAvatar,omitempty" xml:"groupMemberAvatar,omitempty"`
	// example:
	//
	// 认真工作,快乐生活
	GroupMemberDynamics *string `json:"groupMemberDynamics,omitempty" xml:"groupMemberDynamics,omitempty"`
	// example:
	//
	// 1107****2120
	GroupMemberId *string `json:"groupMemberId,omitempty" xml:"groupMemberId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Foo
	GroupMemberName *string `json:"groupMemberName,omitempty" xml:"groupMemberName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	GroupMemberType *int32 `json:"groupMemberType,omitempty" xml:"groupMemberType,omitempty"`
}

func (s QueryGroupMemberResponseBodyGroupMembers) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMemberResponseBodyGroupMembers) GoString() string {
	return s.String()
}

func (s *QueryGroupMemberResponseBodyGroupMembers) SetGroupMemberAvatar(v string) *QueryGroupMemberResponseBodyGroupMembers {
	s.GroupMemberAvatar = &v
	return s
}

func (s *QueryGroupMemberResponseBodyGroupMembers) SetGroupMemberDynamics(v string) *QueryGroupMemberResponseBodyGroupMembers {
	s.GroupMemberDynamics = &v
	return s
}

func (s *QueryGroupMemberResponseBodyGroupMembers) SetGroupMemberId(v string) *QueryGroupMemberResponseBodyGroupMembers {
	s.GroupMemberId = &v
	return s
}

func (s *QueryGroupMemberResponseBodyGroupMembers) SetGroupMemberName(v string) *QueryGroupMemberResponseBodyGroupMembers {
	s.GroupMemberName = &v
	return s
}

func (s *QueryGroupMemberResponseBodyGroupMembers) SetGroupMemberType(v int32) *QueryGroupMemberResponseBodyGroupMembers {
	s.GroupMemberType = &v
	return s
}

type QueryGroupMemberResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGroupMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *QueryGroupMemberResponse) SetHeaders(v map[string]*string) *QueryGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *QueryGroupMemberResponse) SetStatusCode(v int32) *QueryGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGroupMemberResponse) SetBody(v *QueryGroupMemberResponseBody) *QueryGroupMemberResponse {
	s.Body = v
	return s
}

type QueryGroupMemberByMemberAuthHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryGroupMemberByMemberAuthHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMemberByMemberAuthHeaders) GoString() string {
	return s.String()
}

func (s *QueryGroupMemberByMemberAuthHeaders) SetCommonHeaders(v map[string]*string) *QueryGroupMemberByMemberAuthHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGroupMemberByMemberAuthHeaders) SetXAcsDingtalkAccessToken(v string) *QueryGroupMemberByMemberAuthHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryGroupMemberByMemberAuthRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// COOLAPP-XXX
	CoolAppCode *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cidXXX
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s QueryGroupMemberByMemberAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMemberByMemberAuthRequest) GoString() string {
	return s.String()
}

func (s *QueryGroupMemberByMemberAuthRequest) SetCoolAppCode(v string) *QueryGroupMemberByMemberAuthRequest {
	s.CoolAppCode = &v
	return s
}

func (s *QueryGroupMemberByMemberAuthRequest) SetOpenConversationId(v string) *QueryGroupMemberByMemberAuthRequest {
	s.OpenConversationId = &v
	return s
}

type QueryGroupMemberByMemberAuthResponseBody struct {
	GroupMemberList []*QueryGroupMemberByMemberAuthResponseBodyGroupMemberList `json:"groupMemberList,omitempty" xml:"groupMemberList,omitempty" type:"Repeated"`
}

func (s QueryGroupMemberByMemberAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMemberByMemberAuthResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGroupMemberByMemberAuthResponseBody) SetGroupMemberList(v []*QueryGroupMemberByMemberAuthResponseBodyGroupMemberList) *QueryGroupMemberByMemberAuthResponseBody {
	s.GroupMemberList = v
	return s
}

type QueryGroupMemberByMemberAuthResponseBodyGroupMemberList struct {
	// example:
	//
	// 张三
	GroupNickName *string `json:"groupNickName,omitempty" xml:"groupNickName,omitempty"`
	// example:
	//
	// 张某某
	OrgName *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	// example:
	//
	// https://xxx
	ProfilePhotoUrl *string `json:"profilePhotoUrl,omitempty" xml:"profilePhotoUrl,omitempty"`
	// example:
	//
	// xxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryGroupMemberByMemberAuthResponseBodyGroupMemberList) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMemberByMemberAuthResponseBodyGroupMemberList) GoString() string {
	return s.String()
}

func (s *QueryGroupMemberByMemberAuthResponseBodyGroupMemberList) SetGroupNickName(v string) *QueryGroupMemberByMemberAuthResponseBodyGroupMemberList {
	s.GroupNickName = &v
	return s
}

func (s *QueryGroupMemberByMemberAuthResponseBodyGroupMemberList) SetOrgName(v string) *QueryGroupMemberByMemberAuthResponseBodyGroupMemberList {
	s.OrgName = &v
	return s
}

func (s *QueryGroupMemberByMemberAuthResponseBodyGroupMemberList) SetProfilePhotoUrl(v string) *QueryGroupMemberByMemberAuthResponseBodyGroupMemberList {
	s.ProfilePhotoUrl = &v
	return s
}

func (s *QueryGroupMemberByMemberAuthResponseBodyGroupMemberList) SetUserId(v string) *QueryGroupMemberByMemberAuthResponseBodyGroupMemberList {
	s.UserId = &v
	return s
}

type QueryGroupMemberByMemberAuthResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGroupMemberByMemberAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGroupMemberByMemberAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMemberByMemberAuthResponse) GoString() string {
	return s.String()
}

func (s *QueryGroupMemberByMemberAuthResponse) SetHeaders(v map[string]*string) *QueryGroupMemberByMemberAuthResponse {
	s.Headers = v
	return s
}

func (s *QueryGroupMemberByMemberAuthResponse) SetStatusCode(v int32) *QueryGroupMemberByMemberAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGroupMemberByMemberAuthResponse) SetBody(v *QueryGroupMemberByMemberAuthResponseBody) *QueryGroupMemberByMemberAuthResponse {
	s.Body = v
	return s
}

type QueryGroupMuteStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryGroupMuteStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMuteStatusHeaders) GoString() string {
	return s.String()
}

func (s *QueryGroupMuteStatusHeaders) SetCommonHeaders(v map[string]*string) *QueryGroupMuteStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGroupMuteStatusHeaders) SetXAcsDingtalkAccessToken(v string) *QueryGroupMuteStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryGroupMuteStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cidCtneF+XyQjcyF2ROdgSeIg==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 004741900
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryGroupMuteStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMuteStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryGroupMuteStatusRequest) SetOpenConversationId(v string) *QueryGroupMuteStatusRequest {
	s.OpenConversationId = &v
	return s
}

func (s *QueryGroupMuteStatusRequest) SetUserId(v string) *QueryGroupMuteStatusRequest {
	s.UserId = &v
	return s
}

type QueryGroupMuteStatusResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	GroupMuteMode *bool `json:"groupMuteMode,omitempty" xml:"groupMuteMode,omitempty"`
	// This parameter is required.
	UserMuteResult *QueryGroupMuteStatusResponseBodyUserMuteResult `json:"userMuteResult,omitempty" xml:"userMuteResult,omitempty" type:"Struct"`
}

func (s QueryGroupMuteStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMuteStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGroupMuteStatusResponseBody) SetGroupMuteMode(v bool) *QueryGroupMuteStatusResponseBody {
	s.GroupMuteMode = &v
	return s
}

func (s *QueryGroupMuteStatusResponseBody) SetUserMuteResult(v *QueryGroupMuteStatusResponseBodyUserMuteResult) *QueryGroupMuteStatusResponseBody {
	s.UserMuteResult = v
	return s
}

type QueryGroupMuteStatusResponseBodyUserMuteResult struct {
	// This parameter is required.
	//
	// example:
	//
	// 1645315682000
	MuteEndTime *int64 `json:"muteEndTime,omitempty" xml:"muteEndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1645315682000
	MuteStartTime *int64 `json:"muteStartTime,omitempty" xml:"muteStartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	UserMuteMode *bool `json:"userMuteMode,omitempty" xml:"userMuteMode,omitempty"`
}

func (s QueryGroupMuteStatusResponseBodyUserMuteResult) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMuteStatusResponseBodyUserMuteResult) GoString() string {
	return s.String()
}

func (s *QueryGroupMuteStatusResponseBodyUserMuteResult) SetMuteEndTime(v int64) *QueryGroupMuteStatusResponseBodyUserMuteResult {
	s.MuteEndTime = &v
	return s
}

func (s *QueryGroupMuteStatusResponseBodyUserMuteResult) SetMuteStartTime(v int64) *QueryGroupMuteStatusResponseBodyUserMuteResult {
	s.MuteStartTime = &v
	return s
}

func (s *QueryGroupMuteStatusResponseBodyUserMuteResult) SetUserMuteMode(v bool) *QueryGroupMuteStatusResponseBodyUserMuteResult {
	s.UserMuteMode = &v
	return s
}

type QueryGroupMuteStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGroupMuteStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGroupMuteStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupMuteStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryGroupMuteStatusResponse) SetHeaders(v map[string]*string) *QueryGroupMuteStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryGroupMuteStatusResponse) SetStatusCode(v int32) *QueryGroupMuteStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGroupMuteStatusResponse) SetBody(v *QueryGroupMuteStatusResponseBody) *QueryGroupMuteStatusResponse {
	s.Body = v
	return s
}

type QueryInnerGroupMemberListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryInnerGroupMemberListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryInnerGroupMemberListHeaders) GoString() string {
	return s.String()
}

func (s *QueryInnerGroupMemberListHeaders) SetCommonHeaders(v map[string]*string) *QueryInnerGroupMemberListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryInnerGroupMemberListHeaders) SetXAcsDingtalkAccessToken(v string) *QueryInnerGroupMemberListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryInnerGroupMemberListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	NextToken *int32 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryInnerGroupMemberListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInnerGroupMemberListRequest) GoString() string {
	return s.String()
}

func (s *QueryInnerGroupMemberListRequest) SetMaxResults(v int32) *QueryInnerGroupMemberListRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryInnerGroupMemberListRequest) SetNextToken(v int32) *QueryInnerGroupMemberListRequest {
	s.NextToken = &v
	return s
}

func (s *QueryInnerGroupMemberListRequest) SetOpenConversationId(v string) *QueryInnerGroupMemberListRequest {
	s.OpenConversationId = &v
	return s
}

func (s *QueryInnerGroupMemberListRequest) SetUserId(v string) *QueryInnerGroupMemberListRequest {
	s.UserId = &v
	return s
}

type QueryInnerGroupMemberListResponseBody struct {
	HasMore   *bool                                        `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List      []*QueryInnerGroupMemberListResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextToken *int32                                       `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Success   *bool                                        `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryInnerGroupMemberListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryInnerGroupMemberListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInnerGroupMemberListResponseBody) SetHasMore(v bool) *QueryInnerGroupMemberListResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryInnerGroupMemberListResponseBody) SetList(v []*QueryInnerGroupMemberListResponseBodyList) *QueryInnerGroupMemberListResponseBody {
	s.List = v
	return s
}

func (s *QueryInnerGroupMemberListResponseBody) SetNextToken(v int32) *QueryInnerGroupMemberListResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryInnerGroupMemberListResponseBody) SetSuccess(v bool) *QueryInnerGroupMemberListResponseBody {
	s.Success = &v
	return s
}

type QueryInnerGroupMemberListResponseBodyList struct {
	Icon     *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	UserId   *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryInnerGroupMemberListResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryInnerGroupMemberListResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryInnerGroupMemberListResponseBodyList) SetIcon(v string) *QueryInnerGroupMemberListResponseBodyList {
	s.Icon = &v
	return s
}

func (s *QueryInnerGroupMemberListResponseBodyList) SetName(v string) *QueryInnerGroupMemberListResponseBodyList {
	s.Name = &v
	return s
}

func (s *QueryInnerGroupMemberListResponseBodyList) SetNickName(v string) *QueryInnerGroupMemberListResponseBodyList {
	s.NickName = &v
	return s
}

func (s *QueryInnerGroupMemberListResponseBodyList) SetUserId(v string) *QueryInnerGroupMemberListResponseBodyList {
	s.UserId = &v
	return s
}

type QueryInnerGroupMemberListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInnerGroupMemberListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInnerGroupMemberListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryInnerGroupMemberListResponse) GoString() string {
	return s.String()
}

func (s *QueryInnerGroupMemberListResponse) SetHeaders(v map[string]*string) *QueryInnerGroupMemberListResponse {
	s.Headers = v
	return s
}

func (s *QueryInnerGroupMemberListResponse) SetStatusCode(v int32) *QueryInnerGroupMemberListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInnerGroupMemberListResponse) SetBody(v *QueryInnerGroupMemberListResponseBody) *QueryInnerGroupMemberListResponse {
	s.Body = v
	return s
}

type QueryInnerGroupRecentListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryInnerGroupRecentListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryInnerGroupRecentListHeaders) GoString() string {
	return s.String()
}

func (s *QueryInnerGroupRecentListHeaders) SetCommonHeaders(v map[string]*string) *QueryInnerGroupRecentListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryInnerGroupRecentListHeaders) SetXAcsDingtalkAccessToken(v string) *QueryInnerGroupRecentListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryInnerGroupRecentListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 015*****
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryInnerGroupRecentListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInnerGroupRecentListRequest) GoString() string {
	return s.String()
}

func (s *QueryInnerGroupRecentListRequest) SetUserId(v string) *QueryInnerGroupRecentListRequest {
	s.UserId = &v
	return s
}

type QueryInnerGroupRecentListResponseBody struct {
	GroupInfos []*QueryInnerGroupRecentListResponseBodyGroupInfos `json:"groupInfos,omitempty" xml:"groupInfos,omitempty" type:"Repeated"`
	Success    *bool                                              `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryInnerGroupRecentListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryInnerGroupRecentListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInnerGroupRecentListResponseBody) SetGroupInfos(v []*QueryInnerGroupRecentListResponseBodyGroupInfos) *QueryInnerGroupRecentListResponseBody {
	s.GroupInfos = v
	return s
}

func (s *QueryInnerGroupRecentListResponseBody) SetSuccess(v bool) *QueryInnerGroupRecentListResponseBody {
	s.Success = &v
	return s
}

type QueryInnerGroupRecentListResponseBodyGroupInfos struct {
	// example:
	//
	// https://static.xxxxxxx
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// 10
	MemberAmount *string `json:"memberAmount,omitempty" xml:"memberAmount,omitempty"`
	// example:
	//
	// cid1e*****==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// 测试群名称
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryInnerGroupRecentListResponseBodyGroupInfos) String() string {
	return tea.Prettify(s)
}

func (s QueryInnerGroupRecentListResponseBodyGroupInfos) GoString() string {
	return s.String()
}

func (s *QueryInnerGroupRecentListResponseBodyGroupInfos) SetIcon(v string) *QueryInnerGroupRecentListResponseBodyGroupInfos {
	s.Icon = &v
	return s
}

func (s *QueryInnerGroupRecentListResponseBodyGroupInfos) SetMemberAmount(v string) *QueryInnerGroupRecentListResponseBodyGroupInfos {
	s.MemberAmount = &v
	return s
}

func (s *QueryInnerGroupRecentListResponseBodyGroupInfos) SetOpenConversationId(v string) *QueryInnerGroupRecentListResponseBodyGroupInfos {
	s.OpenConversationId = &v
	return s
}

func (s *QueryInnerGroupRecentListResponseBodyGroupInfos) SetTitle(v string) *QueryInnerGroupRecentListResponseBodyGroupInfos {
	s.Title = &v
	return s
}

type QueryInnerGroupRecentListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInnerGroupRecentListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInnerGroupRecentListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryInnerGroupRecentListResponse) GoString() string {
	return s.String()
}

func (s *QueryInnerGroupRecentListResponse) SetHeaders(v map[string]*string) *QueryInnerGroupRecentListResponse {
	s.Headers = v
	return s
}

func (s *QueryInnerGroupRecentListResponse) SetStatusCode(v int32) *QueryInnerGroupRecentListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInnerGroupRecentListResponse) SetBody(v *QueryInnerGroupRecentListResponseBody) *QueryInnerGroupRecentListResponse {
	s.Body = v
	return s
}

type QueryMembersOfGroupRoleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMembersOfGroupRoleHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMembersOfGroupRoleHeaders) GoString() string {
	return s.String()
}

func (s *QueryMembersOfGroupRoleHeaders) SetCommonHeaders(v map[string]*string) *QueryMembersOfGroupRoleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMembersOfGroupRoleHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMembersOfGroupRoleHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMembersOfGroupRoleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cidXXXXXXX
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// roleXXXXX
	OpenRoleId *string `json:"openRoleId,omitempty" xml:"openRoleId,omitempty"`
	// example:
	//
	// 1621502140000
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s QueryMembersOfGroupRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMembersOfGroupRoleRequest) GoString() string {
	return s.String()
}

func (s *QueryMembersOfGroupRoleRequest) SetOpenConversationId(v string) *QueryMembersOfGroupRoleRequest {
	s.OpenConversationId = &v
	return s
}

func (s *QueryMembersOfGroupRoleRequest) SetOpenRoleId(v string) *QueryMembersOfGroupRoleRequest {
	s.OpenRoleId = &v
	return s
}

func (s *QueryMembersOfGroupRoleRequest) SetTimestamp(v int64) *QueryMembersOfGroupRoleRequest {
	s.Timestamp = &v
	return s
}

type QueryMembersOfGroupRoleResponseBody struct {
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s QueryMembersOfGroupRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMembersOfGroupRoleResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMembersOfGroupRoleResponseBody) SetUserIds(v []*string) *QueryMembersOfGroupRoleResponseBody {
	s.UserIds = v
	return s
}

type QueryMembersOfGroupRoleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMembersOfGroupRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMembersOfGroupRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMembersOfGroupRoleResponse) GoString() string {
	return s.String()
}

func (s *QueryMembersOfGroupRoleResponse) SetHeaders(v map[string]*string) *QueryMembersOfGroupRoleResponse {
	s.Headers = v
	return s
}

func (s *QueryMembersOfGroupRoleResponse) SetStatusCode(v int32) *QueryMembersOfGroupRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMembersOfGroupRoleResponse) SetBody(v *QueryMembersOfGroupRoleResponseBody) *QueryMembersOfGroupRoleResponse {
	s.Body = v
	return s
}

type QueryMessageSendResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMessageSendResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMessageSendResultHeaders) GoString() string {
	return s.String()
}

func (s *QueryMessageSendResultHeaders) SetCommonHeaders(v map[string]*string) *QueryMessageSendResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMessageSendResultHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMessageSendResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMessageSendResultRequest struct {
	// example:
	//
	// dhowhi23ohdh==
	OpenTaskId *string `json:"openTaskId,omitempty" xml:"openTaskId,omitempty"`
}

func (s QueryMessageSendResultRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMessageSendResultRequest) GoString() string {
	return s.String()
}

func (s *QueryMessageSendResultRequest) SetOpenTaskId(v string) *QueryMessageSendResultRequest {
	s.OpenTaskId = &v
	return s
}

type QueryMessageSendResultResponseBody struct {
	Result  *QueryMessageSendResultResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryMessageSendResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMessageSendResultResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMessageSendResultResponseBody) SetResult(v *QueryMessageSendResultResponseBodyResult) *QueryMessageSendResultResponseBody {
	s.Result = v
	return s
}

func (s *QueryMessageSendResultResponseBody) SetSuccess(v bool) *QueryMessageSendResultResponseBody {
	s.Success = &v
	return s
}

type QueryMessageSendResultResponseBodyResult struct {
	// example:
	//
	// msghcuh234
	OpenMessageId *string `json:"openMessageId,omitempty" xml:"openMessageId,omitempty"`
	// example:
	//
	// 1
	SendStatus *int32 `json:"sendStatus,omitempty" xml:"sendStatus,omitempty"`
}

func (s QueryMessageSendResultResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryMessageSendResultResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryMessageSendResultResponseBodyResult) SetOpenMessageId(v string) *QueryMessageSendResultResponseBodyResult {
	s.OpenMessageId = &v
	return s
}

func (s *QueryMessageSendResultResponseBodyResult) SetSendStatus(v int32) *QueryMessageSendResultResponseBodyResult {
	s.SendStatus = &v
	return s
}

type QueryMessageSendResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMessageSendResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMessageSendResultResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMessageSendResultResponse) GoString() string {
	return s.String()
}

func (s *QueryMessageSendResultResponse) SetHeaders(v map[string]*string) *QueryMessageSendResultResponse {
	s.Headers = v
	return s
}

func (s *QueryMessageSendResultResponse) SetStatusCode(v int32) *QueryMessageSendResultResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMessageSendResultResponse) SetBody(v *QueryMessageSendResultResponseBody) *QueryMessageSendResultResponse {
	s.Body = v
	return s
}

type QueryOpenConversationReceiveUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOpenConversationReceiveUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOpenConversationReceiveUserHeaders) GoString() string {
	return s.String()
}

func (s *QueryOpenConversationReceiveUserHeaders) SetCommonHeaders(v map[string]*string) *QueryOpenConversationReceiveUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOpenConversationReceiveUserHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOpenConversationReceiveUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOpenConversationReceiveUserRequest struct {
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	SendUserId         *string `json:"sendUserId,omitempty" xml:"sendUserId,omitempty"`
}

func (s QueryOpenConversationReceiveUserRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOpenConversationReceiveUserRequest) GoString() string {
	return s.String()
}

func (s *QueryOpenConversationReceiveUserRequest) SetOpenConversationId(v string) *QueryOpenConversationReceiveUserRequest {
	s.OpenConversationId = &v
	return s
}

func (s *QueryOpenConversationReceiveUserRequest) SetSendUserId(v string) *QueryOpenConversationReceiveUserRequest {
	s.SendUserId = &v
	return s
}

type QueryOpenConversationReceiveUserResponseBody struct {
	Result  *QueryOpenConversationReceiveUserResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                               `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryOpenConversationReceiveUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOpenConversationReceiveUserResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOpenConversationReceiveUserResponseBody) SetResult(v *QueryOpenConversationReceiveUserResponseBodyResult) *QueryOpenConversationReceiveUserResponseBody {
	s.Result = v
	return s
}

func (s *QueryOpenConversationReceiveUserResponseBody) SetSuccess(v bool) *QueryOpenConversationReceiveUserResponseBody {
	s.Success = &v
	return s
}

type QueryOpenConversationReceiveUserResponseBodyResult struct {
	ReceiveUser *QueryOpenConversationReceiveUserResponseBodyResultReceiveUser `json:"receiveUser,omitempty" xml:"receiveUser,omitempty" type:"Struct"`
}

func (s QueryOpenConversationReceiveUserResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryOpenConversationReceiveUserResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryOpenConversationReceiveUserResponseBodyResult) SetReceiveUser(v *QueryOpenConversationReceiveUserResponseBodyResultReceiveUser) *QueryOpenConversationReceiveUserResponseBodyResult {
	s.ReceiveUser = v
	return s
}

type QueryOpenConversationReceiveUserResponseBodyResultReceiveUser struct {
	// This parameter is required.
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryOpenConversationReceiveUserResponseBodyResultReceiveUser) String() string {
	return tea.Prettify(s)
}

func (s QueryOpenConversationReceiveUserResponseBodyResultReceiveUser) GoString() string {
	return s.String()
}

func (s *QueryOpenConversationReceiveUserResponseBodyResultReceiveUser) SetIcon(v string) *QueryOpenConversationReceiveUserResponseBodyResultReceiveUser {
	s.Icon = &v
	return s
}

func (s *QueryOpenConversationReceiveUserResponseBodyResultReceiveUser) SetName(v string) *QueryOpenConversationReceiveUserResponseBodyResultReceiveUser {
	s.Name = &v
	return s
}

func (s *QueryOpenConversationReceiveUserResponseBodyResultReceiveUser) SetNickName(v string) *QueryOpenConversationReceiveUserResponseBodyResultReceiveUser {
	s.NickName = &v
	return s
}

func (s *QueryOpenConversationReceiveUserResponseBodyResultReceiveUser) SetUserId(v string) *QueryOpenConversationReceiveUserResponseBodyResultReceiveUser {
	s.UserId = &v
	return s
}

type QueryOpenConversationReceiveUserResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOpenConversationReceiveUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOpenConversationReceiveUserResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOpenConversationReceiveUserResponse) GoString() string {
	return s.String()
}

func (s *QueryOpenConversationReceiveUserResponse) SetHeaders(v map[string]*string) *QueryOpenConversationReceiveUserResponse {
	s.Headers = v
	return s
}

func (s *QueryOpenConversationReceiveUserResponse) SetStatusCode(v int32) *QueryOpenConversationReceiveUserResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOpenConversationReceiveUserResponse) SetBody(v *QueryOpenConversationReceiveUserResponseBody) *QueryOpenConversationReceiveUserResponse {
	s.Body = v
	return s
}

type QueryOpenGroupBaseInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOpenGroupBaseInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOpenGroupBaseInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryOpenGroupBaseInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryOpenGroupBaseInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOpenGroupBaseInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOpenGroupBaseInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOpenGroupBaseInfoRequest struct {
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s QueryOpenGroupBaseInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOpenGroupBaseInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryOpenGroupBaseInfoRequest) SetOpenConversationId(v string) *QueryOpenGroupBaseInfoRequest {
	s.OpenConversationId = &v
	return s
}

type QueryOpenGroupBaseInfoResponseBody struct {
	Result  *QueryOpenGroupBaseInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryOpenGroupBaseInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOpenGroupBaseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOpenGroupBaseInfoResponseBody) SetResult(v *QueryOpenGroupBaseInfoResponseBodyResult) *QueryOpenGroupBaseInfoResponseBody {
	s.Result = v
	return s
}

func (s *QueryOpenGroupBaseInfoResponseBody) SetSuccess(v bool) *QueryOpenGroupBaseInfoResponseBody {
	s.Success = &v
	return s
}

type QueryOpenGroupBaseInfoResponseBodyResult struct {
	Icon               *string `json:"icon,omitempty" xml:"icon,omitempty"`
	MemberCount        *int32  `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	Tag                *string `json:"tag,omitempty" xml:"tag,omitempty"`
	Title              *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryOpenGroupBaseInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryOpenGroupBaseInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryOpenGroupBaseInfoResponseBodyResult) SetIcon(v string) *QueryOpenGroupBaseInfoResponseBodyResult {
	s.Icon = &v
	return s
}

func (s *QueryOpenGroupBaseInfoResponseBodyResult) SetMemberCount(v int32) *QueryOpenGroupBaseInfoResponseBodyResult {
	s.MemberCount = &v
	return s
}

func (s *QueryOpenGroupBaseInfoResponseBodyResult) SetOpenConversationId(v string) *QueryOpenGroupBaseInfoResponseBodyResult {
	s.OpenConversationId = &v
	return s
}

func (s *QueryOpenGroupBaseInfoResponseBodyResult) SetTag(v string) *QueryOpenGroupBaseInfoResponseBodyResult {
	s.Tag = &v
	return s
}

func (s *QueryOpenGroupBaseInfoResponseBodyResult) SetTitle(v string) *QueryOpenGroupBaseInfoResponseBodyResult {
	s.Title = &v
	return s
}

type QueryOpenGroupBaseInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOpenGroupBaseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOpenGroupBaseInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOpenGroupBaseInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryOpenGroupBaseInfoResponse) SetHeaders(v map[string]*string) *QueryOpenGroupBaseInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryOpenGroupBaseInfoResponse) SetStatusCode(v int32) *QueryOpenGroupBaseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOpenGroupBaseInfoResponse) SetBody(v *QueryOpenGroupBaseInfoResponseBody) *QueryOpenGroupBaseInfoResponse {
	s.Body = v
	return s
}

type QueryPersonalMessageReadStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryPersonalMessageReadStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryPersonalMessageReadStatusHeaders) GoString() string {
	return s.String()
}

func (s *QueryPersonalMessageReadStatusHeaders) SetCommonHeaders(v map[string]*string) *QueryPersonalMessageReadStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryPersonalMessageReadStatusHeaders) SetXAcsDingtalkAccessToken(v string) *QueryPersonalMessageReadStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryPersonalMessageReadStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cidQGfKJCXMfVxZxxx3ZL0Qlw==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// msghnezLi8wb6pGqMsadhj9n0yw==
	OpenMessageId *string `json:"openMessageId,omitempty" xml:"openMessageId,omitempty"`
}

func (s QueryPersonalMessageReadStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPersonalMessageReadStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryPersonalMessageReadStatusRequest) SetOpenConversationId(v string) *QueryPersonalMessageReadStatusRequest {
	s.OpenConversationId = &v
	return s
}

func (s *QueryPersonalMessageReadStatusRequest) SetOpenMessageId(v string) *QueryPersonalMessageReadStatusRequest {
	s.OpenMessageId = &v
	return s
}

type QueryPersonalMessageReadStatusResponseBody struct {
	Result  *QueryPersonalMessageReadStatusResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryPersonalMessageReadStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPersonalMessageReadStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPersonalMessageReadStatusResponseBody) SetResult(v *QueryPersonalMessageReadStatusResponseBodyResult) *QueryPersonalMessageReadStatusResponseBody {
	s.Result = v
	return s
}

func (s *QueryPersonalMessageReadStatusResponseBody) SetSuccess(v bool) *QueryPersonalMessageReadStatusResponseBody {
	s.Success = &v
	return s
}

type QueryPersonalMessageReadStatusResponseBodyResult struct {
	MessageReadInfoList []*QueryPersonalMessageReadStatusResponseBodyResultMessageReadInfoList `json:"messageReadInfoList,omitempty" xml:"messageReadInfoList,omitempty" type:"Repeated"`
}

func (s QueryPersonalMessageReadStatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryPersonalMessageReadStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryPersonalMessageReadStatusResponseBodyResult) SetMessageReadInfoList(v []*QueryPersonalMessageReadStatusResponseBodyResultMessageReadInfoList) *QueryPersonalMessageReadStatusResponseBodyResult {
	s.MessageReadInfoList = v
	return s
}

type QueryPersonalMessageReadStatusResponseBodyResultMessageReadInfoList struct {
	ReadStatus *string `json:"readStatus,omitempty" xml:"readStatus,omitempty"`
	UserId     *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryPersonalMessageReadStatusResponseBodyResultMessageReadInfoList) String() string {
	return tea.Prettify(s)
}

func (s QueryPersonalMessageReadStatusResponseBodyResultMessageReadInfoList) GoString() string {
	return s.String()
}

func (s *QueryPersonalMessageReadStatusResponseBodyResultMessageReadInfoList) SetReadStatus(v string) *QueryPersonalMessageReadStatusResponseBodyResultMessageReadInfoList {
	s.ReadStatus = &v
	return s
}

func (s *QueryPersonalMessageReadStatusResponseBodyResultMessageReadInfoList) SetUserId(v string) *QueryPersonalMessageReadStatusResponseBodyResultMessageReadInfoList {
	s.UserId = &v
	return s
}

type QueryPersonalMessageReadStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPersonalMessageReadStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPersonalMessageReadStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPersonalMessageReadStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryPersonalMessageReadStatusResponse) SetHeaders(v map[string]*string) *QueryPersonalMessageReadStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryPersonalMessageReadStatusResponse) SetStatusCode(v int32) *QueryPersonalMessageReadStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPersonalMessageReadStatusResponse) SetBody(v *QueryPersonalMessageReadStatusResponseBody) *QueryPersonalMessageReadStatusResponse {
	s.Body = v
	return s
}

type QueryRecentConversationsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryRecentConversationsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryRecentConversationsHeaders) GoString() string {
	return s.String()
}

func (s *QueryRecentConversationsHeaders) SetCommonHeaders(v map[string]*string) *QueryRecentConversationsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryRecentConversationsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryRecentConversationsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryRecentConversationsRequest struct {
	OnlyHuman      *bool   `json:"onlyHuman,omitempty" xml:"onlyHuman,omitempty"`
	OnlyInnerGroup *bool   `json:"onlyInnerGroup,omitempty" xml:"onlyInnerGroup,omitempty"`
	UserId         *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryRecentConversationsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRecentConversationsRequest) GoString() string {
	return s.String()
}

func (s *QueryRecentConversationsRequest) SetOnlyHuman(v bool) *QueryRecentConversationsRequest {
	s.OnlyHuman = &v
	return s
}

func (s *QueryRecentConversationsRequest) SetOnlyInnerGroup(v bool) *QueryRecentConversationsRequest {
	s.OnlyInnerGroup = &v
	return s
}

func (s *QueryRecentConversationsRequest) SetUserId(v string) *QueryRecentConversationsRequest {
	s.UserId = &v
	return s
}

type QueryRecentConversationsResponseBody struct {
	Result  *QueryRecentConversationsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                       `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryRecentConversationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRecentConversationsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecentConversationsResponseBody) SetResult(v *QueryRecentConversationsResponseBodyResult) *QueryRecentConversationsResponseBody {
	s.Result = v
	return s
}

func (s *QueryRecentConversationsResponseBody) SetSuccess(v bool) *QueryRecentConversationsResponseBody {
	s.Success = &v
	return s
}

type QueryRecentConversationsResponseBodyResult struct {
	ConversationList []*QueryRecentConversationsResponseBodyResultConversationList `json:"conversationList,omitempty" xml:"conversationList,omitempty" type:"Repeated"`
}

func (s QueryRecentConversationsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryRecentConversationsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryRecentConversationsResponseBodyResult) SetConversationList(v []*QueryRecentConversationsResponseBodyResultConversationList) *QueryRecentConversationsResponseBodyResult {
	s.ConversationList = v
	return s
}

type QueryRecentConversationsResponseBodyResultConversationList struct {
	// This parameter is required.
	ConversationType *int32 `json:"conversationType,omitempty" xml:"conversationType,omitempty"`
	// This parameter is required.
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// This parameter is required.
	MemberCount *string `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryRecentConversationsResponseBodyResultConversationList) String() string {
	return tea.Prettify(s)
}

func (s QueryRecentConversationsResponseBodyResultConversationList) GoString() string {
	return s.String()
}

func (s *QueryRecentConversationsResponseBodyResultConversationList) SetConversationType(v int32) *QueryRecentConversationsResponseBodyResultConversationList {
	s.ConversationType = &v
	return s
}

func (s *QueryRecentConversationsResponseBodyResultConversationList) SetIcon(v string) *QueryRecentConversationsResponseBodyResultConversationList {
	s.Icon = &v
	return s
}

func (s *QueryRecentConversationsResponseBodyResultConversationList) SetMemberCount(v string) *QueryRecentConversationsResponseBodyResultConversationList {
	s.MemberCount = &v
	return s
}

func (s *QueryRecentConversationsResponseBodyResultConversationList) SetName(v string) *QueryRecentConversationsResponseBodyResultConversationList {
	s.Name = &v
	return s
}

func (s *QueryRecentConversationsResponseBodyResultConversationList) SetNickName(v string) *QueryRecentConversationsResponseBodyResultConversationList {
	s.NickName = &v
	return s
}

func (s *QueryRecentConversationsResponseBodyResultConversationList) SetOpenConversationId(v string) *QueryRecentConversationsResponseBodyResultConversationList {
	s.OpenConversationId = &v
	return s
}

func (s *QueryRecentConversationsResponseBodyResultConversationList) SetTitle(v string) *QueryRecentConversationsResponseBodyResultConversationList {
	s.Title = &v
	return s
}

func (s *QueryRecentConversationsResponseBodyResultConversationList) SetUserId(v string) *QueryRecentConversationsResponseBodyResultConversationList {
	s.UserId = &v
	return s
}

type QueryRecentConversationsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRecentConversationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRecentConversationsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRecentConversationsResponse) GoString() string {
	return s.String()
}

func (s *QueryRecentConversationsResponse) SetHeaders(v map[string]*string) *QueryRecentConversationsResponse {
	s.Headers = v
	return s
}

func (s *QueryRecentConversationsResponse) SetStatusCode(v int32) *QueryRecentConversationsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRecentConversationsResponse) SetBody(v *QueryRecentConversationsResponseBody) *QueryRecentConversationsResponse {
	s.Body = v
	return s
}

type QuerySceneGroupTemplateRobotHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QuerySceneGroupTemplateRobotHeaders) String() string {
	return tea.Prettify(s)
}

func (s QuerySceneGroupTemplateRobotHeaders) GoString() string {
	return s.String()
}

func (s *QuerySceneGroupTemplateRobotHeaders) SetCommonHeaders(v map[string]*string) *QuerySceneGroupTemplateRobotHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QuerySceneGroupTemplateRobotHeaders) SetXAcsDingtalkAccessToken(v string) *QuerySceneGroupTemplateRobotHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QuerySceneGroupTemplateRobotRequest struct {
	// example:
	//
	// cidCtneF+XyQjcyF2ROdgSeIg==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// ding5nbbeXXXXXXX
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s QuerySceneGroupTemplateRobotRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySceneGroupTemplateRobotRequest) GoString() string {
	return s.String()
}

func (s *QuerySceneGroupTemplateRobotRequest) SetOpenConversationId(v string) *QuerySceneGroupTemplateRobotRequest {
	s.OpenConversationId = &v
	return s
}

func (s *QuerySceneGroupTemplateRobotRequest) SetRobotCode(v string) *QuerySceneGroupTemplateRobotRequest {
	s.RobotCode = &v
	return s
}

type QuerySceneGroupTemplateRobotResponseBody struct {
	Result  *QuerySceneGroupTemplateRobotResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QuerySceneGroupTemplateRobotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySceneGroupTemplateRobotResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySceneGroupTemplateRobotResponseBody) SetResult(v *QuerySceneGroupTemplateRobotResponseBodyResult) *QuerySceneGroupTemplateRobotResponseBody {
	s.Result = v
	return s
}

func (s *QuerySceneGroupTemplateRobotResponseBody) SetSuccess(v bool) *QuerySceneGroupTemplateRobotResponseBody {
	s.Success = &v
	return s
}

type QuerySceneGroupTemplateRobotResponseBodyResult struct {
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId  *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QuerySceneGroupTemplateRobotResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QuerySceneGroupTemplateRobotResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QuerySceneGroupTemplateRobotResponseBodyResult) SetUnionId(v string) *QuerySceneGroupTemplateRobotResponseBodyResult {
	s.UnionId = &v
	return s
}

func (s *QuerySceneGroupTemplateRobotResponseBodyResult) SetUserId(v string) *QuerySceneGroupTemplateRobotResponseBodyResult {
	s.UserId = &v
	return s
}

type QuerySceneGroupTemplateRobotResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySceneGroupTemplateRobotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySceneGroupTemplateRobotResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySceneGroupTemplateRobotResponse) GoString() string {
	return s.String()
}

func (s *QuerySceneGroupTemplateRobotResponse) SetHeaders(v map[string]*string) *QuerySceneGroupTemplateRobotResponse {
	s.Headers = v
	return s
}

func (s *QuerySceneGroupTemplateRobotResponse) SetStatusCode(v int32) *QuerySceneGroupTemplateRobotResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySceneGroupTemplateRobotResponse) SetBody(v *QuerySceneGroupTemplateRobotResponseBody) *QuerySceneGroupTemplateRobotResponse {
	s.Body = v
	return s
}

type QuerySingleGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QuerySingleGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleGroupHeaders) GoString() string {
	return s.String()
}

func (s *QuerySingleGroupHeaders) SetCommonHeaders(v map[string]*string) *QuerySingleGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QuerySingleGroupHeaders) SetXAcsDingtalkAccessToken(v string) *QuerySingleGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QuerySingleGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1745****8777
	GroupMembers []*QuerySingleGroupRequestGroupMembers `json:"groupMembers,omitempty" xml:"groupMembers,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 14da****2760
	GroupTemplateId *string `json:"groupTemplateId,omitempty" xml:"groupTemplateId,omitempty"`
}

func (s QuerySingleGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleGroupRequest) GoString() string {
	return s.String()
}

func (s *QuerySingleGroupRequest) SetGroupMembers(v []*QuerySingleGroupRequestGroupMembers) *QuerySingleGroupRequest {
	s.GroupMembers = v
	return s
}

func (s *QuerySingleGroupRequest) SetGroupTemplateId(v string) *QuerySingleGroupRequest {
	s.GroupTemplateId = &v
	return s
}

type QuerySingleGroupRequestGroupMembers struct {
	// This parameter is required.
	//
	// example:
	//
	// 1107****2120
	AppUserId *string `json:"appUserId,omitempty" xml:"appUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1745****8778
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QuerySingleGroupRequestGroupMembers) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleGroupRequestGroupMembers) GoString() string {
	return s.String()
}

func (s *QuerySingleGroupRequestGroupMembers) SetAppUserId(v string) *QuerySingleGroupRequestGroupMembers {
	s.AppUserId = &v
	return s
}

func (s *QuerySingleGroupRequestGroupMembers) SetUserId(v string) *QuerySingleGroupRequestGroupMembers {
	s.UserId = &v
	return s
}

type QuerySingleGroupResponseBody struct {
	// This parameter is required.
	OpenConversations []*QuerySingleGroupResponseBodyOpenConversations `json:"openConversations,omitempty" xml:"openConversations,omitempty" type:"Repeated"`
}

func (s QuerySingleGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySingleGroupResponseBody) SetOpenConversations(v []*QuerySingleGroupResponseBodyOpenConversations) *QuerySingleGroupResponseBody {
	s.OpenConversations = v
	return s
}

type QuerySingleGroupResponseBodyOpenConversations struct {
	// This parameter is required.
	//
	// example:
	//
	// 1107****2120
	AppUserId *string `json:"appUserId,omitempty" xml:"appUserId,omitempty"`
	// example:
	//
	// 14da****2760
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1745****8778
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QuerySingleGroupResponseBodyOpenConversations) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleGroupResponseBodyOpenConversations) GoString() string {
	return s.String()
}

func (s *QuerySingleGroupResponseBodyOpenConversations) SetAppUserId(v string) *QuerySingleGroupResponseBodyOpenConversations {
	s.AppUserId = &v
	return s
}

func (s *QuerySingleGroupResponseBodyOpenConversations) SetOpenConversationId(v string) *QuerySingleGroupResponseBodyOpenConversations {
	s.OpenConversationId = &v
	return s
}

func (s *QuerySingleGroupResponseBodyOpenConversations) SetUserId(v string) *QuerySingleGroupResponseBodyOpenConversations {
	s.UserId = &v
	return s
}

type QuerySingleGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySingleGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySingleGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySingleGroupResponse) GoString() string {
	return s.String()
}

func (s *QuerySingleGroupResponse) SetHeaders(v map[string]*string) *QuerySingleGroupResponse {
	s.Headers = v
	return s
}

func (s *QuerySingleGroupResponse) SetStatusCode(v int32) *QuerySingleGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySingleGroupResponse) SetBody(v *QuerySingleGroupResponseBody) *QuerySingleGroupResponse {
	s.Body = v
	return s
}

type QueryUnReadMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUnReadMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUnReadMessageHeaders) GoString() string {
	return s.String()
}

func (s *QueryUnReadMessageHeaders) SetCommonHeaders(v map[string]*string) *QueryUnReadMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUnReadMessageHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUnReadMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUnReadMessageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1107****2120
	AppUserId *string `json:"appUserId,omitempty" xml:"appUserId,omitempty"`
	// example:
	//
	// 1745****8777
	OpenConversationIds []*string `json:"openConversationIds,omitempty" xml:"openConversationIds,omitempty" type:"Repeated"`
}

func (s QueryUnReadMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUnReadMessageRequest) GoString() string {
	return s.String()
}

func (s *QueryUnReadMessageRequest) SetAppUserId(v string) *QueryUnReadMessageRequest {
	s.AppUserId = &v
	return s
}

func (s *QueryUnReadMessageRequest) SetOpenConversationIds(v []*string) *QueryUnReadMessageRequest {
	s.OpenConversationIds = v
	return s
}

type QueryUnReadMessageResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	UnReadCount *int64                                       `json:"unReadCount,omitempty" xml:"unReadCount,omitempty"`
	UnReadItems []*QueryUnReadMessageResponseBodyUnReadItems `json:"unReadItems,omitempty" xml:"unReadItems,omitempty" type:"Repeated"`
}

func (s QueryUnReadMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUnReadMessageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUnReadMessageResponseBody) SetUnReadCount(v int64) *QueryUnReadMessageResponseBody {
	s.UnReadCount = &v
	return s
}

func (s *QueryUnReadMessageResponseBody) SetUnReadItems(v []*QueryUnReadMessageResponseBodyUnReadItems) *QueryUnReadMessageResponseBody {
	s.UnReadItems = v
	return s
}

type QueryUnReadMessageResponseBodyUnReadItems struct {
	// example:
	//
	// 14da****2760
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// 10
	UnReadCount *int64 `json:"unReadCount,omitempty" xml:"unReadCount,omitempty"`
}

func (s QueryUnReadMessageResponseBodyUnReadItems) String() string {
	return tea.Prettify(s)
}

func (s QueryUnReadMessageResponseBodyUnReadItems) GoString() string {
	return s.String()
}

func (s *QueryUnReadMessageResponseBodyUnReadItems) SetOpenConversationId(v string) *QueryUnReadMessageResponseBodyUnReadItems {
	s.OpenConversationId = &v
	return s
}

func (s *QueryUnReadMessageResponseBodyUnReadItems) SetUnReadCount(v int64) *QueryUnReadMessageResponseBodyUnReadItems {
	s.UnReadCount = &v
	return s
}

type QueryUnReadMessageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUnReadMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUnReadMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUnReadMessageResponse) GoString() string {
	return s.String()
}

func (s *QueryUnReadMessageResponse) SetHeaders(v map[string]*string) *QueryUnReadMessageResponse {
	s.Headers = v
	return s
}

func (s *QueryUnReadMessageResponse) SetStatusCode(v int32) *QueryUnReadMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUnReadMessageResponse) SetBody(v *QueryUnReadMessageResponseBody) *QueryUnReadMessageResponse {
	s.Body = v
	return s
}

type QueryUnfurlingRegisterCreatorHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUnfurlingRegisterCreatorHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfurlingRegisterCreatorHeaders) GoString() string {
	return s.String()
}

func (s *QueryUnfurlingRegisterCreatorHeaders) SetCommonHeaders(v map[string]*string) *QueryUnfurlingRegisterCreatorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUnfurlingRegisterCreatorHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUnfurlingRegisterCreatorHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUnfurlingRegisterCreatorRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.dingtalk.com
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /a
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s QueryUnfurlingRegisterCreatorRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfurlingRegisterCreatorRequest) GoString() string {
	return s.String()
}

func (s *QueryUnfurlingRegisterCreatorRequest) SetDomain(v string) *QueryUnfurlingRegisterCreatorRequest {
	s.Domain = &v
	return s
}

func (s *QueryUnfurlingRegisterCreatorRequest) SetPath(v string) *QueryUnfurlingRegisterCreatorRequest {
	s.Path = &v
	return s
}

type QueryUnfurlingRegisterCreatorResponseBody struct {
	Data    *QueryUnfurlingRegisterCreatorResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Success *bool                                          `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryUnfurlingRegisterCreatorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfurlingRegisterCreatorResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUnfurlingRegisterCreatorResponseBody) SetData(v *QueryUnfurlingRegisterCreatorResponseBodyData) *QueryUnfurlingRegisterCreatorResponseBody {
	s.Data = v
	return s
}

func (s *QueryUnfurlingRegisterCreatorResponseBody) SetSuccess(v bool) *QueryUnfurlingRegisterCreatorResponseBody {
	s.Success = &v
	return s
}

type QueryUnfurlingRegisterCreatorResponseBodyData struct {
	AppId         *string `json:"appId,omitempty" xml:"appId,omitempty"`
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	Id            *int64  `json:"id,omitempty" xml:"id,omitempty"`
}

func (s QueryUnfurlingRegisterCreatorResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfurlingRegisterCreatorResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryUnfurlingRegisterCreatorResponseBodyData) SetAppId(v string) *QueryUnfurlingRegisterCreatorResponseBodyData {
	s.AppId = &v
	return s
}

func (s *QueryUnfurlingRegisterCreatorResponseBodyData) SetCreatorUserId(v string) *QueryUnfurlingRegisterCreatorResponseBodyData {
	s.CreatorUserId = &v
	return s
}

func (s *QueryUnfurlingRegisterCreatorResponseBodyData) SetId(v int64) *QueryUnfurlingRegisterCreatorResponseBodyData {
	s.Id = &v
	return s
}

type QueryUnfurlingRegisterCreatorResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUnfurlingRegisterCreatorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUnfurlingRegisterCreatorResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfurlingRegisterCreatorResponse) GoString() string {
	return s.String()
}

func (s *QueryUnfurlingRegisterCreatorResponse) SetHeaders(v map[string]*string) *QueryUnfurlingRegisterCreatorResponse {
	s.Headers = v
	return s
}

func (s *QueryUnfurlingRegisterCreatorResponse) SetStatusCode(v int32) *QueryUnfurlingRegisterCreatorResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUnfurlingRegisterCreatorResponse) SetBody(v *QueryUnfurlingRegisterCreatorResponseBody) *QueryUnfurlingRegisterCreatorResponse {
	s.Body = v
	return s
}

type QueryUnfurlingRegisterInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUnfurlingRegisterInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfurlingRegisterInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryUnfurlingRegisterInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryUnfurlingRegisterInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUnfurlingRegisterInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUnfurlingRegisterInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUnfurlingRegisterInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3102xxxxxxx
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 0
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryUnfurlingRegisterInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfurlingRegisterInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryUnfurlingRegisterInfoRequest) SetAppId(v string) *QueryUnfurlingRegisterInfoRequest {
	s.AppId = &v
	return s
}

func (s *QueryUnfurlingRegisterInfoRequest) SetMaxResults(v int32) *QueryUnfurlingRegisterInfoRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryUnfurlingRegisterInfoRequest) SetNextToken(v int64) *QueryUnfurlingRegisterInfoRequest {
	s.NextToken = &v
	return s
}

type QueryUnfurlingRegisterInfoResponseBody struct {
	HasMore   *bool                                         `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List      []*QueryUnfurlingRegisterInfoResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextToken *int64                                        `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Success   *bool                                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryUnfurlingRegisterInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfurlingRegisterInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUnfurlingRegisterInfoResponseBody) SetHasMore(v bool) *QueryUnfurlingRegisterInfoResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryUnfurlingRegisterInfoResponseBody) SetList(v []*QueryUnfurlingRegisterInfoResponseBodyList) *QueryUnfurlingRegisterInfoResponseBody {
	s.List = v
	return s
}

func (s *QueryUnfurlingRegisterInfoResponseBody) SetNextToken(v int64) *QueryUnfurlingRegisterInfoResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryUnfurlingRegisterInfoResponseBody) SetSuccess(v bool) *QueryUnfurlingRegisterInfoResponseBody {
	s.Success = &v
	return s
}

type QueryUnfurlingRegisterInfoResponseBodyList struct {
	ApiSecret       *string   `json:"apiSecret,omitempty" xml:"apiSecret,omitempty"`
	AppId           *string   `json:"appId,omitempty" xml:"appId,omitempty"`
	AppName         *string   `json:"appName,omitempty" xml:"appName,omitempty"`
	CallbackType    *int32    `json:"callbackType,omitempty" xml:"callbackType,omitempty"`
	CallbackUrl     *string   `json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`
	CardTemplateId  *string   `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	CreatorUserId   *string   `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	Domain          *string   `json:"domain,omitempty" xml:"domain,omitempty"`
	GrayGroupIdList []*string `json:"grayGroupIdList,omitempty" xml:"grayGroupIdList,omitempty" type:"Repeated"`
	GrayUserIdList  []*string `json:"grayUserIdList,omitempty" xml:"grayUserIdList,omitempty" type:"Repeated"`
	HsfMethodName   *string   `json:"hsfMethodName,omitempty" xml:"hsfMethodName,omitempty"`
	HsfServiceName  *string   `json:"hsfServiceName,omitempty" xml:"hsfServiceName,omitempty"`
	HsfVersion      *string   `json:"hsfVersion,omitempty" xml:"hsfVersion,omitempty"`
	Id              *int64    `json:"id,omitempty" xml:"id,omitempty"`
	Path            *string   `json:"path,omitempty" xml:"path,omitempty"`
	RuleDesc        *string   `json:"ruleDesc,omitempty" xml:"ruleDesc,omitempty"`
	RuleMatchType   *int32    `json:"ruleMatchType,omitempty" xml:"ruleMatchType,omitempty"`
	Status          *int32    `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryUnfurlingRegisterInfoResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfurlingRegisterInfoResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryUnfurlingRegisterInfoResponseBodyList) SetApiSecret(v string) *QueryUnfurlingRegisterInfoResponseBodyList {
	s.ApiSecret = &v
	return s
}

func (s *QueryUnfurlingRegisterInfoResponseBodyList) SetAppId(v string) *QueryUnfurlingRegisterInfoResponseBodyList {
	s.AppId = &v
	return s
}

func (s *QueryUnfurlingRegisterInfoResponseBodyList) SetAppName(v string) *QueryUnfurlingRegisterInfoResponseBodyList {
	s.AppName = &v
	return s
}

func (s *QueryUnfurlingRegisterInfoResponseBodyList) SetCallbackType(v int32) *QueryUnfurlingRegisterInfoResponseBodyList {
	s.CallbackType = &v
	return s
}

func (s *QueryUnfurlingRegisterInfoResponseBodyList) SetCallbackUrl(v string) *QueryUnfurlingRegisterInfoResponseBodyList {
	s.CallbackUrl = &v
	return s
}

func (s *QueryUnfurlingRegisterInfoResponseBodyList) SetCardTemplateId(v string) *QueryUnfurlingRegisterInfoResponseBodyList {
	s.CardTemplateId = &v
	return s
}

func (s *QueryUnfurlingRegisterInfoResponseBodyList) SetCreatorUserId(v string) *QueryUnfurlingRegisterInfoResponseBodyList {
	s.CreatorUserId = &v
	return s
}

func (s *QueryUnfurlingRegisterInfoResponseBodyList) SetDomain(v string) *QueryUnfurlingRegisterInfoResponseBodyList {
	s.Domain = &v
	return s
}

func (s *QueryUnfurlingRegisterInfoResponseBodyList) SetGrayGroupIdList(v []*string) *QueryUnfurlingRegisterInfoResponseBodyList {
	s.GrayGroupIdList = v
	return s
}

func (s *QueryUnfurlingRegisterInfoResponseBodyList) SetGrayUserIdList(v []*string) *QueryUnfurlingRegisterInfoResponseBodyList {
	s.GrayUserIdList = v
	return s
}

func (s *QueryUnfurlingRegisterInfoResponseBodyList) SetHsfMethodName(v string) *QueryUnfurlingRegisterInfoResponseBodyList {
	s.HsfMethodName = &v
	return s
}

func (s *QueryUnfurlingRegisterInfoResponseBodyList) SetHsfServiceName(v string) *QueryUnfurlingRegisterInfoResponseBodyList {
	s.HsfServiceName = &v
	return s
}

func (s *QueryUnfurlingRegisterInfoResponseBodyList) SetHsfVersion(v string) *QueryUnfurlingRegisterInfoResponseBodyList {
	s.HsfVersion = &v
	return s
}

func (s *QueryUnfurlingRegisterInfoResponseBodyList) SetId(v int64) *QueryUnfurlingRegisterInfoResponseBodyList {
	s.Id = &v
	return s
}

func (s *QueryUnfurlingRegisterInfoResponseBodyList) SetPath(v string) *QueryUnfurlingRegisterInfoResponseBodyList {
	s.Path = &v
	return s
}

func (s *QueryUnfurlingRegisterInfoResponseBodyList) SetRuleDesc(v string) *QueryUnfurlingRegisterInfoResponseBodyList {
	s.RuleDesc = &v
	return s
}

func (s *QueryUnfurlingRegisterInfoResponseBodyList) SetRuleMatchType(v int32) *QueryUnfurlingRegisterInfoResponseBodyList {
	s.RuleMatchType = &v
	return s
}

func (s *QueryUnfurlingRegisterInfoResponseBodyList) SetStatus(v int32) *QueryUnfurlingRegisterInfoResponseBodyList {
	s.Status = &v
	return s
}

type QueryUnfurlingRegisterInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUnfurlingRegisterInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUnfurlingRegisterInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfurlingRegisterInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryUnfurlingRegisterInfoResponse) SetHeaders(v map[string]*string) *QueryUnfurlingRegisterInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryUnfurlingRegisterInfoResponse) SetStatusCode(v int32) *QueryUnfurlingRegisterInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUnfurlingRegisterInfoResponse) SetBody(v *QueryUnfurlingRegisterInfoResponseBody) *QueryUnfurlingRegisterInfoResponse {
	s.Body = v
	return s
}

type QueryUserViewGroupLastMessageTimeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserViewGroupLastMessageTimeHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserViewGroupLastMessageTimeHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserViewGroupLastMessageTimeHeaders) SetCommonHeaders(v map[string]*string) *QueryUserViewGroupLastMessageTimeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserViewGroupLastMessageTimeHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserViewGroupLastMessageTimeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserViewGroupLastMessageTimeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cidc4iLyQBuHFQRvxxxnz204Q==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s QueryUserViewGroupLastMessageTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserViewGroupLastMessageTimeRequest) GoString() string {
	return s.String()
}

func (s *QueryUserViewGroupLastMessageTimeRequest) SetOpenConversationId(v string) *QueryUserViewGroupLastMessageTimeRequest {
	s.OpenConversationId = &v
	return s
}

type QueryUserViewGroupLastMessageTimeResponseBody struct {
	Result  *QueryUserViewGroupLastMessageTimeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                                `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryUserViewGroupLastMessageTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserViewGroupLastMessageTimeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserViewGroupLastMessageTimeResponseBody) SetResult(v *QueryUserViewGroupLastMessageTimeResponseBodyResult) *QueryUserViewGroupLastMessageTimeResponseBody {
	s.Result = v
	return s
}

func (s *QueryUserViewGroupLastMessageTimeResponseBody) SetSuccess(v bool) *QueryUserViewGroupLastMessageTimeResponseBody {
	s.Success = &v
	return s
}

type QueryUserViewGroupLastMessageTimeResponseBodyResult struct {
	Time *int64 `json:"time,omitempty" xml:"time,omitempty"`
}

func (s QueryUserViewGroupLastMessageTimeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryUserViewGroupLastMessageTimeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryUserViewGroupLastMessageTimeResponseBodyResult) SetTime(v int64) *QueryUserViewGroupLastMessageTimeResponseBodyResult {
	s.Time = &v
	return s
}

type QueryUserViewGroupLastMessageTimeResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserViewGroupLastMessageTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserViewGroupLastMessageTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserViewGroupLastMessageTimeResponse) GoString() string {
	return s.String()
}

func (s *QueryUserViewGroupLastMessageTimeResponse) SetHeaders(v map[string]*string) *QueryUserViewGroupLastMessageTimeResponse {
	s.Headers = v
	return s
}

func (s *QueryUserViewGroupLastMessageTimeResponse) SetStatusCode(v int32) *QueryUserViewGroupLastMessageTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserViewGroupLastMessageTimeResponse) SetBody(v *QueryUserViewGroupLastMessageTimeResponseBody) *QueryUserViewGroupLastMessageTimeResponse {
	s.Body = v
	return s
}

type ReadPersonalMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ReadPersonalMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s ReadPersonalMessageHeaders) GoString() string {
	return s.String()
}

func (s *ReadPersonalMessageHeaders) SetCommonHeaders(v map[string]*string) *ReadPersonalMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ReadPersonalMessageHeaders) SetXAcsDingtalkAccessToken(v string) *ReadPersonalMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ReadPersonalMessageRequest struct {
	// This parameter is required.
	DingOpenConversationMessageIdArray []*ReadPersonalMessageRequestDingOpenConversationMessageIdArray `json:"dingOpenConversationMessageIdArray,omitempty" xml:"dingOpenConversationMessageIdArray,omitempty" type:"Repeated"`
}

func (s ReadPersonalMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s ReadPersonalMessageRequest) GoString() string {
	return s.String()
}

func (s *ReadPersonalMessageRequest) SetDingOpenConversationMessageIdArray(v []*ReadPersonalMessageRequestDingOpenConversationMessageIdArray) *ReadPersonalMessageRequest {
	s.DingOpenConversationMessageIdArray = v
	return s
}

type ReadPersonalMessageRequestDingOpenConversationMessageIdArray struct {
	// This parameter is required.
	//
	// example:
	//
	// cidQGfKJCXMfVxZxxx3ZL0Qlw
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// msghnezLi8wb6pGqMsadhj9n0yw
	OpenMessageId *string `json:"openMessageId,omitempty" xml:"openMessageId,omitempty"`
}

func (s ReadPersonalMessageRequestDingOpenConversationMessageIdArray) String() string {
	return tea.Prettify(s)
}

func (s ReadPersonalMessageRequestDingOpenConversationMessageIdArray) GoString() string {
	return s.String()
}

func (s *ReadPersonalMessageRequestDingOpenConversationMessageIdArray) SetOpenConversationId(v string) *ReadPersonalMessageRequestDingOpenConversationMessageIdArray {
	s.OpenConversationId = &v
	return s
}

func (s *ReadPersonalMessageRequestDingOpenConversationMessageIdArray) SetOpenMessageId(v string) *ReadPersonalMessageRequestDingOpenConversationMessageIdArray {
	s.OpenMessageId = &v
	return s
}

type ReadPersonalMessageResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ReadPersonalMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReadPersonalMessageResponseBody) GoString() string {
	return s.String()
}

func (s *ReadPersonalMessageResponseBody) SetSuccess(v bool) *ReadPersonalMessageResponseBody {
	s.Success = &v
	return s
}

type ReadPersonalMessageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadPersonalMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadPersonalMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s ReadPersonalMessageResponse) GoString() string {
	return s.String()
}

func (s *ReadPersonalMessageResponse) SetHeaders(v map[string]*string) *ReadPersonalMessageResponse {
	s.Headers = v
	return s
}

func (s *ReadPersonalMessageResponse) SetStatusCode(v int32) *ReadPersonalMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadPersonalMessageResponse) SetBody(v *ReadPersonalMessageResponseBody) *ReadPersonalMessageResponse {
	s.Body = v
	return s
}

type RecallPersonalMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RecallPersonalMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s RecallPersonalMessageHeaders) GoString() string {
	return s.String()
}

func (s *RecallPersonalMessageHeaders) SetCommonHeaders(v map[string]*string) *RecallPersonalMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RecallPersonalMessageHeaders) SetXAcsDingtalkAccessToken(v string) *RecallPersonalMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RecallPersonalMessageRequest struct {
	// example:
	//
	// cidxxxx3451=
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// msgxxx112
	OpenMessageId *string `json:"openMessageId,omitempty" xml:"openMessageId,omitempty"`
}

func (s RecallPersonalMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s RecallPersonalMessageRequest) GoString() string {
	return s.String()
}

func (s *RecallPersonalMessageRequest) SetOpenConversationId(v string) *RecallPersonalMessageRequest {
	s.OpenConversationId = &v
	return s
}

func (s *RecallPersonalMessageRequest) SetOpenMessageId(v string) *RecallPersonalMessageRequest {
	s.OpenMessageId = &v
	return s
}

type RecallPersonalMessageResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RecallPersonalMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecallPersonalMessageResponseBody) GoString() string {
	return s.String()
}

func (s *RecallPersonalMessageResponseBody) SetSuccess(v bool) *RecallPersonalMessageResponseBody {
	s.Success = &v
	return s
}

type RecallPersonalMessageResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecallPersonalMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecallPersonalMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s RecallPersonalMessageResponse) GoString() string {
	return s.String()
}

func (s *RecallPersonalMessageResponse) SetHeaders(v map[string]*string) *RecallPersonalMessageResponse {
	s.Headers = v
	return s
}

func (s *RecallPersonalMessageResponse) SetStatusCode(v int32) *RecallPersonalMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *RecallPersonalMessageResponse) SetBody(v *RecallPersonalMessageResponseBody) *RecallPersonalMessageResponse {
	s.Body = v
	return s
}

type ReleaseUnfurlingRegisterHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ReleaseUnfurlingRegisterHeaders) String() string {
	return tea.Prettify(s)
}

func (s ReleaseUnfurlingRegisterHeaders) GoString() string {
	return s.String()
}

func (s *ReleaseUnfurlingRegisterHeaders) SetCommonHeaders(v map[string]*string) *ReleaseUnfurlingRegisterHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ReleaseUnfurlingRegisterHeaders) SetXAcsDingtalkAccessToken(v string) *ReleaseUnfurlingRegisterHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ReleaseUnfurlingRegisterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3102xxxxxxx
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 37xxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ReleaseUnfurlingRegisterRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseUnfurlingRegisterRequest) GoString() string {
	return s.String()
}

func (s *ReleaseUnfurlingRegisterRequest) SetAppId(v string) *ReleaseUnfurlingRegisterRequest {
	s.AppId = &v
	return s
}

func (s *ReleaseUnfurlingRegisterRequest) SetId(v int64) *ReleaseUnfurlingRegisterRequest {
	s.Id = &v
	return s
}

func (s *ReleaseUnfurlingRegisterRequest) SetUserId(v string) *ReleaseUnfurlingRegisterRequest {
	s.UserId = &v
	return s
}

type ReleaseUnfurlingRegisterResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ReleaseUnfurlingRegisterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseUnfurlingRegisterResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseUnfurlingRegisterResponseBody) SetSuccess(v bool) *ReleaseUnfurlingRegisterResponseBody {
	s.Success = &v
	return s
}

type ReleaseUnfurlingRegisterResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseUnfurlingRegisterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseUnfurlingRegisterResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseUnfurlingRegisterResponse) GoString() string {
	return s.String()
}

func (s *ReleaseUnfurlingRegisterResponse) SetHeaders(v map[string]*string) *ReleaseUnfurlingRegisterResponse {
	s.Headers = v
	return s
}

func (s *ReleaseUnfurlingRegisterResponse) SetStatusCode(v int32) *ReleaseUnfurlingRegisterResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseUnfurlingRegisterResponse) SetBody(v *ReleaseUnfurlingRegisterResponseBody) *ReleaseUnfurlingRegisterResponse {
	s.Body = v
	return s
}

type RemoveRobotFromConversationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RemoveRobotFromConversationHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveRobotFromConversationHeaders) GoString() string {
	return s.String()
}

func (s *RemoveRobotFromConversationHeaders) SetCommonHeaders(v map[string]*string) *RemoveRobotFromConversationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveRobotFromConversationHeaders) SetXAcsDingtalkAccessToken(v string) *RemoveRobotFromConversationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RemoveRobotFromConversationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	ChatBotUserId *string `json:"chatBotUserId,omitempty" xml:"chatBotUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cid123cd
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s RemoveRobotFromConversationRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveRobotFromConversationRequest) GoString() string {
	return s.String()
}

func (s *RemoveRobotFromConversationRequest) SetChatBotUserId(v string) *RemoveRobotFromConversationRequest {
	s.ChatBotUserId = &v
	return s
}

func (s *RemoveRobotFromConversationRequest) SetOpenConversationId(v string) *RemoveRobotFromConversationRequest {
	s.OpenConversationId = &v
	return s
}

type RemoveRobotFromConversationResponseBody struct {
	ChatBotUserId *string `json:"chatBotUserId,omitempty" xml:"chatBotUserId,omitempty"`
}

func (s RemoveRobotFromConversationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveRobotFromConversationResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveRobotFromConversationResponseBody) SetChatBotUserId(v string) *RemoveRobotFromConversationResponseBody {
	s.ChatBotUserId = &v
	return s
}

type RemoveRobotFromConversationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveRobotFromConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveRobotFromConversationResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveRobotFromConversationResponse) GoString() string {
	return s.String()
}

func (s *RemoveRobotFromConversationResponse) SetHeaders(v map[string]*string) *RemoveRobotFromConversationResponse {
	s.Headers = v
	return s
}

func (s *RemoveRobotFromConversationResponse) SetStatusCode(v int32) *RemoveRobotFromConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveRobotFromConversationResponse) SetBody(v *RemoveRobotFromConversationResponseBody) *RemoveRobotFromConversationResponse {
	s.Body = v
	return s
}

type SearchInnerGroupsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchInnerGroupsHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchInnerGroupsHeaders) GoString() string {
	return s.String()
}

func (s *SearchInnerGroupsHeaders) SetCommonHeaders(v map[string]*string) *SearchInnerGroupsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchInnerGroupsHeaders) SetXAcsDingtalkAccessToken(v string) *SearchInnerGroupsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchInnerGroupsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试关键词
	SearchKey *string `json:"searchKey,omitempty" xml:"searchKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 015*****
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchInnerGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchInnerGroupsRequest) GoString() string {
	return s.String()
}

func (s *SearchInnerGroupsRequest) SetMaxResults(v int32) *SearchInnerGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchInnerGroupsRequest) SetSearchKey(v string) *SearchInnerGroupsRequest {
	s.SearchKey = &v
	return s
}

func (s *SearchInnerGroupsRequest) SetUserId(v string) *SearchInnerGroupsRequest {
	s.UserId = &v
	return s
}

type SearchInnerGroupsResponseBody struct {
	GroupInfos []*SearchInnerGroupsResponseBodyGroupInfos `json:"groupInfos,omitempty" xml:"groupInfos,omitempty" type:"Repeated"`
}

func (s SearchInnerGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchInnerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchInnerGroupsResponseBody) SetGroupInfos(v []*SearchInnerGroupsResponseBodyGroupInfos) *SearchInnerGroupsResponseBody {
	s.GroupInfos = v
	return s
}

type SearchInnerGroupsResponseBodyGroupInfos struct {
	// example:
	//
	// @lAD*****
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// 10
	MemberAmount *string `json:"memberAmount,omitempty" xml:"memberAmount,omitempty"`
	// example:
	//
	// cid13*****==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// 测试群名称
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SearchInnerGroupsResponseBodyGroupInfos) String() string {
	return tea.Prettify(s)
}

func (s SearchInnerGroupsResponseBodyGroupInfos) GoString() string {
	return s.String()
}

func (s *SearchInnerGroupsResponseBodyGroupInfos) SetIcon(v string) *SearchInnerGroupsResponseBodyGroupInfos {
	s.Icon = &v
	return s
}

func (s *SearchInnerGroupsResponseBodyGroupInfos) SetMemberAmount(v string) *SearchInnerGroupsResponseBodyGroupInfos {
	s.MemberAmount = &v
	return s
}

func (s *SearchInnerGroupsResponseBodyGroupInfos) SetOpenConversationId(v string) *SearchInnerGroupsResponseBodyGroupInfos {
	s.OpenConversationId = &v
	return s
}

func (s *SearchInnerGroupsResponseBodyGroupInfos) SetTitle(v string) *SearchInnerGroupsResponseBodyGroupInfos {
	s.Title = &v
	return s
}

type SearchInnerGroupsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchInnerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchInnerGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchInnerGroupsResponse) GoString() string {
	return s.String()
}

func (s *SearchInnerGroupsResponse) SetHeaders(v map[string]*string) *SearchInnerGroupsResponse {
	s.Headers = v
	return s
}

func (s *SearchInnerGroupsResponse) SetStatusCode(v int32) *SearchInnerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchInnerGroupsResponse) SetBody(v *SearchInnerGroupsResponseBody) *SearchInnerGroupsResponse {
	s.Body = v
	return s
}

type SendInteractiveCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendInteractiveCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendInteractiveCardHeaders) GoString() string {
	return s.String()
}

func (s *SendInteractiveCardHeaders) SetCommonHeaders(v map[string]*string) *SendInteractiveCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendInteractiveCardHeaders) SetXAcsDingtalkAccessToken(v string) *SendInteractiveCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendInteractiveCardRequest struct {
	AtOpenIds        map[string]*string `json:"atOpenIds,omitempty" xml:"atOpenIds,omitempty"`
	CallbackRouteKey *string            `json:"callbackRouteKey,omitempty" xml:"callbackRouteKey,omitempty"`
	// This parameter is required.
	CardData    *SendInteractiveCardRequestCardData    `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	CardOptions *SendInteractiveCardRequestCardOptions `json:"cardOptions,omitempty" xml:"cardOptions,omitempty" type:"Struct"`
	// This parameter is required.
	CardTemplateId *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	ChatBotId      *string `json:"chatBotId,omitempty" xml:"chatBotId,omitempty"`
	// This parameter is required.
	ConversationType   *int32  `json:"conversationType,omitempty" xml:"conversationType,omitempty"`
	DigitalWorkerCode  *string `json:"digitalWorkerCode,omitempty" xml:"digitalWorkerCode,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	OutTrackId         *string                      `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	PrivateData        map[string]*PrivateDataValue `json:"privateData,omitempty" xml:"privateData,omitempty"`
	PullStrategy       *bool                        `json:"pullStrategy,omitempty" xml:"pullStrategy,omitempty"`
	ReceiverUserIdList []*string                    `json:"receiverUserIdList,omitempty" xml:"receiverUserIdList,omitempty" type:"Repeated"`
	RobotCode          *string                      `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	UserIdType         *int32                       `json:"userIdType,omitempty" xml:"userIdType,omitempty"`
}

func (s SendInteractiveCardRequest) String() string {
	return tea.Prettify(s)
}

func (s SendInteractiveCardRequest) GoString() string {
	return s.String()
}

func (s *SendInteractiveCardRequest) SetAtOpenIds(v map[string]*string) *SendInteractiveCardRequest {
	s.AtOpenIds = v
	return s
}

func (s *SendInteractiveCardRequest) SetCallbackRouteKey(v string) *SendInteractiveCardRequest {
	s.CallbackRouteKey = &v
	return s
}

func (s *SendInteractiveCardRequest) SetCardData(v *SendInteractiveCardRequestCardData) *SendInteractiveCardRequest {
	s.CardData = v
	return s
}

func (s *SendInteractiveCardRequest) SetCardOptions(v *SendInteractiveCardRequestCardOptions) *SendInteractiveCardRequest {
	s.CardOptions = v
	return s
}

func (s *SendInteractiveCardRequest) SetCardTemplateId(v string) *SendInteractiveCardRequest {
	s.CardTemplateId = &v
	return s
}

func (s *SendInteractiveCardRequest) SetChatBotId(v string) *SendInteractiveCardRequest {
	s.ChatBotId = &v
	return s
}

func (s *SendInteractiveCardRequest) SetConversationType(v int32) *SendInteractiveCardRequest {
	s.ConversationType = &v
	return s
}

func (s *SendInteractiveCardRequest) SetDigitalWorkerCode(v string) *SendInteractiveCardRequest {
	s.DigitalWorkerCode = &v
	return s
}

func (s *SendInteractiveCardRequest) SetOpenConversationId(v string) *SendInteractiveCardRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SendInteractiveCardRequest) SetOutTrackId(v string) *SendInteractiveCardRequest {
	s.OutTrackId = &v
	return s
}

func (s *SendInteractiveCardRequest) SetPrivateData(v map[string]*PrivateDataValue) *SendInteractiveCardRequest {
	s.PrivateData = v
	return s
}

func (s *SendInteractiveCardRequest) SetPullStrategy(v bool) *SendInteractiveCardRequest {
	s.PullStrategy = &v
	return s
}

func (s *SendInteractiveCardRequest) SetReceiverUserIdList(v []*string) *SendInteractiveCardRequest {
	s.ReceiverUserIdList = v
	return s
}

func (s *SendInteractiveCardRequest) SetRobotCode(v string) *SendInteractiveCardRequest {
	s.RobotCode = &v
	return s
}

func (s *SendInteractiveCardRequest) SetUserIdType(v int32) *SendInteractiveCardRequest {
	s.UserIdType = &v
	return s
}

type SendInteractiveCardRequestCardData struct {
	CardMediaIdParamMap map[string]*string `json:"cardMediaIdParamMap,omitempty" xml:"cardMediaIdParamMap,omitempty"`
	CardParamMap        map[string]*string `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
}

func (s SendInteractiveCardRequestCardData) String() string {
	return tea.Prettify(s)
}

func (s SendInteractiveCardRequestCardData) GoString() string {
	return s.String()
}

func (s *SendInteractiveCardRequestCardData) SetCardMediaIdParamMap(v map[string]*string) *SendInteractiveCardRequestCardData {
	s.CardMediaIdParamMap = v
	return s
}

func (s *SendInteractiveCardRequestCardData) SetCardParamMap(v map[string]*string) *SendInteractiveCardRequestCardData {
	s.CardParamMap = v
	return s
}

type SendInteractiveCardRequestCardOptions struct {
	SupportForward *bool `json:"supportForward,omitempty" xml:"supportForward,omitempty"`
}

func (s SendInteractiveCardRequestCardOptions) String() string {
	return tea.Prettify(s)
}

func (s SendInteractiveCardRequestCardOptions) GoString() string {
	return s.String()
}

func (s *SendInteractiveCardRequestCardOptions) SetSupportForward(v bool) *SendInteractiveCardRequestCardOptions {
	s.SupportForward = &v
	return s
}

type SendInteractiveCardResponseBody struct {
	Result  *SendInteractiveCardResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SendInteractiveCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendInteractiveCardResponseBody) GoString() string {
	return s.String()
}

func (s *SendInteractiveCardResponseBody) SetResult(v *SendInteractiveCardResponseBodyResult) *SendInteractiveCardResponseBody {
	s.Result = v
	return s
}

func (s *SendInteractiveCardResponseBody) SetSuccess(v bool) *SendInteractiveCardResponseBody {
	s.Success = &v
	return s
}

type SendInteractiveCardResponseBodyResult struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxxx
	ProcessQueryKey *string `json:"processQueryKey,omitempty" xml:"processQueryKey,omitempty"`
}

func (s SendInteractiveCardResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SendInteractiveCardResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SendInteractiveCardResponseBodyResult) SetProcessQueryKey(v string) *SendInteractiveCardResponseBodyResult {
	s.ProcessQueryKey = &v
	return s
}

type SendInteractiveCardResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendInteractiveCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendInteractiveCardResponse) String() string {
	return tea.Prettify(s)
}

func (s SendInteractiveCardResponse) GoString() string {
	return s.String()
}

func (s *SendInteractiveCardResponse) SetHeaders(v map[string]*string) *SendInteractiveCardResponse {
	s.Headers = v
	return s
}

func (s *SendInteractiveCardResponse) SetStatusCode(v int32) *SendInteractiveCardResponse {
	s.StatusCode = &v
	return s
}

func (s *SendInteractiveCardResponse) SetBody(v *SendInteractiveCardResponseBody) *SendInteractiveCardResponse {
	s.Body = v
	return s
}

type SendOTOInteractiveCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendOTOInteractiveCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendOTOInteractiveCardHeaders) GoString() string {
	return s.String()
}

func (s *SendOTOInteractiveCardHeaders) SetCommonHeaders(v map[string]*string) *SendOTOInteractiveCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendOTOInteractiveCardHeaders) SetXAcsDingtalkAccessToken(v string) *SendOTOInteractiveCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendOTOInteractiveCardRequest struct {
	AtOpenIds        map[string]*string `json:"atOpenIds,omitempty" xml:"atOpenIds,omitempty"`
	CallbackRouteKey *string            `json:"callbackRouteKey,omitempty" xml:"callbackRouteKey,omitempty"`
	// This parameter is required.
	CardData    *SendOTOInteractiveCardRequestCardData    `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	CardOptions *SendOTOInteractiveCardRequestCardOptions `json:"cardOptions,omitempty" xml:"cardOptions,omitempty" type:"Struct"`
	// This parameter is required.
	CardTemplateId     *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	OutTrackId         *string                      `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	PrivateData        map[string]*PrivateDataValue `json:"privateData,omitempty" xml:"privateData,omitempty"`
	PullStrategy       *bool                        `json:"pullStrategy,omitempty" xml:"pullStrategy,omitempty"`
	ReceiverUserIdList []*string                    `json:"receiverUserIdList,omitempty" xml:"receiverUserIdList,omitempty" type:"Repeated"`
	RobotCode          *string                      `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	UserIdType         *int32                       `json:"userIdType,omitempty" xml:"userIdType,omitempty"`
}

func (s SendOTOInteractiveCardRequest) String() string {
	return tea.Prettify(s)
}

func (s SendOTOInteractiveCardRequest) GoString() string {
	return s.String()
}

func (s *SendOTOInteractiveCardRequest) SetAtOpenIds(v map[string]*string) *SendOTOInteractiveCardRequest {
	s.AtOpenIds = v
	return s
}

func (s *SendOTOInteractiveCardRequest) SetCallbackRouteKey(v string) *SendOTOInteractiveCardRequest {
	s.CallbackRouteKey = &v
	return s
}

func (s *SendOTOInteractiveCardRequest) SetCardData(v *SendOTOInteractiveCardRequestCardData) *SendOTOInteractiveCardRequest {
	s.CardData = v
	return s
}

func (s *SendOTOInteractiveCardRequest) SetCardOptions(v *SendOTOInteractiveCardRequestCardOptions) *SendOTOInteractiveCardRequest {
	s.CardOptions = v
	return s
}

func (s *SendOTOInteractiveCardRequest) SetCardTemplateId(v string) *SendOTOInteractiveCardRequest {
	s.CardTemplateId = &v
	return s
}

func (s *SendOTOInteractiveCardRequest) SetOpenConversationId(v string) *SendOTOInteractiveCardRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SendOTOInteractiveCardRequest) SetOutTrackId(v string) *SendOTOInteractiveCardRequest {
	s.OutTrackId = &v
	return s
}

func (s *SendOTOInteractiveCardRequest) SetPrivateData(v map[string]*PrivateDataValue) *SendOTOInteractiveCardRequest {
	s.PrivateData = v
	return s
}

func (s *SendOTOInteractiveCardRequest) SetPullStrategy(v bool) *SendOTOInteractiveCardRequest {
	s.PullStrategy = &v
	return s
}

func (s *SendOTOInteractiveCardRequest) SetReceiverUserIdList(v []*string) *SendOTOInteractiveCardRequest {
	s.ReceiverUserIdList = v
	return s
}

func (s *SendOTOInteractiveCardRequest) SetRobotCode(v string) *SendOTOInteractiveCardRequest {
	s.RobotCode = &v
	return s
}

func (s *SendOTOInteractiveCardRequest) SetUserIdType(v int32) *SendOTOInteractiveCardRequest {
	s.UserIdType = &v
	return s
}

type SendOTOInteractiveCardRequestCardData struct {
	CardParamMap map[string]*string `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
}

func (s SendOTOInteractiveCardRequestCardData) String() string {
	return tea.Prettify(s)
}

func (s SendOTOInteractiveCardRequestCardData) GoString() string {
	return s.String()
}

func (s *SendOTOInteractiveCardRequestCardData) SetCardParamMap(v map[string]*string) *SendOTOInteractiveCardRequestCardData {
	s.CardParamMap = v
	return s
}

type SendOTOInteractiveCardRequestCardOptions struct {
	SupportForward *bool `json:"supportForward,omitempty" xml:"supportForward,omitempty"`
}

func (s SendOTOInteractiveCardRequestCardOptions) String() string {
	return tea.Prettify(s)
}

func (s SendOTOInteractiveCardRequestCardOptions) GoString() string {
	return s.String()
}

func (s *SendOTOInteractiveCardRequestCardOptions) SetSupportForward(v bool) *SendOTOInteractiveCardRequestCardOptions {
	s.SupportForward = &v
	return s
}

type SendOTOInteractiveCardResponseBody struct {
	Result  *SendOTOInteractiveCardResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SendOTOInteractiveCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendOTOInteractiveCardResponseBody) GoString() string {
	return s.String()
}

func (s *SendOTOInteractiveCardResponseBody) SetResult(v *SendOTOInteractiveCardResponseBodyResult) *SendOTOInteractiveCardResponseBody {
	s.Result = v
	return s
}

func (s *SendOTOInteractiveCardResponseBody) SetSuccess(v bool) *SendOTOInteractiveCardResponseBody {
	s.Success = &v
	return s
}

type SendOTOInteractiveCardResponseBodyResult struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxxx
	ProcessQueryKey *string `json:"processQueryKey,omitempty" xml:"processQueryKey,omitempty"`
}

func (s SendOTOInteractiveCardResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SendOTOInteractiveCardResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SendOTOInteractiveCardResponseBodyResult) SetProcessQueryKey(v string) *SendOTOInteractiveCardResponseBodyResult {
	s.ProcessQueryKey = &v
	return s
}

type SendOTOInteractiveCardResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendOTOInteractiveCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendOTOInteractiveCardResponse) String() string {
	return tea.Prettify(s)
}

func (s SendOTOInteractiveCardResponse) GoString() string {
	return s.String()
}

func (s *SendOTOInteractiveCardResponse) SetHeaders(v map[string]*string) *SendOTOInteractiveCardResponse {
	s.Headers = v
	return s
}

func (s *SendOTOInteractiveCardResponse) SetStatusCode(v int32) *SendOTOInteractiveCardResponse {
	s.StatusCode = &v
	return s
}

func (s *SendOTOInteractiveCardResponse) SetBody(v *SendOTOInteractiveCardResponseBody) *SendOTOInteractiveCardResponse {
	s.Body = v
	return s
}

type SendPersonalMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendPersonalMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendPersonalMessageHeaders) GoString() string {
	return s.String()
}

func (s *SendPersonalMessageHeaders) SetCommonHeaders(v map[string]*string) *SendPersonalMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendPersonalMessageHeaders) SetXAcsDingtalkAccessToken(v string) *SendPersonalMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendPersonalMessageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {"content":"月会通知<@all> ","at":{"atUserIds":[],"isAtAll":true}}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// text
	MsgType *string `json:"msgType,omitempty" xml:"msgType,omitempty"`
	// example:
	//
	// cidc4iLyQBuHFQRvzxznz204Q==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// 1662055829854977
	ReceiverUserId *string `json:"receiverUserId,omitempty" xml:"receiverUserId,omitempty"`
}

func (s SendPersonalMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendPersonalMessageRequest) GoString() string {
	return s.String()
}

func (s *SendPersonalMessageRequest) SetContent(v string) *SendPersonalMessageRequest {
	s.Content = &v
	return s
}

func (s *SendPersonalMessageRequest) SetMsgType(v string) *SendPersonalMessageRequest {
	s.MsgType = &v
	return s
}

func (s *SendPersonalMessageRequest) SetOpenConversationId(v string) *SendPersonalMessageRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SendPersonalMessageRequest) SetReceiverUserId(v string) *SendPersonalMessageRequest {
	s.ReceiverUserId = &v
	return s
}

type SendPersonalMessageResponseBody struct {
	Result  *SendPersonalMessageResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *string                                `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SendPersonalMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendPersonalMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendPersonalMessageResponseBody) SetResult(v *SendPersonalMessageResponseBodyResult) *SendPersonalMessageResponseBody {
	s.Result = v
	return s
}

func (s *SendPersonalMessageResponseBody) SetSuccess(v string) *SendPersonalMessageResponseBody {
	s.Success = &v
	return s
}

type SendPersonalMessageResponseBodyResult struct {
	OpenTaskId *string `json:"openTaskId,omitempty" xml:"openTaskId,omitempty"`
}

func (s SendPersonalMessageResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SendPersonalMessageResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SendPersonalMessageResponseBodyResult) SetOpenTaskId(v string) *SendPersonalMessageResponseBodyResult {
	s.OpenTaskId = &v
	return s
}

type SendPersonalMessageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendPersonalMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendPersonalMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendPersonalMessageResponse) GoString() string {
	return s.String()
}

func (s *SendPersonalMessageResponse) SetHeaders(v map[string]*string) *SendPersonalMessageResponse {
	s.Headers = v
	return s
}

func (s *SendPersonalMessageResponse) SetStatusCode(v int32) *SendPersonalMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendPersonalMessageResponse) SetBody(v *SendPersonalMessageResponseBody) *SendPersonalMessageResponse {
	s.Body = v
	return s
}

type SendRobotInteractiveCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendRobotInteractiveCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendRobotInteractiveCardHeaders) GoString() string {
	return s.String()
}

func (s *SendRobotInteractiveCardHeaders) SetCommonHeaders(v map[string]*string) *SendRobotInteractiveCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendRobotInteractiveCardHeaders) SetXAcsDingtalkAccessToken(v string) *SendRobotInteractiveCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendRobotInteractiveCardRequest struct {
	// example:
	//
	// https://xxx
	CallbackUrl *string `json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cardXXXX01
	CardBizId *string `json:"cardBizId,omitempty" xml:"cardBizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 根据具体的cardTemplateId参考文档格式
	CardData *string `json:"cardData,omitempty" xml:"cardData,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxxxxxx
	CardTemplateId *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	// example:
	//
	// cidXXXX
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	PullStrategy       *bool   `json:"pullStrategy,omitempty" xml:"pullStrategy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxxxx
	RobotCode   *string                                     `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	SendOptions *SendRobotInteractiveCardRequestSendOptions `json:"sendOptions,omitempty" xml:"sendOptions,omitempty" type:"Struct"`
	// example:
	//
	// 以userId为例：{"userId":"userId0001"}；以unionId为例{"unionId":"unionId001"}
	SingleChatReceiver    *string `json:"singleChatReceiver,omitempty" xml:"singleChatReceiver,omitempty"`
	UnionIdPrivateDataMap *string `json:"unionIdPrivateDataMap,omitempty" xml:"unionIdPrivateDataMap,omitempty"`
	UserIdPrivateDataMap  *string `json:"userIdPrivateDataMap,omitempty" xml:"userIdPrivateDataMap,omitempty"`
}

func (s SendRobotInteractiveCardRequest) String() string {
	return tea.Prettify(s)
}

func (s SendRobotInteractiveCardRequest) GoString() string {
	return s.String()
}

func (s *SendRobotInteractiveCardRequest) SetCallbackUrl(v string) *SendRobotInteractiveCardRequest {
	s.CallbackUrl = &v
	return s
}

func (s *SendRobotInteractiveCardRequest) SetCardBizId(v string) *SendRobotInteractiveCardRequest {
	s.CardBizId = &v
	return s
}

func (s *SendRobotInteractiveCardRequest) SetCardData(v string) *SendRobotInteractiveCardRequest {
	s.CardData = &v
	return s
}

func (s *SendRobotInteractiveCardRequest) SetCardTemplateId(v string) *SendRobotInteractiveCardRequest {
	s.CardTemplateId = &v
	return s
}

func (s *SendRobotInteractiveCardRequest) SetOpenConversationId(v string) *SendRobotInteractiveCardRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SendRobotInteractiveCardRequest) SetPullStrategy(v bool) *SendRobotInteractiveCardRequest {
	s.PullStrategy = &v
	return s
}

func (s *SendRobotInteractiveCardRequest) SetRobotCode(v string) *SendRobotInteractiveCardRequest {
	s.RobotCode = &v
	return s
}

func (s *SendRobotInteractiveCardRequest) SetSendOptions(v *SendRobotInteractiveCardRequestSendOptions) *SendRobotInteractiveCardRequest {
	s.SendOptions = v
	return s
}

func (s *SendRobotInteractiveCardRequest) SetSingleChatReceiver(v string) *SendRobotInteractiveCardRequest {
	s.SingleChatReceiver = &v
	return s
}

func (s *SendRobotInteractiveCardRequest) SetUnionIdPrivateDataMap(v string) *SendRobotInteractiveCardRequest {
	s.UnionIdPrivateDataMap = &v
	return s
}

func (s *SendRobotInteractiveCardRequest) SetUserIdPrivateDataMap(v string) *SendRobotInteractiveCardRequest {
	s.UserIdPrivateDataMap = &v
	return s
}

type SendRobotInteractiveCardRequestSendOptions struct {
	// example:
	//
	// true
	AtAll *bool `json:"atAll,omitempty" xml:"atAll,omitempty"`
	// example:
	//
	// [{"nickName":"张三","userId":"userId0001"},{"nickName":"李四","unionId":"unionId001"}]
	AtUserListJson *string `json:"atUserListJson,omitempty" xml:"atUserListJson,omitempty"`
	// example:
	//
	// {}
	CardPropertyJson *string `json:"cardPropertyJson,omitempty" xml:"cardPropertyJson,omitempty"`
	// example:
	//
	// [{"userId":"userId0001"},{"unionId":"unionId001"}]
	ReceiverListJson *string `json:"receiverListJson,omitempty" xml:"receiverListJson,omitempty"`
}

func (s SendRobotInteractiveCardRequestSendOptions) String() string {
	return tea.Prettify(s)
}

func (s SendRobotInteractiveCardRequestSendOptions) GoString() string {
	return s.String()
}

func (s *SendRobotInteractiveCardRequestSendOptions) SetAtAll(v bool) *SendRobotInteractiveCardRequestSendOptions {
	s.AtAll = &v
	return s
}

func (s *SendRobotInteractiveCardRequestSendOptions) SetAtUserListJson(v string) *SendRobotInteractiveCardRequestSendOptions {
	s.AtUserListJson = &v
	return s
}

func (s *SendRobotInteractiveCardRequestSendOptions) SetCardPropertyJson(v string) *SendRobotInteractiveCardRequestSendOptions {
	s.CardPropertyJson = &v
	return s
}

func (s *SendRobotInteractiveCardRequestSendOptions) SetReceiverListJson(v string) *SendRobotInteractiveCardRequestSendOptions {
	s.ReceiverListJson = &v
	return s
}

type SendRobotInteractiveCardResponseBody struct {
	// example:
	//
	// xxxxxx
	ProcessQueryKey *string `json:"processQueryKey,omitempty" xml:"processQueryKey,omitempty"`
}

func (s SendRobotInteractiveCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendRobotInteractiveCardResponseBody) GoString() string {
	return s.String()
}

func (s *SendRobotInteractiveCardResponseBody) SetProcessQueryKey(v string) *SendRobotInteractiveCardResponseBody {
	s.ProcessQueryKey = &v
	return s
}

type SendRobotInteractiveCardResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendRobotInteractiveCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendRobotInteractiveCardResponse) String() string {
	return tea.Prettify(s)
}

func (s SendRobotInteractiveCardResponse) GoString() string {
	return s.String()
}

func (s *SendRobotInteractiveCardResponse) SetHeaders(v map[string]*string) *SendRobotInteractiveCardResponse {
	s.Headers = v
	return s
}

func (s *SendRobotInteractiveCardResponse) SetStatusCode(v int32) *SendRobotInteractiveCardResponse {
	s.StatusCode = &v
	return s
}

func (s *SendRobotInteractiveCardResponse) SetBody(v *SendRobotInteractiveCardResponseBody) *SendRobotInteractiveCardResponse {
	s.Body = v
	return s
}

type SendRobotMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendRobotMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendRobotMessageHeaders) GoString() string {
	return s.String()
}

func (s *SendRobotMessageHeaders) SetCommonHeaders(v map[string]*string) *SendRobotMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendRobotMessageHeaders) SetXAcsDingtalkAccessToken(v string) *SendRobotMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendRobotMessageRequest struct {
	AtAll *bool `json:"atAll,omitempty" xml:"atAll,omitempty"`
	// example:
	//
	// 1107****2120
	AtAppUserId *string `json:"atAppUserId,omitempty" xml:"atAppUserId,omitempty"`
	// example:
	//
	// 1107****2120
	AtDingUserId *string `json:"atDingUserId,omitempty" xml:"atDingUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// { "content": "我就是我, 是不一样的烟火"}
	MsgContent *string `json:"msgContent,omitempty" xml:"msgContent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// text
	MsgType *string `json:"msgType,omitempty" xml:"msgType,omitempty"`
	// This parameter is required.
	OpenConversationIds []*string `json:"openConversationIds,omitempty" xml:"openConversationIds,omitempty" type:"Repeated"`
	// example:
	//
	// kelian-custom-service-robot-101
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s SendRobotMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendRobotMessageRequest) GoString() string {
	return s.String()
}

func (s *SendRobotMessageRequest) SetAtAll(v bool) *SendRobotMessageRequest {
	s.AtAll = &v
	return s
}

func (s *SendRobotMessageRequest) SetAtAppUserId(v string) *SendRobotMessageRequest {
	s.AtAppUserId = &v
	return s
}

func (s *SendRobotMessageRequest) SetAtDingUserId(v string) *SendRobotMessageRequest {
	s.AtDingUserId = &v
	return s
}

func (s *SendRobotMessageRequest) SetMsgContent(v string) *SendRobotMessageRequest {
	s.MsgContent = &v
	return s
}

func (s *SendRobotMessageRequest) SetMsgType(v string) *SendRobotMessageRequest {
	s.MsgType = &v
	return s
}

func (s *SendRobotMessageRequest) SetOpenConversationIds(v []*string) *SendRobotMessageRequest {
	s.OpenConversationIds = v
	return s
}

func (s *SendRobotMessageRequest) SetRobotCode(v string) *SendRobotMessageRequest {
	s.RobotCode = &v
	return s
}

type SendRobotMessageResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SendRobotMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendRobotMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendRobotMessageResponseBody) SetSuccess(v bool) *SendRobotMessageResponseBody {
	s.Success = &v
	return s
}

type SendRobotMessageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendRobotMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendRobotMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendRobotMessageResponse) GoString() string {
	return s.String()
}

func (s *SendRobotMessageResponse) SetHeaders(v map[string]*string) *SendRobotMessageResponse {
	s.Headers = v
	return s
}

func (s *SendRobotMessageResponse) SetStatusCode(v int32) *SendRobotMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendRobotMessageResponse) SetBody(v *SendRobotMessageResponseBody) *SendRobotMessageResponse {
	s.Body = v
	return s
}

type SendTemplateInteractiveCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendTemplateInteractiveCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendTemplateInteractiveCardHeaders) GoString() string {
	return s.String()
}

func (s *SendTemplateInteractiveCardHeaders) SetCommonHeaders(v map[string]*string) *SendTemplateInteractiveCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendTemplateInteractiveCardHeaders) SetXAcsDingtalkAccessToken(v string) *SendTemplateInteractiveCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendTemplateInteractiveCardRequest struct {
	// example:
	//
	// https://xxxx.com/.../
	CallbackUrl *string `json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 根据具体的cardTemplateId参考文档格式
	CardData *string `json:"cardData,omitempty" xml:"cardData,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TuWenCard01
	CardTemplateId *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	// example:
	//
	// cidXXXX
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cardXXXX01
	OutTrackId *string `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxxxx
	RobotCode   *string                                        `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	SendOptions *SendTemplateInteractiveCardRequestSendOptions `json:"sendOptions,omitempty" xml:"sendOptions,omitempty" type:"Struct"`
	// example:
	//
	// 以userId为例：{"userId":"userId0001"}；以unionId为例{"unionId":"unionId001"}
	SingleChatReceiver *string `json:"singleChatReceiver,omitempty" xml:"singleChatReceiver,omitempty"`
}

func (s SendTemplateInteractiveCardRequest) String() string {
	return tea.Prettify(s)
}

func (s SendTemplateInteractiveCardRequest) GoString() string {
	return s.String()
}

func (s *SendTemplateInteractiveCardRequest) SetCallbackUrl(v string) *SendTemplateInteractiveCardRequest {
	s.CallbackUrl = &v
	return s
}

func (s *SendTemplateInteractiveCardRequest) SetCardData(v string) *SendTemplateInteractiveCardRequest {
	s.CardData = &v
	return s
}

func (s *SendTemplateInteractiveCardRequest) SetCardTemplateId(v string) *SendTemplateInteractiveCardRequest {
	s.CardTemplateId = &v
	return s
}

func (s *SendTemplateInteractiveCardRequest) SetOpenConversationId(v string) *SendTemplateInteractiveCardRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SendTemplateInteractiveCardRequest) SetOutTrackId(v string) *SendTemplateInteractiveCardRequest {
	s.OutTrackId = &v
	return s
}

func (s *SendTemplateInteractiveCardRequest) SetRobotCode(v string) *SendTemplateInteractiveCardRequest {
	s.RobotCode = &v
	return s
}

func (s *SendTemplateInteractiveCardRequest) SetSendOptions(v *SendTemplateInteractiveCardRequestSendOptions) *SendTemplateInteractiveCardRequest {
	s.SendOptions = v
	return s
}

func (s *SendTemplateInteractiveCardRequest) SetSingleChatReceiver(v string) *SendTemplateInteractiveCardRequest {
	s.SingleChatReceiver = &v
	return s
}

type SendTemplateInteractiveCardRequestSendOptions struct {
	// example:
	//
	// true
	AtAll *bool `json:"atAll,omitempty" xml:"atAll,omitempty"`
	// example:
	//
	// [{"nickName":"张三","userId":"userId0001"},{"nickName":"李四","unionId":"unionId001"}]
	AtUserListJson *string `json:"atUserListJson,omitempty" xml:"atUserListJson,omitempty"`
	// example:
	//
	// {}
	CardPropertyJson *string `json:"cardPropertyJson,omitempty" xml:"cardPropertyJson,omitempty"`
	// example:
	//
	// [{"userId":"userId0001"},{"unionId":"unionId001"}]
	ReceiverListJson *string `json:"receiverListJson,omitempty" xml:"receiverListJson,omitempty"`
}

func (s SendTemplateInteractiveCardRequestSendOptions) String() string {
	return tea.Prettify(s)
}

func (s SendTemplateInteractiveCardRequestSendOptions) GoString() string {
	return s.String()
}

func (s *SendTemplateInteractiveCardRequestSendOptions) SetAtAll(v bool) *SendTemplateInteractiveCardRequestSendOptions {
	s.AtAll = &v
	return s
}

func (s *SendTemplateInteractiveCardRequestSendOptions) SetAtUserListJson(v string) *SendTemplateInteractiveCardRequestSendOptions {
	s.AtUserListJson = &v
	return s
}

func (s *SendTemplateInteractiveCardRequestSendOptions) SetCardPropertyJson(v string) *SendTemplateInteractiveCardRequestSendOptions {
	s.CardPropertyJson = &v
	return s
}

func (s *SendTemplateInteractiveCardRequestSendOptions) SetReceiverListJson(v string) *SendTemplateInteractiveCardRequestSendOptions {
	s.ReceiverListJson = &v
	return s
}

type SendTemplateInteractiveCardResponseBody struct {
	// example:
	//
	// xxxxxx
	ProcessQueryKey *string `json:"processQueryKey,omitempty" xml:"processQueryKey,omitempty"`
}

func (s SendTemplateInteractiveCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendTemplateInteractiveCardResponseBody) GoString() string {
	return s.String()
}

func (s *SendTemplateInteractiveCardResponseBody) SetProcessQueryKey(v string) *SendTemplateInteractiveCardResponseBody {
	s.ProcessQueryKey = &v
	return s
}

type SendTemplateInteractiveCardResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendTemplateInteractiveCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendTemplateInteractiveCardResponse) String() string {
	return tea.Prettify(s)
}

func (s SendTemplateInteractiveCardResponse) GoString() string {
	return s.String()
}

func (s *SendTemplateInteractiveCardResponse) SetHeaders(v map[string]*string) *SendTemplateInteractiveCardResponse {
	s.Headers = v
	return s
}

func (s *SendTemplateInteractiveCardResponse) SetStatusCode(v int32) *SendTemplateInteractiveCardResponse {
	s.StatusCode = &v
	return s
}

func (s *SendTemplateInteractiveCardResponse) SetBody(v *SendTemplateInteractiveCardResponseBody) *SendTemplateInteractiveCardResponse {
	s.Body = v
	return s
}

type SetRightPanelHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SetRightPanelHeaders) String() string {
	return tea.Prettify(s)
}

func (s SetRightPanelHeaders) GoString() string {
	return s.String()
}

func (s *SetRightPanelHeaders) SetCommonHeaders(v map[string]*string) *SetRightPanelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetRightPanelHeaders) SetXAcsDingtalkAccessToken(v string) *SetRightPanelHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SetRightPanelRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ciddjxhgdDXSAAXXXXX
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	RightPanelClosePermitted *bool `json:"rightPanelClosePermitted,omitempty" xml:"rightPanelClosePermitted,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	RightPanelOpenStatus *int32 `json:"rightPanelOpenStatus,omitempty" xml:"rightPanelOpenStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 侧边栏标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	WebWndParams *SetRightPanelRequestWebWndParams `json:"webWndParams,omitempty" xml:"webWndParams,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 500
	Width *int32 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s SetRightPanelRequest) String() string {
	return tea.Prettify(s)
}

func (s SetRightPanelRequest) GoString() string {
	return s.String()
}

func (s *SetRightPanelRequest) SetOpenConversationId(v string) *SetRightPanelRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SetRightPanelRequest) SetRightPanelClosePermitted(v bool) *SetRightPanelRequest {
	s.RightPanelClosePermitted = &v
	return s
}

func (s *SetRightPanelRequest) SetRightPanelOpenStatus(v int32) *SetRightPanelRequest {
	s.RightPanelOpenStatus = &v
	return s
}

func (s *SetRightPanelRequest) SetTitle(v string) *SetRightPanelRequest {
	s.Title = &v
	return s
}

func (s *SetRightPanelRequest) SetWebWndParams(v *SetRightPanelRequestWebWndParams) *SetRightPanelRequest {
	s.WebWndParams = v
	return s
}

func (s *SetRightPanelRequest) SetWidth(v int32) *SetRightPanelRequest {
	s.Width = &v
	return s
}

type SetRightPanelRequestWebWndParams struct {
	// This parameter is required.
	//
	// example:
	//
	// https://www.dingtalk.com/
	TargetURL *string `json:"targetURL,omitempty" xml:"targetURL,omitempty"`
}

func (s SetRightPanelRequestWebWndParams) String() string {
	return tea.Prettify(s)
}

func (s SetRightPanelRequestWebWndParams) GoString() string {
	return s.String()
}

func (s *SetRightPanelRequestWebWndParams) SetTargetURL(v string) *SetRightPanelRequestWebWndParams {
	s.TargetURL = &v
	return s
}

type SetRightPanelResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SetRightPanelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetRightPanelResponseBody) GoString() string {
	return s.String()
}

func (s *SetRightPanelResponseBody) SetSuccess(v bool) *SetRightPanelResponseBody {
	s.Success = &v
	return s
}

type SetRightPanelResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetRightPanelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetRightPanelResponse) String() string {
	return tea.Prettify(s)
}

func (s SetRightPanelResponse) GoString() string {
	return s.String()
}

func (s *SetRightPanelResponse) SetHeaders(v map[string]*string) *SetRightPanelResponse {
	s.Headers = v
	return s
}

func (s *SetRightPanelResponse) SetStatusCode(v int32) *SetRightPanelResponse {
	s.StatusCode = &v
	return s
}

func (s *SetRightPanelResponse) SetBody(v *SetRightPanelResponseBody) *SetRightPanelResponse {
	s.Body = v
	return s
}

type SuperAdminApplyTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SuperAdminApplyTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s SuperAdminApplyTemplateHeaders) GoString() string {
	return s.String()
}

func (s *SuperAdminApplyTemplateHeaders) SetCommonHeaders(v map[string]*string) *SuperAdminApplyTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SuperAdminApplyTemplateHeaders) SetXAcsDingtalkAccessToken(v string) *SuperAdminApplyTemplateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SuperAdminApplyTemplateRequest struct {
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	OwnerUserId *string `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	// This parameter is required.
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s SuperAdminApplyTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s SuperAdminApplyTemplateRequest) GoString() string {
	return s.String()
}

func (s *SuperAdminApplyTemplateRequest) SetOpenConversationId(v string) *SuperAdminApplyTemplateRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SuperAdminApplyTemplateRequest) SetOwnerUserId(v string) *SuperAdminApplyTemplateRequest {
	s.OwnerUserId = &v
	return s
}

func (s *SuperAdminApplyTemplateRequest) SetTemplateId(v string) *SuperAdminApplyTemplateRequest {
	s.TemplateId = &v
	return s
}

type SuperAdminApplyTemplateResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SuperAdminApplyTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SuperAdminApplyTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *SuperAdminApplyTemplateResponseBody) SetSuccess(v bool) *SuperAdminApplyTemplateResponseBody {
	s.Success = &v
	return s
}

type SuperAdminApplyTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuperAdminApplyTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuperAdminApplyTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s SuperAdminApplyTemplateResponse) GoString() string {
	return s.String()
}

func (s *SuperAdminApplyTemplateResponse) SetHeaders(v map[string]*string) *SuperAdminApplyTemplateResponse {
	s.Headers = v
	return s
}

func (s *SuperAdminApplyTemplateResponse) SetStatusCode(v int32) *SuperAdminApplyTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *SuperAdminApplyTemplateResponse) SetBody(v *SuperAdminApplyTemplateResponseBody) *SuperAdminApplyTemplateResponse {
	s.Body = v
	return s
}

type SuperAdminCloseTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SuperAdminCloseTemplateHeaders) String() string {
	return tea.Prettify(s)
}

func (s SuperAdminCloseTemplateHeaders) GoString() string {
	return s.String()
}

func (s *SuperAdminCloseTemplateHeaders) SetCommonHeaders(v map[string]*string) *SuperAdminCloseTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SuperAdminCloseTemplateHeaders) SetXAcsDingtalkAccessToken(v string) *SuperAdminCloseTemplateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SuperAdminCloseTemplateRequest struct {
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	OwnerUserId *string `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	// This parameter is required.
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s SuperAdminCloseTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s SuperAdminCloseTemplateRequest) GoString() string {
	return s.String()
}

func (s *SuperAdminCloseTemplateRequest) SetOpenConversationId(v string) *SuperAdminCloseTemplateRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SuperAdminCloseTemplateRequest) SetOwnerUserId(v string) *SuperAdminCloseTemplateRequest {
	s.OwnerUserId = &v
	return s
}

func (s *SuperAdminCloseTemplateRequest) SetTemplateId(v string) *SuperAdminCloseTemplateRequest {
	s.TemplateId = &v
	return s
}

type SuperAdminCloseTemplateResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SuperAdminCloseTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SuperAdminCloseTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *SuperAdminCloseTemplateResponseBody) SetSuccess(v bool) *SuperAdminCloseTemplateResponseBody {
	s.Success = &v
	return s
}

type SuperAdminCloseTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuperAdminCloseTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuperAdminCloseTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s SuperAdminCloseTemplateResponse) GoString() string {
	return s.String()
}

func (s *SuperAdminCloseTemplateResponse) SetHeaders(v map[string]*string) *SuperAdminCloseTemplateResponse {
	s.Headers = v
	return s
}

func (s *SuperAdminCloseTemplateResponse) SetStatusCode(v int32) *SuperAdminCloseTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *SuperAdminCloseTemplateResponse) SetBody(v *SuperAdminCloseTemplateResponseBody) *SuperAdminCloseTemplateResponse {
	s.Body = v
	return s
}

type TopboxCloseHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TopboxCloseHeaders) String() string {
	return tea.Prettify(s)
}

func (s TopboxCloseHeaders) GoString() string {
	return s.String()
}

func (s *TopboxCloseHeaders) SetCommonHeaders(v map[string]*string) *TopboxCloseHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TopboxCloseHeaders) SetXAcsDingtalkAccessToken(v string) *TopboxCloseHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TopboxCloseRequest struct {
	ConversationType *int32  `json:"conversationType,omitempty" xml:"conversationType,omitempty"`
	CoolAppCode      *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	// example:
	//
	// xxxx
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	OutTrackId         *string   `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	ReceiverUserIdList []*string `json:"receiverUserIdList,omitempty" xml:"receiverUserIdList,omitempty" type:"Repeated"`
	RobotCode          *string   `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s TopboxCloseRequest) String() string {
	return tea.Prettify(s)
}

func (s TopboxCloseRequest) GoString() string {
	return s.String()
}

func (s *TopboxCloseRequest) SetConversationType(v int32) *TopboxCloseRequest {
	s.ConversationType = &v
	return s
}

func (s *TopboxCloseRequest) SetCoolAppCode(v string) *TopboxCloseRequest {
	s.CoolAppCode = &v
	return s
}

func (s *TopboxCloseRequest) SetOpenConversationId(v string) *TopboxCloseRequest {
	s.OpenConversationId = &v
	return s
}

func (s *TopboxCloseRequest) SetOutTrackId(v string) *TopboxCloseRequest {
	s.OutTrackId = &v
	return s
}

func (s *TopboxCloseRequest) SetReceiverUserIdList(v []*string) *TopboxCloseRequest {
	s.ReceiverUserIdList = v
	return s
}

func (s *TopboxCloseRequest) SetRobotCode(v string) *TopboxCloseRequest {
	s.RobotCode = &v
	return s
}

type TopboxCloseResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s TopboxCloseResponse) String() string {
	return tea.Prettify(s)
}

func (s TopboxCloseResponse) GoString() string {
	return s.String()
}

func (s *TopboxCloseResponse) SetHeaders(v map[string]*string) *TopboxCloseResponse {
	s.Headers = v
	return s
}

func (s *TopboxCloseResponse) SetStatusCode(v int32) *TopboxCloseResponse {
	s.StatusCode = &v
	return s
}

type TopboxOpenHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TopboxOpenHeaders) String() string {
	return tea.Prettify(s)
}

func (s TopboxOpenHeaders) GoString() string {
	return s.String()
}

func (s *TopboxOpenHeaders) SetCommonHeaders(v map[string]*string) *TopboxOpenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TopboxOpenHeaders) SetXAcsDingtalkAccessToken(v string) *TopboxOpenHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TopboxOpenRequest struct {
	ConversationType *int32  `json:"conversationType,omitempty" xml:"conversationType,omitempty"`
	CoolAppCode      *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	// example:
	//
	// 1850042969000
	ExpiredTime *int64 `json:"expiredTime,omitempty" xml:"expiredTime,omitempty"`
	// example:
	//
	// xxxx
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	OutTrackId *string `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	// example:
	//
	// ios|win
	Platforms *string `json:"platforms,omitempty" xml:"platforms,omitempty"`
	// if can be null:
	// true
	ReceiverUserIdList []*string `json:"receiverUserIdList,omitempty" xml:"receiverUserIdList,omitempty" type:"Repeated"`
	RobotCode          *string   `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s TopboxOpenRequest) String() string {
	return tea.Prettify(s)
}

func (s TopboxOpenRequest) GoString() string {
	return s.String()
}

func (s *TopboxOpenRequest) SetConversationType(v int32) *TopboxOpenRequest {
	s.ConversationType = &v
	return s
}

func (s *TopboxOpenRequest) SetCoolAppCode(v string) *TopboxOpenRequest {
	s.CoolAppCode = &v
	return s
}

func (s *TopboxOpenRequest) SetExpiredTime(v int64) *TopboxOpenRequest {
	s.ExpiredTime = &v
	return s
}

func (s *TopboxOpenRequest) SetOpenConversationId(v string) *TopboxOpenRequest {
	s.OpenConversationId = &v
	return s
}

func (s *TopboxOpenRequest) SetOutTrackId(v string) *TopboxOpenRequest {
	s.OutTrackId = &v
	return s
}

func (s *TopboxOpenRequest) SetPlatforms(v string) *TopboxOpenRequest {
	s.Platforms = &v
	return s
}

func (s *TopboxOpenRequest) SetReceiverUserIdList(v []*string) *TopboxOpenRequest {
	s.ReceiverUserIdList = v
	return s
}

func (s *TopboxOpenRequest) SetRobotCode(v string) *TopboxOpenRequest {
	s.RobotCode = &v
	return s
}

type TopboxOpenResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s TopboxOpenResponse) String() string {
	return tea.Prettify(s)
}

func (s TopboxOpenResponse) GoString() string {
	return s.String()
}

func (s *TopboxOpenResponse) SetHeaders(v map[string]*string) *TopboxOpenResponse {
	s.Headers = v
	return s
}

func (s *TopboxOpenResponse) SetStatusCode(v int32) *TopboxOpenResponse {
	s.StatusCode = &v
	return s
}

type UpdateClientServiceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateClientServiceHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateClientServiceHeaders) GoString() string {
	return s.String()
}

func (s *UpdateClientServiceHeaders) SetCommonHeaders(v map[string]*string) *UpdateClientServiceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateClientServiceHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateClientServiceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateClientServiceRequest struct {
	// example:
	//
	// https://***.png
	AvatarUrl *string `json:"avatarUrl,omitempty" xml:"avatarUrl,omitempty"`
	// example:
	//
	// false
	ResetAvatar *bool `json:"resetAvatar,omitempty" xml:"resetAvatar,omitempty"`
	// example:
	//
	// false
	ResetUserName *bool `json:"resetUserName,omitempty" xml:"resetUserName,omitempty"`
	// This parameter is required.
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
	// example:
	//
	// test
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s UpdateClientServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateClientServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateClientServiceRequest) SetAvatarUrl(v string) *UpdateClientServiceRequest {
	s.AvatarUrl = &v
	return s
}

func (s *UpdateClientServiceRequest) SetResetAvatar(v bool) *UpdateClientServiceRequest {
	s.ResetAvatar = &v
	return s
}

func (s *UpdateClientServiceRequest) SetResetUserName(v bool) *UpdateClientServiceRequest {
	s.ResetUserName = &v
	return s
}

func (s *UpdateClientServiceRequest) SetUserIds(v []*string) *UpdateClientServiceRequest {
	s.UserIds = v
	return s
}

func (s *UpdateClientServiceRequest) SetUserName(v string) *UpdateClientServiceRequest {
	s.UserName = &v
	return s
}

type UpdateClientServiceResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateClientServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateClientServiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClientServiceResponseBody) SetResult(v bool) *UpdateClientServiceResponseBody {
	s.Result = &v
	return s
}

type UpdateClientServiceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateClientServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateClientServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateClientServiceResponse) GoString() string {
	return s.String()
}

func (s *UpdateClientServiceResponse) SetHeaders(v map[string]*string) *UpdateClientServiceResponse {
	s.Headers = v
	return s
}

func (s *UpdateClientServiceResponse) SetStatusCode(v int32) *UpdateClientServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateClientServiceResponse) SetBody(v *UpdateClientServiceResponseBody) *UpdateClientServiceResponse {
	s.Body = v
	return s
}

type UpdateGroupAvatarHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateGroupAvatarHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupAvatarHeaders) GoString() string {
	return s.String()
}

func (s *UpdateGroupAvatarHeaders) SetCommonHeaders(v map[string]*string) *UpdateGroupAvatarHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateGroupAvatarHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateGroupAvatarHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateGroupAvatarRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://***.png
	GroupAvatar *string `json:"groupAvatar,omitempty" xml:"groupAvatar,omitempty"`
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s UpdateGroupAvatarRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupAvatarRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupAvatarRequest) SetGroupAvatar(v string) *UpdateGroupAvatarRequest {
	s.GroupAvatar = &v
	return s
}

func (s *UpdateGroupAvatarRequest) SetOpenConversationId(v string) *UpdateGroupAvatarRequest {
	s.OpenConversationId = &v
	return s
}

type UpdateGroupAvatarResponseBody struct {
	// This parameter is required.
	NewGroupAvatar *string `json:"newGroupAvatar,omitempty" xml:"newGroupAvatar,omitempty"`
}

func (s UpdateGroupAvatarResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupAvatarResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupAvatarResponseBody) SetNewGroupAvatar(v string) *UpdateGroupAvatarResponseBody {
	s.NewGroupAvatar = &v
	return s
}

type UpdateGroupAvatarResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGroupAvatarResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGroupAvatarResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupAvatarResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupAvatarResponse) SetHeaders(v map[string]*string) *UpdateGroupAvatarResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupAvatarResponse) SetStatusCode(v int32) *UpdateGroupAvatarResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupAvatarResponse) SetBody(v *UpdateGroupAvatarResponseBody) *UpdateGroupAvatarResponse {
	s.Body = v
	return s
}

type UpdateGroupNameHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateGroupNameHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupNameHeaders) GoString() string {
	return s.String()
}

func (s *UpdateGroupNameHeaders) SetCommonHeaders(v map[string]*string) *UpdateGroupNameHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateGroupNameHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateGroupNameHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateGroupNameRequest struct {
	// This parameter is required.
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s UpdateGroupNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupNameRequest) SetGroupName(v string) *UpdateGroupNameRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateGroupNameRequest) SetOpenConversationId(v string) *UpdateGroupNameRequest {
	s.OpenConversationId = &v
	return s
}

type UpdateGroupNameResponseBody struct {
	NewGroupName *string `json:"newGroupName,omitempty" xml:"newGroupName,omitempty"`
}

func (s UpdateGroupNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupNameResponseBody) SetNewGroupName(v string) *UpdateGroupNameResponseBody {
	s.NewGroupName = &v
	return s
}

type UpdateGroupNameResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGroupNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGroupNameResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupNameResponse) SetHeaders(v map[string]*string) *UpdateGroupNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupNameResponse) SetStatusCode(v int32) *UpdateGroupNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupNameResponse) SetBody(v *UpdateGroupNameResponseBody) *UpdateGroupNameResponse {
	s.Body = v
	return s
}

type UpdateGroupPermissionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateGroupPermissionHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupPermissionHeaders) GoString() string {
	return s.String()
}

func (s *UpdateGroupPermissionHeaders) SetCommonHeaders(v map[string]*string) *UpdateGroupPermissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateGroupPermissionHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateGroupPermissionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateGroupPermissionRequest struct {
	// example:
	//
	// cidXXXXXXX
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	PermissionGroup    *string `json:"permissionGroup,omitempty" xml:"permissionGroup,omitempty"`
	Status             *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateGroupPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupPermissionRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupPermissionRequest) SetOpenConversationId(v string) *UpdateGroupPermissionRequest {
	s.OpenConversationId = &v
	return s
}

func (s *UpdateGroupPermissionRequest) SetPermissionGroup(v string) *UpdateGroupPermissionRequest {
	s.PermissionGroup = &v
	return s
}

func (s *UpdateGroupPermissionRequest) SetStatus(v string) *UpdateGroupPermissionRequest {
	s.Status = &v
	return s
}

type UpdateGroupPermissionResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateGroupPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupPermissionResponseBody) SetSuccess(v bool) *UpdateGroupPermissionResponseBody {
	s.Success = &v
	return s
}

type UpdateGroupPermissionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGroupPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGroupPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupPermissionResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupPermissionResponse) SetHeaders(v map[string]*string) *UpdateGroupPermissionResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupPermissionResponse) SetStatusCode(v int32) *UpdateGroupPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupPermissionResponse) SetBody(v *UpdateGroupPermissionResponseBody) *UpdateGroupPermissionResponse {
	s.Body = v
	return s
}

type UpdateGroupSubAdminHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateGroupSubAdminHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupSubAdminHeaders) GoString() string {
	return s.String()
}

func (s *UpdateGroupSubAdminHeaders) SetCommonHeaders(v map[string]*string) *UpdateGroupSubAdminHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateGroupSubAdminHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateGroupSubAdminHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateGroupSubAdminRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cidXXXXXXX
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	Role *int64 `json:"role,omitempty" xml:"role,omitempty"`
	// This parameter is required.
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s UpdateGroupSubAdminRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupSubAdminRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupSubAdminRequest) SetOpenConversationId(v string) *UpdateGroupSubAdminRequest {
	s.OpenConversationId = &v
	return s
}

func (s *UpdateGroupSubAdminRequest) SetRole(v int64) *UpdateGroupSubAdminRequest {
	s.Role = &v
	return s
}

func (s *UpdateGroupSubAdminRequest) SetUserIds(v []*string) *UpdateGroupSubAdminRequest {
	s.UserIds = v
	return s
}

type UpdateGroupSubAdminResponseBody struct {
	// This parameter is required.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateGroupSubAdminResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupSubAdminResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupSubAdminResponseBody) SetSuccess(v bool) *UpdateGroupSubAdminResponseBody {
	s.Success = &v
	return s
}

type UpdateGroupSubAdminResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGroupSubAdminResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGroupSubAdminResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupSubAdminResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupSubAdminResponse) SetHeaders(v map[string]*string) *UpdateGroupSubAdminResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupSubAdminResponse) SetStatusCode(v int32) *UpdateGroupSubAdminResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupSubAdminResponse) SetBody(v *UpdateGroupSubAdminResponseBody) *UpdateGroupSubAdminResponse {
	s.Body = v
	return s
}

type UpdateInteractiveCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateInteractiveCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateInteractiveCardHeaders) GoString() string {
	return s.String()
}

func (s *UpdateInteractiveCardHeaders) SetCommonHeaders(v map[string]*string) *UpdateInteractiveCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateInteractiveCardHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateInteractiveCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateInteractiveCardRequest struct {
	CardData    *UpdateInteractiveCardRequestCardData    `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	CardOptions *UpdateInteractiveCardRequestCardOptions `json:"cardOptions,omitempty" xml:"cardOptions,omitempty" type:"Struct"`
	OutTrackId  *string                                  `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	PrivateData map[string]*PrivateDataValue             `json:"privateData,omitempty" xml:"privateData,omitempty"`
	UserIdType  *int32                                   `json:"userIdType,omitempty" xml:"userIdType,omitempty"`
}

func (s UpdateInteractiveCardRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInteractiveCardRequest) GoString() string {
	return s.String()
}

func (s *UpdateInteractiveCardRequest) SetCardData(v *UpdateInteractiveCardRequestCardData) *UpdateInteractiveCardRequest {
	s.CardData = v
	return s
}

func (s *UpdateInteractiveCardRequest) SetCardOptions(v *UpdateInteractiveCardRequestCardOptions) *UpdateInteractiveCardRequest {
	s.CardOptions = v
	return s
}

func (s *UpdateInteractiveCardRequest) SetOutTrackId(v string) *UpdateInteractiveCardRequest {
	s.OutTrackId = &v
	return s
}

func (s *UpdateInteractiveCardRequest) SetPrivateData(v map[string]*PrivateDataValue) *UpdateInteractiveCardRequest {
	s.PrivateData = v
	return s
}

func (s *UpdateInteractiveCardRequest) SetUserIdType(v int32) *UpdateInteractiveCardRequest {
	s.UserIdType = &v
	return s
}

type UpdateInteractiveCardRequestCardData struct {
	CardMediaIdParamMap map[string]*string `json:"cardMediaIdParamMap,omitempty" xml:"cardMediaIdParamMap,omitempty"`
	CardParamMap        map[string]*string `json:"cardParamMap,omitempty" xml:"cardParamMap,omitempty"`
}

func (s UpdateInteractiveCardRequestCardData) String() string {
	return tea.Prettify(s)
}

func (s UpdateInteractiveCardRequestCardData) GoString() string {
	return s.String()
}

func (s *UpdateInteractiveCardRequestCardData) SetCardMediaIdParamMap(v map[string]*string) *UpdateInteractiveCardRequestCardData {
	s.CardMediaIdParamMap = v
	return s
}

func (s *UpdateInteractiveCardRequestCardData) SetCardParamMap(v map[string]*string) *UpdateInteractiveCardRequestCardData {
	s.CardParamMap = v
	return s
}

type UpdateInteractiveCardRequestCardOptions struct {
	UpdateCardDataByKey    *bool `json:"updateCardDataByKey,omitempty" xml:"updateCardDataByKey,omitempty"`
	UpdatePrivateDataByKey *bool `json:"updatePrivateDataByKey,omitempty" xml:"updatePrivateDataByKey,omitempty"`
}

func (s UpdateInteractiveCardRequestCardOptions) String() string {
	return tea.Prettify(s)
}

func (s UpdateInteractiveCardRequestCardOptions) GoString() string {
	return s.String()
}

func (s *UpdateInteractiveCardRequestCardOptions) SetUpdateCardDataByKey(v bool) *UpdateInteractiveCardRequestCardOptions {
	s.UpdateCardDataByKey = &v
	return s
}

func (s *UpdateInteractiveCardRequestCardOptions) SetUpdatePrivateDataByKey(v bool) *UpdateInteractiveCardRequestCardOptions {
	s.UpdatePrivateDataByKey = &v
	return s
}

type UpdateInteractiveCardResponseBody struct {
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateInteractiveCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInteractiveCardResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInteractiveCardResponseBody) SetSuccess(v string) *UpdateInteractiveCardResponseBody {
	s.Success = &v
	return s
}

type UpdateInteractiveCardResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInteractiveCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInteractiveCardResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInteractiveCardResponse) GoString() string {
	return s.String()
}

func (s *UpdateInteractiveCardResponse) SetHeaders(v map[string]*string) *UpdateInteractiveCardResponse {
	s.Headers = v
	return s
}

func (s *UpdateInteractiveCardResponse) SetStatusCode(v int32) *UpdateInteractiveCardResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInteractiveCardResponse) SetBody(v *UpdateInteractiveCardResponseBody) *UpdateInteractiveCardResponse {
	s.Body = v
	return s
}

type UpdateMemberBanWordsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateMemberBanWordsHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemberBanWordsHeaders) GoString() string {
	return s.String()
}

func (s *UpdateMemberBanWordsHeaders) SetCommonHeaders(v map[string]*string) *UpdateMemberBanWordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateMemberBanWordsHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateMemberBanWordsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateMemberBanWordsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 300000
	MuteDuration *int64 `json:"muteDuration,omitempty" xml:"muteDuration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	MuteStatus *int32 `json:"muteStatus,omitempty" xml:"muteStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cid5d5uM3XEw3gxbNc/n7EQ4g==
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s UpdateMemberBanWordsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemberBanWordsRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemberBanWordsRequest) SetMuteDuration(v int64) *UpdateMemberBanWordsRequest {
	s.MuteDuration = &v
	return s
}

func (s *UpdateMemberBanWordsRequest) SetMuteStatus(v int32) *UpdateMemberBanWordsRequest {
	s.MuteStatus = &v
	return s
}

func (s *UpdateMemberBanWordsRequest) SetOpenConversationId(v string) *UpdateMemberBanWordsRequest {
	s.OpenConversationId = &v
	return s
}

func (s *UpdateMemberBanWordsRequest) SetUserIdList(v []*string) *UpdateMemberBanWordsRequest {
	s.UserIdList = v
	return s
}

type UpdateMemberBanWordsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateMemberBanWordsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemberBanWordsResponse) GoString() string {
	return s.String()
}

func (s *UpdateMemberBanWordsResponse) SetHeaders(v map[string]*string) *UpdateMemberBanWordsResponse {
	s.Headers = v
	return s
}

func (s *UpdateMemberBanWordsResponse) SetStatusCode(v int32) *UpdateMemberBanWordsResponse {
	s.StatusCode = &v
	return s
}

type UpdateMemberGroupNickHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateMemberGroupNickHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemberGroupNickHeaders) GoString() string {
	return s.String()
}

func (s *UpdateMemberGroupNickHeaders) SetCommonHeaders(v map[string]*string) *UpdateMemberGroupNickHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateMemberGroupNickHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateMemberGroupNickHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateMemberGroupNickRequest struct {
	// This parameter is required.
	GroupNick *string `json:"groupNick,omitempty" xml:"groupNick,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cidXXXXXXX
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateMemberGroupNickRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemberGroupNickRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemberGroupNickRequest) SetGroupNick(v string) *UpdateMemberGroupNickRequest {
	s.GroupNick = &v
	return s
}

func (s *UpdateMemberGroupNickRequest) SetOpenConversationId(v string) *UpdateMemberGroupNickRequest {
	s.OpenConversationId = &v
	return s
}

func (s *UpdateMemberGroupNickRequest) SetUserId(v string) *UpdateMemberGroupNickRequest {
	s.UserId = &v
	return s
}

type UpdateMemberGroupNickResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateMemberGroupNickResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemberGroupNickResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMemberGroupNickResponseBody) SetSuccess(v bool) *UpdateMemberGroupNickResponseBody {
	s.Success = &v
	return s
}

type UpdateMemberGroupNickResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMemberGroupNickResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMemberGroupNickResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemberGroupNickResponse) GoString() string {
	return s.String()
}

func (s *UpdateMemberGroupNickResponse) SetHeaders(v map[string]*string) *UpdateMemberGroupNickResponse {
	s.Headers = v
	return s
}

func (s *UpdateMemberGroupNickResponse) SetStatusCode(v int32) *UpdateMemberGroupNickResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMemberGroupNickResponse) SetBody(v *UpdateMemberGroupNickResponseBody) *UpdateMemberGroupNickResponse {
	s.Body = v
	return s
}

type UpdateRobotInOrgHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateRobotInOrgHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateRobotInOrgHeaders) GoString() string {
	return s.String()
}

func (s *UpdateRobotInOrgHeaders) SetCommonHeaders(v map[string]*string) *UpdateRobotInOrgHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateRobotInOrgHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateRobotInOrgHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateRobotInOrgRequest struct {
	// example:
	//
	// 小加
	Brief *string `json:"brief,omitempty" xml:"brief,omitempty"`
	// example:
	//
	// 小加
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// @lALPDe7s26Bre
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// 小加
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 123
	OutgoingToken *string `json:"outgoingToken,omitempty" xml:"outgoingToken,omitempty"`
	// example:
	//
	// https://*.com
	OutgoingUrl *string `json:"outgoingUrl,omitempty" xml:"outgoingUrl,omitempty"`
	// example:
	//
	// @lALPDe7s26Bre
	PreviewMediaId *string `json:"previewMediaId,omitempty" xml:"previewMediaId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s UpdateRobotInOrgRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRobotInOrgRequest) GoString() string {
	return s.String()
}

func (s *UpdateRobotInOrgRequest) SetBrief(v string) *UpdateRobotInOrgRequest {
	s.Brief = &v
	return s
}

func (s *UpdateRobotInOrgRequest) SetDescription(v string) *UpdateRobotInOrgRequest {
	s.Description = &v
	return s
}

func (s *UpdateRobotInOrgRequest) SetIcon(v string) *UpdateRobotInOrgRequest {
	s.Icon = &v
	return s
}

func (s *UpdateRobotInOrgRequest) SetName(v string) *UpdateRobotInOrgRequest {
	s.Name = &v
	return s
}

func (s *UpdateRobotInOrgRequest) SetOutgoingToken(v string) *UpdateRobotInOrgRequest {
	s.OutgoingToken = &v
	return s
}

func (s *UpdateRobotInOrgRequest) SetOutgoingUrl(v string) *UpdateRobotInOrgRequest {
	s.OutgoingUrl = &v
	return s
}

func (s *UpdateRobotInOrgRequest) SetPreviewMediaId(v string) *UpdateRobotInOrgRequest {
	s.PreviewMediaId = &v
	return s
}

func (s *UpdateRobotInOrgRequest) SetRobotCode(v string) *UpdateRobotInOrgRequest {
	s.RobotCode = &v
	return s
}

type UpdateRobotInOrgResponseBody struct {
	RobotCode *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
}

func (s UpdateRobotInOrgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRobotInOrgResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRobotInOrgResponseBody) SetRobotCode(v string) *UpdateRobotInOrgResponseBody {
	s.RobotCode = &v
	return s
}

type UpdateRobotInOrgResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRobotInOrgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRobotInOrgResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRobotInOrgResponse) GoString() string {
	return s.String()
}

func (s *UpdateRobotInOrgResponse) SetHeaders(v map[string]*string) *UpdateRobotInOrgResponse {
	s.Headers = v
	return s
}

func (s *UpdateRobotInOrgResponse) SetStatusCode(v int32) *UpdateRobotInOrgResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRobotInOrgResponse) SetBody(v *UpdateRobotInOrgResponseBody) *UpdateRobotInOrgResponse {
	s.Body = v
	return s
}

type UpdateRobotInteractiveCardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateRobotInteractiveCardHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateRobotInteractiveCardHeaders) GoString() string {
	return s.String()
}

func (s *UpdateRobotInteractiveCardHeaders) SetCommonHeaders(v map[string]*string) *UpdateRobotInteractiveCardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateRobotInteractiveCardHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateRobotInteractiveCardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateRobotInteractiveCardRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cardXXXX01
	CardBizId *string `json:"cardBizId,omitempty" xml:"cardBizId,omitempty"`
	// example:
	//
	// 根据具体的cardTemplateId参考文档格式
	CardData              *string                                         `json:"cardData,omitempty" xml:"cardData,omitempty"`
	UnionIdPrivateDataMap *string                                         `json:"unionIdPrivateDataMap,omitempty" xml:"unionIdPrivateDataMap,omitempty"`
	UpdateOptions         *UpdateRobotInteractiveCardRequestUpdateOptions `json:"updateOptions,omitempty" xml:"updateOptions,omitempty" type:"Struct"`
	UserIdPrivateDataMap  *string                                         `json:"userIdPrivateDataMap,omitempty" xml:"userIdPrivateDataMap,omitempty"`
}

func (s UpdateRobotInteractiveCardRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRobotInteractiveCardRequest) GoString() string {
	return s.String()
}

func (s *UpdateRobotInteractiveCardRequest) SetCardBizId(v string) *UpdateRobotInteractiveCardRequest {
	s.CardBizId = &v
	return s
}

func (s *UpdateRobotInteractiveCardRequest) SetCardData(v string) *UpdateRobotInteractiveCardRequest {
	s.CardData = &v
	return s
}

func (s *UpdateRobotInteractiveCardRequest) SetUnionIdPrivateDataMap(v string) *UpdateRobotInteractiveCardRequest {
	s.UnionIdPrivateDataMap = &v
	return s
}

func (s *UpdateRobotInteractiveCardRequest) SetUpdateOptions(v *UpdateRobotInteractiveCardRequestUpdateOptions) *UpdateRobotInteractiveCardRequest {
	s.UpdateOptions = v
	return s
}

func (s *UpdateRobotInteractiveCardRequest) SetUserIdPrivateDataMap(v string) *UpdateRobotInteractiveCardRequest {
	s.UserIdPrivateDataMap = &v
	return s
}

type UpdateRobotInteractiveCardRequestUpdateOptions struct {
	UpdateCardDataByKey    *bool `json:"updateCardDataByKey,omitempty" xml:"updateCardDataByKey,omitempty"`
	UpdatePrivateDataByKey *bool `json:"updatePrivateDataByKey,omitempty" xml:"updatePrivateDataByKey,omitempty"`
}

func (s UpdateRobotInteractiveCardRequestUpdateOptions) String() string {
	return tea.Prettify(s)
}

func (s UpdateRobotInteractiveCardRequestUpdateOptions) GoString() string {
	return s.String()
}

func (s *UpdateRobotInteractiveCardRequestUpdateOptions) SetUpdateCardDataByKey(v bool) *UpdateRobotInteractiveCardRequestUpdateOptions {
	s.UpdateCardDataByKey = &v
	return s
}

func (s *UpdateRobotInteractiveCardRequestUpdateOptions) SetUpdatePrivateDataByKey(v bool) *UpdateRobotInteractiveCardRequestUpdateOptions {
	s.UpdatePrivateDataByKey = &v
	return s
}

type UpdateRobotInteractiveCardResponseBody struct {
	// example:
	//
	// xxxxxx
	ProcessQueryKey *string `json:"processQueryKey,omitempty" xml:"processQueryKey,omitempty"`
}

func (s UpdateRobotInteractiveCardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRobotInteractiveCardResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRobotInteractiveCardResponseBody) SetProcessQueryKey(v string) *UpdateRobotInteractiveCardResponseBody {
	s.ProcessQueryKey = &v
	return s
}

type UpdateRobotInteractiveCardResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRobotInteractiveCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRobotInteractiveCardResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRobotInteractiveCardResponse) GoString() string {
	return s.String()
}

func (s *UpdateRobotInteractiveCardResponse) SetHeaders(v map[string]*string) *UpdateRobotInteractiveCardResponse {
	s.Headers = v
	return s
}

func (s *UpdateRobotInteractiveCardResponse) SetStatusCode(v int32) *UpdateRobotInteractiveCardResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRobotInteractiveCardResponse) SetBody(v *UpdateRobotInteractiveCardResponseBody) *UpdateRobotInteractiveCardResponse {
	s.Body = v
	return s
}

type UpdateSceneGroupTemplateMessageOpenStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateSceneGroupTemplateMessageOpenStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateSceneGroupTemplateMessageOpenStatusHeaders) GoString() string {
	return s.String()
}

func (s *UpdateSceneGroupTemplateMessageOpenStatusHeaders) SetCommonHeaders(v map[string]*string) *UpdateSceneGroupTemplateMessageOpenStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateSceneGroupTemplateMessageOpenStatusHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateSceneGroupTemplateMessageOpenStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateSceneGroupTemplateMessageOpenStatusRequest struct {
	// This parameter is required.
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	TemplateIdList []*string `json:"templateIdList,omitempty" xml:"templateIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateSceneGroupTemplateMessageOpenStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSceneGroupTemplateMessageOpenStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateSceneGroupTemplateMessageOpenStatusRequest) SetStatus(v int32) *UpdateSceneGroupTemplateMessageOpenStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateSceneGroupTemplateMessageOpenStatusRequest) SetTemplateIdList(v []*string) *UpdateSceneGroupTemplateMessageOpenStatusRequest {
	s.TemplateIdList = v
	return s
}

func (s *UpdateSceneGroupTemplateMessageOpenStatusRequest) SetUserId(v string) *UpdateSceneGroupTemplateMessageOpenStatusRequest {
	s.UserId = &v
	return s
}

type UpdateSceneGroupTemplateMessageOpenStatusResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateSceneGroupTemplateMessageOpenStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSceneGroupTemplateMessageOpenStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSceneGroupTemplateMessageOpenStatusResponseBody) SetSuccess(v bool) *UpdateSceneGroupTemplateMessageOpenStatusResponseBody {
	s.Success = &v
	return s
}

type UpdateSceneGroupTemplateMessageOpenStatusResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSceneGroupTemplateMessageOpenStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSceneGroupTemplateMessageOpenStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSceneGroupTemplateMessageOpenStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateSceneGroupTemplateMessageOpenStatusResponse) SetHeaders(v map[string]*string) *UpdateSceneGroupTemplateMessageOpenStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateSceneGroupTemplateMessageOpenStatusResponse) SetStatusCode(v int32) *UpdateSceneGroupTemplateMessageOpenStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSceneGroupTemplateMessageOpenStatusResponse) SetBody(v *UpdateSceneGroupTemplateMessageOpenStatusResponseBody) *UpdateSceneGroupTemplateMessageOpenStatusResponse {
	s.Body = v
	return s
}

type UpdateTheGroupRolesOfGroupMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateTheGroupRolesOfGroupMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateTheGroupRolesOfGroupMemberHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTheGroupRolesOfGroupMemberHeaders) SetCommonHeaders(v map[string]*string) *UpdateTheGroupRolesOfGroupMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTheGroupRolesOfGroupMemberHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateTheGroupRolesOfGroupMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateTheGroupRolesOfGroupMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cidXXXXXXX
	OpenConversationId *string   `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	OpenRoleIds        []*string `json:"openRoleIds,omitempty" xml:"openRoleIds,omitempty" type:"Repeated"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateTheGroupRolesOfGroupMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTheGroupRolesOfGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *UpdateTheGroupRolesOfGroupMemberRequest) SetOpenConversationId(v string) *UpdateTheGroupRolesOfGroupMemberRequest {
	s.OpenConversationId = &v
	return s
}

func (s *UpdateTheGroupRolesOfGroupMemberRequest) SetOpenRoleIds(v []*string) *UpdateTheGroupRolesOfGroupMemberRequest {
	s.OpenRoleIds = v
	return s
}

func (s *UpdateTheGroupRolesOfGroupMemberRequest) SetUserId(v string) *UpdateTheGroupRolesOfGroupMemberRequest {
	s.UserId = &v
	return s
}

type UpdateTheGroupRolesOfGroupMemberResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateTheGroupRolesOfGroupMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTheGroupRolesOfGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTheGroupRolesOfGroupMemberResponseBody) SetSuccess(v bool) *UpdateTheGroupRolesOfGroupMemberResponseBody {
	s.Success = &v
	return s
}

type UpdateTheGroupRolesOfGroupMemberResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTheGroupRolesOfGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTheGroupRolesOfGroupMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTheGroupRolesOfGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *UpdateTheGroupRolesOfGroupMemberResponse) SetHeaders(v map[string]*string) *UpdateTheGroupRolesOfGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *UpdateTheGroupRolesOfGroupMemberResponse) SetStatusCode(v int32) *UpdateTheGroupRolesOfGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTheGroupRolesOfGroupMemberResponse) SetBody(v *UpdateTheGroupRolesOfGroupMemberResponseBody) *UpdateTheGroupRolesOfGroupMemberResponse {
	s.Body = v
	return s
}

type UpdateUnfurlingRegisterHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateUnfurlingRegisterHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateUnfurlingRegisterHeaders) GoString() string {
	return s.String()
}

func (s *UpdateUnfurlingRegisterHeaders) SetCommonHeaders(v map[string]*string) *UpdateUnfurlingRegisterHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateUnfurlingRegisterHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateUnfurlingRegisterHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateUnfurlingRegisterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123xxxx
	ApiSecret *string `json:"apiSecret,omitempty" xml:"apiSecret,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3102xxxxxxx
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://xxx.xxx.com/api/dingtalk/link_unfurling
	CallbackUrl *string `json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// d7b9xxx-xxx-xxxx-xxxx-xxxxxxx.schema
	CardTemplateId *string `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// www.dingtalk.com
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /a
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// 规则描述
	RuleDesc *string `json:"ruleDesc,omitempty" xml:"ruleDesc,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	RuleMatchType *int32 `json:"ruleMatchType,omitempty" xml:"ruleMatchType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 37xxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateUnfurlingRegisterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUnfurlingRegisterRequest) GoString() string {
	return s.String()
}

func (s *UpdateUnfurlingRegisterRequest) SetApiSecret(v string) *UpdateUnfurlingRegisterRequest {
	s.ApiSecret = &v
	return s
}

func (s *UpdateUnfurlingRegisterRequest) SetAppId(v string) *UpdateUnfurlingRegisterRequest {
	s.AppId = &v
	return s
}

func (s *UpdateUnfurlingRegisterRequest) SetCallbackUrl(v string) *UpdateUnfurlingRegisterRequest {
	s.CallbackUrl = &v
	return s
}

func (s *UpdateUnfurlingRegisterRequest) SetCardTemplateId(v string) *UpdateUnfurlingRegisterRequest {
	s.CardTemplateId = &v
	return s
}

func (s *UpdateUnfurlingRegisterRequest) SetDomain(v string) *UpdateUnfurlingRegisterRequest {
	s.Domain = &v
	return s
}

func (s *UpdateUnfurlingRegisterRequest) SetId(v int64) *UpdateUnfurlingRegisterRequest {
	s.Id = &v
	return s
}

func (s *UpdateUnfurlingRegisterRequest) SetPath(v string) *UpdateUnfurlingRegisterRequest {
	s.Path = &v
	return s
}

func (s *UpdateUnfurlingRegisterRequest) SetRuleDesc(v string) *UpdateUnfurlingRegisterRequest {
	s.RuleDesc = &v
	return s
}

func (s *UpdateUnfurlingRegisterRequest) SetRuleMatchType(v int32) *UpdateUnfurlingRegisterRequest {
	s.RuleMatchType = &v
	return s
}

func (s *UpdateUnfurlingRegisterRequest) SetUserId(v string) *UpdateUnfurlingRegisterRequest {
	s.UserId = &v
	return s
}

type UpdateUnfurlingRegisterResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateUnfurlingRegisterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUnfurlingRegisterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUnfurlingRegisterResponseBody) SetSuccess(v bool) *UpdateUnfurlingRegisterResponseBody {
	s.Success = &v
	return s
}

type UpdateUnfurlingRegisterResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUnfurlingRegisterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUnfurlingRegisterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUnfurlingRegisterResponse) GoString() string {
	return s.String()
}

func (s *UpdateUnfurlingRegisterResponse) SetHeaders(v map[string]*string) *UpdateUnfurlingRegisterResponse {
	s.Headers = v
	return s
}

func (s *UpdateUnfurlingRegisterResponse) SetStatusCode(v int32) *UpdateUnfurlingRegisterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUnfurlingRegisterResponse) SetBody(v *UpdateUnfurlingRegisterResponseBody) *UpdateUnfurlingRegisterResponse {
	s.Body = v
	return s
}

type UpdateUnfurlingRegisterStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateUnfurlingRegisterStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateUnfurlingRegisterStatusHeaders) GoString() string {
	return s.String()
}

func (s *UpdateUnfurlingRegisterStatusHeaders) SetCommonHeaders(v map[string]*string) *UpdateUnfurlingRegisterStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateUnfurlingRegisterStatusHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateUnfurlingRegisterStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateUnfurlingRegisterStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3102xxxxxxx
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 37xxxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateUnfurlingRegisterStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUnfurlingRegisterStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateUnfurlingRegisterStatusRequest) SetAppId(v string) *UpdateUnfurlingRegisterStatusRequest {
	s.AppId = &v
	return s
}

func (s *UpdateUnfurlingRegisterStatusRequest) SetId(v int64) *UpdateUnfurlingRegisterStatusRequest {
	s.Id = &v
	return s
}

func (s *UpdateUnfurlingRegisterStatusRequest) SetStatus(v int32) *UpdateUnfurlingRegisterStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateUnfurlingRegisterStatusRequest) SetUserId(v string) *UpdateUnfurlingRegisterStatusRequest {
	s.UserId = &v
	return s
}

type UpdateUnfurlingRegisterStatusResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateUnfurlingRegisterStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUnfurlingRegisterStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUnfurlingRegisterStatusResponseBody) SetSuccess(v bool) *UpdateUnfurlingRegisterStatusResponseBody {
	s.Success = &v
	return s
}

type UpdateUnfurlingRegisterStatusResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUnfurlingRegisterStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUnfurlingRegisterStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUnfurlingRegisterStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateUnfurlingRegisterStatusResponse) SetHeaders(v map[string]*string) *UpdateUnfurlingRegisterStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateUnfurlingRegisterStatusResponse) SetStatusCode(v int32) *UpdateUnfurlingRegisterStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUnfurlingRegisterStatusResponse) SetBody(v *UpdateUnfurlingRegisterStatusResponseBody) *UpdateUnfurlingRegisterStatusResponse {
	s.Body = v
	return s
}

type UpgradeToExternalGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpgradeToExternalGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpgradeToExternalGroupHeaders) GoString() string {
	return s.String()
}

func (s *UpgradeToExternalGroupHeaders) SetCommonHeaders(v map[string]*string) *UpgradeToExternalGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpgradeToExternalGroupHeaders) SetXAcsDingtalkAccessToken(v string) *UpgradeToExternalGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpgradeToExternalGroupRequest struct {
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s UpgradeToExternalGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeToExternalGroupRequest) GoString() string {
	return s.String()
}

func (s *UpgradeToExternalGroupRequest) SetOpenConversationId(v string) *UpgradeToExternalGroupRequest {
	s.OpenConversationId = &v
	return s
}

func (s *UpgradeToExternalGroupRequest) SetTemplateId(v string) *UpgradeToExternalGroupRequest {
	s.TemplateId = &v
	return s
}

type UpgradeToExternalGroupResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpgradeToExternalGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeToExternalGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeToExternalGroupResponseBody) SetSuccess(v bool) *UpgradeToExternalGroupResponseBody {
	s.Success = &v
	return s
}

type UpgradeToExternalGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeToExternalGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeToExternalGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeToExternalGroupResponse) GoString() string {
	return s.String()
}

func (s *UpgradeToExternalGroupResponse) SetHeaders(v map[string]*string) *UpgradeToExternalGroupResponse {
	s.Headers = v
	return s
}

func (s *UpgradeToExternalGroupResponse) SetStatusCode(v int32) *UpgradeToExternalGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeToExternalGroupResponse) SetBody(v *UpgradeToExternalGroupResponseBody) *UpgradeToExternalGroupResponse {
	s.Body = v
	return s
}

type UpgradeToServiceGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpgradeToServiceGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpgradeToServiceGroupHeaders) GoString() string {
	return s.String()
}

func (s *UpgradeToServiceGroupHeaders) SetCommonHeaders(v map[string]*string) *UpgradeToServiceGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpgradeToServiceGroupHeaders) SetXAcsDingtalkAccessToken(v string) *UpgradeToServiceGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpgradeToServiceGroupRequest struct {
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s UpgradeToServiceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeToServiceGroupRequest) GoString() string {
	return s.String()
}

func (s *UpgradeToServiceGroupRequest) SetOpenConversationId(v string) *UpgradeToServiceGroupRequest {
	s.OpenConversationId = &v
	return s
}

func (s *UpgradeToServiceGroupRequest) SetTemplateId(v string) *UpgradeToServiceGroupRequest {
	s.TemplateId = &v
	return s
}

type UpgradeToServiceGroupResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpgradeToServiceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeToServiceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeToServiceGroupResponseBody) SetSuccess(v bool) *UpgradeToServiceGroupResponseBody {
	s.Success = &v
	return s
}

type UpgradeToServiceGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeToServiceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeToServiceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeToServiceGroupResponse) GoString() string {
	return s.String()
}

func (s *UpgradeToServiceGroupResponse) SetHeaders(v map[string]*string) *UpgradeToServiceGroupResponse {
	s.Headers = v
	return s
}

func (s *UpgradeToServiceGroupResponse) SetStatusCode(v int32) *UpgradeToServiceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeToServiceGroupResponse) SetBody(v *UpgradeToServiceGroupResponseBody) *UpgradeToServiceGroupResponse {
	s.Body = v
	return s
}

type AddGroupMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddGroupMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMemberHeaders) GoString() string {
	return s.String()
}

func (s *AddGroupMemberHeaders) SetCommonHeaders(v map[string]*string) *AddGroupMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddGroupMemberHeaders) SetXAcsDingtalkAccessToken(v string) *AddGroupMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddGroupMemberRequest struct {
	AppUserIds []*string `json:"appUserIds,omitempty" xml:"appUserIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 14da****2760
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1745****8777
	OperatorId *string   `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	UserIds    []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s AddGroupMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *AddGroupMemberRequest) SetAppUserIds(v []*string) *AddGroupMemberRequest {
	s.AppUserIds = v
	return s
}

func (s *AddGroupMemberRequest) SetOpenConversationId(v string) *AddGroupMemberRequest {
	s.OpenConversationId = &v
	return s
}

func (s *AddGroupMemberRequest) SetOperatorId(v string) *AddGroupMemberRequest {
	s.OperatorId = &v
	return s
}

func (s *AddGroupMemberRequest) SetUserIds(v []*string) *AddGroupMemberRequest {
	s.UserIds = v
	return s
}

type AddGroupMemberResponseBody struct {
	// This parameter is required.
	AppUserIds []*string `json:"appUserIds,omitempty" xml:"appUserIds,omitempty" type:"Repeated"`
	// This parameter is required.
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s AddGroupMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddGroupMemberResponseBody) SetAppUserIds(v []*string) *AddGroupMemberResponseBody {
	s.AppUserIds = v
	return s
}

func (s *AddGroupMemberResponseBody) SetUserIds(v []*string) *AddGroupMemberResponseBody {
	s.UserIds = v
	return s
}

type AddGroupMemberResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddGroupMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *AddGroupMemberResponse) SetHeaders(v map[string]*string) *AddGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *AddGroupMemberResponse) SetStatusCode(v int32) *AddGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *AddGroupMemberResponse) SetBody(v *AddGroupMemberResponseBody) *AddGroupMemberResponse {
	s.Body = v
	return s
}

type RemoveGroupMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RemoveGroupMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupMemberHeaders) GoString() string {
	return s.String()
}

func (s *RemoveGroupMemberHeaders) SetCommonHeaders(v map[string]*string) *RemoveGroupMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveGroupMemberHeaders) SetXAcsDingtalkAccessToken(v string) *RemoveGroupMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RemoveGroupMemberRequest struct {
	AppUserIds []*string `json:"appUserIds,omitempty" xml:"appUserIds,omitempty" type:"Repeated"`
	// This parameter is required.
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1745****8777
	OperatorId *string   `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	UserIds    []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s RemoveGroupMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *RemoveGroupMemberRequest) SetAppUserIds(v []*string) *RemoveGroupMemberRequest {
	s.AppUserIds = v
	return s
}

func (s *RemoveGroupMemberRequest) SetOpenConversationId(v string) *RemoveGroupMemberRequest {
	s.OpenConversationId = &v
	return s
}

func (s *RemoveGroupMemberRequest) SetOperatorId(v string) *RemoveGroupMemberRequest {
	s.OperatorId = &v
	return s
}

func (s *RemoveGroupMemberRequest) SetUserIds(v []*string) *RemoveGroupMemberRequest {
	s.UserIds = v
	return s
}

type RemoveGroupMemberResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 移除成功
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s RemoveGroupMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveGroupMemberResponseBody) SetMessage(v string) *RemoveGroupMemberResponseBody {
	s.Message = &v
	return s
}

type RemoveGroupMemberResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveGroupMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *RemoveGroupMemberResponse) SetHeaders(v map[string]*string) *RemoveGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *RemoveGroupMemberResponse) SetStatusCode(v int32) *RemoveGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveGroupMemberResponse) SetBody(v *RemoveGroupMemberResponseBody) *RemoveGroupMemberResponse {
	s.Body = v
	return s
}

type SendDingMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendDingMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendDingMessageHeaders) GoString() string {
	return s.String()
}

func (s *SendDingMessageHeaders) SetCommonHeaders(v map[string]*string) *SendDingMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendDingMessageHeaders) SetXAcsDingtalkAccessToken(v string) *SendDingMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendDingMessageRequest struct {
	// This parameter is required.
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"msg_type":"text","text":"hello world"}
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// text
	MessageType        *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// 1107****2120
	ReceiverId *string `json:"receiverId,omitempty" xml:"receiverId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1745****8777
	SenderId *string `json:"senderId,omitempty" xml:"senderId,omitempty"`
}

func (s SendDingMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendDingMessageRequest) GoString() string {
	return s.String()
}

func (s *SendDingMessageRequest) SetCode(v string) *SendDingMessageRequest {
	s.Code = &v
	return s
}

func (s *SendDingMessageRequest) SetMessage(v string) *SendDingMessageRequest {
	s.Message = &v
	return s
}

func (s *SendDingMessageRequest) SetMessageType(v string) *SendDingMessageRequest {
	s.MessageType = &v
	return s
}

func (s *SendDingMessageRequest) SetOpenConversationId(v string) *SendDingMessageRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SendDingMessageRequest) SetReceiverId(v string) *SendDingMessageRequest {
	s.ReceiverId = &v
	return s
}

func (s *SendDingMessageRequest) SetSenderId(v string) *SendDingMessageRequest {
	s.SenderId = &v
	return s
}

type SendDingMessageResponseBody struct {
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SendDingMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendDingMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendDingMessageResponseBody) SetRequestId(v string) *SendDingMessageResponseBody {
	s.RequestId = &v
	return s
}

type SendDingMessageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendDingMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendDingMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendDingMessageResponse) GoString() string {
	return s.String()
}

func (s *SendDingMessageResponse) SetHeaders(v map[string]*string) *SendDingMessageResponse {
	s.Headers = v
	return s
}

func (s *SendDingMessageResponse) SetStatusCode(v int32) *SendDingMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendDingMessageResponse) SetBody(v *SendDingMessageResponseBody) *SendDingMessageResponse {
	s.Body = v
	return s
}

type SendMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendMessageHeaders) GoString() string {
	return s.String()
}

func (s *SendMessageHeaders) SetCommonHeaders(v map[string]*string) *SendMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendMessageHeaders) SetXAcsDingtalkAccessToken(v string) *SendMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendMessageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {"msg_type":"text","text":"hello world"}
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// text
	MessageType        *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// 1745****8777
	ReceiverId *string `json:"receiverId,omitempty" xml:"receiverId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1107****2120
	SenderId *string `json:"senderId,omitempty" xml:"senderId,omitempty"`
	// example:
	//
	// {     "9d801647a64******59c9da0207":"[{\"action_url\":\"http://www.baidu.com\",\"title\":\"一个按钮\"},{\"action_url\":\"http://www.baidu.com\",\"title\":\"两个按钮\"}]",     "9d801647a6******59c9da020342":"[{\"action_url\":\"http://www.baidu.com\",\"title\":\"一个按钮\"},{\"action_url\":\"http://www.baidu.com\",\"title\":\"两个按钮\"}]" }
	SourceInfos map[string]interface{} `json:"sourceInfos,omitempty" xml:"sourceInfos,omitempty"`
}

func (s SendMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageRequest) GoString() string {
	return s.String()
}

func (s *SendMessageRequest) SetMessage(v string) *SendMessageRequest {
	s.Message = &v
	return s
}

func (s *SendMessageRequest) SetMessageType(v string) *SendMessageRequest {
	s.MessageType = &v
	return s
}

func (s *SendMessageRequest) SetOpenConversationId(v string) *SendMessageRequest {
	s.OpenConversationId = &v
	return s
}

func (s *SendMessageRequest) SetReceiverId(v string) *SendMessageRequest {
	s.ReceiverId = &v
	return s
}

func (s *SendMessageRequest) SetSenderId(v string) *SendMessageRequest {
	s.SenderId = &v
	return s
}

func (s *SendMessageRequest) SetSourceInfos(v map[string]interface{}) *SendMessageRequest {
	s.SourceInfos = v
	return s
}

type SendMessageResponseBody struct {
	// This parameter is required.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SendMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageResponseBody) SetRequestId(v string) *SendMessageResponseBody {
	s.RequestId = &v
	return s
}

type SendMessageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponse) GoString() string {
	return s.String()
}

func (s *SendMessageResponse) SetHeaders(v map[string]*string) *SendMessageResponse {
	s.Headers = v
	return s
}

func (s *SendMessageResponse) SetStatusCode(v int32) *SendMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendMessageResponse) SetBody(v *SendMessageResponseBody) *SendMessageResponse {
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
// 添加企业文字表情
//
// @param request - AddOrgTextEmotionRequest
//
// @param headers - AddOrgTextEmotionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddOrgTextEmotionResponse
func (client *Client) AddOrgTextEmotionWithOptions(request *AddOrgTextEmotionRequest, headers *AddOrgTextEmotionHeaders, runtime *util.RuntimeOptions) (_result *AddOrgTextEmotionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackgroundMediaId)) {
		body["backgroundMediaId"] = request.BackgroundMediaId
	}

	if !tea.BoolValue(util.IsUnset(request.BackgroundMediaIdForPanel)) {
		body["backgroundMediaIdForPanel"] = request.BackgroundMediaIdForPanel
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.EmotionName)) {
		body["emotionName"] = request.EmotionName
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
		Action:      tea.String("AddOrgTextEmotion"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/organizations/textEmotions"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddOrgTextEmotionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加企业文字表情
//
// @param request - AddOrgTextEmotionRequest
//
// @return AddOrgTextEmotionResponse
func (client *Client) AddOrgTextEmotion(request *AddOrgTextEmotionRequest) (_result *AddOrgTextEmotionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddOrgTextEmotionHeaders{}
	_result = &AddOrgTextEmotionResponse{}
	_body, _err := client.AddOrgTextEmotionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加机器人到会话
//
// @param request - AddRobotToConversationRequest
//
// @param headers - AddRobotToConversationHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRobotToConversationResponse
func (client *Client) AddRobotToConversationWithOptions(request *AddRobotToConversationRequest, headers *AddRobotToConversationHeaders, runtime *util.RuntimeOptions) (_result *AddRobotToConversationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		body["icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
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
		Action:      tea.String("AddRobotToConversation"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/conversations/robots"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddRobotToConversationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加机器人到会话
//
// @param request - AddRobotToConversationRequest
//
// @return AddRobotToConversationResponse
func (client *Client) AddRobotToConversation(request *AddRobotToConversationRequest) (_result *AddRobotToConversationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddRobotToConversationHeaders{}
	_result = &AddRobotToConversationResponse{}
	_body, _err := client.AddRobotToConversationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增链接增强注册规则
//
// @param request - AddUnfurlingRegisterRequest
//
// @param headers - AddUnfurlingRegisterHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUnfurlingRegisterResponse
func (client *Client) AddUnfurlingRegisterWithOptions(request *AddUnfurlingRegisterRequest, headers *AddUnfurlingRegisterHeaders, runtime *util.RuntimeOptions) (_result *AddUnfurlingRegisterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiSecret)) {
		body["apiSecret"] = request.ApiSecret
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackUrl)) {
		body["callbackUrl"] = request.CallbackUrl
	}

	if !tea.BoolValue(util.IsUnset(request.CardTemplateId)) {
		body["cardTemplateId"] = request.CardTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		body["domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		body["path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.RuleDesc)) {
		body["ruleDesc"] = request.RuleDesc
	}

	if !tea.BoolValue(util.IsUnset(request.RuleMatchType)) {
		body["ruleMatchType"] = request.RuleMatchType
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
		Action:      tea.String("AddUnfurlingRegister"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/unfurling/rules"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddUnfurlingRegisterResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增链接增强注册规则
//
// @param request - AddUnfurlingRegisterRequest
//
// @return AddUnfurlingRegisterResponse
func (client *Client) AddUnfurlingRegister(request *AddUnfurlingRegisterRequest) (_result *AddUnfurlingRegisterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddUnfurlingRegisterHeaders{}
	_result = &AddUnfurlingRegisterResponse{}
	_body, _err := client.AddUnfurlingRegisterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 自动开通钉钉客联微应用
//
// @param headers - AutoOpenDingTalkConnectHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AutoOpenDingTalkConnectResponse
func (client *Client) AutoOpenDingTalkConnectWithOptions(headers *AutoOpenDingTalkConnectHeaders, runtime *util.RuntimeOptions) (_result *AutoOpenDingTalkConnectResponse, _err error) {
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
		Action:      tea.String("AutoOpenDingTalkConnect"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/interconnections/apps/open"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AutoOpenDingTalkConnectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 自动开通钉钉客联微应用
//
// @return AutoOpenDingTalkConnectResponse
func (client *Client) AutoOpenDingTalkConnect() (_result *AutoOpenDingTalkConnectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AutoOpenDingTalkConnectHeaders{}
	_result = &AutoOpenDingTalkConnectResponse{}
	_body, _err := client.AutoOpenDingTalkConnectWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询家校群消息详情
//
// @param request - BatchQueryFamilySchoolMessageRequest
//
// @param headers - BatchQueryFamilySchoolMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchQueryFamilySchoolMessageResponse
func (client *Client) BatchQueryFamilySchoolMessageWithOptions(request *BatchQueryFamilySchoolMessageRequest, headers *BatchQueryFamilySchoolMessageHeaders, runtime *util.RuntimeOptions) (_result *BatchQueryFamilySchoolMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenMessageIds)) {
		body["openMessageIds"] = request.OpenMessageIds
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("BatchQueryFamilySchoolMessage"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/conversations/familySchools/messages/batchQuery"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchQueryFamilySchoolMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询家校群消息详情
//
// @param request - BatchQueryFamilySchoolMessageRequest
//
// @return BatchQueryFamilySchoolMessageResponse
func (client *Client) BatchQueryFamilySchoolMessage(request *BatchQueryFamilySchoolMessageRequest) (_result *BatchQueryFamilySchoolMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchQueryFamilySchoolMessageHeaders{}
	_result = &BatchQueryFamilySchoolMessageResponse{}
	_body, _err := client.BatchQueryFamilySchoolMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询群成员
//
// @param request - BatchQueryGroupMemberRequest
//
// @param headers - BatchQueryGroupMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchQueryGroupMemberResponse
func (client *Client) BatchQueryGroupMemberWithOptions(request *BatchQueryGroupMemberRequest, headers *BatchQueryGroupMemberHeaders, runtime *util.RuntimeOptions) (_result *BatchQueryGroupMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CoolAppCode)) {
		body["coolAppCode"] = request.CoolAppCode
	}

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
		Action:      tea.String("BatchQueryGroupMember"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/sceneGroups/members/batchQuery"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchQueryGroupMemberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询群成员
//
// @param request - BatchQueryGroupMemberRequest
//
// @return BatchQueryGroupMemberResponse
func (client *Client) BatchQueryGroupMember(request *BatchQueryGroupMemberRequest) (_result *BatchQueryGroupMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchQueryGroupMemberHeaders{}
	_result = &BatchQueryGroupMemberResponse{}
	_body, _err := client.BatchQueryGroupMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 钉钉互动卡片模板构建动作
//
// @param request - CardTemplateBuildActionRequest
//
// @param headers - CardTemplateBuildActionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CardTemplateBuildActionResponse
func (client *Client) CardTemplateBuildActionWithOptions(request *CardTemplateBuildActionRequest, headers *CardTemplateBuildActionHeaders, runtime *util.RuntimeOptions) (_result *CardTemplateBuildActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		body["action"] = request.Action
	}

	if !tea.BoolValue(util.IsUnset(request.CardTemplateJson)) {
		body["cardTemplateJson"] = request.CardTemplateJson
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
		Action:      tea.String("CardTemplateBuildAction"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/interactiveCards/templates/buildAction"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CardTemplateBuildActionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 钉钉互动卡片模板构建动作
//
// @param request - CardTemplateBuildActionRequest
//
// @return CardTemplateBuildActionResponse
func (client *Client) CardTemplateBuildAction(request *CardTemplateBuildActionRequest) (_result *CardTemplateBuildActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CardTemplateBuildActionHeaders{}
	_result = &CardTemplateBuildActionResponse{}
	_body, _err := client.CardTemplateBuildActionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更换群主
//
// @param request - ChangeGroupOwnerRequest
//
// @param headers - ChangeGroupOwnerHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeGroupOwnerResponse
func (client *Client) ChangeGroupOwnerWithOptions(request *ChangeGroupOwnerRequest, headers *ChangeGroupOwnerHeaders, runtime *util.RuntimeOptions) (_result *ChangeGroupOwnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupOwnerId)) {
		body["groupOwnerId"] = request.GroupOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupOwnerType)) {
		body["groupOwnerType"] = request.GroupOwnerType
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
		Action:      tea.String("ChangeGroupOwner"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/interconnections/groups/owners"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeGroupOwnerResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更换群主
//
// @param request - ChangeGroupOwnerRequest
//
// @return ChangeGroupOwnerResponse
func (client *Client) ChangeGroupOwner(request *ChangeGroupOwnerRequest) (_result *ChangeGroupOwnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ChangeGroupOwnerHeaders{}
	_result = &ChangeGroupOwnerResponse{}
	_body, _err := client.ChangeGroupOwnerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 会话开放的ChatId转OpenConversationId
//
// @param headers - ChatIdToOpenConversationIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatIdToOpenConversationIdResponse
func (client *Client) ChatIdToOpenConversationIdWithOptions(chatId *string, headers *ChatIdToOpenConversationIdHeaders, runtime *util.RuntimeOptions) (_result *ChatIdToOpenConversationIdResponse, _err error) {
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
		Action:      tea.String("ChatIdToOpenConversationId"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/chat/" + tea.StringValue(chatId) + "/convertToOpenConversationId"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ChatIdToOpenConversationIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 会话开放的ChatId转OpenConversationId
//
// @return ChatIdToOpenConversationIdResponse
func (client *Client) ChatIdToOpenConversationId(chatId *string) (_result *ChatIdToOpenConversationIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ChatIdToOpenConversationIdHeaders{}
	_result = &ChatIdToOpenConversationIdResponse{}
	_body, _err := client.ChatIdToOpenConversationIdWithOptions(chatId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置群管理员
//
// @param request - ChatSubAdminUpdateRequest
//
// @param headers - ChatSubAdminUpdateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatSubAdminUpdateResponse
func (client *Client) ChatSubAdminUpdateWithOptions(request *ChatSubAdminUpdateRequest, headers *ChatSubAdminUpdateHeaders, runtime *util.RuntimeOptions) (_result *ChatSubAdminUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		body["role"] = request.Role
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
		Action:      tea.String("ChatSubAdminUpdate"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/subAdministrators"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ChatSubAdminUpdateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置群管理员
//
// @param request - ChatSubAdminUpdateRequest
//
// @return ChatSubAdminUpdateResponse
func (client *Client) ChatSubAdminUpdate(request *ChatSubAdminUpdateRequest) (_result *ChatSubAdminUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ChatSubAdminUpdateHeaders{}
	_result = &ChatSubAdminUpdateResponse{}
	_body, _err := client.ChatSubAdminUpdateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户是否为企业内部群成员
//
// @param request - CheckUserIsGroupMemberRequest
//
// @param headers - CheckUserIsGroupMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckUserIsGroupMemberResponse
func (client *Client) CheckUserIsGroupMemberWithOptions(request *CheckUserIsGroupMemberRequest, headers *CheckUserIsGroupMemberHeaders, runtime *util.RuntimeOptions) (_result *CheckUserIsGroupMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
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
		Action:      tea.String("CheckUserIsGroupMember"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/innerGroups/members/check"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckUserIsGroupMemberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户是否为企业内部群成员
//
// @param request - CheckUserIsGroupMemberRequest
//
// @return CheckUserIsGroupMemberResponse
func (client *Client) CheckUserIsGroupMember(request *CheckUserIsGroupMemberRequest) (_result *CheckUserIsGroupMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CheckUserIsGroupMemberHeaders{}
	_result = &CheckUserIsGroupMemberResponse{}
	_body, _err := client.CheckUserIsGroupMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 链接增强规则拷贝
//
// @param request - CopyUnfurlingRegisterRequest
//
// @param headers - CopyUnfurlingRegisterHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyUnfurlingRegisterResponse
func (client *Client) CopyUnfurlingRegisterWithOptions(request *CopyUnfurlingRegisterRequest, headers *CopyUnfurlingRegisterHeaders, runtime *util.RuntimeOptions) (_result *CopyUnfurlingRegisterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
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
		Action:      tea.String("CopyUnfurlingRegister"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/unfurling/rules/copy"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CopyUnfurlingRegisterResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 链接增强规则拷贝
//
// @param request - CopyUnfurlingRegisterRequest
//
// @return CopyUnfurlingRegisterResponse
func (client *Client) CopyUnfurlingRegister(request *CopyUnfurlingRegisterRequest) (_result *CopyUnfurlingRegisterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CopyUnfurlingRegisterHeaders{}
	_result = &CopyUnfurlingRegisterResponse{}
	_body, _err := client.CopyUnfurlingRegisterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询消息开放群模板下群计数
//
// @param request - CountOpenMsgSceneGroupsRequest
//
// @param headers - CountOpenMsgSceneGroupsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CountOpenMsgSceneGroupsResponse
func (client *Client) CountOpenMsgSceneGroupsWithOptions(request *CountOpenMsgSceneGroupsRequest, headers *CountOpenMsgSceneGroupsHeaders, runtime *util.RuntimeOptions) (_result *CountOpenMsgSceneGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["templateId"] = request.TemplateId
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
		Action:      tea.String("CountOpenMsgSceneGroups"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/openMsgSceneGroups/templates/counts/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CountOpenMsgSceneGroupsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询消息开放群模板下群计数
//
// @param request - CountOpenMsgSceneGroupsRequest
//
// @return CountOpenMsgSceneGroupsResponse
func (client *Client) CountOpenMsgSceneGroups(request *CountOpenMsgSceneGroupsRequest) (_result *CountOpenMsgSceneGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CountOpenMsgSceneGroupsHeaders{}
	_result = &CountOpenMsgSceneGroupsResponse{}
	_body, _err := client.CountOpenMsgSceneGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取企业下消息开放场景群数量
//
// @param headers - CountOrgMessageOpenSceneGroupsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CountOrgMessageOpenSceneGroupsResponse
func (client *Client) CountOrgMessageOpenSceneGroupsWithOptions(headers *CountOrgMessageOpenSceneGroupsHeaders, runtime *util.RuntimeOptions) (_result *CountOrgMessageOpenSceneGroupsResponse, _err error) {
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
		Action:      tea.String("CountOrgMessageOpenSceneGroups"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/chats/sceneGroups/counts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CountOrgMessageOpenSceneGroupsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取企业下消息开放场景群数量
//
// @return CountOrgMessageOpenSceneGroupsResponse
func (client *Client) CountOrgMessageOpenSceneGroups() (_result *CountOrgMessageOpenSceneGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CountOrgMessageOpenSceneGroupsHeaders{}
	_result = &CountOrgMessageOpenSceneGroupsResponse{}
	_body, _err := client.CountOrgMessageOpenSceneGroupsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询群模板关联的群数量
//
// @param headers - CountSceneGroupsByTemplateIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CountSceneGroupsByTemplateIdResponse
func (client *Client) CountSceneGroupsByTemplateIdWithOptions(templateId *string, headers *CountSceneGroupsByTemplateIdHeaders, runtime *util.RuntimeOptions) (_result *CountSceneGroupsByTemplateIdResponse, _err error) {
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
		Action:      tea.String("CountSceneGroupsByTemplateId"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/chats/sceneGroups/templates/" + tea.StringValue(templateId) + "/counts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CountSceneGroupsByTemplateIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询群模板关联的群数量
//
// @return CountSceneGroupsByTemplateIdResponse
func (client *Client) CountSceneGroupsByTemplateId(templateId *string) (_result *CountSceneGroupsByTemplateIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CountSceneGroupsByTemplateIdHeaders{}
	_result = &CountSceneGroupsByTemplateIdResponse{}
	_body, _err := client.CountSceneGroupsByTemplateIdWithOptions(templateId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建钉外两人群
//
// @param request - CreateCoupleGroupConversationRequest
//
// @param headers - CreateCoupleGroupConversationHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCoupleGroupConversationResponse
func (client *Client) CreateCoupleGroupConversationWithOptions(request *CreateCoupleGroupConversationRequest, headers *CreateCoupleGroupConversationHeaders, runtime *util.RuntimeOptions) (_result *CreateCoupleGroupConversationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUserId)) {
		body["appUserId"] = request.AppUserId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupAvatar)) {
		body["groupAvatar"] = request.GroupAvatar
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupOwnerId)) {
		body["groupOwnerId"] = request.GroupOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupTemplateId)) {
		body["groupTemplateId"] = request.GroupTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
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
		Action:      tea.String("CreateCoupleGroupConversation"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/interconnections/coupleGroups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCoupleGroupConversationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建钉外两人群
//
// @param request - CreateCoupleGroupConversationRequest
//
// @return CreateCoupleGroupConversationResponse
func (client *Client) CreateCoupleGroupConversation(request *CreateCoupleGroupConversationRequest) (_result *CreateCoupleGroupConversationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateCoupleGroupConversationHeaders{}
	_result = &CreateCoupleGroupConversationResponse{}
	_body, _err := client.CreateCoupleGroupConversationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建互通群（支持普通互通群、跨钉两人群）
//
// @param request - CreateGroupConversationRequest
//
// @param headers - CreateGroupConversationHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupConversationResponse
func (client *Client) CreateGroupConversationWithOptions(request *CreateGroupConversationRequest, headers *CreateGroupConversationHeaders, runtime *util.RuntimeOptions) (_result *CreateGroupConversationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUserIds)) {
		body["appUserIds"] = request.AppUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.GroupAvatar)) {
		body["groupAvatar"] = request.GroupAvatar
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupOwnerId)) {
		body["groupOwnerId"] = request.GroupOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupOwnerType)) {
		body["groupOwnerType"] = request.GroupOwnerType
	}

	if !tea.BoolValue(util.IsUnset(request.GroupTemplateId)) {
		body["groupTemplateId"] = request.GroupTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
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
		Action:      tea.String("CreateGroupConversation"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/interconnections/groups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGroupConversationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建互通群（支持普通互通群、跨钉两人群）
//
// @param request - CreateGroupConversationRequest
//
// @return CreateGroupConversationResponse
func (client *Client) CreateGroupConversation(request *CreateGroupConversationRequest) (_result *CreateGroupConversationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateGroupConversationHeaders{}
	_result = &CreateGroupConversationResponse{}
	_body, _err := client.CreateGroupConversationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建钉外账号
//
// @param request - CreateInterconnectionRequest
//
// @param headers - CreateInterconnectionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInterconnectionResponse
func (client *Client) CreateInterconnectionWithOptions(request *CreateInterconnectionRequest, headers *CreateInterconnectionHeaders, runtime *util.RuntimeOptions) (_result *CreateInterconnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Interconnections)) {
		body["interconnections"] = request.Interconnections
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
		Action:      tea.String("CreateInterconnection"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/interconnections"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInterconnectionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建钉外账号
//
// @param request - CreateInterconnectionRequest
//
// @return CreateInterconnectionResponse
func (client *Client) CreateInterconnection(request *CreateInterconnectionRequest) (_result *CreateInterconnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateInterconnectionHeaders{}
	_result = &CreateInterconnectionResponse{}
	_body, _err := client.CreateInterconnectionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建场景群会话
//
// @param request - CreateSceneGroupConversationRequest
//
// @param headers - CreateSceneGroupConversationHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSceneGroupConversationResponse
func (client *Client) CreateSceneGroupConversationWithOptions(request *CreateSceneGroupConversationRequest, headers *CreateSceneGroupConversationHeaders, runtime *util.RuntimeOptions) (_result *CreateSceneGroupConversationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Features)) {
		body["features"] = request.Features
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupOwnerId)) {
		body["groupOwnerId"] = request.GroupOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		body["icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.ManagementOptions)) {
		body["managementOptions"] = request.ManagementOptions
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["templateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdList)) {
		body["userIdList"] = request.UserIdList
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["uuid"] = request.Uuid
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
		Action:      tea.String("CreateSceneGroupConversation"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/sceneGroups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSceneGroupConversationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建场景群会话
//
// @param request - CreateSceneGroupConversationRequest
//
// @return CreateSceneGroupConversationResponse
func (client *Client) CreateSceneGroupConversation(request *CreateSceneGroupConversationRequest) (_result *CreateSceneGroupConversationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateSceneGroupConversationHeaders{}
	_result = &CreateSceneGroupConversationResponse{}
	_body, _err := client.CreateSceneGroupConversationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建店铺群
//
// @param request - CreateStoreGroupConversationRequest
//
// @param headers - CreateStoreGroupConversationHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStoreGroupConversationResponse
func (client *Client) CreateStoreGroupConversationWithOptions(request *CreateStoreGroupConversationRequest, headers *CreateStoreGroupConversationHeaders, runtime *util.RuntimeOptions) (_result *CreateStoreGroupConversationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUserId)) {
		body["appUserId"] = request.AppUserId
	}

	if !tea.BoolValue(util.IsUnset(request.BusinessUniqueKey)) {
		body["businessUniqueKey"] = request.BusinessUniqueKey
	}

	if !tea.BoolValue(util.IsUnset(request.GroupAvatar)) {
		body["groupAvatar"] = request.GroupAvatar
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupTemplateId)) {
		body["groupTemplateId"] = request.GroupTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
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
		Action:      tea.String("CreateStoreGroupConversation"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/interconnections/storeGroups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateStoreGroupConversationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建店铺群
//
// @param request - CreateStoreGroupConversationRequest
//
// @return CreateStoreGroupConversationResponse
func (client *Client) CreateStoreGroupConversation(request *CreateStoreGroupConversationRequest) (_result *CreateStoreGroupConversationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateStoreGroupConversationHeaders{}
	_result = &CreateStoreGroupConversationResponse{}
	_body, _err := client.CreateStoreGroupConversationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 链接增强规则调试
//
// @param request - DebugUnfurlingRegisterRequest
//
// @param headers - DebugUnfurlingRegisterHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DebugUnfurlingRegisterResponse
func (client *Client) DebugUnfurlingRegisterWithOptions(request *DebugUnfurlingRegisterRequest, headers *DebugUnfurlingRegisterHeaders, runtime *util.RuntimeOptions) (_result *DebugUnfurlingRegisterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.GrayGroupIdList)) {
		body["grayGroupIdList"] = request.GrayGroupIdList
	}

	if !tea.BoolValue(util.IsUnset(request.GrayUserIdList)) {
		body["grayUserIdList"] = request.GrayUserIdList
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
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
		Action:      tea.String("DebugUnfurlingRegister"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/unfurling/rules/debug"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DebugUnfurlingRegisterResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 链接增强规则调试
//
// @param request - DebugUnfurlingRegisterRequest
//
// @return DebugUnfurlingRegisterResponse
func (client *Client) DebugUnfurlingRegister(request *DebugUnfurlingRegisterRequest) (_result *DebugUnfurlingRegisterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DebugUnfurlingRegisterHeaders{}
	_result = &DebugUnfurlingRegisterResponse{}
	_body, _err := client.DebugUnfurlingRegisterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除企业文字表情
//
// @param request - DeleteOrgTextEmotionRequest
//
// @param headers - DeleteOrgTextEmotionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOrgTextEmotionResponse
func (client *Client) DeleteOrgTextEmotionWithOptions(request *DeleteOrgTextEmotionRequest, headers *DeleteOrgTextEmotionHeaders, runtime *util.RuntimeOptions) (_result *DeleteOrgTextEmotionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.EmotionIds)) {
		body["emotionIds"] = request.EmotionIds
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
		Action:      tea.String("DeleteOrgTextEmotion"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/organizations/textEmotions/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteOrgTextEmotionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除企业文字表情
//
// @param request - DeleteOrgTextEmotionRequest
//
// @return DeleteOrgTextEmotionResponse
func (client *Client) DeleteOrgTextEmotion(request *DeleteOrgTextEmotionRequest) (_result *DeleteOrgTextEmotionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteOrgTextEmotionHeaders{}
	_result = &DeleteOrgTextEmotionResponse{}
	_body, _err := client.DeleteOrgTextEmotionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解散互通群
//
// @param request - DismissGroupConversationRequest
//
// @param headers - DismissGroupConversationHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DismissGroupConversationResponse
func (client *Client) DismissGroupConversationWithOptions(request *DismissGroupConversationRequest, headers *DismissGroupConversationHeaders, runtime *util.RuntimeOptions) (_result *DismissGroupConversationResponse, _err error) {
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
		Action:      tea.String("DismissGroupConversation"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/interconnections/groups/dismiss"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DismissGroupConversationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解散互通群
//
// @param request - DismissGroupConversationRequest
//
// @return DismissGroupConversationResponse
func (client *Client) DismissGroupConversation(request *DismissGroupConversationRequest) (_result *DismissGroupConversationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DismissGroupConversationHeaders{}
	_result = &DismissGroupConversationResponse{}
	_body, _err := client.DismissGroupConversationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 冻结群
//
// @param request - FreezeGroupRequest
//
// @param headers - FreezeGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FreezeGroupResponse
func (client *Client) FreezeGroupWithOptions(request *FreezeGroupRequest, headers *FreezeGroupHeaders, runtime *util.RuntimeOptions) (_result *FreezeGroupResponse, _err error) {
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
		Action:      tea.String("FreezeGroup"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/chats/sceneGroups/freeze"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &FreezeGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 冻结群
//
// @param request - FreezeGroupRequest
//
// @return FreezeGroupResponse
func (client *Client) FreezeGroup(request *FreezeGroupRequest) (_result *FreezeGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &FreezeGroupHeaders{}
	_result = &FreezeGroupResponse{}
	_body, _err := client.FreezeGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建ToB会话地址
//
// @param request - GetConversationUrlRequest
//
// @param headers - GetConversationUrlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConversationUrlResponse
func (client *Client) GetConversationUrlWithOptions(request *GetConversationUrlRequest, headers *GetConversationUrlHeaders, runtime *util.RuntimeOptions) (_result *GetConversationUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUserId)) {
		body["appUserId"] = request.AppUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelCode)) {
		body["channelCode"] = request.ChannelCode
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		body["deviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
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
		Action:      tea.String("GetConversationUrl"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/conversations/urls"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConversationUrlResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建ToB会话地址
//
// @param request - GetConversationUrlRequest
//
// @return GetConversationUrlResponse
func (client *Client) GetConversationUrl(request *GetConversationUrlRequest) (_result *GetConversationUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetConversationUrlHeaders{}
	_result = &GetConversationUrlResponse{}
	_body, _err := client.GetConversationUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户家校群消息(图片&视频Z&富文本)
//
// @param request - GetFamilySchoolConversationMsgRequest
//
// @param headers - GetFamilySchoolConversationMsgHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFamilySchoolConversationMsgResponse
func (client *Client) GetFamilySchoolConversationMsgWithOptions(request *GetFamilySchoolConversationMsgRequest, headers *GetFamilySchoolConversationMsgHeaders, runtime *util.RuntimeOptions) (_result *GetFamilySchoolConversationMsgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.MsgTypes)) {
		body["msgTypes"] = request.MsgTypes
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("GetFamilySchoolConversationMsg"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/conversations/familySchools/messages/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFamilySchoolConversationMsgResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户家校群消息(图片&视频Z&富文本)
//
// @param request - GetFamilySchoolConversationMsgRequest
//
// @return GetFamilySchoolConversationMsgResponse
func (client *Client) GetFamilySchoolConversationMsg(request *GetFamilySchoolConversationMsgRequest) (_result *GetFamilySchoolConversationMsgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFamilySchoolConversationMsgHeaders{}
	_result = &GetFamilySchoolConversationMsgResponse{}
	_body, _err := client.GetFamilySchoolConversationMsgWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户家校群
//
// @param request - GetFamilySchoolConversationsRequest
//
// @param headers - GetFamilySchoolConversationsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFamilySchoolConversationsResponse
func (client *Client) GetFamilySchoolConversationsWithOptions(request *GetFamilySchoolConversationsRequest, headers *GetFamilySchoolConversationsHeaders, runtime *util.RuntimeOptions) (_result *GetFamilySchoolConversationsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("GetFamilySchoolConversations"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/conversations/familySchools/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFamilySchoolConversationsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户家校群
//
// @param request - GetFamilySchoolConversationsRequest
//
// @return GetFamilySchoolConversationsResponse
func (client *Client) GetFamilySchoolConversations(request *GetFamilySchoolConversationsRequest) (_result *GetFamilySchoolConversationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFamilySchoolConversationsHeaders{}
	_result = &GetFamilySchoolConversationsResponse{}
	_body, _err := client.GetFamilySchoolConversationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询企业内部群成员
//
// @param request - GetInnerGroupMembersRequest
//
// @param headers - GetInnerGroupMembersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInnerGroupMembersResponse
func (client *Client) GetInnerGroupMembersWithOptions(request *GetInnerGroupMembersRequest, headers *GetInnerGroupMembersHeaders, runtime *util.RuntimeOptions) (_result *GetInnerGroupMembersResponse, _err error) {
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
		Action:      tea.String("GetInnerGroupMembers"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/innerGroups/members/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInnerGroupMembersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询企业内部群成员
//
// @param request - GetInnerGroupMembersRequest
//
// @return GetInnerGroupMembersResponse
func (client *Client) GetInnerGroupMembers(request *GetInnerGroupMembersRequest) (_result *GetInnerGroupMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetInnerGroupMembersHeaders{}
	_result = &GetInnerGroupMembersResponse{}
	_body, _err := client.GetInnerGroupMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建客联互通会话地址
//
// @param request - GetInterconnectionUrlRequest
//
// @param headers - GetInterconnectionUrlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInterconnectionUrlResponse
func (client *Client) GetInterconnectionUrlWithOptions(request *GetInterconnectionUrlRequest, headers *GetInterconnectionUrlHeaders, runtime *util.RuntimeOptions) (_result *GetInterconnectionUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUserAvatar)) {
		body["appUserAvatar"] = request.AppUserAvatar
	}

	if !tea.BoolValue(util.IsUnset(request.AppUserAvatarType)) {
		body["appUserAvatarType"] = request.AppUserAvatarType
	}

	if !tea.BoolValue(util.IsUnset(request.AppUserId)) {
		body["appUserId"] = request.AppUserId
	}

	if !tea.BoolValue(util.IsUnset(request.AppUserMobileNumber)) {
		body["appUserMobileNumber"] = request.AppUserMobileNumber
	}

	if !tea.BoolValue(util.IsUnset(request.AppUserName)) {
		body["appUserName"] = request.AppUserName
	}

	if !tea.BoolValue(util.IsUnset(request.MsgPageType)) {
		body["msgPageType"] = request.MsgPageType
	}

	if !tea.BoolValue(util.IsUnset(request.QrCode)) {
		body["qrCode"] = request.QrCode
	}

	if !tea.BoolValue(util.IsUnset(request.Signature)) {
		body["signature"] = request.Signature
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCode)) {
		body["sourceCode"] = request.SourceCode
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		body["sourceType"] = request.SourceType
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
		Action:      tea.String("GetInterconnectionUrl"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/interconnections/sessions/urls"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInterconnectionUrlResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建客联互通会话地址
//
// @param request - GetInterconnectionUrlRequest
//
// @return GetInterconnectionUrlResponse
func (client *Client) GetInterconnectionUrl(request *GetInterconnectionUrlRequest) (_result *GetInterconnectionUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetInterconnectionUrlHeaders{}
	_result = &GetInterconnectionUrlResponse{}
	_body, _err := client.GetInterconnectionUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询最近活跃的企业内部群列表
//
// @param request - GetNewestInnerGroupsRequest
//
// @param headers - GetNewestInnerGroupsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNewestInnerGroupsResponse
func (client *Client) GetNewestInnerGroupsWithOptions(request *GetNewestInnerGroupsRequest, headers *GetNewestInnerGroupsHeaders, runtime *util.RuntimeOptions) (_result *GetNewestInnerGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      tea.String("GetNewestInnerGroups"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/activities/innerGroups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNewestInnerGroupsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询最近活跃的企业内部群列表
//
// @param request - GetNewestInnerGroupsRequest
//
// @return GetNewestInnerGroupsResponse
func (client *Client) GetNewestInnerGroups(request *GetNewestInnerGroupsRequest) (_result *GetNewestInnerGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetNewestInnerGroupsHeaders{}
	_result = &GetNewestInnerGroupsResponse{}
	_body, _err := client.GetNewestInnerGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询群简要信息
//
// @param request - GetSceneGroupInfoRequest
//
// @param headers - GetSceneGroupInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSceneGroupInfoResponse
func (client *Client) GetSceneGroupInfoWithOptions(request *GetSceneGroupInfoRequest, headers *GetSceneGroupInfoHeaders, runtime *util.RuntimeOptions) (_result *GetSceneGroupInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CoolAppCode)) {
		body["coolAppCode"] = request.CoolAppCode
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
		Action:      tea.String("GetSceneGroupInfo"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/sceneGroups/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSceneGroupInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询群简要信息
//
// @param request - GetSceneGroupInfoRequest
//
// @return GetSceneGroupInfoResponse
func (client *Client) GetSceneGroupInfo(request *GetSceneGroupInfoRequest) (_result *GetSceneGroupInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSceneGroupInfoHeaders{}
	_result = &GetSceneGroupInfoResponse{}
	_body, _err := client.GetSceneGroupInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询群成员
//
// @param request - GetSceneGroupMembersRequest
//
// @param headers - GetSceneGroupMembersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSceneGroupMembersResponse
func (client *Client) GetSceneGroupMembersWithOptions(request *GetSceneGroupMembersRequest, headers *GetSceneGroupMembersHeaders, runtime *util.RuntimeOptions) (_result *GetSceneGroupMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CoolAppCode)) {
		body["coolAppCode"] = request.CoolAppCode
	}

	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		body["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["size"] = request.Size
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
		Action:      tea.String("GetSceneGroupMembers"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/sceneGroups/members/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSceneGroupMembersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询群成员
//
// @param request - GetSceneGroupMembersRequest
//
// @return GetSceneGroupMembersResponse
func (client *Client) GetSceneGroupMembers(request *GetSceneGroupMembersRequest) (_result *GetSceneGroupMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSceneGroupMembersHeaders{}
	_result = &GetSceneGroupMembersResponse{}
	_body, _err := client.GetSceneGroupMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询场景群模板消息存档能力开启状态
//
// @param headers - GetSceneGroupTemplateMessageOpenStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSceneGroupTemplateMessageOpenStatusResponse
func (client *Client) GetSceneGroupTemplateMessageOpenStatusWithOptions(templateId *string, headers *GetSceneGroupTemplateMessageOpenStatusHeaders, runtime *util.RuntimeOptions) (_result *GetSceneGroupTemplateMessageOpenStatusResponse, _err error) {
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
		Action:      tea.String("GetSceneGroupTemplateMessageOpenStatus"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/chats/sceneGroups/templates/" + tea.StringValue(templateId) + "/messageOpenStatuses"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSceneGroupTemplateMessageOpenStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询场景群模板消息存档能力开启状态
//
// @return GetSceneGroupTemplateMessageOpenStatusResponse
func (client *Client) GetSceneGroupTemplateMessageOpenStatus(templateId *string) (_result *GetSceneGroupTemplateMessageOpenStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSceneGroupTemplateMessageOpenStatusHeaders{}
	_result = &GetSceneGroupTemplateMessageOpenStatusResponse{}
	_body, _err := client.GetSceneGroupTemplateMessageOpenStatusWithOptions(templateId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取单聊会话的OpenConversationId
//
// @param request - GetSingleChatOpenConversationIdRequest
//
// @param headers - GetSingleChatOpenConversationIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSingleChatOpenConversationIdResponse
func (client *Client) GetSingleChatOpenConversationIdWithOptions(request *GetSingleChatOpenConversationIdRequest, headers *GetSingleChatOpenConversationIdHeaders, runtime *util.RuntimeOptions) (_result *GetSingleChatOpenConversationIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId1)) {
		body["userId1"] = request.UserId1
	}

	if !tea.BoolValue(util.IsUnset(request.UserId2)) {
		body["userId2"] = request.UserId2
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
		Action:      tea.String("GetSingleChatOpenConversationId"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/privateChats/openConversationId/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSingleChatOpenConversationIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取单聊会话的OpenConversationId
//
// @param request - GetSingleChatOpenConversationIdRequest
//
// @return GetSingleChatOpenConversationIdResponse
func (client *Client) GetSingleChatOpenConversationId(request *GetSingleChatOpenConversationIdRequest) (_result *GetSingleChatOpenConversationIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSingleChatOpenConversationIdHeaders{}
	_result = &GetSingleChatOpenConversationIdResponse{}
	_body, _err := client.GetSingleChatOpenConversationIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询群信息（超管接口）
//
// @param request - GetSuperAdminOpenSceneGroupInfoRequest
//
// @param headers - GetSuperAdminOpenSceneGroupInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSuperAdminOpenSceneGroupInfoResponse
func (client *Client) GetSuperAdminOpenSceneGroupInfoWithOptions(request *GetSuperAdminOpenSceneGroupInfoRequest, headers *GetSuperAdminOpenSceneGroupInfoHeaders, runtime *util.RuntimeOptions) (_result *GetSuperAdminOpenSceneGroupInfoResponse, _err error) {
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
		Action:      tea.String("GetSuperAdminOpenSceneGroupInfo"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/chats/sceneGroups/groupInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSuperAdminOpenSceneGroupInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询群信息（超管接口）
//
// @param request - GetSuperAdminOpenSceneGroupInfoRequest
//
// @return GetSuperAdminOpenSceneGroupInfoResponse
func (client *Client) GetSuperAdminOpenSceneGroupInfo(request *GetSuperAdminOpenSceneGroupInfoRequest) (_result *GetSuperAdminOpenSceneGroupInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSuperAdminOpenSceneGroupInfoHeaders{}
	_result = &GetSuperAdminOpenSceneGroupInfoResponse{}
	_body, _err := client.GetSuperAdminOpenSceneGroupInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 群禁言
//
// @param request - GroupBanWordsRequest
//
// @param headers - GroupBanWordsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GroupBanWordsResponse
func (client *Client) GroupBanWordsWithOptions(request *GroupBanWordsRequest, headers *GroupBanWordsHeaders, runtime *util.RuntimeOptions) (_result *GroupBanWordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BanWordsMode)) {
		body["banWordsMode"] = request.BanWordsMode
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		body["options"] = request.Options
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
		Action:      tea.String("GroupBanWords"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/groups/words/ban"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("none"),
	}
	_result = &GroupBanWordsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 群禁言
//
// @param request - GroupBanWordsRequest
//
// @return GroupBanWordsResponse
func (client *Client) GroupBanWords(request *GroupBanWordsRequest) (_result *GroupBanWordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GroupBanWordsHeaders{}
	_result = &GroupBanWordsResponse{}
	_body, _err := client.GroupBanWordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 群容量扩容询价
//
// @param request - GroupCapacityInquiryRequest
//
// @param headers - GroupCapacityInquiryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GroupCapacityInquiryResponse
func (client *Client) GroupCapacityInquiryWithOptions(request *GroupCapacityInquiryRequest, headers *GroupCapacityInquiryHeaders, runtime *util.RuntimeOptions) (_result *GroupCapacityInquiryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EffectiveDuration)) {
		body["effectiveDuration"] = request.EffectiveDuration
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		body["options"] = request.Options
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCapacity)) {
		body["targetCapacity"] = request.TargetCapacity
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
		Action:      tea.String("GroupCapacityInquiry"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/groups/capacities/inquiries/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GroupCapacityInquiryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 群容量扩容询价
//
// @param request - GroupCapacityInquiryRequest
//
// @return GroupCapacityInquiryResponse
func (client *Client) GroupCapacityInquiry(request *GroupCapacityInquiryRequest) (_result *GroupCapacityInquiryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GroupCapacityInquiryHeaders{}
	_result = &GroupCapacityInquiryResponse{}
	_body, _err := client.GroupCapacityInquiryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 群容量扩容确认下单
//
// @param request - GroupCapacityOrderConfirmRequest
//
// @param headers - GroupCapacityOrderConfirmHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GroupCapacityOrderConfirmResponse
func (client *Client) GroupCapacityOrderConfirmWithOptions(request *GroupCapacityOrderConfirmRequest, headers *GroupCapacityOrderConfirmHeaders, runtime *util.RuntimeOptions) (_result *GroupCapacityOrderConfirmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		body["orderId"] = request.OrderId
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
		Action:      tea.String("GroupCapacityOrderConfirm"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/groups/capacities/orders/confirm"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GroupCapacityOrderConfirmResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 群容量扩容确认下单
//
// @param request - GroupCapacityOrderConfirmRequest
//
// @return GroupCapacityOrderConfirmResponse
func (client *Client) GroupCapacityOrderConfirm(request *GroupCapacityOrderConfirmRequest) (_result *GroupCapacityOrderConfirmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GroupCapacityOrderConfirmHeaders{}
	_result = &GroupCapacityOrderConfirmResponse{}
	_body, _err := client.GroupCapacityOrderConfirmWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 群容量请求扩容下单
//
// @param request - GroupCapacityOrderPlaceRequest
//
// @param headers - GroupCapacityOrderPlaceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GroupCapacityOrderPlaceResponse
func (client *Client) GroupCapacityOrderPlaceWithOptions(request *GroupCapacityOrderPlaceRequest, headers *GroupCapacityOrderPlaceHeaders, runtime *util.RuntimeOptions) (_result *GroupCapacityOrderPlaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActualPrice)) {
		body["actualPrice"] = request.ActualPrice
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentCapacity)) {
		body["currentCapacity"] = request.CurrentCapacity
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentEffectUntil)) {
		body["currentEffectUntil"] = request.CurrentEffectUntil
	}

	if !tea.BoolValue(util.IsUnset(request.Discount)) {
		body["discount"] = request.Discount
	}

	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		body["extInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.MarkedPrice)) {
		body["markedPrice"] = request.MarkedPrice
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCapacity)) {
		body["targetCapacity"] = request.TargetCapacity
	}

	if !tea.BoolValue(util.IsUnset(request.TargetEffectUntil)) {
		body["targetEffectUntil"] = request.TargetEffectUntil
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		body["token"] = request.Token
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
		Action:      tea.String("GroupCapacityOrderPlace"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/groups/capacities/orders/place"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GroupCapacityOrderPlaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 群容量请求扩容下单
//
// @param request - GroupCapacityOrderPlaceRequest
//
// @return GroupCapacityOrderPlaceResponse
func (client *Client) GroupCapacityOrderPlace(request *GroupCapacityOrderPlaceRequest) (_result *GroupCapacityOrderPlaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GroupCapacityOrderPlaceHeaders{}
	_result = &GroupCapacityOrderPlaceResponse{}
	_body, _err := client.GroupCapacityOrderPlaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据群链接、群号等检索条件，查询群信息
//
// @param request - GroupManageQueryRequest
//
// @param headers - GroupManageQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GroupManageQueryResponse
func (client *Client) GroupManageQueryWithOptions(request *GroupManageQueryRequest, headers *GroupManageQueryHeaders, runtime *util.RuntimeOptions) (_result *GroupManageQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreatedAfter)) {
		body["createdAfter"] = request.CreatedAfter
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["groupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupMemberSamples)) {
		body["groupMemberSamples"] = request.GroupMemberSamples
	}

	if !tea.BoolValue(util.IsUnset(request.GroupOwner)) {
		body["groupOwner"] = request.GroupOwner
	}

	if !tea.BoolValue(util.IsUnset(request.GroupTitleKeywords)) {
		body["groupTitleKeywords"] = request.GroupTitleKeywords
	}

	if !tea.BoolValue(util.IsUnset(request.GroupUrl)) {
		body["groupUrl"] = request.GroupUrl
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.MembersOver)) {
		body["membersOver"] = request.MembersOver
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
		Action:      tea.String("GroupManageQuery"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/groups/managements/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GroupManageQueryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据群链接、群号等检索条件，查询群信息
//
// @param request - GroupManageQueryRequest
//
// @return GroupManageQueryResponse
func (client *Client) GroupManageQuery(request *GroupManageQueryRequest) (_result *GroupManageQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GroupManageQueryHeaders{}
	_result = &GroupManageQueryResponse{}
	_body, _err := client.GroupManageQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 群管理缩容
//
// @param request - GroupManageReduceRequest
//
// @param headers - GroupManageReduceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GroupManageReduceResponse
func (client *Client) GroupManageReduceWithOptions(request *GroupManageReduceRequest, headers *GroupManageReduceHeaders, runtime *util.RuntimeOptions) (_result *GroupManageReduceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CapacityLimit)) {
		body["capacityLimit"] = request.CapacityLimit
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		body["options"] = request.Options
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
		Action:      tea.String("GroupManageReduce"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/groups/capacities/reduce"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("none"),
	}
	_result = &GroupManageReduceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 群管理缩容
//
// @param request - GroupManageReduceRequest
//
// @return GroupManageReduceResponse
func (client *Client) GroupManageReduce(request *GroupManageReduceRequest) (_result *GroupManageReduceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GroupManageReduceHeaders{}
	_result = &GroupManageReduceResponse{}
	_body, _err := client.GroupManageReduceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导入群聊会话
//
// @param request - ImportGroupChatRequest
//
// @param headers - ImportGroupChatHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportGroupChatResponse
func (client *Client) ImportGroupChatWithOptions(request *ImportGroupChatRequest, headers *ImportGroupChatHeaders, runtime *util.RuntimeOptions) (_result *ImportGroupChatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdminIds)) {
		body["adminIds"] = request.AdminIds
	}

	if !tea.BoolValue(util.IsUnset(request.CreateAt)) {
		body["createAt"] = request.CreateAt
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		body["icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.ImportUuid)) {
		body["importUuid"] = request.ImportUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Owner)) {
		body["owner"] = request.Owner
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["templateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.UserList)) {
		body["userList"] = request.UserList
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
		Action:      tea.String("ImportGroupChat"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/groupChats/import"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportGroupChatResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导入群聊会话
//
// @param request - ImportGroupChatRequest
//
// @return ImportGroupChatResponse
func (client *Client) ImportGroupChat(request *ImportGroupChatRequest) (_result *ImportGroupChatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ImportGroupChatHeaders{}
	_result = &ImportGroupChatResponse{}
	_body, _err := client.ImportGroupChatWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导入消息
//
// @param request - ImportMessageRequest
//
// @param headers - ImportMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportMessageResponse
func (client *Client) ImportMessageWithOptions(request *ImportMessageRequest, headers *ImportMessageHeaders, runtime *util.RuntimeOptions) (_result *ImportMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTime)) {
		body["createTime"] = request.CreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.ImportUuid)) {
		body["importUuid"] = request.ImportUuid
	}

	if !tea.BoolValue(util.IsUnset(request.MsgReadStatusSetting)) {
		body["msgReadStatusSetting"] = request.MsgReadStatusSetting
	}

	if !tea.BoolValue(util.IsUnset(request.MsgType)) {
		body["msgType"] = request.MsgType
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Receivers)) {
		body["receivers"] = request.Receivers
	}

	if !tea.BoolValue(util.IsUnset(request.SenderId)) {
		body["senderId"] = request.SenderId
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
		Action:      tea.String("ImportMessage"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/messages/import"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导入消息
//
// @param request - ImportMessageRequest
//
// @return ImportMessageResponse
func (client *Client) ImportMessage(request *ImportMessageRequest) (_result *ImportMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ImportMessageHeaders{}
	_result = &ImportMessageResponse{}
	_body, _err := client.ImportMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 安装机器人到组织
//
// @param request - InstallRobotToOrgRequest
//
// @param headers - InstallRobotToOrgHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallRobotToOrgResponse
func (client *Client) InstallRobotToOrgWithOptions(request *InstallRobotToOrgRequest, headers *InstallRobotToOrgHeaders, runtime *util.RuntimeOptions) (_result *InstallRobotToOrgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Brief)) {
		body["brief"] = request.Brief
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		body["icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OutgoingToken)) {
		body["outgoingToken"] = request.OutgoingToken
	}

	if !tea.BoolValue(util.IsUnset(request.OutgoingUrl)) {
		body["outgoingUrl"] = request.OutgoingUrl
	}

	if !tea.BoolValue(util.IsUnset(request.PreviewMediaId)) {
		body["previewMediaId"] = request.PreviewMediaId
	}

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
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
		Action:      tea.String("InstallRobotToOrg"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/organizations/robots/install"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallRobotToOrgResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 安装机器人到组织
//
// @param request - InstallRobotToOrgRequest
//
// @return InstallRobotToOrgResponse
func (client *Client) InstallRobotToOrg(request *InstallRobotToOrgRequest) (_result *InstallRobotToOrgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InstallRobotToOrgHeaders{}
	_result = &InstallRobotToOrgResponse{}
	_body, _err := client.InstallRobotToOrgWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建可交互式实例
//
// @param request - InteractiveCardCreateInstanceRequest
//
// @param headers - InteractiveCardCreateInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InteractiveCardCreateInstanceResponse
func (client *Client) InteractiveCardCreateInstanceWithOptions(request *InteractiveCardCreateInstanceRequest, headers *InteractiveCardCreateInstanceHeaders, runtime *util.RuntimeOptions) (_result *InteractiveCardCreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallbackRouteKey)) {
		body["callbackRouteKey"] = request.CallbackRouteKey
	}

	if !tea.BoolValue(util.IsUnset(request.CardData)) {
		body["cardData"] = request.CardData
	}

	if !tea.BoolValue(util.IsUnset(request.CardTemplateId)) {
		body["cardTemplateId"] = request.CardTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.ChatBotId)) {
		body["chatBotId"] = request.ChatBotId
	}

	if !tea.BoolValue(util.IsUnset(request.ConversationType)) {
		body["conversationType"] = request.ConversationType
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OutTrackId)) {
		body["outTrackId"] = request.OutTrackId
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateData)) {
		body["privateData"] = request.PrivateData
	}

	if !tea.BoolValue(util.IsUnset(request.PullStrategy)) {
		body["pullStrategy"] = request.PullStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUserIdList)) {
		body["receiverUserIdList"] = request.ReceiverUserIdList
	}

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
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
		Action:      tea.String("InteractiveCardCreateInstance"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/interactiveCards/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &InteractiveCardCreateInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建可交互式实例
//
// @param request - InteractiveCardCreateInstanceRequest
//
// @return InteractiveCardCreateInstanceResponse
func (client *Client) InteractiveCardCreateInstance(request *InteractiveCardCreateInstanceRequest) (_result *InteractiveCardCreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InteractiveCardCreateInstanceHeaders{}
	_result = &InteractiveCardCreateInstanceResponse{}
	_body, _err := client.InteractiveCardCreateInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查组织下所有的场景群模版列表
//
// @param request - ListGroupTemplatesByOrgIdRequest
//
// @param headers - ListGroupTemplatesByOrgIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupTemplatesByOrgIdResponse
func (client *Client) ListGroupTemplatesByOrgIdWithOptions(request *ListGroupTemplatesByOrgIdRequest, headers *ListGroupTemplatesByOrgIdHeaders, runtime *util.RuntimeOptions) (_result *ListGroupTemplatesByOrgIdResponse, _err error) {
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
		Action:      tea.String("ListGroupTemplatesByOrgId"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/chats/sceneGroups/templates/lists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGroupTemplatesByOrgIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查组织下所有的场景群模版列表
//
// @param request - ListGroupTemplatesByOrgIdRequest
//
// @return ListGroupTemplatesByOrgIdResponse
func (client *Client) ListGroupTemplatesByOrgId(request *ListGroupTemplatesByOrgIdRequest) (_result *ListGroupTemplatesByOrgIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListGroupTemplatesByOrgIdHeaders{}
	_result = &ListGroupTemplatesByOrgIdResponse{}
	_body, _err := client.ListGroupTemplatesByOrgIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 拉取企业的所有文字表情，包含正常使用的、已经删除了的、安全审核不通过的文字表情
//
// @param headers - ListOrgTextEmotionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrgTextEmotionResponse
func (client *Client) ListOrgTextEmotionWithOptions(headers *ListOrgTextEmotionHeaders, runtime *util.RuntimeOptions) (_result *ListOrgTextEmotionResponse, _err error) {
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
		Action:      tea.String("ListOrgTextEmotion"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/organizations/textEmotions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOrgTextEmotionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 拉取企业的所有文字表情，包含正常使用的、已经删除了的、安全审核不通过的文字表情
//
// @return ListOrgTextEmotionResponse
func (client *Client) ListOrgTextEmotion() (_result *ListOrgTextEmotionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListOrgTextEmotionHeaders{}
	_result = &ListOrgTextEmotionResponse{}
	_body, _err := client.ListOrgTextEmotionWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据模板id查询关联的群
//
// @param request - ListSceneGroupsByTemplateIdRequest
//
// @param headers - ListSceneGroupsByTemplateIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSceneGroupsByTemplateIdResponse
func (client *Client) ListSceneGroupsByTemplateIdWithOptions(templateId *string, request *ListSceneGroupsByTemplateIdRequest, headers *ListSceneGroupsByTemplateIdHeaders, runtime *util.RuntimeOptions) (_result *ListSceneGroupsByTemplateIdResponse, _err error) {
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
		Action:      tea.String("ListSceneGroupsByTemplateId"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/chats/sceneGroups/templates/" + tea.StringValue(templateId) + "/lists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSceneGroupsByTemplateIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据模板id查询关联的群
//
// @param request - ListSceneGroupsByTemplateIdRequest
//
// @return ListSceneGroupsByTemplateIdResponse
func (client *Client) ListSceneGroupsByTemplateId(templateId *string, request *ListSceneGroupsByTemplateIdRequest) (_result *ListSceneGroupsByTemplateIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListSceneGroupsByTemplateIdHeaders{}
	_result = &ListSceneGroupsByTemplateIdResponse{}
	_body, _err := client.ListSceneGroupsByTemplateIdWithOptions(templateId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 链接增强规则下线
//
// @param request - OfflineUnfurlingRegisterRequest
//
// @param headers - OfflineUnfurlingRegisterHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineUnfurlingRegisterResponse
func (client *Client) OfflineUnfurlingRegisterWithOptions(request *OfflineUnfurlingRegisterRequest, headers *OfflineUnfurlingRegisterHeaders, runtime *util.RuntimeOptions) (_result *OfflineUnfurlingRegisterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
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
		Action:      tea.String("OfflineUnfurlingRegister"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/unfurling/rules/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &OfflineUnfurlingRegisterResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 链接增强规则下线
//
// @param request - OfflineUnfurlingRegisterRequest
//
// @return OfflineUnfurlingRegisterResponse
func (client *Client) OfflineUnfurlingRegister(request *OfflineUnfurlingRegisterRequest) (_result *OfflineUnfurlingRegisterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OfflineUnfurlingRegisterHeaders{}
	_result = &OfflineUnfurlingRegisterResponse{}
	_body, _err := client.OfflineUnfurlingRegisterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开放场景群新增群角色
//
// @param request - OpenGroupRoleAddRequest
//
// @param headers - OpenGroupRoleAddHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenGroupRoleAddResponse
func (client *Client) OpenGroupRoleAddWithOptions(request *OpenGroupRoleAddRequest, headers *OpenGroupRoleAddHeaders, runtime *util.RuntimeOptions) (_result *OpenGroupRoleAddResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		body["roleName"] = request.RoleName
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
		Action:      tea.String("OpenGroupRoleAdd"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/groups/roles"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenGroupRoleAddResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开放场景群新增群角色
//
// @param request - OpenGroupRoleAddRequest
//
// @return OpenGroupRoleAddResponse
func (client *Client) OpenGroupRoleAdd(request *OpenGroupRoleAddRequest) (_result *OpenGroupRoleAddResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OpenGroupRoleAddHeaders{}
	_result = &OpenGroupRoleAddResponse{}
	_body, _err := client.OpenGroupRoleAddWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开放场景群群角色查询
//
// @param request - OpenGroupRoleQueryRequest
//
// @param headers - OpenGroupRoleQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenGroupRoleQueryResponse
func (client *Client) OpenGroupRoleQueryWithOptions(request *OpenGroupRoleQueryRequest, headers *OpenGroupRoleQueryHeaders, runtime *util.RuntimeOptions) (_result *OpenGroupRoleQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
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
		Action:      tea.String("OpenGroupRoleQuery"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/groups/roles/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenGroupRoleQueryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开放场景群群角色查询
//
// @param request - OpenGroupRoleQueryRequest
//
// @return OpenGroupRoleQueryResponse
func (client *Client) OpenGroupRoleQuery(request *OpenGroupRoleQueryRequest) (_result *OpenGroupRoleQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OpenGroupRoleQueryHeaders{}
	_result = &OpenGroupRoleQueryResponse{}
	_body, _err := client.OpenGroupRoleQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开放场景群群角色移除
//
// @param request - OpenGroupRoleRemoveRequest
//
// @param headers - OpenGroupRoleRemoveHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenGroupRoleRemoveResponse
func (client *Client) OpenGroupRoleRemoveWithOptions(request *OpenGroupRoleRemoveRequest, headers *OpenGroupRoleRemoveHeaders, runtime *util.RuntimeOptions) (_result *OpenGroupRoleRemoveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenRoleId)) {
		body["openRoleId"] = request.OpenRoleId
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
		Action:      tea.String("OpenGroupRoleRemove"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/groups/roles/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenGroupRoleRemoveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开放场景群群角色移除
//
// @param request - OpenGroupRoleRemoveRequest
//
// @return OpenGroupRoleRemoveResponse
func (client *Client) OpenGroupRoleRemove(request *OpenGroupRoleRemoveRequest) (_result *OpenGroupRoleRemoveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OpenGroupRoleRemoveHeaders{}
	_result = &OpenGroupRoleRemoveResponse{}
	_body, _err := client.OpenGroupRoleRemoveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开放场景群群角色变更
//
// @param request - OpenGroupRoleUpdateRequest
//
// @param headers - OpenGroupRoleUpdateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenGroupRoleUpdateResponse
func (client *Client) OpenGroupRoleUpdateWithOptions(request *OpenGroupRoleUpdateRequest, headers *OpenGroupRoleUpdateHeaders, runtime *util.RuntimeOptions) (_result *OpenGroupRoleUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenRoleId)) {
		body["openRoleId"] = request.OpenRoleId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		body["roleName"] = request.RoleName
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
		Action:      tea.String("OpenGroupRoleUpdate"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/groups/roles"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenGroupRoleUpdateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开放场景群群角色变更
//
// @param request - OpenGroupRoleUpdateRequest
//
// @return OpenGroupRoleUpdateResponse
func (client *Client) OpenGroupRoleUpdate(request *OpenGroupRoleUpdateRequest) (_result *OpenGroupRoleUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OpenGroupRoleUpdateHeaders{}
	_result = &OpenGroupRoleUpdateResponse{}
	_body, _err := client.OpenGroupRoleUpdateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开放场景群群成员的群角色信息查询
//
// @param request - OpenGroupUserRoleQueryRequest
//
// @param headers - OpenGroupUserRoleQueryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenGroupUserRoleQueryResponse
func (client *Client) OpenGroupUserRoleQueryWithOptions(request *OpenGroupUserRoleQueryRequest, headers *OpenGroupUserRoleQueryHeaders, runtime *util.RuntimeOptions) (_result *OpenGroupUserRoleQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.ViewedUserId)) {
		body["viewedUserId"] = request.ViewedUserId
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
		Action:      tea.String("OpenGroupUserRoleQuery"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/groups/users/roles/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenGroupUserRoleQueryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开放场景群群成员的群角色信息查询
//
// @param request - OpenGroupUserRoleQueryRequest
//
// @return OpenGroupUserRoleQueryResponse
func (client *Client) OpenGroupUserRoleQuery(request *OpenGroupUserRoleQueryRequest) (_result *OpenGroupUserRoleQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OpenGroupUserRoleQueryHeaders{}
	_result = &OpenGroupUserRoleQueryResponse{}
	_body, _err := client.OpenGroupUserRoleQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 内部群转部门群
//
// @param request - OpenInnerGroupTransferToDeptGroupRequest
//
// @param headers - OpenInnerGroupTransferToDeptGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenInnerGroupTransferToDeptGroupResponse
func (client *Client) OpenInnerGroupTransferToDeptGroupWithOptions(request *OpenInnerGroupTransferToDeptGroupRequest, headers *OpenInnerGroupTransferToDeptGroupHeaders, runtime *util.RuntimeOptions) (_result *OpenInnerGroupTransferToDeptGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
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
		Action:      tea.String("OpenInnerGroupTransferToDeptGroup"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/innerGroups/transferToDeptGroups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenInnerGroupTransferToDeptGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 内部群转部门群
//
// @param request - OpenInnerGroupTransferToDeptGroupRequest
//
// @return OpenInnerGroupTransferToDeptGroupResponse
func (client *Client) OpenInnerGroupTransferToDeptGroup(request *OpenInnerGroupTransferToDeptGroupRequest) (_result *OpenInnerGroupTransferToDeptGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OpenInnerGroupTransferToDeptGroupHeaders{}
	_result = &OpenInnerGroupTransferToDeptGroupResponse{}
	_body, _err := client.OpenInnerGroupTransferToDeptGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 群搜索
//
// @param request - OpenSearchGroupListRequest
//
// @param headers - OpenSearchGroupListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenSearchGroupListResponse
func (client *Client) OpenSearchGroupListWithOptions(request *OpenSearchGroupListRequest, headers *OpenSearchGroupListHeaders, runtime *util.RuntimeOptions) (_result *OpenSearchGroupListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		body["keyword"] = request.Keyword
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
		Action:      tea.String("OpenSearchGroupList"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/groups/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenSearchGroupListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 群搜索
//
// @param request - OpenSearchGroupListRequest
//
// @return OpenSearchGroupListResponse
func (client *Client) OpenSearchGroupList(request *OpenSearchGroupListRequest) (_result *OpenSearchGroupListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OpenSearchGroupListHeaders{}
	_result = &OpenSearchGroupListResponse{}
	_body, _err := client.OpenSearchGroupListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 以个人身份发送卡片消息
//
// @param request - OpenUserSendCardMessageRequest
//
// @param headers - OpenUserSendCardMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenUserSendCardMessageResponse
func (client *Client) OpenUserSendCardMessageWithOptions(request *OpenUserSendCardMessageRequest, headers *OpenUserSendCardMessageHeaders, runtime *util.RuntimeOptions) (_result *OpenUserSendCardMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CardContent)) {
		body["cardContent"] = request.CardContent
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiveUserId)) {
		body["receiveUserId"] = request.ReceiveUserId
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
		Action:      tea.String("OpenUserSendCardMessage"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/cardMessages/users/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenUserSendCardMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 以个人身份发送卡片消息
//
// @param request - OpenUserSendCardMessageRequest
//
// @return OpenUserSendCardMessageResponse
func (client *Client) OpenUserSendCardMessage(request *OpenUserSendCardMessageRequest) (_result *OpenUserSendCardMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OpenUserSendCardMessageHeaders{}
	_result = &OpenUserSendCardMessageResponse{}
	_body, _err := client.OpenUserSendCardMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 以用户身份发送卡片消息
//
// @param request - PersonalSendCardMessageRequest
//
// @param headers - PersonalSendCardMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PersonalSendCardMessageResponse
func (client *Client) PersonalSendCardMessageWithOptions(request *PersonalSendCardMessageRequest, headers *PersonalSendCardMessageHeaders, runtime *util.RuntimeOptions) (_result *PersonalSendCardMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AtUserIds)) {
		body["atUserIds"] = request.AtUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.CardContent)) {
		body["cardContent"] = request.CardContent
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiveUserId)) {
		body["receiveUserId"] = request.ReceiveUserId
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
		Action:      tea.String("PersonalSendCardMessage"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/me/messages/cards/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PersonalSendCardMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 以用户身份发送卡片消息
//
// @param request - PersonalSendCardMessageRequest
//
// @return PersonalSendCardMessageResponse
func (client *Client) PersonalSendCardMessage(request *PersonalSendCardMessageRequest) (_result *PersonalSendCardMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PersonalSendCardMessageHeaders{}
	_result = &PersonalSendCardMessageResponse{}
	_body, _err := client.PersonalSendCardMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 成员授权场景下查询群信息
//
// @param request - QueryGroupInfoByMemberAuthRequest
//
// @param headers - QueryGroupInfoByMemberAuthHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGroupInfoByMemberAuthResponse
func (client *Client) QueryGroupInfoByMemberAuthWithOptions(request *QueryGroupInfoByMemberAuthRequest, headers *QueryGroupInfoByMemberAuthHeaders, runtime *util.RuntimeOptions) (_result *QueryGroupInfoByMemberAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CoolAppCode)) {
		body["coolAppCode"] = request.CoolAppCode
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
		Action:      tea.String("QueryGroupInfoByMemberAuth"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/memberAuthorizations/groups/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGroupInfoByMemberAuthResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 成员授权场景下查询群信息
//
// @param request - QueryGroupInfoByMemberAuthRequest
//
// @return QueryGroupInfoByMemberAuthResponse
func (client *Client) QueryGroupInfoByMemberAuth(request *QueryGroupInfoByMemberAuthRequest) (_result *QueryGroupInfoByMemberAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryGroupInfoByMemberAuthHeaders{}
	_result = &QueryGroupInfoByMemberAuthResponse{}
	_body, _err := client.QueryGroupInfoByMemberAuthWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询群成员列表
//
// @param request - QueryGroupMemberRequest
//
// @param headers - QueryGroupMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGroupMemberResponse
func (client *Client) QueryGroupMemberWithOptions(request *QueryGroupMemberRequest, headers *QueryGroupMemberHeaders, runtime *util.RuntimeOptions) (_result *QueryGroupMemberResponse, _err error) {
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
		Action:      tea.String("QueryGroupMember"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/interconnections/conversations/members"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGroupMemberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询群成员列表
//
// @param request - QueryGroupMemberRequest
//
// @return QueryGroupMemberResponse
func (client *Client) QueryGroupMember(request *QueryGroupMemberRequest) (_result *QueryGroupMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryGroupMemberHeaders{}
	_result = &QueryGroupMemberResponse{}
	_body, _err := client.QueryGroupMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 成员授权场景下查询群成员
//
// @param request - QueryGroupMemberByMemberAuthRequest
//
// @param headers - QueryGroupMemberByMemberAuthHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGroupMemberByMemberAuthResponse
func (client *Client) QueryGroupMemberByMemberAuthWithOptions(request *QueryGroupMemberByMemberAuthRequest, headers *QueryGroupMemberByMemberAuthHeaders, runtime *util.RuntimeOptions) (_result *QueryGroupMemberByMemberAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CoolAppCode)) {
		body["coolAppCode"] = request.CoolAppCode
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
		Action:      tea.String("QueryGroupMemberByMemberAuth"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/memberAuthorizations/groups/members/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGroupMemberByMemberAuthResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 成员授权场景下查询群成员
//
// @param request - QueryGroupMemberByMemberAuthRequest
//
// @return QueryGroupMemberByMemberAuthResponse
func (client *Client) QueryGroupMemberByMemberAuth(request *QueryGroupMemberByMemberAuthRequest) (_result *QueryGroupMemberByMemberAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryGroupMemberByMemberAuthHeaders{}
	_result = &QueryGroupMemberByMemberAuthResponse{}
	_body, _err := client.QueryGroupMemberByMemberAuthWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询群禁言状态
//
// @param request - QueryGroupMuteStatusRequest
//
// @param headers - QueryGroupMuteStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGroupMuteStatusResponse
func (client *Client) QueryGroupMuteStatusWithOptions(request *QueryGroupMuteStatusRequest, headers *QueryGroupMuteStatusHeaders, runtime *util.RuntimeOptions) (_result *QueryGroupMuteStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		query["openConversationId"] = request.OpenConversationId
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
		Action:      tea.String("QueryGroupMuteStatus"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/sceneGroups/muteSettings"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGroupMuteStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询群禁言状态
//
// @param request - QueryGroupMuteStatusRequest
//
// @return QueryGroupMuteStatusResponse
func (client *Client) QueryGroupMuteStatus(request *QueryGroupMuteStatusRequest) (_result *QueryGroupMuteStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryGroupMuteStatusHeaders{}
	_result = &QueryGroupMuteStatusResponse{}
	_body, _err := client.QueryGroupMuteStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 读取群成员列表
//
// @param request - QueryInnerGroupMemberListRequest
//
// @param headers - QueryInnerGroupMemberListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInnerGroupMemberListResponse
func (client *Client) QueryInnerGroupMemberListWithOptions(request *QueryInnerGroupMemberListRequest, headers *QueryInnerGroupMemberListHeaders, runtime *util.RuntimeOptions) (_result *QueryInnerGroupMemberListResponse, _err error) {
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
		Action:      tea.String("QueryInnerGroupMemberList"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/innerGroups/memberLists/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryInnerGroupMemberListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 读取群成员列表
//
// @param request - QueryInnerGroupMemberListRequest
//
// @return QueryInnerGroupMemberListResponse
func (client *Client) QueryInnerGroupMemberList(request *QueryInnerGroupMemberListRequest) (_result *QueryInnerGroupMemberListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryInnerGroupMemberListHeaders{}
	_result = &QueryInnerGroupMemberListResponse{}
	_body, _err := client.QueryInnerGroupMemberListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询最近活跃的企业内部群列表
//
// @param request - QueryInnerGroupRecentListRequest
//
// @param headers - QueryInnerGroupRecentListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInnerGroupRecentListResponse
func (client *Client) QueryInnerGroupRecentListWithOptions(request *QueryInnerGroupRecentListRequest, headers *QueryInnerGroupRecentListHeaders, runtime *util.RuntimeOptions) (_result *QueryInnerGroupRecentListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      tea.String("QueryInnerGroupRecentList"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/innerGroups/recentLists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryInnerGroupRecentListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询最近活跃的企业内部群列表
//
// @param request - QueryInnerGroupRecentListRequest
//
// @return QueryInnerGroupRecentListResponse
func (client *Client) QueryInnerGroupRecentList(request *QueryInnerGroupRecentListRequest) (_result *QueryInnerGroupRecentListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryInnerGroupRecentListHeaders{}
	_result = &QueryInnerGroupRecentListResponse{}
	_body, _err := client.QueryInnerGroupRecentListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询群内具有指定群角色的所有群成员
//
// @param request - QueryMembersOfGroupRoleRequest
//
// @param headers - QueryMembersOfGroupRoleHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMembersOfGroupRoleResponse
func (client *Client) QueryMembersOfGroupRoleWithOptions(request *QueryMembersOfGroupRoleRequest, headers *QueryMembersOfGroupRoleHeaders, runtime *util.RuntimeOptions) (_result *QueryMembersOfGroupRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenRoleId)) {
		body["openRoleId"] = request.OpenRoleId
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamp)) {
		body["timestamp"] = request.Timestamp
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
		Action:      tea.String("QueryMembersOfGroupRole"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/sceneGroups/roles/members/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMembersOfGroupRoleResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询群内具有指定群角色的所有群成员
//
// @param request - QueryMembersOfGroupRoleRequest
//
// @return QueryMembersOfGroupRoleResponse
func (client *Client) QueryMembersOfGroupRole(request *QueryMembersOfGroupRoleRequest) (_result *QueryMembersOfGroupRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMembersOfGroupRoleHeaders{}
	_result = &QueryMembersOfGroupRoleResponse{}
	_body, _err := client.QueryMembersOfGroupRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据openTaskId查询消息发送结果
//
// @param request - QueryMessageSendResultRequest
//
// @param headers - QueryMessageSendResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMessageSendResultResponse
func (client *Client) QueryMessageSendResultWithOptions(request *QueryMessageSendResultRequest, headers *QueryMessageSendResultHeaders, runtime *util.RuntimeOptions) (_result *QueryMessageSendResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenTaskId)) {
		body["openTaskId"] = request.OpenTaskId
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
		Action:      tea.String("QueryMessageSendResult"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/messages/sendResults/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMessageSendResultResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据openTaskId查询消息发送结果
//
// @param request - QueryMessageSendResultRequest
//
// @return QueryMessageSendResultResponse
func (client *Client) QueryMessageSendResult(request *QueryMessageSendResultRequest) (_result *QueryMessageSendResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMessageSendResultHeaders{}
	_result = &QueryMessageSendResultResponse{}
	_body, _err := client.QueryMessageSendResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据单聊会话及发送方获取接收方用户信息
//
// @param request - QueryOpenConversationReceiveUserRequest
//
// @param headers - QueryOpenConversationReceiveUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOpenConversationReceiveUserResponse
func (client *Client) QueryOpenConversationReceiveUserWithOptions(request *QueryOpenConversationReceiveUserRequest, headers *QueryOpenConversationReceiveUserHeaders, runtime *util.RuntimeOptions) (_result *QueryOpenConversationReceiveUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.SendUserId)) {
		body["sendUserId"] = request.SendUserId
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
		Action:      tea.String("QueryOpenConversationReceiveUser"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/otoChat/receiveUsers/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOpenConversationReceiveUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据单聊会话及发送方获取接收方用户信息
//
// @param request - QueryOpenConversationReceiveUserRequest
//
// @return QueryOpenConversationReceiveUserResponse
func (client *Client) QueryOpenConversationReceiveUser(request *QueryOpenConversationReceiveUserRequest) (_result *QueryOpenConversationReceiveUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOpenConversationReceiveUserHeaders{}
	_result = &QueryOpenConversationReceiveUserResponse{}
	_body, _err := client.QueryOpenConversationReceiveUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取群基础信息
//
// @param request - QueryOpenGroupBaseInfoRequest
//
// @param headers - QueryOpenGroupBaseInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOpenGroupBaseInfoResponse
func (client *Client) QueryOpenGroupBaseInfoWithOptions(request *QueryOpenGroupBaseInfoRequest, headers *QueryOpenGroupBaseInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryOpenGroupBaseInfoResponse, _err error) {
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
		Action:      tea.String("QueryOpenGroupBaseInfo"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/groups/baseInfos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOpenGroupBaseInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取群基础信息
//
// @param request - QueryOpenGroupBaseInfoRequest
//
// @return QueryOpenGroupBaseInfoResponse
func (client *Client) QueryOpenGroupBaseInfo(request *QueryOpenGroupBaseInfoRequest) (_result *QueryOpenGroupBaseInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOpenGroupBaseInfoHeaders{}
	_result = &QueryOpenGroupBaseInfoResponse{}
	_body, _err := client.QueryOpenGroupBaseInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用户身份查询消息已读未读状态
//
// @param request - QueryPersonalMessageReadStatusRequest
//
// @param headers - QueryPersonalMessageReadStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPersonalMessageReadStatusResponse
func (client *Client) QueryPersonalMessageReadStatusWithOptions(request *QueryPersonalMessageReadStatusRequest, headers *QueryPersonalMessageReadStatusHeaders, runtime *util.RuntimeOptions) (_result *QueryPersonalMessageReadStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenMessageId)) {
		body["openMessageId"] = request.OpenMessageId
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
		Action:      tea.String("QueryPersonalMessageReadStatus"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/me/messages/readStatuses/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPersonalMessageReadStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用户身份查询消息已读未读状态
//
// @param request - QueryPersonalMessageReadStatusRequest
//
// @return QueryPersonalMessageReadStatusResponse
func (client *Client) QueryPersonalMessageReadStatus(request *QueryPersonalMessageReadStatusRequest) (_result *QueryPersonalMessageReadStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryPersonalMessageReadStatusHeaders{}
	_result = &QueryPersonalMessageReadStatusResponse{}
	_body, _err := client.QueryPersonalMessageReadStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取最近联系人及群组
//
// @param request - QueryRecentConversationsRequest
//
// @param headers - QueryRecentConversationsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRecentConversationsResponse
func (client *Client) QueryRecentConversationsWithOptions(request *QueryRecentConversationsRequest, headers *QueryRecentConversationsHeaders, runtime *util.RuntimeOptions) (_result *QueryRecentConversationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OnlyHuman)) {
		body["onlyHuman"] = request.OnlyHuman
	}

	if !tea.BoolValue(util.IsUnset(request.OnlyInnerGroup)) {
		body["onlyInnerGroup"] = request.OnlyInnerGroup
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
		Action:      tea.String("QueryRecentConversations"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/conversations/recentLists/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRecentConversationsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取最近联系人及群组
//
// @param request - QueryRecentConversationsRequest
//
// @return QueryRecentConversationsResponse
func (client *Client) QueryRecentConversations(request *QueryRecentConversationsRequest) (_result *QueryRecentConversationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryRecentConversationsHeaders{}
	_result = &QueryRecentConversationsResponse{}
	_body, _err := client.QueryRecentConversationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询群内群模板机器人
//
// @param request - QuerySceneGroupTemplateRobotRequest
//
// @param headers - QuerySceneGroupTemplateRobotHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySceneGroupTemplateRobotResponse
func (client *Client) QuerySceneGroupTemplateRobotWithOptions(request *QuerySceneGroupTemplateRobotRequest, headers *QuerySceneGroupTemplateRobotHeaders, runtime *util.RuntimeOptions) (_result *QuerySceneGroupTemplateRobotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		query["openConversationId"] = request.OpenConversationId
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
		Action:      tea.String("QuerySceneGroupTemplateRobot"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/sceneGroups/templates/robots"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySceneGroupTemplateRobotResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询群内群模板机器人
//
// @param request - QuerySceneGroupTemplateRobotRequest
//
// @return QuerySceneGroupTemplateRobotResponse
func (client *Client) QuerySceneGroupTemplateRobot(request *QuerySceneGroupTemplateRobotRequest) (_result *QuerySceneGroupTemplateRobotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QuerySceneGroupTemplateRobotHeaders{}
	_result = &QuerySceneGroupTemplateRobotResponse{}
	_body, _err := client.QuerySceneGroupTemplateRobotWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询群信息
//
// @param request - QuerySingleGroupRequest
//
// @param headers - QuerySingleGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySingleGroupResponse
func (client *Client) QuerySingleGroupWithOptions(request *QuerySingleGroupRequest, headers *QuerySingleGroupHeaders, runtime *util.RuntimeOptions) (_result *QuerySingleGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupMembers)) {
		body["groupMembers"] = request.GroupMembers
	}

	if !tea.BoolValue(util.IsUnset(request.GroupTemplateId)) {
		body["groupTemplateId"] = request.GroupTemplateId
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
		Action:      tea.String("QuerySingleGroup"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/interconnections/doubleGroups/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySingleGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询群信息
//
// @param request - QuerySingleGroupRequest
//
// @return QuerySingleGroupResponse
func (client *Client) QuerySingleGroup(request *QuerySingleGroupRequest) (_result *QuerySingleGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QuerySingleGroupHeaders{}
	_result = &QuerySingleGroupResponse{}
	_body, _err := client.QuerySingleGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询未读消息数
//
// @param request - QueryUnReadMessageRequest
//
// @param headers - QueryUnReadMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUnReadMessageResponse
func (client *Client) QueryUnReadMessageWithOptions(request *QueryUnReadMessageRequest, headers *QueryUnReadMessageHeaders, runtime *util.RuntimeOptions) (_result *QueryUnReadMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUserId)) {
		body["appUserId"] = request.AppUserId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationIds)) {
		body["openConversationIds"] = request.OpenConversationIds
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
		Action:      tea.String("QueryUnReadMessage"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/interconnections/unReadMsgs/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUnReadMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询未读消息数
//
// @param request - QueryUnReadMessageRequest
//
// @return QueryUnReadMessageResponse
func (client *Client) QueryUnReadMessage(request *QueryUnReadMessageRequest) (_result *QueryUnReadMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUnReadMessageHeaders{}
	_result = &QueryUnReadMessageResponse{}
	_body, _err := client.QueryUnReadMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询链接查询链接增强注册信息创建者
//
// @param request - QueryUnfurlingRegisterCreatorRequest
//
// @param headers - QueryUnfurlingRegisterCreatorHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUnfurlingRegisterCreatorResponse
func (client *Client) QueryUnfurlingRegisterCreatorWithOptions(request *QueryUnfurlingRegisterCreatorRequest, headers *QueryUnfurlingRegisterCreatorHeaders, runtime *util.RuntimeOptions) (_result *QueryUnfurlingRegisterCreatorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["path"] = request.Path
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
		Action:      tea.String("QueryUnfurlingRegisterCreator"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/unfurling/rules/creators"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUnfurlingRegisterCreatorResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询链接查询链接增强注册信息创建者
//
// @param request - QueryUnfurlingRegisterCreatorRequest
//
// @return QueryUnfurlingRegisterCreatorResponse
func (client *Client) QueryUnfurlingRegisterCreator(request *QueryUnfurlingRegisterCreatorRequest) (_result *QueryUnfurlingRegisterCreatorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUnfurlingRegisterCreatorHeaders{}
	_result = &QueryUnfurlingRegisterCreatorResponse{}
	_body, _err := client.QueryUnfurlingRegisterCreatorWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询链接增强注册信息列表
//
// @param request - QueryUnfurlingRegisterInfoRequest
//
// @param headers - QueryUnfurlingRegisterInfoHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUnfurlingRegisterInfoResponse
func (client *Client) QueryUnfurlingRegisterInfoWithOptions(request *QueryUnfurlingRegisterInfoRequest, headers *QueryUnfurlingRegisterInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryUnfurlingRegisterInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["appId"] = request.AppId
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
		Action:      tea.String("QueryUnfurlingRegisterInfo"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/unfurling/rules"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUnfurlingRegisterInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询链接增强注册信息列表
//
// @param request - QueryUnfurlingRegisterInfoRequest
//
// @return QueryUnfurlingRegisterInfoResponse
func (client *Client) QueryUnfurlingRegisterInfo(request *QueryUnfurlingRegisterInfoRequest) (_result *QueryUnfurlingRegisterInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUnfurlingRegisterInfoHeaders{}
	_result = &QueryUnfurlingRegisterInfoResponse{}
	_body, _err := client.QueryUnfurlingRegisterInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询群主视角群LastMessage时间
//
// @param request - QueryUserViewGroupLastMessageTimeRequest
//
// @param headers - QueryUserViewGroupLastMessageTimeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserViewGroupLastMessageTimeResponse
func (client *Client) QueryUserViewGroupLastMessageTimeWithOptions(request *QueryUserViewGroupLastMessageTimeRequest, headers *QueryUserViewGroupLastMessageTimeHeaders, runtime *util.RuntimeOptions) (_result *QueryUserViewGroupLastMessageTimeResponse, _err error) {
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
		Action:      tea.String("QueryUserViewGroupLastMessageTime"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/chats/lastMessageTime/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserViewGroupLastMessageTimeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询群主视角群LastMessage时间
//
// @param request - QueryUserViewGroupLastMessageTimeRequest
//
// @return QueryUserViewGroupLastMessageTimeResponse
func (client *Client) QueryUserViewGroupLastMessageTime(request *QueryUserViewGroupLastMessageTimeRequest) (_result *QueryUserViewGroupLastMessageTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserViewGroupLastMessageTimeHeaders{}
	_result = &QueryUserViewGroupLastMessageTimeResponse{}
	_body, _err := client.QueryUserViewGroupLastMessageTimeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用户身份设置消息状态为已读
//
// @param request - ReadPersonalMessageRequest
//
// @param headers - ReadPersonalMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadPersonalMessageResponse
func (client *Client) ReadPersonalMessageWithOptions(request *ReadPersonalMessageRequest, headers *ReadPersonalMessageHeaders, runtime *util.RuntimeOptions) (_result *ReadPersonalMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingOpenConversationMessageIdArray)) {
		body["dingOpenConversationMessageIdArray"] = request.DingOpenConversationMessageIdArray
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
		Action:      tea.String("ReadPersonalMessage"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/me/messages/readStatuses/set"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ReadPersonalMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用户身份设置消息状态为已读
//
// @param request - ReadPersonalMessageRequest
//
// @return ReadPersonalMessageResponse
func (client *Client) ReadPersonalMessage(request *ReadPersonalMessageRequest) (_result *ReadPersonalMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ReadPersonalMessageHeaders{}
	_result = &ReadPersonalMessageResponse{}
	_body, _err := client.ReadPersonalMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用户身份撤回消息
//
// @param request - RecallPersonalMessageRequest
//
// @param headers - RecallPersonalMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecallPersonalMessageResponse
func (client *Client) RecallPersonalMessageWithOptions(request *RecallPersonalMessageRequest, headers *RecallPersonalMessageHeaders, runtime *util.RuntimeOptions) (_result *RecallPersonalMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenMessageId)) {
		body["openMessageId"] = request.OpenMessageId
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
		Action:      tea.String("RecallPersonalMessage"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/me/messages/recall"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RecallPersonalMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用户身份撤回消息
//
// @param request - RecallPersonalMessageRequest
//
// @return RecallPersonalMessageResponse
func (client *Client) RecallPersonalMessage(request *RecallPersonalMessageRequest) (_result *RecallPersonalMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RecallPersonalMessageHeaders{}
	_result = &RecallPersonalMessageResponse{}
	_body, _err := client.RecallPersonalMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 链接增强规则发布
//
// @param request - ReleaseUnfurlingRegisterRequest
//
// @param headers - ReleaseUnfurlingRegisterHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseUnfurlingRegisterResponse
func (client *Client) ReleaseUnfurlingRegisterWithOptions(request *ReleaseUnfurlingRegisterRequest, headers *ReleaseUnfurlingRegisterHeaders, runtime *util.RuntimeOptions) (_result *ReleaseUnfurlingRegisterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
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
		Action:      tea.String("ReleaseUnfurlingRegister"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/unfurling/rules/publish"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseUnfurlingRegisterResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 链接增强规则发布
//
// @param request - ReleaseUnfurlingRegisterRequest
//
// @return ReleaseUnfurlingRegisterResponse
func (client *Client) ReleaseUnfurlingRegister(request *ReleaseUnfurlingRegisterRequest) (_result *ReleaseUnfurlingRegisterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ReleaseUnfurlingRegisterHeaders{}
	_result = &ReleaseUnfurlingRegisterResponse{}
	_body, _err := client.ReleaseUnfurlingRegisterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移除会话机器人
//
// @param request - RemoveRobotFromConversationRequest
//
// @param headers - RemoveRobotFromConversationHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveRobotFromConversationResponse
func (client *Client) RemoveRobotFromConversationWithOptions(request *RemoveRobotFromConversationRequest, headers *RemoveRobotFromConversationHeaders, runtime *util.RuntimeOptions) (_result *RemoveRobotFromConversationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChatBotUserId)) {
		body["chatBotUserId"] = request.ChatBotUserId
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
		Action:      tea.String("RemoveRobotFromConversation"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/conversations/robots/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveRobotFromConversationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移除会话机器人
//
// @param request - RemoveRobotFromConversationRequest
//
// @return RemoveRobotFromConversationResponse
func (client *Client) RemoveRobotFromConversation(request *RemoveRobotFromConversationRequest) (_result *RemoveRobotFromConversationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveRobotFromConversationHeaders{}
	_result = &RemoveRobotFromConversationResponse{}
	_body, _err := client.RemoveRobotFromConversationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据关键词搜索企业内部群
//
// @param request - SearchInnerGroupsRequest
//
// @param headers - SearchInnerGroupsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchInnerGroupsResponse
func (client *Client) SearchInnerGroupsWithOptions(request *SearchInnerGroupsRequest, headers *SearchInnerGroupsHeaders, runtime *util.RuntimeOptions) (_result *SearchInnerGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["searchKey"] = request.SearchKey
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
		Action:      tea.String("SearchInnerGroups"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/innerGroups/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchInnerGroupsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据关键词搜索企业内部群
//
// @param request - SearchInnerGroupsRequest
//
// @return SearchInnerGroupsResponse
func (client *Client) SearchInnerGroups(request *SearchInnerGroupsRequest) (_result *SearchInnerGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchInnerGroupsHeaders{}
	_result = &SearchInnerGroupsResponse{}
	_body, _err := client.SearchInnerGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发送可交互式动态卡片
//
// @param request - SendInteractiveCardRequest
//
// @param headers - SendInteractiveCardHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendInteractiveCardResponse
func (client *Client) SendInteractiveCardWithOptions(request *SendInteractiveCardRequest, headers *SendInteractiveCardHeaders, runtime *util.RuntimeOptions) (_result *SendInteractiveCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AtOpenIds)) {
		body["atOpenIds"] = request.AtOpenIds
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackRouteKey)) {
		body["callbackRouteKey"] = request.CallbackRouteKey
	}

	if !tea.BoolValue(util.IsUnset(request.CardData)) {
		body["cardData"] = request.CardData
	}

	if !tea.BoolValue(util.IsUnset(request.CardOptions)) {
		body["cardOptions"] = request.CardOptions
	}

	if !tea.BoolValue(util.IsUnset(request.CardTemplateId)) {
		body["cardTemplateId"] = request.CardTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.ChatBotId)) {
		body["chatBotId"] = request.ChatBotId
	}

	if !tea.BoolValue(util.IsUnset(request.ConversationType)) {
		body["conversationType"] = request.ConversationType
	}

	if !tea.BoolValue(util.IsUnset(request.DigitalWorkerCode)) {
		body["digitalWorkerCode"] = request.DigitalWorkerCode
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OutTrackId)) {
		body["outTrackId"] = request.OutTrackId
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateData)) {
		body["privateData"] = request.PrivateData
	}

	if !tea.BoolValue(util.IsUnset(request.PullStrategy)) {
		body["pullStrategy"] = request.PullStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUserIdList)) {
		body["receiverUserIdList"] = request.ReceiverUserIdList
	}

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
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
		Action:      tea.String("SendInteractiveCard"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/interactiveCards/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendInteractiveCardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送可交互式动态卡片
//
// @param request - SendInteractiveCardRequest
//
// @return SendInteractiveCardResponse
func (client *Client) SendInteractiveCard(request *SendInteractiveCardRequest) (_result *SendInteractiveCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendInteractiveCardHeaders{}
	_result = &SendInteractiveCardResponse{}
	_body, _err := client.SendInteractiveCardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 人与人单聊发送可交互式动态卡片
//
// @param request - SendOTOInteractiveCardRequest
//
// @param headers - SendOTOInteractiveCardHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendOTOInteractiveCardResponse
func (client *Client) SendOTOInteractiveCardWithOptions(request *SendOTOInteractiveCardRequest, headers *SendOTOInteractiveCardHeaders, runtime *util.RuntimeOptions) (_result *SendOTOInteractiveCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AtOpenIds)) {
		body["atOpenIds"] = request.AtOpenIds
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackRouteKey)) {
		body["callbackRouteKey"] = request.CallbackRouteKey
	}

	if !tea.BoolValue(util.IsUnset(request.CardData)) {
		body["cardData"] = request.CardData
	}

	if !tea.BoolValue(util.IsUnset(request.CardOptions)) {
		body["cardOptions"] = request.CardOptions
	}

	if !tea.BoolValue(util.IsUnset(request.CardTemplateId)) {
		body["cardTemplateId"] = request.CardTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OutTrackId)) {
		body["outTrackId"] = request.OutTrackId
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateData)) {
		body["privateData"] = request.PrivateData
	}

	if !tea.BoolValue(util.IsUnset(request.PullStrategy)) {
		body["pullStrategy"] = request.PullStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUserIdList)) {
		body["receiverUserIdList"] = request.ReceiverUserIdList
	}

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
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
		Action:      tea.String("SendOTOInteractiveCard"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/privateChat/interactiveCards/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendOTOInteractiveCardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人与人单聊发送可交互式动态卡片
//
// @param request - SendOTOInteractiveCardRequest
//
// @return SendOTOInteractiveCardResponse
func (client *Client) SendOTOInteractiveCard(request *SendOTOInteractiveCardRequest) (_result *SendOTOInteractiveCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendOTOInteractiveCardHeaders{}
	_result = &SendOTOInteractiveCardResponse{}
	_body, _err := client.SendOTOInteractiveCardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 委托权限发消息
//
// @param request - SendPersonalMessageRequest
//
// @param headers - SendPersonalMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendPersonalMessageResponse
func (client *Client) SendPersonalMessageWithOptions(request *SendPersonalMessageRequest, headers *SendPersonalMessageHeaders, runtime *util.RuntimeOptions) (_result *SendPersonalMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.MsgType)) {
		body["msgType"] = request.MsgType
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
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
		Action:      tea.String("SendPersonalMessage"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/me/messages/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendPersonalMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 委托权限发消息
//
// @param request - SendPersonalMessageRequest
//
// @return SendPersonalMessageResponse
func (client *Client) SendPersonalMessage(request *SendPersonalMessageRequest) (_result *SendPersonalMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendPersonalMessageHeaders{}
	_result = &SendPersonalMessageResponse{}
	_body, _err := client.SendPersonalMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 机器人发送互动卡片（普通版）
//
// @param request - SendRobotInteractiveCardRequest
//
// @param headers - SendRobotInteractiveCardHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendRobotInteractiveCardResponse
func (client *Client) SendRobotInteractiveCardWithOptions(request *SendRobotInteractiveCardRequest, headers *SendRobotInteractiveCardHeaders, runtime *util.RuntimeOptions) (_result *SendRobotInteractiveCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallbackUrl)) {
		body["callbackUrl"] = request.CallbackUrl
	}

	if !tea.BoolValue(util.IsUnset(request.CardBizId)) {
		body["cardBizId"] = request.CardBizId
	}

	if !tea.BoolValue(util.IsUnset(request.CardData)) {
		body["cardData"] = request.CardData
	}

	if !tea.BoolValue(util.IsUnset(request.CardTemplateId)) {
		body["cardTemplateId"] = request.CardTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.PullStrategy)) {
		body["pullStrategy"] = request.PullStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
	}

	if !tea.BoolValue(util.IsUnset(request.SendOptions)) {
		body["sendOptions"] = request.SendOptions
	}

	if !tea.BoolValue(util.IsUnset(request.SingleChatReceiver)) {
		body["singleChatReceiver"] = request.SingleChatReceiver
	}

	if !tea.BoolValue(util.IsUnset(request.UnionIdPrivateDataMap)) {
		body["unionIdPrivateDataMap"] = request.UnionIdPrivateDataMap
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdPrivateDataMap)) {
		body["userIdPrivateDataMap"] = request.UserIdPrivateDataMap
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
		Action:      tea.String("SendRobotInteractiveCard"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/v1.0/robot/interactiveCards/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendRobotInteractiveCardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机器人发送互动卡片（普通版）
//
// @param request - SendRobotInteractiveCardRequest
//
// @return SendRobotInteractiveCardResponse
func (client *Client) SendRobotInteractiveCard(request *SendRobotInteractiveCardRequest) (_result *SendRobotInteractiveCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendRobotInteractiveCardHeaders{}
	_result = &SendRobotInteractiveCardResponse{}
	_body, _err := client.SendRobotInteractiveCardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 机器人发送消息
//
// @param request - SendRobotMessageRequest
//
// @param headers - SendRobotMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendRobotMessageResponse
func (client *Client) SendRobotMessageWithOptions(request *SendRobotMessageRequest, headers *SendRobotMessageHeaders, runtime *util.RuntimeOptions) (_result *SendRobotMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AtAll)) {
		body["atAll"] = request.AtAll
	}

	if !tea.BoolValue(util.IsUnset(request.AtAppUserId)) {
		body["atAppUserId"] = request.AtAppUserId
	}

	if !tea.BoolValue(util.IsUnset(request.AtDingUserId)) {
		body["atDingUserId"] = request.AtDingUserId
	}

	if !tea.BoolValue(util.IsUnset(request.MsgContent)) {
		body["msgContent"] = request.MsgContent
	}

	if !tea.BoolValue(util.IsUnset(request.MsgType)) {
		body["msgType"] = request.MsgType
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationIds)) {
		body["openConversationIds"] = request.OpenConversationIds
	}

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
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
		Action:      tea.String("SendRobotMessage"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/interconnections/robotMessages/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendRobotMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机器人发送消息
//
// @param request - SendRobotMessageRequest
//
// @return SendRobotMessageResponse
func (client *Client) SendRobotMessage(request *SendRobotMessageRequest) (_result *SendRobotMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendRobotMessageHeaders{}
	_result = &SendRobotMessageResponse{}
	_body, _err := client.SendRobotMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发送模板响应式可交互式卡片
//
// @param request - SendTemplateInteractiveCardRequest
//
// @param headers - SendTemplateInteractiveCardHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendTemplateInteractiveCardResponse
func (client *Client) SendTemplateInteractiveCardWithOptions(request *SendTemplateInteractiveCardRequest, headers *SendTemplateInteractiveCardHeaders, runtime *util.RuntimeOptions) (_result *SendTemplateInteractiveCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallbackUrl)) {
		body["callbackUrl"] = request.CallbackUrl
	}

	if !tea.BoolValue(util.IsUnset(request.CardData)) {
		body["cardData"] = request.CardData
	}

	if !tea.BoolValue(util.IsUnset(request.CardTemplateId)) {
		body["cardTemplateId"] = request.CardTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OutTrackId)) {
		body["outTrackId"] = request.OutTrackId
	}

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
	}

	if !tea.BoolValue(util.IsUnset(request.SendOptions)) {
		body["sendOptions"] = request.SendOptions
	}

	if !tea.BoolValue(util.IsUnset(request.SingleChatReceiver)) {
		body["singleChatReceiver"] = request.SingleChatReceiver
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
		Action:      tea.String("SendTemplateInteractiveCard"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/interactiveCards/templates/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendTemplateInteractiveCardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送模板响应式可交互式卡片
//
// @param request - SendTemplateInteractiveCardRequest
//
// @return SendTemplateInteractiveCardResponse
func (client *Client) SendTemplateInteractiveCard(request *SendTemplateInteractiveCardRequest) (_result *SendTemplateInteractiveCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendTemplateInteractiveCardHeaders{}
	_result = &SendTemplateInteractiveCardResponse{}
	_body, _err := client.SendTemplateInteractiveCardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置侧边栏
//
// @param request - SetRightPanelRequest
//
// @param headers - SetRightPanelHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetRightPanelResponse
func (client *Client) SetRightPanelWithOptions(request *SetRightPanelRequest, headers *SetRightPanelHeaders, runtime *util.RuntimeOptions) (_result *SetRightPanelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.RightPanelClosePermitted)) {
		body["rightPanelClosePermitted"] = request.RightPanelClosePermitted
	}

	if !tea.BoolValue(util.IsUnset(request.RightPanelOpenStatus)) {
		body["rightPanelOpenStatus"] = request.RightPanelOpenStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.WebWndParams)) {
		body["webWndParams"] = request.WebWndParams
	}

	if !tea.BoolValue(util.IsUnset(request.Width)) {
		body["width"] = request.Width
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
		Action:      tea.String("SetRightPanel"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/rightPanels/set"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SetRightPanelResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置侧边栏
//
// @param request - SetRightPanelRequest
//
// @return SetRightPanelResponse
func (client *Client) SetRightPanel(request *SetRightPanelRequest) (_result *SetRightPanelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SetRightPanelHeaders{}
	_result = &SetRightPanelResponse{}
	_body, _err := client.SetRightPanelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启用群模板(超管接口)
//
// @param request - SuperAdminApplyTemplateRequest
//
// @param headers - SuperAdminApplyTemplateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuperAdminApplyTemplateResponse
func (client *Client) SuperAdminApplyTemplateWithOptions(request *SuperAdminApplyTemplateRequest, headers *SuperAdminApplyTemplateHeaders, runtime *util.RuntimeOptions) (_result *SuperAdminApplyTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUserId)) {
		body["ownerUserId"] = request.OwnerUserId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["templateId"] = request.TemplateId
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
		Action:      tea.String("SuperAdminApplyTemplate"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/chats/scenegroups/templates/apply"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SuperAdminApplyTemplateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启用群模板(超管接口)
//
// @param request - SuperAdminApplyTemplateRequest
//
// @return SuperAdminApplyTemplateResponse
func (client *Client) SuperAdminApplyTemplate(request *SuperAdminApplyTemplateRequest) (_result *SuperAdminApplyTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SuperAdminApplyTemplateHeaders{}
	_result = &SuperAdminApplyTemplateResponse{}
	_body, _err := client.SuperAdminApplyTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停用群模板（超管接口）
//
// @param request - SuperAdminCloseTemplateRequest
//
// @param headers - SuperAdminCloseTemplateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuperAdminCloseTemplateResponse
func (client *Client) SuperAdminCloseTemplateWithOptions(request *SuperAdminCloseTemplateRequest, headers *SuperAdminCloseTemplateHeaders, runtime *util.RuntimeOptions) (_result *SuperAdminCloseTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUserId)) {
		body["ownerUserId"] = request.OwnerUserId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["templateId"] = request.TemplateId
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
		Action:      tea.String("SuperAdminCloseTemplate"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/chats/scenegroups/templates/close"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SuperAdminCloseTemplateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停用群模板（超管接口）
//
// @param request - SuperAdminCloseTemplateRequest
//
// @return SuperAdminCloseTemplateResponse
func (client *Client) SuperAdminCloseTemplate(request *SuperAdminCloseTemplateRequest) (_result *SuperAdminCloseTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SuperAdminCloseTemplateHeaders{}
	_result = &SuperAdminCloseTemplateResponse{}
	_body, _err := client.SuperAdminCloseTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 钉钉吊顶卡片关闭
//
// @param request - TopboxCloseRequest
//
// @param headers - TopboxCloseHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TopboxCloseResponse
func (client *Client) TopboxCloseWithOptions(request *TopboxCloseRequest, headers *TopboxCloseHeaders, runtime *util.RuntimeOptions) (_result *TopboxCloseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationType)) {
		body["conversationType"] = request.ConversationType
	}

	if !tea.BoolValue(util.IsUnset(request.CoolAppCode)) {
		body["coolAppCode"] = request.CoolAppCode
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OutTrackId)) {
		body["outTrackId"] = request.OutTrackId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUserIdList)) {
		body["receiverUserIdList"] = request.ReceiverUserIdList
	}

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
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
		Action:      tea.String("TopboxClose"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/topBoxes/close"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("none"),
	}
	_result = &TopboxCloseResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 钉钉吊顶卡片关闭
//
// @param request - TopboxCloseRequest
//
// @return TopboxCloseResponse
func (client *Client) TopboxClose(request *TopboxCloseRequest) (_result *TopboxCloseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TopboxCloseHeaders{}
	_result = &TopboxCloseResponse{}
	_body, _err := client.TopboxCloseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 钉钉吊顶卡片开启
//
// @param request - TopboxOpenRequest
//
// @param headers - TopboxOpenHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TopboxOpenResponse
func (client *Client) TopboxOpenWithOptions(request *TopboxOpenRequest, headers *TopboxOpenHeaders, runtime *util.RuntimeOptions) (_result *TopboxOpenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationType)) {
		body["conversationType"] = request.ConversationType
	}

	if !tea.BoolValue(util.IsUnset(request.CoolAppCode)) {
		body["coolAppCode"] = request.CoolAppCode
	}

	if !tea.BoolValue(util.IsUnset(request.ExpiredTime)) {
		body["expiredTime"] = request.ExpiredTime
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OutTrackId)) {
		body["outTrackId"] = request.OutTrackId
	}

	if !tea.BoolValue(util.IsUnset(request.Platforms)) {
		body["platforms"] = request.Platforms
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUserIdList)) {
		body["receiverUserIdList"] = request.ReceiverUserIdList
	}

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
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
		Action:      tea.String("TopboxOpen"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/topBoxes/open"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("none"),
	}
	_result = &TopboxOpenResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 钉钉吊顶卡片开启
//
// @param request - TopboxOpenRequest
//
// @return TopboxOpenResponse
func (client *Client) TopboxOpen(request *TopboxOpenRequest) (_result *TopboxOpenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TopboxOpenHeaders{}
	_result = &TopboxOpenResponse{}
	_body, _err := client.TopboxOpenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新钉内用户C端展示的头像和名称（互通群、钉内两人群）
//
// @param request - UpdateClientServiceRequest
//
// @param headers - UpdateClientServiceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateClientServiceResponse
func (client *Client) UpdateClientServiceWithOptions(request *UpdateClientServiceRequest, headers *UpdateClientServiceHeaders, runtime *util.RuntimeOptions) (_result *UpdateClientServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AvatarUrl)) {
		body["avatarUrl"] = request.AvatarUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ResetAvatar)) {
		body["resetAvatar"] = request.ResetAvatar
	}

	if !tea.BoolValue(util.IsUnset(request.ResetUserName)) {
		body["resetUserName"] = request.ResetUserName
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["userName"] = request.UserName
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
		Action:      tea.String("UpdateClientService"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/interconnections/clientServices/avatarAndName"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateClientServiceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新钉内用户C端展示的头像和名称（互通群、钉内两人群）
//
// @param request - UpdateClientServiceRequest
//
// @return UpdateClientServiceResponse
func (client *Client) UpdateClientService(request *UpdateClientServiceRequest) (_result *UpdateClientServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateClientServiceHeaders{}
	_result = &UpdateClientServiceResponse{}
	_body, _err := client.UpdateClientServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改群头像
//
// @param request - UpdateGroupAvatarRequest
//
// @param headers - UpdateGroupAvatarHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupAvatarResponse
func (client *Client) UpdateGroupAvatarWithOptions(request *UpdateGroupAvatarRequest, headers *UpdateGroupAvatarHeaders, runtime *util.RuntimeOptions) (_result *UpdateGroupAvatarResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupAvatar)) {
		body["groupAvatar"] = request.GroupAvatar
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
		Action:      tea.String("UpdateGroupAvatar"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/interconnections/groups/avatars"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGroupAvatarResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改群头像
//
// @param request - UpdateGroupAvatarRequest
//
// @return UpdateGroupAvatarResponse
func (client *Client) UpdateGroupAvatar(request *UpdateGroupAvatarRequest) (_result *UpdateGroupAvatarResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateGroupAvatarHeaders{}
	_result = &UpdateGroupAvatarResponse{}
	_body, _err := client.UpdateGroupAvatarWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改群名称
//
// @param request - UpdateGroupNameRequest
//
// @param headers - UpdateGroupNameHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupNameResponse
func (client *Client) UpdateGroupNameWithOptions(request *UpdateGroupNameRequest, headers *UpdateGroupNameHeaders, runtime *util.RuntimeOptions) (_result *UpdateGroupNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
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
		Action:      tea.String("UpdateGroupName"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/interconnections/groups/names"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGroupNameResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改群名称
//
// @param request - UpdateGroupNameRequest
//
// @return UpdateGroupNameResponse
func (client *Client) UpdateGroupName(request *UpdateGroupNameRequest) (_result *UpdateGroupNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateGroupNameHeaders{}
	_result = &UpdateGroupNameResponse{}
	_body, _err := client.UpdateGroupNameWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置场景群权限项
//
// @param request - UpdateGroupPermissionRequest
//
// @param headers - UpdateGroupPermissionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupPermissionResponse
func (client *Client) UpdateGroupPermissionWithOptions(request *UpdateGroupPermissionRequest, headers *UpdateGroupPermissionHeaders, runtime *util.RuntimeOptions) (_result *UpdateGroupPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.PermissionGroup)) {
		body["permissionGroup"] = request.PermissionGroup
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
		Action:      tea.String("UpdateGroupPermission"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/sceneGroups/permissions"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGroupPermissionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置场景群权限项
//
// @param request - UpdateGroupPermissionRequest
//
// @return UpdateGroupPermissionResponse
func (client *Client) UpdateGroupPermission(request *UpdateGroupPermissionRequest) (_result *UpdateGroupPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateGroupPermissionHeaders{}
	_result = &UpdateGroupPermissionResponse{}
	_body, _err := client.UpdateGroupPermissionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新群管理员
//
// @param request - UpdateGroupSubAdminRequest
//
// @param headers - UpdateGroupSubAdminHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupSubAdminResponse
func (client *Client) UpdateGroupSubAdminWithOptions(request *UpdateGroupSubAdminRequest, headers *UpdateGroupSubAdminHeaders, runtime *util.RuntimeOptions) (_result *UpdateGroupSubAdminResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		body["role"] = request.Role
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
		Action:      tea.String("UpdateGroupSubAdmin"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/sceneGroups/subAdmins"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGroupSubAdminResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新群管理员
//
// @param request - UpdateGroupSubAdminRequest
//
// @return UpdateGroupSubAdminResponse
func (client *Client) UpdateGroupSubAdmin(request *UpdateGroupSubAdminRequest) (_result *UpdateGroupSubAdminResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateGroupSubAdminHeaders{}
	_result = &UpdateGroupSubAdminResponse{}
	_body, _err := client.UpdateGroupSubAdminWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新可交互式动态卡片
//
// @param request - UpdateInteractiveCardRequest
//
// @param headers - UpdateInteractiveCardHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInteractiveCardResponse
func (client *Client) UpdateInteractiveCardWithOptions(request *UpdateInteractiveCardRequest, headers *UpdateInteractiveCardHeaders, runtime *util.RuntimeOptions) (_result *UpdateInteractiveCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CardData)) {
		body["cardData"] = request.CardData
	}

	if !tea.BoolValue(util.IsUnset(request.CardOptions)) {
		body["cardOptions"] = request.CardOptions
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
		Action:      tea.String("UpdateInteractiveCard"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/interactiveCards"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInteractiveCardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新可交互式动态卡片
//
// @param request - UpdateInteractiveCardRequest
//
// @return UpdateInteractiveCardResponse
func (client *Client) UpdateInteractiveCard(request *UpdateInteractiveCardRequest) (_result *UpdateInteractiveCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateInteractiveCardHeaders{}
	_result = &UpdateInteractiveCardResponse{}
	_body, _err := client.UpdateInteractiveCardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置群成员禁言状态
//
// @param request - UpdateMemberBanWordsRequest
//
// @param headers - UpdateMemberBanWordsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMemberBanWordsResponse
func (client *Client) UpdateMemberBanWordsWithOptions(request *UpdateMemberBanWordsRequest, headers *UpdateMemberBanWordsHeaders, runtime *util.RuntimeOptions) (_result *UpdateMemberBanWordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MuteDuration)) {
		body["muteDuration"] = request.MuteDuration
	}

	if !tea.BoolValue(util.IsUnset(request.MuteStatus)) {
		body["muteStatus"] = request.MuteStatus
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
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
		Action:      tea.String("UpdateMemberBanWords"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/sceneGroups/muteMembers/set"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateMemberBanWordsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置群成员禁言状态
//
// @param request - UpdateMemberBanWordsRequest
//
// @return UpdateMemberBanWordsResponse
func (client *Client) UpdateMemberBanWords(request *UpdateMemberBanWordsRequest) (_result *UpdateMemberBanWordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateMemberBanWordsHeaders{}
	_result = &UpdateMemberBanWordsResponse{}
	_body, _err := client.UpdateMemberBanWordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新群成员的群昵称
//
// @param request - UpdateMemberGroupNickRequest
//
// @param headers - UpdateMemberGroupNickHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMemberGroupNickResponse
func (client *Client) UpdateMemberGroupNickWithOptions(request *UpdateMemberGroupNickRequest, headers *UpdateMemberGroupNickHeaders, runtime *util.RuntimeOptions) (_result *UpdateMemberGroupNickResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupNick)) {
		body["groupNick"] = request.GroupNick
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
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
		Action:      tea.String("UpdateMemberGroupNick"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/sceneGroups/members/groupNicks"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMemberGroupNickResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新群成员的群昵称
//
// @param request - UpdateMemberGroupNickRequest
//
// @return UpdateMemberGroupNickResponse
func (client *Client) UpdateMemberGroupNick(request *UpdateMemberGroupNickRequest) (_result *UpdateMemberGroupNickResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateMemberGroupNickHeaders{}
	_result = &UpdateMemberGroupNickResponse{}
	_body, _err := client.UpdateMemberGroupNickWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改组织里的机器人
//
// @param request - UpdateRobotInOrgRequest
//
// @param headers - UpdateRobotInOrgHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRobotInOrgResponse
func (client *Client) UpdateRobotInOrgWithOptions(request *UpdateRobotInOrgRequest, headers *UpdateRobotInOrgHeaders, runtime *util.RuntimeOptions) (_result *UpdateRobotInOrgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Brief)) {
		body["brief"] = request.Brief
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		body["icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OutgoingToken)) {
		body["outgoingToken"] = request.OutgoingToken
	}

	if !tea.BoolValue(util.IsUnset(request.OutgoingUrl)) {
		body["outgoingUrl"] = request.OutgoingUrl
	}

	if !tea.BoolValue(util.IsUnset(request.PreviewMediaId)) {
		body["previewMediaId"] = request.PreviewMediaId
	}

	if !tea.BoolValue(util.IsUnset(request.RobotCode)) {
		body["robotCode"] = request.RobotCode
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
		Action:      tea.String("UpdateRobotInOrg"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/organizations/robots"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRobotInOrgResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改组织里的机器人
//
// @param request - UpdateRobotInOrgRequest
//
// @return UpdateRobotInOrgResponse
func (client *Client) UpdateRobotInOrg(request *UpdateRobotInOrgRequest) (_result *UpdateRobotInOrgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateRobotInOrgHeaders{}
	_result = &UpdateRobotInOrgResponse{}
	_body, _err := client.UpdateRobotInOrgWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 机器人更新可交互式卡片(个人、企业)
//
// @param request - UpdateRobotInteractiveCardRequest
//
// @param headers - UpdateRobotInteractiveCardHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRobotInteractiveCardResponse
func (client *Client) UpdateRobotInteractiveCardWithOptions(request *UpdateRobotInteractiveCardRequest, headers *UpdateRobotInteractiveCardHeaders, runtime *util.RuntimeOptions) (_result *UpdateRobotInteractiveCardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CardBizId)) {
		body["cardBizId"] = request.CardBizId
	}

	if !tea.BoolValue(util.IsUnset(request.CardData)) {
		body["cardData"] = request.CardData
	}

	if !tea.BoolValue(util.IsUnset(request.UnionIdPrivateDataMap)) {
		body["unionIdPrivateDataMap"] = request.UnionIdPrivateDataMap
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateOptions)) {
		body["updateOptions"] = request.UpdateOptions
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdPrivateDataMap)) {
		body["userIdPrivateDataMap"] = request.UserIdPrivateDataMap
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
		Action:      tea.String("UpdateRobotInteractiveCard"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/robots/interactiveCards"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRobotInteractiveCardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 机器人更新可交互式卡片(个人、企业)
//
// @param request - UpdateRobotInteractiveCardRequest
//
// @return UpdateRobotInteractiveCardResponse
func (client *Client) UpdateRobotInteractiveCard(request *UpdateRobotInteractiveCardRequest) (_result *UpdateRobotInteractiveCardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateRobotInteractiveCardHeaders{}
	_result = &UpdateRobotInteractiveCardResponse{}
	_body, _err := client.UpdateRobotInteractiveCardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改场景群模板消息存档能力开启状态
//
// @param request - UpdateSceneGroupTemplateMessageOpenStatusRequest
//
// @param headers - UpdateSceneGroupTemplateMessageOpenStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSceneGroupTemplateMessageOpenStatusResponse
func (client *Client) UpdateSceneGroupTemplateMessageOpenStatusWithOptions(request *UpdateSceneGroupTemplateMessageOpenStatusRequest, headers *UpdateSceneGroupTemplateMessageOpenStatusHeaders, runtime *util.RuntimeOptions) (_result *UpdateSceneGroupTemplateMessageOpenStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateIdList)) {
		body["templateIdList"] = request.TemplateIdList
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
		Action:      tea.String("UpdateSceneGroupTemplateMessageOpenStatus"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/chats/sceneGroups/templates/messageOpenStatuses"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSceneGroupTemplateMessageOpenStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改场景群模板消息存档能力开启状态
//
// @param request - UpdateSceneGroupTemplateMessageOpenStatusRequest
//
// @return UpdateSceneGroupTemplateMessageOpenStatusResponse
func (client *Client) UpdateSceneGroupTemplateMessageOpenStatus(request *UpdateSceneGroupTemplateMessageOpenStatusRequest) (_result *UpdateSceneGroupTemplateMessageOpenStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateSceneGroupTemplateMessageOpenStatusHeaders{}
	_result = &UpdateSceneGroupTemplateMessageOpenStatusResponse{}
	_body, _err := client.UpdateSceneGroupTemplateMessageOpenStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置群成员的群角色
//
// @param request - UpdateTheGroupRolesOfGroupMemberRequest
//
// @param headers - UpdateTheGroupRolesOfGroupMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTheGroupRolesOfGroupMemberResponse
func (client *Client) UpdateTheGroupRolesOfGroupMemberWithOptions(request *UpdateTheGroupRolesOfGroupMemberRequest, headers *UpdateTheGroupRolesOfGroupMemberHeaders, runtime *util.RuntimeOptions) (_result *UpdateTheGroupRolesOfGroupMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenRoleIds)) {
		body["openRoleIds"] = request.OpenRoleIds
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
		Action:      tea.String("UpdateTheGroupRolesOfGroupMember"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/sceneGroups/members/groupRoles"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTheGroupRolesOfGroupMemberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置群成员的群角色
//
// @param request - UpdateTheGroupRolesOfGroupMemberRequest
//
// @return UpdateTheGroupRolesOfGroupMemberResponse
func (client *Client) UpdateTheGroupRolesOfGroupMember(request *UpdateTheGroupRolesOfGroupMemberRequest) (_result *UpdateTheGroupRolesOfGroupMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateTheGroupRolesOfGroupMemberHeaders{}
	_result = &UpdateTheGroupRolesOfGroupMemberResponse{}
	_body, _err := client.UpdateTheGroupRolesOfGroupMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑链接增强注册规则
//
// @param request - UpdateUnfurlingRegisterRequest
//
// @param headers - UpdateUnfurlingRegisterHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUnfurlingRegisterResponse
func (client *Client) UpdateUnfurlingRegisterWithOptions(request *UpdateUnfurlingRegisterRequest, headers *UpdateUnfurlingRegisterHeaders, runtime *util.RuntimeOptions) (_result *UpdateUnfurlingRegisterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiSecret)) {
		body["apiSecret"] = request.ApiSecret
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackUrl)) {
		body["callbackUrl"] = request.CallbackUrl
	}

	if !tea.BoolValue(util.IsUnset(request.CardTemplateId)) {
		body["cardTemplateId"] = request.CardTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		body["domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		body["path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.RuleDesc)) {
		body["ruleDesc"] = request.RuleDesc
	}

	if !tea.BoolValue(util.IsUnset(request.RuleMatchType)) {
		body["ruleMatchType"] = request.RuleMatchType
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
		Action:      tea.String("UpdateUnfurlingRegister"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/unfurling/rules"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUnfurlingRegisterResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑链接增强注册规则
//
// @param request - UpdateUnfurlingRegisterRequest
//
// @return UpdateUnfurlingRegisterResponse
func (client *Client) UpdateUnfurlingRegister(request *UpdateUnfurlingRegisterRequest) (_result *UpdateUnfurlingRegisterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateUnfurlingRegisterHeaders{}
	_result = &UpdateUnfurlingRegisterResponse{}
	_body, _err := client.UpdateUnfurlingRegisterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 链接增强规则状态更新
//
// @param request - UpdateUnfurlingRegisterStatusRequest
//
// @param headers - UpdateUnfurlingRegisterStatusHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUnfurlingRegisterStatusResponse
func (client *Client) UpdateUnfurlingRegisterStatusWithOptions(request *UpdateUnfurlingRegisterStatusRequest, headers *UpdateUnfurlingRegisterStatusHeaders, runtime *util.RuntimeOptions) (_result *UpdateUnfurlingRegisterStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
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
		Action:      tea.String("UpdateUnfurlingRegisterStatus"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/unfurling/rules/statuses"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUnfurlingRegisterStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 链接增强规则状态更新
//
// @param request - UpdateUnfurlingRegisterStatusRequest
//
// @return UpdateUnfurlingRegisterStatusResponse
func (client *Client) UpdateUnfurlingRegisterStatus(request *UpdateUnfurlingRegisterStatusRequest) (_result *UpdateUnfurlingRegisterStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateUnfurlingRegisterStatusHeaders{}
	_result = &UpdateUnfurlingRegisterStatusResponse{}
	_body, _err := client.UpdateUnfurlingRegisterStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 升级群为外部群
//
// @param request - UpgradeToExternalGroupRequest
//
// @param headers - UpgradeToExternalGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeToExternalGroupResponse
func (client *Client) UpgradeToExternalGroupWithOptions(request *UpgradeToExternalGroupRequest, headers *UpgradeToExternalGroupHeaders, runtime *util.RuntimeOptions) (_result *UpgradeToExternalGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["templateId"] = request.TemplateId
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
		Action:      tea.String("UpgradeToExternalGroup"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/chats/sceneGroups/upgradeToExternalGroup"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeToExternalGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 升级群为外部群
//
// @param request - UpgradeToExternalGroupRequest
//
// @return UpgradeToExternalGroupResponse
func (client *Client) UpgradeToExternalGroup(request *UpgradeToExternalGroupRequest) (_result *UpgradeToExternalGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpgradeToExternalGroupHeaders{}
	_result = &UpgradeToExternalGroupResponse{}
	_body, _err := client.UpgradeToExternalGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 升级为B2C群
//
// @param request - UpgradeToServiceGroupRequest
//
// @param headers - UpgradeToServiceGroupHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeToServiceGroupResponse
func (client *Client) UpgradeToServiceGroupWithOptions(request *UpgradeToServiceGroupRequest, headers *UpgradeToServiceGroupHeaders, runtime *util.RuntimeOptions) (_result *UpgradeToServiceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["templateId"] = request.TemplateId
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
		Action:      tea.String("UpgradeToServiceGroup"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/chats/sceneGroups/upgradeToServiceGroup"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeToServiceGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 升级为B2C群
//
// @param request - UpgradeToServiceGroupRequest
//
// @return UpgradeToServiceGroupResponse
func (client *Client) UpgradeToServiceGroup(request *UpgradeToServiceGroupRequest) (_result *UpgradeToServiceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpgradeToServiceGroupHeaders{}
	_result = &UpgradeToServiceGroupResponse{}
	_body, _err := client.UpgradeToServiceGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加群成员
//
// @param request - AddGroupMemberRequest
//
// @param headers - AddGroupMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGroupMemberResponse
func (client *Client) AddGroupMemberWithOptions(request *AddGroupMemberRequest, headers *AddGroupMemberHeaders, runtime *util.RuntimeOptions) (_result *AddGroupMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUserIds)) {
		body["appUserIds"] = request.AppUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
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
		Action:      tea.String("addGroupMember"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/interconnections/groups/members"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddGroupMemberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加群成员
//
// @param request - AddGroupMemberRequest
//
// @return AddGroupMemberResponse
func (client *Client) AddGroupMember(request *AddGroupMemberRequest) (_result *AddGroupMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddGroupMemberHeaders{}
	_result = &AddGroupMemberResponse{}
	_body, _err := client.AddGroupMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移除群成员
//
// @param request - RemoveGroupMemberRequest
//
// @param headers - RemoveGroupMemberHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveGroupMemberResponse
func (client *Client) RemoveGroupMemberWithOptions(request *RemoveGroupMemberRequest, headers *RemoveGroupMemberHeaders, runtime *util.RuntimeOptions) (_result *RemoveGroupMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUserIds)) {
		body["appUserIds"] = request.AppUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["operatorId"] = request.OperatorId
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
		Action:      tea.String("removeGroupMember"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/interconnections/groups/members/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveGroupMemberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移除群成员
//
// @param request - RemoveGroupMemberRequest
//
// @return RemoveGroupMemberResponse
func (client *Client) RemoveGroupMember(request *RemoveGroupMemberRequest) (_result *RemoveGroupMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveGroupMemberHeaders{}
	_result = &RemoveGroupMemberResponse{}
	_body, _err := client.RemoveGroupMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发送ToC消息
//
// @param request - SendDingMessageRequest
//
// @param headers - SendDingMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendDingMessageResponse
func (client *Client) SendDingMessageWithOptions(request *SendDingMessageRequest, headers *SendDingMessageHeaders, runtime *util.RuntimeOptions) (_result *SendDingMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		body["message"] = request.Message
	}

	if !tea.BoolValue(util.IsUnset(request.MessageType)) {
		body["messageType"] = request.MessageType
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverId)) {
		body["receiverId"] = request.ReceiverId
	}

	if !tea.BoolValue(util.IsUnset(request.SenderId)) {
		body["senderId"] = request.SenderId
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
		Action:      tea.String("sendDingMessage"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/interconnections/dingMessages/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendDingMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送ToC消息
//
// @param request - SendDingMessageRequest
//
// @return SendDingMessageResponse
func (client *Client) SendDingMessage(request *SendDingMessageRequest) (_result *SendDingMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendDingMessageHeaders{}
	_result = &SendDingMessageResponse{}
	_body, _err := client.SendDingMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发送ToB消息
//
// @param request - SendMessageRequest
//
// @param headers - SendMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendMessageResponse
func (client *Client) SendMessageWithOptions(request *SendMessageRequest, headers *SendMessageHeaders, runtime *util.RuntimeOptions) (_result *SendMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Message)) {
		body["message"] = request.Message
	}

	if !tea.BoolValue(util.IsUnset(request.MessageType)) {
		body["messageType"] = request.MessageType
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverId)) {
		body["receiverId"] = request.ReceiverId
	}

	if !tea.BoolValue(util.IsUnset(request.SenderId)) {
		body["senderId"] = request.SenderId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceInfos)) {
		body["sourceInfos"] = request.SourceInfos
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
		Action:      tea.String("sendMessage"),
		Version:     tea.String("im_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/im/interconnections/messages/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送ToB消息
//
// @param request - SendMessageRequest
//
// @return SendMessageResponse
func (client *Client) SendMessage(request *SendMessageRequest) (_result *SendMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendMessageHeaders{}
	_result = &SendMessageResponse{}
	_body, _err := client.SendMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
