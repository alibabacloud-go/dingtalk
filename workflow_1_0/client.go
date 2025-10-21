// This file is auto-generated, don't edit it. Thanks.
package workflow_1_0

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	gatewayclient "github.com/alibabacloud-go/gateway-dingtalk/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AvaliableTemplate struct {
	// This parameter is required.
	//
	// example:
	//
	// 出差申请
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-abcd
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

type FormComponent struct {
	Children []*FormComponent `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// TextField
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// This parameter is required.
	Props *FormComponentProps `json:"props,omitempty" xml:"props,omitempty"`
}

func (s FormComponent) String() string {
	return tea.Prettify(s)
}

func (s FormComponent) GoString() string {
	return s.String()
}

func (s *FormComponent) SetChildren(v []*FormComponent) *FormComponent {
	s.Children = v
	return s
}

func (s *FormComponent) SetComponentType(v string) *FormComponent {
	s.ComponentType = &v
	return s
}

func (s *FormComponent) SetProps(v *FormComponentProps) *FormComponent {
	s.Props = v
	return s
}

type FormComponentProps struct {
	// example:
	//
	// 增加明细
	ActionName   *string `json:"actionName,omitempty" xml:"actionName,omitempty"`
	AddressModel *string `json:"addressModel,omitempty" xml:"addressModel,omitempty"`
	// example:
	//
	// top
	Align *string `json:"align,omitempty" xml:"align,omitempty"`
	// example:
	//
	// true
	AsyncCondition     *bool                `json:"asyncCondition,omitempty" xml:"asyncCondition,omitempty"`
	AvailableTemplates []*AvaliableTemplate `json:"availableTemplates,omitempty" xml:"availableTemplates,omitempty" type:"Repeated"`
	// example:
	//
	// finance_name
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// example:
	//
	// attendance.leave
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 0
	Choice *string `json:"choice,omitempty" xml:"choice,omitempty"`
	// example:
	//
	// custom_view
	CommonBizType *string `json:"commonBizType,omitempty" xml:"commonBizType,omitempty"`
	// example:
	//
	// TextField-abcd
	ComponentId *string `json:"componentId,omitempty" xml:"componentId,omitempty"`
	// example:
	//
	// 我是说明文字控件
	Content    *string         `json:"content,omitempty" xml:"content,omitempty"`
	DataSource *FormDataSource `json:"dataSource,omitempty" xml:"dataSource,omitempty"`
	// example:
	//
	// true
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// example:
	//
	// true
	Duration *bool `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 时长
	DurationLabel *string `json:"durationLabel,omitempty" xml:"durationLabel,omitempty"`
	// example:
	//
	// yyyy-MM-dd
	Format  *string `json:"format,omitempty" xml:"format,omitempty"`
	Formula *string `json:"formula,omitempty" xml:"formula,omitempty"`
	// example:
	//
	// true
	Invisible *bool `json:"invisible,omitempty" xml:"invisible,omitempty"`
	// example:
	//
	// 姓名
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// example:
	//
	// 5
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// example:
	//
	// http://www.
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// example:
	//
	// 20
	MaxLength *int32 `json:"maxLength,omitempty" xml:"maxLength,omitempty"`
	// example:
	//
	// phone_tel
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// example:
	//
	// true
	Multiple *bool           `json:"multiple,omitempty" xml:"multiple,omitempty"`
	Options  []*SelectOption `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// example:
	//
	// 请输入
	Placeholder *string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 2
	Precision *int32 `json:"precision,omitempty" xml:"precision,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 1
	Print *string `json:"print,omitempty" xml:"print,omitempty"`
	// example:
	//
	// true
	Required  *bool                          `json:"required,omitempty" xml:"required,omitempty"`
	StatField []*FormComponentPropsStatField `json:"statField,omitempty" xml:"statField,omitempty" type:"Repeated"`
	// example:
	//
	// table
	TableViewMode *string `json:"tableViewMode,omitempty" xml:"tableViewMode,omitempty"`
	// example:
	//
	// 天
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 1
	Upper *string `json:"upper,omitempty" xml:"upper,omitempty"`
	// example:
	//
	// true
	VerticalPrint *bool `json:"verticalPrint,omitempty" xml:"verticalPrint,omitempty"`
}

func (s FormComponentProps) String() string {
	return tea.Prettify(s)
}

func (s FormComponentProps) GoString() string {
	return s.String()
}

func (s *FormComponentProps) SetActionName(v string) *FormComponentProps {
	s.ActionName = &v
	return s
}

func (s *FormComponentProps) SetAddressModel(v string) *FormComponentProps {
	s.AddressModel = &v
	return s
}

func (s *FormComponentProps) SetAlign(v string) *FormComponentProps {
	s.Align = &v
	return s
}

func (s *FormComponentProps) SetAsyncCondition(v bool) *FormComponentProps {
	s.AsyncCondition = &v
	return s
}

func (s *FormComponentProps) SetAvailableTemplates(v []*AvaliableTemplate) *FormComponentProps {
	s.AvailableTemplates = v
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

func (s *FormComponentProps) SetChoice(v string) *FormComponentProps {
	s.Choice = &v
	return s
}

func (s *FormComponentProps) SetCommonBizType(v string) *FormComponentProps {
	s.CommonBizType = &v
	return s
}

func (s *FormComponentProps) SetComponentId(v string) *FormComponentProps {
	s.ComponentId = &v
	return s
}

func (s *FormComponentProps) SetContent(v string) *FormComponentProps {
	s.Content = &v
	return s
}

func (s *FormComponentProps) SetDataSource(v *FormDataSource) *FormComponentProps {
	s.DataSource = v
	return s
}

func (s *FormComponentProps) SetDisabled(v bool) *FormComponentProps {
	s.Disabled = &v
	return s
}

func (s *FormComponentProps) SetDuration(v bool) *FormComponentProps {
	s.Duration = &v
	return s
}

func (s *FormComponentProps) SetDurationLabel(v string) *FormComponentProps {
	s.DurationLabel = &v
	return s
}

func (s *FormComponentProps) SetFormat(v string) *FormComponentProps {
	s.Format = &v
	return s
}

func (s *FormComponentProps) SetFormula(v string) *FormComponentProps {
	s.Formula = &v
	return s
}

func (s *FormComponentProps) SetInvisible(v bool) *FormComponentProps {
	s.Invisible = &v
	return s
}

func (s *FormComponentProps) SetLabel(v string) *FormComponentProps {
	s.Label = &v
	return s
}

func (s *FormComponentProps) SetLimit(v int32) *FormComponentProps {
	s.Limit = &v
	return s
}

func (s *FormComponentProps) SetLink(v string) *FormComponentProps {
	s.Link = &v
	return s
}

func (s *FormComponentProps) SetMaxLength(v int32) *FormComponentProps {
	s.MaxLength = &v
	return s
}

func (s *FormComponentProps) SetMode(v string) *FormComponentProps {
	s.Mode = &v
	return s
}

func (s *FormComponentProps) SetMultiple(v bool) *FormComponentProps {
	s.Multiple = &v
	return s
}

func (s *FormComponentProps) SetOptions(v []*SelectOption) *FormComponentProps {
	s.Options = v
	return s
}

func (s *FormComponentProps) SetPlaceholder(v string) *FormComponentProps {
	s.Placeholder = &v
	return s
}

func (s *FormComponentProps) SetPrecision(v int32) *FormComponentProps {
	s.Precision = &v
	return s
}

func (s *FormComponentProps) SetPrint(v string) *FormComponentProps {
	s.Print = &v
	return s
}

func (s *FormComponentProps) SetRequired(v bool) *FormComponentProps {
	s.Required = &v
	return s
}

func (s *FormComponentProps) SetStatField(v []*FormComponentPropsStatField) *FormComponentProps {
	s.StatField = v
	return s
}

func (s *FormComponentProps) SetTableViewMode(v string) *FormComponentProps {
	s.TableViewMode = &v
	return s
}

func (s *FormComponentProps) SetUnit(v string) *FormComponentProps {
	s.Unit = &v
	return s
}

func (s *FormComponentProps) SetUpper(v string) *FormComponentProps {
	s.Upper = &v
	return s
}

func (s *FormComponentProps) SetVerticalPrint(v bool) *FormComponentProps {
	s.VerticalPrint = &v
	return s
}

type FormComponentPropsStatField struct {
	// This parameter is required.
	//
	// example:
	//
	// NumberField-abcd
	ComponentId *string `json:"componentId,omitempty" xml:"componentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 金额
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 1
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
	// This parameter is required.
	Target *FormDataSourceTarget `json:"target,omitempty" xml:"target,omitempty" type:"Struct"`
	// example:
	//
	// form
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FormDataSource) String() string {
	return tea.Prettify(s)
}

func (s FormDataSource) GoString() string {
	return s.String()
}

func (s *FormDataSource) SetTarget(v *FormDataSourceTarget) *FormDataSource {
	s.Target = v
	return s
}

func (s *FormDataSource) SetType(v string) *FormDataSource {
	s.Type = &v
	return s
}

type FormDataSourceTarget struct {
	// This parameter is required.
	//
	// example:
	//
	// 0
	AppType *int32 `json:"appType,omitempty" xml:"appType,omitempty"`
	// example:
	//
	// SWAPP-abcd
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
}

func (s FormDataSourceTarget) String() string {
	return tea.Prettify(s)
}

func (s FormDataSourceTarget) GoString() string {
	return s.String()
}

func (s *FormDataSourceTarget) SetAppType(v int32) *FormDataSourceTarget {
	s.AppType = &v
	return s
}

func (s *FormDataSourceTarget) SetAppUuid(v string) *FormDataSourceTarget {
	s.AppUuid = &v
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
	// example:
	//
	// finance
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// 财务
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

type ResultValue struct {
	Result  *bool   `json:"result,omitempty" xml:"result,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ResultValue) String() string {
	return tea.Prettify(s)
}

func (s ResultValue) GoString() string {
	return s.String()
}

func (s *ResultValue) SetResult(v bool) *ResultValue {
	s.Result = &v
	return s
}

func (s *ResultValue) SetMessage(v string) *ResultValue {
	s.Message = &v
	return s
}

type AddApproveDentryAuthHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddApproveDentryAuthHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddApproveDentryAuthHeaders) GoString() string {
	return s.String()
}

func (s *AddApproveDentryAuthHeaders) SetCommonHeaders(v map[string]*string) *AddApproveDentryAuthHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddApproveDentryAuthHeaders) SetXAcsDingtalkAccessToken(v string) *AddApproveDentryAuthHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddApproveDentryAuthRequest struct {
	// This parameter is required.
	FileInfos []*AddApproveDentryAuthRequestFileInfos `json:"fileInfos,omitempty" xml:"fileInfos,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// user123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AddApproveDentryAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s AddApproveDentryAuthRequest) GoString() string {
	return s.String()
}

func (s *AddApproveDentryAuthRequest) SetFileInfos(v []*AddApproveDentryAuthRequestFileInfos) *AddApproveDentryAuthRequest {
	s.FileInfos = v
	return s
}

func (s *AddApproveDentryAuthRequest) SetUserId(v string) *AddApproveDentryAuthRequest {
	s.UserId = &v
	return s
}

type AddApproveDentryAuthRequestFileInfos struct {
	// This parameter is required.
	//
	// example:
	//
	// B1oQixxxx
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111
	SpaceId *int64 `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s AddApproveDentryAuthRequestFileInfos) String() string {
	return tea.Prettify(s)
}

func (s AddApproveDentryAuthRequestFileInfos) GoString() string {
	return s.String()
}

func (s *AddApproveDentryAuthRequestFileInfos) SetFileId(v string) *AddApproveDentryAuthRequestFileInfos {
	s.FileId = &v
	return s
}

func (s *AddApproveDentryAuthRequestFileInfos) SetSpaceId(v int64) *AddApproveDentryAuthRequestFileInfos {
	s.SpaceId = &v
	return s
}

type AddApproveDentryAuthResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddApproveDentryAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddApproveDentryAuthResponseBody) GoString() string {
	return s.String()
}

func (s *AddApproveDentryAuthResponseBody) SetResult(v bool) *AddApproveDentryAuthResponseBody {
	s.Result = &v
	return s
}

func (s *AddApproveDentryAuthResponseBody) SetSuccess(v bool) *AddApproveDentryAuthResponseBody {
	s.Success = &v
	return s
}

type AddApproveDentryAuthResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddApproveDentryAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddApproveDentryAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s AddApproveDentryAuthResponse) GoString() string {
	return s.String()
}

func (s *AddApproveDentryAuthResponse) SetHeaders(v map[string]*string) *AddApproveDentryAuthResponse {
	s.Headers = v
	return s
}

func (s *AddApproveDentryAuthResponse) SetStatusCode(v int32) *AddApproveDentryAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *AddApproveDentryAuthResponse) SetBody(v *AddApproveDentryAuthResponseBody) *AddApproveDentryAuthResponse {
	s.Body = v
	return s
}

type AddProcessInstanceCommentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s AddProcessInstanceCommentHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddProcessInstanceCommentHeaders) GoString() string {
	return s.String()
}

func (s *AddProcessInstanceCommentHeaders) SetCommonHeaders(v map[string]*string) *AddProcessInstanceCommentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddProcessInstanceCommentHeaders) SetXAcsDingtalkAccessToken(v string) *AddProcessInstanceCommentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type AddProcessInstanceCommentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// user123
	CommentUserId *string                               `json:"commentUserId,omitempty" xml:"commentUserId,omitempty"`
	File          *AddProcessInstanceCommentRequestFile `json:"file,omitempty" xml:"file,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// a171de6c-8bxxxx
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 同意。
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s AddProcessInstanceCommentRequest) String() string {
	return tea.Prettify(s)
}

func (s AddProcessInstanceCommentRequest) GoString() string {
	return s.String()
}

func (s *AddProcessInstanceCommentRequest) SetCommentUserId(v string) *AddProcessInstanceCommentRequest {
	s.CommentUserId = &v
	return s
}

func (s *AddProcessInstanceCommentRequest) SetFile(v *AddProcessInstanceCommentRequestFile) *AddProcessInstanceCommentRequest {
	s.File = v
	return s
}

func (s *AddProcessInstanceCommentRequest) SetProcessInstanceId(v string) *AddProcessInstanceCommentRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *AddProcessInstanceCommentRequest) SetText(v string) *AddProcessInstanceCommentRequest {
	s.Text = &v
	return s
}

type AddProcessInstanceCommentRequestFile struct {
	Attachments []*AddProcessInstanceCommentRequestFileAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	Photos      []*string                                          `json:"photos,omitempty" xml:"photos,omitempty" type:"Repeated"`
}

func (s AddProcessInstanceCommentRequestFile) String() string {
	return tea.Prettify(s)
}

func (s AddProcessInstanceCommentRequestFile) GoString() string {
	return s.String()
}

func (s *AddProcessInstanceCommentRequestFile) SetAttachments(v []*AddProcessInstanceCommentRequestFileAttachments) *AddProcessInstanceCommentRequestFile {
	s.Attachments = v
	return s
}

func (s *AddProcessInstanceCommentRequestFile) SetPhotos(v []*string) *AddProcessInstanceCommentRequestFile {
	s.Photos = v
	return s
}

type AddProcessInstanceCommentRequestFileAttachments struct {
	// example:
	//
	// B1oQixxxx
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// example:
	//
	// 文件名称。
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// 1024
	FileSize *string `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// example:
	//
	// file
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// 123
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s AddProcessInstanceCommentRequestFileAttachments) String() string {
	return tea.Prettify(s)
}

func (s AddProcessInstanceCommentRequestFileAttachments) GoString() string {
	return s.String()
}

func (s *AddProcessInstanceCommentRequestFileAttachments) SetFileId(v string) *AddProcessInstanceCommentRequestFileAttachments {
	s.FileId = &v
	return s
}

func (s *AddProcessInstanceCommentRequestFileAttachments) SetFileName(v string) *AddProcessInstanceCommentRequestFileAttachments {
	s.FileName = &v
	return s
}

func (s *AddProcessInstanceCommentRequestFileAttachments) SetFileSize(v string) *AddProcessInstanceCommentRequestFileAttachments {
	s.FileSize = &v
	return s
}

func (s *AddProcessInstanceCommentRequestFileAttachments) SetFileType(v string) *AddProcessInstanceCommentRequestFileAttachments {
	s.FileType = &v
	return s
}

func (s *AddProcessInstanceCommentRequestFileAttachments) SetSpaceId(v string) *AddProcessInstanceCommentRequestFileAttachments {
	s.SpaceId = &v
	return s
}

type AddProcessInstanceCommentResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddProcessInstanceCommentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddProcessInstanceCommentResponseBody) GoString() string {
	return s.String()
}

func (s *AddProcessInstanceCommentResponseBody) SetResult(v bool) *AddProcessInstanceCommentResponseBody {
	s.Result = &v
	return s
}

func (s *AddProcessInstanceCommentResponseBody) SetSuccess(v bool) *AddProcessInstanceCommentResponseBody {
	s.Success = &v
	return s
}

type AddProcessInstanceCommentResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddProcessInstanceCommentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddProcessInstanceCommentResponse) String() string {
	return tea.Prettify(s)
}

func (s AddProcessInstanceCommentResponse) GoString() string {
	return s.String()
}

func (s *AddProcessInstanceCommentResponse) SetHeaders(v map[string]*string) *AddProcessInstanceCommentResponse {
	s.Headers = v
	return s
}

func (s *AddProcessInstanceCommentResponse) SetStatusCode(v int32) *AddProcessInstanceCommentResponse {
	s.StatusCode = &v
	return s
}

func (s *AddProcessInstanceCommentResponse) SetBody(v *AddProcessInstanceCommentResponseBody) *AddProcessInstanceCommentResponse {
	s.Body = v
	return s
}

type ArchiveProcessInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ArchiveProcessInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s ArchiveProcessInstanceHeaders) GoString() string {
	return s.String()
}

func (s *ArchiveProcessInstanceHeaders) SetCommonHeaders(v map[string]*string) *ArchiveProcessInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ArchiveProcessInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *ArchiveProcessInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ArchiveProcessInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 133743186427339452
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a171de6c-8bxxxx
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
}

func (s ArchiveProcessInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ArchiveProcessInstanceRequest) GoString() string {
	return s.String()
}

func (s *ArchiveProcessInstanceRequest) SetOpUserId(v string) *ArchiveProcessInstanceRequest {
	s.OpUserId = &v
	return s
}

func (s *ArchiveProcessInstanceRequest) SetProcessInstanceId(v string) *ArchiveProcessInstanceRequest {
	s.ProcessInstanceId = &v
	return s
}

type ArchiveProcessInstanceResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ArchiveProcessInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ArchiveProcessInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ArchiveProcessInstanceResponseBody) SetResult(v bool) *ArchiveProcessInstanceResponseBody {
	s.Result = &v
	return s
}

func (s *ArchiveProcessInstanceResponseBody) SetSuccess(v bool) *ArchiveProcessInstanceResponseBody {
	s.Success = &v
	return s
}

type ArchiveProcessInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ArchiveProcessInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ArchiveProcessInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ArchiveProcessInstanceResponse) GoString() string {
	return s.String()
}

func (s *ArchiveProcessInstanceResponse) SetHeaders(v map[string]*string) *ArchiveProcessInstanceResponse {
	s.Headers = v
	return s
}

func (s *ArchiveProcessInstanceResponse) SetStatusCode(v int32) *ArchiveProcessInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ArchiveProcessInstanceResponse) SetBody(v *ArchiveProcessInstanceResponseBody) *ArchiveProcessInstanceResponse {
	s.Body = v
	return s
}

type BatchExecuteProcessInstancesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchExecuteProcessInstancesHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchExecuteProcessInstancesHeaders) GoString() string {
	return s.String()
}

func (s *BatchExecuteProcessInstancesHeaders) SetCommonHeaders(v map[string]*string) *BatchExecuteProcessInstancesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchExecuteProcessInstancesHeaders) SetXAcsDingtalkAccessToken(v string) *BatchExecuteProcessInstancesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchExecuteProcessInstancesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 67583405630
	ActionerUserId *string `json:"actionerUserId,omitempty" xml:"actionerUserId,omitempty"`
	Remark         *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agree
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// This parameter is required.
	TaskInfoList []*BatchExecuteProcessInstancesRequestTaskInfoList `json:"taskInfoList,omitempty" xml:"taskInfoList,omitempty" type:"Repeated"`
}

func (s BatchExecuteProcessInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchExecuteProcessInstancesRequest) GoString() string {
	return s.String()
}

func (s *BatchExecuteProcessInstancesRequest) SetActionerUserId(v string) *BatchExecuteProcessInstancesRequest {
	s.ActionerUserId = &v
	return s
}

func (s *BatchExecuteProcessInstancesRequest) SetRemark(v string) *BatchExecuteProcessInstancesRequest {
	s.Remark = &v
	return s
}

func (s *BatchExecuteProcessInstancesRequest) SetResult(v string) *BatchExecuteProcessInstancesRequest {
	s.Result = &v
	return s
}

func (s *BatchExecuteProcessInstancesRequest) SetTaskInfoList(v []*BatchExecuteProcessInstancesRequestTaskInfoList) *BatchExecuteProcessInstancesRequest {
	s.TaskInfoList = v
	return s
}

type BatchExecuteProcessInstancesRequestTaskInfoList struct {
	// This parameter is required.
	//
	// example:
	//
	// a171de6c-8bxxxx
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// This parameter is required.
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s BatchExecuteProcessInstancesRequestTaskInfoList) String() string {
	return tea.Prettify(s)
}

func (s BatchExecuteProcessInstancesRequestTaskInfoList) GoString() string {
	return s.String()
}

func (s *BatchExecuteProcessInstancesRequestTaskInfoList) SetProcessInstanceId(v string) *BatchExecuteProcessInstancesRequestTaskInfoList {
	s.ProcessInstanceId = &v
	return s
}

func (s *BatchExecuteProcessInstancesRequestTaskInfoList) SetTaskId(v int64) *BatchExecuteProcessInstancesRequestTaskInfoList {
	s.TaskId = &v
	return s
}

type BatchExecuteProcessInstancesResponseBody struct {
	Result  map[string]*ResultValue `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool                   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchExecuteProcessInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchExecuteProcessInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchExecuteProcessInstancesResponseBody) SetResult(v map[string]*ResultValue) *BatchExecuteProcessInstancesResponseBody {
	s.Result = v
	return s
}

func (s *BatchExecuteProcessInstancesResponseBody) SetSuccess(v bool) *BatchExecuteProcessInstancesResponseBody {
	s.Success = &v
	return s
}

type BatchExecuteProcessInstancesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchExecuteProcessInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchExecuteProcessInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchExecuteProcessInstancesResponse) GoString() string {
	return s.String()
}

func (s *BatchExecuteProcessInstancesResponse) SetHeaders(v map[string]*string) *BatchExecuteProcessInstancesResponse {
	s.Headers = v
	return s
}

func (s *BatchExecuteProcessInstancesResponse) SetStatusCode(v int32) *BatchExecuteProcessInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchExecuteProcessInstancesResponse) SetBody(v *BatchExecuteProcessInstancesResponseBody) *BatchExecuteProcessInstancesResponse {
	s.Body = v
	return s
}

type BatchTasksRedirectHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchTasksRedirectHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchTasksRedirectHeaders) GoString() string {
	return s.String()
}

func (s *BatchTasksRedirectHeaders) SetCommonHeaders(v map[string]*string) *BatchTasksRedirectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchTasksRedirectHeaders) SetXAcsDingtalkAccessToken(v string) *BatchTasksRedirectHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchTasksRedirectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// staffId-B
	HandoverUserId *string `json:"handoverUserId,omitempty" xml:"handoverUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// manager-12
	ManagerUserId *string `json:"managerUserId,omitempty" xml:"managerUserId,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	TaskIds []*int64 `json:"taskIds,omitempty" xml:"taskIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// staffId-A
	TransfereeUserId *string `json:"transfereeUserId,omitempty" xml:"transfereeUserId,omitempty"`
}

func (s BatchTasksRedirectRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchTasksRedirectRequest) GoString() string {
	return s.String()
}

func (s *BatchTasksRedirectRequest) SetHandoverUserId(v string) *BatchTasksRedirectRequest {
	s.HandoverUserId = &v
	return s
}

func (s *BatchTasksRedirectRequest) SetManagerUserId(v string) *BatchTasksRedirectRequest {
	s.ManagerUserId = &v
	return s
}

func (s *BatchTasksRedirectRequest) SetTaskIds(v []*int64) *BatchTasksRedirectRequest {
	s.TaskIds = v
	return s
}

func (s *BatchTasksRedirectRequest) SetTransfereeUserId(v string) *BatchTasksRedirectRequest {
	s.TransfereeUserId = &v
	return s
}

type BatchTasksRedirectResponseBody struct {
	Result  *BatchTasksRedirectResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                 `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchTasksRedirectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchTasksRedirectResponseBody) GoString() string {
	return s.String()
}

func (s *BatchTasksRedirectResponseBody) SetResult(v *BatchTasksRedirectResponseBodyResult) *BatchTasksRedirectResponseBody {
	s.Result = v
	return s
}

func (s *BatchTasksRedirectResponseBody) SetSuccess(v bool) *BatchTasksRedirectResponseBody {
	s.Success = &v
	return s
}

type BatchTasksRedirectResponseBodyResult struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	FailCount *int64 `json:"failCount,omitempty" xml:"failCount,omitempty"`
	// This parameter is required.
	RedirectResults []*BatchTasksRedirectResponseBodyResultRedirectResults `json:"redirectResults,omitempty" xml:"redirectResults,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s BatchTasksRedirectResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s BatchTasksRedirectResponseBodyResult) GoString() string {
	return s.String()
}

func (s *BatchTasksRedirectResponseBodyResult) SetFailCount(v int64) *BatchTasksRedirectResponseBodyResult {
	s.FailCount = &v
	return s
}

func (s *BatchTasksRedirectResponseBodyResult) SetRedirectResults(v []*BatchTasksRedirectResponseBodyResultRedirectResults) *BatchTasksRedirectResponseBodyResult {
	s.RedirectResults = v
	return s
}

func (s *BatchTasksRedirectResponseBodyResult) SetTotalCount(v int64) *BatchTasksRedirectResponseBodyResult {
	s.TotalCount = &v
	return s
}

type BatchTasksRedirectResponseBodyResultRedirectResults struct {
	// example:
	//
	// 外部流程不允许转交
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s BatchTasksRedirectResponseBodyResultRedirectResults) String() string {
	return tea.Prettify(s)
}

func (s BatchTasksRedirectResponseBodyResultRedirectResults) GoString() string {
	return s.String()
}

func (s *BatchTasksRedirectResponseBodyResultRedirectResults) SetErrorMsg(v string) *BatchTasksRedirectResponseBodyResultRedirectResults {
	s.ErrorMsg = &v
	return s
}

func (s *BatchTasksRedirectResponseBodyResultRedirectResults) SetSuccess(v bool) *BatchTasksRedirectResponseBodyResultRedirectResults {
	s.Success = &v
	return s
}

func (s *BatchTasksRedirectResponseBodyResultRedirectResults) SetTaskId(v int64) *BatchTasksRedirectResponseBodyResultRedirectResults {
	s.TaskId = &v
	return s
}

type BatchTasksRedirectResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchTasksRedirectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchTasksRedirectResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchTasksRedirectResponse) GoString() string {
	return s.String()
}

func (s *BatchTasksRedirectResponse) SetHeaders(v map[string]*string) *BatchTasksRedirectResponse {
	s.Headers = v
	return s
}

func (s *BatchTasksRedirectResponse) SetStatusCode(v int32) *BatchTasksRedirectResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchTasksRedirectResponse) SetBody(v *BatchTasksRedirectResponseBody) *BatchTasksRedirectResponse {
	s.Body = v
	return s
}

type BatchUpdateProcessInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s BatchUpdateProcessInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateProcessInstanceHeaders) GoString() string {
	return s.String()
}

func (s *BatchUpdateProcessInstanceHeaders) SetCommonHeaders(v map[string]*string) *BatchUpdateProcessInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchUpdateProcessInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *BatchUpdateProcessInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type BatchUpdateProcessInstanceRequest struct {
	UpdateProcessInstanceRequests []*BatchUpdateProcessInstanceRequestUpdateProcessInstanceRequests `json:"updateProcessInstanceRequests,omitempty" xml:"updateProcessInstanceRequests,omitempty" type:"Repeated"`
}

func (s BatchUpdateProcessInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateProcessInstanceRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateProcessInstanceRequest) SetUpdateProcessInstanceRequests(v []*BatchUpdateProcessInstanceRequestUpdateProcessInstanceRequests) *BatchUpdateProcessInstanceRequest {
	s.UpdateProcessInstanceRequests = v
	return s
}

type BatchUpdateProcessInstanceRequestUpdateProcessInstanceRequests struct {
	Notifiers []*BatchUpdateProcessInstanceRequestUpdateProcessInstanceRequestsNotifiers `json:"notifiers,omitempty" xml:"notifiers,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// abc
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// example:
	//
	// agree
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// COMPLETED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s BatchUpdateProcessInstanceRequestUpdateProcessInstanceRequests) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateProcessInstanceRequestUpdateProcessInstanceRequests) GoString() string {
	return s.String()
}

func (s *BatchUpdateProcessInstanceRequestUpdateProcessInstanceRequests) SetNotifiers(v []*BatchUpdateProcessInstanceRequestUpdateProcessInstanceRequestsNotifiers) *BatchUpdateProcessInstanceRequestUpdateProcessInstanceRequests {
	s.Notifiers = v
	return s
}

func (s *BatchUpdateProcessInstanceRequestUpdateProcessInstanceRequests) SetProcessInstanceId(v string) *BatchUpdateProcessInstanceRequestUpdateProcessInstanceRequests {
	s.ProcessInstanceId = &v
	return s
}

func (s *BatchUpdateProcessInstanceRequestUpdateProcessInstanceRequests) SetResult(v string) *BatchUpdateProcessInstanceRequestUpdateProcessInstanceRequests {
	s.Result = &v
	return s
}

func (s *BatchUpdateProcessInstanceRequestUpdateProcessInstanceRequests) SetStatus(v string) *BatchUpdateProcessInstanceRequestUpdateProcessInstanceRequests {
	s.Status = &v
	return s
}

type BatchUpdateProcessInstanceRequestUpdateProcessInstanceRequestsNotifiers struct {
	// This parameter is required.
	//
	// example:
	//
	// manager001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s BatchUpdateProcessInstanceRequestUpdateProcessInstanceRequestsNotifiers) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateProcessInstanceRequestUpdateProcessInstanceRequestsNotifiers) GoString() string {
	return s.String()
}

func (s *BatchUpdateProcessInstanceRequestUpdateProcessInstanceRequestsNotifiers) SetUserId(v string) *BatchUpdateProcessInstanceRequestUpdateProcessInstanceRequestsNotifiers {
	s.UserId = &v
	return s
}

type BatchUpdateProcessInstanceResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchUpdateProcessInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateProcessInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdateProcessInstanceResponseBody) SetSuccess(v bool) *BatchUpdateProcessInstanceResponseBody {
	s.Success = &v
	return s
}

type BatchUpdateProcessInstanceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchUpdateProcessInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchUpdateProcessInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateProcessInstanceResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateProcessInstanceResponse) SetHeaders(v map[string]*string) *BatchUpdateProcessInstanceResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateProcessInstanceResponse) SetStatusCode(v int32) *BatchUpdateProcessInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUpdateProcessInstanceResponse) SetBody(v *BatchUpdateProcessInstanceResponseBody) *BatchUpdateProcessInstanceResponse {
	s.Body = v
	return s
}

type CancelIntegratedTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CancelIntegratedTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s CancelIntegratedTaskHeaders) GoString() string {
	return s.String()
}

func (s *CancelIntegratedTaskHeaders) SetCommonHeaders(v map[string]*string) *CancelIntegratedTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CancelIntegratedTaskHeaders) SetXAcsDingtalkAccessToken(v string) *CancelIntegratedTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CancelIntegratedTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// act_xxxx
	ActivityId *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
	// if can be null:
	// true
	ActivityIds []*string `json:"activityIds,omitempty" xml:"activityIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// tPr_FB_mT_xxxxxxxxx2hQ05201655306463
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
}

func (s CancelIntegratedTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelIntegratedTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelIntegratedTaskRequest) SetActivityId(v string) *CancelIntegratedTaskRequest {
	s.ActivityId = &v
	return s
}

func (s *CancelIntegratedTaskRequest) SetActivityIds(v []*string) *CancelIntegratedTaskRequest {
	s.ActivityIds = v
	return s
}

func (s *CancelIntegratedTaskRequest) SetProcessInstanceId(v string) *CancelIntegratedTaskRequest {
	s.ProcessInstanceId = &v
	return s
}

type CancelIntegratedTaskResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CancelIntegratedTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelIntegratedTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelIntegratedTaskResponseBody) SetSuccess(v bool) *CancelIntegratedTaskResponseBody {
	s.Success = &v
	return s
}

type CancelIntegratedTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelIntegratedTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelIntegratedTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelIntegratedTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelIntegratedTaskResponse) SetHeaders(v map[string]*string) *CancelIntegratedTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelIntegratedTaskResponse) SetStatusCode(v int32) *CancelIntegratedTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelIntegratedTaskResponse) SetBody(v *CancelIntegratedTaskResponseBody) *CancelIntegratedTaskResponse {
	s.Body = v
	return s
}

type CleanProcessDataHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CleanProcessDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s CleanProcessDataHeaders) GoString() string {
	return s.String()
}

func (s *CleanProcessDataHeaders) SetCommonHeaders(v map[string]*string) *CleanProcessDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CleanProcessDataHeaders) SetXAcsDingtalkAccessToken(v string) *CleanProcessDataHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CleanProcessDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ding1234
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-EF6YJL35
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
}

func (s CleanProcessDataRequest) String() string {
	return tea.Prettify(s)
}

func (s CleanProcessDataRequest) GoString() string {
	return s.String()
}

func (s *CleanProcessDataRequest) SetCorpId(v string) *CleanProcessDataRequest {
	s.CorpId = &v
	return s
}

func (s *CleanProcessDataRequest) SetProcessCode(v string) *CleanProcessDataRequest {
	s.ProcessCode = &v
	return s
}

type CleanProcessDataResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CleanProcessDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CleanProcessDataResponseBody) GoString() string {
	return s.String()
}

func (s *CleanProcessDataResponseBody) SetSuccess(v bool) *CleanProcessDataResponseBody {
	s.Success = &v
	return s
}

type CleanProcessDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CleanProcessDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CleanProcessDataResponse) String() string {
	return tea.Prettify(s)
}

func (s CleanProcessDataResponse) GoString() string {
	return s.String()
}

func (s *CleanProcessDataResponse) SetHeaders(v map[string]*string) *CleanProcessDataResponse {
	s.Headers = v
	return s
}

func (s *CleanProcessDataResponse) SetStatusCode(v int32) *CleanProcessDataResponse {
	s.StatusCode = &v
	return s
}

func (s *CleanProcessDataResponse) SetBody(v *CleanProcessDataResponseBody) *CleanProcessDataResponse {
	s.Body = v
	return s
}

type CopyProcessHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CopyProcessHeaders) String() string {
	return tea.Prettify(s)
}

func (s CopyProcessHeaders) GoString() string {
	return s.String()
}

func (s *CopyProcessHeaders) SetCommonHeaders(v map[string]*string) *CopyProcessHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CopyProcessHeaders) SetXAcsDingtalkAccessToken(v string) *CopyProcessHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CopyProcessRequest struct {
	CopyOptions *CopyProcessRequestCopyOptions `json:"copyOptions,omitempty" xml:"copyOptions,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// dingabc
	SourceCorpId *string `json:"sourceCorpId,omitempty" xml:"sourceCorpId,omitempty"`
	// This parameter is required.
	SourceProcessVOList []*CopyProcessRequestSourceProcessVOList `json:"sourceProcessVOList,omitempty" xml:"sourceProcessVOList,omitempty" type:"Repeated"`
}

func (s CopyProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s CopyProcessRequest) GoString() string {
	return s.String()
}

func (s *CopyProcessRequest) SetCopyOptions(v *CopyProcessRequestCopyOptions) *CopyProcessRequest {
	s.CopyOptions = v
	return s
}

func (s *CopyProcessRequest) SetSourceCorpId(v string) *CopyProcessRequest {
	s.SourceCorpId = &v
	return s
}

func (s *CopyProcessRequest) SetSourceProcessVOList(v []*CopyProcessRequestSourceProcessVOList) *CopyProcessRequest {
	s.SourceProcessVOList = v
	return s
}

type CopyProcessRequestCopyOptions struct {
	// example:
	//
	// 1
	CopyType *int32 `json:"copyType,omitempty" xml:"copyType,omitempty"`
}

func (s CopyProcessRequestCopyOptions) String() string {
	return tea.Prettify(s)
}

func (s CopyProcessRequestCopyOptions) GoString() string {
	return s.String()
}

func (s *CopyProcessRequestCopyOptions) SetCopyType(v int32) *CopyProcessRequestCopyOptions {
	s.CopyType = &v
	return s
}

type CopyProcessRequestSourceProcessVOList struct {
	// example:
	//
	// abc
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// example:
	//
	// abc
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// proc-abc
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
}

func (s CopyProcessRequestSourceProcessVOList) String() string {
	return tea.Prettify(s)
}

func (s CopyProcessRequestSourceProcessVOList) GoString() string {
	return s.String()
}

func (s *CopyProcessRequestSourceProcessVOList) SetBizType(v string) *CopyProcessRequestSourceProcessVOList {
	s.BizType = &v
	return s
}

func (s *CopyProcessRequestSourceProcessVOList) SetName(v string) *CopyProcessRequestSourceProcessVOList {
	s.Name = &v
	return s
}

func (s *CopyProcessRequestSourceProcessVOList) SetProcessCode(v string) *CopyProcessRequestSourceProcessVOList {
	s.ProcessCode = &v
	return s
}

type CopyProcessResponseBody struct {
	Result []*CopyProcessResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s CopyProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CopyProcessResponseBody) GoString() string {
	return s.String()
}

func (s *CopyProcessResponseBody) SetResult(v []*CopyProcessResponseBodyResult) *CopyProcessResponseBody {
	s.Result = v
	return s
}

type CopyProcessResponseBodyResult struct {
	// example:
	//
	// abc
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// example:
	//
	// abc
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// proc-abc
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
}

func (s CopyProcessResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CopyProcessResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CopyProcessResponseBodyResult) SetBizType(v string) *CopyProcessResponseBodyResult {
	s.BizType = &v
	return s
}

func (s *CopyProcessResponseBodyResult) SetName(v string) *CopyProcessResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CopyProcessResponseBodyResult) SetProcessCode(v string) *CopyProcessResponseBodyResult {
	s.ProcessCode = &v
	return s
}

type CopyProcessResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s CopyProcessResponse) GoString() string {
	return s.String()
}

func (s *CopyProcessResponse) SetHeaders(v map[string]*string) *CopyProcessResponse {
	s.Headers = v
	return s
}

func (s *CopyProcessResponse) SetStatusCode(v int32) *CopyProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyProcessResponse) SetBody(v *CopyProcessResponseBody) *CopyProcessResponse {
	s.Body = v
	return s
}

type CreateIntegratedTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s CreateIntegratedTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateIntegratedTaskHeaders) GoString() string {
	return s.String()
}

func (s *CreateIntegratedTaskHeaders) SetCommonHeaders(v map[string]*string) *CreateIntegratedTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateIntegratedTaskHeaders) SetXAcsDingtalkAccessToken(v string) *CreateIntegratedTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type CreateIntegratedTaskRequest struct {
	// example:
	//
	// act_xxxx
	ActivityId    *string                                   `json:"activityId,omitempty" xml:"activityId,omitempty"`
	FeatureConfig *CreateIntegratedTaskRequestFeatureConfig `json:"featureConfig,omitempty" xml:"featureConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// tPr_FB_mT_xxxxxxxxx2hQ05201655306463
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// This parameter is required.
	Tasks []*CreateIntegratedTaskRequestTasks `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Repeated"`
}

func (s CreateIntegratedTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIntegratedTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateIntegratedTaskRequest) SetActivityId(v string) *CreateIntegratedTaskRequest {
	s.ActivityId = &v
	return s
}

func (s *CreateIntegratedTaskRequest) SetFeatureConfig(v *CreateIntegratedTaskRequestFeatureConfig) *CreateIntegratedTaskRequest {
	s.FeatureConfig = v
	return s
}

func (s *CreateIntegratedTaskRequest) SetProcessInstanceId(v string) *CreateIntegratedTaskRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *CreateIntegratedTaskRequest) SetTasks(v []*CreateIntegratedTaskRequestTasks) *CreateIntegratedTaskRequest {
	s.Tasks = v
	return s
}

type CreateIntegratedTaskRequestFeatureConfig struct {
	Features []*CreateIntegratedTaskRequestFeatureConfigFeatures `json:"features,omitempty" xml:"features,omitempty" type:"Repeated"`
}

func (s CreateIntegratedTaskRequestFeatureConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateIntegratedTaskRequestFeatureConfig) GoString() string {
	return s.String()
}

func (s *CreateIntegratedTaskRequestFeatureConfig) SetFeatures(v []*CreateIntegratedTaskRequestFeatureConfigFeatures) *CreateIntegratedTaskRequestFeatureConfig {
	s.Features = v
	return s
}

type CreateIntegratedTaskRequestFeatureConfigFeatures struct {
	Callback *CreateIntegratedTaskRequestFeatureConfigFeaturesCallback `json:"callback,omitempty" xml:"callback,omitempty" type:"Struct"`
	// if can be null:
	// true
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// example:
	//
	// www.dingtalk.com
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	// example:
	//
	// abc
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// www.dingtalk.com
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	// example:
	//
	// ORIGIN
	RunType *string `json:"runType,omitempty" xml:"runType,omitempty"`
}

func (s CreateIntegratedTaskRequestFeatureConfigFeatures) String() string {
	return tea.Prettify(s)
}

func (s CreateIntegratedTaskRequestFeatureConfigFeatures) GoString() string {
	return s.String()
}

func (s *CreateIntegratedTaskRequestFeatureConfigFeatures) SetCallback(v *CreateIntegratedTaskRequestFeatureConfigFeaturesCallback) *CreateIntegratedTaskRequestFeatureConfigFeatures {
	s.Callback = v
	return s
}

func (s *CreateIntegratedTaskRequestFeatureConfigFeatures) SetConfig(v string) *CreateIntegratedTaskRequestFeatureConfigFeatures {
	s.Config = &v
	return s
}

func (s *CreateIntegratedTaskRequestFeatureConfigFeatures) SetMobileUrl(v string) *CreateIntegratedTaskRequestFeatureConfigFeatures {
	s.MobileUrl = &v
	return s
}

func (s *CreateIntegratedTaskRequestFeatureConfigFeatures) SetName(v string) *CreateIntegratedTaskRequestFeatureConfigFeatures {
	s.Name = &v
	return s
}

func (s *CreateIntegratedTaskRequestFeatureConfigFeatures) SetPcUrl(v string) *CreateIntegratedTaskRequestFeatureConfigFeatures {
	s.PcUrl = &v
	return s
}

func (s *CreateIntegratedTaskRequestFeatureConfigFeatures) SetRunType(v string) *CreateIntegratedTaskRequestFeatureConfigFeatures {
	s.RunType = &v
	return s
}

type CreateIntegratedTaskRequestFeatureConfigFeaturesCallback struct {
	// example:
	//
	// abc
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// example:
	//
	// abc
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CreateIntegratedTaskRequestFeatureConfigFeaturesCallback) String() string {
	return tea.Prettify(s)
}

func (s CreateIntegratedTaskRequestFeatureConfigFeaturesCallback) GoString() string {
	return s.String()
}

func (s *CreateIntegratedTaskRequestFeatureConfigFeaturesCallback) SetApiKey(v string) *CreateIntegratedTaskRequestFeatureConfigFeaturesCallback {
	s.ApiKey = &v
	return s
}

func (s *CreateIntegratedTaskRequestFeatureConfigFeaturesCallback) SetAppUuid(v string) *CreateIntegratedTaskRequestFeatureConfigFeaturesCallback {
	s.AppUuid = &v
	return s
}

func (s *CreateIntegratedTaskRequestFeatureConfigFeaturesCallback) SetVersion(v string) *CreateIntegratedTaskRequestFeatureConfigFeaturesCallback {
	s.Version = &v
	return s
}

type CreateIntegratedTaskRequestTasks struct {
	// example:
	//
	// {\"id\":\"12345\"}
	CustomData *string `json:"customData,omitempty" xml:"customData,omitempty"`
	// example:
	//
	// 1758643200000
	DueTimestamp *int64 `json:"dueTimestamp,omitempty" xml:"dueTimestamp,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// example:
	//
	// manager001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateIntegratedTaskRequestTasks) String() string {
	return tea.Prettify(s)
}

func (s CreateIntegratedTaskRequestTasks) GoString() string {
	return s.String()
}

func (s *CreateIntegratedTaskRequestTasks) SetCustomData(v string) *CreateIntegratedTaskRequestTasks {
	s.CustomData = &v
	return s
}

func (s *CreateIntegratedTaskRequestTasks) SetDueTimestamp(v int64) *CreateIntegratedTaskRequestTasks {
	s.DueTimestamp = &v
	return s
}

func (s *CreateIntegratedTaskRequestTasks) SetUrl(v string) *CreateIntegratedTaskRequestTasks {
	s.Url = &v
	return s
}

func (s *CreateIntegratedTaskRequestTasks) SetUserId(v string) *CreateIntegratedTaskRequestTasks {
	s.UserId = &v
	return s
}

type CreateIntegratedTaskResponseBody struct {
	Result  []*CreateIntegratedTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success *bool                                     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateIntegratedTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIntegratedTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIntegratedTaskResponseBody) SetResult(v []*CreateIntegratedTaskResponseBodyResult) *CreateIntegratedTaskResponseBody {
	s.Result = v
	return s
}

func (s *CreateIntegratedTaskResponseBody) SetSuccess(v bool) *CreateIntegratedTaskResponseBody {
	s.Success = &v
	return s
}

type CreateIntegratedTaskResponseBodyResult struct {
	TaskId *int64  `json:"taskId,omitempty" xml:"taskId,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateIntegratedTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateIntegratedTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateIntegratedTaskResponseBodyResult) SetTaskId(v int64) *CreateIntegratedTaskResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *CreateIntegratedTaskResponseBodyResult) SetUserId(v string) *CreateIntegratedTaskResponseBodyResult {
	s.UserId = &v
	return s
}

type CreateIntegratedTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIntegratedTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIntegratedTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIntegratedTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateIntegratedTaskResponse) SetHeaders(v map[string]*string) *CreateIntegratedTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateIntegratedTaskResponse) SetStatusCode(v int32) *CreateIntegratedTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIntegratedTaskResponse) SetBody(v *CreateIntegratedTaskResponseBody) *CreateIntegratedTaskResponse {
	s.Body = v
	return s
}

type DeleteDirHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteDirHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirHeaders) GoString() string {
	return s.String()
}

func (s *DeleteDirHeaders) SetCommonHeaders(v map[string]*string) *DeleteDirHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteDirHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteDirHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteDirRequest struct {
	// example:
	//
	// oaDirIdxxx
	DirId *string `json:"dirId,omitempty" xml:"dirId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user001
	OperateUserId *string `json:"operateUserId,omitempty" xml:"operateUserId,omitempty"`
}

func (s DeleteDirRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirRequest) GoString() string {
	return s.String()
}

func (s *DeleteDirRequest) SetDirId(v string) *DeleteDirRequest {
	s.DirId = &v
	return s
}

func (s *DeleteDirRequest) SetOperateUserId(v string) *DeleteDirRequest {
	s.OperateUserId = &v
	return s
}

type DeleteDirResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteDirResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDirResponseBody) SetSuccess(v bool) *DeleteDirResponseBody {
	s.Success = &v
	return s
}

type DeleteDirResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDirResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDirResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDirResponse) GoString() string {
	return s.String()
}

func (s *DeleteDirResponse) SetHeaders(v map[string]*string) *DeleteDirResponse {
	s.Headers = v
	return s
}

func (s *DeleteDirResponse) SetStatusCode(v int32) *DeleteDirResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDirResponse) SetBody(v *DeleteDirResponseBody) *DeleteDirResponse {
	s.Body = v
	return s
}

type DeleteProcessHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s DeleteProcessHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteProcessHeaders) GoString() string {
	return s.String()
}

func (s *DeleteProcessHeaders) SetCommonHeaders(v map[string]*string) *DeleteProcessHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteProcessHeaders) SetXAcsDingtalkAccessToken(v string) *DeleteProcessHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type DeleteProcessRequest struct {
	CleanRunningTask *bool `json:"cleanRunningTask,omitempty" xml:"cleanRunningTask,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// proc-abc
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
}

func (s DeleteProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProcessRequest) GoString() string {
	return s.String()
}

func (s *DeleteProcessRequest) SetCleanRunningTask(v bool) *DeleteProcessRequest {
	s.CleanRunningTask = &v
	return s
}

func (s *DeleteProcessRequest) SetProcessCode(v string) *DeleteProcessRequest {
	s.ProcessCode = &v
	return s
}

type DeleteProcessResponseBody struct {
	Result *DeleteProcessResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DeleteProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProcessResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProcessResponseBody) SetResult(v *DeleteProcessResponseBodyResult) *DeleteProcessResponseBody {
	s.Result = v
	return s
}

type DeleteProcessResponseBodyResult struct {
	// example:
	//
	// proc-abc
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
}

func (s DeleteProcessResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteProcessResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteProcessResponseBodyResult) SetProcessCode(v string) *DeleteProcessResponseBodyResult {
	s.ProcessCode = &v
	return s
}

type DeleteProcessResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProcessResponse) GoString() string {
	return s.String()
}

func (s *DeleteProcessResponse) SetHeaders(v map[string]*string) *DeleteProcessResponse {
	s.Headers = v
	return s
}

func (s *DeleteProcessResponse) SetStatusCode(v int32) *DeleteProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProcessResponse) SetBody(v *DeleteProcessResponseBody) *DeleteProcessResponse {
	s.Body = v
	return s
}

type ExecuteProcessInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ExecuteProcessInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s ExecuteProcessInstanceHeaders) GoString() string {
	return s.String()
}

func (s *ExecuteProcessInstanceHeaders) SetCommonHeaders(v map[string]*string) *ExecuteProcessInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ExecuteProcessInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *ExecuteProcessInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ExecuteProcessInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 133743186427339452
	ActionerUserId *string                            `json:"actionerUserId,omitempty" xml:"actionerUserId,omitempty"`
	File           *ExecuteProcessInstanceRequestFile `json:"file,omitempty" xml:"file,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// a171de6c-8bxxxx
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// example:
	//
	// 同意。
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agree
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 67583405630
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ExecuteProcessInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteProcessInstanceRequest) GoString() string {
	return s.String()
}

func (s *ExecuteProcessInstanceRequest) SetActionerUserId(v string) *ExecuteProcessInstanceRequest {
	s.ActionerUserId = &v
	return s
}

func (s *ExecuteProcessInstanceRequest) SetFile(v *ExecuteProcessInstanceRequestFile) *ExecuteProcessInstanceRequest {
	s.File = v
	return s
}

func (s *ExecuteProcessInstanceRequest) SetProcessInstanceId(v string) *ExecuteProcessInstanceRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *ExecuteProcessInstanceRequest) SetRemark(v string) *ExecuteProcessInstanceRequest {
	s.Remark = &v
	return s
}

func (s *ExecuteProcessInstanceRequest) SetResult(v string) *ExecuteProcessInstanceRequest {
	s.Result = &v
	return s
}

func (s *ExecuteProcessInstanceRequest) SetTaskId(v int64) *ExecuteProcessInstanceRequest {
	s.TaskId = &v
	return s
}

type ExecuteProcessInstanceRequestFile struct {
	Attachments []*ExecuteProcessInstanceRequestFileAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	Photos      []*string                                       `json:"photos,omitempty" xml:"photos,omitempty" type:"Repeated"`
}

func (s ExecuteProcessInstanceRequestFile) String() string {
	return tea.Prettify(s)
}

func (s ExecuteProcessInstanceRequestFile) GoString() string {
	return s.String()
}

func (s *ExecuteProcessInstanceRequestFile) SetAttachments(v []*ExecuteProcessInstanceRequestFileAttachments) *ExecuteProcessInstanceRequestFile {
	s.Attachments = v
	return s
}

func (s *ExecuteProcessInstanceRequestFile) SetPhotos(v []*string) *ExecuteProcessInstanceRequestFile {
	s.Photos = v
	return s
}

type ExecuteProcessInstanceRequestFileAttachments struct {
	// example:
	//
	// B1oQixxxx
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// example:
	//
	// 文件名称。
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// 1024
	FileSize *string `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// example:
	//
	// file
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// 123
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s ExecuteProcessInstanceRequestFileAttachments) String() string {
	return tea.Prettify(s)
}

func (s ExecuteProcessInstanceRequestFileAttachments) GoString() string {
	return s.String()
}

func (s *ExecuteProcessInstanceRequestFileAttachments) SetFileId(v string) *ExecuteProcessInstanceRequestFileAttachments {
	s.FileId = &v
	return s
}

func (s *ExecuteProcessInstanceRequestFileAttachments) SetFileName(v string) *ExecuteProcessInstanceRequestFileAttachments {
	s.FileName = &v
	return s
}

func (s *ExecuteProcessInstanceRequestFileAttachments) SetFileSize(v string) *ExecuteProcessInstanceRequestFileAttachments {
	s.FileSize = &v
	return s
}

func (s *ExecuteProcessInstanceRequestFileAttachments) SetFileType(v string) *ExecuteProcessInstanceRequestFileAttachments {
	s.FileType = &v
	return s
}

func (s *ExecuteProcessInstanceRequestFileAttachments) SetSpaceId(v string) *ExecuteProcessInstanceRequestFileAttachments {
	s.SpaceId = &v
	return s
}

type ExecuteProcessInstanceResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ExecuteProcessInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteProcessInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteProcessInstanceResponseBody) SetResult(v bool) *ExecuteProcessInstanceResponseBody {
	s.Result = &v
	return s
}

func (s *ExecuteProcessInstanceResponseBody) SetSuccess(v bool) *ExecuteProcessInstanceResponseBody {
	s.Success = &v
	return s
}

type ExecuteProcessInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteProcessInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteProcessInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteProcessInstanceResponse) GoString() string {
	return s.String()
}

func (s *ExecuteProcessInstanceResponse) SetHeaders(v map[string]*string) *ExecuteProcessInstanceResponse {
	s.Headers = v
	return s
}

func (s *ExecuteProcessInstanceResponse) SetStatusCode(v int32) *ExecuteProcessInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteProcessInstanceResponse) SetBody(v *ExecuteProcessInstanceResponseBody) *ExecuteProcessInstanceResponse {
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
	// example:
	//
	// 用于员工差旅费用报销使用
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	FormComponents []*FormComponent `json:"formComponents,omitempty" xml:"formComponents,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 出差报销审批
	Name           *string                          `json:"name,omitempty" xml:"name,omitempty"`
	ProcessCode    *string                          `json:"processCode,omitempty" xml:"processCode,omitempty"`
	TemplateConfig *FormCreateRequestTemplateConfig `json:"templateConfig,omitempty" xml:"templateConfig,omitempty" type:"Struct"`
}

func (s FormCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequest) GoString() string {
	return s.String()
}

func (s *FormCreateRequest) SetDescription(v string) *FormCreateRequest {
	s.Description = &v
	return s
}

func (s *FormCreateRequest) SetFormComponents(v []*FormComponent) *FormCreateRequest {
	s.FormComponents = v
	return s
}

func (s *FormCreateRequest) SetName(v string) *FormCreateRequest {
	s.Name = &v
	return s
}

func (s *FormCreateRequest) SetProcessCode(v string) *FormCreateRequest {
	s.ProcessCode = &v
	return s
}

func (s *FormCreateRequest) SetTemplateConfig(v *FormCreateRequestTemplateConfig) *FormCreateRequest {
	s.TemplateConfig = v
	return s
}

type FormCreateRequestTemplateConfig struct {
	// example:
	//
	// abcd
	DirId *string `json:"dirId,omitempty" xml:"dirId,omitempty"`
	// example:
	//
	// true
	DisableDeleteProcess *bool `json:"disableDeleteProcess,omitempty" xml:"disableDeleteProcess,omitempty"`
	// example:
	//
	// true
	DisableFormEdit *bool `json:"disableFormEdit,omitempty" xml:"disableFormEdit,omitempty"`
	// example:
	//
	// true
	DisableHomepage *bool `json:"disableHomepage,omitempty" xml:"disableHomepage,omitempty"`
	// example:
	//
	// true
	DisableResubmit *bool `json:"disableResubmit,omitempty" xml:"disableResubmit,omitempty"`
	// example:
	//
	// true
	DisableStopProcessButton *bool `json:"disableStopProcessButton,omitempty" xml:"disableStopProcessButton,omitempty"`
	// example:
	//
	// true
	Hidden *bool `json:"hidden,omitempty" xml:"hidden,omitempty"`
	// example:
	//
	// efgh
	OriginDirId *string `json:"originDirId,omitempty" xml:"originDirId,omitempty"`
}

func (s FormCreateRequestTemplateConfig) String() string {
	return tea.Prettify(s)
}

func (s FormCreateRequestTemplateConfig) GoString() string {
	return s.String()
}

func (s *FormCreateRequestTemplateConfig) SetDirId(v string) *FormCreateRequestTemplateConfig {
	s.DirId = &v
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

func (s *FormCreateRequestTemplateConfig) SetDisableHomepage(v bool) *FormCreateRequestTemplateConfig {
	s.DisableHomepage = &v
	return s
}

func (s *FormCreateRequestTemplateConfig) SetDisableResubmit(v bool) *FormCreateRequestTemplateConfig {
	s.DisableResubmit = &v
	return s
}

func (s *FormCreateRequestTemplateConfig) SetDisableStopProcessButton(v bool) *FormCreateRequestTemplateConfig {
	s.DisableStopProcessButton = &v
	return s
}

func (s *FormCreateRequestTemplateConfig) SetHidden(v bool) *FormCreateRequestTemplateConfig {
	s.Hidden = &v
	return s
}

func (s *FormCreateRequestTemplateConfig) SetOriginDirId(v string) *FormCreateRequestTemplateConfig {
	s.OriginDirId = &v
	return s
}

type FormCreateResponseBody struct {
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
	// This parameter is required.
	//
	// example:
	//
	// PROC-abcdef-example
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FormCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *FormCreateResponse) SetStatusCode(v int32) *FormCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *FormCreateResponse) SetBody(v *FormCreateResponseBody) *FormCreateResponse {
	s.Body = v
	return s
}

type GetAttachmentSpaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetAttachmentSpaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetAttachmentSpaceHeaders) GoString() string {
	return s.String()
}

func (s *GetAttachmentSpaceHeaders) SetCommonHeaders(v map[string]*string) *GetAttachmentSpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAttachmentSpaceHeaders) SetXAcsDingtalkAccessToken(v string) *GetAttachmentSpaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetAttachmentSpaceRequest struct {
	// example:
	//
	// 8345000
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetAttachmentSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAttachmentSpaceRequest) GoString() string {
	return s.String()
}

func (s *GetAttachmentSpaceRequest) SetAgentId(v int64) *GetAttachmentSpaceRequest {
	s.AgentId = &v
	return s
}

func (s *GetAttachmentSpaceRequest) SetUserId(v string) *GetAttachmentSpaceRequest {
	s.UserId = &v
	return s
}

type GetAttachmentSpaceResponseBody struct {
	Result *GetAttachmentSpaceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetAttachmentSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAttachmentSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetAttachmentSpaceResponseBody) SetResult(v *GetAttachmentSpaceResponseBodyResult) *GetAttachmentSpaceResponseBody {
	s.Result = v
	return s
}

func (s *GetAttachmentSpaceResponseBody) SetSuccess(v bool) *GetAttachmentSpaceResponseBody {
	s.Success = &v
	return s
}

type GetAttachmentSpaceResponseBodyResult struct {
	// example:
	//
	// 3996960664
	SpaceId *int64 `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s GetAttachmentSpaceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetAttachmentSpaceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAttachmentSpaceResponseBodyResult) SetSpaceId(v int64) *GetAttachmentSpaceResponseBodyResult {
	s.SpaceId = &v
	return s
}

type GetAttachmentSpaceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAttachmentSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAttachmentSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAttachmentSpaceResponse) GoString() string {
	return s.String()
}

func (s *GetAttachmentSpaceResponse) SetHeaders(v map[string]*string) *GetAttachmentSpaceResponse {
	s.Headers = v
	return s
}

func (s *GetAttachmentSpaceResponse) SetStatusCode(v int32) *GetAttachmentSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAttachmentSpaceResponse) SetBody(v *GetAttachmentSpaceResponseBody) *GetAttachmentSpaceResponse {
	s.Body = v
	return s
}

type GetConditionFormComponentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetConditionFormComponentHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetConditionFormComponentHeaders) GoString() string {
	return s.String()
}

func (s *GetConditionFormComponentHeaders) SetCommonHeaders(v map[string]*string) *GetConditionFormComponentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetConditionFormComponentHeaders) SetXAcsDingtalkAccessToken(v string) *GetConditionFormComponentHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetConditionFormComponentRequest struct {
	// example:
	//
	// 10
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-xxx
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
}

func (s GetConditionFormComponentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConditionFormComponentRequest) GoString() string {
	return s.String()
}

func (s *GetConditionFormComponentRequest) SetAgentId(v int64) *GetConditionFormComponentRequest {
	s.AgentId = &v
	return s
}

func (s *GetConditionFormComponentRequest) SetProcessCode(v string) *GetConditionFormComponentRequest {
	s.ProcessCode = &v
	return s
}

type GetConditionFormComponentResponseBody struct {
	Result []*GetConditionFormComponentResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetConditionFormComponentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConditionFormComponentResponseBody) GoString() string {
	return s.String()
}

func (s *GetConditionFormComponentResponseBody) SetResult(v []*GetConditionFormComponentResponseBodyResult) *GetConditionFormComponentResponseBody {
	s.Result = v
	return s
}

type GetConditionFormComponentResponseBodyResult struct {
	// This parameter is required.
	//
	// example:
	//
	// TextField
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 输入框
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
}

func (s GetConditionFormComponentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetConditionFormComponentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetConditionFormComponentResponseBodyResult) SetId(v string) *GetConditionFormComponentResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetConditionFormComponentResponseBodyResult) SetLabel(v string) *GetConditionFormComponentResponseBodyResult {
	s.Label = &v
	return s
}

type GetConditionFormComponentResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConditionFormComponentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConditionFormComponentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConditionFormComponentResponse) GoString() string {
	return s.String()
}

func (s *GetConditionFormComponentResponse) SetHeaders(v map[string]*string) *GetConditionFormComponentResponse {
	s.Headers = v
	return s
}

func (s *GetConditionFormComponentResponse) SetStatusCode(v int32) *GetConditionFormComponentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConditionFormComponentResponse) SetBody(v *GetConditionFormComponentResponseBody) *GetConditionFormComponentResponse {
	s.Body = v
	return s
}

type GetCrmProcCodesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetCrmProcCodesHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetCrmProcCodesHeaders) GoString() string {
	return s.String()
}

func (s *GetCrmProcCodesHeaders) SetCommonHeaders(v map[string]*string) *GetCrmProcCodesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCrmProcCodesHeaders) SetXAcsDingtalkAccessToken(v string) *GetCrmProcCodesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetCrmProcCodesResponseBody struct {
	Result []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetCrmProcCodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCrmProcCodesResponseBody) GoString() string {
	return s.String()
}

func (s *GetCrmProcCodesResponseBody) SetResult(v []*string) *GetCrmProcCodesResponseBody {
	s.Result = v
	return s
}

type GetCrmProcCodesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCrmProcCodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCrmProcCodesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCrmProcCodesResponse) GoString() string {
	return s.String()
}

func (s *GetCrmProcCodesResponse) SetHeaders(v map[string]*string) *GetCrmProcCodesResponse {
	s.Headers = v
	return s
}

func (s *GetCrmProcCodesResponse) SetStatusCode(v int32) *GetCrmProcCodesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCrmProcCodesResponse) SetBody(v *GetCrmProcCodesResponseBody) *GetCrmProcCodesResponse {
	s.Body = v
	return s
}

type GetFieldModifiedHistoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetFieldModifiedHistoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetFieldModifiedHistoryHeaders) GoString() string {
	return s.String()
}

func (s *GetFieldModifiedHistoryHeaders) SetCommonHeaders(v map[string]*string) *GetFieldModifiedHistoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFieldModifiedHistoryHeaders) SetXAcsDingtalkAccessToken(v string) *GetFieldModifiedHistoryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetFieldModifiedHistoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// TextField-abcd
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// proc-FF6Y2xxxx
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
}

func (s GetFieldModifiedHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFieldModifiedHistoryRequest) GoString() string {
	return s.String()
}

func (s *GetFieldModifiedHistoryRequest) SetFieldId(v string) *GetFieldModifiedHistoryRequest {
	s.FieldId = &v
	return s
}

func (s *GetFieldModifiedHistoryRequest) SetProcessInstanceId(v string) *GetFieldModifiedHistoryRequest {
	s.ProcessInstanceId = &v
	return s
}

type GetFieldModifiedHistoryResponseBody struct {
	Result []*GetFieldModifiedHistoryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetFieldModifiedHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFieldModifiedHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetFieldModifiedHistoryResponseBody) SetResult(v []*GetFieldModifiedHistoryResponseBodyResult) *GetFieldModifiedHistoryResponseBody {
	s.Result = v
	return s
}

func (s *GetFieldModifiedHistoryResponseBody) SetSuccess(v bool) *GetFieldModifiedHistoryResponseBody {
	s.Success = &v
	return s
}

type GetFieldModifiedHistoryResponseBodyResult struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2024-04-02T11:52Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// TextField-abcd
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// example:
	//
	// 钉钉1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// userId1
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// 从 111 修改到 222
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetFieldModifiedHistoryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFieldModifiedHistoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFieldModifiedHistoryResponseBodyResult) SetCreateTime(v string) *GetFieldModifiedHistoryResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetFieldModifiedHistoryResponseBodyResult) SetFieldId(v string) *GetFieldModifiedHistoryResponseBodyResult {
	s.FieldId = &v
	return s
}

func (s *GetFieldModifiedHistoryResponseBodyResult) SetName(v string) *GetFieldModifiedHistoryResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetFieldModifiedHistoryResponseBodyResult) SetUserId(v string) *GetFieldModifiedHistoryResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *GetFieldModifiedHistoryResponseBodyResult) SetValue(v string) *GetFieldModifiedHistoryResponseBodyResult {
	s.Value = &v
	return s
}

type GetFieldModifiedHistoryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFieldModifiedHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFieldModifiedHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFieldModifiedHistoryResponse) GoString() string {
	return s.String()
}

func (s *GetFieldModifiedHistoryResponse) SetHeaders(v map[string]*string) *GetFieldModifiedHistoryResponse {
	s.Headers = v
	return s
}

func (s *GetFieldModifiedHistoryResponse) SetStatusCode(v int32) *GetFieldModifiedHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFieldModifiedHistoryResponse) SetBody(v *GetFieldModifiedHistoryResponseBody) *GetFieldModifiedHistoryResponse {
	s.Body = v
	return s
}

type GetHandSignDownloadUrlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetHandSignDownloadUrlHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetHandSignDownloadUrlHeaders) GoString() string {
	return s.String()
}

func (s *GetHandSignDownloadUrlHeaders) SetCommonHeaders(v map[string]*string) *GetHandSignDownloadUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHandSignDownloadUrlHeaders) SetXAcsDingtalkAccessToken(v string) *GetHandSignDownloadUrlHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetHandSignDownloadUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// AzBltRlvKukX3Wxxxx
	HandSignToken *string `json:"handSignToken,omitempty" xml:"handSignToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ag187wewxxxx
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
}

func (s GetHandSignDownloadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHandSignDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetHandSignDownloadUrlRequest) SetHandSignToken(v string) *GetHandSignDownloadUrlRequest {
	s.HandSignToken = &v
	return s
}

func (s *GetHandSignDownloadUrlRequest) SetProcessInstanceId(v string) *GetHandSignDownloadUrlRequest {
	s.ProcessInstanceId = &v
	return s
}

type GetHandSignDownloadUrlResponseBody struct {
	Result  *GetHandSignDownloadUrlResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *string                                   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetHandSignDownloadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHandSignDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetHandSignDownloadUrlResponseBody) SetResult(v *GetHandSignDownloadUrlResponseBodyResult) *GetHandSignDownloadUrlResponseBody {
	s.Result = v
	return s
}

func (s *GetHandSignDownloadUrlResponseBody) SetSuccess(v string) *GetHandSignDownloadUrlResponseBody {
	s.Success = &v
	return s
}

type GetHandSignDownloadUrlResponseBodyResult struct {
	// example:
	//
	// https://dingding-file-zjk.oss-cn-zhangjiakouxxxx
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	ExpireIn    *int64  `json:"expireIn,omitempty" xml:"expireIn,omitempty"`
}

func (s GetHandSignDownloadUrlResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetHandSignDownloadUrlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetHandSignDownloadUrlResponseBodyResult) SetDownloadUrl(v string) *GetHandSignDownloadUrlResponseBodyResult {
	s.DownloadUrl = &v
	return s
}

func (s *GetHandSignDownloadUrlResponseBodyResult) SetExpireIn(v int64) *GetHandSignDownloadUrlResponseBodyResult {
	s.ExpireIn = &v
	return s
}

type GetHandSignDownloadUrlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHandSignDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHandSignDownloadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHandSignDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetHandSignDownloadUrlResponse) SetHeaders(v map[string]*string) *GetHandSignDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetHandSignDownloadUrlResponse) SetStatusCode(v int32) *GetHandSignDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHandSignDownloadUrlResponse) SetBody(v *GetHandSignDownloadUrlResponseBody) *GetHandSignDownloadUrlResponse {
	s.Body = v
	return s
}

type GetManageProcessByStaffIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetManageProcessByStaffIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetManageProcessByStaffIdHeaders) GoString() string {
	return s.String()
}

func (s *GetManageProcessByStaffIdHeaders) SetCommonHeaders(v map[string]*string) *GetManageProcessByStaffIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetManageProcessByStaffIdHeaders) SetXAcsDingtalkAccessToken(v string) *GetManageProcessByStaffIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetManageProcessByStaffIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// manager7078
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetManageProcessByStaffIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetManageProcessByStaffIdRequest) GoString() string {
	return s.String()
}

func (s *GetManageProcessByStaffIdRequest) SetUserId(v string) *GetManageProcessByStaffIdRequest {
	s.UserId = &v
	return s
}

type GetManageProcessByStaffIdResponseBody struct {
	Result []*GetManageProcessByStaffIdResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetManageProcessByStaffIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetManageProcessByStaffIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetManageProcessByStaffIdResponseBody) SetResult(v []*GetManageProcessByStaffIdResponseBodyResult) *GetManageProcessByStaffIdResponseBody {
	s.Result = v
	return s
}

func (s *GetManageProcessByStaffIdResponseBody) SetSuccess(v bool) *GetManageProcessByStaffIdResponseBody {
	s.Success = &v
	return s
}

type GetManageProcessByStaffIdResponseBodyResult struct {
	// example:
	//
	// 0
	AttendanceType *int32 `json:"attendanceType,omitempty" xml:"attendanceType,omitempty"`
	// example:
	//
	// 通用审批
	FlowTitle *string `json:"flowTitle,omitempty" xml:"flowTitle,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2020-07-14 14:24:59
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// common
	IconName *string `json:"iconName,omitempty" xml:"iconName,omitempty"`
	// example:
	//
	// https://gw.alicdn.com/tfs/xxxx-112-112.png
	IconUrl *string `json:"iconUrl,omitempty" xml:"iconUrl,omitempty"`
	// example:
	//
	// true
	NewProcess *bool `json:"newProcess,omitempty" xml:"newProcess,omitempty"`
	// example:
	//
	// PROC-44E84FC1-16E2-4A69-BB3C-xxxx
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
}

func (s GetManageProcessByStaffIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetManageProcessByStaffIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetManageProcessByStaffIdResponseBodyResult) SetAttendanceType(v int32) *GetManageProcessByStaffIdResponseBodyResult {
	s.AttendanceType = &v
	return s
}

func (s *GetManageProcessByStaffIdResponseBodyResult) SetFlowTitle(v string) *GetManageProcessByStaffIdResponseBodyResult {
	s.FlowTitle = &v
	return s
}

func (s *GetManageProcessByStaffIdResponseBodyResult) SetGmtModified(v string) *GetManageProcessByStaffIdResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *GetManageProcessByStaffIdResponseBodyResult) SetIconName(v string) *GetManageProcessByStaffIdResponseBodyResult {
	s.IconName = &v
	return s
}

func (s *GetManageProcessByStaffIdResponseBodyResult) SetIconUrl(v string) *GetManageProcessByStaffIdResponseBodyResult {
	s.IconUrl = &v
	return s
}

func (s *GetManageProcessByStaffIdResponseBodyResult) SetNewProcess(v bool) *GetManageProcessByStaffIdResponseBodyResult {
	s.NewProcess = &v
	return s
}

func (s *GetManageProcessByStaffIdResponseBodyResult) SetProcessCode(v string) *GetManageProcessByStaffIdResponseBodyResult {
	s.ProcessCode = &v
	return s
}

type GetManageProcessByStaffIdResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetManageProcessByStaffIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetManageProcessByStaffIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetManageProcessByStaffIdResponse) GoString() string {
	return s.String()
}

func (s *GetManageProcessByStaffIdResponse) SetHeaders(v map[string]*string) *GetManageProcessByStaffIdResponse {
	s.Headers = v
	return s
}

func (s *GetManageProcessByStaffIdResponse) SetStatusCode(v int32) *GetManageProcessByStaffIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetManageProcessByStaffIdResponse) SetBody(v *GetManageProcessByStaffIdResponseBody) *GetManageProcessByStaffIdResponse {
	s.Body = v
	return s
}

type GetProcessCodeByNameHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetProcessCodeByNameHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetProcessCodeByNameHeaders) GoString() string {
	return s.String()
}

func (s *GetProcessCodeByNameHeaders) SetCommonHeaders(v map[string]*string) *GetProcessCodeByNameHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetProcessCodeByNameHeaders) SetXAcsDingtalkAccessToken(v string) *GetProcessCodeByNameHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetProcessCodeByNameRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// abc
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetProcessCodeByNameRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProcessCodeByNameRequest) GoString() string {
	return s.String()
}

func (s *GetProcessCodeByNameRequest) SetName(v string) *GetProcessCodeByNameRequest {
	s.Name = &v
	return s
}

type GetProcessCodeByNameResponseBody struct {
	Result *GetProcessCodeByNameResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetProcessCodeByNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProcessCodeByNameResponseBody) GoString() string {
	return s.String()
}

func (s *GetProcessCodeByNameResponseBody) SetResult(v *GetProcessCodeByNameResponseBodyResult) *GetProcessCodeByNameResponseBody {
	s.Result = v
	return s
}

type GetProcessCodeByNameResponseBodyResult struct {
	// This parameter is required.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2024-03-22T11:50Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-abcdef-example
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
}

func (s GetProcessCodeByNameResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetProcessCodeByNameResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetProcessCodeByNameResponseBodyResult) SetGmtModified(v string) *GetProcessCodeByNameResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *GetProcessCodeByNameResponseBodyResult) SetProcessCode(v string) *GetProcessCodeByNameResponseBodyResult {
	s.ProcessCode = &v
	return s
}

type GetProcessCodeByNameResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProcessCodeByNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProcessCodeByNameResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProcessCodeByNameResponse) GoString() string {
	return s.String()
}

func (s *GetProcessCodeByNameResponse) SetHeaders(v map[string]*string) *GetProcessCodeByNameResponse {
	s.Headers = v
	return s
}

func (s *GetProcessCodeByNameResponse) SetStatusCode(v int32) *GetProcessCodeByNameResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProcessCodeByNameResponse) SetBody(v *GetProcessCodeByNameResponseBody) *GetProcessCodeByNameResponse {
	s.Body = v
	return s
}

type GetProcessConfigHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetProcessConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetProcessConfigHeaders) GoString() string {
	return s.String()
}

func (s *GetProcessConfigHeaders) SetCommonHeaders(v map[string]*string) *GetProcessConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetProcessConfigHeaders) SetXAcsDingtalkAccessToken(v string) *GetProcessConfigHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetProcessConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// PROC-BEFC22B7-EA64-4336-86EB-AB773AA2EB12
	ProcCode *string `json:"procCode,omitempty" xml:"procCode,omitempty"`
}

func (s GetProcessConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProcessConfigRequest) GoString() string {
	return s.String()
}

func (s *GetProcessConfigRequest) SetProcCode(v string) *GetProcessConfigRequest {
	s.ProcCode = &v
	return s
}

type GetProcessConfigResponseBody struct {
	Result *GetProcessConfigResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetProcessConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProcessConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetProcessConfigResponseBody) SetResult(v *GetProcessConfigResponseBodyResult) *GetProcessConfigResponseBody {
	s.Result = v
	return s
}

type GetProcessConfigResponseBodyResult struct {
	AbstractGenRule []*string `json:"abstractGenRule,omitempty" xml:"abstractGenRule,omitempty" type:"Repeated"`
	// example:
	//
	// {"sid_instStart":[{"fieldId":"TextField-K2AD4O5B","fieldBehavior":"HIDDEN","componentName":"TextField","disableBehaviors":[]}],"1918_5cd3":[{"fieldId":"TextField-K2AD4O5B","fieldBehavior":"HIDDEN","componentName":"TextField","disableBehaviors":[]}],"d01c_a677":[{"fieldId":"TextField-K2AD4O5B","fieldBehavior":"NORMAL","componentName":"TextField","disableBehaviors":[]}]}
	ActivityAuth               *string `json:"activityAuth,omitempty" xml:"activityAuth,omitempty"`
	AllowRevoke                *bool   `json:"allowRevoke,omitempty" xml:"allowRevoke,omitempty"`
	AppendEnable               *bool   `json:"appendEnable,omitempty" xml:"appendEnable,omitempty"`
	AutoExecuteOriginatorTasks *bool   `json:"autoExecuteOriginatorTasks,omitempty" xml:"autoExecuteOriginatorTasks,omitempty"`
	// example:
	//
	// alitrip.business
	BizCategoryId *string `json:"bizCategoryId,omitempty" xml:"bizCategoryId,omitempty"`
	// example:
	//
	// crm_customer
	BizType     *string                                        `json:"bizType,omitempty" xml:"bizType,omitempty"`
	CommentConf *GetProcessConfigResponseBodyResultCommentConf `json:"commentConf,omitempty" xml:"commentConf,omitempty" type:"Struct"`
	// example:
	//
	// continuousFirst
	DuplicateRemoval *string `json:"duplicateRemoval,omitempty" xml:"duplicateRemoval,omitempty"`
	// example:
	//
	// {"items":[]}
	FormSchema   *string                                         `json:"formSchema,omitempty" xml:"formSchema,omitempty"`
	HandSignConf *GetProcessConfigResponseBodyResultHandSignConf `json:"handSignConf,omitempty" xml:"handSignConf,omitempty" type:"Struct"`
	Managers     []*string                                       `json:"managers,omitempty" xml:"managers,omitempty" type:"Repeated"`
	// example:
	//
	// 模板名称
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	ProcessAppType *bool   `json:"processAppType,omitempty" xml:"processAppType,omitempty"`
	// example:
	//
	// {"type":"","properties":{},"childNode":{}}
	ProcessConfig        *string                                                 `json:"processConfig,omitempty" xml:"processConfig,omitempty"`
	StaticProc           *bool                                                   `json:"staticProc,omitempty" xml:"staticProc,omitempty"`
	SubstituteSubmitConf *GetProcessConfigResponseBodyResultSubstituteSubmitConf `json:"substituteSubmitConf,omitempty" xml:"substituteSubmitConf,omitempty" type:"Struct"`
	TitleGenRule         *GetProcessConfigResponseBodyResultTitleGenRule         `json:"titleGenRule,omitempty" xml:"titleGenRule,omitempty" type:"Struct"`
	Visibility           []*GetProcessConfigResponseBodyResultVisibility         `json:"visibility,omitempty" xml:"visibility,omitempty" type:"Repeated"`
}

func (s GetProcessConfigResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetProcessConfigResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetProcessConfigResponseBodyResult) SetAbstractGenRule(v []*string) *GetProcessConfigResponseBodyResult {
	s.AbstractGenRule = v
	return s
}

func (s *GetProcessConfigResponseBodyResult) SetActivityAuth(v string) *GetProcessConfigResponseBodyResult {
	s.ActivityAuth = &v
	return s
}

func (s *GetProcessConfigResponseBodyResult) SetAllowRevoke(v bool) *GetProcessConfigResponseBodyResult {
	s.AllowRevoke = &v
	return s
}

func (s *GetProcessConfigResponseBodyResult) SetAppendEnable(v bool) *GetProcessConfigResponseBodyResult {
	s.AppendEnable = &v
	return s
}

func (s *GetProcessConfigResponseBodyResult) SetAutoExecuteOriginatorTasks(v bool) *GetProcessConfigResponseBodyResult {
	s.AutoExecuteOriginatorTasks = &v
	return s
}

func (s *GetProcessConfigResponseBodyResult) SetBizCategoryId(v string) *GetProcessConfigResponseBodyResult {
	s.BizCategoryId = &v
	return s
}

func (s *GetProcessConfigResponseBodyResult) SetBizType(v string) *GetProcessConfigResponseBodyResult {
	s.BizType = &v
	return s
}

func (s *GetProcessConfigResponseBodyResult) SetCommentConf(v *GetProcessConfigResponseBodyResultCommentConf) *GetProcessConfigResponseBodyResult {
	s.CommentConf = v
	return s
}

func (s *GetProcessConfigResponseBodyResult) SetDuplicateRemoval(v string) *GetProcessConfigResponseBodyResult {
	s.DuplicateRemoval = &v
	return s
}

func (s *GetProcessConfigResponseBodyResult) SetFormSchema(v string) *GetProcessConfigResponseBodyResult {
	s.FormSchema = &v
	return s
}

func (s *GetProcessConfigResponseBodyResult) SetHandSignConf(v *GetProcessConfigResponseBodyResultHandSignConf) *GetProcessConfigResponseBodyResult {
	s.HandSignConf = v
	return s
}

func (s *GetProcessConfigResponseBodyResult) SetManagers(v []*string) *GetProcessConfigResponseBodyResult {
	s.Managers = v
	return s
}

func (s *GetProcessConfigResponseBodyResult) SetName(v string) *GetProcessConfigResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetProcessConfigResponseBodyResult) SetProcessAppType(v bool) *GetProcessConfigResponseBodyResult {
	s.ProcessAppType = &v
	return s
}

func (s *GetProcessConfigResponseBodyResult) SetProcessConfig(v string) *GetProcessConfigResponseBodyResult {
	s.ProcessConfig = &v
	return s
}

func (s *GetProcessConfigResponseBodyResult) SetStaticProc(v bool) *GetProcessConfigResponseBodyResult {
	s.StaticProc = &v
	return s
}

func (s *GetProcessConfigResponseBodyResult) SetSubstituteSubmitConf(v *GetProcessConfigResponseBodyResultSubstituteSubmitConf) *GetProcessConfigResponseBodyResult {
	s.SubstituteSubmitConf = v
	return s
}

func (s *GetProcessConfigResponseBodyResult) SetTitleGenRule(v *GetProcessConfigResponseBodyResultTitleGenRule) *GetProcessConfigResponseBodyResult {
	s.TitleGenRule = v
	return s
}

func (s *GetProcessConfigResponseBodyResult) SetVisibility(v []*GetProcessConfigResponseBodyResultVisibility) *GetProcessConfigResponseBodyResult {
	s.Visibility = v
	return s
}

type GetProcessConfigResponseBodyResultCommentConf struct {
	// example:
	//
	// 评论描述
	CommentDescription       *string `json:"commentDescription,omitempty" xml:"commentDescription,omitempty"`
	CommentHiddenForProposer *bool   `json:"commentHiddenForProposer,omitempty" xml:"commentHiddenForProposer,omitempty"`
	CommentRequired          *bool   `json:"commentRequired,omitempty" xml:"commentRequired,omitempty"`
}

func (s GetProcessConfigResponseBodyResultCommentConf) String() string {
	return tea.Prettify(s)
}

func (s GetProcessConfigResponseBodyResultCommentConf) GoString() string {
	return s.String()
}

func (s *GetProcessConfigResponseBodyResultCommentConf) SetCommentDescription(v string) *GetProcessConfigResponseBodyResultCommentConf {
	s.CommentDescription = &v
	return s
}

func (s *GetProcessConfigResponseBodyResultCommentConf) SetCommentHiddenForProposer(v bool) *GetProcessConfigResponseBodyResultCommentConf {
	s.CommentHiddenForProposer = &v
	return s
}

func (s *GetProcessConfigResponseBodyResultCommentConf) SetCommentRequired(v bool) *GetProcessConfigResponseBodyResultCommentConf {
	s.CommentRequired = &v
	return s
}

type GetProcessConfigResponseBodyResultHandSignConf struct {
	HandSignEnable *bool `json:"handSignEnable,omitempty" xml:"handSignEnable,omitempty"`
	ResignEnable   *bool `json:"resignEnable,omitempty" xml:"resignEnable,omitempty"`
}

func (s GetProcessConfigResponseBodyResultHandSignConf) String() string {
	return tea.Prettify(s)
}

func (s GetProcessConfigResponseBodyResultHandSignConf) GoString() string {
	return s.String()
}

func (s *GetProcessConfigResponseBodyResultHandSignConf) SetHandSignEnable(v bool) *GetProcessConfigResponseBodyResultHandSignConf {
	s.HandSignEnable = &v
	return s
}

func (s *GetProcessConfigResponseBodyResultHandSignConf) SetResignEnable(v bool) *GetProcessConfigResponseBodyResultHandSignConf {
	s.ResignEnable = &v
	return s
}

type GetProcessConfigResponseBodyResultSubstituteSubmitConf struct {
	Enable        *bool                                                                  `json:"enable,omitempty" xml:"enable,omitempty"`
	SubmitterList []*GetProcessConfigResponseBodyResultSubstituteSubmitConfSubmitterList `json:"submitterList,omitempty" xml:"submitterList,omitempty" type:"Repeated"`
}

func (s GetProcessConfigResponseBodyResultSubstituteSubmitConf) String() string {
	return tea.Prettify(s)
}

func (s GetProcessConfigResponseBodyResultSubstituteSubmitConf) GoString() string {
	return s.String()
}

func (s *GetProcessConfigResponseBodyResultSubstituteSubmitConf) SetEnable(v bool) *GetProcessConfigResponseBodyResultSubstituteSubmitConf {
	s.Enable = &v
	return s
}

func (s *GetProcessConfigResponseBodyResultSubstituteSubmitConf) SetSubmitterList(v []*GetProcessConfigResponseBodyResultSubstituteSubmitConfSubmitterList) *GetProcessConfigResponseBodyResultSubstituteSubmitConf {
	s.SubmitterList = v
	return s
}

type GetProcessConfigResponseBodyResultSubstituteSubmitConfSubmitterList struct {
	// example:
	//
	// 钉三多
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// approval
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// manager1234
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetProcessConfigResponseBodyResultSubstituteSubmitConfSubmitterList) String() string {
	return tea.Prettify(s)
}

func (s GetProcessConfigResponseBodyResultSubstituteSubmitConfSubmitterList) GoString() string {
	return s.String()
}

func (s *GetProcessConfigResponseBodyResultSubstituteSubmitConfSubmitterList) SetName(v string) *GetProcessConfigResponseBodyResultSubstituteSubmitConfSubmitterList {
	s.Name = &v
	return s
}

func (s *GetProcessConfigResponseBodyResultSubstituteSubmitConfSubmitterList) SetType(v string) *GetProcessConfigResponseBodyResultSubstituteSubmitConfSubmitterList {
	s.Type = &v
	return s
}

func (s *GetProcessConfigResponseBodyResultSubstituteSubmitConfSubmitterList) SetValue(v string) *GetProcessConfigResponseBodyResultSubstituteSubmitConfSubmitterList {
	s.Value = &v
	return s
}

type GetProcessConfigResponseBodyResultTitleGenRule struct {
	// example:
	//
	// #{originator}#{formName}#{createTime}
	Express *string `json:"express,omitempty" xml:"express,omitempty"`
	// example:
	//
	// 2
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetProcessConfigResponseBodyResultTitleGenRule) String() string {
	return tea.Prettify(s)
}

func (s GetProcessConfigResponseBodyResultTitleGenRule) GoString() string {
	return s.String()
}

func (s *GetProcessConfigResponseBodyResultTitleGenRule) SetExpress(v string) *GetProcessConfigResponseBodyResultTitleGenRule {
	s.Express = &v
	return s
}

func (s *GetProcessConfigResponseBodyResultTitleGenRule) SetType(v int32) *GetProcessConfigResponseBodyResultTitleGenRule {
	s.Type = &v
	return s
}

type GetProcessConfigResponseBodyResultVisibility struct {
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// manager345
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetProcessConfigResponseBodyResultVisibility) String() string {
	return tea.Prettify(s)
}

func (s GetProcessConfigResponseBodyResultVisibility) GoString() string {
	return s.String()
}

func (s *GetProcessConfigResponseBodyResultVisibility) SetType(v int32) *GetProcessConfigResponseBodyResultVisibility {
	s.Type = &v
	return s
}

func (s *GetProcessConfigResponseBodyResultVisibility) SetValue(v string) *GetProcessConfigResponseBodyResultVisibility {
	s.Value = &v
	return s
}

type GetProcessConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProcessConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProcessConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProcessConfigResponse) GoString() string {
	return s.String()
}

func (s *GetProcessConfigResponse) SetHeaders(v map[string]*string) *GetProcessConfigResponse {
	s.Headers = v
	return s
}

func (s *GetProcessConfigResponse) SetStatusCode(v int32) *GetProcessConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProcessConfigResponse) SetBody(v *GetProcessConfigResponseBody) *GetProcessConfigResponse {
	s.Body = v
	return s
}

type GetProcessInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetProcessInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetProcessInstanceHeaders) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceHeaders) SetCommonHeaders(v map[string]*string) *GetProcessInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetProcessInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *GetProcessInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetProcessInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a171de6c-8bxxxx
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
}

func (s GetProcessInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProcessInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceRequest) SetProcessInstanceId(v string) *GetProcessInstanceRequest {
	s.ProcessInstanceId = &v
	return s
}

type GetProcessInstanceResponseBody struct {
	Result *GetProcessInstanceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetProcessInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProcessInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceResponseBody) SetResult(v *GetProcessInstanceResponseBodyResult) *GetProcessInstanceResponseBody {
	s.Result = v
	return s
}

func (s *GetProcessInstanceResponseBody) SetSuccess(v string) *GetProcessInstanceResponseBody {
	s.Success = &v
	return s
}

type GetProcessInstanceResponseBodyResult struct {
	ApproverUserIds []*string `json:"approverUserIds,omitempty" xml:"approverUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// ["instance1","instance2"]
	AttachedProcessInstanceIds []*string `json:"attachedProcessInstanceIds,omitempty" xml:"attachedProcessInstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// MODIFY
	BizAction *string `json:"bizAction,omitempty" xml:"bizAction,omitempty"`
	// example:
	//
	// {"mykey": "myData"}
	BizData *string `json:"bizData,omitempty" xml:"bizData,omitempty"`
	// example:
	//
	// 111
	BusinessId *string   `json:"businessId,omitempty" xml:"businessId,omitempty"`
	CcUserIds  []*string `json:"ccUserIds,omitempty" xml:"ccUserIds,omitempty" type:"Repeated"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-08-31T11:52Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 2022-08-31T11:52Z
	FinishTime          *string                                                    `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	FormComponentValues []*GetProcessInstanceResponseBodyResultFormComponentValues `json:"formComponentValues,omitempty" xml:"formComponentValues,omitempty" type:"Repeated"`
	// example:
	//
	// AG3U12xWRFex63h6bCPUWw10221698052827
	MainProcessInstanceId *string                                                 `json:"mainProcessInstanceId,omitempty" xml:"mainProcessInstanceId,omitempty"`
	OperationRecords      []*GetProcessInstanceResponseBodyResultOperationRecords `json:"operationRecords,omitempty" xml:"operationRecords,omitempty" type:"Repeated"`
	// example:
	//
	// -1
	OriginatorDeptId *string `json:"originatorDeptId,omitempty" xml:"originatorDeptId,omitempty"`
	// example:
	//
	// 测试
	OriginatorDeptName *string `json:"originatorDeptName,omitempty" xml:"originatorDeptName,omitempty"`
	// example:
	//
	// manager1
	OriginatorUserId *string `json:"originatorUserId,omitempty" xml:"originatorUserId,omitempty"`
	// example:
	//
	// agree
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// NEW
	Status *string                                      `json:"status,omitempty" xml:"status,omitempty"`
	Tasks  []*GetProcessInstanceResponseBodyResultTasks `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Repeated"`
	// example:
	//
	// xx提交的请假申请
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetProcessInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetProcessInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceResponseBodyResult) SetApproverUserIds(v []*string) *GetProcessInstanceResponseBodyResult {
	s.ApproverUserIds = v
	return s
}

func (s *GetProcessInstanceResponseBodyResult) SetAttachedProcessInstanceIds(v []*string) *GetProcessInstanceResponseBodyResult {
	s.AttachedProcessInstanceIds = v
	return s
}

func (s *GetProcessInstanceResponseBodyResult) SetBizAction(v string) *GetProcessInstanceResponseBodyResult {
	s.BizAction = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResult) SetBizData(v string) *GetProcessInstanceResponseBodyResult {
	s.BizData = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResult) SetBusinessId(v string) *GetProcessInstanceResponseBodyResult {
	s.BusinessId = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResult) SetCcUserIds(v []*string) *GetProcessInstanceResponseBodyResult {
	s.CcUserIds = v
	return s
}

func (s *GetProcessInstanceResponseBodyResult) SetCreateTime(v string) *GetProcessInstanceResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResult) SetFinishTime(v string) *GetProcessInstanceResponseBodyResult {
	s.FinishTime = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResult) SetFormComponentValues(v []*GetProcessInstanceResponseBodyResultFormComponentValues) *GetProcessInstanceResponseBodyResult {
	s.FormComponentValues = v
	return s
}

func (s *GetProcessInstanceResponseBodyResult) SetMainProcessInstanceId(v string) *GetProcessInstanceResponseBodyResult {
	s.MainProcessInstanceId = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResult) SetOperationRecords(v []*GetProcessInstanceResponseBodyResultOperationRecords) *GetProcessInstanceResponseBodyResult {
	s.OperationRecords = v
	return s
}

func (s *GetProcessInstanceResponseBodyResult) SetOriginatorDeptId(v string) *GetProcessInstanceResponseBodyResult {
	s.OriginatorDeptId = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResult) SetOriginatorDeptName(v string) *GetProcessInstanceResponseBodyResult {
	s.OriginatorDeptName = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResult) SetOriginatorUserId(v string) *GetProcessInstanceResponseBodyResult {
	s.OriginatorUserId = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResult) SetResult(v string) *GetProcessInstanceResponseBodyResult {
	s.Result = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResult) SetStatus(v string) *GetProcessInstanceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResult) SetTasks(v []*GetProcessInstanceResponseBodyResultTasks) *GetProcessInstanceResponseBodyResult {
	s.Tasks = v
	return s
}

func (s *GetProcessInstanceResponseBodyResult) SetTitle(v string) *GetProcessInstanceResponseBodyResult {
	s.Title = &v
	return s
}

type GetProcessInstanceResponseBodyResultFormComponentValues struct {
	// example:
	//
	// TextField-bizAlias
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// example:
	//
	// DDSelectField
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// example:
	//
	// 示例值
	ExtValue *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	// example:
	//
	// DDHolidayField-J2Bxxxx
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 组件1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 示例值
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetProcessInstanceResponseBodyResultFormComponentValues) String() string {
	return tea.Prettify(s)
}

func (s GetProcessInstanceResponseBodyResultFormComponentValues) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceResponseBodyResultFormComponentValues) SetBizAlias(v string) *GetProcessInstanceResponseBodyResultFormComponentValues {
	s.BizAlias = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResultFormComponentValues) SetComponentType(v string) *GetProcessInstanceResponseBodyResultFormComponentValues {
	s.ComponentType = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResultFormComponentValues) SetExtValue(v string) *GetProcessInstanceResponseBodyResultFormComponentValues {
	s.ExtValue = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResultFormComponentValues) SetId(v string) *GetProcessInstanceResponseBodyResultFormComponentValues {
	s.Id = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResultFormComponentValues) SetName(v string) *GetProcessInstanceResponseBodyResultFormComponentValues {
	s.Name = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResultFormComponentValues) SetValue(v string) *GetProcessInstanceResponseBodyResultFormComponentValues {
	s.Value = &v
	return s
}

type GetProcessInstanceResponseBodyResultOperationRecords struct {
	ActivityId  *string                                                            `json:"activityId,omitempty" xml:"activityId,omitempty"`
	Attachments []*GetProcessInstanceResponseBodyResultOperationRecordsAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	CcUserIds   []*string                                                          `json:"ccUserIds,omitempty" xml:"ccUserIds,omitempty" type:"Repeated"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-08-31T11:52Z
	Date   *string   `json:"date,omitempty" xml:"date,omitempty"`
	Images []*string `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	// example:
	//
	// 评论
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// AGREE
	Result   *string `json:"result,omitempty" xml:"result,omitempty"`
	ShowName *string `json:"showName,omitempty" xml:"showName,omitempty"`
	// example:
	//
	// EXECUTE_TASK_NORMAL
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// manager1
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetProcessInstanceResponseBodyResultOperationRecords) String() string {
	return tea.Prettify(s)
}

func (s GetProcessInstanceResponseBodyResultOperationRecords) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceResponseBodyResultOperationRecords) SetActivityId(v string) *GetProcessInstanceResponseBodyResultOperationRecords {
	s.ActivityId = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResultOperationRecords) SetAttachments(v []*GetProcessInstanceResponseBodyResultOperationRecordsAttachments) *GetProcessInstanceResponseBodyResultOperationRecords {
	s.Attachments = v
	return s
}

func (s *GetProcessInstanceResponseBodyResultOperationRecords) SetCcUserIds(v []*string) *GetProcessInstanceResponseBodyResultOperationRecords {
	s.CcUserIds = v
	return s
}

func (s *GetProcessInstanceResponseBodyResultOperationRecords) SetDate(v string) *GetProcessInstanceResponseBodyResultOperationRecords {
	s.Date = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResultOperationRecords) SetImages(v []*string) *GetProcessInstanceResponseBodyResultOperationRecords {
	s.Images = v
	return s
}

func (s *GetProcessInstanceResponseBodyResultOperationRecords) SetRemark(v string) *GetProcessInstanceResponseBodyResultOperationRecords {
	s.Remark = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResultOperationRecords) SetResult(v string) *GetProcessInstanceResponseBodyResultOperationRecords {
	s.Result = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResultOperationRecords) SetShowName(v string) *GetProcessInstanceResponseBodyResultOperationRecords {
	s.ShowName = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResultOperationRecords) SetType(v string) *GetProcessInstanceResponseBodyResultOperationRecords {
	s.Type = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResultOperationRecords) SetUserId(v string) *GetProcessInstanceResponseBodyResultOperationRecords {
	s.UserId = &v
	return s
}

type GetProcessInstanceResponseBodyResultOperationRecordsAttachments struct {
	// example:
	//
	// 111
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// example:
	//
	// 学历证明
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// 1024
	FileSize *string `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	SpaceId  *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s GetProcessInstanceResponseBodyResultOperationRecordsAttachments) String() string {
	return tea.Prettify(s)
}

func (s GetProcessInstanceResponseBodyResultOperationRecordsAttachments) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceResponseBodyResultOperationRecordsAttachments) SetFileId(v string) *GetProcessInstanceResponseBodyResultOperationRecordsAttachments {
	s.FileId = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResultOperationRecordsAttachments) SetFileName(v string) *GetProcessInstanceResponseBodyResultOperationRecordsAttachments {
	s.FileName = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResultOperationRecordsAttachments) SetFileSize(v string) *GetProcessInstanceResponseBodyResultOperationRecordsAttachments {
	s.FileSize = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResultOperationRecordsAttachments) SetFileType(v string) *GetProcessInstanceResponseBodyResultOperationRecordsAttachments {
	s.FileType = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResultOperationRecordsAttachments) SetSpaceId(v string) *GetProcessInstanceResponseBodyResultOperationRecordsAttachments {
	s.SpaceId = &v
	return s
}

type GetProcessInstanceResponseBodyResultTasks struct {
	// example:
	//
	// 111
	ActivityId *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-08-31T11:52Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-08-31T11:52Z
	FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// example:
	//
	// https://www.xxxx.com
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	// example:
	//
	// https://www.xxxx.com
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	// example:
	//
	// 111
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// example:
	//
	// REDIRECTED
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// NEW
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 111
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// manager1
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetProcessInstanceResponseBodyResultTasks) String() string {
	return tea.Prettify(s)
}

func (s GetProcessInstanceResponseBodyResultTasks) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceResponseBodyResultTasks) SetActivityId(v string) *GetProcessInstanceResponseBodyResultTasks {
	s.ActivityId = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResultTasks) SetCreateTime(v string) *GetProcessInstanceResponseBodyResultTasks {
	s.CreateTime = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResultTasks) SetFinishTime(v string) *GetProcessInstanceResponseBodyResultTasks {
	s.FinishTime = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResultTasks) SetMobileUrl(v string) *GetProcessInstanceResponseBodyResultTasks {
	s.MobileUrl = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResultTasks) SetPcUrl(v string) *GetProcessInstanceResponseBodyResultTasks {
	s.PcUrl = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResultTasks) SetProcessInstanceId(v string) *GetProcessInstanceResponseBodyResultTasks {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResultTasks) SetResult(v string) *GetProcessInstanceResponseBodyResultTasks {
	s.Result = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResultTasks) SetStatus(v string) *GetProcessInstanceResponseBodyResultTasks {
	s.Status = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResultTasks) SetTaskId(v int64) *GetProcessInstanceResponseBodyResultTasks {
	s.TaskId = &v
	return s
}

func (s *GetProcessInstanceResponseBodyResultTasks) SetUserId(v string) *GetProcessInstanceResponseBodyResultTasks {
	s.UserId = &v
	return s
}

type GetProcessInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProcessInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProcessInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProcessInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceResponse) SetHeaders(v map[string]*string) *GetProcessInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetProcessInstanceResponse) SetStatusCode(v int32) *GetProcessInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProcessInstanceResponse) SetBody(v *GetProcessInstanceResponseBody) *GetProcessInstanceResponse {
	s.Body = v
	return s
}

type GetProcessInstanceWithExtraHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetProcessInstanceWithExtraHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetProcessInstanceWithExtraHeaders) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceWithExtraHeaders) SetCommonHeaders(v map[string]*string) *GetProcessInstanceWithExtraHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetProcessInstanceWithExtraHeaders) SetXAcsDingtalkAccessToken(v string) *GetProcessInstanceWithExtraHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetProcessInstanceWithExtraRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a171de6c-8bxxxx
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
}

func (s GetProcessInstanceWithExtraRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProcessInstanceWithExtraRequest) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceWithExtraRequest) SetProcessInstanceId(v string) *GetProcessInstanceWithExtraRequest {
	s.ProcessInstanceId = &v
	return s
}

type GetProcessInstanceWithExtraResponseBody struct {
	Result *GetProcessInstanceWithExtraResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetProcessInstanceWithExtraResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProcessInstanceWithExtraResponseBody) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceWithExtraResponseBody) SetResult(v *GetProcessInstanceWithExtraResponseBodyResult) *GetProcessInstanceWithExtraResponseBody {
	s.Result = v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBody) SetSuccess(v string) *GetProcessInstanceWithExtraResponseBody {
	s.Success = &v
	return s
}

type GetProcessInstanceWithExtraResponseBodyResult struct {
	ApproverUserIds            []*string `json:"approverUserIds,omitempty" xml:"approverUserIds,omitempty" type:"Repeated"`
	AttachedProcessInstanceIds []*string `json:"attachedProcessInstanceIds,omitempty" xml:"attachedProcessInstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// MODIFY
	BizAction *string `json:"bizAction,omitempty" xml:"bizAction,omitempty"`
	BizData   *string `json:"bizData,omitempty" xml:"bizData,omitempty"`
	// example:
	//
	// 20240505xxxx
	BusinessId *string   `json:"businessId,omitempty" xml:"businessId,omitempty"`
	CcUserIds  []*string `json:"ccUserIds,omitempty" xml:"ccUserIds,omitempty" type:"Repeated"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-08-31T11:52Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 2022-08-31T11:52Z
	FinishTime          *string                                                             `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	FormComponentValues []*GetProcessInstanceWithExtraResponseBodyResultFormComponentValues `json:"formComponentValues,omitempty" xml:"formComponentValues,omitempty" type:"Repeated"`
	// example:
	//
	// fvdsfxxxxxx
	MainProcessInstanceId *string                                                          `json:"mainProcessInstanceId,omitempty" xml:"mainProcessInstanceId,omitempty"`
	OperationRecords      []*GetProcessInstanceWithExtraResponseBodyResultOperationRecords `json:"operationRecords,omitempty" xml:"operationRecords,omitempty" type:"Repeated"`
	// example:
	//
	// 25489
	OriginatorDeptId *string `json:"originatorDeptId,omitempty" xml:"originatorDeptId,omitempty"`
	// example:
	//
	// 测试部门
	OriginatorDeptName *string `json:"originatorDeptName,omitempty" xml:"originatorDeptName,omitempty"`
	// example:
	//
	// manager1
	OriginatorUserId *string `json:"originatorUserId,omitempty" xml:"originatorUserId,omitempty"`
	// example:
	//
	// agree
	Result *string                                               `json:"result,omitempty" xml:"result,omitempty"`
	Status *string                                               `json:"status,omitempty" xml:"status,omitempty"`
	Tasks  []*GetProcessInstanceWithExtraResponseBodyResultTasks `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Repeated"`
	// example:
	//
	// xx提交的请假申请
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetProcessInstanceWithExtraResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetProcessInstanceWithExtraResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceWithExtraResponseBodyResult) SetApproverUserIds(v []*string) *GetProcessInstanceWithExtraResponseBodyResult {
	s.ApproverUserIds = v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResult) SetAttachedProcessInstanceIds(v []*string) *GetProcessInstanceWithExtraResponseBodyResult {
	s.AttachedProcessInstanceIds = v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResult) SetBizAction(v string) *GetProcessInstanceWithExtraResponseBodyResult {
	s.BizAction = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResult) SetBizData(v string) *GetProcessInstanceWithExtraResponseBodyResult {
	s.BizData = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResult) SetBusinessId(v string) *GetProcessInstanceWithExtraResponseBodyResult {
	s.BusinessId = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResult) SetCcUserIds(v []*string) *GetProcessInstanceWithExtraResponseBodyResult {
	s.CcUserIds = v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResult) SetCreateTime(v string) *GetProcessInstanceWithExtraResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResult) SetFinishTime(v string) *GetProcessInstanceWithExtraResponseBodyResult {
	s.FinishTime = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResult) SetFormComponentValues(v []*GetProcessInstanceWithExtraResponseBodyResultFormComponentValues) *GetProcessInstanceWithExtraResponseBodyResult {
	s.FormComponentValues = v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResult) SetMainProcessInstanceId(v string) *GetProcessInstanceWithExtraResponseBodyResult {
	s.MainProcessInstanceId = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResult) SetOperationRecords(v []*GetProcessInstanceWithExtraResponseBodyResultOperationRecords) *GetProcessInstanceWithExtraResponseBodyResult {
	s.OperationRecords = v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResult) SetOriginatorDeptId(v string) *GetProcessInstanceWithExtraResponseBodyResult {
	s.OriginatorDeptId = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResult) SetOriginatorDeptName(v string) *GetProcessInstanceWithExtraResponseBodyResult {
	s.OriginatorDeptName = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResult) SetOriginatorUserId(v string) *GetProcessInstanceWithExtraResponseBodyResult {
	s.OriginatorUserId = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResult) SetResult(v string) *GetProcessInstanceWithExtraResponseBodyResult {
	s.Result = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResult) SetStatus(v string) *GetProcessInstanceWithExtraResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResult) SetTasks(v []*GetProcessInstanceWithExtraResponseBodyResultTasks) *GetProcessInstanceWithExtraResponseBodyResult {
	s.Tasks = v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResult) SetTitle(v string) *GetProcessInstanceWithExtraResponseBodyResult {
	s.Title = &v
	return s
}

type GetProcessInstanceWithExtraResponseBodyResultFormComponentValues struct {
	BizAlias      *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	ExtValue      *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	// example:
	//
	// DDHolidayField-J2Bxxxx
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 组件1
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetProcessInstanceWithExtraResponseBodyResultFormComponentValues) String() string {
	return tea.Prettify(s)
}

func (s GetProcessInstanceWithExtraResponseBodyResultFormComponentValues) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceWithExtraResponseBodyResultFormComponentValues) SetBizAlias(v string) *GetProcessInstanceWithExtraResponseBodyResultFormComponentValues {
	s.BizAlias = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultFormComponentValues) SetComponentType(v string) *GetProcessInstanceWithExtraResponseBodyResultFormComponentValues {
	s.ComponentType = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultFormComponentValues) SetExtValue(v string) *GetProcessInstanceWithExtraResponseBodyResultFormComponentValues {
	s.ExtValue = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultFormComponentValues) SetId(v string) *GetProcessInstanceWithExtraResponseBodyResultFormComponentValues {
	s.Id = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultFormComponentValues) SetName(v string) *GetProcessInstanceWithExtraResponseBodyResultFormComponentValues {
	s.Name = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultFormComponentValues) SetValue(v string) *GetProcessInstanceWithExtraResponseBodyResultFormComponentValues {
	s.Value = &v
	return s
}

type GetProcessInstanceWithExtraResponseBodyResultOperationRecords struct {
	// example:
	//
	// aacc_ddee
	ActivityId  *string                                                                     `json:"activityId,omitempty" xml:"activityId,omitempty"`
	Attachments []*GetProcessInstanceWithExtraResponseBodyResultOperationRecordsAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	CcUserIds   []*string                                                                   `json:"ccUserIds,omitempty" xml:"ccUserIds,omitempty" type:"Repeated"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-08-31T11:52Z
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// example:
	//
	// AzBltRlvKukX3WsbLxpDnTZskRNK5GtVfuDlDQ_Qxsp
	HandSignToken *string   `json:"handSignToken,omitempty" xml:"handSignToken,omitempty"`
	Images        []*string `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	Remark        *string   `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// AGREE
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// 审批人节点
	ShowName *string `json:"showName,omitempty" xml:"showName,omitempty"`
	// example:
	//
	// EXECUTE_TASK_NORMAL
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// manager123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetProcessInstanceWithExtraResponseBodyResultOperationRecords) String() string {
	return tea.Prettify(s)
}

func (s GetProcessInstanceWithExtraResponseBodyResultOperationRecords) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceWithExtraResponseBodyResultOperationRecords) SetActivityId(v string) *GetProcessInstanceWithExtraResponseBodyResultOperationRecords {
	s.ActivityId = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultOperationRecords) SetAttachments(v []*GetProcessInstanceWithExtraResponseBodyResultOperationRecordsAttachments) *GetProcessInstanceWithExtraResponseBodyResultOperationRecords {
	s.Attachments = v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultOperationRecords) SetCcUserIds(v []*string) *GetProcessInstanceWithExtraResponseBodyResultOperationRecords {
	s.CcUserIds = v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultOperationRecords) SetDate(v string) *GetProcessInstanceWithExtraResponseBodyResultOperationRecords {
	s.Date = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultOperationRecords) SetHandSignToken(v string) *GetProcessInstanceWithExtraResponseBodyResultOperationRecords {
	s.HandSignToken = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultOperationRecords) SetImages(v []*string) *GetProcessInstanceWithExtraResponseBodyResultOperationRecords {
	s.Images = v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultOperationRecords) SetRemark(v string) *GetProcessInstanceWithExtraResponseBodyResultOperationRecords {
	s.Remark = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultOperationRecords) SetResult(v string) *GetProcessInstanceWithExtraResponseBodyResultOperationRecords {
	s.Result = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultOperationRecords) SetShowName(v string) *GetProcessInstanceWithExtraResponseBodyResultOperationRecords {
	s.ShowName = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultOperationRecords) SetType(v string) *GetProcessInstanceWithExtraResponseBodyResultOperationRecords {
	s.Type = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultOperationRecords) SetUserId(v string) *GetProcessInstanceWithExtraResponseBodyResultOperationRecords {
	s.UserId = &v
	return s
}

type GetProcessInstanceWithExtraResponseBodyResultOperationRecordsAttachments struct {
	// example:
	//
	// 12345
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// example:
	//
	// 学历证明
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// 50000
	FileSize *string `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// example:
	//
	// .pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// 158469
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s GetProcessInstanceWithExtraResponseBodyResultOperationRecordsAttachments) String() string {
	return tea.Prettify(s)
}

func (s GetProcessInstanceWithExtraResponseBodyResultOperationRecordsAttachments) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceWithExtraResponseBodyResultOperationRecordsAttachments) SetFileId(v string) *GetProcessInstanceWithExtraResponseBodyResultOperationRecordsAttachments {
	s.FileId = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultOperationRecordsAttachments) SetFileName(v string) *GetProcessInstanceWithExtraResponseBodyResultOperationRecordsAttachments {
	s.FileName = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultOperationRecordsAttachments) SetFileSize(v string) *GetProcessInstanceWithExtraResponseBodyResultOperationRecordsAttachments {
	s.FileSize = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultOperationRecordsAttachments) SetFileType(v string) *GetProcessInstanceWithExtraResponseBodyResultOperationRecordsAttachments {
	s.FileType = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultOperationRecordsAttachments) SetSpaceId(v string) *GetProcessInstanceWithExtraResponseBodyResultOperationRecordsAttachments {
	s.SpaceId = &v
	return s
}

type GetProcessInstanceWithExtraResponseBodyResultTasks struct {
	// example:
	//
	// aabb_ccdd
	ActivityId *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2024-06-12T14:17Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2024-06-12T14:17Z
	FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// example:
	//
	// aflow.dingtalk.com?procInsId=lTGxxx
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	// example:
	//
	// aflow.dingtalk.com?procInsId=lTGxxx
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	// example:
	//
	// fewferxxxx
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// example:
	//
	// REDIRECTED
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// CANCELED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	TaskId *int64  `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// manager123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetProcessInstanceWithExtraResponseBodyResultTasks) String() string {
	return tea.Prettify(s)
}

func (s GetProcessInstanceWithExtraResponseBodyResultTasks) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceWithExtraResponseBodyResultTasks) SetActivityId(v string) *GetProcessInstanceWithExtraResponseBodyResultTasks {
	s.ActivityId = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultTasks) SetCreateTime(v string) *GetProcessInstanceWithExtraResponseBodyResultTasks {
	s.CreateTime = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultTasks) SetFinishTime(v string) *GetProcessInstanceWithExtraResponseBodyResultTasks {
	s.FinishTime = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultTasks) SetMobileUrl(v string) *GetProcessInstanceWithExtraResponseBodyResultTasks {
	s.MobileUrl = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultTasks) SetPcUrl(v string) *GetProcessInstanceWithExtraResponseBodyResultTasks {
	s.PcUrl = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultTasks) SetProcessInstanceId(v string) *GetProcessInstanceWithExtraResponseBodyResultTasks {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultTasks) SetResult(v string) *GetProcessInstanceWithExtraResponseBodyResultTasks {
	s.Result = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultTasks) SetStatus(v string) *GetProcessInstanceWithExtraResponseBodyResultTasks {
	s.Status = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultTasks) SetTaskId(v int64) *GetProcessInstanceWithExtraResponseBodyResultTasks {
	s.TaskId = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponseBodyResultTasks) SetUserId(v string) *GetProcessInstanceWithExtraResponseBodyResultTasks {
	s.UserId = &v
	return s
}

type GetProcessInstanceWithExtraResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProcessInstanceWithExtraResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProcessInstanceWithExtraResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProcessInstanceWithExtraResponse) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceWithExtraResponse) SetHeaders(v map[string]*string) *GetProcessInstanceWithExtraResponse {
	s.Headers = v
	return s
}

func (s *GetProcessInstanceWithExtraResponse) SetStatusCode(v int32) *GetProcessInstanceWithExtraResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProcessInstanceWithExtraResponse) SetBody(v *GetProcessInstanceWithExtraResponseBody) *GetProcessInstanceWithExtraResponse {
	s.Body = v
	return s
}

type GetSchemaAndProcessconfigBatchllyHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSchemaAndProcessconfigBatchllyHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSchemaAndProcessconfigBatchllyHeaders) GoString() string {
	return s.String()
}

func (s *GetSchemaAndProcessconfigBatchllyHeaders) SetCommonHeaders(v map[string]*string) *GetSchemaAndProcessconfigBatchllyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSchemaAndProcessconfigBatchllyHeaders) SetXAcsDingtalkAccessToken(v string) *GetSchemaAndProcessconfigBatchllyHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSchemaAndProcessconfigBatchllyRequest struct {
	ProcessCodes []*string `json:"processCodes,omitempty" xml:"processCodes,omitempty" type:"Repeated"`
}

func (s GetSchemaAndProcessconfigBatchllyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSchemaAndProcessconfigBatchllyRequest) GoString() string {
	return s.String()
}

func (s *GetSchemaAndProcessconfigBatchllyRequest) SetProcessCodes(v []*string) *GetSchemaAndProcessconfigBatchllyRequest {
	s.ProcessCodes = v
	return s
}

type GetSchemaAndProcessconfigBatchllyShrinkRequest struct {
	ProcessCodesShrink *string `json:"processCodes,omitempty" xml:"processCodes,omitempty"`
}

func (s GetSchemaAndProcessconfigBatchllyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSchemaAndProcessconfigBatchllyShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSchemaAndProcessconfigBatchllyShrinkRequest) SetProcessCodesShrink(v string) *GetSchemaAndProcessconfigBatchllyShrinkRequest {
	s.ProcessCodesShrink = &v
	return s
}

type GetSchemaAndProcessconfigBatchllyResponseBody struct {
	Result []*GetSchemaAndProcessconfigBatchllyResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetSchemaAndProcessconfigBatchllyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSchemaAndProcessconfigBatchllyResponseBody) GoString() string {
	return s.String()
}

func (s *GetSchemaAndProcessconfigBatchllyResponseBody) SetResult(v []*GetSchemaAndProcessconfigBatchllyResponseBodyResult) *GetSchemaAndProcessconfigBatchllyResponseBody {
	s.Result = v
	return s
}

type GetSchemaAndProcessconfigBatchllyResponseBodyResult struct {
	AppUuid       *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	BizCategoryId *string `json:"bizCategoryId,omitempty" xml:"bizCategoryId,omitempty"`
	CreateTime    *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	FormUuid      *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	ManagerList   *string `json:"managerList,omitempty" xml:"managerList,omitempty"`
	Memo          *string `json:"memo,omitempty" xml:"memo,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	ProcessCode   *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	ProcessConfig *string `json:"processConfig,omitempty" xml:"processConfig,omitempty"`
	ProcessId     *int64  `json:"processId,omitempty" xml:"processId,omitempty"`
	Properties    *string `json:"properties,omitempty" xml:"properties,omitempty"`
	SchemaContent *string `json:"schemaContent,omitempty" xml:"schemaContent,omitempty"`
	VisibleScope  *string `json:"visibleScope,omitempty" xml:"visibleScope,omitempty"`
}

func (s GetSchemaAndProcessconfigBatchllyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetSchemaAndProcessconfigBatchllyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSchemaAndProcessconfigBatchllyResponseBodyResult) SetAppUuid(v string) *GetSchemaAndProcessconfigBatchllyResponseBodyResult {
	s.AppUuid = &v
	return s
}

func (s *GetSchemaAndProcessconfigBatchllyResponseBodyResult) SetBizCategoryId(v string) *GetSchemaAndProcessconfigBatchllyResponseBodyResult {
	s.BizCategoryId = &v
	return s
}

func (s *GetSchemaAndProcessconfigBatchllyResponseBodyResult) SetCreateTime(v string) *GetSchemaAndProcessconfigBatchllyResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetSchemaAndProcessconfigBatchllyResponseBodyResult) SetCreatorUserId(v string) *GetSchemaAndProcessconfigBatchllyResponseBodyResult {
	s.CreatorUserId = &v
	return s
}

func (s *GetSchemaAndProcessconfigBatchllyResponseBodyResult) SetFormUuid(v string) *GetSchemaAndProcessconfigBatchllyResponseBodyResult {
	s.FormUuid = &v
	return s
}

func (s *GetSchemaAndProcessconfigBatchllyResponseBodyResult) SetManagerList(v string) *GetSchemaAndProcessconfigBatchllyResponseBodyResult {
	s.ManagerList = &v
	return s
}

func (s *GetSchemaAndProcessconfigBatchllyResponseBodyResult) SetMemo(v string) *GetSchemaAndProcessconfigBatchllyResponseBodyResult {
	s.Memo = &v
	return s
}

func (s *GetSchemaAndProcessconfigBatchllyResponseBodyResult) SetName(v string) *GetSchemaAndProcessconfigBatchllyResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetSchemaAndProcessconfigBatchllyResponseBodyResult) SetProcessCode(v string) *GetSchemaAndProcessconfigBatchllyResponseBodyResult {
	s.ProcessCode = &v
	return s
}

func (s *GetSchemaAndProcessconfigBatchllyResponseBodyResult) SetProcessConfig(v string) *GetSchemaAndProcessconfigBatchllyResponseBodyResult {
	s.ProcessConfig = &v
	return s
}

func (s *GetSchemaAndProcessconfigBatchllyResponseBodyResult) SetProcessId(v int64) *GetSchemaAndProcessconfigBatchllyResponseBodyResult {
	s.ProcessId = &v
	return s
}

func (s *GetSchemaAndProcessconfigBatchllyResponseBodyResult) SetProperties(v string) *GetSchemaAndProcessconfigBatchllyResponseBodyResult {
	s.Properties = &v
	return s
}

func (s *GetSchemaAndProcessconfigBatchllyResponseBodyResult) SetSchemaContent(v string) *GetSchemaAndProcessconfigBatchllyResponseBodyResult {
	s.SchemaContent = &v
	return s
}

func (s *GetSchemaAndProcessconfigBatchllyResponseBodyResult) SetVisibleScope(v string) *GetSchemaAndProcessconfigBatchllyResponseBodyResult {
	s.VisibleScope = &v
	return s
}

type GetSchemaAndProcessconfigBatchllyResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSchemaAndProcessconfigBatchllyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSchemaAndProcessconfigBatchllyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSchemaAndProcessconfigBatchllyResponse) GoString() string {
	return s.String()
}

func (s *GetSchemaAndProcessconfigBatchllyResponse) SetHeaders(v map[string]*string) *GetSchemaAndProcessconfigBatchllyResponse {
	s.Headers = v
	return s
}

func (s *GetSchemaAndProcessconfigBatchllyResponse) SetStatusCode(v int32) *GetSchemaAndProcessconfigBatchllyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSchemaAndProcessconfigBatchllyResponse) SetBody(v *GetSchemaAndProcessconfigBatchllyResponseBody) *GetSchemaAndProcessconfigBatchllyResponse {
	s.Body = v
	return s
}

type GetSpaceWithDownloadAuthHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetSpaceWithDownloadAuthHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceWithDownloadAuthHeaders) GoString() string {
	return s.String()
}

func (s *GetSpaceWithDownloadAuthHeaders) SetCommonHeaders(v map[string]*string) *GetSpaceWithDownloadAuthHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSpaceWithDownloadAuthHeaders) SetXAcsDingtalkAccessToken(v string) *GetSpaceWithDownloadAuthHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetSpaceWithDownloadAuthRequest struct {
	// example:
	//
	// 8345000
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111
	FileId     *string   `json:"fileId,omitempty" xml:"fileId,omitempty"`
	FileIdList []*string `json:"fileIdList,omitempty" xml:"fileIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// a17444d1-075b-4a4d-xxxx
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// if can be null:
	// true
	WithCommentAttatchment *bool `json:"withCommentAttatchment,omitempty" xml:"withCommentAttatchment,omitempty"`
}

func (s GetSpaceWithDownloadAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceWithDownloadAuthRequest) GoString() string {
	return s.String()
}

func (s *GetSpaceWithDownloadAuthRequest) SetAgentId(v int64) *GetSpaceWithDownloadAuthRequest {
	s.AgentId = &v
	return s
}

func (s *GetSpaceWithDownloadAuthRequest) SetFileId(v string) *GetSpaceWithDownloadAuthRequest {
	s.FileId = &v
	return s
}

func (s *GetSpaceWithDownloadAuthRequest) SetFileIdList(v []*string) *GetSpaceWithDownloadAuthRequest {
	s.FileIdList = v
	return s
}

func (s *GetSpaceWithDownloadAuthRequest) SetProcessInstanceId(v string) *GetSpaceWithDownloadAuthRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetSpaceWithDownloadAuthRequest) SetUserId(v string) *GetSpaceWithDownloadAuthRequest {
	s.UserId = &v
	return s
}

func (s *GetSpaceWithDownloadAuthRequest) SetWithCommentAttatchment(v bool) *GetSpaceWithDownloadAuthRequest {
	s.WithCommentAttatchment = &v
	return s
}

type GetSpaceWithDownloadAuthResponseBody struct {
	Result *GetSpaceWithDownloadAuthResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetSpaceWithDownloadAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceWithDownloadAuthResponseBody) GoString() string {
	return s.String()
}

func (s *GetSpaceWithDownloadAuthResponseBody) SetResult(v *GetSpaceWithDownloadAuthResponseBodyResult) *GetSpaceWithDownloadAuthResponseBody {
	s.Result = v
	return s
}

func (s *GetSpaceWithDownloadAuthResponseBody) SetSuccess(v bool) *GetSpaceWithDownloadAuthResponseBody {
	s.Success = &v
	return s
}

type GetSpaceWithDownloadAuthResponseBodyResult struct {
	// example:
	//
	// 3996960664
	SpaceId *int64 `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s GetSpaceWithDownloadAuthResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceWithDownloadAuthResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSpaceWithDownloadAuthResponseBodyResult) SetSpaceId(v int64) *GetSpaceWithDownloadAuthResponseBodyResult {
	s.SpaceId = &v
	return s
}

type GetSpaceWithDownloadAuthResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSpaceWithDownloadAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSpaceWithDownloadAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSpaceWithDownloadAuthResponse) GoString() string {
	return s.String()
}

func (s *GetSpaceWithDownloadAuthResponse) SetHeaders(v map[string]*string) *GetSpaceWithDownloadAuthResponse {
	s.Headers = v
	return s
}

func (s *GetSpaceWithDownloadAuthResponse) SetStatusCode(v int32) *GetSpaceWithDownloadAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSpaceWithDownloadAuthResponse) SetBody(v *GetSpaceWithDownloadAuthResponseBody) *GetSpaceWithDownloadAuthResponse {
	s.Body = v
	return s
}

type GetUserTodoTaskSumHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GetUserTodoTaskSumHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUserTodoTaskSumHeaders) GoString() string {
	return s.String()
}

func (s *GetUserTodoTaskSumHeaders) SetCommonHeaders(v map[string]*string) *GetUserTodoTaskSumHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserTodoTaskSumHeaders) SetXAcsDingtalkAccessToken(v string) *GetUserTodoTaskSumHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GetUserTodoTaskSumRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// manager123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUserTodoTaskSumRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserTodoTaskSumRequest) GoString() string {
	return s.String()
}

func (s *GetUserTodoTaskSumRequest) SetUserId(v string) *GetUserTodoTaskSumRequest {
	s.UserId = &v
	return s
}

type GetUserTodoTaskSumResponseBody struct {
	// example:
	//
	// 10
	Result *int32 `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetUserTodoTaskSumResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserTodoTaskSumResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserTodoTaskSumResponseBody) SetResult(v int32) *GetUserTodoTaskSumResponseBody {
	s.Result = &v
	return s
}

type GetUserTodoTaskSumResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserTodoTaskSumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserTodoTaskSumResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserTodoTaskSumResponse) GoString() string {
	return s.String()
}

func (s *GetUserTodoTaskSumResponse) SetHeaders(v map[string]*string) *GetUserTodoTaskSumResponse {
	s.Headers = v
	return s
}

func (s *GetUserTodoTaskSumResponse) SetStatusCode(v int32) *GetUserTodoTaskSumResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserTodoTaskSumResponse) SetBody(v *GetUserTodoTaskSumResponseBody) *GetUserTodoTaskSumResponse {
	s.Body = v
	return s
}

type GrantCspaceAuthorizationHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GrantCspaceAuthorizationHeaders) String() string {
	return tea.Prettify(s)
}

func (s GrantCspaceAuthorizationHeaders) GoString() string {
	return s.String()
}

func (s *GrantCspaceAuthorizationHeaders) SetCommonHeaders(v map[string]*string) *GrantCspaceAuthorizationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GrantCspaceAuthorizationHeaders) SetXAcsDingtalkAccessToken(v string) *GrantCspaceAuthorizationHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GrantCspaceAuthorizationRequest struct {
	// example:
	//
	// 3600
	DurationSeconds *int64 `json:"durationSeconds,omitempty" xml:"durationSeconds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 163xxxx658
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// add
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 26652461xxxx5992
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GrantCspaceAuthorizationRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantCspaceAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *GrantCspaceAuthorizationRequest) SetDurationSeconds(v int64) *GrantCspaceAuthorizationRequest {
	s.DurationSeconds = &v
	return s
}

func (s *GrantCspaceAuthorizationRequest) SetSpaceId(v string) *GrantCspaceAuthorizationRequest {
	s.SpaceId = &v
	return s
}

func (s *GrantCspaceAuthorizationRequest) SetType(v string) *GrantCspaceAuthorizationRequest {
	s.Type = &v
	return s
}

func (s *GrantCspaceAuthorizationRequest) SetUserId(v string) *GrantCspaceAuthorizationRequest {
	s.UserId = &v
	return s
}

type GrantCspaceAuthorizationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s GrantCspaceAuthorizationResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantCspaceAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *GrantCspaceAuthorizationResponse) SetHeaders(v map[string]*string) *GrantCspaceAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *GrantCspaceAuthorizationResponse) SetStatusCode(v int32) *GrantCspaceAuthorizationResponse {
	s.StatusCode = &v
	return s
}

type GrantProcessInstanceForDownloadFileHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s GrantProcessInstanceForDownloadFileHeaders) String() string {
	return tea.Prettify(s)
}

func (s GrantProcessInstanceForDownloadFileHeaders) GoString() string {
	return s.String()
}

func (s *GrantProcessInstanceForDownloadFileHeaders) SetCommonHeaders(v map[string]*string) *GrantProcessInstanceForDownloadFileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GrantProcessInstanceForDownloadFileHeaders) SetXAcsDingtalkAccessToken(v string) *GrantProcessInstanceForDownloadFileHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type GrantProcessInstanceForDownloadFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 111
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a17444d1-075b-4a4d-xxxx
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// if can be null:
	// true
	WithCommentAttatchment *bool `json:"withCommentAttatchment,omitempty" xml:"withCommentAttatchment,omitempty"`
}

func (s GrantProcessInstanceForDownloadFileRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantProcessInstanceForDownloadFileRequest) GoString() string {
	return s.String()
}

func (s *GrantProcessInstanceForDownloadFileRequest) SetFileId(v string) *GrantProcessInstanceForDownloadFileRequest {
	s.FileId = &v
	return s
}

func (s *GrantProcessInstanceForDownloadFileRequest) SetProcessInstanceId(v string) *GrantProcessInstanceForDownloadFileRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *GrantProcessInstanceForDownloadFileRequest) SetWithCommentAttatchment(v bool) *GrantProcessInstanceForDownloadFileRequest {
	s.WithCommentAttatchment = &v
	return s
}

type GrantProcessInstanceForDownloadFileResponseBody struct {
	Result *GrantProcessInstanceForDownloadFileResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GrantProcessInstanceForDownloadFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantProcessInstanceForDownloadFileResponseBody) GoString() string {
	return s.String()
}

func (s *GrantProcessInstanceForDownloadFileResponseBody) SetResult(v *GrantProcessInstanceForDownloadFileResponseBodyResult) *GrantProcessInstanceForDownloadFileResponseBody {
	s.Result = v
	return s
}

func (s *GrantProcessInstanceForDownloadFileResponseBody) SetSuccess(v bool) *GrantProcessInstanceForDownloadFileResponseBody {
	s.Success = &v
	return s
}

type GrantProcessInstanceForDownloadFileResponseBodyResult struct {
	// example:
	//
	// http://lippi-space-zjk.oss-cn-zhangjiakou.aliyuncs.com/xxxxx
	DownloadUri *string `json:"downloadUri,omitempty" xml:"downloadUri,omitempty"`
	// example:
	//
	// 26748422566
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// example:
	//
	// 3996960664
	SpaceId *int64 `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s GrantProcessInstanceForDownloadFileResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GrantProcessInstanceForDownloadFileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GrantProcessInstanceForDownloadFileResponseBodyResult) SetDownloadUri(v string) *GrantProcessInstanceForDownloadFileResponseBodyResult {
	s.DownloadUri = &v
	return s
}

func (s *GrantProcessInstanceForDownloadFileResponseBodyResult) SetFileId(v string) *GrantProcessInstanceForDownloadFileResponseBodyResult {
	s.FileId = &v
	return s
}

func (s *GrantProcessInstanceForDownloadFileResponseBodyResult) SetSpaceId(v int64) *GrantProcessInstanceForDownloadFileResponseBodyResult {
	s.SpaceId = &v
	return s
}

type GrantProcessInstanceForDownloadFileResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantProcessInstanceForDownloadFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantProcessInstanceForDownloadFileResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantProcessInstanceForDownloadFileResponse) GoString() string {
	return s.String()
}

func (s *GrantProcessInstanceForDownloadFileResponse) SetHeaders(v map[string]*string) *GrantProcessInstanceForDownloadFileResponse {
	s.Headers = v
	return s
}

func (s *GrantProcessInstanceForDownloadFileResponse) SetStatusCode(v int32) *GrantProcessInstanceForDownloadFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantProcessInstanceForDownloadFileResponse) SetBody(v *GrantProcessInstanceForDownloadFileResponseBody) *GrantProcessInstanceForDownloadFileResponse {
	s.Body = v
	return s
}

type InsertOrUpdateDirHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InsertOrUpdateDirHeaders) String() string {
	return tea.Prettify(s)
}

func (s InsertOrUpdateDirHeaders) GoString() string {
	return s.String()
}

func (s *InsertOrUpdateDirHeaders) SetCommonHeaders(v map[string]*string) *InsertOrUpdateDirHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsertOrUpdateDirHeaders) SetXAcsDingtalkAccessToken(v string) *InsertOrUpdateDirHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InsertOrUpdateDirRequest struct {
	// example:
	//
	// administeration
	BizGroup *string `json:"bizGroup,omitempty" xml:"bizGroup,omitempty"`
	// example:
	//
	// 分组描述信息
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 行政管理
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {\"en_US\":\"test\",\"ja_JP\":\"test\",\"vi_VN\":\"test\",\"zh_CN\":\"测试\",\"zh_HK\":\"测试\",\"zh_TW\":\"测试\"}
	Name18n *string `json:"name18n,omitempty" xml:"name18n,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user001
	OperateUserId *string `json:"operateUserId,omitempty" xml:"operateUserId,omitempty"`
}

func (s InsertOrUpdateDirRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertOrUpdateDirRequest) GoString() string {
	return s.String()
}

func (s *InsertOrUpdateDirRequest) SetBizGroup(v string) *InsertOrUpdateDirRequest {
	s.BizGroup = &v
	return s
}

func (s *InsertOrUpdateDirRequest) SetDescription(v string) *InsertOrUpdateDirRequest {
	s.Description = &v
	return s
}

func (s *InsertOrUpdateDirRequest) SetName(v string) *InsertOrUpdateDirRequest {
	s.Name = &v
	return s
}

func (s *InsertOrUpdateDirRequest) SetName18n(v string) *InsertOrUpdateDirRequest {
	s.Name18n = &v
	return s
}

func (s *InsertOrUpdateDirRequest) SetOperateUserId(v string) *InsertOrUpdateDirRequest {
	s.OperateUserId = &v
	return s
}

type InsertOrUpdateDirResponseBody struct {
	Result  *InsertOrUpdateDirResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                `json:"success,omitempty" xml:"success,omitempty"`
}

func (s InsertOrUpdateDirResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertOrUpdateDirResponseBody) GoString() string {
	return s.String()
}

func (s *InsertOrUpdateDirResponseBody) SetResult(v *InsertOrUpdateDirResponseBodyResult) *InsertOrUpdateDirResponseBody {
	s.Result = v
	return s
}

func (s *InsertOrUpdateDirResponseBody) SetSuccess(v bool) *InsertOrUpdateDirResponseBody {
	s.Success = &v
	return s
}

type InsertOrUpdateDirResponseBodyResult struct {
	// example:
	//
	// {应用appId}_administeration
	BizGroup *string `json:"bizGroup,omitempty" xml:"bizGroup,omitempty"`
	// example:
	//
	// oaDirIdxxx
	DirId *string `json:"dirId,omitempty" xml:"dirId,omitempty"`
}

func (s InsertOrUpdateDirResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s InsertOrUpdateDirResponseBodyResult) GoString() string {
	return s.String()
}

func (s *InsertOrUpdateDirResponseBodyResult) SetBizGroup(v string) *InsertOrUpdateDirResponseBodyResult {
	s.BizGroup = &v
	return s
}

func (s *InsertOrUpdateDirResponseBodyResult) SetDirId(v string) *InsertOrUpdateDirResponseBodyResult {
	s.DirId = &v
	return s
}

type InsertOrUpdateDirResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertOrUpdateDirResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertOrUpdateDirResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertOrUpdateDirResponse) GoString() string {
	return s.String()
}

func (s *InsertOrUpdateDirResponse) SetHeaders(v map[string]*string) *InsertOrUpdateDirResponse {
	s.Headers = v
	return s
}

func (s *InsertOrUpdateDirResponse) SetStatusCode(v int32) *InsertOrUpdateDirResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertOrUpdateDirResponse) SetBody(v *InsertOrUpdateDirResponseBody) *InsertOrUpdateDirResponse {
	s.Body = v
	return s
}

type InstallAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s InstallAppHeaders) String() string {
	return tea.Prettify(s)
}

func (s InstallAppHeaders) GoString() string {
	return s.String()
}

func (s *InstallAppHeaders) SetCommonHeaders(v map[string]*string) *InstallAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InstallAppHeaders) SetXAcsDingtalkAccessToken(v string) *InstallAppHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type InstallAppRequest struct {
	BizGroup *string `json:"bizGroup,omitempty" xml:"bizGroup,omitempty"`
	// This parameter is required.
	InstallOption *InstallAppRequestInstallOption `json:"installOption,omitempty" xml:"installOption,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// finance
	SourceDirName *string `json:"sourceDirName,omitempty" xml:"sourceDirName,omitempty"`
}

func (s InstallAppRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallAppRequest) GoString() string {
	return s.String()
}

func (s *InstallAppRequest) SetBizGroup(v string) *InstallAppRequest {
	s.BizGroup = &v
	return s
}

func (s *InstallAppRequest) SetInstallOption(v *InstallAppRequestInstallOption) *InstallAppRequest {
	s.InstallOption = v
	return s
}

func (s *InstallAppRequest) SetSourceDirName(v string) *InstallAppRequest {
	s.SourceDirName = &v
	return s
}

type InstallAppRequestInstallOption struct {
	IsSync *bool `json:"isSync,omitempty" xml:"isSync,omitempty"`
}

func (s InstallAppRequestInstallOption) String() string {
	return tea.Prettify(s)
}

func (s InstallAppRequestInstallOption) GoString() string {
	return s.String()
}

func (s *InstallAppRequestInstallOption) SetIsSync(v bool) *InstallAppRequestInstallOption {
	s.IsSync = &v
	return s
}

type InstallAppResponseBody struct {
	Result []*InstallAppResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s InstallAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallAppResponseBody) GoString() string {
	return s.String()
}

func (s *InstallAppResponseBody) SetResult(v []*InstallAppResponseBodyResult) *InstallAppResponseBody {
	s.Result = v
	return s
}

type InstallAppResponseBodyResult struct {
	// example:
	//
	// abc
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// example:
	//
	// abc
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// PROC-ABC
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
}

func (s InstallAppResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s InstallAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *InstallAppResponseBodyResult) SetBizType(v string) *InstallAppResponseBodyResult {
	s.BizType = &v
	return s
}

func (s *InstallAppResponseBodyResult) SetName(v string) *InstallAppResponseBodyResult {
	s.Name = &v
	return s
}

func (s *InstallAppResponseBodyResult) SetProcessCode(v string) *InstallAppResponseBodyResult {
	s.ProcessCode = &v
	return s
}

type InstallAppResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallAppResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallAppResponse) GoString() string {
	return s.String()
}

func (s *InstallAppResponse) SetHeaders(v map[string]*string) *InstallAppResponse {
	s.Headers = v
	return s
}

func (s *InstallAppResponse) SetStatusCode(v int32) *InstallAppResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallAppResponse) SetBody(v *InstallAppResponseBody) *InstallAppResponse {
	s.Body = v
	return s
}

type ListProcessInstanceIdsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListProcessInstanceIdsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListProcessInstanceIdsHeaders) GoString() string {
	return s.String()
}

func (s *ListProcessInstanceIdsHeaders) SetCommonHeaders(v map[string]*string) *ListProcessInstanceIdsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListProcessInstanceIdsHeaders) SetXAcsDingtalkAccessToken(v string) *ListProcessInstanceIdsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListProcessInstanceIdsRequest struct {
	// example:
	//
	// 1496678400000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-FF6Y2xxxx
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1496678400000
	StartTime *int64    `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Statuses  []*string `json:"statuses,omitempty" xml:"statuses,omitempty" type:"Repeated"`
	UserIds   []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
}

func (s ListProcessInstanceIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProcessInstanceIdsRequest) GoString() string {
	return s.String()
}

func (s *ListProcessInstanceIdsRequest) SetEndTime(v int64) *ListProcessInstanceIdsRequest {
	s.EndTime = &v
	return s
}

func (s *ListProcessInstanceIdsRequest) SetMaxResults(v int64) *ListProcessInstanceIdsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProcessInstanceIdsRequest) SetNextToken(v int64) *ListProcessInstanceIdsRequest {
	s.NextToken = &v
	return s
}

func (s *ListProcessInstanceIdsRequest) SetProcessCode(v string) *ListProcessInstanceIdsRequest {
	s.ProcessCode = &v
	return s
}

func (s *ListProcessInstanceIdsRequest) SetStartTime(v int64) *ListProcessInstanceIdsRequest {
	s.StartTime = &v
	return s
}

func (s *ListProcessInstanceIdsRequest) SetStatuses(v []*string) *ListProcessInstanceIdsRequest {
	s.Statuses = v
	return s
}

func (s *ListProcessInstanceIdsRequest) SetUserIds(v []*string) *ListProcessInstanceIdsRequest {
	s.UserIds = v
	return s
}

type ListProcessInstanceIdsResponseBody struct {
	Result *ListProcessInstanceIdsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListProcessInstanceIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProcessInstanceIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProcessInstanceIdsResponseBody) SetResult(v *ListProcessInstanceIdsResponseBodyResult) *ListProcessInstanceIdsResponseBody {
	s.Result = v
	return s
}

func (s *ListProcessInstanceIdsResponseBody) SetSuccess(v bool) *ListProcessInstanceIdsResponseBody {
	s.Success = &v
	return s
}

type ListProcessInstanceIdsResponseBodyResult struct {
	List []*string `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListProcessInstanceIdsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListProcessInstanceIdsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListProcessInstanceIdsResponseBodyResult) SetList(v []*string) *ListProcessInstanceIdsResponseBodyResult {
	s.List = v
	return s
}

func (s *ListProcessInstanceIdsResponseBodyResult) SetNextToken(v string) *ListProcessInstanceIdsResponseBodyResult {
	s.NextToken = &v
	return s
}

type ListProcessInstanceIdsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProcessInstanceIdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProcessInstanceIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProcessInstanceIdsResponse) GoString() string {
	return s.String()
}

func (s *ListProcessInstanceIdsResponse) SetHeaders(v map[string]*string) *ListProcessInstanceIdsResponse {
	s.Headers = v
	return s
}

func (s *ListProcessInstanceIdsResponse) SetStatusCode(v int32) *ListProcessInstanceIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProcessInstanceIdsResponse) SetBody(v *ListProcessInstanceIdsResponseBody) *ListProcessInstanceIdsResponse {
	s.Body = v
	return s
}

type ListTodoWorkRecordsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListTodoWorkRecordsHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListTodoWorkRecordsHeaders) GoString() string {
	return s.String()
}

func (s *ListTodoWorkRecordsHeaders) SetCommonHeaders(v map[string]*string) *ListTodoWorkRecordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListTodoWorkRecordsHeaders) SetXAcsDingtalkAccessToken(v string) *ListTodoWorkRecordsHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListTodoWorkRecordsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	NextToken *int32 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// manager001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListTodoWorkRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTodoWorkRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListTodoWorkRecordsRequest) SetMaxResults(v int32) *ListTodoWorkRecordsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTodoWorkRecordsRequest) SetNextToken(v int32) *ListTodoWorkRecordsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTodoWorkRecordsRequest) SetStatus(v int32) *ListTodoWorkRecordsRequest {
	s.Status = &v
	return s
}

func (s *ListTodoWorkRecordsRequest) SetUserId(v string) *ListTodoWorkRecordsRequest {
	s.UserId = &v
	return s
}

type ListTodoWorkRecordsResponseBody struct {
	Result *ListTodoWorkRecordsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ListTodoWorkRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTodoWorkRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTodoWorkRecordsResponseBody) SetResult(v *ListTodoWorkRecordsResponseBodyResult) *ListTodoWorkRecordsResponseBody {
	s.Result = v
	return s
}

type ListTodoWorkRecordsResponseBodyResult struct {
	List      []*ListTodoWorkRecordsResponseBodyResultList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	NextToken *int64                                       `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListTodoWorkRecordsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListTodoWorkRecordsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListTodoWorkRecordsResponseBodyResult) SetList(v []*ListTodoWorkRecordsResponseBodyResultList) *ListTodoWorkRecordsResponseBodyResult {
	s.List = v
	return s
}

func (s *ListTodoWorkRecordsResponseBodyResult) SetNextToken(v int64) *ListTodoWorkRecordsResponseBodyResult {
	s.NextToken = &v
	return s
}

type ListTodoWorkRecordsResponseBodyResultList struct {
	Forms []*ListTodoWorkRecordsResponseBodyResultListForms `json:"forms,omitempty" xml:"forms,omitempty" type:"Repeated"`
	// example:
	//
	// Siw2WNVZS4KiUt3tTmaNKg04*****809950
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 1234567
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// xxx提交的入职审批
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListTodoWorkRecordsResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s ListTodoWorkRecordsResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *ListTodoWorkRecordsResponseBodyResultList) SetForms(v []*ListTodoWorkRecordsResponseBodyResultListForms) *ListTodoWorkRecordsResponseBodyResultList {
	s.Forms = v
	return s
}

func (s *ListTodoWorkRecordsResponseBodyResultList) SetInstanceId(v string) *ListTodoWorkRecordsResponseBodyResultList {
	s.InstanceId = &v
	return s
}

func (s *ListTodoWorkRecordsResponseBodyResultList) SetTaskId(v int64) *ListTodoWorkRecordsResponseBodyResultList {
	s.TaskId = &v
	return s
}

func (s *ListTodoWorkRecordsResponseBodyResultList) SetTitle(v string) *ListTodoWorkRecordsResponseBodyResultList {
	s.Title = &v
	return s
}

func (s *ListTodoWorkRecordsResponseBodyResultList) SetUrl(v string) *ListTodoWorkRecordsResponseBodyResultList {
	s.Url = &v
	return s
}

type ListTodoWorkRecordsResponseBodyResultListForms struct {
	// example:
	//
	// 钉三多
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 入职员工姓名
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListTodoWorkRecordsResponseBodyResultListForms) String() string {
	return tea.Prettify(s)
}

func (s ListTodoWorkRecordsResponseBodyResultListForms) GoString() string {
	return s.String()
}

func (s *ListTodoWorkRecordsResponseBodyResultListForms) SetContent(v string) *ListTodoWorkRecordsResponseBodyResultListForms {
	s.Content = &v
	return s
}

func (s *ListTodoWorkRecordsResponseBodyResultListForms) SetTitle(v string) *ListTodoWorkRecordsResponseBodyResultListForms {
	s.Title = &v
	return s
}

type ListTodoWorkRecordsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTodoWorkRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTodoWorkRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTodoWorkRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListTodoWorkRecordsResponse) SetHeaders(v map[string]*string) *ListTodoWorkRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListTodoWorkRecordsResponse) SetStatusCode(v int32) *ListTodoWorkRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTodoWorkRecordsResponse) SetBody(v *ListTodoWorkRecordsResponseBody) *ListTodoWorkRecordsResponse {
	s.Body = v
	return s
}

type ListUserVisibleBpmsProcessesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s ListUserVisibleBpmsProcessesHeaders) String() string {
	return tea.Prettify(s)
}

func (s ListUserVisibleBpmsProcessesHeaders) GoString() string {
	return s.String()
}

func (s *ListUserVisibleBpmsProcessesHeaders) SetCommonHeaders(v map[string]*string) *ListUserVisibleBpmsProcessesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListUserVisibleBpmsProcessesHeaders) SetXAcsDingtalkAccessToken(v string) *ListUserVisibleBpmsProcessesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type ListUserVisibleBpmsProcessesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	NextToken *int64 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// manager7078
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListUserVisibleBpmsProcessesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserVisibleBpmsProcessesRequest) GoString() string {
	return s.String()
}

func (s *ListUserVisibleBpmsProcessesRequest) SetMaxResults(v int64) *ListUserVisibleBpmsProcessesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUserVisibleBpmsProcessesRequest) SetNextToken(v int64) *ListUserVisibleBpmsProcessesRequest {
	s.NextToken = &v
	return s
}

func (s *ListUserVisibleBpmsProcessesRequest) SetUserId(v string) *ListUserVisibleBpmsProcessesRequest {
	s.UserId = &v
	return s
}

type ListUserVisibleBpmsProcessesResponseBody struct {
	Result *ListUserVisibleBpmsProcessesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ListUserVisibleBpmsProcessesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserVisibleBpmsProcessesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserVisibleBpmsProcessesResponseBody) SetResult(v *ListUserVisibleBpmsProcessesResponseBodyResult) *ListUserVisibleBpmsProcessesResponseBody {
	s.Result = v
	return s
}

type ListUserVisibleBpmsProcessesResponseBodyResult struct {
	// example:
	//
	// 10
	NextToken   *int64                                                       `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ProcessList []*ListUserVisibleBpmsProcessesResponseBodyResultProcessList `json:"processList,omitempty" xml:"processList,omitempty" type:"Repeated"`
}

func (s ListUserVisibleBpmsProcessesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListUserVisibleBpmsProcessesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListUserVisibleBpmsProcessesResponseBodyResult) SetNextToken(v int64) *ListUserVisibleBpmsProcessesResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *ListUserVisibleBpmsProcessesResponseBodyResult) SetProcessList(v []*ListUserVisibleBpmsProcessesResponseBodyResultProcessList) *ListUserVisibleBpmsProcessesResponseBodyResult {
	s.ProcessList = v
	return s
}

type ListUserVisibleBpmsProcessesResponseBodyResultProcessList struct {
	// example:
	//
	// 12347899
	DirId *string `json:"dirId,omitempty" xml:"dirId,omitempty"`
	// example:
	//
	// 财务管理
	DirName *string `json:"dirName,omitempty" xml:"dirName,omitempty"`
	// example:
	//
	// https://gw.xxxx/T-102-102.png
	IconUrl *string `json:"iconUrl,omitempty" xml:"iconUrl,omitempty"`
	// example:
	//
	// 物品领用
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// PROC-YMLA1-xxxx-11WFJ-1
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// example:
	//
	// https://aflow.dingtalk.com/dingtalk/mobile/homepage.htm?xxxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListUserVisibleBpmsProcessesResponseBodyResultProcessList) String() string {
	return tea.Prettify(s)
}

func (s ListUserVisibleBpmsProcessesResponseBodyResultProcessList) GoString() string {
	return s.String()
}

func (s *ListUserVisibleBpmsProcessesResponseBodyResultProcessList) SetDirId(v string) *ListUserVisibleBpmsProcessesResponseBodyResultProcessList {
	s.DirId = &v
	return s
}

func (s *ListUserVisibleBpmsProcessesResponseBodyResultProcessList) SetDirName(v string) *ListUserVisibleBpmsProcessesResponseBodyResultProcessList {
	s.DirName = &v
	return s
}

func (s *ListUserVisibleBpmsProcessesResponseBodyResultProcessList) SetIconUrl(v string) *ListUserVisibleBpmsProcessesResponseBodyResultProcessList {
	s.IconUrl = &v
	return s
}

func (s *ListUserVisibleBpmsProcessesResponseBodyResultProcessList) SetName(v string) *ListUserVisibleBpmsProcessesResponseBodyResultProcessList {
	s.Name = &v
	return s
}

func (s *ListUserVisibleBpmsProcessesResponseBodyResultProcessList) SetProcessCode(v string) *ListUserVisibleBpmsProcessesResponseBodyResultProcessList {
	s.ProcessCode = &v
	return s
}

func (s *ListUserVisibleBpmsProcessesResponseBodyResultProcessList) SetUrl(v string) *ListUserVisibleBpmsProcessesResponseBodyResultProcessList {
	s.Url = &v
	return s
}

type ListUserVisibleBpmsProcessesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserVisibleBpmsProcessesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserVisibleBpmsProcessesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserVisibleBpmsProcessesResponse) GoString() string {
	return s.String()
}

func (s *ListUserVisibleBpmsProcessesResponse) SetHeaders(v map[string]*string) *ListUserVisibleBpmsProcessesResponse {
	s.Headers = v
	return s
}

func (s *ListUserVisibleBpmsProcessesResponse) SetStatusCode(v int32) *ListUserVisibleBpmsProcessesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserVisibleBpmsProcessesResponse) SetBody(v *ListUserVisibleBpmsProcessesResponseBody) *ListUserVisibleBpmsProcessesResponse {
	s.Body = v
	return s
}

type PagesExportInstancesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PagesExportInstancesHeaders) String() string {
	return tea.Prettify(s)
}

func (s PagesExportInstancesHeaders) GoString() string {
	return s.String()
}

func (s *PagesExportInstancesHeaders) SetCommonHeaders(v map[string]*string) *PagesExportInstancesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PagesExportInstancesHeaders) SetXAcsDingtalkAccessToken(v string) *PagesExportInstancesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PagesExportInstancesRequest struct {
	EndTimeInMills *int64 `json:"endTimeInMills,omitempty" xml:"endTimeInMills,omitempty"`
	// This parameter is required.
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OrderBy   *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// This parameter is required.
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// This parameter is required.
	StartTimeInMills *int64  `json:"startTimeInMills,omitempty" xml:"startTimeInMills,omitempty"`
	Status           *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PagesExportInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s PagesExportInstancesRequest) GoString() string {
	return s.String()
}

func (s *PagesExportInstancesRequest) SetEndTimeInMills(v int64) *PagesExportInstancesRequest {
	s.EndTimeInMills = &v
	return s
}

func (s *PagesExportInstancesRequest) SetMaxResults(v int32) *PagesExportInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *PagesExportInstancesRequest) SetNextToken(v string) *PagesExportInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *PagesExportInstancesRequest) SetOrderBy(v string) *PagesExportInstancesRequest {
	s.OrderBy = &v
	return s
}

func (s *PagesExportInstancesRequest) SetProcessCode(v string) *PagesExportInstancesRequest {
	s.ProcessCode = &v
	return s
}

func (s *PagesExportInstancesRequest) SetStartTimeInMills(v int64) *PagesExportInstancesRequest {
	s.StartTimeInMills = &v
	return s
}

func (s *PagesExportInstancesRequest) SetStatus(v string) *PagesExportInstancesRequest {
	s.Status = &v
	return s
}

type PagesExportInstancesResponseBody struct {
	Result *PagesExportInstancesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s PagesExportInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PagesExportInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *PagesExportInstancesResponseBody) SetResult(v *PagesExportInstancesResponseBodyResult) *PagesExportInstancesResponseBody {
	s.Result = v
	return s
}

type PagesExportInstancesResponseBodyResult struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	HasMore *bool                                         `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List    []*PagesExportInstancesResponseBodyResultList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s PagesExportInstancesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PagesExportInstancesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PagesExportInstancesResponseBodyResult) SetHasMore(v bool) *PagesExportInstancesResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResult) SetList(v []*PagesExportInstancesResponseBodyResultList) *PagesExportInstancesResponseBodyResult {
	s.List = v
	return s
}

func (s *PagesExportInstancesResponseBodyResult) SetNextToken(v string) *PagesExportInstancesResponseBodyResult {
	s.NextToken = &v
	return s
}

type PagesExportInstancesResponseBodyResultList struct {
	// This parameter is required.
	//
	// example:
	//
	// cdef-dae2fd2-example
	AttachedProcessInstanceIds *string `json:"attachedProcessInstanceIds,omitempty" xml:"attachedProcessInstanceIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 202110111558000355024
	BusinessId *string `json:"businessId,omitempty" xml:"businessId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1635165470201
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1633795200000
	FinishTime *int64 `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// This parameter is required.
	FormComponentValues []*PagesExportInstancesResponseBodyResultListFormComponentValues `json:"formComponentValues,omitempty" xml:"formComponentValues,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// dcdse-dae2fd2-example
	MainProcessInstanceId *string                                                       `json:"mainProcessInstanceId,omitempty" xml:"mainProcessInstanceId,omitempty"`
	OperationRecords      []*PagesExportInstancesResponseBodyResultListOperationRecords `json:"operationRecords,omitempty" xml:"operationRecords,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 默认-1，企业根部门
	OriginatorDeptId *string `json:"originatorDeptId,omitempty" xml:"originatorDeptId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// staff1234
	OriginatorUserid *string `json:"originatorUserid,omitempty" xml:"originatorUserid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abcdse-dse-example
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AGREE同意，REFUSE拒绝
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RUNNING审批中、TERMINATED撤销、COMPLETED审批完成、CANCELED取消
	Status *string                                            `json:"status,omitempty" xml:"status,omitempty"`
	Tasks  []*PagesExportInstancesResponseBodyResultListTasks `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 员工A提交的小肖审批单
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s PagesExportInstancesResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s PagesExportInstancesResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *PagesExportInstancesResponseBodyResultList) SetAttachedProcessInstanceIds(v string) *PagesExportInstancesResponseBodyResultList {
	s.AttachedProcessInstanceIds = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultList) SetBusinessId(v string) *PagesExportInstancesResponseBodyResultList {
	s.BusinessId = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultList) SetCreateTime(v int64) *PagesExportInstancesResponseBodyResultList {
	s.CreateTime = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultList) SetFinishTime(v int64) *PagesExportInstancesResponseBodyResultList {
	s.FinishTime = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultList) SetFormComponentValues(v []*PagesExportInstancesResponseBodyResultListFormComponentValues) *PagesExportInstancesResponseBodyResultList {
	s.FormComponentValues = v
	return s
}

func (s *PagesExportInstancesResponseBodyResultList) SetMainProcessInstanceId(v string) *PagesExportInstancesResponseBodyResultList {
	s.MainProcessInstanceId = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultList) SetOperationRecords(v []*PagesExportInstancesResponseBodyResultListOperationRecords) *PagesExportInstancesResponseBodyResultList {
	s.OperationRecords = v
	return s
}

func (s *PagesExportInstancesResponseBodyResultList) SetOriginatorDeptId(v string) *PagesExportInstancesResponseBodyResultList {
	s.OriginatorDeptId = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultList) SetOriginatorUserid(v string) *PagesExportInstancesResponseBodyResultList {
	s.OriginatorUserid = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultList) SetProcessInstanceId(v string) *PagesExportInstancesResponseBodyResultList {
	s.ProcessInstanceId = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultList) SetResult(v string) *PagesExportInstancesResponseBodyResultList {
	s.Result = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultList) SetStatus(v string) *PagesExportInstancesResponseBodyResultList {
	s.Status = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultList) SetTasks(v []*PagesExportInstancesResponseBodyResultListTasks) *PagesExportInstancesResponseBodyResultList {
	s.Tasks = v
	return s
}

func (s *PagesExportInstancesResponseBodyResultList) SetTitle(v string) *PagesExportInstancesResponseBodyResultList {
	s.Title = &v
	return s
}

type PagesExportInstancesResponseBodyResultListFormComponentValues struct {
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// example:
	//
	// {"staffId":"abcd"}
	ExtValue *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TextField-a32bcdef
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 张三
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s PagesExportInstancesResponseBodyResultListFormComponentValues) String() string {
	return tea.Prettify(s)
}

func (s PagesExportInstancesResponseBodyResultListFormComponentValues) GoString() string {
	return s.String()
}

func (s *PagesExportInstancesResponseBodyResultListFormComponentValues) SetComponentName(v string) *PagesExportInstancesResponseBodyResultListFormComponentValues {
	s.ComponentName = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultListFormComponentValues) SetExtValue(v string) *PagesExportInstancesResponseBodyResultListFormComponentValues {
	s.ExtValue = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultListFormComponentValues) SetId(v string) *PagesExportInstancesResponseBodyResultListFormComponentValues {
	s.Id = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultListFormComponentValues) SetName(v string) *PagesExportInstancesResponseBodyResultListFormComponentValues {
	s.Name = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultListFormComponentValues) SetValue(v string) *PagesExportInstancesResponseBodyResultListFormComponentValues {
	s.Value = &v
	return s
}

type PagesExportInstancesResponseBodyResultListOperationRecords struct {
	// example:
	//
	// 1234_abcd
	ActivityId  *string                                                                  `json:"activityId,omitempty" xml:"activityId,omitempty"`
	Attachments []*PagesExportInstancesResponseBodyResultListOperationRecordsAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// example:
	//
	// []
	Images []*string `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	// example:
	//
	// EXECUTE_TASK_NORMAL（正常执行任务），EXECUTE_TASK_AGENT（代理人执行任务），APPEND_TASK_BEFORE（前加签任务），APPEND_TASK_AFTER（后加签任务），REDIRECT_TASK（转交任务），START_PROCESS_INSTANCE（发起流程实例），TERMINATE_PROCESS_INSTANCE（终止(撤销)流程实例），FINISH_PROCESS_INSTANCE（结束流程实例），ADD_REMARK（添加评论）
	OperationType *string `json:"operationType,omitempty" xml:"operationType,omitempty"`
	// example:
	//
	// 同意
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// AGREE（同意），REFUSE（拒绝），NONE（未知）
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// 12345
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 1657522271000
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// example:
	//
	// manager1
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PagesExportInstancesResponseBodyResultListOperationRecords) String() string {
	return tea.Prettify(s)
}

func (s PagesExportInstancesResponseBodyResultListOperationRecords) GoString() string {
	return s.String()
}

func (s *PagesExportInstancesResponseBodyResultListOperationRecords) SetActivityId(v string) *PagesExportInstancesResponseBodyResultListOperationRecords {
	s.ActivityId = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultListOperationRecords) SetAttachments(v []*PagesExportInstancesResponseBodyResultListOperationRecordsAttachments) *PagesExportInstancesResponseBodyResultListOperationRecords {
	s.Attachments = v
	return s
}

func (s *PagesExportInstancesResponseBodyResultListOperationRecords) SetImages(v []*string) *PagesExportInstancesResponseBodyResultListOperationRecords {
	s.Images = v
	return s
}

func (s *PagesExportInstancesResponseBodyResultListOperationRecords) SetOperationType(v string) *PagesExportInstancesResponseBodyResultListOperationRecords {
	s.OperationType = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultListOperationRecords) SetRemark(v string) *PagesExportInstancesResponseBodyResultListOperationRecords {
	s.Remark = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultListOperationRecords) SetResult(v string) *PagesExportInstancesResponseBodyResultListOperationRecords {
	s.Result = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultListOperationRecords) SetTaskId(v int64) *PagesExportInstancesResponseBodyResultListOperationRecords {
	s.TaskId = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultListOperationRecords) SetTimestamp(v int64) *PagesExportInstancesResponseBodyResultListOperationRecords {
	s.Timestamp = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultListOperationRecords) SetUserId(v string) *PagesExportInstancesResponseBodyResultListOperationRecords {
	s.UserId = &v
	return s
}

type PagesExportInstancesResponseBodyResultListOperationRecordsAttachments struct {
	// example:
	//
	// 1234567
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// example:
	//
	// 附件
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// 123
	FileSize *string `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
}

func (s PagesExportInstancesResponseBodyResultListOperationRecordsAttachments) String() string {
	return tea.Prettify(s)
}

func (s PagesExportInstancesResponseBodyResultListOperationRecordsAttachments) GoString() string {
	return s.String()
}

func (s *PagesExportInstancesResponseBodyResultListOperationRecordsAttachments) SetFileId(v string) *PagesExportInstancesResponseBodyResultListOperationRecordsAttachments {
	s.FileId = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultListOperationRecordsAttachments) SetFileName(v string) *PagesExportInstancesResponseBodyResultListOperationRecordsAttachments {
	s.FileName = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultListOperationRecordsAttachments) SetFileSize(v string) *PagesExportInstancesResponseBodyResultListOperationRecordsAttachments {
	s.FileSize = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultListOperationRecordsAttachments) SetFileType(v string) *PagesExportInstancesResponseBodyResultListOperationRecordsAttachments {
	s.FileType = &v
	return s
}

type PagesExportInstancesResponseBodyResultListTasks struct {
	// example:
	//
	// 1234_abcd
	ActivityId *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
	// example:
	//
	// 1657522271000
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// example:
	//
	// 1657522271000
	FinishTimestamp *int64 `json:"finishTimestamp,omitempty" xml:"finishTimestamp,omitempty"`
	// example:
	//
	// 分为AGREE（同意），REFUSE（拒绝），REDIRECTED（转交）
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// NEW（未启动），RUNNING（处理中），PAUSED（暂停），CANCELED（取消），COMPLETED（完成），TERMINATED（终止）
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// staff1234
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PagesExportInstancesResponseBodyResultListTasks) String() string {
	return tea.Prettify(s)
}

func (s PagesExportInstancesResponseBodyResultListTasks) GoString() string {
	return s.String()
}

func (s *PagesExportInstancesResponseBodyResultListTasks) SetActivityId(v string) *PagesExportInstancesResponseBodyResultListTasks {
	s.ActivityId = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultListTasks) SetCreateTimestamp(v int64) *PagesExportInstancesResponseBodyResultListTasks {
	s.CreateTimestamp = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultListTasks) SetFinishTimestamp(v int64) *PagesExportInstancesResponseBodyResultListTasks {
	s.FinishTimestamp = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultListTasks) SetResult(v string) *PagesExportInstancesResponseBodyResultListTasks {
	s.Result = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultListTasks) SetStatus(v string) *PagesExportInstancesResponseBodyResultListTasks {
	s.Status = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultListTasks) SetTaskId(v int64) *PagesExportInstancesResponseBodyResultListTasks {
	s.TaskId = &v
	return s
}

func (s *PagesExportInstancesResponseBodyResultListTasks) SetUserId(v string) *PagesExportInstancesResponseBodyResultListTasks {
	s.UserId = &v
	return s
}

type PagesExportInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PagesExportInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PagesExportInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s PagesExportInstancesResponse) GoString() string {
	return s.String()
}

func (s *PagesExportInstancesResponse) SetHeaders(v map[string]*string) *PagesExportInstancesResponse {
	s.Headers = v
	return s
}

func (s *PagesExportInstancesResponse) SetStatusCode(v int32) *PagesExportInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *PagesExportInstancesResponse) SetBody(v *PagesExportInstancesResponseBody) *PagesExportInstancesResponse {
	s.Body = v
	return s
}

type PremiumAddApproveDentryAuthHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumAddApproveDentryAuthHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumAddApproveDentryAuthHeaders) GoString() string {
	return s.String()
}

func (s *PremiumAddApproveDentryAuthHeaders) SetCommonHeaders(v map[string]*string) *PremiumAddApproveDentryAuthHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumAddApproveDentryAuthHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumAddApproveDentryAuthHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumAddApproveDentryAuthRequest struct {
	// This parameter is required.
	FileInfos []*PremiumAddApproveDentryAuthRequestFileInfos `json:"fileInfos,omitempty" xml:"fileInfos,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// user123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PremiumAddApproveDentryAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumAddApproveDentryAuthRequest) GoString() string {
	return s.String()
}

func (s *PremiumAddApproveDentryAuthRequest) SetFileInfos(v []*PremiumAddApproveDentryAuthRequestFileInfos) *PremiumAddApproveDentryAuthRequest {
	s.FileInfos = v
	return s
}

func (s *PremiumAddApproveDentryAuthRequest) SetUserId(v string) *PremiumAddApproveDentryAuthRequest {
	s.UserId = &v
	return s
}

type PremiumAddApproveDentryAuthRequestFileInfos struct {
	// This parameter is required.
	//
	// example:
	//
	// B1oQixxxx
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111
	SpaceId *int64 `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s PremiumAddApproveDentryAuthRequestFileInfos) String() string {
	return tea.Prettify(s)
}

func (s PremiumAddApproveDentryAuthRequestFileInfos) GoString() string {
	return s.String()
}

func (s *PremiumAddApproveDentryAuthRequestFileInfos) SetFileId(v string) *PremiumAddApproveDentryAuthRequestFileInfos {
	s.FileId = &v
	return s
}

func (s *PremiumAddApproveDentryAuthRequestFileInfos) SetSpaceId(v int64) *PremiumAddApproveDentryAuthRequestFileInfos {
	s.SpaceId = &v
	return s
}

type PremiumAddApproveDentryAuthResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PremiumAddApproveDentryAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumAddApproveDentryAuthResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumAddApproveDentryAuthResponseBody) SetResult(v bool) *PremiumAddApproveDentryAuthResponseBody {
	s.Result = &v
	return s
}

func (s *PremiumAddApproveDentryAuthResponseBody) SetSuccess(v bool) *PremiumAddApproveDentryAuthResponseBody {
	s.Success = &v
	return s
}

type PremiumAddApproveDentryAuthResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumAddApproveDentryAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumAddApproveDentryAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumAddApproveDentryAuthResponse) GoString() string {
	return s.String()
}

func (s *PremiumAddApproveDentryAuthResponse) SetHeaders(v map[string]*string) *PremiumAddApproveDentryAuthResponse {
	s.Headers = v
	return s
}

func (s *PremiumAddApproveDentryAuthResponse) SetStatusCode(v int32) *PremiumAddApproveDentryAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumAddApproveDentryAuthResponse) SetBody(v *PremiumAddApproveDentryAuthResponseBody) *PremiumAddApproveDentryAuthResponse {
	s.Body = v
	return s
}

type PremiumAppendTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumAppendTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumAppendTaskHeaders) GoString() string {
	return s.String()
}

func (s *PremiumAppendTaskHeaders) SetCommonHeaders(v map[string]*string) *PremiumAppendTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumAppendTaskHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumAppendTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumAppendTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ALL
	ActivateType *string `json:"activateType,omitempty" xml:"activateType,omitempty"`
	AgreeAll     *bool   `json:"agreeAll,omitempty" xml:"agreeAll,omitempty"`
	// This parameter is required.
	AppenderUserIds []*string `json:"appenderUserIds,omitempty" xml:"appenderUserIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// manager001
	OperateUserId *string `json:"operateUserId,omitempty" xml:"operateUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// processInstanceId123
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// example:
	//
	// 请XX帮忙审批一下
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// after
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PremiumAppendTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumAppendTaskRequest) GoString() string {
	return s.String()
}

func (s *PremiumAppendTaskRequest) SetActivateType(v string) *PremiumAppendTaskRequest {
	s.ActivateType = &v
	return s
}

func (s *PremiumAppendTaskRequest) SetAgreeAll(v bool) *PremiumAppendTaskRequest {
	s.AgreeAll = &v
	return s
}

func (s *PremiumAppendTaskRequest) SetAppenderUserIds(v []*string) *PremiumAppendTaskRequest {
	s.AppenderUserIds = v
	return s
}

func (s *PremiumAppendTaskRequest) SetOperateUserId(v string) *PremiumAppendTaskRequest {
	s.OperateUserId = &v
	return s
}

func (s *PremiumAppendTaskRequest) SetProcessInstanceId(v string) *PremiumAppendTaskRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *PremiumAppendTaskRequest) SetRemark(v string) *PremiumAppendTaskRequest {
	s.Remark = &v
	return s
}

func (s *PremiumAppendTaskRequest) SetTaskId(v int64) *PremiumAppendTaskRequest {
	s.TaskId = &v
	return s
}

func (s *PremiumAppendTaskRequest) SetType(v string) *PremiumAppendTaskRequest {
	s.Type = &v
	return s
}

type PremiumAppendTaskResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s PremiumAppendTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumAppendTaskResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumAppendTaskResponseBody) SetResult(v bool) *PremiumAppendTaskResponseBody {
	s.Result = &v
	return s
}

type PremiumAppendTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumAppendTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumAppendTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumAppendTaskResponse) GoString() string {
	return s.String()
}

func (s *PremiumAppendTaskResponse) SetHeaders(v map[string]*string) *PremiumAppendTaskResponse {
	s.Headers = v
	return s
}

func (s *PremiumAppendTaskResponse) SetStatusCode(v int32) *PremiumAppendTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumAppendTaskResponse) SetBody(v *PremiumAppendTaskResponseBody) *PremiumAppendTaskResponse {
	s.Body = v
	return s
}

type PremiumBatchExecuteProcessInstancesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumBatchExecuteProcessInstancesHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumBatchExecuteProcessInstancesHeaders) GoString() string {
	return s.String()
}

func (s *PremiumBatchExecuteProcessInstancesHeaders) SetCommonHeaders(v map[string]*string) *PremiumBatchExecuteProcessInstancesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumBatchExecuteProcessInstancesHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumBatchExecuteProcessInstancesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumBatchExecuteProcessInstancesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 67583405630
	ActionerUserId *string `json:"actionerUserId,omitempty" xml:"actionerUserId,omitempty"`
	Remark         *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agree
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// This parameter is required.
	TaskInfoList []*PremiumBatchExecuteProcessInstancesRequestTaskInfoList `json:"taskInfoList,omitempty" xml:"taskInfoList,omitempty" type:"Repeated"`
}

func (s PremiumBatchExecuteProcessInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumBatchExecuteProcessInstancesRequest) GoString() string {
	return s.String()
}

func (s *PremiumBatchExecuteProcessInstancesRequest) SetActionerUserId(v string) *PremiumBatchExecuteProcessInstancesRequest {
	s.ActionerUserId = &v
	return s
}

func (s *PremiumBatchExecuteProcessInstancesRequest) SetRemark(v string) *PremiumBatchExecuteProcessInstancesRequest {
	s.Remark = &v
	return s
}

func (s *PremiumBatchExecuteProcessInstancesRequest) SetResult(v string) *PremiumBatchExecuteProcessInstancesRequest {
	s.Result = &v
	return s
}

func (s *PremiumBatchExecuteProcessInstancesRequest) SetTaskInfoList(v []*PremiumBatchExecuteProcessInstancesRequestTaskInfoList) *PremiumBatchExecuteProcessInstancesRequest {
	s.TaskInfoList = v
	return s
}

type PremiumBatchExecuteProcessInstancesRequestTaskInfoList struct {
	// This parameter is required.
	//
	// example:
	//
	// a171de6c-8bxxxx
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// This parameter is required.
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s PremiumBatchExecuteProcessInstancesRequestTaskInfoList) String() string {
	return tea.Prettify(s)
}

func (s PremiumBatchExecuteProcessInstancesRequestTaskInfoList) GoString() string {
	return s.String()
}

func (s *PremiumBatchExecuteProcessInstancesRequestTaskInfoList) SetProcessInstanceId(v string) *PremiumBatchExecuteProcessInstancesRequestTaskInfoList {
	s.ProcessInstanceId = &v
	return s
}

func (s *PremiumBatchExecuteProcessInstancesRequestTaskInfoList) SetTaskId(v int64) *PremiumBatchExecuteProcessInstancesRequestTaskInfoList {
	s.TaskId = &v
	return s
}

type PremiumBatchExecuteProcessInstancesResponseBody struct {
	Result  map[string]*ResultValue `json:"result,omitempty" xml:"result,omitempty"`
	Success *bool                   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PremiumBatchExecuteProcessInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumBatchExecuteProcessInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumBatchExecuteProcessInstancesResponseBody) SetResult(v map[string]*ResultValue) *PremiumBatchExecuteProcessInstancesResponseBody {
	s.Result = v
	return s
}

func (s *PremiumBatchExecuteProcessInstancesResponseBody) SetSuccess(v bool) *PremiumBatchExecuteProcessInstancesResponseBody {
	s.Success = &v
	return s
}

type PremiumBatchExecuteProcessInstancesResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumBatchExecuteProcessInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumBatchExecuteProcessInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumBatchExecuteProcessInstancesResponse) GoString() string {
	return s.String()
}

func (s *PremiumBatchExecuteProcessInstancesResponse) SetHeaders(v map[string]*string) *PremiumBatchExecuteProcessInstancesResponse {
	s.Headers = v
	return s
}

func (s *PremiumBatchExecuteProcessInstancesResponse) SetStatusCode(v int32) *PremiumBatchExecuteProcessInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumBatchExecuteProcessInstancesResponse) SetBody(v *PremiumBatchExecuteProcessInstancesResponseBody) *PremiumBatchExecuteProcessInstancesResponse {
	s.Body = v
	return s
}

type PremiumDelDirHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumDelDirHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumDelDirHeaders) GoString() string {
	return s.String()
}

func (s *PremiumDelDirHeaders) SetCommonHeaders(v map[string]*string) *PremiumDelDirHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumDelDirHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumDelDirHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumDelDirRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// oaDirIdxxx
	DirId *string `json:"dirId,omitempty" xml:"dirId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user001
	OperateUserId *string `json:"operateUserId,omitempty" xml:"operateUserId,omitempty"`
}

func (s PremiumDelDirRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumDelDirRequest) GoString() string {
	return s.String()
}

func (s *PremiumDelDirRequest) SetDirId(v string) *PremiumDelDirRequest {
	s.DirId = &v
	return s
}

func (s *PremiumDelDirRequest) SetOperateUserId(v string) *PremiumDelDirRequest {
	s.OperateUserId = &v
	return s
}

type PremiumDelDirResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PremiumDelDirResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumDelDirResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumDelDirResponseBody) SetSuccess(v bool) *PremiumDelDirResponseBody {
	s.Success = &v
	return s
}

type PremiumDelDirResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumDelDirResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumDelDirResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumDelDirResponse) GoString() string {
	return s.String()
}

func (s *PremiumDelDirResponse) SetHeaders(v map[string]*string) *PremiumDelDirResponse {
	s.Headers = v
	return s
}

func (s *PremiumDelDirResponse) SetStatusCode(v int32) *PremiumDelDirResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumDelDirResponse) SetBody(v *PremiumDelDirResponseBody) *PremiumDelDirResponse {
	s.Body = v
	return s
}

type PremiumDeleteFormInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumDeleteFormInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumDeleteFormInstanceHeaders) GoString() string {
	return s.String()
}

func (s *PremiumDeleteFormInstanceHeaders) SetCommonHeaders(v map[string]*string) *PremiumDeleteFormInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumDeleteFormInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumDeleteFormInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumDeleteFormInstanceRequest struct {
	FormInstanceIds []*string `json:"formInstanceIds,omitempty" xml:"formInstanceIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-EF6YJL35P2-SCKICSB7P750S0YISYKV3-xxxx-1
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	UserId      *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PremiumDeleteFormInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumDeleteFormInstanceRequest) GoString() string {
	return s.String()
}

func (s *PremiumDeleteFormInstanceRequest) SetFormInstanceIds(v []*string) *PremiumDeleteFormInstanceRequest {
	s.FormInstanceIds = v
	return s
}

func (s *PremiumDeleteFormInstanceRequest) SetProcessCode(v string) *PremiumDeleteFormInstanceRequest {
	s.ProcessCode = &v
	return s
}

func (s *PremiumDeleteFormInstanceRequest) SetUserId(v string) *PremiumDeleteFormInstanceRequest {
	s.UserId = &v
	return s
}

type PremiumDeleteFormInstanceResponseBody struct {
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PremiumDeleteFormInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumDeleteFormInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumDeleteFormInstanceResponseBody) SetSuccess(v string) *PremiumDeleteFormInstanceResponseBody {
	s.Success = &v
	return s
}

type PremiumDeleteFormInstanceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumDeleteFormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumDeleteFormInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumDeleteFormInstanceResponse) GoString() string {
	return s.String()
}

func (s *PremiumDeleteFormInstanceResponse) SetHeaders(v map[string]*string) *PremiumDeleteFormInstanceResponse {
	s.Headers = v
	return s
}

func (s *PremiumDeleteFormInstanceResponse) SetStatusCode(v int32) *PremiumDeleteFormInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumDeleteFormInstanceResponse) SetBody(v *PremiumDeleteFormInstanceResponseBody) *PremiumDeleteFormInstanceResponse {
	s.Body = v
	return s
}

type PremiumGetAttachmentSpaceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumGetAttachmentSpaceHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetAttachmentSpaceHeaders) GoString() string {
	return s.String()
}

func (s *PremiumGetAttachmentSpaceHeaders) SetCommonHeaders(v map[string]*string) *PremiumGetAttachmentSpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumGetAttachmentSpaceHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumGetAttachmentSpaceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumGetAttachmentSpaceRequest struct {
	// example:
	//
	// 8345000
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PremiumGetAttachmentSpaceRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetAttachmentSpaceRequest) GoString() string {
	return s.String()
}

func (s *PremiumGetAttachmentSpaceRequest) SetAgentId(v int64) *PremiumGetAttachmentSpaceRequest {
	s.AgentId = &v
	return s
}

func (s *PremiumGetAttachmentSpaceRequest) SetUserId(v string) *PremiumGetAttachmentSpaceRequest {
	s.UserId = &v
	return s
}

type PremiumGetAttachmentSpaceResponseBody struct {
	Result *PremiumGetAttachmentSpaceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PremiumGetAttachmentSpaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetAttachmentSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumGetAttachmentSpaceResponseBody) SetResult(v *PremiumGetAttachmentSpaceResponseBodyResult) *PremiumGetAttachmentSpaceResponseBody {
	s.Result = v
	return s
}

func (s *PremiumGetAttachmentSpaceResponseBody) SetSuccess(v bool) *PremiumGetAttachmentSpaceResponseBody {
	s.Success = &v
	return s
}

type PremiumGetAttachmentSpaceResponseBodyResult struct {
	// example:
	//
	// 3996960664
	SpaceId *int64 `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s PremiumGetAttachmentSpaceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetAttachmentSpaceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PremiumGetAttachmentSpaceResponseBodyResult) SetSpaceId(v int64) *PremiumGetAttachmentSpaceResponseBodyResult {
	s.SpaceId = &v
	return s
}

type PremiumGetAttachmentSpaceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumGetAttachmentSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumGetAttachmentSpaceResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetAttachmentSpaceResponse) GoString() string {
	return s.String()
}

func (s *PremiumGetAttachmentSpaceResponse) SetHeaders(v map[string]*string) *PremiumGetAttachmentSpaceResponse {
	s.Headers = v
	return s
}

func (s *PremiumGetAttachmentSpaceResponse) SetStatusCode(v int32) *PremiumGetAttachmentSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumGetAttachmentSpaceResponse) SetBody(v *PremiumGetAttachmentSpaceResponseBody) *PremiumGetAttachmentSpaceResponse {
	s.Body = v
	return s
}

type PremiumGetDoneTasksHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumGetDoneTasksHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetDoneTasksHeaders) GoString() string {
	return s.String()
}

func (s *PremiumGetDoneTasksHeaders) SetCommonHeaders(v map[string]*string) *PremiumGetDoneTasksHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumGetDoneTasksHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumGetDoneTasksHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumGetDoneTasksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// manager001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PremiumGetDoneTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetDoneTasksRequest) GoString() string {
	return s.String()
}

func (s *PremiumGetDoneTasksRequest) SetPageNumber(v int32) *PremiumGetDoneTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *PremiumGetDoneTasksRequest) SetPageSize(v int32) *PremiumGetDoneTasksRequest {
	s.PageSize = &v
	return s
}

func (s *PremiumGetDoneTasksRequest) SetUserId(v string) *PremiumGetDoneTasksRequest {
	s.UserId = &v
	return s
}

type PremiumGetDoneTasksResponseBody struct {
	Result  *PremiumGetDoneTasksResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PremiumGetDoneTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetDoneTasksResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumGetDoneTasksResponseBody) SetResult(v *PremiumGetDoneTasksResponseBodyResult) *PremiumGetDoneTasksResponseBody {
	s.Result = v
	return s
}

func (s *PremiumGetDoneTasksResponseBody) SetSuccess(v bool) *PremiumGetDoneTasksResponseBody {
	s.Success = &v
	return s
}

type PremiumGetDoneTasksResponseBodyResult struct {
	HasMore *bool                                        `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List    []*PremiumGetDoneTasksResponseBodyResultList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s PremiumGetDoneTasksResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetDoneTasksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PremiumGetDoneTasksResponseBodyResult) SetHasMore(v bool) *PremiumGetDoneTasksResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *PremiumGetDoneTasksResponseBodyResult) SetList(v []*PremiumGetDoneTasksResponseBodyResultList) *PremiumGetDoneTasksResponseBodyResult {
	s.List = v
	return s
}

type PremiumGetDoneTasksResponseBodyResultList struct {
	ActivityId      *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
	FormMassage     *string `json:"formMassage,omitempty" xml:"formMassage,omitempty"`
	OriginatorId    *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	OriginatorName  *string `json:"originatorName,omitempty" xml:"originatorName,omitempty"`
	OriginatorPhoto *string `json:"originatorPhoto,omitempty" xml:"originatorPhoto,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	ProcessCreateTime *string `json:"processCreateTime,omitempty" xml:"processCreateTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	ProcessEndTime    *string `json:"processEndTime,omitempty" xml:"processEndTime,omitempty"`
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	ProcessType       *int32  `json:"processType,omitempty" xml:"processType,omitempty"`
	// example:
	//
	// agree
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	Title  *string `json:"title,omitempty" xml:"title,omitempty"`
	Url    *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s PremiumGetDoneTasksResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetDoneTasksResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *PremiumGetDoneTasksResponseBodyResultList) SetActivityId(v string) *PremiumGetDoneTasksResponseBodyResultList {
	s.ActivityId = &v
	return s
}

func (s *PremiumGetDoneTasksResponseBodyResultList) SetFormMassage(v string) *PremiumGetDoneTasksResponseBodyResultList {
	s.FormMassage = &v
	return s
}

func (s *PremiumGetDoneTasksResponseBodyResultList) SetOriginatorId(v string) *PremiumGetDoneTasksResponseBodyResultList {
	s.OriginatorId = &v
	return s
}

func (s *PremiumGetDoneTasksResponseBodyResultList) SetOriginatorName(v string) *PremiumGetDoneTasksResponseBodyResultList {
	s.OriginatorName = &v
	return s
}

func (s *PremiumGetDoneTasksResponseBodyResultList) SetOriginatorPhoto(v string) *PremiumGetDoneTasksResponseBodyResultList {
	s.OriginatorPhoto = &v
	return s
}

func (s *PremiumGetDoneTasksResponseBodyResultList) SetProcessCreateTime(v string) *PremiumGetDoneTasksResponseBodyResultList {
	s.ProcessCreateTime = &v
	return s
}

func (s *PremiumGetDoneTasksResponseBodyResultList) SetProcessEndTime(v string) *PremiumGetDoneTasksResponseBodyResultList {
	s.ProcessEndTime = &v
	return s
}

func (s *PremiumGetDoneTasksResponseBodyResultList) SetProcessInstanceId(v string) *PremiumGetDoneTasksResponseBodyResultList {
	s.ProcessInstanceId = &v
	return s
}

func (s *PremiumGetDoneTasksResponseBodyResultList) SetProcessType(v int32) *PremiumGetDoneTasksResponseBodyResultList {
	s.ProcessType = &v
	return s
}

func (s *PremiumGetDoneTasksResponseBodyResultList) SetResult(v string) *PremiumGetDoneTasksResponseBodyResultList {
	s.Result = &v
	return s
}

func (s *PremiumGetDoneTasksResponseBodyResultList) SetStatus(v string) *PremiumGetDoneTasksResponseBodyResultList {
	s.Status = &v
	return s
}

func (s *PremiumGetDoneTasksResponseBodyResultList) SetTaskId(v string) *PremiumGetDoneTasksResponseBodyResultList {
	s.TaskId = &v
	return s
}

func (s *PremiumGetDoneTasksResponseBodyResultList) SetTitle(v string) *PremiumGetDoneTasksResponseBodyResultList {
	s.Title = &v
	return s
}

func (s *PremiumGetDoneTasksResponseBodyResultList) SetUrl(v string) *PremiumGetDoneTasksResponseBodyResultList {
	s.Url = &v
	return s
}

type PremiumGetDoneTasksResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumGetDoneTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumGetDoneTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetDoneTasksResponse) GoString() string {
	return s.String()
}

func (s *PremiumGetDoneTasksResponse) SetHeaders(v map[string]*string) *PremiumGetDoneTasksResponse {
	s.Headers = v
	return s
}

func (s *PremiumGetDoneTasksResponse) SetStatusCode(v int32) *PremiumGetDoneTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumGetDoneTasksResponse) SetBody(v *PremiumGetDoneTasksResponseBody) *PremiumGetDoneTasksResponse {
	s.Body = v
	return s
}

type PremiumGetFieldModifiedHistoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumGetFieldModifiedHistoryHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFieldModifiedHistoryHeaders) GoString() string {
	return s.String()
}

func (s *PremiumGetFieldModifiedHistoryHeaders) SetCommonHeaders(v map[string]*string) *PremiumGetFieldModifiedHistoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumGetFieldModifiedHistoryHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumGetFieldModifiedHistoryHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumGetFieldModifiedHistoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// TextField-abcd
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// proc-FF6Y2xxxx
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
}

func (s PremiumGetFieldModifiedHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFieldModifiedHistoryRequest) GoString() string {
	return s.String()
}

func (s *PremiumGetFieldModifiedHistoryRequest) SetFieldId(v string) *PremiumGetFieldModifiedHistoryRequest {
	s.FieldId = &v
	return s
}

func (s *PremiumGetFieldModifiedHistoryRequest) SetProcessInstanceId(v string) *PremiumGetFieldModifiedHistoryRequest {
	s.ProcessInstanceId = &v
	return s
}

type PremiumGetFieldModifiedHistoryResponseBody struct {
	Result []*PremiumGetFieldModifiedHistoryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PremiumGetFieldModifiedHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFieldModifiedHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumGetFieldModifiedHistoryResponseBody) SetResult(v []*PremiumGetFieldModifiedHistoryResponseBodyResult) *PremiumGetFieldModifiedHistoryResponseBody {
	s.Result = v
	return s
}

func (s *PremiumGetFieldModifiedHistoryResponseBody) SetSuccess(v bool) *PremiumGetFieldModifiedHistoryResponseBody {
	s.Success = &v
	return s
}

type PremiumGetFieldModifiedHistoryResponseBodyResult struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2024-04-02T11:52Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// TextField-abcd
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
	// example:
	//
	// 钉钉1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// userId1
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// 从 111 修改到 222
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s PremiumGetFieldModifiedHistoryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFieldModifiedHistoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PremiumGetFieldModifiedHistoryResponseBodyResult) SetCreateTime(v string) *PremiumGetFieldModifiedHistoryResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *PremiumGetFieldModifiedHistoryResponseBodyResult) SetFieldId(v string) *PremiumGetFieldModifiedHistoryResponseBodyResult {
	s.FieldId = &v
	return s
}

func (s *PremiumGetFieldModifiedHistoryResponseBodyResult) SetName(v string) *PremiumGetFieldModifiedHistoryResponseBodyResult {
	s.Name = &v
	return s
}

func (s *PremiumGetFieldModifiedHistoryResponseBodyResult) SetUserId(v string) *PremiumGetFieldModifiedHistoryResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *PremiumGetFieldModifiedHistoryResponseBodyResult) SetValue(v string) *PremiumGetFieldModifiedHistoryResponseBodyResult {
	s.Value = &v
	return s
}

type PremiumGetFieldModifiedHistoryResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumGetFieldModifiedHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumGetFieldModifiedHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFieldModifiedHistoryResponse) GoString() string {
	return s.String()
}

func (s *PremiumGetFieldModifiedHistoryResponse) SetHeaders(v map[string]*string) *PremiumGetFieldModifiedHistoryResponse {
	s.Headers = v
	return s
}

func (s *PremiumGetFieldModifiedHistoryResponse) SetStatusCode(v int32) *PremiumGetFieldModifiedHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumGetFieldModifiedHistoryResponse) SetBody(v *PremiumGetFieldModifiedHistoryResponseBody) *PremiumGetFieldModifiedHistoryResponse {
	s.Body = v
	return s
}

type PremiumGetFormInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumGetFormInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormInstanceHeaders) GoString() string {
	return s.String()
}

func (s *PremiumGetFormInstanceHeaders) SetCommonHeaders(v map[string]*string) *PremiumGetFormInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumGetFormInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumGetFormInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumGetFormInstanceRequest struct {
	// example:
	//
	// SWAPP-dfeacds-example
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-abcdef-example
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 951a8-8828-430c-b3e-example
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
}

func (s PremiumGetFormInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormInstanceRequest) GoString() string {
	return s.String()
}

func (s *PremiumGetFormInstanceRequest) SetAppUuid(v string) *PremiumGetFormInstanceRequest {
	s.AppUuid = &v
	return s
}

func (s *PremiumGetFormInstanceRequest) SetFormCode(v string) *PremiumGetFormInstanceRequest {
	s.FormCode = &v
	return s
}

func (s *PremiumGetFormInstanceRequest) SetFormInstanceId(v string) *PremiumGetFormInstanceRequest {
	s.FormInstanceId = &v
	return s
}

type PremiumGetFormInstanceResponseBody struct {
	// example:
	//
	// SWAPP-dfeacds-example
	AppUuid    *string                `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// example:
	//
	// 1631870043000
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 00003
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-abcdef-example
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// This parameter is required.
	FormInstDataList []*PremiumGetFormInstanceResponseBodyFormInstDataList `json:"formInstDataList,omitempty" xml:"formInstDataList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 951a8-8828-430c-b3e-example
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// example:
	//
	// 000025
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// example:
	//
	// 1631870043000
	ModifyTimestamp *int64 `json:"modifyTimestamp,omitempty" xml:"modifyTimestamp,omitempty"`
	// example:
	//
	// PROC-abcdef-example
	OutBizCode *string `json:"outBizCode,omitempty" xml:"outBizCode,omitempty"`
	// example:
	//
	// 951a8-8828-430c-b3e-example
	OutInstanceId *string `json:"outInstanceId,omitempty" xml:"outInstanceId,omitempty"`
	// example:
	//
	// xxx提交的表单数据
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s PremiumGetFormInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumGetFormInstanceResponseBody) SetAppUuid(v string) *PremiumGetFormInstanceResponseBody {
	s.AppUuid = &v
	return s
}

func (s *PremiumGetFormInstanceResponseBody) SetAttributes(v map[string]interface{}) *PremiumGetFormInstanceResponseBody {
	s.Attributes = v
	return s
}

func (s *PremiumGetFormInstanceResponseBody) SetCreateTimestamp(v int64) *PremiumGetFormInstanceResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *PremiumGetFormInstanceResponseBody) SetCreator(v string) *PremiumGetFormInstanceResponseBody {
	s.Creator = &v
	return s
}

func (s *PremiumGetFormInstanceResponseBody) SetFormCode(v string) *PremiumGetFormInstanceResponseBody {
	s.FormCode = &v
	return s
}

func (s *PremiumGetFormInstanceResponseBody) SetFormInstDataList(v []*PremiumGetFormInstanceResponseBodyFormInstDataList) *PremiumGetFormInstanceResponseBody {
	s.FormInstDataList = v
	return s
}

func (s *PremiumGetFormInstanceResponseBody) SetFormInstanceId(v string) *PremiumGetFormInstanceResponseBody {
	s.FormInstanceId = &v
	return s
}

func (s *PremiumGetFormInstanceResponseBody) SetModifier(v string) *PremiumGetFormInstanceResponseBody {
	s.Modifier = &v
	return s
}

func (s *PremiumGetFormInstanceResponseBody) SetModifyTimestamp(v int64) *PremiumGetFormInstanceResponseBody {
	s.ModifyTimestamp = &v
	return s
}

func (s *PremiumGetFormInstanceResponseBody) SetOutBizCode(v string) *PremiumGetFormInstanceResponseBody {
	s.OutBizCode = &v
	return s
}

func (s *PremiumGetFormInstanceResponseBody) SetOutInstanceId(v string) *PremiumGetFormInstanceResponseBody {
	s.OutInstanceId = &v
	return s
}

func (s *PremiumGetFormInstanceResponseBody) SetTitle(v string) *PremiumGetFormInstanceResponseBody {
	s.Title = &v
	return s
}

type PremiumGetFormInstanceResponseBodyFormInstDataList struct {
	// This parameter is required.
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// This parameter is required.
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// This parameter is required.
	ExtendValue *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	// This parameter is required.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s PremiumGetFormInstanceResponseBodyFormInstDataList) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormInstanceResponseBodyFormInstDataList) GoString() string {
	return s.String()
}

func (s *PremiumGetFormInstanceResponseBodyFormInstDataList) SetBizAlias(v string) *PremiumGetFormInstanceResponseBodyFormInstDataList {
	s.BizAlias = &v
	return s
}

func (s *PremiumGetFormInstanceResponseBodyFormInstDataList) SetComponentType(v string) *PremiumGetFormInstanceResponseBodyFormInstDataList {
	s.ComponentType = &v
	return s
}

func (s *PremiumGetFormInstanceResponseBodyFormInstDataList) SetExtendValue(v string) *PremiumGetFormInstanceResponseBodyFormInstDataList {
	s.ExtendValue = &v
	return s
}

func (s *PremiumGetFormInstanceResponseBodyFormInstDataList) SetKey(v string) *PremiumGetFormInstanceResponseBodyFormInstDataList {
	s.Key = &v
	return s
}

func (s *PremiumGetFormInstanceResponseBodyFormInstDataList) SetLabel(v string) *PremiumGetFormInstanceResponseBodyFormInstDataList {
	s.Label = &v
	return s
}

func (s *PremiumGetFormInstanceResponseBodyFormInstDataList) SetValue(v string) *PremiumGetFormInstanceResponseBodyFormInstDataList {
	s.Value = &v
	return s
}

type PremiumGetFormInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumGetFormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumGetFormInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormInstanceResponse) GoString() string {
	return s.String()
}

func (s *PremiumGetFormInstanceResponse) SetHeaders(v map[string]*string) *PremiumGetFormInstanceResponse {
	s.Headers = v
	return s
}

func (s *PremiumGetFormInstanceResponse) SetStatusCode(v int32) *PremiumGetFormInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumGetFormInstanceResponse) SetBody(v *PremiumGetFormInstanceResponseBody) *PremiumGetFormInstanceResponse {
	s.Body = v
	return s
}

type PremiumGetFormInstancesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumGetFormInstancesHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormInstancesHeaders) GoString() string {
	return s.String()
}

func (s *PremiumGetFormInstancesHeaders) SetCommonHeaders(v map[string]*string) *PremiumGetFormInstancesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumGetFormInstancesHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumGetFormInstancesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumGetFormInstancesRequest struct {
	// example:
	//
	// SWAPP-dacdsa-example
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-daccea-example
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 100010
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s PremiumGetFormInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormInstancesRequest) GoString() string {
	return s.String()
}

func (s *PremiumGetFormInstancesRequest) SetAppUuid(v string) *PremiumGetFormInstancesRequest {
	s.AppUuid = &v
	return s
}

func (s *PremiumGetFormInstancesRequest) SetFormCode(v string) *PremiumGetFormInstancesRequest {
	s.FormCode = &v
	return s
}

func (s *PremiumGetFormInstancesRequest) SetMaxResults(v int32) *PremiumGetFormInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *PremiumGetFormInstancesRequest) SetNextToken(v string) *PremiumGetFormInstancesRequest {
	s.NextToken = &v
	return s
}

type PremiumGetFormInstancesResponseBody struct {
	// This parameter is required.
	Result *PremiumGetFormInstancesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s PremiumGetFormInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumGetFormInstancesResponseBody) SetResult(v *PremiumGetFormInstancesResponseBodyResult) *PremiumGetFormInstancesResponseBody {
	s.Result = v
	return s
}

type PremiumGetFormInstancesResponseBodyResult struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	Values []*PremiumGetFormInstancesResponseBodyResultValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s PremiumGetFormInstancesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormInstancesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PremiumGetFormInstancesResponseBodyResult) SetHasMore(v bool) *PremiumGetFormInstancesResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *PremiumGetFormInstancesResponseBodyResult) SetMaxResults(v int64) *PremiumGetFormInstancesResponseBodyResult {
	s.MaxResults = &v
	return s
}

func (s *PremiumGetFormInstancesResponseBodyResult) SetNextToken(v string) *PremiumGetFormInstancesResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *PremiumGetFormInstancesResponseBodyResult) SetValues(v []*PremiumGetFormInstancesResponseBodyResultValues) *PremiumGetFormInstancesResponseBodyResult {
	s.Values = v
	return s
}

type PremiumGetFormInstancesResponseBodyResultValues struct {
	// example:
	//
	// SWAPP-abcd-example
	AppUuid    *string                `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1635151039000
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30314512
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-abcd-example
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// This parameter is required.
	FormInstDataList []*PremiumGetFormInstancesResponseBodyResultValuesFormInstDataList `json:"formInstDataList,omitempty" xml:"formInstDataList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// abcd-eaf-acde12f
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 032142312
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// example:
	//
	// 1635151039000
	ModifyTimestamp *int64 `json:"modifyTimestamp,omitempty" xml:"modifyTimestamp,omitempty"`
	// example:
	//
	// abcd
	OutBizCode *string `json:"outBizCode,omitempty" xml:"outBizCode,omitempty"`
	// example:
	//
	// 323
	OutInstanceId *string `json:"outInstanceId,omitempty" xml:"outInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx提交的数据
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s PremiumGetFormInstancesResponseBodyResultValues) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormInstancesResponseBodyResultValues) GoString() string {
	return s.String()
}

func (s *PremiumGetFormInstancesResponseBodyResultValues) SetAppUuid(v string) *PremiumGetFormInstancesResponseBodyResultValues {
	s.AppUuid = &v
	return s
}

func (s *PremiumGetFormInstancesResponseBodyResultValues) SetAttributes(v map[string]interface{}) *PremiumGetFormInstancesResponseBodyResultValues {
	s.Attributes = v
	return s
}

func (s *PremiumGetFormInstancesResponseBodyResultValues) SetCreateTimestamp(v int64) *PremiumGetFormInstancesResponseBodyResultValues {
	s.CreateTimestamp = &v
	return s
}

func (s *PremiumGetFormInstancesResponseBodyResultValues) SetCreator(v string) *PremiumGetFormInstancesResponseBodyResultValues {
	s.Creator = &v
	return s
}

func (s *PremiumGetFormInstancesResponseBodyResultValues) SetFormCode(v string) *PremiumGetFormInstancesResponseBodyResultValues {
	s.FormCode = &v
	return s
}

func (s *PremiumGetFormInstancesResponseBodyResultValues) SetFormInstDataList(v []*PremiumGetFormInstancesResponseBodyResultValuesFormInstDataList) *PremiumGetFormInstancesResponseBodyResultValues {
	s.FormInstDataList = v
	return s
}

func (s *PremiumGetFormInstancesResponseBodyResultValues) SetFormInstanceId(v string) *PremiumGetFormInstancesResponseBodyResultValues {
	s.FormInstanceId = &v
	return s
}

func (s *PremiumGetFormInstancesResponseBodyResultValues) SetModifier(v string) *PremiumGetFormInstancesResponseBodyResultValues {
	s.Modifier = &v
	return s
}

func (s *PremiumGetFormInstancesResponseBodyResultValues) SetModifyTimestamp(v int64) *PremiumGetFormInstancesResponseBodyResultValues {
	s.ModifyTimestamp = &v
	return s
}

func (s *PremiumGetFormInstancesResponseBodyResultValues) SetOutBizCode(v string) *PremiumGetFormInstancesResponseBodyResultValues {
	s.OutBizCode = &v
	return s
}

func (s *PremiumGetFormInstancesResponseBodyResultValues) SetOutInstanceId(v string) *PremiumGetFormInstancesResponseBodyResultValues {
	s.OutInstanceId = &v
	return s
}

func (s *PremiumGetFormInstancesResponseBodyResultValues) SetTitle(v string) *PremiumGetFormInstancesResponseBodyResultValues {
	s.Title = &v
	return s
}

type PremiumGetFormInstancesResponseBodyResultValuesFormInstDataList struct {
	// example:
	//
	// staff_name
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 具体参见审批控件列表
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// example:
	//
	// {"key":"value}
	ExtendValue *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TextField-abcdefg
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 员工姓名
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 张三
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s PremiumGetFormInstancesResponseBodyResultValuesFormInstDataList) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormInstancesResponseBodyResultValuesFormInstDataList) GoString() string {
	return s.String()
}

func (s *PremiumGetFormInstancesResponseBodyResultValuesFormInstDataList) SetBizAlias(v string) *PremiumGetFormInstancesResponseBodyResultValuesFormInstDataList {
	s.BizAlias = &v
	return s
}

func (s *PremiumGetFormInstancesResponseBodyResultValuesFormInstDataList) SetComponentType(v string) *PremiumGetFormInstancesResponseBodyResultValuesFormInstDataList {
	s.ComponentType = &v
	return s
}

func (s *PremiumGetFormInstancesResponseBodyResultValuesFormInstDataList) SetExtendValue(v string) *PremiumGetFormInstancesResponseBodyResultValuesFormInstDataList {
	s.ExtendValue = &v
	return s
}

func (s *PremiumGetFormInstancesResponseBodyResultValuesFormInstDataList) SetKey(v string) *PremiumGetFormInstancesResponseBodyResultValuesFormInstDataList {
	s.Key = &v
	return s
}

func (s *PremiumGetFormInstancesResponseBodyResultValuesFormInstDataList) SetLabel(v string) *PremiumGetFormInstancesResponseBodyResultValuesFormInstDataList {
	s.Label = &v
	return s
}

func (s *PremiumGetFormInstancesResponseBodyResultValuesFormInstDataList) SetValue(v string) *PremiumGetFormInstancesResponseBodyResultValuesFormInstDataList {
	s.Value = &v
	return s
}

type PremiumGetFormInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumGetFormInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumGetFormInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormInstancesResponse) GoString() string {
	return s.String()
}

func (s *PremiumGetFormInstancesResponse) SetHeaders(v map[string]*string) *PremiumGetFormInstancesResponse {
	s.Headers = v
	return s
}

func (s *PremiumGetFormInstancesResponse) SetStatusCode(v int32) *PremiumGetFormInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumGetFormInstancesResponse) SetBody(v *PremiumGetFormInstancesResponseBody) *PremiumGetFormInstancesResponse {
	s.Body = v
	return s
}

type PremiumGetFormSchemaHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumGetFormSchemaHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormSchemaHeaders) GoString() string {
	return s.String()
}

func (s *PremiumGetFormSchemaHeaders) SetCommonHeaders(v map[string]*string) *PremiumGetFormSchemaHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumGetFormSchemaHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumGetFormSchemaHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumGetFormSchemaRequest struct {
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-17428B8C-6C60-xxxx-924C-64F1037AE067
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
}

func (s PremiumGetFormSchemaRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormSchemaRequest) GoString() string {
	return s.String()
}

func (s *PremiumGetFormSchemaRequest) SetAppUuid(v string) *PremiumGetFormSchemaRequest {
	s.AppUuid = &v
	return s
}

func (s *PremiumGetFormSchemaRequest) SetProcessCode(v string) *PremiumGetFormSchemaRequest {
	s.ProcessCode = &v
	return s
}

type PremiumGetFormSchemaResponseBody struct {
	// This parameter is required.
	Result *PremiumGetFormSchemaResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s PremiumGetFormSchemaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumGetFormSchemaResponseBody) SetResult(v *PremiumGetFormSchemaResponseBodyResult) *PremiumGetFormSchemaResponseBody {
	s.Result = v
	return s
}

type PremiumGetFormSchemaResponseBodyResult struct {
	// This parameter is required.
	//
	// example:
	//
	// 0
	AppType *int32 `json:"appType,omitempty" xml:"appType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 26652461xxxx5992
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-17428B8C-6C60-470E-xxxx-64F1037AE067
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-12-01T10:49Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-12-01T10:49Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// null
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// xxxx
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// example:
	//
	// 示例模板
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	SchemaContent *PremiumGetFormSchemaResponseBodyResultSchemaContent `json:"schemaContent,omitempty" xml:"schemaContent,omitempty" type:"Struct"`
	// example:
	//
	// PUBLISHED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PremiumGetFormSchemaResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormSchemaResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PremiumGetFormSchemaResponseBodyResult) SetAppType(v int32) *PremiumGetFormSchemaResponseBodyResult {
	s.AppType = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResult) SetCreatorUserId(v string) *PremiumGetFormSchemaResponseBodyResult {
	s.CreatorUserId = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResult) SetFormCode(v string) *PremiumGetFormSchemaResponseBodyResult {
	s.FormCode = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResult) SetGmtCreate(v string) *PremiumGetFormSchemaResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResult) SetGmtModified(v string) *PremiumGetFormSchemaResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResult) SetIcon(v string) *PremiumGetFormSchemaResponseBodyResult {
	s.Icon = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResult) SetMemo(v string) *PremiumGetFormSchemaResponseBodyResult {
	s.Memo = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResult) SetName(v string) *PremiumGetFormSchemaResponseBodyResult {
	s.Name = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResult) SetSchemaContent(v *PremiumGetFormSchemaResponseBodyResultSchemaContent) *PremiumGetFormSchemaResponseBodyResult {
	s.SchemaContent = v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResult) SetStatus(v string) *PremiumGetFormSchemaResponseBodyResult {
	s.Status = &v
	return s
}

type PremiumGetFormSchemaResponseBodyResultSchemaContent struct {
	// example:
	//
	// common
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// This parameter is required.
	Items []*PremiumGetFormSchemaResponseBodyResultSchemaContentItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 示例模板
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s PremiumGetFormSchemaResponseBodyResultSchemaContent) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormSchemaResponseBodyResultSchemaContent) GoString() string {
	return s.String()
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContent) SetIcon(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContent {
	s.Icon = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContent) SetItems(v []*PremiumGetFormSchemaResponseBodyResultSchemaContentItems) *PremiumGetFormSchemaResponseBodyResultSchemaContent {
	s.Items = v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContent) SetTitle(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContent {
	s.Title = &v
	return s
}

type PremiumGetFormSchemaResponseBodyResultSchemaContentItems struct {
	Children []*PremiumGetFormSchemaResponseBodyResultSchemaContentItemsChildren `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// TextField
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// This parameter is required.
	Props *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
}

func (s PremiumGetFormSchemaResponseBodyResultSchemaContentItems) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormSchemaResponseBodyResultSchemaContentItems) GoString() string {
	return s.String()
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItems) SetChildren(v []*PremiumGetFormSchemaResponseBodyResultSchemaContentItemsChildren) *PremiumGetFormSchemaResponseBodyResultSchemaContentItems {
	s.Children = v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItems) SetComponentName(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItems {
	s.ComponentName = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItems) SetProps(v *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) *PremiumGetFormSchemaResponseBodyResultSchemaContentItems {
	s.Props = v
	return s
}

type PremiumGetFormSchemaResponseBodyResultSchemaContentItemsChildren struct {
	// This parameter is required.
	//
	// example:
	//
	// TextField
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// This parameter is required.
	Props *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsChildrenProps `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
}

func (s PremiumGetFormSchemaResponseBodyResultSchemaContentItemsChildren) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormSchemaResponseBodyResultSchemaContentItemsChildren) GoString() string {
	return s.String()
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsChildren) SetComponentName(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsChildren {
	s.ComponentName = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsChildren) SetProps(v *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsChildrenProps) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsChildren {
	s.Props = v
	return s
}

type PremiumGetFormSchemaResponseBodyResultSchemaContentItemsChildrenProps struct {
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// This parameter is required.
	Id       *string   `json:"id,omitempty" xml:"id,omitempty"`
	Label    *string   `json:"label,omitempty" xml:"label,omitempty"`
	Options  []*string `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	Required *bool     `json:"required,omitempty" xml:"required,omitempty"`
}

func (s PremiumGetFormSchemaResponseBodyResultSchemaContentItemsChildrenProps) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormSchemaResponseBodyResultSchemaContentItemsChildrenProps) GoString() string {
	return s.String()
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsChildrenProps) SetBizAlias(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsChildrenProps {
	s.BizAlias = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsChildrenProps) SetId(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsChildrenProps {
	s.Id = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsChildrenProps) SetLabel(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsChildrenProps {
	s.Label = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsChildrenProps) SetOptions(v []*string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsChildrenProps {
	s.Options = v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsChildrenProps) SetRequired(v bool) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsChildrenProps {
	s.Required = &v
	return s
}

type PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps struct {
	// example:
	//
	// 添加
	ActionName *string `json:"actionName,omitempty" xml:"actionName,omitempty"`
	// example:
	//
	// top
	Align *string `json:"align,omitempty" xml:"align,omitempty"`
	// example:
	//
	// 1234567
	AppId *int64 `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// true
	AsyncCondition *bool `json:"asyncCondition,omitempty" xml:"asyncCondition,omitempty"`
	// example:
	//
	// 请假
	AttendTypeLabel *string                                                                         `json:"attendTypeLabel,omitempty" xml:"attendTypeLabel,omitempty"`
	BehaviorLinkage []*PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsBehaviorLinkage `json:"behaviorLinkage,omitempty" xml:"behaviorLinkage,omitempty" type:"Repeated"`
	// example:
	//
	// 我的单行输入框
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// example:
	//
	// hrm.xxxx
	BizType           *string          `json:"bizType,omitempty" xml:"bizType,omitempty"`
	ChildFieldVisible map[string]*bool `json:"childFieldVisible,omitempty" xml:"childFieldVisible,omitempty"`
	// example:
	//
	// 1
	Choice *int32 `json:"choice,omitempty" xml:"choice,omitempty"`
	// example:
	//
	// xxxx
	CommonBizType *string `json:"commonBizType,omitempty" xml:"commonBizType,omitempty"`
	// example:
	//
	// true
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// example:
	//
	// true
	Duration *bool `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// xxxx
	DurationLabel *string `json:"durationLabel,omitempty" xml:"durationLabel,omitempty"`
	// example:
	//
	// true
	ESign *bool `json:"eSign,omitempty" xml:"eSign,omitempty"`
	// example:
	//
	// true
	Extract *bool `json:"extract,omitempty" xml:"extract,omitempty"`
	// example:
	//
	// xxxx
	FieldsInfo *string `json:"fieldsInfo,omitempty" xml:"fieldsInfo,omitempty"`
	// example:
	//
	// yyyy-MM-dd
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// example:
	//
	// xxxx
	Formula *string `json:"formula,omitempty" xml:"formula,omitempty"`
	// example:
	//
	// true
	Hidden *bool `json:"hidden,omitempty" xml:"hidden,omitempty"`
	// example:
	//
	// true
	HiddenInApprovalDetail *bool `json:"hiddenInApprovalDetail,omitempty" xml:"hiddenInApprovalDetail,omitempty"`
	// example:
	//
	// true
	HideLabel *bool `json:"hideLabel,omitempty" xml:"hideLabel,omitempty"`
	// example:
	//
	// "[{\"name\":\"\open"}]"
	HolidayOptions []map[string]*string `json:"holidayOptions,omitempty" xml:"holidayOptions,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// TextField-K2AD4O5B
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 单行输入框
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// example:
	//
	// true
	LabelEditableFreeze *bool `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	// example:
	//
	// xxxx
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// example:
	//
	// xxxx
	MainTitle *string `json:"mainTitle,omitempty" xml:"mainTitle,omitempty"`
	// example:
	//
	// 1
	NotPrint *string `json:"notPrint,omitempty" xml:"notPrint,omitempty"`
	// example:
	//
	// 1
	NotUpper   *string                                                                    `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	ObjOptions []*PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsObjOptions `json:"objOptions,omitempty" xml:"objOptions,omitempty" type:"Repeated"`
	Options    []*string                                                                  `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// example:
	//
	// true
	PayEnable *bool `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	// example:
	//
	// 请输入文字
	Placeholder *string                                                            `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	Push        *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsPush `json:"push,omitempty" xml:"push,omitempty" type:"Struct"`
	// example:
	//
	// true
	PushToAttendance *bool `json:"pushToAttendance,omitempty" xml:"pushToAttendance,omitempty"`
	// example:
	//
	// 1
	PushToCalendar *int32 `json:"pushToCalendar,omitempty" xml:"pushToCalendar,omitempty"`
	// example:
	//
	// true
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// example:
	//
	// true
	RequiredEditableFreeze *bool `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	// example:
	//
	// true
	ShowAttendOptions *bool `json:"showAttendOptions,omitempty" xml:"showAttendOptions,omitempty"`
	// example:
	//
	// true
	StaffStatusEnabled *bool                                                                     `json:"staffStatusEnabled,omitempty" xml:"staffStatusEnabled,omitempty"`
	StatField          []*PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsStatField `json:"statField,omitempty" xml:"statField,omitempty" type:"Repeated"`
	// example:
	//
	// list
	TableViewMode *string `json:"tableViewMode,omitempty" xml:"tableViewMode,omitempty"`
	// example:
	//
	// 天
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// example:
	//
	// true
	UseCalendar *bool `json:"useCalendar,omitempty" xml:"useCalendar,omitempty"`
	// example:
	//
	// true
	VerticalPrint *bool `json:"verticalPrint,omitempty" xml:"verticalPrint,omitempty"`
}

func (s PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) GoString() string {
	return s.String()
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetActionName(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.ActionName = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetAlign(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.Align = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetAppId(v int64) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.AppId = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetAsyncCondition(v bool) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.AsyncCondition = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetAttendTypeLabel(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.AttendTypeLabel = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetBehaviorLinkage(v []*PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsBehaviorLinkage) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.BehaviorLinkage = v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetBizAlias(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.BizAlias = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetBizType(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.BizType = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetChildFieldVisible(v map[string]*bool) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.ChildFieldVisible = v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetChoice(v int32) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.Choice = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetCommonBizType(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.CommonBizType = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetDisabled(v bool) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.Disabled = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetDuration(v bool) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.Duration = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetDurationLabel(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.DurationLabel = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetESign(v bool) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.ESign = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetExtract(v bool) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.Extract = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetFieldsInfo(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.FieldsInfo = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetFormat(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.Format = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetFormula(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.Formula = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetHidden(v bool) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.Hidden = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetHiddenInApprovalDetail(v bool) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.HiddenInApprovalDetail = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetHideLabel(v bool) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.HideLabel = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetHolidayOptions(v []map[string]*string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.HolidayOptions = v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetId(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.Id = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetLabel(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.Label = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetLabelEditableFreeze(v bool) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.LabelEditableFreeze = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetLink(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.Link = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetMainTitle(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.MainTitle = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetNotPrint(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.NotPrint = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetNotUpper(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.NotUpper = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetObjOptions(v []*PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsObjOptions) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.ObjOptions = v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetOptions(v []*string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.Options = v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetPayEnable(v bool) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.PayEnable = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetPlaceholder(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.Placeholder = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetPush(v *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsPush) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.Push = v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetPushToAttendance(v bool) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.PushToAttendance = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetPushToCalendar(v int32) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.PushToCalendar = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetRequired(v bool) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.Required = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetRequiredEditableFreeze(v bool) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.RequiredEditableFreeze = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetShowAttendOptions(v bool) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.ShowAttendOptions = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetStaffStatusEnabled(v bool) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.StaffStatusEnabled = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetStatField(v []*PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsStatField) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.StatField = v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetTableViewMode(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.TableViewMode = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetUnit(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.Unit = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetUseCalendar(v bool) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.UseCalendar = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps) SetVerticalPrint(v bool) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsProps {
	s.VerticalPrint = &v
	return s
}

type PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsBehaviorLinkage struct {
	Targets []*PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets `json:"targets,omitempty" xml:"targets,omitempty" type:"Repeated"`
	// example:
	//
	// xxxx
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsBehaviorLinkage) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsBehaviorLinkage) GoString() string {
	return s.String()
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsBehaviorLinkage) SetTargets(v []*PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsBehaviorLinkage {
	s.Targets = v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsBehaviorLinkage) SetValue(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsBehaviorLinkage {
	s.Value = &v
	return s
}

type PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets struct {
	// example:
	//
	// xxxx
	Behavior *string `json:"behavior,omitempty" xml:"behavior,omitempty"`
	// example:
	//
	// TextField-K2AD4O5B
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
}

func (s PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets) GoString() string {
	return s.String()
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets) SetBehavior(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets {
	s.Behavior = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets) SetFieldId(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets {
	s.FieldId = &v
	return s
}

type PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsObjOptions struct {
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsObjOptions) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsObjOptions) GoString() string {
	return s.String()
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsObjOptions) SetValue(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsObjOptions {
	s.Value = &v
	return s
}

type PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsPush struct {
	// example:
	//
	// 1
	AttendanceRule *int32 `json:"attendanceRule,omitempty" xml:"attendanceRule,omitempty"`
	// example:
	//
	// 1
	PushSwitch *int32 `json:"pushSwitch,omitempty" xml:"pushSwitch,omitempty"`
	// example:
	//
	// xxxx
	PushTag *string `json:"pushTag,omitempty" xml:"pushTag,omitempty"`
}

func (s PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsPush) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsPush) GoString() string {
	return s.String()
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsPush) SetAttendanceRule(v int32) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsPush {
	s.AttendanceRule = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsPush) SetPushSwitch(v int32) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsPush {
	s.PushSwitch = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsPush) SetPushTag(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsPush {
	s.PushTag = &v
	return s
}

type PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsStatField struct {
	// example:
	//
	// TextField-K2AD4O5B
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 单行输入框
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// example:
	//
	// xxxx
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// example:
	//
	// true
	Upper *bool `json:"upper,omitempty" xml:"upper,omitempty"`
}

func (s PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsStatField) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsStatField) GoString() string {
	return s.String()
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsStatField) SetId(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsStatField {
	s.Id = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsStatField) SetLabel(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsStatField {
	s.Label = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsStatField) SetUnit(v string) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsStatField {
	s.Unit = &v
	return s
}

func (s *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsStatField) SetUpper(v bool) *PremiumGetFormSchemaResponseBodyResultSchemaContentItemsPropsStatField {
	s.Upper = &v
	return s
}

type PremiumGetFormSchemaResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumGetFormSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumGetFormSchemaResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetFormSchemaResponse) GoString() string {
	return s.String()
}

func (s *PremiumGetFormSchemaResponse) SetHeaders(v map[string]*string) *PremiumGetFormSchemaResponse {
	s.Headers = v
	return s
}

func (s *PremiumGetFormSchemaResponse) SetStatusCode(v int32) *PremiumGetFormSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumGetFormSchemaResponse) SetBody(v *PremiumGetFormSchemaResponseBody) *PremiumGetFormSchemaResponse {
	s.Body = v
	return s
}

type PremiumGetInstFieldSettingHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumGetInstFieldSettingHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetInstFieldSettingHeaders) GoString() string {
	return s.String()
}

func (s *PremiumGetInstFieldSettingHeaders) SetCommonHeaders(v map[string]*string) *PremiumGetInstFieldSettingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumGetInstFieldSettingHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumGetInstFieldSettingHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumGetInstFieldSettingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// proc-FF6Y2xxxx
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// userId123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PremiumGetInstFieldSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetInstFieldSettingRequest) GoString() string {
	return s.String()
}

func (s *PremiumGetInstFieldSettingRequest) SetProcessInstanceId(v string) *PremiumGetInstFieldSettingRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *PremiumGetInstFieldSettingRequest) SetUserId(v string) *PremiumGetInstFieldSettingRequest {
	s.UserId = &v
	return s
}

type PremiumGetInstFieldSettingResponseBody struct {
	Result []*PremiumGetInstFieldSettingResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PremiumGetInstFieldSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetInstFieldSettingResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumGetInstFieldSettingResponseBody) SetResult(v []*PremiumGetInstFieldSettingResponseBodyResult) *PremiumGetInstFieldSettingResponseBody {
	s.Result = v
	return s
}

func (s *PremiumGetInstFieldSettingResponseBody) SetSuccess(v bool) *PremiumGetInstFieldSettingResponseBody {
	s.Success = &v
	return s
}

type PremiumGetInstFieldSettingResponseBodyResult struct {
	// example:
	//
	// TextField
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// example:
	//
	// READONLY
	FieldBehavior *string `json:"fieldBehavior,omitempty" xml:"fieldBehavior,omitempty"`
	// example:
	//
	// TextField-abcd
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
}

func (s PremiumGetInstFieldSettingResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetInstFieldSettingResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PremiumGetInstFieldSettingResponseBodyResult) SetComponentType(v string) *PremiumGetInstFieldSettingResponseBodyResult {
	s.ComponentType = &v
	return s
}

func (s *PremiumGetInstFieldSettingResponseBodyResult) SetFieldBehavior(v string) *PremiumGetInstFieldSettingResponseBodyResult {
	s.FieldBehavior = &v
	return s
}

func (s *PremiumGetInstFieldSettingResponseBodyResult) SetFieldId(v string) *PremiumGetInstFieldSettingResponseBodyResult {
	s.FieldId = &v
	return s
}

type PremiumGetInstFieldSettingResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumGetInstFieldSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumGetInstFieldSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetInstFieldSettingResponse) GoString() string {
	return s.String()
}

func (s *PremiumGetInstFieldSettingResponse) SetHeaders(v map[string]*string) *PremiumGetInstFieldSettingResponse {
	s.Headers = v
	return s
}

func (s *PremiumGetInstFieldSettingResponse) SetStatusCode(v int32) *PremiumGetInstFieldSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumGetInstFieldSettingResponse) SetBody(v *PremiumGetInstFieldSettingResponseBody) *PremiumGetInstFieldSettingResponse {
	s.Body = v
	return s
}

type PremiumGetNoticedInstancesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumGetNoticedInstancesHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetNoticedInstancesHeaders) GoString() string {
	return s.String()
}

func (s *PremiumGetNoticedInstancesHeaders) SetCommonHeaders(v map[string]*string) *PremiumGetNoticedInstancesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumGetNoticedInstancesHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumGetNoticedInstancesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumGetNoticedInstancesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// manager001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PremiumGetNoticedInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetNoticedInstancesRequest) GoString() string {
	return s.String()
}

func (s *PremiumGetNoticedInstancesRequest) SetPageNumber(v int32) *PremiumGetNoticedInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *PremiumGetNoticedInstancesRequest) SetPageSize(v int32) *PremiumGetNoticedInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *PremiumGetNoticedInstancesRequest) SetUserId(v string) *PremiumGetNoticedInstancesRequest {
	s.UserId = &v
	return s
}

type PremiumGetNoticedInstancesResponseBody struct {
	Result  *PremiumGetNoticedInstancesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PremiumGetNoticedInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetNoticedInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumGetNoticedInstancesResponseBody) SetResult(v *PremiumGetNoticedInstancesResponseBodyResult) *PremiumGetNoticedInstancesResponseBody {
	s.Result = v
	return s
}

func (s *PremiumGetNoticedInstancesResponseBody) SetSuccess(v bool) *PremiumGetNoticedInstancesResponseBody {
	s.Success = &v
	return s
}

type PremiumGetNoticedInstancesResponseBodyResult struct {
	HasMore *bool                                               `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List    []*PremiumGetNoticedInstancesResponseBodyResultList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s PremiumGetNoticedInstancesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetNoticedInstancesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PremiumGetNoticedInstancesResponseBodyResult) SetHasMore(v bool) *PremiumGetNoticedInstancesResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *PremiumGetNoticedInstancesResponseBodyResult) SetList(v []*PremiumGetNoticedInstancesResponseBodyResultList) *PremiumGetNoticedInstancesResponseBodyResult {
	s.List = v
	return s
}

type PremiumGetNoticedInstancesResponseBodyResultList struct {
	FormMassage     *string `json:"formMassage,omitempty" xml:"formMassage,omitempty"`
	OriginatorId    *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	OriginatorName  *string `json:"originatorName,omitempty" xml:"originatorName,omitempty"`
	OriginatorPhoto *string `json:"originatorPhoto,omitempty" xml:"originatorPhoto,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	ProcessCreateTime *string `json:"processCreateTime,omitempty" xml:"processCreateTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	ProcessEndTime    *string `json:"processEndTime,omitempty" xml:"processEndTime,omitempty"`
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	ProcessType       *int32  `json:"processType,omitempty" xml:"processType,omitempty"`
	// example:
	//
	// agree
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	Title  *string `json:"title,omitempty" xml:"title,omitempty"`
	Url    *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s PremiumGetNoticedInstancesResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetNoticedInstancesResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *PremiumGetNoticedInstancesResponseBodyResultList) SetFormMassage(v string) *PremiumGetNoticedInstancesResponseBodyResultList {
	s.FormMassage = &v
	return s
}

func (s *PremiumGetNoticedInstancesResponseBodyResultList) SetOriginatorId(v string) *PremiumGetNoticedInstancesResponseBodyResultList {
	s.OriginatorId = &v
	return s
}

func (s *PremiumGetNoticedInstancesResponseBodyResultList) SetOriginatorName(v string) *PremiumGetNoticedInstancesResponseBodyResultList {
	s.OriginatorName = &v
	return s
}

func (s *PremiumGetNoticedInstancesResponseBodyResultList) SetOriginatorPhoto(v string) *PremiumGetNoticedInstancesResponseBodyResultList {
	s.OriginatorPhoto = &v
	return s
}

func (s *PremiumGetNoticedInstancesResponseBodyResultList) SetProcessCreateTime(v string) *PremiumGetNoticedInstancesResponseBodyResultList {
	s.ProcessCreateTime = &v
	return s
}

func (s *PremiumGetNoticedInstancesResponseBodyResultList) SetProcessEndTime(v string) *PremiumGetNoticedInstancesResponseBodyResultList {
	s.ProcessEndTime = &v
	return s
}

func (s *PremiumGetNoticedInstancesResponseBodyResultList) SetProcessInstanceId(v string) *PremiumGetNoticedInstancesResponseBodyResultList {
	s.ProcessInstanceId = &v
	return s
}

func (s *PremiumGetNoticedInstancesResponseBodyResultList) SetProcessType(v int32) *PremiumGetNoticedInstancesResponseBodyResultList {
	s.ProcessType = &v
	return s
}

func (s *PremiumGetNoticedInstancesResponseBodyResultList) SetResult(v string) *PremiumGetNoticedInstancesResponseBodyResultList {
	s.Result = &v
	return s
}

func (s *PremiumGetNoticedInstancesResponseBodyResultList) SetStatus(v string) *PremiumGetNoticedInstancesResponseBodyResultList {
	s.Status = &v
	return s
}

func (s *PremiumGetNoticedInstancesResponseBodyResultList) SetTitle(v string) *PremiumGetNoticedInstancesResponseBodyResultList {
	s.Title = &v
	return s
}

func (s *PremiumGetNoticedInstancesResponseBodyResultList) SetUrl(v string) *PremiumGetNoticedInstancesResponseBodyResultList {
	s.Url = &v
	return s
}

type PremiumGetNoticedInstancesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumGetNoticedInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumGetNoticedInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetNoticedInstancesResponse) GoString() string {
	return s.String()
}

func (s *PremiumGetNoticedInstancesResponse) SetHeaders(v map[string]*string) *PremiumGetNoticedInstancesResponse {
	s.Headers = v
	return s
}

func (s *PremiumGetNoticedInstancesResponse) SetStatusCode(v int32) *PremiumGetNoticedInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumGetNoticedInstancesResponse) SetBody(v *PremiumGetNoticedInstancesResponseBody) *PremiumGetNoticedInstancesResponse {
	s.Body = v
	return s
}

type PremiumGetProcessInstancesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumGetProcessInstancesHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetProcessInstancesHeaders) GoString() string {
	return s.String()
}

func (s *PremiumGetProcessInstancesHeaders) SetCommonHeaders(v map[string]*string) *PremiumGetProcessInstancesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumGetProcessInstancesHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumGetProcessInstancesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumGetProcessInstancesRequest struct {
	// example:
	//
	// SWAPP-4C2F4B-example
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// example:
	//
	// 1633795200000
	EndTimeInMills *int64 `json:"endTimeInMills,omitempty" xml:"endTimeInMills,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-C53-example
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1631289600000
	StartTimeInMills *int64 `json:"startTimeInMills,omitempty" xml:"startTimeInMills,omitempty"`
}

func (s PremiumGetProcessInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetProcessInstancesRequest) GoString() string {
	return s.String()
}

func (s *PremiumGetProcessInstancesRequest) SetAppUuid(v string) *PremiumGetProcessInstancesRequest {
	s.AppUuid = &v
	return s
}

func (s *PremiumGetProcessInstancesRequest) SetEndTimeInMills(v int64) *PremiumGetProcessInstancesRequest {
	s.EndTimeInMills = &v
	return s
}

func (s *PremiumGetProcessInstancesRequest) SetMaxResults(v int64) *PremiumGetProcessInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *PremiumGetProcessInstancesRequest) SetNextToken(v string) *PremiumGetProcessInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *PremiumGetProcessInstancesRequest) SetProcessCode(v string) *PremiumGetProcessInstancesRequest {
	s.ProcessCode = &v
	return s
}

func (s *PremiumGetProcessInstancesRequest) SetStartTimeInMills(v int64) *PremiumGetProcessInstancesRequest {
	s.StartTimeInMills = &v
	return s
}

type PremiumGetProcessInstancesResponseBody struct {
	Result *PremiumGetProcessInstancesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s PremiumGetProcessInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetProcessInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumGetProcessInstancesResponseBody) SetResult(v *PremiumGetProcessInstancesResponseBodyResult) *PremiumGetProcessInstancesResponseBody {
	s.Result = v
	return s
}

type PremiumGetProcessInstancesResponseBodyResult struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	HasMore *bool                                               `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List    []*PremiumGetProcessInstancesResponseBodyResultList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s PremiumGetProcessInstancesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetProcessInstancesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PremiumGetProcessInstancesResponseBodyResult) SetHasMore(v bool) *PremiumGetProcessInstancesResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResult) SetList(v []*PremiumGetProcessInstancesResponseBodyResultList) *PremiumGetProcessInstancesResponseBodyResult {
	s.List = v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResult) SetMaxResults(v int64) *PremiumGetProcessInstancesResponseBodyResult {
	s.MaxResults = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResult) SetNextToken(v string) *PremiumGetProcessInstancesResponseBodyResult {
	s.NextToken = &v
	return s
}

type PremiumGetProcessInstancesResponseBodyResultList struct {
	// This parameter is required.
	//
	// example:
	//
	// cdef-dae2fd2-example
	AttachedProcessInstanceIds *string `json:"attachedProcessInstanceIds,omitempty" xml:"attachedProcessInstanceIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 202110111558000355024
	BusinessId *string `json:"businessId,omitempty" xml:"businessId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1635165470201
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1633795200000
	FinishTime *int64 `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// This parameter is required.
	FormComponentValues []*PremiumGetProcessInstancesResponseBodyResultListFormComponentValues `json:"formComponentValues,omitempty" xml:"formComponentValues,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// dcdse-dae2fd2-example
	MainProcessInstanceId *string                                                             `json:"mainProcessInstanceId,omitempty" xml:"mainProcessInstanceId,omitempty"`
	OperationRecords      []*PremiumGetProcessInstancesResponseBodyResultListOperationRecords `json:"operationRecords,omitempty" xml:"operationRecords,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 默认-1，企业根部门
	OriginatorDeptId *string `json:"originatorDeptId,omitempty" xml:"originatorDeptId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// staff1234
	OriginatorUserid *string `json:"originatorUserid,omitempty" xml:"originatorUserid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abcdse-dse-example
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AGREE同意，REFUSE拒绝
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RUNNING审批中、TERMINATED撤销、COMPLETED审批完成、CANCELED取消
	Status *string                                                  `json:"status,omitempty" xml:"status,omitempty"`
	Tasks  []*PremiumGetProcessInstancesResponseBodyResultListTasks `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 员工A提交的小肖审批单
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s PremiumGetProcessInstancesResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetProcessInstancesResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *PremiumGetProcessInstancesResponseBodyResultList) SetAttachedProcessInstanceIds(v string) *PremiumGetProcessInstancesResponseBodyResultList {
	s.AttachedProcessInstanceIds = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultList) SetBusinessId(v string) *PremiumGetProcessInstancesResponseBodyResultList {
	s.BusinessId = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultList) SetCreateTime(v int64) *PremiumGetProcessInstancesResponseBodyResultList {
	s.CreateTime = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultList) SetFinishTime(v int64) *PremiumGetProcessInstancesResponseBodyResultList {
	s.FinishTime = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultList) SetFormComponentValues(v []*PremiumGetProcessInstancesResponseBodyResultListFormComponentValues) *PremiumGetProcessInstancesResponseBodyResultList {
	s.FormComponentValues = v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultList) SetMainProcessInstanceId(v string) *PremiumGetProcessInstancesResponseBodyResultList {
	s.MainProcessInstanceId = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultList) SetOperationRecords(v []*PremiumGetProcessInstancesResponseBodyResultListOperationRecords) *PremiumGetProcessInstancesResponseBodyResultList {
	s.OperationRecords = v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultList) SetOriginatorDeptId(v string) *PremiumGetProcessInstancesResponseBodyResultList {
	s.OriginatorDeptId = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultList) SetOriginatorUserid(v string) *PremiumGetProcessInstancesResponseBodyResultList {
	s.OriginatorUserid = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultList) SetProcessInstanceId(v string) *PremiumGetProcessInstancesResponseBodyResultList {
	s.ProcessInstanceId = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultList) SetResult(v string) *PremiumGetProcessInstancesResponseBodyResultList {
	s.Result = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultList) SetStatus(v string) *PremiumGetProcessInstancesResponseBodyResultList {
	s.Status = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultList) SetTasks(v []*PremiumGetProcessInstancesResponseBodyResultListTasks) *PremiumGetProcessInstancesResponseBodyResultList {
	s.Tasks = v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultList) SetTitle(v string) *PremiumGetProcessInstancesResponseBodyResultList {
	s.Title = &v
	return s
}

type PremiumGetProcessInstancesResponseBodyResultListFormComponentValues struct {
	// example:
	//
	// {"staffId":"abcd"}
	ExtValue *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TextField-a32bcdef
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 张三
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s PremiumGetProcessInstancesResponseBodyResultListFormComponentValues) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetProcessInstancesResponseBodyResultListFormComponentValues) GoString() string {
	return s.String()
}

func (s *PremiumGetProcessInstancesResponseBodyResultListFormComponentValues) SetExtValue(v string) *PremiumGetProcessInstancesResponseBodyResultListFormComponentValues {
	s.ExtValue = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultListFormComponentValues) SetId(v string) *PremiumGetProcessInstancesResponseBodyResultListFormComponentValues {
	s.Id = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultListFormComponentValues) SetName(v string) *PremiumGetProcessInstancesResponseBodyResultListFormComponentValues {
	s.Name = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultListFormComponentValues) SetValue(v string) *PremiumGetProcessInstancesResponseBodyResultListFormComponentValues {
	s.Value = &v
	return s
}

type PremiumGetProcessInstancesResponseBodyResultListOperationRecords struct {
	Attachments []*PremiumGetProcessInstancesResponseBodyResultListOperationRecordsAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// example:
	//
	// EXECUTE_TASK_NORMAL（正常执行任务），EXECUTE_TASK_AGENT（代理人执行任务），APPEND_TASK_BEFORE（前加签任务），APPEND_TASK_AFTER（后加签任务），REDIRECT_TASK（转交任务），START_PROCESS_INSTANCE（发起流程实例），TERMINATE_PROCESS_INSTANCE（终止(撤销)流程实例），FINISH_PROCESS_INSTANCE（结束流程实例），ADD_REMARK（添加评论）
	OperationType *string `json:"operationType,omitempty" xml:"operationType,omitempty"`
	// example:
	//
	// 同意
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// AGREE（同意），REFUSE（拒绝），NONE（未知）
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// 1657522271000
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// example:
	//
	// manager1
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PremiumGetProcessInstancesResponseBodyResultListOperationRecords) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetProcessInstancesResponseBodyResultListOperationRecords) GoString() string {
	return s.String()
}

func (s *PremiumGetProcessInstancesResponseBodyResultListOperationRecords) SetAttachments(v []*PremiumGetProcessInstancesResponseBodyResultListOperationRecordsAttachments) *PremiumGetProcessInstancesResponseBodyResultListOperationRecords {
	s.Attachments = v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultListOperationRecords) SetOperationType(v string) *PremiumGetProcessInstancesResponseBodyResultListOperationRecords {
	s.OperationType = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultListOperationRecords) SetRemark(v string) *PremiumGetProcessInstancesResponseBodyResultListOperationRecords {
	s.Remark = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultListOperationRecords) SetResult(v string) *PremiumGetProcessInstancesResponseBodyResultListOperationRecords {
	s.Result = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultListOperationRecords) SetTimestamp(v int64) *PremiumGetProcessInstancesResponseBodyResultListOperationRecords {
	s.Timestamp = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultListOperationRecords) SetUserId(v string) *PremiumGetProcessInstancesResponseBodyResultListOperationRecords {
	s.UserId = &v
	return s
}

type PremiumGetProcessInstancesResponseBodyResultListOperationRecordsAttachments struct {
	// example:
	//
	// 1234567
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// example:
	//
	// 附件
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// 123
	FileSize *string `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
}

func (s PremiumGetProcessInstancesResponseBodyResultListOperationRecordsAttachments) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetProcessInstancesResponseBodyResultListOperationRecordsAttachments) GoString() string {
	return s.String()
}

func (s *PremiumGetProcessInstancesResponseBodyResultListOperationRecordsAttachments) SetFileId(v string) *PremiumGetProcessInstancesResponseBodyResultListOperationRecordsAttachments {
	s.FileId = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultListOperationRecordsAttachments) SetFileName(v string) *PremiumGetProcessInstancesResponseBodyResultListOperationRecordsAttachments {
	s.FileName = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultListOperationRecordsAttachments) SetFileSize(v string) *PremiumGetProcessInstancesResponseBodyResultListOperationRecordsAttachments {
	s.FileSize = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultListOperationRecordsAttachments) SetFileType(v string) *PremiumGetProcessInstancesResponseBodyResultListOperationRecordsAttachments {
	s.FileType = &v
	return s
}

type PremiumGetProcessInstancesResponseBodyResultListTasks struct {
	// example:
	//
	// 1234_abcd
	ActivityId *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
	// example:
	//
	// 1657522271000
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// example:
	//
	// 1657522271000
	FinishTimestamp *int64 `json:"finishTimestamp,omitempty" xml:"finishTimestamp,omitempty"`
	// example:
	//
	// 分为AGREE（同意），REFUSE（拒绝），REDIRECTED（转交）
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// NEW（未启动），RUNNING（处理中），PAUSED（暂停），CANCELED（取消），COMPLETED（完成），TERMINATED（终止）
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// staff1234
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PremiumGetProcessInstancesResponseBodyResultListTasks) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetProcessInstancesResponseBodyResultListTasks) GoString() string {
	return s.String()
}

func (s *PremiumGetProcessInstancesResponseBodyResultListTasks) SetActivityId(v string) *PremiumGetProcessInstancesResponseBodyResultListTasks {
	s.ActivityId = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultListTasks) SetCreateTimestamp(v int64) *PremiumGetProcessInstancesResponseBodyResultListTasks {
	s.CreateTimestamp = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultListTasks) SetFinishTimestamp(v int64) *PremiumGetProcessInstancesResponseBodyResultListTasks {
	s.FinishTimestamp = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultListTasks) SetResult(v string) *PremiumGetProcessInstancesResponseBodyResultListTasks {
	s.Result = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultListTasks) SetStatus(v string) *PremiumGetProcessInstancesResponseBodyResultListTasks {
	s.Status = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultListTasks) SetTaskId(v int64) *PremiumGetProcessInstancesResponseBodyResultListTasks {
	s.TaskId = &v
	return s
}

func (s *PremiumGetProcessInstancesResponseBodyResultListTasks) SetUserId(v string) *PremiumGetProcessInstancesResponseBodyResultListTasks {
	s.UserId = &v
	return s
}

type PremiumGetProcessInstancesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumGetProcessInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumGetProcessInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetProcessInstancesResponse) GoString() string {
	return s.String()
}

func (s *PremiumGetProcessInstancesResponse) SetHeaders(v map[string]*string) *PremiumGetProcessInstancesResponse {
	s.Headers = v
	return s
}

func (s *PremiumGetProcessInstancesResponse) SetStatusCode(v int32) *PremiumGetProcessInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumGetProcessInstancesResponse) SetBody(v *PremiumGetProcessInstancesResponseBody) *PremiumGetProcessInstancesResponse {
	s.Body = v
	return s
}

type PremiumGetSpaceWithDownloadAuthHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumGetSpaceWithDownloadAuthHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetSpaceWithDownloadAuthHeaders) GoString() string {
	return s.String()
}

func (s *PremiumGetSpaceWithDownloadAuthHeaders) SetCommonHeaders(v map[string]*string) *PremiumGetSpaceWithDownloadAuthHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumGetSpaceWithDownloadAuthHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumGetSpaceWithDownloadAuthHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumGetSpaceWithDownloadAuthRequest struct {
	// example:
	//
	// 8345000
	AgentId *int64 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111
	FileId     *string   `json:"fileId,omitempty" xml:"fileId,omitempty"`
	FileIdList []*string `json:"fileIdList,omitempty" xml:"fileIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// a17444d1-075b-4a4d-xxxx
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// if can be null:
	// true
	WithCommentAttatchment *bool `json:"withCommentAttatchment,omitempty" xml:"withCommentAttatchment,omitempty"`
}

func (s PremiumGetSpaceWithDownloadAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetSpaceWithDownloadAuthRequest) GoString() string {
	return s.String()
}

func (s *PremiumGetSpaceWithDownloadAuthRequest) SetAgentId(v int64) *PremiumGetSpaceWithDownloadAuthRequest {
	s.AgentId = &v
	return s
}

func (s *PremiumGetSpaceWithDownloadAuthRequest) SetFileId(v string) *PremiumGetSpaceWithDownloadAuthRequest {
	s.FileId = &v
	return s
}

func (s *PremiumGetSpaceWithDownloadAuthRequest) SetFileIdList(v []*string) *PremiumGetSpaceWithDownloadAuthRequest {
	s.FileIdList = v
	return s
}

func (s *PremiumGetSpaceWithDownloadAuthRequest) SetProcessInstanceId(v string) *PremiumGetSpaceWithDownloadAuthRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *PremiumGetSpaceWithDownloadAuthRequest) SetUserId(v string) *PremiumGetSpaceWithDownloadAuthRequest {
	s.UserId = &v
	return s
}

func (s *PremiumGetSpaceWithDownloadAuthRequest) SetWithCommentAttatchment(v bool) *PremiumGetSpaceWithDownloadAuthRequest {
	s.WithCommentAttatchment = &v
	return s
}

type PremiumGetSpaceWithDownloadAuthResponseBody struct {
	Result *PremiumGetSpaceWithDownloadAuthResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PremiumGetSpaceWithDownloadAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetSpaceWithDownloadAuthResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumGetSpaceWithDownloadAuthResponseBody) SetResult(v *PremiumGetSpaceWithDownloadAuthResponseBodyResult) *PremiumGetSpaceWithDownloadAuthResponseBody {
	s.Result = v
	return s
}

func (s *PremiumGetSpaceWithDownloadAuthResponseBody) SetSuccess(v bool) *PremiumGetSpaceWithDownloadAuthResponseBody {
	s.Success = &v
	return s
}

type PremiumGetSpaceWithDownloadAuthResponseBodyResult struct {
	// example:
	//
	// 3996960664
	SpaceId *int64 `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s PremiumGetSpaceWithDownloadAuthResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetSpaceWithDownloadAuthResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PremiumGetSpaceWithDownloadAuthResponseBodyResult) SetSpaceId(v int64) *PremiumGetSpaceWithDownloadAuthResponseBodyResult {
	s.SpaceId = &v
	return s
}

type PremiumGetSpaceWithDownloadAuthResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumGetSpaceWithDownloadAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumGetSpaceWithDownloadAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetSpaceWithDownloadAuthResponse) GoString() string {
	return s.String()
}

func (s *PremiumGetSpaceWithDownloadAuthResponse) SetHeaders(v map[string]*string) *PremiumGetSpaceWithDownloadAuthResponse {
	s.Headers = v
	return s
}

func (s *PremiumGetSpaceWithDownloadAuthResponse) SetStatusCode(v int32) *PremiumGetSpaceWithDownloadAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumGetSpaceWithDownloadAuthResponse) SetBody(v *PremiumGetSpaceWithDownloadAuthResponseBody) *PremiumGetSpaceWithDownloadAuthResponse {
	s.Body = v
	return s
}

type PremiumGetSubmittedInstancesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumGetSubmittedInstancesHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetSubmittedInstancesHeaders) GoString() string {
	return s.String()
}

func (s *PremiumGetSubmittedInstancesHeaders) SetCommonHeaders(v map[string]*string) *PremiumGetSubmittedInstancesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumGetSubmittedInstancesHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumGetSubmittedInstancesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumGetSubmittedInstancesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// manager001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PremiumGetSubmittedInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetSubmittedInstancesRequest) GoString() string {
	return s.String()
}

func (s *PremiumGetSubmittedInstancesRequest) SetPageNumber(v int32) *PremiumGetSubmittedInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *PremiumGetSubmittedInstancesRequest) SetPageSize(v int32) *PremiumGetSubmittedInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *PremiumGetSubmittedInstancesRequest) SetUserId(v string) *PremiumGetSubmittedInstancesRequest {
	s.UserId = &v
	return s
}

type PremiumGetSubmittedInstancesResponseBody struct {
	Result  *PremiumGetSubmittedInstancesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PremiumGetSubmittedInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetSubmittedInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumGetSubmittedInstancesResponseBody) SetResult(v *PremiumGetSubmittedInstancesResponseBodyResult) *PremiumGetSubmittedInstancesResponseBody {
	s.Result = v
	return s
}

func (s *PremiumGetSubmittedInstancesResponseBody) SetSuccess(v bool) *PremiumGetSubmittedInstancesResponseBody {
	s.Success = &v
	return s
}

type PremiumGetSubmittedInstancesResponseBodyResult struct {
	HasMore *bool                                                 `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List    []*PremiumGetSubmittedInstancesResponseBodyResultList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s PremiumGetSubmittedInstancesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetSubmittedInstancesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PremiumGetSubmittedInstancesResponseBodyResult) SetHasMore(v bool) *PremiumGetSubmittedInstancesResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *PremiumGetSubmittedInstancesResponseBodyResult) SetList(v []*PremiumGetSubmittedInstancesResponseBodyResultList) *PremiumGetSubmittedInstancesResponseBodyResult {
	s.List = v
	return s
}

type PremiumGetSubmittedInstancesResponseBodyResultList struct {
	AppType         *int32  `json:"appType,omitempty" xml:"appType,omitempty"`
	FormMassage     *string `json:"formMassage,omitempty" xml:"formMassage,omitempty"`
	OriginatorId    *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	OriginatorName  *string `json:"originatorName,omitempty" xml:"originatorName,omitempty"`
	OriginatorPhoto *string `json:"originatorPhoto,omitempty" xml:"originatorPhoto,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	ProcessCreateTime *string `json:"processCreateTime,omitempty" xml:"processCreateTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	ProcessEndTime    *string `json:"processEndTime,omitempty" xml:"processEndTime,omitempty"`
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	ProcessType       *int32  `json:"processType,omitempty" xml:"processType,omitempty"`
	// example:
	//
	// agree
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	Title  *string `json:"title,omitempty" xml:"title,omitempty"`
	Url    *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s PremiumGetSubmittedInstancesResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetSubmittedInstancesResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *PremiumGetSubmittedInstancesResponseBodyResultList) SetAppType(v int32) *PremiumGetSubmittedInstancesResponseBodyResultList {
	s.AppType = &v
	return s
}

func (s *PremiumGetSubmittedInstancesResponseBodyResultList) SetFormMassage(v string) *PremiumGetSubmittedInstancesResponseBodyResultList {
	s.FormMassage = &v
	return s
}

func (s *PremiumGetSubmittedInstancesResponseBodyResultList) SetOriginatorId(v string) *PremiumGetSubmittedInstancesResponseBodyResultList {
	s.OriginatorId = &v
	return s
}

func (s *PremiumGetSubmittedInstancesResponseBodyResultList) SetOriginatorName(v string) *PremiumGetSubmittedInstancesResponseBodyResultList {
	s.OriginatorName = &v
	return s
}

func (s *PremiumGetSubmittedInstancesResponseBodyResultList) SetOriginatorPhoto(v string) *PremiumGetSubmittedInstancesResponseBodyResultList {
	s.OriginatorPhoto = &v
	return s
}

func (s *PremiumGetSubmittedInstancesResponseBodyResultList) SetProcessCreateTime(v string) *PremiumGetSubmittedInstancesResponseBodyResultList {
	s.ProcessCreateTime = &v
	return s
}

func (s *PremiumGetSubmittedInstancesResponseBodyResultList) SetProcessEndTime(v string) *PremiumGetSubmittedInstancesResponseBodyResultList {
	s.ProcessEndTime = &v
	return s
}

func (s *PremiumGetSubmittedInstancesResponseBodyResultList) SetProcessInstanceId(v string) *PremiumGetSubmittedInstancesResponseBodyResultList {
	s.ProcessInstanceId = &v
	return s
}

func (s *PremiumGetSubmittedInstancesResponseBodyResultList) SetProcessType(v int32) *PremiumGetSubmittedInstancesResponseBodyResultList {
	s.ProcessType = &v
	return s
}

func (s *PremiumGetSubmittedInstancesResponseBodyResultList) SetResult(v string) *PremiumGetSubmittedInstancesResponseBodyResultList {
	s.Result = &v
	return s
}

func (s *PremiumGetSubmittedInstancesResponseBodyResultList) SetStatus(v string) *PremiumGetSubmittedInstancesResponseBodyResultList {
	s.Status = &v
	return s
}

func (s *PremiumGetSubmittedInstancesResponseBodyResultList) SetTitle(v string) *PremiumGetSubmittedInstancesResponseBodyResultList {
	s.Title = &v
	return s
}

func (s *PremiumGetSubmittedInstancesResponseBodyResultList) SetUrl(v string) *PremiumGetSubmittedInstancesResponseBodyResultList {
	s.Url = &v
	return s
}

type PremiumGetSubmittedInstancesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumGetSubmittedInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumGetSubmittedInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetSubmittedInstancesResponse) GoString() string {
	return s.String()
}

func (s *PremiumGetSubmittedInstancesResponse) SetHeaders(v map[string]*string) *PremiumGetSubmittedInstancesResponse {
	s.Headers = v
	return s
}

func (s *PremiumGetSubmittedInstancesResponse) SetStatusCode(v int32) *PremiumGetSubmittedInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumGetSubmittedInstancesResponse) SetBody(v *PremiumGetSubmittedInstancesResponseBody) *PremiumGetSubmittedInstancesResponse {
	s.Body = v
	return s
}

type PremiumGetTodoTasksHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumGetTodoTasksHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetTodoTasksHeaders) GoString() string {
	return s.String()
}

func (s *PremiumGetTodoTasksHeaders) SetCommonHeaders(v map[string]*string) *PremiumGetTodoTasksHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumGetTodoTasksHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumGetTodoTasksHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumGetTodoTasksRequest struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	CreateBefore *string `json:"createBefore,omitempty" xml:"createBefore,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// manager001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PremiumGetTodoTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetTodoTasksRequest) GoString() string {
	return s.String()
}

func (s *PremiumGetTodoTasksRequest) SetCreateBefore(v string) *PremiumGetTodoTasksRequest {
	s.CreateBefore = &v
	return s
}

func (s *PremiumGetTodoTasksRequest) SetPageNumber(v int32) *PremiumGetTodoTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *PremiumGetTodoTasksRequest) SetPageSize(v int32) *PremiumGetTodoTasksRequest {
	s.PageSize = &v
	return s
}

func (s *PremiumGetTodoTasksRequest) SetUserId(v string) *PremiumGetTodoTasksRequest {
	s.UserId = &v
	return s
}

type PremiumGetTodoTasksResponseBody struct {
	Result *PremiumGetTodoTasksResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s PremiumGetTodoTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetTodoTasksResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumGetTodoTasksResponseBody) SetResult(v *PremiumGetTodoTasksResponseBodyResult) *PremiumGetTodoTasksResponseBody {
	s.Result = v
	return s
}

type PremiumGetTodoTasksResponseBodyResult struct {
	HasMore *bool                                        `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List    []*PremiumGetTodoTasksResponseBodyResultList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	Success *bool                                        `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PremiumGetTodoTasksResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetTodoTasksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PremiumGetTodoTasksResponseBodyResult) SetHasMore(v bool) *PremiumGetTodoTasksResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *PremiumGetTodoTasksResponseBodyResult) SetList(v []*PremiumGetTodoTasksResponseBodyResultList) *PremiumGetTodoTasksResponseBodyResult {
	s.List = v
	return s
}

func (s *PremiumGetTodoTasksResponseBodyResult) SetSuccess(v bool) *PremiumGetTodoTasksResponseBodyResult {
	s.Success = &v
	return s
}

type PremiumGetTodoTasksResponseBodyResultList struct {
	ActivityId      *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
	AppType         *int32  `json:"appType,omitempty" xml:"appType,omitempty"`
	FormMassage     *string `json:"formMassage,omitempty" xml:"formMassage,omitempty"`
	OriginatorId    *string `json:"originatorId,omitempty" xml:"originatorId,omitempty"`
	OriginatorName  *string `json:"originatorName,omitempty" xml:"originatorName,omitempty"`
	OriginatorPhoto *string `json:"originatorPhoto,omitempty" xml:"originatorPhoto,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	ProcessCreateTime *string `json:"processCreateTime,omitempty" xml:"processCreateTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	ProcessEndTime    *string `json:"processEndTime,omitempty" xml:"processEndTime,omitempty"`
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	ProcessType       *int32  `json:"processType,omitempty" xml:"processType,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	Title  *string `json:"title,omitempty" xml:"title,omitempty"`
	Url    *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s PremiumGetTodoTasksResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetTodoTasksResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *PremiumGetTodoTasksResponseBodyResultList) SetActivityId(v string) *PremiumGetTodoTasksResponseBodyResultList {
	s.ActivityId = &v
	return s
}

func (s *PremiumGetTodoTasksResponseBodyResultList) SetAppType(v int32) *PremiumGetTodoTasksResponseBodyResultList {
	s.AppType = &v
	return s
}

func (s *PremiumGetTodoTasksResponseBodyResultList) SetFormMassage(v string) *PremiumGetTodoTasksResponseBodyResultList {
	s.FormMassage = &v
	return s
}

func (s *PremiumGetTodoTasksResponseBodyResultList) SetOriginatorId(v string) *PremiumGetTodoTasksResponseBodyResultList {
	s.OriginatorId = &v
	return s
}

func (s *PremiumGetTodoTasksResponseBodyResultList) SetOriginatorName(v string) *PremiumGetTodoTasksResponseBodyResultList {
	s.OriginatorName = &v
	return s
}

func (s *PremiumGetTodoTasksResponseBodyResultList) SetOriginatorPhoto(v string) *PremiumGetTodoTasksResponseBodyResultList {
	s.OriginatorPhoto = &v
	return s
}

func (s *PremiumGetTodoTasksResponseBodyResultList) SetProcessCreateTime(v string) *PremiumGetTodoTasksResponseBodyResultList {
	s.ProcessCreateTime = &v
	return s
}

func (s *PremiumGetTodoTasksResponseBodyResultList) SetProcessEndTime(v string) *PremiumGetTodoTasksResponseBodyResultList {
	s.ProcessEndTime = &v
	return s
}

func (s *PremiumGetTodoTasksResponseBodyResultList) SetProcessInstanceId(v string) *PremiumGetTodoTasksResponseBodyResultList {
	s.ProcessInstanceId = &v
	return s
}

func (s *PremiumGetTodoTasksResponseBodyResultList) SetProcessType(v int32) *PremiumGetTodoTasksResponseBodyResultList {
	s.ProcessType = &v
	return s
}

func (s *PremiumGetTodoTasksResponseBodyResultList) SetStatus(v string) *PremiumGetTodoTasksResponseBodyResultList {
	s.Status = &v
	return s
}

func (s *PremiumGetTodoTasksResponseBodyResultList) SetTaskId(v string) *PremiumGetTodoTasksResponseBodyResultList {
	s.TaskId = &v
	return s
}

func (s *PremiumGetTodoTasksResponseBodyResultList) SetTitle(v string) *PremiumGetTodoTasksResponseBodyResultList {
	s.Title = &v
	return s
}

func (s *PremiumGetTodoTasksResponseBodyResultList) SetUrl(v string) *PremiumGetTodoTasksResponseBodyResultList {
	s.Url = &v
	return s
}

type PremiumGetTodoTasksResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumGetTodoTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumGetTodoTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumGetTodoTasksResponse) GoString() string {
	return s.String()
}

func (s *PremiumGetTodoTasksResponse) SetHeaders(v map[string]*string) *PremiumGetTodoTasksResponse {
	s.Headers = v
	return s
}

func (s *PremiumGetTodoTasksResponse) SetStatusCode(v int32) *PremiumGetTodoTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumGetTodoTasksResponse) SetBody(v *PremiumGetTodoTasksResponseBody) *PremiumGetTodoTasksResponse {
	s.Body = v
	return s
}

type PremiumGrantProcessInstanceForDownloadFileHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumGrantProcessInstanceForDownloadFileHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumGrantProcessInstanceForDownloadFileHeaders) GoString() string {
	return s.String()
}

func (s *PremiumGrantProcessInstanceForDownloadFileHeaders) SetCommonHeaders(v map[string]*string) *PremiumGrantProcessInstanceForDownloadFileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumGrantProcessInstanceForDownloadFileHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumGrantProcessInstanceForDownloadFileHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumGrantProcessInstanceForDownloadFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 111
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a17444d1-075b-4a4d-xxxx
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// if can be null:
	// true
	WithCommentAttatchment *bool `json:"withCommentAttatchment,omitempty" xml:"withCommentAttatchment,omitempty"`
}

func (s PremiumGrantProcessInstanceForDownloadFileRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumGrantProcessInstanceForDownloadFileRequest) GoString() string {
	return s.String()
}

func (s *PremiumGrantProcessInstanceForDownloadFileRequest) SetFileId(v string) *PremiumGrantProcessInstanceForDownloadFileRequest {
	s.FileId = &v
	return s
}

func (s *PremiumGrantProcessInstanceForDownloadFileRequest) SetProcessInstanceId(v string) *PremiumGrantProcessInstanceForDownloadFileRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *PremiumGrantProcessInstanceForDownloadFileRequest) SetWithCommentAttatchment(v bool) *PremiumGrantProcessInstanceForDownloadFileRequest {
	s.WithCommentAttatchment = &v
	return s
}

type PremiumGrantProcessInstanceForDownloadFileResponseBody struct {
	Result *PremiumGrantProcessInstanceForDownloadFileResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PremiumGrantProcessInstanceForDownloadFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumGrantProcessInstanceForDownloadFileResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumGrantProcessInstanceForDownloadFileResponseBody) SetResult(v *PremiumGrantProcessInstanceForDownloadFileResponseBodyResult) *PremiumGrantProcessInstanceForDownloadFileResponseBody {
	s.Result = v
	return s
}

func (s *PremiumGrantProcessInstanceForDownloadFileResponseBody) SetSuccess(v bool) *PremiumGrantProcessInstanceForDownloadFileResponseBody {
	s.Success = &v
	return s
}

type PremiumGrantProcessInstanceForDownloadFileResponseBodyResult struct {
	// example:
	//
	// http://lippi-space-zjk.oss-cn-zhangjiakou.aliyuncs.com/xxxxx
	DownloadUri *string `json:"downloadUri,omitempty" xml:"downloadUri,omitempty"`
	// example:
	//
	// 26748422566
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// example:
	//
	// 3996960664
	SpaceId *int64 `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s PremiumGrantProcessInstanceForDownloadFileResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PremiumGrantProcessInstanceForDownloadFileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PremiumGrantProcessInstanceForDownloadFileResponseBodyResult) SetDownloadUri(v string) *PremiumGrantProcessInstanceForDownloadFileResponseBodyResult {
	s.DownloadUri = &v
	return s
}

func (s *PremiumGrantProcessInstanceForDownloadFileResponseBodyResult) SetFileId(v string) *PremiumGrantProcessInstanceForDownloadFileResponseBodyResult {
	s.FileId = &v
	return s
}

func (s *PremiumGrantProcessInstanceForDownloadFileResponseBodyResult) SetSpaceId(v int64) *PremiumGrantProcessInstanceForDownloadFileResponseBodyResult {
	s.SpaceId = &v
	return s
}

type PremiumGrantProcessInstanceForDownloadFileResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumGrantProcessInstanceForDownloadFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumGrantProcessInstanceForDownloadFileResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumGrantProcessInstanceForDownloadFileResponse) GoString() string {
	return s.String()
}

func (s *PremiumGrantProcessInstanceForDownloadFileResponse) SetHeaders(v map[string]*string) *PremiumGrantProcessInstanceForDownloadFileResponse {
	s.Headers = v
	return s
}

func (s *PremiumGrantProcessInstanceForDownloadFileResponse) SetStatusCode(v int32) *PremiumGrantProcessInstanceForDownloadFileResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumGrantProcessInstanceForDownloadFileResponse) SetBody(v *PremiumGrantProcessInstanceForDownloadFileResponseBody) *PremiumGrantProcessInstanceForDownloadFileResponse {
	s.Body = v
	return s
}

type PremiumInsertOrUpdateDirHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumInsertOrUpdateDirHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumInsertOrUpdateDirHeaders) GoString() string {
	return s.String()
}

func (s *PremiumInsertOrUpdateDirHeaders) SetCommonHeaders(v map[string]*string) *PremiumInsertOrUpdateDirHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumInsertOrUpdateDirHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumInsertOrUpdateDirHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumInsertOrUpdateDirRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// administeration
	BizGroup *string `json:"bizGroup,omitempty" xml:"bizGroup,omitempty"`
	// example:
	//
	// 分组描述信息
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 行政管理
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {\"en_US\":\"test\",\"ja_JP\":\"test\",\"vi_VN\":\"test\",\"zh_CN\":\"测试\",\"zh_HK\":\"测试\",\"zh_TW\":\"测试\"}
	Name18n *string `json:"name18n,omitempty" xml:"name18n,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user001
	OperateUserId *string `json:"operateUserId,omitempty" xml:"operateUserId,omitempty"`
}

func (s PremiumInsertOrUpdateDirRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumInsertOrUpdateDirRequest) GoString() string {
	return s.String()
}

func (s *PremiumInsertOrUpdateDirRequest) SetBizGroup(v string) *PremiumInsertOrUpdateDirRequest {
	s.BizGroup = &v
	return s
}

func (s *PremiumInsertOrUpdateDirRequest) SetDescription(v string) *PremiumInsertOrUpdateDirRequest {
	s.Description = &v
	return s
}

func (s *PremiumInsertOrUpdateDirRequest) SetName(v string) *PremiumInsertOrUpdateDirRequest {
	s.Name = &v
	return s
}

func (s *PremiumInsertOrUpdateDirRequest) SetName18n(v string) *PremiumInsertOrUpdateDirRequest {
	s.Name18n = &v
	return s
}

func (s *PremiumInsertOrUpdateDirRequest) SetOperateUserId(v string) *PremiumInsertOrUpdateDirRequest {
	s.OperateUserId = &v
	return s
}

type PremiumInsertOrUpdateDirResponseBody struct {
	Result  *PremiumInsertOrUpdateDirResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                       `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PremiumInsertOrUpdateDirResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumInsertOrUpdateDirResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumInsertOrUpdateDirResponseBody) SetResult(v *PremiumInsertOrUpdateDirResponseBodyResult) *PremiumInsertOrUpdateDirResponseBody {
	s.Result = v
	return s
}

func (s *PremiumInsertOrUpdateDirResponseBody) SetSuccess(v bool) *PremiumInsertOrUpdateDirResponseBody {
	s.Success = &v
	return s
}

type PremiumInsertOrUpdateDirResponseBodyResult struct {
	// example:
	//
	// {应用appId}_administeration
	BizGroup *string `json:"bizGroup,omitempty" xml:"bizGroup,omitempty"`
	// example:
	//
	// oaDirIdxxx
	DirId *string `json:"dirId,omitempty" xml:"dirId,omitempty"`
}

func (s PremiumInsertOrUpdateDirResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PremiumInsertOrUpdateDirResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PremiumInsertOrUpdateDirResponseBodyResult) SetBizGroup(v string) *PremiumInsertOrUpdateDirResponseBodyResult {
	s.BizGroup = &v
	return s
}

func (s *PremiumInsertOrUpdateDirResponseBodyResult) SetDirId(v string) *PremiumInsertOrUpdateDirResponseBodyResult {
	s.DirId = &v
	return s
}

type PremiumInsertOrUpdateDirResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumInsertOrUpdateDirResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumInsertOrUpdateDirResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumInsertOrUpdateDirResponse) GoString() string {
	return s.String()
}

func (s *PremiumInsertOrUpdateDirResponse) SetHeaders(v map[string]*string) *PremiumInsertOrUpdateDirResponse {
	s.Headers = v
	return s
}

func (s *PremiumInsertOrUpdateDirResponse) SetStatusCode(v int32) *PremiumInsertOrUpdateDirResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumInsertOrUpdateDirResponse) SetBody(v *PremiumInsertOrUpdateDirResponseBody) *PremiumInsertOrUpdateDirResponse {
	s.Body = v
	return s
}

type PremiumQuerySchemaAndProcessByCodeListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumQuerySchemaAndProcessByCodeListHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumQuerySchemaAndProcessByCodeListHeaders) GoString() string {
	return s.String()
}

func (s *PremiumQuerySchemaAndProcessByCodeListHeaders) SetCommonHeaders(v map[string]*string) *PremiumQuerySchemaAndProcessByCodeListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumQuerySchemaAndProcessByCodeListHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumQuerySchemaAndProcessByCodeListHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumQuerySchemaAndProcessByCodeListRequest struct {
	// This parameter is required.
	//
	// if can be null:
	// false
	ProcessCodes []*string `json:"processCodes,omitempty" xml:"processCodes,omitempty" type:"Repeated"`
}

func (s PremiumQuerySchemaAndProcessByCodeListRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumQuerySchemaAndProcessByCodeListRequest) GoString() string {
	return s.String()
}

func (s *PremiumQuerySchemaAndProcessByCodeListRequest) SetProcessCodes(v []*string) *PremiumQuerySchemaAndProcessByCodeListRequest {
	s.ProcessCodes = v
	return s
}

type PremiumQuerySchemaAndProcessByCodeListResponseBody struct {
	Result []*PremiumQuerySchemaAndProcessByCodeListResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PremiumQuerySchemaAndProcessByCodeListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumQuerySchemaAndProcessByCodeListResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumQuerySchemaAndProcessByCodeListResponseBody) SetResult(v []*PremiumQuerySchemaAndProcessByCodeListResponseBodyResult) *PremiumQuerySchemaAndProcessByCodeListResponseBody {
	s.Result = v
	return s
}

func (s *PremiumQuerySchemaAndProcessByCodeListResponseBody) SetSuccess(v bool) *PremiumQuerySchemaAndProcessByCodeListResponseBody {
	s.Success = &v
	return s
}

type PremiumQuerySchemaAndProcessByCodeListResponseBodyResult struct {
	// example:
	//
	// ding123
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// example:
	//
	// hrm.xxx
	BizCategoryId *string `json:"bizCategoryId,omitempty" xml:"bizCategoryId,omitempty"`
	// example:
	//
	// 1638326995000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// userId123
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// example:
	//
	// FORM-28215C3E-00E3-4118-xxxx-4091F828AF2F
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// https//:xxx
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// 模板描述1
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// example:
	//
	// userId123
	ModifierUserId *string `json:"modifierUserId,omitempty" xml:"modifierUserId,omitempty"`
	// example:
	//
	// 1638326995000
	ModifyTime *int64 `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// example:
	//
	// 示例模板
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// PROC-17428B8C-6C60-470E-xxxx-64F1037AE067
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// example:
	//
	// {\"name\":\"发起人\",\"type\":\"start\",\"nodeId\":\"sid-startevent\",\"childNode\":{\"name\":\"审批人\",\"prevId\":\"sid-startevent\",\"type\":\"approver\",\"nodeId\":\"sid-1234_5678\",\"properties\":{\"activateType\":\"ONE_BY_ONE\",\"approvalType\":\"MANUAL\",\"actionerRules\":[{\"select\":[\"allStaff\"],\"range\":{\"approvals\":[],\"labels\":[]},\"type\":\"target_select\",\"key\":\"manual_sid-1234_5678_30a8_b373\",\"multi\":1}],\"agreeAll\":false},\"childNode\":{\"name\":\"抄送人\",\"prevId\":\"sid-1234_5678\",\"type\":\"notifier\",\"nodeId\":\"9955_7e43\",\"properties\":{\"actionerRules\":[{\"select\":[\"allStaff\"],\"range\":{},\"type\":\"target_select\",\"key\":\"manual_9955_7e43_0c96_39d8\",\"multi\":1}]}}}}
	ProcessConfig *string `json:"processConfig,omitempty" xml:"processConfig,omitempty"`
	ProcessId     *int64  `json:"processId,omitempty" xml:"processId,omitempty"`
	// example:
	//
	// {\"commentHiddenForProposer\":\"\",\"commentRequired\":\"\",\"icon\":\"timefades#red\",\"commentDescription\":\"\",\"description\":\"支持地址控件\",\"title\":\"官方OA审批-POP-2025-0109\",\"items\":[{\"componentName\":\"TimeAndLocationField\",\"props\":{\"label\":[\"当前时间\",\"当前地点\"],\"id\":\"TimeAndLocationField_1CVHM5TIIWR9C\",\"required\":false}},{\"componentName\":\"TextField\",\"props\":{\"placeholder\":\"请输入\",\"label\":\"单行输入框\",\"id\":\"TextField_17EZKEGSOCTC0\",\"required\":false}}]}
	SchemaContent *string `json:"schemaContent,omitempty" xml:"schemaContent,omitempty"`
	// example:
	//
	// PUBLISHED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PremiumQuerySchemaAndProcessByCodeListResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PremiumQuerySchemaAndProcessByCodeListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult) SetAppUuid(v string) *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult {
	s.AppUuid = &v
	return s
}

func (s *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult) SetBizCategoryId(v string) *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult {
	s.BizCategoryId = &v
	return s
}

func (s *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult) SetCreateTime(v int64) *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult) SetCreatorUserId(v string) *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult {
	s.CreatorUserId = &v
	return s
}

func (s *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult) SetFormUuid(v string) *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult {
	s.FormUuid = &v
	return s
}

func (s *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult) SetIcon(v string) *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult {
	s.Icon = &v
	return s
}

func (s *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult) SetMemo(v string) *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult {
	s.Memo = &v
	return s
}

func (s *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult) SetModifierUserId(v string) *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult {
	s.ModifierUserId = &v
	return s
}

func (s *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult) SetModifyTime(v int64) *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult {
	s.ModifyTime = &v
	return s
}

func (s *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult) SetName(v string) *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult {
	s.Name = &v
	return s
}

func (s *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult) SetProcessCode(v string) *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult {
	s.ProcessCode = &v
	return s
}

func (s *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult) SetProcessConfig(v string) *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult {
	s.ProcessConfig = &v
	return s
}

func (s *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult) SetProcessId(v int64) *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult {
	s.ProcessId = &v
	return s
}

func (s *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult) SetSchemaContent(v string) *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult {
	s.SchemaContent = &v
	return s
}

func (s *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult) SetStatus(v string) *PremiumQuerySchemaAndProcessByCodeListResponseBodyResult {
	s.Status = &v
	return s
}

type PremiumQuerySchemaAndProcessByCodeListResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumQuerySchemaAndProcessByCodeListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumQuerySchemaAndProcessByCodeListResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumQuerySchemaAndProcessByCodeListResponse) GoString() string {
	return s.String()
}

func (s *PremiumQuerySchemaAndProcessByCodeListResponse) SetHeaders(v map[string]*string) *PremiumQuerySchemaAndProcessByCodeListResponse {
	s.Headers = v
	return s
}

func (s *PremiumQuerySchemaAndProcessByCodeListResponse) SetStatusCode(v int32) *PremiumQuerySchemaAndProcessByCodeListResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumQuerySchemaAndProcessByCodeListResponse) SetBody(v *PremiumQuerySchemaAndProcessByCodeListResponseBody) *PremiumQuerySchemaAndProcessByCodeListResponse {
	s.Body = v
	return s
}

type PremiumQueryTodoTasksByManagerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumQueryTodoTasksByManagerHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumQueryTodoTasksByManagerHeaders) GoString() string {
	return s.String()
}

func (s *PremiumQueryTodoTasksByManagerHeaders) SetCommonHeaders(v map[string]*string) *PremiumQueryTodoTasksByManagerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumQueryTodoTasksByManagerHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumQueryTodoTasksByManagerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumQueryTodoTasksByManagerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// staffId123
	ActionerUserId *string `json:"actionerUserId,omitempty" xml:"actionerUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// manager123
	ManagerUserId *string `json:"managerUserId,omitempty" xml:"managerUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	NextToken *int32 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s PremiumQueryTodoTasksByManagerRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumQueryTodoTasksByManagerRequest) GoString() string {
	return s.String()
}

func (s *PremiumQueryTodoTasksByManagerRequest) SetActionerUserId(v string) *PremiumQueryTodoTasksByManagerRequest {
	s.ActionerUserId = &v
	return s
}

func (s *PremiumQueryTodoTasksByManagerRequest) SetManagerUserId(v string) *PremiumQueryTodoTasksByManagerRequest {
	s.ManagerUserId = &v
	return s
}

func (s *PremiumQueryTodoTasksByManagerRequest) SetMaxResults(v int32) *PremiumQueryTodoTasksByManagerRequest {
	s.MaxResults = &v
	return s
}

func (s *PremiumQueryTodoTasksByManagerRequest) SetNextToken(v int32) *PremiumQueryTodoTasksByManagerRequest {
	s.NextToken = &v
	return s
}

type PremiumQueryTodoTasksByManagerResponseBody struct {
	Result *PremiumQueryTodoTasksByManagerResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s PremiumQueryTodoTasksByManagerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumQueryTodoTasksByManagerResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumQueryTodoTasksByManagerResponseBody) SetResult(v *PremiumQueryTodoTasksByManagerResponseBodyResult) *PremiumQueryTodoTasksByManagerResponseBody {
	s.Result = v
	return s
}

type PremiumQueryTodoTasksByManagerResponseBodyResult struct {
	HasMore *bool                                                   `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List    []*PremiumQueryTodoTasksByManagerResponseBodyResultList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s PremiumQueryTodoTasksByManagerResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PremiumQueryTodoTasksByManagerResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PremiumQueryTodoTasksByManagerResponseBodyResult) SetHasMore(v bool) *PremiumQueryTodoTasksByManagerResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *PremiumQueryTodoTasksByManagerResponseBodyResult) SetList(v []*PremiumQueryTodoTasksByManagerResponseBodyResultList) *PremiumQueryTodoTasksByManagerResponseBodyResult {
	s.List = v
	return s
}

type PremiumQueryTodoTasksByManagerResponseBodyResultList struct {
	// example:
	//
	// RUNNING
	BusinessId  *string `json:"businessId,omitempty" xml:"businessId,omitempty"`
	CanRedirect *bool   `json:"canRedirect,omitempty" xml:"canRedirect,omitempty"`
	CreateTime  *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// act_0001
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// example:
	//
	// Siw2WNVZS4KiUt3tTmaNKg04*****809950
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// example:
	//
	// 1234567
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// manager001
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 2022-10-17T15:12Z
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PremiumQueryTodoTasksByManagerResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s PremiumQueryTodoTasksByManagerResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *PremiumQueryTodoTasksByManagerResponseBodyResultList) SetBusinessId(v string) *PremiumQueryTodoTasksByManagerResponseBodyResultList {
	s.BusinessId = &v
	return s
}

func (s *PremiumQueryTodoTasksByManagerResponseBodyResultList) SetCanRedirect(v bool) *PremiumQueryTodoTasksByManagerResponseBodyResultList {
	s.CanRedirect = &v
	return s
}

func (s *PremiumQueryTodoTasksByManagerResponseBodyResultList) SetCreateTime(v int64) *PremiumQueryTodoTasksByManagerResponseBodyResultList {
	s.CreateTime = &v
	return s
}

func (s *PremiumQueryTodoTasksByManagerResponseBodyResultList) SetProcessCode(v string) *PremiumQueryTodoTasksByManagerResponseBodyResultList {
	s.ProcessCode = &v
	return s
}

func (s *PremiumQueryTodoTasksByManagerResponseBodyResultList) SetProcessInstanceId(v string) *PremiumQueryTodoTasksByManagerResponseBodyResultList {
	s.ProcessInstanceId = &v
	return s
}

func (s *PremiumQueryTodoTasksByManagerResponseBodyResultList) SetTaskId(v int64) *PremiumQueryTodoTasksByManagerResponseBodyResultList {
	s.TaskId = &v
	return s
}

func (s *PremiumQueryTodoTasksByManagerResponseBodyResultList) SetTitle(v string) *PremiumQueryTodoTasksByManagerResponseBodyResultList {
	s.Title = &v
	return s
}

func (s *PremiumQueryTodoTasksByManagerResponseBodyResultList) SetUserId(v string) *PremiumQueryTodoTasksByManagerResponseBodyResultList {
	s.UserId = &v
	return s
}

type PremiumQueryTodoTasksByManagerResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumQueryTodoTasksByManagerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumQueryTodoTasksByManagerResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumQueryTodoTasksByManagerResponse) GoString() string {
	return s.String()
}

func (s *PremiumQueryTodoTasksByManagerResponse) SetHeaders(v map[string]*string) *PremiumQueryTodoTasksByManagerResponse {
	s.Headers = v
	return s
}

func (s *PremiumQueryTodoTasksByManagerResponse) SetStatusCode(v int32) *PremiumQueryTodoTasksByManagerResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumQueryTodoTasksByManagerResponse) SetBody(v *PremiumQueryTodoTasksByManagerResponseBody) *PremiumQueryTodoTasksByManagerResponse {
	s.Body = v
	return s
}

type PremiumRedirectTasksByManagerHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumRedirectTasksByManagerHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumRedirectTasksByManagerHeaders) GoString() string {
	return s.String()
}

func (s *PremiumRedirectTasksByManagerHeaders) SetCommonHeaders(v map[string]*string) *PremiumRedirectTasksByManagerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumRedirectTasksByManagerHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumRedirectTasksByManagerHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumRedirectTasksByManagerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// staffId-B
	HandoverUserId *string `json:"handoverUserId,omitempty" xml:"handoverUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// manager-12
	ManagerUserId *string `json:"managerUserId,omitempty" xml:"managerUserId,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	TaskIds []*int64 `json:"taskIds,omitempty" xml:"taskIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// staffId-A
	TransfereeUserId *string `json:"transfereeUserId,omitempty" xml:"transfereeUserId,omitempty"`
}

func (s PremiumRedirectTasksByManagerRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumRedirectTasksByManagerRequest) GoString() string {
	return s.String()
}

func (s *PremiumRedirectTasksByManagerRequest) SetHandoverUserId(v string) *PremiumRedirectTasksByManagerRequest {
	s.HandoverUserId = &v
	return s
}

func (s *PremiumRedirectTasksByManagerRequest) SetManagerUserId(v string) *PremiumRedirectTasksByManagerRequest {
	s.ManagerUserId = &v
	return s
}

func (s *PremiumRedirectTasksByManagerRequest) SetTaskIds(v []*int64) *PremiumRedirectTasksByManagerRequest {
	s.TaskIds = v
	return s
}

func (s *PremiumRedirectTasksByManagerRequest) SetTransfereeUserId(v string) *PremiumRedirectTasksByManagerRequest {
	s.TransfereeUserId = &v
	return s
}

type PremiumRedirectTasksByManagerResponseBody struct {
	Result  *PremiumRedirectTasksByManagerResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Success *bool                                            `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PremiumRedirectTasksByManagerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumRedirectTasksByManagerResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumRedirectTasksByManagerResponseBody) SetResult(v *PremiumRedirectTasksByManagerResponseBodyResult) *PremiumRedirectTasksByManagerResponseBody {
	s.Result = v
	return s
}

func (s *PremiumRedirectTasksByManagerResponseBody) SetSuccess(v bool) *PremiumRedirectTasksByManagerResponseBody {
	s.Success = &v
	return s
}

type PremiumRedirectTasksByManagerResponseBodyResult struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	FailCount *int64 `json:"failCount,omitempty" xml:"failCount,omitempty"`
	// This parameter is required.
	RedirectResults []*PremiumRedirectTasksByManagerResponseBodyResultRedirectResults `json:"redirectResults,omitempty" xml:"redirectResults,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s PremiumRedirectTasksByManagerResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PremiumRedirectTasksByManagerResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PremiumRedirectTasksByManagerResponseBodyResult) SetFailCount(v int64) *PremiumRedirectTasksByManagerResponseBodyResult {
	s.FailCount = &v
	return s
}

func (s *PremiumRedirectTasksByManagerResponseBodyResult) SetRedirectResults(v []*PremiumRedirectTasksByManagerResponseBodyResultRedirectResults) *PremiumRedirectTasksByManagerResponseBodyResult {
	s.RedirectResults = v
	return s
}

func (s *PremiumRedirectTasksByManagerResponseBodyResult) SetTotalCount(v int64) *PremiumRedirectTasksByManagerResponseBodyResult {
	s.TotalCount = &v
	return s
}

type PremiumRedirectTasksByManagerResponseBodyResultRedirectResults struct {
	// example:
	//
	// 外部流程不允许转交
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s PremiumRedirectTasksByManagerResponseBodyResultRedirectResults) String() string {
	return tea.Prettify(s)
}

func (s PremiumRedirectTasksByManagerResponseBodyResultRedirectResults) GoString() string {
	return s.String()
}

func (s *PremiumRedirectTasksByManagerResponseBodyResultRedirectResults) SetErrorMsg(v string) *PremiumRedirectTasksByManagerResponseBodyResultRedirectResults {
	s.ErrorMsg = &v
	return s
}

func (s *PremiumRedirectTasksByManagerResponseBodyResultRedirectResults) SetSuccess(v bool) *PremiumRedirectTasksByManagerResponseBodyResultRedirectResults {
	s.Success = &v
	return s
}

func (s *PremiumRedirectTasksByManagerResponseBodyResultRedirectResults) SetTaskId(v int64) *PremiumRedirectTasksByManagerResponseBodyResultRedirectResults {
	s.TaskId = &v
	return s
}

type PremiumRedirectTasksByManagerResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumRedirectTasksByManagerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumRedirectTasksByManagerResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumRedirectTasksByManagerResponse) GoString() string {
	return s.String()
}

func (s *PremiumRedirectTasksByManagerResponse) SetHeaders(v map[string]*string) *PremiumRedirectTasksByManagerResponse {
	s.Headers = v
	return s
}

func (s *PremiumRedirectTasksByManagerResponse) SetStatusCode(v int32) *PremiumRedirectTasksByManagerResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumRedirectTasksByManagerResponse) SetBody(v *PremiumRedirectTasksByManagerResponseBody) *PremiumRedirectTasksByManagerResponse {
	s.Body = v
	return s
}

type PremiumRevertTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumRevertTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumRevertTaskHeaders) GoString() string {
	return s.String()
}

func (s *PremiumRevertTaskHeaders) SetCommonHeaders(v map[string]*string) *PremiumRevertTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumRevertTaskHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumRevertTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumRevertTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// manager001
	OperateUserId *string `json:"operateUserId,omitempty" xml:"operateUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// processInstanceId123
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// example:
	//
	// 退回到审批人（上一步）
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// REVERT_FOR_APPROVAL
	RevertAction *string `json:"revertAction,omitempty" xml:"revertAction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// d3aa_1974
	TargetActivityId *string `json:"targetActivityId,omitempty" xml:"targetActivityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s PremiumRevertTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumRevertTaskRequest) GoString() string {
	return s.String()
}

func (s *PremiumRevertTaskRequest) SetOperateUserId(v string) *PremiumRevertTaskRequest {
	s.OperateUserId = &v
	return s
}

func (s *PremiumRevertTaskRequest) SetProcessInstanceId(v string) *PremiumRevertTaskRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *PremiumRevertTaskRequest) SetRemark(v string) *PremiumRevertTaskRequest {
	s.Remark = &v
	return s
}

func (s *PremiumRevertTaskRequest) SetRevertAction(v string) *PremiumRevertTaskRequest {
	s.RevertAction = &v
	return s
}

func (s *PremiumRevertTaskRequest) SetTargetActivityId(v string) *PremiumRevertTaskRequest {
	s.TargetActivityId = &v
	return s
}

func (s *PremiumRevertTaskRequest) SetTaskId(v int64) *PremiumRevertTaskRequest {
	s.TaskId = &v
	return s
}

type PremiumRevertTaskResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s PremiumRevertTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumRevertTaskResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumRevertTaskResponseBody) SetResult(v bool) *PremiumRevertTaskResponseBody {
	s.Result = &v
	return s
}

type PremiumRevertTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumRevertTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumRevertTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumRevertTaskResponse) GoString() string {
	return s.String()
}

func (s *PremiumRevertTaskResponse) SetHeaders(v map[string]*string) *PremiumRevertTaskResponse {
	s.Headers = v
	return s
}

func (s *PremiumRevertTaskResponse) SetStatusCode(v int32) *PremiumRevertTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumRevertTaskResponse) SetBody(v *PremiumRevertTaskResponseBody) *PremiumRevertTaskResponse {
	s.Body = v
	return s
}

type PremiumSaveFormHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumSaveFormHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveFormHeaders) GoString() string {
	return s.String()
}

func (s *PremiumSaveFormHeaders) SetCommonHeaders(v map[string]*string) *PremiumSaveFormHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumSaveFormHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumSaveFormHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumSaveFormRequest struct {
	// example:
	//
	// 用于员工差旅费用报销使用
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	FormComponents []*FormComponent `json:"formComponents,omitempty" xml:"formComponents,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 出差报销审批
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// proc-abc
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PremiumSaveFormRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveFormRequest) GoString() string {
	return s.String()
}

func (s *PremiumSaveFormRequest) SetDescription(v string) *PremiumSaveFormRequest {
	s.Description = &v
	return s
}

func (s *PremiumSaveFormRequest) SetFormComponents(v []*FormComponent) *PremiumSaveFormRequest {
	s.FormComponents = v
	return s
}

func (s *PremiumSaveFormRequest) SetName(v string) *PremiumSaveFormRequest {
	s.Name = &v
	return s
}

func (s *PremiumSaveFormRequest) SetProcessCode(v string) *PremiumSaveFormRequest {
	s.ProcessCode = &v
	return s
}

func (s *PremiumSaveFormRequest) SetUserId(v string) *PremiumSaveFormRequest {
	s.UserId = &v
	return s
}

type PremiumSaveFormResponseBody struct {
	Result *PremiumSaveFormResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s PremiumSaveFormResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveFormResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumSaveFormResponseBody) SetResult(v *PremiumSaveFormResponseBodyResult) *PremiumSaveFormResponseBody {
	s.Result = v
	return s
}

type PremiumSaveFormResponseBodyResult struct {
	// This parameter is required.
	//
	// example:
	//
	// PROC-abcdef-example
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
}

func (s PremiumSaveFormResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveFormResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PremiumSaveFormResponseBodyResult) SetProcessCode(v string) *PremiumSaveFormResponseBodyResult {
	s.ProcessCode = &v
	return s
}

type PremiumSaveFormResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumSaveFormResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumSaveFormResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveFormResponse) GoString() string {
	return s.String()
}

func (s *PremiumSaveFormResponse) SetHeaders(v map[string]*string) *PremiumSaveFormResponse {
	s.Headers = v
	return s
}

func (s *PremiumSaveFormResponse) SetStatusCode(v int32) *PremiumSaveFormResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumSaveFormResponse) SetBody(v *PremiumSaveFormResponseBody) *PremiumSaveFormResponse {
	s.Body = v
	return s
}

type PremiumSaveFormInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumSaveFormInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveFormInstanceHeaders) GoString() string {
	return s.String()
}

func (s *PremiumSaveFormInstanceHeaders) SetCommonHeaders(v map[string]*string) *PremiumSaveFormInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumSaveFormInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumSaveFormInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumSaveFormInstanceRequest struct {
	// This parameter is required.
	FormComponentValueList []*PremiumSaveFormInstanceRequestFormComponentValueList `json:"formComponentValueList,omitempty" xml:"formComponentValueList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// manager432
	OriginatorUserId *string `json:"originatorUserId,omitempty" xml:"originatorUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-EF6YJL35P2-SCKICSB7P750S0YISYKV3-xxxx-1
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
}

func (s PremiumSaveFormInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveFormInstanceRequest) GoString() string {
	return s.String()
}

func (s *PremiumSaveFormInstanceRequest) SetFormComponentValueList(v []*PremiumSaveFormInstanceRequestFormComponentValueList) *PremiumSaveFormInstanceRequest {
	s.FormComponentValueList = v
	return s
}

func (s *PremiumSaveFormInstanceRequest) SetOriginatorUserId(v string) *PremiumSaveFormInstanceRequest {
	s.OriginatorUserId = &v
	return s
}

func (s *PremiumSaveFormInstanceRequest) SetProcessCode(v string) *PremiumSaveFormInstanceRequest {
	s.ProcessCode = &v
	return s
}

type PremiumSaveFormInstanceRequestFormComponentValueList struct {
	// example:
	//
	// Phone
	BizAlias      *string                                                        `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	ComponentType *string                                                        `json:"componentType,omitempty" xml:"componentType,omitempty"`
	Details       []*PremiumSaveFormInstanceRequestFormComponentValueListDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
	// example:
	//
	// 总个数:1
	ExtValue *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	// example:
	//
	// PhoneField_IZI2LP8QF6O0
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PhoneField
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123xxxxxxxx
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s PremiumSaveFormInstanceRequestFormComponentValueList) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveFormInstanceRequestFormComponentValueList) GoString() string {
	return s.String()
}

func (s *PremiumSaveFormInstanceRequestFormComponentValueList) SetBizAlias(v string) *PremiumSaveFormInstanceRequestFormComponentValueList {
	s.BizAlias = &v
	return s
}

func (s *PremiumSaveFormInstanceRequestFormComponentValueList) SetComponentType(v string) *PremiumSaveFormInstanceRequestFormComponentValueList {
	s.ComponentType = &v
	return s
}

func (s *PremiumSaveFormInstanceRequestFormComponentValueList) SetDetails(v []*PremiumSaveFormInstanceRequestFormComponentValueListDetails) *PremiumSaveFormInstanceRequestFormComponentValueList {
	s.Details = v
	return s
}

func (s *PremiumSaveFormInstanceRequestFormComponentValueList) SetExtValue(v string) *PremiumSaveFormInstanceRequestFormComponentValueList {
	s.ExtValue = &v
	return s
}

func (s *PremiumSaveFormInstanceRequestFormComponentValueList) SetId(v string) *PremiumSaveFormInstanceRequestFormComponentValueList {
	s.Id = &v
	return s
}

func (s *PremiumSaveFormInstanceRequestFormComponentValueList) SetName(v string) *PremiumSaveFormInstanceRequestFormComponentValueList {
	s.Name = &v
	return s
}

func (s *PremiumSaveFormInstanceRequestFormComponentValueList) SetValue(v string) *PremiumSaveFormInstanceRequestFormComponentValueList {
	s.Value = &v
	return s
}

type PremiumSaveFormInstanceRequestFormComponentValueListDetails struct {
	// example:
	//
	// Phone
	BizAlias *string                                                               `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	Details  []*PremiumSaveFormInstanceRequestFormComponentValueListDetailsDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
	// example:
	//
	// 总个数:1
	ExtValue *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	// example:
	//
	// PhoneField_IZI2LP8QF6O0
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// PhoneField
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 123xxxxxxxx
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s PremiumSaveFormInstanceRequestFormComponentValueListDetails) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveFormInstanceRequestFormComponentValueListDetails) GoString() string {
	return s.String()
}

func (s *PremiumSaveFormInstanceRequestFormComponentValueListDetails) SetBizAlias(v string) *PremiumSaveFormInstanceRequestFormComponentValueListDetails {
	s.BizAlias = &v
	return s
}

func (s *PremiumSaveFormInstanceRequestFormComponentValueListDetails) SetDetails(v []*PremiumSaveFormInstanceRequestFormComponentValueListDetailsDetails) *PremiumSaveFormInstanceRequestFormComponentValueListDetails {
	s.Details = v
	return s
}

func (s *PremiumSaveFormInstanceRequestFormComponentValueListDetails) SetExtValue(v string) *PremiumSaveFormInstanceRequestFormComponentValueListDetails {
	s.ExtValue = &v
	return s
}

func (s *PremiumSaveFormInstanceRequestFormComponentValueListDetails) SetId(v string) *PremiumSaveFormInstanceRequestFormComponentValueListDetails {
	s.Id = &v
	return s
}

func (s *PremiumSaveFormInstanceRequestFormComponentValueListDetails) SetName(v string) *PremiumSaveFormInstanceRequestFormComponentValueListDetails {
	s.Name = &v
	return s
}

func (s *PremiumSaveFormInstanceRequestFormComponentValueListDetails) SetValue(v string) *PremiumSaveFormInstanceRequestFormComponentValueListDetails {
	s.Value = &v
	return s
}

type PremiumSaveFormInstanceRequestFormComponentValueListDetailsDetails struct {
	// example:
	//
	// Phone
	BizAlias      *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// example:
	//
	// 总个数:1
	ExtValue *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	// example:
	//
	// PhoneField_IZI2LP8QF6O0
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// PhoneField
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 123xxxxxxxx
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s PremiumSaveFormInstanceRequestFormComponentValueListDetailsDetails) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveFormInstanceRequestFormComponentValueListDetailsDetails) GoString() string {
	return s.String()
}

func (s *PremiumSaveFormInstanceRequestFormComponentValueListDetailsDetails) SetBizAlias(v string) *PremiumSaveFormInstanceRequestFormComponentValueListDetailsDetails {
	s.BizAlias = &v
	return s
}

func (s *PremiumSaveFormInstanceRequestFormComponentValueListDetailsDetails) SetComponentType(v string) *PremiumSaveFormInstanceRequestFormComponentValueListDetailsDetails {
	s.ComponentType = &v
	return s
}

func (s *PremiumSaveFormInstanceRequestFormComponentValueListDetailsDetails) SetExtValue(v string) *PremiumSaveFormInstanceRequestFormComponentValueListDetailsDetails {
	s.ExtValue = &v
	return s
}

func (s *PremiumSaveFormInstanceRequestFormComponentValueListDetailsDetails) SetId(v string) *PremiumSaveFormInstanceRequestFormComponentValueListDetailsDetails {
	s.Id = &v
	return s
}

func (s *PremiumSaveFormInstanceRequestFormComponentValueListDetailsDetails) SetName(v string) *PremiumSaveFormInstanceRequestFormComponentValueListDetailsDetails {
	s.Name = &v
	return s
}

func (s *PremiumSaveFormInstanceRequestFormComponentValueListDetailsDetails) SetValue(v string) *PremiumSaveFormInstanceRequestFormComponentValueListDetailsDetails {
	s.Value = &v
	return s
}

type PremiumSaveFormInstanceResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 91ef1076-c3ed-4a78-a7a5-fa29ef2d6252
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s PremiumSaveFormInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveFormInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumSaveFormInstanceResponseBody) SetInstanceId(v string) *PremiumSaveFormInstanceResponseBody {
	s.InstanceId = &v
	return s
}

type PremiumSaveFormInstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumSaveFormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumSaveFormInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveFormInstanceResponse) GoString() string {
	return s.String()
}

func (s *PremiumSaveFormInstanceResponse) SetHeaders(v map[string]*string) *PremiumSaveFormInstanceResponse {
	s.Headers = v
	return s
}

func (s *PremiumSaveFormInstanceResponse) SetStatusCode(v int32) *PremiumSaveFormInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumSaveFormInstanceResponse) SetBody(v *PremiumSaveFormInstanceResponseBody) *PremiumSaveFormInstanceResponse {
	s.Body = v
	return s
}

type PremiumSaveIntegratedProcessHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumSaveIntegratedProcessHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedProcessHeaders) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedProcessHeaders) SetCommonHeaders(v map[string]*string) *PremiumSaveIntegratedProcessHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumSaveIntegratedProcessHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumSaveIntegratedProcessHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumSaveIntegratedProcessRequest struct {
	// example:
	//
	// 用于员工差旅费用报销使用
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	FormComponents []*FormComponent `json:"formComponents,omitempty" xml:"formComponents,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 出差报销审批
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// proc-abc
	ProcessCode          *string                                                  `json:"processCode,omitempty" xml:"processCode,omitempty"`
	ProcessFeatureConfig *PremiumSaveIntegratedProcessRequestProcessFeatureConfig `json:"processFeatureConfig,omitempty" xml:"processFeatureConfig,omitempty" type:"Struct"`
	// Deprecated
	//
	// if can be null:
	// true
	TemplateConfig *PremiumSaveIntegratedProcessRequestTemplateConfig `json:"templateConfig,omitempty" xml:"templateConfig,omitempty" type:"Struct"`
}

func (s PremiumSaveIntegratedProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedProcessRequest) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedProcessRequest) SetDescription(v string) *PremiumSaveIntegratedProcessRequest {
	s.Description = &v
	return s
}

func (s *PremiumSaveIntegratedProcessRequest) SetFormComponents(v []*FormComponent) *PremiumSaveIntegratedProcessRequest {
	s.FormComponents = v
	return s
}

func (s *PremiumSaveIntegratedProcessRequest) SetName(v string) *PremiumSaveIntegratedProcessRequest {
	s.Name = &v
	return s
}

func (s *PremiumSaveIntegratedProcessRequest) SetProcessCode(v string) *PremiumSaveIntegratedProcessRequest {
	s.ProcessCode = &v
	return s
}

func (s *PremiumSaveIntegratedProcessRequest) SetProcessFeatureConfig(v *PremiumSaveIntegratedProcessRequestProcessFeatureConfig) *PremiumSaveIntegratedProcessRequest {
	s.ProcessFeatureConfig = v
	return s
}

func (s *PremiumSaveIntegratedProcessRequest) SetTemplateConfig(v *PremiumSaveIntegratedProcessRequestTemplateConfig) *PremiumSaveIntegratedProcessRequest {
	s.TemplateConfig = v
	return s
}

type PremiumSaveIntegratedProcessRequestProcessFeatureConfig struct {
	Features []*PremiumSaveIntegratedProcessRequestProcessFeatureConfigFeatures `json:"features,omitempty" xml:"features,omitempty" type:"Repeated"`
}

func (s PremiumSaveIntegratedProcessRequestProcessFeatureConfig) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedProcessRequestProcessFeatureConfig) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedProcessRequestProcessFeatureConfig) SetFeatures(v []*PremiumSaveIntegratedProcessRequestProcessFeatureConfigFeatures) *PremiumSaveIntegratedProcessRequestProcessFeatureConfig {
	s.Features = v
	return s
}

type PremiumSaveIntegratedProcessRequestProcessFeatureConfigFeatures struct {
	Callback *PremiumSaveIntegratedProcessRequestProcessFeatureConfigFeaturesCallback `json:"callback,omitempty" xml:"callback,omitempty" type:"Struct"`
	// example:
	//
	// www.dingtalk.com
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	// example:
	//
	// abc
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// www.dingtalk.com
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	// example:
	//
	// ORIGIN
	RunType *string `json:"runType,omitempty" xml:"runType,omitempty"`
}

func (s PremiumSaveIntegratedProcessRequestProcessFeatureConfigFeatures) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedProcessRequestProcessFeatureConfigFeatures) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedProcessRequestProcessFeatureConfigFeatures) SetCallback(v *PremiumSaveIntegratedProcessRequestProcessFeatureConfigFeaturesCallback) *PremiumSaveIntegratedProcessRequestProcessFeatureConfigFeatures {
	s.Callback = v
	return s
}

func (s *PremiumSaveIntegratedProcessRequestProcessFeatureConfigFeatures) SetMobileUrl(v string) *PremiumSaveIntegratedProcessRequestProcessFeatureConfigFeatures {
	s.MobileUrl = &v
	return s
}

func (s *PremiumSaveIntegratedProcessRequestProcessFeatureConfigFeatures) SetName(v string) *PremiumSaveIntegratedProcessRequestProcessFeatureConfigFeatures {
	s.Name = &v
	return s
}

func (s *PremiumSaveIntegratedProcessRequestProcessFeatureConfigFeatures) SetPcUrl(v string) *PremiumSaveIntegratedProcessRequestProcessFeatureConfigFeatures {
	s.PcUrl = &v
	return s
}

func (s *PremiumSaveIntegratedProcessRequestProcessFeatureConfigFeatures) SetRunType(v string) *PremiumSaveIntegratedProcessRequestProcessFeatureConfigFeatures {
	s.RunType = &v
	return s
}

type PremiumSaveIntegratedProcessRequestProcessFeatureConfigFeaturesCallback struct {
	// example:
	//
	// abc
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// example:
	//
	// abc
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s PremiumSaveIntegratedProcessRequestProcessFeatureConfigFeaturesCallback) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedProcessRequestProcessFeatureConfigFeaturesCallback) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedProcessRequestProcessFeatureConfigFeaturesCallback) SetApiKey(v string) *PremiumSaveIntegratedProcessRequestProcessFeatureConfigFeaturesCallback {
	s.ApiKey = &v
	return s
}

func (s *PremiumSaveIntegratedProcessRequestProcessFeatureConfigFeaturesCallback) SetAppUuid(v string) *PremiumSaveIntegratedProcessRequestProcessFeatureConfigFeaturesCallback {
	s.AppUuid = &v
	return s
}

func (s *PremiumSaveIntegratedProcessRequestProcessFeatureConfigFeaturesCallback) SetVersion(v string) *PremiumSaveIntegratedProcessRequestProcessFeatureConfigFeaturesCallback {
	s.Version = &v
	return s
}

type PremiumSaveIntegratedProcessRequestTemplateConfig struct {
	// Deprecated
	//
	// example:
	//
	// https://open.dingtalk.com/
	CreateInstanceMobileUrl *string `json:"createInstanceMobileUrl,omitempty" xml:"createInstanceMobileUrl,omitempty"`
	// Deprecated
	//
	// example:
	//
	// https://open.dingtalk.com/
	CreateInstancePcUrl *string `json:"createInstancePcUrl,omitempty" xml:"createInstancePcUrl,omitempty"`
	// if can be null:
	// true
	DisableSendCard *bool `json:"disableSendCard,omitempty" xml:"disableSendCard,omitempty"`
	// example:
	//
	// true
	Hidden *bool `json:"hidden,omitempty" xml:"hidden,omitempty"`
	// Deprecated
	//
	// if can be null:
	// true
	//
	// example:
	//
	// https://open.dingtalk.com/
	TemplateEditUrl *string `json:"templateEditUrl,omitempty" xml:"templateEditUrl,omitempty"`
}

func (s PremiumSaveIntegratedProcessRequestTemplateConfig) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedProcessRequestTemplateConfig) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedProcessRequestTemplateConfig) SetCreateInstanceMobileUrl(v string) *PremiumSaveIntegratedProcessRequestTemplateConfig {
	s.CreateInstanceMobileUrl = &v
	return s
}

func (s *PremiumSaveIntegratedProcessRequestTemplateConfig) SetCreateInstancePcUrl(v string) *PremiumSaveIntegratedProcessRequestTemplateConfig {
	s.CreateInstancePcUrl = &v
	return s
}

func (s *PremiumSaveIntegratedProcessRequestTemplateConfig) SetDisableSendCard(v bool) *PremiumSaveIntegratedProcessRequestTemplateConfig {
	s.DisableSendCard = &v
	return s
}

func (s *PremiumSaveIntegratedProcessRequestTemplateConfig) SetHidden(v bool) *PremiumSaveIntegratedProcessRequestTemplateConfig {
	s.Hidden = &v
	return s
}

func (s *PremiumSaveIntegratedProcessRequestTemplateConfig) SetTemplateEditUrl(v string) *PremiumSaveIntegratedProcessRequestTemplateConfig {
	s.TemplateEditUrl = &v
	return s
}

type PremiumSaveIntegratedProcessResponseBody struct {
	Result *PremiumSaveIntegratedProcessResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s PremiumSaveIntegratedProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedProcessResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedProcessResponseBody) SetResult(v *PremiumSaveIntegratedProcessResponseBodyResult) *PremiumSaveIntegratedProcessResponseBody {
	s.Result = v
	return s
}

type PremiumSaveIntegratedProcessResponseBodyResult struct {
	// This parameter is required.
	//
	// example:
	//
	// PROC-abcdef-example
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
}

func (s PremiumSaveIntegratedProcessResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedProcessResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedProcessResponseBodyResult) SetProcessCode(v string) *PremiumSaveIntegratedProcessResponseBodyResult {
	s.ProcessCode = &v
	return s
}

type PremiumSaveIntegratedProcessResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumSaveIntegratedProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumSaveIntegratedProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedProcessResponse) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedProcessResponse) SetHeaders(v map[string]*string) *PremiumSaveIntegratedProcessResponse {
	s.Headers = v
	return s
}

func (s *PremiumSaveIntegratedProcessResponse) SetStatusCode(v int32) *PremiumSaveIntegratedProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumSaveIntegratedProcessResponse) SetBody(v *PremiumSaveIntegratedProcessResponseBody) *PremiumSaveIntegratedProcessResponse {
	s.Body = v
	return s
}

type PremiumSaveIntegratedProcessInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumSaveIntegratedProcessInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedProcessInstanceHeaders) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedProcessInstanceHeaders) SetCommonHeaders(v map[string]*string) *PremiumSaveIntegratedProcessInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumSaveIntegratedProcessInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumSaveIntegratedProcessInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumSaveIntegratedProcessInstanceRequest struct {
	// example:
	//
	// "{\"mykey\": \"myData\"}"
	BizData                *string                                                              `json:"bizData,omitempty" xml:"bizData,omitempty"`
	FeatureConfig          *PremiumSaveIntegratedProcessInstanceRequestFeatureConfig            `json:"featureConfig,omitempty" xml:"featureConfig,omitempty" type:"Struct"`
	FormComponentValueList []*PremiumSaveIntegratedProcessInstanceRequestFormComponentValueList `json:"formComponentValueList,omitempty" xml:"formComponentValueList,omitempty" type:"Repeated"`
	Notifiers              []*PremiumSaveIntegratedProcessInstanceRequestNotifiers              `json:"notifiers,omitempty" xml:"notifiers,omitempty" type:"Repeated"`
	// This parameter is required.
	OriginatorUserId *string `json:"originatorUserId,omitempty" xml:"originatorUserId,omitempty"`
	// This parameter is required.
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://www.dingtalk.com/
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s PremiumSaveIntegratedProcessInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedProcessInstanceRequest) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedProcessInstanceRequest) SetBizData(v string) *PremiumSaveIntegratedProcessInstanceRequest {
	s.BizData = &v
	return s
}

func (s *PremiumSaveIntegratedProcessInstanceRequest) SetFeatureConfig(v *PremiumSaveIntegratedProcessInstanceRequestFeatureConfig) *PremiumSaveIntegratedProcessInstanceRequest {
	s.FeatureConfig = v
	return s
}

func (s *PremiumSaveIntegratedProcessInstanceRequest) SetFormComponentValueList(v []*PremiumSaveIntegratedProcessInstanceRequestFormComponentValueList) *PremiumSaveIntegratedProcessInstanceRequest {
	s.FormComponentValueList = v
	return s
}

func (s *PremiumSaveIntegratedProcessInstanceRequest) SetNotifiers(v []*PremiumSaveIntegratedProcessInstanceRequestNotifiers) *PremiumSaveIntegratedProcessInstanceRequest {
	s.Notifiers = v
	return s
}

func (s *PremiumSaveIntegratedProcessInstanceRequest) SetOriginatorUserId(v string) *PremiumSaveIntegratedProcessInstanceRequest {
	s.OriginatorUserId = &v
	return s
}

func (s *PremiumSaveIntegratedProcessInstanceRequest) SetProcessCode(v string) *PremiumSaveIntegratedProcessInstanceRequest {
	s.ProcessCode = &v
	return s
}

func (s *PremiumSaveIntegratedProcessInstanceRequest) SetTitle(v string) *PremiumSaveIntegratedProcessInstanceRequest {
	s.Title = &v
	return s
}

func (s *PremiumSaveIntegratedProcessInstanceRequest) SetUrl(v string) *PremiumSaveIntegratedProcessInstanceRequest {
	s.Url = &v
	return s
}

type PremiumSaveIntegratedProcessInstanceRequestFeatureConfig struct {
	Features []*PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeatures `json:"features,omitempty" xml:"features,omitempty" type:"Repeated"`
}

func (s PremiumSaveIntegratedProcessInstanceRequestFeatureConfig) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedProcessInstanceRequestFeatureConfig) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedProcessInstanceRequestFeatureConfig) SetFeatures(v []*PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeatures) *PremiumSaveIntegratedProcessInstanceRequestFeatureConfig {
	s.Features = v
	return s
}

type PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeatures struct {
	Callback *PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeaturesCallback `json:"callback,omitempty" xml:"callback,omitempty" type:"Struct"`
	// if can be null:
	// true
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// example:
	//
	// www.dingtalk.com
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	// example:
	//
	// abc
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// www.dingtalk.com
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	// example:
	//
	// ORIGIN
	RunType *string `json:"runType,omitempty" xml:"runType,omitempty"`
}

func (s PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeatures) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeatures) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeatures) SetCallback(v *PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeaturesCallback) *PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeatures {
	s.Callback = v
	return s
}

func (s *PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeatures) SetConfig(v string) *PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeatures {
	s.Config = &v
	return s
}

func (s *PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeatures) SetMobileUrl(v string) *PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeatures {
	s.MobileUrl = &v
	return s
}

func (s *PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeatures) SetName(v string) *PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeatures {
	s.Name = &v
	return s
}

func (s *PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeatures) SetPcUrl(v string) *PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeatures {
	s.PcUrl = &v
	return s
}

func (s *PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeatures) SetRunType(v string) *PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeatures {
	s.RunType = &v
	return s
}

type PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeaturesCallback struct {
	// example:
	//
	// abc
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// example:
	//
	// abc
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeaturesCallback) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeaturesCallback) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeaturesCallback) SetApiKey(v string) *PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeaturesCallback {
	s.ApiKey = &v
	return s
}

func (s *PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeaturesCallback) SetAppUuid(v string) *PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeaturesCallback {
	s.AppUuid = &v
	return s
}

func (s *PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeaturesCallback) SetVersion(v string) *PremiumSaveIntegratedProcessInstanceRequestFeatureConfigFeaturesCallback {
	s.Version = &v
	return s
}

type PremiumSaveIntegratedProcessInstanceRequestFormComponentValueList struct {
	BizAlias      *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	ExtValue      *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	Id            *string `json:"id,omitempty" xml:"id,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	Value         *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s PremiumSaveIntegratedProcessInstanceRequestFormComponentValueList) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedProcessInstanceRequestFormComponentValueList) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedProcessInstanceRequestFormComponentValueList) SetBizAlias(v string) *PremiumSaveIntegratedProcessInstanceRequestFormComponentValueList {
	s.BizAlias = &v
	return s
}

func (s *PremiumSaveIntegratedProcessInstanceRequestFormComponentValueList) SetComponentType(v string) *PremiumSaveIntegratedProcessInstanceRequestFormComponentValueList {
	s.ComponentType = &v
	return s
}

func (s *PremiumSaveIntegratedProcessInstanceRequestFormComponentValueList) SetExtValue(v string) *PremiumSaveIntegratedProcessInstanceRequestFormComponentValueList {
	s.ExtValue = &v
	return s
}

func (s *PremiumSaveIntegratedProcessInstanceRequestFormComponentValueList) SetId(v string) *PremiumSaveIntegratedProcessInstanceRequestFormComponentValueList {
	s.Id = &v
	return s
}

func (s *PremiumSaveIntegratedProcessInstanceRequestFormComponentValueList) SetName(v string) *PremiumSaveIntegratedProcessInstanceRequestFormComponentValueList {
	s.Name = &v
	return s
}

func (s *PremiumSaveIntegratedProcessInstanceRequestFormComponentValueList) SetValue(v string) *PremiumSaveIntegratedProcessInstanceRequestFormComponentValueList {
	s.Value = &v
	return s
}

type PremiumSaveIntegratedProcessInstanceRequestNotifiers struct {
	// example:
	//
	// start
	Position *string `json:"position,omitempty" xml:"position,omitempty"`
	// example:
	//
	// manager001
	Userid *string `json:"userid,omitempty" xml:"userid,omitempty"`
}

func (s PremiumSaveIntegratedProcessInstanceRequestNotifiers) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedProcessInstanceRequestNotifiers) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedProcessInstanceRequestNotifiers) SetPosition(v string) *PremiumSaveIntegratedProcessInstanceRequestNotifiers {
	s.Position = &v
	return s
}

func (s *PremiumSaveIntegratedProcessInstanceRequestNotifiers) SetUserid(v string) *PremiumSaveIntegratedProcessInstanceRequestNotifiers {
	s.Userid = &v
	return s
}

type PremiumSaveIntegratedProcessInstanceResponseBody struct {
	Result *PremiumSaveIntegratedProcessInstanceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s PremiumSaveIntegratedProcessInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedProcessInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedProcessInstanceResponseBody) SetResult(v *PremiumSaveIntegratedProcessInstanceResponseBodyResult) *PremiumSaveIntegratedProcessInstanceResponseBody {
	s.Result = v
	return s
}

type PremiumSaveIntegratedProcessInstanceResponseBodyResult struct {
	// example:
	//
	// proc-abc
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
}

func (s PremiumSaveIntegratedProcessInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedProcessInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedProcessInstanceResponseBodyResult) SetProcessInstanceId(v string) *PremiumSaveIntegratedProcessInstanceResponseBodyResult {
	s.ProcessInstanceId = &v
	return s
}

type PremiumSaveIntegratedProcessInstanceResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumSaveIntegratedProcessInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumSaveIntegratedProcessInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedProcessInstanceResponse) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedProcessInstanceResponse) SetHeaders(v map[string]*string) *PremiumSaveIntegratedProcessInstanceResponse {
	s.Headers = v
	return s
}

func (s *PremiumSaveIntegratedProcessInstanceResponse) SetStatusCode(v int32) *PremiumSaveIntegratedProcessInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumSaveIntegratedProcessInstanceResponse) SetBody(v *PremiumSaveIntegratedProcessInstanceResponseBody) *PremiumSaveIntegratedProcessInstanceResponse {
	s.Body = v
	return s
}

type PremiumSaveIntegratedTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumSaveIntegratedTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedTaskHeaders) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedTaskHeaders) SetCommonHeaders(v map[string]*string) *PremiumSaveIntegratedTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumSaveIntegratedTaskHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumSaveIntegratedTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumSaveIntegratedTaskRequest struct {
	// example:
	//
	// act_xxxx
	ActivityId    *string                                        `json:"activityId,omitempty" xml:"activityId,omitempty"`
	FeatureConfig *PremiumSaveIntegratedTaskRequestFeatureConfig `json:"featureConfig,omitempty" xml:"featureConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// tPr_FB_mT_xxxxxxxxx2hQ05201655306463
	ProcessInstanceId *string                                     `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	TaskConfig        *PremiumSaveIntegratedTaskRequestTaskConfig `json:"taskConfig,omitempty" xml:"taskConfig,omitempty" type:"Struct"`
	// This parameter is required.
	Tasks []*PremiumSaveIntegratedTaskRequestTasks `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Repeated"`
}

func (s PremiumSaveIntegratedTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedTaskRequest) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedTaskRequest) SetActivityId(v string) *PremiumSaveIntegratedTaskRequest {
	s.ActivityId = &v
	return s
}

func (s *PremiumSaveIntegratedTaskRequest) SetFeatureConfig(v *PremiumSaveIntegratedTaskRequestFeatureConfig) *PremiumSaveIntegratedTaskRequest {
	s.FeatureConfig = v
	return s
}

func (s *PremiumSaveIntegratedTaskRequest) SetProcessInstanceId(v string) *PremiumSaveIntegratedTaskRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *PremiumSaveIntegratedTaskRequest) SetTaskConfig(v *PremiumSaveIntegratedTaskRequestTaskConfig) *PremiumSaveIntegratedTaskRequest {
	s.TaskConfig = v
	return s
}

func (s *PremiumSaveIntegratedTaskRequest) SetTasks(v []*PremiumSaveIntegratedTaskRequestTasks) *PremiumSaveIntegratedTaskRequest {
	s.Tasks = v
	return s
}

type PremiumSaveIntegratedTaskRequestFeatureConfig struct {
	Features []*PremiumSaveIntegratedTaskRequestFeatureConfigFeatures `json:"features,omitempty" xml:"features,omitempty" type:"Repeated"`
}

func (s PremiumSaveIntegratedTaskRequestFeatureConfig) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedTaskRequestFeatureConfig) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedTaskRequestFeatureConfig) SetFeatures(v []*PremiumSaveIntegratedTaskRequestFeatureConfigFeatures) *PremiumSaveIntegratedTaskRequestFeatureConfig {
	s.Features = v
	return s
}

type PremiumSaveIntegratedTaskRequestFeatureConfigFeatures struct {
	Callback *PremiumSaveIntegratedTaskRequestFeatureConfigFeaturesCallback `json:"callback,omitempty" xml:"callback,omitempty" type:"Struct"`
	// if can be null:
	// true
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// example:
	//
	// www.dingtalk.com
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	// example:
	//
	// abc
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// www.dingtalk.com
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	// example:
	//
	// ORIGIN
	RunType *string `json:"runType,omitempty" xml:"runType,omitempty"`
}

func (s PremiumSaveIntegratedTaskRequestFeatureConfigFeatures) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedTaskRequestFeatureConfigFeatures) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedTaskRequestFeatureConfigFeatures) SetCallback(v *PremiumSaveIntegratedTaskRequestFeatureConfigFeaturesCallback) *PremiumSaveIntegratedTaskRequestFeatureConfigFeatures {
	s.Callback = v
	return s
}

func (s *PremiumSaveIntegratedTaskRequestFeatureConfigFeatures) SetConfig(v string) *PremiumSaveIntegratedTaskRequestFeatureConfigFeatures {
	s.Config = &v
	return s
}

func (s *PremiumSaveIntegratedTaskRequestFeatureConfigFeatures) SetMobileUrl(v string) *PremiumSaveIntegratedTaskRequestFeatureConfigFeatures {
	s.MobileUrl = &v
	return s
}

func (s *PremiumSaveIntegratedTaskRequestFeatureConfigFeatures) SetName(v string) *PremiumSaveIntegratedTaskRequestFeatureConfigFeatures {
	s.Name = &v
	return s
}

func (s *PremiumSaveIntegratedTaskRequestFeatureConfigFeatures) SetPcUrl(v string) *PremiumSaveIntegratedTaskRequestFeatureConfigFeatures {
	s.PcUrl = &v
	return s
}

func (s *PremiumSaveIntegratedTaskRequestFeatureConfigFeatures) SetRunType(v string) *PremiumSaveIntegratedTaskRequestFeatureConfigFeatures {
	s.RunType = &v
	return s
}

type PremiumSaveIntegratedTaskRequestFeatureConfigFeaturesCallback struct {
	// example:
	//
	// abc
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// example:
	//
	// abc
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s PremiumSaveIntegratedTaskRequestFeatureConfigFeaturesCallback) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedTaskRequestFeatureConfigFeaturesCallback) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedTaskRequestFeatureConfigFeaturesCallback) SetApiKey(v string) *PremiumSaveIntegratedTaskRequestFeatureConfigFeaturesCallback {
	s.ApiKey = &v
	return s
}

func (s *PremiumSaveIntegratedTaskRequestFeatureConfigFeaturesCallback) SetAppUuid(v string) *PremiumSaveIntegratedTaskRequestFeatureConfigFeaturesCallback {
	s.AppUuid = &v
	return s
}

func (s *PremiumSaveIntegratedTaskRequestFeatureConfigFeaturesCallback) SetVersion(v string) *PremiumSaveIntegratedTaskRequestFeatureConfigFeaturesCallback {
	s.Version = &v
	return s
}

type PremiumSaveIntegratedTaskRequestTaskConfig struct {
	DisableSendCard *bool `json:"disableSendCard,omitempty" xml:"disableSendCard,omitempty"`
}

func (s PremiumSaveIntegratedTaskRequestTaskConfig) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedTaskRequestTaskConfig) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedTaskRequestTaskConfig) SetDisableSendCard(v bool) *PremiumSaveIntegratedTaskRequestTaskConfig {
	s.DisableSendCard = &v
	return s
}

type PremiumSaveIntegratedTaskRequestTasks struct {
	// example:
	//
	// {\"id\":\"12345\"}
	CustomData *string `json:"customData,omitempty" xml:"customData,omitempty"`
	// example:
	//
	// 1758643200000
	DueTimestamp *int64 `json:"dueTimestamp,omitempty" xml:"dueTimestamp,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// example:
	//
	// manager001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PremiumSaveIntegratedTaskRequestTasks) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedTaskRequestTasks) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedTaskRequestTasks) SetCustomData(v string) *PremiumSaveIntegratedTaskRequestTasks {
	s.CustomData = &v
	return s
}

func (s *PremiumSaveIntegratedTaskRequestTasks) SetDueTimestamp(v int64) *PremiumSaveIntegratedTaskRequestTasks {
	s.DueTimestamp = &v
	return s
}

func (s *PremiumSaveIntegratedTaskRequestTasks) SetUrl(v string) *PremiumSaveIntegratedTaskRequestTasks {
	s.Url = &v
	return s
}

func (s *PremiumSaveIntegratedTaskRequestTasks) SetUserId(v string) *PremiumSaveIntegratedTaskRequestTasks {
	s.UserId = &v
	return s
}

type PremiumSaveIntegratedTaskResponseBody struct {
	Result  []*PremiumSaveIntegratedTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	Success *bool                                          `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PremiumSaveIntegratedTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedTaskResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedTaskResponseBody) SetResult(v []*PremiumSaveIntegratedTaskResponseBodyResult) *PremiumSaveIntegratedTaskResponseBody {
	s.Result = v
	return s
}

func (s *PremiumSaveIntegratedTaskResponseBody) SetSuccess(v bool) *PremiumSaveIntegratedTaskResponseBody {
	s.Success = &v
	return s
}

type PremiumSaveIntegratedTaskResponseBodyResult struct {
	TaskId *int64  `json:"taskId,omitempty" xml:"taskId,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s PremiumSaveIntegratedTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedTaskResponseBodyResult) SetTaskId(v int64) *PremiumSaveIntegratedTaskResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *PremiumSaveIntegratedTaskResponseBodyResult) SetUserId(v string) *PremiumSaveIntegratedTaskResponseBodyResult {
	s.UserId = &v
	return s
}

type PremiumSaveIntegratedTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumSaveIntegratedTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumSaveIntegratedTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumSaveIntegratedTaskResponse) GoString() string {
	return s.String()
}

func (s *PremiumSaveIntegratedTaskResponse) SetHeaders(v map[string]*string) *PremiumSaveIntegratedTaskResponse {
	s.Headers = v
	return s
}

func (s *PremiumSaveIntegratedTaskResponse) SetStatusCode(v int32) *PremiumSaveIntegratedTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumSaveIntegratedTaskResponse) SetBody(v *PremiumSaveIntegratedTaskResponseBody) *PremiumSaveIntegratedTaskResponse {
	s.Body = v
	return s
}

type PremiumUpdateFormInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumUpdateFormInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumUpdateFormInstanceHeaders) GoString() string {
	return s.String()
}

func (s *PremiumUpdateFormInstanceHeaders) SetCommonHeaders(v map[string]*string) *PremiumUpdateFormInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumUpdateFormInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumUpdateFormInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumUpdateFormInstanceRequest struct {
	// This parameter is required.
	FormComponentValueList []*PremiumUpdateFormInstanceRequestFormComponentValueList `json:"formComponentValueList,omitempty" xml:"formComponentValueList,omitempty" type:"Repeated"`
	FormInstanceIds        []*string                                                 `json:"formInstanceIds,omitempty" xml:"formInstanceIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// manager432
	OriginatorUserId *string `json:"originatorUserId,omitempty" xml:"originatorUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-EF6YJL35P2-SCKICSB7P750S0YISYKV3-xxxx-1
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
}

func (s PremiumUpdateFormInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumUpdateFormInstanceRequest) GoString() string {
	return s.String()
}

func (s *PremiumUpdateFormInstanceRequest) SetFormComponentValueList(v []*PremiumUpdateFormInstanceRequestFormComponentValueList) *PremiumUpdateFormInstanceRequest {
	s.FormComponentValueList = v
	return s
}

func (s *PremiumUpdateFormInstanceRequest) SetFormInstanceIds(v []*string) *PremiumUpdateFormInstanceRequest {
	s.FormInstanceIds = v
	return s
}

func (s *PremiumUpdateFormInstanceRequest) SetOriginatorUserId(v string) *PremiumUpdateFormInstanceRequest {
	s.OriginatorUserId = &v
	return s
}

func (s *PremiumUpdateFormInstanceRequest) SetProcessCode(v string) *PremiumUpdateFormInstanceRequest {
	s.ProcessCode = &v
	return s
}

type PremiumUpdateFormInstanceRequestFormComponentValueList struct {
	// example:
	//
	// Phone
	BizAlias      *string                                                          `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	ComponentType *string                                                          `json:"componentType,omitempty" xml:"componentType,omitempty"`
	Details       []*PremiumUpdateFormInstanceRequestFormComponentValueListDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
	// example:
	//
	// 总个数:1
	ExtValue *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	// example:
	//
	// PhoneField_IZI2LP8QF6O0
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PhoneField
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123xxxxxxxx
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s PremiumUpdateFormInstanceRequestFormComponentValueList) String() string {
	return tea.Prettify(s)
}

func (s PremiumUpdateFormInstanceRequestFormComponentValueList) GoString() string {
	return s.String()
}

func (s *PremiumUpdateFormInstanceRequestFormComponentValueList) SetBizAlias(v string) *PremiumUpdateFormInstanceRequestFormComponentValueList {
	s.BizAlias = &v
	return s
}

func (s *PremiumUpdateFormInstanceRequestFormComponentValueList) SetComponentType(v string) *PremiumUpdateFormInstanceRequestFormComponentValueList {
	s.ComponentType = &v
	return s
}

func (s *PremiumUpdateFormInstanceRequestFormComponentValueList) SetDetails(v []*PremiumUpdateFormInstanceRequestFormComponentValueListDetails) *PremiumUpdateFormInstanceRequestFormComponentValueList {
	s.Details = v
	return s
}

func (s *PremiumUpdateFormInstanceRequestFormComponentValueList) SetExtValue(v string) *PremiumUpdateFormInstanceRequestFormComponentValueList {
	s.ExtValue = &v
	return s
}

func (s *PremiumUpdateFormInstanceRequestFormComponentValueList) SetId(v string) *PremiumUpdateFormInstanceRequestFormComponentValueList {
	s.Id = &v
	return s
}

func (s *PremiumUpdateFormInstanceRequestFormComponentValueList) SetName(v string) *PremiumUpdateFormInstanceRequestFormComponentValueList {
	s.Name = &v
	return s
}

func (s *PremiumUpdateFormInstanceRequestFormComponentValueList) SetValue(v string) *PremiumUpdateFormInstanceRequestFormComponentValueList {
	s.Value = &v
	return s
}

type PremiumUpdateFormInstanceRequestFormComponentValueListDetails struct {
	// example:
	//
	// Phone
	BizAlias *string                                                                 `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	Details  []*PremiumUpdateFormInstanceRequestFormComponentValueListDetailsDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
	// example:
	//
	// 总个数:1
	ExtValue *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	// example:
	//
	// PhoneField_IZI2LP8QF6O0
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// PhoneField
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 123xxxxxxxx
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s PremiumUpdateFormInstanceRequestFormComponentValueListDetails) String() string {
	return tea.Prettify(s)
}

func (s PremiumUpdateFormInstanceRequestFormComponentValueListDetails) GoString() string {
	return s.String()
}

func (s *PremiumUpdateFormInstanceRequestFormComponentValueListDetails) SetBizAlias(v string) *PremiumUpdateFormInstanceRequestFormComponentValueListDetails {
	s.BizAlias = &v
	return s
}

func (s *PremiumUpdateFormInstanceRequestFormComponentValueListDetails) SetDetails(v []*PremiumUpdateFormInstanceRequestFormComponentValueListDetailsDetails) *PremiumUpdateFormInstanceRequestFormComponentValueListDetails {
	s.Details = v
	return s
}

func (s *PremiumUpdateFormInstanceRequestFormComponentValueListDetails) SetExtValue(v string) *PremiumUpdateFormInstanceRequestFormComponentValueListDetails {
	s.ExtValue = &v
	return s
}

func (s *PremiumUpdateFormInstanceRequestFormComponentValueListDetails) SetId(v string) *PremiumUpdateFormInstanceRequestFormComponentValueListDetails {
	s.Id = &v
	return s
}

func (s *PremiumUpdateFormInstanceRequestFormComponentValueListDetails) SetName(v string) *PremiumUpdateFormInstanceRequestFormComponentValueListDetails {
	s.Name = &v
	return s
}

func (s *PremiumUpdateFormInstanceRequestFormComponentValueListDetails) SetValue(v string) *PremiumUpdateFormInstanceRequestFormComponentValueListDetails {
	s.Value = &v
	return s
}

type PremiumUpdateFormInstanceRequestFormComponentValueListDetailsDetails struct {
	// example:
	//
	// Phone
	BizAlias      *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// example:
	//
	// 总个数:1
	ExtValue *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	// example:
	//
	// PhoneField_IZI2LP8QF6O0
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// PhoneField
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 123xxxxxxxx
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s PremiumUpdateFormInstanceRequestFormComponentValueListDetailsDetails) String() string {
	return tea.Prettify(s)
}

func (s PremiumUpdateFormInstanceRequestFormComponentValueListDetailsDetails) GoString() string {
	return s.String()
}

func (s *PremiumUpdateFormInstanceRequestFormComponentValueListDetailsDetails) SetBizAlias(v string) *PremiumUpdateFormInstanceRequestFormComponentValueListDetailsDetails {
	s.BizAlias = &v
	return s
}

func (s *PremiumUpdateFormInstanceRequestFormComponentValueListDetailsDetails) SetComponentType(v string) *PremiumUpdateFormInstanceRequestFormComponentValueListDetailsDetails {
	s.ComponentType = &v
	return s
}

func (s *PremiumUpdateFormInstanceRequestFormComponentValueListDetailsDetails) SetExtValue(v string) *PremiumUpdateFormInstanceRequestFormComponentValueListDetailsDetails {
	s.ExtValue = &v
	return s
}

func (s *PremiumUpdateFormInstanceRequestFormComponentValueListDetailsDetails) SetId(v string) *PremiumUpdateFormInstanceRequestFormComponentValueListDetailsDetails {
	s.Id = &v
	return s
}

func (s *PremiumUpdateFormInstanceRequestFormComponentValueListDetailsDetails) SetName(v string) *PremiumUpdateFormInstanceRequestFormComponentValueListDetailsDetails {
	s.Name = &v
	return s
}

func (s *PremiumUpdateFormInstanceRequestFormComponentValueListDetailsDetails) SetValue(v string) *PremiumUpdateFormInstanceRequestFormComponentValueListDetailsDetails {
	s.Value = &v
	return s
}

type PremiumUpdateFormInstanceResponseBody struct {
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s PremiumUpdateFormInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumUpdateFormInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumUpdateFormInstanceResponseBody) SetInstanceId(v string) *PremiumUpdateFormInstanceResponseBody {
	s.InstanceId = &v
	return s
}

type PremiumUpdateFormInstanceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumUpdateFormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumUpdateFormInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumUpdateFormInstanceResponse) GoString() string {
	return s.String()
}

func (s *PremiumUpdateFormInstanceResponse) SetHeaders(v map[string]*string) *PremiumUpdateFormInstanceResponse {
	s.Headers = v
	return s
}

func (s *PremiumUpdateFormInstanceResponse) SetStatusCode(v int32) *PremiumUpdateFormInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumUpdateFormInstanceResponse) SetBody(v *PremiumUpdateFormInstanceResponseBody) *PremiumUpdateFormInstanceResponse {
	s.Body = v
	return s
}

type PremiumUpdateProcessInstanceVariablesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s PremiumUpdateProcessInstanceVariablesHeaders) String() string {
	return tea.Prettify(s)
}

func (s PremiumUpdateProcessInstanceVariablesHeaders) GoString() string {
	return s.String()
}

func (s *PremiumUpdateProcessInstanceVariablesHeaders) SetCommonHeaders(v map[string]*string) *PremiumUpdateProcessInstanceVariablesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PremiumUpdateProcessInstanceVariablesHeaders) SetXAcsDingtalkAccessToken(v string) *PremiumUpdateProcessInstanceVariablesHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type PremiumUpdateProcessInstanceVariablesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// manager432
	OpUserId *string `json:"opUserId,omitempty" xml:"opUserId,omitempty"`
	// example:
	//
	// PROC-EF6YJL35P2-SCKICSB7P750S0YISYKV3-xxxx-1
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// processInstanceId-1
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// example:
	//
	// processInstanceId-1
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// This parameter is required.
	Variables []*PremiumUpdateProcessInstanceVariablesRequestVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s PremiumUpdateProcessInstanceVariablesRequest) String() string {
	return tea.Prettify(s)
}

func (s PremiumUpdateProcessInstanceVariablesRequest) GoString() string {
	return s.String()
}

func (s *PremiumUpdateProcessInstanceVariablesRequest) SetOpUserId(v string) *PremiumUpdateProcessInstanceVariablesRequest {
	s.OpUserId = &v
	return s
}

func (s *PremiumUpdateProcessInstanceVariablesRequest) SetProcessCode(v string) *PremiumUpdateProcessInstanceVariablesRequest {
	s.ProcessCode = &v
	return s
}

func (s *PremiumUpdateProcessInstanceVariablesRequest) SetProcessInstanceId(v string) *PremiumUpdateProcessInstanceVariablesRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *PremiumUpdateProcessInstanceVariablesRequest) SetRemark(v string) *PremiumUpdateProcessInstanceVariablesRequest {
	s.Remark = &v
	return s
}

func (s *PremiumUpdateProcessInstanceVariablesRequest) SetVariables(v []*PremiumUpdateProcessInstanceVariablesRequestVariables) *PremiumUpdateProcessInstanceVariablesRequest {
	s.Variables = v
	return s
}

type PremiumUpdateProcessInstanceVariablesRequestVariables struct {
	// example:
	//
	// Phone
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// example:
	//
	// 总个数:1
	ExtValue *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PhoneField_IZI2LP8QF6O0
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123xxxxxxxx
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s PremiumUpdateProcessInstanceVariablesRequestVariables) String() string {
	return tea.Prettify(s)
}

func (s PremiumUpdateProcessInstanceVariablesRequestVariables) GoString() string {
	return s.String()
}

func (s *PremiumUpdateProcessInstanceVariablesRequestVariables) SetBizAlias(v string) *PremiumUpdateProcessInstanceVariablesRequestVariables {
	s.BizAlias = &v
	return s
}

func (s *PremiumUpdateProcessInstanceVariablesRequestVariables) SetExtValue(v string) *PremiumUpdateProcessInstanceVariablesRequestVariables {
	s.ExtValue = &v
	return s
}

func (s *PremiumUpdateProcessInstanceVariablesRequestVariables) SetId(v string) *PremiumUpdateProcessInstanceVariablesRequestVariables {
	s.Id = &v
	return s
}

func (s *PremiumUpdateProcessInstanceVariablesRequestVariables) SetValue(v string) *PremiumUpdateProcessInstanceVariablesRequestVariables {
	s.Value = &v
	return s
}

type PremiumUpdateProcessInstanceVariablesResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s PremiumUpdateProcessInstanceVariablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PremiumUpdateProcessInstanceVariablesResponseBody) GoString() string {
	return s.String()
}

func (s *PremiumUpdateProcessInstanceVariablesResponseBody) SetResult(v bool) *PremiumUpdateProcessInstanceVariablesResponseBody {
	s.Result = &v
	return s
}

type PremiumUpdateProcessInstanceVariablesResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PremiumUpdateProcessInstanceVariablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PremiumUpdateProcessInstanceVariablesResponse) String() string {
	return tea.Prettify(s)
}

func (s PremiumUpdateProcessInstanceVariablesResponse) GoString() string {
	return s.String()
}

func (s *PremiumUpdateProcessInstanceVariablesResponse) SetHeaders(v map[string]*string) *PremiumUpdateProcessInstanceVariablesResponse {
	s.Headers = v
	return s
}

func (s *PremiumUpdateProcessInstanceVariablesResponse) SetStatusCode(v int32) *PremiumUpdateProcessInstanceVariablesResponse {
	s.StatusCode = &v
	return s
}

func (s *PremiumUpdateProcessInstanceVariablesResponse) SetBody(v *PremiumUpdateProcessInstanceVariablesResponseBody) *PremiumUpdateProcessInstanceVariablesResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// 1
	DeptId *int32 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// This parameter is required.
	FormComponentValues []*ProcessForecastRequestFormComponentValues `json:"formComponentValues,omitempty" xml:"formComponentValues,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-EF6YJL35P2-SCKICSB7P750S0YISYKV3-xxxx-1
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// manager432
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ProcessForecastRequest) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastRequest) GoString() string {
	return s.String()
}

func (s *ProcessForecastRequest) SetDeptId(v int32) *ProcessForecastRequest {
	s.DeptId = &v
	return s
}

func (s *ProcessForecastRequest) SetFormComponentValues(v []*ProcessForecastRequestFormComponentValues) *ProcessForecastRequest {
	s.FormComponentValues = v
	return s
}

func (s *ProcessForecastRequest) SetProcessCode(v string) *ProcessForecastRequest {
	s.ProcessCode = &v
	return s
}

func (s *ProcessForecastRequest) SetUserId(v string) *ProcessForecastRequest {
	s.UserId = &v
	return s
}

type ProcessForecastRequestFormComponentValues struct {
	// example:
	//
	// Phone
	BizAlias      *string                                             `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	ComponentType *string                                             `json:"componentType,omitempty" xml:"componentType,omitempty"`
	Details       []*ProcessForecastRequestFormComponentValuesDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
	// example:
	//
	// 总个数:1
	ExtValue *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	// example:
	//
	// PhoneField_IZI2LP8QF6O0
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PhoneField
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123xxxxxxxx
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ProcessForecastRequestFormComponentValues) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastRequestFormComponentValues) GoString() string {
	return s.String()
}

func (s *ProcessForecastRequestFormComponentValues) SetBizAlias(v string) *ProcessForecastRequestFormComponentValues {
	s.BizAlias = &v
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

func (s *ProcessForecastRequestFormComponentValues) SetExtValue(v string) *ProcessForecastRequestFormComponentValues {
	s.ExtValue = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValues) SetId(v string) *ProcessForecastRequestFormComponentValues {
	s.Id = &v
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

type ProcessForecastRequestFormComponentValuesDetails struct {
	// example:
	//
	// Phone
	BizAlias *string                                                    `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	Details  []*ProcessForecastRequestFormComponentValuesDetailsDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
	// example:
	//
	// 总个数:1
	ExtValue *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	// example:
	//
	// PhoneField_IZI2LP8QF6O0
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// PhoneField
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 123xxxxxxxx
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ProcessForecastRequestFormComponentValuesDetails) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastRequestFormComponentValuesDetails) GoString() string {
	return s.String()
}

func (s *ProcessForecastRequestFormComponentValuesDetails) SetBizAlias(v string) *ProcessForecastRequestFormComponentValuesDetails {
	s.BizAlias = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValuesDetails) SetDetails(v []*ProcessForecastRequestFormComponentValuesDetailsDetails) *ProcessForecastRequestFormComponentValuesDetails {
	s.Details = v
	return s
}

func (s *ProcessForecastRequestFormComponentValuesDetails) SetExtValue(v string) *ProcessForecastRequestFormComponentValuesDetails {
	s.ExtValue = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValuesDetails) SetId(v string) *ProcessForecastRequestFormComponentValuesDetails {
	s.Id = &v
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

type ProcessForecastRequestFormComponentValuesDetailsDetails struct {
	// example:
	//
	// Phone
	BizAlias      *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// example:
	//
	// 总个数:1
	ExtValue *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	// example:
	//
	// PhoneField_IZI2LP8QF6O0
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// PhoneField
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 123xxxxxxxx
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ProcessForecastRequestFormComponentValuesDetailsDetails) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastRequestFormComponentValuesDetailsDetails) GoString() string {
	return s.String()
}

func (s *ProcessForecastRequestFormComponentValuesDetailsDetails) SetBizAlias(v string) *ProcessForecastRequestFormComponentValuesDetailsDetails {
	s.BizAlias = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValuesDetailsDetails) SetComponentType(v string) *ProcessForecastRequestFormComponentValuesDetailsDetails {
	s.ComponentType = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValuesDetailsDetails) SetExtValue(v string) *ProcessForecastRequestFormComponentValuesDetailsDetails {
	s.ExtValue = &v
	return s
}

func (s *ProcessForecastRequestFormComponentValuesDetailsDetails) SetId(v string) *ProcessForecastRequestFormComponentValuesDetailsDetails {
	s.Id = &v
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

type ProcessForecastResponseBody struct {
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
	// This parameter is required.
	//
	// example:
	//
	// true
	IsForecastSuccess *bool `json:"isForecastSuccess,omitempty" xml:"isForecastSuccess,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	IsStaticWorkflow *bool `json:"isStaticWorkflow,omitempty" xml:"isStaticWorkflow,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-2B60E506-D6CB-43F3-B661-359B27F90947
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 63657309999
	ProcessId *int64 `json:"processId,omitempty" xml:"processId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2665246100805992
	UserId                *string                                                   `json:"userId,omitempty" xml:"userId,omitempty"`
	WorkflowActivityRules []*ProcessForecastResponseBodyResultWorkflowActivityRules `json:"workflowActivityRules,omitempty" xml:"workflowActivityRules,omitempty" type:"Repeated"`
	// This parameter is required.
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

func (s *ProcessForecastResponseBodyResult) SetIsStaticWorkflow(v bool) *ProcessForecastResponseBodyResult {
	s.IsStaticWorkflow = &v
	return s
}

func (s *ProcessForecastResponseBodyResult) SetProcessCode(v string) *ProcessForecastResponseBodyResult {
	s.ProcessCode = &v
	return s
}

func (s *ProcessForecastResponseBodyResult) SetProcessId(v int64) *ProcessForecastResponseBodyResult {
	s.ProcessId = &v
	return s
}

func (s *ProcessForecastResponseBodyResult) SetUserId(v string) *ProcessForecastResponseBodyResult {
	s.UserId = &v
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
	ActivityActioners []*ProcessForecastResponseBodyResultWorkflowActivityRulesActivityActioners `json:"activityActioners,omitempty" xml:"activityActioners,omitempty" type:"Repeated"`
	// example:
	//
	// 1918_5cd3
	ActivityId *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
	// example:
	//
	// 审批人
	ActivityName *string `json:"activityName,omitempty" xml:"activityName,omitempty"`
	// example:
	//
	// 包括 target_select、target_approval 等
	ActivityType *string `json:"activityType,omitempty" xml:"activityType,omitempty"`
	// example:
	//
	// true
	IsTargetSelect *bool `json:"isTargetSelect,omitempty" xml:"isTargetSelect,omitempty"`
	// example:
	//
	// 1918_5cd3
	PrevActivityId *string                                                              `json:"prevActivityId,omitempty" xml:"prevActivityId,omitempty"`
	WorkflowActor  *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor `json:"workflowActor,omitempty" xml:"workflowActor,omitempty" type:"Struct"`
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRules) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRules) GoString() string {
	return s.String()
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRules) SetActivityActioners(v []*ProcessForecastResponseBodyResultWorkflowActivityRulesActivityActioners) *ProcessForecastResponseBodyResultWorkflowActivityRules {
	s.ActivityActioners = v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRules) SetActivityId(v string) *ProcessForecastResponseBodyResultWorkflowActivityRules {
	s.ActivityId = &v
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

func (s *ProcessForecastResponseBodyResultWorkflowActivityRules) SetPrevActivityId(v string) *ProcessForecastResponseBodyResultWorkflowActivityRules {
	s.PrevActivityId = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRules) SetWorkflowActor(v *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) *ProcessForecastResponseBodyResultWorkflowActivityRules {
	s.WorkflowActor = v
	return s
}

type ProcessForecastResponseBodyResultWorkflowActivityRulesActivityActioners struct {
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRulesActivityActioners) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRulesActivityActioners) GoString() string {
	return s.String()
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesActivityActioners) SetAvatar(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesActivityActioners {
	s.Avatar = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesActivityActioners) SetName(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesActivityActioners {
	s.Name = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesActivityActioners) SetUserId(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesActivityActioners {
	s.UserId = &v
	return s
}

type ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor struct {
	// example:
	//
	// ALL:并行，ONE_BY_ONE:串行
	ActorActivateType *string `json:"actorActivateType,omitempty" xml:"actorActivateType,omitempty"`
	// example:
	//
	// manual_e203_14a3_895a_45ad
	ActorKey            *string                                                                                 `json:"actorKey,omitempty" xml:"actorKey,omitempty"`
	ActorSelectionRange *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRange `json:"actorSelectionRange,omitempty" xml:"actorSelectionRange,omitempty" type:"Struct"`
	// example:
	//
	// allStaff：全公司，approvals：指定成员，labels：角色
	ActorSelectionType *string `json:"actorSelectionType,omitempty" xml:"actorSelectionType,omitempty"`
	// example:
	//
	// approver:审批人，notifier:抄送人，audit：办理人
	ActorType *string `json:"actorType,omitempty" xml:"actorType,omitempty"`
	// example:
	//
	// true
	AllowedMulti *bool `json:"allowedMulti,omitempty" xml:"allowedMulti,omitempty"`
	// example:
	//
	// ONE_BY_ONE：依次审批，AND：会签审批，OR：或签审批
	ApprovalMethod *string `json:"approvalMethod,omitempty" xml:"approvalMethod,omitempty"`
	// example:
	//
	// MANUAL:人工审批，AUTO_AGREE:自动通过，AUTO_REFUSE:自动拒绝
	ApprovalType *string `json:"approvalType,omitempty" xml:"approvalType,omitempty"`
	// example:
	//
	// true
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) GoString() string {
	return s.String()
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetActorActivateType(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.ActorActivateType = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetActorKey(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.ActorKey = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetActorSelectionRange(v *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRange) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.ActorSelectionRange = v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetActorSelectionType(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.ActorSelectionType = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetActorType(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.ActorType = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetAllowedMulti(v bool) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.AllowedMulti = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetApprovalMethod(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.ApprovalMethod = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetApprovalType(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.ApprovalType = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor) SetRequired(v bool) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActor {
	s.Required = &v
	return s
}

type ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRange struct {
	Approvals []*ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeApprovals `json:"approvals,omitempty" xml:"approvals,omitempty" type:"Repeated"`
	Labels    []*ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeLabels    `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
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
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	WorkNo   *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeApprovals) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeApprovals) GoString() string {
	return s.String()
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeApprovals) SetUserName(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeApprovals {
	s.UserName = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeApprovals) SetWorkNo(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeApprovals {
	s.WorkNo = &v
	return s
}

type ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeLabels struct {
	LabelNames *string `json:"labelNames,omitempty" xml:"labelNames,omitempty"`
	Labels     *string `json:"labels,omitempty" xml:"labels,omitempty"`
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeLabels) String() string {
	return tea.Prettify(s)
}

func (s ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeLabels) GoString() string {
	return s.String()
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeLabels) SetLabelNames(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeLabels {
	s.LabelNames = &v
	return s
}

func (s *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeLabels) SetLabels(v string) *ProcessForecastResponseBodyResultWorkflowActivityRulesWorkflowActorActorSelectionRangeLabels {
	s.Labels = &v
	return s
}

type ProcessForecastResponseBodyResultWorkflowForecastNodes struct {
	// This parameter is required.
	//
	// example:
	//
	// 1cc3_959a
	ActivityId *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// line-random-1cc3_959a-831a_607b
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProcessForecastResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ProcessForecastResponse) SetStatusCode(v int32) *ProcessForecastResponse {
	s.StatusCode = &v
	return s
}

func (s *ProcessForecastResponse) SetBody(v *ProcessForecastResponseBody) *ProcessForecastResponse {
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
	// example:
	//
	// SWAPP-dacdsa-example
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-daccea-example
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 100010
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryAllFormInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllFormInstancesRequest) GoString() string {
	return s.String()
}

func (s *QueryAllFormInstancesRequest) SetAppUuid(v string) *QueryAllFormInstancesRequest {
	s.AppUuid = &v
	return s
}

func (s *QueryAllFormInstancesRequest) SetFormCode(v string) *QueryAllFormInstancesRequest {
	s.FormCode = &v
	return s
}

func (s *QueryAllFormInstancesRequest) SetMaxResults(v int32) *QueryAllFormInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryAllFormInstancesRequest) SetNextToken(v string) *QueryAllFormInstancesRequest {
	s.NextToken = &v
	return s
}

type QueryAllFormInstancesResponseBody struct {
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	Values []*QueryAllFormInstancesResponseBodyResultValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s QueryAllFormInstancesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryAllFormInstancesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryAllFormInstancesResponseBodyResult) SetHasMore(v bool) *QueryAllFormInstancesResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResult) SetMaxResults(v int64) *QueryAllFormInstancesResponseBodyResult {
	s.MaxResults = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResult) SetNextToken(v string) *QueryAllFormInstancesResponseBodyResult {
	s.NextToken = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResult) SetValues(v []*QueryAllFormInstancesResponseBodyResultValues) *QueryAllFormInstancesResponseBodyResult {
	s.Values = v
	return s
}

type QueryAllFormInstancesResponseBodyResultValues struct {
	// example:
	//
	// SWAPP-abcd-example
	AppUuid    *string                `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1635151039000
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30314512
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-abcd-example
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// This parameter is required.
	FormInstDataList []*QueryAllFormInstancesResponseBodyResultValuesFormInstDataList `json:"formInstDataList,omitempty" xml:"formInstDataList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// abcd-eaf-acde12f
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 032142312
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// example:
	//
	// 1635151039000
	ModifyTimestamp *int64 `json:"modifyTimestamp,omitempty" xml:"modifyTimestamp,omitempty"`
	// example:
	//
	// abcd
	OutBizCode *string `json:"outBizCode,omitempty" xml:"outBizCode,omitempty"`
	// example:
	//
	// 323
	OutInstanceId *string `json:"outInstanceId,omitempty" xml:"outInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx提交的数据
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryAllFormInstancesResponseBodyResultValues) String() string {
	return tea.Prettify(s)
}

func (s QueryAllFormInstancesResponseBodyResultValues) GoString() string {
	return s.String()
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetAppUuid(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.AppUuid = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetAttributes(v map[string]interface{}) *QueryAllFormInstancesResponseBodyResultValues {
	s.Attributes = v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetCreateTimestamp(v int64) *QueryAllFormInstancesResponseBodyResultValues {
	s.CreateTimestamp = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetCreator(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.Creator = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetFormCode(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.FormCode = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetFormInstDataList(v []*QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) *QueryAllFormInstancesResponseBodyResultValues {
	s.FormInstDataList = v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetFormInstanceId(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.FormInstanceId = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetModifier(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.Modifier = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetModifyTimestamp(v int64) *QueryAllFormInstancesResponseBodyResultValues {
	s.ModifyTimestamp = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetOutBizCode(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.OutBizCode = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetOutInstanceId(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.OutInstanceId = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValues) SetTitle(v string) *QueryAllFormInstancesResponseBodyResultValues {
	s.Title = &v
	return s
}

type QueryAllFormInstancesResponseBodyResultValuesFormInstDataList struct {
	// example:
	//
	// staff_name
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 具体参见审批控件列表
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// example:
	//
	// {"key":"value}
	ExtendValue *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TextField-abcdefg
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 员工姓名
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 张三
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) GoString() string {
	return s.String()
}

func (s *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) SetBizAlias(v string) *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList {
	s.BizAlias = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) SetComponentType(v string) *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList {
	s.ComponentType = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) SetExtendValue(v string) *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList {
	s.ExtendValue = &v
	return s
}

func (s *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList) SetKey(v string) *QueryAllFormInstancesResponseBodyResultValuesFormInstDataList {
	s.Key = &v
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

type QueryAllFormInstancesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAllFormInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryAllFormInstancesResponse) SetStatusCode(v int32) *QueryAllFormInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAllFormInstancesResponse) SetBody(v *QueryAllFormInstancesResponseBody) *QueryAllFormInstancesResponse {
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
	// example:
	//
	// SWAPP-4C2F4B-example
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// example:
	//
	// 1633795200000
	EndTimeInMills *int64 `json:"endTimeInMills,omitempty" xml:"endTimeInMills,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-C53-example
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1631289600000
	StartTimeInMills *int64 `json:"startTimeInMills,omitempty" xml:"startTimeInMills,omitempty"`
}

func (s QueryAllProcessInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesRequest) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesRequest) SetAppUuid(v string) *QueryAllProcessInstancesRequest {
	s.AppUuid = &v
	return s
}

func (s *QueryAllProcessInstancesRequest) SetEndTimeInMills(v int64) *QueryAllProcessInstancesRequest {
	s.EndTimeInMills = &v
	return s
}

func (s *QueryAllProcessInstancesRequest) SetMaxResults(v int64) *QueryAllProcessInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryAllProcessInstancesRequest) SetNextToken(v string) *QueryAllProcessInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *QueryAllProcessInstancesRequest) SetProcessCode(v string) *QueryAllProcessInstancesRequest {
	s.ProcessCode = &v
	return s
}

func (s *QueryAllProcessInstancesRequest) SetStartTimeInMills(v int64) *QueryAllProcessInstancesRequest {
	s.StartTimeInMills = &v
	return s
}

type QueryAllProcessInstancesResponseBody struct {
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
	// This parameter is required.
	//
	// example:
	//
	// true
	HasMore *bool                                             `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List    []*QueryAllProcessInstancesResponseBodyResultList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryAllProcessInstancesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesResponseBodyResult) SetHasMore(v bool) *QueryAllProcessInstancesResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResult) SetList(v []*QueryAllProcessInstancesResponseBodyResultList) *QueryAllProcessInstancesResponseBodyResult {
	s.List = v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResult) SetMaxResults(v int64) *QueryAllProcessInstancesResponseBodyResult {
	s.MaxResults = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResult) SetNextToken(v string) *QueryAllProcessInstancesResponseBodyResult {
	s.NextToken = &v
	return s
}

type QueryAllProcessInstancesResponseBodyResultList struct {
	// This parameter is required.
	//
	// example:
	//
	// cdef-dae2fd2-example
	AttachedProcessInstanceIds *string `json:"attachedProcessInstanceIds,omitempty" xml:"attachedProcessInstanceIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 202110111558000355024
	BusinessId *string `json:"businessId,omitempty" xml:"businessId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1635165470201
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1633795200000
	FinishTime *int64 `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// This parameter is required.
	FormComponentValues []*QueryAllProcessInstancesResponseBodyResultListFormComponentValues `json:"formComponentValues,omitempty" xml:"formComponentValues,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// dcdse-dae2fd2-example
	MainProcessInstanceId *string                                                           `json:"mainProcessInstanceId,omitempty" xml:"mainProcessInstanceId,omitempty"`
	OperationRecords      []*QueryAllProcessInstancesResponseBodyResultListOperationRecords `json:"operationRecords,omitempty" xml:"operationRecords,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 默认-1，企业根部门
	OriginatorDeptId *string `json:"originatorDeptId,omitempty" xml:"originatorDeptId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// staff1234
	OriginatorUserid *string `json:"originatorUserid,omitempty" xml:"originatorUserid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abcdse-dse-example
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AGREE同意，REFUSE拒绝
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RUNNING审批中、TERMINATED撤销、COMPLETED审批完成、CANCELED取消
	Status *string                                                `json:"status,omitempty" xml:"status,omitempty"`
	Tasks  []*QueryAllProcessInstancesResponseBodyResultListTasks `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 员工A提交的小肖审批单
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryAllProcessInstancesResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetAttachedProcessInstanceIds(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.AttachedProcessInstanceIds = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetBusinessId(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.BusinessId = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetCreateTime(v int64) *QueryAllProcessInstancesResponseBodyResultList {
	s.CreateTime = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetFinishTime(v int64) *QueryAllProcessInstancesResponseBodyResultList {
	s.FinishTime = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetFormComponentValues(v []*QueryAllProcessInstancesResponseBodyResultListFormComponentValues) *QueryAllProcessInstancesResponseBodyResultList {
	s.FormComponentValues = v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetMainProcessInstanceId(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.MainProcessInstanceId = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetOperationRecords(v []*QueryAllProcessInstancesResponseBodyResultListOperationRecords) *QueryAllProcessInstancesResponseBodyResultList {
	s.OperationRecords = v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetOriginatorDeptId(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.OriginatorDeptId = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetOriginatorUserid(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.OriginatorUserid = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetProcessInstanceId(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.ProcessInstanceId = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetResult(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.Result = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetStatus(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.Status = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetTasks(v []*QueryAllProcessInstancesResponseBodyResultListTasks) *QueryAllProcessInstancesResponseBodyResultList {
	s.Tasks = v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultList) SetTitle(v string) *QueryAllProcessInstancesResponseBodyResultList {
	s.Title = &v
	return s
}

type QueryAllProcessInstancesResponseBodyResultListFormComponentValues struct {
	// example:
	//
	// {"staffId":"abcd"}
	ExtValue *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TextField-a32bcdef
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 张三
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s QueryAllProcessInstancesResponseBodyResultListFormComponentValues) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesResponseBodyResultListFormComponentValues) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesResponseBodyResultListFormComponentValues) SetExtValue(v string) *QueryAllProcessInstancesResponseBodyResultListFormComponentValues {
	s.ExtValue = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultListFormComponentValues) SetId(v string) *QueryAllProcessInstancesResponseBodyResultListFormComponentValues {
	s.Id = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultListFormComponentValues) SetName(v string) *QueryAllProcessInstancesResponseBodyResultListFormComponentValues {
	s.Name = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultListFormComponentValues) SetValue(v string) *QueryAllProcessInstancesResponseBodyResultListFormComponentValues {
	s.Value = &v
	return s
}

type QueryAllProcessInstancesResponseBodyResultListOperationRecords struct {
	Attachments []*QueryAllProcessInstancesResponseBodyResultListOperationRecordsAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	// example:
	//
	// EXECUTE_TASK_NORMAL（正常执行任务），EXECUTE_TASK_AGENT（代理人执行任务），APPEND_TASK_BEFORE（前加签任务），APPEND_TASK_AFTER（后加签任务），REDIRECT_TASK（转交任务），START_PROCESS_INSTANCE（发起流程实例），TERMINATE_PROCESS_INSTANCE（终止(撤销)流程实例），FINISH_PROCESS_INSTANCE（结束流程实例），ADD_REMARK（添加评论）
	OperationType *string `json:"operationType,omitempty" xml:"operationType,omitempty"`
	// example:
	//
	// 同意
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// AGREE（同意），REFUSE（拒绝），NONE（未知）
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// 1657522271000
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// example:
	//
	// manager1
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryAllProcessInstancesResponseBodyResultListOperationRecords) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesResponseBodyResultListOperationRecords) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesResponseBodyResultListOperationRecords) SetAttachments(v []*QueryAllProcessInstancesResponseBodyResultListOperationRecordsAttachments) *QueryAllProcessInstancesResponseBodyResultListOperationRecords {
	s.Attachments = v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultListOperationRecords) SetOperationType(v string) *QueryAllProcessInstancesResponseBodyResultListOperationRecords {
	s.OperationType = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultListOperationRecords) SetRemark(v string) *QueryAllProcessInstancesResponseBodyResultListOperationRecords {
	s.Remark = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultListOperationRecords) SetResult(v string) *QueryAllProcessInstancesResponseBodyResultListOperationRecords {
	s.Result = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultListOperationRecords) SetTimestamp(v int64) *QueryAllProcessInstancesResponseBodyResultListOperationRecords {
	s.Timestamp = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultListOperationRecords) SetUserId(v string) *QueryAllProcessInstancesResponseBodyResultListOperationRecords {
	s.UserId = &v
	return s
}

type QueryAllProcessInstancesResponseBodyResultListOperationRecordsAttachments struct {
	// example:
	//
	// 1234567
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// example:
	//
	// 附件
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// 123
	FileSize *string `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
}

func (s QueryAllProcessInstancesResponseBodyResultListOperationRecordsAttachments) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesResponseBodyResultListOperationRecordsAttachments) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesResponseBodyResultListOperationRecordsAttachments) SetFileId(v string) *QueryAllProcessInstancesResponseBodyResultListOperationRecordsAttachments {
	s.FileId = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultListOperationRecordsAttachments) SetFileName(v string) *QueryAllProcessInstancesResponseBodyResultListOperationRecordsAttachments {
	s.FileName = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultListOperationRecordsAttachments) SetFileSize(v string) *QueryAllProcessInstancesResponseBodyResultListOperationRecordsAttachments {
	s.FileSize = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultListOperationRecordsAttachments) SetFileType(v string) *QueryAllProcessInstancesResponseBodyResultListOperationRecordsAttachments {
	s.FileType = &v
	return s
}

type QueryAllProcessInstancesResponseBodyResultListTasks struct {
	// example:
	//
	// 1234_abcd
	ActivityId *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
	// example:
	//
	// 1657522271000
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// example:
	//
	// 1657522271000
	FinishTimestamp *int64 `json:"finishTimestamp,omitempty" xml:"finishTimestamp,omitempty"`
	// example:
	//
	// 分为AGREE（同意），REFUSE（拒绝），REDIRECTED（转交）
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// NEW（未启动），RUNNING（处理中），PAUSED（暂停），CANCELED（取消），COMPLETED（完成），TERMINATED（终止）
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// staff1234
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryAllProcessInstancesResponseBodyResultListTasks) String() string {
	return tea.Prettify(s)
}

func (s QueryAllProcessInstancesResponseBodyResultListTasks) GoString() string {
	return s.String()
}

func (s *QueryAllProcessInstancesResponseBodyResultListTasks) SetActivityId(v string) *QueryAllProcessInstancesResponseBodyResultListTasks {
	s.ActivityId = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultListTasks) SetCreateTimestamp(v int64) *QueryAllProcessInstancesResponseBodyResultListTasks {
	s.CreateTimestamp = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultListTasks) SetFinishTimestamp(v int64) *QueryAllProcessInstancesResponseBodyResultListTasks {
	s.FinishTimestamp = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultListTasks) SetResult(v string) *QueryAllProcessInstancesResponseBodyResultListTasks {
	s.Result = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultListTasks) SetStatus(v string) *QueryAllProcessInstancesResponseBodyResultListTasks {
	s.Status = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultListTasks) SetTaskId(v int64) *QueryAllProcessInstancesResponseBodyResultListTasks {
	s.TaskId = &v
	return s
}

func (s *QueryAllProcessInstancesResponseBodyResultListTasks) SetUserId(v string) *QueryAllProcessInstancesResponseBodyResultListTasks {
	s.UserId = &v
	return s
}

type QueryAllProcessInstancesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAllProcessInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryAllProcessInstancesResponse) SetStatusCode(v int32) *QueryAllProcessInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAllProcessInstancesResponse) SetBody(v *QueryAllProcessInstancesResponseBody) *QueryAllProcessInstancesResponse {
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
	// This parameter is required.
	//
	// example:
	//
	// SWAPP-abcdef-example
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// 应用类型
	AppType *int32 `json:"appType,omitempty" xml:"appType,omitempty"`
	// example:
	//
	// SWAPP-abcdef-example
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// example:
	//
	// 表单业务标识
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 1635151039000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 02501234567890
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-abcdef-example
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-example
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// example:
	//
	// 用于收集休假信息
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// example:
	//
	// 1635151039000
	ModifedTime *int64 `json:"modifedTime,omitempty" xml:"modifedTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 休假申请
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 02501234567890
	OwnerId *string `json:"ownerId,omitempty" xml:"ownerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PUBLISHED(启用), INVALID(停用), SAVED(草稿)
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryFormByBizTypeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryFormByBizTypeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryFormByBizTypeResponseBodyResult) SetAppType(v int32) *QueryFormByBizTypeResponseBodyResult {
	s.AppType = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetAppUuid(v string) *QueryFormByBizTypeResponseBodyResult {
	s.AppUuid = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetBizType(v string) *QueryFormByBizTypeResponseBodyResult {
	s.BizType = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetContent(v string) *QueryFormByBizTypeResponseBodyResult {
	s.Content = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetCreateTime(v int64) *QueryFormByBizTypeResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetCreator(v string) *QueryFormByBizTypeResponseBodyResult {
	s.Creator = &v
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

func (s *QueryFormByBizTypeResponseBodyResult) SetMemo(v string) *QueryFormByBizTypeResponseBodyResult {
	s.Memo = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetModifedTime(v int64) *QueryFormByBizTypeResponseBodyResult {
	s.ModifedTime = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetName(v string) *QueryFormByBizTypeResponseBodyResult {
	s.Name = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetOwnerId(v string) *QueryFormByBizTypeResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *QueryFormByBizTypeResponseBodyResult) SetStatus(v string) *QueryFormByBizTypeResponseBodyResult {
	s.Status = &v
	return s
}

type QueryFormByBizTypeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFormByBizTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryFormByBizTypeResponse) SetStatusCode(v int32) *QueryFormByBizTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFormByBizTypeResponse) SetBody(v *QueryFormByBizTypeResponseBody) *QueryFormByBizTypeResponse {
	s.Body = v
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
	// example:
	//
	// SWAPP-dfeacds-example
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-abcdef-example
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 951a8-8828-430c-b3e-example
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
}

func (s QueryFormInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFormInstanceRequest) GoString() string {
	return s.String()
}

func (s *QueryFormInstanceRequest) SetAppUuid(v string) *QueryFormInstanceRequest {
	s.AppUuid = &v
	return s
}

func (s *QueryFormInstanceRequest) SetFormCode(v string) *QueryFormInstanceRequest {
	s.FormCode = &v
	return s
}

func (s *QueryFormInstanceRequest) SetFormInstanceId(v string) *QueryFormInstanceRequest {
	s.FormInstanceId = &v
	return s
}

type QueryFormInstanceResponseBody struct {
	// example:
	//
	// SWAPP-dfeacds-example
	AppUuid    *string                `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// example:
	//
	// 1631870043000
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 00003
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-abcdef-example
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// This parameter is required.
	FormInstDataList []*QueryFormInstanceResponseBodyFormInstDataList `json:"formInstDataList,omitempty" xml:"formInstDataList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 951a8-8828-430c-b3e-example
	FormInstanceId *string `json:"formInstanceId,omitempty" xml:"formInstanceId,omitempty"`
	// example:
	//
	// 000025
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// example:
	//
	// 1631870043000
	ModifyTimestamp *int64 `json:"modifyTimestamp,omitempty" xml:"modifyTimestamp,omitempty"`
	// example:
	//
	// PROC-abcdef-example
	OutBizCode *string `json:"outBizCode,omitempty" xml:"outBizCode,omitempty"`
	// example:
	//
	// 951a8-8828-430c-b3e-example
	OutInstanceId *string `json:"outInstanceId,omitempty" xml:"outInstanceId,omitempty"`
	// example:
	//
	// xxx提交的表单数据
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QueryFormInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFormInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFormInstanceResponseBody) SetAppUuid(v string) *QueryFormInstanceResponseBody {
	s.AppUuid = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetAttributes(v map[string]interface{}) *QueryFormInstanceResponseBody {
	s.Attributes = v
	return s
}

func (s *QueryFormInstanceResponseBody) SetCreateTimestamp(v int64) *QueryFormInstanceResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetCreator(v string) *QueryFormInstanceResponseBody {
	s.Creator = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetFormCode(v string) *QueryFormInstanceResponseBody {
	s.FormCode = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetFormInstDataList(v []*QueryFormInstanceResponseBodyFormInstDataList) *QueryFormInstanceResponseBody {
	s.FormInstDataList = v
	return s
}

func (s *QueryFormInstanceResponseBody) SetFormInstanceId(v string) *QueryFormInstanceResponseBody {
	s.FormInstanceId = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetModifier(v string) *QueryFormInstanceResponseBody {
	s.Modifier = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetModifyTimestamp(v int64) *QueryFormInstanceResponseBody {
	s.ModifyTimestamp = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetOutBizCode(v string) *QueryFormInstanceResponseBody {
	s.OutBizCode = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetOutInstanceId(v string) *QueryFormInstanceResponseBody {
	s.OutInstanceId = &v
	return s
}

func (s *QueryFormInstanceResponseBody) SetTitle(v string) *QueryFormInstanceResponseBody {
	s.Title = &v
	return s
}

type QueryFormInstanceResponseBodyFormInstDataList struct {
	// This parameter is required.
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// This parameter is required.
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// This parameter is required.
	ExtendValue *string `json:"extendValue,omitempty" xml:"extendValue,omitempty"`
	// This parameter is required.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s QueryFormInstanceResponseBodyFormInstDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryFormInstanceResponseBodyFormInstDataList) GoString() string {
	return s.String()
}

func (s *QueryFormInstanceResponseBodyFormInstDataList) SetBizAlias(v string) *QueryFormInstanceResponseBodyFormInstDataList {
	s.BizAlias = &v
	return s
}

func (s *QueryFormInstanceResponseBodyFormInstDataList) SetComponentType(v string) *QueryFormInstanceResponseBodyFormInstDataList {
	s.ComponentType = &v
	return s
}

func (s *QueryFormInstanceResponseBodyFormInstDataList) SetExtendValue(v string) *QueryFormInstanceResponseBodyFormInstDataList {
	s.ExtendValue = &v
	return s
}

func (s *QueryFormInstanceResponseBodyFormInstDataList) SetKey(v string) *QueryFormInstanceResponseBodyFormInstDataList {
	s.Key = &v
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

type QueryFormInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryFormInstanceResponse) SetStatusCode(v int32) *QueryFormInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFormInstanceResponse) SetBody(v *QueryFormInstanceResponseBody) *QueryFormInstanceResponse {
	s.Body = v
	return s
}

type QueryIntegratedTodoTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryIntegratedTodoTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryIntegratedTodoTaskHeaders) GoString() string {
	return s.String()
}

func (s *QueryIntegratedTodoTaskHeaders) SetCommonHeaders(v map[string]*string) *QueryIntegratedTodoTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryIntegratedTodoTaskHeaders) SetXAcsDingtalkAccessToken(v string) *QueryIntegratedTodoTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryIntegratedTodoTaskRequest struct {
	// example:
	//
	// 1660036833411
	CreateBefore *int64 `json:"createBefore,omitempty" xml:"createBefore,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// manager001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryIntegratedTodoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryIntegratedTodoTaskRequest) GoString() string {
	return s.String()
}

func (s *QueryIntegratedTodoTaskRequest) SetCreateBefore(v int64) *QueryIntegratedTodoTaskRequest {
	s.CreateBefore = &v
	return s
}

func (s *QueryIntegratedTodoTaskRequest) SetPageNumber(v int32) *QueryIntegratedTodoTaskRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryIntegratedTodoTaskRequest) SetPageSize(v int32) *QueryIntegratedTodoTaskRequest {
	s.PageSize = &v
	return s
}

func (s *QueryIntegratedTodoTaskRequest) SetUserId(v string) *QueryIntegratedTodoTaskRequest {
	s.UserId = &v
	return s
}

type QueryIntegratedTodoTaskResponseBody struct {
	Result *QueryIntegratedTodoTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QueryIntegratedTodoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryIntegratedTodoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *QueryIntegratedTodoTaskResponseBody) SetResult(v *QueryIntegratedTodoTaskResponseBodyResult) *QueryIntegratedTodoTaskResponseBody {
	s.Result = v
	return s
}

type QueryIntegratedTodoTaskResponseBodyResult struct {
	HasMore *bool                                            `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List    []*QueryIntegratedTodoTaskResponseBodyResultList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s QueryIntegratedTodoTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryIntegratedTodoTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryIntegratedTodoTaskResponseBodyResult) SetHasMore(v bool) *QueryIntegratedTodoTaskResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *QueryIntegratedTodoTaskResponseBodyResult) SetList(v []*QueryIntegratedTodoTaskResponseBodyResultList) *QueryIntegratedTodoTaskResponseBodyResult {
	s.List = v
	return s
}

type QueryIntegratedTodoTaskResponseBodyResultList struct {
	// example:
	//
	// act_0001
	ActivityId *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
	// example:
	//
	// 2022-10-17T15:12Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 2022-10-17T15:12Z
	FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// example:
	//
	// Siw2WNVZS4KiUt3tTmaNKg04*****809950
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// example:
	//
	// agree
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1234567
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// manager001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryIntegratedTodoTaskResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s QueryIntegratedTodoTaskResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *QueryIntegratedTodoTaskResponseBodyResultList) SetActivityId(v string) *QueryIntegratedTodoTaskResponseBodyResultList {
	s.ActivityId = &v
	return s
}

func (s *QueryIntegratedTodoTaskResponseBodyResultList) SetCreateTime(v string) *QueryIntegratedTodoTaskResponseBodyResultList {
	s.CreateTime = &v
	return s
}

func (s *QueryIntegratedTodoTaskResponseBodyResultList) SetFinishTime(v string) *QueryIntegratedTodoTaskResponseBodyResultList {
	s.FinishTime = &v
	return s
}

func (s *QueryIntegratedTodoTaskResponseBodyResultList) SetProcessInstanceId(v string) *QueryIntegratedTodoTaskResponseBodyResultList {
	s.ProcessInstanceId = &v
	return s
}

func (s *QueryIntegratedTodoTaskResponseBodyResultList) SetResult(v string) *QueryIntegratedTodoTaskResponseBodyResultList {
	s.Result = &v
	return s
}

func (s *QueryIntegratedTodoTaskResponseBodyResultList) SetStatus(v string) *QueryIntegratedTodoTaskResponseBodyResultList {
	s.Status = &v
	return s
}

func (s *QueryIntegratedTodoTaskResponseBodyResultList) SetTaskId(v int64) *QueryIntegratedTodoTaskResponseBodyResultList {
	s.TaskId = &v
	return s
}

func (s *QueryIntegratedTodoTaskResponseBodyResultList) SetUserId(v string) *QueryIntegratedTodoTaskResponseBodyResultList {
	s.UserId = &v
	return s
}

type QueryIntegratedTodoTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryIntegratedTodoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryIntegratedTodoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryIntegratedTodoTaskResponse) GoString() string {
	return s.String()
}

func (s *QueryIntegratedTodoTaskResponse) SetHeaders(v map[string]*string) *QueryIntegratedTodoTaskResponse {
	s.Headers = v
	return s
}

func (s *QueryIntegratedTodoTaskResponse) SetStatusCode(v int32) *QueryIntegratedTodoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryIntegratedTodoTaskResponse) SetBody(v *QueryIntegratedTodoTaskResponseBody) *QueryIntegratedTodoTaskResponse {
	s.Body = v
	return s
}

type QueryProcessByBizCategoryIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QueryProcessByBizCategoryIdHeaders) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessByBizCategoryIdHeaders) GoString() string {
	return s.String()
}

func (s *QueryProcessByBizCategoryIdHeaders) SetCommonHeaders(v map[string]*string) *QueryProcessByBizCategoryIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryProcessByBizCategoryIdHeaders) SetXAcsDingtalkAccessToken(v string) *QueryProcessByBizCategoryIdHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QueryProcessByBizCategoryIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// abc
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// example:
	//
	// manager123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s QueryProcessByBizCategoryIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessByBizCategoryIdRequest) GoString() string {
	return s.String()
}

func (s *QueryProcessByBizCategoryIdRequest) SetBizType(v string) *QueryProcessByBizCategoryIdRequest {
	s.BizType = &v
	return s
}

func (s *QueryProcessByBizCategoryIdRequest) SetUserId(v string) *QueryProcessByBizCategoryIdRequest {
	s.UserId = &v
	return s
}

type QueryProcessByBizCategoryIdResponseBody struct {
	Result []*QueryProcessByBizCategoryIdResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s QueryProcessByBizCategoryIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessByBizCategoryIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryProcessByBizCategoryIdResponseBody) SetResult(v []*QueryProcessByBizCategoryIdResponseBodyResult) *QueryProcessByBizCategoryIdResponseBody {
	s.Result = v
	return s
}

type QueryProcessByBizCategoryIdResponseBodyResult struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	Status      *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryProcessByBizCategoryIdResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessByBizCategoryIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryProcessByBizCategoryIdResponseBodyResult) SetDescription(v string) *QueryProcessByBizCategoryIdResponseBodyResult {
	s.Description = &v
	return s
}

func (s *QueryProcessByBizCategoryIdResponseBodyResult) SetName(v string) *QueryProcessByBizCategoryIdResponseBodyResult {
	s.Name = &v
	return s
}

func (s *QueryProcessByBizCategoryIdResponseBodyResult) SetProcessCode(v string) *QueryProcessByBizCategoryIdResponseBodyResult {
	s.ProcessCode = &v
	return s
}

func (s *QueryProcessByBizCategoryIdResponseBodyResult) SetStatus(v string) *QueryProcessByBizCategoryIdResponseBodyResult {
	s.Status = &v
	return s
}

type QueryProcessByBizCategoryIdResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryProcessByBizCategoryIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryProcessByBizCategoryIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessByBizCategoryIdResponse) GoString() string {
	return s.String()
}

func (s *QueryProcessByBizCategoryIdResponse) SetHeaders(v map[string]*string) *QueryProcessByBizCategoryIdResponse {
	s.Headers = v
	return s
}

func (s *QueryProcessByBizCategoryIdResponse) SetStatusCode(v int32) *QueryProcessByBizCategoryIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryProcessByBizCategoryIdResponse) SetBody(v *QueryProcessByBizCategoryIdResponseBody) *QueryProcessByBizCategoryIdResponse {
	s.Body = v
	return s
}

type QuerySchemaAndProcessHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s QuerySchemaAndProcessHeaders) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaAndProcessHeaders) GoString() string {
	return s.String()
}

func (s *QuerySchemaAndProcessHeaders) SetCommonHeaders(v map[string]*string) *QuerySchemaAndProcessHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QuerySchemaAndProcessHeaders) SetXAcsDingtalkAccessToken(v string) *QuerySchemaAndProcessHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type QuerySchemaAndProcessRequest struct {
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-17428B8C-6C60-xxxx-924C-64F1037AE067
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
}

func (s QuerySchemaAndProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaAndProcessRequest) GoString() string {
	return s.String()
}

func (s *QuerySchemaAndProcessRequest) SetAppUuid(v string) *QuerySchemaAndProcessRequest {
	s.AppUuid = &v
	return s
}

func (s *QuerySchemaAndProcessRequest) SetProcessCode(v string) *QuerySchemaAndProcessRequest {
	s.ProcessCode = &v
	return s
}

type QuerySchemaAndProcessResponseBody struct {
	Result *QuerySchemaAndProcessResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s QuerySchemaAndProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaAndProcessResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySchemaAndProcessResponseBody) SetResult(v *QuerySchemaAndProcessResponseBodyResult) *QuerySchemaAndProcessResponseBody {
	s.Result = v
	return s
}

type QuerySchemaAndProcessResponseBodyResult struct {
	AppType        *int32  `json:"appType,omitempty" xml:"appType,omitempty"`
	Content        *string `json:"content,omitempty" xml:"content,omitempty"`
	HandSignEnable *string `json:"handSignEnable,omitempty" xml:"handSignEnable,omitempty"`
	IconUrl        *string `json:"iconUrl,omitempty" xml:"iconUrl,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	ProcessConfig  *string `json:"processConfig,omitempty" xml:"processConfig,omitempty"`
}

func (s QuerySchemaAndProcessResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaAndProcessResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QuerySchemaAndProcessResponseBodyResult) SetAppType(v int32) *QuerySchemaAndProcessResponseBodyResult {
	s.AppType = &v
	return s
}

func (s *QuerySchemaAndProcessResponseBodyResult) SetContent(v string) *QuerySchemaAndProcessResponseBodyResult {
	s.Content = &v
	return s
}

func (s *QuerySchemaAndProcessResponseBodyResult) SetHandSignEnable(v string) *QuerySchemaAndProcessResponseBodyResult {
	s.HandSignEnable = &v
	return s
}

func (s *QuerySchemaAndProcessResponseBodyResult) SetIconUrl(v string) *QuerySchemaAndProcessResponseBodyResult {
	s.IconUrl = &v
	return s
}

func (s *QuerySchemaAndProcessResponseBodyResult) SetName(v string) *QuerySchemaAndProcessResponseBodyResult {
	s.Name = &v
	return s
}

func (s *QuerySchemaAndProcessResponseBodyResult) SetProcessConfig(v string) *QuerySchemaAndProcessResponseBodyResult {
	s.ProcessConfig = &v
	return s
}

type QuerySchemaAndProcessResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySchemaAndProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySchemaAndProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaAndProcessResponse) GoString() string {
	return s.String()
}

func (s *QuerySchemaAndProcessResponse) SetHeaders(v map[string]*string) *QuerySchemaAndProcessResponse {
	s.Headers = v
	return s
}

func (s *QuerySchemaAndProcessResponse) SetStatusCode(v int32) *QuerySchemaAndProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySchemaAndProcessResponse) SetBody(v *QuerySchemaAndProcessResponseBody) *QuerySchemaAndProcessResponse {
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
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-17428B8C-6C60-xxxx-924C-64F1037AE067
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
}

func (s QuerySchemaByProcessCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaByProcessCodeRequest) GoString() string {
	return s.String()
}

func (s *QuerySchemaByProcessCodeRequest) SetAppUuid(v string) *QuerySchemaByProcessCodeRequest {
	s.AppUuid = &v
	return s
}

func (s *QuerySchemaByProcessCodeRequest) SetProcessCode(v string) *QuerySchemaByProcessCodeRequest {
	s.ProcessCode = &v
	return s
}

type QuerySchemaByProcessCodeResponseBody struct {
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// 0
	AppType *int32 `json:"appType,omitempty" xml:"appType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// example:
	//
	// hrm.xxxx
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 26652461xxxx5992
	CreatorUserId *string `json:"creatorUserId,omitempty" xml:"creatorUserId,omitempty"`
	// example:
	//
	// null
	CustomSetting *string `json:"customSetting,omitempty" xml:"customSetting,omitempty"`
	// example:
	//
	// 0
	EngineType *int32 `json:"engineType,omitempty" xml:"engineType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-17428B8C-6C60-470E-xxxx-64F1037AE067
	FormCode *string `json:"formCode,omitempty" xml:"formCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-28215C3E-00E3-4118-xxxx-4091F828AF2F
	FormUuid *string `json:"formUuid,omitempty" xml:"formUuid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-12-01T10:49Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-12-01T10:49Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// null
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// 1
	ListOrder *int32 `json:"listOrder,omitempty" xml:"listOrder,omitempty"`
	// example:
	//
	// xxxx
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// example:
	//
	// 示例模板
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 25xxxx01
	OwnerIdType *string `json:"ownerIdType,omitempty" xml:"ownerIdType,omitempty"`
	// example:
	//
	// inner
	ProcType *string `json:"procType,omitempty" xml:"procType,omitempty"`
	// This parameter is required.
	SchemaContent *QuerySchemaByProcessCodeResponseBodyResultSchemaContent `json:"schemaContent,omitempty" xml:"schemaContent,omitempty" type:"Struct"`
	// example:
	//
	// PUBLISHED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// PRIVATE
	VisibleRange *string `json:"visibleRange,omitempty" xml:"visibleRange,omitempty"`
}

func (s QuerySchemaByProcessCodeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaByProcessCodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetAppType(v int32) *QuerySchemaByProcessCodeResponseBodyResult {
	s.AppType = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetAppUuid(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.AppUuid = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetBizType(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.BizType = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetCreatorUserId(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.CreatorUserId = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetCustomSetting(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.CustomSetting = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetEngineType(v int32) *QuerySchemaByProcessCodeResponseBodyResult {
	s.EngineType = &v
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

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetGmtCreate(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetGmtModified(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetIcon(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.Icon = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetListOrder(v int32) *QuerySchemaByProcessCodeResponseBodyResult {
	s.ListOrder = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetMemo(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.Memo = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetName(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.Name = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetOwnerIdType(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.OwnerIdType = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetProcType(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.ProcType = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetSchemaContent(v *QuerySchemaByProcessCodeResponseBodyResultSchemaContent) *QuerySchemaByProcessCodeResponseBodyResult {
	s.SchemaContent = v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetStatus(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.Status = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResult) SetVisibleRange(v string) *QuerySchemaByProcessCodeResponseBodyResult {
	s.VisibleRange = &v
	return s
}

type QuerySchemaByProcessCodeResponseBodyResultSchemaContent struct {
	// example:
	//
	// common
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// This parameter is required.
	Items []*QuerySchemaByProcessCodeResponseBodyResultSchemaContentItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 示例模板
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContent) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContent) GoString() string {
	return s.String()
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContent) SetIcon(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContent {
	s.Icon = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContent) SetItems(v []*QuerySchemaByProcessCodeResponseBodyResultSchemaContentItems) *QuerySchemaByProcessCodeResponseBodyResultSchemaContent {
	s.Items = v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContent) SetTitle(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContent {
	s.Title = &v
	return s
}

type QuerySchemaByProcessCodeResponseBodyResultSchemaContentItems struct {
	Children []*QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsChildren `json:"children,omitempty" xml:"children,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// TextField
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// This parameter is required.
	Props *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItems) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItems) GoString() string {
	return s.String()
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItems) SetChildren(v []*QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsChildren) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItems {
	s.Children = v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItems) SetComponentName(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItems {
	s.ComponentName = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItems) SetProps(v *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItems {
	s.Props = v
	return s
}

type QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsChildren struct {
	// This parameter is required.
	//
	// example:
	//
	// TextField
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// This parameter is required.
	Props *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsChildrenProps `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsChildren) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsChildren) GoString() string {
	return s.String()
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsChildren) SetComponentName(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsChildren {
	s.ComponentName = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsChildren) SetProps(v *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsChildrenProps) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsChildren {
	s.Props = v
	return s
}

type QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsChildrenProps struct {
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// This parameter is required.
	Id       *string   `json:"id,omitempty" xml:"id,omitempty"`
	Label    *string   `json:"label,omitempty" xml:"label,omitempty"`
	Options  []*string `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	Required *bool     `json:"required,omitempty" xml:"required,omitempty"`
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsChildrenProps) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsChildrenProps) GoString() string {
	return s.String()
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsChildrenProps) SetBizAlias(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsChildrenProps {
	s.BizAlias = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsChildrenProps) SetId(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsChildrenProps {
	s.Id = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsChildrenProps) SetLabel(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsChildrenProps {
	s.Label = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsChildrenProps) SetOptions(v []*string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsChildrenProps {
	s.Options = v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsChildrenProps) SetRequired(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsChildrenProps {
	s.Required = &v
	return s
}

type QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps struct {
	// example:
	//
	// 添加
	ActionName *string `json:"actionName,omitempty" xml:"actionName,omitempty"`
	// example:
	//
	// top
	Align *string `json:"align,omitempty" xml:"align,omitempty"`
	// example:
	//
	// 1234567
	AppId *int64 `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// true
	AsyncCondition *bool `json:"asyncCondition,omitempty" xml:"asyncCondition,omitempty"`
	// example:
	//
	// 请假
	AttendTypeLabel *string                                                                             `json:"attendTypeLabel,omitempty" xml:"attendTypeLabel,omitempty"`
	BehaviorLinkage []*QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkage `json:"behaviorLinkage,omitempty" xml:"behaviorLinkage,omitempty" type:"Repeated"`
	// example:
	//
	// 我的单行输入框
	BizAlias *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	// example:
	//
	// hrm.xxxx
	BizType           *string          `json:"bizType,omitempty" xml:"bizType,omitempty"`
	ChildFieldVisible map[string]*bool `json:"childFieldVisible,omitempty" xml:"childFieldVisible,omitempty"`
	// example:
	//
	// 1
	Choice *int32 `json:"choice,omitempty" xml:"choice,omitempty"`
	// example:
	//
	// xxxx
	CommonBizType *string `json:"commonBizType,omitempty" xml:"commonBizType,omitempty"`
	// example:
	//
	// true
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// example:
	//
	// true
	Duration *bool `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// xxxx
	DurationLabel *string `json:"durationLabel,omitempty" xml:"durationLabel,omitempty"`
	// example:
	//
	// true
	ESign *bool `json:"eSign,omitempty" xml:"eSign,omitempty"`
	// example:
	//
	// true
	Extract *bool `json:"extract,omitempty" xml:"extract,omitempty"`
	// example:
	//
	// xxxx
	FieldsInfo *string `json:"fieldsInfo,omitempty" xml:"fieldsInfo,omitempty"`
	// example:
	//
	// yyyy-MM-dd
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// example:
	//
	// xxxx
	Formula *string `json:"formula,omitempty" xml:"formula,omitempty"`
	// example:
	//
	// true
	Hidden *bool `json:"hidden,omitempty" xml:"hidden,omitempty"`
	// example:
	//
	// true
	HiddenInApprovalDetail *bool `json:"hiddenInApprovalDetail,omitempty" xml:"hiddenInApprovalDetail,omitempty"`
	// example:
	//
	// true
	HideLabel *bool `json:"hideLabel,omitempty" xml:"hideLabel,omitempty"`
	// example:
	//
	// "[{\"name\":\"\open"}]"
	HolidayOptions []map[string]*string `json:"holidayOptions,omitempty" xml:"holidayOptions,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// TextField-K2AD4O5B
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 单行输入框
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// example:
	//
	// true
	LabelEditableFreeze *bool `json:"labelEditableFreeze,omitempty" xml:"labelEditableFreeze,omitempty"`
	// example:
	//
	// xxxx
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// example:
	//
	// xxxx
	MainTitle *string `json:"mainTitle,omitempty" xml:"mainTitle,omitempty"`
	// example:
	//
	// 1
	NotPrint *string `json:"notPrint,omitempty" xml:"notPrint,omitempty"`
	// example:
	//
	// 1
	NotUpper   *string                                                                        `json:"notUpper,omitempty" xml:"notUpper,omitempty"`
	ObjOptions []*QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsObjOptions `json:"objOptions,omitempty" xml:"objOptions,omitempty" type:"Repeated"`
	Options    []*string                                                                      `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// example:
	//
	// true
	PayEnable *bool `json:"payEnable,omitempty" xml:"payEnable,omitempty"`
	// example:
	//
	// 请输入文字
	Placeholder *string                                                                `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	Push        *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsPush `json:"push,omitempty" xml:"push,omitempty" type:"Struct"`
	// example:
	//
	// true
	PushToAttendance *bool `json:"pushToAttendance,omitempty" xml:"pushToAttendance,omitempty"`
	// example:
	//
	// 1
	PushToCalendar *int32 `json:"pushToCalendar,omitempty" xml:"pushToCalendar,omitempty"`
	// example:
	//
	// true
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// example:
	//
	// true
	RequiredEditableFreeze *bool `json:"requiredEditableFreeze,omitempty" xml:"requiredEditableFreeze,omitempty"`
	// example:
	//
	// true
	ShowAttendOptions *bool `json:"showAttendOptions,omitempty" xml:"showAttendOptions,omitempty"`
	// example:
	//
	// true
	StaffStatusEnabled *bool                                                                         `json:"staffStatusEnabled,omitempty" xml:"staffStatusEnabled,omitempty"`
	StatField          []*QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsStatField `json:"statField,omitempty" xml:"statField,omitempty" type:"Repeated"`
	// example:
	//
	// list
	TableViewMode *string `json:"tableViewMode,omitempty" xml:"tableViewMode,omitempty"`
	// example:
	//
	// 天
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// example:
	//
	// true
	UseCalendar *bool `json:"useCalendar,omitempty" xml:"useCalendar,omitempty"`
	// example:
	//
	// true
	VerticalPrint *bool `json:"verticalPrint,omitempty" xml:"verticalPrint,omitempty"`
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) GoString() string {
	return s.String()
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetActionName(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.ActionName = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetAlign(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Align = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetAppId(v int64) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.AppId = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetAsyncCondition(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.AsyncCondition = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetAttendTypeLabel(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.AttendTypeLabel = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetBehaviorLinkage(v []*QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkage) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.BehaviorLinkage = v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetBizAlias(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.BizAlias = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetBizType(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.BizType = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetChildFieldVisible(v map[string]*bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.ChildFieldVisible = v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetChoice(v int32) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Choice = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetCommonBizType(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.CommonBizType = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetDisabled(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Disabled = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetDuration(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Duration = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetDurationLabel(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.DurationLabel = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetESign(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.ESign = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetExtract(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Extract = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetFieldsInfo(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.FieldsInfo = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetFormat(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Format = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetFormula(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Formula = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetHidden(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Hidden = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetHiddenInApprovalDetail(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.HiddenInApprovalDetail = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetHideLabel(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.HideLabel = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetHolidayOptions(v []map[string]*string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.HolidayOptions = v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetId(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Id = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetLabel(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Label = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetLabelEditableFreeze(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.LabelEditableFreeze = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetLink(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Link = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetMainTitle(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.MainTitle = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetNotPrint(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.NotPrint = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetNotUpper(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.NotUpper = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetObjOptions(v []*QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsObjOptions) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.ObjOptions = v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetOptions(v []*string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Options = v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetPayEnable(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.PayEnable = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetPlaceholder(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Placeholder = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetPush(v *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsPush) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Push = v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetPushToAttendance(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.PushToAttendance = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetPushToCalendar(v int32) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.PushToCalendar = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetRequired(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Required = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetRequiredEditableFreeze(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.RequiredEditableFreeze = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetShowAttendOptions(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.ShowAttendOptions = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetStaffStatusEnabled(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.StaffStatusEnabled = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetStatField(v []*QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsStatField) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.StatField = v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetTableViewMode(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.TableViewMode = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetUnit(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.Unit = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetUseCalendar(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.UseCalendar = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps) SetVerticalPrint(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsProps {
	s.VerticalPrint = &v
	return s
}

type QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkage struct {
	Targets []*QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets `json:"targets,omitempty" xml:"targets,omitempty" type:"Repeated"`
	// example:
	//
	// xxxx
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkage) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkage) GoString() string {
	return s.String()
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkage) SetTargets(v []*QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkage {
	s.Targets = v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkage) SetValue(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkage {
	s.Value = &v
	return s
}

type QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets struct {
	// example:
	//
	// xxxx
	Behavior *string `json:"behavior,omitempty" xml:"behavior,omitempty"`
	// example:
	//
	// TextField-K2AD4O5B
	FieldId *string `json:"fieldId,omitempty" xml:"fieldId,omitempty"`
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets) GoString() string {
	return s.String()
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets) SetBehavior(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets {
	s.Behavior = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets) SetFieldId(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsBehaviorLinkageTargets {
	s.FieldId = &v
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
	// example:
	//
	// 1
	AttendanceRule *int32 `json:"attendanceRule,omitempty" xml:"attendanceRule,omitempty"`
	// example:
	//
	// 1
	PushSwitch *int32 `json:"pushSwitch,omitempty" xml:"pushSwitch,omitempty"`
	// example:
	//
	// xxxx
	PushTag *string `json:"pushTag,omitempty" xml:"pushTag,omitempty"`
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsPush) String() string {
	return tea.Prettify(s)
}

func (s QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsPush) GoString() string {
	return s.String()
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsPush) SetAttendanceRule(v int32) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsPush {
	s.AttendanceRule = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsPush) SetPushSwitch(v int32) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsPush {
	s.PushSwitch = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsPush) SetPushTag(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsPush {
	s.PushTag = &v
	return s
}

type QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsStatField struct {
	// example:
	//
	// TextField-K2AD4O5B
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 单行输入框
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// example:
	//
	// xxxx
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// example:
	//
	// true
	Upper *bool `json:"upper,omitempty" xml:"upper,omitempty"`
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

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsStatField) SetUnit(v string) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsStatField {
	s.Unit = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsStatField) SetUpper(v bool) *QuerySchemaByProcessCodeResponseBodyResultSchemaContentItemsPropsStatField {
	s.Upper = &v
	return s
}

type QuerySchemaByProcessCodeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySchemaByProcessCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QuerySchemaByProcessCodeResponse) SetStatusCode(v int32) *QuerySchemaByProcessCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySchemaByProcessCodeResponse) SetBody(v *QuerySchemaByProcessCodeResponseBody) *QuerySchemaByProcessCodeResponse {
	s.Body = v
	return s
}

type RedirectWorkflowTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s RedirectWorkflowTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s RedirectWorkflowTaskHeaders) GoString() string {
	return s.String()
}

func (s *RedirectWorkflowTaskHeaders) SetCommonHeaders(v map[string]*string) *RedirectWorkflowTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RedirectWorkflowTaskHeaders) SetXAcsDingtalkAccessToken(v string) *RedirectWorkflowTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type RedirectWorkflowTaskRequest struct {
	// example:
	//
	// test
	ActionName *string                          `json:"actionName,omitempty" xml:"actionName,omitempty"`
	File       *RedirectWorkflowTaskRequestFile `json:"file,omitempty" xml:"file,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// manager001
	OperateUserId *string `json:"operateUserId,omitempty" xml:"operateUserId,omitempty"`
	// example:
	//
	// 请XX帮忙审批一下
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// manager001
	ToUserId *string `json:"toUserId,omitempty" xml:"toUserId,omitempty"`
}

func (s RedirectWorkflowTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s RedirectWorkflowTaskRequest) GoString() string {
	return s.String()
}

func (s *RedirectWorkflowTaskRequest) SetActionName(v string) *RedirectWorkflowTaskRequest {
	s.ActionName = &v
	return s
}

func (s *RedirectWorkflowTaskRequest) SetFile(v *RedirectWorkflowTaskRequestFile) *RedirectWorkflowTaskRequest {
	s.File = v
	return s
}

func (s *RedirectWorkflowTaskRequest) SetOperateUserId(v string) *RedirectWorkflowTaskRequest {
	s.OperateUserId = &v
	return s
}

func (s *RedirectWorkflowTaskRequest) SetRemark(v string) *RedirectWorkflowTaskRequest {
	s.Remark = &v
	return s
}

func (s *RedirectWorkflowTaskRequest) SetTaskId(v int64) *RedirectWorkflowTaskRequest {
	s.TaskId = &v
	return s
}

func (s *RedirectWorkflowTaskRequest) SetToUserId(v string) *RedirectWorkflowTaskRequest {
	s.ToUserId = &v
	return s
}

type RedirectWorkflowTaskRequestFile struct {
	Attachments []*RedirectWorkflowTaskRequestFileAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
	Photos      []*string                                     `json:"photos,omitempty" xml:"photos,omitempty" type:"Repeated"`
}

func (s RedirectWorkflowTaskRequestFile) String() string {
	return tea.Prettify(s)
}

func (s RedirectWorkflowTaskRequestFile) GoString() string {
	return s.String()
}

func (s *RedirectWorkflowTaskRequestFile) SetAttachments(v []*RedirectWorkflowTaskRequestFileAttachments) *RedirectWorkflowTaskRequestFile {
	s.Attachments = v
	return s
}

func (s *RedirectWorkflowTaskRequestFile) SetPhotos(v []*string) *RedirectWorkflowTaskRequestFile {
	s.Photos = v
	return s
}

type RedirectWorkflowTaskRequestFileAttachments struct {
	// example:
	//
	// B1oQixxxx
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// example:
	//
	// 文件名称。
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// 1024
	FileSize *string `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// example:
	//
	// file
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// 123
	SpaceId *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
}

func (s RedirectWorkflowTaskRequestFileAttachments) String() string {
	return tea.Prettify(s)
}

func (s RedirectWorkflowTaskRequestFileAttachments) GoString() string {
	return s.String()
}

func (s *RedirectWorkflowTaskRequestFileAttachments) SetFileId(v string) *RedirectWorkflowTaskRequestFileAttachments {
	s.FileId = &v
	return s
}

func (s *RedirectWorkflowTaskRequestFileAttachments) SetFileName(v string) *RedirectWorkflowTaskRequestFileAttachments {
	s.FileName = &v
	return s
}

func (s *RedirectWorkflowTaskRequestFileAttachments) SetFileSize(v string) *RedirectWorkflowTaskRequestFileAttachments {
	s.FileSize = &v
	return s
}

func (s *RedirectWorkflowTaskRequestFileAttachments) SetFileType(v string) *RedirectWorkflowTaskRequestFileAttachments {
	s.FileType = &v
	return s
}

func (s *RedirectWorkflowTaskRequestFileAttachments) SetSpaceId(v string) *RedirectWorkflowTaskRequestFileAttachments {
	s.SpaceId = &v
	return s
}

type RedirectWorkflowTaskResponseBody struct {
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RedirectWorkflowTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RedirectWorkflowTaskResponseBody) GoString() string {
	return s.String()
}

func (s *RedirectWorkflowTaskResponseBody) SetResult(v bool) *RedirectWorkflowTaskResponseBody {
	s.Result = &v
	return s
}

type RedirectWorkflowTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RedirectWorkflowTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RedirectWorkflowTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s RedirectWorkflowTaskResponse) GoString() string {
	return s.String()
}

func (s *RedirectWorkflowTaskResponse) SetHeaders(v map[string]*string) *RedirectWorkflowTaskResponse {
	s.Headers = v
	return s
}

func (s *RedirectWorkflowTaskResponse) SetStatusCode(v int32) *RedirectWorkflowTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *RedirectWorkflowTaskResponse) SetBody(v *RedirectWorkflowTaskResponseBody) *RedirectWorkflowTaskResponse {
	s.Body = v
	return s
}

type SaveIntegratedInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SaveIntegratedInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveIntegratedInstanceHeaders) GoString() string {
	return s.String()
}

func (s *SaveIntegratedInstanceHeaders) SetCommonHeaders(v map[string]*string) *SaveIntegratedInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveIntegratedInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *SaveIntegratedInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SaveIntegratedInstanceRequest struct {
	// example:
	//
	// "{\"mykey\": \"myData\"}"
	BizData                *string                                                `json:"bizData,omitempty" xml:"bizData,omitempty"`
	FeatureConfig          *SaveIntegratedInstanceRequestFeatureConfig            `json:"featureConfig,omitempty" xml:"featureConfig,omitempty" type:"Struct"`
	FormComponentValueList []*SaveIntegratedInstanceRequestFormComponentValueList `json:"formComponentValueList,omitempty" xml:"formComponentValueList,omitempty" type:"Repeated"`
	Notifiers              []*SaveIntegratedInstanceRequestNotifiers              `json:"notifiers,omitempty" xml:"notifiers,omitempty" type:"Repeated"`
	// This parameter is required.
	OriginatorUserId *string `json:"originatorUserId,omitempty" xml:"originatorUserId,omitempty"`
	// This parameter is required.
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://www.dingtalk.com/
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SaveIntegratedInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveIntegratedInstanceRequest) GoString() string {
	return s.String()
}

func (s *SaveIntegratedInstanceRequest) SetBizData(v string) *SaveIntegratedInstanceRequest {
	s.BizData = &v
	return s
}

func (s *SaveIntegratedInstanceRequest) SetFeatureConfig(v *SaveIntegratedInstanceRequestFeatureConfig) *SaveIntegratedInstanceRequest {
	s.FeatureConfig = v
	return s
}

func (s *SaveIntegratedInstanceRequest) SetFormComponentValueList(v []*SaveIntegratedInstanceRequestFormComponentValueList) *SaveIntegratedInstanceRequest {
	s.FormComponentValueList = v
	return s
}

func (s *SaveIntegratedInstanceRequest) SetNotifiers(v []*SaveIntegratedInstanceRequestNotifiers) *SaveIntegratedInstanceRequest {
	s.Notifiers = v
	return s
}

func (s *SaveIntegratedInstanceRequest) SetOriginatorUserId(v string) *SaveIntegratedInstanceRequest {
	s.OriginatorUserId = &v
	return s
}

func (s *SaveIntegratedInstanceRequest) SetProcessCode(v string) *SaveIntegratedInstanceRequest {
	s.ProcessCode = &v
	return s
}

func (s *SaveIntegratedInstanceRequest) SetTitle(v string) *SaveIntegratedInstanceRequest {
	s.Title = &v
	return s
}

func (s *SaveIntegratedInstanceRequest) SetUrl(v string) *SaveIntegratedInstanceRequest {
	s.Url = &v
	return s
}

type SaveIntegratedInstanceRequestFeatureConfig struct {
	Features []*SaveIntegratedInstanceRequestFeatureConfigFeatures `json:"features,omitempty" xml:"features,omitempty" type:"Repeated"`
}

func (s SaveIntegratedInstanceRequestFeatureConfig) String() string {
	return tea.Prettify(s)
}

func (s SaveIntegratedInstanceRequestFeatureConfig) GoString() string {
	return s.String()
}

func (s *SaveIntegratedInstanceRequestFeatureConfig) SetFeatures(v []*SaveIntegratedInstanceRequestFeatureConfigFeatures) *SaveIntegratedInstanceRequestFeatureConfig {
	s.Features = v
	return s
}

type SaveIntegratedInstanceRequestFeatureConfigFeatures struct {
	Callback *SaveIntegratedInstanceRequestFeatureConfigFeaturesCallback `json:"callback,omitempty" xml:"callback,omitempty" type:"Struct"`
	// if can be null:
	// true
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// example:
	//
	// www.dingtalk.com
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	// example:
	//
	// abc
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// www.dingtalk.com
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	// example:
	//
	// ORIGIN
	RunType *string `json:"runType,omitempty" xml:"runType,omitempty"`
}

func (s SaveIntegratedInstanceRequestFeatureConfigFeatures) String() string {
	return tea.Prettify(s)
}

func (s SaveIntegratedInstanceRequestFeatureConfigFeatures) GoString() string {
	return s.String()
}

func (s *SaveIntegratedInstanceRequestFeatureConfigFeatures) SetCallback(v *SaveIntegratedInstanceRequestFeatureConfigFeaturesCallback) *SaveIntegratedInstanceRequestFeatureConfigFeatures {
	s.Callback = v
	return s
}

func (s *SaveIntegratedInstanceRequestFeatureConfigFeatures) SetConfig(v string) *SaveIntegratedInstanceRequestFeatureConfigFeatures {
	s.Config = &v
	return s
}

func (s *SaveIntegratedInstanceRequestFeatureConfigFeatures) SetMobileUrl(v string) *SaveIntegratedInstanceRequestFeatureConfigFeatures {
	s.MobileUrl = &v
	return s
}

func (s *SaveIntegratedInstanceRequestFeatureConfigFeatures) SetName(v string) *SaveIntegratedInstanceRequestFeatureConfigFeatures {
	s.Name = &v
	return s
}

func (s *SaveIntegratedInstanceRequestFeatureConfigFeatures) SetPcUrl(v string) *SaveIntegratedInstanceRequestFeatureConfigFeatures {
	s.PcUrl = &v
	return s
}

func (s *SaveIntegratedInstanceRequestFeatureConfigFeatures) SetRunType(v string) *SaveIntegratedInstanceRequestFeatureConfigFeatures {
	s.RunType = &v
	return s
}

type SaveIntegratedInstanceRequestFeatureConfigFeaturesCallback struct {
	// example:
	//
	// abc
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// example:
	//
	// abc
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s SaveIntegratedInstanceRequestFeatureConfigFeaturesCallback) String() string {
	return tea.Prettify(s)
}

func (s SaveIntegratedInstanceRequestFeatureConfigFeaturesCallback) GoString() string {
	return s.String()
}

func (s *SaveIntegratedInstanceRequestFeatureConfigFeaturesCallback) SetApiKey(v string) *SaveIntegratedInstanceRequestFeatureConfigFeaturesCallback {
	s.ApiKey = &v
	return s
}

func (s *SaveIntegratedInstanceRequestFeatureConfigFeaturesCallback) SetAppUuid(v string) *SaveIntegratedInstanceRequestFeatureConfigFeaturesCallback {
	s.AppUuid = &v
	return s
}

func (s *SaveIntegratedInstanceRequestFeatureConfigFeaturesCallback) SetVersion(v string) *SaveIntegratedInstanceRequestFeatureConfigFeaturesCallback {
	s.Version = &v
	return s
}

type SaveIntegratedInstanceRequestFormComponentValueList struct {
	BizAlias      *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	ExtValue      *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	Id            *string `json:"id,omitempty" xml:"id,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	Value         *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s SaveIntegratedInstanceRequestFormComponentValueList) String() string {
	return tea.Prettify(s)
}

func (s SaveIntegratedInstanceRequestFormComponentValueList) GoString() string {
	return s.String()
}

func (s *SaveIntegratedInstanceRequestFormComponentValueList) SetBizAlias(v string) *SaveIntegratedInstanceRequestFormComponentValueList {
	s.BizAlias = &v
	return s
}

func (s *SaveIntegratedInstanceRequestFormComponentValueList) SetComponentType(v string) *SaveIntegratedInstanceRequestFormComponentValueList {
	s.ComponentType = &v
	return s
}

func (s *SaveIntegratedInstanceRequestFormComponentValueList) SetExtValue(v string) *SaveIntegratedInstanceRequestFormComponentValueList {
	s.ExtValue = &v
	return s
}

func (s *SaveIntegratedInstanceRequestFormComponentValueList) SetId(v string) *SaveIntegratedInstanceRequestFormComponentValueList {
	s.Id = &v
	return s
}

func (s *SaveIntegratedInstanceRequestFormComponentValueList) SetName(v string) *SaveIntegratedInstanceRequestFormComponentValueList {
	s.Name = &v
	return s
}

func (s *SaveIntegratedInstanceRequestFormComponentValueList) SetValue(v string) *SaveIntegratedInstanceRequestFormComponentValueList {
	s.Value = &v
	return s
}

type SaveIntegratedInstanceRequestNotifiers struct {
	// example:
	//
	// start
	Position *string `json:"position,omitempty" xml:"position,omitempty"`
	// example:
	//
	// manager001
	Userid *string `json:"userid,omitempty" xml:"userid,omitempty"`
}

func (s SaveIntegratedInstanceRequestNotifiers) String() string {
	return tea.Prettify(s)
}

func (s SaveIntegratedInstanceRequestNotifiers) GoString() string {
	return s.String()
}

func (s *SaveIntegratedInstanceRequestNotifiers) SetPosition(v string) *SaveIntegratedInstanceRequestNotifiers {
	s.Position = &v
	return s
}

func (s *SaveIntegratedInstanceRequestNotifiers) SetUserid(v string) *SaveIntegratedInstanceRequestNotifiers {
	s.Userid = &v
	return s
}

type SaveIntegratedInstanceResponseBody struct {
	Result *SaveIntegratedInstanceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s SaveIntegratedInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveIntegratedInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *SaveIntegratedInstanceResponseBody) SetResult(v *SaveIntegratedInstanceResponseBodyResult) *SaveIntegratedInstanceResponseBody {
	s.Result = v
	return s
}

type SaveIntegratedInstanceResponseBodyResult struct {
	// example:
	//
	// proc-abc
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
}

func (s SaveIntegratedInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SaveIntegratedInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SaveIntegratedInstanceResponseBodyResult) SetProcessInstanceId(v string) *SaveIntegratedInstanceResponseBodyResult {
	s.ProcessInstanceId = &v
	return s
}

type SaveIntegratedInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveIntegratedInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveIntegratedInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveIntegratedInstanceResponse) GoString() string {
	return s.String()
}

func (s *SaveIntegratedInstanceResponse) SetHeaders(v map[string]*string) *SaveIntegratedInstanceResponse {
	s.Headers = v
	return s
}

func (s *SaveIntegratedInstanceResponse) SetStatusCode(v int32) *SaveIntegratedInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveIntegratedInstanceResponse) SetBody(v *SaveIntegratedInstanceResponseBody) *SaveIntegratedInstanceResponse {
	s.Body = v
	return s
}

type SaveProcessHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s SaveProcessHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveProcessHeaders) GoString() string {
	return s.String()
}

func (s *SaveProcessHeaders) SetCommonHeaders(v map[string]*string) *SaveProcessHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveProcessHeaders) SetXAcsDingtalkAccessToken(v string) *SaveProcessHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type SaveProcessRequest struct {
	// example:
	//
	// 用于员工差旅费用报销使用
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	FormComponents []*FormComponent `json:"formComponents,omitempty" xml:"formComponents,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 出差报销审批
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// proc-abc
	ProcessCode          *string                                 `json:"processCode,omitempty" xml:"processCode,omitempty"`
	ProcessFeatureConfig *SaveProcessRequestProcessFeatureConfig `json:"processFeatureConfig,omitempty" xml:"processFeatureConfig,omitempty" type:"Struct"`
	// Deprecated
	//
	// if can be null:
	// true
	TemplateConfig *SaveProcessRequestTemplateConfig `json:"templateConfig,omitempty" xml:"templateConfig,omitempty" type:"Struct"`
}

func (s SaveProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveProcessRequest) GoString() string {
	return s.String()
}

func (s *SaveProcessRequest) SetDescription(v string) *SaveProcessRequest {
	s.Description = &v
	return s
}

func (s *SaveProcessRequest) SetFormComponents(v []*FormComponent) *SaveProcessRequest {
	s.FormComponents = v
	return s
}

func (s *SaveProcessRequest) SetName(v string) *SaveProcessRequest {
	s.Name = &v
	return s
}

func (s *SaveProcessRequest) SetProcessCode(v string) *SaveProcessRequest {
	s.ProcessCode = &v
	return s
}

func (s *SaveProcessRequest) SetProcessFeatureConfig(v *SaveProcessRequestProcessFeatureConfig) *SaveProcessRequest {
	s.ProcessFeatureConfig = v
	return s
}

func (s *SaveProcessRequest) SetTemplateConfig(v *SaveProcessRequestTemplateConfig) *SaveProcessRequest {
	s.TemplateConfig = v
	return s
}

type SaveProcessRequestProcessFeatureConfig struct {
	Features []*SaveProcessRequestProcessFeatureConfigFeatures `json:"features,omitempty" xml:"features,omitempty" type:"Repeated"`
}

func (s SaveProcessRequestProcessFeatureConfig) String() string {
	return tea.Prettify(s)
}

func (s SaveProcessRequestProcessFeatureConfig) GoString() string {
	return s.String()
}

func (s *SaveProcessRequestProcessFeatureConfig) SetFeatures(v []*SaveProcessRequestProcessFeatureConfigFeatures) *SaveProcessRequestProcessFeatureConfig {
	s.Features = v
	return s
}

type SaveProcessRequestProcessFeatureConfigFeatures struct {
	Callback *SaveProcessRequestProcessFeatureConfigFeaturesCallback `json:"callback,omitempty" xml:"callback,omitempty" type:"Struct"`
	// if can be null:
	// true
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// example:
	//
	// www.dingtalk.com
	MobileUrl *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	// example:
	//
	// abc
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// www.dingtalk.com
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	// example:
	//
	// ORIGIN
	RunType *string `json:"runType,omitempty" xml:"runType,omitempty"`
}

func (s SaveProcessRequestProcessFeatureConfigFeatures) String() string {
	return tea.Prettify(s)
}

func (s SaveProcessRequestProcessFeatureConfigFeatures) GoString() string {
	return s.String()
}

func (s *SaveProcessRequestProcessFeatureConfigFeatures) SetCallback(v *SaveProcessRequestProcessFeatureConfigFeaturesCallback) *SaveProcessRequestProcessFeatureConfigFeatures {
	s.Callback = v
	return s
}

func (s *SaveProcessRequestProcessFeatureConfigFeatures) SetConfig(v string) *SaveProcessRequestProcessFeatureConfigFeatures {
	s.Config = &v
	return s
}

func (s *SaveProcessRequestProcessFeatureConfigFeatures) SetMobileUrl(v string) *SaveProcessRequestProcessFeatureConfigFeatures {
	s.MobileUrl = &v
	return s
}

func (s *SaveProcessRequestProcessFeatureConfigFeatures) SetName(v string) *SaveProcessRequestProcessFeatureConfigFeatures {
	s.Name = &v
	return s
}

func (s *SaveProcessRequestProcessFeatureConfigFeatures) SetPcUrl(v string) *SaveProcessRequestProcessFeatureConfigFeatures {
	s.PcUrl = &v
	return s
}

func (s *SaveProcessRequestProcessFeatureConfigFeatures) SetRunType(v string) *SaveProcessRequestProcessFeatureConfigFeatures {
	s.RunType = &v
	return s
}

type SaveProcessRequestProcessFeatureConfigFeaturesCallback struct {
	// example:
	//
	// abc
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// example:
	//
	// abc
	AppUuid *string `json:"appUuid,omitempty" xml:"appUuid,omitempty"`
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s SaveProcessRequestProcessFeatureConfigFeaturesCallback) String() string {
	return tea.Prettify(s)
}

func (s SaveProcessRequestProcessFeatureConfigFeaturesCallback) GoString() string {
	return s.String()
}

func (s *SaveProcessRequestProcessFeatureConfigFeaturesCallback) SetApiKey(v string) *SaveProcessRequestProcessFeatureConfigFeaturesCallback {
	s.ApiKey = &v
	return s
}

func (s *SaveProcessRequestProcessFeatureConfigFeaturesCallback) SetAppUuid(v string) *SaveProcessRequestProcessFeatureConfigFeaturesCallback {
	s.AppUuid = &v
	return s
}

func (s *SaveProcessRequestProcessFeatureConfigFeaturesCallback) SetVersion(v string) *SaveProcessRequestProcessFeatureConfigFeaturesCallback {
	s.Version = &v
	return s
}

type SaveProcessRequestTemplateConfig struct {
	// Deprecated
	//
	// example:
	//
	// https://open.dingtalk.com/
	CreateInstanceMobileUrl *string `json:"createInstanceMobileUrl,omitempty" xml:"createInstanceMobileUrl,omitempty"`
	// Deprecated
	//
	// example:
	//
	// https://open.dingtalk.com/
	CreateInstancePcUrl *string `json:"createInstancePcUrl,omitempty" xml:"createInstancePcUrl,omitempty"`
	// if can be null:
	// true
	DisableSendCard *bool `json:"disableSendCard,omitempty" xml:"disableSendCard,omitempty"`
	// example:
	//
	// true
	Hidden *bool `json:"hidden,omitempty" xml:"hidden,omitempty"`
	// Deprecated
	//
	// if can be null:
	// true
	//
	// example:
	//
	// https://open.dingtalk.com/
	TemplateEditUrl *string `json:"templateEditUrl,omitempty" xml:"templateEditUrl,omitempty"`
}

func (s SaveProcessRequestTemplateConfig) String() string {
	return tea.Prettify(s)
}

func (s SaveProcessRequestTemplateConfig) GoString() string {
	return s.String()
}

func (s *SaveProcessRequestTemplateConfig) SetCreateInstanceMobileUrl(v string) *SaveProcessRequestTemplateConfig {
	s.CreateInstanceMobileUrl = &v
	return s
}

func (s *SaveProcessRequestTemplateConfig) SetCreateInstancePcUrl(v string) *SaveProcessRequestTemplateConfig {
	s.CreateInstancePcUrl = &v
	return s
}

func (s *SaveProcessRequestTemplateConfig) SetDisableSendCard(v bool) *SaveProcessRequestTemplateConfig {
	s.DisableSendCard = &v
	return s
}

func (s *SaveProcessRequestTemplateConfig) SetHidden(v bool) *SaveProcessRequestTemplateConfig {
	s.Hidden = &v
	return s
}

func (s *SaveProcessRequestTemplateConfig) SetTemplateEditUrl(v string) *SaveProcessRequestTemplateConfig {
	s.TemplateEditUrl = &v
	return s
}

type SaveProcessResponseBody struct {
	Result *SaveProcessResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s SaveProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveProcessResponseBody) GoString() string {
	return s.String()
}

func (s *SaveProcessResponseBody) SetResult(v *SaveProcessResponseBodyResult) *SaveProcessResponseBody {
	s.Result = v
	return s
}

type SaveProcessResponseBodyResult struct {
	// This parameter is required.
	//
	// example:
	//
	// PROC-abcdef-example
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
}

func (s SaveProcessResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s SaveProcessResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SaveProcessResponseBodyResult) SetProcessCode(v string) *SaveProcessResponseBodyResult {
	s.ProcessCode = &v
	return s
}

type SaveProcessResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveProcessResponse) GoString() string {
	return s.String()
}

func (s *SaveProcessResponse) SetHeaders(v map[string]*string) *SaveProcessResponse {
	s.Headers = v
	return s
}

func (s *SaveProcessResponse) SetStatusCode(v int32) *SaveProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveProcessResponse) SetBody(v *SaveProcessResponseBody) *SaveProcessResponse {
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
	Approvers []*StartProcessInstanceRequestApprovers `json:"approvers,omitempty" xml:"approvers,omitempty" type:"Repeated"`
	// example:
	//
	// https://www.dingtalk.com/
	BizDetailPageUrl *string   `json:"bizDetailPageUrl,omitempty" xml:"bizDetailPageUrl,omitempty"`
	CcList           []*string `json:"ccList,omitempty" xml:"ccList,omitempty" type:"Repeated"`
	// example:
	//
	// START、FINISH、START_FINISH
	CcPosition *string `json:"ccPosition,omitempty" xml:"ccPosition,omitempty"`
	// example:
	//
	// 1
	DeptId *int64 `json:"deptId,omitempty" xml:"deptId,omitempty"`
	// This parameter is required.
	FormComponentValues []*StartProcessInstanceRequestFormComponentValues `json:"formComponentValues,omitempty" xml:"formComponentValues,omitempty" type:"Repeated"`
	// example:
	//
	// 41605932
	MicroappAgentId *int64 `json:"microappAgentId,omitempty" xml:"microappAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// manager432
	OriginatorUserId *string `json:"originatorUserId,omitempty" xml:"originatorUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROC-EF6YJL35P2-SCKICSB7P750S0YISYKV3-xxxx-1
	ProcessCode           *string                                             `json:"processCode,omitempty" xml:"processCode,omitempty"`
	TargetSelectActioners []*StartProcessInstanceRequestTargetSelectActioners `json:"targetSelectActioners,omitempty" xml:"targetSelectActioners,omitempty" type:"Repeated"`
}

func (s StartProcessInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartProcessInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceRequest) SetApprovers(v []*StartProcessInstanceRequestApprovers) *StartProcessInstanceRequest {
	s.Approvers = v
	return s
}

func (s *StartProcessInstanceRequest) SetBizDetailPageUrl(v string) *StartProcessInstanceRequest {
	s.BizDetailPageUrl = &v
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

func (s *StartProcessInstanceRequest) SetDeptId(v int64) *StartProcessInstanceRequest {
	s.DeptId = &v
	return s
}

func (s *StartProcessInstanceRequest) SetFormComponentValues(v []*StartProcessInstanceRequestFormComponentValues) *StartProcessInstanceRequest {
	s.FormComponentValues = v
	return s
}

func (s *StartProcessInstanceRequest) SetMicroappAgentId(v int64) *StartProcessInstanceRequest {
	s.MicroappAgentId = &v
	return s
}

func (s *StartProcessInstanceRequest) SetOriginatorUserId(v string) *StartProcessInstanceRequest {
	s.OriginatorUserId = &v
	return s
}

func (s *StartProcessInstanceRequest) SetProcessCode(v string) *StartProcessInstanceRequest {
	s.ProcessCode = &v
	return s
}

func (s *StartProcessInstanceRequest) SetTargetSelectActioners(v []*StartProcessInstanceRequestTargetSelectActioners) *StartProcessInstanceRequest {
	s.TargetSelectActioners = v
	return s
}

type StartProcessInstanceRequestApprovers struct {
	// example:
	//
	// 会签：AND；或签：OR；单人：NONE
	ActionType *string   `json:"actionType,omitempty" xml:"actionType,omitempty"`
	UserIds    []*string `json:"userIds,omitempty" xml:"userIds,omitempty" type:"Repeated"`
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

type StartProcessInstanceRequestFormComponentValues struct {
	// example:
	//
	// Phone
	BizAlias      *string                                                  `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	ComponentType *string                                                  `json:"componentType,omitempty" xml:"componentType,omitempty"`
	Details       []*StartProcessInstanceRequestFormComponentValuesDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
	// example:
	//
	// 总个数:1
	ExtValue *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	// example:
	//
	// PhoneField_IZI2LP8QF6O0
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PhoneField
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123xxxxxxxx
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s StartProcessInstanceRequestFormComponentValues) String() string {
	return tea.Prettify(s)
}

func (s StartProcessInstanceRequestFormComponentValues) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceRequestFormComponentValues) SetBizAlias(v string) *StartProcessInstanceRequestFormComponentValues {
	s.BizAlias = &v
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

func (s *StartProcessInstanceRequestFormComponentValues) SetExtValue(v string) *StartProcessInstanceRequestFormComponentValues {
	s.ExtValue = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValues) SetId(v string) *StartProcessInstanceRequestFormComponentValues {
	s.Id = &v
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

type StartProcessInstanceRequestFormComponentValuesDetails struct {
	// example:
	//
	// Phone
	BizAlias *string                                                         `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	Details  []*StartProcessInstanceRequestFormComponentValuesDetailsDetails `json:"details,omitempty" xml:"details,omitempty" type:"Repeated"`
	// example:
	//
	// 总个数:1
	ExtValue *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	// example:
	//
	// PhoneField_IZI2LP8QF6O0
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// PhoneField
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 123xxxxxxxx
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s StartProcessInstanceRequestFormComponentValuesDetails) String() string {
	return tea.Prettify(s)
}

func (s StartProcessInstanceRequestFormComponentValuesDetails) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceRequestFormComponentValuesDetails) SetBizAlias(v string) *StartProcessInstanceRequestFormComponentValuesDetails {
	s.BizAlias = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValuesDetails) SetDetails(v []*StartProcessInstanceRequestFormComponentValuesDetailsDetails) *StartProcessInstanceRequestFormComponentValuesDetails {
	s.Details = v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValuesDetails) SetExtValue(v string) *StartProcessInstanceRequestFormComponentValuesDetails {
	s.ExtValue = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValuesDetails) SetId(v string) *StartProcessInstanceRequestFormComponentValuesDetails {
	s.Id = &v
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

type StartProcessInstanceRequestFormComponentValuesDetailsDetails struct {
	// example:
	//
	// Phone
	BizAlias      *string `json:"bizAlias,omitempty" xml:"bizAlias,omitempty"`
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// example:
	//
	// 总个数:1
	ExtValue *string `json:"extValue,omitempty" xml:"extValue,omitempty"`
	// example:
	//
	// PhoneField_IZI2LP8QF6O0
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// PhoneField
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 123xxxxxxxx
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s StartProcessInstanceRequestFormComponentValuesDetailsDetails) String() string {
	return tea.Prettify(s)
}

func (s StartProcessInstanceRequestFormComponentValuesDetailsDetails) GoString() string {
	return s.String()
}

func (s *StartProcessInstanceRequestFormComponentValuesDetailsDetails) SetBizAlias(v string) *StartProcessInstanceRequestFormComponentValuesDetailsDetails {
	s.BizAlias = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValuesDetailsDetails) SetComponentType(v string) *StartProcessInstanceRequestFormComponentValuesDetailsDetails {
	s.ComponentType = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValuesDetailsDetails) SetExtValue(v string) *StartProcessInstanceRequestFormComponentValuesDetailsDetails {
	s.ExtValue = &v
	return s
}

func (s *StartProcessInstanceRequestFormComponentValuesDetailsDetails) SetId(v string) *StartProcessInstanceRequestFormComponentValuesDetailsDetails {
	s.Id = &v
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

type StartProcessInstanceRequestTargetSelectActioners struct {
	// example:
	//
	// manual_1918_5cd3_5e19_6a98
	ActionerKey     *string   `json:"actionerKey,omitempty" xml:"actionerKey,omitempty"`
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

type StartProcessInstanceResponseBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 91ef1076-c3ed-4a78-a7a5-fa29ef2d6252
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartProcessInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *StartProcessInstanceResponse) SetStatusCode(v int32) *StartProcessInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartProcessInstanceResponse) SetBody(v *StartProcessInstanceResponseBody) *StartProcessInstanceResponse {
	s.Body = v
	return s
}

type TerminateProcessInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TerminateProcessInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s TerminateProcessInstanceHeaders) GoString() string {
	return s.String()
}

func (s *TerminateProcessInstanceHeaders) SetCommonHeaders(v map[string]*string) *TerminateProcessInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TerminateProcessInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *TerminateProcessInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TerminateProcessInstanceRequest struct {
	IsSystem *bool `json:"isSystem,omitempty" xml:"isSystem,omitempty"`
	// example:
	//
	// 133743186427339452
	OperatingUserId *string `json:"operatingUserId,omitempty" xml:"operatingUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a171de6c-8bxxxx
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// example:
	//
	// 终止说明。
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s TerminateProcessInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s TerminateProcessInstanceRequest) GoString() string {
	return s.String()
}

func (s *TerminateProcessInstanceRequest) SetIsSystem(v bool) *TerminateProcessInstanceRequest {
	s.IsSystem = &v
	return s
}

func (s *TerminateProcessInstanceRequest) SetOperatingUserId(v string) *TerminateProcessInstanceRequest {
	s.OperatingUserId = &v
	return s
}

func (s *TerminateProcessInstanceRequest) SetProcessInstanceId(v string) *TerminateProcessInstanceRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *TerminateProcessInstanceRequest) SetRemark(v string) *TerminateProcessInstanceRequest {
	s.Remark = &v
	return s
}

type TerminateProcessInstanceResponseBody struct {
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TerminateProcessInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TerminateProcessInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *TerminateProcessInstanceResponseBody) SetResult(v bool) *TerminateProcessInstanceResponseBody {
	s.Result = &v
	return s
}

func (s *TerminateProcessInstanceResponseBody) SetSuccess(v bool) *TerminateProcessInstanceResponseBody {
	s.Success = &v
	return s
}

type TerminateProcessInstanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TerminateProcessInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TerminateProcessInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s TerminateProcessInstanceResponse) GoString() string {
	return s.String()
}

func (s *TerminateProcessInstanceResponse) SetHeaders(v map[string]*string) *TerminateProcessInstanceResponse {
	s.Headers = v
	return s
}

func (s *TerminateProcessInstanceResponse) SetStatusCode(v int32) *TerminateProcessInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *TerminateProcessInstanceResponse) SetBody(v *TerminateProcessInstanceResponseBody) *TerminateProcessInstanceResponse {
	s.Body = v
	return s
}

type TodoTasksHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s TodoTasksHeaders) String() string {
	return tea.Prettify(s)
}

func (s TodoTasksHeaders) GoString() string {
	return s.String()
}

func (s *TodoTasksHeaders) SetCommonHeaders(v map[string]*string) *TodoTasksHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TodoTasksHeaders) SetXAcsDingtalkAccessToken(v string) *TodoTasksHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type TodoTasksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// staffId123
	ActionerUserId *string `json:"actionerUserId,omitempty" xml:"actionerUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// manager123
	ManagerUserId *string `json:"managerUserId,omitempty" xml:"managerUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// This parameter is required.
	NextToken *int32 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s TodoTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s TodoTasksRequest) GoString() string {
	return s.String()
}

func (s *TodoTasksRequest) SetActionerUserId(v string) *TodoTasksRequest {
	s.ActionerUserId = &v
	return s
}

func (s *TodoTasksRequest) SetManagerUserId(v string) *TodoTasksRequest {
	s.ManagerUserId = &v
	return s
}

func (s *TodoTasksRequest) SetMaxResults(v int32) *TodoTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *TodoTasksRequest) SetNextToken(v int32) *TodoTasksRequest {
	s.NextToken = &v
	return s
}

type TodoTasksResponseBody struct {
	Result *TodoTasksResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s TodoTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TodoTasksResponseBody) GoString() string {
	return s.String()
}

func (s *TodoTasksResponseBody) SetResult(v *TodoTasksResponseBodyResult) *TodoTasksResponseBody {
	s.Result = v
	return s
}

type TodoTasksResponseBodyResult struct {
	HasMore *bool                              `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	List    []*TodoTasksResponseBodyResultList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s TodoTasksResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s TodoTasksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *TodoTasksResponseBodyResult) SetHasMore(v bool) *TodoTasksResponseBodyResult {
	s.HasMore = &v
	return s
}

func (s *TodoTasksResponseBodyResult) SetList(v []*TodoTasksResponseBodyResultList) *TodoTasksResponseBodyResult {
	s.List = v
	return s
}

type TodoTasksResponseBodyResultList struct {
	// example:
	//
	// RUNNING
	BusinessId  *string `json:"businessId,omitempty" xml:"businessId,omitempty"`
	CanRedirect *bool   `json:"canRedirect,omitempty" xml:"canRedirect,omitempty"`
	CreateTime  *int64  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// act_0001
	ProcessCode *string `json:"processCode,omitempty" xml:"processCode,omitempty"`
	// example:
	//
	// Siw2WNVZS4KiUt3tTmaNKg04*****809950
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// example:
	//
	// 1234567
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// manager001
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 2022-10-17T15:12Z
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s TodoTasksResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s TodoTasksResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *TodoTasksResponseBodyResultList) SetBusinessId(v string) *TodoTasksResponseBodyResultList {
	s.BusinessId = &v
	return s
}

func (s *TodoTasksResponseBodyResultList) SetCanRedirect(v bool) *TodoTasksResponseBodyResultList {
	s.CanRedirect = &v
	return s
}

func (s *TodoTasksResponseBodyResultList) SetCreateTime(v int64) *TodoTasksResponseBodyResultList {
	s.CreateTime = &v
	return s
}

func (s *TodoTasksResponseBodyResultList) SetProcessCode(v string) *TodoTasksResponseBodyResultList {
	s.ProcessCode = &v
	return s
}

func (s *TodoTasksResponseBodyResultList) SetProcessInstanceId(v string) *TodoTasksResponseBodyResultList {
	s.ProcessInstanceId = &v
	return s
}

func (s *TodoTasksResponseBodyResultList) SetTaskId(v int64) *TodoTasksResponseBodyResultList {
	s.TaskId = &v
	return s
}

func (s *TodoTasksResponseBodyResultList) SetTitle(v string) *TodoTasksResponseBodyResultList {
	s.Title = &v
	return s
}

func (s *TodoTasksResponseBodyResultList) SetUserId(v string) *TodoTasksResponseBodyResultList {
	s.UserId = &v
	return s
}

type TodoTasksResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TodoTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TodoTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s TodoTasksResponse) GoString() string {
	return s.String()
}

func (s *TodoTasksResponse) SetHeaders(v map[string]*string) *TodoTasksResponse {
	s.Headers = v
	return s
}

func (s *TodoTasksResponse) SetStatusCode(v int32) *TodoTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *TodoTasksResponse) SetBody(v *TodoTasksResponseBody) *TodoTasksResponse {
	s.Body = v
	return s
}

type UpdateIntegratedTaskHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateIntegratedTaskHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateIntegratedTaskHeaders) GoString() string {
	return s.String()
}

func (s *UpdateIntegratedTaskHeaders) SetCommonHeaders(v map[string]*string) *UpdateIntegratedTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateIntegratedTaskHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateIntegratedTaskHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateIntegratedTaskRequest struct {
	// if can be null:
	// false
	//
	// example:
	//
	// tPr_FB_mT_xxxxxxxxx2hQ05201655306463
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// This parameter is required.
	Tasks []*UpdateIntegratedTaskRequestTasks `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Repeated"`
}

func (s UpdateIntegratedTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIntegratedTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateIntegratedTaskRequest) SetProcessInstanceId(v string) *UpdateIntegratedTaskRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *UpdateIntegratedTaskRequest) SetTasks(v []*UpdateIntegratedTaskRequestTasks) *UpdateIntegratedTaskRequest {
	s.Tasks = v
	return s
}

type UpdateIntegratedTaskRequestTasks struct {
	// example:
	//
	// AGREE
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// COMPLETED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// true
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s UpdateIntegratedTaskRequestTasks) String() string {
	return tea.Prettify(s)
}

func (s UpdateIntegratedTaskRequestTasks) GoString() string {
	return s.String()
}

func (s *UpdateIntegratedTaskRequestTasks) SetResult(v string) *UpdateIntegratedTaskRequestTasks {
	s.Result = &v
	return s
}

func (s *UpdateIntegratedTaskRequestTasks) SetStatus(v string) *UpdateIntegratedTaskRequestTasks {
	s.Status = &v
	return s
}

func (s *UpdateIntegratedTaskRequestTasks) SetTaskId(v int64) *UpdateIntegratedTaskRequestTasks {
	s.TaskId = &v
	return s
}

type UpdateIntegratedTaskResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateIntegratedTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIntegratedTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIntegratedTaskResponseBody) SetSuccess(v bool) *UpdateIntegratedTaskResponseBody {
	s.Success = &v
	return s
}

type UpdateIntegratedTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIntegratedTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIntegratedTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIntegratedTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateIntegratedTaskResponse) SetHeaders(v map[string]*string) *UpdateIntegratedTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateIntegratedTaskResponse) SetStatusCode(v int32) *UpdateIntegratedTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIntegratedTaskResponse) SetBody(v *UpdateIntegratedTaskResponseBody) *UpdateIntegratedTaskResponse {
	s.Body = v
	return s
}

type UpdateProcessInstanceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsDingtalkAccessToken *string            `json:"x-acs-dingtalk-access-token,omitempty" xml:"x-acs-dingtalk-access-token,omitempty"`
}

func (s UpdateProcessInstanceHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateProcessInstanceHeaders) GoString() string {
	return s.String()
}

func (s *UpdateProcessInstanceHeaders) SetCommonHeaders(v map[string]*string) *UpdateProcessInstanceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateProcessInstanceHeaders) SetXAcsDingtalkAccessToken(v string) *UpdateProcessInstanceHeaders {
	s.XAcsDingtalkAccessToken = &v
	return s
}

type UpdateProcessInstanceRequest struct {
	Notifiers []*UpdateProcessInstanceRequestNotifiers `json:"notifiers,omitempty" xml:"notifiers,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// abc
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	// example:
	//
	// agree
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// COMPLETED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateProcessInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProcessInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateProcessInstanceRequest) SetNotifiers(v []*UpdateProcessInstanceRequestNotifiers) *UpdateProcessInstanceRequest {
	s.Notifiers = v
	return s
}

func (s *UpdateProcessInstanceRequest) SetProcessInstanceId(v string) *UpdateProcessInstanceRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *UpdateProcessInstanceRequest) SetResult(v string) *UpdateProcessInstanceRequest {
	s.Result = &v
	return s
}

func (s *UpdateProcessInstanceRequest) SetStatus(v string) *UpdateProcessInstanceRequest {
	s.Status = &v
	return s
}

type UpdateProcessInstanceRequestNotifiers struct {
	// This parameter is required.
	//
	// example:
	//
	// manager001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateProcessInstanceRequestNotifiers) String() string {
	return tea.Prettify(s)
}

func (s UpdateProcessInstanceRequestNotifiers) GoString() string {
	return s.String()
}

func (s *UpdateProcessInstanceRequestNotifiers) SetUserId(v string) *UpdateProcessInstanceRequestNotifiers {
	s.UserId = &v
	return s
}

type UpdateProcessInstanceResponseBody struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateProcessInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProcessInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProcessInstanceResponseBody) SetSuccess(v bool) *UpdateProcessInstanceResponseBody {
	s.Success = &v
	return s
}

type UpdateProcessInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProcessInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProcessInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProcessInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateProcessInstanceResponse) SetHeaders(v map[string]*string) *UpdateProcessInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateProcessInstanceResponse) SetStatusCode(v int32) *UpdateProcessInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProcessInstanceResponse) SetBody(v *UpdateProcessInstanceResponseBody) *UpdateProcessInstanceResponse {
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
// 授权下载审批钉盘文件
//
// @param request - AddApproveDentryAuthRequest
//
// @param headers - AddApproveDentryAuthHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddApproveDentryAuthResponse
func (client *Client) AddApproveDentryAuthWithOptions(request *AddApproveDentryAuthRequest, headers *AddApproveDentryAuthHeaders, runtime *util.RuntimeOptions) (_result *AddApproveDentryAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileInfos)) {
		body["fileInfos"] = request.FileInfos
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
		Action:      tea.String("AddApproveDentryAuth"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processInstances/spaces/files/authDownload"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddApproveDentryAuthResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 授权下载审批钉盘文件
//
// @param request - AddApproveDentryAuthRequest
//
// @return AddApproveDentryAuthResponse
func (client *Client) AddApproveDentryAuth(request *AddApproveDentryAuthRequest) (_result *AddApproveDentryAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddApproveDentryAuthHeaders{}
	_result = &AddApproveDentryAuthResponse{}
	_body, _err := client.AddApproveDentryAuthWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加审批评论
//
// @param request - AddProcessInstanceCommentRequest
//
// @param headers - AddProcessInstanceCommentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddProcessInstanceCommentResponse
func (client *Client) AddProcessInstanceCommentWithOptions(request *AddProcessInstanceCommentRequest, headers *AddProcessInstanceCommentHeaders, runtime *util.RuntimeOptions) (_result *AddProcessInstanceCommentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommentUserId)) {
		body["commentUserId"] = request.CommentUserId
	}

	if !tea.BoolValue(util.IsUnset(request.File)) {
		body["file"] = request.File
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		body["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		body["text"] = request.Text
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
		Action:      tea.String("AddProcessInstanceComment"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processInstances/comments"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &AddProcessInstanceCommentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加审批评论
//
// @param request - AddProcessInstanceCommentRequest
//
// @return AddProcessInstanceCommentResponse
func (client *Client) AddProcessInstanceComment(request *AddProcessInstanceCommentRequest) (_result *AddProcessInstanceCommentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddProcessInstanceCommentHeaders{}
	_result = &AddProcessInstanceCommentResponse{}
	_body, _err := client.AddProcessInstanceCommentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 归档审批实例(OA高级版专享)
//
// @param request - ArchiveProcessInstanceRequest
//
// @param headers - ArchiveProcessInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ArchiveProcessInstanceResponse
func (client *Client) ArchiveProcessInstanceWithOptions(request *ArchiveProcessInstanceRequest, headers *ArchiveProcessInstanceHeaders, runtime *util.RuntimeOptions) (_result *ArchiveProcessInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

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
		Action:      tea.String("ArchiveProcessInstance"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/processInstances/archive"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ArchiveProcessInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 归档审批实例(OA高级版专享)
//
// @param request - ArchiveProcessInstanceRequest
//
// @return ArchiveProcessInstanceResponse
func (client *Client) ArchiveProcessInstance(request *ArchiveProcessInstanceRequest) (_result *ArchiveProcessInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ArchiveProcessInstanceHeaders{}
	_result = &ArchiveProcessInstanceResponse{}
	_body, _err := client.ArchiveProcessInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量同意或拒绝审批任务
//
// @param request - BatchExecuteProcessInstancesRequest
//
// @param headers - BatchExecuteProcessInstancesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchExecuteProcessInstancesResponse
func (client *Client) BatchExecuteProcessInstancesWithOptions(request *BatchExecuteProcessInstancesRequest, headers *BatchExecuteProcessInstancesHeaders, runtime *util.RuntimeOptions) (_result *BatchExecuteProcessInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionerUserId)) {
		body["actionerUserId"] = request.ActionerUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Result)) {
		body["result"] = request.Result
	}

	if !tea.BoolValue(util.IsUnset(request.TaskInfoList)) {
		body["taskInfoList"] = request.TaskInfoList
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
		Action:      tea.String("BatchExecuteProcessInstances"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processInstances/batchExecute"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchExecuteProcessInstancesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量同意或拒绝审批任务
//
// @param request - BatchExecuteProcessInstancesRequest
//
// @return BatchExecuteProcessInstancesResponse
func (client *Client) BatchExecuteProcessInstances(request *BatchExecuteProcessInstancesRequest) (_result *BatchExecuteProcessInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchExecuteProcessInstancesHeaders{}
	_result = &BatchExecuteProcessInstancesResponse{}
	_body, _err := client.BatchExecuteProcessInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量流程审批任务转交
//
// @param request - BatchTasksRedirectRequest
//
// @param headers - BatchTasksRedirectHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchTasksRedirectResponse
func (client *Client) BatchTasksRedirectWithOptions(request *BatchTasksRedirectRequest, headers *BatchTasksRedirectHeaders, runtime *util.RuntimeOptions) (_result *BatchTasksRedirectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HandoverUserId)) {
		body["handoverUserId"] = request.HandoverUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerUserId)) {
		body["managerUserId"] = request.ManagerUserId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskIds)) {
		body["taskIds"] = request.TaskIds
	}

	if !tea.BoolValue(util.IsUnset(request.TransfereeUserId)) {
		body["transfereeUserId"] = request.TransfereeUserId
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
		Action:      tea.String("BatchTasksRedirect"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/tasks/batchRedirect"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchTasksRedirectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量流程审批任务转交
//
// @param request - BatchTasksRedirectRequest
//
// @return BatchTasksRedirectResponse
func (client *Client) BatchTasksRedirect(request *BatchTasksRedirectRequest) (_result *BatchTasksRedirectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchTasksRedirectHeaders{}
	_result = &BatchTasksRedirectResponse{}
	_body, _err := client.BatchTasksRedirectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量更新实例状态
//
// @param request - BatchUpdateProcessInstanceRequest
//
// @param headers - BatchUpdateProcessInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUpdateProcessInstanceResponse
func (client *Client) BatchUpdateProcessInstanceWithOptions(request *BatchUpdateProcessInstanceRequest, headers *BatchUpdateProcessInstanceHeaders, runtime *util.RuntimeOptions) (_result *BatchUpdateProcessInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UpdateProcessInstanceRequests)) {
		body["updateProcessInstanceRequests"] = request.UpdateProcessInstanceRequests
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
		Action:      tea.String("BatchUpdateProcessInstance"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processCentres/instances/batch"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchUpdateProcessInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量更新实例状态
//
// @param request - BatchUpdateProcessInstanceRequest
//
// @return BatchUpdateProcessInstanceResponse
func (client *Client) BatchUpdateProcessInstance(request *BatchUpdateProcessInstanceRequest) (_result *BatchUpdateProcessInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &BatchUpdateProcessInstanceHeaders{}
	_result = &BatchUpdateProcessInstanceResponse{}
	_body, _err := client.BatchUpdateProcessInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量取消流程中心待处理任务
//
// @param request - CancelIntegratedTaskRequest
//
// @param headers - CancelIntegratedTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelIntegratedTaskResponse
func (client *Client) CancelIntegratedTaskWithOptions(request *CancelIntegratedTaskRequest, headers *CancelIntegratedTaskHeaders, runtime *util.RuntimeOptions) (_result *CancelIntegratedTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActivityId)) {
		body["activityId"] = request.ActivityId
	}

	if !tea.BoolValue(util.IsUnset(request.ActivityIds)) {
		body["activityIds"] = request.ActivityIds
	}

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
		Action:      tea.String("CancelIntegratedTask"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processCentres/tasks/cancel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelIntegratedTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量取消流程中心待处理任务
//
// @param request - CancelIntegratedTaskRequest
//
// @return CancelIntegratedTaskResponse
func (client *Client) CancelIntegratedTask(request *CancelIntegratedTaskRequest) (_result *CancelIntegratedTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CancelIntegratedTaskHeaders{}
	_result = &CancelIntegratedTaskResponse{}
	_body, _err := client.CancelIntegratedTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 清理审批数据
//
// @param request - CleanProcessDataRequest
//
// @param headers - CleanProcessDataHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CleanProcessDataResponse
func (client *Client) CleanProcessDataWithOptions(request *CleanProcessDataRequest, headers *CleanProcessDataHeaders, runtime *util.RuntimeOptions) (_result *CleanProcessDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpId)) {
		body["corpId"] = request.CorpId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		body["processCode"] = request.ProcessCode
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
		Action:      tea.String("CleanProcessData"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processes/clean"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CleanProcessDataResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 清理审批数据
//
// @param request - CleanProcessDataRequest
//
// @return CleanProcessDataResponse
func (client *Client) CleanProcessData(request *CleanProcessDataRequest) (_result *CleanProcessDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CleanProcessDataHeaders{}
	_result = &CleanProcessDataResponse{}
	_body, _err := client.CleanProcessDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 复制审批流
//
// @param request - CopyProcessRequest
//
// @param headers - CopyProcessHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyProcessResponse
func (client *Client) CopyProcessWithOptions(request *CopyProcessRequest, headers *CopyProcessHeaders, runtime *util.RuntimeOptions) (_result *CopyProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CopyOptions)) {
		body["copyOptions"] = request.CopyOptions
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCorpId)) {
		body["sourceCorpId"] = request.SourceCorpId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceProcessVOList)) {
		body["sourceProcessVOList"] = request.SourceProcessVOList
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
		Action:      tea.String("CopyProcess"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processes/copy"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CopyProcessResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 复制审批流
//
// @param request - CopyProcessRequest
//
// @return CopyProcessResponse
func (client *Client) CopyProcess(request *CopyProcessRequest) (_result *CopyProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CopyProcessHeaders{}
	_result = &CopyProcessResponse{}
	_body, _err := client.CopyProcessWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建流程中心待处理任务
//
// @param request - CreateIntegratedTaskRequest
//
// @param headers - CreateIntegratedTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIntegratedTaskResponse
func (client *Client) CreateIntegratedTaskWithOptions(request *CreateIntegratedTaskRequest, headers *CreateIntegratedTaskHeaders, runtime *util.RuntimeOptions) (_result *CreateIntegratedTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActivityId)) {
		body["activityId"] = request.ActivityId
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureConfig)) {
		body["featureConfig"] = request.FeatureConfig
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		body["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Tasks)) {
		body["tasks"] = request.Tasks
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
		Action:      tea.String("CreateIntegratedTask"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processCentres/tasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIntegratedTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建流程中心待处理任务
//
// @param request - CreateIntegratedTaskRequest
//
// @return CreateIntegratedTaskResponse
func (client *Client) CreateIntegratedTask(request *CreateIntegratedTaskRequest) (_result *CreateIntegratedTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateIntegratedTaskHeaders{}
	_result = &CreateIntegratedTaskResponse{}
	_body, _err := client.CreateIntegratedTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除分组
//
// @param request - DeleteDirRequest
//
// @param headers - DeleteDirHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDirResponse
func (client *Client) DeleteDirWithOptions(request *DeleteDirRequest, headers *DeleteDirHeaders, runtime *util.RuntimeOptions) (_result *DeleteDirResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirId)) {
		query["dirId"] = request.DirId
	}

	if !tea.BoolValue(util.IsUnset(request.OperateUserId)) {
		query["operateUserId"] = request.OperateUserId
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
		Action:      tea.String("DeleteDir"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processCentres/directories"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDirResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除分组
//
// @param request - DeleteDirRequest
//
// @return DeleteDirResponse
func (client *Client) DeleteDir(request *DeleteDirRequest) (_result *DeleteDirResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteDirHeaders{}
	_result = &DeleteDirResponse{}
	_body, _err := client.DeleteDirWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除模板
//
// @param request - DeleteProcessRequest
//
// @param headers - DeleteProcessHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProcessResponse
func (client *Client) DeleteProcessWithOptions(request *DeleteProcessRequest, headers *DeleteProcessHeaders, runtime *util.RuntimeOptions) (_result *DeleteProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CleanRunningTask)) {
		query["cleanRunningTask"] = request.CleanRunningTask
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		query["processCode"] = request.ProcessCode
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
		Action:      tea.String("DeleteProcess"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processCentres/schemas"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProcessResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除模板
//
// @param request - DeleteProcessRequest
//
// @return DeleteProcessResponse
func (client *Client) DeleteProcess(request *DeleteProcessRequest) (_result *DeleteProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteProcessHeaders{}
	_result = &DeleteProcessResponse{}
	_body, _err := client.DeleteProcessWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 同意或拒绝审批任务
//
// @param request - ExecuteProcessInstanceRequest
//
// @param headers - ExecuteProcessInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteProcessInstanceResponse
func (client *Client) ExecuteProcessInstanceWithOptions(request *ExecuteProcessInstanceRequest, headers *ExecuteProcessInstanceHeaders, runtime *util.RuntimeOptions) (_result *ExecuteProcessInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionerUserId)) {
		body["actionerUserId"] = request.ActionerUserId
	}

	if !tea.BoolValue(util.IsUnset(request.File)) {
		body["file"] = request.File
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		body["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Result)) {
		body["result"] = request.Result
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
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
		Action:      tea.String("ExecuteProcessInstance"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processInstances/execute"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteProcessInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同意或拒绝审批任务
//
// @param request - ExecuteProcessInstanceRequest
//
// @return ExecuteProcessInstanceResponse
func (client *Client) ExecuteProcessInstance(request *ExecuteProcessInstanceRequest) (_result *ExecuteProcessInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ExecuteProcessInstanceHeaders{}
	_result = &ExecuteProcessInstanceResponse{}
	_body, _err := client.ExecuteProcessInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建或更新审批表单模板
//
// @param request - FormCreateRequest
//
// @param headers - FormCreateHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FormCreateResponse
func (client *Client) FormCreateWithOptions(request *FormCreateRequest, headers *FormCreateHeaders, runtime *util.RuntimeOptions) (_result *FormCreateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FormComponents)) {
		body["formComponents"] = request.FormComponents
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		body["processCode"] = request.ProcessCode
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateConfig)) {
		body["templateConfig"] = request.TemplateConfig
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
		Action:      tea.String("FormCreate"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/forms"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &FormCreateResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建或更新审批表单模板
//
// @param request - FormCreateRequest
//
// @return FormCreateResponse
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

// Summary:
//
// 获取审批钉盘空间信息
//
// @param request - GetAttachmentSpaceRequest
//
// @param headers - GetAttachmentSpaceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAttachmentSpaceResponse
func (client *Client) GetAttachmentSpaceWithOptions(request *GetAttachmentSpaceRequest, headers *GetAttachmentSpaceHeaders, runtime *util.RuntimeOptions) (_result *GetAttachmentSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		body["agentId"] = request.AgentId
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
		Action:      tea.String("GetAttachmentSpace"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processInstances/spaces/infos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAttachmentSpaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取审批钉盘空间信息
//
// @param request - GetAttachmentSpaceRequest
//
// @return GetAttachmentSpaceResponse
func (client *Client) GetAttachmentSpace(request *GetAttachmentSpaceRequest) (_result *GetAttachmentSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetAttachmentSpaceHeaders{}
	_result = &GetAttachmentSpaceResponse{}
	_body, _err := client.GetAttachmentSpaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询已设置为条件的表单组件
//
// @param request - GetConditionFormComponentRequest
//
// @param headers - GetConditionFormComponentHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConditionFormComponentResponse
func (client *Client) GetConditionFormComponentWithOptions(request *GetConditionFormComponentRequest, headers *GetConditionFormComponentHeaders, runtime *util.RuntimeOptions) (_result *GetConditionFormComponentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		query["agentId"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		query["processCode"] = request.ProcessCode
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
		Action:      tea.String("GetConditionFormComponent"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processes/conditions/components"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConditionFormComponentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询已设置为条件的表单组件
//
// @param request - GetConditionFormComponentRequest
//
// @return GetConditionFormComponentResponse
func (client *Client) GetConditionFormComponent(request *GetConditionFormComponentRequest) (_result *GetConditionFormComponentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetConditionFormComponentHeaders{}
	_result = &GetConditionFormComponentResponse{}
	_body, _err := client.GetConditionFormComponentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取CRM所有流程code
//
// @param headers - GetCrmProcCodesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCrmProcCodesResponse
func (client *Client) GetCrmProcCodesWithOptions(headers *GetCrmProcCodesHeaders, runtime *util.RuntimeOptions) (_result *GetCrmProcCodesResponse, _err error) {
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
		Action:      tea.String("GetCrmProcCodes"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/crm/processes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCrmProcCodesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取CRM所有流程code
//
// @return GetCrmProcCodesResponse
func (client *Client) GetCrmProcCodes() (_result *GetCrmProcCodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetCrmProcCodesHeaders{}
	_result = &GetCrmProcCodesResponse{}
	_body, _err := client.GetCrmProcCodesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取表单字段修改历史
//
// @param request - GetFieldModifiedHistoryRequest
//
// @param headers - GetFieldModifiedHistoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFieldModifiedHistoryResponse
func (client *Client) GetFieldModifiedHistoryWithOptions(request *GetFieldModifiedHistoryRequest, headers *GetFieldModifiedHistoryHeaders, runtime *util.RuntimeOptions) (_result *GetFieldModifiedHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FieldId)) {
		body["fieldId"] = request.FieldId
	}

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
		Action:      tea.String("GetFieldModifiedHistory"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processes/fields/modifiedRecords/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFieldModifiedHistoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取表单字段修改历史
//
// @param request - GetFieldModifiedHistoryRequest
//
// @return GetFieldModifiedHistoryResponse
func (client *Client) GetFieldModifiedHistory(request *GetFieldModifiedHistoryRequest) (_result *GetFieldModifiedHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetFieldModifiedHistoryHeaders{}
	_result = &GetFieldModifiedHistoryResponse{}
	_body, _err := client.GetFieldModifiedHistoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取手写签名的下载链接
//
// @param request - GetHandSignDownloadUrlRequest
//
// @param headers - GetHandSignDownloadUrlHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHandSignDownloadUrlResponse
func (client *Client) GetHandSignDownloadUrlWithOptions(request *GetHandSignDownloadUrlRequest, headers *GetHandSignDownloadUrlHeaders, runtime *util.RuntimeOptions) (_result *GetHandSignDownloadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HandSignToken)) {
		body["handSignToken"] = request.HandSignToken
	}

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
		Action:      tea.String("GetHandSignDownloadUrl"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/processInstances/handSigns/downloadUrls/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHandSignDownloadUrlResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取手写签名的下载链接
//
// @param request - GetHandSignDownloadUrlRequest
//
// @return GetHandSignDownloadUrlResponse
func (client *Client) GetHandSignDownloadUrl(request *GetHandSignDownloadUrlRequest) (_result *GetHandSignDownloadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetHandSignDownloadUrlHeaders{}
	_result = &GetHandSignDownloadUrlResponse{}
	_body, _err := client.GetHandSignDownloadUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取当前企业所有可管理的表单
//
// @param request - GetManageProcessByStaffIdRequest
//
// @param headers - GetManageProcessByStaffIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetManageProcessByStaffIdResponse
func (client *Client) GetManageProcessByStaffIdWithOptions(request *GetManageProcessByStaffIdRequest, headers *GetManageProcessByStaffIdHeaders, runtime *util.RuntimeOptions) (_result *GetManageProcessByStaffIdResponse, _err error) {
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
		Action:      tea.String("GetManageProcessByStaffId"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processes/managements/templates"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetManageProcessByStaffIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取当前企业所有可管理的表单
//
// @param request - GetManageProcessByStaffIdRequest
//
// @return GetManageProcessByStaffIdResponse
func (client *Client) GetManageProcessByStaffId(request *GetManageProcessByStaffIdRequest) (_result *GetManageProcessByStaffIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetManageProcessByStaffIdHeaders{}
	_result = &GetManageProcessByStaffIdResponse{}
	_body, _err := client.GetManageProcessByStaffIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取模板code
//
// @param request - GetProcessCodeByNameRequest
//
// @param headers - GetProcessCodeByNameHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProcessCodeByNameResponse
func (client *Client) GetProcessCodeByNameWithOptions(request *GetProcessCodeByNameRequest, headers *GetProcessCodeByNameHeaders, runtime *util.RuntimeOptions) (_result *GetProcessCodeByNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
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
		Action:      tea.String("GetProcessCodeByName"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processCentres/schemaNames/processCodes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProcessCodeByNameResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取模板code
//
// @param request - GetProcessCodeByNameRequest
//
// @return GetProcessCodeByNameResponse
func (client *Client) GetProcessCodeByName(request *GetProcessCodeByNameRequest) (_result *GetProcessCodeByNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetProcessCodeByNameHeaders{}
	_result = &GetProcessCodeByNameResponse{}
	_body, _err := client.GetProcessCodeByNameWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取流程配置
//
// @param request - GetProcessConfigRequest
//
// @param headers - GetProcessConfigHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProcessConfigResponse
func (client *Client) GetProcessConfigWithOptions(request *GetProcessConfigRequest, headers *GetProcessConfigHeaders, runtime *util.RuntimeOptions) (_result *GetProcessConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProcCode)) {
		query["procCode"] = request.ProcCode
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
		Action:      tea.String("GetProcessConfig"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/crm/processes/configurations"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProcessConfigResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取流程配置
//
// @param request - GetProcessConfigRequest
//
// @return GetProcessConfigResponse
func (client *Client) GetProcessConfig(request *GetProcessConfigRequest) (_result *GetProcessConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetProcessConfigHeaders{}
	_result = &GetProcessConfigResponse{}
	_body, _err := client.GetProcessConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取单个审批实例详情
//
// @param request - GetProcessInstanceRequest
//
// @param headers - GetProcessInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProcessInstanceResponse
func (client *Client) GetProcessInstanceWithOptions(request *GetProcessInstanceRequest, headers *GetProcessInstanceHeaders, runtime *util.RuntimeOptions) (_result *GetProcessInstanceResponse, _err error) {
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
		Action:      tea.String("GetProcessInstance"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processInstances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProcessInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取单个审批实例详情
//
// @param request - GetProcessInstanceRequest
//
// @return GetProcessInstanceResponse
func (client *Client) GetProcessInstance(request *GetProcessInstanceRequest) (_result *GetProcessInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetProcessInstanceHeaders{}
	_result = &GetProcessInstanceResponse{}
	_body, _err := client.GetProcessInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取审批单详情高级接口，可以返回审批流程中的手写签名密码消息
//
// @param request - GetProcessInstanceWithExtraRequest
//
// @param headers - GetProcessInstanceWithExtraHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProcessInstanceWithExtraResponse
func (client *Client) GetProcessInstanceWithExtraWithOptions(request *GetProcessInstanceWithExtraRequest, headers *GetProcessInstanceWithExtraHeaders, runtime *util.RuntimeOptions) (_result *GetProcessInstanceWithExtraResponse, _err error) {
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
		Action:      tea.String("GetProcessInstanceWithExtra"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/processInstances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProcessInstanceWithExtraResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取审批单详情高级接口，可以返回审批流程中的手写签名密码消息
//
// @param request - GetProcessInstanceWithExtraRequest
//
// @return GetProcessInstanceWithExtraResponse
func (client *Client) GetProcessInstanceWithExtra(request *GetProcessInstanceWithExtraRequest) (_result *GetProcessInstanceWithExtraResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetProcessInstanceWithExtraHeaders{}
	_result = &GetProcessInstanceWithExtraResponse{}
	_body, _err := client.GetProcessInstanceWithExtraWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据模版code列表批量查询模板最新表单和流程配置
//
// @param tmpReq - GetSchemaAndProcessconfigBatchllyRequest
//
// @param headers - GetSchemaAndProcessconfigBatchllyHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSchemaAndProcessconfigBatchllyResponse
func (client *Client) GetSchemaAndProcessconfigBatchllyWithOptions(tmpReq *GetSchemaAndProcessconfigBatchllyRequest, headers *GetSchemaAndProcessconfigBatchllyHeaders, runtime *util.RuntimeOptions) (_result *GetSchemaAndProcessconfigBatchllyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetSchemaAndProcessconfigBatchllyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ProcessCodes)) {
		request.ProcessCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProcessCodes, tea.String("processCodes"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProcessCodesShrink)) {
		query["processCodes"] = request.ProcessCodesShrink
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
		Action:      tea.String("GetSchemaAndProcessconfigBatchlly"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processes/templates/batchQuery"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSchemaAndProcessconfigBatchllyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据模版code列表批量查询模板最新表单和流程配置
//
// @param request - GetSchemaAndProcessconfigBatchllyRequest
//
// @return GetSchemaAndProcessconfigBatchllyResponse
func (client *Client) GetSchemaAndProcessconfigBatchlly(request *GetSchemaAndProcessconfigBatchllyRequest) (_result *GetSchemaAndProcessconfigBatchllyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSchemaAndProcessconfigBatchllyHeaders{}
	_result = &GetSchemaAndProcessconfigBatchllyResponse{}
	_body, _err := client.GetSchemaAndProcessconfigBatchllyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 授权预览审批附件
//
// @param request - GetSpaceWithDownloadAuthRequest
//
// @param headers - GetSpaceWithDownloadAuthHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSpaceWithDownloadAuthResponse
func (client *Client) GetSpaceWithDownloadAuthWithOptions(request *GetSpaceWithDownloadAuthRequest, headers *GetSpaceWithDownloadAuthHeaders, runtime *util.RuntimeOptions) (_result *GetSpaceWithDownloadAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		body["agentId"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["fileId"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.FileIdList)) {
		body["fileIdList"] = request.FileIdList
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		body["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.WithCommentAttatchment)) {
		body["withCommentAttatchment"] = request.WithCommentAttatchment
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
		Action:      tea.String("GetSpaceWithDownloadAuth"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processInstances/spaces/authPreview"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSpaceWithDownloadAuthResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 授权预览审批附件
//
// @param request - GetSpaceWithDownloadAuthRequest
//
// @return GetSpaceWithDownloadAuthResponse
func (client *Client) GetSpaceWithDownloadAuth(request *GetSpaceWithDownloadAuthRequest) (_result *GetSpaceWithDownloadAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetSpaceWithDownloadAuthHeaders{}
	_result = &GetSpaceWithDownloadAuthResponse{}
	_body, _err := client.GetSpaceWithDownloadAuthWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户待审批数量
//
// @param request - GetUserTodoTaskSumRequest
//
// @param headers - GetUserTodoTaskSumHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserTodoTaskSumResponse
func (client *Client) GetUserTodoTaskSumWithOptions(request *GetUserTodoTaskSumRequest, headers *GetUserTodoTaskSumHeaders, runtime *util.RuntimeOptions) (_result *GetUserTodoTaskSumResponse, _err error) {
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
		Action:      tea.String("GetUserTodoTaskSum"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processes/todoTasks/numbers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserTodoTaskSumResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户待审批数量
//
// @param request - GetUserTodoTaskSumRequest
//
// @return GetUserTodoTaskSumResponse
func (client *Client) GetUserTodoTaskSum(request *GetUserTodoTaskSumRequest) (_result *GetUserTodoTaskSumResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUserTodoTaskSumHeaders{}
	_result = &GetUserTodoTaskSumResponse{}
	_body, _err := client.GetUserTodoTaskSumWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 授权用户钉盘空间权限
//
// @param request - GrantCspaceAuthorizationRequest
//
// @param headers - GrantCspaceAuthorizationHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantCspaceAuthorizationResponse
func (client *Client) GrantCspaceAuthorizationWithOptions(request *GrantCspaceAuthorizationRequest, headers *GrantCspaceAuthorizationHeaders, runtime *util.RuntimeOptions) (_result *GrantCspaceAuthorizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DurationSeconds)) {
		body["durationSeconds"] = request.DurationSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.SpaceId)) {
		body["spaceId"] = request.SpaceId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
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
		Action:      tea.String("GrantCspaceAuthorization"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/spaces/authorize"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("none"),
	}
	_result = &GrantCspaceAuthorizationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 授权用户钉盘空间权限
//
// @param request - GrantCspaceAuthorizationRequest
//
// @return GrantCspaceAuthorizationResponse
func (client *Client) GrantCspaceAuthorization(request *GrantCspaceAuthorizationRequest) (_result *GrantCspaceAuthorizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GrantCspaceAuthorizationHeaders{}
	_result = &GrantCspaceAuthorizationResponse{}
	_body, _err := client.GrantCspaceAuthorizationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 下载审批附件
//
// @param request - GrantProcessInstanceForDownloadFileRequest
//
// @param headers - GrantProcessInstanceForDownloadFileHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantProcessInstanceForDownloadFileResponse
func (client *Client) GrantProcessInstanceForDownloadFileWithOptions(request *GrantProcessInstanceForDownloadFileRequest, headers *GrantProcessInstanceForDownloadFileHeaders, runtime *util.RuntimeOptions) (_result *GrantProcessInstanceForDownloadFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["fileId"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		body["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.WithCommentAttatchment)) {
		body["withCommentAttatchment"] = request.WithCommentAttatchment
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
		Action:      tea.String("GrantProcessInstanceForDownloadFile"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processInstances/spaces/files/urls/download"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &GrantProcessInstanceForDownloadFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 下载审批附件
//
// @param request - GrantProcessInstanceForDownloadFileRequest
//
// @return GrantProcessInstanceForDownloadFileResponse
func (client *Client) GrantProcessInstanceForDownloadFile(request *GrantProcessInstanceForDownloadFileRequest) (_result *GrantProcessInstanceForDownloadFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GrantProcessInstanceForDownloadFileHeaders{}
	_result = &GrantProcessInstanceForDownloadFileResponse{}
	_body, _err := client.GrantProcessInstanceForDownloadFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建或更新分组
//
// @param request - InsertOrUpdateDirRequest
//
// @param headers - InsertOrUpdateDirHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertOrUpdateDirResponse
func (client *Client) InsertOrUpdateDirWithOptions(request *InsertOrUpdateDirRequest, headers *InsertOrUpdateDirHeaders, runtime *util.RuntimeOptions) (_result *InsertOrUpdateDirResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizGroup)) {
		body["bizGroup"] = request.BizGroup
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Name18n)) {
		body["name18n"] = request.Name18n
	}

	if !tea.BoolValue(util.IsUnset(request.OperateUserId)) {
		body["operateUserId"] = request.OperateUserId
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
		Action:      tea.String("InsertOrUpdateDir"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processCentres/directories"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &InsertOrUpdateDirResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建或更新分组
//
// @param request - InsertOrUpdateDirRequest
//
// @return InsertOrUpdateDirResponse
func (client *Client) InsertOrUpdateDir(request *InsertOrUpdateDirRequest) (_result *InsertOrUpdateDirResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InsertOrUpdateDirHeaders{}
	_result = &InsertOrUpdateDirResponse{}
	_body, _err := client.InsertOrUpdateDirWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 应用安装
//
// @param request - InstallAppRequest
//
// @param headers - InstallAppHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallAppResponse
func (client *Client) InstallAppWithOptions(request *InstallAppRequest, headers *InstallAppHeaders, runtime *util.RuntimeOptions) (_result *InstallAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizGroup)) {
		body["bizGroup"] = request.BizGroup
	}

	if !tea.BoolValue(util.IsUnset(request.InstallOption)) {
		body["installOption"] = request.InstallOption
	}

	if !tea.BoolValue(util.IsUnset(request.SourceDirName)) {
		body["sourceDirName"] = request.SourceDirName
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
		Action:      tea.String("InstallApp"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processes/apps/install"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallAppResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 应用安装
//
// @param request - InstallAppRequest
//
// @return InstallAppResponse
func (client *Client) InstallApp(request *InstallAppRequest) (_result *InstallAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InstallAppHeaders{}
	_result = &InstallAppResponse{}
	_body, _err := client.InstallAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取审批实例ID列表
//
// @param request - ListProcessInstanceIdsRequest
//
// @param headers - ListProcessInstanceIdsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProcessInstanceIdsResponse
func (client *Client) ListProcessInstanceIdsWithOptions(request *ListProcessInstanceIdsRequest, headers *ListProcessInstanceIdsHeaders, runtime *util.RuntimeOptions) (_result *ListProcessInstanceIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		body["processCode"] = request.ProcessCode
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Statuses)) {
		body["statuses"] = request.Statuses
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
		Action:      tea.String("ListProcessInstanceIds"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processes/instanceIds/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProcessInstanceIdsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取审批实例ID列表
//
// @param request - ListProcessInstanceIdsRequest
//
// @return ListProcessInstanceIdsResponse
func (client *Client) ListProcessInstanceIds(request *ListProcessInstanceIdsRequest) (_result *ListProcessInstanceIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListProcessInstanceIdsHeaders{}
	_result = &ListProcessInstanceIdsResponse{}
	_body, _err := client.ListProcessInstanceIdsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户待办事项
//
// @param request - ListTodoWorkRecordsRequest
//
// @param headers - ListTodoWorkRecordsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTodoWorkRecordsResponse
func (client *Client) ListTodoWorkRecordsWithOptions(request *ListTodoWorkRecordsRequest, headers *ListTodoWorkRecordsHeaders, runtime *util.RuntimeOptions) (_result *ListTodoWorkRecordsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
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
		Action:      tea.String("ListTodoWorkRecords"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/workRecords/todoTasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTodoWorkRecordsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户待办事项
//
// @param request - ListTodoWorkRecordsRequest
//
// @return ListTodoWorkRecordsResponse
func (client *Client) ListTodoWorkRecords(request *ListTodoWorkRecordsRequest) (_result *ListTodoWorkRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListTodoWorkRecordsHeaders{}
	_result = &ListTodoWorkRecordsResponse{}
	_body, _err := client.ListTodoWorkRecordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定用户可见的审批表单列表
//
// @param request - ListUserVisibleBpmsProcessesRequest
//
// @param headers - ListUserVisibleBpmsProcessesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserVisibleBpmsProcessesResponse
func (client *Client) ListUserVisibleBpmsProcessesWithOptions(request *ListUserVisibleBpmsProcessesRequest, headers *ListUserVisibleBpmsProcessesHeaders, runtime *util.RuntimeOptions) (_result *ListUserVisibleBpmsProcessesResponse, _err error) {
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
		Action:      tea.String("ListUserVisibleBpmsProcesses"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processes/userVisibilities/templates"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserVisibleBpmsProcessesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定用户可见的审批表单列表
//
// @param request - ListUserVisibleBpmsProcessesRequest
//
// @return ListUserVisibleBpmsProcessesResponse
func (client *Client) ListUserVisibleBpmsProcesses(request *ListUserVisibleBpmsProcessesRequest) (_result *ListUserVisibleBpmsProcessesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &ListUserVisibleBpmsProcessesHeaders{}
	_result = &ListUserVisibleBpmsProcessesResponse{}
	_body, _err := client.ListUserVisibleBpmsProcessesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询实例数据
//
// @param request - PagesExportInstancesRequest
//
// @param headers - PagesExportInstancesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PagesExportInstancesResponse
func (client *Client) PagesExportInstancesWithOptions(request *PagesExportInstancesRequest, headers *PagesExportInstancesHeaders, runtime *util.RuntimeOptions) (_result *PagesExportInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTimeInMills)) {
		query["endTimeInMills"] = request.EndTimeInMills
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["orderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		query["processCode"] = request.ProcessCode
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimeInMills)) {
		query["startTimeInMills"] = request.StartTimeInMills
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
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
		Action:      tea.String("PagesExportInstances"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/instances/datas"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PagesExportInstancesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询实例数据
//
// @param request - PagesExportInstancesRequest
//
// @return PagesExportInstancesResponse
func (client *Client) PagesExportInstances(request *PagesExportInstancesRequest) (_result *PagesExportInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PagesExportInstancesHeaders{}
	_result = &PagesExportInstancesResponse{}
	_body, _err := client.PagesExportInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 授权下载审批钉盘文件(OA高级版专享)
//
// @param request - PremiumAddApproveDentryAuthRequest
//
// @param headers - PremiumAddApproveDentryAuthHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumAddApproveDentryAuthResponse
func (client *Client) PremiumAddApproveDentryAuthWithOptions(request *PremiumAddApproveDentryAuthRequest, headers *PremiumAddApproveDentryAuthHeaders, runtime *util.RuntimeOptions) (_result *PremiumAddApproveDentryAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileInfos)) {
		body["fileInfos"] = request.FileInfos
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
		Action:      tea.String("PremiumAddApproveDentryAuth"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/processInstances/spaces/files/authDownload"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumAddApproveDentryAuthResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 授权下载审批钉盘文件(OA高级版专享)
//
// @param request - PremiumAddApproveDentryAuthRequest
//
// @return PremiumAddApproveDentryAuthResponse
func (client *Client) PremiumAddApproveDentryAuth(request *PremiumAddApproveDentryAuthRequest) (_result *PremiumAddApproveDentryAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumAddApproveDentryAuthHeaders{}
	_result = &PremiumAddApproveDentryAuthResponse{}
	_body, _err := client.PremiumAddApproveDentryAuthWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 加签OA审批任务(OA高级版专享)
//
// @param request - PremiumAppendTaskRequest
//
// @param headers - PremiumAppendTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumAppendTaskResponse
func (client *Client) PremiumAppendTaskWithOptions(request *PremiumAppendTaskRequest, headers *PremiumAppendTaskHeaders, runtime *util.RuntimeOptions) (_result *PremiumAppendTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActivateType)) {
		body["activateType"] = request.ActivateType
	}

	if !tea.BoolValue(util.IsUnset(request.AgreeAll)) {
		body["agreeAll"] = request.AgreeAll
	}

	if !tea.BoolValue(util.IsUnset(request.AppenderUserIds)) {
		body["appenderUserIds"] = request.AppenderUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.OperateUserId)) {
		body["operateUserId"] = request.OperateUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		body["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
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
		Action:      tea.String("PremiumAppendTask"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/tasks/append"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumAppendTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 加签OA审批任务(OA高级版专享)
//
// @param request - PremiumAppendTaskRequest
//
// @return PremiumAppendTaskResponse
func (client *Client) PremiumAppendTask(request *PremiumAppendTaskRequest) (_result *PremiumAppendTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumAppendTaskHeaders{}
	_result = &PremiumAppendTaskResponse{}
	_body, _err := client.PremiumAppendTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量同意或拒绝审批任务(OA高级版专享接口)
//
// @param request - PremiumBatchExecuteProcessInstancesRequest
//
// @param headers - PremiumBatchExecuteProcessInstancesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumBatchExecuteProcessInstancesResponse
func (client *Client) PremiumBatchExecuteProcessInstancesWithOptions(request *PremiumBatchExecuteProcessInstancesRequest, headers *PremiumBatchExecuteProcessInstancesHeaders, runtime *util.RuntimeOptions) (_result *PremiumBatchExecuteProcessInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionerUserId)) {
		body["actionerUserId"] = request.ActionerUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Result)) {
		body["result"] = request.Result
	}

	if !tea.BoolValue(util.IsUnset(request.TaskInfoList)) {
		body["taskInfoList"] = request.TaskInfoList
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
		Action:      tea.String("PremiumBatchExecuteProcessInstances"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/processInstances/batchExecute"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumBatchExecuteProcessInstancesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量同意或拒绝审批任务(OA高级版专享接口)
//
// @param request - PremiumBatchExecuteProcessInstancesRequest
//
// @return PremiumBatchExecuteProcessInstancesResponse
func (client *Client) PremiumBatchExecuteProcessInstances(request *PremiumBatchExecuteProcessInstancesRequest) (_result *PremiumBatchExecuteProcessInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumBatchExecuteProcessInstancesHeaders{}
	_result = &PremiumBatchExecuteProcessInstancesResponse{}
	_body, _err := client.PremiumBatchExecuteProcessInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除业务分组(高级版专享接口)
//
// @param request - PremiumDelDirRequest
//
// @param headers - PremiumDelDirHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumDelDirResponse
func (client *Client) PremiumDelDirWithOptions(request *PremiumDelDirRequest, headers *PremiumDelDirHeaders, runtime *util.RuntimeOptions) (_result *PremiumDelDirResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DirId)) {
		query["dirId"] = request.DirId
	}

	if !tea.BoolValue(util.IsUnset(request.OperateUserId)) {
		query["operateUserId"] = request.OperateUserId
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
		Action:      tea.String("PremiumDelDir"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/processCentres/directories"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumDelDirResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除业务分组(高级版专享接口)
//
// @param request - PremiumDelDirRequest
//
// @return PremiumDelDirResponse
func (client *Client) PremiumDelDir(request *PremiumDelDirRequest) (_result *PremiumDelDirResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumDelDirHeaders{}
	_result = &PremiumDelDirResponse{}
	_body, _err := client.PremiumDelDirWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除数据表单实例(OA高级版专享)
//
// @param request - PremiumDeleteFormInstanceRequest
//
// @param headers - PremiumDeleteFormInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumDeleteFormInstanceResponse
func (client *Client) PremiumDeleteFormInstanceWithOptions(request *PremiumDeleteFormInstanceRequest, headers *PremiumDeleteFormInstanceHeaders, runtime *util.RuntimeOptions) (_result *PremiumDeleteFormInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FormInstanceIds)) {
		body["formInstanceIds"] = request.FormInstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		body["processCode"] = request.ProcessCode
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
		Action:      tea.String("PremiumDeleteFormInstance"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/dataForms/formInstances/remove"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumDeleteFormInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除数据表单实例(OA高级版专享)
//
// @param request - PremiumDeleteFormInstanceRequest
//
// @return PremiumDeleteFormInstanceResponse
func (client *Client) PremiumDeleteFormInstance(request *PremiumDeleteFormInstanceRequest) (_result *PremiumDeleteFormInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumDeleteFormInstanceHeaders{}
	_result = &PremiumDeleteFormInstanceResponse{}
	_body, _err := client.PremiumDeleteFormInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取审批钉盘空间信息(OA高级版专享)
//
// @param request - PremiumGetAttachmentSpaceRequest
//
// @param headers - PremiumGetAttachmentSpaceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumGetAttachmentSpaceResponse
func (client *Client) PremiumGetAttachmentSpaceWithOptions(request *PremiumGetAttachmentSpaceRequest, headers *PremiumGetAttachmentSpaceHeaders, runtime *util.RuntimeOptions) (_result *PremiumGetAttachmentSpaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		body["agentId"] = request.AgentId
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
		Action:      tea.String("PremiumGetAttachmentSpace"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/processInstances/spaces/infos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumGetAttachmentSpaceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取审批钉盘空间信息(OA高级版专享)
//
// @param request - PremiumGetAttachmentSpaceRequest
//
// @return PremiumGetAttachmentSpaceResponse
func (client *Client) PremiumGetAttachmentSpace(request *PremiumGetAttachmentSpaceRequest) (_result *PremiumGetAttachmentSpaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumGetAttachmentSpaceHeaders{}
	_result = &PremiumGetAttachmentSpaceResponse{}
	_body, _err := client.PremiumGetAttachmentSpaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询审批中心已处理任务列表(OA高级版专享接口)
//
// @param request - PremiumGetDoneTasksRequest
//
// @param headers - PremiumGetDoneTasksHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumGetDoneTasksResponse
func (client *Client) PremiumGetDoneTasksWithOptions(request *PremiumGetDoneTasksRequest, headers *PremiumGetDoneTasksHeaders, runtime *util.RuntimeOptions) (_result *PremiumGetDoneTasksResponse, _err error) {
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
		Action:      tea.String("PremiumGetDoneTasks"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/processCentres/doneTasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumGetDoneTasksResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询审批中心已处理任务列表(OA高级版专享接口)
//
// @param request - PremiumGetDoneTasksRequest
//
// @return PremiumGetDoneTasksResponse
func (client *Client) PremiumGetDoneTasks(request *PremiumGetDoneTasksRequest) (_result *PremiumGetDoneTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumGetDoneTasksHeaders{}
	_result = &PremiumGetDoneTasksResponse{}
	_body, _err := client.PremiumGetDoneTasksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取字段修改历史(高级版专享接口)
//
// @param request - PremiumGetFieldModifiedHistoryRequest
//
// @param headers - PremiumGetFieldModifiedHistoryHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumGetFieldModifiedHistoryResponse
func (client *Client) PremiumGetFieldModifiedHistoryWithOptions(request *PremiumGetFieldModifiedHistoryRequest, headers *PremiumGetFieldModifiedHistoryHeaders, runtime *util.RuntimeOptions) (_result *PremiumGetFieldModifiedHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FieldId)) {
		body["fieldId"] = request.FieldId
	}

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
		Action:      tea.String("PremiumGetFieldModifiedHistory"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/processes/fields/modifiedRecords/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumGetFieldModifiedHistoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取字段修改历史(高级版专享接口)
//
// @param request - PremiumGetFieldModifiedHistoryRequest
//
// @return PremiumGetFieldModifiedHistoryResponse
func (client *Client) PremiumGetFieldModifiedHistory(request *PremiumGetFieldModifiedHistoryRequest) (_result *PremiumGetFieldModifiedHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumGetFieldModifiedHistoryHeaders{}
	_result = &PremiumGetFieldModifiedHistoryResponse{}
	_body, _err := client.PremiumGetFieldModifiedHistoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取单个数据表单实例详情(OA高级版专享)
//
// @param request - PremiumGetFormInstanceRequest
//
// @param headers - PremiumGetFormInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumGetFormInstanceResponse
func (client *Client) PremiumGetFormInstanceWithOptions(request *PremiumGetFormInstanceRequest, headers *PremiumGetFormInstanceHeaders, runtime *util.RuntimeOptions) (_result *PremiumGetFormInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUuid)) {
		query["appUuid"] = request.AppUuid
	}

	if !tea.BoolValue(util.IsUnset(request.FormCode)) {
		query["formCode"] = request.FormCode
	}

	if !tea.BoolValue(util.IsUnset(request.FormInstanceId)) {
		query["formInstanceId"] = request.FormInstanceId
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
		Action:      tea.String("PremiumGetFormInstance"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/dataForms/formInstances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumGetFormInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取单个数据表单实例详情(OA高级版专享)
//
// @param request - PremiumGetFormInstanceRequest
//
// @return PremiumGetFormInstanceResponse
func (client *Client) PremiumGetFormInstance(request *PremiumGetFormInstanceRequest) (_result *PremiumGetFormInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumGetFormInstanceHeaders{}
	_result = &PremiumGetFormInstanceResponse{}
	_body, _err := client.PremiumGetFormInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据formCode分页获取数据表单实例(OA高级版专享)
//
// @param request - PremiumGetFormInstancesRequest
//
// @param headers - PremiumGetFormInstancesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumGetFormInstancesResponse
func (client *Client) PremiumGetFormInstancesWithOptions(request *PremiumGetFormInstancesRequest, headers *PremiumGetFormInstancesHeaders, runtime *util.RuntimeOptions) (_result *PremiumGetFormInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUuid)) {
		query["appUuid"] = request.AppUuid
	}

	if !tea.BoolValue(util.IsUnset(request.FormCode)) {
		query["formCode"] = request.FormCode
	}

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
		Action:      tea.String("PremiumGetFormInstances"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/dataForms/formInstances/pages"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumGetFormInstancesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据formCode分页获取数据表单实例(OA高级版专享)
//
// @param request - PremiumGetFormInstancesRequest
//
// @return PremiumGetFormInstancesResponse
func (client *Client) PremiumGetFormInstances(request *PremiumGetFormInstancesRequest) (_result *PremiumGetFormInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumGetFormInstancesHeaders{}
	_result = &PremiumGetFormInstancesResponse{}
	_body, _err := client.PremiumGetFormInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过formCode获取数据表单schema(OA高级版专享)
//
// @param request - PremiumGetFormSchemaRequest
//
// @param headers - PremiumGetFormSchemaHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumGetFormSchemaResponse
func (client *Client) PremiumGetFormSchemaWithOptions(request *PremiumGetFormSchemaRequest, headers *PremiumGetFormSchemaHeaders, runtime *util.RuntimeOptions) (_result *PremiumGetFormSchemaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUuid)) {
		query["appUuid"] = request.AppUuid
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		query["processCode"] = request.ProcessCode
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
		Action:      tea.String("PremiumGetFormSchema"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/dataForms/schema/formCodes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumGetFormSchemaResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过formCode获取数据表单schema(OA高级版专享)
//
// @param request - PremiumGetFormSchemaRequest
//
// @return PremiumGetFormSchemaResponse
func (client *Client) PremiumGetFormSchema(request *PremiumGetFormSchemaRequest) (_result *PremiumGetFormSchemaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumGetFormSchemaHeaders{}
	_result = &PremiumGetFormSchemaResponse{}
	_body, _err := client.PremiumGetFormSchemaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取流程表单字段操作权限(高级版专享接口)
//
// @param request - PremiumGetInstFieldSettingRequest
//
// @param headers - PremiumGetInstFieldSettingHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumGetInstFieldSettingResponse
func (client *Client) PremiumGetInstFieldSettingWithOptions(request *PremiumGetInstFieldSettingRequest, headers *PremiumGetInstFieldSettingHeaders, runtime *util.RuntimeOptions) (_result *PremiumGetInstFieldSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PremiumGetInstFieldSetting"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/processes/fields/settings/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumGetInstFieldSettingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取流程表单字段操作权限(高级版专享接口)
//
// @param request - PremiumGetInstFieldSettingRequest
//
// @return PremiumGetInstFieldSettingResponse
func (client *Client) PremiumGetInstFieldSetting(request *PremiumGetInstFieldSettingRequest) (_result *PremiumGetInstFieldSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumGetInstFieldSettingHeaders{}
	_result = &PremiumGetInstFieldSettingResponse{}
	_body, _err := client.PremiumGetInstFieldSettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询审批中心我收到的实例列表(OA高级版专享接口)
//
// @param request - PremiumGetNoticedInstancesRequest
//
// @param headers - PremiumGetNoticedInstancesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumGetNoticedInstancesResponse
func (client *Client) PremiumGetNoticedInstancesWithOptions(request *PremiumGetNoticedInstancesRequest, headers *PremiumGetNoticedInstancesHeaders, runtime *util.RuntimeOptions) (_result *PremiumGetNoticedInstancesResponse, _err error) {
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
		Action:      tea.String("PremiumGetNoticedInstances"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/processCentres/noticedInstances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumGetNoticedInstancesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询审批中心我收到的实例列表(OA高级版专享接口)
//
// @param request - PremiumGetNoticedInstancesRequest
//
// @return PremiumGetNoticedInstancesResponse
func (client *Client) PremiumGetNoticedInstances(request *PremiumGetNoticedInstancesRequest) (_result *PremiumGetNoticedInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumGetNoticedInstancesHeaders{}
	_result = &PremiumGetNoticedInstancesResponse{}
	_body, _err := client.PremiumGetNoticedInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据processCode分页获取审批流程数据(高级版专享接口)
//
// @param request - PremiumGetProcessInstancesRequest
//
// @param headers - PremiumGetProcessInstancesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumGetProcessInstancesResponse
func (client *Client) PremiumGetProcessInstancesWithOptions(request *PremiumGetProcessInstancesRequest, headers *PremiumGetProcessInstancesHeaders, runtime *util.RuntimeOptions) (_result *PremiumGetProcessInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUuid)) {
		query["appUuid"] = request.AppUuid
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimeInMills)) {
		query["endTimeInMills"] = request.EndTimeInMills
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		query["processCode"] = request.ProcessCode
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimeInMills)) {
		query["startTimeInMills"] = request.StartTimeInMills
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
		Action:      tea.String("PremiumGetProcessInstances"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/processes/pages/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumGetProcessInstancesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据processCode分页获取审批流程数据(高级版专享接口)
//
// @param request - PremiumGetProcessInstancesRequest
//
// @return PremiumGetProcessInstancesResponse
func (client *Client) PremiumGetProcessInstances(request *PremiumGetProcessInstancesRequest) (_result *PremiumGetProcessInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumGetProcessInstancesHeaders{}
	_result = &PremiumGetProcessInstancesResponse{}
	_body, _err := client.PremiumGetProcessInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 授权预览审批附件(OA高级版专享)
//
// @param request - PremiumGetSpaceWithDownloadAuthRequest
//
// @param headers - PremiumGetSpaceWithDownloadAuthHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumGetSpaceWithDownloadAuthResponse
func (client *Client) PremiumGetSpaceWithDownloadAuthWithOptions(request *PremiumGetSpaceWithDownloadAuthRequest, headers *PremiumGetSpaceWithDownloadAuthHeaders, runtime *util.RuntimeOptions) (_result *PremiumGetSpaceWithDownloadAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		body["agentId"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["fileId"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.FileIdList)) {
		body["fileIdList"] = request.FileIdList
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		body["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.WithCommentAttatchment)) {
		body["withCommentAttatchment"] = request.WithCommentAttatchment
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
		Action:      tea.String("PremiumGetSpaceWithDownloadAuth"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/processInstances/spaces/authPreview"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumGetSpaceWithDownloadAuthResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 授权预览审批附件(OA高级版专享)
//
// @param request - PremiumGetSpaceWithDownloadAuthRequest
//
// @return PremiumGetSpaceWithDownloadAuthResponse
func (client *Client) PremiumGetSpaceWithDownloadAuth(request *PremiumGetSpaceWithDownloadAuthRequest) (_result *PremiumGetSpaceWithDownloadAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumGetSpaceWithDownloadAuthHeaders{}
	_result = &PremiumGetSpaceWithDownloadAuthResponse{}
	_body, _err := client.PremiumGetSpaceWithDownloadAuthWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询审批中心已发起实例列表(OA高级版专享接口)
//
// @param request - PremiumGetSubmittedInstancesRequest
//
// @param headers - PremiumGetSubmittedInstancesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumGetSubmittedInstancesResponse
func (client *Client) PremiumGetSubmittedInstancesWithOptions(request *PremiumGetSubmittedInstancesRequest, headers *PremiumGetSubmittedInstancesHeaders, runtime *util.RuntimeOptions) (_result *PremiumGetSubmittedInstancesResponse, _err error) {
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
		Action:      tea.String("PremiumGetSubmittedInstances"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/processCentres/submittedInstances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumGetSubmittedInstancesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询审批中心已发起实例列表(OA高级版专享接口)
//
// @param request - PremiumGetSubmittedInstancesRequest
//
// @return PremiumGetSubmittedInstancesResponse
func (client *Client) PremiumGetSubmittedInstances(request *PremiumGetSubmittedInstancesRequest) (_result *PremiumGetSubmittedInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumGetSubmittedInstancesHeaders{}
	_result = &PremiumGetSubmittedInstancesResponse{}
	_body, _err := client.PremiumGetSubmittedInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询审批中心待处理任务列表(OA高级版专享接口)
//
// @param request - PremiumGetTodoTasksRequest
//
// @param headers - PremiumGetTodoTasksHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumGetTodoTasksResponse
func (client *Client) PremiumGetTodoTasksWithOptions(request *PremiumGetTodoTasksRequest, headers *PremiumGetTodoTasksHeaders, runtime *util.RuntimeOptions) (_result *PremiumGetTodoTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateBefore)) {
		query["createBefore"] = request.CreateBefore
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
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
		Action:      tea.String("PremiumGetTodoTasks"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/processCentres/todoTasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumGetTodoTasksResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询审批中心待处理任务列表(OA高级版专享接口)
//
// @param request - PremiumGetTodoTasksRequest
//
// @return PremiumGetTodoTasksResponse
func (client *Client) PremiumGetTodoTasks(request *PremiumGetTodoTasksRequest) (_result *PremiumGetTodoTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumGetTodoTasksHeaders{}
	_result = &PremiumGetTodoTasksResponse{}
	_body, _err := client.PremiumGetTodoTasksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 下载审批附件(OA高级版专享)
//
// @param request - PremiumGrantProcessInstanceForDownloadFileRequest
//
// @param headers - PremiumGrantProcessInstanceForDownloadFileHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumGrantProcessInstanceForDownloadFileResponse
func (client *Client) PremiumGrantProcessInstanceForDownloadFileWithOptions(request *PremiumGrantProcessInstanceForDownloadFileRequest, headers *PremiumGrantProcessInstanceForDownloadFileHeaders, runtime *util.RuntimeOptions) (_result *PremiumGrantProcessInstanceForDownloadFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["fileId"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		body["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.WithCommentAttatchment)) {
		body["withCommentAttatchment"] = request.WithCommentAttatchment
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
		Action:      tea.String("PremiumGrantProcessInstanceForDownloadFile"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/processInstances/spaces/files/urls/download"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumGrantProcessInstanceForDownloadFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 下载审批附件(OA高级版专享)
//
// @param request - PremiumGrantProcessInstanceForDownloadFileRequest
//
// @return PremiumGrantProcessInstanceForDownloadFileResponse
func (client *Client) PremiumGrantProcessInstanceForDownloadFile(request *PremiumGrantProcessInstanceForDownloadFileRequest) (_result *PremiumGrantProcessInstanceForDownloadFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumGrantProcessInstanceForDownloadFileHeaders{}
	_result = &PremiumGrantProcessInstanceForDownloadFileResponse{}
	_body, _err := client.PremiumGrantProcessInstanceForDownloadFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建或更新分组(高级版专享接口)
//
// @param request - PremiumInsertOrUpdateDirRequest
//
// @param headers - PremiumInsertOrUpdateDirHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumInsertOrUpdateDirResponse
func (client *Client) PremiumInsertOrUpdateDirWithOptions(request *PremiumInsertOrUpdateDirRequest, headers *PremiumInsertOrUpdateDirHeaders, runtime *util.RuntimeOptions) (_result *PremiumInsertOrUpdateDirResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizGroup)) {
		body["bizGroup"] = request.BizGroup
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Name18n)) {
		body["name18n"] = request.Name18n
	}

	if !tea.BoolValue(util.IsUnset(request.OperateUserId)) {
		body["operateUserId"] = request.OperateUserId
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
		Action:      tea.String("PremiumInsertOrUpdateDir"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/processCentres/directories"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumInsertOrUpdateDirResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建或更新分组(高级版专享接口)
//
// @param request - PremiumInsertOrUpdateDirRequest
//
// @return PremiumInsertOrUpdateDirResponse
func (client *Client) PremiumInsertOrUpdateDir(request *PremiumInsertOrUpdateDirRequest) (_result *PremiumInsertOrUpdateDirResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumInsertOrUpdateDirHeaders{}
	_result = &PremiumInsertOrUpdateDirResponse{}
	_body, _err := client.PremiumInsertOrUpdateDirWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量获取审批模板信息（包含表单和流程配置信息）(高级版专享接口)
//
// @param request - PremiumQuerySchemaAndProcessByCodeListRequest
//
// @param headers - PremiumQuerySchemaAndProcessByCodeListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumQuerySchemaAndProcessByCodeListResponse
func (client *Client) PremiumQuerySchemaAndProcessByCodeListWithOptions(request *PremiumQuerySchemaAndProcessByCodeListRequest, headers *PremiumQuerySchemaAndProcessByCodeListHeaders, runtime *util.RuntimeOptions) (_result *PremiumQuerySchemaAndProcessByCodeListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProcessCodes)) {
		body["processCodes"] = request.ProcessCodes
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
		Action:      tea.String("PremiumQuerySchemaAndProcessByCodeList"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/processes/schemas/batchQuery"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumQuerySchemaAndProcessByCodeListResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量获取审批模板信息（包含表单和流程配置信息）(高级版专享接口)
//
// @param request - PremiumQuerySchemaAndProcessByCodeListRequest
//
// @return PremiumQuerySchemaAndProcessByCodeListResponse
func (client *Client) PremiumQuerySchemaAndProcessByCodeList(request *PremiumQuerySchemaAndProcessByCodeListRequest) (_result *PremiumQuerySchemaAndProcessByCodeListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumQuerySchemaAndProcessByCodeListHeaders{}
	_result = &PremiumQuerySchemaAndProcessByCodeListResponse{}
	_body, _err := client.PremiumQuerySchemaAndProcessByCodeListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 流程转交待处理任务查询(高级版专享接口)
//
// @param request - PremiumQueryTodoTasksByManagerRequest
//
// @param headers - PremiumQueryTodoTasksByManagerHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumQueryTodoTasksByManagerResponse
func (client *Client) PremiumQueryTodoTasksByManagerWithOptions(request *PremiumQueryTodoTasksByManagerRequest, headers *PremiumQueryTodoTasksByManagerHeaders, runtime *util.RuntimeOptions) (_result *PremiumQueryTodoTasksByManagerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionerUserId)) {
		query["actionerUserId"] = request.ActionerUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerUserId)) {
		query["managerUserId"] = request.ManagerUserId
	}

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
		Action:      tea.String("PremiumQueryTodoTasksByManager"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/tasks/todoTasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumQueryTodoTasksByManagerResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 流程转交待处理任务查询(高级版专享接口)
//
// @param request - PremiumQueryTodoTasksByManagerRequest
//
// @return PremiumQueryTodoTasksByManagerResponse
func (client *Client) PremiumQueryTodoTasksByManager(request *PremiumQueryTodoTasksByManagerRequest) (_result *PremiumQueryTodoTasksByManagerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumQueryTodoTasksByManagerHeaders{}
	_result = &PremiumQueryTodoTasksByManagerResponse{}
	_body, _err := client.PremiumQueryTodoTasksByManagerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量流程审批任务转交(高级版专享接口)
//
// @param request - PremiumRedirectTasksByManagerRequest
//
// @param headers - PremiumRedirectTasksByManagerHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumRedirectTasksByManagerResponse
func (client *Client) PremiumRedirectTasksByManagerWithOptions(request *PremiumRedirectTasksByManagerRequest, headers *PremiumRedirectTasksByManagerHeaders, runtime *util.RuntimeOptions) (_result *PremiumRedirectTasksByManagerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HandoverUserId)) {
		body["handoverUserId"] = request.HandoverUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerUserId)) {
		body["managerUserId"] = request.ManagerUserId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskIds)) {
		body["taskIds"] = request.TaskIds
	}

	if !tea.BoolValue(util.IsUnset(request.TransfereeUserId)) {
		body["transfereeUserId"] = request.TransfereeUserId
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
		Action:      tea.String("PremiumRedirectTasksByManager"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/tasks/batchRedirect"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumRedirectTasksByManagerResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量流程审批任务转交(高级版专享接口)
//
// @param request - PremiumRedirectTasksByManagerRequest
//
// @return PremiumRedirectTasksByManagerResponse
func (client *Client) PremiumRedirectTasksByManager(request *PremiumRedirectTasksByManagerRequest) (_result *PremiumRedirectTasksByManagerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumRedirectTasksByManagerHeaders{}
	_result = &PremiumRedirectTasksByManagerResponse{}
	_body, _err := client.PremiumRedirectTasksByManagerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 退回OA审批任务(OA高级版专享)
//
// @param request - PremiumRevertTaskRequest
//
// @param headers - PremiumRevertTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumRevertTaskResponse
func (client *Client) PremiumRevertTaskWithOptions(request *PremiumRevertTaskRequest, headers *PremiumRevertTaskHeaders, runtime *util.RuntimeOptions) (_result *PremiumRevertTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperateUserId)) {
		body["operateUserId"] = request.OperateUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		body["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.RevertAction)) {
		body["revertAction"] = request.RevertAction
	}

	if !tea.BoolValue(util.IsUnset(request.TargetActivityId)) {
		body["targetActivityId"] = request.TargetActivityId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
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
		Action:      tea.String("PremiumRevertTask"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/tasks/revert"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumRevertTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 退回OA审批任务(OA高级版专享)
//
// @param request - PremiumRevertTaskRequest
//
// @return PremiumRevertTaskResponse
func (client *Client) PremiumRevertTask(request *PremiumRevertTaskRequest) (_result *PremiumRevertTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumRevertTaskHeaders{}
	_result = &PremiumRevertTaskResponse{}
	_body, _err := client.PremiumRevertTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建或更新数据表单模板(OA高级版专享)
//
// @param request - PremiumSaveFormRequest
//
// @param headers - PremiumSaveFormHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumSaveFormResponse
func (client *Client) PremiumSaveFormWithOptions(request *PremiumSaveFormRequest, headers *PremiumSaveFormHeaders, runtime *util.RuntimeOptions) (_result *PremiumSaveFormResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FormComponents)) {
		body["formComponents"] = request.FormComponents
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		body["processCode"] = request.ProcessCode
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
		Action:      tea.String("PremiumSaveForm"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/dataForms/templates"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumSaveFormResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建或更新数据表单模板(OA高级版专享)
//
// @param request - PremiumSaveFormRequest
//
// @return PremiumSaveFormResponse
func (client *Client) PremiumSaveForm(request *PremiumSaveFormRequest) (_result *PremiumSaveFormResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumSaveFormHeaders{}
	_result = &PremiumSaveFormResponse{}
	_body, _err := client.PremiumSaveFormWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建数据表单实例(OA高级版专享)
//
// @param request - PremiumSaveFormInstanceRequest
//
// @param headers - PremiumSaveFormInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumSaveFormInstanceResponse
func (client *Client) PremiumSaveFormInstanceWithOptions(request *PremiumSaveFormInstanceRequest, headers *PremiumSaveFormInstanceHeaders, runtime *util.RuntimeOptions) (_result *PremiumSaveFormInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FormComponentValueList)) {
		body["formComponentValueList"] = request.FormComponentValueList
	}

	if !tea.BoolValue(util.IsUnset(request.OriginatorUserId)) {
		body["originatorUserId"] = request.OriginatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		body["processCode"] = request.ProcessCode
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
		Action:      tea.String("PremiumSaveFormInstance"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/dataForms/formInstances/save"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumSaveFormInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据表单实例(OA高级版专享)
//
// @param request - PremiumSaveFormInstanceRequest
//
// @return PremiumSaveFormInstanceResponse
func (client *Client) PremiumSaveFormInstance(request *PremiumSaveFormInstanceRequest) (_result *PremiumSaveFormInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumSaveFormInstanceHeaders{}
	_result = &PremiumSaveFormInstanceResponse{}
	_body, _err := client.PremiumSaveFormInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建或更新流程中心外部集成模板(高级版专享接口)
//
// @param request - PremiumSaveIntegratedProcessRequest
//
// @param headers - PremiumSaveIntegratedProcessHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumSaveIntegratedProcessResponse
func (client *Client) PremiumSaveIntegratedProcessWithOptions(request *PremiumSaveIntegratedProcessRequest, headers *PremiumSaveIntegratedProcessHeaders, runtime *util.RuntimeOptions) (_result *PremiumSaveIntegratedProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FormComponents)) {
		body["formComponents"] = request.FormComponents
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		body["processCode"] = request.ProcessCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessFeatureConfig)) {
		body["processFeatureConfig"] = request.ProcessFeatureConfig
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateConfig)) {
		body["templateConfig"] = request.TemplateConfig
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
		Action:      tea.String("PremiumSaveIntegratedProcess"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/processCentres/schemas"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumSaveIntegratedProcessResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建或更新流程中心外部集成模板(高级版专享接口)
//
// @param request - PremiumSaveIntegratedProcessRequest
//
// @return PremiumSaveIntegratedProcessResponse
func (client *Client) PremiumSaveIntegratedProcess(request *PremiumSaveIntegratedProcessRequest) (_result *PremiumSaveIntegratedProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumSaveIntegratedProcessHeaders{}
	_result = &PremiumSaveIntegratedProcessResponse{}
	_body, _err := client.PremiumSaveIntegratedProcessWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建流程中心外部集成实例(高级版专享接口)
//
// @param request - PremiumSaveIntegratedProcessInstanceRequest
//
// @param headers - PremiumSaveIntegratedProcessInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumSaveIntegratedProcessInstanceResponse
func (client *Client) PremiumSaveIntegratedProcessInstanceWithOptions(request *PremiumSaveIntegratedProcessInstanceRequest, headers *PremiumSaveIntegratedProcessInstanceHeaders, runtime *util.RuntimeOptions) (_result *PremiumSaveIntegratedProcessInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizData)) {
		body["bizData"] = request.BizData
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureConfig)) {
		body["featureConfig"] = request.FeatureConfig
	}

	if !tea.BoolValue(util.IsUnset(request.FormComponentValueList)) {
		body["formComponentValueList"] = request.FormComponentValueList
	}

	if !tea.BoolValue(util.IsUnset(request.Notifiers)) {
		body["notifiers"] = request.Notifiers
	}

	if !tea.BoolValue(util.IsUnset(request.OriginatorUserId)) {
		body["originatorUserId"] = request.OriginatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		body["processCode"] = request.ProcessCode
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["url"] = request.Url
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
		Action:      tea.String("PremiumSaveIntegratedProcessInstance"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/processCentres/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumSaveIntegratedProcessInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建流程中心外部集成实例(高级版专享接口)
//
// @param request - PremiumSaveIntegratedProcessInstanceRequest
//
// @return PremiumSaveIntegratedProcessInstanceResponse
func (client *Client) PremiumSaveIntegratedProcessInstance(request *PremiumSaveIntegratedProcessInstanceRequest) (_result *PremiumSaveIntegratedProcessInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumSaveIntegratedProcessInstanceHeaders{}
	_result = &PremiumSaveIntegratedProcessInstanceResponse{}
	_body, _err := client.PremiumSaveIntegratedProcessInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建流程中心外部集成待处理任务(高级版专享接口)
//
// @param request - PremiumSaveIntegratedTaskRequest
//
// @param headers - PremiumSaveIntegratedTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumSaveIntegratedTaskResponse
func (client *Client) PremiumSaveIntegratedTaskWithOptions(request *PremiumSaveIntegratedTaskRequest, headers *PremiumSaveIntegratedTaskHeaders, runtime *util.RuntimeOptions) (_result *PremiumSaveIntegratedTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActivityId)) {
		body["activityId"] = request.ActivityId
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureConfig)) {
		body["featureConfig"] = request.FeatureConfig
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		body["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskConfig)) {
		body["taskConfig"] = request.TaskConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Tasks)) {
		body["tasks"] = request.Tasks
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
		Action:      tea.String("PremiumSaveIntegratedTask"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/processCentres/tasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumSaveIntegratedTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建流程中心外部集成待处理任务(高级版专享接口)
//
// @param request - PremiumSaveIntegratedTaskRequest
//
// @return PremiumSaveIntegratedTaskResponse
func (client *Client) PremiumSaveIntegratedTask(request *PremiumSaveIntegratedTaskRequest) (_result *PremiumSaveIntegratedTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumSaveIntegratedTaskHeaders{}
	_result = &PremiumSaveIntegratedTaskResponse{}
	_body, _err := client.PremiumSaveIntegratedTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新数据表单实例(OA高级版专享)
//
// @param request - PremiumUpdateFormInstanceRequest
//
// @param headers - PremiumUpdateFormInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumUpdateFormInstanceResponse
func (client *Client) PremiumUpdateFormInstanceWithOptions(request *PremiumUpdateFormInstanceRequest, headers *PremiumUpdateFormInstanceHeaders, runtime *util.RuntimeOptions) (_result *PremiumUpdateFormInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FormComponentValueList)) {
		body["formComponentValueList"] = request.FormComponentValueList
	}

	if !tea.BoolValue(util.IsUnset(request.FormInstanceIds)) {
		body["formInstanceIds"] = request.FormInstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.OriginatorUserId)) {
		body["originatorUserId"] = request.OriginatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		body["processCode"] = request.ProcessCode
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
		Action:      tea.String("PremiumUpdateFormInstance"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/dataForms/formInstances"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumUpdateFormInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新数据表单实例(OA高级版专享)
//
// @param request - PremiumUpdateFormInstanceRequest
//
// @return PremiumUpdateFormInstanceResponse
func (client *Client) PremiumUpdateFormInstance(request *PremiumUpdateFormInstanceRequest) (_result *PremiumUpdateFormInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumUpdateFormInstanceHeaders{}
	_result = &PremiumUpdateFormInstanceResponse{}
	_body, _err := client.PremiumUpdateFormInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新审批实例(OA高级版专享)
//
// @param request - PremiumUpdateProcessInstanceVariablesRequest
//
// @param headers - PremiumUpdateProcessInstanceVariablesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PremiumUpdateProcessInstanceVariablesResponse
func (client *Client) PremiumUpdateProcessInstanceVariablesWithOptions(request *PremiumUpdateProcessInstanceVariablesRequest, headers *PremiumUpdateProcessInstanceVariablesHeaders, runtime *util.RuntimeOptions) (_result *PremiumUpdateProcessInstanceVariablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpUserId)) {
		body["opUserId"] = request.OpUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		body["processCode"] = request.ProcessCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		body["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Variables)) {
		body["variables"] = request.Variables
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
		Action:      tea.String("PremiumUpdateProcessInstanceVariables"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/premium/processInstances"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &PremiumUpdateProcessInstanceVariablesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新审批实例(OA高级版专享)
//
// @param request - PremiumUpdateProcessInstanceVariablesRequest
//
// @return PremiumUpdateProcessInstanceVariablesResponse
func (client *Client) PremiumUpdateProcessInstanceVariables(request *PremiumUpdateProcessInstanceVariablesRequest) (_result *PremiumUpdateProcessInstanceVariablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PremiumUpdateProcessInstanceVariablesHeaders{}
	_result = &PremiumUpdateProcessInstanceVariablesResponse{}
	_body, _err := client.PremiumUpdateProcessInstanceVariablesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 审批流程预测
//
// @param request - ProcessForecastRequest
//
// @param headers - ProcessForecastHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ProcessForecastResponse
func (client *Client) ProcessForecastWithOptions(request *ProcessForecastRequest, headers *ProcessForecastHeaders, runtime *util.RuntimeOptions) (_result *ProcessForecastResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.FormComponentValues)) {
		body["formComponentValues"] = request.FormComponentValues
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		body["processCode"] = request.ProcessCode
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
		Action:      tea.String("ProcessForecast"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processes/forecast"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &ProcessForecastResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 审批流程预测
//
// @param request - ProcessForecastRequest
//
// @return ProcessForecastResponse
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

// Summary:
//
// 根据processCode分页获取表单数据
//
// @param request - QueryAllFormInstancesRequest
//
// @param headers - QueryAllFormInstancesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAllFormInstancesResponse
func (client *Client) QueryAllFormInstancesWithOptions(request *QueryAllFormInstancesRequest, headers *QueryAllFormInstancesHeaders, runtime *util.RuntimeOptions) (_result *QueryAllFormInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUuid)) {
		query["appUuid"] = request.AppUuid
	}

	if !tea.BoolValue(util.IsUnset(request.FormCode)) {
		query["formCode"] = request.FormCode
	}

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
		Action:      tea.String("QueryAllFormInstances"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/forms/pages/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAllFormInstancesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据processCode分页获取表单数据
//
// @param request - QueryAllFormInstancesRequest
//
// @return QueryAllFormInstancesResponse
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

// Summary:
//
// 批量查询审批流程实例
//
// @param request - QueryAllProcessInstancesRequest
//
// @param headers - QueryAllProcessInstancesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAllProcessInstancesResponse
func (client *Client) QueryAllProcessInstancesWithOptions(request *QueryAllProcessInstancesRequest, headers *QueryAllProcessInstancesHeaders, runtime *util.RuntimeOptions) (_result *QueryAllProcessInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUuid)) {
		query["appUuid"] = request.AppUuid
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimeInMills)) {
		query["endTimeInMills"] = request.EndTimeInMills
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		query["processCode"] = request.ProcessCode
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimeInMills)) {
		query["startTimeInMills"] = request.StartTimeInMills
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
		Action:      tea.String("QueryAllProcessInstances"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processes/pages/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAllProcessInstancesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询审批流程实例
//
// @param request - QueryAllProcessInstancesRequest
//
// @return QueryAllProcessInstancesResponse
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

// Summary:
//
// 根据业务标识查询表单描述信息
//
// @param request - QueryFormByBizTypeRequest
//
// @param headers - QueryFormByBizTypeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFormByBizTypeResponse
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
		realHeaders["x-acs-dingtalk-access-token"] = util.ToJSONString(headers.XAcsDingtalkAccessToken)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryFormByBizType"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/forms/forminfos/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryFormByBizTypeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据业务标识查询表单描述信息
//
// @param request - QueryFormByBizTypeRequest
//
// @return QueryFormByBizTypeResponse
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

// Summary:
//
// 查询数据表单
//
// @param request - QueryFormInstanceRequest
//
// @param headers - QueryFormInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryFormInstanceResponse
func (client *Client) QueryFormInstanceWithOptions(request *QueryFormInstanceRequest, headers *QueryFormInstanceHeaders, runtime *util.RuntimeOptions) (_result *QueryFormInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUuid)) {
		query["appUuid"] = request.AppUuid
	}

	if !tea.BoolValue(util.IsUnset(request.FormCode)) {
		query["formCode"] = request.FormCode
	}

	if !tea.BoolValue(util.IsUnset(request.FormInstanceId)) {
		query["formInstanceId"] = request.FormInstanceId
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
		Action:      tea.String("QueryFormInstance"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/forms/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryFormInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据表单
//
// @param request - QueryFormInstanceRequest
//
// @return QueryFormInstanceResponse
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

// Summary:
//
// 查询通过流程中心集成的OA审批任务
//
// @param request - QueryIntegratedTodoTaskRequest
//
// @param headers - QueryIntegratedTodoTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryIntegratedTodoTaskResponse
func (client *Client) QueryIntegratedTodoTaskWithOptions(request *QueryIntegratedTodoTaskRequest, headers *QueryIntegratedTodoTaskHeaders, runtime *util.RuntimeOptions) (_result *QueryIntegratedTodoTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateBefore)) {
		query["createBefore"] = request.CreateBefore
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
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
		Action:      tea.String("QueryIntegratedTodoTask"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processCentres/todoTasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryIntegratedTodoTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询通过流程中心集成的OA审批任务
//
// @param request - QueryIntegratedTodoTaskRequest
//
// @return QueryIntegratedTodoTaskResponse
func (client *Client) QueryIntegratedTodoTask(request *QueryIntegratedTodoTaskRequest) (_result *QueryIntegratedTodoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryIntegratedTodoTaskHeaders{}
	_result = &QueryIntegratedTodoTaskResponse{}
	_body, _err := client.QueryIntegratedTodoTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据业务标识查询模板
//
// @param request - QueryProcessByBizCategoryIdRequest
//
// @param headers - QueryProcessByBizCategoryIdHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryProcessByBizCategoryIdResponse
func (client *Client) QueryProcessByBizCategoryIdWithOptions(request *QueryProcessByBizCategoryIdRequest, headers *QueryProcessByBizCategoryIdHeaders, runtime *util.RuntimeOptions) (_result *QueryProcessByBizCategoryIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["bizType"] = request.BizType
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
		Action:      tea.String("QueryProcessByBizCategoryId"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processes/categories/templates"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryProcessByBizCategoryIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据业务标识查询模板
//
// @param request - QueryProcessByBizCategoryIdRequest
//
// @return QueryProcessByBizCategoryIdResponse
func (client *Client) QueryProcessByBizCategoryId(request *QueryProcessByBizCategoryIdRequest) (_result *QueryProcessByBizCategoryIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QueryProcessByBizCategoryIdHeaders{}
	_result = &QueryProcessByBizCategoryIdResponse{}
	_body, _err := client.QueryProcessByBizCategoryIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 蓝凌获取schema和流程信息
//
// @param request - QuerySchemaAndProcessRequest
//
// @param headers - QuerySchemaAndProcessHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySchemaAndProcessResponse
func (client *Client) QuerySchemaAndProcessWithOptions(request *QuerySchemaAndProcessRequest, headers *QuerySchemaAndProcessHeaders, runtime *util.RuntimeOptions) (_result *QuerySchemaAndProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUuid)) {
		query["appUuid"] = request.AppUuid
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		query["processCode"] = request.ProcessCode
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
		Action:      tea.String("QuerySchemaAndProcess"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/forms/schemaAndProcess"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySchemaAndProcessResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 蓝凌获取schema和流程信息
//
// @param request - QuerySchemaAndProcessRequest
//
// @return QuerySchemaAndProcessResponse
func (client *Client) QuerySchemaAndProcess(request *QuerySchemaAndProcessRequest) (_result *QuerySchemaAndProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &QuerySchemaAndProcessHeaders{}
	_result = &QuerySchemaAndProcessResponse{}
	_body, _err := client.QuerySchemaAndProcessWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过 processCode 获取表单 schema 信息
//
// @param request - QuerySchemaByProcessCodeRequest
//
// @param headers - QuerySchemaByProcessCodeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySchemaByProcessCodeResponse
func (client *Client) QuerySchemaByProcessCodeWithOptions(request *QuerySchemaByProcessCodeRequest, headers *QuerySchemaByProcessCodeHeaders, runtime *util.RuntimeOptions) (_result *QuerySchemaByProcessCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppUuid)) {
		query["appUuid"] = request.AppUuid
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		query["processCode"] = request.ProcessCode
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
		Action:      tea.String("QuerySchemaByProcessCode"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/forms/schemas/processCodes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySchemaByProcessCodeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过 processCode 获取表单 schema 信息
//
// @param request - QuerySchemaByProcessCodeRequest
//
// @return QuerySchemaByProcessCodeResponse
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

// Summary:
//
// 转交OA审批任务
//
// @param request - RedirectWorkflowTaskRequest
//
// @param headers - RedirectWorkflowTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RedirectWorkflowTaskResponse
func (client *Client) RedirectWorkflowTaskWithOptions(request *RedirectWorkflowTaskRequest, headers *RedirectWorkflowTaskHeaders, runtime *util.RuntimeOptions) (_result *RedirectWorkflowTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionName)) {
		body["actionName"] = request.ActionName
	}

	if !tea.BoolValue(util.IsUnset(request.File)) {
		body["file"] = request.File
	}

	if !tea.BoolValue(util.IsUnset(request.OperateUserId)) {
		body["operateUserId"] = request.OperateUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.ToUserId)) {
		body["toUserId"] = request.ToUserId
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
		Action:      tea.String("RedirectWorkflowTask"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/tasks/redirect"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &RedirectWorkflowTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 转交OA审批任务
//
// @param request - RedirectWorkflowTaskRequest
//
// @return RedirectWorkflowTaskResponse
func (client *Client) RedirectWorkflowTask(request *RedirectWorkflowTaskRequest) (_result *RedirectWorkflowTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &RedirectWorkflowTaskHeaders{}
	_result = &RedirectWorkflowTaskResponse{}
	_body, _err := client.RedirectWorkflowTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建实例
//
// @param request - SaveIntegratedInstanceRequest
//
// @param headers - SaveIntegratedInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveIntegratedInstanceResponse
func (client *Client) SaveIntegratedInstanceWithOptions(request *SaveIntegratedInstanceRequest, headers *SaveIntegratedInstanceHeaders, runtime *util.RuntimeOptions) (_result *SaveIntegratedInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizData)) {
		body["bizData"] = request.BizData
	}

	if !tea.BoolValue(util.IsUnset(request.FeatureConfig)) {
		body["featureConfig"] = request.FeatureConfig
	}

	if !tea.BoolValue(util.IsUnset(request.FormComponentValueList)) {
		body["formComponentValueList"] = request.FormComponentValueList
	}

	if !tea.BoolValue(util.IsUnset(request.Notifiers)) {
		body["notifiers"] = request.Notifiers
	}

	if !tea.BoolValue(util.IsUnset(request.OriginatorUserId)) {
		body["originatorUserId"] = request.OriginatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		body["processCode"] = request.ProcessCode
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["url"] = request.Url
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
		Action:      tea.String("SaveIntegratedInstance"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processCentres/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveIntegratedInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建实例
//
// @param request - SaveIntegratedInstanceRequest
//
// @return SaveIntegratedInstanceResponse
func (client *Client) SaveIntegratedInstance(request *SaveIntegratedInstanceRequest) (_result *SaveIntegratedInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveIntegratedInstanceHeaders{}
	_result = &SaveIntegratedInstanceResponse{}
	_body, _err := client.SaveIntegratedInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建或更新审批模板
//
// @param request - SaveProcessRequest
//
// @param headers - SaveProcessHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveProcessResponse
func (client *Client) SaveProcessWithOptions(request *SaveProcessRequest, headers *SaveProcessHeaders, runtime *util.RuntimeOptions) (_result *SaveProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FormComponents)) {
		body["formComponents"] = request.FormComponents
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		body["processCode"] = request.ProcessCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessFeatureConfig)) {
		body["processFeatureConfig"] = request.ProcessFeatureConfig
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateConfig)) {
		body["templateConfig"] = request.TemplateConfig
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
		Action:      tea.String("SaveProcess"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processCentres/schemas"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveProcessResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建或更新审批模板
//
// @param request - SaveProcessRequest
//
// @return SaveProcessResponse
func (client *Client) SaveProcess(request *SaveProcessRequest) (_result *SaveProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveProcessHeaders{}
	_result = &SaveProcessResponse{}
	_body, _err := client.SaveProcessWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建审批实例
//
// @param request - StartProcessInstanceRequest
//
// @param headers - StartProcessInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartProcessInstanceResponse
func (client *Client) StartProcessInstanceWithOptions(request *StartProcessInstanceRequest, headers *StartProcessInstanceHeaders, runtime *util.RuntimeOptions) (_result *StartProcessInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Approvers)) {
		body["approvers"] = request.Approvers
	}

	if !tea.BoolValue(util.IsUnset(request.BizDetailPageUrl)) {
		body["bizDetailPageUrl"] = request.BizDetailPageUrl
	}

	if !tea.BoolValue(util.IsUnset(request.CcList)) {
		body["ccList"] = request.CcList
	}

	if !tea.BoolValue(util.IsUnset(request.CcPosition)) {
		body["ccPosition"] = request.CcPosition
	}

	if !tea.BoolValue(util.IsUnset(request.DeptId)) {
		body["deptId"] = request.DeptId
	}

	if !tea.BoolValue(util.IsUnset(request.FormComponentValues)) {
		body["formComponentValues"] = request.FormComponentValues
	}

	if !tea.BoolValue(util.IsUnset(request.MicroappAgentId)) {
		body["microappAgentId"] = request.MicroappAgentId
	}

	if !tea.BoolValue(util.IsUnset(request.OriginatorUserId)) {
		body["originatorUserId"] = request.OriginatorUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessCode)) {
		body["processCode"] = request.ProcessCode
	}

	if !tea.BoolValue(util.IsUnset(request.TargetSelectActioners)) {
		body["targetSelectActioners"] = request.TargetSelectActioners
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
		Action:      tea.String("StartProcessInstance"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processInstances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &StartProcessInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建审批实例
//
// @param request - StartProcessInstanceRequest
//
// @return StartProcessInstanceResponse
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

// Summary:
//
// 撤销审批实例
//
// @param request - TerminateProcessInstanceRequest
//
// @param headers - TerminateProcessInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TerminateProcessInstanceResponse
func (client *Client) TerminateProcessInstanceWithOptions(request *TerminateProcessInstanceRequest, headers *TerminateProcessInstanceHeaders, runtime *util.RuntimeOptions) (_result *TerminateProcessInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsSystem)) {
		body["isSystem"] = request.IsSystem
	}

	if !tea.BoolValue(util.IsUnset(request.OperatingUserId)) {
		body["operatingUserId"] = request.OperatingUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		body["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
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
		Action:      tea.String("TerminateProcessInstance"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processInstances/terminate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TerminateProcessInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 撤销审批实例
//
// @param request - TerminateProcessInstanceRequest
//
// @return TerminateProcessInstanceResponse
func (client *Client) TerminateProcessInstance(request *TerminateProcessInstanceRequest) (_result *TerminateProcessInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TerminateProcessInstanceHeaders{}
	_result = &TerminateProcessInstanceResponse{}
	_body, _err := client.TerminateProcessInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 流程转交待处理任务查询
//
// @param request - TodoTasksRequest
//
// @param headers - TodoTasksHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TodoTasksResponse
func (client *Client) TodoTasksWithOptions(request *TodoTasksRequest, headers *TodoTasksHeaders, runtime *util.RuntimeOptions) (_result *TodoTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionerUserId)) {
		query["actionerUserId"] = request.ActionerUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerUserId)) {
		query["managerUserId"] = request.ManagerUserId
	}

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
		Action:      tea.String("TodoTasks"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/tasks/todoTasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &TodoTasksResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 流程转交待处理任务查询
//
// @param request - TodoTasksRequest
//
// @return TodoTasksResponse
func (client *Client) TodoTasks(request *TodoTasksRequest) (_result *TodoTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &TodoTasksHeaders{}
	_result = &TodoTasksResponse{}
	_body, _err := client.TodoTasksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新流程中心任务状态
//
// @param request - UpdateIntegratedTaskRequest
//
// @param headers - UpdateIntegratedTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIntegratedTaskResponse
func (client *Client) UpdateIntegratedTaskWithOptions(request *UpdateIntegratedTaskRequest, headers *UpdateIntegratedTaskHeaders, runtime *util.RuntimeOptions) (_result *UpdateIntegratedTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		body["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Tasks)) {
		body["tasks"] = request.Tasks
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
		Action:      tea.String("UpdateIntegratedTask"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processCentres/tasks"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateIntegratedTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新流程中心任务状态
//
// @param request - UpdateIntegratedTaskRequest
//
// @return UpdateIntegratedTaskResponse
func (client *Client) UpdateIntegratedTask(request *UpdateIntegratedTaskRequest) (_result *UpdateIntegratedTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateIntegratedTaskHeaders{}
	_result = &UpdateIntegratedTaskResponse{}
	_body, _err := client.UpdateIntegratedTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新实例状态
//
// @param request - UpdateProcessInstanceRequest
//
// @param headers - UpdateProcessInstanceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProcessInstanceResponse
func (client *Client) UpdateProcessInstanceWithOptions(request *UpdateProcessInstanceRequest, headers *UpdateProcessInstanceHeaders, runtime *util.RuntimeOptions) (_result *UpdateProcessInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Notifiers)) {
		body["notifiers"] = request.Notifiers
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessInstanceId)) {
		body["processInstanceId"] = request.ProcessInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Result)) {
		body["result"] = request.Result
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
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
		Action:      tea.String("UpdateProcessInstance"),
		Version:     tea.String("workflow_1.0"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/v1.0/workflow/processCentres/instances"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("none"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProcessInstanceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新实例状态
//
// @param request - UpdateProcessInstanceRequest
//
// @return UpdateProcessInstanceResponse
func (client *Client) UpdateProcessInstance(request *UpdateProcessInstanceRequest) (_result *UpdateProcessInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateProcessInstanceHeaders{}
	_result = &UpdateProcessInstanceResponse{}
	_body, _err := client.UpdateProcessInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
