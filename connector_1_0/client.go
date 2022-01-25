// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package connector_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateActionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateActionHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateActionHeaders) GoString() string {
	return s.String()
}

func (s *CreateActionHeaders) SetCommonHeaders(v map[string]*string) *CreateActionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateActionHeaders) SetXAcsDingtalkAccessToken(v string) *CreateActionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateActionRequest struct {
	ActionInfo []*CreateActionRequestActionInfo `json:"actionInfo,omitempty" xml:"actionInfo,omitempty" type:"Repeated"`
}

func (s CreateActionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateActionRequest) GoString() string {
	return s.String()
}

func (s *CreateActionRequest) SetActionInfo(v []*CreateActionRequestActionInfo) *CreateActionRequest {
	s.ActionInfo = v
	return s
}

type CreateActionRequestActionInfo struct {
	// api请求url path，结合Connector上的apiDomain使用
	ApiPath *string `json:"apiPath,omitempty" xml:"apiPath,omitempty"`
	// action维度的apiSecret
	ApiSecret *string `json:"apiSecret,omitempty" xml:"apiSecret,omitempty"`
	// action维度的api请求url
	ApiUrl *string `json:"apiUrl,omitempty" xml:"apiUrl,omitempty"`
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 连接平台连接器id
	DingConnectorId *string `json:"dingConnectorId,omitempty" xml:"dingConnectorId,omitempty"`
	// 入参schema
	InputSchema *string `json:"inputSchema,omitempty" xml:"inputSchema,omitempty"`
	// 服务商的执行事件Id
	IntegratorActionId *string `json:"integratorActionId,omitempty" xml:"integratorActionId,omitempty"`
	// 服务商的连接器Id
	IntegratorConnectorId *string `json:"integratorConnectorId,omitempty" xml:"integratorConnectorId,omitempty"`
	// 名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 出参schema
	OutputSchema *string `json:"outputSchema,omitempty" xml:"outputSchema,omitempty"`
}

func (s CreateActionRequestActionInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateActionRequestActionInfo) GoString() string {
	return s.String()
}

func (s *CreateActionRequestActionInfo) SetApiPath(v string) *CreateActionRequestActionInfo {
	s.ApiPath = &v
	return s
}

func (s *CreateActionRequestActionInfo) SetApiSecret(v string) *CreateActionRequestActionInfo {
	s.ApiSecret = &v
	return s
}

func (s *CreateActionRequestActionInfo) SetApiUrl(v string) *CreateActionRequestActionInfo {
	s.ApiUrl = &v
	return s
}

func (s *CreateActionRequestActionInfo) SetDescription(v string) *CreateActionRequestActionInfo {
	s.Description = &v
	return s
}

func (s *CreateActionRequestActionInfo) SetDingConnectorId(v string) *CreateActionRequestActionInfo {
	s.DingConnectorId = &v
	return s
}

func (s *CreateActionRequestActionInfo) SetInputSchema(v string) *CreateActionRequestActionInfo {
	s.InputSchema = &v
	return s
}

func (s *CreateActionRequestActionInfo) SetIntegratorActionId(v string) *CreateActionRequestActionInfo {
	s.IntegratorActionId = &v
	return s
}

func (s *CreateActionRequestActionInfo) SetIntegratorConnectorId(v string) *CreateActionRequestActionInfo {
	s.IntegratorConnectorId = &v
	return s
}

func (s *CreateActionRequestActionInfo) SetName(v string) *CreateActionRequestActionInfo {
	s.Name = &v
	return s
}

func (s *CreateActionRequestActionInfo) SetOutputSchema(v string) *CreateActionRequestActionInfo {
	s.OutputSchema = &v
	return s
}

type CreateActionResponseBody struct {
	// Id of the request
	Item []*CreateActionResponseBodyItem `json:"item,omitempty" xml:"item,omitempty" type:"Repeated"`
}

func (s CreateActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateActionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateActionResponseBody) SetItem(v []*CreateActionResponseBodyItem) *CreateActionResponseBody {
	s.Item = v
	return s
}

type CreateActionResponseBodyItem struct {
	// 连接平台执行事件id
	DingActionId *string `json:"dingActionId,omitempty" xml:"dingActionId,omitempty"`
	// 连接平台连接器id
	DingConnectorId *string `json:"dingConnectorId,omitempty" xml:"dingConnectorId,omitempty"`
	// 服务商的执行事件id
	IntegratorActionId *string `json:"integratorActionId,omitempty" xml:"integratorActionId,omitempty"`
	// 服务商的连接器Id
	IntegratorConnectorId *string `json:"integratorConnectorId,omitempty" xml:"integratorConnectorId,omitempty"`
	// 错误码
	SubErrCode *string `json:"subErrCode,omitempty" xml:"subErrCode,omitempty"`
	// 错误信息
	SubErrMsg *string `json:"subErrMsg,omitempty" xml:"subErrMsg,omitempty"`
	// 是否执行成功
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateActionResponseBodyItem) String() string {
	return tea.Prettify(s)
}

