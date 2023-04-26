// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package content_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateFeedHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateFeedHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedHeaders) GoString() string {
	return s.String()
}

func (s *CreateFeedHeaders) SetCommonHeaders(v map[string]*string) *CreateFeedHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateFeedHeaders) SetXAcsDingtalkAccessToken(v string) *CreateFeedHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateFeedRequest struct {
	CourseInfo   *CreateFeedRequestCourseInfo `json:"courseInfo,omitempty" xml:"courseInfo,omitempty" type:"Struct"`
	CreateUserId *string                      `json:"createUserId,omitempty" xml:"createUserId,omitempty"`
	FeedInfo     *CreateFeedRequestFeedInfo   `json:"feedInfo,omitempty" xml:"feedInfo,omitempty" type:"Struct"`
}

func (s CreateFeedRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedRequest) GoString() string {
	return s.String()
}

func (s *CreateFeedRequest) SetCourseInfo(v *CreateFeedRequestCourseInfo) *CreateFeedRequest {
	s.CourseInfo = v
	return s
}

func (s *CreateFeedRequest) SetCreateUserId(v string) *CreateFeedRequest {
	s.CreateUserId = &v
	return s
}

func (s *CreateFeedRequest) SetFeedInfo(v *CreateFeedRequestFeedInfo) *CreateFeedRequest {
	s.FeedInfo = v
	return s
}

type CreateFeedRequestCourseInfo struct {
	LectorUserInfo *CreateFeedRequestCourseInfoLectorUserInfo `json:"lectorUserInfo,omitempty" xml:"lectorUserInfo,omitempty" type:"Struct"`
	PayInfo        *CreateFeedRequestCourseInfoPayInfo        `json:"payInfo,omitempty" xml:"payInfo,omitempty" type:"Struct"`
	StudyGroupName *string                                    `json:"studyGroupName,omitempty" xml:"studyGroupName,omitempty"`
}

func (s CreateFeedRequestCourseInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedRequestCourseInfo) GoString() string {
	return s.String()
}

func (s *CreateFeedRequestCourseInfo) SetLectorUserInfo(v *CreateFeedRequestCourseInfoLectorUserInfo) *CreateFeedRequestCourseInfo {
	s.LectorUserInfo = v
	return s
}

func (s *CreateFeedRequestCourseInfo) SetPayInfo(v *CreateFeedRequestCourseInfoPayInfo) *CreateFeedRequestCourseInfo {
	s.PayInfo = v
	return s
}

func (s *CreateFeedRequestCourseInfo) SetStudyGroupName(v string) *CreateFeedRequestCourseInfo {
	s.StudyGroupName = &v
	return s
}

type CreateFeedRequestCourseInfoLectorUserInfo struct {
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateFeedRequestCourseInfoLectorUserInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedRequestCourseInfoLectorUserInfo) GoString() string {
	return s.String()
}

func (s *CreateFeedRequestCourseInfoLectorUserInfo) SetAvatar(v string) *CreateFeedRequestCourseInfoLectorUserInfo {
	s.Avatar = &v
	return s
}

func (s *CreateFeedRequestCourseInfoLectorUserInfo) SetName(v string) *CreateFeedRequestCourseInfoLectorUserInfo {
	s.Name = &v
	return s
}

func (s *CreateFeedRequestCourseInfoLectorUserInfo) SetUserId(v string) *CreateFeedRequestCourseInfoLectorUserInfo {
	s.UserId = &v
	return s
}

type CreateFeedRequestCourseInfoPayInfo struct {
	CsUserInfo   *CreateFeedRequestCourseInfoPayInfoCsUserInfo   `json:"csUserInfo,omitempty" xml:"csUserInfo,omitempty" type:"Struct"`
	DiscountInfo *CreateFeedRequestCourseInfoPayInfoDiscountInfo `json:"discountInfo,omitempty" xml:"discountInfo,omitempty" type:"Struct"`
	Price        *int64                                          `json:"price,omitempty" xml:"price,omitempty"`
}

func (s CreateFeedRequestCourseInfoPayInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedRequestCourseInfoPayInfo) GoString() string {
	return s.String()
}

func (s *CreateFeedRequestCourseInfoPayInfo) SetCsUserInfo(v *CreateFeedRequestCourseInfoPayInfoCsUserInfo) *CreateFeedRequestCourseInfoPayInfo {
	s.CsUserInfo = v
	return s
}

func (s *CreateFeedRequestCourseInfoPayInfo) SetDiscountInfo(v *CreateFeedRequestCourseInfoPayInfoDiscountInfo) *CreateFeedRequestCourseInfoPayInfo {
	s.DiscountInfo = v
	return s
}

func (s *CreateFeedRequestCourseInfoPayInfo) SetPrice(v int64) *CreateFeedRequestCourseInfoPayInfo {
	s.Price = &v
	return s
}

type CreateFeedRequestCourseInfoPayInfoCsUserInfo struct {
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateFeedRequestCourseInfoPayInfoCsUserInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedRequestCourseInfoPayInfoCsUserInfo) GoString() string {
	return s.String()
}

func (s *CreateFeedRequestCourseInfoPayInfoCsUserInfo) SetAvatar(v string) *CreateFeedRequestCourseInfoPayInfoCsUserInfo {
	s.Avatar = &v
	return s
}

func (s *CreateFeedRequestCourseInfoPayInfoCsUserInfo) SetName(v string) *CreateFeedRequestCourseInfoPayInfoCsUserInfo {
	s.Name = &v
	return s
}

func (s *CreateFeedRequestCourseInfoPayInfoCsUserInfo) SetUserId(v string) *CreateFeedRequestCourseInfoPayInfoCsUserInfo {
	s.UserId = &v
	return s
}

type CreateFeedRequestCourseInfoPayInfoDiscountInfo struct {
	EndTimeMillis   *int64 `json:"endTimeMillis,omitempty" xml:"endTimeMillis,omitempty"`
	Price           *int64 `json:"price,omitempty" xml:"price,omitempty"`
	StartTimeMillis *int64 `json:"startTimeMillis,omitempty" xml:"startTimeMillis,omitempty"`
}

func (s CreateFeedRequestCourseInfoPayInfoDiscountInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedRequestCourseInfoPayInfoDiscountInfo) GoString() string {
	return s.String()
}

func (s *CreateFeedRequestCourseInfoPayInfoDiscountInfo) SetEndTimeMillis(v int64) *CreateFeedRequestCourseInfoPayInfoDiscountInfo {
	s.EndTimeMillis = &v
	return s
}

func (s *CreateFeedRequestCourseInfoPayInfoDiscountInfo) SetPrice(v int64) *CreateFeedRequestCourseInfoPayInfoDiscountInfo {
	s.Price = &v
	return s
}

func (s *CreateFeedRequestCourseInfoPayInfoDiscountInfo) SetStartTimeMillis(v int64) *CreateFeedRequestCourseInfoPayInfoDiscountInfo {
	s.StartTimeMillis = &v
	return s
}

