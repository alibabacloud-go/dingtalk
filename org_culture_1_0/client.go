// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package org_culture_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ConsumeUserPointsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ConsumeUserPointsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ConsumeUserPointsHeaders) GoString() string {
	return s.String()
}

func (s *ConsumeUserPointsHeaders) SetCommonHeaders(v map[string]*string) *ConsumeUserPointsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ConsumeUserPointsHeaders) SetXAcsDingtalkAccessToken(v string) *ConsumeUserPointsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ConsumeUserPointsRequest struct {
	// 扣减积分数量，1～1000000
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 幂等外部ID，最大长度32个字符
	OutId *string `json:"outId,omitempty" xml:"outId,omitempty"`
	// 备注，最长32个字符
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s ConsumeUserPointsRequest) String() string {
	return tea.Prettify(s)
}

func (s ConsumeUserPointsRequest) GoString() string {
	return s.String()
}

func (s *ConsumeUserPointsRequest) SetAmount(v int64) *ConsumeUserPointsRequest {
	s.Amount = &v
	return s
}

func (s *ConsumeUserPointsRequest) SetOutId(v string) *ConsumeUserPointsRequest {
	s.OutId = &v
	return s
}

func (s *ConsumeUserPointsRequest) SetRemark(v string) *ConsumeUserPointsRequest {
	s.Remark = &v
	return s
}

type ConsumeUserPointsResponseBody struct {
	// 响应数据
	Result *ConsumeUserPointsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// 请求响应状态
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ConsumeUserPointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConsumeUserPointsResponseBody) GoString() string {
	return s.String()
}

func (s *ConsumeUserPointsResponseBody) SetResult(v *ConsumeUserPointsResponseBodyResult) *ConsumeUserPointsResponseBody {
	s.Result = v
	return s
}

func (s *ConsumeUserPointsResponseBody) SetSuccess(v bool) *ConsumeUserPointsResponseBody {
	s.Success = &v
	return s
}

type ConsumeUserPointsResponseBodyResult struct {
	// 扣减后剩余积分数量
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

func (s ConsumeUserPointsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ConsumeUserPointsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ConsumeUserPointsResponseBodyResult) SetAmount(v int64) *ConsumeUserPointsResponseBodyResult {
	s.Amount = &v
	return s
}

type ConsumeUserPointsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConsumeUserPointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConsumeUserPointsResponse) String() string {
	return tea.Prettify(s)
}

func (s ConsumeUserPointsResponse) GoString() string {
	return s.String()
}

func (s *ConsumeUserPointsResponse) SetHeaders(v map[string]*string) *ConsumeUserPointsResponse {
	s.Headers = v
	return s
}

func (s *ConsumeUserPointsResponse) SetBody(v *ConsumeUserPointsResponseBody) *ConsumeUserPointsResponse {
	s.Body = v
	return s
}

type GrantHonorHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GrantHonorHeaders) String() string {
	return tea.Prettify(s)
}

func (s GrantHonorHeaders) GoString() string {
	return s.String()
}

