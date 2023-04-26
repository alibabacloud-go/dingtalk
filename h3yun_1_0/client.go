// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package h3yun_1_0

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BatchInsertBizObjectHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchInsertBizObjectHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertBizObjectHeaders) GoString() string {
	return s.String()
}

func (s *BatchInsertBizObjectHeaders) SetCommonHeaders(v map[string]*string) *BatchInsertBizObjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchInsertBizObjectHeaders) SetXAcsDingtalkAccessToken(v string) *BatchInsertBizObjectHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchInsertBizObjectRequest struct {
	BizObjectJsonArray []*string `json:"bizObjectJsonArray,omitempty" xml:"bizObjectJsonArray,omitempty" type:"Repeated"`
	IsDraft            *bool     `json:"isDraft,omitempty" xml:"isDraft,omitempty"`
	OpUserId           *string   `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	SchemaCode         *string   `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
}

func (s BatchInsertBizObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertBizObjectRequest) GoString() string {
	return s.String()
}

func (s *BatchInsertBizObjectRequest) SetBizObjectJsonArray(v []*string) *BatchInsertBizObjectRequest {
	s.BizObjectJsonArray = v
	return s
}

func (s *BatchInsertBizObjectRequest) SetIsDraft(v bool) *BatchInsertBizObjectRequest {
	s.IsDraft = &v
	return s
}

func (s *BatchInsertBizObjectRequest) SetOpUserId(v string) *BatchInsertBizObjectRequest {
	s.OpUserId = &v
	return s
}

func (s *BatchInsertBizObjectRequest) SetSchemaCode(v string) *BatchInsertBizObjectRequest {
	s.SchemaCode = &v
	return s
}

type BatchInsertBizObjectResponseBody struct {
	Code    *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Data    *BatchInsertBizObjectResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                               `json:"message,omitempty" xml:"message,omitempty"`
}

func (s BatchInsertBizObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertBizObjectResponseBody) GoString() string {
	return s.String()
}

func (s *BatchInsertBizObjectResponseBody) SetCode(v string) *BatchInsertBizObjectResponseBody {
	s.Code = &v
	return s
}

func (s *BatchInsertBizObjectResponseBody) SetData(v *BatchInsertBizObjectResponseBodyData) *BatchInsertBizObjectResponseBody {
	s.Data = v
	return s
}

func (s *BatchInsertBizObjectResponseBody) SetMessage(v string) *BatchInsertBizObjectResponseBody {
	s.Message = &v
	return s
}

type BatchInsertBizObjectResponseBodyData struct {
	BizObjectIds   []*string `json:"bizObjectIds,omitempty" xml:"bizObjectIds,omitempty" type:"Repeated"`
	FailedDatas    []*string `json:"failedDatas,omitempty" xml:"failedDatas,omitempty" type:"Repeated"`
	FailedMessages []*string `json:"failedMessages,omitempty" xml:"failedMessages,omitempty" type:"Repeated"`
	ProcessIds     []*string `json:"processIds,omitempty" xml:"processIds,omitempty" type:"Repeated"`
}

func (s BatchInsertBizObjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertBizObjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchInsertBizObjectResponseBodyData) SetBizObjectIds(v []*string) *BatchInsertBizObjectResponseBodyData {
	s.BizObjectIds = v
	return s
}

func (s *BatchInsertBizObjectResponseBodyData) SetFailedDatas(v []*string) *BatchInsertBizObjectResponseBodyData {
	s.FailedDatas = v
	return s
}

func (s *BatchInsertBizObjectResponseBodyData) SetFailedMessages(v []*string) *BatchInsertBizObjectResponseBodyData {
	s.FailedMessages = v
	return s
}

func (s *BatchInsertBizObjectResponseBodyData) SetProcessIds(v []*string) *BatchInsertBizObjectResponseBodyData {
	s.ProcessIds = v
	return s
}

type BatchInsertBizObjectResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchInsertBizObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchInsertBizObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertBizObjectResponse) GoString() string {
	return s.String()
}

func (s *BatchInsertBizObjectResponse) SetHeaders(v map[string]*string) *BatchInsertBizObjectResponse {
	s.Headers = v
	return s
}

func (s *BatchInsertBizObjectResponse) SetStatusCode(v int32) *BatchInsertBizObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchInsertBizObjectResponse) SetBody(v *BatchInsertBizObjectResponseBody) *BatchInsertBizObjectResponse {
	s.Body = v
	return s
}

type CancelProcessInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CancelProcessInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s CancelProcessInstanceHeaders) GoString() string {
	return s.String()
}

func (s *CancelProcessInstanceHeaders) SetCommonHeaders(v map[string]*string) *CancelProcessInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CancelProcessInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *CancelProcessInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CancelProcessInstanceRequest struct {
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
}

func (s CancelProcessInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelProcessInstanceRequest) GoString() string {
	return s.String()
}

func (s *CancelProcessInstanceRequest) SetProcessInstanceId(v string) *CancelProcessInstanceRequest {
	s.ProcessInstanceId = &v
	return s
}

type CancelProcessInstanceResponseBody struct {
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CancelProcessInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelProcessInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CancelProcessInstanceResponseBody) SetCode(v string) *CancelProcessInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CancelProcessInstanceResponseBody) SetMessage(v string) *CancelProcessInstanceResponseBody {
	s.Message = &v
	return s
}

type CancelProcessInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelProcessInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelProcessInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelProcessInstanceResponse) GoString() string {
	return s.String()
}

func (s *CancelProcessInstanceResponse) SetHeaders(v map[string]*string) *CancelProcessInstanceResponse {
	s.Headers = v
	return s
}

func (s *CancelProcessInstanceResponse) SetStatusCode(v int32) *CancelProcessInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelProcessInstanceResponse) SetBody(v *CancelProcessInstanceResponseBody) *CancelProcessInstanceResponse {
	s.Body = v
	return s
}

type CreateBizObjectHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateBizObjectHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateBizObjectHeaders) GoString() string {
	return s.String()
}

func (s *CreateBizObjectHeaders) SetCommonHeaders(v map[string]*string) *CreateBizObjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateBizObjectHeaders) SetXAcsDingtalkAccessToken(v string) *CreateBizObjectHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateBizObjectRequest struct {
	BizObjectJson *string `json:"bizObjectJson,omitempty" xml:"bizObjectJson,omitempty"`
	IsDraft       *bool   `json:"isDraft,omitempty" xml:"isDraft,omitempty"`
	OpUserId      *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	SchemaCode    *string `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
}

func (s CreateBizObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBizObjectRequest) GoString() string {
	return s.String()
}

func (s *CreateBizObjectRequest) SetBizObjectJson(v string) *CreateBizObjectRequest {
	s.BizObjectJson = &v
	return s
}

func (s *CreateBizObjectRequest) SetIsDraft(v bool) *CreateBizObjectRequest {
	s.IsDraft = &v
	return s
}

func (s *CreateBizObjectRequest) SetOpUserId(v string) *CreateBizObjectRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateBizObjectRequest) SetSchemaCode(v string) *CreateBizObjectRequest {
	s.SchemaCode = &v
	return s
}

type CreateBizObjectResponseBody struct {
	Code    *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data    *CreateBizObjectResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                          `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CreateBizObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBizObjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBizObjectResponseBody) SetCode(v string) *CreateBizObjectResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBizObjectResponseBody) SetData(v *CreateBizObjectResponseBodyData) *CreateBizObjectResponseBody {
	s.Data = v
	return s
}

func (s *CreateBizObjectResponseBody) SetMessage(v string) *CreateBizObjectResponseBody {
	s.Message = &v
	return s
}

type CreateBizObjectResponseBodyData struct {
	BizObjectId       *string `json:"bizObjectId,omitempty" xml:"bizObjectId,omitempty"`
	FormUsageType     *string `json:"formUsageType,omitempty" xml:"formUsageType,omitempty"`
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	SchemaCode        *string `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
}

func (s CreateBizObjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateBizObjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateBizObjectResponseBodyData) SetBizObjectId(v string) *CreateBizObjectResponseBodyData {
	s.BizObjectId = &v
	return s
}

func (s *CreateBizObjectResponseBodyData) SetFormUsageType(v string) *CreateBizObjectResponseBodyData {
	s.FormUsageType = &v
	return s
}

func (s *CreateBizObjectResponseBodyData) SetProcessInstanceId(v string) *CreateBizObjectResponseBodyData {
	s.ProcessInstanceId = &v
	return s
}

func (s *CreateBizObjectResponseBodyData) SetSchemaCode(v string) *CreateBizObjectResponseBodyData {
	s.SchemaCode = &v
	return s
}

type CreateBizObjectResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateBizObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBizObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBizObjectResponse) GoString() string {
	return s.String()
}

func (s *CreateBizObjectResponse) SetHeaders(v map[string]*string) *CreateBizObjectResponse {
	s.Headers = v
	return s
}

func (s *CreateBizObjectResponse) SetStatusCode(v int32) *CreateBizObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBizObjectResponse) SetBody(v *CreateBizObjectResponseBody) *CreateBizObjectResponse {
	s.Body = v
	return s
}

type CreateProcessesInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateProcessesInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessesInstanceHeaders) GoString() string {
	return s.String()
}

func (s *CreateProcessesInstanceHeaders) SetCommonHeaders(v map[string]*string) *CreateProcessesInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateProcessesInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *CreateProcessesInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateProcessesInstanceRequest struct {
	BizObjectId *string `json:"bizObjectId,omitempty" xml:"bizObjectId,omitempty"`
	OpUserId    *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	SchemaCode  *string `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
}

func (s CreateProcessesInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessesInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateProcessesInstanceRequest) SetBizObjectId(v string) *CreateProcessesInstanceRequest {
	s.BizObjectId = &v
	return s
}

func (s *CreateProcessesInstanceRequest) SetOpUserId(v string) *CreateProcessesInstanceRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateProcessesInstanceRequest) SetSchemaCode(v string) *CreateProcessesInstanceRequest {
	s.SchemaCode = &v
	return s
}

type CreateProcessesInstanceResponseBody struct {
	Code    *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Data    *CreateProcessesInstanceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                                  `json:"message,omitempty" xml:"message,omitempty"`
}

func (s CreateProcessesInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessesInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProcessesInstanceResponseBody) SetCode(v string) *CreateProcessesInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateProcessesInstanceResponseBody) SetData(v *CreateProcessesInstanceResponseBodyData) *CreateProcessesInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreateProcessesInstanceResponseBody) SetMessage(v string) *CreateProcessesInstanceResponseBody {
	s.Message = &v
	return s
}

type CreateProcessesInstanceResponseBodyData struct {
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
}

func (s CreateProcessesInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessesInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateProcessesInstanceResponseBodyData) SetProcessInstanceId(v string) *CreateProcessesInstanceResponseBodyData {
	s.ProcessInstanceId = &v
	return s
}

type CreateProcessesInstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateProcessesInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProcessesInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProcessesInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateProcessesInstanceResponse) SetHeaders(v map[string]*string) *CreateProcessesInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateProcessesInstanceResponse) SetStatusCode(v int32) *CreateProcessesInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProcessesInstanceResponse) SetBody(v *CreateProcessesInstanceResponseBody) *CreateProcessesInstanceResponse {
	s.Body = v
	return s
}

type DeleteBizObjectHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteBizObjectHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteBizObjectHeaders) GoString() string {
	return s.String()
}

func (s *DeleteBizObjectHeaders) SetCommonHeaders(v map[string]*string) *DeleteBizObjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteBizObjectHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteBizObjectHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteBizObjectRequest struct {
	BizObjectId *string `json:"bizObjectId,omitempty" xml:"bizObjectId,omitempty"`
	SchemaCode  *string `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
}

func (s DeleteBizObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBizObjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteBizObjectRequest) SetBizObjectId(v string) *DeleteBizObjectRequest {
	s.BizObjectId = &v
	return s
}

func (s *DeleteBizObjectRequest) SetSchemaCode(v string) *DeleteBizObjectRequest {
	s.SchemaCode = &v
	return s
}

type DeleteBizObjectResponseBody struct {
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s DeleteBizObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBizObjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBizObjectResponseBody) SetCode(v string) *DeleteBizObjectResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteBizObjectResponseBody) SetMessage(v string) *DeleteBizObjectResponseBody {
	s.Message = &v
	return s
}

type DeleteBizObjectResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteBizObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBizObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBizObjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteBizObjectResponse) SetHeaders(v map[string]*string) *DeleteBizObjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteBizObjectResponse) SetStatusCode(v int32) *DeleteBizObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBizObjectResponse) SetBody(v *DeleteBizObjectResponseBody) *DeleteBizObjectResponse {
	s.Body = v
	return s
}

type DeleteProcessesInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteProcessesInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteProcessesInstanceHeaders) GoString() string {
	return s.String()
}

func (s *DeleteProcessesInstanceHeaders) SetCommonHeaders(v map[string]*string) *DeleteProcessesInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteProcessesInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteProcessesInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteProcessesInstanceRequest struct {
	IsAutoUpdateBizObject *bool   `json:"isAutoUpdateBizObject,omitempty" xml:"isAutoUpdateBizObject,omitempty"`
	ProcessInstanceId     *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
}

func (s DeleteProcessesInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProcessesInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteProcessesInstanceRequest) SetIsAutoUpdateBizObject(v bool) *DeleteProcessesInstanceRequest {
	s.IsAutoUpdateBizObject = &v
	return s
}

func (s *DeleteProcessesInstanceRequest) SetProcessInstanceId(v string) *DeleteProcessesInstanceRequest {
	s.ProcessInstanceId = &v
	return s
}

type DeleteProcessesInstanceResponseBody struct {
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s DeleteProcessesInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProcessesInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProcessesInstanceResponseBody) SetCode(v string) *DeleteProcessesInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteProcessesInstanceResponseBody) SetMessage(v string) *DeleteProcessesInstanceResponseBody {
	s.Message = &v
	return s
}

type DeleteProcessesInstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteProcessesInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProcessesInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProcessesInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteProcessesInstanceResponse) SetHeaders(v map[string]*string) *DeleteProcessesInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteProcessesInstanceResponse) SetStatusCode(v int32) *DeleteProcessesInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProcessesInstanceResponse) SetBody(v *DeleteProcessesInstanceResponseBody) *DeleteProcessesInstanceResponse {
	s.Body = v
	return s
}

type GetAppsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAppsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAppsHeaders) GoString() string {
	return s.String()
}

func (s *GetAppsHeaders) SetCommonHeaders(v map[string]*string) *GetAppsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAppsHeaders) SetXAcsDingtalkAccessToken(v string) *GetAppsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAppsRequest struct {
	QueryType *string   `json:"queryType,omitempty" xml:"queryType,omitempty"`
	Values    []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppsRequest) GoString() string {
	return s.String()
}

func (s *GetAppsRequest) SetQueryType(v string) *GetAppsRequest {
	s.QueryType = &v
	return s
}

func (s *GetAppsRequest) SetValues(v []*string) *GetAppsRequest {
	s.Values = v
	return s
}

type GetAppsResponseBody struct {
	Code    *string                    `json:"code,omitempty" xml:"code,omitempty"`
	Data    []*GetAppsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Message *string                    `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppsResponseBody) SetCode(v string) *GetAppsResponseBody {
	s.Code = &v
	return s
}

func (s *GetAppsResponseBody) SetData(v []*GetAppsResponseBodyData) *GetAppsResponseBody {
	s.Data = v
	return s
}

func (s *GetAppsResponseBody) SetMessage(v string) *GetAppsResponseBody {
	s.Message = &v
	return s
}

type GetAppsResponseBodyData struct {
	AppCode     *string `json:"appCode,omitempty" xml:"appCode,omitempty"`
	AppSource   *string `json:"appSource,omitempty" xml:"appSource,omitempty"`
	AppState    *string `json:"appState,omitempty" xml:"appState,omitempty"`
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Solution    *string `json:"solution,omitempty" xml:"solution,omitempty"`
}

func (s GetAppsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAppsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAppsResponseBodyData) SetAppCode(v string) *GetAppsResponseBodyData {
	s.AppCode = &v
	return s
}

func (s *GetAppsResponseBodyData) SetAppSource(v string) *GetAppsResponseBodyData {
	s.AppSource = &v
	return s
}

func (s *GetAppsResponseBodyData) SetAppState(v string) *GetAppsResponseBodyData {
	s.AppState = &v
	return s
}

func (s *GetAppsResponseBodyData) SetDisplayName(v string) *GetAppsResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *GetAppsResponseBodyData) SetSolution(v string) *GetAppsResponseBodyData {
	s.Solution = &v
	return s
}

type GetAppsResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppsResponse) GoString() string {
	return s.String()
}

func (s *GetAppsResponse) SetHeaders(v map[string]*string) *GetAppsResponse {
	s.Headers = v
	return s
}

func (s *GetAppsResponse) SetStatusCode(v int32) *GetAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppsResponse) SetBody(v *GetAppsResponseBody) *GetAppsResponse {
	s.Body = v
	return s
}

type GetAttachmentTemporaryUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAttachmentTemporaryUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAttachmentTemporaryUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetAttachmentTemporaryUrlHeaders) SetCommonHeaders(v map[string]*string) *GetAttachmentTemporaryUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAttachmentTemporaryUrlHeaders) SetXAcsDingtalkAccessToken(v string) *GetAttachmentTemporaryUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAttachmentTemporaryUrlRequest struct {
	AttachmentId *string `json:"attachmentId,omitempty" xml:"attachmentId,omitempty"`
}

func (s GetAttachmentTemporaryUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAttachmentTemporaryUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAttachmentTemporaryUrlRequest) SetAttachmentId(v string) *GetAttachmentTemporaryUrlRequest {
	s.AttachmentId = &v
	return s
}

type GetAttachmentTemporaryUrlResponseBody struct {
	Code    *string                                    `json:"code,omitempty" xml:"code,omitempty"`
	Data    *GetAttachmentTemporaryUrlResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                                    `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetAttachmentTemporaryUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAttachmentTemporaryUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAttachmentTemporaryUrlResponseBody) SetCode(v string) *GetAttachmentTemporaryUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetAttachmentTemporaryUrlResponseBody) SetData(v *GetAttachmentTemporaryUrlResponseBodyData) *GetAttachmentTemporaryUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetAttachmentTemporaryUrlResponseBody) SetMessage(v string) *GetAttachmentTemporaryUrlResponseBody {
	s.Message = &v
	return s
}

type GetAttachmentTemporaryUrlResponseBodyData struct {
	AttachmentUrl *string `json:"attachmentUrl,omitempty" xml:"attachmentUrl,omitempty"`
}

func (s GetAttachmentTemporaryUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAttachmentTemporaryUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAttachmentTemporaryUrlResponseBodyData) SetAttachmentUrl(v string) *GetAttachmentTemporaryUrlResponseBodyData {
	s.AttachmentUrl = &v
	return s
}

type GetAttachmentTemporaryUrlResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAttachmentTemporaryUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAttachmentTemporaryUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAttachmentTemporaryUrlResponse) GoString() string {
	return s.String()
}

func (s *GetAttachmentTemporaryUrlResponse) SetHeaders(v map[string]*string) *GetAttachmentTemporaryUrlResponse {
	s.Headers = v
	return s
}

func (s *GetAttachmentTemporaryUrlResponse) SetStatusCode(v int32) *GetAttachmentTemporaryUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAttachmentTemporaryUrlResponse) SetBody(v *GetAttachmentTemporaryUrlResponseBody) *GetAttachmentTemporaryUrlResponse {
	s.Body = v
	return s
}

type GetOrganizationsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOrganizationsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationsHeaders) GoString() string {
	return s.String()
}

func (s *GetOrganizationsHeaders) SetCommonHeaders(v map[string]*string) *GetOrganizationsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOrganizationsHeaders) SetXAcsDingtalkAccessToken(v string) *GetOrganizationsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOrganizationsRequest struct {
	DepartmentId *string `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
}

func (s GetOrganizationsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationsRequest) GoString() string {
	return s.String()
}

func (s *GetOrganizationsRequest) SetDepartmentId(v string) *GetOrganizationsRequest {
	s.DepartmentId = &v
	return s
}

type GetOrganizationsResponseBody struct {
	Code    *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data    []*GetOrganizationsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Message *string                             `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetOrganizationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationsResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrganizationsResponseBody) SetCode(v string) *GetOrganizationsResponseBody {
	s.Code = &v
	return s
}

func (s *GetOrganizationsResponseBody) SetData(v []*GetOrganizationsResponseBodyData) *GetOrganizationsResponseBody {
	s.Data = v
	return s
}

func (s *GetOrganizationsResponseBody) SetMessage(v string) *GetOrganizationsResponseBody {
	s.Message = &v
	return s
}

type GetOrganizationsResponseBodyData struct {
	Code        *string `json:"code,omitempty" xml:"code,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Id          *string `json:"id,omitempty" xml:"id,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	ParentId    *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	SortKey     *int64  `json:"sortKey,omitempty" xml:"sortKey,omitempty"`
	UnitType    *string `json:"unitType,omitempty" xml:"unitType,omitempty"`
}

func (s GetOrganizationsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOrganizationsResponseBodyData) SetCode(v string) *GetOrganizationsResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetOrganizationsResponseBodyData) SetDescription(v string) *GetOrganizationsResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetOrganizationsResponseBodyData) SetId(v string) *GetOrganizationsResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetOrganizationsResponseBodyData) SetName(v string) *GetOrganizationsResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetOrganizationsResponseBodyData) SetParentId(v string) *GetOrganizationsResponseBodyData {
	s.ParentId = &v
	return s
}

func (s *GetOrganizationsResponseBodyData) SetSortKey(v int64) *GetOrganizationsResponseBodyData {
	s.SortKey = &v
	return s
}

