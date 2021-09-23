// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package yida_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ValidateOrderUpgradeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ValidateOrderUpgradeHeaders) String() string {
	return tea.Prettify(s)
}

func (s ValidateOrderUpgradeHeaders) GoString() string {
	return s.String()
}

func (s *ValidateOrderUpgradeHeaders) SetCommonHeaders(v map[string]*string) *ValidateOrderUpgradeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ValidateOrderUpgradeHeaders) SetXAcsDingtalkAccessToken(v string) *ValidateOrderUpgradeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ValidateOrderUpgradeRequest struct {
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	AccessKey  *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	CallerUid  *string `json:"callerUid,omitempty" xml:"callerUid,omitempty"`
}

func (s ValidateOrderUpgradeRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateOrderUpgradeRequest) GoString() string {
	return s.String()
}

func (s *ValidateOrderUpgradeRequest) SetInstanceId(v string) *ValidateOrderUpgradeRequest {
	s.InstanceId = &v
	return s
}

func (s *ValidateOrderUpgradeRequest) SetAccessKey(v string) *ValidateOrderUpgradeRequest {
	s.AccessKey = &v
	return s
}

func (s *ValidateOrderUpgradeRequest) SetCallerUid(v string) *ValidateOrderUpgradeRequest {
	s.CallerUid = &v
	return s
}

type ValidateOrderUpgradeResponseBody struct {
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// status
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ValidateOrderUpgradeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateOrderUpgradeResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateOrderUpgradeResponseBody) SetMessage(v string) *ValidateOrderUpgradeResponseBody {
	s.Message = &v
	return s
}

func (s *ValidateOrderUpgradeResponseBody) SetStatus(v int32) *ValidateOrderUpgradeResponseBody {
	s.Status = &v
	return s
}

type ValidateOrderUpgradeResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ValidateOrderUpgradeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateOrderUpgradeResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateOrderUpgradeResponse) GoString() string {
	return s.String()
}

func (s *ValidateOrderUpgradeResponse) SetHeaders(v map[string]*string) *ValidateOrderUpgradeResponse {
	s.Headers = v
	return s
}

func (s *ValidateOrderUpgradeResponse) SetBody(v *ValidateOrderUpgradeResponseBody) *ValidateOrderUpgradeResponse {
	s.Body = v
	return s
}

type GetCorpLevelByAccountIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCorpLevelByAccountIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCorpLevelByAccountIdHeaders) GoString() string {
	return s.String()
}

func (s *GetCorpLevelByAccountIdHeaders) SetCommonHeaders(v map[string]*string) *GetCorpLevelByAccountIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCorpLevelByAccountIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetCorpLevelByAccountIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCorpLevelByAccountIdRequest struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetCorpLevelByAccountIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCorpLevelByAccountIdRequest) GoString() string {
	return s.String()
}

func (s *GetCorpLevelByAccountIdRequest) SetAccountId(v string) *GetCorpLevelByAccountIdRequest {
	s.AccountId = &v
	return s
}

type GetCorpLevelByAccountIdResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetCorpLevelByAccountIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCorpLevelByAccountIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetCorpLevelByAccountIdResponseBody) SetResult(v string) *GetCorpLevelByAccountIdResponseBody {
	s.Result = &v
	return s
}

type GetCorpLevelByAccountIdResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCorpLevelByAccountIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCorpLevelByAccountIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCorpLevelByAccountIdResponse) GoString() string {
	return s.String()
}

func (s *GetCorpLevelByAccountIdResponse) SetHeaders(v map[string]*string) *GetCorpLevelByAccountIdResponse {
	s.Headers = v
	return s
}

func (s *GetCorpLevelByAccountIdResponse) SetBody(v *GetCorpLevelByAccountIdResponseBody) *GetCorpLevelByAccountIdResponse {
	s.Body = v
	return s
}

type UpdateStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateStatusHeaders) GoString() string {
	return s.String()
}

func (s *UpdateStatusHeaders) SetCommonHeaders(v map[string]*string) *UpdateStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateStatusHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateStatusRequest struct {
	ImportSequence *string  `json:"importSequence,omitempty" xml:"importSequence,omitempty"`
	ErrorLines     []*int32 `json:"errorLines,omitempty" xml:"errorLines,omitempty" type:"Repeated"`
	AppType        *string  `json:"appType,omitempty" xml:"appType,omitempty"`
	SystemToken    *string  `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	Language       *string  `json:"language,omitempty" xml:"language,omitempty"`
	UserId         *string  `json:"userId,omitempty" xml:"userId,omitempty"`
	Status         *string  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateStatusRequest) SetImportSequence(v string) *UpdateStatusRequest {
	s.ImportSequence = &v
	return s
}

func (s *UpdateStatusRequest) SetErrorLines(v []*int32) *UpdateStatusRequest {
	s.ErrorLines = v
	return s
}

func (s *UpdateStatusRequest) SetAppType(v string) *UpdateStatusRequest {
	s.AppType = &v
	return s
}

func (s *UpdateStatusRequest) SetSystemToken(v string) *UpdateStatusRequest {
	s.SystemToken = &v
	return s
}

func (s *UpdateStatusRequest) SetLanguage(v string) *UpdateStatusRequest {
	s.Language = &v
	return s
}

func (s *UpdateStatusRequest) SetUserId(v string) *UpdateStatusRequest {
	s.UserId = &v
	return s
}

func (s *UpdateStatusRequest) SetStatus(v string) *UpdateStatusRequest {
	s.Status = &v
	return s
}

type UpdateStatusResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpdateStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateStatusResponse) SetHeaders(v map[string]*string) *UpdateStatusResponse {
	s.Headers = v
	return s
}

type ExecutePlatformTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ExecutePlatformTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s ExecutePlatformTaskHeaders) GoString() string {
	return s.String()
}

func (s *ExecutePlatformTaskHeaders) SetCommonHeaders(v map[string]*string) *ExecutePlatformTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ExecutePlatformTaskHeaders) SetXAcsDingtalkAccessToken(v string) *ExecutePlatformTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ExecutePlatformTaskRequest struct {
	// 审批结果
	OutResult *string `json:"outResult,omitempty" xml:"outResult,omitempty"`
	// 是否不执行校验&关联操作
	NoExecuteExpressions *string `json:"noExecuteExpressions,omitempty" xml:"noExecuteExpressions,omitempty"`
	// 应用ID
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 更新的表单数据
	FormDataJson *string `json:"formDataJson,omitempty" xml:"formDataJson,omitempty"`
	// 应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 语言
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 审批意见
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 流程实例ID
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// 钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecutePlatformTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecutePlatformTaskRequest) GoString() string {
	return s.String()
}

func (s *ExecutePlatformTaskRequest) SetOutResult(v string) *ExecutePlatformTaskRequest {
	s.OutResult = &v
	return s
}

func (s *ExecutePlatformTaskRequest) SetNoExecuteExpressions(v string) *ExecutePlatformTaskRequest {
	s.NoExecuteExpressions = &v
	return s
}

func (s *ExecutePlatformTaskRequest) SetAppType(v string) *ExecutePlatformTaskRequest {
	s.AppType = &v
	return s
}

func (s *ExecutePlatformTaskRequest) SetFormDataJson(v string) *ExecutePlatformTaskRequest {
	s.FormDataJson = &v
	return s
}

func (s *ExecutePlatformTaskRequest) SetSystemToken(v string) *ExecutePlatformTaskRequest {
	s.SystemToken = &v
	return s
}

func (s *ExecutePlatformTaskRequest) SetLanguage(v string) *ExecutePlatformTaskRequest {
	s.Language = &v
	return s
}

func (s *ExecutePlatformTaskRequest) SetRemark(v string) *ExecutePlatformTaskRequest {
	s.Remark = &v
	return s
}

func (s *ExecutePlatformTaskRequest) SetProcessInstanceId(v string) *ExecutePlatformTaskRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *ExecutePlatformTaskRequest) SetUserId(v string) *ExecutePlatformTaskRequest {
	s.UserId = &v
	return s
}

type ExecutePlatformTaskResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ExecutePlatformTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecutePlatformTaskResponse) GoString() string {
	return s.String()
}

func (s *ExecutePlatformTaskResponse) SetHeaders(v map[string]*string) *ExecutePlatformTaskResponse {
	s.Headers = v
	return s
}

type SaveFormRemarkHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SaveFormRemarkHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveFormRemarkHeaders) GoString() string {
	return s.String()
}

func (s *SaveFormRemarkHeaders) SetCommonHeaders(v map[string]*string) *SaveFormRemarkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveFormRemarkHeaders) SetXAcsDingtalkAccessToken(v string) *SaveFormRemarkHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SaveFormRemarkRequest struct {
	// 应用ID
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 对评论进行回复
	ReplyId *int64 `json:"replyId,omitempty" xml:"replyId,omitempty"`
	// 语言
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 实例ID
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// 评论人钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 将评论内容通过钉钉发给指定用户
	AtUserId *string `json:"atUserId,omitempty" xml:"atUserId,omitempty"`
	// 评论内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s SaveFormRemarkRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveFormRemarkRequest) GoString() string {
	return s.String()
}

func (s *SaveFormRemarkRequest) SetAppType(v string) *SaveFormRemarkRequest {
	s.AppType = &v
	return s
}

func (s *SaveFormRemarkRequest) SetSystemToken(v string) *SaveFormRemarkRequest {
	s.SystemToken = &v
	return s
}

func (s *SaveFormRemarkRequest) SetReplyId(v int64) *SaveFormRemarkRequest {
	s.ReplyId = &v
	return s
}

func (s *SaveFormRemarkRequest) SetLanguage(v string) *SaveFormRemarkRequest {
	s.Language = &v
	return s
}

func (s *SaveFormRemarkRequest) SetFormInstanceId(v string) *SaveFormRemarkRequest {
	s.FormInstanceId = &v
	return s
}

func (s *SaveFormRemarkRequest) SetUserId(v string) *SaveFormRemarkRequest {
	s.UserId = &v
	return s
}

func (s *SaveFormRemarkRequest) SetAtUserId(v string) *SaveFormRemarkRequest {
	s.AtUserId = &v
	return s
}

func (s *SaveFormRemarkRequest) SetContent(v string) *SaveFormRemarkRequest {
	s.Content = &v
	return s
}

type SaveFormRemarkResponseBody struct {
	// 评论的ID
	Result *int64 `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SaveFormRemarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveFormRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *SaveFormRemarkResponseBody) SetResult(v int64) *SaveFormRemarkResponseBody {
	s.Result = &v
	return s
}

type SaveFormRemarkResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveFormRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveFormRemarkResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveFormRemarkResponse) GoString() string {
	return s.String()
}

func (s *SaveFormRemarkResponse) SetHeaders(v map[string]*string) *SaveFormRemarkResponse {
	s.Headers = v
	return s
}

func (s *SaveFormRemarkResponse) SetBody(v *SaveFormRemarkResponseBody) *SaveFormRemarkResponse {
	s.Body = v
	return s
}

type SearchFormDatasHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchFormDatasHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasHeaders) GoString() string {
	return s.String()
}

func (s *SearchFormDatasHeaders) SetCommonHeaders(v map[string]*string) *SearchFormDatasHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchFormDatasHeaders) SetXAcsDingtalkAccessToken(v string) *SearchFormDatasHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchFormDatasRequest struct {
	// 应用编码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 应用秘钥。在应用数据中获取。
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 语言。可选值：zh_CN/en_US 默认：zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 表单ID
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 根据表单内组件值查询
	SearchFieldJson *string `json:"searchFieldJson,omitempty" xml:"searchFieldJson,omitempty"`
	// 当前页
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// 每页记录数
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 根据数据提交人工号查询
	OriginatorId *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	// createFrom和createTo两个时间构造一个时间段。查询在该时间段创建的数据列表, 字符串格式，且为yyyy-MM-DD格式
	CreateFromTimeGMT *string `json:"createFromTimeGMT,omitempty" xml:"createFromTimeGMT,omitempty"`
	// createFrom和createTo两个时间构造一个时间段。查询在该时间段创建的数据列表。字符串格式，且为yyyy-MM-DD格式。 和createFrom一起，相当于查询在 2018-01-01到2018-01-31之间(包含01和31号)创建的数据。
	CreateToTimeGMT *string `json:"createToTimeGMT,omitempty" xml:"createToTimeGMT,omitempty"`
	// modifiedFrom和modifiedTo构成一个时间段，查询在该时间段有修改的数据列表。字符串格式，且为yyyy-MM-DD格式
	ModifiedFromTimeGMT *string `json:"modifiedFromTimeGMT,omitempty" xml:"modifiedFromTimeGMT,omitempty"`
	// modifiedFrom和modifiedTo构成一个时间段，查询在该时间段有修改的数据列表。字符串格式，且为yyyy-MM-DD格式。 和modifiedFrom一起，相当于查询在 2018-01-01到2018-01-31之间(包含01和31号)被修改的数据。
	ModifiedToTimeGMT *string `json:"modifiedToTimeGMT,omitempty" xml:"modifiedToTimeGMT,omitempty"`
	// 指定排序字段
	DynamicOrder *string `json:"dynamicOrder,omitempty" xml:"dynamicOrder,omitempty"`
}

func (s SearchFormDatasRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasRequest) GoString() string {
	return s.String()
}

func (s *SearchFormDatasRequest) SetAppType(v string) *SearchFormDatasRequest {
	s.AppType = &v
	return s
}

func (s *SearchFormDatasRequest) SetSystemToken(v string) *SearchFormDatasRequest {
	s.SystemToken = &v
	return s
}

func (s *SearchFormDatasRequest) SetUserId(v string) *SearchFormDatasRequest {
	s.UserId = &v
	return s
}

func (s *SearchFormDatasRequest) SetLanguage(v string) *SearchFormDatasRequest {
	s.Language = &v
	return s
}

func (s *SearchFormDatasRequest) SetFormUuid(v string) *SearchFormDatasRequest {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDatasRequest) SetSearchFieldJson(v string) *SearchFormDatasRequest {
	s.SearchFieldJson = &v
	return s
}

func (s *SearchFormDatasRequest) SetCurrentPage(v int32) *SearchFormDatasRequest {
	s.CurrentPage = &v
	return s
}

func (s *SearchFormDatasRequest) SetPageSize(v int32) *SearchFormDatasRequest {
	s.PageSize = &v
	return s
}

func (s *SearchFormDatasRequest) SetOriginatorId(v string) *SearchFormDatasRequest {
	s.OriginatorId = &v
	return s
}

func (s *SearchFormDatasRequest) SetCreateFromTimeGMT(v string) *SearchFormDatasRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *SearchFormDatasRequest) SetCreateToTimeGMT(v string) *SearchFormDatasRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *SearchFormDatasRequest) SetModifiedFromTimeGMT(v string) *SearchFormDatasRequest {
	s.ModifiedFromTimeGMT = &v
	return s
}

