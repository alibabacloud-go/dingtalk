// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package search_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BatchInsertSearchItemHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchInsertSearchItemHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertSearchItemHeaders) GoString() string {
	return s.String()
}

func (s *BatchInsertSearchItemHeaders) SetCommonHeaders(v map[string]*string) *BatchInsertSearchItemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchInsertSearchItemHeaders) SetXAcsDingtalkAccessToken(v string) *BatchInsertSearchItemHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchInsertSearchItemRequest struct {
	SearchItemModels []*BatchInsertSearchItemRequestSearchItemModels `json:"searchItemModels,omitempty" xml:"searchItemModels,omitempty" type:"Repeated"`
}

func (s BatchInsertSearchItemRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertSearchItemRequest) GoString() string {
	return s.String()
}

func (s *BatchInsertSearchItemRequest) SetSearchItemModels(v []*BatchInsertSearchItemRequestSearchItemModels) *BatchInsertSearchItemRequest {
	s.SearchItemModels = v
	return s
}

type BatchInsertSearchItemRequestSearchItemModels struct {
	// 数据项的脚注，长度不能超过64
	Footer *string `json:"footer,omitempty" xml:"footer,omitempty"`
	// 数据项的头像，长度不能超过512
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 数据项的id，tabId和orgId相同的情况下，itemId唯一标识一条数据项，长度不能超过128
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// 数据项的移动端跳转url地址，若同时配置默认url和mobileUrl，移动端跳转链接优先使用mobileUrl
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	// 数据项的PC端跳转url地址，若同时配置默认url和pcUrl，PC端跳转链接优先使用pcUrl
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	// 数据项的摘要，长度不能超过64
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 数据项的标题，长度不能超过16
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 数据项的默认url地址，若mobileUrl或pcUrl没有配置，则使用该url地址，默认url和mobileUrl、pcUrl至少配置其中一个
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s BatchInsertSearchItemRequestSearchItemModels) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertSearchItemRequestSearchItemModels) GoString() string {
	return s.String()
}

func (s *BatchInsertSearchItemRequestSearchItemModels) SetFooter(v string) *BatchInsertSearchItemRequestSearchItemModels {
	s.Footer = &v
	return s
}

func (s *BatchInsertSearchItemRequestSearchItemModels) SetIcon(v string) *BatchInsertSearchItemRequestSearchItemModels {
	s.Icon = &v
	return s
}

func (s *BatchInsertSearchItemRequestSearchItemModels) SetItemId(v string) *BatchInsertSearchItemRequestSearchItemModels {
	s.ItemId = &v
	return s
}

func (s *BatchInsertSearchItemRequestSearchItemModels) SetMobileUrl(v string) *BatchInsertSearchItemRequestSearchItemModels {
	s.MobileUrl = &v
	return s
}

func (s *BatchInsertSearchItemRequestSearchItemModels) SetPcUrl(v string) *BatchInsertSearchItemRequestSearchItemModels {
	s.PcUrl = &v
	return s
}

func (s *BatchInsertSearchItemRequestSearchItemModels) SetSummary(v string) *BatchInsertSearchItemRequestSearchItemModels {
	s.Summary = &v
	return s
}

func (s *BatchInsertSearchItemRequestSearchItemModels) SetTitle(v string) *BatchInsertSearchItemRequestSearchItemModels {
	s.Title = &v
	return s
}

func (s *BatchInsertSearchItemRequestSearchItemModels) SetUrl(v string) *BatchInsertSearchItemRequestSearchItemModels {
	s.Url = &v
	return s
}

type BatchInsertSearchItemResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s BatchInsertSearchItemResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertSearchItemResponse) GoString() string {
	return s.String()
}

func (s *BatchInsertSearchItemResponse) SetHeaders(v map[string]*string) *BatchInsertSearchItemResponse {
	s.Headers = v
	return s
}

type CreateSearchTabHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateSearchTabHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateSearchTabHeaders) GoString() string {
	return s.String()
}

