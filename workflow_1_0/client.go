// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package workflow_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type FormComponent struct {
	// 控件类型
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// 控件属性
	Props *FormComponentProps `json:"props,omitempty" xml:"props,omitempty"`
	// 子控件集合
	Children []*FormComponent `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
}

func (s FormComponent) String() string {
	return tea.Prettify(s)
}

func (s FormComponent) GoString() string {
	return s.String()
}

func (s *FormComponent) SetComponentType(v string) *FormComponent {
	s.ComponentType = &v
	return s
}

func (s *FormComponent) SetProps(v *FormComponentProps) *FormComponent {
	s.Props = v
	return s
}

func (s *FormComponent) SetChildren(v []*FormComponent) *FormComponent {
	s.Children = v
	return s
}

type AvaliableTemplate struct {
	// 表单名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 表单模板processCode
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
}

func (s AvaliableTemplate) String() string {
	return tea.Prettify(s)
}

func (s AvaliableTemplate) GoString() string {
	return s.String()
}

func (s *AvaliableTemplate) SetName(v string) *AvaliableTemplate {
	s.Name = &v
	return s
}

func (s *AvaliableTemplate) SetProcessCode(v string) *AvaliableTemplate {
	s.ProcessCode = &v
	return s
}

type FormComponentProps struct {
	// 控件表单内唯一id
	ComponentId *string `json:"componentId,omitempty" xml:"componentId,omitempty"`
	// 控件标题
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// 套件中控件是否可设置为分条件字段
	AsyncCondition *bool `json:"asyncCondition,omitempty" xml:"asyncCondition,omitempty"`
	// 是否必填
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// 说明文字控件内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 时间格式
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// 金额控件是否需要大写，1不需要大写，其他需要大写
	Upper *string `json:"upper,omitempty" xml:"upper,omitempty"`
	// 时间单位（天、小时）
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 输入提示
	Placeholder *string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// 业务别名
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// 套件的业务标识
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 是否自动计算时长
	Duration *bool `json:"duration,omitempty" xml:"duration,omitempty"`
	// 联系人控件是否支持多选，1多选，0单选
	Choice *string `json:"choice,omitempty" xml:"choice,omitempty"`
	// 是否不可编辑
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// 文字提示控件显示方式:top|middle|bottom
	Align *string `json:"align,omitempty" xml:"align,omitempty"`
	// 是否隐藏字段
	Invisible *bool `json:"invisible,omitempty" xml:"invisible,omitempty"`
	// 说明文字控件链接地址
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// 明细打印方式false横向，true纵向
	VerticalPrint *bool `json:"verticalPrint,omitempty" xml:"verticalPrint,omitempty"`
	// 公式
	Formula *string `json:"formula,omitempty" xml:"formula,omitempty"`
	// 自定义控件渲染标识
	CommonBizType *string `json:"commonBizType,omitempty" xml:"commonBizType,omitempty"`
	// 单选多选控件选项列表
	Options []*SelectOption `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// 字段是否可打印，1打印，0不打印，默认打印
	Print *string `json:"print,omitempty" xml:"print,omitempty"`
	// 明细控件数据汇总统计
	StatField []*FormComponentPropsStatField `json:"statField,omitempty" xml:"statField,omitempty" type:"Repeated"`
	// 关联数据源配置
	DataSource *FormDataSource `json:"dataSource,omitempty" xml:"dataSource,omitempty"`
	// 地址控件模式city省市,district省市区,street省市区街道
	AddressModel *string `json:"addressModel,omitempty" xml:"addressModel,omitempty"`
	// 部门控件是否可多选
	Multiple *bool `json:"multiple,omitempty" xml:"multiple,omitempty"`
	// 评分控件上限
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 关联审批单控件限定模板列表
	AvailableTemplates []*AvaliableTemplate `json:"availableTemplates,omitempty" xml:"availableTemplates,omitempty" type:"Repeated"`
	// 明细填写方式，table（表格）、list（列表）
	TableViewMode *string `json:"tableViewMode,omitempty" xml:"tableViewMode,omitempty"`
}

func (s FormComponentProps) String() string {
	return tea.Prettify(s)
}

func (s FormComponentProps) GoString() string {
	return s.String()
}

func (s *FormComponentProps) SetComponentId(v string) *FormComponentProps {
	s.ComponentId = &v
	return s
}

func (s *FormComponentProps) SetLabel(v string) *FormComponentProps {
	s.Label = &v
	return s
}

func (s *FormComponentProps) SetAsyncCondition(v bool) *FormComponentProps {
	s.AsyncCondition = &v
	return s
}

func (s *FormComponentProps) SetRequired(v bool) *FormComponentProps {
	s.Required = &v
	return s
}

func (s *FormComponentProps) SetContent(v string) *FormComponentProps {
	s.Content = &v
	return s
}

func (s *FormComponentProps) SetFormat(v string) *FormComponentProps {
	s.Format = &v
	return s
}

func (s *FormComponentProps) SetUpper(v string) *FormComponentProps {
	s.Upper = &v
	return s
}

func (s *FormComponentProps) SetUnit(v string) *FormComponentProps {
	s.Unit = &v
	return s
}

func (s *FormComponentProps) SetPlaceholder(v string) *FormComponentProps {
	s.Placeholder = &v
	return s
}

func (s *FormComponentProps) SetBizAlias(v string) *FormComponentProps {
	s.BizAlias = &v
	return s
}

func (s *FormComponentProps) SetBizType(v string) *FormComponentProps {
	s.BizType = &v
	return s
}

func (s *FormComponentProps) SetDuration(v bool) *FormComponentProps {
	s.Duration = &v
	return s
}

func (s *FormComponentProps) SetChoice(v string) *FormComponentProps {
	s.Choice = &v
	return s
}

func (s *FormComponentProps) SetDisabled(v bool) *FormComponentProps {
	s.Disabled = &v
	return s
}

func (s *FormComponentProps) SetAlign(v string) *FormComponentProps {
	s.Align = &v
	return s
}

func (s *FormComponentProps) SetInvisible(v bool) *FormComponentProps {
	s.Invisible = &v
	return s
}

func (s *FormComponentProps) SetLink(v string) *FormComponentProps {
	s.Link = &v
	return s
}

func (s *FormComponentProps) SetVerticalPrint(v bool) *FormComponentProps {
	s.VerticalPrint = &v
	return s
}

func (s *FormComponentProps) SetFormula(v string) *FormComponentProps {
	s.Formula = &v
	return s
}

func (s *FormComponentProps) SetCommonBizType(v string) *FormComponentProps {
	s.CommonBizType = &v
	return s
}

func (s *FormComponentProps) SetOptions(v []*SelectOption) *FormComponentProps {
	s.Options = v
	return s
}

func (s *FormComponentProps) SetPrint(v string) *FormComponentProps {
	s.Print = &v
	return s
}

func (s *FormComponentProps) SetStatField(v []*FormComponentPropsStatField) *FormComponentProps {
	s.StatField = v
	return s
}

func (s *FormComponentProps) SetDataSource(v *FormDataSource) *FormComponentProps {
	s.DataSource = v
	return s
}

func (s *FormComponentProps) SetAddressModel(v string) *FormComponentProps {
	s.AddressModel = &v
	return s
}

func (s *FormComponentProps) SetMultiple(v bool) *FormComponentProps {
	s.Multiple = &v
	return s
}

func (s *FormComponentProps) SetLimit(v int32) *FormComponentProps {
	s.Limit = &v
	return s
}

func (s *FormComponentProps) SetAvailableTemplates(v []*AvaliableTemplate) *FormComponentProps {
	s.AvailableTemplates = v
	return s
}

func (s *FormComponentProps) SetTableViewMode(v string) *FormComponentProps {
	s.TableViewMode = &v
	return s
}

type FormComponentPropsStatField struct {
	// 需要统计的明细控件内子控件id
	ComponentId *string `json:"componentId,omitempty" xml:"componentId,omitempty"`
	// 子控件标题
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// 金额控件是否需要大写，1不需要大写，其他需要大写
	Upper *string `json:"upper,omitempty" xml:"upper,omitempty"`
}

func (s FormComponentPropsStatField) String() string {
	return tea.Prettify(s)
}

func (s FormComponentPropsStatField) GoString() string {
	return s.String()
}

func (s *FormComponentPropsStatField) SetComponentId(v string) *FormComponentPropsStatField {
	s.ComponentId = &v
	return s
}

func (s *FormComponentPropsStatField) SetLabel(v string) *FormComponentPropsStatField {
	s.Label = &v
	return s
}

func (s *FormComponentPropsStatField) SetUpper(v string) *FormComponentPropsStatField {
	s.Upper = &v
	return s
}

type FormDataSource struct {
	// 关联类型，form关联表单
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 关联表单信息
	Target *FormDataSourceTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
}

func (s FormDataSource) String() string {
	return tea.Prettify(s)
}

func (s FormDataSource) GoString() string {
	return s.String()
}

func (s *FormDataSource) SetType(v string) *FormDataSource {
	s.Type = &v
	return s
}

func (s *FormDataSource) SetTarget(v *FormDataSourceTarget) *FormDataSource {
	s.Target = v
	return s
}

type FormDataSourceTarget struct {
	// 应用appUuid
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// 表单类型，0流程表单
	AppType *int32 `json:"appType,omitempty" xml:"appType,omitempty"`
	// 关联表单业务标识
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 关联表单的formCode
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
}

func (s FormDataSourceTarget) String() string {
	return tea.Prettify(s)
}

func (s FormDataSourceTarget) GoString() string {
	return s.String()
}

func (s *FormDataSourceTarget) SetAppUuid(v string) *FormDataSourceTarget {
	s.AppUuid = &v
	return s
}

func (s *FormDataSourceTarget) SetAppType(v int32) *FormDataSourceTarget {
	s.AppType = &v
	return s
}

func (s *FormDataSourceTarget) SetBizType(v string) *FormDataSourceTarget {
	s.BizType = &v
	return s
}

func (s *FormDataSourceTarget) SetFormCode(v string) *FormDataSourceTarget {
	s.FormCode = &v
	return s
}

type SelectOption struct {
	// 每一个选项的唯一键
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// 每一个选项的值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s SelectOption) String() string {
	return tea.Prettify(s)
}

func (s SelectOption) GoString() string {
	return s.String()
}

func (s *SelectOption) SetKey(v string) *SelectOption {
	s.Key = &v
	return s
}

func (s *SelectOption) SetValue(v string) *SelectOption {
	s.Value = &v
	return s
}

type QueryFormInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryFormInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryFormInstanceHeaders) GoString() string {
	return s.String()
}

