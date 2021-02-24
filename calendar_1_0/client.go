// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package calendar_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DeleteEventHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteEventHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventHeaders) GoString() string {
	return s.String()
}

func (s *DeleteEventHeaders) SetCommonHeaders(v map[string]*string) *DeleteEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteEventHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteEventHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteEventResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEventResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventResponse) SetHeaders(v map[string]*string) *DeleteEventResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventResponse) SetBody(v map[string]interface{}) *DeleteEventResponse {
	s.Body = v
	return s
}

type RespondEventHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RespondEventHeaders) String() string {
	return tea.Prettify(s)
}

func (s RespondEventHeaders) GoString() string {
	return s.String()
}

func (s *RespondEventHeaders) SetCommonHeaders(v map[string]*string) *RespondEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RespondEventHeaders) SetXAcsDingtalkAccessToken(v string) *RespondEventHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RespondEventResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RespondEventResponse) String() string {
	return tea.Prettify(s)
}

func (s RespondEventResponse) GoString() string {
	return s.String()
}

func (s *RespondEventResponse) SetHeaders(v map[string]*string) *RespondEventResponse {
	s.Headers = v
	return s
}

func (s *RespondEventResponse) SetBody(v map[string]interface{}) *RespondEventResponse {
	s.Body = v
	return s
}

type ListEventsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListEventsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListEventsHeaders) GoString() string {
	return s.String()
}

func (s *ListEventsHeaders) SetCommonHeaders(v map[string]*string) *ListEventsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListEventsHeaders) SetXAcsDingtalkAccessToken(v string) *ListEventsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListEventsResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEventsResponse) GoString() string {
	return s.String()
}

func (s *ListEventsResponse) SetHeaders(v map[string]*string) *ListEventsResponse {
	s.Headers = v
	return s
}

func (s *ListEventsResponse) SetBody(v map[string]interface{}) *ListEventsResponse {
	s.Body = v
	return s
}

type RemoveAttendeeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RemoveAttendeeHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveAttendeeHeaders) GoString() string {
	return s.String()
}

func (s *RemoveAttendeeHeaders) SetCommonHeaders(v map[string]*string) *RemoveAttendeeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveAttendeeHeaders) SetXAcsDingtalkAccessToken(v string) *RemoveAttendeeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RemoveAttendeeResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveAttendeeResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveAttendeeResponse) GoString() string {
	return s.String()
}

func (s *RemoveAttendeeResponse) SetHeaders(v map[string]*string) *RemoveAttendeeResponse {
	s.Headers = v
	return s
}

func (s *RemoveAttendeeResponse) SetBody(v map[string]interface{}) *RemoveAttendeeResponse {
	s.Body = v
	return s
}

type AddAttendeeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddAttendeeHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddAttendeeHeaders) GoString() string {
	return s.String()
}

func (s *AddAttendeeHeaders) SetCommonHeaders(v map[string]*string) *AddAttendeeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddAttendeeHeaders) SetXAcsDingtalkAccessToken(v string) *AddAttendeeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddAttendeeResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddAttendeeResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAttendeeResponse) GoString() string {
	return s.String()
}

func (s *AddAttendeeResponse) SetHeaders(v map[string]*string) *AddAttendeeResponse {
	s.Headers = v
	return s
}

func (s *AddAttendeeResponse) SetBody(v map[string]interface{}) *AddAttendeeResponse {
	s.Body = v
	return s
}

type GetEventHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetEventHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetEventHeaders) GoString() string {
	return s.String()
}

func (s *GetEventHeaders) SetCommonHeaders(v map[string]*string) *GetEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetEventHeaders) SetXAcsDingtalkAccessToken(v string) *GetEventHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetEventResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEventResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEventResponse) GoString() string {
	return s.String()
}

func (s *GetEventResponse) SetHeaders(v map[string]*string) *GetEventResponse {
	s.Headers = v
	return s
}

func (s *GetEventResponse) SetBody(v map[string]interface{}) *GetEventResponse {
	s.Body = v
	return s
}

type PatchEventHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PatchEventHeaders) String() string {
	return tea.Prettify(s)
}

func (s PatchEventHeaders) GoString() string {
	return s.String()
}

func (s *PatchEventHeaders) SetCommonHeaders(v map[string]*string) *PatchEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PatchEventHeaders) SetXAcsDingtalkAccessToken(v string) *PatchEventHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PatchEventResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PatchEventResponse) String() string {
	return tea.Prettify(s)
}

func (s PatchEventResponse) GoString() string {
	return s.String()
}

func (s *PatchEventResponse) SetHeaders(v map[string]*string) *PatchEventResponse {
	s.Headers = v
	return s
}

func (s *PatchEventResponse) SetBody(v map[string]interface{}) *PatchEventResponse {
	s.Body = v
	return s
}

type CreateEventHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateEventHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateEventHeaders) GoString() string {
	return s.String()
}

