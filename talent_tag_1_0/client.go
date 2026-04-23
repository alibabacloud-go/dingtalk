// This file is auto-generated, don't edit it. Thanks.
package talent_tag_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type TalentV2AddCustomTagHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TalentV2AddCustomTagHeaders) String() string {
	return tea.Prettify(s)
}

func (s TalentV2AddCustomTagHeaders) GoString() string {
	return s.String()
}

func (s *TalentV2AddCustomTagHeaders) SetCommonHeaders(v map[string]*string) *TalentV2AddCustomTagHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TalentV2AddCustomTagHeaders) SetXAcsDingtalkAccessToken(v string) *TalentV2AddCustomTagHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TalentV2AddCustomTagRequest struct {
	SortOrder *int32  `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	TagCode   *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
	// This parameter is required.
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s TalentV2AddCustomTagRequest) String() string {
	return tea.Prettify(s)
}

func (s TalentV2AddCustomTagRequest) GoString() string {
	return s.String()
}

func (s *TalentV2AddCustomTagRequest) SetSortOrder(v int32) *TalentV2AddCustomTagRequest {
	s.SortOrder = &v
	return s
}

func (s *TalentV2AddCustomTagRequest) SetTagCode(v string) *TalentV2AddCustomTagRequest {
	s.TagCode = &v
	return s
}

func (s *TalentV2AddCustomTagRequest) SetTagName(v string) *TalentV2AddCustomTagRequest {
	s.TagName = &v
	return s
}

func (s *TalentV2AddCustomTagRequest) SetUserId(v string) *TalentV2AddCustomTagRequest {
	s.UserId = &v
	return s
}

type TalentV2AddCustomTagResponseBody struct {
	Result *TalentV2AddCustomTagResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s TalentV2AddCustomTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TalentV2AddCustomTagResponseBody) GoString() string {
	return s.String()
}

func (s *TalentV2AddCustomTagResponseBody) SetResult(v *TalentV2AddCustomTagResponseBodyResult) *TalentV2AddCustomTagResponseBody {
	s.Result = v
	return s
}

type TalentV2AddCustomTagResponseBodyResult struct {
	TagCode *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s TalentV2AddCustomTagResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s TalentV2AddCustomTagResponseBodyResult) GoString() string {
	return s.String()
}

func (s *TalentV2AddCustomTagResponseBodyResult) SetTagCode(v string) *TalentV2AddCustomTagResponseBodyResult {
	s.TagCode = &v
	return s
}

func (s *TalentV2AddCustomTagResponseBodyResult) SetTagName(v string) *TalentV2AddCustomTagResponseBodyResult {
	s.TagName = &v
	return s
}

type TalentV2AddCustomTagResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TalentV2AddCustomTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TalentV2AddCustomTagResponse) String() string {
	return tea.Prettify(s)
}

func (s TalentV2AddCustomTagResponse) GoString() string {
	return s.String()
}

func (s *TalentV2AddCustomTagResponse) SetHeaders(v map[string]*string) *TalentV2AddCustomTagResponse {
	s.Headers = v
	return s
}

func (s *TalentV2AddCustomTagResponse) SetStatusCode(v int32) *TalentV2AddCustomTagResponse {
	s.StatusCode = &v
	return s
}

func (s *TalentV2AddCustomTagResponse) SetBody(v *TalentV2AddCustomTagResponseBody) *TalentV2AddCustomTagResponse {
	s.Body = v
	return s
}

type TalentV2AddObjectiveTagHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TalentV2AddObjectiveTagHeaders) String() string {
	return tea.Prettify(s)
}

func (s TalentV2AddObjectiveTagHeaders) GoString() string {
	return s.String()
}

func (s *TalentV2AddObjectiveTagHeaders) SetCommonHeaders(v map[string]*string) *TalentV2AddObjectiveTagHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TalentV2AddObjectiveTagHeaders) SetXAcsDingtalkAccessToken(v string) *TalentV2AddObjectiveTagHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TalentV2AddObjectiveTagRequest struct {
	SortOrder *int32  `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	TagCode   *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
	// This parameter is required.
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s TalentV2AddObjectiveTagRequest) String() string {
	return tea.Prettify(s)
}

func (s TalentV2AddObjectiveTagRequest) GoString() string {
	return s.String()
}

func (s *TalentV2AddObjectiveTagRequest) SetSortOrder(v int32) *TalentV2AddObjectiveTagRequest {
	s.SortOrder = &v
	return s
}

func (s *TalentV2AddObjectiveTagRequest) SetTagCode(v string) *TalentV2AddObjectiveTagRequest {
	s.TagCode = &v
	return s
}

func (s *TalentV2AddObjectiveTagRequest) SetTagName(v string) *TalentV2AddObjectiveTagRequest {
	s.TagName = &v
	return s
}

func (s *TalentV2AddObjectiveTagRequest) SetUserId(v string) *TalentV2AddObjectiveTagRequest {
	s.UserId = &v
	return s
}

type TalentV2AddObjectiveTagResponseBody struct {
	Result *TalentV2AddObjectiveTagResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s TalentV2AddObjectiveTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TalentV2AddObjectiveTagResponseBody) GoString() string {
	return s.String()
}

func (s *TalentV2AddObjectiveTagResponseBody) SetResult(v *TalentV2AddObjectiveTagResponseBodyResult) *TalentV2AddObjectiveTagResponseBody {
	s.Result = v
	return s
}

type TalentV2AddObjectiveTagResponseBodyResult struct {
	TagCode *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s TalentV2AddObjectiveTagResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s TalentV2AddObjectiveTagResponseBodyResult) GoString() string {
	return s.String()
}

func (s *TalentV2AddObjectiveTagResponseBodyResult) SetTagCode(v string) *TalentV2AddObjectiveTagResponseBodyResult {
	s.TagCode = &v
	return s
}

func (s *TalentV2AddObjectiveTagResponseBodyResult) SetTagName(v string) *TalentV2AddObjectiveTagResponseBodyResult {
	s.TagName = &v
	return s
}

type TalentV2AddObjectiveTagResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TalentV2AddObjectiveTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TalentV2AddObjectiveTagResponse) String() string {
	return tea.Prettify(s)
}

func (s TalentV2AddObjectiveTagResponse) GoString() string {
	return s.String()
}

func (s *TalentV2AddObjectiveTagResponse) SetHeaders(v map[string]*string) *TalentV2AddObjectiveTagResponse {
	s.Headers = v
	return s
}

func (s *TalentV2AddObjectiveTagResponse) SetStatusCode(v int32) *TalentV2AddObjectiveTagResponse {
	s.StatusCode = &v
	return s
}

func (s *TalentV2AddObjectiveTagResponse) SetBody(v *TalentV2AddObjectiveTagResponseBody) *TalentV2AddObjectiveTagResponse {
	s.Body = v
	return s
}

type TalentV2AddPersonalityTagHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TalentV2AddPersonalityTagHeaders) String() string {
	return tea.Prettify(s)
}

func (s TalentV2AddPersonalityTagHeaders) GoString() string {
	return s.String()
}