func (s *CreateSearchTabHeaders) SetCommonHeaders(v map[string]*string) *CreateSearchTabHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSearchTabHeaders) SetXAcsDingtalkAccessToken(v string) *CreateSearchTabHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateSearchTabRequest struct {
	// 数据源名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 数据源优先级，数值越小优先级越高
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// 数据来源,非必填,默认来源为钉钉搜索内部引擎
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 数据源状态，1表示上线，0表示下线
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateSearchTabRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSearchTabRequest) GoString() string {
	return s.String()
}

func (s *CreateSearchTabRequest) SetName(v string) *CreateSearchTabRequest {
	s.Name = &v
	return s
}

func (s *CreateSearchTabRequest) SetPriority(v int32) *CreateSearchTabRequest {
	s.Priority = &v
	return s
}

func (s *CreateSearchTabRequest) SetSource(v string) *CreateSearchTabRequest {
	s.Source = &v
	return s
}

func (s *CreateSearchTabRequest) SetStatus(v int32) *CreateSearchTabRequest {
	s.Status = &v
	return s
}

type CreateSearchTabResponseBody struct {
	// 数据源的id,范围为3000-4000
	TabId *int32 `json:"tabId,omitempty" xml:"tabId,omitempty"`
}

func (s CreateSearchTabResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSearchTabResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSearchTabResponseBody) SetTabId(v int32) *CreateSearchTabResponseBody {
	s.TabId = &v
	return s
}

type CreateSearchTabResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSearchTabResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSearchTabResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSearchTabResponse) GoString() string {
	return s.String()
}

func (s *CreateSearchTabResponse) SetHeaders(v map[string]*string) *CreateSearchTabResponse {
	s.Headers = v
	return s
}

func (s *CreateSearchTabResponse) SetBody(v *CreateSearchTabResponseBody) *CreateSearchTabResponse {
	s.Body = v
	return s
}

type DeleteSearchItemHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteSearchItemHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteSearchItemHeaders) GoString() string {
	return s.String()
}

func (s *DeleteSearchItemHeaders) SetCommonHeaders(v map[string]*string) *DeleteSearchItemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteSearchItemHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteSearchItemHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteSearchItemResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteSearchItemResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSearchItemResponse) GoString() string {
	return s.String()
}

func (s *DeleteSearchItemResponse) SetHeaders(v map[string]*string) *DeleteSearchItemResponse {
	s.Headers = v
	return s
}

type DeleteSearchTabHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteSearchTabHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteSearchTabHeaders) GoString() string {
	return s.String()
}

func (s *DeleteSearchTabHeaders) SetCommonHeaders(v map[string]*string) *DeleteSearchTabHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteSearchTabHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteSearchTabHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteSearchTabResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteSearchTabResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSearchTabResponse) GoString() string {
	return s.String()
}

func (s *DeleteSearchTabResponse) SetHeaders(v map[string]*string) *DeleteSearchTabResponse {
	s.Headers = v
	return s
}

type GetSearchItemHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSearchItemHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSearchItemHeaders) GoString() string {
	return s.String()
}

func (s *GetSearchItemHeaders) SetCommonHeaders(v map[string]*string) *GetSearchItemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSearchItemHeaders) SetXAcsDingtalkAccessToken(v string) *GetSearchItemHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSearchItemResponseBody struct {
	// 数据项的脚注
	Footer *string `json:"footer,omitempty" xml:"footer,omitempty"`
	// 创建时间
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 修改时间
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 数据项的头像
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 数据项的id,tabId和orgId相同的情况下，itemId唯一标识一条数据项
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// 数据项的移动端跳转url地址，若同时配置默认url和mobileUrl，移动端跳转链接优先使用mobileUrl
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	// 数据项的PC端跳转url地址，若同时配置默认url和pcUrl，PC端跳转链接优先使用pcUrl
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	// 数据项的摘要
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 数据源id
	TabId *int32 `json:"tabId,omitempty" xml:"tabId,omitempty"`
	// 数据项的标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 数据项的默认url地址，若mobileUrl或pcUrl没有配置，则使用该url地址，默认url和mobileUrl、pcUrl至少配置其中一个
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetSearchItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSearchItemResponseBody) GoString() string {
	return s.String()
}

func (s *GetSearchItemResponseBody) SetFooter(v string) *GetSearchItemResponseBody {
	s.Footer = &v
	return s
}

