// This file is auto-generated, don't edit it. Thanks.
package blackboard_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetBlackboardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetBlackboardHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetBlackboardHeaders) GoString() string {
	return s.String()
}

func (s *GetBlackboardHeaders) SetCommonHeaders(v map[string]*string) *GetBlackboardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetBlackboardHeaders) SetXAcsDingtalkAccessToken(v string) *GetBlackboardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetBlackboardRequest struct {
	// example:
	//
	// ca80xxxx0a04
	BlackboardId *string `json:"blackboardId,omitempty" xml:"blackboardId,omitempty"`
	// example:
	//
	// manager01
	OperationUserId *string `json:"operationUserId,omitempty" xml:"operationUserId,omitempty"`
}

func (s GetBlackboardRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBlackboardRequest) GoString() string {
	return s.String()
}

func (s *GetBlackboardRequest) SetBlackboardId(v string) *GetBlackboardRequest {
	s.BlackboardId = &v
	return s
}

func (s *GetBlackboardRequest) SetOperationUserId(v string) *GetBlackboardRequest {
	s.OperationUserId = &v
	return s
}

type GetBlackboardResponseBody struct {
	Attachments []*GetBlackboardResponseBodyAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// example:
	//
	// example_category_id
	CategoryId *string `json:"categoryId,omitempty" xml:"categoryId,omitempty"`
	// example:
	//
	// 分类示例
	CategoryName *string `json:"categoryName,omitempty" xml:"categoryName,omitempty"`
	// example:
	//
	// 公告内容示例
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// https://down.dingtalk.com/ddmedia/xxxx.png?ddFrom=blackboard.pic
	CoverPicUrl *string                              `json:"coverPicUrl,omitempty" xml:"coverPicUrl,omitempty"`
	DepNameList []*string                            `json:"depNameList,omitempty" xml:"depNameList,omitempty" type:"Repeated"`
	DeptList    []*GetBlackboardResponseBodyDeptList `json:"deptList,omitempty" xml:"deptList,omitempty" type:"Repeated"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2025-01-01 00:00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// fbeaxxxxxxxxxxxxxxxxxxxxxxxxe292
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 0
	IsPushTop *int64 `json:"isPushTop,omitempty" xml:"isPushTop,omitempty"`
	// example:
	//
	// 0
	PrivateLevel *int64 `json:"privateLevel,omitempty" xml:"privateLevel,omitempty"`
	// example:
	//
	// 1
	ReadCount *int64 `json:"readCount,omitempty" xml:"readCount,omitempty"`
	// example:
	//
	// manager01
	SenderStaffId *string `json:"senderStaffId,omitempty" xml:"senderStaffId,omitempty"`
	// example:
	//
	// 公告标题实例
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 1
	UnReadCount  *int64                               `json:"unReadCount,omitempty" xml:"unReadCount,omitempty"`
	UserList     []*GetBlackboardResponseBodyUserList `json:"userList,omitempty" xml:"userList,omitempty" type:"Repeated"`
	UserNameList []*string                            `json:"userNameList,omitempty" xml:"userNameList,omitempty" type:"Repeated"`
}

func (s GetBlackboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBlackboardResponseBody) GoString() string {
	return s.String()
}

func (s *GetBlackboardResponseBody) SetAttachments(v []*GetBlackboardResponseBodyAttachments) *GetBlackboardResponseBody {
	s.Attachments = v
	return s
}

func (s *GetBlackboardResponseBody) SetCategoryId(v string) *GetBlackboardResponseBody {
	s.CategoryId = &v
	return s
}

func (s *GetBlackboardResponseBody) SetCategoryName(v string) *GetBlackboardResponseBody {
	s.CategoryName = &v
	return s
}

func (s *GetBlackboardResponseBody) SetContent(v string) *GetBlackboardResponseBody {
	s.Content = &v
	return s
}

func (s *GetBlackboardResponseBody) SetCoverPicUrl(v string) *GetBlackboardResponseBody {
	s.CoverPicUrl = &v
	return s
}

func (s *GetBlackboardResponseBody) SetDepNameList(v []*string) *GetBlackboardResponseBody {
	s.DepNameList = v
	return s
}

func (s *GetBlackboardResponseBody) SetDeptList(v []*GetBlackboardResponseBodyDeptList) *GetBlackboardResponseBody {
	s.DeptList = v
	return s
}

func (s *GetBlackboardResponseBody) SetGmtCreate(v string) *GetBlackboardResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetBlackboardResponseBody) SetGmtModified(v string) *GetBlackboardResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetBlackboardResponseBody) SetId(v string) *GetBlackboardResponseBody {
	s.Id = &v
	return s
}

func (s *GetBlackboardResponseBody) SetIsPushTop(v int64) *GetBlackboardResponseBody {
	s.IsPushTop = &v
	return s
}

func (s *GetBlackboardResponseBody) SetPrivateLevel(v int64) *GetBlackboardResponseBody {
	s.PrivateLevel = &v
	return s
}

func (s *GetBlackboardResponseBody) SetReadCount(v int64) *GetBlackboardResponseBody {
	s.ReadCount = &v
	return s
}

func (s *GetBlackboardResponseBody) SetSenderStaffId(v string) *GetBlackboardResponseBody {
	s.SenderStaffId = &v
	return s
}

func (s *GetBlackboardResponseBody) SetTitle(v string) *GetBlackboardResponseBody {
	s.Title = &v
	return s
}

func (s *GetBlackboardResponseBody) SetUnReadCount(v int64) *GetBlackboardResponseBody {
	s.UnReadCount = &v
	return s
}

func (s *GetBlackboardResponseBody) SetUserList(v []*GetBlackboardResponseBodyUserList) *GetBlackboardResponseBody {
	s.UserList = v
	return s
}

func (s *GetBlackboardResponseBody) SetUserNameList(v []*string) *GetBlackboardResponseBody {
	s.UserNameList = v
	return s
}

type GetBlackboardResponseBodyAttachments struct {
	// example:
	//
	// xxxx
	DentryId *string `json:"dentryId,omitempty" xml:"dentryId,omitempty"`
	// example:
	//
	// 附件.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// xxxx
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s GetBlackboardResponseBodyAttachments) String() string {
	return tea.Prettify(s)
}

func (s GetBlackboardResponseBodyAttachments) GoString() string {
	return s.String()
}

func (s *GetBlackboardResponseBodyAttachments) SetDentryId(v string) *GetBlackboardResponseBodyAttachments {
	s.DentryId = &v
	return s
}

func (s *GetBlackboardResponseBodyAttachments) SetFileName(v string) *GetBlackboardResponseBodyAttachments {
	s.FileName = &v
	return s
}

func (s *GetBlackboardResponseBodyAttachments) SetFileType(v string) *GetBlackboardResponseBodyAttachments {
	s.FileType = &v
	return s
}

func (s *GetBlackboardResponseBodyAttachments) SetSpaceId(v string) *GetBlackboardResponseBodyAttachments {
	s.SpaceId = &v
	return s
}

type GetBlackboardResponseBodyDeptList struct {
	// example:
	//
	// example_dept_id
	DeptId *string `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// example:
	//
	// xxxx部门
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetBlackboardResponseBodyDeptList) String() string {
	return tea.Prettify(s)
}