func (s *SearchFormDatasRequest) SetModifiedToTimeGMT(v string) *SearchFormDatasRequest {
	s.ModifiedToTimeGMT = &v
	return s
}

func (s *SearchFormDatasRequest) SetDynamicOrder(v string) *SearchFormDatasRequest {
	s.DynamicOrder = &v
	return s
}

type SearchFormDatasResponseBody struct {
	// 当前页
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// 符合条件的实例总数
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 实例详情列表
	Data []*SearchFormDatasResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s SearchFormDatasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponseBody) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBody) SetCurrentPage(v int64) *SearchFormDatasResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *SearchFormDatasResponseBody) SetTotalCount(v int64) *SearchFormDatasResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchFormDatasResponseBody) SetData(v []*SearchFormDatasResponseBodyData) *SearchFormDatasResponseBody {
	s.Data = v
	return s
}

type SearchFormDatasResponseBodyData struct {
	// 实体主键id
	DataId *int64 `json:"dataId,omitempty" xml:"dataId,omitempty"`
	// 表单实例ID
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// 数据创建时间
	CreatedTimeGMT *string `json:"createdTimeGMT,omitempty" xml:"createdTimeGMT,omitempty"`
	// 最近修改时间
	ModifiedTimeGMT *string `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	// 表单id
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 模型id
	ModelUuid *string `json:"modelUuid,omitempty" xml:"modelUuid,omitempty"`
	// 发起人
	Originator *SearchFormDatasResponseBodyDataOriginator `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	// 修改者
	ModifyUser *SearchFormDatasResponseBodyDataModifyUser `json:"modifyUser,omitempty" xml:"modifyUser,omitempty" type:"Struct"`
	// 表单数据
	FormData map[string]interface{} `json:"formData,omitempty" xml:"formData,omitempty"`
	// 标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 流水号
	SerialNo *string `json:"serialNo,omitempty" xml:"serialNo,omitempty"`
	// 表单实例原始格式值
	InstanceValue *string `json:"instanceValue,omitempty" xml:"instanceValue,omitempty"`
	// 数据版本
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 创建人
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// 修改人
	ModifierUserId *string `json:"modifierUserId,omitempty" xml:"modifierUserId,omitempty"`
	// 批次号
	Sequence *string `json:"sequence,omitempty" xml:"sequence,omitempty"`
}

func (s SearchFormDatasResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBodyData) SetDataId(v int64) *SearchFormDatasResponseBodyData {
	s.DataId = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetFormInstanceId(v string) *SearchFormDatasResponseBodyData {
	s.FormInstanceId = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetCreatedTimeGMT(v string) *SearchFormDatasResponseBodyData {
	s.CreatedTimeGMT = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetModifiedTimeGMT(v string) *SearchFormDatasResponseBodyData {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetFormUuid(v string) *SearchFormDatasResponseBodyData {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetModelUuid(v string) *SearchFormDatasResponseBodyData {
	s.ModelUuid = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetOriginator(v *SearchFormDatasResponseBodyDataOriginator) *SearchFormDatasResponseBodyData {
	s.Originator = v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetModifyUser(v *SearchFormDatasResponseBodyDataModifyUser) *SearchFormDatasResponseBodyData {
	s.ModifyUser = v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetFormData(v map[string]interface{}) *SearchFormDatasResponseBodyData {
	s.FormData = v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetTitle(v string) *SearchFormDatasResponseBodyData {
	s.Title = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetSerialNo(v string) *SearchFormDatasResponseBodyData {
	s.SerialNo = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetInstanceValue(v string) *SearchFormDatasResponseBodyData {
	s.InstanceValue = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetVersion(v int64) *SearchFormDatasResponseBodyData {
	s.Version = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetCreatorUserId(v string) *SearchFormDatasResponseBodyData {
	s.CreatorUserId = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetModifierUserId(v string) *SearchFormDatasResponseBodyData {
	s.ModifierUserId = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetSequence(v string) *SearchFormDatasResponseBodyData {
	s.Sequence = &v
	return s
}

type SearchFormDatasResponseBodyDataOriginator struct {
	// 用户工号
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用户名
	UserName *SearchFormDatasResponseBodyDataOriginatorUserName `json:"userName,omitempty" xml:"userName,omitempty" type:"Struct"`
	// 部门名称
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 邮箱
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
}

func (s SearchFormDatasResponseBodyDataOriginator) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponseBodyDataOriginator) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBodyDataOriginator) SetUserId(v string) *SearchFormDatasResponseBodyDataOriginator {
	s.UserId = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataOriginator) SetUserName(v *SearchFormDatasResponseBodyDataOriginatorUserName) *SearchFormDatasResponseBodyDataOriginator {
	s.UserName = v
	return s
}

func (s *SearchFormDatasResponseBodyDataOriginator) SetDepartmentName(v string) *SearchFormDatasResponseBodyDataOriginator {
	s.DepartmentName = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataOriginator) SetEmail(v string) *SearchFormDatasResponseBodyDataOriginator {
	s.Email = &v
	return s
}

type SearchFormDatasResponseBodyDataOriginatorUserName struct {
	// 中文名称
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// 英文名称
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	// 国际化
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SearchFormDatasResponseBodyDataOriginatorUserName) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponseBodyDataOriginatorUserName) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBodyDataOriginatorUserName) SetNameInChinese(v string) *SearchFormDatasResponseBodyDataOriginatorUserName {
	s.NameInChinese = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataOriginatorUserName) SetNameInEnglish(v string) *SearchFormDatasResponseBodyDataOriginatorUserName {
	s.NameInEnglish = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataOriginatorUserName) SetType(v string) *SearchFormDatasResponseBodyDataOriginatorUserName {
	s.Type = &v
	return s
}

type SearchFormDatasResponseBodyDataModifyUser struct {
	// 用户工号
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用户名
	UserName *SearchFormDatasResponseBodyDataModifyUserUserName `json:"userName,omitempty" xml:"userName,omitempty" type:"Struct"`
	// 部门名称
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 邮箱
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
}

func (s SearchFormDatasResponseBodyDataModifyUser) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponseBodyDataModifyUser) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBodyDataModifyUser) SetUserId(v string) *SearchFormDatasResponseBodyDataModifyUser {
	s.UserId = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataModifyUser) SetUserName(v *SearchFormDatasResponseBodyDataModifyUserUserName) *SearchFormDatasResponseBodyDataModifyUser {
	s.UserName = v
	return s
}

func (s *SearchFormDatasResponseBodyDataModifyUser) SetDepartmentName(v string) *SearchFormDatasResponseBodyDataModifyUser {
	s.DepartmentName = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataModifyUser) SetEmail(v string) *SearchFormDatasResponseBodyDataModifyUser {
	s.Email = &v
	return s
}

type SearchFormDatasResponseBodyDataModifyUserUserName struct {
	// 中文名称
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// 英文名称
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	// 国际化
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SearchFormDatasResponseBodyDataModifyUserUserName) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponseBodyDataModifyUserUserName) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBodyDataModifyUserUserName) SetNameInChinese(v string) *SearchFormDatasResponseBodyDataModifyUserUserName {
	s.NameInChinese = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataModifyUserUserName) SetNameInEnglish(v string) *SearchFormDatasResponseBodyDataModifyUserUserName {
	s.NameInEnglish = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataModifyUserUserName) SetType(v string) *SearchFormDatasResponseBodyDataModifyUserUserName {
	s.Type = &v
	return s
}

type SearchFormDatasResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchFormDatasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchFormDatasResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponse) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponse) SetHeaders(v map[string]*string) *SearchFormDatasResponse {
	s.Headers = v
	return s
}

func (s *SearchFormDatasResponse) SetBody(v *SearchFormDatasResponseBody) *SearchFormDatasResponse {
	s.Body = v
	return s
}

type SearchActivationCodeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchActivationCodeHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchActivationCodeHeaders) GoString() string {
	return s.String()
}

func (s *SearchActivationCodeHeaders) SetCommonHeaders(v map[string]*string) *SearchActivationCodeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchActivationCodeHeaders) SetXAcsDingtalkAccessToken(v string) *SearchActivationCodeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchActivationCodeRequest struct {
	// 访问key
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// 用户的uid
	CallerUid *string `json:"callerUid,omitempty" xml:"callerUid,omitempty"`
}

func (s SearchActivationCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchActivationCodeRequest) GoString() string {
	return s.String()
}

func (s *SearchActivationCodeRequest) SetAccessKey(v string) *SearchActivationCodeRequest {
	s.AccessKey = &v
	return s
}

func (s *SearchActivationCodeRequest) SetCallerUid(v string) *SearchActivationCodeRequest {
	s.CallerUid = &v
	return s
}

type SearchActivationCodeResponseBody struct {
	// instanceId
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// activationCode
	ActivationCode *string `json:"activationCode,omitempty" xml:"activationCode,omitempty"`
	// authType
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty"`
	// expireTime
	ExpireTimeGMT *string `json:"expireTimeGMT,omitempty" xml:"expireTimeGMT,omitempty"`
	// status
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SearchActivationCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchActivationCodeResponseBody) GoString() string {
	return s.String()
}

func (s *SearchActivationCodeResponseBody) SetInstanceId(v string) *SearchActivationCodeResponseBody {
	s.InstanceId = &v
	return s
}

func (s *SearchActivationCodeResponseBody) SetActivationCode(v string) *SearchActivationCodeResponseBody {
	s.ActivationCode = &v
	return s
}

func (s *SearchActivationCodeResponseBody) SetAuthType(v string) *SearchActivationCodeResponseBody {
	s.AuthType = &v
	return s
}

func (s *SearchActivationCodeResponseBody) SetExpireTimeGMT(v string) *SearchActivationCodeResponseBody {
	s.ExpireTimeGMT = &v
	return s
}

func (s *SearchActivationCodeResponseBody) SetStatus(v int32) *SearchActivationCodeResponseBody {
	s.Status = &v
	return s
}

type SearchActivationCodeResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchActivationCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchActivationCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchActivationCodeResponse) GoString() string {
	return s.String()
}

func (s *SearchActivationCodeResponse) SetHeaders(v map[string]*string) *SearchActivationCodeResponse {
	s.Headers = v
	return s
}

func (s *SearchActivationCodeResponse) SetBody(v *SearchActivationCodeResponseBody) *SearchActivationCodeResponse {
	s.Body = v
	return s
}

type SearchEmployeeFieldValuesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchEmployeeFieldValuesHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchEmployeeFieldValuesHeaders) GoString() string {
	return s.String()
}

func (s *SearchEmployeeFieldValuesHeaders) SetCommonHeaders(v map[string]*string) *SearchEmployeeFieldValuesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchEmployeeFieldValuesHeaders) SetXAcsDingtalkAccessToken(v string) *SearchEmployeeFieldValuesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchEmployeeFieldValuesRequest struct {
	TargetFieldJson     *string `json:"targetFieldJson,omitempty" xml:"targetFieldJson,omitempty"`
	FormUuid            *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	AppType             *string `json:"appType,omitempty" xml:"appType,omitempty"`
	ModifiedToTimeGMT   *string `json:"modifiedToTimeGMT,omitempty" xml:"modifiedToTimeGMT,omitempty"`
	SystemToken         *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	ModifiedFromTimeGMT *string `json:"modifiedFromTimeGMT,omitempty" xml:"modifiedFromTimeGMT,omitempty"`
	Language            *string `json:"language,omitempty" xml:"language,omitempty"`
	SearchFieldJson     *string `json:"searchFieldJson,omitempty" xml:"searchFieldJson,omitempty"`
	OriginatorId        *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	UserId              *string `json:"userId,omitempty" xml:"userId,omitempty"`
	CreateToTimeGMT     *string `json:"createToTimeGMT,omitempty" xml:"createToTimeGMT,omitempty"`
	CreateFromTimeGMT   *string `json:"createFromTimeGMT,omitempty" xml:"createFromTimeGMT,omitempty"`
}

func (s SearchEmployeeFieldValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchEmployeeFieldValuesRequest) GoString() string {
	return s.String()
}