func (s *TalentV2AddPersonalityTagHeaders) SetCommonHeaders(v map[string]*string) *TalentV2AddPersonalityTagHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TalentV2AddPersonalityTagHeaders) SetXAcsDingtalkAccessToken(v string) *TalentV2AddPersonalityTagHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TalentV2AddPersonalityTagRequest struct {
	CategoryCode *string `json:"categoryCode,omitempty" xml:"categoryCode,omitempty"`
	// This parameter is required.
	CategoryName      *string `json:"categoryName,omitempty" xml:"categoryName,omitempty"`
	CategorySortOrder *int32  `json:"categorySortOrder,omitempty" xml:"categorySortOrder,omitempty"`
	SortOrder         *int32  `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	TagCode           *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
	// This parameter is required.
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s TalentV2AddPersonalityTagRequest) String() string {
	return tea.Prettify(s)
}

func (s TalentV2AddPersonalityTagRequest) GoString() string {
	return s.String()
}

func (s *TalentV2AddPersonalityTagRequest) SetCategoryCode(v string) *TalentV2AddPersonalityTagRequest {
	s.CategoryCode = &v
	return s
}

func (s *TalentV2AddPersonalityTagRequest) SetCategoryName(v string) *TalentV2AddPersonalityTagRequest {
	s.CategoryName = &v
	return s
}

func (s *TalentV2AddPersonalityTagRequest) SetCategorySortOrder(v int32) *TalentV2AddPersonalityTagRequest {
	s.CategorySortOrder = &v
	return s
}

func (s *TalentV2AddPersonalityTagRequest) SetSortOrder(v int32) *TalentV2AddPersonalityTagRequest {
	s.SortOrder = &v
	return s
}

func (s *TalentV2AddPersonalityTagRequest) SetTagCode(v string) *TalentV2AddPersonalityTagRequest {
	s.TagCode = &v
	return s
}

func (s *TalentV2AddPersonalityTagRequest) SetTagName(v string) *TalentV2AddPersonalityTagRequest {
	s.TagName = &v
	return s
}

type TalentV2AddPersonalityTagResponseBody struct {
	Result *TalentV2AddPersonalityTagResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s TalentV2AddPersonalityTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TalentV2AddPersonalityTagResponseBody) GoString() string {
	return s.String()
}

func (s *TalentV2AddPersonalityTagResponseBody) SetResult(v *TalentV2AddPersonalityTagResponseBodyResult) *TalentV2AddPersonalityTagResponseBody {
	s.Result = v
	return s
}

type TalentV2AddPersonalityTagResponseBodyResult struct {
	CategoryCode *string `json:"categoryCode,omitempty" xml:"categoryCode,omitempty"`
	CategoryName *string `json:"categoryName,omitempty" xml:"categoryName,omitempty"`
	TagCode      *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
	TagName      *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s TalentV2AddPersonalityTagResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s TalentV2AddPersonalityTagResponseBodyResult) GoString() string {
	return s.String()
}

func (s *TalentV2AddPersonalityTagResponseBodyResult) SetCategoryCode(v string) *TalentV2AddPersonalityTagResponseBodyResult {
	s.CategoryCode = &v
	return s
}

func (s *TalentV2AddPersonalityTagResponseBodyResult) SetCategoryName(v string) *TalentV2AddPersonalityTagResponseBodyResult {
	s.CategoryName = &v
	return s
}

func (s *TalentV2AddPersonalityTagResponseBodyResult) SetTagCode(v string) *TalentV2AddPersonalityTagResponseBodyResult {
	s.TagCode = &v
	return s
}

func (s *TalentV2AddPersonalityTagResponseBodyResult) SetTagName(v string) *TalentV2AddPersonalityTagResponseBodyResult {
	s.TagName = &v
	return s
}

type TalentV2AddPersonalityTagResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TalentV2AddPersonalityTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TalentV2AddPersonalityTagResponse) String() string {
	return tea.Prettify(s)
}

func (s TalentV2AddPersonalityTagResponse) GoString() string {
	return s.String()
}

func (s *TalentV2AddPersonalityTagResponse) SetHeaders(v map[string]*string) *TalentV2AddPersonalityTagResponse {
	s.Headers = v
	return s
}

func (s *TalentV2AddPersonalityTagResponse) SetStatusCode(v int32) *TalentV2AddPersonalityTagResponse {
	s.StatusCode = &v
	return s
}

func (s *TalentV2AddPersonalityTagResponse) SetBody(v *TalentV2AddPersonalityTagResponseBody) *TalentV2AddPersonalityTagResponse {
	s.Body = v
	return s
}

type TalentV2DeleteCustomTagHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TalentV2DeleteCustomTagHeaders) String() string {
	return tea.Prettify(s)
}

func (s TalentV2DeleteCustomTagHeaders) GoString() string {
	return s.String()
}

func (s *TalentV2DeleteCustomTagHeaders) SetCommonHeaders(v map[string]*string) *TalentV2DeleteCustomTagHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TalentV2DeleteCustomTagHeaders) SetXAcsDingtalkAccessToken(v string) *TalentV2DeleteCustomTagHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TalentV2DeleteCustomTagRequest struct {
	// This parameter is required.
	TagCode *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s TalentV2DeleteCustomTagRequest) String() string {
	return tea.Prettify(s)
}

func (s TalentV2DeleteCustomTagRequest) GoString() string {
	return s.String()
}

func (s *TalentV2DeleteCustomTagRequest) SetTagCode(v string) *TalentV2DeleteCustomTagRequest {
	s.TagCode = &v
	return s
}

func (s *TalentV2DeleteCustomTagRequest) SetUserId(v string) *TalentV2DeleteCustomTagRequest {
	s.UserId = &v
	return s
}

type TalentV2DeleteCustomTagResponseBody struct {
	Result *TalentV2DeleteCustomTagResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s TalentV2DeleteCustomTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TalentV2DeleteCustomTagResponseBody) GoString() string {
	return s.String()
}

func (s *TalentV2DeleteCustomTagResponseBody) SetResult(v *TalentV2DeleteCustomTagResponseBodyResult) *TalentV2DeleteCustomTagResponseBody {
	s.Result = v
	return s
}

type TalentV2DeleteCustomTagResponseBodyResult struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TalentV2DeleteCustomTagResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s TalentV2DeleteCustomTagResponseBodyResult) GoString() string {
	return s.String()
}

func (s *TalentV2DeleteCustomTagResponseBodyResult) SetSuccess(v bool) *TalentV2DeleteCustomTagResponseBodyResult {
	s.Success = &v
	return s
}

type TalentV2DeleteCustomTagResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TalentV2DeleteCustomTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TalentV2DeleteCustomTagResponse) String() string {
	return tea.Prettify(s)
}

func (s TalentV2DeleteCustomTagResponse) GoString() string {
	return s.String()
}

func (s *TalentV2DeleteCustomTagResponse) SetHeaders(v map[string]*string) *TalentV2DeleteCustomTagResponse {
	s.Headers = v
	return s
}

func (s *TalentV2DeleteCustomTagResponse) SetStatusCode(v int32) *TalentV2DeleteCustomTagResponse {
	s.StatusCode = &v
	return s
}

func (s *TalentV2DeleteCustomTagResponse) SetBody(v *TalentV2DeleteCustomTagResponseBody) *TalentV2DeleteCustomTagResponse {
	s.Body = v
	return s
}

type TalentV2DeleteObjectiveTagHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TalentV2DeleteObjectiveTagHeaders) String() string {
	return tea.Prettify(s)
}

func (s TalentV2DeleteObjectiveTagHeaders) GoString() string {
	return s.String()
}

func (s *TalentV2DeleteObjectiveTagHeaders) SetCommonHeaders(v map[string]*string) *TalentV2DeleteObjectiveTagHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TalentV2DeleteObjectiveTagHeaders) SetXAcsDingtalkAccessToken(v string) *TalentV2DeleteObjectiveTagHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TalentV2DeleteObjectiveTagRequest struct {
	// This parameter is required.
	TagCode *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s TalentV2DeleteObjectiveTagRequest) String() string {
	return tea.Prettify(s)
}

func (s TalentV2DeleteObjectiveTagRequest) GoString() string {
	return s.String()
}

func (s *TalentV2DeleteObjectiveTagRequest) SetTagCode(v string) *TalentV2DeleteObjectiveTagRequest {
	s.TagCode = &v
	return s
}

func (s *TalentV2DeleteObjectiveTagRequest) SetUserId(v string) *TalentV2DeleteObjectiveTagRequest {
	s.UserId = &v
	return s
}

type TalentV2DeleteObjectiveTagResponseBody struct {
	Result *TalentV2DeleteObjectiveTagResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s TalentV2DeleteObjectiveTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TalentV2DeleteObjectiveTagResponseBody) GoString() string {
	return s.String()
}

func (s *TalentV2DeleteObjectiveTagResponseBody) SetResult(v *TalentV2DeleteObjectiveTagResponseBodyResult) *TalentV2DeleteObjectiveTagResponseBody {
	s.Result = v
	return s
}

type TalentV2DeleteObjectiveTagResponseBodyResult struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TalentV2DeleteObjectiveTagResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s TalentV2DeleteObjectiveTagResponseBodyResult) GoString() string {
	return s.String()
}

func (s *TalentV2DeleteObjectiveTagResponseBodyResult) SetSuccess(v bool) *TalentV2DeleteObjectiveTagResponseBodyResult {
	s.Success = &v
	return s
}

type TalentV2DeleteObjectiveTagResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TalentV2DeleteObjectiveTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TalentV2DeleteObjectiveTagResponse) String() string {
	return tea.Prettify(s)
}

func (s TalentV2DeleteObjectiveTagResponse) GoString() string {
	return s.String()
}

func (s *TalentV2DeleteObjectiveTagResponse) SetHeaders(v map[string]*string) *TalentV2DeleteObjectiveTagResponse {
	s.Headers = v
	return s
}

func (s *TalentV2DeleteObjectiveTagResponse) SetStatusCode(v int32) *TalentV2DeleteObjectiveTagResponse {
	s.StatusCode = &v
	return s
}

func (s *TalentV2DeleteObjectiveTagResponse) SetBody(v *TalentV2DeleteObjectiveTagResponseBody) *TalentV2DeleteObjectiveTagResponse {
	s.Body = v
	return s
}

type TalentV2DeletePersonalityTagHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TalentV2DeletePersonalityTagHeaders) String() string {
	return tea.Prettify(s)
}

func (s TalentV2DeletePersonalityTagHeaders) GoString() string {
	return s.String()
}

func (s *TalentV2DeletePersonalityTagHeaders) SetCommonHeaders(v map[string]*string) *TalentV2DeletePersonalityTagHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TalentV2DeletePersonalityTagHeaders) SetXAcsDingtalkAccessToken(v string) *TalentV2DeletePersonalityTagHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TalentV2DeletePersonalityTagRequest struct {
	// This parameter is required.
	TagCode *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
}

func (s TalentV2DeletePersonalityTagRequest) String() string {
	return tea.Prettify(s)
}

func (s TalentV2DeletePersonalityTagRequest) GoString() string {
	return s.String()
}

func (s *TalentV2DeletePersonalityTagRequest) SetTagCode(v string) *TalentV2DeletePersonalityTagRequest {
	s.TagCode = &v
	return s
}

type TalentV2DeletePersonalityTagResponseBody struct {
	Result *TalentV2DeletePersonalityTagResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s TalentV2DeletePersonalityTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TalentV2DeletePersonalityTagResponseBody) GoString() string {
	return s.String()
}

func (s *TalentV2DeletePersonalityTagResponseBody) SetResult(v *TalentV2DeletePersonalityTagResponseBodyResult) *TalentV2DeletePersonalityTagResponseBody {
	s.Result = v
	return s
}

type TalentV2DeletePersonalityTagResponseBodyResult struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TalentV2DeletePersonalityTagResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s TalentV2DeletePersonalityTagResponseBodyResult) GoString() string {
	return s.String()
}

func (s *TalentV2DeletePersonalityTagResponseBodyResult) SetSuccess(v bool) *TalentV2DeletePersonalityTagResponseBodyResult {
	s.Success = &v
	return s
}

type TalentV2DeletePersonalityTagResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TalentV2DeletePersonalityTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TalentV2DeletePersonalityTagResponse) String() string {
	return tea.Prettify(s)
}

func (s TalentV2DeletePersonalityTagResponse) GoString() string {
	return s.String()
}

func (s *TalentV2DeletePersonalityTagResponse) SetHeaders(v map[string]*string) *TalentV2DeletePersonalityTagResponse {
	s.Headers = v
	return s
}

func (s *TalentV2DeletePersonalityTagResponse) SetStatusCode(v int32) *TalentV2DeletePersonalityTagResponse {
	s.StatusCode = &v
	return s
}

func (s *TalentV2DeletePersonalityTagResponse) SetBody(v *TalentV2DeletePersonalityTagResponseBody) *TalentV2DeletePersonalityTagResponse {
	s.Body = v
	return s
}

type TalentV2LikeTagHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TalentV2LikeTagHeaders) String() string {
	return tea.Prettify(s)
}

func (s TalentV2LikeTagHeaders) GoString() string {
	return s.String()
}

func (s *TalentV2LikeTagHeaders) SetCommonHeaders(v map[string]*string) *TalentV2LikeTagHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TalentV2LikeTagHeaders) SetXAcsDingtalkAccessToken(v string) *TalentV2LikeTagHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TalentV2LikeTagRequest struct {
	// This parameter is required.
	ActionType *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// This parameter is required.
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// This parameter is required.
	TagCode *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s TalentV2LikeTagRequest) String() string {
	return tea.Prettify(s)
}

func (s TalentV2LikeTagRequest) GoString() string {
	return s.String()
}

func (s *TalentV2LikeTagRequest) SetActionType(v string) *TalentV2LikeTagRequest {
	s.ActionType = &v
	return s
}

func (s *TalentV2LikeTagRequest) SetOperatorUserId(v string) *TalentV2LikeTagRequest {
	s.OperatorUserId = &v
	return s
}

func (s *TalentV2LikeTagRequest) SetTagCode(v string) *TalentV2LikeTagRequest {
	s.TagCode = &v
	return s
}

func (s *TalentV2LikeTagRequest) SetUserId(v string) *TalentV2LikeTagRequest {
	s.UserId = &v
	return s
}

type TalentV2LikeTagResponseBody struct {
	Result *TalentV2LikeTagResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s TalentV2LikeTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TalentV2LikeTagResponseBody) GoString() string {
	return s.String()
}

func (s *TalentV2LikeTagResponseBody) SetResult(v *TalentV2LikeTagResponseBodyResult) *TalentV2LikeTagResponseBody {
	s.Result = v
	return s
}

type TalentV2LikeTagResponseBodyResult struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TalentV2LikeTagResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s TalentV2LikeTagResponseBodyResult) GoString() string {
	return s.String()
}

func (s *TalentV2LikeTagResponseBodyResult) SetSuccess(v bool) *TalentV2LikeTagResponseBodyResult {
	s.Success = &v
	return s
}

type TalentV2LikeTagResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TalentV2LikeTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TalentV2LikeTagResponse) String() string {
	return tea.Prettify(s)
}

func (s TalentV2LikeTagResponse) GoString() string {
	return s.String()
}

func (s *TalentV2LikeTagResponse) SetHeaders(v map[string]*string) *TalentV2LikeTagResponse {
	s.Headers = v
	return s
}

func (s *TalentV2LikeTagResponse) SetStatusCode(v int32) *TalentV2LikeTagResponse {
	s.StatusCode = &v
	return s
}

func (s *TalentV2LikeTagResponse) SetBody(v *TalentV2LikeTagResponseBody) *TalentV2LikeTagResponse {
	s.Body = v
	return s
}

type TalentV2QueryCustomTagHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TalentV2QueryCustomTagHeaders) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryCustomTagHeaders) GoString() string {
	return s.String()
}

func (s *TalentV2QueryCustomTagHeaders) SetCommonHeaders(v map[string]*string) *TalentV2QueryCustomTagHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TalentV2QueryCustomTagHeaders) SetXAcsDingtalkAccessToken(v string) *TalentV2QueryCustomTagHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TalentV2QueryCustomTagRequest struct {
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s TalentV2QueryCustomTagRequest) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryCustomTagRequest) GoString() string {
	return s.String()
}

func (s *TalentV2QueryCustomTagRequest) SetUserId(v string) *TalentV2QueryCustomTagRequest {
	s.UserId = &v
	return s
}

type TalentV2QueryCustomTagResponseBody struct {
	Result *TalentV2QueryCustomTagResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s TalentV2QueryCustomTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryCustomTagResponseBody) GoString() string {
	return s.String()
}

func (s *TalentV2QueryCustomTagResponseBody) SetResult(v *TalentV2QueryCustomTagResponseBodyResult) *TalentV2QueryCustomTagResponseBody {
	s.Result = v
	return s
}

type TalentV2QueryCustomTagResponseBodyResult struct {
	Tags []*TalentV2QueryCustomTagResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s TalentV2QueryCustomTagResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryCustomTagResponseBodyResult) GoString() string {
	return s.String()
}

func (s *TalentV2QueryCustomTagResponseBodyResult) SetTags(v []*TalentV2QueryCustomTagResponseBodyResultTags) *TalentV2QueryCustomTagResponseBodyResult {
	s.Tags = v
	return s
}

type TalentV2QueryCustomTagResponseBodyResultTags struct {
	SortOrder *int32  `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	TagCode   *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
	TagName   *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s TalentV2QueryCustomTagResponseBodyResultTags) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryCustomTagResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *TalentV2QueryCustomTagResponseBodyResultTags) SetSortOrder(v int32) *TalentV2QueryCustomTagResponseBodyResultTags {
	s.SortOrder = &v
	return s
}