func (s *GrantHonorHeaders) SetCommonHeaders(v map[string]*string) *GrantHonorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GrantHonorHeaders) SetXAcsDingtalkAccessToken(v string) *GrantHonorHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GrantHonorRequest struct {
	// 有效期到期时间 时间戳. 会处理成到期那天的23:59:59秒的时间戳
	ExpirationTime *int64 `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
	// 颁奖词，最多可以填50字
	GrantReason *string `json:"grantReason,omitempty" xml:"grantReason,omitempty"`
	// 颁奖人名字，最多15个字
	GranterName *string `json:"granterName,omitempty" xml:"granterName,omitempty"`
	// 是否使用官宣号发送内网动态
	NoticeAnnouncer *bool `json:"noticeAnnouncer,omitempty" xml:"noticeAnnouncer,omitempty"`
	// 是否触达单聊会话通知
	NoticeSingle *bool `json:"noticeSingle,omitempty" xml:"noticeSingle,omitempty"`
	// 接受人userId
	ReceiverUserIds []*string `json:"receiverUserIds,omitempty" xml:"receiverUserIds,omitempty" type:"Repeated"`
	// 发送人userId
	SenderUserId *string `json:"senderUserId,omitempty" xml:"senderUserId,omitempty"`
}

func (s GrantHonorRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantHonorRequest) GoString() string {
	return s.String()
}

func (s *GrantHonorRequest) SetExpirationTime(v int64) *GrantHonorRequest {
	s.ExpirationTime = &v
	return s
}

func (s *GrantHonorRequest) SetGrantReason(v string) *GrantHonorRequest {
	s.GrantReason = &v
	return s
}

func (s *GrantHonorRequest) SetGranterName(v string) *GrantHonorRequest {
	s.GranterName = &v
	return s
}

func (s *GrantHonorRequest) SetNoticeAnnouncer(v bool) *GrantHonorRequest {
	s.NoticeAnnouncer = &v
	return s
}

func (s *GrantHonorRequest) SetNoticeSingle(v bool) *GrantHonorRequest {
	s.NoticeSingle = &v
	return s
}

func (s *GrantHonorRequest) SetReceiverUserIds(v []*string) *GrantHonorRequest {
	s.ReceiverUserIds = v
	return s
}

func (s *GrantHonorRequest) SetSenderUserId(v string) *GrantHonorRequest {
	s.SenderUserId = &v
	return s
}

type GrantHonorResponseBody struct {
	Result  *GrantHonorResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GrantHonorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantHonorResponseBody) GoString() string {
	return s.String()
}

func (s *GrantHonorResponseBody) SetResult(v *GrantHonorResponseBodyResult) *GrantHonorResponseBody {
	s.Result = v
	return s
}

func (s *GrantHonorResponseBody) SetSuccess(v bool) *GrantHonorResponseBody {
	s.Success = &v
	return s
}

type GrantHonorResponseBodyResult struct {
	// 失败的userId
	FailedUserIds []*string `json:"failedUserIds,omitempty" xml:"failedUserIds,omitempty" type:"Repeated"`
	// 成功的userId
	SuccessUserIds []*string `json:"successUserIds,omitempty" xml:"successUserIds,omitempty" type:"Repeated"`
}

func (s GrantHonorResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GrantHonorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GrantHonorResponseBodyResult) SetFailedUserIds(v []*string) *GrantHonorResponseBodyResult {
	s.FailedUserIds = v
	return s
}

func (s *GrantHonorResponseBodyResult) SetSuccessUserIds(v []*string) *GrantHonorResponseBodyResult {
	s.SuccessUserIds = v
	return s
}

type GrantHonorResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GrantHonorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GrantHonorResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantHonorResponse) GoString() string {
	return s.String()
}

func (s *GrantHonorResponse) SetHeaders(v map[string]*string) *GrantHonorResponse {
	s.Headers = v
	return s
}

func (s *GrantHonorResponse) SetBody(v *GrantHonorResponseBody) *GrantHonorResponse {
	s.Body = v
	return s
}

type QueryCorpPointsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryCorpPointsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryCorpPointsHeaders) GoString() string {
	return s.String()
}

func (s *QueryCorpPointsHeaders) SetCommonHeaders(v map[string]*string) *QueryCorpPointsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryCorpPointsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryCorpPointsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryCorpPointsRequest struct {
	// 操作用户ID
	OptUserId *string `json:"optUserId,omitempty" xml:"optUserId,omitempty"`
}

func (s QueryCorpPointsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCorpPointsRequest) GoString() string {
	return s.String()
}

func (s *QueryCorpPointsRequest) SetOptUserId(v string) *QueryCorpPointsRequest {
	s.OptUserId = &v
	return s
}

type QueryCorpPointsResponseBody struct {
	// 响应数据
	Result *QueryCorpPointsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// 请求响应状态
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryCorpPointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCorpPointsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCorpPointsResponseBody) SetResult(v *QueryCorpPointsResponseBodyResult) *QueryCorpPointsResponseBody {
	s.Result = v
	return s
}

func (s *QueryCorpPointsResponseBody) SetSuccess(v bool) *QueryCorpPointsResponseBody {
	s.Success = &v
	return s
}

type QueryCorpPointsResponseBodyResult struct {
	// 企业员工可用于兑换积分总额
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

func (s QueryCorpPointsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryCorpPointsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryCorpPointsResponseBodyResult) SetAmount(v int64) *QueryCorpPointsResponseBodyResult {
	s.Amount = &v
	return s
}

type QueryCorpPointsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCorpPointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCorpPointsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCorpPointsResponse) GoString() string {
	return s.String()
}

func (s *QueryCorpPointsResponse) SetHeaders(v map[string]*string) *QueryCorpPointsResponse {
	s.Headers = v
	return s
}

func (s *QueryCorpPointsResponse) SetBody(v *QueryCorpPointsResponseBody) *QueryCorpPointsResponse {
	s.Body = v
	return s
}

type QueryOrgHonorsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryOrgHonorsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgHonorsHeaders) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsHeaders) SetCommonHeaders(v map[string]*string) *QueryOrgHonorsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOrgHonorsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryOrgHonorsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryOrgHonorsRequest struct {
	// 分页获取数据时，数据的数量，默认为20，最大可传入100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 分页获取数据的标记，第一页调用时传0，非第一页传入上次调用本接口返回值中的nextToken
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryOrgHonorsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgHonorsRequest) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsRequest) SetMaxResults(v int32) *QueryOrgHonorsRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryOrgHonorsRequest) SetNextToken(v string) *QueryOrgHonorsRequest {
	s.NextToken = &v
	return s
}

type QueryOrgHonorsResponseBody struct {
	Result  *QueryOrgHonorsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryOrgHonorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgHonorsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsResponseBody) SetResult(v *QueryOrgHonorsResponseBodyResult) *QueryOrgHonorsResponseBody {
	s.Result = v
	return s
}

func (s *QueryOrgHonorsResponseBody) SetSuccess(v bool) *QueryOrgHonorsResponseBody {
	s.Success = &v
	return s
}

type QueryOrgHonorsResponseBodyResult struct {
	// 下次获取数据的游标
	NextToken  *string                                       `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OpenHonors []*QueryOrgHonorsResponseBodyResultOpenHonors `json:"openHonors,omitempty" xml:"openHonors,omitempty" type:"Repeated"`
}