type CreateFeedRequestFeedInfo struct {
	ActionType         *int32                                    `json:"actionType,omitempty" xml:"actionType,omitempty"`
	BelongsTo          *int32                                    `json:"belongsTo,omitempty" xml:"belongsTo,omitempty"`
	FeedCategory       *int64                                    `json:"feedCategory,omitempty" xml:"feedCategory,omitempty"`
	FeedId             *string                                   `json:"feedId,omitempty" xml:"feedId,omitempty"`
	FeedTag            *string                                   `json:"feedTag,omitempty" xml:"feedTag,omitempty"`
	FeedType           *int32                                    `json:"feedType,omitempty" xml:"feedType,omitempty"`
	IndustryId         *int64                                    `json:"industryId,omitempty" xml:"industryId,omitempty"`
	Introduction       *string                                   `json:"introduction,omitempty" xml:"introduction,omitempty"`
	IntroductionPicUrl *string                                   `json:"introductionPicUrl,omitempty" xml:"introductionPicUrl,omitempty"`
	McnId              *string                                   `json:"mcnId,omitempty" xml:"mcnId,omitempty"`
	MediaContents      []*CreateFeedRequestFeedInfoMediaContents `json:"mediaContents,omitempty" xml:"mediaContents,omitempty" type:"Repeated"`
	Recommends         []*CreateFeedRequestFeedInfoRecommends    `json:"recommends,omitempty" xml:"recommends,omitempty" type:"Repeated"`
	ThumbUrl           *string                                   `json:"thumbUrl,omitempty" xml:"thumbUrl,omitempty"`
	Title              *string                                   `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateFeedRequestFeedInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedRequestFeedInfo) GoString() string {
	return s.String()
}

func (s *CreateFeedRequestFeedInfo) SetActionType(v int32) *CreateFeedRequestFeedInfo {
	s.ActionType = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetBelongsTo(v int32) *CreateFeedRequestFeedInfo {
	s.BelongsTo = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetFeedCategory(v int64) *CreateFeedRequestFeedInfo {
	s.FeedCategory = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetFeedId(v string) *CreateFeedRequestFeedInfo {
	s.FeedId = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetFeedTag(v string) *CreateFeedRequestFeedInfo {
	s.FeedTag = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetFeedType(v int32) *CreateFeedRequestFeedInfo {
	s.FeedType = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetIndustryId(v int64) *CreateFeedRequestFeedInfo {
	s.IndustryId = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetIntroduction(v string) *CreateFeedRequestFeedInfo {
	s.Introduction = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetIntroductionPicUrl(v string) *CreateFeedRequestFeedInfo {
	s.IntroductionPicUrl = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetMcnId(v string) *CreateFeedRequestFeedInfo {
	s.McnId = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetMediaContents(v []*CreateFeedRequestFeedInfoMediaContents) *CreateFeedRequestFeedInfo {
	s.MediaContents = v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetRecommends(v []*CreateFeedRequestFeedInfoRecommends) *CreateFeedRequestFeedInfo {
	s.Recommends = v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetThumbUrl(v string) *CreateFeedRequestFeedInfo {
	s.ThumbUrl = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetTitle(v string) *CreateFeedRequestFeedInfo {
	s.Title = &v
	return s
}

type CreateFeedRequestFeedInfoMediaContents struct {
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	Title   *string `json:"title,omitempty" xml:"title,omitempty"`
	Type    *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateFeedRequestFeedInfoMediaContents) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedRequestFeedInfoMediaContents) GoString() string {
	return s.String()
}

func (s *CreateFeedRequestFeedInfoMediaContents) SetMediaId(v string) *CreateFeedRequestFeedInfoMediaContents {
	s.MediaId = &v
	return s
}

func (s *CreateFeedRequestFeedInfoMediaContents) SetTitle(v string) *CreateFeedRequestFeedInfoMediaContents {
	s.Title = &v
	return s
}

func (s *CreateFeedRequestFeedInfoMediaContents) SetType(v int32) *CreateFeedRequestFeedInfoMediaContents {
	s.Type = &v
	return s
}

type CreateFeedRequestFeedInfoRecommends struct {
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	Type     *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateFeedRequestFeedInfoRecommends) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedRequestFeedInfoRecommends) GoString() string {
	return s.String()
}

func (s *CreateFeedRequestFeedInfoRecommends) SetObjectId(v string) *CreateFeedRequestFeedInfoRecommends {
	s.ObjectId = &v
	return s
}

func (s *CreateFeedRequestFeedInfoRecommends) SetType(v int32) *CreateFeedRequestFeedInfoRecommends {
	s.Type = &v
	return s
}

type CreateFeedResponseBody struct {
	FeedId *string `json:"feedId,omitempty" xml:"feedId,omitempty"`
}

func (s CreateFeedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFeedResponseBody) SetFeedId(v string) *CreateFeedResponseBody {
	s.FeedId = &v
	return s
}

type CreateFeedResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFeedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFeedResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedResponse) GoString() string {
	return s.String()
}

func (s *CreateFeedResponse) SetHeaders(v map[string]*string) *CreateFeedResponse {
	s.Headers = v
	return s
}

func (s *CreateFeedResponse) SetStatusCode(v int32) *CreateFeedResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFeedResponse) SetBody(v *CreateFeedResponseBody) *CreateFeedResponse {
	s.Body = v
	return s
}

type GetFeedHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFeedHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFeedHeaders) GoString() string {
	return s.String()
}

func (s *GetFeedHeaders) SetCommonHeaders(v map[string]*string) *GetFeedHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFeedHeaders) SetXAcsDingtalkAccessToken(v string) *GetFeedHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFeedRequest struct {
	McnId *string `json:"mcnId,omitempty" xml:"mcnId,omitempty"`
}

func (s GetFeedRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFeedRequest) GoString() string {
	return s.String()
}

func (s *GetFeedRequest) SetMcnId(v string) *GetFeedRequest {
	s.McnId = &v
	return s
}

type GetFeedResponseBody struct {
	FeedId   *string                        `json:"feedId,omitempty" xml:"feedId,omitempty"`
	FeedItem []*GetFeedResponseBodyFeedItem `json:"feedItem,omitempty" xml:"feedItem,omitempty" type:"Repeated"`
}

func (s GetFeedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFeedResponseBody) GoString() string {
	return s.String()
}

func (s *GetFeedResponseBody) SetFeedId(v string) *GetFeedResponseBody {
	s.FeedId = &v
	return s
}

func (s *GetFeedResponseBody) SetFeedItem(v []*GetFeedResponseBodyFeedItem) *GetFeedResponseBody {
	s.FeedItem = v
	return s
}

type GetFeedResponseBodyFeedItem struct {
	DurationMillis  *int64  `json:"durationMillis,omitempty" xml:"durationMillis,omitempty"`
	FeedContentType *int32  `json:"feedContentType,omitempty" xml:"feedContentType,omitempty"`
	ItemId          *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	Title           *string `json:"title,omitempty" xml:"title,omitempty"`
	Url             *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetFeedResponseBodyFeedItem) String() string {
	return tea.Prettify(s)
}

func (s GetFeedResponseBodyFeedItem) GoString() string {
	return s.String()
}

func (s *GetFeedResponseBodyFeedItem) SetDurationMillis(v int64) *GetFeedResponseBodyFeedItem {
	s.DurationMillis = &v
	return s
}

func (s *GetFeedResponseBodyFeedItem) SetFeedContentType(v int32) *GetFeedResponseBodyFeedItem {
	s.FeedContentType = &v
	return s
}

func (s *GetFeedResponseBodyFeedItem) SetItemId(v string) *GetFeedResponseBodyFeedItem {
	s.ItemId = &v
	return s
}

func (s *GetFeedResponseBodyFeedItem) SetTitle(v string) *GetFeedResponseBodyFeedItem {
	s.Title = &v
	return s
}

func (s *GetFeedResponseBodyFeedItem) SetUrl(v string) *GetFeedResponseBodyFeedItem {
	s.Url = &v
	return s
}

type GetFeedResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFeedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFeedResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFeedResponse) GoString() string {
	return s.String()
}

func (s *GetFeedResponse) SetHeaders(v map[string]*string) *GetFeedResponse {
	s.Headers = v
	return s
}

func (s *GetFeedResponse) SetStatusCode(v int32) *GetFeedResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFeedResponse) SetBody(v *GetFeedResponseBody) *GetFeedResponse {
	s.Body = v
	return s
}

type GetMediaCerficateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetMediaCerficateHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetMediaCerficateHeaders) GoString() string {
	return s.String()
}

func (s *GetMediaCerficateHeaders) SetCommonHeaders(v map[string]*string) *GetMediaCerficateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMediaCerficateHeaders) SetXAcsDingtalkAccessToken(v string) *GetMediaCerficateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetMediaCerficateRequest struct {
	FileName          *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	McnId             *string `json:"mcnId,omitempty" xml:"mcnId,omitempty"`
	MediaId           *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	MediaIntroduction *string `json:"mediaIntroduction,omitempty" xml:"mediaIntroduction,omitempty"`
	MediaTitle        *string `json:"mediaTitle,omitempty" xml:"mediaTitle,omitempty"`
	ThumbUrl          *string `json:"thumbUrl,omitempty" xml:"thumbUrl,omitempty"`
	UserId            *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetMediaCerficateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMediaCerficateRequest) GoString() string {
	return s.String()
}

func (s *GetMediaCerficateRequest) SetFileName(v string) *GetMediaCerficateRequest {
	s.FileName = &v
	return s
}

func (s *GetMediaCerficateRequest) SetMcnId(v string) *GetMediaCerficateRequest {
	s.McnId = &v
	return s
}

func (s *GetMediaCerficateRequest) SetMediaId(v string) *GetMediaCerficateRequest {
	s.MediaId = &v
	return s
}

func (s *GetMediaCerficateRequest) SetMediaIntroduction(v string) *GetMediaCerficateRequest {
	s.MediaIntroduction = &v
	return s
}

func (s *GetMediaCerficateRequest) SetMediaTitle(v string) *GetMediaCerficateRequest {
	s.MediaTitle = &v
	return s
}

func (s *GetMediaCerficateRequest) SetThumbUrl(v string) *GetMediaCerficateRequest {
	s.ThumbUrl = &v
	return s
}

func (s *GetMediaCerficateRequest) SetUserId(v string) *GetMediaCerficateRequest {
	s.UserId = &v
	return s
}

type GetMediaCerficateResponseBody struct {
	MediaId            *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	OssAccessKeyId     *string `json:"ossAccessKeyId,omitempty" xml:"ossAccessKeyId,omitempty"`
	OssAccessKeySecret *string `json:"ossAccessKeySecret,omitempty" xml:"ossAccessKeySecret,omitempty"`
	OssBucketName      *string `json:"ossBucketName,omitempty" xml:"ossBucketName,omitempty"`
	OssEndpoint        *string `json:"ossEndpoint,omitempty" xml:"ossEndpoint,omitempty"`
	OssExpiration      *string `json:"ossExpiration,omitempty" xml:"ossExpiration,omitempty"`
	OssFileName        *string `json:"ossFileName,omitempty" xml:"ossFileName,omitempty"`
	OssSecurityToken   *string `json:"ossSecurityToken,omitempty" xml:"ossSecurityToken,omitempty"`
}

func (s GetMediaCerficateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMediaCerficateResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaCerficateResponseBody) SetMediaId(v string) *GetMediaCerficateResponseBody {
	s.MediaId = &v
	return s
}

func (s *GetMediaCerficateResponseBody) SetOssAccessKeyId(v string) *GetMediaCerficateResponseBody {
	s.OssAccessKeyId = &v
	return s
}

func (s *GetMediaCerficateResponseBody) SetOssAccessKeySecret(v string) *GetMediaCerficateResponseBody {
	s.OssAccessKeySecret = &v
	return s
}

func (s *GetMediaCerficateResponseBody) SetOssBucketName(v string) *GetMediaCerficateResponseBody {
	s.OssBucketName = &v
	return s
}

func (s *GetMediaCerficateResponseBody) SetOssEndpoint(v string) *GetMediaCerficateResponseBody {
	s.OssEndpoint = &v
	return s
}

func (s *GetMediaCerficateResponseBody) SetOssExpiration(v string) *GetMediaCerficateResponseBody {
	s.OssExpiration = &v
	return s
}

func (s *GetMediaCerficateResponseBody) SetOssFileName(v string) *GetMediaCerficateResponseBody {
	s.OssFileName = &v
	return s
}

func (s *GetMediaCerficateResponseBody) SetOssSecurityToken(v string) *GetMediaCerficateResponseBody {
	s.OssSecurityToken = &v
	return s
}

type GetMediaCerficateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMediaCerficateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMediaCerficateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMediaCerficateResponse) GoString() string {
	return s.String()
}

func (s *GetMediaCerficateResponse) SetHeaders(v map[string]*string) *GetMediaCerficateResponse {
	s.Headers = v
	return s
}

func (s *GetMediaCerficateResponse) SetStatusCode(v int32) *GetMediaCerficateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaCerficateResponse) SetBody(v *GetMediaCerficateResponseBody) *GetMediaCerficateResponse {
	s.Body = v
	return s
}

type ListItemUserDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListItemUserDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListItemUserDataHeaders) GoString() string {
	return s.String()
}

func (s *ListItemUserDataHeaders) SetCommonHeaders(v map[string]*string) *ListItemUserDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListItemUserDataHeaders) SetXAcsDingtalkAccessToken(v string) *ListItemUserDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListItemUserDataRequest struct {
	Body []*string `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s ListItemUserDataRequest) String() string {
	return tea.Prettify(s)
}

