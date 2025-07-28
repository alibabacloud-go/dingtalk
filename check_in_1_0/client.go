// This file is auto-generated, don't edit it. Thanks.
package check_in_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetCheckinRecordByUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCheckinRecordByUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCheckinRecordByUserHeaders) GoString() string {
	return s.String()
}

func (s *GetCheckinRecordByUserHeaders) SetCommonHeaders(v map[string]*string) *GetCheckinRecordByUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCheckinRecordByUserHeaders) SetXAcsDingtalkAccessToken(v string) *GetCheckinRecordByUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCheckinRecordByUserRequest struct {
	// This parameter is required.
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// This parameter is required.
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// This parameter is required.
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s GetCheckinRecordByUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCheckinRecordByUserRequest) GoString() string {
	return s.String()
}

func (s *GetCheckinRecordByUserRequest) SetEndTime(v int64) *GetCheckinRecordByUserRequest {
	s.EndTime = &v
	return s
}

func (s *GetCheckinRecordByUserRequest) SetMaxResults(v int64) *GetCheckinRecordByUserRequest {
	s.MaxResults = &v
	return s
}

func (s *GetCheckinRecordByUserRequest) SetNextToken(v int64) *GetCheckinRecordByUserRequest {
	s.NextToken = &v
	return s
}

func (s *GetCheckinRecordByUserRequest) SetOperatorUserId(v string) *GetCheckinRecordByUserRequest {
	s.OperatorUserId = &v
	return s
}

func (s *GetCheckinRecordByUserRequest) SetStartTime(v int64) *GetCheckinRecordByUserRequest {
	s.StartTime = &v
	return s
}

func (s *GetCheckinRecordByUserRequest) SetUserIdList(v []*string) *GetCheckinRecordByUserRequest {
	s.UserIdList = v
	return s
}

type GetCheckinRecordByUserResponseBody struct {
	Result *GetCheckinRecordByUserResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetCheckinRecordByUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCheckinRecordByUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetCheckinRecordByUserResponseBody) SetResult(v *GetCheckinRecordByUserResponseBodyResult) *GetCheckinRecordByUserResponseBody {
	s.Result = v
	return s
}

