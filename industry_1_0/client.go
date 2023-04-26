// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package industry_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CampusAddRenterMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CampusAddRenterMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s CampusAddRenterMemberHeaders) GoString() string {
	return s.String()
}

func (s *CampusAddRenterMemberHeaders) SetCommonHeaders(v map[string]*string) *CampusAddRenterMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CampusAddRenterMemberHeaders) SetXAcsDingtalkAccessToken(v string) *CampusAddRenterMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CampusAddRenterMemberRequest struct {
	Extend   *string `json:"extend,omitempty" xml:"extend,omitempty"`
	Mobile   *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	RenterId *int64  `json:"renterId,omitempty" xml:"renterId,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CampusAddRenterMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s CampusAddRenterMemberRequest) GoString() string {
	return s.String()
}

func (s *CampusAddRenterMemberRequest) SetExtend(v string) *CampusAddRenterMemberRequest {
	s.Extend = &v
	return s
}

func (s *CampusAddRenterMemberRequest) SetMobile(v string) *CampusAddRenterMemberRequest {
	s.Mobile = &v
	return s
}

func (s *CampusAddRenterMemberRequest) SetName(v string) *CampusAddRenterMemberRequest {
	s.Name = &v
	return s
}

func (s *CampusAddRenterMemberRequest) SetRenterId(v int64) *CampusAddRenterMemberRequest {
	s.RenterId = &v
	return s
}

func (s *CampusAddRenterMemberRequest) SetType(v string) *CampusAddRenterMemberRequest {
	s.Type = &v
	return s
}

type CampusAddRenterMemberResponseBody struct {
	UnionId   *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
	UserState *string `json:"userState,omitempty" xml:"userState,omitempty"`
}

func (s CampusAddRenterMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CampusAddRenterMemberResponseBody) GoString() string {
	return s.String()
}

func (s *CampusAddRenterMemberResponseBody) SetUnionId(v string) *CampusAddRenterMemberResponseBody {
	s.UnionId = &v
	return s
}

func (s *CampusAddRenterMemberResponseBody) SetUserId(v string) *CampusAddRenterMemberResponseBody {
	s.UserId = &v
	return s
}

func (s *CampusAddRenterMemberResponseBody) SetUserState(v string) *CampusAddRenterMemberResponseBody {
	s.UserState = &v
	return s
}

type CampusAddRenterMemberResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CampusAddRenterMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CampusAddRenterMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s CampusAddRenterMemberResponse) GoString() string {
	return s.String()
}

func (s *CampusAddRenterMemberResponse) SetHeaders(v map[string]*string) *CampusAddRenterMemberResponse {
	s.Headers = v
	return s
}

func (s *CampusAddRenterMemberResponse) SetStatusCode(v int32) *CampusAddRenterMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *CampusAddRenterMemberResponse) SetBody(v *CampusAddRenterMemberResponseBody) *CampusAddRenterMemberResponse {
	s.Body = v
	return s
}

type CampusCreateCampusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CampusCreateCampusHeaders) String() string {
	return tea.Prettify(s)
}

func (s CampusCreateCampusHeaders) GoString() string {
	return s.String()
}

func (s *CampusCreateCampusHeaders) SetCommonHeaders(v map[string]*string) *CampusCreateCampusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CampusCreateCampusHeaders) SetXAcsDingtalkAccessToken(v string) *CampusCreateCampusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CampusCreateCampusRequest struct {
	Address              *string  `json:"address,omitempty" xml:"address,omitempty"`
	Area                 *float64 `json:"area,omitempty" xml:"area,omitempty"`
	BelongProjectGroupId *int64   `json:"belongProjectGroupId,omitempty" xml:"belongProjectGroupId,omitempty"`
	CampusName           *string  `json:"campusName,omitempty" xml:"campusName,omitempty"`
	Capacity             *int32   `json:"capacity,omitempty" xml:"capacity,omitempty"`
	CityId               *int32   `json:"cityId,omitempty" xml:"cityId,omitempty"`
	Country              *string  `json:"country,omitempty" xml:"country,omitempty"`
	CountyId             *int32   `json:"countyId,omitempty" xml:"countyId,omitempty"`
	CreatorUnionId       *string  `json:"creatorUnionId,omitempty" xml:"creatorUnionId,omitempty"`
	Description          *string  `json:"description,omitempty" xml:"description,omitempty"`
	Extend               *string  `json:"extend,omitempty" xml:"extend,omitempty"`
	Location             *string  `json:"location,omitempty" xml:"location,omitempty"`
	OrderEndTime         *int64   `json:"orderEndTime,omitempty" xml:"orderEndTime,omitempty"`
	OrderInfo            *string  `json:"orderInfo,omitempty" xml:"orderInfo,omitempty"`
	OrderStartTime       *int64   `json:"orderStartTime,omitempty" xml:"orderStartTime,omitempty"`
	ProvId               *int32   `json:"provId,omitempty" xml:"provId,omitempty"`
	Telephone            *string  `json:"telephone,omitempty" xml:"telephone,omitempty"`
}

func (s CampusCreateCampusRequest) String() string {
	return tea.Prettify(s)
}

func (s CampusCreateCampusRequest) GoString() string {
	return s.String()
}

func (s *CampusCreateCampusRequest) SetAddress(v string) *CampusCreateCampusRequest {
	s.Address = &v
	return s
}

func (s *CampusCreateCampusRequest) SetArea(v float64) *CampusCreateCampusRequest {
	s.Area = &v
	return s
}

func (s *CampusCreateCampusRequest) SetBelongProjectGroupId(v int64) *CampusCreateCampusRequest {
	s.BelongProjectGroupId = &v
	return s
}

func (s *CampusCreateCampusRequest) SetCampusName(v string) *CampusCreateCampusRequest {
	s.CampusName = &v
	return s
}

func (s *CampusCreateCampusRequest) SetCapacity(v int32) *CampusCreateCampusRequest {
	s.Capacity = &v
	return s
}

func (s *CampusCreateCampusRequest) SetCityId(v int32) *CampusCreateCampusRequest {
	s.CityId = &v
	return s
}

func (s *CampusCreateCampusRequest) SetCountry(v string) *CampusCreateCampusRequest {
	s.Country = &v
	return s
}

func (s *CampusCreateCampusRequest) SetCountyId(v int32) *CampusCreateCampusRequest {
	s.CountyId = &v
	return s
}

func (s *CampusCreateCampusRequest) SetCreatorUnionId(v string) *CampusCreateCampusRequest {
	s.CreatorUnionId = &v
	return s
}

func (s *CampusCreateCampusRequest) SetDescription(v string) *CampusCreateCampusRequest {
	s.Description = &v
	return s
}

func (s *CampusCreateCampusRequest) SetExtend(v string) *CampusCreateCampusRequest {
	s.Extend = &v
	return s
}

func (s *CampusCreateCampusRequest) SetLocation(v string) *CampusCreateCampusRequest {
	s.Location = &v
	return s
}

func (s *CampusCreateCampusRequest) SetOrderEndTime(v int64) *CampusCreateCampusRequest {
	s.OrderEndTime = &v
	return s
}

func (s *CampusCreateCampusRequest) SetOrderInfo(v string) *CampusCreateCampusRequest {
	s.OrderInfo = &v
	return s
}

func (s *CampusCreateCampusRequest) SetOrderStartTime(v int64) *CampusCreateCampusRequest {
	s.OrderStartTime = &v
	return s
}

func (s *CampusCreateCampusRequest) SetProvId(v int32) *CampusCreateCampusRequest {
	s.ProvId = &v
	return s
}

func (s *CampusCreateCampusRequest) SetTelephone(v string) *CampusCreateCampusRequest {
	s.Telephone = &v
	return s
}

type CampusCreateCampusResponseBody struct {
	CampusCorpId *string `json:"campusCorpId,omitempty" xml:"campusCorpId,omitempty"`
	CampusDeptId *string `json:"campusDeptId,omitempty" xml:"campusDeptId,omitempty"`
}

func (s CampusCreateCampusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CampusCreateCampusResponseBody) GoString() string {
	return s.String()
}

func (s *CampusCreateCampusResponseBody) SetCampusCorpId(v string) *CampusCreateCampusResponseBody {
	s.CampusCorpId = &v
	return s
}

func (s *CampusCreateCampusResponseBody) SetCampusDeptId(v string) *CampusCreateCampusResponseBody {
	s.CampusDeptId = &v
	return s
}

type CampusCreateCampusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CampusCreateCampusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CampusCreateCampusResponse) String() string {
	return tea.Prettify(s)
}

func (s CampusCreateCampusResponse) GoString() string {
	return s.String()
}

func (s *CampusCreateCampusResponse) SetHeaders(v map[string]*string) *CampusCreateCampusResponse {
	s.Headers = v
	return s
}

func (s *CampusCreateCampusResponse) SetStatusCode(v int32) *CampusCreateCampusResponse {
	s.StatusCode = &v
	return s
}

func (s *CampusCreateCampusResponse) SetBody(v *CampusCreateCampusResponseBody) *CampusCreateCampusResponse {
	s.Body = v
	return s
}

type CampusCreateCampusGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CampusCreateCampusGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s CampusCreateCampusGroupHeaders) GoString() string {
	return s.String()
}

func (s *CampusCreateCampusGroupHeaders) SetCommonHeaders(v map[string]*string) *CampusCreateCampusGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CampusCreateCampusGroupHeaders) SetXAcsDingtalkAccessToken(v string) *CampusCreateCampusGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CampusCreateCampusGroupRequest struct {
	Extend *string `json:"extend,omitempty" xml:"extend,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CampusCreateCampusGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CampusCreateCampusGroupRequest) GoString() string {
	return s.String()
}

func (s *CampusCreateCampusGroupRequest) SetExtend(v string) *CampusCreateCampusGroupRequest {
	s.Extend = &v
	return s
}

func (s *CampusCreateCampusGroupRequest) SetName(v string) *CampusCreateCampusGroupRequest {
	s.Name = &v
	return s
}

type CampusCreateCampusGroupResponseBody struct {
	GroupId *int64 `json:"groupId,omitempty" xml:"groupId,omitempty"`
}

func (s CampusCreateCampusGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CampusCreateCampusGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CampusCreateCampusGroupResponseBody) SetGroupId(v int64) *CampusCreateCampusGroupResponseBody {
	s.GroupId = &v
	return s
}

type CampusCreateCampusGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CampusCreateCampusGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CampusCreateCampusGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CampusCreateCampusGroupResponse) GoString() string {
	return s.String()
}

func (s *CampusCreateCampusGroupResponse) SetHeaders(v map[string]*string) *CampusCreateCampusGroupResponse {
	s.Headers = v
	return s
}

func (s *CampusCreateCampusGroupResponse) SetStatusCode(v int32) *CampusCreateCampusGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CampusCreateCampusGroupResponse) SetBody(v *CampusCreateCampusGroupResponseBody) *CampusCreateCampusGroupResponse {
	s.Body = v
	return s
}

type CampusCreateRenterHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CampusCreateRenterHeaders) String() string {
	return tea.Prettify(s)
}

func (s CampusCreateRenterHeaders) GoString() string {
	return s.String()
}

func (s *CampusCreateRenterHeaders) SetCommonHeaders(v map[string]*string) *CampusCreateRenterHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CampusCreateRenterHeaders) SetXAcsDingtalkAccessToken(v string) *CampusCreateRenterHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CampusCreateRenterRequest struct {
	CreditCode *string `json:"creditCode,omitempty" xml:"creditCode,omitempty"`
	EndTime    *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Extend     *string `json:"extend,omitempty" xml:"extend,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	StartTime  *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	State      *int32  `json:"state,omitempty" xml:"state,omitempty"`
}

func (s CampusCreateRenterRequest) String() string {
	return tea.Prettify(s)
}

func (s CampusCreateRenterRequest) GoString() string {
	return s.String()
}

func (s *CampusCreateRenterRequest) SetCreditCode(v string) *CampusCreateRenterRequest {
	s.CreditCode = &v
	return s
}

func (s *CampusCreateRenterRequest) SetEndTime(v int64) *CampusCreateRenterRequest {
	s.EndTime = &v
	return s
}

func (s *CampusCreateRenterRequest) SetExtend(v string) *CampusCreateRenterRequest {
	s.Extend = &v
	return s
}

func (s *CampusCreateRenterRequest) SetName(v string) *CampusCreateRenterRequest {
	s.Name = &v
	return s
}

func (s *CampusCreateRenterRequest) SetStartTime(v int64) *CampusCreateRenterRequest {
	s.StartTime = &v
	return s
}

func (s *CampusCreateRenterRequest) SetState(v int32) *CampusCreateRenterRequest {
	s.State = &v
	return s
}

type CampusCreateRenterResponseBody struct {
	RenterId *string `json:"renterId,omitempty" xml:"renterId,omitempty"`
}

func (s CampusCreateRenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CampusCreateRenterResponseBody) GoString() string {
	return s.String()
}

func (s *CampusCreateRenterResponseBody) SetRenterId(v string) *CampusCreateRenterResponseBody {
	s.RenterId = &v
	return s
}

type CampusCreateRenterResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CampusCreateRenterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CampusCreateRenterResponse) String() string {
	return tea.Prettify(s)
}

func (s CampusCreateRenterResponse) GoString() string {
	return s.String()
}

func (s *CampusCreateRenterResponse) SetHeaders(v map[string]*string) *CampusCreateRenterResponse {
	s.Headers = v
	return s
}

func (s *CampusCreateRenterResponse) SetStatusCode(v int32) *CampusCreateRenterResponse {
	s.StatusCode = &v
	return s
}

func (s *CampusCreateRenterResponse) SetBody(v *CampusCreateRenterResponseBody) *CampusCreateRenterResponse {
	s.Body = v
	return s
}

type CampusDelRenterMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CampusDelRenterMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s CampusDelRenterMemberHeaders) GoString() string {
	return s.String()
}

func (s *CampusDelRenterMemberHeaders) SetCommonHeaders(v map[string]*string) *CampusDelRenterMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CampusDelRenterMemberHeaders) SetXAcsDingtalkAccessToken(v string) *CampusDelRenterMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CampusDelRenterMemberRequest struct {
	RenterId *int64  `json:"renterId,omitempty" xml:"renterId,omitempty"`
	UnionId  *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s CampusDelRenterMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s CampusDelRenterMemberRequest) GoString() string {
	return s.String()
}

func (s *CampusDelRenterMemberRequest) SetRenterId(v int64) *CampusDelRenterMemberRequest {
	s.RenterId = &v
	return s
}

func (s *CampusDelRenterMemberRequest) SetUnionId(v string) *CampusDelRenterMemberRequest {
	s.UnionId = &v
	return s
}

type CampusDelRenterMemberResponseBody struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s CampusDelRenterMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CampusDelRenterMemberResponseBody) GoString() string {
	return s.String()
}

func (s *CampusDelRenterMemberResponseBody) SetContent(v string) *CampusDelRenterMemberResponseBody {
	s.Content = &v
	return s
}

type CampusDelRenterMemberResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CampusDelRenterMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CampusDelRenterMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s CampusDelRenterMemberResponse) GoString() string {
	return s.String()
}

func (s *CampusDelRenterMemberResponse) SetHeaders(v map[string]*string) *CampusDelRenterMemberResponse {
	s.Headers = v
	return s
}

func (s *CampusDelRenterMemberResponse) SetStatusCode(v int32) *CampusDelRenterMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *CampusDelRenterMemberResponse) SetBody(v *CampusDelRenterMemberResponseBody) *CampusDelRenterMemberResponse {
	s.Body = v
	return s
}

type CampusDeleteCampusGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CampusDeleteCampusGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s CampusDeleteCampusGroupHeaders) GoString() string {
	return s.String()
}

func (s *CampusDeleteCampusGroupHeaders) SetCommonHeaders(v map[string]*string) *CampusDeleteCampusGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CampusDeleteCampusGroupHeaders) SetXAcsDingtalkAccessToken(v string) *CampusDeleteCampusGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CampusDeleteCampusGroupRequest struct {
	CampusProjectGroupId *int64 `json:"campusProjectGroupId,omitempty" xml:"campusProjectGroupId,omitempty"`
}

func (s CampusDeleteCampusGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CampusDeleteCampusGroupRequest) GoString() string {
	return s.String()
}

func (s *CampusDeleteCampusGroupRequest) SetCampusProjectGroupId(v int64) *CampusDeleteCampusGroupRequest {
	s.CampusProjectGroupId = &v
	return s
}

type CampusDeleteCampusGroupResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CampusDeleteCampusGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CampusDeleteCampusGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CampusDeleteCampusGroupResponseBody) SetSuccess(v bool) *CampusDeleteCampusGroupResponseBody {
	s.Success = &v
	return s
}

type CampusDeleteCampusGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CampusDeleteCampusGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CampusDeleteCampusGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CampusDeleteCampusGroupResponse) GoString() string {
	return s.String()
}

func (s *CampusDeleteCampusGroupResponse) SetHeaders(v map[string]*string) *CampusDeleteCampusGroupResponse {
	s.Headers = v
	return s
}

func (s *CampusDeleteCampusGroupResponse) SetStatusCode(v int32) *CampusDeleteCampusGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CampusDeleteCampusGroupResponse) SetBody(v *CampusDeleteCampusGroupResponseBody) *CampusDeleteCampusGroupResponse {
	s.Body = v
	return s
}

type CampusDeleteRenterHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CampusDeleteRenterHeaders) String() string {
	return tea.Prettify(s)
}

func (s CampusDeleteRenterHeaders) GoString() string {
	return s.String()
}

func (s *CampusDeleteRenterHeaders) SetCommonHeaders(v map[string]*string) *CampusDeleteRenterHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CampusDeleteRenterHeaders) SetXAcsDingtalkAccessToken(v string) *CampusDeleteRenterHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CampusDeleteRenterRequest struct {
	RenterId *int64 `json:"renterId,omitempty" xml:"renterId,omitempty"`
}

func (s CampusDeleteRenterRequest) String() string {
	return tea.Prettify(s)
}

func (s CampusDeleteRenterRequest) GoString() string {
	return s.String()
}

func (s *CampusDeleteRenterRequest) SetRenterId(v int64) *CampusDeleteRenterRequest {
	s.RenterId = &v
	return s
}

type CampusDeleteRenterResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s CampusDeleteRenterResponse) String() string {
	return tea.Prettify(s)
}

func (s CampusDeleteRenterResponse) GoString() string {
	return s.String()
}

func (s *CampusDeleteRenterResponse) SetHeaders(v map[string]*string) *CampusDeleteRenterResponse {
	s.Headers = v
	return s
}

func (s *CampusDeleteRenterResponse) SetStatusCode(v int32) *CampusDeleteRenterResponse {
	s.StatusCode = &v
	return s
}

type CampusGetCampusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CampusGetCampusHeaders) String() string {
	return tea.Prettify(s)
}

func (s CampusGetCampusHeaders) GoString() string {
	return s.String()
}

func (s *CampusGetCampusHeaders) SetCommonHeaders(v map[string]*string) *CampusGetCampusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CampusGetCampusHeaders) SetXAcsDingtalkAccessToken(v string) *CampusGetCampusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CampusGetCampusRequest struct {
	CampusDeptId *int64 `json:"campusDeptId,omitempty" xml:"campusDeptId,omitempty"`
}

func (s CampusGetCampusRequest) String() string {
	return tea.Prettify(s)
}

func (s CampusGetCampusRequest) GoString() string {
	return s.String()
}

func (s *CampusGetCampusRequest) SetCampusDeptId(v int64) *CampusGetCampusRequest {
	s.CampusDeptId = &v
	return s
}

type CampusGetCampusResponseBody struct {
	Address              *string  `json:"address,omitempty" xml:"address,omitempty"`
	Area                 *float64 `json:"area,omitempty" xml:"area,omitempty"`
	BelongProjectGroupId *string  `json:"belongProjectGroupId,omitempty" xml:"belongProjectGroupId,omitempty"`
	CampusCorpId         *string  `json:"campusCorpId,omitempty" xml:"campusCorpId,omitempty"`
	CampusDeptId         *int64   `json:"campusDeptId,omitempty" xml:"campusDeptId,omitempty"`
	CampusName           *string  `json:"campusName,omitempty" xml:"campusName,omitempty"`
	Capacity             *string  `json:"capacity,omitempty" xml:"capacity,omitempty"`
	CityId               *int32   `json:"cityId,omitempty" xml:"cityId,omitempty"`
	Country              *string  `json:"country,omitempty" xml:"country,omitempty"`
	CountyId             *int32   `json:"countyId,omitempty" xml:"countyId,omitempty"`
	Description          *string  `json:"description,omitempty" xml:"description,omitempty"`
	Extend               *string  `json:"extend,omitempty" xml:"extend,omitempty"`
	Location             *string  `json:"location,omitempty" xml:"location,omitempty"`
	OrderEndTime         *int64   `json:"orderEndTime,omitempty" xml:"orderEndTime,omitempty"`
	OrderInfo            *string  `json:"orderInfo,omitempty" xml:"orderInfo,omitempty"`
	OrderStartTime       *int64   `json:"orderStartTime,omitempty" xml:"orderStartTime,omitempty"`
	ProvId               *int32   `json:"provId,omitempty" xml:"provId,omitempty"`
	Telephone            *string  `json:"telephone,omitempty" xml:"telephone,omitempty"`
}

func (s CampusGetCampusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CampusGetCampusResponseBody) GoString() string {
	return s.String()
}

func (s *CampusGetCampusResponseBody) SetAddress(v string) *CampusGetCampusResponseBody {
	s.Address = &v
	return s
}

func (s *CampusGetCampusResponseBody) SetArea(v float64) *CampusGetCampusResponseBody {
	s.Area = &v
	return s
}

func (s *CampusGetCampusResponseBody) SetBelongProjectGroupId(v string) *CampusGetCampusResponseBody {
	s.BelongProjectGroupId = &v
	return s
}

func (s *CampusGetCampusResponseBody) SetCampusCorpId(v string) *CampusGetCampusResponseBody {
	s.CampusCorpId = &v
	return s
}

func (s *CampusGetCampusResponseBody) SetCampusDeptId(v int64) *CampusGetCampusResponseBody {
	s.CampusDeptId = &v
	return s
}

func (s *CampusGetCampusResponseBody) SetCampusName(v string) *CampusGetCampusResponseBody {
	s.CampusName = &v
	return s
}

func (s *CampusGetCampusResponseBody) SetCapacity(v string) *CampusGetCampusResponseBody {
	s.Capacity = &v
	return s
}

func (s *CampusGetCampusResponseBody) SetCityId(v int32) *CampusGetCampusResponseBody {
	s.CityId = &v
	return s
}

func (s *CampusGetCampusResponseBody) SetCountry(v string) *CampusGetCampusResponseBody {
	s.Country = &v
	return s
}

func (s *CampusGetCampusResponseBody) SetCountyId(v int32) *CampusGetCampusResponseBody {
	s.CountyId = &v
	return s
}

func (s *CampusGetCampusResponseBody) SetDescription(v string) *CampusGetCampusResponseBody {
	s.Description = &v
	return s
}

func (s *CampusGetCampusResponseBody) SetExtend(v string) *CampusGetCampusResponseBody {
	s.Extend = &v
	return s
}

func (s *CampusGetCampusResponseBody) SetLocation(v string) *CampusGetCampusResponseBody {
	s.Location = &v
	return s
}

func (s *CampusGetCampusResponseBody) SetOrderEndTime(v int64) *CampusGetCampusResponseBody {
	s.OrderEndTime = &v
	return s
}

func (s *CampusGetCampusResponseBody) SetOrderInfo(v string) *CampusGetCampusResponseBody {
	s.OrderInfo = &v
	return s
}

func (s *CampusGetCampusResponseBody) SetOrderStartTime(v int64) *CampusGetCampusResponseBody {
	s.OrderStartTime = &v
	return s
}

func (s *CampusGetCampusResponseBody) SetProvId(v int32) *CampusGetCampusResponseBody {
	s.ProvId = &v
	return s
}

func (s *CampusGetCampusResponseBody) SetTelephone(v string) *CampusGetCampusResponseBody {
	s.Telephone = &v
	return s
}

type CampusGetCampusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CampusGetCampusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CampusGetCampusResponse) String() string {
	return tea.Prettify(s)
}

func (s CampusGetCampusResponse) GoString() string {
	return s.String()
}

func (s *CampusGetCampusResponse) SetHeaders(v map[string]*string) *CampusGetCampusResponse {
	s.Headers = v
	return s
}

func (s *CampusGetCampusResponse) SetStatusCode(v int32) *CampusGetCampusResponse {
	s.StatusCode = &v
	return s
}

func (s *CampusGetCampusResponse) SetBody(v *CampusGetCampusResponseBody) *CampusGetCampusResponse {
	s.Body = v
	return s
}

type CampusGetCampusGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CampusGetCampusGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s CampusGetCampusGroupHeaders) GoString() string {
	return s.String()
}

func (s *CampusGetCampusGroupHeaders) SetCommonHeaders(v map[string]*string) *CampusGetCampusGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CampusGetCampusGroupHeaders) SetXAcsDingtalkAccessToken(v string) *CampusGetCampusGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CampusGetCampusGroupRequest struct {
	GroupId *int64 `json:"groupId,omitempty" xml:"groupId,omitempty"`
}

func (s CampusGetCampusGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CampusGetCampusGroupRequest) GoString() string {
	return s.String()
}

func (s *CampusGetCampusGroupRequest) SetGroupId(v int64) *CampusGetCampusGroupRequest {
	s.GroupId = &v
	return s
}

type CampusGetCampusGroupResponseBody struct {
	Extend           *string `json:"extend,omitempty" xml:"extend,omitempty"`
	ProjectGroupName *string `json:"projectGroupName,omitempty" xml:"projectGroupName,omitempty"`
}

func (s CampusGetCampusGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CampusGetCampusGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CampusGetCampusGroupResponseBody) SetExtend(v string) *CampusGetCampusGroupResponseBody {
	s.Extend = &v
	return s
}

func (s *CampusGetCampusGroupResponseBody) SetProjectGroupName(v string) *CampusGetCampusGroupResponseBody {
	s.ProjectGroupName = &v
	return s
}

type CampusGetCampusGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CampusGetCampusGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CampusGetCampusGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CampusGetCampusGroupResponse) GoString() string {
	return s.String()
}

func (s *CampusGetCampusGroupResponse) SetHeaders(v map[string]*string) *CampusGetCampusGroupResponse {
	s.Headers = v
	return s
}

func (s *CampusGetCampusGroupResponse) SetStatusCode(v int32) *CampusGetCampusGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CampusGetCampusGroupResponse) SetBody(v *CampusGetCampusGroupResponseBody) *CampusGetCampusGroupResponse {
	s.Body = v
	return s
}

type CampusGetRenterHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CampusGetRenterHeaders) String() string {
	return tea.Prettify(s)
}

func (s CampusGetRenterHeaders) GoString() string {
	return s.String()
}

func (s *CampusGetRenterHeaders) SetCommonHeaders(v map[string]*string) *CampusGetRenterHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CampusGetRenterHeaders) SetXAcsDingtalkAccessToken(v string) *CampusGetRenterHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CampusGetRenterRequest struct {
	RenterId *int64 `json:"renterId,omitempty" xml:"renterId,omitempty"`
}

func (s CampusGetRenterRequest) String() string {
	return tea.Prettify(s)
}

func (s CampusGetRenterRequest) GoString() string {
	return s.String()
}

func (s *CampusGetRenterRequest) SetRenterId(v int64) *CampusGetRenterRequest {
	s.RenterId = &v
	return s
}

type CampusGetRenterResponseBody struct {
	BindRenterCorpId *string `json:"bindRenterCorpId,omitempty" xml:"bindRenterCorpId,omitempty"`
	BindTime         *int64  `json:"bindTime,omitempty" xml:"bindTime,omitempty"`
	CreditCode       *string `json:"creditCode,omitempty" xml:"creditCode,omitempty"`
	EndTime          *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Extend           *string `json:"extend,omitempty" xml:"extend,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
	RenterDeptId     *int64  `json:"renterDeptId,omitempty" xml:"renterDeptId,omitempty"`
	StartTime        *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	State            *int32  `json:"state,omitempty" xml:"state,omitempty"`
}

func (s CampusGetRenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CampusGetRenterResponseBody) GoString() string {
	return s.String()
}

func (s *CampusGetRenterResponseBody) SetBindRenterCorpId(v string) *CampusGetRenterResponseBody {
	s.BindRenterCorpId = &v
	return s
}

func (s *CampusGetRenterResponseBody) SetBindTime(v int64) *CampusGetRenterResponseBody {
	s.BindTime = &v
	return s
}

func (s *CampusGetRenterResponseBody) SetCreditCode(v string) *CampusGetRenterResponseBody {
	s.CreditCode = &v
	return s
}

func (s *CampusGetRenterResponseBody) SetEndTime(v int64) *CampusGetRenterResponseBody {
	s.EndTime = &v
	return s
}

func (s *CampusGetRenterResponseBody) SetExtend(v string) *CampusGetRenterResponseBody {
	s.Extend = &v
	return s
}

func (s *CampusGetRenterResponseBody) SetName(v string) *CampusGetRenterResponseBody {
	s.Name = &v
	return s
}

func (s *CampusGetRenterResponseBody) SetRenterDeptId(v int64) *CampusGetRenterResponseBody {
	s.RenterDeptId = &v
	return s
}

func (s *CampusGetRenterResponseBody) SetStartTime(v int64) *CampusGetRenterResponseBody {
	s.StartTime = &v
	return s
}

func (s *CampusGetRenterResponseBody) SetState(v int32) *CampusGetRenterResponseBody {
	s.State = &v
	return s
}

type CampusGetRenterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CampusGetRenterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CampusGetRenterResponse) String() string {
	return tea.Prettify(s)
}

func (s CampusGetRenterResponse) GoString() string {
	return s.String()
}

func (s *CampusGetRenterResponse) SetHeaders(v map[string]*string) *CampusGetRenterResponse {
	s.Headers = v
	return s
}

func (s *CampusGetRenterResponse) SetStatusCode(v int32) *CampusGetRenterResponse {
	s.StatusCode = &v
	return s
}

func (s *CampusGetRenterResponse) SetBody(v *CampusGetRenterResponseBody) *CampusGetRenterResponse {
	s.Body = v
	return s
}

type CampusGetRenterMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CampusGetRenterMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s CampusGetRenterMemberHeaders) GoString() string {
	return s.String()
}

func (s *CampusGetRenterMemberHeaders) SetCommonHeaders(v map[string]*string) *CampusGetRenterMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CampusGetRenterMemberHeaders) SetXAcsDingtalkAccessToken(v string) *CampusGetRenterMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CampusGetRenterMemberRequest struct {
	RenterId *int64  `json:"renterId,omitempty" xml:"renterId,omitempty"`
	UnionId  *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s CampusGetRenterMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s CampusGetRenterMemberRequest) GoString() string {
	return s.String()
}

func (s *CampusGetRenterMemberRequest) SetRenterId(v int64) *CampusGetRenterMemberRequest {
	s.RenterId = &v
	return s
}

func (s *CampusGetRenterMemberRequest) SetUnionId(v string) *CampusGetRenterMemberRequest {
	s.UnionId = &v
	return s
}

type CampusGetRenterMemberResponseBody struct {
	Extend      *string `json:"extend,omitempty" xml:"extend,omitempty"`
	InviteState *int32  `json:"inviteState,omitempty" xml:"inviteState,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	State       *string `json:"state,omitempty" xml:"state,omitempty"`
	Type        *string `json:"type,omitempty" xml:"type,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CampusGetRenterMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CampusGetRenterMemberResponseBody) GoString() string {
	return s.String()
}

func (s *CampusGetRenterMemberResponseBody) SetExtend(v string) *CampusGetRenterMemberResponseBody {
	s.Extend = &v
	return s
}

func (s *CampusGetRenterMemberResponseBody) SetInviteState(v int32) *CampusGetRenterMemberResponseBody {
	s.InviteState = &v
	return s
}

func (s *CampusGetRenterMemberResponseBody) SetName(v string) *CampusGetRenterMemberResponseBody {
	s.Name = &v
	return s
}

func (s *CampusGetRenterMemberResponseBody) SetState(v string) *CampusGetRenterMemberResponseBody {
	s.State = &v
	return s
}

func (s *CampusGetRenterMemberResponseBody) SetType(v string) *CampusGetRenterMemberResponseBody {
	s.Type = &v
	return s
}

func (s *CampusGetRenterMemberResponseBody) SetUserId(v string) *CampusGetRenterMemberResponseBody {
	s.UserId = &v
	return s
}

type CampusGetRenterMemberResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CampusGetRenterMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CampusGetRenterMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s CampusGetRenterMemberResponse) GoString() string {
	return s.String()
}

func (s *CampusGetRenterMemberResponse) SetHeaders(v map[string]*string) *CampusGetRenterMemberResponse {
	s.Headers = v
	return s
}

func (s *CampusGetRenterMemberResponse) SetStatusCode(v int32) *CampusGetRenterMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *CampusGetRenterMemberResponse) SetBody(v *CampusGetRenterMemberResponseBody) *CampusGetRenterMemberResponse {
	s.Body = v
	return s
}

type CampusListCampusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CampusListCampusHeaders) String() string {
	return tea.Prettify(s)
}

func (s CampusListCampusHeaders) GoString() string {
	return s.String()
}

func (s *CampusListCampusHeaders) SetCommonHeaders(v map[string]*string) *CampusListCampusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CampusListCampusHeaders) SetXAcsDingtalkAccessToken(v string) *CampusListCampusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CampusListCampusRequest struct {
	GroupDeptId *int64 `json:"groupDeptId,omitempty" xml:"groupDeptId,omitempty"`
}

func (s CampusListCampusRequest) String() string {
	return tea.Prettify(s)
}

func (s CampusListCampusRequest) GoString() string {
	return s.String()
}

func (s *CampusListCampusRequest) SetGroupDeptId(v int64) *CampusListCampusRequest {
	s.GroupDeptId = &v
	return s
}

type CampusListCampusResponseBody struct {
	Result []*CampusListCampusResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s CampusListCampusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CampusListCampusResponseBody) GoString() string {
	return s.String()
}

func (s *CampusListCampusResponseBody) SetResult(v []*CampusListCampusResponseBodyResult) *CampusListCampusResponseBody {
	s.Result = v
	return s
}

type CampusListCampusResponseBodyResult struct {
	Address              *string  `json:"address,omitempty" xml:"address,omitempty"`
	Area                 *float64 `json:"area,omitempty" xml:"area,omitempty"`
	BelongProjectGroupId *int64   `json:"belongProjectGroupId,omitempty" xml:"belongProjectGroupId,omitempty"`
	CampusCorpId         *string  `json:"campusCorpId,omitempty" xml:"campusCorpId,omitempty"`
	CampusDeptId         *int64   `json:"campusDeptId,omitempty" xml:"campusDeptId,omitempty"`
	CampusName           *string  `json:"campusName,omitempty" xml:"campusName,omitempty"`
	CityId               *int32   `json:"cityId,omitempty" xml:"cityId,omitempty"`
	Country              *string  `json:"country,omitempty" xml:"country,omitempty"`
	CountyId             *int32   `json:"countyId,omitempty" xml:"countyId,omitempty"`
	Description          *string  `json:"description,omitempty" xml:"description,omitempty"`
	Extend               *string  `json:"extend,omitempty" xml:"extend,omitempty"`
	Location             *string  `json:"location,omitempty" xml:"location,omitempty"`
	OrderEndTime         *int64   `json:"orderEndTime,omitempty" xml:"orderEndTime,omitempty"`
	OrderInfo            *string  `json:"orderInfo,omitempty" xml:"orderInfo,omitempty"`
	OrderStartTime       *int64   `json:"orderStartTime,omitempty" xml:"orderStartTime,omitempty"`
	ProvId               *int32   `json:"provId,omitempty" xml:"provId,omitempty"`
	Telephone            *string  `json:"telephone,omitempty" xml:"telephone,omitempty"`
}

func (s CampusListCampusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CampusListCampusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CampusListCampusResponseBodyResult) SetAddress(v string) *CampusListCampusResponseBodyResult {
	s.Address = &v
	return s
}

func (s *CampusListCampusResponseBodyResult) SetArea(v float64) *CampusListCampusResponseBodyResult {
	s.Area = &v
	return s
}

func (s *CampusListCampusResponseBodyResult) SetBelongProjectGroupId(v int64) *CampusListCampusResponseBodyResult {
	s.BelongProjectGroupId = &v
	return s
}

func (s *CampusListCampusResponseBodyResult) SetCampusCorpId(v string) *CampusListCampusResponseBodyResult {
	s.CampusCorpId = &v
	return s
}

func (s *CampusListCampusResponseBodyResult) SetCampusDeptId(v int64) *CampusListCampusResponseBodyResult {
	s.CampusDeptId = &v
	return s
}

func (s *CampusListCampusResponseBodyResult) SetCampusName(v string) *CampusListCampusResponseBodyResult {
	s.CampusName = &v
	return s
}

func (s *CampusListCampusResponseBodyResult) SetCityId(v int32) *CampusListCampusResponseBodyResult {
	s.CityId = &v
	return s
}

func (s *CampusListCampusResponseBodyResult) SetCountry(v string) *CampusListCampusResponseBodyResult {
	s.Country = &v
	return s
}

func (s *CampusListCampusResponseBodyResult) SetCountyId(v int32) *CampusListCampusResponseBodyResult {
	s.CountyId = &v
	return s
}

func (s *CampusListCampusResponseBodyResult) SetDescription(v string) *CampusListCampusResponseBodyResult {
	s.Description = &v
	return s
}

func (s *CampusListCampusResponseBodyResult) SetExtend(v string) *CampusListCampusResponseBodyResult {
	s.Extend = &v
	return s
}

func (s *CampusListCampusResponseBodyResult) SetLocation(v string) *CampusListCampusResponseBodyResult {
	s.Location = &v
	return s
}

func (s *CampusListCampusResponseBodyResult) SetOrderEndTime(v int64) *CampusListCampusResponseBodyResult {
	s.OrderEndTime = &v
	return s
}

func (s *CampusListCampusResponseBodyResult) SetOrderInfo(v string) *CampusListCampusResponseBodyResult {
	s.OrderInfo = &v
	return s
}

func (s *CampusListCampusResponseBodyResult) SetOrderStartTime(v int64) *CampusListCampusResponseBodyResult {
	s.OrderStartTime = &v
	return s
}

func (s *CampusListCampusResponseBodyResult) SetProvId(v int32) *CampusListCampusResponseBodyResult {
	s.ProvId = &v
	return s
}

func (s *CampusListCampusResponseBodyResult) SetTelephone(v string) *CampusListCampusResponseBodyResult {
	s.Telephone = &v
	return s
}

type CampusListCampusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CampusListCampusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CampusListCampusResponse) String() string {
	return tea.Prettify(s)
}

func (s CampusListCampusResponse) GoString() string {
	return s.String()
}

func (s *CampusListCampusResponse) SetHeaders(v map[string]*string) *CampusListCampusResponse {
	s.Headers = v
	return s
}

func (s *CampusListCampusResponse) SetStatusCode(v int32) *CampusListCampusResponse {
	s.StatusCode = &v
	return s
}

func (s *CampusListCampusResponse) SetBody(v *CampusListCampusResponseBody) *CampusListCampusResponse {
	s.Body = v
	return s
}

type CampusListCampusGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CampusListCampusGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s CampusListCampusGroupHeaders) GoString() string {
	return s.String()
}

func (s *CampusListCampusGroupHeaders) SetCommonHeaders(v map[string]*string) *CampusListCampusGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CampusListCampusGroupHeaders) SetXAcsDingtalkAccessToken(v string) *CampusListCampusGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CampusListCampusGroupResponseBody struct {
	Result []*CampusListCampusGroupResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s CampusListCampusGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CampusListCampusGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CampusListCampusGroupResponseBody) SetResult(v []*CampusListCampusGroupResponseBodyResult) *CampusListCampusGroupResponseBody {
	s.Result = v
	return s
}

type CampusListCampusGroupResponseBodyResult struct {
	Extend      *string `json:"extend,omitempty" xml:"extend,omitempty"`
	GroupDeptId *int64  `json:"groupDeptId,omitempty" xml:"groupDeptId,omitempty"`
	GroupName   *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
}

func (s CampusListCampusGroupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CampusListCampusGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CampusListCampusGroupResponseBodyResult) SetExtend(v string) *CampusListCampusGroupResponseBodyResult {
	s.Extend = &v
	return s
}

func (s *CampusListCampusGroupResponseBodyResult) SetGroupDeptId(v int64) *CampusListCampusGroupResponseBodyResult {
	s.GroupDeptId = &v
	return s
}

func (s *CampusListCampusGroupResponseBodyResult) SetGroupName(v string) *CampusListCampusGroupResponseBodyResult {
	s.GroupName = &v
	return s
}

type CampusListCampusGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CampusListCampusGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CampusListCampusGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CampusListCampusGroupResponse) GoString() string {
	return s.String()
}

func (s *CampusListCampusGroupResponse) SetHeaders(v map[string]*string) *CampusListCampusGroupResponse {
	s.Headers = v
	return s
}

func (s *CampusListCampusGroupResponse) SetStatusCode(v int32) *CampusListCampusGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CampusListCampusGroupResponse) SetBody(v *CampusListCampusGroupResponseBody) *CampusListCampusGroupResponse {
	s.Body = v
	return s
}

type CampusListRenterHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CampusListRenterHeaders) String() string {
	return tea.Prettify(s)
}

func (s CampusListRenterHeaders) GoString() string {
	return s.String()
}

func (s *CampusListRenterHeaders) SetCommonHeaders(v map[string]*string) *CampusListRenterHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CampusListRenterHeaders) SetXAcsDingtalkAccessToken(v string) *CampusListRenterHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CampusListRenterResponseBody struct {
	Result []*CampusListRenterResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s CampusListRenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CampusListRenterResponseBody) GoString() string {
	return s.String()
}

func (s *CampusListRenterResponseBody) SetResult(v []*CampusListRenterResponseBodyResult) *CampusListRenterResponseBody {
	s.Result = v
	return s
}

type CampusListRenterResponseBodyResult struct {
	BindRenterCorpId *string `json:"bindRenterCorpId,omitempty" xml:"bindRenterCorpId,omitempty"`
	BindTime         *int64  `json:"bindTime,omitempty" xml:"bindTime,omitempty"`
	CreditCode       *string `json:"creditCode,omitempty" xml:"creditCode,omitempty"`
	EndTime          *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Extend           *string `json:"extend,omitempty" xml:"extend,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
	RenterDeptId     *int64  `json:"renterDeptId,omitempty" xml:"renterDeptId,omitempty"`
	StartTime        *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	State            *int32  `json:"state,omitempty" xml:"state,omitempty"`
}

func (s CampusListRenterResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CampusListRenterResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CampusListRenterResponseBodyResult) SetBindRenterCorpId(v string) *CampusListRenterResponseBodyResult {
	s.BindRenterCorpId = &v
	return s
}

func (s *CampusListRenterResponseBodyResult) SetBindTime(v int64) *CampusListRenterResponseBodyResult {
	s.BindTime = &v
	return s
}

func (s *CampusListRenterResponseBodyResult) SetCreditCode(v string) *CampusListRenterResponseBodyResult {
	s.CreditCode = &v
	return s
}

func (s *CampusListRenterResponseBodyResult) SetEndTime(v int64) *CampusListRenterResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *CampusListRenterResponseBodyResult) SetExtend(v string) *CampusListRenterResponseBodyResult {
	s.Extend = &v
	return s
}

func (s *CampusListRenterResponseBodyResult) SetName(v string) *CampusListRenterResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CampusListRenterResponseBodyResult) SetRenterDeptId(v int64) *CampusListRenterResponseBodyResult {
	s.RenterDeptId = &v
	return s
}

func (s *CampusListRenterResponseBodyResult) SetStartTime(v int64) *CampusListRenterResponseBodyResult {
	s.StartTime = &v
	return s
}

func (s *CampusListRenterResponseBodyResult) SetState(v int32) *CampusListRenterResponseBodyResult {
	s.State = &v
	return s
}

type CampusListRenterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CampusListRenterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CampusListRenterResponse) String() string {
	return tea.Prettify(s)
}

func (s CampusListRenterResponse) GoString() string {
	return s.String()
}

func (s *CampusListRenterResponse) SetHeaders(v map[string]*string) *CampusListRenterResponse {
	s.Headers = v
	return s
}

func (s *CampusListRenterResponse) SetStatusCode(v int32) *CampusListRenterResponse {
	s.StatusCode = &v
	return s
}

func (s *CampusListRenterResponse) SetBody(v *CampusListRenterResponseBody) *CampusListRenterResponse {
	s.Body = v
	return s
}

type CampusListRenterMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CampusListRenterMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s CampusListRenterMembersHeaders) GoString() string {
	return s.String()
}

func (s *CampusListRenterMembersHeaders) SetCommonHeaders(v map[string]*string) *CampusListRenterMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CampusListRenterMembersHeaders) SetXAcsDingtalkAccessToken(v string) *CampusListRenterMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CampusListRenterMembersRequest struct {
	RenterId *int64 `json:"renterId,omitempty" xml:"renterId,omitempty"`
}

func (s CampusListRenterMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s CampusListRenterMembersRequest) GoString() string {
	return s.String()
}

func (s *CampusListRenterMembersRequest) SetRenterId(v int64) *CampusListRenterMembersRequest {
	s.RenterId = &v
	return s
}

type CampusListRenterMembersResponseBody struct {
	Result []*CampusListRenterMembersResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s CampusListRenterMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CampusListRenterMembersResponseBody) GoString() string {
	return s.String()
}

func (s *CampusListRenterMembersResponseBody) SetResult(v []*CampusListRenterMembersResponseBodyResult) *CampusListRenterMembersResponseBody {
	s.Result = v
	return s
}

type CampusListRenterMembersResponseBodyResult struct {
	Extend      *string `json:"extend,omitempty" xml:"extend,omitempty"`
	InviteState *string `json:"inviteState,omitempty" xml:"inviteState,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	State       *string `json:"state,omitempty" xml:"state,omitempty"`
	Type        *string `json:"type,omitempty" xml:"type,omitempty"`
	UnionId     *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CampusListRenterMembersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CampusListRenterMembersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CampusListRenterMembersResponseBodyResult) SetExtend(v string) *CampusListRenterMembersResponseBodyResult {
	s.Extend = &v
	return s
}

func (s *CampusListRenterMembersResponseBodyResult) SetInviteState(v string) *CampusListRenterMembersResponseBodyResult {
	s.InviteState = &v
	return s
}

func (s *CampusListRenterMembersResponseBodyResult) SetName(v string) *CampusListRenterMembersResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CampusListRenterMembersResponseBodyResult) SetState(v string) *CampusListRenterMembersResponseBodyResult {
	s.State = &v
	return s
}

func (s *CampusListRenterMembersResponseBodyResult) SetType(v string) *CampusListRenterMembersResponseBodyResult {
	s.Type = &v
	return s
}

func (s *CampusListRenterMembersResponseBodyResult) SetUnionId(v string) *CampusListRenterMembersResponseBodyResult {
	s.UnionId = &v
	return s
}

func (s *CampusListRenterMembersResponseBodyResult) SetUserId(v string) *CampusListRenterMembersResponseBodyResult {
	s.UserId = &v
	return s
}

type CampusListRenterMembersResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CampusListRenterMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CampusListRenterMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s CampusListRenterMembersResponse) GoString() string {
	return s.String()
}

func (s *CampusListRenterMembersResponse) SetHeaders(v map[string]*string) *CampusListRenterMembersResponse {
	s.Headers = v
	return s
}

func (s *CampusListRenterMembersResponse) SetStatusCode(v int32) *CampusListRenterMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *CampusListRenterMembersResponse) SetBody(v *CampusListRenterMembersResponseBody) *CampusListRenterMembersResponse {
	s.Body = v
	return s
}

type CampusUpdateCampusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CampusUpdateCampusHeaders) String() string {
	return tea.Prettify(s)
}

func (s CampusUpdateCampusHeaders) GoString() string {
	return s.String()
}

func (s *CampusUpdateCampusHeaders) SetCommonHeaders(v map[string]*string) *CampusUpdateCampusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CampusUpdateCampusHeaders) SetXAcsDingtalkAccessToken(v string) *CampusUpdateCampusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CampusUpdateCampusRequest struct {
	Address              *string  `json:"address,omitempty" xml:"address,omitempty"`
	Area                 *float64 `json:"area,omitempty" xml:"area,omitempty"`
	BelongProjectGroupId *int64   `json:"belongProjectGroupId,omitempty" xml:"belongProjectGroupId,omitempty"`
	CampusDeptId         *int64   `json:"campusDeptId,omitempty" xml:"campusDeptId,omitempty"`
	CampusName           *string  `json:"campusName,omitempty" xml:"campusName,omitempty"`
	Capacity             *int32   `json:"capacity,omitempty" xml:"capacity,omitempty"`
	CityId               *int32   `json:"cityId,omitempty" xml:"cityId,omitempty"`
	Country              *string  `json:"country,omitempty" xml:"country,omitempty"`
	CountyId             *int32   `json:"countyId,omitempty" xml:"countyId,omitempty"`
	Description          *string  `json:"description,omitempty" xml:"description,omitempty"`
	Extend               *string  `json:"extend,omitempty" xml:"extend,omitempty"`
	OrderEndTime         *int64   `json:"orderEndTime,omitempty" xml:"orderEndTime,omitempty"`
	OrderInfo            *int64   `json:"orderInfo,omitempty" xml:"orderInfo,omitempty"`
	OrderStartTime       *int64   `json:"orderStartTime,omitempty" xml:"orderStartTime,omitempty"`
	ProvId               *int32   `json:"provId,omitempty" xml:"provId,omitempty"`
	Telephone            *string  `json:"telephone,omitempty" xml:"telephone,omitempty"`
}

func (s CampusUpdateCampusRequest) String() string {
	return tea.Prettify(s)
}

func (s CampusUpdateCampusRequest) GoString() string {
	return s.String()
}

func (s *CampusUpdateCampusRequest) SetAddress(v string) *CampusUpdateCampusRequest {
	s.Address = &v
	return s
}

func (s *CampusUpdateCampusRequest) SetArea(v float64) *CampusUpdateCampusRequest {
	s.Area = &v
	return s
}

func (s *CampusUpdateCampusRequest) SetBelongProjectGroupId(v int64) *CampusUpdateCampusRequest {
	s.BelongProjectGroupId = &v
	return s
}

func (s *CampusUpdateCampusRequest) SetCampusDeptId(v int64) *CampusUpdateCampusRequest {
	s.CampusDeptId = &v
	return s
}

func (s *CampusUpdateCampusRequest) SetCampusName(v string) *CampusUpdateCampusRequest {
	s.CampusName = &v
	return s
}

func (s *CampusUpdateCampusRequest) SetCapacity(v int32) *CampusUpdateCampusRequest {
	s.Capacity = &v
	return s
}

func (s *CampusUpdateCampusRequest) SetCityId(v int32) *CampusUpdateCampusRequest {
	s.CityId = &v
	return s
}

func (s *CampusUpdateCampusRequest) SetCountry(v string) *CampusUpdateCampusRequest {
	s.Country = &v
	return s
}

func (s *CampusUpdateCampusRequest) SetCountyId(v int32) *CampusUpdateCampusRequest {
	s.CountyId = &v
	return s
}

func (s *CampusUpdateCampusRequest) SetDescription(v string) *CampusUpdateCampusRequest {
	s.Description = &v
	return s
}

func (s *CampusUpdateCampusRequest) SetExtend(v string) *CampusUpdateCampusRequest {
	s.Extend = &v
	return s
}

func (s *CampusUpdateCampusRequest) SetOrderEndTime(v int64) *CampusUpdateCampusRequest {
	s.OrderEndTime = &v
	return s
}

func (s *CampusUpdateCampusRequest) SetOrderInfo(v int64) *CampusUpdateCampusRequest {
	s.OrderInfo = &v
	return s
}

func (s *CampusUpdateCampusRequest) SetOrderStartTime(v int64) *CampusUpdateCampusRequest {
	s.OrderStartTime = &v
	return s
}

func (s *CampusUpdateCampusRequest) SetProvId(v int32) *CampusUpdateCampusRequest {
	s.ProvId = &v
	return s
}

func (s *CampusUpdateCampusRequest) SetTelephone(v string) *CampusUpdateCampusRequest {
	s.Telephone = &v
	return s
}

type CampusUpdateCampusResponseBody struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s CampusUpdateCampusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CampusUpdateCampusResponseBody) GoString() string {
	return s.String()
}

func (s *CampusUpdateCampusResponseBody) SetContent(v string) *CampusUpdateCampusResponseBody {
	s.Content = &v
	return s
}

type CampusUpdateCampusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CampusUpdateCampusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CampusUpdateCampusResponse) String() string {
	return tea.Prettify(s)
}

func (s CampusUpdateCampusResponse) GoString() string {
	return s.String()
}

func (s *CampusUpdateCampusResponse) SetHeaders(v map[string]*string) *CampusUpdateCampusResponse {
	s.Headers = v
	return s
}

func (s *CampusUpdateCampusResponse) SetStatusCode(v int32) *CampusUpdateCampusResponse {
	s.StatusCode = &v
	return s
}

func (s *CampusUpdateCampusResponse) SetBody(v *CampusUpdateCampusResponseBody) *CampusUpdateCampusResponse {
	s.Body = v
	return s
}

type CampusUpdateCampusGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CampusUpdateCampusGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s CampusUpdateCampusGroupHeaders) GoString() string {
	return s.String()
}

func (s *CampusUpdateCampusGroupHeaders) SetCommonHeaders(v map[string]*string) *CampusUpdateCampusGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CampusUpdateCampusGroupHeaders) SetXAcsDingtalkAccessToken(v string) *CampusUpdateCampusGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CampusUpdateCampusGroupRequest struct {
	CampusProjectGroupId *int64  `json:"campusProjectGroupId,omitempty" xml:"campusProjectGroupId,omitempty"`
	Extend               *string `json:"extend,omitempty" xml:"extend,omitempty"`
	Name                 *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CampusUpdateCampusGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CampusUpdateCampusGroupRequest) GoString() string {
	return s.String()
}

func (s *CampusUpdateCampusGroupRequest) SetCampusProjectGroupId(v int64) *CampusUpdateCampusGroupRequest {
	s.CampusProjectGroupId = &v
	return s
}

func (s *CampusUpdateCampusGroupRequest) SetExtend(v string) *CampusUpdateCampusGroupRequest {
	s.Extend = &v
	return s
}

func (s *CampusUpdateCampusGroupRequest) SetName(v string) *CampusUpdateCampusGroupRequest {
	s.Name = &v
	return s
}

type CampusUpdateCampusGroupResponseBody struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s CampusUpdateCampusGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CampusUpdateCampusGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CampusUpdateCampusGroupResponseBody) SetContent(v string) *CampusUpdateCampusGroupResponseBody {
	s.Content = &v
	return s
}

type CampusUpdateCampusGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CampusUpdateCampusGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CampusUpdateCampusGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CampusUpdateCampusGroupResponse) GoString() string {
	return s.String()
}

func (s *CampusUpdateCampusGroupResponse) SetHeaders(v map[string]*string) *CampusUpdateCampusGroupResponse {
	s.Headers = v
	return s
}

func (s *CampusUpdateCampusGroupResponse) SetStatusCode(v int32) *CampusUpdateCampusGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CampusUpdateCampusGroupResponse) SetBody(v *CampusUpdateCampusGroupResponseBody) *CampusUpdateCampusGroupResponse {
	s.Body = v
	return s
}

type CampusUpdateRenterHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CampusUpdateRenterHeaders) String() string {
	return tea.Prettify(s)
}

func (s CampusUpdateRenterHeaders) GoString() string {
	return s.String()
}

func (s *CampusUpdateRenterHeaders) SetCommonHeaders(v map[string]*string) *CampusUpdateRenterHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CampusUpdateRenterHeaders) SetXAcsDingtalkAccessToken(v string) *CampusUpdateRenterHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CampusUpdateRenterRequest struct {
	CreditCode *string `json:"creditCode,omitempty" xml:"creditCode,omitempty"`
	EndTime    *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Extend     *string `json:"extend,omitempty" xml:"extend,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	RenterId   *int64  `json:"renterId,omitempty" xml:"renterId,omitempty"`
	StartTime  *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	State      *int32  `json:"state,omitempty" xml:"state,omitempty"`
}

func (s CampusUpdateRenterRequest) String() string {
	return tea.Prettify(s)
}

func (s CampusUpdateRenterRequest) GoString() string {
	return s.String()
}

func (s *CampusUpdateRenterRequest) SetCreditCode(v string) *CampusUpdateRenterRequest {
	s.CreditCode = &v
	return s
}

func (s *CampusUpdateRenterRequest) SetEndTime(v int64) *CampusUpdateRenterRequest {
	s.EndTime = &v
	return s
}

func (s *CampusUpdateRenterRequest) SetExtend(v string) *CampusUpdateRenterRequest {
	s.Extend = &v
	return s
}

func (s *CampusUpdateRenterRequest) SetName(v string) *CampusUpdateRenterRequest {
	s.Name = &v
	return s
}

func (s *CampusUpdateRenterRequest) SetRenterId(v int64) *CampusUpdateRenterRequest {
	s.RenterId = &v
	return s
}

func (s *CampusUpdateRenterRequest) SetStartTime(v int64) *CampusUpdateRenterRequest {
	s.StartTime = &v
	return s
}

func (s *CampusUpdateRenterRequest) SetState(v int32) *CampusUpdateRenterRequest {
	s.State = &v
	return s
}

type CampusUpdateRenterResponseBody struct {
	RenterId *string `json:"renterId,omitempty" xml:"renterId,omitempty"`
}

func (s CampusUpdateRenterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CampusUpdateRenterResponseBody) GoString() string {
	return s.String()
}

func (s *CampusUpdateRenterResponseBody) SetRenterId(v string) *CampusUpdateRenterResponseBody {
	s.RenterId = &v
	return s
}

type CampusUpdateRenterResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CampusUpdateRenterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CampusUpdateRenterResponse) String() string {
	return tea.Prettify(s)
}

func (s CampusUpdateRenterResponse) GoString() string {
	return s.String()
}

func (s *CampusUpdateRenterResponse) SetHeaders(v map[string]*string) *CampusUpdateRenterResponse {
	s.Headers = v
	return s
}

func (s *CampusUpdateRenterResponse) SetStatusCode(v int32) *CampusUpdateRenterResponse {
	s.StatusCode = &v
	return s
}

func (s *CampusUpdateRenterResponse) SetBody(v *CampusUpdateRenterResponseBody) *CampusUpdateRenterResponse {
	s.Body = v
	return s
}

type CampusUpdateRenterMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CampusUpdateRenterMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s CampusUpdateRenterMemberHeaders) GoString() string {
	return s.String()
}

func (s *CampusUpdateRenterMemberHeaders) SetCommonHeaders(v map[string]*string) *CampusUpdateRenterMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CampusUpdateRenterMemberHeaders) SetXAcsDingtalkAccessToken(v string) *CampusUpdateRenterMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CampusUpdateRenterMemberRequest struct {
	Extend   *string `json:"extend,omitempty" xml:"extend,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	RenterId *int64  `json:"renterId,omitempty" xml:"renterId,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
	UnionId  *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
}

func (s CampusUpdateRenterMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s CampusUpdateRenterMemberRequest) GoString() string {
	return s.String()
}

func (s *CampusUpdateRenterMemberRequest) SetExtend(v string) *CampusUpdateRenterMemberRequest {
	s.Extend = &v
	return s
}

func (s *CampusUpdateRenterMemberRequest) SetName(v string) *CampusUpdateRenterMemberRequest {
	s.Name = &v
	return s
}

func (s *CampusUpdateRenterMemberRequest) SetRenterId(v int64) *CampusUpdateRenterMemberRequest {
	s.RenterId = &v
	return s
}

func (s *CampusUpdateRenterMemberRequest) SetType(v string) *CampusUpdateRenterMemberRequest {
	s.Type = &v
	return s
}

func (s *CampusUpdateRenterMemberRequest) SetUnionId(v string) *CampusUpdateRenterMemberRequest {
	s.UnionId = &v
	return s
}

type CampusUpdateRenterMemberResponseBody struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s CampusUpdateRenterMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CampusUpdateRenterMemberResponseBody) GoString() string {
	return s.String()
}

func (s *CampusUpdateRenterMemberResponseBody) SetContent(v string) *CampusUpdateRenterMemberResponseBody {
	s.Content = &v
	return s
}

type CampusUpdateRenterMemberResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CampusUpdateRenterMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CampusUpdateRenterMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s CampusUpdateRenterMemberResponse) GoString() string {
	return s.String()
}

func (s *CampusUpdateRenterMemberResponse) SetHeaders(v map[string]*string) *CampusUpdateRenterMemberResponse {
	s.Headers = v
	return s
}

func (s *CampusUpdateRenterMemberResponse) SetStatusCode(v int32) *CampusUpdateRenterMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *CampusUpdateRenterMemberResponse) SetBody(v *CampusUpdateRenterMemberResponseBody) *CampusUpdateRenterMemberResponse {
	s.Body = v
	return s
}

type CollegeActiveCollegeDeptGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CollegeActiveCollegeDeptGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s CollegeActiveCollegeDeptGroupHeaders) GoString() string {
	return s.String()
}

func (s *CollegeActiveCollegeDeptGroupHeaders) SetCommonHeaders(v map[string]*string) *CollegeActiveCollegeDeptGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CollegeActiveCollegeDeptGroupHeaders) SetXAcsDingtalkAccessToken(v string) *CollegeActiveCollegeDeptGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CollegeActiveCollegeDeptGroupRequest struct {
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s CollegeActiveCollegeDeptGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CollegeActiveCollegeDeptGroupRequest) GoString() string {
	return s.String()
}

func (s *CollegeActiveCollegeDeptGroupRequest) SetDeptId(v int64) *CollegeActiveCollegeDeptGroupRequest {
	s.DeptId = &v
	return s
}

type CollegeActiveCollegeDeptGroupResponseBody struct {
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s CollegeActiveCollegeDeptGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CollegeActiveCollegeDeptGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CollegeActiveCollegeDeptGroupResponseBody) SetOpenConversationId(v string) *CollegeActiveCollegeDeptGroupResponseBody {
	s.OpenConversationId = &v
	return s
}

type CollegeActiveCollegeDeptGroupResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CollegeActiveCollegeDeptGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CollegeActiveCollegeDeptGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CollegeActiveCollegeDeptGroupResponse) GoString() string {
	return s.String()
}

func (s *CollegeActiveCollegeDeptGroupResponse) SetHeaders(v map[string]*string) *CollegeActiveCollegeDeptGroupResponse {
	s.Headers = v
	return s
}

func (s *CollegeActiveCollegeDeptGroupResponse) SetStatusCode(v int32) *CollegeActiveCollegeDeptGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CollegeActiveCollegeDeptGroupResponse) SetBody(v *CollegeActiveCollegeDeptGroupResponseBody) *CollegeActiveCollegeDeptGroupResponse {
	s.Body = v
	return s
}

type CollegeAddCollegeDeptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CollegeAddCollegeDeptHeaders) String() string {
	return tea.Prettify(s)
}

func (s CollegeAddCollegeDeptHeaders) GoString() string {
	return s.String()
}

func (s *CollegeAddCollegeDeptHeaders) SetCommonHeaders(v map[string]*string) *CollegeAddCollegeDeptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CollegeAddCollegeDeptHeaders) SetXAcsDingtalkAccessToken(v string) *CollegeAddCollegeDeptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CollegeAddCollegeDeptRequest struct {
	DeptName   *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	DeptType   *string `json:"deptType,omitempty" xml:"deptType,omitempty"`
	SortFactor *int64  `json:"sortFactor,omitempty" xml:"sortFactor,omitempty"`
	SuperId    *int64  `json:"superId,omitempty" xml:"superId,omitempty"`
}

func (s CollegeAddCollegeDeptRequest) String() string {
	return tea.Prettify(s)
}

func (s CollegeAddCollegeDeptRequest) GoString() string {
	return s.String()
}

func (s *CollegeAddCollegeDeptRequest) SetDeptName(v string) *CollegeAddCollegeDeptRequest {
	s.DeptName = &v
	return s
}

func (s *CollegeAddCollegeDeptRequest) SetDeptType(v string) *CollegeAddCollegeDeptRequest {
	s.DeptType = &v
	return s
}

func (s *CollegeAddCollegeDeptRequest) SetSortFactor(v int64) *CollegeAddCollegeDeptRequest {
	s.SortFactor = &v
	return s
}

func (s *CollegeAddCollegeDeptRequest) SetSuperId(v int64) *CollegeAddCollegeDeptRequest {
	s.SuperId = &v
	return s
}

type CollegeAddCollegeDeptResponseBody struct {
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s CollegeAddCollegeDeptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CollegeAddCollegeDeptResponseBody) GoString() string {
	return s.String()
}

func (s *CollegeAddCollegeDeptResponseBody) SetDeptId(v int64) *CollegeAddCollegeDeptResponseBody {
	s.DeptId = &v
	return s
}

type CollegeAddCollegeDeptResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CollegeAddCollegeDeptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CollegeAddCollegeDeptResponse) String() string {
	return tea.Prettify(s)
}

func (s CollegeAddCollegeDeptResponse) GoString() string {
	return s.String()
}

func (s *CollegeAddCollegeDeptResponse) SetHeaders(v map[string]*string) *CollegeAddCollegeDeptResponse {
	s.Headers = v
	return s
}

func (s *CollegeAddCollegeDeptResponse) SetStatusCode(v int32) *CollegeAddCollegeDeptResponse {
	s.StatusCode = &v
	return s
}

func (s *CollegeAddCollegeDeptResponse) SetBody(v *CollegeAddCollegeDeptResponseBody) *CollegeAddCollegeDeptResponse {
	s.Body = v
	return s
}

type CollegeAddManagerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CollegeAddManagerHeaders) String() string {
	return tea.Prettify(s)
}

func (s CollegeAddManagerHeaders) GoString() string {
	return s.String()
}

func (s *CollegeAddManagerHeaders) SetCommonHeaders(v map[string]*string) *CollegeAddManagerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CollegeAddManagerHeaders) SetXAcsDingtalkAccessToken(v string) *CollegeAddManagerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CollegeAddManagerRequest struct {
	DeptId *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CollegeAddManagerRequest) String() string {
	return tea.Prettify(s)
}

func (s CollegeAddManagerRequest) GoString() string {
	return s.String()
}

func (s *CollegeAddManagerRequest) SetDeptId(v int64) *CollegeAddManagerRequest {
	s.DeptId = &v
	return s
}

func (s *CollegeAddManagerRequest) SetUserId(v string) *CollegeAddManagerRequest {
	s.UserId = &v
	return s
}

type CollegeAddManagerResponseBody struct {
	IsSuccessful *bool `json:"isSuccessful,omitempty" xml:"isSuccessful,omitempty"`
}

func (s CollegeAddManagerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CollegeAddManagerResponseBody) GoString() string {
	return s.String()
}

func (s *CollegeAddManagerResponseBody) SetIsSuccessful(v bool) *CollegeAddManagerResponseBody {
	s.IsSuccessful = &v
	return s
}

type CollegeAddManagerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CollegeAddManagerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CollegeAddManagerResponse) String() string {
	return tea.Prettify(s)
}

func (s CollegeAddManagerResponse) GoString() string {
	return s.String()
}

func (s *CollegeAddManagerResponse) SetHeaders(v map[string]*string) *CollegeAddManagerResponse {
	s.Headers = v
	return s
}

func (s *CollegeAddManagerResponse) SetStatusCode(v int32) *CollegeAddManagerResponse {
	s.StatusCode = &v
	return s
}

func (s *CollegeAddManagerResponse) SetBody(v *CollegeAddManagerResponseBody) *CollegeAddManagerResponse {
	s.Body = v
	return s
}

type CollegeAddStudentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CollegeAddStudentHeaders) String() string {
	return tea.Prettify(s)
}

func (s CollegeAddStudentHeaders) GoString() string {
	return s.String()
}

func (s *CollegeAddStudentHeaders) SetCommonHeaders(v map[string]*string) *CollegeAddStudentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CollegeAddStudentHeaders) SetXAcsDingtalkAccessToken(v string) *CollegeAddStudentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CollegeAddStudentRequest struct {
	DeptId        *int64             `json:"deptId,omitempty" xml:"deptId,omitempty"`
	EmpExtension  map[string]*string `json:"empExtension,omitempty" xml:"empExtension,omitempty"`
	Gender        *string            `json:"gender,omitempty" xml:"gender,omitempty"`
	IdentifyId    *string            `json:"identifyId,omitempty" xml:"identifyId,omitempty"`
	Mobile        *string            `json:"mobile,omitempty" xml:"mobile,omitempty"`
	StartYear     *string            `json:"startYear,omitempty" xml:"startYear,omitempty"`
	StudentName   *string            `json:"studentName,omitempty" xml:"studentName,omitempty"`
	StudentNumber *string            `json:"studentNumber,omitempty" xml:"studentNumber,omitempty"`
	UserId        *string            `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CollegeAddStudentRequest) String() string {
	return tea.Prettify(s)
}

func (s CollegeAddStudentRequest) GoString() string {
	return s.String()
}

func (s *CollegeAddStudentRequest) SetDeptId(v int64) *CollegeAddStudentRequest {
	s.DeptId = &v
	return s
}

func (s *CollegeAddStudentRequest) SetEmpExtension(v map[string]*string) *CollegeAddStudentRequest {
	s.EmpExtension = v
	return s
}

func (s *CollegeAddStudentRequest) SetGender(v string) *CollegeAddStudentRequest {
	s.Gender = &v
	return s
}

func (s *CollegeAddStudentRequest) SetIdentifyId(v string) *CollegeAddStudentRequest {
	s.IdentifyId = &v
	return s
}

func (s *CollegeAddStudentRequest) SetMobile(v string) *CollegeAddStudentRequest {
	s.Mobile = &v
	return s
}

func (s *CollegeAddStudentRequest) SetStartYear(v string) *CollegeAddStudentRequest {
	s.StartYear = &v
	return s
}

func (s *CollegeAddStudentRequest) SetStudentName(v string) *CollegeAddStudentRequest {
	s.StudentName = &v
	return s
}

func (s *CollegeAddStudentRequest) SetStudentNumber(v string) *CollegeAddStudentRequest {
	s.StudentNumber = &v
	return s
}

func (s *CollegeAddStudentRequest) SetUserId(v string) *CollegeAddStudentRequest {
	s.UserId = &v
	return s
}

type CollegeAddStudentResponseBody struct {
	DingMemberStatus *string `json:"dingMemberStatus,omitempty" xml:"dingMemberStatus,omitempty"`
	IsActive         *bool   `json:"isActive,omitempty" xml:"isActive,omitempty"`
	StudentId        *int64  `json:"studentId,omitempty" xml:"studentId,omitempty"`
	UnionId          *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId           *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CollegeAddStudentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CollegeAddStudentResponseBody) GoString() string {
	return s.String()
}

func (s *CollegeAddStudentResponseBody) SetDingMemberStatus(v string) *CollegeAddStudentResponseBody {
	s.DingMemberStatus = &v
	return s
}

func (s *CollegeAddStudentResponseBody) SetIsActive(v bool) *CollegeAddStudentResponseBody {
	s.IsActive = &v
	return s
}

func (s *CollegeAddStudentResponseBody) SetStudentId(v int64) *CollegeAddStudentResponseBody {
	s.StudentId = &v
	return s
}

func (s *CollegeAddStudentResponseBody) SetUnionId(v string) *CollegeAddStudentResponseBody {
	s.UnionId = &v
	return s
}

func (s *CollegeAddStudentResponseBody) SetUserId(v string) *CollegeAddStudentResponseBody {
	s.UserId = &v
	return s
}

type CollegeAddStudentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CollegeAddStudentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CollegeAddStudentResponse) String() string {
	return tea.Prettify(s)
}

func (s CollegeAddStudentResponse) GoString() string {
	return s.String()
}

func (s *CollegeAddStudentResponse) SetHeaders(v map[string]*string) *CollegeAddStudentResponse {
	s.Headers = v
	return s
}

func (s *CollegeAddStudentResponse) SetStatusCode(v int32) *CollegeAddStudentResponse {
	s.StatusCode = &v
	return s
}

func (s *CollegeAddStudentResponse) SetBody(v *CollegeAddStudentResponseBody) *CollegeAddStudentResponse {
	s.Body = v
	return s
}

type CollegeChangeStudentDeptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CollegeChangeStudentDeptHeaders) String() string {
	return tea.Prettify(s)
}

func (s CollegeChangeStudentDeptHeaders) GoString() string {
	return s.String()
}

func (s *CollegeChangeStudentDeptHeaders) SetCommonHeaders(v map[string]*string) *CollegeChangeStudentDeptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CollegeChangeStudentDeptHeaders) SetXAcsDingtalkAccessToken(v string) *CollegeChangeStudentDeptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CollegeChangeStudentDeptRequest struct {
	DeptId    *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	NewDeptId *int64 `json:"newDeptId,omitempty" xml:"newDeptId,omitempty"`
	StudentId *int64 `json:"studentId,omitempty" xml:"studentId,omitempty"`
}

func (s CollegeChangeStudentDeptRequest) String() string {
	return tea.Prettify(s)
}

func (s CollegeChangeStudentDeptRequest) GoString() string {
	return s.String()
}

func (s *CollegeChangeStudentDeptRequest) SetDeptId(v int64) *CollegeChangeStudentDeptRequest {
	s.DeptId = &v
	return s
}

func (s *CollegeChangeStudentDeptRequest) SetNewDeptId(v int64) *CollegeChangeStudentDeptRequest {
	s.NewDeptId = &v
	return s
}

func (s *CollegeChangeStudentDeptRequest) SetStudentId(v int64) *CollegeChangeStudentDeptRequest {
	s.StudentId = &v
	return s
}

type CollegeChangeStudentDeptResponseBody struct {
	IsSuccessful *bool `json:"isSuccessful,omitempty" xml:"isSuccessful,omitempty"`
}

func (s CollegeChangeStudentDeptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CollegeChangeStudentDeptResponseBody) GoString() string {
	return s.String()
}

func (s *CollegeChangeStudentDeptResponseBody) SetIsSuccessful(v bool) *CollegeChangeStudentDeptResponseBody {
	s.IsSuccessful = &v
	return s
}

type CollegeChangeStudentDeptResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CollegeChangeStudentDeptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CollegeChangeStudentDeptResponse) String() string {
	return tea.Prettify(s)
}

func (s CollegeChangeStudentDeptResponse) GoString() string {
	return s.String()
}

func (s *CollegeChangeStudentDeptResponse) SetHeaders(v map[string]*string) *CollegeChangeStudentDeptResponse {
	s.Headers = v
	return s
}

func (s *CollegeChangeStudentDeptResponse) SetStatusCode(v int32) *CollegeChangeStudentDeptResponse {
	s.StatusCode = &v
	return s
}

func (s *CollegeChangeStudentDeptResponse) SetBody(v *CollegeChangeStudentDeptResponseBody) *CollegeChangeStudentDeptResponse {
	s.Body = v
	return s
}

type CollegeDeleteCollegeDeptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CollegeDeleteCollegeDeptHeaders) String() string {
	return tea.Prettify(s)
}

func (s CollegeDeleteCollegeDeptHeaders) GoString() string {
	return s.String()
}

func (s *CollegeDeleteCollegeDeptHeaders) SetCommonHeaders(v map[string]*string) *CollegeDeleteCollegeDeptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CollegeDeleteCollegeDeptHeaders) SetXAcsDingtalkAccessToken(v string) *CollegeDeleteCollegeDeptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CollegeDeleteCollegeDeptRequest struct {
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s CollegeDeleteCollegeDeptRequest) String() string {
	return tea.Prettify(s)
}

func (s CollegeDeleteCollegeDeptRequest) GoString() string {
	return s.String()
}

func (s *CollegeDeleteCollegeDeptRequest) SetDeptId(v int64) *CollegeDeleteCollegeDeptRequest {
	s.DeptId = &v
	return s
}

type CollegeDeleteCollegeDeptResponseBody struct {
	IsSuccessful *bool `json:"isSuccessful,omitempty" xml:"isSuccessful,omitempty"`
}

func (s CollegeDeleteCollegeDeptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CollegeDeleteCollegeDeptResponseBody) GoString() string {
	return s.String()
}

func (s *CollegeDeleteCollegeDeptResponseBody) SetIsSuccessful(v bool) *CollegeDeleteCollegeDeptResponseBody {
	s.IsSuccessful = &v
	return s
}

type CollegeDeleteCollegeDeptResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CollegeDeleteCollegeDeptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CollegeDeleteCollegeDeptResponse) String() string {
	return tea.Prettify(s)
}

func (s CollegeDeleteCollegeDeptResponse) GoString() string {
	return s.String()
}

func (s *CollegeDeleteCollegeDeptResponse) SetHeaders(v map[string]*string) *CollegeDeleteCollegeDeptResponse {
	s.Headers = v
	return s
}

func (s *CollegeDeleteCollegeDeptResponse) SetStatusCode(v int32) *CollegeDeleteCollegeDeptResponse {
	s.StatusCode = &v
	return s
}

func (s *CollegeDeleteCollegeDeptResponse) SetBody(v *CollegeDeleteCollegeDeptResponseBody) *CollegeDeleteCollegeDeptResponse {
	s.Body = v
	return s
}

type CollegeListCollegeSubDeptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CollegeListCollegeSubDeptHeaders) String() string {
	return tea.Prettify(s)
}

func (s CollegeListCollegeSubDeptHeaders) GoString() string {
	return s.String()
}

func (s *CollegeListCollegeSubDeptHeaders) SetCommonHeaders(v map[string]*string) *CollegeListCollegeSubDeptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CollegeListCollegeSubDeptHeaders) SetXAcsDingtalkAccessToken(v string) *CollegeListCollegeSubDeptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CollegeListCollegeSubDeptRequest struct {
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s CollegeListCollegeSubDeptRequest) String() string {
	return tea.Prettify(s)
}

func (s CollegeListCollegeSubDeptRequest) GoString() string {
	return s.String()
}

func (s *CollegeListCollegeSubDeptRequest) SetDeptId(v int64) *CollegeListCollegeSubDeptRequest {
	s.DeptId = &v
	return s
}

type CollegeListCollegeSubDeptResponseBody struct {
	CollegeDeptInfoSimpleList []*CollegeListCollegeSubDeptResponseBodyCollegeDeptInfoSimpleList `json:"collegeDeptInfoSimpleList,omitempty" xml:"collegeDeptInfoSimpleList,omitempty" type:"Repeated"`
}

func (s CollegeListCollegeSubDeptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CollegeListCollegeSubDeptResponseBody) GoString() string {
	return s.String()
}

func (s *CollegeListCollegeSubDeptResponseBody) SetCollegeDeptInfoSimpleList(v []*CollegeListCollegeSubDeptResponseBodyCollegeDeptInfoSimpleList) *CollegeListCollegeSubDeptResponseBody {
	s.CollegeDeptInfoSimpleList = v
	return s
}

type CollegeListCollegeSubDeptResponseBodyCollegeDeptInfoSimpleList struct {
	DeptId   *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	DeptType *string `json:"deptType,omitempty" xml:"deptType,omitempty"`
}

func (s CollegeListCollegeSubDeptResponseBodyCollegeDeptInfoSimpleList) String() string {
	return tea.Prettify(s)
}

func (s CollegeListCollegeSubDeptResponseBodyCollegeDeptInfoSimpleList) GoString() string {
	return s.String()
}

func (s *CollegeListCollegeSubDeptResponseBodyCollegeDeptInfoSimpleList) SetDeptId(v int64) *CollegeListCollegeSubDeptResponseBodyCollegeDeptInfoSimpleList {
	s.DeptId = &v
	return s
}

func (s *CollegeListCollegeSubDeptResponseBodyCollegeDeptInfoSimpleList) SetDeptName(v string) *CollegeListCollegeSubDeptResponseBodyCollegeDeptInfoSimpleList {
	s.DeptName = &v
	return s
}

func (s *CollegeListCollegeSubDeptResponseBodyCollegeDeptInfoSimpleList) SetDeptType(v string) *CollegeListCollegeSubDeptResponseBodyCollegeDeptInfoSimpleList {
	s.DeptType = &v
	return s
}

type CollegeListCollegeSubDeptResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CollegeListCollegeSubDeptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CollegeListCollegeSubDeptResponse) String() string {
	return tea.Prettify(s)
}

func (s CollegeListCollegeSubDeptResponse) GoString() string {
	return s.String()
}

func (s *CollegeListCollegeSubDeptResponse) SetHeaders(v map[string]*string) *CollegeListCollegeSubDeptResponse {
	s.Headers = v
	return s
}

func (s *CollegeListCollegeSubDeptResponse) SetStatusCode(v int32) *CollegeListCollegeSubDeptResponse {
	s.StatusCode = &v
	return s
}

func (s *CollegeListCollegeSubDeptResponse) SetBody(v *CollegeListCollegeSubDeptResponseBody) *CollegeListCollegeSubDeptResponse {
	s.Body = v
	return s
}

type CollegeListDeptManagerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CollegeListDeptManagerHeaders) String() string {
	return tea.Prettify(s)
}

func (s CollegeListDeptManagerHeaders) GoString() string {
	return s.String()
}

func (s *CollegeListDeptManagerHeaders) SetCommonHeaders(v map[string]*string) *CollegeListDeptManagerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CollegeListDeptManagerHeaders) SetXAcsDingtalkAccessToken(v string) *CollegeListDeptManagerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CollegeListDeptManagerRequest struct {
	DeptId     *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s CollegeListDeptManagerRequest) String() string {
	return tea.Prettify(s)
}

func (s CollegeListDeptManagerRequest) GoString() string {
	return s.String()
}

func (s *CollegeListDeptManagerRequest) SetDeptId(v int64) *CollegeListDeptManagerRequest {
	s.DeptId = &v
	return s
}

func (s *CollegeListDeptManagerRequest) SetPageNumber(v int64) *CollegeListDeptManagerRequest {
	s.PageNumber = &v
	return s
}

func (s *CollegeListDeptManagerRequest) SetPageSize(v int64) *CollegeListDeptManagerRequest {
	s.PageSize = &v
	return s
}

type CollegeListDeptManagerResponseBody struct {
	ManagerInfoSimpleList []*CollegeListDeptManagerResponseBodyManagerInfoSimpleList `json:"managerInfoSimpleList,omitempty" xml:"managerInfoSimpleList,omitempty" type:"Repeated"`
	TotalCount            *int64                                                     `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s CollegeListDeptManagerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CollegeListDeptManagerResponseBody) GoString() string {
	return s.String()
}

func (s *CollegeListDeptManagerResponseBody) SetManagerInfoSimpleList(v []*CollegeListDeptManagerResponseBodyManagerInfoSimpleList) *CollegeListDeptManagerResponseBody {
	s.ManagerInfoSimpleList = v
	return s
}

func (s *CollegeListDeptManagerResponseBody) SetTotalCount(v int64) *CollegeListDeptManagerResponseBody {
	s.TotalCount = &v
	return s
}

type CollegeListDeptManagerResponseBodyManagerInfoSimpleList struct {
	IsActive *bool   `json:"isActive,omitempty" xml:"isActive,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId   *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CollegeListDeptManagerResponseBodyManagerInfoSimpleList) String() string {
	return tea.Prettify(s)
}

func (s CollegeListDeptManagerResponseBodyManagerInfoSimpleList) GoString() string {
	return s.String()
}

func (s *CollegeListDeptManagerResponseBodyManagerInfoSimpleList) SetIsActive(v bool) *CollegeListDeptManagerResponseBodyManagerInfoSimpleList {
	s.IsActive = &v
	return s
}

func (s *CollegeListDeptManagerResponseBodyManagerInfoSimpleList) SetName(v string) *CollegeListDeptManagerResponseBodyManagerInfoSimpleList {
	s.Name = &v
	return s
}

func (s *CollegeListDeptManagerResponseBodyManagerInfoSimpleList) SetUserId(v string) *CollegeListDeptManagerResponseBodyManagerInfoSimpleList {
	s.UserId = &v
	return s
}

type CollegeListDeptManagerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CollegeListDeptManagerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CollegeListDeptManagerResponse) String() string {
	return tea.Prettify(s)
}

func (s CollegeListDeptManagerResponse) GoString() string {
	return s.String()
}

func (s *CollegeListDeptManagerResponse) SetHeaders(v map[string]*string) *CollegeListDeptManagerResponse {
	s.Headers = v
	return s
}

func (s *CollegeListDeptManagerResponse) SetStatusCode(v int32) *CollegeListDeptManagerResponse {
	s.StatusCode = &v
	return s
}

func (s *CollegeListDeptManagerResponse) SetBody(v *CollegeListDeptManagerResponseBody) *CollegeListDeptManagerResponse {
	s.Body = v
	return s
}

type CollegeListStudentInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CollegeListStudentInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s CollegeListStudentInfoHeaders) GoString() string {
	return s.String()
}

func (s *CollegeListStudentInfoHeaders) SetCommonHeaders(v map[string]*string) *CollegeListStudentInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CollegeListStudentInfoHeaders) SetXAcsDingtalkAccessToken(v string) *CollegeListStudentInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CollegeListStudentInfoRequest struct {
	DeptId            *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	DingStudentStatus *string `json:"dingStudentStatus,omitempty" xml:"dingStudentStatus,omitempty"`
	PageNumber        *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize          *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s CollegeListStudentInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s CollegeListStudentInfoRequest) GoString() string {
	return s.String()
}

func (s *CollegeListStudentInfoRequest) SetDeptId(v int64) *CollegeListStudentInfoRequest {
	s.DeptId = &v
	return s
}

func (s *CollegeListStudentInfoRequest) SetDingStudentStatus(v string) *CollegeListStudentInfoRequest {
	s.DingStudentStatus = &v
	return s
}

func (s *CollegeListStudentInfoRequest) SetPageNumber(v int64) *CollegeListStudentInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *CollegeListStudentInfoRequest) SetPageSize(v int64) *CollegeListStudentInfoRequest {
	s.PageSize = &v
	return s
}

type CollegeListStudentInfoResponseBody struct {
	StudentInfoSimpleList []*CollegeListStudentInfoResponseBodyStudentInfoSimpleList `json:"studentInfoSimpleList,omitempty" xml:"studentInfoSimpleList,omitempty" type:"Repeated"`
	TotalCount            *int64                                                     `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s CollegeListStudentInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CollegeListStudentInfoResponseBody) GoString() string {
	return s.String()
}

func (s *CollegeListStudentInfoResponseBody) SetStudentInfoSimpleList(v []*CollegeListStudentInfoResponseBodyStudentInfoSimpleList) *CollegeListStudentInfoResponseBody {
	s.StudentInfoSimpleList = v
	return s
}

func (s *CollegeListStudentInfoResponseBody) SetTotalCount(v int64) *CollegeListStudentInfoResponseBody {
	s.TotalCount = &v
	return s
}

type CollegeListStudentInfoResponseBodyStudentInfoSimpleList struct {
	DingMemberStatus *string `json:"dingMemberStatus,omitempty" xml:"dingMemberStatus,omitempty"`
	IsActive         *bool   `json:"isActive,omitempty" xml:"isActive,omitempty"`
	StudentId        *int64  `json:"studentId,omitempty" xml:"studentId,omitempty"`
	StudentName      *string `json:"studentName,omitempty" xml:"studentName,omitempty"`
	UnionId          *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId           *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CollegeListStudentInfoResponseBodyStudentInfoSimpleList) String() string {
	return tea.Prettify(s)
}

func (s CollegeListStudentInfoResponseBodyStudentInfoSimpleList) GoString() string {
	return s.String()
}

func (s *CollegeListStudentInfoResponseBodyStudentInfoSimpleList) SetDingMemberStatus(v string) *CollegeListStudentInfoResponseBodyStudentInfoSimpleList {
	s.DingMemberStatus = &v
	return s
}

func (s *CollegeListStudentInfoResponseBodyStudentInfoSimpleList) SetIsActive(v bool) *CollegeListStudentInfoResponseBodyStudentInfoSimpleList {
	s.IsActive = &v
	return s
}

func (s *CollegeListStudentInfoResponseBodyStudentInfoSimpleList) SetStudentId(v int64) *CollegeListStudentInfoResponseBodyStudentInfoSimpleList {
	s.StudentId = &v
	return s
}

func (s *CollegeListStudentInfoResponseBodyStudentInfoSimpleList) SetStudentName(v string) *CollegeListStudentInfoResponseBodyStudentInfoSimpleList {
	s.StudentName = &v
	return s
}

func (s *CollegeListStudentInfoResponseBodyStudentInfoSimpleList) SetUnionId(v string) *CollegeListStudentInfoResponseBodyStudentInfoSimpleList {
	s.UnionId = &v
	return s
}

func (s *CollegeListStudentInfoResponseBodyStudentInfoSimpleList) SetUserId(v string) *CollegeListStudentInfoResponseBodyStudentInfoSimpleList {
	s.UserId = &v
	return s
}

type CollegeListStudentInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CollegeListStudentInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CollegeListStudentInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s CollegeListStudentInfoResponse) GoString() string {
	return s.String()
}

func (s *CollegeListStudentInfoResponse) SetHeaders(v map[string]*string) *CollegeListStudentInfoResponse {
	s.Headers = v
	return s
}

func (s *CollegeListStudentInfoResponse) SetStatusCode(v int32) *CollegeListStudentInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *CollegeListStudentInfoResponse) SetBody(v *CollegeListStudentInfoResponseBody) *CollegeListStudentInfoResponse {
	s.Body = v
	return s
}

type CollegeListUncheckedStudentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CollegeListUncheckedStudentHeaders) String() string {
	return tea.Prettify(s)
}

func (s CollegeListUncheckedStudentHeaders) GoString() string {
	return s.String()
}

func (s *CollegeListUncheckedStudentHeaders) SetCommonHeaders(v map[string]*string) *CollegeListUncheckedStudentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CollegeListUncheckedStudentHeaders) SetXAcsDingtalkAccessToken(v string) *CollegeListUncheckedStudentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CollegeListUncheckedStudentRequest struct {
	DeptId     *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s CollegeListUncheckedStudentRequest) String() string {
	return tea.Prettify(s)
}

func (s CollegeListUncheckedStudentRequest) GoString() string {
	return s.String()
}

func (s *CollegeListUncheckedStudentRequest) SetDeptId(v int64) *CollegeListUncheckedStudentRequest {
	s.DeptId = &v
	return s
}

func (s *CollegeListUncheckedStudentRequest) SetPageNumber(v int64) *CollegeListUncheckedStudentRequest {
	s.PageNumber = &v
	return s
}

func (s *CollegeListUncheckedStudentRequest) SetPageSize(v int64) *CollegeListUncheckedStudentRequest {
	s.PageSize = &v
	return s
}

type CollegeListUncheckedStudentResponseBody struct {
	StudentInfoSimpleList []*CollegeListUncheckedStudentResponseBodyStudentInfoSimpleList `json:"studentInfoSimpleList,omitempty" xml:"studentInfoSimpleList,omitempty" type:"Repeated"`
	TotalCount            *int64                                                          `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s CollegeListUncheckedStudentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CollegeListUncheckedStudentResponseBody) GoString() string {
	return s.String()
}

func (s *CollegeListUncheckedStudentResponseBody) SetStudentInfoSimpleList(v []*CollegeListUncheckedStudentResponseBodyStudentInfoSimpleList) *CollegeListUncheckedStudentResponseBody {
	s.StudentInfoSimpleList = v
	return s
}

func (s *CollegeListUncheckedStudentResponseBody) SetTotalCount(v int64) *CollegeListUncheckedStudentResponseBody {
	s.TotalCount = &v
	return s
}

type CollegeListUncheckedStudentResponseBodyStudentInfoSimpleList struct {
	DingMemberStatus *string `json:"dingMemberStatus,omitempty" xml:"dingMemberStatus,omitempty"`
	IsActive         *bool   `json:"isActive,omitempty" xml:"isActive,omitempty"`
	StudentId        *int64  `json:"studentId,omitempty" xml:"studentId,omitempty"`
	StudentName      *string `json:"studentName,omitempty" xml:"studentName,omitempty"`
	UnionId          *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId           *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CollegeListUncheckedStudentResponseBodyStudentInfoSimpleList) String() string {
	return tea.Prettify(s)
}

func (s CollegeListUncheckedStudentResponseBodyStudentInfoSimpleList) GoString() string {
	return s.String()
}

func (s *CollegeListUncheckedStudentResponseBodyStudentInfoSimpleList) SetDingMemberStatus(v string) *CollegeListUncheckedStudentResponseBodyStudentInfoSimpleList {
	s.DingMemberStatus = &v
	return s
}

func (s *CollegeListUncheckedStudentResponseBodyStudentInfoSimpleList) SetIsActive(v bool) *CollegeListUncheckedStudentResponseBodyStudentInfoSimpleList {
	s.IsActive = &v
	return s
}

func (s *CollegeListUncheckedStudentResponseBodyStudentInfoSimpleList) SetStudentId(v int64) *CollegeListUncheckedStudentResponseBodyStudentInfoSimpleList {
	s.StudentId = &v
	return s
}

func (s *CollegeListUncheckedStudentResponseBodyStudentInfoSimpleList) SetStudentName(v string) *CollegeListUncheckedStudentResponseBodyStudentInfoSimpleList {
	s.StudentName = &v
	return s
}

func (s *CollegeListUncheckedStudentResponseBodyStudentInfoSimpleList) SetUnionId(v string) *CollegeListUncheckedStudentResponseBodyStudentInfoSimpleList {
	s.UnionId = &v
	return s
}

func (s *CollegeListUncheckedStudentResponseBodyStudentInfoSimpleList) SetUserId(v string) *CollegeListUncheckedStudentResponseBodyStudentInfoSimpleList {
	s.UserId = &v
	return s
}

type CollegeListUncheckedStudentResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CollegeListUncheckedStudentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CollegeListUncheckedStudentResponse) String() string {
	return tea.Prettify(s)
}

func (s CollegeListUncheckedStudentResponse) GoString() string {
	return s.String()
}

func (s *CollegeListUncheckedStudentResponse) SetHeaders(v map[string]*string) *CollegeListUncheckedStudentResponse {
	s.Headers = v
	return s
}

func (s *CollegeListUncheckedStudentResponse) SetStatusCode(v int32) *CollegeListUncheckedStudentResponse {
	s.StatusCode = &v
	return s
}

func (s *CollegeListUncheckedStudentResponse) SetBody(v *CollegeListUncheckedStudentResponseBody) *CollegeListUncheckedStudentResponse {
	s.Body = v
	return s
}

type CollegeQueryCollegeDeptGroupInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CollegeQueryCollegeDeptGroupInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s CollegeQueryCollegeDeptGroupInfoHeaders) GoString() string {
	return s.String()
}

func (s *CollegeQueryCollegeDeptGroupInfoHeaders) SetCommonHeaders(v map[string]*string) *CollegeQueryCollegeDeptGroupInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CollegeQueryCollegeDeptGroupInfoHeaders) SetXAcsDingtalkAccessToken(v string) *CollegeQueryCollegeDeptGroupInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CollegeQueryCollegeDeptGroupInfoRequest struct {
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s CollegeQueryCollegeDeptGroupInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s CollegeQueryCollegeDeptGroupInfoRequest) GoString() string {
	return s.String()
}

func (s *CollegeQueryCollegeDeptGroupInfoRequest) SetDeptId(v int64) *CollegeQueryCollegeDeptGroupInfoRequest {
	s.DeptId = &v
	return s
}

type CollegeQueryCollegeDeptGroupInfoResponseBody struct {
	GroupName          *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
}

func (s CollegeQueryCollegeDeptGroupInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CollegeQueryCollegeDeptGroupInfoResponseBody) GoString() string {
	return s.String()
}

func (s *CollegeQueryCollegeDeptGroupInfoResponseBody) SetGroupName(v string) *CollegeQueryCollegeDeptGroupInfoResponseBody {
	s.GroupName = &v
	return s
}

func (s *CollegeQueryCollegeDeptGroupInfoResponseBody) SetOpenConversationId(v string) *CollegeQueryCollegeDeptGroupInfoResponseBody {
	s.OpenConversationId = &v
	return s
}

type CollegeQueryCollegeDeptGroupInfoResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CollegeQueryCollegeDeptGroupInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CollegeQueryCollegeDeptGroupInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s CollegeQueryCollegeDeptGroupInfoResponse) GoString() string {
	return s.String()
}

func (s *CollegeQueryCollegeDeptGroupInfoResponse) SetHeaders(v map[string]*string) *CollegeQueryCollegeDeptGroupInfoResponse {
	s.Headers = v
	return s
}

func (s *CollegeQueryCollegeDeptGroupInfoResponse) SetStatusCode(v int32) *CollegeQueryCollegeDeptGroupInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *CollegeQueryCollegeDeptGroupInfoResponse) SetBody(v *CollegeQueryCollegeDeptGroupInfoResponseBody) *CollegeQueryCollegeDeptGroupInfoResponse {
	s.Body = v
	return s
}

type CollegeQueryCollegeDeptInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CollegeQueryCollegeDeptInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s CollegeQueryCollegeDeptInfoHeaders) GoString() string {
	return s.String()
}

func (s *CollegeQueryCollegeDeptInfoHeaders) SetCommonHeaders(v map[string]*string) *CollegeQueryCollegeDeptInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CollegeQueryCollegeDeptInfoHeaders) SetXAcsDingtalkAccessToken(v string) *CollegeQueryCollegeDeptInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CollegeQueryCollegeDeptInfoRequest struct {
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s CollegeQueryCollegeDeptInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s CollegeQueryCollegeDeptInfoRequest) GoString() string {
	return s.String()
}

func (s *CollegeQueryCollegeDeptInfoRequest) SetDeptId(v int64) *CollegeQueryCollegeDeptInfoRequest {
	s.DeptId = &v
	return s
}

type CollegeQueryCollegeDeptInfoResponseBody struct {
	DeptId     *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	DeptName   *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	DeptType   *string `json:"deptType,omitempty" xml:"deptType,omitempty"`
	SortFactor *int64  `json:"sortFactor,omitempty" xml:"sortFactor,omitempty"`
	SuperId    *int64  `json:"superId,omitempty" xml:"superId,omitempty"`
}

func (s CollegeQueryCollegeDeptInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CollegeQueryCollegeDeptInfoResponseBody) GoString() string {
	return s.String()
}

func (s *CollegeQueryCollegeDeptInfoResponseBody) SetDeptId(v int64) *CollegeQueryCollegeDeptInfoResponseBody {
	s.DeptId = &v
	return s
}

func (s *CollegeQueryCollegeDeptInfoResponseBody) SetDeptName(v string) *CollegeQueryCollegeDeptInfoResponseBody {
	s.DeptName = &v
	return s
}

func (s *CollegeQueryCollegeDeptInfoResponseBody) SetDeptType(v string) *CollegeQueryCollegeDeptInfoResponseBody {
	s.DeptType = &v
	return s
}

func (s *CollegeQueryCollegeDeptInfoResponseBody) SetSortFactor(v int64) *CollegeQueryCollegeDeptInfoResponseBody {
	s.SortFactor = &v
	return s
}

func (s *CollegeQueryCollegeDeptInfoResponseBody) SetSuperId(v int64) *CollegeQueryCollegeDeptInfoResponseBody {
	s.SuperId = &v
	return s
}

type CollegeQueryCollegeDeptInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CollegeQueryCollegeDeptInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CollegeQueryCollegeDeptInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s CollegeQueryCollegeDeptInfoResponse) GoString() string {
	return s.String()
}

func (s *CollegeQueryCollegeDeptInfoResponse) SetHeaders(v map[string]*string) *CollegeQueryCollegeDeptInfoResponse {
	s.Headers = v
	return s
}

func (s *CollegeQueryCollegeDeptInfoResponse) SetStatusCode(v int32) *CollegeQueryCollegeDeptInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *CollegeQueryCollegeDeptInfoResponse) SetBody(v *CollegeQueryCollegeDeptInfoResponseBody) *CollegeQueryCollegeDeptInfoResponse {
	s.Body = v
	return s
}

type CollegeQueryStudentInfoByDeptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CollegeQueryStudentInfoByDeptHeaders) String() string {
	return tea.Prettify(s)
}

func (s CollegeQueryStudentInfoByDeptHeaders) GoString() string {
	return s.String()
}

func (s *CollegeQueryStudentInfoByDeptHeaders) SetCommonHeaders(v map[string]*string) *CollegeQueryStudentInfoByDeptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CollegeQueryStudentInfoByDeptHeaders) SetXAcsDingtalkAccessToken(v string) *CollegeQueryStudentInfoByDeptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CollegeQueryStudentInfoByDeptRequest struct {
	DeptId    *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	StudentId *int64 `json:"studentId,omitempty" xml:"studentId,omitempty"`
}

func (s CollegeQueryStudentInfoByDeptRequest) String() string {
	return tea.Prettify(s)
}

func (s CollegeQueryStudentInfoByDeptRequest) GoString() string {
	return s.String()
}

func (s *CollegeQueryStudentInfoByDeptRequest) SetDeptId(v int64) *CollegeQueryStudentInfoByDeptRequest {
	s.DeptId = &v
	return s
}

func (s *CollegeQueryStudentInfoByDeptRequest) SetStudentId(v int64) *CollegeQueryStudentInfoByDeptRequest {
	s.StudentId = &v
	return s
}

type CollegeQueryStudentInfoByDeptResponseBody struct {
	DeptId           *int64                 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	DingMemberStatus *string                `json:"dingMemberStatus,omitempty" xml:"dingMemberStatus,omitempty"`
	EmpExtension     map[string]interface{} `json:"empExtension,omitempty" xml:"empExtension,omitempty"`
	Gender           *string                `json:"gender,omitempty" xml:"gender,omitempty"`
	IdentifyId       *string                `json:"identifyId,omitempty" xml:"identifyId,omitempty"`
	IsActive         *bool                  `json:"isActive,omitempty" xml:"isActive,omitempty"`
	StartYear        *string                `json:"startYear,omitempty" xml:"startYear,omitempty"`
	StudentId        *int64                 `json:"studentId,omitempty" xml:"studentId,omitempty"`
	StudentName      *string                `json:"studentName,omitempty" xml:"studentName,omitempty"`
	StudentNumber    *string                `json:"studentNumber,omitempty" xml:"studentNumber,omitempty"`
	UnionId          *string                `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId           *string                `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CollegeQueryStudentInfoByDeptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CollegeQueryStudentInfoByDeptResponseBody) GoString() string {
	return s.String()
}

func (s *CollegeQueryStudentInfoByDeptResponseBody) SetDeptId(v int64) *CollegeQueryStudentInfoByDeptResponseBody {
	s.DeptId = &v
	return s
}

func (s *CollegeQueryStudentInfoByDeptResponseBody) SetDingMemberStatus(v string) *CollegeQueryStudentInfoByDeptResponseBody {
	s.DingMemberStatus = &v
	return s
}

func (s *CollegeQueryStudentInfoByDeptResponseBody) SetEmpExtension(v map[string]interface{}) *CollegeQueryStudentInfoByDeptResponseBody {
	s.EmpExtension = v
	return s
}

func (s *CollegeQueryStudentInfoByDeptResponseBody) SetGender(v string) *CollegeQueryStudentInfoByDeptResponseBody {
	s.Gender = &v
	return s
}

func (s *CollegeQueryStudentInfoByDeptResponseBody) SetIdentifyId(v string) *CollegeQueryStudentInfoByDeptResponseBody {
	s.IdentifyId = &v
	return s
}

func (s *CollegeQueryStudentInfoByDeptResponseBody) SetIsActive(v bool) *CollegeQueryStudentInfoByDeptResponseBody {
	s.IsActive = &v
	return s
}

func (s *CollegeQueryStudentInfoByDeptResponseBody) SetStartYear(v string) *CollegeQueryStudentInfoByDeptResponseBody {
	s.StartYear = &v
	return s
}

func (s *CollegeQueryStudentInfoByDeptResponseBody) SetStudentId(v int64) *CollegeQueryStudentInfoByDeptResponseBody {
	s.StudentId = &v
	return s
}

func (s *CollegeQueryStudentInfoByDeptResponseBody) SetStudentName(v string) *CollegeQueryStudentInfoByDeptResponseBody {
	s.StudentName = &v
	return s
}

func (s *CollegeQueryStudentInfoByDeptResponseBody) SetStudentNumber(v string) *CollegeQueryStudentInfoByDeptResponseBody {
	s.StudentNumber = &v
	return s
}

func (s *CollegeQueryStudentInfoByDeptResponseBody) SetUnionId(v string) *CollegeQueryStudentInfoByDeptResponseBody {
	s.UnionId = &v
	return s
}

func (s *CollegeQueryStudentInfoByDeptResponseBody) SetUserId(v string) *CollegeQueryStudentInfoByDeptResponseBody {
	s.UserId = &v
	return s
}

type CollegeQueryStudentInfoByDeptResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CollegeQueryStudentInfoByDeptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CollegeQueryStudentInfoByDeptResponse) String() string {
	return tea.Prettify(s)
}

func (s CollegeQueryStudentInfoByDeptResponse) GoString() string {
	return s.String()
}

func (s *CollegeQueryStudentInfoByDeptResponse) SetHeaders(v map[string]*string) *CollegeQueryStudentInfoByDeptResponse {
	s.Headers = v
	return s
}

func (s *CollegeQueryStudentInfoByDeptResponse) SetStatusCode(v int32) *CollegeQueryStudentInfoByDeptResponse {
	s.StatusCode = &v
	return s
}

func (s *CollegeQueryStudentInfoByDeptResponse) SetBody(v *CollegeQueryStudentInfoByDeptResponseBody) *CollegeQueryStudentInfoByDeptResponse {
	s.Body = v
	return s
}

type CollegeQueryStudentInfoByMobileHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CollegeQueryStudentInfoByMobileHeaders) String() string {
	return tea.Prettify(s)
}

func (s CollegeQueryStudentInfoByMobileHeaders) GoString() string {
	return s.String()
}

func (s *CollegeQueryStudentInfoByMobileHeaders) SetCommonHeaders(v map[string]*string) *CollegeQueryStudentInfoByMobileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CollegeQueryStudentInfoByMobileHeaders) SetXAcsDingtalkAccessToken(v string) *CollegeQueryStudentInfoByMobileHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CollegeQueryStudentInfoByMobileRequest struct {
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
}

func (s CollegeQueryStudentInfoByMobileRequest) String() string {
	return tea.Prettify(s)
}

func (s CollegeQueryStudentInfoByMobileRequest) GoString() string {
	return s.String()
}

func (s *CollegeQueryStudentInfoByMobileRequest) SetMobile(v string) *CollegeQueryStudentInfoByMobileRequest {
	s.Mobile = &v
	return s
}

type CollegeQueryStudentInfoByMobileResponseBody struct {
	DeptStudentInfoList []*CollegeQueryStudentInfoByMobileResponseBodyDeptStudentInfoList `json:"deptStudentInfoList,omitempty" xml:"deptStudentInfoList,omitempty" type:"Repeated"`
	DingMemberStatus    *string                                                           `json:"dingMemberStatus,omitempty" xml:"dingMemberStatus,omitempty"`
	EmpExtension        map[string]interface{}                                            `json:"empExtension,omitempty" xml:"empExtension,omitempty"`
	Gender              *string                                                           `json:"gender,omitempty" xml:"gender,omitempty"`
	IdentifyId          *string                                                           `json:"identifyId,omitempty" xml:"identifyId,omitempty"`
	IsActive            *bool                                                             `json:"isActive,omitempty" xml:"isActive,omitempty"`
	StartYear           *string                                                           `json:"startYear,omitempty" xml:"startYear,omitempty"`
	StudentId           *int64                                                            `json:"studentId,omitempty" xml:"studentId,omitempty"`
	StudentName         *string                                                           `json:"studentName,omitempty" xml:"studentName,omitempty"`
	UnionId             *string                                                           `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId              *string                                                           `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CollegeQueryStudentInfoByMobileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CollegeQueryStudentInfoByMobileResponseBody) GoString() string {
	return s.String()
}

func (s *CollegeQueryStudentInfoByMobileResponseBody) SetDeptStudentInfoList(v []*CollegeQueryStudentInfoByMobileResponseBodyDeptStudentInfoList) *CollegeQueryStudentInfoByMobileResponseBody {
	s.DeptStudentInfoList = v
	return s
}

func (s *CollegeQueryStudentInfoByMobileResponseBody) SetDingMemberStatus(v string) *CollegeQueryStudentInfoByMobileResponseBody {
	s.DingMemberStatus = &v
	return s
}

func (s *CollegeQueryStudentInfoByMobileResponseBody) SetEmpExtension(v map[string]interface{}) *CollegeQueryStudentInfoByMobileResponseBody {
	s.EmpExtension = v
	return s
}

func (s *CollegeQueryStudentInfoByMobileResponseBody) SetGender(v string) *CollegeQueryStudentInfoByMobileResponseBody {
	s.Gender = &v
	return s
}

func (s *CollegeQueryStudentInfoByMobileResponseBody) SetIdentifyId(v string) *CollegeQueryStudentInfoByMobileResponseBody {
	s.IdentifyId = &v
	return s
}

func (s *CollegeQueryStudentInfoByMobileResponseBody) SetIsActive(v bool) *CollegeQueryStudentInfoByMobileResponseBody {
	s.IsActive = &v
	return s
}

func (s *CollegeQueryStudentInfoByMobileResponseBody) SetStartYear(v string) *CollegeQueryStudentInfoByMobileResponseBody {
	s.StartYear = &v
	return s
}

func (s *CollegeQueryStudentInfoByMobileResponseBody) SetStudentId(v int64) *CollegeQueryStudentInfoByMobileResponseBody {
	s.StudentId = &v
	return s
}

func (s *CollegeQueryStudentInfoByMobileResponseBody) SetStudentName(v string) *CollegeQueryStudentInfoByMobileResponseBody {
	s.StudentName = &v
	return s
}

func (s *CollegeQueryStudentInfoByMobileResponseBody) SetUnionId(v string) *CollegeQueryStudentInfoByMobileResponseBody {
	s.UnionId = &v
	return s
}

func (s *CollegeQueryStudentInfoByMobileResponseBody) SetUserId(v string) *CollegeQueryStudentInfoByMobileResponseBody {
	s.UserId = &v
	return s
}

type CollegeQueryStudentInfoByMobileResponseBodyDeptStudentInfoList struct {
	DeptId        *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	MemberType    *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
	StudentNumber *string `json:"studentNumber,omitempty" xml:"studentNumber,omitempty"`
}

func (s CollegeQueryStudentInfoByMobileResponseBodyDeptStudentInfoList) String() string {
	return tea.Prettify(s)
}

func (s CollegeQueryStudentInfoByMobileResponseBodyDeptStudentInfoList) GoString() string {
	return s.String()
}

func (s *CollegeQueryStudentInfoByMobileResponseBodyDeptStudentInfoList) SetDeptId(v int64) *CollegeQueryStudentInfoByMobileResponseBodyDeptStudentInfoList {
	s.DeptId = &v
	return s
}

func (s *CollegeQueryStudentInfoByMobileResponseBodyDeptStudentInfoList) SetMemberType(v string) *CollegeQueryStudentInfoByMobileResponseBodyDeptStudentInfoList {
	s.MemberType = &v
	return s
}

func (s *CollegeQueryStudentInfoByMobileResponseBodyDeptStudentInfoList) SetStudentNumber(v string) *CollegeQueryStudentInfoByMobileResponseBodyDeptStudentInfoList {
	s.StudentNumber = &v
	return s
}

type CollegeQueryStudentInfoByMobileResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CollegeQueryStudentInfoByMobileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CollegeQueryStudentInfoByMobileResponse) String() string {
	return tea.Prettify(s)
}

func (s CollegeQueryStudentInfoByMobileResponse) GoString() string {
	return s.String()
}

func (s *CollegeQueryStudentInfoByMobileResponse) SetHeaders(v map[string]*string) *CollegeQueryStudentInfoByMobileResponse {
	s.Headers = v
	return s
}

func (s *CollegeQueryStudentInfoByMobileResponse) SetStatusCode(v int32) *CollegeQueryStudentInfoByMobileResponse {
	s.StatusCode = &v
	return s
}

func (s *CollegeQueryStudentInfoByMobileResponse) SetBody(v *CollegeQueryStudentInfoByMobileResponseBody) *CollegeQueryStudentInfoByMobileResponse {
	s.Body = v
	return s
}

type CollegeQueryStudentInfoByStudentIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CollegeQueryStudentInfoByStudentIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s CollegeQueryStudentInfoByStudentIdHeaders) GoString() string {
	return s.String()
}

func (s *CollegeQueryStudentInfoByStudentIdHeaders) SetCommonHeaders(v map[string]*string) *CollegeQueryStudentInfoByStudentIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CollegeQueryStudentInfoByStudentIdHeaders) SetXAcsDingtalkAccessToken(v string) *CollegeQueryStudentInfoByStudentIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CollegeQueryStudentInfoByStudentIdRequest struct {
	StudentId *int64 `json:"studentId,omitempty" xml:"studentId,omitempty"`
}

func (s CollegeQueryStudentInfoByStudentIdRequest) String() string {
	return tea.Prettify(s)
}

func (s CollegeQueryStudentInfoByStudentIdRequest) GoString() string {
	return s.String()
}

func (s *CollegeQueryStudentInfoByStudentIdRequest) SetStudentId(v int64) *CollegeQueryStudentInfoByStudentIdRequest {
	s.StudentId = &v
	return s
}

type CollegeQueryStudentInfoByStudentIdResponseBody struct {
	DeptStudentInfoList []*CollegeQueryStudentInfoByStudentIdResponseBodyDeptStudentInfoList `json:"deptStudentInfoList,omitempty" xml:"deptStudentInfoList,omitempty" type:"Repeated"`
	DingMemberStatus    *string                                                              `json:"dingMemberStatus,omitempty" xml:"dingMemberStatus,omitempty"`
	EmpExtension        map[string]interface{}                                               `json:"empExtension,omitempty" xml:"empExtension,omitempty"`
	Gender              *string                                                              `json:"gender,omitempty" xml:"gender,omitempty"`
	IdentifyId          *string                                                              `json:"identifyId,omitempty" xml:"identifyId,omitempty"`
	IsActive            *bool                                                                `json:"isActive,omitempty" xml:"isActive,omitempty"`
	StartYear           *string                                                              `json:"startYear,omitempty" xml:"startYear,omitempty"`
	StudentId           *int64                                                               `json:"studentId,omitempty" xml:"studentId,omitempty"`
	StudentName         *string                                                              `json:"studentName,omitempty" xml:"studentName,omitempty"`
	UnionId             *string                                                              `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId              *string                                                              `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CollegeQueryStudentInfoByStudentIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CollegeQueryStudentInfoByStudentIdResponseBody) GoString() string {
	return s.String()
}

func (s *CollegeQueryStudentInfoByStudentIdResponseBody) SetDeptStudentInfoList(v []*CollegeQueryStudentInfoByStudentIdResponseBodyDeptStudentInfoList) *CollegeQueryStudentInfoByStudentIdResponseBody {
	s.DeptStudentInfoList = v
	return s
}

func (s *CollegeQueryStudentInfoByStudentIdResponseBody) SetDingMemberStatus(v string) *CollegeQueryStudentInfoByStudentIdResponseBody {
	s.DingMemberStatus = &v
	return s
}

func (s *CollegeQueryStudentInfoByStudentIdResponseBody) SetEmpExtension(v map[string]interface{}) *CollegeQueryStudentInfoByStudentIdResponseBody {
	s.EmpExtension = v
	return s
}

func (s *CollegeQueryStudentInfoByStudentIdResponseBody) SetGender(v string) *CollegeQueryStudentInfoByStudentIdResponseBody {
	s.Gender = &v
	return s
}

func (s *CollegeQueryStudentInfoByStudentIdResponseBody) SetIdentifyId(v string) *CollegeQueryStudentInfoByStudentIdResponseBody {
	s.IdentifyId = &v
	return s
}

func (s *CollegeQueryStudentInfoByStudentIdResponseBody) SetIsActive(v bool) *CollegeQueryStudentInfoByStudentIdResponseBody {
	s.IsActive = &v
	return s
}

func (s *CollegeQueryStudentInfoByStudentIdResponseBody) SetStartYear(v string) *CollegeQueryStudentInfoByStudentIdResponseBody {
	s.StartYear = &v
	return s
}

func (s *CollegeQueryStudentInfoByStudentIdResponseBody) SetStudentId(v int64) *CollegeQueryStudentInfoByStudentIdResponseBody {
	s.StudentId = &v
	return s
}

func (s *CollegeQueryStudentInfoByStudentIdResponseBody) SetStudentName(v string) *CollegeQueryStudentInfoByStudentIdResponseBody {
	s.StudentName = &v
	return s
}

func (s *CollegeQueryStudentInfoByStudentIdResponseBody) SetUnionId(v string) *CollegeQueryStudentInfoByStudentIdResponseBody {
	s.UnionId = &v
	return s
}

func (s *CollegeQueryStudentInfoByStudentIdResponseBody) SetUserId(v string) *CollegeQueryStudentInfoByStudentIdResponseBody {
	s.UserId = &v
	return s
}

type CollegeQueryStudentInfoByStudentIdResponseBodyDeptStudentInfoList struct {
	DeptId        *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	MemberType    *string `json:"memberType,omitempty" xml:"memberType,omitempty"`
	StudentNumber *string `json:"studentNumber,omitempty" xml:"studentNumber,omitempty"`
}

func (s CollegeQueryStudentInfoByStudentIdResponseBodyDeptStudentInfoList) String() string {
	return tea.Prettify(s)
}

func (s CollegeQueryStudentInfoByStudentIdResponseBodyDeptStudentInfoList) GoString() string {
	return s.String()
}

func (s *CollegeQueryStudentInfoByStudentIdResponseBodyDeptStudentInfoList) SetDeptId(v int64) *CollegeQueryStudentInfoByStudentIdResponseBodyDeptStudentInfoList {
	s.DeptId = &v
	return s
}

func (s *CollegeQueryStudentInfoByStudentIdResponseBodyDeptStudentInfoList) SetMemberType(v string) *CollegeQueryStudentInfoByStudentIdResponseBodyDeptStudentInfoList {
	s.MemberType = &v
	return s
}

func (s *CollegeQueryStudentInfoByStudentIdResponseBodyDeptStudentInfoList) SetStudentNumber(v string) *CollegeQueryStudentInfoByStudentIdResponseBodyDeptStudentInfoList {
	s.StudentNumber = &v
	return s
}

type CollegeQueryStudentInfoByStudentIdResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CollegeQueryStudentInfoByStudentIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CollegeQueryStudentInfoByStudentIdResponse) String() string {
	return tea.Prettify(s)
}

func (s CollegeQueryStudentInfoByStudentIdResponse) GoString() string {
	return s.String()
}

func (s *CollegeQueryStudentInfoByStudentIdResponse) SetHeaders(v map[string]*string) *CollegeQueryStudentInfoByStudentIdResponse {
	s.Headers = v
	return s
}

func (s *CollegeQueryStudentInfoByStudentIdResponse) SetStatusCode(v int32) *CollegeQueryStudentInfoByStudentIdResponse {
	s.StatusCode = &v
	return s
}

func (s *CollegeQueryStudentInfoByStudentIdResponse) SetBody(v *CollegeQueryStudentInfoByStudentIdResponseBody) *CollegeQueryStudentInfoByStudentIdResponse {
	s.Body = v
	return s
}

type CollegeRemoveManagerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CollegeRemoveManagerHeaders) String() string {
	return tea.Prettify(s)
}

func (s CollegeRemoveManagerHeaders) GoString() string {
	return s.String()
}

func (s *CollegeRemoveManagerHeaders) SetCommonHeaders(v map[string]*string) *CollegeRemoveManagerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CollegeRemoveManagerHeaders) SetXAcsDingtalkAccessToken(v string) *CollegeRemoveManagerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CollegeRemoveManagerRequest struct {
	DeptId  *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	IsForce *bool   `json:"isForce,omitempty" xml:"isForce,omitempty"`
	UserId  *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CollegeRemoveManagerRequest) String() string {
	return tea.Prettify(s)
}

func (s CollegeRemoveManagerRequest) GoString() string {
	return s.String()
}

func (s *CollegeRemoveManagerRequest) SetDeptId(v int64) *CollegeRemoveManagerRequest {
	s.DeptId = &v
	return s
}

func (s *CollegeRemoveManagerRequest) SetIsForce(v bool) *CollegeRemoveManagerRequest {
	s.IsForce = &v
	return s
}

func (s *CollegeRemoveManagerRequest) SetUserId(v string) *CollegeRemoveManagerRequest {
	s.UserId = &v
	return s
}

type CollegeRemoveManagerResponseBody struct {
	IsSuccessful *bool `json:"isSuccessful,omitempty" xml:"isSuccessful,omitempty"`
}

func (s CollegeRemoveManagerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CollegeRemoveManagerResponseBody) GoString() string {
	return s.String()
}

func (s *CollegeRemoveManagerResponseBody) SetIsSuccessful(v bool) *CollegeRemoveManagerResponseBody {
	s.IsSuccessful = &v
	return s
}

type CollegeRemoveManagerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CollegeRemoveManagerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CollegeRemoveManagerResponse) String() string {
	return tea.Prettify(s)
}

func (s CollegeRemoveManagerResponse) GoString() string {
	return s.String()
}

func (s *CollegeRemoveManagerResponse) SetHeaders(v map[string]*string) *CollegeRemoveManagerResponse {
	s.Headers = v
	return s
}

func (s *CollegeRemoveManagerResponse) SetStatusCode(v int32) *CollegeRemoveManagerResponse {
	s.StatusCode = &v
	return s
}

func (s *CollegeRemoveManagerResponse) SetBody(v *CollegeRemoveManagerResponseBody) *CollegeRemoveManagerResponse {
	s.Body = v
	return s
}

type CollegeRemoveStudentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CollegeRemoveStudentHeaders) String() string {
	return tea.Prettify(s)
}

func (s CollegeRemoveStudentHeaders) GoString() string {
	return s.String()
}

func (s *CollegeRemoveStudentHeaders) SetCommonHeaders(v map[string]*string) *CollegeRemoveStudentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CollegeRemoveStudentHeaders) SetXAcsDingtalkAccessToken(v string) *CollegeRemoveStudentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CollegeRemoveStudentRequest struct {
	DeptId    *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	StudentId *int64 `json:"studentId,omitempty" xml:"studentId,omitempty"`
}

func (s CollegeRemoveStudentRequest) String() string {
	return tea.Prettify(s)
}

func (s CollegeRemoveStudentRequest) GoString() string {
	return s.String()
}

func (s *CollegeRemoveStudentRequest) SetDeptId(v int64) *CollegeRemoveStudentRequest {
	s.DeptId = &v
	return s
}

func (s *CollegeRemoveStudentRequest) SetStudentId(v int64) *CollegeRemoveStudentRequest {
	s.StudentId = &v
	return s
}

type CollegeRemoveStudentResponseBody struct {
	IsSuccessful *bool `json:"isSuccessful,omitempty" xml:"isSuccessful,omitempty"`
}

func (s CollegeRemoveStudentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CollegeRemoveStudentResponseBody) GoString() string {
	return s.String()
}

func (s *CollegeRemoveStudentResponseBody) SetIsSuccessful(v bool) *CollegeRemoveStudentResponseBody {
	s.IsSuccessful = &v
	return s
}

type CollegeRemoveStudentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CollegeRemoveStudentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CollegeRemoveStudentResponse) String() string {
	return tea.Prettify(s)
}

func (s CollegeRemoveStudentResponse) GoString() string {
	return s.String()
}

func (s *CollegeRemoveStudentResponse) SetHeaders(v map[string]*string) *CollegeRemoveStudentResponse {
	s.Headers = v
	return s
}

func (s *CollegeRemoveStudentResponse) SetStatusCode(v int32) *CollegeRemoveStudentResponse {
	s.StatusCode = &v
	return s
}

func (s *CollegeRemoveStudentResponse) SetBody(v *CollegeRemoveStudentResponseBody) *CollegeRemoveStudentResponse {
	s.Body = v
	return s
}

type CollegeUpdateCollegeDeptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CollegeUpdateCollegeDeptHeaders) String() string {
	return tea.Prettify(s)
}

func (s CollegeUpdateCollegeDeptHeaders) GoString() string {
	return s.String()
}

func (s *CollegeUpdateCollegeDeptHeaders) SetCommonHeaders(v map[string]*string) *CollegeUpdateCollegeDeptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CollegeUpdateCollegeDeptHeaders) SetXAcsDingtalkAccessToken(v string) *CollegeUpdateCollegeDeptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CollegeUpdateCollegeDeptRequest struct {
	DeptId     *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	DeptName   *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	SortFactor *int64  `json:"sortFactor,omitempty" xml:"sortFactor,omitempty"`
	SuperId    *int64  `json:"superId,omitempty" xml:"superId,omitempty"`
}

func (s CollegeUpdateCollegeDeptRequest) String() string {
	return tea.Prettify(s)
}

func (s CollegeUpdateCollegeDeptRequest) GoString() string {
	return s.String()
}

func (s *CollegeUpdateCollegeDeptRequest) SetDeptId(v int64) *CollegeUpdateCollegeDeptRequest {
	s.DeptId = &v
	return s
}

func (s *CollegeUpdateCollegeDeptRequest) SetDeptName(v string) *CollegeUpdateCollegeDeptRequest {
	s.DeptName = &v
	return s
}

func (s *CollegeUpdateCollegeDeptRequest) SetSortFactor(v int64) *CollegeUpdateCollegeDeptRequest {
	s.SortFactor = &v
	return s
}

func (s *CollegeUpdateCollegeDeptRequest) SetSuperId(v int64) *CollegeUpdateCollegeDeptRequest {
	s.SuperId = &v
	return s
}

type CollegeUpdateCollegeDeptResponseBody struct {
	IsSuccessful *bool `json:"isSuccessful,omitempty" xml:"isSuccessful,omitempty"`
}

func (s CollegeUpdateCollegeDeptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CollegeUpdateCollegeDeptResponseBody) GoString() string {
	return s.String()
}

func (s *CollegeUpdateCollegeDeptResponseBody) SetIsSuccessful(v bool) *CollegeUpdateCollegeDeptResponseBody {
	s.IsSuccessful = &v
	return s
}

type CollegeUpdateCollegeDeptResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CollegeUpdateCollegeDeptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CollegeUpdateCollegeDeptResponse) String() string {
	return tea.Prettify(s)
}

func (s CollegeUpdateCollegeDeptResponse) GoString() string {
	return s.String()
}

func (s *CollegeUpdateCollegeDeptResponse) SetHeaders(v map[string]*string) *CollegeUpdateCollegeDeptResponse {
	s.Headers = v
	return s
}

func (s *CollegeUpdateCollegeDeptResponse) SetStatusCode(v int32) *CollegeUpdateCollegeDeptResponse {
	s.StatusCode = &v
	return s
}

func (s *CollegeUpdateCollegeDeptResponse) SetBody(v *CollegeUpdateCollegeDeptResponseBody) *CollegeUpdateCollegeDeptResponse {
	s.Body = v
	return s
}

type CollegeUpdateStudentDeptInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CollegeUpdateStudentDeptInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s CollegeUpdateStudentDeptInfoHeaders) GoString() string {
	return s.String()
}

func (s *CollegeUpdateStudentDeptInfoHeaders) SetCommonHeaders(v map[string]*string) *CollegeUpdateStudentDeptInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CollegeUpdateStudentDeptInfoHeaders) SetXAcsDingtalkAccessToken(v string) *CollegeUpdateStudentDeptInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CollegeUpdateStudentDeptInfoRequest struct {
	DeptId        *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	StudentId     *int64  `json:"studentId,omitempty" xml:"studentId,omitempty"`
	StudentNumber *string `json:"studentNumber,omitempty" xml:"studentNumber,omitempty"`
}

func (s CollegeUpdateStudentDeptInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s CollegeUpdateStudentDeptInfoRequest) GoString() string {
	return s.String()
}

func (s *CollegeUpdateStudentDeptInfoRequest) SetDeptId(v int64) *CollegeUpdateStudentDeptInfoRequest {
	s.DeptId = &v
	return s
}

func (s *CollegeUpdateStudentDeptInfoRequest) SetStudentId(v int64) *CollegeUpdateStudentDeptInfoRequest {
	s.StudentId = &v
	return s
}

func (s *CollegeUpdateStudentDeptInfoRequest) SetStudentNumber(v string) *CollegeUpdateStudentDeptInfoRequest {
	s.StudentNumber = &v
	return s
}

type CollegeUpdateStudentDeptInfoResponseBody struct {
	IsSuccessful *bool `json:"isSuccessful,omitempty" xml:"isSuccessful,omitempty"`
}

func (s CollegeUpdateStudentDeptInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CollegeUpdateStudentDeptInfoResponseBody) GoString() string {
	return s.String()
}

func (s *CollegeUpdateStudentDeptInfoResponseBody) SetIsSuccessful(v bool) *CollegeUpdateStudentDeptInfoResponseBody {
	s.IsSuccessful = &v
	return s
}

type CollegeUpdateStudentDeptInfoResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CollegeUpdateStudentDeptInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CollegeUpdateStudentDeptInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s CollegeUpdateStudentDeptInfoResponse) GoString() string {
	return s.String()
}

func (s *CollegeUpdateStudentDeptInfoResponse) SetHeaders(v map[string]*string) *CollegeUpdateStudentDeptInfoResponse {
	s.Headers = v
	return s
}

func (s *CollegeUpdateStudentDeptInfoResponse) SetStatusCode(v int32) *CollegeUpdateStudentDeptInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *CollegeUpdateStudentDeptInfoResponse) SetBody(v *CollegeUpdateStudentDeptInfoResponseBody) *CollegeUpdateStudentDeptInfoResponse {
	s.Body = v
	return s
}

type CollegeUpdateStudentInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CollegeUpdateStudentInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s CollegeUpdateStudentInfoHeaders) GoString() string {
	return s.String()
}

func (s *CollegeUpdateStudentInfoHeaders) SetCommonHeaders(v map[string]*string) *CollegeUpdateStudentInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CollegeUpdateStudentInfoHeaders) SetXAcsDingtalkAccessToken(v string) *CollegeUpdateStudentInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CollegeUpdateStudentInfoRequest struct {
	EmpExtension map[string]*string `json:"empExtension,omitempty" xml:"empExtension,omitempty"`
	Gender       *string            `json:"gender,omitempty" xml:"gender,omitempty"`
	IdentifyId   *string            `json:"identifyId,omitempty" xml:"identifyId,omitempty"`
	StartYear    *string            `json:"startYear,omitempty" xml:"startYear,omitempty"`
	StudentId    *int64             `json:"studentId,omitempty" xml:"studentId,omitempty"`
	StudentName  *string            `json:"studentName,omitempty" xml:"studentName,omitempty"`
}

func (s CollegeUpdateStudentInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s CollegeUpdateStudentInfoRequest) GoString() string {
	return s.String()
}

func (s *CollegeUpdateStudentInfoRequest) SetEmpExtension(v map[string]*string) *CollegeUpdateStudentInfoRequest {
	s.EmpExtension = v
	return s
}

func (s *CollegeUpdateStudentInfoRequest) SetGender(v string) *CollegeUpdateStudentInfoRequest {
	s.Gender = &v
	return s
}

func (s *CollegeUpdateStudentInfoRequest) SetIdentifyId(v string) *CollegeUpdateStudentInfoRequest {
	s.IdentifyId = &v
	return s
}

func (s *CollegeUpdateStudentInfoRequest) SetStartYear(v string) *CollegeUpdateStudentInfoRequest {
	s.StartYear = &v
	return s
}

func (s *CollegeUpdateStudentInfoRequest) SetStudentId(v int64) *CollegeUpdateStudentInfoRequest {
	s.StudentId = &v
	return s
}

func (s *CollegeUpdateStudentInfoRequest) SetStudentName(v string) *CollegeUpdateStudentInfoRequest {
	s.StudentName = &v
	return s
}

type CollegeUpdateStudentInfoResponseBody struct {
	IsSuccessful *bool `json:"isSuccessful,omitempty" xml:"isSuccessful,omitempty"`
}

func (s CollegeUpdateStudentInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CollegeUpdateStudentInfoResponseBody) GoString() string {
	return s.String()
}

func (s *CollegeUpdateStudentInfoResponseBody) SetIsSuccessful(v bool) *CollegeUpdateStudentInfoResponseBody {
	s.IsSuccessful = &v
	return s
}

type CollegeUpdateStudentInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CollegeUpdateStudentInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CollegeUpdateStudentInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s CollegeUpdateStudentInfoResponse) GoString() string {
	return s.String()
}

func (s *CollegeUpdateStudentInfoResponse) SetHeaders(v map[string]*string) *CollegeUpdateStudentInfoResponse {
	s.Headers = v
	return s
}

func (s *CollegeUpdateStudentInfoResponse) SetStatusCode(v int32) *CollegeUpdateStudentInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *CollegeUpdateStudentInfoResponse) SetBody(v *CollegeUpdateStudentInfoResponseBody) *CollegeUpdateStudentInfoResponse {
	s.Body = v
	return s
}

type CollegeUpdateStudentMoblieHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CollegeUpdateStudentMoblieHeaders) String() string {
	return tea.Prettify(s)
}

func (s CollegeUpdateStudentMoblieHeaders) GoString() string {
	return s.String()
}

func (s *CollegeUpdateStudentMoblieHeaders) SetCommonHeaders(v map[string]*string) *CollegeUpdateStudentMoblieHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CollegeUpdateStudentMoblieHeaders) SetXAcsDingtalkAccessToken(v string) *CollegeUpdateStudentMoblieHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CollegeUpdateStudentMoblieRequest struct {
	IsForce   *bool   `json:"isForce,omitempty" xml:"isForce,omitempty"`
	NewMobile *string `json:"newMobile,omitempty" xml:"newMobile,omitempty"`
	StudentId *int64  `json:"studentId,omitempty" xml:"studentId,omitempty"`
}

func (s CollegeUpdateStudentMoblieRequest) String() string {
	return tea.Prettify(s)
}

func (s CollegeUpdateStudentMoblieRequest) GoString() string {
	return s.String()
}

func (s *CollegeUpdateStudentMoblieRequest) SetIsForce(v bool) *CollegeUpdateStudentMoblieRequest {
	s.IsForce = &v
	return s
}

func (s *CollegeUpdateStudentMoblieRequest) SetNewMobile(v string) *CollegeUpdateStudentMoblieRequest {
	s.NewMobile = &v
	return s
}

func (s *CollegeUpdateStudentMoblieRequest) SetStudentId(v int64) *CollegeUpdateStudentMoblieRequest {
	s.StudentId = &v
	return s
}

type CollegeUpdateStudentMoblieResponseBody struct {
	UpdateResult *string `json:"updateResult,omitempty" xml:"updateResult,omitempty"`
}

func (s CollegeUpdateStudentMoblieResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CollegeUpdateStudentMoblieResponseBody) GoString() string {
	return s.String()
}

func (s *CollegeUpdateStudentMoblieResponseBody) SetUpdateResult(v string) *CollegeUpdateStudentMoblieResponseBody {
	s.UpdateResult = &v
	return s
}

type CollegeUpdateStudentMoblieResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CollegeUpdateStudentMoblieResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CollegeUpdateStudentMoblieResponse) String() string {
	return tea.Prettify(s)
}

func (s CollegeUpdateStudentMoblieResponse) GoString() string {
	return s.String()
}

func (s *CollegeUpdateStudentMoblieResponse) SetHeaders(v map[string]*string) *CollegeUpdateStudentMoblieResponse {
	s.Headers = v
	return s
}

func (s *CollegeUpdateStudentMoblieResponse) SetStatusCode(v int32) *CollegeUpdateStudentMoblieResponse {
	s.StatusCode = &v
	return s
}

func (s *CollegeUpdateStudentMoblieResponse) SetBody(v *CollegeUpdateStudentMoblieResponseBody) *CollegeUpdateStudentMoblieResponse {
	s.Body = v
	return s
}

type CustomizeContactCreateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CustomizeContactCreateHeaders) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactCreateHeaders) GoString() string {
	return s.String()
}

func (s *CustomizeContactCreateHeaders) SetCommonHeaders(v map[string]*string) *CustomizeContactCreateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CustomizeContactCreateHeaders) SetXAcsDingtalkAccessToken(v string) *CustomizeContactCreateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CustomizeContactCreateRequest struct {
	ManagerIdList []*string `json:"managerIdList,omitempty" xml:"managerIdList,omitempty" type:"Repeated"`
	Name          *string   `json:"name,omitempty" xml:"name,omitempty"`
	Order         *int64    `json:"order,omitempty" xml:"order,omitempty"`
}

func (s CustomizeContactCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactCreateRequest) GoString() string {
	return s.String()
}

func (s *CustomizeContactCreateRequest) SetManagerIdList(v []*string) *CustomizeContactCreateRequest {
	s.ManagerIdList = v
	return s
}

func (s *CustomizeContactCreateRequest) SetName(v string) *CustomizeContactCreateRequest {
	s.Name = &v
	return s
}

func (s *CustomizeContactCreateRequest) SetOrder(v int64) *CustomizeContactCreateRequest {
	s.Order = &v
	return s
}

type CustomizeContactCreateResponseBody struct {
	Content *CustomizeContactCreateResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
}

func (s CustomizeContactCreateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactCreateResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeContactCreateResponseBody) SetContent(v *CustomizeContactCreateResponseBodyContent) *CustomizeContactCreateResponseBody {
	s.Content = v
	return s
}

type CustomizeContactCreateResponseBodyContent struct {
	Code       *string `json:"code,omitempty" xml:"code,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	Order      *int64  `json:"order,omitempty" xml:"order,omitempty"`
	RootDeptId *int64  `json:"rootDeptId,omitempty" xml:"rootDeptId,omitempty"`
}

func (s CustomizeContactCreateResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactCreateResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CustomizeContactCreateResponseBodyContent) SetCode(v string) *CustomizeContactCreateResponseBodyContent {
	s.Code = &v
	return s
}

func (s *CustomizeContactCreateResponseBodyContent) SetName(v string) *CustomizeContactCreateResponseBodyContent {
	s.Name = &v
	return s
}

func (s *CustomizeContactCreateResponseBodyContent) SetOrder(v int64) *CustomizeContactCreateResponseBodyContent {
	s.Order = &v
	return s
}

func (s *CustomizeContactCreateResponseBodyContent) SetRootDeptId(v int64) *CustomizeContactCreateResponseBodyContent {
	s.RootDeptId = &v
	return s
}

type CustomizeContactCreateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CustomizeContactCreateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeContactCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactCreateResponse) GoString() string {
	return s.String()
}

func (s *CustomizeContactCreateResponse) SetHeaders(v map[string]*string) *CustomizeContactCreateResponse {
	s.Headers = v
	return s
}

func (s *CustomizeContactCreateResponse) SetStatusCode(v int32) *CustomizeContactCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomizeContactCreateResponse) SetBody(v *CustomizeContactCreateResponseBody) *CustomizeContactCreateResponse {
	s.Body = v
	return s
}

type CustomizeContactDeleteHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CustomizeContactDeleteHeaders) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeleteHeaders) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeleteHeaders) SetCommonHeaders(v map[string]*string) *CustomizeContactDeleteHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CustomizeContactDeleteHeaders) SetXAcsDingtalkAccessToken(v string) *CustomizeContactDeleteHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CustomizeContactDeleteRequest struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s CustomizeContactDeleteRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeleteRequest) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeleteRequest) SetCode(v string) *CustomizeContactDeleteRequest {
	s.Code = &v
	return s
}

type CustomizeContactDeleteResponseBody struct {
	Content *bool `json:"content,omitempty" xml:"content,omitempty"`
}

func (s CustomizeContactDeleteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeleteResponseBody) SetContent(v bool) *CustomizeContactDeleteResponseBody {
	s.Content = &v
	return s
}

type CustomizeContactDeleteResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CustomizeContactDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeContactDeleteResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeleteResponse) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeleteResponse) SetHeaders(v map[string]*string) *CustomizeContactDeleteResponse {
	s.Headers = v
	return s
}

func (s *CustomizeContactDeleteResponse) SetStatusCode(v int32) *CustomizeContactDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomizeContactDeleteResponse) SetBody(v *CustomizeContactDeleteResponseBody) *CustomizeContactDeleteResponse {
	s.Body = v
	return s
}

type CustomizeContactDeptCreateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CustomizeContactDeptCreateHeaders) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptCreateHeaders) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptCreateHeaders) SetCommonHeaders(v map[string]*string) *CustomizeContactDeptCreateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CustomizeContactDeptCreateHeaders) SetXAcsDingtalkAccessToken(v string) *CustomizeContactDeptCreateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CustomizeContactDeptCreateRequest struct {
	Code          *string   `json:"code,omitempty" xml:"code,omitempty"`
	ManagerIdList []*string `json:"managerIdList,omitempty" xml:"managerIdList,omitempty" type:"Repeated"`
	Name          *string   `json:"name,omitempty" xml:"name,omitempty"`
	Order         *int64    `json:"order,omitempty" xml:"order,omitempty"`
	ParentDeptId  *int64    `json:"parentDeptId,omitempty" xml:"parentDeptId,omitempty"`
	RefId         *int64    `json:"refId,omitempty" xml:"refId,omitempty"`
	Type          *int64    `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CustomizeContactDeptCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptCreateRequest) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptCreateRequest) SetCode(v string) *CustomizeContactDeptCreateRequest {
	s.Code = &v
	return s
}

func (s *CustomizeContactDeptCreateRequest) SetManagerIdList(v []*string) *CustomizeContactDeptCreateRequest {
	s.ManagerIdList = v
	return s
}

func (s *CustomizeContactDeptCreateRequest) SetName(v string) *CustomizeContactDeptCreateRequest {
	s.Name = &v
	return s
}

func (s *CustomizeContactDeptCreateRequest) SetOrder(v int64) *CustomizeContactDeptCreateRequest {
	s.Order = &v
	return s
}

func (s *CustomizeContactDeptCreateRequest) SetParentDeptId(v int64) *CustomizeContactDeptCreateRequest {
	s.ParentDeptId = &v
	return s
}

func (s *CustomizeContactDeptCreateRequest) SetRefId(v int64) *CustomizeContactDeptCreateRequest {
	s.RefId = &v
	return s
}

func (s *CustomizeContactDeptCreateRequest) SetType(v int64) *CustomizeContactDeptCreateRequest {
	s.Type = &v
	return s
}

type CustomizeContactDeptCreateResponseBody struct {
	Content *int64 `json:"content,omitempty" xml:"content,omitempty"`
}

func (s CustomizeContactDeptCreateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptCreateResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptCreateResponseBody) SetContent(v int64) *CustomizeContactDeptCreateResponseBody {
	s.Content = &v
	return s
}

type CustomizeContactDeptCreateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CustomizeContactDeptCreateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeContactDeptCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptCreateResponse) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptCreateResponse) SetHeaders(v map[string]*string) *CustomizeContactDeptCreateResponse {
	s.Headers = v
	return s
}

func (s *CustomizeContactDeptCreateResponse) SetStatusCode(v int32) *CustomizeContactDeptCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomizeContactDeptCreateResponse) SetBody(v *CustomizeContactDeptCreateResponseBody) *CustomizeContactDeptCreateResponse {
	s.Body = v
	return s
}

type CustomizeContactDeptDeleteHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CustomizeContactDeptDeleteHeaders) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptDeleteHeaders) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptDeleteHeaders) SetCommonHeaders(v map[string]*string) *CustomizeContactDeptDeleteHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CustomizeContactDeptDeleteHeaders) SetXAcsDingtalkAccessToken(v string) *CustomizeContactDeptDeleteHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CustomizeContactDeptDeleteRequest struct {
	Code   *string `json:"code,omitempty" xml:"code,omitempty"`
	DeptId *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s CustomizeContactDeptDeleteRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptDeleteRequest) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptDeleteRequest) SetCode(v string) *CustomizeContactDeptDeleteRequest {
	s.Code = &v
	return s
}

func (s *CustomizeContactDeptDeleteRequest) SetDeptId(v int64) *CustomizeContactDeptDeleteRequest {
	s.DeptId = &v
	return s
}

type CustomizeContactDeptDeleteResponseBody struct {
	Content *bool `json:"content,omitempty" xml:"content,omitempty"`
}

func (s CustomizeContactDeptDeleteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptDeleteResponseBody) SetContent(v bool) *CustomizeContactDeptDeleteResponseBody {
	s.Content = &v
	return s
}

type CustomizeContactDeptDeleteResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CustomizeContactDeptDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeContactDeptDeleteResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptDeleteResponse) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptDeleteResponse) SetHeaders(v map[string]*string) *CustomizeContactDeptDeleteResponse {
	s.Headers = v
	return s
}

func (s *CustomizeContactDeptDeleteResponse) SetStatusCode(v int32) *CustomizeContactDeptDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomizeContactDeptDeleteResponse) SetBody(v *CustomizeContactDeptDeleteResponseBody) *CustomizeContactDeptDeleteResponse {
	s.Body = v
	return s
}

type CustomizeContactDeptGroupCreateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CustomizeContactDeptGroupCreateHeaders) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptGroupCreateHeaders) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptGroupCreateHeaders) SetCommonHeaders(v map[string]*string) *CustomizeContactDeptGroupCreateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CustomizeContactDeptGroupCreateHeaders) SetXAcsDingtalkAccessToken(v string) *CustomizeContactDeptGroupCreateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CustomizeContactDeptGroupCreateRequest struct {
	Code   *string `json:"code,omitempty" xml:"code,omitempty"`
	DeptId *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s CustomizeContactDeptGroupCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptGroupCreateRequest) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptGroupCreateRequest) SetCode(v string) *CustomizeContactDeptGroupCreateRequest {
	s.Code = &v
	return s
}

func (s *CustomizeContactDeptGroupCreateRequest) SetDeptId(v int64) *CustomizeContactDeptGroupCreateRequest {
	s.DeptId = &v
	return s
}

type CustomizeContactDeptGroupCreateResponseBody struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s CustomizeContactDeptGroupCreateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptGroupCreateResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptGroupCreateResponseBody) SetContent(v string) *CustomizeContactDeptGroupCreateResponseBody {
	s.Content = &v
	return s
}

type CustomizeContactDeptGroupCreateResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CustomizeContactDeptGroupCreateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeContactDeptGroupCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptGroupCreateResponse) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptGroupCreateResponse) SetHeaders(v map[string]*string) *CustomizeContactDeptGroupCreateResponse {
	s.Headers = v
	return s
}

func (s *CustomizeContactDeptGroupCreateResponse) SetStatusCode(v int32) *CustomizeContactDeptGroupCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomizeContactDeptGroupCreateResponse) SetBody(v *CustomizeContactDeptGroupCreateResponseBody) *CustomizeContactDeptGroupCreateResponse {
	s.Body = v
	return s
}

type CustomizeContactDeptInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CustomizeContactDeptInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptInfoHeaders) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptInfoHeaders) SetCommonHeaders(v map[string]*string) *CustomizeContactDeptInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CustomizeContactDeptInfoHeaders) SetXAcsDingtalkAccessToken(v string) *CustomizeContactDeptInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CustomizeContactDeptInfoRequest struct {
	Code   *string `json:"code,omitempty" xml:"code,omitempty"`
	DeptId *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s CustomizeContactDeptInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptInfoRequest) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptInfoRequest) SetCode(v string) *CustomizeContactDeptInfoRequest {
	s.Code = &v
	return s
}

func (s *CustomizeContactDeptInfoRequest) SetDeptId(v int64) *CustomizeContactDeptInfoRequest {
	s.DeptId = &v
	return s
}

type CustomizeContactDeptInfoResponseBody struct {
	Content *CustomizeContactDeptInfoResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
}

func (s CustomizeContactDeptInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptInfoResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptInfoResponseBody) SetContent(v *CustomizeContactDeptInfoResponseBodyContent) *CustomizeContactDeptInfoResponseBody {
	s.Content = v
	return s
}

type CustomizeContactDeptInfoResponseBodyContent struct {
	Code          *string   `json:"code,omitempty" xml:"code,omitempty"`
	Id            *int64    `json:"id,omitempty" xml:"id,omitempty"`
	ManagerIdList []*string `json:"managerIdList,omitempty" xml:"managerIdList,omitempty" type:"Repeated"`
	Name          *string   `json:"name,omitempty" xml:"name,omitempty"`
	Order         *int64    `json:"order,omitempty" xml:"order,omitempty"`
	ParentDeptId  *int64    `json:"parentDeptId,omitempty" xml:"parentDeptId,omitempty"`
	RefId         *int64    `json:"refId,omitempty" xml:"refId,omitempty"`
	Type          *int64    `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CustomizeContactDeptInfoResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptInfoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptInfoResponseBodyContent) SetCode(v string) *CustomizeContactDeptInfoResponseBodyContent {
	s.Code = &v
	return s
}

func (s *CustomizeContactDeptInfoResponseBodyContent) SetId(v int64) *CustomizeContactDeptInfoResponseBodyContent {
	s.Id = &v
	return s
}

func (s *CustomizeContactDeptInfoResponseBodyContent) SetManagerIdList(v []*string) *CustomizeContactDeptInfoResponseBodyContent {
	s.ManagerIdList = v
	return s
}

func (s *CustomizeContactDeptInfoResponseBodyContent) SetName(v string) *CustomizeContactDeptInfoResponseBodyContent {
	s.Name = &v
	return s
}

func (s *CustomizeContactDeptInfoResponseBodyContent) SetOrder(v int64) *CustomizeContactDeptInfoResponseBodyContent {
	s.Order = &v
	return s
}

func (s *CustomizeContactDeptInfoResponseBodyContent) SetParentDeptId(v int64) *CustomizeContactDeptInfoResponseBodyContent {
	s.ParentDeptId = &v
	return s
}

func (s *CustomizeContactDeptInfoResponseBodyContent) SetRefId(v int64) *CustomizeContactDeptInfoResponseBodyContent {
	s.RefId = &v
	return s
}

func (s *CustomizeContactDeptInfoResponseBodyContent) SetType(v int64) *CustomizeContactDeptInfoResponseBodyContent {
	s.Type = &v
	return s
}

type CustomizeContactDeptInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CustomizeContactDeptInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeContactDeptInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptInfoResponse) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptInfoResponse) SetHeaders(v map[string]*string) *CustomizeContactDeptInfoResponse {
	s.Headers = v
	return s
}

func (s *CustomizeContactDeptInfoResponse) SetStatusCode(v int32) *CustomizeContactDeptInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomizeContactDeptInfoResponse) SetBody(v *CustomizeContactDeptInfoResponseBody) *CustomizeContactDeptInfoResponse {
	s.Body = v
	return s
}

type CustomizeContactDeptListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CustomizeContactDeptListHeaders) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptListHeaders) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptListHeaders) SetCommonHeaders(v map[string]*string) *CustomizeContactDeptListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CustomizeContactDeptListHeaders) SetXAcsDingtalkAccessToken(v string) *CustomizeContactDeptListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CustomizeContactDeptListRequest struct {
	Code   *string `json:"code,omitempty" xml:"code,omitempty"`
	DeptId *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s CustomizeContactDeptListRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptListRequest) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptListRequest) SetCode(v string) *CustomizeContactDeptListRequest {
	s.Code = &v
	return s
}

func (s *CustomizeContactDeptListRequest) SetDeptId(v int64) *CustomizeContactDeptListRequest {
	s.DeptId = &v
	return s
}

type CustomizeContactDeptListResponseBody struct {
	Content []*CustomizeContactDeptListResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s CustomizeContactDeptListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptListResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptListResponseBody) SetContent(v []*CustomizeContactDeptListResponseBodyContent) *CustomizeContactDeptListResponseBody {
	s.Content = v
	return s
}

type CustomizeContactDeptListResponseBodyContent struct {
	Code          *string   `json:"code,omitempty" xml:"code,omitempty"`
	Id            *int64    `json:"id,omitempty" xml:"id,omitempty"`
	ManagerIdList []*string `json:"managerIdList,omitempty" xml:"managerIdList,omitempty" type:"Repeated"`
	Name          *string   `json:"name,omitempty" xml:"name,omitempty"`
	Order         *int64    `json:"order,omitempty" xml:"order,omitempty"`
	ParentDeptId  *int64    `json:"parentDeptId,omitempty" xml:"parentDeptId,omitempty"`
	RefId         *int64    `json:"refId,omitempty" xml:"refId,omitempty"`
	Type          *int64    `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CustomizeContactDeptListResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptListResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptListResponseBodyContent) SetCode(v string) *CustomizeContactDeptListResponseBodyContent {
	s.Code = &v
	return s
}

func (s *CustomizeContactDeptListResponseBodyContent) SetId(v int64) *CustomizeContactDeptListResponseBodyContent {
	s.Id = &v
	return s
}

func (s *CustomizeContactDeptListResponseBodyContent) SetManagerIdList(v []*string) *CustomizeContactDeptListResponseBodyContent {
	s.ManagerIdList = v
	return s
}

func (s *CustomizeContactDeptListResponseBodyContent) SetName(v string) *CustomizeContactDeptListResponseBodyContent {
	s.Name = &v
	return s
}

func (s *CustomizeContactDeptListResponseBodyContent) SetOrder(v int64) *CustomizeContactDeptListResponseBodyContent {
	s.Order = &v
	return s
}

func (s *CustomizeContactDeptListResponseBodyContent) SetParentDeptId(v int64) *CustomizeContactDeptListResponseBodyContent {
	s.ParentDeptId = &v
	return s
}

func (s *CustomizeContactDeptListResponseBodyContent) SetRefId(v int64) *CustomizeContactDeptListResponseBodyContent {
	s.RefId = &v
	return s
}

func (s *CustomizeContactDeptListResponseBodyContent) SetType(v int64) *CustomizeContactDeptListResponseBodyContent {
	s.Type = &v
	return s
}

type CustomizeContactDeptListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CustomizeContactDeptListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeContactDeptListResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptListResponse) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptListResponse) SetHeaders(v map[string]*string) *CustomizeContactDeptListResponse {
	s.Headers = v
	return s
}

func (s *CustomizeContactDeptListResponse) SetStatusCode(v int32) *CustomizeContactDeptListResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomizeContactDeptListResponse) SetBody(v *CustomizeContactDeptListResponseBody) *CustomizeContactDeptListResponse {
	s.Body = v
	return s
}

type CustomizeContactDeptUpdateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CustomizeContactDeptUpdateHeaders) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptUpdateHeaders) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptUpdateHeaders) SetCommonHeaders(v map[string]*string) *CustomizeContactDeptUpdateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CustomizeContactDeptUpdateHeaders) SetXAcsDingtalkAccessToken(v string) *CustomizeContactDeptUpdateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CustomizeContactDeptUpdateRequest struct {
	Code          *string   `json:"code,omitempty" xml:"code,omitempty"`
	DeptId        *int64    `json:"deptId,omitempty" xml:"deptId,omitempty"`
	ManagerIdList []*string `json:"managerIdList,omitempty" xml:"managerIdList,omitempty" type:"Repeated"`
	Name          *string   `json:"name,omitempty" xml:"name,omitempty"`
	Order         *int64    `json:"order,omitempty" xml:"order,omitempty"`
	ParentDeptId  *int64    `json:"parentDeptId,omitempty" xml:"parentDeptId,omitempty"`
}

func (s CustomizeContactDeptUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptUpdateRequest) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptUpdateRequest) SetCode(v string) *CustomizeContactDeptUpdateRequest {
	s.Code = &v
	return s
}

func (s *CustomizeContactDeptUpdateRequest) SetDeptId(v int64) *CustomizeContactDeptUpdateRequest {
	s.DeptId = &v
	return s
}

func (s *CustomizeContactDeptUpdateRequest) SetManagerIdList(v []*string) *CustomizeContactDeptUpdateRequest {
	s.ManagerIdList = v
	return s
}

func (s *CustomizeContactDeptUpdateRequest) SetName(v string) *CustomizeContactDeptUpdateRequest {
	s.Name = &v
	return s
}

func (s *CustomizeContactDeptUpdateRequest) SetOrder(v int64) *CustomizeContactDeptUpdateRequest {
	s.Order = &v
	return s
}

func (s *CustomizeContactDeptUpdateRequest) SetParentDeptId(v int64) *CustomizeContactDeptUpdateRequest {
	s.ParentDeptId = &v
	return s
}

type CustomizeContactDeptUpdateResponseBody struct {
	Content *int64 `json:"content,omitempty" xml:"content,omitempty"`
}

func (s CustomizeContactDeptUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptUpdateResponseBody) SetContent(v int64) *CustomizeContactDeptUpdateResponseBody {
	s.Content = &v
	return s
}

type CustomizeContactDeptUpdateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CustomizeContactDeptUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeContactDeptUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactDeptUpdateResponse) GoString() string {
	return s.String()
}

func (s *CustomizeContactDeptUpdateResponse) SetHeaders(v map[string]*string) *CustomizeContactDeptUpdateResponse {
	s.Headers = v
	return s
}

func (s *CustomizeContactDeptUpdateResponse) SetStatusCode(v int32) *CustomizeContactDeptUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomizeContactDeptUpdateResponse) SetBody(v *CustomizeContactDeptUpdateResponseBody) *CustomizeContactDeptUpdateResponse {
	s.Body = v
	return s
}

type CustomizeContactEmpAddHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CustomizeContactEmpAddHeaders) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactEmpAddHeaders) GoString() string {
	return s.String()
}

func (s *CustomizeContactEmpAddHeaders) SetCommonHeaders(v map[string]*string) *CustomizeContactEmpAddHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CustomizeContactEmpAddHeaders) SetXAcsDingtalkAccessToken(v string) *CustomizeContactEmpAddHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CustomizeContactEmpAddRequest struct {
	Code       *string   `json:"code,omitempty" xml:"code,omitempty"`
	DeptId     *int64    `json:"deptId,omitempty" xml:"deptId,omitempty"`
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s CustomizeContactEmpAddRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactEmpAddRequest) GoString() string {
	return s.String()
}

func (s *CustomizeContactEmpAddRequest) SetCode(v string) *CustomizeContactEmpAddRequest {
	s.Code = &v
	return s
}

func (s *CustomizeContactEmpAddRequest) SetDeptId(v int64) *CustomizeContactEmpAddRequest {
	s.DeptId = &v
	return s
}

func (s *CustomizeContactEmpAddRequest) SetUserIdList(v []*string) *CustomizeContactEmpAddRequest {
	s.UserIdList = v
	return s
}

type CustomizeContactEmpAddResponseBody struct {
	Content *bool `json:"content,omitempty" xml:"content,omitempty"`
}

func (s CustomizeContactEmpAddResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactEmpAddResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeContactEmpAddResponseBody) SetContent(v bool) *CustomizeContactEmpAddResponseBody {
	s.Content = &v
	return s
}

type CustomizeContactEmpAddResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CustomizeContactEmpAddResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeContactEmpAddResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactEmpAddResponse) GoString() string {
	return s.String()
}

func (s *CustomizeContactEmpAddResponse) SetHeaders(v map[string]*string) *CustomizeContactEmpAddResponse {
	s.Headers = v
	return s
}

func (s *CustomizeContactEmpAddResponse) SetStatusCode(v int32) *CustomizeContactEmpAddResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomizeContactEmpAddResponse) SetBody(v *CustomizeContactEmpAddResponseBody) *CustomizeContactEmpAddResponse {
	s.Body = v
	return s
}

type CustomizeContactEmpDeleteHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CustomizeContactEmpDeleteHeaders) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactEmpDeleteHeaders) GoString() string {
	return s.String()
}

func (s *CustomizeContactEmpDeleteHeaders) SetCommonHeaders(v map[string]*string) *CustomizeContactEmpDeleteHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CustomizeContactEmpDeleteHeaders) SetXAcsDingtalkAccessToken(v string) *CustomizeContactEmpDeleteHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CustomizeContactEmpDeleteRequest struct {
	Code       *string   `json:"code,omitempty" xml:"code,omitempty"`
	DeptId     *int64    `json:"deptId,omitempty" xml:"deptId,omitempty"`
	UserIdList []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s CustomizeContactEmpDeleteRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactEmpDeleteRequest) GoString() string {
	return s.String()
}

func (s *CustomizeContactEmpDeleteRequest) SetCode(v string) *CustomizeContactEmpDeleteRequest {
	s.Code = &v
	return s
}

func (s *CustomizeContactEmpDeleteRequest) SetDeptId(v int64) *CustomizeContactEmpDeleteRequest {
	s.DeptId = &v
	return s
}

func (s *CustomizeContactEmpDeleteRequest) SetUserIdList(v []*string) *CustomizeContactEmpDeleteRequest {
	s.UserIdList = v
	return s
}

type CustomizeContactEmpDeleteResponseBody struct {
	Content *bool `json:"content,omitempty" xml:"content,omitempty"`
}

func (s CustomizeContactEmpDeleteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactEmpDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeContactEmpDeleteResponseBody) SetContent(v bool) *CustomizeContactEmpDeleteResponseBody {
	s.Content = &v
	return s
}

type CustomizeContactEmpDeleteResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CustomizeContactEmpDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeContactEmpDeleteResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactEmpDeleteResponse) GoString() string {
	return s.String()
}

func (s *CustomizeContactEmpDeleteResponse) SetHeaders(v map[string]*string) *CustomizeContactEmpDeleteResponse {
	s.Headers = v
	return s
}

func (s *CustomizeContactEmpDeleteResponse) SetStatusCode(v int32) *CustomizeContactEmpDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomizeContactEmpDeleteResponse) SetBody(v *CustomizeContactEmpDeleteResponseBody) *CustomizeContactEmpDeleteResponse {
	s.Body = v
	return s
}

type CustomizeContactEmpListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CustomizeContactEmpListHeaders) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactEmpListHeaders) GoString() string {
	return s.String()
}

func (s *CustomizeContactEmpListHeaders) SetCommonHeaders(v map[string]*string) *CustomizeContactEmpListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CustomizeContactEmpListHeaders) SetXAcsDingtalkAccessToken(v string) *CustomizeContactEmpListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CustomizeContactEmpListRequest struct {
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s CustomizeContactEmpListRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactEmpListRequest) GoString() string {
	return s.String()
}

func (s *CustomizeContactEmpListRequest) SetDeptId(v int64) *CustomizeContactEmpListRequest {
	s.DeptId = &v
	return s
}

type CustomizeContactEmpListResponseBody struct {
	Content []*CustomizeContactEmpListResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s CustomizeContactEmpListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactEmpListResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeContactEmpListResponseBody) SetContent(v []*CustomizeContactEmpListResponseBodyContent) *CustomizeContactEmpListResponseBody {
	s.Content = v
	return s
}

type CustomizeContactEmpListResponseBodyContent struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CustomizeContactEmpListResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactEmpListResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CustomizeContactEmpListResponseBodyContent) SetName(v string) *CustomizeContactEmpListResponseBodyContent {
	s.Name = &v
	return s
}

func (s *CustomizeContactEmpListResponseBodyContent) SetUserId(v string) *CustomizeContactEmpListResponseBodyContent {
	s.UserId = &v
	return s
}

type CustomizeContactEmpListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CustomizeContactEmpListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeContactEmpListResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactEmpListResponse) GoString() string {
	return s.String()
}

func (s *CustomizeContactEmpListResponse) SetHeaders(v map[string]*string) *CustomizeContactEmpListResponse {
	s.Headers = v
	return s
}

func (s *CustomizeContactEmpListResponse) SetStatusCode(v int32) *CustomizeContactEmpListResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomizeContactEmpListResponse) SetBody(v *CustomizeContactEmpListResponseBody) *CustomizeContactEmpListResponse {
	s.Body = v
	return s
}

type CustomizeContactListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CustomizeContactListHeaders) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactListHeaders) GoString() string {
	return s.String()
}

func (s *CustomizeContactListHeaders) SetCommonHeaders(v map[string]*string) *CustomizeContactListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CustomizeContactListHeaders) SetXAcsDingtalkAccessToken(v string) *CustomizeContactListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CustomizeContactListResponseBody struct {
	Content []*CustomizeContactListResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s CustomizeContactListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactListResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeContactListResponseBody) SetContent(v []*CustomizeContactListResponseBodyContent) *CustomizeContactListResponseBody {
	s.Content = v
	return s
}

type CustomizeContactListResponseBodyContent struct {
	Code       *string `json:"code,omitempty" xml:"code,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	Order      *int64  `json:"order,omitempty" xml:"order,omitempty"`
	RootDeptId *int64  `json:"rootDeptId,omitempty" xml:"rootDeptId,omitempty"`
}

func (s CustomizeContactListResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactListResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CustomizeContactListResponseBodyContent) SetCode(v string) *CustomizeContactListResponseBodyContent {
	s.Code = &v
	return s
}

func (s *CustomizeContactListResponseBodyContent) SetName(v string) *CustomizeContactListResponseBodyContent {
	s.Name = &v
	return s
}

func (s *CustomizeContactListResponseBodyContent) SetOrder(v int64) *CustomizeContactListResponseBodyContent {
	s.Order = &v
	return s
}

func (s *CustomizeContactListResponseBodyContent) SetRootDeptId(v int64) *CustomizeContactListResponseBodyContent {
	s.RootDeptId = &v
	return s
}

type CustomizeContactListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CustomizeContactListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeContactListResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactListResponse) GoString() string {
	return s.String()
}

func (s *CustomizeContactListResponse) SetHeaders(v map[string]*string) *CustomizeContactListResponse {
	s.Headers = v
	return s
}

func (s *CustomizeContactListResponse) SetStatusCode(v int32) *CustomizeContactListResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomizeContactListResponse) SetBody(v *CustomizeContactListResponseBody) *CustomizeContactListResponse {
	s.Body = v
	return s
}

type CustomizeContactUpdateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CustomizeContactUpdateHeaders) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactUpdateHeaders) GoString() string {
	return s.String()
}

func (s *CustomizeContactUpdateHeaders) SetCommonHeaders(v map[string]*string) *CustomizeContactUpdateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CustomizeContactUpdateHeaders) SetXAcsDingtalkAccessToken(v string) *CustomizeContactUpdateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CustomizeContactUpdateRequest struct {
	Code          *string   `json:"code,omitempty" xml:"code,omitempty"`
	ManagerIdList []*string `json:"managerIdList,omitempty" xml:"managerIdList,omitempty" type:"Repeated"`
	Name          *string   `json:"name,omitempty" xml:"name,omitempty"`
	Order         *int64    `json:"order,omitempty" xml:"order,omitempty"`
}

func (s CustomizeContactUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactUpdateRequest) GoString() string {
	return s.String()
}

func (s *CustomizeContactUpdateRequest) SetCode(v string) *CustomizeContactUpdateRequest {
	s.Code = &v
	return s
}

func (s *CustomizeContactUpdateRequest) SetManagerIdList(v []*string) *CustomizeContactUpdateRequest {
	s.ManagerIdList = v
	return s
}

func (s *CustomizeContactUpdateRequest) SetName(v string) *CustomizeContactUpdateRequest {
	s.Name = &v
	return s
}

func (s *CustomizeContactUpdateRequest) SetOrder(v int64) *CustomizeContactUpdateRequest {
	s.Order = &v
	return s
}

type CustomizeContactUpdateResponseBody struct {
	Content *bool `json:"content,omitempty" xml:"content,omitempty"`
}

func (s CustomizeContactUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *CustomizeContactUpdateResponseBody) SetContent(v bool) *CustomizeContactUpdateResponseBody {
	s.Content = &v
	return s
}

type CustomizeContactUpdateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CustomizeContactUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CustomizeContactUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s CustomizeContactUpdateResponse) GoString() string {
	return s.String()
}

func (s *CustomizeContactUpdateResponse) SetHeaders(v map[string]*string) *CustomizeContactUpdateResponse {
	s.Headers = v
	return s
}

func (s *CustomizeContactUpdateResponse) SetStatusCode(v int32) *CustomizeContactUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomizeContactUpdateResponse) SetBody(v *CustomizeContactUpdateResponseBody) *CustomizeContactUpdateResponse {
	s.Body = v
	return s
}

type DIgitalStoreMessagePushHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DIgitalStoreMessagePushHeaders) String() string {
	return tea.Prettify(s)
}

func (s DIgitalStoreMessagePushHeaders) GoString() string {
	return s.String()
}

func (s *DIgitalStoreMessagePushHeaders) SetCommonHeaders(v map[string]*string) *DIgitalStoreMessagePushHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DIgitalStoreMessagePushHeaders) SetXAcsDingtalkAccessToken(v string) *DIgitalStoreMessagePushHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DIgitalStoreMessagePushRequest struct {
	MessageDataList []*DIgitalStoreMessagePushRequestMessageDataList `json:"messageDataList,omitempty" xml:"messageDataList,omitempty" type:"Repeated"`
}

func (s DIgitalStoreMessagePushRequest) String() string {
	return tea.Prettify(s)
}

func (s DIgitalStoreMessagePushRequest) GoString() string {
	return s.String()
}

func (s *DIgitalStoreMessagePushRequest) SetMessageDataList(v []*DIgitalStoreMessagePushRequestMessageDataList) *DIgitalStoreMessagePushRequest {
	s.MessageDataList = v
	return s
}

type DIgitalStoreMessagePushRequestMessageDataList struct {
	CallbackKey   *string `json:"callbackKey,omitempty" xml:"callbackKey,omitempty"`
	Content       *string `json:"content,omitempty" xml:"content,omitempty"`
	NewCard       *bool   `json:"newCard,omitempty" xml:"newCard,omitempty"`
	OutTraceId    *string `json:"outTraceId,omitempty" xml:"outTraceId,omitempty"`
	SceneCardCode *string `json:"sceneCardCode,omitempty" xml:"sceneCardCode,omitempty"`
	SceneScope    *int64  `json:"sceneScope,omitempty" xml:"sceneScope,omitempty"`
	SendNow       *bool   `json:"sendNow,omitempty" xml:"sendNow,omitempty"`
}

func (s DIgitalStoreMessagePushRequestMessageDataList) String() string {
	return tea.Prettify(s)
}

func (s DIgitalStoreMessagePushRequestMessageDataList) GoString() string {
	return s.String()
}

func (s *DIgitalStoreMessagePushRequestMessageDataList) SetCallbackKey(v string) *DIgitalStoreMessagePushRequestMessageDataList {
	s.CallbackKey = &v
	return s
}

func (s *DIgitalStoreMessagePushRequestMessageDataList) SetContent(v string) *DIgitalStoreMessagePushRequestMessageDataList {
	s.Content = &v
	return s
}

func (s *DIgitalStoreMessagePushRequestMessageDataList) SetNewCard(v bool) *DIgitalStoreMessagePushRequestMessageDataList {
	s.NewCard = &v
	return s
}

func (s *DIgitalStoreMessagePushRequestMessageDataList) SetOutTraceId(v string) *DIgitalStoreMessagePushRequestMessageDataList {
	s.OutTraceId = &v
	return s
}

func (s *DIgitalStoreMessagePushRequestMessageDataList) SetSceneCardCode(v string) *DIgitalStoreMessagePushRequestMessageDataList {
	s.SceneCardCode = &v
	return s
}

func (s *DIgitalStoreMessagePushRequestMessageDataList) SetSceneScope(v int64) *DIgitalStoreMessagePushRequestMessageDataList {
	s.SceneScope = &v
	return s
}

func (s *DIgitalStoreMessagePushRequestMessageDataList) SetSendNow(v bool) *DIgitalStoreMessagePushRequestMessageDataList {
	s.SendNow = &v
	return s
}

type DIgitalStoreMessagePushShrinkRequest struct {
	MessageDataListShrink *string `json:"messageDataList,omitempty" xml:"messageDataList,omitempty"`
}

func (s DIgitalStoreMessagePushShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DIgitalStoreMessagePushShrinkRequest) GoString() string {
	return s.String()
}

func (s *DIgitalStoreMessagePushShrinkRequest) SetMessageDataListShrink(v string) *DIgitalStoreMessagePushShrinkRequest {
	s.MessageDataListShrink = &v
	return s
}

type DIgitalStoreMessagePushResponseBody struct {
	Content *bool `json:"content,omitempty" xml:"content,omitempty"`
}

func (s DIgitalStoreMessagePushResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DIgitalStoreMessagePushResponseBody) GoString() string {
	return s.String()
}

func (s *DIgitalStoreMessagePushResponseBody) SetContent(v bool) *DIgitalStoreMessagePushResponseBody {
	s.Content = &v
	return s
}

type DIgitalStoreMessagePushResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DIgitalStoreMessagePushResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DIgitalStoreMessagePushResponse) String() string {
	return tea.Prettify(s)
}

func (s DIgitalStoreMessagePushResponse) GoString() string {
	return s.String()
}

func (s *DIgitalStoreMessagePushResponse) SetHeaders(v map[string]*string) *DIgitalStoreMessagePushResponse {
	s.Headers = v
	return s
}

func (s *DIgitalStoreMessagePushResponse) SetStatusCode(v int32) *DIgitalStoreMessagePushResponse {
	s.StatusCode = &v
	return s
}

func (s *DIgitalStoreMessagePushResponse) SetBody(v *DIgitalStoreMessagePushResponseBody) *DIgitalStoreMessagePushResponse {
	s.Body = v
	return s
}

type DigitalStoreContactInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DigitalStoreContactInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreContactInfoHeaders) GoString() string {
	return s.String()
}

func (s *DigitalStoreContactInfoHeaders) SetCommonHeaders(v map[string]*string) *DigitalStoreContactInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DigitalStoreContactInfoHeaders) SetXAcsDingtalkAccessToken(v string) *DigitalStoreContactInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DigitalStoreContactInfoResponseBody struct {
	Code       *string `json:"code,omitempty" xml:"code,omitempty"`
	DingDeptId *int64  `json:"dingDeptId,omitempty" xml:"dingDeptId,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	RootDeptId *int64  `json:"rootDeptId,omitempty" xml:"rootDeptId,omitempty"`
}

func (s DigitalStoreContactInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreContactInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DigitalStoreContactInfoResponseBody) SetCode(v string) *DigitalStoreContactInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DigitalStoreContactInfoResponseBody) SetDingDeptId(v int64) *DigitalStoreContactInfoResponseBody {
	s.DingDeptId = &v
	return s
}

func (s *DigitalStoreContactInfoResponseBody) SetName(v string) *DigitalStoreContactInfoResponseBody {
	s.Name = &v
	return s
}

func (s *DigitalStoreContactInfoResponseBody) SetRootDeptId(v int64) *DigitalStoreContactInfoResponseBody {
	s.RootDeptId = &v
	return s
}

type DigitalStoreContactInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DigitalStoreContactInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DigitalStoreContactInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreContactInfoResponse) GoString() string {
	return s.String()
}

func (s *DigitalStoreContactInfoResponse) SetHeaders(v map[string]*string) *DigitalStoreContactInfoResponse {
	s.Headers = v
	return s
}

func (s *DigitalStoreContactInfoResponse) SetStatusCode(v int32) *DigitalStoreContactInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DigitalStoreContactInfoResponse) SetBody(v *DigitalStoreContactInfoResponseBody) *DigitalStoreContactInfoResponse {
	s.Body = v
	return s
}

type DigitalStoreConversationsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DigitalStoreConversationsHeaders) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreConversationsHeaders) GoString() string {
	return s.String()
}

func (s *DigitalStoreConversationsHeaders) SetCommonHeaders(v map[string]*string) *DigitalStoreConversationsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DigitalStoreConversationsHeaders) SetXAcsDingtalkAccessToken(v string) *DigitalStoreConversationsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DigitalStoreConversationsRequest struct {
	ConversationTitle *string `json:"conversationTitle,omitempty" xml:"conversationTitle,omitempty"`
	ConversationType  *string `json:"conversationType,omitempty" xml:"conversationType,omitempty"`
}

func (s DigitalStoreConversationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreConversationsRequest) GoString() string {
	return s.String()
}

func (s *DigitalStoreConversationsRequest) SetConversationTitle(v string) *DigitalStoreConversationsRequest {
	s.ConversationTitle = &v
	return s
}

func (s *DigitalStoreConversationsRequest) SetConversationType(v string) *DigitalStoreConversationsRequest {
	s.ConversationType = &v
	return s
}

type DigitalStoreConversationsResponseBody struct {
	Content []*DigitalStoreConversationsResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s DigitalStoreConversationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreConversationsResponseBody) GoString() string {
	return s.String()
}

func (s *DigitalStoreConversationsResponseBody) SetContent(v []*DigitalStoreConversationsResponseBodyContent) *DigitalStoreConversationsResponseBody {
	s.Content = v
	return s
}

type DigitalStoreConversationsResponseBodyContent struct {
	ConversationTitle *string `json:"conversationTitle,omitempty" xml:"conversationTitle,omitempty"`
	ConversationType  *string `json:"conversationType,omitempty" xml:"conversationType,omitempty"`
	Id                *int64  `json:"id,omitempty" xml:"id,omitempty"`
}

func (s DigitalStoreConversationsResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreConversationsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DigitalStoreConversationsResponseBodyContent) SetConversationTitle(v string) *DigitalStoreConversationsResponseBodyContent {
	s.ConversationTitle = &v
	return s
}

func (s *DigitalStoreConversationsResponseBodyContent) SetConversationType(v string) *DigitalStoreConversationsResponseBodyContent {
	s.ConversationType = &v
	return s
}

func (s *DigitalStoreConversationsResponseBodyContent) SetId(v int64) *DigitalStoreConversationsResponseBodyContent {
	s.Id = &v
	return s
}

type DigitalStoreConversationsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DigitalStoreConversationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DigitalStoreConversationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreConversationsResponse) GoString() string {
	return s.String()
}

func (s *DigitalStoreConversationsResponse) SetHeaders(v map[string]*string) *DigitalStoreConversationsResponse {
	s.Headers = v
	return s
}

func (s *DigitalStoreConversationsResponse) SetStatusCode(v int32) *DigitalStoreConversationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DigitalStoreConversationsResponse) SetBody(v *DigitalStoreConversationsResponseBody) *DigitalStoreConversationsResponse {
	s.Body = v
	return s
}

type DigitalStoreGroupInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DigitalStoreGroupInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreGroupInfoHeaders) GoString() string {
	return s.String()
}

func (s *DigitalStoreGroupInfoHeaders) SetCommonHeaders(v map[string]*string) *DigitalStoreGroupInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DigitalStoreGroupInfoHeaders) SetXAcsDingtalkAccessToken(v string) *DigitalStoreGroupInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DigitalStoreGroupInfoRequest struct {
	GroupId *int64 `json:"groupId,omitempty" xml:"groupId,omitempty"`
}

func (s DigitalStoreGroupInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreGroupInfoRequest) GoString() string {
	return s.String()
}

func (s *DigitalStoreGroupInfoRequest) SetGroupId(v int64) *DigitalStoreGroupInfoRequest {
	s.GroupId = &v
	return s
}

type DigitalStoreGroupInfoResponseBody struct {
	GroupId     *int64   `json:"groupId,omitempty" xml:"groupId,omitempty"`
	GroupName   *string  `json:"groupName,omitempty" xml:"groupName,omitempty"`
	StoreIdList []*int64 `json:"storeIdList,omitempty" xml:"storeIdList,omitempty" type:"Repeated"`
}

func (s DigitalStoreGroupInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreGroupInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DigitalStoreGroupInfoResponseBody) SetGroupId(v int64) *DigitalStoreGroupInfoResponseBody {
	s.GroupId = &v
	return s
}

func (s *DigitalStoreGroupInfoResponseBody) SetGroupName(v string) *DigitalStoreGroupInfoResponseBody {
	s.GroupName = &v
	return s
}

func (s *DigitalStoreGroupInfoResponseBody) SetStoreIdList(v []*int64) *DigitalStoreGroupInfoResponseBody {
	s.StoreIdList = v
	return s
}

type DigitalStoreGroupInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DigitalStoreGroupInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DigitalStoreGroupInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreGroupInfoResponse) GoString() string {
	return s.String()
}

func (s *DigitalStoreGroupInfoResponse) SetHeaders(v map[string]*string) *DigitalStoreGroupInfoResponse {
	s.Headers = v
	return s
}

func (s *DigitalStoreGroupInfoResponse) SetStatusCode(v int32) *DigitalStoreGroupInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DigitalStoreGroupInfoResponse) SetBody(v *DigitalStoreGroupInfoResponseBody) *DigitalStoreGroupInfoResponse {
	s.Body = v
	return s
}

type DigitalStoreGroupsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DigitalStoreGroupsHeaders) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreGroupsHeaders) GoString() string {
	return s.String()
}

func (s *DigitalStoreGroupsHeaders) SetCommonHeaders(v map[string]*string) *DigitalStoreGroupsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DigitalStoreGroupsHeaders) SetXAcsDingtalkAccessToken(v string) *DigitalStoreGroupsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DigitalStoreGroupsResponseBody struct {
	Content []*DigitalStoreGroupsResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s DigitalStoreGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DigitalStoreGroupsResponseBody) SetContent(v []*DigitalStoreGroupsResponseBodyContent) *DigitalStoreGroupsResponseBody {
	s.Content = v
	return s
}

type DigitalStoreGroupsResponseBodyContent struct {
	GroupId   *int64  `json:"groupId,omitempty" xml:"groupId,omitempty"`
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
}

func (s DigitalStoreGroupsResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreGroupsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DigitalStoreGroupsResponseBodyContent) SetGroupId(v int64) *DigitalStoreGroupsResponseBodyContent {
	s.GroupId = &v
	return s
}

func (s *DigitalStoreGroupsResponseBodyContent) SetGroupName(v string) *DigitalStoreGroupsResponseBodyContent {
	s.GroupName = &v
	return s
}

type DigitalStoreGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DigitalStoreGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DigitalStoreGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreGroupsResponse) GoString() string {
	return s.String()
}

func (s *DigitalStoreGroupsResponse) SetHeaders(v map[string]*string) *DigitalStoreGroupsResponse {
	s.Headers = v
	return s
}

func (s *DigitalStoreGroupsResponse) SetStatusCode(v int32) *DigitalStoreGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DigitalStoreGroupsResponse) SetBody(v *DigitalStoreGroupsResponseBody) *DigitalStoreGroupsResponse {
	s.Body = v
	return s
}

type DigitalStoreNodeInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DigitalStoreNodeInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreNodeInfoHeaders) GoString() string {
	return s.String()
}

func (s *DigitalStoreNodeInfoHeaders) SetCommonHeaders(v map[string]*string) *DigitalStoreNodeInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DigitalStoreNodeInfoHeaders) SetXAcsDingtalkAccessToken(v string) *DigitalStoreNodeInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DigitalStoreNodeInfoRequest struct {
	Code   *string `json:"code,omitempty" xml:"code,omitempty"`
	NodeId *int64  `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
}

func (s DigitalStoreNodeInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreNodeInfoRequest) GoString() string {
	return s.String()
}

func (s *DigitalStoreNodeInfoRequest) SetCode(v string) *DigitalStoreNodeInfoRequest {
	s.Code = &v
	return s
}

func (s *DigitalStoreNodeInfoRequest) SetNodeId(v int64) *DigitalStoreNodeInfoRequest {
	s.NodeId = &v
	return s
}

type DigitalStoreNodeInfoResponseBody struct {
	DingDeptId *int64  `json:"dingDeptId,omitempty" xml:"dingDeptId,omitempty"`
	Id         *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	ParentId   *int64  `json:"parentId,omitempty" xml:"parentId,omitempty"`
	Type       *int64  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DigitalStoreNodeInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreNodeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DigitalStoreNodeInfoResponseBody) SetDingDeptId(v int64) *DigitalStoreNodeInfoResponseBody {
	s.DingDeptId = &v
	return s
}

func (s *DigitalStoreNodeInfoResponseBody) SetId(v int64) *DigitalStoreNodeInfoResponseBody {
	s.Id = &v
	return s
}

func (s *DigitalStoreNodeInfoResponseBody) SetName(v string) *DigitalStoreNodeInfoResponseBody {
	s.Name = &v
	return s
}

func (s *DigitalStoreNodeInfoResponseBody) SetParentId(v int64) *DigitalStoreNodeInfoResponseBody {
	s.ParentId = &v
	return s
}

func (s *DigitalStoreNodeInfoResponseBody) SetType(v int64) *DigitalStoreNodeInfoResponseBody {
	s.Type = &v
	return s
}

type DigitalStoreNodeInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DigitalStoreNodeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DigitalStoreNodeInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreNodeInfoResponse) GoString() string {
	return s.String()
}

func (s *DigitalStoreNodeInfoResponse) SetHeaders(v map[string]*string) *DigitalStoreNodeInfoResponse {
	s.Headers = v
	return s
}

func (s *DigitalStoreNodeInfoResponse) SetStatusCode(v int32) *DigitalStoreNodeInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DigitalStoreNodeInfoResponse) SetBody(v *DigitalStoreNodeInfoResponseBody) *DigitalStoreNodeInfoResponse {
	s.Body = v
	return s
}

type DigitalStoreRightsInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DigitalStoreRightsInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreRightsInfoHeaders) GoString() string {
	return s.String()
}

func (s *DigitalStoreRightsInfoHeaders) SetCommonHeaders(v map[string]*string) *DigitalStoreRightsInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DigitalStoreRightsInfoHeaders) SetXAcsDingtalkAccessToken(v string) *DigitalStoreRightsInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DigitalStoreRightsInfoResponseBody struct {
	EndTime    *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Quantity   *int64  `json:"quantity,omitempty" xml:"quantity,omitempty"`
	RightsCode *string `json:"rightsCode,omitempty" xml:"rightsCode,omitempty"`
	RightsName *string `json:"rightsName,omitempty" xml:"rightsName,omitempty"`
	StartTime  *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s DigitalStoreRightsInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreRightsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DigitalStoreRightsInfoResponseBody) SetEndTime(v int64) *DigitalStoreRightsInfoResponseBody {
	s.EndTime = &v
	return s
}

func (s *DigitalStoreRightsInfoResponseBody) SetQuantity(v int64) *DigitalStoreRightsInfoResponseBody {
	s.Quantity = &v
	return s
}

func (s *DigitalStoreRightsInfoResponseBody) SetRightsCode(v string) *DigitalStoreRightsInfoResponseBody {
	s.RightsCode = &v
	return s
}

func (s *DigitalStoreRightsInfoResponseBody) SetRightsName(v string) *DigitalStoreRightsInfoResponseBody {
	s.RightsName = &v
	return s
}

func (s *DigitalStoreRightsInfoResponseBody) SetStartTime(v int64) *DigitalStoreRightsInfoResponseBody {
	s.StartTime = &v
	return s
}

type DigitalStoreRightsInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DigitalStoreRightsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DigitalStoreRightsInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreRightsInfoResponse) GoString() string {
	return s.String()
}

func (s *DigitalStoreRightsInfoResponse) SetHeaders(v map[string]*string) *DigitalStoreRightsInfoResponse {
	s.Headers = v
	return s
}

func (s *DigitalStoreRightsInfoResponse) SetStatusCode(v int32) *DigitalStoreRightsInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DigitalStoreRightsInfoResponse) SetBody(v *DigitalStoreRightsInfoResponseBody) *DigitalStoreRightsInfoResponse {
	s.Body = v
	return s
}

type DigitalStoreRolesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DigitalStoreRolesHeaders) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreRolesHeaders) GoString() string {
	return s.String()
}

func (s *DigitalStoreRolesHeaders) SetCommonHeaders(v map[string]*string) *DigitalStoreRolesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DigitalStoreRolesHeaders) SetXAcsDingtalkAccessToken(v string) *DigitalStoreRolesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DigitalStoreRolesResponseBody struct {
	Content []*DigitalStoreRolesResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s DigitalStoreRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreRolesResponseBody) GoString() string {
	return s.String()
}

func (s *DigitalStoreRolesResponseBody) SetContent(v []*DigitalStoreRolesResponseBodyContent) *DigitalStoreRolesResponseBody {
	s.Content = v
	return s
}

type DigitalStoreRolesResponseBodyContent struct {
	Level    *int64  `json:"level,omitempty" xml:"level,omitempty"`
	RoleCode *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	RoleId   *int64  `json:"roleId,omitempty" xml:"roleId,omitempty"`
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
}

func (s DigitalStoreRolesResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreRolesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DigitalStoreRolesResponseBodyContent) SetLevel(v int64) *DigitalStoreRolesResponseBodyContent {
	s.Level = &v
	return s
}

func (s *DigitalStoreRolesResponseBodyContent) SetRoleCode(v string) *DigitalStoreRolesResponseBodyContent {
	s.RoleCode = &v
	return s
}

func (s *DigitalStoreRolesResponseBodyContent) SetRoleId(v int64) *DigitalStoreRolesResponseBodyContent {
	s.RoleId = &v
	return s
}

func (s *DigitalStoreRolesResponseBodyContent) SetRoleName(v string) *DigitalStoreRolesResponseBodyContent {
	s.RoleName = &v
	return s
}

type DigitalStoreRolesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DigitalStoreRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DigitalStoreRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreRolesResponse) GoString() string {
	return s.String()
}

func (s *DigitalStoreRolesResponse) SetHeaders(v map[string]*string) *DigitalStoreRolesResponse {
	s.Headers = v
	return s
}

func (s *DigitalStoreRolesResponse) SetStatusCode(v int32) *DigitalStoreRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *DigitalStoreRolesResponse) SetBody(v *DigitalStoreRolesResponseBody) *DigitalStoreRolesResponse {
	s.Body = v
	return s
}

type DigitalStoreSceneScopeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DigitalStoreSceneScopeHeaders) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreSceneScopeHeaders) GoString() string {
	return s.String()
}

func (s *DigitalStoreSceneScopeHeaders) SetCommonHeaders(v map[string]*string) *DigitalStoreSceneScopeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DigitalStoreSceneScopeHeaders) SetXAcsDingtalkAccessToken(v string) *DigitalStoreSceneScopeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DigitalStoreSceneScopeRequest struct {
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	SceneCode          *string `json:"sceneCode,omitempty" xml:"sceneCode,omitempty"`
}

func (s DigitalStoreSceneScopeRequest) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreSceneScopeRequest) GoString() string {
	return s.String()
}

func (s *DigitalStoreSceneScopeRequest) SetOpenConversationId(v string) *DigitalStoreSceneScopeRequest {
	s.OpenConversationId = &v
	return s
}

func (s *DigitalStoreSceneScopeRequest) SetSceneCode(v string) *DigitalStoreSceneScopeRequest {
	s.SceneCode = &v
	return s
}

type DigitalStoreSceneScopeResponseBody struct {
	GroupConversationType *string `json:"groupConversationType,omitempty" xml:"groupConversationType,omitempty"`
	ScopeId               *int64  `json:"scopeId,omitempty" xml:"scopeId,omitempty"`
}

func (s DigitalStoreSceneScopeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreSceneScopeResponseBody) GoString() string {
	return s.String()
}

func (s *DigitalStoreSceneScopeResponseBody) SetGroupConversationType(v string) *DigitalStoreSceneScopeResponseBody {
	s.GroupConversationType = &v
	return s
}

func (s *DigitalStoreSceneScopeResponseBody) SetScopeId(v int64) *DigitalStoreSceneScopeResponseBody {
	s.ScopeId = &v
	return s
}

type DigitalStoreSceneScopeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DigitalStoreSceneScopeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DigitalStoreSceneScopeResponse) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreSceneScopeResponse) GoString() string {
	return s.String()
}

func (s *DigitalStoreSceneScopeResponse) SetHeaders(v map[string]*string) *DigitalStoreSceneScopeResponse {
	s.Headers = v
	return s
}

func (s *DigitalStoreSceneScopeResponse) SetStatusCode(v int32) *DigitalStoreSceneScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *DigitalStoreSceneScopeResponse) SetBody(v *DigitalStoreSceneScopeResponseBody) *DigitalStoreSceneScopeResponse {
	s.Body = v
	return s
}

type DigitalStoreStoreInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DigitalStoreStoreInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreStoreInfoHeaders) GoString() string {
	return s.String()
}

func (s *DigitalStoreStoreInfoHeaders) SetCommonHeaders(v map[string]*string) *DigitalStoreStoreInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DigitalStoreStoreInfoHeaders) SetXAcsDingtalkAccessToken(v string) *DigitalStoreStoreInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DigitalStoreStoreInfoRequest struct {
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	StoreId *int64  `json:"storeId,omitempty" xml:"storeId,omitempty"`
}

func (s DigitalStoreStoreInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreStoreInfoRequest) GoString() string {
	return s.String()
}

func (s *DigitalStoreStoreInfoRequest) SetCode(v string) *DigitalStoreStoreInfoRequest {
	s.Code = &v
	return s
}

func (s *DigitalStoreStoreInfoRequest) SetStoreId(v int64) *DigitalStoreStoreInfoRequest {
	s.StoreId = &v
	return s
}

type DigitalStoreStoreInfoResponseBody struct {
	Address         *string `json:"address,omitempty" xml:"address,omitempty"`
	BusinessHours   *string `json:"businessHours,omitempty" xml:"businessHours,omitempty"`
	DingDeptId      *int64  `json:"dingDeptId,omitempty" xml:"dingDeptId,omitempty"`
	Latitude        *string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	LocationAddress *string `json:"locationAddress,omitempty" xml:"locationAddress,omitempty"`
	Longitude       *string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	ParentId        *int64  `json:"parentId,omitempty" xml:"parentId,omitempty"`
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
	StoreAcreage    *string `json:"storeAcreage,omitempty" xml:"storeAcreage,omitempty"`
	StoreBandwidth  *string `json:"storeBandwidth,omitempty" xml:"storeBandwidth,omitempty"`
	StoreCode       *string `json:"storeCode,omitempty" xml:"storeCode,omitempty"`
	StoreId         *int64  `json:"storeId,omitempty" xml:"storeId,omitempty"`
	Telephone       *string `json:"telephone,omitempty" xml:"telephone,omitempty"`
}

func (s DigitalStoreStoreInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreStoreInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DigitalStoreStoreInfoResponseBody) SetAddress(v string) *DigitalStoreStoreInfoResponseBody {
	s.Address = &v
	return s
}

func (s *DigitalStoreStoreInfoResponseBody) SetBusinessHours(v string) *DigitalStoreStoreInfoResponseBody {
	s.BusinessHours = &v
	return s
}

func (s *DigitalStoreStoreInfoResponseBody) SetDingDeptId(v int64) *DigitalStoreStoreInfoResponseBody {
	s.DingDeptId = &v
	return s
}

func (s *DigitalStoreStoreInfoResponseBody) SetLatitude(v string) *DigitalStoreStoreInfoResponseBody {
	s.Latitude = &v
	return s
}

func (s *DigitalStoreStoreInfoResponseBody) SetLocationAddress(v string) *DigitalStoreStoreInfoResponseBody {
	s.LocationAddress = &v
	return s
}

func (s *DigitalStoreStoreInfoResponseBody) SetLongitude(v string) *DigitalStoreStoreInfoResponseBody {
	s.Longitude = &v
	return s
}

func (s *DigitalStoreStoreInfoResponseBody) SetName(v string) *DigitalStoreStoreInfoResponseBody {
	s.Name = &v
	return s
}

func (s *DigitalStoreStoreInfoResponseBody) SetParentId(v int64) *DigitalStoreStoreInfoResponseBody {
	s.ParentId = &v
	return s
}

func (s *DigitalStoreStoreInfoResponseBody) SetStatus(v string) *DigitalStoreStoreInfoResponseBody {
	s.Status = &v
	return s
}

func (s *DigitalStoreStoreInfoResponseBody) SetStoreAcreage(v string) *DigitalStoreStoreInfoResponseBody {
	s.StoreAcreage = &v
	return s
}

func (s *DigitalStoreStoreInfoResponseBody) SetStoreBandwidth(v string) *DigitalStoreStoreInfoResponseBody {
	s.StoreBandwidth = &v
	return s
}

func (s *DigitalStoreStoreInfoResponseBody) SetStoreCode(v string) *DigitalStoreStoreInfoResponseBody {
	s.StoreCode = &v
	return s
}

func (s *DigitalStoreStoreInfoResponseBody) SetStoreId(v int64) *DigitalStoreStoreInfoResponseBody {
	s.StoreId = &v
	return s
}

func (s *DigitalStoreStoreInfoResponseBody) SetTelephone(v string) *DigitalStoreStoreInfoResponseBody {
	s.Telephone = &v
	return s
}

type DigitalStoreStoreInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DigitalStoreStoreInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DigitalStoreStoreInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreStoreInfoResponse) GoString() string {
	return s.String()
}

func (s *DigitalStoreStoreInfoResponse) SetHeaders(v map[string]*string) *DigitalStoreStoreInfoResponse {
	s.Headers = v
	return s
}

func (s *DigitalStoreStoreInfoResponse) SetStatusCode(v int32) *DigitalStoreStoreInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DigitalStoreStoreInfoResponse) SetBody(v *DigitalStoreStoreInfoResponseBody) *DigitalStoreStoreInfoResponse {
	s.Body = v
	return s
}

type DigitalStoreSubNodesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DigitalStoreSubNodesHeaders) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreSubNodesHeaders) GoString() string {
	return s.String()
}

func (s *DigitalStoreSubNodesHeaders) SetCommonHeaders(v map[string]*string) *DigitalStoreSubNodesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DigitalStoreSubNodesHeaders) SetXAcsDingtalkAccessToken(v string) *DigitalStoreSubNodesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DigitalStoreSubNodesRequest struct {
	Code   *string `json:"code,omitempty" xml:"code,omitempty"`
	NodeId *int64  `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
}

func (s DigitalStoreSubNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreSubNodesRequest) GoString() string {
	return s.String()
}

func (s *DigitalStoreSubNodesRequest) SetCode(v string) *DigitalStoreSubNodesRequest {
	s.Code = &v
	return s
}

func (s *DigitalStoreSubNodesRequest) SetNodeId(v int64) *DigitalStoreSubNodesRequest {
	s.NodeId = &v
	return s
}

type DigitalStoreSubNodesResponseBody struct {
	Content []*DigitalStoreSubNodesResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s DigitalStoreSubNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreSubNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DigitalStoreSubNodesResponseBody) SetContent(v []*DigitalStoreSubNodesResponseBodyContent) *DigitalStoreSubNodesResponseBody {
	s.Content = v
	return s
}

type DigitalStoreSubNodesResponseBodyContent struct {
	DingDeptId *int64  `json:"dingDeptId,omitempty" xml:"dingDeptId,omitempty"`
	Id         *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	ParentId   *int64  `json:"parentId,omitempty" xml:"parentId,omitempty"`
	Type       *int64  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DigitalStoreSubNodesResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreSubNodesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DigitalStoreSubNodesResponseBodyContent) SetDingDeptId(v int64) *DigitalStoreSubNodesResponseBodyContent {
	s.DingDeptId = &v
	return s
}

func (s *DigitalStoreSubNodesResponseBodyContent) SetId(v int64) *DigitalStoreSubNodesResponseBodyContent {
	s.Id = &v
	return s
}

func (s *DigitalStoreSubNodesResponseBodyContent) SetName(v string) *DigitalStoreSubNodesResponseBodyContent {
	s.Name = &v
	return s
}

func (s *DigitalStoreSubNodesResponseBodyContent) SetParentId(v int64) *DigitalStoreSubNodesResponseBodyContent {
	s.ParentId = &v
	return s
}

func (s *DigitalStoreSubNodesResponseBodyContent) SetType(v int64) *DigitalStoreSubNodesResponseBodyContent {
	s.Type = &v
	return s
}

type DigitalStoreSubNodesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DigitalStoreSubNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DigitalStoreSubNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreSubNodesResponse) GoString() string {
	return s.String()
}

func (s *DigitalStoreSubNodesResponse) SetHeaders(v map[string]*string) *DigitalStoreSubNodesResponse {
	s.Headers = v
	return s
}

func (s *DigitalStoreSubNodesResponse) SetStatusCode(v int32) *DigitalStoreSubNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DigitalStoreSubNodesResponse) SetBody(v *DigitalStoreSubNodesResponseBody) *DigitalStoreSubNodesResponse {
	s.Body = v
	return s
}

type DigitalStoreUpdateAuthInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DigitalStoreUpdateAuthInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreUpdateAuthInfoHeaders) GoString() string {
	return s.String()
}

func (s *DigitalStoreUpdateAuthInfoHeaders) SetCommonHeaders(v map[string]*string) *DigitalStoreUpdateAuthInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DigitalStoreUpdateAuthInfoHeaders) SetXAcsDingtalkAccessToken(v string) *DigitalStoreUpdateAuthInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DigitalStoreUpdateAuthInfoRequest struct {
	UpdateUserList []*DigitalStoreUpdateAuthInfoRequestUpdateUserList `json:"updateUserList,omitempty" xml:"updateUserList,omitempty" type:"Repeated"`
}

func (s DigitalStoreUpdateAuthInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreUpdateAuthInfoRequest) GoString() string {
	return s.String()
}

func (s *DigitalStoreUpdateAuthInfoRequest) SetUpdateUserList(v []*DigitalStoreUpdateAuthInfoRequestUpdateUserList) *DigitalStoreUpdateAuthInfoRequest {
	s.UpdateUserList = v
	return s
}

type DigitalStoreUpdateAuthInfoRequestUpdateUserList struct {
	RoleList     []*DigitalStoreUpdateAuthInfoRequestUpdateUserListRoleList     `json:"roleList,omitempty" xml:"roleList,omitempty" type:"Repeated"`
	UserAuthList []*DigitalStoreUpdateAuthInfoRequestUpdateUserListUserAuthList `json:"userAuthList,omitempty" xml:"userAuthList,omitempty" type:"Repeated"`
	UserId       *string                                                        `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DigitalStoreUpdateAuthInfoRequestUpdateUserList) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreUpdateAuthInfoRequestUpdateUserList) GoString() string {
	return s.String()
}

func (s *DigitalStoreUpdateAuthInfoRequestUpdateUserList) SetRoleList(v []*DigitalStoreUpdateAuthInfoRequestUpdateUserListRoleList) *DigitalStoreUpdateAuthInfoRequestUpdateUserList {
	s.RoleList = v
	return s
}

func (s *DigitalStoreUpdateAuthInfoRequestUpdateUserList) SetUserAuthList(v []*DigitalStoreUpdateAuthInfoRequestUpdateUserListUserAuthList) *DigitalStoreUpdateAuthInfoRequestUpdateUserList {
	s.UserAuthList = v
	return s
}

func (s *DigitalStoreUpdateAuthInfoRequestUpdateUserList) SetUserId(v string) *DigitalStoreUpdateAuthInfoRequestUpdateUserList {
	s.UserId = &v
	return s
}

type DigitalStoreUpdateAuthInfoRequestUpdateUserListRoleList struct {
	RoleName     *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	SourceRoleId *string `json:"sourceRoleId,omitempty" xml:"sourceRoleId,omitempty"`
}

func (s DigitalStoreUpdateAuthInfoRequestUpdateUserListRoleList) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreUpdateAuthInfoRequestUpdateUserListRoleList) GoString() string {
	return s.String()
}

func (s *DigitalStoreUpdateAuthInfoRequestUpdateUserListRoleList) SetRoleName(v string) *DigitalStoreUpdateAuthInfoRequestUpdateUserListRoleList {
	s.RoleName = &v
	return s
}

func (s *DigitalStoreUpdateAuthInfoRequestUpdateUserListRoleList) SetSourceRoleId(v string) *DigitalStoreUpdateAuthInfoRequestUpdateUserListRoleList {
	s.SourceRoleId = &v
	return s
}

type DigitalStoreUpdateAuthInfoRequestUpdateUserListUserAuthList struct {
	DingDeptId   *string `json:"dingDeptId,omitempty" xml:"dingDeptId,omitempty"`
	SourceDeptId *string `json:"sourceDeptId,omitempty" xml:"sourceDeptId,omitempty"`
}

func (s DigitalStoreUpdateAuthInfoRequestUpdateUserListUserAuthList) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreUpdateAuthInfoRequestUpdateUserListUserAuthList) GoString() string {
	return s.String()
}

func (s *DigitalStoreUpdateAuthInfoRequestUpdateUserListUserAuthList) SetDingDeptId(v string) *DigitalStoreUpdateAuthInfoRequestUpdateUserListUserAuthList {
	s.DingDeptId = &v
	return s
}

func (s *DigitalStoreUpdateAuthInfoRequestUpdateUserListUserAuthList) SetSourceDeptId(v string) *DigitalStoreUpdateAuthInfoRequestUpdateUserListUserAuthList {
	s.SourceDeptId = &v
	return s
}

type DigitalStoreUpdateAuthInfoResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DigitalStoreUpdateAuthInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreUpdateAuthInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DigitalStoreUpdateAuthInfoResponseBody) SetResult(v bool) *DigitalStoreUpdateAuthInfoResponseBody {
	s.Result = &v
	return s
}

type DigitalStoreUpdateAuthInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DigitalStoreUpdateAuthInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DigitalStoreUpdateAuthInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreUpdateAuthInfoResponse) GoString() string {
	return s.String()
}

func (s *DigitalStoreUpdateAuthInfoResponse) SetHeaders(v map[string]*string) *DigitalStoreUpdateAuthInfoResponse {
	s.Headers = v
	return s
}

func (s *DigitalStoreUpdateAuthInfoResponse) SetStatusCode(v int32) *DigitalStoreUpdateAuthInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DigitalStoreUpdateAuthInfoResponse) SetBody(v *DigitalStoreUpdateAuthInfoResponseBody) *DigitalStoreUpdateAuthInfoResponse {
	s.Body = v
	return s
}

type DigitalStoreUserInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DigitalStoreUserInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreUserInfoHeaders) GoString() string {
	return s.String()
}

func (s *DigitalStoreUserInfoHeaders) SetCommonHeaders(v map[string]*string) *DigitalStoreUserInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DigitalStoreUserInfoHeaders) SetXAcsDingtalkAccessToken(v string) *DigitalStoreUserInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DigitalStoreUserInfoRequest struct {
	Code   *string `json:"code,omitempty" xml:"code,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DigitalStoreUserInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreUserInfoRequest) GoString() string {
	return s.String()
}

func (s *DigitalStoreUserInfoRequest) SetCode(v string) *DigitalStoreUserInfoRequest {
	s.Code = &v
	return s
}

func (s *DigitalStoreUserInfoRequest) SetUserId(v string) *DigitalStoreUserInfoRequest {
	s.UserId = &v
	return s
}

type DigitalStoreUserInfoResponseBody struct {
	Name       *string  `json:"name,omitempty" xml:"name,omitempty"`
	RoleIdList []*int64 `json:"roleIdList,omitempty" xml:"roleIdList,omitempty" type:"Repeated"`
	ScopeList  []*int64 `json:"scopeList,omitempty" xml:"scopeList,omitempty" type:"Repeated"`
	StoreList  []*int64 `json:"storeList,omitempty" xml:"storeList,omitempty" type:"Repeated"`
	UserId     *string  `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DigitalStoreUserInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DigitalStoreUserInfoResponseBody) SetName(v string) *DigitalStoreUserInfoResponseBody {
	s.Name = &v
	return s
}

func (s *DigitalStoreUserInfoResponseBody) SetRoleIdList(v []*int64) *DigitalStoreUserInfoResponseBody {
	s.RoleIdList = v
	return s
}

func (s *DigitalStoreUserInfoResponseBody) SetScopeList(v []*int64) *DigitalStoreUserInfoResponseBody {
	s.ScopeList = v
	return s
}

func (s *DigitalStoreUserInfoResponseBody) SetStoreList(v []*int64) *DigitalStoreUserInfoResponseBody {
	s.StoreList = v
	return s
}

func (s *DigitalStoreUserInfoResponseBody) SetUserId(v string) *DigitalStoreUserInfoResponseBody {
	s.UserId = &v
	return s
}

type DigitalStoreUserInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DigitalStoreUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DigitalStoreUserInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreUserInfoResponse) GoString() string {
	return s.String()
}

func (s *DigitalStoreUserInfoResponse) SetHeaders(v map[string]*string) *DigitalStoreUserInfoResponse {
	s.Headers = v
	return s
}

func (s *DigitalStoreUserInfoResponse) SetStatusCode(v int32) *DigitalStoreUserInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DigitalStoreUserInfoResponse) SetBody(v *DigitalStoreUserInfoResponseBody) *DigitalStoreUserInfoResponse {
	s.Body = v
	return s
}

type DigitalStoreUsersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DigitalStoreUsersHeaders) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreUsersHeaders) GoString() string {
	return s.String()
}

func (s *DigitalStoreUsersHeaders) SetCommonHeaders(v map[string]*string) *DigitalStoreUsersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DigitalStoreUsersHeaders) SetXAcsDingtalkAccessToken(v string) *DigitalStoreUsersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DigitalStoreUsersRequest struct {
	Code   *string `json:"code,omitempty" xml:"code,omitempty"`
	NodeId *int64  `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
}

func (s DigitalStoreUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreUsersRequest) GoString() string {
	return s.String()
}

func (s *DigitalStoreUsersRequest) SetCode(v string) *DigitalStoreUsersRequest {
	s.Code = &v
	return s
}

func (s *DigitalStoreUsersRequest) SetNodeId(v int64) *DigitalStoreUsersRequest {
	s.NodeId = &v
	return s
}

type DigitalStoreUsersResponseBody struct {
	Content []*DigitalStoreUsersResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s DigitalStoreUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DigitalStoreUsersResponseBody) SetContent(v []*DigitalStoreUsersResponseBodyContent) *DigitalStoreUsersResponseBody {
	s.Content = v
	return s
}

type DigitalStoreUsersResponseBodyContent struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DigitalStoreUsersResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreUsersResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DigitalStoreUsersResponseBodyContent) SetName(v string) *DigitalStoreUsersResponseBodyContent {
	s.Name = &v
	return s
}

func (s *DigitalStoreUsersResponseBodyContent) SetUserId(v string) *DigitalStoreUsersResponseBodyContent {
	s.UserId = &v
	return s
}

type DigitalStoreUsersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DigitalStoreUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DigitalStoreUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s DigitalStoreUsersResponse) GoString() string {
	return s.String()
}

func (s *DigitalStoreUsersResponse) SetHeaders(v map[string]*string) *DigitalStoreUsersResponse {
	s.Headers = v
	return s
}

func (s *DigitalStoreUsersResponse) SetStatusCode(v int32) *DigitalStoreUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DigitalStoreUsersResponse) SetBody(v *DigitalStoreUsersResponseBody) *DigitalStoreUsersResponse {
	s.Body = v
	return s
}

type ExternalQueryExternalAppOrgsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ExternalQueryExternalAppOrgsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ExternalQueryExternalAppOrgsHeaders) GoString() string {
	return s.String()
}

func (s *ExternalQueryExternalAppOrgsHeaders) SetCommonHeaders(v map[string]*string) *ExternalQueryExternalAppOrgsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ExternalQueryExternalAppOrgsHeaders) SetXAcsDingtalkAccessToken(v string) *ExternalQueryExternalAppOrgsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ExternalQueryExternalAppOrgsRequest struct {
	ExternalType *string `json:"externalType,omitempty" xml:"externalType,omitempty"`
}

func (s ExternalQueryExternalAppOrgsRequest) String() string {
	return tea.Prettify(s)
}

func (s ExternalQueryExternalAppOrgsRequest) GoString() string {
	return s.String()
}

func (s *ExternalQueryExternalAppOrgsRequest) SetExternalType(v string) *ExternalQueryExternalAppOrgsRequest {
	s.ExternalType = &v
	return s
}

type ExternalQueryExternalAppOrgsResponseBody struct {
	Result []*ExternalQueryExternalAppOrgsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ExternalQueryExternalAppOrgsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExternalQueryExternalAppOrgsResponseBody) GoString() string {
	return s.String()
}

func (s *ExternalQueryExternalAppOrgsResponseBody) SetResult(v []*ExternalQueryExternalAppOrgsResponseBodyResult) *ExternalQueryExternalAppOrgsResponseBody {
	s.Result = v
	return s
}

type ExternalQueryExternalAppOrgsResponseBodyResult struct {
	CorpId   *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	CorpName *string `json:"corpName,omitempty" xml:"corpName,omitempty"`
}

func (s ExternalQueryExternalAppOrgsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ExternalQueryExternalAppOrgsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ExternalQueryExternalAppOrgsResponseBodyResult) SetCorpId(v string) *ExternalQueryExternalAppOrgsResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *ExternalQueryExternalAppOrgsResponseBodyResult) SetCorpName(v string) *ExternalQueryExternalAppOrgsResponseBodyResult {
	s.CorpName = &v
	return s
}

type ExternalQueryExternalAppOrgsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExternalQueryExternalAppOrgsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExternalQueryExternalAppOrgsResponse) String() string {
	return tea.Prettify(s)
}

func (s ExternalQueryExternalAppOrgsResponse) GoString() string {
	return s.String()
}

func (s *ExternalQueryExternalAppOrgsResponse) SetHeaders(v map[string]*string) *ExternalQueryExternalAppOrgsResponse {
	s.Headers = v
	return s
}

func (s *ExternalQueryExternalAppOrgsResponse) SetStatusCode(v int32) *ExternalQueryExternalAppOrgsResponse {
	s.StatusCode = &v
	return s
}

func (s *ExternalQueryExternalAppOrgsResponse) SetBody(v *ExternalQueryExternalAppOrgsResponseBody) *ExternalQueryExternalAppOrgsResponse {
	s.Body = v
	return s
}

type ExternalQueryExternalBelongMainOrgHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ExternalQueryExternalBelongMainOrgHeaders) String() string {
	return tea.Prettify(s)
}

func (s ExternalQueryExternalBelongMainOrgHeaders) GoString() string {
	return s.String()
}

func (s *ExternalQueryExternalBelongMainOrgHeaders) SetCommonHeaders(v map[string]*string) *ExternalQueryExternalBelongMainOrgHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ExternalQueryExternalBelongMainOrgHeaders) SetXAcsDingtalkAccessToken(v string) *ExternalQueryExternalBelongMainOrgHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ExternalQueryExternalBelongMainOrgRequest struct {
	ExternalType *string `json:"externalType,omitempty" xml:"externalType,omitempty"`
}

func (s ExternalQueryExternalBelongMainOrgRequest) String() string {
	return tea.Prettify(s)
}

func (s ExternalQueryExternalBelongMainOrgRequest) GoString() string {
	return s.String()
}

func (s *ExternalQueryExternalBelongMainOrgRequest) SetExternalType(v string) *ExternalQueryExternalBelongMainOrgRequest {
	s.ExternalType = &v
	return s
}

type ExternalQueryExternalBelongMainOrgResponseBody struct {
	CorpId   *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	CorpName *string `json:"corpName,omitempty" xml:"corpName,omitempty"`
}

func (s ExternalQueryExternalBelongMainOrgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExternalQueryExternalBelongMainOrgResponseBody) GoString() string {
	return s.String()
}

func (s *ExternalQueryExternalBelongMainOrgResponseBody) SetCorpId(v string) *ExternalQueryExternalBelongMainOrgResponseBody {
	s.CorpId = &v
	return s
}

func (s *ExternalQueryExternalBelongMainOrgResponseBody) SetCorpName(v string) *ExternalQueryExternalBelongMainOrgResponseBody {
	s.CorpName = &v
	return s
}

type ExternalQueryExternalBelongMainOrgResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExternalQueryExternalBelongMainOrgResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExternalQueryExternalBelongMainOrgResponse) String() string {
	return tea.Prettify(s)
}

func (s ExternalQueryExternalBelongMainOrgResponse) GoString() string {
	return s.String()
}

func (s *ExternalQueryExternalBelongMainOrgResponse) SetHeaders(v map[string]*string) *ExternalQueryExternalBelongMainOrgResponse {
	s.Headers = v
	return s
}

func (s *ExternalQueryExternalBelongMainOrgResponse) SetStatusCode(v int32) *ExternalQueryExternalBelongMainOrgResponse {
	s.StatusCode = &v
	return s
}

func (s *ExternalQueryExternalBelongMainOrgResponse) SetBody(v *ExternalQueryExternalBelongMainOrgResponseBody) *ExternalQueryExternalBelongMainOrgResponse {
	s.Body = v
	return s
}

type ExternalQueryExternalOrgsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ExternalQueryExternalOrgsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ExternalQueryExternalOrgsHeaders) GoString() string {
	return s.String()
}

func (s *ExternalQueryExternalOrgsHeaders) SetCommonHeaders(v map[string]*string) *ExternalQueryExternalOrgsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ExternalQueryExternalOrgsHeaders) SetXAcsDingtalkAccessToken(v string) *ExternalQueryExternalOrgsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ExternalQueryExternalOrgsRequest struct {
	ExternalType *string `json:"externalType,omitempty" xml:"externalType,omitempty"`
}

func (s ExternalQueryExternalOrgsRequest) String() string {
	return tea.Prettify(s)
}

func (s ExternalQueryExternalOrgsRequest) GoString() string {
	return s.String()
}

func (s *ExternalQueryExternalOrgsRequest) SetExternalType(v string) *ExternalQueryExternalOrgsRequest {
	s.ExternalType = &v
	return s
}

type ExternalQueryExternalOrgsResponseBody struct {
	Result []*ExternalQueryExternalOrgsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ExternalQueryExternalOrgsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExternalQueryExternalOrgsResponseBody) GoString() string {
	return s.String()
}

func (s *ExternalQueryExternalOrgsResponseBody) SetResult(v []*ExternalQueryExternalOrgsResponseBodyResult) *ExternalQueryExternalOrgsResponseBody {
	s.Result = v
	return s
}

type ExternalQueryExternalOrgsResponseBodyResult struct {
	CorpId   *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	CorpName *string `json:"corpName,omitempty" xml:"corpName,omitempty"`
}

func (s ExternalQueryExternalOrgsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ExternalQueryExternalOrgsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ExternalQueryExternalOrgsResponseBodyResult) SetCorpId(v string) *ExternalQueryExternalOrgsResponseBodyResult {
	s.CorpId = &v
	return s
}

func (s *ExternalQueryExternalOrgsResponseBodyResult) SetCorpName(v string) *ExternalQueryExternalOrgsResponseBodyResult {
	s.CorpName = &v
	return s
}

type ExternalQueryExternalOrgsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExternalQueryExternalOrgsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExternalQueryExternalOrgsResponse) String() string {
	return tea.Prettify(s)
}

func (s ExternalQueryExternalOrgsResponse) GoString() string {
	return s.String()
}

func (s *ExternalQueryExternalOrgsResponse) SetHeaders(v map[string]*string) *ExternalQueryExternalOrgsResponse {
	s.Headers = v
	return s
}

func (s *ExternalQueryExternalOrgsResponse) SetStatusCode(v int32) *ExternalQueryExternalOrgsResponse {
	s.StatusCode = &v
	return s
}

func (s *ExternalQueryExternalOrgsResponse) SetBody(v *ExternalQueryExternalOrgsResponseBody) *ExternalQueryExternalOrgsResponse {
	s.Body = v
	return s
}

type HospitalDataCheckHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s HospitalDataCheckHeaders) String() string {
	return tea.Prettify(s)
}

func (s HospitalDataCheckHeaders) GoString() string {
	return s.String()
}

func (s *HospitalDataCheckHeaders) SetCommonHeaders(v map[string]*string) *HospitalDataCheckHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HospitalDataCheckHeaders) SetXAcsDingtalkAccessToken(v string) *HospitalDataCheckHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type HospitalDataCheckRequest struct {
	AllDeptCount      *int64 `json:"allDeptCount,omitempty" xml:"allDeptCount,omitempty"`
	AllDeptUserCount  *int64 `json:"allDeptUserCount,omitempty" xml:"allDeptUserCount,omitempty"`
	AllGroupCount     *int64 `json:"allGroupCount,omitempty" xml:"allGroupCount,omitempty"`
	AllGroupUserCount *int64 `json:"allGroupUserCount,omitempty" xml:"allGroupUserCount,omitempty"`
	DeptCount         *int64 `json:"deptCount,omitempty" xml:"deptCount,omitempty"`
	DeptUserCount     *int64 `json:"deptUserCount,omitempty" xml:"deptUserCount,omitempty"`
	GroupCount        *int64 `json:"groupCount,omitempty" xml:"groupCount,omitempty"`
	GroupUserCount    *int64 `json:"groupUserCount,omitempty" xml:"groupUserCount,omitempty"`
}

func (s HospitalDataCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s HospitalDataCheckRequest) GoString() string {
	return s.String()
}

func (s *HospitalDataCheckRequest) SetAllDeptCount(v int64) *HospitalDataCheckRequest {
	s.AllDeptCount = &v
	return s
}

func (s *HospitalDataCheckRequest) SetAllDeptUserCount(v int64) *HospitalDataCheckRequest {
	s.AllDeptUserCount = &v
	return s
}

func (s *HospitalDataCheckRequest) SetAllGroupCount(v int64) *HospitalDataCheckRequest {
	s.AllGroupCount = &v
	return s
}

func (s *HospitalDataCheckRequest) SetAllGroupUserCount(v int64) *HospitalDataCheckRequest {
	s.AllGroupUserCount = &v
	return s
}

func (s *HospitalDataCheckRequest) SetDeptCount(v int64) *HospitalDataCheckRequest {
	s.DeptCount = &v
	return s
}

func (s *HospitalDataCheckRequest) SetDeptUserCount(v int64) *HospitalDataCheckRequest {
	s.DeptUserCount = &v
	return s
}

func (s *HospitalDataCheckRequest) SetGroupCount(v int64) *HospitalDataCheckRequest {
	s.GroupCount = &v
	return s
}

func (s *HospitalDataCheckRequest) SetGroupUserCount(v int64) *HospitalDataCheckRequest {
	s.GroupUserCount = &v
	return s
}

type HospitalDataCheckResponseBody struct {
	AllDeptCount      *int64 `json:"allDeptCount,omitempty" xml:"allDeptCount,omitempty"`
	AllDeptUserCount  *int64 `json:"allDeptUserCount,omitempty" xml:"allDeptUserCount,omitempty"`
	AllGroupCount     *int64 `json:"allGroupCount,omitempty" xml:"allGroupCount,omitempty"`
	AllGroupUserCount *int64 `json:"allGroupUserCount,omitempty" xml:"allGroupUserCount,omitempty"`
	DeptCount         *int64 `json:"deptCount,omitempty" xml:"deptCount,omitempty"`
	DeptUserCount     *int64 `json:"deptUserCount,omitempty" xml:"deptUserCount,omitempty"`
	GroupCount        *int64 `json:"groupCount,omitempty" xml:"groupCount,omitempty"`
	GroupUserCount    *int64 `json:"groupUserCount,omitempty" xml:"groupUserCount,omitempty"`
	Match             *bool  `json:"match,omitempty" xml:"match,omitempty"`
}

func (s HospitalDataCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HospitalDataCheckResponseBody) GoString() string {
	return s.String()
}

func (s *HospitalDataCheckResponseBody) SetAllDeptCount(v int64) *HospitalDataCheckResponseBody {
	s.AllDeptCount = &v
	return s
}

func (s *HospitalDataCheckResponseBody) SetAllDeptUserCount(v int64) *HospitalDataCheckResponseBody {
	s.AllDeptUserCount = &v
	return s
}

func (s *HospitalDataCheckResponseBody) SetAllGroupCount(v int64) *HospitalDataCheckResponseBody {
	s.AllGroupCount = &v
	return s
}

func (s *HospitalDataCheckResponseBody) SetAllGroupUserCount(v int64) *HospitalDataCheckResponseBody {
	s.AllGroupUserCount = &v
	return s
}

func (s *HospitalDataCheckResponseBody) SetDeptCount(v int64) *HospitalDataCheckResponseBody {
	s.DeptCount = &v
	return s
}

func (s *HospitalDataCheckResponseBody) SetDeptUserCount(v int64) *HospitalDataCheckResponseBody {
	s.DeptUserCount = &v
	return s
}

func (s *HospitalDataCheckResponseBody) SetGroupCount(v int64) *HospitalDataCheckResponseBody {
	s.GroupCount = &v
	return s
}

func (s *HospitalDataCheckResponseBody) SetGroupUserCount(v int64) *HospitalDataCheckResponseBody {
	s.GroupUserCount = &v
	return s
}

func (s *HospitalDataCheckResponseBody) SetMatch(v bool) *HospitalDataCheckResponseBody {
	s.Match = &v
	return s
}

type HospitalDataCheckResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *HospitalDataCheckResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HospitalDataCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s HospitalDataCheckResponse) GoString() string {
	return s.String()
}

func (s *HospitalDataCheckResponse) SetHeaders(v map[string]*string) *HospitalDataCheckResponse {
	s.Headers = v
	return s
}

func (s *HospitalDataCheckResponse) SetStatusCode(v int32) *HospitalDataCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *HospitalDataCheckResponse) SetBody(v *HospitalDataCheckResponseBody) *HospitalDataCheckResponse {
	s.Body = v
	return s
}

type IndustryManufactureCommonEventHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IndustryManufactureCommonEventHeaders) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureCommonEventHeaders) GoString() string {
	return s.String()
}

func (s *IndustryManufactureCommonEventHeaders) SetCommonHeaders(v map[string]*string) *IndustryManufactureCommonEventHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IndustryManufactureCommonEventHeaders) SetXAcsDingtalkAccessToken(v string) *IndustryManufactureCommonEventHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IndustryManufactureCommonEventRequest struct {
	Action    *string                `json:"action,omitempty" xml:"action,omitempty"`
	AppKey    *string                `json:"appKey,omitempty" xml:"appKey,omitempty"`
	BizData   map[string]interface{} `json:"bizData,omitempty" xml:"bizData,omitempty"`
	EventType []*string              `json:"eventType,omitempty" xml:"eventType,omitempty" type:"Repeated"`
}

func (s IndustryManufactureCommonEventRequest) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureCommonEventRequest) GoString() string {
	return s.String()
}

func (s *IndustryManufactureCommonEventRequest) SetAction(v string) *IndustryManufactureCommonEventRequest {
	s.Action = &v
	return s
}

func (s *IndustryManufactureCommonEventRequest) SetAppKey(v string) *IndustryManufactureCommonEventRequest {
	s.AppKey = &v
	return s
}

func (s *IndustryManufactureCommonEventRequest) SetBizData(v map[string]interface{}) *IndustryManufactureCommonEventRequest {
	s.BizData = v
	return s
}

func (s *IndustryManufactureCommonEventRequest) SetEventType(v []*string) *IndustryManufactureCommonEventRequest {
	s.EventType = v
	return s
}

type IndustryManufactureCommonEventResponseBody struct {
	ErrorMsg  *string                                           `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	RequestId *string                                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *IndustryManufactureCommonEventResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s IndustryManufactureCommonEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureCommonEventResponseBody) GoString() string {
	return s.String()
}

func (s *IndustryManufactureCommonEventResponseBody) SetErrorMsg(v string) *IndustryManufactureCommonEventResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *IndustryManufactureCommonEventResponseBody) SetRequestId(v string) *IndustryManufactureCommonEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *IndustryManufactureCommonEventResponseBody) SetResult(v *IndustryManufactureCommonEventResponseBodyResult) *IndustryManufactureCommonEventResponseBody {
	s.Result = v
	return s
}

type IndustryManufactureCommonEventResponseBodyResult struct {
	Content  *string `json:"content,omitempty" xml:"content,omitempty"`
	HttpCode *string `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
}

func (s IndustryManufactureCommonEventResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureCommonEventResponseBodyResult) GoString() string {
	return s.String()
}

func (s *IndustryManufactureCommonEventResponseBodyResult) SetContent(v string) *IndustryManufactureCommonEventResponseBodyResult {
	s.Content = &v
	return s
}

func (s *IndustryManufactureCommonEventResponseBodyResult) SetHttpCode(v string) *IndustryManufactureCommonEventResponseBodyResult {
	s.HttpCode = &v
	return s
}

type IndustryManufactureCommonEventResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IndustryManufactureCommonEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndustryManufactureCommonEventResponse) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureCommonEventResponse) GoString() string {
	return s.String()
}

func (s *IndustryManufactureCommonEventResponse) SetHeaders(v map[string]*string) *IndustryManufactureCommonEventResponse {
	s.Headers = v
	return s
}

func (s *IndustryManufactureCommonEventResponse) SetStatusCode(v int32) *IndustryManufactureCommonEventResponse {
	s.StatusCode = &v
	return s
}

func (s *IndustryManufactureCommonEventResponse) SetBody(v *IndustryManufactureCommonEventResponseBody) *IndustryManufactureCommonEventResponse {
	s.Body = v
	return s
}

type IndustryManufactureCostRecordListGetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IndustryManufactureCostRecordListGetHeaders) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureCostRecordListGetHeaders) GoString() string {
	return s.String()
}

func (s *IndustryManufactureCostRecordListGetHeaders) SetCommonHeaders(v map[string]*string) *IndustryManufactureCostRecordListGetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IndustryManufactureCostRecordListGetHeaders) SetXAcsDingtalkAccessToken(v string) *IndustryManufactureCostRecordListGetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IndustryManufactureCostRecordListGetRequest struct {
	AppId            *int64   `json:"appId,omitempty" xml:"appId,omitempty"`
	AppIds           []*int64 `json:"appIds,omitempty" xml:"appIds,omitempty" type:"Repeated"`
	AppName          *string  `json:"appName,omitempty" xml:"appName,omitempty"`
	CorpId           *string  `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Cursor           *int64   `json:"cursor,omitempty" xml:"cursor,omitempty"`
	EndTime          *int64   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	InstanceId       *string  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	IsvOrgId         *int64   `json:"isvOrgId,omitempty" xml:"isvOrgId,omitempty"`
	MaterialNo       *string  `json:"materialNo,omitempty" xml:"materialNo,omitempty"`
	MicroappAgentId  *int64   `json:"microappAgentId,omitempty" xml:"microappAgentId,omitempty"`
	OrderNo          *string  `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	OrgId            *int64   `json:"orgId,omitempty" xml:"orgId,omitempty"`
	PageNumber       *int64   `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize         *int32   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ProductionTaskNo *string  `json:"productionTaskNo,omitempty" xml:"productionTaskNo,omitempty"`
	StartTime        *int64   `json:"startTime,omitempty" xml:"startTime,omitempty"`
	SuiteKey         *string  `json:"suiteKey,omitempty" xml:"suiteKey,omitempty"`
	TokenGrantType   *int32   `json:"tokenGrantType,omitempty" xml:"tokenGrantType,omitempty"`
}

func (s IndustryManufactureCostRecordListGetRequest) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureCostRecordListGetRequest) GoString() string {
	return s.String()
}

func (s *IndustryManufactureCostRecordListGetRequest) SetAppId(v int64) *IndustryManufactureCostRecordListGetRequest {
	s.AppId = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetAppIds(v []*int64) *IndustryManufactureCostRecordListGetRequest {
	s.AppIds = v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetAppName(v string) *IndustryManufactureCostRecordListGetRequest {
	s.AppName = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetCorpId(v string) *IndustryManufactureCostRecordListGetRequest {
	s.CorpId = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetCursor(v int64) *IndustryManufactureCostRecordListGetRequest {
	s.Cursor = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetEndTime(v int64) *IndustryManufactureCostRecordListGetRequest {
	s.EndTime = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetInstanceId(v string) *IndustryManufactureCostRecordListGetRequest {
	s.InstanceId = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetIsvOrgId(v int64) *IndustryManufactureCostRecordListGetRequest {
	s.IsvOrgId = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetMaterialNo(v string) *IndustryManufactureCostRecordListGetRequest {
	s.MaterialNo = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetMicroappAgentId(v int64) *IndustryManufactureCostRecordListGetRequest {
	s.MicroappAgentId = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetOrderNo(v string) *IndustryManufactureCostRecordListGetRequest {
	s.OrderNo = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetOrgId(v int64) *IndustryManufactureCostRecordListGetRequest {
	s.OrgId = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetPageNumber(v int64) *IndustryManufactureCostRecordListGetRequest {
	s.PageNumber = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetPageSize(v int32) *IndustryManufactureCostRecordListGetRequest {
	s.PageSize = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetProductionTaskNo(v string) *IndustryManufactureCostRecordListGetRequest {
	s.ProductionTaskNo = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetStartTime(v int64) *IndustryManufactureCostRecordListGetRequest {
	s.StartTime = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetSuiteKey(v string) *IndustryManufactureCostRecordListGetRequest {
	s.SuiteKey = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetRequest) SetTokenGrantType(v int32) *IndustryManufactureCostRecordListGetRequest {
	s.TokenGrantType = &v
	return s
}

type IndustryManufactureCostRecordListGetResponseBody struct {
	HasMore    *bool                                                   `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*IndustryManufactureCostRecordListGetResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextCursor *int64                                                  `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	TotalCount *int64                                                  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s IndustryManufactureCostRecordListGetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureCostRecordListGetResponseBody) GoString() string {
	return s.String()
}

func (s *IndustryManufactureCostRecordListGetResponseBody) SetHasMore(v bool) *IndustryManufactureCostRecordListGetResponseBody {
	s.HasMore = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBody) SetList(v []*IndustryManufactureCostRecordListGetResponseBodyList) *IndustryManufactureCostRecordListGetResponseBody {
	s.List = v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBody) SetNextCursor(v int64) *IndustryManufactureCostRecordListGetResponseBody {
	s.NextCursor = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBody) SetTotalCount(v int64) *IndustryManufactureCostRecordListGetResponseBody {
	s.TotalCount = &v
	return s
}

type IndustryManufactureCostRecordListGetResponseBodyList struct {
	CorpId               *string  `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Count                *float32 `json:"count,omitempty" xml:"count,omitempty"`
	Ext                  *string  `json:"ext,omitempty" xml:"ext,omitempty"`
	GmtCreate            *int64   `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified          *int64   `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	InstanceId           *string  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	IsDeleted            *string  `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	MaterialCostRecordNo *string  `json:"materialCostRecordNo,omitempty" xml:"materialCostRecordNo,omitempty"`
	MaterialName         *string  `json:"materialName,omitempty" xml:"materialName,omitempty"`
	MaterialNo           *string  `json:"materialNo,omitempty" xml:"materialNo,omitempty"`
	Memo                 *string  `json:"memo,omitempty" xml:"memo,omitempty"`
	OrderNo              *string  `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	Price                *float32 `json:"price,omitempty" xml:"price,omitempty"`
	ProcessCode          *string  `json:"processCode,omitempty" xml:"processCode,omitempty"`
	ProductionTaskNo     *string  `json:"productionTaskNo,omitempty" xml:"productionTaskNo,omitempty"`
	RealCount            *float32 `json:"realCount,omitempty" xml:"realCount,omitempty"`
	RealPrice            *float32 `json:"realPrice,omitempty" xml:"realPrice,omitempty"`
	Type                 *string  `json:"type,omitempty" xml:"type,omitempty"`
	Unit                 *string  `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s IndustryManufactureCostRecordListGetResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureCostRecordListGetResponseBodyList) GoString() string {
	return s.String()
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetCorpId(v string) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.CorpId = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetCount(v float32) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.Count = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetExt(v string) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.Ext = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetGmtCreate(v int64) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.GmtCreate = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetGmtModified(v int64) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.GmtModified = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetInstanceId(v string) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.InstanceId = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetIsDeleted(v string) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.IsDeleted = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetMaterialCostRecordNo(v string) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.MaterialCostRecordNo = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetMaterialName(v string) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.MaterialName = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetMaterialNo(v string) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.MaterialNo = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetMemo(v string) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.Memo = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetOrderNo(v string) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.OrderNo = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetPrice(v float32) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.Price = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetProcessCode(v string) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.ProcessCode = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetProductionTaskNo(v string) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.ProductionTaskNo = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetRealCount(v float32) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.RealCount = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetRealPrice(v float32) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.RealPrice = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetType(v string) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.Type = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponseBodyList) SetUnit(v string) *IndustryManufactureCostRecordListGetResponseBodyList {
	s.Unit = &v
	return s
}

type IndustryManufactureCostRecordListGetResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IndustryManufactureCostRecordListGetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndustryManufactureCostRecordListGetResponse) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureCostRecordListGetResponse) GoString() string {
	return s.String()
}

func (s *IndustryManufactureCostRecordListGetResponse) SetHeaders(v map[string]*string) *IndustryManufactureCostRecordListGetResponse {
	s.Headers = v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponse) SetStatusCode(v int32) *IndustryManufactureCostRecordListGetResponse {
	s.StatusCode = &v
	return s
}

func (s *IndustryManufactureCostRecordListGetResponse) SetBody(v *IndustryManufactureCostRecordListGetResponseBody) *IndustryManufactureCostRecordListGetResponse {
	s.Body = v
	return s
}

type IndustryManufactureFeeListGetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IndustryManufactureFeeListGetHeaders) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureFeeListGetHeaders) GoString() string {
	return s.String()
}

func (s *IndustryManufactureFeeListGetHeaders) SetCommonHeaders(v map[string]*string) *IndustryManufactureFeeListGetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IndustryManufactureFeeListGetHeaders) SetXAcsDingtalkAccessToken(v string) *IndustryManufactureFeeListGetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IndustryManufactureFeeListGetRequest struct {
	AppId            *int64   `json:"appId,omitempty" xml:"appId,omitempty"`
	AppIds           []*int64 `json:"appIds,omitempty" xml:"appIds,omitempty" type:"Repeated"`
	AppName          *string  `json:"appName,omitempty" xml:"appName,omitempty"`
	CorpId           *string  `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Cursor           *int64   `json:"cursor,omitempty" xml:"cursor,omitempty"`
	EndTime          *int64   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	IsvOrgId         *int64   `json:"isvOrgId,omitempty" xml:"isvOrgId,omitempty"`
	MaterialNo       *string  `json:"materialNo,omitempty" xml:"materialNo,omitempty"`
	MicroappAgentId  *int64   `json:"microappAgentId,omitempty" xml:"microappAgentId,omitempty"`
	OrgId            *int64   `json:"orgId,omitempty" xml:"orgId,omitempty"`
	PageNumber       *int64   `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize         *int32   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ProductionTaskNo *string  `json:"productionTaskNo,omitempty" xml:"productionTaskNo,omitempty"`
	StartTime        *int64   `json:"startTime,omitempty" xml:"startTime,omitempty"`
	SuiteKey         *string  `json:"suiteKey,omitempty" xml:"suiteKey,omitempty"`
	TokenGrantType   *int32   `json:"tokenGrantType,omitempty" xml:"tokenGrantType,omitempty"`
	Type             *string  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s IndustryManufactureFeeListGetRequest) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureFeeListGetRequest) GoString() string {
	return s.String()
}

func (s *IndustryManufactureFeeListGetRequest) SetAppId(v int64) *IndustryManufactureFeeListGetRequest {
	s.AppId = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetAppIds(v []*int64) *IndustryManufactureFeeListGetRequest {
	s.AppIds = v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetAppName(v string) *IndustryManufactureFeeListGetRequest {
	s.AppName = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetCorpId(v string) *IndustryManufactureFeeListGetRequest {
	s.CorpId = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetCursor(v int64) *IndustryManufactureFeeListGetRequest {
	s.Cursor = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetEndTime(v int64) *IndustryManufactureFeeListGetRequest {
	s.EndTime = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetIsvOrgId(v int64) *IndustryManufactureFeeListGetRequest {
	s.IsvOrgId = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetMaterialNo(v string) *IndustryManufactureFeeListGetRequest {
	s.MaterialNo = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetMicroappAgentId(v int64) *IndustryManufactureFeeListGetRequest {
	s.MicroappAgentId = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetOrgId(v int64) *IndustryManufactureFeeListGetRequest {
	s.OrgId = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetPageNumber(v int64) *IndustryManufactureFeeListGetRequest {
	s.PageNumber = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetPageSize(v int32) *IndustryManufactureFeeListGetRequest {
	s.PageSize = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetProductionTaskNo(v string) *IndustryManufactureFeeListGetRequest {
	s.ProductionTaskNo = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetStartTime(v int64) *IndustryManufactureFeeListGetRequest {
	s.StartTime = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetSuiteKey(v string) *IndustryManufactureFeeListGetRequest {
	s.SuiteKey = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetTokenGrantType(v int32) *IndustryManufactureFeeListGetRequest {
	s.TokenGrantType = &v
	return s
}

func (s *IndustryManufactureFeeListGetRequest) SetType(v string) *IndustryManufactureFeeListGetRequest {
	s.Type = &v
	return s
}

type IndustryManufactureFeeListGetResponseBody struct {
	HasMore    *bool                                            `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*IndustryManufactureFeeListGetResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextCursor *int64                                           `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	TotalCount *int64                                           `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s IndustryManufactureFeeListGetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureFeeListGetResponseBody) GoString() string {
	return s.String()
}

func (s *IndustryManufactureFeeListGetResponseBody) SetHasMore(v bool) *IndustryManufactureFeeListGetResponseBody {
	s.HasMore = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBody) SetList(v []*IndustryManufactureFeeListGetResponseBodyList) *IndustryManufactureFeeListGetResponseBody {
	s.List = v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBody) SetNextCursor(v int64) *IndustryManufactureFeeListGetResponseBody {
	s.NextCursor = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBody) SetTotalCount(v int64) *IndustryManufactureFeeListGetResponseBody {
	s.TotalCount = &v
	return s
}

type IndustryManufactureFeeListGetResponseBodyList struct {
	Amount           *string  `json:"amount,omitempty" xml:"amount,omitempty"`
	CorpId           *string  `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Count            *float32 `json:"count,omitempty" xml:"count,omitempty"`
	Ext              *string  `json:"ext,omitempty" xml:"ext,omitempty"`
	GmtCreate        *int64   `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified      *int64   `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id               *int64   `json:"id,omitempty" xml:"id,omitempty"`
	InstanceId       *string  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	IsDeleted        *string  `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	MaterialName     *string  `json:"materialName,omitempty" xml:"materialName,omitempty"`
	MaterialNo       *string  `json:"materialNo,omitempty" xml:"materialNo,omitempty"`
	PerAmount        *float32 `json:"perAmount,omitempty" xml:"perAmount,omitempty"`
	ProcessCode      *string  `json:"processCode,omitempty" xml:"processCode,omitempty"`
	ProductionTaskNo *string  `json:"productionTaskNo,omitempty" xml:"productionTaskNo,omitempty"`
	Title            *string  `json:"title,omitempty" xml:"title,omitempty"`
	Type             *string  `json:"type,omitempty" xml:"type,omitempty"`
	Unit             *string  `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s IndustryManufactureFeeListGetResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureFeeListGetResponseBodyList) GoString() string {
	return s.String()
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetAmount(v string) *IndustryManufactureFeeListGetResponseBodyList {
	s.Amount = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetCorpId(v string) *IndustryManufactureFeeListGetResponseBodyList {
	s.CorpId = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetCount(v float32) *IndustryManufactureFeeListGetResponseBodyList {
	s.Count = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetExt(v string) *IndustryManufactureFeeListGetResponseBodyList {
	s.Ext = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetGmtCreate(v int64) *IndustryManufactureFeeListGetResponseBodyList {
	s.GmtCreate = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetGmtModified(v int64) *IndustryManufactureFeeListGetResponseBodyList {
	s.GmtModified = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetId(v int64) *IndustryManufactureFeeListGetResponseBodyList {
	s.Id = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetInstanceId(v string) *IndustryManufactureFeeListGetResponseBodyList {
	s.InstanceId = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetIsDeleted(v string) *IndustryManufactureFeeListGetResponseBodyList {
	s.IsDeleted = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetMaterialName(v string) *IndustryManufactureFeeListGetResponseBodyList {
	s.MaterialName = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetMaterialNo(v string) *IndustryManufactureFeeListGetResponseBodyList {
	s.MaterialNo = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetPerAmount(v float32) *IndustryManufactureFeeListGetResponseBodyList {
	s.PerAmount = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetProcessCode(v string) *IndustryManufactureFeeListGetResponseBodyList {
	s.ProcessCode = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetProductionTaskNo(v string) *IndustryManufactureFeeListGetResponseBodyList {
	s.ProductionTaskNo = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetTitle(v string) *IndustryManufactureFeeListGetResponseBodyList {
	s.Title = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetType(v string) *IndustryManufactureFeeListGetResponseBodyList {
	s.Type = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponseBodyList) SetUnit(v string) *IndustryManufactureFeeListGetResponseBodyList {
	s.Unit = &v
	return s
}

type IndustryManufactureFeeListGetResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IndustryManufactureFeeListGetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndustryManufactureFeeListGetResponse) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureFeeListGetResponse) GoString() string {
	return s.String()
}

func (s *IndustryManufactureFeeListGetResponse) SetHeaders(v map[string]*string) *IndustryManufactureFeeListGetResponse {
	s.Headers = v
	return s
}

func (s *IndustryManufactureFeeListGetResponse) SetStatusCode(v int32) *IndustryManufactureFeeListGetResponse {
	s.StatusCode = &v
	return s
}

func (s *IndustryManufactureFeeListGetResponse) SetBody(v *IndustryManufactureFeeListGetResponseBody) *IndustryManufactureFeeListGetResponse {
	s.Body = v
	return s
}

type IndustryManufactureLabourCostHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IndustryManufactureLabourCostHeaders) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureLabourCostHeaders) GoString() string {
	return s.String()
}

func (s *IndustryManufactureLabourCostHeaders) SetCommonHeaders(v map[string]*string) *IndustryManufactureLabourCostHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IndustryManufactureLabourCostHeaders) SetXAcsDingtalkAccessToken(v string) *IndustryManufactureLabourCostHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IndustryManufactureLabourCostRequest struct {
	AppId           *int64   `json:"appId,omitempty" xml:"appId,omitempty"`
	AppIds          []*int64 `json:"appIds,omitempty" xml:"appIds,omitempty" type:"Repeated"`
	AppName         *string  `json:"appName,omitempty" xml:"appName,omitempty"`
	CorpId          *string  `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Cursor          *int64   `json:"cursor,omitempty" xml:"cursor,omitempty"`
	EndTime         *int64   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	IsvOrgId        *string  `json:"isvOrgId,omitempty" xml:"isvOrgId,omitempty"`
	MaterialNo      *string  `json:"materialNo,omitempty" xml:"materialNo,omitempty"`
	MicroappAgentId *int64   `json:"microappAgentId,omitempty" xml:"microappAgentId,omitempty"`
	OrgId           *int64   `json:"orgId,omitempty" xml:"orgId,omitempty"`
	PageNumber      *int64   `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize        *int32   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ProcessNo       *string  `json:"processNo,omitempty" xml:"processNo,omitempty"`
	StartTime       *int64   `json:"startTime,omitempty" xml:"startTime,omitempty"`
	SuiteKey        *string  `json:"suiteKey,omitempty" xml:"suiteKey,omitempty"`
	TokenGrantType  *int32   `json:"tokenGrantType,omitempty" xml:"tokenGrantType,omitempty"`
}

func (s IndustryManufactureLabourCostRequest) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureLabourCostRequest) GoString() string {
	return s.String()
}

func (s *IndustryManufactureLabourCostRequest) SetAppId(v int64) *IndustryManufactureLabourCostRequest {
	s.AppId = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetAppIds(v []*int64) *IndustryManufactureLabourCostRequest {
	s.AppIds = v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetAppName(v string) *IndustryManufactureLabourCostRequest {
	s.AppName = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetCorpId(v string) *IndustryManufactureLabourCostRequest {
	s.CorpId = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetCursor(v int64) *IndustryManufactureLabourCostRequest {
	s.Cursor = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetEndTime(v int64) *IndustryManufactureLabourCostRequest {
	s.EndTime = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetIsvOrgId(v string) *IndustryManufactureLabourCostRequest {
	s.IsvOrgId = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetMaterialNo(v string) *IndustryManufactureLabourCostRequest {
	s.MaterialNo = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetMicroappAgentId(v int64) *IndustryManufactureLabourCostRequest {
	s.MicroappAgentId = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetOrgId(v int64) *IndustryManufactureLabourCostRequest {
	s.OrgId = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetPageNumber(v int64) *IndustryManufactureLabourCostRequest {
	s.PageNumber = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetPageSize(v int32) *IndustryManufactureLabourCostRequest {
	s.PageSize = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetProcessNo(v string) *IndustryManufactureLabourCostRequest {
	s.ProcessNo = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetStartTime(v int64) *IndustryManufactureLabourCostRequest {
	s.StartTime = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetSuiteKey(v string) *IndustryManufactureLabourCostRequest {
	s.SuiteKey = &v
	return s
}

func (s *IndustryManufactureLabourCostRequest) SetTokenGrantType(v int32) *IndustryManufactureLabourCostRequest {
	s.TokenGrantType = &v
	return s
}

type IndustryManufactureLabourCostResponseBody struct {
	HasMore    *bool                                            `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*IndustryManufactureLabourCostResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextCursor *int64                                           `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	TotalCount *int64                                           `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s IndustryManufactureLabourCostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureLabourCostResponseBody) GoString() string {
	return s.String()
}

func (s *IndustryManufactureLabourCostResponseBody) SetHasMore(v bool) *IndustryManufactureLabourCostResponseBody {
	s.HasMore = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBody) SetList(v []*IndustryManufactureLabourCostResponseBodyList) *IndustryManufactureLabourCostResponseBody {
	s.List = v
	return s
}

func (s *IndustryManufactureLabourCostResponseBody) SetNextCursor(v int64) *IndustryManufactureLabourCostResponseBody {
	s.NextCursor = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBody) SetTotalCount(v int64) *IndustryManufactureLabourCostResponseBody {
	s.TotalCount = &v
	return s
}

type IndustryManufactureLabourCostResponseBodyList struct {
	CorpId             *string  `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Ext                *string  `json:"ext,omitempty" xml:"ext,omitempty"`
	GmtCreate          *int64   `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified        *int64   `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	InstanceId         *string  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	LabourCostName     *string  `json:"labourCostName,omitempty" xml:"labourCostName,omitempty"`
	LabourCostNo       *string  `json:"labourCostNo,omitempty" xml:"labourCostNo,omitempty"`
	MaterialName       *string  `json:"materialName,omitempty" xml:"materialName,omitempty"`
	MaterialNo         *string  `json:"materialNo,omitempty" xml:"materialNo,omitempty"`
	ProcessCode        *string  `json:"processCode,omitempty" xml:"processCode,omitempty"`
	ProcessName        *string  `json:"processName,omitempty" xml:"processName,omitempty"`
	ProcessNo          *string  `json:"processNo,omitempty" xml:"processNo,omitempty"`
	QualifiedPrice     *float32 `json:"qualifiedPrice,omitempty" xml:"qualifiedPrice,omitempty"`
	UnQualifiedInfo    *string  `json:"unQualifiedInfo,omitempty" xml:"unQualifiedInfo,omitempty"`
	UnQualifiedPrice1  *float32 `json:"unQualifiedPrice1,omitempty" xml:"unQualifiedPrice1,omitempty"`
	UnQualifiedPrice2  *float32 `json:"unQualifiedPrice2,omitempty" xml:"unQualifiedPrice2,omitempty"`
	UnQualifiedReason1 *string  `json:"unQualifiedReason1,omitempty" xml:"unQualifiedReason1,omitempty"`
	UnQualifiedReason2 *string  `json:"unQualifiedReason2,omitempty" xml:"unQualifiedReason2,omitempty"`
}

func (s IndustryManufactureLabourCostResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureLabourCostResponseBodyList) GoString() string {
	return s.String()
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetCorpId(v string) *IndustryManufactureLabourCostResponseBodyList {
	s.CorpId = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetExt(v string) *IndustryManufactureLabourCostResponseBodyList {
	s.Ext = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetGmtCreate(v int64) *IndustryManufactureLabourCostResponseBodyList {
	s.GmtCreate = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetGmtModified(v int64) *IndustryManufactureLabourCostResponseBodyList {
	s.GmtModified = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetInstanceId(v string) *IndustryManufactureLabourCostResponseBodyList {
	s.InstanceId = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetLabourCostName(v string) *IndustryManufactureLabourCostResponseBodyList {
	s.LabourCostName = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetLabourCostNo(v string) *IndustryManufactureLabourCostResponseBodyList {
	s.LabourCostNo = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetMaterialName(v string) *IndustryManufactureLabourCostResponseBodyList {
	s.MaterialName = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetMaterialNo(v string) *IndustryManufactureLabourCostResponseBodyList {
	s.MaterialNo = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetProcessCode(v string) *IndustryManufactureLabourCostResponseBodyList {
	s.ProcessCode = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetProcessName(v string) *IndustryManufactureLabourCostResponseBodyList {
	s.ProcessName = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetProcessNo(v string) *IndustryManufactureLabourCostResponseBodyList {
	s.ProcessNo = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetQualifiedPrice(v float32) *IndustryManufactureLabourCostResponseBodyList {
	s.QualifiedPrice = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetUnQualifiedInfo(v string) *IndustryManufactureLabourCostResponseBodyList {
	s.UnQualifiedInfo = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetUnQualifiedPrice1(v float32) *IndustryManufactureLabourCostResponseBodyList {
	s.UnQualifiedPrice1 = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetUnQualifiedPrice2(v float32) *IndustryManufactureLabourCostResponseBodyList {
	s.UnQualifiedPrice2 = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetUnQualifiedReason1(v string) *IndustryManufactureLabourCostResponseBodyList {
	s.UnQualifiedReason1 = &v
	return s
}

func (s *IndustryManufactureLabourCostResponseBodyList) SetUnQualifiedReason2(v string) *IndustryManufactureLabourCostResponseBodyList {
	s.UnQualifiedReason2 = &v
	return s
}

type IndustryManufactureLabourCostResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IndustryManufactureLabourCostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndustryManufactureLabourCostResponse) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureLabourCostResponse) GoString() string {
	return s.String()
}

func (s *IndustryManufactureLabourCostResponse) SetHeaders(v map[string]*string) *IndustryManufactureLabourCostResponse {
	s.Headers = v
	return s
}

func (s *IndustryManufactureLabourCostResponse) SetStatusCode(v int32) *IndustryManufactureLabourCostResponse {
	s.StatusCode = &v
	return s
}

func (s *IndustryManufactureLabourCostResponse) SetBody(v *IndustryManufactureLabourCostResponseBody) *IndustryManufactureLabourCostResponse {
	s.Body = v
	return s
}

type IndustryManufactureMaterialListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IndustryManufactureMaterialListHeaders) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMaterialListHeaders) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMaterialListHeaders) SetCommonHeaders(v map[string]*string) *IndustryManufactureMaterialListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IndustryManufactureMaterialListHeaders) SetXAcsDingtalkAccessToken(v string) *IndustryManufactureMaterialListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IndustryManufactureMaterialListRequest struct {
	AppId           *int64   `json:"appId,omitempty" xml:"appId,omitempty"`
	AppIds          []*int64 `json:"appIds,omitempty" xml:"appIds,omitempty" type:"Repeated"`
	AppName         *string  `json:"appName,omitempty" xml:"appName,omitempty"`
	CorpId          *string  `json:"corpId,omitempty" xml:"corpId,omitempty"`
	CurrentPage     *int64   `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	Cursor          *int64   `json:"cursor,omitempty" xml:"cursor,omitempty"`
	EndTime         *int64   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	InstanceId      *string  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	IsvOrgId        *int64   `json:"isvOrgId,omitempty" xml:"isvOrgId,omitempty"`
	MaterialNo      *string  `json:"materialNo,omitempty" xml:"materialNo,omitempty"`
	MicroappAgentId *int64   `json:"microappAgentId,omitempty" xml:"microappAgentId,omitempty"`
	OrgId           *int64   `json:"orgId,omitempty" xml:"orgId,omitempty"`
	PageSize        *int32   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StartTime       *int64   `json:"startTime,omitempty" xml:"startTime,omitempty"`
	SuiteKey        *string  `json:"suiteKey,omitempty" xml:"suiteKey,omitempty"`
	TokenGrantType  *int64   `json:"tokenGrantType,omitempty" xml:"tokenGrantType,omitempty"`
}

func (s IndustryManufactureMaterialListRequest) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMaterialListRequest) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMaterialListRequest) SetAppId(v int64) *IndustryManufactureMaterialListRequest {
	s.AppId = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetAppIds(v []*int64) *IndustryManufactureMaterialListRequest {
	s.AppIds = v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetAppName(v string) *IndustryManufactureMaterialListRequest {
	s.AppName = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetCorpId(v string) *IndustryManufactureMaterialListRequest {
	s.CorpId = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetCurrentPage(v int64) *IndustryManufactureMaterialListRequest {
	s.CurrentPage = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetCursor(v int64) *IndustryManufactureMaterialListRequest {
	s.Cursor = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetEndTime(v int64) *IndustryManufactureMaterialListRequest {
	s.EndTime = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetInstanceId(v string) *IndustryManufactureMaterialListRequest {
	s.InstanceId = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetIsvOrgId(v int64) *IndustryManufactureMaterialListRequest {
	s.IsvOrgId = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetMaterialNo(v string) *IndustryManufactureMaterialListRequest {
	s.MaterialNo = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetMicroappAgentId(v int64) *IndustryManufactureMaterialListRequest {
	s.MicroappAgentId = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetOrgId(v int64) *IndustryManufactureMaterialListRequest {
	s.OrgId = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetPageSize(v int32) *IndustryManufactureMaterialListRequest {
	s.PageSize = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetStartTime(v int64) *IndustryManufactureMaterialListRequest {
	s.StartTime = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetSuiteKey(v string) *IndustryManufactureMaterialListRequest {
	s.SuiteKey = &v
	return s
}

func (s *IndustryManufactureMaterialListRequest) SetTokenGrantType(v int64) *IndustryManufactureMaterialListRequest {
	s.TokenGrantType = &v
	return s
}

type IndustryManufactureMaterialListResponseBody struct {
	HasMore    *bool                                              `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*IndustryManufactureMaterialListResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextCursor *int64                                             `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	TotalCount *int64                                             `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s IndustryManufactureMaterialListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMaterialListResponseBody) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMaterialListResponseBody) SetHasMore(v bool) *IndustryManufactureMaterialListResponseBody {
	s.HasMore = &v
	return s
}

func (s *IndustryManufactureMaterialListResponseBody) SetList(v []*IndustryManufactureMaterialListResponseBodyList) *IndustryManufactureMaterialListResponseBody {
	s.List = v
	return s
}

func (s *IndustryManufactureMaterialListResponseBody) SetNextCursor(v int64) *IndustryManufactureMaterialListResponseBody {
	s.NextCursor = &v
	return s
}

func (s *IndustryManufactureMaterialListResponseBody) SetTotalCount(v int64) *IndustryManufactureMaterialListResponseBody {
	s.TotalCount = &v
	return s
}

type IndustryManufactureMaterialListResponseBodyList struct {
	Category      *string  `json:"category,omitempty" xml:"category,omitempty"`
	CorpId        *string  `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Ext           *string  `json:"ext,omitempty" xml:"ext,omitempty"`
	InstanceId    *string  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	MaterialName  *string  `json:"materialName,omitempty" xml:"materialName,omitempty"`
	MaterialNo    *string  `json:"materialNo,omitempty" xml:"materialNo,omitempty"`
	ProcessCode   *string  `json:"processCode,omitempty" xml:"processCode,omitempty"`
	Specification *string  `json:"specification,omitempty" xml:"specification,omitempty"`
	StockMaxWarn  *float32 `json:"stockMaxWarn,omitempty" xml:"stockMaxWarn,omitempty"`
	StockMinWarn  *float32 `json:"stockMinWarn,omitempty" xml:"stockMinWarn,omitempty"`
	Type          *string  `json:"type,omitempty" xml:"type,omitempty"`
	Unit          *string  `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s IndustryManufactureMaterialListResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMaterialListResponseBodyList) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMaterialListResponseBodyList) SetCategory(v string) *IndustryManufactureMaterialListResponseBodyList {
	s.Category = &v
	return s
}

func (s *IndustryManufactureMaterialListResponseBodyList) SetCorpId(v string) *IndustryManufactureMaterialListResponseBodyList {
	s.CorpId = &v
	return s
}

func (s *IndustryManufactureMaterialListResponseBodyList) SetExt(v string) *IndustryManufactureMaterialListResponseBodyList {
	s.Ext = &v
	return s
}

func (s *IndustryManufactureMaterialListResponseBodyList) SetInstanceId(v string) *IndustryManufactureMaterialListResponseBodyList {
	s.InstanceId = &v
	return s
}

func (s *IndustryManufactureMaterialListResponseBodyList) SetMaterialName(v string) *IndustryManufactureMaterialListResponseBodyList {
	s.MaterialName = &v
	return s
}

func (s *IndustryManufactureMaterialListResponseBodyList) SetMaterialNo(v string) *IndustryManufactureMaterialListResponseBodyList {
	s.MaterialNo = &v
	return s
}

func (s *IndustryManufactureMaterialListResponseBodyList) SetProcessCode(v string) *IndustryManufactureMaterialListResponseBodyList {
	s.ProcessCode = &v
	return s
}

func (s *IndustryManufactureMaterialListResponseBodyList) SetSpecification(v string) *IndustryManufactureMaterialListResponseBodyList {
	s.Specification = &v
	return s
}

func (s *IndustryManufactureMaterialListResponseBodyList) SetStockMaxWarn(v float32) *IndustryManufactureMaterialListResponseBodyList {
	s.StockMaxWarn = &v
	return s
}

func (s *IndustryManufactureMaterialListResponseBodyList) SetStockMinWarn(v float32) *IndustryManufactureMaterialListResponseBodyList {
	s.StockMinWarn = &v
	return s
}

func (s *IndustryManufactureMaterialListResponseBodyList) SetType(v string) *IndustryManufactureMaterialListResponseBodyList {
	s.Type = &v
	return s
}

func (s *IndustryManufactureMaterialListResponseBodyList) SetUnit(v string) *IndustryManufactureMaterialListResponseBodyList {
	s.Unit = &v
	return s
}

type IndustryManufactureMaterialListResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IndustryManufactureMaterialListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndustryManufactureMaterialListResponse) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMaterialListResponse) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMaterialListResponse) SetHeaders(v map[string]*string) *IndustryManufactureMaterialListResponse {
	s.Headers = v
	return s
}

func (s *IndustryManufactureMaterialListResponse) SetStatusCode(v int32) *IndustryManufactureMaterialListResponse {
	s.StatusCode = &v
	return s
}

func (s *IndustryManufactureMaterialListResponse) SetBody(v *IndustryManufactureMaterialListResponseBody) *IndustryManufactureMaterialListResponse {
	s.Body = v
	return s
}

type IndustryManufactureMesDispatchTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IndustryManufactureMesDispatchTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesDispatchTaskHeaders) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesDispatchTaskHeaders) SetCommonHeaders(v map[string]*string) *IndustryManufactureMesDispatchTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IndustryManufactureMesDispatchTaskHeaders) SetXAcsDingtalkAccessToken(v string) *IndustryManufactureMesDispatchTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IndustryManufactureMesDispatchTaskRequest struct {
	Action               *string `json:"action,omitempty" xml:"action,omitempty"`
	AppKey               *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	BaseDataName         *string `json:"baseDataName,omitempty" xml:"baseDataName,omitempty"`
	DefectsAmount        *string `json:"defectsAmount,omitempty" xml:"defectsAmount,omitempty"`
	DispatchStaffName    *string `json:"dispatchStaffName,omitempty" xml:"dispatchStaffName,omitempty"`
	DispatchStaffNo      *string `json:"dispatchStaffNo,omitempty" xml:"dispatchStaffNo,omitempty"`
	FineAmount           *string `json:"fineAmount,omitempty" xml:"fineAmount,omitempty"`
	Overdue              *int32  `json:"overdue,omitempty" xml:"overdue,omitempty"`
	PlanQuantity         *int64  `json:"planQuantity,omitempty" xml:"planQuantity,omitempty"`
	Priority             *int32  `json:"priority,omitempty" xml:"priority,omitempty"`
	ProcessName          *string `json:"processName,omitempty" xml:"processName,omitempty"`
	ProcessUuid          *string `json:"processUuid,omitempty" xml:"processUuid,omitempty"`
	ProductCode          *string `json:"productCode,omitempty" xml:"productCode,omitempty"`
	ProductName          *string `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductSpecification *string `json:"productSpecification,omitempty" xml:"productSpecification,omitempty"`
	ProjectCode          *string `json:"projectCode,omitempty" xml:"projectCode,omitempty"`
	ProjectId            *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	ProjectStatus        *string `json:"projectStatus,omitempty" xml:"projectStatus,omitempty"`
	TaskOperators        *string `json:"taskOperators,omitempty" xml:"taskOperators,omitempty"`
	TaskPlanEndTime      *string `json:"taskPlanEndTime,omitempty" xml:"taskPlanEndTime,omitempty"`
	TaskPlanStartTime    *string `json:"taskPlanStartTime,omitempty" xml:"taskPlanStartTime,omitempty"`
	TaskStatus           *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
	TaskType             *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	TeamId               *string `json:"teamId,omitempty" xml:"teamId,omitempty"`
	Uuid                 *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s IndustryManufactureMesDispatchTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesDispatchTaskRequest) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesDispatchTaskRequest) SetAction(v string) *IndustryManufactureMesDispatchTaskRequest {
	s.Action = &v
	return s
}

func (s *IndustryManufactureMesDispatchTaskRequest) SetAppKey(v string) *IndustryManufactureMesDispatchTaskRequest {
	s.AppKey = &v
	return s
}

func (s *IndustryManufactureMesDispatchTaskRequest) SetBaseDataName(v string) *IndustryManufactureMesDispatchTaskRequest {
	s.BaseDataName = &v
	return s
}

func (s *IndustryManufactureMesDispatchTaskRequest) SetDefectsAmount(v string) *IndustryManufactureMesDispatchTaskRequest {
	s.DefectsAmount = &v
	return s
}

func (s *IndustryManufactureMesDispatchTaskRequest) SetDispatchStaffName(v string) *IndustryManufactureMesDispatchTaskRequest {
	s.DispatchStaffName = &v
	return s
}

func (s *IndustryManufactureMesDispatchTaskRequest) SetDispatchStaffNo(v string) *IndustryManufactureMesDispatchTaskRequest {
	s.DispatchStaffNo = &v
	return s
}

func (s *IndustryManufactureMesDispatchTaskRequest) SetFineAmount(v string) *IndustryManufactureMesDispatchTaskRequest {
	s.FineAmount = &v
	return s
}

func (s *IndustryManufactureMesDispatchTaskRequest) SetOverdue(v int32) *IndustryManufactureMesDispatchTaskRequest {
	s.Overdue = &v
	return s
}

func (s *IndustryManufactureMesDispatchTaskRequest) SetPlanQuantity(v int64) *IndustryManufactureMesDispatchTaskRequest {
	s.PlanQuantity = &v
	return s
}

func (s *IndustryManufactureMesDispatchTaskRequest) SetPriority(v int32) *IndustryManufactureMesDispatchTaskRequest {
	s.Priority = &v
	return s
}

func (s *IndustryManufactureMesDispatchTaskRequest) SetProcessName(v string) *IndustryManufactureMesDispatchTaskRequest {
	s.ProcessName = &v
	return s
}

func (s *IndustryManufactureMesDispatchTaskRequest) SetProcessUuid(v string) *IndustryManufactureMesDispatchTaskRequest {
	s.ProcessUuid = &v
	return s
}

func (s *IndustryManufactureMesDispatchTaskRequest) SetProductCode(v string) *IndustryManufactureMesDispatchTaskRequest {
	s.ProductCode = &v
	return s
}

func (s *IndustryManufactureMesDispatchTaskRequest) SetProductName(v string) *IndustryManufactureMesDispatchTaskRequest {
	s.ProductName = &v
	return s
}

func (s *IndustryManufactureMesDispatchTaskRequest) SetProductSpecification(v string) *IndustryManufactureMesDispatchTaskRequest {
	s.ProductSpecification = &v
	return s
}

func (s *IndustryManufactureMesDispatchTaskRequest) SetProjectCode(v string) *IndustryManufactureMesDispatchTaskRequest {
	s.ProjectCode = &v
	return s
}

func (s *IndustryManufactureMesDispatchTaskRequest) SetProjectId(v string) *IndustryManufactureMesDispatchTaskRequest {
	s.ProjectId = &v
	return s
}

func (s *IndustryManufactureMesDispatchTaskRequest) SetProjectStatus(v string) *IndustryManufactureMesDispatchTaskRequest {
	s.ProjectStatus = &v
	return s
}

func (s *IndustryManufactureMesDispatchTaskRequest) SetTaskOperators(v string) *IndustryManufactureMesDispatchTaskRequest {
	s.TaskOperators = &v
	return s
}

func (s *IndustryManufactureMesDispatchTaskRequest) SetTaskPlanEndTime(v string) *IndustryManufactureMesDispatchTaskRequest {
	s.TaskPlanEndTime = &v
	return s
}

func (s *IndustryManufactureMesDispatchTaskRequest) SetTaskPlanStartTime(v string) *IndustryManufactureMesDispatchTaskRequest {
	s.TaskPlanStartTime = &v
	return s
}

func (s *IndustryManufactureMesDispatchTaskRequest) SetTaskStatus(v string) *IndustryManufactureMesDispatchTaskRequest {
	s.TaskStatus = &v
	return s
}

func (s *IndustryManufactureMesDispatchTaskRequest) SetTaskType(v string) *IndustryManufactureMesDispatchTaskRequest {
	s.TaskType = &v
	return s
}

func (s *IndustryManufactureMesDispatchTaskRequest) SetTeamId(v string) *IndustryManufactureMesDispatchTaskRequest {
	s.TeamId = &v
	return s
}

func (s *IndustryManufactureMesDispatchTaskRequest) SetUuid(v string) *IndustryManufactureMesDispatchTaskRequest {
	s.Uuid = &v
	return s
}

type IndustryManufactureMesDispatchTaskResponseBody struct {
	Result *IndustryManufactureMesDispatchTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s IndustryManufactureMesDispatchTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesDispatchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesDispatchTaskResponseBody) SetResult(v *IndustryManufactureMesDispatchTaskResponseBodyResult) *IndustryManufactureMesDispatchTaskResponseBody {
	s.Result = v
	return s
}

type IndustryManufactureMesDispatchTaskResponseBodyResult struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s IndustryManufactureMesDispatchTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesDispatchTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesDispatchTaskResponseBodyResult) SetContent(v string) *IndustryManufactureMesDispatchTaskResponseBodyResult {
	s.Content = &v
	return s
}

type IndustryManufactureMesDispatchTaskResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IndustryManufactureMesDispatchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndustryManufactureMesDispatchTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesDispatchTaskResponse) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesDispatchTaskResponse) SetHeaders(v map[string]*string) *IndustryManufactureMesDispatchTaskResponse {
	s.Headers = v
	return s
}

func (s *IndustryManufactureMesDispatchTaskResponse) SetStatusCode(v int32) *IndustryManufactureMesDispatchTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *IndustryManufactureMesDispatchTaskResponse) SetBody(v *IndustryManufactureMesDispatchTaskResponseBody) *IndustryManufactureMesDispatchTaskResponse {
	s.Body = v
	return s
}

type IndustryManufactureMesMaterialHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IndustryManufactureMesMaterialHeaders) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesMaterialHeaders) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesMaterialHeaders) SetCommonHeaders(v map[string]*string) *IndustryManufactureMesMaterialHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IndustryManufactureMesMaterialHeaders) SetXAcsDingtalkAccessToken(v string) *IndustryManufactureMesMaterialHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IndustryManufactureMesMaterialRequest struct {
	Action               *string                                            `json:"action,omitempty" xml:"action,omitempty"`
	AppKey               *string                                            `json:"appKey,omitempty" xml:"appKey,omitempty"`
	BaseDataName         *string                                            `json:"baseDataName,omitempty" xml:"baseDataName,omitempty"`
	Category             *string                                            `json:"category,omitempty" xml:"category,omitempty"`
	ExtendData           []*IndustryManufactureMesMaterialRequestExtendData `json:"extendData,omitempty" xml:"extendData,omitempty" type:"Repeated"`
	ProductCode          *string                                            `json:"productCode,omitempty" xml:"productCode,omitempty"`
	ProductName          *string                                            `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductSpecification *string                                            `json:"productSpecification,omitempty" xml:"productSpecification,omitempty"`
	Prop                 *string                                            `json:"prop,omitempty" xml:"prop,omitempty"`
	Unit                 *string                                            `json:"unit,omitempty" xml:"unit,omitempty"`
	Uuid                 *string                                            `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s IndustryManufactureMesMaterialRequest) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesMaterialRequest) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesMaterialRequest) SetAction(v string) *IndustryManufactureMesMaterialRequest {
	s.Action = &v
	return s
}

func (s *IndustryManufactureMesMaterialRequest) SetAppKey(v string) *IndustryManufactureMesMaterialRequest {
	s.AppKey = &v
	return s
}

func (s *IndustryManufactureMesMaterialRequest) SetBaseDataName(v string) *IndustryManufactureMesMaterialRequest {
	s.BaseDataName = &v
	return s
}

func (s *IndustryManufactureMesMaterialRequest) SetCategory(v string) *IndustryManufactureMesMaterialRequest {
	s.Category = &v
	return s
}

func (s *IndustryManufactureMesMaterialRequest) SetExtendData(v []*IndustryManufactureMesMaterialRequestExtendData) *IndustryManufactureMesMaterialRequest {
	s.ExtendData = v
	return s
}

func (s *IndustryManufactureMesMaterialRequest) SetProductCode(v string) *IndustryManufactureMesMaterialRequest {
	s.ProductCode = &v
	return s
}

func (s *IndustryManufactureMesMaterialRequest) SetProductName(v string) *IndustryManufactureMesMaterialRequest {
	s.ProductName = &v
	return s
}

func (s *IndustryManufactureMesMaterialRequest) SetProductSpecification(v string) *IndustryManufactureMesMaterialRequest {
	s.ProductSpecification = &v
	return s
}

func (s *IndustryManufactureMesMaterialRequest) SetProp(v string) *IndustryManufactureMesMaterialRequest {
	s.Prop = &v
	return s
}

func (s *IndustryManufactureMesMaterialRequest) SetUnit(v string) *IndustryManufactureMesMaterialRequest {
	s.Unit = &v
	return s
}

func (s *IndustryManufactureMesMaterialRequest) SetUuid(v string) *IndustryManufactureMesMaterialRequest {
	s.Uuid = &v
	return s
}

type IndustryManufactureMesMaterialRequestExtendData struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Value     *string `json:"value,omitempty" xml:"value,omitempty"`
	ValueType *string `json:"valueType,omitempty" xml:"valueType,omitempty"`
}

func (s IndustryManufactureMesMaterialRequestExtendData) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesMaterialRequestExtendData) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesMaterialRequestExtendData) SetCode(v string) *IndustryManufactureMesMaterialRequestExtendData {
	s.Code = &v
	return s
}

func (s *IndustryManufactureMesMaterialRequestExtendData) SetName(v string) *IndustryManufactureMesMaterialRequestExtendData {
	s.Name = &v
	return s
}

func (s *IndustryManufactureMesMaterialRequestExtendData) SetValue(v string) *IndustryManufactureMesMaterialRequestExtendData {
	s.Value = &v
	return s
}

func (s *IndustryManufactureMesMaterialRequestExtendData) SetValueType(v string) *IndustryManufactureMesMaterialRequestExtendData {
	s.ValueType = &v
	return s
}

type IndustryManufactureMesMaterialResponseBody struct {
	Result *IndustryManufactureMesMaterialResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s IndustryManufactureMesMaterialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesMaterialResponseBody) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesMaterialResponseBody) SetResult(v *IndustryManufactureMesMaterialResponseBodyResult) *IndustryManufactureMesMaterialResponseBody {
	s.Result = v
	return s
}

type IndustryManufactureMesMaterialResponseBodyResult struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s IndustryManufactureMesMaterialResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesMaterialResponseBodyResult) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesMaterialResponseBodyResult) SetContent(v string) *IndustryManufactureMesMaterialResponseBodyResult {
	s.Content = &v
	return s
}

type IndustryManufactureMesMaterialResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IndustryManufactureMesMaterialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndustryManufactureMesMaterialResponse) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesMaterialResponse) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesMaterialResponse) SetHeaders(v map[string]*string) *IndustryManufactureMesMaterialResponse {
	s.Headers = v
	return s
}

func (s *IndustryManufactureMesMaterialResponse) SetStatusCode(v int32) *IndustryManufactureMesMaterialResponse {
	s.StatusCode = &v
	return s
}

func (s *IndustryManufactureMesMaterialResponse) SetBody(v *IndustryManufactureMesMaterialResponseBody) *IndustryManufactureMesMaterialResponse {
	s.Body = v
	return s
}

type IndustryManufactureMesOutPlanHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IndustryManufactureMesOutPlanHeaders) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesOutPlanHeaders) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesOutPlanHeaders) SetCommonHeaders(v map[string]*string) *IndustryManufactureMesOutPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IndustryManufactureMesOutPlanHeaders) SetXAcsDingtalkAccessToken(v string) *IndustryManufactureMesOutPlanHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IndustryManufactureMesOutPlanRequest struct {
	ApprovalStatus            *string `json:"approvalStatus,omitempty" xml:"approvalStatus,omitempty"`
	Approver                  *string `json:"approver,omitempty" xml:"approver,omitempty"`
	BaseDataName              *string `json:"baseDataName,omitempty" xml:"baseDataName,omitempty"`
	OutSourceProjectCode      *string `json:"outSourceProjectCode,omitempty" xml:"outSourceProjectCode,omitempty"`
	OutSourceTeamId           *string `json:"outSourceTeamId,omitempty" xml:"outSourceTeamId,omitempty"`
	Price                     *string `json:"price,omitempty" xml:"price,omitempty"`
	ProcessIdentificationCode *string `json:"processIdentificationCode,omitempty" xml:"processIdentificationCode,omitempty"`
	ProcessUuids              *string `json:"processUuids,omitempty" xml:"processUuids,omitempty"`
	ProductCode               *string `json:"productCode,omitempty" xml:"productCode,omitempty"`
	ProductName               *string `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductSpecification      *string `json:"productSpecification,omitempty" xml:"productSpecification,omitempty"`
	ProjectCode               *string `json:"projectCode,omitempty" xml:"projectCode,omitempty"`
	ProjectId                 *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	SendPlanQuantity          *string `json:"sendPlanQuantity,omitempty" xml:"sendPlanQuantity,omitempty"`
	SupplierCode              *string `json:"supplierCode,omitempty" xml:"supplierCode,omitempty"`
	SupplierName              *string `json:"supplierName,omitempty" xml:"supplierName,omitempty"`
	TotalWage                 *string `json:"totalWage,omitempty" xml:"totalWage,omitempty"`
	Uuid                      *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s IndustryManufactureMesOutPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesOutPlanRequest) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesOutPlanRequest) SetApprovalStatus(v string) *IndustryManufactureMesOutPlanRequest {
	s.ApprovalStatus = &v
	return s
}

func (s *IndustryManufactureMesOutPlanRequest) SetApprover(v string) *IndustryManufactureMesOutPlanRequest {
	s.Approver = &v
	return s
}

func (s *IndustryManufactureMesOutPlanRequest) SetBaseDataName(v string) *IndustryManufactureMesOutPlanRequest {
	s.BaseDataName = &v
	return s
}

func (s *IndustryManufactureMesOutPlanRequest) SetOutSourceProjectCode(v string) *IndustryManufactureMesOutPlanRequest {
	s.OutSourceProjectCode = &v
	return s
}

func (s *IndustryManufactureMesOutPlanRequest) SetOutSourceTeamId(v string) *IndustryManufactureMesOutPlanRequest {
	s.OutSourceTeamId = &v
	return s
}

func (s *IndustryManufactureMesOutPlanRequest) SetPrice(v string) *IndustryManufactureMesOutPlanRequest {
	s.Price = &v
	return s
}

func (s *IndustryManufactureMesOutPlanRequest) SetProcessIdentificationCode(v string) *IndustryManufactureMesOutPlanRequest {
	s.ProcessIdentificationCode = &v
	return s
}

func (s *IndustryManufactureMesOutPlanRequest) SetProcessUuids(v string) *IndustryManufactureMesOutPlanRequest {
	s.ProcessUuids = &v
	return s
}

func (s *IndustryManufactureMesOutPlanRequest) SetProductCode(v string) *IndustryManufactureMesOutPlanRequest {
	s.ProductCode = &v
	return s
}

func (s *IndustryManufactureMesOutPlanRequest) SetProductName(v string) *IndustryManufactureMesOutPlanRequest {
	s.ProductName = &v
	return s
}

func (s *IndustryManufactureMesOutPlanRequest) SetProductSpecification(v string) *IndustryManufactureMesOutPlanRequest {
	s.ProductSpecification = &v
	return s
}

func (s *IndustryManufactureMesOutPlanRequest) SetProjectCode(v string) *IndustryManufactureMesOutPlanRequest {
	s.ProjectCode = &v
	return s
}

func (s *IndustryManufactureMesOutPlanRequest) SetProjectId(v string) *IndustryManufactureMesOutPlanRequest {
	s.ProjectId = &v
	return s
}

func (s *IndustryManufactureMesOutPlanRequest) SetSendPlanQuantity(v string) *IndustryManufactureMesOutPlanRequest {
	s.SendPlanQuantity = &v
	return s
}

func (s *IndustryManufactureMesOutPlanRequest) SetSupplierCode(v string) *IndustryManufactureMesOutPlanRequest {
	s.SupplierCode = &v
	return s
}

func (s *IndustryManufactureMesOutPlanRequest) SetSupplierName(v string) *IndustryManufactureMesOutPlanRequest {
	s.SupplierName = &v
	return s
}

func (s *IndustryManufactureMesOutPlanRequest) SetTotalWage(v string) *IndustryManufactureMesOutPlanRequest {
	s.TotalWage = &v
	return s
}

func (s *IndustryManufactureMesOutPlanRequest) SetUuid(v string) *IndustryManufactureMesOutPlanRequest {
	s.Uuid = &v
	return s
}

type IndustryManufactureMesOutPlanResponseBody struct {
	Result *IndustryManufactureMesOutPlanResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s IndustryManufactureMesOutPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesOutPlanResponseBody) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesOutPlanResponseBody) SetResult(v *IndustryManufactureMesOutPlanResponseBodyResult) *IndustryManufactureMesOutPlanResponseBody {
	s.Result = v
	return s
}

type IndustryManufactureMesOutPlanResponseBodyResult struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s IndustryManufactureMesOutPlanResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesOutPlanResponseBodyResult) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesOutPlanResponseBodyResult) SetContent(v string) *IndustryManufactureMesOutPlanResponseBodyResult {
	s.Content = &v
	return s
}

type IndustryManufactureMesOutPlanResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IndustryManufactureMesOutPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndustryManufactureMesOutPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesOutPlanResponse) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesOutPlanResponse) SetHeaders(v map[string]*string) *IndustryManufactureMesOutPlanResponse {
	s.Headers = v
	return s
}

func (s *IndustryManufactureMesOutPlanResponse) SetStatusCode(v int32) *IndustryManufactureMesOutPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *IndustryManufactureMesOutPlanResponse) SetBody(v *IndustryManufactureMesOutPlanResponseBody) *IndustryManufactureMesOutPlanResponse {
	s.Body = v
	return s
}

type IndustryManufactureMesOutputHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IndustryManufactureMesOutputHeaders) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesOutputHeaders) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesOutputHeaders) SetCommonHeaders(v map[string]*string) *IndustryManufactureMesOutputHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IndustryManufactureMesOutputHeaders) SetXAcsDingtalkAccessToken(v string) *IndustryManufactureMesOutputHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IndustryManufactureMesOutputRequest struct {
	Action               *string `json:"action,omitempty" xml:"action,omitempty"`
	AppKey               *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	ApproveStatus        *string `json:"approveStatus,omitempty" xml:"approveStatus,omitempty"`
	BaseDataName         *string `json:"baseDataName,omitempty" xml:"baseDataName,omitempty"`
	DefectsAmount        *string `json:"defectsAmount,omitempty" xml:"defectsAmount,omitempty"`
	DefectsReason        *string `json:"defectsReason,omitempty" xml:"defectsReason,omitempty"`
	FineAmount           *string `json:"fineAmount,omitempty" xml:"fineAmount,omitempty"`
	HasQualityTest       *string `json:"hasQualityTest,omitempty" xml:"hasQualityTest,omitempty"`
	Overdue              *int32  `json:"overdue,omitempty" xml:"overdue,omitempty"`
	PlanQuantity         *int64  `json:"planQuantity,omitempty" xml:"planQuantity,omitempty"`
	Priority             *int32  `json:"priority,omitempty" xml:"priority,omitempty"`
	ProcessName          *string `json:"processName,omitempty" xml:"processName,omitempty"`
	ProcessUuid          *string `json:"processUuid,omitempty" xml:"processUuid,omitempty"`
	ProductCode          *string `json:"productCode,omitempty" xml:"productCode,omitempty"`
	ProductName          *string `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductSpecification *string `json:"productSpecification,omitempty" xml:"productSpecification,omitempty"`
	ProjectCode          *string `json:"projectCode,omitempty" xml:"projectCode,omitempty"`
	ProjectId            *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	ProjectStatus        *string `json:"projectStatus,omitempty" xml:"projectStatus,omitempty"`
	QualityTestStatus    *string `json:"qualityTestStatus,omitempty" xml:"qualityTestStatus,omitempty"`
	TaskPlanEndTime      *string `json:"taskPlanEndTime,omitempty" xml:"taskPlanEndTime,omitempty"`
	TaskPlanStartTime    *string `json:"taskPlanStartTime,omitempty" xml:"taskPlanStartTime,omitempty"`
	TaskStatus           *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
	TaskType             *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	TaskUuid             *string `json:"taskUuid,omitempty" xml:"taskUuid,omitempty"`
	TeamId               *string `json:"teamId,omitempty" xml:"teamId,omitempty"`
	UserId               *string `json:"userId,omitempty" xml:"userId,omitempty"`
	UserName             *string `json:"userName,omitempty" xml:"userName,omitempty"`
	Uuid                 *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s IndustryManufactureMesOutputRequest) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesOutputRequest) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesOutputRequest) SetAction(v string) *IndustryManufactureMesOutputRequest {
	s.Action = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetAppKey(v string) *IndustryManufactureMesOutputRequest {
	s.AppKey = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetApproveStatus(v string) *IndustryManufactureMesOutputRequest {
	s.ApproveStatus = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetBaseDataName(v string) *IndustryManufactureMesOutputRequest {
	s.BaseDataName = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetDefectsAmount(v string) *IndustryManufactureMesOutputRequest {
	s.DefectsAmount = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetDefectsReason(v string) *IndustryManufactureMesOutputRequest {
	s.DefectsReason = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetFineAmount(v string) *IndustryManufactureMesOutputRequest {
	s.FineAmount = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetHasQualityTest(v string) *IndustryManufactureMesOutputRequest {
	s.HasQualityTest = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetOverdue(v int32) *IndustryManufactureMesOutputRequest {
	s.Overdue = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetPlanQuantity(v int64) *IndustryManufactureMesOutputRequest {
	s.PlanQuantity = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetPriority(v int32) *IndustryManufactureMesOutputRequest {
	s.Priority = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetProcessName(v string) *IndustryManufactureMesOutputRequest {
	s.ProcessName = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetProcessUuid(v string) *IndustryManufactureMesOutputRequest {
	s.ProcessUuid = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetProductCode(v string) *IndustryManufactureMesOutputRequest {
	s.ProductCode = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetProductName(v string) *IndustryManufactureMesOutputRequest {
	s.ProductName = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetProductSpecification(v string) *IndustryManufactureMesOutputRequest {
	s.ProductSpecification = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetProjectCode(v string) *IndustryManufactureMesOutputRequest {
	s.ProjectCode = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetProjectId(v string) *IndustryManufactureMesOutputRequest {
	s.ProjectId = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetProjectStatus(v string) *IndustryManufactureMesOutputRequest {
	s.ProjectStatus = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetQualityTestStatus(v string) *IndustryManufactureMesOutputRequest {
	s.QualityTestStatus = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetTaskPlanEndTime(v string) *IndustryManufactureMesOutputRequest {
	s.TaskPlanEndTime = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetTaskPlanStartTime(v string) *IndustryManufactureMesOutputRequest {
	s.TaskPlanStartTime = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetTaskStatus(v string) *IndustryManufactureMesOutputRequest {
	s.TaskStatus = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetTaskType(v string) *IndustryManufactureMesOutputRequest {
	s.TaskType = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetTaskUuid(v string) *IndustryManufactureMesOutputRequest {
	s.TaskUuid = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetTeamId(v string) *IndustryManufactureMesOutputRequest {
	s.TeamId = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetUserId(v string) *IndustryManufactureMesOutputRequest {
	s.UserId = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetUserName(v string) *IndustryManufactureMesOutputRequest {
	s.UserName = &v
	return s
}

func (s *IndustryManufactureMesOutputRequest) SetUuid(v string) *IndustryManufactureMesOutputRequest {
	s.Uuid = &v
	return s
}

type IndustryManufactureMesOutputResponseBody struct {
	Result *IndustryManufactureMesOutputResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s IndustryManufactureMesOutputResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesOutputResponseBody) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesOutputResponseBody) SetResult(v *IndustryManufactureMesOutputResponseBodyResult) *IndustryManufactureMesOutputResponseBody {
	s.Result = v
	return s
}

type IndustryManufactureMesOutputResponseBodyResult struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s IndustryManufactureMesOutputResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesOutputResponseBodyResult) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesOutputResponseBodyResult) SetContent(v string) *IndustryManufactureMesOutputResponseBodyResult {
	s.Content = &v
	return s
}

type IndustryManufactureMesOutputResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IndustryManufactureMesOutputResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndustryManufactureMesOutputResponse) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesOutputResponse) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesOutputResponse) SetHeaders(v map[string]*string) *IndustryManufactureMesOutputResponse {
	s.Headers = v
	return s
}

func (s *IndustryManufactureMesOutputResponse) SetStatusCode(v int32) *IndustryManufactureMesOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *IndustryManufactureMesOutputResponse) SetBody(v *IndustryManufactureMesOutputResponseBody) *IndustryManufactureMesOutputResponse {
	s.Body = v
	return s
}

type IndustryManufactureMesProcessHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IndustryManufactureMesProcessHeaders) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesProcessHeaders) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesProcessHeaders) SetCommonHeaders(v map[string]*string) *IndustryManufactureMesProcessHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IndustryManufactureMesProcessHeaders) SetXAcsDingtalkAccessToken(v string) *IndustryManufactureMesProcessHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IndustryManufactureMesProcessRequest struct {
	Action          *string                                           `json:"action,omitempty" xml:"action,omitempty"`
	AppKey          *string                                           `json:"appKey,omitempty" xml:"appKey,omitempty"`
	BaseDataName    *string                                           `json:"baseDataName,omitempty" xml:"baseDataName,omitempty"`
	ExtendData      []*IndustryManufactureMesProcessRequestExtendData `json:"extendData,omitempty" xml:"extendData,omitempty" type:"Repeated"`
	Name            *string                                           `json:"name,omitempty" xml:"name,omitempty"`
	NeedDispatch    *string                                           `json:"needDispatch,omitempty" xml:"needDispatch,omitempty"`
	NeedQualityTest *string                                           `json:"needQualityTest,omitempty" xml:"needQualityTest,omitempty"`
	No              *string                                           `json:"no,omitempty" xml:"no,omitempty"`
	Price           *string                                           `json:"price,omitempty" xml:"price,omitempty"`
	Prop            *string                                           `json:"prop,omitempty" xml:"prop,omitempty"`
	Remark          *string                                           `json:"remark,omitempty" xml:"remark,omitempty"`
	Sop             *string                                           `json:"sop,omitempty" xml:"sop,omitempty"`
	Uuid            *string                                           `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s IndustryManufactureMesProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesProcessRequest) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesProcessRequest) SetAction(v string) *IndustryManufactureMesProcessRequest {
	s.Action = &v
	return s
}

func (s *IndustryManufactureMesProcessRequest) SetAppKey(v string) *IndustryManufactureMesProcessRequest {
	s.AppKey = &v
	return s
}

func (s *IndustryManufactureMesProcessRequest) SetBaseDataName(v string) *IndustryManufactureMesProcessRequest {
	s.BaseDataName = &v
	return s
}

func (s *IndustryManufactureMesProcessRequest) SetExtendData(v []*IndustryManufactureMesProcessRequestExtendData) *IndustryManufactureMesProcessRequest {
	s.ExtendData = v
	return s
}

func (s *IndustryManufactureMesProcessRequest) SetName(v string) *IndustryManufactureMesProcessRequest {
	s.Name = &v
	return s
}

func (s *IndustryManufactureMesProcessRequest) SetNeedDispatch(v string) *IndustryManufactureMesProcessRequest {
	s.NeedDispatch = &v
	return s
}

func (s *IndustryManufactureMesProcessRequest) SetNeedQualityTest(v string) *IndustryManufactureMesProcessRequest {
	s.NeedQualityTest = &v
	return s
}

func (s *IndustryManufactureMesProcessRequest) SetNo(v string) *IndustryManufactureMesProcessRequest {
	s.No = &v
	return s
}

func (s *IndustryManufactureMesProcessRequest) SetPrice(v string) *IndustryManufactureMesProcessRequest {
	s.Price = &v
	return s
}

func (s *IndustryManufactureMesProcessRequest) SetProp(v string) *IndustryManufactureMesProcessRequest {
	s.Prop = &v
	return s
}

func (s *IndustryManufactureMesProcessRequest) SetRemark(v string) *IndustryManufactureMesProcessRequest {
	s.Remark = &v
	return s
}

func (s *IndustryManufactureMesProcessRequest) SetSop(v string) *IndustryManufactureMesProcessRequest {
	s.Sop = &v
	return s
}

func (s *IndustryManufactureMesProcessRequest) SetUuid(v string) *IndustryManufactureMesProcessRequest {
	s.Uuid = &v
	return s
}

type IndustryManufactureMesProcessRequestExtendData struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Value     *string `json:"value,omitempty" xml:"value,omitempty"`
	ValueType *string `json:"valueType,omitempty" xml:"valueType,omitempty"`
}

func (s IndustryManufactureMesProcessRequestExtendData) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesProcessRequestExtendData) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesProcessRequestExtendData) SetCode(v string) *IndustryManufactureMesProcessRequestExtendData {
	s.Code = &v
	return s
}

func (s *IndustryManufactureMesProcessRequestExtendData) SetName(v string) *IndustryManufactureMesProcessRequestExtendData {
	s.Name = &v
	return s
}

func (s *IndustryManufactureMesProcessRequestExtendData) SetValue(v string) *IndustryManufactureMesProcessRequestExtendData {
	s.Value = &v
	return s
}

func (s *IndustryManufactureMesProcessRequestExtendData) SetValueType(v string) *IndustryManufactureMesProcessRequestExtendData {
	s.ValueType = &v
	return s
}

type IndustryManufactureMesProcessResponseBody struct {
	Result *IndustryManufactureMesProcessResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s IndustryManufactureMesProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesProcessResponseBody) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesProcessResponseBody) SetResult(v *IndustryManufactureMesProcessResponseBodyResult) *IndustryManufactureMesProcessResponseBody {
	s.Result = v
	return s
}

type IndustryManufactureMesProcessResponseBodyResult struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s IndustryManufactureMesProcessResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesProcessResponseBodyResult) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesProcessResponseBodyResult) SetContent(v string) *IndustryManufactureMesProcessResponseBodyResult {
	s.Content = &v
	return s
}

type IndustryManufactureMesProcessResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IndustryManufactureMesProcessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndustryManufactureMesProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesProcessResponse) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesProcessResponse) SetHeaders(v map[string]*string) *IndustryManufactureMesProcessResponse {
	s.Headers = v
	return s
}

func (s *IndustryManufactureMesProcessResponse) SetStatusCode(v int32) *IndustryManufactureMesProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *IndustryManufactureMesProcessResponse) SetBody(v *IndustryManufactureMesProcessResponseBody) *IndustryManufactureMesProcessResponse {
	s.Body = v
	return s
}

type IndustryManufactureMesProductionPlanHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IndustryManufactureMesProductionPlanHeaders) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesProductionPlanHeaders) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesProductionPlanHeaders) SetCommonHeaders(v map[string]*string) *IndustryManufactureMesProductionPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IndustryManufactureMesProductionPlanHeaders) SetXAcsDingtalkAccessToken(v string) *IndustryManufactureMesProductionPlanHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IndustryManufactureMesProductionPlanRequest struct {
	Action               *string                                                  `json:"action,omitempty" xml:"action,omitempty"`
	ActualEndTime        *string                                                  `json:"actualEndTime,omitempty" xml:"actualEndTime,omitempty"`
	ActualStartTime      *string                                                  `json:"actualStartTime,omitempty" xml:"actualStartTime,omitempty"`
	AppKey               *string                                                  `json:"appKey,omitempty" xml:"appKey,omitempty"`
	BaseDataName         *string                                                  `json:"baseDataName,omitempty" xml:"baseDataName,omitempty"`
	BomUuid              *string                                                  `json:"bomUuid,omitempty" xml:"bomUuid,omitempty"`
	Events               []*string                                                `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
	ExtendData           []*IndustryManufactureMesProductionPlanRequestExtendData `json:"extendData,omitempty" xml:"extendData,omitempty" type:"Repeated"`
	No                   *string                                                  `json:"no,omitempty" xml:"no,omitempty"`
	Overdue              *string                                                  `json:"overdue,omitempty" xml:"overdue,omitempty"`
	PlanEndTime          *string                                                  `json:"planEndTime,omitempty" xml:"planEndTime,omitempty"`
	PlanQuantity         *string                                                  `json:"planQuantity,omitempty" xml:"planQuantity,omitempty"`
	PlanStartTime        *string                                                  `json:"planStartTime,omitempty" xml:"planStartTime,omitempty"`
	ProcessUuids         *string                                                  `json:"processUuids,omitempty" xml:"processUuids,omitempty"`
	ProductCode          *string                                                  `json:"productCode,omitempty" xml:"productCode,omitempty"`
	ProductName          *string                                                  `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductSpecification *string                                                  `json:"productSpecification,omitempty" xml:"productSpecification,omitempty"`
	QualifiedQuantity    *string                                                  `json:"qualifiedQuantity,omitempty" xml:"qualifiedQuantity,omitempty"`
	SellOrderNo          *string                                                  `json:"sellOrderNo,omitempty" xml:"sellOrderNo,omitempty"`
	Status               *string                                                  `json:"status,omitempty" xml:"status,omitempty"`
	TeamList             *string                                                  `json:"teamList,omitempty" xml:"teamList,omitempty"`
	Title                *string                                                  `json:"title,omitempty" xml:"title,omitempty"`
	Type                 *string                                                  `json:"type,omitempty" xml:"type,omitempty"`
	Unit                 *string                                                  `json:"unit,omitempty" xml:"unit,omitempty"`
	Uuid                 *string                                                  `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s IndustryManufactureMesProductionPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesProductionPlanRequest) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesProductionPlanRequest) SetAction(v string) *IndustryManufactureMesProductionPlanRequest {
	s.Action = &v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequest) SetActualEndTime(v string) *IndustryManufactureMesProductionPlanRequest {
	s.ActualEndTime = &v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequest) SetActualStartTime(v string) *IndustryManufactureMesProductionPlanRequest {
	s.ActualStartTime = &v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequest) SetAppKey(v string) *IndustryManufactureMesProductionPlanRequest {
	s.AppKey = &v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequest) SetBaseDataName(v string) *IndustryManufactureMesProductionPlanRequest {
	s.BaseDataName = &v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequest) SetBomUuid(v string) *IndustryManufactureMesProductionPlanRequest {
	s.BomUuid = &v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequest) SetEvents(v []*string) *IndustryManufactureMesProductionPlanRequest {
	s.Events = v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequest) SetExtendData(v []*IndustryManufactureMesProductionPlanRequestExtendData) *IndustryManufactureMesProductionPlanRequest {
	s.ExtendData = v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequest) SetNo(v string) *IndustryManufactureMesProductionPlanRequest {
	s.No = &v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequest) SetOverdue(v string) *IndustryManufactureMesProductionPlanRequest {
	s.Overdue = &v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequest) SetPlanEndTime(v string) *IndustryManufactureMesProductionPlanRequest {
	s.PlanEndTime = &v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequest) SetPlanQuantity(v string) *IndustryManufactureMesProductionPlanRequest {
	s.PlanQuantity = &v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequest) SetPlanStartTime(v string) *IndustryManufactureMesProductionPlanRequest {
	s.PlanStartTime = &v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequest) SetProcessUuids(v string) *IndustryManufactureMesProductionPlanRequest {
	s.ProcessUuids = &v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequest) SetProductCode(v string) *IndustryManufactureMesProductionPlanRequest {
	s.ProductCode = &v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequest) SetProductName(v string) *IndustryManufactureMesProductionPlanRequest {
	s.ProductName = &v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequest) SetProductSpecification(v string) *IndustryManufactureMesProductionPlanRequest {
	s.ProductSpecification = &v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequest) SetQualifiedQuantity(v string) *IndustryManufactureMesProductionPlanRequest {
	s.QualifiedQuantity = &v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequest) SetSellOrderNo(v string) *IndustryManufactureMesProductionPlanRequest {
	s.SellOrderNo = &v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequest) SetStatus(v string) *IndustryManufactureMesProductionPlanRequest {
	s.Status = &v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequest) SetTeamList(v string) *IndustryManufactureMesProductionPlanRequest {
	s.TeamList = &v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequest) SetTitle(v string) *IndustryManufactureMesProductionPlanRequest {
	s.Title = &v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequest) SetType(v string) *IndustryManufactureMesProductionPlanRequest {
	s.Type = &v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequest) SetUnit(v string) *IndustryManufactureMesProductionPlanRequest {
	s.Unit = &v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequest) SetUuid(v string) *IndustryManufactureMesProductionPlanRequest {
	s.Uuid = &v
	return s
}

type IndustryManufactureMesProductionPlanRequestExtendData struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Value     *string `json:"value,omitempty" xml:"value,omitempty"`
	ValueType *string `json:"valueType,omitempty" xml:"valueType,omitempty"`
}

func (s IndustryManufactureMesProductionPlanRequestExtendData) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesProductionPlanRequestExtendData) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesProductionPlanRequestExtendData) SetCode(v string) *IndustryManufactureMesProductionPlanRequestExtendData {
	s.Code = &v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequestExtendData) SetName(v string) *IndustryManufactureMesProductionPlanRequestExtendData {
	s.Name = &v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequestExtendData) SetValue(v string) *IndustryManufactureMesProductionPlanRequestExtendData {
	s.Value = &v
	return s
}

func (s *IndustryManufactureMesProductionPlanRequestExtendData) SetValueType(v string) *IndustryManufactureMesProductionPlanRequestExtendData {
	s.ValueType = &v
	return s
}

type IndustryManufactureMesProductionPlanResponseBody struct {
	Result *IndustryManufactureMesProductionPlanResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s IndustryManufactureMesProductionPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesProductionPlanResponseBody) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesProductionPlanResponseBody) SetResult(v *IndustryManufactureMesProductionPlanResponseBodyResult) *IndustryManufactureMesProductionPlanResponseBody {
	s.Result = v
	return s
}

type IndustryManufactureMesProductionPlanResponseBodyResult struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s IndustryManufactureMesProductionPlanResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesProductionPlanResponseBodyResult) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesProductionPlanResponseBodyResult) SetContent(v string) *IndustryManufactureMesProductionPlanResponseBodyResult {
	s.Content = &v
	return s
}

type IndustryManufactureMesProductionPlanResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IndustryManufactureMesProductionPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndustryManufactureMesProductionPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesProductionPlanResponse) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesProductionPlanResponse) SetHeaders(v map[string]*string) *IndustryManufactureMesProductionPlanResponse {
	s.Headers = v
	return s
}

func (s *IndustryManufactureMesProductionPlanResponse) SetStatusCode(v int32) *IndustryManufactureMesProductionPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *IndustryManufactureMesProductionPlanResponse) SetBody(v *IndustryManufactureMesProductionPlanResponseBody) *IndustryManufactureMesProductionPlanResponse {
	s.Body = v
	return s
}

type IndustryManufactureMesSubCooperationTeamHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IndustryManufactureMesSubCooperationTeamHeaders) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesSubCooperationTeamHeaders) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesSubCooperationTeamHeaders) SetCommonHeaders(v map[string]*string) *IndustryManufactureMesSubCooperationTeamHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IndustryManufactureMesSubCooperationTeamHeaders) SetXAcsDingtalkAccessToken(v string) *IndustryManufactureMesSubCooperationTeamHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IndustryManufactureMesSubCooperationTeamRequest struct {
	Action       *string                                                        `json:"action,omitempty" xml:"action,omitempty"`
	AppKey       *string                                                        `json:"appKey,omitempty" xml:"appKey,omitempty"`
	BaseDataName *string                                                        `json:"baseDataName,omitempty" xml:"baseDataName,omitempty"`
	Events       []*string                                                      `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
	ExtendData   []*IndustryManufactureMesSubCooperationTeamRequestExtendData   `json:"extendData,omitempty" xml:"extendData,omitempty" type:"Repeated"`
	GroupPlugins []*IndustryManufactureMesSubCooperationTeamRequestGroupPlugins `json:"groupPlugins,omitempty" xml:"groupPlugins,omitempty" type:"Repeated"`
	GroupType    *string                                                        `json:"groupType,omitempty" xml:"groupType,omitempty"`
	Leaders      []*IndustryManufactureMesSubCooperationTeamRequestLeaders      `json:"leaders,omitempty" xml:"leaders,omitempty" type:"Repeated"`
	Members      []*IndustryManufactureMesSubCooperationTeamRequestMembers      `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	Name         *string                                                        `json:"name,omitempty" xml:"name,omitempty"`
	OutCorpId    *string                                                        `json:"outCorpId,omitempty" xml:"outCorpId,omitempty"`
	ProcessIds   []*string                                                      `json:"processIds,omitempty" xml:"processIds,omitempty" type:"Repeated"`
	Uuid         *string                                                        `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s IndustryManufactureMesSubCooperationTeamRequest) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesSubCooperationTeamRequest) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesSubCooperationTeamRequest) SetAction(v string) *IndustryManufactureMesSubCooperationTeamRequest {
	s.Action = &v
	return s
}

func (s *IndustryManufactureMesSubCooperationTeamRequest) SetAppKey(v string) *IndustryManufactureMesSubCooperationTeamRequest {
	s.AppKey = &v
	return s
}

func (s *IndustryManufactureMesSubCooperationTeamRequest) SetBaseDataName(v string) *IndustryManufactureMesSubCooperationTeamRequest {
	s.BaseDataName = &v
	return s
}

func (s *IndustryManufactureMesSubCooperationTeamRequest) SetEvents(v []*string) *IndustryManufactureMesSubCooperationTeamRequest {
	s.Events = v
	return s
}

func (s *IndustryManufactureMesSubCooperationTeamRequest) SetExtendData(v []*IndustryManufactureMesSubCooperationTeamRequestExtendData) *IndustryManufactureMesSubCooperationTeamRequest {
	s.ExtendData = v
	return s
}

func (s *IndustryManufactureMesSubCooperationTeamRequest) SetGroupPlugins(v []*IndustryManufactureMesSubCooperationTeamRequestGroupPlugins) *IndustryManufactureMesSubCooperationTeamRequest {
	s.GroupPlugins = v
	return s
}

func (s *IndustryManufactureMesSubCooperationTeamRequest) SetGroupType(v string) *IndustryManufactureMesSubCooperationTeamRequest {
	s.GroupType = &v
	return s
}

func (s *IndustryManufactureMesSubCooperationTeamRequest) SetLeaders(v []*IndustryManufactureMesSubCooperationTeamRequestLeaders) *IndustryManufactureMesSubCooperationTeamRequest {
	s.Leaders = v
	return s
}

func (s *IndustryManufactureMesSubCooperationTeamRequest) SetMembers(v []*IndustryManufactureMesSubCooperationTeamRequestMembers) *IndustryManufactureMesSubCooperationTeamRequest {
	s.Members = v
	return s
}

func (s *IndustryManufactureMesSubCooperationTeamRequest) SetName(v string) *IndustryManufactureMesSubCooperationTeamRequest {
	s.Name = &v
	return s
}

func (s *IndustryManufactureMesSubCooperationTeamRequest) SetOutCorpId(v string) *IndustryManufactureMesSubCooperationTeamRequest {
	s.OutCorpId = &v
	return s
}

func (s *IndustryManufactureMesSubCooperationTeamRequest) SetProcessIds(v []*string) *IndustryManufactureMesSubCooperationTeamRequest {
	s.ProcessIds = v
	return s
}

func (s *IndustryManufactureMesSubCooperationTeamRequest) SetUuid(v string) *IndustryManufactureMesSubCooperationTeamRequest {
	s.Uuid = &v
	return s
}

type IndustryManufactureMesSubCooperationTeamRequestExtendData struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Value     *string `json:"value,omitempty" xml:"value,omitempty"`
	ValueType *string `json:"valueType,omitempty" xml:"valueType,omitempty"`
}

func (s IndustryManufactureMesSubCooperationTeamRequestExtendData) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesSubCooperationTeamRequestExtendData) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesSubCooperationTeamRequestExtendData) SetCode(v string) *IndustryManufactureMesSubCooperationTeamRequestExtendData {
	s.Code = &v
	return s
}

func (s *IndustryManufactureMesSubCooperationTeamRequestExtendData) SetName(v string) *IndustryManufactureMesSubCooperationTeamRequestExtendData {
	s.Name = &v
	return s
}

func (s *IndustryManufactureMesSubCooperationTeamRequestExtendData) SetValue(v string) *IndustryManufactureMesSubCooperationTeamRequestExtendData {
	s.Value = &v
	return s
}

func (s *IndustryManufactureMesSubCooperationTeamRequestExtendData) SetValueType(v string) *IndustryManufactureMesSubCooperationTeamRequestExtendData {
	s.ValueType = &v
	return s
}

type IndustryManufactureMesSubCooperationTeamRequestGroupPlugins struct {
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s IndustryManufactureMesSubCooperationTeamRequestGroupPlugins) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesSubCooperationTeamRequestGroupPlugins) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesSubCooperationTeamRequestGroupPlugins) SetLabel(v string) *IndustryManufactureMesSubCooperationTeamRequestGroupPlugins {
	s.Label = &v
	return s
}

func (s *IndustryManufactureMesSubCooperationTeamRequestGroupPlugins) SetValue(v string) *IndustryManufactureMesSubCooperationTeamRequestGroupPlugins {
	s.Value = &v
	return s
}

type IndustryManufactureMesSubCooperationTeamRequestLeaders struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s IndustryManufactureMesSubCooperationTeamRequestLeaders) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesSubCooperationTeamRequestLeaders) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesSubCooperationTeamRequestLeaders) SetName(v string) *IndustryManufactureMesSubCooperationTeamRequestLeaders {
	s.Name = &v
	return s
}

func (s *IndustryManufactureMesSubCooperationTeamRequestLeaders) SetUserId(v string) *IndustryManufactureMesSubCooperationTeamRequestLeaders {
	s.UserId = &v
	return s
}

type IndustryManufactureMesSubCooperationTeamRequestMembers struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s IndustryManufactureMesSubCooperationTeamRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesSubCooperationTeamRequestMembers) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesSubCooperationTeamRequestMembers) SetName(v string) *IndustryManufactureMesSubCooperationTeamRequestMembers {
	s.Name = &v
	return s
}

func (s *IndustryManufactureMesSubCooperationTeamRequestMembers) SetUserId(v string) *IndustryManufactureMesSubCooperationTeamRequestMembers {
	s.UserId = &v
	return s
}

type IndustryManufactureMesSubCooperationTeamResponseBody struct {
	Result *IndustryManufactureMesSubCooperationTeamResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s IndustryManufactureMesSubCooperationTeamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesSubCooperationTeamResponseBody) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesSubCooperationTeamResponseBody) SetResult(v *IndustryManufactureMesSubCooperationTeamResponseBodyResult) *IndustryManufactureMesSubCooperationTeamResponseBody {
	s.Result = v
	return s
}

type IndustryManufactureMesSubCooperationTeamResponseBodyResult struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s IndustryManufactureMesSubCooperationTeamResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesSubCooperationTeamResponseBodyResult) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesSubCooperationTeamResponseBodyResult) SetContent(v string) *IndustryManufactureMesSubCooperationTeamResponseBodyResult {
	s.Content = &v
	return s
}

type IndustryManufactureMesSubCooperationTeamResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IndustryManufactureMesSubCooperationTeamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndustryManufactureMesSubCooperationTeamResponse) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesSubCooperationTeamResponse) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesSubCooperationTeamResponse) SetHeaders(v map[string]*string) *IndustryManufactureMesSubCooperationTeamResponse {
	s.Headers = v
	return s
}

func (s *IndustryManufactureMesSubCooperationTeamResponse) SetStatusCode(v int32) *IndustryManufactureMesSubCooperationTeamResponse {
	s.StatusCode = &v
	return s
}

func (s *IndustryManufactureMesSubCooperationTeamResponse) SetBody(v *IndustryManufactureMesSubCooperationTeamResponseBody) *IndustryManufactureMesSubCooperationTeamResponse {
	s.Body = v
	return s
}

type IndustryManufactureMesTeamMgmtHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IndustryManufactureMesTeamMgmtHeaders) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesTeamMgmtHeaders) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesTeamMgmtHeaders) SetCommonHeaders(v map[string]*string) *IndustryManufactureMesTeamMgmtHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IndustryManufactureMesTeamMgmtHeaders) SetXAcsDingtalkAccessToken(v string) *IndustryManufactureMesTeamMgmtHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IndustryManufactureMesTeamMgmtRequest struct {
	Action       *string                                              `json:"action,omitempty" xml:"action,omitempty"`
	AppKey       *string                                              `json:"appKey,omitempty" xml:"appKey,omitempty"`
	BaseDataName *string                                              `json:"baseDataName,omitempty" xml:"baseDataName,omitempty"`
	Events       []*string                                            `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
	ExtendData   []*IndustryManufactureMesTeamMgmtRequestExtendData   `json:"extendData,omitempty" xml:"extendData,omitempty" type:"Repeated"`
	GroupConfig  map[string]interface{}                               `json:"groupConfig,omitempty" xml:"groupConfig,omitempty"`
	GroupPlugins []*IndustryManufactureMesTeamMgmtRequestGroupPlugins `json:"groupPlugins,omitempty" xml:"groupPlugins,omitempty" type:"Repeated"`
	GroupType    *string                                              `json:"groupType,omitempty" xml:"groupType,omitempty"`
	Id           *string                                              `json:"id,omitempty" xml:"id,omitempty"`
	Leaders      []*IndustryManufactureMesTeamMgmtRequestLeaders      `json:"leaders,omitempty" xml:"leaders,omitempty" type:"Repeated"`
	Members      []*IndustryManufactureMesTeamMgmtRequestMembers      `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	Name         *string                                              `json:"name,omitempty" xml:"name,omitempty"`
	ProcessIds   []*string                                            `json:"processIds,omitempty" xml:"processIds,omitempty" type:"Repeated"`
	TagKey       *string                                              `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	TagValues    []*string                                            `json:"tagValues,omitempty" xml:"tagValues,omitempty" type:"Repeated"`
}

func (s IndustryManufactureMesTeamMgmtRequest) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesTeamMgmtRequest) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesTeamMgmtRequest) SetAction(v string) *IndustryManufactureMesTeamMgmtRequest {
	s.Action = &v
	return s
}

func (s *IndustryManufactureMesTeamMgmtRequest) SetAppKey(v string) *IndustryManufactureMesTeamMgmtRequest {
	s.AppKey = &v
	return s
}

func (s *IndustryManufactureMesTeamMgmtRequest) SetBaseDataName(v string) *IndustryManufactureMesTeamMgmtRequest {
	s.BaseDataName = &v
	return s
}

func (s *IndustryManufactureMesTeamMgmtRequest) SetEvents(v []*string) *IndustryManufactureMesTeamMgmtRequest {
	s.Events = v
	return s
}

func (s *IndustryManufactureMesTeamMgmtRequest) SetExtendData(v []*IndustryManufactureMesTeamMgmtRequestExtendData) *IndustryManufactureMesTeamMgmtRequest {
	s.ExtendData = v
	return s
}

func (s *IndustryManufactureMesTeamMgmtRequest) SetGroupConfig(v map[string]interface{}) *IndustryManufactureMesTeamMgmtRequest {
	s.GroupConfig = v
	return s
}

func (s *IndustryManufactureMesTeamMgmtRequest) SetGroupPlugins(v []*IndustryManufactureMesTeamMgmtRequestGroupPlugins) *IndustryManufactureMesTeamMgmtRequest {
	s.GroupPlugins = v
	return s
}

func (s *IndustryManufactureMesTeamMgmtRequest) SetGroupType(v string) *IndustryManufactureMesTeamMgmtRequest {
	s.GroupType = &v
	return s
}

func (s *IndustryManufactureMesTeamMgmtRequest) SetId(v string) *IndustryManufactureMesTeamMgmtRequest {
	s.Id = &v
	return s
}

func (s *IndustryManufactureMesTeamMgmtRequest) SetLeaders(v []*IndustryManufactureMesTeamMgmtRequestLeaders) *IndustryManufactureMesTeamMgmtRequest {
	s.Leaders = v
	return s
}

func (s *IndustryManufactureMesTeamMgmtRequest) SetMembers(v []*IndustryManufactureMesTeamMgmtRequestMembers) *IndustryManufactureMesTeamMgmtRequest {
	s.Members = v
	return s
}

func (s *IndustryManufactureMesTeamMgmtRequest) SetName(v string) *IndustryManufactureMesTeamMgmtRequest {
	s.Name = &v
	return s
}

func (s *IndustryManufactureMesTeamMgmtRequest) SetProcessIds(v []*string) *IndustryManufactureMesTeamMgmtRequest {
	s.ProcessIds = v
	return s
}

func (s *IndustryManufactureMesTeamMgmtRequest) SetTagKey(v string) *IndustryManufactureMesTeamMgmtRequest {
	s.TagKey = &v
	return s
}

func (s *IndustryManufactureMesTeamMgmtRequest) SetTagValues(v []*string) *IndustryManufactureMesTeamMgmtRequest {
	s.TagValues = v
	return s
}

type IndustryManufactureMesTeamMgmtRequestExtendData struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Value     *string `json:"value,omitempty" xml:"value,omitempty"`
	ValueType *string `json:"valueType,omitempty" xml:"valueType,omitempty"`
}

func (s IndustryManufactureMesTeamMgmtRequestExtendData) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesTeamMgmtRequestExtendData) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesTeamMgmtRequestExtendData) SetCode(v string) *IndustryManufactureMesTeamMgmtRequestExtendData {
	s.Code = &v
	return s
}

func (s *IndustryManufactureMesTeamMgmtRequestExtendData) SetName(v string) *IndustryManufactureMesTeamMgmtRequestExtendData {
	s.Name = &v
	return s
}

func (s *IndustryManufactureMesTeamMgmtRequestExtendData) SetValue(v string) *IndustryManufactureMesTeamMgmtRequestExtendData {
	s.Value = &v
	return s
}

func (s *IndustryManufactureMesTeamMgmtRequestExtendData) SetValueType(v string) *IndustryManufactureMesTeamMgmtRequestExtendData {
	s.ValueType = &v
	return s
}

type IndustryManufactureMesTeamMgmtRequestGroupPlugins struct {
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s IndustryManufactureMesTeamMgmtRequestGroupPlugins) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesTeamMgmtRequestGroupPlugins) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesTeamMgmtRequestGroupPlugins) SetLabel(v string) *IndustryManufactureMesTeamMgmtRequestGroupPlugins {
	s.Label = &v
	return s
}

func (s *IndustryManufactureMesTeamMgmtRequestGroupPlugins) SetValue(v string) *IndustryManufactureMesTeamMgmtRequestGroupPlugins {
	s.Value = &v
	return s
}

type IndustryManufactureMesTeamMgmtRequestLeaders struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s IndustryManufactureMesTeamMgmtRequestLeaders) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesTeamMgmtRequestLeaders) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesTeamMgmtRequestLeaders) SetName(v string) *IndustryManufactureMesTeamMgmtRequestLeaders {
	s.Name = &v
	return s
}

func (s *IndustryManufactureMesTeamMgmtRequestLeaders) SetUserId(v string) *IndustryManufactureMesTeamMgmtRequestLeaders {
	s.UserId = &v
	return s
}

type IndustryManufactureMesTeamMgmtRequestMembers struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s IndustryManufactureMesTeamMgmtRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesTeamMgmtRequestMembers) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesTeamMgmtRequestMembers) SetName(v string) *IndustryManufactureMesTeamMgmtRequestMembers {
	s.Name = &v
	return s
}

func (s *IndustryManufactureMesTeamMgmtRequestMembers) SetUserId(v string) *IndustryManufactureMesTeamMgmtRequestMembers {
	s.UserId = &v
	return s
}

type IndustryManufactureMesTeamMgmtResponseBody struct {
	DingOpenErrcode *int32                                            `json:"dingOpenErrcode,omitempty" xml:"dingOpenErrcode,omitempty"`
	ErrorMsg        *string                                           `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Result          *IndustryManufactureMesTeamMgmtResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s IndustryManufactureMesTeamMgmtResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesTeamMgmtResponseBody) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesTeamMgmtResponseBody) SetDingOpenErrcode(v int32) *IndustryManufactureMesTeamMgmtResponseBody {
	s.DingOpenErrcode = &v
	return s
}

func (s *IndustryManufactureMesTeamMgmtResponseBody) SetErrorMsg(v string) *IndustryManufactureMesTeamMgmtResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *IndustryManufactureMesTeamMgmtResponseBody) SetResult(v *IndustryManufactureMesTeamMgmtResponseBodyResult) *IndustryManufactureMesTeamMgmtResponseBody {
	s.Result = v
	return s
}

type IndustryManufactureMesTeamMgmtResponseBodyResult struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s IndustryManufactureMesTeamMgmtResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesTeamMgmtResponseBodyResult) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesTeamMgmtResponseBodyResult) SetContent(v string) *IndustryManufactureMesTeamMgmtResponseBodyResult {
	s.Content = &v
	return s
}

type IndustryManufactureMesTeamMgmtResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IndustryManufactureMesTeamMgmtResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndustryManufactureMesTeamMgmtResponse) String() string {
	return tea.Prettify(s)
}

func (s IndustryManufactureMesTeamMgmtResponse) GoString() string {
	return s.String()
}

func (s *IndustryManufactureMesTeamMgmtResponse) SetHeaders(v map[string]*string) *IndustryManufactureMesTeamMgmtResponse {
	s.Headers = v
	return s
}

func (s *IndustryManufactureMesTeamMgmtResponse) SetStatusCode(v int32) *IndustryManufactureMesTeamMgmtResponse {
	s.StatusCode = &v
	return s
}

func (s *IndustryManufactureMesTeamMgmtResponse) SetBody(v *IndustryManufactureMesTeamMgmtResponseBody) *IndustryManufactureMesTeamMgmtResponse {
	s.Body = v
	return s
}

type IndustryMmanufactureMaterialCostGetHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s IndustryMmanufactureMaterialCostGetHeaders) String() string {
	return tea.Prettify(s)
}

func (s IndustryMmanufactureMaterialCostGetHeaders) GoString() string {
	return s.String()
}

func (s *IndustryMmanufactureMaterialCostGetHeaders) SetCommonHeaders(v map[string]*string) *IndustryMmanufactureMaterialCostGetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetHeaders) SetXAcsDingtalkAccessToken(v string) *IndustryMmanufactureMaterialCostGetHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type IndustryMmanufactureMaterialCostGetRequest struct {
	AppId           *int64   `json:"appId,omitempty" xml:"appId,omitempty"`
	AppIds          []*int64 `json:"appIds,omitempty" xml:"appIds,omitempty" type:"Repeated"`
	AppName         *string  `json:"appName,omitempty" xml:"appName,omitempty"`
	CorpId          *string  `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Cursor          *int64   `json:"cursor,omitempty" xml:"cursor,omitempty"`
	EndTime         *int64   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	InstanceId      *string  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	IsvOrgId        *int64   `json:"isvOrgId,omitempty" xml:"isvOrgId,omitempty"`
	MaterialNo      *string  `json:"materialNo,omitempty" xml:"materialNo,omitempty"`
	MicroappAgentId *int64   `json:"microappAgentId,omitempty" xml:"microappAgentId,omitempty"`
	OrgId           *int64   `json:"orgId,omitempty" xml:"orgId,omitempty"`
	PageNumber      *int64   `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize        *int32   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	StartTime       *int64   `json:"startTime,omitempty" xml:"startTime,omitempty"`
	SuiteKey        *string  `json:"suiteKey,omitempty" xml:"suiteKey,omitempty"`
	TokenGrantType  *int32   `json:"tokenGrantType,omitempty" xml:"tokenGrantType,omitempty"`
}

func (s IndustryMmanufactureMaterialCostGetRequest) String() string {
	return tea.Prettify(s)
}

func (s IndustryMmanufactureMaterialCostGetRequest) GoString() string {
	return s.String()
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetAppId(v int64) *IndustryMmanufactureMaterialCostGetRequest {
	s.AppId = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetAppIds(v []*int64) *IndustryMmanufactureMaterialCostGetRequest {
	s.AppIds = v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetAppName(v string) *IndustryMmanufactureMaterialCostGetRequest {
	s.AppName = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetCorpId(v string) *IndustryMmanufactureMaterialCostGetRequest {
	s.CorpId = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetCursor(v int64) *IndustryMmanufactureMaterialCostGetRequest {
	s.Cursor = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetEndTime(v int64) *IndustryMmanufactureMaterialCostGetRequest {
	s.EndTime = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetInstanceId(v string) *IndustryMmanufactureMaterialCostGetRequest {
	s.InstanceId = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetIsvOrgId(v int64) *IndustryMmanufactureMaterialCostGetRequest {
	s.IsvOrgId = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetMaterialNo(v string) *IndustryMmanufactureMaterialCostGetRequest {
	s.MaterialNo = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetMicroappAgentId(v int64) *IndustryMmanufactureMaterialCostGetRequest {
	s.MicroappAgentId = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetOrgId(v int64) *IndustryMmanufactureMaterialCostGetRequest {
	s.OrgId = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetPageNumber(v int64) *IndustryMmanufactureMaterialCostGetRequest {
	s.PageNumber = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetPageSize(v int32) *IndustryMmanufactureMaterialCostGetRequest {
	s.PageSize = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetStartTime(v int64) *IndustryMmanufactureMaterialCostGetRequest {
	s.StartTime = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetSuiteKey(v string) *IndustryMmanufactureMaterialCostGetRequest {
	s.SuiteKey = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetRequest) SetTokenGrantType(v int32) *IndustryMmanufactureMaterialCostGetRequest {
	s.TokenGrantType = &v
	return s
}

type IndustryMmanufactureMaterialCostGetResponseBody struct {
	HasMore    *bool                                                  `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List       []*IndustryMmanufactureMaterialCostGetResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextCursor *int64                                                 `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
	TotalCount *int64                                                 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s IndustryMmanufactureMaterialCostGetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IndustryMmanufactureMaterialCostGetResponseBody) GoString() string {
	return s.String()
}

func (s *IndustryMmanufactureMaterialCostGetResponseBody) SetHasMore(v bool) *IndustryMmanufactureMaterialCostGetResponseBody {
	s.HasMore = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBody) SetList(v []*IndustryMmanufactureMaterialCostGetResponseBodyList) *IndustryMmanufactureMaterialCostGetResponseBody {
	s.List = v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBody) SetNextCursor(v int64) *IndustryMmanufactureMaterialCostGetResponseBody {
	s.NextCursor = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBody) SetTotalCount(v int64) *IndustryMmanufactureMaterialCostGetResponseBody {
	s.TotalCount = &v
	return s
}

type IndustryMmanufactureMaterialCostGetResponseBodyList struct {
	ActPrice       *float32 `json:"actPrice,omitempty" xml:"actPrice,omitempty"`
	CorpId         *string  `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Count          *float32 `json:"count,omitempty" xml:"count,omitempty"`
	Ext            *string  `json:"ext,omitempty" xml:"ext,omitempty"`
	GmtCreate      *int64   `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *int64   `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	InstanceId     *string  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	MaterialCostNo *string  `json:"materialCostNo,omitempty" xml:"materialCostNo,omitempty"`
	MaterialName   *string  `json:"materialName,omitempty" xml:"materialName,omitempty"`
	MaterialNo     *string  `json:"materialNo,omitempty" xml:"materialNo,omitempty"`
	Memo           *string  `json:"memo,omitempty" xml:"memo,omitempty"`
	Price          *float32 `json:"price,omitempty" xml:"price,omitempty"`
	ProcessCode    *string  `json:"processCode,omitempty" xml:"processCode,omitempty"`
	Unit           *string  `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s IndustryMmanufactureMaterialCostGetResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s IndustryMmanufactureMaterialCostGetResponseBodyList) GoString() string {
	return s.String()
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetActPrice(v float32) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.ActPrice = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetCorpId(v string) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.CorpId = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetCount(v float32) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.Count = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetExt(v string) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.Ext = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetGmtCreate(v int64) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.GmtCreate = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetGmtModified(v int64) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.GmtModified = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetInstanceId(v string) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.InstanceId = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetMaterialCostNo(v string) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.MaterialCostNo = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetMaterialName(v string) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.MaterialName = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetMaterialNo(v string) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.MaterialNo = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetMemo(v string) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.Memo = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetPrice(v float32) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.Price = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetProcessCode(v string) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.ProcessCode = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponseBodyList) SetUnit(v string) *IndustryMmanufactureMaterialCostGetResponseBodyList {
	s.Unit = &v
	return s
}

type IndustryMmanufactureMaterialCostGetResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IndustryMmanufactureMaterialCostGetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IndustryMmanufactureMaterialCostGetResponse) String() string {
	return tea.Prettify(s)
}

func (s IndustryMmanufactureMaterialCostGetResponse) GoString() string {
	return s.String()
}

func (s *IndustryMmanufactureMaterialCostGetResponse) SetHeaders(v map[string]*string) *IndustryMmanufactureMaterialCostGetResponse {
	s.Headers = v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponse) SetStatusCode(v int32) *IndustryMmanufactureMaterialCostGetResponse {
	s.StatusCode = &v
	return s
}

func (s *IndustryMmanufactureMaterialCostGetResponse) SetBody(v *IndustryMmanufactureMaterialCostGetResponseBody) *IndustryMmanufactureMaterialCostGetResponse {
	s.Body = v
	return s
}

type PushDingMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PushDingMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s PushDingMessageHeaders) GoString() string {
	return s.String()
}

func (s *PushDingMessageHeaders) SetCommonHeaders(v map[string]*string) *PushDingMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PushDingMessageHeaders) SetXAcsDingtalkAccessToken(v string) *PushDingMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PushDingMessageRequest struct {
	AppId       *int64    `json:"appId,omitempty" xml:"appId,omitempty"`
	Content     *string   `json:"content,omitempty" xml:"content,omitempty"`
	MessageType *string   `json:"messageType,omitempty" xml:"messageType,omitempty"`
	MessageUrl  *string   `json:"messageUrl,omitempty" xml:"messageUrl,omitempty"`
	PictureUrl  *string   `json:"pictureUrl,omitempty" xml:"pictureUrl,omitempty"`
	SingleTitle *string   `json:"singleTitle,omitempty" xml:"singleTitle,omitempty"`
	SingleUrl   *string   `json:"singleUrl,omitempty" xml:"singleUrl,omitempty"`
	Title       *string   `json:"title,omitempty" xml:"title,omitempty"`
	UserIdList  []*string `json:"userIdList,omitempty" xml:"userIdList,omitempty" type:"Repeated"`
}

func (s PushDingMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s PushDingMessageRequest) GoString() string {
	return s.String()
}

func (s *PushDingMessageRequest) SetAppId(v int64) *PushDingMessageRequest {
	s.AppId = &v
	return s
}

func (s *PushDingMessageRequest) SetContent(v string) *PushDingMessageRequest {
	s.Content = &v
	return s
}

func (s *PushDingMessageRequest) SetMessageType(v string) *PushDingMessageRequest {
	s.MessageType = &v
	return s
}

func (s *PushDingMessageRequest) SetMessageUrl(v string) *PushDingMessageRequest {
	s.MessageUrl = &v
	return s
}

func (s *PushDingMessageRequest) SetPictureUrl(v string) *PushDingMessageRequest {
	s.PictureUrl = &v
	return s
}

func (s *PushDingMessageRequest) SetSingleTitle(v string) *PushDingMessageRequest {
	s.SingleTitle = &v
	return s
}

func (s *PushDingMessageRequest) SetSingleUrl(v string) *PushDingMessageRequest {
	s.SingleUrl = &v
	return s
}

func (s *PushDingMessageRequest) SetTitle(v string) *PushDingMessageRequest {
	s.Title = &v
	return s
}

func (s *PushDingMessageRequest) SetUserIdList(v []*string) *PushDingMessageRequest {
	s.UserIdList = v
	return s
}

type PushDingMessageResponseBody struct {
	Content *int64 `json:"content,omitempty" xml:"content,omitempty"`
	Success *bool  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PushDingMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushDingMessageResponseBody) GoString() string {
	return s.String()
}

func (s *PushDingMessageResponseBody) SetContent(v int64) *PushDingMessageResponseBody {
	s.Content = &v
	return s
}

func (s *PushDingMessageResponseBody) SetSuccess(v bool) *PushDingMessageResponseBody {
	s.Success = &v
	return s
}

type PushDingMessageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PushDingMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushDingMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s PushDingMessageResponse) GoString() string {
	return s.String()
}

func (s *PushDingMessageResponse) SetHeaders(v map[string]*string) *PushDingMessageResponse {
	s.Headers = v
	return s
}

func (s *PushDingMessageResponse) SetStatusCode(v int32) *PushDingMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *PushDingMessageResponse) SetBody(v *PushDingMessageResponseBody) *PushDingMessageResponse {
	s.Body = v
	return s
}

type QueryAllDepartmentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllDepartmentHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentHeaders) SetCommonHeaders(v map[string]*string) *QueryAllDepartmentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllDepartmentHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllDepartmentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllDepartmentRequest struct {
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryAllDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentRequest) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentRequest) SetPageNumber(v int32) *QueryAllDepartmentRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryAllDepartmentRequest) SetPageSize(v int32) *QueryAllDepartmentRequest {
	s.PageSize = &v
	return s
}

type QueryAllDepartmentResponseBody struct {
	Content     []*QueryAllDepartmentResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	CurrentPage *int32                                   `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	TotalCount  *int64                                   `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	TotalPages  *int32                                   `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
}

func (s QueryAllDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentResponseBody) SetContent(v []*QueryAllDepartmentResponseBodyContent) *QueryAllDepartmentResponseBody {
	s.Content = v
	return s
}

func (s *QueryAllDepartmentResponseBody) SetCurrentPage(v int32) *QueryAllDepartmentResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryAllDepartmentResponseBody) SetTotalCount(v int64) *QueryAllDepartmentResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryAllDepartmentResponseBody) SetTotalPages(v int32) *QueryAllDepartmentResponseBody {
	s.TotalPages = &v
	return s
}

type QueryAllDepartmentResponseBodyContent struct {
	DeptAndExt      *QueryAllDepartmentResponseBodyContentDeptAndExt        `json:"deptAndExt,omitempty" xml:"deptAndExt,omitempty" type:"Struct"`
	GroupAndExtList []*QueryAllDepartmentResponseBodyContentGroupAndExtList `json:"groupAndExtList,omitempty" xml:"groupAndExtList,omitempty" type:"Repeated"`
	Id              *int64                                                  `json:"id,omitempty" xml:"id,omitempty"`
	Name            *string                                                 `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryAllDepartmentResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentResponseBodyContent) SetDeptAndExt(v *QueryAllDepartmentResponseBodyContentDeptAndExt) *QueryAllDepartmentResponseBodyContent {
	s.DeptAndExt = v
	return s
}

func (s *QueryAllDepartmentResponseBodyContent) SetGroupAndExtList(v []*QueryAllDepartmentResponseBodyContentGroupAndExtList) *QueryAllDepartmentResponseBodyContent {
	s.GroupAndExtList = v
	return s
}

func (s *QueryAllDepartmentResponseBodyContent) SetId(v int64) *QueryAllDepartmentResponseBodyContent {
	s.Id = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContent) SetName(v string) *QueryAllDepartmentResponseBodyContent {
	s.Name = &v
	return s
}

type QueryAllDepartmentResponseBodyContentDeptAndExt struct {
	Department  *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment    `json:"department,omitempty" xml:"department,omitempty" type:"Struct"`
	ExtendInfos []*QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos `json:"extendInfos,omitempty" xml:"extendInfos,omitempty" type:"Repeated"`
}

func (s QueryAllDepartmentResponseBodyContentDeptAndExt) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentResponseBodyContentDeptAndExt) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExt) SetDepartment(v *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) *QueryAllDepartmentResponseBodyContentDeptAndExt {
	s.Department = v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExt) SetExtendInfos(v []*QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos) *QueryAllDepartmentResponseBodyContentDeptAndExt {
	s.ExtendInfos = v
	return s
}

type QueryAllDepartmentResponseBodyContentDeptAndExtDepartment struct {
	DeptCode       *string  `json:"deptCode,omitempty" xml:"deptCode,omitempty"`
	DeptName       *string  `json:"deptName,omitempty" xml:"deptName,omitempty"`
	DeptOrder      *int64   `json:"deptOrder,omitempty" xml:"deptOrder,omitempty"`
	DeptStatus     *int32   `json:"deptStatus,omitempty" xml:"deptStatus,omitempty"`
	DeptType       *int32   `json:"deptType,omitempty" xml:"deptType,omitempty"`
	GmtCreateStr   *string  `json:"gmtCreateStr,omitempty" xml:"gmtCreateStr,omitempty"`
	GmtModifiedStr *string  `json:"gmtModifiedStr,omitempty" xml:"gmtModifiedStr,omitempty"`
	Id             *int64   `json:"id,omitempty" xml:"id,omitempty"`
	Name           *string  `json:"name,omitempty" xml:"name,omitempty"`
	ParentDeptCode *string  `json:"parentDeptCode,omitempty" xml:"parentDeptCode,omitempty"`
	Remark         *string  `json:"remark,omitempty" xml:"remark,omitempty"`
	WardIdList     []*int64 `json:"wardIdList,omitempty" xml:"wardIdList,omitempty" type:"Repeated"`
}

func (s QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) SetDeptCode(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment {
	s.DeptCode = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) SetDeptName(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment {
	s.DeptName = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) SetDeptOrder(v int64) *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment {
	s.DeptOrder = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) SetDeptStatus(v int32) *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment {
	s.DeptStatus = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) SetDeptType(v int32) *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment {
	s.DeptType = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) SetGmtCreateStr(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment {
	s.GmtCreateStr = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) SetGmtModifiedStr(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment {
	s.GmtModifiedStr = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) SetId(v int64) *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment {
	s.Id = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) SetName(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment {
	s.Name = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) SetParentDeptCode(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment {
	s.ParentDeptCode = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) SetRemark(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment {
	s.Remark = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment) SetWardIdList(v []*int64) *QueryAllDepartmentResponseBodyContentDeptAndExtDepartment {
	s.WardIdList = v
	return s
}

type QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos struct {
	DeptCode              *string `json:"deptCode,omitempty" xml:"deptCode,omitempty"`
	DeptExtendDisplayName *string `json:"deptExtendDisplayName,omitempty" xml:"deptExtendDisplayName,omitempty"`
	DeptExtendKey         *string `json:"deptExtendKey,omitempty" xml:"deptExtendKey,omitempty"`
	DeptExtendValue       *string `json:"deptExtendValue,omitempty" xml:"deptExtendValue,omitempty"`
	GmtCreateStr          *string `json:"gmtCreateStr,omitempty" xml:"gmtCreateStr,omitempty"`
	GmtModifiedStr        *string `json:"gmtModifiedStr,omitempty" xml:"gmtModifiedStr,omitempty"`
	Id                    *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Status                *int32  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos) SetDeptCode(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos {
	s.DeptCode = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos) SetDeptExtendDisplayName(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos {
	s.DeptExtendDisplayName = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos) SetDeptExtendKey(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos {
	s.DeptExtendKey = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos) SetDeptExtendValue(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos {
	s.DeptExtendValue = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos) SetGmtCreateStr(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos {
	s.GmtCreateStr = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos) SetGmtModifiedStr(v string) *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos {
	s.GmtModifiedStr = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos) SetId(v int64) *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos {
	s.Id = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos) SetStatus(v int32) *QueryAllDepartmentResponseBodyContentDeptAndExtExtendInfos {
	s.Status = &v
	return s
}

type QueryAllDepartmentResponseBodyContentGroupAndExtList struct {
	ExtendInfos []*QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos `json:"extendInfos,omitempty" xml:"extendInfos,omitempty" type:"Repeated"`
	Group       *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup         `json:"group,omitempty" xml:"group,omitempty" type:"Struct"`
}

func (s QueryAllDepartmentResponseBodyContentGroupAndExtList) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentResponseBodyContentGroupAndExtList) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtList) SetExtendInfos(v []*QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos) *QueryAllDepartmentResponseBodyContentGroupAndExtList {
	s.ExtendInfos = v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtList) SetGroup(v *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup) *QueryAllDepartmentResponseBodyContentGroupAndExtList {
	s.Group = v
	return s
}

type QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos struct {
	DeptCode              *string `json:"deptCode,omitempty" xml:"deptCode,omitempty"`
	DeptExtendDisplayName *string `json:"deptExtendDisplayName,omitempty" xml:"deptExtendDisplayName,omitempty"`
	DeptExtendKey         *string `json:"deptExtendKey,omitempty" xml:"deptExtendKey,omitempty"`
	DeptExtendValue       *string `json:"deptExtendValue,omitempty" xml:"deptExtendValue,omitempty"`
	GmtCreateStr          *string `json:"gmtCreateStr,omitempty" xml:"gmtCreateStr,omitempty"`
	GmtModifiedStr        *string `json:"gmtModifiedStr,omitempty" xml:"gmtModifiedStr,omitempty"`
	Id                    *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Status                *int32  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos) SetDeptCode(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos {
	s.DeptCode = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos) SetDeptExtendDisplayName(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos {
	s.DeptExtendDisplayName = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos) SetDeptExtendKey(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos {
	s.DeptExtendKey = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos) SetDeptExtendValue(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos {
	s.DeptExtendValue = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos) SetGmtCreateStr(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos {
	s.GmtCreateStr = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos) SetGmtModifiedStr(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos {
	s.GmtModifiedStr = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos) SetId(v int64) *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos {
	s.Id = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos) SetStatus(v int32) *QueryAllDepartmentResponseBodyContentGroupAndExtListExtendInfos {
	s.Status = &v
	return s
}

type QueryAllDepartmentResponseBodyContentGroupAndExtListGroup struct {
	DeptId         *int64                                                           `json:"deptId,omitempty" xml:"deptId,omitempty"`
	DeptStatus     *int32                                                           `json:"deptStatus,omitempty" xml:"deptStatus,omitempty"`
	GmtCreateStr   *string                                                          `json:"gmtCreateStr,omitempty" xml:"gmtCreateStr,omitempty"`
	GmtModifiedStr *string                                                          `json:"gmtModifiedStr,omitempty" xml:"gmtModifiedStr,omitempty"`
	Id             *int64                                                           `json:"id,omitempty" xml:"id,omitempty"`
	Leader         *QueryAllDepartmentResponseBodyContentGroupAndExtListGroupLeader `json:"leader,omitempty" xml:"leader,omitempty" type:"Struct"`
	Name           *string                                                          `json:"name,omitempty" xml:"name,omitempty"`
	ParentDeptCode *string                                                          `json:"parentDeptCode,omitempty" xml:"parentDeptCode,omitempty"`
	Remark         *string                                                          `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s QueryAllDepartmentResponseBodyContentGroupAndExtListGroup) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentResponseBodyContentGroupAndExtListGroup) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup) SetDeptId(v int64) *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup {
	s.DeptId = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup) SetDeptStatus(v int32) *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup {
	s.DeptStatus = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup) SetGmtCreateStr(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup {
	s.GmtCreateStr = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup) SetGmtModifiedStr(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup {
	s.GmtModifiedStr = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup) SetId(v int64) *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup {
	s.Id = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup) SetLeader(v *QueryAllDepartmentResponseBodyContentGroupAndExtListGroupLeader) *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup {
	s.Leader = v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup) SetName(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup {
	s.Name = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup) SetParentDeptCode(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup {
	s.ParentDeptCode = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup) SetRemark(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListGroup {
	s.Remark = &v
	return s
}

type QueryAllDepartmentResponseBodyContentGroupAndExtListGroupLeader struct {
	JobNumber *string `json:"jobNumber,omitempty" xml:"jobNumber,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryAllDepartmentResponseBodyContentGroupAndExtListGroupLeader) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentResponseBodyContentGroupAndExtListGroupLeader) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListGroupLeader) SetJobNumber(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListGroupLeader {
	s.JobNumber = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListGroupLeader) SetName(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListGroupLeader {
	s.Name = &v
	return s
}

func (s *QueryAllDepartmentResponseBodyContentGroupAndExtListGroupLeader) SetUserId(v string) *QueryAllDepartmentResponseBodyContentGroupAndExtListGroupLeader {
	s.UserId = &v
	return s
}

type QueryAllDepartmentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryAllDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDepartmentResponse) GoString() string {
	return s.String()
}

func (s *QueryAllDepartmentResponse) SetHeaders(v map[string]*string) *QueryAllDepartmentResponse {
	s.Headers = v
	return s
}

func (s *QueryAllDepartmentResponse) SetStatusCode(v int32) *QueryAllDepartmentResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAllDepartmentResponse) SetBody(v *QueryAllDepartmentResponseBody) *QueryAllDepartmentResponse {
	s.Body = v
	return s
}

type QueryAllDoctorsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllDoctorsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDoctorsHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllDoctorsHeaders) SetCommonHeaders(v map[string]*string) *QueryAllDoctorsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllDoctorsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllDoctorsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllDoctorsRequest struct {
	MonthMark *string `json:"monthMark,omitempty" xml:"monthMark,omitempty"`
	PageNum   *int32  `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize  *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryAllDoctorsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDoctorsRequest) GoString() string {
	return s.String()
}

func (s *QueryAllDoctorsRequest) SetMonthMark(v string) *QueryAllDoctorsRequest {
	s.MonthMark = &v
	return s
}

func (s *QueryAllDoctorsRequest) SetPageNum(v int32) *QueryAllDoctorsRequest {
	s.PageNum = &v
	return s
}

func (s *QueryAllDoctorsRequest) SetPageSize(v int32) *QueryAllDoctorsRequest {
	s.PageSize = &v
	return s
}

type QueryAllDoctorsResponseBody struct {
	Content     []*QueryAllDoctorsResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	CurrentPage *int32                                `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	TotalCount  *int64                                `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	TotalPages  *int32                                `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
}

func (s QueryAllDoctorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDoctorsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllDoctorsResponseBody) SetContent(v []*QueryAllDoctorsResponseBodyContent) *QueryAllDoctorsResponseBody {
	s.Content = v
	return s
}

func (s *QueryAllDoctorsResponseBody) SetCurrentPage(v int32) *QueryAllDoctorsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryAllDoctorsResponseBody) SetTotalCount(v int64) *QueryAllDoctorsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryAllDoctorsResponseBody) SetTotalPages(v int32) *QueryAllDoctorsResponseBody {
	s.TotalPages = &v
	return s
}

type QueryAllDoctorsResponseBodyContent struct {
	AssessGroupId   *string `json:"assessGroupId,omitempty" xml:"assessGroupId,omitempty"`
	AssessGroupName *string `json:"assessGroupName,omitempty" xml:"assessGroupName,omitempty"`
	DeptCode        *string `json:"deptCode,omitempty" xml:"deptCode,omitempty"`
	DeptType        *string `json:"deptType,omitempty" xml:"deptType,omitempty"`
	GmtCreateStr    *string `json:"gmtCreateStr,omitempty" xml:"gmtCreateStr,omitempty"`
	GmtModifiedStr  *string `json:"gmtModifiedStr,omitempty" xml:"gmtModifiedStr,omitempty"`
	Id              *int64  `json:"id,omitempty" xml:"id,omitempty"`
	JobNum          *string `json:"jobNum,omitempty" xml:"jobNum,omitempty"`
	Status          *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Uid             *string `json:"uid,omitempty" xml:"uid,omitempty"`
	UserCode        *string `json:"userCode,omitempty" xml:"userCode,omitempty"`
	UserName        *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s QueryAllDoctorsResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDoctorsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryAllDoctorsResponseBodyContent) SetAssessGroupId(v string) *QueryAllDoctorsResponseBodyContent {
	s.AssessGroupId = &v
	return s
}

func (s *QueryAllDoctorsResponseBodyContent) SetAssessGroupName(v string) *QueryAllDoctorsResponseBodyContent {
	s.AssessGroupName = &v
	return s
}

func (s *QueryAllDoctorsResponseBodyContent) SetDeptCode(v string) *QueryAllDoctorsResponseBodyContent {
	s.DeptCode = &v
	return s
}

func (s *QueryAllDoctorsResponseBodyContent) SetDeptType(v string) *QueryAllDoctorsResponseBodyContent {
	s.DeptType = &v
	return s
}

func (s *QueryAllDoctorsResponseBodyContent) SetGmtCreateStr(v string) *QueryAllDoctorsResponseBodyContent {
	s.GmtCreateStr = &v
	return s
}

func (s *QueryAllDoctorsResponseBodyContent) SetGmtModifiedStr(v string) *QueryAllDoctorsResponseBodyContent {
	s.GmtModifiedStr = &v
	return s
}

func (s *QueryAllDoctorsResponseBodyContent) SetId(v int64) *QueryAllDoctorsResponseBodyContent {
	s.Id = &v
	return s
}

func (s *QueryAllDoctorsResponseBodyContent) SetJobNum(v string) *QueryAllDoctorsResponseBodyContent {
	s.JobNum = &v
	return s
}

func (s *QueryAllDoctorsResponseBodyContent) SetStatus(v int32) *QueryAllDoctorsResponseBodyContent {
	s.Status = &v
	return s
}

func (s *QueryAllDoctorsResponseBodyContent) SetUid(v string) *QueryAllDoctorsResponseBodyContent {
	s.Uid = &v
	return s
}

func (s *QueryAllDoctorsResponseBodyContent) SetUserCode(v string) *QueryAllDoctorsResponseBodyContent {
	s.UserCode = &v
	return s
}

func (s *QueryAllDoctorsResponseBodyContent) SetUserName(v string) *QueryAllDoctorsResponseBodyContent {
	s.UserName = &v
	return s
}

type QueryAllDoctorsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryAllDoctorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllDoctorsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllDoctorsResponse) GoString() string {
	return s.String()
}

func (s *QueryAllDoctorsResponse) SetHeaders(v map[string]*string) *QueryAllDoctorsResponse {
	s.Headers = v
	return s
}

func (s *QueryAllDoctorsResponse) SetStatusCode(v int32) *QueryAllDoctorsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAllDoctorsResponse) SetBody(v *QueryAllDoctorsResponseBody) *QueryAllDoctorsResponse {
	s.Body = v
	return s
}

type QueryAllGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllGroupHeaders) SetCommonHeaders(v map[string]*string) *QueryAllGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllGroupHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllGroupRequest struct {
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryAllGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupRequest) GoString() string {
	return s.String()
}

func (s *QueryAllGroupRequest) SetPageNumber(v int32) *QueryAllGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryAllGroupRequest) SetPageSize(v int32) *QueryAllGroupRequest {
	s.PageSize = &v
	return s
}

type QueryAllGroupResponseBody struct {
	Content     []*QueryAllGroupResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	CurrentPage *int64                              `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	TotalCount  *int64                              `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	TotalPages  *int64                              `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
}

func (s QueryAllGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllGroupResponseBody) SetContent(v []*QueryAllGroupResponseBodyContent) *QueryAllGroupResponseBody {
	s.Content = v
	return s
}

func (s *QueryAllGroupResponseBody) SetCurrentPage(v int64) *QueryAllGroupResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryAllGroupResponseBody) SetTotalCount(v int64) *QueryAllGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryAllGroupResponseBody) SetTotalPages(v int64) *QueryAllGroupResponseBody {
	s.TotalPages = &v
	return s
}

type QueryAllGroupResponseBodyContent struct {
	DeptId *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	Id     *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryAllGroupResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryAllGroupResponseBodyContent) SetDeptId(v int64) *QueryAllGroupResponseBodyContent {
	s.DeptId = &v
	return s
}

func (s *QueryAllGroupResponseBodyContent) SetId(v int64) *QueryAllGroupResponseBodyContent {
	s.Id = &v
	return s
}

func (s *QueryAllGroupResponseBodyContent) SetName(v string) *QueryAllGroupResponseBodyContent {
	s.Name = &v
	return s
}

type QueryAllGroupResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryAllGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupResponse) GoString() string {
	return s.String()
}

func (s *QueryAllGroupResponse) SetHeaders(v map[string]*string) *QueryAllGroupResponse {
	s.Headers = v
	return s
}

func (s *QueryAllGroupResponse) SetStatusCode(v int32) *QueryAllGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAllGroupResponse) SetBody(v *QueryAllGroupResponseBody) *QueryAllGroupResponse {
	s.Body = v
	return s
}

type QueryAllGroupsInDeptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllGroupsInDeptHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupsInDeptHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllGroupsInDeptHeaders) SetCommonHeaders(v map[string]*string) *QueryAllGroupsInDeptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllGroupsInDeptHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllGroupsInDeptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllGroupsInDeptRequest struct {
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryAllGroupsInDeptRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupsInDeptRequest) GoString() string {
	return s.String()
}

func (s *QueryAllGroupsInDeptRequest) SetPageNumber(v int32) *QueryAllGroupsInDeptRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryAllGroupsInDeptRequest) SetPageSize(v int32) *QueryAllGroupsInDeptRequest {
	s.PageSize = &v
	return s
}

type QueryAllGroupsInDeptResponseBody struct {
	Content     []*QueryAllGroupsInDeptResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	CurrentPage *int64                                     `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	TotalCount  *int64                                     `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	TotalPages  *int64                                     `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
}

func (s QueryAllGroupsInDeptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupsInDeptResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllGroupsInDeptResponseBody) SetContent(v []*QueryAllGroupsInDeptResponseBodyContent) *QueryAllGroupsInDeptResponseBody {
	s.Content = v
	return s
}

func (s *QueryAllGroupsInDeptResponseBody) SetCurrentPage(v int64) *QueryAllGroupsInDeptResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryAllGroupsInDeptResponseBody) SetTotalCount(v int64) *QueryAllGroupsInDeptResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryAllGroupsInDeptResponseBody) SetTotalPages(v int64) *QueryAllGroupsInDeptResponseBody {
	s.TotalPages = &v
	return s
}

type QueryAllGroupsInDeptResponseBodyContent struct {
	DeptId *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	Id     *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryAllGroupsInDeptResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupsInDeptResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryAllGroupsInDeptResponseBodyContent) SetDeptId(v int64) *QueryAllGroupsInDeptResponseBodyContent {
	s.DeptId = &v
	return s
}

func (s *QueryAllGroupsInDeptResponseBodyContent) SetId(v int64) *QueryAllGroupsInDeptResponseBodyContent {
	s.Id = &v
	return s
}

func (s *QueryAllGroupsInDeptResponseBodyContent) SetName(v string) *QueryAllGroupsInDeptResponseBodyContent {
	s.Name = &v
	return s
}

type QueryAllGroupsInDeptResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryAllGroupsInDeptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllGroupsInDeptResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllGroupsInDeptResponse) GoString() string {
	return s.String()
}

func (s *QueryAllGroupsInDeptResponse) SetHeaders(v map[string]*string) *QueryAllGroupsInDeptResponse {
	s.Headers = v
	return s
}

func (s *QueryAllGroupsInDeptResponse) SetStatusCode(v int32) *QueryAllGroupsInDeptResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAllGroupsInDeptResponse) SetBody(v *QueryAllGroupsInDeptResponseBody) *QueryAllGroupsInDeptResponse {
	s.Body = v
	return s
}

type QueryAllMemberByDeptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllMemberByDeptHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByDeptHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByDeptHeaders) SetCommonHeaders(v map[string]*string) *QueryAllMemberByDeptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllMemberByDeptHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllMemberByDeptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllMemberByDeptRequest struct {
	MonthMark  *string `json:"monthMark,omitempty" xml:"monthMark,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryAllMemberByDeptRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByDeptRequest) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByDeptRequest) SetMonthMark(v string) *QueryAllMemberByDeptRequest {
	s.MonthMark = &v
	return s
}

func (s *QueryAllMemberByDeptRequest) SetPageNumber(v int32) *QueryAllMemberByDeptRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryAllMemberByDeptRequest) SetPageSize(v int32) *QueryAllMemberByDeptRequest {
	s.PageSize = &v
	return s
}

type QueryAllMemberByDeptResponseBody struct {
	Content     []*QueryAllMemberByDeptResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	CurrentPage *int32                                     `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	TotalCount  *int64                                     `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	TotalPages  *int32                                     `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
}

func (s QueryAllMemberByDeptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByDeptResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByDeptResponseBody) SetContent(v []*QueryAllMemberByDeptResponseBodyContent) *QueryAllMemberByDeptResponseBody {
	s.Content = v
	return s
}

func (s *QueryAllMemberByDeptResponseBody) SetCurrentPage(v int32) *QueryAllMemberByDeptResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryAllMemberByDeptResponseBody) SetTotalCount(v int64) *QueryAllMemberByDeptResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryAllMemberByDeptResponseBody) SetTotalPages(v int32) *QueryAllMemberByDeptResponseBody {
	s.TotalPages = &v
	return s
}

type QueryAllMemberByDeptResponseBodyContent struct {
	JobNum   *string `json:"jobNum,omitempty" xml:"jobNum,omitempty"`
	Uid      *string `json:"uid,omitempty" xml:"uid,omitempty"`
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s QueryAllMemberByDeptResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByDeptResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByDeptResponseBodyContent) SetJobNum(v string) *QueryAllMemberByDeptResponseBodyContent {
	s.JobNum = &v
	return s
}

func (s *QueryAllMemberByDeptResponseBodyContent) SetUid(v string) *QueryAllMemberByDeptResponseBodyContent {
	s.Uid = &v
	return s
}

func (s *QueryAllMemberByDeptResponseBodyContent) SetUserName(v string) *QueryAllMemberByDeptResponseBodyContent {
	s.UserName = &v
	return s
}

type QueryAllMemberByDeptResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryAllMemberByDeptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllMemberByDeptResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByDeptResponse) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByDeptResponse) SetHeaders(v map[string]*string) *QueryAllMemberByDeptResponse {
	s.Headers = v
	return s
}

func (s *QueryAllMemberByDeptResponse) SetStatusCode(v int32) *QueryAllMemberByDeptResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAllMemberByDeptResponse) SetBody(v *QueryAllMemberByDeptResponseBody) *QueryAllMemberByDeptResponse {
	s.Body = v
	return s
}

type QueryAllMemberByGroupHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllMemberByGroupHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByGroupHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByGroupHeaders) SetCommonHeaders(v map[string]*string) *QueryAllMemberByGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllMemberByGroupHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllMemberByGroupHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllMemberByGroupRequest struct {
	MonthMark  *string `json:"monthMark,omitempty" xml:"monthMark,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryAllMemberByGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByGroupRequest) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByGroupRequest) SetMonthMark(v string) *QueryAllMemberByGroupRequest {
	s.MonthMark = &v
	return s
}

func (s *QueryAllMemberByGroupRequest) SetPageNumber(v int32) *QueryAllMemberByGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryAllMemberByGroupRequest) SetPageSize(v int32) *QueryAllMemberByGroupRequest {
	s.PageSize = &v
	return s
}

type QueryAllMemberByGroupResponseBody struct {
	Content     []*QueryAllMemberByGroupResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	CurrentPage *int32                                      `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	TotalCount  *int64                                      `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	TotalPages  *int32                                      `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
}

func (s QueryAllMemberByGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByGroupResponseBody) SetContent(v []*QueryAllMemberByGroupResponseBodyContent) *QueryAllMemberByGroupResponseBody {
	s.Content = v
	return s
}

func (s *QueryAllMemberByGroupResponseBody) SetCurrentPage(v int32) *QueryAllMemberByGroupResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryAllMemberByGroupResponseBody) SetTotalCount(v int64) *QueryAllMemberByGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryAllMemberByGroupResponseBody) SetTotalPages(v int32) *QueryAllMemberByGroupResponseBody {
	s.TotalPages = &v
	return s
}

type QueryAllMemberByGroupResponseBodyContent struct {
	JobNum   *string `json:"jobNum,omitempty" xml:"jobNum,omitempty"`
	Uid      *string `json:"uid,omitempty" xml:"uid,omitempty"`
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s QueryAllMemberByGroupResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByGroupResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByGroupResponseBodyContent) SetJobNum(v string) *QueryAllMemberByGroupResponseBodyContent {
	s.JobNum = &v
	return s
}

func (s *QueryAllMemberByGroupResponseBodyContent) SetUid(v string) *QueryAllMemberByGroupResponseBodyContent {
	s.Uid = &v
	return s
}

func (s *QueryAllMemberByGroupResponseBodyContent) SetUserName(v string) *QueryAllMemberByGroupResponseBodyContent {
	s.UserName = &v
	return s
}

type QueryAllMemberByGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryAllMemberByGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllMemberByGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllMemberByGroupResponse) GoString() string {
	return s.String()
}

func (s *QueryAllMemberByGroupResponse) SetHeaders(v map[string]*string) *QueryAllMemberByGroupResponse {
	s.Headers = v
	return s
}

func (s *QueryAllMemberByGroupResponse) SetStatusCode(v int32) *QueryAllMemberByGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAllMemberByGroupResponse) SetBody(v *QueryAllMemberByGroupResponseBody) *QueryAllMemberByGroupResponse {
	s.Body = v
	return s
}

type QueryBizOptLogHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryBizOptLogHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryBizOptLogHeaders) GoString() string {
	return s.String()
}

func (s *QueryBizOptLogHeaders) SetCommonHeaders(v map[string]*string) *QueryBizOptLogHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryBizOptLogHeaders) SetXAcsDingtalkAccessToken(v string) *QueryBizOptLogHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryBizOptLogRequest struct {
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryBizOptLogRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBizOptLogRequest) GoString() string {
	return s.String()
}

func (s *QueryBizOptLogRequest) SetMaxResults(v int32) *QueryBizOptLogRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryBizOptLogRequest) SetNextToken(v int64) *QueryBizOptLogRequest {
	s.NextToken = &v
	return s
}

type QueryBizOptLogResponseBody struct {
	Content   []*QueryBizOptLogResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	NextToken *int64                               `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryBizOptLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBizOptLogResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBizOptLogResponseBody) SetContent(v []*QueryBizOptLogResponseBodyContent) *QueryBizOptLogResponseBody {
	s.Content = v
	return s
}

func (s *QueryBizOptLogResponseBody) SetNextToken(v int64) *QueryBizOptLogResponseBody {
	s.NextToken = &v
	return s
}

type QueryBizOptLogResponseBodyContent struct {
	BizType            *int32  `json:"bizType,omitempty" xml:"bizType,omitempty"`
	DataType           *int32  `json:"dataType,omitempty" xml:"dataType,omitempty"`
	Id                 *int64  `json:"id,omitempty" xml:"id,omitempty"`
	OptAfterData       *string `json:"optAfterData,omitempty" xml:"optAfterData,omitempty"`
	OptBeforeData      *string `json:"optBeforeData,omitempty" xml:"optBeforeData,omitempty"`
	OptBizType         *int32  `json:"optBizType,omitempty" xml:"optBizType,omitempty"`
	OptExtend          *string `json:"optExtend,omitempty" xml:"optExtend,omitempty"`
	OptJobNumber       *string `json:"optJobNumber,omitempty" xml:"optJobNumber,omitempty"`
	OptObjectCode      *string `json:"optObjectCode,omitempty" xml:"optObjectCode,omitempty"`
	OptObjectName      *string `json:"optObjectName,omitempty" xml:"optObjectName,omitempty"`
	OptObjectUserJobNo *string `json:"optObjectUserJobNo,omitempty" xml:"optObjectUserJobNo,omitempty"`
	OptSuccess         *int32  `json:"optSuccess,omitempty" xml:"optSuccess,omitempty"`
	OptTime            *int64  `json:"optTime,omitempty" xml:"optTime,omitempty"`
	OptType            *int32  `json:"optType,omitempty" xml:"optType,omitempty"`
	OptUserCode        *string `json:"optUserCode,omitempty" xml:"optUserCode,omitempty"`
	OptUserName        *string `json:"optUserName,omitempty" xml:"optUserName,omitempty"`
	Remark             *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s QueryBizOptLogResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryBizOptLogResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryBizOptLogResponseBodyContent) SetBizType(v int32) *QueryBizOptLogResponseBodyContent {
	s.BizType = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetDataType(v int32) *QueryBizOptLogResponseBodyContent {
	s.DataType = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetId(v int64) *QueryBizOptLogResponseBodyContent {
	s.Id = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptAfterData(v string) *QueryBizOptLogResponseBodyContent {
	s.OptAfterData = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptBeforeData(v string) *QueryBizOptLogResponseBodyContent {
	s.OptBeforeData = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptBizType(v int32) *QueryBizOptLogResponseBodyContent {
	s.OptBizType = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptExtend(v string) *QueryBizOptLogResponseBodyContent {
	s.OptExtend = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptJobNumber(v string) *QueryBizOptLogResponseBodyContent {
	s.OptJobNumber = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptObjectCode(v string) *QueryBizOptLogResponseBodyContent {
	s.OptObjectCode = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptObjectName(v string) *QueryBizOptLogResponseBodyContent {
	s.OptObjectName = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptObjectUserJobNo(v string) *QueryBizOptLogResponseBodyContent {
	s.OptObjectUserJobNo = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptSuccess(v int32) *QueryBizOptLogResponseBodyContent {
	s.OptSuccess = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptTime(v int64) *QueryBizOptLogResponseBodyContent {
	s.OptTime = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptType(v int32) *QueryBizOptLogResponseBodyContent {
	s.OptType = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptUserCode(v string) *QueryBizOptLogResponseBodyContent {
	s.OptUserCode = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetOptUserName(v string) *QueryBizOptLogResponseBodyContent {
	s.OptUserName = &v
	return s
}

func (s *QueryBizOptLogResponseBodyContent) SetRemark(v string) *QueryBizOptLogResponseBodyContent {
	s.Remark = &v
	return s
}

type QueryBizOptLogResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryBizOptLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBizOptLogResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBizOptLogResponse) GoString() string {
	return s.String()
}

func (s *QueryBizOptLogResponse) SetHeaders(v map[string]*string) *QueryBizOptLogResponse {
	s.Headers = v
	return s
}

func (s *QueryBizOptLogResponse) SetStatusCode(v int32) *QueryBizOptLogResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBizOptLogResponse) SetBody(v *QueryBizOptLogResponseBody) *QueryBizOptLogResponse {
	s.Body = v
	return s
}

type QueryDepartmentExtendInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDepartmentExtendInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentExtendInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryDepartmentExtendInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryDepartmentExtendInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDepartmentExtendInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDepartmentExtendInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDepartmentExtendInfoRequest struct {
	DeptCode *int64  `json:"deptCode,omitempty" xml:"deptCode,omitempty"`
	PropCode *string `json:"propCode,omitempty" xml:"propCode,omitempty"`
}

func (s QueryDepartmentExtendInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentExtendInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryDepartmentExtendInfoRequest) SetDeptCode(v int64) *QueryDepartmentExtendInfoRequest {
	s.DeptCode = &v
	return s
}

func (s *QueryDepartmentExtendInfoRequest) SetPropCode(v string) *QueryDepartmentExtendInfoRequest {
	s.PropCode = &v
	return s
}

type QueryDepartmentExtendInfoResponseBody struct {
	Content []*QueryDepartmentExtendInfoResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s QueryDepartmentExtendInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentExtendInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDepartmentExtendInfoResponseBody) SetContent(v []*QueryDepartmentExtendInfoResponseBodyContent) *QueryDepartmentExtendInfoResponseBody {
	s.Content = v
	return s
}

type QueryDepartmentExtendInfoResponseBodyContent struct {
	DeptCode              *string `json:"deptCode,omitempty" xml:"deptCode,omitempty"`
	DeptExtendDisplayName *string `json:"deptExtendDisplayName,omitempty" xml:"deptExtendDisplayName,omitempty"`
	DeptExtendKey         *string `json:"deptExtendKey,omitempty" xml:"deptExtendKey,omitempty"`
	DeptExtendValue       *string `json:"deptExtendValue,omitempty" xml:"deptExtendValue,omitempty"`
	GmtCreateStr          *string `json:"gmtCreateStr,omitempty" xml:"gmtCreateStr,omitempty"`
	GmtModifiedStr        *string `json:"gmtModifiedStr,omitempty" xml:"gmtModifiedStr,omitempty"`
	Id                    *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Status                *int32  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryDepartmentExtendInfoResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentExtendInfoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryDepartmentExtendInfoResponseBodyContent) SetDeptCode(v string) *QueryDepartmentExtendInfoResponseBodyContent {
	s.DeptCode = &v
	return s
}

func (s *QueryDepartmentExtendInfoResponseBodyContent) SetDeptExtendDisplayName(v string) *QueryDepartmentExtendInfoResponseBodyContent {
	s.DeptExtendDisplayName = &v
	return s
}

func (s *QueryDepartmentExtendInfoResponseBodyContent) SetDeptExtendKey(v string) *QueryDepartmentExtendInfoResponseBodyContent {
	s.DeptExtendKey = &v
	return s
}

func (s *QueryDepartmentExtendInfoResponseBodyContent) SetDeptExtendValue(v string) *QueryDepartmentExtendInfoResponseBodyContent {
	s.DeptExtendValue = &v
	return s
}

func (s *QueryDepartmentExtendInfoResponseBodyContent) SetGmtCreateStr(v string) *QueryDepartmentExtendInfoResponseBodyContent {
	s.GmtCreateStr = &v
	return s
}

func (s *QueryDepartmentExtendInfoResponseBodyContent) SetGmtModifiedStr(v string) *QueryDepartmentExtendInfoResponseBodyContent {
	s.GmtModifiedStr = &v
	return s
}

func (s *QueryDepartmentExtendInfoResponseBodyContent) SetId(v int64) *QueryDepartmentExtendInfoResponseBodyContent {
	s.Id = &v
	return s
}

func (s *QueryDepartmentExtendInfoResponseBodyContent) SetStatus(v int32) *QueryDepartmentExtendInfoResponseBodyContent {
	s.Status = &v
	return s
}

type QueryDepartmentExtendInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDepartmentExtendInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDepartmentExtendInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentExtendInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryDepartmentExtendInfoResponse) SetHeaders(v map[string]*string) *QueryDepartmentExtendInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryDepartmentExtendInfoResponse) SetStatusCode(v int32) *QueryDepartmentExtendInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDepartmentExtendInfoResponse) SetBody(v *QueryDepartmentExtendInfoResponseBody) *QueryDepartmentExtendInfoResponse {
	s.Body = v
	return s
}

type QueryDepartmentInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDepartmentInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryDepartmentInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryDepartmentInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDepartmentInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDepartmentInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDepartmentInfoResponseBody struct {
	Content *QueryDepartmentInfoResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
}

func (s QueryDepartmentInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDepartmentInfoResponseBody) SetContent(v *QueryDepartmentInfoResponseBodyContent) *QueryDepartmentInfoResponseBody {
	s.Content = v
	return s
}

type QueryDepartmentInfoResponseBodyContent struct {
	Department  *QueryDepartmentInfoResponseBodyContentDepartment    `json:"department,omitempty" xml:"department,omitempty" type:"Struct"`
	ExtendInfos []*QueryDepartmentInfoResponseBodyContentExtendInfos `json:"extendInfos,omitempty" xml:"extendInfos,omitempty" type:"Repeated"`
}

func (s QueryDepartmentInfoResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentInfoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryDepartmentInfoResponseBodyContent) SetDepartment(v *QueryDepartmentInfoResponseBodyContentDepartment) *QueryDepartmentInfoResponseBodyContent {
	s.Department = v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContent) SetExtendInfos(v []*QueryDepartmentInfoResponseBodyContentExtendInfos) *QueryDepartmentInfoResponseBodyContent {
	s.ExtendInfos = v
	return s
}

type QueryDepartmentInfoResponseBodyContentDepartment struct {
	DeptCode       *string  `json:"deptCode,omitempty" xml:"deptCode,omitempty"`
	DeptName       *string  `json:"deptName,omitempty" xml:"deptName,omitempty"`
	DeptOrder      *int64   `json:"deptOrder,omitempty" xml:"deptOrder,omitempty"`
	DeptStatus     *int32   `json:"deptStatus,omitempty" xml:"deptStatus,omitempty"`
	DeptType       *int32   `json:"deptType,omitempty" xml:"deptType,omitempty"`
	GmtCreateStr   *string  `json:"gmtCreateStr,omitempty" xml:"gmtCreateStr,omitempty"`
	GmtModifiedStr *string  `json:"gmtModifiedStr,omitempty" xml:"gmtModifiedStr,omitempty"`
	Id             *int64   `json:"id,omitempty" xml:"id,omitempty"`
	Name           *string  `json:"name,omitempty" xml:"name,omitempty"`
	ParentDeptCode *string  `json:"parentDeptCode,omitempty" xml:"parentDeptCode,omitempty"`
	Remark         *string  `json:"remark,omitempty" xml:"remark,omitempty"`
	WardIdList     []*int64 `json:"wardIdList,omitempty" xml:"wardIdList,omitempty" type:"Repeated"`
}

func (s QueryDepartmentInfoResponseBodyContentDepartment) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentInfoResponseBodyContentDepartment) GoString() string {
	return s.String()
}

func (s *QueryDepartmentInfoResponseBodyContentDepartment) SetDeptCode(v string) *QueryDepartmentInfoResponseBodyContentDepartment {
	s.DeptCode = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentDepartment) SetDeptName(v string) *QueryDepartmentInfoResponseBodyContentDepartment {
	s.DeptName = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentDepartment) SetDeptOrder(v int64) *QueryDepartmentInfoResponseBodyContentDepartment {
	s.DeptOrder = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentDepartment) SetDeptStatus(v int32) *QueryDepartmentInfoResponseBodyContentDepartment {
	s.DeptStatus = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentDepartment) SetDeptType(v int32) *QueryDepartmentInfoResponseBodyContentDepartment {
	s.DeptType = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentDepartment) SetGmtCreateStr(v string) *QueryDepartmentInfoResponseBodyContentDepartment {
	s.GmtCreateStr = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentDepartment) SetGmtModifiedStr(v string) *QueryDepartmentInfoResponseBodyContentDepartment {
	s.GmtModifiedStr = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentDepartment) SetId(v int64) *QueryDepartmentInfoResponseBodyContentDepartment {
	s.Id = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentDepartment) SetName(v string) *QueryDepartmentInfoResponseBodyContentDepartment {
	s.Name = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentDepartment) SetParentDeptCode(v string) *QueryDepartmentInfoResponseBodyContentDepartment {
	s.ParentDeptCode = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentDepartment) SetRemark(v string) *QueryDepartmentInfoResponseBodyContentDepartment {
	s.Remark = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentDepartment) SetWardIdList(v []*int64) *QueryDepartmentInfoResponseBodyContentDepartment {
	s.WardIdList = v
	return s
}

type QueryDepartmentInfoResponseBodyContentExtendInfos struct {
	DeptCode              *string `json:"deptCode,omitempty" xml:"deptCode,omitempty"`
	DeptExtendDisplayName *string `json:"deptExtendDisplayName,omitempty" xml:"deptExtendDisplayName,omitempty"`
	DeptExtendKey         *string `json:"deptExtendKey,omitempty" xml:"deptExtendKey,omitempty"`
	DeptExtendValue       *string `json:"deptExtendValue,omitempty" xml:"deptExtendValue,omitempty"`
	GmtCreateStr          *string `json:"gmtCreateStr,omitempty" xml:"gmtCreateStr,omitempty"`
	GmtModifiedStr        *string `json:"gmtModifiedStr,omitempty" xml:"gmtModifiedStr,omitempty"`
	Id                    *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Status                *int32  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryDepartmentInfoResponseBodyContentExtendInfos) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentInfoResponseBodyContentExtendInfos) GoString() string {
	return s.String()
}

func (s *QueryDepartmentInfoResponseBodyContentExtendInfos) SetDeptCode(v string) *QueryDepartmentInfoResponseBodyContentExtendInfos {
	s.DeptCode = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentExtendInfos) SetDeptExtendDisplayName(v string) *QueryDepartmentInfoResponseBodyContentExtendInfos {
	s.DeptExtendDisplayName = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentExtendInfos) SetDeptExtendKey(v string) *QueryDepartmentInfoResponseBodyContentExtendInfos {
	s.DeptExtendKey = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentExtendInfos) SetDeptExtendValue(v string) *QueryDepartmentInfoResponseBodyContentExtendInfos {
	s.DeptExtendValue = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentExtendInfos) SetGmtCreateStr(v string) *QueryDepartmentInfoResponseBodyContentExtendInfos {
	s.GmtCreateStr = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentExtendInfos) SetGmtModifiedStr(v string) *QueryDepartmentInfoResponseBodyContentExtendInfos {
	s.GmtModifiedStr = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentExtendInfos) SetId(v int64) *QueryDepartmentInfoResponseBodyContentExtendInfos {
	s.Id = &v
	return s
}

func (s *QueryDepartmentInfoResponseBodyContentExtendInfos) SetStatus(v int32) *QueryDepartmentInfoResponseBodyContentExtendInfos {
	s.Status = &v
	return s
}

type QueryDepartmentInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDepartmentInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDepartmentInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDepartmentInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryDepartmentInfoResponse) SetHeaders(v map[string]*string) *QueryDepartmentInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryDepartmentInfoResponse) SetStatusCode(v int32) *QueryDepartmentInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDepartmentInfoResponse) SetBody(v *QueryDepartmentInfoResponseBody) *QueryDepartmentInfoResponse {
	s.Body = v
	return s
}

type QueryDoctorDetailsByJobNumberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryDoctorDetailsByJobNumberHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryDoctorDetailsByJobNumberHeaders) GoString() string {
	return s.String()
}

func (s *QueryDoctorDetailsByJobNumberHeaders) SetCommonHeaders(v map[string]*string) *QueryDoctorDetailsByJobNumberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDoctorDetailsByJobNumberHeaders) SetXAcsDingtalkAccessToken(v string) *QueryDoctorDetailsByJobNumberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryDoctorDetailsByJobNumberRequest struct {
	MonthMark *string `json:"monthMark,omitempty" xml:"monthMark,omitempty"`
}

func (s QueryDoctorDetailsByJobNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDoctorDetailsByJobNumberRequest) GoString() string {
	return s.String()
}

func (s *QueryDoctorDetailsByJobNumberRequest) SetMonthMark(v string) *QueryDoctorDetailsByJobNumberRequest {
	s.MonthMark = &v
	return s
}

type QueryDoctorDetailsByJobNumberResponseBody struct {
	Content *QueryDoctorDetailsByJobNumberResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
}

func (s QueryDoctorDetailsByJobNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDoctorDetailsByJobNumberResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDoctorDetailsByJobNumberResponseBody) SetContent(v *QueryDoctorDetailsByJobNumberResponseBodyContent) *QueryDoctorDetailsByJobNumberResponseBody {
	s.Content = v
	return s
}

type QueryDoctorDetailsByJobNumberResponseBodyContent struct {
	DeptList          []*QueryDoctorDetailsByJobNumberResponseBodyContentDeptList        `json:"deptList,omitempty" xml:"deptList,omitempty" type:"Repeated"`
	GroupList         []*QueryDoctorDetailsByJobNumberResponseBodyContentGroupList       `json:"groupList,omitempty" xml:"groupList,omitempty" type:"Repeated"`
	JobNumber         *string                                                            `json:"jobNumber,omitempty" xml:"jobNumber,omitempty"`
	JobStatus         []*QueryDoctorDetailsByJobNumberResponseBodyContentJobStatus       `json:"jobStatus,omitempty" xml:"jobStatus,omitempty" type:"Repeated"`
	ProfessionalTitle *QueryDoctorDetailsByJobNumberResponseBodyContentProfessionalTitle `json:"professionalTitle,omitempty" xml:"professionalTitle,omitempty" type:"Struct"`
	UserId            *string                                                            `json:"userId,omitempty" xml:"userId,omitempty"`
	UserName          *string                                                            `json:"userName,omitempty" xml:"userName,omitempty"`
	UserProbList      []*QueryDoctorDetailsByJobNumberResponseBodyContentUserProbList    `json:"userProbList,omitempty" xml:"userProbList,omitempty" type:"Repeated"`
}

func (s QueryDoctorDetailsByJobNumberResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryDoctorDetailsByJobNumberResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContent) SetDeptList(v []*QueryDoctorDetailsByJobNumberResponseBodyContentDeptList) *QueryDoctorDetailsByJobNumberResponseBodyContent {
	s.DeptList = v
	return s
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContent) SetGroupList(v []*QueryDoctorDetailsByJobNumberResponseBodyContentGroupList) *QueryDoctorDetailsByJobNumberResponseBodyContent {
	s.GroupList = v
	return s
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContent) SetJobNumber(v string) *QueryDoctorDetailsByJobNumberResponseBodyContent {
	s.JobNumber = &v
	return s
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContent) SetJobStatus(v []*QueryDoctorDetailsByJobNumberResponseBodyContentJobStatus) *QueryDoctorDetailsByJobNumberResponseBodyContent {
	s.JobStatus = v
	return s
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContent) SetProfessionalTitle(v *QueryDoctorDetailsByJobNumberResponseBodyContentProfessionalTitle) *QueryDoctorDetailsByJobNumberResponseBodyContent {
	s.ProfessionalTitle = v
	return s
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContent) SetUserId(v string) *QueryDoctorDetailsByJobNumberResponseBodyContent {
	s.UserId = &v
	return s
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContent) SetUserName(v string) *QueryDoctorDetailsByJobNumberResponseBodyContent {
	s.UserName = &v
	return s
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContent) SetUserProbList(v []*QueryDoctorDetailsByJobNumberResponseBodyContentUserProbList) *QueryDoctorDetailsByJobNumberResponseBodyContent {
	s.UserProbList = v
	return s
}

type QueryDoctorDetailsByJobNumberResponseBodyContentDeptList struct {
	CategoryName *string `json:"categoryName,omitempty" xml:"categoryName,omitempty"`
	DeptId       *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	DeptName     *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	GmtCreate    *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified  *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	RelationId   *int64  `json:"relationId,omitempty" xml:"relationId,omitempty"`
}

func (s QueryDoctorDetailsByJobNumberResponseBodyContentDeptList) String() string {
	return tea.Prettify(s)
}

func (s QueryDoctorDetailsByJobNumberResponseBodyContentDeptList) GoString() string {
	return s.String()
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContentDeptList) SetCategoryName(v string) *QueryDoctorDetailsByJobNumberResponseBodyContentDeptList {
	s.CategoryName = &v
	return s
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContentDeptList) SetDeptId(v int64) *QueryDoctorDetailsByJobNumberResponseBodyContentDeptList {
	s.DeptId = &v
	return s
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContentDeptList) SetDeptName(v string) *QueryDoctorDetailsByJobNumberResponseBodyContentDeptList {
	s.DeptName = &v
	return s
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContentDeptList) SetGmtCreate(v string) *QueryDoctorDetailsByJobNumberResponseBodyContentDeptList {
	s.GmtCreate = &v
	return s
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContentDeptList) SetGmtModified(v string) *QueryDoctorDetailsByJobNumberResponseBodyContentDeptList {
	s.GmtModified = &v
	return s
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContentDeptList) SetRelationId(v int64) *QueryDoctorDetailsByJobNumberResponseBodyContentDeptList {
	s.RelationId = &v
	return s
}

type QueryDoctorDetailsByJobNumberResponseBodyContentGroupList struct {
	DeptId        *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	DeptName      *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	GroupId       *int64  `json:"groupId,omitempty" xml:"groupId,omitempty"`
	GroupName     *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	IsAssessGroup *string `json:"isAssessGroup,omitempty" xml:"isAssessGroup,omitempty"`
	IsLeader      *bool   `json:"isLeader,omitempty" xml:"isLeader,omitempty"`
	RelationId    *int64  `json:"relationId,omitempty" xml:"relationId,omitempty"`
}

func (s QueryDoctorDetailsByJobNumberResponseBodyContentGroupList) String() string {
	return tea.Prettify(s)
}

func (s QueryDoctorDetailsByJobNumberResponseBodyContentGroupList) GoString() string {
	return s.String()
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContentGroupList) SetDeptId(v int64) *QueryDoctorDetailsByJobNumberResponseBodyContentGroupList {
	s.DeptId = &v
	return s
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContentGroupList) SetDeptName(v string) *QueryDoctorDetailsByJobNumberResponseBodyContentGroupList {
	s.DeptName = &v
	return s
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContentGroupList) SetGroupId(v int64) *QueryDoctorDetailsByJobNumberResponseBodyContentGroupList {
	s.GroupId = &v
	return s
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContentGroupList) SetGroupName(v string) *QueryDoctorDetailsByJobNumberResponseBodyContentGroupList {
	s.GroupName = &v
	return s
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContentGroupList) SetIsAssessGroup(v string) *QueryDoctorDetailsByJobNumberResponseBodyContentGroupList {
	s.IsAssessGroup = &v
	return s
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContentGroupList) SetIsLeader(v bool) *QueryDoctorDetailsByJobNumberResponseBodyContentGroupList {
	s.IsLeader = &v
	return s
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContentGroupList) SetRelationId(v int64) *QueryDoctorDetailsByJobNumberResponseBodyContentGroupList {
	s.RelationId = &v
	return s
}

type QueryDoctorDetailsByJobNumberResponseBodyContentJobStatus struct {
	Code       *string `json:"code,omitempty" xml:"code,omitempty"`
	StatusName *string `json:"statusName,omitempty" xml:"statusName,omitempty"`
}

func (s QueryDoctorDetailsByJobNumberResponseBodyContentJobStatus) String() string {
	return tea.Prettify(s)
}

func (s QueryDoctorDetailsByJobNumberResponseBodyContentJobStatus) GoString() string {
	return s.String()
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContentJobStatus) SetCode(v string) *QueryDoctorDetailsByJobNumberResponseBodyContentJobStatus {
	s.Code = &v
	return s
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContentJobStatus) SetStatusName(v string) *QueryDoctorDetailsByJobNumberResponseBodyContentJobStatus {
	s.StatusName = &v
	return s
}

type QueryDoctorDetailsByJobNumberResponseBodyContentProfessionalTitle struct {
	Code                      *string `json:"code,omitempty" xml:"code,omitempty"`
	ProfessionalTitleCategory *string `json:"professionalTitleCategory,omitempty" xml:"professionalTitleCategory,omitempty"`
	ProfessionalTitleDetail   *string `json:"professionalTitleDetail,omitempty" xml:"professionalTitleDetail,omitempty"`
}

func (s QueryDoctorDetailsByJobNumberResponseBodyContentProfessionalTitle) String() string {
	return tea.Prettify(s)
}

func (s QueryDoctorDetailsByJobNumberResponseBodyContentProfessionalTitle) GoString() string {
	return s.String()
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContentProfessionalTitle) SetCode(v string) *QueryDoctorDetailsByJobNumberResponseBodyContentProfessionalTitle {
	s.Code = &v
	return s
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContentProfessionalTitle) SetProfessionalTitleCategory(v string) *QueryDoctorDetailsByJobNumberResponseBodyContentProfessionalTitle {
	s.ProfessionalTitleCategory = &v
	return s
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContentProfessionalTitle) SetProfessionalTitleDetail(v string) *QueryDoctorDetailsByJobNumberResponseBodyContentProfessionalTitle {
	s.ProfessionalTitleDetail = &v
	return s
}

type QueryDoctorDetailsByJobNumberResponseBodyContentUserProbList struct {
	Code             *string `json:"code,omitempty" xml:"code,omitempty"`
	UserPropertyName *string `json:"userPropertyName,omitempty" xml:"userPropertyName,omitempty"`
}

func (s QueryDoctorDetailsByJobNumberResponseBodyContentUserProbList) String() string {
	return tea.Prettify(s)
}

func (s QueryDoctorDetailsByJobNumberResponseBodyContentUserProbList) GoString() string {
	return s.String()
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContentUserProbList) SetCode(v string) *QueryDoctorDetailsByJobNumberResponseBodyContentUserProbList {
	s.Code = &v
	return s
}

func (s *QueryDoctorDetailsByJobNumberResponseBodyContentUserProbList) SetUserPropertyName(v string) *QueryDoctorDetailsByJobNumberResponseBodyContentUserProbList {
	s.UserPropertyName = &v
	return s
}

type QueryDoctorDetailsByJobNumberResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDoctorDetailsByJobNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDoctorDetailsByJobNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDoctorDetailsByJobNumberResponse) GoString() string {
	return s.String()
}

func (s *QueryDoctorDetailsByJobNumberResponse) SetHeaders(v map[string]*string) *QueryDoctorDetailsByJobNumberResponse {
	s.Headers = v
	return s
}

func (s *QueryDoctorDetailsByJobNumberResponse) SetStatusCode(v int32) *QueryDoctorDetailsByJobNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDoctorDetailsByJobNumberResponse) SetBody(v *QueryDoctorDetailsByJobNumberResponseBody) *QueryDoctorDetailsByJobNumberResponse {
	s.Body = v
	return s
}

type QueryGroupInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryGroupInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryGroupInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGroupInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryGroupInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryGroupInfoResponseBody struct {
	Content *QueryGroupInfoResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
}

func (s QueryGroupInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoResponseBody) SetContent(v *QueryGroupInfoResponseBodyContent) *QueryGroupInfoResponseBody {
	s.Content = v
	return s
}

type QueryGroupInfoResponseBodyContent struct {
	ExtendInfos []*QueryGroupInfoResponseBodyContentExtendInfos `json:"extendInfos,omitempty" xml:"extendInfos,omitempty" type:"Repeated"`
	Group       *QueryGroupInfoResponseBodyContentGroup         `json:"group,omitempty" xml:"group,omitempty" type:"Struct"`
}

func (s QueryGroupInfoResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoResponseBodyContent) SetExtendInfos(v []*QueryGroupInfoResponseBodyContentExtendInfos) *QueryGroupInfoResponseBodyContent {
	s.ExtendInfos = v
	return s
}

func (s *QueryGroupInfoResponseBodyContent) SetGroup(v *QueryGroupInfoResponseBodyContentGroup) *QueryGroupInfoResponseBodyContent {
	s.Group = v
	return s
}

type QueryGroupInfoResponseBodyContentExtendInfos struct {
	DeptCode              *string `json:"deptCode,omitempty" xml:"deptCode,omitempty"`
	DeptExtendDisplayName *string `json:"deptExtendDisplayName,omitempty" xml:"deptExtendDisplayName,omitempty"`
	DeptExtendKey         *string `json:"deptExtendKey,omitempty" xml:"deptExtendKey,omitempty"`
	DeptExtendValue       *string `json:"deptExtendValue,omitempty" xml:"deptExtendValue,omitempty"`
	GmtCreateStr          *string `json:"gmtCreateStr,omitempty" xml:"gmtCreateStr,omitempty"`
	GmtModifiedStr        *string `json:"gmtModifiedStr,omitempty" xml:"gmtModifiedStr,omitempty"`
	Id                    *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Status                *int32  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryGroupInfoResponseBodyContentExtendInfos) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoResponseBodyContentExtendInfos) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoResponseBodyContentExtendInfos) SetDeptCode(v string) *QueryGroupInfoResponseBodyContentExtendInfos {
	s.DeptCode = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentExtendInfos) SetDeptExtendDisplayName(v string) *QueryGroupInfoResponseBodyContentExtendInfos {
	s.DeptExtendDisplayName = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentExtendInfos) SetDeptExtendKey(v string) *QueryGroupInfoResponseBodyContentExtendInfos {
	s.DeptExtendKey = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentExtendInfos) SetDeptExtendValue(v string) *QueryGroupInfoResponseBodyContentExtendInfos {
	s.DeptExtendValue = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentExtendInfos) SetGmtCreateStr(v string) *QueryGroupInfoResponseBodyContentExtendInfos {
	s.GmtCreateStr = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentExtendInfos) SetGmtModifiedStr(v string) *QueryGroupInfoResponseBodyContentExtendInfos {
	s.GmtModifiedStr = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentExtendInfos) SetId(v int64) *QueryGroupInfoResponseBodyContentExtendInfos {
	s.Id = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentExtendInfos) SetStatus(v int32) *QueryGroupInfoResponseBodyContentExtendInfos {
	s.Status = &v
	return s
}

type QueryGroupInfoResponseBodyContentGroup struct {
	DeptId         *int64                                        `json:"deptId,omitempty" xml:"deptId,omitempty"`
	DeptStatus     *int32                                        `json:"deptStatus,omitempty" xml:"deptStatus,omitempty"`
	GmtCreateStr   *string                                       `json:"gmtCreateStr,omitempty" xml:"gmtCreateStr,omitempty"`
	GmtModifiedStr *string                                       `json:"gmtModifiedStr,omitempty" xml:"gmtModifiedStr,omitempty"`
	Id             *int64                                        `json:"id,omitempty" xml:"id,omitempty"`
	Leader         *QueryGroupInfoResponseBodyContentGroupLeader `json:"leader,omitempty" xml:"leader,omitempty" type:"Struct"`
	Name           *string                                       `json:"name,omitempty" xml:"name,omitempty"`
	ParentDeptCode *string                                       `json:"parentDeptCode,omitempty" xml:"parentDeptCode,omitempty"`
	Remark         *string                                       `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s QueryGroupInfoResponseBodyContentGroup) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoResponseBodyContentGroup) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoResponseBodyContentGroup) SetDeptId(v int64) *QueryGroupInfoResponseBodyContentGroup {
	s.DeptId = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentGroup) SetDeptStatus(v int32) *QueryGroupInfoResponseBodyContentGroup {
	s.DeptStatus = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentGroup) SetGmtCreateStr(v string) *QueryGroupInfoResponseBodyContentGroup {
	s.GmtCreateStr = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentGroup) SetGmtModifiedStr(v string) *QueryGroupInfoResponseBodyContentGroup {
	s.GmtModifiedStr = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentGroup) SetId(v int64) *QueryGroupInfoResponseBodyContentGroup {
	s.Id = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentGroup) SetLeader(v *QueryGroupInfoResponseBodyContentGroupLeader) *QueryGroupInfoResponseBodyContentGroup {
	s.Leader = v
	return s
}

func (s *QueryGroupInfoResponseBodyContentGroup) SetName(v string) *QueryGroupInfoResponseBodyContentGroup {
	s.Name = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentGroup) SetParentDeptCode(v string) *QueryGroupInfoResponseBodyContentGroup {
	s.ParentDeptCode = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentGroup) SetRemark(v string) *QueryGroupInfoResponseBodyContentGroup {
	s.Remark = &v
	return s
}

type QueryGroupInfoResponseBodyContentGroupLeader struct {
	JobNumber *string `json:"jobNumber,omitempty" xml:"jobNumber,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryGroupInfoResponseBodyContentGroupLeader) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoResponseBodyContentGroupLeader) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoResponseBodyContentGroupLeader) SetJobNumber(v string) *QueryGroupInfoResponseBodyContentGroupLeader {
	s.JobNumber = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentGroupLeader) SetName(v string) *QueryGroupInfoResponseBodyContentGroupLeader {
	s.Name = &v
	return s
}

func (s *QueryGroupInfoResponseBodyContentGroupLeader) SetUserId(v string) *QueryGroupInfoResponseBodyContentGroupLeader {
	s.UserId = &v
	return s
}

type QueryGroupInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryGroupInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryGroupInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGroupInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryGroupInfoResponse) SetHeaders(v map[string]*string) *QueryGroupInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryGroupInfoResponse) SetStatusCode(v int32) *QueryGroupInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGroupInfoResponse) SetBody(v *QueryGroupInfoResponseBody) *QueryGroupInfoResponse {
	s.Body = v
	return s
}

type QueryHospitalDistrictInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryHospitalDistrictInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalDistrictInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryHospitalDistrictInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryHospitalDistrictInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryHospitalDistrictInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryHospitalDistrictInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryHospitalDistrictInfoRequest struct {
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryHospitalDistrictInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalDistrictInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryHospitalDistrictInfoRequest) SetPageNumber(v int32) *QueryHospitalDistrictInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryHospitalDistrictInfoRequest) SetPageSize(v int32) *QueryHospitalDistrictInfoRequest {
	s.PageSize = &v
	return s
}

type QueryHospitalDistrictInfoResponseBody struct {
	Content     []*QueryHospitalDistrictInfoResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	CurrentPage *int64                                          `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	TotalCount  *int64                                          `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	TotalPages  *int32                                          `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
}

func (s QueryHospitalDistrictInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalDistrictInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHospitalDistrictInfoResponseBody) SetContent(v []*QueryHospitalDistrictInfoResponseBodyContent) *QueryHospitalDistrictInfoResponseBody {
	s.Content = v
	return s
}

func (s *QueryHospitalDistrictInfoResponseBody) SetCurrentPage(v int64) *QueryHospitalDistrictInfoResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryHospitalDistrictInfoResponseBody) SetTotalCount(v int64) *QueryHospitalDistrictInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryHospitalDistrictInfoResponseBody) SetTotalPages(v int32) *QueryHospitalDistrictInfoResponseBody {
	s.TotalPages = &v
	return s
}

type QueryHospitalDistrictInfoResponseBodyContent struct {
	Address          *string `json:"address,omitempty" xml:"address,omitempty"`
	Deleted          *int32  `json:"deleted,omitempty" xml:"deleted,omitempty"`
	DistrictName     *string `json:"districtName,omitempty" xml:"districtName,omitempty"`
	DistrictType     *int32  `json:"districtType,omitempty" xml:"districtType,omitempty"`
	GmtCreate        *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified      *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id               *int64  `json:"id,omitempty" xml:"id,omitempty"`
	ParentDistrictId *int64  `json:"parentDistrictId,omitempty" xml:"parentDistrictId,omitempty"`
}

func (s QueryHospitalDistrictInfoResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalDistrictInfoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryHospitalDistrictInfoResponseBodyContent) SetAddress(v string) *QueryHospitalDistrictInfoResponseBodyContent {
	s.Address = &v
	return s
}

func (s *QueryHospitalDistrictInfoResponseBodyContent) SetDeleted(v int32) *QueryHospitalDistrictInfoResponseBodyContent {
	s.Deleted = &v
	return s
}

func (s *QueryHospitalDistrictInfoResponseBodyContent) SetDistrictName(v string) *QueryHospitalDistrictInfoResponseBodyContent {
	s.DistrictName = &v
	return s
}

func (s *QueryHospitalDistrictInfoResponseBodyContent) SetDistrictType(v int32) *QueryHospitalDistrictInfoResponseBodyContent {
	s.DistrictType = &v
	return s
}

func (s *QueryHospitalDistrictInfoResponseBodyContent) SetGmtCreate(v string) *QueryHospitalDistrictInfoResponseBodyContent {
	s.GmtCreate = &v
	return s
}

func (s *QueryHospitalDistrictInfoResponseBodyContent) SetGmtModified(v string) *QueryHospitalDistrictInfoResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *QueryHospitalDistrictInfoResponseBodyContent) SetId(v int64) *QueryHospitalDistrictInfoResponseBodyContent {
	s.Id = &v
	return s
}

func (s *QueryHospitalDistrictInfoResponseBodyContent) SetParentDistrictId(v int64) *QueryHospitalDistrictInfoResponseBodyContent {
	s.ParentDistrictId = &v
	return s
}

type QueryHospitalDistrictInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryHospitalDistrictInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryHospitalDistrictInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalDistrictInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryHospitalDistrictInfoResponse) SetHeaders(v map[string]*string) *QueryHospitalDistrictInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryHospitalDistrictInfoResponse) SetStatusCode(v int32) *QueryHospitalDistrictInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHospitalDistrictInfoResponse) SetBody(v *QueryHospitalDistrictInfoResponseBody) *QueryHospitalDistrictInfoResponse {
	s.Body = v
	return s
}

type QueryHospitalRoleUserInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryHospitalRoleUserInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalRoleUserInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryHospitalRoleUserInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryHospitalRoleUserInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryHospitalRoleUserInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryHospitalRoleUserInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryHospitalRoleUserInfoRequest struct {
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s QueryHospitalRoleUserInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalRoleUserInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryHospitalRoleUserInfoRequest) SetPageNumber(v int64) *QueryHospitalRoleUserInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryHospitalRoleUserInfoRequest) SetPageSize(v int64) *QueryHospitalRoleUserInfoRequest {
	s.PageSize = &v
	return s
}

type QueryHospitalRoleUserInfoResponseBody struct {
	Content     []*QueryHospitalRoleUserInfoResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	CurrentPage *int32                                          `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	TotalCount  *int64                                          `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	TotalPages  *int32                                          `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
}

func (s QueryHospitalRoleUserInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalRoleUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHospitalRoleUserInfoResponseBody) SetContent(v []*QueryHospitalRoleUserInfoResponseBodyContent) *QueryHospitalRoleUserInfoResponseBody {
	s.Content = v
	return s
}

func (s *QueryHospitalRoleUserInfoResponseBody) SetCurrentPage(v int32) *QueryHospitalRoleUserInfoResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryHospitalRoleUserInfoResponseBody) SetTotalCount(v int64) *QueryHospitalRoleUserInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryHospitalRoleUserInfoResponseBody) SetTotalPages(v int32) *QueryHospitalRoleUserInfoResponseBody {
	s.TotalPages = &v
	return s
}

type QueryHospitalRoleUserInfoResponseBodyContent struct {
	GmtCreate   *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	JobNumber   *string `json:"jobNumber,omitempty" xml:"jobNumber,omitempty"`
	RoleCode    *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	RoleName    *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	Status      *int32  `json:"status,omitempty" xml:"status,omitempty"`
	UserCode    *string `json:"userCode,omitempty" xml:"userCode,omitempty"`
	UserName    *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s QueryHospitalRoleUserInfoResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalRoleUserInfoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryHospitalRoleUserInfoResponseBodyContent) SetGmtCreate(v string) *QueryHospitalRoleUserInfoResponseBodyContent {
	s.GmtCreate = &v
	return s
}

func (s *QueryHospitalRoleUserInfoResponseBodyContent) SetGmtModified(v string) *QueryHospitalRoleUserInfoResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *QueryHospitalRoleUserInfoResponseBodyContent) SetJobNumber(v string) *QueryHospitalRoleUserInfoResponseBodyContent {
	s.JobNumber = &v
	return s
}

func (s *QueryHospitalRoleUserInfoResponseBodyContent) SetRoleCode(v string) *QueryHospitalRoleUserInfoResponseBodyContent {
	s.RoleCode = &v
	return s
}

func (s *QueryHospitalRoleUserInfoResponseBodyContent) SetRoleName(v string) *QueryHospitalRoleUserInfoResponseBodyContent {
	s.RoleName = &v
	return s
}

func (s *QueryHospitalRoleUserInfoResponseBodyContent) SetStatus(v int32) *QueryHospitalRoleUserInfoResponseBodyContent {
	s.Status = &v
	return s
}

func (s *QueryHospitalRoleUserInfoResponseBodyContent) SetUserCode(v string) *QueryHospitalRoleUserInfoResponseBodyContent {
	s.UserCode = &v
	return s
}

func (s *QueryHospitalRoleUserInfoResponseBodyContent) SetUserName(v string) *QueryHospitalRoleUserInfoResponseBodyContent {
	s.UserName = &v
	return s
}

type QueryHospitalRoleUserInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryHospitalRoleUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryHospitalRoleUserInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalRoleUserInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryHospitalRoleUserInfoResponse) SetHeaders(v map[string]*string) *QueryHospitalRoleUserInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryHospitalRoleUserInfoResponse) SetStatusCode(v int32) *QueryHospitalRoleUserInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHospitalRoleUserInfoResponse) SetBody(v *QueryHospitalRoleUserInfoResponseBody) *QueryHospitalRoleUserInfoResponse {
	s.Body = v
	return s
}

type QueryHospitalRolesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryHospitalRolesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalRolesHeaders) GoString() string {
	return s.String()
}

func (s *QueryHospitalRolesHeaders) SetCommonHeaders(v map[string]*string) *QueryHospitalRolesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryHospitalRolesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryHospitalRolesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryHospitalRolesResponseBody struct {
	Content []*QueryHospitalRolesResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s QueryHospitalRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalRolesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHospitalRolesResponseBody) SetContent(v []*QueryHospitalRolesResponseBodyContent) *QueryHospitalRolesResponseBody {
	s.Content = v
	return s
}

type QueryHospitalRolesResponseBodyContent struct {
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	Id        *int64  `json:"id,omitempty" xml:"id,omitempty"`
	IsDeleted *int64  `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	ReadOnly  *int64  `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
	Remark    *string `json:"remark,omitempty" xml:"remark,omitempty"`
	RoleCode  *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	RoleName  *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	Sort      *int64  `json:"sort,omitempty" xml:"sort,omitempty"`
}

func (s QueryHospitalRolesResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalRolesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryHospitalRolesResponseBodyContent) SetGmtCreate(v string) *QueryHospitalRolesResponseBodyContent {
	s.GmtCreate = &v
	return s
}

func (s *QueryHospitalRolesResponseBodyContent) SetId(v int64) *QueryHospitalRolesResponseBodyContent {
	s.Id = &v
	return s
}

func (s *QueryHospitalRolesResponseBodyContent) SetIsDeleted(v int64) *QueryHospitalRolesResponseBodyContent {
	s.IsDeleted = &v
	return s
}

func (s *QueryHospitalRolesResponseBodyContent) SetReadOnly(v int64) *QueryHospitalRolesResponseBodyContent {
	s.ReadOnly = &v
	return s
}

func (s *QueryHospitalRolesResponseBodyContent) SetRemark(v string) *QueryHospitalRolesResponseBodyContent {
	s.Remark = &v
	return s
}

func (s *QueryHospitalRolesResponseBodyContent) SetRoleCode(v string) *QueryHospitalRolesResponseBodyContent {
	s.RoleCode = &v
	return s
}

func (s *QueryHospitalRolesResponseBodyContent) SetRoleName(v string) *QueryHospitalRolesResponseBodyContent {
	s.RoleName = &v
	return s
}

func (s *QueryHospitalRolesResponseBodyContent) SetSort(v int64) *QueryHospitalRolesResponseBodyContent {
	s.Sort = &v
	return s
}

type QueryHospitalRolesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryHospitalRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryHospitalRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHospitalRolesResponse) GoString() string {
	return s.String()
}

func (s *QueryHospitalRolesResponse) SetHeaders(v map[string]*string) *QueryHospitalRolesResponse {
	s.Headers = v
	return s
}

func (s *QueryHospitalRolesResponse) SetStatusCode(v int32) *QueryHospitalRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHospitalRolesResponse) SetBody(v *QueryHospitalRolesResponseBody) *QueryHospitalRolesResponse {
	s.Body = v
	return s
}

type QueryJobCodeDictionaryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryJobCodeDictionaryHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryJobCodeDictionaryHeaders) GoString() string {
	return s.String()
}

func (s *QueryJobCodeDictionaryHeaders) SetCommonHeaders(v map[string]*string) *QueryJobCodeDictionaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryJobCodeDictionaryHeaders) SetXAcsDingtalkAccessToken(v string) *QueryJobCodeDictionaryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryJobCodeDictionaryResponseBody struct {
	Content []*QueryJobCodeDictionaryResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s QueryJobCodeDictionaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryJobCodeDictionaryResponseBody) GoString() string {
	return s.String()
}

func (s *QueryJobCodeDictionaryResponseBody) SetContent(v []*QueryJobCodeDictionaryResponseBodyContent) *QueryJobCodeDictionaryResponseBody {
	s.Content = v
	return s
}

type QueryJobCodeDictionaryResponseBodyContent struct {
	Category    *string `json:"category,omitempty" xml:"category,omitempty"`
	Code        *string `json:"code,omitempty" xml:"code,omitempty"`
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	DoctorType  *string `json:"doctorType,omitempty" xml:"doctorType,omitempty"`
}

func (s QueryJobCodeDictionaryResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryJobCodeDictionaryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryJobCodeDictionaryResponseBodyContent) SetCategory(v string) *QueryJobCodeDictionaryResponseBodyContent {
	s.Category = &v
	return s
}

func (s *QueryJobCodeDictionaryResponseBodyContent) SetCode(v string) *QueryJobCodeDictionaryResponseBodyContent {
	s.Code = &v
	return s
}

func (s *QueryJobCodeDictionaryResponseBodyContent) SetDisplayName(v string) *QueryJobCodeDictionaryResponseBodyContent {
	s.DisplayName = &v
	return s
}

func (s *QueryJobCodeDictionaryResponseBodyContent) SetDoctorType(v string) *QueryJobCodeDictionaryResponseBodyContent {
	s.DoctorType = &v
	return s
}

type QueryJobCodeDictionaryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryJobCodeDictionaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryJobCodeDictionaryResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryJobCodeDictionaryResponse) GoString() string {
	return s.String()
}

func (s *QueryJobCodeDictionaryResponse) SetHeaders(v map[string]*string) *QueryJobCodeDictionaryResponse {
	s.Headers = v
	return s
}

func (s *QueryJobCodeDictionaryResponse) SetStatusCode(v int32) *QueryJobCodeDictionaryResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryJobCodeDictionaryResponse) SetBody(v *QueryJobCodeDictionaryResponseBody) *QueryJobCodeDictionaryResponse {
	s.Body = v
	return s
}

type QueryJobStatusCodeDictionaryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryJobStatusCodeDictionaryHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryJobStatusCodeDictionaryHeaders) GoString() string {
	return s.String()
}

func (s *QueryJobStatusCodeDictionaryHeaders) SetCommonHeaders(v map[string]*string) *QueryJobStatusCodeDictionaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryJobStatusCodeDictionaryHeaders) SetXAcsDingtalkAccessToken(v string) *QueryJobStatusCodeDictionaryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryJobStatusCodeDictionaryResponseBody struct {
	Content []*QueryJobStatusCodeDictionaryResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s QueryJobStatusCodeDictionaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryJobStatusCodeDictionaryResponseBody) GoString() string {
	return s.String()
}

func (s *QueryJobStatusCodeDictionaryResponseBody) SetContent(v []*QueryJobStatusCodeDictionaryResponseBodyContent) *QueryJobStatusCodeDictionaryResponseBody {
	s.Content = v
	return s
}

type QueryJobStatusCodeDictionaryResponseBodyContent struct {
	Category    *string `json:"category,omitempty" xml:"category,omitempty"`
	Code        *string `json:"code,omitempty" xml:"code,omitempty"`
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s QueryJobStatusCodeDictionaryResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryJobStatusCodeDictionaryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryJobStatusCodeDictionaryResponseBodyContent) SetCategory(v string) *QueryJobStatusCodeDictionaryResponseBodyContent {
	s.Category = &v
	return s
}

func (s *QueryJobStatusCodeDictionaryResponseBodyContent) SetCode(v string) *QueryJobStatusCodeDictionaryResponseBodyContent {
	s.Code = &v
	return s
}

func (s *QueryJobStatusCodeDictionaryResponseBodyContent) SetDisplayName(v string) *QueryJobStatusCodeDictionaryResponseBodyContent {
	s.DisplayName = &v
	return s
}

type QueryJobStatusCodeDictionaryResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryJobStatusCodeDictionaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryJobStatusCodeDictionaryResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryJobStatusCodeDictionaryResponse) GoString() string {
	return s.String()
}

func (s *QueryJobStatusCodeDictionaryResponse) SetHeaders(v map[string]*string) *QueryJobStatusCodeDictionaryResponse {
	s.Headers = v
	return s
}

func (s *QueryJobStatusCodeDictionaryResponse) SetStatusCode(v int32) *QueryJobStatusCodeDictionaryResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryJobStatusCodeDictionaryResponse) SetBody(v *QueryJobStatusCodeDictionaryResponseBody) *QueryJobStatusCodeDictionaryResponse {
	s.Body = v
	return s
}

type QueryMedicalEventsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryMedicalEventsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryMedicalEventsHeaders) GoString() string {
	return s.String()
}

func (s *QueryMedicalEventsHeaders) SetCommonHeaders(v map[string]*string) *QueryMedicalEventsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMedicalEventsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryMedicalEventsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryMedicalEventsResponseBody struct {
	Content    []*QueryMedicalEventsResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	Success    *bool                                    `json:"success,omitempty" xml:"success,omitempty"`
	TotalCount *int64                                   `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryMedicalEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMedicalEventsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMedicalEventsResponseBody) SetContent(v []*QueryMedicalEventsResponseBodyContent) *QueryMedicalEventsResponseBody {
	s.Content = v
	return s
}

func (s *QueryMedicalEventsResponseBody) SetSuccess(v bool) *QueryMedicalEventsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMedicalEventsResponseBody) SetTotalCount(v int64) *QueryMedicalEventsResponseBody {
	s.TotalCount = &v
	return s
}

type QueryMedicalEventsResponseBodyContent struct {
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	EventId *int64  `json:"eventId,omitempty" xml:"eventId,omitempty"`
}

func (s QueryMedicalEventsResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryMedicalEventsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryMedicalEventsResponseBodyContent) SetCode(v string) *QueryMedicalEventsResponseBodyContent {
	s.Code = &v
	return s
}

func (s *QueryMedicalEventsResponseBodyContent) SetContent(v string) *QueryMedicalEventsResponseBodyContent {
	s.Content = &v
	return s
}

func (s *QueryMedicalEventsResponseBodyContent) SetEventId(v int64) *QueryMedicalEventsResponseBodyContent {
	s.EventId = &v
	return s
}

type QueryMedicalEventsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMedicalEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMedicalEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMedicalEventsResponse) GoString() string {
	return s.String()
}

func (s *QueryMedicalEventsResponse) SetHeaders(v map[string]*string) *QueryMedicalEventsResponse {
	s.Headers = v
	return s
}

func (s *QueryMedicalEventsResponse) SetStatusCode(v int32) *QueryMedicalEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMedicalEventsResponse) SetBody(v *QueryMedicalEventsResponseBody) *QueryMedicalEventsResponse {
	s.Body = v
	return s
}

type QueryUserCredentialsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserCredentialsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserCredentialsHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserCredentialsHeaders) SetCommonHeaders(v map[string]*string) *QueryUserCredentialsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserCredentialsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserCredentialsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserCredentialsRequest struct {
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s QueryUserCredentialsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserCredentialsRequest) GoString() string {
	return s.String()
}

func (s *QueryUserCredentialsRequest) SetUserIds(v []*string) *QueryUserCredentialsRequest {
	s.UserIds = v
	return s
}

type QueryUserCredentialsResponseBody struct {
	Content []*QueryUserCredentialsResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s QueryUserCredentialsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserCredentialsResponseBody) SetContent(v []*QueryUserCredentialsResponseBodyContent) *QueryUserCredentialsResponseBody {
	s.Content = v
	return s
}

type QueryUserCredentialsResponseBodyContent struct {
	CredentialList []*QueryUserCredentialsResponseBodyContentCredentialList `json:"credentialList,omitempty" xml:"credentialList,omitempty" type:"Repeated"`
	UserId         *string                                                  `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryUserCredentialsResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryUserCredentialsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryUserCredentialsResponseBodyContent) SetCredentialList(v []*QueryUserCredentialsResponseBodyContentCredentialList) *QueryUserCredentialsResponseBodyContent {
	s.CredentialList = v
	return s
}

func (s *QueryUserCredentialsResponseBodyContent) SetUserId(v string) *QueryUserCredentialsResponseBodyContent {
	s.UserId = &v
	return s
}

type QueryUserCredentialsResponseBodyContentCredentialList struct {
	CredentialName *string `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	CredentialType *int32  `json:"credentialType,omitempty" xml:"credentialType,omitempty"`
	TermOfValidity *string `json:"termOfValidity,omitempty" xml:"termOfValidity,omitempty"`
}

func (s QueryUserCredentialsResponseBodyContentCredentialList) String() string {
	return tea.Prettify(s)
}

func (s QueryUserCredentialsResponseBodyContentCredentialList) GoString() string {
	return s.String()
}

func (s *QueryUserCredentialsResponseBodyContentCredentialList) SetCredentialName(v string) *QueryUserCredentialsResponseBodyContentCredentialList {
	s.CredentialName = &v
	return s
}

func (s *QueryUserCredentialsResponseBodyContentCredentialList) SetCredentialType(v int32) *QueryUserCredentialsResponseBodyContentCredentialList {
	s.CredentialType = &v
	return s
}

func (s *QueryUserCredentialsResponseBodyContentCredentialList) SetTermOfValidity(v string) *QueryUserCredentialsResponseBodyContentCredentialList {
	s.TermOfValidity = &v
	return s
}

type QueryUserCredentialsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUserCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserCredentialsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserCredentialsResponse) GoString() string {
	return s.String()
}

func (s *QueryUserCredentialsResponse) SetHeaders(v map[string]*string) *QueryUserCredentialsResponse {
	s.Headers = v
	return s
}

func (s *QueryUserCredentialsResponse) SetStatusCode(v int32) *QueryUserCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserCredentialsResponse) SetBody(v *QueryUserCredentialsResponseBody) *QueryUserCredentialsResponse {
	s.Body = v
	return s
}

type QueryUserExtInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserExtInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserExtInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserExtInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryUserExtInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserExtInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserExtInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserExtInfoResponseBody struct {
	Content []*QueryUserExtInfoResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s QueryUserExtInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserExtInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserExtInfoResponseBody) SetContent(v []*QueryUserExtInfoResponseBodyContent) *QueryUserExtInfoResponseBody {
	s.Content = v
	return s
}

type QueryUserExtInfoResponseBodyContent struct {
	GmtCreate             *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified           *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Status                *int32  `json:"status,omitempty" xml:"status,omitempty"`
	UserCode              *string `json:"userCode,omitempty" xml:"userCode,omitempty"`
	UserExtendDisplayName *string `json:"userExtendDisplayName,omitempty" xml:"userExtendDisplayName,omitempty"`
	UserExtendKey         *string `json:"userExtendKey,omitempty" xml:"userExtendKey,omitempty"`
	UserExtendValue       *string `json:"userExtendValue,omitempty" xml:"userExtendValue,omitempty"`
}

func (s QueryUserExtInfoResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryUserExtInfoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryUserExtInfoResponseBodyContent) SetGmtCreate(v string) *QueryUserExtInfoResponseBodyContent {
	s.GmtCreate = &v
	return s
}

func (s *QueryUserExtInfoResponseBodyContent) SetGmtModified(v string) *QueryUserExtInfoResponseBodyContent {
	s.GmtModified = &v
	return s
}

func (s *QueryUserExtInfoResponseBodyContent) SetStatus(v int32) *QueryUserExtInfoResponseBodyContent {
	s.Status = &v
	return s
}

func (s *QueryUserExtInfoResponseBodyContent) SetUserCode(v string) *QueryUserExtInfoResponseBodyContent {
	s.UserCode = &v
	return s
}

func (s *QueryUserExtInfoResponseBodyContent) SetUserExtendDisplayName(v string) *QueryUserExtInfoResponseBodyContent {
	s.UserExtendDisplayName = &v
	return s
}

func (s *QueryUserExtInfoResponseBodyContent) SetUserExtendKey(v string) *QueryUserExtInfoResponseBodyContent {
	s.UserExtendKey = &v
	return s
}

func (s *QueryUserExtInfoResponseBodyContent) SetUserExtendValue(v string) *QueryUserExtInfoResponseBodyContent {
	s.UserExtendValue = &v
	return s
}

type QueryUserExtInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUserExtInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserExtInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserExtInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryUserExtInfoResponse) SetHeaders(v map[string]*string) *QueryUserExtInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryUserExtInfoResponse) SetStatusCode(v int32) *QueryUserExtInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserExtInfoResponse) SetBody(v *QueryUserExtInfoResponseBody) *QueryUserExtInfoResponse {
	s.Body = v
	return s
}

type QueryUserExtendValuesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserExtendValuesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserExtendValuesHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserExtendValuesHeaders) SetCommonHeaders(v map[string]*string) *QueryUserExtendValuesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserExtendValuesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserExtendValuesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserExtendValuesRequest struct {
	UserExtendKey *string   `json:"userExtendKey,omitempty" xml:"userExtendKey,omitempty"`
	UserIds       []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s QueryUserExtendValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserExtendValuesRequest) GoString() string {
	return s.String()
}

func (s *QueryUserExtendValuesRequest) SetUserExtendKey(v string) *QueryUserExtendValuesRequest {
	s.UserExtendKey = &v
	return s
}

func (s *QueryUserExtendValuesRequest) SetUserIds(v []*string) *QueryUserExtendValuesRequest {
	s.UserIds = v
	return s
}

type QueryUserExtendValuesResponseBody struct {
	Content    []*QueryUserExtendValuesResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	Success    *bool                                       `json:"success,omitempty" xml:"success,omitempty"`
	TotalCount *int64                                      `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s QueryUserExtendValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserExtendValuesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserExtendValuesResponseBody) SetContent(v []*QueryUserExtendValuesResponseBodyContent) *QueryUserExtendValuesResponseBody {
	s.Content = v
	return s
}

func (s *QueryUserExtendValuesResponseBody) SetSuccess(v bool) *QueryUserExtendValuesResponseBody {
	s.Success = &v
	return s
}

func (s *QueryUserExtendValuesResponseBody) SetTotalCount(v int64) *QueryUserExtendValuesResponseBody {
	s.TotalCount = &v
	return s
}

type QueryUserExtendValuesResponseBodyContent struct {
	UserCode              *string `json:"userCode,omitempty" xml:"userCode,omitempty"`
	UserExtendDisplayName *string `json:"userExtendDisplayName,omitempty" xml:"userExtendDisplayName,omitempty"`
	UserExtendKey         *string `json:"userExtendKey,omitempty" xml:"userExtendKey,omitempty"`
	UserExtendValue       *string `json:"userExtendValue,omitempty" xml:"userExtendValue,omitempty"`
}

func (s QueryUserExtendValuesResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryUserExtendValuesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryUserExtendValuesResponseBodyContent) SetUserCode(v string) *QueryUserExtendValuesResponseBodyContent {
	s.UserCode = &v
	return s
}

func (s *QueryUserExtendValuesResponseBodyContent) SetUserExtendDisplayName(v string) *QueryUserExtendValuesResponseBodyContent {
	s.UserExtendDisplayName = &v
	return s
}

func (s *QueryUserExtendValuesResponseBodyContent) SetUserExtendKey(v string) *QueryUserExtendValuesResponseBodyContent {
	s.UserExtendKey = &v
	return s
}

func (s *QueryUserExtendValuesResponseBodyContent) SetUserExtendValue(v string) *QueryUserExtendValuesResponseBodyContent {
	s.UserExtendValue = &v
	return s
}

type QueryUserExtendValuesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUserExtendValuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserExtendValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserExtendValuesResponse) GoString() string {
	return s.String()
}

func (s *QueryUserExtendValuesResponse) SetHeaders(v map[string]*string) *QueryUserExtendValuesResponse {
	s.Headers = v
	return s
}

func (s *QueryUserExtendValuesResponse) SetStatusCode(v int32) *QueryUserExtendValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserExtendValuesResponse) SetBody(v *QueryUserExtendValuesResponseBody) *QueryUserExtendValuesResponse {
	s.Body = v
	return s
}

type QueryUserInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryUserInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserInfoHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserInfoRequest struct {
	MonthMark *string `json:"monthMark,omitempty" xml:"monthMark,omitempty"`
}

func (s QueryUserInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryUserInfoRequest) SetMonthMark(v string) *QueryUserInfoRequest {
	s.MonthMark = &v
	return s
}

type QueryUserInfoResponseBody struct {
	Content *QueryUserInfoResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
}

func (s QueryUserInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponseBody) SetContent(v *QueryUserInfoResponseBodyContent) *QueryUserInfoResponseBody {
	s.Content = v
	return s
}

type QueryUserInfoResponseBodyContent struct {
	Comments      *string                                          `json:"comments,omitempty" xml:"comments,omitempty"`
	Dept          []*QueryUserInfoResponseBodyContentDept          `json:"dept,omitempty" xml:"dept,omitempty" type:"Repeated"`
	Group         []*QueryUserInfoResponseBodyContentGroup         `json:"group,omitempty" xml:"group,omitempty" type:"Repeated"`
	Job           *QueryUserInfoResponseBodyContentJob             `json:"job,omitempty" xml:"job,omitempty" type:"Struct"`
	JobNum        *string                                          `json:"jobNum,omitempty" xml:"jobNum,omitempty"`
	JobStatus     *QueryUserInfoResponseBodyContentJobStatus       `json:"jobStatus,omitempty" xml:"jobStatus,omitempty" type:"Struct"`
	JobStatusList []*QueryUserInfoResponseBodyContentJobStatusList `json:"jobStatusList,omitempty" xml:"jobStatusList,omitempty" type:"Repeated"`
	Uid           *string                                          `json:"uid,omitempty" xml:"uid,omitempty"`
	UserName      *string                                          `json:"userName,omitempty" xml:"userName,omitempty"`
	UserProb      *QueryUserInfoResponseBodyContentUserProb        `json:"userProb,omitempty" xml:"userProb,omitempty" type:"Struct"`
}

func (s QueryUserInfoResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponseBodyContent) SetComments(v string) *QueryUserInfoResponseBodyContent {
	s.Comments = &v
	return s
}

func (s *QueryUserInfoResponseBodyContent) SetDept(v []*QueryUserInfoResponseBodyContentDept) *QueryUserInfoResponseBodyContent {
	s.Dept = v
	return s
}

func (s *QueryUserInfoResponseBodyContent) SetGroup(v []*QueryUserInfoResponseBodyContentGroup) *QueryUserInfoResponseBodyContent {
	s.Group = v
	return s
}

func (s *QueryUserInfoResponseBodyContent) SetJob(v *QueryUserInfoResponseBodyContentJob) *QueryUserInfoResponseBodyContent {
	s.Job = v
	return s
}

func (s *QueryUserInfoResponseBodyContent) SetJobNum(v string) *QueryUserInfoResponseBodyContent {
	s.JobNum = &v
	return s
}

func (s *QueryUserInfoResponseBodyContent) SetJobStatus(v *QueryUserInfoResponseBodyContentJobStatus) *QueryUserInfoResponseBodyContent {
	s.JobStatus = v
	return s
}

func (s *QueryUserInfoResponseBodyContent) SetJobStatusList(v []*QueryUserInfoResponseBodyContentJobStatusList) *QueryUserInfoResponseBodyContent {
	s.JobStatusList = v
	return s
}

func (s *QueryUserInfoResponseBodyContent) SetUid(v string) *QueryUserInfoResponseBodyContent {
	s.Uid = &v
	return s
}

func (s *QueryUserInfoResponseBodyContent) SetUserName(v string) *QueryUserInfoResponseBodyContent {
	s.UserName = &v
	return s
}

func (s *QueryUserInfoResponseBodyContent) SetUserProb(v *QueryUserInfoResponseBodyContentUserProb) *QueryUserInfoResponseBodyContent {
	s.UserProb = v
	return s
}

type QueryUserInfoResponseBodyContentDept struct {
	GmtCreateStr   *string `json:"gmtCreateStr,omitempty" xml:"gmtCreateStr,omitempty"`
	GmtModifiedStr *string `json:"gmtModifiedStr,omitempty" xml:"gmtModifiedStr,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	RelId          *int64  `json:"relId,omitempty" xml:"relId,omitempty"`
}

func (s QueryUserInfoResponseBodyContentDept) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponseBodyContentDept) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponseBodyContentDept) SetGmtCreateStr(v string) *QueryUserInfoResponseBodyContentDept {
	s.GmtCreateStr = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentDept) SetGmtModifiedStr(v string) *QueryUserInfoResponseBodyContentDept {
	s.GmtModifiedStr = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentDept) SetId(v int64) *QueryUserInfoResponseBodyContentDept {
	s.Id = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentDept) SetName(v string) *QueryUserInfoResponseBodyContentDept {
	s.Name = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentDept) SetRelId(v int64) *QueryUserInfoResponseBodyContentDept {
	s.RelId = &v
	return s
}

type QueryUserInfoResponseBodyContentGroup struct {
	DeptId         *int64  `json:"deptId,omitempty" xml:"deptId,omitempty"`
	DeptName       *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	GmtCreateStr   *string `json:"gmtCreateStr,omitempty" xml:"gmtCreateStr,omitempty"`
	GmtModifiedStr *string `json:"gmtModifiedStr,omitempty" xml:"gmtModifiedStr,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	RelId          *int64  `json:"relId,omitempty" xml:"relId,omitempty"`
}

func (s QueryUserInfoResponseBodyContentGroup) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponseBodyContentGroup) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponseBodyContentGroup) SetDeptId(v int64) *QueryUserInfoResponseBodyContentGroup {
	s.DeptId = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentGroup) SetDeptName(v string) *QueryUserInfoResponseBodyContentGroup {
	s.DeptName = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentGroup) SetGmtCreateStr(v string) *QueryUserInfoResponseBodyContentGroup {
	s.GmtCreateStr = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentGroup) SetGmtModifiedStr(v string) *QueryUserInfoResponseBodyContentGroup {
	s.GmtModifiedStr = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentGroup) SetId(v int64) *QueryUserInfoResponseBodyContentGroup {
	s.Id = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentGroup) SetName(v string) *QueryUserInfoResponseBodyContentGroup {
	s.Name = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentGroup) SetRelId(v int64) *QueryUserInfoResponseBodyContentGroup {
	s.RelId = &v
	return s
}

type QueryUserInfoResponseBodyContentJob struct {
	BizType     *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	Category    *string `json:"category,omitempty" xml:"category,omitempty"`
	Code        *string `json:"code,omitempty" xml:"code,omitempty"`
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s QueryUserInfoResponseBodyContentJob) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponseBodyContentJob) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponseBodyContentJob) SetBizType(v string) *QueryUserInfoResponseBodyContentJob {
	s.BizType = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentJob) SetCategory(v string) *QueryUserInfoResponseBodyContentJob {
	s.Category = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentJob) SetCode(v string) *QueryUserInfoResponseBodyContentJob {
	s.Code = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentJob) SetDisplayName(v string) *QueryUserInfoResponseBodyContentJob {
	s.DisplayName = &v
	return s
}

type QueryUserInfoResponseBodyContentJobStatus struct {
	BizType     *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	Category    *string `json:"category,omitempty" xml:"category,omitempty"`
	Code        *string `json:"code,omitempty" xml:"code,omitempty"`
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s QueryUserInfoResponseBodyContentJobStatus) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponseBodyContentJobStatus) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponseBodyContentJobStatus) SetBizType(v string) *QueryUserInfoResponseBodyContentJobStatus {
	s.BizType = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentJobStatus) SetCategory(v string) *QueryUserInfoResponseBodyContentJobStatus {
	s.Category = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentJobStatus) SetCode(v string) *QueryUserInfoResponseBodyContentJobStatus {
	s.Code = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentJobStatus) SetDisplayName(v string) *QueryUserInfoResponseBodyContentJobStatus {
	s.DisplayName = &v
	return s
}

type QueryUserInfoResponseBodyContentJobStatusList struct {
	BizType     *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	Category    *string `json:"category,omitempty" xml:"category,omitempty"`
	Code        *string `json:"code,omitempty" xml:"code,omitempty"`
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s QueryUserInfoResponseBodyContentJobStatusList) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponseBodyContentJobStatusList) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponseBodyContentJobStatusList) SetBizType(v string) *QueryUserInfoResponseBodyContentJobStatusList {
	s.BizType = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentJobStatusList) SetCategory(v string) *QueryUserInfoResponseBodyContentJobStatusList {
	s.Category = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentJobStatusList) SetCode(v string) *QueryUserInfoResponseBodyContentJobStatusList {
	s.Code = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentJobStatusList) SetDisplayName(v string) *QueryUserInfoResponseBodyContentJobStatusList {
	s.DisplayName = &v
	return s
}

type QueryUserInfoResponseBodyContentUserProb struct {
	BizType     *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	Category    *string `json:"category,omitempty" xml:"category,omitempty"`
	Code        *string `json:"code,omitempty" xml:"code,omitempty"`
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s QueryUserInfoResponseBodyContentUserProb) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponseBodyContentUserProb) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponseBodyContentUserProb) SetBizType(v string) *QueryUserInfoResponseBodyContentUserProb {
	s.BizType = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentUserProb) SetCategory(v string) *QueryUserInfoResponseBodyContentUserProb {
	s.Category = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentUserProb) SetCode(v string) *QueryUserInfoResponseBodyContentUserProb {
	s.Code = &v
	return s
}

func (s *QueryUserInfoResponseBodyContentUserProb) SetDisplayName(v string) *QueryUserInfoResponseBodyContentUserProb {
	s.DisplayName = &v
	return s
}

type QueryUserInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryUserInfoResponse) SetHeaders(v map[string]*string) *QueryUserInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryUserInfoResponse) SetStatusCode(v int32) *QueryUserInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserInfoResponse) SetBody(v *QueryUserInfoResponseBody) *QueryUserInfoResponse {
	s.Body = v
	return s
}

type QueryUserProbCodeDictionaryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserProbCodeDictionaryHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserProbCodeDictionaryHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserProbCodeDictionaryHeaders) SetCommonHeaders(v map[string]*string) *QueryUserProbCodeDictionaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserProbCodeDictionaryHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserProbCodeDictionaryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserProbCodeDictionaryResponseBody struct {
	Content []*QueryUserProbCodeDictionaryResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s QueryUserProbCodeDictionaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserProbCodeDictionaryResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserProbCodeDictionaryResponseBody) SetContent(v []*QueryUserProbCodeDictionaryResponseBodyContent) *QueryUserProbCodeDictionaryResponseBody {
	s.Content = v
	return s
}

type QueryUserProbCodeDictionaryResponseBodyContent struct {
	Category    *string `json:"category,omitempty" xml:"category,omitempty"`
	Code        *string `json:"code,omitempty" xml:"code,omitempty"`
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s QueryUserProbCodeDictionaryResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryUserProbCodeDictionaryResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryUserProbCodeDictionaryResponseBodyContent) SetCategory(v string) *QueryUserProbCodeDictionaryResponseBodyContent {
	s.Category = &v
	return s
}

func (s *QueryUserProbCodeDictionaryResponseBodyContent) SetCode(v string) *QueryUserProbCodeDictionaryResponseBodyContent {
	s.Code = &v
	return s
}

func (s *QueryUserProbCodeDictionaryResponseBodyContent) SetDisplayName(v string) *QueryUserProbCodeDictionaryResponseBodyContent {
	s.DisplayName = &v
	return s
}

type QueryUserProbCodeDictionaryResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUserProbCodeDictionaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserProbCodeDictionaryResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserProbCodeDictionaryResponse) GoString() string {
	return s.String()
}

func (s *QueryUserProbCodeDictionaryResponse) SetHeaders(v map[string]*string) *QueryUserProbCodeDictionaryResponse {
	s.Headers = v
	return s
}

func (s *QueryUserProbCodeDictionaryResponse) SetStatusCode(v int32) *QueryUserProbCodeDictionaryResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserProbCodeDictionaryResponse) SetBody(v *QueryUserProbCodeDictionaryResponseBody) *QueryUserProbCodeDictionaryResponse {
	s.Body = v
	return s
}

type QueryUserRolesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryUserRolesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryUserRolesHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserRolesHeaders) SetCommonHeaders(v map[string]*string) *QueryUserRolesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserRolesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryUserRolesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryUserRolesResponseBody struct {
	Content []*QueryUserRolesResponseBodyContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
}

func (s QueryUserRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserRolesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserRolesResponseBody) SetContent(v []*QueryUserRolesResponseBodyContent) *QueryUserRolesResponseBody {
	s.Content = v
	return s
}

type QueryUserRolesResponseBodyContent struct {
	RoleCode *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
}

func (s QueryUserRolesResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s QueryUserRolesResponseBodyContent) GoString() string {
	return s.String()
}

func (s *QueryUserRolesResponseBodyContent) SetRoleCode(v string) *QueryUserRolesResponseBodyContent {
	s.RoleCode = &v
	return s
}

func (s *QueryUserRolesResponseBodyContent) SetRoleName(v string) *QueryUserRolesResponseBodyContent {
	s.RoleName = &v
	return s
}

type QueryUserRolesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryUserRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserRolesResponse) GoString() string {
	return s.String()
}

func (s *QueryUserRolesResponse) SetHeaders(v map[string]*string) *QueryUserRolesResponse {
	s.Headers = v
	return s
}

func (s *QueryUserRolesResponse) SetStatusCode(v int32) *QueryUserRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserRolesResponse) SetBody(v *QueryUserRolesResponseBody) *QueryUserRolesResponse {
	s.Body = v
	return s
}

type SaveUserExtendValuesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SaveUserExtendValuesHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveUserExtendValuesHeaders) GoString() string {
	return s.String()
}

func (s *SaveUserExtendValuesHeaders) SetCommonHeaders(v map[string]*string) *SaveUserExtendValuesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveUserExtendValuesHeaders) SetXAcsDingtalkAccessToken(v string) *SaveUserExtendValuesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SaveUserExtendValuesRequest struct {
	UserDisplayName *string `json:"userDisplayName,omitempty" xml:"userDisplayName,omitempty"`
	UserExtendKey   *string `json:"userExtendKey,omitempty" xml:"userExtendKey,omitempty"`
	UserExtendValue *string `json:"userExtendValue,omitempty" xml:"userExtendValue,omitempty"`
}

func (s SaveUserExtendValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveUserExtendValuesRequest) GoString() string {
	return s.String()
}

func (s *SaveUserExtendValuesRequest) SetUserDisplayName(v string) *SaveUserExtendValuesRequest {
	s.UserDisplayName = &v
	return s
}

func (s *SaveUserExtendValuesRequest) SetUserExtendKey(v string) *SaveUserExtendValuesRequest {
	s.UserExtendKey = &v
	return s
}

func (s *SaveUserExtendValuesRequest) SetUserExtendValue(v string) *SaveUserExtendValuesRequest {
	s.UserExtendValue = &v
	return s
}

type SaveUserExtendValuesResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SaveUserExtendValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveUserExtendValuesResponseBody) GoString() string {
	return s.String()
}

func (s *SaveUserExtendValuesResponseBody) SetSuccess(v bool) *SaveUserExtendValuesResponseBody {
	s.Success = &v
	return s
}

type SaveUserExtendValuesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveUserExtendValuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveUserExtendValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveUserExtendValuesResponse) GoString() string {
	return s.String()
}

func (s *SaveUserExtendValuesResponse) SetHeaders(v map[string]*string) *SaveUserExtendValuesResponse {
	s.Headers = v
	return s
}

func (s *SaveUserExtendValuesResponse) SetStatusCode(v int32) *SaveUserExtendValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveUserExtendValuesResponse) SetBody(v *SaveUserExtendValuesResponseBody) *SaveUserExtendValuesResponse {
	s.Body = v
	return s
}

type SupplyAddDeptHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SupplyAddDeptHeaders) String() string {
	return tea.Prettify(s)
}

func (s SupplyAddDeptHeaders) GoString() string {
	return s.String()
}

func (s *SupplyAddDeptHeaders) SetCommonHeaders(v map[string]*string) *SupplyAddDeptHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SupplyAddDeptHeaders) SetXAcsDingtalkAccessToken(v string) *SupplyAddDeptHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SupplyAddDeptRequest struct {
	DeptName       *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	PartnerNumber  *string `json:"partnerNumber,omitempty" xml:"partnerNumber,omitempty"`
	SuperDeptId    *int64  `json:"superDeptId,omitempty" xml:"superDeptId,omitempty"`
	SupplyDeptType *string `json:"supplyDeptType,omitempty" xml:"supplyDeptType,omitempty"`
}

func (s SupplyAddDeptRequest) String() string {
	return tea.Prettify(s)
}

func (s SupplyAddDeptRequest) GoString() string {
	return s.String()
}

func (s *SupplyAddDeptRequest) SetDeptName(v string) *SupplyAddDeptRequest {
	s.DeptName = &v
	return s
}

func (s *SupplyAddDeptRequest) SetPartnerNumber(v string) *SupplyAddDeptRequest {
	s.PartnerNumber = &v
	return s
}

func (s *SupplyAddDeptRequest) SetSuperDeptId(v int64) *SupplyAddDeptRequest {
	s.SuperDeptId = &v
	return s
}

func (s *SupplyAddDeptRequest) SetSupplyDeptType(v string) *SupplyAddDeptRequest {
	s.SupplyDeptType = &v
	return s
}

type SupplyAddDeptResponseBody struct {
	Result *SupplyAddDeptResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s SupplyAddDeptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SupplyAddDeptResponseBody) GoString() string {
	return s.String()
}

func (s *SupplyAddDeptResponseBody) SetResult(v *SupplyAddDeptResponseBodyResult) *SupplyAddDeptResponseBody {
	s.Result = v
	return s
}

type SupplyAddDeptResponseBodyResult struct {
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
}

func (s SupplyAddDeptResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SupplyAddDeptResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SupplyAddDeptResponseBodyResult) SetDeptId(v int64) *SupplyAddDeptResponseBodyResult {
	s.DeptId = &v
	return s
}

type SupplyAddDeptResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SupplyAddDeptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SupplyAddDeptResponse) String() string {
	return tea.Prettify(s)
}

func (s SupplyAddDeptResponse) GoString() string {
	return s.String()
}

func (s *SupplyAddDeptResponse) SetHeaders(v map[string]*string) *SupplyAddDeptResponse {
	s.Headers = v
	return s
}

func (s *SupplyAddDeptResponse) SetStatusCode(v int32) *SupplyAddDeptResponse {
	s.StatusCode = &v
	return s
}

func (s *SupplyAddDeptResponse) SetBody(v *SupplyAddDeptResponseBody) *SupplyAddDeptResponse {
	s.Body = v
	return s
}

type SupplyAddMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SupplyAddMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s SupplyAddMemberHeaders) GoString() string {
	return s.String()
}

func (s *SupplyAddMemberHeaders) SetCommonHeaders(v map[string]*string) *SupplyAddMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SupplyAddMemberHeaders) SetXAcsDingtalkAccessToken(v string) *SupplyAddMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SupplyAddMemberRequest struct {
	IsPartnerManager *bool   `json:"isPartnerManager,omitempty" xml:"isPartnerManager,omitempty"`
	MemberMobile     *string `json:"memberMobile,omitempty" xml:"memberMobile,omitempty"`
	MemberName       *string `json:"memberName,omitempty" xml:"memberName,omitempty"`
	MemberWorkNumber *string `json:"memberWorkNumber,omitempty" xml:"memberWorkNumber,omitempty"`
	SupplyDeptId     *int64  `json:"supplyDeptId,omitempty" xml:"supplyDeptId,omitempty"`
}

func (s SupplyAddMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s SupplyAddMemberRequest) GoString() string {
	return s.String()
}

func (s *SupplyAddMemberRequest) SetIsPartnerManager(v bool) *SupplyAddMemberRequest {
	s.IsPartnerManager = &v
	return s
}

func (s *SupplyAddMemberRequest) SetMemberMobile(v string) *SupplyAddMemberRequest {
	s.MemberMobile = &v
	return s
}

func (s *SupplyAddMemberRequest) SetMemberName(v string) *SupplyAddMemberRequest {
	s.MemberName = &v
	return s
}

func (s *SupplyAddMemberRequest) SetMemberWorkNumber(v string) *SupplyAddMemberRequest {
	s.MemberWorkNumber = &v
	return s
}

func (s *SupplyAddMemberRequest) SetSupplyDeptId(v int64) *SupplyAddMemberRequest {
	s.SupplyDeptId = &v
	return s
}

type SupplyAddMemberResponseBody struct {
	Result *SupplyAddMemberResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s SupplyAddMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SupplyAddMemberResponseBody) GoString() string {
	return s.String()
}

func (s *SupplyAddMemberResponseBody) SetResult(v *SupplyAddMemberResponseBodyResult) *SupplyAddMemberResponseBody {
	s.Result = v
	return s
}

type SupplyAddMemberResponseBodyResult struct {
	DingMemberStatus *string `json:"dingMemberStatus,omitempty" xml:"dingMemberStatus,omitempty"`
	UnionId          *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId           *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SupplyAddMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SupplyAddMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SupplyAddMemberResponseBodyResult) SetDingMemberStatus(v string) *SupplyAddMemberResponseBodyResult {
	s.DingMemberStatus = &v
	return s
}

func (s *SupplyAddMemberResponseBodyResult) SetUnionId(v string) *SupplyAddMemberResponseBodyResult {
	s.UnionId = &v
	return s
}

func (s *SupplyAddMemberResponseBodyResult) SetUserId(v string) *SupplyAddMemberResponseBodyResult {
	s.UserId = &v
	return s
}

type SupplyAddMemberResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SupplyAddMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SupplyAddMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s SupplyAddMemberResponse) GoString() string {
	return s.String()
}

func (s *SupplyAddMemberResponse) SetHeaders(v map[string]*string) *SupplyAddMemberResponse {
	s.Headers = v
	return s
}

func (s *SupplyAddMemberResponse) SetStatusCode(v int32) *SupplyAddMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *SupplyAddMemberResponse) SetBody(v *SupplyAddMemberResponseBody) *SupplyAddMemberResponse {
	s.Body = v
	return s
}

type SupplyGetMemberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SupplyGetMemberHeaders) String() string {
	return tea.Prettify(s)
}

func (s SupplyGetMemberHeaders) GoString() string {
	return s.String()
}

func (s *SupplyGetMemberHeaders) SetCommonHeaders(v map[string]*string) *SupplyGetMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SupplyGetMemberHeaders) SetXAcsDingtalkAccessToken(v string) *SupplyGetMemberHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SupplyGetMemberRequest struct {
	UnionId *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId  *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SupplyGetMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s SupplyGetMemberRequest) GoString() string {
	return s.String()
}

func (s *SupplyGetMemberRequest) SetUnionId(v string) *SupplyGetMemberRequest {
	s.UnionId = &v
	return s
}

func (s *SupplyGetMemberRequest) SetUserId(v string) *SupplyGetMemberRequest {
	s.UserId = &v
	return s
}

type SupplyGetMemberResponseBody struct {
	Result *SupplyGetMemberResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s SupplyGetMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SupplyGetMemberResponseBody) GoString() string {
	return s.String()
}

func (s *SupplyGetMemberResponseBody) SetResult(v *SupplyGetMemberResponseBodyResult) *SupplyGetMemberResponseBody {
	s.Result = v
	return s
}

type SupplyGetMemberResponseBodyResult struct {
	DingMemberStatus *string `json:"dingMemberStatus,omitempty" xml:"dingMemberStatus,omitempty"`
	IsActive         *bool   `json:"isActive,omitempty" xml:"isActive,omitempty"`
	MemberName       *string `json:"memberName,omitempty" xml:"memberName,omitempty"`
	MemberWorkNumber *string `json:"memberWorkNumber,omitempty" xml:"memberWorkNumber,omitempty"`
}

func (s SupplyGetMemberResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SupplyGetMemberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SupplyGetMemberResponseBodyResult) SetDingMemberStatus(v string) *SupplyGetMemberResponseBodyResult {
	s.DingMemberStatus = &v
	return s
}

func (s *SupplyGetMemberResponseBodyResult) SetIsActive(v bool) *SupplyGetMemberResponseBodyResult {
	s.IsActive = &v
	return s
}

func (s *SupplyGetMemberResponseBodyResult) SetMemberName(v string) *SupplyGetMemberResponseBodyResult {
	s.MemberName = &v
	return s
}

func (s *SupplyGetMemberResponseBodyResult) SetMemberWorkNumber(v string) *SupplyGetMemberResponseBodyResult {
	s.MemberWorkNumber = &v
	return s
}

type SupplyGetMemberResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SupplyGetMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SupplyGetMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s SupplyGetMemberResponse) GoString() string {
	return s.String()
}

func (s *SupplyGetMemberResponse) SetHeaders(v map[string]*string) *SupplyGetMemberResponse {
	s.Headers = v
	return s
}

func (s *SupplyGetMemberResponse) SetStatusCode(v int32) *SupplyGetMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *SupplyGetMemberResponse) SetBody(v *SupplyGetMemberResponseBody) *SupplyGetMemberResponse {
	s.Body = v
	return s
}

type SupplyListDeptMembersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SupplyListDeptMembersHeaders) String() string {
	return tea.Prettify(s)
}

func (s SupplyListDeptMembersHeaders) GoString() string {
	return s.String()
}

func (s *SupplyListDeptMembersHeaders) SetCommonHeaders(v map[string]*string) *SupplyListDeptMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SupplyListDeptMembersHeaders) SetXAcsDingtalkAccessToken(v string) *SupplyListDeptMembersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SupplyListDeptMembersRequest struct {
	PageNumber   *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize     *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	SupplyDeptId *int64 `json:"supplyDeptId,omitempty" xml:"supplyDeptId,omitempty"`
}

func (s SupplyListDeptMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s SupplyListDeptMembersRequest) GoString() string {
	return s.String()
}

func (s *SupplyListDeptMembersRequest) SetPageNumber(v int64) *SupplyListDeptMembersRequest {
	s.PageNumber = &v
	return s
}

func (s *SupplyListDeptMembersRequest) SetPageSize(v int64) *SupplyListDeptMembersRequest {
	s.PageSize = &v
	return s
}

func (s *SupplyListDeptMembersRequest) SetSupplyDeptId(v int64) *SupplyListDeptMembersRequest {
	s.SupplyDeptId = &v
	return s
}

type SupplyListDeptMembersResponseBody struct {
	HasMore *bool                                    `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List    []*SupplyListDeptMembersResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s SupplyListDeptMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SupplyListDeptMembersResponseBody) GoString() string {
	return s.String()
}

func (s *SupplyListDeptMembersResponseBody) SetHasMore(v bool) *SupplyListDeptMembersResponseBody {
	s.HasMore = &v
	return s
}

func (s *SupplyListDeptMembersResponseBody) SetList(v []*SupplyListDeptMembersResponseBodyList) *SupplyListDeptMembersResponseBody {
	s.List = v
	return s
}

type SupplyListDeptMembersResponseBodyList struct {
	DingMemberStatus *string `json:"dingMemberStatus,omitempty" xml:"dingMemberStatus,omitempty"`
	IsActive         *bool   `json:"isActive,omitempty" xml:"isActive,omitempty"`
	MemberName       *string `json:"memberName,omitempty" xml:"memberName,omitempty"`
	MemberWorkNumber *string `json:"memberWorkNumber,omitempty" xml:"memberWorkNumber,omitempty"`
	UnionId          *string `json:"unionId,omitempty" xml:"unionId,omitempty"`
	UserId           *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SupplyListDeptMembersResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s SupplyListDeptMembersResponseBodyList) GoString() string {
	return s.String()
}

func (s *SupplyListDeptMembersResponseBodyList) SetDingMemberStatus(v string) *SupplyListDeptMembersResponseBodyList {
	s.DingMemberStatus = &v
	return s
}

func (s *SupplyListDeptMembersResponseBodyList) SetIsActive(v bool) *SupplyListDeptMembersResponseBodyList {
	s.IsActive = &v
	return s
}

func (s *SupplyListDeptMembersResponseBodyList) SetMemberName(v string) *SupplyListDeptMembersResponseBodyList {
	s.MemberName = &v
	return s
}

func (s *SupplyListDeptMembersResponseBodyList) SetMemberWorkNumber(v string) *SupplyListDeptMembersResponseBodyList {
	s.MemberWorkNumber = &v
	return s
}

func (s *SupplyListDeptMembersResponseBodyList) SetUnionId(v string) *SupplyListDeptMembersResponseBodyList {
	s.UnionId = &v
	return s
}

func (s *SupplyListDeptMembersResponseBodyList) SetUserId(v string) *SupplyListDeptMembersResponseBodyList {
	s.UserId = &v
	return s
}

type SupplyListDeptMembersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SupplyListDeptMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SupplyListDeptMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s SupplyListDeptMembersResponse) GoString() string {
	return s.String()
}

func (s *SupplyListDeptMembersResponse) SetHeaders(v map[string]*string) *SupplyListDeptMembersResponse {
	s.Headers = v
	return s
}

func (s *SupplyListDeptMembersResponse) SetStatusCode(v int32) *SupplyListDeptMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *SupplyListDeptMembersResponse) SetBody(v *SupplyListDeptMembersResponseBody) *SupplyListDeptMembersResponse {
	s.Body = v
	return s
}

type UpdateUserExtendInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateUserExtendInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserExtendInfoHeaders) GoString() string {
	return s.String()
}

func (s *UpdateUserExtendInfoHeaders) SetCommonHeaders(v map[string]*string) *UpdateUserExtendInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateUserExtendInfoHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateUserExtendInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateUserExtendInfoRequest struct {
	Comments      *string   `json:"comments,omitempty" xml:"comments,omitempty"`
	JobCode       *string   `json:"jobCode,omitempty" xml:"jobCode,omitempty"`
	JobStatusCode []*string `json:"jobStatusCode,omitempty" xml:"jobStatusCode,omitempty" type:"Repeated"`
	UserProbCode  *string   `json:"userProbCode,omitempty" xml:"userProbCode,omitempty"`
}

func (s UpdateUserExtendInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserExtendInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserExtendInfoRequest) SetComments(v string) *UpdateUserExtendInfoRequest {
	s.Comments = &v
	return s
}

func (s *UpdateUserExtendInfoRequest) SetJobCode(v string) *UpdateUserExtendInfoRequest {
	s.JobCode = &v
	return s
}

func (s *UpdateUserExtendInfoRequest) SetJobStatusCode(v []*string) *UpdateUserExtendInfoRequest {
	s.JobStatusCode = v
	return s
}

func (s *UpdateUserExtendInfoRequest) SetUserProbCode(v string) *UpdateUserExtendInfoRequest {
	s.UserProbCode = &v
	return s
}

type UpdateUserExtendInfoResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateUserExtendInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserExtendInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserExtendInfoResponse) SetHeaders(v map[string]*string) *UpdateUserExtendInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserExtendInfoResponse) SetStatusCode(v int32) *UpdateUserExtendInfoResponse {
	s.StatusCode = &v
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

func (client *Client) CampusAddRenterMemberWithOptions(request *CampusAddRenterMemberRequest, headers *CampusAddRenterMemberHeaders, runtime *util.RuntimeOptions) (_result *CampusAddRenterMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Extend)) {
		body["extend"] = request.Extend
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		body["mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RenterId)) {
		body["renterId"] = request.RenterId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
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
		Action:      tea.String("CampusAddRenterMember"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/campuses/renters/members"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CampusAddRenterMemberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CampusAddRenterMember(request *CampusAddRenterMemberRequest) (_result *CampusAddRenterMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CampusAddRenterMemberHeaders{}
	_result = &CampusAddRenterMemberResponse{}
	_body, _err := client.CampusAddRenterMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CampusCreateCampusWithOptions(request *CampusCreateCampusRequest, headers *CampusCreateCampusHeaders, runtime *util.RuntimeOptions) (_result *CampusCreateCampusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		body["address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.Area)) {
		body["area"] = request.Area
	}

	if !tea.BoolValue(util.IsUnset(request.BelongProjectGroupId)) {
		body["belongProjectGroupId"] = request.BelongProjectGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.CampusName)) {
		body["campusName"] = request.CampusName
	}

	if !tea.BoolValue(util.IsUnset(request.Capacity)) {
		body["capacity"] = request.Capacity
	}

	if !tea.BoolValue(util.IsUnset(request.CityId)) {
		body["cityId"] = request.CityId
	}

	if !tea.BoolValue(util.IsUnset(request.Country)) {
		body["country"] = request.Country
	}

	if !tea.BoolValue(util.IsUnset(request.CountyId)) {
		body["countyId"] = request.CountyId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorUnionId)) {
		body["creatorUnionId"] = request.CreatorUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Extend)) {
		body["extend"] = request.Extend
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.OrderEndTime)) {
		body["orderEndTime"] = request.OrderEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OrderInfo)) {
		body["orderInfo"] = request.OrderInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OrderStartTime)) {
		body["orderStartTime"] = request.OrderStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.ProvId)) {
		body["provId"] = request.ProvId
	}

	if !tea.BoolValue(util.IsUnset(request.Telephone)) {
		body["telephone"] = request.Telephone
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
		Action:      tea.String("CampusCreateCampus"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/campuses/projects"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CampusCreateCampusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CampusCreateCampus(request *CampusCreateCampusRequest) (_result *CampusCreateCampusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CampusCreateCampusHeaders{}
	_result = &CampusCreateCampusResponse{}
	_body, _err := client.CampusCreateCampusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CampusCreateCampusGroupWithOptions(request *CampusCreateCampusGroupRequest, headers *CampusCreateCampusGroupHeaders, runtime *util.RuntimeOptions) (_result *CampusCreateCampusGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Extend)) {
		body["extend"] = request.Extend
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
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
		Action:      tea.String("CampusCreateCampusGroup"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/campuses/projects/groups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CampusCreateCampusGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CampusCreateCampusGroup(request *CampusCreateCampusGroupRequest) (_result *CampusCreateCampusGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CampusCreateCampusGroupHeaders{}
	_result = &CampusCreateCampusGroupResponse{}
	_body, _err := client.CampusCreateCampusGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CampusCreateRenterWithOptions(request *CampusCreateRenterRequest, headers *CampusCreateRenterHeaders, runtime *util.RuntimeOptions) (_result *CampusCreateRenterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreditCode)) {
		body["creditCode"] = request.CreditCode
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Extend)) {
		body["extend"] = request.Extend
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		body["state"] = request.State
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
		Action:      tea.String("CampusCreateRenter"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/campuses/renters"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CampusCreateRenterResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CampusCreateRenter(request *CampusCreateRenterRequest) (_result *CampusCreateRenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CampusCreateRenterHeaders{}
	_result = &CampusCreateRenterResponse{}
	_body, _err := client.CampusCreateRenterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CampusDelRenterMemberWithOptions(request *CampusDelRenterMemberRequest, headers *CampusDelRenterMemberHeaders, runtime *util.RuntimeOptions) (_result *CampusDelRenterMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RenterId)) {
		query["renterId"] = request.RenterId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
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
		Action:      tea.String("CampusDelRenterMember"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/campuses/renters/members"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CampusDelRenterMemberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CampusDelRenterMember(request *CampusDelRenterMemberRequest) (_result *CampusDelRenterMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CampusDelRenterMemberHeaders{}
	_result = &CampusDelRenterMemberResponse{}
	_body, _err := client.CampusDelRenterMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CampusDeleteCampusGroupWithOptions(request *CampusDeleteCampusGroupRequest, headers *CampusDeleteCampusGroupHeaders, runtime *util.RuntimeOptions) (_result *CampusDeleteCampusGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CampusProjectGroupId)) {
		query["campusProjectGroupId"] = request.CampusProjectGroupId
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
		Action:      tea.String("CampusDeleteCampusGroup"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/campuses/projects/groups"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CampusDeleteCampusGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CampusDeleteCampusGroup(request *CampusDeleteCampusGroupRequest) (_result *CampusDeleteCampusGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CampusDeleteCampusGroupHeaders{}
	_result = &CampusDeleteCampusGroupResponse{}
	_body, _err := client.CampusDeleteCampusGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CampusDeleteRenterWithOptions(request *CampusDeleteRenterRequest, headers *CampusDeleteRenterHeaders, runtime *util.RuntimeOptions) (_result *CampusDeleteRenterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RenterId)) {
		query["renterId"] = request.RenterId
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
		Action:      tea.String("CampusDeleteRenter"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/campuses/renters"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("none"),
	}
	_result = &CampusDeleteRenterResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CampusDeleteRenter(request *CampusDeleteRenterRequest) (_result *CampusDeleteRenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CampusDeleteRenterHeaders{}
	_result = &CampusDeleteRenterResponse{}
	_body, _err := client.CampusDeleteRenterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CampusGetCampusWithOptions(request *CampusGetCampusRequest, headers *CampusGetCampusHeaders, runtime *util.RuntimeOptions) (_result *CampusGetCampusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CampusDeptId)) {
		query["campusDeptId"] = request.CampusDeptId
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
		Action:      tea.String("CampusGetCampus"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/campuses/projectInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CampusGetCampusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CampusGetCampus(request *CampusGetCampusRequest) (_result *CampusGetCampusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CampusGetCampusHeaders{}
	_result = &CampusGetCampusResponse{}
	_body, _err := client.CampusGetCampusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CampusGetCampusGroupWithOptions(request *CampusGetCampusGroupRequest, headers *CampusGetCampusGroupHeaders, runtime *util.RuntimeOptions) (_result *CampusGetCampusGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["groupId"] = request.GroupId
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
		Action:      tea.String("CampusGetCampusGroup"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/campuses/projects/groupInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CampusGetCampusGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CampusGetCampusGroup(request *CampusGetCampusGroupRequest) (_result *CampusGetCampusGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CampusGetCampusGroupHeaders{}
	_result = &CampusGetCampusGroupResponse{}
	_body, _err := client.CampusGetCampusGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CampusGetRenterWithOptions(request *CampusGetRenterRequest, headers *CampusGetRenterHeaders, runtime *util.RuntimeOptions) (_result *CampusGetRenterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RenterId)) {
		query["renterId"] = request.RenterId
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
		Action:      tea.String("CampusGetRenter"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/campuses/renterInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CampusGetRenterResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CampusGetRenter(request *CampusGetRenterRequest) (_result *CampusGetRenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CampusGetRenterHeaders{}
	_result = &CampusGetRenterResponse{}
	_body, _err := client.CampusGetRenterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CampusGetRenterMemberWithOptions(request *CampusGetRenterMemberRequest, headers *CampusGetRenterMemberHeaders, runtime *util.RuntimeOptions) (_result *CampusGetRenterMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RenterId)) {
		query["renterId"] = request.RenterId
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
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
		Action:      tea.String("CampusGetRenterMember"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/campuses/renters/memberInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CampusGetRenterMemberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CampusGetRenterMember(request *CampusGetRenterMemberRequest) (_result *CampusGetRenterMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CampusGetRenterMemberHeaders{}
	_result = &CampusGetRenterMemberResponse{}
	_body, _err := client.CampusGetRenterMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CampusListCampusWithOptions(request *CampusListCampusRequest, headers *CampusListCampusHeaders, runtime *util.RuntimeOptions) (_result *CampusListCampusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupDeptId)) {
		query["groupDeptId"] = request.GroupDeptId
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
		Action:      tea.String("CampusListCampus"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/campuses/projects"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CampusListCampusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CampusListCampus(request *CampusListCampusRequest) (_result *CampusListCampusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CampusListCampusHeaders{}
	_result = &CampusListCampusResponse{}
	_body, _err := client.CampusListCampusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CampusListCampusGroupWithOptions(headers *CampusListCampusGroupHeaders, runtime *util.RuntimeOptions) (_result *CampusListCampusGroupResponse, _err error) {
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
		Action:      tea.String("CampusListCampusGroup"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/campuses/projects/groups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CampusListCampusGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CampusListCampusGroup() (_result *CampusListCampusGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CampusListCampusGroupHeaders{}
	_result = &CampusListCampusGroupResponse{}
	_body, _err := client.CampusListCampusGroupWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CampusListRenterWithOptions(headers *CampusListRenterHeaders, runtime *util.RuntimeOptions) (_result *CampusListRenterResponse, _err error) {
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
		Action:      tea.String("CampusListRenter"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/campuses/renters"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CampusListRenterResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CampusListRenter() (_result *CampusListRenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CampusListRenterHeaders{}
	_result = &CampusListRenterResponse{}
	_body, _err := client.CampusListRenterWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CampusListRenterMembersWithOptions(request *CampusListRenterMembersRequest, headers *CampusListRenterMembersHeaders, runtime *util.RuntimeOptions) (_result *CampusListRenterMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RenterId)) {
		query["renterId"] = request.RenterId
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
		Action:      tea.String("CampusListRenterMembers"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/campuses/renters/members"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CampusListRenterMembersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CampusListRenterMembers(request *CampusListRenterMembersRequest) (_result *CampusListRenterMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CampusListRenterMembersHeaders{}
	_result = &CampusListRenterMembersResponse{}
	_body, _err := client.CampusListRenterMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CampusUpdateCampusWithOptions(request *CampusUpdateCampusRequest, headers *CampusUpdateCampusHeaders, runtime *util.RuntimeOptions) (_result *CampusUpdateCampusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		body["address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.Area)) {
		body["area"] = request.Area
	}

	if !tea.BoolValue(util.IsUnset(request.BelongProjectGroupId)) {
		body["belongProjectGroupId"] = request.BelongProjectGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.CampusDeptId)) {
		body["campusDeptId"] = request.CampusDeptId
	}

	if !tea.BoolValue(util.IsUnset(request.CampusName)) {
		body["campusName"] = request.CampusName
	}

	if !tea.BoolValue(util.IsUnset(request.Capacity)) {
		body["capacity"] = request.Capacity
	}

	if !tea.BoolValue(util.IsUnset(request.CityId)) {
		body["cityId"] = request.CityId
	}

	if !tea.BoolValue(util.IsUnset(request.Country)) {
		body["country"] = request.Country
	}

	if !tea.BoolValue(util.IsUnset(request.CountyId)) {
		body["countyId"] = request.CountyId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Extend)) {
		body["extend"] = request.Extend
	}

	if !tea.BoolValue(util.IsUnset(request.OrderEndTime)) {
		body["orderEndTime"] = request.OrderEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OrderInfo)) {
		body["orderInfo"] = request.OrderInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OrderStartTime)) {
		body["orderStartTime"] = request.OrderStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.ProvId)) {
		body["provId"] = request.ProvId
	}

	if !tea.BoolValue(util.IsUnset(request.Telephone)) {
		body["telephone"] = request.Telephone
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
		Action:      tea.String("CampusUpdateCampus"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/campuses/projects"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CampusUpdateCampusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CampusUpdateCampus(request *CampusUpdateCampusRequest) (_result *CampusUpdateCampusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CampusUpdateCampusHeaders{}
	_result = &CampusUpdateCampusResponse{}
	_body, _err := client.CampusUpdateCampusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CampusUpdateCampusGroupWithOptions(request *CampusUpdateCampusGroupRequest, headers *CampusUpdateCampusGroupHeaders, runtime *util.RuntimeOptions) (_result *CampusUpdateCampusGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CampusProjectGroupId)) {
		body["campusProjectGroupId"] = request.CampusProjectGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Extend)) {
		body["extend"] = request.Extend
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
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
		Action:      tea.String("CampusUpdateCampusGroup"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/campuses/projects/groups"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CampusUpdateCampusGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CampusUpdateCampusGroup(request *CampusUpdateCampusGroupRequest) (_result *CampusUpdateCampusGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CampusUpdateCampusGroupHeaders{}
	_result = &CampusUpdateCampusGroupResponse{}
	_body, _err := client.CampusUpdateCampusGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CampusUpdateRenterWithOptions(request *CampusUpdateRenterRequest, headers *CampusUpdateRenterHeaders, runtime *util.RuntimeOptions) (_result *CampusUpdateRenterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreditCode)) {
		body["creditCode"] = request.CreditCode
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Extend)) {
		body["extend"] = request.Extend
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RenterId)) {
		body["renterId"] = request.RenterId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		body["state"] = request.State
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
		Action:      tea.String("CampusUpdateRenter"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/campuses/renters"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CampusUpdateRenterResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CampusUpdateRenter(request *CampusUpdateRenterRequest) (_result *CampusUpdateRenterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CampusUpdateRenterHeaders{}
	_result = &CampusUpdateRenterResponse{}
	_body, _err := client.CampusUpdateRenterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CampusUpdateRenterMemberWithOptions(request *CampusUpdateRenterMemberRequest, headers *CampusUpdateRenterMemberHeaders, runtime *util.RuntimeOptions) (_result *CampusUpdateRenterMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Extend)) {
		body["extend"] = request.Extend
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RenterId)) {
		body["renterId"] = request.RenterId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		body["unionId"] = request.UnionId
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
		Action:      tea.String("CampusUpdateRenterMember"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/campuses/renters/members"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CampusUpdateRenterMemberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CampusUpdateRenterMember(request *CampusUpdateRenterMemberRequest) (_result *CampusUpdateRenterMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CampusUpdateRenterMemberHeaders{}
	_result = &CampusUpdateRenterMemberResponse{}
	_body, _err := client.CampusUpdateRenterMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CollegeActiveCollegeDeptGroupWithOptions(request *CollegeActiveCollegeDeptGroupRequest, headers *CollegeActiveCollegeDeptGroupHeaders, runtime *util.RuntimeOptions) (_result *CollegeActiveCollegeDeptGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
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
		Action:      tea.String("CollegeActiveCollegeDeptGroup"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/colleges/depts/groups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CollegeActiveCollegeDeptGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CollegeActiveCollegeDeptGroup(request *CollegeActiveCollegeDeptGroupRequest) (_result *CollegeActiveCollegeDeptGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CollegeActiveCollegeDeptGroupHeaders{}
	_result = &CollegeActiveCollegeDeptGroupResponse{}
	_body, _err := client.CollegeActiveCollegeDeptGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CollegeAddCollegeDeptWithOptions(request *CollegeAddCollegeDeptRequest, headers *CollegeAddCollegeDeptHeaders, runtime *util.RuntimeOptions) (_result *CollegeAddCollegeDeptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptName)) {
		query["deptName"] = request.DeptName
	}

	if !tea.BoolValue(util.IsUnset(request.DeptType)) {
		query["deptType"] = request.DeptType
	}

	if !tea.BoolValue(util.IsUnset(request.SortFactor)) {
		query["sortFactor"] = request.SortFactor
	}

	if !tea.BoolValue(util.IsUnset(request.SuperId)) {
		query["superId"] = request.SuperId
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
		Action:      tea.String("CollegeAddCollegeDept"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/colleges/depts"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CollegeAddCollegeDeptResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CollegeAddCollegeDept(request *CollegeAddCollegeDeptRequest) (_result *CollegeAddCollegeDeptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CollegeAddCollegeDeptHeaders{}
	_result = &CollegeAddCollegeDeptResponse{}
	_body, _err := client.CollegeAddCollegeDeptWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CollegeAddManagerWithOptions(request *CollegeAddManagerRequest, headers *CollegeAddManagerHeaders, runtime *util.RuntimeOptions) (_result *CollegeAddManagerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
		Action:      tea.String("CollegeAddManager"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/colleges/members/depts/managers"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CollegeAddManagerResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CollegeAddManager(request *CollegeAddManagerRequest) (_result *CollegeAddManagerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CollegeAddManagerHeaders{}
	_result = &CollegeAddManagerResponse{}
	_body, _err := client.CollegeAddManagerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CollegeAddStudentWithOptions(request *CollegeAddStudentRequest, headers *CollegeAddStudentHeaders, runtime *util.RuntimeOptions) (_result *CollegeAddStudentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.EmpExtension)) {
		body["empExtension"] = request.EmpExtension
	}

	if !tea.BoolValue(util.IsUnset(request.Gender)) {
		body["gender"] = request.Gender
	}

	if !tea.BoolValue(util.IsUnset(request.IdentifyId)) {
		body["identifyId"] = request.IdentifyId
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		body["mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.StartYear)) {
		body["startYear"] = request.StartYear
	}

	if !tea.BoolValue(util.IsUnset(request.StudentName)) {
		body["studentName"] = request.StudentName
	}

	if !tea.BoolValue(util.IsUnset(request.StudentNumber)) {
		body["studentNumber"] = request.StudentNumber
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
		Action:      tea.String("CollegeAddStudent"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/colleges/members/depts/students"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CollegeAddStudentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CollegeAddStudent(request *CollegeAddStudentRequest) (_result *CollegeAddStudentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CollegeAddStudentHeaders{}
	_result = &CollegeAddStudentResponse{}
	_body, _err := client.CollegeAddStudentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CollegeChangeStudentDeptWithOptions(request *CollegeChangeStudentDeptRequest, headers *CollegeChangeStudentDeptHeaders, runtime *util.RuntimeOptions) (_result *CollegeChangeStudentDeptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.NewDeptId)) {
		query["newDeptId"] = request.NewDeptId
	}

	if !tea.BoolValue(util.IsUnset(request.StudentId)) {
		query["studentId"] = request.StudentId
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
		Action:      tea.String("CollegeChangeStudentDept"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/colleges/members/depts/students/move"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CollegeChangeStudentDeptResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CollegeChangeStudentDept(request *CollegeChangeStudentDeptRequest) (_result *CollegeChangeStudentDeptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CollegeChangeStudentDeptHeaders{}
	_result = &CollegeChangeStudentDeptResponse{}
	_body, _err := client.CollegeChangeStudentDeptWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CollegeDeleteCollegeDeptWithOptions(request *CollegeDeleteCollegeDeptRequest, headers *CollegeDeleteCollegeDeptHeaders, runtime *util.RuntimeOptions) (_result *CollegeDeleteCollegeDeptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
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
		Action:      tea.String("CollegeDeleteCollegeDept"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/colleges/depts"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CollegeDeleteCollegeDeptResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CollegeDeleteCollegeDept(request *CollegeDeleteCollegeDeptRequest) (_result *CollegeDeleteCollegeDeptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CollegeDeleteCollegeDeptHeaders{}
	_result = &CollegeDeleteCollegeDeptResponse{}
	_body, _err := client.CollegeDeleteCollegeDeptWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CollegeListCollegeSubDeptWithOptions(request *CollegeListCollegeSubDeptRequest, headers *CollegeListCollegeSubDeptHeaders, runtime *util.RuntimeOptions) (_result *CollegeListCollegeSubDeptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
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
		Action:      tea.String("CollegeListCollegeSubDept"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/colleges/subDepts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CollegeListCollegeSubDeptResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CollegeListCollegeSubDept(request *CollegeListCollegeSubDeptRequest) (_result *CollegeListCollegeSubDeptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CollegeListCollegeSubDeptHeaders{}
	_result = &CollegeListCollegeSubDeptResponse{}
	_body, _err := client.CollegeListCollegeSubDeptWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CollegeListDeptManagerWithOptions(request *CollegeListDeptManagerRequest, headers *CollegeListDeptManagerHeaders, runtime *util.RuntimeOptions) (_result *CollegeListDeptManagerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
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
		Action:      tea.String("CollegeListDeptManager"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/colleges/members/depts/managers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CollegeListDeptManagerResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CollegeListDeptManager(request *CollegeListDeptManagerRequest) (_result *CollegeListDeptManagerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CollegeListDeptManagerHeaders{}
	_result = &CollegeListDeptManagerResponse{}
	_body, _err := client.CollegeListDeptManagerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CollegeListStudentInfoWithOptions(request *CollegeListStudentInfoRequest, headers *CollegeListStudentInfoHeaders, runtime *util.RuntimeOptions) (_result *CollegeListStudentInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.DingStudentStatus)) {
		query["dingStudentStatus"] = request.DingStudentStatus
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
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
		Action:      tea.String("CollegeListStudentInfo"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/colleges/members/depts/students"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CollegeListStudentInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CollegeListStudentInfo(request *CollegeListStudentInfoRequest) (_result *CollegeListStudentInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CollegeListStudentInfoHeaders{}
	_result = &CollegeListStudentInfoResponse{}
	_body, _err := client.CollegeListStudentInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CollegeListUncheckedStudentWithOptions(request *CollegeListUncheckedStudentRequest, headers *CollegeListUncheckedStudentHeaders, runtime *util.RuntimeOptions) (_result *CollegeListUncheckedStudentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
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
		Action:      tea.String("CollegeListUncheckedStudent"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/colleges/members/organizations/unjoinedStudents"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CollegeListUncheckedStudentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CollegeListUncheckedStudent(request *CollegeListUncheckedStudentRequest) (_result *CollegeListUncheckedStudentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CollegeListUncheckedStudentHeaders{}
	_result = &CollegeListUncheckedStudentResponse{}
	_body, _err := client.CollegeListUncheckedStudentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CollegeQueryCollegeDeptGroupInfoWithOptions(request *CollegeQueryCollegeDeptGroupInfoRequest, headers *CollegeQueryCollegeDeptGroupInfoHeaders, runtime *util.RuntimeOptions) (_result *CollegeQueryCollegeDeptGroupInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
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
		Action:      tea.String("CollegeQueryCollegeDeptGroupInfo"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/colleges/depts/groupInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CollegeQueryCollegeDeptGroupInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CollegeQueryCollegeDeptGroupInfo(request *CollegeQueryCollegeDeptGroupInfoRequest) (_result *CollegeQueryCollegeDeptGroupInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CollegeQueryCollegeDeptGroupInfoHeaders{}
	_result = &CollegeQueryCollegeDeptGroupInfoResponse{}
	_body, _err := client.CollegeQueryCollegeDeptGroupInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CollegeQueryCollegeDeptInfoWithOptions(request *CollegeQueryCollegeDeptInfoRequest, headers *CollegeQueryCollegeDeptInfoHeaders, runtime *util.RuntimeOptions) (_result *CollegeQueryCollegeDeptInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
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
		Action:      tea.String("CollegeQueryCollegeDeptInfo"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/colleges/deptInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CollegeQueryCollegeDeptInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CollegeQueryCollegeDeptInfo(request *CollegeQueryCollegeDeptInfoRequest) (_result *CollegeQueryCollegeDeptInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CollegeQueryCollegeDeptInfoHeaders{}
	_result = &CollegeQueryCollegeDeptInfoResponse{}
	_body, _err := client.CollegeQueryCollegeDeptInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CollegeQueryStudentInfoByDeptWithOptions(request *CollegeQueryStudentInfoByDeptRequest, headers *CollegeQueryStudentInfoByDeptHeaders, runtime *util.RuntimeOptions) (_result *CollegeQueryStudentInfoByDeptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.StudentId)) {
		query["studentId"] = request.StudentId
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
		Action:      tea.String("CollegeQueryStudentInfoByDept"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/colleges/members/depts/studentinfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CollegeQueryStudentInfoByDeptResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CollegeQueryStudentInfoByDept(request *CollegeQueryStudentInfoByDeptRequest) (_result *CollegeQueryStudentInfoByDeptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CollegeQueryStudentInfoByDeptHeaders{}
	_result = &CollegeQueryStudentInfoByDeptResponse{}
	_body, _err := client.CollegeQueryStudentInfoByDeptWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CollegeQueryStudentInfoByMobileWithOptions(request *CollegeQueryStudentInfoByMobileRequest, headers *CollegeQueryStudentInfoByMobileHeaders, runtime *util.RuntimeOptions) (_result *CollegeQueryStudentInfoByMobileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["mobile"] = request.Mobile
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
		Action:      tea.String("CollegeQueryStudentInfoByMobile"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/colleges/members/students/mobiles"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CollegeQueryStudentInfoByMobileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CollegeQueryStudentInfoByMobile(request *CollegeQueryStudentInfoByMobileRequest) (_result *CollegeQueryStudentInfoByMobileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CollegeQueryStudentInfoByMobileHeaders{}
	_result = &CollegeQueryStudentInfoByMobileResponse{}
	_body, _err := client.CollegeQueryStudentInfoByMobileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CollegeQueryStudentInfoByStudentIdWithOptions(request *CollegeQueryStudentInfoByStudentIdRequest, headers *CollegeQueryStudentInfoByStudentIdHeaders, runtime *util.RuntimeOptions) (_result *CollegeQueryStudentInfoByStudentIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StudentId)) {
		query["studentId"] = request.StudentId
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
		Action:      tea.String("CollegeQueryStudentInfoByStudentId"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/colleges/members/students"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CollegeQueryStudentInfoByStudentIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CollegeQueryStudentInfoByStudentId(request *CollegeQueryStudentInfoByStudentIdRequest) (_result *CollegeQueryStudentInfoByStudentIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CollegeQueryStudentInfoByStudentIdHeaders{}
	_result = &CollegeQueryStudentInfoByStudentIdResponse{}
	_body, _err := client.CollegeQueryStudentInfoByStudentIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CollegeRemoveManagerWithOptions(request *CollegeRemoveManagerRequest, headers *CollegeRemoveManagerHeaders, runtime *util.RuntimeOptions) (_result *CollegeRemoveManagerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.IsForce)) {
		query["isForce"] = request.IsForce
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
		Action:      tea.String("CollegeRemoveManager"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/colleges/members/managers"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CollegeRemoveManagerResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CollegeRemoveManager(request *CollegeRemoveManagerRequest) (_result *CollegeRemoveManagerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CollegeRemoveManagerHeaders{}
	_result = &CollegeRemoveManagerResponse{}
	_body, _err := client.CollegeRemoveManagerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CollegeRemoveStudentWithOptions(request *CollegeRemoveStudentRequest, headers *CollegeRemoveStudentHeaders, runtime *util.RuntimeOptions) (_result *CollegeRemoveStudentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.StudentId)) {
		query["studentId"] = request.StudentId
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
		Action:      tea.String("CollegeRemoveStudent"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/colleges/members/depts/students"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CollegeRemoveStudentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CollegeRemoveStudent(request *CollegeRemoveStudentRequest) (_result *CollegeRemoveStudentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CollegeRemoveStudentHeaders{}
	_result = &CollegeRemoveStudentResponse{}
	_body, _err := client.CollegeRemoveStudentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CollegeUpdateCollegeDeptWithOptions(request *CollegeUpdateCollegeDeptRequest, headers *CollegeUpdateCollegeDeptHeaders, runtime *util.RuntimeOptions) (_result *CollegeUpdateCollegeDeptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.DeptName)) {
		query["deptName"] = request.DeptName
	}

	if !tea.BoolValue(util.IsUnset(request.SortFactor)) {
		query["sortFactor"] = request.SortFactor
	}

	if !tea.BoolValue(util.IsUnset(request.SuperId)) {
		query["superId"] = request.SuperId
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
		Action:      tea.String("CollegeUpdateCollegeDept"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/colleges/depts"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CollegeUpdateCollegeDeptResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CollegeUpdateCollegeDept(request *CollegeUpdateCollegeDeptRequest) (_result *CollegeUpdateCollegeDeptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CollegeUpdateCollegeDeptHeaders{}
	_result = &CollegeUpdateCollegeDeptResponse{}
	_body, _err := client.CollegeUpdateCollegeDeptWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CollegeUpdateStudentDeptInfoWithOptions(request *CollegeUpdateStudentDeptInfoRequest, headers *CollegeUpdateStudentDeptInfoHeaders, runtime *util.RuntimeOptions) (_result *CollegeUpdateStudentDeptInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.StudentId)) {
		query["studentId"] = request.StudentId
	}

	if !tea.BoolValue(util.IsUnset(request.StudentNumber)) {
		query["studentNumber"] = request.StudentNumber
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
		Action:      tea.String("CollegeUpdateStudentDeptInfo"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/colleges/members/deptInfos"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CollegeUpdateStudentDeptInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CollegeUpdateStudentDeptInfo(request *CollegeUpdateStudentDeptInfoRequest) (_result *CollegeUpdateStudentDeptInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CollegeUpdateStudentDeptInfoHeaders{}
	_result = &CollegeUpdateStudentDeptInfoResponse{}
	_body, _err := client.CollegeUpdateStudentDeptInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CollegeUpdateStudentInfoWithOptions(request *CollegeUpdateStudentInfoRequest, headers *CollegeUpdateStudentInfoHeaders, runtime *util.RuntimeOptions) (_result *CollegeUpdateStudentInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EmpExtension)) {
		body["empExtension"] = request.EmpExtension
	}

	if !tea.BoolValue(util.IsUnset(request.Gender)) {
		body["gender"] = request.Gender
	}

	if !tea.BoolValue(util.IsUnset(request.IdentifyId)) {
		body["identifyId"] = request.IdentifyId
	}

	if !tea.BoolValue(util.IsUnset(request.StartYear)) {
		body["startYear"] = request.StartYear
	}

	if !tea.BoolValue(util.IsUnset(request.StudentId)) {
		body["studentId"] = request.StudentId
	}

	if !tea.BoolValue(util.IsUnset(request.StudentName)) {
		body["studentName"] = request.StudentName
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
		Action:      tea.String("CollegeUpdateStudentInfo"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/colleges/members/depts/students"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CollegeUpdateStudentInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CollegeUpdateStudentInfo(request *CollegeUpdateStudentInfoRequest) (_result *CollegeUpdateStudentInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CollegeUpdateStudentInfoHeaders{}
	_result = &CollegeUpdateStudentInfoResponse{}
	_body, _err := client.CollegeUpdateStudentInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CollegeUpdateStudentMoblieWithOptions(request *CollegeUpdateStudentMoblieRequest, headers *CollegeUpdateStudentMoblieHeaders, runtime *util.RuntimeOptions) (_result *CollegeUpdateStudentMoblieResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsForce)) {
		query["isForce"] = request.IsForce
	}

	if !tea.BoolValue(util.IsUnset(request.NewMobile)) {
		query["newMobile"] = request.NewMobile
	}

	if !tea.BoolValue(util.IsUnset(request.StudentId)) {
		query["studentId"] = request.StudentId
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
		Action:      tea.String("CollegeUpdateStudentMoblie"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/colleges/members/students/mobiles"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CollegeUpdateStudentMoblieResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CollegeUpdateStudentMoblie(request *CollegeUpdateStudentMoblieRequest) (_result *CollegeUpdateStudentMoblieResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CollegeUpdateStudentMoblieHeaders{}
	_result = &CollegeUpdateStudentMoblieResponse{}
	_body, _err := client.CollegeUpdateStudentMoblieWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeContactCreateWithOptions(request *CustomizeContactCreateRequest, headers *CustomizeContactCreateHeaders, runtime *util.RuntimeOptions) (_result *CustomizeContactCreateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ManagerIdList)) {
		body["managerIdList"] = request.ManagerIdList
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["order"] = request.Order
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
		Action:      tea.String("CustomizeContactCreate"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/customizations/contacts"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CustomizeContactCreateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeContactCreate(request *CustomizeContactCreateRequest) (_result *CustomizeContactCreateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CustomizeContactCreateHeaders{}
	_result = &CustomizeContactCreateResponse{}
	_body, _err := client.CustomizeContactCreateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeContactDeleteWithOptions(request *CustomizeContactDeleteRequest, headers *CustomizeContactDeleteHeaders, runtime *util.RuntimeOptions) (_result *CustomizeContactDeleteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["code"] = request.Code
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
		Action:      tea.String("CustomizeContactDelete"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/customizations/contacts"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CustomizeContactDeleteResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeContactDelete(request *CustomizeContactDeleteRequest) (_result *CustomizeContactDeleteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CustomizeContactDeleteHeaders{}
	_result = &CustomizeContactDeleteResponse{}
	_body, _err := client.CustomizeContactDeleteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeContactDeptCreateWithOptions(request *CustomizeContactDeptCreateRequest, headers *CustomizeContactDeptCreateHeaders, runtime *util.RuntimeOptions) (_result *CustomizeContactDeptCreateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerIdList)) {
		body["managerIdList"] = request.ManagerIdList
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.ParentDeptId)) {
		body["parentDeptId"] = request.ParentDeptId
	}

	if !tea.BoolValue(util.IsUnset(request.RefId)) {
		body["refId"] = request.RefId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
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
		Action:      tea.String("CustomizeContactDeptCreate"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/customizations/departments"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CustomizeContactDeptCreateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeContactDeptCreate(request *CustomizeContactDeptCreateRequest) (_result *CustomizeContactDeptCreateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CustomizeContactDeptCreateHeaders{}
	_result = &CustomizeContactDeptCreateResponse{}
	_body, _err := client.CustomizeContactDeptCreateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeContactDeptDeleteWithOptions(request *CustomizeContactDeptDeleteRequest, headers *CustomizeContactDeptDeleteHeaders, runtime *util.RuntimeOptions) (_result *CustomizeContactDeptDeleteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
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
		Action:      tea.String("CustomizeContactDeptDelete"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/customizations/departments"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CustomizeContactDeptDeleteResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeContactDeptDelete(request *CustomizeContactDeptDeleteRequest) (_result *CustomizeContactDeptDeleteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CustomizeContactDeptDeleteHeaders{}
	_result = &CustomizeContactDeptDeleteResponse{}
	_body, _err := client.CustomizeContactDeptDeleteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeContactDeptGroupCreateWithOptions(request *CustomizeContactDeptGroupCreateRequest, headers *CustomizeContactDeptGroupCreateHeaders, runtime *util.RuntimeOptions) (_result *CustomizeContactDeptGroupCreateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
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
		Action:      tea.String("CustomizeContactDeptGroupCreate"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/customizations/departmentGroups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CustomizeContactDeptGroupCreateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeContactDeptGroupCreate(request *CustomizeContactDeptGroupCreateRequest) (_result *CustomizeContactDeptGroupCreateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CustomizeContactDeptGroupCreateHeaders{}
	_result = &CustomizeContactDeptGroupCreateResponse{}
	_body, _err := client.CustomizeContactDeptGroupCreateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeContactDeptInfoWithOptions(request *CustomizeContactDeptInfoRequest, headers *CustomizeContactDeptInfoHeaders, runtime *util.RuntimeOptions) (_result *CustomizeContactDeptInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
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
		Action:      tea.String("CustomizeContactDeptInfo"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/customizations/departments"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CustomizeContactDeptInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeContactDeptInfo(request *CustomizeContactDeptInfoRequest) (_result *CustomizeContactDeptInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CustomizeContactDeptInfoHeaders{}
	_result = &CustomizeContactDeptInfoResponse{}
	_body, _err := client.CustomizeContactDeptInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeContactDeptListWithOptions(request *CustomizeContactDeptListRequest, headers *CustomizeContactDeptListHeaders, runtime *util.RuntimeOptions) (_result *CustomizeContactDeptListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
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
		Action:      tea.String("CustomizeContactDeptList"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/customizations/subsidiaryDepartments"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CustomizeContactDeptListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeContactDeptList(request *CustomizeContactDeptListRequest) (_result *CustomizeContactDeptListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CustomizeContactDeptListHeaders{}
	_result = &CustomizeContactDeptListResponse{}
	_body, _err := client.CustomizeContactDeptListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeContactDeptUpdateWithOptions(request *CustomizeContactDeptUpdateRequest, headers *CustomizeContactDeptUpdateHeaders, runtime *util.RuntimeOptions) (_result *CustomizeContactDeptUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerIdList)) {
		body["managerIdList"] = request.ManagerIdList
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.ParentDeptId)) {
		body["parentDeptId"] = request.ParentDeptId
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
		Action:      tea.String("CustomizeContactDeptUpdate"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/customizations/departments"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CustomizeContactDeptUpdateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeContactDeptUpdate(request *CustomizeContactDeptUpdateRequest) (_result *CustomizeContactDeptUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CustomizeContactDeptUpdateHeaders{}
	_result = &CustomizeContactDeptUpdateResponse{}
	_body, _err := client.CustomizeContactDeptUpdateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeContactEmpAddWithOptions(request *CustomizeContactEmpAddRequest, headers *CustomizeContactEmpAddHeaders, runtime *util.RuntimeOptions) (_result *CustomizeContactEmpAddResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
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
		Action:      tea.String("CustomizeContactEmpAdd"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/customizations/users"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CustomizeContactEmpAddResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeContactEmpAdd(request *CustomizeContactEmpAddRequest) (_result *CustomizeContactEmpAddResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CustomizeContactEmpAddHeaders{}
	_result = &CustomizeContactEmpAddResponse{}
	_body, _err := client.CustomizeContactEmpAddWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeContactEmpDeleteWithOptions(request *CustomizeContactEmpDeleteRequest, headers *CustomizeContactEmpDeleteHeaders, runtime *util.RuntimeOptions) (_result *CustomizeContactEmpDeleteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
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
		Action:      tea.String("CustomizeContactEmpDelete"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/customizations/users/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CustomizeContactEmpDeleteResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeContactEmpDelete(request *CustomizeContactEmpDeleteRequest) (_result *CustomizeContactEmpDeleteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CustomizeContactEmpDeleteHeaders{}
	_result = &CustomizeContactEmpDeleteResponse{}
	_body, _err := client.CustomizeContactEmpDeleteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeContactEmpListWithOptions(request *CustomizeContactEmpListRequest, headers *CustomizeContactEmpListHeaders, runtime *util.RuntimeOptions) (_result *CustomizeContactEmpListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		query["deptId"] = request.DeptId
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
		Action:      tea.String("CustomizeContactEmpList"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/customizations/users"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CustomizeContactEmpListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeContactEmpList(request *CustomizeContactEmpListRequest) (_result *CustomizeContactEmpListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CustomizeContactEmpListHeaders{}
	_result = &CustomizeContactEmpListResponse{}
	_body, _err := client.CustomizeContactEmpListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeContactListWithOptions(headers *CustomizeContactListHeaders, runtime *util.RuntimeOptions) (_result *CustomizeContactListResponse, _err error) {
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
		Action:      tea.String("CustomizeContactList"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/customizations/contacts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CustomizeContactListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeContactList() (_result *CustomizeContactListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CustomizeContactListHeaders{}
	_result = &CustomizeContactListResponse{}
	_body, _err := client.CustomizeContactListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CustomizeContactUpdateWithOptions(request *CustomizeContactUpdateRequest, headers *CustomizeContactUpdateHeaders, runtime *util.RuntimeOptions) (_result *CustomizeContactUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerIdList)) {
		body["managerIdList"] = request.ManagerIdList
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["order"] = request.Order
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
		Action:      tea.String("CustomizeContactUpdate"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/customizations/contacts"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CustomizeContactUpdateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CustomizeContactUpdate(request *CustomizeContactUpdateRequest) (_result *CustomizeContactUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CustomizeContactUpdateHeaders{}
	_result = &CustomizeContactUpdateResponse{}
	_body, _err := client.CustomizeContactUpdateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DIgitalStoreMessagePushWithOptions(tmpReq *DIgitalStoreMessagePushRequest, headers *DIgitalStoreMessagePushHeaders, runtime *util.RuntimeOptions) (_result *DIgitalStoreMessagePushResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DIgitalStoreMessagePushShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.MessageDataList)) {
		request.MessageDataListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MessageDataList, tea.String("messageDataList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MessageDataListShrink)) {
		query["messageDataList"] = request.MessageDataListShrink
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
		Action:      tea.String("DIgitalStoreMessagePush"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/digitalStores/messages/push"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DIgitalStoreMessagePushResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DIgitalStoreMessagePush(request *DIgitalStoreMessagePushRequest) (_result *DIgitalStoreMessagePushResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DIgitalStoreMessagePushHeaders{}
	_result = &DIgitalStoreMessagePushResponse{}
	_body, _err := client.DIgitalStoreMessagePushWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DigitalStoreContactInfoWithOptions(headers *DigitalStoreContactInfoHeaders, runtime *util.RuntimeOptions) (_result *DigitalStoreContactInfoResponse, _err error) {
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
		Action:      tea.String("DigitalStoreContactInfo"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/digitalStores/contactInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DigitalStoreContactInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DigitalStoreContactInfo() (_result *DigitalStoreContactInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DigitalStoreContactInfoHeaders{}
	_result = &DigitalStoreContactInfoResponse{}
	_body, _err := client.DigitalStoreContactInfoWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DigitalStoreConversationsWithOptions(request *DigitalStoreConversationsRequest, headers *DigitalStoreConversationsHeaders, runtime *util.RuntimeOptions) (_result *DigitalStoreConversationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConversationTitle)) {
		query["conversationTitle"] = request.ConversationTitle
	}

	if !tea.BoolValue(util.IsUnset(request.ConversationType)) {
		query["conversationType"] = request.ConversationType
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
		Action:      tea.String("DigitalStoreConversations"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/digitalStores/conversations"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DigitalStoreConversationsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DigitalStoreConversations(request *DigitalStoreConversationsRequest) (_result *DigitalStoreConversationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DigitalStoreConversationsHeaders{}
	_result = &DigitalStoreConversationsResponse{}
	_body, _err := client.DigitalStoreConversationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DigitalStoreGroupInfoWithOptions(request *DigitalStoreGroupInfoRequest, headers *DigitalStoreGroupInfoHeaders, runtime *util.RuntimeOptions) (_result *DigitalStoreGroupInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["groupId"] = request.GroupId
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
		Action:      tea.String("DigitalStoreGroupInfo"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/digitalStores/groupInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DigitalStoreGroupInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DigitalStoreGroupInfo(request *DigitalStoreGroupInfoRequest) (_result *DigitalStoreGroupInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DigitalStoreGroupInfoHeaders{}
	_result = &DigitalStoreGroupInfoResponse{}
	_body, _err := client.DigitalStoreGroupInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DigitalStoreGroupsWithOptions(headers *DigitalStoreGroupsHeaders, runtime *util.RuntimeOptions) (_result *DigitalStoreGroupsResponse, _err error) {
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
		Action:      tea.String("DigitalStoreGroups"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/digitalStores/groups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DigitalStoreGroupsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DigitalStoreGroups() (_result *DigitalStoreGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DigitalStoreGroupsHeaders{}
	_result = &DigitalStoreGroupsResponse{}
	_body, _err := client.DigitalStoreGroupsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DigitalStoreNodeInfoWithOptions(request *DigitalStoreNodeInfoRequest, headers *DigitalStoreNodeInfoHeaders, runtime *util.RuntimeOptions) (_result *DigitalStoreNodeInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["nodeId"] = request.NodeId
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
		Action:      tea.String("DigitalStoreNodeInfo"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/digitalStores/nodeInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DigitalStoreNodeInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DigitalStoreNodeInfo(request *DigitalStoreNodeInfoRequest) (_result *DigitalStoreNodeInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DigitalStoreNodeInfoHeaders{}
	_result = &DigitalStoreNodeInfoResponse{}
	_body, _err := client.DigitalStoreNodeInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DigitalStoreRightsInfoWithOptions(headers *DigitalStoreRightsInfoHeaders, runtime *util.RuntimeOptions) (_result *DigitalStoreRightsInfoResponse, _err error) {
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
		Action:      tea.String("DigitalStoreRightsInfo"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/digitalStores/rightsInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DigitalStoreRightsInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DigitalStoreRightsInfo() (_result *DigitalStoreRightsInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DigitalStoreRightsInfoHeaders{}
	_result = &DigitalStoreRightsInfoResponse{}
	_body, _err := client.DigitalStoreRightsInfoWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DigitalStoreRolesWithOptions(headers *DigitalStoreRolesHeaders, runtime *util.RuntimeOptions) (_result *DigitalStoreRolesResponse, _err error) {
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
		Action:      tea.String("DigitalStoreRoles"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/digitalStores/roles"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DigitalStoreRolesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DigitalStoreRoles() (_result *DigitalStoreRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DigitalStoreRolesHeaders{}
	_result = &DigitalStoreRolesResponse{}
	_body, _err := client.DigitalStoreRolesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DigitalStoreSceneScopeWithOptions(request *DigitalStoreSceneScopeRequest, headers *DigitalStoreSceneScopeHeaders, runtime *util.RuntimeOptions) (_result *DigitalStoreSceneScopeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		query["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneCode)) {
		query["sceneCode"] = request.SceneCode
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
		Action:      tea.String("DigitalStoreSceneScope"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/digitalStores/sceneScopes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DigitalStoreSceneScopeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DigitalStoreSceneScope(request *DigitalStoreSceneScopeRequest) (_result *DigitalStoreSceneScopeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DigitalStoreSceneScopeHeaders{}
	_result = &DigitalStoreSceneScopeResponse{}
	_body, _err := client.DigitalStoreSceneScopeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DigitalStoreStoreInfoWithOptions(request *DigitalStoreStoreInfoRequest, headers *DigitalStoreStoreInfoHeaders, runtime *util.RuntimeOptions) (_result *DigitalStoreStoreInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.StoreId)) {
		query["storeId"] = request.StoreId
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
		Action:      tea.String("DigitalStoreStoreInfo"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/digitalStores/storeInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DigitalStoreStoreInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DigitalStoreStoreInfo(request *DigitalStoreStoreInfoRequest) (_result *DigitalStoreStoreInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DigitalStoreStoreInfoHeaders{}
	_result = &DigitalStoreStoreInfoResponse{}
	_body, _err := client.DigitalStoreStoreInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DigitalStoreSubNodesWithOptions(request *DigitalStoreSubNodesRequest, headers *DigitalStoreSubNodesHeaders, runtime *util.RuntimeOptions) (_result *DigitalStoreSubNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["nodeId"] = request.NodeId
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
		Action:      tea.String("DigitalStoreSubNodes"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/digitalStores/subsidiaryNodes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DigitalStoreSubNodesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DigitalStoreSubNodes(request *DigitalStoreSubNodesRequest) (_result *DigitalStoreSubNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DigitalStoreSubNodesHeaders{}
	_result = &DigitalStoreSubNodesResponse{}
	_body, _err := client.DigitalStoreSubNodesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DigitalStoreUpdateAuthInfoWithOptions(request *DigitalStoreUpdateAuthInfoRequest, headers *DigitalStoreUpdateAuthInfoHeaders, runtime *util.RuntimeOptions) (_result *DigitalStoreUpdateAuthInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UpdateUserList)) {
		body["updateUserList"] = request.UpdateUserList
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
		Action:      tea.String("DigitalStoreUpdateAuthInfo"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/digitalStores/authInfos"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DigitalStoreUpdateAuthInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DigitalStoreUpdateAuthInfo(request *DigitalStoreUpdateAuthInfoRequest) (_result *DigitalStoreUpdateAuthInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DigitalStoreUpdateAuthInfoHeaders{}
	_result = &DigitalStoreUpdateAuthInfoResponse{}
	_body, _err := client.DigitalStoreUpdateAuthInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DigitalStoreUserInfoWithOptions(request *DigitalStoreUserInfoRequest, headers *DigitalStoreUserInfoHeaders, runtime *util.RuntimeOptions) (_result *DigitalStoreUserInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
		Action:      tea.String("DigitalStoreUserInfo"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/digitalStores/userInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DigitalStoreUserInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DigitalStoreUserInfo(request *DigitalStoreUserInfoRequest) (_result *DigitalStoreUserInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DigitalStoreUserInfoHeaders{}
	_result = &DigitalStoreUserInfoResponse{}
	_body, _err := client.DigitalStoreUserInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DigitalStoreUsersWithOptions(request *DigitalStoreUsersRequest, headers *DigitalStoreUsersHeaders, runtime *util.RuntimeOptions) (_result *DigitalStoreUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		query["nodeId"] = request.NodeId
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
		Action:      tea.String("DigitalStoreUsers"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/digitalStores/nodes/users"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DigitalStoreUsersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DigitalStoreUsers(request *DigitalStoreUsersRequest) (_result *DigitalStoreUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DigitalStoreUsersHeaders{}
	_result = &DigitalStoreUsersResponse{}
	_body, _err := client.DigitalStoreUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExternalQueryExternalAppOrgsWithOptions(request *ExternalQueryExternalAppOrgsRequest, headers *ExternalQueryExternalAppOrgsHeaders, runtime *util.RuntimeOptions) (_result *ExternalQueryExternalAppOrgsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExternalType)) {
		query["externalType"] = request.ExternalType
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
		Action:      tea.String("ExternalQueryExternalAppOrgs"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/externals/apps/organizations"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ExternalQueryExternalAppOrgsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExternalQueryExternalAppOrgs(request *ExternalQueryExternalAppOrgsRequest) (_result *ExternalQueryExternalAppOrgsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ExternalQueryExternalAppOrgsHeaders{}
	_result = &ExternalQueryExternalAppOrgsResponse{}
	_body, _err := client.ExternalQueryExternalAppOrgsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExternalQueryExternalBelongMainOrgWithOptions(request *ExternalQueryExternalBelongMainOrgRequest, headers *ExternalQueryExternalBelongMainOrgHeaders, runtime *util.RuntimeOptions) (_result *ExternalQueryExternalBelongMainOrgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExternalType)) {
		query["externalType"] = request.ExternalType
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
		Action:      tea.String("ExternalQueryExternalBelongMainOrg"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/externals/attributions/masterOrganizations"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ExternalQueryExternalBelongMainOrgResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExternalQueryExternalBelongMainOrg(request *ExternalQueryExternalBelongMainOrgRequest) (_result *ExternalQueryExternalBelongMainOrgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ExternalQueryExternalBelongMainOrgHeaders{}
	_result = &ExternalQueryExternalBelongMainOrgResponse{}
	_body, _err := client.ExternalQueryExternalBelongMainOrgWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExternalQueryExternalOrgsWithOptions(request *ExternalQueryExternalOrgsRequest, headers *ExternalQueryExternalOrgsHeaders, runtime *util.RuntimeOptions) (_result *ExternalQueryExternalOrgsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExternalType)) {
		query["externalType"] = request.ExternalType
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
		Action:      tea.String("ExternalQueryExternalOrgs"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/externals/organizations"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ExternalQueryExternalOrgsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExternalQueryExternalOrgs(request *ExternalQueryExternalOrgsRequest) (_result *ExternalQueryExternalOrgsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ExternalQueryExternalOrgsHeaders{}
	_result = &ExternalQueryExternalOrgsResponse{}
	_body, _err := client.ExternalQueryExternalOrgsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HospitalDataCheckWithOptions(request *HospitalDataCheckRequest, headers *HospitalDataCheckHeaders, runtime *util.RuntimeOptions) (_result *HospitalDataCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllDeptCount)) {
		body["allDeptCount"] = request.AllDeptCount
	}

	if !tea.BoolValue(util.IsUnset(request.AllDeptUserCount)) {
		body["allDeptUserCount"] = request.AllDeptUserCount
	}

	if !tea.BoolValue(util.IsUnset(request.AllGroupCount)) {
		body["allGroupCount"] = request.AllGroupCount
	}

	if !tea.BoolValue(util.IsUnset(request.AllGroupUserCount)) {
		body["allGroupUserCount"] = request.AllGroupUserCount
	}

	if !tea.BoolValue(util.IsUnset(request.DeptCount)) {
		body["deptCount"] = request.DeptCount
	}

	if !tea.BoolValue(util.IsUnset(request.DeptUserCount)) {
		body["deptUserCount"] = request.DeptUserCount
	}

	if !tea.BoolValue(util.IsUnset(request.GroupCount)) {
		body["groupCount"] = request.GroupCount
	}

	if !tea.BoolValue(util.IsUnset(request.GroupUserCount)) {
		body["groupUserCount"] = request.GroupUserCount
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
		Action:      tea.String("HospitalDataCheck"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/medicals/datas/check"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &HospitalDataCheckResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HospitalDataCheck(request *HospitalDataCheckRequest) (_result *HospitalDataCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HospitalDataCheckHeaders{}
	_result = &HospitalDataCheckResponse{}
	_body, _err := client.HospitalDataCheckWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndustryManufactureCommonEventWithOptions(request *IndustryManufactureCommonEventRequest, headers *IndustryManufactureCommonEventHeaders, runtime *util.RuntimeOptions) (_result *IndustryManufactureCommonEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		body["action"] = request.Action
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["appKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.BizData)) {
		body["bizData"] = request.BizData
	}

	if !tea.BoolValue(util.IsUnset(request.EventType)) {
		body["eventType"] = request.EventType
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
		Action:      tea.String("IndustryManufactureCommonEvent"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/manufacturing/bases/commons/events"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &IndustryManufactureCommonEventResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IndustryManufactureCommonEvent(request *IndustryManufactureCommonEventRequest) (_result *IndustryManufactureCommonEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IndustryManufactureCommonEventHeaders{}
	_result = &IndustryManufactureCommonEventResponse{}
	_body, _err := client.IndustryManufactureCommonEventWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndustryManufactureCostRecordListGetWithOptions(request *IndustryManufactureCostRecordListGetRequest, headers *IndustryManufactureCostRecordListGetHeaders, runtime *util.RuntimeOptions) (_result *IndustryManufactureCostRecordListGetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppIds)) {
		body["appIds"] = request.AppIds
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["appName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		body["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsvOrgId)) {
		body["isvOrgId"] = request.IsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.MaterialNo)) {
		body["materialNo"] = request.MaterialNo
	}

	if !tea.BoolValue(util.IsUnset(request.MicroappAgentId)) {
		body["microappAgentId"] = request.MicroappAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNo)) {
		body["orderNo"] = request.OrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["orgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductionTaskNo)) {
		body["productionTaskNo"] = request.ProductionTaskNo
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.SuiteKey)) {
		body["suiteKey"] = request.SuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.TokenGrantType)) {
		body["tokenGrantType"] = request.TokenGrantType
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
		Action:      tea.String("IndustryManufactureCostRecordListGet"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/manufactures/materialCostRecords/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &IndustryManufactureCostRecordListGetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IndustryManufactureCostRecordListGet(request *IndustryManufactureCostRecordListGetRequest) (_result *IndustryManufactureCostRecordListGetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IndustryManufactureCostRecordListGetHeaders{}
	_result = &IndustryManufactureCostRecordListGetResponse{}
	_body, _err := client.IndustryManufactureCostRecordListGetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndustryManufactureFeeListGetWithOptions(request *IndustryManufactureFeeListGetRequest, headers *IndustryManufactureFeeListGetHeaders, runtime *util.RuntimeOptions) (_result *IndustryManufactureFeeListGetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppIds)) {
		body["appIds"] = request.AppIds
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["appName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		body["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IsvOrgId)) {
		body["isvOrgId"] = request.IsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.MaterialNo)) {
		body["materialNo"] = request.MaterialNo
	}

	if !tea.BoolValue(util.IsUnset(request.MicroappAgentId)) {
		body["microappAgentId"] = request.MicroappAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["orgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductionTaskNo)) {
		body["productionTaskNo"] = request.ProductionTaskNo
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.SuiteKey)) {
		body["suiteKey"] = request.SuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.TokenGrantType)) {
		body["tokenGrantType"] = request.TokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
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
		Action:      tea.String("IndustryManufactureFeeListGet"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/manufactures/fees/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &IndustryManufactureFeeListGetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IndustryManufactureFeeListGet(request *IndustryManufactureFeeListGetRequest) (_result *IndustryManufactureFeeListGetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IndustryManufactureFeeListGetHeaders{}
	_result = &IndustryManufactureFeeListGetResponse{}
	_body, _err := client.IndustryManufactureFeeListGetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndustryManufactureLabourCostWithOptions(request *IndustryManufactureLabourCostRequest, headers *IndustryManufactureLabourCostHeaders, runtime *util.RuntimeOptions) (_result *IndustryManufactureLabourCostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppIds)) {
		body["appIds"] = request.AppIds
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["appName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		body["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IsvOrgId)) {
		body["isvOrgId"] = request.IsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.MaterialNo)) {
		body["materialNo"] = request.MaterialNo
	}

	if !tea.BoolValue(util.IsUnset(request.MicroappAgentId)) {
		body["microappAgentId"] = request.MicroappAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["orgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessNo)) {
		body["processNo"] = request.ProcessNo
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.SuiteKey)) {
		body["suiteKey"] = request.SuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.TokenGrantType)) {
		body["tokenGrantType"] = request.TokenGrantType
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
		Action:      tea.String("IndustryManufactureLabourCost"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/manufactures/labourCosts/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &IndustryManufactureLabourCostResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IndustryManufactureLabourCost(request *IndustryManufactureLabourCostRequest) (_result *IndustryManufactureLabourCostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IndustryManufactureLabourCostHeaders{}
	_result = &IndustryManufactureLabourCostResponse{}
	_body, _err := client.IndustryManufactureLabourCostWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndustryManufactureMaterialListWithOptions(request *IndustryManufactureMaterialListRequest, headers *IndustryManufactureMaterialListHeaders, runtime *util.RuntimeOptions) (_result *IndustryManufactureMaterialListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppIds)) {
		body["appIds"] = request.AppIds
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["appName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["currentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		body["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsvOrgId)) {
		body["isvOrgId"] = request.IsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.MaterialNo)) {
		body["materialNo"] = request.MaterialNo
	}

	if !tea.BoolValue(util.IsUnset(request.MicroappAgentId)) {
		body["microappAgentId"] = request.MicroappAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["orgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.SuiteKey)) {
		body["suiteKey"] = request.SuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.TokenGrantType)) {
		body["tokenGrantType"] = request.TokenGrantType
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
		Action:      tea.String("IndustryManufactureMaterialList"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/manufactures/materials/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &IndustryManufactureMaterialListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IndustryManufactureMaterialList(request *IndustryManufactureMaterialListRequest) (_result *IndustryManufactureMaterialListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IndustryManufactureMaterialListHeaders{}
	_result = &IndustryManufactureMaterialListResponse{}
	_body, _err := client.IndustryManufactureMaterialListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndustryManufactureMesDispatchTaskWithOptions(request *IndustryManufactureMesDispatchTaskRequest, headers *IndustryManufactureMesDispatchTaskHeaders, runtime *util.RuntimeOptions) (_result *IndustryManufactureMesDispatchTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		body["action"] = request.Action
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["appKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.BaseDataName)) {
		body["baseDataName"] = request.BaseDataName
	}

	if !tea.BoolValue(util.IsUnset(request.DefectsAmount)) {
		body["defectsAmount"] = request.DefectsAmount
	}

	if !tea.BoolValue(util.IsUnset(request.DispatchStaffName)) {
		body["dispatchStaffName"] = request.DispatchStaffName
	}

	if !tea.BoolValue(util.IsUnset(request.DispatchStaffNo)) {
		body["dispatchStaffNo"] = request.DispatchStaffNo
	}

	if !tea.BoolValue(util.IsUnset(request.FineAmount)) {
		body["fineAmount"] = request.FineAmount
	}

	if !tea.BoolValue(util.IsUnset(request.Overdue)) {
		body["overdue"] = request.Overdue
	}

	if !tea.BoolValue(util.IsUnset(request.PlanQuantity)) {
		body["planQuantity"] = request.PlanQuantity
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessName)) {
		body["processName"] = request.ProcessName
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessUuid)) {
		body["processUuid"] = request.ProcessUuid
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["productCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		body["productName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductSpecification)) {
		body["productSpecification"] = request.ProductSpecification
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectCode)) {
		body["projectCode"] = request.ProjectCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["projectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectStatus)) {
		body["projectStatus"] = request.ProjectStatus
	}

	if !tea.BoolValue(util.IsUnset(request.TaskOperators)) {
		body["taskOperators"] = request.TaskOperators
	}

	if !tea.BoolValue(util.IsUnset(request.TaskPlanEndTime)) {
		body["taskPlanEndTime"] = request.TaskPlanEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskPlanStartTime)) {
		body["taskPlanStartTime"] = request.TaskPlanStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskStatus)) {
		body["taskStatus"] = request.TaskStatus
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		body["taskType"] = request.TaskType
	}

	if !tea.BoolValue(util.IsUnset(request.TeamId)) {
		body["teamId"] = request.TeamId
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["uuid"] = request.Uuid
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
		Action:      tea.String("IndustryManufactureMesDispatchTask"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/manufacturings/dispatchTasks/manage"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &IndustryManufactureMesDispatchTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IndustryManufactureMesDispatchTask(request *IndustryManufactureMesDispatchTaskRequest) (_result *IndustryManufactureMesDispatchTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IndustryManufactureMesDispatchTaskHeaders{}
	_result = &IndustryManufactureMesDispatchTaskResponse{}
	_body, _err := client.IndustryManufactureMesDispatchTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndustryManufactureMesMaterialWithOptions(request *IndustryManufactureMesMaterialRequest, headers *IndustryManufactureMesMaterialHeaders, runtime *util.RuntimeOptions) (_result *IndustryManufactureMesMaterialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		body["action"] = request.Action
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["appKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.BaseDataName)) {
		body["baseDataName"] = request.BaseDataName
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendData)) {
		body["extendData"] = request.ExtendData
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["productCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		body["productName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductSpecification)) {
		body["productSpecification"] = request.ProductSpecification
	}

	if !tea.BoolValue(util.IsUnset(request.Prop)) {
		body["prop"] = request.Prop
	}

	if !tea.BoolValue(util.IsUnset(request.Unit)) {
		body["unit"] = request.Unit
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["uuid"] = request.Uuid
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
		Action:      tea.String("IndustryManufactureMesMaterial"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/manufacturings/materials/manage"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &IndustryManufactureMesMaterialResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IndustryManufactureMesMaterial(request *IndustryManufactureMesMaterialRequest) (_result *IndustryManufactureMesMaterialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IndustryManufactureMesMaterialHeaders{}
	_result = &IndustryManufactureMesMaterialResponse{}
	_body, _err := client.IndustryManufactureMesMaterialWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndustryManufactureMesOutPlanWithOptions(request *IndustryManufactureMesOutPlanRequest, headers *IndustryManufactureMesOutPlanHeaders, runtime *util.RuntimeOptions) (_result *IndustryManufactureMesOutPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApprovalStatus)) {
		body["approvalStatus"] = request.ApprovalStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Approver)) {
		body["approver"] = request.Approver
	}

	if !tea.BoolValue(util.IsUnset(request.BaseDataName)) {
		body["baseDataName"] = request.BaseDataName
	}

	if !tea.BoolValue(util.IsUnset(request.OutSourceProjectCode)) {
		body["outSourceProjectCode"] = request.OutSourceProjectCode
	}

	if !tea.BoolValue(util.IsUnset(request.OutSourceTeamId)) {
		body["outSourceTeamId"] = request.OutSourceTeamId
	}

	if !tea.BoolValue(util.IsUnset(request.Price)) {
		body["price"] = request.Price
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessIdentificationCode)) {
		body["processIdentificationCode"] = request.ProcessIdentificationCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessUuids)) {
		body["processUuids"] = request.ProcessUuids
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["productCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		body["productName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductSpecification)) {
		body["productSpecification"] = request.ProductSpecification
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectCode)) {
		body["projectCode"] = request.ProjectCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["projectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SendPlanQuantity)) {
		body["sendPlanQuantity"] = request.SendPlanQuantity
	}

	if !tea.BoolValue(util.IsUnset(request.SupplierCode)) {
		body["supplierCode"] = request.SupplierCode
	}

	if !tea.BoolValue(util.IsUnset(request.SupplierName)) {
		body["supplierName"] = request.SupplierName
	}

	if !tea.BoolValue(util.IsUnset(request.TotalWage)) {
		body["totalWage"] = request.TotalWage
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["uuid"] = request.Uuid
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
		Action:      tea.String("IndustryManufactureMesOutPlan"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/manufacturings/outPlans/manage"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &IndustryManufactureMesOutPlanResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IndustryManufactureMesOutPlan(request *IndustryManufactureMesOutPlanRequest) (_result *IndustryManufactureMesOutPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IndustryManufactureMesOutPlanHeaders{}
	_result = &IndustryManufactureMesOutPlanResponse{}
	_body, _err := client.IndustryManufactureMesOutPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndustryManufactureMesOutputWithOptions(request *IndustryManufactureMesOutputRequest, headers *IndustryManufactureMesOutputHeaders, runtime *util.RuntimeOptions) (_result *IndustryManufactureMesOutputResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		body["action"] = request.Action
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["appKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.ApproveStatus)) {
		body["approveStatus"] = request.ApproveStatus
	}

	if !tea.BoolValue(util.IsUnset(request.BaseDataName)) {
		body["baseDataName"] = request.BaseDataName
	}

	if !tea.BoolValue(util.IsUnset(request.DefectsAmount)) {
		body["defectsAmount"] = request.DefectsAmount
	}

	if !tea.BoolValue(util.IsUnset(request.DefectsReason)) {
		body["defectsReason"] = request.DefectsReason
	}

	if !tea.BoolValue(util.IsUnset(request.FineAmount)) {
		body["fineAmount"] = request.FineAmount
	}

	if !tea.BoolValue(util.IsUnset(request.HasQualityTest)) {
		body["hasQualityTest"] = request.HasQualityTest
	}

	if !tea.BoolValue(util.IsUnset(request.Overdue)) {
		body["overdue"] = request.Overdue
	}

	if !tea.BoolValue(util.IsUnset(request.PlanQuantity)) {
		body["planQuantity"] = request.PlanQuantity
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessName)) {
		body["processName"] = request.ProcessName
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessUuid)) {
		body["processUuid"] = request.ProcessUuid
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["productCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		body["productName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductSpecification)) {
		body["productSpecification"] = request.ProductSpecification
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectCode)) {
		body["projectCode"] = request.ProjectCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["projectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectStatus)) {
		body["projectStatus"] = request.ProjectStatus
	}

	if !tea.BoolValue(util.IsUnset(request.QualityTestStatus)) {
		body["qualityTestStatus"] = request.QualityTestStatus
	}

	if !tea.BoolValue(util.IsUnset(request.TaskPlanEndTime)) {
		body["taskPlanEndTime"] = request.TaskPlanEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskPlanStartTime)) {
		body["taskPlanStartTime"] = request.TaskPlanStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskStatus)) {
		body["taskStatus"] = request.TaskStatus
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		body["taskType"] = request.TaskType
	}

	if !tea.BoolValue(util.IsUnset(request.TaskUuid)) {
		body["taskUuid"] = request.TaskUuid
	}

	if !tea.BoolValue(util.IsUnset(request.TeamId)) {
		body["teamId"] = request.TeamId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["userName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["uuid"] = request.Uuid
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
		Action:      tea.String("IndustryManufactureMesOutput"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/manufacturings/outputs/manage"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &IndustryManufactureMesOutputResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IndustryManufactureMesOutput(request *IndustryManufactureMesOutputRequest) (_result *IndustryManufactureMesOutputResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IndustryManufactureMesOutputHeaders{}
	_result = &IndustryManufactureMesOutputResponse{}
	_body, _err := client.IndustryManufactureMesOutputWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndustryManufactureMesProcessWithOptions(request *IndustryManufactureMesProcessRequest, headers *IndustryManufactureMesProcessHeaders, runtime *util.RuntimeOptions) (_result *IndustryManufactureMesProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		body["action"] = request.Action
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["appKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.BaseDataName)) {
		body["baseDataName"] = request.BaseDataName
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendData)) {
		body["extendData"] = request.ExtendData
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NeedDispatch)) {
		body["needDispatch"] = request.NeedDispatch
	}

	if !tea.BoolValue(util.IsUnset(request.NeedQualityTest)) {
		body["needQualityTest"] = request.NeedQualityTest
	}

	if !tea.BoolValue(util.IsUnset(request.No)) {
		body["no"] = request.No
	}

	if !tea.BoolValue(util.IsUnset(request.Price)) {
		body["price"] = request.Price
	}

	if !tea.BoolValue(util.IsUnset(request.Prop)) {
		body["prop"] = request.Prop
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Sop)) {
		body["sop"] = request.Sop
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["uuid"] = request.Uuid
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
		Action:      tea.String("IndustryManufactureMesProcess"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/manufacturings/processes/manage"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &IndustryManufactureMesProcessResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IndustryManufactureMesProcess(request *IndustryManufactureMesProcessRequest) (_result *IndustryManufactureMesProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IndustryManufactureMesProcessHeaders{}
	_result = &IndustryManufactureMesProcessResponse{}
	_body, _err := client.IndustryManufactureMesProcessWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndustryManufactureMesProductionPlanWithOptions(request *IndustryManufactureMesProductionPlanRequest, headers *IndustryManufactureMesProductionPlanHeaders, runtime *util.RuntimeOptions) (_result *IndustryManufactureMesProductionPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		body["action"] = request.Action
	}

	if !tea.BoolValue(util.IsUnset(request.ActualEndTime)) {
		body["actualEndTime"] = request.ActualEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ActualStartTime)) {
		body["actualStartTime"] = request.ActualStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["appKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.BaseDataName)) {
		body["baseDataName"] = request.BaseDataName
	}

	if !tea.BoolValue(util.IsUnset(request.BomUuid)) {
		body["bomUuid"] = request.BomUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Events)) {
		body["events"] = request.Events
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendData)) {
		body["extendData"] = request.ExtendData
	}

	if !tea.BoolValue(util.IsUnset(request.No)) {
		body["no"] = request.No
	}

	if !tea.BoolValue(util.IsUnset(request.Overdue)) {
		body["overdue"] = request.Overdue
	}

	if !tea.BoolValue(util.IsUnset(request.PlanEndTime)) {
		body["planEndTime"] = request.PlanEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PlanQuantity)) {
		body["planQuantity"] = request.PlanQuantity
	}

	if !tea.BoolValue(util.IsUnset(request.PlanStartTime)) {
		body["planStartTime"] = request.PlanStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessUuids)) {
		body["processUuids"] = request.ProcessUuids
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["productCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		body["productName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductSpecification)) {
		body["productSpecification"] = request.ProductSpecification
	}

	if !tea.BoolValue(util.IsUnset(request.QualifiedQuantity)) {
		body["qualifiedQuantity"] = request.QualifiedQuantity
	}

	if !tea.BoolValue(util.IsUnset(request.SellOrderNo)) {
		body["sellOrderNo"] = request.SellOrderNo
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TeamList)) {
		body["teamList"] = request.TeamList
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Unit)) {
		body["unit"] = request.Unit
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["uuid"] = request.Uuid
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
		Action:      tea.String("IndustryManufactureMesProductionPlan"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/manufacturings/productionPlans/manage"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &IndustryManufactureMesProductionPlanResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IndustryManufactureMesProductionPlan(request *IndustryManufactureMesProductionPlanRequest) (_result *IndustryManufactureMesProductionPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IndustryManufactureMesProductionPlanHeaders{}
	_result = &IndustryManufactureMesProductionPlanResponse{}
	_body, _err := client.IndustryManufactureMesProductionPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndustryManufactureMesSubCooperationTeamWithOptions(request *IndustryManufactureMesSubCooperationTeamRequest, headers *IndustryManufactureMesSubCooperationTeamHeaders, runtime *util.RuntimeOptions) (_result *IndustryManufactureMesSubCooperationTeamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		body["action"] = request.Action
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["appKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.BaseDataName)) {
		body["baseDataName"] = request.BaseDataName
	}

	if !tea.BoolValue(util.IsUnset(request.Events)) {
		body["events"] = request.Events
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendData)) {
		body["extendData"] = request.ExtendData
	}

	if !tea.BoolValue(util.IsUnset(request.GroupPlugins)) {
		body["groupPlugins"] = request.GroupPlugins
	}

	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		body["groupType"] = request.GroupType
	}

	if !tea.BoolValue(util.IsUnset(request.Leaders)) {
		body["leaders"] = request.Leaders
	}

	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OutCorpId)) {
		body["outCorpId"] = request.OutCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessIds)) {
		body["processIds"] = request.ProcessIds
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["uuid"] = request.Uuid
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
		Action:      tea.String("IndustryManufactureMesSubCooperationTeam"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/manufacturings/outTeams/manage"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &IndustryManufactureMesSubCooperationTeamResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IndustryManufactureMesSubCooperationTeam(request *IndustryManufactureMesSubCooperationTeamRequest) (_result *IndustryManufactureMesSubCooperationTeamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IndustryManufactureMesSubCooperationTeamHeaders{}
	_result = &IndustryManufactureMesSubCooperationTeamResponse{}
	_body, _err := client.IndustryManufactureMesSubCooperationTeamWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndustryManufactureMesTeamMgmtWithOptions(request *IndustryManufactureMesTeamMgmtRequest, headers *IndustryManufactureMesTeamMgmtHeaders, runtime *util.RuntimeOptions) (_result *IndustryManufactureMesTeamMgmtResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		body["action"] = request.Action
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["appKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.BaseDataName)) {
		body["baseDataName"] = request.BaseDataName
	}

	if !tea.BoolValue(util.IsUnset(request.Events)) {
		body["events"] = request.Events
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendData)) {
		body["extendData"] = request.ExtendData
	}

	if !tea.BoolValue(util.IsUnset(request.GroupConfig)) {
		body["groupConfig"] = request.GroupConfig
	}

	if !tea.BoolValue(util.IsUnset(request.GroupPlugins)) {
		body["groupPlugins"] = request.GroupPlugins
	}

	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		body["groupType"] = request.GroupType
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Leaders)) {
		body["leaders"] = request.Leaders
	}

	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessIds)) {
		body["processIds"] = request.ProcessIds
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		body["tagKey"] = request.TagKey
	}

	if !tea.BoolValue(util.IsUnset(request.TagValues)) {
		body["tagValues"] = request.TagValues
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
		Action:      tea.String("IndustryManufactureMesTeamMgmt"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/manufacturing/base/data/team"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &IndustryManufactureMesTeamMgmtResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IndustryManufactureMesTeamMgmt(request *IndustryManufactureMesTeamMgmtRequest) (_result *IndustryManufactureMesTeamMgmtResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IndustryManufactureMesTeamMgmtHeaders{}
	_result = &IndustryManufactureMesTeamMgmtResponse{}
	_body, _err := client.IndustryManufactureMesTeamMgmtWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IndustryMmanufactureMaterialCostGetWithOptions(request *IndustryMmanufactureMaterialCostGetRequest, headers *IndustryMmanufactureMaterialCostGetHeaders, runtime *util.RuntimeOptions) (_result *IndustryMmanufactureMaterialCostGetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppIds)) {
		body["appIds"] = request.AppIds
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["appName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		body["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsvOrgId)) {
		body["isvOrgId"] = request.IsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.MaterialNo)) {
		body["materialNo"] = request.MaterialNo
	}

	if !tea.BoolValue(util.IsUnset(request.MicroappAgentId)) {
		body["microappAgentId"] = request.MicroappAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["orgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.SuiteKey)) {
		body["suiteKey"] = request.SuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.TokenGrantType)) {
		body["tokenGrantType"] = request.TokenGrantType
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
		Action:      tea.String("IndustryMmanufactureMaterialCostGet"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/manufactures/base/materialCosts/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &IndustryMmanufactureMaterialCostGetResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IndustryMmanufactureMaterialCostGet(request *IndustryMmanufactureMaterialCostGetRequest) (_result *IndustryMmanufactureMaterialCostGetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &IndustryMmanufactureMaterialCostGetHeaders{}
	_result = &IndustryMmanufactureMaterialCostGetResponse{}
	_body, _err := client.IndustryMmanufactureMaterialCostGetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushDingMessageWithOptions(request *PushDingMessageRequest, headers *PushDingMessageHeaders, runtime *util.RuntimeOptions) (_result *PushDingMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.MessageType)) {
		body["messageType"] = request.MessageType
	}

	if !tea.BoolValue(util.IsUnset(request.MessageUrl)) {
		body["messageUrl"] = request.MessageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.PictureUrl)) {
		body["pictureUrl"] = request.PictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SingleTitle)) {
		body["singleTitle"] = request.SingleTitle
	}

	if !tea.BoolValue(util.IsUnset(request.SingleUrl)) {
		body["singleUrl"] = request.SingleUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
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
		Action:      tea.String("PushDingMessage"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/works/notice"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PushDingMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushDingMessage(request *PushDingMessageRequest) (_result *PushDingMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PushDingMessageHeaders{}
	_result = &PushDingMessageResponse{}
	_body, _err := client.PushDingMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllDepartmentWithOptions(request *QueryAllDepartmentRequest, headers *QueryAllDepartmentHeaders, runtime *util.RuntimeOptions) (_result *QueryAllDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
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
		Action:      tea.String("QueryAllDepartment"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/medicals/departments"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAllDepartmentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllDepartment(request *QueryAllDepartmentRequest) (_result *QueryAllDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllDepartmentHeaders{}
	_result = &QueryAllDepartmentResponse{}
	_body, _err := client.QueryAllDepartmentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllDoctorsWithOptions(request *QueryAllDoctorsRequest, headers *QueryAllDoctorsHeaders, runtime *util.RuntimeOptions) (_result *QueryAllDoctorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MonthMark)) {
		query["monthMark"] = request.MonthMark
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
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
		Action:      tea.String("QueryAllDoctors"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/medicals/doctors"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAllDoctorsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllDoctors(request *QueryAllDoctorsRequest) (_result *QueryAllDoctorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllDoctorsHeaders{}
	_result = &QueryAllDoctorsResponse{}
	_body, _err := client.QueryAllDoctorsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllGroupWithOptions(request *QueryAllGroupRequest, headers *QueryAllGroupHeaders, runtime *util.RuntimeOptions) (_result *QueryAllGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
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
		Action:      tea.String("QueryAllGroup"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/medicals/groups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAllGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllGroup(request *QueryAllGroupRequest) (_result *QueryAllGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllGroupHeaders{}
	_result = &QueryAllGroupResponse{}
	_body, _err := client.QueryAllGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllGroupsInDeptWithOptions(deptId *string, request *QueryAllGroupsInDeptRequest, headers *QueryAllGroupsInDeptHeaders, runtime *util.RuntimeOptions) (_result *QueryAllGroupsInDeptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
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
		Action:      tea.String("QueryAllGroupsInDept"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/medicals/departments/" + tea.StringValue(deptId) + "/groups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAllGroupsInDeptResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllGroupsInDept(deptId *string, request *QueryAllGroupsInDeptRequest) (_result *QueryAllGroupsInDeptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllGroupsInDeptHeaders{}
	_result = &QueryAllGroupsInDeptResponse{}
	_body, _err := client.QueryAllGroupsInDeptWithOptions(deptId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllMemberByDeptWithOptions(deptId *string, request *QueryAllMemberByDeptRequest, headers *QueryAllMemberByDeptHeaders, runtime *util.RuntimeOptions) (_result *QueryAllMemberByDeptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MonthMark)) {
		query["monthMark"] = request.MonthMark
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
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
		Action:      tea.String("QueryAllMemberByDept"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/medicals/departments/" + tea.StringValue(deptId) + "/members"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAllMemberByDeptResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllMemberByDept(deptId *string, request *QueryAllMemberByDeptRequest) (_result *QueryAllMemberByDeptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllMemberByDeptHeaders{}
	_result = &QueryAllMemberByDeptResponse{}
	_body, _err := client.QueryAllMemberByDeptWithOptions(deptId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllMemberByGroupWithOptions(groupId *string, request *QueryAllMemberByGroupRequest, headers *QueryAllMemberByGroupHeaders, runtime *util.RuntimeOptions) (_result *QueryAllMemberByGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MonthMark)) {
		query["monthMark"] = request.MonthMark
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
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
		Action:      tea.String("QueryAllMemberByGroup"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/medicals/groups/" + tea.StringValue(groupId) + "/members"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAllMemberByGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllMemberByGroup(groupId *string, request *QueryAllMemberByGroupRequest) (_result *QueryAllMemberByGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllMemberByGroupHeaders{}
	_result = &QueryAllMemberByGroupResponse{}
	_body, _err := client.QueryAllMemberByGroupWithOptions(groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBizOptLogWithOptions(request *QueryBizOptLogRequest, headers *QueryBizOptLogHeaders, runtime *util.RuntimeOptions) (_result *QueryBizOptLogResponse, _err error) {
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
		Action:      tea.String("QueryBizOptLog"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/medicals/bizOptLogs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBizOptLogResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBizOptLog(request *QueryBizOptLogRequest) (_result *QueryBizOptLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryBizOptLogHeaders{}
	_result = &QueryBizOptLogResponse{}
	_body, _err := client.QueryBizOptLogWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDepartmentExtendInfoWithOptions(request *QueryDepartmentExtendInfoRequest, headers *QueryDepartmentExtendInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryDepartmentExtendInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptCode)) {
		query["deptCode"] = request.DeptCode
	}

	if !tea.BoolValue(util.IsUnset(request.PropCode)) {
		query["propCode"] = request.PropCode
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
		Action:      tea.String("QueryDepartmentExtendInfo"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/medicals/departments/extensions/infos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDepartmentExtendInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDepartmentExtendInfo(request *QueryDepartmentExtendInfoRequest) (_result *QueryDepartmentExtendInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDepartmentExtendInfoHeaders{}
	_result = &QueryDepartmentExtendInfoResponse{}
	_body, _err := client.QueryDepartmentExtendInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDepartmentInfoWithOptions(deptId *string, headers *QueryDepartmentInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryDepartmentInfoResponse, _err error) {
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
		Action:      tea.String("QueryDepartmentInfo"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/medicals/departments/" + tea.StringValue(deptId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDepartmentInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDepartmentInfo(deptId *string) (_result *QueryDepartmentInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDepartmentInfoHeaders{}
	_result = &QueryDepartmentInfoResponse{}
	_body, _err := client.QueryDepartmentInfoWithOptions(deptId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDoctorDetailsByJobNumberWithOptions(jobNumber *string, request *QueryDoctorDetailsByJobNumberRequest, headers *QueryDoctorDetailsByJobNumberHeaders, runtime *util.RuntimeOptions) (_result *QueryDoctorDetailsByJobNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MonthMark)) {
		query["monthMark"] = request.MonthMark
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
		Action:      tea.String("QueryDoctorDetailsByJobNumber"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/medicals/doctors/" + tea.StringValue(jobNumber)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDoctorDetailsByJobNumberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDoctorDetailsByJobNumber(jobNumber *string, request *QueryDoctorDetailsByJobNumberRequest) (_result *QueryDoctorDetailsByJobNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryDoctorDetailsByJobNumberHeaders{}
	_result = &QueryDoctorDetailsByJobNumberResponse{}
	_body, _err := client.QueryDoctorDetailsByJobNumberWithOptions(jobNumber, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryGroupInfoWithOptions(groupId *string, headers *QueryGroupInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryGroupInfoResponse, _err error) {
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
		Action:      tea.String("QueryGroupInfo"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/medicals/groups/" + tea.StringValue(groupId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGroupInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryGroupInfo(groupId *string) (_result *QueryGroupInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryGroupInfoHeaders{}
	_result = &QueryGroupInfoResponse{}
	_body, _err := client.QueryGroupInfoWithOptions(groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryHospitalDistrictInfoWithOptions(request *QueryHospitalDistrictInfoRequest, headers *QueryHospitalDistrictInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryHospitalDistrictInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
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
		Action:      tea.String("QueryHospitalDistrictInfo"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/medicals/districts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryHospitalDistrictInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryHospitalDistrictInfo(request *QueryHospitalDistrictInfoRequest) (_result *QueryHospitalDistrictInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryHospitalDistrictInfoHeaders{}
	_result = &QueryHospitalDistrictInfoResponse{}
	_body, _err := client.QueryHospitalDistrictInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryHospitalRoleUserInfoWithOptions(request *QueryHospitalRoleUserInfoRequest, headers *QueryHospitalRoleUserInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryHospitalRoleUserInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
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
		Action:      tea.String("QueryHospitalRoleUserInfo"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/medicals/roles/userInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryHospitalRoleUserInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryHospitalRoleUserInfo(request *QueryHospitalRoleUserInfoRequest) (_result *QueryHospitalRoleUserInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryHospitalRoleUserInfoHeaders{}
	_result = &QueryHospitalRoleUserInfoResponse{}
	_body, _err := client.QueryHospitalRoleUserInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryHospitalRolesWithOptions(headers *QueryHospitalRolesHeaders, runtime *util.RuntimeOptions) (_result *QueryHospitalRolesResponse, _err error) {
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
		Action:      tea.String("QueryHospitalRoles"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/medicals/roles"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryHospitalRolesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryHospitalRoles() (_result *QueryHospitalRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryHospitalRolesHeaders{}
	_result = &QueryHospitalRolesResponse{}
	_body, _err := client.QueryHospitalRolesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryJobCodeDictionaryWithOptions(headers *QueryJobCodeDictionaryHeaders, runtime *util.RuntimeOptions) (_result *QueryJobCodeDictionaryResponse, _err error) {
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
		Action:      tea.String("QueryJobCodeDictionary"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/medicals/jobCodes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryJobCodeDictionaryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryJobCodeDictionary() (_result *QueryJobCodeDictionaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryJobCodeDictionaryHeaders{}
	_result = &QueryJobCodeDictionaryResponse{}
	_body, _err := client.QueryJobCodeDictionaryWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryJobStatusCodeDictionaryWithOptions(headers *QueryJobStatusCodeDictionaryHeaders, runtime *util.RuntimeOptions) (_result *QueryJobStatusCodeDictionaryResponse, _err error) {
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
		Action:      tea.String("QueryJobStatusCodeDictionary"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/medicals/jobStatusCodes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryJobStatusCodeDictionaryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryJobStatusCodeDictionary() (_result *QueryJobStatusCodeDictionaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryJobStatusCodeDictionaryHeaders{}
	_result = &QueryJobStatusCodeDictionaryResponse{}
	_body, _err := client.QueryJobStatusCodeDictionaryWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMedicalEventsWithOptions(headers *QueryMedicalEventsHeaders, runtime *util.RuntimeOptions) (_result *QueryMedicalEventsResponse, _err error) {
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
		Action:      tea.String("QueryMedicalEvents"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/medicals/events"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMedicalEventsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMedicalEvents() (_result *QueryMedicalEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryMedicalEventsHeaders{}
	_result = &QueryMedicalEventsResponse{}
	_body, _err := client.QueryMedicalEventsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserCredentialsWithOptions(request *QueryUserCredentialsRequest, headers *QueryUserCredentialsHeaders, runtime *util.RuntimeOptions) (_result *QueryUserCredentialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
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
		Action:      tea.String("QueryUserCredentials"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/medicals/users/credentials/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserCredentialsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserCredentials(request *QueryUserCredentialsRequest) (_result *QueryUserCredentialsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserCredentialsHeaders{}
	_result = &QueryUserCredentialsResponse{}
	_body, _err := client.QueryUserCredentialsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserExtInfoWithOptions(userId *string, headers *QueryUserExtInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryUserExtInfoResponse, _err error) {
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
		Action:      tea.String("QueryUserExtInfo"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/medicals/users/" + tea.StringValue(userId) + "/extInfos"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserExtInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserExtInfo(userId *string) (_result *QueryUserExtInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserExtInfoHeaders{}
	_result = &QueryUserExtInfoResponse{}
	_body, _err := client.QueryUserExtInfoWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserExtendValuesWithOptions(request *QueryUserExtendValuesRequest, headers *QueryUserExtendValuesHeaders, runtime *util.RuntimeOptions) (_result *QueryUserExtendValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserExtendKey)) {
		body["userExtendKey"] = request.UserExtendKey
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["userIds"] = request.UserIds
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
		Action:      tea.String("QueryUserExtendValues"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/medicals/users/extends/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserExtendValuesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserExtendValues(request *QueryUserExtendValuesRequest) (_result *QueryUserExtendValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserExtendValuesHeaders{}
	_result = &QueryUserExtendValuesResponse{}
	_body, _err := client.QueryUserExtendValuesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserInfoWithOptions(userId *string, request *QueryUserInfoRequest, headers *QueryUserInfoHeaders, runtime *util.RuntimeOptions) (_result *QueryUserInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MonthMark)) {
		query["monthMark"] = request.MonthMark
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
		Action:      tea.String("QueryUserInfo"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/medicals/users/" + tea.StringValue(userId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserInfo(userId *string, request *QueryUserInfoRequest) (_result *QueryUserInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserInfoHeaders{}
	_result = &QueryUserInfoResponse{}
	_body, _err := client.QueryUserInfoWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserProbCodeDictionaryWithOptions(headers *QueryUserProbCodeDictionaryHeaders, runtime *util.RuntimeOptions) (_result *QueryUserProbCodeDictionaryResponse, _err error) {
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
		Action:      tea.String("QueryUserProbCodeDictionary"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/medicals/userProbCodes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserProbCodeDictionaryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserProbCodeDictionary() (_result *QueryUserProbCodeDictionaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserProbCodeDictionaryHeaders{}
	_result = &QueryUserProbCodeDictionaryResponse{}
	_body, _err := client.QueryUserProbCodeDictionaryWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserRolesWithOptions(userId *string, headers *QueryUserRolesHeaders, runtime *util.RuntimeOptions) (_result *QueryUserRolesResponse, _err error) {
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
		Action:      tea.String("QueryUserRoles"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/medicals/users/" + tea.StringValue(userId) + "/roles"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUserRolesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserRoles(userId *string) (_result *QueryUserRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryUserRolesHeaders{}
	_result = &QueryUserRolesResponse{}
	_body, _err := client.QueryUserRolesWithOptions(userId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveUserExtendValuesWithOptions(userId *string, request *SaveUserExtendValuesRequest, headers *SaveUserExtendValuesHeaders, runtime *util.RuntimeOptions) (_result *SaveUserExtendValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserDisplayName)) {
		query["userDisplayName"] = request.UserDisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.UserExtendKey)) {
		query["userExtendKey"] = request.UserExtendKey
	}

	if !tea.BoolValue(util.IsUnset(request.UserExtendValue)) {
		query["userExtendValue"] = request.UserExtendValue
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
		Action:      tea.String("SaveUserExtendValues"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/medicals/users/" + tea.StringValue(userId) + "/extends"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveUserExtendValuesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveUserExtendValues(userId *string, request *SaveUserExtendValuesRequest) (_result *SaveUserExtendValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveUserExtendValuesHeaders{}
	_result = &SaveUserExtendValuesResponse{}
	_body, _err := client.SaveUserExtendValuesWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SupplyAddDeptWithOptions(request *SupplyAddDeptRequest, headers *SupplyAddDeptHeaders, runtime *util.RuntimeOptions) (_result *SupplyAddDeptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptName)) {
		query["deptName"] = request.DeptName
	}

	if !tea.BoolValue(util.IsUnset(request.PartnerNumber)) {
		query["partnerNumber"] = request.PartnerNumber
	}

	if !tea.BoolValue(util.IsUnset(request.SuperDeptId)) {
		query["superDeptId"] = request.SuperDeptId
	}

	if !tea.BoolValue(util.IsUnset(request.SupplyDeptType)) {
		query["supplyDeptType"] = request.SupplyDeptType
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
		Action:      tea.String("SupplyAddDept"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/supplyChains/departments"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SupplyAddDeptResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SupplyAddDept(request *SupplyAddDeptRequest) (_result *SupplyAddDeptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SupplyAddDeptHeaders{}
	_result = &SupplyAddDeptResponse{}
	_body, _err := client.SupplyAddDeptWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SupplyAddMemberWithOptions(request *SupplyAddMemberRequest, headers *SupplyAddMemberHeaders, runtime *util.RuntimeOptions) (_result *SupplyAddMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsPartnerManager)) {
		query["isPartnerManager"] = request.IsPartnerManager
	}

	if !tea.BoolValue(util.IsUnset(request.MemberMobile)) {
		query["memberMobile"] = request.MemberMobile
	}

	if !tea.BoolValue(util.IsUnset(request.MemberName)) {
		query["memberName"] = request.MemberName
	}

	if !tea.BoolValue(util.IsUnset(request.MemberWorkNumber)) {
		query["memberWorkNumber"] = request.MemberWorkNumber
	}

	if !tea.BoolValue(util.IsUnset(request.SupplyDeptId)) {
		query["supplyDeptId"] = request.SupplyDeptId
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
		Action:      tea.String("SupplyAddMember"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/supplyChains/members"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SupplyAddMemberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SupplyAddMember(request *SupplyAddMemberRequest) (_result *SupplyAddMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SupplyAddMemberHeaders{}
	_result = &SupplyAddMemberResponse{}
	_body, _err := client.SupplyAddMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SupplyGetMemberWithOptions(request *SupplyGetMemberRequest, headers *SupplyGetMemberHeaders, runtime *util.RuntimeOptions) (_result *SupplyGetMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UnionId)) {
		query["unionId"] = request.UnionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
		Action:      tea.String("SupplyGetMember"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/supplyChains/members"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SupplyGetMemberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SupplyGetMember(request *SupplyGetMemberRequest) (_result *SupplyGetMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SupplyGetMemberHeaders{}
	_result = &SupplyGetMemberResponse{}
	_body, _err := client.SupplyGetMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SupplyListDeptMembersWithOptions(request *SupplyListDeptMembersRequest, headers *SupplyListDeptMembersHeaders, runtime *util.RuntimeOptions) (_result *SupplyListDeptMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SupplyDeptId)) {
		query["supplyDeptId"] = request.SupplyDeptId
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
		Action:      tea.String("SupplyListDeptMembers"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/supplyChains/departments/members"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SupplyListDeptMembersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SupplyListDeptMembers(request *SupplyListDeptMembersRequest) (_result *SupplyListDeptMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SupplyListDeptMembersHeaders{}
	_result = &SupplyListDeptMembersResponse{}
	_body, _err := client.SupplyListDeptMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUserExtendInfoWithOptions(userId *string, request *UpdateUserExtendInfoRequest, headers *UpdateUserExtendInfoHeaders, runtime *util.RuntimeOptions) (_result *UpdateUserExtendInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Comments)) {
		body["comments"] = request.Comments
	}

	if !tea.BoolValue(util.IsUnset(request.JobCode)) {
		body["jobCode"] = request.JobCode
	}

	if !tea.BoolValue(util.IsUnset(request.JobStatusCode)) {
		body["jobStatusCode"] = request.JobStatusCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserProbCode)) {
		body["userProbCode"] = request.UserProbCode
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
		Action:      tea.String("UpdateUserExtendInfo"),
		Version:     tea.String("industry_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/industry/medicals/users/" + tea.StringValue(userId) + "/extInfos"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserExtendInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUserExtendInfo(userId *string, request *UpdateUserExtendInfoRequest) (_result *UpdateUserExtendInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateUserExtendInfoHeaders{}
	_result = &UpdateUserExtendInfoResponse{}
	_body, _err := client.UpdateUserExtendInfoWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
