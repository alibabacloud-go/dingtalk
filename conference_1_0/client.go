// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package conference_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CloseVideoConferenceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CloseVideoConferenceHeaders) String() string {
	return tea.Prettify(s)
}

func (s CloseVideoConferenceHeaders) GoString() string {
	return s.String()
}

func (s *CloseVideoConferenceHeaders) SetCommonHeaders(v map[string]*string) *CloseVideoConferenceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CloseVideoConferenceHeaders) SetXAcsDingtalkAccessToken(v string) *CloseVideoConferenceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CloseVideoConferenceRequest struct {
	// 员工在当前开发者企业账号范围内的唯一标识
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s CloseVideoConferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseVideoConferenceRequest) GoString() string {
	return s.String()
}

func (s *CloseVideoConferenceRequest) SetUnionId(v string) *CloseVideoConferenceRequest {
	s.UnionId = &v
	return s
}

type CloseVideoConferenceResponseBody struct {
	Cause *string `json:"cause,omitempty" xml:"cause,omitempty"`
	Code  *int64  `json:"code,omitempty" xml:"code,omitempty"`
}

func (s CloseVideoConferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseVideoConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *CloseVideoConferenceResponseBody) SetCause(v string) *CloseVideoConferenceResponseBody {
	s.Cause = &v
	return s
}

func (s *CloseVideoConferenceResponseBody) SetCode(v int64) *CloseVideoConferenceResponseBody {
	s.Code = &v
	return s
}

type CloseVideoConferenceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CloseVideoConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloseVideoConferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseVideoConferenceResponse) GoString() string {
	return s.String()
}

func (s *CloseVideoConferenceResponse) SetHeaders(v map[string]*string) *CloseVideoConferenceResponse {
	s.Headers = v
	return s
}

func (s *CloseVideoConferenceResponse) SetBody(v *CloseVideoConferenceResponseBody) *CloseVideoConferenceResponse {
	s.Body = v
	return s
}

type CreateVideoConferenceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateVideoConferenceHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoConferenceHeaders) GoString() string {
	return s.String()
}

func (s *CreateVideoConferenceHeaders) SetCommonHeaders(v map[string]*string) *CreateVideoConferenceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateVideoConferenceHeaders) SetXAcsDingtalkAccessToken(v string) *CreateVideoConferenceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateVideoConferenceRequest struct {
	// 会议主题： 文字，不超过20中文
	ConfTitle *string `json:"confTitle,omitempty" xml:"confTitle,omitempty"`
	// 是否邀请主叫
	InviteCaller *bool `json:"inviteCaller,omitempty" xml:"inviteCaller,omitempty"`
	// 邀请参会人员UID列表（必须好友或同事）
	InviteUserIds []*string `json:"inviteUserIds,omitempty" xml:"inviteUserIds,omitempty" type:"Repeated"`
	// 会议发起人UID
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateVideoConferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoConferenceRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoConferenceRequest) SetConfTitle(v string) *CreateVideoConferenceRequest {
	s.ConfTitle = &v
	return s
}

func (s *CreateVideoConferenceRequest) SetInviteCaller(v bool) *CreateVideoConferenceRequest {
	s.InviteCaller = &v
	return s
}

func (s *CreateVideoConferenceRequest) SetInviteUserIds(v []*string) *CreateVideoConferenceRequest {
	s.InviteUserIds = v
	return s
}

func (s *CreateVideoConferenceRequest) SetUserId(v string) *CreateVideoConferenceRequest {
	s.UserId = &v
	return s
}

type CreateVideoConferenceResponseBody struct {
	// conferenceId
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// 会议密码
	ConferencePassword *string `json:"conferencePassword,omitempty" xml:"conferencePassword,omitempty"`
	// 入会链接
	ExternalLinkUrl *string `json:"externalLinkUrl,omitempty" xml:"externalLinkUrl,omitempty"`
	// 主持人密码
	HostPassword *string `json:"hostPassword,omitempty" xml:"hostPassword,omitempty"`
	// 电话入会号码
	PhoneNumbers []*string `json:"phoneNumbers,omitempty" xml:"phoneNumbers,omitempty" type:"Repeated"`
	RoomCode     *string   `json:"roomCode,omitempty" xml:"roomCode,omitempty"`
}

func (s CreateVideoConferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVideoConferenceResponseBody) SetConferenceId(v string) *CreateVideoConferenceResponseBody {
	s.ConferenceId = &v
	return s
}

func (s *CreateVideoConferenceResponseBody) SetConferencePassword(v string) *CreateVideoConferenceResponseBody {
	s.ConferencePassword = &v
	return s
}

func (s *CreateVideoConferenceResponseBody) SetExternalLinkUrl(v string) *CreateVideoConferenceResponseBody {
	s.ExternalLinkUrl = &v
	return s
}

func (s *CreateVideoConferenceResponseBody) SetHostPassword(v string) *CreateVideoConferenceResponseBody {
	s.HostPassword = &v
	return s
}

func (s *CreateVideoConferenceResponseBody) SetPhoneNumbers(v []*string) *CreateVideoConferenceResponseBody {
	s.PhoneNumbers = v
	return s
}

func (s *CreateVideoConferenceResponseBody) SetRoomCode(v string) *CreateVideoConferenceResponseBody {
	s.RoomCode = &v
	return s
}

type CreateVideoConferenceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateVideoConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVideoConferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVideoConferenceResponse) GoString() string {
	return s.String()
}

func (s *CreateVideoConferenceResponse) SetHeaders(v map[string]*string) *CreateVideoConferenceResponse {
	s.Headers = v
	return s
}

func (s *CreateVideoConferenceResponse) SetBody(v *CreateVideoConferenceResponseBody) *CreateVideoConferenceResponse {
	s.Body = v
	return s
}

type QueryCloudRecordTextHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCloudRecordTextHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordTextHeaders) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordTextHeaders) SetCommonHeaders(v map[string]*string) *QueryCloudRecordTextHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCloudRecordTextHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCloudRecordTextHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCloudRecordTextRequest struct {
	// 0-向前查询，1-向后查询 。 向前查询：此次查询按照时间由小到大的顺序进行。
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// 单页查询的最大条目数，最多2000
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 游标，第一次查询可为空，之后每次带上一次的游标
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryCloudRecordTextRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordTextRequest) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordTextRequest) SetDirection(v string) *QueryCloudRecordTextRequest {
	s.Direction = &v
	return s
}

func (s *QueryCloudRecordTextRequest) SetMaxResults(v int64) *QueryCloudRecordTextRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryCloudRecordTextRequest) SetNextToken(v int64) *QueryCloudRecordTextRequest {
	s.NextToken = &v
	return s
}

func (s *QueryCloudRecordTextRequest) SetStartTime(v int64) *QueryCloudRecordTextRequest {
	s.StartTime = &v
	return s
}

func (s *QueryCloudRecordTextRequest) SetUnionId(v string) *QueryCloudRecordTextRequest {
	s.UnionId = &v
	return s
}

type QueryCloudRecordTextResponseBody struct {
	// 是否有更多
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 段落列表
	ParagraphList []*QueryCloudRecordTextResponseBodyParagraphList `json:"paragraphList,omitempty" xml:"paragraphList,omitempty" type:"Repeated"`
}