type GetCheckinRecordByUserResponseBodyResult struct {
	NextToken *int64                                              `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	PageList  []*GetCheckinRecordByUserResponseBodyResultPageList `json:"pageList,omitempty" xml:"pageList,omitempty" type:"Repeated"`
}

func (s GetCheckinRecordByUserResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetCheckinRecordByUserResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetCheckinRecordByUserResponseBodyResult) SetNextToken(v int64) *GetCheckinRecordByUserResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *GetCheckinRecordByUserResponseBodyResult) SetPageList(v []*GetCheckinRecordByUserResponseBodyResultPageList) *GetCheckinRecordByUserResponseBodyResult {
	s.PageList = v
	return s
}

type GetCheckinRecordByUserResponseBodyResultPageList struct {
	CheckinTime    *int64                                                            `json:"checkinTime,omitempty" xml:"checkinTime,omitempty"`
	CustomDataList []*GetCheckinRecordByUserResponseBodyResultPageListCustomDataList `json:"customDataList,omitempty" xml:"customDataList,omitempty" type:"Repeated"`
	DetailPlace    *string                                                           `json:"detailPlace,omitempty" xml:"detailPlace,omitempty"`
	ImageList      []*string                                                         `json:"imageList,omitempty" xml:"imageList,omitempty" type:"Repeated"`
	Place          *string                                                           `json:"place,omitempty" xml:"place,omitempty"`
	Remark         *string                                                           `json:"remark,omitempty" xml:"remark,omitempty"`
	UserId         *string                                                           `json:"userId,omitempty" xml:"userId,omitempty"`
	VisitUser      *string                                                           `json:"visitUser,omitempty" xml:"visitUser,omitempty"`
}

func (s GetCheckinRecordByUserResponseBodyResultPageList) String() string {
	return tea.Prettify(s)
}

func (s GetCheckinRecordByUserResponseBodyResultPageList) GoString() string {
	return s.String()
}

func (s *GetCheckinRecordByUserResponseBodyResultPageList) SetCheckinTime(v int64) *GetCheckinRecordByUserResponseBodyResultPageList {
	s.CheckinTime = &v
	return s
}

func (s *GetCheckinRecordByUserResponseBodyResultPageList) SetCustomDataList(v []*GetCheckinRecordByUserResponseBodyResultPageListCustomDataList) *GetCheckinRecordByUserResponseBodyResultPageList {
	s.CustomDataList = v
	return s
}

func (s *GetCheckinRecordByUserResponseBodyResultPageList) SetDetailPlace(v string) *GetCheckinRecordByUserResponseBodyResultPageList {
	s.DetailPlace = &v
	return s
}

func (s *GetCheckinRecordByUserResponseBodyResultPageList) SetImageList(v []*string) *GetCheckinRecordByUserResponseBodyResultPageList {
	s.ImageList = v
	return s
}

func (s *GetCheckinRecordByUserResponseBodyResultPageList) SetPlace(v string) *GetCheckinRecordByUserResponseBodyResultPageList {
	s.Place = &v
	return s
}

func (s *GetCheckinRecordByUserResponseBodyResultPageList) SetRemark(v string) *GetCheckinRecordByUserResponseBodyResultPageList {
	s.Remark = &v
	return s
}

func (s *GetCheckinRecordByUserResponseBodyResultPageList) SetUserId(v string) *GetCheckinRecordByUserResponseBodyResultPageList {
	s.UserId = &v
	return s
}

func (s *GetCheckinRecordByUserResponseBodyResultPageList) SetVisitUser(v string) *GetCheckinRecordByUserResponseBodyResultPageList {
	s.VisitUser = &v
	return s
}

type GetCheckinRecordByUserResponseBodyResultPageListCustomDataList struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetCheckinRecordByUserResponseBodyResultPageListCustomDataList) String() string {
	return tea.Prettify(s)
}

func (s GetCheckinRecordByUserResponseBodyResultPageListCustomDataList) GoString() string {
	return s.String()
}

func (s *GetCheckinRecordByUserResponseBodyResultPageListCustomDataList) SetKey(v string) *GetCheckinRecordByUserResponseBodyResultPageListCustomDataList {
	s.Key = &v
	return s
}

func (s *GetCheckinRecordByUserResponseBodyResultPageListCustomDataList) SetLabel(v string) *GetCheckinRecordByUserResponseBodyResultPageListCustomDataList {
	s.Label = &v
	return s
}

func (s *GetCheckinRecordByUserResponseBodyResultPageListCustomDataList) SetValue(v string) *GetCheckinRecordByUserResponseBodyResultPageListCustomDataList {
	s.Value = &v
	return s
}

type GetCheckinRecordByUserResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCheckinRecordByUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCheckinRecordByUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCheckinRecordByUserResponse) GoString() string {
	return s.String()
}

func (s *GetCheckinRecordByUserResponse) SetHeaders(v map[string]*string) *GetCheckinRecordByUserResponse {
	s.Headers = v
	return s
}

func (s *GetCheckinRecordByUserResponse) SetStatusCode(v int32) *GetCheckinRecordByUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCheckinRecordByUserResponse) SetBody(v *GetCheckinRecordByUserResponseBody) *GetCheckinRecordByUserResponse {
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
// 调用本接口，获取用户签到记录
//
// @param request - GetCheckinRecordByUserRequest
//
// @param headers - GetCheckinRecordByUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCheckinRecordByUserResponse
func (client *Client) GetCheckinRecordByUserWithOptions(request *GetCheckinRecordByUserRequest, headers *GetCheckinRecordByUserHeaders, runtime *util.RuntimeOptions) (_result *GetCheckinRecordByUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdList)) {
		body["userIdList"] = request.UserIdList
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
	params := &openapi.Params{
		Action:      tea.String("GetCheckinRecordByUser"),
		Version:     tea.String("checkIn_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/checkIn/records/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCheckinRecordByUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用本接口，获取用户签到记录
//
// @param request - GetCheckinRecordByUserRequest
//
// @return GetCheckinRecordByUserResponse
func (client *Client) GetCheckinRecordByUser(request *GetCheckinRecordByUserRequest) (_result *GetCheckinRecordByUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCheckinRecordByUserHeaders{}
	_result = &GetCheckinRecordByUserResponse{}
	_body, _err := client.GetCheckinRecordByUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
