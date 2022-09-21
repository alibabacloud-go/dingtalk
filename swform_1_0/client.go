// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package swform_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetFormInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFormInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFormInstanceHeaders) GoString() string {
	return s.String()
}

func (s *GetFormInstanceHeaders) SetCommonHeaders(v map[string]*string) *GetFormInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFormInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *GetFormInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFormInstanceRequest struct {
	// 填表类型。0表示通用填表，1表示教育版填表
	BizType *int32 `json:"bizType,omitempty" xml:"bizType,omitempty"`
}

func (s GetFormInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFormInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetFormInstanceRequest) SetBizType(v int32) *GetFormInstanceRequest {
	s.BizType = &v
	return s
}

type GetFormInstanceResponseBody struct {
	// 返回结果对象。
	Result *GetFormInstanceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// 是否成功。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetFormInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFormInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetFormInstanceResponseBody) SetResult(v *GetFormInstanceResponseBodyResult) *GetFormInstanceResponseBody {
	s.Result = v
	return s
}

func (s *GetFormInstanceResponseBody) SetSuccess(v bool) *GetFormInstanceResponseBody {
	s.Success = &v
	return s
}

type GetFormInstanceResponseBodyResult struct {
	// 创建时间。iso8601格式。
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 创建者userid
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 填表code，用此code可调接口获取填表实例列表。
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// 表单内容列表。
	Forms []*GetFormInstanceResponseBodyResultForms `json:"forms,omitempty" xml:"forms,omitempty" type:"Repeated"`
	// 更新时间。iso8601格式。
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// 填表名称。
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetFormInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFormInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFormInstanceResponseBodyResult) SetCreateTime(v string) *GetFormInstanceResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetFormInstanceResponseBodyResult) SetCreator(v string) *GetFormInstanceResponseBodyResult {
	s.Creator = &v
	return s
}

func (s *GetFormInstanceResponseBodyResult) SetFormCode(v string) *GetFormInstanceResponseBodyResult {
	s.FormCode = &v
	return s
}

func (s *GetFormInstanceResponseBodyResult) SetForms(v []*GetFormInstanceResponseBodyResultForms) *GetFormInstanceResponseBodyResult {
	s.Forms = v
	return s
}

func (s *GetFormInstanceResponseBodyResult) SetModifyTime(v string) *GetFormInstanceResponseBodyResult {
	s.ModifyTime = &v
	return s
}

func (s *GetFormInstanceResponseBodyResult) SetTitle(v string) *GetFormInstanceResponseBodyResult {
	s.Title = &v
	return s
}

type GetFormInstanceResponseBodyResultForms struct {
	// 表单控件key。
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// 表单主题。  当label字段为空或不存在时，忽略这个label和value。
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// 表单的值。
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetFormInstanceResponseBodyResultForms) String() string {
	return tea.Prettify(s)
}

func (s GetFormInstanceResponseBodyResultForms) GoString() string {
	return s.String()
}

func (s *GetFormInstanceResponseBodyResultForms) SetKey(v string) *GetFormInstanceResponseBodyResultForms {
	s.Key = &v
	return s
}

func (s *GetFormInstanceResponseBodyResultForms) SetLabel(v string) *GetFormInstanceResponseBodyResultForms {
	s.Label = &v
	return s
}

func (s *GetFormInstanceResponseBodyResultForms) SetValue(v string) *GetFormInstanceResponseBodyResultForms {
	s.Value = &v
	return s
}

type GetFormInstanceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFormInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFormInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetFormInstanceResponse) SetHeaders(v map[string]*string) *GetFormInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetFormInstanceResponse) SetBody(v *GetFormInstanceResponseBody) *GetFormInstanceResponse {
	s.Body = v
	return s
}

type ListFormInstancesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListFormInstancesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListFormInstancesHeaders) GoString() string {
	return s.String()
}