func (s QueryCloudRecordTextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordTextResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordTextResponseBody) SetHasMore(v bool) *QueryCloudRecordTextResponseBody {
	s.HasMore = &v
	return s
}

func (s *QueryCloudRecordTextResponseBody) SetParagraphList(v []*QueryCloudRecordTextResponseBodyParagraphList) *QueryCloudRecordTextResponseBody {
	s.ParagraphList = v
	return s
}

type QueryCloudRecordTextResponseBodyParagraphList struct {
	// 结束时间，毫秒
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 游标，下次查询时使用
	NextTtoken *int64 `json:"nextTtoken,omitempty" xml:"nextTtoken,omitempty"`
	// 发言人昵称
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// 段落内容
	Paragraph *string `json:"paragraph,omitempty" xml:"paragraph,omitempty"`
	// 云录制id
	RecordId *int64 `json:"recordId,omitempty" xml:"recordId,omitempty"`
	// 句子列表
	SentenceList []*QueryCloudRecordTextResponseBodyParagraphListSentenceList `json:"sentenceList,omitempty" xml:"sentenceList,omitempty" type:"Repeated"`
	// 开始时间，毫秒
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 状态，暂不解析
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 发言人unionId
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryCloudRecordTextResponseBodyParagraphList) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordTextResponseBodyParagraphList) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetEndTime(v int64) *QueryCloudRecordTextResponseBodyParagraphList {
	s.EndTime = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetNextTtoken(v int64) *QueryCloudRecordTextResponseBodyParagraphList {
	s.NextTtoken = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetNickName(v string) *QueryCloudRecordTextResponseBodyParagraphList {
	s.NickName = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetParagraph(v string) *QueryCloudRecordTextResponseBodyParagraphList {
	s.Paragraph = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetRecordId(v int64) *QueryCloudRecordTextResponseBodyParagraphList {
	s.RecordId = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetSentenceList(v []*QueryCloudRecordTextResponseBodyParagraphListSentenceList) *QueryCloudRecordTextResponseBodyParagraphList {
	s.SentenceList = v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetStartTime(v int64) *QueryCloudRecordTextResponseBodyParagraphList {
	s.StartTime = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetStatus(v int64) *QueryCloudRecordTextResponseBodyParagraphList {
	s.Status = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphList) SetUnionId(v string) *QueryCloudRecordTextResponseBodyParagraphList {
	s.UnionId = &v
	return s
}

type QueryCloudRecordTextResponseBodyParagraphListSentenceList struct {
	// 结束时间
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 句子
	Sentence *string `json:"sentence,omitempty" xml:"sentence,omitempty"`
	// 开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 用户unionId
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	// 单词列表
	WordList []*QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList `json:"wordList,omitempty" xml:"wordList,omitempty" type:"Repeated"`
}

func (s QueryCloudRecordTextResponseBodyParagraphListSentenceList) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordTextResponseBodyParagraphListSentenceList) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceList) SetEndTime(v int64) *QueryCloudRecordTextResponseBodyParagraphListSentenceList {
	s.EndTime = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceList) SetSentence(v string) *QueryCloudRecordTextResponseBodyParagraphListSentenceList {
	s.Sentence = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceList) SetStartTime(v int64) *QueryCloudRecordTextResponseBodyParagraphListSentenceList {
	s.StartTime = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceList) SetUnionId(v string) *QueryCloudRecordTextResponseBodyParagraphListSentenceList {
	s.UnionId = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceList) SetWordList(v []*QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList) *QueryCloudRecordTextResponseBodyParagraphListSentenceList {
	s.WordList = v
	return s
}

type QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList struct {
	// 结束时间
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 单词
	Word *string `json:"word,omitempty" xml:"word,omitempty"`
	// 单词id
	WordId *string `json:"wordId,omitempty" xml:"wordId,omitempty"`
}

func (s QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList) SetEndTime(v int64) *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList {
	s.EndTime = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList) SetStartTime(v int64) *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList {
	s.StartTime = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList) SetWord(v string) *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList {
	s.Word = &v
	return s
}

func (s *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList) SetWordId(v string) *QueryCloudRecordTextResponseBodyParagraphListSentenceListWordList {
	s.WordId = &v
	return s
}

type QueryCloudRecordTextResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCloudRecordTextResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCloudRecordTextResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordTextResponse) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordTextResponse) SetHeaders(v map[string]*string) *QueryCloudRecordTextResponse {
	s.Headers = v
	return s
}

func (s *QueryCloudRecordTextResponse) SetBody(v *QueryCloudRecordTextResponseBody) *QueryCloudRecordTextResponse {
	s.Body = v
	return s
}

type QueryCloudRecordVideoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCloudRecordVideoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordVideoHeaders) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoHeaders) SetCommonHeaders(v map[string]*string) *QueryCloudRecordVideoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCloudRecordVideoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCloudRecordVideoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCloudRecordVideoRequest struct {
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryCloudRecordVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordVideoRequest) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoRequest) SetUnionId(v string) *QueryCloudRecordVideoRequest {
	s.UnionId = &v
	return s
}

type QueryCloudRecordVideoResponseBody struct {
	// 视频列表
	VideoList []*QueryCloudRecordVideoResponseBodyVideoList `json:"videoList,omitempty" xml:"videoList,omitempty" type:"Repeated"`
}

func (s QueryCloudRecordVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordVideoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoResponseBody) SetVideoList(v []*QueryCloudRecordVideoResponseBodyVideoList) *QueryCloudRecordVideoResponseBody {
	s.VideoList = v
	return s
}

