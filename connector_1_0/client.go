// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package connector_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type SyncDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SyncDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s SyncDataHeaders) GoString() string {
	return s.String()
}

func (s *SyncDataHeaders) SetCommonHeaders(v map[string]*string) *SyncDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncDataHeaders) SetXAcsDingtalkAccessToken(v string) *SyncDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SyncDataRequest struct {
	TriggerDataList []*SyncDataRequestTriggerDataList `json:"triggerDataList,omitempty" xml:"triggerDataList,omitempty" type:"Repeated"`
}

func (s SyncDataRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncDataRequest) GoString() string {
	return s.String()
}

func (s *SyncDataRequest) SetTriggerDataList(v []*SyncDataRequestTriggerDataList) *SyncDataRequest {
	s.TriggerDataList = v
	return s
}

type SyncDataRequestTriggerDataList struct {
	TriggerId       *string `json:"triggerId,omitempty" xml:"triggerId,omitempty"`
	CustomTriggerId *string `json:"customTriggerId,omitempty" xml:"customTriggerId,omitempty"`
	JsonData        *string `json:"jsonData,omitempty" xml:"jsonData,omitempty"`
	DataGmtCreate   *int64  `json:"dataGmtCreate,omitempty" xml:"dataGmtCreate,omitempty"`
	DataGmtModified *int64  `json:"dataGmtModified,omitempty" xml:"dataGmtModified,omitempty"`
	Action          *string `json:"action,omitempty" xml:"action,omitempty"`
}

func (s SyncDataRequestTriggerDataList) String() string {
	return tea.Prettify(s)
}

func (s SyncDataRequestTriggerDataList) GoString() string {
	return s.String()
}

func (s *SyncDataRequestTriggerDataList) SetTriggerId(v string) *SyncDataRequestTriggerDataList {
	s.TriggerId = &v
	return s
}

func (s *SyncDataRequestTriggerDataList) SetCustomTriggerId(v string) *SyncDataRequestTriggerDataList {
	s.CustomTriggerId = &v
	return s
}

func (s *SyncDataRequestTriggerDataList) SetJsonData(v string) *SyncDataRequestTriggerDataList {
	s.JsonData = &v
	return s
}

func (s *SyncDataRequestTriggerDataList) SetDataGmtCreate(v int64) *SyncDataRequestTriggerDataList {
	s.DataGmtCreate = &v
	return s
}

func (s *SyncDataRequestTriggerDataList) SetDataGmtModified(v int64) *SyncDataRequestTriggerDataList {
	s.DataGmtModified = &v
	return s
}

func (s *SyncDataRequestTriggerDataList) SetAction(v string) *SyncDataRequestTriggerDataList {
	s.Action = &v
	return s
}

type SyncDataResponseBody struct {
	// resultList
	List []*SyncDataResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s SyncDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncDataResponseBody) GoString() string {
	return s.String()
}

func (s *SyncDataResponseBody) SetList(v []*SyncDataResponseBodyList) *SyncDataResponseBody {
	s.List = v
	return s
}

type SyncDataResponseBodyList struct {
	TriggerId     *string `json:"triggerId,omitempty" xml:"triggerId,omitempty"`
	BizPrimaryKey *string `json:"bizPrimaryKey,omitempty" xml:"bizPrimaryKey,omitempty"`
	Success       *bool   `json:"success,omitempty" xml:"success,omitempty"`
	SubErrCode    *string `json:"subErrCode,omitempty" xml:"subErrCode,omitempty"`
	SubErrMsg     *string `json:"subErrMsg,omitempty" xml:"subErrMsg,omitempty"`
}

func (s SyncDataResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s SyncDataResponseBodyList) GoString() string {
	return s.String()
}

func (s *SyncDataResponseBodyList) SetTriggerId(v string) *SyncDataResponseBodyList {
	s.TriggerId = &v
	return s
}

func (s *SyncDataResponseBodyList) SetBizPrimaryKey(v string) *SyncDataResponseBodyList {
	s.BizPrimaryKey = &v
	return s
}

func (s *SyncDataResponseBodyList) SetSuccess(v bool) *SyncDataResponseBodyList {
	s.Success = &v
	return s
}

func (s *SyncDataResponseBodyList) SetSubErrCode(v string) *SyncDataResponseBodyList {
	s.SubErrCode = &v
	return s
}

func (s *SyncDataResponseBodyList) SetSubErrMsg(v string) *SyncDataResponseBodyList {
	s.SubErrMsg = &v
	return s
}

type SyncDataResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SyncDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncDataResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncDataResponse) GoString() string {
	return s.String()
}

func (s *SyncDataResponse) SetHeaders(v map[string]*string) *SyncDataResponse {
	s.Headers = v
	return s
}

func (s *SyncDataResponse) SetBody(v *SyncDataResponseBody) *SyncDataResponse {
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

func (client *Client) SyncData(request *SyncDataRequest) (_result *SyncDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SyncDataHeaders{}
	_result = &SyncDataResponse{}
	_body, _err := client.SyncDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncDataWithOptions(request *SyncDataRequest, headers *SyncDataHeaders, runtime *util.RuntimeOptions) (_result *SyncDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TriggerDataList)) {
		body["triggerDataList"] = request.TriggerDataList
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &SyncDataResponse{}
	_body, _err := client.DoROARequest(tea.String("SyncData"), tea.String("connector_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/connector/triggers/data/sync"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