func (s GetBlackboardResponseBodyDeptList) GoString() string {
	return s.String()
}

func (s *GetBlackboardResponseBodyDeptList) SetDeptId(v string) *GetBlackboardResponseBodyDeptList {
	s.DeptId = &v
	return s
}

func (s *GetBlackboardResponseBodyDeptList) SetName(v string) *GetBlackboardResponseBodyDeptList {
	s.Name = &v
	return s
}

type GetBlackboardResponseBodyUserList struct {
	// example:
	//
	// dingxxxx
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// example:
	//
	// 示例员工名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// manager01
	StaffId *string `json:"staffId,omitempty" xml:"staffId,omitempty"`
}

func (s GetBlackboardResponseBodyUserList) String() string {
	return tea.Prettify(s)
}

func (s GetBlackboardResponseBodyUserList) GoString() string {
	return s.String()
}

func (s *GetBlackboardResponseBodyUserList) SetCorpId(v string) *GetBlackboardResponseBodyUserList {
	s.CorpId = &v
	return s
}

func (s *GetBlackboardResponseBodyUserList) SetName(v string) *GetBlackboardResponseBodyUserList {
	s.Name = &v
	return s
}

func (s *GetBlackboardResponseBodyUserList) SetStaffId(v string) *GetBlackboardResponseBodyUserList {
	s.StaffId = &v
	return s
}

type GetBlackboardResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBlackboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBlackboardResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBlackboardResponse) GoString() string {
	return s.String()
}

func (s *GetBlackboardResponse) SetHeaders(v map[string]*string) *GetBlackboardResponse {
	s.Headers = v
	return s
}

