// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package wiki_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type WikiWordsDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s WikiWordsDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s WikiWordsDetailHeaders) GoString() string {
	return s.String()
}

func (s *WikiWordsDetailHeaders) SetCommonHeaders(v map[string]*string) *WikiWordsDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *WikiWordsDetailHeaders) SetXAcsDingtalkAccessToken(v string) *WikiWordsDetailHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type WikiWordsDetailRequest struct {
	WordName *string `json:"wordName,omitempty" xml:"wordName,omitempty"`
}

func (s WikiWordsDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s WikiWordsDetailRequest) GoString() string {
	return s.String()
}

func (s *WikiWordsDetailRequest) SetWordName(v string) *WikiWordsDetailRequest {
	s.WordName = &v
	return s
}

type WikiWordsDetailResponseBody struct {
	Data    []*WikiWordsDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	ErrMsg  *string                            `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                              `json:"success,omitempty" xml:"success,omitempty"`
}

func (s WikiWordsDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WikiWordsDetailResponseBody) GoString() string {
	return s.String()
}

func (s *WikiWordsDetailResponseBody) SetData(v []*WikiWordsDetailResponseBodyData) *WikiWordsDetailResponseBody {
	s.Data = v
	return s
}

func (s *WikiWordsDetailResponseBody) SetErrMsg(v string) *WikiWordsDetailResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *WikiWordsDetailResponseBody) SetSuccess(v bool) *WikiWordsDetailResponseBody {
	s.Success = &v
	return s
}

type WikiWordsDetailResponseBodyData struct {
	AppLink              []*WikiWordsDetailResponseBodyDataAppLink     `json:"appLink,omitempty" xml:"appLink,omitempty" type:"Repeated"`
	ApproveName          *string                                       `json:"approveName,omitempty" xml:"approveName,omitempty"`
	Contacts             []*string                                     `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	CreatorName          *string                                       `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	GmtCreate            *int64                                        `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModify            *int64                                        `json:"gmtModify,omitempty" xml:"gmtModify,omitempty"`
	HighLightWordAlias   []*string                                     `json:"highLightWordAlias,omitempty" xml:"highLightWordAlias,omitempty" type:"Repeated"`
	ImHighLight          *bool                                         `json:"imHighLight,omitempty" xml:"imHighLight,omitempty"`
	OrgName              *string                                       `json:"orgName,omitempty" xml:"orgName,omitempty"`
	RelatedDoc           []*WikiWordsDetailResponseBodyDataRelatedDoc  `json:"relatedDoc,omitempty" xml:"relatedDoc,omitempty" type:"Repeated"`
	RelatedLink          []*WikiWordsDetailResponseBodyDataRelatedLink `json:"relatedLink,omitempty" xml:"relatedLink,omitempty" type:"Repeated"`
	SimHighLight         *bool                                         `json:"simHighLight,omitempty" xml:"simHighLight,omitempty"`
	SimpleWordParaphrase *string                                       `json:"simpleWordParaphrase,omitempty" xml:"simpleWordParaphrase,omitempty"`
	TagsList             []*string                                     `json:"tagsList,omitempty" xml:"tagsList,omitempty" type:"Repeated"`
	UpdaterName          *string                                       `json:"updaterName,omitempty" xml:"updaterName,omitempty"`
	Uuid                 *int64                                        `json:"uuid,omitempty" xml:"uuid,omitempty"`
	WordAlias            []*string                                     `json:"wordAlias,omitempty" xml:"wordAlias,omitempty" type:"Repeated"`
	WordFullName         *string                                       `json:"wordFullName,omitempty" xml:"wordFullName,omitempty"`
	WordName             *string                                       `json:"wordName,omitempty" xml:"wordName,omitempty"`
	WordParaphrase       *string                                       `json:"wordParaphrase,omitempty" xml:"wordParaphrase,omitempty"`
}