func (s ListItemUserDataRequest) GoString() string {
	return s.String()
}

func (s *ListItemUserDataRequest) SetBody(v []*string) *ListItemUserDataRequest {
	s.Body = v
	return s
}

type ListItemUserDataResponseBody struct {
	StudyInfos []*ListItemUserDataResponseBodyStudyInfos `json:"studyInfos,omitempty" xml:"studyInfos,omitempty" type:"Repeated"`
}

func (s ListItemUserDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListItemUserDataResponseBody) GoString() string {
	return s.String()
}

func (s *ListItemUserDataResponseBody) SetStudyInfos(v []*ListItemUserDataResponseBodyStudyInfos) *ListItemUserDataResponseBody {
	s.StudyInfos = v
	return s
}

type ListItemUserDataResponseBodyStudyInfos struct {
	DurationMillis *int64  `json:"durationMillis,omitempty" xml:"durationMillis,omitempty"`
	Uid            *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ListItemUserDataResponseBodyStudyInfos) String() string {
	return tea.Prettify(s)
}

func (s ListItemUserDataResponseBodyStudyInfos) GoString() string {
	return s.String()
}

func (s *ListItemUserDataResponseBodyStudyInfos) SetDurationMillis(v int64) *ListItemUserDataResponseBodyStudyInfos {
	s.DurationMillis = &v
	return s
}