func (s *QueryFormInstanceHeaders) SetCommonHeaders(v map[string]*string) *QueryFormInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryFormInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *QueryFormInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryFormInstanceRequest struct {
	// 表单实例id
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// 表单模板Code
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// 应用搭建id
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
}

func (s QueryFormInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFormInstanceRequest) GoString() string {
	return s.String()
}

func (s *QueryFormInstanceRequest) SetFormInstanceId(v string) *QueryFormInstanceRequest {
	s.FormInstanceId = &v
	return s
}

func (s *QueryFormInstanceRequest) SetFormCode(v string) *QueryFormInstanceRequest {
	s.FormCode = &v
	return s
}

func (s *QueryFormInstanceRequest) SetAppUuid(v string) *QueryFormInstanceRequest {
	s.AppUuid = &v
	return s
}

type QueryFormInstanceResponseBody struct {
	// 实例id
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// 表单数据
	FormInstDataList []*QueryFormInstanceResponseBodyFormInstDataList `json:"formInstDataList,omitempty" xml:"formInstDataList,omitempty" type:"Repeated"`
	// 应用搭建id
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// 表单模板id
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// 表单标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 创建人
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 修改人
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// 实例创建时间戳
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// 实例最近修改时间戳
	ModifyTimestamp *int64 `json:"modifyTimestamp,omitempty" xml:"modifyTimestamp,omitempty"`
	// 外联业务实例id
	OutInstanceId *string `json:"outInstanceId,omitempty" xml:"outInstanceId,omitempty"`
	// 外联业务code
	OutBizCode *string `json:"outBizCode,omitempty" xml:"outBizCode,omitempty"`
	// 扩展信息
	Attributes map[string]interface{} `json:"attributes,omitempty" xml:"attributes,omitempty"`
}

