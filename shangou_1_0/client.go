// This file is auto-generated, don't edit it. Thanks.
package shangou_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddCateringCommentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddCateringCommentHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddCateringCommentHeaders) GoString() string {
	return s.String()
}

func (s *AddCateringCommentHeaders) SetCommonHeaders(v map[string]*string) *AddCateringCommentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddCateringCommentHeaders) SetXAcsDingtalkAccessToken(v string) *AddCateringCommentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddCateringCommentRequest struct {
	// This parameter is required.
	CommentId   *string `json:"commentId,omitempty" xml:"commentId,omitempty"`
	OrderId     *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	RateContent *string `json:"rateContent,omitempty" xml:"rateContent,omitempty"`
	RatedAt     *int64  `json:"ratedAt,omitempty" xml:"ratedAt,omitempty"`
	Rating      *int32  `json:"rating,omitempty" xml:"rating,omitempty"`
	// This parameter is required.
	ShopId *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
	// This parameter is required.
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	Status *int32  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s AddCateringCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCateringCommentRequest) GoString() string {
	return s.String()
}

func (s *AddCateringCommentRequest) SetCommentId(v string) *AddCateringCommentRequest {
	s.CommentId = &v
	return s
}

func (s *AddCateringCommentRequest) SetOrderId(v string) *AddCateringCommentRequest {
	s.OrderId = &v
	return s
}

func (s *AddCateringCommentRequest) SetRateContent(v string) *AddCateringCommentRequest {
	s.RateContent = &v
	return s
}

func (s *AddCateringCommentRequest) SetRatedAt(v int64) *AddCateringCommentRequest {
	s.RatedAt = &v
	return s
}

func (s *AddCateringCommentRequest) SetRating(v int32) *AddCateringCommentRequest {
	s.Rating = &v
	return s
}

func (s *AddCateringCommentRequest) SetShopId(v string) *AddCateringCommentRequest {
	s.ShopId = &v
	return s
}

func (s *AddCateringCommentRequest) SetSource(v string) *AddCateringCommentRequest {
	s.Source = &v
	return s
}

func (s *AddCateringCommentRequest) SetStatus(v int32) *AddCateringCommentRequest {
	s.Status = &v
	return s
}

type AddCateringCommentResponseBody struct {
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result    *string `json:"result,omitempty" xml:"result,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddCateringCommentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCateringCommentResponseBody) GoString() string {
	return s.String()
}

func (s *AddCateringCommentResponseBody) SetErrorCode(v string) *AddCateringCommentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddCateringCommentResponseBody) SetErrorMsg(v string) *AddCateringCommentResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *AddCateringCommentResponseBody) SetResult(v string) *AddCateringCommentResponseBody {
	s.Result = &v
	return s
}

func (s *AddCateringCommentResponseBody) SetSuccess(v bool) *AddCateringCommentResponseBody {
	s.Success = &v
	return s
}

type AddCateringCommentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCateringCommentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCateringCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCateringCommentResponse) GoString() string {
	return s.String()
}

func (s *AddCateringCommentResponse) SetHeaders(v map[string]*string) *AddCateringCommentResponse {
	s.Headers = v
	return s
}

func (s *AddCateringCommentResponse) SetStatusCode(v int32) *AddCateringCommentResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCateringCommentResponse) SetBody(v *AddCateringCommentResponseBody) *AddCateringCommentResponse {
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
// 新增餐饮评价数据
//
// @param request - AddCateringCommentRequest
//
// @param headers - AddCateringCommentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCateringCommentResponse
func (client *Client) AddCateringCommentWithOptions(request *AddCateringCommentRequest, headers *AddCateringCommentHeaders, runtime *util.RuntimeOptions) (_result *AddCateringCommentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommentId)) {
		body["commentId"] = request.CommentId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		body["orderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.RateContent)) {
		body["rateContent"] = request.RateContent
	}

	if !tea.BoolValue(util.IsUnset(request.RatedAt)) {
		body["ratedAt"] = request.RatedAt
	}

	if !tea.BoolValue(util.IsUnset(request.Rating)) {
		body["rating"] = request.Rating
	}

	if !tea.BoolValue(util.IsUnset(request.ShopId)) {
		body["shopId"] = request.ShopId
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
	params := &openapi.Params{
		Action:      tea.String("AddCateringComment"),
		Version:     tea.String("shangou_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/shangou/comment/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddCateringCommentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增餐饮评价数据
//
// @param request - AddCateringCommentRequest
//
// @return AddCateringCommentResponse
func (client *Client) AddCateringComment(request *AddCateringCommentRequest) (_result *AddCateringCommentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddCateringCommentHeaders{}
	_result = &AddCateringCommentResponse{}
	_body, _err := client.AddCateringCommentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