func (s *SearchEmployeeFieldValuesRequest) SetTargetFieldJson(v string) *SearchEmployeeFieldValuesRequest {
	s.TargetFieldJson = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetFormUuid(v string) *SearchEmployeeFieldValuesRequest {
	s.FormUuid = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetAppType(v string) *SearchEmployeeFieldValuesRequest {
	s.AppType = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetModifiedToTimeGMT(v string) *SearchEmployeeFieldValuesRequest {
	s.ModifiedToTimeGMT = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetSystemToken(v string) *SearchEmployeeFieldValuesRequest {
	s.SystemToken = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetModifiedFromTimeGMT(v string) *SearchEmployeeFieldValuesRequest {
	s.ModifiedFromTimeGMT = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetLanguage(v string) *SearchEmployeeFieldValuesRequest {
	s.Language = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetSearchFieldJson(v string) *SearchEmployeeFieldValuesRequest {
	s.SearchFieldJson = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetOriginatorId(v string) *SearchEmployeeFieldValuesRequest {
	s.OriginatorId = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetUserId(v string) *SearchEmployeeFieldValuesRequest {
	s.UserId = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetCreateToTimeGMT(v string) *SearchEmployeeFieldValuesRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetCreateFromTimeGMT(v string) *SearchEmployeeFieldValuesRequest {
	s.CreateFromTimeGMT = &v
	return s
}

type SearchEmployeeFieldValuesResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SearchEmployeeFieldValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchEmployeeFieldValuesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchEmployeeFieldValuesResponseBody) SetResult(v string) *SearchEmployeeFieldValuesResponseBody {
	s.Result = &v
	return s
}

type SearchEmployeeFieldValuesResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchEmployeeFieldValuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchEmployeeFieldValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchEmployeeFieldValuesResponse) GoString() string {
	return s.String()
}

func (s *SearchEmployeeFieldValuesResponse) SetHeaders(v map[string]*string) *SearchEmployeeFieldValuesResponse {
	s.Headers = v
	return s
}

func (s *SearchEmployeeFieldValuesResponse) SetBody(v *SearchEmployeeFieldValuesResponseBody) *SearchEmployeeFieldValuesResponse {
	s.Body = v
	return s
}

type UpdateFormDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateFormDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateFormDataHeaders) GoString() string {
	return s.String()
}

func (s *UpdateFormDataHeaders) SetCommonHeaders(v map[string]*string) *UpdateFormDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateFormDataHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateFormDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateFormDataRequest struct {
	// 应用编码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 应用秘钥。在应用数据中获取。
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 语言。可选值：zh_CN/en_US 默认：zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 要更新的表单数据ID
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// 使用最新的表单版本进行更新。默认为false
	UseLatestVersion *bool `json:"useLatestVersion,omitempty" xml:"useLatestVersion,omitempty"`
	// 要更新的表单组件值。参数有的组件更新，没有的组件保持不变。 明细的值只能统一更新，无法只更新明细下某个组件的值
	UpdateFormDataJson *string `json:"updateFormDataJson,omitempty" xml:"updateFormDataJson,omitempty"`
}

func (s UpdateFormDataRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFormDataRequest) GoString() string {
	return s.String()
}

func (s *UpdateFormDataRequest) SetAppType(v string) *UpdateFormDataRequest {
	s.AppType = &v
	return s
}

func (s *UpdateFormDataRequest) SetSystemToken(v string) *UpdateFormDataRequest {
	s.SystemToken = &v
	return s
}

func (s *UpdateFormDataRequest) SetUserId(v string) *UpdateFormDataRequest {
	s.UserId = &v
	return s
}

func (s *UpdateFormDataRequest) SetLanguage(v string) *UpdateFormDataRequest {
	s.Language = &v
	return s
}

func (s *UpdateFormDataRequest) SetFormInstanceId(v string) *UpdateFormDataRequest {
	s.FormInstanceId = &v
	return s
}

func (s *UpdateFormDataRequest) SetUseLatestVersion(v bool) *UpdateFormDataRequest {
	s.UseLatestVersion = &v
	return s
}

func (s *UpdateFormDataRequest) SetUpdateFormDataJson(v string) *UpdateFormDataRequest {
	s.UpdateFormDataJson = &v
	return s
}

type UpdateFormDataResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpdateFormDataResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFormDataResponse) GoString() string {
	return s.String()
}

func (s *UpdateFormDataResponse) SetHeaders(v map[string]*string) *UpdateFormDataResponse {
	s.Headers = v
	return s
}

type GetOperationRecordsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOperationRecordsHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOperationRecordsHeaders) GoString() string {
	return s.String()
}

func (s *GetOperationRecordsHeaders) SetCommonHeaders(v map[string]*string) *GetOperationRecordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOperationRecordsHeaders) SetXAcsDingtalkAccessToken(v string) *GetOperationRecordsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOperationRecordsRequest struct {
	// 应用ID
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 语言
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 流程实例ID
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
}

func (s GetOperationRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOperationRecordsRequest) GoString() string {
	return s.String()
}

func (s *GetOperationRecordsRequest) SetAppType(v string) *GetOperationRecordsRequest {
	s.AppType = &v
	return s
}

func (s *GetOperationRecordsRequest) SetSystemToken(v string) *GetOperationRecordsRequest {
	s.SystemToken = &v
	return s
}

func (s *GetOperationRecordsRequest) SetUserId(v string) *GetOperationRecordsRequest {
	s.UserId = &v
	return s
}

func (s *GetOperationRecordsRequest) SetLanguage(v string) *GetOperationRecordsRequest {
	s.Language = &v
	return s
}

func (s *GetOperationRecordsRequest) SetProcessInstanceId(v string) *GetOperationRecordsRequest {
	s.ProcessInstanceId = &v
	return s
}

type GetOperationRecordsResponseBody struct {
	// 流程实例操作记录数组
	Result []*GetOperationRecordsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetOperationRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOperationRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *GetOperationRecordsResponseBody) SetResult(v []*GetOperationRecordsResponseBodyResult) *GetOperationRecordsResponseBody {
	s.Result = v
	return s
}

type GetOperationRecordsResponseBodyResult struct {
	// processInstanceId
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// showName
	ShowName *string `json:"showName,omitempty" xml:"showName,omitempty"`
	// operatorNick
	OperatorNickName *string `json:"operatorNickName,omitempty" xml:"operatorNickName,omitempty"`
	// activeTime
	ActiveTimeGMT *string `json:"activeTimeGMT,omitempty" xml:"activeTimeGMT,omitempty"`
	// operateTime
	OperateTimeGMT *string `json:"operateTimeGMT,omitempty" xml:"operateTimeGMT,omitempty"`
	// operateType
	OperateType *string `json:"operateType,omitempty" xml:"operateType,omitempty"`
	// operatorStatus
	OperatorStatus *string `json:"operatorStatus,omitempty" xml:"operatorStatus,omitempty"`
	// remark
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// taskHoldTime
	TaskHoldTimeGMT *int64 `json:"taskHoldTimeGMT,omitempty" xml:"taskHoldTimeGMT,omitempty"`
	// type
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// operatorName
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	// operator
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// activityId
	ActivityId *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
	// taskType
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// taskExecuteType
	TaskExecuteType *string `json:"taskExecuteType,omitempty" xml:"taskExecuteType,omitempty"`
	// size
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// operatorDisplayName
	OperatorDisplayName *string `json:"operatorDisplayName,omitempty" xml:"operatorDisplayName,omitempty"`
	// files
	Files *string `json:"files,omitempty" xml:"files,omitempty"`
	// action
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// actionExt
	ActionExit *string `json:"actionExit,omitempty" xml:"actionExit,omitempty"`
	// id
	DataId *int64 `json:"dataId,omitempty" xml:"dataId,omitempty"`
	// taskId
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// digitalSign
	DigitalSign *string `json:"digitalSign,omitempty" xml:"digitalSign,omitempty"`
	// operatorPhotoUrl
	OperatorPhotoUrl *string `json:"operatorPhotoUrl,omitempty" xml:"operatorPhotoUrl,omitempty"`
}

func (s GetOperationRecordsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetOperationRecordsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetOperationRecordsResponseBodyResult) SetProcessInstanceId(v string) *GetOperationRecordsResponseBodyResult {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetShowName(v string) *GetOperationRecordsResponseBodyResult {
	s.ShowName = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetOperatorNickName(v string) *GetOperationRecordsResponseBodyResult {
	s.OperatorNickName = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetActiveTimeGMT(v string) *GetOperationRecordsResponseBodyResult {
	s.ActiveTimeGMT = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetOperateTimeGMT(v string) *GetOperationRecordsResponseBodyResult {
	s.OperateTimeGMT = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetOperateType(v string) *GetOperationRecordsResponseBodyResult {
	s.OperateType = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetOperatorStatus(v string) *GetOperationRecordsResponseBodyResult {
	s.OperatorStatus = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetRemark(v string) *GetOperationRecordsResponseBodyResult {
	s.Remark = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetTaskHoldTimeGMT(v int64) *GetOperationRecordsResponseBodyResult {
	s.TaskHoldTimeGMT = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetType(v string) *GetOperationRecordsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetOperatorName(v string) *GetOperationRecordsResponseBodyResult {
	s.OperatorName = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetOperatorUserId(v string) *GetOperationRecordsResponseBodyResult {
	s.OperatorUserId = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetActivityId(v string) *GetOperationRecordsResponseBodyResult {
	s.ActivityId = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetTaskType(v string) *GetOperationRecordsResponseBodyResult {
	s.TaskType = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetTaskExecuteType(v string) *GetOperationRecordsResponseBodyResult {
	s.TaskExecuteType = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetSize(v int32) *GetOperationRecordsResponseBodyResult {
	s.Size = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetOperatorDisplayName(v string) *GetOperationRecordsResponseBodyResult {
	s.OperatorDisplayName = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetFiles(v string) *GetOperationRecordsResponseBodyResult {
	s.Files = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetAction(v string) *GetOperationRecordsResponseBodyResult {
	s.Action = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetActionExit(v string) *GetOperationRecordsResponseBodyResult {
	s.ActionExit = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetDataId(v int64) *GetOperationRecordsResponseBodyResult {
	s.DataId = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetTaskId(v string) *GetOperationRecordsResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetDigitalSign(v string) *GetOperationRecordsResponseBodyResult {
	s.DigitalSign = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetOperatorPhotoUrl(v string) *GetOperationRecordsResponseBodyResult {
	s.OperatorPhotoUrl = &v
	return s
}

type GetOperationRecordsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOperationRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOperationRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOperationRecordsResponse) GoString() string {
	return s.String()
}

func (s *GetOperationRecordsResponse) SetHeaders(v map[string]*string) *GetOperationRecordsResponse {
	s.Headers = v
	return s
}

func (s *GetOperationRecordsResponse) SetBody(v *GetOperationRecordsResponseBody) *GetOperationRecordsResponse {
	s.Body = v
	return s
}

type GetPlatformResourceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPlatformResourceHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPlatformResourceHeaders) GoString() string {
	return s.String()
}

func (s *GetPlatformResourceHeaders) SetCommonHeaders(v map[string]*string) *GetPlatformResourceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPlatformResourceHeaders) SetXAcsDingtalkAccessToken(v string) *GetPlatformResourceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPlatformResourceRequest struct {
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	AccessKey  *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	CallerUid  *string `json:"callerUid,omitempty" xml:"callerUid,omitempty"`
}

func (s GetPlatformResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPlatformResourceRequest) GoString() string {
	return s.String()
}

func (s *GetPlatformResourceRequest) SetInstanceId(v string) *GetPlatformResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetPlatformResourceRequest) SetAccessKey(v string) *GetPlatformResourceRequest {
	s.AccessKey = &v
	return s
}

func (s *GetPlatformResourceRequest) SetCallerUid(v string) *GetPlatformResourceRequest {
	s.CallerUid = &v
	return s
}

type GetPlatformResourceResponseBody struct {
	// appTotalAmount
	AppTotalAmount *int32 `json:"appTotalAmount,omitempty" xml:"appTotalAmount,omitempty"`
	// instanceTotalAmount
	InstanceTotalAmount *int64 `json:"instanceTotalAmount,omitempty" xml:"instanceTotalAmount,omitempty"`
	// instanceUsageAmount
	InstanceUsageAmount *int64 `json:"instanceUsageAmount,omitempty" xml:"instanceUsageAmount,omitempty"`
	// accountUsageAmount
	AccountUsageAmount *int32 `json:"accountUsageAmount,omitempty" xml:"accountUsageAmount,omitempty"`
	// accountTotalAmount
	AccountTotalAmount *int32 `json:"accountTotalAmount,omitempty" xml:"accountTotalAmount,omitempty"`
	// pluginUsageAmount
	PluginUsageAmount *int64 `json:"pluginUsageAmount,omitempty" xml:"pluginUsageAmount,omitempty"`
	// attachmentTotalAmount
	AttachmentTotalAmount *int64 `json:"attachmentTotalAmount,omitempty" xml:"attachmentTotalAmount,omitempty"`
	// attachmentUsageAmount
	AttachmentUsageAmount *int64 `json:"attachmentUsageAmount,omitempty" xml:"attachmentUsageAmount,omitempty"`
}

func (s GetPlatformResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPlatformResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetPlatformResourceResponseBody) SetAppTotalAmount(v int32) *GetPlatformResourceResponseBody {
	s.AppTotalAmount = &v
	return s
}

func (s *GetPlatformResourceResponseBody) SetInstanceTotalAmount(v int64) *GetPlatformResourceResponseBody {
	s.InstanceTotalAmount = &v
	return s
}

func (s *GetPlatformResourceResponseBody) SetInstanceUsageAmount(v int64) *GetPlatformResourceResponseBody {
	s.InstanceUsageAmount = &v
	return s
}

func (s *GetPlatformResourceResponseBody) SetAccountUsageAmount(v int32) *GetPlatformResourceResponseBody {
	s.AccountUsageAmount = &v
	return s
}

func (s *GetPlatformResourceResponseBody) SetAccountTotalAmount(v int32) *GetPlatformResourceResponseBody {
	s.AccountTotalAmount = &v
	return s
}

func (s *GetPlatformResourceResponseBody) SetPluginUsageAmount(v int64) *GetPlatformResourceResponseBody {
	s.PluginUsageAmount = &v
	return s
}

func (s *GetPlatformResourceResponseBody) SetAttachmentTotalAmount(v int64) *GetPlatformResourceResponseBody {
	s.AttachmentTotalAmount = &v
	return s
}

func (s *GetPlatformResourceResponseBody) SetAttachmentUsageAmount(v int64) *GetPlatformResourceResponseBody {
	s.AttachmentUsageAmount = &v
	return s
}

type GetPlatformResourceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPlatformResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPlatformResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPlatformResourceResponse) GoString() string {
	return s.String()
}

func (s *GetPlatformResourceResponse) SetHeaders(v map[string]*string) *GetPlatformResourceResponse {
	s.Headers = v
	return s
}

func (s *GetPlatformResourceResponse) SetBody(v *GetPlatformResourceResponseBody) *GetPlatformResourceResponse {
	s.Body = v
	return s
}

type GetRunningTasksHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetRunningTasksHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetRunningTasksHeaders) GoString() string {
	return s.String()
}

func (s *GetRunningTasksHeaders) SetCommonHeaders(v map[string]*string) *GetRunningTasksHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRunningTasksHeaders) SetXAcsDingtalkAccessToken(v string) *GetRunningTasksHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetRunningTasksRequest struct {
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	AppType           *string `json:"appType,omitempty" xml:"appType,omitempty"`
	SystemToken       *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	Language          *string `json:"language,omitempty" xml:"language,omitempty"`
	UserId            *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetRunningTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRunningTasksRequest) GoString() string {
	return s.String()
}

func (s *GetRunningTasksRequest) SetProcessInstanceId(v string) *GetRunningTasksRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetRunningTasksRequest) SetAppType(v string) *GetRunningTasksRequest {
	s.AppType = &v
	return s
}

func (s *GetRunningTasksRequest) SetSystemToken(v string) *GetRunningTasksRequest {
	s.SystemToken = &v
	return s
}

func (s *GetRunningTasksRequest) SetLanguage(v string) *GetRunningTasksRequest {
	s.Language = &v
	return s
}

func (s *GetRunningTasksRequest) SetUserId(v string) *GetRunningTasksRequest {
	s.UserId = &v
	return s
}

type GetRunningTasksResponseBody struct {
	Result []*GetRunningTasksResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetRunningTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRunningTasksResponseBody) GoString() string {
	return s.String()
}

func (s *GetRunningTasksResponseBody) SetResult(v []*GetRunningTasksResponseBodyResult) *GetRunningTasksResponseBody {
	s.Result = v
	return s
}

type GetRunningTasksResponseBodyResult struct {
	// createTime
	CreateTimeGMT *string `json:"createTimeGMT,omitempty" xml:"createTimeGMT,omitempty"`
	// activityId
	ActivityId *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
	// processInstanceId
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// taskType
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// titleEn
	TitleInEnglish *string `json:"titleInEnglish,omitempty" xml:"titleInEnglish,omitempty"`
	// activeTime
	ActiveTimeGMT *string `json:"activeTimeGMT,omitempty" xml:"activeTimeGMT,omitempty"`
	// actualActionerId
	ActualActionerId *string `json:"actualActionerId,omitempty" xml:"actualActionerId,omitempty"`
	// originatorId
	OriginatorId *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	// finishTime
	FinishTimeGMT *string `json:"finishTimeGMT,omitempty" xml:"finishTimeGMT,omitempty"`
	// title
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// taskId
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// status
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetRunningTasksResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetRunningTasksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRunningTasksResponseBodyResult) SetCreateTimeGMT(v string) *GetRunningTasksResponseBodyResult {
	s.CreateTimeGMT = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetActivityId(v string) *GetRunningTasksResponseBodyResult {
	s.ActivityId = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetProcessInstanceId(v string) *GetRunningTasksResponseBodyResult {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetTaskType(v string) *GetRunningTasksResponseBodyResult {
	s.TaskType = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetTitleInEnglish(v string) *GetRunningTasksResponseBodyResult {
	s.TitleInEnglish = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetActiveTimeGMT(v string) *GetRunningTasksResponseBodyResult {
	s.ActiveTimeGMT = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetActualActionerId(v string) *GetRunningTasksResponseBodyResult {
	s.ActualActionerId = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetOriginatorId(v string) *GetRunningTasksResponseBodyResult {
	s.OriginatorId = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetFinishTimeGMT(v string) *GetRunningTasksResponseBodyResult {
	s.FinishTimeGMT = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetTitle(v string) *GetRunningTasksResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetTaskId(v string) *GetRunningTasksResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetStatus(v string) *GetRunningTasksResponseBodyResult {
	s.Status = &v
	return s
}

type GetRunningTasksResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRunningTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRunningTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRunningTasksResponse) GoString() string {
	return s.String()
}

func (s *GetRunningTasksResponse) SetHeaders(v map[string]*string) *GetRunningTasksResponse {
	s.Headers = v
	return s
}

func (s *GetRunningTasksResponse) SetBody(v *GetRunningTasksResponseBody) *GetRunningTasksResponse {
	s.Body = v
	return s
}

type ListNavigationByFormTypeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListNavigationByFormTypeHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListNavigationByFormTypeHeaders) GoString() string {
	return s.String()
}

func (s *ListNavigationByFormTypeHeaders) SetCommonHeaders(v map[string]*string) *ListNavigationByFormTypeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListNavigationByFormTypeHeaders) SetXAcsDingtalkAccessToken(v string) *ListNavigationByFormTypeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListNavigationByFormTypeRequest struct {
	// 应用ID
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 评论人钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 语言
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 页面类型
	FormType *string `json:"formType,omitempty" xml:"formType,omitempty"`
}

func (s ListNavigationByFormTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNavigationByFormTypeRequest) GoString() string {
	return s.String()
}

func (s *ListNavigationByFormTypeRequest) SetAppType(v string) *ListNavigationByFormTypeRequest {
	s.AppType = &v
	return s
}

func (s *ListNavigationByFormTypeRequest) SetSystemToken(v string) *ListNavigationByFormTypeRequest {
	s.SystemToken = &v
	return s
}

func (s *ListNavigationByFormTypeRequest) SetUserId(v string) *ListNavigationByFormTypeRequest {
	s.UserId = &v
	return s
}

func (s *ListNavigationByFormTypeRequest) SetLanguage(v string) *ListNavigationByFormTypeRequest {
	s.Language = &v
	return s
}

func (s *ListNavigationByFormTypeRequest) SetFormType(v string) *ListNavigationByFormTypeRequest {
	s.FormType = &v
	return s
}

type ListNavigationByFormTypeResponseBody struct {
	Result []*ListNavigationByFormTypeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListNavigationByFormTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNavigationByFormTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ListNavigationByFormTypeResponseBody) SetResult(v []*ListNavigationByFormTypeResponseBodyResult) *ListNavigationByFormTypeResponseBody {
	s.Result = v
	return s
}

type ListNavigationByFormTypeResponseBodyResult struct {
	// title
	Title *ListNavigationByFormTypeResponseBodyResultTitle `json:"title,omitempty" xml:"title,omitempty" type:"Struct"`
	// processCode
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// formUuid
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
}

func (s ListNavigationByFormTypeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListNavigationByFormTypeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListNavigationByFormTypeResponseBodyResult) SetTitle(v *ListNavigationByFormTypeResponseBodyResultTitle) *ListNavigationByFormTypeResponseBodyResult {
	s.Title = v
	return s
}

func (s *ListNavigationByFormTypeResponseBodyResult) SetProcessCode(v string) *ListNavigationByFormTypeResponseBodyResult {
	s.ProcessCode = &v
	return s
}

func (s *ListNavigationByFormTypeResponseBodyResult) SetFormUuid(v string) *ListNavigationByFormTypeResponseBodyResult {
	s.FormUuid = &v
	return s
}

type ListNavigationByFormTypeResponseBodyResultTitle struct {
	// 英文名称
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	// type
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 中文名称
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
}

func (s ListNavigationByFormTypeResponseBodyResultTitle) String() string {
	return tea.Prettify(s)
}

func (s ListNavigationByFormTypeResponseBodyResultTitle) GoString() string {
	return s.String()
}

func (s *ListNavigationByFormTypeResponseBodyResultTitle) SetNameInEnglish(v string) *ListNavigationByFormTypeResponseBodyResultTitle {
	s.NameInEnglish = &v
	return s
}

func (s *ListNavigationByFormTypeResponseBodyResultTitle) SetType(v string) *ListNavigationByFormTypeResponseBodyResultTitle {
	s.Type = &v
	return s
}

func (s *ListNavigationByFormTypeResponseBodyResultTitle) SetNameInChinese(v string) *ListNavigationByFormTypeResponseBodyResultTitle {
	s.NameInChinese = &v
	return s
}

type ListNavigationByFormTypeResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListNavigationByFormTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNavigationByFormTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNavigationByFormTypeResponse) GoString() string {
	return s.String()
}

func (s *ListNavigationByFormTypeResponse) SetHeaders(v map[string]*string) *ListNavigationByFormTypeResponse {
	s.Headers = v
	return s
}

func (s *ListNavigationByFormTypeResponse) SetBody(v *ListNavigationByFormTypeResponseBody) *ListNavigationByFormTypeResponse {
	s.Body = v
	return s
}

type TerminateInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TerminateInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s TerminateInstanceHeaders) GoString() string {
	return s.String()
}

func (s *TerminateInstanceHeaders) SetCommonHeaders(v map[string]*string) *TerminateInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TerminateInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *TerminateInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TerminateInstanceRequest struct {
	// 应用ID
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 语言
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 流程实例ID
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
}

func (s TerminateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s TerminateInstanceRequest) GoString() string {
	return s.String()
}

func (s *TerminateInstanceRequest) SetAppType(v string) *TerminateInstanceRequest {
	s.AppType = &v
	return s
}

func (s *TerminateInstanceRequest) SetSystemToken(v string) *TerminateInstanceRequest {
	s.SystemToken = &v
	return s
}

func (s *TerminateInstanceRequest) SetUserId(v string) *TerminateInstanceRequest {
	s.UserId = &v
	return s
}

func (s *TerminateInstanceRequest) SetLanguage(v string) *TerminateInstanceRequest {
	s.Language = &v
	return s
}

func (s *TerminateInstanceRequest) SetProcessInstanceId(v string) *TerminateInstanceRequest {
	s.ProcessInstanceId = &v
	return s
}

type TerminateInstanceResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s TerminateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s TerminateInstanceResponse) GoString() string {
	return s.String()
}

func (s *TerminateInstanceResponse) SetHeaders(v map[string]*string) *TerminateInstanceResponse {
	s.Headers = v
	return s
}

type ExpireCommodityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ExpireCommodityHeaders) String() string {
	return tea.Prettify(s)
}

func (s ExpireCommodityHeaders) GoString() string {
	return s.String()
}

func (s *ExpireCommodityHeaders) SetCommonHeaders(v map[string]*string) *ExpireCommodityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ExpireCommodityHeaders) SetXAcsDingtalkAccessToken(v string) *ExpireCommodityHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ExpireCommodityRequest struct {
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	AccessKey  *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	CallerUid  *string `json:"callerUid,omitempty" xml:"callerUid,omitempty"`
}

func (s ExpireCommodityRequest) String() string {
	return tea.Prettify(s)
}

func (s ExpireCommodityRequest) GoString() string {
	return s.String()
}

func (s *ExpireCommodityRequest) SetInstanceId(v string) *ExpireCommodityRequest {
	s.InstanceId = &v
	return s
}

func (s *ExpireCommodityRequest) SetAccessKey(v string) *ExpireCommodityRequest {
	s.AccessKey = &v
	return s
}

func (s *ExpireCommodityRequest) SetCallerUid(v string) *ExpireCommodityRequest {
	s.CallerUid = &v
	return s
}

type ExpireCommodityResponseBody struct {
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExpireCommodityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExpireCommodityResponseBody) GoString() string {
	return s.String()
}

func (s *ExpireCommodityResponseBody) SetMessage(v string) *ExpireCommodityResponseBody {
	s.Message = &v
	return s
}

func (s *ExpireCommodityResponseBody) SetSuccess(v bool) *ExpireCommodityResponseBody {
	s.Success = &v
	return s
}

type ExpireCommodityResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExpireCommodityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExpireCommodityResponse) String() string {
	return tea.Prettify(s)
}

func (s ExpireCommodityResponse) GoString() string {
	return s.String()
}

func (s *ExpireCommodityResponse) SetHeaders(v map[string]*string) *ExpireCommodityResponse {
	s.Headers = v
	return s
}

func (s *ExpireCommodityResponse) SetBody(v *ExpireCommodityResponseBody) *ExpireCommodityResponse {
	s.Body = v
	return s
}

type ValidateOrderBuyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ValidateOrderBuyHeaders) String() string {
	return tea.Prettify(s)
}

func (s ValidateOrderBuyHeaders) GoString() string {
	return s.String()
}

func (s *ValidateOrderBuyHeaders) SetCommonHeaders(v map[string]*string) *ValidateOrderBuyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ValidateOrderBuyHeaders) SetXAcsDingtalkAccessToken(v string) *ValidateOrderBuyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ValidateOrderBuyRequest struct {
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	CallerUid *string `json:"callerUid,omitempty" xml:"callerUid,omitempty"`
}

func (s ValidateOrderBuyRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateOrderBuyRequest) GoString() string {
	return s.String()
}

func (s *ValidateOrderBuyRequest) SetAccessKey(v string) *ValidateOrderBuyRequest {
	s.AccessKey = &v
	return s
}

func (s *ValidateOrderBuyRequest) SetCallerUid(v string) *ValidateOrderBuyRequest {
	s.CallerUid = &v
	return s
}

type ValidateOrderBuyResponseBody struct {
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// status
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ValidateOrderBuyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateOrderBuyResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateOrderBuyResponseBody) SetMessage(v string) *ValidateOrderBuyResponseBody {
	s.Message = &v
	return s
}

func (s *ValidateOrderBuyResponseBody) SetStatus(v int32) *ValidateOrderBuyResponseBody {
	s.Status = &v
	return s
}

type ValidateOrderBuyResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ValidateOrderBuyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateOrderBuyResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateOrderBuyResponse) GoString() string {
	return s.String()
}

func (s *ValidateOrderBuyResponse) SetHeaders(v map[string]*string) *ValidateOrderBuyResponse {
	s.Headers = v
	return s
}

func (s *ValidateOrderBuyResponse) SetBody(v *ValidateOrderBuyResponseBody) *ValidateOrderBuyResponse {
	s.Body = v
	return s
}

type SaveFormDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SaveFormDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveFormDataHeaders) GoString() string {
	return s.String()
}

func (s *SaveFormDataHeaders) SetCommonHeaders(v map[string]*string) *SaveFormDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveFormDataHeaders) SetXAcsDingtalkAccessToken(v string) *SaveFormDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SaveFormDataRequest struct {
	// 应用编码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 应用秘钥。在应用数据中获取。
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 语言。可选值：zh_CN/en_US 默认：zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 表单ID
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 表单数据
	FormDataJson *string `json:"formDataJson,omitempty" xml:"formDataJson,omitempty"`
}

func (s SaveFormDataRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveFormDataRequest) GoString() string {
	return s.String()
}

func (s *SaveFormDataRequest) SetAppType(v string) *SaveFormDataRequest {
	s.AppType = &v
	return s
}

func (s *SaveFormDataRequest) SetSystemToken(v string) *SaveFormDataRequest {
	s.SystemToken = &v
	return s
}

func (s *SaveFormDataRequest) SetUserId(v string) *SaveFormDataRequest {
	s.UserId = &v
	return s
}

func (s *SaveFormDataRequest) SetLanguage(v string) *SaveFormDataRequest {
	s.Language = &v
	return s
}

func (s *SaveFormDataRequest) SetFormUuid(v string) *SaveFormDataRequest {
	s.FormUuid = &v
	return s
}

func (s *SaveFormDataRequest) SetFormDataJson(v string) *SaveFormDataRequest {
	s.FormDataJson = &v
	return s
}

type SaveFormDataResponseBody struct {
	// 表单实例ID
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SaveFormDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveFormDataResponseBody) GoString() string {
	return s.String()
}

func (s *SaveFormDataResponseBody) SetResult(v string) *SaveFormDataResponseBody {
	s.Result = &v
	return s
}

type SaveFormDataResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveFormDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveFormDataResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveFormDataResponse) GoString() string {
	return s.String()
}

func (s *SaveFormDataResponse) SetHeaders(v map[string]*string) *SaveFormDataResponse {
	s.Headers = v
	return s
}

func (s *SaveFormDataResponse) SetBody(v *SaveFormDataResponseBody) *SaveFormDataResponse {
	s.Body = v
	return s
}

type DeleteFormDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteFormDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteFormDataHeaders) GoString() string {
	return s.String()
}

func (s *DeleteFormDataHeaders) SetCommonHeaders(v map[string]*string) *DeleteFormDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteFormDataHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteFormDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteFormDataRequest struct {
	// 应用ID
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 语言
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 要删除的表单数据ID
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
}

func (s DeleteFormDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFormDataRequest) GoString() string {
	return s.String()
}

func (s *DeleteFormDataRequest) SetAppType(v string) *DeleteFormDataRequest {
	s.AppType = &v
	return s
}

func (s *DeleteFormDataRequest) SetSystemToken(v string) *DeleteFormDataRequest {
	s.SystemToken = &v
	return s
}

func (s *DeleteFormDataRequest) SetUserId(v string) *DeleteFormDataRequest {
	s.UserId = &v
	return s
}

func (s *DeleteFormDataRequest) SetLanguage(v string) *DeleteFormDataRequest {
	s.Language = &v
	return s
}

func (s *DeleteFormDataRequest) SetFormInstanceId(v string) *DeleteFormDataRequest {
	s.FormInstanceId = &v
	return s
}

type DeleteFormDataResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteFormDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFormDataResponse) GoString() string {
	return s.String()
}

func (s *DeleteFormDataResponse) SetHeaders(v map[string]*string) *DeleteFormDataResponse {
	s.Headers = v
	return s
}

type UpdateInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceHeaders) GoString() string {
	return s.String()
}

func (s *UpdateInstanceHeaders) SetCommonHeaders(v map[string]*string) *UpdateInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateInstanceRequest struct {
	// 实例ID
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// 应用ID
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 更新的表单数据
	UpdateFormDataJson *string `json:"updateFormDataJson,omitempty" xml:"updateFormDataJson,omitempty"`
	// 应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 语言环境
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) SetProcessInstanceId(v string) *UpdateInstanceRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *UpdateInstanceRequest) SetAppType(v string) *UpdateInstanceRequest {
	s.AppType = &v
	return s
}

func (s *UpdateInstanceRequest) SetUpdateFormDataJson(v string) *UpdateInstanceRequest {
	s.UpdateFormDataJson = &v
	return s
}

func (s *UpdateInstanceRequest) SetSystemToken(v string) *UpdateInstanceRequest {
	s.SystemToken = &v
	return s
}

func (s *UpdateInstanceRequest) SetLanguage(v string) *UpdateInstanceRequest {
	s.Language = &v
	return s
}

func (s *UpdateInstanceRequest) SetUserId(v string) *UpdateInstanceRequest {
	s.UserId = &v
	return s
}

type UpdateInstanceResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UpdateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponse) SetHeaders(v map[string]*string) *UpdateInstanceResponse {
	s.Headers = v
	return s
}

type ListCommodityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListCommodityHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListCommodityHeaders) GoString() string {
	return s.String()
}

func (s *ListCommodityHeaders) SetCommonHeaders(v map[string]*string) *ListCommodityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListCommodityHeaders) SetXAcsDingtalkAccessToken(v string) *ListCommodityHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListCommodityRequest struct {
	// accessKey
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// pageSize
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// callerUid
	CallerUid *string `json:"callerUid,omitempty" xml:"callerUid,omitempty"`
	// currentPage
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
}

func (s ListCommodityRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCommodityRequest) GoString() string {
	return s.String()
}

func (s *ListCommodityRequest) SetAccessKey(v string) *ListCommodityRequest {
	s.AccessKey = &v
	return s
}

func (s *ListCommodityRequest) SetPageSize(v int32) *ListCommodityRequest {
	s.PageSize = &v
	return s
}

func (s *ListCommodityRequest) SetCallerUid(v string) *ListCommodityRequest {
	s.CallerUid = &v
	return s
}

func (s *ListCommodityRequest) SetCurrentPage(v int32) *ListCommodityRequest {
	s.CurrentPage = &v
	return s
}

type ListCommodityResponseBody struct {
	// 分页大小
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// commodityVOList
	CommodityVOList []*ListCommodityResponseBodyCommodityVOList `json:"commodityVOList,omitempty" xml:"commodityVOList,omitempty" type:"Repeated"`
	// 当前第几页
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// 总数量
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListCommodityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCommodityResponseBody) GoString() string {
	return s.String()
}

func (s *ListCommodityResponseBody) SetPageSize(v int32) *ListCommodityResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCommodityResponseBody) SetCommodityVOList(v []*ListCommodityResponseBodyCommodityVOList) *ListCommodityResponseBody {
	s.CommodityVOList = v
	return s
}

func (s *ListCommodityResponseBody) SetCurrentPage(v int32) *ListCommodityResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListCommodityResponseBody) SetTotalCount(v int32) *ListCommodityResponseBody {
	s.TotalCount = &v
	return s
}

type ListCommodityResponseBodyCommodityVOList struct {
	// instanceId
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// buyDate
	BuyDateGMT *string `json:"buyDateGMT,omitempty" xml:"buyDateGMT,omitempty"`
	// expireDate
	ExpireDateGMT *string `json:"expireDateGMT,omitempty" xml:"expireDateGMT,omitempty"`
	// activationCode
	ActivationCode *string `json:"activationCode,omitempty" xml:"activationCode,omitempty"`
	// accountNum
	AccountNumber *int32 `json:"accountNumber,omitempty" xml:"accountNumber,omitempty"`
	// accountDistributionNumber
	AccountDistributionNumber *int32 `json:"accountDistributionNumber,omitempty" xml:"accountDistributionNumber,omitempty"`
	// version
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
	// status
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListCommodityResponseBodyCommodityVOList) String() string {
	return tea.Prettify(s)
}

func (s ListCommodityResponseBodyCommodityVOList) GoString() string {
	return s.String()
}

func (s *ListCommodityResponseBodyCommodityVOList) SetInstanceId(v string) *ListCommodityResponseBodyCommodityVOList {
	s.InstanceId = &v
	return s
}

func (s *ListCommodityResponseBodyCommodityVOList) SetBuyDateGMT(v string) *ListCommodityResponseBodyCommodityVOList {
	s.BuyDateGMT = &v
	return s
}

func (s *ListCommodityResponseBodyCommodityVOList) SetExpireDateGMT(v string) *ListCommodityResponseBodyCommodityVOList {
	s.ExpireDateGMT = &v
	return s
}

func (s *ListCommodityResponseBodyCommodityVOList) SetActivationCode(v string) *ListCommodityResponseBodyCommodityVOList {
	s.ActivationCode = &v
	return s
}

func (s *ListCommodityResponseBodyCommodityVOList) SetAccountNumber(v int32) *ListCommodityResponseBodyCommodityVOList {
	s.AccountNumber = &v
	return s
}

func (s *ListCommodityResponseBodyCommodityVOList) SetAccountDistributionNumber(v int32) *ListCommodityResponseBodyCommodityVOList {
	s.AccountDistributionNumber = &v
	return s
}

func (s *ListCommodityResponseBodyCommodityVOList) SetVersion(v int32) *ListCommodityResponseBodyCommodityVOList {
	s.Version = &v
	return s
}

func (s *ListCommodityResponseBodyCommodityVOList) SetStatus(v string) *ListCommodityResponseBodyCommodityVOList {
	s.Status = &v
	return s
}

type ListCommodityResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCommodityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCommodityResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCommodityResponse) GoString() string {
	return s.String()
}

func (s *ListCommodityResponse) SetHeaders(v map[string]*string) *ListCommodityResponse {
	s.Headers = v
	return s
}

func (s *ListCommodityResponse) SetBody(v *ListCommodityResponseBody) *ListCommodityResponse {
	s.Body = v
	return s
}

type GetApplicationAuthorizationServicePlatformResourceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetApplicationAuthorizationServicePlatformResourceHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationAuthorizationServicePlatformResourceHeaders) GoString() string {
	return s.String()
}

