// This file is auto-generated, don't edit it. Thanks.
package group_blackboard_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateGroupBlackboardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateGroupBlackboardHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupBlackboardHeaders) GoString() string {
	return s.String()
}

func (s *CreateGroupBlackboardHeaders) SetCommonHeaders(v map[string]*string) *CreateGroupBlackboardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateGroupBlackboardHeaders) SetXAcsDingtalkAccessToken(v string) *CreateGroupBlackboardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateGroupBlackboardRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 这是一条群公告
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cid123456
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// example:
	//
	// false
	SendDing *bool `json:"sendDing,omitempty" xml:"sendDing,omitempty"`
	// example:
	//
	// false
	Sticky *bool `json:"sticky,omitempty" xml:"sticky,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	UniqueId *string `json:"uniqueId,omitempty" xml:"uniqueId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateGroupBlackboardRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupBlackboardRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupBlackboardRequest) SetContent(v string) *CreateGroupBlackboardRequest {
	s.Content = &v
	return s
}

func (s *CreateGroupBlackboardRequest) SetOpenConversationId(v string) *CreateGroupBlackboardRequest {
	s.OpenConversationId = &v
	return s
}

func (s *CreateGroupBlackboardRequest) SetSendDing(v bool) *CreateGroupBlackboardRequest {
	s.SendDing = &v
	return s
}

func (s *CreateGroupBlackboardRequest) SetSticky(v bool) *CreateGroupBlackboardRequest {
	s.Sticky = &v
	return s
}

func (s *CreateGroupBlackboardRequest) SetUniqueId(v string) *CreateGroupBlackboardRequest {
	s.UniqueId = &v
	return s
}

func (s *CreateGroupBlackboardRequest) SetUserId(v string) *CreateGroupBlackboardRequest {
	s.UserId = &v
	return s
}

type CreateGroupBlackboardResponseBody struct {
	// example:
	//
	// 123456
	DataId *string `json:"dataId,omitempty" xml:"dataId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateGroupBlackboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupBlackboardResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupBlackboardResponseBody) SetDataId(v string) *CreateGroupBlackboardResponseBody {
	s.DataId = &v
	return s
}

func (s *CreateGroupBlackboardResponseBody) SetSuccess(v bool) *CreateGroupBlackboardResponseBody {
	s.Success = &v
	return s
}

type CreateGroupBlackboardResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupBlackboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGroupBlackboardResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupBlackboardResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupBlackboardResponse) SetHeaders(v map[string]*string) *CreateGroupBlackboardResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupBlackboardResponse) SetStatusCode(v int32) *CreateGroupBlackboardResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupBlackboardResponse) SetBody(v *CreateGroupBlackboardResponseBody) *CreateGroupBlackboardResponse {
	s.Body = v
	return s
}

type CreateGroupBlackboardNewHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateGroupBlackboardNewHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupBlackboardNewHeaders) GoString() string {
	return s.String()
}

func (s *CreateGroupBlackboardNewHeaders) SetCommonHeaders(v map[string]*string) *CreateGroupBlackboardNewHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateGroupBlackboardNewHeaders) SetXAcsDingtalkAccessToken(v string) *CreateGroupBlackboardNewHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateGroupBlackboardNewRequest struct {
	Content            *string `json:"content,omitempty" xml:"content,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	SendDing           *bool   `json:"sendDing,omitempty" xml:"sendDing,omitempty"`
	Sticky             *bool   `json:"sticky,omitempty" xml:"sticky,omitempty"`
	UniqueId           *string `json:"uniqueId,omitempty" xml:"uniqueId,omitempty"`
	UserId             *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateGroupBlackboardNewRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupBlackboardNewRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupBlackboardNewRequest) SetContent(v string) *CreateGroupBlackboardNewRequest {
	s.Content = &v
	return s
}

func (s *CreateGroupBlackboardNewRequest) SetOpenConversationId(v string) *CreateGroupBlackboardNewRequest {
	s.OpenConversationId = &v
	return s
}

func (s *CreateGroupBlackboardNewRequest) SetSendDing(v bool) *CreateGroupBlackboardNewRequest {
	s.SendDing = &v
	return s
}

func (s *CreateGroupBlackboardNewRequest) SetSticky(v bool) *CreateGroupBlackboardNewRequest {
	s.Sticky = &v
	return s
}