type QueryCloudRecordVideoResponseBodyVideoList struct {
	// 录制持续时间
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 录制结束时间
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 文件大小
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// 媒体文件id，唯一
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	// 音视频云录制Id，多份视频recordId一样
	RecordId *string `json:"recordId,omitempty" xml:"recordId,omitempty"`
	// 记录类型,0-普通录制，1-合成的文件
	RecordType *int64 `json:"recordType,omitempty" xml:"recordType,omitempty"`
	// 媒体文件所在集群id
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// 录制开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 录制人UnionId
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryCloudRecordVideoResponseBodyVideoList) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordVideoResponseBodyVideoList) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetDuration(v int64) *QueryCloudRecordVideoResponseBodyVideoList {
	s.Duration = &v
	return s
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetEndTime(v int64) *QueryCloudRecordVideoResponseBodyVideoList {
	s.EndTime = &v
	return s
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetFileSize(v int64) *QueryCloudRecordVideoResponseBodyVideoList {
	s.FileSize = &v
	return s
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetMediaId(v string) *QueryCloudRecordVideoResponseBodyVideoList {
	s.MediaId = &v
	return s
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetRecordId(v string) *QueryCloudRecordVideoResponseBodyVideoList {
	s.RecordId = &v
	return s
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetRecordType(v int64) *QueryCloudRecordVideoResponseBodyVideoList {
	s.RecordType = &v
	return s
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetRegionId(v string) *QueryCloudRecordVideoResponseBodyVideoList {
	s.RegionId = &v
	return s
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetStartTime(v int64) *QueryCloudRecordVideoResponseBodyVideoList {
	s.StartTime = &v
	return s
}

func (s *QueryCloudRecordVideoResponseBodyVideoList) SetUnionId(v string) *QueryCloudRecordVideoResponseBodyVideoList {
	s.UnionId = &v
	return s
}

type QueryCloudRecordVideoResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCloudRecordVideoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCloudRecordVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordVideoResponse) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoResponse) SetHeaders(v map[string]*string) *QueryCloudRecordVideoResponse {
	s.Headers = v
	return s
}

func (s *QueryCloudRecordVideoResponse) SetBody(v *QueryCloudRecordVideoResponseBody) *QueryCloudRecordVideoResponse {
	s.Body = v
	return s
}

type QueryCloudRecordVideoPlayInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCloudRecordVideoPlayInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordVideoPlayInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoPlayInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryCloudRecordVideoPlayInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCloudRecordVideoPlayInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCloudRecordVideoPlayInfoRequest struct {
	// 媒体文件id
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	// 集群id
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s QueryCloudRecordVideoPlayInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordVideoPlayInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoPlayInfoRequest) SetMediaId(v string) *QueryCloudRecordVideoPlayInfoRequest {
	s.MediaId = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoRequest) SetRegionId(v string) *QueryCloudRecordVideoPlayInfoRequest {
	s.RegionId = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoRequest) SetUnionId(v string) *QueryCloudRecordVideoPlayInfoRequest {
	s.UnionId = &v
	return s
}

type QueryCloudRecordVideoPlayInfoResponseBody struct {
	// 时长
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 大小
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// MP4格式下载链接
	Mp4FileUrl *string `json:"mp4FileUrl,omitempty" xml:"mp4FileUrl,omitempty"`
	// 在线播放链接
	PlayUrl *string `json:"playUrl,omitempty" xml:"playUrl,omitempty"`
	// 状态
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryCloudRecordVideoPlayInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordVideoPlayInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoPlayInfoResponseBody) SetDuration(v int64) *QueryCloudRecordVideoPlayInfoResponseBody {
	s.Duration = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoResponseBody) SetFileSize(v int64) *QueryCloudRecordVideoPlayInfoResponseBody {
	s.FileSize = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoResponseBody) SetMp4FileUrl(v string) *QueryCloudRecordVideoPlayInfoResponseBody {
	s.Mp4FileUrl = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoResponseBody) SetPlayUrl(v string) *QueryCloudRecordVideoPlayInfoResponseBody {
	s.PlayUrl = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoResponseBody) SetStatus(v int64) *QueryCloudRecordVideoPlayInfoResponseBody {
	s.Status = &v
	return s
}

type QueryCloudRecordVideoPlayInfoResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCloudRecordVideoPlayInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCloudRecordVideoPlayInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCloudRecordVideoPlayInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoPlayInfoResponse) SetHeaders(v map[string]*string) *QueryCloudRecordVideoPlayInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoResponse) SetBody(v *QueryCloudRecordVideoPlayInfoResponseBody) *QueryCloudRecordVideoPlayInfoResponse {
	s.Body = v
	return s
}

type QueryConferenceInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryConferenceInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryConferenceInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryConferenceInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryConferenceInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryConferenceInfoResponseBody struct {
	// 会议信息结构体
	ConfInfo *QueryConferenceInfoResponseBodyConfInfo `json:"confInfo,omitempty" xml:"confInfo,omitempty" type:"Struct"`
}

func (s QueryConferenceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoResponseBody) SetConfInfo(v *QueryConferenceInfoResponseBodyConfInfo) *QueryConferenceInfoResponseBody {
	s.ConfInfo = v
	return s
}

type QueryConferenceInfoResponseBodyConfInfo struct {
	// 当前在会人数
	ActiveNum *int32 `json:"activeNum,omitempty" xml:"activeNum,omitempty"`
	// 累积入会人数
	AttendNum *int32 `json:"attendNum,omitempty" xml:"attendNum,omitempty"`
	// 会议时长
	ConfDuration *int64 `json:"confDuration,omitempty" xml:"confDuration,omitempty"`
	// 会议id
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// 会议创建人unionId
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// 会议创建人昵称
	CreatorNick *string `json:"creatorNick,omitempty" xml:"creatorNick,omitempty"`
	// 会议结束时间
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 会议web入会链接
	ExternalLinkUrl *string `json:"externalLinkUrl,omitempty" xml:"externalLinkUrl,omitempty"`
	// 邀请人数
	InvitedNum *int32 `json:"invitedNum,omitempty" xml:"invitedNum,omitempty"`
	// 会议码
	RoomCode *string `json:"roomCode,omitempty" xml:"roomCode,omitempty"`
	// 会议开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 会议状态
	// 0 初始化
	// 1 开始
	// 2 结束
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// 会议标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryConferenceInfoResponseBodyConfInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoResponseBodyConfInfo) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetActiveNum(v int32) *QueryConferenceInfoResponseBodyConfInfo {
	s.ActiveNum = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetAttendNum(v int32) *QueryConferenceInfoResponseBodyConfInfo {
	s.AttendNum = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetConfDuration(v int64) *QueryConferenceInfoResponseBodyConfInfo {
	s.ConfDuration = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetConferenceId(v string) *QueryConferenceInfoResponseBodyConfInfo {
	s.ConferenceId = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetCreatorId(v string) *QueryConferenceInfoResponseBodyConfInfo {
	s.CreatorId = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetCreatorNick(v string) *QueryConferenceInfoResponseBodyConfInfo {
	s.CreatorNick = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetEndTime(v int64) *QueryConferenceInfoResponseBodyConfInfo {
	s.EndTime = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetExternalLinkUrl(v string) *QueryConferenceInfoResponseBodyConfInfo {
	s.ExternalLinkUrl = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetInvitedNum(v int32) *QueryConferenceInfoResponseBodyConfInfo {
	s.InvitedNum = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetRoomCode(v string) *QueryConferenceInfoResponseBodyConfInfo {
	s.RoomCode = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetStartTime(v int64) *QueryConferenceInfoResponseBodyConfInfo {
	s.StartTime = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetStatus(v int32) *QueryConferenceInfoResponseBodyConfInfo {
	s.Status = &v
	return s
}

func (s *QueryConferenceInfoResponseBodyConfInfo) SetTitle(v string) *QueryConferenceInfoResponseBodyConfInfo {
	s.Title = &v
	return s
}

type QueryConferenceInfoResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryConferenceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryConferenceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoResponse) SetHeaders(v map[string]*string) *QueryConferenceInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryConferenceInfoResponse) SetBody(v *QueryConferenceInfoResponseBody) *QueryConferenceInfoResponse {
	s.Body = v
	return s
}

type QueryConferenceInfoBatchHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryConferenceInfoBatchHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoBatchHeaders) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoBatchHeaders) SetCommonHeaders(v map[string]*string) *QueryConferenceInfoBatchHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryConferenceInfoBatchHeaders) SetXAcsDingtalkAccessToken(v string) *QueryConferenceInfoBatchHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryConferenceInfoBatchRequest struct {
	ConferenceIdList []*string `json:"conferenceIdList,omitempty" xml:"conferenceIdList,omitempty" type:"Repeated"`
}

func (s QueryConferenceInfoBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoBatchRequest) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoBatchRequest) SetConferenceIdList(v []*string) *QueryConferenceInfoBatchRequest {
	s.ConferenceIdList = v
	return s
}

type QueryConferenceInfoBatchResponseBody struct {
	// 会议详情列表
	Infos []*QueryConferenceInfoBatchResponseBodyInfos `json:"infos,omitempty" xml:"infos,omitempty" type:"Repeated"`
}

func (s QueryConferenceInfoBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoBatchResponseBody) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoBatchResponseBody) SetInfos(v []*QueryConferenceInfoBatchResponseBodyInfos) *QueryConferenceInfoBatchResponseBody {
	s.Infos = v
	return s
}

type QueryConferenceInfoBatchResponseBodyInfos struct {
	// 会议iD
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// 媒体状态
	MediaStatus *int64 `json:"mediaStatus,omitempty" xml:"mediaStatus,omitempty"`
	// 会议开始时间
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 会议状态
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 会议名称
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 参会用户列表
	UserList []*QueryConferenceInfoBatchResponseBodyInfosUserList `json:"userList,omitempty" xml:"userList,omitempty" type:"Repeated"`
}

func (s QueryConferenceInfoBatchResponseBodyInfos) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoBatchResponseBodyInfos) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoBatchResponseBodyInfos) SetConferenceId(v string) *QueryConferenceInfoBatchResponseBodyInfos {
	s.ConferenceId = &v
	return s
}

func (s *QueryConferenceInfoBatchResponseBodyInfos) SetMediaStatus(v int64) *QueryConferenceInfoBatchResponseBodyInfos {
	s.MediaStatus = &v
	return s
}

func (s *QueryConferenceInfoBatchResponseBodyInfos) SetStartTime(v int64) *QueryConferenceInfoBatchResponseBodyInfos {
	s.StartTime = &v
	return s
}

func (s *QueryConferenceInfoBatchResponseBodyInfos) SetStatus(v int64) *QueryConferenceInfoBatchResponseBodyInfos {
	s.Status = &v
	return s
}

func (s *QueryConferenceInfoBatchResponseBodyInfos) SetTitle(v string) *QueryConferenceInfoBatchResponseBodyInfos {
	s.Title = &v
	return s
}

func (s *QueryConferenceInfoBatchResponseBodyInfos) SetUserList(v []*QueryConferenceInfoBatchResponseBodyInfosUserList) *QueryConferenceInfoBatchResponseBodyInfos {
	s.UserList = v
	return s
}

type QueryConferenceInfoBatchResponseBodyInfosUserList struct {
	// 在会状态
	AttendStatus *int64 `json:"attendStatus,omitempty" xml:"attendStatus,omitempty"`
	// 摄像头状态
	CameraStatus *int64 `json:"cameraStatus,omitempty" xml:"cameraStatus,omitempty"`
	// 麦克风状态
	MicStatus *int64 `json:"micStatus,omitempty" xml:"micStatus,omitempty"`
	// 名称
	Nick *string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 拒绝原因
	RejectDescription *string `json:"rejectDescription,omitempty" xml:"rejectDescription,omitempty"`
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryConferenceInfoBatchResponseBodyInfosUserList) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoBatchResponseBodyInfosUserList) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoBatchResponseBodyInfosUserList) SetAttendStatus(v int64) *QueryConferenceInfoBatchResponseBodyInfosUserList {
	s.AttendStatus = &v
	return s
}

func (s *QueryConferenceInfoBatchResponseBodyInfosUserList) SetCameraStatus(v int64) *QueryConferenceInfoBatchResponseBodyInfosUserList {
	s.CameraStatus = &v
	return s
}

func (s *QueryConferenceInfoBatchResponseBodyInfosUserList) SetMicStatus(v int64) *QueryConferenceInfoBatchResponseBodyInfosUserList {
	s.MicStatus = &v
	return s
}

func (s *QueryConferenceInfoBatchResponseBodyInfosUserList) SetNick(v string) *QueryConferenceInfoBatchResponseBodyInfosUserList {
	s.Nick = &v
	return s
}

func (s *QueryConferenceInfoBatchResponseBodyInfosUserList) SetRejectDescription(v string) *QueryConferenceInfoBatchResponseBodyInfosUserList {
	s.RejectDescription = &v
	return s
}

func (s *QueryConferenceInfoBatchResponseBodyInfosUserList) SetUserId(v string) *QueryConferenceInfoBatchResponseBodyInfosUserList {
	s.UserId = &v
	return s
}

type QueryConferenceInfoBatchResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryConferenceInfoBatchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryConferenceInfoBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceInfoBatchResponse) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoBatchResponse) SetHeaders(v map[string]*string) *QueryConferenceInfoBatchResponse {
	s.Headers = v
	return s
}

func (s *QueryConferenceInfoBatchResponse) SetBody(v *QueryConferenceInfoBatchResponseBody) *QueryConferenceInfoBatchResponse {
	s.Body = v
	return s
}

type QueryConferenceMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryConferenceMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceMembersHeaders) GoString() string {
	return s.String()
}

