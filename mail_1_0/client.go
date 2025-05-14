// This file is auto-generated, don't edit it. Thanks.
package mail_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type DraftMessage struct {
	// This parameter is required.
	BccRecipients []*Recipient `json:"bccRecipients,omitempty" xml:"bccRecipients,omitempty" type:"Repeated"`
	// This parameter is required.
	Body *DraftMessageBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// This parameter is required.
	CcRecipients []*Recipient `json:"ccRecipients,omitempty" xml:"ccRecipients,omitempty" type:"Repeated"`
	// This parameter is required.
	From *Recipient `json:"from,omitempty" xml:"from,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 由RFC5322定义的邮件头集合
	InternetMessageHeaders map[string]interface{} `json:"internetMessageHeaders,omitempty" xml:"internetMessageHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// uniqid@dingtalk.com
	InternetMessageId io.Reader `json:"internetMessageId,omitempty" xml:"internetMessageId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	IsReadReceiptRequested *bool `json:"isReadReceiptRequested,omitempty" xml:"isReadReceiptRequested,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PRY_NORMAL
	Priority io.Reader `json:"priority,omitempty" xml:"priority,omitempty"`
	// This parameter is required.
	ReplyTo *Recipient `json:"replyTo,omitempty" xml:"replyTo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 主题
	Subject io.Reader `json:"subject,omitempty" xml:"subject,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 一般取邮件正文的前一段
	Summary io.Reader `json:"summary,omitempty" xml:"summary,omitempty"`
	// This parameter is required.
	Tags []io.Reader `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// This parameter is required.
	ToRecipients []*Recipient `json:"toRecipients,omitempty" xml:"toRecipients,omitempty" type:"Repeated"`
}

func (s DraftMessage) String() string {
	return tea.Prettify(s)
}

func (s DraftMessage) GoString() string {
	return s.String()
}

func (s *DraftMessage) SetBccRecipients(v []*Recipient) *DraftMessage {
	s.BccRecipients = v
	return s
}

func (s *DraftMessage) SetBody(v *DraftMessageBody) *DraftMessage {
	s.Body = v
	return s
}

func (s *DraftMessage) SetCcRecipients(v []*Recipient) *DraftMessage {
	s.CcRecipients = v
	return s
}

func (s *DraftMessage) SetFrom(v *Recipient) *DraftMessage {
	s.From = v
	return s
}

func (s *DraftMessage) SetInternetMessageHeaders(v map[string]interface{}) *DraftMessage {
	s.InternetMessageHeaders = v
	return s
}

func (s *DraftMessage) SetInternetMessageId(v io.Reader) *DraftMessage {
	s.InternetMessageId = v
	return s
}

func (s *DraftMessage) SetIsReadReceiptRequested(v bool) *DraftMessage {
	s.IsReadReceiptRequested = &v
	return s
}

func (s *DraftMessage) SetPriority(v io.Reader) *DraftMessage {
	s.Priority = v
	return s
}

func (s *DraftMessage) SetReplyTo(v *Recipient) *DraftMessage {
	s.ReplyTo = v
	return s
}

func (s *DraftMessage) SetSubject(v io.Reader) *DraftMessage {
	s.Subject = v
	return s
}

func (s *DraftMessage) SetSummary(v io.Reader) *DraftMessage {
	s.Summary = v
	return s
}

func (s *DraftMessage) SetTags(v []io.Reader) *DraftMessage {
	s.Tags = v
	return s
}

func (s *DraftMessage) SetToRecipients(v []*Recipient) *DraftMessage {
	s.ToRecipients = v
	return s
}

type DraftMessageBody struct {
	// This parameter is required.
	BodyHtml io.Reader `json:"bodyHtml,omitempty" xml:"bodyHtml,omitempty"`
	// This parameter is required.
	BodyText io.Reader `json:"bodyText,omitempty" xml:"bodyText,omitempty"`
}

func (s DraftMessageBody) String() string {
	return tea.Prettify(s)
}

func (s DraftMessageBody) GoString() string {
	return s.String()
}

func (s *DraftMessageBody) SetBodyHtml(v io.Reader) *DraftMessageBody {
	s.BodyHtml = v
	return s
}

func (s *DraftMessageBody) SetBodyText(v io.Reader) *DraftMessageBody {
	s.BodyText = v
	return s
}

type Message struct {
	// This parameter is required.
	BccRecipients []*Recipient `json:"bccRecipients,omitempty" xml:"bccRecipients,omitempty" type:"Repeated"`
	// This parameter is required.
	CcRecipients []*Recipient `json:"ccRecipients,omitempty" xml:"ccRecipients,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// conversationid
	ConversationId io.Reader `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	FolderId io.Reader `json:"folderId,omitempty" xml:"folderId,omitempty"`
	// This parameter is required.
	From *Recipient `json:"from,omitempty" xml:"from,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	HasAttachments *bool `json:"hasAttachments,omitempty" xml:"hasAttachments,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mailid
	Id io.Reader `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 由RFC5322定义的邮件头集合
	InternetMessageHeaders map[string]interface{} `json:"internetMessageHeaders,omitempty" xml:"internetMessageHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// uniqid@dingtalk.com
	InternetMessageId io.Reader `json:"internetMessageId,omitempty" xml:"internetMessageId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	IsForwarded *bool `json:"isForwarded,omitempty" xml:"isForwarded,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	IsRead *bool `json:"isRead,omitempty" xml:"isRead,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	IsReadReceiptRequested *bool `json:"isReadReceiptRequested,omitempty" xml:"isReadReceiptRequested,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	IsReplied *bool `json:"isReplied,omitempty" xml:"isReplied,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-10-01T00:00:00Z
	LastModifiedDateTime io.Reader `json:"lastModifiedDateTime,omitempty" xml:"lastModifiedDateTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PRY_NORMAL
	Priority io.Reader `json:"priority,omitempty" xml:"priority,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-10-01T00:00:00Z
	ReceivedDateTime io.Reader `json:"receivedDateTime,omitempty" xml:"receivedDateTime,omitempty"`
	// This parameter is required.
	ReplyTo *Recipient `json:"replyTo,omitempty" xml:"replyTo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-10-01T00:00:00Z
	SentDateTime io.Reader `json:"sentDateTime,omitempty" xml:"sentDateTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 主题
	Subject io.Reader `json:"subject,omitempty" xml:"subject,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 一般取邮件正文的前一段
	Summary io.Reader `json:"summary,omitempty" xml:"summary,omitempty"`
	// This parameter is required.
	Tags []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// This parameter is required.
	ToRecipients []*Recipient `json:"toRecipients,omitempty" xml:"toRecipients,omitempty" type:"Repeated"`
}

func (s Message) String() string {
	return tea.Prettify(s)
}

func (s Message) GoString() string {
	return s.String()
}

func (s *Message) SetBccRecipients(v []*Recipient) *Message {
	s.BccRecipients = v
	return s
}

func (s *Message) SetCcRecipients(v []*Recipient) *Message {
	s.CcRecipients = v
	return s
}

func (s *Message) SetConversationId(v io.Reader) *Message {
	s.ConversationId = v
	return s
}

func (s *Message) SetFolderId(v io.Reader) *Message {
	s.FolderId = v
	return s
}

func (s *Message) SetFrom(v *Recipient) *Message {
	s.From = v
	return s
}

func (s *Message) SetHasAttachments(v bool) *Message {
	s.HasAttachments = &v
	return s
}

func (s *Message) SetId(v io.Reader) *Message {
	s.Id = v
	return s
}

func (s *Message) SetInternetMessageHeaders(v map[string]interface{}) *Message {
	s.InternetMessageHeaders = v
	return s
}

func (s *Message) SetInternetMessageId(v io.Reader) *Message {
	s.InternetMessageId = v
	return s
}

func (s *Message) SetIsForwarded(v bool) *Message {
	s.IsForwarded = &v
	return s
}

func (s *Message) SetIsRead(v bool) *Message {
	s.IsRead = &v
	return s
}

func (s *Message) SetIsReadReceiptRequested(v bool) *Message {
	s.IsReadReceiptRequested = &v
	return s
}

func (s *Message) SetIsReplied(v bool) *Message {
	s.IsReplied = &v
	return s
}

func (s *Message) SetLastModifiedDateTime(v io.Reader) *Message {
	s.LastModifiedDateTime = v
	return s
}

func (s *Message) SetPriority(v io.Reader) *Message {
	s.Priority = v
	return s
}

func (s *Message) SetReceivedDateTime(v io.Reader) *Message {
	s.ReceivedDateTime = v
	return s
}

func (s *Message) SetReplyTo(v *Recipient) *Message {
	s.ReplyTo = v
	return s
}

func (s *Message) SetSentDateTime(v io.Reader) *Message {
	s.SentDateTime = v
	return s
}

func (s *Message) SetSubject(v io.Reader) *Message {
	s.Subject = v
	return s
}

func (s *Message) SetSummary(v io.Reader) *Message {
	s.Summary = v
	return s
}

func (s *Message) SetTags(v []*string) *Message {
	s.Tags = v
	return s
}

func (s *Message) SetToRecipients(v []*Recipient) *Message {
	s.ToRecipients = v
	return s
}

type Recipient struct {
	// example:
	//
	// zhangsan@b.com
	Email io.Reader `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// ZhangSan
	Name io.Reader `json:"name,omitempty" xml:"name,omitempty"`
}