func (s *CreateGroupBlackboardNewRequest) SetUniqueId(v string) *CreateGroupBlackboardNewRequest {
	s.UniqueId = &v
	return s
}

func (s *CreateGroupBlackboardNewRequest) SetUserId(v string) *CreateGroupBlackboardNewRequest {
	s.UserId = &v
	return s
}

type CreateGroupBlackboardNewResponseBody struct {
	DataId  *string `json:"dataId,omitempty" xml:"dataId,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateGroupBlackboardNewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupBlackboardNewResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupBlackboardNewResponseBody) SetDataId(v string) *CreateGroupBlackboardNewResponseBody {
	s.DataId = &v
	return s
}

func (s *CreateGroupBlackboardNewResponseBody) SetSuccess(v bool) *CreateGroupBlackboardNewResponseBody {
	s.Success = &v
	return s
}

type CreateGroupBlackboardNewResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupBlackboardNewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGroupBlackboardNewResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupBlackboardNewResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupBlackboardNewResponse) SetHeaders(v map[string]*string) *CreateGroupBlackboardNewResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupBlackboardNewResponse) SetStatusCode(v int32) *CreateGroupBlackboardNewResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupBlackboardNewResponse) SetBody(v *CreateGroupBlackboardNewResponseBody) *CreateGroupBlackboardNewResponse {
	s.Body = v
	return s
}

type DeleteGroupBlackboardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteGroupBlackboardHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupBlackboardHeaders) GoString() string {
	return s.String()
}

func (s *DeleteGroupBlackboardHeaders) SetCommonHeaders(v map[string]*string) *DeleteGroupBlackboardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteGroupBlackboardHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteGroupBlackboardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteGroupBlackboardRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e3b4f5
	DataId *string `json:"dataId,omitempty" xml:"dataId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cid123456
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeleteGroupBlackboardRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupBlackboardRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupBlackboardRequest) SetDataId(v string) *DeleteGroupBlackboardRequest {
	s.DataId = &v
	return s
}

func (s *DeleteGroupBlackboardRequest) SetOpenConversationId(v string) *DeleteGroupBlackboardRequest {
	s.OpenConversationId = &v
	return s
}

func (s *DeleteGroupBlackboardRequest) SetUserId(v string) *DeleteGroupBlackboardRequest {
	s.UserId = &v
	return s
}

type DeleteGroupBlackboardResponseBody struct {
	// example:
	//
	// true
	IsDeleted *bool `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteGroupBlackboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupBlackboardResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupBlackboardResponseBody) SetIsDeleted(v bool) *DeleteGroupBlackboardResponseBody {
	s.IsDeleted = &v
	return s
}

func (s *DeleteGroupBlackboardResponseBody) SetSuccess(v bool) *DeleteGroupBlackboardResponseBody {
	s.Success = &v
	return s
}

type DeleteGroupBlackboardResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGroupBlackboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGroupBlackboardResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupBlackboardResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupBlackboardResponse) SetHeaders(v map[string]*string) *DeleteGroupBlackboardResponse {
	s.Headers = v
	return s
}

func (s *DeleteGroupBlackboardResponse) SetStatusCode(v int32) *DeleteGroupBlackboardResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGroupBlackboardResponse) SetBody(v *DeleteGroupBlackboardResponseBody) *DeleteGroupBlackboardResponse {
	s.Body = v
	return s
}

type DeleteGroupBlackboardNewHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteGroupBlackboardNewHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupBlackboardNewHeaders) GoString() string {
	return s.String()
}

func (s *DeleteGroupBlackboardNewHeaders) SetCommonHeaders(v map[string]*string) *DeleteGroupBlackboardNewHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteGroupBlackboardNewHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteGroupBlackboardNewHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteGroupBlackboardNewRequest struct {
	DataId             *string `json:"dataId,omitempty" xml:"dataId,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	UserId             *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeleteGroupBlackboardNewRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupBlackboardNewRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupBlackboardNewRequest) SetDataId(v string) *DeleteGroupBlackboardNewRequest {
	s.DataId = &v
	return s
}

func (s *DeleteGroupBlackboardNewRequest) SetOpenConversationId(v string) *DeleteGroupBlackboardNewRequest {
	s.OpenConversationId = &v
	return s
}