func (s *QueryConferenceMembersHeaders) SetCommonHeaders(v map[string]*string) *QueryConferenceMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryConferenceMembersHeaders) SetXAcsDingtalkAccessToken(v string) *QueryConferenceMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryConferenceMembersRequest struct {
	// 返回的最大结果数
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 分页token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryConferenceMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceMembersRequest) GoString() string {
	return s.String()
}

func (s *QueryConferenceMembersRequest) SetMaxResults(v int32) *QueryConferenceMembersRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryConferenceMembersRequest) SetNextToken(v string) *QueryConferenceMembersRequest {
	s.NextToken = &v
	return s
}

type QueryConferenceMembersResponseBody struct {
	// 成员列表
	MemberModels []*QueryConferenceMembersResponseBodyMemberModels `json:"memberModels,omitempty" xml:"memberModels,omitempty" type:"Repeated"`
	// 分页查询下一页token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 本次返回结果数
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryConferenceMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceMembersResponseBody) GoString() string {
	return s.String()
}

func (s *QueryConferenceMembersResponseBody) SetMemberModels(v []*QueryConferenceMembersResponseBodyMemberModels) *QueryConferenceMembersResponseBody {
	s.MemberModels = v
	return s
}

func (s *QueryConferenceMembersResponseBody) SetNextToken(v string) *QueryConferenceMembersResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryConferenceMembersResponseBody) SetTotalCount(v int32) *QueryConferenceMembersResponseBody {
	s.TotalCount = &v
	return s
}

type QueryConferenceMembersResponseBodyMemberModels struct {
	// 成员状态
	// 1 初始化
	// 2 呼叫中
	// 3 活跃（在会）
	// 4 入会失败（拒接等）
	// 5 被踢
	// 6 离会
	AttendStatus *int32 `json:"attendStatus,omitempty" xml:"attendStatus,omitempty"`
	// 是否为联席主持人
	CoHost *bool `json:"coHost,omitempty" xml:"coHost,omitempty"`
	// 会议id
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// 在会时长
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 是否为主持人
	Host *bool `json:"host,omitempty" xml:"host,omitempty"`
	// 入会时间
	JoinTime *int64 `json:"joinTime,omitempty" xml:"joinTime,omitempty"`
	// 离会时间
	LeaveTime *int64 `json:"leaveTime,omitempty" xml:"leaveTime,omitempty"`
	// 是否为非会议所属企业内成员
	OuterOrgMember *bool `json:"outerOrgMember,omitempty" xml:"outerOrgMember,omitempty"`
	// 是否为pstn入会
	PstnJoin *bool `json:"pstnJoin,omitempty" xml:"pstnJoin,omitempty"`
	// 用户unionId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 成员昵称
	UserNick *string `json:"userNick,omitempty" xml:"userNick,omitempty"`
}