func (s QueryFormInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFormInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFormInstanceResponseBody) SetFormInstanceId(v string) *QueryFormInstanceResponseBody {
	s.FormInstanceId = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetFormInstDataList(v []*QueryFormInstanceResponseBodyFormInstDataList) *QueryFormInstanceResponseBody {
	s.FormInstDataList = v
	return s
}

func (s *QueryFormInstanceResponseBody) SetAppUuid(v string) *QueryFormInstanceResponseBody {
	s.AppUuid = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetFormCode(v string) *QueryFormInstanceResponseBody {
	s.FormCode = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetTitle(v string) *QueryFormInstanceResponseBody {
	s.Title = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetCreator(v string) *QueryFormInstanceResponseBody {
	s.Creator = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetModifier(v string) *QueryFormInstanceResponseBody {
	s.Modifier = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetCreateTimestamp(v int64) *QueryFormInstanceResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetModifyTimestamp(v int64) *QueryFormInstanceResponseBody {
	s.ModifyTimestamp = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetOutInstanceId(v string) *QueryFormInstanceResponseBody {
	s.OutInstanceId = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetOutBizCode(v string) *QueryFormInstanceResponseBody {
	s.OutBizCode = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetAttributes(v map[string]interface{}) *QueryFormInstanceResponseBody {
	s.Attributes = v
	return s
}

type QueryFormInstanceResponseBodyFormInstDataList struct {
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	BizAlias      *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	ExtendValue   *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	Label         *string `json:"label,omitempty" xml:"label,omitempty"`
	Value         *string `json:"value,omitempty" xml:"value,omitempty"`
	Key           *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s QueryFormInstanceResponseBodyFormInstDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryFormInstanceResponseBodyFormInstDataList) GoString() string {
	return s.String()
}

func (s *QueryFormInstanceResponseBodyFormInstDataList) SetComponentType(v string) *QueryFormInstanceResponseBodyFormInstDataList {
	s.ComponentType = &v
	return s
}

func (s *QueryFormInstanceResponseBodyFormInstDataList) SetBizAlias(v string) *QueryFormInstanceResponseBodyFormInstDataList {
	s.BizAlias = &v
	return s
}

func (s *QueryFormInstanceResponseBodyFormInstDataList) SetExtendValue(v string) *QueryFormInstanceResponseBodyFormInstDataList {
	s.ExtendValue = &v
	return s
}

func (s *QueryFormInstanceResponseBodyFormInstDataList) SetLabel(v string) *QueryFormInstanceResponseBodyFormInstDataList {
	s.Label = &v
	return s
}

func (s *QueryFormInstanceResponseBodyFormInstDataList) SetValue(v string) *QueryFormInstanceResponseBodyFormInstDataList {
	s.Value = &v
	return s
}

func (s *QueryFormInstanceResponseBodyFormInstDataList) SetKey(v string) *QueryFormInstanceResponseBodyFormInstDataList {
	s.Key = &v
	return s
}

type QueryFormInstanceResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryFormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryFormInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFormInstanceResponse) GoString() string {
	return s.String()
}

func (s *QueryFormInstanceResponse) SetHeaders(v map[string]*string) *QueryFormInstanceResponse {
	s.Headers = v
	return s
}

func (s *QueryFormInstanceResponse) SetBody(v *QueryFormInstanceResponseBody) *QueryFormInstanceResponse {
	s.Body = v
	return s
}

type ProcessForecastHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ProcessForecastHeaders) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastHeaders) GoString() string {
	return s.String()
}

func (s *ProcessForecastHeaders) SetCommonHeaders(v map[string]*string) *ProcessForecastHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ProcessForecastHeaders) SetXAcsDingtalkAccessToken(v string) *ProcessForecastHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ProcessForecastRequest struct {
	DingCorpId         *string `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 审批流的唯一码
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// 部门ID
	DeptId *int32 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 审批发起人的userId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 表单数据内容，控件列表
	FormComponentValues []*ProcessForecastRequestFormComponentValues `json:"formComponentValues,omitempty" xml:"formComponentValues,omitempty" type:"Repeated"`
}

func (s ProcessForecastRequest) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastRequest) GoString() string {
	return s.String()
}

func (s *ProcessForecastRequest) SetDingCorpId(v string) *ProcessForecastRequest {
	s.DingCorpId = &v
	return s
}

func (s *ProcessForecastRequest) SetDingOrgId(v int64) *ProcessForecastRequest {
	s.DingOrgId = &v
	return s
}

func (s *ProcessForecastRequest) SetDingIsvOrgId(v int64) *ProcessForecastRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *ProcessForecastRequest) SetDingSuiteKey(v string) *ProcessForecastRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *ProcessForecastRequest) SetDingTokenGrantType(v int64) *ProcessForecastRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *ProcessForecastRequest) SetRequestId(v string) *ProcessForecastRequest {
	s.RequestId = &v
	return s
}

func (s *ProcessForecastRequest) SetProcessCode(v string) *ProcessForecastRequest {
	s.ProcessCode = &v
	return s
}

func (s *ProcessForecastRequest) SetDeptId(v int32) *ProcessForecastRequest {
	s.DeptId = &v
	return s
}

func (s *ProcessForecastRequest) SetUserId(v string) *ProcessForecastRequest {
	s.UserId = &v
	return s
}

func (s *ProcessForecastRequest) SetFormComponentValues(v []*ProcessForecastRequestFormComponentValues) *ProcessForecastRequest {
	s.FormComponentValues = v
	return s
}

type ProcessForecastRequestFormComponentValues struct {
	// 控件id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 控件别名
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// 控件名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 控件值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 控件扩展值
	ExtValue *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	// 控件类型
	ComponentType *string                                             `json:"componentType,omitempty" xml:"componentType,omitempty"`
	Details       []*ProcessForecastRequestFormComponentValuesDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
}

func (s ProcessForecastRequestFormComponentValues) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastRequestFormComponentValues) GoString() string {
	return s.String()
}

func (s *ProcessForecastRequestFormComponentValues) SetId(v string) *ProcessForecastRequestFormComponentValues {
	s.Id = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValues) SetBizAlias(v string) *ProcessForecastRequestFormComponentValues {
	s.BizAlias = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValues) SetName(v string) *ProcessForecastRequestFormComponentValues {
	s.Name = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValues) SetValue(v string) *ProcessForecastRequestFormComponentValues {
	s.Value = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValues) SetExtValue(v string) *ProcessForecastRequestFormComponentValues {
	s.ExtValue = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValues) SetComponentType(v string) *ProcessForecastRequestFormComponentValues {
	s.ComponentType = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValues) SetDetails(v []*ProcessForecastRequestFormComponentValuesDetails) *ProcessForecastRequestFormComponentValues {
	s.Details = v
	return s
}

type ProcessForecastRequestFormComponentValuesDetails struct {
	// 控件id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 控件别名
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// 控件名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 控件值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 控件扩展值
	ExtValue *string                                                    `json:"extValue,omitempty" xml:"extValue,omitempty"`
	Details  []*ProcessForecastRequestFormComponentValuesDetailsDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
}

func (s ProcessForecastRequestFormComponentValuesDetails) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastRequestFormComponentValuesDetails) GoString() string {
	return s.String()
}

func (s *ProcessForecastRequestFormComponentValuesDetails) SetId(v string) *ProcessForecastRequestFormComponentValuesDetails {
	s.Id = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValuesDetails) SetBizAlias(v string) *ProcessForecastRequestFormComponentValuesDetails {
	s.BizAlias = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValuesDetails) SetName(v string) *ProcessForecastRequestFormComponentValuesDetails {
	s.Name = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValuesDetails) SetValue(v string) *ProcessForecastRequestFormComponentValuesDetails {
	s.Value = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValuesDetails) SetExtValue(v string) *ProcessForecastRequestFormComponentValuesDetails {
	s.ExtValue = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValuesDetails) SetDetails(v []*ProcessForecastRequestFormComponentValuesDetailsDetails) *ProcessForecastRequestFormComponentValuesDetails {
	s.Details = v
	return s
}

type ProcessForecastRequestFormComponentValuesDetailsDetails struct {
	// 控件id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 控件别名
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// 控件名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 控件值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 控件扩展值
	ExtValue      *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
}

func (s ProcessForecastRequestFormComponentValuesDetailsDetails) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastRequestFormComponentValuesDetailsDetails) GoString() string {
	return s.String()
}

func (s *ProcessForecastRequestFormComponentValuesDetailsDetails) SetId(v string) *ProcessForecastRequestFormComponentValuesDetailsDetails {
	s.Id = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValuesDetailsDetails) SetBizAlias(v string) *ProcessForecastRequestFormComponentValuesDetailsDetails {
	s.BizAlias = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValuesDetailsDetails) SetName(v string) *ProcessForecastRequestFormComponentValuesDetailsDetails {
	s.Name = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValuesDetailsDetails) SetValue(v string) *ProcessForecastRequestFormComponentValuesDetailsDetails {
	s.Value = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValuesDetailsDetails) SetExtValue(v string) *ProcessForecastRequestFormComponentValuesDetailsDetails {
	s.ExtValue = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValuesDetailsDetails) SetComponentType(v string) *ProcessForecastRequestFormComponentValuesDetailsDetails {
	s.ComponentType = &v
	return s
}

type ProcessForecastResponseBody struct {
	// 返回结果
	Result *ProcessForecastResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ProcessForecastResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastResponseBody) GoString() string {
	return s.String()
}

func (s *ProcessForecastResponseBody) SetResult(v *ProcessForecastResponseBodyResult) *ProcessForecastResponseBody {
	s.Result = v
	return s
}

type ProcessForecastResponseBodyResult struct {
	// 是否预测成功
	IsForecastSuccess *bool `json:"isForecastSuccess,omitempty" xml:"isForecastSuccess,omitempty"`
	// 流程 code
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// 用户 id
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// 流程 id
	ProcessId *int64 `json:"processId,omitempty" xml:"processId,omitempty"`
	// 是否静态流程
	IsStaticWorkflow      *bool                                                     `json:"isStaticWorkflow,omitempty" xml:"isStaticWorkflow,omitempty"`
	WorkflowActivityRules []*ProcessForecastResponseBodyResultWorkflowActivityRules `json:"workflowActivityRules,omitempty" xml:"workflowActivityRules,omitempty" type:"Repeated"`
	WorkflowForecastNodes []*ProcessForecastResponseBodyResultWorkflowForecastNodes `json:"workflowForecastNodes,omitempty" xml:"workflowForecastNodes,omitempty" type:"Repeated"`
}

func (s ProcessForecastResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ProcessForecastResponseBodyResult) SetIsForecastSuccess(v bool) *ProcessForecastResponseBodyResult {
	s.IsForecastSuccess = &v
	return s
}

func (s *ProcessForecastResponseBodyResult) SetProcessCode(v string) *ProcessForecastResponseBodyResult {
	s.ProcessCode = &v
	return s
}

func (s *ProcessForecastResponseBodyResult) SetUserId(v string) *ProcessForecastResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *ProcessForecastResponseBodyResult) SetProcessId(v int64) *ProcessForecastResponseBodyResult {
	s.ProcessId = &v
	return s
}

func (s *ProcessForecastResponseBodyResult) SetIsStaticWorkflow(v bool) *ProcessForecastResponseBodyResult {
	s.IsStaticWorkflow = &v
	return s
}

func (s *ProcessForecastResponseBodyResult) SetWorkflowActivityRules(v []*ProcessForecastResponseBodyResultWorkflowActivityRules) *ProcessForecastResponseBodyResult {
	s.WorkflowActivityRules = v
	return s
}

func (s *ProcessForecastResponseBodyResult) SetWorkflowForecastNodes(v []*ProcessForecastResponseBodyResultWorkflowForecastNodes) *ProcessForecastResponseBodyResult {
	s.WorkflowForecastNodes = v
	return s
}

type ProcessForecastResponseBodyResultWorkflowActivityRules struct {
	// 节点 id
	ActivityId *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
	// 流程中前一个节点的 id
	PrevActivityId *string `json:"prevActivityId,omitempty" xml:"prevActivityId,omitempty"`
	// 节点名称
	ActivityName *string `json:"activityName,omitempty" xml:"activityName,omitempty"`
	// 规则类型
	ActivityType *string `json:"activityType,omitempty" xml:"activityType,omitempty"`
	// 是否自选审批节点
	IsTargetSelect *bool `json:"isTargetSelect,omitempty" xml:"isTargetSelect,omitempty"`
	// 节点操作人信息
	WorkflowActor *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor `json:"workflowActor,omitempty" xml:"workflowActor,omitempty" type:"Struct"`
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRules) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRules) GoString() string {
	return s.String()
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRules) SetActivityId(v string) *ProcessForecastResponseBodyResultWorkflowActivityRules {
	s.ActivityId = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRules) SetPrevActivityId(v string) *ProcessForecastResponseBodyResultWorkflowActivityRules {
	s.PrevActivityId = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRules) SetActivityName(v string) *ProcessForecastResponseBodyResultWorkflowActivityRules {
	s.ActivityName = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRules) SetActivityType(v string) *ProcessForecastResponseBodyResultWorkflowActivityRules {
	s.ActivityType = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRules) SetIsTargetSelect(v bool) *ProcessForecastResponseBodyResultWorkflowActivityRules {
	s.IsTargetSelect = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRules) SetWorkflowActor(v *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) *ProcessForecastResponseBodyResultWorkflowActivityRules {
	s.WorkflowActor = v
	return s
}

type ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor struct {
	// 节点操作人 key
	ActorKey *string `json:"actorKey,omitempty" xml:"actorKey,omitempty"`
	// 节点操作人类型
	ActorType *string `json:"actorType,omitempty" xml:"actorType,omitempty"`
	// 节点操作人选择范围类型
	ActorSelectionType *string `json:"actorSelectionType,omitempty" xml:"actorSelectionType,omitempty"`
	// 节点操作人选择范围
	ActorSelectionRange *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRange `json:"actorSelectionRange,omitempty" xml:"actorSelectionRange,omitempty" type:"Struct"`
	// 是否允许多选，还是仅允许选一人
	AllowedMulti *bool `json:"allowedMulti,omitempty" xml:"allowedMulti,omitempty"`
	// 节点审批类型
	ApprovalType *string `json:"approvalType,omitempty" xml:"approvalType,omitempty"`
	// 节点审批方式
	ApprovalMethod *string `json:"approvalMethod,omitempty" xml:"approvalMethod,omitempty"`
	// 节点激活类型
	ActorActivateType *string `json:"actorActivateType,omitempty" xml:"actorActivateType,omitempty"`
	// 该审批人节点在发起审批时是否必填
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) GoString() string {
	return s.String()
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetActorKey(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.ActorKey = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetActorType(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.ActorType = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetActorSelectionType(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.ActorSelectionType = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetActorSelectionRange(v *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRange) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.ActorSelectionRange = v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetAllowedMulti(v bool) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.AllowedMulti = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetApprovalType(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.ApprovalType = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetApprovalMethod(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.ApprovalMethod = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetActorActivateType(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.ActorActivateType = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetRequired(v bool) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.Required = &v
	return s
}

type ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRange struct {
	// 审批指定成员
	Approvals []*ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeApprovals `json:"approvals,omitempty" xml:"approvals,omitempty" type:"Repeated"`
	// 审批指定角色
	Labels []*ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeLabels `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRange) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRange) GoString() string {
	return s.String()
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRange) SetApprovals(v []*ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeApprovals) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRange {
	s.Approvals = v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRange) SetLabels(v []*ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeLabels) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRange {
	s.Labels = v
	return s
}

type ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeApprovals struct {
	// 员工 userId
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
	// 员工姓名
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeApprovals) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeApprovals) GoString() string {
	return s.String()
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeApprovals) SetWorkNo(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeApprovals {
	s.WorkNo = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeApprovals) SetUserName(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeApprovals {
	s.UserName = &v
	return s
}

type ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeLabels struct {
	// 角色 id
	Labels *string `json:"labels,omitempty" xml:"labels,omitempty"`
	// 角色名字
	LabelNames *string `json:"labelNames,omitempty" xml:"labelNames,omitempty"`
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeLabels) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeLabels) GoString() string {
	return s.String()
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeLabels) SetLabels(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeLabels {
	s.Labels = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeLabels) SetLabelNames(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeLabels {
	s.LabelNames = &v
	return s
}

type ProcessForecastResponseBodyResultWorkflowForecastNodes struct {
	// 节点 id
	ActivityId *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
	// 节点出线 id
	OutId *string `json:"outId,omitempty" xml:"outId,omitempty"`
}

func (s ProcessForecastResponseBodyResultWorkflowForecastNodes) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastResponseBodyResultWorkflowForecastNodes) GoString() string {
	return s.String()
}

func (s *ProcessForecastResponseBodyResultWorkflowForecastNodes) SetActivityId(v string) *ProcessForecastResponseBodyResultWorkflowForecastNodes {
	s.ActivityId = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowForecastNodes) SetOutId(v string) *ProcessForecastResponseBodyResultWorkflowForecastNodes {
	s.OutId = &v
	return s
}

type ProcessForecastResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ProcessForecastResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ProcessForecastResponse) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastResponse) GoString() string {
	return s.String()
}

func (s *ProcessForecastResponse) SetHeaders(v map[string]*string) *ProcessForecastResponse {
	s.Headers = v
	return s
}

func (s *ProcessForecastResponse) SetBody(v *ProcessForecastResponseBody) *ProcessForecastResponse {
	s.Body = v
	return s
}

type QueryAllProcessInstancesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllProcessInstancesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesHeaders) SetCommonHeaders(v map[string]*string) *QueryAllProcessInstancesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllProcessInstancesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllProcessInstancesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllProcessInstancesRequest struct {
	// 分页起始值
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 分页大小
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 开始时间
	StartTimeInMills *int64 `json:"startTimeInMills,omitempty" xml:"startTimeInMills,omitempty"`
	// 结束时间
	EndTimeInMills *int64 `json:"endTimeInMills,omitempty" xml:"endTimeInMills,omitempty"`
	// 模板编码
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// 应用编码
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
}

func (s QueryAllProcessInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesRequest) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesRequest) SetNextToken(v string) *QueryAllProcessInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *QueryAllProcessInstancesRequest) SetMaxResults(v int64) *QueryAllProcessInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryAllProcessInstancesRequest) SetStartTimeInMills(v int64) *QueryAllProcessInstancesRequest {
	s.StartTimeInMills = &v
	return s
}

func (s *QueryAllProcessInstancesRequest) SetEndTimeInMills(v int64) *QueryAllProcessInstancesRequest {
	s.EndTimeInMills = &v
	return s
}

func (s *QueryAllProcessInstancesRequest) SetProcessCode(v string) *QueryAllProcessInstancesRequest {
	s.ProcessCode = &v
	return s
}

func (s *QueryAllProcessInstancesRequest) SetAppUuid(v string) *QueryAllProcessInstancesRequest {
	s.AppUuid = &v
	return s
}

type QueryAllProcessInstancesResponseBody struct {
	// result
	Result *QueryAllProcessInstancesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryAllProcessInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesResponseBody) SetResult(v *QueryAllProcessInstancesResponseBodyResult) *QueryAllProcessInstancesResponseBody {
	s.Result = v
	return s
}