func (s *GetApplicationAuthorizationServicePlatformResourceHeaders) SetCommonHeaders(v map[string]*string) *GetApplicationAuthorizationServicePlatformResourceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetApplicationAuthorizationServicePlatformResourceHeaders) SetXAcsDingtalkAccessToken(v string) *GetApplicationAuthorizationServicePlatformResourceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetApplicationAuthorizationServicePlatformResourceRequest struct {
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	AccessKey  *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	CallerUid  *string `json:"callerUid,omitempty" xml:"callerUid,omitempty"`
}

func (s GetApplicationAuthorizationServicePlatformResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationAuthorizationServicePlatformResourceRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationAuthorizationServicePlatformResourceRequest) SetInstanceId(v string) *GetApplicationAuthorizationServicePlatformResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetApplicationAuthorizationServicePlatformResourceRequest) SetAccessKey(v string) *GetApplicationAuthorizationServicePlatformResourceRequest {
	s.AccessKey = &v
	return s
}

func (s *GetApplicationAuthorizationServicePlatformResourceRequest) SetCallerUid(v string) *GetApplicationAuthorizationServicePlatformResourceRequest {
	s.CallerUid = &v
	return s
}

type GetApplicationAuthorizationServicePlatformResourceResponseBody struct {
	// appTotalAmount
	AppTotalAmount *int32 `json:"appTotalAmount,omitempty" xml:"appTotalAmount,omitempty"`
	// instanceId
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// instanceTotalAmount
	InstanceTotalAmount *int64 `json:"instanceTotalAmount,omitempty" xml:"instanceTotalAmount,omitempty"`
	// instanceUsageAmount
	InstanceUsageAmount *int64 `json:"instanceUsageAmount,omitempty" xml:"instanceUsageAmount,omitempty"`
	// accountUsageAmount
	AccountUsageAmount *int32 `json:"accountUsageAmount,omitempty" xml:"accountUsageAmount,omitempty"`
	// accountTotalAmount
	AccountTotalAmount *int32 `json:"accountTotalAmount,omitempty" xml:"accountTotalAmount,omitempty"`
	// pluginUsageAmount
	PluginUsageAmount *int64 `json:"pluginUsageAmount,omitempty" xml:"pluginUsageAmount,omitempty"`
	// attachmentTotalAmount
	AttachmentTotalAmount *int64 `json:"attachmentTotalAmount,omitempty" xml:"attachmentTotalAmount,omitempty"`
	// attachmentUsageAmount
	AttachmentUsageAmount *int64 `json:"attachmentUsageAmount,omitempty" xml:"attachmentUsageAmount,omitempty"`
}