func (s QueryConferenceMembersResponseBodyMemberModels) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceMembersResponseBodyMemberModels) GoString() string {
	return s.String()
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetAttendStatus(v int32) *QueryConferenceMembersResponseBodyMemberModels {
	s.AttendStatus = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetCoHost(v bool) *QueryConferenceMembersResponseBodyMemberModels {
	s.CoHost = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetConferenceId(v string) *QueryConferenceMembersResponseBodyMemberModels {
	s.ConferenceId = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetDuration(v int64) *QueryConferenceMembersResponseBodyMemberModels {
	s.Duration = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetHost(v bool) *QueryConferenceMembersResponseBodyMemberModels {
	s.Host = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetJoinTime(v int64) *QueryConferenceMembersResponseBodyMemberModels {
	s.JoinTime = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetLeaveTime(v int64) *QueryConferenceMembersResponseBodyMemberModels {
	s.LeaveTime = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetOuterOrgMember(v bool) *QueryConferenceMembersResponseBodyMemberModels {
	s.OuterOrgMember = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetPstnJoin(v bool) *QueryConferenceMembersResponseBodyMemberModels {
	s.PstnJoin = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetUserId(v string) *QueryConferenceMembersResponseBodyMemberModels {
	s.UserId = &v
	return s
}

func (s *QueryConferenceMembersResponseBodyMemberModels) SetUserNick(v string) *QueryConferenceMembersResponseBodyMemberModels {
	s.UserNick = &v
	return s
}

type QueryConferenceMembersResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryConferenceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryConferenceMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryConferenceMembersResponse) GoString() string {
	return s.String()
}

func (s *QueryConferenceMembersResponse) SetHeaders(v map[string]*string) *QueryConferenceMembersResponse {
	s.Headers = v
	return s
}

func (s *QueryConferenceMembersResponse) SetBody(v *QueryConferenceMembersResponseBody) *QueryConferenceMembersResponse {
	s.Body = v
	return s
}

type StartCloudRecordHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s StartCloudRecordHeaders) String() string {
	return tea.Prettify(s)
}

func (s StartCloudRecordHeaders) GoString() string {
	return s.String()
}

func (s *StartCloudRecordHeaders) SetCommonHeaders(v map[string]*string) *StartCloudRecordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartCloudRecordHeaders) SetXAcsDingtalkAccessToken(v string) *StartCloudRecordHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type StartCloudRecordRequest struct {
	// 录制模版
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// 小窗位置
	SmallWindowPosition *string `json:"smallWindowPosition,omitempty" xml:"smallWindowPosition,omitempty"`
	// 会议发起人UID
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s StartCloudRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s StartCloudRecordRequest) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequest) SetMode(v string) *StartCloudRecordRequest {
	s.Mode = &v
	return s
}

func (s *StartCloudRecordRequest) SetSmallWindowPosition(v string) *StartCloudRecordRequest {
	s.SmallWindowPosition = &v
	return s
}

func (s *StartCloudRecordRequest) SetUnionId(v string) *StartCloudRecordRequest {
	s.UnionId = &v
	return s
}

type StartCloudRecordResponseBody struct {
	// 返回码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s StartCloudRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartCloudRecordResponseBody) GoString() string {
	return s.String()
}

func (s *StartCloudRecordResponseBody) SetCode(v string) *StartCloudRecordResponseBody {
	s.Code = &v
	return s
}

type StartCloudRecordResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartCloudRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartCloudRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s StartCloudRecordResponse) GoString() string {
	return s.String()
}

func (s *StartCloudRecordResponse) SetHeaders(v map[string]*string) *StartCloudRecordResponse {
	s.Headers = v
	return s
}

func (s *StartCloudRecordResponse) SetBody(v *StartCloudRecordResponseBody) *StartCloudRecordResponse {
	s.Body = v
	return s
}

type StartStreamOutHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s StartStreamOutHeaders) String() string {
	return tea.Prettify(s)
}

func (s StartStreamOutHeaders) GoString() string {
	return s.String()
}

func (s *StartStreamOutHeaders) SetCommonHeaders(v map[string]*string) *StartStreamOutHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartStreamOutHeaders) SetXAcsDingtalkAccessToken(v string) *StartStreamOutHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type StartStreamOutRequest struct {
	// 布局
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// 是否需要主持人加入后才允许推流
	NeedHostJoin *bool `json:"needHostJoin,omitempty" xml:"needHostJoin,omitempty"`
	// 小窗位置
	SmallWindowPosition *string `json:"smallWindowPosition,omitempty" xml:"smallWindowPosition,omitempty"`
	StreamName          *string `json:"streamName,omitempty" xml:"streamName,omitempty"`
	// 推流地址列表, 最多10个，需要以RTMP开头
	StreamUrlList []*string `json:"streamUrlList,omitempty" xml:"streamUrlList,omitempty" type:"Repeated"`
	// 主持人UID
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s StartStreamOutRequest) String() string {
	return tea.Prettify(s)
}

func (s StartStreamOutRequest) GoString() string {
	return s.String()
}

func (s *StartStreamOutRequest) SetMode(v string) *StartStreamOutRequest {
	s.Mode = &v
	return s
}

func (s *StartStreamOutRequest) SetNeedHostJoin(v bool) *StartStreamOutRequest {
	s.NeedHostJoin = &v
	return s
}

func (s *StartStreamOutRequest) SetSmallWindowPosition(v string) *StartStreamOutRequest {
	s.SmallWindowPosition = &v
	return s
}

func (s *StartStreamOutRequest) SetStreamName(v string) *StartStreamOutRequest {
	s.StreamName = &v
	return s
}

func (s *StartStreamOutRequest) SetStreamUrlList(v []*string) *StartStreamOutRequest {
	s.StreamUrlList = v
	return s
}

func (s *StartStreamOutRequest) SetUnionId(v string) *StartStreamOutRequest {
	s.UnionId = &v
	return s
}

type StartStreamOutResponseBody struct {
	// 失败的地址与失败原因映射
	FailStreamMap map[string]interface{} `json:"failStreamMap,omitempty" xml:"failStreamMap,omitempty"`
	// 成功推流地址与liveId映射
	SuccessStreamMap map[string]interface{} `json:"successStreamMap,omitempty" xml:"successStreamMap,omitempty"`
}

func (s StartStreamOutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartStreamOutResponseBody) GoString() string {
	return s.String()
}

func (s *StartStreamOutResponseBody) SetFailStreamMap(v map[string]interface{}) *StartStreamOutResponseBody {
	s.FailStreamMap = v
	return s
}

func (s *StartStreamOutResponseBody) SetSuccessStreamMap(v map[string]interface{}) *StartStreamOutResponseBody {
	s.SuccessStreamMap = v
	return s
}

type StartStreamOutResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartStreamOutResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartStreamOutResponse) String() string {
	return tea.Prettify(s)
}

func (s StartStreamOutResponse) GoString() string {
	return s.String()
}

func (s *StartStreamOutResponse) SetHeaders(v map[string]*string) *StartStreamOutResponse {
	s.Headers = v
	return s
}

func (s *StartStreamOutResponse) SetBody(v *StartStreamOutResponseBody) *StartStreamOutResponse {
	s.Body = v
	return s
}

type StopCloudRecordHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s StopCloudRecordHeaders) String() string {
	return tea.Prettify(s)
}

func (s StopCloudRecordHeaders) GoString() string {
	return s.String()
}

func (s *StopCloudRecordHeaders) SetCommonHeaders(v map[string]*string) *StopCloudRecordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StopCloudRecordHeaders) SetXAcsDingtalkAccessToken(v string) *StopCloudRecordHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type StopCloudRecordRequest struct {
	// 主持人uid
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s StopCloudRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s StopCloudRecordRequest) GoString() string {
	return s.String()
}