func (s Recipient) String() string {
	return tea.Prettify(s)
}

func (s Recipient) GoString() string {
	return s.String()
}

func (s *Recipient) SetEmail(v io.Reader) *Recipient {
	s.Email = v
	return s
}

func (s *Recipient) SetName(v io.Reader) *Recipient {
	s.Name = v
	return s
}

type CreateMailFolderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateMailFolderHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateMailFolderHeaders) GoString() string {
	return s.String()
}

func (s *CreateMailFolderHeaders) SetCommonHeaders(v map[string]*string) *CreateMailFolderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateMailFolderHeaders) SetXAcsDingtalkAccessToken(v string) *CreateMailFolderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateMailFolderRequest struct {
	// This parameter is required.
	DisplayName *string                `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Extensions  map[string]interface{} `json:"extensions,omitempty" xml:"extensions,omitempty"`
	FolerId     *string                `json:"folerId,omitempty" xml:"folerId,omitempty"`
}

func (s CreateMailFolderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMailFolderRequest) GoString() string {
	return s.String()
}

func (s *CreateMailFolderRequest) SetDisplayName(v string) *CreateMailFolderRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateMailFolderRequest) SetExtensions(v map[string]interface{}) *CreateMailFolderRequest {
	s.Extensions = v
	return s
}

func (s *CreateMailFolderRequest) SetFolerId(v string) *CreateMailFolderRequest {
	s.FolerId = &v
	return s
}

type CreateMailFolderResponseBody struct {
	Folder    *CreateMailFolderResponseBodyFolder `json:"folder,omitempty" xml:"folder,omitempty" type:"Struct"`
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateMailFolderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMailFolderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMailFolderResponseBody) SetFolder(v *CreateMailFolderResponseBodyFolder) *CreateMailFolderResponseBody {
	s.Folder = v
	return s
}

func (s *CreateMailFolderResponseBody) SetRequestId(v string) *CreateMailFolderResponseBody {
	s.RequestId = &v
	return s
}

type CreateMailFolderResponseBodyFolder struct {
	ChildFolderCount *int32                 `json:"childFolderCount,omitempty" xml:"childFolderCount,omitempty"`
	DisplayName      *string                `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Extensions       map[string]interface{} `json:"extensions,omitempty" xml:"extensions,omitempty"`
	Id               *string                `json:"id,omitempty" xml:"id,omitempty"`
	ParentFolderId   *string                `json:"parentFolderId,omitempty" xml:"parentFolderId,omitempty"`
	TotalItemCount   *int32                 `json:"totalItemCount,omitempty" xml:"totalItemCount,omitempty"`
	UnreadItemCount  *int32                 `json:"unreadItemCount,omitempty" xml:"unreadItemCount,omitempty"`
}