func (s QueryOrgHonorsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgHonorsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsResponseBodyResult) SetNextToken(v string) *QueryOrgHonorsResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *QueryOrgHonorsResponseBodyResult) SetOpenHonors(v []*QueryOrgHonorsResponseBodyResultOpenHonors) *QueryOrgHonorsResponseBodyResult {
	s.OpenHonors = v
	return s
}

type QueryOrgHonorsResponseBodyResultOpenHonors struct {
	// 荣誉含义
	HonorDesc *string `json:"honorDesc,omitempty" xml:"honorDesc,omitempty"`
	// 荣誉id
	HonorId *int64 `json:"honorId,omitempty" xml:"honorId,omitempty"`
	// 荣誉图片url
	HonorImgUrl *string `json:"honorImgUrl,omitempty" xml:"honorImgUrl,omitempty"`
	// 荣誉名字
	HonorName *string `json:"honorName,omitempty" xml:"honorName,omitempty"`
	// 荣誉附赠的挂件图url
	HonorPendantImgUrl *string `json:"honorPendantImgUrl,omitempty" xml:"honorPendantImgUrl,omitempty"`
}

func (s QueryOrgHonorsResponseBodyResultOpenHonors) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgHonorsResponseBodyResultOpenHonors) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsResponseBodyResultOpenHonors) SetHonorDesc(v string) *QueryOrgHonorsResponseBodyResultOpenHonors {
	s.HonorDesc = &v
	return s
}

func (s *QueryOrgHonorsResponseBodyResultOpenHonors) SetHonorId(v int64) *QueryOrgHonorsResponseBodyResultOpenHonors {
	s.HonorId = &v
	return s
}