func (s *DeleteGroupBlackboardNewRequest) SetUserId(v string) *DeleteGroupBlackboardNewRequest {
	s.UserId = &v
	return s
}

type DeleteGroupBlackboardNewResponseBody struct {
	IsDeleted *bool `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	Success   *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteGroupBlackboardNewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupBlackboardNewResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupBlackboardNewResponseBody) SetIsDeleted(v bool) *DeleteGroupBlackboardNewResponseBody {
	s.IsDeleted = &v
	return s
}

func (s *DeleteGroupBlackboardNewResponseBody) SetSuccess(v bool) *DeleteGroupBlackboardNewResponseBody {
	s.Success = &v
	return s
}

type DeleteGroupBlackboardNewResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGroupBlackboardNewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGroupBlackboardNewResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupBlackboardNewResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupBlackboardNewResponse) SetHeaders(v map[string]*string) *DeleteGroupBlackboardNewResponse {
	s.Headers = v
	return s
}

func (s *DeleteGroupBlackboardNewResponse) SetStatusCode(v int32) *DeleteGroupBlackboardNewResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGroupBlackboardNewResponse) SetBody(v *DeleteGroupBlackboardNewResponseBody) *DeleteGroupBlackboardNewResponse {
	s.Body = v
	return s
}

type EditGroupBlackboardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s EditGroupBlackboardHeaders) String() string {
	return tea.Prettify(s)
}

func (s EditGroupBlackboardHeaders) GoString() string {
	return s.String()
}

func (s *EditGroupBlackboardHeaders) SetCommonHeaders(v map[string]*string) *EditGroupBlackboardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *EditGroupBlackboardHeaders) SetXAcsDingtalkAccessToken(v string) *EditGroupBlackboardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type EditGroupBlackboardRequest struct {
	Content            *string `json:"content,omitempty" xml:"content,omitempty"`
	DataId             *string `json:"dataId,omitempty" xml:"dataId,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	Sticky             *bool   `json:"sticky,omitempty" xml:"sticky,omitempty"`
	UserId             *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s EditGroupBlackboardRequest) String() string {
	return tea.Prettify(s)
}

func (s EditGroupBlackboardRequest) GoString() string {
	return s.String()
}

func (s *EditGroupBlackboardRequest) SetContent(v string) *EditGroupBlackboardRequest {
	s.Content = &v
	return s
}

func (s *EditGroupBlackboardRequest) SetDataId(v string) *EditGroupBlackboardRequest {
	s.DataId = &v
	return s
}

func (s *EditGroupBlackboardRequest) SetOpenConversationId(v string) *EditGroupBlackboardRequest {
	s.OpenConversationId = &v
	return s
}

func (s *EditGroupBlackboardRequest) SetSticky(v bool) *EditGroupBlackboardRequest {
	s.Sticky = &v
	return s
}

func (s *EditGroupBlackboardRequest) SetUserId(v string) *EditGroupBlackboardRequest {
	s.UserId = &v
	return s
}

type EditGroupBlackboardResponseBody struct {
	DataId  *string `json:"dataId,omitempty" xml:"dataId,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s EditGroupBlackboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditGroupBlackboardResponseBody) GoString() string {
	return s.String()
}

func (s *EditGroupBlackboardResponseBody) SetDataId(v string) *EditGroupBlackboardResponseBody {
	s.DataId = &v
	return s
}

func (s *EditGroupBlackboardResponseBody) SetSuccess(v bool) *EditGroupBlackboardResponseBody {
	s.Success = &v
	return s
}

type EditGroupBlackboardResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EditGroupBlackboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditGroupBlackboardResponse) String() string {
	return tea.Prettify(s)
}

func (s EditGroupBlackboardResponse) GoString() string {
	return s.String()
}

func (s *EditGroupBlackboardResponse) SetHeaders(v map[string]*string) *EditGroupBlackboardResponse {
	s.Headers = v
	return s
}

func (s *EditGroupBlackboardResponse) SetStatusCode(v int32) *EditGroupBlackboardResponse {
	s.StatusCode = &v
	return s
}

func (s *EditGroupBlackboardResponse) SetBody(v *EditGroupBlackboardResponseBody) *EditGroupBlackboardResponse {
	s.Body = v
	return s
}

type GetGroupBlackboardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetGroupBlackboardHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetGroupBlackboardHeaders) GoString() string {
	return s.String()
}

func (s *GetGroupBlackboardHeaders) SetCommonHeaders(v map[string]*string) *GetGroupBlackboardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetGroupBlackboardHeaders) SetXAcsDingtalkAccessToken(v string) *GetGroupBlackboardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetGroupBlackboardRequest struct {
	DataId             *string `json:"dataId,omitempty" xml:"dataId,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	UserId             *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetGroupBlackboardRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGroupBlackboardRequest) GoString() string {
	return s.String()
}

