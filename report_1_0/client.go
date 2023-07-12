// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package report_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateTemplatesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateTemplatesHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplatesHeaders) GoString() string {
	return s.String()
}

func (s *CreateTemplatesHeaders) SetCommonHeaders(v map[string]*string) *CreateTemplatesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTemplatesHeaders) SetXAcsDingtalkAccessToken(v string) *CreateTemplatesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateTemplatesRequest struct {
	AllowAddReceivers           *bool                           `json:"allowAddReceivers,omitempty" xml:"allowAddReceivers,omitempty"`
	AllowEdit                   *bool                           `json:"allowEdit,omitempty" xml:"allowEdit,omitempty"`
	AllowGetLocation            *bool                           `json:"allowGetLocation,omitempty" xml:"allowGetLocation,omitempty"`
	AuthDeptIds                 []*string                       `json:"authDeptIds,omitempty" xml:"authDeptIds,omitempty" type:"Repeated"`
	AuthUserIds                 []*string                       `json:"authUserIds,omitempty" xml:"authUserIds,omitempty" type:"Repeated"`
	Creator                     *string                         `json:"creator,omitempty" xml:"creator,omitempty"`
	DefaultReceivedCids         []*string                       `json:"defaultReceivedCids,omitempty" xml:"defaultReceivedCids,omitempty" type:"Repeated"`
	DefaultReceivedMasterLevels []*string                       `json:"defaultReceivedMasterLevels,omitempty" xml:"defaultReceivedMasterLevels,omitempty" type:"Repeated"`
	DefaultReceivers            []*string                       `json:"defaultReceivers,omitempty" xml:"defaultReceivers,omitempty" type:"Repeated"`
	Fields                      []*CreateTemplatesRequestFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	Logo                        *string                         `json:"logo,omitempty" xml:"logo,omitempty"`
	MaxWordCount                *int32                          `json:"maxWordCount,omitempty" xml:"maxWordCount,omitempty"`
	MinWordCount                *int32                          `json:"minWordCount,omitempty" xml:"minWordCount,omitempty"`
	Name                        *string                         `json:"name,omitempty" xml:"name,omitempty"`
	TemplateManagers            []*string                       `json:"templateManagers,omitempty" xml:"templateManagers,omitempty" type:"Repeated"`
}

func (s CreateTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplatesRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplatesRequest) SetAllowAddReceivers(v bool) *CreateTemplatesRequest {
	s.AllowAddReceivers = &v
	return s
}

func (s *CreateTemplatesRequest) SetAllowEdit(v bool) *CreateTemplatesRequest {
	s.AllowEdit = &v
	return s
}

func (s *CreateTemplatesRequest) SetAllowGetLocation(v bool) *CreateTemplatesRequest {
	s.AllowGetLocation = &v
	return s
}

func (s *CreateTemplatesRequest) SetAuthDeptIds(v []*string) *CreateTemplatesRequest {
	s.AuthDeptIds = v
	return s
}

func (s *CreateTemplatesRequest) SetAuthUserIds(v []*string) *CreateTemplatesRequest {
	s.AuthUserIds = v
	return s
}

func (s *CreateTemplatesRequest) SetCreator(v string) *CreateTemplatesRequest {
	s.Creator = &v
	return s
}

func (s *CreateTemplatesRequest) SetDefaultReceivedCids(v []*string) *CreateTemplatesRequest {
	s.DefaultReceivedCids = v
	return s
}

func (s *CreateTemplatesRequest) SetDefaultReceivedMasterLevels(v []*string) *CreateTemplatesRequest {
	s.DefaultReceivedMasterLevels = v
	return s
}

func (s *CreateTemplatesRequest) SetDefaultReceivers(v []*string) *CreateTemplatesRequest {
	s.DefaultReceivers = v
	return s
}

func (s *CreateTemplatesRequest) SetFields(v []*CreateTemplatesRequestFields) *CreateTemplatesRequest {
	s.Fields = v
	return s
}

func (s *CreateTemplatesRequest) SetLogo(v string) *CreateTemplatesRequest {
	s.Logo = &v
	return s
}

func (s *CreateTemplatesRequest) SetMaxWordCount(v int32) *CreateTemplatesRequest {
	s.MaxWordCount = &v
	return s
}

func (s *CreateTemplatesRequest) SetMinWordCount(v int32) *CreateTemplatesRequest {
	s.MinWordCount = &v
	return s
}

func (s *CreateTemplatesRequest) SetName(v string) *CreateTemplatesRequest {
	s.Name = &v
	return s
}

func (s *CreateTemplatesRequest) SetTemplateManagers(v []*string) *CreateTemplatesRequest {
	s.TemplateManagers = v
	return s
}

