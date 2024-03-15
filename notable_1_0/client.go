// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package notable_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateSheetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateSheetHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateSheetHeaders) GoString() string {
	return s.String()
}

func (s *CreateSheetHeaders) SetCommonHeaders(v map[string]*string) *CreateSheetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSheetHeaders) SetXAcsDingtalkAccessToken(v string) *CreateSheetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateSheetRequest struct {
	Fields []*CreateSheetRequestFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
}

func (s CreateSheetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSheetRequest) GoString() string {
	return s.String()
}

func (s *CreateSheetRequest) SetFields(v []*CreateSheetRequestFields) *CreateSheetRequest {
	s.Fields = v
	return s
}

type CreateSheetRequestFields struct {
	Name     *string                `json:"name,omitempty" xml:"name,omitempty"`
	Property map[string]interface{} `json:"property,omitempty" xml:"property,omitempty"`
	Type     *string                `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateSheetRequestFields) String() string {
	return tea.Prettify(s)
}

func (s CreateSheetRequestFields) GoString() string {
	return s.String()
}

func (s *CreateSheetRequestFields) SetName(v string) *CreateSheetRequestFields {
	s.Name = &v
	return s
}

func (s *CreateSheetRequestFields) SetProperty(v map[string]interface{}) *CreateSheetRequestFields {
	s.Property = v
	return s
}

func (s *CreateSheetRequestFields) SetType(v string) *CreateSheetRequestFields {
	s.Type = &v
	return s
}

type CreateSheetResponseBody struct {
	Id   *string `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateSheetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSheetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSheetResponseBody) SetId(v string) *CreateSheetResponseBody {
	s.Id = &v
	return s
}

func (s *CreateSheetResponseBody) SetName(v string) *CreateSheetResponseBody {
	s.Name = &v
	return s
}

type CreateSheetResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSheetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSheetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSheetResponse) GoString() string {
	return s.String()
}

func (s *CreateSheetResponse) SetHeaders(v map[string]*string) *CreateSheetResponse {
	s.Headers = v
	return s
}

func (s *CreateSheetResponse) SetStatusCode(v int32) *CreateSheetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSheetResponse) SetBody(v *CreateSheetResponseBody) *CreateSheetResponse {
	s.Body = v
	return s
}

type DeleteSheetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteSheetHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteSheetHeaders) GoString() string {
	return s.String()
}

func (s *DeleteSheetHeaders) SetCommonHeaders(v map[string]*string) *DeleteSheetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteSheetHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteSheetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteSheetResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteSheetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSheetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSheetResponseBody) SetSuccess(v bool) *DeleteSheetResponseBody {
	s.Success = &v
	return s
}

type DeleteSheetResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSheetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSheetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSheetResponse) GoString() string {
	return s.String()
}

func (s *DeleteSheetResponse) SetHeaders(v map[string]*string) *DeleteSheetResponse {
	s.Headers = v
	return s
}

func (s *DeleteSheetResponse) SetStatusCode(v int32) *DeleteSheetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSheetResponse) SetBody(v *DeleteSheetResponseBody) *DeleteSheetResponse {
	s.Body = v
	return s
}

type GetAllSheetsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAllSheetsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAllSheetsHeaders) GoString() string {
	return s.String()
}

func (s *GetAllSheetsHeaders) SetCommonHeaders(v map[string]*string) *GetAllSheetsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAllSheetsHeaders) SetXAcsDingtalkAccessToken(v string) *GetAllSheetsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAllSheetsResponseBody struct {
	Value []*GetAllSheetsResponseBodyValue `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s GetAllSheetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAllSheetsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllSheetsResponseBody) SetValue(v []*GetAllSheetsResponseBodyValue) *GetAllSheetsResponseBody {
	s.Value = v
	return s
}

type GetAllSheetsResponseBodyValue struct {
	Fields []*GetAllSheetsResponseBodyValueFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	Id     *string                                `json:"id,omitempty" xml:"id,omitempty"`
	Name   *string                                `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetAllSheetsResponseBodyValue) String() string {
	return tea.Prettify(s)
}

func (s GetAllSheetsResponseBodyValue) GoString() string {
	return s.String()
}

func (s *GetAllSheetsResponseBodyValue) SetFields(v []*GetAllSheetsResponseBodyValueFields) *GetAllSheetsResponseBodyValue {
	s.Fields = v
	return s
}

func (s *GetAllSheetsResponseBodyValue) SetId(v string) *GetAllSheetsResponseBodyValue {
	s.Id = &v
	return s
}

func (s *GetAllSheetsResponseBodyValue) SetName(v string) *GetAllSheetsResponseBodyValue {
	s.Name = &v
	return s
}