func (s *ListFormInstancesHeaders) SetCommonHeaders(v map[string]*string) *ListFormInstancesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListFormInstancesHeaders) SetXAcsDingtalkAccessToken(v string) *ListFormInstancesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListFormInstancesRequest struct {
	// 时间，必须是YYYY-MM-DD的格式。循环填表才需要传这个参数。
	ActionDate *string `json:"actionDate,omitempty" xml:"actionDate,omitempty"`
	// 填表类型。0表示通用填表，1表示教育版填表
	BizType *int32 `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 分页大小，最大100。
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 分页起始，从0开始。当返回结果中hasMore为false时，表示没有下一页了。否则取返回结果中nextToken的值作为下一次请求的offset。
	NextToken *int32 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListFormInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFormInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListFormInstancesRequest) SetActionDate(v string) *ListFormInstancesRequest {
	s.ActionDate = &v
	return s
}

func (s *ListFormInstancesRequest) SetBizType(v int32) *ListFormInstancesRequest {
	s.BizType = &v
	return s
}

func (s *ListFormInstancesRequest) SetMaxResults(v int32) *ListFormInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListFormInstancesRequest) SetNextToken(v int32) *ListFormInstancesRequest {
	s.NextToken = &v
	return s
}

type ListFormInstancesResponseBody struct {
	// 返回结果对象。
	Result  *ListFormInstancesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListFormInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFormInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFormInstancesResponseBody) SetResult(v *ListFormInstancesResponseBodyResult) *ListFormInstancesResponseBody {
	s.Result = v
	return s
}

func (s *ListFormInstancesResponseBody) SetSuccess(v bool) *ListFormInstancesResponseBody {
	s.Success = &v
	return s
}

type ListFormInstancesResponseBodyResult struct {
	// 是否还有下一页数据。
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 填表实例列表。
	List []*ListFormInstancesResponseBodyResultList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// 下一次分页offset的值。
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListFormInstancesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListFormInstancesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListFormInstancesResponseBodyResult) SetHasMore(v bool) *ListFormInstancesResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *ListFormInstancesResponseBodyResult) SetList(v []*ListFormInstancesResponseBodyResultList) *ListFormInstancesResponseBodyResult {
	s.List = v
	return s
}

func (s *ListFormInstancesResponseBodyResult) SetNextToken(v int64) *ListFormInstancesResponseBodyResult {
	s.NextToken = &v
	return s
}

type ListFormInstancesResponseBodyResultList struct {
	// 创建时间。iso8601格式。
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 填表code，用此code可调接口获取填表列表。
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// 实例ID。
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// 表单内容列表。
	Forms []*ListFormInstancesResponseBodyResultListForms `json:"forms,omitempty" xml:"forms,omitempty" type:"Repeated"`
	// 更新时间。iso8601格式。
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// 学生班级ID。
	StudentClassId *string `json:"studentClassId,omitempty" xml:"studentClassId,omitempty"`
	// 学生班级名称。
	StudentClassName *string `json:"studentClassName,omitempty" xml:"studentClassName,omitempty"`
	// 学生名称。
	StudentName *string `json:"studentName,omitempty" xml:"studentName,omitempty"`
	// 学生ID。
	StudentUserId *string `json:"studentUserId,omitempty" xml:"studentUserId,omitempty"`
	// 提交人的userid。
	SubmitterUserId *string `json:"submitterUserId,omitempty" xml:"submitterUserId,omitempty"`
	// 提交人姓名。
	SubmitterUserName *string `json:"submitterUserName,omitempty" xml:"submitterUserName,omitempty"`
	// 填表名称。
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListFormInstancesResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s ListFormInstancesResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *ListFormInstancesResponseBodyResultList) SetCreateTime(v string) *ListFormInstancesResponseBodyResultList {
	s.CreateTime = &v
	return s
}

func (s *ListFormInstancesResponseBodyResultList) SetFormCode(v string) *ListFormInstancesResponseBodyResultList {
	s.FormCode = &v
	return s
}

func (s *ListFormInstancesResponseBodyResultList) SetFormInstanceId(v string) *ListFormInstancesResponseBodyResultList {
	s.FormInstanceId = &v
	return s
}

func (s *ListFormInstancesResponseBodyResultList) SetForms(v []*ListFormInstancesResponseBodyResultListForms) *ListFormInstancesResponseBodyResultList {
	s.Forms = v
	return s
}

func (s *ListFormInstancesResponseBodyResultList) SetModifyTime(v string) *ListFormInstancesResponseBodyResultList {
	s.ModifyTime = &v
	return s
}

func (s *ListFormInstancesResponseBodyResultList) SetStudentClassId(v string) *ListFormInstancesResponseBodyResultList {
	s.StudentClassId = &v
	return s
}

func (s *ListFormInstancesResponseBodyResultList) SetStudentClassName(v string) *ListFormInstancesResponseBodyResultList {
	s.StudentClassName = &v
	return s
}

func (s *ListFormInstancesResponseBodyResultList) SetStudentName(v string) *ListFormInstancesResponseBodyResultList {
	s.StudentName = &v
	return s
}

func (s *ListFormInstancesResponseBodyResultList) SetStudentUserId(v string) *ListFormInstancesResponseBodyResultList {
	s.StudentUserId = &v
	return s
}

func (s *ListFormInstancesResponseBodyResultList) SetSubmitterUserId(v string) *ListFormInstancesResponseBodyResultList {
	s.SubmitterUserId = &v
	return s
}

func (s *ListFormInstancesResponseBodyResultList) SetSubmitterUserName(v string) *ListFormInstancesResponseBodyResultList {
	s.SubmitterUserName = &v
	return s
}

func (s *ListFormInstancesResponseBodyResultList) SetTitle(v string) *ListFormInstancesResponseBodyResultList {
	s.Title = &v
	return s
}

type ListFormInstancesResponseBodyResultListForms struct {
	// 表单控件key。
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// 表单主题。  当label字段为空或不存在时，忽略这个label和value。
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// 表单的值。
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListFormInstancesResponseBodyResultListForms) String() string {
	return tea.Prettify(s)
}

func (s ListFormInstancesResponseBodyResultListForms) GoString() string {
	return s.String()
}

func (s *ListFormInstancesResponseBodyResultListForms) SetKey(v string) *ListFormInstancesResponseBodyResultListForms {
	s.Key = &v
	return s
}

func (s *ListFormInstancesResponseBodyResultListForms) SetLabel(v string) *ListFormInstancesResponseBodyResultListForms {
	s.Label = &v
	return s
}

func (s *ListFormInstancesResponseBodyResultListForms) SetValue(v string) *ListFormInstancesResponseBodyResultListForms {
	s.Value = &v
	return s
}

type ListFormInstancesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFormInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFormInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFormInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListFormInstancesResponse) SetHeaders(v map[string]*string) *ListFormInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListFormInstancesResponse) SetBody(v *ListFormInstancesResponseBody) *ListFormInstancesResponse {
	s.Body = v
	return s
}

type ListFormSchemasByCreatorHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListFormSchemasByCreatorHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListFormSchemasByCreatorHeaders) GoString() string {
	return s.String()
}

func (s *ListFormSchemasByCreatorHeaders) SetCommonHeaders(v map[string]*string) *ListFormSchemasByCreatorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListFormSchemasByCreatorHeaders) SetXAcsDingtalkAccessToken(v string) *ListFormSchemasByCreatorHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListFormSchemasByCreatorRequest struct {
	// 填表类型。0表示通用填表，1表示教育版填表
	BizType *int32 `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 填表创建人userid
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 分页大小，最大200
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 分页游标，从0开始。后续取返回结果中nextToken的值。
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListFormSchemasByCreatorRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFormSchemasByCreatorRequest) GoString() string {
	return s.String()
}