func (s CreateActionResponseBodyItem) GoString() string {
	return s.String()
}

func (s *CreateActionResponseBodyItem) SetDingActionId(v string) *CreateActionResponseBodyItem {
	s.DingActionId = &v
	return s
}

func (s *CreateActionResponseBodyItem) SetDingConnectorId(v string) *CreateActionResponseBodyItem {
	s.DingConnectorId = &v
	return s
}

func (s *CreateActionResponseBodyItem) SetIntegratorActionId(v string) *CreateActionResponseBodyItem {
	s.IntegratorActionId = &v
	return s
}

func (s *CreateActionResponseBodyItem) SetIntegratorConnectorId(v string) *CreateActionResponseBodyItem {
	s.IntegratorConnectorId = &v
	return s
}

func (s *CreateActionResponseBodyItem) SetSubErrCode(v string) *CreateActionResponseBodyItem {
	s.SubErrCode = &v
	return s
}

func (s *CreateActionResponseBodyItem) SetSubErrMsg(v string) *CreateActionResponseBodyItem {
	s.SubErrMsg = &v
	return s
}

func (s *CreateActionResponseBodyItem) SetSuccess(v string) *CreateActionResponseBodyItem {
	s.Success = &v
	return s
}

type CreateActionResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateActionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateActionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateActionResponse) GoString() string {
	return s.String()
}

func (s *CreateActionResponse) SetHeaders(v map[string]*string) *CreateActionResponse {
	s.Headers = v
	return s
}

func (s *CreateActionResponse) SetBody(v *CreateActionResponseBody) *CreateActionResponse {
	s.Body = v
	return s
}

type CreateConnectorHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateConnectorHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectorHeaders) GoString() string {
	return s.String()
}

func (s *CreateConnectorHeaders) SetCommonHeaders(v map[string]*string) *CreateConnectorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateConnectorHeaders) SetXAcsDingtalkAccessToken(v string) *CreateConnectorHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateConnectorRequest struct {
	ConnectorInfo []*CreateConnectorRequestConnectorInfo `json:"connectorInfo,omitempty" xml:"connectorInfo,omitempty" type:"Repeated"`
}

func (s CreateConnectorRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectorRequest) GoString() string {
	return s.String()
}

func (s *CreateConnectorRequest) SetConnectorInfo(v []*CreateConnectorRequestConnectorInfo) *CreateConnectorRequest {
	s.ConnectorInfo = v
	return s
}