type GetAllSheetsResponseBodyValueFields struct {
	Id       *string                `json:"id,omitempty" xml:"id,omitempty"`
	Name     *string                `json:"name,omitempty" xml:"name,omitempty"`
	Property map[string]interface{} `json:"property,omitempty" xml:"property,omitempty"`
	Type     *string                `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetAllSheetsResponseBodyValueFields) String() string {
	return tea.Prettify(s)
}

func (s GetAllSheetsResponseBodyValueFields) GoString() string {
	return s.String()
}

func (s *GetAllSheetsResponseBodyValueFields) SetId(v string) *GetAllSheetsResponseBodyValueFields {
	s.Id = &v
	return s
}

func (s *GetAllSheetsResponseBodyValueFields) SetName(v string) *GetAllSheetsResponseBodyValueFields {
	s.Name = &v
	return s
}

func (s *GetAllSheetsResponseBodyValueFields) SetProperty(v map[string]interface{}) *GetAllSheetsResponseBodyValueFields {
	s.Property = v
	return s
}

func (s *GetAllSheetsResponseBodyValueFields) SetType(v string) *GetAllSheetsResponseBodyValueFields {
	s.Type = &v
	return s
}

type GetAllSheetsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAllSheetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAllSheetsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllSheetsResponse) GoString() string {
	return s.String()
}

func (s *GetAllSheetsResponse) SetHeaders(v map[string]*string) *GetAllSheetsResponse {
	s.Headers = v
	return s
}

func (s *GetAllSheetsResponse) SetStatusCode(v int32) *GetAllSheetsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAllSheetsResponse) SetBody(v *GetAllSheetsResponseBody) *GetAllSheetsResponse {
	s.Body = v
	return s
}

type GetRecordHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetRecordHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetRecordHeaders) GoString() string {
	return s.String()
}

func (s *GetRecordHeaders) SetCommonHeaders(v map[string]*string) *GetRecordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRecordHeaders) SetXAcsDingtalkAccessToken(v string) *GetRecordHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetRecordResponseBody struct {
	Fields map[string]interface{} `json:"fields,omitempty" xml:"fields,omitempty"`
	Id     *string                `json:"id,omitempty" xml:"id,omitempty"`
}

func (s GetRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecordResponseBody) SetFields(v map[string]interface{}) *GetRecordResponseBody {
	s.Fields = v
	return s
}

func (s *GetRecordResponseBody) SetId(v string) *GetRecordResponseBody {
	s.Id = &v
	return s
}

type GetRecordResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRecordResponse) GoString() string {
	return s.String()
}

func (s *GetRecordResponse) SetHeaders(v map[string]*string) *GetRecordResponse {
	s.Headers = v
	return s
}

func (s *GetRecordResponse) SetStatusCode(v int32) *GetRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecordResponse) SetBody(v *GetRecordResponseBody) *GetRecordResponse {
	s.Body = v
	return s
}

type GetRecordsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetRecordsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetRecordsHeaders) GoString() string {
	return s.String()
}

func (s *GetRecordsHeaders) SetCommonHeaders(v map[string]*string) *GetRecordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRecordsHeaders) SetXAcsDingtalkAccessToken(v string) *GetRecordsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetRecordsRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s GetRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRecordsRequest) GoString() string {
	return s.String()
}

func (s *GetRecordsRequest) SetMaxResults(v int32) *GetRecordsRequest {
	s.MaxResults = &v
	return s
}

func (s *GetRecordsRequest) SetNextToken(v string) *GetRecordsRequest {
	s.NextToken = &v
	return s
}

type GetRecordsResponseBody struct {
	HasMore   *bool                            `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextToken *string                          `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Records   []*GetRecordsResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
}

func (s GetRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecordsResponseBody) SetHasMore(v bool) *GetRecordsResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetRecordsResponseBody) SetNextToken(v string) *GetRecordsResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetRecordsResponseBody) SetRecords(v []*GetRecordsResponseBodyRecords) *GetRecordsResponseBody {
	s.Records = v
	return s
}

type GetRecordsResponseBodyRecords struct {
	Fields map[string]interface{} `json:"fields,omitempty" xml:"fields,omitempty"`
	Id     *string                `json:"id,omitempty" xml:"id,omitempty"`
}

func (s GetRecordsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s GetRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *GetRecordsResponseBodyRecords) SetFields(v map[string]interface{}) *GetRecordsResponseBodyRecords {
	s.Fields = v
	return s
}

func (s *GetRecordsResponseBodyRecords) SetId(v string) *GetRecordsResponseBodyRecords {
	s.Id = &v
	return s
}

type GetRecordsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRecordsResponse) GoString() string {
	return s.String()
}

func (s *GetRecordsResponse) SetHeaders(v map[string]*string) *GetRecordsResponse {
	s.Headers = v
	return s
}

func (s *GetRecordsResponse) SetStatusCode(v int32) *GetRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecordsResponse) SetBody(v *GetRecordsResponseBody) *GetRecordsResponse {
	s.Body = v
	return s
}

type GetSheetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSheetHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSheetHeaders) GoString() string {
	return s.String()
}

func (s *GetSheetHeaders) SetCommonHeaders(v map[string]*string) *GetSheetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSheetHeaders) SetXAcsDingtalkAccessToken(v string) *GetSheetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSheetResponseBody struct {
	Fields []*GetSheetResponseBodyFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
}

func (s GetSheetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSheetResponseBody) GoString() string {
	return s.String()
}

func (s *GetSheetResponseBody) SetFields(v []*GetSheetResponseBodyFields) *GetSheetResponseBody {
	s.Fields = v
	return s
}

type GetSheetResponseBodyFields struct {
	Id       *string                `json:"id,omitempty" xml:"id,omitempty"`
	Name     *string                `json:"name,omitempty" xml:"name,omitempty"`
	Property map[string]interface{} `json:"property,omitempty" xml:"property,omitempty"`
	Type     *string                `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetSheetResponseBodyFields) String() string {
	return tea.Prettify(s)
}

