// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package crm_2_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetRelationUkSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetRelationUkSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetRelationUkSettingHeaders) GoString() string {
	return s.String()
}

func (s *GetRelationUkSettingHeaders) SetCommonHeaders(v map[string]*string) *GetRelationUkSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRelationUkSettingHeaders) SetXAcsDingtalkAccessToken(v string) *GetRelationUkSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetRelationUkSettingRequest struct {
	// 关系类型。
	RelationType *string `json:"relationType,omitempty" xml:"relationType,omitempty"`
}

func (s GetRelationUkSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRelationUkSettingRequest) GoString() string {
	return s.String()
}

func (s *GetRelationUkSettingRequest) SetRelationType(v string) *GetRelationUkSettingRequest {
	s.RelationType = &v
	return s
}

type GetRelationUkSettingResponseBody struct {
	Result *GetRelationUkSettingResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetRelationUkSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRelationUkSettingResponseBody) GoString() string {
	return s.String()
}

func (s *GetRelationUkSettingResponseBody) SetResult(v *GetRelationUkSettingResponseBodyResult) *GetRelationUkSettingResponseBody {
	s.Result = v
	return s
}

type GetRelationUkSettingResponseBodyResult struct {
	FormUkSettings []*GetRelationUkSettingResponseBodyResultFormUkSettings `json:"formUkSettings,omitempty" xml:"formUkSettings,omitempty" type:"Repeated"`
	// 查重列表表头字段id列表。
	HeaderFieldIds []*string `json:"headerFieldIds,omitempty" xml:"headerFieldIds,omitempty" type:"Repeated"`
}

func (s GetRelationUkSettingResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetRelationUkSettingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRelationUkSettingResponseBodyResult) SetFormUkSettings(v []*GetRelationUkSettingResponseBodyResultFormUkSettings) *GetRelationUkSettingResponseBodyResult {
	s.FormUkSettings = v
	return s
}

func (s *GetRelationUkSettingResponseBodyResult) SetHeaderFieldIds(v []*string) *GetRelationUkSettingResponseBodyResult {
	s.HeaderFieldIds = v
	return s
}

type GetRelationUkSettingResponseBodyResultFormUkSettings struct {
	FieldList []*GetRelationUkSettingResponseBodyResultFormUkSettingsFieldList `json:"fieldList,omitempty" xml:"fieldList,omitempty" type:"Repeated"`
}

func (s GetRelationUkSettingResponseBodyResultFormUkSettings) String() string {
	return tea.Prettify(s)
}

func (s GetRelationUkSettingResponseBodyResultFormUkSettings) GoString() string {
	return s.String()
}

func (s *GetRelationUkSettingResponseBodyResultFormUkSettings) SetFieldList(v []*GetRelationUkSettingResponseBodyResultFormUkSettingsFieldList) *GetRelationUkSettingResponseBodyResultFormUkSettings {
	s.FieldList = v
	return s
}

type GetRelationUkSettingResponseBodyResultFormUkSettingsFieldList struct {
	// 查重字段的bizAlias。
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// 查重字段的字段id。
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
}

func (s GetRelationUkSettingResponseBodyResultFormUkSettingsFieldList) String() string {
	return tea.Prettify(s)
}

func (s GetRelationUkSettingResponseBodyResultFormUkSettingsFieldList) GoString() string {
	return s.String()
}

func (s *GetRelationUkSettingResponseBodyResultFormUkSettingsFieldList) SetBizAlias(v string) *GetRelationUkSettingResponseBodyResultFormUkSettingsFieldList {
	s.BizAlias = &v
	return s
}

func (s *GetRelationUkSettingResponseBodyResultFormUkSettingsFieldList) SetFieldId(v string) *GetRelationUkSettingResponseBodyResultFormUkSettingsFieldList {
	s.FieldId = &v
	return s
}

type GetRelationUkSettingResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRelationUkSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRelationUkSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRelationUkSettingResponse) GoString() string {
	return s.String()
}

func (s *GetRelationUkSettingResponse) SetHeaders(v map[string]*string) *GetRelationUkSettingResponse {
	s.Headers = v
	return s
}

func (s *GetRelationUkSettingResponse) SetBody(v *GetRelationUkSettingResponseBody) *GetRelationUkSettingResponse {
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

func (client *Client) GetRelationUkSetting(request *GetRelationUkSettingRequest) (_result *GetRelationUkSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetRelationUkSettingHeaders{}
	_result = &GetRelationUkSettingResponse{}
	_body, _err := client.GetRelationUkSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRelationUkSettingWithOptions(request *GetRelationUkSettingRequest, headers *GetRelationUkSettingHeaders, runtime *util.RuntimeOptions) (_result *GetRelationUkSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RelationType)) {
		query["relationType"] = request.RelationType
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
	_result = &GetRelationUkSettingResponse{}
	_body, _err := client.DoROARequest(tea.String("GetRelationUkSetting"), tea.String("crm_2.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v2.0/crm/relationUkSettings"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
