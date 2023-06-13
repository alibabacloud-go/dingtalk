// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package activity_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateActivityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateActivityHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateActivityHeaders) GoString() string {
	return s.String()
}

func (s *CreateActivityHeaders) SetCommonHeaders(v map[string]*string) *CreateActivityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateActivityHeaders) SetXAcsDingtalkAccessToken(v string) *CreateActivityHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateActivityRequest struct {
	Detail *CreateActivityRequestDetail `json:"detail,omitempty" xml:"detail,omitempty" type:"Struct"`
}

func (s CreateActivityRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateActivityRequest) GoString() string {
	return s.String()
}

func (s *CreateActivityRequest) SetDetail(v *CreateActivityRequestDetail) *CreateActivityRequest {
	s.Detail = v
	return s
}

type CreateActivityRequestDetail struct {
	Address       *CreateActivityRequestDetailAddress `json:"address,omitempty" xml:"address,omitempty" type:"Struct"`
	BannerMediaId *string                             `json:"bannerMediaId,omitempty" xml:"bannerMediaId,omitempty"`
	EndTime       *int64                              `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ForeignId     *string                             `json:"foreignId,omitempty" xml:"foreignId,omitempty"`
	Industry      *string                             `json:"industry,omitempty" xml:"industry,omitempty"`
	RoleName      *string                             `json:"roleName,omitempty" xml:"roleName,omitempty"`
	Source        *string                             `json:"source,omitempty" xml:"source,omitempty"`
	StartTime     *int64                              `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Title         *string                             `json:"title,omitempty" xml:"title,omitempty"`
	Type          *int32                              `json:"type,omitempty" xml:"type,omitempty"`
	Url           *string                             `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateActivityRequestDetail) String() string {
	return tea.Prettify(s)
}

func (s CreateActivityRequestDetail) GoString() string {
	return s.String()
}

func (s *CreateActivityRequestDetail) SetAddress(v *CreateActivityRequestDetailAddress) *CreateActivityRequestDetail {
	s.Address = v
	return s
}

func (s *CreateActivityRequestDetail) SetBannerMediaId(v string) *CreateActivityRequestDetail {
	s.BannerMediaId = &v
	return s
}

func (s *CreateActivityRequestDetail) SetEndTime(v int64) *CreateActivityRequestDetail {
	s.EndTime = &v
	return s
}

func (s *CreateActivityRequestDetail) SetForeignId(v string) *CreateActivityRequestDetail {
	s.ForeignId = &v
	return s
}

func (s *CreateActivityRequestDetail) SetIndustry(v string) *CreateActivityRequestDetail {
	s.Industry = &v
	return s
}

func (s *CreateActivityRequestDetail) SetRoleName(v string) *CreateActivityRequestDetail {
	s.RoleName = &v
	return s
}

func (s *CreateActivityRequestDetail) SetSource(v string) *CreateActivityRequestDetail {
	s.Source = &v
	return s
}

func (s *CreateActivityRequestDetail) SetStartTime(v int64) *CreateActivityRequestDetail {
	s.StartTime = &v
	return s
}

func (s *CreateActivityRequestDetail) SetTitle(v string) *CreateActivityRequestDetail {
	s.Title = &v
	return s
}

func (s *CreateActivityRequestDetail) SetType(v int32) *CreateActivityRequestDetail {
	s.Type = &v
	return s
}

func (s *CreateActivityRequestDetail) SetUrl(v string) *CreateActivityRequestDetail {
	s.Url = &v
	return s
}

type CreateActivityRequestDetailAddress struct {
	District *string `json:"district,omitempty" xml:"district,omitempty"`
	Lat      *string `json:"lat,omitempty" xml:"lat,omitempty"`
	Lng      *string `json:"lng,omitempty" xml:"lng,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateActivityRequestDetailAddress) String() string {
	return tea.Prettify(s)
}

func (s CreateActivityRequestDetailAddress) GoString() string {
	return s.String()
}

func (s *CreateActivityRequestDetailAddress) SetDistrict(v string) *CreateActivityRequestDetailAddress {
	s.District = &v
	return s
}

func (s *CreateActivityRequestDetailAddress) SetLat(v string) *CreateActivityRequestDetailAddress {
	s.Lat = &v
	return s
}

func (s *CreateActivityRequestDetailAddress) SetLng(v string) *CreateActivityRequestDetailAddress {
	s.Lng = &v
	return s
}

func (s *CreateActivityRequestDetailAddress) SetName(v string) *CreateActivityRequestDetailAddress {
	s.Name = &v
	return s
}

type CreateActivityResponseBody struct {
	ActivityId *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
}

func (s CreateActivityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateActivityResponseBody) GoString() string {
	return s.String()
}

func (s *CreateActivityResponseBody) SetActivityId(v string) *CreateActivityResponseBody {
	s.ActivityId = &v
	return s
}

type CreateActivityResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateActivityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateActivityResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateActivityResponse) GoString() string {
	return s.String()
}