func (s *GetOrganizationsResponseBodyData) SetUnitType(v string) *GetOrganizationsResponseBodyData {
	s.UnitType = &v
	return s
}

type GetOrganizationsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOrganizationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOrganizationsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationsResponse) GoString() string {
	return s.String()
}

func (s *GetOrganizationsResponse) SetHeaders(v map[string]*string) *GetOrganizationsResponse {
	s.Headers = v
	return s
}

func (s *GetOrganizationsResponse) SetStatusCode(v int32) *GetOrganizationsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrganizationsResponse) SetBody(v *GetOrganizationsResponseBody) *GetOrganizationsResponse {
	s.Body = v
	return s
}

type GetRoleUsersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetRoleUsersHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetRoleUsersHeaders) GoString() string {
	return s.String()
}

func (s *GetRoleUsersHeaders) SetCommonHeaders(v map[string]*string) *GetRoleUsersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRoleUsersHeaders) SetXAcsDingtalkAccessToken(v string) *GetRoleUsersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetRoleUsersRequest struct {
	RoleId *string `json:"roleId,omitempty" xml:"roleId,omitempty"`
}

func (s GetRoleUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRoleUsersRequest) GoString() string {
	return s.String()
}

func (s *GetRoleUsersRequest) SetRoleId(v string) *GetRoleUsersRequest {
	s.RoleId = &v
	return s
}

type GetRoleUsersResponseBody struct {
	Code    *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data    []*GetRoleUsersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Message *string                         `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetRoleUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRoleUsersResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoleUsersResponseBody) SetCode(v string) *GetRoleUsersResponseBody {
	s.Code = &v
	return s
}

func (s *GetRoleUsersResponseBody) SetData(v []*GetRoleUsersResponseBodyData) *GetRoleUsersResponseBody {
	s.Data = v
	return s
}

func (s *GetRoleUsersResponseBody) SetMessage(v string) *GetRoleUsersResponseBody {
	s.Message = &v
	return s
}

type GetRoleUsersResponseBodyData struct {
	Code              *string   `json:"code,omitempty" xml:"code,omitempty"`
	DepartmentId      *string   `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	DepartmentName    *string   `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	Description       *string   `json:"description,omitempty" xml:"description,omitempty"`
	DomainType        *string   `json:"domainType,omitempty" xml:"domainType,omitempty"`
	Email             *string   `json:"email,omitempty" xml:"email,omitempty"`
	Mobile            *string   `json:"mobile,omitempty" xml:"mobile,omitempty"`
	Name              *string   `json:"name,omitempty" xml:"name,omitempty"`
	PartDepartmentIds []*string `json:"partDepartmentIds,omitempty" xml:"partDepartmentIds,omitempty" type:"Repeated"`
	Sex               *string   `json:"sex,omitempty" xml:"sex,omitempty"`
	SortKey           *int64    `json:"sortKey,omitempty" xml:"sortKey,omitempty"`
	UserId            *string   `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetRoleUsersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRoleUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRoleUsersResponseBodyData) SetCode(v string) *GetRoleUsersResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetRoleUsersResponseBodyData) SetDepartmentId(v string) *GetRoleUsersResponseBodyData {
	s.DepartmentId = &v
	return s
}

func (s *GetRoleUsersResponseBodyData) SetDepartmentName(v string) *GetRoleUsersResponseBodyData {
	s.DepartmentName = &v
	return s
}

func (s *GetRoleUsersResponseBodyData) SetDescription(v string) *GetRoleUsersResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetRoleUsersResponseBodyData) SetDomainType(v string) *GetRoleUsersResponseBodyData {
	s.DomainType = &v
	return s
}

func (s *GetRoleUsersResponseBodyData) SetEmail(v string) *GetRoleUsersResponseBodyData {
	s.Email = &v
	return s
}

func (s *GetRoleUsersResponseBodyData) SetMobile(v string) *GetRoleUsersResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *GetRoleUsersResponseBodyData) SetName(v string) *GetRoleUsersResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetRoleUsersResponseBodyData) SetPartDepartmentIds(v []*string) *GetRoleUsersResponseBodyData {
	s.PartDepartmentIds = v
	return s
}

func (s *GetRoleUsersResponseBodyData) SetSex(v string) *GetRoleUsersResponseBodyData {
	s.Sex = &v
	return s
}

func (s *GetRoleUsersResponseBodyData) SetSortKey(v int64) *GetRoleUsersResponseBodyData {
	s.SortKey = &v
	return s
}

func (s *GetRoleUsersResponseBodyData) SetUserId(v string) *GetRoleUsersResponseBodyData {
	s.UserId = &v
	return s
}

type GetRoleUsersResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRoleUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRoleUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRoleUsersResponse) GoString() string {
	return s.String()
}

func (s *GetRoleUsersResponse) SetHeaders(v map[string]*string) *GetRoleUsersResponse {
	s.Headers = v
	return s
}

func (s *GetRoleUsersResponse) SetStatusCode(v int32) *GetRoleUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRoleUsersResponse) SetBody(v *GetRoleUsersResponseBody) *GetRoleUsersResponse {
	s.Body = v
	return s
}

type GetRolesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetRolesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetRolesHeaders) GoString() string {
	return s.String()
}

func (s *GetRolesHeaders) SetCommonHeaders(v map[string]*string) *GetRolesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRolesHeaders) SetXAcsDingtalkAccessToken(v string) *GetRolesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetRolesResponseBody struct {
	Code    *string                   `json:"code,omitempty" xml:"code,omitempty"`
	Data    *GetRolesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                   `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRolesResponseBody) GoString() string {
	return s.String()
}

func (s *GetRolesResponseBody) SetCode(v string) *GetRolesResponseBody {
	s.Code = &v
	return s
}

func (s *GetRolesResponseBody) SetData(v *GetRolesResponseBodyData) *GetRolesResponseBody {
	s.Data = v
	return s
}

func (s *GetRolesResponseBody) SetMessage(v string) *GetRolesResponseBody {
	s.Message = &v
	return s
}

type GetRolesResponseBodyData struct {
	RoleGroups []*GetRolesResponseBodyDataRoleGroups `json:"roleGroups,omitempty" xml:"roleGroups,omitempty" type:"Repeated"`
	Roles      []*GetRolesResponseBodyDataRoles      `json:"roles,omitempty" xml:"roles,omitempty" type:"Repeated"`
}

func (s GetRolesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRolesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRolesResponseBodyData) SetRoleGroups(v []*GetRolesResponseBodyDataRoleGroups) *GetRolesResponseBodyData {
	s.RoleGroups = v
	return s
}

func (s *GetRolesResponseBodyData) SetRoles(v []*GetRolesResponseBodyDataRoles) *GetRolesResponseBodyData {
	s.Roles = v
	return s
}

type GetRolesResponseBodyDataRoleGroups struct {
	CompanyId   *string `json:"companyId,omitempty" xml:"companyId,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	GroupCode   *string `json:"groupCode,omitempty" xml:"groupCode,omitempty"`
	GroupId     *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	GroupName   *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	State       *string `json:"state,omitempty" xml:"state,omitempty"`
	Visibility  *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
}

func (s GetRolesResponseBodyDataRoleGroups) String() string {
	return tea.Prettify(s)
}

func (s GetRolesResponseBodyDataRoleGroups) GoString() string {
	return s.String()
}

func (s *GetRolesResponseBodyDataRoleGroups) SetCompanyId(v string) *GetRolesResponseBodyDataRoleGroups {
	s.CompanyId = &v
	return s
}

func (s *GetRolesResponseBodyDataRoleGroups) SetDescription(v string) *GetRolesResponseBodyDataRoleGroups {
	s.Description = &v
	return s
}

func (s *GetRolesResponseBodyDataRoleGroups) SetGroupCode(v string) *GetRolesResponseBodyDataRoleGroups {
	s.GroupCode = &v
	return s
}

func (s *GetRolesResponseBodyDataRoleGroups) SetGroupId(v string) *GetRolesResponseBodyDataRoleGroups {
	s.GroupId = &v
	return s
}

func (s *GetRolesResponseBodyDataRoleGroups) SetGroupName(v string) *GetRolesResponseBodyDataRoleGroups {
	s.GroupName = &v
	return s
}

func (s *GetRolesResponseBodyDataRoleGroups) SetState(v string) *GetRolesResponseBodyDataRoleGroups {
	s.State = &v
	return s
}

func (s *GetRolesResponseBodyDataRoleGroups) SetVisibility(v string) *GetRolesResponseBodyDataRoleGroups {
	s.Visibility = &v
	return s
}

type GetRolesResponseBodyDataRoles struct {
	CompanyId   *string `json:"companyId,omitempty" xml:"companyId,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	GroupId     *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	RoleCode    *string `json:"roleCode,omitempty" xml:"roleCode,omitempty"`
	RoleId      *string `json:"roleId,omitempty" xml:"roleId,omitempty"`
	RoleName    *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	State       *string `json:"state,omitempty" xml:"state,omitempty"`
	Visibility  *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
}

func (s GetRolesResponseBodyDataRoles) String() string {
	return tea.Prettify(s)
}

func (s GetRolesResponseBodyDataRoles) GoString() string {
	return s.String()
}

func (s *GetRolesResponseBodyDataRoles) SetCompanyId(v string) *GetRolesResponseBodyDataRoles {
	s.CompanyId = &v
	return s
}

func (s *GetRolesResponseBodyDataRoles) SetDescription(v string) *GetRolesResponseBodyDataRoles {
	s.Description = &v
	return s
}

func (s *GetRolesResponseBodyDataRoles) SetGroupId(v string) *GetRolesResponseBodyDataRoles {
	s.GroupId = &v
	return s
}

func (s *GetRolesResponseBodyDataRoles) SetRoleCode(v string) *GetRolesResponseBodyDataRoles {
	s.RoleCode = &v
	return s
}

func (s *GetRolesResponseBodyDataRoles) SetRoleId(v string) *GetRolesResponseBodyDataRoles {
	s.RoleId = &v
	return s
}

func (s *GetRolesResponseBodyDataRoles) SetRoleName(v string) *GetRolesResponseBodyDataRoles {
	s.RoleName = &v
	return s
}

func (s *GetRolesResponseBodyDataRoles) SetState(v string) *GetRolesResponseBodyDataRoles {
	s.State = &v
	return s
}

func (s *GetRolesResponseBodyDataRoles) SetVisibility(v string) *GetRolesResponseBodyDataRoles {
	s.Visibility = &v
	return s
}

type GetRolesResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRolesResponse) GoString() string {
	return s.String()
}

func (s *GetRolesResponse) SetHeaders(v map[string]*string) *GetRolesResponse {
	s.Headers = v
	return s
}

func (s *GetRolesResponse) SetStatusCode(v int32) *GetRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRolesResponse) SetBody(v *GetRolesResponseBody) *GetRolesResponse {
	s.Body = v
	return s
}

type GetUploadUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUploadUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUploadUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetUploadUrlHeaders) SetCommonHeaders(v map[string]*string) *GetUploadUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUploadUrlHeaders) SetXAcsDingtalkAccessToken(v string) *GetUploadUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUploadUrlRequest struct {
	BizObjectId *string `json:"bizObjectId,omitempty" xml:"bizObjectId,omitempty"`
	FieldName   *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	IsOverwrite *bool   `json:"isOverwrite,omitempty" xml:"isOverwrite,omitempty"`
	SchemaCode  *string `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
}