func (s *TalentV2QueryCustomTagResponseBodyResultTags) SetTagCode(v string) *TalentV2QueryCustomTagResponseBodyResultTags {
	s.TagCode = &v
	return s
}

func (s *TalentV2QueryCustomTagResponseBodyResultTags) SetTagName(v string) *TalentV2QueryCustomTagResponseBodyResultTags {
	s.TagName = &v
	return s
}

type TalentV2QueryCustomTagResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TalentV2QueryCustomTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TalentV2QueryCustomTagResponse) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryCustomTagResponse) GoString() string {
	return s.String()
}

func (s *TalentV2QueryCustomTagResponse) SetHeaders(v map[string]*string) *TalentV2QueryCustomTagResponse {
	s.Headers = v
	return s
}

func (s *TalentV2QueryCustomTagResponse) SetStatusCode(v int32) *TalentV2QueryCustomTagResponse {
	s.StatusCode = &v
	return s
}

func (s *TalentV2QueryCustomTagResponse) SetBody(v *TalentV2QueryCustomTagResponseBody) *TalentV2QueryCustomTagResponse {
	s.Body = v
	return s
}

type TalentV2QueryObjectiveTagHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TalentV2QueryObjectiveTagHeaders) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryObjectiveTagHeaders) GoString() string {
	return s.String()
}