type QueryAllProcessInstancesResponseBodyResult struct {
	// 下次获取数据的游标
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 是否有更多数据
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 总数
	MaxResults *int64                                            `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	List       []*QueryAllProcessInstancesResponseBodyResultList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s QueryAllProcessInstancesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesResponseBodyResult) SetNextToken(v string) *QueryAllProcessInstancesResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResult) SetHasMore(v bool) *QueryAllProcessInstancesResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResult) SetMaxResults(v int64) *QueryAllProcessInstancesResponseBodyResult {
	s.MaxResults = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResult) SetList(v []*QueryAllProcessInstancesResponseBodyResultList) *QueryAllProcessInstancesResponseBodyResult {
	s.List = v
	return s
}

type QueryAllProcessInstancesResponseBodyResultList struct {
	// 流程实例ID
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// 主单实例Id
	MainProcessInstanceId *string `json:"mainProcessInstanceId,omitempty" xml:"mainProcessInstanceId,omitempty"`
	// 审批结束时间
	FinishTime *int64 `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// 附属单信息
	AttachedProcessInstanceIds *string `json:"attachedProcessInstanceIds,omitempty" xml:"attachedProcessInstanceIds,omitempty"`
	// 审批单编号
	BusinessId *string `json:"businessId,omitempty" xml:"businessId,omitempty"`
	// 审批单标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 发起人部门id
	OriginatorDeptId *string `json:"originatorDeptId,omitempty" xml:"originatorDeptId,omitempty"`
	// 审批结果
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// 审批单创建时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 发起者userId
	OriginatorUserid *string `json:"originatorUserid,omitempty" xml:"originatorUserid,omitempty"`
	// 审批单状态
	Status              *string                                                              `json:"status,omitempty" xml:"status,omitempty"`
	FormComponentValues []*QueryAllProcessInstancesResponseBodyResultListFormComponentValues `json:"formComponentValues,omitempty" xml:"formComponentValues,omitempty" type:"Repeated"`
}

func (s QueryAllProcessInstancesResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetProcessInstanceId(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.ProcessInstanceId = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetMainProcessInstanceId(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.MainProcessInstanceId = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetFinishTime(v int64) *QueryAllProcessInstancesResponseBodyResultList {
	s.FinishTime = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetAttachedProcessInstanceIds(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.AttachedProcessInstanceIds = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetBusinessId(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.BusinessId = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetTitle(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.Title = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetOriginatorDeptId(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.OriginatorDeptId = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetResult(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.Result = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetCreateTime(v int64) *QueryAllProcessInstancesResponseBodyResultList {
	s.CreateTime = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetOriginatorUserid(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.OriginatorUserid = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetStatus(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.Status = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetFormComponentValues(v []*QueryAllProcessInstancesResponseBodyResultListFormComponentValues) *QueryAllProcessInstancesResponseBodyResultList {
	s.FormComponentValues = v
	return s
}

type QueryAllProcessInstancesResponseBodyResultListFormComponentValues struct {
	// 控件名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 控件id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 控件数据
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 控件扩展数据
	ExtValue *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
}

func (s QueryAllProcessInstancesResponseBodyResultListFormComponentValues) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesResponseBodyResultListFormComponentValues) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesResponseBodyResultListFormComponentValues) SetName(v string) *QueryAllProcessInstancesResponseBodyResultListFormComponentValues {
	s.Name = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultListFormComponentValues) SetId(v string) *QueryAllProcessInstancesResponseBodyResultListFormComponentValues {
	s.Id = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultListFormComponentValues) SetValue(v string) *QueryAllProcessInstancesResponseBodyResultListFormComponentValues {
	s.Value = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultListFormComponentValues) SetExtValue(v string) *QueryAllProcessInstancesResponseBodyResultListFormComponentValues {
	s.ExtValue = &v
	return s
}

type QueryAllProcessInstancesResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAllProcessInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllProcessInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesResponse) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesResponse) SetHeaders(v map[string]*string) *QueryAllProcessInstancesResponse {
	s.Headers = v
	return s
}

func (s *QueryAllProcessInstancesResponse) SetBody(v *QueryAllProcessInstancesResponseBody) *QueryAllProcessInstancesResponse {
	s.Body = v
	return s
}

type QueryAllFormInstancesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryAllFormInstancesHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryAllFormInstancesHeaders) GoString() string {
	return s.String()
}

func (s *QueryAllFormInstancesHeaders) SetCommonHeaders(v map[string]*string) *QueryAllFormInstancesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryAllFormInstancesHeaders) SetXAcsDingtalkAccessToken(v string) *QueryAllFormInstancesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryAllFormInstancesRequest struct {
	// 分页游标，第一次调用传空或者null
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 翻页size
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 应用搭建id
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// 表单模板id
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
}

func (s QueryAllFormInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllFormInstancesRequest) GoString() string {
	return s.String()
}

func (s *QueryAllFormInstancesRequest) SetNextToken(v string) *QueryAllFormInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *QueryAllFormInstancesRequest) SetMaxResults(v int32) *QueryAllFormInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryAllFormInstancesRequest) SetAppUuid(v string) *QueryAllFormInstancesRequest {
	s.AppUuid = &v
	return s
}

func (s *QueryAllFormInstancesRequest) SetFormCode(v string) *QueryAllFormInstancesRequest {
	s.FormCode = &v
	return s
}

type QueryAllFormInstancesResponseBody struct {
	// 分页结果
	Result *QueryAllFormInstancesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryAllFormInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllFormInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllFormInstancesResponseBody) SetResult(v *QueryAllFormInstancesResponseBodyResult) *QueryAllFormInstancesResponseBody {
	s.Result = v
	return s
}

type QueryAllFormInstancesResponseBodyResult struct {
	// 下一页的游标
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 是否有更多数据
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 分页大小
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 表单列表
	Values []*QueryAllFormInstancesResponseBodyResultValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s QueryAllFormInstancesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryAllFormInstancesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryAllFormInstancesResponseBodyResult) SetNextToken(v string) *QueryAllFormInstancesResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResult) SetHasMore(v bool) *QueryAllFormInstancesResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResult) SetMaxResults(v int64) *QueryAllFormInstancesResponseBodyResult {
	s.MaxResults = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResult) SetValues(v []*QueryAllFormInstancesResponseBodyResultValues) *QueryAllFormInstancesResponseBodyResult {
	s.Values = v
	return s
}

type QueryAllFormInstancesResponseBodyResultValues struct {
	// 表单实例id
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// 应用搭建id
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// 表单模板code
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// 标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 创建人
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 修改人
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// 创建时间
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// 修改时间
	ModifyTimestamp *int64 `json:"modifyTimestamp,omitempty" xml:"modifyTimestamp,omitempty"`
	// 外部实例编码
	OutInstanceId *string `json:"outInstanceId,omitempty" xml:"outInstanceId,omitempty"`
	// 外部业务编码
	OutBizCode *string `json:"outBizCode,omitempty" xml:"outBizCode,omitempty"`
	// 扩展信息
	Attributes map[string]interface{} `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 表单实例数据
	FormInstDataList []*QueryAllFormInstancesResponseBodyResultValuesFormInstDataList `json:"formInstDataList,omitempty" xml:"formInstDataList,omitempty" type:"Repeated"`
}

func (s QueryAllFormInstancesResponseBodyResultValues) String() string {
	return tea.Prettify(s)
}

func (s QueryAllFormInstancesResponseBodyResultValues) GoString() string {
	return s.String()
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetFormInstanceId(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.FormInstanceId = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetAppUuid(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.AppUuid = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetFormCode(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.FormCode = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetTitle(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.Title = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetCreator(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.Creator = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetModifier(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.Modifier = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetCreateTimestamp(v int64) *QueryAllFormInstancesResponseBodyResultValues {
	s.CreateTimestamp = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetModifyTimestamp(v int64) *QueryAllFormInstancesResponseBodyResultValues {
	s.ModifyTimestamp = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetOutInstanceId(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.OutInstanceId = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetOutBizCode(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.OutBizCode = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetAttributes(v map[string]interface{}) *QueryAllFormInstancesResponseBodyResultValues {
	s.Attributes = v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetFormInstDataList(v []*QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) *QueryAllFormInstancesResponseBodyResultValues {
	s.FormInstDataList = v
	return s
}

type QueryAllFormInstancesResponseBodyResultValuesFormInstDataList struct {
	// 控件类型
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// 控件别名
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// 表单控件扩展数据
	ExtendValue *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	// 控件名称
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// 控件填写的数据
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 控件唯一id
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) GoString() string {
	return s.String()
}

func (s *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) SetComponentType(v string) *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList {
	s.ComponentType = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) SetBizAlias(v string) *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList {
	s.BizAlias = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) SetExtendValue(v string) *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList {
	s.ExtendValue = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) SetLabel(v string) *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList {
	s.Label = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) SetValue(v string) *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList {
	s.Value = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) SetKey(v string) *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList {
	s.Key = &v
	return s
}

type QueryAllFormInstancesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAllFormInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllFormInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllFormInstancesResponse) GoString() string {
	return s.String()
}

func (s *QueryAllFormInstancesResponse) SetHeaders(v map[string]*string) *QueryAllFormInstancesResponse {
	s.Headers = v
	return s
}

func (s *QueryAllFormInstancesResponse) SetBody(v *QueryAllFormInstancesResponseBody) *QueryAllFormInstancesResponse {
	s.Body = v
	return s
}

type QueryFormByBizTypeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryFormByBizTypeHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryFormByBizTypeHeaders) GoString() string {
	return s.String()
}

func (s *QueryFormByBizTypeHeaders) SetCommonHeaders(v map[string]*string) *QueryFormByBizTypeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryFormByBizTypeHeaders) SetXAcsDingtalkAccessToken(v string) *QueryFormByBizTypeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryFormByBizTypeRequest struct {
	// 应用搭建id
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// 表单业务标识
	BizTypes []*string `json:"bizTypes,omitempty" xml:"bizTypes,omitempty" type:"Repeated"`
}

func (s QueryFormByBizTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFormByBizTypeRequest) GoString() string {
	return s.String()
}

func (s *QueryFormByBizTypeRequest) SetAppUuid(v string) *QueryFormByBizTypeRequest {
	s.AppUuid = &v
	return s
}

func (s *QueryFormByBizTypeRequest) SetBizTypes(v []*string) *QueryFormByBizTypeRequest {
	s.BizTypes = v
	return s
}

type QueryFormByBizTypeResponseBody struct {
	// 模板列表
	Result []*QueryFormByBizTypeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryFormByBizTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFormByBizTypeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFormByBizTypeResponseBody) SetResult(v []*QueryFormByBizTypeResponseBodyResult) *QueryFormByBizTypeResponseBody {
	s.Result = v
	return s
}

type QueryFormByBizTypeResponseBodyResult struct {
	// 创建人
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 应用搭建id
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// 模板code
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// 表单uuid
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 模板名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 模板描述
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 数据归属id
	OwnerId *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// 表单类型，0为流程表单，1为数据表单
	AppType *int32 `json:"appType,omitempty" xml:"appType,omitempty"`
	// 业务标识
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 模板状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 创建时间
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 修改时间
	ModifedTime *int64 `json:"modifedTime,omitempty" xml:"modifedTime,omitempty"`
	// 表单控件描述
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s QueryFormByBizTypeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryFormByBizTypeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryFormByBizTypeResponseBodyResult) SetCreator(v string) *QueryFormByBizTypeResponseBodyResult {
	s.Creator = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetAppUuid(v string) *QueryFormByBizTypeResponseBodyResult {
	s.AppUuid = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetFormCode(v string) *QueryFormByBizTypeResponseBodyResult {
	s.FormCode = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetFormUuid(v string) *QueryFormByBizTypeResponseBodyResult {
	s.FormUuid = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetName(v string) *QueryFormByBizTypeResponseBodyResult {
	s.Name = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetMemo(v string) *QueryFormByBizTypeResponseBodyResult {
	s.Memo = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetOwnerId(v string) *QueryFormByBizTypeResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetAppType(v int32) *QueryFormByBizTypeResponseBodyResult {
	s.AppType = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetBizType(v string) *QueryFormByBizTypeResponseBodyResult {
	s.BizType = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetStatus(v string) *QueryFormByBizTypeResponseBodyResult {
	s.Status = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetCreateTime(v int64) *QueryFormByBizTypeResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetModifedTime(v int64) *QueryFormByBizTypeResponseBodyResult {
	s.ModifedTime = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetContent(v string) *QueryFormByBizTypeResponseBodyResult {
	s.Content = &v
	return s
}

type QueryFormByBizTypeResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryFormByBizTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryFormByBizTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFormByBizTypeResponse) GoString() string {
	return s.String()
}

func (s *QueryFormByBizTypeResponse) SetHeaders(v map[string]*string) *QueryFormByBizTypeResponse {
	s.Headers = v
	return s
}

func (s *QueryFormByBizTypeResponse) SetBody(v *QueryFormByBizTypeResponseBody) *QueryFormByBizTypeResponse {
	s.Body = v
	return s
}

type FormCreateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s FormCreateHeaders) String() string {
	return tea.Prettify(s)
}

func (s FormCreateHeaders) GoString() string {
	return s.String()
}

func (s *FormCreateHeaders) SetCommonHeaders(v map[string]*string) *FormCreateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FormCreateHeaders) SetXAcsDingtalkAccessToken(v string) *FormCreateHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type FormCreateRequest struct {
	DingCorpId         *string `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	DingOrgId          *int64  `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingIsvOrgId       *int64  `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingSuiteKey       *string `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType *int64  `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ProcessCode        *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// 表单模板名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 表单模板描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 表单控件列表
	FormComponents []*FormComponent `json:"formComponents,omitempty" xml:"formComponents,omitempty" type:"Repeated"`
	// 模板配置信息
	TemplateConfig *FormCreateRequestTemplateConfig `json:"templateConfig,omitempty" xml:"templateConfig,omitempty" type:"Struct"`
}

func (s FormCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequest) GoString() string {
	return s.String()
}

func (s *FormCreateRequest) SetDingCorpId(v string) *FormCreateRequest {
	s.DingCorpId = &v
	return s
}

func (s *FormCreateRequest) SetDingOrgId(v int64) *FormCreateRequest {
	s.DingOrgId = &v
	return s
}

func (s *FormCreateRequest) SetDingIsvOrgId(v int64) *FormCreateRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *FormCreateRequest) SetDingSuiteKey(v string) *FormCreateRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *FormCreateRequest) SetDingTokenGrantType(v int64) *FormCreateRequest {
	s.DingTokenGrantType = &v
	return s
}

