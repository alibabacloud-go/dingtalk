// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package content_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
	// 自定义视频封面的URL地址
	ThumbUrl *string `json:"thumbUrl,omitempty" xml:"thumbUrl,omitempty"`
	// 视频的文件名称,必须带扩展名,支持的扩展名参考:https://help.aliyun.com/document_detail/55396.htm?spm=a2c4g.11186623.2.11.2d385d4aG2IkCZ#title-j7o-gr4-c7a
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 如果传入该值，表示续订该mediaId对应的上传凭证 ;否则将视为新建一个视频上传连接和凭证
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	// 视频的标题
	MediaTitle *string `json:"mediaTitle,omitempty" xml:"mediaTitle,omitempty"`
	// 入驻账号Id(请联系钉钉接口同学获取)
	McnId *string `json:"mcnId,omitempty" xml:"mcnId,omitempty"`
	// 操作人的用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 视频介绍
	MediaIntroduction *string `json:"mediaIntroduction,omitempty" xml:"mediaIntroduction,omitempty"`
}

func (s GetMediaCerficateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMediaCerficateRequest) GoString() string {
	return s.String()
}

func (s *GetMediaCerficateRequest) SetThumbUrl(v string) *GetMediaCerficateRequest {
	s.ThumbUrl = &v
	return s
}

func (s *GetMediaCerficateRequest) SetFileName(v string) *GetMediaCerficateRequest {
	s.FileName = &v
	return s
}

func (s *GetMediaCerficateRequest) SetMediaId(v string) *GetMediaCerficateRequest {
	s.MediaId = &v
	return s
}

func (s *GetMediaCerficateRequest) SetMediaTitle(v string) *GetMediaCerficateRequest {
	s.MediaTitle = &v
	return s
}

func (s *GetMediaCerficateRequest) SetMcnId(v string) *GetMediaCerficateRequest {
	s.McnId = &v
	return s
}

func (s *GetMediaCerficateRequest) SetUserId(v string) *GetMediaCerficateRequest {
	s.UserId = &v
	return s
}

func (s *GetMediaCerficateRequest) SetMediaIntroduction(v string) *GetMediaCerficateRequest {
	s.MediaIntroduction = &v
	return s
}