func (s *CreateEventHeaders) SetCommonHeaders(v map[string]*string) *CreateEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateEventHeaders) SetXAcsDingtalkAccessToken(v string) *CreateEventHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateEventResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEventResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEventResponse) GoString() string {
	return s.String()
}

func (s *CreateEventResponse) SetHeaders(v map[string]*string) *CreateEventResponse {
	s.Headers = v
	return s
}

func (s *CreateEventResponse) SetBody(v map[string]interface{}) *CreateEventResponse {
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

func (client *Client) DeleteEvent(userId *string, calendarId *string, eventId *string) (_result *DeleteEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteEventHeaders{}
	_result = &DeleteEventResponse{}
	_body, _err := client.DeleteEventWithOptions(userId, calendarId, eventId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEventWithOptions(userId *string, calendarId *string, eventId *string, headers *DeleteEventHeaders, runtime *util.RuntimeOptions) (_result *DeleteEventResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &DeleteEventResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteEvent"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/events/"+tea.StringValue(eventId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RespondEvent(userId *string, calendarId *string, eventId *string) (_result *RespondEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RespondEventHeaders{}
	_result = &RespondEventResponse{}
	_body, _err := client.RespondEventWithOptions(userId, calendarId, eventId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RespondEventWithOptions(userId *string, calendarId *string, eventId *string, headers *RespondEventHeaders, runtime *util.RuntimeOptions) (_result *RespondEventResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &RespondEventResponse{}
	_body, _err := client.DoROARequest(tea.String("RespondEvent"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/events/"+tea.StringValue(eventId)+"/respond"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEvents(userId *string, calendarId *string) (_result *ListEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListEventsHeaders{}
	_result = &ListEventsResponse{}
	_body, _err := client.ListEventsWithOptions(userId, calendarId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEventsWithOptions(userId *string, calendarId *string, headers *ListEventsHeaders, runtime *util.RuntimeOptions) (_result *ListEventsResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &ListEventsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListEvents"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/events"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveAttendee(userId *string, calendarId *string, eventId *string) (_result *RemoveAttendeeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveAttendeeHeaders{}
	_result = &RemoveAttendeeResponse{}
	_body, _err := client.RemoveAttendeeWithOptions(userId, calendarId, eventId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveAttendeeWithOptions(userId *string, calendarId *string, eventId *string, headers *RemoveAttendeeHeaders, runtime *util.RuntimeOptions) (_result *RemoveAttendeeResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &RemoveAttendeeResponse{}
	_body, _err := client.DoROARequest(tea.String("RemoveAttendee"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/events/"+tea.StringValue(eventId)+"/attendees"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddAttendee(userId *string, calendarId *string, eventId *string) (_result *AddAttendeeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddAttendeeHeaders{}
	_result = &AddAttendeeResponse{}
	_body, _err := client.AddAttendeeWithOptions(userId, calendarId, eventId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddAttendeeWithOptions(userId *string, calendarId *string, eventId *string, headers *AddAttendeeHeaders, runtime *util.RuntimeOptions) (_result *AddAttendeeResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &AddAttendeeResponse{}
	_body, _err := client.DoROARequest(tea.String("AddAttendee"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/events/"+tea.StringValue(eventId)+"/attendees"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEvent(userId *string, calendarId *string, eventId *string) (_result *GetEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetEventHeaders{}
	_result = &GetEventResponse{}
	_body, _err := client.GetEventWithOptions(userId, calendarId, eventId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEventWithOptions(userId *string, calendarId *string, eventId *string, headers *GetEventHeaders, runtime *util.RuntimeOptions) (_result *GetEventResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &GetEventResponse{}
	_body, _err := client.DoROARequest(tea.String("GetEvent"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/events/"+tea.StringValue(eventId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PatchEvent(userId *string, calendarId *string, eventId *string) (_result *PatchEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PatchEventHeaders{}
	_result = &PatchEventResponse{}
	_body, _err := client.PatchEventWithOptions(userId, calendarId, eventId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PatchEventWithOptions(userId *string, calendarId *string, eventId *string, headers *PatchEventHeaders, runtime *util.RuntimeOptions) (_result *PatchEventResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &PatchEventResponse{}
	_body, _err := client.DoROARequest(tea.String("PatchEvent"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/events/"+tea.StringValue(eventId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEvent(userId *string, calendarId *string) (_result *CreateEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateEventHeaders{}
	_result = &CreateEventResponse{}
	_body, _err := client.CreateEventWithOptions(userId, calendarId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEventWithOptions(userId *string, calendarId *string, headers *CreateEventHeaders, runtime *util.RuntimeOptions) (_result *CreateEventResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.XAcsDingtalkAccessToken)) {
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &CreateEventResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateEvent"), tea.String("calendar_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/calendar/users/"+tea.StringValue(userId)+"/calendars/"+tea.StringValue(calendarId)+"/events"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