type CreateTemplatesRequestFields struct {
	DataType  *int32                                 `json:"dataType,omitempty" xml:"dataType,omitempty"`
	DataValue *CreateTemplatesRequestFieldsDataValue `json:"dataValue,omitempty" xml:"dataValue,omitempty" type:"Struct"`
	FieldName *string                                `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	Need      *bool                                  `json:"need,omitempty" xml:"need,omitempty"`
	Order     *int32                                 `json:"order,omitempty" xml:"order,omitempty"`
	Sort      *int32                                 `json:"sort,omitempty" xml:"sort,omitempty"`
}

func (s CreateTemplatesRequestFields) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplatesRequestFields) GoString() string {
	return s.String()
}

func (s *CreateTemplatesRequestFields) SetDataType(v int32) *CreateTemplatesRequestFields {
	s.DataType = &v
	return s
}

func (s *CreateTemplatesRequestFields) SetDataValue(v *CreateTemplatesRequestFieldsDataValue) *CreateTemplatesRequestFields {
	s.DataValue = v
	return s
}

func (s *CreateTemplatesRequestFields) SetFieldName(v string) *CreateTemplatesRequestFields {
	s.FieldName = &v
	return s
}

func (s *CreateTemplatesRequestFields) SetNeed(v bool) *CreateTemplatesRequestFields {
	s.Need = &v
	return s
}

func (s *CreateTemplatesRequestFields) SetOrder(v int32) *CreateTemplatesRequestFields {
	s.Order = &v
	return s
}

func (s *CreateTemplatesRequestFields) SetSort(v int32) *CreateTemplatesRequestFields {
	s.Sort = &v
	return s
}

type CreateTemplatesRequestFieldsDataValue struct {
	OpenInfo    *CreateTemplatesRequestFieldsDataValueOpenInfo `json:"openInfo,omitempty" xml:"openInfo,omitempty" type:"Struct"`
	Options     []*string                                      `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	Placeholder *string                                        `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
}

func (s CreateTemplatesRequestFieldsDataValue) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplatesRequestFieldsDataValue) GoString() string {
	return s.String()
}

func (s *CreateTemplatesRequestFieldsDataValue) SetOpenInfo(v *CreateTemplatesRequestFieldsDataValueOpenInfo) *CreateTemplatesRequestFieldsDataValue {
	s.OpenInfo = v
	return s
}

func (s *CreateTemplatesRequestFieldsDataValue) SetOptions(v []*string) *CreateTemplatesRequestFieldsDataValue {
	s.Options = v
	return s
}

func (s *CreateTemplatesRequestFieldsDataValue) SetPlaceholder(v string) *CreateTemplatesRequestFieldsDataValue {
	s.Placeholder = &v
	return s
}

type CreateTemplatesRequestFieldsDataValueOpenInfo struct {
	Attribute map[string]*string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	OpenId    *string            `json:"openId,omitempty" xml:"openId,omitempty"`
}

func (s CreateTemplatesRequestFieldsDataValueOpenInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplatesRequestFieldsDataValueOpenInfo) GoString() string {
	return s.String()
}

func (s *CreateTemplatesRequestFieldsDataValueOpenInfo) SetAttribute(v map[string]*string) *CreateTemplatesRequestFieldsDataValueOpenInfo {
	s.Attribute = v
	return s
}

func (s *CreateTemplatesRequestFieldsDataValueOpenInfo) SetOpenId(v string) *CreateTemplatesRequestFieldsDataValueOpenInfo {
	s.OpenId = &v
	return s
}

type CreateTemplatesResponseBody struct {
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s CreateTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplatesResponseBody) SetTemplateId(v string) *CreateTemplatesResponseBody {
	s.TemplateId = &v
	return s
}

type CreateTemplatesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplatesResponse) GoString() string {
	return s.String()
}

func (s *CreateTemplatesResponse) SetHeaders(v map[string]*string) *CreateTemplatesResponse {
	s.Headers = v
	return s
}

func (s *CreateTemplatesResponse) SetStatusCode(v int32) *CreateTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTemplatesResponse) SetBody(v *CreateTemplatesResponseBody) *CreateTemplatesResponse {
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

func (client *Client) CreateTemplatesWithOptions(request *CreateTemplatesRequest, headers *CreateTemplatesHeaders, runtime *util.RuntimeOptions) (_result *CreateTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowAddReceivers)) {
		body["allowAddReceivers"] = request.AllowAddReceivers
	}

	if !tea.BoolValue(util.IsUnset(request.AllowEdit)) {
		body["allowEdit"] = request.AllowEdit
	}

	if !tea.BoolValue(util.IsUnset(request.AllowGetLocation)) {
		body["allowGetLocation"] = request.AllowGetLocation
	}

	if !tea.BoolValue(util.IsUnset(request.AuthDeptIds)) {
		body["authDeptIds"] = request.AuthDeptIds
	}

	if !tea.BoolValue(util.IsUnset(request.AuthUserIds)) {
		body["authUserIds"] = request.AuthUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.Creator)) {
		body["creator"] = request.Creator
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultReceivedCids)) {
		body["defaultReceivedCids"] = request.DefaultReceivedCids
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultReceivedMasterLevels)) {
		body["defaultReceivedMasterLevels"] = request.DefaultReceivedMasterLevels
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultReceivers)) {
		body["defaultReceivers"] = request.DefaultReceivers
	}

	if !tea.BoolValue(util.IsUnset(request.Fields)) {
		body["fields"] = request.Fields
	}

	if !tea.BoolValue(util.IsUnset(request.Logo)) {
		body["logo"] = request.Logo
	}

	if !tea.BoolValue(util.IsUnset(request.MaxWordCount)) {
		body["maxWordCount"] = request.MaxWordCount
	}

	if !tea.BoolValue(util.IsUnset(request.MinWordCount)) {
		body["minWordCount"] = request.MinWordCount
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateManagers)) {
		body["templateManagers"] = request.TemplateManagers
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
		Action:      tea.String("CreateTemplates"),
		Version:     tea.String("report_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/report/templates"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTemplatesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTemplates(request *CreateTemplatesRequest) (_result *CreateTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateTemplatesHeaders{}
	_result = &CreateTemplatesResponse{}
	_body, _err := client.CreateTemplatesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