func (s *ListItemUserDataResponseBodyStudyInfos) SetUid(v string) *ListItemUserDataResponseBodyStudyInfos {
	s.Uid = &v
	return s
}

type ListItemUserDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListItemUserDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListItemUserDataResponse) String() string {
	return tea.Prettify(s)
}

func (s ListItemUserDataResponse) GoString() string {
	return s.String()
}

func (s *ListItemUserDataResponse) SetHeaders(v map[string]*string) *ListItemUserDataResponse {
	s.Headers = v
	return s
}

func (s *ListItemUserDataResponse) SetStatusCode(v int32) *ListItemUserDataResponse {
	s.StatusCode = &v
	return s
}

func (s *ListItemUserDataResponse) SetBody(v *ListItemUserDataResponseBody) *ListItemUserDataResponse {
	s.Body = v
	return s
}

type PageFeedHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PageFeedHeaders) String() string {
	return tea.Prettify(s)
}

func (s PageFeedHeaders) GoString() string {
	return s.String()
}

func (s *PageFeedHeaders) SetCommonHeaders(v map[string]*string) *PageFeedHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PageFeedHeaders) SetXAcsDingtalkAccessToken(v string) *PageFeedHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PageFeedRequest struct {
	Body       []*string `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	MaxResults *int32    `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	McnId      *string   `json:"mcnId,omitempty" xml:"mcnId,omitempty"`
	NextToken  *int32    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s PageFeedRequest) String() string {
	return tea.Prettify(s)
}