func (s CreateMailFolderResponseBodyFolder) String() string {
	return tea.Prettify(s)
}

func (s CreateMailFolderResponseBodyFolder) GoString() string {
	return s.String()
}

func (s *CreateMailFolderResponseBodyFolder) SetChildFolderCount(v int32) *CreateMailFolderResponseBodyFolder {
	s.ChildFolderCount = &v
	return s
}

func (s *CreateMailFolderResponseBodyFolder) SetDisplayName(v string) *CreateMailFolderResponseBodyFolder {
	s.DisplayName = &v
	return s
}

func (s *CreateMailFolderResponseBodyFolder) SetExtensions(v map[string]interface{}) *CreateMailFolderResponseBodyFolder {
	s.Extensions = v
	return s
}

func (s *CreateMailFolderResponseBodyFolder) SetId(v string) *CreateMailFolderResponseBodyFolder {
	s.Id = &v
	return s
}

func (s *CreateMailFolderResponseBodyFolder) SetParentFolderId(v string) *CreateMailFolderResponseBodyFolder {
	s.ParentFolderId = &v
	return s
}

func (s *CreateMailFolderResponseBodyFolder) SetTotalItemCount(v int32) *CreateMailFolderResponseBodyFolder {
	s.TotalItemCount = &v
	return s
}

func (s *CreateMailFolderResponseBodyFolder) SetUnreadItemCount(v int32) *CreateMailFolderResponseBodyFolder {
	s.UnreadItemCount = &v
	return s
}

type CreateMailFolderResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMailFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMailFolderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMailFolderResponse) GoString() string {
	return s.String()
}

func (s *CreateMailFolderResponse) SetHeaders(v map[string]*string) *CreateMailFolderResponse {
	s.Headers = v
	return s
}

func (s *CreateMailFolderResponse) SetStatusCode(v int32) *CreateMailFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMailFolderResponse) SetBody(v *CreateMailFolderResponseBody) *CreateMailFolderResponse {
	s.Body = v
	return s
}

type CreateMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateMessageHeaders) GoString() string {
	return s.String()
}

func (s *CreateMessageHeaders) SetCommonHeaders(v map[string]*string) *CreateMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateMessageHeaders) SetXAcsDingtalkAccessToken(v string) *CreateMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateMessageRequest struct {
	Message *DraftMessage `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CreateMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMessageRequest) GoString() string {
	return s.String()
}

func (s *CreateMessageRequest) SetMessage(v *DraftMessage) *CreateMessageRequest {
	s.Message = v
	return s
}

type CreateMessageResponseBody struct {
	Message   *Message `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string  `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMessageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMessageResponseBody) SetMessage(v *Message) *CreateMessageResponseBody {
	s.Message = v
	return s
}

func (s *CreateMessageResponseBody) SetRequestId(v string) *CreateMessageResponseBody {
	s.RequestId = &v
	return s
}

type CreateMessageResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMessageResponse) GoString() string {
	return s.String()
}

func (s *CreateMessageResponse) SetHeaders(v map[string]*string) *CreateMessageResponse {
	s.Headers = v
	return s
}

func (s *CreateMessageResponse) SetStatusCode(v int32) *CreateMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMessageResponse) SetBody(v *CreateMessageResponseBody) *CreateMessageResponse {
	s.Body = v
	return s
}

type CreateUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateUserHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateUserHeaders) GoString() string {
	return s.String()
}

func (s *CreateUserHeaders) SetCommonHeaders(v map[string]*string) *CreateUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateUserHeaders) SetXAcsDingtalkAccessToken(v string) *CreateUserHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateUserRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// user@yourcompany.org
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// This parameter is required.
	EmployeeType *string `json:"employeeType,omitempty" xml:"employeeType,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// password
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
}

func (s CreateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) SetEmail(v string) *CreateUserRequest {
	s.Email = &v
	return s
}

func (s *CreateUserRequest) SetEmployeeType(v string) *CreateUserRequest {
	s.EmployeeType = &v
	return s
}