func (s *QueryOrgHonorsResponseBodyResultOpenHonors) SetHonorImgUrl(v string) *QueryOrgHonorsResponseBodyResultOpenHonors {
	s.HonorImgUrl = &v
	return s
}

func (s *QueryOrgHonorsResponseBodyResultOpenHonors) SetHonorName(v string) *QueryOrgHonorsResponseBodyResultOpenHonors {
	s.HonorName = &v
	return s
}

func (s *QueryOrgHonorsResponseBodyResultOpenHonors) SetHonorPendantImgUrl(v string) *QueryOrgHonorsResponseBodyResultOpenHonors {
	s.HonorPendantImgUrl = &v
	return s
}

type QueryOrgHonorsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryOrgHonorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOrgHonorsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrgHonorsResponse) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsResponse) SetHeaders(v map[string]*string) *QueryOrgHonorsResponse {
	s.Headers = v
	return s
}

func (s *QueryOrgHonorsResponse) SetBody(v *QueryOrgHonorsResponseBody) *QueryOrgHonorsResponse {
	s.Body = v
	return s
}

type QueryUserHonorsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserHonorsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserHonorsHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsHeaders) SetCommonHeaders(v map[string]*string) *QueryUserHonorsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserHonorsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserHonorsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserHonorsRequest struct {
	// 查询数据的条数，默认查询20条，最大可传100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 分页查询的标记，查询第一页时传0，非第一页时传入上次调用本接口返回值中的nextToken
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryUserHonorsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserHonorsRequest) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsRequest) SetMaxResults(v int32) *QueryUserHonorsRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryUserHonorsRequest) SetNextToken(v string) *QueryUserHonorsRequest {
	s.NextToken = &v
	return s
}

type QueryUserHonorsResponseBody struct {
	Result  *QueryUserHonorsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                              `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryUserHonorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserHonorsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsResponseBody) SetResult(v *QueryUserHonorsResponseBodyResult) *QueryUserHonorsResponseBody {
	s.Result = v
	return s
}

func (s *QueryUserHonorsResponseBody) SetSuccess(v bool) *QueryUserHonorsResponseBody {
	s.Success = &v
	return s
}

type QueryUserHonorsResponseBodyResult struct {
	Honors    []*QueryUserHonorsResponseBodyResultHonors `json:"honors,omitempty" xml:"honors,omitempty" type:"Repeated"`
	NextToken *string                                    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryUserHonorsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryUserHonorsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsResponseBodyResult) SetHonors(v []*QueryUserHonorsResponseBodyResultHonors) *QueryUserHonorsResponseBodyResult {
	s.Honors = v
	return s
}

func (s *QueryUserHonorsResponseBodyResult) SetNextToken(v string) *QueryUserHonorsResponseBodyResult {
	s.NextToken = &v
	return s
}

type QueryUserHonorsResponseBodyResultHonors struct {
	// 有效期截止时间点，没有该属性则为永久生效
	ExpirationTime *int64 `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
	// 授予历史记录 包含用户及授予时间
	GrantHistory []*QueryUserHonorsResponseBodyResultHonorsGrantHistory `json:"grantHistory,omitempty" xml:"grantHistory,omitempty" type:"Repeated"`
	// 荣誉含义
	HonorDesc *string `json:"honorDesc,omitempty" xml:"honorDesc,omitempty"`
	// 必须。荣誉id
	HonorId *string `json:"honorId,omitempty" xml:"honorId,omitempty"`
	// 必须。荣誉名字
	HonorName *string `json:"honorName,omitempty" xml:"honorName,omitempty"`
}

func (s QueryUserHonorsResponseBodyResultHonors) String() string {
	return tea.Prettify(s)
}

func (s QueryUserHonorsResponseBodyResultHonors) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsResponseBodyResultHonors) SetExpirationTime(v int64) *QueryUserHonorsResponseBodyResultHonors {
	s.ExpirationTime = &v
	return s
}

