// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package yida_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BatchGetFormDataByIdListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchGetFormDataByIdListHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchGetFormDataByIdListHeaders) GoString() string {
	return s.String()
}

func (s *BatchGetFormDataByIdListHeaders) SetCommonHeaders(v map[string]*string) *BatchGetFormDataByIdListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchGetFormDataByIdListHeaders) SetXAcsDingtalkAccessToken(v string) *BatchGetFormDataByIdListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchGetFormDataByIdListRequest struct {
	// 宜搭应用编码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 表单实例id列表
	FormInstanceIdList []*string `json:"formInstanceIdList,omitempty" xml:"formInstanceIdList,omitempty" type:"Repeated"`
	// 表单编码
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 是否需要宜搭表单组件格式的实例数据
	NeedFormInstanceValue *bool `json:"needFormInstanceValue,omitempty" xml:"needFormInstanceValue,omitempty"`
	// 宜搭应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchGetFormDataByIdListRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetFormDataByIdListRequest) GoString() string {
	return s.String()
}

func (s *BatchGetFormDataByIdListRequest) SetAppType(v string) *BatchGetFormDataByIdListRequest {
	s.AppType = &v
	return s
}

func (s *BatchGetFormDataByIdListRequest) SetFormInstanceIdList(v []*string) *BatchGetFormDataByIdListRequest {
	s.FormInstanceIdList = v
	return s
}

func (s *BatchGetFormDataByIdListRequest) SetFormUuid(v string) *BatchGetFormDataByIdListRequest {
	s.FormUuid = &v
	return s
}

func (s *BatchGetFormDataByIdListRequest) SetNeedFormInstanceValue(v bool) *BatchGetFormDataByIdListRequest {
	s.NeedFormInstanceValue = &v
	return s
}

func (s *BatchGetFormDataByIdListRequest) SetSystemToken(v string) *BatchGetFormDataByIdListRequest {
	s.SystemToken = &v
	return s
}

func (s *BatchGetFormDataByIdListRequest) SetUserId(v string) *BatchGetFormDataByIdListRequest {
	s.UserId = &v
	return s
}

type BatchGetFormDataByIdListResponseBody struct {
	// 表单实例数据
	Result []*BatchGetFormDataByIdListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s BatchGetFormDataByIdListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchGetFormDataByIdListResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetFormDataByIdListResponseBody) SetResult(v []*BatchGetFormDataByIdListResponseBodyResult) *BatchGetFormDataByIdListResponseBody {
	s.Result = v
	return s
}

type BatchGetFormDataByIdListResponseBodyResult struct {
	// 创建时间
	CreateTimeGMT *string `json:"createTimeGMT,omitempty" xml:"createTimeGMT,omitempty"`
	// 创建者的userId
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// 表单实例数据
	FormData map[string]interface{} `json:"formData,omitempty" xml:"formData,omitempty"`
	// 表单实例id
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// 表单编码
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 数据库表记录主键id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 实例数据
	InstanceValue *string `json:"instanceValue,omitempty" xml:"instanceValue,omitempty"`
	// 修改时间
	ModifiedTimeGMT *string `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	// 修改者的钉钉userId
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// 修改者
	ModifyUser *BatchGetFormDataByIdListResponseBodyResultModifyUser `json:"modifyUser,omitempty" xml:"modifyUser,omitempty" type:"Struct"`
	// 表单提交人
	Originator *BatchGetFormDataByIdListResponseBodyResultOriginator `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	// 该表单实例对应的批量导入的批次号(如果是通过批量导入创建的)
	Sequence *string `json:"sequence,omitempty" xml:"sequence,omitempty"`
	// 流水号
	SerialNumber *string `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
	// 标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 该表单实例对应的表单schema版本
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s BatchGetFormDataByIdListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s BatchGetFormDataByIdListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetCreateTimeGMT(v string) *BatchGetFormDataByIdListResponseBodyResult {
	s.CreateTimeGMT = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetCreatorUserId(v string) *BatchGetFormDataByIdListResponseBodyResult {
	s.CreatorUserId = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetFormData(v map[string]interface{}) *BatchGetFormDataByIdListResponseBodyResult {
	s.FormData = v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetFormInstanceId(v string) *BatchGetFormDataByIdListResponseBodyResult {
	s.FormInstanceId = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetFormUuid(v string) *BatchGetFormDataByIdListResponseBodyResult {
	s.FormUuid = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetId(v int64) *BatchGetFormDataByIdListResponseBodyResult {
	s.Id = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetInstanceValue(v string) *BatchGetFormDataByIdListResponseBodyResult {
	s.InstanceValue = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetModifiedTimeGMT(v string) *BatchGetFormDataByIdListResponseBodyResult {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetModifier(v string) *BatchGetFormDataByIdListResponseBodyResult {
	s.Modifier = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetModifyUser(v *BatchGetFormDataByIdListResponseBodyResultModifyUser) *BatchGetFormDataByIdListResponseBodyResult {
	s.ModifyUser = v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetOriginator(v *BatchGetFormDataByIdListResponseBodyResultOriginator) *BatchGetFormDataByIdListResponseBodyResult {
	s.Originator = v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetSequence(v string) *BatchGetFormDataByIdListResponseBodyResult {
	s.Sequence = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetSerialNumber(v string) *BatchGetFormDataByIdListResponseBodyResult {
	s.SerialNumber = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetTitle(v string) *BatchGetFormDataByIdListResponseBodyResult {
	s.Title = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResult) SetVersion(v int64) *BatchGetFormDataByIdListResponseBodyResult {
	s.Version = &v
	return s
}

type BatchGetFormDataByIdListResponseBodyResultModifyUser struct {
	// 名称
	Name *BatchGetFormDataByIdListResponseBodyResultModifyUserName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchGetFormDataByIdListResponseBodyResultModifyUser) String() string {
	return tea.Prettify(s)
}

func (s BatchGetFormDataByIdListResponseBodyResultModifyUser) GoString() string {
	return s.String()
}

func (s *BatchGetFormDataByIdListResponseBodyResultModifyUser) SetName(v *BatchGetFormDataByIdListResponseBodyResultModifyUserName) *BatchGetFormDataByIdListResponseBodyResultModifyUser {
	s.Name = v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResultModifyUser) SetUserId(v string) *BatchGetFormDataByIdListResponseBodyResultModifyUser {
	s.UserId = &v
	return s
}

type BatchGetFormDataByIdListResponseBodyResultModifyUserName struct {
	// 中文名称
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// 英文名称
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
}

func (s BatchGetFormDataByIdListResponseBodyResultModifyUserName) String() string {
	return tea.Prettify(s)
}

func (s BatchGetFormDataByIdListResponseBodyResultModifyUserName) GoString() string {
	return s.String()
}

func (s *BatchGetFormDataByIdListResponseBodyResultModifyUserName) SetNameInChinese(v string) *BatchGetFormDataByIdListResponseBodyResultModifyUserName {
	s.NameInChinese = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResultModifyUserName) SetNameInEnglish(v string) *BatchGetFormDataByIdListResponseBodyResultModifyUserName {
	s.NameInEnglish = &v
	return s
}

type BatchGetFormDataByIdListResponseBodyResultOriginator struct {
	// 名称
	Name *BatchGetFormDataByIdListResponseBodyResultOriginatorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchGetFormDataByIdListResponseBodyResultOriginator) String() string {
	return tea.Prettify(s)
}

func (s BatchGetFormDataByIdListResponseBodyResultOriginator) GoString() string {
	return s.String()
}

func (s *BatchGetFormDataByIdListResponseBodyResultOriginator) SetName(v *BatchGetFormDataByIdListResponseBodyResultOriginatorName) *BatchGetFormDataByIdListResponseBodyResultOriginator {
	s.Name = v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResultOriginator) SetUserId(v string) *BatchGetFormDataByIdListResponseBodyResultOriginator {
	s.UserId = &v
	return s
}

type BatchGetFormDataByIdListResponseBodyResultOriginatorName struct {
	// 中文名称
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// 英文名称
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
}

func (s BatchGetFormDataByIdListResponseBodyResultOriginatorName) String() string {
	return tea.Prettify(s)
}

func (s BatchGetFormDataByIdListResponseBodyResultOriginatorName) GoString() string {
	return s.String()
}

func (s *BatchGetFormDataByIdListResponseBodyResultOriginatorName) SetNameInChinese(v string) *BatchGetFormDataByIdListResponseBodyResultOriginatorName {
	s.NameInChinese = &v
	return s
}

func (s *BatchGetFormDataByIdListResponseBodyResultOriginatorName) SetNameInEnglish(v string) *BatchGetFormDataByIdListResponseBodyResultOriginatorName {
	s.NameInEnglish = &v
	return s
}

type BatchGetFormDataByIdListResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchGetFormDataByIdListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchGetFormDataByIdListResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchGetFormDataByIdListResponse) GoString() string {
	return s.String()
}

func (s *BatchGetFormDataByIdListResponse) SetHeaders(v map[string]*string) *BatchGetFormDataByIdListResponse {
	s.Headers = v
	return s
}

func (s *BatchGetFormDataByIdListResponse) SetBody(v *BatchGetFormDataByIdListResponseBody) *BatchGetFormDataByIdListResponse {
	s.Body = v
	return s
}

type BatchRemovalByFormInstanceIdListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchRemovalByFormInstanceIdListHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchRemovalByFormInstanceIdListHeaders) GoString() string {
	return s.String()
}

func (s *BatchRemovalByFormInstanceIdListHeaders) SetCommonHeaders(v map[string]*string) *BatchRemovalByFormInstanceIdListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchRemovalByFormInstanceIdListHeaders) SetXAcsDingtalkAccessToken(v string) *BatchRemovalByFormInstanceIdListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchRemovalByFormInstanceIdListRequest struct {
	// 宜搭应用编码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 是否需要宜搭服务端异步执行该任务(选择异步执行删除操作，那么OpenAPI调用会立即返回，并且宜搭服务端会继续执行删除操作直至结束，且允许的单次删除数据量上限更大)
	AsynchronousExecution *bool `json:"asynchronousExecution,omitempty" xml:"asynchronousExecution,omitempty"`
	// 是否需要触发表单绑定的校验规则、关联业务规则和第三方服务回调（如果您的业务无必要执行这些，那么请填false以降低API的耗时以及获得更大的单次删除数据量上限）
	ExecuteExpression *bool `json:"executeExpression,omitempty" xml:"executeExpression,omitempty"`
	// 表单实例id列表
	FormInstanceIdList []*string `json:"formInstanceIdList,omitempty" xml:"formInstanceIdList,omitempty" type:"Repeated"`
	// 表单编码
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 宜搭应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchRemovalByFormInstanceIdListRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchRemovalByFormInstanceIdListRequest) GoString() string {
	return s.String()
}

func (s *BatchRemovalByFormInstanceIdListRequest) SetAppType(v string) *BatchRemovalByFormInstanceIdListRequest {
	s.AppType = &v
	return s
}

func (s *BatchRemovalByFormInstanceIdListRequest) SetAsynchronousExecution(v bool) *BatchRemovalByFormInstanceIdListRequest {
	s.AsynchronousExecution = &v
	return s
}

func (s *BatchRemovalByFormInstanceIdListRequest) SetExecuteExpression(v bool) *BatchRemovalByFormInstanceIdListRequest {
	s.ExecuteExpression = &v
	return s
}

func (s *BatchRemovalByFormInstanceIdListRequest) SetFormInstanceIdList(v []*string) *BatchRemovalByFormInstanceIdListRequest {
	s.FormInstanceIdList = v
	return s
}

func (s *BatchRemovalByFormInstanceIdListRequest) SetFormUuid(v string) *BatchRemovalByFormInstanceIdListRequest {
	s.FormUuid = &v
	return s
}

func (s *BatchRemovalByFormInstanceIdListRequest) SetSystemToken(v string) *BatchRemovalByFormInstanceIdListRequest {
	s.SystemToken = &v
	return s
}

func (s *BatchRemovalByFormInstanceIdListRequest) SetUserId(v string) *BatchRemovalByFormInstanceIdListRequest {
	s.UserId = &v
	return s
}

type BatchRemovalByFormInstanceIdListResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s BatchRemovalByFormInstanceIdListResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchRemovalByFormInstanceIdListResponse) GoString() string {
	return s.String()
}

func (s *BatchRemovalByFormInstanceIdListResponse) SetHeaders(v map[string]*string) *BatchRemovalByFormInstanceIdListResponse {
	s.Headers = v
	return s
}

type BatchSaveFormDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchSaveFormDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchSaveFormDataHeaders) GoString() string {
	return s.String()
}

func (s *BatchSaveFormDataHeaders) SetCommonHeaders(v map[string]*string) *BatchSaveFormDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchSaveFormDataHeaders) SetXAcsDingtalkAccessToken(v string) *BatchSaveFormDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchSaveFormDataRequest struct {
	// 宜搭应用编码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 是否需要宜搭服务端异步执行该任务(如果选择异步创建表单实例，那么OpenAPI调用会立即返回，并且宜搭服务端会继续执行保存操作直至结束，且允许的单次保存数据量上限更大)
	AsynchronousExecution *bool `json:"asynchronousExecution,omitempty" xml:"asynchronousExecution,omitempty"`
	// 表单实例数据列表。表单实例数据的结构请参考 https://www.yuque.com/yida/support/agb8im#f26a51f429f9f19aa0b5b3ee847ac556_h3_31
	FormDataJsonList []*string `json:"formDataJsonList,omitempty" xml:"formDataJsonList,omitempty" type:"Repeated"`
	// 表单编码
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 批量保存多条表单实例数据发生异常时是否跳过异常的表单实例并继续保存下一个表单实例数据。当noExecuteExpression为false时此参数才生效。
	KeepRunningAfterException *bool `json:"keepRunningAfterException,omitempty" xml:"keepRunningAfterException,omitempty"`
	// 是否不触发表单绑定的校验规则、关联业务规则和第三方服务回调（如果您的业务无必要执行这些，那么请填true以减小API的耗时以及获得更大的单次保存数据量上限）
	NoExecuteExpression *bool `json:"noExecuteExpression,omitempty" xml:"noExecuteExpression,omitempty"`
	// 宜搭应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchSaveFormDataRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSaveFormDataRequest) GoString() string {
	return s.String()
}

func (s *BatchSaveFormDataRequest) SetAppType(v string) *BatchSaveFormDataRequest {
	s.AppType = &v
	return s
}

func (s *BatchSaveFormDataRequest) SetAsynchronousExecution(v bool) *BatchSaveFormDataRequest {
	s.AsynchronousExecution = &v
	return s
}

func (s *BatchSaveFormDataRequest) SetFormDataJsonList(v []*string) *BatchSaveFormDataRequest {
	s.FormDataJsonList = v
	return s
}

func (s *BatchSaveFormDataRequest) SetFormUuid(v string) *BatchSaveFormDataRequest {
	s.FormUuid = &v
	return s
}

func (s *BatchSaveFormDataRequest) SetKeepRunningAfterException(v bool) *BatchSaveFormDataRequest {
	s.KeepRunningAfterException = &v
	return s
}

func (s *BatchSaveFormDataRequest) SetNoExecuteExpression(v bool) *BatchSaveFormDataRequest {
	s.NoExecuteExpression = &v
	return s
}

func (s *BatchSaveFormDataRequest) SetSystemToken(v string) *BatchSaveFormDataRequest {
	s.SystemToken = &v
	return s
}

func (s *BatchSaveFormDataRequest) SetUserId(v string) *BatchSaveFormDataRequest {
	s.UserId = &v
	return s
}

type BatchSaveFormDataResponseBody struct {
	// 新增的表单实例的id列表
	Result []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s BatchSaveFormDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSaveFormDataResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSaveFormDataResponseBody) SetResult(v []*string) *BatchSaveFormDataResponseBody {
	s.Result = v
	return s
}

type BatchSaveFormDataResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchSaveFormDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchSaveFormDataResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchSaveFormDataResponse) GoString() string {
	return s.String()
}

func (s *BatchSaveFormDataResponse) SetHeaders(v map[string]*string) *BatchSaveFormDataResponse {
	s.Headers = v
	return s
}

func (s *BatchSaveFormDataResponse) SetBody(v *BatchSaveFormDataResponseBody) *BatchSaveFormDataResponse {
	s.Body = v
	return s
}

type BatchUpdateFormDataByInstanceIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchUpdateFormDataByInstanceIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateFormDataByInstanceIdHeaders) GoString() string {
	return s.String()
}

func (s *BatchUpdateFormDataByInstanceIdHeaders) SetCommonHeaders(v map[string]*string) *BatchUpdateFormDataByInstanceIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdHeaders) SetXAcsDingtalkAccessToken(v string) *BatchUpdateFormDataByInstanceIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchUpdateFormDataByInstanceIdRequest struct {
	// 宜搭应用编码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 是否需要宜搭服务端异步执行该任务(选择异步执行那么OpenAPI调用会立即返回，并且任务会在宜搭服务端继续执行直至结束，且允许的单次更新数据量上限更大)
	AsynchronousExecution *bool `json:"asynchronousExecution,omitempty" xml:"asynchronousExecution,omitempty"`
	// 表单实例id列表
	FormInstanceIdList []*string `json:"formInstanceIdList,omitempty" xml:"formInstanceIdList,omitempty" type:"Repeated"`
	// 表单编码
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 是否忽略空值
	IgnoreEmpty *bool `json:"ignoreEmpty,omitempty" xml:"ignoreEmpty,omitempty"`
	// 是否不触发校验规则、关联业务规则和第三方服务回调（如果您的业务无必要触发这些那么请填true以增大单次更新允许的数据量上限以及API的执行速度）
	NoExecuteExpression *bool `json:"noExecuteExpression,omitempty" xml:"noExecuteExpression,omitempty"`
	// 宜搭应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 用于更新表单实例的数据, 格式为json字符串, 能解析成Map结构, 解析得到的Map的键为表单组件id, 值为表单组件值。详情参考 https://www.yuque.com/yida/support/agb8im#f26a51f429f9f19aa0b5b3ee847ac556_h3_31
	UpdateFormDataJson *string `json:"updateFormDataJson,omitempty" xml:"updateFormDataJson,omitempty"`
	// 是否使用最新的表单schema版本, 默认false
	UseLatestFormSchemaVersion *bool `json:"useLatestFormSchemaVersion,omitempty" xml:"useLatestFormSchemaVersion,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchUpdateFormDataByInstanceIdRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateFormDataByInstanceIdRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateFormDataByInstanceIdRequest) SetAppType(v string) *BatchUpdateFormDataByInstanceIdRequest {
	s.AppType = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdRequest) SetAsynchronousExecution(v bool) *BatchUpdateFormDataByInstanceIdRequest {
	s.AsynchronousExecution = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdRequest) SetFormInstanceIdList(v []*string) *BatchUpdateFormDataByInstanceIdRequest {
	s.FormInstanceIdList = v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdRequest) SetFormUuid(v string) *BatchUpdateFormDataByInstanceIdRequest {
	s.FormUuid = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdRequest) SetIgnoreEmpty(v bool) *BatchUpdateFormDataByInstanceIdRequest {
	s.IgnoreEmpty = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdRequest) SetNoExecuteExpression(v bool) *BatchUpdateFormDataByInstanceIdRequest {
	s.NoExecuteExpression = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdRequest) SetSystemToken(v string) *BatchUpdateFormDataByInstanceIdRequest {
	s.SystemToken = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdRequest) SetUpdateFormDataJson(v string) *BatchUpdateFormDataByInstanceIdRequest {
	s.UpdateFormDataJson = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdRequest) SetUseLatestFormSchemaVersion(v bool) *BatchUpdateFormDataByInstanceIdRequest {
	s.UseLatestFormSchemaVersion = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdRequest) SetUserId(v string) *BatchUpdateFormDataByInstanceIdRequest {
	s.UserId = &v
	return s
}

type BatchUpdateFormDataByInstanceIdResponseBody struct {
	// 成功更新的表单实例的id
	Result []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s BatchUpdateFormDataByInstanceIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateFormDataByInstanceIdResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdateFormDataByInstanceIdResponseBody) SetResult(v []*string) *BatchUpdateFormDataByInstanceIdResponseBody {
	s.Result = v
	return s
}

type BatchUpdateFormDataByInstanceIdResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchUpdateFormDataByInstanceIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchUpdateFormDataByInstanceIdResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateFormDataByInstanceIdResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateFormDataByInstanceIdResponse) SetHeaders(v map[string]*string) *BatchUpdateFormDataByInstanceIdResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdResponse) SetBody(v *BatchUpdateFormDataByInstanceIdResponseBody) *BatchUpdateFormDataByInstanceIdResponse {
	s.Body = v
	return s
}

type BatchUpdateFormDataByInstanceMapHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchUpdateFormDataByInstanceMapHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateFormDataByInstanceMapHeaders) GoString() string {
	return s.String()
}

func (s *BatchUpdateFormDataByInstanceMapHeaders) SetCommonHeaders(v map[string]*string) *BatchUpdateFormDataByInstanceMapHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapHeaders) SetXAcsDingtalkAccessToken(v string) *BatchUpdateFormDataByInstanceMapHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchUpdateFormDataByInstanceMapRequest struct {
	// 宜搭应用编码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 该任务是否需要服务端异步执行(选择异步执行那么OpenAPI调用会立即返回并且任务在宜搭服务端继续执行，可支持更大的单次更新数据量上限)
	AsynchronousExecution *bool `json:"asynchronousExecution,omitempty" xml:"asynchronousExecution,omitempty"`
	// 表单编码
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 是否忽略空值
	IgnoreEmpty *bool `json:"ignoreEmpty,omitempty" xml:"ignoreEmpty,omitempty"`
	// 是否不需要触发表单绑定的校验规则、关联业务规则和第三方服务回调（如果您的业务无必要执行这些，那么请填true以减小API的耗时以及更大的单次更新数据量上限）
	NoExecuteExpression *bool `json:"noExecuteExpression,omitempty" xml:"noExecuteExpression,omitempty"`
	// 宜搭应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 表单实例数据, json字符串, 可以解析成Map, 解析后得到的Map的键是表单实例id, 值是表单实例更新值json字符串。具体结构请参考 https://www.yuque.com/yida/support/agb8im#f26a51f429f9f19aa0b5b3ee847ac556_h3_31
	UpdateFormDataJsonMap map[string]interface{} `json:"updateFormDataJsonMap,omitempty" xml:"updateFormDataJsonMap,omitempty"`
	// 是否使用最新的表单schema版本, 默认false
	UseLatestFormSchemaVersion *bool `json:"useLatestFormSchemaVersion,omitempty" xml:"useLatestFormSchemaVersion,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchUpdateFormDataByInstanceMapRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateFormDataByInstanceMapRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateFormDataByInstanceMapRequest) SetAppType(v string) *BatchUpdateFormDataByInstanceMapRequest {
	s.AppType = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapRequest) SetAsynchronousExecution(v bool) *BatchUpdateFormDataByInstanceMapRequest {
	s.AsynchronousExecution = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapRequest) SetFormUuid(v string) *BatchUpdateFormDataByInstanceMapRequest {
	s.FormUuid = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapRequest) SetIgnoreEmpty(v bool) *BatchUpdateFormDataByInstanceMapRequest {
	s.IgnoreEmpty = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapRequest) SetNoExecuteExpression(v bool) *BatchUpdateFormDataByInstanceMapRequest {
	s.NoExecuteExpression = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapRequest) SetSystemToken(v string) *BatchUpdateFormDataByInstanceMapRequest {
	s.SystemToken = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapRequest) SetUpdateFormDataJsonMap(v map[string]interface{}) *BatchUpdateFormDataByInstanceMapRequest {
	s.UpdateFormDataJsonMap = v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapRequest) SetUseLatestFormSchemaVersion(v bool) *BatchUpdateFormDataByInstanceMapRequest {
	s.UseLatestFormSchemaVersion = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapRequest) SetUserId(v string) *BatchUpdateFormDataByInstanceMapRequest {
	s.UserId = &v
	return s
}

type BatchUpdateFormDataByInstanceMapResponseBody struct {
	// 更新成功的表单实例ID
	Result []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s BatchUpdateFormDataByInstanceMapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateFormDataByInstanceMapResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdateFormDataByInstanceMapResponseBody) SetResult(v []*string) *BatchUpdateFormDataByInstanceMapResponseBody {
	s.Result = v
	return s
}

type BatchUpdateFormDataByInstanceMapResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchUpdateFormDataByInstanceMapResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchUpdateFormDataByInstanceMapResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateFormDataByInstanceMapResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateFormDataByInstanceMapResponse) SetHeaders(v map[string]*string) *BatchUpdateFormDataByInstanceMapResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapResponse) SetBody(v *BatchUpdateFormDataByInstanceMapResponseBody) *BatchUpdateFormDataByInstanceMapResponse {
	s.Body = v
	return s
}

type BuyAuthorizationOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BuyAuthorizationOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s BuyAuthorizationOrderHeaders) GoString() string {
	return s.String()
}

func (s *BuyAuthorizationOrderHeaders) SetCommonHeaders(v map[string]*string) *BuyAuthorizationOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BuyAuthorizationOrderHeaders) SetXAcsDingtalkAccessToken(v string) *BuyAuthorizationOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BuyAuthorizationOrderRequest struct {
	// 访问秘钥
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// 账户号
	AccountNumber *string `json:"accountNumber,omitempty" xml:"accountNumber,omitempty"`
	// 开始时间
	BeginTimeGMT *int64 `json:"beginTimeGMT,omitempty" xml:"beginTimeGMT,omitempty"`
	// 调用者unionId
	CallerUnionId *string `json:"callerUnionId,omitempty" xml:"callerUnionId,omitempty"`
	// 收费类型
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// 商业类型
	CommerceType *string `json:"commerceType,omitempty" xml:"commerceType,omitempty"`
	// 商品类型
	CommodityType *string `json:"commodityType,omitempty" xml:"commodityType,omitempty"`
	// 结束时间
	EndTimeGMT *int64 `json:"endTimeGMT,omitempty" xml:"endTimeGMT,omitempty"`
	// 实例id
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// 实例名称
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// 阿里云产品码
	ProduceCode *string `json:"produceCode,omitempty" xml:"produceCode,omitempty"`
}

func (s BuyAuthorizationOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s BuyAuthorizationOrderRequest) GoString() string {
	return s.String()
}

func (s *BuyAuthorizationOrderRequest) SetAccessKey(v string) *BuyAuthorizationOrderRequest {
	s.AccessKey = &v
	return s
}

func (s *BuyAuthorizationOrderRequest) SetAccountNumber(v string) *BuyAuthorizationOrderRequest {
	s.AccountNumber = &v
	return s
}

func (s *BuyAuthorizationOrderRequest) SetBeginTimeGMT(v int64) *BuyAuthorizationOrderRequest {
	s.BeginTimeGMT = &v
	return s
}

func (s *BuyAuthorizationOrderRequest) SetCallerUnionId(v string) *BuyAuthorizationOrderRequest {
	s.CallerUnionId = &v
	return s
}

func (s *BuyAuthorizationOrderRequest) SetChargeType(v string) *BuyAuthorizationOrderRequest {
	s.ChargeType = &v
	return s
}

func (s *BuyAuthorizationOrderRequest) SetCommerceType(v string) *BuyAuthorizationOrderRequest {
	s.CommerceType = &v
	return s
}

func (s *BuyAuthorizationOrderRequest) SetCommodityType(v string) *BuyAuthorizationOrderRequest {
	s.CommodityType = &v
	return s
}

func (s *BuyAuthorizationOrderRequest) SetEndTimeGMT(v int64) *BuyAuthorizationOrderRequest {
	s.EndTimeGMT = &v
	return s
}

func (s *BuyAuthorizationOrderRequest) SetInstanceId(v string) *BuyAuthorizationOrderRequest {
	s.InstanceId = &v
	return s
}

func (s *BuyAuthorizationOrderRequest) SetInstanceName(v string) *BuyAuthorizationOrderRequest {
	s.InstanceName = &v
	return s
}

func (s *BuyAuthorizationOrderRequest) SetProduceCode(v string) *BuyAuthorizationOrderRequest {
	s.ProduceCode = &v
	return s
}

type BuyAuthorizationOrderResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s BuyAuthorizationOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BuyAuthorizationOrderResponseBody) GoString() string {
	return s.String()
}

func (s *BuyAuthorizationOrderResponseBody) SetResult(v bool) *BuyAuthorizationOrderResponseBody {
	s.Result = &v
	return s
}

type BuyAuthorizationOrderResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BuyAuthorizationOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BuyAuthorizationOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s BuyAuthorizationOrderResponse) GoString() string {
	return s.String()
}

func (s *BuyAuthorizationOrderResponse) SetHeaders(v map[string]*string) *BuyAuthorizationOrderResponse {
	s.Headers = v
	return s
}

func (s *BuyAuthorizationOrderResponse) SetBody(v *BuyAuthorizationOrderResponseBody) *BuyAuthorizationOrderResponse {
	s.Body = v
	return s
}

type BuyFreshOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BuyFreshOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s BuyFreshOrderHeaders) GoString() string {
	return s.String()
}

func (s *BuyFreshOrderHeaders) SetCommonHeaders(v map[string]*string) *BuyFreshOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BuyFreshOrderHeaders) SetXAcsDingtalkAccessToken(v string) *BuyFreshOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BuyFreshOrderRequest struct {
	// 访问秘钥
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// 账户号
	AccountNumber *string `json:"accountNumber,omitempty" xml:"accountNumber,omitempty"`
	// 开始时间
	BeginTimeGMT *int64 `json:"beginTimeGMT,omitempty" xml:"beginTimeGMT,omitempty"`
	// 调用者unionId
	CallerUnionId *string `json:"callerUnionId,omitempty" xml:"callerUnionId,omitempty"`
	// 收费类型
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// 商业类型
	CommerceType *string `json:"commerceType,omitempty" xml:"commerceType,omitempty"`
	// 商品类型
	CommodityType *string `json:"commodityType,omitempty" xml:"commodityType,omitempty"`
	// 结束时间
	EndTimeGMT *int64 `json:"endTimeGMT,omitempty" xml:"endTimeGMT,omitempty"`
	// 实例id
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// 实例名称
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// 阿里云产品码
	ProduceCode *string `json:"produceCode,omitempty" xml:"produceCode,omitempty"`
}

func (s BuyFreshOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s BuyFreshOrderRequest) GoString() string {
	return s.String()
}

func (s *BuyFreshOrderRequest) SetAccessKey(v string) *BuyFreshOrderRequest {
	s.AccessKey = &v
	return s
}

func (s *BuyFreshOrderRequest) SetAccountNumber(v string) *BuyFreshOrderRequest {
	s.AccountNumber = &v
	return s
}

func (s *BuyFreshOrderRequest) SetBeginTimeGMT(v int64) *BuyFreshOrderRequest {
	s.BeginTimeGMT = &v
	return s
}

func (s *BuyFreshOrderRequest) SetCallerUnionId(v string) *BuyFreshOrderRequest {
	s.CallerUnionId = &v
	return s
}

func (s *BuyFreshOrderRequest) SetChargeType(v string) *BuyFreshOrderRequest {
	s.ChargeType = &v
	return s
}

func (s *BuyFreshOrderRequest) SetCommerceType(v string) *BuyFreshOrderRequest {
	s.CommerceType = &v
	return s
}

func (s *BuyFreshOrderRequest) SetCommodityType(v string) *BuyFreshOrderRequest {
	s.CommodityType = &v
	return s
}

func (s *BuyFreshOrderRequest) SetEndTimeGMT(v int64) *BuyFreshOrderRequest {
	s.EndTimeGMT = &v
	return s
}

func (s *BuyFreshOrderRequest) SetInstanceId(v string) *BuyFreshOrderRequest {
	s.InstanceId = &v
	return s
}

func (s *BuyFreshOrderRequest) SetInstanceName(v string) *BuyFreshOrderRequest {
	s.InstanceName = &v
	return s
}

func (s *BuyFreshOrderRequest) SetProduceCode(v string) *BuyFreshOrderRequest {
	s.ProduceCode = &v
	return s
}

type BuyFreshOrderResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s BuyFreshOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BuyFreshOrderResponseBody) GoString() string {
	return s.String()
}

func (s *BuyFreshOrderResponseBody) SetResult(v bool) *BuyFreshOrderResponseBody {
	s.Result = &v
	return s
}

type BuyFreshOrderResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BuyFreshOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BuyFreshOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s BuyFreshOrderResponse) GoString() string {
	return s.String()
}

func (s *BuyFreshOrderResponse) SetHeaders(v map[string]*string) *BuyFreshOrderResponse {
	s.Headers = v
	return s
}

func (s *BuyFreshOrderResponse) SetBody(v *BuyFreshOrderResponseBody) *BuyFreshOrderResponse {
	s.Body = v
	return s
}

type CheckCloudAccountStatusHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CheckCloudAccountStatusHeaders) String() string {
	return tea.Prettify(s)
}

func (s CheckCloudAccountStatusHeaders) GoString() string {
	return s.String()
}

func (s *CheckCloudAccountStatusHeaders) SetCommonHeaders(v map[string]*string) *CheckCloudAccountStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CheckCloudAccountStatusHeaders) SetXAcsDingtalkAccessToken(v string) *CheckCloudAccountStatusHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CheckCloudAccountStatusRequest struct {
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
}

func (s CheckCloudAccountStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckCloudAccountStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckCloudAccountStatusRequest) SetAccessKey(v string) *CheckCloudAccountStatusRequest {
	s.AccessKey = &v
	return s
}

type CheckCloudAccountStatusResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CheckCloudAccountStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckCloudAccountStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckCloudAccountStatusResponseBody) SetResult(v bool) *CheckCloudAccountStatusResponseBody {
	s.Result = &v
	return s
}

type CheckCloudAccountStatusResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckCloudAccountStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckCloudAccountStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckCloudAccountStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckCloudAccountStatusResponse) SetHeaders(v map[string]*string) *CheckCloudAccountStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckCloudAccountStatusResponse) SetBody(v *CheckCloudAccountStatusResponseBody) *CheckCloudAccountStatusResponse {
	s.Body = v
	return s
}

type CreateOrUpdateFormDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateOrUpdateFormDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateFormDataHeaders) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateFormDataHeaders) SetCommonHeaders(v map[string]*string) *CreateOrUpdateFormDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateOrUpdateFormDataHeaders) SetXAcsDingtalkAccessToken(v string) *CreateOrUpdateFormDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateOrUpdateFormDataRequest struct {
	// 宜搭应用编码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 宜搭表单实例数据 json格式，如果存在满足检索条件的表单实例数据则用此值增量更新满足检索条件的的表单实例数据，否则用此值创建一条表单实例。表单实例数据的结构请参考 https://www.yuque.com/yida/support/agb8im#f26a51f429f9f19aa0b5b3ee847ac556_h3_31
	FormDataJson *string `json:"formDataJson,omitempty" xml:"formDataJson,omitempty"`
	// 表单编码
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 是否不触发校验规则、关联业务规则和第三方服务回调
	NoExecuteExpression *bool `json:"noExecuteExpression,omitempty" xml:"noExecuteExpression,omitempty"`
	// 用于检索表单实例数据的检索条件
	SearchCondition *string `json:"searchCondition,omitempty" xml:"searchCondition,omitempty"`
	// 宜搭应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateOrUpdateFormDataRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateFormDataRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateFormDataRequest) SetAppType(v string) *CreateOrUpdateFormDataRequest {
	s.AppType = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) SetFormDataJson(v string) *CreateOrUpdateFormDataRequest {
	s.FormDataJson = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) SetFormUuid(v string) *CreateOrUpdateFormDataRequest {
	s.FormUuid = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) SetNoExecuteExpression(v bool) *CreateOrUpdateFormDataRequest {
	s.NoExecuteExpression = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) SetSearchCondition(v string) *CreateOrUpdateFormDataRequest {
	s.SearchCondition = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) SetSystemToken(v string) *CreateOrUpdateFormDataRequest {
	s.SystemToken = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) SetUserId(v string) *CreateOrUpdateFormDataRequest {
	s.UserId = &v
	return s
}

type CreateOrUpdateFormDataResponseBody struct {
	// 新增的或者更新的表单实例id列表
	Result []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s CreateOrUpdateFormDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateFormDataResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateFormDataResponseBody) SetResult(v []*string) *CreateOrUpdateFormDataResponseBody {
	s.Result = v
	return s
}

type CreateOrUpdateFormDataResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateOrUpdateFormDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOrUpdateFormDataResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateFormDataResponse) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateFormDataResponse) SetHeaders(v map[string]*string) *CreateOrUpdateFormDataResponse {
	s.Headers = v
	return s
}

func (s *CreateOrUpdateFormDataResponse) SetBody(v *CreateOrUpdateFormDataResponseBody) *CreateOrUpdateFormDataResponse {
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
	// 要删除的表单数据ID
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// 语言
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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

func (s *DeleteFormDataRequest) SetFormInstanceId(v string) *DeleteFormDataRequest {
	s.FormInstanceId = &v
	return s
}

func (s *DeleteFormDataRequest) SetLanguage(v string) *DeleteFormDataRequest {
	s.Language = &v
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
	// 语言
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 流程实例ID
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// 应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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

func (s *DeleteInstanceRequest) SetLanguage(v string) *DeleteInstanceRequest {
	s.Language = &v
	return s
}

func (s *DeleteInstanceRequest) SetProcessInstanceId(v string) *DeleteInstanceRequest {
	s.ProcessInstanceId = &v
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
	AppType     *string `json:"appType,omitempty" xml:"appType,omitempty"`
	Language    *string `json:"language,omitempty" xml:"language,omitempty"`
	Sequence    *string `json:"sequence,omitempty" xml:"sequence,omitempty"`
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeleteSequenceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSequenceRequest) GoString() string {
	return s.String()
}

func (s *DeleteSequenceRequest) SetAppType(v string) *DeleteSequenceRequest {
	s.AppType = &v
	return s
}

func (s *DeleteSequenceRequest) SetLanguage(v string) *DeleteSequenceRequest {
	s.Language = &v
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

func (s *DeleteSequenceRequest) SetUserId(v string) *DeleteSequenceRequest {
	s.UserId = &v
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

type DeployFunctionCallbackHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeployFunctionCallbackHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeployFunctionCallbackHeaders) GoString() string {
	return s.String()
}

func (s *DeployFunctionCallbackHeaders) SetCommonHeaders(v map[string]*string) *DeployFunctionCallbackHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeployFunctionCallbackHeaders) SetXAcsDingtalkAccessToken(v string) *DeployFunctionCallbackHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeployFunctionCallbackRequest struct {
	// 云应用id
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// 自定义域名
	CustomDomain *string `json:"customDomain,omitempty" xml:"customDomain,omitempty"`
	// 部署阶段
	DeployStage *string `json:"deployStage,omitempty" xml:"deployStage,omitempty"`
	// api网关实例的AppKey
	GateWayAppKey *string `json:"gateWayAppKey,omitempty" xml:"gateWayAppKey,omitempty"`
	// api网关实例的APPSecret
	GateWayAppSecret *string `json:"gateWayAppSecret,omitempty" xml:"gateWayAppSecret,omitempty"`
	// api网关二级域名
	GateWayDomain *string `json:"gateWayDomain,omitempty" xml:"gateWayDomain,omitempty"`
}

func (s DeployFunctionCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployFunctionCallbackRequest) GoString() string {
	return s.String()
}

func (s *DeployFunctionCallbackRequest) SetAppId(v string) *DeployFunctionCallbackRequest {
	s.AppId = &v
	return s
}

func (s *DeployFunctionCallbackRequest) SetCustomDomain(v string) *DeployFunctionCallbackRequest {
	s.CustomDomain = &v
	return s
}

func (s *DeployFunctionCallbackRequest) SetDeployStage(v string) *DeployFunctionCallbackRequest {
	s.DeployStage = &v
	return s
}

func (s *DeployFunctionCallbackRequest) SetGateWayAppKey(v string) *DeployFunctionCallbackRequest {
	s.GateWayAppKey = &v
	return s
}

func (s *DeployFunctionCallbackRequest) SetGateWayAppSecret(v string) *DeployFunctionCallbackRequest {
	s.GateWayAppSecret = &v
	return s
}

func (s *DeployFunctionCallbackRequest) SetGateWayDomain(v string) *DeployFunctionCallbackRequest {
	s.GateWayDomain = &v
	return s
}

type DeployFunctionCallbackResponseBody struct {
	// 是否处理成功
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeployFunctionCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployFunctionCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *DeployFunctionCallbackResponseBody) SetResult(v bool) *DeployFunctionCallbackResponseBody {
	s.Result = &v
	return s
}

type DeployFunctionCallbackResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeployFunctionCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeployFunctionCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployFunctionCallbackResponse) GoString() string {
	return s.String()
}

func (s *DeployFunctionCallbackResponse) SetHeaders(v map[string]*string) *DeployFunctionCallbackResponse {
	s.Headers = v
	return s
}

func (s *DeployFunctionCallbackResponse) SetBody(v *DeployFunctionCallbackResponseBody) *DeployFunctionCallbackResponse {
	s.Body = v
	return s
}

type ExecuteBatchTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ExecuteBatchTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s ExecuteBatchTaskHeaders) GoString() string {
	return s.String()
}

func (s *ExecuteBatchTaskHeaders) SetCommonHeaders(v map[string]*string) *ExecuteBatchTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ExecuteBatchTaskHeaders) SetXAcsDingtalkAccessToken(v string) *ExecuteBatchTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ExecuteBatchTaskRequest struct {
	// 宜搭应用编码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 审批动作
	OutResult *string `json:"outResult,omitempty" xml:"outResult,omitempty"`
	// 审批意见
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 宜搭应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// taskInfoList
	TaskInformationList *string `json:"taskInformationList,omitempty" xml:"taskInformationList,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteBatchTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteBatchTaskRequest) GoString() string {
	return s.String()
}

func (s *ExecuteBatchTaskRequest) SetAppType(v string) *ExecuteBatchTaskRequest {
	s.AppType = &v
	return s
}

func (s *ExecuteBatchTaskRequest) SetOutResult(v string) *ExecuteBatchTaskRequest {
	s.OutResult = &v
	return s
}

func (s *ExecuteBatchTaskRequest) SetRemark(v string) *ExecuteBatchTaskRequest {
	s.Remark = &v
	return s
}

func (s *ExecuteBatchTaskRequest) SetSystemToken(v string) *ExecuteBatchTaskRequest {
	s.SystemToken = &v
	return s
}

func (s *ExecuteBatchTaskRequest) SetTaskInformationList(v string) *ExecuteBatchTaskRequest {
	s.TaskInformationList = &v
	return s
}

func (s *ExecuteBatchTaskRequest) SetUserId(v string) *ExecuteBatchTaskRequest {
	s.UserId = &v
	return s
}

type ExecuteBatchTaskResponseBody struct {
	// 审批失败的任务数
	FailNumber *int32 `json:"failNumber,omitempty" xml:"failNumber,omitempty"`
	// 审批成功的任务数
	SuccessNumber *int32 `json:"successNumber,omitempty" xml:"successNumber,omitempty"`
	// 总任务数
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ExecuteBatchTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteBatchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteBatchTaskResponseBody) SetFailNumber(v int32) *ExecuteBatchTaskResponseBody {
	s.FailNumber = &v
	return s
}

func (s *ExecuteBatchTaskResponseBody) SetSuccessNumber(v int32) *ExecuteBatchTaskResponseBody {
	s.SuccessNumber = &v
	return s
}

func (s *ExecuteBatchTaskResponseBody) SetTotal(v int32) *ExecuteBatchTaskResponseBody {
	s.Total = &v
	return s
}

type ExecuteBatchTaskResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExecuteBatchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteBatchTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteBatchTaskResponse) GoString() string {
	return s.String()
}

func (s *ExecuteBatchTaskResponse) SetHeaders(v map[string]*string) *ExecuteBatchTaskResponse {
	s.Headers = v
	return s
}

func (s *ExecuteBatchTaskResponse) SetBody(v *ExecuteBatchTaskResponseBody) *ExecuteBatchTaskResponse {
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
	AppType     *string `json:"appType,omitempty" xml:"appType,omitempty"`
	Data        *string `json:"data,omitempty" xml:"data,omitempty"`
	Language    *string `json:"language,omitempty" xml:"language,omitempty"`
	ServiceId   *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteCustomApiRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteCustomApiRequest) GoString() string {
	return s.String()
}

func (s *ExecuteCustomApiRequest) SetAppType(v string) *ExecuteCustomApiRequest {
	s.AppType = &v
	return s
}

func (s *ExecuteCustomApiRequest) SetData(v string) *ExecuteCustomApiRequest {
	s.Data = &v
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

func (s *ExecuteCustomApiRequest) SetSystemToken(v string) *ExecuteCustomApiRequest {
	s.SystemToken = &v
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
	// 应用ID
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 更新的表单数据
	FormDataJson *string `json:"formDataJson,omitempty" xml:"formDataJson,omitempty"`
	// 语言
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 是否不执行校验&关联操作
	NoExecuteExpressions *string `json:"noExecuteExpressions,omitempty" xml:"noExecuteExpressions,omitempty"`
	// 审批结果
	OutResult *string `json:"outResult,omitempty" xml:"outResult,omitempty"`
	// 流程实例ID
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// 审批意见
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecutePlatformTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecutePlatformTaskRequest) GoString() string {
	return s.String()
}

func (s *ExecutePlatformTaskRequest) SetAppType(v string) *ExecutePlatformTaskRequest {
	s.AppType = &v
	return s
}

func (s *ExecutePlatformTaskRequest) SetFormDataJson(v string) *ExecutePlatformTaskRequest {
	s.FormDataJson = &v
	return s
}

func (s *ExecutePlatformTaskRequest) SetLanguage(v string) *ExecutePlatformTaskRequest {
	s.Language = &v
	return s
}

func (s *ExecutePlatformTaskRequest) SetNoExecuteExpressions(v string) *ExecutePlatformTaskRequest {
	s.NoExecuteExpressions = &v
	return s
}

func (s *ExecutePlatformTaskRequest) SetOutResult(v string) *ExecutePlatformTaskRequest {
	s.OutResult = &v
	return s
}

func (s *ExecutePlatformTaskRequest) SetProcessInstanceId(v string) *ExecutePlatformTaskRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *ExecutePlatformTaskRequest) SetRemark(v string) *ExecutePlatformTaskRequest {
	s.Remark = &v
	return s
}

func (s *ExecutePlatformTaskRequest) SetSystemToken(v string) *ExecutePlatformTaskRequest {
	s.SystemToken = &v
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
	// 应用ID
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 电子签名
	DigitalSignUrl *string `json:"digitalSignUrl,omitempty" xml:"digitalSignUrl,omitempty"`
	// 更新的表单值
	FormDataJson *string `json:"formDataJson,omitempty" xml:"formDataJson,omitempty"`
	// 语言
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 是否不执行校验&关联操作
	NoExecuteExpressions *string `json:"noExecuteExpressions,omitempty" xml:"noExecuteExpressions,omitempty"`
	// 审批结果
	OutResult *string `json:"outResult,omitempty" xml:"outResult,omitempty"`
	// 实例ID
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// 审批意见
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 任务ID
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// 钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteTaskRequest) GoString() string {
	return s.String()
}

func (s *ExecuteTaskRequest) SetAppType(v string) *ExecuteTaskRequest {
	s.AppType = &v
	return s
}

func (s *ExecuteTaskRequest) SetDigitalSignUrl(v string) *ExecuteTaskRequest {
	s.DigitalSignUrl = &v
	return s
}

func (s *ExecuteTaskRequest) SetFormDataJson(v string) *ExecuteTaskRequest {
	s.FormDataJson = &v
	return s
}

func (s *ExecuteTaskRequest) SetLanguage(v string) *ExecuteTaskRequest {
	s.Language = &v
	return s
}

func (s *ExecuteTaskRequest) SetNoExecuteExpressions(v string) *ExecuteTaskRequest {
	s.NoExecuteExpressions = &v
	return s
}

func (s *ExecuteTaskRequest) SetOutResult(v string) *ExecuteTaskRequest {
	s.OutResult = &v
	return s
}

func (s *ExecuteTaskRequest) SetProcessInstanceId(v string) *ExecuteTaskRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *ExecuteTaskRequest) SetRemark(v string) *ExecuteTaskRequest {
	s.Remark = &v
	return s
}

func (s *ExecuteTaskRequest) SetSystemToken(v string) *ExecuteTaskRequest {
	s.SystemToken = &v
	return s
}

func (s *ExecuteTaskRequest) SetTaskId(v int64) *ExecuteTaskRequest {
	s.TaskId = &v
	return s
}

func (s *ExecuteTaskRequest) SetUserId(v string) *ExecuteTaskRequest {
	s.UserId = &v
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
	AccessKey  *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	CallerUid  *string `json:"callerUid,omitempty" xml:"callerUid,omitempty"`
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s ExpireCommodityRequest) String() string {
	return tea.Prettify(s)
}

func (s ExpireCommodityRequest) GoString() string {
	return s.String()
}

func (s *ExpireCommodityRequest) SetAccessKey(v string) *ExpireCommodityRequest {
	s.AccessKey = &v
	return s
}

func (s *ExpireCommodityRequest) SetCallerUid(v string) *ExpireCommodityRequest {
	s.CallerUid = &v
	return s
}

func (s *ExpireCommodityRequest) SetInstanceId(v string) *ExpireCommodityRequest {
	s.InstanceId = &v
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

type GetActivationCodeByCallerUnionIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetActivationCodeByCallerUnionIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetActivationCodeByCallerUnionIdHeaders) GoString() string {
	return s.String()
}

func (s *GetActivationCodeByCallerUnionIdHeaders) SetCommonHeaders(v map[string]*string) *GetActivationCodeByCallerUnionIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetActivationCodeByCallerUnionIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetActivationCodeByCallerUnionIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetActivationCodeByCallerUnionIdRequest struct {
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
}

func (s GetActivationCodeByCallerUnionIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetActivationCodeByCallerUnionIdRequest) GoString() string {
	return s.String()
}

func (s *GetActivationCodeByCallerUnionIdRequest) SetAccessKey(v string) *GetActivationCodeByCallerUnionIdRequest {
	s.AccessKey = &v
	return s
}

type GetActivationCodeByCallerUnionIdResponseBody struct {
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetActivationCodeByCallerUnionIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetActivationCodeByCallerUnionIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetActivationCodeByCallerUnionIdResponseBody) SetResult(v string) *GetActivationCodeByCallerUnionIdResponseBody {
	s.Result = &v
	return s
}

type GetActivationCodeByCallerUnionIdResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetActivationCodeByCallerUnionIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetActivationCodeByCallerUnionIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetActivationCodeByCallerUnionIdResponse) GoString() string {
	return s.String()
}

func (s *GetActivationCodeByCallerUnionIdResponse) SetHeaders(v map[string]*string) *GetActivationCodeByCallerUnionIdResponse {
	s.Headers = v
	return s
}

func (s *GetActivationCodeByCallerUnionIdResponse) SetBody(v *GetActivationCodeByCallerUnionIdResponseBody) *GetActivationCodeByCallerUnionIdResponse {
	s.Body = v
	return s
}

type GetActivityButtonListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetActivityButtonListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetActivityButtonListHeaders) GoString() string {
	return s.String()
}

func (s *GetActivityButtonListHeaders) SetCommonHeaders(v map[string]*string) *GetActivityButtonListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetActivityButtonListHeaders) SetXAcsDingtalkAccessToken(v string) *GetActivityButtonListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetActivityButtonListRequest struct {
	// 语言环境
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetActivityButtonListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetActivityButtonListRequest) GoString() string {
	return s.String()
}

func (s *GetActivityButtonListRequest) SetLanguage(v string) *GetActivityButtonListRequest {
	s.Language = &v
	return s
}

func (s *GetActivityButtonListRequest) SetSystemToken(v string) *GetActivityButtonListRequest {
	s.SystemToken = &v
	return s
}

func (s *GetActivityButtonListRequest) SetUserId(v string) *GetActivityButtonListRequest {
	s.UserId = &v
	return s
}

type GetActivityButtonListResponseBody struct {
	Result []*GetActivityButtonListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetActivityButtonListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetActivityButtonListResponseBody) GoString() string {
	return s.String()
}

func (s *GetActivityButtonListResponseBody) SetResult(v []*GetActivityButtonListResponseBodyResult) *GetActivityButtonListResponseBody {
	s.Result = v
	return s
}

type GetActivityButtonListResponseBodyResult struct {
	// alias
	AliasInChinese *string `json:"aliasInChinese,omitempty" xml:"aliasInChinese,omitempty"`
	// aliasEn
	AliasInEnglish *string `json:"aliasInEnglish,omitempty" xml:"aliasInEnglish,omitempty"`
}

func (s GetActivityButtonListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetActivityButtonListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetActivityButtonListResponseBodyResult) SetAliasInChinese(v string) *GetActivityButtonListResponseBodyResult {
	s.AliasInChinese = &v
	return s
}

func (s *GetActivityButtonListResponseBodyResult) SetAliasInEnglish(v string) *GetActivityButtonListResponseBodyResult {
	s.AliasInEnglish = &v
	return s
}

type GetActivityButtonListResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetActivityButtonListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetActivityButtonListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetActivityButtonListResponse) GoString() string {
	return s.String()
}

func (s *GetActivityButtonListResponse) SetHeaders(v map[string]*string) *GetActivityButtonListResponse {
	s.Headers = v
	return s
}

func (s *GetActivityButtonListResponse) SetBody(v *GetActivityButtonListResponseBody) *GetActivityButtonListResponse {
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
	AppType     *string `json:"appType,omitempty" xml:"appType,omitempty"`
	Language    *string `json:"language,omitempty" xml:"language,omitempty"`
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetActivityListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetActivityListRequest) GoString() string {
	return s.String()
}

func (s *GetActivityListRequest) SetAppType(v string) *GetActivityListRequest {
	s.AppType = &v
	return s
}

func (s *GetActivityListRequest) SetLanguage(v string) *GetActivityListRequest {
	s.Language = &v
	return s
}

func (s *GetActivityListRequest) SetProcessCode(v string) *GetActivityListRequest {
	s.ProcessCode = &v
	return s
}

func (s *GetActivityListRequest) SetSystemToken(v string) *GetActivityListRequest {
	s.SystemToken = &v
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
	// activityId
	ActivityId *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
	// activityName
	ActivityName *string `json:"activityName,omitempty" xml:"activityName,omitempty"`
	// activityNameEn
	ActivityNameInEnglish *string `json:"activityNameInEnglish,omitempty" xml:"activityNameInEnglish,omitempty"`
}

func (s GetActivityListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetActivityListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetActivityListResponseBodyResult) SetActivityId(v string) *GetActivityListResponseBodyResult {
	s.ActivityId = &v
	return s
}

func (s *GetActivityListResponseBodyResult) SetActivityName(v string) *GetActivityListResponseBodyResult {
	s.ActivityName = &v
	return s
}

func (s *GetActivityListResponseBodyResult) SetActivityNameInEnglish(v string) *GetActivityListResponseBodyResult {
	s.ActivityNameInEnglish = &v
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
	AccessKey  *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	CallerUid  *string `json:"callerUid,omitempty" xml:"callerUid,omitempty"`
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s GetApplicationAuthorizationServicePlatformResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationAuthorizationServicePlatformResourceRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationAuthorizationServicePlatformResourceRequest) SetAccessKey(v string) *GetApplicationAuthorizationServicePlatformResourceRequest {
	s.AccessKey = &v
	return s
}

func (s *GetApplicationAuthorizationServicePlatformResourceRequest) SetCallerUid(v string) *GetApplicationAuthorizationServicePlatformResourceRequest {
	s.CallerUid = &v
	return s
}

func (s *GetApplicationAuthorizationServicePlatformResourceRequest) SetInstanceId(v string) *GetApplicationAuthorizationServicePlatformResourceRequest {
	s.InstanceId = &v
	return s
}

type GetApplicationAuthorizationServicePlatformResourceResponseBody struct {
	// accountTotalAmount
	AccountTotalAmount *int32 `json:"accountTotalAmount,omitempty" xml:"accountTotalAmount,omitempty"`
	// accountUsageAmount
	AccountUsageAmount *int32 `json:"accountUsageAmount,omitempty" xml:"accountUsageAmount,omitempty"`
	// appTotalAmount
	AppTotalAmount *int32 `json:"appTotalAmount,omitempty" xml:"appTotalAmount,omitempty"`
	// attachmentTotalAmount
	AttachmentTotalAmount *int64 `json:"attachmentTotalAmount,omitempty" xml:"attachmentTotalAmount,omitempty"`
	// attachmentUsageAmount
	AttachmentUsageAmount *int64 `json:"attachmentUsageAmount,omitempty" xml:"attachmentUsageAmount,omitempty"`
	// instanceId
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// instanceTotalAmount
	InstanceTotalAmount *int64 `json:"instanceTotalAmount,omitempty" xml:"instanceTotalAmount,omitempty"`
	// instanceUsageAmount
	InstanceUsageAmount *int64 `json:"instanceUsageAmount,omitempty" xml:"instanceUsageAmount,omitempty"`
	// pluginUsageAmount
	PluginUsageAmount *int64 `json:"pluginUsageAmount,omitempty" xml:"pluginUsageAmount,omitempty"`
}

func (s GetApplicationAuthorizationServicePlatformResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationAuthorizationServicePlatformResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationAuthorizationServicePlatformResourceResponseBody) SetAccountTotalAmount(v int32) *GetApplicationAuthorizationServicePlatformResourceResponseBody {
	s.AccountTotalAmount = &v
	return s
}

func (s *GetApplicationAuthorizationServicePlatformResourceResponseBody) SetAccountUsageAmount(v int32) *GetApplicationAuthorizationServicePlatformResourceResponseBody {
	s.AccountUsageAmount = &v
	return s
}

func (s *GetApplicationAuthorizationServicePlatformResourceResponseBody) SetAppTotalAmount(v int32) *GetApplicationAuthorizationServicePlatformResourceResponseBody {
	s.AppTotalAmount = &v
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

func (s *GetApplicationAuthorizationServicePlatformResourceResponseBody) SetPluginUsageAmount(v int64) *GetApplicationAuthorizationServicePlatformResourceResponseBody {
	s.PluginUsageAmount = &v
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

type GetCorpAccomplishmentTasksHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCorpAccomplishmentTasksHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCorpAccomplishmentTasksHeaders) GoString() string {
	return s.String()
}

func (s *GetCorpAccomplishmentTasksHeaders) SetCommonHeaders(v map[string]*string) *GetCorpAccomplishmentTasksHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCorpAccomplishmentTasksHeaders) SetXAcsDingtalkAccessToken(v string) *GetCorpAccomplishmentTasksHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCorpAccomplishmentTasksRequest struct {
	// 应用标识列表
	AppTypes *string `json:"appTypes,omitempty" xml:"appTypes,omitempty"`
	// 创建时间开始
	CreateFromTimeGMT *int64 `json:"createFromTimeGMT,omitempty" xml:"createFromTimeGMT,omitempty"`
	// 创建时间结束
	CreateToTimeGMT *int64 `json:"createToTimeGMT,omitempty" xml:"createToTimeGMT,omitempty"`
	// 关键词
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 语言环境
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 当前页
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页记录数
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 流程code列表
	ProcessCodes *string `json:"processCodes,omitempty" xml:"processCodes,omitempty"`
	// 验权token
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s GetCorpAccomplishmentTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCorpAccomplishmentTasksRequest) GoString() string {
	return s.String()
}

func (s *GetCorpAccomplishmentTasksRequest) SetAppTypes(v string) *GetCorpAccomplishmentTasksRequest {
	s.AppTypes = &v
	return s
}

func (s *GetCorpAccomplishmentTasksRequest) SetCreateFromTimeGMT(v int64) *GetCorpAccomplishmentTasksRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *GetCorpAccomplishmentTasksRequest) SetCreateToTimeGMT(v int64) *GetCorpAccomplishmentTasksRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *GetCorpAccomplishmentTasksRequest) SetKeyword(v string) *GetCorpAccomplishmentTasksRequest {
	s.Keyword = &v
	return s
}

func (s *GetCorpAccomplishmentTasksRequest) SetLanguage(v string) *GetCorpAccomplishmentTasksRequest {
	s.Language = &v
	return s
}

func (s *GetCorpAccomplishmentTasksRequest) SetPageNumber(v int32) *GetCorpAccomplishmentTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *GetCorpAccomplishmentTasksRequest) SetPageSize(v int32) *GetCorpAccomplishmentTasksRequest {
	s.PageSize = &v
	return s
}

func (s *GetCorpAccomplishmentTasksRequest) SetProcessCodes(v string) *GetCorpAccomplishmentTasksRequest {
	s.ProcessCodes = &v
	return s
}

func (s *GetCorpAccomplishmentTasksRequest) SetToken(v string) *GetCorpAccomplishmentTasksRequest {
	s.Token = &v
	return s
}

type GetCorpAccomplishmentTasksResponseBody struct {
	// data
	Data []*GetCorpAccomplishmentTasksResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 当前第几页
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 总数量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetCorpAccomplishmentTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCorpAccomplishmentTasksResponseBody) GoString() string {
	return s.String()
}

func (s *GetCorpAccomplishmentTasksResponseBody) SetData(v []*GetCorpAccomplishmentTasksResponseBodyData) *GetCorpAccomplishmentTasksResponseBody {
	s.Data = v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBody) SetPageNumber(v int64) *GetCorpAccomplishmentTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBody) SetTotalCount(v int64) *GetCorpAccomplishmentTasksResponseBody {
	s.TotalCount = &v
	return s
}

type GetCorpAccomplishmentTasksResponseBodyData struct {
	// activeTime
	ActiveTimeGMT *string `json:"activeTimeGMT,omitempty" xml:"activeTimeGMT,omitempty"`
	// actualActionerId
	ActualActionerId *string `json:"actualActionerId,omitempty" xml:"actualActionerId,omitempty"`
	// appType
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// createTime
	CreateTimeGMT *string `json:"createTimeGMT,omitempty" xml:"createTimeGMT,omitempty"`
	// finishTime
	FinishTimeGMT *string `json:"finishTimeGMT,omitempty" xml:"finishTimeGMT,omitempty"`
	// originatorEmail
	OriginatorEmail *string `json:"originatorEmail,omitempty" xml:"originatorEmail,omitempty"`
	// originatorId
	OriginatorId *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	// originatorName
	OriginatorName *string `json:"originatorName,omitempty" xml:"originatorName,omitempty"`
	// originatorNameEn
	OriginatorNameInEnglish *string `json:"originatorNameInEnglish,omitempty" xml:"originatorNameInEnglish,omitempty"`
	// originatorNickName
	OriginatorNickName *string `json:"originatorNickName,omitempty" xml:"originatorNickName,omitempty"`
	// originatorNickNameEn
	OriginatorNickNameInEnglish *string `json:"originatorNickNameInEnglish,omitempty" xml:"originatorNickNameInEnglish,omitempty"`
	// originatorPhoto
	OriginatorPhoto *string `json:"originatorPhoto,omitempty" xml:"originatorPhoto,omitempty"`
	// outResult
	OutResult *string `json:"outResult,omitempty" xml:"outResult,omitempty"`
	// outResultName
	OutResultName *string `json:"outResultName,omitempty" xml:"outResultName,omitempty"`
	// processInstanceId
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// status
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// taskId
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// taskType
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// title
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// titleEn
	TitleInEnglish *string `json:"titleInEnglish,omitempty" xml:"titleInEnglish,omitempty"`
}

func (s GetCorpAccomplishmentTasksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCorpAccomplishmentTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetActiveTimeGMT(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.ActiveTimeGMT = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetActualActionerId(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.ActualActionerId = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetAppType(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.AppType = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetCreateTimeGMT(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.CreateTimeGMT = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetFinishTimeGMT(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.FinishTimeGMT = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetOriginatorEmail(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.OriginatorEmail = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetOriginatorId(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.OriginatorId = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetOriginatorName(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.OriginatorName = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetOriginatorNameInEnglish(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.OriginatorNameInEnglish = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetOriginatorNickName(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.OriginatorNickName = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetOriginatorNickNameInEnglish(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.OriginatorNickNameInEnglish = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetOriginatorPhoto(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.OriginatorPhoto = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetOutResult(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.OutResult = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetOutResultName(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.OutResultName = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetProcessInstanceId(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetStatus(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetTaskId(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetTaskType(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetTitle(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetCorpAccomplishmentTasksResponseBodyData) SetTitleInEnglish(v string) *GetCorpAccomplishmentTasksResponseBodyData {
	s.TitleInEnglish = &v
	return s
}

type GetCorpAccomplishmentTasksResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCorpAccomplishmentTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCorpAccomplishmentTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCorpAccomplishmentTasksResponse) GoString() string {
	return s.String()
}

func (s *GetCorpAccomplishmentTasksResponse) SetHeaders(v map[string]*string) *GetCorpAccomplishmentTasksResponse {
	s.Headers = v
	return s
}

func (s *GetCorpAccomplishmentTasksResponse) SetBody(v *GetCorpAccomplishmentTasksResponseBody) *GetCorpAccomplishmentTasksResponse {
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

type GetCorpTasksHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCorpTasksHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCorpTasksHeaders) GoString() string {
	return s.String()
}

func (s *GetCorpTasksHeaders) SetCommonHeaders(v map[string]*string) *GetCorpTasksHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCorpTasksHeaders) SetXAcsDingtalkAccessToken(v string) *GetCorpTasksHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCorpTasksRequest struct {
	// 应用标识列表
	AppTypes *string `json:"appTypes,omitempty" xml:"appTypes,omitempty"`
	// 企业ID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 创建时间开始
	CreateFromTimeGMT *int64 `json:"createFromTimeGMT,omitempty" xml:"createFromTimeGMT,omitempty"`
	// 创建时间结束
	CreateToTimeGMT *int64 `json:"createToTimeGMT,omitempty" xml:"createToTimeGMT,omitempty"`
	// 关键词
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 语言环境
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 当前页
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页记录数
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 流程code列表
	ProcessCodes *string `json:"processCodes,omitempty" xml:"processCodes,omitempty"`
	// 验权token
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// 钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetCorpTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCorpTasksRequest) GoString() string {
	return s.String()
}

func (s *GetCorpTasksRequest) SetAppTypes(v string) *GetCorpTasksRequest {
	s.AppTypes = &v
	return s
}

func (s *GetCorpTasksRequest) SetCorpId(v string) *GetCorpTasksRequest {
	s.CorpId = &v
	return s
}

func (s *GetCorpTasksRequest) SetCreateFromTimeGMT(v int64) *GetCorpTasksRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *GetCorpTasksRequest) SetCreateToTimeGMT(v int64) *GetCorpTasksRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *GetCorpTasksRequest) SetKeyword(v string) *GetCorpTasksRequest {
	s.Keyword = &v
	return s
}

func (s *GetCorpTasksRequest) SetLanguage(v string) *GetCorpTasksRequest {
	s.Language = &v
	return s
}

func (s *GetCorpTasksRequest) SetPageNumber(v int32) *GetCorpTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *GetCorpTasksRequest) SetPageSize(v int32) *GetCorpTasksRequest {
	s.PageSize = &v
	return s
}

func (s *GetCorpTasksRequest) SetProcessCodes(v string) *GetCorpTasksRequest {
	s.ProcessCodes = &v
	return s
}

func (s *GetCorpTasksRequest) SetToken(v string) *GetCorpTasksRequest {
	s.Token = &v
	return s
}

func (s *GetCorpTasksRequest) SetUserId(v string) *GetCorpTasksRequest {
	s.UserId = &v
	return s
}

type GetCorpTasksResponseBody struct {
	// data
	Data []*GetCorpTasksResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 当前第几页
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 总数量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetCorpTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCorpTasksResponseBody) GoString() string {
	return s.String()
}

func (s *GetCorpTasksResponseBody) SetData(v []*GetCorpTasksResponseBodyData) *GetCorpTasksResponseBody {
	s.Data = v
	return s
}

func (s *GetCorpTasksResponseBody) SetPageNumber(v int64) *GetCorpTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetCorpTasksResponseBody) SetTotalCount(v int64) *GetCorpTasksResponseBody {
	s.TotalCount = &v
	return s
}

type GetCorpTasksResponseBodyData struct {
	// activeTime
	ActiveTimeGMT *string `json:"activeTimeGMT,omitempty" xml:"activeTimeGMT,omitempty"`
	// actualActionerId
	ActualActionerId *string `json:"actualActionerId,omitempty" xml:"actualActionerId,omitempty"`
	// appType
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// createTime
	CreateTimeGMT *string `json:"createTimeGMT,omitempty" xml:"createTimeGMT,omitempty"`
	// finishTime
	FinishTimeGMT *string `json:"finishTimeGMT,omitempty" xml:"finishTimeGMT,omitempty"`
	// originatorEmail
	OriginatorEmail *string `json:"originatorEmail,omitempty" xml:"originatorEmail,omitempty"`
	// originatorId
	OriginatorId *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	// originatorName
	OriginatorName *string `json:"originatorName,omitempty" xml:"originatorName,omitempty"`
	// originatorNameEn
	OriginatorNameInEnglish *string `json:"originatorNameInEnglish,omitempty" xml:"originatorNameInEnglish,omitempty"`
	// originatorNickName
	OriginatorNickName *string `json:"originatorNickName,omitempty" xml:"originatorNickName,omitempty"`
	// originatorNickNameEn
	OriginatorNickNameEn *string `json:"originatorNickNameEn,omitempty" xml:"originatorNickNameEn,omitempty"`
	// originatorPhoto
	OriginatorPhoto *string `json:"originatorPhoto,omitempty" xml:"originatorPhoto,omitempty"`
	// outResult
	OutResult *string `json:"outResult,omitempty" xml:"outResult,omitempty"`
	// outResultName
	OutResultName *string `json:"outResultName,omitempty" xml:"outResultName,omitempty"`
	// processInstanceId
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// status
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// taskId
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// taskType
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// title
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// titleEn
	TitleInEnglish *string `json:"titleInEnglish,omitempty" xml:"titleInEnglish,omitempty"`
}

func (s GetCorpTasksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCorpTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCorpTasksResponseBodyData) SetActiveTimeGMT(v string) *GetCorpTasksResponseBodyData {
	s.ActiveTimeGMT = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetActualActionerId(v string) *GetCorpTasksResponseBodyData {
	s.ActualActionerId = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetAppType(v string) *GetCorpTasksResponseBodyData {
	s.AppType = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetCreateTimeGMT(v string) *GetCorpTasksResponseBodyData {
	s.CreateTimeGMT = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetFinishTimeGMT(v string) *GetCorpTasksResponseBodyData {
	s.FinishTimeGMT = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetOriginatorEmail(v string) *GetCorpTasksResponseBodyData {
	s.OriginatorEmail = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetOriginatorId(v string) *GetCorpTasksResponseBodyData {
	s.OriginatorId = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetOriginatorName(v string) *GetCorpTasksResponseBodyData {
	s.OriginatorName = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetOriginatorNameInEnglish(v string) *GetCorpTasksResponseBodyData {
	s.OriginatorNameInEnglish = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetOriginatorNickName(v string) *GetCorpTasksResponseBodyData {
	s.OriginatorNickName = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetOriginatorNickNameEn(v string) *GetCorpTasksResponseBodyData {
	s.OriginatorNickNameEn = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetOriginatorPhoto(v string) *GetCorpTasksResponseBodyData {
	s.OriginatorPhoto = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetOutResult(v string) *GetCorpTasksResponseBodyData {
	s.OutResult = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetOutResultName(v string) *GetCorpTasksResponseBodyData {
	s.OutResultName = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetProcessInstanceId(v string) *GetCorpTasksResponseBodyData {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetStatus(v string) *GetCorpTasksResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetTaskId(v string) *GetCorpTasksResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetTaskType(v string) *GetCorpTasksResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetTitle(v string) *GetCorpTasksResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetCorpTasksResponseBodyData) SetTitleInEnglish(v string) *GetCorpTasksResponseBodyData {
	s.TitleInEnglish = &v
	return s
}

type GetCorpTasksResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCorpTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCorpTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCorpTasksResponse) GoString() string {
	return s.String()
}

func (s *GetCorpTasksResponse) SetHeaders(v map[string]*string) *GetCorpTasksResponse {
	s.Headers = v
	return s
}

func (s *GetCorpTasksResponse) SetBody(v *GetCorpTasksResponseBody) *GetCorpTasksResponse {
	s.Body = v
	return s
}

type GetFieldDefByUuidHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFieldDefByUuidHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFieldDefByUuidHeaders) GoString() string {
	return s.String()
}

func (s *GetFieldDefByUuidHeaders) SetCommonHeaders(v map[string]*string) *GetFieldDefByUuidHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFieldDefByUuidHeaders) SetXAcsDingtalkAccessToken(v string) *GetFieldDefByUuidHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFieldDefByUuidRequest struct {
	// 应用编码。应用唯一标识。如：APP_XXX
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 表单唯一标识
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 应用秘钥。在应用设置-部署运维-应用密钥中获取。
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 操作人userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetFieldDefByUuidRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFieldDefByUuidRequest) GoString() string {
	return s.String()
}

func (s *GetFieldDefByUuidRequest) SetAppType(v string) *GetFieldDefByUuidRequest {
	s.AppType = &v
	return s
}

func (s *GetFieldDefByUuidRequest) SetFormUuid(v string) *GetFieldDefByUuidRequest {
	s.FormUuid = &v
	return s
}

func (s *GetFieldDefByUuidRequest) SetSystemToken(v string) *GetFieldDefByUuidRequest {
	s.SystemToken = &v
	return s
}

func (s *GetFieldDefByUuidRequest) SetUserId(v string) *GetFieldDefByUuidRequest {
	s.UserId = &v
	return s
}

type GetFieldDefByUuidResponseBody struct {
	Result []*GetFieldDefByUuidResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// 是否成功，true代表成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetFieldDefByUuidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFieldDefByUuidResponseBody) GoString() string {
	return s.String()
}

func (s *GetFieldDefByUuidResponseBody) SetResult(v []*GetFieldDefByUuidResponseBodyResult) *GetFieldDefByUuidResponseBody {
	s.Result = v
	return s
}

func (s *GetFieldDefByUuidResponseBody) SetSuccess(v bool) *GetFieldDefByUuidResponseBody {
	s.Success = &v
	return s
}

type GetFieldDefByUuidResponseBodyResult struct {
	// 组件展示状态
	// 普通NORMAL/禁用DISABLED/只读READONLY/隐藏HIDDEN
	Behavior *string `json:"behavior,omitempty" xml:"behavior,omitempty"`
	// 子组件信息
	Children *string `json:"children,omitempty" xml:"children,omitempty"`
	// 组件类型，如文本类型：TextField
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// 字段ID，字段唯一标识
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// 字段名称。符合国际化标准。
	Label interface{} `json:"label,omitempty" xml:"label,omitempty"`
	//  组件属性
	Props interface{} `json:"props,omitempty" xml:"props,omitempty"`
}

func (s GetFieldDefByUuidResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFieldDefByUuidResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFieldDefByUuidResponseBodyResult) SetBehavior(v string) *GetFieldDefByUuidResponseBodyResult {
	s.Behavior = &v
	return s
}

func (s *GetFieldDefByUuidResponseBodyResult) SetChildren(v string) *GetFieldDefByUuidResponseBodyResult {
	s.Children = &v
	return s
}

func (s *GetFieldDefByUuidResponseBodyResult) SetComponentName(v string) *GetFieldDefByUuidResponseBodyResult {
	s.ComponentName = &v
	return s
}

func (s *GetFieldDefByUuidResponseBodyResult) SetFieldId(v string) *GetFieldDefByUuidResponseBodyResult {
	s.FieldId = &v
	return s
}

func (s *GetFieldDefByUuidResponseBodyResult) SetLabel(v interface{}) *GetFieldDefByUuidResponseBodyResult {
	s.Label = v
	return s
}

func (s *GetFieldDefByUuidResponseBodyResult) SetProps(v interface{}) *GetFieldDefByUuidResponseBodyResult {
	s.Props = v
	return s
}

type GetFieldDefByUuidResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFieldDefByUuidResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFieldDefByUuidResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFieldDefByUuidResponse) GoString() string {
	return s.String()
}

func (s *GetFieldDefByUuidResponse) SetHeaders(v map[string]*string) *GetFieldDefByUuidResponse {
	s.Headers = v
	return s
}

func (s *GetFieldDefByUuidResponse) SetBody(v *GetFieldDefByUuidResponseBody) *GetFieldDefByUuidResponse {
	s.Body = v
	return s
}

type GetFormComponentDefinitionListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFormComponentDefinitionListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFormComponentDefinitionListHeaders) GoString() string {
	return s.String()
}

func (s *GetFormComponentDefinitionListHeaders) SetCommonHeaders(v map[string]*string) *GetFormComponentDefinitionListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFormComponentDefinitionListHeaders) SetXAcsDingtalkAccessToken(v string) *GetFormComponentDefinitionListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFormComponentDefinitionListRequest struct {
	// 语言
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 表单版本
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetFormComponentDefinitionListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFormComponentDefinitionListRequest) GoString() string {
	return s.String()
}

func (s *GetFormComponentDefinitionListRequest) SetLanguage(v string) *GetFormComponentDefinitionListRequest {
	s.Language = &v
	return s
}

func (s *GetFormComponentDefinitionListRequest) SetSystemToken(v string) *GetFormComponentDefinitionListRequest {
	s.SystemToken = &v
	return s
}

func (s *GetFormComponentDefinitionListRequest) SetUserId(v string) *GetFormComponentDefinitionListRequest {
	s.UserId = &v
	return s
}

func (s *GetFormComponentDefinitionListRequest) SetVersion(v int64) *GetFormComponentDefinitionListRequest {
	s.Version = &v
	return s
}

type GetFormComponentDefinitionListResponseBody struct {
	Result []*GetFormComponentDefinitionListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetFormComponentDefinitionListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFormComponentDefinitionListResponseBody) GoString() string {
	return s.String()
}

func (s *GetFormComponentDefinitionListResponseBody) SetResult(v []*GetFormComponentDefinitionListResponseBodyResult) *GetFormComponentDefinitionListResponseBody {
	s.Result = v
	return s
}

type GetFormComponentDefinitionListResponseBodyResult struct {
	// componentName
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// key
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// label
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// parentId
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
}

func (s GetFormComponentDefinitionListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFormComponentDefinitionListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFormComponentDefinitionListResponseBodyResult) SetComponentName(v string) *GetFormComponentDefinitionListResponseBodyResult {
	s.ComponentName = &v
	return s
}

func (s *GetFormComponentDefinitionListResponseBodyResult) SetFieldId(v string) *GetFormComponentDefinitionListResponseBodyResult {
	s.FieldId = &v
	return s
}

func (s *GetFormComponentDefinitionListResponseBodyResult) SetLabel(v string) *GetFormComponentDefinitionListResponseBodyResult {
	s.Label = &v
	return s
}

func (s *GetFormComponentDefinitionListResponseBodyResult) SetParentId(v string) *GetFormComponentDefinitionListResponseBodyResult {
	s.ParentId = &v
	return s
}

type GetFormComponentDefinitionListResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFormComponentDefinitionListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFormComponentDefinitionListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFormComponentDefinitionListResponse) GoString() string {
	return s.String()
}

func (s *GetFormComponentDefinitionListResponse) SetHeaders(v map[string]*string) *GetFormComponentDefinitionListResponse {
	s.Headers = v
	return s
}

func (s *GetFormComponentDefinitionListResponse) SetBody(v *GetFormComponentDefinitionListResponseBody) *GetFormComponentDefinitionListResponse {
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
	// 语言。可选值：zh_CN/en_US 默认：zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 应用秘钥。在应用数据中获取。
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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

func (s *GetFormDataByIDRequest) SetLanguage(v string) *GetFormDataByIDRequest {
	s.Language = &v
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

type GetFormDataByIDResponseBody struct {
	// 表单数据详情
	FormData map[string]interface{} `json:"formData,omitempty" xml:"formData,omitempty"`
	// 表单实例ID
	FormInstId *string `json:"formInstId,omitempty" xml:"formInstId,omitempty"`
	// 表单ID
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 最后修改时间
	ModifiedTimeGMT *string `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	// 发起人详情
	Originator *GetFormDataByIDResponseBodyOriginator `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
}

func (s GetFormDataByIDResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDResponseBody) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDResponseBody) SetFormData(v map[string]interface{}) *GetFormDataByIDResponseBody {
	s.FormData = v
	return s
}

func (s *GetFormDataByIDResponseBody) SetFormInstId(v string) *GetFormDataByIDResponseBody {
	s.FormInstId = &v
	return s
}

func (s *GetFormDataByIDResponseBody) SetFormUuid(v string) *GetFormDataByIDResponseBody {
	s.FormUuid = &v
	return s
}

func (s *GetFormDataByIDResponseBody) SetModifiedTimeGMT(v string) *GetFormDataByIDResponseBody {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *GetFormDataByIDResponseBody) SetOriginator(v *GetFormDataByIDResponseBodyOriginator) *GetFormDataByIDResponseBody {
	s.Originator = v
	return s
}

type GetFormDataByIDResponseBodyOriginator struct {
	// 部门名称
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 邮箱
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// 用户名
	Name *GetFormDataByIDResponseBodyOriginatorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	// 用户工号
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetFormDataByIDResponseBodyOriginator) String() string {
	return tea.Prettify(s)
}

func (s GetFormDataByIDResponseBodyOriginator) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDResponseBodyOriginator) SetDepartmentName(v string) *GetFormDataByIDResponseBodyOriginator {
	s.DepartmentName = &v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginator) SetEmail(v string) *GetFormDataByIDResponseBodyOriginator {
	s.Email = &v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginator) SetName(v *GetFormDataByIDResponseBodyOriginatorName) *GetFormDataByIDResponseBodyOriginator {
	s.Name = v
	return s
}

func (s *GetFormDataByIDResponseBodyOriginator) SetUserId(v string) *GetFormDataByIDResponseBodyOriginator {
	s.UserId = &v
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

type GetFormListInAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFormListInAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFormListInAppHeaders) GoString() string {
	return s.String()
}

func (s *GetFormListInAppHeaders) SetCommonHeaders(v map[string]*string) *GetFormListInAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFormListInAppHeaders) SetXAcsDingtalkAccessToken(v string) *GetFormListInAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFormListInAppRequest struct {
	// 应用编码。应用唯一标识。如：APP_XXX
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 支持两种表单类型。
	// receipt :  普通单据表单
	// process ： 流程表单
	// 如需查询多种类型，可用英文逗号分隔。
	// 不传时默认单据和流程均返回。
	FormTypes *string `json:"formTypes,omitempty" xml:"formTypes,omitempty"`
	// 页码，不填默认为1。
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 单页返回的条目数，最大值100。
	// 不填默认为100。
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 应用秘钥。在应用设置-部署运维-应用密钥中获取。
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 操作人userId。
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetFormListInAppRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFormListInAppRequest) GoString() string {
	return s.String()
}

func (s *GetFormListInAppRequest) SetAppType(v string) *GetFormListInAppRequest {
	s.AppType = &v
	return s
}

func (s *GetFormListInAppRequest) SetFormTypes(v string) *GetFormListInAppRequest {
	s.FormTypes = &v
	return s
}

func (s *GetFormListInAppRequest) SetPageNumber(v int32) *GetFormListInAppRequest {
	s.PageNumber = &v
	return s
}

func (s *GetFormListInAppRequest) SetPageSize(v int32) *GetFormListInAppRequest {
	s.PageSize = &v
	return s
}

func (s *GetFormListInAppRequest) SetSystemToken(v string) *GetFormListInAppRequest {
	s.SystemToken = &v
	return s
}

func (s *GetFormListInAppRequest) SetUserId(v string) *GetFormListInAppRequest {
	s.UserId = &v
	return s
}

type GetFormListInAppResponseBody struct {
	// 接口返回的结果
	Result *GetFormListInAppResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// 是否成功，true代表成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetFormListInAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFormListInAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetFormListInAppResponseBody) SetResult(v *GetFormListInAppResponseBodyResult) *GetFormListInAppResponseBody {
	s.Result = v
	return s
}

func (s *GetFormListInAppResponseBody) SetSuccess(v bool) *GetFormListInAppResponseBody {
	s.Success = &v
	return s
}

type GetFormListInAppResponseBodyResult struct {
	// 分页参数，当前页码
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// 返回的表单列表
	Data []*GetFormListInAppResponseBodyResultData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 符合条件的总数目
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetFormListInAppResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFormListInAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFormListInAppResponseBodyResult) SetCurrentPage(v int32) *GetFormListInAppResponseBodyResult {
	s.CurrentPage = &v
	return s
}

func (s *GetFormListInAppResponseBodyResult) SetData(v []*GetFormListInAppResponseBodyResultData) *GetFormListInAppResponseBodyResult {
	s.Data = v
	return s
}

func (s *GetFormListInAppResponseBodyResult) SetTotalCount(v int32) *GetFormListInAppResponseBodyResult {
	s.TotalCount = &v
	return s
}

type GetFormListInAppResponseBodyResultData struct {
	Creator   *string                                      `json:"creator,omitempty" xml:"creator,omitempty"`
	FormType  *string                                      `json:"formType,omitempty" xml:"formType,omitempty"`
	FormUuid  *string                                      `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	GmtCreate *string                                      `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	Title     *GetFormListInAppResponseBodyResultDataTitle `json:"title,omitempty" xml:"title,omitempty" type:"Struct"`
}

func (s GetFormListInAppResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s GetFormListInAppResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *GetFormListInAppResponseBodyResultData) SetCreator(v string) *GetFormListInAppResponseBodyResultData {
	s.Creator = &v
	return s
}

func (s *GetFormListInAppResponseBodyResultData) SetFormType(v string) *GetFormListInAppResponseBodyResultData {
	s.FormType = &v
	return s
}

func (s *GetFormListInAppResponseBodyResultData) SetFormUuid(v string) *GetFormListInAppResponseBodyResultData {
	s.FormUuid = &v
	return s
}

func (s *GetFormListInAppResponseBodyResultData) SetGmtCreate(v string) *GetFormListInAppResponseBodyResultData {
	s.GmtCreate = &v
	return s
}

func (s *GetFormListInAppResponseBodyResultData) SetTitle(v *GetFormListInAppResponseBodyResultDataTitle) *GetFormListInAppResponseBodyResultData {
	s.Title = v
	return s
}

type GetFormListInAppResponseBodyResultDataTitle struct {
	EnUS *string `json:"enUS,omitempty" xml:"enUS,omitempty"`
	ZhCN *string `json:"zhCN,omitempty" xml:"zhCN,omitempty"`
}

func (s GetFormListInAppResponseBodyResultDataTitle) String() string {
	return tea.Prettify(s)
}

func (s GetFormListInAppResponseBodyResultDataTitle) GoString() string {
	return s.String()
}

func (s *GetFormListInAppResponseBodyResultDataTitle) SetEnUS(v string) *GetFormListInAppResponseBodyResultDataTitle {
	s.EnUS = &v
	return s
}

func (s *GetFormListInAppResponseBodyResultDataTitle) SetZhCN(v string) *GetFormListInAppResponseBodyResultDataTitle {
	s.ZhCN = &v
	return s
}

type GetFormListInAppResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFormListInAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFormListInAppResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFormListInAppResponse) GoString() string {
	return s.String()
}

func (s *GetFormListInAppResponse) SetHeaders(v map[string]*string) *GetFormListInAppResponse {
	s.Headers = v
	return s
}

func (s *GetFormListInAppResponse) SetBody(v *GetFormListInAppResponseBody) *GetFormListInAppResponse {
	s.Body = v
	return s
}

type GetInstanceByIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetInstanceByIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceByIdHeaders) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdHeaders) SetCommonHeaders(v map[string]*string) *GetInstanceByIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInstanceByIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetInstanceByIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetInstanceByIdRequest struct {
	// 应用ID
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 语言
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetInstanceByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceByIdRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdRequest) SetAppType(v string) *GetInstanceByIdRequest {
	s.AppType = &v
	return s
}

func (s *GetInstanceByIdRequest) SetLanguage(v string) *GetInstanceByIdRequest {
	s.Language = &v
	return s
}

func (s *GetInstanceByIdRequest) SetSystemToken(v string) *GetInstanceByIdRequest {
	s.SystemToken = &v
	return s
}

func (s *GetInstanceByIdRequest) SetUserId(v string) *GetInstanceByIdRequest {
	s.UserId = &v
	return s
}

type GetInstanceByIdResponseBody struct {
	// actionExecutor
	ActionExecutor []*GetInstanceByIdResponseBodyActionExecutor `json:"actionExecutor,omitempty" xml:"actionExecutor,omitempty" type:"Repeated"`
	// approvedResult
	ApprovedResult *string `json:"approvedResult,omitempty" xml:"approvedResult,omitempty"`
	// 创建时间
	CreateTimeGMT *string `json:"createTimeGMT,omitempty" xml:"createTimeGMT,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// formUuid
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// instanceStatus
	InstanceStatus *string `json:"instanceStatus,omitempty" xml:"instanceStatus,omitempty"`
	// 修改时间
	ModifiedTimeGMT *string `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	// originator
	Originator *GetInstanceByIdResponseBodyOriginator `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	// processCode
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// processInstanceId
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// title
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// version
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetInstanceByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdResponseBody) SetActionExecutor(v []*GetInstanceByIdResponseBodyActionExecutor) *GetInstanceByIdResponseBody {
	s.ActionExecutor = v
	return s
}

func (s *GetInstanceByIdResponseBody) SetApprovedResult(v string) *GetInstanceByIdResponseBody {
	s.ApprovedResult = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetCreateTimeGMT(v string) *GetInstanceByIdResponseBody {
	s.CreateTimeGMT = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetData(v map[string]interface{}) *GetInstanceByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceByIdResponseBody) SetFormUuid(v string) *GetInstanceByIdResponseBody {
	s.FormUuid = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetInstanceStatus(v string) *GetInstanceByIdResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetModifiedTimeGMT(v string) *GetInstanceByIdResponseBody {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetOriginator(v *GetInstanceByIdResponseBodyOriginator) *GetInstanceByIdResponseBody {
	s.Originator = v
	return s
}

func (s *GetInstanceByIdResponseBody) SetProcessCode(v string) *GetInstanceByIdResponseBody {
	s.ProcessCode = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetProcessInstanceId(v string) *GetInstanceByIdResponseBody {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetTitle(v string) *GetInstanceByIdResponseBody {
	s.Title = &v
	return s
}

func (s *GetInstanceByIdResponseBody) SetVersion(v int64) *GetInstanceByIdResponseBody {
	s.Version = &v
	return s
}

type GetInstanceByIdResponseBodyActionExecutor struct {
	// deptName
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// email
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// name
	Name *GetInstanceByIdResponseBodyActionExecutorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	// userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetInstanceByIdResponseBodyActionExecutor) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceByIdResponseBodyActionExecutor) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdResponseBodyActionExecutor) SetDeptName(v string) *GetInstanceByIdResponseBodyActionExecutor {
	s.DeptName = &v
	return s
}

func (s *GetInstanceByIdResponseBodyActionExecutor) SetEmail(v string) *GetInstanceByIdResponseBodyActionExecutor {
	s.Email = &v
	return s
}

func (s *GetInstanceByIdResponseBodyActionExecutor) SetName(v *GetInstanceByIdResponseBodyActionExecutorName) *GetInstanceByIdResponseBodyActionExecutor {
	s.Name = v
	return s
}

func (s *GetInstanceByIdResponseBodyActionExecutor) SetUserId(v string) *GetInstanceByIdResponseBodyActionExecutor {
	s.UserId = &v
	return s
}

type GetInstanceByIdResponseBodyActionExecutorName struct {
	// 中文名称
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// 英文名称
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	// type
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetInstanceByIdResponseBodyActionExecutorName) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceByIdResponseBodyActionExecutorName) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdResponseBodyActionExecutorName) SetNameInChinese(v string) *GetInstanceByIdResponseBodyActionExecutorName {
	s.NameInChinese = &v
	return s
}

func (s *GetInstanceByIdResponseBodyActionExecutorName) SetNameInEnglish(v string) *GetInstanceByIdResponseBodyActionExecutorName {
	s.NameInEnglish = &v
	return s
}

func (s *GetInstanceByIdResponseBodyActionExecutorName) SetType(v string) *GetInstanceByIdResponseBodyActionExecutorName {
	s.Type = &v
	return s
}

type GetInstanceByIdResponseBodyOriginator struct {
	// deptName
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// email
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// name
	Name *GetInstanceByIdResponseBodyOriginatorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	// userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetInstanceByIdResponseBodyOriginator) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceByIdResponseBodyOriginator) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdResponseBodyOriginator) SetDeptName(v string) *GetInstanceByIdResponseBodyOriginator {
	s.DeptName = &v
	return s
}

func (s *GetInstanceByIdResponseBodyOriginator) SetEmail(v string) *GetInstanceByIdResponseBodyOriginator {
	s.Email = &v
	return s
}

func (s *GetInstanceByIdResponseBodyOriginator) SetName(v *GetInstanceByIdResponseBodyOriginatorName) *GetInstanceByIdResponseBodyOriginator {
	s.Name = v
	return s
}

func (s *GetInstanceByIdResponseBodyOriginator) SetUserId(v string) *GetInstanceByIdResponseBodyOriginator {
	s.UserId = &v
	return s
}

type GetInstanceByIdResponseBodyOriginatorName struct {
	// 中文名称
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// 英文名称
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	// type
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetInstanceByIdResponseBodyOriginatorName) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceByIdResponseBodyOriginatorName) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdResponseBodyOriginatorName) SetNameInChinese(v string) *GetInstanceByIdResponseBodyOriginatorName {
	s.NameInChinese = &v
	return s
}

func (s *GetInstanceByIdResponseBodyOriginatorName) SetNameInEnglish(v string) *GetInstanceByIdResponseBodyOriginatorName {
	s.NameInEnglish = &v
	return s
}

func (s *GetInstanceByIdResponseBodyOriginatorName) SetType(v string) *GetInstanceByIdResponseBodyOriginatorName {
	s.Type = &v
	return s
}

type GetInstanceByIdResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstanceByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceByIdResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdResponse) SetHeaders(v map[string]*string) *GetInstanceByIdResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceByIdResponse) SetBody(v *GetInstanceByIdResponseBody) *GetInstanceByIdResponse {
	s.Body = v
	return s
}

type GetInstanceIdListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetInstanceIdListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceIdListHeaders) GoString() string {
	return s.String()
}

func (s *GetInstanceIdListHeaders) SetCommonHeaders(v map[string]*string) *GetInstanceIdListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInstanceIdListHeaders) SetXAcsDingtalkAccessToken(v string) *GetInstanceIdListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetInstanceIdListRequest struct {
	// 应用ID
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 流程审批结果
	ApprovedResult *string `json:"approvedResult,omitempty" xml:"approvedResult,omitempty"`
	// createFrom和createTo两个时间构造一个时间段。查询在该时间段创建的数据列表
	CreateFromTimeGMT *string `json:"createFromTimeGMT,omitempty" xml:"createFromTimeGMT,omitempty"`
	// createFrom和createTo两个时间构造一个时间段。查询在该时间段创建的数据列表。
	CreateToTimeGMT *string `json:"createToTimeGMT,omitempty" xml:"createToTimeGMT,omitempty"`
	// 表单ID
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 实例状态
	InstanceStatus *string `json:"instanceStatus,omitempty" xml:"instanceStatus,omitempty"`
	// 语言
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// modifiedFrom和modifiedTo构成一个时间段，查询在该时间段有修改的数据列表
	ModifiedFromTimeGMT *string `json:"modifiedFromTimeGMT,omitempty" xml:"modifiedFromTimeGMT,omitempty"`
	// modifiedFrom和modifiedTo构成一个时间段，查询在该时间段有修改的数据列表。
	ModifiedToTimeGMT *string `json:"modifiedToTimeGMT,omitempty" xml:"modifiedToTimeGMT,omitempty"`
	// 根据流程发起人工号查询
	OriginatorId *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	// 根据表单内组件值查询
	SearchFieldJson *string `json:"searchFieldJson,omitempty" xml:"searchFieldJson,omitempty"`
	// 应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 任务ID
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// 钉钉的userId
	UserId     *string `json:"userId,omitempty" xml:"userId,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetInstanceIdListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceIdListRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceIdListRequest) SetAppType(v string) *GetInstanceIdListRequest {
	s.AppType = &v
	return s
}

func (s *GetInstanceIdListRequest) SetApprovedResult(v string) *GetInstanceIdListRequest {
	s.ApprovedResult = &v
	return s
}

func (s *GetInstanceIdListRequest) SetCreateFromTimeGMT(v string) *GetInstanceIdListRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *GetInstanceIdListRequest) SetCreateToTimeGMT(v string) *GetInstanceIdListRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *GetInstanceIdListRequest) SetFormUuid(v string) *GetInstanceIdListRequest {
	s.FormUuid = &v
	return s
}

func (s *GetInstanceIdListRequest) SetInstanceStatus(v string) *GetInstanceIdListRequest {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstanceIdListRequest) SetLanguage(v string) *GetInstanceIdListRequest {
	s.Language = &v
	return s
}

func (s *GetInstanceIdListRequest) SetModifiedFromTimeGMT(v string) *GetInstanceIdListRequest {
	s.ModifiedFromTimeGMT = &v
	return s
}

func (s *GetInstanceIdListRequest) SetModifiedToTimeGMT(v string) *GetInstanceIdListRequest {
	s.ModifiedToTimeGMT = &v
	return s
}

func (s *GetInstanceIdListRequest) SetOriginatorId(v string) *GetInstanceIdListRequest {
	s.OriginatorId = &v
	return s
}

func (s *GetInstanceIdListRequest) SetSearchFieldJson(v string) *GetInstanceIdListRequest {
	s.SearchFieldJson = &v
	return s
}

func (s *GetInstanceIdListRequest) SetSystemToken(v string) *GetInstanceIdListRequest {
	s.SystemToken = &v
	return s
}

func (s *GetInstanceIdListRequest) SetTaskId(v string) *GetInstanceIdListRequest {
	s.TaskId = &v
	return s
}

func (s *GetInstanceIdListRequest) SetUserId(v string) *GetInstanceIdListRequest {
	s.UserId = &v
	return s
}

func (s *GetInstanceIdListRequest) SetPageNumber(v int32) *GetInstanceIdListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetInstanceIdListRequest) SetPageSize(v int32) *GetInstanceIdListRequest {
	s.PageSize = &v
	return s
}

type GetInstanceIdListResponseBody struct {
	// data
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 当前第几页
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 总数量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetInstanceIdListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceIdListResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceIdListResponseBody) SetData(v []*string) *GetInstanceIdListResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceIdListResponseBody) SetPageNumber(v int64) *GetInstanceIdListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetInstanceIdListResponseBody) SetTotalCount(v int64) *GetInstanceIdListResponseBody {
	s.TotalCount = &v
	return s
}

type GetInstanceIdListResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstanceIdListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceIdListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceIdListResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceIdListResponse) SetHeaders(v map[string]*string) *GetInstanceIdListResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceIdListResponse) SetBody(v *GetInstanceIdListResponseBody) *GetInstanceIdListResponse {
	s.Body = v
	return s
}

type GetInstancesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetInstancesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesHeaders) GoString() string {
	return s.String()
}

func (s *GetInstancesHeaders) SetCommonHeaders(v map[string]*string) *GetInstancesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInstancesHeaders) SetXAcsDingtalkAccessToken(v string) *GetInstancesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetInstancesRequest struct {
	// 应用编码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 审批结果。可选值为：agree, disagree。 分别表示：同意， 拒绝。
	ApprovedResult *string `json:"approvedResult,omitempty" xml:"approvedResult,omitempty"`
	// createFrom和createTo两个时间构造一个时间段。查询在该时间段创建的数据列表, 字符串格式，且为yyyy-MM-DD格式
	CreateFromTimeGMT *string `json:"createFromTimeGMT,omitempty" xml:"createFromTimeGMT,omitempty"`
	// createFrom和createTo两个时间构造一个时间段。查询在该时间段创建的数据列表。字符串格式，且为yyyy-MM-DD格式。 和createFrom一起，相当于查询在 2018-01-01到2018-01-31之间(包含01和31号)创建的数据。
	CreateToTimeGMT *string `json:"createToTimeGMT,omitempty" xml:"createToTimeGMT,omitempty"`
	// 表单ID
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 实例状态, 可选值为：RUNNING,TERMINATED,COMPLETED,ERROR。 分别代表：运行中，已终止，已完成，异常。
	InstanceStatus *string `json:"instanceStatus,omitempty" xml:"instanceStatus,omitempty"`
	// 语言。可选值：zh_CN/en_US 默认：zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// modifiedFrom和modifiedTo构成一个时间段，查询在该时间段有修改的数据列表。字符串格式，且为yyyy-MM-DD格式
	ModifiedFromTimeGMT *string `json:"modifiedFromTimeGMT,omitempty" xml:"modifiedFromTimeGMT,omitempty"`
	// modifiedFrom和modifiedTo构成一个时间段，查询在该时间段有修改的数据列表。字符串格式，且为yyyy-MM-DD格式。 和modifiedFrom一起，相当于查询在 2018-01-01到2018-01-31之间(包含01和31号)被修改的数据。
	ModifiedToTimeGMT *string `json:"modifiedToTimeGMT,omitempty" xml:"modifiedToTimeGMT,omitempty"`
	// 排序规则
	OrderConfigJson *string `json:"orderConfigJson,omitempty" xml:"orderConfigJson,omitempty"`
	// 数据提交人的钉钉userId
	OriginatorId *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	// 查询过滤条件，支持2种模式的过滤规则
	SearchFieldJson *string `json:"searchFieldJson,omitempty" xml:"searchFieldJson,omitempty"`
	// 应用秘钥。在应用数据中获取。
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 任务ID。一般用不到。
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 当前第几页
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页多少条数据
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesRequest) GoString() string {
	return s.String()
}

func (s *GetInstancesRequest) SetAppType(v string) *GetInstancesRequest {
	s.AppType = &v
	return s
}

func (s *GetInstancesRequest) SetApprovedResult(v string) *GetInstancesRequest {
	s.ApprovedResult = &v
	return s
}

func (s *GetInstancesRequest) SetCreateFromTimeGMT(v string) *GetInstancesRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *GetInstancesRequest) SetCreateToTimeGMT(v string) *GetInstancesRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *GetInstancesRequest) SetFormUuid(v string) *GetInstancesRequest {
	s.FormUuid = &v
	return s
}

func (s *GetInstancesRequest) SetInstanceStatus(v string) *GetInstancesRequest {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstancesRequest) SetLanguage(v string) *GetInstancesRequest {
	s.Language = &v
	return s
}

func (s *GetInstancesRequest) SetModifiedFromTimeGMT(v string) *GetInstancesRequest {
	s.ModifiedFromTimeGMT = &v
	return s
}

func (s *GetInstancesRequest) SetModifiedToTimeGMT(v string) *GetInstancesRequest {
	s.ModifiedToTimeGMT = &v
	return s
}

func (s *GetInstancesRequest) SetOrderConfigJson(v string) *GetInstancesRequest {
	s.OrderConfigJson = &v
	return s
}

func (s *GetInstancesRequest) SetOriginatorId(v string) *GetInstancesRequest {
	s.OriginatorId = &v
	return s
}

func (s *GetInstancesRequest) SetSearchFieldJson(v string) *GetInstancesRequest {
	s.SearchFieldJson = &v
	return s
}

func (s *GetInstancesRequest) SetSystemToken(v string) *GetInstancesRequest {
	s.SystemToken = &v
	return s
}

func (s *GetInstancesRequest) SetTaskId(v string) *GetInstancesRequest {
	s.TaskId = &v
	return s
}

func (s *GetInstancesRequest) SetUserId(v string) *GetInstancesRequest {
	s.UserId = &v
	return s
}

func (s *GetInstancesRequest) SetPageNumber(v int32) *GetInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *GetInstancesRequest) SetPageSize(v int32) *GetInstancesRequest {
	s.PageSize = &v
	return s
}

type GetInstancesResponseBody struct {
	// data
	Data []*GetInstancesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 当前第几页
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 总数量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstancesResponseBody) SetData(v []*GetInstancesResponseBodyData) *GetInstancesResponseBody {
	s.Data = v
	return s
}

func (s *GetInstancesResponseBody) SetPageNumber(v int64) *GetInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetInstancesResponseBody) SetTotalCount(v int64) *GetInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type GetInstancesResponseBodyData struct {
	// actioners
	ActionExecutor []*GetInstancesResponseBodyDataActionExecutor `json:"actionExecutor,omitempty" xml:"actionExecutor,omitempty" type:"Repeated"`
	// approvedResult
	ApprovedResult *string `json:"approvedResult,omitempty" xml:"approvedResult,omitempty"`
	// 创建时间
	CreateTimeGMT *string `json:"createTimeGMT,omitempty" xml:"createTimeGMT,omitempty"`
	// data
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// formUuid
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// instanceStatus
	InstanceStatus *string `json:"instanceStatus,omitempty" xml:"instanceStatus,omitempty"`
	// 修改时间
	ModifiedTimeGMT *string `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	// originator
	Originator *GetInstancesResponseBodyDataOriginator `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	// processCode
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// processInstanceId
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// title
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// version
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstancesResponseBodyData) SetActionExecutor(v []*GetInstancesResponseBodyDataActionExecutor) *GetInstancesResponseBodyData {
	s.ActionExecutor = v
	return s
}

func (s *GetInstancesResponseBodyData) SetApprovedResult(v string) *GetInstancesResponseBodyData {
	s.ApprovedResult = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetCreateTimeGMT(v string) *GetInstancesResponseBodyData {
	s.CreateTimeGMT = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetData(v map[string]interface{}) *GetInstancesResponseBodyData {
	s.Data = v
	return s
}

func (s *GetInstancesResponseBodyData) SetFormUuid(v string) *GetInstancesResponseBodyData {
	s.FormUuid = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetInstanceStatus(v string) *GetInstancesResponseBodyData {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetModifiedTimeGMT(v string) *GetInstancesResponseBodyData {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetOriginator(v *GetInstancesResponseBodyDataOriginator) *GetInstancesResponseBodyData {
	s.Originator = v
	return s
}

func (s *GetInstancesResponseBodyData) SetProcessCode(v string) *GetInstancesResponseBodyData {
	s.ProcessCode = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetProcessInstanceId(v string) *GetInstancesResponseBodyData {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetTitle(v string) *GetInstancesResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetInstancesResponseBodyData) SetVersion(v int64) *GetInstancesResponseBodyData {
	s.Version = &v
	return s
}

type GetInstancesResponseBodyDataActionExecutor struct {
	// deptName
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// email
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// name
	Name *GetInstancesResponseBodyDataActionExecutorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	// userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetInstancesResponseBodyDataActionExecutor) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesResponseBodyDataActionExecutor) GoString() string {
	return s.String()
}

func (s *GetInstancesResponseBodyDataActionExecutor) SetDeptName(v string) *GetInstancesResponseBodyDataActionExecutor {
	s.DeptName = &v
	return s
}

func (s *GetInstancesResponseBodyDataActionExecutor) SetEmail(v string) *GetInstancesResponseBodyDataActionExecutor {
	s.Email = &v
	return s
}

func (s *GetInstancesResponseBodyDataActionExecutor) SetName(v *GetInstancesResponseBodyDataActionExecutorName) *GetInstancesResponseBodyDataActionExecutor {
	s.Name = v
	return s
}

func (s *GetInstancesResponseBodyDataActionExecutor) SetUserId(v string) *GetInstancesResponseBodyDataActionExecutor {
	s.UserId = &v
	return s
}

type GetInstancesResponseBodyDataActionExecutorName struct {
	// 中文名称
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// 英文名称
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	// type
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetInstancesResponseBodyDataActionExecutorName) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesResponseBodyDataActionExecutorName) GoString() string {
	return s.String()
}

func (s *GetInstancesResponseBodyDataActionExecutorName) SetNameInChinese(v string) *GetInstancesResponseBodyDataActionExecutorName {
	s.NameInChinese = &v
	return s
}

func (s *GetInstancesResponseBodyDataActionExecutorName) SetNameInEnglish(v string) *GetInstancesResponseBodyDataActionExecutorName {
	s.NameInEnglish = &v
	return s
}

func (s *GetInstancesResponseBodyDataActionExecutorName) SetType(v string) *GetInstancesResponseBodyDataActionExecutorName {
	s.Type = &v
	return s
}

type GetInstancesResponseBodyDataOriginator struct {
	// deptName
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// email
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// name
	Name *GetInstancesResponseBodyDataOriginatorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	// userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetInstancesResponseBodyDataOriginator) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesResponseBodyDataOriginator) GoString() string {
	return s.String()
}

func (s *GetInstancesResponseBodyDataOriginator) SetDeptName(v string) *GetInstancesResponseBodyDataOriginator {
	s.DeptName = &v
	return s
}

func (s *GetInstancesResponseBodyDataOriginator) SetEmail(v string) *GetInstancesResponseBodyDataOriginator {
	s.Email = &v
	return s
}

func (s *GetInstancesResponseBodyDataOriginator) SetName(v *GetInstancesResponseBodyDataOriginatorName) *GetInstancesResponseBodyDataOriginator {
	s.Name = v
	return s
}

func (s *GetInstancesResponseBodyDataOriginator) SetUserId(v string) *GetInstancesResponseBodyDataOriginator {
	s.UserId = &v
	return s
}

type GetInstancesResponseBodyDataOriginatorName struct {
	// 中文名称
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// 英文名称
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	// type
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetInstancesResponseBodyDataOriginatorName) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesResponseBodyDataOriginatorName) GoString() string {
	return s.String()
}

func (s *GetInstancesResponseBodyDataOriginatorName) SetNameInChinese(v string) *GetInstancesResponseBodyDataOriginatorName {
	s.NameInChinese = &v
	return s
}

func (s *GetInstancesResponseBodyDataOriginatorName) SetNameInEnglish(v string) *GetInstancesResponseBodyDataOriginatorName {
	s.NameInEnglish = &v
	return s
}

func (s *GetInstancesResponseBodyDataOriginatorName) SetType(v string) *GetInstancesResponseBodyDataOriginatorName {
	s.Type = &v
	return s
}

type GetInstancesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesResponse) GoString() string {
	return s.String()
}

func (s *GetInstancesResponse) SetHeaders(v map[string]*string) *GetInstancesResponse {
	s.Headers = v
	return s
}

func (s *GetInstancesResponse) SetBody(v *GetInstancesResponseBody) *GetInstancesResponse {
	s.Body = v
	return s
}

type GetInstancesByIdListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetInstancesByIdListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesByIdListHeaders) GoString() string {
	return s.String()
}

func (s *GetInstancesByIdListHeaders) SetCommonHeaders(v map[string]*string) *GetInstancesByIdListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInstancesByIdListHeaders) SetXAcsDingtalkAccessToken(v string) *GetInstancesByIdListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetInstancesByIdListRequest struct {
	// 应用ID
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 语言
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 流程实例ID列表，多个用,分割
	ProcessInstanceIds *string `json:"processInstanceIds,omitempty" xml:"processInstanceIds,omitempty"`
	// 应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetInstancesByIdListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesByIdListRequest) GoString() string {
	return s.String()
}

func (s *GetInstancesByIdListRequest) SetAppType(v string) *GetInstancesByIdListRequest {
	s.AppType = &v
	return s
}

func (s *GetInstancesByIdListRequest) SetLanguage(v string) *GetInstancesByIdListRequest {
	s.Language = &v
	return s
}

func (s *GetInstancesByIdListRequest) SetProcessInstanceIds(v string) *GetInstancesByIdListRequest {
	s.ProcessInstanceIds = &v
	return s
}

func (s *GetInstancesByIdListRequest) SetSystemToken(v string) *GetInstancesByIdListRequest {
	s.SystemToken = &v
	return s
}

func (s *GetInstancesByIdListRequest) SetUserId(v string) *GetInstancesByIdListRequest {
	s.UserId = &v
	return s
}

type GetInstancesByIdListResponseBody struct {
	// 流程实例列表
	Result []*GetInstancesByIdListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetInstancesByIdListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesByIdListResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstancesByIdListResponseBody) SetResult(v []*GetInstancesByIdListResponseBodyResult) *GetInstancesByIdListResponseBody {
	s.Result = v
	return s
}

type GetInstancesByIdListResponseBodyResult struct {
	// 流程实例当前任务执行人列表
	ActionExecutor []*GetInstancesByIdListResponseBodyResultActionExecutor `json:"actionExecutor,omitempty" xml:"actionExecutor,omitempty" type:"Repeated"`
	// 流程结束时的审批结论
	ApprovedResult *string `json:"approvedResult,omitempty" xml:"approvedResult,omitempty"`
	// 表单数据
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// 流程表单ID
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 实例状态
	InstanceStatus *string `json:"instanceStatus,omitempty" xml:"instanceStatus,omitempty"`
	// 发起人信息
	Originator *GetInstancesByIdListResponseBodyResultOriginator `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	// 流程Code
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// 实例ID
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// 实例标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetInstancesByIdListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesByIdListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetInstancesByIdListResponseBodyResult) SetActionExecutor(v []*GetInstancesByIdListResponseBodyResultActionExecutor) *GetInstancesByIdListResponseBodyResult {
	s.ActionExecutor = v
	return s
}

func (s *GetInstancesByIdListResponseBodyResult) SetApprovedResult(v string) *GetInstancesByIdListResponseBodyResult {
	s.ApprovedResult = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResult) SetData(v map[string]interface{}) *GetInstancesByIdListResponseBodyResult {
	s.Data = v
	return s
}

func (s *GetInstancesByIdListResponseBodyResult) SetFormUuid(v string) *GetInstancesByIdListResponseBodyResult {
	s.FormUuid = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResult) SetInstanceStatus(v string) *GetInstancesByIdListResponseBodyResult {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResult) SetOriginator(v *GetInstancesByIdListResponseBodyResultOriginator) *GetInstancesByIdListResponseBodyResult {
	s.Originator = v
	return s
}

func (s *GetInstancesByIdListResponseBodyResult) SetProcessCode(v string) *GetInstancesByIdListResponseBodyResult {
	s.ProcessCode = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResult) SetProcessInstanceId(v string) *GetInstancesByIdListResponseBodyResult {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResult) SetTitle(v string) *GetInstancesByIdListResponseBodyResult {
	s.Title = &v
	return s
}

type GetInstancesByIdListResponseBodyResultActionExecutor struct {
	// 部门名称
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 邮箱
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// 用户名
	Name *GetInstancesByIdListResponseBodyResultActionExecutorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	// 用户工号
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetInstancesByIdListResponseBodyResultActionExecutor) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesByIdListResponseBodyResultActionExecutor) GoString() string {
	return s.String()
}

func (s *GetInstancesByIdListResponseBodyResultActionExecutor) SetDepartmentName(v string) *GetInstancesByIdListResponseBodyResultActionExecutor {
	s.DepartmentName = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResultActionExecutor) SetEmail(v string) *GetInstancesByIdListResponseBodyResultActionExecutor {
	s.Email = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResultActionExecutor) SetName(v *GetInstancesByIdListResponseBodyResultActionExecutorName) *GetInstancesByIdListResponseBodyResultActionExecutor {
	s.Name = v
	return s
}

func (s *GetInstancesByIdListResponseBodyResultActionExecutor) SetUserId(v string) *GetInstancesByIdListResponseBodyResultActionExecutor {
	s.UserId = &v
	return s
}

type GetInstancesByIdListResponseBodyResultActionExecutorName struct {
	// 中文名称
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// 英文名称
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	// 国际化
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetInstancesByIdListResponseBodyResultActionExecutorName) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesByIdListResponseBodyResultActionExecutorName) GoString() string {
	return s.String()
}

func (s *GetInstancesByIdListResponseBodyResultActionExecutorName) SetNameInChinese(v string) *GetInstancesByIdListResponseBodyResultActionExecutorName {
	s.NameInChinese = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResultActionExecutorName) SetNameInEnglish(v string) *GetInstancesByIdListResponseBodyResultActionExecutorName {
	s.NameInEnglish = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResultActionExecutorName) SetType(v string) *GetInstancesByIdListResponseBodyResultActionExecutorName {
	s.Type = &v
	return s
}

type GetInstancesByIdListResponseBodyResultOriginator struct {
	// 部门名称
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 邮箱
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// 用户名
	Name *GetInstancesByIdListResponseBodyResultOriginatorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	// 用户工号
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetInstancesByIdListResponseBodyResultOriginator) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesByIdListResponseBodyResultOriginator) GoString() string {
	return s.String()
}

func (s *GetInstancesByIdListResponseBodyResultOriginator) SetDepartmentName(v string) *GetInstancesByIdListResponseBodyResultOriginator {
	s.DepartmentName = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResultOriginator) SetEmail(v string) *GetInstancesByIdListResponseBodyResultOriginator {
	s.Email = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResultOriginator) SetName(v *GetInstancesByIdListResponseBodyResultOriginatorName) *GetInstancesByIdListResponseBodyResultOriginator {
	s.Name = v
	return s
}

func (s *GetInstancesByIdListResponseBodyResultOriginator) SetUserId(v string) *GetInstancesByIdListResponseBodyResultOriginator {
	s.UserId = &v
	return s
}

type GetInstancesByIdListResponseBodyResultOriginatorName struct {
	// 中文名称
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// 英文名称
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	// 国际化
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetInstancesByIdListResponseBodyResultOriginatorName) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesByIdListResponseBodyResultOriginatorName) GoString() string {
	return s.String()
}

func (s *GetInstancesByIdListResponseBodyResultOriginatorName) SetNameInChinese(v string) *GetInstancesByIdListResponseBodyResultOriginatorName {
	s.NameInChinese = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResultOriginatorName) SetNameInEnglish(v string) *GetInstancesByIdListResponseBodyResultOriginatorName {
	s.NameInEnglish = &v
	return s
}

func (s *GetInstancesByIdListResponseBodyResultOriginatorName) SetType(v string) *GetInstancesByIdListResponseBodyResultOriginatorName {
	s.Type = &v
	return s
}

type GetInstancesByIdListResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstancesByIdListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstancesByIdListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstancesByIdListResponse) GoString() string {
	return s.String()
}

func (s *GetInstancesByIdListResponse) SetHeaders(v map[string]*string) *GetInstancesByIdListResponse {
	s.Headers = v
	return s
}

func (s *GetInstancesByIdListResponse) SetBody(v *GetInstancesByIdListResponseBody) *GetInstancesByIdListResponse {
	s.Body = v
	return s
}

type GetMeCorpSubmissionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetMeCorpSubmissionHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetMeCorpSubmissionHeaders) GoString() string {
	return s.String()
}

func (s *GetMeCorpSubmissionHeaders) SetCommonHeaders(v map[string]*string) *GetMeCorpSubmissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMeCorpSubmissionHeaders) SetXAcsDingtalkAccessToken(v string) *GetMeCorpSubmissionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetMeCorpSubmissionRequest struct {
	// 应用标识列表
	AppTypes *string `json:"appTypes,omitempty" xml:"appTypes,omitempty"`
	// 企业ID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 创建时间开始
	CreateFromTimeGMT *int64 `json:"createFromTimeGMT,omitempty" xml:"createFromTimeGMT,omitempty"`
	// 创建时间结束
	CreateToTimeGMT *int64 `json:"createToTimeGMT,omitempty" xml:"createToTimeGMT,omitempty"`
	// 关键词
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 语言环境
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 当前页
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页记录数
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 流程code列表
	ProcessCodes *string `json:"processCodes,omitempty" xml:"processCodes,omitempty"`
	// 验权token
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s GetMeCorpSubmissionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMeCorpSubmissionRequest) GoString() string {
	return s.String()
}

func (s *GetMeCorpSubmissionRequest) SetAppTypes(v string) *GetMeCorpSubmissionRequest {
	s.AppTypes = &v
	return s
}

func (s *GetMeCorpSubmissionRequest) SetCorpId(v string) *GetMeCorpSubmissionRequest {
	s.CorpId = &v
	return s
}

func (s *GetMeCorpSubmissionRequest) SetCreateFromTimeGMT(v int64) *GetMeCorpSubmissionRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *GetMeCorpSubmissionRequest) SetCreateToTimeGMT(v int64) *GetMeCorpSubmissionRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *GetMeCorpSubmissionRequest) SetKeyword(v string) *GetMeCorpSubmissionRequest {
	s.Keyword = &v
	return s
}

func (s *GetMeCorpSubmissionRequest) SetLanguage(v string) *GetMeCorpSubmissionRequest {
	s.Language = &v
	return s
}

func (s *GetMeCorpSubmissionRequest) SetPageNumber(v int32) *GetMeCorpSubmissionRequest {
	s.PageNumber = &v
	return s
}

func (s *GetMeCorpSubmissionRequest) SetPageSize(v int32) *GetMeCorpSubmissionRequest {
	s.PageSize = &v
	return s
}

func (s *GetMeCorpSubmissionRequest) SetProcessCodes(v string) *GetMeCorpSubmissionRequest {
	s.ProcessCodes = &v
	return s
}

func (s *GetMeCorpSubmissionRequest) SetToken(v string) *GetMeCorpSubmissionRequest {
	s.Token = &v
	return s
}

type GetMeCorpSubmissionResponseBody struct {
	// data
	Data []*GetMeCorpSubmissionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 当前第几页
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 总数量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetMeCorpSubmissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMeCorpSubmissionResponseBody) GoString() string {
	return s.String()
}

func (s *GetMeCorpSubmissionResponseBody) SetData(v []*GetMeCorpSubmissionResponseBodyData) *GetMeCorpSubmissionResponseBody {
	s.Data = v
	return s
}

func (s *GetMeCorpSubmissionResponseBody) SetPageNumber(v int64) *GetMeCorpSubmissionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBody) SetTotalCount(v int64) *GetMeCorpSubmissionResponseBody {
	s.TotalCount = &v
	return s
}

type GetMeCorpSubmissionResponseBodyData struct {
	// actioner
	Actioner []*GetMeCorpSubmissionResponseBodyDataActioner `json:"actioner,omitempty" xml:"actioner,omitempty" type:"Repeated"`
	// actionerId
	ActionerId []*string `json:"actionerId,omitempty" xml:"actionerId,omitempty" type:"Repeated"`
	// actionerName
	ActionerName []*string `json:"actionerName,omitempty" xml:"actionerName,omitempty" type:"Repeated"`
	// appType
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// createTime
	CreateTimeGMT *string `json:"createTimeGMT,omitempty" xml:"createTimeGMT,omitempty"`
	// currentActivityInstances
	CurrentActivityInstances []*GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances `json:"currentActivityInstances,omitempty" xml:"currentActivityInstances,omitempty" type:"Repeated"`
	// dataMap
	DataMap map[string]interface{} `json:"dataMap,omitempty" xml:"dataMap,omitempty"`
	// dataType
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// finishTime
	FinishTimeGMT *string `json:"finishTimeGMT,omitempty" xml:"finishTimeGMT,omitempty"`
	// formInstanceId
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// formUuid
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// instValue
	InstanceValue *string `json:"instanceValue,omitempty" xml:"instanceValue,omitempty"`
	// modifiedTime
	ModifiedTimeGMT *string `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	// originatorAvatar
	OriginatorAvatar *string `json:"originatorAvatar,omitempty" xml:"originatorAvatar,omitempty"`
	// originatorDisplayName
	OriginatorDisplayName *string `json:"originatorDisplayName,omitempty" xml:"originatorDisplayName,omitempty"`
	// originatorId
	OriginatorId *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	// processApprovedResult
	ProcessApprovedResult *string `json:"processApprovedResult,omitempty" xml:"processApprovedResult,omitempty"`
	// processApprovedResultText
	ProcessApprovedResultText *string `json:"processApprovedResultText,omitempty" xml:"processApprovedResultText,omitempty"`
	// processCode
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// processId
	ProcessId *int64 `json:"processId,omitempty" xml:"processId,omitempty"`
	// processInstanceId
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// processInstanceStatus
	ProcessInstanceStatus *string `json:"processInstanceStatus,omitempty" xml:"processInstanceStatus,omitempty"`
	// processInstanceStatusText
	ProcessInstanceStatusText *string `json:"processInstanceStatusText,omitempty" xml:"processInstanceStatusText,omitempty"`
	// processName
	ProcessName *string `json:"processName,omitempty" xml:"processName,omitempty"`
	// title
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// version
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetMeCorpSubmissionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMeCorpSubmissionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMeCorpSubmissionResponseBodyData) SetActioner(v []*GetMeCorpSubmissionResponseBodyDataActioner) *GetMeCorpSubmissionResponseBodyData {
	s.Actioner = v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetActionerId(v []*string) *GetMeCorpSubmissionResponseBodyData {
	s.ActionerId = v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetActionerName(v []*string) *GetMeCorpSubmissionResponseBodyData {
	s.ActionerName = v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetAppType(v string) *GetMeCorpSubmissionResponseBodyData {
	s.AppType = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetCreateTimeGMT(v string) *GetMeCorpSubmissionResponseBodyData {
	s.CreateTimeGMT = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetCurrentActivityInstances(v []*GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances) *GetMeCorpSubmissionResponseBodyData {
	s.CurrentActivityInstances = v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetDataMap(v map[string]interface{}) *GetMeCorpSubmissionResponseBodyData {
	s.DataMap = v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetDataType(v string) *GetMeCorpSubmissionResponseBodyData {
	s.DataType = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetFinishTimeGMT(v string) *GetMeCorpSubmissionResponseBodyData {
	s.FinishTimeGMT = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetFormInstanceId(v string) *GetMeCorpSubmissionResponseBodyData {
	s.FormInstanceId = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetFormUuid(v string) *GetMeCorpSubmissionResponseBodyData {
	s.FormUuid = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetInstanceValue(v string) *GetMeCorpSubmissionResponseBodyData {
	s.InstanceValue = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetModifiedTimeGMT(v string) *GetMeCorpSubmissionResponseBodyData {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetOriginatorAvatar(v string) *GetMeCorpSubmissionResponseBodyData {
	s.OriginatorAvatar = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetOriginatorDisplayName(v string) *GetMeCorpSubmissionResponseBodyData {
	s.OriginatorDisplayName = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetOriginatorId(v string) *GetMeCorpSubmissionResponseBodyData {
	s.OriginatorId = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetProcessApprovedResult(v string) *GetMeCorpSubmissionResponseBodyData {
	s.ProcessApprovedResult = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetProcessApprovedResultText(v string) *GetMeCorpSubmissionResponseBodyData {
	s.ProcessApprovedResultText = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetProcessCode(v string) *GetMeCorpSubmissionResponseBodyData {
	s.ProcessCode = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetProcessId(v int64) *GetMeCorpSubmissionResponseBodyData {
	s.ProcessId = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetProcessInstanceId(v string) *GetMeCorpSubmissionResponseBodyData {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetProcessInstanceStatus(v string) *GetMeCorpSubmissionResponseBodyData {
	s.ProcessInstanceStatus = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetProcessInstanceStatusText(v string) *GetMeCorpSubmissionResponseBodyData {
	s.ProcessInstanceStatusText = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetProcessName(v string) *GetMeCorpSubmissionResponseBodyData {
	s.ProcessName = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetTitle(v string) *GetMeCorpSubmissionResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyData) SetVersion(v int64) *GetMeCorpSubmissionResponseBodyData {
	s.Version = &v
	return s
}

type GetMeCorpSubmissionResponseBodyDataActioner struct {
	// buName
	BuName *string `json:"buName,omitempty" xml:"buName,omitempty"`
	// email
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// empType
	EmployeeType *string `json:"employeeType,omitempty" xml:"employeeType,omitempty"`
	// employeeTypeInformation
	EmployeeTypeInformation *string `json:"employeeTypeInformation,omitempty" xml:"employeeTypeInformation,omitempty"`
	// hrgWorkNo
	HumanResourceGroupWorkNumber *string `json:"humanResourceGroupWorkNumber,omitempty" xml:"humanResourceGroupWorkNumber,omitempty"`
	// isSystemAdmin
	IsSystemAdmin *bool `json:"isSystemAdmin,omitempty" xml:"isSystemAdmin,omitempty"`
	// level
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// nickName
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// orderNum
	OrderNumber *string `json:"orderNumber,omitempty" xml:"orderNumber,omitempty"`
	// personalPhoto
	PersonalPhoto *string `json:"personalPhoto,omitempty" xml:"personalPhoto,omitempty"`
	// personalPhotoUrl
	PersonalPhotoUrl *string `json:"personalPhotoUrl,omitempty" xml:"personalPhotoUrl,omitempty"`
	// pinyinNameAll
	PinyinNameAll *string `json:"pinyinNameAll,omitempty" xml:"pinyinNameAll,omitempty"`
	// pinyinNick
	PinyinNickName *string `json:"pinyinNickName,omitempty" xml:"pinyinNickName,omitempty"`
	// state
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// superUserId
	SuperUserId *string `json:"superUserId,omitempty" xml:"superUserId,omitempty"`
	// tbWang
	TbWang *string `json:"tbWang,omitempty" xml:"tbWang,omitempty"`
	// userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetMeCorpSubmissionResponseBodyDataActioner) String() string {
	return tea.Prettify(s)
}

func (s GetMeCorpSubmissionResponseBodyDataActioner) GoString() string {
	return s.String()
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetBuName(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.BuName = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetEmail(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.Email = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetEmployeeType(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.EmployeeType = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetEmployeeTypeInformation(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.EmployeeTypeInformation = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetHumanResourceGroupWorkNumber(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.HumanResourceGroupWorkNumber = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetIsSystemAdmin(v bool) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.IsSystemAdmin = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetLevel(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.Level = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetName(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.Name = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetNickName(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.NickName = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetOrderNumber(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.OrderNumber = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetPersonalPhoto(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.PersonalPhoto = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetPersonalPhotoUrl(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.PersonalPhotoUrl = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetPinyinNameAll(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.PinyinNameAll = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetPinyinNickName(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.PinyinNickName = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetState(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.State = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetSuperUserId(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.SuperUserId = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetTbWang(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.TbWang = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataActioner) SetUserId(v string) *GetMeCorpSubmissionResponseBodyDataActioner {
	s.UserId = &v
	return s
}

type GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances struct {
	// activityId
	ActivityId *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
	// activityInstanceStatus
	ActivityInstanceStatus *string `json:"activityInstanceStatus,omitempty" xml:"activityInstanceStatus,omitempty"`
	// activityName
	ActivityName *string `json:"activityName,omitempty" xml:"activityName,omitempty"`
	// activityNameEn
	ActivityNameEn *string `json:"activityNameEn,omitempty" xml:"activityNameEn,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances) String() string {
	return tea.Prettify(s)
}

func (s GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances) GoString() string {
	return s.String()
}

func (s *GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances) SetActivityId(v string) *GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances {
	s.ActivityId = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances) SetActivityInstanceStatus(v string) *GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances {
	s.ActivityInstanceStatus = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances) SetActivityName(v string) *GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances {
	s.ActivityName = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances) SetActivityNameEn(v string) *GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances {
	s.ActivityNameEn = &v
	return s
}

func (s *GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances) SetId(v int64) *GetMeCorpSubmissionResponseBodyDataCurrentActivityInstances {
	s.Id = &v
	return s
}

type GetMeCorpSubmissionResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMeCorpSubmissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMeCorpSubmissionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMeCorpSubmissionResponse) GoString() string {
	return s.String()
}

func (s *GetMeCorpSubmissionResponse) SetHeaders(v map[string]*string) *GetMeCorpSubmissionResponse {
	s.Headers = v
	return s
}

func (s *GetMeCorpSubmissionResponse) SetBody(v *GetMeCorpSubmissionResponseBody) *GetMeCorpSubmissionResponse {
	s.Body = v
	return s
}

type GetNotifyMeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetNotifyMeHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetNotifyMeHeaders) GoString() string {
	return s.String()
}

func (s *GetNotifyMeHeaders) SetCommonHeaders(v map[string]*string) *GetNotifyMeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetNotifyMeHeaders) SetXAcsDingtalkAccessToken(v string) *GetNotifyMeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetNotifyMeRequest struct {
	// 应用标识列表
	AppTypes *string `json:"appTypes,omitempty" xml:"appTypes,omitempty"`
	// 企业ID
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 抄送到达时间开始
	CreateFromTimeGMT *int64 `json:"createFromTimeGMT,omitempty" xml:"createFromTimeGMT,omitempty"`
	// 抄送到达时间结束
	CreateToTimeGMT *int64 `json:"createToTimeGMT,omitempty" xml:"createToTimeGMT,omitempty"`
	// 数据提交时间开始
	InstanceCreateFromTimeGMT *int64 `json:"instanceCreateFromTimeGMT,omitempty" xml:"instanceCreateFromTimeGMT,omitempty"`
	// 数据提交时间结束
	InstanceCreateToTimeGMT *int64 `json:"instanceCreateToTimeGMT,omitempty" xml:"instanceCreateToTimeGMT,omitempty"`
	// 表单中组件数据模糊搜
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 语言环境
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 当前页
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页记录数
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 流程code列表
	ProcessCodes *string `json:"processCodes,omitempty" xml:"processCodes,omitempty"`
	// 验权token
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s GetNotifyMeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNotifyMeRequest) GoString() string {
	return s.String()
}

func (s *GetNotifyMeRequest) SetAppTypes(v string) *GetNotifyMeRequest {
	s.AppTypes = &v
	return s
}

func (s *GetNotifyMeRequest) SetCorpId(v string) *GetNotifyMeRequest {
	s.CorpId = &v
	return s
}

func (s *GetNotifyMeRequest) SetCreateFromTimeGMT(v int64) *GetNotifyMeRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *GetNotifyMeRequest) SetCreateToTimeGMT(v int64) *GetNotifyMeRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *GetNotifyMeRequest) SetInstanceCreateFromTimeGMT(v int64) *GetNotifyMeRequest {
	s.InstanceCreateFromTimeGMT = &v
	return s
}

func (s *GetNotifyMeRequest) SetInstanceCreateToTimeGMT(v int64) *GetNotifyMeRequest {
	s.InstanceCreateToTimeGMT = &v
	return s
}

func (s *GetNotifyMeRequest) SetKeyword(v string) *GetNotifyMeRequest {
	s.Keyword = &v
	return s
}

func (s *GetNotifyMeRequest) SetLanguage(v string) *GetNotifyMeRequest {
	s.Language = &v
	return s
}

func (s *GetNotifyMeRequest) SetPageNumber(v int32) *GetNotifyMeRequest {
	s.PageNumber = &v
	return s
}

func (s *GetNotifyMeRequest) SetPageSize(v int32) *GetNotifyMeRequest {
	s.PageSize = &v
	return s
}

func (s *GetNotifyMeRequest) SetProcessCodes(v string) *GetNotifyMeRequest {
	s.ProcessCodes = &v
	return s
}

func (s *GetNotifyMeRequest) SetToken(v string) *GetNotifyMeRequest {
	s.Token = &v
	return s
}

type GetNotifyMeResponseBody struct {
	// data
	Data []*GetNotifyMeResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 当前第几页
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 总数量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetNotifyMeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNotifyMeResponseBody) GoString() string {
	return s.String()
}

func (s *GetNotifyMeResponseBody) SetData(v []*GetNotifyMeResponseBodyData) *GetNotifyMeResponseBody {
	s.Data = v
	return s
}

func (s *GetNotifyMeResponseBody) SetPageNumber(v int64) *GetNotifyMeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetNotifyMeResponseBody) SetTotalCount(v int64) *GetNotifyMeResponseBody {
	s.TotalCount = &v
	return s
}

type GetNotifyMeResponseBodyData struct {
	// activityId
	ActivityId *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
	// appType
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// corpId
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 创建时间
	CreateTimeGMT *string `json:"createTimeGMT,omitempty" xml:"createTimeGMT,omitempty"`
	// 创建者的userId
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// 流程实例id
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// instStatus
	InstStatus *string `json:"instStatus,omitempty" xml:"instStatus,omitempty"`
	// mobileUrl
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	// 修改时间
	ModifiedTimeGMT *string `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	// processCode
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// title
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// titleEn
	TitleInEnglish *string `json:"titleInEnglish,omitempty" xml:"titleInEnglish,omitempty"`
	// url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetNotifyMeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetNotifyMeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNotifyMeResponseBodyData) SetActivityId(v string) *GetNotifyMeResponseBodyData {
	s.ActivityId = &v
	return s
}

func (s *GetNotifyMeResponseBodyData) SetAppType(v string) *GetNotifyMeResponseBodyData {
	s.AppType = &v
	return s
}

func (s *GetNotifyMeResponseBodyData) SetCorpId(v string) *GetNotifyMeResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *GetNotifyMeResponseBodyData) SetCreateTimeGMT(v string) *GetNotifyMeResponseBodyData {
	s.CreateTimeGMT = &v
	return s
}

func (s *GetNotifyMeResponseBodyData) SetCreatorUserId(v string) *GetNotifyMeResponseBodyData {
	s.CreatorUserId = &v
	return s
}

func (s *GetNotifyMeResponseBodyData) SetFormInstanceId(v string) *GetNotifyMeResponseBodyData {
	s.FormInstanceId = &v
	return s
}

func (s *GetNotifyMeResponseBodyData) SetInstStatus(v string) *GetNotifyMeResponseBodyData {
	s.InstStatus = &v
	return s
}

func (s *GetNotifyMeResponseBodyData) SetMobileUrl(v string) *GetNotifyMeResponseBodyData {
	s.MobileUrl = &v
	return s
}

func (s *GetNotifyMeResponseBodyData) SetModifiedTimeGMT(v string) *GetNotifyMeResponseBodyData {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *GetNotifyMeResponseBodyData) SetProcessCode(v string) *GetNotifyMeResponseBodyData {
	s.ProcessCode = &v
	return s
}

func (s *GetNotifyMeResponseBodyData) SetTitle(v string) *GetNotifyMeResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetNotifyMeResponseBodyData) SetTitleInEnglish(v string) *GetNotifyMeResponseBodyData {
	s.TitleInEnglish = &v
	return s
}

func (s *GetNotifyMeResponseBodyData) SetUrl(v string) *GetNotifyMeResponseBodyData {
	s.Url = &v
	return s
}

type GetNotifyMeResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetNotifyMeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNotifyMeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNotifyMeResponse) GoString() string {
	return s.String()
}

func (s *GetNotifyMeResponse) SetHeaders(v map[string]*string) *GetNotifyMeResponse {
	s.Headers = v
	return s
}

func (s *GetNotifyMeResponse) SetBody(v *GetNotifyMeResponseBody) *GetNotifyMeResponse {
	s.Body = v
	return s
}

type GetOpenUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetOpenUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetOpenUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetOpenUrlHeaders) SetCommonHeaders(v map[string]*string) *GetOpenUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOpenUrlHeaders) SetXAcsDingtalkAccessToken(v string) *GetOpenUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetOpenUrlRequest struct {
	// 宜搭附件地址
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// 语言
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 临时地址多久失效,单位毫秒
	Timeout *int64 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// 钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetOpenUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOpenUrlRequest) GoString() string {
	return s.String()
}

func (s *GetOpenUrlRequest) SetFileUrl(v string) *GetOpenUrlRequest {
	s.FileUrl = &v
	return s
}

func (s *GetOpenUrlRequest) SetLanguage(v string) *GetOpenUrlRequest {
	s.Language = &v
	return s
}

func (s *GetOpenUrlRequest) SetSystemToken(v string) *GetOpenUrlRequest {
	s.SystemToken = &v
	return s
}

func (s *GetOpenUrlRequest) SetTimeout(v int64) *GetOpenUrlRequest {
	s.Timeout = &v
	return s
}

func (s *GetOpenUrlRequest) SetUserId(v string) *GetOpenUrlRequest {
	s.UserId = &v
	return s
}

type GetOpenUrlResponseBody struct {
	// 临时免登地址url
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetOpenUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOpenUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpenUrlResponseBody) SetResult(v string) *GetOpenUrlResponseBody {
	s.Result = &v
	return s
}

type GetOpenUrlResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOpenUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOpenUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOpenUrlResponse) GoString() string {
	return s.String()
}

func (s *GetOpenUrlResponse) SetHeaders(v map[string]*string) *GetOpenUrlResponse {
	s.Headers = v
	return s
}

func (s *GetOpenUrlResponse) SetBody(v *GetOpenUrlResponseBody) *GetOpenUrlResponse {
	s.Body = v
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
	// 语言
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 流程实例ID
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// 应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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

func (s *GetOperationRecordsRequest) SetLanguage(v string) *GetOperationRecordsRequest {
	s.Language = &v
	return s
}

func (s *GetOperationRecordsRequest) SetProcessInstanceId(v string) *GetOperationRecordsRequest {
	s.ProcessInstanceId = &v
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
	// action
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// actionExt
	ActionExit *string `json:"actionExit,omitempty" xml:"actionExit,omitempty"`
	// activeTime
	ActiveTimeGMT *string `json:"activeTimeGMT,omitempty" xml:"activeTimeGMT,omitempty"`
	// activityId
	ActivityId *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
	// id
	DataId *int64 `json:"dataId,omitempty" xml:"dataId,omitempty"`
	// digitalSign
	DigitalSign *string `json:"digitalSign,omitempty" xml:"digitalSign,omitempty"`
	// files
	Files *string `json:"files,omitempty" xml:"files,omitempty"`
	// operateTime
	OperateTimeGMT *string `json:"operateTimeGMT,omitempty" xml:"operateTimeGMT,omitempty"`
	// operateType
	OperateType *string `json:"operateType,omitempty" xml:"operateType,omitempty"`
	// operatorDisplayName
	OperatorDisplayName *string `json:"operatorDisplayName,omitempty" xml:"operatorDisplayName,omitempty"`
	// operatorName
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	// operatorNick
	OperatorNickName *string `json:"operatorNickName,omitempty" xml:"operatorNickName,omitempty"`
	// operatorPhotoUrl
	OperatorPhotoUrl *string `json:"operatorPhotoUrl,omitempty" xml:"operatorPhotoUrl,omitempty"`
	// operatorStatus
	OperatorStatus *string `json:"operatorStatus,omitempty" xml:"operatorStatus,omitempty"`
	// operator
	OperatorUserId *string `json:"operatorUserId,omitempty" xml:"operatorUserId,omitempty"`
	// processInstanceId
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// remark
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// showName
	ShowName *string `json:"showName,omitempty" xml:"showName,omitempty"`
	// size
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// taskExecuteType
	TaskExecuteType *string `json:"taskExecuteType,omitempty" xml:"taskExecuteType,omitempty"`
	// taskHoldTime
	TaskHoldTimeGMT *int64 `json:"taskHoldTimeGMT,omitempty" xml:"taskHoldTimeGMT,omitempty"`
	// taskId
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// taskType
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// type
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetOperationRecordsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetOperationRecordsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetOperationRecordsResponseBodyResult) SetAction(v string) *GetOperationRecordsResponseBodyResult {
	s.Action = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetActionExit(v string) *GetOperationRecordsResponseBodyResult {
	s.ActionExit = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetActiveTimeGMT(v string) *GetOperationRecordsResponseBodyResult {
	s.ActiveTimeGMT = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetActivityId(v string) *GetOperationRecordsResponseBodyResult {
	s.ActivityId = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetDataId(v int64) *GetOperationRecordsResponseBodyResult {
	s.DataId = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetDigitalSign(v string) *GetOperationRecordsResponseBodyResult {
	s.DigitalSign = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetFiles(v string) *GetOperationRecordsResponseBodyResult {
	s.Files = &v
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

func (s *GetOperationRecordsResponseBodyResult) SetOperatorDisplayName(v string) *GetOperationRecordsResponseBodyResult {
	s.OperatorDisplayName = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetOperatorName(v string) *GetOperationRecordsResponseBodyResult {
	s.OperatorName = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetOperatorNickName(v string) *GetOperationRecordsResponseBodyResult {
	s.OperatorNickName = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetOperatorPhotoUrl(v string) *GetOperationRecordsResponseBodyResult {
	s.OperatorPhotoUrl = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetOperatorStatus(v string) *GetOperationRecordsResponseBodyResult {
	s.OperatorStatus = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetOperatorUserId(v string) *GetOperationRecordsResponseBodyResult {
	s.OperatorUserId = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetProcessInstanceId(v string) *GetOperationRecordsResponseBodyResult {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetRemark(v string) *GetOperationRecordsResponseBodyResult {
	s.Remark = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetShowName(v string) *GetOperationRecordsResponseBodyResult {
	s.ShowName = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetSize(v int32) *GetOperationRecordsResponseBodyResult {
	s.Size = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetTaskExecuteType(v string) *GetOperationRecordsResponseBodyResult {
	s.TaskExecuteType = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetTaskHoldTimeGMT(v int64) *GetOperationRecordsResponseBodyResult {
	s.TaskHoldTimeGMT = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetTaskId(v string) *GetOperationRecordsResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetTaskType(v string) *GetOperationRecordsResponseBodyResult {
	s.TaskType = &v
	return s
}

func (s *GetOperationRecordsResponseBodyResult) SetType(v string) *GetOperationRecordsResponseBodyResult {
	s.Type = &v
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
	AccessKey  *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	CallerUid  *string `json:"callerUid,omitempty" xml:"callerUid,omitempty"`
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s GetPlatformResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPlatformResourceRequest) GoString() string {
	return s.String()
}

func (s *GetPlatformResourceRequest) SetAccessKey(v string) *GetPlatformResourceRequest {
	s.AccessKey = &v
	return s
}

func (s *GetPlatformResourceRequest) SetCallerUid(v string) *GetPlatformResourceRequest {
	s.CallerUid = &v
	return s
}

func (s *GetPlatformResourceRequest) SetInstanceId(v string) *GetPlatformResourceRequest {
	s.InstanceId = &v
	return s
}

type GetPlatformResourceResponseBody struct {
	// accountTotalAmount
	AccountTotalAmount *int32 `json:"accountTotalAmount,omitempty" xml:"accountTotalAmount,omitempty"`
	// accountUsageAmount
	AccountUsageAmount *int32 `json:"accountUsageAmount,omitempty" xml:"accountUsageAmount,omitempty"`
	// appTotalAmount
	AppTotalAmount *int32 `json:"appTotalAmount,omitempty" xml:"appTotalAmount,omitempty"`
	// attachmentTotalAmount
	AttachmentTotalAmount *int64 `json:"attachmentTotalAmount,omitempty" xml:"attachmentTotalAmount,omitempty"`
	// attachmentUsageAmount
	AttachmentUsageAmount *int64 `json:"attachmentUsageAmount,omitempty" xml:"attachmentUsageAmount,omitempty"`
	// instanceTotalAmount
	InstanceTotalAmount *int64 `json:"instanceTotalAmount,omitempty" xml:"instanceTotalAmount,omitempty"`
	// instanceUsageAmount
	InstanceUsageAmount *int64 `json:"instanceUsageAmount,omitempty" xml:"instanceUsageAmount,omitempty"`
	// pluginUsageAmount
	PluginUsageAmount *int64 `json:"pluginUsageAmount,omitempty" xml:"pluginUsageAmount,omitempty"`
}

func (s GetPlatformResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPlatformResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetPlatformResourceResponseBody) SetAccountTotalAmount(v int32) *GetPlatformResourceResponseBody {
	s.AccountTotalAmount = &v
	return s
}

func (s *GetPlatformResourceResponseBody) SetAccountUsageAmount(v int32) *GetPlatformResourceResponseBody {
	s.AccountUsageAmount = &v
	return s
}

func (s *GetPlatformResourceResponseBody) SetAppTotalAmount(v int32) *GetPlatformResourceResponseBody {
	s.AppTotalAmount = &v
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

func (s *GetPlatformResourceResponseBody) SetInstanceTotalAmount(v int64) *GetPlatformResourceResponseBody {
	s.InstanceTotalAmount = &v
	return s
}

func (s *GetPlatformResourceResponseBody) SetInstanceUsageAmount(v int64) *GetPlatformResourceResponseBody {
	s.InstanceUsageAmount = &v
	return s
}

func (s *GetPlatformResourceResponseBody) SetPluginUsageAmount(v int64) *GetPlatformResourceResponseBody {
	s.PluginUsageAmount = &v
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

type GetPrintAppInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPrintAppInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPrintAppInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetPrintAppInfoHeaders) SetCommonHeaders(v map[string]*string) *GetPrintAppInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPrintAppInfoHeaders) SetXAcsDingtalkAccessToken(v string) *GetPrintAppInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPrintAppInfoRequest struct {
	// 搜索关键字
	NameLike *string `json:"nameLike,omitempty" xml:"nameLike,omitempty"`
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetPrintAppInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPrintAppInfoRequest) GoString() string {
	return s.String()
}

func (s *GetPrintAppInfoRequest) SetNameLike(v string) *GetPrintAppInfoRequest {
	s.NameLike = &v
	return s
}

func (s *GetPrintAppInfoRequest) SetUserId(v string) *GetPrintAppInfoRequest {
	s.UserId = &v
	return s
}

type GetPrintAppInfoResponseBody struct {
	Result []*GetPrintAppInfoResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetPrintAppInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPrintAppInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetPrintAppInfoResponseBody) SetResult(v []*GetPrintAppInfoResponseBodyResult) *GetPrintAppInfoResponseBody {
	s.Result = v
	return s
}

type GetPrintAppInfoResponseBodyResult struct {
	// 应用名称
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// appType
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// formInfoList
	FormInfoList []*GetPrintAppInfoResponseBodyResultFormInfoList `json:"formInfoList,omitempty" xml:"formInfoList,omitempty" type:"Repeated"`
	// 图标链接
	IconUrl *string `json:"iconUrl,omitempty" xml:"iconUrl,omitempty"`
}

func (s GetPrintAppInfoResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetPrintAppInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetPrintAppInfoResponseBodyResult) SetAppName(v string) *GetPrintAppInfoResponseBodyResult {
	s.AppName = &v
	return s
}

func (s *GetPrintAppInfoResponseBodyResult) SetAppType(v string) *GetPrintAppInfoResponseBodyResult {
	s.AppType = &v
	return s
}

func (s *GetPrintAppInfoResponseBodyResult) SetFormInfoList(v []*GetPrintAppInfoResponseBodyResultFormInfoList) *GetPrintAppInfoResponseBodyResult {
	s.FormInfoList = v
	return s
}

func (s *GetPrintAppInfoResponseBodyResult) SetIconUrl(v string) *GetPrintAppInfoResponseBodyResult {
	s.IconUrl = &v
	return s
}

type GetPrintAppInfoResponseBodyResultFormInfoList struct {
	// formName
	FormName *string `json:"formName,omitempty" xml:"formName,omitempty"`
	// formUuid
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
}

func (s GetPrintAppInfoResponseBodyResultFormInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetPrintAppInfoResponseBodyResultFormInfoList) GoString() string {
	return s.String()
}

func (s *GetPrintAppInfoResponseBodyResultFormInfoList) SetFormName(v string) *GetPrintAppInfoResponseBodyResultFormInfoList {
	s.FormName = &v
	return s
}

func (s *GetPrintAppInfoResponseBodyResultFormInfoList) SetFormUuid(v string) *GetPrintAppInfoResponseBodyResultFormInfoList {
	s.FormUuid = &v
	return s
}

type GetPrintAppInfoResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPrintAppInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPrintAppInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPrintAppInfoResponse) GoString() string {
	return s.String()
}

func (s *GetPrintAppInfoResponse) SetHeaders(v map[string]*string) *GetPrintAppInfoResponse {
	s.Headers = v
	return s
}

func (s *GetPrintAppInfoResponse) SetBody(v *GetPrintAppInfoResponseBody) *GetPrintAppInfoResponse {
	s.Body = v
	return s
}

type GetPrintDictionaryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetPrintDictionaryHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetPrintDictionaryHeaders) GoString() string {
	return s.String()
}

func (s *GetPrintDictionaryHeaders) SetCommonHeaders(v map[string]*string) *GetPrintDictionaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPrintDictionaryHeaders) SetXAcsDingtalkAccessToken(v string) *GetPrintDictionaryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetPrintDictionaryRequest struct {
	// 应用代码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 表单id
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 版本
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetPrintDictionaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPrintDictionaryRequest) GoString() string {
	return s.String()
}

func (s *GetPrintDictionaryRequest) SetAppType(v string) *GetPrintDictionaryRequest {
	s.AppType = &v
	return s
}

func (s *GetPrintDictionaryRequest) SetFormUuid(v string) *GetPrintDictionaryRequest {
	s.FormUuid = &v
	return s
}

func (s *GetPrintDictionaryRequest) SetUserId(v string) *GetPrintDictionaryRequest {
	s.UserId = &v
	return s
}

func (s *GetPrintDictionaryRequest) SetVersion(v int64) *GetPrintDictionaryRequest {
	s.Version = &v
	return s
}

type GetPrintDictionaryResponseBody struct {
	// Id of the request
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetPrintDictionaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPrintDictionaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetPrintDictionaryResponseBody) SetResult(v string) *GetPrintDictionaryResponseBody {
	s.Result = &v
	return s
}

type GetPrintDictionaryResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPrintDictionaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPrintDictionaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPrintDictionaryResponse) GoString() string {
	return s.String()
}

func (s *GetPrintDictionaryResponse) SetHeaders(v map[string]*string) *GetPrintDictionaryResponse {
	s.Headers = v
	return s
}

func (s *GetPrintDictionaryResponse) SetBody(v *GetPrintDictionaryResponseBody) *GetPrintDictionaryResponse {
	s.Body = v
	return s
}

type GetProcessDefinitionHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetProcessDefinitionHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetProcessDefinitionHeaders) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionHeaders) SetCommonHeaders(v map[string]*string) *GetProcessDefinitionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetProcessDefinitionHeaders) SetXAcsDingtalkAccessToken(v string) *GetProcessDefinitionHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetProcessDefinitionRequest struct {
	AppType     *string `json:"appType,omitempty" xml:"appType,omitempty"`
	CorpId      *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	GroupId     *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	Language    *string `json:"language,omitempty" xml:"language,omitempty"`
	NameSpace   *string `json:"nameSpace,omitempty" xml:"nameSpace,omitempty"`
	OrderNumber *string `json:"orderNumber,omitempty" xml:"orderNumber,omitempty"`
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	SystemType  *string `json:"systemType,omitempty" xml:"systemType,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetProcessDefinitionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProcessDefinitionRequest) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionRequest) SetAppType(v string) *GetProcessDefinitionRequest {
	s.AppType = &v
	return s
}

func (s *GetProcessDefinitionRequest) SetCorpId(v string) *GetProcessDefinitionRequest {
	s.CorpId = &v
	return s
}

func (s *GetProcessDefinitionRequest) SetGroupId(v string) *GetProcessDefinitionRequest {
	s.GroupId = &v
	return s
}

func (s *GetProcessDefinitionRequest) SetLanguage(v string) *GetProcessDefinitionRequest {
	s.Language = &v
	return s
}

func (s *GetProcessDefinitionRequest) SetNameSpace(v string) *GetProcessDefinitionRequest {
	s.NameSpace = &v
	return s
}

func (s *GetProcessDefinitionRequest) SetOrderNumber(v string) *GetProcessDefinitionRequest {
	s.OrderNumber = &v
	return s
}

func (s *GetProcessDefinitionRequest) SetSystemToken(v string) *GetProcessDefinitionRequest {
	s.SystemToken = &v
	return s
}

func (s *GetProcessDefinitionRequest) SetSystemType(v string) *GetProcessDefinitionRequest {
	s.SystemType = &v
	return s
}

func (s *GetProcessDefinitionRequest) SetUserId(v string) *GetProcessDefinitionRequest {
	s.UserId = &v
	return s
}

type GetProcessDefinitionResponseBody struct {
	// formUuid
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// originator
	Originator *GetProcessDefinitionResponseBodyOriginator `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	// outResult
	OutResult *string `json:"outResult,omitempty" xml:"outResult,omitempty"`
	// owners
	Owners []*GetProcessDefinitionResponseBodyOwners `json:"owners,omitempty" xml:"owners,omitempty" type:"Repeated"`
	// processId
	ProcessId *string `json:"processId,omitempty" xml:"processId,omitempty"`
	// processInstanceId
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// status
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// tasks
	Tasks []*GetProcessDefinitionResponseBodyTasks `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Repeated"`
	// title
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// variables
	Variables map[string]interface{} `json:"variables,omitempty" xml:"variables,omitempty"`
}

func (s GetProcessDefinitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProcessDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionResponseBody) SetFormUuid(v string) *GetProcessDefinitionResponseBody {
	s.FormUuid = &v
	return s
}

func (s *GetProcessDefinitionResponseBody) SetOriginator(v *GetProcessDefinitionResponseBodyOriginator) *GetProcessDefinitionResponseBody {
	s.Originator = v
	return s
}

func (s *GetProcessDefinitionResponseBody) SetOutResult(v string) *GetProcessDefinitionResponseBody {
	s.OutResult = &v
	return s
}

func (s *GetProcessDefinitionResponseBody) SetOwners(v []*GetProcessDefinitionResponseBodyOwners) *GetProcessDefinitionResponseBody {
	s.Owners = v
	return s
}

func (s *GetProcessDefinitionResponseBody) SetProcessId(v string) *GetProcessDefinitionResponseBody {
	s.ProcessId = &v
	return s
}

func (s *GetProcessDefinitionResponseBody) SetProcessInstanceId(v string) *GetProcessDefinitionResponseBody {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetProcessDefinitionResponseBody) SetStatus(v string) *GetProcessDefinitionResponseBody {
	s.Status = &v
	return s
}

func (s *GetProcessDefinitionResponseBody) SetTasks(v []*GetProcessDefinitionResponseBodyTasks) *GetProcessDefinitionResponseBody {
	s.Tasks = v
	return s
}

func (s *GetProcessDefinitionResponseBody) SetTitle(v string) *GetProcessDefinitionResponseBody {
	s.Title = &v
	return s
}

func (s *GetProcessDefinitionResponseBody) SetVariables(v map[string]interface{}) *GetProcessDefinitionResponseBody {
	s.Variables = v
	return s
}

type GetProcessDefinitionResponseBodyOriginator struct {
	// departmentDescription
	DepartmentDescription *string `json:"departmentDescription,omitempty" xml:"departmentDescription,omitempty"`
	// displayEnName
	DisplayEnName *string `json:"displayEnName,omitempty" xml:"displayEnName,omitempty"`
	// displayName
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// masterDataDepartments
	MasterDataDepartments []*GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments `json:"masterDataDepartments,omitempty" xml:"masterDataDepartments,omitempty" type:"Repeated"`
	// orderNum
	OrderNumber *string `json:"orderNumber,omitempty" xml:"orderNumber,omitempty"`
	// personalPhoto
	PersonalPhoto *string `json:"personalPhoto,omitempty" xml:"personalPhoto,omitempty"`
	// status
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// tbWang
	TbWang *string `json:"tbWang,omitempty" xml:"tbWang,omitempty"`
	// userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// userInfo
	UserInfo *string `json:"userInfo,omitempty" xml:"userInfo,omitempty"`
}

func (s GetProcessDefinitionResponseBodyOriginator) String() string {
	return tea.Prettify(s)
}

func (s GetProcessDefinitionResponseBodyOriginator) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionResponseBodyOriginator) SetDepartmentDescription(v string) *GetProcessDefinitionResponseBodyOriginator {
	s.DepartmentDescription = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginator) SetDisplayEnName(v string) *GetProcessDefinitionResponseBodyOriginator {
	s.DisplayEnName = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginator) SetDisplayName(v string) *GetProcessDefinitionResponseBodyOriginator {
	s.DisplayName = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginator) SetMasterDataDepartments(v []*GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) *GetProcessDefinitionResponseBodyOriginator {
	s.MasterDataDepartments = v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginator) SetOrderNumber(v string) *GetProcessDefinitionResponseBodyOriginator {
	s.OrderNumber = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginator) SetPersonalPhoto(v string) *GetProcessDefinitionResponseBodyOriginator {
	s.PersonalPhoto = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginator) SetStatus(v string) *GetProcessDefinitionResponseBodyOriginator {
	s.Status = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginator) SetTbWang(v string) *GetProcessDefinitionResponseBodyOriginator {
	s.TbWang = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginator) SetUserId(v string) *GetProcessDefinitionResponseBodyOriginator {
	s.UserId = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginator) SetUserInfo(v string) *GetProcessDefinitionResponseBodyOriginator {
	s.UserInfo = &v
	return s
}

type GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments struct {
	// deptName
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// deptNameEn
	DeptNameInEnglish *string `json:"deptNameInEnglish,omitempty" xml:"deptNameInEnglish,omitempty"`
	// deptNo
	DeptNo *string `json:"deptNo,omitempty" xml:"deptNo,omitempty"`
	// deptPath
	DeptPath *string `json:"deptPath,omitempty" xml:"deptPath,omitempty"`
	// humanSourceGroupOrderNum
	HumanSourceGroupOrderNumber *string `json:"humanSourceGroupOrderNumber,omitempty" xml:"humanSourceGroupOrderNumber,omitempty"`
	// humanSourceGroupWorkNo
	HumanSourceGroupWorkNo *string `json:"humanSourceGroupWorkNo,omitempty" xml:"humanSourceGroupWorkNo,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// masterWorkNo
	MasterWorkNo *string `json:"masterWorkNo,omitempty" xml:"masterWorkNo,omitempty"`
}

func (s GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) String() string {
	return tea.Prettify(s)
}

func (s GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) SetDeptName(v string) *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments {
	s.DeptName = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) SetDeptNameInEnglish(v string) *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments {
	s.DeptNameInEnglish = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) SetDeptNo(v string) *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments {
	s.DeptNo = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) SetDeptPath(v string) *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments {
	s.DeptPath = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) SetHumanSourceGroupOrderNumber(v string) *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments {
	s.HumanSourceGroupOrderNumber = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) SetHumanSourceGroupWorkNo(v string) *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments {
	s.HumanSourceGroupWorkNo = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) SetId(v int64) *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments {
	s.Id = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments) SetMasterWorkNo(v string) *GetProcessDefinitionResponseBodyOriginatorMasterDataDepartments {
	s.MasterWorkNo = &v
	return s
}

type GetProcessDefinitionResponseBodyOwners struct {
	// departmentDescription
	DepartmentDescription *string `json:"departmentDescription,omitempty" xml:"departmentDescription,omitempty"`
	// displayEnName
	DisplayEnName *string `json:"displayEnName,omitempty" xml:"displayEnName,omitempty"`
	// displayName
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// masterDataDepartments
	MasterDataDepartments []*GetProcessDefinitionResponseBodyOwnersMasterDataDepartments `json:"masterDataDepartments,omitempty" xml:"masterDataDepartments,omitempty" type:"Repeated"`
	// orderNum
	OrderNumber *string `json:"orderNumber,omitempty" xml:"orderNumber,omitempty"`
	// personalPhoto
	PersonalPhoto *string `json:"personalPhoto,omitempty" xml:"personalPhoto,omitempty"`
	// status
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// tbWang
	TbWang *string `json:"tbWang,omitempty" xml:"tbWang,omitempty"`
	// userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// userInfo
	UserInfo *string `json:"userInfo,omitempty" xml:"userInfo,omitempty"`
}

func (s GetProcessDefinitionResponseBodyOwners) String() string {
	return tea.Prettify(s)
}

func (s GetProcessDefinitionResponseBodyOwners) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionResponseBodyOwners) SetDepartmentDescription(v string) *GetProcessDefinitionResponseBodyOwners {
	s.DepartmentDescription = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwners) SetDisplayEnName(v string) *GetProcessDefinitionResponseBodyOwners {
	s.DisplayEnName = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwners) SetDisplayName(v string) *GetProcessDefinitionResponseBodyOwners {
	s.DisplayName = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwners) SetMasterDataDepartments(v []*GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) *GetProcessDefinitionResponseBodyOwners {
	s.MasterDataDepartments = v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwners) SetOrderNumber(v string) *GetProcessDefinitionResponseBodyOwners {
	s.OrderNumber = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwners) SetPersonalPhoto(v string) *GetProcessDefinitionResponseBodyOwners {
	s.PersonalPhoto = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwners) SetStatus(v string) *GetProcessDefinitionResponseBodyOwners {
	s.Status = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwners) SetTbWang(v string) *GetProcessDefinitionResponseBodyOwners {
	s.TbWang = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwners) SetUserId(v string) *GetProcessDefinitionResponseBodyOwners {
	s.UserId = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwners) SetUserInfo(v string) *GetProcessDefinitionResponseBodyOwners {
	s.UserInfo = &v
	return s
}

type GetProcessDefinitionResponseBodyOwnersMasterDataDepartments struct {
	// deptName
	DeptName *string `json:"deptName,omitempty" xml:"deptName,omitempty"`
	// deptNameEn
	DeptNameInEnglish *string `json:"deptNameInEnglish,omitempty" xml:"deptNameInEnglish,omitempty"`
	// deptNo
	DeptNo *string `json:"deptNo,omitempty" xml:"deptNo,omitempty"`
	// deptPath
	DeptPath *string `json:"deptPath,omitempty" xml:"deptPath,omitempty"`
	// humanSourceGroupOrderNum
	HumanSourceGroupOrderNumber *string `json:"humanSourceGroupOrderNumber,omitempty" xml:"humanSourceGroupOrderNumber,omitempty"`
	// humanSourceGroupWorkNo
	HumanSourceGroupWorkNo *string `json:"humanSourceGroupWorkNo,omitempty" xml:"humanSourceGroupWorkNo,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// masterWorkNo
	MasterWorkNo *string `json:"masterWorkNo,omitempty" xml:"masterWorkNo,omitempty"`
}

func (s GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) String() string {
	return tea.Prettify(s)
}

func (s GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) SetDeptName(v string) *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments {
	s.DeptName = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) SetDeptNameInEnglish(v string) *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments {
	s.DeptNameInEnglish = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) SetDeptNo(v string) *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments {
	s.DeptNo = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) SetDeptPath(v string) *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments {
	s.DeptPath = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) SetHumanSourceGroupOrderNumber(v string) *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments {
	s.HumanSourceGroupOrderNumber = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) SetHumanSourceGroupWorkNo(v string) *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments {
	s.HumanSourceGroupWorkNo = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) SetId(v int64) *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments {
	s.Id = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments) SetMasterWorkNo(v string) *GetProcessDefinitionResponseBodyOwnersMasterDataDepartments {
	s.MasterWorkNo = &v
	return s
}

type GetProcessDefinitionResponseBodyTasks struct {
	// actioner
	ActionerId *string `json:"actionerId,omitempty" xml:"actionerId,omitempty"`
	// activity
	Activity *GetProcessDefinitionResponseBodyTasksActivity `json:"activity,omitempty" xml:"activity,omitempty" type:"Struct"`
	// status
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// taskId
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetProcessDefinitionResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s GetProcessDefinitionResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionResponseBodyTasks) SetActionerId(v string) *GetProcessDefinitionResponseBodyTasks {
	s.ActionerId = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyTasks) SetActivity(v *GetProcessDefinitionResponseBodyTasksActivity) *GetProcessDefinitionResponseBodyTasks {
	s.Activity = v
	return s
}

func (s *GetProcessDefinitionResponseBodyTasks) SetStatus(v string) *GetProcessDefinitionResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyTasks) SetTaskId(v int64) *GetProcessDefinitionResponseBodyTasks {
	s.TaskId = &v
	return s
}

type GetProcessDefinitionResponseBodyTasksActivity struct {
	// activityId
	ActivityId *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
	// activityInstanceStatus
	ActivityInstanceStatus *string `json:"activityInstanceStatus,omitempty" xml:"activityInstanceStatus,omitempty"`
	// activityName
	ActivityName *string `json:"activityName,omitempty" xml:"activityName,omitempty"`
	// activityNameEn
	ActivityNameInEnglish *string `json:"activityNameInEnglish,omitempty" xml:"activityNameInEnglish,omitempty"`
	// id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s GetProcessDefinitionResponseBodyTasksActivity) String() string {
	return tea.Prettify(s)
}

func (s GetProcessDefinitionResponseBodyTasksActivity) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionResponseBodyTasksActivity) SetActivityId(v string) *GetProcessDefinitionResponseBodyTasksActivity {
	s.ActivityId = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyTasksActivity) SetActivityInstanceStatus(v string) *GetProcessDefinitionResponseBodyTasksActivity {
	s.ActivityInstanceStatus = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyTasksActivity) SetActivityName(v string) *GetProcessDefinitionResponseBodyTasksActivity {
	s.ActivityName = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyTasksActivity) SetActivityNameInEnglish(v string) *GetProcessDefinitionResponseBodyTasksActivity {
	s.ActivityNameInEnglish = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyTasksActivity) SetId(v int64) *GetProcessDefinitionResponseBodyTasksActivity {
	s.Id = &v
	return s
}

type GetProcessDefinitionResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProcessDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProcessDefinitionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProcessDefinitionResponse) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionResponse) SetHeaders(v map[string]*string) *GetProcessDefinitionResponse {
	s.Headers = v
	return s
}

func (s *GetProcessDefinitionResponse) SetBody(v *GetProcessDefinitionResponseBody) *GetProcessDefinitionResponse {
	s.Body = v
	return s
}

type GetRunningTaskListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetRunningTaskListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetRunningTaskListHeaders) GoString() string {
	return s.String()
}

func (s *GetRunningTaskListHeaders) SetCommonHeaders(v map[string]*string) *GetRunningTaskListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRunningTaskListHeaders) SetXAcsDingtalkAccessToken(v string) *GetRunningTaskListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetRunningTaskListRequest struct {
	// appType
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 流程实例id列表, 逗号分隔
	ProcessInstanceIdList *string `json:"processInstanceIdList,omitempty" xml:"processInstanceIdList,omitempty"`
	// systemToken
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 用户所属的企业id
	UserCorpId *string `json:"userCorpId,omitempty" xml:"userCorpId,omitempty"`
	// userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetRunningTaskListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRunningTaskListRequest) GoString() string {
	return s.String()
}

func (s *GetRunningTaskListRequest) SetAppType(v string) *GetRunningTaskListRequest {
	s.AppType = &v
	return s
}

func (s *GetRunningTaskListRequest) SetProcessInstanceIdList(v string) *GetRunningTaskListRequest {
	s.ProcessInstanceIdList = &v
	return s
}

func (s *GetRunningTaskListRequest) SetSystemToken(v string) *GetRunningTaskListRequest {
	s.SystemToken = &v
	return s
}

func (s *GetRunningTaskListRequest) SetUserCorpId(v string) *GetRunningTaskListRequest {
	s.UserCorpId = &v
	return s
}

func (s *GetRunningTaskListRequest) SetUserId(v string) *GetRunningTaskListRequest {
	s.UserId = &v
	return s
}

type GetRunningTaskListResponseBody struct {
	Result []*GetRunningTaskListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetRunningTaskListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRunningTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *GetRunningTaskListResponseBody) SetResult(v []*GetRunningTaskListResponseBodyResult) *GetRunningTaskListResponseBody {
	s.Result = v
	return s
}

type GetRunningTaskListResponseBodyResult struct {
	// 激活时间
	ActiveTimeGMT *string `json:"activeTimeGMT,omitempty" xml:"activeTimeGMT,omitempty"`
	// 实际执行人id
	ActualActionExecutorId *string `json:"actualActionExecutorId,omitempty" xml:"actualActionExecutorId,omitempty"`
	// appType
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 创建时间
	CreateTimeGMT *string `json:"createTimeGMT,omitempty" xml:"createTimeGMT,omitempty"`
	// 结束时间
	FinishTimeGMT *string `json:"finishTimeGMT,omitempty" xml:"finishTimeGMT,omitempty"`
	// originatorEmail
	OriginatorEmail *string `json:"originatorEmail,omitempty" xml:"originatorEmail,omitempty"`
	// originatorId
	OriginatorId *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	// originatorName
	OriginatorName *string `json:"originatorName,omitempty" xml:"originatorName,omitempty"`
	// originatorNameEn
	OriginatorNameInEnglish *string `json:"originatorNameInEnglish,omitempty" xml:"originatorNameInEnglish,omitempty"`
	// originatorNickName
	OriginatorNickName *string `json:"originatorNickName,omitempty" xml:"originatorNickName,omitempty"`
	// originatorNickNameEn
	OriginatorNickNameInEnglish *string `json:"originatorNickNameInEnglish,omitempty" xml:"originatorNickNameInEnglish,omitempty"`
	// originatorPhoto
	OriginatorPhoto *string `json:"originatorPhoto,omitempty" xml:"originatorPhoto,omitempty"`
	// outResult
	OutResult *string `json:"outResult,omitempty" xml:"outResult,omitempty"`
	// outResultName
	OutResultName *string `json:"outResultName,omitempty" xml:"outResultName,omitempty"`
	// processInstanceId
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// 状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 任务id
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// 任务类型
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// 标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 标题英文
	TitleInEnglish *string `json:"titleInEnglish,omitempty" xml:"titleInEnglish,omitempty"`
}

func (s GetRunningTaskListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetRunningTaskListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRunningTaskListResponseBodyResult) SetActiveTimeGMT(v string) *GetRunningTaskListResponseBodyResult {
	s.ActiveTimeGMT = &v
	return s
}

func (s *GetRunningTaskListResponseBodyResult) SetActualActionExecutorId(v string) *GetRunningTaskListResponseBodyResult {
	s.ActualActionExecutorId = &v
	return s
}

func (s *GetRunningTaskListResponseBodyResult) SetAppType(v string) *GetRunningTaskListResponseBodyResult {
	s.AppType = &v
	return s
}

func (s *GetRunningTaskListResponseBodyResult) SetCreateTimeGMT(v string) *GetRunningTaskListResponseBodyResult {
	s.CreateTimeGMT = &v
	return s
}

func (s *GetRunningTaskListResponseBodyResult) SetFinishTimeGMT(v string) *GetRunningTaskListResponseBodyResult {
	s.FinishTimeGMT = &v
	return s
}

func (s *GetRunningTaskListResponseBodyResult) SetOriginatorEmail(v string) *GetRunningTaskListResponseBodyResult {
	s.OriginatorEmail = &v
	return s
}

func (s *GetRunningTaskListResponseBodyResult) SetOriginatorId(v string) *GetRunningTaskListResponseBodyResult {
	s.OriginatorId = &v
	return s
}

func (s *GetRunningTaskListResponseBodyResult) SetOriginatorName(v string) *GetRunningTaskListResponseBodyResult {
	s.OriginatorName = &v
	return s
}

func (s *GetRunningTaskListResponseBodyResult) SetOriginatorNameInEnglish(v string) *GetRunningTaskListResponseBodyResult {
	s.OriginatorNameInEnglish = &v
	return s
}

func (s *GetRunningTaskListResponseBodyResult) SetOriginatorNickName(v string) *GetRunningTaskListResponseBodyResult {
	s.OriginatorNickName = &v
	return s
}

func (s *GetRunningTaskListResponseBodyResult) SetOriginatorNickNameInEnglish(v string) *GetRunningTaskListResponseBodyResult {
	s.OriginatorNickNameInEnglish = &v
	return s
}

func (s *GetRunningTaskListResponseBodyResult) SetOriginatorPhoto(v string) *GetRunningTaskListResponseBodyResult {
	s.OriginatorPhoto = &v
	return s
}

func (s *GetRunningTaskListResponseBodyResult) SetOutResult(v string) *GetRunningTaskListResponseBodyResult {
	s.OutResult = &v
	return s
}

func (s *GetRunningTaskListResponseBodyResult) SetOutResultName(v string) *GetRunningTaskListResponseBodyResult {
	s.OutResultName = &v
	return s
}

func (s *GetRunningTaskListResponseBodyResult) SetProcessInstanceId(v string) *GetRunningTaskListResponseBodyResult {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetRunningTaskListResponseBodyResult) SetStatus(v string) *GetRunningTaskListResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetRunningTaskListResponseBodyResult) SetTaskId(v string) *GetRunningTaskListResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *GetRunningTaskListResponseBodyResult) SetTaskType(v string) *GetRunningTaskListResponseBodyResult {
	s.TaskType = &v
	return s
}

func (s *GetRunningTaskListResponseBodyResult) SetTitle(v string) *GetRunningTaskListResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetRunningTaskListResponseBodyResult) SetTitleInEnglish(v string) *GetRunningTaskListResponseBodyResult {
	s.TitleInEnglish = &v
	return s
}

type GetRunningTaskListResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRunningTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRunningTaskListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRunningTaskListResponse) GoString() string {
	return s.String()
}

func (s *GetRunningTaskListResponse) SetHeaders(v map[string]*string) *GetRunningTaskListResponse {
	s.Headers = v
	return s
}

func (s *GetRunningTaskListResponse) SetBody(v *GetRunningTaskListResponseBody) *GetRunningTaskListResponse {
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
	AppType           *string `json:"appType,omitempty" xml:"appType,omitempty"`
	Language          *string `json:"language,omitempty" xml:"language,omitempty"`
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	SystemToken       *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	UserId            *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetRunningTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRunningTasksRequest) GoString() string {
	return s.String()
}

func (s *GetRunningTasksRequest) SetAppType(v string) *GetRunningTasksRequest {
	s.AppType = &v
	return s
}

func (s *GetRunningTasksRequest) SetLanguage(v string) *GetRunningTasksRequest {
	s.Language = &v
	return s
}

func (s *GetRunningTasksRequest) SetProcessInstanceId(v string) *GetRunningTasksRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetRunningTasksRequest) SetSystemToken(v string) *GetRunningTasksRequest {
	s.SystemToken = &v
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
	// activeTime
	ActiveTimeGMT *string `json:"activeTimeGMT,omitempty" xml:"activeTimeGMT,omitempty"`
	// activityId
	ActivityId *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
	// actualActionerId
	ActualActionerId *string `json:"actualActionerId,omitempty" xml:"actualActionerId,omitempty"`
	// createTime
	CreateTimeGMT *string `json:"createTimeGMT,omitempty" xml:"createTimeGMT,omitempty"`
	// finishTime
	FinishTimeGMT *string `json:"finishTimeGMT,omitempty" xml:"finishTimeGMT,omitempty"`
	// originatorId
	OriginatorId *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	// processInstanceId
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// status
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// taskId
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// taskType
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// title
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// titleEn
	TitleInEnglish *string `json:"titleInEnglish,omitempty" xml:"titleInEnglish,omitempty"`
}

func (s GetRunningTasksResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetRunningTasksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRunningTasksResponseBodyResult) SetActiveTimeGMT(v string) *GetRunningTasksResponseBodyResult {
	s.ActiveTimeGMT = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetActivityId(v string) *GetRunningTasksResponseBodyResult {
	s.ActivityId = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetActualActionerId(v string) *GetRunningTasksResponseBodyResult {
	s.ActualActionerId = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetCreateTimeGMT(v string) *GetRunningTasksResponseBodyResult {
	s.CreateTimeGMT = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetFinishTimeGMT(v string) *GetRunningTasksResponseBodyResult {
	s.FinishTimeGMT = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetOriginatorId(v string) *GetRunningTasksResponseBodyResult {
	s.OriginatorId = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetProcessInstanceId(v string) *GetRunningTasksResponseBodyResult {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetStatus(v string) *GetRunningTasksResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetTaskId(v string) *GetRunningTasksResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetTaskType(v string) *GetRunningTasksResponseBodyResult {
	s.TaskType = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetTitle(v string) *GetRunningTasksResponseBodyResult {
	s.Title = &v
	return s
}

func (s *GetRunningTasksResponseBodyResult) SetTitleInEnglish(v string) *GetRunningTasksResponseBodyResult {
	s.TitleInEnglish = &v
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
	// accountId
	AccountId *int64 `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// corpList
	CorpList []*GetSaleUserInfoByUserIdResponseBodyCorpList `json:"corpList,omitempty" xml:"corpList,omitempty" type:"Repeated"`
	// userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// userName
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s GetSaleUserInfoByUserIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSaleUserInfoByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetSaleUserInfoByUserIdResponseBody) SetAccountId(v int64) *GetSaleUserInfoByUserIdResponseBody {
	s.AccountId = &v
	return s
}

func (s *GetSaleUserInfoByUserIdResponseBody) SetCorpList(v []*GetSaleUserInfoByUserIdResponseBodyCorpList) *GetSaleUserInfoByUserIdResponseBody {
	s.CorpList = v
	return s
}

func (s *GetSaleUserInfoByUserIdResponseBody) SetUserId(v string) *GetSaleUserInfoByUserIdResponseBody {
	s.UserId = &v
	return s
}

func (s *GetSaleUserInfoByUserIdResponseBody) SetUserName(v string) *GetSaleUserInfoByUserIdResponseBody {
	s.UserName = &v
	return s
}

type GetSaleUserInfoByUserIdResponseBodyCorpList struct {
	// corpId
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// corpName
	CorpName *string `json:"corpName,omitempty" xml:"corpName,omitempty"`
	// namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
}

func (s GetSaleUserInfoByUserIdResponseBodyCorpList) String() string {
	return tea.Prettify(s)
}

func (s GetSaleUserInfoByUserIdResponseBodyCorpList) GoString() string {
	return s.String()
}

func (s *GetSaleUserInfoByUserIdResponseBodyCorpList) SetCorpId(v string) *GetSaleUserInfoByUserIdResponseBodyCorpList {
	s.CorpId = &v
	return s
}

func (s *GetSaleUserInfoByUserIdResponseBodyCorpList) SetCorpName(v string) *GetSaleUserInfoByUserIdResponseBodyCorpList {
	s.CorpName = &v
	return s
}

func (s *GetSaleUserInfoByUserIdResponseBodyCorpList) SetNamespace(v string) *GetSaleUserInfoByUserIdResponseBodyCorpList {
	s.Namespace = &v
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

type GetTaskCopiesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetTaskCopiesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetTaskCopiesHeaders) GoString() string {
	return s.String()
}

func (s *GetTaskCopiesHeaders) SetCommonHeaders(v map[string]*string) *GetTaskCopiesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTaskCopiesHeaders) SetXAcsDingtalkAccessToken(v string) *GetTaskCopiesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetTaskCopiesRequest struct {
	// 应用ID
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 创建时间开始; 时间戳
	CreateFromTimeGMT *int64 `json:"createFromTimeGMT,omitempty" xml:"createFromTimeGMT,omitempty"`
	// 创建时间结束; 时间戳
	CreateToTimeGMT *int64 `json:"createToTimeGMT,omitempty" xml:"createToTimeGMT,omitempty"`
	// 关键词
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 语言环境; 可选值：zh_CN/en_US
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 当前页; 必须大于0 默认1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页记录数; 必须大于0 默认10 最大值：100
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 流程code列表
	ProcessCodes *string `json:"processCodes,omitempty" xml:"processCodes,omitempty"`
	// 验权token
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetTaskCopiesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskCopiesRequest) GoString() string {
	return s.String()
}

func (s *GetTaskCopiesRequest) SetAppType(v string) *GetTaskCopiesRequest {
	s.AppType = &v
	return s
}

func (s *GetTaskCopiesRequest) SetCreateFromTimeGMT(v int64) *GetTaskCopiesRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *GetTaskCopiesRequest) SetCreateToTimeGMT(v int64) *GetTaskCopiesRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *GetTaskCopiesRequest) SetKeyword(v string) *GetTaskCopiesRequest {
	s.Keyword = &v
	return s
}

func (s *GetTaskCopiesRequest) SetLanguage(v string) *GetTaskCopiesRequest {
	s.Language = &v
	return s
}

func (s *GetTaskCopiesRequest) SetPageNumber(v int32) *GetTaskCopiesRequest {
	s.PageNumber = &v
	return s
}

func (s *GetTaskCopiesRequest) SetPageSize(v int32) *GetTaskCopiesRequest {
	s.PageSize = &v
	return s
}

func (s *GetTaskCopiesRequest) SetProcessCodes(v string) *GetTaskCopiesRequest {
	s.ProcessCodes = &v
	return s
}

func (s *GetTaskCopiesRequest) SetSystemToken(v string) *GetTaskCopiesRequest {
	s.SystemToken = &v
	return s
}

func (s *GetTaskCopiesRequest) SetUserId(v string) *GetTaskCopiesRequest {
	s.UserId = &v
	return s
}

type GetTaskCopiesResponseBody struct {
	// 数据
	Data []*GetTaskCopiesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 当前第几页
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 总数量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetTaskCopiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskCopiesResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskCopiesResponseBody) SetData(v []*GetTaskCopiesResponseBodyData) *GetTaskCopiesResponseBody {
	s.Data = v
	return s
}

func (s *GetTaskCopiesResponseBody) SetPageNumber(v int64) *GetTaskCopiesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetTaskCopiesResponseBody) SetTotalCount(v int64) *GetTaskCopiesResponseBody {
	s.TotalCount = &v
	return s
}

type GetTaskCopiesResponseBodyData struct {
	// actionerId
	ActionExecutorId []*string `json:"actionExecutorId,omitempty" xml:"actionExecutorId,omitempty" type:"Repeated"`
	// actionerName
	ActionExecutorName []*string `json:"actionExecutorName,omitempty" xml:"actionExecutorName,omitempty" type:"Repeated"`
	// appType
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// carbonActivityId
	CarbonActivityId *string `json:"carbonActivityId,omitempty" xml:"carbonActivityId,omitempty"`
	// 创建时间
	CreateTimeGMT *string `json:"createTimeGMT,omitempty" xml:"createTimeGMT,omitempty"`
	// currentActivityInstances
	CurrentActivityInstances []*GetTaskCopiesResponseBodyDataCurrentActivityInstances `json:"currentActivityInstances,omitempty" xml:"currentActivityInstances,omitempty" type:"Repeated"`
	// dataMap
	DataMap map[string]interface{} `json:"dataMap,omitempty" xml:"dataMap,omitempty"`
	// dataType
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// 结束时间
	FinishTimeGMT *string `json:"finishTimeGMT,omitempty" xml:"finishTimeGMT,omitempty"`
	// formInstanceId
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// formUuid
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 实例数据
	InstanceValue *string `json:"instanceValue,omitempty" xml:"instanceValue,omitempty"`
	// modifiedTime
	ModifiedTimeGMT *string `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	// originatorAvatar
	OriginatorAvatar *string `json:"originatorAvatar,omitempty" xml:"originatorAvatar,omitempty"`
	// originatorDisplayName
	OriginatorDisplayName *string `json:"originatorDisplayName,omitempty" xml:"originatorDisplayName,omitempty"`
	// originatorId
	OriginatorId *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	// processApprovedResult
	ProcessApprovedResult *string `json:"processApprovedResult,omitempty" xml:"processApprovedResult,omitempty"`
	// processApprovedResultText
	ProcessApprovedResultText *string `json:"processApprovedResultText,omitempty" xml:"processApprovedResultText,omitempty"`
	// processCode
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// 流程id
	ProcessId *int64 `json:"processId,omitempty" xml:"processId,omitempty"`
	// processInstanceId
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// processInstanceStatus
	ProcessInstanceStatus *string `json:"processInstanceStatus,omitempty" xml:"processInstanceStatus,omitempty"`
	// processInstanceStatusText
	ProcessInstanceStatusText *string `json:"processInstanceStatusText,omitempty" xml:"processInstanceStatusText,omitempty"`
	// processName
	ProcessName *string `json:"processName,omitempty" xml:"processName,omitempty"`
	// 序列号
	SerialNumber *string `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
	// 标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 版本
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetTaskCopiesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTaskCopiesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTaskCopiesResponseBodyData) SetActionExecutorId(v []*string) *GetTaskCopiesResponseBodyData {
	s.ActionExecutorId = v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetActionExecutorName(v []*string) *GetTaskCopiesResponseBodyData {
	s.ActionExecutorName = v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetAppType(v string) *GetTaskCopiesResponseBodyData {
	s.AppType = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetCarbonActivityId(v string) *GetTaskCopiesResponseBodyData {
	s.CarbonActivityId = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetCreateTimeGMT(v string) *GetTaskCopiesResponseBodyData {
	s.CreateTimeGMT = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetCurrentActivityInstances(v []*GetTaskCopiesResponseBodyDataCurrentActivityInstances) *GetTaskCopiesResponseBodyData {
	s.CurrentActivityInstances = v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetDataMap(v map[string]interface{}) *GetTaskCopiesResponseBodyData {
	s.DataMap = v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetDataType(v string) *GetTaskCopiesResponseBodyData {
	s.DataType = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetFinishTimeGMT(v string) *GetTaskCopiesResponseBodyData {
	s.FinishTimeGMT = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetFormInstanceId(v string) *GetTaskCopiesResponseBodyData {
	s.FormInstanceId = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetFormUuid(v string) *GetTaskCopiesResponseBodyData {
	s.FormUuid = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetInstanceValue(v string) *GetTaskCopiesResponseBodyData {
	s.InstanceValue = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetModifiedTimeGMT(v string) *GetTaskCopiesResponseBodyData {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetOriginatorAvatar(v string) *GetTaskCopiesResponseBodyData {
	s.OriginatorAvatar = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetOriginatorDisplayName(v string) *GetTaskCopiesResponseBodyData {
	s.OriginatorDisplayName = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetOriginatorId(v string) *GetTaskCopiesResponseBodyData {
	s.OriginatorId = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetProcessApprovedResult(v string) *GetTaskCopiesResponseBodyData {
	s.ProcessApprovedResult = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetProcessApprovedResultText(v string) *GetTaskCopiesResponseBodyData {
	s.ProcessApprovedResultText = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetProcessCode(v string) *GetTaskCopiesResponseBodyData {
	s.ProcessCode = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetProcessId(v int64) *GetTaskCopiesResponseBodyData {
	s.ProcessId = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetProcessInstanceId(v string) *GetTaskCopiesResponseBodyData {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetProcessInstanceStatus(v string) *GetTaskCopiesResponseBodyData {
	s.ProcessInstanceStatus = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetProcessInstanceStatusText(v string) *GetTaskCopiesResponseBodyData {
	s.ProcessInstanceStatusText = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetProcessName(v string) *GetTaskCopiesResponseBodyData {
	s.ProcessName = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetSerialNumber(v string) *GetTaskCopiesResponseBodyData {
	s.SerialNumber = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetTitle(v string) *GetTaskCopiesResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetTaskCopiesResponseBodyData) SetVersion(v int64) *GetTaskCopiesResponseBodyData {
	s.Version = &v
	return s
}

type GetTaskCopiesResponseBodyDataCurrentActivityInstances struct {
	// 节点id
	ActivityId *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
	// 节点实例状态
	ActivityInstanceStatus *string `json:"activityInstanceStatus,omitempty" xml:"activityInstanceStatus,omitempty"`
	// 节点名称
	ActivityName *string `json:"activityName,omitempty" xml:"activityName,omitempty"`
	// 节点英文名称
	ActivityNameInEnglish *string `json:"activityNameInEnglish,omitempty" xml:"activityNameInEnglish,omitempty"`
	// 数据id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s GetTaskCopiesResponseBodyDataCurrentActivityInstances) String() string {
	return tea.Prettify(s)
}

func (s GetTaskCopiesResponseBodyDataCurrentActivityInstances) GoString() string {
	return s.String()
}

func (s *GetTaskCopiesResponseBodyDataCurrentActivityInstances) SetActivityId(v string) *GetTaskCopiesResponseBodyDataCurrentActivityInstances {
	s.ActivityId = &v
	return s
}

func (s *GetTaskCopiesResponseBodyDataCurrentActivityInstances) SetActivityInstanceStatus(v string) *GetTaskCopiesResponseBodyDataCurrentActivityInstances {
	s.ActivityInstanceStatus = &v
	return s
}

func (s *GetTaskCopiesResponseBodyDataCurrentActivityInstances) SetActivityName(v string) *GetTaskCopiesResponseBodyDataCurrentActivityInstances {
	s.ActivityName = &v
	return s
}

func (s *GetTaskCopiesResponseBodyDataCurrentActivityInstances) SetActivityNameInEnglish(v string) *GetTaskCopiesResponseBodyDataCurrentActivityInstances {
	s.ActivityNameInEnglish = &v
	return s
}

func (s *GetTaskCopiesResponseBodyDataCurrentActivityInstances) SetId(v int64) *GetTaskCopiesResponseBodyDataCurrentActivityInstances {
	s.Id = &v
	return s
}

type GetTaskCopiesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTaskCopiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTaskCopiesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskCopiesResponse) GoString() string {
	return s.String()
}

func (s *GetTaskCopiesResponse) SetHeaders(v map[string]*string) *GetTaskCopiesResponse {
	s.Headers = v
	return s
}

func (s *GetTaskCopiesResponse) SetBody(v *GetTaskCopiesResponseBody) *GetTaskCopiesResponse {
	s.Body = v
	return s
}

type ListApplicationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListApplicationHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationHeaders) GoString() string {
	return s.String()
}

func (s *ListApplicationHeaders) SetCommonHeaders(v map[string]*string) *ListApplicationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListApplicationHeaders) SetXAcsDingtalkAccessToken(v string) *ListApplicationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListApplicationRequest struct {
	// 应用过滤条件, 不填则获取发布到了宜搭应用中心的宜搭应用, 填createdByMe获取我创建的宜搭应用, 填managedByMe获取我管理的宜搭应用
	AppFilter *string `json:"appFilter,omitempty" xml:"appFilter,omitempty"`
	// 应用名称检索关键词
	AppNameSearchKeyword *string `json:"appNameSearchKeyword,omitempty" xml:"appNameSearchKeyword,omitempty"`
	// 钉钉组织id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 第几页
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页大小
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// corpId+userId+CorpToken做md5加密计算生成的字符串(每个企业有自己的唯一corpToken), 获取具体计算详情请联系宜搭 dingtalk://dingtalkclient/action/sendmsg?dingtalk_id=somjffs
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationRequest) SetAppFilter(v string) *ListApplicationRequest {
	s.AppFilter = &v
	return s
}

func (s *ListApplicationRequest) SetAppNameSearchKeyword(v string) *ListApplicationRequest {
	s.AppNameSearchKeyword = &v
	return s
}

func (s *ListApplicationRequest) SetCorpId(v string) *ListApplicationRequest {
	s.CorpId = &v
	return s
}

func (s *ListApplicationRequest) SetPageNumber(v int32) *ListApplicationRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationRequest) SetPageSize(v int32) *ListApplicationRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationRequest) SetToken(v string) *ListApplicationRequest {
	s.Token = &v
	return s
}

func (s *ListApplicationRequest) SetUserId(v string) *ListApplicationRequest {
	s.UserId = &v
	return s
}

type ListApplicationResponseBody struct {
	// 数据
	Data []*ListApplicationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 当前第几页
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 总数量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationResponseBody) SetData(v []*ListApplicationResponseBodyData) *ListApplicationResponseBody {
	s.Data = v
	return s
}

func (s *ListApplicationResponseBody) SetPageNumber(v int64) *ListApplicationResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationResponseBody) SetTotalCount(v int64) *ListApplicationResponseBody {
	s.TotalCount = &v
	return s
}

type ListApplicationResponseBodyData struct {
	// 宜搭应用配置
	AppConfig *string `json:"appConfig,omitempty" xml:"appConfig,omitempty"`
	// 宜搭应用编码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 应用状态
	ApplicationStatus *string `json:"applicationStatus,omitempty" xml:"applicationStatus,omitempty"`
	// 钉钉组织id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 创建者的userId
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// 描述信息
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 宜搭图标编码
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 是否被删除了
	Inexistence *string `json:"inexistence,omitempty" xml:"inexistence,omitempty"`
	// 名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 子组织的钉钉CorpId
	SubCorpId *string `json:"subCorpId,omitempty" xml:"subCorpId,omitempty"`
}

func (s ListApplicationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListApplicationResponseBodyData) SetAppConfig(v string) *ListApplicationResponseBodyData {
	s.AppConfig = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetAppType(v string) *ListApplicationResponseBodyData {
	s.AppType = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetApplicationStatus(v string) *ListApplicationResponseBodyData {
	s.ApplicationStatus = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetCorpId(v string) *ListApplicationResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetCreatorUserId(v string) *ListApplicationResponseBodyData {
	s.CreatorUserId = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetDescription(v string) *ListApplicationResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetIcon(v string) *ListApplicationResponseBodyData {
	s.Icon = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetInexistence(v string) *ListApplicationResponseBodyData {
	s.Inexistence = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetName(v string) *ListApplicationResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetSubCorpId(v string) *ListApplicationResponseBodyData {
	s.SubCorpId = &v
	return s
}

type ListApplicationResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationResponse) SetHeaders(v map[string]*string) *ListApplicationResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationResponse) SetBody(v *ListApplicationResponseBody) *ListApplicationResponse {
	s.Body = v
	return s
}

type ListApplicationAuthorizationServiceApplicationInformationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListApplicationAuthorizationServiceApplicationInformationHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationAuthorizationServiceApplicationInformationHeaders) GoString() string {
	return s.String()
}

func (s *ListApplicationAuthorizationServiceApplicationInformationHeaders) SetCommonHeaders(v map[string]*string) *ListApplicationAuthorizationServiceApplicationInformationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListApplicationAuthorizationServiceApplicationInformationHeaders) SetXAcsDingtalkAccessToken(v string) *ListApplicationAuthorizationServiceApplicationInformationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListApplicationAuthorizationServiceApplicationInformationRequest struct {
	AccessKey     *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	CallerUnionId *string `json:"callerUnionId,omitempty" xml:"callerUnionId,omitempty"`
	PageNumber    *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize      *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListApplicationAuthorizationServiceApplicationInformationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationAuthorizationServiceApplicationInformationRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationAuthorizationServiceApplicationInformationRequest) SetAccessKey(v string) *ListApplicationAuthorizationServiceApplicationInformationRequest {
	s.AccessKey = &v
	return s
}

func (s *ListApplicationAuthorizationServiceApplicationInformationRequest) SetCallerUnionId(v string) *ListApplicationAuthorizationServiceApplicationInformationRequest {
	s.CallerUnionId = &v
	return s
}

func (s *ListApplicationAuthorizationServiceApplicationInformationRequest) SetPageNumber(v int32) *ListApplicationAuthorizationServiceApplicationInformationRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationAuthorizationServiceApplicationInformationRequest) SetPageSize(v int32) *ListApplicationAuthorizationServiceApplicationInformationRequest {
	s.PageSize = &v
	return s
}

type ListApplicationAuthorizationServiceApplicationInformationResponseBody struct {
	// applicationInformation
	ApplicationInformation []*ListApplicationAuthorizationServiceApplicationInformationResponseBodyApplicationInformation `json:"applicationInformation,omitempty" xml:"applicationInformation,omitempty" type:"Repeated"`
	// 当前第几页
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页大小
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 总数量
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListApplicationAuthorizationServiceApplicationInformationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationAuthorizationServiceApplicationInformationResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationAuthorizationServiceApplicationInformationResponseBody) SetApplicationInformation(v []*ListApplicationAuthorizationServiceApplicationInformationResponseBodyApplicationInformation) *ListApplicationAuthorizationServiceApplicationInformationResponseBody {
	s.ApplicationInformation = v
	return s
}

func (s *ListApplicationAuthorizationServiceApplicationInformationResponseBody) SetPageNumber(v int32) *ListApplicationAuthorizationServiceApplicationInformationResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationAuthorizationServiceApplicationInformationResponseBody) SetPageSize(v int32) *ListApplicationAuthorizationServiceApplicationInformationResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListApplicationAuthorizationServiceApplicationInformationResponseBody) SetTotalCount(v int32) *ListApplicationAuthorizationServiceApplicationInformationResponseBody {
	s.TotalCount = &v
	return s
}

type ListApplicationAuthorizationServiceApplicationInformationResponseBodyApplicationInformation struct {
	// appName
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// appType
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// attachmentUsageAmount
	AttachmentUsageAmount *int64 `json:"attachmentUsageAmount,omitempty" xml:"attachmentUsageAmount,omitempty"`
	// instanceUsageAmount
	InstanceUsageAmount *int64 `json:"instanceUsageAmount,omitempty" xml:"instanceUsageAmount,omitempty"`
	// usagePlugins
	UsagePlugins []*ListApplicationAuthorizationServiceApplicationInformationResponseBodyApplicationInformationUsagePlugins `json:"usagePlugins,omitempty" xml:"usagePlugins,omitempty" type:"Repeated"`
}

func (s ListApplicationAuthorizationServiceApplicationInformationResponseBodyApplicationInformation) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationAuthorizationServiceApplicationInformationResponseBodyApplicationInformation) GoString() string {
	return s.String()
}

func (s *ListApplicationAuthorizationServiceApplicationInformationResponseBodyApplicationInformation) SetAppName(v string) *ListApplicationAuthorizationServiceApplicationInformationResponseBodyApplicationInformation {
	s.AppName = &v
	return s
}

func (s *ListApplicationAuthorizationServiceApplicationInformationResponseBodyApplicationInformation) SetAppType(v string) *ListApplicationAuthorizationServiceApplicationInformationResponseBodyApplicationInformation {
	s.AppType = &v
	return s
}

func (s *ListApplicationAuthorizationServiceApplicationInformationResponseBodyApplicationInformation) SetAttachmentUsageAmount(v int64) *ListApplicationAuthorizationServiceApplicationInformationResponseBodyApplicationInformation {
	s.AttachmentUsageAmount = &v
	return s
}

func (s *ListApplicationAuthorizationServiceApplicationInformationResponseBodyApplicationInformation) SetInstanceUsageAmount(v int64) *ListApplicationAuthorizationServiceApplicationInformationResponseBodyApplicationInformation {
	s.InstanceUsageAmount = &v
	return s
}

func (s *ListApplicationAuthorizationServiceApplicationInformationResponseBodyApplicationInformation) SetUsagePlugins(v []*ListApplicationAuthorizationServiceApplicationInformationResponseBodyApplicationInformationUsagePlugins) *ListApplicationAuthorizationServiceApplicationInformationResponseBodyApplicationInformation {
	s.UsagePlugins = v
	return s
}

type ListApplicationAuthorizationServiceApplicationInformationResponseBodyApplicationInformationUsagePlugins struct {
	// iconUrl
	IconUrl *string `json:"iconUrl,omitempty" xml:"iconUrl,omitempty"`
	// pluginName
	PluginName *string `json:"pluginName,omitempty" xml:"pluginName,omitempty"`
}

func (s ListApplicationAuthorizationServiceApplicationInformationResponseBodyApplicationInformationUsagePlugins) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationAuthorizationServiceApplicationInformationResponseBodyApplicationInformationUsagePlugins) GoString() string {
	return s.String()
}

func (s *ListApplicationAuthorizationServiceApplicationInformationResponseBodyApplicationInformationUsagePlugins) SetIconUrl(v string) *ListApplicationAuthorizationServiceApplicationInformationResponseBodyApplicationInformationUsagePlugins {
	s.IconUrl = &v
	return s
}

func (s *ListApplicationAuthorizationServiceApplicationInformationResponseBodyApplicationInformationUsagePlugins) SetPluginName(v string) *ListApplicationAuthorizationServiceApplicationInformationResponseBodyApplicationInformationUsagePlugins {
	s.PluginName = &v
	return s
}

type ListApplicationAuthorizationServiceApplicationInformationResponse struct {
	Headers map[string]*string                                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListApplicationAuthorizationServiceApplicationInformationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListApplicationAuthorizationServiceApplicationInformationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationAuthorizationServiceApplicationInformationResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationAuthorizationServiceApplicationInformationResponse) SetHeaders(v map[string]*string) *ListApplicationAuthorizationServiceApplicationInformationResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationAuthorizationServiceApplicationInformationResponse) SetBody(v *ListApplicationAuthorizationServiceApplicationInformationResponseBody) *ListApplicationAuthorizationServiceApplicationInformationResponse {
	s.Body = v
	return s
}

type ListApplicationAuthorizationServiceConnectorInformationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListApplicationAuthorizationServiceConnectorInformationHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationAuthorizationServiceConnectorInformationHeaders) GoString() string {
	return s.String()
}

func (s *ListApplicationAuthorizationServiceConnectorInformationHeaders) SetCommonHeaders(v map[string]*string) *ListApplicationAuthorizationServiceConnectorInformationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListApplicationAuthorizationServiceConnectorInformationHeaders) SetXAcsDingtalkAccessToken(v string) *ListApplicationAuthorizationServiceConnectorInformationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListApplicationAuthorizationServiceConnectorInformationRequest struct {
	AccessKey  *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	CallerUid  *string `json:"callerUid,omitempty" xml:"callerUid,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListApplicationAuthorizationServiceConnectorInformationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationAuthorizationServiceConnectorInformationRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationAuthorizationServiceConnectorInformationRequest) SetAccessKey(v string) *ListApplicationAuthorizationServiceConnectorInformationRequest {
	s.AccessKey = &v
	return s
}

func (s *ListApplicationAuthorizationServiceConnectorInformationRequest) SetCallerUid(v string) *ListApplicationAuthorizationServiceConnectorInformationRequest {
	s.CallerUid = &v
	return s
}

func (s *ListApplicationAuthorizationServiceConnectorInformationRequest) SetPageNumber(v int32) *ListApplicationAuthorizationServiceConnectorInformationRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationAuthorizationServiceConnectorInformationRequest) SetPageSize(v int32) *ListApplicationAuthorizationServiceConnectorInformationRequest {
	s.PageSize = &v
	return s
}

type ListApplicationAuthorizationServiceConnectorInformationResponseBody struct {
	// 当前第几页
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页大小
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// pluginInfos
	PlugInformation []*ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformation `json:"plugInformation,omitempty" xml:"plugInformation,omitempty" type:"Repeated"`
	// 总数量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListApplicationAuthorizationServiceConnectorInformationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationAuthorizationServiceConnectorInformationResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationAuthorizationServiceConnectorInformationResponseBody) SetPageNumber(v int32) *ListApplicationAuthorizationServiceConnectorInformationResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationAuthorizationServiceConnectorInformationResponseBody) SetPageSize(v int32) *ListApplicationAuthorizationServiceConnectorInformationResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListApplicationAuthorizationServiceConnectorInformationResponseBody) SetPlugInformation(v []*ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformation) *ListApplicationAuthorizationServiceConnectorInformationResponseBody {
	s.PlugInformation = v
	return s
}

func (s *ListApplicationAuthorizationServiceConnectorInformationResponseBody) SetTotalCount(v int64) *ListApplicationAuthorizationServiceConnectorInformationResponseBody {
	s.TotalCount = &v
	return s
}

type ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformation struct {
	// apps
	Applications []*ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformationApplications `json:"applications,omitempty" xml:"applications,omitempty" type:"Repeated"`
	// iconUrl
	IconUrl *string `json:"iconUrl,omitempty" xml:"iconUrl,omitempty"`
	// pluginName
	PlugName *string `json:"plugName,omitempty" xml:"plugName,omitempty"`
	// pluginPayType
	PlugPayType *int32 `json:"plugPayType,omitempty" xml:"plugPayType,omitempty"`
	// pluginStatus
	PlugStatus *int32 `json:"plugStatus,omitempty" xml:"plugStatus,omitempty"`
	// pluginTotalAmount
	PlugTotalAmount *int64 `json:"plugTotalAmount,omitempty" xml:"plugTotalAmount,omitempty"`
	// pluginUsageAmount
	PlugUsageAmount *int64 `json:"plugUsageAmount,omitempty" xml:"plugUsageAmount,omitempty"`
	// pluginUuid
	PlugUuid *string `json:"plugUuid,omitempty" xml:"plugUuid,omitempty"`
}

func (s ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformation) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformation) GoString() string {
	return s.String()
}

func (s *ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformation) SetApplications(v []*ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformationApplications) *ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformation {
	s.Applications = v
	return s
}

func (s *ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformation) SetIconUrl(v string) *ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformation {
	s.IconUrl = &v
	return s
}

func (s *ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformation) SetPlugName(v string) *ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformation {
	s.PlugName = &v
	return s
}

func (s *ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformation) SetPlugPayType(v int32) *ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformation {
	s.PlugPayType = &v
	return s
}

func (s *ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformation) SetPlugStatus(v int32) *ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformation {
	s.PlugStatus = &v
	return s
}

func (s *ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformation) SetPlugTotalAmount(v int64) *ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformation {
	s.PlugTotalAmount = &v
	return s
}

func (s *ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformation) SetPlugUsageAmount(v int64) *ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformation {
	s.PlugUsageAmount = &v
	return s
}

func (s *ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformation) SetPlugUuid(v string) *ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformation {
	s.PlugUuid = &v
	return s
}

type ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformationApplications struct {
	// appName
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
}

func (s ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformationApplications) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformationApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformationApplications) SetAppName(v string) *ListApplicationAuthorizationServiceConnectorInformationResponseBodyPlugInformationApplications {
	s.AppName = &v
	return s
}

type ListApplicationAuthorizationServiceConnectorInformationResponse struct {
	Headers map[string]*string                                                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListApplicationAuthorizationServiceConnectorInformationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListApplicationAuthorizationServiceConnectorInformationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationAuthorizationServiceConnectorInformationResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationAuthorizationServiceConnectorInformationResponse) SetHeaders(v map[string]*string) *ListApplicationAuthorizationServiceConnectorInformationResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationAuthorizationServiceConnectorInformationResponse) SetBody(v *ListApplicationAuthorizationServiceConnectorInformationResponseBody) *ListApplicationAuthorizationServiceConnectorInformationResponse {
	s.Body = v
	return s
}

type ListApplicationInformationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListApplicationInformationHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationInformationHeaders) GoString() string {
	return s.String()
}

func (s *ListApplicationInformationHeaders) SetCommonHeaders(v map[string]*string) *ListApplicationInformationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListApplicationInformationHeaders) SetXAcsDingtalkAccessToken(v string) *ListApplicationInformationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListApplicationInformationRequest struct {
	AccessKey  *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	CallerUid  *string `json:"callerUid,omitempty" xml:"callerUid,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListApplicationInformationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationInformationRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationInformationRequest) SetAccessKey(v string) *ListApplicationInformationRequest {
	s.AccessKey = &v
	return s
}

func (s *ListApplicationInformationRequest) SetCallerUid(v string) *ListApplicationInformationRequest {
	s.CallerUid = &v
	return s
}

func (s *ListApplicationInformationRequest) SetPageNumber(v int32) *ListApplicationInformationRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationInformationRequest) SetPageSize(v int32) *ListApplicationInformationRequest {
	s.PageSize = &v
	return s
}

type ListApplicationInformationResponseBody struct {
	// applicationInformation
	ApplicationInformation []*ListApplicationInformationResponseBodyApplicationInformation `json:"applicationInformation,omitempty" xml:"applicationInformation,omitempty" type:"Repeated"`
	// 当前第几页
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页大小
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 总数量
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListApplicationInformationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationInformationResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationInformationResponseBody) SetApplicationInformation(v []*ListApplicationInformationResponseBodyApplicationInformation) *ListApplicationInformationResponseBody {
	s.ApplicationInformation = v
	return s
}

func (s *ListApplicationInformationResponseBody) SetPageNumber(v int32) *ListApplicationInformationResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationInformationResponseBody) SetPageSize(v int32) *ListApplicationInformationResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListApplicationInformationResponseBody) SetTotalCount(v int32) *ListApplicationInformationResponseBody {
	s.TotalCount = &v
	return s
}

type ListApplicationInformationResponseBodyApplicationInformation struct {
	// appName
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
	// appType
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// attachmentUsageAmount
	AttachmentUsageAmount *int64 `json:"attachmentUsageAmount,omitempty" xml:"attachmentUsageAmount,omitempty"`
	// instanceUsageAmount
	InstanceUsageAmount *int64 `json:"instanceUsageAmount,omitempty" xml:"instanceUsageAmount,omitempty"`
	// usagePlugins
	UsagePlugins []*ListApplicationInformationResponseBodyApplicationInformationUsagePlugins `json:"usagePlugins,omitempty" xml:"usagePlugins,omitempty" type:"Repeated"`
}

func (s ListApplicationInformationResponseBodyApplicationInformation) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationInformationResponseBodyApplicationInformation) GoString() string {
	return s.String()
}

func (s *ListApplicationInformationResponseBodyApplicationInformation) SetAppName(v string) *ListApplicationInformationResponseBodyApplicationInformation {
	s.AppName = &v
	return s
}

func (s *ListApplicationInformationResponseBodyApplicationInformation) SetAppType(v string) *ListApplicationInformationResponseBodyApplicationInformation {
	s.AppType = &v
	return s
}

func (s *ListApplicationInformationResponseBodyApplicationInformation) SetAttachmentUsageAmount(v int64) *ListApplicationInformationResponseBodyApplicationInformation {
	s.AttachmentUsageAmount = &v
	return s
}

func (s *ListApplicationInformationResponseBodyApplicationInformation) SetInstanceUsageAmount(v int64) *ListApplicationInformationResponseBodyApplicationInformation {
	s.InstanceUsageAmount = &v
	return s
}

func (s *ListApplicationInformationResponseBodyApplicationInformation) SetUsagePlugins(v []*ListApplicationInformationResponseBodyApplicationInformationUsagePlugins) *ListApplicationInformationResponseBodyApplicationInformation {
	s.UsagePlugins = v
	return s
}

type ListApplicationInformationResponseBodyApplicationInformationUsagePlugins struct {
	// iconUrl
	IconUrl *string `json:"iconUrl,omitempty" xml:"iconUrl,omitempty"`
	// pluginName
	PluginName *string `json:"pluginName,omitempty" xml:"pluginName,omitempty"`
}

func (s ListApplicationInformationResponseBodyApplicationInformationUsagePlugins) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationInformationResponseBodyApplicationInformationUsagePlugins) GoString() string {
	return s.String()
}

func (s *ListApplicationInformationResponseBodyApplicationInformationUsagePlugins) SetIconUrl(v string) *ListApplicationInformationResponseBodyApplicationInformationUsagePlugins {
	s.IconUrl = &v
	return s
}

func (s *ListApplicationInformationResponseBodyApplicationInformationUsagePlugins) SetPluginName(v string) *ListApplicationInformationResponseBodyApplicationInformationUsagePlugins {
	s.PluginName = &v
	return s
}

type ListApplicationInformationResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListApplicationInformationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListApplicationInformationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationInformationResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationInformationResponse) SetHeaders(v map[string]*string) *ListApplicationInformationResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationInformationResponse) SetBody(v *ListApplicationInformationResponseBody) *ListApplicationInformationResponse {
	s.Body = v
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
	// callerUid
	CallerUid *string `json:"callerUid,omitempty" xml:"callerUid,omitempty"`
	// currentPage
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// pageSize
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
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

func (s *ListCommodityRequest) SetCallerUid(v string) *ListCommodityRequest {
	s.CallerUid = &v
	return s
}

func (s *ListCommodityRequest) SetPageNumber(v int32) *ListCommodityRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCommodityRequest) SetPageSize(v int32) *ListCommodityRequest {
	s.PageSize = &v
	return s
}

type ListCommodityResponseBody struct {
	// commodityVOList
	CommodityVOList []*ListCommodityResponseBodyCommodityVOList `json:"commodityVOList,omitempty" xml:"commodityVOList,omitempty" type:"Repeated"`
	// 当前第几页
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页大小
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 总数量
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListCommodityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCommodityResponseBody) GoString() string {
	return s.String()
}

func (s *ListCommodityResponseBody) SetCommodityVOList(v []*ListCommodityResponseBodyCommodityVOList) *ListCommodityResponseBody {
	s.CommodityVOList = v
	return s
}

func (s *ListCommodityResponseBody) SetPageNumber(v int32) *ListCommodityResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCommodityResponseBody) SetPageSize(v int32) *ListCommodityResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCommodityResponseBody) SetTotalCount(v int32) *ListCommodityResponseBody {
	s.TotalCount = &v
	return s
}

type ListCommodityResponseBodyCommodityVOList struct {
	// accountDistributionNumber
	AccountDistributionNumber *int32 `json:"accountDistributionNumber,omitempty" xml:"accountDistributionNumber,omitempty"`
	// accountNum
	AccountNumber *int32 `json:"accountNumber,omitempty" xml:"accountNumber,omitempty"`
	// activationCode
	ActivationCode *string `json:"activationCode,omitempty" xml:"activationCode,omitempty"`
	// buyDate
	BuyDateGMT *string `json:"buyDateGMT,omitempty" xml:"buyDateGMT,omitempty"`
	// expireDate
	ExpireDateGMT *string `json:"expireDateGMT,omitempty" xml:"expireDateGMT,omitempty"`
	// instanceId
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// status
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// version
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListCommodityResponseBodyCommodityVOList) String() string {
	return tea.Prettify(s)
}

func (s ListCommodityResponseBodyCommodityVOList) GoString() string {
	return s.String()
}

func (s *ListCommodityResponseBodyCommodityVOList) SetAccountDistributionNumber(v int32) *ListCommodityResponseBodyCommodityVOList {
	s.AccountDistributionNumber = &v
	return s
}

func (s *ListCommodityResponseBodyCommodityVOList) SetAccountNumber(v int32) *ListCommodityResponseBodyCommodityVOList {
	s.AccountNumber = &v
	return s
}

func (s *ListCommodityResponseBodyCommodityVOList) SetActivationCode(v string) *ListCommodityResponseBodyCommodityVOList {
	s.ActivationCode = &v
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

func (s *ListCommodityResponseBodyCommodityVOList) SetInstanceId(v string) *ListCommodityResponseBodyCommodityVOList {
	s.InstanceId = &v
	return s
}

func (s *ListCommodityResponseBodyCommodityVOList) SetStatus(v string) *ListCommodityResponseBodyCommodityVOList {
	s.Status = &v
	return s
}

func (s *ListCommodityResponseBodyCommodityVOList) SetVersion(v int32) *ListCommodityResponseBodyCommodityVOList {
	s.Version = &v
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

type ListConnectorInformationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListConnectorInformationHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListConnectorInformationHeaders) GoString() string {
	return s.String()
}

func (s *ListConnectorInformationHeaders) SetCommonHeaders(v map[string]*string) *ListConnectorInformationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListConnectorInformationHeaders) SetXAcsDingtalkAccessToken(v string) *ListConnectorInformationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListConnectorInformationRequest struct {
	AccessKey  *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	CallerUid  *string `json:"callerUid,omitempty" xml:"callerUid,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListConnectorInformationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConnectorInformationRequest) GoString() string {
	return s.String()
}

func (s *ListConnectorInformationRequest) SetAccessKey(v string) *ListConnectorInformationRequest {
	s.AccessKey = &v
	return s
}

func (s *ListConnectorInformationRequest) SetCallerUid(v string) *ListConnectorInformationRequest {
	s.CallerUid = &v
	return s
}

func (s *ListConnectorInformationRequest) SetPageNumber(v int32) *ListConnectorInformationRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConnectorInformationRequest) SetPageSize(v int32) *ListConnectorInformationRequest {
	s.PageSize = &v
	return s
}

type ListConnectorInformationResponseBody struct {
	// 当前第几页
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页大小
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// pluginInfos
	PluginInfos []*ListConnectorInformationResponseBodyPluginInfos `json:"pluginInfos,omitempty" xml:"pluginInfos,omitempty" type:"Repeated"`
	// 总数量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListConnectorInformationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConnectorInformationResponseBody) GoString() string {
	return s.String()
}

func (s *ListConnectorInformationResponseBody) SetPageNumber(v int32) *ListConnectorInformationResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListConnectorInformationResponseBody) SetPageSize(v int32) *ListConnectorInformationResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListConnectorInformationResponseBody) SetPluginInfos(v []*ListConnectorInformationResponseBodyPluginInfos) *ListConnectorInformationResponseBody {
	s.PluginInfos = v
	return s
}

func (s *ListConnectorInformationResponseBody) SetTotalCount(v int64) *ListConnectorInformationResponseBody {
	s.TotalCount = &v
	return s
}

type ListConnectorInformationResponseBodyPluginInfos struct {
	// apps
	Apps []*ListConnectorInformationResponseBodyPluginInfosApps `json:"apps,omitempty" xml:"apps,omitempty" type:"Repeated"`
	// iconUrl
	IconUrl *string `json:"iconUrl,omitempty" xml:"iconUrl,omitempty"`
	// pluginName
	PluginName *string `json:"pluginName,omitempty" xml:"pluginName,omitempty"`
	// pluginPayType
	PluginPayType *int32 `json:"pluginPayType,omitempty" xml:"pluginPayType,omitempty"`
	// pluginStatus
	PluginStatus *int32 `json:"pluginStatus,omitempty" xml:"pluginStatus,omitempty"`
	// pluginTotalAmount
	PluginTotalAmount *int64 `json:"pluginTotalAmount,omitempty" xml:"pluginTotalAmount,omitempty"`
	// pluginUsageAmount
	PluginUsageAmount *int64 `json:"pluginUsageAmount,omitempty" xml:"pluginUsageAmount,omitempty"`
	// pluginUuid
	PluginUuid *string `json:"pluginUuid,omitempty" xml:"pluginUuid,omitempty"`
}

func (s ListConnectorInformationResponseBodyPluginInfos) String() string {
	return tea.Prettify(s)
}

func (s ListConnectorInformationResponseBodyPluginInfos) GoString() string {
	return s.String()
}

func (s *ListConnectorInformationResponseBodyPluginInfos) SetApps(v []*ListConnectorInformationResponseBodyPluginInfosApps) *ListConnectorInformationResponseBodyPluginInfos {
	s.Apps = v
	return s
}

func (s *ListConnectorInformationResponseBodyPluginInfos) SetIconUrl(v string) *ListConnectorInformationResponseBodyPluginInfos {
	s.IconUrl = &v
	return s
}

func (s *ListConnectorInformationResponseBodyPluginInfos) SetPluginName(v string) *ListConnectorInformationResponseBodyPluginInfos {
	s.PluginName = &v
	return s
}

func (s *ListConnectorInformationResponseBodyPluginInfos) SetPluginPayType(v int32) *ListConnectorInformationResponseBodyPluginInfos {
	s.PluginPayType = &v
	return s
}

func (s *ListConnectorInformationResponseBodyPluginInfos) SetPluginStatus(v int32) *ListConnectorInformationResponseBodyPluginInfos {
	s.PluginStatus = &v
	return s
}

func (s *ListConnectorInformationResponseBodyPluginInfos) SetPluginTotalAmount(v int64) *ListConnectorInformationResponseBodyPluginInfos {
	s.PluginTotalAmount = &v
	return s
}

func (s *ListConnectorInformationResponseBodyPluginInfos) SetPluginUsageAmount(v int64) *ListConnectorInformationResponseBodyPluginInfos {
	s.PluginUsageAmount = &v
	return s
}

func (s *ListConnectorInformationResponseBodyPluginInfos) SetPluginUuid(v string) *ListConnectorInformationResponseBodyPluginInfos {
	s.PluginUuid = &v
	return s
}

type ListConnectorInformationResponseBodyPluginInfosApps struct {
	// appName
	AppName *string `json:"appName,omitempty" xml:"appName,omitempty"`
}

func (s ListConnectorInformationResponseBodyPluginInfosApps) String() string {
	return tea.Prettify(s)
}

func (s ListConnectorInformationResponseBodyPluginInfosApps) GoString() string {
	return s.String()
}

func (s *ListConnectorInformationResponseBodyPluginInfosApps) SetAppName(v string) *ListConnectorInformationResponseBodyPluginInfosApps {
	s.AppName = &v
	return s
}

type ListConnectorInformationResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListConnectorInformationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConnectorInformationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConnectorInformationResponse) GoString() string {
	return s.String()
}

func (s *ListConnectorInformationResponse) SetHeaders(v map[string]*string) *ListConnectorInformationResponse {
	s.Headers = v
	return s
}

func (s *ListConnectorInformationResponse) SetBody(v *ListConnectorInformationResponseBody) *ListConnectorInformationResponse {
	s.Body = v
	return s
}

type ListFormRemarksHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListFormRemarksHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListFormRemarksHeaders) GoString() string {
	return s.String()
}

func (s *ListFormRemarksHeaders) SetCommonHeaders(v map[string]*string) *ListFormRemarksHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListFormRemarksHeaders) SetXAcsDingtalkAccessToken(v string) *ListFormRemarksHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListFormRemarksRequest struct {
	// 宜搭应用编码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 表单实例id列表
	FormInstanceIdList []*string `json:"formInstanceIdList,omitempty" xml:"formInstanceIdList,omitempty" type:"Repeated"`
	// 表单编码
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 宜搭应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListFormRemarksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFormRemarksRequest) GoString() string {
	return s.String()
}

func (s *ListFormRemarksRequest) SetAppType(v string) *ListFormRemarksRequest {
	s.AppType = &v
	return s
}

func (s *ListFormRemarksRequest) SetFormInstanceIdList(v []*string) *ListFormRemarksRequest {
	s.FormInstanceIdList = v
	return s
}

func (s *ListFormRemarksRequest) SetFormUuid(v string) *ListFormRemarksRequest {
	s.FormUuid = &v
	return s
}

func (s *ListFormRemarksRequest) SetSystemToken(v string) *ListFormRemarksRequest {
	s.SystemToken = &v
	return s
}

func (s *ListFormRemarksRequest) SetUserId(v string) *ListFormRemarksRequest {
	s.UserId = &v
	return s
}

type ListFormRemarksResponseBody struct {
	// formRemarkVoMap
	FormRemarkVoMap map[string]interface{} `json:"formRemarkVoMap,omitempty" xml:"formRemarkVoMap,omitempty"`
}

func (s ListFormRemarksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFormRemarksResponseBody) GoString() string {
	return s.String()
}

func (s *ListFormRemarksResponseBody) SetFormRemarkVoMap(v map[string]interface{}) *ListFormRemarksResponseBody {
	s.FormRemarkVoMap = v
	return s
}

type ListFormRemarksResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFormRemarksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFormRemarksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFormRemarksResponse) GoString() string {
	return s.String()
}

func (s *ListFormRemarksResponse) SetHeaders(v map[string]*string) *ListFormRemarksResponse {
	s.Headers = v
	return s
}

func (s *ListFormRemarksResponse) SetBody(v *ListFormRemarksResponseBody) *ListFormRemarksResponse {
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
	// 页面类型
	FormType *string `json:"formType,omitempty" xml:"formType,omitempty"`
	// 语言
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 评论人钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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

func (s *ListNavigationByFormTypeRequest) SetFormType(v string) *ListNavigationByFormTypeRequest {
	s.FormType = &v
	return s
}

func (s *ListNavigationByFormTypeRequest) SetLanguage(v string) *ListNavigationByFormTypeRequest {
	s.Language = &v
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
	// formUuid
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// processCode
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// title
	Title *ListNavigationByFormTypeResponseBodyResultTitle `json:"title,omitempty" xml:"title,omitempty" type:"Struct"`
}

func (s ListNavigationByFormTypeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListNavigationByFormTypeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListNavigationByFormTypeResponseBodyResult) SetFormUuid(v string) *ListNavigationByFormTypeResponseBodyResult {
	s.FormUuid = &v
	return s
}

func (s *ListNavigationByFormTypeResponseBodyResult) SetProcessCode(v string) *ListNavigationByFormTypeResponseBodyResult {
	s.ProcessCode = &v
	return s
}

func (s *ListNavigationByFormTypeResponseBodyResult) SetTitle(v *ListNavigationByFormTypeResponseBodyResultTitle) *ListNavigationByFormTypeResponseBodyResult {
	s.Title = v
	return s
}

type ListNavigationByFormTypeResponseBodyResultTitle struct {
	// 中文名称
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// 英文名称
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
	// type
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListNavigationByFormTypeResponseBodyResultTitle) String() string {
	return tea.Prettify(s)
}

func (s ListNavigationByFormTypeResponseBodyResultTitle) GoString() string {
	return s.String()
}

func (s *ListNavigationByFormTypeResponseBodyResultTitle) SetNameInChinese(v string) *ListNavigationByFormTypeResponseBodyResultTitle {
	s.NameInChinese = &v
	return s
}

func (s *ListNavigationByFormTypeResponseBodyResultTitle) SetNameInEnglish(v string) *ListNavigationByFormTypeResponseBodyResultTitle {
	s.NameInEnglish = &v
	return s
}

func (s *ListNavigationByFormTypeResponseBodyResultTitle) SetType(v string) *ListNavigationByFormTypeResponseBodyResultTitle {
	s.Type = &v
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

type ListOperationLogsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListOperationLogsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListOperationLogsHeaders) GoString() string {
	return s.String()
}

func (s *ListOperationLogsHeaders) SetCommonHeaders(v map[string]*string) *ListOperationLogsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListOperationLogsHeaders) SetXAcsDingtalkAccessToken(v string) *ListOperationLogsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListOperationLogsRequest struct {
	// 宜搭应用编码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 表单实例id列表
	FormInstanceIdList []*string `json:"formInstanceIdList,omitempty" xml:"formInstanceIdList,omitempty" type:"Repeated"`
	// 表单编码
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 宜搭应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListOperationLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOperationLogsRequest) GoString() string {
	return s.String()
}

func (s *ListOperationLogsRequest) SetAppType(v string) *ListOperationLogsRequest {
	s.AppType = &v
	return s
}

func (s *ListOperationLogsRequest) SetFormInstanceIdList(v []*string) *ListOperationLogsRequest {
	s.FormInstanceIdList = v
	return s
}

func (s *ListOperationLogsRequest) SetFormUuid(v string) *ListOperationLogsRequest {
	s.FormUuid = &v
	return s
}

func (s *ListOperationLogsRequest) SetSystemToken(v string) *ListOperationLogsRequest {
	s.SystemToken = &v
	return s
}

func (s *ListOperationLogsRequest) SetUserId(v string) *ListOperationLogsRequest {
	s.UserId = &v
	return s
}

type ListOperationLogsResponseBody struct {
	// 操作记录对象
	OperationLogMap map[string]interface{} `json:"operationLogMap,omitempty" xml:"operationLogMap,omitempty"`
}

func (s ListOperationLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOperationLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOperationLogsResponseBody) SetOperationLogMap(v map[string]interface{}) *ListOperationLogsResponseBody {
	s.OperationLogMap = v
	return s
}

type ListOperationLogsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOperationLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOperationLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOperationLogsResponse) GoString() string {
	return s.String()
}

func (s *ListOperationLogsResponse) SetHeaders(v map[string]*string) *ListOperationLogsResponse {
	s.Headers = v
	return s
}

func (s *ListOperationLogsResponse) SetBody(v *ListOperationLogsResponseBody) *ListOperationLogsResponse {
	s.Body = v
	return s
}

type ListTableDataByFormInstanceIdTableIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListTableDataByFormInstanceIdTableIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListTableDataByFormInstanceIdTableIdHeaders) GoString() string {
	return s.String()
}

func (s *ListTableDataByFormInstanceIdTableIdHeaders) SetCommonHeaders(v map[string]*string) *ListTableDataByFormInstanceIdTableIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdHeaders) SetXAcsDingtalkAccessToken(v string) *ListTableDataByFormInstanceIdTableIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListTableDataByFormInstanceIdTableIdRequest struct {
	// 应用编码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 表单ID
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 当前页
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页记录数
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 需要查找的子表单组件的唯一标识
	TableFieldId *string `json:"tableFieldId,omitempty" xml:"tableFieldId,omitempty"`
	// 钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListTableDataByFormInstanceIdTableIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTableDataByFormInstanceIdTableIdRequest) GoString() string {
	return s.String()
}

func (s *ListTableDataByFormInstanceIdTableIdRequest) SetAppType(v string) *ListTableDataByFormInstanceIdTableIdRequest {
	s.AppType = &v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdRequest) SetFormUuid(v string) *ListTableDataByFormInstanceIdTableIdRequest {
	s.FormUuid = &v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdRequest) SetPageNumber(v int32) *ListTableDataByFormInstanceIdTableIdRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdRequest) SetPageSize(v int32) *ListTableDataByFormInstanceIdTableIdRequest {
	s.PageSize = &v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdRequest) SetSystemToken(v string) *ListTableDataByFormInstanceIdTableIdRequest {
	s.SystemToken = &v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdRequest) SetTableFieldId(v string) *ListTableDataByFormInstanceIdTableIdRequest {
	s.TableFieldId = &v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdRequest) SetUserId(v string) *ListTableDataByFormInstanceIdTableIdRequest {
	s.UserId = &v
	return s
}

type ListTableDataByFormInstanceIdTableIdResponseBody struct {
	// data
	Data []map[string]interface{} `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 当前第几页
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 总数量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListTableDataByFormInstanceIdTableIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTableDataByFormInstanceIdTableIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListTableDataByFormInstanceIdTableIdResponseBody) SetData(v []map[string]interface{}) *ListTableDataByFormInstanceIdTableIdResponseBody {
	s.Data = v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdResponseBody) SetPageNumber(v int64) *ListTableDataByFormInstanceIdTableIdResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdResponseBody) SetTotalCount(v int64) *ListTableDataByFormInstanceIdTableIdResponseBody {
	s.TotalCount = &v
	return s
}

type ListTableDataByFormInstanceIdTableIdResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTableDataByFormInstanceIdTableIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTableDataByFormInstanceIdTableIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTableDataByFormInstanceIdTableIdResponse) GoString() string {
	return s.String()
}

func (s *ListTableDataByFormInstanceIdTableIdResponse) SetHeaders(v map[string]*string) *ListTableDataByFormInstanceIdTableIdResponse {
	s.Headers = v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdResponse) SetBody(v *ListTableDataByFormInstanceIdTableIdResponseBody) *ListTableDataByFormInstanceIdTableIdResponse {
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

type NotifyAuthorizationResultHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s NotifyAuthorizationResultHeaders) String() string {
	return tea.Prettify(s)
}

func (s NotifyAuthorizationResultHeaders) GoString() string {
	return s.String()
}

func (s *NotifyAuthorizationResultHeaders) SetCommonHeaders(v map[string]*string) *NotifyAuthorizationResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *NotifyAuthorizationResultHeaders) SetXAcsDingtalkAccessToken(v string) *NotifyAuthorizationResultHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type NotifyAuthorizationResultRequest struct {
	AccessKey     *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	AccountNumber *string `json:"accountNumber,omitempty" xml:"accountNumber,omitempty"`
	BeginTimeGMT  *int64  `json:"beginTimeGMT,omitempty" xml:"beginTimeGMT,omitempty"`
	CallerUid     *string `json:"callerUid,omitempty" xml:"callerUid,omitempty"`
	ChargeType    *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	CommerceType  *string `json:"commerceType,omitempty" xml:"commerceType,omitempty"`
	CommodityType *string `json:"commodityType,omitempty" xml:"commodityType,omitempty"`
	EndTimeGMT    *int64  `json:"endTimeGMT,omitempty" xml:"endTimeGMT,omitempty"`
	InstanceId    *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	InstanceName  *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// 阿里云产品code
	ProduceCode *string `json:"produceCode,omitempty" xml:"produceCode,omitempty"`
}

func (s NotifyAuthorizationResultRequest) String() string {
	return tea.Prettify(s)
}

func (s NotifyAuthorizationResultRequest) GoString() string {
	return s.String()
}

func (s *NotifyAuthorizationResultRequest) SetAccessKey(v string) *NotifyAuthorizationResultRequest {
	s.AccessKey = &v
	return s
}

func (s *NotifyAuthorizationResultRequest) SetAccountNumber(v string) *NotifyAuthorizationResultRequest {
	s.AccountNumber = &v
	return s
}

func (s *NotifyAuthorizationResultRequest) SetBeginTimeGMT(v int64) *NotifyAuthorizationResultRequest {
	s.BeginTimeGMT = &v
	return s
}

func (s *NotifyAuthorizationResultRequest) SetCallerUid(v string) *NotifyAuthorizationResultRequest {
	s.CallerUid = &v
	return s
}

func (s *NotifyAuthorizationResultRequest) SetChargeType(v string) *NotifyAuthorizationResultRequest {
	s.ChargeType = &v
	return s
}

func (s *NotifyAuthorizationResultRequest) SetCommerceType(v string) *NotifyAuthorizationResultRequest {
	s.CommerceType = &v
	return s
}

func (s *NotifyAuthorizationResultRequest) SetCommodityType(v string) *NotifyAuthorizationResultRequest {
	s.CommodityType = &v
	return s
}

func (s *NotifyAuthorizationResultRequest) SetEndTimeGMT(v int64) *NotifyAuthorizationResultRequest {
	s.EndTimeGMT = &v
	return s
}

func (s *NotifyAuthorizationResultRequest) SetInstanceId(v string) *NotifyAuthorizationResultRequest {
	s.InstanceId = &v
	return s
}

func (s *NotifyAuthorizationResultRequest) SetInstanceName(v string) *NotifyAuthorizationResultRequest {
	s.InstanceName = &v
	return s
}

func (s *NotifyAuthorizationResultRequest) SetProduceCode(v string) *NotifyAuthorizationResultRequest {
	s.ProduceCode = &v
	return s
}

type NotifyAuthorizationResultResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s NotifyAuthorizationResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NotifyAuthorizationResultResponseBody) GoString() string {
	return s.String()
}

func (s *NotifyAuthorizationResultResponseBody) SetResult(v bool) *NotifyAuthorizationResultResponseBody {
	s.Result = &v
	return s
}

type NotifyAuthorizationResultResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *NotifyAuthorizationResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s NotifyAuthorizationResultResponse) String() string {
	return tea.Prettify(s)
}

func (s NotifyAuthorizationResultResponse) GoString() string {
	return s.String()
}

func (s *NotifyAuthorizationResultResponse) SetHeaders(v map[string]*string) *NotifyAuthorizationResultResponse {
	s.Headers = v
	return s
}

func (s *NotifyAuthorizationResultResponse) SetBody(v *NotifyAuthorizationResultResponseBody) *NotifyAuthorizationResultResponse {
	s.Body = v
	return s
}

type PageFormBaseInfosHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PageFormBaseInfosHeaders) String() string {
	return tea.Prettify(s)
}

func (s PageFormBaseInfosHeaders) GoString() string {
	return s.String()
}

func (s *PageFormBaseInfosHeaders) SetCommonHeaders(v map[string]*string) *PageFormBaseInfosHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PageFormBaseInfosHeaders) SetXAcsDingtalkAccessToken(v string) *PageFormBaseInfosHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PageFormBaseInfosRequest struct {
	// 应用编码
	AppKey *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	// 表单类型列表，可传"process", "receipt"
	FormTypeList []*string `json:"formTypeList,omitempty" xml:"formTypeList,omitempty" type:"Repeated"`
	// 语言。可选值：zh_CN/en_US 默认：zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 当前页码
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// 每页数量，最大100
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 应用秘钥。在应用数据中获取。
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PageFormBaseInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s PageFormBaseInfosRequest) GoString() string {
	return s.String()
}

func (s *PageFormBaseInfosRequest) SetAppKey(v string) *PageFormBaseInfosRequest {
	s.AppKey = &v
	return s
}

func (s *PageFormBaseInfosRequest) SetFormTypeList(v []*string) *PageFormBaseInfosRequest {
	s.FormTypeList = v
	return s
}

func (s *PageFormBaseInfosRequest) SetLanguage(v string) *PageFormBaseInfosRequest {
	s.Language = &v
	return s
}

func (s *PageFormBaseInfosRequest) SetPageIndex(v int32) *PageFormBaseInfosRequest {
	s.PageIndex = &v
	return s
}

func (s *PageFormBaseInfosRequest) SetPageSize(v int32) *PageFormBaseInfosRequest {
	s.PageSize = &v
	return s
}

func (s *PageFormBaseInfosRequest) SetSystemToken(v string) *PageFormBaseInfosRequest {
	s.SystemToken = &v
	return s
}

func (s *PageFormBaseInfosRequest) SetUserId(v string) *PageFormBaseInfosRequest {
	s.UserId = &v
	return s
}

type PageFormBaseInfosResponseBody struct {
	// 结果集
	Result *PageFormBaseInfosResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PageFormBaseInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageFormBaseInfosResponseBody) GoString() string {
	return s.String()
}

func (s *PageFormBaseInfosResponseBody) SetResult(v *PageFormBaseInfosResponseBodyResult) *PageFormBaseInfosResponseBody {
	s.Result = v
	return s
}

func (s *PageFormBaseInfosResponseBody) SetSuccess(v bool) *PageFormBaseInfosResponseBody {
	s.Success = &v
	return s
}

type PageFormBaseInfosResponseBodyResult struct {
	// 当前页
	CurrentPage *int32                                     `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	Data        []*PageFormBaseInfosResponseBodyResultData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 总行数
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s PageFormBaseInfosResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PageFormBaseInfosResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PageFormBaseInfosResponseBodyResult) SetCurrentPage(v int32) *PageFormBaseInfosResponseBodyResult {
	s.CurrentPage = &v
	return s
}

func (s *PageFormBaseInfosResponseBodyResult) SetData(v []*PageFormBaseInfosResponseBodyResultData) *PageFormBaseInfosResponseBodyResult {
	s.Data = v
	return s
}

func (s *PageFormBaseInfosResponseBodyResult) SetTotalCount(v int32) *PageFormBaseInfosResponseBodyResult {
	s.TotalCount = &v
	return s
}

type PageFormBaseInfosResponseBodyResultData struct {
	Creator   *string                                       `json:"creator,omitempty" xml:"creator,omitempty"`
	FormType  *string                                       `json:"formType,omitempty" xml:"formType,omitempty"`
	FormUuid  *string                                       `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	GmtCreate *string                                       `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	Title     *PageFormBaseInfosResponseBodyResultDataTitle `json:"title,omitempty" xml:"title,omitempty" type:"Struct"`
}

func (s PageFormBaseInfosResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s PageFormBaseInfosResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *PageFormBaseInfosResponseBodyResultData) SetCreator(v string) *PageFormBaseInfosResponseBodyResultData {
	s.Creator = &v
	return s
}

func (s *PageFormBaseInfosResponseBodyResultData) SetFormType(v string) *PageFormBaseInfosResponseBodyResultData {
	s.FormType = &v
	return s
}

func (s *PageFormBaseInfosResponseBodyResultData) SetFormUuid(v string) *PageFormBaseInfosResponseBodyResultData {
	s.FormUuid = &v
	return s
}

func (s *PageFormBaseInfosResponseBodyResultData) SetGmtCreate(v string) *PageFormBaseInfosResponseBodyResultData {
	s.GmtCreate = &v
	return s
}

func (s *PageFormBaseInfosResponseBodyResultData) SetTitle(v *PageFormBaseInfosResponseBodyResultDataTitle) *PageFormBaseInfosResponseBodyResultData {
	s.Title = v
	return s
}

type PageFormBaseInfosResponseBodyResultDataTitle struct {
	EnUS *string `json:"enUS,omitempty" xml:"enUS,omitempty"`
	ZhCN *string `json:"zhCN,omitempty" xml:"zhCN,omitempty"`
}

func (s PageFormBaseInfosResponseBodyResultDataTitle) String() string {
	return tea.Prettify(s)
}

func (s PageFormBaseInfosResponseBodyResultDataTitle) GoString() string {
	return s.String()
}

func (s *PageFormBaseInfosResponseBodyResultDataTitle) SetEnUS(v string) *PageFormBaseInfosResponseBodyResultDataTitle {
	s.EnUS = &v
	return s
}

func (s *PageFormBaseInfosResponseBodyResultDataTitle) SetZhCN(v string) *PageFormBaseInfosResponseBodyResultDataTitle {
	s.ZhCN = &v
	return s
}

type PageFormBaseInfosResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PageFormBaseInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PageFormBaseInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s PageFormBaseInfosResponse) GoString() string {
	return s.String()
}

func (s *PageFormBaseInfosResponse) SetHeaders(v map[string]*string) *PageFormBaseInfosResponse {
	s.Headers = v
	return s
}

func (s *PageFormBaseInfosResponse) SetBody(v *PageFormBaseInfosResponseBody) *PageFormBaseInfosResponse {
	s.Body = v
	return s
}

type QueryServiceRecordHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryServiceRecordHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceRecordHeaders) GoString() string {
	return s.String()
}

func (s *QueryServiceRecordHeaders) SetCommonHeaders(v map[string]*string) *QueryServiceRecordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryServiceRecordHeaders) SetXAcsDingtalkAccessToken(v string) *QueryServiceRecordHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryServiceRecordRequest struct {
	// 宜搭应用编码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 表单编码
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 服务类型
	HookType *string `json:"hookType,omitempty" xml:"hookType,omitempty"`
	// 本次服务调用的唯一ID
	HookUuid *string `json:"hookUuid,omitempty" xml:"hookUuid,omitempty"`
	// 表单实例ID
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// 服务在此时间之后调用的
	InvokeAfterDateGMT *string `json:"invokeAfterDateGMT,omitempty" xml:"invokeAfterDateGMT,omitempty"`
	// 服务在此时间之前调用的
	InvokeBeforeDateGMT *string `json:"invokeBeforeDateGMT,omitempty" xml:"invokeBeforeDateGMT,omitempty"`
	// 服务调用状态
	InvokeStatus *string `json:"invokeStatus,omitempty" xml:"invokeStatus,omitempty"`
	// 分页第几页
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页大小
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 服务调用地址包含的部分字符串，用于模糊搜索
	RequestUrl *string `json:"requestUrl,omitempty" xml:"requestUrl,omitempty"`
	// 被重试的服务调用唯一ID(此次服务调用是重试哪个执行失败的服务调用)
	SourceUuid *string `json:"sourceUuid,omitempty" xml:"sourceUuid,omitempty"`
	// 服务调用是否成功(不传此参数则查询全部的)
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 宜搭应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 操作人的钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryServiceRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceRecordRequest) GoString() string {
	return s.String()
}

func (s *QueryServiceRecordRequest) SetAppType(v string) *QueryServiceRecordRequest {
	s.AppType = &v
	return s
}

func (s *QueryServiceRecordRequest) SetFormUuid(v string) *QueryServiceRecordRequest {
	s.FormUuid = &v
	return s
}

func (s *QueryServiceRecordRequest) SetHookType(v string) *QueryServiceRecordRequest {
	s.HookType = &v
	return s
}

func (s *QueryServiceRecordRequest) SetHookUuid(v string) *QueryServiceRecordRequest {
	s.HookUuid = &v
	return s
}

func (s *QueryServiceRecordRequest) SetInstanceId(v string) *QueryServiceRecordRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryServiceRecordRequest) SetInvokeAfterDateGMT(v string) *QueryServiceRecordRequest {
	s.InvokeAfterDateGMT = &v
	return s
}

func (s *QueryServiceRecordRequest) SetInvokeBeforeDateGMT(v string) *QueryServiceRecordRequest {
	s.InvokeBeforeDateGMT = &v
	return s
}

func (s *QueryServiceRecordRequest) SetInvokeStatus(v string) *QueryServiceRecordRequest {
	s.InvokeStatus = &v
	return s
}

func (s *QueryServiceRecordRequest) SetPageNumber(v int32) *QueryServiceRecordRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryServiceRecordRequest) SetPageSize(v int32) *QueryServiceRecordRequest {
	s.PageSize = &v
	return s
}

func (s *QueryServiceRecordRequest) SetRequestUrl(v string) *QueryServiceRecordRequest {
	s.RequestUrl = &v
	return s
}

func (s *QueryServiceRecordRequest) SetSourceUuid(v string) *QueryServiceRecordRequest {
	s.SourceUuid = &v
	return s
}

func (s *QueryServiceRecordRequest) SetSuccess(v bool) *QueryServiceRecordRequest {
	s.Success = &v
	return s
}

func (s *QueryServiceRecordRequest) SetSystemToken(v string) *QueryServiceRecordRequest {
	s.SystemToken = &v
	return s
}

func (s *QueryServiceRecordRequest) SetUserId(v string) *QueryServiceRecordRequest {
	s.UserId = &v
	return s
}

type QueryServiceRecordResponseBody struct {
	// 总数量
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 服务调用记录数组
	Values []*QueryServiceRecordResponseBodyValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s QueryServiceRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceRecordResponseBody) GoString() string {
	return s.String()
}

func (s *QueryServiceRecordResponseBody) SetTotalCount(v int32) *QueryServiceRecordResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryServiceRecordResponseBody) SetValues(v []*QueryServiceRecordResponseBodyValues) *QueryServiceRecordResponseBody {
	s.Values = v
	return s
}

type QueryServiceRecordResponseBodyValues struct {
	// 表单实例id
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// 表单编码
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 服务类型
	HookType *string `json:"hookType,omitempty" xml:"hookType,omitempty"`
	// 本次服务调用的唯一ID
	HookUuid *string `json:"hookUuid,omitempty" xml:"hookUuid,omitempty"`
	// 服务调用的实际入参
	InvokeParameter *string `json:"invokeParameter,omitempty" xml:"invokeParameter,omitempty"`
	// 服务调用的返回结果
	InvokeResult *string `json:"invokeResult,omitempty" xml:"invokeResult,omitempty"`
	// 服务调用状态
	InvokeStatus *string `json:"invokeStatus,omitempty" xml:"invokeStatus,omitempty"`
	// 服务调用是否成功
	InvokeSuccess *string `json:"invokeSuccess,omitempty" xml:"invokeSuccess,omitempty"`
	// 服务调用地址
	InvokeUrl *string `json:"invokeUrl,omitempty" xml:"invokeUrl,omitempty"`
	// 宜搭调用目标服务时传的实际参数
	ServiceContent *string `json:"serviceContent,omitempty" xml:"serviceContent,omitempty"`
	// 服务名称
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// 服务调用的实际入参
	ServiceParameter *string `json:"serviceParameter,omitempty" xml:"serviceParameter,omitempty"`
	// 重试的服务调用唯一ID(此次服务调用是重试哪个执行失败的服务调用)
	SourceUuid *string `json:"sourceUuid,omitempty" xml:"sourceUuid,omitempty"`
}

func (s QueryServiceRecordResponseBodyValues) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceRecordResponseBodyValues) GoString() string {
	return s.String()
}

func (s *QueryServiceRecordResponseBodyValues) SetFormInstanceId(v string) *QueryServiceRecordResponseBodyValues {
	s.FormInstanceId = &v
	return s
}

func (s *QueryServiceRecordResponseBodyValues) SetFormUuid(v string) *QueryServiceRecordResponseBodyValues {
	s.FormUuid = &v
	return s
}

func (s *QueryServiceRecordResponseBodyValues) SetHookType(v string) *QueryServiceRecordResponseBodyValues {
	s.HookType = &v
	return s
}

func (s *QueryServiceRecordResponseBodyValues) SetHookUuid(v string) *QueryServiceRecordResponseBodyValues {
	s.HookUuid = &v
	return s
}

func (s *QueryServiceRecordResponseBodyValues) SetInvokeParameter(v string) *QueryServiceRecordResponseBodyValues {
	s.InvokeParameter = &v
	return s
}

func (s *QueryServiceRecordResponseBodyValues) SetInvokeResult(v string) *QueryServiceRecordResponseBodyValues {
	s.InvokeResult = &v
	return s
}

func (s *QueryServiceRecordResponseBodyValues) SetInvokeStatus(v string) *QueryServiceRecordResponseBodyValues {
	s.InvokeStatus = &v
	return s
}

func (s *QueryServiceRecordResponseBodyValues) SetInvokeSuccess(v string) *QueryServiceRecordResponseBodyValues {
	s.InvokeSuccess = &v
	return s
}

func (s *QueryServiceRecordResponseBodyValues) SetInvokeUrl(v string) *QueryServiceRecordResponseBodyValues {
	s.InvokeUrl = &v
	return s
}

func (s *QueryServiceRecordResponseBodyValues) SetServiceContent(v string) *QueryServiceRecordResponseBodyValues {
	s.ServiceContent = &v
	return s
}

func (s *QueryServiceRecordResponseBodyValues) SetServiceName(v string) *QueryServiceRecordResponseBodyValues {
	s.ServiceName = &v
	return s
}

func (s *QueryServiceRecordResponseBodyValues) SetServiceParameter(v string) *QueryServiceRecordResponseBodyValues {
	s.ServiceParameter = &v
	return s
}

func (s *QueryServiceRecordResponseBodyValues) SetSourceUuid(v string) *QueryServiceRecordResponseBodyValues {
	s.SourceUuid = &v
	return s
}

type QueryServiceRecordResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryServiceRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryServiceRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceRecordResponse) GoString() string {
	return s.String()
}

func (s *QueryServiceRecordResponse) SetHeaders(v map[string]*string) *QueryServiceRecordResponse {
	s.Headers = v
	return s
}

func (s *QueryServiceRecordResponse) SetBody(v *QueryServiceRecordResponseBody) *QueryServiceRecordResponse {
	s.Body = v
	return s
}

type RedirectTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RedirectTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s RedirectTaskHeaders) GoString() string {
	return s.String()
}

func (s *RedirectTaskHeaders) SetCommonHeaders(v map[string]*string) *RedirectTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RedirectTaskHeaders) SetXAcsDingtalkAccessToken(v string) *RedirectTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RedirectTaskRequest struct {
	// 应用ID
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 是否应用管理员进行转交; ●
	// 可选项 : y/n
	//
	// ●
	// y,则userId必须传应用管理员工号，或者yida_pub_account
	//
	// ●
	// n, userId必须传任务的当前执行人
	ByManager *string `json:"byManager,omitempty" xml:"byManager,omitempty"`
	// 语言环境; 可选值：zh_CN/en_US
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 新的任务处理人工号
	NowActionExecutorId *string `json:"nowActionExecutorId,omitempty" xml:"nowActionExecutorId,omitempty"`
	// 实例ID
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// 转交备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 验权token; 在应用数据中获取。
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 任务ID
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// 钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s RedirectTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s RedirectTaskRequest) GoString() string {
	return s.String()
}

func (s *RedirectTaskRequest) SetAppType(v string) *RedirectTaskRequest {
	s.AppType = &v
	return s
}

func (s *RedirectTaskRequest) SetByManager(v string) *RedirectTaskRequest {
	s.ByManager = &v
	return s
}

func (s *RedirectTaskRequest) SetLanguage(v string) *RedirectTaskRequest {
	s.Language = &v
	return s
}

func (s *RedirectTaskRequest) SetNowActionExecutorId(v string) *RedirectTaskRequest {
	s.NowActionExecutorId = &v
	return s
}

func (s *RedirectTaskRequest) SetProcessInstanceId(v string) *RedirectTaskRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *RedirectTaskRequest) SetRemark(v string) *RedirectTaskRequest {
	s.Remark = &v
	return s
}

func (s *RedirectTaskRequest) SetSystemToken(v string) *RedirectTaskRequest {
	s.SystemToken = &v
	return s
}

func (s *RedirectTaskRequest) SetTaskId(v int64) *RedirectTaskRequest {
	s.TaskId = &v
	return s
}

func (s *RedirectTaskRequest) SetUserId(v string) *RedirectTaskRequest {
	s.UserId = &v
	return s
}

type RedirectTaskResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s RedirectTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s RedirectTaskResponse) GoString() string {
	return s.String()
}

func (s *RedirectTaskResponse) SetHeaders(v map[string]*string) *RedirectTaskResponse {
	s.Headers = v
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
	AccessKey  *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	CallerUid  *string `json:"callerUid,omitempty" xml:"callerUid,omitempty"`
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s RefundCommodityRequest) String() string {
	return tea.Prettify(s)
}

func (s RefundCommodityRequest) GoString() string {
	return s.String()
}

func (s *RefundCommodityRequest) SetAccessKey(v string) *RefundCommodityRequest {
	s.AccessKey = &v
	return s
}

func (s *RefundCommodityRequest) SetCallerUid(v string) *RefundCommodityRequest {
	s.CallerUid = &v
	return s
}

func (s *RefundCommodityRequest) SetInstanceId(v string) *RefundCommodityRequest {
	s.InstanceId = &v
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

type RegisterAccountsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RegisterAccountsHeaders) String() string {
	return tea.Prettify(s)
}

func (s RegisterAccountsHeaders) GoString() string {
	return s.String()
}

func (s *RegisterAccountsHeaders) SetCommonHeaders(v map[string]*string) *RegisterAccountsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RegisterAccountsHeaders) SetXAcsDingtalkAccessToken(v string) *RegisterAccountsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RegisterAccountsRequest struct {
	// 访问秘钥
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// 激活码
	ActiveCode *string `json:"activeCode,omitempty" xml:"activeCode,omitempty"`
	// 组织id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
}

func (s RegisterAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterAccountsRequest) GoString() string {
	return s.String()
}

func (s *RegisterAccountsRequest) SetAccessKey(v string) *RegisterAccountsRequest {
	s.AccessKey = &v
	return s
}

func (s *RegisterAccountsRequest) SetActiveCode(v string) *RegisterAccountsRequest {
	s.ActiveCode = &v
	return s
}

func (s *RegisterAccountsRequest) SetCorpId(v string) *RegisterAccountsRequest {
	s.CorpId = &v
	return s
}

type RegisterAccountsResponseBody struct {
	// 实例id
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s RegisterAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterAccountsResponseBody) SetInstanceId(v string) *RegisterAccountsResponseBody {
	s.InstanceId = &v
	return s
}

type RegisterAccountsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterAccountsResponse) GoString() string {
	return s.String()
}

func (s *RegisterAccountsResponse) SetHeaders(v map[string]*string) *RegisterAccountsResponse {
	s.Headers = v
	return s
}

func (s *RegisterAccountsResponse) SetBody(v *RegisterAccountsResponseBody) *RegisterAccountsResponse {
	s.Body = v
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
	AccessKey  *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	CallerUid  *string `json:"callerUid,omitempty" xml:"callerUid,omitempty"`
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s ReleaseCommodityRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseCommodityRequest) GoString() string {
	return s.String()
}

func (s *ReleaseCommodityRequest) SetAccessKey(v string) *ReleaseCommodityRequest {
	s.AccessKey = &v
	return s
}

func (s *ReleaseCommodityRequest) SetCallerUid(v string) *ReleaseCommodityRequest {
	s.CallerUid = &v
	return s
}

func (s *ReleaseCommodityRequest) SetInstanceId(v string) *ReleaseCommodityRequest {
	s.InstanceId = &v
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

type RemoveTenantResourceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RemoveTenantResourceHeaders) String() string {
	return tea.Prettify(s)
}

func (s RemoveTenantResourceHeaders) GoString() string {
	return s.String()
}

func (s *RemoveTenantResourceHeaders) SetCommonHeaders(v map[string]*string) *RemoveTenantResourceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveTenantResourceHeaders) SetXAcsDingtalkAccessToken(v string) *RemoveTenantResourceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RemoveTenantResourceRequest struct {
	// accessKey
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
}

func (s RemoveTenantResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveTenantResourceRequest) GoString() string {
	return s.String()
}

func (s *RemoveTenantResourceRequest) SetAccessKey(v string) *RemoveTenantResourceRequest {
	s.AccessKey = &v
	return s
}

type RemoveTenantResourceResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RemoveTenantResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveTenantResourceResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveTenantResourceResponseBody) SetResult(v bool) *RemoveTenantResourceResponseBody {
	s.Result = &v
	return s
}

type RemoveTenantResourceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveTenantResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveTenantResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveTenantResourceResponse) GoString() string {
	return s.String()
}

func (s *RemoveTenantResourceResponse) SetHeaders(v map[string]*string) *RemoveTenantResourceResponse {
	s.Headers = v
	return s
}

func (s *RemoveTenantResourceResponse) SetBody(v *RemoveTenantResourceResponseBody) *RemoveTenantResourceResponse {
	s.Body = v
	return s
}

type RenderBatchCallbackHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RenderBatchCallbackHeaders) String() string {
	return tea.Prettify(s)
}

func (s RenderBatchCallbackHeaders) GoString() string {
	return s.String()
}

func (s *RenderBatchCallbackHeaders) SetCommonHeaders(v map[string]*string) *RenderBatchCallbackHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RenderBatchCallbackHeaders) SetXAcsDingtalkAccessToken(v string) *RenderBatchCallbackHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RenderBatchCallbackRequest struct {
	// appType
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 组织id
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// 文件大小
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// language
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 名称空间
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// oss文件链接
	OssUrl *string `json:"ossUrl,omitempty" xml:"ossUrl,omitempty"`
	// 流水号
	SequenceId *string `json:"sequenceId,omitempty" xml:"sequenceId,omitempty"`
	// 源
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// systemToken
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 时间区域
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
	// userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s RenderBatchCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s RenderBatchCallbackRequest) GoString() string {
	return s.String()
}

func (s *RenderBatchCallbackRequest) SetAppType(v string) *RenderBatchCallbackRequest {
	s.AppType = &v
	return s
}

func (s *RenderBatchCallbackRequest) SetCorpId(v string) *RenderBatchCallbackRequest {
	s.CorpId = &v
	return s
}

func (s *RenderBatchCallbackRequest) SetFileSize(v int64) *RenderBatchCallbackRequest {
	s.FileSize = &v
	return s
}

func (s *RenderBatchCallbackRequest) SetLanguage(v string) *RenderBatchCallbackRequest {
	s.Language = &v
	return s
}

func (s *RenderBatchCallbackRequest) SetNamespace(v string) *RenderBatchCallbackRequest {
	s.Namespace = &v
	return s
}

func (s *RenderBatchCallbackRequest) SetOssUrl(v string) *RenderBatchCallbackRequest {
	s.OssUrl = &v
	return s
}

func (s *RenderBatchCallbackRequest) SetSequenceId(v string) *RenderBatchCallbackRequest {
	s.SequenceId = &v
	return s
}

func (s *RenderBatchCallbackRequest) SetSource(v string) *RenderBatchCallbackRequest {
	s.Source = &v
	return s
}

func (s *RenderBatchCallbackRequest) SetStatus(v string) *RenderBatchCallbackRequest {
	s.Status = &v
	return s
}

func (s *RenderBatchCallbackRequest) SetSystemToken(v string) *RenderBatchCallbackRequest {
	s.SystemToken = &v
	return s
}

func (s *RenderBatchCallbackRequest) SetTimeZone(v string) *RenderBatchCallbackRequest {
	s.TimeZone = &v
	return s
}

func (s *RenderBatchCallbackRequest) SetUserId(v string) *RenderBatchCallbackRequest {
	s.UserId = &v
	return s
}

type RenderBatchCallbackResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s RenderBatchCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s RenderBatchCallbackResponse) GoString() string {
	return s.String()
}

func (s *RenderBatchCallbackResponse) SetHeaders(v map[string]*string) *RenderBatchCallbackResponse {
	s.Headers = v
	return s
}

type RenewApplicationAuthorizationServiceOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RenewApplicationAuthorizationServiceOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s RenewApplicationAuthorizationServiceOrderHeaders) GoString() string {
	return s.String()
}

func (s *RenewApplicationAuthorizationServiceOrderHeaders) SetCommonHeaders(v map[string]*string) *RenewApplicationAuthorizationServiceOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RenewApplicationAuthorizationServiceOrderHeaders) SetXAcsDingtalkAccessToken(v string) *RenewApplicationAuthorizationServiceOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RenewApplicationAuthorizationServiceOrderRequest struct {
	// 访问秘钥
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// 调用者unionId
	CallerUnionId *string `json:"callerUnionId,omitempty" xml:"callerUnionId,omitempty"`
	// 结束时间
	EndTimeGMT *int64 `json:"endTimeGMT,omitempty" xml:"endTimeGMT,omitempty"`
	// 实例id
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s RenewApplicationAuthorizationServiceOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewApplicationAuthorizationServiceOrderRequest) GoString() string {
	return s.String()
}

func (s *RenewApplicationAuthorizationServiceOrderRequest) SetAccessKey(v string) *RenewApplicationAuthorizationServiceOrderRequest {
	s.AccessKey = &v
	return s
}

func (s *RenewApplicationAuthorizationServiceOrderRequest) SetCallerUnionId(v string) *RenewApplicationAuthorizationServiceOrderRequest {
	s.CallerUnionId = &v
	return s
}

func (s *RenewApplicationAuthorizationServiceOrderRequest) SetEndTimeGMT(v int64) *RenewApplicationAuthorizationServiceOrderRequest {
	s.EndTimeGMT = &v
	return s
}

func (s *RenewApplicationAuthorizationServiceOrderRequest) SetInstanceId(v string) *RenewApplicationAuthorizationServiceOrderRequest {
	s.InstanceId = &v
	return s
}

type RenewApplicationAuthorizationServiceOrderResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RenewApplicationAuthorizationServiceOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewApplicationAuthorizationServiceOrderResponseBody) GoString() string {
	return s.String()
}

func (s *RenewApplicationAuthorizationServiceOrderResponseBody) SetResult(v bool) *RenewApplicationAuthorizationServiceOrderResponseBody {
	s.Result = &v
	return s
}

type RenewApplicationAuthorizationServiceOrderResponse struct {
	Headers map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenewApplicationAuthorizationServiceOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenewApplicationAuthorizationServiceOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewApplicationAuthorizationServiceOrderResponse) GoString() string {
	return s.String()
}

func (s *RenewApplicationAuthorizationServiceOrderResponse) SetHeaders(v map[string]*string) *RenewApplicationAuthorizationServiceOrderResponse {
	s.Headers = v
	return s
}

func (s *RenewApplicationAuthorizationServiceOrderResponse) SetBody(v *RenewApplicationAuthorizationServiceOrderResponseBody) *RenewApplicationAuthorizationServiceOrderResponse {
	s.Body = v
	return s
}

type RenewTenantOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RenewTenantOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s RenewTenantOrderHeaders) GoString() string {
	return s.String()
}

func (s *RenewTenantOrderHeaders) SetCommonHeaders(v map[string]*string) *RenewTenantOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RenewTenantOrderHeaders) SetXAcsDingtalkAccessToken(v string) *RenewTenantOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RenewTenantOrderRequest struct {
	// 访问秘钥
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// 调用者unionId
	CallerUnionId *string `json:"callerUnionId,omitempty" xml:"callerUnionId,omitempty"`
	// 结束时间
	EndTimeGMT *int64 `json:"endTimeGMT,omitempty" xml:"endTimeGMT,omitempty"`
}

func (s RenewTenantOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewTenantOrderRequest) GoString() string {
	return s.String()
}

func (s *RenewTenantOrderRequest) SetAccessKey(v string) *RenewTenantOrderRequest {
	s.AccessKey = &v
	return s
}

func (s *RenewTenantOrderRequest) SetCallerUnionId(v string) *RenewTenantOrderRequest {
	s.CallerUnionId = &v
	return s
}

func (s *RenewTenantOrderRequest) SetEndTimeGMT(v int64) *RenewTenantOrderRequest {
	s.EndTimeGMT = &v
	return s
}

type RenewTenantOrderResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RenewTenantOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewTenantOrderResponseBody) GoString() string {
	return s.String()
}

func (s *RenewTenantOrderResponseBody) SetResult(v bool) *RenewTenantOrderResponseBody {
	s.Result = &v
	return s
}

type RenewTenantOrderResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenewTenantOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenewTenantOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewTenantOrderResponse) GoString() string {
	return s.String()
}

func (s *RenewTenantOrderResponse) SetHeaders(v map[string]*string) *RenewTenantOrderResponse {
	s.Headers = v
	return s
}

func (s *RenewTenantOrderResponse) SetBody(v *RenewTenantOrderResponseBody) *RenewTenantOrderResponse {
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
	// 表单数据
	FormDataJson *string `json:"formDataJson,omitempty" xml:"formDataJson,omitempty"`
	// 表单ID
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 语言。可选值：zh_CN/en_US 默认：zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 应用秘钥。在应用数据中获取。
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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

func (s *SaveFormDataRequest) SetFormDataJson(v string) *SaveFormDataRequest {
	s.FormDataJson = &v
	return s
}

func (s *SaveFormDataRequest) SetFormUuid(v string) *SaveFormDataRequest {
	s.FormUuid = &v
	return s
}

func (s *SaveFormDataRequest) SetLanguage(v string) *SaveFormDataRequest {
	s.Language = &v
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
	// 将评论内容通过钉钉发给指定用户, 逗号分隔
	AtUserId *string `json:"atUserId,omitempty" xml:"atUserId,omitempty"`
	// 评论内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 实例ID
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// 语言
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 对评论进行回复
	ReplyId *int64 `json:"replyId,omitempty" xml:"replyId,omitempty"`
	// 应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 评论人钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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

func (s *SaveFormRemarkRequest) SetAtUserId(v string) *SaveFormRemarkRequest {
	s.AtUserId = &v
	return s
}

func (s *SaveFormRemarkRequest) SetContent(v string) *SaveFormRemarkRequest {
	s.Content = &v
	return s
}

func (s *SaveFormRemarkRequest) SetFormInstanceId(v string) *SaveFormRemarkRequest {
	s.FormInstanceId = &v
	return s
}

func (s *SaveFormRemarkRequest) SetLanguage(v string) *SaveFormRemarkRequest {
	s.Language = &v
	return s
}

func (s *SaveFormRemarkRequest) SetReplyId(v int64) *SaveFormRemarkRequest {
	s.ReplyId = &v
	return s
}

func (s *SaveFormRemarkRequest) SetSystemToken(v string) *SaveFormRemarkRequest {
	s.SystemToken = &v
	return s
}

func (s *SaveFormRemarkRequest) SetUserId(v string) *SaveFormRemarkRequest {
	s.UserId = &v
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

type SavePrintTplDetailInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SavePrintTplDetailInfoHeaders) String() string {
	return tea.Prettify(s)
}

func (s SavePrintTplDetailInfoHeaders) GoString() string {
	return s.String()
}

func (s *SavePrintTplDetailInfoHeaders) SetCommonHeaders(v map[string]*string) *SavePrintTplDetailInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SavePrintTplDetailInfoHeaders) SetXAcsDingtalkAccessToken(v string) *SavePrintTplDetailInfoHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SavePrintTplDetailInfoRequest struct {
	// 应用代码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 模板描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 文件名配置
	FileNameConfig *string `json:"fileNameConfig,omitempty" xml:"fileNameConfig,omitempty"`
	// 表单id
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 表单版本
	FormVersion *int32 `json:"formVersion,omitempty" xml:"formVersion,omitempty"`
	// 模板的其他配置信息
	Setting *string `json:"setting,omitempty" xml:"setting,omitempty"`
	// 打印模板id
	TemplateId *int64 `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// 模板标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 用户id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 模板的VM
	Vm *string `json:"vm,omitempty" xml:"vm,omitempty"`
}

func (s SavePrintTplDetailInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SavePrintTplDetailInfoRequest) GoString() string {
	return s.String()
}

func (s *SavePrintTplDetailInfoRequest) SetAppType(v string) *SavePrintTplDetailInfoRequest {
	s.AppType = &v
	return s
}

func (s *SavePrintTplDetailInfoRequest) SetDescription(v string) *SavePrintTplDetailInfoRequest {
	s.Description = &v
	return s
}

func (s *SavePrintTplDetailInfoRequest) SetFileNameConfig(v string) *SavePrintTplDetailInfoRequest {
	s.FileNameConfig = &v
	return s
}

func (s *SavePrintTplDetailInfoRequest) SetFormUuid(v string) *SavePrintTplDetailInfoRequest {
	s.FormUuid = &v
	return s
}

func (s *SavePrintTplDetailInfoRequest) SetFormVersion(v int32) *SavePrintTplDetailInfoRequest {
	s.FormVersion = &v
	return s
}

func (s *SavePrintTplDetailInfoRequest) SetSetting(v string) *SavePrintTplDetailInfoRequest {
	s.Setting = &v
	return s
}

func (s *SavePrintTplDetailInfoRequest) SetTemplateId(v int64) *SavePrintTplDetailInfoRequest {
	s.TemplateId = &v
	return s
}

func (s *SavePrintTplDetailInfoRequest) SetTitle(v string) *SavePrintTplDetailInfoRequest {
	s.Title = &v
	return s
}

func (s *SavePrintTplDetailInfoRequest) SetUserId(v string) *SavePrintTplDetailInfoRequest {
	s.UserId = &v
	return s
}

func (s *SavePrintTplDetailInfoRequest) SetVm(v string) *SavePrintTplDetailInfoRequest {
	s.Vm = &v
	return s
}

type SavePrintTplDetailInfoResponseBody struct {
	// 模板id
	Result *int64 `json:"result,omitempty" xml:"result,omitempty"`
}

func (s SavePrintTplDetailInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SavePrintTplDetailInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SavePrintTplDetailInfoResponseBody) SetResult(v int64) *SavePrintTplDetailInfoResponseBody {
	s.Result = &v
	return s
}

type SavePrintTplDetailInfoResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SavePrintTplDetailInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SavePrintTplDetailInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SavePrintTplDetailInfoResponse) GoString() string {
	return s.String()
}

func (s *SavePrintTplDetailInfoResponse) SetHeaders(v map[string]*string) *SavePrintTplDetailInfoResponse {
	s.Headers = v
	return s
}

func (s *SavePrintTplDetailInfoResponse) SetBody(v *SavePrintTplDetailInfoResponseBody) *SavePrintTplDetailInfoResponse {
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
	// activationCode
	ActivationCode *string `json:"activationCode,omitempty" xml:"activationCode,omitempty"`
	// authType
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty"`
	// expireTime
	ExpireTimeGMT *string `json:"expireTimeGMT,omitempty" xml:"expireTimeGMT,omitempty"`
	// instanceId
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// status
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SearchActivationCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchActivationCodeResponseBody) GoString() string {
	return s.String()
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

func (s *SearchActivationCodeResponseBody) SetInstanceId(v string) *SearchActivationCodeResponseBody {
	s.InstanceId = &v
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
	AppType             *string `json:"appType,omitempty" xml:"appType,omitempty"`
	CreateFromTimeGMT   *string `json:"createFromTimeGMT,omitempty" xml:"createFromTimeGMT,omitempty"`
	CreateToTimeGMT     *string `json:"createToTimeGMT,omitempty" xml:"createToTimeGMT,omitempty"`
	FormUuid            *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	Language            *string `json:"language,omitempty" xml:"language,omitempty"`
	ModifiedFromTimeGMT *string `json:"modifiedFromTimeGMT,omitempty" xml:"modifiedFromTimeGMT,omitempty"`
	ModifiedToTimeGMT   *string `json:"modifiedToTimeGMT,omitempty" xml:"modifiedToTimeGMT,omitempty"`
	OriginatorId        *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	SearchFieldJson     *string `json:"searchFieldJson,omitempty" xml:"searchFieldJson,omitempty"`
	SystemToken         *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	TargetFieldJson     *string `json:"targetFieldJson,omitempty" xml:"targetFieldJson,omitempty"`
	UserId              *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchEmployeeFieldValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchEmployeeFieldValuesRequest) GoString() string {
	return s.String()
}

func (s *SearchEmployeeFieldValuesRequest) SetAppType(v string) *SearchEmployeeFieldValuesRequest {
	s.AppType = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetCreateFromTimeGMT(v string) *SearchEmployeeFieldValuesRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetCreateToTimeGMT(v string) *SearchEmployeeFieldValuesRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetFormUuid(v string) *SearchEmployeeFieldValuesRequest {
	s.FormUuid = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetLanguage(v string) *SearchEmployeeFieldValuesRequest {
	s.Language = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetModifiedFromTimeGMT(v string) *SearchEmployeeFieldValuesRequest {
	s.ModifiedFromTimeGMT = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetModifiedToTimeGMT(v string) *SearchEmployeeFieldValuesRequest {
	s.ModifiedToTimeGMT = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetOriginatorId(v string) *SearchEmployeeFieldValuesRequest {
	s.OriginatorId = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetSearchFieldJson(v string) *SearchEmployeeFieldValuesRequest {
	s.SearchFieldJson = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetSystemToken(v string) *SearchEmployeeFieldValuesRequest {
	s.SystemToken = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetTargetFieldJson(v string) *SearchEmployeeFieldValuesRequest {
	s.TargetFieldJson = &v
	return s
}

func (s *SearchEmployeeFieldValuesRequest) SetUserId(v string) *SearchEmployeeFieldValuesRequest {
	s.UserId = &v
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

type SearchFormDataIdListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchFormDataIdListHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataIdListHeaders) GoString() string {
	return s.String()
}

func (s *SearchFormDataIdListHeaders) SetCommonHeaders(v map[string]*string) *SearchFormDataIdListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchFormDataIdListHeaders) SetXAcsDingtalkAccessToken(v string) *SearchFormDataIdListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchFormDataIdListRequest struct {
	// createFrom和createTo两个时间构造一个时间段。查询在该时间段创建的数据列表
	CreateFromTimeGMT *string `json:"createFromTimeGMT,omitempty" xml:"createFromTimeGMT,omitempty"`
	// createFrom和createTo两个时间构造一个时间段。查询在该时间段创建的数据列表。
	CreateToTimeGMT *string `json:"createToTimeGMT,omitempty" xml:"createToTimeGMT,omitempty"`
	// 语言
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// modifiedFrom和modifiedTo构成一个时间段，查询在该时间段有修改的数据列表
	ModifiedFromTimeGMT *string `json:"modifiedFromTimeGMT,omitempty" xml:"modifiedFromTimeGMT,omitempty"`
	// modifiedFrom和modifiedTo构成一个时间段，查询在该时间段有修改的数据列表。
	ModifiedToTimeGMT *string `json:"modifiedToTimeGMT,omitempty" xml:"modifiedToTimeGMT,omitempty"`
	// 根据数据提交人工号查询
	OriginatorId *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	// 根据表单内组件值查询
	SearchFieldJson *string `json:"searchFieldJson,omitempty" xml:"searchFieldJson,omitempty"`
	// 应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉的userId
	UserId     *string `json:"userId,omitempty" xml:"userId,omitempty"`
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s SearchFormDataIdListRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataIdListRequest) GoString() string {
	return s.String()
}

func (s *SearchFormDataIdListRequest) SetCreateFromTimeGMT(v string) *SearchFormDataIdListRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetCreateToTimeGMT(v string) *SearchFormDataIdListRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetLanguage(v string) *SearchFormDataIdListRequest {
	s.Language = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetModifiedFromTimeGMT(v string) *SearchFormDataIdListRequest {
	s.ModifiedFromTimeGMT = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetModifiedToTimeGMT(v string) *SearchFormDataIdListRequest {
	s.ModifiedToTimeGMT = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetOriginatorId(v string) *SearchFormDataIdListRequest {
	s.OriginatorId = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetSearchFieldJson(v string) *SearchFormDataIdListRequest {
	s.SearchFieldJson = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetSystemToken(v string) *SearchFormDataIdListRequest {
	s.SystemToken = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetUserId(v string) *SearchFormDataIdListRequest {
	s.UserId = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetPageNumber(v int32) *SearchFormDataIdListRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchFormDataIdListRequest) SetPageSize(v int32) *SearchFormDataIdListRequest {
	s.PageSize = &v
	return s
}

type SearchFormDataIdListResponseBody struct {
	// data
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 当前第几页
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 总数量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s SearchFormDataIdListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataIdListResponseBody) GoString() string {
	return s.String()
}

func (s *SearchFormDataIdListResponseBody) SetData(v []*string) *SearchFormDataIdListResponseBody {
	s.Data = v
	return s
}

func (s *SearchFormDataIdListResponseBody) SetPageNumber(v int64) *SearchFormDataIdListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchFormDataIdListResponseBody) SetTotalCount(v int64) *SearchFormDataIdListResponseBody {
	s.TotalCount = &v
	return s
}

type SearchFormDataIdListResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchFormDataIdListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchFormDataIdListResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataIdListResponse) GoString() string {
	return s.String()
}

func (s *SearchFormDataIdListResponse) SetHeaders(v map[string]*string) *SearchFormDataIdListResponse {
	s.Headers = v
	return s
}

func (s *SearchFormDataIdListResponse) SetBody(v *SearchFormDataIdListResponseBody) *SearchFormDataIdListResponse {
	s.Body = v
	return s
}

type SearchFormDataRemovalTableDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchFormDataRemovalTableDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataRemovalTableDataHeaders) GoString() string {
	return s.String()
}

func (s *SearchFormDataRemovalTableDataHeaders) SetCommonHeaders(v map[string]*string) *SearchFormDataRemovalTableDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchFormDataRemovalTableDataHeaders) SetXAcsDingtalkAccessToken(v string) *SearchFormDataRemovalTableDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchFormDataRemovalTableDataRequest struct {
	// 宜搭应用编码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 创建时间起始值
	CreateFromTimeGMT *string `json:"createFromTimeGMT,omitempty" xml:"createFromTimeGMT,omitempty"`
	// 创建时间终止值
	CreateToTimeGMT *string `json:"createToTimeGMT,omitempty" xml:"createToTimeGMT,omitempty"`
	// 表单编码
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 修改时间起始值
	ModifiedFromTimeGMT *string `json:"modifiedFromTimeGMT,omitempty" xml:"modifiedFromTimeGMT,omitempty"`
	// 修改时间终止值
	ModifiedToTimeGMT *string `json:"modifiedToTimeGMT,omitempty" xml:"modifiedToTimeGMT,omitempty"`
	// 排序规则, json格式数据
	OrderConfigJson *string `json:"orderConfigJson,omitempty" xml:"orderConfigJson,omitempty"`
	// 表单提交人的钉钉userId
	OriginatorId *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	// 当前第几页
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页大小
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 用于模糊查询表单实例的查询或精确查询的条件
	SearchFieldJson *string `json:"searchFieldJson,omitempty" xml:"searchFieldJson,omitempty"`
	// 宜搭应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchFormDataRemovalTableDataRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataRemovalTableDataRequest) GoString() string {
	return s.String()
}

func (s *SearchFormDataRemovalTableDataRequest) SetAppType(v string) *SearchFormDataRemovalTableDataRequest {
	s.AppType = &v
	return s
}

func (s *SearchFormDataRemovalTableDataRequest) SetCreateFromTimeGMT(v string) *SearchFormDataRemovalTableDataRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *SearchFormDataRemovalTableDataRequest) SetCreateToTimeGMT(v string) *SearchFormDataRemovalTableDataRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *SearchFormDataRemovalTableDataRequest) SetFormUuid(v string) *SearchFormDataRemovalTableDataRequest {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDataRemovalTableDataRequest) SetModifiedFromTimeGMT(v string) *SearchFormDataRemovalTableDataRequest {
	s.ModifiedFromTimeGMT = &v
	return s
}

func (s *SearchFormDataRemovalTableDataRequest) SetModifiedToTimeGMT(v string) *SearchFormDataRemovalTableDataRequest {
	s.ModifiedToTimeGMT = &v
	return s
}

func (s *SearchFormDataRemovalTableDataRequest) SetOrderConfigJson(v string) *SearchFormDataRemovalTableDataRequest {
	s.OrderConfigJson = &v
	return s
}

func (s *SearchFormDataRemovalTableDataRequest) SetOriginatorId(v string) *SearchFormDataRemovalTableDataRequest {
	s.OriginatorId = &v
	return s
}

func (s *SearchFormDataRemovalTableDataRequest) SetPageNumber(v int32) *SearchFormDataRemovalTableDataRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchFormDataRemovalTableDataRequest) SetPageSize(v int32) *SearchFormDataRemovalTableDataRequest {
	s.PageSize = &v
	return s
}

func (s *SearchFormDataRemovalTableDataRequest) SetSearchFieldJson(v string) *SearchFormDataRemovalTableDataRequest {
	s.SearchFieldJson = &v
	return s
}

func (s *SearchFormDataRemovalTableDataRequest) SetSystemToken(v string) *SearchFormDataRemovalTableDataRequest {
	s.SystemToken = &v
	return s
}

func (s *SearchFormDataRemovalTableDataRequest) SetUserId(v string) *SearchFormDataRemovalTableDataRequest {
	s.UserId = &v
	return s
}

type SearchFormDataRemovalTableDataResponseBody struct {
	// 数据
	Data []*SearchFormDataRemovalTableDataResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 是否还有数据
	HasMoreData *bool `json:"hasMoreData,omitempty" xml:"hasMoreData,omitempty"`
	// 当前第几页
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 总数量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s SearchFormDataRemovalTableDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataRemovalTableDataResponseBody) GoString() string {
	return s.String()
}

func (s *SearchFormDataRemovalTableDataResponseBody) SetData(v []*SearchFormDataRemovalTableDataResponseBodyData) *SearchFormDataRemovalTableDataResponseBody {
	s.Data = v
	return s
}

func (s *SearchFormDataRemovalTableDataResponseBody) SetHasMoreData(v bool) *SearchFormDataRemovalTableDataResponseBody {
	s.HasMoreData = &v
	return s
}

func (s *SearchFormDataRemovalTableDataResponseBody) SetPageNumber(v int64) *SearchFormDataRemovalTableDataResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchFormDataRemovalTableDataResponseBody) SetTotalCount(v int64) *SearchFormDataRemovalTableDataResponseBody {
	s.TotalCount = &v
	return s
}

type SearchFormDataRemovalTableDataResponseBodyData struct {
	// 创建时间
	CreateTimeGMT *string `json:"createTimeGMT,omitempty" xml:"createTimeGMT,omitempty"`
	// 创建者的userId
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// 表单实例数据。结构说明参考  https://www.yuque.com/yida/support/agb8im#jksEx
	FormData map[string]interface{} `json:"formData,omitempty" xml:"formData,omitempty"`
	// 表单实例id
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// 表单编码
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 数据库表记录的主键id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 实例数据
	InstanceValue *string `json:"instanceValue,omitempty" xml:"instanceValue,omitempty"`
	// 修改时间
	ModifiedTimeGMT *string `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	// 修改者的钉钉userId
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// 修改者
	ModifyUser *SearchFormDataRemovalTableDataResponseBodyDataModifyUser `json:"modifyUser,omitempty" xml:"modifyUser,omitempty" type:"Struct"`
	// 发起人
	Originator *SearchFormDataRemovalTableDataResponseBodyDataOriginator `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	// 一次批量导入对应的批次号
	Sequence *string `json:"sequence,omitempty" xml:"sequence,omitempty"`
	// 流水号
	SerialNumber *string `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
	// 标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 该实例对应的表单schema版本
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s SearchFormDataRemovalTableDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataRemovalTableDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchFormDataRemovalTableDataResponseBodyData) SetCreateTimeGMT(v string) *SearchFormDataRemovalTableDataResponseBodyData {
	s.CreateTimeGMT = &v
	return s
}

func (s *SearchFormDataRemovalTableDataResponseBodyData) SetCreatorUserId(v string) *SearchFormDataRemovalTableDataResponseBodyData {
	s.CreatorUserId = &v
	return s
}

func (s *SearchFormDataRemovalTableDataResponseBodyData) SetFormData(v map[string]interface{}) *SearchFormDataRemovalTableDataResponseBodyData {
	s.FormData = v
	return s
}

func (s *SearchFormDataRemovalTableDataResponseBodyData) SetFormInstanceId(v string) *SearchFormDataRemovalTableDataResponseBodyData {
	s.FormInstanceId = &v
	return s
}

func (s *SearchFormDataRemovalTableDataResponseBodyData) SetFormUuid(v string) *SearchFormDataRemovalTableDataResponseBodyData {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDataRemovalTableDataResponseBodyData) SetId(v int64) *SearchFormDataRemovalTableDataResponseBodyData {
	s.Id = &v
	return s
}

func (s *SearchFormDataRemovalTableDataResponseBodyData) SetInstanceValue(v string) *SearchFormDataRemovalTableDataResponseBodyData {
	s.InstanceValue = &v
	return s
}

func (s *SearchFormDataRemovalTableDataResponseBodyData) SetModifiedTimeGMT(v string) *SearchFormDataRemovalTableDataResponseBodyData {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *SearchFormDataRemovalTableDataResponseBodyData) SetModifier(v string) *SearchFormDataRemovalTableDataResponseBodyData {
	s.Modifier = &v
	return s
}

func (s *SearchFormDataRemovalTableDataResponseBodyData) SetModifyUser(v *SearchFormDataRemovalTableDataResponseBodyDataModifyUser) *SearchFormDataRemovalTableDataResponseBodyData {
	s.ModifyUser = v
	return s
}

func (s *SearchFormDataRemovalTableDataResponseBodyData) SetOriginator(v *SearchFormDataRemovalTableDataResponseBodyDataOriginator) *SearchFormDataRemovalTableDataResponseBodyData {
	s.Originator = v
	return s
}

func (s *SearchFormDataRemovalTableDataResponseBodyData) SetSequence(v string) *SearchFormDataRemovalTableDataResponseBodyData {
	s.Sequence = &v
	return s
}

func (s *SearchFormDataRemovalTableDataResponseBodyData) SetSerialNumber(v string) *SearchFormDataRemovalTableDataResponseBodyData {
	s.SerialNumber = &v
	return s
}

func (s *SearchFormDataRemovalTableDataResponseBodyData) SetTitle(v string) *SearchFormDataRemovalTableDataResponseBodyData {
	s.Title = &v
	return s
}

func (s *SearchFormDataRemovalTableDataResponseBodyData) SetVersion(v int64) *SearchFormDataRemovalTableDataResponseBodyData {
	s.Version = &v
	return s
}

type SearchFormDataRemovalTableDataResponseBodyDataModifyUser struct {
	// 部门名称
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 电子邮箱
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// 名称
	Name *SearchFormDataRemovalTableDataResponseBodyDataModifyUserName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchFormDataRemovalTableDataResponseBodyDataModifyUser) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataRemovalTableDataResponseBodyDataModifyUser) GoString() string {
	return s.String()
}

func (s *SearchFormDataRemovalTableDataResponseBodyDataModifyUser) SetDepartmentName(v string) *SearchFormDataRemovalTableDataResponseBodyDataModifyUser {
	s.DepartmentName = &v
	return s
}

func (s *SearchFormDataRemovalTableDataResponseBodyDataModifyUser) SetEmail(v string) *SearchFormDataRemovalTableDataResponseBodyDataModifyUser {
	s.Email = &v
	return s
}

func (s *SearchFormDataRemovalTableDataResponseBodyDataModifyUser) SetName(v *SearchFormDataRemovalTableDataResponseBodyDataModifyUserName) *SearchFormDataRemovalTableDataResponseBodyDataModifyUser {
	s.Name = v
	return s
}

func (s *SearchFormDataRemovalTableDataResponseBodyDataModifyUser) SetUserId(v string) *SearchFormDataRemovalTableDataResponseBodyDataModifyUser {
	s.UserId = &v
	return s
}

type SearchFormDataRemovalTableDataResponseBodyDataModifyUserName struct {
	// 中文名称
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// 英文名称
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
}

func (s SearchFormDataRemovalTableDataResponseBodyDataModifyUserName) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataRemovalTableDataResponseBodyDataModifyUserName) GoString() string {
	return s.String()
}

func (s *SearchFormDataRemovalTableDataResponseBodyDataModifyUserName) SetNameInChinese(v string) *SearchFormDataRemovalTableDataResponseBodyDataModifyUserName {
	s.NameInChinese = &v
	return s
}

func (s *SearchFormDataRemovalTableDataResponseBodyDataModifyUserName) SetNameInEnglish(v string) *SearchFormDataRemovalTableDataResponseBodyDataModifyUserName {
	s.NameInEnglish = &v
	return s
}

type SearchFormDataRemovalTableDataResponseBodyDataOriginator struct {
	// 部门名称
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 电子邮箱
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// 名称
	Name *SearchFormDataRemovalTableDataResponseBodyDataOriginatorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchFormDataRemovalTableDataResponseBodyDataOriginator) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataRemovalTableDataResponseBodyDataOriginator) GoString() string {
	return s.String()
}

func (s *SearchFormDataRemovalTableDataResponseBodyDataOriginator) SetDepartmentName(v string) *SearchFormDataRemovalTableDataResponseBodyDataOriginator {
	s.DepartmentName = &v
	return s
}

func (s *SearchFormDataRemovalTableDataResponseBodyDataOriginator) SetEmail(v string) *SearchFormDataRemovalTableDataResponseBodyDataOriginator {
	s.Email = &v
	return s
}

func (s *SearchFormDataRemovalTableDataResponseBodyDataOriginator) SetName(v *SearchFormDataRemovalTableDataResponseBodyDataOriginatorName) *SearchFormDataRemovalTableDataResponseBodyDataOriginator {
	s.Name = v
	return s
}

func (s *SearchFormDataRemovalTableDataResponseBodyDataOriginator) SetUserId(v string) *SearchFormDataRemovalTableDataResponseBodyDataOriginator {
	s.UserId = &v
	return s
}

type SearchFormDataRemovalTableDataResponseBodyDataOriginatorName struct {
	// 中文名称
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// 英文名称
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
}

func (s SearchFormDataRemovalTableDataResponseBodyDataOriginatorName) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataRemovalTableDataResponseBodyDataOriginatorName) GoString() string {
	return s.String()
}

func (s *SearchFormDataRemovalTableDataResponseBodyDataOriginatorName) SetNameInChinese(v string) *SearchFormDataRemovalTableDataResponseBodyDataOriginatorName {
	s.NameInChinese = &v
	return s
}

func (s *SearchFormDataRemovalTableDataResponseBodyDataOriginatorName) SetNameInEnglish(v string) *SearchFormDataRemovalTableDataResponseBodyDataOriginatorName {
	s.NameInEnglish = &v
	return s
}

type SearchFormDataRemovalTableDataResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchFormDataRemovalTableDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchFormDataRemovalTableDataResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataRemovalTableDataResponse) GoString() string {
	return s.String()
}

func (s *SearchFormDataRemovalTableDataResponse) SetHeaders(v map[string]*string) *SearchFormDataRemovalTableDataResponse {
	s.Headers = v
	return s
}

func (s *SearchFormDataRemovalTableDataResponse) SetBody(v *SearchFormDataRemovalTableDataResponseBody) *SearchFormDataRemovalTableDataResponse {
	s.Body = v
	return s
}

type SearchFormDataSecondGenerationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchFormDataSecondGenerationHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationHeaders) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationHeaders) SetCommonHeaders(v map[string]*string) *SearchFormDataSecondGenerationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchFormDataSecondGenerationHeaders) SetXAcsDingtalkAccessToken(v string) *SearchFormDataSecondGenerationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchFormDataSecondGenerationRequest struct {
	// 宜搭应用编码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 创建时间起始值
	CreateFromTimeGMT *string `json:"createFromTimeGMT,omitempty" xml:"createFromTimeGMT,omitempty"`
	// 创建时间终止值
	CreateToTimeGMT *string `json:"createToTimeGMT,omitempty" xml:"createToTimeGMT,omitempty"`
	// 表单编码
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 修改时间起始值
	ModifiedFromTimeGMT *string `json:"modifiedFromTimeGMT,omitempty" xml:"modifiedFromTimeGMT,omitempty"`
	// 修改时间终止值
	ModifiedToTimeGMT *string `json:"modifiedToTimeGMT,omitempty" xml:"modifiedToTimeGMT,omitempty"`
	// 排序规则
	OrderConfigJson *string `json:"orderConfigJson,omitempty" xml:"orderConfigJson,omitempty"`
	// 表单提交人的钉钉userId
	OriginatorId *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	// 当前第几页
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页大小
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 用于检索表单实例数据的检索条件
	SearchCondition *string `json:"searchCondition,omitempty" xml:"searchCondition,omitempty"`
	// 宜搭应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchFormDataSecondGenerationRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationRequest) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationRequest) SetAppType(v string) *SearchFormDataSecondGenerationRequest {
	s.AppType = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetCreateFromTimeGMT(v string) *SearchFormDataSecondGenerationRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetCreateToTimeGMT(v string) *SearchFormDataSecondGenerationRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetFormUuid(v string) *SearchFormDataSecondGenerationRequest {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetModifiedFromTimeGMT(v string) *SearchFormDataSecondGenerationRequest {
	s.ModifiedFromTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetModifiedToTimeGMT(v string) *SearchFormDataSecondGenerationRequest {
	s.ModifiedToTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetOrderConfigJson(v string) *SearchFormDataSecondGenerationRequest {
	s.OrderConfigJson = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetOriginatorId(v string) *SearchFormDataSecondGenerationRequest {
	s.OriginatorId = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetPageNumber(v int32) *SearchFormDataSecondGenerationRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetPageSize(v int32) *SearchFormDataSecondGenerationRequest {
	s.PageSize = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetSearchCondition(v string) *SearchFormDataSecondGenerationRequest {
	s.SearchCondition = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetSystemToken(v string) *SearchFormDataSecondGenerationRequest {
	s.SystemToken = &v
	return s
}

func (s *SearchFormDataSecondGenerationRequest) SetUserId(v string) *SearchFormDataSecondGenerationRequest {
	s.UserId = &v
	return s
}

type SearchFormDataSecondGenerationResponseBody struct {
	// 表单实例数据
	Data []*SearchFormDataSecondGenerationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 当前第几页
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 总数量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s SearchFormDataSecondGenerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponseBody) SetData(v []*SearchFormDataSecondGenerationResponseBodyData) *SearchFormDataSecondGenerationResponseBody {
	s.Data = v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBody) SetPageNumber(v int64) *SearchFormDataSecondGenerationResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBody) SetTotalCount(v int64) *SearchFormDataSecondGenerationResponseBody {
	s.TotalCount = &v
	return s
}

type SearchFormDataSecondGenerationResponseBodyData struct {
	// 创建时间
	CreateTimeGMT *string `json:"createTimeGMT,omitempty" xml:"createTimeGMT,omitempty"`
	// 创建者的userId
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// 表单实例数据以Map结构展示。结构说明参考  https://www.yuque.com/yida/support/agb8im#jksEx
	FormData map[string]interface{} `json:"formData,omitempty" xml:"formData,omitempty"`
	// 表单实例id
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// 表单编码
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 数据库表记录主键id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 表单实例数据以宜搭组件值格式展示
	InstanceValue *string `json:"instanceValue,omitempty" xml:"instanceValue,omitempty"`
	// 修改时间
	ModifiedTimeGMT *string `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	// 修改者的钉钉userId
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// 修改者
	ModifyUser *SearchFormDataSecondGenerationResponseBodyDataModifyUser `json:"modifyUser,omitempty" xml:"modifyUser,omitempty" type:"Struct"`
	// 发起人
	Originator *SearchFormDataSecondGenerationResponseBodyDataOriginator `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	// 此表单实例所对应的批量导入批次号(如果该表单实例是通过批量导入创建的)
	Sequence *string `json:"sequence,omitempty" xml:"sequence,omitempty"`
	// 流水号
	SerialNumber *string `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
	// 标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 表单实例对应的表单schema版本
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s SearchFormDataSecondGenerationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetCreateTimeGMT(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.CreateTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetCreatorUserId(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.CreatorUserId = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetFormData(v map[string]interface{}) *SearchFormDataSecondGenerationResponseBodyData {
	s.FormData = v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetFormInstanceId(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.FormInstanceId = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetFormUuid(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetId(v int64) *SearchFormDataSecondGenerationResponseBodyData {
	s.Id = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetInstanceValue(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.InstanceValue = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetModifiedTimeGMT(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetModifier(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.Modifier = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetModifyUser(v *SearchFormDataSecondGenerationResponseBodyDataModifyUser) *SearchFormDataSecondGenerationResponseBodyData {
	s.ModifyUser = v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetOriginator(v *SearchFormDataSecondGenerationResponseBodyDataOriginator) *SearchFormDataSecondGenerationResponseBodyData {
	s.Originator = v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetSequence(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.Sequence = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetSerialNumber(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.SerialNumber = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetTitle(v string) *SearchFormDataSecondGenerationResponseBodyData {
	s.Title = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyData) SetVersion(v int64) *SearchFormDataSecondGenerationResponseBodyData {
	s.Version = &v
	return s
}

type SearchFormDataSecondGenerationResponseBodyDataModifyUser struct {
	// 名称
	Name *SearchFormDataSecondGenerationResponseBodyDataModifyUserName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchFormDataSecondGenerationResponseBodyDataModifyUser) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponseBodyDataModifyUser) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponseBodyDataModifyUser) SetName(v *SearchFormDataSecondGenerationResponseBodyDataModifyUserName) *SearchFormDataSecondGenerationResponseBodyDataModifyUser {
	s.Name = v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyDataModifyUser) SetUserId(v string) *SearchFormDataSecondGenerationResponseBodyDataModifyUser {
	s.UserId = &v
	return s
}

type SearchFormDataSecondGenerationResponseBodyDataModifyUserName struct {
	// 中文名称
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// 英文名称
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
}

func (s SearchFormDataSecondGenerationResponseBodyDataModifyUserName) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponseBodyDataModifyUserName) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponseBodyDataModifyUserName) SetNameInChinese(v string) *SearchFormDataSecondGenerationResponseBodyDataModifyUserName {
	s.NameInChinese = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyDataModifyUserName) SetNameInEnglish(v string) *SearchFormDataSecondGenerationResponseBodyDataModifyUserName {
	s.NameInEnglish = &v
	return s
}

type SearchFormDataSecondGenerationResponseBodyDataOriginator struct {
	// 名称
	Name *SearchFormDataSecondGenerationResponseBodyDataOriginatorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchFormDataSecondGenerationResponseBodyDataOriginator) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponseBodyDataOriginator) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponseBodyDataOriginator) SetName(v *SearchFormDataSecondGenerationResponseBodyDataOriginatorName) *SearchFormDataSecondGenerationResponseBodyDataOriginator {
	s.Name = v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyDataOriginator) SetUserId(v string) *SearchFormDataSecondGenerationResponseBodyDataOriginator {
	s.UserId = &v
	return s
}

type SearchFormDataSecondGenerationResponseBodyDataOriginatorName struct {
	// 中文名称
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// 英文名称
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
}

func (s SearchFormDataSecondGenerationResponseBodyDataOriginatorName) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponseBodyDataOriginatorName) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponseBodyDataOriginatorName) SetNameInChinese(v string) *SearchFormDataSecondGenerationResponseBodyDataOriginatorName {
	s.NameInChinese = &v
	return s
}

func (s *SearchFormDataSecondGenerationResponseBodyDataOriginatorName) SetNameInEnglish(v string) *SearchFormDataSecondGenerationResponseBodyDataOriginatorName {
	s.NameInEnglish = &v
	return s
}

type SearchFormDataSecondGenerationResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchFormDataSecondGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchFormDataSecondGenerationResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationResponse) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationResponse) SetHeaders(v map[string]*string) *SearchFormDataSecondGenerationResponse {
	s.Headers = v
	return s
}

func (s *SearchFormDataSecondGenerationResponse) SetBody(v *SearchFormDataSecondGenerationResponseBody) *SearchFormDataSecondGenerationResponse {
	s.Body = v
	return s
}

type SearchFormDataSecondGenerationNoTableFieldHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SearchFormDataSecondGenerationNoTableFieldHeaders) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationNoTableFieldHeaders) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationNoTableFieldHeaders) SetCommonHeaders(v map[string]*string) *SearchFormDataSecondGenerationNoTableFieldHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldHeaders) SetXAcsDingtalkAccessToken(v string) *SearchFormDataSecondGenerationNoTableFieldHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SearchFormDataSecondGenerationNoTableFieldRequest struct {
	// 宜搭应用编码
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 创建时间起始值
	CreateFromTimeGMT *string `json:"createFromTimeGMT,omitempty" xml:"createFromTimeGMT,omitempty"`
	// 创建时间终止值
	CreateToTimeGMT *string `json:"createToTimeGMT,omitempty" xml:"createToTimeGMT,omitempty"`
	// 表单编码
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 修改时间起始值
	ModifiedFromTimeGMT *string `json:"modifiedFromTimeGMT,omitempty" xml:"modifiedFromTimeGMT,omitempty"`
	// 修改时间终止值
	ModifiedToTimeGMT *string `json:"modifiedToTimeGMT,omitempty" xml:"modifiedToTimeGMT,omitempty"`
	// 排序规则
	OrderConfigJson *string `json:"orderConfigJson,omitempty" xml:"orderConfigJson,omitempty"`
	// 表单提交人的钉钉userId
	OriginatorId *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	// 当前第几页
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 分页大小
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 用于检索表单实例数据的检索条件
	SearchCondition *string `json:"searchCondition,omitempty" xml:"searchCondition,omitempty"`
	// 宜搭应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchFormDataSecondGenerationNoTableFieldRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationNoTableFieldRequest) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) SetAppType(v string) *SearchFormDataSecondGenerationNoTableFieldRequest {
	s.AppType = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) SetCreateFromTimeGMT(v string) *SearchFormDataSecondGenerationNoTableFieldRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) SetCreateToTimeGMT(v string) *SearchFormDataSecondGenerationNoTableFieldRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) SetFormUuid(v string) *SearchFormDataSecondGenerationNoTableFieldRequest {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) SetModifiedFromTimeGMT(v string) *SearchFormDataSecondGenerationNoTableFieldRequest {
	s.ModifiedFromTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) SetModifiedToTimeGMT(v string) *SearchFormDataSecondGenerationNoTableFieldRequest {
	s.ModifiedToTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) SetOrderConfigJson(v string) *SearchFormDataSecondGenerationNoTableFieldRequest {
	s.OrderConfigJson = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) SetOriginatorId(v string) *SearchFormDataSecondGenerationNoTableFieldRequest {
	s.OriginatorId = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) SetPageNumber(v int32) *SearchFormDataSecondGenerationNoTableFieldRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) SetPageSize(v int32) *SearchFormDataSecondGenerationNoTableFieldRequest {
	s.PageSize = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) SetSearchCondition(v string) *SearchFormDataSecondGenerationNoTableFieldRequest {
	s.SearchCondition = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) SetSystemToken(v string) *SearchFormDataSecondGenerationNoTableFieldRequest {
	s.SystemToken = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldRequest) SetUserId(v string) *SearchFormDataSecondGenerationNoTableFieldRequest {
	s.UserId = &v
	return s
}

type SearchFormDataSecondGenerationNoTableFieldResponseBody struct {
	// 数据
	Data []*SearchFormDataSecondGenerationNoTableFieldResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 当前第几页
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 总数量
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s SearchFormDataSecondGenerationNoTableFieldResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationNoTableFieldResponseBody) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBody) SetData(v []*SearchFormDataSecondGenerationNoTableFieldResponseBodyData) *SearchFormDataSecondGenerationNoTableFieldResponseBody {
	s.Data = v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBody) SetPageNumber(v int64) *SearchFormDataSecondGenerationNoTableFieldResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBody) SetTotalCount(v int64) *SearchFormDataSecondGenerationNoTableFieldResponseBody {
	s.TotalCount = &v
	return s
}

type SearchFormDataSecondGenerationNoTableFieldResponseBodyData struct {
	// 创建时间
	CreateTimeGMT *string `json:"createTimeGMT,omitempty" xml:"createTimeGMT,omitempty"`
	// 创建者的userId
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// 表单实例数据。结构说明参考 https://www.yuque.com/yida/support/agb8im#jksEx
	FormData map[string]interface{} `json:"formData,omitempty" xml:"formData,omitempty"`
	// 表单实例id
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// 表单编码
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 数据库表记录主键id
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 表单实例数据以组件值格式展示
	InstanceValue *string `json:"instanceValue,omitempty" xml:"instanceValue,omitempty"`
	// 修改时间
	ModifiedTimeGMT *string `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	// 修改者的钉钉userId
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// 修改者
	ModifyUser *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUser `json:"modifyUser,omitempty" xml:"modifyUser,omitempty" type:"Struct"`
	// 表单实例提交人
	Originator *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginator `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	// 此表单实例所对应的批量导入批次号(如果该表单实例是通过批量导入创建的)
	Sequence *string `json:"sequence,omitempty" xml:"sequence,omitempty"`
	// 流水号
	SerialNumber *string `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
	// 标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 该表单实例对应的表单schema版本
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s SearchFormDataSecondGenerationNoTableFieldResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationNoTableFieldResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetCreateTimeGMT(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.CreateTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetCreatorUserId(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.CreatorUserId = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetFormData(v map[string]interface{}) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.FormData = v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetFormInstanceId(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.FormInstanceId = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetFormUuid(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetId(v int64) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.Id = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetInstanceValue(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.InstanceValue = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetModifiedTimeGMT(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetModifier(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.Modifier = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetModifyUser(v *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUser) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.ModifyUser = v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetOriginator(v *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginator) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.Originator = v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetSequence(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.Sequence = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetSerialNumber(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.SerialNumber = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetTitle(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.Title = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyData) SetVersion(v int64) *SearchFormDataSecondGenerationNoTableFieldResponseBodyData {
	s.Version = &v
	return s
}

type SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUser struct {
	// 名称
	Name *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUserName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUser) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUser) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUser) SetName(v *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUserName) *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUser {
	s.Name = v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUser) SetUserId(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUser {
	s.UserId = &v
	return s
}

type SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUserName struct {
	// 中文名称
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// 英文名称
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
}

func (s SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUserName) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUserName) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUserName) SetNameInChinese(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUserName {
	s.NameInChinese = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUserName) SetNameInEnglish(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataModifyUserName {
	s.NameInEnglish = &v
	return s
}

type SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginator struct {
	// 名称
	Name *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginatorName `json:"name,omitempty" xml:"name,omitempty" type:"Struct"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginator) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginator) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginator) SetName(v *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginatorName) *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginator {
	s.Name = v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginator) SetUserId(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginator {
	s.UserId = &v
	return s
}

type SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginatorName struct {
	// 中文名称
	NameInChinese *string `json:"nameInChinese,omitempty" xml:"nameInChinese,omitempty"`
	// 英文名称
	NameInEnglish *string `json:"nameInEnglish,omitempty" xml:"nameInEnglish,omitempty"`
}

func (s SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginatorName) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginatorName) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginatorName) SetNameInChinese(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginatorName {
	s.NameInChinese = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginatorName) SetNameInEnglish(v string) *SearchFormDataSecondGenerationNoTableFieldResponseBodyDataOriginatorName {
	s.NameInEnglish = &v
	return s
}

type SearchFormDataSecondGenerationNoTableFieldResponse struct {
	Headers map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchFormDataSecondGenerationNoTableFieldResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchFormDataSecondGenerationNoTableFieldResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDataSecondGenerationNoTableFieldResponse) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponse) SetHeaders(v map[string]*string) *SearchFormDataSecondGenerationNoTableFieldResponse {
	s.Headers = v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldResponse) SetBody(v *SearchFormDataSecondGenerationNoTableFieldResponseBody) *SearchFormDataSecondGenerationNoTableFieldResponse {
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
	// createFrom和createTo两个时间构造一个时间段。查询在该时间段创建的数据列表, 字符串格式，且为yyyy-MM-DD格式
	CreateFromTimeGMT *string `json:"createFromTimeGMT,omitempty" xml:"createFromTimeGMT,omitempty"`
	// createFrom和createTo两个时间构造一个时间段。查询在该时间段创建的数据列表。字符串格式，且为yyyy-MM-DD格式。 和createFrom一起，相当于查询在 2018-01-01到2018-01-31之间(包含01和31号)创建的数据。
	CreateToTimeGMT *string `json:"createToTimeGMT,omitempty" xml:"createToTimeGMT,omitempty"`
	// 当前页
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// 指定排序字段
	DynamicOrder *string `json:"dynamicOrder,omitempty" xml:"dynamicOrder,omitempty"`
	// 表单ID
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 语言。可选值：zh_CN/en_US 默认：zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// modifiedFrom和modifiedTo构成一个时间段，查询在该时间段有修改的数据列表。字符串格式，且为yyyy-MM-DD格式
	ModifiedFromTimeGMT *string `json:"modifiedFromTimeGMT,omitempty" xml:"modifiedFromTimeGMT,omitempty"`
	// modifiedFrom和modifiedTo构成一个时间段，查询在该时间段有修改的数据列表。字符串格式，且为yyyy-MM-DD格式。 和modifiedFrom一起，相当于查询在 2018-01-01到2018-01-31之间(包含01和31号)被修改的数据。
	ModifiedToTimeGMT *string `json:"modifiedToTimeGMT,omitempty" xml:"modifiedToTimeGMT,omitempty"`
	// 根据数据提交人工号查询
	OriginatorId *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	// 每页记录数
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 根据表单内组件值查询
	SearchFieldJson *string `json:"searchFieldJson,omitempty" xml:"searchFieldJson,omitempty"`
	// 应用秘钥。在应用数据中获取。
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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

func (s *SearchFormDatasRequest) SetCreateFromTimeGMT(v string) *SearchFormDatasRequest {
	s.CreateFromTimeGMT = &v
	return s
}

func (s *SearchFormDatasRequest) SetCreateToTimeGMT(v string) *SearchFormDatasRequest {
	s.CreateToTimeGMT = &v
	return s
}

func (s *SearchFormDatasRequest) SetCurrentPage(v int32) *SearchFormDatasRequest {
	s.CurrentPage = &v
	return s
}

func (s *SearchFormDatasRequest) SetDynamicOrder(v string) *SearchFormDatasRequest {
	s.DynamicOrder = &v
	return s
}

func (s *SearchFormDatasRequest) SetFormUuid(v string) *SearchFormDatasRequest {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDatasRequest) SetLanguage(v string) *SearchFormDatasRequest {
	s.Language = &v
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

func (s *SearchFormDatasRequest) SetOriginatorId(v string) *SearchFormDatasRequest {
	s.OriginatorId = &v
	return s
}

func (s *SearchFormDatasRequest) SetPageSize(v int32) *SearchFormDatasRequest {
	s.PageSize = &v
	return s
}

func (s *SearchFormDatasRequest) SetSearchFieldJson(v string) *SearchFormDatasRequest {
	s.SearchFieldJson = &v
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

type SearchFormDatasResponseBody struct {
	// 当前页
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// 实例详情列表
	Data []*SearchFormDatasResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 符合条件的实例总数
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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

func (s *SearchFormDatasResponseBody) SetData(v []*SearchFormDatasResponseBodyData) *SearchFormDatasResponseBody {
	s.Data = v
	return s
}

func (s *SearchFormDatasResponseBody) SetTotalCount(v int64) *SearchFormDatasResponseBody {
	s.TotalCount = &v
	return s
}

type SearchFormDatasResponseBodyData struct {
	// 数据创建时间
	CreatedTimeGMT *string `json:"createdTimeGMT,omitempty" xml:"createdTimeGMT,omitempty"`
	// 创建人
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// 实体主键id
	DataId *int64 `json:"dataId,omitempty" xml:"dataId,omitempty"`
	// 表单数据
	FormData map[string]interface{} `json:"formData,omitempty" xml:"formData,omitempty"`
	// 表单实例ID
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// 表单id
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 表单实例原始格式值
	InstanceValue *string `json:"instanceValue,omitempty" xml:"instanceValue,omitempty"`
	// 模型id
	ModelUuid *string `json:"modelUuid,omitempty" xml:"modelUuid,omitempty"`
	// 最近修改时间
	ModifiedTimeGMT *string `json:"modifiedTimeGMT,omitempty" xml:"modifiedTimeGMT,omitempty"`
	// 修改人
	ModifierUserId *string `json:"modifierUserId,omitempty" xml:"modifierUserId,omitempty"`
	// 修改者
	ModifyUser *SearchFormDatasResponseBodyDataModifyUser `json:"modifyUser,omitempty" xml:"modifyUser,omitempty" type:"Struct"`
	// 发起人
	Originator *SearchFormDatasResponseBodyDataOriginator `json:"originator,omitempty" xml:"originator,omitempty" type:"Struct"`
	// 批次号
	Sequence *string `json:"sequence,omitempty" xml:"sequence,omitempty"`
	// 流水号
	SerialNo *string `json:"serialNo,omitempty" xml:"serialNo,omitempty"`
	// 标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 数据版本
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s SearchFormDatasResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBodyData) SetCreatedTimeGMT(v string) *SearchFormDatasResponseBodyData {
	s.CreatedTimeGMT = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetCreatorUserId(v string) *SearchFormDatasResponseBodyData {
	s.CreatorUserId = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetDataId(v int64) *SearchFormDatasResponseBodyData {
	s.DataId = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetFormData(v map[string]interface{}) *SearchFormDatasResponseBodyData {
	s.FormData = v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetFormInstanceId(v string) *SearchFormDatasResponseBodyData {
	s.FormInstanceId = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetFormUuid(v string) *SearchFormDatasResponseBodyData {
	s.FormUuid = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetInstanceValue(v string) *SearchFormDatasResponseBodyData {
	s.InstanceValue = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetModelUuid(v string) *SearchFormDatasResponseBodyData {
	s.ModelUuid = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetModifiedTimeGMT(v string) *SearchFormDatasResponseBodyData {
	s.ModifiedTimeGMT = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetModifierUserId(v string) *SearchFormDatasResponseBodyData {
	s.ModifierUserId = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetModifyUser(v *SearchFormDatasResponseBodyDataModifyUser) *SearchFormDatasResponseBodyData {
	s.ModifyUser = v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetOriginator(v *SearchFormDatasResponseBodyDataOriginator) *SearchFormDatasResponseBodyData {
	s.Originator = v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetSequence(v string) *SearchFormDatasResponseBodyData {
	s.Sequence = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetSerialNo(v string) *SearchFormDatasResponseBodyData {
	s.SerialNo = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetTitle(v string) *SearchFormDatasResponseBodyData {
	s.Title = &v
	return s
}

func (s *SearchFormDatasResponseBodyData) SetVersion(v int64) *SearchFormDatasResponseBodyData {
	s.Version = &v
	return s
}

type SearchFormDatasResponseBodyDataModifyUser struct {
	// 部门名称
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 邮箱
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// 用户工号
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用户名
	UserName *SearchFormDatasResponseBodyDataModifyUserUserName `json:"userName,omitempty" xml:"userName,omitempty" type:"Struct"`
}

func (s SearchFormDatasResponseBodyDataModifyUser) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponseBodyDataModifyUser) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBodyDataModifyUser) SetDepartmentName(v string) *SearchFormDatasResponseBodyDataModifyUser {
	s.DepartmentName = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataModifyUser) SetEmail(v string) *SearchFormDatasResponseBodyDataModifyUser {
	s.Email = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataModifyUser) SetUserId(v string) *SearchFormDatasResponseBodyDataModifyUser {
	s.UserId = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataModifyUser) SetUserName(v *SearchFormDatasResponseBodyDataModifyUserUserName) *SearchFormDatasResponseBodyDataModifyUser {
	s.UserName = v
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

type SearchFormDatasResponseBodyDataOriginator struct {
	// 部门名称
	DepartmentName *string `json:"departmentName,omitempty" xml:"departmentName,omitempty"`
	// 邮箱
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// 用户工号
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 用户名
	UserName *SearchFormDatasResponseBodyDataOriginatorUserName `json:"userName,omitempty" xml:"userName,omitempty" type:"Struct"`
}

func (s SearchFormDatasResponseBodyDataOriginator) String() string {
	return tea.Prettify(s)
}

func (s SearchFormDatasResponseBodyDataOriginator) GoString() string {
	return s.String()
}

func (s *SearchFormDatasResponseBodyDataOriginator) SetDepartmentName(v string) *SearchFormDatasResponseBodyDataOriginator {
	s.DepartmentName = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataOriginator) SetEmail(v string) *SearchFormDatasResponseBodyDataOriginator {
	s.Email = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataOriginator) SetUserId(v string) *SearchFormDatasResponseBodyDataOriginator {
	s.UserId = &v
	return s
}

func (s *SearchFormDatasResponseBodyDataOriginator) SetUserName(v *SearchFormDatasResponseBodyDataOriginatorUserName) *SearchFormDatasResponseBodyDataOriginator {
	s.UserName = v
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
	// 发起人所在部门编号
	DepartmentId *string `json:"departmentId,omitempty" xml:"departmentId,omitempty"`
	// 表单数据
	FormDataJson *string `json:"formDataJson,omitempty" xml:"formDataJson,omitempty"`
	// 表单唯一编码
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 语言。可选值：zh_CN/en_US 默认：zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 流程编码
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// 应用秘钥。在应用数据中获取。
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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

func (s *StartInstanceRequest) SetDepartmentId(v string) *StartInstanceRequest {
	s.DepartmentId = &v
	return s
}

func (s *StartInstanceRequest) SetFormDataJson(v string) *StartInstanceRequest {
	s.FormDataJson = &v
	return s
}

func (s *StartInstanceRequest) SetFormUuid(v string) *StartInstanceRequest {
	s.FormUuid = &v
	return s
}

func (s *StartInstanceRequest) SetLanguage(v string) *StartInstanceRequest {
	s.Language = &v
	return s
}

func (s *StartInstanceRequest) SetProcessCode(v string) *StartInstanceRequest {
	s.ProcessCode = &v
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

type TerminateCloudAuthorizationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TerminateCloudAuthorizationHeaders) String() string {
	return tea.Prettify(s)
}

func (s TerminateCloudAuthorizationHeaders) GoString() string {
	return s.String()
}

func (s *TerminateCloudAuthorizationHeaders) SetCommonHeaders(v map[string]*string) *TerminateCloudAuthorizationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TerminateCloudAuthorizationHeaders) SetXAcsDingtalkAccessToken(v string) *TerminateCloudAuthorizationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TerminateCloudAuthorizationRequest struct {
	// 访问秘钥
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// 调用者unionId
	CallerUnionId *string `json:"callerUnionId,omitempty" xml:"callerUnionId,omitempty"`
	// 实例id
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s TerminateCloudAuthorizationRequest) String() string {
	return tea.Prettify(s)
}

func (s TerminateCloudAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *TerminateCloudAuthorizationRequest) SetAccessKey(v string) *TerminateCloudAuthorizationRequest {
	s.AccessKey = &v
	return s
}

func (s *TerminateCloudAuthorizationRequest) SetCallerUnionId(v string) *TerminateCloudAuthorizationRequest {
	s.CallerUnionId = &v
	return s
}

func (s *TerminateCloudAuthorizationRequest) SetInstanceId(v string) *TerminateCloudAuthorizationRequest {
	s.InstanceId = &v
	return s
}

type TerminateCloudAuthorizationResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s TerminateCloudAuthorizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TerminateCloudAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *TerminateCloudAuthorizationResponseBody) SetResult(v bool) *TerminateCloudAuthorizationResponseBody {
	s.Result = &v
	return s
}

type TerminateCloudAuthorizationResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TerminateCloudAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TerminateCloudAuthorizationResponse) String() string {
	return tea.Prettify(s)
}

func (s TerminateCloudAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *TerminateCloudAuthorizationResponse) SetHeaders(v map[string]*string) *TerminateCloudAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *TerminateCloudAuthorizationResponse) SetBody(v *TerminateCloudAuthorizationResponseBody) *TerminateCloudAuthorizationResponse {
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
	// 语言
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 流程实例ID
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// 应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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

func (s *TerminateInstanceRequest) SetLanguage(v string) *TerminateInstanceRequest {
	s.Language = &v
	return s
}

func (s *TerminateInstanceRequest) SetProcessInstanceId(v string) *TerminateInstanceRequest {
	s.ProcessInstanceId = &v
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

type UpdateCloudAccountInformationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateCloudAccountInformationHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateCloudAccountInformationHeaders) GoString() string {
	return s.String()
}

func (s *UpdateCloudAccountInformationHeaders) SetCommonHeaders(v map[string]*string) *UpdateCloudAccountInformationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateCloudAccountInformationHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateCloudAccountInformationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateCloudAccountInformationRequest struct {
	// 访问秘钥
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// 账户号
	AccountNumber *string `json:"accountNumber,omitempty" xml:"accountNumber,omitempty"`
	// 调用者unionId
	CallerUnionId *string `json:"callerUnionId,omitempty" xml:"callerUnionId,omitempty"`
	// 商品类型
	CommodityType *string `json:"commodityType,omitempty" xml:"commodityType,omitempty"`
}

func (s UpdateCloudAccountInformationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCloudAccountInformationRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudAccountInformationRequest) SetAccessKey(v string) *UpdateCloudAccountInformationRequest {
	s.AccessKey = &v
	return s
}

func (s *UpdateCloudAccountInformationRequest) SetAccountNumber(v string) *UpdateCloudAccountInformationRequest {
	s.AccountNumber = &v
	return s
}

func (s *UpdateCloudAccountInformationRequest) SetCallerUnionId(v string) *UpdateCloudAccountInformationRequest {
	s.CallerUnionId = &v
	return s
}

func (s *UpdateCloudAccountInformationRequest) SetCommodityType(v string) *UpdateCloudAccountInformationRequest {
	s.CommodityType = &v
	return s
}

type UpdateCloudAccountInformationResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateCloudAccountInformationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCloudAccountInformationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloudAccountInformationResponseBody) SetResult(v bool) *UpdateCloudAccountInformationResponseBody {
	s.Result = &v
	return s
}

type UpdateCloudAccountInformationResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCloudAccountInformationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCloudAccountInformationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCloudAccountInformationResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloudAccountInformationResponse) SetHeaders(v map[string]*string) *UpdateCloudAccountInformationResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloudAccountInformationResponse) SetBody(v *UpdateCloudAccountInformationResponseBody) *UpdateCloudAccountInformationResponse {
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
	// 要更新的表单数据ID
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// 语言。可选值：zh_CN/en_US 默认：zh_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 应用秘钥。在应用数据中获取。
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 要更新的表单组件值。参数有的组件更新，没有的组件保持不变。 明细的值只能统一更新，无法只更新明细下某个组件的值
	UpdateFormDataJson *string `json:"updateFormDataJson,omitempty" xml:"updateFormDataJson,omitempty"`
	// 使用最新的表单版本进行更新。默认为false
	UseLatestVersion *bool `json:"useLatestVersion,omitempty" xml:"useLatestVersion,omitempty"`
	// 钉钉userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
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

func (s *UpdateFormDataRequest) SetFormInstanceId(v string) *UpdateFormDataRequest {
	s.FormInstanceId = &v
	return s
}

func (s *UpdateFormDataRequest) SetLanguage(v string) *UpdateFormDataRequest {
	s.Language = &v
	return s
}

func (s *UpdateFormDataRequest) SetSystemToken(v string) *UpdateFormDataRequest {
	s.SystemToken = &v
	return s
}

func (s *UpdateFormDataRequest) SetUpdateFormDataJson(v string) *UpdateFormDataRequest {
	s.UpdateFormDataJson = &v
	return s
}

func (s *UpdateFormDataRequest) SetUseLatestVersion(v bool) *UpdateFormDataRequest {
	s.UseLatestVersion = &v
	return s
}

func (s *UpdateFormDataRequest) SetUserId(v string) *UpdateFormDataRequest {
	s.UserId = &v
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
	// 应用ID
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// 语言环境
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// 实例ID
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// 应用秘钥
	SystemToken *string `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	// 更新的表单数据
	UpdateFormDataJson *string `json:"updateFormDataJson,omitempty" xml:"updateFormDataJson,omitempty"`
	// 钉钉的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) SetAppType(v string) *UpdateInstanceRequest {
	s.AppType = &v
	return s
}

func (s *UpdateInstanceRequest) SetLanguage(v string) *UpdateInstanceRequest {
	s.Language = &v
	return s
}

func (s *UpdateInstanceRequest) SetProcessInstanceId(v string) *UpdateInstanceRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *UpdateInstanceRequest) SetSystemToken(v string) *UpdateInstanceRequest {
	s.SystemToken = &v
	return s
}

func (s *UpdateInstanceRequest) SetUpdateFormDataJson(v string) *UpdateInstanceRequest {
	s.UpdateFormDataJson = &v
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
	AppType        *string  `json:"appType,omitempty" xml:"appType,omitempty"`
	ErrorLines     []*int32 `json:"errorLines,omitempty" xml:"errorLines,omitempty" type:"Repeated"`
	ImportSequence *string  `json:"importSequence,omitempty" xml:"importSequence,omitempty"`
	Language       *string  `json:"language,omitempty" xml:"language,omitempty"`
	Status         *string  `json:"status,omitempty" xml:"status,omitempty"`
	SystemToken    *string  `json:"systemToken,omitempty" xml:"systemToken,omitempty"`
	UserId         *string  `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateStatusRequest) SetAppType(v string) *UpdateStatusRequest {
	s.AppType = &v
	return s
}

func (s *UpdateStatusRequest) SetErrorLines(v []*int32) *UpdateStatusRequest {
	s.ErrorLines = v
	return s
}

func (s *UpdateStatusRequest) SetImportSequence(v string) *UpdateStatusRequest {
	s.ImportSequence = &v
	return s
}

func (s *UpdateStatusRequest) SetLanguage(v string) *UpdateStatusRequest {
	s.Language = &v
	return s
}

func (s *UpdateStatusRequest) SetStatus(v string) *UpdateStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateStatusRequest) SetSystemToken(v string) *UpdateStatusRequest {
	s.SystemToken = &v
	return s
}

func (s *UpdateStatusRequest) SetUserId(v string) *UpdateStatusRequest {
	s.UserId = &v
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

type UpgradeTenantInformationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpgradeTenantInformationHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpgradeTenantInformationHeaders) GoString() string {
	return s.String()
}

func (s *UpgradeTenantInformationHeaders) SetCommonHeaders(v map[string]*string) *UpgradeTenantInformationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpgradeTenantInformationHeaders) SetXAcsDingtalkAccessToken(v string) *UpgradeTenantInformationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpgradeTenantInformationRequest struct {
	// 访问秘钥
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// 账户号
	AccountNumber *string `json:"accountNumber,omitempty" xml:"accountNumber,omitempty"`
	// 调用者unionId
	CallerUnionId *string `json:"callerUnionId,omitempty" xml:"callerUnionId,omitempty"`
	// 商品类型
	CommodityType *string `json:"commodityType,omitempty" xml:"commodityType,omitempty"`
}

func (s UpgradeTenantInformationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeTenantInformationRequest) GoString() string {
	return s.String()
}

func (s *UpgradeTenantInformationRequest) SetAccessKey(v string) *UpgradeTenantInformationRequest {
	s.AccessKey = &v
	return s
}

func (s *UpgradeTenantInformationRequest) SetAccountNumber(v string) *UpgradeTenantInformationRequest {
	s.AccountNumber = &v
	return s
}

func (s *UpgradeTenantInformationRequest) SetCallerUnionId(v string) *UpgradeTenantInformationRequest {
	s.CallerUnionId = &v
	return s
}

func (s *UpgradeTenantInformationRequest) SetCommodityType(v string) *UpgradeTenantInformationRequest {
	s.CommodityType = &v
	return s
}

type UpgradeTenantInformationResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpgradeTenantInformationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeTenantInformationResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeTenantInformationResponseBody) SetResult(v bool) *UpgradeTenantInformationResponseBody {
	s.Result = &v
	return s
}

type UpgradeTenantInformationResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeTenantInformationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeTenantInformationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeTenantInformationResponse) GoString() string {
	return s.String()
}

func (s *UpgradeTenantInformationResponse) SetHeaders(v map[string]*string) *UpgradeTenantInformationResponse {
	s.Headers = v
	return s
}

func (s *UpgradeTenantInformationResponse) SetBody(v *UpgradeTenantInformationResponseBody) *UpgradeTenantInformationResponse {
	s.Body = v
	return s
}

type ValidateApplicationAuthorizationOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ValidateApplicationAuthorizationOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s ValidateApplicationAuthorizationOrderHeaders) GoString() string {
	return s.String()
}

func (s *ValidateApplicationAuthorizationOrderHeaders) SetCommonHeaders(v map[string]*string) *ValidateApplicationAuthorizationOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ValidateApplicationAuthorizationOrderHeaders) SetXAcsDingtalkAccessToken(v string) *ValidateApplicationAuthorizationOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ValidateApplicationAuthorizationOrderRequest struct {
	AccessKey     *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	CallerUnionId *string `json:"callerUnionId,omitempty" xml:"callerUnionId,omitempty"`
}

func (s ValidateApplicationAuthorizationOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateApplicationAuthorizationOrderRequest) GoString() string {
	return s.String()
}

func (s *ValidateApplicationAuthorizationOrderRequest) SetAccessKey(v string) *ValidateApplicationAuthorizationOrderRequest {
	s.AccessKey = &v
	return s
}

func (s *ValidateApplicationAuthorizationOrderRequest) SetCallerUnionId(v string) *ValidateApplicationAuthorizationOrderRequest {
	s.CallerUnionId = &v
	return s
}

type ValidateApplicationAuthorizationOrderResponseBody struct {
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// status
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ValidateApplicationAuthorizationOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateApplicationAuthorizationOrderResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateApplicationAuthorizationOrderResponseBody) SetMessage(v string) *ValidateApplicationAuthorizationOrderResponseBody {
	s.Message = &v
	return s
}

func (s *ValidateApplicationAuthorizationOrderResponseBody) SetStatus(v int32) *ValidateApplicationAuthorizationOrderResponseBody {
	s.Status = &v
	return s
}

type ValidateApplicationAuthorizationOrderResponse struct {
	Headers map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ValidateApplicationAuthorizationOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateApplicationAuthorizationOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateApplicationAuthorizationOrderResponse) GoString() string {
	return s.String()
}

func (s *ValidateApplicationAuthorizationOrderResponse) SetHeaders(v map[string]*string) *ValidateApplicationAuthorizationOrderResponse {
	s.Headers = v
	return s
}

func (s *ValidateApplicationAuthorizationOrderResponse) SetBody(v *ValidateApplicationAuthorizationOrderResponseBody) *ValidateApplicationAuthorizationOrderResponse {
	s.Body = v
	return s
}

type ValidateApplicationAuthorizationServiceOrderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ValidateApplicationAuthorizationServiceOrderHeaders) String() string {
	return tea.Prettify(s)
}

func (s ValidateApplicationAuthorizationServiceOrderHeaders) GoString() string {
	return s.String()
}

func (s *ValidateApplicationAuthorizationServiceOrderHeaders) SetCommonHeaders(v map[string]*string) *ValidateApplicationAuthorizationServiceOrderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ValidateApplicationAuthorizationServiceOrderHeaders) SetXAcsDingtalkAccessToken(v string) *ValidateApplicationAuthorizationServiceOrderHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ValidateApplicationAuthorizationServiceOrderRequest struct {
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
}

func (s ValidateApplicationAuthorizationServiceOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateApplicationAuthorizationServiceOrderRequest) GoString() string {
	return s.String()
}

func (s *ValidateApplicationAuthorizationServiceOrderRequest) SetAccessKey(v string) *ValidateApplicationAuthorizationServiceOrderRequest {
	s.AccessKey = &v
	return s
}

type ValidateApplicationAuthorizationServiceOrderResponseBody struct {
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// status
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ValidateApplicationAuthorizationServiceOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateApplicationAuthorizationServiceOrderResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateApplicationAuthorizationServiceOrderResponseBody) SetMessage(v string) *ValidateApplicationAuthorizationServiceOrderResponseBody {
	s.Message = &v
	return s
}

func (s *ValidateApplicationAuthorizationServiceOrderResponseBody) SetStatus(v int32) *ValidateApplicationAuthorizationServiceOrderResponseBody {
	s.Status = &v
	return s
}

type ValidateApplicationAuthorizationServiceOrderResponse struct {
	Headers map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ValidateApplicationAuthorizationServiceOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateApplicationAuthorizationServiceOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateApplicationAuthorizationServiceOrderResponse) GoString() string {
	return s.String()
}

func (s *ValidateApplicationAuthorizationServiceOrderResponse) SetHeaders(v map[string]*string) *ValidateApplicationAuthorizationServiceOrderResponse {
	s.Headers = v
	return s
}

func (s *ValidateApplicationAuthorizationServiceOrderResponse) SetBody(v *ValidateApplicationAuthorizationServiceOrderResponseBody) *ValidateApplicationAuthorizationServiceOrderResponse {
	s.Body = v
	return s
}

type ValidateApplicationServiceOrderUpgradeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ValidateApplicationServiceOrderUpgradeHeaders) String() string {
	return tea.Prettify(s)
}

func (s ValidateApplicationServiceOrderUpgradeHeaders) GoString() string {
	return s.String()
}

func (s *ValidateApplicationServiceOrderUpgradeHeaders) SetCommonHeaders(v map[string]*string) *ValidateApplicationServiceOrderUpgradeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ValidateApplicationServiceOrderUpgradeHeaders) SetXAcsDingtalkAccessToken(v string) *ValidateApplicationServiceOrderUpgradeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ValidateApplicationServiceOrderUpgradeRequest struct {
	// accessKey
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
}

func (s ValidateApplicationServiceOrderUpgradeRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateApplicationServiceOrderUpgradeRequest) GoString() string {
	return s.String()
}

func (s *ValidateApplicationServiceOrderUpgradeRequest) SetAccessKey(v string) *ValidateApplicationServiceOrderUpgradeRequest {
	s.AccessKey = &v
	return s
}

type ValidateApplicationServiceOrderUpgradeResponseBody struct {
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// status
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ValidateApplicationServiceOrderUpgradeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateApplicationServiceOrderUpgradeResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateApplicationServiceOrderUpgradeResponseBody) SetMessage(v string) *ValidateApplicationServiceOrderUpgradeResponseBody {
	s.Message = &v
	return s
}

func (s *ValidateApplicationServiceOrderUpgradeResponseBody) SetStatus(v int32) *ValidateApplicationServiceOrderUpgradeResponseBody {
	s.Status = &v
	return s
}

type ValidateApplicationServiceOrderUpgradeResponse struct {
	Headers map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ValidateApplicationServiceOrderUpgradeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateApplicationServiceOrderUpgradeResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateApplicationServiceOrderUpgradeResponse) GoString() string {
	return s.String()
}

func (s *ValidateApplicationServiceOrderUpgradeResponse) SetHeaders(v map[string]*string) *ValidateApplicationServiceOrderUpgradeResponse {
	s.Headers = v
	return s
}

func (s *ValidateApplicationServiceOrderUpgradeResponse) SetBody(v *ValidateApplicationServiceOrderUpgradeResponseBody) *ValidateApplicationServiceOrderUpgradeResponse {
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

type ValidateOrderUpdateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ValidateOrderUpdateHeaders) String() string {
	return tea.Prettify(s)
}

func (s ValidateOrderUpdateHeaders) GoString() string {
	return s.String()
}

func (s *ValidateOrderUpdateHeaders) SetCommonHeaders(v map[string]*string) *ValidateOrderUpdateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ValidateOrderUpdateHeaders) SetXAcsDingtalkAccessToken(v string) *ValidateOrderUpdateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ValidateOrderUpdateRequest struct {
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	CallerUid *string `json:"callerUid,omitempty" xml:"callerUid,omitempty"`
}

func (s ValidateOrderUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateOrderUpdateRequest) GoString() string {
	return s.String()
}

func (s *ValidateOrderUpdateRequest) SetAccessKey(v string) *ValidateOrderUpdateRequest {
	s.AccessKey = &v
	return s
}

func (s *ValidateOrderUpdateRequest) SetCallerUid(v string) *ValidateOrderUpdateRequest {
	s.CallerUid = &v
	return s
}

type ValidateOrderUpdateResponseBody struct {
	// message
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// status
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ValidateOrderUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateOrderUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateOrderUpdateResponseBody) SetMessage(v string) *ValidateOrderUpdateResponseBody {
	s.Message = &v
	return s
}

func (s *ValidateOrderUpdateResponseBody) SetStatus(v int32) *ValidateOrderUpdateResponseBody {
	s.Status = &v
	return s
}

type ValidateOrderUpdateResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ValidateOrderUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateOrderUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateOrderUpdateResponse) GoString() string {
	return s.String()
}

func (s *ValidateOrderUpdateResponse) SetHeaders(v map[string]*string) *ValidateOrderUpdateResponse {
	s.Headers = v
	return s
}

func (s *ValidateOrderUpdateResponse) SetBody(v *ValidateOrderUpdateResponseBody) *ValidateOrderUpdateResponse {
	s.Body = v
	return s
}

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
	AccessKey  *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	CallerUid  *string `json:"callerUid,omitempty" xml:"callerUid,omitempty"`
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s ValidateOrderUpgradeRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateOrderUpgradeRequest) GoString() string {
	return s.String()
}

func (s *ValidateOrderUpgradeRequest) SetAccessKey(v string) *ValidateOrderUpgradeRequest {
	s.AccessKey = &v
	return s
}

func (s *ValidateOrderUpgradeRequest) SetCallerUid(v string) *ValidateOrderUpgradeRequest {
	s.CallerUid = &v
	return s
}

func (s *ValidateOrderUpgradeRequest) SetInstanceId(v string) *ValidateOrderUpgradeRequest {
	s.InstanceId = &v
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

func (client *Client) BatchGetFormDataByIdList(request *BatchGetFormDataByIdListRequest) (_result *BatchGetFormDataByIdListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchGetFormDataByIdListHeaders{}
	_result = &BatchGetFormDataByIdListResponse{}
	_body, _err := client.BatchGetFormDataByIdListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchGetFormDataByIdListWithOptions(request *BatchGetFormDataByIdListRequest, headers *BatchGetFormDataByIdListHeaders, runtime *util.RuntimeOptions) (_result *BatchGetFormDataByIdListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.FormInstanceIdList)) {
		body["formInstanceIdList"] = request.FormInstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.NeedFormInstanceValue)) {
		body["needFormInstanceValue"] = request.NeedFormInstanceValue
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
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
	_result = &BatchGetFormDataByIdListResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchGetFormDataByIdList"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/forms/instances/ids/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchRemovalByFormInstanceIdList(request *BatchRemovalByFormInstanceIdListRequest) (_result *BatchRemovalByFormInstanceIdListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchRemovalByFormInstanceIdListHeaders{}
	_result = &BatchRemovalByFormInstanceIdListResponse{}
	_body, _err := client.BatchRemovalByFormInstanceIdListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchRemovalByFormInstanceIdListWithOptions(request *BatchRemovalByFormInstanceIdListRequest, headers *BatchRemovalByFormInstanceIdListHeaders, runtime *util.RuntimeOptions) (_result *BatchRemovalByFormInstanceIdListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.AsynchronousExecution)) {
		body["asynchronousExecution"] = request.AsynchronousExecution
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteExpression)) {
		body["executeExpression"] = request.ExecuteExpression
	}

	if !tea.BoolValue(util.IsUnset(request.FormInstanceIdList)) {
		body["formInstanceIdList"] = request.FormInstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
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
	_result = &BatchRemovalByFormInstanceIdListResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchRemovalByFormInstanceIdList"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/forms/instances/batchRemove"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchSaveFormData(request *BatchSaveFormDataRequest) (_result *BatchSaveFormDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchSaveFormDataHeaders{}
	_result = &BatchSaveFormDataResponse{}
	_body, _err := client.BatchSaveFormDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchSaveFormDataWithOptions(request *BatchSaveFormDataRequest, headers *BatchSaveFormDataHeaders, runtime *util.RuntimeOptions) (_result *BatchSaveFormDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.AsynchronousExecution)) {
		body["asynchronousExecution"] = request.AsynchronousExecution
	}

	if !tea.BoolValue(util.IsUnset(request.FormDataJsonList)) {
		body["formDataJsonList"] = request.FormDataJsonList
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.KeepRunningAfterException)) {
		body["keepRunningAfterException"] = request.KeepRunningAfterException
	}

	if !tea.BoolValue(util.IsUnset(request.NoExecuteExpression)) {
		body["noExecuteExpression"] = request.NoExecuteExpression
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
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
	_result = &BatchSaveFormDataResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchSaveFormData"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/forms/instances/batchSave"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchUpdateFormDataByInstanceId(request *BatchUpdateFormDataByInstanceIdRequest) (_result *BatchUpdateFormDataByInstanceIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchUpdateFormDataByInstanceIdHeaders{}
	_result = &BatchUpdateFormDataByInstanceIdResponse{}
	_body, _err := client.BatchUpdateFormDataByInstanceIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchUpdateFormDataByInstanceIdWithOptions(request *BatchUpdateFormDataByInstanceIdRequest, headers *BatchUpdateFormDataByInstanceIdHeaders, runtime *util.RuntimeOptions) (_result *BatchUpdateFormDataByInstanceIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.AsynchronousExecution)) {
		body["asynchronousExecution"] = request.AsynchronousExecution
	}

	if !tea.BoolValue(util.IsUnset(request.FormInstanceIdList)) {
		body["formInstanceIdList"] = request.FormInstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.IgnoreEmpty)) {
		body["ignoreEmpty"] = request.IgnoreEmpty
	}

	if !tea.BoolValue(util.IsUnset(request.NoExecuteExpression)) {
		body["noExecuteExpression"] = request.NoExecuteExpression
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateFormDataJson)) {
		body["updateFormDataJson"] = request.UpdateFormDataJson
	}

	if !tea.BoolValue(util.IsUnset(request.UseLatestFormSchemaVersion)) {
		body["useLatestFormSchemaVersion"] = request.UseLatestFormSchemaVersion
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
	_result = &BatchUpdateFormDataByInstanceIdResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchUpdateFormDataByInstanceId"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/yida/forms/instances/components"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchUpdateFormDataByInstanceMap(request *BatchUpdateFormDataByInstanceMapRequest) (_result *BatchUpdateFormDataByInstanceMapResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchUpdateFormDataByInstanceMapHeaders{}
	_result = &BatchUpdateFormDataByInstanceMapResponse{}
	_body, _err := client.BatchUpdateFormDataByInstanceMapWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchUpdateFormDataByInstanceMapWithOptions(request *BatchUpdateFormDataByInstanceMapRequest, headers *BatchUpdateFormDataByInstanceMapHeaders, runtime *util.RuntimeOptions) (_result *BatchUpdateFormDataByInstanceMapResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.AsynchronousExecution)) {
		body["asynchronousExecution"] = request.AsynchronousExecution
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.IgnoreEmpty)) {
		body["ignoreEmpty"] = request.IgnoreEmpty
	}

	if !tea.BoolValue(util.IsUnset(request.NoExecuteExpression)) {
		body["noExecuteExpression"] = request.NoExecuteExpression
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateFormDataJsonMap)) {
		body["updateFormDataJsonMap"] = request.UpdateFormDataJsonMap
	}

	if !tea.BoolValue(util.IsUnset(request.UseLatestFormSchemaVersion)) {
		body["useLatestFormSchemaVersion"] = request.UseLatestFormSchemaVersion
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
	_result = &BatchUpdateFormDataByInstanceMapResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchUpdateFormDataByInstanceMap"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/yida/forms/instances/datas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BuyAuthorizationOrder(request *BuyAuthorizationOrderRequest) (_result *BuyAuthorizationOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BuyAuthorizationOrderHeaders{}
	_result = &BuyAuthorizationOrderResponse{}
	_body, _err := client.BuyAuthorizationOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BuyAuthorizationOrderWithOptions(request *BuyAuthorizationOrderRequest, headers *BuyAuthorizationOrderHeaders, runtime *util.RuntimeOptions) (_result *BuyAuthorizationOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		body["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.AccountNumber)) {
		body["accountNumber"] = request.AccountNumber
	}

	if !tea.BoolValue(util.IsUnset(request.BeginTimeGMT)) {
		body["beginTimeGMT"] = request.BeginTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUnionId)) {
		body["callerUnionId"] = request.CallerUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		body["chargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.CommerceType)) {
		body["commerceType"] = request.CommerceType
	}

	if !tea.BoolValue(util.IsUnset(request.CommodityType)) {
		body["commodityType"] = request.CommodityType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimeGMT)) {
		body["endTimeGMT"] = request.EndTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["instanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.ProduceCode)) {
		body["produceCode"] = request.ProduceCode
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
	_result = &BuyAuthorizationOrderResponse{}
	_body, _err := client.DoROARequest(tea.String("BuyAuthorizationOrder"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/appAuthorizations/order"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BuyFreshOrder(request *BuyFreshOrderRequest) (_result *BuyFreshOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BuyFreshOrderHeaders{}
	_result = &BuyFreshOrderResponse{}
	_body, _err := client.BuyFreshOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BuyFreshOrderWithOptions(request *BuyFreshOrderRequest, headers *BuyFreshOrderHeaders, runtime *util.RuntimeOptions) (_result *BuyFreshOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		body["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.AccountNumber)) {
		body["accountNumber"] = request.AccountNumber
	}

	if !tea.BoolValue(util.IsUnset(request.BeginTimeGMT)) {
		body["beginTimeGMT"] = request.BeginTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUnionId)) {
		body["callerUnionId"] = request.CallerUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		body["chargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.CommerceType)) {
		body["commerceType"] = request.CommerceType
	}

	if !tea.BoolValue(util.IsUnset(request.CommodityType)) {
		body["commodityType"] = request.CommodityType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimeGMT)) {
		body["endTimeGMT"] = request.EndTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["instanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.ProduceCode)) {
		body["produceCode"] = request.ProduceCode
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
	_result = &BuyFreshOrderResponse{}
	_body, _err := client.DoROARequest(tea.String("BuyFreshOrder"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/apps/freshOrders"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckCloudAccountStatus(callerUid *string, request *CheckCloudAccountStatusRequest) (_result *CheckCloudAccountStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CheckCloudAccountStatusHeaders{}
	_result = &CheckCloudAccountStatusResponse{}
	_body, _err := client.CheckCloudAccountStatusWithOptions(callerUid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckCloudAccountStatusWithOptions(callerUid *string, request *CheckCloudAccountStatusRequest, headers *CheckCloudAccountStatusHeaders, runtime *util.RuntimeOptions) (_result *CheckCloudAccountStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	callerUid = openapiutil.GetEncodeParam(callerUid)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
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
	_result = &CheckCloudAccountStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("CheckCloudAccountStatus"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/apps/cloudAccountStatus/"+tea.StringValue(callerUid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOrUpdateFormData(request *CreateOrUpdateFormDataRequest) (_result *CreateOrUpdateFormDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateOrUpdateFormDataHeaders{}
	_result = &CreateOrUpdateFormDataResponse{}
	_body, _err := client.CreateOrUpdateFormDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOrUpdateFormDataWithOptions(request *CreateOrUpdateFormDataRequest, headers *CreateOrUpdateFormDataHeaders, runtime *util.RuntimeOptions) (_result *CreateOrUpdateFormDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.FormDataJson)) {
		body["formDataJson"] = request.FormDataJson
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.NoExecuteExpression)) {
		body["noExecuteExpression"] = request.NoExecuteExpression
	}

	if !tea.BoolValue(util.IsUnset(request.SearchCondition)) {
		body["searchCondition"] = request.SearchCondition
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
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
	_result = &CreateOrUpdateFormDataResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateOrUpdateFormData"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/forms/instances/insertOrUpdate"), tea.String("json"), req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.FormInstanceId)) {
		query["formInstanceId"] = request.FormInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
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
	_result = &DeleteFormDataResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteFormData"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/yida/forms/instances"), tea.String("none"), req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		query["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
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
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteInstance"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/yida/processes/instances"), tea.String("none"), req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Sequence)) {
		query["sequence"] = request.Sequence
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
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
	_result = &DeleteSequenceResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteSequence"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/yida/forms/deleteSequence"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeployFunctionCallback(request *DeployFunctionCallbackRequest) (_result *DeployFunctionCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeployFunctionCallbackHeaders{}
	_result = &DeployFunctionCallbackResponse{}
	_body, _err := client.DeployFunctionCallbackWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeployFunctionCallbackWithOptions(request *DeployFunctionCallbackRequest, headers *DeployFunctionCallbackHeaders, runtime *util.RuntimeOptions) (_result *DeployFunctionCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomDomain)) {
		body["customDomain"] = request.CustomDomain
	}

	if !tea.BoolValue(util.IsUnset(request.DeployStage)) {
		body["deployStage"] = request.DeployStage
	}

	if !tea.BoolValue(util.IsUnset(request.GateWayAppKey)) {
		body["gateWayAppKey"] = request.GateWayAppKey
	}

	if !tea.BoolValue(util.IsUnset(request.GateWayAppSecret)) {
		body["gateWayAppSecret"] = request.GateWayAppSecret
	}

	if !tea.BoolValue(util.IsUnset(request.GateWayDomain)) {
		body["gateWayDomain"] = request.GateWayDomain
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
	_result = &DeployFunctionCallbackResponse{}
	_body, _err := client.DoROARequest(tea.String("DeployFunctionCallback"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/functionComputeConnectors/completeDeployments/notify"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteBatchTask(request *ExecuteBatchTaskRequest) (_result *ExecuteBatchTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ExecuteBatchTaskHeaders{}
	_result = &ExecuteBatchTaskResponse{}
	_body, _err := client.ExecuteBatchTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteBatchTaskWithOptions(request *ExecuteBatchTaskRequest, headers *ExecuteBatchTaskHeaders, runtime *util.RuntimeOptions) (_result *ExecuteBatchTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.OutResult)) {
		body["outResult"] = request.OutResult
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.TaskInformationList)) {
		body["taskInformationList"] = request.TaskInformationList
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
	_result = &ExecuteBatchTaskResponse{}
	_body, _err := client.DoROARequest(tea.String("ExecuteBatchTask"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/tasks/batches/execute"), tea.String("json"), req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		query["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["serviceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
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
	_result = &ExecuteCustomApiResponse{}
	_body, _err := client.DoROARequest(tea.String("ExecuteCustomApi"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/apps/customApi/execute"), tea.String("json"), req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.FormDataJson)) {
		body["formDataJson"] = request.FormDataJson
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.NoExecuteExpressions)) {
		body["noExecuteExpressions"] = request.NoExecuteExpressions
	}

	if !tea.BoolValue(util.IsUnset(request.OutResult)) {
		body["outResult"] = request.OutResult
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		body["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
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
	_result = &ExecutePlatformTaskResponse{}
	_body, _err := client.DoROARequest(tea.String("ExecutePlatformTask"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/tasks/platformTasks/execute"), tea.String("none"), req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.DigitalSignUrl)) {
		body["digitalSignUrl"] = request.DigitalSignUrl
	}

	if !tea.BoolValue(util.IsUnset(request.FormDataJson)) {
		body["formDataJson"] = request.FormDataJson
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.NoExecuteExpressions)) {
		body["noExecuteExpressions"] = request.NoExecuteExpressions
	}

	if !tea.BoolValue(util.IsUnset(request.OutResult)) {
		body["outResult"] = request.OutResult
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		body["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
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
	_result = &ExecuteTaskResponse{}
	_body, _err := client.DoROARequest(tea.String("ExecuteTask"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/tasks/execute"), tea.String("none"), req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUid)) {
		query["callerUid"] = request.CallerUid
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
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
	_result = &ExpireCommodityResponse{}
	_body, _err := client.DoROARequest(tea.String("ExpireCommodity"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/yida/appAuth/commodities/expire"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetActivationCodeByCallerUnionId(callerUid *string, request *GetActivationCodeByCallerUnionIdRequest) (_result *GetActivationCodeByCallerUnionIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetActivationCodeByCallerUnionIdHeaders{}
	_result = &GetActivationCodeByCallerUnionIdResponse{}
	_body, _err := client.GetActivationCodeByCallerUnionIdWithOptions(callerUid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetActivationCodeByCallerUnionIdWithOptions(callerUid *string, request *GetActivationCodeByCallerUnionIdRequest, headers *GetActivationCodeByCallerUnionIdHeaders, runtime *util.RuntimeOptions) (_result *GetActivationCodeByCallerUnionIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	callerUid = openapiutil.GetEncodeParam(callerUid)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
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
	_result = &GetActivationCodeByCallerUnionIdResponse{}
	_body, _err := client.DoROARequest(tea.String("GetActivationCodeByCallerUnionId"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/applications/activationCodes/"+tea.StringValue(callerUid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetActivityButtonList(appType *string, processCode *string, activityId *string, request *GetActivityButtonListRequest) (_result *GetActivityButtonListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetActivityButtonListHeaders{}
	_result = &GetActivityButtonListResponse{}
	_body, _err := client.GetActivityButtonListWithOptions(appType, processCode, activityId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetActivityButtonListWithOptions(appType *string, processCode *string, activityId *string, request *GetActivityButtonListRequest, headers *GetActivityButtonListHeaders, runtime *util.RuntimeOptions) (_result *GetActivityButtonListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appType = openapiutil.GetEncodeParam(appType)
	processCode = openapiutil.GetEncodeParam(processCode)
	activityId = openapiutil.GetEncodeParam(activityId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
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
	_result = &GetActivityButtonListResponse{}
	_body, _err := client.DoROARequest(tea.String("GetActivityButtonList"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/processDefinitions/buttons/"+tea.StringValue(appType)+"/"+tea.StringValue(processCode)+"/"+tea.StringValue(activityId)), tea.String("json"), req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		query["processCode"] = request.ProcessCode
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
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
	_result = &GetActivityListResponse{}
	_body, _err := client.DoROARequest(tea.String("GetActivityList"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/processes/activities"), tea.String("json"), req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUid)) {
		query["callerUid"] = request.CallerUid
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
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
	_result = &GetApplicationAuthorizationServicePlatformResourceResponse{}
	_body, _err := client.DoROARequest(tea.String("GetApplicationAuthorizationServicePlatformResource"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/authorization/platformResources"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCorpAccomplishmentTasks(corpId *string, userId *string, request *GetCorpAccomplishmentTasksRequest) (_result *GetCorpAccomplishmentTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCorpAccomplishmentTasksHeaders{}
	_result = &GetCorpAccomplishmentTasksResponse{}
	_body, _err := client.GetCorpAccomplishmentTasksWithOptions(corpId, userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCorpAccomplishmentTasksWithOptions(corpId *string, userId *string, request *GetCorpAccomplishmentTasksRequest, headers *GetCorpAccomplishmentTasksHeaders, runtime *util.RuntimeOptions) (_result *GetCorpAccomplishmentTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	corpId = openapiutil.GetEncodeParam(corpId)
	userId = openapiutil.GetEncodeParam(userId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppTypes)) {
		query["appTypes"] = request.AppTypes
	}

	if !tea.BoolValue(util.IsUnset(request.CreateFromTimeGMT)) {
		query["createFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CreateToTimeGMT)) {
		query["createToTimeGMT"] = request.CreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCodes)) {
		query["processCodes"] = request.ProcessCodes
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["token"] = request.Token
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
	_result = &GetCorpAccomplishmentTasksResponse{}
	_body, _err := client.DoROARequest(tea.String("GetCorpAccomplishmentTasks"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/tasks/completedTasks/"+tea.StringValue(corpId)+"/"+tea.StringValue(userId)), tea.String("json"), req, runtime)
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
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

func (client *Client) GetCorpTasks(request *GetCorpTasksRequest) (_result *GetCorpTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCorpTasksHeaders{}
	_result = &GetCorpTasksResponse{}
	_body, _err := client.GetCorpTasksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCorpTasksWithOptions(request *GetCorpTasksRequest, headers *GetCorpTasksHeaders, runtime *util.RuntimeOptions) (_result *GetCorpTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppTypes)) {
		query["appTypes"] = request.AppTypes
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.CreateFromTimeGMT)) {
		query["createFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CreateToTimeGMT)) {
		query["createToTimeGMT"] = request.CreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCodes)) {
		query["processCodes"] = request.ProcessCodes
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["token"] = request.Token
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
	_result = &GetCorpTasksResponse{}
	_body, _err := client.DoROARequest(tea.String("GetCorpTasks"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/corpTasks"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFieldDefByUuid(request *GetFieldDefByUuidRequest) (_result *GetFieldDefByUuidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFieldDefByUuidHeaders{}
	_result = &GetFieldDefByUuidResponse{}
	_body, _err := client.GetFieldDefByUuidWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFieldDefByUuidWithOptions(request *GetFieldDefByUuidRequest, headers *GetFieldDefByUuidHeaders, runtime *util.RuntimeOptions) (_result *GetFieldDefByUuidResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		query["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
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
	_result = &GetFieldDefByUuidResponse{}
	_body, _err := client.DoROARequest(tea.String("GetFieldDefByUuid"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/forms/formFields"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFormComponentDefinitionList(appType *string, formUuid *string, request *GetFormComponentDefinitionListRequest) (_result *GetFormComponentDefinitionListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFormComponentDefinitionListHeaders{}
	_result = &GetFormComponentDefinitionListResponse{}
	_body, _err := client.GetFormComponentDefinitionListWithOptions(appType, formUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFormComponentDefinitionListWithOptions(appType *string, formUuid *string, request *GetFormComponentDefinitionListRequest, headers *GetFormComponentDefinitionListHeaders, runtime *util.RuntimeOptions) (_result *GetFormComponentDefinitionListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appType = openapiutil.GetEncodeParam(appType)
	formUuid = openapiutil.GetEncodeParam(formUuid)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		query["version"] = request.Version
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
	_result = &GetFormComponentDefinitionListResponse{}
	_body, _err := client.DoROARequest(tea.String("GetFormComponentDefinitionList"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/forms/definitions/"+tea.StringValue(appType)+"/"+tea.StringValue(formUuid)), tea.String("json"), req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
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
	_result = &GetFormDataByIDResponse{}
	_body, _err := client.DoROARequest(tea.String("GetFormDataByID"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/forms/instances/"+tea.StringValue(id)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFormListInApp(request *GetFormListInAppRequest) (_result *GetFormListInAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFormListInAppHeaders{}
	_result = &GetFormListInAppResponse{}
	_body, _err := client.GetFormListInAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFormListInAppWithOptions(request *GetFormListInAppRequest, headers *GetFormListInAppHeaders, runtime *util.RuntimeOptions) (_result *GetFormListInAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.FormTypes)) {
		query["formTypes"] = request.FormTypes
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
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
	_result = &GetFormListInAppResponse{}
	_body, _err := client.DoROARequest(tea.String("GetFormListInApp"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/forms"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstanceById(id *string, request *GetInstanceByIdRequest) (_result *GetInstanceByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetInstanceByIdHeaders{}
	_result = &GetInstanceByIdResponse{}
	_body, _err := client.GetInstanceByIdWithOptions(id, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceByIdWithOptions(id *string, request *GetInstanceByIdRequest, headers *GetInstanceByIdHeaders, runtime *util.RuntimeOptions) (_result *GetInstanceByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	id = openapiutil.GetEncodeParam(id)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
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
	_result = &GetInstanceByIdResponse{}
	_body, _err := client.DoROARequest(tea.String("GetInstanceById"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/processes/instancesInfos/"+tea.StringValue(id)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstanceIdList(request *GetInstanceIdListRequest) (_result *GetInstanceIdListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetInstanceIdListHeaders{}
	_result = &GetInstanceIdListResponse{}
	_body, _err := client.GetInstanceIdListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceIdListWithOptions(request *GetInstanceIdListRequest, headers *GetInstanceIdListHeaders, runtime *util.RuntimeOptions) (_result *GetInstanceIdListResponse, _err error) {
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

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.ApprovedResult)) {
		body["approvedResult"] = request.ApprovedResult
	}

	if !tea.BoolValue(util.IsUnset(request.CreateFromTimeGMT)) {
		body["createFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CreateToTimeGMT)) {
		body["createToTimeGMT"] = request.CreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceStatus)) {
		body["instanceStatus"] = request.InstanceStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedFromTimeGMT)) {
		body["modifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedToTimeGMT)) {
		body["modifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.OriginatorId)) {
		body["originatorId"] = request.OriginatorId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchFieldJson)) {
		body["searchFieldJson"] = request.SearchFieldJson
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
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
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &GetInstanceIdListResponse{}
	_body, _err := client.DoROARequest(tea.String("GetInstanceIdList"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/processes/instanceIds"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstances(request *GetInstancesRequest) (_result *GetInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetInstancesHeaders{}
	_result = &GetInstancesResponse{}
	_body, _err := client.GetInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstancesWithOptions(request *GetInstancesRequest, headers *GetInstancesHeaders, runtime *util.RuntimeOptions) (_result *GetInstancesResponse, _err error) {
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

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.ApprovedResult)) {
		body["approvedResult"] = request.ApprovedResult
	}

	if !tea.BoolValue(util.IsUnset(request.CreateFromTimeGMT)) {
		body["createFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CreateToTimeGMT)) {
		body["createToTimeGMT"] = request.CreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceStatus)) {
		body["instanceStatus"] = request.InstanceStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedFromTimeGMT)) {
		body["modifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedToTimeGMT)) {
		body["modifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.OrderConfigJson)) {
		body["orderConfigJson"] = request.OrderConfigJson
	}

	if !tea.BoolValue(util.IsUnset(request.OriginatorId)) {
		body["originatorId"] = request.OriginatorId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchFieldJson)) {
		body["searchFieldJson"] = request.SearchFieldJson
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
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
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &GetInstancesResponse{}
	_body, _err := client.DoROARequest(tea.String("GetInstances"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/processes/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstancesByIdList(request *GetInstancesByIdListRequest) (_result *GetInstancesByIdListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetInstancesByIdListHeaders{}
	_result = &GetInstancesByIdListResponse{}
	_body, _err := client.GetInstancesByIdListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstancesByIdListWithOptions(request *GetInstancesByIdListRequest, headers *GetInstancesByIdListHeaders, runtime *util.RuntimeOptions) (_result *GetInstancesByIdListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceIds)) {
		query["processInstanceIds"] = request.ProcessInstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
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
	_result = &GetInstancesByIdListResponse{}
	_body, _err := client.DoROARequest(tea.String("GetInstancesByIdList"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/processes/instances/searchWithIds"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMeCorpSubmission(userId *string, request *GetMeCorpSubmissionRequest) (_result *GetMeCorpSubmissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetMeCorpSubmissionHeaders{}
	_result = &GetMeCorpSubmissionResponse{}
	_body, _err := client.GetMeCorpSubmissionWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMeCorpSubmissionWithOptions(userId *string, request *GetMeCorpSubmissionRequest, headers *GetMeCorpSubmissionHeaders, runtime *util.RuntimeOptions) (_result *GetMeCorpSubmissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppTypes)) {
		query["appTypes"] = request.AppTypes
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.CreateFromTimeGMT)) {
		query["createFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CreateToTimeGMT)) {
		query["createToTimeGMT"] = request.CreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCodes)) {
		query["processCodes"] = request.ProcessCodes
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["token"] = request.Token
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
	_result = &GetMeCorpSubmissionResponse{}
	_body, _err := client.DoROARequest(tea.String("GetMeCorpSubmission"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/tasks/myCorpSubmission/"+tea.StringValue(userId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNotifyMe(userId *string, request *GetNotifyMeRequest) (_result *GetNotifyMeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetNotifyMeHeaders{}
	_result = &GetNotifyMeResponse{}
	_body, _err := client.GetNotifyMeWithOptions(userId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNotifyMeWithOptions(userId *string, request *GetNotifyMeRequest, headers *GetNotifyMeHeaders, runtime *util.RuntimeOptions) (_result *GetNotifyMeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	userId = openapiutil.GetEncodeParam(userId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppTypes)) {
		query["appTypes"] = request.AppTypes
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.CreateFromTimeGMT)) {
		query["createFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CreateToTimeGMT)) {
		query["createToTimeGMT"] = request.CreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceCreateFromTimeGMT)) {
		query["instanceCreateFromTimeGMT"] = request.InstanceCreateFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceCreateToTimeGMT)) {
		query["instanceCreateToTimeGMT"] = request.InstanceCreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCodes)) {
		query["processCodes"] = request.ProcessCodes
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["token"] = request.Token
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
	_result = &GetNotifyMeResponse{}
	_body, _err := client.DoROARequest(tea.String("GetNotifyMe"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/corpNotifications/"+tea.StringValue(userId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOpenUrl(appType *string, request *GetOpenUrlRequest) (_result *GetOpenUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetOpenUrlHeaders{}
	_result = &GetOpenUrlResponse{}
	_body, _err := client.GetOpenUrlWithOptions(appType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOpenUrlWithOptions(appType *string, request *GetOpenUrlRequest, headers *GetOpenUrlHeaders, runtime *util.RuntimeOptions) (_result *GetOpenUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appType = openapiutil.GetEncodeParam(appType)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		query["fileUrl"] = request.FileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		query["timeout"] = request.Timeout
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
	_result = &GetOpenUrlResponse{}
	_body, _err := client.DoROARequest(tea.String("GetOpenUrl"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/apps/temporaryUrls/"+tea.StringValue(appType)), tea.String("json"), req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		query["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
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
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUid)) {
		query["callerUid"] = request.CallerUid
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
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
	_result = &GetPlatformResourceResponse{}
	_body, _err := client.DoROARequest(tea.String("GetPlatformResource"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/apps/platformResources"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPrintAppInfo(request *GetPrintAppInfoRequest) (_result *GetPrintAppInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPrintAppInfoHeaders{}
	_result = &GetPrintAppInfoResponse{}
	_body, _err := client.GetPrintAppInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPrintAppInfoWithOptions(request *GetPrintAppInfoRequest, headers *GetPrintAppInfoHeaders, runtime *util.RuntimeOptions) (_result *GetPrintAppInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NameLike)) {
		query["nameLike"] = request.NameLike
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
	_result = &GetPrintAppInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetPrintAppInfo"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/printTemplates/printAppInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPrintDictionary(request *GetPrintDictionaryRequest) (_result *GetPrintDictionaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetPrintDictionaryHeaders{}
	_result = &GetPrintDictionaryResponse{}
	_body, _err := client.GetPrintDictionaryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPrintDictionaryWithOptions(request *GetPrintDictionaryRequest, headers *GetPrintDictionaryHeaders, runtime *util.RuntimeOptions) (_result *GetPrintDictionaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		query["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		query["version"] = request.Version
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
	_result = &GetPrintDictionaryResponse{}
	_body, _err := client.DoROARequest(tea.String("GetPrintDictionary"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/printTemplates/printDictionaries"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProcessDefinition(processInstanceId *string, request *GetProcessDefinitionRequest) (_result *GetProcessDefinitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetProcessDefinitionHeaders{}
	_result = &GetProcessDefinitionResponse{}
	_body, _err := client.GetProcessDefinitionWithOptions(processInstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProcessDefinitionWithOptions(processInstanceId *string, request *GetProcessDefinitionRequest, headers *GetProcessDefinitionHeaders, runtime *util.RuntimeOptions) (_result *GetProcessDefinitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	processInstanceId = openapiutil.GetEncodeParam(processInstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["groupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.NameSpace)) {
		query["nameSpace"] = request.NameSpace
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNumber)) {
		query["orderNumber"] = request.OrderNumber
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.SystemType)) {
		query["systemType"] = request.SystemType
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
	_result = &GetProcessDefinitionResponse{}
	_body, _err := client.DoROARequest(tea.String("GetProcessDefinition"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/processes/definitions/"+tea.StringValue(processInstanceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRunningTaskList(request *GetRunningTaskListRequest) (_result *GetRunningTaskListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetRunningTaskListHeaders{}
	_result = &GetRunningTaskListResponse{}
	_body, _err := client.GetRunningTaskListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRunningTaskListWithOptions(request *GetRunningTaskListRequest, headers *GetRunningTaskListHeaders, runtime *util.RuntimeOptions) (_result *GetRunningTaskListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceIdList)) {
		body["processInstanceIdList"] = request.ProcessInstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UserCorpId)) {
		body["userCorpId"] = request.UserCorpId
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
	_result = &GetRunningTaskListResponse{}
	_body, _err := client.DoROARequest(tea.String("GetRunningTaskList"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/tasks/runningTasks/query"), tea.String("json"), req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		query["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
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
	_result = &GetRunningTasksResponse{}
	_body, _err := client.DoROARequest(tea.String("GetRunningTasks"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/processes/tasks/getRunningTasks"), tea.String("json"), req, runtime)
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
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

func (client *Client) GetTaskCopies(request *GetTaskCopiesRequest) (_result *GetTaskCopiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetTaskCopiesHeaders{}
	_result = &GetTaskCopiesResponse{}
	_body, _err := client.GetTaskCopiesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTaskCopiesWithOptions(request *GetTaskCopiesRequest, headers *GetTaskCopiesHeaders, runtime *util.RuntimeOptions) (_result *GetTaskCopiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.CreateFromTimeGMT)) {
		query["createFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CreateToTimeGMT)) {
		query["createToTimeGMT"] = request.CreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCodes)) {
		query["processCodes"] = request.ProcessCodes
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
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
	_result = &GetTaskCopiesResponse{}
	_body, _err := client.DoROARequest(tea.String("GetTaskCopies"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/tasks/taskCopies"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApplication(request *ListApplicationRequest) (_result *ListApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListApplicationHeaders{}
	_result = &ListApplicationResponse{}
	_body, _err := client.ListApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListApplicationWithOptions(request *ListApplicationRequest, headers *ListApplicationHeaders, runtime *util.RuntimeOptions) (_result *ListApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppFilter)) {
		query["appFilter"] = request.AppFilter
	}

	if !tea.BoolValue(util.IsUnset(request.AppNameSearchKeyword)) {
		query["appNameSearchKeyword"] = request.AppNameSearchKeyword
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		query["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["token"] = request.Token
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
	_result = &ListApplicationResponse{}
	_body, _err := client.DoROARequest(tea.String("ListApplication"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/organizations/applications"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApplicationAuthorizationServiceApplicationInformation(instanceId *string, request *ListApplicationAuthorizationServiceApplicationInformationRequest) (_result *ListApplicationAuthorizationServiceApplicationInformationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListApplicationAuthorizationServiceApplicationInformationHeaders{}
	_result = &ListApplicationAuthorizationServiceApplicationInformationResponse{}
	_body, _err := client.ListApplicationAuthorizationServiceApplicationInformationWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListApplicationAuthorizationServiceApplicationInformationWithOptions(instanceId *string, request *ListApplicationAuthorizationServiceApplicationInformationRequest, headers *ListApplicationAuthorizationServiceApplicationInformationHeaders, runtime *util.RuntimeOptions) (_result *ListApplicationAuthorizationServiceApplicationInformationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	instanceId = openapiutil.GetEncodeParam(instanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUnionId)) {
		query["callerUnionId"] = request.CallerUnionId
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
	_result = &ListApplicationAuthorizationServiceApplicationInformationResponse{}
	_body, _err := client.DoROARequest(tea.String("ListApplicationAuthorizationServiceApplicationInformation"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/authorizations/applicationInfos/"+tea.StringValue(instanceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApplicationAuthorizationServiceConnectorInformation(instanceId *string, request *ListApplicationAuthorizationServiceConnectorInformationRequest) (_result *ListApplicationAuthorizationServiceConnectorInformationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListApplicationAuthorizationServiceConnectorInformationHeaders{}
	_result = &ListApplicationAuthorizationServiceConnectorInformationResponse{}
	_body, _err := client.ListApplicationAuthorizationServiceConnectorInformationWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListApplicationAuthorizationServiceConnectorInformationWithOptions(instanceId *string, request *ListApplicationAuthorizationServiceConnectorInformationRequest, headers *ListApplicationAuthorizationServiceConnectorInformationHeaders, runtime *util.RuntimeOptions) (_result *ListApplicationAuthorizationServiceConnectorInformationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	instanceId = openapiutil.GetEncodeParam(instanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUid)) {
		query["callerUid"] = request.CallerUid
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
	_result = &ListApplicationAuthorizationServiceConnectorInformationResponse{}
	_body, _err := client.DoROARequest(tea.String("ListApplicationAuthorizationServiceConnectorInformation"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/applicationAuthorizations/plugs/"+tea.StringValue(instanceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApplicationInformation(instanceId *string, request *ListApplicationInformationRequest) (_result *ListApplicationInformationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListApplicationInformationHeaders{}
	_result = &ListApplicationInformationResponse{}
	_body, _err := client.ListApplicationInformationWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListApplicationInformationWithOptions(instanceId *string, request *ListApplicationInformationRequest, headers *ListApplicationInformationHeaders, runtime *util.RuntimeOptions) (_result *ListApplicationInformationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	instanceId = openapiutil.GetEncodeParam(instanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUid)) {
		query["callerUid"] = request.CallerUid
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
	_result = &ListApplicationInformationResponse{}
	_body, _err := client.DoROARequest(tea.String("ListApplicationInformation"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/apps/infos/"+tea.StringValue(instanceId)), tea.String("json"), req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.CallerUid)) {
		query["callerUid"] = request.CallerUid
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
	_result = &ListCommodityResponse{}
	_body, _err := client.DoROARequest(tea.String("ListCommodity"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/appAuth/commodities"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConnectorInformation(instanceId *string, request *ListConnectorInformationRequest) (_result *ListConnectorInformationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListConnectorInformationHeaders{}
	_result = &ListConnectorInformationResponse{}
	_body, _err := client.ListConnectorInformationWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConnectorInformationWithOptions(instanceId *string, request *ListConnectorInformationRequest, headers *ListConnectorInformationHeaders, runtime *util.RuntimeOptions) (_result *ListConnectorInformationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	instanceId = openapiutil.GetEncodeParam(instanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUid)) {
		query["callerUid"] = request.CallerUid
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
	_result = &ListConnectorInformationResponse{}
	_body, _err := client.DoROARequest(tea.String("ListConnectorInformation"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/plugins/infos/"+tea.StringValue(instanceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFormRemarks(request *ListFormRemarksRequest) (_result *ListFormRemarksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListFormRemarksHeaders{}
	_result = &ListFormRemarksResponse{}
	_body, _err := client.ListFormRemarksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFormRemarksWithOptions(request *ListFormRemarksRequest, headers *ListFormRemarksHeaders, runtime *util.RuntimeOptions) (_result *ListFormRemarksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.FormInstanceIdList)) {
		body["formInstanceIdList"] = request.FormInstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
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
	_result = &ListFormRemarksResponse{}
	_body, _err := client.DoROARequest(tea.String("ListFormRemarks"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/forms/remarks/query"), tea.String("json"), req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.FormType)) {
		query["formType"] = request.FormType
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
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
	_result = &ListNavigationByFormTypeResponse{}
	_body, _err := client.DoROARequest(tea.String("ListNavigationByFormType"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/apps/navigations"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOperationLogs(request *ListOperationLogsRequest) (_result *ListOperationLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListOperationLogsHeaders{}
	_result = &ListOperationLogsResponse{}
	_body, _err := client.ListOperationLogsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOperationLogsWithOptions(request *ListOperationLogsRequest, headers *ListOperationLogsHeaders, runtime *util.RuntimeOptions) (_result *ListOperationLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.FormInstanceIdList)) {
		body["formInstanceIdList"] = request.FormInstanceIdList
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
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
	_result = &ListOperationLogsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListOperationLogs"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/forms/operationsLogs/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTableDataByFormInstanceIdTableId(formInstanceId *string, request *ListTableDataByFormInstanceIdTableIdRequest) (_result *ListTableDataByFormInstanceIdTableIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListTableDataByFormInstanceIdTableIdHeaders{}
	_result = &ListTableDataByFormInstanceIdTableIdResponse{}
	_body, _err := client.ListTableDataByFormInstanceIdTableIdWithOptions(formInstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTableDataByFormInstanceIdTableIdWithOptions(formInstanceId *string, request *ListTableDataByFormInstanceIdTableIdRequest, headers *ListTableDataByFormInstanceIdTableIdHeaders, runtime *util.RuntimeOptions) (_result *ListTableDataByFormInstanceIdTableIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	formInstanceId = openapiutil.GetEncodeParam(formInstanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		query["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.TableFieldId)) {
		query["tableFieldId"] = request.TableFieldId
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
	_result = &ListTableDataByFormInstanceIdTableIdResponse{}
	_body, _err := client.DoROARequest(tea.String("ListTableDataByFormInstanceIdTableId"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/forms/innerTables/"+tea.StringValue(formInstanceId)), tea.String("json"), req, runtime)
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
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

func (client *Client) NotifyAuthorizationResult(request *NotifyAuthorizationResultRequest) (_result *NotifyAuthorizationResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &NotifyAuthorizationResultHeaders{}
	_result = &NotifyAuthorizationResultResponse{}
	_body, _err := client.NotifyAuthorizationResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) NotifyAuthorizationResultWithOptions(request *NotifyAuthorizationResultRequest, headers *NotifyAuthorizationResultHeaders, runtime *util.RuntimeOptions) (_result *NotifyAuthorizationResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		body["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.AccountNumber)) {
		body["accountNumber"] = request.AccountNumber
	}

	if !tea.BoolValue(util.IsUnset(request.BeginTimeGMT)) {
		body["beginTimeGMT"] = request.BeginTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUid)) {
		body["callerUid"] = request.CallerUid
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		body["chargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.CommerceType)) {
		body["commerceType"] = request.CommerceType
	}

	if !tea.BoolValue(util.IsUnset(request.CommodityType)) {
		body["commodityType"] = request.CommodityType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimeGMT)) {
		body["endTimeGMT"] = request.EndTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["instanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.ProduceCode)) {
		body["produceCode"] = request.ProduceCode
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
	_result = &NotifyAuthorizationResultResponse{}
	_body, _err := client.DoROARequest(tea.String("NotifyAuthorizationResult"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/apps/authorizationResults/notify"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PageFormBaseInfos(request *PageFormBaseInfosRequest) (_result *PageFormBaseInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PageFormBaseInfosHeaders{}
	_result = &PageFormBaseInfosResponse{}
	_body, _err := client.PageFormBaseInfosWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PageFormBaseInfosWithOptions(request *PageFormBaseInfosRequest, headers *PageFormBaseInfosHeaders, runtime *util.RuntimeOptions) (_result *PageFormBaseInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["appKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.FormTypeList)) {
		body["formTypeList"] = request.FormTypeList
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		body["pageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
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
	_result = &PageFormBaseInfosResponse{}
	_body, _err := client.DoROARequest(tea.String("PageFormBaseInfos"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/apps/forms/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryServiceRecord(request *QueryServiceRecordRequest) (_result *QueryServiceRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryServiceRecordHeaders{}
	_result = &QueryServiceRecordResponse{}
	_body, _err := client.QueryServiceRecordWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryServiceRecordWithOptions(request *QueryServiceRecordRequest, headers *QueryServiceRecordHeaders, runtime *util.RuntimeOptions) (_result *QueryServiceRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		query["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.HookType)) {
		query["hookType"] = request.HookType
	}

	if !tea.BoolValue(util.IsUnset(request.HookUuid)) {
		query["hookUuid"] = request.HookUuid
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InvokeAfterDateGMT)) {
		query["invokeAfterDateGMT"] = request.InvokeAfterDateGMT
	}

	if !tea.BoolValue(util.IsUnset(request.InvokeBeforeDateGMT)) {
		query["invokeBeforeDateGMT"] = request.InvokeBeforeDateGMT
	}

	if !tea.BoolValue(util.IsUnset(request.InvokeStatus)) {
		query["invokeStatus"] = request.InvokeStatus
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RequestUrl)) {
		query["requestUrl"] = request.RequestUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SourceUuid)) {
		query["sourceUuid"] = request.SourceUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Success)) {
		query["success"] = request.Success
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
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
	_result = &QueryServiceRecordResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryServiceRecord"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/services/invocationRecords"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RedirectTask(request *RedirectTaskRequest) (_result *RedirectTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RedirectTaskHeaders{}
	_result = &RedirectTaskResponse{}
	_body, _err := client.RedirectTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RedirectTaskWithOptions(request *RedirectTaskRequest, headers *RedirectTaskHeaders, runtime *util.RuntimeOptions) (_result *RedirectTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.ByManager)) {
		body["byManager"] = request.ByManager
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.NowActionExecutorId)) {
		body["nowActionExecutorId"] = request.NowActionExecutorId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		body["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
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
	_result = &RedirectTaskResponse{}
	_body, _err := client.DoROARequest(tea.String("RedirectTask"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/tasks/redirect"), tea.String("none"), req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUid)) {
		query["callerUid"] = request.CallerUid
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
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
	_result = &RefundCommodityResponse{}
	_body, _err := client.DoROARequest(tea.String("RefundCommodity"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/appAuth/commodities/refund"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterAccounts(request *RegisterAccountsRequest) (_result *RegisterAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RegisterAccountsHeaders{}
	_result = &RegisterAccountsResponse{}
	_body, _err := client.RegisterAccountsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterAccountsWithOptions(request *RegisterAccountsRequest, headers *RegisterAccountsHeaders, runtime *util.RuntimeOptions) (_result *RegisterAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		body["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.ActiveCode)) {
		body["activeCode"] = request.ActiveCode
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
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
	_result = &RegisterAccountsResponse{}
	_body, _err := client.DoROARequest(tea.String("RegisterAccounts"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/applicationAuthorizations/accounts/register"), tea.String("json"), req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUid)) {
		query["callerUid"] = request.CallerUid
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
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
	_result = &ReleaseCommodityResponse{}
	_body, _err := client.DoROARequest(tea.String("ReleaseCommodity"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/yida/appAuth/commodities/release"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveTenantResource(callerUid *string, request *RemoveTenantResourceRequest) (_result *RemoveTenantResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RemoveTenantResourceHeaders{}
	_result = &RemoveTenantResourceResponse{}
	_body, _err := client.RemoveTenantResourceWithOptions(callerUid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveTenantResourceWithOptions(callerUid *string, request *RemoveTenantResourceRequest, headers *RemoveTenantResourceHeaders, runtime *util.RuntimeOptions) (_result *RemoveTenantResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	callerUid = openapiutil.GetEncodeParam(callerUid)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
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
	_result = &RemoveTenantResourceResponse{}
	_body, _err := client.DoROARequest(tea.String("RemoveTenantResource"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1.0/yida/applications/tenantRelatedResources/"+tea.StringValue(callerUid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenderBatchCallback(request *RenderBatchCallbackRequest) (_result *RenderBatchCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RenderBatchCallbackHeaders{}
	_result = &RenderBatchCallbackResponse{}
	_body, _err := client.RenderBatchCallbackWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenderBatchCallbackWithOptions(request *RenderBatchCallbackRequest, headers *RenderBatchCallbackHeaders, runtime *util.RuntimeOptions) (_result *RenderBatchCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.FileSize)) {
		body["fileSize"] = request.FileSize
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		body["namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.OssUrl)) {
		body["ossUrl"] = request.OssUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SequenceId)) {
		body["sequenceId"] = request.SequenceId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.TimeZone)) {
		body["timeZone"] = request.TimeZone
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
	_result = &RenderBatchCallbackResponse{}
	_body, _err := client.DoROARequest(tea.String("RenderBatchCallback"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/printings/callbacks/batch"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewApplicationAuthorizationServiceOrder(request *RenewApplicationAuthorizationServiceOrderRequest) (_result *RenewApplicationAuthorizationServiceOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RenewApplicationAuthorizationServiceOrderHeaders{}
	_result = &RenewApplicationAuthorizationServiceOrderResponse{}
	_body, _err := client.RenewApplicationAuthorizationServiceOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewApplicationAuthorizationServiceOrderWithOptions(request *RenewApplicationAuthorizationServiceOrderRequest, headers *RenewApplicationAuthorizationServiceOrderHeaders, runtime *util.RuntimeOptions) (_result *RenewApplicationAuthorizationServiceOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		body["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUnionId)) {
		body["callerUnionId"] = request.CallerUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimeGMT)) {
		body["endTimeGMT"] = request.EndTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
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
	_result = &RenewApplicationAuthorizationServiceOrderResponse{}
	_body, _err := client.DoROARequest(tea.String("RenewApplicationAuthorizationServiceOrder"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/applicationAuthorizations/orders/renew"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewTenantOrder(request *RenewTenantOrderRequest) (_result *RenewTenantOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RenewTenantOrderHeaders{}
	_result = &RenewTenantOrderResponse{}
	_body, _err := client.RenewTenantOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewTenantOrderWithOptions(request *RenewTenantOrderRequest, headers *RenewTenantOrderHeaders, runtime *util.RuntimeOptions) (_result *RenewTenantOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		body["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUnionId)) {
		body["callerUnionId"] = request.CallerUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimeGMT)) {
		body["endTimeGMT"] = request.EndTimeGMT
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
	_result = &RenewTenantOrderResponse{}
	_body, _err := client.DoROARequest(tea.String("RenewTenantOrder"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/apps/tenants/reorder"), tea.String("json"), req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.FormDataJson)) {
		body["formDataJson"] = request.FormDataJson
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
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
	_result = &SaveFormDataResponse{}
	_body, _err := client.DoROARequest(tea.String("SaveFormData"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/forms/instances"), tea.String("json"), req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.AtUserId)) {
		body["atUserId"] = request.AtUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.FormInstanceId)) {
		body["formInstanceId"] = request.FormInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ReplyId)) {
		body["replyId"] = request.ReplyId
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
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
	_result = &SaveFormRemarkResponse{}
	_body, _err := client.DoROARequest(tea.String("SaveFormRemark"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/forms/remarks"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SavePrintTplDetailInfo(request *SavePrintTplDetailInfoRequest) (_result *SavePrintTplDetailInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SavePrintTplDetailInfoHeaders{}
	_result = &SavePrintTplDetailInfoResponse{}
	_body, _err := client.SavePrintTplDetailInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SavePrintTplDetailInfoWithOptions(request *SavePrintTplDetailInfoRequest, headers *SavePrintTplDetailInfoHeaders, runtime *util.RuntimeOptions) (_result *SavePrintTplDetailInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileNameConfig)) {
		body["fileNameConfig"] = request.FileNameConfig
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.FormVersion)) {
		body["formVersion"] = request.FormVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Setting)) {
		body["setting"] = request.Setting
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["templateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Vm)) {
		body["vm"] = request.Vm
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
	_result = &SavePrintTplDetailInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("SavePrintTplDetailInfo"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/printTemplates/printTplDetailInfos"), tea.String("json"), req, runtime)
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
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
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.CreateFromTimeGMT)) {
		body["createFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CreateToTimeGMT)) {
		body["createToTimeGMT"] = request.CreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedFromTimeGMT)) {
		body["modifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedToTimeGMT)) {
		body["modifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.OriginatorId)) {
		body["originatorId"] = request.OriginatorId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchFieldJson)) {
		body["searchFieldJson"] = request.SearchFieldJson
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.TargetFieldJson)) {
		body["targetFieldJson"] = request.TargetFieldJson
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
	_result = &SearchEmployeeFieldValuesResponse{}
	_body, _err := client.DoROARequest(tea.String("SearchEmployeeFieldValues"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/forms/employeeFields"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchFormDataIdList(appType *string, formUuid *string, request *SearchFormDataIdListRequest) (_result *SearchFormDataIdListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchFormDataIdListHeaders{}
	_result = &SearchFormDataIdListResponse{}
	_body, _err := client.SearchFormDataIdListWithOptions(appType, formUuid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchFormDataIdListWithOptions(appType *string, formUuid *string, request *SearchFormDataIdListRequest, headers *SearchFormDataIdListHeaders, runtime *util.RuntimeOptions) (_result *SearchFormDataIdListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	appType = openapiutil.GetEncodeParam(appType)
	formUuid = openapiutil.GetEncodeParam(formUuid)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateFromTimeGMT)) {
		body["createFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CreateToTimeGMT)) {
		body["createToTimeGMT"] = request.CreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedFromTimeGMT)) {
		body["modifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedToTimeGMT)) {
		body["modifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.OriginatorId)) {
		body["originatorId"] = request.OriginatorId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchFieldJson)) {
		body["searchFieldJson"] = request.SearchFieldJson
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
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
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &SearchFormDataIdListResponse{}
	_body, _err := client.DoROARequest(tea.String("SearchFormDataIdList"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/forms/instances/ids/"+tea.StringValue(appType)+"/"+tea.StringValue(formUuid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchFormDataRemovalTableData(request *SearchFormDataRemovalTableDataRequest) (_result *SearchFormDataRemovalTableDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchFormDataRemovalTableDataHeaders{}
	_result = &SearchFormDataRemovalTableDataResponse{}
	_body, _err := client.SearchFormDataRemovalTableDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchFormDataRemovalTableDataWithOptions(request *SearchFormDataRemovalTableDataRequest, headers *SearchFormDataRemovalTableDataHeaders, runtime *util.RuntimeOptions) (_result *SearchFormDataRemovalTableDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.CreateFromTimeGMT)) {
		body["createFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CreateToTimeGMT)) {
		body["createToTimeGMT"] = request.CreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedFromTimeGMT)) {
		body["modifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedToTimeGMT)) {
		body["modifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.OrderConfigJson)) {
		body["orderConfigJson"] = request.OrderConfigJson
	}

	if !tea.BoolValue(util.IsUnset(request.OriginatorId)) {
		body["originatorId"] = request.OriginatorId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchFieldJson)) {
		body["searchFieldJson"] = request.SearchFieldJson
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
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
	_result = &SearchFormDataRemovalTableDataResponse{}
	_body, _err := client.DoROARequest(tea.String("SearchFormDataRemovalTableData"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/forms/instances/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchFormDataSecondGeneration(request *SearchFormDataSecondGenerationRequest) (_result *SearchFormDataSecondGenerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchFormDataSecondGenerationHeaders{}
	_result = &SearchFormDataSecondGenerationResponse{}
	_body, _err := client.SearchFormDataSecondGenerationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchFormDataSecondGenerationWithOptions(request *SearchFormDataSecondGenerationRequest, headers *SearchFormDataSecondGenerationHeaders, runtime *util.RuntimeOptions) (_result *SearchFormDataSecondGenerationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.CreateFromTimeGMT)) {
		body["createFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CreateToTimeGMT)) {
		body["createToTimeGMT"] = request.CreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedFromTimeGMT)) {
		body["modifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedToTimeGMT)) {
		body["modifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.OrderConfigJson)) {
		body["orderConfigJson"] = request.OrderConfigJson
	}

	if !tea.BoolValue(util.IsUnset(request.OriginatorId)) {
		body["originatorId"] = request.OriginatorId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchCondition)) {
		body["searchCondition"] = request.SearchCondition
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
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
	_result = &SearchFormDataSecondGenerationResponse{}
	_body, _err := client.DoROARequest(tea.String("SearchFormDataSecondGeneration"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/forms/instances/advances/queryAll"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchFormDataSecondGenerationNoTableField(request *SearchFormDataSecondGenerationNoTableFieldRequest) (_result *SearchFormDataSecondGenerationNoTableFieldResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SearchFormDataSecondGenerationNoTableFieldHeaders{}
	_result = &SearchFormDataSecondGenerationNoTableFieldResponse{}
	_body, _err := client.SearchFormDataSecondGenerationNoTableFieldWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchFormDataSecondGenerationNoTableFieldWithOptions(request *SearchFormDataSecondGenerationNoTableFieldRequest, headers *SearchFormDataSecondGenerationNoTableFieldHeaders, runtime *util.RuntimeOptions) (_result *SearchFormDataSecondGenerationNoTableFieldResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.CreateFromTimeGMT)) {
		body["createFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CreateToTimeGMT)) {
		body["createToTimeGMT"] = request.CreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedFromTimeGMT)) {
		body["modifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedToTimeGMT)) {
		body["modifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.OrderConfigJson)) {
		body["orderConfigJson"] = request.OrderConfigJson
	}

	if !tea.BoolValue(util.IsUnset(request.OriginatorId)) {
		body["originatorId"] = request.OriginatorId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchCondition)) {
		body["searchCondition"] = request.SearchCondition
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
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
	_result = &SearchFormDataSecondGenerationNoTableFieldResponse{}
	_body, _err := client.DoROARequest(tea.String("SearchFormDataSecondGenerationNoTableField"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/forms/instances/advances/query"), tea.String("json"), req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.CreateFromTimeGMT)) {
		body["createFromTimeGMT"] = request.CreateFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CreateToTimeGMT)) {
		body["createToTimeGMT"] = request.CreateToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["currentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DynamicOrder)) {
		body["dynamicOrder"] = request.DynamicOrder
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedFromTimeGMT)) {
		body["modifiedFromTimeGMT"] = request.ModifiedFromTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedToTimeGMT)) {
		body["modifiedToTimeGMT"] = request.ModifiedToTimeGMT
	}

	if !tea.BoolValue(util.IsUnset(request.OriginatorId)) {
		body["originatorId"] = request.OriginatorId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchFieldJson)) {
		body["searchFieldJson"] = request.SearchFieldJson
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
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
	_result = &SearchFormDatasResponse{}
	_body, _err := client.DoROARequest(tea.String("SearchFormDatas"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/forms/instances/search"), tea.String("json"), req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		body["departmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.FormDataJson)) {
		body["formDataJson"] = request.FormDataJson
	}

	if !tea.BoolValue(util.IsUnset(request.FormUuid)) {
		body["formUuid"] = request.FormUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		body["processCode"] = request.ProcessCode
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
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
	_result = &StartInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("StartInstance"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/processes/instances/start"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TerminateCloudAuthorization(request *TerminateCloudAuthorizationRequest) (_result *TerminateCloudAuthorizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TerminateCloudAuthorizationHeaders{}
	_result = &TerminateCloudAuthorizationResponse{}
	_body, _err := client.TerminateCloudAuthorizationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TerminateCloudAuthorizationWithOptions(request *TerminateCloudAuthorizationRequest, headers *TerminateCloudAuthorizationHeaders, runtime *util.RuntimeOptions) (_result *TerminateCloudAuthorizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		body["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUnionId)) {
		body["callerUnionId"] = request.CallerUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
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
	_result = &TerminateCloudAuthorizationResponse{}
	_body, _err := client.DoROARequest(tea.String("TerminateCloudAuthorization"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/yida/apps/cloudAuthorizations/terminate"), tea.String("json"), req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		query["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		query["systemToken"] = request.SystemToken
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
	_result = &TerminateInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("TerminateInstance"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/yida/processes/instances/terminate"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCloudAccountInformation(request *UpdateCloudAccountInformationRequest) (_result *UpdateCloudAccountInformationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateCloudAccountInformationHeaders{}
	_result = &UpdateCloudAccountInformationResponse{}
	_body, _err := client.UpdateCloudAccountInformationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCloudAccountInformationWithOptions(request *UpdateCloudAccountInformationRequest, headers *UpdateCloudAccountInformationHeaders, runtime *util.RuntimeOptions) (_result *UpdateCloudAccountInformationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		body["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.AccountNumber)) {
		body["accountNumber"] = request.AccountNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUnionId)) {
		body["callerUnionId"] = request.CallerUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.CommodityType)) {
		body["commodityType"] = request.CommodityType
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
	_result = &UpdateCloudAccountInformationResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateCloudAccountInformation"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/yida/apps/cloudAccountInfos"), tea.String("json"), req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.FormInstanceId)) {
		body["formInstanceId"] = request.FormInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateFormDataJson)) {
		body["updateFormDataJson"] = request.UpdateFormDataJson
	}

	if !tea.BoolValue(util.IsUnset(request.UseLatestVersion)) {
		body["useLatestVersion"] = request.UseLatestVersion
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
	_result = &UpdateFormDataResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateFormData"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/yida/forms/instances"), tea.String("none"), req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		body["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateFormDataJson)) {
		body["updateFormDataJson"] = request.UpdateFormDataJson
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
	_result = &UpdateInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateInstance"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/yida/processes/instances"), tea.String("none"), req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.ErrorLines)) {
		body["errorLines"] = request.ErrorLines
	}

	if !tea.BoolValue(util.IsUnset(request.ImportSequence)) {
		body["importSequence"] = request.ImportSequence
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.SystemToken)) {
		body["systemToken"] = request.SystemToken
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
	_result = &UpdateStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateStatus"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/yida/forms/status"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeTenantInformation(request *UpgradeTenantInformationRequest) (_result *UpgradeTenantInformationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpgradeTenantInformationHeaders{}
	_result = &UpgradeTenantInformationResponse{}
	_body, _err := client.UpgradeTenantInformationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeTenantInformationWithOptions(request *UpgradeTenantInformationRequest, headers *UpgradeTenantInformationHeaders, runtime *util.RuntimeOptions) (_result *UpgradeTenantInformationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		body["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.AccountNumber)) {
		body["accountNumber"] = request.AccountNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUnionId)) {
		body["callerUnionId"] = request.CallerUnionId
	}

	if !tea.BoolValue(util.IsUnset(request.CommodityType)) {
		body["commodityType"] = request.CommodityType
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
	_result = &UpgradeTenantInformationResponse{}
	_body, _err := client.DoROARequest(tea.String("UpgradeTenantInformation"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("PUT"), tea.String("AK"), tea.String("/v1.0/yida/apps/tenantInfos"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateApplicationAuthorizationOrder(instanceId *string, request *ValidateApplicationAuthorizationOrderRequest) (_result *ValidateApplicationAuthorizationOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ValidateApplicationAuthorizationOrderHeaders{}
	_result = &ValidateApplicationAuthorizationOrderResponse{}
	_body, _err := client.ValidateApplicationAuthorizationOrderWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateApplicationAuthorizationOrderWithOptions(instanceId *string, request *ValidateApplicationAuthorizationOrderRequest, headers *ValidateApplicationAuthorizationOrderHeaders, runtime *util.RuntimeOptions) (_result *ValidateApplicationAuthorizationOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	instanceId = openapiutil.GetEncodeParam(instanceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUnionId)) {
		query["callerUnionId"] = request.CallerUnionId
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
	_result = &ValidateApplicationAuthorizationOrderResponse{}
	_body, _err := client.DoROARequest(tea.String("ValidateApplicationAuthorizationOrder"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/applicationOrderUpdateAuthorizations/"+tea.StringValue(instanceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateApplicationAuthorizationServiceOrder(callerUid *string, request *ValidateApplicationAuthorizationServiceOrderRequest) (_result *ValidateApplicationAuthorizationServiceOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ValidateApplicationAuthorizationServiceOrderHeaders{}
	_result = &ValidateApplicationAuthorizationServiceOrderResponse{}
	_body, _err := client.ValidateApplicationAuthorizationServiceOrderWithOptions(callerUid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateApplicationAuthorizationServiceOrderWithOptions(callerUid *string, request *ValidateApplicationAuthorizationServiceOrderRequest, headers *ValidateApplicationAuthorizationServiceOrderHeaders, runtime *util.RuntimeOptions) (_result *ValidateApplicationAuthorizationServiceOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	callerUid = openapiutil.GetEncodeParam(callerUid)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
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
	_result = &ValidateApplicationAuthorizationServiceOrderResponse{}
	_body, _err := client.DoROARequest(tea.String("ValidateApplicationAuthorizationServiceOrder"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/appsAuthorizations/freshOrderInfoReviews/"+tea.StringValue(callerUid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateApplicationServiceOrderUpgrade(callerUnionid *string, request *ValidateApplicationServiceOrderUpgradeRequest) (_result *ValidateApplicationServiceOrderUpgradeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ValidateApplicationServiceOrderUpgradeHeaders{}
	_result = &ValidateApplicationServiceOrderUpgradeResponse{}
	_body, _err := client.ValidateApplicationServiceOrderUpgradeWithOptions(callerUnionid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateApplicationServiceOrderUpgradeWithOptions(callerUnionid *string, request *ValidateApplicationServiceOrderUpgradeRequest, headers *ValidateApplicationServiceOrderUpgradeHeaders, runtime *util.RuntimeOptions) (_result *ValidateApplicationServiceOrderUpgradeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	callerUnionid = openapiutil.GetEncodeParam(callerUnionid)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
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
	_result = &ValidateApplicationServiceOrderUpgradeResponse{}
	_body, _err := client.DoROARequest(tea.String("ValidateApplicationServiceOrderUpgrade"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/applications/orderValidations/"+tea.StringValue(callerUnionid)), tea.String("json"), req, runtime)
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
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

func (client *Client) ValidateOrderUpdate(instanceId *string, request *ValidateOrderUpdateRequest) (_result *ValidateOrderUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ValidateOrderUpdateHeaders{}
	_result = &ValidateOrderUpdateResponse{}
	_body, _err := client.ValidateOrderUpdateWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateOrderUpdateWithOptions(instanceId *string, request *ValidateOrderUpdateRequest, headers *ValidateOrderUpdateHeaders, runtime *util.RuntimeOptions) (_result *ValidateOrderUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	instanceId = openapiutil.GetEncodeParam(instanceId)
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &ValidateOrderUpdateResponse{}
	_body, _err := client.DoROARequest(tea.String("ValidateOrderUpdate"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/orders/renewalReviews/"+tea.StringValue(instanceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
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
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		query["accessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.CallerUid)) {
		query["callerUid"] = request.CallerUid
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
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
	_result = &ValidateOrderUpgradeResponse{}
	_body, _err := client.DoROARequest(tea.String("ValidateOrderUpgrade"), tea.String("yida_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/yida/apps/orderUpgrade/validate"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