func (s PageFeedRequest) GoString() string {
	return s.String()
}

func (s *PageFeedRequest) SetBody(v []*string) *PageFeedRequest {
	s.Body = v
	return s
}

func (s *PageFeedRequest) SetMaxResults(v int32) *PageFeedRequest {
	s.MaxResults = &v
	return s
}

func (s *PageFeedRequest) SetMcnId(v string) *PageFeedRequest {
	s.McnId = &v
	return s
}

func (s *PageFeedRequest) SetNextToken(v int32) *PageFeedRequest {
	s.NextToken = &v
	return s
}

type PageFeedResponseBody struct {
	FeedList   []*PageFeedResponseBodyFeedList `json:"feedList,omitempty" xml:"feedList,omitempty" type:"Repeated"`
	HasNext    *bool                           `json:"hasNext,omitempty" xml:"hasNext,omitempty"`
	NextCursor *int32                          `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
}

func (s PageFeedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageFeedResponseBody) GoString() string {
	return s.String()
}

func (s *PageFeedResponseBody) SetFeedList(v []*PageFeedResponseBodyFeedList) *PageFeedResponseBody {
	s.FeedList = v
	return s
}

func (s *PageFeedResponseBody) SetHasNext(v bool) *PageFeedResponseBody {
	s.HasNext = &v
	return s
}

func (s *PageFeedResponseBody) SetNextCursor(v int32) *PageFeedResponseBody {
	s.NextCursor = &v
	return s
}

type PageFeedResponseBodyFeedList struct {
	FeedCategory *string `json:"feedCategory,omitempty" xml:"feedCategory,omitempty"`
	FeedId       *string `json:"feedId,omitempty" xml:"feedId,omitempty"`
	FeedType     *int32  `json:"feedType,omitempty" xml:"feedType,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	ThumbUrl     *string `json:"thumbUrl,omitempty" xml:"thumbUrl,omitempty"`
	Url          *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s PageFeedResponseBodyFeedList) String() string {
	return tea.Prettify(s)
}