func (s *GetSearchItemResponseBody) SetGmtCreate(v string) *GetSearchItemResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetSearchItemResponseBody) SetGmtModified(v string) *GetSearchItemResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetSearchItemResponseBody) SetIcon(v string) *GetSearchItemResponseBody {
	s.Icon = &v
	return s
}

func (s *GetSearchItemResponseBody) SetItemId(v string) *GetSearchItemResponseBody {
	s.ItemId = &v
	return s
}

func (s *GetSearchItemResponseBody) SetMobileUrl(v string) *GetSearchItemResponseBody {
	s.MobileUrl = &v
	return s
}

func (s *GetSearchItemResponseBody) SetPcUrl(v string) *GetSearchItemResponseBody {
	s.PcUrl = &v
	return s
}

func (s *GetSearchItemResponseBody) SetSummary(v string) *GetSearchItemResponseBody {
	s.Summary = &v
	return s
}

func (s *GetSearchItemResponseBody) SetTabId(v int32) *GetSearchItemResponseBody {
	s.TabId = &v
	return s
}

func (s *GetSearchItemResponseBody) SetTitle(v string) *GetSearchItemResponseBody {
	s.Title = &v
	return s
}

func (s *GetSearchItemResponseBody) SetUrl(v string) *GetSearchItemResponseBody {
	s.Url = &v
	return s
}

type GetSearchItemResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSearchItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSearchItemResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSearchItemResponse) GoString() string {
	return s.String()
}

func (s *GetSearchItemResponse) SetHeaders(v map[string]*string) *GetSearchItemResponse {
	s.Headers = v
	return s
}

func (s *GetSearchItemResponse) SetBody(v *GetSearchItemResponseBody) *GetSearchItemResponse {
	s.Body = v
	return s
}

type GetSearchItemsByKeyWordHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSearchItemsByKeyWordHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSearchItemsByKeyWordHeaders) GoString() string {
	return s.String()
}

func (s *GetSearchItemsByKeyWordHeaders) SetCommonHeaders(v map[string]*string) *GetSearchItemsByKeyWordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSearchItemsByKeyWordHeaders) SetXAcsDingtalkAccessToken(v string) *GetSearchItemsByKeyWordHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSearchItemsByKeyWordRequest struct {
	// 搜索关键词
	KeyWord *string `json:"keyWord,omitempty" xml:"keyWord,omitempty"`
	// 一次性请求的item数量
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 加密偏移量，第一次请求取“0”值，后续请求根据接口返回的nextToken值进行填写
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetSearchItemsByKeyWordRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSearchItemsByKeyWordRequest) GoString() string {
	return s.String()
}

func (s *GetSearchItemsByKeyWordRequest) SetKeyWord(v string) *GetSearchItemsByKeyWordRequest {
	s.KeyWord = &v
	return s
}

func (s *GetSearchItemsByKeyWordRequest) SetMaxResults(v int32) *GetSearchItemsByKeyWordRequest {
	s.MaxResults = &v
	return s
}

func (s *GetSearchItemsByKeyWordRequest) SetNextToken(v string) *GetSearchItemsByKeyWordRequest {
	s.NextToken = &v
	return s
}

type GetSearchItemsByKeyWordResponseBody struct {
	// 下一次请求的加密offset，若为空则代表item已经读取完毕
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 本次请求条件下的item总量
	TotalCount *int32                                      `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	Value      []*GetSearchItemsByKeyWordResponseBodyValue `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s GetSearchItemsByKeyWordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSearchItemsByKeyWordResponseBody) GoString() string {
	return s.String()
}

func (s *GetSearchItemsByKeyWordResponseBody) SetNextToken(v string) *GetSearchItemsByKeyWordResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetSearchItemsByKeyWordResponseBody) SetTotalCount(v int32) *GetSearchItemsByKeyWordResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetSearchItemsByKeyWordResponseBody) SetValue(v []*GetSearchItemsByKeyWordResponseBodyValue) *GetSearchItemsByKeyWordResponseBody {
	s.Value = v
	return s
}

