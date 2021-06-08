// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package exclusive_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetGroupActiveInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetGroupActiveInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetGroupActiveInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetGroupActiveInfoHeaders) SetCommonHeaders(v map[string]*string) *GetGroupActiveInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetGroupActiveInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetGroupActiveInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetGroupActiveInfoRequest struct {
	// 统计日期
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
	// 钉钉群组id
	DingGroupId *string `json:"dingGroupId,omitempty" xml:"dingGroupId,omitempty"`
	// 分页起始页
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页大小
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetGroupActiveInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGroupActiveInfoRequest) GoString() string {
	return s.String()
}

func (s *GetGroupActiveInfoRequest) SetStatDate(v string) *GetGroupActiveInfoRequest {
	s.StatDate = &v
	return s
}

func (s *GetGroupActiveInfoRequest) SetDingGroupId(v string) *GetGroupActiveInfoRequest {
	s.DingGroupId = &v
	return s
}

func (s *GetGroupActiveInfoRequest) SetPageNumber(v int64) *GetGroupActiveInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *GetGroupActiveInfoRequest) SetPageSize(v int64) *GetGroupActiveInfoRequest {
	s.PageSize = &v
	return s
}

type GetGroupActiveInfoResponseBody struct {
	Data       []*GetGroupActiveInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	TotalCount *int64                                `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetGroupActiveInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGroupActiveInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupActiveInfoResponseBody) SetData(v []*GetGroupActiveInfoResponseBodyData) *GetGroupActiveInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetGroupActiveInfoResponseBody) SetTotalCount(v int64) *GetGroupActiveInfoResponseBody {
	s.TotalCount = &v
	return s
}

type GetGroupActiveInfoResponseBodyData struct {
	// 统计时间
	StatDate *string `json:"statDate,omitempty" xml:"statDate,omitempty"`
	// 群组id
	DingGroupId *string `json:"dingGroupId,omitempty" xml:"dingGroupId,omitempty"`
	// 群组创建时间
	GroupCreateTime *string `json:"groupCreateTime,omitempty" xml:"groupCreateTime,omitempty"`
	// 群组创建用户id
	GroupCreateUserId *string `json:"groupCreateUserId,omitempty" xml:"groupCreateUserId,omitempty"`
	// 群组创建用户姓名
	GroupCreateUserName *string `json:"groupCreateUserName,omitempty" xml:"groupCreateUserName,omitempty"`
	// 群名称
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// 群类型：1-全员群，2-部门群，3-（其他）内部群，4-场景群
	GroupType *int64 `json:"groupType,omitempty" xml:"groupType,omitempty"`
	// 最近1天群人数
	GroupUserCnt1d *int32 `json:"groupUserCnt1d,omitempty" xml:"groupUserCnt1d,omitempty"`
	// 最近1天发消息人数
	SendMessageUserCnt1d *int64 `json:"sendMessageUserCnt1d,omitempty" xml:"sendMessageUserCnt1d,omitempty"`
	// 最近1天发消息次数
	SendMessageCnt1d *int64 `json:"sendMessageCnt1d,omitempty" xml:"sendMessageCnt1d,omitempty"`
	// 最近1天打开群人数
	OpenConvUv1d *int32 `json:"openConvUv1d,omitempty" xml:"openConvUv1d,omitempty"`
}

func (s GetGroupActiveInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetGroupActiveInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGroupActiveInfoResponseBodyData) SetStatDate(v string) *GetGroupActiveInfoResponseBodyData {
	s.StatDate = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetDingGroupId(v string) *GetGroupActiveInfoResponseBodyData {
	s.DingGroupId = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetGroupCreateTime(v string) *GetGroupActiveInfoResponseBodyData {
	s.GroupCreateTime = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetGroupCreateUserId(v string) *GetGroupActiveInfoResponseBodyData {
	s.GroupCreateUserId = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetGroupCreateUserName(v string) *GetGroupActiveInfoResponseBodyData {
	s.GroupCreateUserName = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetGroupName(v string) *GetGroupActiveInfoResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetGroupType(v int64) *GetGroupActiveInfoResponseBodyData {
	s.GroupType = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetGroupUserCnt1d(v int32) *GetGroupActiveInfoResponseBodyData {
	s.GroupUserCnt1d = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetSendMessageUserCnt1d(v int64) *GetGroupActiveInfoResponseBodyData {
	s.SendMessageUserCnt1d = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetSendMessageCnt1d(v int64) *GetGroupActiveInfoResponseBodyData {
	s.SendMessageCnt1d = &v
	return s
}

func (s *GetGroupActiveInfoResponseBodyData) SetOpenConvUv1d(v int32) *GetGroupActiveInfoResponseBodyData {
	s.OpenConvUv1d = &v
	return s
}

type GetGroupActiveInfoResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetGroupActiveInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGroupActiveInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGroupActiveInfoResponse) GoString() string {
	return s.String()
}

func (s *GetGroupActiveInfoResponse) SetHeaders(v map[string]*string) *GetGroupActiveInfoResponse {
	s.Headers = v
	return s
}

func (s *GetGroupActiveInfoResponse) SetBody(v *GetGroupActiveInfoResponseBody) *GetGroupActiveInfoResponse {
	s.Body = v
	return s
}

type SearchOrgInnerGroupInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchOrgInnerGroupInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchOrgInnerGroupInfoHeaders) GoString() string {
	return s.String()
}

func (s *SearchOrgInnerGroupInfoHeaders) SetCommonHeaders(v map[string]*string) *SearchOrgInnerGroupInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchOrgInnerGroupInfoHeaders) SetXAcsDingtalkAccessToken(v string) *SearchOrgInnerGroupInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchOrgInnerGroupInfoRequest struct {
	// groupMembersCntEnd
	GroupMembersCountEnd *int32 `json:"groupMembersCountEnd,omitempty" xml:"groupMembersCountEnd,omitempty"`
	// syncToDingpan
	SyncToDingpan *int32 `json:"syncToDingpan,omitempty" xml:"syncToDingpan,omitempty"`
	// groupOwner
	GroupOwner *string `json:"groupOwner,omitempty" xml:"groupOwner,omitempty"`
	// createTimeEnd
	CreateTimeEnd *int64 `json:"createTimeEnd,omitempty" xml:"createTimeEnd,omitempty"`
	// pageSize
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// createTimeStart
	CreateTimeStart *int64 `json:"createTimeStart,omitempty" xml:"createTimeStart,omitempty"`
	// uuid
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// groupMembersCntStart
	GroupMembersCountStart *int32 `json:"groupMembersCountStart,omitempty" xml:"groupMembersCountStart,omitempty"`
	// lastActiveTimeEnd
	LastActiveTimeEnd *int64 `json:"lastActiveTimeEnd,omitempty" xml:"lastActiveTimeEnd,omitempty"`
	// operatorUserId
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// groupName
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// pageStart
	PageStart *int32 `json:"pageStart,omitempty" xml:"pageStart,omitempty"`
	// lastActiveTimeStart
	LastActiveTimeStart *int64 `json:"lastActiveTimeStart,omitempty" xml:"lastActiveTimeStart,omitempty"`
}

func (s SearchOrgInnerGroupInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchOrgInnerGroupInfoRequest) GoString() string {
	return s.String()
}

func (s *SearchOrgInnerGroupInfoRequest) SetGroupMembersCountEnd(v int32) *SearchOrgInnerGroupInfoRequest {
	s.GroupMembersCountEnd = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetSyncToDingpan(v int32) *SearchOrgInnerGroupInfoRequest {
	s.SyncToDingpan = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetGroupOwner(v string) *SearchOrgInnerGroupInfoRequest {
	s.GroupOwner = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetCreateTimeEnd(v int64) *SearchOrgInnerGroupInfoRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetPageSize(v int32) *SearchOrgInnerGroupInfoRequest {
	s.PageSize = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetCreateTimeStart(v int64) *SearchOrgInnerGroupInfoRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetUuid(v string) *SearchOrgInnerGroupInfoRequest {
	s.Uuid = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetGroupMembersCountStart(v int32) *SearchOrgInnerGroupInfoRequest {
	s.GroupMembersCountStart = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetLastActiveTimeEnd(v int64) *SearchOrgInnerGroupInfoRequest {
	s.LastActiveTimeEnd = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetOperatorUserId(v string) *SearchOrgInnerGroupInfoRequest {
	s.OperatorUserId = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetGroupName(v string) *SearchOrgInnerGroupInfoRequest {
	s.GroupName = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetPageStart(v int32) *SearchOrgInnerGroupInfoRequest {
	s.PageStart = &v
	return s
}

func (s *SearchOrgInnerGroupInfoRequest) SetLastActiveTimeStart(v int64) *SearchOrgInnerGroupInfoRequest {
	s.LastActiveTimeStart = &v
	return s
}

type SearchOrgInnerGroupInfoResponseBody struct {
	TotalCount *int64                                      `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	ItemCount  *int32                                      `json:"itemCount,omitempty" xml:"itemCount,omitempty"`
	Items      []*SearchOrgInnerGroupInfoResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s SearchOrgInnerGroupInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchOrgInnerGroupInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SearchOrgInnerGroupInfoResponseBody) SetTotalCount(v int64) *SearchOrgInnerGroupInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBody) SetItemCount(v int32) *SearchOrgInnerGroupInfoResponseBody {
	s.ItemCount = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBody) SetItems(v []*SearchOrgInnerGroupInfoResponseBodyItems) *SearchOrgInnerGroupInfoResponseBody {
	s.Items = v
	return s
}