func (s WikiWordsDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s WikiWordsDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *WikiWordsDetailResponseBodyData) SetAppLink(v []*WikiWordsDetailResponseBodyDataAppLink) *WikiWordsDetailResponseBodyData {
	s.AppLink = v
	return s
}

func (s *WikiWordsDetailResponseBodyData) SetApproveName(v string) *WikiWordsDetailResponseBodyData {
	s.ApproveName = &v
	return s
}

func (s *WikiWordsDetailResponseBodyData) SetContacts(v []*string) *WikiWordsDetailResponseBodyData {
	s.Contacts = v
	return s
}

func (s *WikiWordsDetailResponseBodyData) SetCreatorName(v string) *WikiWordsDetailResponseBodyData {
	s.CreatorName = &v
	return s
}

func (s *WikiWordsDetailResponseBodyData) SetGmtCreate(v int64) *WikiWordsDetailResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *WikiWordsDetailResponseBodyData) SetGmtModify(v int64) *WikiWordsDetailResponseBodyData {
	s.GmtModify = &v
	return s
}

func (s *WikiWordsDetailResponseBodyData) SetHighLightWordAlias(v []*string) *WikiWordsDetailResponseBodyData {
	s.HighLightWordAlias = v
	return s
}

func (s *WikiWordsDetailResponseBodyData) SetImHighLight(v bool) *WikiWordsDetailResponseBodyData {
	s.ImHighLight = &v
	return s
}

func (s *WikiWordsDetailResponseBodyData) SetOrgName(v string) *WikiWordsDetailResponseBodyData {
	s.OrgName = &v
	return s
}

func (s *WikiWordsDetailResponseBodyData) SetRelatedDoc(v []*WikiWordsDetailResponseBodyDataRelatedDoc) *WikiWordsDetailResponseBodyData {
	s.RelatedDoc = v
	return s
}

func (s *WikiWordsDetailResponseBodyData) SetRelatedLink(v []*WikiWordsDetailResponseBodyDataRelatedLink) *WikiWordsDetailResponseBodyData {
	s.RelatedLink = v
	return s
}

func (s *WikiWordsDetailResponseBodyData) SetSimHighLight(v bool) *WikiWordsDetailResponseBodyData {
	s.SimHighLight = &v
	return s
}

func (s *WikiWordsDetailResponseBodyData) SetSimpleWordParaphrase(v string) *WikiWordsDetailResponseBodyData {
	s.SimpleWordParaphrase = &v
	return s
}

func (s *WikiWordsDetailResponseBodyData) SetTagsList(v []*string) *WikiWordsDetailResponseBodyData {
	s.TagsList = v
	return s
}

func (s *WikiWordsDetailResponseBodyData) SetUpdaterName(v string) *WikiWordsDetailResponseBodyData {
	s.UpdaterName = &v
	return s
}

func (s *WikiWordsDetailResponseBodyData) SetUuid(v int64) *WikiWordsDetailResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *WikiWordsDetailResponseBodyData) SetWordAlias(v []*string) *WikiWordsDetailResponseBodyData {
	s.WordAlias = v
	return s
}

func (s *WikiWordsDetailResponseBodyData) SetWordFullName(v string) *WikiWordsDetailResponseBodyData {
	s.WordFullName = &v
	return s
}

func (s *WikiWordsDetailResponseBodyData) SetWordName(v string) *WikiWordsDetailResponseBodyData {
	s.WordName = &v
	return s
}

func (s *WikiWordsDetailResponseBodyData) SetWordParaphrase(v string) *WikiWordsDetailResponseBodyData {
	s.WordParaphrase = &v
	return s
}

type WikiWordsDetailResponseBodyDataAppLink struct {
	AppId     *int64  `json:"appId,omitempty" xml:"appId,omitempty"`
	AppName   *string `json:"appName,omitempty" xml:"appName,omitempty"`
	IconLink  *string `json:"iconLink,omitempty" xml:"iconLink,omitempty"`
	PcLink    *string `json:"pcLink,omitempty" xml:"pcLink,omitempty"`
	PhoneLink *string `json:"phoneLink,omitempty" xml:"phoneLink,omitempty"`
}