func (s *ListFormSchemasByCreatorRequest) SetBizType(v int32) *ListFormSchemasByCreatorRequest {
	s.BizType = &v
	return s
}

func (s *ListFormSchemasByCreatorRequest) SetCreator(v string) *ListFormSchemasByCreatorRequest {
	s.Creator = &v
	return s
}

func (s *ListFormSchemasByCreatorRequest) SetMaxResults(v int32) *ListFormSchemasByCreatorRequest {
	s.MaxResults = &v
	return s
}

func (s *ListFormSchemasByCreatorRequest) SetNextToken(v int64) *ListFormSchemasByCreatorRequest {
	s.NextToken = &v
	return s
}

type ListFormSchemasByCreatorResponseBody struct {
	// 返回结果对象。
	Result *ListFormSchemasByCreatorResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// 是否成功。
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListFormSchemasByCreatorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFormSchemasByCreatorResponseBody) GoString() string {
	return s.String()
}

func (s *ListFormSchemasByCreatorResponseBody) SetResult(v *ListFormSchemasByCreatorResponseBodyResult) *ListFormSchemasByCreatorResponseBody {
	s.Result = v
	return s
}

func (s *ListFormSchemasByCreatorResponseBody) SetSuccess(v bool) *ListFormSchemasByCreatorResponseBody {
	s.Success = &v
	return s
}