func (s *CreateUserRequest) SetName(v string) *CreateUserRequest {
	s.Name = &v
	return s
}

func (s *CreateUserRequest) SetPassword(v string) *CreateUserRequest {
	s.Password = &v
	return s
}

type CreateUserResponseBody struct {
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
}

func (s CreateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBody) SetEmail(v string) *CreateUserResponseBody {
	s.Email = &v
	return s
}

type CreateUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponse) GoString() string {
	return s.String()
}

func (s *CreateUserResponse) SetHeaders(v map[string]*string) *CreateUserResponse {
	s.Headers = v
	return s
}

func (s *CreateUserResponse) SetStatusCode(v int32) *CreateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserResponse) SetBody(v *CreateUserResponseBody) *CreateUserResponse {
	s.Body = v
	return s
}

type DeleteMailFolderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteMailFolderHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteMailFolderHeaders) GoString() string {
	return s.String()
}

func (s *DeleteMailFolderHeaders) SetCommonHeaders(v map[string]*string) *DeleteMailFolderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteMailFolderHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteMailFolderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteMailFolderResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteMailFolderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMailFolderResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMailFolderResponseBody) SetRequestId(v string) *DeleteMailFolderResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMailFolderResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMailFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMailFolderResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMailFolderResponse) GoString() string {
	return s.String()
}

func (s *DeleteMailFolderResponse) SetHeaders(v map[string]*string) *DeleteMailFolderResponse {
	s.Headers = v
	return s
}

func (s *DeleteMailFolderResponse) SetStatusCode(v int32) *DeleteMailFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMailFolderResponse) SetBody(v *DeleteMailFolderResponseBody) *DeleteMailFolderResponse {
	s.Body = v
	return s
}

type DeleteMessagesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteMessagesHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteMessagesHeaders) GoString() string {
	return s.String()
}

func (s *DeleteMessagesHeaders) SetCommonHeaders(v map[string]*string) *DeleteMessagesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteMessagesHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteMessagesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteMessagesRequest struct {
	DeleteType *string `json:"deleteType,omitempty" xml:"deleteType,omitempty"`
	// This parameter is required.
	Ids []*string `json:"ids,omitempty" xml:"ids,omitempty" type:"Repeated"`
}

func (s DeleteMessagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMessagesRequest) GoString() string {
	return s.String()
}

func (s *DeleteMessagesRequest) SetDeleteType(v string) *DeleteMessagesRequest {
	s.DeleteType = &v
	return s
}

func (s *DeleteMessagesRequest) SetIds(v []*string) *DeleteMessagesRequest {
	s.Ids = v
	return s
}

type DeleteMessagesResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteMessagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMessagesResponseBody) SetRequestId(v string) *DeleteMessagesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMessagesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMessagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMessagesResponse) GoString() string {
	return s.String()
}

func (s *DeleteMessagesResponse) SetHeaders(v map[string]*string) *DeleteMessagesResponse {
	s.Headers = v
	return s
}

func (s *DeleteMessagesResponse) SetStatusCode(v int32) *DeleteMessagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMessagesResponse) SetBody(v *DeleteMessagesResponseBody) *DeleteMessagesResponse {
	s.Body = v
	return s
}

type GetMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetMessageHeaders) GoString() string {
	return s.String()
}

func (s *GetMessageHeaders) SetCommonHeaders(v map[string]*string) *GetMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMessageHeaders) SetXAcsDingtalkAccessToken(v string) *GetMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetMessageRequest struct {
	SelectFields *string `json:"selectFields,omitempty" xml:"selectFields,omitempty"`
}

func (s GetMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMessageRequest) GoString() string {
	return s.String()
}

func (s *GetMessageRequest) SetSelectFields(v string) *GetMessageRequest {
	s.SelectFields = &v
	return s
}

type GetMessageResponseBody struct {
	Message   *Message `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string  `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMessageResponseBody) GoString() string {
	return s.String()
}

func (s *GetMessageResponseBody) SetMessage(v *Message) *GetMessageResponseBody {
	s.Message = v
	return s
}

func (s *GetMessageResponseBody) SetRequestId(v string) *GetMessageResponseBody {
	s.RequestId = &v
	return s
}

type GetMessageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMessageResponse) GoString() string {
	return s.String()
}

func (s *GetMessageResponse) SetHeaders(v map[string]*string) *GetMessageResponse {
	s.Headers = v
	return s
}

func (s *GetMessageResponse) SetStatusCode(v int32) *GetMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMessageResponse) SetBody(v *GetMessageResponseBody) *GetMessageResponse {
	s.Body = v
	return s
}

type ListMailFoldersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListMailFoldersHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListMailFoldersHeaders) GoString() string {
	return s.String()
}

func (s *ListMailFoldersHeaders) SetCommonHeaders(v map[string]*string) *ListMailFoldersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListMailFoldersHeaders) SetXAcsDingtalkAccessToken(v string) *ListMailFoldersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListMailFoldersRequest struct {
	FolderId *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
}

func (s ListMailFoldersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMailFoldersRequest) GoString() string {
	return s.String()
}

func (s *ListMailFoldersRequest) SetFolderId(v string) *ListMailFoldersRequest {
	s.FolderId = &v
	return s
}

type ListMailFoldersResponseBody struct {
	Folders []*ListMailFoldersResponseBodyFolders `json:"folders,omitempty" xml:"folders,omitempty" type:"Repeated"`
}