func (s *FormCreateRequest) SetRequestId(v string) *FormCreateRequest {
	s.RequestId = &v
	return s
}

func (s *FormCreateRequest) SetProcessCode(v string) *FormCreateRequest {
	s.ProcessCode = &v
	return s
}

func (s *FormCreateRequest) SetName(v string) *FormCreateRequest {
	s.Name = &v
	return s
}

func (s *FormCreateRequest) SetDescription(v string) *FormCreateRequest {
	s.Description = &v
	return s
}

func (s *FormCreateRequest) SetFormComponents(v []*FormComponent) *FormCreateRequest {
	s.FormComponents = v
	return s
}

func (s *FormCreateRequest) SetTemplateConfig(v *FormCreateRequestTemplateConfig) *FormCreateRequest {
	s.TemplateConfig = v
	return s
}

type FormCreateRequestTemplateConfig struct {
	// 禁用停止按钮
	DisableStopProcessButton *bool `json:"disableStopProcessButton,omitempty" xml:"disableStopProcessButton,omitempty"`
	// 审批场景内隐藏模板
	Hidden *bool `json:"hidden,omitempty" xml:"hidden,omitempty"`
	// 禁用模板删除按钮
	DisableDeleteProcess *bool `json:"disableDeleteProcess,omitempty" xml:"disableDeleteProcess,omitempty"`
	// 禁用表单编辑
	DisableFormEdit *bool `json:"disableFormEdit,omitempty" xml:"disableFormEdit,omitempty"`
	// 禁用再次提交
	DisableResubmit *bool `json:"disableResubmit,omitempty" xml:"disableResubmit,omitempty"`
	// 首页工作台是否可见
	DisableHomepage *bool `json:"disableHomepage,omitempty" xml:"disableHomepage,omitempty"`
	// 更新后模板目录id
	DirId *string `json:"dirId,omitempty" xml:"dirId,omitempty"`
	// 源模板目录id
	OriginDirId *string `json:"originDirId,omitempty" xml:"originDirId,omitempty"`
}

func (s FormCreateRequestTemplateConfig) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestTemplateConfig) GoString() string {
	return s.String()
}

func (s *FormCreateRequestTemplateConfig) SetDisableStopProcessButton(v bool) *FormCreateRequestTemplateConfig {
	s.DisableStopProcessButton = &v
	return s
}

func (s *FormCreateRequestTemplateConfig) SetHidden(v bool) *FormCreateRequestTemplateConfig {
	s.Hidden = &v
	return s
}

func (s *FormCreateRequestTemplateConfig) SetDisableDeleteProcess(v bool) *FormCreateRequestTemplateConfig {
	s.DisableDeleteProcess = &v
	return s
}

func (s *FormCreateRequestTemplateConfig) SetDisableFormEdit(v bool) *FormCreateRequestTemplateConfig {
	s.DisableFormEdit = &v
	return s
}

func (s *FormCreateRequestTemplateConfig) SetDisableResubmit(v bool) *FormCreateRequestTemplateConfig {
	s.DisableResubmit = &v
	return s
}

func (s *FormCreateRequestTemplateConfig) SetDisableHomepage(v bool) *FormCreateRequestTemplateConfig {
	s.DisableHomepage = &v
	return s
}

func (s *FormCreateRequestTemplateConfig) SetDirId(v string) *FormCreateRequestTemplateConfig {
	s.DirId = &v
	return s
}

func (s *FormCreateRequestTemplateConfig) SetOriginDirId(v string) *FormCreateRequestTemplateConfig {
	s.OriginDirId = &v
	return s
}

type FormCreateResponseBody struct {
	// 表单模板信息
	Result *FormCreateResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s FormCreateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FormCreateResponseBody) GoString() string {
	return s.String()
}

func (s *FormCreateResponseBody) SetResult(v *FormCreateResponseBodyResult) *FormCreateResponseBody {
	s.Result = v
	return s
}

type FormCreateResponseBodyResult struct {
	// 保存或更新的表单code
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
}

func (s FormCreateResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s FormCreateResponseBodyResult) GoString() string {
	return s.String()
}

func (s *FormCreateResponseBodyResult) SetProcessCode(v string) *FormCreateResponseBodyResult {
	s.ProcessCode = &v
	return s
}

type FormCreateResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FormCreateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FormCreateResponse) String() string {
	return tea.Prettify(s)
}

func (s FormCreateResponse) GoString() string {
	return s.String()
}

func (s *FormCreateResponse) SetHeaders(v map[string]*string) *FormCreateResponse {
	s.Headers = v
	return s
}

func (s *FormCreateResponse) SetBody(v *FormCreateResponseBody) *FormCreateResponse {
	s.Body = v
	return s
}

type QuerySchemaByProcessCodeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QuerySchemaByProcessCodeHeaders) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaByProcessCodeHeaders) GoString() string {
	return s.String()
}

func (s *QuerySchemaByProcessCodeHeaders) SetCommonHeaders(v map[string]*string) *QuerySchemaByProcessCodeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QuerySchemaByProcessCodeHeaders) SetXAcsDingtalkAccessToken(v string) *QuerySchemaByProcessCodeHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QuerySchemaByProcessCodeRequest struct {
	// 表单的唯一码
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
}

func (s QuerySchemaByProcessCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaByProcessCodeRequest) GoString() string {
	return s.String()
}

func (s *QuerySchemaByProcessCodeRequest) SetProcessCode(v string) *QuerySchemaByProcessCodeRequest {
	s.ProcessCode = &v
	return s
}

type QuerySchemaByProcessCodeResponseBody struct {
	// 返回结果详情。
	Result *QuerySchemaByProcessCodeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QuerySchemaByProcessCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaByProcessCodeResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySchemaByProcessCodeResponseBody) SetResult(v *QuerySchemaByProcessCodeResponseBodyResult) *QuerySchemaByProcessCodeResponseBody {
	s.Result = v
	return s
}