type SearchOrgInnerGroupInfoResponseBodyItems struct {
	OpenConversationId      *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	GroupOwner              *string `json:"groupOwner,omitempty" xml:"groupOwner,omitempty"`
	GroupName               *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	GroupAdminsCount        *int32  `json:"groupAdminsCount,omitempty" xml:"groupAdminsCount,omitempty"`
	GroupMembersCount       *int32  `json:"groupMembersCount,omitempty" xml:"groupMembersCount,omitempty"`
	GroupCreateTime         *int64  `json:"groupCreateTime,omitempty" xml:"groupCreateTime,omitempty"`
	GroupLastActiveTime     *int64  `json:"groupLastActiveTime,omitempty" xml:"groupLastActiveTime,omitempty"`
	GroupLastActiveTimeShow *string `json:"groupLastActiveTimeShow,omitempty" xml:"groupLastActiveTimeShow,omitempty"`
	SyncToDingpan           *int32  `json:"syncToDingpan,omitempty" xml:"syncToDingpan,omitempty"`
	UsedQuota               *int64  `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
	GroupOwnerUserId        *string `json:"groupOwnerUserId,omitempty" xml:"groupOwnerUserId,omitempty"`
	Status                  *int32  `json:"status,omitempty" xml:"status,omitempty"`
	TemplateId              *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	TemplateName            *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
}

func (s SearchOrgInnerGroupInfoResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s SearchOrgInnerGroupInfoResponseBodyItems) GoString() string {
	return s.String()
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetOpenConversationId(v string) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.OpenConversationId = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetGroupOwner(v string) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.GroupOwner = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetGroupName(v string) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.GroupName = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetGroupAdminsCount(v int32) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.GroupAdminsCount = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetGroupMembersCount(v int32) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.GroupMembersCount = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetGroupCreateTime(v int64) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.GroupCreateTime = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetGroupLastActiveTime(v int64) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.GroupLastActiveTime = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetGroupLastActiveTimeShow(v string) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.GroupLastActiveTimeShow = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetSyncToDingpan(v int32) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.SyncToDingpan = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetUsedQuota(v int64) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.UsedQuota = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetGroupOwnerUserId(v string) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.GroupOwnerUserId = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetStatus(v int32) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.Status = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetTemplateId(v string) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.TemplateId = &v
	return s
}

func (s *SearchOrgInnerGroupInfoResponseBodyItems) SetTemplateName(v string) *SearchOrgInnerGroupInfoResponseBodyItems {
	s.TemplateName = &v
	return s
}

type SearchOrgInnerGroupInfoResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchOrgInnerGroupInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchOrgInnerGroupInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchOrgInnerGroupInfoResponse) GoString() string {
	return s.String()
}

func (s *SearchOrgInnerGroupInfoResponse) SetHeaders(v map[string]*string) *SearchOrgInnerGroupInfoResponse {
	s.Headers = v
	return s
}

func (s *SearchOrgInnerGroupInfoResponse) SetBody(v *SearchOrgInnerGroupInfoResponseBody) *SearchOrgInnerGroupInfoResponse {
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

func (client *Client) GetGroupActiveInfo(request *GetGroupActiveInfoRequest) (_result *GetGroupActiveInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetGroupActiveInfoHeaders{}
	_result = &GetGroupActiveInfoResponse{}
	_body, _err := client.GetGroupActiveInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGroupActiveInfoWithOptions(request *GetGroupActiveInfoRequest, headers *GetGroupActiveInfoHeaders, runtime *util.RuntimeOptions) (_result *GetGroupActiveInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatDate)) {
		query["statDate"] = request.StatDate
	}

	if !tea.BoolValue(util.IsUnset(request.DingGroupId)) {
		query["dingGroupId"] = request.DingGroupId
	}

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
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &GetGroupActiveInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetGroupActiveInfo"), tea.String("exclusive_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/exclusive/data/activeGroups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchOrgInnerGroupInfo(request *SearchOrgInnerGroupInfoRequest) (_result *SearchOrgInnerGroupInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchOrgInnerGroupInfoHeaders{}
	_result = &SearchOrgInnerGroupInfoResponse{}
	_body, _err := client.SearchOrgInnerGroupInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchOrgInnerGroupInfoWithOptions(request *SearchOrgInnerGroupInfoRequest, headers *SearchOrgInnerGroupInfoHeaders, runtime *util.RuntimeOptions) (_result *SearchOrgInnerGroupInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupMembersCountEnd)) {
		query["groupMembersCountEnd"] = request.GroupMembersCountEnd
	}

	if !tea.BoolValue(util.IsUnset(request.SyncToDingpan)) {
		query["syncToDingpan"] = request.SyncToDingpan
	}

	if !tea.BoolValue(util.IsUnset(request.GroupOwner)) {
		query["groupOwner"] = request.GroupOwner
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimeEnd)) {
		query["createTimeEnd"] = request.CreateTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimeStart)) {
		query["createTimeStart"] = request.CreateTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["uuid"] = request.Uuid
	}

	if !tea.BoolValue(util.IsUnset(request.GroupMembersCountStart)) {
		query["groupMembersCountStart"] = request.GroupMembersCountStart
	}

	if !tea.BoolValue(util.IsUnset(request.LastActiveTimeEnd)) {
		query["lastActiveTimeEnd"] = request.LastActiveTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		query["operatorUserId"] = request.OperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.PageStart)) {
		query["pageStart"] = request.PageStart
	}

	if !tea.BoolValue(util.IsUnset(request.LastActiveTimeStart)) {
		query["lastActiveTimeStart"] = request.LastActiveTimeStart
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
	_result = &SearchOrgInnerGroupInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("SearchOrgInnerGroupInfo"), tea.String("exclusive_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/exclusive/securities/orgGroupInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