func (s *StopCloudRecordRequest) SetUnionId(v string) *StopCloudRecordRequest {
	s.UnionId = &v
	return s
}

type StopCloudRecordResponseBody struct {
	// 返回码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s StopCloudRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopCloudRecordResponseBody) GoString() string {
	return s.String()
}

func (s *StopCloudRecordResponseBody) SetCode(v string) *StopCloudRecordResponseBody {
	s.Code = &v
	return s
}

type StopCloudRecordResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopCloudRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopCloudRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s StopCloudRecordResponse) GoString() string {
	return s.String()
}

func (s *StopCloudRecordResponse) SetHeaders(v map[string]*string) *StopCloudRecordResponse {
	s.Headers = v
	return s
}

func (s *StopCloudRecordResponse) SetBody(v *StopCloudRecordResponseBody) *StopCloudRecordResponse {
	s.Body = v
	return s
}

type StopStreamOutHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s StopStreamOutHeaders) String() string {
	return tea.Prettify(s)
}

func (s StopStreamOutHeaders) GoString() string {
	return s.String()
}

func (s *StopStreamOutHeaders) SetCommonHeaders(v map[string]*string) *StopStreamOutHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StopStreamOutHeaders) SetXAcsDingtalkAccessToken(v string) *StopStreamOutHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type StopStreamOutRequest struct {
	// 是否停止所有流，为true时，则不理会streamId字段
	StopAllStream *bool `json:"stopAllStream,omitempty" xml:"stopAllStream,omitempty"`
	// 流id
	StreamId *string `json:"streamId,omitempty" xml:"streamId,omitempty"`
	// 用户id
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s StopStreamOutRequest) String() string {
	return tea.Prettify(s)
}

func (s StopStreamOutRequest) GoString() string {
	return s.String()
}

func (s *StopStreamOutRequest) SetStopAllStream(v bool) *StopStreamOutRequest {
	s.StopAllStream = &v
	return s
}

func (s *StopStreamOutRequest) SetStreamId(v string) *StopStreamOutRequest {
	s.StreamId = &v
	return s
}

func (s *StopStreamOutRequest) SetUnionId(v string) *StopStreamOutRequest {
	s.UnionId = &v
	return s
}

type StopStreamOutResponseBody struct {
	// conferenceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s StopStreamOutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopStreamOutResponseBody) GoString() string {
	return s.String()
}

func (s *StopStreamOutResponseBody) SetCode(v string) *StopStreamOutResponseBody {
	s.Code = &v
	return s
}

type StopStreamOutResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopStreamOutResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopStreamOutResponse) String() string {
	return tea.Prettify(s)
}

func (s StopStreamOutResponse) GoString() string {
	return s.String()
}

func (s *StopStreamOutResponse) SetHeaders(v map[string]*string) *StopStreamOutResponse {
	s.Headers = v
	return s
}

func (s *StopStreamOutResponse) SetBody(v *StopStreamOutResponseBody) *StopStreamOutResponse {
	s.Body = v
	return s
}

type UpdateVideoConferenceExtInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateVideoConferenceExtInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateVideoConferenceExtInfoHeaders) GoString() string {
	return s.String()
}

func (s *UpdateVideoConferenceExtInfoHeaders) SetCommonHeaders(v map[string]*string) *UpdateVideoConferenceExtInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateVideoConferenceExtInfoHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateVideoConferenceExtInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateVideoConferenceExtInfoResponseBody struct {
	// 失败原因
	Case *string `json:"case,omitempty" xml:"case,omitempty"`
	// 返回编码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s UpdateVideoConferenceExtInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVideoConferenceExtInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVideoConferenceExtInfoResponseBody) SetCase(v string) *UpdateVideoConferenceExtInfoResponseBody {
	s.Case = &v
	return s
}

func (s *UpdateVideoConferenceExtInfoResponseBody) SetCode(v string) *UpdateVideoConferenceExtInfoResponseBody {
	s.Code = &v
	return s
}

type UpdateVideoConferenceExtInfoResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateVideoConferenceExtInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateVideoConferenceExtInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVideoConferenceExtInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateVideoConferenceExtInfoResponse) SetHeaders(v map[string]*string) *UpdateVideoConferenceExtInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateVideoConferenceExtInfoResponse) SetBody(v *UpdateVideoConferenceExtInfoResponseBody) *UpdateVideoConferenceExtInfoResponse {
	s.Body = v
	return s
}

type UpdateVideoConferenceSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateVideoConferenceSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateVideoConferenceSettingHeaders) GoString() string {
	return s.String()
}

func (s *UpdateVideoConferenceSettingHeaders) SetCommonHeaders(v map[string]*string) *UpdateVideoConferenceSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateVideoConferenceSettingHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateVideoConferenceSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateVideoConferenceSettingRequest struct {
	// 允许参会人员取消静音
	AllowUnmuteSelf *bool `json:"allowUnmuteSelf,omitempty" xml:"allowUnmuteSelf,omitempty"`
	// 主持人离会，是否自动转移主持人角色
	AutoTransferHost *bool `json:"autoTransferHost,omitempty" xml:"autoTransferHost,omitempty"`
	// 禁止共享屏幕
	ForbiddenShareScreen *bool `json:"forbiddenShareScreen,omitempty" xml:"forbiddenShareScreen,omitempty"`
	// 锁定会议，禁止邀请入会
	LockConference *bool `json:"lockConference,omitempty" xml:"lockConference,omitempty"`
	// 全员静音
	MuteAll *bool `json:"muteAll,omitempty" xml:"muteAll,omitempty"`
	// 仅允许企业内员工加入会议
	OnlyInternalEmployeesJoin *bool `json:"onlyInternalEmployeesJoin,omitempty" xml:"onlyInternalEmployeesJoin,omitempty"`
}

func (s UpdateVideoConferenceSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVideoConferenceSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateVideoConferenceSettingRequest) SetAllowUnmuteSelf(v bool) *UpdateVideoConferenceSettingRequest {
	s.AllowUnmuteSelf = &v
	return s
}