type GetSearchItemsByKeyWordResponseBodyValue struct {
	// 数据项的脚注
	Footer *string `json:"footer,omitempty" xml:"footer,omitempty"`
	// 创建时间
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 修改时间
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 数据项的头像
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 数据项的id,tabId和orgId相同的情况下，itemId唯一标识一条数据项
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// 数据项的移动端跳转url地址，若同时配置默认url和mobileUrl，移动端跳转链接优先使用mobileUrl
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	// 数据项的PC端跳转url地址，若同时配置默认url和pcUrl，PC端跳转链接优先使用pcUrl
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	// 数据项的摘要
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 数据源id
	TabId *int32 `json:"tabId,omitempty" xml:"tabId,omitempty"`
	// 数据项的标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 数据项的默认url地址，若mobileUrl或pcUrl没有配置，则使用该url地址，默认url和mobileUrl、pcUrl至少配置其中一个
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetSearchItemsByKeyWordResponseBodyValue) String() string {
	return tea.Prettify(s)
}

func (s GetSearchItemsByKeyWordResponseBodyValue) GoString() string {
	return s.String()
}

func (s *GetSearchItemsByKeyWordResponseBodyValue) SetFooter(v string) *GetSearchItemsByKeyWordResponseBodyValue {
	s.Footer = &v
	return s
}

func (s *GetSearchItemsByKeyWordResponseBodyValue) SetGmtCreate(v string) *GetSearchItemsByKeyWordResponseBodyValue {
	s.GmtCreate = &v
	return s
}

func (s *GetSearchItemsByKeyWordResponseBodyValue) SetGmtModified(v string) *GetSearchItemsByKeyWordResponseBodyValue {
	s.GmtModified = &v
	return s
}

func (s *GetSearchItemsByKeyWordResponseBodyValue) SetIcon(v string) *GetSearchItemsByKeyWordResponseBodyValue {
	s.Icon = &v
	return s
}

func (s *GetSearchItemsByKeyWordResponseBodyValue) SetItemId(v string) *GetSearchItemsByKeyWordResponseBodyValue {
	s.ItemId = &v
	return s
}

func (s *GetSearchItemsByKeyWordResponseBodyValue) SetMobileUrl(v string) *GetSearchItemsByKeyWordResponseBodyValue {
	s.MobileUrl = &v
	return s
}

func (s *GetSearchItemsByKeyWordResponseBodyValue) SetPcUrl(v string) *GetSearchItemsByKeyWordResponseBodyValue {
	s.PcUrl = &v
	return s
}

func (s *GetSearchItemsByKeyWordResponseBodyValue) SetSummary(v string) *GetSearchItemsByKeyWordResponseBodyValue {
	s.Summary = &v
	return s
}

func (s *GetSearchItemsByKeyWordResponseBodyValue) SetTabId(v int32) *GetSearchItemsByKeyWordResponseBodyValue {
	s.TabId = &v
	return s
}

func (s *GetSearchItemsByKeyWordResponseBodyValue) SetTitle(v string) *GetSearchItemsByKeyWordResponseBodyValue {
	s.Title = &v
	return s
}

func (s *GetSearchItemsByKeyWordResponseBodyValue) SetUrl(v string) *GetSearchItemsByKeyWordResponseBodyValue {
	s.Url = &v
	return s
}

type GetSearchItemsByKeyWordResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSearchItemsByKeyWordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSearchItemsByKeyWordResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSearchItemsByKeyWordResponse) GoString() string {
	return s.String()
}

func (s *GetSearchItemsByKeyWordResponse) SetHeaders(v map[string]*string) *GetSearchItemsByKeyWordResponse {
	s.Headers = v
	return s
}

func (s *GetSearchItemsByKeyWordResponse) SetBody(v *GetSearchItemsByKeyWordResponseBody) *GetSearchItemsByKeyWordResponse {
	s.Body = v
	return s
}

type GetSearchTabHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSearchTabHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSearchTabHeaders) GoString() string {
	return s.String()
}

func (s *GetSearchTabHeaders) SetCommonHeaders(v map[string]*string) *GetSearchTabHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSearchTabHeaders) SetXAcsDingtalkAccessToken(v string) *GetSearchTabHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSearchTabResponseBody struct {
	// 创建时间
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 修改时间
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 数据源名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 数据源优先级，数值越小优先级越高
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// 数据来源,非必填,默认来源为钉钉搜索内部引擎
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 数据源状态，1表示上线，0表示下线
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// 数据源的id,范围为3000-4000
	TabId *int32 `json:"tabId,omitempty" xml:"tabId,omitempty"`
}