func (s ListMailFoldersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMailFoldersResponseBody) GoString() string {
	return s.String()
}

func (s *ListMailFoldersResponseBody) SetFolders(v []*ListMailFoldersResponseBodyFolders) *ListMailFoldersResponseBody {
	s.Folders = v
	return s
}

type ListMailFoldersResponseBodyFolders struct {
	// This parameter is required.
	ChildFolderCount *int32 `json:"childFolderCount,omitempty" xml:"childFolderCount,omitempty"`
	// This parameter is required.
	DisplayName *string            `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Extensions  map[string]*string `json:"extensions,omitempty" xml:"extensions,omitempty"`
	// This parameter is required.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	ParentFolderId *string `json:"parentFolderId,omitempty" xml:"parentFolderId,omitempty"`
	// This parameter is required.
	TotalItemCount *int32 `json:"totalItemCount,omitempty" xml:"totalItemCount,omitempty"`
	// This parameter is required.
	UnreadItemCount *int32 `json:"unreadItemCount,omitempty" xml:"unreadItemCount,omitempty"`
}

func (s ListMailFoldersResponseBodyFolders) String() string {
	return tea.Prettify(s)
}

func (s ListMailFoldersResponseBodyFolders) GoString() string {
	return s.String()
}

func (s *ListMailFoldersResponseBodyFolders) SetChildFolderCount(v int32) *ListMailFoldersResponseBodyFolders {
	s.ChildFolderCount = &v
	return s
}

func (s *ListMailFoldersResponseBodyFolders) SetDisplayName(v string) *ListMailFoldersResponseBodyFolders {
	s.DisplayName = &v
	return s
}

func (s *ListMailFoldersResponseBodyFolders) SetExtensions(v map[string]*string) *ListMailFoldersResponseBodyFolders {
	s.Extensions = v
	return s
}

func (s *ListMailFoldersResponseBodyFolders) SetId(v string) *ListMailFoldersResponseBodyFolders {
	s.Id = &v
	return s
}

func (s *ListMailFoldersResponseBodyFolders) SetParentFolderId(v string) *ListMailFoldersResponseBodyFolders {
	s.ParentFolderId = &v
	return s
}

func (s *ListMailFoldersResponseBodyFolders) SetTotalItemCount(v int32) *ListMailFoldersResponseBodyFolders {
	s.TotalItemCount = &v
	return s
}

func (s *ListMailFoldersResponseBodyFolders) SetUnreadItemCount(v int32) *ListMailFoldersResponseBodyFolders {
	s.UnreadItemCount = &v
	return s
}

type ListMailFoldersResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMailFoldersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMailFoldersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMailFoldersResponse) GoString() string {
	return s.String()
}

func (s *ListMailFoldersResponse) SetHeaders(v map[string]*string) *ListMailFoldersResponse {
	s.Headers = v
	return s
}

func (s *ListMailFoldersResponse) SetStatusCode(v int32) *ListMailFoldersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMailFoldersResponse) SetBody(v *ListMailFoldersResponseBody) *ListMailFoldersResponse {
	s.Body = v
	return s
}

type ListMessagesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListMessagesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListMessagesHeaders) GoString() string {
	return s.String()
}

func (s *ListMessagesHeaders) SetCommonHeaders(v map[string]*string) *ListMessagesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListMessagesHeaders) SetXAcsDingtalkAccessToken(v string) *ListMessagesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListMessagesRequest struct {
	// This parameter is required.
	MaxResults   *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken    *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	SelectFields *string `json:"selectFields,omitempty" xml:"selectFields,omitempty"`
}

func (s ListMessagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMessagesRequest) GoString() string {
	return s.String()
}

func (s *ListMessagesRequest) SetMaxResults(v int32) *ListMessagesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMessagesRequest) SetNextToken(v string) *ListMessagesRequest {
	s.NextToken = &v
	return s
}

func (s *ListMessagesRequest) SetSelectFields(v string) *ListMessagesRequest {
	s.SelectFields = &v
	return s
}

type ListMessagesResponseBody struct {
	HasMore   *bool      `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	Messages  []*Message `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	NextToken *string    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId *string    `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMessagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMessagesResponseBody) SetHasMore(v bool) *ListMessagesResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListMessagesResponseBody) SetMessages(v []*Message) *ListMessagesResponseBody {
	s.Messages = v
	return s
}

func (s *ListMessagesResponseBody) SetNextToken(v string) *ListMessagesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMessagesResponseBody) SetRequestId(v string) *ListMessagesResponseBody {
	s.RequestId = &v
	return s
}

type ListMessagesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMessagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMessagesResponse) GoString() string {
	return s.String()
}

func (s *ListMessagesResponse) SetHeaders(v map[string]*string) *ListMessagesResponse {
	s.Headers = v
	return s
}

func (s *ListMessagesResponse) SetStatusCode(v int32) *ListMessagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMessagesResponse) SetBody(v *ListMessagesResponseBody) *ListMessagesResponse {
	s.Body = v
	return s
}

type MoveMailFolderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s MoveMailFolderHeaders) String() string {
	return tea.Prettify(s)
}

func (s MoveMailFolderHeaders) GoString() string {
	return s.String()
}

func (s *MoveMailFolderHeaders) SetCommonHeaders(v map[string]*string) *MoveMailFolderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MoveMailFolderHeaders) SetXAcsDingtalkAccessToken(v string) *MoveMailFolderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type MoveMailFolderRequest struct {
	DestinationFolderId *string `json:"destinationFolderId,omitempty" xml:"destinationFolderId,omitempty"`
}