type GetMediaCerficateResponseBody struct {
	// 媒体文件ID
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	// OSS区域地址
	OssEndpoint *string `json:"ossEndpoint,omitempty" xml:"ossEndpoint,omitempty"`
	// 上传授权密钥ID
	OssAccessKeyId *string `json:"ossAccessKeyId,omitempty" xml:"ossAccessKeyId,omitempty"`
	// 上传授权密钥
	OssAccessKeySecret *string `json:"ossAccessKeySecret,omitempty" xml:"ossAccessKeySecret,omitempty"`
	// 上传授权安全令牌
	OssSecurityToken *string `json:"ossSecurityToken,omitempty" xml:"ossSecurityToken,omitempty"`
	// OSS Bucket名称
	OssBucketName *string `json:"ossBucketName,omitempty" xml:"ossBucketName,omitempty"`
	// 分配的媒体文件名
	OssFileName *string `json:"ossFileName,omitempty" xml:"ossFileName,omitempty"`
	// 凭证有效时间(单位秒)，当上传凭证过期时，需要重新使用本次返回的videoId重新调用接口进行凭证刷新
	OssExpiration *string `json:"ossExpiration,omitempty" xml:"ossExpiration,omitempty"`
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

func (s *GetMediaCerficateResponseBody) SetOssEndpoint(v string) *GetMediaCerficateResponseBody {
	s.OssEndpoint = &v
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

func (s *GetMediaCerficateResponseBody) SetOssSecurityToken(v string) *GetMediaCerficateResponseBody {
	s.OssSecurityToken = &v
	return s
}

func (s *GetMediaCerficateResponseBody) SetOssBucketName(v string) *GetMediaCerficateResponseBody {
	s.OssBucketName = &v
	return s
}

func (s *GetMediaCerficateResponseBody) SetOssFileName(v string) *GetMediaCerficateResponseBody {
	s.OssFileName = &v
	return s
}

func (s *GetMediaCerficateResponseBody) SetOssExpiration(v string) *GetMediaCerficateResponseBody {
	s.OssExpiration = &v
	return s
}

type GetMediaCerficateResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMediaCerficateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetMediaCerficateResponse) SetBody(v *GetMediaCerficateResponseBody) *GetMediaCerficateResponse {
	s.Body = v
	return s
}

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
	// 课程相关信息
	CourseInfo *CreateFeedRequestCourseInfo `json:"courseInfo,omitempty" xml:"courseInfo,omitempty" type:"Struct"`
	// 内容相关信息
	FeedInfo *CreateFeedRequestFeedInfo `json:"feedInfo,omitempty" xml:"feedInfo,omitempty" type:"Struct"`
	// 发布者的用户Id
	CreateUserId *string `json:"createUserId,omitempty" xml:"createUserId,omitempty"`
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

func (s *CreateFeedRequest) SetFeedInfo(v *CreateFeedRequestFeedInfo) *CreateFeedRequest {
	s.FeedInfo = v
	return s
}

func (s *CreateFeedRequest) SetCreateUserId(v string) *CreateFeedRequest {
	s.CreateUserId = &v
	return s
}

type CreateFeedRequestCourseInfo struct {
	// 创建一个和该课程绑定的学习群和圈子
	StudyGroupName *string `json:"studyGroupName,omitempty" xml:"studyGroupName,omitempty"`
	// 讲师身份信息
	LectorUserInfo *CreateFeedRequestCourseInfoLectorUserInfo `json:"lectorUserInfo,omitempty" xml:"lectorUserInfo,omitempty" type:"Struct"`
	// 支付信息
	PayInfo *CreateFeedRequestCourseInfoPayInfo `json:"payInfo,omitempty" xml:"payInfo,omitempty" type:"Struct"`
}

func (s CreateFeedRequestCourseInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedRequestCourseInfo) GoString() string {
	return s.String()
}

func (s *CreateFeedRequestCourseInfo) SetStudyGroupName(v string) *CreateFeedRequestCourseInfo {
	s.StudyGroupName = &v
	return s
}

func (s *CreateFeedRequestCourseInfo) SetLectorUserInfo(v *CreateFeedRequestCourseInfoLectorUserInfo) *CreateFeedRequestCourseInfo {
	s.LectorUserInfo = v
	return s
}

func (s *CreateFeedRequestCourseInfo) SetPayInfo(v *CreateFeedRequestCourseInfoPayInfo) *CreateFeedRequestCourseInfo {
	s.PayInfo = v
	return s
}

type CreateFeedRequestCourseInfoLectorUserInfo struct {
	// 讲师头像链接
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// 讲师用户Id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 讲师用户名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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

func (s *CreateFeedRequestCourseInfoLectorUserInfo) SetUserId(v string) *CreateFeedRequestCourseInfoLectorUserInfo {
	s.UserId = &v
	return s
}

func (s *CreateFeedRequestCourseInfoLectorUserInfo) SetName(v string) *CreateFeedRequestCourseInfoLectorUserInfo {
	s.Name = &v
	return s
}

type CreateFeedRequestCourseInfoPayInfo struct {
	// 客服身份信息
	CsUserInfo *CreateFeedRequestCourseInfoPayInfoCsUserInfo `json:"csUserInfo,omitempty" xml:"csUserInfo,omitempty" type:"Struct"`
	// 商品的默认情况下非打折时的价格，单位为分
	Price *int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 课程打折信息
	DiscountInfo *CreateFeedRequestCourseInfoPayInfoDiscountInfo `json:"discountInfo,omitempty" xml:"discountInfo,omitempty" type:"Struct"`
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

func (s *CreateFeedRequestCourseInfoPayInfo) SetPrice(v int64) *CreateFeedRequestCourseInfoPayInfo {
	s.Price = &v
	return s
}

func (s *CreateFeedRequestCourseInfoPayInfo) SetDiscountInfo(v *CreateFeedRequestCourseInfoPayInfoDiscountInfo) *CreateFeedRequestCourseInfoPayInfo {
	s.DiscountInfo = v
	return s
}

type CreateFeedRequestCourseInfoPayInfoCsUserInfo struct {
	// 客服头像链接
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// 客服用户Id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 客服用户名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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

func (s *CreateFeedRequestCourseInfoPayInfoCsUserInfo) SetUserId(v string) *CreateFeedRequestCourseInfoPayInfoCsUserInfo {
	s.UserId = &v
	return s
}

func (s *CreateFeedRequestCourseInfoPayInfoCsUserInfo) SetName(v string) *CreateFeedRequestCourseInfoPayInfoCsUserInfo {
	s.Name = &v
	return s
}

type CreateFeedRequestCourseInfoPayInfoDiscountInfo struct {
	// 打折开始的时间，时间戳精确到毫秒，时间为东八区时间
	StartTimeMillis *int64 `json:"startTimeMillis,omitempty" xml:"startTimeMillis,omitempty"`
	// 打折时商品的价格，单位为分
	Price *int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 打折结束的时间，时间戳精确到毫秒，时间为东八区时间
	EndTimeMillis *int64 `json:"endTimeMillis,omitempty" xml:"endTimeMillis,omitempty"`
}

func (s CreateFeedRequestCourseInfoPayInfoDiscountInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedRequestCourseInfoPayInfoDiscountInfo) GoString() string {
	return s.String()
}

func (s *CreateFeedRequestCourseInfoPayInfoDiscountInfo) SetStartTimeMillis(v int64) *CreateFeedRequestCourseInfoPayInfoDiscountInfo {
	s.StartTimeMillis = &v
	return s
}

func (s *CreateFeedRequestCourseInfoPayInfoDiscountInfo) SetPrice(v int64) *CreateFeedRequestCourseInfoPayInfoDiscountInfo {
	s.Price = &v
	return s
}

func (s *CreateFeedRequestCourseInfoPayInfoDiscountInfo) SetEndTimeMillis(v int64) *CreateFeedRequestCourseInfoPayInfoDiscountInfo {
	s.EndTimeMillis = &v
	return s
}

type CreateFeedRequestFeedInfo struct {
	// 请求的行为，是保存还是发布,1为save，2为publish，是修改还是新建取决于feedId是否为空
	ActionType *int32 `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// 一个课程下可以有多个视频或音频教程
	MediaContents []*CreateFeedRequestFeedInfoMediaContents `json:"mediaContents,omitempty" xml:"mediaContents,omitempty" type:"Repeated"`
	// 内容分类，课程分类列表详情请见附录中的列表
	FeedCategory *int64 `json:"feedCategory,omitempty" xml:"feedCategory,omitempty"`
	// 版权所属:1：自有， 2.代理， 3.钉钉
	BelongsTo *int32 `json:"belongsTo,omitempty" xml:"belongsTo,omitempty"`
	// 行业划分，决定了展示的页面的不同，例如学习中心下的职场、教育、商学院的划分,目前支持的行业id有：10001：职场学堂，10008：K12教育，10023：新职业，10024：钉钉培训
	IndustryId *int64 `json:"industryId,omitempty" xml:"industryId,omitempty"`
	// 课程的封面Url
	ThumbUrl *string `json:"thumbUrl,omitempty" xml:"thumbUrl,omitempty"`
	// 推荐内容集合
	Recommends []*CreateFeedRequestFeedInfoRecommends `json:"recommends,omitempty" xml:"recommends,omitempty" type:"Repeated"`
	// 可选参数，当feedId不为null时，表示对当前课程进行修改，否则为新建一个课程
	FeedId *string `json:"feedId,omitempty" xml:"feedId,omitempty"`
	// 内容的标题（标题不能重复）
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 内容类别,限制只能使用4种类型：0：免费 4：平价 5：专栏 6：训练营
	FeedType *int32 `json:"feedType,omitempty" xml:"feedType,omitempty"`
	// 课程的描述
	Introduction *string `json:"introduction,omitempty" xml:"introduction,omitempty"`
	// 课程的文字标签
	FeedTag *string `json:"feedTag,omitempty" xml:"feedTag,omitempty"`
	// 入驻账号Id(请联系钉钉接口同学获取)
	McnId *string `json:"mcnId,omitempty" xml:"mcnId,omitempty"`
	// 课程简介用的图片
	IntroductionPicUrl *string `json:"introductionPicUrl,omitempty" xml:"introductionPicUrl,omitempty"`
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

func (s *CreateFeedRequestFeedInfo) SetMediaContents(v []*CreateFeedRequestFeedInfoMediaContents) *CreateFeedRequestFeedInfo {
	s.MediaContents = v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetFeedCategory(v int64) *CreateFeedRequestFeedInfo {
	s.FeedCategory = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetBelongsTo(v int32) *CreateFeedRequestFeedInfo {
	s.BelongsTo = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetIndustryId(v int64) *CreateFeedRequestFeedInfo {
	s.IndustryId = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetThumbUrl(v string) *CreateFeedRequestFeedInfo {
	s.ThumbUrl = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetRecommends(v []*CreateFeedRequestFeedInfoRecommends) *CreateFeedRequestFeedInfo {
	s.Recommends = v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetFeedId(v string) *CreateFeedRequestFeedInfo {
	s.FeedId = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetTitle(v string) *CreateFeedRequestFeedInfo {
	s.Title = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetFeedType(v int32) *CreateFeedRequestFeedInfo {
	s.FeedType = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetIntroduction(v string) *CreateFeedRequestFeedInfo {
	s.Introduction = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetFeedTag(v string) *CreateFeedRequestFeedInfo {
	s.FeedTag = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetMcnId(v string) *CreateFeedRequestFeedInfo {
	s.McnId = &v
	return s
}

func (s *CreateFeedRequestFeedInfo) SetIntroductionPicUrl(v string) *CreateFeedRequestFeedInfo {
	s.IntroductionPicUrl = &v
	return s
}

type CreateFeedRequestFeedInfoMediaContents struct {
	// 媒体的类型，当前该接口只支持video和audio,2:视频,3:音频
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// 媒体的mediaId，唯一对应oss中的一个视频或者音频
	MediaId *string `json:"mediaId,omitempty" xml:"mediaId,omitempty"`
	// 媒体的标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateFeedRequestFeedInfoMediaContents) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedRequestFeedInfoMediaContents) GoString() string {
	return s.String()
}

func (s *CreateFeedRequestFeedInfoMediaContents) SetType(v int32) *CreateFeedRequestFeedInfoMediaContents {
	s.Type = &v
	return s
}

func (s *CreateFeedRequestFeedInfoMediaContents) SetMediaId(v string) *CreateFeedRequestFeedInfoMediaContents {
	s.MediaId = &v
	return s
}

func (s *CreateFeedRequestFeedInfoMediaContents) SetTitle(v string) *CreateFeedRequestFeedInfoMediaContents {
	s.Title = &v
	return s
}

type CreateFeedRequestFeedInfoRecommends struct {
	// 推荐物品的类别,0:课程,1:微应用
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// 推荐物品的id，可以时feedId或者是微应用Id
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
}

func (s CreateFeedRequestFeedInfoRecommends) String() string {
	return tea.Prettify(s)
}

func (s CreateFeedRequestFeedInfoRecommends) GoString() string {
	return s.String()
}

func (s *CreateFeedRequestFeedInfoRecommends) SetType(v int32) *CreateFeedRequestFeedInfoRecommends {
	s.Type = &v
	return s
}

func (s *CreateFeedRequestFeedInfoRecommends) SetObjectId(v string) *CreateFeedRequestFeedInfoRecommends {
	s.ObjectId = &v
	return s
}

type CreateFeedResponseBody struct {
	// 创建内容的内容Id
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
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFeedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateFeedResponse) SetBody(v *CreateFeedResponseBody) *CreateFeedResponse {
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

func (client *Client) GetMediaCerficateWithOptions(request *GetMediaCerficateRequest, headers *GetMediaCerficateHeaders, runtime *util.RuntimeOptions) (_result *GetMediaCerficateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ThumbUrl)) {
		query["thumbUrl"] = request.ThumbUrl
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.MediaId)) {
		query["mediaId"] = request.MediaId
	}

	if !tea.BoolValue(util.IsUnset(request.MediaTitle)) {
		query["mediaTitle"] = request.MediaTitle
	}

	if !tea.BoolValue(util.IsUnset(request.McnId)) {
		query["mcnId"] = request.McnId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.MediaIntroduction)) {
		query["mediaIntroduction"] = request.MediaIntroduction
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
	_result = &GetMediaCerficateResponse{}
	_body, _err := client.DoROARequest(tea.String("GetMediaCerficate"), tea.String("content_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/content/media/cerficates"), tea.String("json"), req, runtime)
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

func (client *Client) CreateFeedWithOptions(request *CreateFeedRequest, headers *CreateFeedHeaders, runtime *util.RuntimeOptions) (_result *CreateFeedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CourseInfo))) {
		body["courseInfo"] = request.CourseInfo
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.FeedInfo))) {
		body["feedInfo"] = request.FeedInfo
	}

	if !tea.BoolValue(util.IsUnset(request.CreateUserId)) {
		body["createUserId"] = request.CreateUserId
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
	_result = &CreateFeedResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateFeed"), tea.String("content_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/content/feeds"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