func (s *TalentV2QueryObjectiveTagHeaders) SetCommonHeaders(v map[string]*string) *TalentV2QueryObjectiveTagHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TalentV2QueryObjectiveTagHeaders) SetXAcsDingtalkAccessToken(v string) *TalentV2QueryObjectiveTagHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TalentV2QueryObjectiveTagRequest struct {
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s TalentV2QueryObjectiveTagRequest) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryObjectiveTagRequest) GoString() string {
	return s.String()
}

func (s *TalentV2QueryObjectiveTagRequest) SetUserId(v string) *TalentV2QueryObjectiveTagRequest {
	s.UserId = &v
	return s
}

type TalentV2QueryObjectiveTagResponseBody struct {
	Result *TalentV2QueryObjectiveTagResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s TalentV2QueryObjectiveTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryObjectiveTagResponseBody) GoString() string {
	return s.String()
}

func (s *TalentV2QueryObjectiveTagResponseBody) SetResult(v *TalentV2QueryObjectiveTagResponseBodyResult) *TalentV2QueryObjectiveTagResponseBody {
	s.Result = v
	return s
}

type TalentV2QueryObjectiveTagResponseBodyResult struct {
	Tags []*TalentV2QueryObjectiveTagResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s TalentV2QueryObjectiveTagResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryObjectiveTagResponseBodyResult) GoString() string {
	return s.String()
}

func (s *TalentV2QueryObjectiveTagResponseBodyResult) SetTags(v []*TalentV2QueryObjectiveTagResponseBodyResultTags) *TalentV2QueryObjectiveTagResponseBodyResult {
	s.Tags = v
	return s
}