func (s GetApplicationAuthorizationServicePlatformResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationAuthorizationServicePlatformResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationAuthorizationServicePlatformResourceResponseBody) SetAppTotalAmount(v int32) *GetApplicationAuthorizationServicePlatformResourceResponseBody {
	s.AppTotalAmount = &v
	return s
}

func (s *GetApplicationAuthorizationServicePlatformResourceResponseBody) SetInstanceId(v string) *GetApplicationAuthorizationServicePlatformResourceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetApplicationAuthorizationServicePlatformResourceResponseBody) SetInstanceTotalAmount(v int64) *GetApplicationAuthorizationServicePlatformResourceResponseBody {
	s.InstanceTotalAmount = &v
	return s
}

func (s *GetApplicationAuthorizationServicePlatformResourceResponseBody) SetInstanceUsageAmount(v int64) *GetApplicationAuthorizationServicePlatformResourceResponseBody {
	s.InstanceUsageAmount = &v
	return s
}

func (s *GetApplicationAuthorizationServicePlatformResourceResponseBody) SetAccountUsageAmount(v int32) *GetApplicationAuthorizationServicePlatformResourceResponseBody {
	s.AccountUsageAmount = &v
	return s
}

func (s *GetApplicationAuthorizationServicePlatformResourceResponseBody) SetAccountTotalAmount(v int32) *GetApplicationAuthorizationServicePlatformResourceResponseBody {
	s.AccountTotalAmount = &v
	return s
}

func (s *GetApplicationAuthorizationServicePlatformResourceResponseBody) SetPluginUsageAmount(v int64) *GetApplicationAuthorizationServicePlatformResourceResponseBody {
	s.PluginUsageAmount = &v
	return s
}

func (s *GetApplicationAuthorizationServicePlatformResourceResponseBody) SetAttachmentTotalAmount(v int64) *GetApplicationAuthorizationServicePlatformResourceResponseBody {
	s.AttachmentTotalAmount = &v
	return s
}

func (s *GetApplicationAuthorizationServicePlatformResourceResponseBody) SetAttachmentUsageAmount(v int64) *GetApplicationAuthorizationServicePlatformResourceResponseBody {
	s.AttachmentUsageAmount = &v
	return s
}

type GetApplicationAuthorizationServicePlatformResourceResponse struct {
	Headers map[string]*string                                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetApplicationAuthorizationServicePlatformResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetApplicationAuthorizationServicePlatformResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationAuthorizationServicePlatformResourceResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationAuthorizationServicePlatformResourceResponse) SetHeaders(v map[string]*string) *GetApplicationAuthorizationServicePlatformResourceResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationAuthorizationServicePlatformResourceResponse) SetBody(v *GetApplicationAuthorizationServicePlatformResourceResponseBody) *GetApplicationAuthorizationServicePlatformResourceResponse {
	s.Body = v
	return s
}

type GetActivityListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetActivityListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetActivityListHeaders) GoString() string {
	return s.String()
}

