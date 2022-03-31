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
	ExpirationTime *string `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
	// 颁奖词，最多可以填50字
	GrantReason *string `json:"grantReason,omitempty" xml:"grantReason,omitempty"`
	// 颁奖人名字，最多15个字
	GranterName *string `json:"granterName,omitempty" xml:"granterName,omitempty"`
	// 是否使用官宣号通知获奖人
	NoticeAnnouncer *bool `json:"noticeAnnouncer,omitempty" xml:"noticeAnnouncer,omitempty"`
	// 是否触达单聊会话通知
	NoticeSingle *bool `json:"noticeSingle,omitempty" xml:"noticeSingle,omitempty"`
	// 接受人staffId
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

func (s *GrantHonorRequest) SetExpirationTime(v string) *GrantHonorRequest {
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
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