type TalentV2QueryObjectiveTagResponseBodyResultTags struct {
	SortOrder *int32  `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	TagCode   *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
	TagName   *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s TalentV2QueryObjectiveTagResponseBodyResultTags) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryObjectiveTagResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *TalentV2QueryObjectiveTagResponseBodyResultTags) SetSortOrder(v int32) *TalentV2QueryObjectiveTagResponseBodyResultTags {
	s.SortOrder = &v
	return s
}

func (s *TalentV2QueryObjectiveTagResponseBodyResultTags) SetTagCode(v string) *TalentV2QueryObjectiveTagResponseBodyResultTags {
	s.TagCode = &v
	return s
}

func (s *TalentV2QueryObjectiveTagResponseBodyResultTags) SetTagName(v string) *TalentV2QueryObjectiveTagResponseBodyResultTags {
	s.TagName = &v
	return s
}

type TalentV2QueryObjectiveTagResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TalentV2QueryObjectiveTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TalentV2QueryObjectiveTagResponse) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryObjectiveTagResponse) GoString() string {
	return s.String()
}

func (s *TalentV2QueryObjectiveTagResponse) SetHeaders(v map[string]*string) *TalentV2QueryObjectiveTagResponse {
	s.Headers = v
	return s
}

func (s *TalentV2QueryObjectiveTagResponse) SetStatusCode(v int32) *TalentV2QueryObjectiveTagResponse {
	s.StatusCode = &v
	return s
}

func (s *TalentV2QueryObjectiveTagResponse) SetBody(v *TalentV2QueryObjectiveTagResponseBody) *TalentV2QueryObjectiveTagResponse {
	s.Body = v
	return s
}

type TalentV2QueryPersonalityTagHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TalentV2QueryPersonalityTagHeaders) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryPersonalityTagHeaders) GoString() string {
	return s.String()
}

func (s *TalentV2QueryPersonalityTagHeaders) SetCommonHeaders(v map[string]*string) *TalentV2QueryPersonalityTagHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TalentV2QueryPersonalityTagHeaders) SetXAcsDingtalkAccessToken(v string) *TalentV2QueryPersonalityTagHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TalentV2QueryPersonalityTagResponseBody struct {
	Result *TalentV2QueryPersonalityTagResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s TalentV2QueryPersonalityTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryPersonalityTagResponseBody) GoString() string {
	return s.String()
}

func (s *TalentV2QueryPersonalityTagResponseBody) SetResult(v *TalentV2QueryPersonalityTagResponseBodyResult) *TalentV2QueryPersonalityTagResponseBody {
	s.Result = v
	return s
}

type TalentV2QueryPersonalityTagResponseBodyResult struct {
	Tags []*TalentV2QueryPersonalityTagResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s TalentV2QueryPersonalityTagResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryPersonalityTagResponseBodyResult) GoString() string {
	return s.String()
}

func (s *TalentV2QueryPersonalityTagResponseBodyResult) SetTags(v []*TalentV2QueryPersonalityTagResponseBodyResultTags) *TalentV2QueryPersonalityTagResponseBodyResult {
	s.Tags = v
	return s
}

type TalentV2QueryPersonalityTagResponseBodyResultTags struct {
	CategoryCode      *string `json:"categoryCode,omitempty" xml:"categoryCode,omitempty"`
	CategoryName      *string `json:"categoryName,omitempty" xml:"categoryName,omitempty"`
	CategorySortOrder *int32  `json:"categorySortOrder,omitempty" xml:"categorySortOrder,omitempty"`
	SortOrder         *int32  `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	TagCode           *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
	TagName           *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s TalentV2QueryPersonalityTagResponseBodyResultTags) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryPersonalityTagResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *TalentV2QueryPersonalityTagResponseBodyResultTags) SetCategoryCode(v string) *TalentV2QueryPersonalityTagResponseBodyResultTags {
	s.CategoryCode = &v
	return s
}

func (s *TalentV2QueryPersonalityTagResponseBodyResultTags) SetCategoryName(v string) *TalentV2QueryPersonalityTagResponseBodyResultTags {
	s.CategoryName = &v
	return s
}

func (s *TalentV2QueryPersonalityTagResponseBodyResultTags) SetCategorySortOrder(v int32) *TalentV2QueryPersonalityTagResponseBodyResultTags {
	s.CategorySortOrder = &v
	return s
}

func (s *TalentV2QueryPersonalityTagResponseBodyResultTags) SetSortOrder(v int32) *TalentV2QueryPersonalityTagResponseBodyResultTags {
	s.SortOrder = &v
	return s
}

func (s *TalentV2QueryPersonalityTagResponseBodyResultTags) SetTagCode(v string) *TalentV2QueryPersonalityTagResponseBodyResultTags {
	s.TagCode = &v
	return s
}

func (s *TalentV2QueryPersonalityTagResponseBodyResultTags) SetTagName(v string) *TalentV2QueryPersonalityTagResponseBodyResultTags {
	s.TagName = &v
	return s
}

type TalentV2QueryPersonalityTagResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TalentV2QueryPersonalityTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TalentV2QueryPersonalityTagResponse) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryPersonalityTagResponse) GoString() string {
	return s.String()
}

func (s *TalentV2QueryPersonalityTagResponse) SetHeaders(v map[string]*string) *TalentV2QueryPersonalityTagResponse {
	s.Headers = v
	return s
}

func (s *TalentV2QueryPersonalityTagResponse) SetStatusCode(v int32) *TalentV2QueryPersonalityTagResponse {
	s.StatusCode = &v
	return s
}

func (s *TalentV2QueryPersonalityTagResponse) SetBody(v *TalentV2QueryPersonalityTagResponseBody) *TalentV2QueryPersonalityTagResponse {
	s.Body = v
	return s
}

type TalentV2QueryTagLikeDetailListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TalentV2QueryTagLikeDetailListHeaders) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryTagLikeDetailListHeaders) GoString() string {
	return s.String()
}

func (s *TalentV2QueryTagLikeDetailListHeaders) SetCommonHeaders(v map[string]*string) *TalentV2QueryTagLikeDetailListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TalentV2QueryTagLikeDetailListHeaders) SetXAcsDingtalkAccessToken(v string) *TalentV2QueryTagLikeDetailListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TalentV2QueryTagLikeDetailListRequest struct {
	Cursor *int64 `json:"cursor,omitempty" xml:"cursor,omitempty"`
	Size   *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// This parameter is required.
	TagCode *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s TalentV2QueryTagLikeDetailListRequest) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryTagLikeDetailListRequest) GoString() string {
	return s.String()
}

func (s *TalentV2QueryTagLikeDetailListRequest) SetCursor(v int64) *TalentV2QueryTagLikeDetailListRequest {
	s.Cursor = &v
	return s
}

func (s *TalentV2QueryTagLikeDetailListRequest) SetSize(v int32) *TalentV2QueryTagLikeDetailListRequest {
	s.Size = &v
	return s
}

func (s *TalentV2QueryTagLikeDetailListRequest) SetTagCode(v string) *TalentV2QueryTagLikeDetailListRequest {
	s.TagCode = &v
	return s
}

func (s *TalentV2QueryTagLikeDetailListRequest) SetUserId(v string) *TalentV2QueryTagLikeDetailListRequest {
	s.UserId = &v
	return s
}

type TalentV2QueryTagLikeDetailListResponseBody struct {
	Result *TalentV2QueryTagLikeDetailListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s TalentV2QueryTagLikeDetailListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryTagLikeDetailListResponseBody) GoString() string {
	return s.String()
}

func (s *TalentV2QueryTagLikeDetailListResponseBody) SetResult(v *TalentV2QueryTagLikeDetailListResponseBodyResult) *TalentV2QueryTagLikeDetailListResponseBody {
	s.Result = v
	return s
}