func (s GetSheetResponseBodyFields) GoString() string {
	return s.String()
}

func (s *GetSheetResponseBodyFields) SetId(v string) *GetSheetResponseBodyFields {
	s.Id = &v
	return s
}

func (s *GetSheetResponseBodyFields) SetName(v string) *GetSheetResponseBodyFields {
	s.Name = &v
	return s
}

func (s *GetSheetResponseBodyFields) SetProperty(v map[string]interface{}) *GetSheetResponseBodyFields {
	s.Property = v
	return s
}

func (s *GetSheetResponseBodyFields) SetType(v string) *GetSheetResponseBodyFields {
	s.Type = &v
	return s
}

type GetSheetResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSheetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSheetResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSheetResponse) GoString() string {
	return s.String()
}

func (s *GetSheetResponse) SetHeaders(v map[string]*string) *GetSheetResponse {
	s.Headers = v
	return s
}

func (s *GetSheetResponse) SetStatusCode(v int32) *GetSheetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSheetResponse) SetBody(v *GetSheetResponseBody) *GetSheetResponse {
	s.Body = v
	return s
}

type InsertRecordsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InsertRecordsHeaders) String() string {
	return tea.Prettify(s)
}

func (s InsertRecordsHeaders) GoString() string {
	return s.String()
}

func (s *InsertRecordsHeaders) SetCommonHeaders(v map[string]*string) *InsertRecordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsertRecordsHeaders) SetXAcsDingtalkAccessToken(v string) *InsertRecordsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InsertRecordsRequest struct {
	Records []*InsertRecordsRequestRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
}

func (s InsertRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertRecordsRequest) GoString() string {
	return s.String()
}

func (s *InsertRecordsRequest) SetRecords(v []*InsertRecordsRequestRecords) *InsertRecordsRequest {
	s.Records = v
	return s
}

type InsertRecordsRequestRecords struct {
	Fields map[string]interface{} `json:"fields,omitempty" xml:"fields,omitempty"`
}

func (s InsertRecordsRequestRecords) String() string {
	return tea.Prettify(s)
}

func (s InsertRecordsRequestRecords) GoString() string {
	return s.String()
}

func (s *InsertRecordsRequestRecords) SetFields(v map[string]interface{}) *InsertRecordsRequestRecords {
	s.Fields = v
	return s
}

type InsertRecordsResponseBody struct {
	Value []*InsertRecordsResponseBodyValue `json:"value,omitempty" xml:"value,omitempty" type:"Repeated"`
}

func (s InsertRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *InsertRecordsResponseBody) SetValue(v []*InsertRecordsResponseBodyValue) *InsertRecordsResponseBody {
	s.Value = v
	return s
}

type InsertRecordsResponseBodyValue struct {
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s InsertRecordsResponseBodyValue) String() string {
	return tea.Prettify(s)
}

func (s InsertRecordsResponseBodyValue) GoString() string {
	return s.String()
}

func (s *InsertRecordsResponseBodyValue) SetId(v string) *InsertRecordsResponseBodyValue {
	s.Id = &v
	return s
}

type InsertRecordsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertRecordsResponse) GoString() string {
	return s.String()
}

func (s *InsertRecordsResponse) SetHeaders(v map[string]*string) *InsertRecordsResponse {
	s.Headers = v
	return s
}