func (s PageFeedResponseBodyFeedList) GoString() string {
	return s.String()
}

func (s *PageFeedResponseBodyFeedList) SetFeedCategory(v string) *PageFeedResponseBodyFeedList {
	s.FeedCategory = &v
	return s
}

func (s *PageFeedResponseBodyFeedList) SetFeedId(v string) *PageFeedResponseBodyFeedList {
	s.FeedId = &v
	return s
}

func (s *PageFeedResponseBodyFeedList) SetFeedType(v int32) *PageFeedResponseBodyFeedList {
	s.FeedType = &v
	return s
}

func (s *PageFeedResponseBodyFeedList) SetName(v string) *PageFeedResponseBodyFeedList {
	s.Name = &v
	return s
}

func (s *PageFeedResponseBodyFeedList) SetThumbUrl(v string) *PageFeedResponseBodyFeedList {
	s.ThumbUrl = &v
	return s
}

func (s *PageFeedResponseBodyFeedList) SetUrl(v string) *PageFeedResponseBodyFeedList {
	s.Url = &v
	return s
}

type PageFeedResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PageFeedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PageFeedResponse) String() string {
	return tea.Prettify(s)
}

func (s PageFeedResponse) GoString() string {
	return s.String()
}

func (s *PageFeedResponse) SetHeaders(v map[string]*string) *PageFeedResponse {
	s.Headers = v
	return s
}

func (s *PageFeedResponse) SetStatusCode(v int32) *PageFeedResponse {
	s.StatusCode = &v
	return s
}