func (s WikiWordsDetailResponseBodyDataAppLink) String() string {
	return tea.Prettify(s)
}

func (s WikiWordsDetailResponseBodyDataAppLink) GoString() string {
	return s.String()
}

func (s *WikiWordsDetailResponseBodyDataAppLink) SetAppId(v int64) *WikiWordsDetailResponseBodyDataAppLink {
	s.AppId = &v
	return s
}

func (s *WikiWordsDetailResponseBodyDataAppLink) SetAppName(v string) *WikiWordsDetailResponseBodyDataAppLink {
	s.AppName = &v
	return s
}

func (s *WikiWordsDetailResponseBodyDataAppLink) SetIconLink(v string) *WikiWordsDetailResponseBodyDataAppLink {
	s.IconLink = &v
	return s
}

func (s *WikiWordsDetailResponseBodyDataAppLink) SetPcLink(v string) *WikiWordsDetailResponseBodyDataAppLink {
	s.PcLink = &v
	return s
}

func (s *WikiWordsDetailResponseBodyDataAppLink) SetPhoneLink(v string) *WikiWordsDetailResponseBodyDataAppLink {
	s.PhoneLink = &v
	return s
}

type WikiWordsDetailResponseBodyDataRelatedDoc struct {
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s WikiWordsDetailResponseBodyDataRelatedDoc) String() string {
	return tea.Prettify(s)
}

func (s WikiWordsDetailResponseBodyDataRelatedDoc) GoString() string {
	return s.String()
}

func (s *WikiWordsDetailResponseBodyDataRelatedDoc) SetLink(v string) *WikiWordsDetailResponseBodyDataRelatedDoc {
	s.Link = &v
	return s
}

func (s *WikiWordsDetailResponseBodyDataRelatedDoc) SetName(v string) *WikiWordsDetailResponseBodyDataRelatedDoc {
	s.Name = &v
	return s
}

func (s *WikiWordsDetailResponseBodyDataRelatedDoc) SetType(v string) *WikiWordsDetailResponseBodyDataRelatedDoc {
	s.Type = &v
	return s
}

type WikiWordsDetailResponseBodyDataRelatedLink struct {
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s WikiWordsDetailResponseBodyDataRelatedLink) String() string {
	return tea.Prettify(s)
}

func (s WikiWordsDetailResponseBodyDataRelatedLink) GoString() string {
	return s.String()
}

func (s *WikiWordsDetailResponseBodyDataRelatedLink) SetLink(v string) *WikiWordsDetailResponseBodyDataRelatedLink {
	s.Link = &v
	return s
}

func (s *WikiWordsDetailResponseBodyDataRelatedLink) SetName(v string) *WikiWordsDetailResponseBodyDataRelatedLink {
	s.Name = &v
	return s
}

func (s *WikiWordsDetailResponseBodyDataRelatedLink) SetType(v string) *WikiWordsDetailResponseBodyDataRelatedLink {
	s.Type = &v
	return s
}

type WikiWordsDetailResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *WikiWordsDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s WikiWordsDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s WikiWordsDetailResponse) GoString() string {
	return s.String()
}

func (s *WikiWordsDetailResponse) SetHeaders(v map[string]*string) *WikiWordsDetailResponse {
	s.Headers = v
	return s
}

func (s *WikiWordsDetailResponse) SetStatusCode(v int32) *WikiWordsDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *WikiWordsDetailResponse) SetBody(v *WikiWordsDetailResponseBody) *WikiWordsDetailResponse {
	s.Body = v
	return s
}

type WikiWordsParseHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s WikiWordsParseHeaders) String() string {
	return tea.Prettify(s)
}

func (s WikiWordsParseHeaders) GoString() string {
	return s.String()
}

func (s *WikiWordsParseHeaders) SetCommonHeaders(v map[string]*string) *WikiWordsParseHeaders {
	s.CommonHeaders = v
	return s
}

