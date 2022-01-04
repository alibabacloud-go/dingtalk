// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package project_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetDeptsByOrgIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	DingAccessTokenType     *string            `json:"dingAccessTokenType,omitempty" xml:"dingAccessTokenType,omitempty"`
	DingOrgId               *string            `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetDeptsByOrgIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetDeptsByOrgIdHeaders) GoString() string {
	return s.String()
}

func (s *GetDeptsByOrgIdHeaders) SetCommonHeaders(v map[string]*string) *GetDeptsByOrgIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeptsByOrgIdHeaders) SetDingAccessTokenType(v string) *GetDeptsByOrgIdHeaders {
	s.DingAccessTokenType = &v
	return s
}

func (s *GetDeptsByOrgIdHeaders) SetDingOrgId(v string) *GetDeptsByOrgIdHeaders {
	s.DingOrgId = &v
	return s
}

func (s *GetDeptsByOrgIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetDeptsByOrgIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetDeptsByOrgIdRequest struct {
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetDeptsByOrgIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeptsByOrgIdRequest) GoString() string {
	return s.String()
}

func (s *GetDeptsByOrgIdRequest) SetMaxResults(v int32) *GetDeptsByOrgIdRequest {
	s.MaxResults = &v
	return s
}

func (s *GetDeptsByOrgIdRequest) SetNextToken(v int64) *GetDeptsByOrgIdRequest {
	s.NextToken = &v
	return s
}

type GetDeptsByOrgIdResponseBody struct {
	// deptList
	DeptList []*GetDeptsByOrgIdResponseBodyDeptList `json:"deptList,omitempty" xml:"deptList,omitempty" type:"Repeated"`
	// hasMore
	HasMore    *bool  `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// nextCursor
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetDeptsByOrgIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeptsByOrgIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeptsByOrgIdResponseBody) SetDeptList(v []*GetDeptsByOrgIdResponseBodyDeptList) *GetDeptsByOrgIdResponseBody {
	s.DeptList = v
	return s
}

func (s *GetDeptsByOrgIdResponseBody) SetHasMore(v bool) *GetDeptsByOrgIdResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetDeptsByOrgIdResponseBody) SetMaxResults(v int32) *GetDeptsByOrgIdResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetDeptsByOrgIdResponseBody) SetNextToken(v int64) *GetDeptsByOrgIdResponseBody {
	s.NextToken = &v
	return s
}

type GetDeptsByOrgIdResponseBodyDeptList struct {
	// id
	DeptId *int64 `json:"dept_id,omitempty" xml:"dept_id,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// parentId
	ParentId *int64 `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
}

func (s GetDeptsByOrgIdResponseBodyDeptList) String() string {
	return tea.Prettify(s)
}

func (s GetDeptsByOrgIdResponseBodyDeptList) GoString() string {
	return s.String()
}

func (s *GetDeptsByOrgIdResponseBodyDeptList) SetDeptId(v int64) *GetDeptsByOrgIdResponseBodyDeptList {
	s.DeptId = &v
	return s
}

func (s *GetDeptsByOrgIdResponseBodyDeptList) SetName(v string) *GetDeptsByOrgIdResponseBodyDeptList {
	s.Name = &v
	return s
}

func (s *GetDeptsByOrgIdResponseBodyDeptList) SetParentId(v int64) *GetDeptsByOrgIdResponseBodyDeptList {
	s.ParentId = &v
	return s
}

type GetDeptsByOrgIdResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDeptsByOrgIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeptsByOrgIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeptsByOrgIdResponse) GoString() string {
	return s.String()
}

func (s *GetDeptsByOrgIdResponse) SetHeaders(v map[string]*string) *GetDeptsByOrgIdResponse {
	s.Headers = v
	return s
}

func (s *GetDeptsByOrgIdResponse) SetBody(v *GetDeptsByOrgIdResponseBody) *GetDeptsByOrgIdResponse {
	s.Body = v
	return s
}

type GetEmpsByOrgIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	DingAccessTokenType     *string            `json:"dingAccessTokenType,omitempty" xml:"dingAccessTokenType,omitempty"`
	DingOrgId               *string            `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetEmpsByOrgIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetEmpsByOrgIdHeaders) GoString() string {
	return s.String()
}

func (s *GetEmpsByOrgIdHeaders) SetCommonHeaders(v map[string]*string) *GetEmpsByOrgIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetEmpsByOrgIdHeaders) SetDingAccessTokenType(v string) *GetEmpsByOrgIdHeaders {
	s.DingAccessTokenType = &v
	return s
}

func (s *GetEmpsByOrgIdHeaders) SetDingOrgId(v string) *GetEmpsByOrgIdHeaders {
	s.DingOrgId = &v
	return s
}

func (s *GetEmpsByOrgIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetEmpsByOrgIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetEmpsByOrgIdRequest struct {
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NeedDept   *bool  `json:"needDept,omitempty" xml:"needDept,omitempty"`
	NextToken  *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetEmpsByOrgIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEmpsByOrgIdRequest) GoString() string {
	return s.String()
}

func (s *GetEmpsByOrgIdRequest) SetMaxResults(v int32) *GetEmpsByOrgIdRequest {
	s.MaxResults = &v
	return s
}

func (s *GetEmpsByOrgIdRequest) SetNeedDept(v bool) *GetEmpsByOrgIdRequest {
	s.NeedDept = &v
	return s
}

func (s *GetEmpsByOrgIdRequest) SetNextToken(v int64) *GetEmpsByOrgIdRequest {
	s.NextToken = &v
	return s
}

type GetEmpsByOrgIdResponseBody struct {
	// empList
	EmpList []*GetEmpsByOrgIdResponseBodyEmpList `json:"empList,omitempty" xml:"empList,omitempty" type:"Repeated"`
	// hasMore
	HasMore   *bool  `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetEmpsByOrgIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEmpsByOrgIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetEmpsByOrgIdResponseBody) SetEmpList(v []*GetEmpsByOrgIdResponseBodyEmpList) *GetEmpsByOrgIdResponseBody {
	s.EmpList = v
	return s
}

func (s *GetEmpsByOrgIdResponseBody) SetHasMore(v bool) *GetEmpsByOrgIdResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetEmpsByOrgIdResponseBody) SetNextToken(v int64) *GetEmpsByOrgIdResponseBody {
	s.NextToken = &v
	return s
}