func (s *GetActivityListHeaders) SetCommonHeaders(v map[string]*string) *GetActivityListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetActivityListHeaders) SetXAcsDingtalkAccessToken(v string) *GetActivityListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetActivityListRequest struct {
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	AppType     *string `json:"appType,omitempty" xml:"appType,omitempty"`
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	Language    *string `json:"language,omitempty" xml:"language,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetActivityListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetActivityListRequest) GoString() string {
	return s.String()
}

func (s *GetActivityListRequest) SetProcessCode(v string) *GetActivityListRequest {
	s.ProcessCode = &v
	return s
}

func (s *GetActivityListRequest) SetAppType(v string) *GetActivityListRequest {
	s.AppType = &v
	return s
}

func (s *GetActivityListRequest) SetSystemToken(v string) *GetActivityListRequest {
	s.SystemToken = &v
	return s
}

func (s *GetActivityListRequest) SetLanguage(v string) *GetActivityListRequest {
	s.Language = &v
	return s
}

func (s *GetActivityListRequest) SetUserId(v string) *GetActivityListRequest {
	s.UserId = &v
	return s
}

type GetActivityListResponseBody struct {
	// Id of the request
	Result []*GetActivityListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetActivityListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetActivityListResponseBody) GoString() string {
	return s.String()
}

func (s *GetActivityListResponseBody) SetResult(v []*GetActivityListResponseBodyResult) *GetActivityListResponseBody {
	s.Result = v
	return s
}

type GetActivityListResponseBodyResult struct {
	// activityName
	ActivityName *string `json:"activityName,omitempty" xml:"activityName,omitempty"`
	// activityNameEn
	ActivityNameInEnglish *string `json:"activityNameInEnglish,omitempty" xml:"activityNameInEnglish,omitempty"`
	// activityId
	ActivityId *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
}

func (s GetActivityListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetActivityListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetActivityListResponseBodyResult) SetActivityName(v string) *GetActivityListResponseBodyResult {
	s.ActivityName = &v
	return s
}

func (s *GetActivityListResponseBodyResult) SetActivityNameInEnglish(v string) *GetActivityListResponseBodyResult {
	s.ActivityNameInEnglish = &v
	return s
}

func (s *GetActivityListResponseBodyResult) SetActivityId(v string) *GetActivityListResponseBodyResult {
	s.ActivityId = &v
	return s
}

type GetActivityListResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetActivityListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetActivityListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetActivityListResponse) GoString() string {
	return s.String()
}

func (s *GetActivityListResponse) SetHeaders(v map[string]*string) *GetActivityListResponse {
	s.Headers = v
	return s
}

func (s *GetActivityListResponse) SetBody(v *GetActivityListResponseBody) *GetActivityListResponse {
	s.Body = v
	return s
}

type GetFormDataByIDHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFormDataByIDHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDHeaders) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDHeaders) SetCommonHeaders(v map[string]*string) *GetFormDataByIDHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFormDataByIDHeaders) SetXAcsDingtalkAccessToken(v string) *GetFormDataByIDHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFormDataByIDRequest struct {
	// 应用编码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 应用秘钥。在应用数据中获取。
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 语言。可选值：zh_CN/en_US 默认：zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s GetFormDataByIDRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDRequest) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDRequest) SetAppType(v string) *GetFormDataByIDRequest {
	s.AppType = &v
	return s
}

func (s *GetFormDataByIDRequest) SetSystemToken(v string) *GetFormDataByIDRequest {
	s.SystemToken = &v
	return s
}

func (s *GetFormDataByIDRequest) SetUserId(v string) *GetFormDataByIDRequest {
	s.UserId = &v
	return s
}

func (s *GetFormDataByIDRequest) SetLanguage(v string) *GetFormDataByIDRequest {
	s.Language = &v
	return s
}

type GetFormDataByIDResponseBody struct {
	// 发起人详情
	Originator *GetFormDataByIDResponseBodyOriginator `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	// 最后修改时间
	ModifiedTimeGMT *string `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	// 表单ID
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 表单实例ID
	FormInstId *string `json:"formInstId,omitempty" xml:"formInstId,omitempty"`
	// 表单数据详情
	FormData map[string]interface{} `json:"formData,omitempty" xml:"formData,omitempty"`
}

func (s GetFormDataByIDResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDResponseBody) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDResponseBody) SetOriginator(v *GetFormDataByIDResponseBodyOriginator) *GetFormDataByIDResponseBody {
	s.Originator = v
	return s
}

func (s *GetFormDataByIDResponseBody) SetModifiedTimeGMT(v string) *GetFormDataByIDResponseBody {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *GetFormDataByIDResponseBody) SetFormUuid(v string) *GetFormDataByIDResponseBody {
	s.FormUuid = &v
	return s
}

func (s *GetFormDataByIDResponseBody) SetFormInstId(v string) *GetFormDataByIDResponseBody {
	s.FormInstId = &v
	return s
}

func (s *GetFormDataByIDResponseBody) SetFormData(v map[string]interface{}) *GetFormDataByIDResponseBody {
	s.FormData = v
	return s
}

type GetFormDataByIDResponseBodyOriginator struct {
	// 用户工号
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用户名
	Name *GetFormDataByIDResponseBodyOriginatorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	// 部门名称
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 邮箱
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
}

func (s GetFormDataByIDResponseBodyOriginator) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDResponseBodyOriginator) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDResponseBodyOriginator) SetUserId(v string) *GetFormDataByIDResponseBodyOriginator {
	s.UserId = &v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginator) SetName(v *GetFormDataByIDResponseBodyOriginatorName) *GetFormDataByIDResponseBodyOriginator {
	s.Name = v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginator) SetDepartmentName(v string) *GetFormDataByIDResponseBodyOriginator {
	s.DepartmentName = &v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginator) SetEmail(v string) *GetFormDataByIDResponseBodyOriginator {
	s.Email = &v
	return s
}

type GetFormDataByIDResponseBodyOriginatorName struct {
	// 中文名称
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// 英文名称
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	// 国际化
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetFormDataByIDResponseBodyOriginatorName) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDResponseBodyOriginatorName) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDResponseBodyOriginatorName) SetNameInChinese(v string) *GetFormDataByIDResponseBodyOriginatorName {
	s.NameInChinese = &v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginatorName) SetNameInEnglish(v string) *GetFormDataByIDResponseBodyOriginatorName {
	s.NameInEnglish = &v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginatorName) SetType(v string) *GetFormDataByIDResponseBodyOriginatorName {
	s.Type = &v
	return s
}

type GetFormDataByIDResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFormDataByIDResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFormDataByIDResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDResponse) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDResponse) SetHeaders(v map[string]*string) *GetFormDataByIDResponse {
	s.Headers = v
	return s
}

func (s *GetFormDataByIDResponse) SetBody(v *GetFormDataByIDResponseBody) *GetFormDataByIDResponse {
	s.Body = v
	return s
}

type ExecuteCustomApiHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ExecuteCustomApiHeaders) String() string {
	return tea.Prettify(s)
}

func (s ExecuteCustomApiHeaders) GoString() string {
	return s.String()
}

func (s *ExecuteCustomApiHeaders) SetCommonHeaders(v map[string]*string) *ExecuteCustomApiHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ExecuteCustomApiHeaders) SetXAcsDingtalkAccessToken(v string) *ExecuteCustomApiHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ExecuteCustomApiRequest struct {
	Data        *string `json:"data,omitempty" xml:"data,omitempty"`
	AppType     *string `json:"appType,omitempty" xml:"appType,omitempty"`
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	Language    *string `json:"language,omitempty" xml:"language,omitempty"`
	ServiceId   *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteCustomApiRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteCustomApiRequest) GoString() string {
	return s.String()
}

func (s *ExecuteCustomApiRequest) SetData(v string) *ExecuteCustomApiRequest {
	s.Data = &v
	return s
}

func (s *ExecuteCustomApiRequest) SetAppType(v string) *ExecuteCustomApiRequest {
	s.AppType = &v
	return s
}

func (s *ExecuteCustomApiRequest) SetSystemToken(v string) *ExecuteCustomApiRequest {
	s.SystemToken = &v
	return s
}

func (s *ExecuteCustomApiRequest) SetLanguage(v string) *ExecuteCustomApiRequest {
	s.Language = &v
	return s
}

func (s *ExecuteCustomApiRequest) SetServiceId(v string) *ExecuteCustomApiRequest {
	s.ServiceId = &v
	return s
}

func (s *ExecuteCustomApiRequest) SetUserId(v string) *ExecuteCustomApiRequest {
	s.UserId = &v
	return s
}

type ExecuteCustomApiResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ExecuteCustomApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteCustomApiResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteCustomApiResponseBody) SetResult(v string) *ExecuteCustomApiResponseBody {
	s.Result = &v
	return s
}

type ExecuteCustomApiResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExecuteCustomApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteCustomApiResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteCustomApiResponse) GoString() string {
	return s.String()
}

func (s *ExecuteCustomApiResponse) SetHeaders(v map[string]*string) *ExecuteCustomApiResponse {
	s.Headers = v
	return s
}

func (s *ExecuteCustomApiResponse) SetBody(v *ExecuteCustomApiResponseBody) *ExecuteCustomApiResponse {
	s.Body = v
	return s
}

type RefundCommodityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RefundCommodityHeaders) String() string {
	return tea.Prettify(s)
}

func (s RefundCommodityHeaders) GoString() string {
	return s.String()
}

func (s *RefundCommodityHeaders) SetCommonHeaders(v map[string]*string) *RefundCommodityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RefundCommodityHeaders) SetXAcsDingtalkAccessToken(v string) *RefundCommodityHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RefundCommodityRequest struct {
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	AccessKey  *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	CallerUid  *string `json:"callerUid,omitempty" xml:"callerUid,omitempty"`
}

func (s RefundCommodityRequest) String() string {
	return tea.Prettify(s)
}

func (s RefundCommodityRequest) GoString() string {
	return s.String()
}

func (s *RefundCommodityRequest) SetInstanceId(v string) *RefundCommodityRequest {
	s.InstanceId = &v
	return s
}

func (s *RefundCommodityRequest) SetAccessKey(v string) *RefundCommodityRequest {
	s.AccessKey = &v
	return s
}

func (s *RefundCommodityRequest) SetCallerUid(v string) *RefundCommodityRequest {
	s.CallerUid = &v
	return s
}

type RefundCommodityResponseBody struct {
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RefundCommodityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefundCommodityResponseBody) GoString() string {
	return s.String()
}

func (s *RefundCommodityResponseBody) SetMessage(v string) *RefundCommodityResponseBody {
	s.Message = &v
	return s
}

func (s *RefundCommodityResponseBody) SetSuccess(v bool) *RefundCommodityResponseBody {
	s.Success = &v
	return s
}

type RefundCommodityResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RefundCommodityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefundCommodityResponse) String() string {
	return tea.Prettify(s)
}

func (s RefundCommodityResponse) GoString() string {
	return s.String()
}

func (s *RefundCommodityResponse) SetHeaders(v map[string]*string) *RefundCommodityResponse {
	s.Headers = v
	return s
}

func (s *RefundCommodityResponse) SetBody(v *RefundCommodityResponseBody) *RefundCommodityResponse {
	s.Body = v
	return s
}

type DeleteSequenceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteSequenceHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteSequenceHeaders) GoString() string {
	return s.String()
}

func (s *DeleteSequenceHeaders) SetCommonHeaders(v map[string]*string) *DeleteSequenceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteSequenceHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteSequenceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteSequenceRequest struct {
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
	Sequence    *string `json:"sequence,omitempty" xml:"sequence,omitempty"`
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	Language    *string `json:"language,omitempty" xml:"language,omitempty"`
	AppType     *string `json:"appType,omitempty" xml:"appType,omitempty"`
}

func (s DeleteSequenceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSequenceRequest) GoString() string {
	return s.String()
}

func (s *DeleteSequenceRequest) SetUserId(v string) *DeleteSequenceRequest {
	s.UserId = &v
	return s
}

func (s *DeleteSequenceRequest) SetSequence(v string) *DeleteSequenceRequest {
	s.Sequence = &v
	return s
}

func (s *DeleteSequenceRequest) SetSystemToken(v string) *DeleteSequenceRequest {
	s.SystemToken = &v
	return s
}

func (s *DeleteSequenceRequest) SetLanguage(v string) *DeleteSequenceRequest {
	s.Language = &v
	return s
}

func (s *DeleteSequenceRequest) SetAppType(v string) *DeleteSequenceRequest {
	s.AppType = &v
	return s
}

type DeleteSequenceResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteSequenceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSequenceResponse) GoString() string {
	return s.String()
}

func (s *DeleteSequenceResponse) SetHeaders(v map[string]*string) *DeleteSequenceResponse {
	s.Headers = v
	return s
}

type ReleaseCommodityHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ReleaseCommodityHeaders) String() string {
	return tea.Prettify(s)
}

func (s ReleaseCommodityHeaders) GoString() string {
	return s.String()
}

func (s *ReleaseCommodityHeaders) SetCommonHeaders(v map[string]*string) *ReleaseCommodityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ReleaseCommodityHeaders) SetXAcsDingtalkAccessToken(v string) *ReleaseCommodityHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ReleaseCommodityRequest struct {
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	AccessKey  *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	CallerUid  *string `json:"callerUid,omitempty" xml:"callerUid,omitempty"`
}

func (s ReleaseCommodityRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseCommodityRequest) GoString() string {
	return s.String()
}

func (s *ReleaseCommodityRequest) SetInstanceId(v string) *ReleaseCommodityRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseCommodityRequest) SetAccessKey(v string) *ReleaseCommodityRequest {
	s.AccessKey = &v
	return s
}

func (s *ReleaseCommodityRequest) SetCallerUid(v string) *ReleaseCommodityRequest {
	s.CallerUid = &v
	return s
}

type ReleaseCommodityResponseBody struct {
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ReleaseCommodityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseCommodityResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseCommodityResponseBody) SetMessage(v string) *ReleaseCommodityResponseBody {
	s.Message = &v
	return s
}

func (s *ReleaseCommodityResponseBody) SetSuccess(v bool) *ReleaseCommodityResponseBody {
	s.Success = &v
	return s
}

type ReleaseCommodityResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleaseCommodityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseCommodityResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseCommodityResponse) GoString() string {
	return s.String()
}

func (s *ReleaseCommodityResponse) SetHeaders(v map[string]*string) *ReleaseCommodityResponse {
	s.Headers = v
	return s
}

func (s *ReleaseCommodityResponse) SetBody(v *ReleaseCommodityResponseBody) *ReleaseCommodityResponse {
	s.Body = v
	return s
}

type LoginCodeGenHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s LoginCodeGenHeaders) String() string {
	return tea.Prettify(s)
}

func (s LoginCodeGenHeaders) GoString() string {
	return s.String()
}

func (s *LoginCodeGenHeaders) SetCommonHeaders(v map[string]*string) *LoginCodeGenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *LoginCodeGenHeaders) SetXAcsDingtalkAccessToken(v string) *LoginCodeGenHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type LoginCodeGenRequest struct {
	// userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s LoginCodeGenRequest) String() string {
	return tea.Prettify(s)
}

func (s LoginCodeGenRequest) GoString() string {
	return s.String()
}

func (s *LoginCodeGenRequest) SetUserId(v string) *LoginCodeGenRequest {
	s.UserId = &v
	return s
}

type LoginCodeGenResponseBody struct {
	// result
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s LoginCodeGenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LoginCodeGenResponseBody) GoString() string {
	return s.String()
}

func (s *LoginCodeGenResponseBody) SetResult(v string) *LoginCodeGenResponseBody {
	s.Result = &v
	return s
}

type LoginCodeGenResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *LoginCodeGenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LoginCodeGenResponse) String() string {
	return tea.Prettify(s)
}

func (s LoginCodeGenResponse) GoString() string {
	return s.String()
}

func (s *LoginCodeGenResponse) SetHeaders(v map[string]*string) *LoginCodeGenResponse {
	s.Headers = v
	return s
}

func (s *LoginCodeGenResponse) SetBody(v *LoginCodeGenResponseBody) *LoginCodeGenResponse {
	s.Body = v
	return s
}

type GetSaleUserInfoByUserIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSaleUserInfoByUserIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSaleUserInfoByUserIdHeaders) GoString() string {
	return s.String()
}

func (s *GetSaleUserInfoByUserIdHeaders) SetCommonHeaders(v map[string]*string) *GetSaleUserInfoByUserIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSaleUserInfoByUserIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetSaleUserInfoByUserIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSaleUserInfoByUserIdRequest struct {
	CorpId    *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	UserId    *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetSaleUserInfoByUserIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSaleUserInfoByUserIdRequest) GoString() string {
	return s.String()
}

func (s *GetSaleUserInfoByUserIdRequest) SetCorpId(v string) *GetSaleUserInfoByUserIdRequest {
	s.CorpId = &v
	return s
}

func (s *GetSaleUserInfoByUserIdRequest) SetNamespace(v string) *GetSaleUserInfoByUserIdRequest {
	s.Namespace = &v
	return s
}

func (s *GetSaleUserInfoByUserIdRequest) SetUserId(v string) *GetSaleUserInfoByUserIdRequest {
	s.UserId = &v
	return s
}

type GetSaleUserInfoByUserIdResponseBody struct {
	// userName
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// accountId
	AccountId *int64 `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// corpList
	CorpList []*GetSaleUserInfoByUserIdResponseBodyCorpList `json:"corpList,omitempty" xml:"corpList,omitempty" type:"Repeated"`
}

func (s GetSaleUserInfoByUserIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSaleUserInfoByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetSaleUserInfoByUserIdResponseBody) SetUserName(v string) *GetSaleUserInfoByUserIdResponseBody {
	s.UserName = &v
	return s
}

func (s *GetSaleUserInfoByUserIdResponseBody) SetUserId(v string) *GetSaleUserInfoByUserIdResponseBody {
	s.UserId = &v
	return s
}