func (s *PageFeedResponse) SetBody(v *PageFeedResponseBody) *PageFeedResponse {
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

func (client *Client) CreateFeedWithOptions(request *CreateFeedRequest, headers *CreateFeedHeaders, runtime *util.RuntimeOptions) (_result *CreateFeedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CourseInfo)) {
		body["courseInfo"] = request.CourseInfo
	}

	if !tea.BoolValue(util.IsUnset(request.CreateUserId)) {
		body["createUserId"] = request.CreateUserId
	}

	if !tea.BoolValue(util.IsUnset(request.FeedInfo)) {
		body["feedInfo"] = request.FeedInfo
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
		Action:      tea.String("CreateFeed"),
		Version:     tea.String("content_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/content/feeds"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFeedResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFeed(request *CreateFeedRequest) (_result *CreateFeedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateFeedHeaders{}
	_result = &CreateFeedResponse{}
	_body, _err := client.CreateFeedWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFeedWithOptions(feedId *string, request *GetFeedRequest, headers *GetFeedHeaders, runtime *util.RuntimeOptions) (_result *GetFeedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.McnId)) {
		query["mcnId"] = request.McnId
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
		Action:      tea.String("GetFeed"),
		Version:     tea.String("content_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/content/feeds/" + tea.StringValue(feedId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFeedResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFeed(feedId *string, request *GetFeedRequest) (_result *GetFeedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFeedHeaders{}
	_result = &GetFeedResponse{}
	_body, _err := client.GetFeedWithOptions(feedId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMediaCerficateWithOptions(request *GetMediaCerficateRequest, headers *GetMediaCerficateHeaders, runtime *util.RuntimeOptions) (_result *GetMediaCerficateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.McnId)) {
		query["mcnId"] = request.McnId
	}

	if !tea.BoolValue(util.IsUnset(request.MediaId)) {
		query["mediaId"] = request.MediaId
	}

	if !tea.BoolValue(util.IsUnset(request.MediaIntroduction)) {
		query["mediaIntroduction"] = request.MediaIntroduction
	}

	if !tea.BoolValue(util.IsUnset(request.MediaTitle)) {
		query["mediaTitle"] = request.MediaTitle
	}

	if !tea.BoolValue(util.IsUnset(request.ThumbUrl)) {
		query["thumbUrl"] = request.ThumbUrl
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
		Action:      tea.String("GetMediaCerficate"),
		Version:     tea.String("content_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/content/media/cerficates"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMediaCerficateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMediaCerficate(request *GetMediaCerficateRequest) (_result *GetMediaCerficateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetMediaCerficateHeaders{}
	_result = &GetMediaCerficateResponse{}
	_body, _err := client.GetMediaCerficateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListItemUserDataWithOptions(itemId *string, request *ListItemUserDataRequest, headers *ListItemUserDataHeaders, runtime *util.RuntimeOptions) (_result *ListItemUserDataResponse, _err error) {
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
		Action:      tea.String("ListItemUserData"),
		Version:     tea.String("content_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/content/feeds/items/" + tea.StringValue(itemId) + "/userStatistics/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListItemUserDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListItemUserData(itemId *string, request *ListItemUserDataRequest) (_result *ListItemUserDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListItemUserDataHeaders{}
	_result = &ListItemUserDataResponse{}
	_body, _err := client.ListItemUserDataWithOptions(itemId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PageFeedWithOptions(request *PageFeedRequest, headers *PageFeedHeaders, runtime *util.RuntimeOptions) (_result *PageFeedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.McnId)) {
		query["mcnId"] = request.McnId
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
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("PageFeed"),
		Version:     tea.String("content_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/content/feeds/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PageFeedResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PageFeed(request *PageFeedRequest) (_result *PageFeedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PageFeedHeaders{}
	_result = &PageFeedResponse{}
	_body, _err := client.PageFeedWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