type QuerySchemaByProcessCodeResponseBodyResult struct {
	// 创建人 userId。
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// 创建人 uid。
	CreatorUid *int64 `json:"creatorUid,omitempty" xml:"creatorUid,omitempty"`
	// 表单应用 uuid 或者 corpId。
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// 表单的唯一码。
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// 表单 uuid。
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// 表单名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 说明文案。
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 数据归属者的 id。
	OwnerId *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// 数据归属者的 id 类型。企业(orgId), 群(cid), 人(uid)。
	OwnerIdType *string `json:"ownerIdType,omitempty" xml:"ownerIdType,omitempty"`
	// 表单 schema 详情。
	SchemaContent *QuerySchemaByProcessCodeResponseBodyResultSchemaContent `json:"schemaContent,omitempty" xml:"schemaContent,omitempty" type:"Struct"`
	// 图标。
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 表单类型。
	AppType *int32 `json:"appType,omitempty" xml:"appType,omitempty"`
	// 代表表单业务含义的类型。
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 引擎类型，表单：0，页面：1
	EngineType *int32 `json:"engineType,omitempty" xml:"engineType,omitempty"`
	// 状态, PUBLISHED(启用), INVALID(停用), SAVED(草稿)
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 排序 id。
	ListOrder *int32 `json:"listOrder,omitempty" xml:"listOrder,omitempty"`
	// 业务自定义设置数据。
	CustomSetting *string `json:"customSetting,omitempty" xml:"customSetting,omitempty"`
	// 目标类型: inner, outer, customer。
	ProcType *string `json:"procType,omitempty" xml:"procType,omitempty"`
	// 可见范围类型。
	VisibleRange *string `json:"visibleRange,omitempty" xml:"visibleRange,omitempty"`
	// 创建时间的时间戳。
	GmtCreate *int32 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 修改时间的时间戳。
	GmtModified *int32 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
}

func (s QuerySchemaByProcessCodeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaByProcessCodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetCreatorUserId(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.CreatorUserId = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetCreatorUid(v int64) *QuerySchemaByProcessCodeResponseBodyResult {
	s.CreatorUid = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetAppUuid(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.AppUuid = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetFormCode(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.FormCode = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetFormUuid(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.FormUuid = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetName(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.Name = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetMemo(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.Memo = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetOwnerId(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetOwnerIdType(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.OwnerIdType = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetSchemaContent(v *QuerySchemaByProcessCodeResponseBodyResultSchemaContent) *QuerySchemaByProcessCodeResponseBodyResult {
	s.SchemaContent = v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetIcon(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.Icon = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetAppType(v int32) *QuerySchemaByProcessCodeResponseBodyResult {
	s.AppType = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetBizType(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.BizType = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetEngineType(v int32) *QuerySchemaByProcessCodeResponseBodyResult {
	s.EngineType = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetStatus(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.Status = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetListOrder(v int32) *QuerySchemaByProcessCodeResponseBodyResult {
	s.ListOrder = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetCustomSetting(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.CustomSetting = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetProcType(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.ProcType = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetVisibleRange(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.VisibleRange = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetGmtCreate(v int32) *QuerySchemaByProcessCodeResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetGmtModified(v int32) *QuerySchemaByProcessCodeResponseBodyResult {
	s.GmtModified = &v
	return s
}

type QuerySchemaByProcessCodeResponseBodyResultSchemaContent struct {
	// 表单名称。
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 图标
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 控件列表。
	Items []*QuerySchemaByProcessCodeResponseBodyResultSchemaContentItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContent) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContent) GoString() string {
	return s.String()
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContent) SetTitle(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContent {
	s.Title = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContent) SetIcon(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContent {
	s.Icon = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContent) SetItems(v []*QuerySchemaByProcessCodeResponseBodyResultSchemaContentItems) *QuerySchemaByProcessCodeResponseBodyResultSchemaContent {
	s.Items = v
	return s
}

type QuerySchemaByProcessCodeResponseBodyResultSchemaContentItems struct {
	// 控件类型，取值：
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// 控件属性。
	Props *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItems) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItems) GoString() string {
	return s.String()
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItems) SetComponentName(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItems {
	s.ComponentName = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItems) SetProps(v *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItems {
	s.Props = v
	return s
}

type QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps struct {
	// 控件 id。
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 控件名称。
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// 控件业务自定义别名。
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// 是否必填。
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// 占位符。
	Placeholder *string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// 单选框选项列表。
	Options []*string `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// ISV 微应用 appId，用于ISV身份权限识别，ISV可获得相应数据。
	AppId *int64 `json:"appId,omitempty" xml:"appId,omitempty"`
	// 兼容字段。
	DurationLabel *string `json:"durationLabel,omitempty" xml:"durationLabel,omitempty"`
	// 是否推送管理日历(DDDateRangeField, 1表示推送, 0表示不推送, 该属性为兼容保留)。
	PushToCalendar *int32 `json:"pushToCalendar,omitempty" xml:"pushToCalendar,omitempty"`
	// textnote的样式，top|middle|bottom。
	Align *string `json:"align,omitempty" xml:"align,omitempty"`
	// 需要计算总和的明细组件
	StatField []*QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsStatField `json:"statField,omitempty" xml:"statField,omitempty" type:"Repeated"`
	// 加班套件4.0新增 加班明细是否隐藏标签。
	HideLabel *bool `json:"hideLabel,omitempty" xml:"hideLabel,omitempty"`
	// 选项内容列表，提供给业务方更多的选择器操作。
	ObjOptions []*QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsObjOptions `json:"objOptions,omitempty" xml:"objOptions,omitempty" type:"Repeated"`
	// 时间格式(DDDateField和DDDateRangeField)。
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// 推送到考勤, 子类型(DDSelectField)。
	PushToAttendance *bool `json:"pushToAttendance,omitempty" xml:"pushToAttendance,omitempty"`
	// label是否可修改 true：不可修改。
	LabelEditableFreeze *bool `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	// 同步到考勤, 表示是否设置为员工状态。
	Push *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsPush `json:"push,omitempty" xml:"push,omitempty" type:"Struct"`
	// common field的commonBizType。
	CommonBizType *string `json:"commonBizType,omitempty" xml:"commonBizType,omitempty"`
	// 必填是否可修改 true：不可修改。
	RequiredEditableFreeze *bool `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	// 数字组件/日期区间组件单位属性。
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 套件值是否打平
	Extract *bool `json:"extract,omitempty" xml:"extract,omitempty"`
	// 说明文案的链接地址。
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// 是否有支付属性。
	PayEnable *bool `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	// 加班套件4.0新增 加班明细是否隐藏。
	Hidden *bool `json:"hidden,omitempty" xml:"hidden,omitempty"`
	// 业务套件类型。
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// 是否开启员工状态。
	StaffStatusEnabled *bool `json:"staffStatusEnabled,omitempty" xml:"staffStatusEnabled,omitempty"`
	// 加班套件4.0新增 加班明细名称。
	ActionName *string `json:"actionName,omitempty" xml:"actionName,omitempty"`
	// 请假、出差、外出、加班类型标签。
	AttendTypeLabel *string `json:"attendTypeLabel,omitempty" xml:"attendTypeLabel,omitempty"`
	// 套件内子组件可见性。
	ChildFieldVisible *bool `json:"childFieldVisible,omitempty" xml:"childFieldVisible,omitempty"`
	// 是否参与打印(1表示不打印, 0表示打印)。
	NotPrint *string `json:"notPrint,omitempty" xml:"notPrint,omitempty"`
	// 明细打印排版方式 false：横向 true：纵向。
	VerticalPrint *bool `json:"verticalPrint,omitempty" xml:"verticalPrint,omitempty"`
	// 是否自动计算时长。
	Duration *bool `json:"duration,omitempty" xml:"duration,omitempty"`
	// 兼容出勤套件类型。
	HolidayOptions *string `json:"holidayOptions,omitempty" xml:"holidayOptions,omitempty"`
	// 是否使用考勤日历。
	UseCalendar *bool `json:"useCalendar,omitempty" xml:"useCalendar,omitempty"`
	// textnote在详情页是否隐藏，true隐藏， false不隐藏
	HiddenInApprovalDetail *bool `json:"hiddenInApprovalDetail,omitempty" xml:"hiddenInApprovalDetail,omitempty"`
	// 是否可编辑。
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// 套件是否开启异步获取分条件规则，true：开启；false：不开启。
	AsyncCondition *bool `json:"asyncCondition,omitempty" xml:"asyncCondition,omitempty"`
	// 表单关联控件列表。
	BehaviorLinkage []*QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkage `json:"behaviorLinkage,omitempty" xml:"behaviorLinkage,omitempty" type:"Repeated"`
	// 兼容出勤套件类型。
	ShowAttendOptions *bool `json:"showAttendOptions,omitempty" xml:"showAttendOptions,omitempty"`
	// 是否需要大写 默认是需要; 1:不需要大写, 空或者0:需要大写。
	NotUpper *string `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	// 关联表单中的fields存储
	FieldsInfo *string `json:"fieldsInfo,omitempty" xml:"fieldsInfo,omitempty"`
	// e签宝专用标识。
	ESign *bool `json:"eSign,omitempty" xml:"eSign,omitempty"`
	// 加班套件4.0新增 加班明细描述。
	MainTitle *string `json:"mainTitle,omitempty" xml:"mainTitle,omitempty"`
	// 公式。
	Formula *string `json:"formula,omitempty" xml:"formula,omitempty"`
	// 内部联系人choice，1表示多选，0表示单选。
	Choice *int32 `json:"choice,omitempty" xml:"choice,omitempty"`
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) GoString() string {
	return s.String()
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetId(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Id = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetLabel(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Label = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetBizAlias(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.BizAlias = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetRequired(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Required = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetPlaceholder(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Placeholder = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetOptions(v []*string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Options = v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetAppId(v int64) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.AppId = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetDurationLabel(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.DurationLabel = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetPushToCalendar(v int32) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.PushToCalendar = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetAlign(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Align = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetStatField(v []*QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsStatField) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.StatField = v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetHideLabel(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.HideLabel = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetObjOptions(v []*QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsObjOptions) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.ObjOptions = v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetFormat(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Format = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetPushToAttendance(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.PushToAttendance = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetLabelEditableFreeze(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.LabelEditableFreeze = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetPush(v *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsPush) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Push = v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetCommonBizType(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.CommonBizType = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetRequiredEditableFreeze(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.RequiredEditableFreeze = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetUnit(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Unit = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetExtract(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Extract = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetLink(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Link = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetPayEnable(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.PayEnable = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetHidden(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Hidden = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetBizType(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.BizType = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetStaffStatusEnabled(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.StaffStatusEnabled = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetActionName(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.ActionName = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetAttendTypeLabel(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.AttendTypeLabel = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetChildFieldVisible(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.ChildFieldVisible = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetNotPrint(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.NotPrint = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetVerticalPrint(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.VerticalPrint = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetDuration(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Duration = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetHolidayOptions(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.HolidayOptions = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetUseCalendar(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.UseCalendar = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetHiddenInApprovalDetail(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.HiddenInApprovalDetail = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetDisabled(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Disabled = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetAsyncCondition(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.AsyncCondition = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetBehaviorLinkage(v []*QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkage) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.BehaviorLinkage = v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetShowAttendOptions(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.ShowAttendOptions = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetNotUpper(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.NotUpper = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetFieldsInfo(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.FieldsInfo = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetESign(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.ESign = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetMainTitle(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.MainTitle = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetFormula(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Formula = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetChoice(v int32) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Choice = &v
	return s
}

type QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsStatField struct {
	// id 值。
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 名称。
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// 大写。
	Upper *bool `json:"upper,omitempty" xml:"upper,omitempty"`
	// 单位。
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsStatField) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsStatField) GoString() string {
	return s.String()
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsStatField) SetId(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsStatField {
	s.Id = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsStatField) SetLabel(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsStatField {
	s.Label = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsStatField) SetUpper(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsStatField {
	s.Upper = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsStatField) SetUnit(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsStatField {
	s.Unit = &v
	return s
}

type QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsObjOptions struct {
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsObjOptions) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsObjOptions) GoString() string {
	return s.String()
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsObjOptions) SetValue(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsObjOptions {
	s.Value = &v
	return s
}

type QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsPush struct {
	// 开启状态(1表示开启, 0表示关闭)
	PushSwitch *int32 `json:"pushSwitch,omitempty" xml:"pushSwitch,omitempty"`
	// 状态显示名称
	PushTag *string `json:"pushTag,omitempty" xml:"pushTag,omitempty"`
	// 考勤类型(1表示请假, 2表示出差, 3表示加班, 4表示外出)
	AttendanceRule *int32 `json:"attendanceRule,omitempty" xml:"attendanceRule,omitempty"`
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsPush) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsPush) GoString() string {
	return s.String()
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsPush) SetPushSwitch(v int32) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsPush {
	s.PushSwitch = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsPush) SetPushTag(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsPush {
	s.PushTag = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsPush) SetAttendanceRule(v int32) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsPush {
	s.AttendanceRule = &v
	return s
}

type QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkage struct {
	// 控件值。
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 关联控件列表。
	Targets []*QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets `json:"targets,omitempty" xml:"targets,omitempty" type:"Repeated"`
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkage) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkage) GoString() string {
	return s.String()
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkage) SetValue(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkage {
	s.Value = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkage) SetTargets(v []*QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkage {
	s.Targets = v
	return s
}

type QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets struct {
	// 字段 id。
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// 行为。
	Behavior *string `json:"behavior,omitempty" xml:"behavior,omitempty"`
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets) GoString() string {
	return s.String()
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets) SetFieldId(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets {
	s.FieldId = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets) SetBehavior(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets {
	s.Behavior = &v
	return s
}

type QuerySchemaByProcessCodeResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySchemaByProcessCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySchemaByProcessCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaByProcessCodeResponse) GoString() string {
	return s.String()
}

func (s *QuerySchemaByProcessCodeResponse) SetHeaders(v map[string]*string) *QuerySchemaByProcessCodeResponse {
	s.Headers = v
	return s
}

func (s *QuerySchemaByProcessCodeResponse) SetBody(v *QuerySchemaByProcessCodeResponseBody) *QuerySchemaByProcessCodeResponse {
	s.Body = v
	return s
}

type StartProcessInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s StartProcessInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s StartProcessInstanceHeaders) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceHeaders) SetCommonHeaders(v map[string]*string) *StartProcessInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartProcessInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *StartProcessInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type StartProcessInstanceRequest struct {
	// 审批发起人的userId
	OriginatorUserId *string `json:"originatorUserId,omitempty" xml:"originatorUserId,omitempty"`
	// 审批流的唯一码
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// 部门ID
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// 企业微应用标识
	MicroappAgentId *int64 `json:"microappAgentId,omitempty" xml:"microappAgentId,omitempty"`
	// 不使用审批流模板时，直接指定审批人列表
	Approvers []*StartProcessInstanceRequestApprovers `json:"approvers,omitempty" xml:"approvers,omitempty" type:"Repeated"`
	// 抄送人userId列表
	CcList []*string `json:"ccList,omitempty" xml:"ccList,omitempty" type:"Repeated"`
	// 抄送时间
	CcPosition *string `json:"ccPosition,omitempty" xml:"ccPosition,omitempty"`
	// 使用审批流模板时，模板上的自选操作人列表
	TargetSelectActioners []*StartProcessInstanceRequestTargetSelectActioners `json:"targetSelectActioners,omitempty" xml:"targetSelectActioners,omitempty" type:"Repeated"`
	// 表单数据内容，控件列表
	FormComponentValues []*StartProcessInstanceRequestFormComponentValues `json:"formComponentValues,omitempty" xml:"formComponentValues,omitempty" type:"Repeated"`
	RequestId           *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DingCorpId          *string                                           `json:"dingCorpId,omitempty" xml:"dingCorpId,omitempty"`
	DingOrgId           *int64                                            `json:"dingOrgId,omitempty" xml:"dingOrgId,omitempty"`
	DingIsvOrgId        *int64                                            `json:"dingIsvOrgId,omitempty" xml:"dingIsvOrgId,omitempty"`
	DingSuiteKey        *string                                           `json:"dingSuiteKey,omitempty" xml:"dingSuiteKey,omitempty"`
	DingTokenGrantType  *int64                                            `json:"dingTokenGrantType,omitempty" xml:"dingTokenGrantType,omitempty"`
}

func (s StartProcessInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartProcessInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceRequest) SetOriginatorUserId(v string) *StartProcessInstanceRequest {
	s.OriginatorUserId = &v
	return s
}

func (s *StartProcessInstanceRequest) SetProcessCode(v string) *StartProcessInstanceRequest {
	s.ProcessCode = &v
	return s
}

func (s *StartProcessInstanceRequest) SetDeptId(v int64) *StartProcessInstanceRequest {
	s.DeptId = &v
	return s
}

func (s *StartProcessInstanceRequest) SetMicroappAgentId(v int64) *StartProcessInstanceRequest {
	s.MicroappAgentId = &v
	return s
}

func (s *StartProcessInstanceRequest) SetApprovers(v []*StartProcessInstanceRequestApprovers) *StartProcessInstanceRequest {
	s.Approvers = v
	return s
}

func (s *StartProcessInstanceRequest) SetCcList(v []*string) *StartProcessInstanceRequest {
	s.CcList = v
	return s
}

func (s *StartProcessInstanceRequest) SetCcPosition(v string) *StartProcessInstanceRequest {
	s.CcPosition = &v
	return s
}

func (s *StartProcessInstanceRequest) SetTargetSelectActioners(v []*StartProcessInstanceRequestTargetSelectActioners) *StartProcessInstanceRequest {
	s.TargetSelectActioners = v
	return s
}

func (s *StartProcessInstanceRequest) SetFormComponentValues(v []*StartProcessInstanceRequestFormComponentValues) *StartProcessInstanceRequest {
	s.FormComponentValues = v
	return s
}

func (s *StartProcessInstanceRequest) SetRequestId(v string) *StartProcessInstanceRequest {
	s.RequestId = &v
	return s
}

func (s *StartProcessInstanceRequest) SetDingCorpId(v string) *StartProcessInstanceRequest {
	s.DingCorpId = &v
	return s
}

func (s *StartProcessInstanceRequest) SetDingOrgId(v int64) *StartProcessInstanceRequest {
	s.DingOrgId = &v
	return s
}

func (s *StartProcessInstanceRequest) SetDingIsvOrgId(v int64) *StartProcessInstanceRequest {
	s.DingIsvOrgId = &v
	return s
}

func (s *StartProcessInstanceRequest) SetDingSuiteKey(v string) *StartProcessInstanceRequest {
	s.DingSuiteKey = &v
	return s
}

func (s *StartProcessInstanceRequest) SetDingTokenGrantType(v int64) *StartProcessInstanceRequest {
	s.DingTokenGrantType = &v
	return s
}

type StartProcessInstanceRequestApprovers struct {
	// 审批类型
	ActionType *string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// 审批人列表
	UserIds []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s StartProcessInstanceRequestApprovers) String() string {
	return tea.Prettify(s)
}

func (s StartProcessInstanceRequestApprovers) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceRequestApprovers) SetActionType(v string) *StartProcessInstanceRequestApprovers {
	s.ActionType = &v
	return s
}

func (s *StartProcessInstanceRequestApprovers) SetUserIds(v []*string) *StartProcessInstanceRequestApprovers {
	s.UserIds = v
	return s
}

type StartProcessInstanceRequestTargetSelectActioners struct {
	// 自选节点的规则key
	ActionerKey *string `json:"actionerKey,omitempty" xml:"actionerKey,omitempty"`
	// 操作人userId列表
	ActionerUserIds []*string `json:"actionerUserIds,omitempty" xml:"actionerUserIds,omitempty" type:"Repeated"`
}

func (s StartProcessInstanceRequestTargetSelectActioners) String() string {
	return tea.Prettify(s)
}

func (s StartProcessInstanceRequestTargetSelectActioners) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceRequestTargetSelectActioners) SetActionerKey(v string) *StartProcessInstanceRequestTargetSelectActioners {
	s.ActionerKey = &v
	return s
}

func (s *StartProcessInstanceRequestTargetSelectActioners) SetActionerUserIds(v []*string) *StartProcessInstanceRequestTargetSelectActioners {
	s.ActionerUserIds = v
	return s
}

type StartProcessInstanceRequestFormComponentValues struct {
	// 控件id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 控件别名
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// 控件名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 控件值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 控件扩展值
	ExtValue *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	// 控件类型
	ComponentType *string                                                  `json:"componentType,omitempty" xml:"componentType,omitempty"`
	Details       []*StartProcessInstanceRequestFormComponentValuesDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
}

func (s StartProcessInstanceRequestFormComponentValues) String() string {
	return tea.Prettify(s)
}

func (s StartProcessInstanceRequestFormComponentValues) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceRequestFormComponentValues) SetId(v string) *StartProcessInstanceRequestFormComponentValues {
	s.Id = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValues) SetBizAlias(v string) *StartProcessInstanceRequestFormComponentValues {
	s.BizAlias = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValues) SetName(v string) *StartProcessInstanceRequestFormComponentValues {
	s.Name = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValues) SetValue(v string) *StartProcessInstanceRequestFormComponentValues {
	s.Value = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValues) SetExtValue(v string) *StartProcessInstanceRequestFormComponentValues {
	s.ExtValue = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValues) SetComponentType(v string) *StartProcessInstanceRequestFormComponentValues {
	s.ComponentType = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValues) SetDetails(v []*StartProcessInstanceRequestFormComponentValuesDetails) *StartProcessInstanceRequestFormComponentValues {
	s.Details = v
	return s
}

type StartProcessInstanceRequestFormComponentValuesDetails struct {
	// 控件id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 控件别名
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// 控件名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 控件值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 控件扩展值
	ExtValue *string                                                         `json:"extValue,omitempty" xml:"extValue,omitempty"`
	Details  []*StartProcessInstanceRequestFormComponentValuesDetailsDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
}

func (s StartProcessInstanceRequestFormComponentValuesDetails) String() string {
	return tea.Prettify(s)
}

func (s StartProcessInstanceRequestFormComponentValuesDetails) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceRequestFormComponentValuesDetails) SetId(v string) *StartProcessInstanceRequestFormComponentValuesDetails {
	s.Id = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValuesDetails) SetBizAlias(v string) *StartProcessInstanceRequestFormComponentValuesDetails {
	s.BizAlias = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValuesDetails) SetName(v string) *StartProcessInstanceRequestFormComponentValuesDetails {
	s.Name = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValuesDetails) SetValue(v string) *StartProcessInstanceRequestFormComponentValuesDetails {
	s.Value = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValuesDetails) SetExtValue(v string) *StartProcessInstanceRequestFormComponentValuesDetails {
	s.ExtValue = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValuesDetails) SetDetails(v []*StartProcessInstanceRequestFormComponentValuesDetailsDetails) *StartProcessInstanceRequestFormComponentValuesDetails {
	s.Details = v
	return s
}

type StartProcessInstanceRequestFormComponentValuesDetailsDetails struct {
	// 控件id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// 控件别名
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// 控件名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 控件值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// 控件扩展值
	ExtValue *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	// 控件类型
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
}

func (s StartProcessInstanceRequestFormComponentValuesDetailsDetails) String() string {
	return tea.Prettify(s)
}

func (s StartProcessInstanceRequestFormComponentValuesDetailsDetails) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceRequestFormComponentValuesDetailsDetails) SetId(v string) *StartProcessInstanceRequestFormComponentValuesDetailsDetails {
	s.Id = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValuesDetailsDetails) SetBizAlias(v string) *StartProcessInstanceRequestFormComponentValuesDetailsDetails {
	s.BizAlias = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValuesDetailsDetails) SetName(v string) *StartProcessInstanceRequestFormComponentValuesDetailsDetails {
	s.Name = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValuesDetailsDetails) SetValue(v string) *StartProcessInstanceRequestFormComponentValuesDetailsDetails {
	s.Value = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValuesDetailsDetails) SetExtValue(v string) *StartProcessInstanceRequestFormComponentValuesDetailsDetails {
	s.ExtValue = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValuesDetailsDetails) SetComponentType(v string) *StartProcessInstanceRequestFormComponentValuesDetailsDetails {
	s.ComponentType = &v
	return s
}

type StartProcessInstanceResponseBody struct {
	// 审批实例id
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s StartProcessInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartProcessInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceResponseBody) SetInstanceId(v string) *StartProcessInstanceResponseBody {
	s.InstanceId = &v
	return s
}

type StartProcessInstanceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartProcessInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartProcessInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartProcessInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceResponse) SetHeaders(v map[string]*string) *StartProcessInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartProcessInstanceResponse) SetBody(v *StartProcessInstanceResponseBody) *StartProcessInstanceResponse {
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

func (client *Client) QueryFormInstance(request *QueryFormInstanceRequest) (_result *QueryFormInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryFormInstanceHeaders{}
	_result = &QueryFormInstanceResponse{}
	_body, _err := client.QueryFormInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFormInstanceWithOptions(request *QueryFormInstanceRequest, headers *QueryFormInstanceHeaders, runtime *util.RuntimeOptions) (_result *QueryFormInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FormInstanceId)) {
		query["formInstanceId"] = request.FormInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.FormCode)) {
		query["formCode"] = request.FormCode
	}

	if !tea.BoolValue(util.IsUnset(request.AppUuid)) {
		query["appUuid"] = request.AppUuid
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
	_result = &QueryFormInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryFormInstance"), tea.String("workflow_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/workflow/forms/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ProcessForecast(request *ProcessForecastRequest) (_result *ProcessForecastResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ProcessForecastHeaders{}
	_result = &ProcessForecastResponse{}
	_body, _err := client.ProcessForecastWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ProcessForecastWithOptions(request *ProcessForecastRequest, headers *ProcessForecastHeaders, runtime *util.RuntimeOptions) (_result *ProcessForecastResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingCorpId)) {
		body["dingCorpId"] = request.DingCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		body["processCode"] = request.ProcessCode
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.FormComponentValues)) {
		body["formComponentValues"] = request.FormComponentValues
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
	_result = &ProcessForecastResponse{}
	_body, _err := client.DoROARequest(tea.String("ProcessForecast"), tea.String("workflow_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/workflow/processes/forecast"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllProcessInstances(request *QueryAllProcessInstancesRequest) (_result *QueryAllProcessInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllProcessInstancesHeaders{}
	_result = &QueryAllProcessInstancesResponse{}
	_body, _err := client.QueryAllProcessInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllProcessInstancesWithOptions(request *QueryAllProcessInstancesRequest, headers *QueryAllProcessInstancesHeaders, runtime *util.RuntimeOptions) (_result *QueryAllProcessInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimeInMills)) {
		query["startTimeInMills"] = request.StartTimeInMills
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimeInMills)) {
		query["endTimeInMills"] = request.EndTimeInMills
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		query["processCode"] = request.ProcessCode
	}

	if !tea.BoolValue(util.IsUnset(request.AppUuid)) {
		query["appUuid"] = request.AppUuid
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
	_result = &QueryAllProcessInstancesResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAllProcessInstances"), tea.String("workflow_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/workflow/processes/pages/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllFormInstances(request *QueryAllFormInstancesRequest) (_result *QueryAllFormInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryAllFormInstancesHeaders{}
	_result = &QueryAllFormInstancesResponse{}
	_body, _err := client.QueryAllFormInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllFormInstancesWithOptions(request *QueryAllFormInstancesRequest, headers *QueryAllFormInstancesHeaders, runtime *util.RuntimeOptions) (_result *QueryAllFormInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.AppUuid)) {
		query["appUuid"] = request.AppUuid
	}

	if !tea.BoolValue(util.IsUnset(request.FormCode)) {
		query["formCode"] = request.FormCode
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
	_result = &QueryAllFormInstancesResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryAllFormInstances"), tea.String("workflow_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/workflow/forms/pages/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFormByBizType(request *QueryFormByBizTypeRequest) (_result *QueryFormByBizTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryFormByBizTypeHeaders{}
	_result = &QueryFormByBizTypeResponse{}
	_body, _err := client.QueryFormByBizTypeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFormByBizTypeWithOptions(request *QueryFormByBizTypeRequest, headers *QueryFormByBizTypeHeaders, runtime *util.RuntimeOptions) (_result *QueryFormByBizTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUuid)) {
		body["appUuid"] = request.AppUuid
	}

	if !tea.BoolValue(util.IsUnset(request.BizTypes)) {
		body["bizTypes"] = request.BizTypes
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
	_result = &QueryFormByBizTypeResponse{}
	_body, _err := client.DoROARequest(tea.String("QueryFormByBizType"), tea.String("workflow_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/workflow/forms/forminfos/query"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FormCreate(request *FormCreateRequest) (_result *FormCreateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &FormCreateHeaders{}
	_result = &FormCreateResponse{}
	_body, _err := client.FormCreateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FormCreateWithOptions(request *FormCreateRequest, headers *FormCreateHeaders, runtime *util.RuntimeOptions) (_result *FormCreateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingCorpId)) {
		body["dingCorpId"] = request.DingCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		body["processCode"] = request.ProcessCode
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FormComponents)) {
		body["formComponents"] = request.FormComponents
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TemplateConfig))) {
		body["templateConfig"] = request.TemplateConfig
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
	_result = &FormCreateResponse{}
	_body, _err := client.DoROARequest(tea.String("FormCreate"), tea.String("workflow_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/workflow/forms"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySchemaByProcessCode(request *QuerySchemaByProcessCodeRequest) (_result *QuerySchemaByProcessCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QuerySchemaByProcessCodeHeaders{}
	_result = &QuerySchemaByProcessCodeResponse{}
	_body, _err := client.QuerySchemaByProcessCodeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySchemaByProcessCodeWithOptions(request *QuerySchemaByProcessCodeRequest, headers *QuerySchemaByProcessCodeHeaders, runtime *util.RuntimeOptions) (_result *QuerySchemaByProcessCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		query["processCode"] = request.ProcessCode
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
	_result = &QuerySchemaByProcessCodeResponse{}
	_body, _err := client.DoROARequest(tea.String("QuerySchemaByProcessCode"), tea.String("workflow_1.0"), tea.String("HTTP"), tea.String("GET"), tea.String("AK"), tea.String("/v1.0/workflow/forms/schemas/processCodes"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartProcessInstance(request *StartProcessInstanceRequest) (_result *StartProcessInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &StartProcessInstanceHeaders{}
	_result = &StartProcessInstanceResponse{}
	_body, _err := client.StartProcessInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartProcessInstanceWithOptions(request *StartProcessInstanceRequest, headers *StartProcessInstanceHeaders, runtime *util.RuntimeOptions) (_result *StartProcessInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OriginatorUserId)) {
		body["originatorUserId"] = request.OriginatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		body["processCode"] = request.ProcessCode
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.MicroappAgentId)) {
		body["microappAgentId"] = request.MicroappAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.Approvers)) {
		body["approvers"] = request.Approvers
	}

	if !tea.BoolValue(util.IsUnset(request.CcList)) {
		body["ccList"] = request.CcList
	}

	if !tea.BoolValue(util.IsUnset(request.CcPosition)) {
		body["ccPosition"] = request.CcPosition
	}

	if !tea.BoolValue(util.IsUnset(request.TargetSelectActioners)) {
		body["targetSelectActioners"] = request.TargetSelectActioners
	}

	if !tea.BoolValue(util.IsUnset(request.FormComponentValues)) {
		body["formComponentValues"] = request.FormComponentValues
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.DingCorpId)) {
		body["dingCorpId"] = request.DingCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.DingOrgId)) {
		body["dingOrgId"] = request.DingOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingIsvOrgId)) {
		body["dingIsvOrgId"] = request.DingIsvOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.DingSuiteKey)) {
		body["dingSuiteKey"] = request.DingSuiteKey
	}

	if !tea.BoolValue(util.IsUnset(request.DingTokenGrantType)) {
		body["dingTokenGrantType"] = request.DingTokenGrantType
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
	_result = &StartProcessInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("StartProcessInstance"), tea.String("workflow_1.0"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/v1.0/workflow/processInstances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
