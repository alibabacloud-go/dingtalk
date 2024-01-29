// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package swform_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
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
	Result  *GetFormInstanceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                              `json:"success,omitempty" xml:"success,omitempty"`
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
	CreateTime *string                                   `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Creator    *string                                   `json:"creator,omitempty" xml:"creator,omitempty"`
	FormCode   *string                                   `json:"formCode,omitempty" xml:"formCode,omitempty"`
	Forms      []*GetFormInstanceResponseBodyResultForms `json:"forms,omitempty" xml:"forms,omitempty" type:"Repeated"`
	ModifyTime *string                                   `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	Title      *string                                   `json:"title,omitempty" xml:"title,omitempty"`
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
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetFormInstanceResponse) SetStatusCode(v int32) *GetFormInstanceResponse {
	s.StatusCode = &v
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
	ActionDate *string `json:"actionDate,omitempty" xml:"actionDate,omitempty"`
	BizType    *int32  `json:"bizType,omitempty" xml:"bizType,omitempty"`
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int32  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	HasMore   *bool                                      `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List      []*ListFormInstancesResponseBodyResultList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextToken *int64                                     `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	CreateTime        *string                                         `json:"createTime,omitempty" xml:"createTime,omitempty"`
	FormCode          *string                                         `json:"formCode,omitempty" xml:"formCode,omitempty"`
	FormInstanceId    *string                                         `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	Forms             []*ListFormInstancesResponseBodyResultListForms `json:"forms,omitempty" xml:"forms,omitempty" type:"Repeated"`
	ModifyTime        *string                                         `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	StudentClassId    *string                                         `json:"studentClassId,omitempty" xml:"studentClassId,omitempty"`
	StudentClassName  *string                                         `json:"studentClassName,omitempty" xml:"studentClassName,omitempty"`
	StudentName       *string                                         `json:"studentName,omitempty" xml:"studentName,omitempty"`
	StudentUserId     *string                                         `json:"studentUserId,omitempty" xml:"studentUserId,omitempty"`
	SubmitterUserId   *string                                         `json:"submitterUserId,omitempty" xml:"submitterUserId,omitempty"`
	SubmitterUserName *string                                         `json:"submitterUserName,omitempty" xml:"submitterUserName,omitempty"`
	Title             *string                                         `json:"title,omitempty" xml:"title,omitempty"`
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
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFormInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListFormInstancesResponse) SetStatusCode(v int32) *ListFormInstancesResponse {
	s.StatusCode = &v
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
	BizType    *int32  `json:"bizType,omitempty" xml:"bizType,omitempty"`
	Creator    *string `json:"creator,omitempty" xml:"creator,omitempty"`
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int64  `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	Result  *ListFormSchemasByCreatorResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                       `json:"success,omitempty" xml:"success,omitempty"`
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
	HasMore   *bool                                             `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List      []*ListFormSchemasByCreatorResponseBodyResultList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextToken *int64                                            `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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
	Creator  *string                                                `json:"creator,omitempty" xml:"creator,omitempty"`
	FormCode *string                                                `json:"formCode,omitempty" xml:"formCode,omitempty"`
	Memo     *string                                                `json:"memo,omitempty" xml:"memo,omitempty"`
	Name     *string                                                `json:"name,omitempty" xml:"name,omitempty"`
	Setting  *ListFormSchemasByCreatorResponseBodyResultListSetting `json:"setting,omitempty" xml:"setting,omitempty" type:"Struct"`
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
	BizType    *int32   `json:"bizType,omitempty" xml:"bizType,omitempty"`
	CreateTime *string  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	EndTime    *string  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	FormType   *int32   `json:"formType,omitempty" xml:"formType,omitempty"`
	LoopDays   []*int32 `json:"loopDays,omitempty" xml:"loopDays,omitempty" type:"Repeated"`
	LoopTime   *string  `json:"loopTime,omitempty" xml:"loopTime,omitempty"`
	Stop       *bool    `json:"stop,omitempty" xml:"stop,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFormSchemasByCreatorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListFormSchemasByCreatorResponse) SetStatusCode(v int32) *ListFormSchemasByCreatorResponse {
	s.StatusCode = &v
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

func (client *Client) GetFormInstanceWithOptions(formInstanceId *string, request *GetFormInstanceRequest, headers *GetFormInstanceHeaders, runtime *util.RuntimeOptions) (_result *GetFormInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
	params := &openapi.Params{
		Action:      tea.String("GetFormInstance"),
		Version:     tea.String("swform_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/swform/instances/" + tea.StringValue(formInstanceId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFormInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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

func (client *Client) ListFormInstancesWithOptions(formCode *string, request *ListFormInstancesRequest, headers *ListFormInstancesHeaders, runtime *util.RuntimeOptions) (_result *ListFormInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
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
	params := &openapi.Params{
		Action:      tea.String("ListFormInstances"),
		Version:     tea.String("swform_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/swform/forms/" + tea.StringValue(formCode) + "/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFormInstancesResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("ListFormSchemasByCreator"),
		Version:     tea.String("swform_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/swform/users/forms"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFormSchemasByCreatorResponse{}
	_body, _err := client.Execute(params, req, runtime)
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