func (s *GetGroupBlackboardRequest) SetDataId(v string) *GetGroupBlackboardRequest {
	s.DataId = &v
	return s
}

func (s *GetGroupBlackboardRequest) SetOpenConversationId(v string) *GetGroupBlackboardRequest {
	s.OpenConversationId = &v
	return s
}

func (s *GetGroupBlackboardRequest) SetUserId(v string) *GetGroupBlackboardRequest {
	s.UserId = &v
	return s
}

type GetGroupBlackboardResponseBody struct {
	Result  *GetGroupBlackboardResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                 `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetGroupBlackboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGroupBlackboardResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupBlackboardResponseBody) SetResult(v *GetGroupBlackboardResponseBodyResult) *GetGroupBlackboardResponseBody {
	s.Result = v
	return s
}

func (s *GetGroupBlackboardResponseBody) SetSuccess(v bool) *GetGroupBlackboardResponseBody {
	s.Success = &v
	return s
}

type GetGroupBlackboardResponseBodyResult struct {
	Content     *string `json:"content,omitempty" xml:"content,omitempty"`
	DataId      *string `json:"dataId,omitempty" xml:"dataId,omitempty"`
	GmtCreate   *int64  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified *int64  `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	ReadCount   *int32  `json:"readCount,omitempty" xml:"readCount,omitempty"`
	Sticky      *bool   `json:"sticky,omitempty" xml:"sticky,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
	UserName    *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s GetGroupBlackboardResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetGroupBlackboardResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetGroupBlackboardResponseBodyResult) SetContent(v string) *GetGroupBlackboardResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetGroupBlackboardResponseBodyResult) SetDataId(v string) *GetGroupBlackboardResponseBodyResult {
	s.DataId = &v
	return s
}

func (s *GetGroupBlackboardResponseBodyResult) SetGmtCreate(v int64) *GetGroupBlackboardResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *GetGroupBlackboardResponseBodyResult) SetGmtModified(v int64) *GetGroupBlackboardResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *GetGroupBlackboardResponseBodyResult) SetReadCount(v int32) *GetGroupBlackboardResponseBodyResult {
	s.ReadCount = &v
	return s
}

func (s *GetGroupBlackboardResponseBodyResult) SetSticky(v bool) *GetGroupBlackboardResponseBodyResult {
	s.Sticky = &v
	return s
}

func (s *GetGroupBlackboardResponseBodyResult) SetTitle(v string) *GetGroupBlackboardResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetGroupBlackboardResponseBodyResult) SetUserId(v string) *GetGroupBlackboardResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *GetGroupBlackboardResponseBodyResult) SetUserName(v string) *GetGroupBlackboardResponseBodyResult {
	s.UserName = &v
	return s
}

type GetGroupBlackboardResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGroupBlackboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGroupBlackboardResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGroupBlackboardResponse) GoString() string {
	return s.String()
}

func (s *GetGroupBlackboardResponse) SetHeaders(v map[string]*string) *GetGroupBlackboardResponse {
	s.Headers = v
	return s
}

func (s *GetGroupBlackboardResponse) SetStatusCode(v int32) *GetGroupBlackboardResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGroupBlackboardResponse) SetBody(v *GetGroupBlackboardResponseBody) *GetGroupBlackboardResponse {
	s.Body = v
	return s
}

type ListGroupBlackboardHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListGroupBlackboardHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListGroupBlackboardHeaders) GoString() string {
	return s.String()
}

func (s *ListGroupBlackboardHeaders) SetCommonHeaders(v map[string]*string) *ListGroupBlackboardHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListGroupBlackboardHeaders) SetXAcsDingtalkAccessToken(v string) *ListGroupBlackboardHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListGroupBlackboardRequest struct {
	NextPageCursor     *string `json:"nextPageCursor,omitempty" xml:"nextPageCursor,omitempty"`
	OpenConversationId *string `json:"openConversationId,omitempty" xml:"openConversationId,omitempty"`
	PageSize           *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	UserId             *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListGroupBlackboardRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupBlackboardRequest) GoString() string {
	return s.String()
}

func (s *ListGroupBlackboardRequest) SetNextPageCursor(v string) *ListGroupBlackboardRequest {
	s.NextPageCursor = &v
	return s
}

func (s *ListGroupBlackboardRequest) SetOpenConversationId(v string) *ListGroupBlackboardRequest {
	s.OpenConversationId = &v
	return s
}

func (s *ListGroupBlackboardRequest) SetPageSize(v int32) *ListGroupBlackboardRequest {
	s.PageSize = &v
	return s
}

func (s *ListGroupBlackboardRequest) SetUserId(v string) *ListGroupBlackboardRequest {
	s.UserId = &v
	return s
}

type ListGroupBlackboardResponseBody struct {
	BlackboardList []*ListGroupBlackboardResponseBodyBlackboardList `json:"blackboardList,omitempty" xml:"blackboardList,omitempty" type:"Repeated"`
	HasMore        *bool                                            `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	NextPageCursor *string                                          `json:"nextPageCursor,omitempty" xml:"nextPageCursor,omitempty"`
	Success        *bool                                            `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListGroupBlackboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupBlackboardResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupBlackboardResponseBody) SetBlackboardList(v []*ListGroupBlackboardResponseBodyBlackboardList) *ListGroupBlackboardResponseBody {
	s.BlackboardList = v
	return s
}

func (s *ListGroupBlackboardResponseBody) SetHasMore(v bool) *ListGroupBlackboardResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListGroupBlackboardResponseBody) SetNextPageCursor(v string) *ListGroupBlackboardResponseBody {
	s.NextPageCursor = &v
	return s
}

func (s *ListGroupBlackboardResponseBody) SetSuccess(v bool) *ListGroupBlackboardResponseBody {
	s.Success = &v
	return s
}

type ListGroupBlackboardResponseBodyBlackboardList struct {
	Content     *string `json:"content,omitempty" xml:"content,omitempty"`
	DataId      *string `json:"dataId,omitempty" xml:"dataId,omitempty"`
	GmtCreate   *int64  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified *int64  `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	ReadCount   *int32  `json:"readCount,omitempty" xml:"readCount,omitempty"`
	Sticky      *bool   `json:"sticky,omitempty" xml:"sticky,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
	UserName    *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s ListGroupBlackboardResponseBodyBlackboardList) String() string {
	return tea.Prettify(s)
}

func (s ListGroupBlackboardResponseBodyBlackboardList) GoString() string {
	return s.String()
}

func (s *ListGroupBlackboardResponseBodyBlackboardList) SetContent(v string) *ListGroupBlackboardResponseBodyBlackboardList {
	s.Content = &v
	return s
}

func (s *ListGroupBlackboardResponseBodyBlackboardList) SetDataId(v string) *ListGroupBlackboardResponseBodyBlackboardList {
	s.DataId = &v
	return s
}

func (s *ListGroupBlackboardResponseBodyBlackboardList) SetGmtCreate(v int64) *ListGroupBlackboardResponseBodyBlackboardList {
	s.GmtCreate = &v
	return s
}

func (s *ListGroupBlackboardResponseBodyBlackboardList) SetGmtModified(v int64) *ListGroupBlackboardResponseBodyBlackboardList {
	s.GmtModified = &v
	return s
}

func (s *ListGroupBlackboardResponseBodyBlackboardList) SetReadCount(v int32) *ListGroupBlackboardResponseBodyBlackboardList {
	s.ReadCount = &v
	return s
}

func (s *ListGroupBlackboardResponseBodyBlackboardList) SetSticky(v bool) *ListGroupBlackboardResponseBodyBlackboardList {
	s.Sticky = &v
	return s
}

func (s *ListGroupBlackboardResponseBodyBlackboardList) SetTitle(v string) *ListGroupBlackboardResponseBodyBlackboardList {
	s.Title = &v
	return s
}

func (s *ListGroupBlackboardResponseBodyBlackboardList) SetUserId(v string) *ListGroupBlackboardResponseBodyBlackboardList {
	s.UserId = &v
	return s
}

func (s *ListGroupBlackboardResponseBodyBlackboardList) SetUserName(v string) *ListGroupBlackboardResponseBodyBlackboardList {
	s.UserName = &v
	return s
}

type ListGroupBlackboardResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupBlackboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupBlackboardResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupBlackboardResponse) GoString() string {
	return s.String()
}

func (s *ListGroupBlackboardResponse) SetHeaders(v map[string]*string) *ListGroupBlackboardResponse {
	s.Headers = v
	return s
}

func (s *ListGroupBlackboardResponse) SetStatusCode(v int32) *ListGroupBlackboardResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupBlackboardResponse) SetBody(v *ListGroupBlackboardResponseBody) *ListGroupBlackboardResponse {
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
// 创建群公告
//
// @param request - CreateGroupBlackboardRequest
//
// @param headers - CreateGroupBlackboardHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupBlackboardResponse
func (client *Client) CreateGroupBlackboardWithOptions(request *CreateGroupBlackboardRequest, headers *CreateGroupBlackboardHeaders, runtime *util.RuntimeOptions) (_result *CreateGroupBlackboardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.SendDing)) {
		body["sendDing"] = request.SendDing
	}

	if !tea.BoolValue(util.IsUnset(request.Sticky)) {
		body["sticky"] = request.Sticky
	}

	if !tea.BoolValue(util.IsUnset(request.UniqueId)) {
		body["uniqueId"] = request.UniqueId
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
		Action:      tea.String("CreateGroupBlackboard"),
		Version:     tea.String("groupBlackboard_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/groupBlackboard/blackboards"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGroupBlackboardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建群公告
//
// @param request - CreateGroupBlackboardRequest
//
// @return CreateGroupBlackboardResponse
func (client *Client) CreateGroupBlackboard(request *CreateGroupBlackboardRequest) (_result *CreateGroupBlackboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateGroupBlackboardHeaders{}
	_result = &CreateGroupBlackboardResponse{}
	_body, _err := client.CreateGroupBlackboardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建群公告
//
// @param request - CreateGroupBlackboardNewRequest
//
// @param headers - CreateGroupBlackboardNewHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupBlackboardNewResponse
func (client *Client) CreateGroupBlackboardNewWithOptions(request *CreateGroupBlackboardNewRequest, headers *CreateGroupBlackboardNewHeaders, runtime *util.RuntimeOptions) (_result *CreateGroupBlackboardNewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.SendDing)) {
		body["sendDing"] = request.SendDing
	}

	if !tea.BoolValue(util.IsUnset(request.Sticky)) {
		body["sticky"] = request.Sticky
	}

	if !tea.BoolValue(util.IsUnset(request.UniqueId)) {
		body["uniqueId"] = request.UniqueId
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
		Action:      tea.String("CreateGroupBlackboardNew"),
		Version:     tea.String("groupBlackboard_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/groupBlackboard/blackboards/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGroupBlackboardNewResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建群公告
//
// @param request - CreateGroupBlackboardNewRequest
//
// @return CreateGroupBlackboardNewResponse
func (client *Client) CreateGroupBlackboardNew(request *CreateGroupBlackboardNewRequest) (_result *CreateGroupBlackboardNewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateGroupBlackboardNewHeaders{}
	_result = &CreateGroupBlackboardNewResponse{}
	_body, _err := client.CreateGroupBlackboardNewWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除群公告
//
// @param request - DeleteGroupBlackboardRequest
//
// @param headers - DeleteGroupBlackboardHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupBlackboardResponse
func (client *Client) DeleteGroupBlackboardWithOptions(request *DeleteGroupBlackboardRequest, headers *DeleteGroupBlackboardHeaders, runtime *util.RuntimeOptions) (_result *DeleteGroupBlackboardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataId)) {
		body["dataId"] = request.DataId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
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
		Action:      tea.String("DeleteGroupBlackboard"),
		Version:     tea.String("groupBlackboard_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/groupBlackboard/blackboards/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGroupBlackboardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除群公告
//
// @param request - DeleteGroupBlackboardRequest
//
// @return DeleteGroupBlackboardResponse
func (client *Client) DeleteGroupBlackboard(request *DeleteGroupBlackboardRequest) (_result *DeleteGroupBlackboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteGroupBlackboardHeaders{}
	_result = &DeleteGroupBlackboardResponse{}
	_body, _err := client.DeleteGroupBlackboardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除群公告
//
// @param request - DeleteGroupBlackboardNewRequest
//
// @param headers - DeleteGroupBlackboardNewHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupBlackboardNewResponse
func (client *Client) DeleteGroupBlackboardNewWithOptions(request *DeleteGroupBlackboardNewRequest, headers *DeleteGroupBlackboardNewHeaders, runtime *util.RuntimeOptions) (_result *DeleteGroupBlackboardNewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataId)) {
		body["dataId"] = request.DataId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
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
		Action:      tea.String("DeleteGroupBlackboardNew"),
		Version:     tea.String("groupBlackboard_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/groupBlackboard/blackboards/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGroupBlackboardNewResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除群公告
//
// @param request - DeleteGroupBlackboardNewRequest
//
// @return DeleteGroupBlackboardNewResponse
func (client *Client) DeleteGroupBlackboardNew(request *DeleteGroupBlackboardNewRequest) (_result *DeleteGroupBlackboardNewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteGroupBlackboardNewHeaders{}
	_result = &DeleteGroupBlackboardNewResponse{}
	_body, _err := client.DeleteGroupBlackboardNewWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑群公告
//
// @param request - EditGroupBlackboardRequest
//
// @param headers - EditGroupBlackboardHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditGroupBlackboardResponse
func (client *Client) EditGroupBlackboardWithOptions(request *EditGroupBlackboardRequest, headers *EditGroupBlackboardHeaders, runtime *util.RuntimeOptions) (_result *EditGroupBlackboardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.DataId)) {
		body["dataId"] = request.DataId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.Sticky)) {
		body["sticky"] = request.Sticky
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
		Action:      tea.String("EditGroupBlackboard"),
		Version:     tea.String("groupBlackboard_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/groupBlackboard/blackboards/edit"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &EditGroupBlackboardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑群公告
//
// @param request - EditGroupBlackboardRequest
//
// @return EditGroupBlackboardResponse
func (client *Client) EditGroupBlackboard(request *EditGroupBlackboardRequest) (_result *EditGroupBlackboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &EditGroupBlackboardHeaders{}
	_result = &EditGroupBlackboardResponse{}
	_body, _err := client.EditGroupBlackboardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询群公告详情
//
// @param request - GetGroupBlackboardRequest
//
// @param headers - GetGroupBlackboardHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupBlackboardResponse
func (client *Client) GetGroupBlackboardWithOptions(request *GetGroupBlackboardRequest, headers *GetGroupBlackboardHeaders, runtime *util.RuntimeOptions) (_result *GetGroupBlackboardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataId)) {
		body["dataId"] = request.DataId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
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
		Action:      tea.String("GetGroupBlackboard"),
		Version:     tea.String("groupBlackboard_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/groupBlackboard/blackboards/get"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGroupBlackboardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询群公告详情
//
// @param request - GetGroupBlackboardRequest
//
// @return GetGroupBlackboardResponse
func (client *Client) GetGroupBlackboard(request *GetGroupBlackboardRequest) (_result *GetGroupBlackboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetGroupBlackboardHeaders{}
	_result = &GetGroupBlackboardResponse{}
	_body, _err := client.GetGroupBlackboardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询群公告列表
//
// @param request - ListGroupBlackboardRequest
//
// @param headers - ListGroupBlackboardHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupBlackboardResponse
func (client *Client) ListGroupBlackboardWithOptions(request *ListGroupBlackboardRequest, headers *ListGroupBlackboardHeaders, runtime *util.RuntimeOptions) (_result *ListGroupBlackboardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextPageCursor)) {
		body["nextPageCursor"] = request.NextPageCursor
	}

	if !tea.BoolValue(util.IsUnset(request.OpenConversationId)) {
		body["openConversationId"] = request.OpenConversationId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
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
		Action:      tea.String("ListGroupBlackboard"),
		Version:     tea.String("groupBlackboard_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/groupBlackboard/blackboards/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGroupBlackboardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询群公告列表
//
// @param request - ListGroupBlackboardRequest
//
// @return ListGroupBlackboardResponse
func (client *Client) ListGroupBlackboard(request *ListGroupBlackboardRequest) (_result *ListGroupBlackboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListGroupBlackboardHeaders{}
	_result = &ListGroupBlackboardResponse{}
	_body, _err := client.ListGroupBlackboardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