func (s GetSearchTabResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSearchTabResponseBody) GoString() string {
	return s.String()
}

func (s *GetSearchTabResponseBody) SetGmtCreate(v string) *GetSearchTabResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetSearchTabResponseBody) SetGmtModified(v string) *GetSearchTabResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetSearchTabResponseBody) SetName(v string) *GetSearchTabResponseBody {
	s.Name = &v
	return s
}

func (s *GetSearchTabResponseBody) SetPriority(v int32) *GetSearchTabResponseBody {
	s.Priority = &v
	return s
}

func (s *GetSearchTabResponseBody) SetSource(v string) *GetSearchTabResponseBody {
	s.Source = &v
	return s
}

func (s *GetSearchTabResponseBody) SetStatus(v int32) *GetSearchTabResponseBody {
	s.Status = &v
	return s
}

func (s *GetSearchTabResponseBody) SetTabId(v int32) *GetSearchTabResponseBody {
	s.TabId = &v
	return s
}

type GetSearchTabResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSearchTabResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSearchTabResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSearchTabResponse) GoString() string {
	return s.String()
}

func (s *GetSearchTabResponse) SetHeaders(v map[string]*string) *GetSearchTabResponse {
	s.Headers = v
	return s
}

func (s *GetSearchTabResponse) SetBody(v *GetSearchTabResponseBody) *GetSearchTabResponse {
	s.Body = v
	return s
}

type InsertSearchItemHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InsertSearchItemHeaders) String() string {
	return tea.Prettify(s)
}

func (s InsertSearchItemHeaders) GoString() string {
	return s.String()
}

func (s *InsertSearchItemHeaders) SetCommonHeaders(v map[string]*string) *InsertSearchItemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsertSearchItemHeaders) SetXAcsDingtalkAccessToken(v string) *InsertSearchItemHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InsertSearchItemRequest struct {
	// 数据项的脚注，长度不能超过64
	Footer *string `json:"footer,omitempty" xml:"footer,omitempty"`
	// 数据项的头像，长度不能超过512
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 数据项的id，tabId和orgId相同的情况下，itemId唯一标识一条数据项，长度不能超过128
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// 数据项的移动端跳转url地址，若同时配置默认url和mobileUrl，移动端跳转链接优先使用mobileUrl
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	// 数据项的PC端跳转url地址，若同时配置默认url和pcUrl，PC端跳转链接优先使用pcUrl
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	// 数据项的摘要，长度不能超过64
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 数据项的标题，长度不能超过16
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 数据项的默认url地址，若mobileUrl或pcUrl没有配置，则使用该url地址，默认url和mobileUrl、pcUrl至少配置其中一个
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s InsertSearchItemRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertSearchItemRequest) GoString() string {
	return s.String()
}

func (s *InsertSearchItemRequest) SetFooter(v string) *InsertSearchItemRequest {
	s.Footer = &v
	return s
}

func (s *InsertSearchItemRequest) SetIcon(v string) *InsertSearchItemRequest {
	s.Icon = &v
	return s
}

func (s *InsertSearchItemRequest) SetItemId(v string) *InsertSearchItemRequest {
	s.ItemId = &v
	return s
}

func (s *InsertSearchItemRequest) SetMobileUrl(v string) *InsertSearchItemRequest {
	s.MobileUrl = &v
	return s
}

func (s *InsertSearchItemRequest) SetPcUrl(v string) *InsertSearchItemRequest {
	s.PcUrl = &v
	return s
}

func (s *InsertSearchItemRequest) SetSummary(v string) *InsertSearchItemRequest {
	s.Summary = &v
	return s
}

func (s *InsertSearchItemRequest) SetTitle(v string) *InsertSearchItemRequest {
	s.Title = &v
	return s
}

func (s *InsertSearchItemRequest) SetUrl(v string) *InsertSearchItemRequest {
	s.Url = &v
	return s
}

type InsertSearchItemResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s InsertSearchItemResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertSearchItemResponse) GoString() string {
	return s.String()
}

func (s *InsertSearchItemResponse) SetHeaders(v map[string]*string) *InsertSearchItemResponse {
	s.Headers = v
	return s
}

type ListSearchTabsByOrgIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListSearchTabsByOrgIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListSearchTabsByOrgIdHeaders) GoString() string {
	return s.String()
}

func (s *ListSearchTabsByOrgIdHeaders) SetCommonHeaders(v map[string]*string) *ListSearchTabsByOrgIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSearchTabsByOrgIdHeaders) SetXAcsDingtalkAccessToken(v string) *ListSearchTabsByOrgIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListSearchTabsByOrgIdResponseBody struct {
	// 该企业拥有的所有数据源信息
	SearchTabResult []*ListSearchTabsByOrgIdResponseBodySearchTabResult `json:"searchTabResult,omitempty" xml:"searchTabResult,omitempty" type:"Repeated"`
}

func (s ListSearchTabsByOrgIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSearchTabsByOrgIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListSearchTabsByOrgIdResponseBody) SetSearchTabResult(v []*ListSearchTabsByOrgIdResponseBodySearchTabResult) *ListSearchTabsByOrgIdResponseBody {
	s.SearchTabResult = v
	return s
}

type ListSearchTabsByOrgIdResponseBodySearchTabResult struct {
	// 创建时间
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 修改时间
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 数据源名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 数据源优先级，数值越小优先级越高
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// 数据来源,非必填,默认来源为钉钉搜索内部引擎
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 数据源状态，1表示上线，0表示下线
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// 数据源的id,范围为3000-4000
	TabId *int32 `json:"tabId,omitempty" xml:"tabId,omitempty"`
}

func (s ListSearchTabsByOrgIdResponseBodySearchTabResult) String() string {
	return tea.Prettify(s)
}

func (s ListSearchTabsByOrgIdResponseBodySearchTabResult) GoString() string {
	return s.String()
}

func (s *ListSearchTabsByOrgIdResponseBodySearchTabResult) SetGmtCreate(v string) *ListSearchTabsByOrgIdResponseBodySearchTabResult {
	s.GmtCreate = &v
	return s
}

func (s *ListSearchTabsByOrgIdResponseBodySearchTabResult) SetGmtModified(v string) *ListSearchTabsByOrgIdResponseBodySearchTabResult {
	s.GmtModified = &v
	return s
}

func (s *ListSearchTabsByOrgIdResponseBodySearchTabResult) SetName(v string) *ListSearchTabsByOrgIdResponseBodySearchTabResult {
	s.Name = &v
	return s
}

func (s *ListSearchTabsByOrgIdResponseBodySearchTabResult) SetPriority(v int32) *ListSearchTabsByOrgIdResponseBodySearchTabResult {
	s.Priority = &v
	return s
}

func (s *ListSearchTabsByOrgIdResponseBodySearchTabResult) SetSource(v string) *ListSearchTabsByOrgIdResponseBodySearchTabResult {
	s.Source = &v
	return s
}

func (s *ListSearchTabsByOrgIdResponseBodySearchTabResult) SetStatus(v int32) *ListSearchTabsByOrgIdResponseBodySearchTabResult {
	s.Status = &v
	return s
}

func (s *ListSearchTabsByOrgIdResponseBodySearchTabResult) SetTabId(v int32) *ListSearchTabsByOrgIdResponseBodySearchTabResult {
	s.TabId = &v
	return s
}

type ListSearchTabsByOrgIdResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSearchTabsByOrgIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSearchTabsByOrgIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSearchTabsByOrgIdResponse) GoString() string {
	return s.String()
}

func (s *ListSearchTabsByOrgIdResponse) SetHeaders(v map[string]*string) *ListSearchTabsByOrgIdResponse {
	s.Headers = v
	return s
}

func (s *ListSearchTabsByOrgIdResponse) SetBody(v *ListSearchTabsByOrgIdResponseBody) *ListSearchTabsByOrgIdResponse {
	s.Body = v
	return s
}

type UpdateSearchTabHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateSearchTabHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateSearchTabHeaders) GoString() string {
	return s.String()
}

func (s *UpdateSearchTabHeaders) SetCommonHeaders(v map[string]*string) *UpdateSearchTabHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateSearchTabHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateSearchTabHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateSearchTabRequest struct {
	// 数据源名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 数据源优先级，数值越小优先级越高
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// 数据来源,非必填,默认来源为钉钉搜索内部引擎
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 数据源状态，1表示上线，0表示下线
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateSearchTabRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSearchTabRequest) GoString() string {
	return s.String()
}