func (s GetUploadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUploadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetUploadUrlRequest) SetBizObjectId(v string) *GetUploadUrlRequest {
	s.BizObjectId = &v
	return s
}

func (s *GetUploadUrlRequest) SetFieldName(v string) *GetUploadUrlRequest {
	s.FieldName = &v
	return s
}

func (s *GetUploadUrlRequest) SetIsOverwrite(v bool) *GetUploadUrlRequest {
	s.IsOverwrite = &v
	return s
}

func (s *GetUploadUrlRequest) SetSchemaCode(v string) *GetUploadUrlRequest {
	s.SchemaCode = &v
	return s
}

type GetUploadUrlResponseBody struct {
	Code    *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data    *GetUploadUrlResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                       `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetUploadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUploadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetUploadUrlResponseBody) SetCode(v string) *GetUploadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetUploadUrlResponseBody) SetData(v *GetUploadUrlResponseBodyData) *GetUploadUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetUploadUrlResponseBody) SetMessage(v string) *GetUploadUrlResponseBody {
	s.Message = &v
	return s
}

type GetUploadUrlResponseBodyData struct {
	UploadUrl *string `json:"uploadUrl,omitempty" xml:"uploadUrl,omitempty"`
}

func (s GetUploadUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUploadUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUploadUrlResponseBodyData) SetUploadUrl(v string) *GetUploadUrlResponseBodyData {
	s.UploadUrl = &v
	return s
}

type GetUploadUrlResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUploadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUploadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUploadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetUploadUrlResponse) SetHeaders(v map[string]*string) *GetUploadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetUploadUrlResponse) SetStatusCode(v int32) *GetUploadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUploadUrlResponse) SetBody(v *GetUploadUrlResponseBody) *GetUploadUrlResponse {
	s.Body = v
	return s
}

type GetUsersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUsersHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUsersHeaders) GoString() string {
	return s.String()
}

func (s *GetUsersHeaders) SetCommonHeaders(v map[string]*string) *GetUsersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUsersHeaders) SetXAcsDingtalkAccessToken(v string) *GetUsersHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUsersRequest struct {
	DepartmentId *string `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	IsRecursive  *bool   `json:"isRecursive,omitempty" xml:"isRecursive,omitempty"`
}

func (s GetUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUsersRequest) GoString() string {
	return s.String()
}

func (s *GetUsersRequest) SetDepartmentId(v string) *GetUsersRequest {
	s.DepartmentId = &v
	return s
}

func (s *GetUsersRequest) SetIsRecursive(v bool) *GetUsersRequest {
	s.IsRecursive = &v
	return s
}

type GetUsersResponseBody struct {
	Code    *string                     `json:"code,omitempty" xml:"code,omitempty"`
	Data    []*GetUsersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Message *string                     `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUsersResponseBody) GoString() string {
	return s.String()
}

func (s *GetUsersResponseBody) SetCode(v string) *GetUsersResponseBody {
	s.Code = &v
	return s
}

func (s *GetUsersResponseBody) SetData(v []*GetUsersResponseBodyData) *GetUsersResponseBody {
	s.Data = v
	return s
}

func (s *GetUsersResponseBody) SetMessage(v string) *GetUsersResponseBody {
	s.Message = &v
	return s
}

type GetUsersResponseBodyData struct {
	Code              *string   `json:"code,omitempty" xml:"code,omitempty"`
	DepartmentId      *string   `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	DepartmentName    *string   `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	Description       *string   `json:"description,omitempty" xml:"description,omitempty"`
	DomainType        *string   `json:"domainType,omitempty" xml:"domainType,omitempty"`
	Email             *string   `json:"email,omitempty" xml:"email,omitempty"`
	Id                *string   `json:"id,omitempty" xml:"id,omitempty"`
	Mobile            *string   `json:"mobile,omitempty" xml:"mobile,omitempty"`
	Name              *string   `json:"name,omitempty" xml:"name,omitempty"`
	PartDepartmentIds []*string `json:"partDepartmentIds,omitempty" xml:"partDepartmentIds,omitempty" type:"Repeated"`
	Sex               *string   `json:"sex,omitempty" xml:"sex,omitempty"`
	SortKey           *int64    `json:"sortKey,omitempty" xml:"sortKey,omitempty"`
}

func (s GetUsersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUsersResponseBodyData) SetCode(v string) *GetUsersResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetUsersResponseBodyData) SetDepartmentId(v string) *GetUsersResponseBodyData {
	s.DepartmentId = &v
	return s
}

func (s *GetUsersResponseBodyData) SetDepartmentName(v string) *GetUsersResponseBodyData {
	s.DepartmentName = &v
	return s
}

func (s *GetUsersResponseBodyData) SetDescription(v string) *GetUsersResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetUsersResponseBodyData) SetDomainType(v string) *GetUsersResponseBodyData {
	s.DomainType = &v
	return s
}

func (s *GetUsersResponseBodyData) SetEmail(v string) *GetUsersResponseBodyData {
	s.Email = &v
	return s
}

func (s *GetUsersResponseBodyData) SetId(v string) *GetUsersResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetUsersResponseBodyData) SetMobile(v string) *GetUsersResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *GetUsersResponseBodyData) SetName(v string) *GetUsersResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetUsersResponseBodyData) SetPartDepartmentIds(v []*string) *GetUsersResponseBodyData {
	s.PartDepartmentIds = v
	return s
}

func (s *GetUsersResponseBodyData) SetSex(v string) *GetUsersResponseBodyData {
	s.Sex = &v
	return s
}

func (s *GetUsersResponseBodyData) SetSortKey(v int64) *GetUsersResponseBodyData {
	s.SortKey = &v
	return s
}

type GetUsersResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUsersResponse) GoString() string {
	return s.String()
}

func (s *GetUsersResponse) SetHeaders(v map[string]*string) *GetUsersResponse {
	s.Headers = v
	return s
}

func (s *GetUsersResponse) SetStatusCode(v int32) *GetUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUsersResponse) SetBody(v *GetUsersResponseBody) *GetUsersResponse {
	s.Body = v
	return s
}

type LoadBizFieldsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s LoadBizFieldsHeaders) String() string {
	return tea.Prettify(s)
}

func (s LoadBizFieldsHeaders) GoString() string {
	return s.String()
}

func (s *LoadBizFieldsHeaders) SetCommonHeaders(v map[string]*string) *LoadBizFieldsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *LoadBizFieldsHeaders) SetXAcsDingtalkAccessToken(v string) *LoadBizFieldsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type LoadBizFieldsRequest struct {
	SchemaCode *string `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
}

func (s LoadBizFieldsRequest) String() string {
	return tea.Prettify(s)
}

func (s LoadBizFieldsRequest) GoString() string {
	return s.String()
}

func (s *LoadBizFieldsRequest) SetSchemaCode(v string) *LoadBizFieldsRequest {
	s.SchemaCode = &v
	return s
}

type LoadBizFieldsResponseBody struct {
	Code    *string                        `json:"code,omitempty" xml:"code,omitempty"`
	Data    *LoadBizFieldsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                        `json:"message,omitempty" xml:"message,omitempty"`
}

func (s LoadBizFieldsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LoadBizFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *LoadBizFieldsResponseBody) SetCode(v string) *LoadBizFieldsResponseBody {
	s.Code = &v
	return s
}

func (s *LoadBizFieldsResponseBody) SetData(v *LoadBizFieldsResponseBodyData) *LoadBizFieldsResponseBody {
	s.Data = v
	return s
}

func (s *LoadBizFieldsResponseBody) SetMessage(v string) *LoadBizFieldsResponseBody {
	s.Message = &v
	return s
}

type LoadBizFieldsResponseBodyData struct {
	ChildForms []*LoadBizFieldsResponseBodyDataChildForms `json:"childForms,omitempty" xml:"childForms,omitempty" type:"Repeated"`
	Fields     []*LoadBizFieldsResponseBodyDataFields     `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	FormName   *string                                    `json:"formName,omitempty" xml:"formName,omitempty"`
	SchemaCode *string                                    `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
}

func (s LoadBizFieldsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s LoadBizFieldsResponseBodyData) GoString() string {
	return s.String()
}

func (s *LoadBizFieldsResponseBodyData) SetChildForms(v []*LoadBizFieldsResponseBodyDataChildForms) *LoadBizFieldsResponseBodyData {
	s.ChildForms = v
	return s
}

func (s *LoadBizFieldsResponseBodyData) SetFields(v []*LoadBizFieldsResponseBodyDataFields) *LoadBizFieldsResponseBodyData {
	s.Fields = v
	return s
}

func (s *LoadBizFieldsResponseBodyData) SetFormName(v string) *LoadBizFieldsResponseBodyData {
	s.FormName = &v
	return s
}

func (s *LoadBizFieldsResponseBodyData) SetSchemaCode(v string) *LoadBizFieldsResponseBodyData {
	s.SchemaCode = &v
	return s
}