func (s *GetSaleUserInfoByUserIdResponseBody) SetAccountId(v int64) *GetSaleUserInfoByUserIdResponseBody {
	s.AccountId = &v
	return s
}

func (s *GetSaleUserInfoByUserIdResponseBody) SetCorpList(v []*GetSaleUserInfoByUserIdResponseBodyCorpList) *GetSaleUserInfoByUserIdResponseBody {
	s.CorpList = v
	return s
}

type GetSaleUserInfoByUserIdResponseBodyCorpList struct {
	// namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// corpId
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// corpName
	CorpName *string `json:"corpName,omitempty" xml:"corpName,omitempty"`
}

func (s GetSaleUserInfoByUserIdResponseBodyCorpList) String() string {
	return tea.Prettify(s)
}

func (s GetSaleUserInfoByUserIdResponseBodyCorpList) GoString() string {
	return s.String()
}

func (s *GetSaleUserInfoByUserIdResponseBodyCorpList) SetNamespace(v string) *GetSaleUserInfoByUserIdResponseBodyCorpList {
	s.Namespace = &v
	return s
}

func (s *GetSaleUserInfoByUserIdResponseBodyCorpList) SetCorpId(v string) *GetSaleUserInfoByUserIdResponseBodyCorpList {
	s.CorpId = &v
	return s
}

func (s *GetSaleUserInfoByUserIdResponseBodyCorpList) SetCorpName(v string) *GetSaleUserInfoByUserIdResponseBodyCorpList {
	s.CorpName = &v
	return s
}

type GetSaleUserInfoByUserIdResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSaleUserInfoByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSaleUserInfoByUserIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSaleUserInfoByUserIdResponse) GoString() string {
	return s.String()
}

func (s *GetSaleUserInfoByUserIdResponse) SetHeaders(v map[string]*string) *GetSaleUserInfoByUserIdResponse {
	s.Headers = v
	return s
}

func (s *GetSaleUserInfoByUserIdResponse) SetBody(v *GetSaleUserInfoByUserIdResponseBody) *GetSaleUserInfoByUserIdResponse {
	s.Body = v
	return s
}

type ExecuteTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ExecuteTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTaskHeaders) GoString() string {
	return s.String()
}

func (s *ExecuteTaskHeaders) SetCommonHeaders(v map[string]*string) *ExecuteTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ExecuteTaskHeaders) SetXAcsDingtalkAccessToken(v string) *ExecuteTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ExecuteTaskRequest struct {
	// 审批结果
	OutResult *string `json:"outResult,omitempty" xml:"outResult,omitempty"`
	// 是否不执行校验&关联操作
	NoExecuteExpressions *string `json:"noExecuteExpressions,omitempty" xml:"noExecuteExpressions,omitempty"`
	// 应用ID
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 更新的表单值
	FormDataJson *string `json:"formDataJson,omitempty" xml:"formDataJson,omitempty"`
	// 应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 语言
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 审批意见
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 实例ID
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// 钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 任务ID
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ExecuteTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTaskRequest) GoString() string {
	return s.String()
}

func (s *ExecuteTaskRequest) SetOutResult(v string) *ExecuteTaskRequest {
	s.OutResult = &v
	return s
}

func (s *ExecuteTaskRequest) SetNoExecuteExpressions(v string) *ExecuteTaskRequest {
	s.NoExecuteExpressions = &v
	return s
}

func (s *ExecuteTaskRequest) SetAppType(v string) *ExecuteTaskRequest {
	s.AppType = &v
	return s
}

func (s *ExecuteTaskRequest) SetFormDataJson(v string) *ExecuteTaskRequest {
	s.FormDataJson = &v
	return s
}

func (s *ExecuteTaskRequest) SetSystemToken(v string) *ExecuteTaskRequest {
	s.SystemToken = &v
	return s
}

func (s *ExecuteTaskRequest) SetLanguage(v string) *ExecuteTaskRequest {
	s.Language = &v
	return s
}

func (s *ExecuteTaskRequest) SetRemark(v string) *ExecuteTaskRequest {
	s.Remark = &v
	return s
}

func (s *ExecuteTaskRequest) SetProcessInstanceId(v string) *ExecuteTaskRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *ExecuteTaskRequest) SetUserId(v string) *ExecuteTaskRequest {
	s.UserId = &v
	return s
}

func (s *ExecuteTaskRequest) SetTaskId(v int64) *ExecuteTaskRequest {
	s.TaskId = &v
	return s
}

type ExecuteTaskResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ExecuteTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTaskResponse) GoString() string {
	return s.String()
}

func (s *ExecuteTaskResponse) SetHeaders(v map[string]*string) *ExecuteTaskResponse {
	s.Headers = v
	return s
}

type StartInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s StartInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceHeaders) GoString() string {
	return s.String()
}