func (s *GetBlackboardResponse) SetStatusCode(v int32) *GetBlackboardResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBlackboardResponse) SetBody(v *GetBlackboardResponseBody) *GetBlackboardResponse {
	s.Body = v
	return s
}

type QueryBlackboardReadUnReadHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryBlackboardReadUnReadHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardReadUnReadHeaders) GoString() string {
	return s.String()
}

func (s *QueryBlackboardReadUnReadHeaders) SetCommonHeaders(v map[string]*string) *QueryBlackboardReadUnReadHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryBlackboardReadUnReadHeaders) SetXAcsDingtalkAccessToken(v string) *QueryBlackboardReadUnReadHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryBlackboardReadUnReadRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 49dc87dc1b30cd099b13a
	BlackboardId *string `json:"blackboardId,omitempty" xml:"blackboardId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 200
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// xb1dc
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// manager01
	OperationUserId *string `json:"operationUserId,omitempty" xml:"operationUserId,omitempty"`
}

func (s QueryBlackboardReadUnReadRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardReadUnReadRequest) GoString() string {
	return s.String()
}

func (s *QueryBlackboardReadUnReadRequest) SetBlackboardId(v string) *QueryBlackboardReadUnReadRequest {
	s.BlackboardId = &v
	return s
}

func (s *QueryBlackboardReadUnReadRequest) SetMaxResults(v int32) *QueryBlackboardReadUnReadRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryBlackboardReadUnReadRequest) SetNextToken(v string) *QueryBlackboardReadUnReadRequest {
	s.NextToken = &v
	return s
}

func (s *QueryBlackboardReadUnReadRequest) SetOperationUserId(v string) *QueryBlackboardReadUnReadRequest {
	s.OperationUserId = &v
	return s
}

type QueryBlackboardReadUnReadResponseBody struct {
	NextToken *string                                       `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Users     []*QueryBlackboardReadUnReadResponseBodyUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s QueryBlackboardReadUnReadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardReadUnReadResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBlackboardReadUnReadResponseBody) SetNextToken(v string) *QueryBlackboardReadUnReadResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryBlackboardReadUnReadResponseBody) SetUsers(v []*QueryBlackboardReadUnReadResponseBodyUsers) *QueryBlackboardReadUnReadResponseBody {
	s.Users = v
	return s
}

type QueryBlackboardReadUnReadResponseBodyUsers struct {
	// example:
	//
	// true
	Read          *string `json:"read,omitempty" xml:"read,omitempty"`
	ReadTimestamp *int64  `json:"readTimestamp,omitempty" xml:"readTimestamp,omitempty"`
	// example:
	//
	// 12039
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryBlackboardReadUnReadResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardReadUnReadResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *QueryBlackboardReadUnReadResponseBodyUsers) SetRead(v string) *QueryBlackboardReadUnReadResponseBodyUsers {
	s.Read = &v
	return s
}

func (s *QueryBlackboardReadUnReadResponseBodyUsers) SetReadTimestamp(v int64) *QueryBlackboardReadUnReadResponseBodyUsers {
	s.ReadTimestamp = &v
	return s
}

func (s *QueryBlackboardReadUnReadResponseBodyUsers) SetUserId(v string) *QueryBlackboardReadUnReadResponseBodyUsers {
	s.UserId = &v
	return s
}

type QueryBlackboardReadUnReadResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBlackboardReadUnReadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBlackboardReadUnReadResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardReadUnReadResponse) GoString() string {
	return s.String()
}

func (s *QueryBlackboardReadUnReadResponse) SetHeaders(v map[string]*string) *QueryBlackboardReadUnReadResponse {
	s.Headers = v
	return s
}

func (s *QueryBlackboardReadUnReadResponse) SetStatusCode(v int32) *QueryBlackboardReadUnReadResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBlackboardReadUnReadResponse) SetBody(v *QueryBlackboardReadUnReadResponseBody) *QueryBlackboardReadUnReadResponse {
	s.Body = v
	return s
}

type QueryBlackboardSpaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryBlackboardSpaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardSpaceHeaders) GoString() string {
	return s.String()
}

func (s *QueryBlackboardSpaceHeaders) SetCommonHeaders(v map[string]*string) *QueryBlackboardSpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryBlackboardSpaceHeaders) SetXAcsDingtalkAccessToken(v string) *QueryBlackboardSpaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryBlackboardSpaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// manager01
	OperationUserId *string `json:"operationUserId,omitempty" xml:"operationUserId,omitempty"`
}

func (s QueryBlackboardSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardSpaceRequest) GoString() string {
	return s.String()
}