func (s *QueryUserHonorsResponseBodyResultHonors) SetGrantHistory(v []*QueryUserHonorsResponseBodyResultHonorsGrantHistory) *QueryUserHonorsResponseBodyResultHonors {
	s.GrantHistory = v
	return s
}

func (s *QueryUserHonorsResponseBodyResultHonors) SetHonorDesc(v string) *QueryUserHonorsResponseBodyResultHonors {
	s.HonorDesc = &v
	return s
}

func (s *QueryUserHonorsResponseBodyResultHonors) SetHonorId(v string) *QueryUserHonorsResponseBodyResultHonors {
	s.HonorId = &v
	return s
}

func (s *QueryUserHonorsResponseBodyResultHonors) SetHonorName(v string) *QueryUserHonorsResponseBodyResultHonors {
	s.HonorName = &v
	return s
}

type QueryUserHonorsResponseBodyResultHonorsGrantHistory struct {
	// 授予时间 时间戳
	GrantTime *int64 `json:"grantTime,omitempty" xml:"grantTime,omitempty"`
	// 必须。荣誉发放人userid
	SenderUserid *string `json:"senderUserid,omitempty" xml:"senderUserid,omitempty"`
}

func (s QueryUserHonorsResponseBodyResultHonorsGrantHistory) String() string {
	return tea.Prettify(s)
}

func (s QueryUserHonorsResponseBodyResultHonorsGrantHistory) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsResponseBodyResultHonorsGrantHistory) SetGrantTime(v int64) *QueryUserHonorsResponseBodyResultHonorsGrantHistory {
	s.GrantTime = &v
	return s
}

func (s *QueryUserHonorsResponseBodyResultHonorsGrantHistory) SetSenderUserid(v string) *QueryUserHonorsResponseBodyResultHonorsGrantHistory {
	s.SenderUserid = &v
	return s
}

type QueryUserHonorsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryUserHonorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserHonorsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserHonorsResponse) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsResponse) SetHeaders(v map[string]*string) *QueryUserHonorsResponse {
	s.Headers = v
	return s
}

func (s *QueryUserHonorsResponse) SetBody(v *QueryUserHonorsResponseBody) *QueryUserHonorsResponse {
	s.Body = v
	return s
}

type QueryUserPointsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserPointsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserPointsHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserPointsHeaders) SetCommonHeaders(v map[string]*string) *QueryUserPointsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserPointsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserPointsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserPointsResponseBody struct {
	// 响应数据
	Result *QueryUserPointsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// 请求响应状态
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryUserPointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserPointsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserPointsResponseBody) SetResult(v *QueryUserPointsResponseBodyResult) *QueryUserPointsResponseBody {
	s.Result = v
	return s
}

func (s *QueryUserPointsResponseBody) SetSuccess(v bool) *QueryUserPointsResponseBody {
	s.Success = &v
	return s
}

type QueryUserPointsResponseBodyResult struct {
	// 员工积分数量
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

func (s QueryUserPointsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryUserPointsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryUserPointsResponseBodyResult) SetAmount(v int64) *QueryUserPointsResponseBodyResult {
	s.Amount = &v
	return s
}

type QueryUserPointsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryUserPointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserPointsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserPointsResponse) GoString() string {
	return s.String()
}

func (s *QueryUserPointsResponse) SetHeaders(v map[string]*string) *QueryUserPointsResponse {
	s.Headers = v
	return s
}