type ListFormSchemasByCreatorResponseBodyResult struct {
	// 是否还有下一页数据。
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 创建的填表列表。
	List []*ListFormSchemasByCreatorResponseBodyResultList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// 下一次分页offset的值。
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListFormSchemasByCreatorResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListFormSchemasByCreatorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListFormSchemasByCreatorResponseBodyResult) SetHasMore(v bool) *ListFormSchemasByCreatorResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *ListFormSchemasByCreatorResponseBodyResult) SetList(v []*ListFormSchemasByCreatorResponseBodyResultList) *ListFormSchemasByCreatorResponseBodyResult {
	s.List = v
	return s
}

func (s *ListFormSchemasByCreatorResponseBodyResult) SetNextToken(v int64) *ListFormSchemasByCreatorResponseBodyResult {
	s.NextToken = &v
	return s
}

type ListFormSchemasByCreatorResponseBodyResultList struct {
	// 创建人。
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 填表code，用此code可调接口获取填表列表。
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// 填表提示。
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 填表名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 填表设置。
	Setting *ListFormSchemasByCreatorResponseBodyResultListSetting `json:"setting,omitempty" xml:"setting,omitempty" type:"Struct"`
}

func (s ListFormSchemasByCreatorResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s ListFormSchemasByCreatorResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *ListFormSchemasByCreatorResponseBodyResultList) SetCreator(v string) *ListFormSchemasByCreatorResponseBodyResultList {
	s.Creator = &v
	return s
}

func (s *ListFormSchemasByCreatorResponseBodyResultList) SetFormCode(v string) *ListFormSchemasByCreatorResponseBodyResultList {
	s.FormCode = &v
	return s
}

func (s *ListFormSchemasByCreatorResponseBodyResultList) SetMemo(v string) *ListFormSchemasByCreatorResponseBodyResultList {
	s.Memo = &v
	return s
}

func (s *ListFormSchemasByCreatorResponseBodyResultList) SetName(v string) *ListFormSchemasByCreatorResponseBodyResultList {
	s.Name = &v
	return s
}

func (s *ListFormSchemasByCreatorResponseBodyResultList) SetSetting(v *ListFormSchemasByCreatorResponseBodyResultListSetting) *ListFormSchemasByCreatorResponseBodyResultList {
	s.Setting = v
	return s
}