func (s MoveMailFolderRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveMailFolderRequest) GoString() string {
	return s.String()
}

func (s *MoveMailFolderRequest) SetDestinationFolderId(v string) *MoveMailFolderRequest {
	s.DestinationFolderId = &v
	return s
}

type MoveMailFolderResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s MoveMailFolderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveMailFolderResponseBody) GoString() string {
	return s.String()
}

func (s *MoveMailFolderResponseBody) SetRequestId(v string) *MoveMailFolderResponseBody {
	s.RequestId = &v
	return s
}

type MoveMailFolderResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveMailFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveMailFolderResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveMailFolderResponse) GoString() string {
	return s.String()
}

func (s *MoveMailFolderResponse) SetHeaders(v map[string]*string) *MoveMailFolderResponse {
	s.Headers = v
	return s
}

func (s *MoveMailFolderResponse) SetStatusCode(v int32) *MoveMailFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveMailFolderResponse) SetBody(v *MoveMailFolderResponseBody) *MoveMailFolderResponse {
	s.Body = v
	return s
}

type SendMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SendMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s SendMessageHeaders) GoString() string {
	return s.String()
}

func (s *SendMessageHeaders) SetCommonHeaders(v map[string]*string) *SendMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendMessageHeaders) SetXAcsDingtalkAccessToken(v string) *SendMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SendMessageRequest struct {
	SaveToSentItems *bool `json:"saveToSentItems,omitempty" xml:"saveToSentItems,omitempty"`
}

func (s SendMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageRequest) GoString() string {
	return s.String()
}

func (s *SendMessageRequest) SetSaveToSentItems(v bool) *SendMessageRequest {
	s.SaveToSentItems = &v
	return s
}

type SendMessageResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SendMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageResponseBody) SetRequestId(v string) *SendMessageResponseBody {
	s.RequestId = &v
	return s
}

type SendMessageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponse) GoString() string {
	return s.String()
}

func (s *SendMessageResponse) SetHeaders(v map[string]*string) *SendMessageResponse {
	s.Headers = v
	return s
}

func (s *SendMessageResponse) SetStatusCode(v int32) *SendMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendMessageResponse) SetBody(v *SendMessageResponseBody) *SendMessageResponse {
	s.Body = v
	return s
}

type UpdateMailFolderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateMailFolderHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateMailFolderHeaders) GoString() string {
	return s.String()
}

func (s *UpdateMailFolderHeaders) SetCommonHeaders(v map[string]*string) *UpdateMailFolderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateMailFolderHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateMailFolderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateMailFolderRequest struct {
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s UpdateMailFolderRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMailFolderRequest) GoString() string {
	return s.String()
}

func (s *UpdateMailFolderRequest) SetDisplayName(v string) *UpdateMailFolderRequest {
	s.DisplayName = &v
	return s
}

type UpdateMailFolderResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateMailFolderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMailFolderResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMailFolderResponseBody) SetRequestId(v string) *UpdateMailFolderResponseBody {
	s.RequestId = &v
	return s
}

type UpdateMailFolderResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMailFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMailFolderResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMailFolderResponse) GoString() string {
	return s.String()
}

func (s *UpdateMailFolderResponse) SetHeaders(v map[string]*string) *UpdateMailFolderResponse {
	s.Headers = v
	return s
}

func (s *UpdateMailFolderResponse) SetStatusCode(v int32) *UpdateMailFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMailFolderResponse) SetBody(v *UpdateMailFolderResponseBody) *UpdateMailFolderResponse {
	s.Body = v
	return s
}

type UpdateMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateMessageHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateMessageHeaders) GoString() string {
	return s.String()
}