type GetEmpsByOrgIdResponseBodyEmpList struct {
	// avatar
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// deptIdList
	DeptIdList []*int64 `json:"dept_id_list,omitempty" xml:"dept_id_list,omitempty" type:"Repeated"`
	// dingId
	DingId *string `json:"dingId,omitempty" xml:"dingId,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// nick
	Nick *string `json:"nick,omitempty" xml:"nick,omitempty"`
	// orgId
	OrgId    *int64  `json:"orgId,omitempty" xml:"orgId,omitempty"`
	Position *string `json:"position,omitempty" xml:"position,omitempty"`
	// unionId
	Unionid *string `json:"unionid,omitempty" xml:"unionid,omitempty"`
	// userid
	Userid *string `json:"userid,omitempty" xml:"userid,omitempty"`
}

func (s GetEmpsByOrgIdResponseBodyEmpList) String() string {
	return tea.Prettify(s)
}

func (s GetEmpsByOrgIdResponseBodyEmpList) GoString() string {
	return s.String()
}

func (s *GetEmpsByOrgIdResponseBodyEmpList) SetAvatar(v string) *GetEmpsByOrgIdResponseBodyEmpList {
	s.Avatar = &v
	return s
}

func (s *GetEmpsByOrgIdResponseBodyEmpList) SetDeptIdList(v []*int64) *GetEmpsByOrgIdResponseBodyEmpList {
	s.DeptIdList = v
	return s
}

func (s *GetEmpsByOrgIdResponseBodyEmpList) SetDingId(v string) *GetEmpsByOrgIdResponseBodyEmpList {
	s.DingId = &v
	return s
}

func (s *GetEmpsByOrgIdResponseBodyEmpList) SetName(v string) *GetEmpsByOrgIdResponseBodyEmpList {
	s.Name = &v
	return s
}

func (s *GetEmpsByOrgIdResponseBodyEmpList) SetNick(v string) *GetEmpsByOrgIdResponseBodyEmpList {
	s.Nick = &v
	return s
}

func (s *GetEmpsByOrgIdResponseBodyEmpList) SetOrgId(v int64) *GetEmpsByOrgIdResponseBodyEmpList {
	s.OrgId = &v
	return s
}

func (s *GetEmpsByOrgIdResponseBodyEmpList) SetPosition(v string) *GetEmpsByOrgIdResponseBodyEmpList {
	s.Position = &v
	return s
}

func (s *GetEmpsByOrgIdResponseBodyEmpList) SetUnionid(v string) *GetEmpsByOrgIdResponseBodyEmpList {
	s.Unionid = &v
	return s
}

func (s *GetEmpsByOrgIdResponseBodyEmpList) SetUserid(v string) *GetEmpsByOrgIdResponseBodyEmpList {
	s.Userid = &v
	return s
}

type GetEmpsByOrgIdResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEmpsByOrgIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEmpsByOrgIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEmpsByOrgIdResponse) GoString() string {
	return s.String()
}

func (s *GetEmpsByOrgIdResponse) SetHeaders(v map[string]*string) *GetEmpsByOrgIdResponse {
	s.Headers = v
	return s
}

func (s *GetEmpsByOrgIdResponse) SetBody(v *GetEmpsByOrgIdResponseBody) *GetEmpsByOrgIdResponse {
	s.Body = v
	return s
}

type GetTbProjectGrayHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	DingAccessTokenType     *string            `json:"dingAccessTokenType,omitempty" xml:"dingAccessTokenType,omitempty"`
	DingCorpId              *string            `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	DingIsvOrgId            *string            `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId               *string            `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingSuiteKey            *string            `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTbProjectGrayHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTbProjectGrayHeaders) GoString() string {
	return s.String()
}

func (s *GetTbProjectGrayHeaders) SetCommonHeaders(v map[string]*string) *GetTbProjectGrayHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTbProjectGrayHeaders) SetDingAccessTokenType(v string) *GetTbProjectGrayHeaders {
	s.DingAccessTokenType = &v
	return s
}

func (s *GetTbProjectGrayHeaders) SetDingCorpId(v string) *GetTbProjectGrayHeaders {
	s.DingCorpId = &v
	return s
}

func (s *GetTbProjectGrayHeaders) SetDingIsvOrgId(v string) *GetTbProjectGrayHeaders {
	s.DingIsvOrgId = &v
	return s
}

func (s *GetTbProjectGrayHeaders) SetDingOrgId(v string) *GetTbProjectGrayHeaders {
	s.DingOrgId = &v
	return s
}

func (s *GetTbProjectGrayHeaders) SetDingSuiteKey(v string) *GetTbProjectGrayHeaders {
	s.DingSuiteKey = &v
	return s
}

func (s *GetTbProjectGrayHeaders) SetXAcsDingtalkAccessToken(v string) *GetTbProjectGrayHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTbProjectGrayRequest struct {
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
}

func (s GetTbProjectGrayRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTbProjectGrayRequest) GoString() string {
	return s.String()
}

func (s *GetTbProjectGrayRequest) SetLabel(v string) *GetTbProjectGrayRequest {
	s.Label = &v
	return s
}

type GetTbProjectGrayResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 是否灰度
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetTbProjectGrayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTbProjectGrayResponseBody) GoString() string {
	return s.String()
}

func (s *GetTbProjectGrayResponseBody) SetRequestId(v string) *GetTbProjectGrayResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTbProjectGrayResponseBody) SetResult(v bool) *GetTbProjectGrayResponseBody {
	s.Result = &v
	return s
}

type GetTbProjectGrayResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTbProjectGrayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTbProjectGrayResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTbProjectGrayResponse) GoString() string {
	return s.String()
}

func (s *GetTbProjectGrayResponse) SetHeaders(v map[string]*string) *GetTbProjectGrayResponse {
	s.Headers = v
	return s
}

func (s *GetTbProjectGrayResponse) SetBody(v *GetTbProjectGrayResponseBody) *GetTbProjectGrayResponse {
	s.Body = v
	return s
}

type GetTbProjectSourceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	DingAccessTokenType     *string            `json:"dingAccessTokenType,omitempty" xml:"dingAccessTokenType,omitempty"`
	DingCorpId              *string            `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	DingIsvOrgId            *string            `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingOrgId               *string            `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingSuiteKey            *string            `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTbProjectSourceHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTbProjectSourceHeaders) GoString() string {
	return s.String()
}

func (s *GetTbProjectSourceHeaders) SetCommonHeaders(v map[string]*string) *GetTbProjectSourceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTbProjectSourceHeaders) SetDingAccessTokenType(v string) *GetTbProjectSourceHeaders {
	s.DingAccessTokenType = &v
	return s
}

func (s *GetTbProjectSourceHeaders) SetDingCorpId(v string) *GetTbProjectSourceHeaders {
	s.DingCorpId = &v
	return s
}

func (s *GetTbProjectSourceHeaders) SetDingIsvOrgId(v string) *GetTbProjectSourceHeaders {
	s.DingIsvOrgId = &v
	return s
}

func (s *GetTbProjectSourceHeaders) SetDingOrgId(v string) *GetTbProjectSourceHeaders {
	s.DingOrgId = &v
	return s
}

func (s *GetTbProjectSourceHeaders) SetDingSuiteKey(v string) *GetTbProjectSourceHeaders {
	s.DingSuiteKey = &v
	return s
}

func (s *GetTbProjectSourceHeaders) SetXAcsDingtalkAccessToken(v string) *GetTbProjectSourceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTbProjectSourceResponseBody struct {
	// 应用安装来源，"0"：来自应用中心，”6“：预安装
	InstallSource *string `json:"installSource,omitempty" xml:"installSource,omitempty"`
}

func (s GetTbProjectSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTbProjectSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetTbProjectSourceResponseBody) SetInstallSource(v string) *GetTbProjectSourceResponseBody {
	s.InstallSource = &v
	return s
}

type GetTbProjectSourceResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTbProjectSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTbProjectSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTbProjectSourceResponse) GoString() string {
	return s.String()
}

func (s *GetTbProjectSourceResponse) SetHeaders(v map[string]*string) *GetTbProjectSourceResponse {
	s.Headers = v
	return s
}

func (s *GetTbProjectSourceResponse) SetBody(v *GetTbProjectSourceResponseBody) *GetTbProjectSourceResponse {
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

func (client *Client) GetDeptsByOrgId(request *GetDeptsByOrgIdRequest) (_result *GetDeptsByOrgIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetDeptsByOrgIdHeaders{}
	_result = &GetDeptsByOrgIdResponse{}
	_body, _err := client.GetDeptsByOrgIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeptsByOrgIdWithOptions(request *GetDeptsByOrgIdRequest, headers *GetDeptsByOrgIdHeaders, runtime *util.RuntimeOptions) (_result *GetDeptsByOrgIdResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(headers.DingAccessTokenType)) {
		realHeaders["dingAccessTokenType"] = util.ToJSONString(headers.DingAccessTokenType)
	}

	if !tea.BoolValue(util.IsUnset(headers.DingOrgId)) {
		realHeaders["dingOrgId"] = util.ToJSONString(headers.DingOrgId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &GetDeptsByOrgIdResponse{}
	_body, _err := client.DoROARequest(tea.String("GetDeptsByOrgId"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/project/orgs/depts"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEmpsByOrgId(request *GetEmpsByOrgIdRequest) (_result *GetEmpsByOrgIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetEmpsByOrgIdHeaders{}
	_result = &GetEmpsByOrgIdResponse{}
	_body, _err := client.GetEmpsByOrgIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEmpsByOrgIdWithOptions(request *GetEmpsByOrgIdRequest, headers *GetEmpsByOrgIdHeaders, runtime *util.RuntimeOptions) (_result *GetEmpsByOrgIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NeedDept)) {
		query["needDept"] = request.NeedDept
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.DingAccessTokenType)) {
		realHeaders["dingAccessTokenType"] = util.ToJSONString(headers.DingAccessTokenType)
	}

	if !tea.BoolValue(util.IsUnset(headers.DingOrgId)) {
		realHeaders["dingOrgId"] = util.ToJSONString(headers.DingOrgId)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &GetEmpsByOrgIdResponse{}
	_body, _err := client.DoROARequest(tea.String("GetEmpsByOrgId"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/project/orgs/employees"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTbProjectGray(request *GetTbProjectGrayRequest) (_result *GetTbProjectGrayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTbProjectGrayHeaders{}
	_result = &GetTbProjectGrayResponse{}
	_body, _err := client.GetTbProjectGrayWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTbProjectGrayWithOptions(request *GetTbProjectGrayRequest, headers *GetTbProjectGrayHeaders, runtime *util.RuntimeOptions) (_result *GetTbProjectGrayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Label)) {
		body["label"] = request.Label
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.DingAccessTokenType)) {
		realHeaders["dingAccessTokenType"] = util.ToJSONString(headers.DingAccessTokenType)
	}

	if !tea.BoolValue(util.IsUnset(headers.DingCorpId)) {
		realHeaders["dingCorpId"] = util.ToJSONString(headers.DingCorpId)
	}

	if !tea.BoolValue(util.IsUnset(headers.DingIsvOrgId)) {
		realHeaders["dingIsvOrgId"] = util.ToJSONString(headers.DingIsvOrgId)
	}

	if !tea.BoolValue(util.IsUnset(headers.DingOrgId)) {
		realHeaders["dingOrgId"] = util.ToJSONString(headers.DingOrgId)
	}

	if !tea.BoolValue(util.IsUnset(headers.DingSuiteKey)) {
		realHeaders["dingSuiteKey"] = util.ToJSONString(headers.DingSuiteKey)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &GetTbProjectGrayResponse{}
	_body, _err := client.DoROARequest(tea.String("GetTbProjectGray"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/project/projects/gray"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTbProjectSource() (_result *GetTbProjectSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTbProjectSourceHeaders{}
	_result = &GetTbProjectSourceResponse{}
	_body, _err := client.GetTbProjectSourceWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTbProjectSourceWithOptions(headers *GetTbProjectSourceHeaders, runtime *util.RuntimeOptions) (_result *GetTbProjectSourceResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.DingAccessTokenType)) {
		realHeaders["dingAccessTokenType"] = util.ToJSONString(headers.DingAccessTokenType)
	}

	if !tea.BoolValue(util.IsUnset(headers.DingCorpId)) {
		realHeaders["dingCorpId"] = util.ToJSONString(headers.DingCorpId)
	}

	if !tea.BoolValue(util.IsUnset(headers.DingIsvOrgId)) {
		realHeaders["dingIsvOrgId"] = util.ToJSONString(headers.DingIsvOrgId)
	}

	if !tea.BoolValue(util.IsUnset(headers.DingOrgId)) {
		realHeaders["dingOrgId"] = util.ToJSONString(headers.DingOrgId)
	}

	if !tea.BoolValue(util.IsUnset(headers.DingSuiteKey)) {
		realHeaders["dingSuiteKey"] = util.ToJSONString(headers.DingSuiteKey)
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &GetTbProjectSourceResponse{}
	_body, _err := client.DoROARequest(tea.String("GetTbProjectSource"), tea.String("project_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/project/projects/source"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