type ListFormSchemasByCreatorResponseBodyResultListSetting struct {
	// 表单类型：  0：一次性填表  1：周期性填表
	BizType *int32 `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 创建时间。iso8601格式。
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 截止时间。iso8601格式。
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 表单类型：  0：一次性填表  1：周期性填表
	FormType *int32 `json:"formType,omitempty" xml:"formType,omitempty"`
	// 填表周期，周一到周日分别用1-7表示。
	LoopDays []*int32 `json:"loopDays,omitempty" xml:"loopDays,omitempty" type:"Repeated"`
	// 循环执行的时间点。
	LoopTime *string `json:"loopTime,omitempty" xml:"loopTime,omitempty"`
	// 填表是否终止的标记。
	Stop *bool `json:"stop,omitempty" xml:"stop,omitempty"`
}

func (s ListFormSchemasByCreatorResponseBodyResultListSetting) String() string {
	return tea.Prettify(s)
}

func (s ListFormSchemasByCreatorResponseBodyResultListSetting) GoString() string {
	return s.String()
}

func (s *ListFormSchemasByCreatorResponseBodyResultListSetting) SetBizType(v int32) *ListFormSchemasByCreatorResponseBodyResultListSetting {
	s.BizType = &v
	return s
}

func (s *ListFormSchemasByCreatorResponseBodyResultListSetting) SetCreateTime(v string) *ListFormSchemasByCreatorResponseBodyResultListSetting {
	s.CreateTime = &v
	return s
}

func (s *ListFormSchemasByCreatorResponseBodyResultListSetting) SetEndTime(v string) *ListFormSchemasByCreatorResponseBodyResultListSetting {
	s.EndTime = &v
	return s
}

func (s *ListFormSchemasByCreatorResponseBodyResultListSetting) SetFormType(v int32) *ListFormSchemasByCreatorResponseBodyResultListSetting {
	s.FormType = &v
	return s
}

func (s *ListFormSchemasByCreatorResponseBodyResultListSetting) SetLoopDays(v []*int32) *ListFormSchemasByCreatorResponseBodyResultListSetting {
	s.LoopDays = v
	return s
}

func (s *ListFormSchemasByCreatorResponseBodyResultListSetting) SetLoopTime(v string) *ListFormSchemasByCreatorResponseBodyResultListSetting {
	s.LoopTime = &v
	return s
}

func (s *ListFormSchemasByCreatorResponseBodyResultListSetting) SetStop(v bool) *ListFormSchemasByCreatorResponseBodyResultListSetting {
	s.Stop = &v
	return s
}

type ListFormSchemasByCreatorResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFormSchemasByCreatorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFormSchemasByCreatorResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFormSchemasByCreatorResponse) GoString() string {
	return s.String()
}

func (s *ListFormSchemasByCreatorResponse) SetHeaders(v map[string]*string) *ListFormSchemasByCreatorResponse {
	s.Headers = v
	return s
}

func (s *ListFormSchemasByCreatorResponse) SetBody(v *ListFormSchemasByCreatorResponseBody) *ListFormSchemasByCreatorResponse {
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

func (client *Client) GetFormInstance(formInstanceId *string, request *GetFormInstanceRequest) (_result *GetFormInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFormInstanceHeaders{}
	_result = &GetFormInstanceResponse{}
	_body, _err := client.GetFormInstanceWithOptions(formInstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFormInstanceWithOptions(formInstanceId *string, request *GetFormInstanceRequest, headers *GetFormInstanceHeaders, runtime *util.RuntimeOptions) (_result *GetFormInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	formInstanceId = openapiutil.GetEncodeParam(formInstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["bizType"] = request.BizType
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
	_result = &GetFormInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("GetFormInstance"), tea.String("swform_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/swform/instances/"+tea.StringValue(formInstanceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFormInstances(formCode *string, request *ListFormInstancesRequest) (_result *ListFormInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListFormInstancesHeaders{}
	_result = &ListFormInstancesResponse{}
	_body, _err := client.ListFormInstancesWithOptions(formCode, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFormInstancesWithOptions(formCode *string, request *ListFormInstancesRequest, headers *ListFormInstancesHeaders, runtime *util.RuntimeOptions) (_result *ListFormInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	formCode = openapiutil.GetEncodeParam(formCode)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionDate)) {
		query["actionDate"] = request.ActionDate
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["bizType"] = request.BizType
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
	_result = &ListFormInstancesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListFormInstances"), tea.String("swform_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/swform/forms/"+tea.StringValue(formCode)+"/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFormSchemasByCreator(request *ListFormSchemasByCreatorRequest) (_result *ListFormSchemasByCreatorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListFormSchemasByCreatorHeaders{}
	_result = &ListFormSchemasByCreatorResponse{}
	_body, _err := client.ListFormSchemasByCreatorWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFormSchemasByCreatorWithOptions(request *ListFormSchemasByCreatorRequest, headers *ListFormSchemasByCreatorHeaders, runtime *util.RuntimeOptions) (_result *ListFormSchemasByCreatorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.Creator)) {
		query["creator"] = request.Creator
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
	_result = &ListFormSchemasByCreatorResponse{}
	_body, _err := client.DoROARequest(tea.String("ListFormSchemasByCreator"), tea.String("swform_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/swform/users/forms"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