func (s *QueryUserPointsResponse) SetBody(v *QueryUserPointsResponseBody) *QueryUserPointsResponse {
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

func (client *Client) ConsumeUserPoints(userId *string, request *ConsumeUserPointsRequest) (_result *ConsumeUserPointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ConsumeUserPointsHeaders{}
	_result = &ConsumeUserPointsResponse{}
	_body, _err := client.ConsumeUserPointsWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConsumeUserPointsWithOptions(userId *string, request *ConsumeUserPointsRequest, headers *ConsumeUserPointsHeaders, runtime *util.RuntimeOptions) (_result *ConsumeUserPointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		body["amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		body["outId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
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
	_result = &ConsumeUserPointsResponse{}
	_body, _err := client.DoROARequest(tea.String("ConsumeUserPoints"), tea.String("orgCulture_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/orgCulture/users/"+tea.StringValue(userId)+"/points/deduct"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GrantHonor(honorId *string, request *GrantHonorRequest) (_result *GrantHonorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GrantHonorHeaders{}
	_result = &GrantHonorResponse{}
	_body, _err := client.GrantHonorWithOptions(honorId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GrantHonorWithOptions(honorId *string, request *GrantHonorRequest, headers *GrantHonorHeaders, runtime *util.RuntimeOptions) (_result *GrantHonorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	honorId = openapiutil.GetEncodeParam(honorId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExpirationTime)) {
		body["expirationTime"] = request.ExpirationTime
	}

	if !tea.BoolValue(util.IsUnset(request.GrantReason)) {
		body["grantReason"] = request.GrantReason
	}

	if !tea.BoolValue(util.IsUnset(request.GranterName)) {
		body["granterName"] = request.GranterName
	}

	if !tea.BoolValue(util.IsUnset(request.NoticeAnnouncer)) {
		body["noticeAnnouncer"] = request.NoticeAnnouncer
	}

	if !tea.BoolValue(util.IsUnset(request.NoticeSingle)) {
		body["noticeSingle"] = request.NoticeSingle
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverUserIds)) {
		body["receiverUserIds"] = request.ReceiverUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.SenderUserId)) {
		body["senderUserId"] = request.SenderUserId
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
	_result = &GrantHonorResponse{}
	_body, _err := client.DoROARequest(tea.String("GrantHonor"), tea.String("orgCulture_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/orgCulture/honors/"+tea.StringValue(honorId)+"/grant"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCorpPoints(request *QueryCorpPointsRequest) (_result *QueryCorpPointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryCorpPointsHeaders{}
	_result = &QueryCorpPointsResponse{}
	_body, _err := client.QueryCorpPointsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCorpPointsWithOptions(request *QueryCorpPointsRequest, headers *QueryCorpPointsHeaders, runtime *util.RuntimeOptions) (_result *QueryCorpPointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OptUserId)) {
		query["optUserId"] = request.OptUserId
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
	_result = &QueryCorpPointsResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryCorpPoints"), tea.String("orgCulture_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/orgCulture/organizations/points"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrgHonors(request *QueryOrgHonorsRequest) (_result *QueryOrgHonorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryOrgHonorsHeaders{}
	_result = &QueryOrgHonorsResponse{}
	_body, _err := client.QueryOrgHonorsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrgHonorsWithOptions(request *QueryOrgHonorsRequest, headers *QueryOrgHonorsHeaders, runtime *util.RuntimeOptions) (_result *QueryOrgHonorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
	_result = &QueryOrgHonorsResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryOrgHonors"), tea.String("orgCulture_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/orgCulture/organizations/honors"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserHonors(userId *string, request *QueryUserHonorsRequest) (_result *QueryUserHonorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserHonorsHeaders{}
	_result = &QueryUserHonorsResponse{}
	_body, _err := client.QueryUserHonorsWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserHonorsWithOptions(userId *string, request *QueryUserHonorsRequest, headers *QueryUserHonorsHeaders, runtime *util.RuntimeOptions) (_result *QueryUserHonorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
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
	_result = &QueryUserHonorsResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryUserHonors"), tea.String("orgCulture_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/orgCulture/honors/users/"+tea.StringValue(userId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserPoints(userId *string) (_result *QueryUserPointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserPointsHeaders{}
	_result = &QueryUserPointsResponse{}
	_body, _err := client.QueryUserPointsWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserPointsWithOptions(userId *string, headers *QueryUserPointsHeaders, runtime *util.RuntimeOptions) (_result *QueryUserPointsResponse, _err error) {
	userId = openapiutil.GetEncodeParam(userId)
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
	_result = &QueryUserPointsResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryUserPoints"), tea.String("orgCulture_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/orgCulture/users/"+tea.StringValue(userId)+"/points"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
