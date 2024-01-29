// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package wms_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type QueryGoodsListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryGoodsListHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryGoodsListHeaders) GoString() string {
	return s.String()
}

func (s *QueryGoodsListHeaders) SetCommonHeaders(v map[string]*string) *QueryGoodsListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGoodsListHeaders) SetXAcsDingtalkAccessToken(v string) *QueryGoodsListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryGoodsListRequest struct {
	EndTimeInMills   *int64 `json:"endTimeInMills,omitempty" xml:"endTimeInMills,omitempty"`
	MaxResults       *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken        *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	StartTimeInMills *int64 `json:"startTimeInMills,omitempty" xml:"startTimeInMills,omitempty"`
}

func (s QueryGoodsListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGoodsListRequest) GoString() string {
	return s.String()
}

func (s *QueryGoodsListRequest) SetEndTimeInMills(v int64) *QueryGoodsListRequest {
	s.EndTimeInMills = &v
	return s
}

func (s *QueryGoodsListRequest) SetMaxResults(v int64) *QueryGoodsListRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryGoodsListRequest) SetNextToken(v int64) *QueryGoodsListRequest {
	s.NextToken = &v
	return s
}

func (s *QueryGoodsListRequest) SetStartTimeInMills(v int64) *QueryGoodsListRequest {
	s.StartTimeInMills = &v
	return s
}

type QueryGoodsListResponseBody struct {
	Result  *QueryGoodsListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryGoodsListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGoodsListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGoodsListResponseBody) SetResult(v *QueryGoodsListResponseBodyResult) *QueryGoodsListResponseBody {
	s.Result = v
	return s
}

func (s *QueryGoodsListResponseBody) SetSuccess(v bool) *QueryGoodsListResponseBody {
	s.Success = &v
	return s
}

type QueryGoodsListResponseBodyResult struct {
	HasMore    *bool                                   `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*QueryGoodsListResponseBodyResultList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	MaxResults *int64                                  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string                                 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryGoodsListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryGoodsListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryGoodsListResponseBodyResult) SetHasMore(v bool) *QueryGoodsListResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *QueryGoodsListResponseBodyResult) SetList(v []*QueryGoodsListResponseBodyResultList) *QueryGoodsListResponseBodyResult {
	s.List = v
	return s
}

func (s *QueryGoodsListResponseBodyResult) SetMaxResults(v int64) *QueryGoodsListResponseBodyResult {
	s.MaxResults = &v
	return s
}

func (s *QueryGoodsListResponseBodyResult) SetNextToken(v string) *QueryGoodsListResponseBodyResult {
	s.NextToken = &v
	return s
}

type QueryGoodsListResponseBodyResultList struct {
	GoodsName    *string `json:"goodsName,omitempty" xml:"goodsName,omitempty"`
	GoodsNo      *string `json:"goodsNo,omitempty" xml:"goodsNo,omitempty"`
	InstanceId   *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	ProductSpecs *string `json:"productSpecs,omitempty" xml:"productSpecs,omitempty"`
	Unit         *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QueryGoodsListResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s QueryGoodsListResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *QueryGoodsListResponseBodyResultList) SetGoodsName(v string) *QueryGoodsListResponseBodyResultList {
	s.GoodsName = &v
	return s
}

func (s *QueryGoodsListResponseBodyResultList) SetGoodsNo(v string) *QueryGoodsListResponseBodyResultList {
	s.GoodsNo = &v
	return s
}

func (s *QueryGoodsListResponseBodyResultList) SetInstanceId(v string) *QueryGoodsListResponseBodyResultList {
	s.InstanceId = &v
	return s
}

func (s *QueryGoodsListResponseBodyResultList) SetProductSpecs(v string) *QueryGoodsListResponseBodyResultList {
	s.ProductSpecs = &v
	return s
}

func (s *QueryGoodsListResponseBodyResultList) SetUnit(v string) *QueryGoodsListResponseBodyResultList {
	s.Unit = &v
	return s
}

type QueryGoodsListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGoodsListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGoodsListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGoodsListResponse) GoString() string {
	return s.String()
}

func (s *QueryGoodsListResponse) SetHeaders(v map[string]*string) *QueryGoodsListResponse {
	s.Headers = v
	return s
}

func (s *QueryGoodsListResponse) SetStatusCode(v int32) *QueryGoodsListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGoodsListResponse) SetBody(v *QueryGoodsListResponseBody) *QueryGoodsListResponse {
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

func (client *Client) QueryGoodsListWithOptions(request *QueryGoodsListRequest, headers *QueryGoodsListHeaders, runtime *util.RuntimeOptions) (_result *QueryGoodsListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimeInMills)) {
		query["endTimeInMills"] = request.EndTimeInMills
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimeInMills)) {
		query["startTimeInMills"] = request.StartTimeInMills
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
		Action:      tea.String("QueryGoodsList"),
		Version:     tea.String("wms_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/wms/goods"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGoodsListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryGoodsList(request *QueryGoodsListRequest) (_result *QueryGoodsListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryGoodsListHeaders{}
	_result = &QueryGoodsListResponse{}
	_body, _err := client.QueryGoodsListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