type LoadBizFieldsResponseBodyDataChildForms struct {
	Fields     []*LoadBizFieldsResponseBodyDataChildFormsFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	FormName   *string                                          `json:"formName,omitempty" xml:"formName,omitempty"`
	SchemaCode *string                                          `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
}

func (s LoadBizFieldsResponseBodyDataChildForms) String() string {
	return tea.Prettify(s)
}

func (s LoadBizFieldsResponseBodyDataChildForms) GoString() string {
	return s.String()
}

func (s *LoadBizFieldsResponseBodyDataChildForms) SetFields(v []*LoadBizFieldsResponseBodyDataChildFormsFields) *LoadBizFieldsResponseBodyDataChildForms {
	s.Fields = v
	return s
}

func (s *LoadBizFieldsResponseBodyDataChildForms) SetFormName(v string) *LoadBizFieldsResponseBodyDataChildForms {
	s.FormName = &v
	return s
}

func (s *LoadBizFieldsResponseBodyDataChildForms) SetSchemaCode(v string) *LoadBizFieldsResponseBodyDataChildForms {
	s.SchemaCode = &v
	return s
}

type LoadBizFieldsResponseBodyDataChildFormsFields struct {
	BizDataType *string `json:"bizDataType,omitempty" xml:"bizDataType,omitempty"`
	FieldName   *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	Label       *string `json:"label,omitempty" xml:"label,omitempty"`
}

func (s LoadBizFieldsResponseBodyDataChildFormsFields) String() string {
	return tea.Prettify(s)
}

func (s LoadBizFieldsResponseBodyDataChildFormsFields) GoString() string {
	return s.String()
}

func (s *LoadBizFieldsResponseBodyDataChildFormsFields) SetBizDataType(v string) *LoadBizFieldsResponseBodyDataChildFormsFields {
	s.BizDataType = &v
	return s
}

func (s *LoadBizFieldsResponseBodyDataChildFormsFields) SetFieldName(v string) *LoadBizFieldsResponseBodyDataChildFormsFields {
	s.FieldName = &v
	return s
}

func (s *LoadBizFieldsResponseBodyDataChildFormsFields) SetLabel(v string) *LoadBizFieldsResponseBodyDataChildFormsFields {
	s.Label = &v
	return s
}

type LoadBizFieldsResponseBodyDataFields struct {
	BizDataType *string `json:"bizDataType,omitempty" xml:"bizDataType,omitempty"`
	FieldName   *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	Label       *string `json:"label,omitempty" xml:"label,omitempty"`
}

func (s LoadBizFieldsResponseBodyDataFields) String() string {
	return tea.Prettify(s)
}

func (s LoadBizFieldsResponseBodyDataFields) GoString() string {
	return s.String()
}

func (s *LoadBizFieldsResponseBodyDataFields) SetBizDataType(v string) *LoadBizFieldsResponseBodyDataFields {
	s.BizDataType = &v
	return s
}

func (s *LoadBizFieldsResponseBodyDataFields) SetFieldName(v string) *LoadBizFieldsResponseBodyDataFields {
	s.FieldName = &v
	return s
}

func (s *LoadBizFieldsResponseBodyDataFields) SetLabel(v string) *LoadBizFieldsResponseBodyDataFields {
	s.Label = &v
	return s
}

type LoadBizFieldsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LoadBizFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LoadBizFieldsResponse) String() string {
	return tea.Prettify(s)
}

func (s LoadBizFieldsResponse) GoString() string {
	return s.String()
}

func (s *LoadBizFieldsResponse) SetHeaders(v map[string]*string) *LoadBizFieldsResponse {
	s.Headers = v
	return s
}

func (s *LoadBizFieldsResponse) SetStatusCode(v int32) *LoadBizFieldsResponse {
	s.StatusCode = &v
	return s
}

func (s *LoadBizFieldsResponse) SetBody(v *LoadBizFieldsResponseBody) *LoadBizFieldsResponse {
	s.Body = v
	return s
}

type LoadBizObjectHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s LoadBizObjectHeaders) String() string {
	return tea.Prettify(s)
}

func (s LoadBizObjectHeaders) GoString() string {
	return s.String()
}

func (s *LoadBizObjectHeaders) SetCommonHeaders(v map[string]*string) *LoadBizObjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *LoadBizObjectHeaders) SetXAcsDingtalkAccessToken(v string) *LoadBizObjectHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type LoadBizObjectRequest struct {
	BizObjectId *string `json:"bizObjectId,omitempty" xml:"bizObjectId,omitempty"`
	SchemaCode  *string `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
}

func (s LoadBizObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s LoadBizObjectRequest) GoString() string {
	return s.String()
}

func (s *LoadBizObjectRequest) SetBizObjectId(v string) *LoadBizObjectRequest {
	s.BizObjectId = &v
	return s
}

func (s *LoadBizObjectRequest) SetSchemaCode(v string) *LoadBizObjectRequest {
	s.SchemaCode = &v
	return s
}

type LoadBizObjectResponseBody struct {
	Code    *string                `json:"code,omitempty" xml:"code,omitempty"`
	Data    map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	Message *string                `json:"message,omitempty" xml:"message,omitempty"`
}

func (s LoadBizObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LoadBizObjectResponseBody) GoString() string {
	return s.String()
}

func (s *LoadBizObjectResponseBody) SetCode(v string) *LoadBizObjectResponseBody {
	s.Code = &v
	return s
}

func (s *LoadBizObjectResponseBody) SetData(v map[string]interface{}) *LoadBizObjectResponseBody {
	s.Data = v
	return s
}

func (s *LoadBizObjectResponseBody) SetMessage(v string) *LoadBizObjectResponseBody {
	s.Message = &v
	return s
}

type LoadBizObjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LoadBizObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LoadBizObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s LoadBizObjectResponse) GoString() string {
	return s.String()
}

func (s *LoadBizObjectResponse) SetHeaders(v map[string]*string) *LoadBizObjectResponse {
	s.Headers = v
	return s
}

func (s *LoadBizObjectResponse) SetStatusCode(v int32) *LoadBizObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *LoadBizObjectResponse) SetBody(v *LoadBizObjectResponseBody) *LoadBizObjectResponse {
	s.Body = v
	return s
}

type LoadBizObjectsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s LoadBizObjectsHeaders) String() string {
	return tea.Prettify(s)
}

func (s LoadBizObjectsHeaders) GoString() string {
	return s.String()
}

func (s *LoadBizObjectsHeaders) SetCommonHeaders(v map[string]*string) *LoadBizObjectsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *LoadBizObjectsHeaders) SetXAcsDingtalkAccessToken(v string) *LoadBizObjectsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type LoadBizObjectsRequest struct {
	MatcherJson  *string                              `json:"matcherJson,omitempty" xml:"matcherJson,omitempty"`
	PageNumber   *int32                               `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize     *int32                               `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ReturnFields []*string                            `json:"returnFields,omitempty" xml:"returnFields,omitempty" type:"Repeated"`
	SchemaCode   *string                              `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
	SortByFields []*LoadBizObjectsRequestSortByFields `json:"sortByFields,omitempty" xml:"sortByFields,omitempty" type:"Repeated"`
}

func (s LoadBizObjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s LoadBizObjectsRequest) GoString() string {
	return s.String()
}

func (s *LoadBizObjectsRequest) SetMatcherJson(v string) *LoadBizObjectsRequest {
	s.MatcherJson = &v
	return s
}

func (s *LoadBizObjectsRequest) SetPageNumber(v int32) *LoadBizObjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *LoadBizObjectsRequest) SetPageSize(v int32) *LoadBizObjectsRequest {
	s.PageSize = &v
	return s
}

func (s *LoadBizObjectsRequest) SetReturnFields(v []*string) *LoadBizObjectsRequest {
	s.ReturnFields = v
	return s
}

func (s *LoadBizObjectsRequest) SetSchemaCode(v string) *LoadBizObjectsRequest {
	s.SchemaCode = &v
	return s
}

func (s *LoadBizObjectsRequest) SetSortByFields(v []*LoadBizObjectsRequestSortByFields) *LoadBizObjectsRequest {
	s.SortByFields = v
	return s
}

type LoadBizObjectsRequestSortByFields struct {
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
}

func (s LoadBizObjectsRequestSortByFields) String() string {
	return tea.Prettify(s)
}

func (s LoadBizObjectsRequestSortByFields) GoString() string {
	return s.String()
}

func (s *LoadBizObjectsRequestSortByFields) SetDirection(v string) *LoadBizObjectsRequestSortByFields {
	s.Direction = &v
	return s
}

func (s *LoadBizObjectsRequestSortByFields) SetFieldName(v string) *LoadBizObjectsRequestSortByFields {
	s.FieldName = &v
	return s
}

type LoadBizObjectsResponseBody struct {
	Code    *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data    *LoadBizObjectsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                         `json:"message,omitempty" xml:"message,omitempty"`
}

func (s LoadBizObjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LoadBizObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *LoadBizObjectsResponseBody) SetCode(v string) *LoadBizObjectsResponseBody {
	s.Code = &v
	return s
}

func (s *LoadBizObjectsResponseBody) SetData(v *LoadBizObjectsResponseBodyData) *LoadBizObjectsResponseBody {
	s.Data = v
	return s
}

func (s *LoadBizObjectsResponseBody) SetMessage(v string) *LoadBizObjectsResponseBody {
	s.Message = &v
	return s
}

type LoadBizObjectsResponseBodyData struct {
	BizObjects []map[string]interface{} `json:"bizObjects,omitempty" xml:"bizObjects,omitempty" type:"Repeated"`
	PageNumber *int32                   `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32                   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	TotalCount *int32                   `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s LoadBizObjectsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s LoadBizObjectsResponseBodyData) GoString() string {
	return s.String()
}

func (s *LoadBizObjectsResponseBodyData) SetBizObjects(v []map[string]interface{}) *LoadBizObjectsResponseBodyData {
	s.BizObjects = v
	return s
}

func (s *LoadBizObjectsResponseBodyData) SetPageNumber(v int32) *LoadBizObjectsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *LoadBizObjectsResponseBodyData) SetPageSize(v int32) *LoadBizObjectsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *LoadBizObjectsResponseBodyData) SetTotalCount(v int32) *LoadBizObjectsResponseBodyData {
	s.TotalCount = &v
	return s
}

type LoadBizObjectsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LoadBizObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LoadBizObjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s LoadBizObjectsResponse) GoString() string {
	return s.String()
}

func (s *LoadBizObjectsResponse) SetHeaders(v map[string]*string) *LoadBizObjectsResponse {
	s.Headers = v
	return s
}

func (s *LoadBizObjectsResponse) SetStatusCode(v int32) *LoadBizObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *LoadBizObjectsResponse) SetBody(v *LoadBizObjectsResponseBody) *LoadBizObjectsResponse {
	s.Body = v
	return s
}

type QueryAppFunctionNodesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAppFunctionNodesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAppFunctionNodesHeaders) GoString() string {
	return s.String()
}

func (s *QueryAppFunctionNodesHeaders) SetCommonHeaders(v map[string]*string) *QueryAppFunctionNodesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAppFunctionNodesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAppFunctionNodesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAppFunctionNodesRequest struct {
	AppCode *string `json:"appCode,omitempty" xml:"appCode,omitempty"`
}

func (s QueryAppFunctionNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAppFunctionNodesRequest) GoString() string {
	return s.String()
}

func (s *QueryAppFunctionNodesRequest) SetAppCode(v string) *QueryAppFunctionNodesRequest {
	s.AppCode = &v
	return s
}

type QueryAppFunctionNodesResponseBody struct {
	Code    *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Data    []*QueryAppFunctionNodesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Message *string                                  `json:"message,omitempty" xml:"message,omitempty"`
}

func (s QueryAppFunctionNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAppFunctionNodesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAppFunctionNodesResponseBody) SetCode(v string) *QueryAppFunctionNodesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAppFunctionNodesResponseBody) SetData(v []*QueryAppFunctionNodesResponseBodyData) *QueryAppFunctionNodesResponseBody {
	s.Data = v
	return s
}

func (s *QueryAppFunctionNodesResponseBody) SetMessage(v string) *QueryAppFunctionNodesResponseBody {
	s.Message = &v
	return s
}

type QueryAppFunctionNodesResponseBodyData struct {
	AppCode         *string `json:"appCode,omitempty" xml:"appCode,omitempty"`
	DisplayName     *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	IsSystem        *bool   `json:"isSystem,omitempty" xml:"isSystem,omitempty"`
	NodeType        *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
	NodeVisibleType *string `json:"nodeVisibleType,omitempty" xml:"nodeVisibleType,omitempty"`
	ParentCode      *string `json:"parentCode,omitempty" xml:"parentCode,omitempty"`
	SchemaCode      *string `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
	SortKey         *int64  `json:"sortKey,omitempty" xml:"sortKey,omitempty"`
	State           *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s QueryAppFunctionNodesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAppFunctionNodesResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAppFunctionNodesResponseBodyData) SetAppCode(v string) *QueryAppFunctionNodesResponseBodyData {
	s.AppCode = &v
	return s
}