func (s *WikiWordsParseHeaders) SetXAcsDingtalkAccessToken(v string) *WikiWordsParseHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type WikiWordsParseRequest struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s WikiWordsParseRequest) String() string {
	return tea.Prettify(s)
}

func (s WikiWordsParseRequest) GoString() string {
	return s.String()
}

func (s *WikiWordsParseRequest) SetContent(v string) *WikiWordsParseRequest {
	s.Content = &v
	return s
}

type WikiWordsParseResponseBody struct {
	Data    []*WikiWordsParseResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	ErrMsg  *string                           `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s WikiWordsParseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WikiWordsParseResponseBody) GoString() string {
	return s.String()
}

func (s *WikiWordsParseResponseBody) SetData(v []*WikiWordsParseResponseBodyData) *WikiWordsParseResponseBody {
	s.Data = v
	return s
}

func (s *WikiWordsParseResponseBody) SetErrMsg(v string) *WikiWordsParseResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *WikiWordsParseResponseBody) SetSuccess(v bool) *WikiWordsParseResponseBody {
	s.Success = &v
	return s
}

type WikiWordsParseResponseBodyData struct {
	EndIndex   *int64  `json:"endIndex,omitempty" xml:"endIndex,omitempty"`
	StartIndex *int64  `json:"startIndex,omitempty" xml:"startIndex,omitempty"`
	WordName   *string `json:"wordName,omitempty" xml:"wordName,omitempty"`
}

func (s WikiWordsParseResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s WikiWordsParseResponseBodyData) GoString() string {
	return s.String()
}

func (s *WikiWordsParseResponseBodyData) SetEndIndex(v int64) *WikiWordsParseResponseBodyData {
	s.EndIndex = &v
	return s
}

func (s *WikiWordsParseResponseBodyData) SetStartIndex(v int64) *WikiWordsParseResponseBodyData {
	s.StartIndex = &v
	return s
}

func (s *WikiWordsParseResponseBodyData) SetWordName(v string) *WikiWordsParseResponseBodyData {
	s.WordName = &v
	return s
}

type WikiWordsParseResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *WikiWordsParseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s WikiWordsParseResponse) String() string {
	return tea.Prettify(s)
}

func (s WikiWordsParseResponse) GoString() string {
	return s.String()
}

func (s *WikiWordsParseResponse) SetHeaders(v map[string]*string) *WikiWordsParseResponse {
	s.Headers = v
	return s
}

func (s *WikiWordsParseResponse) SetStatusCode(v int32) *WikiWordsParseResponse {
	s.StatusCode = &v
	return s
}

func (s *WikiWordsParseResponse) SetBody(v *WikiWordsParseResponseBody) *WikiWordsParseResponse {
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

func (client *Client) WikiWordsDetailWithOptions(request *WikiWordsDetailRequest, headers *WikiWordsDetailHeaders, runtime *util.RuntimeOptions) (_result *WikiWordsDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WordName)) {
		query["wordName"] = request.WordName
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
		Action:      tea.String("WikiWordsDetail"),
		Version:     tea.String("wiki_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/wiki/words/details"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &WikiWordsDetailResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) WikiWordsDetail(request *WikiWordsDetailRequest) (_result *WikiWordsDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &WikiWordsDetailHeaders{}
	_result = &WikiWordsDetailResponse{}
	_body, _err := client.WikiWordsDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) WikiWordsParseWithOptions(request *WikiWordsParseRequest, headers *WikiWordsParseHeaders, runtime *util.RuntimeOptions) (_result *WikiWordsParseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
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
		Action:      tea.String("WikiWordsParse"),
		Version:     tea.String("wiki_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/wiki/words/parse"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &WikiWordsParseResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) WikiWordsParse(request *WikiWordsParseRequest) (_result *WikiWordsParseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &WikiWordsParseHeaders{}
	_result = &WikiWordsParseResponse{}
	_body, _err := client.WikiWordsParseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