func (s *InsertRecordsResponse) SetStatusCode(v int32) *InsertRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertRecordsResponse) SetBody(v *InsertRecordsResponseBody) *InsertRecordsResponse {
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

func (client *Client) CreateSheetWithOptions(baseId *string, name *string, request *CreateSheetRequest, headers *CreateSheetHeaders, runtime *util.RuntimeOptions) (_result *CreateSheetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Fields)) {
		body["fields"] = request.Fields
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
		Action:      tea.String("CreateSheet"),
		Version:     tea.String("notable_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/notable/bases/" + tea.StringValue(baseId) + "/sheets"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSheetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSheet(baseId *string, name *string, request *CreateSheetRequest) (_result *CreateSheetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateSheetHeaders{}
	_result = &CreateSheetResponse{}
	_body, _err := client.CreateSheetWithOptions(baseId, name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSheetWithOptions(baseId *string, sheetIdOrName *string, headers *DeleteSheetHeaders, runtime *util.RuntimeOptions) (_result *DeleteSheetResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("DeleteSheet"),
		Version:     tea.String("notable_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/notable/bases/" + tea.StringValue(baseId) + "/sheets/" + tea.StringValue(sheetIdOrName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSheetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSheet(baseId *string, sheetIdOrName *string) (_result *DeleteSheetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteSheetHeaders{}
	_result = &DeleteSheetResponse{}
	_body, _err := client.DeleteSheetWithOptions(baseId, sheetIdOrName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAllSheetsWithOptions(baseId *string, headers *GetAllSheetsHeaders, runtime *util.RuntimeOptions) (_result *GetAllSheetsResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("GetAllSheets"),
		Version:     tea.String("notable_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/notable/bases/" + tea.StringValue(baseId) + "/sheets"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAllSheetsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAllSheets(baseId *string) (_result *GetAllSheetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAllSheetsHeaders{}
	_result = &GetAllSheetsResponse{}
	_body, _err := client.GetAllSheetsWithOptions(baseId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRecordWithOptions(baseId *string, sheetIdOrName *string, recordId *string, headers *GetRecordHeaders, runtime *util.RuntimeOptions) (_result *GetRecordResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("GetRecord"),
		Version:     tea.String("notable_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/notable/bases/" + tea.StringValue(baseId) + "/sheets/" + tea.StringValue(sheetIdOrName) + "/records/" + tea.StringValue(recordId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRecordResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRecord(baseId *string, sheetIdOrName *string, recordId *string) (_result *GetRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetRecordHeaders{}
	_result = &GetRecordResponse{}
	_body, _err := client.GetRecordWithOptions(baseId, sheetIdOrName, recordId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRecordsWithOptions(baseId *string, sheetIdOrName *string, request *GetRecordsRequest, headers *GetRecordsHeaders, runtime *util.RuntimeOptions) (_result *GetRecordsResponse, _err error) {
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
		Action:      tea.String("GetRecords"),
		Version:     tea.String("notable_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/notable/bases/" + tea.StringValue(baseId) + "/sheets/" + tea.StringValue(sheetIdOrName) + "/records"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRecordsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRecords(baseId *string, sheetIdOrName *string, request *GetRecordsRequest) (_result *GetRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetRecordsHeaders{}
	_result = &GetRecordsResponse{}
	_body, _err := client.GetRecordsWithOptions(baseId, sheetIdOrName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSheetWithOptions(baseId *string, sheetIdOrName *string, headers *GetSheetHeaders, runtime *util.RuntimeOptions) (_result *GetSheetResponse, _err error) {
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
	params := &openapi.Params{
		Action:      tea.String("GetSheet"),
		Version:     tea.String("notable_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/notable/bases/" + tea.StringValue(baseId) + "/sheets/" + tea.StringValue(sheetIdOrName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSheetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSheet(baseId *string, sheetIdOrName *string) (_result *GetSheetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSheetHeaders{}
	_result = &GetSheetResponse{}
	_body, _err := client.GetSheetWithOptions(baseId, sheetIdOrName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertRecordsWithOptions(baseId *string, sheetIdOrName *string, request *InsertRecordsRequest, headers *InsertRecordsHeaders, runtime *util.RuntimeOptions) (_result *InsertRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Records)) {
		body["records"] = request.Records
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
		Action:      tea.String("InsertRecords"),
		Version:     tea.String("notable_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/notable/bases/" + tea.StringValue(baseId) + "/sheets/" + tea.StringValue(sheetIdOrName) + "/records"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &InsertRecordsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertRecords(baseId *string, sheetIdOrName *string, request *InsertRecordsRequest) (_result *InsertRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InsertRecordsHeaders{}
	_result = &InsertRecordsResponse{}
	_body, _err := client.InsertRecordsWithOptions(baseId, sheetIdOrName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