func (s *QueryBlackboardSpaceRequest) SetOperationUserId(v string) *QueryBlackboardSpaceRequest {
	s.OperationUserId = &v
	return s
}

type QueryBlackboardSpaceResponseBody struct {
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s QueryBlackboardSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBlackboardSpaceResponseBody) SetSpaceId(v string) *QueryBlackboardSpaceResponseBody {
	s.SpaceId = &v
	return s
}

type QueryBlackboardSpaceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBlackboardSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBlackboardSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBlackboardSpaceResponse) GoString() string {
	return s.String()
}

func (s *QueryBlackboardSpaceResponse) SetHeaders(v map[string]*string) *QueryBlackboardSpaceResponse {
	s.Headers = v
	return s
}

func (s *QueryBlackboardSpaceResponse) SetStatusCode(v int32) *QueryBlackboardSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBlackboardSpaceResponse) SetBody(v *QueryBlackboardSpaceResponseBody) *QueryBlackboardSpaceResponse {
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
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	client.EndpointRule = tea.String("")
	if tea.BoolValue(util.Empty(client.Endpoint)) {
		client.Endpoint = tea.String("api.dingtalk.com")
	}

	return nil
}

// Summary:
//
// 获取公告详情
//
// @param request - GetBlackboardRequest
//
// @param headers - GetBlackboardHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBlackboardResponse
func (client *Client) GetBlackboardWithOptions(request *GetBlackboardRequest, headers *GetBlackboardHeaders, runtime *util.RuntimeOptions) (_result *GetBlackboardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BlackboardId)) {
		query["blackboardId"] = request.BlackboardId
	}

	if !tea.BoolValue(util.IsUnset(request.OperationUserId)) {
		query["operationUserId"] = request.OperationUserId
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
		Action:      tea.String("GetBlackboard"),
		Version:     tea.String("blackboard_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/blackboard/get_blackboard"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBlackboardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取公告详情
//
// @param request - GetBlackboardRequest
//
// @return GetBlackboardResponse
func (client *Client) GetBlackboard(request *GetBlackboardRequest) (_result *GetBlackboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetBlackboardHeaders{}
	_result = &GetBlackboardResponse{}
	_body, _err := client.GetBlackboardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询公告已读未读人员列表
//
// @param request - QueryBlackboardReadUnReadRequest
//
// @param headers - QueryBlackboardReadUnReadHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBlackboardReadUnReadResponse
func (client *Client) QueryBlackboardReadUnReadWithOptions(request *QueryBlackboardReadUnReadRequest, headers *QueryBlackboardReadUnReadHeaders, runtime *util.RuntimeOptions) (_result *QueryBlackboardReadUnReadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BlackboardId)) {
		query["blackboardId"] = request.BlackboardId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OperationUserId)) {
		query["operationUserId"] = request.OperationUserId
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
		Action:      tea.String("QueryBlackboardReadUnRead"),
		Version:     tea.String("blackboard_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/blackboard/readers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBlackboardReadUnReadResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询公告已读未读人员列表
//
// @param request - QueryBlackboardReadUnReadRequest
//
// @return QueryBlackboardReadUnReadResponse
func (client *Client) QueryBlackboardReadUnRead(request *QueryBlackboardReadUnReadRequest) (_result *QueryBlackboardReadUnReadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryBlackboardReadUnReadHeaders{}
	_result = &QueryBlackboardReadUnReadResponse{}
	_body, _err := client.QueryBlackboardReadUnReadWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取公告钉盘空间信息
//
// @param request - QueryBlackboardSpaceRequest
//
// @param headers - QueryBlackboardSpaceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBlackboardSpaceResponse
func (client *Client) QueryBlackboardSpaceWithOptions(request *QueryBlackboardSpaceRequest, headers *QueryBlackboardSpaceHeaders, runtime *util.RuntimeOptions) (_result *QueryBlackboardSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperationUserId)) {
		query["operationUserId"] = request.OperationUserId
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
		Action:      tea.String("QueryBlackboardSpace"),
		Version:     tea.String("blackboard_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/blackboard/spaces"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBlackboardSpaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取公告钉盘空间信息
//
// @param request - QueryBlackboardSpaceRequest
//
// @return QueryBlackboardSpaceResponse
func (client *Client) QueryBlackboardSpace(request *QueryBlackboardSpaceRequest) (_result *QueryBlackboardSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryBlackboardSpaceHeaders{}
	_result = &QueryBlackboardSpaceResponse{}
	_body, _err := client.QueryBlackboardSpaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