func (s *UpdateSearchTabRequest) SetName(v string) *UpdateSearchTabRequest {
	s.Name = &v
	return s
}

func (s *UpdateSearchTabRequest) SetPriority(v int32) *UpdateSearchTabRequest {
	s.Priority = &v
	return s
}

func (s *UpdateSearchTabRequest) SetSource(v string) *UpdateSearchTabRequest {
	s.Source = &v
	return s
}

func (s *UpdateSearchTabRequest) SetStatus(v int32) *UpdateSearchTabRequest {
	s.Status = &v
	return s
}

type UpdateSearchTabResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpdateSearchTabResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSearchTabResponse) GoString() string {
	return s.String()
}

func (s *UpdateSearchTabResponse) SetHeaders(v map[string]*string) *UpdateSearchTabResponse {
	s.Headers = v
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

func (client *Client) BatchInsertSearchItem(tabId *string, request *BatchInsertSearchItemRequest) (_result *BatchInsertSearchItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchInsertSearchItemHeaders{}
	_result = &BatchInsertSearchItemResponse{}
	_body, _err := client.BatchInsertSearchItemWithOptions(tabId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchInsertSearchItemWithOptions(tabId *string, request *BatchInsertSearchItemRequest, headers *BatchInsertSearchItemHeaders, runtime *util.RuntimeOptions) (_result *BatchInsertSearchItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	tabId = openapiutil.GetEncodeParam(tabId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SearchItemModels)) {
		body["searchItemModels"] = request.SearchItemModels
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
	_result = &BatchInsertSearchItemResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchInsertSearchItem"), tea.String("search_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/search/tabs/"+tea.StringValue(tabId)+"/items/batch"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSearchTab(request *CreateSearchTabRequest) (_result *CreateSearchTabResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateSearchTabHeaders{}
	_result = &CreateSearchTabResponse{}
	_body, _err := client.CreateSearchTabWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSearchTabWithOptions(request *CreateSearchTabRequest, headers *CreateSearchTabHeaders, runtime *util.RuntimeOptions) (_result *CreateSearchTabResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
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
	_result = &CreateSearchTabResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateSearchTab"), tea.String("search_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/search/tabs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSearchItem(tabId *string, itemId *string) (_result *DeleteSearchItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteSearchItemHeaders{}
	_result = &DeleteSearchItemResponse{}
	_body, _err := client.DeleteSearchItemWithOptions(tabId, itemId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSearchItemWithOptions(tabId *string, itemId *string, headers *DeleteSearchItemHeaders, runtime *util.RuntimeOptions) (_result *DeleteSearchItemResponse, _err error) {
	tabId = openapiutil.GetEncodeParam(tabId)
	itemId = openapiutil.GetEncodeParam(itemId)
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
	_result = &DeleteSearchItemResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteSearchItem"), tea.String("search_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/search/tabs/"+tea.StringValue(tabId)+"/items/"+tea.StringValue(itemId)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSearchTab(tabId *string) (_result *DeleteSearchTabResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteSearchTabHeaders{}
	_result = &DeleteSearchTabResponse{}
	_body, _err := client.DeleteSearchTabWithOptions(tabId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSearchTabWithOptions(tabId *string, headers *DeleteSearchTabHeaders, runtime *util.RuntimeOptions) (_result *DeleteSearchTabResponse, _err error) {
	tabId = openapiutil.GetEncodeParam(tabId)
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
	_result = &DeleteSearchTabResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteSearchTab"), tea.String("search_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/search/tabs/"+tea.StringValue(tabId)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSearchItem(tabId *string, itemId *string) (_result *GetSearchItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSearchItemHeaders{}
	_result = &GetSearchItemResponse{}
	_body, _err := client.GetSearchItemWithOptions(tabId, itemId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSearchItemWithOptions(tabId *string, itemId *string, headers *GetSearchItemHeaders, runtime *util.RuntimeOptions) (_result *GetSearchItemResponse, _err error) {
	tabId = openapiutil.GetEncodeParam(tabId)
	itemId = openapiutil.GetEncodeParam(itemId)
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
	_result = &GetSearchItemResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSearchItem"), tea.String("search_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/search/tabs/"+tea.StringValue(tabId)+"/items/"+tea.StringValue(itemId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSearchItemsByKeyWord(tabId *string, request *GetSearchItemsByKeyWordRequest) (_result *GetSearchItemsByKeyWordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSearchItemsByKeyWordHeaders{}
	_result = &GetSearchItemsByKeyWordResponse{}
	_body, _err := client.GetSearchItemsByKeyWordWithOptions(tabId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSearchItemsByKeyWordWithOptions(tabId *string, request *GetSearchItemsByKeyWordRequest, headers *GetSearchItemsByKeyWordHeaders, runtime *util.RuntimeOptions) (_result *GetSearchItemsByKeyWordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	tabId = openapiutil.GetEncodeParam(tabId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyWord)) {
		query["keyWord"] = request.KeyWord
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
	_result = &GetSearchItemsByKeyWordResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSearchItemsByKeyWord"), tea.String("search_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/search/tabs/"+tea.StringValue(tabId)+"/items"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSearchTab(tabId *string) (_result *GetSearchTabResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSearchTabHeaders{}
	_result = &GetSearchTabResponse{}
	_body, _err := client.GetSearchTabWithOptions(tabId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSearchTabWithOptions(tabId *string, headers *GetSearchTabHeaders, runtime *util.RuntimeOptions) (_result *GetSearchTabResponse, _err error) {
	tabId = openapiutil.GetEncodeParam(tabId)
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
	_result = &GetSearchTabResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSearchTab"), tea.String("search_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/search/tabs/"+tea.StringValue(tabId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertSearchItem(tabId *string, request *InsertSearchItemRequest) (_result *InsertSearchItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InsertSearchItemHeaders{}
	_result = &InsertSearchItemResponse{}
	_body, _err := client.InsertSearchItemWithOptions(tabId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertSearchItemWithOptions(tabId *string, request *InsertSearchItemRequest, headers *InsertSearchItemHeaders, runtime *util.RuntimeOptions) (_result *InsertSearchItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	tabId = openapiutil.GetEncodeParam(tabId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Footer)) {
		body["footer"] = request.Footer
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		body["icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.ItemId)) {
		body["itemId"] = request.ItemId
	}

	if !tea.BoolValue(util.IsUnset(request.MobileUrl)) {
		body["mobileUrl"] = request.MobileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.PcUrl)) {
		body["pcUrl"] = request.PcUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Summary)) {
		body["summary"] = request.Summary
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["url"] = request.Url
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
	_result = &InsertSearchItemResponse{}
	_body, _err := client.DoROARequest(tea.String("InsertSearchItem"), tea.String("search_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/search/tabs/"+tea.StringValue(tabId)+"/items"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSearchTabsByOrgId() (_result *ListSearchTabsByOrgIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListSearchTabsByOrgIdHeaders{}
	_result = &ListSearchTabsByOrgIdResponse{}
	_body, _err := client.ListSearchTabsByOrgIdWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSearchTabsByOrgIdWithOptions(headers *ListSearchTabsByOrgIdHeaders, runtime *util.RuntimeOptions) (_result *ListSearchTabsByOrgIdResponse, _err error) {
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
	_result = &ListSearchTabsByOrgIdResponse{}
	_body, _err := client.DoROARequest(tea.String("ListSearchTabsByOrgId"), tea.String("search_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/search/tabs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSearchTab(tabId *string, request *UpdateSearchTabRequest) (_result *UpdateSearchTabResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateSearchTabHeaders{}
	_result = &UpdateSearchTabResponse{}
	_body, _err := client.UpdateSearchTabWithOptions(tabId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSearchTabWithOptions(tabId *string, request *UpdateSearchTabRequest, headers *UpdateSearchTabHeaders, runtime *util.RuntimeOptions) (_result *UpdateSearchTabResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	tabId = openapiutil.GetEncodeParam(tabId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
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
	_result = &UpdateSearchTabResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateSearchTab"), tea.String("search_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/search/tabs/"+tea.StringValue(tabId)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