type CreateConnectorRequestConnectorInfo struct {
	// 连接器下action api请求url域名，当baseVariables中无apiDomain，该项必填
	ApiDomain *string `json:"apiDomain,omitempty" xml:"apiDomain,omitempty"`
	// 连接器下action api请求时使用http加密签名，当baseVariables无apiSecret，该项必填
	ApiSecret *string `json:"apiSecret,omitempty" xml:"apiSecret,omitempty"`
	AppId     *int64  `json:"appId,omitempty" xml:"appId,omitempty"`
	// 变量列表。目前支持将apiDomain、apiSecret声明为基础变量。若声明为变量，则对应属性可不传值
	BaseVariables         []*string `json:"baseVariables,omitempty" xml:"baseVariables,omitempty" type:"Repeated"`
	Description           *string   `json:"description,omitempty" xml:"description,omitempty"`
	IconMediaId           *string   `json:"iconMediaId,omitempty" xml:"iconMediaId,omitempty"`
	IntegratorConnectorId *string   `json:"integratorConnectorId,omitempty" xml:"integratorConnectorId,omitempty"`
	Name                  *string   `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateConnectorRequestConnectorInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectorRequestConnectorInfo) GoString() string {
	return s.String()
}

func (s *CreateConnectorRequestConnectorInfo) SetApiDomain(v string) *CreateConnectorRequestConnectorInfo {
	s.ApiDomain = &v
	return s
}

func (s *CreateConnectorRequestConnectorInfo) SetApiSecret(v string) *CreateConnectorRequestConnectorInfo {
	s.ApiSecret = &v
	return s
}

func (s *CreateConnectorRequestConnectorInfo) SetAppId(v int64) *CreateConnectorRequestConnectorInfo {
	s.AppId = &v
	return s
}

func (s *CreateConnectorRequestConnectorInfo) SetBaseVariables(v []*string) *CreateConnectorRequestConnectorInfo {
	s.BaseVariables = v
	return s
}

func (s *CreateConnectorRequestConnectorInfo) SetDescription(v string) *CreateConnectorRequestConnectorInfo {
	s.Description = &v
	return s
}

func (s *CreateConnectorRequestConnectorInfo) SetIconMediaId(v string) *CreateConnectorRequestConnectorInfo {
	s.IconMediaId = &v
	return s
}

func (s *CreateConnectorRequestConnectorInfo) SetIntegratorConnectorId(v string) *CreateConnectorRequestConnectorInfo {
	s.IntegratorConnectorId = &v
	return s
}

func (s *CreateConnectorRequestConnectorInfo) SetName(v string) *CreateConnectorRequestConnectorInfo {
	s.Name = &v
	return s
}

type CreateConnectorResponseBody struct {
	// responseUnitList
	Item []*CreateConnectorResponseBodyItem `json:"item,omitempty" xml:"item,omitempty" type:"Repeated"`
}

func (s CreateConnectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConnectorResponseBody) SetItem(v []*CreateConnectorResponseBodyItem) *CreateConnectorResponseBody {
	s.Item = v
	return s
}

type CreateConnectorResponseBodyItem struct {
	// 连接平台connectorId
	DingConnectorId *string `json:"dingConnectorId,omitempty" xml:"dingConnectorId,omitempty"`
	// 服务商连接器connectorId
	IntegratorConnectorId *string `json:"integratorConnectorId,omitempty" xml:"integratorConnectorId,omitempty"`
	// 错误码
	SubErrCode *string `json:"subErrCode,omitempty" xml:"subErrCode,omitempty"`
	// 错误信息
	SubErrMsg *string `json:"subErrMsg,omitempty" xml:"subErrMsg,omitempty"`
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateConnectorResponseBodyItem) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectorResponseBodyItem) GoString() string {
	return s.String()
}

func (s *CreateConnectorResponseBodyItem) SetDingConnectorId(v string) *CreateConnectorResponseBodyItem {
	s.DingConnectorId = &v
	return s
}

func (s *CreateConnectorResponseBodyItem) SetIntegratorConnectorId(v string) *CreateConnectorResponseBodyItem {
	s.IntegratorConnectorId = &v
	return s
}

func (s *CreateConnectorResponseBodyItem) SetSubErrCode(v string) *CreateConnectorResponseBodyItem {
	s.SubErrCode = &v
	return s
}

func (s *CreateConnectorResponseBodyItem) SetSubErrMsg(v string) *CreateConnectorResponseBodyItem {
	s.SubErrMsg = &v
	return s
}

func (s *CreateConnectorResponseBodyItem) SetSuccess(v bool) *CreateConnectorResponseBodyItem {
	s.Success = &v
	return s
}

type CreateConnectorResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateConnectorResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConnectorResponse) GoString() string {
	return s.String()
}

func (s *CreateConnectorResponse) SetHeaders(v map[string]*string) *CreateConnectorResponse {
	s.Headers = v
	return s
}

func (s *CreateConnectorResponse) SetBody(v *CreateConnectorResponseBody) *CreateConnectorResponse {
	s.Body = v
	return s
}

type CreateTriggerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateTriggerHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerHeaders) GoString() string {
	return s.String()
}

func (s *CreateTriggerHeaders) SetCommonHeaders(v map[string]*string) *CreateTriggerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTriggerHeaders) SetXAcsDingtalkAccessToken(v string) *CreateTriggerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateTriggerRequest struct {
	TriggerInfo []*CreateTriggerRequestTriggerInfo `json:"triggerInfo,omitempty" xml:"triggerInfo,omitempty" type:"Repeated"`
}

func (s CreateTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerRequest) GoString() string {
	return s.String()
}

func (s *CreateTriggerRequest) SetTriggerInfo(v []*CreateTriggerRequestTriggerInfo) *CreateTriggerRequest {
	s.TriggerInfo = v
	return s
}

type CreateTriggerRequestTriggerInfo struct {
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 连接平台连接器id
	DingConnectorId *string `json:"dingConnectorId,omitempty" xml:"dingConnectorId,omitempty"`
	// 入参jsonSchema
	InputSchema *string `json:"inputSchema,omitempty" xml:"inputSchema,omitempty"`
	// 服务商的连接器Id，优先使用dingConnectorId，其次使用integratorConnectorId
	IntegratorConnectorId *string `json:"integratorConnectorId,omitempty" xml:"integratorConnectorId,omitempty"`
	// 服务商的触发事件Id
	IntegratorTriggerId *string `json:"integratorTriggerId,omitempty" xml:"integratorTriggerId,omitempty"`
	// 名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateTriggerRequestTriggerInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerRequestTriggerInfo) GoString() string {
	return s.String()
}

func (s *CreateTriggerRequestTriggerInfo) SetDescription(v string) *CreateTriggerRequestTriggerInfo {
	s.Description = &v
	return s
}

func (s *CreateTriggerRequestTriggerInfo) SetDingConnectorId(v string) *CreateTriggerRequestTriggerInfo {
	s.DingConnectorId = &v
	return s
}

func (s *CreateTriggerRequestTriggerInfo) SetInputSchema(v string) *CreateTriggerRequestTriggerInfo {
	s.InputSchema = &v
	return s
}

func (s *CreateTriggerRequestTriggerInfo) SetIntegratorConnectorId(v string) *CreateTriggerRequestTriggerInfo {
	s.IntegratorConnectorId = &v
	return s
}

func (s *CreateTriggerRequestTriggerInfo) SetIntegratorTriggerId(v string) *CreateTriggerRequestTriggerInfo {
	s.IntegratorTriggerId = &v
	return s
}

func (s *CreateTriggerRequestTriggerInfo) SetName(v string) *CreateTriggerRequestTriggerInfo {
	s.Name = &v
	return s
}

type CreateTriggerResponseBody struct {
	// Id of the request
	Item []*CreateTriggerResponseBodyItem `json:"item,omitempty" xml:"item,omitempty" type:"Repeated"`
}

func (s CreateTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTriggerResponseBody) SetItem(v []*CreateTriggerResponseBodyItem) *CreateTriggerResponseBody {
	s.Item = v
	return s
}

type CreateTriggerResponseBodyItem struct {
	// 连接平台连接器id
	DingConnectorId *string `json:"dingConnectorId,omitempty" xml:"dingConnectorId,omitempty"`
	// 连接平台触发事件id
	DingTriggerId *string `json:"dingTriggerId,omitempty" xml:"dingTriggerId,omitempty"`
	// 服务商的连接器Id
	IntegratorConnectorId *string `json:"integratorConnectorId,omitempty" xml:"integratorConnectorId,omitempty"`
	// 服务商的触发事件id
	IntegratorTriggerId *string `json:"integratorTriggerId,omitempty" xml:"integratorTriggerId,omitempty"`
	// 错误码
	SubErrCode *string `json:"subErrCode,omitempty" xml:"subErrCode,omitempty"`
	// 错误信息
	SubErrMsg *string `json:"subErrMsg,omitempty" xml:"subErrMsg,omitempty"`
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateTriggerResponseBodyItem) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerResponseBodyItem) GoString() string {
	return s.String()
}

func (s *CreateTriggerResponseBodyItem) SetDingConnectorId(v string) *CreateTriggerResponseBodyItem {
	s.DingConnectorId = &v
	return s
}

func (s *CreateTriggerResponseBodyItem) SetDingTriggerId(v string) *CreateTriggerResponseBodyItem {
	s.DingTriggerId = &v
	return s
}

func (s *CreateTriggerResponseBodyItem) SetIntegratorConnectorId(v string) *CreateTriggerResponseBodyItem {
	s.IntegratorConnectorId = &v
	return s
}

func (s *CreateTriggerResponseBodyItem) SetIntegratorTriggerId(v string) *CreateTriggerResponseBodyItem {
	s.IntegratorTriggerId = &v
	return s
}

func (s *CreateTriggerResponseBodyItem) SetSubErrCode(v string) *CreateTriggerResponseBodyItem {
	s.SubErrCode = &v
	return s
}

func (s *CreateTriggerResponseBodyItem) SetSubErrMsg(v string) *CreateTriggerResponseBodyItem {
	s.SubErrMsg = &v
	return s
}

func (s *CreateTriggerResponseBodyItem) SetSuccess(v bool) *CreateTriggerResponseBodyItem {
	s.Success = &v
	return s
}

type CreateTriggerResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTriggerResponse) GoString() string {
	return s.String()
}

func (s *CreateTriggerResponse) SetHeaders(v map[string]*string) *CreateTriggerResponse {
	s.Headers = v
	return s
}

func (s *CreateTriggerResponse) SetBody(v *CreateTriggerResponseBody) *CreateTriggerResponse {
	s.Body = v
	return s
}

type PullDataByPageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PullDataByPageHeaders) String() string {
	return tea.Prettify(s)
}

func (s PullDataByPageHeaders) GoString() string {
	return s.String()
}

func (s *PullDataByPageHeaders) SetCommonHeaders(v map[string]*string) *PullDataByPageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PullDataByPageHeaders) SetXAcsDingtalkAccessToken(v string) *PullDataByPageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PullDataByPageRequest struct {
	// 同步数据的应用id，isv应用传isv应用id，企业自建应用传agentId。
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// 要拉取的主数据模型id。
	DataModelId *string `json:"dataModelId,omitempty" xml:"dataModelId,omitempty"`
	// 用于过滤时间范围的字段，包含数据创建时间(dataGmtCreate)和数据修改时间(dataGmtModified)，如不传则不过滤。
	DatetimeFilterField *string `json:"datetimeFilterField,omitempty" xml:"datetimeFilterField,omitempty"`
	// 当配置了datetimeFilterField字段后，数据的时间终点，如果不传则按最新一条数据作为终点。
	MaxDatetime *int64 `json:"maxDatetime,omitempty" xml:"maxDatetime,omitempty"`
	// 单次获取的最大记录条数，最大限制100条。
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 当配置了datetimeFilterField字段后，数据的时间起点，如果不传则将最早一条数据作为起点。
	MinDatetime *int64 `json:"minDatetime,omitempty" xml:"minDatetime,omitempty"`
	// 用于翻页的游标，如果为空则从第一条数据开始查询。
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s PullDataByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s PullDataByPageRequest) GoString() string {
	return s.String()
}

func (s *PullDataByPageRequest) SetAppId(v string) *PullDataByPageRequest {
	s.AppId = &v
	return s
}

func (s *PullDataByPageRequest) SetDataModelId(v string) *PullDataByPageRequest {
	s.DataModelId = &v
	return s
}

func (s *PullDataByPageRequest) SetDatetimeFilterField(v string) *PullDataByPageRequest {
	s.DatetimeFilterField = &v
	return s
}

func (s *PullDataByPageRequest) SetMaxDatetime(v int64) *PullDataByPageRequest {
	s.MaxDatetime = &v
	return s
}

func (s *PullDataByPageRequest) SetMaxResults(v int64) *PullDataByPageRequest {
	s.MaxResults = &v
	return s
}

func (s *PullDataByPageRequest) SetMinDatetime(v int64) *PullDataByPageRequest {
	s.MinDatetime = &v
	return s
}

func (s *PullDataByPageRequest) SetNextToken(v string) *PullDataByPageRequest {
	s.NextToken = &v
	return s
}

type PullDataByPageResponseBody struct {
	// resultList
	List []*PullDataByPageResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// 单次获取的最大记录条数。
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 用于查看下一页数据的游标，如果为空则说明没有更多数据了。
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s PullDataByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PullDataByPageResponseBody) GoString() string {
	return s.String()
}

func (s *PullDataByPageResponseBody) SetList(v []*PullDataByPageResponseBodyList) *PullDataByPageResponseBody {
	s.List = v
	return s
}

func (s *PullDataByPageResponseBody) SetMaxResults(v int64) *PullDataByPageResponseBody {
	s.MaxResults = &v
	return s
}

func (s *PullDataByPageResponseBody) SetNextToken(v string) *PullDataByPageResponseBody {
	s.NextToken = &v
	return s
}

type PullDataByPageResponseBodyList struct {
	// 创建数据的应用id。
	DataCreateAppId *string `json:"dataCreateAppId,omitempty" xml:"dataCreateAppId,omitempty"`
	// 创建数据的应用类型，isv应用为premium_microapp。
	DataCreateAppType *string `json:"dataCreateAppType,omitempty" xml:"dataCreateAppType,omitempty"`
	// 数据创建时间。
	DataGmtCreate *int64 `json:"dataGmtCreate,omitempty" xml:"dataGmtCreate,omitempty"`
	// 数据最后修改时间。
	DataGmtModified *int64 `json:"dataGmtModified,omitempty" xml:"dataGmtModified,omitempty"`
	// 最后修改数据的应用id。
	DataModifiedAppId *string `json:"dataModifiedAppId,omitempty" xml:"dataModifiedAppId,omitempty"`
	// 最后修改数据的应用类型，取值同dataCreateAppType。
	DataModifiedAppType *string `json:"dataModifiedAppType,omitempty" xml:"dataModifiedAppType,omitempty"`
	// 数据完整内容。
	JsonData *string `json:"jsonData,omitempty" xml:"jsonData,omitempty"`
}

func (s PullDataByPageResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s PullDataByPageResponseBodyList) GoString() string {
	return s.String()
}

func (s *PullDataByPageResponseBodyList) SetDataCreateAppId(v string) *PullDataByPageResponseBodyList {
	s.DataCreateAppId = &v
	return s
}

func (s *PullDataByPageResponseBodyList) SetDataCreateAppType(v string) *PullDataByPageResponseBodyList {
	s.DataCreateAppType = &v
	return s
}

func (s *PullDataByPageResponseBodyList) SetDataGmtCreate(v int64) *PullDataByPageResponseBodyList {
	s.DataGmtCreate = &v
	return s
}

func (s *PullDataByPageResponseBodyList) SetDataGmtModified(v int64) *PullDataByPageResponseBodyList {
	s.DataGmtModified = &v
	return s
}

func (s *PullDataByPageResponseBodyList) SetDataModifiedAppId(v string) *PullDataByPageResponseBodyList {
	s.DataModifiedAppId = &v
	return s
}

func (s *PullDataByPageResponseBodyList) SetDataModifiedAppType(v string) *PullDataByPageResponseBodyList {
	s.DataModifiedAppType = &v
	return s
}

func (s *PullDataByPageResponseBodyList) SetJsonData(v string) *PullDataByPageResponseBodyList {
	s.JsonData = &v
	return s
}

type PullDataByPageResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PullDataByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PullDataByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s PullDataByPageResponse) GoString() string {
	return s.String()
}

func (s *PullDataByPageResponse) SetHeaders(v map[string]*string) *PullDataByPageResponse {
	s.Headers = v
	return s
}

func (s *PullDataByPageResponse) SetBody(v *PullDataByPageResponseBody) *PullDataByPageResponse {
	s.Body = v
	return s
}

type PullDataByPkHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PullDataByPkHeaders) String() string {
	return tea.Prettify(s)
}

func (s PullDataByPkHeaders) GoString() string {
	return s.String()
}

func (s *PullDataByPkHeaders) SetCommonHeaders(v map[string]*string) *PullDataByPkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PullDataByPkHeaders) SetXAcsDingtalkAccessToken(v string) *PullDataByPkHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PullDataByPkRequest struct {
	// 同步数据的应用id，isv应用传isv应用id，企业自建应用传agentId。
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// 数据的主键字段值。
	PrimaryKey *string `json:"primaryKey,omitempty" xml:"primaryKey,omitempty"`
}

func (s PullDataByPkRequest) String() string {
	return tea.Prettify(s)
}

func (s PullDataByPkRequest) GoString() string {
	return s.String()
}

func (s *PullDataByPkRequest) SetAppId(v string) *PullDataByPkRequest {
	s.AppId = &v
	return s
}

func (s *PullDataByPkRequest) SetPrimaryKey(v string) *PullDataByPkRequest {
	s.PrimaryKey = &v
	return s
}

type PullDataByPkResponseBody struct {
	// 创建数据的应用id。
	DataCreateAppId *string `json:"dataCreateAppId,omitempty" xml:"dataCreateAppId,omitempty"`
	// 创建数据的应用类型，isv应用为premium_microapp。
	DataCreateAppType *string `json:"dataCreateAppType,omitempty" xml:"dataCreateAppType,omitempty"`
	// 数据创建时间。
	DataGmtCreate *int64 `json:"dataGmtCreate,omitempty" xml:"dataGmtCreate,omitempty"`
	// 数据最后修改时间。
	DataGmtModified *int64 `json:"dataGmtModified,omitempty" xml:"dataGmtModified,omitempty"`
	// 最后修改数据的应用id。
	DataModifiedAppId *string `json:"dataModifiedAppId,omitempty" xml:"dataModifiedAppId,omitempty"`
	// 最后修改数据的应用类型，取值同dataCreateAppType。
	DataModifiedAppType *string `json:"dataModifiedAppType,omitempty" xml:"dataModifiedAppType,omitempty"`
	// 数据完整内容。
	JsonData *string `json:"jsonData,omitempty" xml:"jsonData,omitempty"`
}

func (s PullDataByPkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PullDataByPkResponseBody) GoString() string {
	return s.String()
}

func (s *PullDataByPkResponseBody) SetDataCreateAppId(v string) *PullDataByPkResponseBody {
	s.DataCreateAppId = &v
	return s
}

func (s *PullDataByPkResponseBody) SetDataCreateAppType(v string) *PullDataByPkResponseBody {
	s.DataCreateAppType = &v
	return s
}

func (s *PullDataByPkResponseBody) SetDataGmtCreate(v int64) *PullDataByPkResponseBody {
	s.DataGmtCreate = &v
	return s
}

func (s *PullDataByPkResponseBody) SetDataGmtModified(v int64) *PullDataByPkResponseBody {
	s.DataGmtModified = &v
	return s
}

func (s *PullDataByPkResponseBody) SetDataModifiedAppId(v string) *PullDataByPkResponseBody {
	s.DataModifiedAppId = &v
	return s
}

func (s *PullDataByPkResponseBody) SetDataModifiedAppType(v string) *PullDataByPkResponseBody {
	s.DataModifiedAppType = &v
	return s
}

func (s *PullDataByPkResponseBody) SetJsonData(v string) *PullDataByPkResponseBody {
	s.JsonData = &v
	return s
}

type PullDataByPkResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PullDataByPkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PullDataByPkResponse) String() string {
	return tea.Prettify(s)
}

func (s PullDataByPkResponse) GoString() string {
	return s.String()
}

func (s *PullDataByPkResponse) SetHeaders(v map[string]*string) *PullDataByPkResponse {
	s.Headers = v
	return s
}

func (s *PullDataByPkResponse) SetBody(v *PullDataByPkResponseBody) *PullDataByPkResponse {
	s.Body = v
	return s
}

type SyncDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SyncDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s SyncDataHeaders) GoString() string {
	return s.String()
}

func (s *SyncDataHeaders) SetCommonHeaders(v map[string]*string) *SyncDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncDataHeaders) SetXAcsDingtalkAccessToken(v string) *SyncDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SyncDataRequest struct {
	// 同步数据的应用id，isv应用传isv应用id，企业自建应用传agentId。
	AppId           *string                           `json:"appId,omitempty" xml:"appId,omitempty"`
	TriggerDataList []*SyncDataRequestTriggerDataList `json:"triggerDataList,omitempty" xml:"triggerDataList,omitempty" type:"Repeated"`
}

func (s SyncDataRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncDataRequest) GoString() string {
	return s.String()
}

func (s *SyncDataRequest) SetAppId(v string) *SyncDataRequest {
	s.AppId = &v
	return s
}

func (s *SyncDataRequest) SetTriggerDataList(v []*SyncDataRequestTriggerDataList) *SyncDataRequest {
	s.TriggerDataList = v
	return s
}

type SyncDataRequestTriggerDataList struct {
	Action          *string `json:"action,omitempty" xml:"action,omitempty"`
	CustomTriggerId *string `json:"customTriggerId,omitempty" xml:"customTriggerId,omitempty"`
	DataGmtCreate   *int64  `json:"dataGmtCreate,omitempty" xml:"dataGmtCreate,omitempty"`
	DataGmtModified *int64  `json:"dataGmtModified,omitempty" xml:"dataGmtModified,omitempty"`
	JsonData        *string `json:"jsonData,omitempty" xml:"jsonData,omitempty"`
	TriggerId       *string `json:"triggerId,omitempty" xml:"triggerId,omitempty"`
}

func (s SyncDataRequestTriggerDataList) String() string {
	return tea.Prettify(s)
}

func (s SyncDataRequestTriggerDataList) GoString() string {
	return s.String()
}

func (s *SyncDataRequestTriggerDataList) SetAction(v string) *SyncDataRequestTriggerDataList {
	s.Action = &v
	return s
}

func (s *SyncDataRequestTriggerDataList) SetCustomTriggerId(v string) *SyncDataRequestTriggerDataList {
	s.CustomTriggerId = &v
	return s
}

func (s *SyncDataRequestTriggerDataList) SetDataGmtCreate(v int64) *SyncDataRequestTriggerDataList {
	s.DataGmtCreate = &v
	return s
}

func (s *SyncDataRequestTriggerDataList) SetDataGmtModified(v int64) *SyncDataRequestTriggerDataList {
	s.DataGmtModified = &v
	return s
}

func (s *SyncDataRequestTriggerDataList) SetJsonData(v string) *SyncDataRequestTriggerDataList {
	s.JsonData = &v
	return s
}

func (s *SyncDataRequestTriggerDataList) SetTriggerId(v string) *SyncDataRequestTriggerDataList {
	s.TriggerId = &v
	return s
}

type SyncDataResponseBody struct {
	// resultList
	List []*SyncDataResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s SyncDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncDataResponseBody) GoString() string {
	return s.String()
}

func (s *SyncDataResponseBody) SetList(v []*SyncDataResponseBodyList) *SyncDataResponseBody {
	s.List = v
	return s
}

type SyncDataResponseBodyList struct {
	BizPrimaryKey *string `json:"bizPrimaryKey,omitempty" xml:"bizPrimaryKey,omitempty"`
	SubErrCode    *string `json:"subErrCode,omitempty" xml:"subErrCode,omitempty"`
	SubErrMsg     *string `json:"subErrMsg,omitempty" xml:"subErrMsg,omitempty"`
	Success       *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TriggerId     *string `json:"triggerId,omitempty" xml:"triggerId,omitempty"`
}

func (s SyncDataResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s SyncDataResponseBodyList) GoString() string {
	return s.String()
}

func (s *SyncDataResponseBodyList) SetBizPrimaryKey(v string) *SyncDataResponseBodyList {
	s.BizPrimaryKey = &v
	return s
}

func (s *SyncDataResponseBodyList) SetSubErrCode(v string) *SyncDataResponseBodyList {
	s.SubErrCode = &v
	return s
}

func (s *SyncDataResponseBodyList) SetSubErrMsg(v string) *SyncDataResponseBodyList {
	s.SubErrMsg = &v
	return s
}

func (s *SyncDataResponseBodyList) SetSuccess(v bool) *SyncDataResponseBodyList {
	s.Success = &v
	return s
}

func (s *SyncDataResponseBodyList) SetTriggerId(v string) *SyncDataResponseBodyList {
	s.TriggerId = &v
	return s
}

type SyncDataResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SyncDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncDataResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncDataResponse) GoString() string {
	return s.String()
}

func (s *SyncDataResponse) SetHeaders(v map[string]*string) *SyncDataResponse {
	s.Headers = v
	return s
}

func (s *SyncDataResponse) SetBody(v *SyncDataResponseBody) *SyncDataResponse {
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

func (client *Client) CreateAction(request *CreateActionRequest) (_result *CreateActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateActionHeaders{}
	_result = &CreateActionResponse{}
	_body, _err := client.CreateActionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateActionWithOptions(request *CreateActionRequest, headers *CreateActionHeaders, runtime *util.RuntimeOptions) (_result *CreateActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionInfo)) {
		body["actionInfo"] = request.ActionInfo
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
	_result = &CreateActionResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateAction"), tea.String("connector_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/connector/actions"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateConnector(request *CreateConnectorRequest) (_result *CreateConnectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateConnectorHeaders{}
	_result = &CreateConnectorResponse{}
	_body, _err := client.CreateConnectorWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateConnectorWithOptions(request *CreateConnectorRequest, headers *CreateConnectorHeaders, runtime *util.RuntimeOptions) (_result *CreateConnectorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectorInfo)) {
		body["connectorInfo"] = request.ConnectorInfo
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
	_result = &CreateConnectorResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateConnector"), tea.String("connector_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/connector/connectors"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTrigger(request *CreateTriggerRequest) (_result *CreateTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateTriggerHeaders{}
	_result = &CreateTriggerResponse{}
	_body, _err := client.CreateTriggerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTriggerWithOptions(request *CreateTriggerRequest, headers *CreateTriggerHeaders, runtime *util.RuntimeOptions) (_result *CreateTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TriggerInfo)) {
		body["triggerInfo"] = request.TriggerInfo
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
	_result = &CreateTriggerResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateTrigger"), tea.String("connector_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/connector/triggers"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PullDataByPage(request *PullDataByPageRequest) (_result *PullDataByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PullDataByPageHeaders{}
	_result = &PullDataByPageResponse{}
	_body, _err := client.PullDataByPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PullDataByPageWithOptions(request *PullDataByPageRequest, headers *PullDataByPageHeaders, runtime *util.RuntimeOptions) (_result *PullDataByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.DataModelId)) {
		query["dataModelId"] = request.DataModelId
	}

	if !tea.BoolValue(util.IsUnset(request.DatetimeFilterField)) {
		query["datetimeFilterField"] = request.DatetimeFilterField
	}

	if !tea.BoolValue(util.IsUnset(request.MaxDatetime)) {
		query["maxDatetime"] = request.MaxDatetime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.MinDatetime)) {
		query["minDatetime"] = request.MinDatetime
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
	_result = &PullDataByPageResponse{}
	_body, _err := client.DoROARequest(tea.String("PullDataByPage"), tea.String("connector_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/connector/data"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PullDataByPk(dataModelId *string, request *PullDataByPkRequest) (_result *PullDataByPkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PullDataByPkHeaders{}
	_result = &PullDataByPkResponse{}
	_body, _err := client.PullDataByPkWithOptions(dataModelId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PullDataByPkWithOptions(dataModelId *string, request *PullDataByPkRequest, headers *PullDataByPkHeaders, runtime *util.RuntimeOptions) (_result *PullDataByPkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	dataModelId = openapiutil.GetEncodeParam(dataModelId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.PrimaryKey)) {
		query["primaryKey"] = request.PrimaryKey
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
	_result = &PullDataByPkResponse{}
	_body, _err := client.DoROARequest(tea.String("PullDataByPk"), tea.String("connector_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/connector/data/"+tea.StringValue(dataModelId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncData(request *SyncDataRequest) (_result *SyncDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SyncDataHeaders{}
	_result = &SyncDataResponse{}
	_body, _err := client.SyncDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncDataWithOptions(request *SyncDataRequest, headers *SyncDataHeaders, runtime *util.RuntimeOptions) (_result *SyncDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerDataList)) {
		body["triggerDataList"] = request.TriggerDataList
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
	_result = &SyncDataResponse{}
	_body, _err := client.DoROARequest(tea.String("SyncData"), tea.String("connector_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/connector/triggers/data/sync"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