func (s *UpdateMessageHeaders) SetCommonHeaders(v map[string]*string) *UpdateMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateMessageHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateMessageHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateMessageRequest struct {
	Message *DraftMessage `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpdateMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMessageRequest) GoString() string {
	return s.String()
}

func (s *UpdateMessageRequest) SetMessage(v *DraftMessage) *UpdateMessageRequest {
	s.Message = v
	return s
}

type UpdateMessageResponseBody struct {
	Message   *Message `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string  `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMessageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMessageResponseBody) SetMessage(v *Message) *UpdateMessageResponseBody {
	s.Message = v
	return s
}

func (s *UpdateMessageResponseBody) SetRequestId(v string) *UpdateMessageResponseBody {
	s.RequestId = &v
	return s
}

type UpdateMessageResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMessageResponse) GoString() string {
	return s.String()
}

func (s *UpdateMessageResponse) SetHeaders(v map[string]*string) *UpdateMessageResponse {
	s.Headers = v
	return s
}

func (s *UpdateMessageResponse) SetStatusCode(v int32) *UpdateMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMessageResponse) SetBody(v *UpdateMessageResponseBody) *UpdateMessageResponse {
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
// 创建邮件文件夹
//
// @param request - CreateMailFolderRequest
//
// @param headers - CreateMailFolderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMailFolderResponse
func (client *Client) CreateMailFolderWithOptions(email *string, request *CreateMailFolderRequest, headers *CreateMailFolderHeaders, runtime *util.RuntimeOptions) (_result *CreateMailFolderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["displayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Extensions)) {
		body["extensions"] = request.Extensions
	}

	if !tea.BoolValue(util.IsUnset(request.FolerId)) {
		body["folerId"] = request.FolerId
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
		Action:      tea.String("CreateMailFolder"),
		Version:     tea.String("mail_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/mail/users/" + tea.StringValue(email) + "/mailFolders"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMailFolderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建邮件文件夹
//
// @param request - CreateMailFolderRequest
//
// @return CreateMailFolderResponse
func (client *Client) CreateMailFolder(email *string, request *CreateMailFolderRequest) (_result *CreateMailFolderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateMailFolderHeaders{}
	_result = &CreateMailFolderResponse{}
	_body, _err := client.CreateMailFolderWithOptions(email, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建草稿
//
// @param request - CreateMessageRequest
//
// @param headers - CreateMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMessageResponse
func (client *Client) CreateMessageWithOptions(email *string, request *CreateMessageRequest, headers *CreateMessageHeaders, runtime *util.RuntimeOptions) (_result *CreateMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Message)) {
		body["message"] = request.Message
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
		Action:      tea.String("CreateMessage"),
		Version:     tea.String("mail_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/mail/users/" + tea.StringValue(email) + "/messages"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建草稿
//
// @param request - CreateMessageRequest
//
// @return CreateMessageResponse
func (client *Client) CreateMessage(email *string, request *CreateMessageRequest) (_result *CreateMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateMessageHeaders{}
	_result = &CreateMessageResponse{}
	_body, _err := client.CreateMessageWithOptions(email, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建企业邮箱用户
//
// @param request - CreateUserRequest
//
// @param headers - CreateUserHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserResponse
func (client *Client) CreateUserWithOptions(request *CreateUserRequest, headers *CreateUserHeaders, runtime *util.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Email)) {
		body["email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.EmployeeType)) {
		body["employeeType"] = request.EmployeeType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["password"] = request.Password
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
		Action:      tea.String("CreateUser"),
		Version:     tea.String("mail_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/mail/users"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建企业邮箱用户
//
// @param request - CreateUserRequest
//
// @return CreateUserResponse
func (client *Client) CreateUser(request *CreateUserRequest) (_result *CreateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateUserHeaders{}
	_result = &CreateUserResponse{}
	_body, _err := client.CreateUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除文件夹
//
// @param headers - DeleteMailFolderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMailFolderResponse
func (client *Client) DeleteMailFolderWithOptions(email *string, id *string, headers *DeleteMailFolderHeaders, runtime *util.RuntimeOptions) (_result *DeleteMailFolderResponse, _err error) {
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
		Action:      tea.String("DeleteMailFolder"),
		Version:     tea.String("mail_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/mail/users/" + tea.StringValue(email) + "/mailFolders/" + tea.StringValue(id)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMailFolderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除文件夹
//
// @return DeleteMailFolderResponse
func (client *Client) DeleteMailFolder(email *string, id *string) (_result *DeleteMailFolderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteMailFolderHeaders{}
	_result = &DeleteMailFolderResponse{}
	_body, _err := client.DeleteMailFolderWithOptions(email, id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量删除邮件
//
// @param request - DeleteMessagesRequest
//
// @param headers - DeleteMessagesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMessagesResponse
func (client *Client) DeleteMessagesWithOptions(email *string, request *DeleteMessagesRequest, headers *DeleteMessagesHeaders, runtime *util.RuntimeOptions) (_result *DeleteMessagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeleteType)) {
		body["deleteType"] = request.DeleteType
	}

	if !tea.BoolValue(util.IsUnset(request.Ids)) {
		body["ids"] = request.Ids
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
		Action:      tea.String("DeleteMessages"),
		Version:     tea.String("mail_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/mail/users/" + tea.StringValue(email) + "/messages/batchDelete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMessagesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除邮件
//
// @param request - DeleteMessagesRequest
//
// @return DeleteMessagesResponse
func (client *Client) DeleteMessages(email *string, request *DeleteMessagesRequest) (_result *DeleteMessagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteMessagesHeaders{}
	_result = &DeleteMessagesResponse{}
	_body, _err := client.DeleteMessagesWithOptions(email, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取邮件元数据
//
// @param request - GetMessageRequest
//
// @param headers - GetMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMessageResponse
func (client *Client) GetMessageWithOptions(email *string, id *string, request *GetMessageRequest, headers *GetMessageHeaders, runtime *util.RuntimeOptions) (_result *GetMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SelectFields)) {
		query["selectFields"] = request.SelectFields
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
		Action:      tea.String("GetMessage"),
		Version:     tea.String("mail_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/mail/users/" + tea.StringValue(email) + "/messages/" + tea.StringValue(id)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取邮件元数据
//
// @param request - GetMessageRequest
//
// @return GetMessageResponse
func (client *Client) GetMessage(email *string, id *string, request *GetMessageRequest) (_result *GetMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetMessageHeaders{}
	_result = &GetMessageResponse{}
	_body, _err := client.GetMessageWithOptions(email, id, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定文件夹的子文件夹列表
//
// @param request - ListMailFoldersRequest
//
// @param headers - ListMailFoldersHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMailFoldersResponse
func (client *Client) ListMailFoldersWithOptions(email *string, request *ListMailFoldersRequest, headers *ListMailFoldersHeaders, runtime *util.RuntimeOptions) (_result *ListMailFoldersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FolderId)) {
		query["folderId"] = request.FolderId
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
		Action:      tea.String("ListMailFolders"),
		Version:     tea.String("mail_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/mail/users/" + tea.StringValue(email) + "/mailFolders"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMailFoldersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定文件夹的子文件夹列表
//
// @param request - ListMailFoldersRequest
//
// @return ListMailFoldersResponse
func (client *Client) ListMailFolders(email *string, request *ListMailFoldersRequest) (_result *ListMailFoldersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListMailFoldersHeaders{}
	_result = &ListMailFoldersResponse{}
	_body, _err := client.ListMailFoldersWithOptions(email, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取邮件列表
//
// @param request - ListMessagesRequest
//
// @param headers - ListMessagesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMessagesResponse
func (client *Client) ListMessagesWithOptions(email *string, folderId *string, request *ListMessagesRequest, headers *ListMessagesHeaders, runtime *util.RuntimeOptions) (_result *ListMessagesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SelectFields)) {
		query["selectFields"] = request.SelectFields
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
		Action:      tea.String("ListMessages"),
		Version:     tea.String("mail_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/mail/users/" + tea.StringValue(email) + "/mailFolders/" + tea.StringValue(folderId) + "/messages"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMessagesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取邮件列表
//
// @param request - ListMessagesRequest
//
// @return ListMessagesResponse
func (client *Client) ListMessages(email *string, folderId *string, request *ListMessagesRequest) (_result *ListMessagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListMessagesHeaders{}
	_result = &ListMessagesResponse{}
	_body, _err := client.ListMessagesWithOptions(email, folderId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移动文件夹
//
// @param request - MoveMailFolderRequest
//
// @param headers - MoveMailFolderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveMailFolderResponse
func (client *Client) MoveMailFolderWithOptions(email *string, id *string, request *MoveMailFolderRequest, headers *MoveMailFolderHeaders, runtime *util.RuntimeOptions) (_result *MoveMailFolderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationFolderId)) {
		body["destinationFolderId"] = request.DestinationFolderId
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
		Action:      tea.String("MoveMailFolder"),
		Version:     tea.String("mail_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/mail/users/" + tea.StringValue(email) + "/mailFolders/" + tea.StringValue(id) + "/move"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveMailFolderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移动文件夹
//
// @param request - MoveMailFolderRequest
//
// @return MoveMailFolderResponse
func (client *Client) MoveMailFolder(email *string, id *string, request *MoveMailFolderRequest) (_result *MoveMailFolderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &MoveMailFolderHeaders{}
	_result = &MoveMailFolderResponse{}
	_body, _err := client.MoveMailFolderWithOptions(email, id, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发送邮件
//
// @param request - SendMessageRequest
//
// @param headers - SendMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendMessageResponse
func (client *Client) SendMessageWithOptions(email *string, id *string, request *SendMessageRequest, headers *SendMessageHeaders, runtime *util.RuntimeOptions) (_result *SendMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SaveToSentItems)) {
		body["saveToSentItems"] = request.SaveToSentItems
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
		Action:      tea.String("SendMessage"),
		Version:     tea.String("mail_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/mail/users/" + tea.StringValue(email) + "/messages/" + tea.StringValue(id) + "/send"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SendMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送邮件
//
// @param request - SendMessageRequest
//
// @return SendMessageResponse
func (client *Client) SendMessage(email *string, id *string, request *SendMessageRequest) (_result *SendMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SendMessageHeaders{}
	_result = &SendMessageResponse{}
	_body, _err := client.SendMessageWithOptions(email, id, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改文件夹信息
//
// @param request - UpdateMailFolderRequest
//
// @param headers - UpdateMailFolderHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMailFolderResponse
func (client *Client) UpdateMailFolderWithOptions(email *string, id *string, request *UpdateMailFolderRequest, headers *UpdateMailFolderHeaders, runtime *util.RuntimeOptions) (_result *UpdateMailFolderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["displayName"] = request.DisplayName
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
		Action:      tea.String("UpdateMailFolder"),
		Version:     tea.String("mail_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/mail/users/" + tea.StringValue(email) + "/mailFolders/" + tea.StringValue(id) + "/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMailFolderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改文件夹信息
//
// @param request - UpdateMailFolderRequest
//
// @return UpdateMailFolderResponse
func (client *Client) UpdateMailFolder(email *string, id *string, request *UpdateMailFolderRequest) (_result *UpdateMailFolderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateMailFolderHeaders{}
	_result = &UpdateMailFolderResponse{}
	_body, _err := client.UpdateMailFolderWithOptions(email, id, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改草稿
//
// @param request - UpdateMessageRequest
//
// @param headers - UpdateMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMessageResponse
func (client *Client) UpdateMessageWithOptions(email *string, id *string, request *UpdateMessageRequest, headers *UpdateMessageHeaders, runtime *util.RuntimeOptions) (_result *UpdateMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Message)) {
		body["message"] = request.Message
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
		Action:      tea.String("UpdateMessage"),
		Version:     tea.String("mail_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/mail/users/" + tea.StringValue(email) + "/messages/" + tea.StringValue(id) + "/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMessageResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改草稿
//
// @param request - UpdateMessageRequest
//
// @return UpdateMessageResponse
func (client *Client) UpdateMessage(email *string, id *string, request *UpdateMessageRequest) (_result *UpdateMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateMessageHeaders{}
	_result = &UpdateMessageResponse{}
	_body, _err := client.UpdateMessageWithOptions(email, id, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