func (s *StartInstanceHeaders) SetCommonHeaders(v map[string]*string) *StartInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *StartInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type StartInstanceRequest struct {
	// 应用编码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 应用秘钥。在应用数据中获取。
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 语言。可选值：zh_CN/en_US 默认：zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 表单唯一编码
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 表单数据
	FormDataJson *string `json:"formDataJson,omitempty" xml:"formDataJson,omitempty"`
	// 流程编码
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// 发起人所在部门编号
	DepartmentId *string `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
}

func (s StartInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartInstanceRequest) SetAppType(v string) *StartInstanceRequest {
	s.AppType = &v
	return s
}

func (s *StartInstanceRequest) SetSystemToken(v string) *StartInstanceRequest {
	s.SystemToken = &v
	return s
}

func (s *StartInstanceRequest) SetUserId(v string) *StartInstanceRequest {
	s.UserId = &v
	return s
}

func (s *StartInstanceRequest) SetLanguage(v string) *StartInstanceRequest {
	s.Language = &v
	return s
}

func (s *StartInstanceRequest) SetFormUuid(v string) *StartInstanceRequest {
	s.FormUuid = &v
	return s
}

func (s *StartInstanceRequest) SetFormDataJson(v string) *StartInstanceRequest {
	s.FormDataJson = &v
	return s
}

func (s *StartInstanceRequest) SetProcessCode(v string) *StartInstanceRequest {
	s.ProcessCode = &v
	return s
}

func (s *StartInstanceRequest) SetDepartmentId(v string) *StartInstanceRequest {
	s.DepartmentId = &v
	return s
}

type StartInstanceResponseBody struct {
	// 流程实例ID
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s StartInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstanceResponseBody) SetResult(v string) *StartInstanceResponseBody {
	s.Result = &v
	return s
}

type StartInstanceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartInstanceResponse) SetHeaders(v map[string]*string) *StartInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartInstanceResponse) SetBody(v *StartInstanceResponseBody) *StartInstanceResponse {
	s.Body = v
	return s
}

type DeleteInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceHeaders) GoString() string {
	return s.String()
}

func (s *DeleteInstanceHeaders) SetCommonHeaders(v map[string]*string) *DeleteInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteInstanceRequest struct {
	// 应用ID
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 语言
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 流程实例ID
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) SetAppType(v string) *DeleteInstanceRequest {
	s.AppType = &v
	return s
}

func (s *DeleteInstanceRequest) SetSystemToken(v string) *DeleteInstanceRequest {
	s.SystemToken = &v
	return s
}

func (s *DeleteInstanceRequest) SetUserId(v string) *DeleteInstanceRequest {
	s.UserId = &v
	return s
}

func (s *DeleteInstanceRequest) SetLanguage(v string) *DeleteInstanceRequest {
	s.Language = &v
	return s
}

func (s *DeleteInstanceRequest) SetProcessInstanceId(v string) *DeleteInstanceRequest {
	s.ProcessInstanceId = &v
	return s
}

type DeleteInstanceResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponse) SetHeaders(v map[string]*string) *DeleteInstanceResponse {
	s.Headers = v
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

func (client *Client) ValidateOrderUpgrade(request *ValidateOrderUpgradeRequest) (_result *ValidateOrderUpgradeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ValidateOrderUpgradeHeaders{}
	_result = &ValidateOrderUpgradeResponse{}
	_body, _err := client.ValidateOrderUpgradeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateOrderUpgradeWithOptions(request *ValidateOrderUpgradeRequest, headers *ValidateOrderUpgradeHeaders, runtime *util.RuntimeOptions) (_result *ValidateOrderUpgradeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUid)) {
		query["callerUid"] = request.CallerUid
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
		Query:   openapiutil.Query(query),
	}
	_result = &ValidateOrderUpgradeResponse{}
	_body, _err := client.DoROARequest(tea.String("ValidateOrderUpgrade"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/apps/orderUpgrade/validate"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCorpLevelByAccountId(request *GetCorpLevelByAccountIdRequest) (_result *GetCorpLevelByAccountIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCorpLevelByAccountIdHeaders{}
	_result = &GetCorpLevelByAccountIdResponse{}
	_body, _err := client.GetCorpLevelByAccountIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCorpLevelByAccountIdWithOptions(request *GetCorpLevelByAccountIdRequest, headers *GetCorpLevelByAccountIdHeaders, runtime *util.RuntimeOptions) (_result *GetCorpLevelByAccountIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["accountId"] = request.AccountId
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
		Query:   openapiutil.Query(query),
	}
	_result = &GetCorpLevelByAccountIdResponse{}
	_body, _err := client.DoROARequest(tea.String("GetCorpLevelByAccountId"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/apps/corpLevel"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateStatus(request *UpdateStatusRequest) (_result *UpdateStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateStatusHeaders{}
	_result = &UpdateStatusResponse{}
	_body, _err := client.UpdateStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateStatusWithOptions(request *UpdateStatusRequest, headers *UpdateStatusHeaders, runtime *util.RuntimeOptions) (_result *UpdateStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImportSequence)) {
		body["importSequence"] = request.ImportSequence
	}

	if !tea.BoolValue(util.IsUnset(request.ErrorLines)) {
		body["errorLines"] = request.ErrorLines
	}

	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
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
	_result = &UpdateStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateStatus"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/yida/forms/status"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecutePlatformTask(request *ExecutePlatformTaskRequest) (_result *ExecutePlatformTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ExecutePlatformTaskHeaders{}
	_result = &ExecutePlatformTaskResponse{}
	_body, _err := client.ExecutePlatformTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecutePlatformTaskWithOptions(request *ExecutePlatformTaskRequest, headers *ExecutePlatformTaskHeaders, runtime *util.RuntimeOptions) (_result *ExecutePlatformTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutResult)) {
		body["outResult"] = request.OutResult
	}

	if !tea.BoolValue(util.IsUnset(request.NoExecuteExpressions)) {
		body["noExecuteExpressions"] = request.NoExecuteExpressions
	}

	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.FormDataJson)) {
		body["formDataJson"] = request.FormDataJson
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		body["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
	_result = &ExecutePlatformTaskResponse{}
	_body, _err := client.DoROARequest(tea.String("ExecutePlatformTask"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/tasks/platformTasks/execute"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveFormRemark(request *SaveFormRemarkRequest) (_result *SaveFormRemarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveFormRemarkHeaders{}
	_result = &SaveFormRemarkResponse{}
	_body, _err := client.SaveFormRemarkWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveFormRemarkWithOptions(request *SaveFormRemarkRequest, headers *SaveFormRemarkHeaders, runtime *util.RuntimeOptions) (_result *SaveFormRemarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.ReplyId)) {
		body["replyId"] = request.ReplyId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.FormInstanceId)) {
		body["formInstanceId"] = request.FormInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.AtUserId)) {
		body["atUserId"] = request.AtUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
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
	_result = &SaveFormRemarkResponse{}
	_body, _err := client.DoROARequest(tea.String("SaveFormRemark"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/forms/remarks"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchFormDatas(request *SearchFormDatasRequest) (_result *SearchFormDatasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchFormDatasHeaders{}
	_result = &SearchFormDatasResponse{}
	_body, _err := client.SearchFormDatasWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchFormDatasWithOptions(request *SearchFormDatasRequest, headers *SearchFormDatasHeaders, runtime *util.RuntimeOptions) (_result *SearchFormDatasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.SearchFieldJson)) {
		body["searchFieldJson"] = request.SearchFieldJson
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["currentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.OriginatorId)) {
		body["originatorId"] = request.OriginatorId
	}

	if !tea.BoolValue(util.IsUnset(request.CreateFromTimeGMT)) {
		body["createFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CreateToTimeGMT)) {
		body["createToTimeGMT"] = request.CreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedFromTimeGMT)) {
		body["modifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedToTimeGMT)) {
		body["modifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.DynamicOrder)) {
		body["dynamicOrder"] = request.DynamicOrder
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
	_result = &SearchFormDatasResponse{}
	_body, _err := client.DoROARequest(tea.String("SearchFormDatas"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/forms/instances/search"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchActivationCode(request *SearchActivationCodeRequest) (_result *SearchActivationCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchActivationCodeHeaders{}
	_result = &SearchActivationCodeResponse{}
	_body, _err := client.SearchActivationCodeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchActivationCodeWithOptions(request *SearchActivationCodeRequest, headers *SearchActivationCodeHeaders, runtime *util.RuntimeOptions) (_result *SearchActivationCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUid)) {
		query["callerUid"] = request.CallerUid
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
		Query:   openapiutil.Query(query),
	}
	_result = &SearchActivationCodeResponse{}
	_body, _err := client.DoROARequest(tea.String("SearchActivationCode"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/apps/activationCode/information"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchEmployeeFieldValues(request *SearchEmployeeFieldValuesRequest) (_result *SearchEmployeeFieldValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchEmployeeFieldValuesHeaders{}
	_result = &SearchEmployeeFieldValuesResponse{}
	_body, _err := client.SearchEmployeeFieldValuesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchEmployeeFieldValuesWithOptions(request *SearchEmployeeFieldValuesRequest, headers *SearchEmployeeFieldValuesHeaders, runtime *util.RuntimeOptions) (_result *SearchEmployeeFieldValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TargetFieldJson)) {
		body["targetFieldJson"] = request.TargetFieldJson
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedToTimeGMT)) {
		body["modifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedFromTimeGMT)) {
		body["modifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.SearchFieldJson)) {
		body["searchFieldJson"] = request.SearchFieldJson
	}

	if !tea.BoolValue(util.IsUnset(request.OriginatorId)) {
		body["originatorId"] = request.OriginatorId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.CreateToTimeGMT)) {
		body["createToTimeGMT"] = request.CreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CreateFromTimeGMT)) {
		body["createFromTimeGMT"] = request.CreateFromTimeGMT
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
	_result = &SearchEmployeeFieldValuesResponse{}
	_body, _err := client.DoROARequest(tea.String("SearchEmployeeFieldValues"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/forms/employeeFields"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFormData(request *UpdateFormDataRequest) (_result *UpdateFormDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateFormDataHeaders{}
	_result = &UpdateFormDataResponse{}
	_body, _err := client.UpdateFormDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFormDataWithOptions(request *UpdateFormDataRequest, headers *UpdateFormDataHeaders, runtime *util.RuntimeOptions) (_result *UpdateFormDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.FormInstanceId)) {
		body["formInstanceId"] = request.FormInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UseLatestVersion)) {
		body["useLatestVersion"] = request.UseLatestVersion
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateFormDataJson)) {
		body["updateFormDataJson"] = request.UpdateFormDataJson
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
	_result = &UpdateFormDataResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateFormData"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/yida/forms/instances"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOperationRecords(request *GetOperationRecordsRequest) (_result *GetOperationRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOperationRecordsHeaders{}
	_result = &GetOperationRecordsResponse{}
	_body, _err := client.GetOperationRecordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOperationRecordsWithOptions(request *GetOperationRecordsRequest, headers *GetOperationRecordsHeaders, runtime *util.RuntimeOptions) (_result *GetOperationRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		query["processInstanceId"] = request.ProcessInstanceId
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
		Query:   openapiutil.Query(query),
	}
	_result = &GetOperationRecordsResponse{}
	_body, _err := client.DoROARequest(tea.String("GetOperationRecords"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/processes/operationRecords"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPlatformResource(request *GetPlatformResourceRequest) (_result *GetPlatformResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPlatformResourceHeaders{}
	_result = &GetPlatformResourceResponse{}
	_body, _err := client.GetPlatformResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPlatformResourceWithOptions(request *GetPlatformResourceRequest, headers *GetPlatformResourceHeaders, runtime *util.RuntimeOptions) (_result *GetPlatformResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUid)) {
		query["callerUid"] = request.CallerUid
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
		Query:   openapiutil.Query(query),
	}
	_result = &GetPlatformResourceResponse{}
	_body, _err := client.DoROARequest(tea.String("GetPlatformResource"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/apps/platformResources"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRunningTasks(request *GetRunningTasksRequest) (_result *GetRunningTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetRunningTasksHeaders{}
	_result = &GetRunningTasksResponse{}
	_body, _err := client.GetRunningTasksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRunningTasksWithOptions(request *GetRunningTasksRequest, headers *GetRunningTasksHeaders, runtime *util.RuntimeOptions) (_result *GetRunningTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		query["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
		Query:   openapiutil.Query(query),
	}
	_result = &GetRunningTasksResponse{}
	_body, _err := client.DoROARequest(tea.String("GetRunningTasks"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/processes/tasks/getRunningTasks"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNavigationByFormType(request *ListNavigationByFormTypeRequest) (_result *ListNavigationByFormTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListNavigationByFormTypeHeaders{}
	_result = &ListNavigationByFormTypeResponse{}
	_body, _err := client.ListNavigationByFormTypeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNavigationByFormTypeWithOptions(request *ListNavigationByFormTypeRequest, headers *ListNavigationByFormTypeHeaders, runtime *util.RuntimeOptions) (_result *ListNavigationByFormTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.FormType)) {
		query["formType"] = request.FormType
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
		Query:   openapiutil.Query(query),
	}
	_result = &ListNavigationByFormTypeResponse{}
	_body, _err := client.DoROARequest(tea.String("ListNavigationByFormType"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/apps/navigations"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TerminateInstance(request *TerminateInstanceRequest) (_result *TerminateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TerminateInstanceHeaders{}
	_result = &TerminateInstanceResponse{}
	_body, _err := client.TerminateInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TerminateInstanceWithOptions(request *TerminateInstanceRequest, headers *TerminateInstanceHeaders, runtime *util.RuntimeOptions) (_result *TerminateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		query["processInstanceId"] = request.ProcessInstanceId
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
		Query:   openapiutil.Query(query),
	}
	_result = &TerminateInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("TerminateInstance"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/yida/processes/instances/terminate"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExpireCommodity(request *ExpireCommodityRequest) (_result *ExpireCommodityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ExpireCommodityHeaders{}
	_result = &ExpireCommodityResponse{}
	_body, _err := client.ExpireCommodityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExpireCommodityWithOptions(request *ExpireCommodityRequest, headers *ExpireCommodityHeaders, runtime *util.RuntimeOptions) (_result *ExpireCommodityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUid)) {
		query["callerUid"] = request.CallerUid
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
		Query:   openapiutil.Query(query),
	}
	_result = &ExpireCommodityResponse{}
	_body, _err := client.DoROARequest(tea.String("ExpireCommodity"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/yida/appAuth/commodities/expire"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateOrderBuy(request *ValidateOrderBuyRequest) (_result *ValidateOrderBuyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ValidateOrderBuyHeaders{}
	_result = &ValidateOrderBuyResponse{}
	_body, _err := client.ValidateOrderBuyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateOrderBuyWithOptions(request *ValidateOrderBuyRequest, headers *ValidateOrderBuyHeaders, runtime *util.RuntimeOptions) (_result *ValidateOrderBuyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUid)) {
		query["callerUid"] = request.CallerUid
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
		Query:   openapiutil.Query(query),
	}
	_result = &ValidateOrderBuyResponse{}
	_body, _err := client.DoROARequest(tea.String("ValidateOrderBuy"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/apps/orderBuy/validate"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveFormData(request *SaveFormDataRequest) (_result *SaveFormDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveFormDataHeaders{}
	_result = &SaveFormDataResponse{}
	_body, _err := client.SaveFormDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveFormDataWithOptions(request *SaveFormDataRequest, headers *SaveFormDataHeaders, runtime *util.RuntimeOptions) (_result *SaveFormDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.FormDataJson)) {
		body["formDataJson"] = request.FormDataJson
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
	_result = &SaveFormDataResponse{}
	_body, _err := client.DoROARequest(tea.String("SaveFormData"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/forms/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFormData(request *DeleteFormDataRequest) (_result *DeleteFormDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteFormDataHeaders{}
	_result = &DeleteFormDataResponse{}
	_body, _err := client.DeleteFormDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFormDataWithOptions(request *DeleteFormDataRequest, headers *DeleteFormDataHeaders, runtime *util.RuntimeOptions) (_result *DeleteFormDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.FormInstanceId)) {
		query["formInstanceId"] = request.FormInstanceId
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
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteFormDataResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteFormData"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/yida/forms/instances"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInstance(request *UpdateInstanceRequest) (_result *UpdateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateInstanceHeaders{}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.UpdateInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInstanceWithOptions(request *UpdateInstanceRequest, headers *UpdateInstanceHeaders, runtime *util.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		body["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateFormDataJson)) {
		body["updateFormDataJson"] = request.UpdateFormDataJson
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
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
	_result = &UpdateInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateInstance"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/yida/processes/instances"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCommodity(request *ListCommodityRequest) (_result *ListCommodityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListCommodityHeaders{}
	_result = &ListCommodityResponse{}
	_body, _err := client.ListCommodityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCommodityWithOptions(request *ListCommodityRequest, headers *ListCommodityHeaders, runtime *util.RuntimeOptions) (_result *ListCommodityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUid)) {
		query["callerUid"] = request.CallerUid
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["currentPage"] = request.CurrentPage
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
		Query:   openapiutil.Query(query),
	}
	_result = &ListCommodityResponse{}
	_body, _err := client.DoROARequest(tea.String("ListCommodity"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/appAuth/commodities"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetApplicationAuthorizationServicePlatformResource(request *GetApplicationAuthorizationServicePlatformResourceRequest) (_result *GetApplicationAuthorizationServicePlatformResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetApplicationAuthorizationServicePlatformResourceHeaders{}
	_result = &GetApplicationAuthorizationServicePlatformResourceResponse{}
	_body, _err := client.GetApplicationAuthorizationServicePlatformResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetApplicationAuthorizationServicePlatformResourceWithOptions(request *GetApplicationAuthorizationServicePlatformResourceRequest, headers *GetApplicationAuthorizationServicePlatformResourceHeaders, runtime *util.RuntimeOptions) (_result *GetApplicationAuthorizationServicePlatformResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUid)) {
		query["callerUid"] = request.CallerUid
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
		Query:   openapiutil.Query(query),
	}
	_result = &GetApplicationAuthorizationServicePlatformResourceResponse{}
	_body, _err := client.DoROARequest(tea.String("GetApplicationAuthorizationServicePlatformResource"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/authorization/platformResources"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetActivityList(request *GetActivityListRequest) (_result *GetActivityListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetActivityListHeaders{}
	_result = &GetActivityListResponse{}
	_body, _err := client.GetActivityListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetActivityListWithOptions(request *GetActivityListRequest, headers *GetActivityListHeaders, runtime *util.RuntimeOptions) (_result *GetActivityListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		query["processCode"] = request.ProcessCode
	}

	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
		Query:   openapiutil.Query(query),
	}
	_result = &GetActivityListResponse{}
	_body, _err := client.DoROARequest(tea.String("GetActivityList"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/processes/activities"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFormDataByID(id *string, request *GetFormDataByIDRequest) (_result *GetFormDataByIDResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFormDataByIDHeaders{}
	_result = &GetFormDataByIDResponse{}
	_body, _err := client.GetFormDataByIDWithOptions(id, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFormDataByIDWithOptions(id *string, request *GetFormDataByIDRequest, headers *GetFormDataByIDHeaders, runtime *util.RuntimeOptions) (_result *GetFormDataByIDResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	id = openapiutil.GetEncodeParam(id)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
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
		Query:   openapiutil.Query(query),
	}
	_result = &GetFormDataByIDResponse{}
	_body, _err := client.DoROARequest(tea.String("GetFormDataByID"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/forms/instances/"+tea.StringValue(id)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteCustomApi(request *ExecuteCustomApiRequest) (_result *ExecuteCustomApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ExecuteCustomApiHeaders{}
	_result = &ExecuteCustomApiResponse{}
	_body, _err := client.ExecuteCustomApiWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteCustomApiWithOptions(request *ExecuteCustomApiRequest, headers *ExecuteCustomApiHeaders, runtime *util.RuntimeOptions) (_result *ExecuteCustomApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		query["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["serviceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
		Query:   openapiutil.Query(query),
	}
	_result = &ExecuteCustomApiResponse{}
	_body, _err := client.DoROARequest(tea.String("ExecuteCustomApi"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/apps/customApi/execute"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefundCommodity(request *RefundCommodityRequest) (_result *RefundCommodityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RefundCommodityHeaders{}
	_result = &RefundCommodityResponse{}
	_body, _err := client.RefundCommodityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefundCommodityWithOptions(request *RefundCommodityRequest, headers *RefundCommodityHeaders, runtime *util.RuntimeOptions) (_result *RefundCommodityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUid)) {
		query["callerUid"] = request.CallerUid
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
		Query:   openapiutil.Query(query),
	}
	_result = &RefundCommodityResponse{}
	_body, _err := client.DoROARequest(tea.String("RefundCommodity"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/appAuth/commodities/refund"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSequence(request *DeleteSequenceRequest) (_result *DeleteSequenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteSequenceHeaders{}
	_result = &DeleteSequenceResponse{}
	_body, _err := client.DeleteSequenceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSequenceWithOptions(request *DeleteSequenceRequest, headers *DeleteSequenceHeaders, runtime *util.RuntimeOptions) (_result *DeleteSequenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Sequence)) {
		query["sequence"] = request.Sequence
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
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
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteSequenceResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteSequence"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/yida/forms/deleteSequence"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseCommodity(request *ReleaseCommodityRequest) (_result *ReleaseCommodityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ReleaseCommodityHeaders{}
	_result = &ReleaseCommodityResponse{}
	_body, _err := client.ReleaseCommodityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseCommodityWithOptions(request *ReleaseCommodityRequest, headers *ReleaseCommodityHeaders, runtime *util.RuntimeOptions) (_result *ReleaseCommodityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUid)) {
		query["callerUid"] = request.CallerUid
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
		Query:   openapiutil.Query(query),
	}
	_result = &ReleaseCommodityResponse{}
	_body, _err := client.DoROARequest(tea.String("ReleaseCommodity"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/yida/appAuth/commodities/release"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LoginCodeGen(request *LoginCodeGenRequest) (_result *LoginCodeGenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &LoginCodeGenHeaders{}
	_result = &LoginCodeGenResponse{}
	_body, _err := client.LoginCodeGenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LoginCodeGenWithOptions(request *LoginCodeGenRequest, headers *LoginCodeGenHeaders, runtime *util.RuntimeOptions) (_result *LoginCodeGenResponse, _err error) {
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
		realHeaders["x-acs-dingtalk-access-token"] = headers.XAcsDingtalkAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &LoginCodeGenResponse{}
	_body, _err := client.DoROARequest(tea.String("LoginCodeGen"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/authorizations/loginCodes"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSaleUserInfoByUserId(request *GetSaleUserInfoByUserIdRequest) (_result *GetSaleUserInfoByUserIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSaleUserInfoByUserIdHeaders{}
	_result = &GetSaleUserInfoByUserIdResponse{}
	_body, _err := client.GetSaleUserInfoByUserIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSaleUserInfoByUserIdWithOptions(request *GetSaleUserInfoByUserIdRequest, headers *GetSaleUserInfoByUserIdHeaders, runtime *util.RuntimeOptions) (_result *GetSaleUserInfoByUserIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
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
		Query:   openapiutil.Query(query),
	}
	_result = &GetSaleUserInfoByUserIdResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSaleUserInfoByUserId"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/apps/saleUserInfo"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteTask(request *ExecuteTaskRequest) (_result *ExecuteTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ExecuteTaskHeaders{}
	_result = &ExecuteTaskResponse{}
	_body, _err := client.ExecuteTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteTaskWithOptions(request *ExecuteTaskRequest, headers *ExecuteTaskHeaders, runtime *util.RuntimeOptions) (_result *ExecuteTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutResult)) {
		body["outResult"] = request.OutResult
	}

	if !tea.BoolValue(util.IsUnset(request.NoExecuteExpressions)) {
		body["noExecuteExpressions"] = request.NoExecuteExpressions
	}

	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.FormDataJson)) {
		body["formDataJson"] = request.FormDataJson
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		body["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
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
	_result = &ExecuteTaskResponse{}
	_body, _err := client.DoROARequest(tea.String("ExecuteTask"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/tasks/execute"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartInstance(request *StartInstanceRequest) (_result *StartInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StartInstanceHeaders{}
	_result = &StartInstanceResponse{}
	_body, _err := client.StartInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartInstanceWithOptions(request *StartInstanceRequest, headers *StartInstanceHeaders, runtime *util.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.FormDataJson)) {
		body["formDataJson"] = request.FormDataJson
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		body["processCode"] = request.ProcessCode
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		body["departmentId"] = request.DepartmentId
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
	_result = &StartInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("StartInstance"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/processes/instances/start"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteInstanceHeaders{}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, headers *DeleteInstanceHeaders, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		query["processInstanceId"] = request.ProcessInstanceId
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
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteInstance"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/yida/processes/instances"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