func (s *CreateActivityResponse) SetHeaders(v map[string]*string) *CreateActivityResponse {
	s.Headers = v
	return s
}

func (s *CreateActivityResponse) SetStatusCode(v int32) *CreateActivityResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateActivityResponse) SetBody(v *CreateActivityResponseBody) *CreateActivityResponse {
	s.Body = v
	return s
}

type ListActivityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListActivityHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListActivityHeaders) GoString() string {
	return s.String()
}

func (s *ListActivityHeaders) SetCommonHeaders(v map[string]*string) *ListActivityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListActivityHeaders) SetXAcsDingtalkAccessToken(v string) *ListActivityHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListActivityRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListActivityRequest) String() string {
	return tea.Prettify(s)
}

func (s ListActivityRequest) GoString() string {
	return s.String()
}

func (s *ListActivityRequest) SetMaxResults(v int32) *ListActivityRequest {
	s.MaxResults = &v
	return s
}

func (s *ListActivityRequest) SetNextToken(v string) *ListActivityRequest {
	s.NextToken = &v
	return s
}

type ListActivityResponseBody struct {
	List       []*ListActivityResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	MaxResults *string                         `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string                         `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListActivityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListActivityResponseBody) GoString() string {
	return s.String()
}

func (s *ListActivityResponseBody) SetList(v []*ListActivityResponseBodyList) *ListActivityResponseBody {
	s.List = v
	return s
}

func (s *ListActivityResponseBody) SetMaxResults(v string) *ListActivityResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListActivityResponseBody) SetNextToken(v string) *ListActivityResponseBody {
	s.NextToken = &v
	return s
}

type ListActivityResponseBodyList struct {
	ActivityId    *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
	BannerMediaId *string `json:"bannerMediaId,omitempty" xml:"bannerMediaId,omitempty"`
	EndTime       *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	ForeignId     *string `json:"foreignId,omitempty" xml:"foreignId,omitempty"`
	StartTime     *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Status        *string `json:"status,omitempty" xml:"status,omitempty"`
	Title         *string `json:"title,omitempty" xml:"title,omitempty"`
	Type          *string `json:"type,omitempty" xml:"type,omitempty"`
	Url           *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListActivityResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s ListActivityResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListActivityResponseBodyList) SetActivityId(v string) *ListActivityResponseBodyList {
	s.ActivityId = &v
	return s
}

func (s *ListActivityResponseBodyList) SetBannerMediaId(v string) *ListActivityResponseBodyList {
	s.BannerMediaId = &v
	return s
}

func (s *ListActivityResponseBodyList) SetEndTime(v int64) *ListActivityResponseBodyList {
	s.EndTime = &v
	return s
}

func (s *ListActivityResponseBodyList) SetForeignId(v string) *ListActivityResponseBodyList {
	s.ForeignId = &v
	return s
}

func (s *ListActivityResponseBodyList) SetStartTime(v int64) *ListActivityResponseBodyList {
	s.StartTime = &v
	return s
}

func (s *ListActivityResponseBodyList) SetStatus(v string) *ListActivityResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListActivityResponseBodyList) SetTitle(v string) *ListActivityResponseBodyList {
	s.Title = &v
	return s
}

func (s *ListActivityResponseBodyList) SetType(v string) *ListActivityResponseBodyList {
	s.Type = &v
	return s
}

func (s *ListActivityResponseBodyList) SetUrl(v string) *ListActivityResponseBodyList {
	s.Url = &v
	return s
}

type ListActivityResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListActivityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListActivityResponse) String() string {
	return tea.Prettify(s)
}

func (s ListActivityResponse) GoString() string {
	return s.String()
}

func (s *ListActivityResponse) SetHeaders(v map[string]*string) *ListActivityResponse {
	s.Headers = v
	return s
}

func (s *ListActivityResponse) SetStatusCode(v int32) *ListActivityResponse {
	s.StatusCode = &v
	return s
}

func (s *ListActivityResponse) SetBody(v *ListActivityResponseBody) *ListActivityResponse {
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

func (client *Client) CreateActivityWithOptions(request *CreateActivityRequest, headers *CreateActivityHeaders, runtime *util.RuntimeOptions) (_result *CreateActivityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Detail)) {
		body["detail"] = request.Detail
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
		Action:      tea.String("CreateActivity"),
		Version:     tea.String("activity_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/activity/meta"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateActivityResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateActivity(request *CreateActivityRequest) (_result *CreateActivityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateActivityHeaders{}
	_result = &CreateActivityResponse{}
	_body, _err := client.CreateActivityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListActivityWithOptions(request *ListActivityRequest, headers *ListActivityHeaders, runtime *util.RuntimeOptions) (_result *ListActivityResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("ListActivity"),
		Version:     tea.String("activity_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/activity/metaLists"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListActivityResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListActivity(request *ListActivityRequest) (_result *ListActivityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListActivityHeaders{}
	_result = &ListActivityResponse{}
	_body, _err := client.ListActivityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