func (s *QueryAppFunctionNodesResponseBodyData) SetDisplayName(v string) *QueryAppFunctionNodesResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *QueryAppFunctionNodesResponseBodyData) SetIsSystem(v bool) *QueryAppFunctionNodesResponseBodyData {
	s.IsSystem = &v
	return s
}

func (s *QueryAppFunctionNodesResponseBodyData) SetNodeType(v string) *QueryAppFunctionNodesResponseBodyData {
	s.NodeType = &v
	return s
}

func (s *QueryAppFunctionNodesResponseBodyData) SetNodeVisibleType(v string) *QueryAppFunctionNodesResponseBodyData {
	s.NodeVisibleType = &v
	return s
}

func (s *QueryAppFunctionNodesResponseBodyData) SetParentCode(v string) *QueryAppFunctionNodesResponseBodyData {
	s.ParentCode = &v
	return s
}

func (s *QueryAppFunctionNodesResponseBodyData) SetSchemaCode(v string) *QueryAppFunctionNodesResponseBodyData {
	s.SchemaCode = &v
	return s
}

func (s *QueryAppFunctionNodesResponseBodyData) SetSortKey(v int64) *QueryAppFunctionNodesResponseBodyData {
	s.SortKey = &v
	return s
}

func (s *QueryAppFunctionNodesResponseBodyData) SetState(v string) *QueryAppFunctionNodesResponseBodyData {
	s.State = &v
	return s
}

type QueryAppFunctionNodesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryAppFunctionNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAppFunctionNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAppFunctionNodesResponse) GoString() string {
	return s.String()
}

func (s *QueryAppFunctionNodesResponse) SetHeaders(v map[string]*string) *QueryAppFunctionNodesResponse {
	s.Headers = v
	return s
}

func (s *QueryAppFunctionNodesResponse) SetStatusCode(v int32) *QueryAppFunctionNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAppFunctionNodesResponse) SetBody(v *QueryAppFunctionNodesResponseBody) *QueryAppFunctionNodesResponse {
	s.Body = v
	return s
}

type QueryProcessesInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryProcessesInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesInstanceHeaders) GoString() string {
	return s.String()
}

func (s *QueryProcessesInstanceHeaders) SetCommonHeaders(v map[string]*string) *QueryProcessesInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryProcessesInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *QueryProcessesInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryProcessesInstanceRequest struct {
	BizObjectId *string `json:"bizObjectId,omitempty" xml:"bizObjectId,omitempty"`
	SchemaCode  *string `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
}

func (s QueryProcessesInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesInstanceRequest) GoString() string {
	return s.String()
}

func (s *QueryProcessesInstanceRequest) SetBizObjectId(v string) *QueryProcessesInstanceRequest {
	s.BizObjectId = &v
	return s
}

func (s *QueryProcessesInstanceRequest) SetSchemaCode(v string) *QueryProcessesInstanceRequest {
	s.SchemaCode = &v
	return s
}

type QueryProcessesInstanceResponseBody struct {
	Code    *string                                   `json:"code,omitempty" xml:"code,omitempty"`
	Data    []*QueryProcessesInstanceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Message *string                                   `json:"message,omitempty" xml:"message,omitempty"`
}

func (s QueryProcessesInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryProcessesInstanceResponseBody) SetCode(v string) *QueryProcessesInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *QueryProcessesInstanceResponseBody) SetData(v []*QueryProcessesInstanceResponseBodyData) *QueryProcessesInstanceResponseBody {
	s.Data = v
	return s
}

func (s *QueryProcessesInstanceResponseBody) SetMessage(v string) *QueryProcessesInstanceResponseBody {
	s.Message = &v
	return s
}

type QueryProcessesInstanceResponseBodyData struct {
	AppCode            *string                                           `json:"appCode,omitempty" xml:"appCode,omitempty"`
	BizObjectId        *string                                           `json:"bizObjectId,omitempty" xml:"bizObjectId,omitempty"`
	CreatedTimeGMT     *string                                           `json:"createdTimeGMT,omitempty" xml:"createdTimeGMT,omitempty"`
	DingTalkProcessId  *string                                           `json:"dingTalkProcessId,omitempty" xml:"dingTalkProcessId,omitempty"`
	FinishTimeGMT      *string                                           `json:"finishTimeGMT,omitempty" xml:"finishTimeGMT,omitempty"`
	Originator         *QueryProcessesInstanceResponseBodyDataOriginator `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	ProcessDisplayName *string                                           `json:"processDisplayName,omitempty" xml:"processDisplayName,omitempty"`
	ProcessInstanceId  *string                                           `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	ProcessVersion     *int32                                            `json:"processVersion,omitempty" xml:"processVersion,omitempty"`
	SchemaCode         *string                                           `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
	StartTimeGMT       *string                                           `json:"startTimeGMT,omitempty" xml:"startTimeGMT,omitempty"`
	State              *string                                           `json:"state,omitempty" xml:"state,omitempty"`
}

func (s QueryProcessesInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryProcessesInstanceResponseBodyData) SetAppCode(v string) *QueryProcessesInstanceResponseBodyData {
	s.AppCode = &v
	return s
}

func (s *QueryProcessesInstanceResponseBodyData) SetBizObjectId(v string) *QueryProcessesInstanceResponseBodyData {
	s.BizObjectId = &v
	return s
}

func (s *QueryProcessesInstanceResponseBodyData) SetCreatedTimeGMT(v string) *QueryProcessesInstanceResponseBodyData {
	s.CreatedTimeGMT = &v
	return s
}

func (s *QueryProcessesInstanceResponseBodyData) SetDingTalkProcessId(v string) *QueryProcessesInstanceResponseBodyData {
	s.DingTalkProcessId = &v
	return s
}

func (s *QueryProcessesInstanceResponseBodyData) SetFinishTimeGMT(v string) *QueryProcessesInstanceResponseBodyData {
	s.FinishTimeGMT = &v
	return s
}

func (s *QueryProcessesInstanceResponseBodyData) SetOriginator(v *QueryProcessesInstanceResponseBodyDataOriginator) *QueryProcessesInstanceResponseBodyData {
	s.Originator = v
	return s
}

func (s *QueryProcessesInstanceResponseBodyData) SetProcessDisplayName(v string) *QueryProcessesInstanceResponseBodyData {
	s.ProcessDisplayName = &v
	return s
}

func (s *QueryProcessesInstanceResponseBodyData) SetProcessInstanceId(v string) *QueryProcessesInstanceResponseBodyData {
	s.ProcessInstanceId = &v
	return s
}

func (s *QueryProcessesInstanceResponseBodyData) SetProcessVersion(v int32) *QueryProcessesInstanceResponseBodyData {
	s.ProcessVersion = &v
	return s
}

func (s *QueryProcessesInstanceResponseBodyData) SetSchemaCode(v string) *QueryProcessesInstanceResponseBodyData {
	s.SchemaCode = &v
	return s
}

func (s *QueryProcessesInstanceResponseBodyData) SetStartTimeGMT(v string) *QueryProcessesInstanceResponseBodyData {
	s.StartTimeGMT = &v
	return s
}

func (s *QueryProcessesInstanceResponseBodyData) SetState(v string) *QueryProcessesInstanceResponseBodyData {
	s.State = &v
	return s
}

type QueryProcessesInstanceResponseBodyDataOriginator struct {
	DepartmentId   *string `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId         *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryProcessesInstanceResponseBodyDataOriginator) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesInstanceResponseBodyDataOriginator) GoString() string {
	return s.String()
}

func (s *QueryProcessesInstanceResponseBodyDataOriginator) SetDepartmentId(v string) *QueryProcessesInstanceResponseBodyDataOriginator {
	s.DepartmentId = &v
	return s
}

func (s *QueryProcessesInstanceResponseBodyDataOriginator) SetDepartmentName(v string) *QueryProcessesInstanceResponseBodyDataOriginator {
	s.DepartmentName = &v
	return s
}

func (s *QueryProcessesInstanceResponseBodyDataOriginator) SetName(v string) *QueryProcessesInstanceResponseBodyDataOriginator {
	s.Name = &v
	return s
}

func (s *QueryProcessesInstanceResponseBodyDataOriginator) SetUserId(v string) *QueryProcessesInstanceResponseBodyDataOriginator {
	s.UserId = &v
	return s
}

type QueryProcessesInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryProcessesInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryProcessesInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesInstanceResponse) GoString() string {
	return s.String()
}

func (s *QueryProcessesInstanceResponse) SetHeaders(v map[string]*string) *QueryProcessesInstanceResponse {
	s.Headers = v
	return s
}

func (s *QueryProcessesInstanceResponse) SetStatusCode(v int32) *QueryProcessesInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryProcessesInstanceResponse) SetBody(v *QueryProcessesInstanceResponseBody) *QueryProcessesInstanceResponse {
	s.Body = v
	return s
}

type QueryProcessesWorkItemsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryProcessesWorkItemsHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesWorkItemsHeaders) GoString() string {
	return s.String()
}

func (s *QueryProcessesWorkItemsHeaders) SetCommonHeaders(v map[string]*string) *QueryProcessesWorkItemsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryProcessesWorkItemsHeaders) SetXAcsDingtalkAccessToken(v string) *QueryProcessesWorkItemsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryProcessesWorkItemsRequest struct {
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
}

func (s QueryProcessesWorkItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesWorkItemsRequest) GoString() string {
	return s.String()
}

func (s *QueryProcessesWorkItemsRequest) SetProcessInstanceId(v string) *QueryProcessesWorkItemsRequest {
	s.ProcessInstanceId = &v
	return s
}

type QueryProcessesWorkItemsResponseBody struct {
	Code    *string                                    `json:"code,omitempty" xml:"code,omitempty"`
	Data    []*QueryProcessesWorkItemsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Message *string                                    `json:"message,omitempty" xml:"message,omitempty"`
}

func (s QueryProcessesWorkItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesWorkItemsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryProcessesWorkItemsResponseBody) SetCode(v string) *QueryProcessesWorkItemsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBody) SetData(v []*QueryProcessesWorkItemsResponseBodyData) *QueryProcessesWorkItemsResponseBody {
	s.Data = v
	return s
}

func (s *QueryProcessesWorkItemsResponseBody) SetMessage(v string) *QueryProcessesWorkItemsResponseBody {
	s.Message = &v
	return s
}

type QueryProcessesWorkItemsResponseBodyData struct {
	ActivityCode      *string                                             `json:"activityCode,omitempty" xml:"activityCode,omitempty"`
	ActivityName      *string                                             `json:"activityName,omitempty" xml:"activityName,omitempty"`
	AppCode           *string                                             `json:"appCode,omitempty" xml:"appCode,omitempty"`
	BizObjectId       *string                                             `json:"bizObjectId,omitempty" xml:"bizObjectId,omitempty"`
	Comment           *string                                             `json:"comment,omitempty" xml:"comment,omitempty"`
	DisplayName       *string                                             `json:"displayName,omitempty" xml:"displayName,omitempty"`
	FinishTimeGMT     *string                                             `json:"finishTimeGMT,omitempty" xml:"finishTimeGMT,omitempty"`
	Finisher          *QueryProcessesWorkItemsResponseBodyDataFinisher    `json:"finisher,omitempty" xml:"finisher,omitempty" type:"Struct"`
	IsApproval        *bool                                               `json:"isApproval,omitempty" xml:"isApproval,omitempty"`
	IsFinish          *bool                                               `json:"isFinish,omitempty" xml:"isFinish,omitempty"`
	Participant       *QueryProcessesWorkItemsResponseBodyDataParticipant `json:"participant,omitempty" xml:"participant,omitempty" type:"Struct"`
	ProcessInstanceId *string                                             `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	ProcessVersion    *string                                             `json:"processVersion,omitempty" xml:"processVersion,omitempty"`
	Receiptor         *QueryProcessesWorkItemsResponseBodyDataReceiptor   `json:"receiptor,omitempty" xml:"receiptor,omitempty" type:"Struct"`
	ReceiveTimeGMT    *string                                             `json:"receiveTimeGMT,omitempty" xml:"receiveTimeGMT,omitempty"`
	SchemaCode        *string                                             `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
	StartTimeGMT      *string                                             `json:"startTimeGMT,omitempty" xml:"startTimeGMT,omitempty"`
	State             *string                                             `json:"state,omitempty" xml:"state,omitempty"`
	WorkItemId        *string                                             `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
	WorkItemType      *string                                             `json:"workItemType,omitempty" xml:"workItemType,omitempty"`
}

func (s QueryProcessesWorkItemsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesWorkItemsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetActivityCode(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.ActivityCode = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetActivityName(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.ActivityName = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetAppCode(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.AppCode = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetBizObjectId(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.BizObjectId = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetComment(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.Comment = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetDisplayName(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetFinishTimeGMT(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.FinishTimeGMT = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetFinisher(v *QueryProcessesWorkItemsResponseBodyDataFinisher) *QueryProcessesWorkItemsResponseBodyData {
	s.Finisher = v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetIsApproval(v bool) *QueryProcessesWorkItemsResponseBodyData {
	s.IsApproval = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetIsFinish(v bool) *QueryProcessesWorkItemsResponseBodyData {
	s.IsFinish = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetParticipant(v *QueryProcessesWorkItemsResponseBodyDataParticipant) *QueryProcessesWorkItemsResponseBodyData {
	s.Participant = v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetProcessInstanceId(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.ProcessInstanceId = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetProcessVersion(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.ProcessVersion = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetReceiptor(v *QueryProcessesWorkItemsResponseBodyDataReceiptor) *QueryProcessesWorkItemsResponseBodyData {
	s.Receiptor = v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetReceiveTimeGMT(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.ReceiveTimeGMT = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetSchemaCode(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.SchemaCode = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetStartTimeGMT(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.StartTimeGMT = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetState(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.State = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetWorkItemId(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.WorkItemId = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyData) SetWorkItemType(v string) *QueryProcessesWorkItemsResponseBodyData {
	s.WorkItemType = &v
	return s
}

type QueryProcessesWorkItemsResponseBodyDataFinisher struct {
	DepartmentId   *string `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId         *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryProcessesWorkItemsResponseBodyDataFinisher) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesWorkItemsResponseBodyDataFinisher) GoString() string {
	return s.String()
}

func (s *QueryProcessesWorkItemsResponseBodyDataFinisher) SetDepartmentId(v string) *QueryProcessesWorkItemsResponseBodyDataFinisher {
	s.DepartmentId = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyDataFinisher) SetDepartmentName(v string) *QueryProcessesWorkItemsResponseBodyDataFinisher {
	s.DepartmentName = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyDataFinisher) SetName(v string) *QueryProcessesWorkItemsResponseBodyDataFinisher {
	s.Name = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyDataFinisher) SetUserId(v string) *QueryProcessesWorkItemsResponseBodyDataFinisher {
	s.UserId = &v
	return s
}

type QueryProcessesWorkItemsResponseBodyDataParticipant struct {
	DepartmentId   *string `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId         *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryProcessesWorkItemsResponseBodyDataParticipant) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesWorkItemsResponseBodyDataParticipant) GoString() string {
	return s.String()
}

func (s *QueryProcessesWorkItemsResponseBodyDataParticipant) SetDepartmentId(v string) *QueryProcessesWorkItemsResponseBodyDataParticipant {
	s.DepartmentId = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyDataParticipant) SetDepartmentName(v string) *QueryProcessesWorkItemsResponseBodyDataParticipant {
	s.DepartmentName = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyDataParticipant) SetName(v string) *QueryProcessesWorkItemsResponseBodyDataParticipant {
	s.Name = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyDataParticipant) SetUserId(v string) *QueryProcessesWorkItemsResponseBodyDataParticipant {
	s.UserId = &v
	return s
}

type QueryProcessesWorkItemsResponseBodyDataReceiptor struct {
	DepartmentId   *string `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId         *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryProcessesWorkItemsResponseBodyDataReceiptor) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesWorkItemsResponseBodyDataReceiptor) GoString() string {
	return s.String()
}

func (s *QueryProcessesWorkItemsResponseBodyDataReceiptor) SetDepartmentId(v string) *QueryProcessesWorkItemsResponseBodyDataReceiptor {
	s.DepartmentId = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyDataReceiptor) SetDepartmentName(v string) *QueryProcessesWorkItemsResponseBodyDataReceiptor {
	s.DepartmentName = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyDataReceiptor) SetName(v string) *QueryProcessesWorkItemsResponseBodyDataReceiptor {
	s.Name = &v
	return s
}

func (s *QueryProcessesWorkItemsResponseBodyDataReceiptor) SetUserId(v string) *QueryProcessesWorkItemsResponseBodyDataReceiptor {
	s.UserId = &v
	return s
}

type QueryProcessesWorkItemsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryProcessesWorkItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryProcessesWorkItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessesWorkItemsResponse) GoString() string {
	return s.String()
}

func (s *QueryProcessesWorkItemsResponse) SetHeaders(v map[string]*string) *QueryProcessesWorkItemsResponse {
	s.Headers = v
	return s
}

func (s *QueryProcessesWorkItemsResponse) SetStatusCode(v int32) *QueryProcessesWorkItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryProcessesWorkItemsResponse) SetBody(v *QueryProcessesWorkItemsResponseBody) *QueryProcessesWorkItemsResponse {
	s.Body = v
	return s
}

type UpdateBizObjectHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateBizObjectHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateBizObjectHeaders) GoString() string {
	return s.String()
}

func (s *UpdateBizObjectHeaders) SetCommonHeaders(v map[string]*string) *UpdateBizObjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateBizObjectHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateBizObjectHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateBizObjectRequest struct {
	BizObjectId   *string `json:"bizObjectId,omitempty" xml:"bizObjectId,omitempty"`
	BizObjectJson *string `json:"bizObjectJson,omitempty" xml:"bizObjectJson,omitempty"`
	SchemaCode    *string `json:"schemaCode,omitempty" xml:"schemaCode,omitempty"`
}

func (s UpdateBizObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBizObjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateBizObjectRequest) SetBizObjectId(v string) *UpdateBizObjectRequest {
	s.BizObjectId = &v
	return s
}

func (s *UpdateBizObjectRequest) SetBizObjectJson(v string) *UpdateBizObjectRequest {
	s.BizObjectJson = &v
	return s
}

func (s *UpdateBizObjectRequest) SetSchemaCode(v string) *UpdateBizObjectRequest {
	s.SchemaCode = &v
	return s
}

type UpdateBizObjectResponseBody struct {
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpdateBizObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBizObjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBizObjectResponseBody) SetCode(v string) *UpdateBizObjectResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateBizObjectResponseBody) SetMessage(v string) *UpdateBizObjectResponseBody {
	s.Message = &v
	return s
}

type UpdateBizObjectResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateBizObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateBizObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBizObjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateBizObjectResponse) SetHeaders(v map[string]*string) *UpdateBizObjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateBizObjectResponse) SetStatusCode(v int32) *UpdateBizObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBizObjectResponse) SetBody(v *UpdateBizObjectResponseBody) *UpdateBizObjectResponse {
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

func (client *Client) BatchInsertBizObjectWithOptions(request *BatchInsertBizObjectRequest, headers *BatchInsertBizObjectHeaders, runtime *util.RuntimeOptions) (_result *BatchInsertBizObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizObjectJsonArray)) {
		body["bizObjectJsonArray"] = request.BizObjectJsonArray
	}

	if !tea.BoolValue(util.IsUnset(request.IsDraft)) {
		body["isDraft"] = request.IsDraft
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaCode)) {
		body["schemaCode"] = request.SchemaCode
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
		Action:      tea.String("BatchInsertBizObject"),
		Version:     tea.String("h3yun_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/h3yun/forms/instances/batch"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchInsertBizObjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchInsertBizObject(request *BatchInsertBizObjectRequest) (_result *BatchInsertBizObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchInsertBizObjectHeaders{}
	_result = &BatchInsertBizObjectResponse{}
	_body, _err := client.BatchInsertBizObjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelProcessInstanceWithOptions(request *CancelProcessInstanceRequest, headers *CancelProcessInstanceHeaders, runtime *util.RuntimeOptions) (_result *CancelProcessInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		body["processInstanceId"] = request.ProcessInstanceId
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
		Action:      tea.String("CancelProcessInstance"),
		Version:     tea.String("h3yun_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/h3yun/processes/instances/cancel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelProcessInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelProcessInstance(request *CancelProcessInstanceRequest) (_result *CancelProcessInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CancelProcessInstanceHeaders{}
	_result = &CancelProcessInstanceResponse{}
	_body, _err := client.CancelProcessInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBizObjectWithOptions(request *CreateBizObjectRequest, headers *CreateBizObjectHeaders, runtime *util.RuntimeOptions) (_result *CreateBizObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizObjectJson)) {
		body["bizObjectJson"] = request.BizObjectJson
	}

	if !tea.BoolValue(util.IsUnset(request.IsDraft)) {
		body["isDraft"] = request.IsDraft
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaCode)) {
		body["schemaCode"] = request.SchemaCode
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
		Action:      tea.String("CreateBizObject"),
		Version:     tea.String("h3yun_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/h3yun/forms/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBizObjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBizObject(request *CreateBizObjectRequest) (_result *CreateBizObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateBizObjectHeaders{}
	_result = &CreateBizObjectResponse{}
	_body, _err := client.CreateBizObjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProcessesInstanceWithOptions(request *CreateProcessesInstanceRequest, headers *CreateProcessesInstanceHeaders, runtime *util.RuntimeOptions) (_result *CreateProcessesInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizObjectId)) {
		body["bizObjectId"] = request.BizObjectId
	}

	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaCode)) {
		body["schemaCode"] = request.SchemaCode
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
		Action:      tea.String("CreateProcessesInstance"),
		Version:     tea.String("h3yun_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/h3yun/processes/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProcessesInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProcessesInstance(request *CreateProcessesInstanceRequest) (_result *CreateProcessesInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateProcessesInstanceHeaders{}
	_result = &CreateProcessesInstanceResponse{}
	_body, _err := client.CreateProcessesInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBizObjectWithOptions(request *DeleteBizObjectRequest, headers *DeleteBizObjectHeaders, runtime *util.RuntimeOptions) (_result *DeleteBizObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizObjectId)) {
		query["bizObjectId"] = request.BizObjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaCode)) {
		query["schemaCode"] = request.SchemaCode
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
		Action:      tea.String("DeleteBizObject"),
		Version:     tea.String("h3yun_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/h3yun/forms/instances"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBizObjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBizObject(request *DeleteBizObjectRequest) (_result *DeleteBizObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteBizObjectHeaders{}
	_result = &DeleteBizObjectResponse{}
	_body, _err := client.DeleteBizObjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProcessesInstanceWithOptions(request *DeleteProcessesInstanceRequest, headers *DeleteProcessesInstanceHeaders, runtime *util.RuntimeOptions) (_result *DeleteProcessesInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsAutoUpdateBizObject)) {
		query["isAutoUpdateBizObject"] = request.IsAutoUpdateBizObject
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		query["processInstanceId"] = request.ProcessInstanceId
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
		Action:      tea.String("DeleteProcessesInstance"),
		Version:     tea.String("h3yun_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/h3yun/processes/instances"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProcessesInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProcessesInstance(request *DeleteProcessesInstanceRequest) (_result *DeleteProcessesInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteProcessesInstanceHeaders{}
	_result = &DeleteProcessesInstanceResponse{}
	_body, _err := client.DeleteProcessesInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppsWithOptions(request *GetAppsRequest, headers *GetAppsHeaders, runtime *util.RuntimeOptions) (_result *GetAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		body["queryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.Values)) {
		body["values"] = request.Values
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
		Action:      tea.String("GetApps"),
		Version:     tea.String("h3yun_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/h3yun/apps/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetApps(request *GetAppsRequest) (_result *GetAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAppsHeaders{}
	_result = &GetAppsResponse{}
	_body, _err := client.GetAppsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAttachmentTemporaryUrlWithOptions(request *GetAttachmentTemporaryUrlRequest, headers *GetAttachmentTemporaryUrlHeaders, runtime *util.RuntimeOptions) (_result *GetAttachmentTemporaryUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttachmentId)) {
		query["attachmentId"] = request.AttachmentId
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
		Action:      tea.String("GetAttachmentTemporaryUrl"),
		Version:     tea.String("h3yun_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/h3yun/attachments/temporaryUrls"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAttachmentTemporaryUrlResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAttachmentTemporaryUrl(request *GetAttachmentTemporaryUrlRequest) (_result *GetAttachmentTemporaryUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAttachmentTemporaryUrlHeaders{}
	_result = &GetAttachmentTemporaryUrlResponse{}
	_body, _err := client.GetAttachmentTemporaryUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOrganizationsWithOptions(request *GetOrganizationsRequest, headers *GetOrganizationsHeaders, runtime *util.RuntimeOptions) (_result *GetOrganizationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["departmentId"] = request.DepartmentId
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
		Action:      tea.String("GetOrganizations"),
		Version:     tea.String("h3yun_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/h3yun/departments"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOrganizationsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOrganizations(request *GetOrganizationsRequest) (_result *GetOrganizationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOrganizationsHeaders{}
	_result = &GetOrganizationsResponse{}
	_body, _err := client.GetOrganizationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRoleUsersWithOptions(request *GetRoleUsersRequest, headers *GetRoleUsersHeaders, runtime *util.RuntimeOptions) (_result *GetRoleUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleId)) {
		query["roleId"] = request.RoleId
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
		Action:      tea.String("GetRoleUsers"),
		Version:     tea.String("h3yun_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/h3yun/roles/roleUsers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRoleUsersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRoleUsers(request *GetRoleUsersRequest) (_result *GetRoleUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetRoleUsersHeaders{}
	_result = &GetRoleUsersResponse{}
	_body, _err := client.GetRoleUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRolesWithOptions(headers *GetRolesHeaders, runtime *util.RuntimeOptions) (_result *GetRolesResponse, _err error) {
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
		Action:      tea.String("GetRoles"),
		Version:     tea.String("h3yun_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/h3yun/roles"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRolesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRoles() (_result *GetRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetRolesHeaders{}
	_result = &GetRolesResponse{}
	_body, _err := client.GetRolesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUploadUrlWithOptions(request *GetUploadUrlRequest, headers *GetUploadUrlHeaders, runtime *util.RuntimeOptions) (_result *GetUploadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizObjectId)) {
		query["bizObjectId"] = request.BizObjectId
	}

	if !tea.BoolValue(util.IsUnset(request.FieldName)) {
		query["fieldName"] = request.FieldName
	}

	if !tea.BoolValue(util.IsUnset(request.IsOverwrite)) {
		query["isOverwrite"] = request.IsOverwrite
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaCode)) {
		query["schemaCode"] = request.SchemaCode
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
		Action:      tea.String("GetUploadUrl"),
		Version:     tea.String("h3yun_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/h3yun/attachments/uploadUrls"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUploadUrlResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUploadUrl(request *GetUploadUrlRequest) (_result *GetUploadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUploadUrlHeaders{}
	_result = &GetUploadUrlResponse{}
	_body, _err := client.GetUploadUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUsersWithOptions(request *GetUsersRequest, headers *GetUsersHeaders, runtime *util.RuntimeOptions) (_result *GetUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["departmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.IsRecursive)) {
		query["isRecursive"] = request.IsRecursive
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
		Action:      tea.String("GetUsers"),
		Version:     tea.String("h3yun_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/h3yun/users"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUsersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUsers(request *GetUsersRequest) (_result *GetUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUsersHeaders{}
	_result = &GetUsersResponse{}
	_body, _err := client.GetUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LoadBizFieldsWithOptions(request *LoadBizFieldsRequest, headers *LoadBizFieldsHeaders, runtime *util.RuntimeOptions) (_result *LoadBizFieldsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SchemaCode)) {
		query["schemaCode"] = request.SchemaCode
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
		Action:      tea.String("LoadBizFields"),
		Version:     tea.String("h3yun_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/h3yun/forms/loadBizFields"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &LoadBizFieldsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LoadBizFields(request *LoadBizFieldsRequest) (_result *LoadBizFieldsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &LoadBizFieldsHeaders{}
	_result = &LoadBizFieldsResponse{}
	_body, _err := client.LoadBizFieldsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LoadBizObjectWithOptions(request *LoadBizObjectRequest, headers *LoadBizObjectHeaders, runtime *util.RuntimeOptions) (_result *LoadBizObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizObjectId)) {
		query["bizObjectId"] = request.BizObjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaCode)) {
		query["schemaCode"] = request.SchemaCode
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
		Action:      tea.String("LoadBizObject"),
		Version:     tea.String("h3yun_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/h3yun/forms/instances/loadInstances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &LoadBizObjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LoadBizObject(request *LoadBizObjectRequest) (_result *LoadBizObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &LoadBizObjectHeaders{}
	_result = &LoadBizObjectResponse{}
	_body, _err := client.LoadBizObjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LoadBizObjectsWithOptions(request *LoadBizObjectsRequest, headers *LoadBizObjectsHeaders, runtime *util.RuntimeOptions) (_result *LoadBizObjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MatcherJson)) {
		body["matcherJson"] = request.MatcherJson
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnFields)) {
		body["returnFields"] = request.ReturnFields
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaCode)) {
		body["schemaCode"] = request.SchemaCode
	}

	if !tea.BoolValue(util.IsUnset(request.SortByFields)) {
		body["sortByFields"] = request.SortByFields
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
		Action:      tea.String("LoadBizObjects"),
		Version:     tea.String("h3yun_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/h3yun/forms/instances/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &LoadBizObjectsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LoadBizObjects(request *LoadBizObjectsRequest) (_result *LoadBizObjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &LoadBizObjectsHeaders{}
	_result = &LoadBizObjectsResponse{}
	_body, _err := client.LoadBizObjectsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAppFunctionNodesWithOptions(request *QueryAppFunctionNodesRequest, headers *QueryAppFunctionNodesHeaders, runtime *util.RuntimeOptions) (_result *QueryAppFunctionNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppCode)) {
		query["appCode"] = request.AppCode
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
		Action:      tea.String("QueryAppFunctionNodes"),
		Version:     tea.String("h3yun_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/h3yun/apps/functionNodes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAppFunctionNodesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAppFunctionNodes(request *QueryAppFunctionNodesRequest) (_result *QueryAppFunctionNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAppFunctionNodesHeaders{}
	_result = &QueryAppFunctionNodesResponse{}
	_body, _err := client.QueryAppFunctionNodesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryProcessesInstanceWithOptions(request *QueryProcessesInstanceRequest, headers *QueryProcessesInstanceHeaders, runtime *util.RuntimeOptions) (_result *QueryProcessesInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizObjectId)) {
		query["bizObjectId"] = request.BizObjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaCode)) {
		query["schemaCode"] = request.SchemaCode
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
		Action:      tea.String("QueryProcessesInstance"),
		Version:     tea.String("h3yun_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/h3yun/processes/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryProcessesInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryProcessesInstance(request *QueryProcessesInstanceRequest) (_result *QueryProcessesInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryProcessesInstanceHeaders{}
	_result = &QueryProcessesInstanceResponse{}
	_body, _err := client.QueryProcessesInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryProcessesWorkItemsWithOptions(request *QueryProcessesWorkItemsRequest, headers *QueryProcessesWorkItemsHeaders, runtime *util.RuntimeOptions) (_result *QueryProcessesWorkItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		query["processInstanceId"] = request.ProcessInstanceId
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
		Action:      tea.String("QueryProcessesWorkItems"),
		Version:     tea.String("h3yun_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/h3yun/processes/workItems"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryProcessesWorkItemsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryProcessesWorkItems(request *QueryProcessesWorkItemsRequest) (_result *QueryProcessesWorkItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryProcessesWorkItemsHeaders{}
	_result = &QueryProcessesWorkItemsResponse{}
	_body, _err := client.QueryProcessesWorkItemsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateBizObjectWithOptions(request *UpdateBizObjectRequest, headers *UpdateBizObjectHeaders, runtime *util.RuntimeOptions) (_result *UpdateBizObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizObjectId)) {
		body["bizObjectId"] = request.BizObjectId
	}

	if !tea.BoolValue(util.IsUnset(request.BizObjectJson)) {
		body["bizObjectJson"] = request.BizObjectJson
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaCode)) {
		body["schemaCode"] = request.SchemaCode
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
		Action:      tea.String("UpdateBizObject"),
		Version:     tea.String("h3yun_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/h3yun/forms/instances"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateBizObjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateBizObject(request *UpdateBizObjectRequest) (_result *UpdateBizObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateBizObjectHeaders{}
	_result = &UpdateBizObjectResponse{}
	_body, _err := client.UpdateBizObjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