func (s *UpdateVideoConferenceSettingRequest) SetAutoTransferHost(v bool) *UpdateVideoConferenceSettingRequest {
	s.AutoTransferHost = &v
	return s
}

func (s *UpdateVideoConferenceSettingRequest) SetForbiddenShareScreen(v bool) *UpdateVideoConferenceSettingRequest {
	s.ForbiddenShareScreen = &v
	return s
}

func (s *UpdateVideoConferenceSettingRequest) SetLockConference(v bool) *UpdateVideoConferenceSettingRequest {
	s.LockConference = &v
	return s
}

func (s *UpdateVideoConferenceSettingRequest) SetMuteAll(v bool) *UpdateVideoConferenceSettingRequest {
	s.MuteAll = &v
	return s
}

func (s *UpdateVideoConferenceSettingRequest) SetOnlyInternalEmployeesJoin(v bool) *UpdateVideoConferenceSettingRequest {
	s.OnlyInternalEmployeesJoin = &v
	return s
}

type UpdateVideoConferenceSettingResponseBody struct {
	// 结果详情
	Case *string `json:"case,omitempty" xml:"case,omitempty"`
	// 返回编码
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s UpdateVideoConferenceSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVideoConferenceSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVideoConferenceSettingResponseBody) SetCase(v string) *UpdateVideoConferenceSettingResponseBody {
	s.Case = &v
	return s
}

func (s *UpdateVideoConferenceSettingResponseBody) SetCode(v string) *UpdateVideoConferenceSettingResponseBody {
	s.Code = &v
	return s
}

type UpdateVideoConferenceSettingResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateVideoConferenceSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateVideoConferenceSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVideoConferenceSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateVideoConferenceSettingResponse) SetHeaders(v map[string]*string) *UpdateVideoConferenceSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateVideoConferenceSettingResponse) SetBody(v *UpdateVideoConferenceSettingResponseBody) *UpdateVideoConferenceSettingResponse {
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

func (client *Client) CloseVideoConference(conferenceId *string, request *CloseVideoConferenceRequest) (_result *CloseVideoConferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CloseVideoConferenceHeaders{}
	_result = &CloseVideoConferenceResponse{}
	_body, _err := client.CloseVideoConferenceWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloseVideoConferenceWithOptions(conferenceId *string, request *CloseVideoConferenceRequest, headers *CloseVideoConferenceHeaders, runtime *util.RuntimeOptions) (_result *CloseVideoConferenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	conferenceId = openapiutil.GetEncodeParam(conferenceId)
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
	_result = &CloseVideoConferenceResponse{}
	_body, _err := client.DoROARequest(tea.String("CloseVideoConference"), tea.String("conference_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/conference/videoConferences/"+tea.StringValue(conferenceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVideoConference(request *CreateVideoConferenceRequest) (_result *CreateVideoConferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateVideoConferenceHeaders{}
	_result = &CreateVideoConferenceResponse{}
	_body, _err := client.CreateVideoConferenceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVideoConferenceWithOptions(request *CreateVideoConferenceRequest, headers *CreateVideoConferenceHeaders, runtime *util.RuntimeOptions) (_result *CreateVideoConferenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfTitle)) {
		body["confTitle"] = request.ConfTitle
	}

	if !tea.BoolValue(util.IsUnset(request.InviteCaller)) {
		body["inviteCaller"] = request.InviteCaller
	}

	if !tea.BoolValue(util.IsUnset(request.InviteUserIds)) {
		body["inviteUserIds"] = request.InviteUserIds
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
	_result = &CreateVideoConferenceResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateVideoConference"), tea.String("conference_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/conference/videoConferences"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCloudRecordText(conferenceId *string, request *QueryCloudRecordTextRequest) (_result *QueryCloudRecordTextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCloudRecordTextHeaders{}
	_result = &QueryCloudRecordTextResponse{}
	_body, _err := client.QueryCloudRecordTextWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCloudRecordTextWithOptions(conferenceId *string, request *QueryCloudRecordTextRequest, headers *QueryCloudRecordTextHeaders, runtime *util.RuntimeOptions) (_result *QueryCloudRecordTextResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	conferenceId = openapiutil.GetEncodeParam(conferenceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
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
	_result = &QueryCloudRecordTextResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryCloudRecordText"), tea.String("conference_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/conference/videoConferences/"+tea.StringValue(conferenceId)+"/cloudRecords/getTexts"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCloudRecordVideo(conferenceId *string, request *QueryCloudRecordVideoRequest) (_result *QueryCloudRecordVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCloudRecordVideoHeaders{}
	_result = &QueryCloudRecordVideoResponse{}
	_body, _err := client.QueryCloudRecordVideoWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCloudRecordVideoWithOptions(conferenceId *string, request *QueryCloudRecordVideoRequest, headers *QueryCloudRecordVideoHeaders, runtime *util.RuntimeOptions) (_result *QueryCloudRecordVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	conferenceId = openapiutil.GetEncodeParam(conferenceId)
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
	_result = &QueryCloudRecordVideoResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryCloudRecordVideo"), tea.String("conference_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/conference/videoConferences/"+tea.StringValue(conferenceId)+"/cloudRecords/getVideos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCloudRecordVideoPlayInfo(conferenceId *string, request *QueryCloudRecordVideoPlayInfoRequest) (_result *QueryCloudRecordVideoPlayInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCloudRecordVideoPlayInfoHeaders{}
	_result = &QueryCloudRecordVideoPlayInfoResponse{}
	_body, _err := client.QueryCloudRecordVideoPlayInfoWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCloudRecordVideoPlayInfoWithOptions(conferenceId *string, request *QueryCloudRecordVideoPlayInfoRequest, headers *QueryCloudRecordVideoPlayInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryCloudRecordVideoPlayInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	conferenceId = openapiutil.GetEncodeParam(conferenceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MediaId)) {
		query["mediaId"] = request.MediaId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
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
	_result = &QueryCloudRecordVideoPlayInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryCloudRecordVideoPlayInfo"), tea.String("conference_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/conference/videoConferences/"+tea.StringValue(conferenceId)+"/cloudRecords/videos/getPlayInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryConferenceInfo(conferenceId *string) (_result *QueryConferenceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryConferenceInfoHeaders{}
	_result = &QueryConferenceInfoResponse{}
	_body, _err := client.QueryConferenceInfoWithOptions(conferenceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryConferenceInfoWithOptions(conferenceId *string, headers *QueryConferenceInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryConferenceInfoResponse, _err error) {
	conferenceId = openapiutil.GetEncodeParam(conferenceId)
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
	_result = &QueryConferenceInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryConferenceInfo"), tea.String("conference_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/conference/videoConferences/"+tea.StringValue(conferenceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryConferenceInfoBatch(request *QueryConferenceInfoBatchRequest) (_result *QueryConferenceInfoBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryConferenceInfoBatchHeaders{}
	_result = &QueryConferenceInfoBatchResponse{}
	_body, _err := client.QueryConferenceInfoBatchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryConferenceInfoBatchWithOptions(request *QueryConferenceInfoBatchRequest, headers *QueryConferenceInfoBatchHeaders, runtime *util.RuntimeOptions) (_result *QueryConferenceInfoBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConferenceIdList)) {
		body["conferenceIdList"] = request.ConferenceIdList
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &QueryConferenceInfoBatchResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryConferenceInfoBatch"), tea.String("conference_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/conference/videoConferences/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryConferenceMembers(conferenceId *string, request *QueryConferenceMembersRequest) (_result *QueryConferenceMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryConferenceMembersHeaders{}
	_result = &QueryConferenceMembersResponse{}
	_body, _err := client.QueryConferenceMembersWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryConferenceMembersWithOptions(conferenceId *string, request *QueryConferenceMembersRequest, headers *QueryConferenceMembersHeaders, runtime *util.RuntimeOptions) (_result *QueryConferenceMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	conferenceId = openapiutil.GetEncodeParam(conferenceId)
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
	_result = &QueryConferenceMembersResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryConferenceMembers"), tea.String("conference_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/conference/videoConferences/"+tea.StringValue(conferenceId)+"/members"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartCloudRecord(conferenceId *string, request *StartCloudRecordRequest) (_result *StartCloudRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StartCloudRecordHeaders{}
	_result = &StartCloudRecordResponse{}
	_body, _err := client.StartCloudRecordWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartCloudRecordWithOptions(conferenceId *string, request *StartCloudRecordRequest, headers *StartCloudRecordHeaders, runtime *util.RuntimeOptions) (_result *StartCloudRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	conferenceId = openapiutil.GetEncodeParam(conferenceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		body["mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.SmallWindowPosition)) {
		body["smallWindowPosition"] = request.SmallWindowPosition
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
	_result = &StartCloudRecordResponse{}
	_body, _err := client.DoROARequest(tea.String("StartCloudRecord"), tea.String("conference_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/conference/videoConferences/"+tea.StringValue(conferenceId)+"/cloudRecords/start"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartStreamOut(conferenceId *string, request *StartStreamOutRequest) (_result *StartStreamOutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StartStreamOutHeaders{}
	_result = &StartStreamOutResponse{}
	_body, _err := client.StartStreamOutWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartStreamOutWithOptions(conferenceId *string, request *StartStreamOutRequest, headers *StartStreamOutHeaders, runtime *util.RuntimeOptions) (_result *StartStreamOutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	conferenceId = openapiutil.GetEncodeParam(conferenceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		body["mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.NeedHostJoin)) {
		body["needHostJoin"] = request.NeedHostJoin
	}

	if !tea.BoolValue(util.IsUnset(request.SmallWindowPosition)) {
		body["smallWindowPosition"] = request.SmallWindowPosition
	}

	if !tea.BoolValue(util.IsUnset(request.StreamName)) {
		body["streamName"] = request.StreamName
	}

	if !tea.BoolValue(util.IsUnset(request.StreamUrlList)) {
		body["streamUrlList"] = request.StreamUrlList
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
	_result = &StartStreamOutResponse{}
	_body, _err := client.DoROARequest(tea.String("StartStreamOut"), tea.String("conference_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/conference/videoConferences/"+tea.StringValue(conferenceId)+"/streamOuts/start"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopCloudRecord(conferenceId *string, request *StopCloudRecordRequest) (_result *StopCloudRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StopCloudRecordHeaders{}
	_result = &StopCloudRecordResponse{}
	_body, _err := client.StopCloudRecordWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopCloudRecordWithOptions(conferenceId *string, request *StopCloudRecordRequest, headers *StopCloudRecordHeaders, runtime *util.RuntimeOptions) (_result *StopCloudRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	conferenceId = openapiutil.GetEncodeParam(conferenceId)
	body := map[string]interface{}{}
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
	_result = &StopCloudRecordResponse{}
	_body, _err := client.DoROARequest(tea.String("StopCloudRecord"), tea.String("conference_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/conference/videoConferences/"+tea.StringValue(conferenceId)+"/cloudRecords/stop"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopStreamOut(conferenceId *string, request *StopStreamOutRequest) (_result *StopStreamOutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StopStreamOutHeaders{}
	_result = &StopStreamOutResponse{}
	_body, _err := client.StopStreamOutWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopStreamOutWithOptions(conferenceId *string, request *StopStreamOutRequest, headers *StopStreamOutHeaders, runtime *util.RuntimeOptions) (_result *StopStreamOutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	conferenceId = openapiutil.GetEncodeParam(conferenceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StopAllStream)) {
		body["stopAllStream"] = request.StopAllStream
	}

	if !tea.BoolValue(util.IsUnset(request.StreamId)) {
		body["streamId"] = request.StreamId
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
	_result = &StopStreamOutResponse{}
	_body, _err := client.DoROARequest(tea.String("StopStreamOut"), tea.String("conference_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/conference/videoConferences/"+tea.StringValue(conferenceId)+"/streamOuts/stop"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateVideoConferenceExtInfo(conferenceId *string) (_result *UpdateVideoConferenceExtInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateVideoConferenceExtInfoHeaders{}
	_result = &UpdateVideoConferenceExtInfoResponse{}
	_body, _err := client.UpdateVideoConferenceExtInfoWithOptions(conferenceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateVideoConferenceExtInfoWithOptions(conferenceId *string, headers *UpdateVideoConferenceExtInfoHeaders, runtime *util.RuntimeOptions) (_result *UpdateVideoConferenceExtInfoResponse, _err error) {
	conferenceId = openapiutil.GetEncodeParam(conferenceId)
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
	_result = &UpdateVideoConferenceExtInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateVideoConferenceExtInfo"), tea.String("conference_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/conference/videoConferences/"+tea.StringValue(conferenceId)+"/extInfo"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateVideoConferenceSetting(conferenceId *string, request *UpdateVideoConferenceSettingRequest) (_result *UpdateVideoConferenceSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateVideoConferenceSettingHeaders{}
	_result = &UpdateVideoConferenceSettingResponse{}
	_body, _err := client.UpdateVideoConferenceSettingWithOptions(conferenceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateVideoConferenceSettingWithOptions(conferenceId *string, request *UpdateVideoConferenceSettingRequest, headers *UpdateVideoConferenceSettingHeaders, runtime *util.RuntimeOptions) (_result *UpdateVideoConferenceSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	conferenceId = openapiutil.GetEncodeParam(conferenceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowUnmuteSelf)) {
		body["allowUnmuteSelf"] = request.AllowUnmuteSelf
	}

	if !tea.BoolValue(util.IsUnset(request.AutoTransferHost)) {
		body["autoTransferHost"] = request.AutoTransferHost
	}

	if !tea.BoolValue(util.IsUnset(request.ForbiddenShareScreen)) {
		body["forbiddenShareScreen"] = request.ForbiddenShareScreen
	}

	if !tea.BoolValue(util.IsUnset(request.LockConference)) {
		body["lockConference"] = request.LockConference
	}

	if !tea.BoolValue(util.IsUnset(request.MuteAll)) {
		body["muteAll"] = request.MuteAll
	}

	if !tea.BoolValue(util.IsUnset(request.OnlyInternalEmployeesJoin)) {
		body["onlyInternalEmployeesJoin"] = request.OnlyInternalEmployeesJoin
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateVideoConferenceSettingResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateVideoConferenceSetting"), tea.String("conference_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/conference/videoConferences/"+tea.StringValue(conferenceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
