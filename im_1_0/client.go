// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package im_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
	BackgroundMediaId         *string `json:"backgroundMediaId,omitempty" xml:"backgroundMediaId,omitempty"`
	BackgroundMediaIdForPanel *string `json:"backgroundMediaIdForPanel,omitempty" xml:"backgroundMediaIdForPanel,omitempty"`
	DeptId                    *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	EmotionName               *string `json:"emotionName,omitempty" xml:"emotionName,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddOrgTextEmotionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Icon               *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	RobotCode          *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddRobotToConversationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AutoOpenDingTalkConnectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	OpenConversationId *string   `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	OpenMessageIds     []*string `json:"openMessageIds,omitempty" xml:"openMessageIds,omitempty" type:"Repeated"`
	UnionId            *string   `json:"unionId,omitempty" xml:"unionId,omitempty"`
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
	OpenMsgId   *string                                                         `json:"openMsgId,omitempty" xml:"openMsgId,omitempty"`
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
	FileName        *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType        *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	MediaId         *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	Size            *string `json:"size,omitempty" xml:"size,omitempty"`
	Url             *string `json:"url,omitempty" xml:"url,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchQueryFamilySchoolMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CoolAppCode        *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	MaxResults         *int64  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken          *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	HasMore       *bool     `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	MemberUserIds []*string `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty" type:"Repeated"`
	NextToken     *string   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Success       *bool     `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchQueryGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Action           *string `json:"action,omitempty" xml:"action,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CardTemplateBuildActionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GroupOwnerId       *string `json:"groupOwnerId,omitempty" xml:"groupOwnerId,omitempty"`
	GroupOwnerType     *int32  `json:"groupOwnerType,omitempty" xml:"groupOwnerType,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ChangeGroupOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ChatIdToOpenConversationIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	OpenConversationId *string   `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	Role               *int32    `json:"role,omitempty" xml:"role,omitempty"`
	UserIds            []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ChatSubAdminUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	UserId             *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckUserIsGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AppUserId       *string `json:"appUserId,omitempty" xml:"appUserId,omitempty"`
	GroupAvatar     *string `json:"groupAvatar,omitempty" xml:"groupAvatar,omitempty"`
	GroupName       *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	GroupOwnerId    *string `json:"groupOwnerId,omitempty" xml:"groupOwnerId,omitempty"`
	GroupTemplateId *string `json:"groupTemplateId,omitempty" xml:"groupTemplateId,omitempty"`
	OperatorId      *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
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
	ConversationId     *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCoupleGroupConversationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AppUserIds      []*string `json:"appUserIds,omitempty" xml:"appUserIds,omitempty" type:"Repeated"`
	GroupAvatar     *string   `json:"groupAvatar,omitempty" xml:"groupAvatar,omitempty"`
	GroupName       *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	GroupOwnerId    *string   `json:"groupOwnerId,omitempty" xml:"groupOwnerId,omitempty"`
	GroupOwnerType  *int32    `json:"groupOwnerType,omitempty" xml:"groupOwnerType,omitempty"`
	GroupTemplateId *string   `json:"groupTemplateId,omitempty" xml:"groupTemplateId,omitempty"`
	OperatorId      *string   `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	UserIds         []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
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
	AppUserIds         []*string `json:"appUserIds,omitempty" xml:"appUserIds,omitempty" type:"Repeated"`
	ConversationId     *string   `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	OpenConversationId *string   `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	UserIds            []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGroupConversationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AppUserAvatar          *string `json:"appUserAvatar,omitempty" xml:"appUserAvatar,omitempty"`
	AppUserAvatarMediaType *int32  `json:"appUserAvatarMediaType,omitempty" xml:"appUserAvatarMediaType,omitempty"`
	AppUserDynamics        *string `json:"appUserDynamics,omitempty" xml:"appUserDynamics,omitempty"`
	AppUserId              *string `json:"appUserId,omitempty" xml:"appUserId,omitempty"`
	AppUserMobile          *string `json:"appUserMobile,omitempty" xml:"appUserMobile,omitempty"`
	AppUserName            *string `json:"appUserName,omitempty" xml:"appUserName,omitempty"`
	ChannelCode            *string `json:"channelCode,omitempty" xml:"channelCode,omitempty"`
	UserId                 *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateInterconnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Features          map[string]*string                                    `json:"features,omitempty" xml:"features,omitempty"`
	GroupName         *string                                               `json:"groupName,omitempty" xml:"groupName,omitempty"`
	GroupOwnerId      *string                                               `json:"groupOwnerId,omitempty" xml:"groupOwnerId,omitempty"`
	Icon              *string                                               `json:"icon,omitempty" xml:"icon,omitempty"`
	ManagementOptions *CreateSceneGroupConversationRequestManagementOptions `json:"managementOptions,omitempty" xml:"managementOptions,omitempty" type:"Struct"`
	TemplateId        *string                                               `json:"templateId,omitempty" xml:"templateId,omitempty"`
	UserIdList        []*string                                             `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
	Uuid              *string                                               `json:"uuid,omitempty" xml:"uuid,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSceneGroupConversationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AppUserId         *string   `json:"appUserId,omitempty" xml:"appUserId,omitempty"`
	BusinessUniqueKey *string   `json:"businessUniqueKey,omitempty" xml:"businessUniqueKey,omitempty"`
	GroupAvatar       *string   `json:"groupAvatar,omitempty" xml:"groupAvatar,omitempty"`
	GroupName         *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	GroupTemplateId   *string   `json:"groupTemplateId,omitempty" xml:"groupTemplateId,omitempty"`
	OperatorId        *string   `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	UserIds           []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
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
	ConversationId     *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateStoreGroupConversationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DeptId     *int64    `json:"deptId,omitempty" xml:"deptId,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteOrgTextEmotionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DismissGroupConversationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AppUserId          *string `json:"appUserId,omitempty" xml:"appUserId,omitempty"`
	ChannelCode        *string `json:"channelCode,omitempty" xml:"channelCode,omitempty"`
	DeviceId           *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	UserId             *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetConversationUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	MaxResults         *int32   `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	MsgTypes           []*int32 `json:"msgTypes,omitempty" xml:"msgTypes,omitempty" type:"Repeated"`
	NextToken          *int64   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OpenConversationId *string  `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	UnionId            *string  `json:"unionId,omitempty" xml:"unionId,omitempty"`
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
	CorpId             *string                                               `json:"corpId,omitempty" xml:"corpId,omitempty"`
	HasMore            *string                                               `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	Messages           []*GetFamilySchoolConversationMsgResponseBodyMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	NextToken          *string                                               `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OpenConversationId *string                                               `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
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
	OpenMsgId   *string                                                          `json:"openMsgId,omitempty" xml:"openMsgId,omitempty"`
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
	FileName        *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType        *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	MediaId         *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	Size            *string `json:"size,omitempty" xml:"size,omitempty"`
	Url             *string `json:"url,omitempty" xml:"url,omitempty"`
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFamilySchoolConversationMsgResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int64  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	UnionId    *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
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
	HasMore       *string                                                  `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken     *string                                                  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	CorpId             *string   `json:"corpId,omitempty" xml:"corpId,omitempty"`
	DeptNameChain      []*string `json:"deptNameChain,omitempty" xml:"deptNameChain,omitempty" type:"Repeated"`
	GroupName          *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	GroupType          *string   `json:"groupType,omitempty" xml:"groupType,omitempty"`
	JoinGroupTime      *int64    `json:"joinGroupTime,omitempty" xml:"joinGroupTime,omitempty"`
	OpenConversationId *string   `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFamilySchoolConversationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	MaxResults         *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken          *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	UserId             *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	HasMore   *bool     `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInnerGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AppUserAvatar       *string `json:"appUserAvatar,omitempty" xml:"appUserAvatar,omitempty"`
	AppUserAvatarType   *int32  `json:"appUserAvatarType,omitempty" xml:"appUserAvatarType,omitempty"`
	AppUserId           *string `json:"appUserId,omitempty" xml:"appUserId,omitempty"`
	AppUserMobileNumber *string `json:"appUserMobileNumber,omitempty" xml:"appUserMobileNumber,omitempty"`
	AppUserName         *string `json:"appUserName,omitempty" xml:"appUserName,omitempty"`
	MsgPageType         *int32  `json:"msgPageType,omitempty" xml:"msgPageType,omitempty"`
	QrCode              *string `json:"qrCode,omitempty" xml:"qrCode,omitempty"`
	Signature           *string `json:"signature,omitempty" xml:"signature,omitempty"`
	SourceCode          *string `json:"sourceCode,omitempty" xml:"sourceCode,omitempty"`
	SourceType          *int32  `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	UserId              *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInterconnectionUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Icon               *string `json:"icon,omitempty" xml:"icon,omitempty"`
	MemberAmount       *string `json:"memberAmount,omitempty" xml:"memberAmount,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	Title              *string `json:"title,omitempty" xml:"title,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetNewestInnerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CoolAppCode        *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
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
	GroupUrl           *string `json:"groupUrl,omitempty" xml:"groupUrl,omitempty"`
	Icon               *string `json:"icon,omitempty" xml:"icon,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSceneGroupInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CoolAppCode        *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	Cursor             *string `json:"cursor,omitempty" xml:"cursor,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	Size               *int64  `json:"size,omitempty" xml:"size,omitempty"`
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
	HasMore       *bool     `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	MemberUserIds []*string `json:"memberUserIds,omitempty" xml:"memberUserIds,omitempty" type:"Repeated"`
	NextCursor    *string   `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	Success       *bool     `json:"success,omitempty" xml:"success,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSceneGroupMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	BanWordsMode       *int32                 `json:"banWordsMode,omitempty" xml:"banWordsMode,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	EffectiveDuration  *string                `json:"effectiveDuration,omitempty" xml:"effectiveDuration,omitempty"`
	OpenConversationId *string                `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	Operator           *string                `json:"operator,omitempty" xml:"operator,omitempty"`
	Options            map[string]interface{} `json:"options,omitempty" xml:"options,omitempty"`
	TargetCapacity     *int32                 `json:"targetCapacity,omitempty" xml:"targetCapacity,omitempty"`
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
	ActualPrice        *int64                 `json:"actualPrice,omitempty" xml:"actualPrice,omitempty"`
	CreatedAt          *int64                 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CurrentCapacity    *int32                 `json:"currentCapacity,omitempty" xml:"currentCapacity,omitempty"`
	CurrentEffectUntil *int64                 `json:"currentEffectUntil,omitempty" xml:"currentEffectUntil,omitempty"`
	Discount           *int32                 `json:"discount,omitempty" xml:"discount,omitempty"`
	ExtInfo            map[string]interface{} `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	GroupOwner         *string                `json:"groupOwner,omitempty" xml:"groupOwner,omitempty"`
	GroupTitle         *string                `json:"groupTitle,omitempty" xml:"groupTitle,omitempty"`
	MarkedPrice        *int64                 `json:"markedPrice,omitempty" xml:"markedPrice,omitempty"`
	MemberCount        *int32                 `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	OpenConversationId *string                `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	Operator           *string                `json:"operator,omitempty" xml:"operator,omitempty"`
	TargetCapacity     *int32                 `json:"targetCapacity,omitempty" xml:"targetCapacity,omitempty"`
	TargetEffectUntil  *int64                 `json:"targetEffectUntil,omitempty" xml:"targetEffectUntil,omitempty"`
	Token              *string                `json:"token,omitempty" xml:"token,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GroupCapacityInquiryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	OrderId  *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GroupCapacityOrderConfirmResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ActualPrice        *int64                 `json:"actualPrice,omitempty" xml:"actualPrice,omitempty"`
	CurrentCapacity    *int32                 `json:"currentCapacity,omitempty" xml:"currentCapacity,omitempty"`
	CurrentEffectUntil *int64                 `json:"currentEffectUntil,omitempty" xml:"currentEffectUntil,omitempty"`
	Discount           *int32                 `json:"discount,omitempty" xml:"discount,omitempty"`
	ExtInfo            map[string]interface{} `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	MarkedPrice        *int64                 `json:"markedPrice,omitempty" xml:"markedPrice,omitempty"`
	OpenConversationId *string                `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	Operator           *string                `json:"operator,omitempty" xml:"operator,omitempty"`
	TargetCapacity     *int32                 `json:"targetCapacity,omitempty" xml:"targetCapacity,omitempty"`
	TargetEffectUntil  *int64                 `json:"targetEffectUntil,omitempty" xml:"targetEffectUntil,omitempty"`
	Token              *string                `json:"token,omitempty" xml:"token,omitempty"`
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
	ActualPrice        *int64             `json:"actualPrice,omitempty" xml:"actualPrice,omitempty"`
	CurrentCapacity    *int32             `json:"currentCapacity,omitempty" xml:"currentCapacity,omitempty"`
	CurrentEffectUntil *int64             `json:"currentEffectUntil,omitempty" xml:"currentEffectUntil,omitempty"`
	Discount           *int32             `json:"discount,omitempty" xml:"discount,omitempty"`
	ExtInfo            map[string]*string `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	MarkedPrice        *int64             `json:"markedPrice,omitempty" xml:"markedPrice,omitempty"`
	OpenConversationId *string            `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	Operator           *string            `json:"operator,omitempty" xml:"operator,omitempty"`
	OrderId            *string            `json:"orderId,omitempty" xml:"orderId,omitempty"`
	TargetCapacity     *int32             `json:"targetCapacity,omitempty" xml:"targetCapacity,omitempty"`
	TargetEffectUntil  *int64             `json:"targetEffectUntil,omitempty" xml:"targetEffectUntil,omitempty"`
	Token              *string            `json:"token,omitempty" xml:"token,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GroupCapacityOrderPlaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CreatedAfter       *int64    `json:"createdAfter,omitempty" xml:"createdAfter,omitempty"`
	GroupId            *string   `json:"groupId,omitempty" xml:"groupId,omitempty"`
	GroupMemberSamples []*string `json:"groupMemberSamples,omitempty" xml:"groupMemberSamples,omitempty" type:"Repeated"`
	GroupOwner         *string   `json:"groupOwner,omitempty" xml:"groupOwner,omitempty"`
	GroupTitleKeywords []*string `json:"groupTitleKeywords,omitempty" xml:"groupTitleKeywords,omitempty" type:"Repeated"`
	GroupUrl           *string   `json:"groupUrl,omitempty" xml:"groupUrl,omitempty"`
	MaxResults         *int32    `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	MembersOver        *int32    `json:"membersOver,omitempty" xml:"membersOver,omitempty"`
	NextToken          *string   `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OpenConversationId *string   `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
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
	HasMore       *bool                                        `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken     *string                                      `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	BanWordsMode       *int32                 `json:"banWordsMode,omitempty" xml:"banWordsMode,omitempty"`
	Capacity           *int32                 `json:"capacity,omitempty" xml:"capacity,omitempty"`
	CreatedAt          *int64                 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	ExtInfo            map[string]interface{} `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	GroupAdminList     []*string              `json:"groupAdminList,omitempty" xml:"groupAdminList,omitempty" type:"Repeated"`
	GroupOwner         *string                `json:"groupOwner,omitempty" xml:"groupOwner,omitempty"`
	GroupTitle         *string                `json:"groupTitle,omitempty" xml:"groupTitle,omitempty"`
	MemberCount        *int32                 `json:"memberCount,omitempty" xml:"memberCount,omitempty"`
	OpenConversationId *string                `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	Type               *string                `json:"type,omitempty" xml:"type,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GroupManageQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CapacityLimit      *int32                 `json:"capacityLimit,omitempty" xml:"capacityLimit,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	Brief          *string `json:"brief,omitempty" xml:"brief,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	Icon           *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	OutgoingToken  *string `json:"outgoingToken,omitempty" xml:"outgoingToken,omitempty"`
	OutgoingUrl    *string `json:"outgoingUrl,omitempty" xml:"outgoingUrl,omitempty"`
	PreviewMediaId *string `json:"previewMediaId,omitempty" xml:"previewMediaId,omitempty"`
	RobotCode      *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InstallRobotToOrgResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CallbackRouteKey   *string                                       `json:"callbackRouteKey,omitempty" xml:"callbackRouteKey,omitempty"`
	CardData           *InteractiveCardCreateInstanceRequestCardData `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	CardTemplateId     *string                                       `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	ChatBotId          *string                                       `json:"chatBotId,omitempty" xml:"chatBotId,omitempty"`
	ConversationType   *int32                                        `json:"conversationType,omitempty" xml:"conversationType,omitempty"`
	OpenConversationId *string                                       `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	OutTrackId         *string                                       `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	PrivateData        map[string]*PrivateDataValue                  `json:"privateData,omitempty" xml:"privateData,omitempty"`
	PullStrategy       *bool                                         `json:"pullStrategy,omitempty" xml:"pullStrategy,omitempty"`
	ReceiverUserIdList []*string                                     `json:"receiverUserIdList,omitempty" xml:"receiverUserIdList,omitempty" type:"Repeated"`
	RobotCode          *string                                       `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	UserIdType         *int32                                        `json:"userIdType,omitempty" xml:"userIdType,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InteractiveCardCreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	BackgroundMediaId         *string `json:"backgroundMediaId,omitempty" xml:"backgroundMediaId,omitempty"`
	BackgroundMediaIdForPanel *string `json:"backgroundMediaIdForPanel,omitempty" xml:"backgroundMediaIdForPanel,omitempty"`
	DeptId                    *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	EmotionId                 *string `json:"emotionId,omitempty" xml:"emotionId,omitempty"`
	EmotionName               *string `json:"emotionName,omitempty" xml:"emotionName,omitempty"`
	Status                    *int32  `json:"status,omitempty" xml:"status,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListOrgTextEmotionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CoolAppCode        *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryGroupInfoByMemberAuthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GroupMembers       []*QueryGroupMemberResponseBodyGroupMembers `json:"groupMembers,omitempty" xml:"groupMembers,omitempty" type:"Repeated"`
	OpenConversationId *string                                     `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
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
	GroupMemberAvatar   *string `json:"groupMemberAvatar,omitempty" xml:"groupMemberAvatar,omitempty"`
	GroupMemberDynamics *string `json:"groupMemberDynamics,omitempty" xml:"groupMemberDynamics,omitempty"`
	GroupMemberId       *string `json:"groupMemberId,omitempty" xml:"groupMemberId,omitempty"`
	GroupMemberName     *string `json:"groupMemberName,omitempty" xml:"groupMemberName,omitempty"`
	GroupMemberType     *int32  `json:"groupMemberType,omitempty" xml:"groupMemberType,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CoolAppCode        *string `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
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
	GroupNickName   *string `json:"groupNickName,omitempty" xml:"groupNickName,omitempty"`
	OrgName         *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
	ProfilePhotoUrl *string `json:"profilePhotoUrl,omitempty" xml:"profilePhotoUrl,omitempty"`
	UserId          *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryGroupMemberByMemberAuthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	UserId             *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	GroupMuteMode  *bool                                           `json:"groupMuteMode,omitempty" xml:"groupMuteMode,omitempty"`
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
	MuteEndTime   *int64 `json:"muteEndTime,omitempty" xml:"muteEndTime,omitempty"`
	MuteStartTime *int64 `json:"muteStartTime,omitempty" xml:"muteStartTime,omitempty"`
	UserMuteMode  *bool  `json:"userMuteMode,omitempty" xml:"userMuteMode,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryGroupMuteStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	OpenRoleId         *string `json:"openRoleId,omitempty" xml:"openRoleId,omitempty"`
	Timestamp          *int64  `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMembersOfGroupRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	RobotCode          *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QuerySceneGroupTemplateRobotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GroupMembers    []*QuerySingleGroupRequestGroupMembers `json:"groupMembers,omitempty" xml:"groupMembers,omitempty" type:"Repeated"`
	GroupTemplateId *string                                `json:"groupTemplateId,omitempty" xml:"groupTemplateId,omitempty"`
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
	AppUserId *string `json:"appUserId,omitempty" xml:"appUserId,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	AppUserId          *string `json:"appUserId,omitempty" xml:"appUserId,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	UserId             *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QuerySingleGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AppUserId           *string   `json:"appUserId,omitempty" xml:"appUserId,omitempty"`
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
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	UnReadCount        *int64  `json:"unReadCount,omitempty" xml:"unReadCount,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUnReadMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ChatBotUserId      *string `json:"chatBotUserId,omitempty" xml:"chatBotUserId,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveRobotFromConversationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	SearchKey  *string `json:"searchKey,omitempty" xml:"searchKey,omitempty"`
	UserId     *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Icon               *string `json:"icon,omitempty" xml:"icon,omitempty"`
	MemberAmount       *string `json:"memberAmount,omitempty" xml:"memberAmount,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	Title              *string `json:"title,omitempty" xml:"title,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchInnerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AtOpenIds          map[string]*string                     `json:"atOpenIds,omitempty" xml:"atOpenIds,omitempty"`
	CallbackRouteKey   *string                                `json:"callbackRouteKey,omitempty" xml:"callbackRouteKey,omitempty"`
	CardData           *SendInteractiveCardRequestCardData    `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	CardOptions        *SendInteractiveCardRequestCardOptions `json:"cardOptions,omitempty" xml:"cardOptions,omitempty" type:"Struct"`
	CardTemplateId     *string                                `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	ChatBotId          *string                                `json:"chatBotId,omitempty" xml:"chatBotId,omitempty"`
	ConversationType   *int32                                 `json:"conversationType,omitempty" xml:"conversationType,omitempty"`
	OpenConversationId *string                                `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	OutTrackId         *string                                `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	PrivateData        map[string]*PrivateDataValue           `json:"privateData,omitempty" xml:"privateData,omitempty"`
	PullStrategy       *bool                                  `json:"pullStrategy,omitempty" xml:"pullStrategy,omitempty"`
	ReceiverUserIdList []*string                              `json:"receiverUserIdList,omitempty" xml:"receiverUserIdList,omitempty" type:"Repeated"`
	RobotCode          *string                                `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	UserIdType         *int32                                 `json:"userIdType,omitempty" xml:"userIdType,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendInteractiveCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AtOpenIds          map[string]*string                        `json:"atOpenIds,omitempty" xml:"atOpenIds,omitempty"`
	CallbackRouteKey   *string                                   `json:"callbackRouteKey,omitempty" xml:"callbackRouteKey,omitempty"`
	CardData           *SendOTOInteractiveCardRequestCardData    `json:"cardData,omitempty" xml:"cardData,omitempty" type:"Struct"`
	CardOptions        *SendOTOInteractiveCardRequestCardOptions `json:"cardOptions,omitempty" xml:"cardOptions,omitempty" type:"Struct"`
	CardTemplateId     *string                                   `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	OpenConversationId *string                                   `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	OutTrackId         *string                                   `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	PrivateData        map[string]*PrivateDataValue              `json:"privateData,omitempty" xml:"privateData,omitempty"`
	PullStrategy       *bool                                     `json:"pullStrategy,omitempty" xml:"pullStrategy,omitempty"`
	ReceiverUserIdList []*string                                 `json:"receiverUserIdList,omitempty" xml:"receiverUserIdList,omitempty" type:"Repeated"`
	RobotCode          *string                                   `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	UserIdType         *int32                                    `json:"userIdType,omitempty" xml:"userIdType,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendOTOInteractiveCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CallbackUrl           *string                                     `json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`
	CardBizId             *string                                     `json:"cardBizId,omitempty" xml:"cardBizId,omitempty"`
	CardData              *string                                     `json:"cardData,omitempty" xml:"cardData,omitempty"`
	CardTemplateId        *string                                     `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	OpenConversationId    *string                                     `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	PullStrategy          *bool                                       `json:"pullStrategy,omitempty" xml:"pullStrategy,omitempty"`
	RobotCode             *string                                     `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	SendOptions           *SendRobotInteractiveCardRequestSendOptions `json:"sendOptions,omitempty" xml:"sendOptions,omitempty" type:"Struct"`
	SingleChatReceiver    *string                                     `json:"singleChatReceiver,omitempty" xml:"singleChatReceiver,omitempty"`
	UnionIdPrivateDataMap *string                                     `json:"unionIdPrivateDataMap,omitempty" xml:"unionIdPrivateDataMap,omitempty"`
	UserIdPrivateDataMap  *string                                     `json:"userIdPrivateDataMap,omitempty" xml:"userIdPrivateDataMap,omitempty"`
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
	AtAll            *bool   `json:"atAll,omitempty" xml:"atAll,omitempty"`
	AtUserListJson   *string `json:"atUserListJson,omitempty" xml:"atUserListJson,omitempty"`
	CardPropertyJson *string `json:"cardPropertyJson,omitempty" xml:"cardPropertyJson,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendRobotInteractiveCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AtAll               *bool     `json:"atAll,omitempty" xml:"atAll,omitempty"`
	AtAppUserId         *string   `json:"atAppUserId,omitempty" xml:"atAppUserId,omitempty"`
	AtDingUserId        *string   `json:"atDingUserId,omitempty" xml:"atDingUserId,omitempty"`
	MsgContent          *string   `json:"msgContent,omitempty" xml:"msgContent,omitempty"`
	MsgType             *string   `json:"msgType,omitempty" xml:"msgType,omitempty"`
	OpenConversationIds []*string `json:"openConversationIds,omitempty" xml:"openConversationIds,omitempty" type:"Repeated"`
	RobotCode           *string   `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendRobotMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CallbackUrl        *string                                        `json:"callbackUrl,omitempty" xml:"callbackUrl,omitempty"`
	CardData           *string                                        `json:"cardData,omitempty" xml:"cardData,omitempty"`
	CardTemplateId     *string                                        `json:"cardTemplateId,omitempty" xml:"cardTemplateId,omitempty"`
	OpenConversationId *string                                        `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	OutTrackId         *string                                        `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	RobotCode          *string                                        `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
	SendOptions        *SendTemplateInteractiveCardRequestSendOptions `json:"sendOptions,omitempty" xml:"sendOptions,omitempty" type:"Struct"`
	SingleChatReceiver *string                                        `json:"singleChatReceiver,omitempty" xml:"singleChatReceiver,omitempty"`
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
	AtAll            *bool   `json:"atAll,omitempty" xml:"atAll,omitempty"`
	AtUserListJson   *string `json:"atUserListJson,omitempty" xml:"atUserListJson,omitempty"`
	CardPropertyJson *string `json:"cardPropertyJson,omitempty" xml:"cardPropertyJson,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendTemplateInteractiveCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	OpenConversationId       *string                           `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	RightPanelClosePermitted *bool                             `json:"rightPanelClosePermitted,omitempty" xml:"rightPanelClosePermitted,omitempty"`
	RightPanelOpenStatus     *int32                            `json:"rightPanelOpenStatus,omitempty" xml:"rightPanelOpenStatus,omitempty"`
	Title                    *string                           `json:"title,omitempty" xml:"title,omitempty"`
	WebWndParams             *SetRightPanelRequestWebWndParams `json:"webWndParams,omitempty" xml:"webWndParams,omitempty" type:"Struct"`
	Width                    *int32                            `json:"width,omitempty" xml:"width,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetRightPanelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ConversationType   *int32    `json:"conversationType,omitempty" xml:"conversationType,omitempty"`
	CoolAppCode        *string   `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	OpenConversationId *string   `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	ConversationType   *int32    `json:"conversationType,omitempty" xml:"conversationType,omitempty"`
	CoolAppCode        *string   `json:"coolAppCode,omitempty" xml:"coolAppCode,omitempty"`
	ExpiredTime        *int64    `json:"expiredTime,omitempty" xml:"expiredTime,omitempty"`
	OpenConversationId *string   `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	OutTrackId         *string   `json:"outTrackId,omitempty" xml:"outTrackId,omitempty"`
	Platforms          *string   `json:"platforms,omitempty" xml:"platforms,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	GroupAvatar        *string `json:"groupAvatar,omitempty" xml:"groupAvatar,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateGroupAvatarResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GroupName          *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateGroupNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateGroupPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	OpenConversationId *string   `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	Role               *int64    `json:"role,omitempty" xml:"role,omitempty"`
	UserIds            []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateGroupSubAdminResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateInteractiveCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	MuteDuration       *int64    `json:"muteDuration,omitempty" xml:"muteDuration,omitempty"`
	MuteStatus         *int32    `json:"muteStatus,omitempty" xml:"muteStatus,omitempty"`
	OpenConversationId *string   `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	UserIdList         []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	GroupNick          *string `json:"groupNick,omitempty" xml:"groupNick,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	UserId             *string `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateMemberGroupNickResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Brief          *string `json:"brief,omitempty" xml:"brief,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	Icon           *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	OutgoingToken  *string `json:"outgoingToken,omitempty" xml:"outgoingToken,omitempty"`
	OutgoingUrl    *string `json:"outgoingUrl,omitempty" xml:"outgoingUrl,omitempty"`
	PreviewMediaId *string `json:"previewMediaId,omitempty" xml:"previewMediaId,omitempty"`
	RobotCode      *string `json:"robotCode,omitempty" xml:"robotCode,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRobotInOrgResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CardBizId             *string                                         `json:"cardBizId,omitempty" xml:"cardBizId,omitempty"`
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRobotInteractiveCardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	OpenConversationId *string   `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	OpenRoleIds        []*string `json:"openRoleIds,omitempty" xml:"openRoleIds,omitempty" type:"Repeated"`
	UserId             *string   `json:"userId,omitempty" xml:"userId,omitempty"`
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
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTheGroupRolesOfGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AppUserIds         []*string `json:"appUserIds,omitempty" xml:"appUserIds,omitempty" type:"Repeated"`
	OpenConversationId *string   `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	OperatorId         *string   `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	UserIds            []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
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
	AppUserIds []*string `json:"appUserIds,omitempty" xml:"appUserIds,omitempty" type:"Repeated"`
	UserIds    []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AppUserIds         []*string `json:"appUserIds,omitempty" xml:"appUserIds,omitempty" type:"Repeated"`
	OpenConversationId *string   `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	OperatorId         *string   `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	UserIds            []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Code               *string `json:"code,omitempty" xml:"code,omitempty"`
	Message            *string `json:"message,omitempty" xml:"message,omitempty"`
	MessageType        *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	ReceiverId         *string `json:"receiverId,omitempty" xml:"receiverId,omitempty"`
	SenderId           *string `json:"senderId,omitempty" xml:"senderId,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendDingMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Message            *string                `json:"message,omitempty" xml:"message,omitempty"`
	MessageType        *string                `json:"messageType,omitempty" xml:"messageType,omitempty"`
	OpenConversationId *string                `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	ReceiverId         *string                `json:"receiverId,omitempty" xml:"receiverId,omitempty"`
	SenderId           *string                `json:"senderId,omitempty" xml:"senderId,omitempty"`
	SourceInfos        map[string]interface{} `json:"sourceInfos,omitempty" xml:"sourceInfos,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