type TalentV2QueryTagLikeDetailListResponseBodyResult struct {
	HasMore     *bool                                                          `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	LikeDetails []*TalentV2QueryTagLikeDetailListResponseBodyResultLikeDetails `json:"likeDetails,omitempty" xml:"likeDetails,omitempty" type:"Repeated"`
	NextCursor  *int64                                                         `json:"nextCursor,omitempty" xml:"nextCursor,omitempty"`
}

func (s TalentV2QueryTagLikeDetailListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryTagLikeDetailListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *TalentV2QueryTagLikeDetailListResponseBodyResult) SetHasMore(v bool) *TalentV2QueryTagLikeDetailListResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *TalentV2QueryTagLikeDetailListResponseBodyResult) SetLikeDetails(v []*TalentV2QueryTagLikeDetailListResponseBodyResultLikeDetails) *TalentV2QueryTagLikeDetailListResponseBodyResult {
	s.LikeDetails = v
	return s
}

func (s *TalentV2QueryTagLikeDetailListResponseBodyResult) SetNextCursor(v int64) *TalentV2QueryTagLikeDetailListResponseBodyResult {
	s.NextCursor = &v
	return s
}

type TalentV2QueryTagLikeDetailListResponseBodyResultLikeDetails struct {
	LikeTimestamp  *int64  `json:"likeTimestamp,omitempty" xml:"likeTimestamp,omitempty"`
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
}

func (s TalentV2QueryTagLikeDetailListResponseBodyResultLikeDetails) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryTagLikeDetailListResponseBodyResultLikeDetails) GoString() string {
	return s.String()
}

func (s *TalentV2QueryTagLikeDetailListResponseBodyResultLikeDetails) SetLikeTimestamp(v int64) *TalentV2QueryTagLikeDetailListResponseBodyResultLikeDetails {
	s.LikeTimestamp = &v
	return s
}

func (s *TalentV2QueryTagLikeDetailListResponseBodyResultLikeDetails) SetOperatorUserId(v string) *TalentV2QueryTagLikeDetailListResponseBodyResultLikeDetails {
	s.OperatorUserId = &v
	return s
}

type TalentV2QueryTagLikeDetailListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TalentV2QueryTagLikeDetailListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TalentV2QueryTagLikeDetailListResponse) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryTagLikeDetailListResponse) GoString() string {
	return s.String()
}

func (s *TalentV2QueryTagLikeDetailListResponse) SetHeaders(v map[string]*string) *TalentV2QueryTagLikeDetailListResponse {
	s.Headers = v
	return s
}

func (s *TalentV2QueryTagLikeDetailListResponse) SetStatusCode(v int32) *TalentV2QueryTagLikeDetailListResponse {
	s.StatusCode = &v
	return s
}

func (s *TalentV2QueryTagLikeDetailListResponse) SetBody(v *TalentV2QueryTagLikeDetailListResponseBody) *TalentV2QueryTagLikeDetailListResponse {
	s.Body = v
	return s
}

type TalentV2QueryTagLikeListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TalentV2QueryTagLikeListHeaders) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryTagLikeListHeaders) GoString() string {
	return s.String()
}

func (s *TalentV2QueryTagLikeListHeaders) SetCommonHeaders(v map[string]*string) *TalentV2QueryTagLikeListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TalentV2QueryTagLikeListHeaders) SetXAcsDingtalkAccessToken(v string) *TalentV2QueryTagLikeListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TalentV2QueryTagLikeListRequest struct {
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s TalentV2QueryTagLikeListRequest) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryTagLikeListRequest) GoString() string {
	return s.String()
}

func (s *TalentV2QueryTagLikeListRequest) SetOperatorUserId(v string) *TalentV2QueryTagLikeListRequest {
	s.OperatorUserId = &v
	return s
}

func (s *TalentV2QueryTagLikeListRequest) SetUserId(v string) *TalentV2QueryTagLikeListRequest {
	s.UserId = &v
	return s
}

type TalentV2QueryTagLikeListResponseBody struct {
	Result *TalentV2QueryTagLikeListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s TalentV2QueryTagLikeListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryTagLikeListResponseBody) GoString() string {
	return s.String()
}

func (s *TalentV2QueryTagLikeListResponseBody) SetResult(v *TalentV2QueryTagLikeListResponseBodyResult) *TalentV2QueryTagLikeListResponseBody {
	s.Result = v
	return s
}

type TalentV2QueryTagLikeListResponseBodyResult struct {
	TagLikes []*TalentV2QueryTagLikeListResponseBodyResultTagLikes `json:"tagLikes,omitempty" xml:"tagLikes,omitempty" type:"Repeated"`
}

func (s TalentV2QueryTagLikeListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryTagLikeListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *TalentV2QueryTagLikeListResponseBodyResult) SetTagLikes(v []*TalentV2QueryTagLikeListResponseBodyResultTagLikes) *TalentV2QueryTagLikeListResponseBodyResult {
	s.TagLikes = v
	return s
}

type TalentV2QueryTagLikeListResponseBodyResultTagLikes struct {
	HasLiked  *bool   `json:"hasLiked,omitempty" xml:"hasLiked,omitempty"`
	LikeCount *int32  `json:"likeCount,omitempty" xml:"likeCount,omitempty"`
	TagCode   *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
	TagName   *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s TalentV2QueryTagLikeListResponseBodyResultTagLikes) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryTagLikeListResponseBodyResultTagLikes) GoString() string {
	return s.String()
}

func (s *TalentV2QueryTagLikeListResponseBodyResultTagLikes) SetHasLiked(v bool) *TalentV2QueryTagLikeListResponseBodyResultTagLikes {
	s.HasLiked = &v
	return s
}

func (s *TalentV2QueryTagLikeListResponseBodyResultTagLikes) SetLikeCount(v int32) *TalentV2QueryTagLikeListResponseBodyResultTagLikes {
	s.LikeCount = &v
	return s
}

func (s *TalentV2QueryTagLikeListResponseBodyResultTagLikes) SetTagCode(v string) *TalentV2QueryTagLikeListResponseBodyResultTagLikes {
	s.TagCode = &v
	return s
}

func (s *TalentV2QueryTagLikeListResponseBodyResultTagLikes) SetTagName(v string) *TalentV2QueryTagLikeListResponseBodyResultTagLikes {
	s.TagName = &v
	return s
}

type TalentV2QueryTagLikeListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TalentV2QueryTagLikeListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TalentV2QueryTagLikeListResponse) String() string {
	return tea.Prettify(s)
}

func (s TalentV2QueryTagLikeListResponse) GoString() string {
	return s.String()
}

func (s *TalentV2QueryTagLikeListResponse) SetHeaders(v map[string]*string) *TalentV2QueryTagLikeListResponse {
	s.Headers = v
	return s
}

func (s *TalentV2QueryTagLikeListResponse) SetStatusCode(v int32) *TalentV2QueryTagLikeListResponse {
	s.StatusCode = &v
	return s
}

func (s *TalentV2QueryTagLikeListResponse) SetBody(v *TalentV2QueryTagLikeListResponseBody) *TalentV2QueryTagLikeListResponse {
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
// 人才标签：添加员工自定义标签
//
// @param request - TalentV2AddCustomTagRequest
//
// @param headers - TalentV2AddCustomTagHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TalentV2AddCustomTagResponse
func (client *Client) TalentV2AddCustomTagWithOptions(request *TalentV2AddCustomTagRequest, headers *TalentV2AddCustomTagHeaders, runtime *util.RuntimeOptions) (_result *TalentV2AddCustomTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		body["sortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.TagCode)) {
		body["tagCode"] = request.TagCode
	}

	if !tea.BoolValue(util.IsUnset(request.TagName)) {
		body["tagName"] = request.TagName
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
		Action:      tea.String("TalentV2AddCustomTag"),
		Version:     tea.String("talentTag_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/talentTag/talentTags/addCustomTag"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TalentV2AddCustomTagResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人才标签：添加员工自定义标签
//
// @param request - TalentV2AddCustomTagRequest
//
// @return TalentV2AddCustomTagResponse
func (client *Client) TalentV2AddCustomTag(request *TalentV2AddCustomTagRequest) (_result *TalentV2AddCustomTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TalentV2AddCustomTagHeaders{}
	_result = &TalentV2AddCustomTagResponse{}
	_body, _err := client.TalentV2AddCustomTagWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 人才标签：添加员工客观标签
//
// @param request - TalentV2AddObjectiveTagRequest
//
// @param headers - TalentV2AddObjectiveTagHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TalentV2AddObjectiveTagResponse
func (client *Client) TalentV2AddObjectiveTagWithOptions(request *TalentV2AddObjectiveTagRequest, headers *TalentV2AddObjectiveTagHeaders, runtime *util.RuntimeOptions) (_result *TalentV2AddObjectiveTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		body["sortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.TagCode)) {
		body["tagCode"] = request.TagCode
	}

	if !tea.BoolValue(util.IsUnset(request.TagName)) {
		body["tagName"] = request.TagName
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
		Action:      tea.String("TalentV2AddObjectiveTag"),
		Version:     tea.String("talentTag_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/talentTag/talentTags/addObjectiveTag"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TalentV2AddObjectiveTagResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人才标签：添加员工客观标签
//
// @param request - TalentV2AddObjectiveTagRequest
//
// @return TalentV2AddObjectiveTagResponse
func (client *Client) TalentV2AddObjectiveTag(request *TalentV2AddObjectiveTagRequest) (_result *TalentV2AddObjectiveTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TalentV2AddObjectiveTagHeaders{}
	_result = &TalentV2AddObjectiveTagResponse{}
	_body, _err := client.TalentV2AddObjectiveTagWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 人才标签：添加企业个性标签
//
// @param request - TalentV2AddPersonalityTagRequest
//
// @param headers - TalentV2AddPersonalityTagHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TalentV2AddPersonalityTagResponse
func (client *Client) TalentV2AddPersonalityTagWithOptions(request *TalentV2AddPersonalityTagRequest, headers *TalentV2AddPersonalityTagHeaders, runtime *util.RuntimeOptions) (_result *TalentV2AddPersonalityTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryCode)) {
		body["categoryCode"] = request.CategoryCode
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryName)) {
		body["categoryName"] = request.CategoryName
	}

	if !tea.BoolValue(util.IsUnset(request.CategorySortOrder)) {
		body["categorySortOrder"] = request.CategorySortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		body["sortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.TagCode)) {
		body["tagCode"] = request.TagCode
	}

	if !tea.BoolValue(util.IsUnset(request.TagName)) {
		body["tagName"] = request.TagName
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
		Action:      tea.String("TalentV2AddPersonalityTag"),
		Version:     tea.String("talentTag_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/talentTag/talentTags/addPersonalityTag"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TalentV2AddPersonalityTagResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人才标签：添加企业个性标签
//
// @param request - TalentV2AddPersonalityTagRequest
//
// @return TalentV2AddPersonalityTagResponse
func (client *Client) TalentV2AddPersonalityTag(request *TalentV2AddPersonalityTagRequest) (_result *TalentV2AddPersonalityTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TalentV2AddPersonalityTagHeaders{}
	_result = &TalentV2AddPersonalityTagResponse{}
	_body, _err := client.TalentV2AddPersonalityTagWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 人才标签：删除员工自定义标签并清除所有点赞记录
//
// @param request - TalentV2DeleteCustomTagRequest
//
// @param headers - TalentV2DeleteCustomTagHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TalentV2DeleteCustomTagResponse
func (client *Client) TalentV2DeleteCustomTagWithOptions(request *TalentV2DeleteCustomTagRequest, headers *TalentV2DeleteCustomTagHeaders, runtime *util.RuntimeOptions) (_result *TalentV2DeleteCustomTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TagCode)) {
		body["tagCode"] = request.TagCode
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
		Action:      tea.String("TalentV2DeleteCustomTag"),
		Version:     tea.String("talentTag_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/talentTag/talentTags/deleteCustomTagWithClearLike"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TalentV2DeleteCustomTagResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人才标签：删除员工自定义标签并清除所有点赞记录
//
// @param request - TalentV2DeleteCustomTagRequest
//
// @return TalentV2DeleteCustomTagResponse
func (client *Client) TalentV2DeleteCustomTag(request *TalentV2DeleteCustomTagRequest) (_result *TalentV2DeleteCustomTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TalentV2DeleteCustomTagHeaders{}
	_result = &TalentV2DeleteCustomTagResponse{}
	_body, _err := client.TalentV2DeleteCustomTagWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 人才标签：删除员工客观标签
//
// @param request - TalentV2DeleteObjectiveTagRequest
//
// @param headers - TalentV2DeleteObjectiveTagHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TalentV2DeleteObjectiveTagResponse
func (client *Client) TalentV2DeleteObjectiveTagWithOptions(request *TalentV2DeleteObjectiveTagRequest, headers *TalentV2DeleteObjectiveTagHeaders, runtime *util.RuntimeOptions) (_result *TalentV2DeleteObjectiveTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TagCode)) {
		body["tagCode"] = request.TagCode
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
		Action:      tea.String("TalentV2DeleteObjectiveTag"),
		Version:     tea.String("talentTag_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/talentTag/talentTags/deleteObjectiveTag"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TalentV2DeleteObjectiveTagResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人才标签：删除员工客观标签
//
// @param request - TalentV2DeleteObjectiveTagRequest
//
// @return TalentV2DeleteObjectiveTagResponse
func (client *Client) TalentV2DeleteObjectiveTag(request *TalentV2DeleteObjectiveTagRequest) (_result *TalentV2DeleteObjectiveTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TalentV2DeleteObjectiveTagHeaders{}
	_result = &TalentV2DeleteObjectiveTagResponse{}
	_body, _err := client.TalentV2DeleteObjectiveTagWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 人才标签：删除企业个性标签
//
// @param request - TalentV2DeletePersonalityTagRequest
//
// @param headers - TalentV2DeletePersonalityTagHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TalentV2DeletePersonalityTagResponse
func (client *Client) TalentV2DeletePersonalityTagWithOptions(request *TalentV2DeletePersonalityTagRequest, headers *TalentV2DeletePersonalityTagHeaders, runtime *util.RuntimeOptions) (_result *TalentV2DeletePersonalityTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TagCode)) {
		body["tagCode"] = request.TagCode
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
		Action:      tea.String("TalentV2DeletePersonalityTag"),
		Version:     tea.String("talentTag_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/talentTag/talentTags/deletePersonalityTag"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TalentV2DeletePersonalityTagResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人才标签：删除企业个性标签
//
// @param request - TalentV2DeletePersonalityTagRequest
//
// @return TalentV2DeletePersonalityTagResponse
func (client *Client) TalentV2DeletePersonalityTag(request *TalentV2DeletePersonalityTagRequest) (_result *TalentV2DeletePersonalityTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TalentV2DeletePersonalityTagHeaders{}
	_result = &TalentV2DeletePersonalityTagResponse{}
	_body, _err := client.TalentV2DeletePersonalityTagWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 人才标签：点赞/取消点赞标签
//
// @param request - TalentV2LikeTagRequest
//
// @param headers - TalentV2LikeTagHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TalentV2LikeTagResponse
func (client *Client) TalentV2LikeTagWithOptions(request *TalentV2LikeTagRequest, headers *TalentV2LikeTagHeaders, runtime *util.RuntimeOptions) (_result *TalentV2LikeTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionType)) {
		body["actionType"] = request.ActionType
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		body["operatorUserId"] = request.OperatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.TagCode)) {
		body["tagCode"] = request.TagCode
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
		Action:      tea.String("TalentV2LikeTag"),
		Version:     tea.String("talentTag_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/talentTag/talentTags/likeTag"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TalentV2LikeTagResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人才标签：点赞/取消点赞标签
//
// @param request - TalentV2LikeTagRequest
//
// @return TalentV2LikeTagResponse
func (client *Client) TalentV2LikeTag(request *TalentV2LikeTagRequest) (_result *TalentV2LikeTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TalentV2LikeTagHeaders{}
	_result = &TalentV2LikeTagResponse{}
	_body, _err := client.TalentV2LikeTagWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 人才标签：查询员工自定义标签
//
// @param request - TalentV2QueryCustomTagRequest
//
// @param headers - TalentV2QueryCustomTagHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TalentV2QueryCustomTagResponse
func (client *Client) TalentV2QueryCustomTagWithOptions(request *TalentV2QueryCustomTagRequest, headers *TalentV2QueryCustomTagHeaders, runtime *util.RuntimeOptions) (_result *TalentV2QueryCustomTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      tea.String("TalentV2QueryCustomTag"),
		Version:     tea.String("talentTag_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/talentTag/talentTags/queryCustomTag"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TalentV2QueryCustomTagResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人才标签：查询员工自定义标签
//
// @param request - TalentV2QueryCustomTagRequest
//
// @return TalentV2QueryCustomTagResponse
func (client *Client) TalentV2QueryCustomTag(request *TalentV2QueryCustomTagRequest) (_result *TalentV2QueryCustomTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TalentV2QueryCustomTagHeaders{}
	_result = &TalentV2QueryCustomTagResponse{}
	_body, _err := client.TalentV2QueryCustomTagWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 人才标签：查询员工客观标签
//
// @param request - TalentV2QueryObjectiveTagRequest
//
// @param headers - TalentV2QueryObjectiveTagHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TalentV2QueryObjectiveTagResponse
func (client *Client) TalentV2QueryObjectiveTagWithOptions(request *TalentV2QueryObjectiveTagRequest, headers *TalentV2QueryObjectiveTagHeaders, runtime *util.RuntimeOptions) (_result *TalentV2QueryObjectiveTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      tea.String("TalentV2QueryObjectiveTag"),
		Version:     tea.String("talentTag_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/talentTag/talentTags/queryObjectiveTag"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TalentV2QueryObjectiveTagResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人才标签：查询员工客观标签
//
// @param request - TalentV2QueryObjectiveTagRequest
//
// @return TalentV2QueryObjectiveTagResponse
func (client *Client) TalentV2QueryObjectiveTag(request *TalentV2QueryObjectiveTagRequest) (_result *TalentV2QueryObjectiveTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TalentV2QueryObjectiveTagHeaders{}
	_result = &TalentV2QueryObjectiveTagResponse{}
	_body, _err := client.TalentV2QueryObjectiveTagWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 人才标签：查询企业个性标签
//
// @param headers - TalentV2QueryPersonalityTagHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TalentV2QueryPersonalityTagResponse
func (client *Client) TalentV2QueryPersonalityTagWithOptions(headers *TalentV2QueryPersonalityTagHeaders, runtime *util.RuntimeOptions) (_result *TalentV2QueryPersonalityTagResponse, _err error) {
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
		Action:      tea.String("TalentV2QueryPersonalityTag"),
		Version:     tea.String("talentTag_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/talentTag/talentTags/queryPersonalityTag"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TalentV2QueryPersonalityTagResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人才标签：查询企业个性标签
//
// @return TalentV2QueryPersonalityTagResponse
func (client *Client) TalentV2QueryPersonalityTag() (_result *TalentV2QueryPersonalityTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TalentV2QueryPersonalityTagHeaders{}
	_result = &TalentV2QueryPersonalityTagResponse{}
	_body, _err := client.TalentV2QueryPersonalityTagWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 人才标签：分页查询指定标签的点赞记录
//
// @param request - TalentV2QueryTagLikeDetailListRequest
//
// @param headers - TalentV2QueryTagLikeDetailListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TalentV2QueryTagLikeDetailListResponse
func (client *Client) TalentV2QueryTagLikeDetailListWithOptions(request *TalentV2QueryTagLikeDetailListRequest, headers *TalentV2QueryTagLikeDetailListHeaders, runtime *util.RuntimeOptions) (_result *TalentV2QueryTagLikeDetailListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		query["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.TagCode)) {
		query["tagCode"] = request.TagCode
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
		Action:      tea.String("TalentV2QueryTagLikeDetailList"),
		Version:     tea.String("talentTag_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/talentTag/talentTags/queryTagLikeDetailList"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TalentV2QueryTagLikeDetailListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人才标签：分页查询指定标签的点赞记录
//
// @param request - TalentV2QueryTagLikeDetailListRequest
//
// @return TalentV2QueryTagLikeDetailListResponse
func (client *Client) TalentV2QueryTagLikeDetailList(request *TalentV2QueryTagLikeDetailListRequest) (_result *TalentV2QueryTagLikeDetailListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TalentV2QueryTagLikeDetailListHeaders{}
	_result = &TalentV2QueryTagLikeDetailListResponse{}
	_body, _err := client.TalentV2QueryTagLikeDetailListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 人才标签：查询点赞标签列表
//
// @param request - TalentV2QueryTagLikeListRequest
//
// @param headers - TalentV2QueryTagLikeListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TalentV2QueryTagLikeListResponse
func (client *Client) TalentV2QueryTagLikeListWithOptions(request *TalentV2QueryTagLikeListRequest, headers *TalentV2QueryTagLikeListHeaders, runtime *util.RuntimeOptions) (_result *TalentV2QueryTagLikeListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorUserId)) {
		query["operatorUserId"] = request.OperatorUserId
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
		Action:      tea.String("TalentV2QueryTagLikeList"),
		Version:     tea.String("talentTag_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/talentTag/talentTags/queryTagLikeList"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TalentV2QueryTagLikeListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人才标签：查询点赞标签列表
//
// @param request - TalentV2QueryTagLikeListRequest
//
// @return TalentV2QueryTagLikeListResponse
func (client *Client) TalentV2QueryTagLikeList(request *TalentV2QueryTagLikeListRequest) (_result *TalentV2QueryTagLikeListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TalentV2QueryTagLikeListHeaders{}
	_result = &TalentV2QueryTagLikeListResponse{}
	_body, _err := client.TalentV2QueryTagLikeListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
